// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_factory

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

var ContractFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"allowList\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"computeAddress\",\"inputs\":[{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAndCall\",\"inputs\":[{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calls\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"CallerNotAllowed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x6080604052346101d75761115880380380610019816101f2565b9283398101906020818303126101d7578051906001600160401b0382116101d7570181601f820112156101d7578051916001600160401b0383116101dc578260051b9160208061006a8186016101f2565b8096815201938201019182116101d757602001915b8183106101b7578333156101a657600180546001600160a01b031916331790556100a960206101f2565b60008152600036813760005b8151811015610124576001906100dd6001600160a01b036100d68386610217565b5116610259565b6100e8575b016100b5565b818060a01b036100f88285610217565b51167f50c35a67b454d38c20800d5b55e320f58f4c9c86a28d8ab20f03045d1a38d99a600080a26100e2565b8260005b8151811015610197576001906101506001600160a01b036101498386610217565b5116610357565b61015b575b01610128565b818060a01b0361016b8285610217565b51167ff7762e85af7b409451f9a76004c5f755642902434eb11351ae67eb9746888b69600080a2610155565b604051610da090816103b88239f35b639b15e16f60e01b60005260046000fd5b82516001600160a01b03811681036101d75781526020928301920161007f565b600080fd5b634e487b7160e01b600052604160045260246000fd5b6040519190601f01601f191682016001600160401b038111838210176101dc57604052565b805182101561022b5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561022b5760005260206000200190600090565b600081815260036020526040902054801561035057600019810181811161033a5760025460001981019190821161033a578181036102e9575b50505060025480156102d357600019016102ad816002610241565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6103226102fa61030b936002610241565b90549060031b1c9283926002610241565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610292565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146103b157600254680100000000000000008110156101dc5761039861030b8260018594016002556002610241565b9055600254906000526003602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c8063181f5a77146108645780633e4b9f7a146107c157806354c8a4f31461061257806379ba5097146105295780638da5cb5b146104d7578063a7cd63b7146103e9578063e4a9848d1461016c5763f2fde38b1461007457600080fd5b346101675760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043573ffffffffffffffffffffffffffffffffffffffff8116809103610167576100cc610ac8565b33811461013d57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101675760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff811161016757366023820112156101675780600401359067ffffffffffffffff82116101675736602483830101116101675760443567ffffffffffffffff8111610167576101f9903690600401610a25565b929091336000526003602052604060002054156103bb5761021e9160243692016109ee565b805115610391578051602435916020016000f59073ffffffffffffffffffffffffffffffffffffffff821691821561036757600093927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe183360301945b8481101561035c5760008160051b8501358781121561035857850180359067ffffffffffffffff82116103545760200181360381136103545781839291839260405192839283378101838152039082885af13d1561034b573d6102dd81610955565b906102eb60405192836108e5565b8152809260203d92013e5b15610304575060010161027b565b906103476040519283927f5c0dee5d000000000000000000000000000000000000000000000000000000008452600484015260406024840152604483019061098f565b0390fd5b606091506102f6565b8280fd5b5080fd5b602082604051908152f35b7f741752c20000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fe7557a52000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b8181106104c157505050816104689103826108e5565b6040519182916020830190602084525180915260408301919060005b818110610492575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610484565b8254845260209093019260019283019201610452565b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760005473ffffffffffffffffffffffffffffffffffffffff811633036105e8577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101675760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff811161016757610661903690600401610a25565b6024359067ffffffffffffffff82116101675761069a6106886106a2933690600401610a25565b949092610693610ac8565b3691610a56565b923691610a56565b60005b825181101561073057806106da73ffffffffffffffffffffffffffffffffffffffff6106d360019487610b13565b5116610b6e565b6106e5575b016106a5565b73ffffffffffffffffffffffffffffffffffffffff6107048286610b13565b51167f50c35a67b454d38c20800d5b55e320f58f4c9c86a28d8ab20f03045d1a38d99a600080a26106df565b5060005b81518110156107bf578061076973ffffffffffffffffffffffffffffffffffffffff61076260019486610b13565b5116610d33565b610774575b01610734565b73ffffffffffffffffffffffffffffffffffffffff6107938285610b13565b51167ff7762e85af7b409451f9a76004c5f755642902434eb11351ae67eb9746888b69600080a261076e565b005b346101675760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101675760043567ffffffffffffffff81116101675736602382011215610167576055600b61082760209336906024816004013591016109ee565b838151910120604051906040820152602435848201523081520160ff81532073ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101675760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610167576108e160408051906108a581836108e5565b601582527f436f6e7472616374466163746f727920312e372e30000000000000000000000060208301525191829160208352602083019061098f565b0390f35b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761092657604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b67ffffffffffffffff811161092657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106109d95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161099a565b9291926109fa82610955565b91610a0860405193846108e5565b829481845281830111610167578281602093846000960137010152565b9181601f840112156101675782359167ffffffffffffffff8311610167576020808501948460051b01011161016757565b90929167ffffffffffffffff8411610926578360051b916020604051610a7e828601826108e5565b809681520192810191821161016757915b818310610a9b57505050565b823573ffffffffffffffffffffffffffffffffffffffff8116810361016757815260209283019201610a8f565b73ffffffffffffffffffffffffffffffffffffffff600154163303610ae957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015610b275760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8054821015610b275760005260206000200190600090565b6000818152600360205260409020548015610d2c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610cfd57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610cfd57818103610c8e575b5050506002548015610c5f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610c1c816002610b56565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610ce5610c9f610cb0936002610b56565b90549060031b1c9283926002610b56565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080610be3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610d8d576002546801000000000000000081101561092657610d74610cb08260018594016002556002610b56565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
}

var ContractFactoryABI = ContractFactoryMetaData.ABI

var ContractFactoryBin = ContractFactoryMetaData.Bin

func DeployContractFactory(auth *bind.TransactOpts, backend bind.ContractBackend, allowList []common.Address) (common.Address, *types.Transaction, *ContractFactory, error) {
	parsed, err := ContractFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractFactoryBin), backend, allowList)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractFactory{address: address, abi: *parsed, ContractFactoryCaller: ContractFactoryCaller{contract: contract}, ContractFactoryTransactor: ContractFactoryTransactor{contract: contract}, ContractFactoryFilterer: ContractFactoryFilterer{contract: contract}}, nil
}

type ContractFactory struct {
	address common.Address
	abi     abi.ABI
	ContractFactoryCaller
	ContractFactoryTransactor
	ContractFactoryFilterer
}

type ContractFactoryCaller struct {
	contract *bind.BoundContract
}

type ContractFactoryTransactor struct {
	contract *bind.BoundContract
}

type ContractFactoryFilterer struct {
	contract *bind.BoundContract
}

type ContractFactorySession struct {
	Contract     *ContractFactory
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ContractFactoryCallerSession struct {
	Contract *ContractFactoryCaller
	CallOpts bind.CallOpts
}

type ContractFactoryTransactorSession struct {
	Contract     *ContractFactoryTransactor
	TransactOpts bind.TransactOpts
}

type ContractFactoryRaw struct {
	Contract *ContractFactory
}

type ContractFactoryCallerRaw struct {
	Contract *ContractFactoryCaller
}

type ContractFactoryTransactorRaw struct {
	Contract *ContractFactoryTransactor
}

func NewContractFactory(address common.Address, backend bind.ContractBackend) (*ContractFactory, error) {
	abi, err := abi.JSON(strings.NewReader(ContractFactoryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindContractFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractFactory{address: address, abi: abi, ContractFactoryCaller: ContractFactoryCaller{contract: contract}, ContractFactoryTransactor: ContractFactoryTransactor{contract: contract}, ContractFactoryFilterer: ContractFactoryFilterer{contract: contract}}, nil
}

func NewContractFactoryCaller(address common.Address, caller bind.ContractCaller) (*ContractFactoryCaller, error) {
	contract, err := bindContractFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryCaller{contract: contract}, nil
}

func NewContractFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractFactoryTransactor, error) {
	contract, err := bindContractFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryTransactor{contract: contract}, nil
}

func NewContractFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFactoryFilterer, error) {
	contract, err := bindContractFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryFilterer{contract: contract}, nil
}

func bindContractFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ContractFactory *ContractFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractFactory.Contract.ContractFactoryCaller.contract.Call(opts, result, method, params...)
}

func (_ContractFactory *ContractFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractFactory.Contract.ContractFactoryTransactor.contract.Transfer(opts)
}

func (_ContractFactory *ContractFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractFactory.Contract.ContractFactoryTransactor.contract.Transact(opts, method, params...)
}

func (_ContractFactory *ContractFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractFactory.Contract.contract.Call(opts, result, method, params...)
}

func (_ContractFactory *ContractFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractFactory.Contract.contract.Transfer(opts)
}

func (_ContractFactory *ContractFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractFactory.Contract.contract.Transact(opts, method, params...)
}

func (_ContractFactory *ContractFactoryCaller) ComputeAddress(opts *bind.CallOpts, creationCode []byte, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractFactory.contract.Call(opts, &out, "computeAddress", creationCode, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ContractFactory *ContractFactorySession) ComputeAddress(creationCode []byte, salt [32]byte) (common.Address, error) {
	return _ContractFactory.Contract.ComputeAddress(&_ContractFactory.CallOpts, creationCode, salt)
}

func (_ContractFactory *ContractFactoryCallerSession) ComputeAddress(creationCode []byte, salt [32]byte) (common.Address, error) {
	return _ContractFactory.Contract.ComputeAddress(&_ContractFactory.CallOpts, creationCode, salt)
}

func (_ContractFactory *ContractFactoryCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractFactory.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_ContractFactory *ContractFactorySession) GetAllowList() ([]common.Address, error) {
	return _ContractFactory.Contract.GetAllowList(&_ContractFactory.CallOpts)
}

func (_ContractFactory *ContractFactoryCallerSession) GetAllowList() ([]common.Address, error) {
	return _ContractFactory.Contract.GetAllowList(&_ContractFactory.CallOpts)
}

func (_ContractFactory *ContractFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ContractFactory *ContractFactorySession) Owner() (common.Address, error) {
	return _ContractFactory.Contract.Owner(&_ContractFactory.CallOpts)
}

func (_ContractFactory *ContractFactoryCallerSession) Owner() (common.Address, error) {
	return _ContractFactory.Contract.Owner(&_ContractFactory.CallOpts)
}

func (_ContractFactory *ContractFactoryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractFactory.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ContractFactory *ContractFactorySession) TypeAndVersion() (string, error) {
	return _ContractFactory.Contract.TypeAndVersion(&_ContractFactory.CallOpts)
}

func (_ContractFactory *ContractFactoryCallerSession) TypeAndVersion() (string, error) {
	return _ContractFactory.Contract.TypeAndVersion(&_ContractFactory.CallOpts)
}

func (_ContractFactory *ContractFactoryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractFactory.contract.Transact(opts, "acceptOwnership")
}

func (_ContractFactory *ContractFactorySession) AcceptOwnership() (*types.Transaction, error) {
	return _ContractFactory.Contract.AcceptOwnership(&_ContractFactory.TransactOpts)
}

func (_ContractFactory *ContractFactoryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ContractFactory.Contract.AcceptOwnership(&_ContractFactory.TransactOpts)
}

func (_ContractFactory *ContractFactoryTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _ContractFactory.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_ContractFactory *ContractFactorySession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _ContractFactory.Contract.ApplyAllowListUpdates(&_ContractFactory.TransactOpts, removes, adds)
}

func (_ContractFactory *ContractFactoryTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _ContractFactory.Contract.ApplyAllowListUpdates(&_ContractFactory.TransactOpts, removes, adds)
}

func (_ContractFactory *ContractFactoryTransactor) CreateAndCall(opts *bind.TransactOpts, creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _ContractFactory.contract.Transact(opts, "createAndCall", creationCode, salt, calls)
}

func (_ContractFactory *ContractFactorySession) CreateAndCall(creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _ContractFactory.Contract.CreateAndCall(&_ContractFactory.TransactOpts, creationCode, salt, calls)
}

func (_ContractFactory *ContractFactoryTransactorSession) CreateAndCall(creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _ContractFactory.Contract.CreateAndCall(&_ContractFactory.TransactOpts, creationCode, salt, calls)
}

func (_ContractFactory *ContractFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ContractFactory.contract.Transact(opts, "transferOwnership", to)
}

func (_ContractFactory *ContractFactorySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ContractFactory.Contract.TransferOwnership(&_ContractFactory.TransactOpts, to)
}

func (_ContractFactory *ContractFactoryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ContractFactory.Contract.TransferOwnership(&_ContractFactory.TransactOpts, to)
}

type ContractFactoryCallerAddedIterator struct {
	Event *ContractFactoryCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractFactoryCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFactoryCallerAdded)
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
		it.Event = new(ContractFactoryCallerAdded)
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

func (it *ContractFactoryCallerAddedIterator) Error() error {
	return it.fail
}

func (it *ContractFactoryCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractFactoryCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_ContractFactory *ContractFactoryFilterer) FilterCallerAdded(opts *bind.FilterOpts, caller []common.Address) (*ContractFactoryCallerAddedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ContractFactory.contract.FilterLogs(opts, "CallerAdded", callerRule)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryCallerAddedIterator{contract: _ContractFactory.contract, event: "CallerAdded", logs: logs, sub: sub}, nil
}

func (_ContractFactory *ContractFactoryFilterer) WatchCallerAdded(opts *bind.WatchOpts, sink chan<- *ContractFactoryCallerAdded, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ContractFactory.contract.WatchLogs(opts, "CallerAdded", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractFactoryCallerAdded)
				if err := _ContractFactory.contract.UnpackLog(event, "CallerAdded", log); err != nil {
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

func (_ContractFactory *ContractFactoryFilterer) ParseCallerAdded(log types.Log) (*ContractFactoryCallerAdded, error) {
	event := new(ContractFactoryCallerAdded)
	if err := _ContractFactory.contract.UnpackLog(event, "CallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractFactoryCallerRemovedIterator struct {
	Event *ContractFactoryCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractFactoryCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFactoryCallerRemoved)
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
		it.Event = new(ContractFactoryCallerRemoved)
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

func (it *ContractFactoryCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *ContractFactoryCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractFactoryCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_ContractFactory *ContractFactoryFilterer) FilterCallerRemoved(opts *bind.FilterOpts, caller []common.Address) (*ContractFactoryCallerRemovedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ContractFactory.contract.FilterLogs(opts, "CallerRemoved", callerRule)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryCallerRemovedIterator{contract: _ContractFactory.contract, event: "CallerRemoved", logs: logs, sub: sub}, nil
}

func (_ContractFactory *ContractFactoryFilterer) WatchCallerRemoved(opts *bind.WatchOpts, sink chan<- *ContractFactoryCallerRemoved, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ContractFactory.contract.WatchLogs(opts, "CallerRemoved", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractFactoryCallerRemoved)
				if err := _ContractFactory.contract.UnpackLog(event, "CallerRemoved", log); err != nil {
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

func (_ContractFactory *ContractFactoryFilterer) ParseCallerRemoved(log types.Log) (*ContractFactoryCallerRemoved, error) {
	event := new(ContractFactoryCallerRemoved)
	if err := _ContractFactory.contract.UnpackLog(event, "CallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractFactoryOwnershipTransferRequestedIterator struct {
	Event *ContractFactoryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractFactoryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFactoryOwnershipTransferRequested)
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
		it.Event = new(ContractFactoryOwnershipTransferRequested)
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

func (it *ContractFactoryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ContractFactoryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractFactoryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ContractFactory *ContractFactoryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractFactoryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractFactory.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryOwnershipTransferRequestedIterator{contract: _ContractFactory.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ContractFactory *ContractFactoryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ContractFactoryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractFactory.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractFactoryOwnershipTransferRequested)
				if err := _ContractFactory.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ContractFactory *ContractFactoryFilterer) ParseOwnershipTransferRequested(log types.Log) (*ContractFactoryOwnershipTransferRequested, error) {
	event := new(ContractFactoryOwnershipTransferRequested)
	if err := _ContractFactory.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractFactoryOwnershipTransferredIterator struct {
	Event *ContractFactoryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractFactoryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFactoryOwnershipTransferred)
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
		it.Event = new(ContractFactoryOwnershipTransferred)
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

func (it *ContractFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ContractFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractFactoryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ContractFactory *ContractFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractFactoryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractFactory.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractFactoryOwnershipTransferredIterator{contract: _ContractFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ContractFactory *ContractFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractFactoryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractFactory.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractFactoryOwnershipTransferred)
				if err := _ContractFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ContractFactory *ContractFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*ContractFactoryOwnershipTransferred, error) {
	event := new(ContractFactoryOwnershipTransferred)
	if err := _ContractFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ContractFactoryCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xf7762e85af7b409451f9a76004c5f755642902434eb11351ae67eb9746888b69")
}

func (ContractFactoryCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0x50c35a67b454d38c20800d5b55e320f58f4c9c86a28d8ab20f03045d1a38d99a")
}

func (ContractFactoryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ContractFactoryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ContractFactory *ContractFactory) Address() common.Address {
	return _ContractFactory.address
}

type ContractFactoryInterface interface {
	ComputeAddress(opts *bind.CallOpts, creationCode []byte, salt [32]byte) (common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	CreateAndCall(opts *bind.TransactOpts, creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterCallerAdded(opts *bind.FilterOpts, caller []common.Address) (*ContractFactoryCallerAddedIterator, error)

	WatchCallerAdded(opts *bind.WatchOpts, sink chan<- *ContractFactoryCallerAdded, caller []common.Address) (event.Subscription, error)

	ParseCallerAdded(log types.Log) (*ContractFactoryCallerAdded, error)

	FilterCallerRemoved(opts *bind.FilterOpts, caller []common.Address) (*ContractFactoryCallerRemovedIterator, error)

	WatchCallerRemoved(opts *bind.WatchOpts, sink chan<- *ContractFactoryCallerRemoved, caller []common.Address) (event.Subscription, error)

	ParseCallerRemoved(log types.Log) (*ContractFactoryCallerRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractFactoryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ContractFactoryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ContractFactoryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractFactoryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractFactoryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ContractFactoryOwnershipTransferred, error)

	Address() common.Address
}
