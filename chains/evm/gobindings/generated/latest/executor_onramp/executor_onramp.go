// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executor_onramp

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type ClientCCV struct {
	CcvAddress common.Address
	Args       []byte
}

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type ExecutorOnRampDynamicConfig struct {
	MaxCCVsPerMsg uint8
}

var ExecutorOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structExecutorOnRamp.DynamicConfig\",\"components\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowedCCVUpdates\",\"inputs\":[{\"name\":\"ccvsToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvsToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainUpdates\",\"inputs\":[{\"name\":\"destChainSelectorsToAdd\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"destChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCCVs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structExecutorOnRamp.DynamicConfig\",\"components\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"requiredCCVs\",\"type\":\"tuple[]\",\"internalType\":\"structClient.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"optionalCCVs\",\"type\":\"tuple[]\",\"internalType\":\"structClient.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_allowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structExecutorOnRamp.DynamicConfig\",\"components\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllowlistUpdated\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVAdded\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVRemoved\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structExecutorOnRamp.DynamicConfig\",\"components\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxCCVs\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMaxPossibleCCVsPerMsg\",\"inputs\":[{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60806040523461011857604051601f61142f38819003918201601f19168301916001600160401b0383118484101761011d578084926020946040528339810103126101185760405190600090602083016001600160401b03811184821017610104576040525160ff8116810361010057825233156100f157600180546001600160a01b03191633179055815160ff1680156100de577f633b135fc0b23f6b6cd84d99d4cb17859fc1ad99a42dce5f90f68b4e1d6bd432602060ff8551168060ff196006541617600655604051908152a16040516112fb90816101348239f35b631f3f959360e01b825260045260249150fd5b639b15e16f60e01b8152600490fd5b5080fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c908163181f5a7714610ca657508063240b96e914610b5d578063336e545a146109565780637437ff9f146108f957806379ba5097146108105780638da5cb5b146107be578063a32845bd1461053b578063a422fdb5146103b8578063a68c61a6146102ca578063d5d9108314610286578063e27717c71461019b5763f2fde38b146100a357600080fd5b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043573ffffffffffffffffffffffffffffffffffffffff8116809103610196576100fb610f10565b33811461016c57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101965760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760006040516101d881610db5565b60043560ff811681036102825781526101ef610f10565b60ff81511680156102575750602060ff7f633b135fc0b23f6b6cd84d99d4cb17859fc1ad99a42dce5f90f68b4e1d6bd432925116807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006006541617600655604051908152a180f35b7f1f3f9593000000000000000000000000000000000000000000000000000000008352600452602482fd5b8280fd5b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657602060ff60015460a01c166040519015158152f35b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610196576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b8181106103a25750505081610349910382610dd1565b6040519182916020830190602084525180915260408301919060005b818110610373575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610365565b8254845260209093019260019283019201610333565b346101965760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff811161019657610407903690600401610e12565b9060243567ffffffffffffffff811161019657610428903690600401610e12565b919092610433610f10565b60005b8181106104ac5750505060005b81811061044c57005b8067ffffffffffffffff61046b6104666001948688610e5b565b610efb565b16610475816111c3565b610481575b5001610443565b7ff74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1600080a28461047a565b67ffffffffffffffff6104c3610466838587610e5b565b16801561050e5790816104d7600193611008565b6104e3575b5001610436565b7f6e9c954f174a6a41806c1779c207ed29eb3266ba1d60230290dd88ee6a8fb65f600080a2866104dc565b7f020a07e50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101965760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff81168091036101965760243567ffffffffffffffff8111610196577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60a091360301126101965760443567ffffffffffffffff8111610196576105dd903690600401610e12565b9160643567ffffffffffffffff8111610196576105fe903690600401610e12565b92909160843567ffffffffffffffff8111610196573660238201121561019657806004013567ffffffffffffffff81116101965736910160240111610196578060005260056020526040600020541561050e575060ff60015460a01c166106e3575b505081018091116106b45760ff600654169081811161068457602060405160008152f35b7ff2d323530000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60005b84811061077557505060005b828110156106605773ffffffffffffffffffffffffffffffffffffffff61072261071d838686610ebb565b610e9a565b1661073a816000526003602052604060002054151590565b1561074857506001016106f2565b7fa409d83e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff61079861071d838886610ebb565b166107b0816000526003602052604060002054151590565b1561074857506001016106e6565b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760005473ffffffffffffffffffffffffffffffffffffffff811633036108cf577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657600060405161093681610db5565b52602060405161094581610db5565b60ff60065416809152604051908152f35b346101965760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101965760043567ffffffffffffffff8111610196576109a5903690600401610e12565b9060243567ffffffffffffffff8111610196576109c6903690600401610e12565b909160443593841515809503610196576109de610f10565b60005b818110610aef5750505060005b818110610a67577f2009642e712cfd664bd243fbca19fbf294816da9b9d1f7f94942e2f9a299fcc76020856001547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000008360a01b16911617600155604051908152a1005b80610a9673ffffffffffffffffffffffffffffffffffffffff610a9061071d6001958789610e5b565b16611062565b610aa1575b016109ee565b73ffffffffffffffffffffffffffffffffffffffff610ac461071d838688610e5b565b167fbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e600080a2610a9b565b73ffffffffffffffffffffffffffffffffffffffff610b1261071d838587610e5b565b168015610748579081610b26600193610f73565b610b32575b50016109e1565b7fba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e600080a287610b2b565b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019657600454610b9881610e43565b90610ba66040519283610dd1565b808252610bb281610e43565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208401920136833760045460005b828110610c395783856040519182916020830190602084525180915260408301919060005b818110610c16575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610c08565b600082821015610c79576020816004849352200167ffffffffffffffff6000915416908651831015610c795750600582901b860160200152600101610be3565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b346101965760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610196576040810181811067ffffffffffffffff821117610d8657604052601881527f4578656375746f724f6e52616d7020312e372e302d6465760000000000000000602082015260405190602082528181519182602083015260005b838110610d6e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001610d2e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117610d8657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610d8657604052565b9181601f840112156101965782359167ffffffffffffffff8311610196576020808501948460051b01011161019657565b67ffffffffffffffff8111610d865760051b60200190565b9190811015610e6b5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036101965790565b9190811015610e6b5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc181360301821215610196570190565b3567ffffffffffffffff811681036101965790565b73ffffffffffffffffffffffffffffffffffffffff600154163303610f3157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8054821015610e6b5760005260206000200190600090565b806000526003602052604060002054156000146110025760025468010000000000000000811015610d8657610fe9610fb48260018594016002556002610f5b565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b806000526005602052604060002054156000146110025760045468010000000000000000811015610d8657611049610fb48260018594016004556004610f5b565b9055600454906000526005602052604060002055600190565b60008181526003602052604090205480156111bc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116106b457600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116106b457818103611182575b5050506002548015611153577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611110816002610f5b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6111a4611193610fb4936002610f5b565b90549060031b1c9283926002610f5b565b905560005260036020526040600020553880806110d7565b5050600090565b60008181526005602052604090205480156111bc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116106b457600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116106b4578181036112b4575b5050506004548015611153577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611271816004610f5b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600455600052600560205260006040812055600190565b6112d66112c5610fb4936004610f5b565b90549060031b1c9283926004610f5b565b9055600052600560205260406000205538808061123856fea164736f6c634300081a000a",
}

var ExecutorOnRampABI = ExecutorOnRampMetaData.ABI

var ExecutorOnRampBin = ExecutorOnRampMetaData.Bin

func DeployExecutorOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, dynamicConfig ExecutorOnRampDynamicConfig) (common.Address, *types.Transaction, *ExecutorOnRamp, error) {
	parsed, err := ExecutorOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorOnRampBin), backend, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutorOnRamp{address: address, abi: *parsed, ExecutorOnRampCaller: ExecutorOnRampCaller{contract: contract}, ExecutorOnRampTransactor: ExecutorOnRampTransactor{contract: contract}, ExecutorOnRampFilterer: ExecutorOnRampFilterer{contract: contract}}, nil
}

type ExecutorOnRamp struct {
	address common.Address
	abi     abi.ABI
	ExecutorOnRampCaller
	ExecutorOnRampTransactor
	ExecutorOnRampFilterer
}

type ExecutorOnRampCaller struct {
	contract *bind.BoundContract
}

type ExecutorOnRampTransactor struct {
	contract *bind.BoundContract
}

type ExecutorOnRampFilterer struct {
	contract *bind.BoundContract
}

type ExecutorOnRampSession struct {
	Contract     *ExecutorOnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ExecutorOnRampCallerSession struct {
	Contract *ExecutorOnRampCaller
	CallOpts bind.CallOpts
}

type ExecutorOnRampTransactorSession struct {
	Contract     *ExecutorOnRampTransactor
	TransactOpts bind.TransactOpts
}

type ExecutorOnRampRaw struct {
	Contract *ExecutorOnRamp
}

type ExecutorOnRampCallerRaw struct {
	Contract *ExecutorOnRampCaller
}

type ExecutorOnRampTransactorRaw struct {
	Contract *ExecutorOnRampTransactor
}

func NewExecutorOnRamp(address common.Address, backend bind.ContractBackend) (*ExecutorOnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(ExecutorOnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindExecutorOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRamp{address: address, abi: abi, ExecutorOnRampCaller: ExecutorOnRampCaller{contract: contract}, ExecutorOnRampTransactor: ExecutorOnRampTransactor{contract: contract}, ExecutorOnRampFilterer: ExecutorOnRampFilterer{contract: contract}}, nil
}

func NewExecutorOnRampCaller(address common.Address, caller bind.ContractCaller) (*ExecutorOnRampCaller, error) {
	contract, err := bindExecutorOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampCaller{contract: contract}, nil
}

func NewExecutorOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorOnRampTransactor, error) {
	contract, err := bindExecutorOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampTransactor{contract: contract}, nil
}

func NewExecutorOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorOnRampFilterer, error) {
	contract, err := bindExecutorOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampFilterer{contract: contract}, nil
}

func bindExecutorOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutorOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ExecutorOnRamp *ExecutorOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorOnRamp.Contract.ExecutorOnRampCaller.contract.Call(opts, result, method, params...)
}

func (_ExecutorOnRamp *ExecutorOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ExecutorOnRampTransactor.contract.Transfer(opts)
}

func (_ExecutorOnRamp *ExecutorOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ExecutorOnRampTransactor.contract.Transact(opts, method, params...)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutorOnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.contract.Transfer(opts)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetAllowedCCVs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getAllowedCCVs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetAllowedCCVs() ([]common.Address, error) {
	return _ExecutorOnRamp.Contract.GetAllowedCCVs(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetAllowedCCVs() ([]common.Address, error) {
	return _ExecutorOnRamp.Contract.GetAllowedCCVs(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetDestChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getDestChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetDestChains() ([]uint64, error) {
	return _ExecutorOnRamp.Contract.GetDestChains(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetDestChains() ([]uint64, error) {
	return _ExecutorOnRamp.Contract.GetDestChains(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (ExecutorOnRampDynamicConfig, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(ExecutorOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ExecutorOnRampDynamicConfig)).(*ExecutorOnRampDynamicConfig)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetDynamicConfig() (ExecutorOnRampDynamicConfig, error) {
	return _ExecutorOnRamp.Contract.GetDynamicConfig(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetDynamicConfig() (ExecutorOnRampDynamicConfig, error) {
	return _ExecutorOnRamp.Contract.GetDynamicConfig(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, requiredCCVs []ClientCCV, optionalCCVs []ClientCCV, arg4 []byte) (*big.Int, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getFee", destChainSelector, arg1, requiredCCVs, optionalCCVs, arg4)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, requiredCCVs []ClientCCV, optionalCCVs []ClientCCV, arg4 []byte) (*big.Int, error) {
	return _ExecutorOnRamp.Contract.GetFee(&_ExecutorOnRamp.CallOpts, destChainSelector, arg1, requiredCCVs, optionalCCVs, arg4)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetFee(destChainSelector uint64, arg1 ClientEVM2AnyMessage, requiredCCVs []ClientCCV, optionalCCVs []ClientCCV, arg4 []byte) (*big.Int, error) {
	return _ExecutorOnRamp.Contract.GetFee(&_ExecutorOnRamp.CallOpts, destChainSelector, arg1, requiredCCVs, optionalCCVs, arg4)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) Owner() (common.Address, error) {
	return _ExecutorOnRamp.Contract.Owner(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) Owner() (common.Address, error) {
	return _ExecutorOnRamp.Contract.Owner(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) SAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "s_allowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) SAllowlistEnabled() (bool, error) {
	return _ExecutorOnRamp.Contract.SAllowlistEnabled(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) SAllowlistEnabled() (bool, error) {
	return _ExecutorOnRamp.Contract.SAllowlistEnabled(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) TypeAndVersion() (string, error) {
	return _ExecutorOnRamp.Contract.TypeAndVersion(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) TypeAndVersion() (string, error) {
	return _ExecutorOnRamp.Contract.TypeAndVersion(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_ExecutorOnRamp *ExecutorOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.AcceptOwnership(&_ExecutorOnRamp.TransactOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.AcceptOwnership(&_ExecutorOnRamp.TransactOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToAdd []common.Address, ccvsToRemove []common.Address, allowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "applyAllowedCCVUpdates", ccvsToAdd, ccvsToRemove, allowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) ApplyAllowedCCVUpdates(ccvsToAdd []common.Address, ccvsToRemove []common.Address, allowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyAllowedCCVUpdates(&_ExecutorOnRamp.TransactOpts, ccvsToAdd, ccvsToRemove, allowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) ApplyAllowedCCVUpdates(ccvsToAdd []common.Address, ccvsToRemove []common.Address, allowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyAllowedCCVUpdates(&_ExecutorOnRamp.TransactOpts, ccvsToAdd, ccvsToRemove, allowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToAdd []uint64, destChainSelectorsToRemove []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "applyDestChainUpdates", destChainSelectorsToAdd, destChainSelectorsToRemove)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) ApplyDestChainUpdates(destChainSelectorsToAdd []uint64, destChainSelectorsToRemove []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyDestChainUpdates(&_ExecutorOnRamp.TransactOpts, destChainSelectorsToAdd, destChainSelectorsToRemove)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) ApplyDestChainUpdates(destChainSelectorsToAdd []uint64, destChainSelectorsToRemove []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyDestChainUpdates(&_ExecutorOnRamp.TransactOpts, destChainSelectorsToAdd, destChainSelectorsToRemove)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig ExecutorOnRampDynamicConfig) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) SetDynamicConfig(dynamicConfig ExecutorOnRampDynamicConfig) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.SetDynamicConfig(&_ExecutorOnRamp.TransactOpts, dynamicConfig)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) SetDynamicConfig(dynamicConfig ExecutorOnRampDynamicConfig) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.SetDynamicConfig(&_ExecutorOnRamp.TransactOpts, dynamicConfig)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.TransferOwnership(&_ExecutorOnRamp.TransactOpts, to)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.TransferOwnership(&_ExecutorOnRamp.TransactOpts, to)
}

type ExecutorOnRampAllowlistUpdatedIterator struct {
	Event *ExecutorOnRampAllowlistUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampAllowlistUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampAllowlistUpdated)
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
		it.Event = new(ExecutorOnRampAllowlistUpdated)
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

func (it *ExecutorOnRampAllowlistUpdatedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampAllowlistUpdated struct {
	Enabled bool
	Raw     types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterAllowlistUpdated(opts *bind.FilterOpts) (*ExecutorOnRampAllowlistUpdatedIterator, error) {

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "AllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampAllowlistUpdatedIterator{contract: _ExecutorOnRamp.contract, event: "AllowlistUpdated", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampAllowlistUpdated) (event.Subscription, error) {

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "AllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampAllowlistUpdated)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "AllowlistUpdated", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseAllowlistUpdated(log types.Log) (*ExecutorOnRampAllowlistUpdated, error) {
	event := new(ExecutorOnRampAllowlistUpdated)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "AllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampCCVAddedIterator struct {
	Event *ExecutorOnRampCCVAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampCCVAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampCCVAdded)
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
		it.Event = new(ExecutorOnRampCCVAdded)
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

func (it *ExecutorOnRampCCVAddedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampCCVAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampCCVAdded struct {
	Ccv common.Address
	Raw types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterCCVAdded(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVAddedIterator, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "CCVAdded", ccvRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampCCVAddedIterator{contract: _ExecutorOnRamp.contract, event: "CCVAdded", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchCCVAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVAdded, ccv []common.Address) (event.Subscription, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "CCVAdded", ccvRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampCCVAdded)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVAdded", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseCCVAdded(log types.Log) (*ExecutorOnRampCCVAdded, error) {
	event := new(ExecutorOnRampCCVAdded)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampCCVRemovedIterator struct {
	Event *ExecutorOnRampCCVRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampCCVRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampCCVRemoved)
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
		it.Event = new(ExecutorOnRampCCVRemoved)
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

func (it *ExecutorOnRampCCVRemovedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampCCVRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampCCVRemoved struct {
	Ccv common.Address
	Raw types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterCCVRemoved(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVRemovedIterator, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "CCVRemoved", ccvRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampCCVRemovedIterator{contract: _ExecutorOnRamp.contract, event: "CCVRemoved", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchCCVRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVRemoved, ccv []common.Address) (event.Subscription, error) {

	var ccvRule []interface{}
	for _, ccvItem := range ccv {
		ccvRule = append(ccvRule, ccvItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "CCVRemoved", ccvRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampCCVRemoved)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVRemoved", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseCCVRemoved(log types.Log) (*ExecutorOnRampCCVRemoved, error) {
	event := new(ExecutorOnRampCCVRemoved)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampConfigSetIterator struct {
	Event *ExecutorOnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampConfigSet)
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
		it.Event = new(ExecutorOnRampConfigSet)
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

func (it *ExecutorOnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampConfigSet struct {
	DynamicConfig ExecutorOnRampDynamicConfig
	Raw           types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*ExecutorOnRampConfigSetIterator, error) {

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampConfigSetIterator{contract: _ExecutorOnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampConfigSet)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseConfigSet(log types.Log) (*ExecutorOnRampConfigSet, error) {
	event := new(ExecutorOnRampConfigSet)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampDestChainAddedIterator struct {
	Event *ExecutorOnRampDestChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampDestChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampDestChainAdded)
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
		it.Event = new(ExecutorOnRampDestChainAdded)
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

func (it *ExecutorOnRampDestChainAddedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampDestChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampDestChainAdded struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainAddedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampDestChainAddedIterator{contract: _ExecutorOnRamp.contract, event: "DestChainAdded", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainAdded, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "DestChainAdded", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampDestChainAdded)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseDestChainAdded(log types.Log) (*ExecutorOnRampDestChainAdded, error) {
	event := new(ExecutorOnRampDestChainAdded)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "DestChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampDestChainRemovedIterator struct {
	Event *ExecutorOnRampDestChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampDestChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampDestChainRemoved)
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
		it.Event = new(ExecutorOnRampDestChainRemoved)
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

func (it *ExecutorOnRampDestChainRemovedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampDestChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampDestChainRemoved struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterDestChainRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainRemovedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "DestChainRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampDestChainRemovedIterator{contract: _ExecutorOnRamp.contract, event: "DestChainRemoved", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchDestChainRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainRemoved, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "DestChainRemoved", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampDestChainRemoved)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "DestChainRemoved", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseDestChainRemoved(log types.Log) (*ExecutorOnRampDestChainRemoved, error) {
	event := new(ExecutorOnRampDestChainRemoved)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "DestChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampOwnershipTransferRequestedIterator struct {
	Event *ExecutorOnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampOwnershipTransferRequested)
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
		it.Event = new(ExecutorOnRampOwnershipTransferRequested)
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

func (it *ExecutorOnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampOwnershipTransferRequestedIterator{contract: _ExecutorOnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampOwnershipTransferRequested)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*ExecutorOnRampOwnershipTransferRequested, error) {
	event := new(ExecutorOnRampOwnershipTransferRequested)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ExecutorOnRampOwnershipTransferredIterator struct {
	Event *ExecutorOnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampOwnershipTransferred)
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
		it.Event = new(ExecutorOnRampOwnershipTransferred)
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

func (it *ExecutorOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampOwnershipTransferredIterator{contract: _ExecutorOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampOwnershipTransferred)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*ExecutorOnRampOwnershipTransferred, error) {
	event := new(ExecutorOnRampOwnershipTransferred)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ExecutorOnRamp *ExecutorOnRamp) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ExecutorOnRamp.abi.Events["AllowlistUpdated"].ID:
		return _ExecutorOnRamp.ParseAllowlistUpdated(log)
	case _ExecutorOnRamp.abi.Events["CCVAdded"].ID:
		return _ExecutorOnRamp.ParseCCVAdded(log)
	case _ExecutorOnRamp.abi.Events["CCVRemoved"].ID:
		return _ExecutorOnRamp.ParseCCVRemoved(log)
	case _ExecutorOnRamp.abi.Events["ConfigSet"].ID:
		return _ExecutorOnRamp.ParseConfigSet(log)
	case _ExecutorOnRamp.abi.Events["DestChainAdded"].ID:
		return _ExecutorOnRamp.ParseDestChainAdded(log)
	case _ExecutorOnRamp.abi.Events["DestChainRemoved"].ID:
		return _ExecutorOnRamp.ParseDestChainRemoved(log)
	case _ExecutorOnRamp.abi.Events["OwnershipTransferRequested"].ID:
		return _ExecutorOnRamp.ParseOwnershipTransferRequested(log)
	case _ExecutorOnRamp.abi.Events["OwnershipTransferred"].ID:
		return _ExecutorOnRamp.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ExecutorOnRampAllowlistUpdated) Topic() common.Hash {
	return common.HexToHash("0x2009642e712cfd664bd243fbca19fbf294816da9b9d1f7f94942e2f9a299fcc7")
}

func (ExecutorOnRampCCVAdded) Topic() common.Hash {
	return common.HexToHash("0xba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e")
}

func (ExecutorOnRampCCVRemoved) Topic() common.Hash {
	return common.HexToHash("0xbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e")
}

func (ExecutorOnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x633b135fc0b23f6b6cd84d99d4cb17859fc1ad99a42dce5f90f68b4e1d6bd432")
}

func (ExecutorOnRampDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x6e9c954f174a6a41806c1779c207ed29eb3266ba1d60230290dd88ee6a8fb65f")
}

func (ExecutorOnRampDestChainRemoved) Topic() common.Hash {
	return common.HexToHash("0xf74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1")
}

func (ExecutorOnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ExecutorOnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ExecutorOnRamp *ExecutorOnRamp) Address() common.Address {
	return _ExecutorOnRamp.address
}

type ExecutorOnRampInterface interface {
	GetAllowedCCVs(opts *bind.CallOpts) ([]common.Address, error)

	GetDestChains(opts *bind.CallOpts) ([]uint64, error)

	GetDynamicConfig(opts *bind.CallOpts) (ExecutorOnRampDynamicConfig, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, requiredCCVs []ClientCCV, optionalCCVs []ClientCCV, arg4 []byte) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToAdd []common.Address, ccvsToRemove []common.Address, allowlistEnabled bool) (*types.Transaction, error)

	ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToAdd []uint64, destChainSelectorsToRemove []uint64) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig ExecutorOnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowlistUpdated(opts *bind.FilterOpts) (*ExecutorOnRampAllowlistUpdatedIterator, error)

	WatchAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampAllowlistUpdated) (event.Subscription, error)

	ParseAllowlistUpdated(log types.Log) (*ExecutorOnRampAllowlistUpdated, error)

	FilterCCVAdded(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVAddedIterator, error)

	WatchCCVAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVAdded, ccv []common.Address) (event.Subscription, error)

	ParseCCVAdded(log types.Log) (*ExecutorOnRampCCVAdded, error)

	FilterCCVRemoved(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVRemovedIterator, error)

	WatchCCVRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVRemoved, ccv []common.Address) (event.Subscription, error)

	ParseCCVRemoved(log types.Log) (*ExecutorOnRampCCVRemoved, error)

	FilterConfigSet(opts *bind.FilterOpts) (*ExecutorOnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*ExecutorOnRampConfigSet, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*ExecutorOnRampDestChainAdded, error)

	FilterDestChainRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainRemovedIterator, error)

	WatchDestChainRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainRemoved(log types.Log) (*ExecutorOnRampDestChainRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ExecutorOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ExecutorOnRampOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
