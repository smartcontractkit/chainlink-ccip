// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package same_address_remote_token_address_provider

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

var SameAddressRemoteTokenAddressProviderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getRemoteTokenAddress\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"remoteToken\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"}]",
	Bin: "0x60808060405234601557610226908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816309bebd6214610122575063181f5a771461003557600080fd5b3461011d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261011d57604051606081019080821067ffffffffffffffff8311176100ee576100ea91604052602f81527f53616d654164647265737352656d6f7465546f6b656e4164647265737350726f60208201527f766964657220322e312e302d646576000000000000000000000000000000000060408201526040519182916020835260208301906101ba565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b3461011d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261011d5760043573ffffffffffffffffffffffffffffffffffffffff811680910361011d5760243567ffffffffffffffff81160361011d57602082015260208152604081019080821067ffffffffffffffff8311176100ee576040829052602080835282916100ea918301905b919082519283825260005b8481106102045750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016101c556fea164736f6c634300081a000a",
}

var SameAddressRemoteTokenAddressProviderABI = SameAddressRemoteTokenAddressProviderMetaData.ABI

var SameAddressRemoteTokenAddressProviderBin = SameAddressRemoteTokenAddressProviderMetaData.Bin

func DeploySameAddressRemoteTokenAddressProvider(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SameAddressRemoteTokenAddressProvider, error) {
	parsed, err := SameAddressRemoteTokenAddressProviderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SameAddressRemoteTokenAddressProviderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SameAddressRemoteTokenAddressProvider{address: address, abi: *parsed, SameAddressRemoteTokenAddressProviderCaller: SameAddressRemoteTokenAddressProviderCaller{contract: contract}, SameAddressRemoteTokenAddressProviderTransactor: SameAddressRemoteTokenAddressProviderTransactor{contract: contract}, SameAddressRemoteTokenAddressProviderFilterer: SameAddressRemoteTokenAddressProviderFilterer{contract: contract}}, nil
}

type SameAddressRemoteTokenAddressProvider struct {
	address common.Address
	abi     abi.ABI
	SameAddressRemoteTokenAddressProviderCaller
	SameAddressRemoteTokenAddressProviderTransactor
	SameAddressRemoteTokenAddressProviderFilterer
}

type SameAddressRemoteTokenAddressProviderCaller struct {
	contract *bind.BoundContract
}

type SameAddressRemoteTokenAddressProviderTransactor struct {
	contract *bind.BoundContract
}

type SameAddressRemoteTokenAddressProviderFilterer struct {
	contract *bind.BoundContract
}

type SameAddressRemoteTokenAddressProviderSession struct {
	Contract     *SameAddressRemoteTokenAddressProvider
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SameAddressRemoteTokenAddressProviderCallerSession struct {
	Contract *SameAddressRemoteTokenAddressProviderCaller
	CallOpts bind.CallOpts
}

type SameAddressRemoteTokenAddressProviderTransactorSession struct {
	Contract     *SameAddressRemoteTokenAddressProviderTransactor
	TransactOpts bind.TransactOpts
}

type SameAddressRemoteTokenAddressProviderRaw struct {
	Contract *SameAddressRemoteTokenAddressProvider
}

type SameAddressRemoteTokenAddressProviderCallerRaw struct {
	Contract *SameAddressRemoteTokenAddressProviderCaller
}

type SameAddressRemoteTokenAddressProviderTransactorRaw struct {
	Contract *SameAddressRemoteTokenAddressProviderTransactor
}

func NewSameAddressRemoteTokenAddressProvider(address common.Address, backend bind.ContractBackend) (*SameAddressRemoteTokenAddressProvider, error) {
	abi, err := abi.JSON(strings.NewReader(SameAddressRemoteTokenAddressProviderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSameAddressRemoteTokenAddressProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SameAddressRemoteTokenAddressProvider{address: address, abi: abi, SameAddressRemoteTokenAddressProviderCaller: SameAddressRemoteTokenAddressProviderCaller{contract: contract}, SameAddressRemoteTokenAddressProviderTransactor: SameAddressRemoteTokenAddressProviderTransactor{contract: contract}, SameAddressRemoteTokenAddressProviderFilterer: SameAddressRemoteTokenAddressProviderFilterer{contract: contract}}, nil
}

func NewSameAddressRemoteTokenAddressProviderCaller(address common.Address, caller bind.ContractCaller) (*SameAddressRemoteTokenAddressProviderCaller, error) {
	contract, err := bindSameAddressRemoteTokenAddressProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SameAddressRemoteTokenAddressProviderCaller{contract: contract}, nil
}

func NewSameAddressRemoteTokenAddressProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*SameAddressRemoteTokenAddressProviderTransactor, error) {
	contract, err := bindSameAddressRemoteTokenAddressProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SameAddressRemoteTokenAddressProviderTransactor{contract: contract}, nil
}

func NewSameAddressRemoteTokenAddressProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*SameAddressRemoteTokenAddressProviderFilterer, error) {
	contract, err := bindSameAddressRemoteTokenAddressProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SameAddressRemoteTokenAddressProviderFilterer{contract: contract}, nil
}

func bindSameAddressRemoteTokenAddressProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SameAddressRemoteTokenAddressProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SameAddressRemoteTokenAddressProvider.Contract.SameAddressRemoteTokenAddressProviderCaller.contract.Call(opts, result, method, params...)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SameAddressRemoteTokenAddressProvider.Contract.SameAddressRemoteTokenAddressProviderTransactor.contract.Transfer(opts)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SameAddressRemoteTokenAddressProvider.Contract.SameAddressRemoteTokenAddressProviderTransactor.contract.Transact(opts, method, params...)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SameAddressRemoteTokenAddressProvider.Contract.contract.Call(opts, result, method, params...)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SameAddressRemoteTokenAddressProvider.Contract.contract.Transfer(opts)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SameAddressRemoteTokenAddressProvider.Contract.contract.Transact(opts, method, params...)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderCaller) GetRemoteTokenAddress(opts *bind.CallOpts, localToken common.Address, arg1 uint64) ([]byte, error) {
	var out []interface{}
	err := _SameAddressRemoteTokenAddressProvider.contract.Call(opts, &out, "getRemoteTokenAddress", localToken, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderSession) GetRemoteTokenAddress(localToken common.Address, arg1 uint64) ([]byte, error) {
	return _SameAddressRemoteTokenAddressProvider.Contract.GetRemoteTokenAddress(&_SameAddressRemoteTokenAddressProvider.CallOpts, localToken, arg1)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderCallerSession) GetRemoteTokenAddress(localToken common.Address, arg1 uint64) ([]byte, error) {
	return _SameAddressRemoteTokenAddressProvider.Contract.GetRemoteTokenAddress(&_SameAddressRemoteTokenAddressProvider.CallOpts, localToken, arg1)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SameAddressRemoteTokenAddressProvider.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderSession) TypeAndVersion() (string, error) {
	return _SameAddressRemoteTokenAddressProvider.Contract.TypeAndVersion(&_SameAddressRemoteTokenAddressProvider.CallOpts)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProviderCallerSession) TypeAndVersion() (string, error) {
	return _SameAddressRemoteTokenAddressProvider.Contract.TypeAndVersion(&_SameAddressRemoteTokenAddressProvider.CallOpts)
}

func (_SameAddressRemoteTokenAddressProvider *SameAddressRemoteTokenAddressProvider) Address() common.Address {
	return _SameAddressRemoteTokenAddressProvider.address
}

type SameAddressRemoteTokenAddressProviderInterface interface {
	GetRemoteTokenAddress(opts *bind.CallOpts, localToken common.Address, arg1 uint64) ([]byte, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Address() common.Address
}
