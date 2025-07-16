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
	Caller  common.Address
	Allowed bool
}

var ERC20LockBoxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"configureAllowedCallers\",\"inputs\":[{\"name\":\"configArgs\",\"type\":\"tuple[]\",\"internalType\":\"structERC20LockBox.AllowedCallerConfigArgs[]\",\"components\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_tokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecipientCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a03460ab57601f6110d538819003918201601f19168301916001600160401b0383118484101760b05780849260209460405283398101031260ab57516001600160a01b03811680820360ab573315609a57600180546001600160a01b031916331790551560895760805260405161100e90816100c782396080518181816103f9015261053c0152f35b6340163c5160e11b60005260046000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c90816310807aa7146109185750806322ffa47a14610774578063341b9e5a1461071b5780633509c27a1461071b57806379ba5097146106325780637d552ea61461046f5780638da5cb5b1461041d5780639608b232146103ae578063a68012581461033f578063bd028e7c1461018d5763f2fde38b1461009857600080fd5b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885773ffffffffffffffffffffffffffffffffffffffff6100e4610a03565b6100ec610b0d565b1633811461015e57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760043567ffffffffffffffff8111610188573660238201121561018857806004013567ffffffffffffffff8111610188576024820191602436918360061b01011161018857610209610b0d565b60005b81811061021557005b6020610222828486610a3d565b013590811515820361018857600191156102c05761026773ffffffffffffffffffffffffffffffffffffffff61026161025c848789610a3d565b610a7c565b16610e80565b610272575b0161020c565b73ffffffffffffffffffffffffffffffffffffffff61029561025c838688610a3d565b167f663c7e9ed36d9138863ef4306bbfcf01f60e1e7ca69b370c53d3094369e2cb02600080a261026c565b6102ec73ffffffffffffffffffffffffffffffffffffffff6102e661025c848789610a3d565b16610cea565b1561026c5773ffffffffffffffffffffffffffffffffffffffff61031461025c838688610a3d565b167fbc0a6e072a312bde289d32bc84e5b758d7c617f734ecc0d69f995b2d7e69be36600080a261026c565b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760206103a473ffffffffffffffffffffffffffffffffffffffff610390610a03565b166000526004602052604060002054151590565b6040519015158152f35b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018857602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101885760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760043560243567ffffffffffffffff8116809103610188576104cd336000526004602052604060002054151590565b156106045781156105da576105616040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015233602482015230604482015283606482015260648152610525608482610a9d565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016610b58565b80600052600260205260406000208054908382018092116105ab57556040519182527f88ab94ac53260736800da5d05843e504231e9d57ea5cc4ce6479495a52fa296d60203393a3005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f8b1fa9dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760005473ffffffffffffffffffffffffffffffffffffffff811633036106f1577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885767ffffffffffffffff61075b610a26565b1660005260026020526020604060002054604051908152f35b346101885760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610188576004356024359073ffffffffffffffffffffffffffffffffffffffff8216809203610188576044359067ffffffffffffffff8216809203610188576107f5336000526004602052604060002054151590565b156106045782156108ee5780156105da5781600052600260205280604060002054106108ab5781600052600260205260406000208054918083039283116105ab577fc6de56eb9f3f126f4b7f2e63a8477225c96fe39e4b742116b8d81f656820c05292602092556108a26040517fa9059cbb000000000000000000000000000000000000000000000000000000008482015286602482015282604482015260448152610525606482610a9d565b604051908152a3005b816000526002602052604060002054917fd236ce5e0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7fd87070520000000000000000000000000000000000000000000000000000000060005260046000fd5b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610188576003549081815260208101809260036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9060005b8181106109ed5750505081610994910382610a9d565b6040519182916020830190602084525180915260408301919060005b8181106109be575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016109b0565b825484526020909301926001928301920161097e565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361018857565b6004359067ffffffffffffffff8216820361018857565b9190811015610a4d5760061b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036101885790565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610ade57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff600154163303610b2e57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff16604091600080845192610b818685610a9d565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260208151910182865af13d15610cc5573d9067ffffffffffffffff8211610ade57610c159360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160192610c0687519485610a9d565b83523d6000602085013e610ee0565b805180610c2157505050565b8160209181010312610188576020015180159081150361018857610c425750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91610c1592606091610ee0565b8054821015610a4d5760005260206000200190600090565b6000818152600460205260409020548015610e79577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116105ab57600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116105ab57818103610e0a575b5050506003548015610ddb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610d98816003610cd2565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610e61610e1b610e2c936003610cd2565b90549060031b1c9283926003610cd2565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526004602052604060002055388080610d5f565b5050600090565b80600052600460205260406000205415600014610eda5760035468010000000000000000811015610ade57610ec1610e2c8260018594016003556003610cd2565b9055600354906000526004602052604060002055600190565b50600090565b91929015610f5b5750815115610ef4575090565b3b15610efd5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015610f6e5750805190602001fd5b604051907f08c379a0000000000000000000000000000000000000000000000000000000008252602060048301528181519182602483015260005b838110610fe95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604480968601015201168101030190fd5b60208282018101516044878401015285935001610fa956fea164736f6c634300081a000a",
}

var ERC20LockBoxABI = ERC20LockBoxMetaData.ABI

var ERC20LockBoxBin = ERC20LockBoxMetaData.Bin

func DeployERC20LockBox(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *ERC20LockBox, error) {
	parsed, err := ERC20LockBoxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20LockBoxBin), backend, token)
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

func (_ERC20LockBox *ERC20LockBoxCaller) GetAllowedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "getAllowedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) GetAllowedCallers() ([]common.Address, error) {
	return _ERC20LockBox.Contract.GetAllowedCallers(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) GetAllowedCallers() ([]common.Address, error) {
	return _ERC20LockBox.Contract.GetAllowedCallers(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCaller) GetBalance(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "getBalance", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) GetBalance(remoteChainSelector uint64) (*big.Int, error) {
	return _ERC20LockBox.Contract.GetBalance(&_ERC20LockBox.CallOpts, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) GetBalance(remoteChainSelector uint64) (*big.Int, error) {
	return _ERC20LockBox.Contract.GetBalance(&_ERC20LockBox.CallOpts, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxCaller) IToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "i_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) IToken() (common.Address, error) {
	return _ERC20LockBox.Contract.IToken(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) IToken() (common.Address, error) {
	return _ERC20LockBox.Contract.IToken(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCaller) IsAllowedCaller(opts *bind.CallOpts, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "isAllowedCaller", caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) IsAllowedCaller(caller common.Address) (bool, error) {
	return _ERC20LockBox.Contract.IsAllowedCaller(&_ERC20LockBox.CallOpts, caller)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) IsAllowedCaller(caller common.Address) (bool, error) {
	return _ERC20LockBox.Contract.IsAllowedCaller(&_ERC20LockBox.CallOpts, caller)
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

func (_ERC20LockBox *ERC20LockBoxCaller) STokenBalances(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "s_tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) STokenBalances(arg0 uint64) (*big.Int, error) {
	return _ERC20LockBox.Contract.STokenBalances(&_ERC20LockBox.CallOpts, arg0)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) STokenBalances(arg0 uint64) (*big.Int, error) {
	return _ERC20LockBox.Contract.STokenBalances(&_ERC20LockBox.CallOpts, arg0)
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

func (_ERC20LockBox *ERC20LockBoxTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "deposit", amount, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxSession) Deposit(amount *big.Int, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Deposit(&_ERC20LockBox.TransactOpts, amount, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) Deposit(amount *big.Int, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Deposit(&_ERC20LockBox.TransactOpts, amount, remoteChainSelector)
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

func (_ERC20LockBox *ERC20LockBoxTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, recipient common.Address, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "withdraw", amount, recipient, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxSession) Withdraw(amount *big.Int, recipient common.Address, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Withdraw(&_ERC20LockBox.TransactOpts, amount, recipient, remoteChainSelector)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) Withdraw(amount *big.Int, recipient common.Address, remoteChainSelector uint64) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Withdraw(&_ERC20LockBox.TransactOpts, amount, recipient, remoteChainSelector)
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
	Caller common.Address
	Raw    types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterAllowedCallerAdded(opts *bind.FilterOpts, caller []common.Address) (*ERC20LockBoxAllowedCallerAddedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "AllowedCallerAdded", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxAllowedCallerAddedIterator{contract: _ERC20LockBox.contract, event: "AllowedCallerAdded", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchAllowedCallerAdded(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerAdded, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "AllowedCallerAdded", callerRule)
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
	Caller common.Address
	Raw    types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterAllowedCallerRemoved(opts *bind.FilterOpts, caller []common.Address) (*ERC20LockBoxAllowedCallerRemovedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "AllowedCallerRemoved", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxAllowedCallerRemovedIterator{contract: _ERC20LockBox.contract, event: "AllowedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchAllowedCallerRemoved(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerRemoved, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "AllowedCallerRemoved", callerRule)
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
	return common.HexToHash("0x663c7e9ed36d9138863ef4306bbfcf01f60e1e7ca69b370c53d3094369e2cb02")
}

func (ERC20LockBoxAllowedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xbc0a6e072a312bde289d32bc84e5b758d7c617f734ecc0d69f995b2d7e69be36")
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
	GetAllowedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetBalance(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	IToken(opts *bind.CallOpts) (common.Address, error)

	IsAllowedCaller(opts *bind.CallOpts, caller common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	STokenBalances(opts *bind.CallOpts, arg0 uint64) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ConfigureAllowedCallers(opts *bind.TransactOpts, configArgs []ERC20LockBoxAllowedCallerConfigArgs) (*types.Transaction, error)

	Deposit(opts *bind.TransactOpts, amount *big.Int, remoteChainSelector uint64) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, amount *big.Int, recipient common.Address, remoteChainSelector uint64) (*types.Transaction, error)

	FilterAllowedCallerAdded(opts *bind.FilterOpts, caller []common.Address) (*ERC20LockBoxAllowedCallerAddedIterator, error)

	WatchAllowedCallerAdded(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerAdded, caller []common.Address) (event.Subscription, error)

	ParseAllowedCallerAdded(log types.Log) (*ERC20LockBoxAllowedCallerAdded, error)

	FilterAllowedCallerRemoved(opts *bind.FilterOpts, caller []common.Address) (*ERC20LockBoxAllowedCallerRemovedIterator, error)

	WatchAllowedCallerRemoved(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAllowedCallerRemoved, caller []common.Address) (event.Subscription, error)

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
