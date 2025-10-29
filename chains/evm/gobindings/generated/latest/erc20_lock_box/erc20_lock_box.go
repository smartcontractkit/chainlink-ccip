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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureAllowedCallers\",\"inputs\":[{\"name\":\"configArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct ERC20LockBox.AllowedCallerConfigArgs[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"i_tokenAdminRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract TokenAdminRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedCaller\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedCallerUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RecipientCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a034608b57601f610c6b38819003918201601f19168301916001600160401b03831184841017609057808492602094604052833981010312608b57516001600160a01b03811690819003608b578015607a57608052604051610bc490816100a7823960805181818160c6015281816103f801526108190152f35b6342bcdf7f60e11b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c8063162d2edd146103a0578063181f5a771461034557806347e7ef241461029457806369328dec146100ea5780638ca86f28146100a65763f901fa161461005e57600080fd5b346100a15760406003193601126100a157610077610754565b602435906001600160a01b03821682036100a157602091610097916107d1565b6040519015158152f35b600080fd5b346100a15760006003193601126100a15760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346100a15760606003193601126100a157610103610754565b602435604435916001600160a01b0383168093036100a1576101258282610953565b821561026a576001600160a01b0316906040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481865afa90811561025e5760009161022c575b508082116101fb575060207f2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398916101f26040517fa9059cbb0000000000000000000000000000000000000000000000000000000084820152866024820152826044820152604481526101ec6064826106b8565b856109a4565b604051908152a3005b907fcf4791810000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b90506020813d602011610256575b81610247602093836106b8565b810103126100a1575184610179565b3d915061023a565b6040513d6000823e3d90fd5b7fd87070520000000000000000000000000000000000000000000000000000000060005260046000fd5b346100a15760406003193601126100a1576102ad610754565b6001600160a01b03602435916102c38382610953565b166103176040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152836064820152606481526103116084826106b8565b826109a4565b6040519182527f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6260203393a3005b346100a15760006003193601126100a15761039c604080519061036881836106b8565b601682527f45524332304c6f636b426f7820312e362e322d646576000000000000000000006020830152519182918261070a565b0390f35b346100a15760206003193601126100a15760043567ffffffffffffffff81116100a157366023820112156100a157806004013567ffffffffffffffff81116100a157602482019160243691606084020101116100a1577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031660005b82811061042c57005b6001600160a01b0361044761044283868861076a565b6107a9565b1690811561068e576040517fbbe4f6db000000000000000000000000000000000000000000000000000000008152826004820152602081602481875afa90811561025e5760009161064a575b5060206001600160a01b03916004604051809481937f8da5cb5b000000000000000000000000000000000000000000000000000000008352165afa801561025e57600090610608575b6001600160a01b0391501633036105da5761050360206104fd83878961076a565b016107a9565b91604061051183878961076a565b0135908115158092036100a1576001938160005260006020528260ff6040806000206000906001600160a01b0386168252602052205416151503610559575b50505001610423565b60206001600160a01b037fb2cc4dde7f9044ba1999f7843e2f9cd1e4ce506f8cc2e16de26ce982bf113fa692846000526000835260408060002060009084841682528552207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff88161790556040519586521693a3858080610550565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020823d8211610642575b81610620602093836106b8565b8101031261063f575061063a6001600160a01b03916107bd565b6104dc565b80fd5b3d9150610613565b906020823d8211610686575b81610663602093836106b8565b8101031261063f5750602061067f6001600160a01b03926107bd565b9150610493565b3d9150610656565b7f802c78a20000000000000000000000000000000000000000000000000000000060005260046000fd5b90601f601f19910116810190811067ffffffffffffffff8211176106db57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b9190916020815282519283602083015260005b84811061073e575050601f19601f8460006040809697860101520116010190565b806020809284010151604082860101520161071d565b600435906001600160a01b03821682036100a157565b919081101561077a576060020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b356001600160a01b03811681036100a15790565b51906001600160a01b03821682036100a157565b906001600160a01b03604051927fcb67e3b100000000000000000000000000000000000000000000000000000000845216918260048201526060816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561025e5760009161089c575b5060406001600160a01b03910151166001600160a01b0382161491821561087057505090565b90915060005260006020526001600160a01b036040600020911660005260205260ff6040600020541690565b6060813d60601161094b575b816108b5606093836106b8565b810103126109475760405191606083019083821067ffffffffffffffff83111761091a57509161090f604080936001600160a01b039582526108f6816107bd565b8452610904602082016107bd565b6020850152016107bd565b82820152915061084a565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b5080fd5b3d91506108a8565b906001600160a01b0382161561068e571561097a576109739033906107d1565b156105da57565b7f8b1fa9dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b6001600160a01b03166040916000808451926109c086856106b8565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260208151910182865af13d15610ae6573d9067ffffffffffffffff82116106db57610a36936020601f19601f8501160192610a27875194856106b8565b83523d6000602085013e610aef565b805180610a4257505050565b81602091810103126100a157602001518015908115036100a157610a635750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91610a36926060915b91929015610b6a5750815115610b03575090565b3b15610b0c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015610b7d5750805190602001fd5b610bb3906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161070a565b0390fdfea164736f6c634300081a000a",
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
