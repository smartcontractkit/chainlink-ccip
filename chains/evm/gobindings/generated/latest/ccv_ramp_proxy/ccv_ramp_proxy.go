// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccv_ramp_proxy

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

var CCVRampProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRamp\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRamp\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rampAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RampSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rampAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRampAddress\",\"inputs\":[{\"name\":\"rampAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownRamp\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b0319163317905561075690816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe6080604052600436101561001a575b3415610640575b600080fd5b60003560e01c8063181f5a771461007a57806379ba5097146100755780638da5cb5b14610070578063b41cc0261461006b578063f0405f45146100665763f2fde38b0361000e5761054d565b6104ce565b610311565b610285565b61019c565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557604051604081019080821067ffffffffffffffff8311176101055761010191604052601681527f43435652616d7050726f787920312e372e302d64657600000000000000000000602082015260405191829182610134565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b9190916020815282519283602083015260005b8481106101865750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b8060208092840101516040828601015201610147565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760005473ffffffffffffffffffffffffffffffffffffffff8116330361025b577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b6004359067ffffffffffffffff8216820361001557565b6044359073ffffffffffffffffffffffffffffffffffffffff8216820361001557565b346100155760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610015576103486102d7565b6024356103536102ee565b61035b6106fe565b67ffffffffffffffff83169081156104a05782156104715773ffffffffffffffffffffffffffffffffffffffff811693841561042d57906103c6846103b76104069467ffffffffffffffff166000526002602052604060002090565b90600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b7f61079add8b3485b65ad33a15cebc6188cca5cb506a21df75784554c7339d3584600080a4005b7f2cc5133b0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff821660045260246000fd5b7fe84925c700000000000000000000000000000000000000000000000000000000600052600483905260246000fd5b507f77b160700000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346100155760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610015576105056102d7565b67ffffffffffffffff60243591166000526002602052604060002090600052602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760043573ffffffffffffffffffffffffffffffffffffffff8116809103610015576105a56106fe565b33811461061657807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b60043560243567ffffffffffffffff8216600052600260205260406000208160005260205273ffffffffffffffffffffffffffffffffffffffff604060002054169173ffffffffffffffffffffffffffffffffffffffff8316156106c357600080843682803733604452818036925af13d6000803e6106be573d6000fd5b3d6000f35b67ffffffffffffffff907f2923343e000000000000000000000000000000000000000000000000000000006000521660045260245260446000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361071f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fdfea164736f6c634300081a000a",
}

var CCVRampProxyABI = CCVRampProxyMetaData.ABI

var CCVRampProxyBin = CCVRampProxyMetaData.Bin

func DeployCCVRampProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CCVRampProxy, error) {
	parsed, err := CCVRampProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCVRampProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCVRampProxy{address: address, abi: *parsed, CCVRampProxyCaller: CCVRampProxyCaller{contract: contract}, CCVRampProxyTransactor: CCVRampProxyTransactor{contract: contract}, CCVRampProxyFilterer: CCVRampProxyFilterer{contract: contract}}, nil
}

type CCVRampProxy struct {
	address common.Address
	abi     abi.ABI
	CCVRampProxyCaller
	CCVRampProxyTransactor
	CCVRampProxyFilterer
}

type CCVRampProxyCaller struct {
	contract *bind.BoundContract
}

type CCVRampProxyTransactor struct {
	contract *bind.BoundContract
}

type CCVRampProxyFilterer struct {
	contract *bind.BoundContract
}

type CCVRampProxySession struct {
	Contract     *CCVRampProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCVRampProxyCallerSession struct {
	Contract *CCVRampProxyCaller
	CallOpts bind.CallOpts
}

type CCVRampProxyTransactorSession struct {
	Contract     *CCVRampProxyTransactor
	TransactOpts bind.TransactOpts
}

type CCVRampProxyRaw struct {
	Contract *CCVRampProxy
}

type CCVRampProxyCallerRaw struct {
	Contract *CCVRampProxyCaller
}

type CCVRampProxyTransactorRaw struct {
	Contract *CCVRampProxyTransactor
}

func NewCCVRampProxy(address common.Address, backend bind.ContractBackend) (*CCVRampProxy, error) {
	abi, err := abi.JSON(strings.NewReader(CCVRampProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCVRampProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCVRampProxy{address: address, abi: abi, CCVRampProxyCaller: CCVRampProxyCaller{contract: contract}, CCVRampProxyTransactor: CCVRampProxyTransactor{contract: contract}, CCVRampProxyFilterer: CCVRampProxyFilterer{contract: contract}}, nil
}

func NewCCVRampProxyCaller(address common.Address, caller bind.ContractCaller) (*CCVRampProxyCaller, error) {
	contract, err := bindCCVRampProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCVRampProxyCaller{contract: contract}, nil
}

func NewCCVRampProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CCVRampProxyTransactor, error) {
	contract, err := bindCCVRampProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCVRampProxyTransactor{contract: contract}, nil
}

func NewCCVRampProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CCVRampProxyFilterer, error) {
	contract, err := bindCCVRampProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCVRampProxyFilterer{contract: contract}, nil
}

func bindCCVRampProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCVRampProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCVRampProxy *CCVRampProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCVRampProxy.Contract.CCVRampProxyCaller.contract.Call(opts, result, method, params...)
}

func (_CCVRampProxy *CCVRampProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.CCVRampProxyTransactor.contract.Transfer(opts)
}

func (_CCVRampProxy *CCVRampProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.CCVRampProxyTransactor.contract.Transact(opts, method, params...)
}

func (_CCVRampProxy *CCVRampProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCVRampProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_CCVRampProxy *CCVRampProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.contract.Transfer(opts)
}

func (_CCVRampProxy *CCVRampProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.contract.Transact(opts, method, params...)
}

func (_CCVRampProxy *CCVRampProxyCaller) GetRamp(opts *bind.CallOpts, remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	var out []interface{}
	err := _CCVRampProxy.contract.Call(opts, &out, "getRamp", remoteChainSelector, version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCVRampProxy *CCVRampProxySession) GetRamp(remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	return _CCVRampProxy.Contract.GetRamp(&_CCVRampProxy.CallOpts, remoteChainSelector, version)
}

func (_CCVRampProxy *CCVRampProxyCallerSession) GetRamp(remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	return _CCVRampProxy.Contract.GetRamp(&_CCVRampProxy.CallOpts, remoteChainSelector, version)
}

func (_CCVRampProxy *CCVRampProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCVRampProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCVRampProxy *CCVRampProxySession) Owner() (common.Address, error) {
	return _CCVRampProxy.Contract.Owner(&_CCVRampProxy.CallOpts)
}

func (_CCVRampProxy *CCVRampProxyCallerSession) Owner() (common.Address, error) {
	return _CCVRampProxy.Contract.Owner(&_CCVRampProxy.CallOpts)
}

func (_CCVRampProxy *CCVRampProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCVRampProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCVRampProxy *CCVRampProxySession) TypeAndVersion() (string, error) {
	return _CCVRampProxy.Contract.TypeAndVersion(&_CCVRampProxy.CallOpts)
}

func (_CCVRampProxy *CCVRampProxyCallerSession) TypeAndVersion() (string, error) {
	return _CCVRampProxy.Contract.TypeAndVersion(&_CCVRampProxy.CallOpts)
}

func (_CCVRampProxy *CCVRampProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVRampProxy.contract.Transact(opts, "acceptOwnership")
}

func (_CCVRampProxy *CCVRampProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _CCVRampProxy.Contract.AcceptOwnership(&_CCVRampProxy.TransactOpts)
}

func (_CCVRampProxy *CCVRampProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCVRampProxy.Contract.AcceptOwnership(&_CCVRampProxy.TransactOpts)
}

func (_CCVRampProxy *CCVRampProxyTransactor) SetRamp(opts *bind.TransactOpts, remoteChainSelector uint64, version [32]byte, rampAddress common.Address) (*types.Transaction, error) {
	return _CCVRampProxy.contract.Transact(opts, "setRamp", remoteChainSelector, version, rampAddress)
}

func (_CCVRampProxy *CCVRampProxySession) SetRamp(remoteChainSelector uint64, version [32]byte, rampAddress common.Address) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.SetRamp(&_CCVRampProxy.TransactOpts, remoteChainSelector, version, rampAddress)
}

func (_CCVRampProxy *CCVRampProxyTransactorSession) SetRamp(remoteChainSelector uint64, version [32]byte, rampAddress common.Address) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.SetRamp(&_CCVRampProxy.TransactOpts, remoteChainSelector, version, rampAddress)
}

func (_CCVRampProxy *CCVRampProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCVRampProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_CCVRampProxy *CCVRampProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.TransferOwnership(&_CCVRampProxy.TransactOpts, to)
}

func (_CCVRampProxy *CCVRampProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.TransferOwnership(&_CCVRampProxy.TransactOpts, to)
}

func (_CCVRampProxy *CCVRampProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CCVRampProxy.contract.RawTransact(opts, calldata)
}

func (_CCVRampProxy *CCVRampProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.Fallback(&_CCVRampProxy.TransactOpts, calldata)
}

func (_CCVRampProxy *CCVRampProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CCVRampProxy.Contract.Fallback(&_CCVRampProxy.TransactOpts, calldata)
}

type CCVRampProxyOwnershipTransferRequestedIterator struct {
	Event *CCVRampProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVRampProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVRampProxyOwnershipTransferRequested)
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
		it.Event = new(CCVRampProxyOwnershipTransferRequested)
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

func (it *CCVRampProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCVRampProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVRampProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCVRampProxy *CCVRampProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVRampProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVRampProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCVRampProxyOwnershipTransferRequestedIterator{contract: _CCVRampProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCVRampProxy *CCVRampProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCVRampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVRampProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVRampProxyOwnershipTransferRequested)
				if err := _CCVRampProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCVRampProxy *CCVRampProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCVRampProxyOwnershipTransferRequested, error) {
	event := new(CCVRampProxyOwnershipTransferRequested)
	if err := _CCVRampProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVRampProxyOwnershipTransferredIterator struct {
	Event *CCVRampProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVRampProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVRampProxyOwnershipTransferred)
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
		it.Event = new(CCVRampProxyOwnershipTransferred)
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

func (it *CCVRampProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCVRampProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVRampProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCVRampProxy *CCVRampProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVRampProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVRampProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCVRampProxyOwnershipTransferredIterator{contract: _CCVRampProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCVRampProxy *CCVRampProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCVRampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVRampProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVRampProxyOwnershipTransferred)
				if err := _CCVRampProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCVRampProxy *CCVRampProxyFilterer) ParseOwnershipTransferred(log types.Log) (*CCVRampProxyOwnershipTransferred, error) {
	event := new(CCVRampProxyOwnershipTransferred)
	if err := _CCVRampProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVRampProxyRampSetIterator struct {
	Event *CCVRampProxyRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVRampProxyRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVRampProxyRampSet)
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
		it.Event = new(CCVRampProxyRampSet)
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

func (it *CCVRampProxyRampSetIterator) Error() error {
	return it.fail
}

func (it *CCVRampProxyRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVRampProxyRampSet struct {
	RemoteChainSelector uint64
	Version             [32]byte
	RampAddress         common.Address
	Raw                 types.Log
}

func (_CCVRampProxy *CCVRampProxyFilterer) FilterRampSet(opts *bind.FilterOpts, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (*CCVRampProxyRampSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var rampAddressRule []interface{}
	for _, rampAddressItem := range rampAddress {
		rampAddressRule = append(rampAddressRule, rampAddressItem)
	}

	logs, sub, err := _CCVRampProxy.contract.FilterLogs(opts, "RampSet", remoteChainSelectorRule, versionRule, rampAddressRule)
	if err != nil {
		return nil, err
	}
	return &CCVRampProxyRampSetIterator{contract: _CCVRampProxy.contract, event: "RampSet", logs: logs, sub: sub}, nil
}

func (_CCVRampProxy *CCVRampProxyFilterer) WatchRampSet(opts *bind.WatchOpts, sink chan<- *CCVRampProxyRampSet, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var rampAddressRule []interface{}
	for _, rampAddressItem := range rampAddress {
		rampAddressRule = append(rampAddressRule, rampAddressItem)
	}

	logs, sub, err := _CCVRampProxy.contract.WatchLogs(opts, "RampSet", remoteChainSelectorRule, versionRule, rampAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVRampProxyRampSet)
				if err := _CCVRampProxy.contract.UnpackLog(event, "RampSet", log); err != nil {
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

func (_CCVRampProxy *CCVRampProxyFilterer) ParseRampSet(log types.Log) (*CCVRampProxyRampSet, error) {
	event := new(CCVRampProxyRampSet)
	if err := _CCVRampProxy.contract.UnpackLog(event, "RampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (CCVRampProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCVRampProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CCVRampProxyRampSet) Topic() common.Hash {
	return common.HexToHash("0x61079add8b3485b65ad33a15cebc6188cca5cb506a21df75784554c7339d3584")
}

func (_CCVRampProxy *CCVRampProxy) Address() common.Address {
	return _CCVRampProxy.address
}

type CCVRampProxyInterface interface {
	GetRamp(opts *bind.CallOpts, remoteChainSelector uint64, version [32]byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetRamp(opts *bind.TransactOpts, remoteChainSelector uint64, version [32]byte, rampAddress common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVRampProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCVRampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCVRampProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVRampProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCVRampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCVRampProxyOwnershipTransferred, error)

	FilterRampSet(opts *bind.FilterOpts, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (*CCVRampProxyRampSetIterator, error)

	WatchRampSet(opts *bind.WatchOpts, sink chan<- *CCVRampProxyRampSet, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (event.Subscription, error)

	ParseRampSet(log types.Log) (*CCVRampProxyRampSet, error)

	Address() common.Address
}
