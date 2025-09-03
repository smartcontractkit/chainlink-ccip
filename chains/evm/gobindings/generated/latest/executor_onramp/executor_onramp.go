// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executor_onramp

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

var ExecutorOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowedCCVUpdates\",\"inputs\":[{\"name\":\"ccvsToRemove\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvsToAdd\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvAllowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainUpdates\",\"inputs\":[{\"name\":\"destChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"destChainSelectorsToAdd\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCCVs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"requiredCCVs\",\"type\":\"tuple[]\",\"internalType\":\"structClient.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"optionalCCVs\",\"type\":\"tuple[]\",\"internalType\":\"structClient.CCV[]\",\"components\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxCCVsPerMsg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCCVAllowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMaxCCVsPerMsg\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CCVAdded\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVAllowlistUpdated\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVRemoved\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainAdded\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxCCVsPerMsgSet\",\"inputs\":[{\"name\":\"maxCCVsPerMsg\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxCCVs\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidCCV\",\"inputs\":[{\"name\":\"ccv\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChain\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidMaxPossibleCCVsPerMsg\",\"inputs\":[{\"name\":\"maxPossibleCCVsPerMsg\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x6080346100e157601f6113d738819003918201601f19168301916001600160401b038311848410176100e6578084926020946040528339810103126100e1575160ff8116908181036100e15733156100d0576001549082156100bb576001600160a81b03199091163360ff60a01b19161760a09190911b60ff60a01b16176001556040519081527fcd39dd44d856487a5d3ff100b17da01d09fd38f56a6bc6c1430458ec9cd31bd890602090a16040516112da90816100fd8239f35b82631f3f959360e01b60005260045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c908163181f5a7714610ca157508063240b96e914610b58578063336e545a1461093d578063394058421461085057806379ba5097146107675780638da5cb5b146107155780639dd50723146106d1578063a32845bd1461044e578063a422fdb5146102cb578063a68c61a6146101dd578063f2fde38b146100ea5763fe3b4b1a146100a357600080fd5b346100e55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e557602060ff60015460a01c16604051908152f35b600080fd5b346100e55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e55760043573ffffffffffffffffffffffffffffffffffffffff81168091036100e557610142610eef565b3381146101b357807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e5576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b8181106102b5575050508161025c910382610db0565b6040519182916020830190602084525180915260408301919060005b818110610286575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610278565b8254845260209093019260019283019201610246565b346100e55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e55760043567ffffffffffffffff81116100e55761031a903690600401610df1565b9060243567ffffffffffffffff81116100e55761033b903690600401610df1565b919092610346610eef565b60005b8181106103f35750505060005b81811061035f57005b67ffffffffffffffff61037b610376838587610e3a565b610eda565b1680156103c657908161038f600193611273565b61039b575b5001610356565b7f6e9c954f174a6a41806c1779c207ed29eb3266ba1d60230290dd88ee6a8fb65f600080a284610394565b7f020a07e50000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8067ffffffffffffffff61040d6103766001948688610e3a565b16610417816110e8565b610423575b5001610349565b7ff74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1600080a28661041c565b346100e55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e55760043567ffffffffffffffff81168091036100e55760243567ffffffffffffffff81116100e5577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60a091360301126100e55760443567ffffffffffffffff81116100e5576104f0903690600401610df1565b9160643567ffffffffffffffff81116100e557610511903690600401610df1565b9160843567ffffffffffffffff81116100e557366023820112156100e557806004013567ffffffffffffffff81116100e557369101602401116100e557806000526005602052604060002054156103c657506001549260ff8460a81c166105f6575b505082018092116105c75760a01c60ff169081811161059757602060405160008152f35b7ff2d323530000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60005b85811061068857505060005b828110156105735773ffffffffffffffffffffffffffffffffffffffff610635610630838686610e9a565b610e79565b1661064d816000526003602052604060002054151590565b1561065b5750600101610605565b7fa409d83e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6106ab610630838986610e9a565b166106c3816000526003602052604060002054151590565b1561065b57506001016105f9565b346100e55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e557602060ff60015460a81c166040519015158152f35b346100e55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100e55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e55760005473ffffffffffffffffffffffffffffffffffffffff81163303610826577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e55760043560ff8116908181036100e557610896610eef565b811561090f577fcd39dd44d856487a5d3ff100b17da01d09fd38f56a6bc6c1430458ec9cd31bd8916020917fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff00000000000000000000000000000000000000006001549260a01b16911617600155604051908152a1005b507f1f3f95930000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346100e55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e55760043567ffffffffffffffff81116100e55761098c903690600401610df1565b9060243567ffffffffffffffff81116100e5576109ad903690600401610df1565b9091604435938415158095036100e5576109c5610eef565b60005b818110610ad05750505060005b818110610a6257836001548160ff8260a81c161515036109f157005b816020917fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff75ff0000000000000000000000000000000000000000007fd9e9ee812485edbbfab1d848c2c025cd0d1da3f7b9dcf38edf78c40ec4810ed89560a81b16911617600155604051908152a1005b73ffffffffffffffffffffffffffffffffffffffff610a85610630838587610e3a565b16801561065b579081610a99600193611213565b610aa5575b50016109d5565b7fba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e600080a285610a9e565b80610aff73ffffffffffffffffffffffffffffffffffffffff610af96106306001958789610e3a565b16610f52565b610b0a575b016109c8565b73ffffffffffffffffffffffffffffffffffffffff610b2d610630838688610e3a565b167fbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e600080a2610b04565b346100e55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e557600454610b9381610e22565b90610ba16040519283610db0565b808252610bad81610e22565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208401920136833760045460005b828110610c345783856040519182916020830190602084525180915260408301919060005b818110610c11575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610c03565b600082821015610c74576020816004849352200167ffffffffffffffff6000915416908651831015610c745750600582901b860160200152600101610bde565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b346100e55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e5576040810181811067ffffffffffffffff821117610d8157604052601881527f4578656375746f724f6e52616d7020312e372e302d6465760000000000000000602082015260405190602082528181519182602083015260005b838110610d695750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001610d29565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610d8157604052565b9181601f840112156100e55782359167ffffffffffffffff83116100e5576020808501948460051b0101116100e557565b67ffffffffffffffff8111610d815760051b60200190565b9190811015610e4a5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036100e55790565b9190811015610e4a5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1813603018212156100e5570190565b3567ffffffffffffffff811681036100e55790565b73ffffffffffffffffffffffffffffffffffffffff600154163303610f1057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8054821015610e4a5760005260206000200190600090565b60008181526003602052604090205480156110e1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116105c757600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116105c757818103611072575b5050506002548015611043577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611000816002610f3a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6110c9611083611094936002610f3a565b90549060031b1c9283926002610f3a565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080610fc7565b5050600090565b60008181526005602052604090205480156110e1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116105c757600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116105c7578181036111d9575b5050506004548015611043577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611196816004610f3a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600455600052600560205260006040812055600190565b6111fb6111ea611094936004610f3a565b90549060031b1c9283926004610f3a565b9055600052600560205260406000205538808061115d565b8060005260036020526040600020541560001461126d5760025468010000000000000000811015610d81576112546110948260018594016002556002610f3a565b9055600254906000526003602052604060002055600190565b50600090565b8060005260056020526040600020541560001461126d5760045468010000000000000000811015610d81576112b46110948260018594016004556004610f3a565b905560045490600052600560205260406000205560019056fea164736f6c634300081a000a",
}

var ExecutorOnRampABI = ExecutorOnRampMetaData.ABI

var ExecutorOnRampBin = ExecutorOnRampMetaData.Bin

func DeployExecutorOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, maxCCVsPerMsg uint8) (common.Address, *types.Transaction, *ExecutorOnRamp, error) {
	parsed, err := ExecutorOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorOnRampBin), backend, maxCCVsPerMsg)
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

func (_ExecutorOnRamp *ExecutorOnRampCaller) GetMaxCCVsPerMsg(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "getMaxCCVsPerMsg")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) GetMaxCCVsPerMsg() (uint8, error) {
	return _ExecutorOnRamp.Contract.GetMaxCCVsPerMsg(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) GetMaxCCVsPerMsg() (uint8, error) {
	return _ExecutorOnRamp.Contract.GetMaxCCVsPerMsg(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCaller) IsCCVAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExecutorOnRamp.contract.Call(opts, &out, "isCCVAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ExecutorOnRamp *ExecutorOnRampSession) IsCCVAllowlistEnabled() (bool, error) {
	return _ExecutorOnRamp.Contract.IsCCVAllowlistEnabled(&_ExecutorOnRamp.CallOpts)
}

func (_ExecutorOnRamp *ExecutorOnRampCallerSession) IsCCVAllowlistEnabled() (bool, error) {
	return _ExecutorOnRamp.Contract.IsCCVAllowlistEnabled(&_ExecutorOnRamp.CallOpts)
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

func (_ExecutorOnRamp *ExecutorOnRampTransactor) ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "applyAllowedCCVUpdates", ccvsToRemove, ccvsToAdd, ccvAllowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) ApplyAllowedCCVUpdates(ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyAllowedCCVUpdates(&_ExecutorOnRamp.TransactOpts, ccvsToRemove, ccvsToAdd, ccvAllowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) ApplyAllowedCCVUpdates(ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyAllowedCCVUpdates(&_ExecutorOnRamp.TransactOpts, ccvsToRemove, ccvsToAdd, ccvAllowlistEnabled)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "applyDestChainUpdates", destChainSelectorsToRemove, destChainSelectorsToAdd)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) ApplyDestChainUpdates(destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyDestChainUpdates(&_ExecutorOnRamp.TransactOpts, destChainSelectorsToRemove, destChainSelectorsToAdd)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) ApplyDestChainUpdates(destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []uint64) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.ApplyDestChainUpdates(&_ExecutorOnRamp.TransactOpts, destChainSelectorsToRemove, destChainSelectorsToAdd)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactor) SetMaxCCVsPerMsg(opts *bind.TransactOpts, maxCCVsPerMsg uint8) (*types.Transaction, error) {
	return _ExecutorOnRamp.contract.Transact(opts, "setMaxCCVsPerMsg", maxCCVsPerMsg)
}

func (_ExecutorOnRamp *ExecutorOnRampSession) SetMaxCCVsPerMsg(maxCCVsPerMsg uint8) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.SetMaxCCVsPerMsg(&_ExecutorOnRamp.TransactOpts, maxCCVsPerMsg)
}

func (_ExecutorOnRamp *ExecutorOnRampTransactorSession) SetMaxCCVsPerMsg(maxCCVsPerMsg uint8) (*types.Transaction, error) {
	return _ExecutorOnRamp.Contract.SetMaxCCVsPerMsg(&_ExecutorOnRamp.TransactOpts, maxCCVsPerMsg)
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

type ExecutorOnRampCCVAllowlistUpdatedIterator struct {
	Event *ExecutorOnRampCCVAllowlistUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampCCVAllowlistUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampCCVAllowlistUpdated)
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
		it.Event = new(ExecutorOnRampCCVAllowlistUpdated)
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

func (it *ExecutorOnRampCCVAllowlistUpdatedIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampCCVAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampCCVAllowlistUpdated struct {
	Enabled bool
	Raw     types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterCCVAllowlistUpdated(opts *bind.FilterOpts) (*ExecutorOnRampCCVAllowlistUpdatedIterator, error) {

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "CCVAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampCCVAllowlistUpdatedIterator{contract: _ExecutorOnRamp.contract, event: "CCVAllowlistUpdated", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchCCVAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVAllowlistUpdated) (event.Subscription, error) {

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "CCVAllowlistUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampCCVAllowlistUpdated)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVAllowlistUpdated", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseCCVAllowlistUpdated(log types.Log) (*ExecutorOnRampCCVAllowlistUpdated, error) {
	event := new(ExecutorOnRampCCVAllowlistUpdated)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "CCVAllowlistUpdated", log); err != nil {
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

type ExecutorOnRampMaxCCVsPerMsgSetIterator struct {
	Event *ExecutorOnRampMaxCCVsPerMsgSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ExecutorOnRampMaxCCVsPerMsgSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutorOnRampMaxCCVsPerMsgSet)
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
		it.Event = new(ExecutorOnRampMaxCCVsPerMsgSet)
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

func (it *ExecutorOnRampMaxCCVsPerMsgSetIterator) Error() error {
	return it.fail
}

func (it *ExecutorOnRampMaxCCVsPerMsgSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ExecutorOnRampMaxCCVsPerMsgSet struct {
	MaxCCVsPerMsg uint8
	Raw           types.Log
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) FilterMaxCCVsPerMsgSet(opts *bind.FilterOpts) (*ExecutorOnRampMaxCCVsPerMsgSetIterator, error) {

	logs, sub, err := _ExecutorOnRamp.contract.FilterLogs(opts, "MaxCCVsPerMsgSet")
	if err != nil {
		return nil, err
	}
	return &ExecutorOnRampMaxCCVsPerMsgSetIterator{contract: _ExecutorOnRamp.contract, event: "MaxCCVsPerMsgSet", logs: logs, sub: sub}, nil
}

func (_ExecutorOnRamp *ExecutorOnRampFilterer) WatchMaxCCVsPerMsgSet(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampMaxCCVsPerMsgSet) (event.Subscription, error) {

	logs, sub, err := _ExecutorOnRamp.contract.WatchLogs(opts, "MaxCCVsPerMsgSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ExecutorOnRampMaxCCVsPerMsgSet)
				if err := _ExecutorOnRamp.contract.UnpackLog(event, "MaxCCVsPerMsgSet", log); err != nil {
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

func (_ExecutorOnRamp *ExecutorOnRampFilterer) ParseMaxCCVsPerMsgSet(log types.Log) (*ExecutorOnRampMaxCCVsPerMsgSet, error) {
	event := new(ExecutorOnRampMaxCCVsPerMsgSet)
	if err := _ExecutorOnRamp.contract.UnpackLog(event, "MaxCCVsPerMsgSet", log); err != nil {
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

func (ExecutorOnRampCCVAdded) Topic() common.Hash {
	return common.HexToHash("0xba540b0c7a674c7f1716e91e0e0a2390ebb27755267c72e0807812b93f3bf00e")
}

func (ExecutorOnRampCCVAllowlistUpdated) Topic() common.Hash {
	return common.HexToHash("0xd9e9ee812485edbbfab1d848c2c025cd0d1da3f7b9dcf38edf78c40ec4810ed8")
}

func (ExecutorOnRampCCVRemoved) Topic() common.Hash {
	return common.HexToHash("0xbc743a2d04de950d86944633fbe825e492514eef584678e9fa97f3e939cf605e")
}

func (ExecutorOnRampDestChainAdded) Topic() common.Hash {
	return common.HexToHash("0x6e9c954f174a6a41806c1779c207ed29eb3266ba1d60230290dd88ee6a8fb65f")
}

func (ExecutorOnRampDestChainRemoved) Topic() common.Hash {
	return common.HexToHash("0xf74668182f6a521d1a362a6bbc8344cac3a467bab207cdabbaf39e503edef6a1")
}

func (ExecutorOnRampMaxCCVsPerMsgSet) Topic() common.Hash {
	return common.HexToHash("0xcd39dd44d856487a5d3ff100b17da01d09fd38f56a6bc6c1430458ec9cd31bd8")
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

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 ClientEVM2AnyMessage, requiredCCVs []ClientCCV, optionalCCVs []ClientCCV, arg4 []byte) (*big.Int, error)

	GetMaxCCVsPerMsg(opts *bind.CallOpts) (uint8, error)

	IsCCVAllowlistEnabled(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowedCCVUpdates(opts *bind.TransactOpts, ccvsToRemove []common.Address, ccvsToAdd []common.Address, ccvAllowlistEnabled bool) (*types.Transaction, error)

	ApplyDestChainUpdates(opts *bind.TransactOpts, destChainSelectorsToRemove []uint64, destChainSelectorsToAdd []uint64) (*types.Transaction, error)

	SetMaxCCVsPerMsg(opts *bind.TransactOpts, maxCCVsPerMsg uint8) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterCCVAdded(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVAddedIterator, error)

	WatchCCVAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVAdded, ccv []common.Address) (event.Subscription, error)

	ParseCCVAdded(log types.Log) (*ExecutorOnRampCCVAdded, error)

	FilterCCVAllowlistUpdated(opts *bind.FilterOpts) (*ExecutorOnRampCCVAllowlistUpdatedIterator, error)

	WatchCCVAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVAllowlistUpdated) (event.Subscription, error)

	ParseCCVAllowlistUpdated(log types.Log) (*ExecutorOnRampCCVAllowlistUpdated, error)

	FilterCCVRemoved(opts *bind.FilterOpts, ccv []common.Address) (*ExecutorOnRampCCVRemovedIterator, error)

	WatchCCVRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampCCVRemoved, ccv []common.Address) (event.Subscription, error)

	ParseCCVRemoved(log types.Log) (*ExecutorOnRampCCVRemoved, error)

	FilterDestChainAdded(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainAddedIterator, error)

	WatchDestChainAdded(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainAdded, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainAdded(log types.Log) (*ExecutorOnRampDestChainAdded, error)

	FilterDestChainRemoved(opts *bind.FilterOpts, destChainSelector []uint64) (*ExecutorOnRampDestChainRemovedIterator, error)

	WatchDestChainRemoved(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampDestChainRemoved, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainRemoved(log types.Log) (*ExecutorOnRampDestChainRemoved, error)

	FilterMaxCCVsPerMsgSet(opts *bind.FilterOpts) (*ExecutorOnRampMaxCCVsPerMsgSetIterator, error)

	WatchMaxCCVsPerMsgSet(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampMaxCCVsPerMsgSet) (event.Subscription, error)

	ParseMaxCCVsPerMsgSet(log types.Log) (*ExecutorOnRampMaxCCVsPerMsgSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ExecutorOnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExecutorOnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutorOnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ExecutorOnRampOwnershipTransferred, error)

	Address() common.Address
}
