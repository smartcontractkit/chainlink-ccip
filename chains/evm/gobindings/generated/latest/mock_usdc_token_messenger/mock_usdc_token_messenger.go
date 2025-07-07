// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_usdc_token_messenger

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"transmitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DESTINATION_TOKEN_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositForBurn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"burnToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"localMessageTransmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"localMessageTransmitterWithRelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMessageTransmitterWithRelay\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageBodyVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_nonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositForBurn\",\"inputs\":[{\"name\":\"burnToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mintRecipient\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationDomain\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"destinationCaller\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"maxFee\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"minFinalityThreshold\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"hookData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x60e0346100b757601f61083b38819003918201601f19168301916001600160401b038311848410176100bc5780849260409485528339810103126100b75780519063ffffffff821682036100b757602001516001600160a01b038116918282036100b757608052600080546001600160401b031916600117905560a05260c05260405161076890816100d382396080518181816102110152610411015260a051816104c4015260c0518181816103b201526106750152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe608080604052600436101561001357600080fd5b600090813560e01c9081632c1219211461047a575080637eccf63e146104355780639cdbb181146103d6578063a250c66a14610367578063d04857b0146100be5763fb8406a91461006357600080fd5b346100bb57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100bb5760206040517f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f68152f35b80fd5b50346100bb5760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100bb5760243560043563ffffffff8216808303610363576044356064359173ffffffffffffffffffffffffffffffffffffffff831680930361035f576084359060a4359263ffffffff841680940361035b5760c4359663ffffffff881680980361031f576040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201528760448201526020816064818d8b5af1801561031457610323575b50853b1561031f576040517f42966c680000000000000000000000000000000000000000000000000000000081528760048201528981602481838b5af18015610314576102fe575b5061027f9084604051917fffffffff000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060e01b1660208401528860248401528460448401528960648401523360848401528b60a48401528b60c48401528b60e48401528b61010484015260e4835261027a610104846104ec565b6105e0565b50604051958652602086015260408501527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f66060850152608084015260a083015260e060c08301528360e0830152836101008301527f6a4c152b4ad8c08f204453d58ef2ac1c0bb69627dd545cf47507d32d036e67d56101003393a480f35b9861030d8161027f939b6104ec565b98906101e4565b6040513d8c823e3d90fd5b8880fd5b6020813d602011610353575b8161033c602093836104ec565b8101031261034f5751801515811461019c575b8980fd5b3d915061032f565b8780fd5b8580fd5b8380fd5b50346100bb57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100bb57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100bb57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100bb57602060405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346100bb57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100bb5767ffffffffffffffff6020915416604051908152f35b9050346104e857817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104e85760209073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5080fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761052d57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b9081602091031261057c575167ffffffffffffffff8116810361057c5790565b600080fd5b919082519283825260005b8481106105cb5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161058c565b90806106e1575063ffffffff60209161065a60405194859384937f0ba469bc0000000000000000000000000000000000000000000000000000000085521660048401527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f66024840152606060448401526064830190610581565b0381600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af19081156106d5576000916106a9575090565b6106cb915060203d6020116106ce575b6106c381836104ec565b81019061055c565b90565b503d6106b9565b6040513d6000823e3d90fd5b9160209161065a63ffffffff9260405195869485947ff7259a750000000000000000000000000000000000000000000000000000000086521660048501527f17c71eed51b181d8ae1908b4743526c6dbf099c201f158a1acd5f6718e82e8f66024850152604484015260806064840152608483019061058156fea164736f6c634300081a000a",
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

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactor) DepositForBurn(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.contract.Transact(opts, "depositForBurn", amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DepositForBurn(&_MockE2EUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerTransactorSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32) (*types.Transaction, error) {
	return _MockE2EUSDCTokenMessenger.Contract.DepositForBurn(&_MockE2EUSDCTokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold)
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
	BurnToken                 common.Address
	Amount                    *big.Int
	Depositor                 common.Address
	MintRecipient             [32]byte
	DestinationDomain         uint32
	DestinationTokenMessenger [32]byte
	DestinationCaller         [32]byte
	MaxFee                    uint32
	MinFinalityThreshold      uint32
	HookData                  []byte
	Raw                       types.Log
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerFilterer) FilterDepositForBurn(opts *bind.FilterOpts, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (*MockE2EUSDCTokenMessengerDepositForBurnIterator, error) {

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

	logs, sub, err := _MockE2EUSDCTokenMessenger.contract.FilterLogs(opts, "DepositForBurn", burnTokenRule, depositorRule, minFinalityThresholdRule)
	if err != nil {
		return nil, err
	}
	return &MockE2EUSDCTokenMessengerDepositForBurnIterator{contract: _MockE2EUSDCTokenMessenger.contract, event: "DepositForBurn", logs: logs, sub: sub}, nil
}

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessengerFilterer) WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTokenMessengerDepositForBurn, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (event.Subscription, error) {

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

	logs, sub, err := _MockE2EUSDCTokenMessenger.contract.WatchLogs(opts, "DepositForBurn", burnTokenRule, depositorRule, minFinalityThresholdRule)
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

func (_MockE2EUSDCTokenMessenger *MockE2EUSDCTokenMessenger) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockE2EUSDCTokenMessenger.abi.Events["DepositForBurn"].ID:
		return _MockE2EUSDCTokenMessenger.ParseDepositForBurn(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockE2EUSDCTokenMessengerDepositForBurn) Topic() common.Hash {
	return common.HexToHash("0x6a4c152b4ad8c08f204453d58ef2ac1c0bb69627dd545cf47507d32d036e67d5")
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

	DepositForBurn(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte, maxFee uint32, minFinalityThreshold uint32) (*types.Transaction, error)

	FilterDepositForBurn(opts *bind.FilterOpts, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (*MockE2EUSDCTokenMessengerDepositForBurnIterator, error)

	WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *MockE2EUSDCTokenMessengerDepositForBurn, burnToken []common.Address, depositor []common.Address, minFinalityThreshold []uint32) (event.Subscription, error)

	ParseDepositForBurn(log types.Log) (*MockE2EUSDCTokenMessengerDepositForBurn, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
