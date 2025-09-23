// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_proxy

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

var VerifierProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"verifierAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setVerifier\",\"inputs\":[{\"name\":\"verifierAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierUpdated\",\"inputs\":[{\"name\":\"oldVerifier\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newVerifier\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6080346100e457601f61056338819003918201601f19168301916001600160401b038311848410176100e9578084926020946040528339810103126100e457516001600160a01b038116908190036100e45733156100d357600180546001600160a01b0319163317905580156100c257600280546001600160a01b03198116831790915560405191906001600160a01b03167f0243549a92b2412f7a3caf7a2e56d65b8821b91345363faa5f57195384065fcc600080a361046390816101008239f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001a575b34156103ce575b600080fd5b60003560e01c806346657fe91461006a5780635437988d1461006557806379ba5097146100605780638da5cb5b1461005b5763f2fde38b0361000e57610307565b6102b5565b6101cc565b610108565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155773ffffffffffffffffffffffffffffffffffffffff6002541660805260206080f35b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60209101126100155760043573ffffffffffffffffffffffffffffffffffffffff811681036100155790565b346100155773ffffffffffffffffffffffffffffffffffffffff61012b366100bb565b61013361040b565b1680156101a25773ffffffffffffffffffffffffffffffffffffffff600254827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600255167f0243549a92b2412f7a3caf7a2e56d65b8821b91345363faa5f57195384065fcc600080a3005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760005473ffffffffffffffffffffffffffffffffffffffff8116330361028b577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100155773ffffffffffffffffffffffffffffffffffffffff61032a366100bb565b61033261040b565b163381146103a457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b60008073ffffffffffffffffffffffffffffffffffffffff600254163682803733600452818036925af13d6000803e610406573d6000fd5b3d6000f35b73ffffffffffffffffffffffffffffffffffffffff60015416330361042c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fdfea164736f6c634300081a000a",
}

var VerifierProxyABI = VerifierProxyMetaData.ABI

var VerifierProxyBin = VerifierProxyMetaData.Bin

func DeployVerifierProxy(auth *bind.TransactOpts, backend bind.ContractBackend, verifierAddress common.Address) (common.Address, *types.Transaction, *VerifierProxy, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierProxyBin), backend, verifierAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierProxy{address: address, abi: *parsed, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

type VerifierProxy struct {
	address common.Address
	abi     abi.ABI
	VerifierProxyCaller
	VerifierProxyTransactor
	VerifierProxyFilterer
}

type VerifierProxyCaller struct {
	contract *bind.BoundContract
}

type VerifierProxyTransactor struct {
	contract *bind.BoundContract
}

type VerifierProxyFilterer struct {
	contract *bind.BoundContract
}

type VerifierProxySession struct {
	Contract     *VerifierProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierProxyCallerSession struct {
	Contract *VerifierProxyCaller
	CallOpts bind.CallOpts
}

type VerifierProxyTransactorSession struct {
	Contract     *VerifierProxyTransactor
	TransactOpts bind.TransactOpts
}

type VerifierProxyRaw struct {
	Contract *VerifierProxy
}

type VerifierProxyCallerRaw struct {
	Contract *VerifierProxyCaller
}

type VerifierProxyTransactorRaw struct {
	Contract *VerifierProxyTransactor
}

func NewVerifierProxy(address common.Address, backend bind.ContractBackend) (*VerifierProxy, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierProxy{address: address, abi: abi, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

func NewVerifierProxyCaller(address common.Address, caller bind.ContractCaller) (*VerifierProxyCaller, error) {
	contract, err := bindVerifierProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCaller{contract: contract}, nil
}

func NewVerifierProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierProxyTransactor, error) {
	contract, err := bindVerifierProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyTransactor{contract: contract}, nil
}

func NewVerifierProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierProxyFilterer, error) {
	contract, err := bindVerifierProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFilterer{contract: contract}, nil
}

func bindVerifierProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierProxy *VerifierProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.VerifierProxyCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetVerifier() (common.Address, error) {
	return _VerifierProxy.Contract.GetVerifier(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetVerifier() (common.Address, error) {
	return _VerifierProxy.Contract.GetVerifier(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierProxy *VerifierProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) SetVerifier(opts *bind.TransactOpts, verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setVerifier", verifierAddress)
}

func (_VerifierProxy *VerifierProxySession) SetVerifier(verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetVerifier(&_VerifierProxy.TransactOpts, verifierAddress)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetVerifier(verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetVerifier(&_VerifierProxy.TransactOpts, verifierAddress)
}

func (_VerifierProxy *VerifierProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierProxy *VerifierProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _VerifierProxy.contract.RawTransact(opts, calldata)
}

func (_VerifierProxy *VerifierProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.Fallback(&_VerifierProxy.TransactOpts, calldata)
}

func (_VerifierProxy *VerifierProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.Fallback(&_VerifierProxy.TransactOpts, calldata)
}

type VerifierProxyOwnershipTransferRequestedIterator struct {
	Event *VerifierProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferRequested)
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
		it.Event = new(VerifierProxyOwnershipTransferRequested)
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

func (it *VerifierProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferRequestedIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferRequested)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error) {
	event := new(VerifierProxyOwnershipTransferRequested)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferredIterator struct {
	Event *VerifierProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferred)
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
		it.Event = new(VerifierProxyOwnershipTransferred)
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

func (it *VerifierProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferredIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferred)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error) {
	event := new(VerifierProxyOwnershipTransferred)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyVerifierUpdatedIterator struct {
	Event *VerifierProxyVerifierUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyVerifierUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyVerifierUpdated)
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
		it.Event = new(VerifierProxyVerifierUpdated)
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

func (it *VerifierProxyVerifierUpdatedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyVerifierUpdated struct {
	OldVerifier common.Address
	NewVerifier common.Address
	Raw         types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterVerifierUpdated(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*VerifierProxyVerifierUpdatedIterator, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "VerifierUpdated", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyVerifierUpdatedIterator{contract: _VerifierProxy.contract, event: "VerifierUpdated", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchVerifierUpdated(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierUpdated, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "VerifierUpdated", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyVerifierUpdated)
				if err := _VerifierProxy.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseVerifierUpdated(log types.Log) (*VerifierProxyVerifierUpdated, error) {
	event := new(VerifierProxyVerifierUpdated)
	if err := _VerifierProxy.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (VerifierProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VerifierProxyVerifierUpdated) Topic() common.Hash {
	return common.HexToHash("0x0243549a92b2412f7a3caf7a2e56d65b8821b91345363faa5f57195384065fcc")
}

func (_VerifierProxy *VerifierProxy) Address() common.Address {
	return _VerifierProxy.address
}

type VerifierProxyInterface interface {
	GetVerifier(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetVerifier(opts *bind.TransactOpts, verifierAddress common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error)

	FilterVerifierUpdated(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*VerifierProxyVerifierUpdatedIterator, error)

	WatchVerifierUpdated(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierUpdated, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error)

	ParseVerifierUpdated(log types.Log) (*VerifierProxyVerifierUpdated, error)

	Address() common.Address
}
