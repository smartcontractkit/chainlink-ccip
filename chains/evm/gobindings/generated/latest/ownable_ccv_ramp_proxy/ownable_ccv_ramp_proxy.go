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

type CCVRampProxySetRampsArgs struct {
	RemoteChainSelector uint64
	RampAddress         common.Address
	Version             [32]byte
}

var OwnableCCVRampProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRamp\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRamps\",\"inputs\":[{\"name\":\"ramps\",\"type\":\"tuple[]\",\"internalType\":\"structCCVRampProxy.SetRampsArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rampAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RampSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rampAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RampNotFound\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60808060405234603d573315602c57600280546001600160a01b0319163317905561088090816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe6080604052600436101561001a575b34156106b9575b600080fd5b60003560e01c806306d9f1ff1461007a578063181f5a771461007557806379ba5097146100705780638da5cb5b1461006b578063f0405f45146100665763f2fde38b0361000e576103d6565b610354565b610302565b610219565b61014b565b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760043567ffffffffffffffff8111610015573660238201121561001557806004013567ffffffffffffffff81116100155736602460608302840101116100155760246100f592016104df565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051906040820182811067ffffffffffffffff82111761014657604052565b6100f7565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557610182610126565b601681527f43435652616d7050726f787920312e372e302d64657600000000000000000000602082015260405190602082528181519182602083015260005b8381106102015750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b602082820181015160408784010152859350016101c1565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760015473ffffffffffffffffffffffffffffffffffffffff811633036102d8577fffffffffffffffffffffffff00000000000000000000000000000000000000006002549133828416176002551660015573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b346100155760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760043567ffffffffffffffff811680910361001557602435906000526000602052604060002090600052602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760043573ffffffffffffffffffffffffffffffffffffffff811681036100155773ffffffffffffffffffffffffffffffffffffffff90610443610777565b163381146104b557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff600254167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906104e8610777565b60005b8181106104f757505050565b61050a6105058284866107c2565b610801565b805167ffffffffffffffff16801561068257506040810190815180156106555750908167ffffffffffffffff73ffffffffffffffffffffffffffffffffffffffff61062761060a60206001989701946105fc61057a875173ffffffffffffffffffffffffffffffffffffffff1690565b6105bc6105ac610592855167ffffffffffffffff1690565b67ffffffffffffffff166000526000602052604060002090565b8a51600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5167ffffffffffffffff1690565b9451935173ffffffffffffffffffffffffffffffffffffffff1690565b1692167f61079add8b3485b65ad33a15cebc6188cca5cb506a21df75784554c7339d3584600080a4016104eb565b7fe84925c70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f77b160700000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b60043560243567ffffffffffffffff8216600052600060205260406000208160005260205273ffffffffffffffffffffffffffffffffffffffff604060002054169173ffffffffffffffffffffffffffffffffffffffff83161561073c57600080843682803733604452818036925af13d6000803e610737573d6000fd5b3d6000f35b67ffffffffffffffff907fc6a64cca000000000000000000000000000000000000000000000000000000006000521660045260245260446000fd5b73ffffffffffffffffffffffffffffffffffffffff60025416330361079857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156107d2576060020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60608136031261001557604051906060820182811067ffffffffffffffff82111761014657604052803567ffffffffffffffff8116810361001557825260208101359073ffffffffffffffffffffffffffffffffffffffff82168203610015576040916020840152013560408201529056fea164736f6c634300081a000a",
}

var OwnableCCVRampProxyABI = OwnableCCVRampProxyMetaData.ABI

var OwnableCCVRampProxyBin = OwnableCCVRampProxyMetaData.Bin

func DeployOwnableCCVRampProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OwnableCCVRampProxy, error) {
	parsed, err := OwnableCCVRampProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnableCCVRampProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OwnableCCVRampProxy{address: address, abi: *parsed, OwnableCCVRampProxyCaller: OwnableCCVRampProxyCaller{contract: contract}, OwnableCCVRampProxyTransactor: OwnableCCVRampProxyTransactor{contract: contract}, OwnableCCVRampProxyFilterer: OwnableCCVRampProxyFilterer{contract: contract}}, nil
}

type OwnableCCVRampProxy struct {
	address common.Address
	abi     abi.ABI
	OwnableCCVRampProxyCaller
	OwnableCCVRampProxyTransactor
	OwnableCCVRampProxyFilterer
}

type OwnableCCVRampProxyCaller struct {
	contract *bind.BoundContract
}

type OwnableCCVRampProxyTransactor struct {
	contract *bind.BoundContract
}

type OwnableCCVRampProxyFilterer struct {
	contract *bind.BoundContract
}

type OwnableCCVRampProxySession struct {
	Contract     *OwnableCCVRampProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OwnableCCVRampProxyCallerSession struct {
	Contract *OwnableCCVRampProxyCaller
	CallOpts bind.CallOpts
}

type OwnableCCVRampProxyTransactorSession struct {
	Contract     *OwnableCCVRampProxyTransactor
	TransactOpts bind.TransactOpts
}

type OwnableCCVRampProxyRaw struct {
	Contract *OwnableCCVRampProxy
}

type OwnableCCVRampProxyCallerRaw struct {
	Contract *OwnableCCVRampProxyCaller
}

type OwnableCCVRampProxyTransactorRaw struct {
	Contract *OwnableCCVRampProxyTransactor
}

func NewOwnableCCVRampProxy(address common.Address, backend bind.ContractBackend) (*OwnableCCVRampProxy, error) {
	abi, err := abi.JSON(strings.NewReader(OwnableCCVRampProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOwnableCCVRampProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableCCVRampProxy{address: address, abi: abi, OwnableCCVRampProxyCaller: OwnableCCVRampProxyCaller{contract: contract}, OwnableCCVRampProxyTransactor: OwnableCCVRampProxyTransactor{contract: contract}, OwnableCCVRampProxyFilterer: OwnableCCVRampProxyFilterer{contract: contract}}, nil
}

func NewOwnableCCVRampProxyCaller(address common.Address, caller bind.ContractCaller) (*OwnableCCVRampProxyCaller, error) {
	contract, err := bindOwnableCCVRampProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCCVRampProxyCaller{contract: contract}, nil
}

func NewOwnableCCVRampProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableCCVRampProxyTransactor, error) {
	contract, err := bindOwnableCCVRampProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCCVRampProxyTransactor{contract: contract}, nil
}

func NewOwnableCCVRampProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableCCVRampProxyFilterer, error) {
	contract, err := bindOwnableCCVRampProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableCCVRampProxyFilterer{contract: contract}, nil
}

func bindOwnableCCVRampProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableCCVRampProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableCCVRampProxy.Contract.OwnableCCVRampProxyCaller.contract.Call(opts, result, method, params...)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.OwnableCCVRampProxyTransactor.contract.Transfer(opts)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.OwnableCCVRampProxyTransactor.contract.Transact(opts, method, params...)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableCCVRampProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.contract.Transfer(opts)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.contract.Transact(opts, method, params...)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyCaller) GetRamp(opts *bind.CallOpts, remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	var out []interface{}
	err := _OwnableCCVRampProxy.contract.Call(opts, &out, "getRamp", remoteChainSelector, version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OwnableCCVRampProxy *OwnableCCVRampProxySession) GetRamp(remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	return _OwnableCCVRampProxy.Contract.GetRamp(&_OwnableCCVRampProxy.CallOpts, remoteChainSelector, version)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyCallerSession) GetRamp(remoteChainSelector uint64, version [32]byte) (common.Address, error) {
	return _OwnableCCVRampProxy.Contract.GetRamp(&_OwnableCCVRampProxy.CallOpts, remoteChainSelector, version)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableCCVRampProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OwnableCCVRampProxy *OwnableCCVRampProxySession) Owner() (common.Address, error) {
	return _OwnableCCVRampProxy.Contract.Owner(&_OwnableCCVRampProxy.CallOpts)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyCallerSession) Owner() (common.Address, error) {
	return _OwnableCCVRampProxy.Contract.Owner(&_OwnableCCVRampProxy.CallOpts)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OwnableCCVRampProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OwnableCCVRampProxy *OwnableCCVRampProxySession) TypeAndVersion() (string, error) {
	return _OwnableCCVRampProxy.Contract.TypeAndVersion(&_OwnableCCVRampProxy.CallOpts)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyCallerSession) TypeAndVersion() (string, error) {
	return _OwnableCCVRampProxy.Contract.TypeAndVersion(&_OwnableCCVRampProxy.CallOpts)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.contract.Transact(opts, "acceptOwnership")
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.AcceptOwnership(&_OwnableCCVRampProxy.TransactOpts)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.AcceptOwnership(&_OwnableCCVRampProxy.TransactOpts)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactor) SetRamps(opts *bind.TransactOpts, ramps []CCVRampProxySetRampsArgs) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.contract.Transact(opts, "setRamps", ramps)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxySession) SetRamps(ramps []CCVRampProxySetRampsArgs) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.SetRamps(&_OwnableCCVRampProxy.TransactOpts, ramps)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactorSession) SetRamps(ramps []CCVRampProxySetRampsArgs) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.SetRamps(&_OwnableCCVRampProxy.TransactOpts, ramps)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.TransferOwnership(&_OwnableCCVRampProxy.TransactOpts, to)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.TransferOwnership(&_OwnableCCVRampProxy.TransactOpts, to)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.contract.RawTransact(opts, calldata)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.Fallback(&_OwnableCCVRampProxy.TransactOpts, calldata)
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OwnableCCVRampProxy.Contract.Fallback(&_OwnableCCVRampProxy.TransactOpts, calldata)
}

type OwnableCCVRampProxyOwnershipTransferRequestedIterator struct {
	Event *OwnableCCVRampProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnableCCVRampProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableCCVRampProxyOwnershipTransferRequested)
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
		it.Event = new(OwnableCCVRampProxyOwnershipTransferRequested)
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

func (it *OwnableCCVRampProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OwnableCCVRampProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnableCCVRampProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableCCVRampProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnableCCVRampProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OwnableCCVRampProxyOwnershipTransferRequestedIterator{contract: _OwnableCCVRampProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OwnableCCVRampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnableCCVRampProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnableCCVRampProxyOwnershipTransferRequested)
				if err := _OwnableCCVRampProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*OwnableCCVRampProxyOwnershipTransferRequested, error) {
	event := new(OwnableCCVRampProxyOwnershipTransferRequested)
	if err := _OwnableCCVRampProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OwnableCCVRampProxyOwnershipTransferredIterator struct {
	Event *OwnableCCVRampProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnableCCVRampProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableCCVRampProxyOwnershipTransferred)
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
		it.Event = new(OwnableCCVRampProxyOwnershipTransferred)
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

func (it *OwnableCCVRampProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OwnableCCVRampProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnableCCVRampProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableCCVRampProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnableCCVRampProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OwnableCCVRampProxyOwnershipTransferredIterator{contract: _OwnableCCVRampProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableCCVRampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OwnableCCVRampProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnableCCVRampProxyOwnershipTransferred)
				if err := _OwnableCCVRampProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableCCVRampProxyOwnershipTransferred, error) {
	event := new(OwnableCCVRampProxyOwnershipTransferred)
	if err := _OwnableCCVRampProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OwnableCCVRampProxyRampSetIterator struct {
	Event *OwnableCCVRampProxyRampSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OwnableCCVRampProxyRampSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableCCVRampProxyRampSet)
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
		it.Event = new(OwnableCCVRampProxyRampSet)
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

func (it *OwnableCCVRampProxyRampSetIterator) Error() error {
	return it.fail
}

func (it *OwnableCCVRampProxyRampSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OwnableCCVRampProxyRampSet struct {
	RemoteChainSelector uint64
	Version             [32]byte
	RampAddress         common.Address
	Raw                 types.Log
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) FilterRampSet(opts *bind.FilterOpts, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (*OwnableCCVRampProxyRampSetIterator, error) {

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

	logs, sub, err := _OwnableCCVRampProxy.contract.FilterLogs(opts, "RampSet", remoteChainSelectorRule, versionRule, rampAddressRule)
	if err != nil {
		return nil, err
	}
	return &OwnableCCVRampProxyRampSetIterator{contract: _OwnableCCVRampProxy.contract, event: "RampSet", logs: logs, sub: sub}, nil
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) WatchRampSet(opts *bind.WatchOpts, sink chan<- *OwnableCCVRampProxyRampSet, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OwnableCCVRampProxy.contract.WatchLogs(opts, "RampSet", remoteChainSelectorRule, versionRule, rampAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OwnableCCVRampProxyRampSet)
				if err := _OwnableCCVRampProxy.contract.UnpackLog(event, "RampSet", log); err != nil {
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

func (_OwnableCCVRampProxy *OwnableCCVRampProxyFilterer) ParseRampSet(log types.Log) (*OwnableCCVRampProxyRampSet, error) {
	event := new(OwnableCCVRampProxyRampSet)
	if err := _OwnableCCVRampProxy.contract.UnpackLog(event, "RampSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (OwnableCCVRampProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OwnableCCVRampProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OwnableCCVRampProxyRampSet) Topic() common.Hash {
	return common.HexToHash("0x61079add8b3485b65ad33a15cebc6188cca5cb506a21df75784554c7339d3584")
}

func (_OwnableCCVRampProxy *OwnableCCVRampProxy) Address() common.Address {
	return _OwnableCCVRampProxy.address
}

type OwnableCCVRampProxyInterface interface {
	GetRamp(opts *bind.CallOpts, remoteChainSelector uint64, version [32]byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetRamps(opts *bind.TransactOpts, ramps []CCVRampProxySetRampsArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableCCVRampProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OwnableCCVRampProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OwnableCCVRampProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OwnableCCVRampProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableCCVRampProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OwnableCCVRampProxyOwnershipTransferred, error)

	FilterRampSet(opts *bind.FilterOpts, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (*OwnableCCVRampProxyRampSetIterator, error)

	WatchRampSet(opts *bind.WatchOpts, sink chan<- *OwnableCCVRampProxyRampSet, remoteChainSelector []uint64, version [][32]byte, rampAddress []common.Address) (event.Subscription, error)

	ParseRampSet(log types.Log) (*OwnableCCVRampProxyRampSet, error)

	Address() common.Address
}
