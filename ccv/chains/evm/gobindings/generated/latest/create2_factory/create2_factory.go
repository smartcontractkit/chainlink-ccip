// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package create2_factory

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

var CREATE2FactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"allowList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"computeAddress\",\"inputs\":[{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAndCall\",\"inputs\":[{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calls\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAndTransferOwnership\",\"inputs\":[{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractDeployed\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"CallerNotAllowed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x6080604052346101d75761132f80380380610019816101f2565b9283398101906020818303126101d7578051906001600160401b0382116101d7570181601f820112156101d7578051916001600160401b0383116101dc578260051b9160208061006a8186016101f2565b8096815201938201019182116101d757602001915b8183106101b7578333156101a657600180546001600160a01b031916331790556100a960206101f2565b60008152600036813760005b8151811015610124576001906100dd6001600160a01b036100d68386610217565b5116610259565b6100e8575b016100b5565b818060a01b036100f88285610217565b51167f50c35a67b454d38c20800d5b55e320f58f4c9c86a28d8ab20f03045d1a38d99a600080a26100e2565b8260005b8151811015610197576001906101506001600160a01b036101498386610217565b5116610357565b61015b575b01610128565b818060a01b0361016b8285610217565b51167ff7762e85af7b409451f9a76004c5f755642902434eb11351ae67eb9746888b69600080a2610155565b604051610f7790816103b88239f35b639b15e16f60e01b60005260046000fd5b82516001600160a01b03811681036101d75781526020928301920161007f565b600080fd5b634e487b7160e01b600052604160045260246000fd5b6040519190601f01601f191682016001600160401b038111838210176101dc57604052565b805182101561022b5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561022b5760005260206000200190600090565b600081815260036020526040902054801561035057600019810181811161033a5760025460001981019190821161033a578181036102e9575b50505060025480156102d357600019016102ad816002610241565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6103226102fa61030b936002610241565b90549060031b1c9283926002610241565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610292565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146103b157600254680100000000000000008110156101dc5761039861030b8260018594016002556002610241565b9055600254906000526003602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a77146108305780633e4b9f7a1461079e57806354c8a4f3146105ef57806379ba50971461050657806381b72a66146103c45780638da5cb5b14610372578063a7cd63b714610284578063e4a9848d146101775763f2fde38b1461007f57600080fd5b346101725760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725760043573ffffffffffffffffffffffffffffffffffffffff8116809103610172576100d7610b3e565b33811461014857807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101725760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725760043567ffffffffffffffff8111610172576101c6903690600401610a40565b906044359167ffffffffffffffff831161017257366023840112156101725782600401356101f381610a6e565b9361020160405195866108b1565b8185526024602086019260051b820101903682116101725760248101925b82841061025457602061023688602435888a610b89565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b833567ffffffffffffffff8111610172576020916102798392602436918701016109f1565b81520193019261021f565b346101725760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610172576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061035c57505050816103039103826108b1565b6040519182916020830190602084525180915260408301919060005b81811061032d575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff1684528594506020938401939092019160010161031f565b82548452602090930192600192830192016102ed565b346101725760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017257602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101725760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725760043567ffffffffffffffff811161017257610413903690600401610a40565b60443573ffffffffffffffffffffffffffffffffffffffff81168091036101725760209260409283519261044785856108b1565b600184527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0850160005b8181106104f75750509173ffffffffffffffffffffffffffffffffffffffff93916104ef938651907ff2fde38b00000000000000000000000000000000000000000000000000000000898301526024820152602481526104d26044826108b1565b6104db84610aee565b526104e583610aee565b5060243591610b89565b915191168152f35b60608682018901528701610471565b346101725760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725760005473ffffffffffffffffffffffffffffffffffffffff811633036105c5577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101725760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725760043567ffffffffffffffff81116101725761063e903690600401610a0f565b6024359067ffffffffffffffff82116101725761067761066561067f933690600401610a0f565b949092610670610b3e565b3691610a86565b923691610a86565b60005b825181101561070d57806106b773ffffffffffffffffffffffffffffffffffffffff6106b060019487610b2a565b5116610d45565b6106c2575b01610682565b73ffffffffffffffffffffffffffffffffffffffff6106e18286610b2a565b51167f50c35a67b454d38c20800d5b55e320f58f4c9c86a28d8ab20f03045d1a38d99a600080a26106bc565b5060005b815181101561079c578061074673ffffffffffffffffffffffffffffffffffffffff61073f60019486610b2a565b5116610f0a565b610751575b01610711565b73ffffffffffffffffffffffffffffffffffffffff6107708285610b2a565b51167ff7762e85af7b409451f9a76004c5f755642902434eb11351ae67eb9746888b69600080a261074b565b005b346101725760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725760043567ffffffffffffffff8111610172576055600b6107f360209336906004016109f1565b838151910120604051906040820152602435848201523081520160ff81532073ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101725760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610172576108ad604080519061087181836108b1565b601482527f43524541544532466163746f727920312e372e3000000000000000000000000060208301525191829160208352602083019061095b565b0390f35b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176108f257604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b67ffffffffffffffff81116108f257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106109a55750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610966565b9291926109c682610921565b916109d460405193846108b1565b829481845281830111610172578281602093846000960137010152565b9080601f8301121561017257816020610a0c933591016109ba565b90565b9181601f840112156101725782359167ffffffffffffffff8311610172576020808501948460051b01011161017257565b9181601f840112156101725782359167ffffffffffffffff8311610172576020838186019501011161017257565b67ffffffffffffffff81116108f25760051b60200190565b9291610a9182610a6e565b93610a9f60405195866108b1565b602085848152019260051b810191821161017257915b818310610ac157505050565b823573ffffffffffffffffffffffffffffffffffffffff8116810361017257815260209283019201610ab5565b805115610afb5760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051821015610afb5760209160051b010190565b73ffffffffffffffffffffffffffffffffffffffff600154163303610b5f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b33600052600360205260406000205415610cff57610ba89136916109ba565b805115610cd5576020815191016000f59073ffffffffffffffffffffffffffffffffffffffff82168015610cab577f8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf600080a260005b8151811015610ca657600080610c148385610b2a565b5160208151910182875af13d15610c9e573d90610c3082610921565b91610c3e60405193846108b1565b82523d6000602084013e5b15610c575750600101610bfe565b90610c9a6040519283927f5c0dee5d000000000000000000000000000000000000000000000000000000008452600484015260406024840152604483019061095b565b0390fd5b606090610c49565b505090565b7f741752c20000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fe7557a52000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b8054821015610afb5760005260206000200190600090565b6000818152600360205260409020548015610f03577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610ed457600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610ed457818103610e65575b5050506002548015610e36577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610df3816002610d2d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610ebc610e76610e87936002610d2d565b90549060031b1c9283926002610d2d565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080610dba565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610f6457600254680100000000000000008110156108f257610f4b610e878260018594016002556002610d2d565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
}

var CREATE2FactoryABI = CREATE2FactoryMetaData.ABI

var CREATE2FactoryBin = CREATE2FactoryMetaData.Bin

func DeployCREATE2Factory(auth *bind.TransactOpts, backend bind.ContractBackend, allowList []common.Address) (common.Address, *types.Transaction, *CREATE2Factory, error) {
	parsed, err := CREATE2FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CREATE2FactoryBin), backend, allowList)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CREATE2Factory{address: address, abi: *parsed, CREATE2FactoryCaller: CREATE2FactoryCaller{contract: contract}, CREATE2FactoryTransactor: CREATE2FactoryTransactor{contract: contract}, CREATE2FactoryFilterer: CREATE2FactoryFilterer{contract: contract}}, nil
}

type CREATE2Factory struct {
	address common.Address
	abi     abi.ABI
	CREATE2FactoryCaller
	CREATE2FactoryTransactor
	CREATE2FactoryFilterer
}

type CREATE2FactoryCaller struct {
	contract *bind.BoundContract
}

type CREATE2FactoryTransactor struct {
	contract *bind.BoundContract
}

type CREATE2FactoryFilterer struct {
	contract *bind.BoundContract
}

type CREATE2FactorySession struct {
	Contract     *CREATE2Factory
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CREATE2FactoryCallerSession struct {
	Contract *CREATE2FactoryCaller
	CallOpts bind.CallOpts
}

type CREATE2FactoryTransactorSession struct {
	Contract     *CREATE2FactoryTransactor
	TransactOpts bind.TransactOpts
}

type CREATE2FactoryRaw struct {
	Contract *CREATE2Factory
}

type CREATE2FactoryCallerRaw struct {
	Contract *CREATE2FactoryCaller
}

type CREATE2FactoryTransactorRaw struct {
	Contract *CREATE2FactoryTransactor
}

func NewCREATE2Factory(address common.Address, backend bind.ContractBackend) (*CREATE2Factory, error) {
	abi, err := abi.JSON(strings.NewReader(CREATE2FactoryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCREATE2Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CREATE2Factory{address: address, abi: abi, CREATE2FactoryCaller: CREATE2FactoryCaller{contract: contract}, CREATE2FactoryTransactor: CREATE2FactoryTransactor{contract: contract}, CREATE2FactoryFilterer: CREATE2FactoryFilterer{contract: contract}}, nil
}

func NewCREATE2FactoryCaller(address common.Address, caller bind.ContractCaller) (*CREATE2FactoryCaller, error) {
	contract, err := bindCREATE2Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CREATE2FactoryCaller{contract: contract}, nil
}

func NewCREATE2FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CREATE2FactoryTransactor, error) {
	contract, err := bindCREATE2Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CREATE2FactoryTransactor{contract: contract}, nil
}

func NewCREATE2FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CREATE2FactoryFilterer, error) {
	contract, err := bindCREATE2Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CREATE2FactoryFilterer{contract: contract}, nil
}

func bindCREATE2Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CREATE2FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CREATE2Factory *CREATE2FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CREATE2Factory.Contract.CREATE2FactoryCaller.contract.Call(opts, result, method, params...)
}

func (_CREATE2Factory *CREATE2FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.CREATE2FactoryTransactor.contract.Transfer(opts)
}

func (_CREATE2Factory *CREATE2FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.CREATE2FactoryTransactor.contract.Transact(opts, method, params...)
}

func (_CREATE2Factory *CREATE2FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CREATE2Factory.Contract.contract.Call(opts, result, method, params...)
}

func (_CREATE2Factory *CREATE2FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.contract.Transfer(opts)
}

func (_CREATE2Factory *CREATE2FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.contract.Transact(opts, method, params...)
}

func (_CREATE2Factory *CREATE2FactoryCaller) ComputeAddress(opts *bind.CallOpts, creationCode []byte, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _CREATE2Factory.contract.Call(opts, &out, "computeAddress", creationCode, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CREATE2Factory *CREATE2FactorySession) ComputeAddress(creationCode []byte, salt [32]byte) (common.Address, error) {
	return _CREATE2Factory.Contract.ComputeAddress(&_CREATE2Factory.CallOpts, creationCode, salt)
}

func (_CREATE2Factory *CREATE2FactoryCallerSession) ComputeAddress(creationCode []byte, salt [32]byte) (common.Address, error) {
	return _CREATE2Factory.Contract.ComputeAddress(&_CREATE2Factory.CallOpts, creationCode, salt)
}

func (_CREATE2Factory *CREATE2FactoryCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CREATE2Factory.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CREATE2Factory *CREATE2FactorySession) GetAllowList() ([]common.Address, error) {
	return _CREATE2Factory.Contract.GetAllowList(&_CREATE2Factory.CallOpts)
}

func (_CREATE2Factory *CREATE2FactoryCallerSession) GetAllowList() ([]common.Address, error) {
	return _CREATE2Factory.Contract.GetAllowList(&_CREATE2Factory.CallOpts)
}

func (_CREATE2Factory *CREATE2FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CREATE2Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CREATE2Factory *CREATE2FactorySession) Owner() (common.Address, error) {
	return _CREATE2Factory.Contract.Owner(&_CREATE2Factory.CallOpts)
}

func (_CREATE2Factory *CREATE2FactoryCallerSession) Owner() (common.Address, error) {
	return _CREATE2Factory.Contract.Owner(&_CREATE2Factory.CallOpts)
}

func (_CREATE2Factory *CREATE2FactoryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CREATE2Factory.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CREATE2Factory *CREATE2FactorySession) TypeAndVersion() (string, error) {
	return _CREATE2Factory.Contract.TypeAndVersion(&_CREATE2Factory.CallOpts)
}

func (_CREATE2Factory *CREATE2FactoryCallerSession) TypeAndVersion() (string, error) {
	return _CREATE2Factory.Contract.TypeAndVersion(&_CREATE2Factory.CallOpts)
}

func (_CREATE2Factory *CREATE2FactoryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CREATE2Factory.contract.Transact(opts, "acceptOwnership")
}

func (_CREATE2Factory *CREATE2FactorySession) AcceptOwnership() (*types.Transaction, error) {
	return _CREATE2Factory.Contract.AcceptOwnership(&_CREATE2Factory.TransactOpts)
}

func (_CREATE2Factory *CREATE2FactoryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CREATE2Factory.Contract.AcceptOwnership(&_CREATE2Factory.TransactOpts)
}

func (_CREATE2Factory *CREATE2FactoryTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_CREATE2Factory *CREATE2FactorySession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.ApplyAllowListUpdates(&_CREATE2Factory.TransactOpts, removes, adds)
}

func (_CREATE2Factory *CREATE2FactoryTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.ApplyAllowListUpdates(&_CREATE2Factory.TransactOpts, removes, adds)
}

func (_CREATE2Factory *CREATE2FactoryTransactor) CreateAndCall(opts *bind.TransactOpts, creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _CREATE2Factory.contract.Transact(opts, "createAndCall", creationCode, salt, calls)
}

func (_CREATE2Factory *CREATE2FactorySession) CreateAndCall(creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.CreateAndCall(&_CREATE2Factory.TransactOpts, creationCode, salt, calls)
}

func (_CREATE2Factory *CREATE2FactoryTransactorSession) CreateAndCall(creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.CreateAndCall(&_CREATE2Factory.TransactOpts, creationCode, salt, calls)
}

func (_CREATE2Factory *CREATE2FactoryTransactor) CreateAndTransferOwnership(opts *bind.TransactOpts, creationCode []byte, salt [32]byte, to common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.contract.Transact(opts, "createAndTransferOwnership", creationCode, salt, to)
}

func (_CREATE2Factory *CREATE2FactorySession) CreateAndTransferOwnership(creationCode []byte, salt [32]byte, to common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.CreateAndTransferOwnership(&_CREATE2Factory.TransactOpts, creationCode, salt, to)
}

func (_CREATE2Factory *CREATE2FactoryTransactorSession) CreateAndTransferOwnership(creationCode []byte, salt [32]byte, to common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.CreateAndTransferOwnership(&_CREATE2Factory.TransactOpts, creationCode, salt, to)
}

func (_CREATE2Factory *CREATE2FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.contract.Transact(opts, "transferOwnership", to)
}

func (_CREATE2Factory *CREATE2FactorySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.TransferOwnership(&_CREATE2Factory.TransactOpts, to)
}

func (_CREATE2Factory *CREATE2FactoryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CREATE2Factory.Contract.TransferOwnership(&_CREATE2Factory.TransactOpts, to)
}

type CREATE2FactoryCallerAddedIterator struct {
	Event *CREATE2FactoryCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CREATE2FactoryCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CREATE2FactoryCallerAdded)
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
		it.Event = new(CREATE2FactoryCallerAdded)
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

func (it *CREATE2FactoryCallerAddedIterator) Error() error {
	return it.fail
}

func (it *CREATE2FactoryCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CREATE2FactoryCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_CREATE2Factory *CREATE2FactoryFilterer) FilterCallerAdded(opts *bind.FilterOpts, caller []common.Address) (*CREATE2FactoryCallerAddedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CREATE2Factory.contract.FilterLogs(opts, "CallerAdded", callerRule)
	if err != nil {
		return nil, err
	}
	return &CREATE2FactoryCallerAddedIterator{contract: _CREATE2Factory.contract, event: "CallerAdded", logs: logs, sub: sub}, nil
}

func (_CREATE2Factory *CREATE2FactoryFilterer) WatchCallerAdded(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryCallerAdded, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CREATE2Factory.contract.WatchLogs(opts, "CallerAdded", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CREATE2FactoryCallerAdded)
				if err := _CREATE2Factory.contract.UnpackLog(event, "CallerAdded", log); err != nil {
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

func (_CREATE2Factory *CREATE2FactoryFilterer) ParseCallerAdded(log types.Log) (*CREATE2FactoryCallerAdded, error) {
	event := new(CREATE2FactoryCallerAdded)
	if err := _CREATE2Factory.contract.UnpackLog(event, "CallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CREATE2FactoryCallerRemovedIterator struct {
	Event *CREATE2FactoryCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CREATE2FactoryCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CREATE2FactoryCallerRemoved)
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
		it.Event = new(CREATE2FactoryCallerRemoved)
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

func (it *CREATE2FactoryCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *CREATE2FactoryCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CREATE2FactoryCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_CREATE2Factory *CREATE2FactoryFilterer) FilterCallerRemoved(opts *bind.FilterOpts, caller []common.Address) (*CREATE2FactoryCallerRemovedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CREATE2Factory.contract.FilterLogs(opts, "CallerRemoved", callerRule)
	if err != nil {
		return nil, err
	}
	return &CREATE2FactoryCallerRemovedIterator{contract: _CREATE2Factory.contract, event: "CallerRemoved", logs: logs, sub: sub}, nil
}

func (_CREATE2Factory *CREATE2FactoryFilterer) WatchCallerRemoved(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryCallerRemoved, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CREATE2Factory.contract.WatchLogs(opts, "CallerRemoved", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CREATE2FactoryCallerRemoved)
				if err := _CREATE2Factory.contract.UnpackLog(event, "CallerRemoved", log); err != nil {
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

func (_CREATE2Factory *CREATE2FactoryFilterer) ParseCallerRemoved(log types.Log) (*CREATE2FactoryCallerRemoved, error) {
	event := new(CREATE2FactoryCallerRemoved)
	if err := _CREATE2Factory.contract.UnpackLog(event, "CallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CREATE2FactoryContractDeployedIterator struct {
	Event *CREATE2FactoryContractDeployed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CREATE2FactoryContractDeployedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CREATE2FactoryContractDeployed)
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
		it.Event = new(CREATE2FactoryContractDeployed)
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

func (it *CREATE2FactoryContractDeployedIterator) Error() error {
	return it.fail
}

func (it *CREATE2FactoryContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CREATE2FactoryContractDeployed struct {
	ContractAddress common.Address
	Raw             types.Log
}

func (_CREATE2Factory *CREATE2FactoryFilterer) FilterContractDeployed(opts *bind.FilterOpts, contractAddress []common.Address) (*CREATE2FactoryContractDeployedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _CREATE2Factory.contract.FilterLogs(opts, "ContractDeployed", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &CREATE2FactoryContractDeployedIterator{contract: _CREATE2Factory.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

func (_CREATE2Factory *CREATE2FactoryFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryContractDeployed, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _CREATE2Factory.contract.WatchLogs(opts, "ContractDeployed", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CREATE2FactoryContractDeployed)
				if err := _CREATE2Factory.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
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

func (_CREATE2Factory *CREATE2FactoryFilterer) ParseContractDeployed(log types.Log) (*CREATE2FactoryContractDeployed, error) {
	event := new(CREATE2FactoryContractDeployed)
	if err := _CREATE2Factory.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CREATE2FactoryOwnershipTransferRequestedIterator struct {
	Event *CREATE2FactoryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CREATE2FactoryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CREATE2FactoryOwnershipTransferRequested)
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
		it.Event = new(CREATE2FactoryOwnershipTransferRequested)
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

func (it *CREATE2FactoryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CREATE2FactoryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CREATE2FactoryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CREATE2Factory *CREATE2FactoryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CREATE2FactoryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CREATE2Factory.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CREATE2FactoryOwnershipTransferRequestedIterator{contract: _CREATE2Factory.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CREATE2Factory *CREATE2FactoryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CREATE2Factory.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CREATE2FactoryOwnershipTransferRequested)
				if err := _CREATE2Factory.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CREATE2Factory *CREATE2FactoryFilterer) ParseOwnershipTransferRequested(log types.Log) (*CREATE2FactoryOwnershipTransferRequested, error) {
	event := new(CREATE2FactoryOwnershipTransferRequested)
	if err := _CREATE2Factory.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CREATE2FactoryOwnershipTransferredIterator struct {
	Event *CREATE2FactoryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CREATE2FactoryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CREATE2FactoryOwnershipTransferred)
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
		it.Event = new(CREATE2FactoryOwnershipTransferred)
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

func (it *CREATE2FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CREATE2FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CREATE2FactoryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CREATE2Factory *CREATE2FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CREATE2FactoryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CREATE2Factory.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CREATE2FactoryOwnershipTransferredIterator{contract: _CREATE2Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CREATE2Factory *CREATE2FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CREATE2Factory.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CREATE2FactoryOwnershipTransferred)
				if err := _CREATE2Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CREATE2Factory *CREATE2FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*CREATE2FactoryOwnershipTransferred, error) {
	event := new(CREATE2FactoryOwnershipTransferred)
	if err := _CREATE2Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (CREATE2FactoryCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xf7762e85af7b409451f9a76004c5f755642902434eb11351ae67eb9746888b69")
}

func (CREATE2FactoryCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0x50c35a67b454d38c20800d5b55e320f58f4c9c86a28d8ab20f03045d1a38d99a")
}

func (CREATE2FactoryContractDeployed) Topic() common.Hash {
	return common.HexToHash("0x8ffcdc15a283d706d38281f500270d8b5a656918f555de0913d7455e3e6bc1bf")
}

func (CREATE2FactoryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CREATE2FactoryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CREATE2Factory *CREATE2Factory) Address() common.Address {
	return _CREATE2Factory.address
}

type CREATE2FactoryInterface interface {
	ComputeAddress(opts *bind.CallOpts, creationCode []byte, salt [32]byte) (common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	CreateAndCall(opts *bind.TransactOpts, creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error)

	CreateAndTransferOwnership(opts *bind.TransactOpts, creationCode []byte, salt [32]byte, to common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterCallerAdded(opts *bind.FilterOpts, caller []common.Address) (*CREATE2FactoryCallerAddedIterator, error)

	WatchCallerAdded(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryCallerAdded, caller []common.Address) (event.Subscription, error)

	ParseCallerAdded(log types.Log) (*CREATE2FactoryCallerAdded, error)

	FilterCallerRemoved(opts *bind.FilterOpts, caller []common.Address) (*CREATE2FactoryCallerRemovedIterator, error)

	WatchCallerRemoved(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryCallerRemoved, caller []common.Address) (event.Subscription, error)

	ParseCallerRemoved(log types.Log) (*CREATE2FactoryCallerRemoved, error)

	FilterContractDeployed(opts *bind.FilterOpts, contractAddress []common.Address) (*CREATE2FactoryContractDeployedIterator, error)

	WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryContractDeployed, contractAddress []common.Address) (event.Subscription, error)

	ParseContractDeployed(log types.Log) (*CREATE2FactoryContractDeployed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CREATE2FactoryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CREATE2FactoryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CREATE2FactoryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CREATE2FactoryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CREATE2FactoryOwnershipTransferred, error)

	Address() common.Address
}
