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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
}

var CCTPMessageTransmitterProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenMessenger\",\"type\":\"address\",\"internalType\":\"contract ITokenMessenger\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_cctpTransmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IMessageTransmitter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransmitterCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a0806040523461026e57602081611147803803809161001f8285610273565b83398101031261026e57516001600160a01b0381169081900361026e5760405190602061004c8184610273565b600083526000368137331561025d57600180546001600160a01b0319163317905560405161007a8282610273565b60008152600036813760408051949085016001600160401b03811186821017610247576040528452808285015260005b8151811015610111576001906001600160a01b036100c88285610296565b5116846100d4826102d8565b6100e1575b5050016100aa565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138846100d9565b5050915160005b83825182101561018e57506001600160a01b036101358284610296565b5116801561017d5784917feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef838361016d6001956103d6565b50604051908152a1019050610118565b6342bcdf7f60e11b60005260046000fd5b604051632c12192160e01b81528181600481885afa91821561023b576000926101f7575b50506001600160a01b03166080819052156101e657604051610d10908161043782396080518181816101ba015261065a0152f35b6324acd18360e21b60005260046000fd5b81813d8311610234575b61020b8183610273565b810103126102305751906001600160a01b038216820361022d575081806101b2565b80fd5b5080fd5b503d610201565b6040513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b639b15e16f60e01b60005260046000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761024757604052565b80518210156102aa5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156102aa5760005260206000200190600090565b60008181526003602052604090205480156103cf5760001981018181116103b9576002546000198101919082116103b957808203610368575b5050506002548015610352576000190161032c8160026102c0565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6103a161037961038a9360026102c0565b90549060031b1c92839260026102c0565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610311565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461043057600254680100000000000000008110156102475761041761038a82600185940160025560026102c0565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c908163181f5a7714610800575080632451a6271461071257806357ecfd281461054357806379ba50971461045a5780638da5cb5b1461040857806391a2749a146101de578063cfc1db061461016f5763f2fde38b1461007757600080fd5b3461016a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043573ffffffffffffffffffffffffffffffffffffffff811680910361016a576100cf610a38565b33811461014057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461016a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043567ffffffffffffffff811161016a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261016a57604051906040820182811067ffffffffffffffff8211176103d957604052806004013567ffffffffffffffff811161016a5761028d9060043691840101610975565b825260248101359067ffffffffffffffff821161016a5760046102b39236920101610975565b602082019081526102c2610a38565b519060005b825181101561033a578073ffffffffffffffffffffffffffffffffffffffff6102f260019386610a83565b51166102fd81610ade565b610309575b50016102c7565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184610302565b505160005b81518110156103d75773ffffffffffffffffffffffffffffffffffffffff6103678284610a83565b51169081156103ad577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef60208361039f600195610ca3565b50604051908152a10161033f565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760005473ffffffffffffffffffffffffffffffffffffffff81163303610519577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a5760043567ffffffffffffffff811161016a57610592903690600401610947565b60243567ffffffffffffffff811161016a576105b2903690600401610947565b929091336000526003602052604060002054156106e45761063f60209361060f9560405196879586957f57ecfd280000000000000000000000000000000000000000000000000000000087526040600488015260448701916109f9565b917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8584030160248601526109f9565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156106d857600091610698575b6020826040519015158152f35b6020813d6020116106d0575b816106b160209383610906565b810103126106cc575180151581036106cc579050602061068b565b5080fd5b3d91506106a4565b6040513d6000823e3d90fd5b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a576040518060206002549283815201809260026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace9060005b8181106107ea5750505081610791910382610906565b6040519182916020830190602084525180915260408301919060005b8181106107bb575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016107ad565b825484526020909301926001928301920161077b565b3461016a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261016a576060810181811067ffffffffffffffff8211176103d957604052602581527f434354504d6573736167655472616e736d697474657250726f787920322e302e60208201527f302d646576000000000000000000000000000000000000000000000000000000604082015260405190602082528181519182602083015260005b8381106108ee5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b602082820181015160408784010152859350016108ae565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103d957604052565b9181601f8401121561016a5782359167ffffffffffffffff831161016a576020838186019501011161016a57565b81601f8201121561016a5780359167ffffffffffffffff83116103d9578260051b91604051936109a86020850186610906565b845260208085019382010191821161016a57602001915b8183106109cc5750505090565b823573ffffffffffffffffffffffffffffffffffffffff8116810361016a578152602092830192016109bf565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b73ffffffffffffffffffffffffffffffffffffffff600154163303610a5957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051821015610a975760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8054821015610a975760005260206000200190600090565b6000818152600360205260409020548015610c9c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111610c6d57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610c6d57808203610bfe575b5050506002548015610bcf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610b8c816002610ac6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b610c55610c0f610c20936002610ac6565b90549060031b1c9283926002610ac6565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080610b53565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610cfd57600254680100000000000000008110156103d957610ce4610c208260018594016002556002610ac6565b9055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
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

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPMessageTransmitterProxy.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _CCTPMessageTransmitterProxy.Contract.GetAllAuthorizedCallers(&_CCTPMessageTransmitterProxy.CallOpts)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _CCTPMessageTransmitterProxy.Contract.GetAllAuthorizedCallers(&_CCTPMessageTransmitterProxy.CallOpts)
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

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxySession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.ApplyAuthorizedCallerUpdates(&_CCTPMessageTransmitterProxy.TransactOpts, authorizedCallerArgs)
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPMessageTransmitterProxy.Contract.ApplyAuthorizedCallerUpdates(&_CCTPMessageTransmitterProxy.TransactOpts, authorizedCallerArgs)
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

type CCTPMessageTransmitterProxyAuthorizedCallerAddedIterator struct {
	Event *CCTPMessageTransmitterProxyAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPMessageTransmitterProxyAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPMessageTransmitterProxyAuthorizedCallerAdded)
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
		it.Event = new(CCTPMessageTransmitterProxyAuthorizedCallerAdded)
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

func (it *CCTPMessageTransmitterProxyAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPMessageTransmitterProxyAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPMessageTransmitterProxyAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*CCTPMessageTransmitterProxyAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyAuthorizedCallerAddedIterator{contract: _CCTPMessageTransmitterProxy.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPMessageTransmitterProxyAuthorizedCallerAdded)
				if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) ParseAuthorizedCallerAdded(log types.Log) (*CCTPMessageTransmitterProxyAuthorizedCallerAdded, error) {
	event := new(CCTPMessageTransmitterProxyAuthorizedCallerAdded)
	if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPMessageTransmitterProxyAuthorizedCallerRemovedIterator struct {
	Event *CCTPMessageTransmitterProxyAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPMessageTransmitterProxyAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPMessageTransmitterProxyAuthorizedCallerRemoved)
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
		it.Event = new(CCTPMessageTransmitterProxyAuthorizedCallerRemoved)
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

func (it *CCTPMessageTransmitterProxyAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPMessageTransmitterProxyAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPMessageTransmitterProxyAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*CCTPMessageTransmitterProxyAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &CCTPMessageTransmitterProxyAuthorizedCallerRemovedIterator{contract: _CCTPMessageTransmitterProxy.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _CCTPMessageTransmitterProxy.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPMessageTransmitterProxyAuthorizedCallerRemoved)
				if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_CCTPMessageTransmitterProxy *CCTPMessageTransmitterProxyFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*CCTPMessageTransmitterProxyAuthorizedCallerRemoved, error) {
	event := new(CCTPMessageTransmitterProxyAuthorizedCallerRemoved)
	if err := _CCTPMessageTransmitterProxy.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (CCTPMessageTransmitterProxyAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (CCTPMessageTransmitterProxyAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
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
	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	ICctpTransmitter(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*CCTPMessageTransmitterProxyAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*CCTPMessageTransmitterProxyAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*CCTPMessageTransmitterProxyAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*CCTPMessageTransmitterProxyAuthorizedCallerRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPMessageTransmitterProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPMessageTransmitterProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPMessageTransmitterProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPMessageTransmitterProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPMessageTransmitterProxyOwnershipTransferred, error)

	Address() common.Address
}
