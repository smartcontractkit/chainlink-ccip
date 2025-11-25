// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cctp_message_transmitter_proxy

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

type CCTPMessageTransmitterProxyAllowedCallerConfigArgs struct {
	Caller  common.Address
	Allowed bool
}

var CCTPMessageTransmitterProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureAllowedCallers\",\"inputs\":[{\"name\":\"configArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct CCTPMessageTransmitterProxy.AllowedCallerConfigArgs[]\",\"components\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_cctpTransmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IMessageTransmitter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AllowedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a0806040523461010f57602081610a30803803809161001f8285610114565b83398101031261010f57516001600160a01b0381169081900361010f5733156100fe57600180546001600160a01b03191633179055604051632c12192160e01b815290602090829060049082905afa9081156100f2576000916100a9575b506001600160a01b03166080526040516108e2908161014e823960805181818161013101526104690152f35b6020813d6020116100ea575b816100c260209383610114565b810103126100e65751906001600160a01b03821682036100e357503861007d565b80fd5b5080fd5b3d91506100b5565b6040513d6000823e3d90fd5b639b15e16f60e01b60005260046000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761013757604052565b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816310807aa7146105cf57508063181f5a771461050857806357ecfd28146103ac57806379ba5097146103245780638da5cb5b146102fd578063a6801258146102b5578063bd028e7c14610155578063cfc1db06146101115763f2fde38b1461008257600080fd5b3461010c57602036600319011261010c576004356001600160a01b03811680910361010c576100af61073a565b3381146100fb57806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b636d6c4ee560e11b60005260046000fd5b600080fd5b3461010c57600036600319011261010c5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461010c57602036600319011261010c5760043567ffffffffffffffff811161010c573660238201121561010c57806004013567ffffffffffffffff811161010c576024820191602436918360061b01011161010c576101b361073a565b60005b8181106101bf57005b60206101cc828486610700565b013590811515820361010c5760019115610250576102046001600160a01b036101fe6101f9848789610700565b610726565b16610875565b61020f575b016101b6565b6001600160a01b036102256101f9838688610700565b167f663c7e9ed36d9138863ef4306bbfcf01f60e1e7ca69b370c53d3094369e2cb02600080a2610209565b61026f6001600160a01b036102696101f9848789610700565b16610777565b15610209576001600160a01b0361028a6101f9838688610700565b167fbc0a6e072a312bde289d32bc84e5b758d7c617f734ecc0d69f995b2d7e69be36600080a2610209565b3461010c57602036600319011261010c576004356001600160a01b03811680910361010c576102f36020916000526003602052604060002054151590565b6040519015158152f35b3461010c57600036600319011261010c5760206001600160a01b0360015416604051908152f35b3461010c57600036600319011261010c576000546001600160a01b038116330361039b5760015490336001600160a01b03198316176001556001600160a01b0319166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b63015aa1e360e11b60005260046000fd5b3461010c57604036600319011261010c5760043567ffffffffffffffff811161010c576103dd9036906004016106b1565b60243567ffffffffffffffff811161010c576103fd9036906004016106b1565b929091610417336000526003602052604060002054151590565b156104f35761045b602093610449956040519687958695630afd9fa560e31b87526040600488015260448701916106df565b848103600319016024860152916106df565b038160006001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af19081156104e7576000916104a7575b6020826040519015158152f35b6020813d6020116104df575b816104c06020938361068f565b810103126104db575180151581036104db579050602061049a565b5080fd5b3d91506104b3565b6040513d6000823e3d90fd5b63472511eb60e11b6000523360045260246000fd5b3461010c57600036600319011261010c576040516060810181811067ffffffffffffffff8211176105b957604052602181527f434354504d6573736167655472616e736d697474657250726f787920312e362e6020820152601960f91b604082015260405190602082528181519182602083015260005b8381106105a15750508160006040809484010152601f80199101168101030190f35b6020828201810151604087840101528593500161057f565b634e487b7160e01b600052604160045260246000fd5b3461010c57600036600319011261010c576002549081815260208101809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b818110610679575050508161062d91038261068f565b6040519182916020830190602084525180915260408301919060005b818110610657575050500390f35b82516001600160a01b0316845285945060209384019390920191600101610649565b8254845260209093019260019283019201610617565b90601f8019910116810190811067ffffffffffffffff8211176105b957604052565b9181601f8401121561010c5782359167ffffffffffffffff831161010c576020838186019501011161010c57565b908060209392818452848401376000828201840152601f01601f1916010190565b91908110156107105760061b0190565b634e487b7160e01b600052603260045260246000fd5b356001600160a01b038116810361010c5790565b6001600160a01b0360015416330361074e57565b6315ae3a6f60e11b60005260046000fd5b80548210156107105760005260206000200190600090565b600081815260036020526040902054801561086e5760001981018181116108585760025460001981019190821161085857818103610807575b50505060025480156107f157600019016107cb81600261075f565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61084061081861082993600261075f565b90549060031b1c928392600261075f565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806107b0565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146108cf57600254680100000000000000008110156105b9576108b6610829826001859401600255600261075f565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
}

var CCTPMessageTransmitterProxyABI = CCTPMessageTransmitterProxyMetaData.ABI

var CCTPMessageTransmitterProxyBin = CCTPMessageTransmitterProxyMetaData.Bin

func DeployCCTPMessageTransmitterProxy(auth *bind.TransactOpts, backend bind.ContractBackend, tokenMessenger common.Address) (common.Address, *types.Transaction, *CCTPMessageTransmitterProxy, error) {
	parsed, err := CCTPMessageTransmitterProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPMessageTransmitterProxyBin), backend, tokenMessenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCTPMessageTransmitterProxy{address: address, abi: *parsed, CCTPMessageTransmitterProxyCaller: CCTPMessageTransmitterProxyCaller{contract: contract}, CCTPMessageTransmitterProxyTransactor: CCTPMessageTransmitterProxyTransactor{contract: contract}, CCTPMessageTransmitterProxyFilterer: CCTPMessageTransmitterProxyFilterer{contract: contract}}, nil
}

type CCTPMessageTransmitterProxy struct {
	address common.Address
	abi     abi.ABI
	CCTPMessageTransmitterProxyCaller
	CCTPMessageTransmitterProxyTransactor
	CCTPMessageTransmitterProxyFilterer
}

type CCTPMessageTransmitterProxyCaller struct {
	contract *bind.BoundContract
}

type CCTPMessageTransmitterProxyTransactor struct {
	contract *bind.BoundContract
}

type CCTPMessageTransmitterProxyFilterer struct {
	contract *bind.BoundContract
}

type CCTPMessageTransmitterProxySession struct {
	Contract     *CCTPMessageTransmitterProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCTPMessageTransmitterProxyCallerSession struct {
	Contract *CCTPMessageTransmitterProxyCaller
	CallOpts bind.CallOpts
}

type CCTPMessageTransmitterProxyTransactorSession struct {
	Contract     *CCTPMessageTransmitterProxyTransactor
	TransactOpts bind.TransactOpts
}

type CCTPMessageTransmitterProxyRaw struct {
	Contract *CCTPMessageTransmitterProxy
}

type CCTPMessageTransmitterProxyCallerRaw struct {
	Contract *CCTPMessageTransmitterProxyCaller
}

type CCTPMessageTransmitterProxyTransactorRaw struct {
	Contract *CCTPMessageTransmitterProxyTransactor
}

func NewCCTPMessageTransmitterProxy(address common.Address, backend bind.ContractBackend) (*CCTPMessageTransmitterProxy, error) {
	abi, err := abi.JSON(strings.NewReader(CCTPMessageTransmitterProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCTPMessageTransmitterProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxy{address: address, abi: abi, CCTPMessageTransmitterProxyCaller: CCTPMessageTransmitterProxyCaller{contract: contract}, CCTPMessageTransmitterProxyTransactor: CCTPMessageTransmitterProxyTransactor{contract: contract}, CCTPMessageTransmitterProxyFilterer: CCTPMessageTransmitterProxyFilterer{contract: contract}}, nil
}

func NewCCTPMessageTransmitterProxyCaller(address common.Address, caller bind.ContractCaller) (*CCTPMessageTransmitterProxyCaller, error) {
	contract, err := bindCCTPMessageTransmitterProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyCaller{contract: contract}, nil
}

func NewCCTPMessageTransmitterProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CCTPMessageTransmitterProxyTransactor, error) {
	contract, err := bindCCTPMessageTransmitterProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyTransactor{contract: contract}, nil
}

func NewCCTPMessageTransmitterProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CCTPMessageTransmitterProxyFilterer, error) {
	contract, err := bindCCTPMessageTransmitterProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyFilterer{contract: contract}, nil
}

func bindCCTPMessageTransmitterProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCTPMessageTransmitterProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPMessageTransmitterProxy.Contract.CCTPMessageTransmitterProxyCaller.contract.Call(opts, result, method, params...)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.CCTPMessageTransmitterProxyTransactor.contract.Transfer(opts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.CCTPMessageTransmitterProxyTransactor.contract.Transact(opts, method, params...)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPMessageTransmitterProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.contract.Transfer(opts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.contract.Transact(opts, method, params...)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCaller) GetAllowedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPMessageTransmitterProxy.contract.Call(opts, &out, "getAllowedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) GetAllowedCallers() ([]common.Address, error) {
	return _CCTPMessageTransmitterProxy.Contract.GetAllowedCallers(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCallerSession) GetAllowedCallers() ([]common.Address, error) {
	return _CCTPMessageTransmitterProxy.Contract.GetAllowedCallers(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCaller) ICctpTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPMessageTransmitterProxy.contract.Call(opts, &out, "i_cctpTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) ICctpTransmitter() (common.Address, error) {
	return _CCTPMessageTransmitterProxy.Contract.ICctpTransmitter(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCallerSession) ICctpTransmitter() (common.Address, error) {
	return _CCTPMessageTransmitterProxy.Contract.ICctpTransmitter(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCaller) IsAllowedCaller(opts *bind.CallOpts, caller common.Address) (bool, error) {
	var out []interface{}
	err := _CCTPMessageTransmitterProxy.contract.Call(opts, &out, "isAllowedCaller", caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) IsAllowedCaller(caller common.Address) (bool, error) {
	return _CCTPMessageTransmitterProxy.Contract.IsAllowedCaller(&_CCTPMessageTransmitterProxy.CallOpts, caller)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCallerSession) IsAllowedCaller(caller common.Address) (bool, error) {
	return _CCTPMessageTransmitterProxy.Contract.IsAllowedCaller(&_CCTPMessageTransmitterProxy.CallOpts, caller)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPMessageTransmitterProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) Owner() (common.Address, error) {
	return _CCTPMessageTransmitterProxy.Contract.Owner(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCallerSession) Owner() (common.Address, error) {
	return _CCTPMessageTransmitterProxy.Contract.Owner(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCTPMessageTransmitterProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) TypeAndVersion() (string, error) {
	return _CCTPMessageTransmitterProxy.Contract.TypeAndVersion(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCallerSession) TypeAndVersion() (string, error) {
	return _CCTPMessageTransmitterProxy.Contract.TypeAndVersion(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.contract.Transact(opts, "acceptOwnership")
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.AcceptOwnership(&_CCTPMessageTransmitterProxy.TransactOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.AcceptOwnership(&_CCTPMessageTransmitterProxy.TransactOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactor) ConfigureAllowedCallers(opts *bind.TransactOpts, configArgs []CCTPMessageTransmitterProxyAllowedCallerConfigArgs) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.contract.Transact(opts, "configureAllowedCallers", configArgs)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) ConfigureAllowedCallers(configArgs []CCTPMessageTransmitterProxyAllowedCallerConfigArgs) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.ConfigureAllowedCallers(&_CCTPMessageTransmitterProxy.TransactOpts, configArgs)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactorSession) ConfigureAllowedCallers(configArgs []CCTPMessageTransmitterProxyAllowedCallerConfigArgs) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.ConfigureAllowedCallers(&_CCTPMessageTransmitterProxy.TransactOpts, configArgs)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.contract.Transact(opts, "receiveMessage", message, attestation)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.ReceiveMessage(&_CCTPMessageTransmitterProxy.TransactOpts, message, attestation)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactorSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.ReceiveMessage(&_CCTPMessageTransmitterProxy.TransactOpts, message, attestation)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.TransferOwnership(&_CCTPMessageTransmitterProxy.TransactOpts, to)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.TransferOwnership(&_CCTPMessageTransmitterProxy.TransactOpts, to)
}

type CCTPMessageTransmitterProxyAllowedCallerAddedIterator struct {
	Event *CCTPMessageTransmitterProxyAllowedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPMessageTransmitterProxyAllowedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPMessageTransmitterProxyAllowedCallerAdded)
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
		it.Event = new(CCTPMessageTransmitterProxyAllowedCallerAdded)
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

func (it *CCTPMessageTransmitterProxyAllowedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPMessageTransmitterProxyAllowedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPMessageTransmitterProxyAllowedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) FilterAllowedCallerAdded(opts *bind.FilterOpts, caller []common.Address) (*CCTPMessageTransmitterProxyAllowedCallerAddedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.FilterLogs(opts, "AllowedCallerAdded", callerRule)
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyAllowedCallerAddedIterator{contract: _CCTPMessageTransmitterProxy.contract, event: "AllowedCallerAdded", logs: logs, sub: sub}, nil
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) WatchAllowedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyAllowedCallerAdded, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.WatchLogs(opts, "AllowedCallerAdded", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPMessageTransmitterProxyAllowedCallerAdded)
				if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "AllowedCallerAdded", log); err != nil {
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

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) ParseAllowedCallerAdded(log types.Log) (*CCTPMessageTransmitterProxyAllowedCallerAdded, error) {
	event := new(CCTPMessageTransmitterProxyAllowedCallerAdded)
	if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "AllowedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPMessageTransmitterProxyAllowedCallerRemovedIterator struct {
	Event *CCTPMessageTransmitterProxyAllowedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPMessageTransmitterProxyAllowedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPMessageTransmitterProxyAllowedCallerRemoved)
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
		it.Event = new(CCTPMessageTransmitterProxyAllowedCallerRemoved)
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

func (it *CCTPMessageTransmitterProxyAllowedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPMessageTransmitterProxyAllowedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPMessageTransmitterProxyAllowedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) FilterAllowedCallerRemoved(opts *bind.FilterOpts, caller []common.Address) (*CCTPMessageTransmitterProxyAllowedCallerRemovedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.FilterLogs(opts, "AllowedCallerRemoved", callerRule)
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyAllowedCallerRemovedIterator{contract: _CCTPMessageTransmitterProxy.contract, event: "AllowedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) WatchAllowedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyAllowedCallerRemoved, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.WatchLogs(opts, "AllowedCallerRemoved", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPMessageTransmitterProxyAllowedCallerRemoved)
				if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "AllowedCallerRemoved", log); err != nil {
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

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) ParseAllowedCallerRemoved(log types.Log) (*CCTPMessageTransmitterProxyAllowedCallerRemoved, error) {
	event := new(CCTPMessageTransmitterProxyAllowedCallerRemoved)
	if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "AllowedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPMessageTransmitterProxyOwnershipTransferRequestedIterator struct {
	Event *CCTPMessageTransmitterProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPMessageTransmitterProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPMessageTransmitterProxyOwnershipTransferRequested)
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
		it.Event = new(CCTPMessageTransmitterProxyOwnershipTransferRequested)
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

func (it *CCTPMessageTransmitterProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCTPMessageTransmitterProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPMessageTransmitterProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPMessageTransmitterProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyOwnershipTransferRequestedIterator{contract: _CCTPMessageTransmitterProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPMessageTransmitterProxyOwnershipTransferRequested)
				if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCTPMessageTransmitterProxyOwnershipTransferRequested, error) {
	event := new(CCTPMessageTransmitterProxyOwnershipTransferRequested)
	if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPMessageTransmitterProxyOwnershipTransferredIterator struct {
	Event *CCTPMessageTransmitterProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPMessageTransmitterProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPMessageTransmitterProxyOwnershipTransferred)
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
		it.Event = new(CCTPMessageTransmitterProxyOwnershipTransferred)
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

func (it *CCTPMessageTransmitterProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCTPMessageTransmitterProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPMessageTransmitterProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPMessageTransmitterProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyOwnershipTransferredIterator{contract: _CCTPMessageTransmitterProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPMessageTransmitterProxyOwnershipTransferred)
				if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) ParseOwnershipTransferred(log types.Log) (*CCTPMessageTransmitterProxyOwnershipTransferred, error) {
	event := new(CCTPMessageTransmitterProxyOwnershipTransferred)
	if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (CCTPMessageTransmitterProxyAllowedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0x663c7e9ed36d9138863ef4306bbfcf01f60e1e7ca69b370c53d3094369e2cb02")
}

func (CCTPMessageTransmitterProxyAllowedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xbc0a6e072a312bde289d32bc84e5b758d7c617f734ecc0d69f995b2d7e69be36")
}

func (CCTPMessageTransmitterProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCTPMessageTransmitterProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxy) Address() common.Address {
	return _CCTPMessageTransmitterProxy.address
}

type CCTPMessageTransmitterProxyInterface interface {
	GetAllowedCallers(opts *bind.CallOpts) ([]common.Address, error)

	ICctpTransmitter(opts *bind.CallOpts) (common.Address, error)

	IsAllowedCaller(opts *bind.CallOpts, caller common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ConfigureAllowedCallers(opts *bind.TransactOpts, configArgs []CCTPMessageTransmitterProxyAllowedCallerConfigArgs) (*types.Transaction, error)

	ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowedCallerAdded(opts *bind.FilterOpts, caller []common.Address) (*CCTPMessageTransmitterProxyAllowedCallerAddedIterator, error)

	WatchAllowedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyAllowedCallerAdded, caller []common.Address) (event.Subscription, error)

	ParseAllowedCallerAdded(log types.Log) (*CCTPMessageTransmitterProxyAllowedCallerAdded, error)

	FilterAllowedCallerRemoved(opts *bind.FilterOpts, caller []common.Address) (*CCTPMessageTransmitterProxyAllowedCallerRemovedIterator, error)

	WatchAllowedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyAllowedCallerRemoved, caller []common.Address) (event.Subscription, error)

	ParseAllowedCallerRemoved(log types.Log) (*CCTPMessageTransmitterProxyAllowedCallerRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPMessageTransmitterProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPMessageTransmitterProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPMessageTransmitterProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPMessageTransmitterProxyOwnershipTransferred, error)

	Address() common.Address
}
