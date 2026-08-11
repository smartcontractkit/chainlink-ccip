// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_token_minter

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

var MockTokenMinterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getLocalToken\",\"inputs\":[{\"name\":\"remoteDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setLocalToken\",\"inputs\":[{\"name\":\"remoteDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"remoteToken\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x60808060405234601557610177908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80636a879ac4146100b0576378a0565e1461003257600080fd5b346100ab5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ab5763ffffffff61006e610157565b1660005260006020526040600020602435600052602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b600080fd5b346100ab5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ab576100e7610157565b6044359073ffffffffffffffffffffffffffffffffffffffff82168092036100ab5763ffffffff16600052600060205260406000206024356000526020526040600020907fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055600080f35b6004359063ffffffff821682036100ab5756fea164736f6c634300081a000a",
}

var MockTokenMinterABI = MockTokenMinterMetaData.ABI

var MockTokenMinterBin = MockTokenMinterMetaData.Bin

func DeployMockTokenMinter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockTokenMinter, error) {
	parsed, err := MockTokenMinterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockTokenMinterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockTokenMinter{address: address, abi: *parsed, MockTokenMinterCaller: MockTokenMinterCaller{contract: contract}, MockTokenMinterTransactor: MockTokenMinterTransactor{contract: contract}, MockTokenMinterFilterer: MockTokenMinterFilterer{contract: contract}}, nil
}

type MockTokenMinter struct {
	address common.Address
	abi     abi.ABI
	MockTokenMinterCaller
	MockTokenMinterTransactor
	MockTokenMinterFilterer
}

type MockTokenMinterCaller struct {
	contract *bind.BoundContract
}

type MockTokenMinterTransactor struct {
	contract *bind.BoundContract
}

type MockTokenMinterFilterer struct {
	contract *bind.BoundContract
}

type MockTokenMinterSession struct {
	Contract     *MockTokenMinter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockTokenMinterCallerSession struct {
	Contract *MockTokenMinterCaller
	CallOpts bind.CallOpts
}

type MockTokenMinterTransactorSession struct {
	Contract     *MockTokenMinterTransactor
	TransactOpts bind.TransactOpts
}

type MockTokenMinterRaw struct {
	Contract *MockTokenMinter
}

type MockTokenMinterCallerRaw struct {
	Contract *MockTokenMinterCaller
}

type MockTokenMinterTransactorRaw struct {
	Contract *MockTokenMinterTransactor
}

func NewMockTokenMinter(address common.Address, backend bind.ContractBackend) (*MockTokenMinter, error) {
	abi, err := abi.JSON(strings.NewReader(MockTokenMinterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockTokenMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockTokenMinter{address: address, abi: abi, MockTokenMinterCaller: MockTokenMinterCaller{contract: contract}, MockTokenMinterTransactor: MockTokenMinterTransactor{contract: contract}, MockTokenMinterFilterer: MockTokenMinterFilterer{contract: contract}}, nil
}

func NewMockTokenMinterCaller(address common.Address, caller bind.ContractCaller) (*MockTokenMinterCaller, error) {
	contract, err := bindMockTokenMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockTokenMinterCaller{contract: contract}, nil
}

func NewMockTokenMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockTokenMinterTransactor, error) {
	contract, err := bindMockTokenMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTokenMinterTransactor{contract: contract}, nil
}

func NewMockTokenMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockTokenMinterFilterer, error) {
	contract, err := bindMockTokenMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockTokenMinterFilterer{contract: contract}, nil
}

func bindMockTokenMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockTokenMinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockTokenMinter *MockTokenMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTokenMinter.Contract.MockTokenMinterCaller.contract.Call(opts, result, method, params...)
}

func (_MockTokenMinter *MockTokenMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTokenMinter.Contract.MockTokenMinterTransactor.contract.Transfer(opts)
}

func (_MockTokenMinter *MockTokenMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTokenMinter.Contract.MockTokenMinterTransactor.contract.Transact(opts, method, params...)
}

func (_MockTokenMinter *MockTokenMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockTokenMinter.Contract.contract.Call(opts, result, method, params...)
}

func (_MockTokenMinter *MockTokenMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockTokenMinter.Contract.contract.Transfer(opts)
}

func (_MockTokenMinter *MockTokenMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockTokenMinter.Contract.contract.Transact(opts, method, params...)
}

func (_MockTokenMinter *MockTokenMinterCaller) GetLocalToken(opts *bind.CallOpts, remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MockTokenMinter.contract.Call(opts, &out, "getLocalToken", remoteDomain, remoteToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockTokenMinter *MockTokenMinterSession) GetLocalToken(remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	return _MockTokenMinter.Contract.GetLocalToken(&_MockTokenMinter.CallOpts, remoteDomain, remoteToken)
}

func (_MockTokenMinter *MockTokenMinterCallerSession) GetLocalToken(remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	return _MockTokenMinter.Contract.GetLocalToken(&_MockTokenMinter.CallOpts, remoteDomain, remoteToken)
}

func (_MockTokenMinter *MockTokenMinterTransactor) SetLocalToken(opts *bind.TransactOpts, remoteDomain uint32, remoteToken [32]byte, localToken common.Address) (*types.Transaction, error) {
	return _MockTokenMinter.contract.Transact(opts, "setLocalToken", remoteDomain, remoteToken, localToken)
}

func (_MockTokenMinter *MockTokenMinterSession) SetLocalToken(remoteDomain uint32, remoteToken [32]byte, localToken common.Address) (*types.Transaction, error) {
	return _MockTokenMinter.Contract.SetLocalToken(&_MockTokenMinter.TransactOpts, remoteDomain, remoteToken, localToken)
}

func (_MockTokenMinter *MockTokenMinterTransactorSession) SetLocalToken(remoteDomain uint32, remoteToken [32]byte, localToken common.Address) (*types.Transaction, error) {
	return _MockTokenMinter.Contract.SetLocalToken(&_MockTokenMinter.TransactOpts, remoteDomain, remoteToken, localToken)
}

func (_MockTokenMinter *MockTokenMinter) Address() common.Address {
	return _MockTokenMinter.address
}

type MockTokenMinterInterface interface {
	GetLocalToken(opts *bind.CallOpts, remoteDomain uint32, remoteToken [32]byte) (common.Address, error)

	SetLocalToken(opts *bind.TransactOpts, remoteDomain uint32, remoteToken [32]byte, localToken common.Address) (*types.Transaction, error)

	Address() common.Address
}
