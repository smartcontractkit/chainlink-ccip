// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_receiver_v2

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

type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

var MockReceiverV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"required\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optional\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCVsAndMinBlockConfirmations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredVerifier\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalVerifiers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minBlockDepth\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmations\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"}]",
	Bin: "0x6080604052346101ff5761074b8038038061001981610204565b9283398101906060818303126101ff5780516001600160401b0381116101ff5782610045918301610229565b60208201519092906001600160401b0381116101ff57604091610069918401610229565b9101519160ff83168093036101ff578051906001600160401b0382116101875768010000000000000000821161018757600054826000558083106101ba575b5060200160008052602060002060005b83811061019d5784518690866001600160401b038211610187576801000000000000000082116101875760015482600155808310610141575b506020016001600052602060002060005b838110610124578460ff1960025416176002556040516104b0908161029b8239f35b82516001600160a01b031681830155602090920191600101610102565b60016000527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf69081019083015b81811061017b57506100f1565b6000815560010161016e565b634e487b7160e01b600052604160045260246000fd5b82516001600160a01b0316818301556020909201916001016100b8565b600080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5639081019083015b8181106101f357506100a8565b600081556001016101e6565b600080fd5b6040519190601f01601f191682016001600160401b0381118382101761018757604052565b81601f820112156101ff578051916001600160401b038311610187578260051b91602080610258818601610204565b8096815201938201019182116101ff57602001915b81831061027a5750505090565b82516001600160a01b03811681036101ff5781526020928301920161026d56fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461036c5750806370709a4d1461013457806385572ffb146100c55763d4d6de231461004b57600080fd5b346100c05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c05760043561ffff811681036100c0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000ff62ffff006002549260081b16911617600255600080f35b600080fd5b346100c05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c05760043567ffffffffffffffff81116100c0577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60a091360301126100c057005b346100c05760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c05760043567ffffffffffffffff8116036100c05760243567ffffffffffffffff81116100c057366023820112156100c057806004013567ffffffffffffffff81116100c057369101602401116100c05760025460405160008054808352818052829160208301917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563915b81811061033d5750505003601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210176102df5760405260405191826001548082526020820190600160005260206000209060005b81811061030e5750505003601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01683019267ffffffffffffffff8411818510176102df576102c861ffff916102ba95604052604051958695608087526080870190610459565b908582036020870152610459565b9160ff8116604085015260081c1660608301520390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b825473ffffffffffffffffffffffffffffffffffffffff16845287945060209093019260019283019201610252565b825473ffffffffffffffffffffffffffffffffffffffff168452859450602090930192600192830192016101ec565b346100c05760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100c057600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036100c057817f70709a4d000000000000000000000000000000000000000000000000000000006020931490811561042f575b8115610405575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836103fe565b7f85572ffb00000000000000000000000000000000000000000000000000000000811491506103f7565b906020808351928381520192019060005b8181106104775750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161046a56fea164736f6c634300081a000a",
}

var MockReceiverV2ABI = MockReceiverV2MetaData.ABI

var MockReceiverV2Bin = MockReceiverV2MetaData.Bin

func DeployMockReceiverV2(auth *bind.TransactOpts, backend bind.ContractBackend, required []common.Address, optional []common.Address, threshold uint8) (common.Address, *types.Transaction, *MockReceiverV2, error) {
	parsed, err := MockReceiverV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockReceiverV2Bin), backend, required, optional, threshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockReceiverV2{address: address, abi: *parsed, MockReceiverV2Caller: MockReceiverV2Caller{contract: contract}, MockReceiverV2Transactor: MockReceiverV2Transactor{contract: contract}, MockReceiverV2Filterer: MockReceiverV2Filterer{contract: contract}}, nil
}

type MockReceiverV2 struct {
	address common.Address
	abi     abi.ABI
	MockReceiverV2Caller
	MockReceiverV2Transactor
	MockReceiverV2Filterer
}

type MockReceiverV2Caller struct {
	contract *bind.BoundContract
}

type MockReceiverV2Transactor struct {
	contract *bind.BoundContract
}

type MockReceiverV2Filterer struct {
	contract *bind.BoundContract
}

type MockReceiverV2Session struct {
	Contract     *MockReceiverV2
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockReceiverV2CallerSession struct {
	Contract *MockReceiverV2Caller
	CallOpts bind.CallOpts
}

type MockReceiverV2TransactorSession struct {
	Contract     *MockReceiverV2Transactor
	TransactOpts bind.TransactOpts
}

type MockReceiverV2Raw struct {
	Contract *MockReceiverV2
}

type MockReceiverV2CallerRaw struct {
	Contract *MockReceiverV2Caller
}

type MockReceiverV2TransactorRaw struct {
	Contract *MockReceiverV2Transactor
}

func NewMockReceiverV2(address common.Address, backend bind.ContractBackend) (*MockReceiverV2, error) {
	abi, err := abi.JSON(strings.NewReader(MockReceiverV2ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockReceiverV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockReceiverV2{address: address, abi: abi, MockReceiverV2Caller: MockReceiverV2Caller{contract: contract}, MockReceiverV2Transactor: MockReceiverV2Transactor{contract: contract}, MockReceiverV2Filterer: MockReceiverV2Filterer{contract: contract}}, nil
}

func NewMockReceiverV2Caller(address common.Address, caller bind.ContractCaller) (*MockReceiverV2Caller, error) {
	contract, err := bindMockReceiverV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockReceiverV2Caller{contract: contract}, nil
}

func NewMockReceiverV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MockReceiverV2Transactor, error) {
	contract, err := bindMockReceiverV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockReceiverV2Transactor{contract: contract}, nil
}

func NewMockReceiverV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MockReceiverV2Filterer, error) {
	contract, err := bindMockReceiverV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockReceiverV2Filterer{contract: contract}, nil
}

func bindMockReceiverV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockReceiverV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockReceiverV2 *MockReceiverV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockReceiverV2.Contract.MockReceiverV2Caller.contract.Call(opts, result, method, params...)
}

func (_MockReceiverV2 *MockReceiverV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockReceiverV2.Contract.MockReceiverV2Transactor.contract.Transfer(opts)
}

func (_MockReceiverV2 *MockReceiverV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockReceiverV2.Contract.MockReceiverV2Transactor.contract.Transact(opts, method, params...)
}

func (_MockReceiverV2 *MockReceiverV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockReceiverV2.Contract.contract.Call(opts, result, method, params...)
}

func (_MockReceiverV2 *MockReceiverV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockReceiverV2.Contract.contract.Transfer(opts)
}

func (_MockReceiverV2 *MockReceiverV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockReceiverV2.Contract.contract.Transact(opts, method, params...)
}

func (_MockReceiverV2 *MockReceiverV2Caller) GetCCVsAndMinBlockConfirmations(opts *bind.CallOpts, arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockConfirmations,

	error) {
	var out []interface{}
	err := _MockReceiverV2.contract.Call(opts, &out, "getCCVsAndMinBlockConfirmations", arg0, arg1)

	outstruct := new(GetCCVsAndMinBlockConfirmations)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredVerifier = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalVerifiers = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Threshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.MinBlockDepth = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

func (_MockReceiverV2 *MockReceiverV2Session) GetCCVsAndMinBlockConfirmations(arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockConfirmations,

	error) {
	return _MockReceiverV2.Contract.GetCCVsAndMinBlockConfirmations(&_MockReceiverV2.CallOpts, arg0, arg1)
}

func (_MockReceiverV2 *MockReceiverV2CallerSession) GetCCVsAndMinBlockConfirmations(arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockConfirmations,

	error) {
	return _MockReceiverV2.Contract.GetCCVsAndMinBlockConfirmations(&_MockReceiverV2.CallOpts, arg0, arg1)
}

func (_MockReceiverV2 *MockReceiverV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockReceiverV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockReceiverV2 *MockReceiverV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockReceiverV2.Contract.SupportsInterface(&_MockReceiverV2.CallOpts, interfaceId)
}

func (_MockReceiverV2 *MockReceiverV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockReceiverV2.Contract.SupportsInterface(&_MockReceiverV2.CallOpts, interfaceId)
}

func (_MockReceiverV2 *MockReceiverV2Transactor) CcipReceive(opts *bind.TransactOpts, arg0 ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MockReceiverV2.contract.Transact(opts, "ccipReceive", arg0)
}

func (_MockReceiverV2 *MockReceiverV2Session) CcipReceive(arg0 ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MockReceiverV2.Contract.CcipReceive(&_MockReceiverV2.TransactOpts, arg0)
}

func (_MockReceiverV2 *MockReceiverV2TransactorSession) CcipReceive(arg0 ClientAny2EVMMessage) (*types.Transaction, error) {
	return _MockReceiverV2.Contract.CcipReceive(&_MockReceiverV2.TransactOpts, arg0)
}

func (_MockReceiverV2 *MockReceiverV2Transactor) SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error) {
	return _MockReceiverV2.contract.Transact(opts, "setMinBlockConfirmations", minBlockConfirmations)
}

func (_MockReceiverV2 *MockReceiverV2Session) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _MockReceiverV2.Contract.SetMinBlockConfirmations(&_MockReceiverV2.TransactOpts, minBlockConfirmations)
}

func (_MockReceiverV2 *MockReceiverV2TransactorSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _MockReceiverV2.Contract.SetMinBlockConfirmations(&_MockReceiverV2.TransactOpts, minBlockConfirmations)
}

type GetCCVsAndMinBlockConfirmations struct {
	RequiredVerifier  []common.Address
	OptionalVerifiers []common.Address
	Threshold         uint8
	MinBlockDepth     uint16
}

func (_MockReceiverV2 *MockReceiverV2) Address() common.Address {
	return _MockReceiverV2.address
}

type MockReceiverV2Interface interface {
	GetCCVsAndMinBlockConfirmations(opts *bind.CallOpts, arg0 uint64, arg1 []byte) (GetCCVsAndMinBlockConfirmations,

		error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	CcipReceive(opts *bind.TransactOpts, arg0 ClientAny2EVMMessage) (*types.Transaction, error)

	SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error)

	Address() common.Address
}
