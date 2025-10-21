// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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

var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTarget\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setTarget\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TargetUpdated\",\"inputs\":[{\"name\":\"oldTarget\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTarget\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6080346100e457601f61056438819003918201601f19168301916001600160401b038311848410176100e9578084926020946040528339810103126100e457516001600160a01b038116908190036100e45733156100d357600180546001600160a01b0319163317905580156100c257600280546001600160a01b03198116831790915560405191906001600160a01b03167f331faca2e54d546f21863baefddc0bc0b9fe7554216f8798e32574255571713f600080a361046490816101008239f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001a575b34156103cf575b600080fd5b60003560e01c8063776d1a011461006a57806379ba5097146100655780638da5cb5b14610060578063f00e6a2a1461005b5763f2fde38b0361000e57610308565b6102b6565b610264565b61017b565b346100155773ffffffffffffffffffffffffffffffffffffffff61008d3661012e565b61009561040c565b1680156101045773ffffffffffffffffffffffffffffffffffffffff600254827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600255167f331faca2e54d546f21863baefddc0bc0b9fe7554216f8798e32574255571713f600080a3005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60209101126100155760043573ffffffffffffffffffffffffffffffffffffffff811681036100155790565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760005473ffffffffffffffffffffffffffffffffffffffff8116330361023a577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b346100155773ffffffffffffffffffffffffffffffffffffffff61032b3661012e565b61033361040c565b163381146103a557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b60008073ffffffffffffffffffffffffffffffffffffffff600254163682803733600452818036925af13d6000803e610407573d6000fd5b3d6000f35b73ffffffffffffffffffffffffffffffffffffffff60015416330361042d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fdfea164736f6c634300081a000a",
}

var ProxyABI = ProxyMetaData.ABI

var ProxyBin = ProxyMetaData.Bin

func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, target common.Address) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyBin), backend, target)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{address: address, abi: *parsed, ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

type Proxy struct {
	address common.Address
	abi     abi.ABI
	ProxyCaller
	ProxyTransactor
	ProxyFilterer
}

type ProxyCaller struct {
	contract *bind.BoundContract
}

type ProxyTransactor struct {
	contract *bind.BoundContract
}

type ProxyFilterer struct {
	contract *bind.BoundContract
}

type ProxySession struct {
	Contract     *Proxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ProxyCallerSession struct {
	Contract *ProxyCaller
	CallOpts bind.CallOpts
}

type ProxyTransactorSession struct {
	Contract     *ProxyTransactor
	TransactOpts bind.TransactOpts
}

type ProxyRaw struct {
	Contract *Proxy
}

type ProxyCallerRaw struct {
	Contract *ProxyCaller
}

type ProxyTransactorRaw struct {
	Contract *ProxyTransactor
}

func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	abi, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{address: address, abi: abi, ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

func (_Proxy *ProxyCaller) GetTarget(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getTarget")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Proxy *ProxySession) GetTarget() (common.Address, error) {
	return _Proxy.Contract.GetTarget(&_Proxy.CallOpts)
}

func (_Proxy *ProxyCallerSession) GetTarget() (common.Address, error) {
	return _Proxy.Contract.GetTarget(&_Proxy.CallOpts)
}

func (_Proxy *ProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Proxy *ProxySession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

func (_Proxy *ProxyCallerSession) Owner() (common.Address, error) {
	return _Proxy.Contract.Owner(&_Proxy.CallOpts)
}

func (_Proxy *ProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "acceptOwnership")
}

func (_Proxy *ProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxy.Contract.AcceptOwnership(&_Proxy.TransactOpts)
}

func (_Proxy *ProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Proxy.Contract.AcceptOwnership(&_Proxy.TransactOpts)
}

func (_Proxy *ProxyTransactor) SetTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setTarget", target)
}

func (_Proxy *ProxySession) SetTarget(target common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetTarget(&_Proxy.TransactOpts, target)
}

func (_Proxy *ProxyTransactorSession) SetTarget(target common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetTarget(&_Proxy.TransactOpts, target)
}

func (_Proxy *ProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "transferOwnership", to)
}

func (_Proxy *ProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.TransferOwnership(&_Proxy.TransactOpts, to)
}

func (_Proxy *ProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.TransferOwnership(&_Proxy.TransactOpts, to)
}

func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

func (_Proxy *ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

func (_Proxy *ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

type ProxyOwnershipTransferRequestedIterator struct {
	Event *ProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOwnershipTransferRequested)
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
		it.Event = new(ProxyOwnershipTransferRequested)
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

func (it *ProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Proxy *ProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ProxyOwnershipTransferRequestedIterator{contract: _Proxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Proxy *ProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ProxyOwnershipTransferRequested)
				if err := _Proxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Proxy *ProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*ProxyOwnershipTransferRequested, error) {
	event := new(ProxyOwnershipTransferRequested)
	if err := _Proxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ProxyOwnershipTransferredIterator struct {
	Event *ProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOwnershipTransferred)
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
		it.Event = new(ProxyOwnershipTransferred)
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

func (it *ProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Proxy *ProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ProxyOwnershipTransferredIterator{contract: _Proxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Proxy *ProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ProxyOwnershipTransferred)
				if err := _Proxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Proxy *ProxyFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyOwnershipTransferred, error) {
	event := new(ProxyOwnershipTransferred)
	if err := _Proxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ProxyTargetUpdatedIterator struct {
	Event *ProxyTargetUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ProxyTargetUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyTargetUpdated)
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
		it.Event = new(ProxyTargetUpdated)
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

func (it *ProxyTargetUpdatedIterator) Error() error {
	return it.fail
}

func (it *ProxyTargetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ProxyTargetUpdated struct {
	OldTarget common.Address
	NewTarget common.Address
	Raw       types.Log
}

func (_Proxy *ProxyFilterer) FilterTargetUpdated(opts *bind.FilterOpts, oldTarget []common.Address, newTarget []common.Address) (*ProxyTargetUpdatedIterator, error) {

	var oldTargetRule []interface{}
	for _, oldTargetItem := range oldTarget {
		oldTargetRule = append(oldTargetRule, oldTargetItem)
	}
	var newTargetRule []interface{}
	for _, newTargetItem := range newTarget {
		newTargetRule = append(newTargetRule, newTargetItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "TargetUpdated", oldTargetRule, newTargetRule)
	if err != nil {
		return nil, err
	}
	return &ProxyTargetUpdatedIterator{contract: _Proxy.contract, event: "TargetUpdated", logs: logs, sub: sub}, nil
}

func (_Proxy *ProxyFilterer) WatchTargetUpdated(opts *bind.WatchOpts, sink chan<- *ProxyTargetUpdated, oldTarget []common.Address, newTarget []common.Address) (event.Subscription, error) {

	var oldTargetRule []interface{}
	for _, oldTargetItem := range oldTarget {
		oldTargetRule = append(oldTargetRule, oldTargetItem)
	}
	var newTargetRule []interface{}
	for _, newTargetItem := range newTarget {
		newTargetRule = append(newTargetRule, newTargetItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "TargetUpdated", oldTargetRule, newTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ProxyTargetUpdated)
				if err := _Proxy.contract.UnpackLog(event, "TargetUpdated", log); err != nil {
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

func (_Proxy *ProxyFilterer) ParseTargetUpdated(log types.Log) (*ProxyTargetUpdated, error) {
	event := new(ProxyTargetUpdated)
	if err := _Proxy.contract.UnpackLog(event, "TargetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ProxyTargetUpdated) Topic() common.Hash {
	return common.HexToHash("0x331faca2e54d546f21863baefddc0bc0b9fe7554216f8798e32574255571713f")
}

func (_Proxy *Proxy) Address() common.Address {
	return _Proxy.address
}

type ProxyInterface interface {
	GetTarget(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ProxyOwnershipTransferred, error)

	FilterTargetUpdated(opts *bind.FilterOpts, oldTarget []common.Address, newTarget []common.Address) (*ProxyTargetUpdatedIterator, error)

	WatchTargetUpdated(opts *bind.WatchOpts, sink chan<- *ProxyTargetUpdated, oldTarget []common.Address, newTarget []common.Address) (event.Subscription, error)

	ParseTargetUpdated(log types.Log) (*ProxyTargetUpdated, error)

	Address() common.Address
}
