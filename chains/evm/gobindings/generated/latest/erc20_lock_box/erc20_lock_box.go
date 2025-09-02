// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20_lock_box

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

type ERC20LockBoxAllowedCallerConfigArgs struct {
	Token   common.Address
	Caller  common.Address
	Allowed bool
}

var ERC20LockBoxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureAllowedCallers\",\"inputs\":[{\"name\":\"configArgs\",\"type\":\"tuple[]\",\"internalType\":\"structERC20LockBox.AllowedCallerConfigArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"i_tokenAdminRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractTokenAdminRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedCaller\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedCallerUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RecipientCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a034608c57601f610eb238819003918201601f19168301916001600160401b03831184841017609157808492602094604052833981010312608c57516001600160a01b03811690819003608c578015607b57608052604051610e0a90816100a8823960805181818161011c015281816104ed01526109f30152f35b6342bcdf7f60e11b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063162d2edd14610477578063181f5a77146103fe57806347e7ef241461032257806369328dec146101405780638ca86f28146100d15763f901fa161461005e57600080fd5b346100cc5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100cc576100956108ed565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036100cc576020916100c291610991565b6040519015158152f35b600080fd5b346100cc5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100cc57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346100cc5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100cc576101776108ed565b6024356044359173ffffffffffffffffffffffffffffffffffffffff83168093036100cc576101a68282610b61565b82156102f85773ffffffffffffffffffffffffffffffffffffffff16906040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481865afa9081156102ec576000916102ba575b50808211610289575060207f2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398916102806040517fa9059cbb00000000000000000000000000000000000000000000000000000000848201528660248201528260448201526044815261027a606482610815565b85610bbf565b604051908152a3005b907fcf4791810000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b90506020813d6020116102e4575b816102d560209383610815565b810103126100cc575184610207565b3d91506102c8565b6040513d6000823e3d90fd5b7fd87070520000000000000000000000000000000000000000000000000000000060005260046000fd5b346100cc5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100cc576103596108ed565b73ffffffffffffffffffffffffffffffffffffffff6024359161037c8382610b61565b166103d06040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152836064820152606481526103ca608482610815565b82610bbf565b6040519182527f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6260203393a3005b346100cc5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100cc57610473604080519061043f8183610815565b601682527f45524332304c6f636b426f7820312e362e322d6465760000000000000000000060208301525191829182610885565b0390f35b346100cc5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100cc5760043567ffffffffffffffff81116100cc57366023820112156100cc57806004013567ffffffffffffffff81116100cc57602482019160243691606084020101116100cc577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1660005b82811061052e57005b73ffffffffffffffffffffffffffffffffffffffff610556610551838688610910565b61094f565b169081156107eb576040517fbbe4f6db000000000000000000000000000000000000000000000000000000008152826004820152602081602481875afa9081156102ec5760009161079a575b50602073ffffffffffffffffffffffffffffffffffffffff916004604051809481937f8da5cb5b000000000000000000000000000000000000000000000000000000008352165afa80156102ec5760009061074b575b73ffffffffffffffffffffffffffffffffffffffff915016330361071d5761062c6020610626838789610910565b0161094f565b91604061063a838789610910565b0135908115158092036100cc576001938160005260006020528260ff60408060002060009073ffffffffffffffffffffffffffffffffffffffff8616825260205220541615150361068f575b50505001610525565b602073ffffffffffffffffffffffffffffffffffffffff7fb2cc4dde7f9044ba1999f7843e2f9cd1e4ce506f8cc2e16de26ce982bf113fa692846000526000835260408060002060009084841682528552207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff88161790556040519586521693a3858080610686565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020823d8211610792575b8161076360209383610815565b8101031261078f575061078a73ffffffffffffffffffffffffffffffffffffffff91610970565b6105f8565b80fd5b3d9150610756565b906020823d82116107e3575b816107b360209383610815565b8101031261078f575060206107dc73ffffffffffffffffffffffffffffffffffffffff92610970565b91506105a2565b3d91506107a6565b7f802c78a20000000000000000000000000000000000000000000000000000000060005260046000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761085657604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b9190916020815282519283602083015260005b8481106108d75750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006040809697860101520116010190565b8060208092840101516040828601015201610898565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036100cc57565b9190811015610920576060020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036100cc5790565b519073ffffffffffffffffffffffffffffffffffffffff821682036100cc57565b9073ffffffffffffffffffffffffffffffffffffffff604051927fcb67e3b1000000000000000000000000000000000000000000000000000000008452169182600482015260608160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156102ec57600091610a9d575b50604073ffffffffffffffffffffffffffffffffffffffff9101511673ffffffffffffffffffffffffffffffffffffffff821614918215610a6457505090565b909150600052600060205273ffffffffffffffffffffffffffffffffffffffff6040600020911660005260205260ff6040600020541690565b6060813d606011610b59575b81610ab660609383610815565b81010312610b555760405191606083019083821067ffffffffffffffff831117610b28575091610b1d6040809373ffffffffffffffffffffffffffffffffffffffff958252610b0481610970565b8452610b1260208201610970565b602085015201610970565b828201529150610a24565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b5080fd5b3d9150610aa9565b9073ffffffffffffffffffffffffffffffffffffffff8216156107eb5715610b9557610b8e903390610991565b1561071d57565b7f8b1fa9dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff16604091600080845192610be88685610815565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260208151910182865af13d15610d2c573d9067ffffffffffffffff821161085657610c7c9360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160192610c6d87519485610815565b83523d6000602085013e610d35565b805180610c8857505050565b81602091810103126100cc57602001518015908115036100cc57610ca95750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91610c7c926060915b91929015610db05750815115610d49575090565b3b15610d525790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015610dc35750805190602001fd5b610df9906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610885565b0390fdfea164736f6c634300081a000a",
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

func (_ERC20LockBox *ERC20LockBoxCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) TypeAndVersion() (string, error) {
	return _ERC20LockBox.Contract.TypeAndVersion(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) TypeAndVersion() (string, error) {
	return _ERC20LockBox.Contract.TypeAndVersion(&_ERC20LockBox.CallOpts)
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

func (_ERC20LockBox *ERC20LockBoxTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "deposit", token, amount)
}

func (_ERC20LockBox *ERC20LockBoxSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Deposit(&_ERC20LockBox.TransactOpts, token, amount)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Deposit(&_ERC20LockBox.TransactOpts, token, amount)
}

func (_ERC20LockBox *ERC20LockBoxTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "withdraw", token, amount, recipient)
}

func (_ERC20LockBox *ERC20LockBoxSession) Withdraw(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Withdraw(&_ERC20LockBox.TransactOpts, token, amount, recipient)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) Withdraw(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Withdraw(&_ERC20LockBox.TransactOpts, token, amount, recipient)
}

type ERC20LockBoxAllowedCallerUpdatedIterator struct {
	Event *ERC20LockBoxAllowedCallerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxAllowedCallerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxAllowedCallerUpdated)
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
		it.Event = new(ERC20LockBoxAllowedCallerUpdated)
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

func (it *ERC20LockBoxAllowedCallerUpdatedIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxAllowedCallerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxAllowedCallerUpdated struct {
	Token   common.Address
	Caller  common.Address
	Allowed bool
	Raw     types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterAllowedCallerUpdated(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*ERC20LockBoxAllowedCallerUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "AllowedCallerUpdated", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxAllowedCallerUpdatedIterator{contract: _ERC20LockBox.contract, event: "AllowedCallerUpdated", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchAllowedCallerUpdated(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerUpdated, token []common.Address, caller []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "AllowedCallerUpdated", tokenRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxAllowedCallerUpdated)
				if err := _ERC20LockBox.contract.UnpackLog(event, "AllowedCallerUpdated", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseAllowedCallerUpdated(log types.Log) (*ERC20LockBoxAllowedCallerUpdated, error) {
	event := new(ERC20LockBoxAllowedCallerUpdated)
	if err := _ERC20LockBox.contract.UnpackLog(event, "AllowedCallerUpdated", log); err != nil {
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
	Token     common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterDeposit(opts *bind.FilterOpts, token []common.Address, depositor []common.Address) (*ERC20LockBoxDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "Deposit", tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxDepositIterator{contract: _ERC20LockBox.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxDeposit, token []common.Address, depositor []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "Deposit", tokenRule, depositorRule)
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
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterWithdrawal(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*ERC20LockBoxWithdrawalIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "Withdrawal", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxWithdrawalIterator{contract: _ERC20LockBox.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxWithdrawal, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "Withdrawal", tokenRule, recipientRule)
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

func (ERC20LockBoxAllowedCallerUpdated) Topic() common.Hash {
	return common.HexToHash("0xb2cc4dde7f9044ba1999f7843e2f9cd1e4ce506f8cc2e16de26ce982bf113fa6")
}

func (ERC20LockBoxDeposit) Topic() common.Hash {
	return common.HexToHash("0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62")
}

func (ERC20LockBoxWithdrawal) Topic() common.Hash {
	return common.HexToHash("0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398")
}

func (_ERC20LockBox *ERC20LockBox) Address() common.Address {
	return _ERC20LockBox.address
}

type ERC20LockBoxInterface interface {
	ITokenAdminRegistry(opts *bind.CallOpts) (common.Address, error)

	IsAllowedCaller(opts *bind.CallOpts, token common.Address, caller common.Address) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	ConfigureAllowedCallers(opts *bind.TransactOpts, configArgs []ERC20LockBoxAllowedCallerConfigArgs) (*types.Transaction, error)

	Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error)

	FilterAllowedCallerUpdated(opts *bind.FilterOpts, token []common.Address, caller []common.Address) (*ERC20LockBoxAllowedCallerUpdatedIterator, error)

	WatchAllowedCallerUpdated(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerUpdated, token []common.Address, caller []common.Address) (event.Subscription, error)

	ParseAllowedCallerUpdated(log types.Log) (*ERC20LockBoxAllowedCallerUpdated, error)

	FilterDeposit(opts *bind.FilterOpts, token []common.Address, depositor []common.Address) (*ERC20LockBoxDepositIterator, error)

	WatchDeposit(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxDeposit, token []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseDeposit(log types.Log) (*ERC20LockBoxDeposit, error)

	FilterWithdrawal(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*ERC20LockBoxWithdrawalIterator, error)

	WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxWithdrawal, token []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseWithdrawal(log types.Log) (*ERC20LockBoxWithdrawal, error)

	Address() common.Address
}
