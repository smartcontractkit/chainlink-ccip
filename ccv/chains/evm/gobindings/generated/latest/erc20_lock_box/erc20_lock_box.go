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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
}

var ERC20LockBoxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTokenSupported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecipientCannotBeZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a0604052346101d9576113cf6020813803918261001c816101de565b9384928339810103126101d957516001600160a01b038116908190036101d957602090610048826101de565b9160008352600036813733156101c857600180546001600160a01b03191633179055610073816101de565b60008152600036813760408051949085016001600160401b038111868210176101b2576040528452808285015260005b815181101561010a576001906001600160a01b036100c18285610203565b5116846100cd82610245565b6100da575b5050016100a3565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138846100d2565b5050915160005b8151811015610182576001600160a01b0361012c8284610203565b5116908115610171577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8583610163600195610343565b50604051908152a101610111565b6342bcdf7f60e11b60005260046000fd5b8280156101715760805260405161102b90816103a482396080518181816105f6015281816109960152610c060152f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101b257604052565b80518210156102175760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156102175760005260206000200190600090565b600081815260036020526040902054801561033c57600019810181811161032657600254600019810191908211610326578082036102d5575b50505060025480156102bf576000190161029981600261022d565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61030e6102e66102f793600261022d565b90549060031b1c928392600261022d565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061027e565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461039d57600254680100000000000000008110156101b2576103846102f7826001859401600255600261022d565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c908163181f5a77146109ba5750806321df0da71461094b5780632451a6271461085d57806374fd18ac1461061b57806375151b631461058c57806379ba5097146104a35780638da5cb5b1461045157806391a2749a14610267578063a36a7fee146101825763f2fde38b1461008d57600080fd5b3461017d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d5773ffffffffffffffffffffffffffffffffffffffff6100d9610a89565b6100e1610cc8565b1633811461015357807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b3461017d5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d576101b9610a89565b6101c1610aac565b5073ffffffffffffffffffffffffffffffffffffffff604435916101e58382610bd3565b166102396040517f23b872dd00000000000000000000000000000000000000000000000000000000602082015233602482015230604482015283606482015260648152610233608482610b0e565b82610d56565b6040519182527f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f6260203393a3005b3461017d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d5760043567ffffffffffffffff811161017d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261017d57604051906102e182610ac3565b806004013567ffffffffffffffff811161017d576103059060043691840101610b4f565b825260248101359067ffffffffffffffff821161017d57600461032b9236920101610b4f565b6020820190815261033a610cc8565b519060005b82518110156103b2578073ffffffffffffffffffffffffffffffffffffffff61036a60019386610d13565b511661037581610df9565b610381575b500161033f565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a18461037a565b505160005b815181101561044f5773ffffffffffffffffffffffffffffffffffffffff6103df8284610d13565b5116908115610425577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef602083610417600195610fbe565b50604051908152a1016103b7565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b005b3461017d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461017d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d5760005473ffffffffffffffffffffffffffffffffffffffff81163303610562577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461017d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d5760206105c5610a89565b73ffffffffffffffffffffffffffffffffffffffff604051911673ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148152f35b3461017d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d57610652610a89565b61065a610aac565b506044356064359173ffffffffffffffffffffffffffffffffffffffff831680930361017d57819061068c8382610bd3565b83156108335773ffffffffffffffffffffffffffffffffffffffff1691604051907f70a08231000000000000000000000000000000000000000000000000000000008252306004830152602082602481875afa918215610827576000926107d0575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff146107c8575b808211610797575060207f2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b63989161078e6040517fa9059cbb000000000000000000000000000000000000000000000000000000008482015286602482015282604482015260448152610788606482610b0e565b85610d56565b604051908152a3005b907fcf4791810000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b905080610716565b90916020823d60201161081f575b816107eb60209383610b0e565b8101031261081c575051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6106ee565b80fd5b3d91506107de565b6040513d6000823e3d90fd5b7fd87070520000000000000000000000000000000000000000000000000000000060005260046000fd5b3461017d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061093557505050816108dc910382610b0e565b6040519182916020830190602084525180915260408301919060005b818110610906575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016108f8565b82548452602090930192600192830192016108c6565b3461017d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461017d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017d576109f281610ac3565b601681527f45524332304c6f636b426f7820322e302e302d64657600000000000000000000602082015260405190602082528181519182602083015260005b838110610a715750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001610a31565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361017d57565b6024359067ffffffffffffffff8216820361017d57565b6040810190811067ffffffffffffffff821117610adf57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610adf57604052565b81601f8201121561017d5780359167ffffffffffffffff8311610adf578260051b9160405193610b826020850186610b0e565b845260208085019382010191821161017d57602001915b818310610ba65750505090565b823573ffffffffffffffffffffffffffffffffffffffff8116810361017d57815260209283019201610b99565b9015610c9e5773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168103610c71575033600052600360205260406000205415610c4357565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fbf16aab60000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f8b1fa9dd0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff600154163303610ce957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015610d275760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000602091828151910182855af115610827576000513d610dd8575073ffffffffffffffffffffffffffffffffffffffff81163b155b610d945750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415610d8d565b8054821015610d275760005260206000200190600090565b6000818152600360205260409020548015610fb7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610f8857600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610f8857808203610f19575b5050506002548015610eea577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610ea7816002610de1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610f70610f2a610f3b936002610de1565b90549060031b1c9283926002610de1565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080610e6e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146110185760025468010000000000000000811015610adf57610fff610f3b8260018594016002556002610de1565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
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

func (_ERC20LockBox *ERC20LockBoxCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _ERC20LockBox.Contract.GetAllAuthorizedCallers(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _ERC20LockBox.Contract.GetAllAuthorizedCallers(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) GetToken() (common.Address, error) {
	return _ERC20LockBox.Contract.GetToken(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) GetToken() (common.Address, error) {
	return _ERC20LockBox.Contract.GetToken(&_ERC20LockBox.CallOpts)
}

func (_ERC20LockBox *ERC20LockBoxCaller) IsTokenSupported(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20LockBox.contract.Call(opts, &out, "isTokenSupported", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ERC20LockBox *ERC20LockBoxSession) IsTokenSupported(token common.Address) (bool, error) {
	return _ERC20LockBox.Contract.IsTokenSupported(&_ERC20LockBox.CallOpts, token)
}

func (_ERC20LockBox *ERC20LockBoxCallerSession) IsTokenSupported(token common.Address) (bool, error) {
	return _ERC20LockBox.Contract.IsTokenSupported(&_ERC20LockBox.CallOpts, token)
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

func (_ERC20LockBox *ERC20LockBoxTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "acceptOwnership")
}

func (_ERC20LockBox *ERC20LockBoxSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC20LockBox.Contract.AcceptOwnership(&_ERC20LockBox.TransactOpts)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC20LockBox.Contract.AcceptOwnership(&_ERC20LockBox.TransactOpts)
}

func (_ERC20LockBox *ERC20LockBoxTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_ERC20LockBox *ERC20LockBoxSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.ApplyAuthorizedCallerUpdates(&_ERC20LockBox.TransactOpts, authorizedCallerArgs)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.ApplyAuthorizedCallerUpdates(&_ERC20LockBox.TransactOpts, authorizedCallerArgs)
}

func (_ERC20LockBox *ERC20LockBoxTransactor) Deposit(opts *bind.TransactOpts, token common.Address, arg1 uint64, amount *big.Int) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "deposit", token, arg1, amount)
}

func (_ERC20LockBox *ERC20LockBoxSession) Deposit(token common.Address, arg1 uint64, amount *big.Int) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Deposit(&_ERC20LockBox.TransactOpts, token, arg1, amount)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) Deposit(token common.Address, arg1 uint64, amount *big.Int) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Deposit(&_ERC20LockBox.TransactOpts, token, arg1, amount)
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

func (_ERC20LockBox *ERC20LockBoxTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, arg1 uint64, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.contract.Transact(opts, "withdraw", token, arg1, amount, recipient)
}

func (_ERC20LockBox *ERC20LockBoxSession) Withdraw(token common.Address, arg1 uint64, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Withdraw(&_ERC20LockBox.TransactOpts, token, arg1, amount, recipient)
}

func (_ERC20LockBox *ERC20LockBoxTransactorSession) Withdraw(token common.Address, arg1 uint64, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20LockBox.Contract.Withdraw(&_ERC20LockBox.TransactOpts, token, arg1, amount, recipient)
}

type ERC20LockBoxAuthorizedCallerAddedIterator struct {
	Event *ERC20LockBoxAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxAuthorizedCallerAdded)
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
		it.Event = new(ERC20LockBoxAuthorizedCallerAdded)
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

func (it *ERC20LockBoxAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*ERC20LockBoxAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxAuthorizedCallerAddedIterator{contract: _ERC20LockBox.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxAuthorizedCallerAdded)
				if err := _ERC20LockBox.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseAuthorizedCallerAdded(log types.Log) (*ERC20LockBoxAuthorizedCallerAdded, error) {
	event := new(ERC20LockBoxAuthorizedCallerAdded)
	if err := _ERC20LockBox.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ERC20LockBoxAuthorizedCallerRemovedIterator struct {
	Event *ERC20LockBoxAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ERC20LockBoxAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LockBoxAuthorizedCallerRemoved)
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
		it.Event = new(ERC20LockBoxAuthorizedCallerRemoved)
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

func (it *ERC20LockBoxAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *ERC20LockBoxAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ERC20LockBoxAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_ERC20LockBox *ERC20LockBoxFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*ERC20LockBoxAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _ERC20LockBox.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &ERC20LockBoxAuthorizedCallerRemovedIterator{contract: _ERC20LockBox.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_ERC20LockBox *ERC20LockBoxFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _ERC20LockBox.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ERC20LockBoxAuthorizedCallerRemoved)
				if err := _ERC20LockBox.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_ERC20LockBox *ERC20LockBoxFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*ERC20LockBoxAuthorizedCallerRemoved, error) {
	event := new(ERC20LockBoxAuthorizedCallerRemoved)
	if err := _ERC20LockBox.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (ERC20LockBoxAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (ERC20LockBoxAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (ERC20LockBoxDeposit) Topic() common.Hash {
	return common.HexToHash("0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62")
}

func (ERC20LockBoxOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ERC20LockBoxOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ERC20LockBoxWithdrawal) Topic() common.Hash {
	return common.HexToHash("0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398")
}

func (_ERC20LockBox *ERC20LockBox) Address() common.Address {
	return _ERC20LockBox.address
}

type ERC20LockBoxInterface interface {
	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsTokenSupported(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	Deposit(opts *bind.TransactOpts, token common.Address, arg1 uint64, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, token common.Address, arg1 uint64, amount *big.Int, recipient common.Address) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*ERC20LockBoxAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*ERC20LockBoxAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*ERC20LockBoxAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*ERC20LockBoxAuthorizedCallerRemoved, error)

	FilterDeposit(opts *bind.FilterOpts, token []common.Address, depositor []common.Address) (*ERC20LockBoxDepositIterator, error)

	WatchDeposit(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxDeposit, token []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseDeposit(log types.Log) (*ERC20LockBoxDeposit, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20LockBoxOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ERC20LockBoxOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20LockBoxOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ERC20LockBoxOwnershipTransferred, error)

	FilterWithdrawal(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*ERC20LockBoxWithdrawalIterator, error)

	WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ERC20LockBoxWithdrawal, token []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseWithdrawal(log types.Log) (*ERC20LockBoxWithdrawal, error)

	Address() common.Address
}
