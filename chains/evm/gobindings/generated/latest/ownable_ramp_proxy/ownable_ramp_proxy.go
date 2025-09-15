// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownable_ccv_ramp_proxy

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

type RampProxySetRampsArgs struct {
	RemoteChainSelector uint64
	RampAddress         common.Address
	Version             [32]byte
}

var OwnableRampProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRamp\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRamps\",\"inputs\":[{\"name\":\"ramps\",\"type\":\"tuple[]\",\"internalType\":\"structRampProxy.SetRampsArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rampAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RampSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rampAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RampNotFound\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60808060405234603d573315602c57600280546001600160a01b0319163317905561088090816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe6080604052600436101561001a575b34156106b9575b600080fd5b60003560e01c806306d9f1ff1461007a578063181f5a771461007557806379ba5097146100705780638da5cb5b1461006b578063f0405f45146100665763f2fde38b0361000e576103d6565b610354565b610302565b610219565b61014b565b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760043567ffffffffffffffff8111610015573660238201121561001557806004013567ffffffffffffffff81116100155736602460608302840101116100155760246100f592016104df565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051906040820182811067ffffffffffffffff82111761014657604052565b6100f7565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557610182610126565b601681527f43435652616d7050726f787920312e372e302d64657600000000000000000000602082015260405190602082528181519182602083015260005b8381106102015750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b602082820181015160408784010152859350016101c1565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760015473ffffffffffffffffffffffffffffffffffffffff811633036102d8577fffffffffffffffffffffffff00000000000000000000000000000000000000006002549133828416176002551660015573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b346100155760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760043567ffffffffffffffff811680910361001557602435906000526000602052604060002090600052602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760043573ffffffffffffffffffffffffffffffffffffffff811681036100155773ffffffffffffffffffffffffffffffffffffffff90610443610777565b163381146104b557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff600254167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906104e8610777565b60005b8181106104f757505050565b61050a6105058284866107c2565b610801565b805167ffffffffffffffff16801561068257506040810190815180156106555750908167ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff61062761060a60206001989701946105fc61057a875173ffffffffffffffffffffffffffffffffffffffff1690565b6105bc6105ac610592855167ffffffffffffffff1690565b67ffffffffffffffff166000526000602052604060002090565b8a51600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5167ffffffffffffffff1690565b9451935173ffffffffffffffffffffffffffffffffffffffff1690565b1692167f61079add8b3485b65ad33a15cebc6188cca5cb506a21df75784554c7339d3584600080a4016104eb565b7fe84925c70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f77b160700000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b60043560243567ffffffffffffffff8216600052600060205260406000208160005260205273ffffffffffffffffffffffffffffffffffffffff604060002054169173ffffffffffffffffffffffffffffffffffffffff83161561073c57600080843682803733604452818036925af13d6000803e610737573d6000fd5b3d6000f35b67ffffffffffffffff907fc6a64cca000000000000000000000000000000000000000000000000000000006000521660045260245260446000fd5b73ffffffffffffffffffffffffffffffffffffffff60025416330361079857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156107d2576060020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60608136031261001557604051906060820182811067ffffffffffffffff82111761014657604052803567ffffffffffffffff8116810361001557825260208101359073ffffffffffffffffffffffffffffffffffffffff82168203610015576040916020840152013560408201529056fea164736f6c634300081a000a",
}

var OwnableRampProxyABI = OwnableRampProxyMetaData.ABI

var OwnableRampProxyBin = OwnableRampProxyMetaData.Bin

func DeployOwnableRampProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OwnableRampProxy, error) {
	parsed, err := OwnableRampProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnableRampProxyBin), backend)
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

func (_OwnableRampProxy *OwnableRampProxyCaller) GetRamp(opts *bind.CallOpts, remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	var out []interface{}
	err := _OwnableRampProxy.contract.Call(opts, &out, "getRamp", remoteChainSelector, version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OwnableRampProxy *OwnableRampProxySession) GetRamp(remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	return _OwnableRampProxy.Contract.GetRamp(&_OwnableRampProxy.CallOpts, remoteChainSelector, version)
}

func (_OwnableRampProxy *OwnableRampProxyCallerSession) GetRamp(remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	return _OwnableRampProxy.Contract.GetRamp(&_OwnableRampProxy.CallOpts, remoteChainSelector, version)
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

func (_OwnableRampProxy *OwnableRampProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OwnableRampProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OwnableRampProxy *OwnableRampProxySession) TypeAndVersion() (string, error) {
	return _OwnableRampProxy.Contract.TypeAndVersion(&_OwnableRampProxy.CallOpts)
}

func (_OwnableRampProxy *OwnableRampProxyCallerSession) TypeAndVersion() (string, error) {
	return _OwnableRampProxy.Contract.TypeAndVersion(&_OwnableRampProxy.CallOpts)
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

func (_OwnableRampProxy *OwnableRampProxyTransactor) SetRamps(opts *bind.TransactOpts, ramps []RampProxySetRampsArgs) (*types.Transaction, error) {
	return _OwnableRampProxy.contract.Transact(opts, "setRamps", ramps)
}

func (_OwnableRampProxy *OwnableRampProxySession) SetRamps(ramps []RampProxySetRampsArgs) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.SetRamps(&_OwnableRampProxy.TransactOpts, ramps)
}

func (_OwnableRampProxy *OwnableRampProxyTransactorSession) SetRamps(ramps []RampProxySetRampsArgs) (*types.Transaction, error) {
	return _OwnableRampProxy.Contract.SetRamps(&_OwnableRampProxy.TransactOpts, ramps)
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

type OwnableRampProxyRampSetIterator struct {
	Event *OwnableRampProxyRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnableRampProxyRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableRampProxyRampSet)
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
		it.Event = new(OwnableRampProxyRampSet)
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

func (it *OwnableRampProxyRampSetIterator) Error() error {
	return it.fail
}

func (it *OwnableRampProxyRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnableRampProxyRampSet struct {
	RemoteChainSelector uint64
	Version             [32]byte
	RampAddress         common.Address
	Raw                 types.Log
}

func (_OwnableRampProxy *OwnableRampProxyFilterer) FilterRampSet(opts *bind.FilterOpts, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (*OwnableRampProxyRampSetIterator, error) {

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

	logs, sub, err := _OwnableRampProxy.contract.FilterLogs(opts, "RampSet", remoteChainSelectorRule, versionRule, rampAddressRule)
	if err != nil {
		return nil, err
	}
	return &OwnableRampProxyRampSetIterator{contract: _OwnableRampProxy.contract, event: "RampSet", logs: logs, sub: sub}, nil
}

func (_OwnableRampProxy *OwnableRampProxyFilterer) WatchRampSet(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyRampSet, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OwnableRampProxy.contract.WatchLogs(opts, "RampSet", remoteChainSelectorRule, versionRule, rampAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnableRampProxyRampSet)
				if err := _OwnableRampProxy.contract.UnpackLog(event, "RampSet", log); err != nil {
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

func (_OwnableRampProxy *OwnableRampProxyFilterer) ParseRampSet(log types.Log) (*OwnableRampProxyRampSet, error) {
	event := new(OwnableRampProxyRampSet)
	if err := _OwnableRampProxy.contract.UnpackLog(event, "RampSet", log); err != nil {
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

func (OwnableRampProxyRampSet) Topic() common.Hash {
	return common.HexToHash("0x61079add8b3485b65ad33a15cebc6188cca5cb506a21df75784554c7339d3584")
}

func (_OwnableRampProxy *OwnableRampProxy) Address() common.Address {
	return _OwnableRampProxy.address
}

type OwnableRampProxyInterface interface {
	GetRamp(opts *bind.CallOpts, remoteChainSelector uint64, version [32]byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetRamps(opts *bind.TransactOpts, ramps []RampProxySetRampsArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableRampProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OwnableRampProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableRampProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OwnableRampProxyOwnershipTransferred, error)

	FilterRampSet(opts *bind.FilterOpts, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (*OwnableRampProxyRampSetIterator, error)

	WatchRampSet(opts *bind.WatchOpts, sink chan<- *OwnableRampProxyRampSet, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (event.Subscription, error)

	ParseRampSet(log types.Log) (*OwnableRampProxyRampSet, error)

	Address() common.Address
}
