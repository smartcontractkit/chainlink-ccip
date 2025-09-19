// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ramp_proxy

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

var RampProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"rampAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_ramp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRamp\",\"inputs\":[{\"name\":\"rampAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RampUpdated\",\"inputs\":[{\"name\":\"oldRamp\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRamp\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6080346100e457601f61055b38819003918201601f19168301916001600160401b038311848410176100e9578084926020946040528339810103126100e457516001600160a01b038116908190036100e45733156100d357600180546001600160a01b0319163317905580156100c257600280546001600160a01b03198116831790915560405191906001600160a01b03167f9a6b1be05a589b6fe11a413ab4892655cec2680cfa59a0eae85fbf90ed987969600080a361045b90816101008239f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001a575b3415610411575b600080fd5b60003560e01c80633d9c3c171461006a578063654d74db1461006557806379ba5097146100605780638da5cb5b1461005b5763f2fde38b0361000e57610349565b6102f7565b61020e565b6101bc565b34610015576100783661016f565b73ffffffffffffffffffffffffffffffffffffffff6001541633036101455773ffffffffffffffffffffffffffffffffffffffff16801561011b5773ffffffffffffffffffffffffffffffffffffffff600254827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600255167f9a6b1be05a589b6fe11a413ab4892655cec2680cfa59a0eae85fbf90ed987969600080a3005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60209101126100155760043573ffffffffffffffffffffffffffffffffffffffff811681036100155790565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760005473ffffffffffffffffffffffffffffffffffffffff811633036102cd577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b34610015576103573661016f565b73ffffffffffffffffffffffffffffffffffffffff60015416908133036101455773ffffffffffffffffffffffffffffffffffffffff16903382146103e757817fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000557fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b60008073ffffffffffffffffffffffffffffffffffffffff600254163682803733600452818036925af13d6000803e610449573d6000fd5b3d6000f3fea164736f6c634300081a000a",
}

var RampProxyABI = RampProxyMetaData.ABI

var RampProxyBin = RampProxyMetaData.Bin

func DeployRampProxy(auth *bind.TransactOpts, backend bind.ContractBackend, rampAddress common.Address) (common.Address, *types.Transaction, *RampProxy, error) {
	parsed, err := RampProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RampProxyBin), backend, rampAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RampProxy{address: address, abi: *parsed, RampProxyCaller: RampProxyCaller{contract: contract}, RampProxyTransactor: RampProxyTransactor{contract: contract}, RampProxyFilterer: RampProxyFilterer{contract: contract}}, nil
}

type RampProxy struct {
	address common.Address
	abi     abi.ABI
	RampProxyCaller
	RampProxyTransactor
	RampProxyFilterer
}

type RampProxyCaller struct {
	contract *bind.BoundContract
}

type RampProxyTransactor struct {
	contract *bind.BoundContract
}

type RampProxyFilterer struct {
	contract *bind.BoundContract
}

type RampProxySession struct {
	Contract     *RampProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RampProxyCallerSession struct {
	Contract *RampProxyCaller
	CallOpts bind.CallOpts
}

type RampProxyTransactorSession struct {
	Contract     *RampProxyTransactor
	TransactOpts bind.TransactOpts
}

type RampProxyRaw struct {
	Contract *RampProxy
}

type RampProxyCallerRaw struct {
	Contract *RampProxyCaller
}

type RampProxyTransactorRaw struct {
	Contract *RampProxyTransactor
}

func NewRampProxy(address common.Address, backend bind.ContractBackend) (*RampProxy, error) {
	abi, err := abi.JSON(strings.NewReader(RampProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRampProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RampProxy{address: address, abi: abi, RampProxyCaller: RampProxyCaller{contract: contract}, RampProxyTransactor: RampProxyTransactor{contract: contract}, RampProxyFilterer: RampProxyFilterer{contract: contract}}, nil
}

func NewRampProxyCaller(address common.Address, caller bind.ContractCaller) (*RampProxyCaller, error) {
	contract, err := bindRampProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RampProxyCaller{contract: contract}, nil
}

func NewRampProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*RampProxyTransactor, error) {
	contract, err := bindRampProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RampProxyTransactor{contract: contract}, nil
}

func NewRampProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*RampProxyFilterer, error) {
	contract, err := bindRampProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RampProxyFilterer{contract: contract}, nil
}

func bindRampProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RampProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_RampProxy *RampProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RampProxy.Contract.RampProxyCaller.contract.Call(opts, result, method, params...)
}

func (_RampProxy *RampProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RampProxy.Contract.RampProxyTransactor.contract.Transfer(opts)
}

func (_RampProxy *RampProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RampProxy.Contract.RampProxyTransactor.contract.Transact(opts, method, params...)
}

func (_RampProxy *RampProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RampProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_RampProxy *RampProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RampProxy.Contract.contract.Transfer(opts)
}

func (_RampProxy *RampProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RampProxy.Contract.contract.Transact(opts, method, params...)
}

func (_RampProxy *RampProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RampProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RampProxy *RampProxySession) Owner() (common.Address, error) {
	return _RampProxy.Contract.Owner(&_RampProxy.CallOpts)
}

func (_RampProxy *RampProxyCallerSession) Owner() (common.Address, error) {
	return _RampProxy.Contract.Owner(&_RampProxy.CallOpts)
}

func (_RampProxy *RampProxyCaller) SRamp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RampProxy.contract.Call(opts, &out, "s_ramp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RampProxy *RampProxySession) SRamp() (common.Address, error) {
	return _RampProxy.Contract.SRamp(&_RampProxy.CallOpts)
}

func (_RampProxy *RampProxyCallerSession) SRamp() (common.Address, error) {
	return _RampProxy.Contract.SRamp(&_RampProxy.CallOpts)
}

func (_RampProxy *RampProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RampProxy.contract.Transact(opts, "acceptOwnership")
}

func (_RampProxy *RampProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _RampProxy.Contract.AcceptOwnership(&_RampProxy.TransactOpts)
}

func (_RampProxy *RampProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RampProxy.Contract.AcceptOwnership(&_RampProxy.TransactOpts)
}

func (_RampProxy *RampProxyTransactor) SetRamp(opts *bind.TransactOpts, rampAddress common.Address) (*types.Transaction, error) {
	return _RampProxy.contract.Transact(opts, "setRamp", rampAddress)
}

func (_RampProxy *RampProxySession) SetRamp(rampAddress common.Address) (*types.Transaction, error) {
	return _RampProxy.Contract.SetRamp(&_RampProxy.TransactOpts, rampAddress)
}

func (_RampProxy *RampProxyTransactorSession) SetRamp(rampAddress common.Address) (*types.Transaction, error) {
	return _RampProxy.Contract.SetRamp(&_RampProxy.TransactOpts, rampAddress)
}

func (_RampProxy *RampProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RampProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_RampProxy *RampProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RampProxy.Contract.TransferOwnership(&_RampProxy.TransactOpts, to)
}

func (_RampProxy *RampProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RampProxy.Contract.TransferOwnership(&_RampProxy.TransactOpts, to)
}

func (_RampProxy *RampProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RampProxy.contract.RawTransact(opts, calldata)
}

func (_RampProxy *RampProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RampProxy.Contract.Fallback(&_RampProxy.TransactOpts, calldata)
}

func (_RampProxy *RampProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RampProxy.Contract.Fallback(&_RampProxy.TransactOpts, calldata)
}

type RampProxyOwnershipTransferRequestedIterator struct {
	Event *RampProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RampProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RampProxyOwnershipTransferRequested)
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
		it.Event = new(RampProxyOwnershipTransferRequested)
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

func (it *RampProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RampProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RampProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RampProxy *RampProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RampProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RampProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RampProxyOwnershipTransferRequestedIterator{contract: _RampProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_RampProxy *RampProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RampProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RampProxyOwnershipTransferRequested)
				if err := _RampProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_RampProxy *RampProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*RampProxyOwnershipTransferRequested, error) {
	event := new(RampProxyOwnershipTransferRequested)
	if err := _RampProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RampProxyOwnershipTransferredIterator struct {
	Event *RampProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RampProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RampProxyOwnershipTransferred)
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
		it.Event = new(RampProxyOwnershipTransferred)
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

func (it *RampProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RampProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RampProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RampProxy *RampProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RampProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RampProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RampProxyOwnershipTransferredIterator{contract: _RampProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_RampProxy *RampProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RampProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RampProxyOwnershipTransferred)
				if err := _RampProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_RampProxy *RampProxyFilterer) ParseOwnershipTransferred(log types.Log) (*RampProxyOwnershipTransferred, error) {
	event := new(RampProxyOwnershipTransferred)
	if err := _RampProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RampProxyRampUpdatedIterator struct {
	Event *RampProxyRampUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RampProxyRampUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RampProxyRampUpdated)
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
		it.Event = new(RampProxyRampUpdated)
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

func (it *RampProxyRampUpdatedIterator) Error() error {
	return it.fail
}

func (it *RampProxyRampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RampProxyRampUpdated struct {
	OldRamp common.Address
	NewRamp common.Address
	Raw     types.Log
}

func (_RampProxy *RampProxyFilterer) FilterRampUpdated(opts *bind.FilterOpts, oldRamp []common.Address, newRamp []common.Address) (*RampProxyRampUpdatedIterator, error) {

	var oldRampRule []interface{}
	for _, oldRampItem := range oldRamp {
		oldRampRule = append(oldRampRule, oldRampItem)
	}
	var newRampRule []interface{}
	for _, newRampItem := range newRamp {
		newRampRule = append(newRampRule, newRampItem)
	}

	logs, sub, err := _RampProxy.contract.FilterLogs(opts, "RampUpdated", oldRampRule, newRampRule)
	if err != nil {
		return nil, err
	}
	return &RampProxyRampUpdatedIterator{contract: _RampProxy.contract, event: "RampUpdated", logs: logs, sub: sub}, nil
}

func (_RampProxy *RampProxyFilterer) WatchRampUpdated(opts *bind.WatchOpts, sink chan<- *RampProxyRampUpdated, oldRamp []common.Address, newRamp []common.Address) (event.Subscription, error) {

	var oldRampRule []interface{}
	for _, oldRampItem := range oldRamp {
		oldRampRule = append(oldRampRule, oldRampItem)
	}
	var newRampRule []interface{}
	for _, newRampItem := range newRamp {
		newRampRule = append(newRampRule, newRampItem)
	}

	logs, sub, err := _RampProxy.contract.WatchLogs(opts, "RampUpdated", oldRampRule, newRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RampProxyRampUpdated)
				if err := _RampProxy.contract.UnpackLog(event, "RampUpdated", log); err != nil {
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

func (_RampProxy *RampProxyFilterer) ParseRampUpdated(log types.Log) (*RampProxyRampUpdated, error) {
	event := new(RampProxyRampUpdated)
	if err := _RampProxy.contract.UnpackLog(event, "RampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (RampProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (RampProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (RampProxyRampUpdated) Topic() common.Hash {
	return common.HexToHash("0x9a6b1be05a589b6fe11a413ab4892655cec2680cfa59a0eae85fbf90ed987969")
}

func (_RampProxy *RampProxy) Address() common.Address {
	return _RampProxy.address
}

type RampProxyInterface interface {
	Owner(opts *bind.CallOpts) (common.Address, error)

	SRamp(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetRamp(opts *bind.TransactOpts, rampAddress common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RampProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*RampProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RampProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*RampProxyOwnershipTransferred, error)

	FilterRampUpdated(opts *bind.FilterOpts, oldRamp []common.Address, newRamp []common.Address) (*RampProxyRampUpdatedIterator, error)

	WatchRampUpdated(opts *bind.WatchOpts, sink chan<- *RampProxyRampUpdated, oldRamp []common.Address, newRamp []common.Address) (event.Subscription, error)

	ParseRampUpdated(log types.Log) (*RampProxyRampUpdated, error)

	Address() common.Address
}
