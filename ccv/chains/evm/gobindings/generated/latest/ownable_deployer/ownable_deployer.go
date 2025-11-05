// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownable_deployer

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

var OwnableDeployerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"computeAddress\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployAndTransferOwnership\",\"inputs\":[{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x608080604052346015576104bf908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b6000803560e01c8062d762e0146102c9578063181f5a77146101fb57638fdc344b1461003d57600080fd5b346101f85760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f85760043567ffffffffffffffff81116101a45761008c903690600401610438565b6040516020810190610101816100d53360243586906034927fffffffffffffffffffffffffffffffffffffffff00000000000000000000000091835260601b1660208201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103c8565b519020908051156101d057805173ffffffffffffffffffffffffffffffffffffffff929160200184f51680156101a857803b156101a457604051917ff2fde38b000000000000000000000000000000000000000000000000000000008352336004840152808360248183865af192831561019757602093610187575b5050604051908152f35b81610191916103c8565b3861017d565b50604051903d90823e3d90fd5b5080fd5b6004827f741752c2000000000000000000000000000000000000000000000000000000008152fd5b6004837f4ca249dc000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346101f857807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f857604080519161023982846103c8565b601b83527f44657465726d696e69737469634465706c6f79657220312e372e3000000000006020840152815192839160208352815191826020850152815b8381106102b2575050828201840152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168101030190f35b602082820181015188830188015287955001610277565b50346101f85760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101f8576004359073ffffffffffffffffffffffffffffffffffffffff821682036101f8576024359067ffffffffffffffff82116101f85760206055600b856103413660048801610438565b90604051610389816100d58882019460443586906034927fffffffffffffffffffffffffffffffffffffffff00000000000000000000000091835260601b1660208201520190565b5190209084815191012090604051916040830152848201523081520160ff81532073ffffffffffffffffffffffffffffffffffffffff60405191168152f35b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761040957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b81601f820112156104ad5780359067ffffffffffffffff8211610409576040519261048b601f84017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001856103c8565b828452602083830101116104ad57816000926020809301838601378301015290565b600080fdfea164736f6c634300081a000a",
}

var OwnableDeployerABI = OwnableDeployerMetaData.ABI

var OwnableDeployerBin = OwnableDeployerMetaData.Bin

func DeployOwnableDeployer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OwnableDeployer, error) {
	parsed, err := OwnableDeployerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnableDeployerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OwnableDeployer{address: address, abi: *parsed, OwnableDeployerCaller: OwnableDeployerCaller{contract: contract}, OwnableDeployerTransactor: OwnableDeployerTransactor{contract: contract}, OwnableDeployerFilterer: OwnableDeployerFilterer{contract: contract}}, nil
}

type OwnableDeployer struct {
	address common.Address
	abi     abi.ABI
	OwnableDeployerCaller
	OwnableDeployerTransactor
	OwnableDeployerFilterer
}

type OwnableDeployerCaller struct {
	contract *bind.BoundContract
}

type OwnableDeployerTransactor struct {
	contract *bind.BoundContract
}

type OwnableDeployerFilterer struct {
	contract *bind.BoundContract
}

type OwnableDeployerSession struct {
	Contract     *OwnableDeployer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OwnableDeployerCallerSession struct {
	Contract *OwnableDeployerCaller
	CallOpts bind.CallOpts
}

type OwnableDeployerTransactorSession struct {
	Contract     *OwnableDeployerTransactor
	TransactOpts bind.TransactOpts
}

type OwnableDeployerRaw struct {
	Contract *OwnableDeployer
}

type OwnableDeployerCallerRaw struct {
	Contract *OwnableDeployerCaller
}

type OwnableDeployerTransactorRaw struct {
	Contract *OwnableDeployerTransactor
}

func NewOwnableDeployer(address common.Address, backend bind.ContractBackend) (*OwnableDeployer, error) {
	abi, err := abi.JSON(strings.NewReader(OwnableDeployerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOwnableDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableDeployer{address: address, abi: abi, OwnableDeployerCaller: OwnableDeployerCaller{contract: contract}, OwnableDeployerTransactor: OwnableDeployerTransactor{contract: contract}, OwnableDeployerFilterer: OwnableDeployerFilterer{contract: contract}}, nil
}

func NewOwnableDeployerCaller(address common.Address, caller bind.ContractCaller) (*OwnableDeployerCaller, error) {
	contract, err := bindOwnableDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableDeployerCaller{contract: contract}, nil
}

func NewOwnableDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableDeployerTransactor, error) {
	contract, err := bindOwnableDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableDeployerTransactor{contract: contract}, nil
}

func NewOwnableDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableDeployerFilterer, error) {
	contract, err := bindOwnableDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableDeployerFilterer{contract: contract}, nil
}

func bindOwnableDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableDeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OwnableDeployer *OwnableDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableDeployer.Contract.OwnableDeployerCaller.contract.Call(opts, result, method, params...)
}

func (_OwnableDeployer *OwnableDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableDeployer.Contract.OwnableDeployerTransactor.contract.Transfer(opts)
}

func (_OwnableDeployer *OwnableDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableDeployer.Contract.OwnableDeployerTransactor.contract.Transact(opts, method, params...)
}

func (_OwnableDeployer *OwnableDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableDeployer.Contract.contract.Call(opts, result, method, params...)
}

func (_OwnableDeployer *OwnableDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableDeployer.Contract.contract.Transfer(opts)
}

func (_OwnableDeployer *OwnableDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableDeployer.Contract.contract.Transact(opts, method, params...)
}

func (_OwnableDeployer *OwnableDeployerCaller) ComputeAddress(opts *bind.CallOpts, sender common.Address, initCode []byte, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _OwnableDeployer.contract.Call(opts, &out, "computeAddress", sender, initCode, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OwnableDeployer *OwnableDeployerSession) ComputeAddress(sender common.Address, initCode []byte, salt [32]byte) (common.Address, error) {
	return _OwnableDeployer.Contract.ComputeAddress(&_OwnableDeployer.CallOpts, sender, initCode, salt)
}

func (_OwnableDeployer *OwnableDeployerCallerSession) ComputeAddress(sender common.Address, initCode []byte, salt [32]byte) (common.Address, error) {
	return _OwnableDeployer.Contract.ComputeAddress(&_OwnableDeployer.CallOpts, sender, initCode, salt)
}

func (_OwnableDeployer *OwnableDeployerCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OwnableDeployer.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OwnableDeployer *OwnableDeployerSession) TypeAndVersion() (string, error) {
	return _OwnableDeployer.Contract.TypeAndVersion(&_OwnableDeployer.CallOpts)
}

func (_OwnableDeployer *OwnableDeployerCallerSession) TypeAndVersion() (string, error) {
	return _OwnableDeployer.Contract.TypeAndVersion(&_OwnableDeployer.CallOpts)
}

func (_OwnableDeployer *OwnableDeployerTransactor) DeployAndTransferOwnership(opts *bind.TransactOpts, initCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _OwnableDeployer.contract.Transact(opts, "deployAndTransferOwnership", initCode, salt)
}

func (_OwnableDeployer *OwnableDeployerSession) DeployAndTransferOwnership(initCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _OwnableDeployer.Contract.DeployAndTransferOwnership(&_OwnableDeployer.TransactOpts, initCode, salt)
}

func (_OwnableDeployer *OwnableDeployerTransactorSession) DeployAndTransferOwnership(initCode []byte, salt [32]byte) (*types.Transaction, error) {
	return _OwnableDeployer.Contract.DeployAndTransferOwnership(&_OwnableDeployer.TransactOpts, initCode, salt)
}

func (_OwnableDeployer *OwnableDeployer) Address() common.Address {
	return _OwnableDeployer.address
}

type OwnableDeployerInterface interface {
	ComputeAddress(opts *bind.CallOpts, sender common.Address, initCode []byte, salt [32]byte) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	DeployAndTransferOwnership(opts *bind.TransactOpts, initCode []byte, salt [32]byte) (*types.Transaction, error)

	Address() common.Address
}
