package main

import (
	"github.com/lagrangedao/go-computing-provider/common"
	"strconv"
	"time"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/lagrangedao/go-computing-provider/conf"
	"github.com/lagrangedao/go-computing-provider/initializer"
	"github.com/lagrangedao/go-computing-provider/routers"
)

func main() {
	logs.GetLogger().Info("Start in computing provider mode.")
	initializer.ProjectInit()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		ValidateHeaders: false,
	}))
	pprof.Register(r)

	v1 := r.Group("/api/v1")
	routers.CPManager(v1.Group("/computing"))

	shutdownChan := make(chan struct{})
	httpStopper, err := common.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port))
	if err != nil {
		logs.GetLogger().Fatal("failed to start cp-api endpoint: %s", err)
	}

	finishCh := common.MonitorShutdown(shutdownChan,
		common.ShutdownHandler{Component: "cp-api", StopFunc: httpStopper},
	)

	<-finishCh
}
