// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_usdc_token_messenger

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

var MockE2EUSDCTokenMessengerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"transmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DESTINATION_TOKEN_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositForBurn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositForBurnWithCaller\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositForBurnWithHook\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"localMessageTransmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localMessageTransmitterWithRelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IMessageTransmitterWithRelay\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageBodyVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositForBurn\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"burnToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositForBurn\",\"inputs\":[{\"name\":\"burnToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x60e0346100c557601f610e3638819003918201601f19168301916001600160401b038311848410176100ca5780849260409485528339810103126100c55780519063ffffffff821682036100c557602001516001600160a01b038116918282036100c557608052600080546001600160401b031916600117905560a05260c052604051610d5590816100e1823960805181818161020c015281816103f00152818161054e0152610832015260a05181610a50015260c0518181816103910152610c620152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b600090813560e01c9081632c12192114610a0657508063779b432d146106de5780637eccf63e146106995780638e0250ee146104145780639cdbb181146103b5578063a250c66a14610346578063f856ddb6146100d45763fb8406a91461007957600080fd5b346100d157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d15760206040517f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f68152f35b80fd5b50346100d15760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d15760043590610110610a78565b9160443561011c610aa3565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905273ffffffffffffffffffffffffffffffffffffffff91909116916084359160208160648189885af1801561030e57610319575b50823b1561030a576040517f42966c68000000000000000000000000000000000000000000000000000000008152846004820152858160248183885af1801561030e576102f5575b5060209563ffffffff9167ffffffffffffffff6102616040517fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060e01b168b8201528760248201528360448201528860648201523360848201526084815261025a60a482610ac6565b8685610bcd565b1696877fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000825416179055604051958652878601521660408401527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f660608401526080830152827f2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c060a03394a4604051908152f35b610300868092610ac6565b61030a57386101d0565b8480fd5b6040513d88823e3d90fd5b61033a9060203d60201161033f575b6103328183610ac6565b810190610b36565b610188565b503d610328565b50346100d157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d157602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100d157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d157602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100d15760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d15760043561044f610a78565b906044359161045c610aa3565b60843573ffffffffffffffffffffffffffffffffffffffff61047c610a90565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790529190931695906020816064818b8b5af1801561068e57610671575b50853b1561066257866040517f42966c680000000000000000000000000000000000000000000000000000000081528660048201528181602481838c5af180156106665761064d575b509363ffffffff94918592506105c26040517fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060e01b1660208201528960248201528260448201528860648201523360848201528a60a48201528a60c48201528a60e48201528a8080378a61010482015260e481526105bb61010482610ac6565b8584610bcd565b5060405196875260208701521660408501527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f66060850152608084015260a43560a084015260e060c08401528460e08401526101008301858080378581525016917f0c8c1cbdc5190613ebd485511d4e2812cfa45eecb79d845893331fedad5130a56101003393a480f35b8161065791610ac6565b610662578638610519565b8680fd5b6040513d84823e3d90fd5b6106899060203d60201161033f576103328183610ac6565b6104d0565b6040513d8a823e3d90fd5b50346100d157807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d15767ffffffffffffffff6020915416604051908152f35b50346100d1576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d15760043561071a610a78565b604435610725610aa3565b60843592610731610a90565b9360e4359167ffffffffffffffff8311610a025736602384011215610a025782600401359367ffffffffffffffff85116109e157602484019360248636920101116109e1576040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810189905273ffffffffffffffffffffffffffffffffffffffff9190911695906020816064818d8b5af180156109d6576109e5575b50853b156109e1576040517f42966c680000000000000000000000000000000000000000000000000000000081528860048201528981602481838b5af180156109d6576109bc575b5090849392916040517f000000000000000000000000000000000000000000000000000000000000000060e01b7fffffffff000000000000000000000000000000000000000000000000000000001660208201528760248201528160448201528960648201523360848201528a60a48201528a60c48201528a60e48201528585610104830137808681018c61010482015203610104017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810182526108ec9082610ac6565b6108f7908484610bcd565b50604051988952602089015263ffffffff166040880152606087017f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f69052608087015260a43560a087015260c0860160e090528160e0870152610100860137848185016101000152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01683019163ffffffff169280339303610100017f0c8c1cbdc5190613ebd485511d4e2812cfa45eecb79d845893331fedad5130a591a480f35b986109cc8187969594939b610ac6565b9890919293610827565b6040513d8c823e3d90fd5b8880fd5b6109fd9060203d60201161033f576103328183610ac6565b6107df565b8780fd5b905034610a7457817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610a745760209073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5080fd5b6024359063ffffffff82168203610a8b57565b600080fd5b60c4359063ffffffff82168203610a8b57565b6064359073ffffffffffffffffffffffffffffffffffffffff82168203610a8b57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610b0757604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90816020910312610a8b57518015158103610a8b5790565b90816020910312610a8b575167ffffffffffffffff81168103610a8b5790565b919082519283825260005b848110610bb85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610b79565b9080610cce575063ffffffff602091610c4760405194859384937f0ba469bc0000000000000000000000000000000000000000000000000000000085521660048401527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f66024840152606060448401526064830190610b6e565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610cc257600091610c96575090565b610cb8915060203d602011610cbb575b610cb08183610ac6565b810190610b4e565b90565b503d610ca6565b6040513d6000823e3d90fd5b91602091610c4763ffffffff9260405195869485947ff7259a750000000000000000000000000000000000000000000000000000000086521660048501527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f660248501526044840152608060648401526084830190610b6e56fea164736f6c634300081a000a",
}

var MockE2EUSDCTokenMessengerABI = MockE2EUSDCTokenMessengerMetaData.ABI

var MockE2EUSDCTokenMessengerBin = MockE2EUSDCTokenMessengerMetaData.Bin

func DeployMockE2EUSDCTokenMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, version uint32, transmitter common.Address) (common.Address, *types.Transaction, *MockE2EUSDCTokenMessenger, error) {
	parsed, err := MockE2EUSDCTokenMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockE2EUSDCTokenMessengerBin), backend, version, transmitter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockE2EUSDCTokenMessenger{address: address, abi: *parsed, MockE2EUSDCTokenMessengerCaller: MockE2EUSDCTokenMessengerCaller{contract: contract}, MockE2EUSDCTokenMessengerTransactor: MockE2EUSDCTokenMessengerTransactor{contract: contract}, MockE2EUSDCTokenMessengerFilterer: MockE2EUSDCTokenMessengerFilterer{contract: contract}}, nil
}

type MockE2EUSDCTokenMessenger struct {
	address common.Address
	abi     abi.ABI
	MockE2EUSDCTokenMessengerCaller
	MockE2EUSDCTokenMessengerTransactor
	MockE2EUSDCTokenMessengerFilterer
}

type MockE2EUSDCTokenMessengerCaller struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTokenMessengerTransactor struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTokenMessengerFilterer struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTokenMessengerSession struct {
	Contract     *MockE2EUSDCTokenMessenger
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockE2EUSDCTokenMessengerCallerSession struct {
	Contract *MockE2EUSDCTokenMessengerCaller
	CallOpts bind.CallOpts
}

type MockE2EUSDCTokenMessengerTransactorSession struct {
	Contract     *MockE2EUSDCTokenMessengerTransactor
	TransactOpts bind.TransactOpts
}

type MockE2EUSDCTokenMessengerRaw struct {
	Contract *MockE2EUSDCTokenMessenger
}

type MockE2EUSDCTokenMessengerCallerRaw struct {
	Contract *MockE2EUSDCTokenMessengerCaller
}

type MockE2EUSDCTokenMessengerTransactorRaw struct {
	Contract *MockE2EUSDCTokenMessengerTransactor
}

func NewMockE2EUSDCTokenMessenger(address common.Address, backend bind.ContractBackend) (*MockE2EUSDCTokenMessenger, error) {
	abi, err := abi.JSON(strings.NewReader(MockE2EUSDCTokenMessengerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockE2EUSDCTokenMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTokenMessenger{address: address, abi: abi, MockE2EUSDCTokenMessengerCaller: MockE2EUSDCTokenMessengerCaller{contract: contract}, MockE2EUSDCTokenMessengerTransactor: MockE2EUSDCTokenMessengerTransactor{contract: contract}, MockE2EUSDCTokenMessengerFilterer: MockE2EUSDCTokenMessengerFilterer{contract: contract}}, nil
}

func NewMockE2EUSDCTokenMessengerCaller(address common.Address, caller bind.ContractCaller) (*MockE2EUSDCTokenMessengerCaller, error) {
	contract, err := bindMockE2EUSDCTokenMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTokenMessengerCaller{contract: contract}, nil
}

func NewMockE2EUSDCTokenMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*MockE2EUSDCTokenMessengerTransactor, error) {
	contract, err := bindMockE2EUSDCTokenMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTokenMessengerTransactor{contract: contract}, nil
}

func NewMockE2EUSDCTokenMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*MockE2EUSDCTokenMessengerFilterer, error) {
	contract, err := bindMockE2EUSDCTokenMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTokenMessengerFilterer{contract: contract}, nil
}

func bindMockE2EUSDCTokenMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockE2EUSDCTokenMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2EUSDCTokenMessenger.Contract.MockE2EUSDCTokenMessengerCaller.contract.Call(opts, result, method, params...)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.MockE2EUSDCTokenMessengerTransactor.contract.Transfer(opts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.MockE2EUSDCTokenMessengerTransactor.contract.Transact(opts, method, params...)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2EUSDCTokenMessenger.Contract.contract.Call(opts, result, method, params...)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.contract.Transfer(opts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.contract.Transact(opts, method, params...)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCaller) DESTINATIONTOKENMESSENGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockE2EUSDCTokenMessenger.contract.Call(opts, &out, "DESTINATION_TOKEN_MESSENGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) DESTINATIONTOKENMESSENGER() ([32]byte, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DESTINATIONTOKENMESSENGER(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCallerSession) DESTINATIONTOKENMESSENGER() ([32]byte, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DESTINATIONTOKENMESSENGER(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCaller) LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2EUSDCTokenMessenger.contract.Call(opts, &out, "localMessageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) LocalMessageTransmitter() (common.Address, error) {
	return _MockE2EUSDCTokenMessenger.Contract.LocalMessageTransmitter(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCallerSession) LocalMessageTransmitter() (common.Address, error) {
	return _MockE2EUSDCTokenMessenger.Contract.LocalMessageTransmitter(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCaller) LocalMessageTransmitterWithRelay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2EUSDCTokenMessenger.contract.Call(opts, &out, "localMessageTransmitterWithRelay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) LocalMessageTransmitterWithRelay() (common.Address, error) {
	return _MockE2EUSDCTokenMessenger.Contract.LocalMessageTransmitterWithRelay(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCallerSession) LocalMessageTransmitterWithRelay() (common.Address, error) {
	return _MockE2EUSDCTokenMessenger.Contract.LocalMessageTransmitterWithRelay(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCaller) MessageBodyVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockE2EUSDCTokenMessenger.contract.Call(opts, &out, "messageBodyVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) MessageBodyVersion() (uint32, error) {
	return _MockE2EUSDCTokenMessenger.Contract.MessageBodyVersion(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCallerSession) MessageBodyVersion() (uint32, error) {
	return _MockE2EUSDCTokenMessenger.Contract.MessageBodyVersion(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCaller) SNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockE2EUSDCTokenMessenger.contract.Call(opts, &out, "s_nonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) SNonce() (uint64, error) {
	return _MockE2EUSDCTokenMessenger.Contract.SNonce(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCallerSession) SNonce() (uint64, error) {
	return _MockE2EUSDCTokenMessenger.Contract.SNonce(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactor) DepositForBurn(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.contract.Transact(opts, "depositForBurn", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DepositForBurn(&_MockE2EUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactorSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DepositForBurn(&_MockE2EUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactor) DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.contract.Transact(opts, "depositForBurnWithCaller", amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DepositForBurnWithCaller(&_MockE2EUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactorSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DepositForBurnWithCaller(&_MockE2EUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactor) DepositForBurnWithHook(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32, hookData []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.contract.Transact(opts, "depositForBurnWithHook", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) DepositForBurnWithHook(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32, hookData []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DepositForBurnWithHook(&_MockE2EUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactorSession) DepositForBurnWithHook(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32, hookData []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DepositForBurnWithHook(&_MockE2EUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold, hookData)
}

type MockE2EUSDCTokenMessengerDepositForBurnIterator struct {
	Event *MockE2EUSDCTokenMessengerDepositForBurn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2EUSDCTokenMessengerDepositForBurnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2EUSDCTokenMessengerDepositForBurn)
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
		it.Event = new(MockE2EUSDCTokenMessengerDepositForBurn)
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

func (it *MockE2EUSDCTokenMessengerDepositForBurnIterator) Error() error {
	return it.fail
}

func (it *MockE2EUSDCTokenMessengerDepositForBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2EUSDCTokenMessengerDepositForBurn struct {
	Nonce                     uint64
	BurnToken                 common.Address
	Amount                    *big.Int
	Depositor                 common.Address
	MintRecipient             [32]byte
	DestinationDomain         uint32
	DestinationTokenMessenger [32]byte
	DestinationCaller         [32]byte
	Raw                       types.Log
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerFilterer) FilterDepositForBurn(opts *bind.FilterOpts, nonce []uint64, burnToken []common.Address, depositor []common.Address) (*MockE2EUSDCTokenMessengerDepositForBurnIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _MockE2EUSDCTokenMessenger.contract.FilterLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTokenMessengerDepositForBurnIterator{contract: _MockE2EUSDCTokenMessenger.contract, event: "DepositForBurn", logs: logs, sub: sub}, nil
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerFilterer) WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTokenMessengerDepositForBurn, nonce []uint64, burnToken []common.Address, depositor []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _MockE2EUSDCTokenMessenger.contract.WatchLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2EUSDCTokenMessengerDepositForBurn)
				if err := _MockE2EUSDCTokenMessenger.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
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

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerFilterer) ParseDepositForBurn(log types.Log) (*MockE2EUSDCTokenMessengerDepositForBurn, error) {
	event := new(MockE2EUSDCTokenMessengerDepositForBurn)
	if err := _MockE2EUSDCTokenMessenger.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2EUSDCTokenMessengerDepositForBurn0Iterator struct {
	Event *MockE2EUSDCTokenMessengerDepositForBurn0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2EUSDCTokenMessengerDepositForBurn0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2EUSDCTokenMessengerDepositForBurn0)
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
		it.Event = new(MockE2EUSDCTokenMessengerDepositForBurn0)
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

func (it *MockE2EUSDCTokenMessengerDepositForBurn0Iterator) Error() error {
	return it.fail
}

func (it *MockE2EUSDCTokenMessengerDepositForBurn0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2EUSDCTokenMessengerDepositForBurn0 struct {
	BurnToken                 common.Address
	Amount                    *big.Int
	Depositor                 common.Address
	MintRecipient             [32]byte
	DestinationDomain         uint32
	DestinationTokenMessenger [32]byte
	DestinationCaller         [32]byte
	MaxFee                    *big.Int
	MinFinalityThreshold      uint32
	HookData                  []byte
	Raw                       types.Log
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerFilterer) FilterDepositForBurn0(opts *bind.FilterOpts, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (*MockE2EUSDCTokenMessengerDepositForBurn0Iterator, error) {

	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	var minFinalityThresholdRule []interface{}
	for _, minFinalityThresholdItem := range minFinalityThreshold {
		minFinalityThresholdRule = append(minFinalityThresholdRule, minFinalityThresholdItem)
	}

	logs, sub, err := _MockE2EUSDCTokenMessenger.contract.FilterLogs(opts, "DepositForBurn0", burnTokenRule, depositorRule, minFinalityThresholdRule)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTokenMessengerDepositForBurn0Iterator{contract: _MockE2EUSDCTokenMessenger.contract, event: "DepositForBurn0", logs: logs, sub: sub}, nil
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerFilterer) WatchDepositForBurn0(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTokenMessengerDepositForBurn0, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (event.Subscription, error) {

	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	var minFinalityThresholdRule []interface{}
	for _, minFinalityThresholdItem := range minFinalityThreshold {
		minFinalityThresholdRule = append(minFinalityThresholdRule, minFinalityThresholdItem)
	}

	logs, sub, err := _MockE2EUSDCTokenMessenger.contract.WatchLogs(opts, "DepositForBurn0", burnTokenRule, depositorRule, minFinalityThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2EUSDCTokenMessengerDepositForBurn0)
				if err := _MockE2EUSDCTokenMessenger.contract.UnpackLog(event, "DepositForBurn0", log); err != nil {
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

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerFilterer) ParseDepositForBurn0(log types.Log) (*MockE2EUSDCTokenMessengerDepositForBurn0, error) {
	event := new(MockE2EUSDCTokenMessengerDepositForBurn0)
	if err := _MockE2EUSDCTokenMessenger.contract.UnpackLog(event, "DepositForBurn0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (MockE2EUSDCTokenMessengerDepositForBurn) Topic() common.Hash {
	return common.HexToHash("0x2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c0")
}

func (MockE2EUSDCTokenMessengerDepositForBurn0) Topic() common.Hash {
	return common.HexToHash("0x0c8c1cbdc5190613ebd485511d4e2812cfa45eecb79d845893331fedad5130a5")
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessenger) Address() common.Address {
	return _MockE2EUSDCTokenMessenger.address
}

type MockE2EUSDCTokenMessengerInterface interface {
	DESTINATIONTOKENMESSENGER(opts *bind.CallOpts) ([32]byte, error)

	LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error)

	LocalMessageTransmitterWithRelay(opts *bind.CallOpts) (common.Address, error)

	MessageBodyVersion(opts *bind.CallOpts) (uint32, error)

	SNonce(opts *bind.CallOpts) (uint64, error)

	DepositForBurn(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32) (*types.Transaction, error)

	DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error)

	DepositForBurnWithHook(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee *big.Int, minFinalityThreshold uint32, hookData []byte) (*types.Transaction, error)

	FilterDepositForBurn(opts *bind.FilterOpts, nonce []uint64, burnToken []common.Address, depositor []common.Address) (*MockE2EUSDCTokenMessengerDepositForBurnIterator, error)

	WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTokenMessengerDepositForBurn, nonce []uint64, burnToken []common.Address, depositor []common.Address) (event.Subscription, error)

	ParseDepositForBurn(log types.Log) (*MockE2EUSDCTokenMessengerDepositForBurn, error)

	FilterDepositForBurn0(opts *bind.FilterOpts, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (*MockE2EUSDCTokenMessengerDepositForBurn0Iterator, error)

	WatchDepositForBurn0(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTokenMessengerDepositForBurn0, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (event.Subscription, error)

	ParseDepositForBurn0(log types.Log) (*MockE2EUSDCTokenMessengerDepositForBurn0, error)

	Address() common.Address
}
