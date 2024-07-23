// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ecp

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EcpSequencerMetaData contains all meta data concerning the EcpSequencer contract.
var EcpSequencerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cpAccounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchTransferredToEscrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferredToEscrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFromEscrow\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cpAccounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferToEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrowBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"getCPBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromEscrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261009860201b60201c565b6100a060201b60201c565b61003c3361016460201b60201c565b60018060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506103b5565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6101726101f360201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036101e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d890610329565b60405180910390fd5b6101f0816100a060201b60201c565b50565b61020161009860201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1661022561027d60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff161461027b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027290610395565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600082825260208201905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006103136026836102a6565b915061031e826102b7565b604082019050919050565b6000602082019050818103600083015261034281610306565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061037f6020836102a6565b915061038a82610349565b602082019050919050565b600060208201905081810360008301526103ae81610372565b9050919050565b611ead806103c46000396000f3fe6080604052600436106100ec5760003560e01c8063922f126a1161008a578063e11e1b0c11610059578063e11e1b0c146102e7578063f2fde38b14610312578063f340fa011461033b578063f3fef3a314610357576100f3565b8063922f126a1461022f5780639e88fd811461026c578063a6f9dae114610295578063be47afda146102be576100f3565b8063429b62e5116100c6578063429b62e51461018757806370480275146101c4578063715018a6146101ed5780638da5cb5b14610204576100f3565b80631485fef9146100f85780631785f53c1461012157806327e235e31461014a576100f3565b366100f357005b600080fd5b34801561010457600080fd5b5061011f600480360381019061011a91906113e5565b610380565b005b34801561012d57600080fd5b50610148600480360381019061014391906114c4565b61071f565b005b34801561015657600080fd5b50610171600480360381019061016c91906114c4565b6107c5565b60405161017e919061150a565b60405180910390f35b34801561019357600080fd5b506101ae60048036038101906101a991906114c4565b6107dd565b6040516101bb9190611540565b60405180910390f35b3480156101d057600080fd5b506101eb60048036038101906101e691906114c4565b6107fd565b005b3480156101f957600080fd5b506102026108a2565b005b34801561021057600080fd5b506102196108b6565b604051610226919061156a565b60405180910390f35b34801561023b57600080fd5b50610256600480360381019061025191906114c4565b6108df565b604051610263919061150a565b60405180910390f35b34801561027857600080fd5b50610293600480360381019061028e91906115b1565b610928565b005b3480156102a157600080fd5b506102bc60048036038101906102b791906114c4565b610b34565b005b3480156102ca57600080fd5b506102e560048036038101906102e091906115f1565b610c1e565b005b3480156102f357600080fd5b506102fc610d0c565b604051610309919061150a565b60405180910390f35b34801561031e57600080fd5b50610339600480360381019061033491906114c4565b610d12565b005b610355600480360381019061035091906114c4565b610d95565b005b34801561036357600080fd5b5061037e600480360381019061037991906115b1565b610eab565b005b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff168061040a57503373ffffffffffffffffffffffffffffffffffffffff166103f26108b6565b73ffffffffffffffffffffffffffffffffffffffff16145b610449576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104409061167b565b60405180910390fd5b818190508484905014610491576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610488906116e7565b60405180910390fd5b60005b848490508110156106c4578282828181106104b2576104b1611707565b5b90506020020135600260008787858181106104d0576104cf611707565b5b90506020020160208101906104e591906114c4565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610561576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055890611782565b60405180910390fd5b82828281811061057457610573611707565b5b905060200201356002600087878581811061059257610591611707565b5b90506020020160208101906105a791906114c4565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105f091906117d1565b9250508190555082828281811061060a57610609611707565b5b90506020020135600360008282546106229190611805565b9250508190555084848281811061063c5761063b611707565b5b905060200201602081019061065191906114c4565b73ffffffffffffffffffffffffffffffffffffffff167f8aa981417e2c3a0dadea5dd2eb5eab703d6d9c074ae83d377ce454dd6a93bfa784848481811061069b5761069a611707565b5b905060200201356040516106af919061150a565b60405180910390a28080600101915050610494565b503373ffffffffffffffffffffffffffffffffffffffff167f5a5b8c52e21fd816e73687ac900bbb3238b4e80ddcb10c216085993299e64ed5858585856040516107119493929190611977565b60405180910390a250505050565b6107276111d6565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f60405160405180910390a250565b60026020528060005260406000206000915090505481565b60016020528060005260406000206000915054906101000a900460ff1681565b6108056111d6565b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33960405160405180910390a250565b6108aa6111d6565b6108b46000611254565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16806109b257503373ffffffffffffffffffffffffffffffffffffffff1661099a6108b6565b73ffffffffffffffffffffffffffffffffffffffff16145b6109f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e89061167b565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610a73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6a90611782565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610ac291906117d1565b925050819055508060036000828254610adb9190611805565b925050819055508173ffffffffffffffffffffffffffffffffffffffff167f8aa981417e2c3a0dadea5dd2eb5eab703d6d9c074ae83d377ce454dd6a93bfa782604051610b28919061150a565b60405180910390a25050565b610b3c6111d6565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ba2906119fe565b60405180910390fd5b6000610bb56108b6565b9050610bc082610d12565b8173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c60405160405180910390a35050565b610c266111d6565b806003541015610c6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6290611a6a565b60405180910390fd5b8060036000828254610c7d91906117d1565b92505081905550610c8c6108b6565b73ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610cd1573d6000803e3d6000fd5b507fd1f01ada3f4d4c0d5d81675b458e3f5c4dda3e79a82d07b03459147644bc7b7b81604051610d01919061150a565b60405180910390a150565b60035481565b610d1a6111d6565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610d89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8090611afc565b60405180910390fd5b610d9281611254565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610e04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dfb90611b68565b60405180910390fd5b34600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e539190611805565b925050819055508073ffffffffffffffffffffffffffffffffffffffff167f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c434604051610ea0919061150a565b60405180910390a250565b6000808373ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610f559190611bf9565b6000604051808303816000865af19150503d8060008114610f92576040519150601f19603f3d011682016040523d82523d6000602084013e610f97565b606091505b509150915081610fdc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fd390611c82565b60405180910390fd5b600081806020019051810190610ff29190611ce0565b905083600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611076576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106d90611d59565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110db90611deb565b60405180910390fd5b83600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461113391906117d1565b925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f19350505050158015611180573d6000803e3d6000fd5b508473ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5856040516111c7919061150a565b60405180910390a25050505050565b6111de611318565b73ffffffffffffffffffffffffffffffffffffffff166111fc6108b6565b73ffffffffffffffffffffffffffffffffffffffff1614611252576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124990611e57565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261134f5761134e61132a565b5b8235905067ffffffffffffffff81111561136c5761136b61132f565b5b60208301915083602082028301111561138857611387611334565b5b9250929050565b60008083601f8401126113a5576113a461132a565b5b8235905067ffffffffffffffff8111156113c2576113c161132f565b5b6020830191508360208202830111156113de576113dd611334565b5b9250929050565b600080600080604085870312156113ff576113fe611320565b5b600085013567ffffffffffffffff81111561141d5761141c611325565b5b61142987828801611339565b9450945050602085013567ffffffffffffffff81111561144c5761144b611325565b5b6114588782880161138f565b925092505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061149182611466565b9050919050565b6114a181611486565b81146114ac57600080fd5b50565b6000813590506114be81611498565b92915050565b6000602082840312156114da576114d9611320565b5b60006114e8848285016114af565b91505092915050565b6000819050919050565b611504816114f1565b82525050565b600060208201905061151f60008301846114fb565b92915050565b60008115159050919050565b61153a81611525565b82525050565b60006020820190506115556000830184611531565b92915050565b61156481611486565b82525050565b600060208201905061157f600083018461155b565b92915050565b61158e816114f1565b811461159957600080fd5b50565b6000813590506115ab81611585565b92915050565b600080604083850312156115c8576115c7611320565b5b60006115d6858286016114af565b92505060206115e78582860161159c565b9150509250929050565b60006020828403121561160757611606611320565b5b60006116158482850161159c565b91505092915050565b600082825260208201905092915050565b7f4e6f7420616e2061646d696e206f72206f776e65720000000000000000000000600082015250565b600061166560158361161e565b91506116708261162f565b602082019050919050565b6000602082019050818103600083015261169481611658565b9050919050565b7f417272617973206c656e677468206d69736d6174636800000000000000000000600082015250565b60006116d160168361161e565b91506116dc8261169b565b602082019050919050565b60006020820190508181036000830152611700816116c4565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f496e73756666696369656e742062616c616e6365000000000000000000000000600082015250565b600061176c60148361161e565b915061177782611736565b602082019050919050565b6000602082019050818103600083015261179b8161175f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117dc826114f1565b91506117e7836114f1565b92508282039050818111156117ff576117fe6117a2565b5b92915050565b6000611810826114f1565b915061181b836114f1565b9250828201905080821115611833576118326117a2565b5b92915050565b600082825260208201905092915050565b6000819050919050565b61185d81611486565b82525050565b600061186f8383611854565b60208301905092915050565b600061188a60208401846114af565b905092915050565b6000602082019050919050565b60006118ab8385611839565b93506118b68261184a565b8060005b858110156118ef576118cc828461187b565b6118d68882611863565b97506118e183611892565b9250506001810190506118ba565b5085925050509392505050565b600082825260208201905092915050565b600080fd5b82818337505050565b600061192783856118fc565b93507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561195a5761195961190d565b5b60208302925061196b838584611912565b82840190509392505050565b6000604082019050818103600083015261199281868861189f565b905081810360208301526119a781848661191b565b905095945050505050565b7f496e76616c6964206e6577206f776e6572206164647265737300000000000000600082015250565b60006119e860198361161e565b91506119f3826119b2565b602082019050919050565b60006020820190508181036000830152611a17816119db565b9050919050565b7f496e73756666696369656e7420657363726f772062616c616e63650000000000600082015250565b6000611a54601b8361161e565b9150611a5f82611a1e565b602082019050919050565b60006020820190508181036000830152611a8381611a47565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611ae660268361161e565b9150611af182611a8a565b604082019050919050565b60006020820190508181036000830152611b1581611ad9565b9050919050565b7f496e76616c6964206163636f756e742061646472657373000000000000000000600082015250565b6000611b5260178361161e565b9150611b5d82611b1c565b602082019050919050565b60006020820190508181036000830152611b8181611b45565b9050919050565b600081519050919050565b600081905092915050565b60005b83811015611bbc578082015181840152602081019050611ba1565b60008484015250505050565b6000611bd382611b88565b611bdd8185611b93565b9350611bed818560208601611b9e565b80840191505092915050565b6000611c058284611bc8565b915081905092915050565b7f4661696c656420746f2063616c6c206765744f776e65722066756e6374696f6e60008201527f206f662043504163636f756e7400000000000000000000000000000000000000602082015250565b6000611c6c602d8361161e565b9150611c7782611c10565b604082019050919050565b60006020820190508181036000830152611c9b81611c5f565b9050919050565b6000611cad82611466565b9050919050565b611cbd81611ca2565b8114611cc857600080fd5b50565b600081519050611cda81611cb4565b92915050565b600060208284031215611cf657611cf5611320565b5b6000611d0484828501611ccb565b91505092915050565b7f576974686472617720616d6f756e7420657863656564732062616c616e636500600082015250565b6000611d43601f8361161e565b9150611d4e82611d0d565b602082019050919050565b60006020820190508181036000830152611d7281611d36565b9050919050565b7f4f6e6c792043502773206f776e65722063616e2077697468647261772074686560008201527f20636f6c6c61746572616c2066756e6473000000000000000000000000000000602082015250565b6000611dd560318361161e565b9150611de082611d79565b604082019050919050565b60006020820190508181036000830152611e0481611dc8565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611e4160208361161e565b9150611e4c82611e0b565b602082019050919050565b60006020820190508181036000830152611e7081611e34565b905091905056fea2646970667358221220846c9bc554d4b679bf5980d0f3356858d7bb024602ef2121ba6deead6fa2446164736f6c63430008190033",
}

// EcpSequencerABI is the input ABI used to generate the binding from.
// Deprecated: Use EcpSequencerMetaData.ABI instead.
var EcpSequencerABI = EcpSequencerMetaData.ABI

// EcpSequencerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EcpSequencerMetaData.Bin instead.
var EcpSequencerBin = EcpSequencerMetaData.Bin

// DeployEcpSequencer deploys a new Ethereum contract, binding an instance of EcpSequencer to it.
func DeployEcpSequencer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EcpSequencer, error) {
	parsed, err := EcpSequencerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EcpSequencerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EcpSequencer{EcpSequencerCaller: EcpSequencerCaller{contract: contract}, EcpSequencerTransactor: EcpSequencerTransactor{contract: contract}, EcpSequencerFilterer: EcpSequencerFilterer{contract: contract}}, nil
}

// EcpSequencer is an auto generated Go binding around an Ethereum contract.
type EcpSequencer struct {
	EcpSequencerCaller     // Read-only binding to the contract
	EcpSequencerTransactor // Write-only binding to the contract
	EcpSequencerFilterer   // Log filterer for contract events
}

// EcpSequencerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EcpSequencerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpSequencerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EcpSequencerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpSequencerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EcpSequencerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EcpSequencerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EcpSequencerSession struct {
	Contract     *EcpSequencer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EcpSequencerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EcpSequencerCallerSession struct {
	Contract *EcpSequencerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EcpSequencerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EcpSequencerTransactorSession struct {
	Contract     *EcpSequencerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EcpSequencerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EcpSequencerRaw struct {
	Contract *EcpSequencer // Generic contract binding to access the raw methods on
}

// EcpSequencerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EcpSequencerCallerRaw struct {
	Contract *EcpSequencerCaller // Generic read-only contract binding to access the raw methods on
}

// EcpSequencerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EcpSequencerTransactorRaw struct {
	Contract *EcpSequencerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEcpSequencer creates a new instance of EcpSequencer, bound to a specific deployed contract.
func NewEcpSequencer(address common.Address, backend bind.ContractBackend) (*EcpSequencer, error) {
	contract, err := bindEcpSequencer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EcpSequencer{EcpSequencerCaller: EcpSequencerCaller{contract: contract}, EcpSequencerTransactor: EcpSequencerTransactor{contract: contract}, EcpSequencerFilterer: EcpSequencerFilterer{contract: contract}}, nil
}

// NewEcpSequencerCaller creates a new read-only instance of EcpSequencer, bound to a specific deployed contract.
func NewEcpSequencerCaller(address common.Address, caller bind.ContractCaller) (*EcpSequencerCaller, error) {
	contract, err := bindEcpSequencer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerCaller{contract: contract}, nil
}

// NewEcpSequencerTransactor creates a new write-only instance of EcpSequencer, bound to a specific deployed contract.
func NewEcpSequencerTransactor(address common.Address, transactor bind.ContractTransactor) (*EcpSequencerTransactor, error) {
	contract, err := bindEcpSequencer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerTransactor{contract: contract}, nil
}

// NewEcpSequencerFilterer creates a new log filterer instance of EcpSequencer, bound to a specific deployed contract.
func NewEcpSequencerFilterer(address common.Address, filterer bind.ContractFilterer) (*EcpSequencerFilterer, error) {
	contract, err := bindEcpSequencer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerFilterer{contract: contract}, nil
}

// bindEcpSequencer binds a generic wrapper to an already deployed contract.
func bindEcpSequencer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EcpSequencerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EcpSequencer *EcpSequencerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EcpSequencer.Contract.EcpSequencerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EcpSequencer *EcpSequencerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpSequencer.Contract.EcpSequencerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EcpSequencer *EcpSequencerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EcpSequencer.Contract.EcpSequencerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EcpSequencer *EcpSequencerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EcpSequencer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EcpSequencer *EcpSequencerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpSequencer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EcpSequencer *EcpSequencerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EcpSequencer.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EcpSequencer *EcpSequencerCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EcpSequencer.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EcpSequencer *EcpSequencerSession) Admins(arg0 common.Address) (bool, error) {
	return _EcpSequencer.Contract.Admins(&_EcpSequencer.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EcpSequencer *EcpSequencerCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _EcpSequencer.Contract.Admins(&_EcpSequencer.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_EcpSequencer *EcpSequencerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EcpSequencer.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_EcpSequencer *EcpSequencerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EcpSequencer.Contract.Balances(&_EcpSequencer.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_EcpSequencer *EcpSequencerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EcpSequencer.Contract.Balances(&_EcpSequencer.CallOpts, arg0)
}

// EscrowBalance is a free data retrieval call binding the contract method 0xe11e1b0c.
//
// Solidity: function escrowBalance() view returns(uint256)
func (_EcpSequencer *EcpSequencerCaller) EscrowBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EcpSequencer.contract.Call(opts, &out, "escrowBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscrowBalance is a free data retrieval call binding the contract method 0xe11e1b0c.
//
// Solidity: function escrowBalance() view returns(uint256)
func (_EcpSequencer *EcpSequencerSession) EscrowBalance() (*big.Int, error) {
	return _EcpSequencer.Contract.EscrowBalance(&_EcpSequencer.CallOpts)
}

// EscrowBalance is a free data retrieval call binding the contract method 0xe11e1b0c.
//
// Solidity: function escrowBalance() view returns(uint256)
func (_EcpSequencer *EcpSequencerCallerSession) EscrowBalance() (*big.Int, error) {
	return _EcpSequencer.Contract.EscrowBalance(&_EcpSequencer.CallOpts)
}

// GetCPBalance is a free data retrieval call binding the contract method 0x922f126a.
//
// Solidity: function getCPBalance(address cpAccount) view returns(uint256)
func (_EcpSequencer *EcpSequencerCaller) GetCPBalance(opts *bind.CallOpts, cpAccount common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EcpSequencer.contract.Call(opts, &out, "getCPBalance", cpAccount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCPBalance is a free data retrieval call binding the contract method 0x922f126a.
//
// Solidity: function getCPBalance(address cpAccount) view returns(uint256)
func (_EcpSequencer *EcpSequencerSession) GetCPBalance(cpAccount common.Address) (*big.Int, error) {
	return _EcpSequencer.Contract.GetCPBalance(&_EcpSequencer.CallOpts, cpAccount)
}

// GetCPBalance is a free data retrieval call binding the contract method 0x922f126a.
//
// Solidity: function getCPBalance(address cpAccount) view returns(uint256)
func (_EcpSequencer *EcpSequencerCallerSession) GetCPBalance(cpAccount common.Address) (*big.Int, error) {
	return _EcpSequencer.Contract.GetCPBalance(&_EcpSequencer.CallOpts, cpAccount)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpSequencer *EcpSequencerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EcpSequencer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpSequencer *EcpSequencerSession) Owner() (common.Address, error) {
	return _EcpSequencer.Contract.Owner(&_EcpSequencer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EcpSequencer *EcpSequencerCallerSession) Owner() (common.Address, error) {
	return _EcpSequencer.Contract.Owner(&_EcpSequencer.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_EcpSequencer *EcpSequencerTransactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_EcpSequencer *EcpSequencerSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.AddAdmin(&_EcpSequencer.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_EcpSequencer *EcpSequencerTransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.AddAdmin(&_EcpSequencer.TransactOpts, _admin)
}

// BatchTransferToEscrow is a paid mutator transaction binding the contract method 0x1485fef9.
//
// Solidity: function batchTransferToEscrow(address[] cpAccounts, uint256[] amounts) returns()
func (_EcpSequencer *EcpSequencerTransactor) BatchTransferToEscrow(opts *bind.TransactOpts, cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "batchTransferToEscrow", cpAccounts, amounts)
}

// BatchTransferToEscrow is a paid mutator transaction binding the contract method 0x1485fef9.
//
// Solidity: function batchTransferToEscrow(address[] cpAccounts, uint256[] amounts) returns()
func (_EcpSequencer *EcpSequencerSession) BatchTransferToEscrow(cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _EcpSequencer.Contract.BatchTransferToEscrow(&_EcpSequencer.TransactOpts, cpAccounts, amounts)
}

// BatchTransferToEscrow is a paid mutator transaction binding the contract method 0x1485fef9.
//
// Solidity: function batchTransferToEscrow(address[] cpAccounts, uint256[] amounts) returns()
func (_EcpSequencer *EcpSequencerTransactorSession) BatchTransferToEscrow(cpAccounts []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _EcpSequencer.Contract.BatchTransferToEscrow(&_EcpSequencer.TransactOpts, cpAccounts, amounts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_EcpSequencer *EcpSequencerTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_EcpSequencer *EcpSequencerSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.ChangeOwner(&_EcpSequencer.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_EcpSequencer *EcpSequencerTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.ChangeOwner(&_EcpSequencer.TransactOpts, newOwner)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address cpAccount) payable returns()
func (_EcpSequencer *EcpSequencerTransactor) Deposit(opts *bind.TransactOpts, cpAccount common.Address) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "deposit", cpAccount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address cpAccount) payable returns()
func (_EcpSequencer *EcpSequencerSession) Deposit(cpAccount common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.Deposit(&_EcpSequencer.TransactOpts, cpAccount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address cpAccount) payable returns()
func (_EcpSequencer *EcpSequencerTransactorSession) Deposit(cpAccount common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.Deposit(&_EcpSequencer.TransactOpts, cpAccount)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_EcpSequencer *EcpSequencerTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_EcpSequencer *EcpSequencerSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.RemoveAdmin(&_EcpSequencer.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_EcpSequencer *EcpSequencerTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.RemoveAdmin(&_EcpSequencer.TransactOpts, _admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpSequencer *EcpSequencerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpSequencer *EcpSequencerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EcpSequencer.Contract.RenounceOwnership(&_EcpSequencer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EcpSequencer *EcpSequencerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EcpSequencer.Contract.RenounceOwnership(&_EcpSequencer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpSequencer *EcpSequencerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpSequencer *EcpSequencerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.TransferOwnership(&_EcpSequencer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EcpSequencer *EcpSequencerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EcpSequencer.Contract.TransferOwnership(&_EcpSequencer.TransactOpts, newOwner)
}

// TransferToEscrow is a paid mutator transaction binding the contract method 0x9e88fd81.
//
// Solidity: function transferToEscrow(address cpAccount, uint256 amount) returns()
func (_EcpSequencer *EcpSequencerTransactor) TransferToEscrow(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "transferToEscrow", cpAccount, amount)
}

// TransferToEscrow is a paid mutator transaction binding the contract method 0x9e88fd81.
//
// Solidity: function transferToEscrow(address cpAccount, uint256 amount) returns()
func (_EcpSequencer *EcpSequencerSession) TransferToEscrow(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.Contract.TransferToEscrow(&_EcpSequencer.TransactOpts, cpAccount, amount)
}

// TransferToEscrow is a paid mutator transaction binding the contract method 0x9e88fd81.
//
// Solidity: function transferToEscrow(address cpAccount, uint256 amount) returns()
func (_EcpSequencer *EcpSequencerTransactorSession) TransferToEscrow(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.Contract.TransferToEscrow(&_EcpSequencer.TransactOpts, cpAccount, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_EcpSequencer *EcpSequencerTransactor) Withdraw(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "withdraw", cpAccount, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_EcpSequencer *EcpSequencerSession) Withdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.Contract.Withdraw(&_EcpSequencer.TransactOpts, cpAccount, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_EcpSequencer *EcpSequencerTransactorSession) Withdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.Contract.Withdraw(&_EcpSequencer.TransactOpts, cpAccount, amount)
}

// WithdrawFromEscrow is a paid mutator transaction binding the contract method 0xbe47afda.
//
// Solidity: function withdrawFromEscrow(uint256 amount) returns()
func (_EcpSequencer *EcpSequencerTransactor) WithdrawFromEscrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.contract.Transact(opts, "withdrawFromEscrow", amount)
}

// WithdrawFromEscrow is a paid mutator transaction binding the contract method 0xbe47afda.
//
// Solidity: function withdrawFromEscrow(uint256 amount) returns()
func (_EcpSequencer *EcpSequencerSession) WithdrawFromEscrow(amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.Contract.WithdrawFromEscrow(&_EcpSequencer.TransactOpts, amount)
}

// WithdrawFromEscrow is a paid mutator transaction binding the contract method 0xbe47afda.
//
// Solidity: function withdrawFromEscrow(uint256 amount) returns()
func (_EcpSequencer *EcpSequencerTransactorSession) WithdrawFromEscrow(amount *big.Int) (*types.Transaction, error) {
	return _EcpSequencer.Contract.WithdrawFromEscrow(&_EcpSequencer.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EcpSequencer *EcpSequencerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EcpSequencer.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EcpSequencer *EcpSequencerSession) Receive() (*types.Transaction, error) {
	return _EcpSequencer.Contract.Receive(&_EcpSequencer.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EcpSequencer *EcpSequencerTransactorSession) Receive() (*types.Transaction, error) {
	return _EcpSequencer.Contract.Receive(&_EcpSequencer.TransactOpts)
}

// EcpSequencerAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the EcpSequencer contract.
type EcpSequencerAdminAddedIterator struct {
	Event *EcpSequencerAdminAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerAdminAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerAdminAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerAdminAdded represents a AdminAdded event raised by the EcpSequencer contract.
type EcpSequencerAdminAdded struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed admin)
func (_EcpSequencer *EcpSequencerFilterer) FilterAdminAdded(opts *bind.FilterOpts, admin []common.Address) (*EcpSequencerAdminAddedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerAdminAddedIterator{contract: _EcpSequencer.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed admin)
func (_EcpSequencer *EcpSequencerFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *EcpSequencerAdminAdded, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerAdminAdded)
				if err := _EcpSequencer.contract.UnpackLog(event, "AdminAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed admin)
func (_EcpSequencer *EcpSequencerFilterer) ParseAdminAdded(log types.Log) (*EcpSequencerAdminAdded, error) {
	event := new(EcpSequencerAdminAdded)
	if err := _EcpSequencer.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpSequencerAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the EcpSequencer contract.
type EcpSequencerAdminRemovedIterator struct {
	Event *EcpSequencerAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerAdminRemoved represents a AdminRemoved event raised by the EcpSequencer contract.
type EcpSequencerAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed admin)
func (_EcpSequencer *EcpSequencerFilterer) FilterAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*EcpSequencerAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerAdminRemovedIterator{contract: _EcpSequencer.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed admin)
func (_EcpSequencer *EcpSequencerFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *EcpSequencerAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerAdminRemoved)
				if err := _EcpSequencer.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed admin)
func (_EcpSequencer *EcpSequencerFilterer) ParseAdminRemoved(log types.Log) (*EcpSequencerAdminRemoved, error) {
	event := new(EcpSequencerAdminRemoved)
	if err := _EcpSequencer.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpSequencerBatchTransferredToEscrowIterator is returned from FilterBatchTransferredToEscrow and is used to iterate over the raw logs and unpacked data for BatchTransferredToEscrow events raised by the EcpSequencer contract.
type EcpSequencerBatchTransferredToEscrowIterator struct {
	Event *EcpSequencerBatchTransferredToEscrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerBatchTransferredToEscrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerBatchTransferredToEscrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerBatchTransferredToEscrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerBatchTransferredToEscrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerBatchTransferredToEscrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerBatchTransferredToEscrow represents a BatchTransferredToEscrow event raised by the EcpSequencer contract.
type EcpSequencerBatchTransferredToEscrow struct {
	Admin      common.Address
	CpAccounts []common.Address
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchTransferredToEscrow is a free log retrieval operation binding the contract event 0x5a5b8c52e21fd816e73687ac900bbb3238b4e80ddcb10c216085993299e64ed5.
//
// Solidity: event BatchTransferredToEscrow(address indexed admin, address[] cpAccounts, uint256[] amounts)
func (_EcpSequencer *EcpSequencerFilterer) FilterBatchTransferredToEscrow(opts *bind.FilterOpts, admin []common.Address) (*EcpSequencerBatchTransferredToEscrowIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "BatchTransferredToEscrow", adminRule)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerBatchTransferredToEscrowIterator{contract: _EcpSequencer.contract, event: "BatchTransferredToEscrow", logs: logs, sub: sub}, nil
}

// WatchBatchTransferredToEscrow is a free log subscription operation binding the contract event 0x5a5b8c52e21fd816e73687ac900bbb3238b4e80ddcb10c216085993299e64ed5.
//
// Solidity: event BatchTransferredToEscrow(address indexed admin, address[] cpAccounts, uint256[] amounts)
func (_EcpSequencer *EcpSequencerFilterer) WatchBatchTransferredToEscrow(opts *bind.WatchOpts, sink chan<- *EcpSequencerBatchTransferredToEscrow, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "BatchTransferredToEscrow", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerBatchTransferredToEscrow)
				if err := _EcpSequencer.contract.UnpackLog(event, "BatchTransferredToEscrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchTransferredToEscrow is a log parse operation binding the contract event 0x5a5b8c52e21fd816e73687ac900bbb3238b4e80ddcb10c216085993299e64ed5.
//
// Solidity: event BatchTransferredToEscrow(address indexed admin, address[] cpAccounts, uint256[] amounts)
func (_EcpSequencer *EcpSequencerFilterer) ParseBatchTransferredToEscrow(log types.Log) (*EcpSequencerBatchTransferredToEscrow, error) {
	event := new(EcpSequencerBatchTransferredToEscrow)
	if err := _EcpSequencer.contract.UnpackLog(event, "BatchTransferredToEscrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpSequencerDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the EcpSequencer contract.
type EcpSequencerDepositedIterator struct {
	Event *EcpSequencerDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerDeposited represents a Deposited event raised by the EcpSequencer contract.
type EcpSequencerDeposited struct {
	CpAccount common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) FilterDeposited(opts *bind.FilterOpts, cpAccount []common.Address) (*EcpSequencerDepositedIterator, error) {

	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "Deposited", cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerDepositedIterator{contract: _EcpSequencer.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EcpSequencerDeposited, cpAccount []common.Address) (event.Subscription, error) {

	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "Deposited", cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerDeposited)
				if err := _EcpSequencer.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) ParseDeposited(log types.Log) (*EcpSequencerDeposited, error) {
	event := new(EcpSequencerDeposited)
	if err := _EcpSequencer.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpSequencerOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the EcpSequencer contract.
type EcpSequencerOwnerChangedIterator struct {
	Event *EcpSequencerOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerOwnerChanged represents a OwnerChanged event raised by the EcpSequencer contract.
type EcpSequencerOwnerChanged struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed previousOwner, address indexed newOwner)
func (_EcpSequencer *EcpSequencerFilterer) FilterOwnerChanged(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EcpSequencerOwnerChangedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "OwnerChanged", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerOwnerChangedIterator{contract: _EcpSequencer.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed previousOwner, address indexed newOwner)
func (_EcpSequencer *EcpSequencerFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *EcpSequencerOwnerChanged, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "OwnerChanged", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerOwnerChanged)
				if err := _EcpSequencer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed previousOwner, address indexed newOwner)
func (_EcpSequencer *EcpSequencerFilterer) ParseOwnerChanged(log types.Log) (*EcpSequencerOwnerChanged, error) {
	event := new(EcpSequencerOwnerChanged)
	if err := _EcpSequencer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpSequencerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EcpSequencer contract.
type EcpSequencerOwnershipTransferredIterator struct {
	Event *EcpSequencerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerOwnershipTransferred represents a OwnershipTransferred event raised by the EcpSequencer contract.
type EcpSequencerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EcpSequencer *EcpSequencerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EcpSequencerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerOwnershipTransferredIterator{contract: _EcpSequencer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EcpSequencer *EcpSequencerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EcpSequencerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerOwnershipTransferred)
				if err := _EcpSequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EcpSequencer *EcpSequencerFilterer) ParseOwnershipTransferred(log types.Log) (*EcpSequencerOwnershipTransferred, error) {
	event := new(EcpSequencerOwnershipTransferred)
	if err := _EcpSequencer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpSequencerTransferredToEscrowIterator is returned from FilterTransferredToEscrow and is used to iterate over the raw logs and unpacked data for TransferredToEscrow events raised by the EcpSequencer contract.
type EcpSequencerTransferredToEscrowIterator struct {
	Event *EcpSequencerTransferredToEscrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerTransferredToEscrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerTransferredToEscrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerTransferredToEscrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerTransferredToEscrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerTransferredToEscrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerTransferredToEscrow represents a TransferredToEscrow event raised by the EcpSequencer contract.
type EcpSequencerTransferredToEscrow struct {
	CpAccount common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferredToEscrow is a free log retrieval operation binding the contract event 0x8aa981417e2c3a0dadea5dd2eb5eab703d6d9c074ae83d377ce454dd6a93bfa7.
//
// Solidity: event TransferredToEscrow(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) FilterTransferredToEscrow(opts *bind.FilterOpts, cpAccount []common.Address) (*EcpSequencerTransferredToEscrowIterator, error) {

	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "TransferredToEscrow", cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerTransferredToEscrowIterator{contract: _EcpSequencer.contract, event: "TransferredToEscrow", logs: logs, sub: sub}, nil
}

// WatchTransferredToEscrow is a free log subscription operation binding the contract event 0x8aa981417e2c3a0dadea5dd2eb5eab703d6d9c074ae83d377ce454dd6a93bfa7.
//
// Solidity: event TransferredToEscrow(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) WatchTransferredToEscrow(opts *bind.WatchOpts, sink chan<- *EcpSequencerTransferredToEscrow, cpAccount []common.Address) (event.Subscription, error) {

	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "TransferredToEscrow", cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerTransferredToEscrow)
				if err := _EcpSequencer.contract.UnpackLog(event, "TransferredToEscrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferredToEscrow is a log parse operation binding the contract event 0x8aa981417e2c3a0dadea5dd2eb5eab703d6d9c074ae83d377ce454dd6a93bfa7.
//
// Solidity: event TransferredToEscrow(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) ParseTransferredToEscrow(log types.Log) (*EcpSequencerTransferredToEscrow, error) {
	event := new(EcpSequencerTransferredToEscrow)
	if err := _EcpSequencer.contract.UnpackLog(event, "TransferredToEscrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpSequencerWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the EcpSequencer contract.
type EcpSequencerWithdrawnIterator struct {
	Event *EcpSequencerWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerWithdrawn represents a Withdrawn event raised by the EcpSequencer contract.
type EcpSequencerWithdrawn struct {
	CpAccount common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) FilterWithdrawn(opts *bind.FilterOpts, cpAccount []common.Address) (*EcpSequencerWithdrawnIterator, error) {

	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "Withdrawn", cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &EcpSequencerWithdrawnIterator{contract: _EcpSequencer.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EcpSequencerWithdrawn, cpAccount []common.Address) (event.Subscription, error) {

	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "Withdrawn", cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerWithdrawn)
				if err := _EcpSequencer.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed cpAccount, uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) ParseWithdrawn(log types.Log) (*EcpSequencerWithdrawn, error) {
	event := new(EcpSequencerWithdrawn)
	if err := _EcpSequencer.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EcpSequencerWithdrawnFromEscrowIterator is returned from FilterWithdrawnFromEscrow and is used to iterate over the raw logs and unpacked data for WithdrawnFromEscrow events raised by the EcpSequencer contract.
type EcpSequencerWithdrawnFromEscrowIterator struct {
	Event *EcpSequencerWithdrawnFromEscrow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EcpSequencerWithdrawnFromEscrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EcpSequencerWithdrawnFromEscrow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EcpSequencerWithdrawnFromEscrow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EcpSequencerWithdrawnFromEscrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EcpSequencerWithdrawnFromEscrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EcpSequencerWithdrawnFromEscrow represents a WithdrawnFromEscrow event raised by the EcpSequencer contract.
type EcpSequencerWithdrawnFromEscrow struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFromEscrow is a free log retrieval operation binding the contract event 0xd1f01ada3f4d4c0d5d81675b458e3f5c4dda3e79a82d07b03459147644bc7b7b.
//
// Solidity: event WithdrawnFromEscrow(uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) FilterWithdrawnFromEscrow(opts *bind.FilterOpts) (*EcpSequencerWithdrawnFromEscrowIterator, error) {

	logs, sub, err := _EcpSequencer.contract.FilterLogs(opts, "WithdrawnFromEscrow")
	if err != nil {
		return nil, err
	}
	return &EcpSequencerWithdrawnFromEscrowIterator{contract: _EcpSequencer.contract, event: "WithdrawnFromEscrow", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFromEscrow is a free log subscription operation binding the contract event 0xd1f01ada3f4d4c0d5d81675b458e3f5c4dda3e79a82d07b03459147644bc7b7b.
//
// Solidity: event WithdrawnFromEscrow(uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) WatchWithdrawnFromEscrow(opts *bind.WatchOpts, sink chan<- *EcpSequencerWithdrawnFromEscrow) (event.Subscription, error) {

	logs, sub, err := _EcpSequencer.contract.WatchLogs(opts, "WithdrawnFromEscrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EcpSequencerWithdrawnFromEscrow)
				if err := _EcpSequencer.contract.UnpackLog(event, "WithdrawnFromEscrow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawnFromEscrow is a log parse operation binding the contract event 0xd1f01ada3f4d4c0d5d81675b458e3f5c4dda3e79a82d07b03459147644bc7b7b.
//
// Solidity: event WithdrawnFromEscrow(uint256 amount)
func (_EcpSequencer *EcpSequencerFilterer) ParseWithdrawnFromEscrow(log types.Log) (*EcpSequencerWithdrawnFromEscrow, error) {
	event := new(EcpSequencerWithdrawnFromEscrow)
	if err := _EcpSequencer.contract.UnpackLog(event, "WithdrawnFromEscrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
