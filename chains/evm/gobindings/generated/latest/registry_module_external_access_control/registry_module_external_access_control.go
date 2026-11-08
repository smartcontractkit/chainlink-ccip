// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry_module_external_access_control

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

var RegistryModuleExternalAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CCIP_TOKEN_REGISTRY_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerAccessControlCCIPTokenRegistryAdmin\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAccessControlDefaultAdmin\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAdminViaGetCCIPAdmin\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAdminViaOwner\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerExternalAccessControlCCIPTokenRegistryAdmin\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerExternalAccessControlDefaultAdmin\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AdministratorRegistered\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"administrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CanOnlySelfRegister\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"roleAdmin\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RequiredRoleNotFound\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a034607e57601f610d9538819003918201601f19168301916001600160401b03831184841017608357808492602094604052833981010312607e57516001600160a01b03811690819003607e578015606d57608052604051610cfb908161009a82396080518161081e0152f35b639fabe1c160e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b60003560e01c908163181f5a7714610574575080632a3caeae1461051b57806359330dc91461046c5780637e752ca1146103a55780639f292e651461026a578063bac898251461019e578063c29c51001461014d5763de4c02781461007757600080fd5b3461014857610085366106ea565b9060005b82811061009257005b6100a56100a082858561075c565b61079b565b906040517f8fd6a6ac00000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff87165afa801561013c576001936101069260009261010c575b506107e8565b01610089565b61012e91925060203d8111610135575b61012681836106a9565b8101906107bc565b9038610100565b503d61011c565b6040513d6000823e3d90fd5b600080fd5b346101485761015b366106ea565b60005b81811061016757005b8061019861017b6100a0600194868861075c565b73ffffffffffffffffffffffffffffffffffffffff811690610935565b0161015e565b34610148576101ac366106ea565b9060005b8281106101b957005b6101c76100a082858561075c565b73ffffffffffffffffffffffffffffffffffffffff8116604051917fa217fddf000000000000000000000000000000000000000000000000000000008352602083600481855afa92831561013c57600093610232575b509161022c9160019493610bf1565b016101b0565b90926020823d8211610262575b8161024c602093836106a9565b8101031261025f5750519161022c61021d565b80fd5b3d915061023f565b3461014857610278366106ea565b9060005b82811061028557005b6102936100a082858561075c565b6040517f50c09be300000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff86165afa801561013c5773ffffffffffffffffffffffffffffffffffffffff91600091610387575b5016604051917fa217fddf000000000000000000000000000000000000000000000000000000008352602083600481855afa92831561013c57600093610352575b509161034c9160019493610bf1565b0161027c565b90926020823d821161037f575b8161036c602093836106a9565b8101031261025f5750519161034c61033d565b3d915061035f565b61039f915060203d81116101355761012681836106a9565b866102fc565b34610148576103b3366106ea565b9060005b8281106103c057005b6103ce6100a082858561075c565b90604051917f50c09be300000000000000000000000000000000000000000000000000000000835260208360048173ffffffffffffffffffffffffffffffffffffffff85165afa90811561013c5773ffffffffffffffffffffffffffffffffffffffff6104489260019560009161044e575b501690610935565b016103b7565b610466915060203d81116101355761012681836106a9565b88610440565b346101485761047a366106ea565b9060005b82811061048757005b6104956100a082858561075c565b906040517f8da5cb5b00000000000000000000000000000000000000000000000000000000815260208160048173ffffffffffffffffffffffffffffffffffffffff87165afa801561013c576001936104f5926000926104fb57506107e8565b0161047e565b61051491925060203d81116101355761012681836106a9565b9087610100565b346101485760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101485760206040517f0c0d1d9add5f8ad50021c3f2749513234d117723135ca28d672e7edf29cd89618152f35b346101485760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610148576060810181811067ffffffffffffffff82111761067a57604052602d81527f52656769737472794d6f64756c6545787465726e616c416363657373436f6e7460208201527f726f6c20322e312e302d64657600000000000000000000000000000000000000604082015260405190602082528181519182602083015260005b8381106106625750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f836000604080968601015201168101030190f35b60208282018101516040878401015285935001610622565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761067a57604052565b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101485760043567ffffffffffffffff811161014857826023820112156101485780600401359267ffffffffffffffff84116101485760248460051b83010111610148576024019190565b919081101561076c5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036101485790565b90816020910312610148575173ffffffffffffffffffffffffffffffffffffffff811681036101485790565b9073ffffffffffffffffffffffffffffffffffffffff16903382036108ee5773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b156101485773ffffffffffffffffffffffffffffffffffffffff906040519182917fe677ae3700000000000000000000000000000000000000000000000000000000835216928360048301528460248301528160446000948580945af180156108e35790827f09590fb70af4b833346363965e043a9339e8c7d378b8a2b903c75c277faec4f993926108d3575b9050a3565b6108dc916106a9565b38816108ce565b6040513d84823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff917fc454d182000000000000000000000000000000000000000000000000000000006000526004521660245260446000fd5b9073ffffffffffffffffffffffffffffffffffffffff166040517f248a9ca30000000000000000000000000000000000000000000000000000000081527f0c0d1d9add5f8ad50021c3f2749513234d117723135ca28d672e7edf29cd89616004820152602081602481855afa90811561013c57600091610bbf575b506040517fa217fddf000000000000000000000000000000000000000000000000000000008152602081600481865afa90811561013c57600091610b8d575b508103610b2157506020604491604051928380927f91d148540000000000000000000000000000000000000000000000000000000082527f0c0d1d9add5f8ad50021c3f2749513234d117723135ca28d672e7edf29cd896160048301523360248301525afa90811561013c57600091610ae2575b5015610a7657610a749033906107e8565b565b73ffffffffffffffffffffffffffffffffffffffff907f86e0b34400000000000000000000000000000000000000000000000000000000600052336004527f0c0d1d9add5f8ad50021c3f2749513234d117723135ca28d672e7edf29cd89616024521660445260646000fd5b6020813d602011610b19575b81610afb602093836106a9565b81010312610b15575190811515820361025f575038610a63565b5080fd5b3d9150610aee565b8273ffffffffffffffffffffffffffffffffffffffff917fcbdeba4c000000000000000000000000000000000000000000000000000000006000527f0c0d1d9add5f8ad50021c3f2749513234d117723135ca28d672e7edf29cd89616004526024521660445260646000fd5b90506020813d602011610bb7575b81610ba8602093836106a9565b810103126101485751386109ef565b3d9150610b9b565b906020823d602011610be9575b81610bd9602093836106a9565b8101031261025f575051386109b0565b3d9150610bcc565b9190602073ffffffffffffffffffffffffffffffffffffffff916044604051809481937f91d14854000000000000000000000000000000000000000000000000000000008352876004840152336024840152165afa90811561013c57600091610cb3575b5015610c675750610a749033906107e8565b73ffffffffffffffffffffffffffffffffffffffff92507f86e0b34400000000000000000000000000000000000000000000000000000000600052336004526024521660445260646000fd5b6020813d602011610ce6575b81610ccc602093836106a9565b81010312610b15575190811515820361025f575038610c55565b3d9150610cbf56fea164736f6c634300081a000a",
}

var RegistryModuleExternalAccessControlABI = RegistryModuleExternalAccessControlMetaData.ABI

var RegistryModuleExternalAccessControlBin = RegistryModuleExternalAccessControlMetaData.Bin

func DeployRegistryModuleExternalAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAdminRegistry common.Address) (common.Address, *types.Transaction, *RegistryModuleExternalAccessControl, error) {
	parsed, err := RegistryModuleExternalAccessControlMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryModuleExternalAccessControlBin), backend, tokenAdminRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RegistryModuleExternalAccessControl{address: address, abi: *parsed, RegistryModuleExternalAccessControlCaller: RegistryModuleExternalAccessControlCaller{contract: contract}, RegistryModuleExternalAccessControlTransactor: RegistryModuleExternalAccessControlTransactor{contract: contract}, RegistryModuleExternalAccessControlFilterer: RegistryModuleExternalAccessControlFilterer{contract: contract}}, nil
}

type RegistryModuleExternalAccessControl struct {
	address common.Address
	abi     abi.ABI
	RegistryModuleExternalAccessControlCaller
	RegistryModuleExternalAccessControlTransactor
	RegistryModuleExternalAccessControlFilterer
}

type RegistryModuleExternalAccessControlCaller struct {
	contract *bind.BoundContract
}

type RegistryModuleExternalAccessControlTransactor struct {
	contract *bind.BoundContract
}

type RegistryModuleExternalAccessControlFilterer struct {
	contract *bind.BoundContract
}

type RegistryModuleExternalAccessControlSession struct {
	Contract     *RegistryModuleExternalAccessControl
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type RegistryModuleExternalAccessControlCallerSession struct {
	Contract *RegistryModuleExternalAccessControlCaller
	CallOpts bind.CallOpts
}

type RegistryModuleExternalAccessControlTransactorSession struct {
	Contract     *RegistryModuleExternalAccessControlTransactor
	TransactOpts bind.TransactOpts
}

type RegistryModuleExternalAccessControlRaw struct {
	Contract *RegistryModuleExternalAccessControl
}

type RegistryModuleExternalAccessControlCallerRaw struct {
	Contract *RegistryModuleExternalAccessControlCaller
}

type RegistryModuleExternalAccessControlTransactorRaw struct {
	Contract *RegistryModuleExternalAccessControlTransactor
}

func NewRegistryModuleExternalAccessControl(address common.Address, backend bind.ContractBackend) (*RegistryModuleExternalAccessControl, error) {
	abi, err := abi.JSON(strings.NewReader(RegistryModuleExternalAccessControlABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindRegistryModuleExternalAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleExternalAccessControl{address: address, abi: abi, RegistryModuleExternalAccessControlCaller: RegistryModuleExternalAccessControlCaller{contract: contract}, RegistryModuleExternalAccessControlTransactor: RegistryModuleExternalAccessControlTransactor{contract: contract}, RegistryModuleExternalAccessControlFilterer: RegistryModuleExternalAccessControlFilterer{contract: contract}}, nil
}

func NewRegistryModuleExternalAccessControlCaller(address common.Address, caller bind.ContractCaller) (*RegistryModuleExternalAccessControlCaller, error) {
	contract, err := bindRegistryModuleExternalAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleExternalAccessControlCaller{contract: contract}, nil
}

func NewRegistryModuleExternalAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryModuleExternalAccessControlTransactor, error) {
	contract, err := bindRegistryModuleExternalAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleExternalAccessControlTransactor{contract: contract}, nil
}

func NewRegistryModuleExternalAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryModuleExternalAccessControlFilterer, error) {
	contract, err := bindRegistryModuleExternalAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleExternalAccessControlFilterer{contract: contract}, nil
}

func bindRegistryModuleExternalAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryModuleExternalAccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryModuleExternalAccessControl.Contract.RegistryModuleExternalAccessControlCaller.contract.Call(opts, result, method, params...)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegistryModuleExternalAccessControlTransactor.contract.Transfer(opts)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegistryModuleExternalAccessControlTransactor.contract.Transact(opts, method, params...)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryModuleExternalAccessControl.Contract.contract.Call(opts, result, method, params...)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.contract.Transfer(opts)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.contract.Transact(opts, method, params...)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlCaller) CCIPTOKENREGISTRYADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegistryModuleExternalAccessControl.contract.Call(opts, &out, "CCIP_TOKEN_REGISTRY_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlSession) CCIPTOKENREGISTRYADMINROLE() ([32]byte, error) {
	return _RegistryModuleExternalAccessControl.Contract.CCIPTOKENREGISTRYADMINROLE(&_RegistryModuleExternalAccessControl.CallOpts)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlCallerSession) CCIPTOKENREGISTRYADMINROLE() ([32]byte, error) {
	return _RegistryModuleExternalAccessControl.Contract.CCIPTOKENREGISTRYADMINROLE(&_RegistryModuleExternalAccessControl.CallOpts)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RegistryModuleExternalAccessControl.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlSession) TypeAndVersion() (string, error) {
	return _RegistryModuleExternalAccessControl.Contract.TypeAndVersion(&_RegistryModuleExternalAccessControl.CallOpts)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlCallerSession) TypeAndVersion() (string, error) {
	return _RegistryModuleExternalAccessControl.Contract.TypeAndVersion(&_RegistryModuleExternalAccessControl.CallOpts)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactor) RegisterAccessControlCCIPTokenRegistryAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.contract.Transact(opts, "registerAccessControlCCIPTokenRegistryAdmin", tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlSession) RegisterAccessControlCCIPTokenRegistryAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterAccessControlCCIPTokenRegistryAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactorSession) RegisterAccessControlCCIPTokenRegistryAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterAccessControlCCIPTokenRegistryAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactor) RegisterAccessControlDefaultAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.contract.Transact(opts, "registerAccessControlDefaultAdmin", tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlSession) RegisterAccessControlDefaultAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterAccessControlDefaultAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactorSession) RegisterAccessControlDefaultAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterAccessControlDefaultAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactor) RegisterAdminViaGetCCIPAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.contract.Transact(opts, "registerAdminViaGetCCIPAdmin", tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlSession) RegisterAdminViaGetCCIPAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterAdminViaGetCCIPAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactorSession) RegisterAdminViaGetCCIPAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterAdminViaGetCCIPAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactor) RegisterAdminViaOwner(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.contract.Transact(opts, "registerAdminViaOwner", tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlSession) RegisterAdminViaOwner(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterAdminViaOwner(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactorSession) RegisterAdminViaOwner(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterAdminViaOwner(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactor) RegisterExternalAccessControlCCIPTokenRegistryAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.contract.Transact(opts, "registerExternalAccessControlCCIPTokenRegistryAdmin", tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlSession) RegisterExternalAccessControlCCIPTokenRegistryAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterExternalAccessControlCCIPTokenRegistryAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactorSession) RegisterExternalAccessControlCCIPTokenRegistryAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterExternalAccessControlCCIPTokenRegistryAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactor) RegisterExternalAccessControlDefaultAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.contract.Transact(opts, "registerExternalAccessControlDefaultAdmin", tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlSession) RegisterExternalAccessControlDefaultAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterExternalAccessControlDefaultAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlTransactorSession) RegisterExternalAccessControlDefaultAdmin(tokens []common.Address) (*types.Transaction, error) {
	return _RegistryModuleExternalAccessControl.Contract.RegisterExternalAccessControlDefaultAdmin(&_RegistryModuleExternalAccessControl.TransactOpts, tokens)
}

type RegistryModuleExternalAccessControlAdministratorRegisteredIterator struct {
	Event *RegistryModuleExternalAccessControlAdministratorRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *RegistryModuleExternalAccessControlAdministratorRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryModuleExternalAccessControlAdministratorRegistered)
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
		it.Event = new(RegistryModuleExternalAccessControlAdministratorRegistered)
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

func (it *RegistryModuleExternalAccessControlAdministratorRegisteredIterator) Error() error {
	return it.fail
}

func (it *RegistryModuleExternalAccessControlAdministratorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type RegistryModuleExternalAccessControlAdministratorRegistered struct {
	Token         common.Address
	Administrator common.Address
	Raw           types.Log
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlFilterer) FilterAdministratorRegistered(opts *bind.FilterOpts, token []common.Address, administrator []common.Address) (*RegistryModuleExternalAccessControlAdministratorRegisteredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var administratorRule []interface{}
	for _, administratorItem := range administrator {
		administratorRule = append(administratorRule, administratorItem)
	}

	logs, sub, err := _RegistryModuleExternalAccessControl.contract.FilterLogs(opts, "AdministratorRegistered", tokenRule, administratorRule)
	if err != nil {
		return nil, err
	}
	return &RegistryModuleExternalAccessControlAdministratorRegisteredIterator{contract: _RegistryModuleExternalAccessControl.contract, event: "AdministratorRegistered", logs: logs, sub: sub}, nil
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlFilterer) WatchAdministratorRegistered(opts *bind.WatchOpts, sink chan<- *RegistryModuleExternalAccessControlAdministratorRegistered, token []common.Address, administrator []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var administratorRule []interface{}
	for _, administratorItem := range administrator {
		administratorRule = append(administratorRule, administratorItem)
	}

	logs, sub, err := _RegistryModuleExternalAccessControl.contract.WatchLogs(opts, "AdministratorRegistered", tokenRule, administratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(RegistryModuleExternalAccessControlAdministratorRegistered)
				if err := _RegistryModuleExternalAccessControl.contract.UnpackLog(event, "AdministratorRegistered", log); err != nil {
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

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControlFilterer) ParseAdministratorRegistered(log types.Log) (*RegistryModuleExternalAccessControlAdministratorRegistered, error) {
	event := new(RegistryModuleExternalAccessControlAdministratorRegistered)
	if err := _RegistryModuleExternalAccessControl.contract.UnpackLog(event, "AdministratorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (RegistryModuleExternalAccessControlAdministratorRegistered) Topic() common.Hash {
	return common.HexToHash("0x09590fb70af4b833346363965e043a9339e8c7d378b8a2b903c75c277faec4f9")
}

func (_RegistryModuleExternalAccessControl *RegistryModuleExternalAccessControl) Address() common.Address {
	return _RegistryModuleExternalAccessControl.address
}

type RegistryModuleExternalAccessControlInterface interface {
	CCIPTOKENREGISTRYADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	RegisterAccessControlCCIPTokenRegistryAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error)

	RegisterAccessControlDefaultAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error)

	RegisterAdminViaGetCCIPAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error)

	RegisterAdminViaOwner(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error)

	RegisterExternalAccessControlCCIPTokenRegistryAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error)

	RegisterExternalAccessControlDefaultAdmin(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error)

	FilterAdministratorRegistered(opts *bind.FilterOpts, token []common.Address, administrator []common.Address) (*RegistryModuleExternalAccessControlAdministratorRegisteredIterator, error)

	WatchAdministratorRegistered(opts *bind.WatchOpts, sink chan<- *RegistryModuleExternalAccessControlAdministratorRegistered, token []common.Address, administrator []common.Address) (event.Subscription, error)

	ParseAdministratorRegistered(log types.Log) (*RegistryModuleExternalAccessControlAdministratorRegistered, error)

	Address() common.Address
}
