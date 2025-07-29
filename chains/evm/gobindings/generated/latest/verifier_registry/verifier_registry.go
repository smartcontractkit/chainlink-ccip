// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_registry

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

var VerifierRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getVerifier\",\"inputs\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b0319163317905561042e90816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe608080604052600436101561001357600080fd5b60003560e01c908163181f5a77146103125750806379ba5097146102295780638da5cb5b146101d7578063eeb7b248146101775763f2fde38b1461005657600080fd5b346101725760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725760043573ffffffffffffffffffffffffffffffffffffffff81168091036101725773ffffffffffffffffffffffffffffffffffffffff600154168033036101485733821461011e57817fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000557fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101725760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610172576004356000526002602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346101725760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017257602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101725760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101725760005473ffffffffffffffffffffffffffffffffffffffff811633036102e8577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101725760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610172576040810181811067ffffffffffffffff8211176103f257604052601a81527f5665726966696572526567697374727920312e372e302d646576000000000000602082015260405190602082528181519182602083015260005b8381106103da5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b6020828201810151604087840101528593500161039a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300081a000a",
}

var VerifierRegistryABI = VerifierRegistryMetaData.ABI

var VerifierRegistryBin = VerifierRegistryMetaData.Bin

func DeployVerifierRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerifierRegistry, error) {
	parsed, err := VerifierRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierRegistry{address: address, abi: *parsed, VerifierRegistryCaller: VerifierRegistryCaller{contract: contract}, VerifierRegistryTransactor: VerifierRegistryTransactor{contract: contract}, VerifierRegistryFilterer: VerifierRegistryFilterer{contract: contract}}, nil
}

type VerifierRegistry struct {
	address common.Address
	abi     abi.ABI
	VerifierRegistryCaller
	VerifierRegistryTransactor
	VerifierRegistryFilterer
}

type VerifierRegistryCaller struct {
	contract *bind.BoundContract
}

type VerifierRegistryTransactor struct {
	contract *bind.BoundContract
}

type VerifierRegistryFilterer struct {
	contract *bind.BoundContract
}

type VerifierRegistrySession struct {
	Contract     *VerifierRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierRegistryCallerSession struct {
	Contract *VerifierRegistryCaller
	CallOpts bind.CallOpts
}

type VerifierRegistryTransactorSession struct {
	Contract     *VerifierRegistryTransactor
	TransactOpts bind.TransactOpts
}

type VerifierRegistryRaw struct {
	Contract *VerifierRegistry
}

type VerifierRegistryCallerRaw struct {
	Contract *VerifierRegistryCaller
}

type VerifierRegistryTransactorRaw struct {
	Contract *VerifierRegistryTransactor
}

func NewVerifierRegistry(address common.Address, backend bind.ContractBackend) (*VerifierRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierRegistry{address: address, abi: abi, VerifierRegistryCaller: VerifierRegistryCaller{contract: contract}, VerifierRegistryTransactor: VerifierRegistryTransactor{contract: contract}, VerifierRegistryFilterer: VerifierRegistryFilterer{contract: contract}}, nil
}

func NewVerifierRegistryCaller(address common.Address, caller bind.ContractCaller) (*VerifierRegistryCaller, error) {
	contract, err := bindVerifierRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierRegistryCaller{contract: contract}, nil
}

func NewVerifierRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierRegistryTransactor, error) {
	contract, err := bindVerifierRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierRegistryTransactor{contract: contract}, nil
}

func NewVerifierRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierRegistryFilterer, error) {
	contract, err := bindVerifierRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierRegistryFilterer{contract: contract}, nil
}

func bindVerifierRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierRegistry *VerifierRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierRegistry.Contract.VerifierRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierRegistry *VerifierRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierRegistry.Contract.VerifierRegistryTransactor.contract.Transfer(opts)
}

func (_VerifierRegistry *VerifierRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierRegistry.Contract.VerifierRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierRegistry *VerifierRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierRegistry *VerifierRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierRegistry.Contract.contract.Transfer(opts)
}

func (_VerifierRegistry *VerifierRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierRegistry *VerifierRegistryCaller) GetVerifier(opts *bind.CallOpts, verifierId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VerifierRegistry.contract.Call(opts, &out, "getVerifier", verifierId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierRegistry *VerifierRegistrySession) GetVerifier(verifierId [32]byte) (common.Address, error) {
	return _VerifierRegistry.Contract.GetVerifier(&_VerifierRegistry.CallOpts, verifierId)
}

func (_VerifierRegistry *VerifierRegistryCallerSession) GetVerifier(verifierId [32]byte) (common.Address, error) {
	return _VerifierRegistry.Contract.GetVerifier(&_VerifierRegistry.CallOpts, verifierId)
}

func (_VerifierRegistry *VerifierRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierRegistry *VerifierRegistrySession) Owner() (common.Address, error) {
	return _VerifierRegistry.Contract.Owner(&_VerifierRegistry.CallOpts)
}

func (_VerifierRegistry *VerifierRegistryCallerSession) Owner() (common.Address, error) {
	return _VerifierRegistry.Contract.Owner(&_VerifierRegistry.CallOpts)
}

func (_VerifierRegistry *VerifierRegistryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifierRegistry.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VerifierRegistry *VerifierRegistrySession) TypeAndVersion() (string, error) {
	return _VerifierRegistry.Contract.TypeAndVersion(&_VerifierRegistry.CallOpts)
}

func (_VerifierRegistry *VerifierRegistryCallerSession) TypeAndVersion() (string, error) {
	return _VerifierRegistry.Contract.TypeAndVersion(&_VerifierRegistry.CallOpts)
}

func (_VerifierRegistry *VerifierRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierRegistry *VerifierRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierRegistry.Contract.AcceptOwnership(&_VerifierRegistry.TransactOpts)
}

func (_VerifierRegistry *VerifierRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierRegistry.Contract.AcceptOwnership(&_VerifierRegistry.TransactOpts)
}

func (_VerifierRegistry *VerifierRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierRegistry *VerifierRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierRegistry.Contract.TransferOwnership(&_VerifierRegistry.TransactOpts, to)
}

func (_VerifierRegistry *VerifierRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierRegistry.Contract.TransferOwnership(&_VerifierRegistry.TransactOpts, to)
}

type VerifierRegistryOwnershipTransferRequestedIterator struct {
	Event *VerifierRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierRegistryOwnershipTransferRequested)
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
		it.Event = new(VerifierRegistryOwnershipTransferRequested)
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

func (it *VerifierRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierRegistry *VerifierRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierRegistryOwnershipTransferRequestedIterator{contract: _VerifierRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierRegistry *VerifierRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierRegistryOwnershipTransferRequested)
				if err := _VerifierRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VerifierRegistry *VerifierRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierRegistryOwnershipTransferRequested, error) {
	event := new(VerifierRegistryOwnershipTransferRequested)
	if err := _VerifierRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierRegistryOwnershipTransferredIterator struct {
	Event *VerifierRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierRegistryOwnershipTransferred)
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
		it.Event = new(VerifierRegistryOwnershipTransferred)
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

func (it *VerifierRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierRegistry *VerifierRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierRegistryOwnershipTransferredIterator{contract: _VerifierRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierRegistry *VerifierRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierRegistryOwnershipTransferred)
				if err := _VerifierRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VerifierRegistry *VerifierRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierRegistryOwnershipTransferred, error) {
	event := new(VerifierRegistryOwnershipTransferred)
	if err := _VerifierRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VerifierRegistry *VerifierRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _VerifierRegistry.ParseOwnershipTransferRequested(log)
	case _VerifierRegistry.abi.Events["OwnershipTransferred"].ID:
		return _VerifierRegistry.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VerifierRegistry *VerifierRegistry) Address() common.Address {
	return _VerifierRegistry.address
}

type VerifierRegistryInterface interface {
	GetVerifier(opts *bind.CallOpts, verifierId [32]byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierRegistryOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
