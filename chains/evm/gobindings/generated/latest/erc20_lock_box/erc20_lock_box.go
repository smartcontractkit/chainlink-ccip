// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_lock_box

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

type ERC20LockBoxAllowedCallerConfigArgs struct {
	Token   common.Address
	Caller  common.Address
	Allowed bool
}

var ERC20LockBoxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureAllowedCallers\",\"inputs\":[{\"name\":\"configArgs\",\"type\":\"tuple[]\",\"internalType\":\"structERC20LockBox.AllowedCallerConfigArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_tokenAdminRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTokenAdminRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedCaller\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedCallerAdded\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowedCallerRemoved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecipientCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAdminRegistryCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a03460b457601f61119138819003918201601f19168301916001600160401b0383118484101760b95780849260209460405283398101031260b457516001600160a01b0381169081900360b457331560a357600180546001600160a01b0319163317905580156092576080526040516110c190816100d0823960805181818161033a015281816107c60152610c6b0152f35b635c20484160e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063162d2edd1461075057806317b1f24214610636578063233671d51461044757806379ba50971461035e5780638ca86f28146102ef5780638da5cb5b1461029d5780638e932c0e1461020d578063f2fde38b146100f25763f901fa161461007f57600080fd5b346100ed5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed576100b6610ade565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036100ed576020916100e391610c0a565b6040519015158152f35b600080fd5b346100ed5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed57610129610ade565b73ffffffffffffffffffffffffffffffffffffffff60015416908133036101e35773ffffffffffffffffffffffffffffffffffffffff16903382146101b957817fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000557fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100ed5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed57610244610ade565b6024359067ffffffffffffffff821682036100ed5773ffffffffffffffffffffffffffffffffffffffff16600052600360205267ffffffffffffffff604060002091166000526020526020604060002054604051908152f35b346100ed5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100ed5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346100ed5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed5760005473ffffffffffffffffffffffffffffffffffffffff8116330361041d577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100ed5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed5761047e610ade565b602435906044359173ffffffffffffffffffffffffffffffffffffffff83168093036100ed576064359167ffffffffffffffff83168093036100ed576104c48282610dc7565b831561060c5773ffffffffffffffffffffffffffffffffffffffff16908160005260036020526040600020836000526020526040600020548181106105d757508160005260036020526040600020836000526020526040600020908154918183039283116105a8577fc6de56eb9f3f126f4b7f2e63a8477225c96fe39e4b742116b8d81f656820c0529360209361059f9255604051907fa9059cbb00000000000000000000000000000000000000000000000000000000858301528760248301528360448301526044825261059a606483610b61565b610e24565b604051908152a3005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90837fd236ce5e0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7fd87070520000000000000000000000000000000000000000000000000000000060005260046000fd5b346100ed5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed5761066d610ade565b602435906044359067ffffffffffffffff82168092036100ed57806106a78473ffffffffffffffffffffffffffffffffffffffff93610dc7565b166106fb6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152846064820152606481526106f5608482610b61565b82610e24565b600052600360205260406000208160005260205260406000208054908382018092116105a857556040519182527f88ab94ac53260736800da5d05843e504231e9d57ea5cc4ce6479495a52fa296d60203393a3005b346100ed5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ed5760043567ffffffffffffffff81116100ed57366023820112156100ed57806004013567ffffffffffffffff81116100ed57602482019160243691606084020101116100ed577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1660005b82811061080757005b73ffffffffffffffffffffffffffffffffffffffff61082f61082a838688610b01565b610b40565b16908115610ab4576040517fc1af6e03000000000000000000000000000000000000000000000000000000008152826004820152336024820152602081604481875afa908115610aa857600091610a7a575b5015610a4c5761089d6020610897838789610b01565b01610b40565b9160406108ab838789610b01565b01359283151584036100ed57600193156109885781600052600260205260ff60408060002060009073ffffffffffffffffffffffffffffffffffffffff8516825260205220541615610901575b50505b016107fe565b73ffffffffffffffffffffffffffffffffffffffff90826000526002602052604080600020600090848416825260205220857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905516907f79404d3b2f81567a5c92b6bc47602aa54236fdaa21f2ec2d127f6a34d96aa84f600080a385806108f8565b81600052600260205260ff60408060002060009073ffffffffffffffffffffffffffffffffffffffff851682526020522054166109c7575b50506108fb565b73ffffffffffffffffffffffffffffffffffffffff908260005260026020526040806000206000908484168252602052207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905516907f5c12668e21792ebf115dda4657f791a5e7717d06c01d91cc7ae97b75e92d02de600080a385806109c0565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b610a9b915060203d8111610aa1575b610a938183610b61565b810190610bd1565b86610881565b503d610a89565b6040513d6000823e3d90fd5b7f802c78a20000000000000000000000000000000000000000000000000000000060005260046000fd5b6004359073ffffffffffffffffffffffffffffffffffffffff821682036100ed57565b9190811015610b11576060020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036100ed5790565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610ba257604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b908160209103126100ed575180151581036100ed5790565b519073ffffffffffffffffffffffffffffffffffffffff821682036100ed57565b9073ffffffffffffffffffffffffffffffffffffffff604051927fcb67e3b10000000000000000000000000000000000000000000000000000000084521680600484015260608360248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa928315610aa857600093610d16575b5073ffffffffffffffffffffffffffffffffffffffff806040850151169216918214928315610cf5575b508215610cd257505090565b909150600052600260205260406000209060005260205260ff6040600020541690565b5173ffffffffffffffffffffffffffffffffffffffff168214925038610cc6565b6060939193813d606011610dbf575b81610d3260609383610b61565b81010312610dbb5760405191606083019083821067ffffffffffffffff831117610d8e5750610d82916040918252610d6981610be9565b8452610d7760208201610be9565b602085015201610be9565b60408201529138610c9c565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b5080fd5b3d9150610d25565b73ffffffffffffffffffffffffffffffffffffffff811615610ab457610dee903390610c0a565b15610a4c5715610dfa57565b7f8b1fa9dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff16604091600080845192610e4d8685610b61565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260208151910182865af13d15610f8a573d9067ffffffffffffffff8211610ba257610ee19360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160192610ed287519485610b61565b83523d6000602085013e610f93565b80519081610eee57505050565b602080610eff938301019101610bd1565b15610f075750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91610ee1926060915b9192901561100e5750815115610fa7575090565b3b15610fb05790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156110215750805190602001fd5b604051907f08c379a0000000000000000000000000000000000000000000000000000000008252602060048301528181519182602483015260005b83811061109c5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604480968601015201168101030190fd5b6020828201810151604487840101528593500161105c56fea164736f6c634300081a000a",
}

var ERC20LockBoxABI = ERC20LockBoxMetaData.ABI

var ERC20LockBoxBin = ERC20LockBoxMetaData.Bin

func DeployERC20LockBox(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAdminRegistry common.Address) (common.Address, *types.Transaction, *ERC20LockBox, error) {
	parsed, err := ERC20LockBoxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20LockBoxBin), backend, tokenAdminRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20LockBox{address: address, abi: *parsed, ERC20LockBoxCaller: ERC20LockBoxCaller{contract: contract}, ERC20LockBoxTransactor: ERC20LockBoxTransactor{contract: contract}, ERC20LockBoxFilterer: ERC20LockBoxFilterer{contract: contract}}, nil
}

type ERC20LockBox struct {
	address common.Address
	abi     abi.ABI
	ERC20LockBoxCaller
	ERC20LockBoxTransactor
	ERC20LockBoxFilterer
}

type ERC20LockBoxCaller struct {
	contract *bind.BoundContract
}

type ERC20LockBoxTransactor struct {
	contract *bind.BoundContract
}

type ERC20LockBoxFilterer struct {
	contract *bind.BoundContract
}

type ERC20LockBoxSession struct {
	Contract     *ERC20LockBox
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ERC20LockBoxCallerSession struct {
	Contract *ERC20LockBoxCaller
	CallOpts bind.CallOpts
}

type ERC20LockBoxTransactorSession struct {
	Contract     *ERC20LockBoxTransactor
	TransactOpts bind.TransactOpts
}

type ERC20LockBoxRaw struct {
	Contract *ERC20LockBox
}

type ERC20LockBoxCallerRaw struct {
	Contract *ERC20LockBoxCaller
}

type ERC20LockBoxTransactorRaw struct {
	Contract *ERC20LockBoxTransactor
}

func NewERC20LockBox(address common.Address, backend bind.ContractBackend) (*ERC20LockBox, error) {
	abi, err := abi.JSON(strings.NewReader(ERC20LockBoxABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindERC20LockBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBox{address: address, abi: abi, ERC20LockBoxCaller: ERC20LockBoxCaller{contract: contract}, ERC20LockBoxTransactor: ERC20LockBoxTransactor{contract: contract}, ERC20LockBoxFilterer: ERC20LockBoxFilterer{contract: contract}}, nil
}

func NewERC20LockBoxCaller(address common.Address, caller bind.ContractCaller) (*ERC20LockBoxCaller, error) {
	contract, err := bindERC20LockBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxCaller{contract: contract}, nil
}

func NewERC20LockBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20LockBoxTransactor, error) {
	contract, err := bindERC20LockBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxTransactor{contract: contract}, nil
}

func NewERC20LockBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20LockBoxFilterer, error) {
	contract, err := bindERC20LockBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxFilterer{contract: contract}, nil
}

func bindERC20LockBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20LockBoxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ERC20LockBox *ERC20LockBoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20LockBox.Contract.ERC20LockBoxCaller.contract.Call(opts, result, method, params...)
}

func (_ERC20LockBox *ERC20LockBoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.ERC20LockBoxTransactor.contract.Transfer(opts)
}

func (_ERC20LockBox *ERC20LockBoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.ERC20LockBoxTransactor.contract.Transact(opts, method, params...)
}

func (_ERC20LockBox *ERC20LockBoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20LockBox.Contract.contract.Call(opts, result, method, params...)
}

func (_ERC20LockBox *ERC20LockBoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.contract.Transfer(opts)
}

func (_ERC20LockBox *ERC20LockBoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.contract.Transact(opts, method, params...)
}

func (_ERC20LockBox *ERC20LockBoxCaller) GetBalance(opts *bind.CallOpts, token common.Address, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "getBalance", token, remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) GetBalance(token common.Address, remoteChainSelector uint64) (*big.Int, error) {
	return _ERC20LockBox.Contract.GetBalance(&_ERC20LockBox.CallOpts, token, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) GetBalance(token common.Address, remoteChainSelector uint64) (*big.Int, error) {
	return _ERC20LockBox.Contract.GetBalance(&_ERC20LockBox.CallOpts, token, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxCaller) ITokenAdminRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "i_tokenAdminRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) ITokenAdminRegistry() (common.Address, error) {
	return _ERC20LockBox.Contract.ITokenAdminRegistry(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) ITokenAdminRegistry() (common.Address, error) {
	return _ERC20LockBox.Contract.ITokenAdminRegistry(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCaller) IsAllowedCaller(opts *bind.CallOpts, token common.Address, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "isAllowedCaller", token, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) IsAllowedCaller(token common.Address, caller common.Address) (bool, error) {
	return _ERC20LockBox.Contract.IsAllowedCaller(&_ERC20LockBox.CallOpts, token, caller)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) IsAllowedCaller(token common.Address, caller common.Address) (bool, error) {
	return _ERC20LockBox.Contract.IsAllowedCaller(&_ERC20LockBox.CallOpts, token, caller)
}

func (_ERC20LockBox *ERC20LockBoxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) Owner() (common.Address, error) {
	return _ERC20LockBox.Contract.Owner(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) Owner() (common.Address, error) {
	return _ERC20LockBox.Contract.Owner(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "acceptOwnership")
}

func (_ERC20LockBox *ERC20LockBoxSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC20LockBox.Contract.AcceptOwnership(&_ERC20LockBox.TransactOpts)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC20LockBox.Contract.AcceptOwnership(&_ERC20LockBox.TransactOpts)
}

func (_ERC20LockBox *ERC20LockBoxTransactor) ConfigureAllowedCallers(opts *bind.TransactOpts, configArgs []ERC20LockBoxAllowedCallerConfigArgs) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "configureAllowedCallers", configArgs)
}

func (_ERC20LockBox *ERC20LockBoxSession) ConfigureAllowedCallers(configArgs []ERC20LockBoxAllowedCallerConfigArgs) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.ConfigureAllowedCallers(&_ERC20LockBox.TransactOpts, configArgs)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) ConfigureAllowedCallers(configArgs []ERC20LockBoxAllowedCallerConfigArgs) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.ConfigureAllowedCallers(&_ERC20LockBox.TransactOpts, configArgs)
}

func (_ERC20LockBox *ERC20LockBoxTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "deposit", token, amount, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxSession) Deposit(token common.Address, amount *big.Int, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Deposit(&_ERC20LockBox.TransactOpts, token, amount, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) Deposit(token common.Address, amount *big.Int, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Deposit(&_ERC20LockBox.TransactOpts, token, amount, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "transferOwnership", to)
}

func (_ERC20LockBox *ERC20LockBoxSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.TransferOwnership(&_ERC20LockBox.TransactOpts, to)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.TransferOwnership(&_ERC20LockBox.TransactOpts, to)
}

func (_ERC20LockBox *ERC20LockBoxTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "withdraw", token, amount, recipient, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxSession) Withdraw(token common.Address, amount *big.Int, recipient common.Address, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Withdraw(&_ERC20LockBox.TransactOpts, token, amount, recipient, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) Withdraw(token common.Address, amount *big.Int, recipient common.Address, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Withdraw(&_ERC20LockBox.TransactOpts, token, amount, recipient, remoteChainSelector)
}

type ERC20LockBoxAllowedCallerAddedIterator struct {
	Event *ERC20LockBoxAllowedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxAllowedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxAllowedCallerAdded)
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
		it.Event = new(ERC20LockBoxAllowedCallerAdded)
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

func (it *ERC20LockBoxAllowedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxAllowedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxAllowedCallerAdded struct {
	Token  common.Address
	Caller common.Address
	Raw    types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterAllowedCallerAdded(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*ERC20LockBoxAllowedCallerAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "AllowedCallerAdded", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxAllowedCallerAddedIterator{contract: _ERC20LockBox.contract, event: "AllowedCallerAdded", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchAllowedCallerAdded(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerAdded, token []common.Address, caller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "AllowedCallerAdded", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxAllowedCallerAdded)
				if err := _ERC20LockBox.contract.UnpackLog(event, "AllowedCallerAdded", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseAllowedCallerAdded(log types.Log) (*ERC20LockBoxAllowedCallerAdded, error) {
	event := new(ERC20LockBoxAllowedCallerAdded)
	if err := _ERC20LockBox.contract.UnpackLog(event, "AllowedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ERC20LockBoxAllowedCallerRemovedIterator struct {
	Event *ERC20LockBoxAllowedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxAllowedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxAllowedCallerRemoved)
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
		it.Event = new(ERC20LockBoxAllowedCallerRemoved)
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

func (it *ERC20LockBoxAllowedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxAllowedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxAllowedCallerRemoved struct {
	Token  common.Address
	Caller common.Address
	Raw    types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterAllowedCallerRemoved(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*ERC20LockBoxAllowedCallerRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "AllowedCallerRemoved", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxAllowedCallerRemovedIterator{contract: _ERC20LockBox.contract, event: "AllowedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchAllowedCallerRemoved(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerRemoved, token []common.Address, caller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "AllowedCallerRemoved", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxAllowedCallerRemoved)
				if err := _ERC20LockBox.contract.UnpackLog(event, "AllowedCallerRemoved", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseAllowedCallerRemoved(log types.Log) (*ERC20LockBoxAllowedCallerRemoved, error) {
	event := new(ERC20LockBoxAllowedCallerRemoved)
	if err := _ERC20LockBox.contract.UnpackLog(event, "AllowedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ERC20LockBoxDepositIterator struct {
	Event *ERC20LockBoxDeposit

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxDepositIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxDeposit)
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
		it.Event = new(ERC20LockBoxDeposit)
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

func (it *ERC20LockBoxDepositIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxDeposit struct {
	RemoteChainSelector uint64
	Depositor           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterDeposit(opts *bind.FilterOpts, remoteChainSelector []uint64, depositor []common.Address) (*ERC20LockBoxDepositIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "Deposit", remoteChainSelectorRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxDepositIterator{contract: _ERC20LockBox.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxDeposit, remoteChainSelector []uint64, depositor []common.Address) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "Deposit", remoteChainSelectorRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxDeposit)
				if err := _ERC20LockBox.contract.UnpackLog(event, "Deposit", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseDeposit(log types.Log) (*ERC20LockBoxDeposit, error) {
	event := new(ERC20LockBoxDeposit)
	if err := _ERC20LockBox.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ERC20LockBoxOwnershipTransferRequestedIterator struct {
	Event *ERC20LockBoxOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxOwnershipTransferRequested)
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
		it.Event = new(ERC20LockBoxOwnershipTransferRequested)
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

func (it *ERC20LockBoxOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20LockBoxOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxOwnershipTransferRequestedIterator{contract: _ERC20LockBox.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxOwnershipTransferRequested)
				if err := _ERC20LockBox.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseOwnershipTransferRequested(log types.Log) (*ERC20LockBoxOwnershipTransferRequested, error) {
	event := new(ERC20LockBoxOwnershipTransferRequested)
	if err := _ERC20LockBox.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ERC20LockBoxOwnershipTransferredIterator struct {
	Event *ERC20LockBoxOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxOwnershipTransferred)
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
		it.Event = new(ERC20LockBoxOwnershipTransferred)
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

func (it *ERC20LockBoxOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20LockBoxOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxOwnershipTransferredIterator{contract: _ERC20LockBox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxOwnershipTransferred)
				if err := _ERC20LockBox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20LockBoxOwnershipTransferred, error) {
	event := new(ERC20LockBoxOwnershipTransferred)
	if err := _ERC20LockBox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ERC20LockBoxWithdrawalIterator struct {
	Event *ERC20LockBoxWithdrawal

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxWithdrawalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxWithdrawal)
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
		it.Event = new(ERC20LockBoxWithdrawal)
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

func (it *ERC20LockBoxWithdrawalIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxWithdrawal struct {
	RemoteChainSelector uint64
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterWithdrawal(opts *bind.FilterOpts, remoteChainSelector []uint64, recipient []common.Address) (*ERC20LockBoxWithdrawalIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "Withdrawal", remoteChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxWithdrawalIterator{contract: _ERC20LockBox.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxWithdrawal, remoteChainSelector []uint64, recipient []common.Address) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "Withdrawal", remoteChainSelectorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxWithdrawal)
				if err := _ERC20LockBox.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseWithdrawal(log types.Log) (*ERC20LockBoxWithdrawal, error) {
	event := new(ERC20LockBoxWithdrawal)
	if err := _ERC20LockBox.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ERC20LockBox *ERC20LockBox) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ERC20LockBox.abi.Events["AllowedCallerAdded"].ID:
		return _ERC20LockBox.ParseAllowedCallerAdded(log)
	case _ERC20LockBox.abi.Events["AllowedCallerRemoved"].ID:
		return _ERC20LockBox.ParseAllowedCallerRemoved(log)
	case _ERC20LockBox.abi.Events["Deposit"].ID:
		return _ERC20LockBox.ParseDeposit(log)
	case _ERC20LockBox.abi.Events["OwnershipTransferRequested"].ID:
		return _ERC20LockBox.ParseOwnershipTransferRequested(log)
	case _ERC20LockBox.abi.Events["OwnershipTransferred"].ID:
		return _ERC20LockBox.ParseOwnershipTransferred(log)
	case _ERC20LockBox.abi.Events["Withdrawal"].ID:
		return _ERC20LockBox.ParseWithdrawal(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ERC20LockBoxAllowedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0x79404d3b2f81567a5c92b6bc47602aa54236fdaa21f2ec2d127f6a34d96aa84f")
}

func (ERC20LockBoxAllowedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0x5c12668e21792ebf115dda4657f791a5e7717d06c01d91cc7ae97b75e92d02de")
}

func (ERC20LockBoxDeposit) Topic() common.Hash {
	return common.HexToHash("0x88ab94ac53260736800da5d05843e504231e9d57ea5cc4ce6479495a52fa296d")
}

func (ERC20LockBoxOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ERC20LockBoxOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ERC20LockBoxWithdrawal) Topic() common.Hash {
	return common.HexToHash("0xc6de56eb9f3f126f4b7f2e63a8477225c96fe39e4b742116b8d81f656820c052")
}

func (_ERC20LockBox *ERC20LockBox) Address() common.Address {
	return _ERC20LockBox.address
}

type ERC20LockBoxInterface interface {
	GetBalance(opts *bind.CallOpts, token common.Address, remoteChainSelector uint64) (*big.Int, error)

	ITokenAdminRegistry(opts *bind.CallOpts) (common.Address, error)

	IsAllowedCaller(opts *bind.CallOpts, token common.Address, caller common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ConfigureAllowedCallers(opts *bind.TransactOpts, configArgs []ERC20LockBoxAllowedCallerConfigArgs) (*types.Transaction, error)

	Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int, remoteChainSelector uint64) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address, remoteChainSelector uint64) (*types.Transaction, error)

	FilterAllowedCallerAdded(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*ERC20LockBoxAllowedCallerAddedIterator, error)

	WatchAllowedCallerAdded(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerAdded, token []common.Address, caller []common.Address) (event.Subscription, error)

	ParseAllowedCallerAdded(log types.Log) (*ERC20LockBoxAllowedCallerAdded, error)

	FilterAllowedCallerRemoved(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*ERC20LockBoxAllowedCallerRemovedIterator, error)

	WatchAllowedCallerRemoved(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerRemoved, token []common.Address, caller []common.Address) (event.Subscription, error)

	ParseAllowedCallerRemoved(log types.Log) (*ERC20LockBoxAllowedCallerRemoved, error)

	FilterDeposit(opts *bind.FilterOpts, remoteChainSelector []uint64, depositor []common.Address) (*ERC20LockBoxDepositIterator, error)

	WatchDeposit(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxDeposit, remoteChainSelector []uint64, depositor []common.Address) (event.Subscription, error)

	ParseDeposit(log types.Log) (*ERC20LockBoxDeposit, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20LockBoxOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ERC20LockBoxOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20LockBoxOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ERC20LockBoxOwnershipTransferred, error)

	FilterWithdrawal(opts *bind.FilterOpts, remoteChainSelector []uint64, recipient []common.Address) (*ERC20LockBoxWithdrawalIterator, error)

	WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxWithdrawal, remoteChainSelector []uint64, recipient []common.Address) (event.Subscription, error)

	ParseWithdrawal(log types.Log) (*ERC20LockBoxWithdrawal, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
