// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownable_ramp_proxy

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

var OwnableRampProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"rampAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_ramp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRamp\",\"inputs\":[{\"name\":\"rampAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RampUpdated\",\"inputs\":[{\"name\":\"oldRamp\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRamp\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6080346100e557601f61056438819003918201601f19168301916001600160401b038311848410176100ea578084926020946040528339810103126100e557516001600160a01b038116908190036100e55780156100d457600080546001600160a01b031981168317825560405192916001600160a01b03909116907f9a6b1be05a589b6fe11a413ab4892655cec2680cfa59a0eae85fbf90ed9879699080a333156100c357600280546001600160a01b0319163317905561046390816101018239f35b639b15e16f60e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001a575b34156103cf575b600080fd5b60003560e01c80633d9c3c171461006a578063654d74db1461006557806379ba5097146100605780638da5cb5b1461005b5763f2fde38b0361000e57610308565b6102b6565b6101cd565b61017b565b346100155773ffffffffffffffffffffffffffffffffffffffff61008d3661012e565b61009561040b565b1680156101045773ffffffffffffffffffffffffffffffffffffffff600054827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600055167f9a6b1be05a589b6fe11a413ab4892655cec2680cfa59a0eae85fbf90ed987969600080a3005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60209101126100155760043573ffffffffffffffffffffffffffffffffffffffff811681036100155790565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60005416604051908152f35b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760015473ffffffffffffffffffffffffffffffffffffffff8116330361028c577fffffffffffffffffffffffff00000000000000000000000000000000000000006002549133828416176002551660015573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b346100155773ffffffffffffffffffffffffffffffffffffffff61032b3661012e565b61033361040b565b163381146103a557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff600254167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b60008073ffffffffffffffffffffffffffffffffffffffff8154163682803733600452818036925af13d6000803e610406573d6000fd5b3d6000f35b73ffffffffffffffffffffffffffffffffffffffff60025416330361042c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fdfea164736f6c634300081a000a",
}

var OwnableRampProxyABI = OwnableRampProxyMetaData.ABI

var OwnableRampProxyBin = OwnableRampProxyMetaData.Bin

func DeployOwnableRampProxy(auth *bind.TransactOpts, backend bind.ContractBackend, rampAddress common.Address) (common.Address, *types.Transaction, *OwnableRampProxy, error) {
	parsed, err := OwnableRampProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnableRampProxyBin), backend, rampAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OwnableRampProxy{address: address, abi: *parsed, OwnableRampProxyCaller: OwnableRampProxyCaller{contract: contract}, OwnableRampProxyTransactor: OwnableRampProxyTransactor{contract: contract}, OwnableRampProxyFilterer: OwnableRampProxyFilterer{contract: contract}}, nil
}

type OwnableRampProxy struct {
	address common.Address
	abi     abi.ABI
	OwnableRampProxyCaller
	OwnableRampProxyTransactor
	OwnableRampProxyFilterer
}

type OwnableRampProxyCaller struct {
	contract *bind.BoundContract
}

type OwnableRampProxyTransactor struct {
	contract *bind.BoundContract
}

type OwnableRampProxyFilterer struct {
	contract *bind.BoundContract
}

type OwnableRampProxySession struct {
	Contract     *OwnableRampProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OwnableRampProxyCallerSession struct {
	Contract *OwnableRampProxyCaller
	CallOpts bind.CallOpts
}

type OwnableRampProxyTransactorSession struct {
	Contract     *OwnableRampProxyTransactor
	TransactOpts bind.TransactOpts
}

type OwnableRampProxyRaw struct {
	Contract *OwnableRampProxy
}

type OwnableRampProxyCallerRaw struct {
	Contract *OwnableRampProxyCaller
}

type OwnableRampProxyTransactorRaw struct {
	Contract *OwnableRampProxyTransactor
}

func NewOwnableRampProxy(address common.Address, backend bind.ContractBackend) (*OwnableRampProxy, error) {
	abi, err := abi.JSON(strings.NewReader(OwnableRampProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOwnableRampProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableRampProxy{address: address, abi: abi, OwnableRampProxyCaller: OwnableRampProxyCaller{contract: contract}, OwnableRampProxyTransactor: OwnableRampProxyTransactor{contract: contract}, OwnableRampProxyFilterer: OwnableRampProxyFilterer{contract: contract}}, nil
}

func NewOwnableRampProxyCaller(address common.Address, caller bind.ContractCaller) (*OwnableRampProxyCaller, error) {
	contract, err := bindOwnableRampProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableRampProxyCaller{contract: contract}, nil
}

func NewOwnableRampProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableRampProxyTransactor, error) {
	contract, err := bindOwnableRampProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableRampProxyTransactor{contract: contract}, nil
}

func NewOwnableRampProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableRampProxyFilterer, error) {
	contract, err := bindOwnableRampProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableRampProxyFilterer{contract: contract}, nil
}

func bindOwnableRampProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableRampProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OwnableRampProxy *OwnableRampProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableRampProxy.Contract.OwnableRampProxyCaller.contract.Call(opts, result, method, params...)
}

func (_OwnableRampProxy *OwnableRampProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.OwnableRampProxyTransactor.contract.Transfer(opts)
}

func (_OwnableRampProxy *OwnableRampProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.OwnableRampProxyTransactor.contract.Transact(opts, method, params...)
}

func (_OwnableRampProxy *OwnableRampProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableRampProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_OwnableRampProxy *OwnableRampProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.contract.Transfer(opts)
}

func (_OwnableRampProxy *OwnableRampProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.contract.Transact(opts, method, params...)
}

func (_OwnableRampProxy *OwnableRampProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableRampProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OwnableRampProxy *OwnableRampProxySession) Owner() (common.Address, error) {
	return _OwnableRampProxy.Contract.Owner(&_OwnableRampProxy.CallOpts)
}

func (_OwnableRampProxy *OwnableRampProxyCallerSession) Owner() (common.Address, error) {
	return _OwnableRampProxy.Contract.Owner(&_OwnableRampProxy.CallOpts)
}

func (_OwnableRampProxy *OwnableRampProxyCaller) SRamp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableRampProxy.contract.Call(opts, &out, "s_ramp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OwnableRampProxy *OwnableRampProxySession) SRamp() (common.Address, error) {
	return _OwnableRampProxy.Contract.SRamp(&_OwnableRampProxy.CallOpts)
}

func (_OwnableRampProxy *OwnableRampProxyCallerSession) SRamp() (common.Address, error) {
	return _OwnableRampProxy.Contract.SRamp(&_OwnableRampProxy.CallOpts)
}

func (_OwnableRampProxy *OwnableRampProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableRampProxy.contract.Transact(opts, "acceptOwnership")
}

func (_OwnableRampProxy *OwnableRampProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.AcceptOwnership(&_OwnableRampProxy.TransactOpts)
}

func (_OwnableRampProxy *OwnableRampProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.AcceptOwnership(&_OwnableRampProxy.TransactOpts)
}

func (_OwnableRampProxy *OwnableRampProxyTransactor) SetRamp(opts *bind.TransactOpts, rampAddress common.Address) (*types.Transaction, error) {
	return _OwnableRampProxy.contract.Transact(opts, "setRamp", rampAddress)
}

func (_OwnableRampProxy *OwnableRampProxySession) SetRamp(rampAddress common.Address) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.SetRamp(&_OwnableRampProxy.TransactOpts, rampAddress)
}

func (_OwnableRampProxy *OwnableRampProxyTransactorSession) SetRamp(rampAddress common.Address) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.SetRamp(&_OwnableRampProxy.TransactOpts, rampAddress)
}

func (_OwnableRampProxy *OwnableRampProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OwnableRampProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_OwnableRampProxy *OwnableRampProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.TransferOwnership(&_OwnableRampProxy.TransactOpts, to)
}

func (_OwnableRampProxy *OwnableRampProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.TransferOwnership(&_OwnableRampProxy.TransactOpts, to)
}

func (_OwnableRampProxy *OwnableRampProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _OwnableRampProxy.contract.RawTransact(opts, calldata)
}

func (_OwnableRampProxy *OwnableRampProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.Fallback(&_OwnableRampProxy.TransactOpts, calldata)
}

func (_OwnableRampProxy *OwnableRampProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.Fallback(&_OwnableRampProxy.TransactOpts, calldata)
}

type OwnableRampProxyOwnershipTransferRequestedIterator struct {
	Event *OwnableRampProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnableRampProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableRampProxyOwnershipTransferRequested)
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
		it.Event = new(OwnableRampProxyOwnershipTransferRequested)
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

func (it *OwnableRampProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OwnableRampProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnableRampProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OwnableRampProxy *OwnableRampProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableRampProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnableRampProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OwnableRampProxyOwnershipTransferRequestedIterator{contract: _OwnableRampProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OwnableRampProxy *OwnableRampProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnableRampProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnableRampProxyOwnershipTransferRequested)
				if err := _OwnableRampProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OwnableRampProxy *OwnableRampProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*OwnableRampProxyOwnershipTransferRequested, error) {
	event := new(OwnableRampProxyOwnershipTransferRequested)
	if err := _OwnableRampProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OwnableRampProxyOwnershipTransferredIterator struct {
	Event *OwnableRampProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnableRampProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableRampProxyOwnershipTransferred)
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
		it.Event = new(OwnableRampProxyOwnershipTransferred)
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

func (it *OwnableRampProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OwnableRampProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnableRampProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OwnableRampProxy *OwnableRampProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableRampProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnableRampProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OwnableRampProxyOwnershipTransferredIterator{contract: _OwnableRampProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OwnableRampProxy *OwnableRampProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnableRampProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnableRampProxyOwnershipTransferred)
				if err := _OwnableRampProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OwnableRampProxy *OwnableRampProxyFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableRampProxyOwnershipTransferred, error) {
	event := new(OwnableRampProxyOwnershipTransferred)
	if err := _OwnableRampProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OwnableRampProxyRampUpdatedIterator struct {
	Event *OwnableRampProxyRampUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnableRampProxyRampUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableRampProxyRampUpdated)
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
		it.Event = new(OwnableRampProxyRampUpdated)
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

func (it *OwnableRampProxyRampUpdatedIterator) Error() error {
	return it.fail
}

func (it *OwnableRampProxyRampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnableRampProxyRampUpdated struct {
	OldRamp common.Address
	NewRamp common.Address
	Raw     types.Log
}

func (_OwnableRampProxy *OwnableRampProxyFilterer) FilterRampUpdated(opts *bind.FilterOpts, oldRamp []common.Address, newRamp []common.Address) (*OwnableRampProxyRampUpdatedIterator, error) {

	var oldRampRule []interface{}
	for _, oldRampItem := range oldRamp {
		oldRampRule = append(oldRampRule, oldRampItem)
	}
	var newRampRule []interface{}
	for _, newRampItem := range newRamp {
		newRampRule = append(newRampRule, newRampItem)
	}

	logs, sub, err := _OwnableRampProxy.contract.FilterLogs(opts, "RampUpdated", oldRampRule, newRampRule)
	if err != nil {
		return nil, err
	}
	return &OwnableRampProxyRampUpdatedIterator{contract: _OwnableRampProxy.contract, event: "RampUpdated", logs: logs, sub: sub}, nil
}

func (_OwnableRampProxy *OwnableRampProxyFilterer) WatchRampUpdated(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyRampUpdated, oldRamp []common.Address, newRamp []common.Address) (event.Subscription, error) {

	var oldRampRule []interface{}
	for _, oldRampItem := range oldRamp {
		oldRampRule = append(oldRampRule, oldRampItem)
	}
	var newRampRule []interface{}
	for _, newRampItem := range newRamp {
		newRampRule = append(newRampRule, newRampItem)
	}

	logs, sub, err := _OwnableRampProxy.contract.WatchLogs(opts, "RampUpdated", oldRampRule, newRampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnableRampProxyRampUpdated)
				if err := _OwnableRampProxy.contract.UnpackLog(event, "RampUpdated", log); err != nil {
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

func (_OwnableRampProxy *OwnableRampProxyFilterer) ParseRampUpdated(log types.Log) (*OwnableRampProxyRampUpdated, error) {
	event := new(OwnableRampProxyRampUpdated)
	if err := _OwnableRampProxy.contract.UnpackLog(event, "RampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (OwnableRampProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OwnableRampProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OwnableRampProxyRampUpdated) Topic() common.Hash {
	return common.HexToHash("0x9a6b1be05a589b6fe11a413ab4892655cec2680cfa59a0eae85fbf90ed987969")
}

func (_OwnableRampProxy *OwnableRampProxy) Address() common.Address {
	return _OwnableRampProxy.address
}

type OwnableRampProxyInterface interface {
	Owner(opts *bind.CallOpts) (common.Address, error)

	SRamp(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetRamp(opts *bind.TransactOpts, rampAddress common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableRampProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OwnableRampProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableRampProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OwnableRampProxyOwnershipTransferred, error)

	FilterRampUpdated(opts *bind.FilterOpts, oldRamp []common.Address, newRamp []common.Address) (*OwnableRampProxyRampUpdatedIterator, error)

	WatchRampUpdated(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyRampUpdated, oldRamp []common.Address, newRamp []common.Address) (event.Subscription, error)

	ParseRampUpdated(log types.Log) (*OwnableRampProxyRampUpdated, error)

	Address() common.Address
}
