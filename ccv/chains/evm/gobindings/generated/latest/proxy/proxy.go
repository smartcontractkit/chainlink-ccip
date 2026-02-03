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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFeeAggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTarget\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFeeAggregator\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTarget\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeAggregatorUpdated\",\"inputs\":[{\"name\":\"oldFeeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newFeeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TargetUpdated\",\"inputs\":[{\"name\":\"oldTarget\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTarget\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60803461013757601f610bb138819003918201601f19168301916001600160401b0383118484101761013c57808492604094855283398101031261013757610052602061004b83610152565b9201610152565b90331561012657600180546001600160a01b031916331790556001600160a01b031690811561011557600280546001600160a01b03198116841790915560405192906001600160a01b03167f331faca2e54d546f21863baefddc0bc0b9fe7554216f8798e32574255571713f600080a3600380546001600160a01b039283166001600160a01b0319821681179092559091167f5f93cfaedcfeead9f6922f03a6557cc9c40dd65f320e80dd4aa68fce736bf723600080a3610a4a90816101678239f35b6342bcdf7f60e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101375756fe6080604052600436101561001a575b341561086f575b600080fd5b60003560e01c806315b358e0146100aa578063181f5a77146100a55780635cb80c5d146100a0578063776d1a011461009b57806379ba5097146100965780638da5cb5b146100915780639cb406c91461008c578063f00e6a2a146100875763f2fde38b0361000e57610627565b6105d5565b610583565b610531565b610448565b610357565b6102da565b610208565b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610015576004356100e581610159565b6100ed6108a8565b73ffffffffffffffffffffffffffffffffffffffff80600354921691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600355167f5f93cfaedcfeead9f6922f03a6557cc9c40dd65f320e80dd4aa68fce736bf723600080a3005b73ffffffffffffffffffffffffffffffffffffffff81160361001557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176101c257604052565b610177565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101c257604052565b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557604051610243816101a6565b600f81527f50726f787920312e372e302d6465760000000000000000000000000000000000602082015260405190602082528181519182602083015260005b8381106102c25750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001610282565b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760043567ffffffffffffffff8111610015573660238201121561001557806004013567ffffffffffffffff8111610015573660248260051b84010111610015576024610355920161071b565b005b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155773ffffffffffffffffffffffffffffffffffffffff6004356103a781610159565b6103af6108a8565b16801561041e5773ffffffffffffffffffffffffffffffffffffffff600254827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600255167f331faca2e54d546f21863baefddc0bc0b9fe7554216f8798e32574255571713f600080a3005b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155760005473ffffffffffffffffffffffffffffffffffffffff81163303610507577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b346100155760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261001557602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b346100155760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100155773ffffffffffffffffffffffffffffffffffffffff60043561067781610159565b61067f6108a8565b163381146106f157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff6003541691821561041e5760005b81811061074b5750505050565b61077a61076161075c8385876108f3565b610932565b73ffffffffffffffffffffffffffffffffffffffff1690565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa801561086a57600194889260009261083a575b50816107ee575b505050500161073e565b8161081e7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e938561082e9461095a565b6040519081529081906020820190565b0390a3388581806107e4565b61085c91925060203d8111610863575b61085481836101c7565b81019061093f565b90386107dd565b503d61084a565b61094e565b60008073ffffffffffffffffffffffffffffffffffffffff6002541636828037818036925af13d6000803e6108a3573d6000fd5b3d6000f35b73ffffffffffffffffffffffffffffffffffffffff6001541633036108c957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156109035760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3561093c81610159565b90565b90816020910312610015575190565b6040513d6000823e3d90fd5b916020916000916040519073ffffffffffffffffffffffffffffffffffffffff858301937fa9059cbb0000000000000000000000000000000000000000000000000000000085521660248301526044820152604481526109bb6064826101c7565b519082855af11561094e576000513d610a34575073ffffffffffffffffffffffffffffffffffffffff81163b155b6109f05750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b600114156109e956fea164736f6c634300081a000a",
}

var ProxyABI = ProxyMetaData.ABI

var ProxyBin = ProxyMetaData.Bin

func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, target common.Address, feeAggregator common.Address) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyBin), backend, target, feeAggregator)
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

func (_Proxy *ProxyCaller) GetFeeAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getFeeAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Proxy *ProxySession) GetFeeAggregator() (common.Address, error) {
	return _Proxy.Contract.GetFeeAggregator(&_Proxy.CallOpts)
}

func (_Proxy *ProxyCallerSession) GetFeeAggregator() (common.Address, error) {
	return _Proxy.Contract.GetFeeAggregator(&_Proxy.CallOpts)
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

func (_Proxy *ProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Proxy *ProxySession) TypeAndVersion() (string, error) {
	return _Proxy.Contract.TypeAndVersion(&_Proxy.CallOpts)
}

func (_Proxy *ProxyCallerSession) TypeAndVersion() (string, error) {
	return _Proxy.Contract.TypeAndVersion(&_Proxy.CallOpts)
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

func (_Proxy *ProxyTransactor) SetFeeAggregator(opts *bind.TransactOpts, feeAggregator common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "setFeeAggregator", feeAggregator)
}

func (_Proxy *ProxySession) SetFeeAggregator(feeAggregator common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetFeeAggregator(&_Proxy.TransactOpts, feeAggregator)
}

func (_Proxy *ProxyTransactorSession) SetFeeAggregator(feeAggregator common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetFeeAggregator(&_Proxy.TransactOpts, feeAggregator)
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

func (_Proxy *ProxyTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_Proxy *ProxySession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.WithdrawFeeTokens(&_Proxy.TransactOpts, feeTokens)
}

func (_Proxy *ProxyTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.WithdrawFeeTokens(&_Proxy.TransactOpts, feeTokens)
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

type ProxyFeeAggregatorUpdatedIterator struct {
	Event *ProxyFeeAggregatorUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ProxyFeeAggregatorUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyFeeAggregatorUpdated)
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
		it.Event = new(ProxyFeeAggregatorUpdated)
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

func (it *ProxyFeeAggregatorUpdatedIterator) Error() error {
	return it.fail
}

func (it *ProxyFeeAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ProxyFeeAggregatorUpdated struct {
	OldFeeAggregator common.Address
	NewFeeAggregator common.Address
	Raw              types.Log
}

func (_Proxy *ProxyFilterer) FilterFeeAggregatorUpdated(opts *bind.FilterOpts, oldFeeAggregator []common.Address, newFeeAggregator []common.Address) (*ProxyFeeAggregatorUpdatedIterator, error) {

	var oldFeeAggregatorRule []interface{}
	for _, oldFeeAggregatorItem := range oldFeeAggregator {
		oldFeeAggregatorRule = append(oldFeeAggregatorRule, oldFeeAggregatorItem)
	}
	var newFeeAggregatorRule []interface{}
	for _, newFeeAggregatorItem := range newFeeAggregator {
		newFeeAggregatorRule = append(newFeeAggregatorRule, newFeeAggregatorItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "FeeAggregatorUpdated", oldFeeAggregatorRule, newFeeAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &ProxyFeeAggregatorUpdatedIterator{contract: _Proxy.contract, event: "FeeAggregatorUpdated", logs: logs, sub: sub}, nil
}

func (_Proxy *ProxyFilterer) WatchFeeAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *ProxyFeeAggregatorUpdated, oldFeeAggregator []common.Address, newFeeAggregator []common.Address) (event.Subscription, error) {

	var oldFeeAggregatorRule []interface{}
	for _, oldFeeAggregatorItem := range oldFeeAggregator {
		oldFeeAggregatorRule = append(oldFeeAggregatorRule, oldFeeAggregatorItem)
	}
	var newFeeAggregatorRule []interface{}
	for _, newFeeAggregatorItem := range newFeeAggregator {
		newFeeAggregatorRule = append(newFeeAggregatorRule, newFeeAggregatorItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "FeeAggregatorUpdated", oldFeeAggregatorRule, newFeeAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ProxyFeeAggregatorUpdated)
				if err := _Proxy.contract.UnpackLog(event, "FeeAggregatorUpdated", log); err != nil {
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

func (_Proxy *ProxyFilterer) ParseFeeAggregatorUpdated(log types.Log) (*ProxyFeeAggregatorUpdated, error) {
	event := new(ProxyFeeAggregatorUpdated)
	if err := _Proxy.contract.UnpackLog(event, "FeeAggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ProxyFeeTokenWithdrawnIterator struct {
	Event *ProxyFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ProxyFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyFeeTokenWithdrawn)
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
		it.Event = new(ProxyFeeTokenWithdrawn)
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

func (it *ProxyFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *ProxyFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ProxyFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_Proxy *ProxyFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*ProxyFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &ProxyFeeTokenWithdrawnIterator{contract: _Proxy.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_Proxy *ProxyFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *ProxyFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ProxyFeeTokenWithdrawn)
				if err := _Proxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_Proxy *ProxyFilterer) ParseFeeTokenWithdrawn(log types.Log) (*ProxyFeeTokenWithdrawn, error) {
	event := new(ProxyFeeTokenWithdrawn)
	if err := _Proxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

func (ProxyFeeAggregatorUpdated) Topic() common.Hash {
	return common.HexToHash("0x5f93cfaedcfeead9f6922f03a6557cc9c40dd65f320e80dd4aa68fce736bf723")
}

func (ProxyFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
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
	GetFeeAggregator(opts *bind.CallOpts) (common.Address, error)

	GetTarget(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetFeeAggregator(opts *bind.TransactOpts, feeAggregator common.Address) (*types.Transaction, error)

	SetTarget(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterFeeAggregatorUpdated(opts *bind.FilterOpts, oldFeeAggregator []common.Address, newFeeAggregator []common.Address) (*ProxyFeeAggregatorUpdatedIterator, error)

	WatchFeeAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *ProxyFeeAggregatorUpdated, oldFeeAggregator []common.Address, newFeeAggregator []common.Address) (event.Subscription, error)

	ParseFeeAggregatorUpdated(log types.Log) (*ProxyFeeAggregatorUpdated, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*ProxyFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *ProxyFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*ProxyFeeTokenWithdrawn, error)

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
