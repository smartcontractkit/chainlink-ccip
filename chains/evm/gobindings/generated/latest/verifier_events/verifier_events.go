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

type InternalAny2EVMMultiProofMessage struct {
	Header            InternalHeader
	Sender            []byte
	Data              []byte
	Receiver          common.Address
	GasLimit          uint32
	TokenAmounts      []InternalAny2EVMMultiProofTokenTransfer
	RequiredVerifiers []InternalRequiredVerifier
}

type InternalAny2EVMMultiProofTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  common.Address
	ExtraData         []byte
	Amount            *big.Int
}

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
	ABI: "[{\"type\":\"function\",\"name\":\"emitCCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVM2AnyCommitVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destChainExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierExtraArgs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredVerifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exposeAny2EVMMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyCommitVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destChainExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierExtraArgs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredVerifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false}]",
	Bin: "0x60808060405234601557610d94908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806329f22dbe1461029157637f50332f1461003257600080fd5b346101c75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043567ffffffffffffffff81116101c7576101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c7576100a8610a86565b906100b63682600401610b0a565b8252608481013567ffffffffffffffff81116101c7576100dc9060043691840101610b78565b602083015260a481013567ffffffffffffffff81116101c7576101059060043691840101610b78565b604083015261011660c48201610b57565b606083015260e481013563ffffffff811681036101c757608083015261010481013567ffffffffffffffff81116101c7578101366023820112156101c757600481013561016a61016582610be9565b610ac6565b91602060048185858152019360051b83010101903682116101c75760248101925b8284106101cc575050505060a083015261012481013567ffffffffffffffff81116101c75760c09160046101c29236920101610c01565b910152005b600080fd5b833567ffffffffffffffff81116101c75760049083010160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082360301126101c757610217610aa6565b91602082013567ffffffffffffffff81116101c75761023c9060203691850101610b78565b835261024a60408301610b57565b602084015260608201359267ffffffffffffffff84116101c75760806020949361027a8695863691840101610b78565b60408401520135606082015281520193019261018b565b346101c75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c75760043567ffffffffffffffff81168091036101c7576024359067ffffffffffffffff82168092036101c75760443567ffffffffffffffff81116101c7576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c75760405190610160820182811067ffffffffffffffff821117610a42576040526103563682600401610b0a565b825261036460848201610b57565b6020830190815260a482013567ffffffffffffffff81116101c75761038f9060043691850101610b78565b916040840192835260c481013567ffffffffffffffff81116101c7576103bb9060043691840101610b78565b926060850193845260e482013567ffffffffffffffff81116101c7576103e79060043691850101610b78565b936080860194855261010483013567ffffffffffffffff81116101c757830194366023870112156101c757600486013561042361016582610be9565b9660206004818a858152019360051b83010101903682116101c75760248101925b828410610a10575050505060a087019586526104636101248501610b57565b9060c0880191825260e08801926101448601358452610100890194610164870135865261018487013567ffffffffffffffff81116101c757870196366023890112156101c75760048801356104ba61016582610be9565b9860206004818c858152019360051b83010101903682116101c75760248101925b82841061090757505050506101208b019788526101a481013567ffffffffffffffff81116101c7573691016004019061051391610c01565b976101408b019889526040519a60208c5251805160208d0152602081015167ffffffffffffffff1660408d0152604081015167ffffffffffffffff1660608d01526060015167ffffffffffffffff1660808c01525173ffffffffffffffffffffffffffffffffffffffff1660a08b01525160c08a016101c090526101e08a0161059b91610d28565b9051908981037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00160e08b01526105d191610d28565b9051908881037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0016101008a015261060891610d28565b9551958781037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001610120890152865180825260208201918160051b810160200198602001926000915b8383106108bc57505050505073ffffffffffffffffffffffffffffffffffffffff905116610140870152516101608601525161018085015251917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848203016101a0850152825180825260208201916020808360051b8301019501926000915b8383106107ea57505050505051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838203016101c0840152815180825260208201916020808360051b8301019401926000915b8383106107575788887f7d6fb821cf54c871623cf9ddb80288c52a51263d358a7125cf8f7a1d9d4ee5618989038aa3005b90919293946020806107db837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260806107ac8584015160a08785015260a0840190610d28565b926040810151604084015267ffffffffffffffff60608201511660608401520151906080818403910152610d28565b97019301930191939290610726565b9091929395602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865289519073ffffffffffffffffffffffffffffffffffffffff825116815273ffffffffffffffffffffffffffffffffffffffff83830151168382015260c0806108a761088b610879604087015160e0604088015260e0870190610d28565b60608701518682036060880152610d28565b6080860151608086015260a086015185820360a0870152610d28565b930151910152980193019301919392906106d2565b90919293996020806108f8838e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0876001970301885251610d28565b9c019301930191939290610652565b833567ffffffffffffffff81116101c75760049083010160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082360301126101c757610952610a86565b9161095f60208301610b57565b835261096d60408301610b57565b6020840152606082013567ffffffffffffffff81116101c7576109969060203691850101610b78565b6040840152608082013567ffffffffffffffff81116101c7576109bf9060203691850101610b78565b606084015260a0820135608084015260c08201359267ffffffffffffffff84116101c75760e0602094936109f98695863691840101610b78565b60a0840152013560c08201528152019301926104db565b833567ffffffffffffffff81116101c757602091610a378392836004369288010101610b78565b815201930192610444565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b359067ffffffffffffffff821682036101c757565b6040519060e0820182811067ffffffffffffffff821117610a4257604052565b604051906080820182811067ffffffffffffffff821117610a4257604052565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f604051930116820182811067ffffffffffffffff821117610a4257604052565b91908260809103126101c757610b506060610b23610aa6565b9380358552610b3460208201610a71565b6020860152610b4560408201610a71565b604086015201610a71565b6060830152565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c757565b81601f820112156101c75780359067ffffffffffffffff8211610a4257610bc660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011601610ac6565b92828452602083830101116101c757816000926020809301838601378301015290565b67ffffffffffffffff8111610a425760051b60200190565b9080601f830112156101c757813591610c1c61016584610be9565b9260208085838152019160051b830101918383116101c75760208101915b838310610c4957505050505090565b823567ffffffffffffffff81116101c75782019060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083880301126101c7576040519060a0820182811067ffffffffffffffff821117610a425760405260208301358252604083013567ffffffffffffffff81116101c757876020610cd192860101610b78565b602083015260608301356040830152610cec60808401610a71565b606083015260a08301359167ffffffffffffffff83116101c757610d1888602080969581960101610b78565b6080820152815201920191610c3a565b919082519283825260005b848110610d725750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610d3356fea164736f6c634300081a000a",
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

func (_VerifierEvents *VerifierEventsCaller) ExposeAny2EVMMessage(opts *bind.CallOpts, message InternalAny2EVMMultiProofMessage) error {
	var out []interface{}
	err := _VerifierEvents.contract.Call(opts, &out, "exposeAny2EVMMessage", message)

	if err != nil {
		return err
	}

	return err

}

func (_VerifierEvents *VerifierEventsSession) ExposeAny2EVMMessage(message InternalAny2EVMMultiProofMessage) error {
	return _VerifierEvents.Contract.ExposeAny2EVMMessage(&_VerifierEvents.CallOpts, message)
}

func (_VerifierEvents *VerifierEventsCallerSession) ExposeAny2EVMMessage(message InternalAny2EVMMultiProofMessage) error {
	return _VerifierEvents.Contract.ExposeAny2EVMMessage(&_VerifierEvents.CallOpts, message)
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
	ExposeAny2EVMMessage(opts *bind.CallOpts, message InternalAny2EVMMultiProofMessage) error

	EmitCCIPMessageSent(opts *bind.TransactOpts, destChainSelector uint64, sequenceNumber uint64, message InternalEVM2AnyCommitVerifierMessage) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierEventsCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierEventsCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*VerifierEventsCCIPMessageSent, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
