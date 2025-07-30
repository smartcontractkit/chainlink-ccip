// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_events

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

type InternalEVM2AnyCommitVerifierMessage struct {
	Header             InternalHeader
	Sender             common.Address
	Data               []byte
	Receiver           []byte
	DestChainExtraArgs []byte
	VerifierExtraArgs  [][]byte
	FeeToken           common.Address
	FeeTokenAmount     *big.Int
	FeeValueJuels      *big.Int
	TokenAmounts       []InternalEVMTokenTransfer
	RequiredVerifiers  []InternalRequiredVerifier
}

type InternalEVMTokenTransfer struct {
	SourceTokenAddress common.Address
	SourcePoolAddress  common.Address
	DestTokenAddress   []byte
	ExtraData          []byte
	Amount             *big.Int
	DestExecData       []byte
	RequiredVerifierId [32]byte
}

type InternalHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

type InternalRequiredVerifier struct {
	VerifierId [32]byte
	Payload    []byte
	FeeAmount  *big.Int
	GasLimit   uint64
	ExtraArgs  []byte
}

var VerifierEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"emitCCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVM2AnyCommitVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destChainExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierExtraArgs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredVerifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyCommitVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destChainExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierExtraArgs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredVerifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false}]",
	Bin: "0x60808060405234601557610b08908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c6329f22dbe1461002757600080fd5b346108155760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126108155760043567ffffffffffffffff8116809103610815576024359067ffffffffffffffff82168092036108155760443567ffffffffffffffff8111610815577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc81360301906101c08212610815576040519161016083019083821067ffffffffffffffff83111761081a5760809160405212610815576040516080810181811067ffffffffffffffff82111761081a576040528160040135815261011a60248301610999565b602082015261012b60448301610999565b604082015261013c60648301610999565b6060820152825261014f608482016109f2565b6020830190815260a482013567ffffffffffffffff81116108155761017a9060043691850101610a13565b916040840192835260c481013567ffffffffffffffff8111610815576101a69060043691840101610a13565b926060850193845260e482013567ffffffffffffffff8111610815576101d29060043691850101610a13565b936080860194855261010483013567ffffffffffffffff8111610815578301943660238701121561081557600486013561021361020e82610a84565b6109ae565b9660206004818a858152019360051b83010101903682116108155760248101925b828410610967575050505060a0870195865261025361012485016109f2565b9060c0880191825260e08801926101448601358452610100890194610164870135865261018487013567ffffffffffffffff811161081557870196366023890112156108155760048801356102aa61020e82610a84565b9860206004818c858152019360051b83010101903682116108155760248101925b82841061084957505050506101208b019788526101a48101359067ffffffffffffffff82116108155701973660238a01121561081557600489013561031261020e82610a84565b9960206004818d858152019360051b83010101903682116108155760248101925b82841061073157505050506101408b019889526040519a60208c5251805160208d0152602081015167ffffffffffffffff1660408d0152604081015167ffffffffffffffff1660608d01526060015167ffffffffffffffff1660808c01525173ffffffffffffffffffffffffffffffffffffffff1660a08b01525160c08a016101c090526101e08a016103c591610a9c565b9051908981037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00160e08b01526103fb91610a9c565b9051908881037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0016101008a015261043291610a9c565b9551958781037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001610120890152865180825260208201918160051b810160200198602001926000915b8383106106e657505050505073ffffffffffffffffffffffffffffffffffffffff905116610140870152516101608601525161018085015251917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848203016101a0850152825180825260208201916020808360051b8301019501926000915b83831061061457505050505051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838203016101c0840152815180825260208201916020808360051b8301019401926000915b8383106105815788887f7d6fb821cf54c871623cf9ddb80288c52a51263d358a7125cf8f7a1d9d4ee5618989038aa3005b9091929394602080610605837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260806105d68584015160a08785015260a0840190610a9c565b926040810151604084015267ffffffffffffffff60608201511660608401520151906080818403910152610a9c565b97019301930191939290610550565b9091929395602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865289519073ffffffffffffffffffffffffffffffffffffffff825116815273ffffffffffffffffffffffffffffffffffffffff83830151168382015260c0806106d16106b56106a3604087015160e0604088015260e0870190610a9c565b60608701518682036060880152610a9c565b6080860151608086015260a086015185820360a0870152610a9c565b930151910152980193019301919392906104fc565b9091929399602080610722838e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0876001970301885251610a9c565b9c01930193019193929061047c565b833567ffffffffffffffff81116108155760049083010160a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08236030112610815576040519160a0830183811067ffffffffffffffff82111761081a5760405260208201358352604082013567ffffffffffffffff8111610815576107bd9060203691850101610a13565b6020840152606082013560408401526107d860808301610999565b606084015260a08201359267ffffffffffffffff8411610815576108056020949385809536920101610a13565b6080820152815201930192610333565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b833567ffffffffffffffff81116108155760049083010160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08236030112610815576040519160e0830183811067ffffffffffffffff82111761081a576040526108b6602083016109f2565b83526108c4604083016109f2565b6020840152606082013567ffffffffffffffff8111610815576108ed9060203691850101610a13565b6040840152608082013567ffffffffffffffff8111610815576109169060203691850101610a13565b606084015260a0820135608084015260c08201359267ffffffffffffffff84116108155760e0602094936109508695863691840101610a13565b60a0840152013560c08201528152019301926102cb565b833567ffffffffffffffff81116108155760209161098e8392836004369288010101610a13565b815201930192610234565b359067ffffffffffffffff8216820361081557565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f604051930116820182811067ffffffffffffffff82111761081a57604052565b359073ffffffffffffffffffffffffffffffffffffffff8216820361081557565b81601f820112156108155780359067ffffffffffffffff821161081a57610a6160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f850116016109ae565b928284526020838301011161081557816000926020809301838601378301015290565b67ffffffffffffffff811161081a5760051b60200190565b919082519283825260005b848110610ae65750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610aa756fea164736f6c634300081a000a",
}

var VerifierEventsABI = VerifierEventsMetaData.ABI

var VerifierEventsBin = VerifierEventsMetaData.Bin

func DeployVerifierEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerifierEvents, error) {
	parsed, err := VerifierEventsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierEvents{address: address, abi: *parsed, VerifierEventsCaller: VerifierEventsCaller{contract: contract}, VerifierEventsTransactor: VerifierEventsTransactor{contract: contract}, VerifierEventsFilterer: VerifierEventsFilterer{contract: contract}}, nil
}

type VerifierEvents struct {
	address common.Address
	abi     abi.ABI
	VerifierEventsCaller
	VerifierEventsTransactor
	VerifierEventsFilterer
}

type VerifierEventsCaller struct {
	contract *bind.BoundContract
}

type VerifierEventsTransactor struct {
	contract *bind.BoundContract
}

type VerifierEventsFilterer struct {
	contract *bind.BoundContract
}

type VerifierEventsSession struct {
	Contract     *VerifierEvents
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierEventsCallerSession struct {
	Contract *VerifierEventsCaller
	CallOpts bind.CallOpts
}

type VerifierEventsTransactorSession struct {
	Contract     *VerifierEventsTransactor
	TransactOpts bind.TransactOpts
}

type VerifierEventsRaw struct {
	Contract *VerifierEvents
}

type VerifierEventsCallerRaw struct {
	Contract *VerifierEventsCaller
}

type VerifierEventsTransactorRaw struct {
	Contract *VerifierEventsTransactor
}

func NewVerifierEvents(address common.Address, backend bind.ContractBackend) (*VerifierEvents, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierEventsABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierEvents{address: address, abi: abi, VerifierEventsCaller: VerifierEventsCaller{contract: contract}, VerifierEventsTransactor: VerifierEventsTransactor{contract: contract}, VerifierEventsFilterer: VerifierEventsFilterer{contract: contract}}, nil
}

func NewVerifierEventsCaller(address common.Address, caller bind.ContractCaller) (*VerifierEventsCaller, error) {
	contract, err := bindVerifierEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierEventsCaller{contract: contract}, nil
}

func NewVerifierEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierEventsTransactor, error) {
	contract, err := bindVerifierEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierEventsTransactor{contract: contract}, nil
}

func NewVerifierEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierEventsFilterer, error) {
	contract, err := bindVerifierEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierEventsFilterer{contract: contract}, nil
}

func bindVerifierEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierEvents *VerifierEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierEvents.Contract.VerifierEventsCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierEvents *VerifierEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierEvents.Contract.VerifierEventsTransactor.contract.Transfer(opts)
}

func (_VerifierEvents *VerifierEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierEvents.Contract.VerifierEventsTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierEvents *VerifierEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierEvents.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierEvents *VerifierEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierEvents.Contract.contract.Transfer(opts)
}

func (_VerifierEvents *VerifierEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierEvents.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierEvents *VerifierEventsTransactor) EmitCCIPMessageSent(opts *bind.TransactOpts, destChainSelector uint64, sequenceNumber uint64, message InternalEVM2AnyCommitVerifierMessage) (*types.Transaction, error) {
	return _VerifierEvents.contract.Transact(opts, "emitCCIPMessageSent", destChainSelector, sequenceNumber, message)
}

func (_VerifierEvents *VerifierEventsSession) EmitCCIPMessageSent(destChainSelector uint64, sequenceNumber uint64, message InternalEVM2AnyCommitVerifierMessage) (*types.Transaction, error) {
	return _VerifierEvents.Contract.EmitCCIPMessageSent(&_VerifierEvents.TransactOpts, destChainSelector, sequenceNumber, message)
}

func (_VerifierEvents *VerifierEventsTransactorSession) EmitCCIPMessageSent(destChainSelector uint64, sequenceNumber uint64, message InternalEVM2AnyCommitVerifierMessage) (*types.Transaction, error) {
	return _VerifierEvents.Contract.EmitCCIPMessageSent(&_VerifierEvents.TransactOpts, destChainSelector, sequenceNumber, message)
}

type VerifierEventsCCIPMessageSentIterator struct {
	Event *VerifierEventsCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierEventsCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierEventsCCIPMessageSent)
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
		it.Event = new(VerifierEventsCCIPMessageSent)
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

func (it *VerifierEventsCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *VerifierEventsCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierEventsCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalEVM2AnyCommitVerifierMessage
	Raw               types.Log
}

func (_VerifierEvents *VerifierEventsFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierEventsCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _VerifierEvents.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &VerifierEventsCCIPMessageSentIterator{contract: _VerifierEvents.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_VerifierEvents *VerifierEventsFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierEventsCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _VerifierEvents.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierEventsCCIPMessageSent)
				if err := _VerifierEvents.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_VerifierEvents *VerifierEventsFilterer) ParseCCIPMessageSent(log types.Log) (*VerifierEventsCCIPMessageSent, error) {
	event := new(VerifierEventsCCIPMessageSent)
	if err := _VerifierEvents.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VerifierEvents *VerifierEvents) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierEvents.abi.Events["CCIPMessageSent"].ID:
		return _VerifierEvents.ParseCCIPMessageSent(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierEventsCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x7d6fb821cf54c871623cf9ddb80288c52a51263d358a7125cf8f7a1d9d4ee561")
}

func (_VerifierEvents *VerifierEvents) Address() common.Address {
	return _VerifierEvents.address
}

type VerifierEventsInterface interface {
	EmitCCIPMessageSent(opts *bind.TransactOpts, destChainSelector uint64, sequenceNumber uint64, message InternalEVM2AnyCommitVerifierMessage) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierEventsCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierEventsCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*VerifierEventsCCIPMessageSent, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
