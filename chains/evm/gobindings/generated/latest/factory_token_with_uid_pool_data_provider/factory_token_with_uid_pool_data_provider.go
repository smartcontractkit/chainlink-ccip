// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory_token_with_uid_pool_data_provider

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

var FactoryTokenWithUIDPoolDataProviderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenFactory\",\"type\":\"address\",\"internalType\":\"contract IResolveTokenAddressFromUID\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdditionalPoolData\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"additionalPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"tokenFactory\",\"type\":\"address\",\"internalType\":\"contract IResolveTokenAddressFromUID\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAdditionalPoolDataValid\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"additionalPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"isValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"InvalidTokenFactory\",\"inputs\":[{\"name\":\"tokenFactory\",\"type\":\"address\",\"internalType\":\"contract IResolveTokenAddressFromUID\"}]},{\"type\":\"error\",\"name\":\"TokenNotFromFactory\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uid\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a034608f57601f6106eb38819003918201601f19168301916001600160401b03831184841017609457808492602094604052833981010312608f57516001600160a01b038116808203608f57803b15607b575060805260405161064090816100ab823960805181818160fc01528181610272015261059a0152f35b6390b6e1e360e01b60005260045260246000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c908163181f5a77146102bd575080639219fcd714610296578063925d6ba2146102275763e4fa169c1461004b57600080fd5b346102225773ffffffffffffffffffffffffffffffffffffffff61006e366103fd565b5050505016604051907ff514ce36000000000000000000000000000000000000000000000000000000008252602082600481845afa9182156101df576000926101eb575b506040517f97bb3ce900000000000000000000000000000000000000000000000000000000815282600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156101df57829173ffffffffffffffffffffffffffffffffffffffff916000916101b0575b5016036101805761017c82604051906020820152602081526101686040826104ce565b60405191829160208352602083019061039e565b0390f35b7fea774b900000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b6101d2915060203d6020116101d8575b6101ca81836104ce565b81019061050f565b38610145565b503d6101c0565b6040513d6000823e3d90fd5b90916020823d60201161021a575b81610206602093836104ce565b8101031261021757505190386100b2565b80fd5b3d91506101f9565b600080fd5b346102225760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022257602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102225760206102b36102a9366103fd565b939150915061053b565b6040519015158152f35b346102225760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022257606081019080821067ffffffffffffffff83111761036f5761017c91604052602d81527f466163746f7279546f6b656e57697468554944506f6f6c4461746150726f766960208201527f64657220322e312e302d64657600000000000000000000000000000000000000604082015260405191829160208352602083019061039e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b919082519283825260005b8481106103e85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016103a9565b60807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102225760043573ffffffffffffffffffffffffffffffffffffffff81168103610222579160243567ffffffffffffffff8116810361022257916044357fffffffff0000000000000000000000000000000000000000000000000000000081168103610222579160643567ffffffffffffffff811161022257826023820112156102225780600401359267ffffffffffffffff84116102225760248483010111610222576024019190565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761036f57604052565b90816020910312610222575173ffffffffffffffffffffffffffffffffffffffff811681036102225790565b916020810361062b57816020918101031261022257604051907f97bb3ce900000000000000000000000000000000000000000000000000000000825235600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156101df5773ffffffffffffffffffffffffffffffffffffffff9160009161060c575b501680151591826105f057505090565b73ffffffffffffffffffffffffffffffffffffffff1614919050565b610625915060203d6020116101d8576101ca81836104ce565b386105e0565b50505060009056fea164736f6c634300081a000a",
}

var FactoryTokenWithUIDPoolDataProviderABI = FactoryTokenWithUIDPoolDataProviderMetaData.ABI

var FactoryTokenWithUIDPoolDataProviderBin = FactoryTokenWithUIDPoolDataProviderMetaData.Bin

func DeployFactoryTokenWithUIDPoolDataProvider(auth *bind.TransactOpts, backend bind.ContractBackend, tokenFactory common.Address) (common.Address, *types.Transaction, *FactoryTokenWithUIDPoolDataProvider, error) {
	parsed, err := FactoryTokenWithUIDPoolDataProviderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FactoryTokenWithUIDPoolDataProviderBin), backend, tokenFactory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FactoryTokenWithUIDPoolDataProvider{address: address, abi: *parsed, FactoryTokenWithUIDPoolDataProviderCaller: FactoryTokenWithUIDPoolDataProviderCaller{contract: contract}, FactoryTokenWithUIDPoolDataProviderTransactor: FactoryTokenWithUIDPoolDataProviderTransactor{contract: contract}, FactoryTokenWithUIDPoolDataProviderFilterer: FactoryTokenWithUIDPoolDataProviderFilterer{contract: contract}}, nil
}

type FactoryTokenWithUIDPoolDataProvider struct {
	address common.Address
	abi     abi.ABI
	FactoryTokenWithUIDPoolDataProviderCaller
	FactoryTokenWithUIDPoolDataProviderTransactor
	FactoryTokenWithUIDPoolDataProviderFilterer
}

type FactoryTokenWithUIDPoolDataProviderCaller struct {
	contract *bind.BoundContract
}

type FactoryTokenWithUIDPoolDataProviderTransactor struct {
	contract *bind.BoundContract
}

type FactoryTokenWithUIDPoolDataProviderFilterer struct {
	contract *bind.BoundContract
}

type FactoryTokenWithUIDPoolDataProviderSession struct {
	Contract     *FactoryTokenWithUIDPoolDataProvider
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type FactoryTokenWithUIDPoolDataProviderCallerSession struct {
	Contract *FactoryTokenWithUIDPoolDataProviderCaller
	CallOpts bind.CallOpts
}

type FactoryTokenWithUIDPoolDataProviderTransactorSession struct {
	Contract     *FactoryTokenWithUIDPoolDataProviderTransactor
	TransactOpts bind.TransactOpts
}

type FactoryTokenWithUIDPoolDataProviderRaw struct {
	Contract *FactoryTokenWithUIDPoolDataProvider
}

type FactoryTokenWithUIDPoolDataProviderCallerRaw struct {
	Contract *FactoryTokenWithUIDPoolDataProviderCaller
}

type FactoryTokenWithUIDPoolDataProviderTransactorRaw struct {
	Contract *FactoryTokenWithUIDPoolDataProviderTransactor
}

func NewFactoryTokenWithUIDPoolDataProvider(address common.Address, backend bind.ContractBackend) (*FactoryTokenWithUIDPoolDataProvider, error) {
	abi, err := abi.JSON(strings.NewReader(FactoryTokenWithUIDPoolDataProviderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindFactoryTokenWithUIDPoolDataProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FactoryTokenWithUIDPoolDataProvider{address: address, abi: abi, FactoryTokenWithUIDPoolDataProviderCaller: FactoryTokenWithUIDPoolDataProviderCaller{contract: contract}, FactoryTokenWithUIDPoolDataProviderTransactor: FactoryTokenWithUIDPoolDataProviderTransactor{contract: contract}, FactoryTokenWithUIDPoolDataProviderFilterer: FactoryTokenWithUIDPoolDataProviderFilterer{contract: contract}}, nil
}

func NewFactoryTokenWithUIDPoolDataProviderCaller(address common.Address, caller bind.ContractCaller) (*FactoryTokenWithUIDPoolDataProviderCaller, error) {
	contract, err := bindFactoryTokenWithUIDPoolDataProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTokenWithUIDPoolDataProviderCaller{contract: contract}, nil
}

func NewFactoryTokenWithUIDPoolDataProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTokenWithUIDPoolDataProviderTransactor, error) {
	contract, err := bindFactoryTokenWithUIDPoolDataProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTokenWithUIDPoolDataProviderTransactor{contract: contract}, nil
}

func NewFactoryTokenWithUIDPoolDataProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryTokenWithUIDPoolDataProviderFilterer, error) {
	contract, err := bindFactoryTokenWithUIDPoolDataProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryTokenWithUIDPoolDataProviderFilterer{contract: contract}, nil
}

func bindFactoryTokenWithUIDPoolDataProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryTokenWithUIDPoolDataProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.FactoryTokenWithUIDPoolDataProviderCaller.contract.Call(opts, result, method, params...)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.FactoryTokenWithUIDPoolDataProviderTransactor.contract.Transfer(opts)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.FactoryTokenWithUIDPoolDataProviderTransactor.contract.Transact(opts, method, params...)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.contract.Call(opts, result, method, params...)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.contract.Transfer(opts)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.contract.Transact(opts, method, params...)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCaller) GetAdditionalPoolData(opts *bind.CallOpts, localToken common.Address, arg1 uint64, arg2 [4]byte, arg3 []byte) ([]byte, error) {
	var out []interface{}
	err := _FactoryTokenWithUIDPoolDataProvider.contract.Call(opts, &out, "getAdditionalPoolData", localToken, arg1, arg2, arg3)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderSession) GetAdditionalPoolData(localToken common.Address, arg1 uint64, arg2 [4]byte, arg3 []byte) ([]byte, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.GetAdditionalPoolData(&_FactoryTokenWithUIDPoolDataProvider.CallOpts, localToken, arg1, arg2, arg3)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCallerSession) GetAdditionalPoolData(localToken common.Address, arg1 uint64, arg2 [4]byte, arg3 []byte) ([]byte, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.GetAdditionalPoolData(&_FactoryTokenWithUIDPoolDataProvider.CallOpts, localToken, arg1, arg2, arg3)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCaller) GetTokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FactoryTokenWithUIDPoolDataProvider.contract.Call(opts, &out, "getTokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderSession) GetTokenFactory() (common.Address, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.GetTokenFactory(&_FactoryTokenWithUIDPoolDataProvider.CallOpts)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCallerSession) GetTokenFactory() (common.Address, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.GetTokenFactory(&_FactoryTokenWithUIDPoolDataProvider.CallOpts)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCaller) IsAdditionalPoolDataValid(opts *bind.CallOpts, localToken common.Address, arg1 uint64, arg2 [4]byte, additionalPoolData []byte) (bool, error) {
	var out []interface{}
	err := _FactoryTokenWithUIDPoolDataProvider.contract.Call(opts, &out, "isAdditionalPoolDataValid", localToken, arg1, arg2, additionalPoolData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderSession) IsAdditionalPoolDataValid(localToken common.Address, arg1 uint64, arg2 [4]byte, additionalPoolData []byte) (bool, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.IsAdditionalPoolDataValid(&_FactoryTokenWithUIDPoolDataProvider.CallOpts, localToken, arg1, arg2, additionalPoolData)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCallerSession) IsAdditionalPoolDataValid(localToken common.Address, arg1 uint64, arg2 [4]byte, additionalPoolData []byte) (bool, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.IsAdditionalPoolDataValid(&_FactoryTokenWithUIDPoolDataProvider.CallOpts, localToken, arg1, arg2, additionalPoolData)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FactoryTokenWithUIDPoolDataProvider.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderSession) TypeAndVersion() (string, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.TypeAndVersion(&_FactoryTokenWithUIDPoolDataProvider.CallOpts)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProviderCallerSession) TypeAndVersion() (string, error) {
	return _FactoryTokenWithUIDPoolDataProvider.Contract.TypeAndVersion(&_FactoryTokenWithUIDPoolDataProvider.CallOpts)
}

func (_FactoryTokenWithUIDPoolDataProvider *FactoryTokenWithUIDPoolDataProvider) Address() common.Address {
	return _FactoryTokenWithUIDPoolDataProvider.address
}

type FactoryTokenWithUIDPoolDataProviderInterface interface {
	GetAdditionalPoolData(opts *bind.CallOpts, localToken common.Address, arg1 uint64, arg2 [4]byte, arg3 []byte) ([]byte, error)

	GetTokenFactory(opts *bind.CallOpts) (common.Address, error)

	IsAdditionalPoolDataValid(opts *bind.CallOpts, localToken common.Address, arg1 uint64, arg2 [4]byte, additionalPoolData []byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Address() common.Address
}
