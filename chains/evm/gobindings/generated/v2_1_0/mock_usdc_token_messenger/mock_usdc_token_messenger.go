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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"transmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DESTINATION_TOKEN_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositForBurn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositForBurnWithCaller\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositForBurnWithHook\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"i_tokenMinter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract MockTokenMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localMessageTransmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localMessageTransmitterWithRelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IMessageTransmitterWithRelay\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localMinter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageBodyVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositForBurn\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"burnToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositForBurn\",\"inputs\":[{\"name\":\"burnToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x6101003461012757601f61110838819003918201601f19168301916001600160401b038311848410176101115780849260409485528339810103126101275780519063ffffffff8216820361012757602001516001600160a01b0381169182820361012757608052600080546001600160401b031916600117905560a05260e0526040516101928082016001600160401b03811183821017610111578291610f76833903906000f080156101055760c052604051610e49908161012d8239608051818181610291015281816104e4015281816106420152610926015260a05181610b44015260c05181818160d90152610416015260e0518181816104850152610d560152f35b6040513d6000823e3d90fd5b634e487b7160e01b600052604160045260246000fd5b600080fdfe608080604052600436101561001357600080fd5b600090813560e01c9081632c12192114610afa57508063779b432d146107d25780637eccf63e1461078d5780638e0250ee146105085780639cdbb181146104a9578063a250c66a1461043a578063cb75c11c146103cb578063f856ddb614610159578063fb8406a9146101005763fffd8cc01461008f57600080fd5b346100fd57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b80fd5b50346100fd57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760206040517f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f68152f35b50346100fd5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760043590610195610b6c565b916044356101a1610b97565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810184905273ffffffffffffffffffffffffffffffffffffffff91909116916084359160208160648189885af180156103935761039e575b50823b1561038f576040517f42966c68000000000000000000000000000000000000000000000000000000008152846004820152858160248183885af180156103935761037a575b5060209563ffffffff9167ffffffffffffffff6102e66040517fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060e01b168b820152876024820152836044820152886064820152336084820152608481526102df60a482610bba565b8685610cc1565b1696877fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000825416179055604051958652878601521660408401527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f660608401526080830152827f2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c060a03394a4604051908152f35b610385868092610bba565b61038f5738610255565b8480fd5b6040513d88823e3d90fd5b6103bf9060203d6020116103c4575b6103b78183610bba565b810190610c2a565b61020d565b503d6103ad565b50346100fd57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100fd57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100fd57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100fd5760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd57600435610543610b6c565b9060443591610550610b97565b60843573ffffffffffffffffffffffffffffffffffffffff610570610b84565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790529190931695906020816064818b8b5af1801561078257610765575b50853b1561075657866040517f42966c680000000000000000000000000000000000000000000000000000000081528660048201528181602481838c5af1801561075a57610741575b509363ffffffff94918592506106b66040517fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060e01b1660208201528960248201528260448201528860648201523360848201528a60a48201528a60c48201528a60e48201528a8080378a61010482015260e481526106af61010482610bba565b8584610cc1565b5060405196875260208701521660408501527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f66060850152608084015260a43560a084015260e060c08401528460e08401526101008301858080378581525016917f0c8c1cbdc5190613ebd485511d4e2812cfa45eecb79d845893331fedad5130a56101003393a480f35b8161074b91610bba565b61075657863861060d565b8680fd5b6040513d84823e3d90fd5b61077d9060203d6020116103c4576103b78183610bba565b6105c4565b6040513d8a823e3d90fd5b50346100fd57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5767ffffffffffffffff6020915416604051908152f35b50346100fd576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100fd5760043561080e610b6c565b604435610819610b97565b60843592610825610b84565b9360e4359167ffffffffffffffff8311610af65736602384011215610af65782600401359367ffffffffffffffff8511610ad55760248401936024863692010111610ad5576040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810189905273ffffffffffffffffffffffffffffffffffffffff9190911695906020816064818d8b5af18015610aca57610ad9575b50853b15610ad5576040517f42966c680000000000000000000000000000000000000000000000000000000081528860048201528981602481838b5af18015610aca57610ab0575b5090849392916040517f000000000000000000000000000000000000000000000000000000000000000060e01b7fffffffff000000000000000000000000000000000000000000000000000000001660208201528760248201528160448201528960648201523360848201528a60a48201528a60c48201528a60e48201528585610104830137808681018c61010482015203610104017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810182526109e09082610bba565b6109eb908484610cc1565b50604051988952602089015263ffffffff166040880152606087017f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f69052608087015260a43560a087015260c0860160e090528160e0870152610100860137848185016101000152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01683019163ffffffff169280339303610100017f0c8c1cbdc5190613ebd485511d4e2812cfa45eecb79d845893331fedad5130a591a480f35b98610ac08187969594939b610bba565b989091929361091b565b6040513d8c823e3d90fd5b8880fd5b610af19060203d6020116103c4576103b78183610bba565b6108d3565b8780fd5b905034610b6857817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610b685760209073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5080fd5b6024359063ffffffff82168203610b7f57565b600080fd5b60c4359063ffffffff82168203610b7f57565b6064359073ffffffffffffffffffffffffffffffffffffffff82168203610b7f57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610bfb57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90816020910312610b7f57518015158103610b7f5790565b90816020910312610b7f575167ffffffffffffffff81168103610b7f5790565b919082519283825260005b848110610cac5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610c6d565b9080610dc2575063ffffffff602091610d3b60405194859384937f0ba469bc0000000000000000000000000000000000000000000000000000000085521660048401527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f66024840152606060448401526064830190610c62565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af1908115610db657600091610d8a575090565b610dac915060203d602011610daf575b610da48183610bba565b810190610c42565b90565b503d610d9a565b6040513d6000823e3d90fd5b91602091610d3b63ffffffff9260405195869485947ff7259a750000000000000000000000000000000000000000000000000000000086521660048501527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f660248501526044840152608060648401526084830190610c6256fea164736f6c634300081a000a60808060405234601557610177908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80636a879ac4146100b0576378a0565e1461003257600080fd5b346100ab5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ab5763ffffffff61006e610157565b1660005260006020526040600020602435600052602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b600080fd5b346100ab5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ab576100e7610157565b6044359073ffffffffffffffffffffffffffffffffffffffff82168092036100ab5763ffffffff16600052600060205260406000206024356000526020526040600020907fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055600080f35b6004359063ffffffff821682036100ab5756fea164736f6c634300081a000a",
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

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCaller) ITokenMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2EUSDCTokenMessenger.contract.Call(opts, &out, "i_tokenMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) ITokenMinter() (common.Address, error) {
	return _MockE2EUSDCTokenMessenger.Contract.ITokenMinter(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCallerSession) ITokenMinter() (common.Address, error) {
	return _MockE2EUSDCTokenMessenger.Contract.ITokenMinter(&_MockE2EUSDCTokenMessenger.CallOpts)
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

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCaller) LocalMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2EUSDCTokenMessenger.contract.Call(opts, &out, "localMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) LocalMinter() (common.Address, error) {
	return _MockE2EUSDCTokenMessenger.Contract.LocalMinter(&_MockE2EUSDCTokenMessenger.CallOpts)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerCallerSession) LocalMinter() (common.Address, error) {
	return _MockE2EUSDCTokenMessenger.Contract.LocalMinter(&_MockE2EUSDCTokenMessenger.CallOpts)
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

	ITokenMinter(opts *bind.CallOpts) (common.Address, error)

	LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error)

	LocalMessageTransmitterWithRelay(opts *bind.CallOpts) (common.Address, error)

	LocalMinter(opts *bind.CallOpts) (common.Address, error)

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
