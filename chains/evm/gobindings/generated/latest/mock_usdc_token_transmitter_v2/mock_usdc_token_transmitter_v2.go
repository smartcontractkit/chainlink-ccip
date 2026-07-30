// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_usdc_token_transmitter_v2

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

var MockE2EUSDCTransmitterCCTPV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_version\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_localDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"localDomain\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextAvailableNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"attestation\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_shouldSucceed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendMessage\",\"inputs\":[{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageBody\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendMessageWithCaller\",\"inputs\":[{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"recipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageBody\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setShouldSucceed\",\"inputs\":[{\"name\":\"shouldSucceed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MessageReceived\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"attestation\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSent\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x60e0346100c857601f610b6c38819003918201601f19168301916001600160401b038311848410176100cd578084926060946040528339810103126100c857610047816100e3565b906040610056602083016100e3565b9101516001600160a01b03811692908390036100c85760805260a052600160ff19600054161760005560c052604051610a7790816100f5823960805181818161011f015281816106b90152610781015260a0518181816101490152818161041801526107ab015260c051816105570152f35b600080fd5b634e487b7160e01b600052604160045260246000fd5b519063ffffffff821682036100c85756fe6080604052600436101561001257600080fd5b6000803560e01c80630ba469bc146106dd57806354fd4d501461067e57806357ecfd28146104c45780637a642935146104845780638371744e1461043c5780638d3638f4146103dd5780639e31ddb61461036e5763f7259a751461007557600080fd5b3461036b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036b576100ac6108e8565b6044359160243560643567ffffffffffffffff8111610367576100d3903690600401610900565b9480156102e3576100e26109dd565b94831561028557866020976101ed946094947fffffffff0000000000000000000000000000000000000000000000000000000097604051988996817f000000000000000000000000000000000000000000000000000000000000000060e01b168e890152817f000000000000000000000000000000000000000000000000000000000000000060e01b16602489015260e01b1660288701527fffffffffffffffff0000000000000000000000000000000000000000000000008b60c01b16602c87015233603487015260548601526074850152848401378101858382015203017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261092e565b604051918483528151918286850152815b838110610271575050827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f846040948585977f8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b0369901015201168101030190a167ffffffffffffffff60405191168152f35b8181018701518582016040015286016101fe565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f526563697069656e74206d757374206265206e6f6e7a65726f000000000000006044820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f44657374696e6174696f6e2063616c6c6572206d757374206265206e6f6e7a6560448201527f726f0000000000000000000000000000000000000000000000000000000000006064820152fd5b8280fd5b80fd5b503461036b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036b576004358015158091036103d95760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00835416911617815580f35b5080fd5b503461036b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036b57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461036b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036b5767ffffffffffffffff6020915460081c16604051908152f35b503461036b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036b5760ff60209154166040519015158152f35b503461036b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036b5760043567ffffffffffffffff81116103d957610514903690600401610900565b9060243567ffffffffffffffff811161067a57610535903690600401610900565b92908160d811610676578473ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b156103d95781906044604051809481937f40c10f1900000000000000000000000000000000000000000000000000000000835260c48a013560601c6004840152600160248401525af1801561066b5791602096939160ff969593610633575b50610624907f245e8e0654076d706bf5dea6eab892c9d7788b4aefa7f4923e837e43d200959b949561061760405195869560408752604087019161099e565b918483038a86015261099e565b0390a154166040519015158152f35b90846106637f245e8e0654076d706bf5dea6eab892c9d7788b4aefa7f4923e837e43d200959b966106249461092e565b9450906105d8565b6040513d88823e3d90fd5b8480fd5b8380fd5b503461036b57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036b57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461036b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261036b576107156108e8565b906024359060443567ffffffffffffffff81116103d95761073a903690600401610900565b90926107446109dd565b93811561028557602095837fffffffff000000000000000000000000000000000000000000000000000000009460949361085095604051978895817f000000000000000000000000000000000000000000000000000000000000000060e01b168d880152817f000000000000000000000000000000000000000000000000000000000000000060e01b16602488015260e01b1660288601527fffffffffffffffff0000000000000000000000000000000000000000000000008a60c01b16602c8601523360348601526054850152876074850152848401378101858382015203017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261092e565b604051918483528151918286850152815b8381106108d4575050827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f846040948585977f8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b0369901015201168101030190a167ffffffffffffffff60405191168152f35b818101870151858201604001528601610861565b6004359063ffffffff821682036108fb57565b600080fd5b9181601f840112156108fb5782359167ffffffffffffffff83116108fb57602083818601950101116108fb57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761096f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60005467ffffffffffffffff8160081c16906001820167ffffffffffffffff8111610a3b5768ffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff9160081b1691161760005590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea164736f6c634300081a000a",
}

var MockE2EUSDCTransmitterCCTPV2ABI = MockE2EUSDCTransmitterCCTPV2MetaData.ABI

var MockE2EUSDCTransmitterCCTPV2Bin = MockE2EUSDCTransmitterCCTPV2MetaData.Bin

func DeployMockE2EUSDCTransmitterCCTPV2(auth *bind.TransactOpts, backend bind.ContractBackend, _version uint32, _localDomain uint32, token common.Address) (common.Address, *types.Transaction, *MockE2EUSDCTransmitterCCTPV2, error) {
	parsed, err := MockE2EUSDCTransmitterCCTPV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockE2EUSDCTransmitterCCTPV2Bin), backend, _version, _localDomain, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockE2EUSDCTransmitterCCTPV2{address: address, abi: *parsed, MockE2EUSDCTransmitterCCTPV2Caller: MockE2EUSDCTransmitterCCTPV2Caller{contract: contract}, MockE2EUSDCTransmitterCCTPV2Transactor: MockE2EUSDCTransmitterCCTPV2Transactor{contract: contract}, MockE2EUSDCTransmitterCCTPV2Filterer: MockE2EUSDCTransmitterCCTPV2Filterer{contract: contract}}, nil
}

type MockE2EUSDCTransmitterCCTPV2 struct {
	address common.Address
	abi     abi.ABI
	MockE2EUSDCTransmitterCCTPV2Caller
	MockE2EUSDCTransmitterCCTPV2Transactor
	MockE2EUSDCTransmitterCCTPV2Filterer
}

type MockE2EUSDCTransmitterCCTPV2Caller struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTransmitterCCTPV2Transactor struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTransmitterCCTPV2Filterer struct {
	contract *bind.BoundContract
}

type MockE2EUSDCTransmitterCCTPV2Session struct {
	Contract     *MockE2EUSDCTransmitterCCTPV2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockE2EUSDCTransmitterCCTPV2CallerSession struct {
	Contract *MockE2EUSDCTransmitterCCTPV2Caller
	CallOpts bind.CallOpts
}

type MockE2EUSDCTransmitterCCTPV2TransactorSession struct {
	Contract     *MockE2EUSDCTransmitterCCTPV2Transactor
	TransactOpts bind.TransactOpts
}

type MockE2EUSDCTransmitterCCTPV2Raw struct {
	Contract *MockE2EUSDCTransmitterCCTPV2
}

type MockE2EUSDCTransmitterCCTPV2CallerRaw struct {
	Contract *MockE2EUSDCTransmitterCCTPV2Caller
}

type MockE2EUSDCTransmitterCCTPV2TransactorRaw struct {
	Contract *MockE2EUSDCTransmitterCCTPV2Transactor
}

func NewMockE2EUSDCTransmitterCCTPV2(address common.Address, backend bind.ContractBackend) (*MockE2EUSDCTransmitterCCTPV2, error) {
	abi, err := abi.JSON(strings.NewReader(MockE2EUSDCTransmitterCCTPV2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockE2EUSDCTransmitterCCTPV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterCCTPV2{address: address, abi: abi, MockE2EUSDCTransmitterCCTPV2Caller: MockE2EUSDCTransmitterCCTPV2Caller{contract: contract}, MockE2EUSDCTransmitterCCTPV2Transactor: MockE2EUSDCTransmitterCCTPV2Transactor{contract: contract}, MockE2EUSDCTransmitterCCTPV2Filterer: MockE2EUSDCTransmitterCCTPV2Filterer{contract: contract}}, nil
}

func NewMockE2EUSDCTransmitterCCTPV2Caller(address common.Address, caller bind.ContractCaller) (*MockE2EUSDCTransmitterCCTPV2Caller, error) {
	contract, err := bindMockE2EUSDCTransmitterCCTPV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterCCTPV2Caller{contract: contract}, nil
}

func NewMockE2EUSDCTransmitterCCTPV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MockE2EUSDCTransmitterCCTPV2Transactor, error) {
	contract, err := bindMockE2EUSDCTransmitterCCTPV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterCCTPV2Transactor{contract: contract}, nil
}

func NewMockE2EUSDCTransmitterCCTPV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MockE2EUSDCTransmitterCCTPV2Filterer, error) {
	contract, err := bindMockE2EUSDCTransmitterCCTPV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterCCTPV2Filterer{contract: contract}, nil
}

func bindMockE2EUSDCTransmitterCCTPV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockE2EUSDCTransmitterCCTPV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.MockE2EUSDCTransmitterCCTPV2Caller.contract.Call(opts, result, method, params...)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.MockE2EUSDCTransmitterCCTPV2Transactor.contract.Transfer(opts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.MockE2EUSDCTransmitterCCTPV2Transactor.contract.Transact(opts, method, params...)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.contract.Call(opts, result, method, params...)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.contract.Transfer(opts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.contract.Transact(opts, method, params...)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Caller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitterCCTPV2.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Session) LocalDomain() (uint32, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.LocalDomain(&_MockE2EUSDCTransmitterCCTPV2.CallOpts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2CallerSession) LocalDomain() (uint32, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.LocalDomain(&_MockE2EUSDCTransmitterCCTPV2.CallOpts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Caller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitterCCTPV2.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Session) NextAvailableNonce() (uint64, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.NextAvailableNonce(&_MockE2EUSDCTransmitterCCTPV2.CallOpts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2CallerSession) NextAvailableNonce() (uint64, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.NextAvailableNonce(&_MockE2EUSDCTransmitterCCTPV2.CallOpts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Caller) SShouldSucceed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitterCCTPV2.contract.Call(opts, &out, "s_shouldSucceed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Session) SShouldSucceed() (bool, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.SShouldSucceed(&_MockE2EUSDCTransmitterCCTPV2.CallOpts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2CallerSession) SShouldSucceed() (bool, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.SShouldSucceed(&_MockE2EUSDCTransmitterCCTPV2.CallOpts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Caller) Version(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockE2EUSDCTransmitterCCTPV2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Session) Version() (uint32, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.Version(&_MockE2EUSDCTransmitterCCTPV2.CallOpts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2CallerSession) Version() (uint32, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.Version(&_MockE2EUSDCTransmitterCCTPV2.CallOpts)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Transactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.contract.Transact(opts, "receiveMessage", message, attestation)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Session) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.ReceiveMessage(&_MockE2EUSDCTransmitterCCTPV2.TransactOpts, message, attestation)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2TransactorSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.ReceiveMessage(&_MockE2EUSDCTransmitterCCTPV2.TransactOpts, message, attestation)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Transactor) SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.contract.Transact(opts, "sendMessage", destinationDomain, recipient, messageBody)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Session) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.SendMessage(&_MockE2EUSDCTransmitterCCTPV2.TransactOpts, destinationDomain, recipient, messageBody)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2TransactorSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.SendMessage(&_MockE2EUSDCTransmitterCCTPV2.TransactOpts, destinationDomain, recipient, messageBody)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Transactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Session) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.SendMessageWithCaller(&_MockE2EUSDCTransmitterCCTPV2.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2TransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.SendMessageWithCaller(&_MockE2EUSDCTransmitterCCTPV2.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Transactor) SetShouldSucceed(opts *bind.TransactOpts, shouldSucceed bool) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.contract.Transact(opts, "setShouldSucceed", shouldSucceed)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Session) SetShouldSucceed(shouldSucceed bool) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.SetShouldSucceed(&_MockE2EUSDCTransmitterCCTPV2.TransactOpts, shouldSucceed)
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2TransactorSession) SetShouldSucceed(shouldSucceed bool) (*types.Transaction, error) {
	return _MockE2EUSDCTransmitterCCTPV2.Contract.SetShouldSucceed(&_MockE2EUSDCTransmitterCCTPV2.TransactOpts, shouldSucceed)
}

type MockE2EUSDCTransmitterCCTPV2MessageReceivedIterator struct {
	Event *MockE2EUSDCTransmitterCCTPV2MessageReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2EUSDCTransmitterCCTPV2MessageReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2EUSDCTransmitterCCTPV2MessageReceived)
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
		it.Event = new(MockE2EUSDCTransmitterCCTPV2MessageReceived)
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

func (it *MockE2EUSDCTransmitterCCTPV2MessageReceivedIterator) Error() error {
	return it.fail
}

func (it *MockE2EUSDCTransmitterCCTPV2MessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2EUSDCTransmitterCCTPV2MessageReceived struct {
	Message     []byte
	Attestation []byte
	Raw         types.Log
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Filterer) FilterMessageReceived(opts *bind.FilterOpts) (*MockE2EUSDCTransmitterCCTPV2MessageReceivedIterator, error) {

	logs, sub, err := _MockE2EUSDCTransmitterCCTPV2.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterCCTPV2MessageReceivedIterator{contract: _MockE2EUSDCTransmitterCCTPV2.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Filterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTransmitterCCTPV2MessageReceived) (event.Subscription, error) {

	logs, sub, err := _MockE2EUSDCTransmitterCCTPV2.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2EUSDCTransmitterCCTPV2MessageReceived)
				if err := _MockE2EUSDCTransmitterCCTPV2.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Filterer) ParseMessageReceived(log types.Log) (*MockE2EUSDCTransmitterCCTPV2MessageReceived, error) {
	event := new(MockE2EUSDCTransmitterCCTPV2MessageReceived)
	if err := _MockE2EUSDCTransmitterCCTPV2.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2EUSDCTransmitterCCTPV2MessageSentIterator struct {
	Event *MockE2EUSDCTransmitterCCTPV2MessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2EUSDCTransmitterCCTPV2MessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2EUSDCTransmitterCCTPV2MessageSent)
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
		it.Event = new(MockE2EUSDCTransmitterCCTPV2MessageSent)
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

func (it *MockE2EUSDCTransmitterCCTPV2MessageSentIterator) Error() error {
	return it.fail
}

func (it *MockE2EUSDCTransmitterCCTPV2MessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2EUSDCTransmitterCCTPV2MessageSent struct {
	Message []byte
	Raw     types.Log
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Filterer) FilterMessageSent(opts *bind.FilterOpts) (*MockE2EUSDCTransmitterCCTPV2MessageSentIterator, error) {

	logs, sub, err := _MockE2EUSDCTransmitterCCTPV2.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTransmitterCCTPV2MessageSentIterator{contract: _MockE2EUSDCTransmitterCCTPV2.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Filterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTransmitterCCTPV2MessageSent) (event.Subscription, error) {

	logs, sub, err := _MockE2EUSDCTransmitterCCTPV2.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2EUSDCTransmitterCCTPV2MessageSent)
				if err := _MockE2EUSDCTransmitterCCTPV2.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2Filterer) ParseMessageSent(log types.Log) (*MockE2EUSDCTransmitterCCTPV2MessageSent, error) {
	event := new(MockE2EUSDCTransmitterCCTPV2MessageSent)
	if err := _MockE2EUSDCTransmitterCCTPV2.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (MockE2EUSDCTransmitterCCTPV2MessageReceived) Topic() common.Hash {
	return common.HexToHash("0x245e8e0654076d706bf5dea6eab892c9d7788b4aefa7f4923e837e43d200959b")
}

func (MockE2EUSDCTransmitterCCTPV2MessageSent) Topic() common.Hash {
	return common.HexToHash("0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036")
}

func (_MockE2EUSDCTransmitterCCTPV2 *MockE2EUSDCTransmitterCCTPV2) Address() common.Address {
	return _MockE2EUSDCTransmitterCCTPV2.address
}

type MockE2EUSDCTransmitterCCTPV2Interface interface {
	LocalDomain(opts *bind.CallOpts) (uint32, error)

	NextAvailableNonce(opts *bind.CallOpts) (uint64, error)

	SShouldSucceed(opts *bind.CallOpts) (bool, error)

	Version(opts *bind.CallOpts) (uint32, error)

	ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error)

	SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error)

	SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error)

	SetShouldSucceed(opts *bind.TransactOpts, shouldSucceed bool) (*types.Transaction, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*MockE2EUSDCTransmitterCCTPV2MessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTransmitterCCTPV2MessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*MockE2EUSDCTransmitterCCTPV2MessageReceived, error)

	FilterMessageSent(opts *bind.FilterOpts) (*MockE2EUSDCTransmitterCCTPV2MessageSentIterator, error)

	WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTransmitterCCTPV2MessageSent) (event.Subscription, error)

	ParseMessageSent(log types.Log) (*MockE2EUSDCTransmitterCCTPV2MessageSent, error)

	Address() common.Address
}
