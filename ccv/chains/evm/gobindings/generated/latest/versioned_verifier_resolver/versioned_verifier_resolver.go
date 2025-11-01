// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package versioned_verifier_resolver

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

type VersionedVerifierResolverInboundImplementationArgs struct {
	Version  [4]byte
	Verifier common.Address
}

type VersionedVerifierResolverOutboundImplementationArgs struct {
	DestChainSelector uint64
	Verifier          common.Address
}

var VersionedVerifierResolverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyInboundImplementationUpdates\",\"inputs\":[{\"name\":\"implementations\",\"type\":\"tuple[]\",\"internalType\":\"struct VersionedVerifierResolver.InboundImplementationArgs[]\",\"components\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyOutboundImplementationUpdates\",\"inputs\":[{\"name\":\"implementations\",\"type\":\"tuple[]\",\"internalType\":\"struct VersionedVerifierResolver.OutboundImplementationArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getInboundImplementation\",\"inputs\":[{\"name\":\"ccvData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInboundImplementationForVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutboundImplementation\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedDestChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedVerifierVersions\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"InboundImplementationRemoved\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundImplementationUpdated\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"prevImpl\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newImpl\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundImplementationRemoved\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundImplementationUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"prevImpl\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newImpl\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCCVDataLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainSelector\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x60808060405234603d573315602c57600180546001600160a01b0319163317905561137090816100438239f35b639b15e16f60e01b60005260046000fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c8063181f5a7714610d3f578063348759c114610c1b5780634c53709d14610a9557806379ba5097146109ac5780637a9c2ef9146107565780638da5cb5b14610704578063958021a7146105df578063b67987eb14610555578063c3a7ded61461045b578063e70769181461018d5763f2fde38b1461009557600080fd5b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760043573ffffffffffffffffffffffffffffffffffffffff8116809103610188576100ed610f56565b33811461015e57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760043567ffffffffffffffff8111610188576101dc903690600401610e99565b6101e4610f56565b60005b8181106101f057005b6101fb818385610f46565b6040813603126101885761020d610e35565b8135917fffffffff0000000000000000000000000000000000000000000000000000000083169283810361018857825261024990602001610eca565b9173ffffffffffffffffffffffffffffffffffffffff602083019380855216156103c357507fffffffff0000000000000000000000000000000000000000000000000000000081511680156103965750606060019392827fffffffff000000000000000000000000000000000000000000000000000000007f240744c957da89d5c44d43838bbc5553c6ec57314f9e62435f9158c45b4e3413945116600052600260205273ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000008160406000205416928285511682825116600052600260205283604060002091167fffffffffffffffffffffffff000000000000000000000000000000000000000082541617905561037782825116611309565b5051169251169060405192835260208301526040820152a15b016101e7565b7fa176027f0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60019392507fffffffff000000000000000000000000000000000000000000000000000000007f5dd8185b50a7df2c96bed0b91303df2507335646714c0d7896403165e4a58013926020926000526002835260406000207fffffffffffffffffffffffff0000000000000000000000000000000000000000815416905561044c8282511661117e565b505116604051908152a1610390565b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760043567ffffffffffffffff8111610188573660238201121561018857806004013567ffffffffffffffff8111610188573660248284010111610188576004811061052b576004116101885760247fffffffff00000000000000000000000000000000000000000000000000000000910135166000526002602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b7f2a3e9ae70000000000000000000000000000000000000000000000000000000060005260046000fd5b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610188576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610188576000526002602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346101885760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760043567ffffffffffffffff81168091036101885760243567ffffffffffffffff811161018857366023820112156101885780600401359067ffffffffffffffff82116106d55761068660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011601610e55565b9180835236602482840101116101885760009281602460209401848301370101526000526005602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760043567ffffffffffffffff8111610188576107a5903690600401610e99565b6107ad610f56565b60005b8181106107b957005b6107c4818385610f46565b604081360312610188576107d6610e35565b81359167ffffffffffffffff8316928381036101885782526107fa90602001610eca565b9173ffffffffffffffffffffffffffffffffffffffff6020830193808552161561092c575067ffffffffffffffff81511680156108ff57506060600193928267ffffffffffffffff7fc12b226506536cd62f34841a87d2333621e547ff4af0f3b13f3ac204bfb47ab1945116600052600560205273ffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff8160406000205416928285511682825116600052600560205283604060002091167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790556108e0828251166112a9565b5051169251169060405192835260208301526040820152a15b016107b0565b7fef75b4cf0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b600193925067ffffffffffffffff7f243416eecc562f47eb105155ee12ae26bb6e8dcbfce4c10e3ee69273e167214a926020926000526005835260406000207fffffffffffffffffffffffff0000000000000000000000000000000000000000815416905561099d82825116610fb9565b505116604051908152a16108f9565b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101885760005473ffffffffffffffffffffffffffffffffffffffff81163303610a6b577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018857600354610ad8610ad382610eeb565b610e55565b90808252610ae581610eeb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376003549060005b818110610b855783856040519182916020830190602084525180915260408301919060005b818110610b4a575050500390f35b82517fffffffff0000000000000000000000000000000000000000000000000000000016845285945060209384019390920191600101610b3c565b600083821015610bee57600390527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810154600191907fffffffff0000000000000000000000000000000000000000000000000000000016610be78288610f03565b5201610b17565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018857600654610c59610ad382610eeb565b90808252610c6681610eeb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376006549060005b818110610cee5783856040519182916020830190602084525180915260408301919060005b818110610ccb575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101610cbd565b600083821015610bee57600690527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f8101546001919067ffffffffffffffff16610d388288610f03565b5201610c98565b346101885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018857610d786060610e55565b602381527f56657273696f6e656456657269666965725265736f6c76657220312e372e302d60208201527f6465760000000000000000000000000000000000000000000000000000000000604082015260405190602082528181519182602083015260005b838110610e1d5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001610ddd565b604051906040820182811067ffffffffffffffff8211176106d557604052565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f604051930116820182811067ffffffffffffffff8211176106d557604052565b9181601f840112156101885782359167ffffffffffffffff8311610188576020808501948460061b01011161018857565b359073ffffffffffffffffffffffffffffffffffffffff8216820361018857565b67ffffffffffffffff81116106d55760051b60200190565b8051821015610f175760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015610f175760061b0190565b73ffffffffffffffffffffffffffffffffffffffff600154163303610f7757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8054821015610f175760005260206000200190600090565b6000818152600760205260409020548015611177577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161114857600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611148578181036110d9575b50505060065480156110aa577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611067816006610fa1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6111306110ea6110fb936006610fa1565b90549060031b1c9283926006610fa1565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600760205260406000205538808061102e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5050600090565b6000818152600460205260409020548015611177577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161114857600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116111485781810361126f575b50505060035480156110aa577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161122c816003610fa1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b6112916112806110fb936003610fa1565b90549060031b1c9283926003610fa1565b905560005260046020526040600020553880806111f3565b8060005260076020526040600020541560001461130357600654680100000000000000008110156106d5576112ea6110fb8260018594016006556006610fa1565b9055600654906000526007602052604060002055600190565b50600090565b8060005260046020526040600020541560001461130357600354680100000000000000008110156106d55761134a6110fb8260018594016003556003610fa1565b905560035490600052600460205260406000205560019056fea164736f6c634300081a000a",
}

var VersionedVerifierResolverABI = VersionedVerifierResolverMetaData.ABI

var VersionedVerifierResolverBin = VersionedVerifierResolverMetaData.Bin

func DeployVersionedVerifierResolver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VersionedVerifierResolver, error) {
	parsed, err := VersionedVerifierResolverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VersionedVerifierResolverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VersionedVerifierResolver{address: address, abi: *parsed, VersionedVerifierResolverCaller: VersionedVerifierResolverCaller{contract: contract}, VersionedVerifierResolverTransactor: VersionedVerifierResolverTransactor{contract: contract}, VersionedVerifierResolverFilterer: VersionedVerifierResolverFilterer{contract: contract}}, nil
}

type VersionedVerifierResolver struct {
	address common.Address
	abi     abi.ABI
	VersionedVerifierResolverCaller
	VersionedVerifierResolverTransactor
	VersionedVerifierResolverFilterer
}

type VersionedVerifierResolverCaller struct {
	contract *bind.BoundContract
}

type VersionedVerifierResolverTransactor struct {
	contract *bind.BoundContract
}

type VersionedVerifierResolverFilterer struct {
	contract *bind.BoundContract
}

type VersionedVerifierResolverSession struct {
	Contract     *VersionedVerifierResolver
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VersionedVerifierResolverCallerSession struct {
	Contract *VersionedVerifierResolverCaller
	CallOpts bind.CallOpts
}

type VersionedVerifierResolverTransactorSession struct {
	Contract     *VersionedVerifierResolverTransactor
	TransactOpts bind.TransactOpts
}

type VersionedVerifierResolverRaw struct {
	Contract *VersionedVerifierResolver
}

type VersionedVerifierResolverCallerRaw struct {
	Contract *VersionedVerifierResolverCaller
}

type VersionedVerifierResolverTransactorRaw struct {
	Contract *VersionedVerifierResolverTransactor
}

func NewVersionedVerifierResolver(address common.Address, backend bind.ContractBackend) (*VersionedVerifierResolver, error) {
	abi, err := abi.JSON(strings.NewReader(VersionedVerifierResolverABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVersionedVerifierResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolver{address: address, abi: abi, VersionedVerifierResolverCaller: VersionedVerifierResolverCaller{contract: contract}, VersionedVerifierResolverTransactor: VersionedVerifierResolverTransactor{contract: contract}, VersionedVerifierResolverFilterer: VersionedVerifierResolverFilterer{contract: contract}}, nil
}

func NewVersionedVerifierResolverCaller(address common.Address, caller bind.ContractCaller) (*VersionedVerifierResolverCaller, error) {
	contract, err := bindVersionedVerifierResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverCaller{contract: contract}, nil
}

func NewVersionedVerifierResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionedVerifierResolverTransactor, error) {
	contract, err := bindVersionedVerifierResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverTransactor{contract: contract}, nil
}

func NewVersionedVerifierResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionedVerifierResolverFilterer, error) {
	contract, err := bindVersionedVerifierResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverFilterer{contract: contract}, nil
}

func bindVersionedVerifierResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VersionedVerifierResolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VersionedVerifierResolver.Contract.VersionedVerifierResolverCaller.contract.Call(opts, result, method, params...)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.VersionedVerifierResolverTransactor.contract.Transfer(opts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.VersionedVerifierResolverTransactor.contract.Transact(opts, method, params...)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VersionedVerifierResolver.Contract.contract.Call(opts, result, method, params...)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.contract.Transfer(opts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.contract.Transact(opts, method, params...)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetInboundImplementation(opts *bind.CallOpts, ccvData []byte) (common.Address, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getInboundImplementation", ccvData)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetInboundImplementation(ccvData []byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetInboundImplementation(&_VersionedVerifierResolver.CallOpts, ccvData)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetInboundImplementation(ccvData []byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetInboundImplementation(&_VersionedVerifierResolver.CallOpts, ccvData)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetInboundImplementationForVersion(opts *bind.CallOpts, version [4]byte) (common.Address, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getInboundImplementationForVersion", version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetInboundImplementationForVersion(version [4]byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetInboundImplementationForVersion(&_VersionedVerifierResolver.CallOpts, version)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetInboundImplementationForVersion(version [4]byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetInboundImplementationForVersion(&_VersionedVerifierResolver.CallOpts, version)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetOutboundImplementation(opts *bind.CallOpts, destChainSelector uint64, arg1 []byte) (common.Address, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getOutboundImplementation", destChainSelector, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetOutboundImplementation(destChainSelector uint64, arg1 []byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetOutboundImplementation(&_VersionedVerifierResolver.CallOpts, destChainSelector, arg1)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetOutboundImplementation(destChainSelector uint64, arg1 []byte) (common.Address, error) {
	return _VersionedVerifierResolver.Contract.GetOutboundImplementation(&_VersionedVerifierResolver.CallOpts, destChainSelector, arg1)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetSupportedDestChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getSupportedDestChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetSupportedDestChains() ([]uint64, error) {
	return _VersionedVerifierResolver.Contract.GetSupportedDestChains(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetSupportedDestChains() ([]uint64, error) {
	return _VersionedVerifierResolver.Contract.GetSupportedDestChains(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) GetSupportedVerifierVersions(opts *bind.CallOpts) ([][4]byte, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "getSupportedVerifierVersions")

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) GetSupportedVerifierVersions() ([][4]byte, error) {
	return _VersionedVerifierResolver.Contract.GetSupportedVerifierVersions(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) GetSupportedVerifierVersions() ([][4]byte, error) {
	return _VersionedVerifierResolver.Contract.GetSupportedVerifierVersions(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) Owner() (common.Address, error) {
	return _VersionedVerifierResolver.Contract.Owner(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) Owner() (common.Address, error) {
	return _VersionedVerifierResolver.Contract.Owner(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VersionedVerifierResolver.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) TypeAndVersion() (string, error) {
	return _VersionedVerifierResolver.Contract.TypeAndVersion(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverCallerSession) TypeAndVersion() (string, error) {
	return _VersionedVerifierResolver.Contract.TypeAndVersion(&_VersionedVerifierResolver.CallOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "acceptOwnership")
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) AcceptOwnership() (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.AcceptOwnership(&_VersionedVerifierResolver.TransactOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.AcceptOwnership(&_VersionedVerifierResolver.TransactOpts)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) ApplyInboundImplementationUpdates(opts *bind.TransactOpts, implementations []VersionedVerifierResolverInboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "applyInboundImplementationUpdates", implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) ApplyInboundImplementationUpdates(implementations []VersionedVerifierResolverInboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.ApplyInboundImplementationUpdates(&_VersionedVerifierResolver.TransactOpts, implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) ApplyInboundImplementationUpdates(implementations []VersionedVerifierResolverInboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.ApplyInboundImplementationUpdates(&_VersionedVerifierResolver.TransactOpts, implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) ApplyOutboundImplementationUpdates(opts *bind.TransactOpts, implementations []VersionedVerifierResolverOutboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "applyOutboundImplementationUpdates", implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) ApplyOutboundImplementationUpdates(implementations []VersionedVerifierResolverOutboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.ApplyOutboundImplementationUpdates(&_VersionedVerifierResolver.TransactOpts, implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) ApplyOutboundImplementationUpdates(implementations []VersionedVerifierResolverOutboundImplementationArgs) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.ApplyOutboundImplementationUpdates(&_VersionedVerifierResolver.TransactOpts, implementations)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.contract.Transact(opts, "transferOwnership", to)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.TransferOwnership(&_VersionedVerifierResolver.TransactOpts, to)
}

func (_VersionedVerifierResolver *VersionedVerifierResolverTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VersionedVerifierResolver.Contract.TransferOwnership(&_VersionedVerifierResolver.TransactOpts, to)
}

type VersionedVerifierResolverInboundImplementationRemovedIterator struct {
	Event *VersionedVerifierResolverInboundImplementationRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverInboundImplementationRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverInboundImplementationRemoved)
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
		it.Event = new(VersionedVerifierResolverInboundImplementationRemoved)
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

func (it *VersionedVerifierResolverInboundImplementationRemovedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverInboundImplementationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverInboundImplementationRemoved struct {
	Version [4]byte
	Raw     types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterInboundImplementationRemoved(opts *bind.FilterOpts) (*VersionedVerifierResolverInboundImplementationRemovedIterator, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "InboundImplementationRemoved")
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverInboundImplementationRemovedIterator{contract: _VersionedVerifierResolver.contract, event: "InboundImplementationRemoved", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchInboundImplementationRemoved(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverInboundImplementationRemoved) (event.Subscription, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "InboundImplementationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverInboundImplementationRemoved)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "InboundImplementationRemoved", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseInboundImplementationRemoved(log types.Log) (*VersionedVerifierResolverInboundImplementationRemoved, error) {
	event := new(VersionedVerifierResolverInboundImplementationRemoved)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "InboundImplementationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverInboundImplementationUpdatedIterator struct {
	Event *VersionedVerifierResolverInboundImplementationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverInboundImplementationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverInboundImplementationUpdated)
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
		it.Event = new(VersionedVerifierResolverInboundImplementationUpdated)
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

func (it *VersionedVerifierResolverInboundImplementationUpdatedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverInboundImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverInboundImplementationUpdated struct {
	Version  [4]byte
	PrevImpl common.Address
	NewImpl  common.Address
	Raw      types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterInboundImplementationUpdated(opts *bind.FilterOpts) (*VersionedVerifierResolverInboundImplementationUpdatedIterator, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "InboundImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverInboundImplementationUpdatedIterator{contract: _VersionedVerifierResolver.contract, event: "InboundImplementationUpdated", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchInboundImplementationUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverInboundImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "InboundImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverInboundImplementationUpdated)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "InboundImplementationUpdated", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseInboundImplementationUpdated(log types.Log) (*VersionedVerifierResolverInboundImplementationUpdated, error) {
	event := new(VersionedVerifierResolverInboundImplementationUpdated)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "InboundImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverOutboundImplementationRemovedIterator struct {
	Event *VersionedVerifierResolverOutboundImplementationRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverOutboundImplementationRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverOutboundImplementationRemoved)
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
		it.Event = new(VersionedVerifierResolverOutboundImplementationRemoved)
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

func (it *VersionedVerifierResolverOutboundImplementationRemovedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverOutboundImplementationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverOutboundImplementationRemoved struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterOutboundImplementationRemoved(opts *bind.FilterOpts) (*VersionedVerifierResolverOutboundImplementationRemovedIterator, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "OutboundImplementationRemoved")
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverOutboundImplementationRemovedIterator{contract: _VersionedVerifierResolver.contract, event: "OutboundImplementationRemoved", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchOutboundImplementationRemoved(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOutboundImplementationRemoved) (event.Subscription, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "OutboundImplementationRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverOutboundImplementationRemoved)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OutboundImplementationRemoved", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseOutboundImplementationRemoved(log types.Log) (*VersionedVerifierResolverOutboundImplementationRemoved, error) {
	event := new(VersionedVerifierResolverOutboundImplementationRemoved)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OutboundImplementationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverOutboundImplementationUpdatedIterator struct {
	Event *VersionedVerifierResolverOutboundImplementationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverOutboundImplementationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverOutboundImplementationUpdated)
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
		it.Event = new(VersionedVerifierResolverOutboundImplementationUpdated)
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

func (it *VersionedVerifierResolverOutboundImplementationUpdatedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverOutboundImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverOutboundImplementationUpdated struct {
	DestChainSelector uint64
	PrevImpl          common.Address
	NewImpl           common.Address
	Raw               types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterOutboundImplementationUpdated(opts *bind.FilterOpts) (*VersionedVerifierResolverOutboundImplementationUpdatedIterator, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "OutboundImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverOutboundImplementationUpdatedIterator{contract: _VersionedVerifierResolver.contract, event: "OutboundImplementationUpdated", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchOutboundImplementationUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOutboundImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "OutboundImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverOutboundImplementationUpdated)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OutboundImplementationUpdated", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseOutboundImplementationUpdated(log types.Log) (*VersionedVerifierResolverOutboundImplementationUpdated, error) {
	event := new(VersionedVerifierResolverOutboundImplementationUpdated)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OutboundImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverOwnershipTransferRequestedIterator struct {
	Event *VersionedVerifierResolverOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverOwnershipTransferRequested)
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
		it.Event = new(VersionedVerifierResolverOwnershipTransferRequested)
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

func (it *VersionedVerifierResolverOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VersionedVerifierResolverOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverOwnershipTransferRequestedIterator{contract: _VersionedVerifierResolver.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverOwnershipTransferRequested)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseOwnershipTransferRequested(log types.Log) (*VersionedVerifierResolverOwnershipTransferRequested, error) {
	event := new(VersionedVerifierResolverOwnershipTransferRequested)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VersionedVerifierResolverOwnershipTransferredIterator struct {
	Event *VersionedVerifierResolverOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VersionedVerifierResolverOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VersionedVerifierResolverOwnershipTransferred)
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
		it.Event = new(VersionedVerifierResolverOwnershipTransferred)
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

func (it *VersionedVerifierResolverOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VersionedVerifierResolverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VersionedVerifierResolverOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VersionedVerifierResolverOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VersionedVerifierResolverOwnershipTransferredIterator{contract: _VersionedVerifierResolver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VersionedVerifierResolver.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VersionedVerifierResolverOwnershipTransferred)
				if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VersionedVerifierResolver *VersionedVerifierResolverFilterer) ParseOwnershipTransferred(log types.Log) (*VersionedVerifierResolverOwnershipTransferred, error) {
	event := new(VersionedVerifierResolverOwnershipTransferred)
	if err := _VersionedVerifierResolver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (VersionedVerifierResolverInboundImplementationRemoved) Topic() common.Hash {
	return common.HexToHash("0x5dd8185b50a7df2c96bed0b91303df2507335646714c0d7896403165e4a58013")
}

func (VersionedVerifierResolverInboundImplementationUpdated) Topic() common.Hash {
	return common.HexToHash("0x240744c957da89d5c44d43838bbc5553c6ec57314f9e62435f9158c45b4e3413")
}

func (VersionedVerifierResolverOutboundImplementationRemoved) Topic() common.Hash {
	return common.HexToHash("0x243416eecc562f47eb105155ee12ae26bb6e8dcbfce4c10e3ee69273e167214a")
}

func (VersionedVerifierResolverOutboundImplementationUpdated) Topic() common.Hash {
	return common.HexToHash("0xc12b226506536cd62f34841a87d2333621e547ff4af0f3b13f3ac204bfb47ab1")
}

func (VersionedVerifierResolverOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VersionedVerifierResolverOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VersionedVerifierResolver *VersionedVerifierResolver) Address() common.Address {
	return _VersionedVerifierResolver.address
}

type VersionedVerifierResolverInterface interface {
	GetInboundImplementation(opts *bind.CallOpts, ccvData []byte) (common.Address, error)

	GetInboundImplementationForVersion(opts *bind.CallOpts, version [4]byte) (common.Address, error)

	GetOutboundImplementation(opts *bind.CallOpts, destChainSelector uint64, arg1 []byte) (common.Address, error)

	GetSupportedDestChains(opts *bind.CallOpts) ([]uint64, error)

	GetSupportedVerifierVersions(opts *bind.CallOpts) ([][4]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyInboundImplementationUpdates(opts *bind.TransactOpts, implementations []VersionedVerifierResolverInboundImplementationArgs) (*types.Transaction, error)

	ApplyOutboundImplementationUpdates(opts *bind.TransactOpts, implementations []VersionedVerifierResolverOutboundImplementationArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterInboundImplementationRemoved(opts *bind.FilterOpts) (*VersionedVerifierResolverInboundImplementationRemovedIterator, error)

	WatchInboundImplementationRemoved(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverInboundImplementationRemoved) (event.Subscription, error)

	ParseInboundImplementationRemoved(log types.Log) (*VersionedVerifierResolverInboundImplementationRemoved, error)

	FilterInboundImplementationUpdated(opts *bind.FilterOpts) (*VersionedVerifierResolverInboundImplementationUpdatedIterator, error)

	WatchInboundImplementationUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverInboundImplementationUpdated) (event.Subscription, error)

	ParseInboundImplementationUpdated(log types.Log) (*VersionedVerifierResolverInboundImplementationUpdated, error)

	FilterOutboundImplementationRemoved(opts *bind.FilterOpts) (*VersionedVerifierResolverOutboundImplementationRemovedIterator, error)

	WatchOutboundImplementationRemoved(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOutboundImplementationRemoved) (event.Subscription, error)

	ParseOutboundImplementationRemoved(log types.Log) (*VersionedVerifierResolverOutboundImplementationRemoved, error)

	FilterOutboundImplementationUpdated(opts *bind.FilterOpts) (*VersionedVerifierResolverOutboundImplementationUpdatedIterator, error)

	WatchOutboundImplementationUpdated(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOutboundImplementationUpdated) (event.Subscription, error)

	ParseOutboundImplementationUpdated(log types.Log) (*VersionedVerifierResolverOutboundImplementationUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VersionedVerifierResolverOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VersionedVerifierResolverOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VersionedVerifierResolverOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VersionedVerifierResolverOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VersionedVerifierResolverOwnershipTransferred, error)

	Address() common.Address
}
