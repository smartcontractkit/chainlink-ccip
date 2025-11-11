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
	ABI: "[{\"type\":\"function\",\"name\":\"computeAddress\",\"inputs\":[{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAndCall\",\"inputs\":[{\"name\":\"creationCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calls\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60808060405234601557610502908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a771461033f5780633e4b9f7a1461029c5763e4a9848d1461003d57600080fd5b346102975760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102975760043567ffffffffffffffff8111610297573660238201121561029757806004013567ffffffffffffffff8111610297573660248284010111610297576044359067ffffffffffffffff821161029757366023830112156102975781600401359267ffffffffffffffff8411610297573660248560051b85010111610297576100f99160243692016104c9565b80511561026d578051602435916020016000f59073ffffffffffffffffffffffffffffffffffffffff821691821561024357600093927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbd83360301945b8481101561023857600060248260051b860101358781121561023457850160248101359067ffffffffffffffff82116102305760440181360381136102305781839291839260405192839283378101838152039082885af13d15610227573d6101c66101c182610430565b6103bd565b908152809260203d92013e5b156101e05750600101610156565b906102236040519283927f5c0dee5d000000000000000000000000000000000000000000000000000000008452600484015260406024840152604483019061046a565b0390fd5b606091506101d2565b8280fd5b5080fd5b602082604051908152f35b7f741752c20000000000000000000000000000000000000000000000000000000060005260046000fd5b7f4ca249dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102975760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102975760043567ffffffffffffffff81116102975736602382011215610297576055600b61030260209336906024816004013591016104c9565b838151910120604051906040820152602435848201523081520160ff81532073ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346102975760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610297576103b961037b60406103bd565b601581527f436f6e7472616374466163746f727920312e372e300000000000000000000000602082015260405191829160208352602083019061046a565b0390f35b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f604051930116820182811067ffffffffffffffff82111761040157604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b67ffffffffffffffff811161040157601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106104b45750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610475565b9291926104d86101c183610430565b93828552828201116102975781600092602092838701378401015256fea164736f6c634300081a000a",
}

var ContractFactoryABI = ContractFactoryMetaData.ABI

var ContractFactoryBin = ContractFactoryMetaData.Bin

func DeployContractFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractFactory, error) {
	parsed, err := ContractFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractFactoryBin), backend)
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

func (_ContractFactory *ContractFactoryTransactor) CreateAndCall(opts *bind.TransactOpts, creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _ContractFactory.contract.Transact(opts, "createAndCall", creationCode, salt, calls)
}

func (_ContractFactory *ContractFactorySession) CreateAndCall(creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _ContractFactory.Contract.CreateAndCall(&_ContractFactory.TransactOpts, creationCode, salt, calls)
}

func (_ContractFactory *ContractFactoryTransactorSession) CreateAndCall(creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error) {
	return _ContractFactory.Contract.CreateAndCall(&_ContractFactory.TransactOpts, creationCode, salt, calls)
}

func (_ContractFactory *ContractFactory) Address() common.Address {
	return _ContractFactory.address
}

type ContractFactoryInterface interface {
	ComputeAddress(opts *bind.CallOpts, creationCode []byte, salt [32]byte) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	CreateAndCall(opts *bind.TransactOpts, creationCode []byte, salt [32]byte, calls [][]byte) (*types.Transaction, error)

	Address() common.Address
}
