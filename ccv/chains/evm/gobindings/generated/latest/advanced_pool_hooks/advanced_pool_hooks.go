// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package advanced_pool_hooks

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

type AdvancedPoolHooksCCVConfigArg struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
}

type PoolLockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      common.Address
	Amount              *big.Int
	LocalToken          common.Address
}

type PoolReleaseOrMintInV1 struct {
	OriginalSender          []byte
	RemoteChainSelector     uint64
	Receiver                common.Address
	SourceDenominatedAmount *big.Int
	LocalToken              common.Address
	SourcePoolAddress       []byte
	SourcePoolData          []byte
	OffchainTokenData       []byte
}

var AdvancedPoolHooksMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct AdvancedPoolHooks.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkAllowList\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getThresholdAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postFlightCheck\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"preflightCheck\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setThresholdAmount\",\"inputs\":[{\"name\":\"thresholdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdAmountSet\",\"inputs\":[{\"name\":\"thresholdAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a06040523461022657611a808038038061001981610241565b9283398101906040818303126102265780516001600160401b03811161022657810182601f82011215610226578051926001600160401b03841161022b578360051b9160208061006a818601610241565b80978152019382010191821161022657602001915b8183106102065760208401518533156101f557600180546001600160a01b031916331790558051151560808190526100dc575b506004556040516116799081610407823960805181818161021201528181610c2b01526110330152f35b60206100e781610241565b600081526000368137608051156101e45760005b8151811015610162576001906001600160a01b036101198285610266565b511684610125826102a8565b610132575b5050016100fb565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1868461012a565b505060005b82518110156101db576001906001600160a01b036101858286610266565b511680156101d55783610197826103a6565b6101a5575b50505b01610167565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1858361019c565b5061019f565b505050816100b2565b6335f4a7b360e01b60005260046000fd5b639b15e16f60e01b60005260046000fd5b82516001600160a01b03811681036102265781526020928301920161007f565b600080fd5b634e487b7160e01b600052604160045260246000fd5b6040519190601f01601f191682016001600160401b0381118382101761022b57604052565b805182101561027a5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561027a5760005260206000200190600090565b600081815260036020526040902054801561039f5760001981018181116103895760025460001981019190821161038957818103610338575b505050600254801561032257600019016102fc816002610290565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61037161034961035a936002610290565b90549060031b1c9283926002610290565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806102e1565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610400576002546801000000000000000081101561022b576103e761035a8260018594016002556002610290565b9055600254906000526003602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806354c8a4f314610b995780635c3af7ca14610ad65780635eff3bf714610a6057806379ba50971461097757806389720a62146108e25780638da5cb5b14610890578063961e2e7c14610826578063a7cd63b714610781578063ce07c7c814610740578063d966866b14610237578063e0351e13146101dc578063f2fde38b146100ec5763f72c071b146100ab57600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576020600454604051908152f35b600080fd5b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff610138610e05565b6101406111c5565b163381146101b257807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610286903690600401610d73565b61028e6111c5565b6000905b80821061029b57005b6102a68282856110b3565b359067ffffffffffffffff821682036100e7576102d16102c78483876110b3565b6020810190611122565b6102eb6102e186858995996110b3565b6040810190611122565b94906103056102fb8887876110b3565b6060810190611122565b61031d6103138a89896110b3565b6080810190611122565b93909861033461032f8d893691610ef2565b611394565b61034261032f368385610ef2565b61035061032f368587610ef2565b61035e61032f36878d610ef2565b6040519b60808d018d811067ffffffffffffffff82111761061057604052610387368983610ef2565b8d528c8b610396368587610ef2565b602083019081526103bb6103ab36898b610ef2565b92604085019384528a3691610ef2565b606084015267ffffffffffffffff8a1660005260056020526040600020925180519067ffffffffffffffff821161061057680100000000000000008211610610578454828655808310610717575b5060200184600052602060002060005b8381106106ed57505050506001839e9c9d9e0190519081519167ffffffffffffffff8311610610576801000000000000000083116106105781548383558084106106bc575b5060200190600052602060002060005b83811061069257505050506002820190519081519167ffffffffffffffff831161061057680100000000000000008311610610578154838355808410610669575b509e9f939495969798999a9b9c9d9e60200190600052602060002060005b83811061063f575050505060036060919e9c9d9e019101519081519167ffffffffffffffff8311610610576801000000000000000083116106105781548383558084106105e7575b5060200190600052602060002060005b8381106105bd57505050506105a26080956105b2956105947fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a9661058667ffffffffffffffff976040519d8d8f9e8f9081520191611176565b918b830360208d0152611176565b9188830360408a0152611176565b9285840360608701521696611176565b0390a2019091610292565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610525565b8260005283602060002091820191015b8181106106045750610515565b600081556001016105f7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016104cd565b8260005283602060002091820191015b81811061068657506104af565b60008155600101610679565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161046e565b9d9f9e9d8260005283602060002091820191015b8181106106e157509f9d9e9f61045e565b600081556001016106d0565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610419565b8560005282602060002091820191015b8181106107345750610409565b60008155600101610727565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761077f61077a610e05565b611031565b005b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760405180602060025491828152019060026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b8181106108105761080c8561080081870382610e99565b60405191829182610e49565b0390f35b82548452602090930192600192830192016107e9565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7577f80dc2a1a49dda9f8bd85c1c376266e011db6448050b7bfd5c2f423e162c1114560206004356108836111c5565b80600455604051908152a1005b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100e75760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757610919610e05565b5060243567ffffffffffffffff811681036100e757610936610dc6565b5060843567ffffffffffffffff81116100e757610957903690600401610dd7565b505060a4359060028210156100e75761080c916108009160443590610f96565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303610a36577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc61010091360301126100e75761077f610db5565b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e75760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e757610b4b610da4565b506044359067ffffffffffffffff82116100e757610b6f6044923690600401610dd7565b5050013573ffffffffffffffffffffffffffffffffffffffff811681036100e75761077f90611031565b346100e75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610be8903690600401610d73565b6024359067ffffffffffffffff82116100e757610c21610c0f610c29933690600401610d73565b949092610c1a6111c5565b3691610ef2565b923691610ef2565b7f000000000000000000000000000000000000000000000000000000000000000015610d495760005b8251811015610cc5578073ffffffffffffffffffffffffffffffffffffffff610c7d60019386611210565b5116610c8881611476565b610c94575b5001610c52565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184610c8d565b5060005b815181101561077f578073ffffffffffffffffffffffffffffffffffffffff610cf460019385611210565b51168015610d4357610d058161160c565b610d12575b505b01610cc9565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183610d0a565b50610d0c565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b6024359061ffff821682036100e757565b6044359061ffff821682036100e757565b6064359061ffff821682036100e757565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036100e757565b359073ffffffffffffffffffffffffffffffffffffffff821682036100e757565b602060408183019282815284518094520192019060005b818110610e6d5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610e60565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761061057604052565b67ffffffffffffffff81116106105760051b60200190565b929190610efe81610eda565b93610f0c6040519586610e99565b602085838152019160051b81019283116100e757905b828210610f2e57505050565b60208091610f3b84610e28565b815201910190610f22565b906020825491828152019160005260206000209060005b818110610f6a5750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201610f5d565b67ffffffffffffffff166000526005602052604060002091600281101561100257600114610fe757610fe491600160405191610fdd83610fd68184610f46565b0384610e99565b0190611260565b90565b610fe491600360405191610fdd83610fd68160028501610f46565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f00000000000000000000000000000000000000000000000000000000000000006110595750565b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054156110865750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b91908110156110f35760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100e7570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b9160209082815201919060005b8181106111905750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff6111b788610e28565b168152019401929101611183565b73ffffffffffffffffffffffffffffffffffffffff6001541633036111e657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80518210156110f35760209160051b010190565b9190820180921161123157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b916004548015159182611389575b5050611278575090565b9061129361128c9260405193848092610f46565b0383610e99565b8151806112a05750905090565b6112ab908251611224565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06112ef6112d986610eda565b956112e76040519788610e99565b808752610eda565b0136602086013760005b8251811015611337578073ffffffffffffffffffffffffffffffffffffffff61132460019386611210565b51166113308288611210565b52016112f9565b509160005b8151811015611384578073ffffffffffffffffffffffffffffffffffffffff61136760019385611210565b511661137d611377838751611224565b88611210565b520161133c565b505050565b10159050388061126e565b805160005b8181106113a557505050565b60018101808211611231575b8281106113c15750600101611399565b73ffffffffffffffffffffffffffffffffffffffff6113e08386611210565b511673ffffffffffffffffffffffffffffffffffffffff6114018387611210565b511614611410576001016113b1565b73ffffffffffffffffffffffffffffffffffffffff61142f8386611210565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80548210156110f35760005260206000200190600090565b6000818152600360205260409020548015611605577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161123157600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161123157818103611596575b5050506002548015611567577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161152481600261145e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6115ed6115a76115b893600261145e565b90549060031b1c928392600261145e565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260036020526040600020553880806114eb565b5050600090565b8060005260036020526040600020541560001461166657600254680100000000000000008110156106105761164d6115b8826001859401600255600261145e565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
}

var AdvancedPoolHooksABI = AdvancedPoolHooksMetaData.ABI

var AdvancedPoolHooksBin = AdvancedPoolHooksMetaData.Bin

func DeployAdvancedPoolHooks(auth *bind.TransactOpts, backend bind.ContractBackend, allowlist []common.Address, thresholdAmountForAdditionalCCVs *big.Int) (common.Address, *types.Transaction, *AdvancedPoolHooks, error) {
	parsed, err := AdvancedPoolHooksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdvancedPoolHooksBin), backend, allowlist, thresholdAmountForAdditionalCCVs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdvancedPoolHooks{address: address, abi: *parsed, AdvancedPoolHooksCaller: AdvancedPoolHooksCaller{contract: contract}, AdvancedPoolHooksTransactor: AdvancedPoolHooksTransactor{contract: contract}, AdvancedPoolHooksFilterer: AdvancedPoolHooksFilterer{contract: contract}}, nil
}

type AdvancedPoolHooks struct {
	address common.Address
	abi     abi.ABI
	AdvancedPoolHooksCaller
	AdvancedPoolHooksTransactor
	AdvancedPoolHooksFilterer
}

type AdvancedPoolHooksCaller struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksTransactor struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksFilterer struct {
	contract *bind.BoundContract
}

type AdvancedPoolHooksSession struct {
	Contract     *AdvancedPoolHooks
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type AdvancedPoolHooksCallerSession struct {
	Contract *AdvancedPoolHooksCaller
	CallOpts bind.CallOpts
}

type AdvancedPoolHooksTransactorSession struct {
	Contract     *AdvancedPoolHooksTransactor
	TransactOpts bind.TransactOpts
}

type AdvancedPoolHooksRaw struct {
	Contract *AdvancedPoolHooks
}

type AdvancedPoolHooksCallerRaw struct {
	Contract *AdvancedPoolHooksCaller
}

type AdvancedPoolHooksTransactorRaw struct {
	Contract *AdvancedPoolHooksTransactor
}

func NewAdvancedPoolHooks(address common.Address, backend bind.ContractBackend) (*AdvancedPoolHooks, error) {
	abi, err := abi.JSON(strings.NewReader(AdvancedPoolHooksABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindAdvancedPoolHooks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooks{address: address, abi: abi, AdvancedPoolHooksCaller: AdvancedPoolHooksCaller{contract: contract}, AdvancedPoolHooksTransactor: AdvancedPoolHooksTransactor{contract: contract}, AdvancedPoolHooksFilterer: AdvancedPoolHooksFilterer{contract: contract}}, nil
}

func NewAdvancedPoolHooksCaller(address common.Address, caller bind.ContractCaller) (*AdvancedPoolHooksCaller, error) {
	contract, err := bindAdvancedPoolHooks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksCaller{contract: contract}, nil
}

func NewAdvancedPoolHooksTransactor(address common.Address, transactor bind.ContractTransactor) (*AdvancedPoolHooksTransactor, error) {
	contract, err := bindAdvancedPoolHooks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksTransactor{contract: contract}, nil
}

func NewAdvancedPoolHooksFilterer(address common.Address, filterer bind.ContractFilterer) (*AdvancedPoolHooksFilterer, error) {
	contract, err := bindAdvancedPoolHooks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksFilterer{contract: contract}, nil
}

func bindAdvancedPoolHooks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdvancedPoolHooksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvancedPoolHooks.Contract.AdvancedPoolHooksCaller.contract.Call(opts, result, method, params...)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.AdvancedPoolHooksTransactor.contract.Transfer(opts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.AdvancedPoolHooksTransactor.contract.Transact(opts, method, params...)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvancedPoolHooks.Contract.contract.Call(opts, result, method, params...)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.contract.Transfer(opts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.contract.Transact(opts, method, params...)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) CheckAllowList(opts *bind.CallOpts, sender common.Address) error {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "checkAllowList", sender)

	if err != nil {
		return err
	}

	return err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) CheckAllowList(sender common.Address) error {
	return _AdvancedPoolHooks.Contract.CheckAllowList(&_AdvancedPoolHooks.CallOpts, sender)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) CheckAllowList(sender common.Address) error {
	return _AdvancedPoolHooks.Contract.CheckAllowList(&_AdvancedPoolHooks.CallOpts, sender)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetAllowList() ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetAllowList(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetAllowList() ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetAllowList(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetAllowListEnabled() (bool, error) {
	return _AdvancedPoolHooks.Contract.GetAllowListEnabled(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetAllowListEnabled() (bool, error) {
	return _AdvancedPoolHooks.Contract.GetAllowListEnabled(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetRequiredCCVs(&_AdvancedPoolHooks.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _AdvancedPoolHooks.Contract.GetRequiredCCVs(&_AdvancedPoolHooks.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) GetThresholdAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "getThresholdAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) GetThresholdAmount() (*big.Int, error) {
	return _AdvancedPoolHooks.Contract.GetThresholdAmount(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) GetThresholdAmount() (*big.Int, error) {
	return _AdvancedPoolHooks.Contract.GetThresholdAmount(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) Owner() (common.Address, error) {
	return _AdvancedPoolHooks.Contract.Owner(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) Owner() (common.Address, error) {
	return _AdvancedPoolHooks.Contract.Owner(&_AdvancedPoolHooks.CallOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) PostFlightCheck(opts *bind.CallOpts, arg0 PoolReleaseOrMintInV1, arg1 *big.Int, arg2 uint16) error {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "postFlightCheck", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) PostFlightCheck(arg0 PoolReleaseOrMintInV1, arg1 *big.Int, arg2 uint16) error {
	return _AdvancedPoolHooks.Contract.PostFlightCheck(&_AdvancedPoolHooks.CallOpts, arg0, arg1, arg2)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) PostFlightCheck(arg0 PoolReleaseOrMintInV1, arg1 *big.Int, arg2 uint16) error {
	return _AdvancedPoolHooks.Contract.PostFlightCheck(&_AdvancedPoolHooks.CallOpts, arg0, arg1, arg2)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCaller) PreflightCheck(opts *bind.CallOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 uint16, arg2 []byte) error {
	var out []interface{}
	err := _AdvancedPoolHooks.contract.Call(opts, &out, "preflightCheck", lockOrBurnIn, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, arg1 uint16, arg2 []byte) error {
	return _AdvancedPoolHooks.Contract.PreflightCheck(&_AdvancedPoolHooks.CallOpts, lockOrBurnIn, arg1, arg2)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksCallerSession) PreflightCheck(lockOrBurnIn PoolLockOrBurnInV1, arg1 uint16, arg2 []byte) error {
	return _AdvancedPoolHooks.Contract.PreflightCheck(&_AdvancedPoolHooks.CallOpts, lockOrBurnIn, arg1, arg2)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "acceptOwnership")
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) AcceptOwnership() (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.AcceptOwnership(&_AdvancedPoolHooks.TransactOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.AcceptOwnership(&_AdvancedPoolHooks.TransactOpts)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyAllowListUpdates(&_AdvancedPoolHooks.TransactOpts, removes, adds)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyAllowListUpdates(&_AdvancedPoolHooks.TransactOpts, removes, adds)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) ApplyCCVConfigUpdates(ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyCCVConfigUpdates(&_AdvancedPoolHooks.TransactOpts, ccvConfigArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.ApplyCCVConfigUpdates(&_AdvancedPoolHooks.TransactOpts, ccvConfigArgs)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) SetThresholdAmount(opts *bind.TransactOpts, thresholdAmount *big.Int) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "setThresholdAmount", thresholdAmount)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) SetThresholdAmount(thresholdAmount *big.Int) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.SetThresholdAmount(&_AdvancedPoolHooks.TransactOpts, thresholdAmount)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) SetThresholdAmount(thresholdAmount *big.Int) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.SetThresholdAmount(&_AdvancedPoolHooks.TransactOpts, thresholdAmount)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.contract.Transact(opts, "transferOwnership", to)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.TransferOwnership(&_AdvancedPoolHooks.TransactOpts, to)
}

func (_AdvancedPoolHooks *AdvancedPoolHooksTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AdvancedPoolHooks.Contract.TransferOwnership(&_AdvancedPoolHooks.TransactOpts, to)
}

type AdvancedPoolHooksAllowListAddIterator struct {
	Event *AdvancedPoolHooksAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksAllowListAdd)
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

	select {
	case log := <-it.logs:
		it.Event = new(AdvancedPoolHooksAllowListAdd)
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

func (it *AdvancedPoolHooksAllowListAddIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListAddIterator, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksAllowListAddIterator{contract: _AdvancedPoolHooks.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksAllowListAdd)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseAllowListAdd(log types.Log) (*AdvancedPoolHooksAllowListAdd, error) {
	event := new(AdvancedPoolHooksAllowListAdd)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksAllowListRemoveIterator struct {
	Event *AdvancedPoolHooksAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksAllowListRemove)
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

	select {
	case log := <-it.logs:
		it.Event = new(AdvancedPoolHooksAllowListRemove)
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

func (it *AdvancedPoolHooksAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListRemoveIterator, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksAllowListRemoveIterator{contract: _AdvancedPoolHooks.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksAllowListRemove)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseAllowListRemove(log types.Log) (*AdvancedPoolHooksAllowListRemove, error) {
	event := new(AdvancedPoolHooksAllowListRemove)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksCCVConfigUpdatedIterator struct {
	Event *AdvancedPoolHooksCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksCCVConfigUpdated)
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

	select {
	case log := <-it.logs:
		it.Event = new(AdvancedPoolHooksCCVConfigUpdated)
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

func (it *AdvancedPoolHooksCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*AdvancedPoolHooksCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksCCVConfigUpdatedIterator{contract: _AdvancedPoolHooks.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksCCVConfigUpdated)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseCCVConfigUpdated(log types.Log) (*AdvancedPoolHooksCCVConfigUpdated, error) {
	event := new(AdvancedPoolHooksCCVConfigUpdated)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksOwnershipTransferRequestedIterator struct {
	Event *AdvancedPoolHooksOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(AdvancedPoolHooksOwnershipTransferRequested)
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

func (it *AdvancedPoolHooksOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedPoolHooksOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksOwnershipTransferRequestedIterator{contract: _AdvancedPoolHooks.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksOwnershipTransferRequested)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseOwnershipTransferRequested(log types.Log) (*AdvancedPoolHooksOwnershipTransferRequested, error) {
	event := new(AdvancedPoolHooksOwnershipTransferRequested)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksOwnershipTransferredIterator struct {
	Event *AdvancedPoolHooksOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(AdvancedPoolHooksOwnershipTransferred)
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

func (it *AdvancedPoolHooksOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedPoolHooksOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksOwnershipTransferredIterator{contract: _AdvancedPoolHooks.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksOwnershipTransferred)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseOwnershipTransferred(log types.Log) (*AdvancedPoolHooksOwnershipTransferred, error) {
	event := new(AdvancedPoolHooksOwnershipTransferred)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type AdvancedPoolHooksThresholdAmountSetIterator struct {
	Event *AdvancedPoolHooksThresholdAmountSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *AdvancedPoolHooksThresholdAmountSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedPoolHooksThresholdAmountSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(AdvancedPoolHooksThresholdAmountSet)
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

func (it *AdvancedPoolHooksThresholdAmountSetIterator) Error() error {
	return it.fail
}

func (it *AdvancedPoolHooksThresholdAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type AdvancedPoolHooksThresholdAmountSet struct {
	ThresholdAmount *big.Int
	Raw             types.Log
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) FilterThresholdAmountSet(opts *bind.FilterOpts) (*AdvancedPoolHooksThresholdAmountSetIterator, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.FilterLogs(opts, "ThresholdAmountSet")
	if err != nil {
		return nil, err
	}
	return &AdvancedPoolHooksThresholdAmountSetIterator{contract: _AdvancedPoolHooks.contract, event: "ThresholdAmountSet", logs: logs, sub: sub}, nil
}

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) WatchThresholdAmountSet(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksThresholdAmountSet) (event.Subscription, error) {

	logs, sub, err := _AdvancedPoolHooks.contract.WatchLogs(opts, "ThresholdAmountSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(AdvancedPoolHooksThresholdAmountSet)
				if err := _AdvancedPoolHooks.contract.UnpackLog(event, "ThresholdAmountSet", log); err != nil {
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

func (_AdvancedPoolHooks *AdvancedPoolHooksFilterer) ParseThresholdAmountSet(log types.Log) (*AdvancedPoolHooksThresholdAmountSet, error) {
	event := new(AdvancedPoolHooksThresholdAmountSet)
	if err := _AdvancedPoolHooks.contract.UnpackLog(event, "ThresholdAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (AdvancedPoolHooksAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (AdvancedPoolHooksAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (AdvancedPoolHooksCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (AdvancedPoolHooksOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (AdvancedPoolHooksOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (AdvancedPoolHooksThresholdAmountSet) Topic() common.Hash {
	return common.HexToHash("0x80dc2a1a49dda9f8bd85c1c376266e011db6448050b7bfd5c2f423e162c11145")
}

func (_AdvancedPoolHooks *AdvancedPoolHooks) Address() common.Address {
	return _AdvancedPoolHooks.address
}

type AdvancedPoolHooksInterface interface {
	CheckAllowList(opts *bind.CallOpts, sender common.Address) error

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetThresholdAmount(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PostFlightCheck(opts *bind.CallOpts, arg0 PoolReleaseOrMintInV1, arg1 *big.Int, arg2 uint16) error

	PreflightCheck(opts *bind.CallOpts, lockOrBurnIn PoolLockOrBurnInV1, arg1 uint16, arg2 []byte) error

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error)

	SetThresholdAmount(opts *bind.TransactOpts, thresholdAmount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*AdvancedPoolHooksAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*AdvancedPoolHooksAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*AdvancedPoolHooksAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*AdvancedPoolHooksCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*AdvancedPoolHooksCCVConfigUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedPoolHooksOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*AdvancedPoolHooksOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedPoolHooksOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*AdvancedPoolHooksOwnershipTransferred, error)

	FilterThresholdAmountSet(opts *bind.FilterOpts) (*AdvancedPoolHooksThresholdAmountSetIterator, error)

	WatchThresholdAmountSet(opts *bind.WatchOpts, sink chan<- *AdvancedPoolHooksThresholdAmountSet) (event.Subscription, error)

	ParseThresholdAmountSet(log types.Log) (*AdvancedPoolHooksThresholdAmountSet, error)

	Address() common.Address
}
