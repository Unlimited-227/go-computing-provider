package computing

import (
	"context"
	"encoding/json"
	stErr "errors"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/lagrangedao/go-computing-provider/build"
	"github.com/lagrangedao/go-computing-provider/conf"
	"github.com/lagrangedao/go-computing-provider/constants"
	"github.com/lagrangedao/go-computing-provider/internal/models"
	"github.com/lagrangedao/go-computing-provider/util"
	"io"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func GetServiceProviderInfo(c *gin.Context) {
	info := new(models.HostInfo)
	info.SwanProviderVersion = build.UserVersion()
	info.OperatingSystem = runtime.GOOS
	info.Architecture = runtime.GOARCH
	info.CPUCores = runtime.NumCPU()
	c.JSON(http.StatusOK, util.CreateSuccessResponse(info))
}

func ReceiveJob(c *gin.Context) {
	var jobData models.JobData
	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.GetLogger().Infof("Job received Data: %+v", jobData)

	available, err := checkResourceAvailable(jobData.JobSourceURI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
	}

	if !available {
		c.JSON(http.StatusOK, util.CreateErrorResponse(util.CheckAvailableResources))
	}

	var hostName string
	var logHost string
	prefixStr := generateString(10)
	if strings.HasPrefix(conf.GetConfig().API.Domain, ".") {
		hostName = prefixStr + conf.GetConfig().API.Domain
		logHost = "log" + conf.GetConfig().API.Domain
	} else {
		hostName = strings.Join([]string{prefixStr, conf.GetConfig().API.Domain}, ".")
		logHost = "log." + conf.GetConfig().API.Domain
	}

	if _, err = celeryService.DelayTask(constants.TASK_DEPLOY, jobData.JobSourceURI, hostName, jobData.Duration, jobData.UUID, jobData.TaskUUID); err != nil {
		logs.GetLogger().Errorf("Failed sync delpoy task, error: %v", err)
		return
	}

	jobData.JobResultURI = fmt.Sprintf("https://%s", hostName)

	multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
	jobSourceUri := jobData.JobSourceURI
	spaceUuid := jobSourceUri[strings.LastIndex(jobSourceUri, "/")+1:]
	wsUrl := fmt.Sprintf("wss://%s:%s/api/v1/computing/lagrange/spaces/log?space_id=%s", logHost, multiAddressSplit[4], spaceUuid)
	jobData.BuildLog = wsUrl + "&type=build"
	jobData.ContainerLog = wsUrl + "&type=container"

	if err = submitJob(&jobData); err != nil {
		jobData.JobResultURI = ""
	}
	logs.GetLogger().Infof("submit job detail: %+v", jobData)
	c.JSON(http.StatusOK, jobData)
}

func submitJob(jobData *models.JobData) error {
	logs.GetLogger().Printf("submitting job...")
	oldMask := syscall.Umask(0)
	defer syscall.Umask(oldMask)

	fileCachePath := conf.GetConfig().MCS.FileCachePath
	folderPath := "jobs"
	jobDetailFile := filepath.Join(folderPath, uuid.NewString()+".json")
	os.MkdirAll(filepath.Join(fileCachePath, folderPath), os.ModePerm)
	taskDetailFilePath := filepath.Join(fileCachePath, jobDetailFile)

	jobData.Status = constants.BiddingSubmitted
	jobData.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	bytes, err := json.Marshal(jobData)
	if err != nil {
		logs.GetLogger().Errorf("Failed Marshal JobData, error: %v", err)
		return err
	}
	if err = os.WriteFile(taskDetailFilePath, bytes, os.ModePerm); err != nil {
		logs.GetLogger().Errorf("Failed jobData write to file, error: %v", err)
		return err
	}

	storageService := NewStorageService()
	mcsOssFile, err := storageService.UploadFileToBucket(jobDetailFile, taskDetailFilePath, true)
	if err != nil {
		logs.GetLogger().Errorf("Failed upload file to bucket, error: %v", err)
		return err
	}
	logs.GetLogger().Infof("jobuuid: %s successfully submitted to IPFS", jobData.UUID)

	gatewayUrl, err := storageService.GetGatewayUrl()
	if err != nil {
		logs.GetLogger().Errorf("Failed get mcs ipfs gatewayUrl, error: %v", err)
		return err
	}
	jobData.JobResultURI = *gatewayUrl + "/ipfs/" + mcsOssFile.PayloadCid
	return nil
}

func RedeployJob(c *gin.Context) {
	var jobData models.JobData

	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.GetLogger().Infof("redeploy Job received: %+v", jobData)

	var hostName string
	if jobData.JobResultURI != "" {
		resp, err := http.Get(jobData.JobResultURI)
		if err != nil {
			logs.GetLogger().Errorf("error making request to Space API: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				logs.GetLogger().Errorf("error closed resp Space API: %+v", err)
			}
		}(resp.Body)
		logs.GetLogger().Infof("Space API response received. Response: %d", resp.StatusCode)
		if resp.StatusCode != http.StatusOK {
			logs.GetLogger().Errorf("space API response not OK. Status Code: %d", resp.StatusCode)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		var hostInfo struct {
			JobResultUri string `json:"job_result_uri"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&hostInfo); err != nil {
			logs.GetLogger().Errorf("error decoding Space API response JSON: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		hostName = strings.ReplaceAll(hostInfo.JobResultUri, "https://", "")
	} else {
		hostName = generateString(10) + conf.GetConfig().API.Domain
	}

	delayTask, err := celeryService.DelayTask(constants.TASK_DEPLOY, jobData.JobResultURI, hostName, jobData.Duration, jobData.UUID, jobData.TaskUUID)
	if err != nil {
		logs.GetLogger().Errorf("Failed sync delpoy task, error: %v", err)
		return
	}
	logs.GetLogger().Infof("delayTask detail info: %+v", delayTask)

	go func() {
		result, err := delayTask.Get(180 * time.Second)
		if err != nil {
			logs.GetLogger().Errorf("Failed get sync task result, error: %v", err)
			return
		}
		logs.GetLogger().Infof("Job: %s, service running successfully, job_result_url: %s", jobData.JobResultURI, result.(string))
	}()

	jobData.JobResultURI = fmt.Sprintf("https://%s", hostName)
	if err = submitJob(&jobData); err != nil {
		jobData.JobResultURI = ""
	}
	c.JSON(http.StatusOK, jobData)
}

func ReNewJob(c *gin.Context) {
	var jobData struct {
		TaskUuid string `json:"task_uuid"`
		Duration int    `json:"duration"`
	}

	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.GetLogger().Infof("renew Job received: %+v", jobData)

	if strings.TrimSpace(jobData.TaskUuid) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: task_uuid"})
		return
	}

	if jobData.Duration == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: duration"})
		return
	}

	conn := redisPool.Get()
	prefix := constants.REDIS_FULL_PREFIX + "*"
	keys, err := redis.Strings(conn.Do("KEYS", prefix))
	if err != nil {
		logs.GetLogger().Errorf("Failed get redis %s prefix, error: %+v", prefix, err)
		return
	}

	var spaceDetail models.CacheSpaceDetail
	for _, key := range keys {
		jobMetadata, err := RetrieveJobMetadata(key)
		if err != nil {
			logs.GetLogger().Errorf("Failed get redis key data, key: %s, error: %+v", key, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "query data failed"})
			return
		}
		if strings.EqualFold(jobMetadata.TaskUuid, jobData.TaskUuid) {
			spaceDetail = jobMetadata
			break
		}
	}

	redisKey := constants.REDIS_FULL_PREFIX + spaceDetail.SpaceUuid
	leftTime := spaceDetail.ExpireTime - time.Now().Unix()
	if leftTime < 0 {
		c.JSON(http.StatusOK, map[string]string{
			"status":  "failed",
			"message": "The job was terminated due to its expiration date",
		})
		return
	} else {
		fullArgs := []interface{}{redisKey}
		fields := map[string]string{
			"wallet_address": spaceDetail.WalletAddress,
			"space_name":     spaceDetail.SpaceName,
			"expire_time":    strconv.Itoa(int(time.Now().Unix()) + int(leftTime) + jobData.Duration),
			"space_uuid":     spaceDetail.SpaceUuid,
			"job_uuid":       spaceDetail.JobUuid,
			"task_type":      spaceDetail.TaskType,
			"deploy_name":    spaceDetail.DeployName,
			"hardware":       spaceDetail.Hardware,
		}

		for key, val := range fields {
			fullArgs = append(fullArgs, key, val)
		}
		redisConn := redisPool.Get()
		defer redisConn.Close()

		redisConn.Do("HSET", fullArgs...)
		redisConn.Do("SET", spaceDetail.SpaceUuid, "wait-delete", "EX", int(leftTime)+jobData.Duration)
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func CancelJob(c *gin.Context) {
	taskUuid := c.Query("task_uuid")
	if taskUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_uuid is required"})
		return
	}

	creatorWallet := c.Query("creator_wallet")
	if creatorWallet == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "creator_wallet is required"})
		return
	}

	conn := redisPool.Get()
	prefix := constants.REDIS_FULL_PREFIX + "*"
	keys, err := redis.Strings(conn.Do("KEYS", prefix))
	if err != nil {
		logs.GetLogger().Errorf("Failed get redis %s prefix, error: %+v", prefix, err)
		return
	}

	var jobDetail models.CacheSpaceDetail
	for _, key := range keys {
		jobMetadata, err := RetrieveJobMetadata(key)
		if err != nil {
			logs.GetLogger().Errorf("Failed get redis key data, key: %s, error: %+v", key, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "query data failed"})
			return
		}
		if strings.EqualFold(jobMetadata.TaskUuid, taskUuid) {
			jobDetail = jobMetadata
			break
		}
	}

	if jobDetail.WalletAddress == "" {
		c.JSON(http.StatusOK, util.CreateSuccessResponse("deleted success"))
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task_uuid: %s, delete space request failed, error: %+v", taskUuid, err)
				return
			}
		}()
		k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(jobDetail.WalletAddress)
		deleteJob(k8sNameSpace, jobDetail.SpaceUuid)
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("deleted success"))
}

func StatisticalSources(c *gin.Context) {
	location, err := getLocation()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed get location info"})
		return
	}

	k8sService := NewK8sService()
	statisticalSources, err := k8sService.StatisticalSources(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.ClusterResource{
		Region:      location,
		ClusterInfo: statisticalSources,
	})
}

func GetSpaceLog(c *gin.Context) {
	spaceUuid := c.Query("space_id")
	logType := c.Query("type")
	if strings.TrimSpace(spaceUuid) == "" {
		logs.GetLogger().Errorf("get space log failed, space_id is empty: %s", spaceUuid)
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: space_id"})
		return
	}

	if strings.TrimSpace(logType) == "" {
		logs.GetLogger().Errorf("get space log failed, type is empty: %s", logType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: type"})
		return
	}

	if strings.TrimSpace(logType) != "build" && strings.TrimSpace(logType) != "container" {
		logs.GetLogger().Errorf("get space log failed, type is build or container, type:: %s", logType)
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: type"})
		return
	}

	redisKey := constants.REDIS_FULL_PREFIX + spaceUuid
	spaceDetail, err := RetrieveJobMetadata(redisKey)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query data failed"})
		return
	}

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logs.GetLogger().Errorf("upgrading connection failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "upgrading connection failed"})
		return
	}
	handleConnection(conn, spaceDetail, logType)
}

func DoProof(c *gin.Context) {
	var proofTask struct {
		Method    string `json:"method"`
		BlockData string `json:"block_data"`
		Exp       int64  `json:"exp"`
	}
	if err := c.ShouldBindJSON(&proofTask); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("do proof task received: %+v", proofTask)

	if strings.TrimSpace(proofTask.Method) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.ProofParamError, "missing required field: method"))
		return
	}
	if proofTask.Method != "mine" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.ProofParamError, "method must be mine"))
		return
	}
	if proofTask.Exp < 0 || proofTask.Exp > 250 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.ProofParamError, "exp range is [0~250]"))
		return
	}

	k8sService := NewK8sService()
	job := &batchv1.Job{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "proof-job-" + generateString(5),
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "worker-container-" + generateString(5),
							Image: "filswan/worker-proof:v1.0",
							Env: []v1.EnvVar{
								{
									Name:  "METHOD",
									Value: proofTask.Method,
								},
								{
									Name:  "BLOCK_DATA",
									Value: proofTask.BlockData,
								},
								{
									Name:  "EXP",
									Value: strconv.Itoa(int(proofTask.Exp)),
								},
							},
						},
					},
					RestartPolicy: "Never",
				},
			},
			BackoffLimit:            new(int32),
			TTLSecondsAfterFinished: new(int32),
		},
	}

	*job.Spec.BackoffLimit = 1
	*job.Spec.TTLSecondsAfterFinished = 30

	createdJob, err := k8sService.k8sClient.BatchV1().Jobs(metaV1.NamespaceDefault).Create(context.TODO(), job, metaV1.CreateOptions{})
	if err != nil {
		logs.GetLogger().Errorf("Failed creating Pod: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofError))
		return
	}

	err = wait.PollImmediate(time.Second*3, time.Minute*5, func() (bool, error) {
		job, err := k8sService.k8sClient.BatchV1().Jobs(metaV1.NamespaceDefault).Get(context.Background(), createdJob.Name, metaV1.GetOptions{})
		if err != nil {
			logs.GetLogger().Errorf("Failed getting Job status: %v\n", err)
			return false, err
		}

		if job.Status.Succeeded > 0 {
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		logs.GetLogger().Errorf("Failed waiting for Job completion: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofError))
		return
	}

	podList, err := k8sService.k8sClient.CoreV1().Pods(metaV1.NamespaceDefault).List(context.Background(), metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("job-name=%s", createdJob.Name),
	})
	if err != nil {
		logs.GetLogger().Errorf("Error getting Pods for Job: %v\n", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofError))
		return
	}

	if len(podList.Items) == 0 {
		logs.GetLogger().Errorf("No Pods found for Job.")
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofError))
		return
	}

	podName := podList.Items[0].Name
	podLog, err := k8sService.k8sClient.CoreV1().Pods(metaV1.NamespaceDefault).GetLogs(podName, &v1.PodLogOptions{}).Stream(context.Background())
	if err != nil {
		logs.GetLogger().Errorf("Failed gettingPod logs: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofReadLogError))
		return
	}
	defer podLog.Close()

	bytes, err := io.ReadAll(podLog)
	if err != nil {
		logs.GetLogger().Errorf("Failed gettingPod logs: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofReadLogError))
		return
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse(string(bytes)))
}

func DoUbiProof(c *gin.Context) {
	var ubiTask struct {
		Wallet   string `json:"wallet"`
		NodeId   string `json:"node_id"`
		TaskUuid string `json:"task_uuid"`
		Task     string `json:"task"`
		TaskType int    `json:"task_type"`
		TaskName string `json:"task_name"`
	}

	if err := c.ShouldBindJSON(&ubiTask); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("do proof task received: %+v", ubiTask)

	if strings.TrimSpace(ubiTask.Wallet) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: wallet"))
		return
	}
	if strings.TrimSpace(ubiTask.NodeId) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: node_id"))
		return
	}

	if strings.TrimSpace(ubiTask.TaskUuid) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: task_uuid"))
		return
	}
	if strings.TrimSpace(ubiTask.Task) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "method must be task"))
		return
	}
	if ubiTask.TaskType != 0 && ubiTask.TaskType != 1 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "the value of task_type is 0 or 1"))
		return
	}
	if strings.TrimSpace(ubiTask.TaskName) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: task_name"))
		return
	}

	inputParamTaskJson, err := util.GetMcsFileByUrl(ubiTask.Task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "the value of task_type is 0 or 1"))
		return
	}

	go func() {
		//var envFilePath string
		//envFilePath = filepath.Join(os.Getenv("CP_PATH"), "fil-c2.env")
		//envVars, err := godotenv.Read(envFilePath)
		//if err != nil {
		//	logs.GetLogger().Errorf("reading fil-c2-env.env")
		//	return
		//}

		k8sService := NewK8sService()
		namespace := "ubi-task"
		filC2SecretName := ubiTask.TaskName
		if err = k8sService.CreateUbiTaskSecret(context.TODO(), namespace, filC2SecretName, inputParamTaskJson); err != nil {
			logs.GetLogger().Errorf("create ubi task configmap failed, error: %v", err)
			return
		}

		urlSplits := strings.Split(conf.GetConfig().API.MultiAddress, "/")
		receiveUrl := fmt.Sprintf("https://%s:%s/api/v1/computing/lagrange/cp/receive/ubi", urlSplits[2], urlSplits[4])

		execCommand := []string{"/bin/sh", "-c", "ubi-bench", "c2"}

		job := &batchv1.Job{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      "ubi-fil-c2-" + generateString(5),
				Namespace: namespace,
			},
			Spec: batchv1.JobSpec{
				Template: v1.PodTemplateSpec{
					Spec: v1.PodSpec{
						//NodeSelector: map[string]string{
						//
						//},
						Containers: []v1.Container{
							{
								Name:  "ubi-task-" + generateString(5),
								Image: "filswan/ubi-worker:v1.0",
								Env: []v1.EnvVar{
									{
										Name:  "RECEIVE_PROOF_URL",
										Value: receiveUrl,
									},
									{
										Name:  "TASK_UUID",
										Value: ubiTask.TaskUuid,
									},
								},
								VolumeMounts: []v1.VolumeMount{
									{
										Name:      "fil-c2-input-volume",
										MountPath: "/var/tmp/fil-c2-param",
									},
									{
										Name:      "proof-params",
										MountPath: "/var/tmp/filecoin-proof-parameters",
									},
								},
								Command: execCommand,
							},
						},
						Volumes: []v1.Volume{
							{
								Name: "proof-params",
								VolumeSource: v1.VolumeSource{
									HostPath: &v1.HostPathVolumeSource{
										Path: "/data/filecoin-proof-parameters",
									},
								},
							},
							{
								Name: "fil-c2-input-volume",
								VolumeSource: v1.VolumeSource{
									Secret: &v1.SecretVolumeSource{
										SecretName: filC2SecretName,
									},
								},
							},
						},
						RestartPolicy: "Never",
					},
				},
				BackoffLimit:            new(int32),
				TTLSecondsAfterFinished: new(int32),
			},
		}

		*job.Spec.BackoffLimit = 1
		*job.Spec.TTLSecondsAfterFinished = 3000

		if _, err = k8sService.k8sClient.BatchV1().Jobs(namespace).Create(context.TODO(), job, metaV1.CreateOptions{}); err != nil {
			logs.GetLogger().Errorf("Failed creating ubi task job: %v", err)
			return
		}
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func ReceiveProof(c *gin.Context) {
	var c2Proof models.Commit2Proof
	if err := c.ShouldBindJSON(&c2Proof); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("task_uuid: %s, C2 proof out received: %+v", c2Proof.TaskUuid, c2Proof)
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func handleConnection(conn *websocket.Conn, spaceDetail models.CacheSpaceDetail, logType string) {
	client := NewWsClient(conn)

	if logType == "build" {
		buildLogPath := filepath.Join("build", spaceDetail.WalletAddress, "spaces", spaceDetail.SpaceName, BuildFileName)
		if _, err := os.Stat(buildLogPath); err != nil {
			client.HandleLogs(strings.NewReader("This space is deployed starting from a image."))
		} else {
			logFile, _ := os.Open(buildLogPath)
			defer logFile.Close()
			client.HandleLogs(logFile)
		}
	} else if logType == "container" {
		k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(spaceDetail.WalletAddress)

		k8sService := NewK8sService()
		pods, err := k8sService.k8sClient.CoreV1().Pods(k8sNameSpace).List(context.TODO(), metaV1.ListOptions{
			LabelSelector: fmt.Sprintf("lad_app=%s", spaceDetail.SpaceUuid),
		})
		if err != nil {
			logs.GetLogger().Errorf("Error listing Pods: %v", err)
			return
		}

		if len(pods.Items) > 0 {
			line := int64(1000)
			containerStatuses := pods.Items[0].Status.ContainerStatuses
			lastIndex := len(containerStatuses) - 1
			req := k8sService.k8sClient.CoreV1().Pods(k8sNameSpace).GetLogs(pods.Items[0].Name, &v1.PodLogOptions{
				Container:  containerStatuses[lastIndex].Name,
				Follow:     true,
				Timestamps: true,
				TailLines:  &line,
			})

			podLogs, err := req.Stream(context.Background())
			if err != nil {
				logs.GetLogger().Errorf("Error opening log stream: %v", err)
				return
			}
			defer podLogs.Close()

			client.HandleLogs(podLogs)
		}
	}
}

func DeploySpaceTask(jobSourceURI, hostName string, duration int, jobUuid string, taskUuid string) string {
	updateJobStatus(jobUuid, models.JobUploadResult)

	var success bool
	var spaceUuid string
	var walletAddress string
	defer func() {
		if !success {
			k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(walletAddress)
			deleteJob(k8sNameSpace, spaceUuid)
		}

		if err := recover(); err != nil {
			logs.GetLogger().Errorf("deploy space task painc, error: %+v", err)
			return
		}
	}()

	spaceDetail, err := getSpaceDetail(jobSourceURI)
	if err != nil {
		logs.GetLogger().Errorln(err)
		return ""
	}

	walletAddress = spaceDetail.Data.Owner.PublicAddress
	spaceName := spaceDetail.Data.Space.Name
	spaceUuid = strings.ToLower(spaceDetail.Data.Space.Uuid)
	spaceHardware := spaceDetail.Data.Space.ActiveOrder.Config

	conn := redisPool.Get()
	fullArgs := []interface{}{constants.REDIS_FULL_PREFIX + spaceUuid}
	fields := map[string]string{
		"wallet_address": walletAddress,
		"space_name":     spaceName,
		"expire_time":    strconv.Itoa(int(time.Now().Unix()) + duration),
		"space_uuid":     spaceUuid,
		"task_uuid":      taskUuid,
	}

	for key, val := range fields {
		fullArgs = append(fullArgs, key, val)
	}
	_, _ = conn.Do("HSET", fullArgs...)

	logs.GetLogger().Infof("uuid: %s, spaceName: %s, hardwareName: %s", spaceUuid, spaceName, spaceHardware.Description)
	if len(spaceHardware.Description) == 0 {
		return ""
	}

	deploy := NewDeploy(jobUuid, hostName, walletAddress, spaceHardware.Description, int64(duration), taskUuid)
	deploy.WithSpaceInfo(spaceUuid, spaceName)

	spacePath := filepath.Join("build", walletAddress, "spaces", spaceName)
	os.RemoveAll(spacePath)
	updateJobStatus(jobUuid, models.JobDownloadSource)
	containsYaml, yamlPath, imagePath, modelsSettingFile, err := BuildSpaceTaskImage(spaceUuid, spaceDetail.Data.Files)
	if err != nil {
		logs.GetLogger().Error(err)
		return ""
	}

	deploy.WithSpacePath(imagePath)
	if len(modelsSettingFile) > 0 {
		err := deploy.WithModelSettingFile(modelsSettingFile).ModelInferenceToK8s()
		if err != nil {
			logs.GetLogger().Error(err)
			return ""
		}
		return hostName
	}

	if containsYaml {
		deploy.WithYamlInfo(yamlPath).YamlToK8s()
	} else {
		imageName, dockerfilePath := BuildImagesByDockerfile(jobUuid, spaceUuid, spaceName, imagePath)
		deploy.WithDockerfile(imageName, dockerfilePath).DockerfileToK8s()
	}
	success = true

	return hostName
}

func deleteJob(namespace, spaceUuid string) error {
	deployName := constants.K8S_DEPLOY_NAME_PREFIX + spaceUuid
	serviceName := constants.K8S_SERVICE_NAME_PREFIX + spaceUuid
	ingressName := constants.K8S_INGRESS_NAME_PREFIX + spaceUuid

	k8sService := NewK8sService()
	if err := k8sService.DeleteIngress(context.TODO(), namespace, ingressName); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete ingress, ingressName: %s, error: %+v", deployName, err)
		return err
	}
	logs.GetLogger().Infof("Deleted ingress %s finished", ingressName)

	if err := k8sService.DeleteService(context.TODO(), namespace, serviceName); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete service, serviceName: %s, error: %+v", serviceName, err)
		return err
	}
	logs.GetLogger().Infof("Deleted service %s finished", serviceName)

	dockerService := NewDockerService()
	deployImageIds, err := k8sService.GetDeploymentImages(context.TODO(), namespace, deployName)
	if err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed get deploy imageIds, deployName: %s, error: %+v", deployName, err)
		return err
	}
	for _, imageId := range deployImageIds {
		dockerService.RemoveImage(imageId)
	}

	if err := k8sService.DeleteDeployment(context.TODO(), namespace, deployName); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete deployment, deployName: %s, error: %+v", deployName, err)
		return err
	}
	time.Sleep(6 * time.Second)
	logs.GetLogger().Infof("Deleted deployment %s finished", deployName)

	if err := k8sService.DeleteDeployRs(context.TODO(), namespace, spaceUuid); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete ReplicaSetsController, spaceUuid: %s, error: %+v", spaceUuid, err)
		return err
	}

	if err := k8sService.DeletePod(context.TODO(), namespace, spaceUuid); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete pods, spaceUuid: %s, error: %+v", spaceUuid, err)
		return err
	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	var count = 0
	for {
		<-ticker.C
		count++
		if count >= 20 {
			break
		}
		getPods, err := k8sService.GetPods(namespace, spaceUuid)
		if err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed get pods form namespace, namepace: %s, error: %+v", namespace, err)
			continue
		}
		if !getPods {
			logs.GetLogger().Infof("Deleted all resource finised. spaceUuid: %s", spaceUuid)
			break
		}
	}
	return nil
}

func downloadModelUrl(namespace, spaceUuid, serviceIp string, podCmd []string) {
	k8sService := NewK8sService()
	podName, err := k8sService.WaitForPodRunning(namespace, spaceUuid, serviceIp)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	if err = k8sService.PodDoCommand(namespace, podName, "", podCmd); err != nil {
		logs.GetLogger().Error(err)
		return
	}
}

func updateJobStatus(jobUuid string, jobStatus models.JobStatus, url ...string) {
	go func() {
		if len(url) > 0 {
			deployingChan <- models.Job{
				Uuid:   jobUuid,
				Status: jobStatus,
				Count:  0,
				Url:    url[0],
			}
		} else {
			deployingChan <- models.Job{
				Uuid:   jobUuid,
				Status: jobStatus,
				Count:  0,
				Url:    "",
			}
		}
	}()
}

func getSpaceDetail(jobSourceURI string) (models.SpaceJSON, error) {
	resp, err := http.Get(jobSourceURI)
	if err != nil {
		return models.SpaceJSON{}, fmt.Errorf("error making request to Space API: %+v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return models.SpaceJSON{}, fmt.Errorf("space API response not OK. Status Code: %d", resp.StatusCode)
	}

	var spaceJson models.SpaceJSON
	if err := json.NewDecoder(resp.Body).Decode(&spaceJson); err != nil {
		return models.SpaceJSON{}, fmt.Errorf("error decoding Space API response JSON: %v", err)
	}
	return spaceJson, nil
}

func checkResourceAvailable(jobSourceURI string) (bool, error) {
	spaceDetail, err := getSpaceDetail(jobSourceURI)
	if err != nil {
		logs.GetLogger().Errorln(err)
		return false, err
	}

	taskType, hardwareDetail := getHardwareDetail(spaceDetail.Data.Space.ActiveOrder.Config.Description)
	k8sService := NewK8sService()

	activePods, err := k8sService.GetAllActivePod(context.TODO())
	if err != nil {
		return false, err
	}

	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return false, err
	}

	nodeGpuSummary, err := k8sService.GetNodeGpuSummary(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed collect k8s gpu, error: %+v", err)
		return false, err
	}

	for _, node := range nodes.Items {
		nodeGpu, remainderResource, _ := GetNodeResource(activePods, &node)
		remainderCpu := remainderResource[ResourceCpu]
		remainderMemory := float64(remainderResource[ResourceMem] / 1024 / 1024 / 1024)
		remainderStorage := float64(remainderResource[ResourceStorage] / 1024 / 1024 / 1024)

		if hardwareDetail.Cpu.Quantity < remainderCpu && float64(hardwareDetail.Memory.Quantity) < remainderMemory && float64(hardwareDetail.Storage.Quantity) < remainderStorage {
			if taskType == "CPU" {
				return true, nil
			} else if taskType == "GPU" {
				gpuName := strings.ReplaceAll(hardwareDetail.Gpu.Unit, " ", "-")
				logs.GetLogger().Infof("gpuName: %s, nodeGpu: %+v, nodeGpuSummary: %+v", gpuName, nodeGpu, nodeGpuSummary)
				usedCount, ok := nodeGpu[gpuName]
				if !ok {
					usedCount = 0
				}

				if usedCount+hardwareDetail.Gpu.Quantity <= nodeGpuSummary[node.Name][gpuName] {
					return true, nil
				}
				return true, nil
			}
		}
	}
	return false, nil
}

func generateString(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	source := characters + numbers
	result := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result[i] = source[rand.Intn(len(source))]
	}
	return string(result)
}

func getLocation() (string, error) {
	publicIpAddress, err := getLocalIPAddress()
	if err != nil {
		return "", err
	}
	logs.GetLogger().Infof("publicIpAddress: %s", publicIpAddress)

	resp, err := http.Get("http://ip-api.com/json/" + publicIpAddress)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ipInfo struct {
		Country     string `json:"country"`
		CountryCode string `json:"countryCode"`
		City        string `json:"city"`
		Region      string `json:"region"`
		RegionName  string `json:"regionName"`
	}
	if err = json.Unmarshal(body, &ipInfo); err != nil {
		return "", err
	}

	return strings.TrimSpace(ipInfo.RegionName) + "-" + ipInfo.CountryCode, nil
}

func getLocalIPAddress() (string, error) {
	req, err := http.NewRequest("GET", "https://ipapi.co/ip", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(ipBytes)), nil
}

var NotFoundRedisKey = stErr.New("not found redis key")

func RetrieveJobMetadata(key string) (models.CacheSpaceDetail, error) {
	redisConn := redisPool.Get()
	defer redisConn.Close()

	exist, err := redis.Int(redisConn.Do("EXISTS", key))
	if err != nil {
		return models.CacheSpaceDetail{}, err
	}
	if exist == 0 {
		return models.CacheSpaceDetail{}, NotFoundRedisKey
	}

	args := append([]interface{}{key}, "wallet_address", "space_name", "expire_time", "space_uuid", "job_uuid",
		"task_type", "deploy_name", "hardware", "url", "task_uuid")
	valuesStr, err := redis.Strings(redisConn.Do("HMGET", args...))
	if err != nil {
		logs.GetLogger().Errorf("Failed get redis key data, key: %s, error: %+v", key, err)
		return models.CacheSpaceDetail{}, err
	}

	var (
		walletAddress string
		spaceName     string
		expireTime    int64
		spaceUuid     string
		jobUuid       string
		taskType      string
		deployName    string
		hardware      string
		url           string
		taskUuid      string
	)

	if len(valuesStr) >= 3 {
		walletAddress = valuesStr[0]
		spaceName = valuesStr[1]
		expireTimeStr := valuesStr[2]
		spaceUuid = valuesStr[3]
		jobUuid = valuesStr[4]
		taskType = valuesStr[5]
		deployName = valuesStr[6]
		hardware = valuesStr[7]
		url = valuesStr[8]
		taskUuid = valuesStr[9]
		expireTime, err = strconv.ParseInt(strings.TrimSpace(expireTimeStr), 10, 64)
		if err != nil {
			logs.GetLogger().Errorf("Failed convert time str: [%s], error: %+v", expireTimeStr, err)
			return models.CacheSpaceDetail{}, err
		}
	}

	return models.CacheSpaceDetail{
		WalletAddress: walletAddress,
		SpaceName:     spaceName,
		SpaceUuid:     spaceUuid,
		ExpireTime:    expireTime,
		JobUuid:       jobUuid,
		TaskType:      taskType,
		DeployName:    deployName,
		Hardware:      hardware,
		Url:           url,
		TaskUuid:      taskUuid,
	}, nil
}
