// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rmn

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

var RMNMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"curseAdmins\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"curse\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"curse\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCursedSubjects\",\"inputs\":[],\"outputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isCursed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uncurse\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uncurse\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"internalType\":\"bytes16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Cursed\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"indexed\":false,\"internalType\":\"bytes16[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Uncursed\",\"inputs\":[{\"name\":\"subjects\",\"type\":\"bytes16[]\",\"indexed\":false,\"internalType\":\"bytes16[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCursed\",\"inputs\":[{\"name\":\"subject\",\"type\":\"bytes16\",\"internalType\":\"bytes16\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60806040523461020f576117008038038061001981610214565b92833981019060208183031261020f578051906001600160401b03821161020f570181601f8201121561020f578051916001600160401b0383116101c8578260051b9160208061006a818601610214565b80968152019382010191821161020f57602001915b8183106101ef578333156101de57600180546001600160a01b031916331790556020906100ab82610214565b60008152600036813760408051929083016001600160401b038111848210176101c8576040528252808383015260005b8151811015610142576001906001600160a01b036100f98285610239565b5116856101058261027b565b610112575b5050016100db565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1858561010a565b50505160005b81518110156101b9576001600160a01b036101638284610239565b51169081156101a8577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef848361019a600195610379565b50604051908152a101610148565b6342bcdf7f60e11b60005260046000fd5b60405161132690816103da8239f35b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b82516001600160a01b038116810361020f5781526020928301920161007f565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176101c857604052565b805182101561024d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561024d5760005260206000200190600090565b600081815260036020526040902054801561037257600019810181811161035c5760025460001981019190821161035c5780820361030b575b50505060025480156102f557600019016102cf816002610263565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61034461031c61032d936002610263565b90549060031b1c9283926002610263565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806102b4565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146103d357600254680100000000000000008110156101c8576103ba61032d8260018594016002556002610263565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c908163181f5a77146109b4575080632451a627146108c65780632cbc26bb14610885578063397796f71461084257806362eed415146107995780636d2d39931461067957806379ba5097146105905780638da5cb5b1461053e57806391a2749a146103565780639a19b32914610263578063d881e092146101be578063f2fde38b146100cb5763f8bb876e146100ae57600080fd5b346100c6576100c46100bf36610bd1565b610e3b565b005b600080fd5b346100c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c65760043573ffffffffffffffffffffffffffffffffffffffff81168091036100c657610123610f7b565b33811461019457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c65760405180602060045491828152019060046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9060005b81811061024d576102498561023d81870382610ace565b60405191829182610c98565b0390f35b8254845260209093019260019283019201610226565b346100c65761027136610bd1565b610279610f7b565b60005b8151811015610321576102ba7fffffffffffffffffffffffffffffffff000000000000000000000000000000006102b38385610e27565b5116611071565b156102c75760010161027c565b6102f2907fffffffffffffffffffffffffffffffff0000000000000000000000000000000092610e27565b51167f73281fa10000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6040517f0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba190806103518582610c98565b0390a1005b346100c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c65760043567ffffffffffffffff81116100c65760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100c657604051906103d082610a83565b806004013567ffffffffffffffff81116100c6576103f49060043691840101610b56565b825260248101359067ffffffffffffffff82116100c657600461041a9236920101610b56565b60208201908152610429610f7b565b519060005b82518110156104a1578073ffffffffffffffffffffffffffffffffffffffff61045960019386610e27565b511661046481611249565b610470575b500161042e565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184610469565b505160005b81518110156100c45773ffffffffffffffffffffffffffffffffffffffff6104ce8284610e27565b5116908115610514577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6020836105066001956111d7565b50604051908152a1016104a6565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c65760005473ffffffffffffffffffffffffffffffffffffffff8116330361064f577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c6576106b0610b0f565b6040908151906106c08383610ace565b600182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083013660208401377fffffffffffffffffffffffffffffffff0000000000000000000000000000000061071783610deb565b91169052610723610f7b565b60005b815181101561076a5761075d7fffffffffffffffffffffffffffffffff000000000000000000000000000000006102b38385610e27565b156102c757600101610726565b82517f0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba190806103518582610c98565b346100c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c6576100c46107d3610b0f565b6040907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08251926108048185610ace565b60018452013660208401377fffffffffffffffffffffffffffffffff0000000000000000000000000000000061083983610deb565b91169052610e3b565b346100c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c657602061087b610d8e565b6040519015158152f35b346100c65760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c657602061087b6108c1610b0f565b610cf4565b346100c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c6576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b81811061099e5750505081610945910382610ace565b6040519182916020830190602084525180915260408301919060005b81811061096f575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610961565b825484526020909301926001928301920161092f565b346100c65760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c6576109ec81610a83565b600981527f524d4e20322e312e300000000000000000000000000000000000000000000000602082015260405190602082528181519182602083015260005b838110610a6b5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001610a2b565b6040810190811067ffffffffffffffff821117610a9f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610a9f57604052565b600435907fffffffffffffffffffffffffffffffff00000000000000000000000000000000821682036100c657565b67ffffffffffffffff8111610a9f5760051b60200190565b9080601f830112156100c657813590610b6e82610b3e565b92610b7c6040519485610ace565b82845260208085019360051b8201019182116100c657602001915b818310610ba45750505090565b823573ffffffffffffffffffffffffffffffffffffffff811681036100c657815260209283019201610b97565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126100c6576004359067ffffffffffffffff82116100c657806023830112156100c657816004013590610c2882610b3e565b92610c366040519485610ace565b8284526024602085019360051b8201019182116100c657602401915b818310610c5f5750505090565b82357fffffffffffffffffffffffffffffffff00000000000000000000000000000000811681036100c657815260209283019201610c52565b602060408183019282815284518094520192019060005b818110610cbc5750505090565b82517fffffffffffffffffffffffffffffffff0000000000000000000000000000000016845260209384019390920191600101610caf565b60045415610d88577fffffffffffffffffffffffffffffffff0000000000000000000000000000000016600052600560205260406000205415801590610d375790565b507f010000000000000000000000000000010000000000000000000000000000000060005260056020527f8f496e4ceafb62bf7f18e44784f657270af67789253a1cc665c8d949978172bc54151590565b50600090565b60045415610de6577f010000000000000000000000000000010000000000000000000000000000000060005260056020527f8f496e4ceafb62bf7f18e44784f657270af67789253a1cc665c8d949978172bc54151590565b600090565b805115610df85760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051821015610df85760209160051b010190565b73ffffffffffffffffffffffffffffffffffffffff600154163303610f3a575b60005b8151811015610f0357610e9c7fffffffffffffffffffffffffffffffff00000000000000000000000000000000610e958385610e27565b5116611210565b15610ea957600101610e5e565b610ed4907fffffffffffffffffffffffffffffffff0000000000000000000000000000000092610e27565b51167f19d5c79b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b50610f357f1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f749160405191829182610c98565b0390a1565b336000526003602052604060002054610e5b577fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff600154163303610f9c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8054821015610df85760005260206000200190600090565b80548015611042577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906110138282610fc6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008181526005602052604090205480156111a5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161117657600454907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161117657818103611107575b5050506110f36004610fde565b600052600560205260006040812055600190565b61115e611118611129936004610fc6565b90549060031b1c9283926004610fc6565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260056020526040600020553880806110e6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b80549068010000000000000000821015610a9f57816111299160016111d394018155610fc6565b9055565b80600052600360205260406000205415600014610d88576111f98160026111ac565b600254906000526003602052604060002055600190565b80600052600560205260406000205415600014610d88576112328160046111ac565b600454906000526005602052604060002055600190565b60008181526003602052604090205480156111a5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161117657600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611176578082036112df575b5050506112cb6002610fde565b600052600360205260006040812055600190565b6113016112f0611129936002610fc6565b90549060031b1c9283926002610fc6565b905560005260036020526040600020553880806112be56fea164736f6c634300081a000a",
}

var RMNABI = RMNMetaData.ABI

var RMNBin = RMNMetaData.Bin

func DeployRMN(auth *bind.TransactOpts, backend bind.ContractBackend, curseAdmins []common.Address) (common.Address, *types.Transaction, *RMN, error) {
	parsed, err := RMNMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RMNBin), backend, curseAdmins)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RMN{address: address, abi: *parsed, RMNCaller: RMNCaller{contract: contract}, RMNTransactor: RMNTransactor{contract: contract}, RMNFilterer: RMNFilterer{contract: contract}}, nil
}

type RMN struct {
	address common.Address
	abi     abi.ABI
	RMNCaller
	RMNTransactor
	RMNFilterer
}

type RMNCaller struct {
	contract *bind.BoundContract
}

type RMNTransactor struct {
	contract *bind.BoundContract
}

type RMNFilterer struct {
	contract *bind.BoundContract
}

type RMNSession struct {
	Contract     *RMN
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RMNCallerSession struct {
	Contract *RMNCaller
	CallOpts bind.CallOpts
}

type RMNTransactorSession struct {
	Contract     *RMNTransactor
	TransactOpts bind.TransactOpts
}

type RMNRaw struct {
	Contract *RMN
}

type RMNCallerRaw struct {
	Contract *RMNCaller
}

type RMNTransactorRaw struct {
	Contract *RMNTransactor
}

func NewRMN(address common.Address, backend bind.ContractBackend) (*RMN, error) {
	abi, err := abi.JSON(strings.NewReader(RMNABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRMN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RMN{address: address, abi: abi, RMNCaller: RMNCaller{contract: contract}, RMNTransactor: RMNTransactor{contract: contract}, RMNFilterer: RMNFilterer{contract: contract}}, nil
}

func NewRMNCaller(address common.Address, caller bind.ContractCaller) (*RMNCaller, error) {
	contract, err := bindRMN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RMNCaller{contract: contract}, nil
}

func NewRMNTransactor(address common.Address, transactor bind.ContractTransactor) (*RMNTransactor, error) {
	contract, err := bindRMN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RMNTransactor{contract: contract}, nil
}

func NewRMNFilterer(address common.Address, filterer bind.ContractFilterer) (*RMNFilterer, error) {
	contract, err := bindRMN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RMNFilterer{contract: contract}, nil
}

func bindRMN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RMNMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_RMN *RMNRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RMN.Contract.RMNCaller.contract.Call(opts, result, method, params...)
}

func (_RMN *RMNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMN.Contract.RMNTransactor.contract.Transfer(opts)
}

func (_RMN *RMNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RMN.Contract.RMNTransactor.contract.Transact(opts, method, params...)
}

func (_RMN *RMNCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RMN.Contract.contract.Call(opts, result, method, params...)
}

func (_RMN *RMNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMN.Contract.contract.Transfer(opts)
}

func (_RMN *RMNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RMN.Contract.contract.Transact(opts, method, params...)
}

func (_RMN *RMNCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RMN.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_RMN *RMNSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _RMN.Contract.GetAllAuthorizedCallers(&_RMN.CallOpts)
}

func (_RMN *RMNCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _RMN.Contract.GetAllAuthorizedCallers(&_RMN.CallOpts)
}

func (_RMN *RMNCaller) GetCursedSubjects(opts *bind.CallOpts) ([][16]byte, error) {
	var out []interface{}
	err := _RMN.contract.Call(opts, &out, "getCursedSubjects")

	if err != nil {
		return *new([][16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][16]byte)).(*[][16]byte)

	return out0, err

}

func (_RMN *RMNSession) GetCursedSubjects() ([][16]byte, error) {
	return _RMN.Contract.GetCursedSubjects(&_RMN.CallOpts)
}

func (_RMN *RMNCallerSession) GetCursedSubjects() ([][16]byte, error) {
	return _RMN.Contract.GetCursedSubjects(&_RMN.CallOpts)
}

func (_RMN *RMNCaller) IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error) {
	var out []interface{}
	err := _RMN.contract.Call(opts, &out, "isCursed", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RMN *RMNSession) IsCursed(subject [16]byte) (bool, error) {
	return _RMN.Contract.IsCursed(&_RMN.CallOpts, subject)
}

func (_RMN *RMNCallerSession) IsCursed(subject [16]byte) (bool, error) {
	return _RMN.Contract.IsCursed(&_RMN.CallOpts, subject)
}

func (_RMN *RMNCaller) IsCursed0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RMN.contract.Call(opts, &out, "isCursed0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RMN *RMNSession) IsCursed0() (bool, error) {
	return _RMN.Contract.IsCursed0(&_RMN.CallOpts)
}

func (_RMN *RMNCallerSession) IsCursed0() (bool, error) {
	return _RMN.Contract.IsCursed0(&_RMN.CallOpts)
}

func (_RMN *RMNCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RMN.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RMN *RMNSession) Owner() (common.Address, error) {
	return _RMN.Contract.Owner(&_RMN.CallOpts)
}

func (_RMN *RMNCallerSession) Owner() (common.Address, error) {
	return _RMN.Contract.Owner(&_RMN.CallOpts)
}

func (_RMN *RMNCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RMN.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_RMN *RMNSession) TypeAndVersion() (string, error) {
	return _RMN.Contract.TypeAndVersion(&_RMN.CallOpts)
}

func (_RMN *RMNCallerSession) TypeAndVersion() (string, error) {
	return _RMN.Contract.TypeAndVersion(&_RMN.CallOpts)
}

func (_RMN *RMNTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RMN.contract.Transact(opts, "acceptOwnership")
}

func (_RMN *RMNSession) AcceptOwnership() (*types.Transaction, error) {
	return _RMN.Contract.AcceptOwnership(&_RMN.TransactOpts)
}

func (_RMN *RMNTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RMN.Contract.AcceptOwnership(&_RMN.TransactOpts)
}

func (_RMN *RMNTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _RMN.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_RMN *RMNSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _RMN.Contract.ApplyAuthorizedCallerUpdates(&_RMN.TransactOpts, authorizedCallerArgs)
}

func (_RMN *RMNTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _RMN.Contract.ApplyAuthorizedCallerUpdates(&_RMN.TransactOpts, authorizedCallerArgs)
}

func (_RMN *RMNTransactor) Curse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error) {
	return _RMN.contract.Transact(opts, "curse", subject)
}

func (_RMN *RMNSession) Curse(subject [16]byte) (*types.Transaction, error) {
	return _RMN.Contract.Curse(&_RMN.TransactOpts, subject)
}

func (_RMN *RMNTransactorSession) Curse(subject [16]byte) (*types.Transaction, error) {
	return _RMN.Contract.Curse(&_RMN.TransactOpts, subject)
}

func (_RMN *RMNTransactor) Curse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error) {
	return _RMN.contract.Transact(opts, "curse0", subjects)
}

func (_RMN *RMNSession) Curse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMN.Contract.Curse0(&_RMN.TransactOpts, subjects)
}

func (_RMN *RMNTransactorSession) Curse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMN.Contract.Curse0(&_RMN.TransactOpts, subjects)
}

func (_RMN *RMNTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RMN.contract.Transact(opts, "transferOwnership", to)
}

func (_RMN *RMNSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMN.Contract.TransferOwnership(&_RMN.TransactOpts, to)
}

func (_RMN *RMNTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _RMN.Contract.TransferOwnership(&_RMN.TransactOpts, to)
}

func (_RMN *RMNTransactor) Uncurse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error) {
	return _RMN.contract.Transact(opts, "uncurse", subject)
}

func (_RMN *RMNSession) Uncurse(subject [16]byte) (*types.Transaction, error) {
	return _RMN.Contract.Uncurse(&_RMN.TransactOpts, subject)
}

func (_RMN *RMNTransactorSession) Uncurse(subject [16]byte) (*types.Transaction, error) {
	return _RMN.Contract.Uncurse(&_RMN.TransactOpts, subject)
}

func (_RMN *RMNTransactor) Uncurse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error) {
	return _RMN.contract.Transact(opts, "uncurse0", subjects)
}

func (_RMN *RMNSession) Uncurse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMN.Contract.Uncurse0(&_RMN.TransactOpts, subjects)
}

func (_RMN *RMNTransactorSession) Uncurse0(subjects [][16]byte) (*types.Transaction, error) {
	return _RMN.Contract.Uncurse0(&_RMN.TransactOpts, subjects)
}

type RMNAuthorizedCallerAddedIterator struct {
	Event *RMNAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNAuthorizedCallerAdded)
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
		it.Event = new(RMNAuthorizedCallerAdded)
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

func (it *RMNAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *RMNAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_RMN *RMNFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*RMNAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _RMN.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &RMNAuthorizedCallerAddedIterator{contract: _RMN.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_RMN *RMNFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *RMNAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _RMN.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNAuthorizedCallerAdded)
				if err := _RMN.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_RMN *RMNFilterer) ParseAuthorizedCallerAdded(log types.Log) (*RMNAuthorizedCallerAdded, error) {
	event := new(RMNAuthorizedCallerAdded)
	if err := _RMN.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNAuthorizedCallerRemovedIterator struct {
	Event *RMNAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNAuthorizedCallerRemoved)
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
		it.Event = new(RMNAuthorizedCallerRemoved)
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

func (it *RMNAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *RMNAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_RMN *RMNFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*RMNAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _RMN.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &RMNAuthorizedCallerRemovedIterator{contract: _RMN.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_RMN *RMNFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *RMNAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _RMN.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNAuthorizedCallerRemoved)
				if err := _RMN.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_RMN *RMNFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*RMNAuthorizedCallerRemoved, error) {
	event := new(RMNAuthorizedCallerRemoved)
	if err := _RMN.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNCursedIterator struct {
	Event *RMNCursed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNCursedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNCursed)
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
		it.Event = new(RMNCursed)
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

func (it *RMNCursedIterator) Error() error {
	return it.fail
}

func (it *RMNCursedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNCursed struct {
	Subjects [][16]byte
	Raw      types.Log
}

func (_RMN *RMNFilterer) FilterCursed(opts *bind.FilterOpts) (*RMNCursedIterator, error) {

	logs, sub, err := _RMN.contract.FilterLogs(opts, "Cursed")
	if err != nil {
		return nil, err
	}
	return &RMNCursedIterator{contract: _RMN.contract, event: "Cursed", logs: logs, sub: sub}, nil
}

func (_RMN *RMNFilterer) WatchCursed(opts *bind.WatchOpts, sink chan<- *RMNCursed) (event.Subscription, error) {

	logs, sub, err := _RMN.contract.WatchLogs(opts, "Cursed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNCursed)
				if err := _RMN.contract.UnpackLog(event, "Cursed", log); err != nil {
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

func (_RMN *RMNFilterer) ParseCursed(log types.Log) (*RMNCursed, error) {
	event := new(RMNCursed)
	if err := _RMN.contract.UnpackLog(event, "Cursed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNOwnershipTransferRequestedIterator struct {
	Event *RMNOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNOwnershipTransferRequested)
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
		it.Event = new(RMNOwnershipTransferRequested)
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

func (it *RMNOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *RMNOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RMN *RMNFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMN.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RMNOwnershipTransferRequestedIterator{contract: _RMN.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_RMN *RMNFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMN.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNOwnershipTransferRequested)
				if err := _RMN.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_RMN *RMNFilterer) ParseOwnershipTransferRequested(log types.Log) (*RMNOwnershipTransferRequested, error) {
	event := new(RMNOwnershipTransferRequested)
	if err := _RMN.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNOwnershipTransferredIterator struct {
	Event *RMNOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNOwnershipTransferred)
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
		it.Event = new(RMNOwnershipTransferred)
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

func (it *RMNOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *RMNOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_RMN *RMNFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMN.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RMNOwnershipTransferredIterator{contract: _RMN.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_RMN *RMNFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RMN.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNOwnershipTransferred)
				if err := _RMN.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_RMN *RMNFilterer) ParseOwnershipTransferred(log types.Log) (*RMNOwnershipTransferred, error) {
	event := new(RMNOwnershipTransferred)
	if err := _RMN.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type RMNUncursedIterator struct {
	Event *RMNUncursed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RMNUncursedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RMNUncursed)
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
		it.Event = new(RMNUncursed)
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

func (it *RMNUncursedIterator) Error() error {
	return it.fail
}

func (it *RMNUncursedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RMNUncursed struct {
	Subjects [][16]byte
	Raw      types.Log
}

func (_RMN *RMNFilterer) FilterUncursed(opts *bind.FilterOpts) (*RMNUncursedIterator, error) {

	logs, sub, err := _RMN.contract.FilterLogs(opts, "Uncursed")
	if err != nil {
		return nil, err
	}
	return &RMNUncursedIterator{contract: _RMN.contract, event: "Uncursed", logs: logs, sub: sub}, nil
}

func (_RMN *RMNFilterer) WatchUncursed(opts *bind.WatchOpts, sink chan<- *RMNUncursed) (event.Subscription, error) {

	logs, sub, err := _RMN.contract.WatchLogs(opts, "Uncursed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RMNUncursed)
				if err := _RMN.contract.UnpackLog(event, "Uncursed", log); err != nil {
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

func (_RMN *RMNFilterer) ParseUncursed(log types.Log) (*RMNUncursed, error) {
	event := new(RMNUncursed)
	if err := _RMN.contract.UnpackLog(event, "Uncursed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (RMNAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (RMNAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (RMNCursed) Topic() common.Hash {
	return common.HexToHash("0x1716e663a90a76d3b6c7e5f680673d1b051454c19c627e184c8daf28f3104f74")
}

func (RMNOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (RMNOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (RMNUncursed) Topic() common.Hash {
	return common.HexToHash("0x0676e709c9cc74fa0519fd78f7c33be0f1b2b0bae0507c724aef7229379c6ba1")
}

func (_RMN *RMN) Address() common.Address {
	return _RMN.address
}

type RMNInterface interface {
	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetCursedSubjects(opts *bind.CallOpts) ([][16]byte, error)

	IsCursed(opts *bind.CallOpts, subject [16]byte) (bool, error)

	IsCursed0(opts *bind.CallOpts) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	Curse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)

	Curse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Uncurse(opts *bind.TransactOpts, subject [16]byte) (*types.Transaction, error)

	Uncurse0(opts *bind.TransactOpts, subjects [][16]byte) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*RMNAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *RMNAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*RMNAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*RMNAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *RMNAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*RMNAuthorizedCallerRemoved, error)

	FilterCursed(opts *bind.FilterOpts) (*RMNCursedIterator, error)

	WatchCursed(opts *bind.WatchOpts, sink chan<- *RMNCursed) (event.Subscription, error)

	ParseCursed(log types.Log) (*RMNCursed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *RMNOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*RMNOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RMNOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RMNOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*RMNOwnershipTransferred, error)

	FilterUncursed(opts *bind.FilterOpts) (*RMNUncursedIterator, error)

	WatchUncursed(opts *bind.WatchOpts, sink chan<- *RMNUncursed) (event.Subscription, error)

	ParseUncursed(log types.Log) (*RMNUncursed, error)

	Address() common.Address
}
