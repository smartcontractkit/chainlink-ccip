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
	ABI: "[{\"type\":\"function\",\"name\":\"emitCCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVM2AnyCommitVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destChainExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierExtraArgs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredVerifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exposeAny2EVMMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"s_messageExecuted\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_numMessagesExecuted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_numMessagesReExecuted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyCommitVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destChainExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifierExtraArgs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"requiredVerifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageExecuted\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false}]",
	Bin: "0x60808060405234601557611291908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816329f22dbe14610291575080632fd6440b14610154578063547cb2501461010e578063710dcfd6146100c55780637b8bb9841461007657637f50332f1461006157600080fd5b346100715761006f36610c8b565b005b600080fd5b346100715760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610071576004356000526001602052602060ff604060002054166040519015158152f35b346100715760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261007157602067ffffffffffffffff60005460401c16604051908152f35b346100715760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261007157602067ffffffffffffffff60005416604051908152f35b34610071577f035593d2600da9bcde7a7f86eef93f32fdd79bbba1a76061d8a0f5095d54329e61024861018636610c8b565b61018f8161123d565b600052600160205260ff6040600020541660001461024d576000547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff6fffffffffffffffff00000000000000006101f267ffffffffffffffff8460401c16611005565b60401b169116176000555b6102068161123d565b6000526001602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905560405191829182611053565b0390a1005b6000547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000067ffffffffffffffff610285818416611005565b169116176000556101fd565b346100715760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100715760043567ffffffffffffffff8116809103610071576024359067ffffffffffffffff8216809203610071576044359267ffffffffffffffff8411610071576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc853603011261007157610160810181811067ffffffffffffffff82111761099e576040526103533685600401610a5b565b815261036160848501610aac565b6020820190815260a485013567ffffffffffffffff81116100715761038c9060043691880101610acd565b946040830195865260c481013567ffffffffffffffff8111610071576103b89060043691840101610acd565b956060840196875260e482013567ffffffffffffffff8111610071576103e49060043691850101610acd565b966080850197885261010483013567ffffffffffffffff8111610071578301973660238a01121561007157600489013561041d81610b42565b9961042b6040519b8c610a1a565b818b5260206004818d019360051b83010101903682116100715760248101925b82841061096c575050505060a0860198895261046a6101248501610aac565b9060c0870191825260e08701926101448601358452610100880194610164870135865261018487013567ffffffffffffffff811161007157870196366023890112156100715760048801356104be81610b42565b986104cc6040519a8b610a1a565b818a5260206004818c019360051b83010101903682116100715760248101925b82841061085f57505050506101208a019788526101a481013567ffffffffffffffff8111610071573691016004019061052491610b5a565b976101408a019889526040519960208b5260208b019051906105719167ffffffffffffffff6060809280518552826020820151166020860152826040820151166040860152015116910152565b5173ffffffffffffffffffffffffffffffffffffffff1660a08a01525160c089016101c090526101e089016105a591610ee7565b9051908881037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00160e08a01526105db91610ee7565b9051908781037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00161010089015261061291610ee7565b9851988681037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001610120880152895180825260208201918160051b81016020019b602001926000915b83831061081357505050505073ffffffffffffffffffffffffffffffffffffffff905116610140860152516101608501525161018084015251907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838703016101a0840152815180875260208701906020808260051b8a01019401916000905b8282106107435788887f7d6fb821cf54c871623cf9ddb80288c52a51263d358a7125cf8f7a1d9d4ee561898061073e8b8b517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016101c0850152610f46565b0390a3005b90919294602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001950301855288519073ffffffffffffffffffffffffffffffffffffffff825116815273ffffffffffffffffffffffffffffffffffffffff83830151168382015260c0806107ff6107e36107d1604087015160e0604088015260e0870190610ee7565b60608701518682036060880152610ee7565b6080860151608086015260a086015185820360a0870152610ee7565b9301519101529701920192019092916106dc565b909192939c6020808f83610850917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0876001970301885251610ee7565b9f01930193019193929061065c565b833567ffffffffffffffff81116100715760049083010160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0823603011261007157604051916108af836109e2565b6108bb60208301610aac565b83526108c960408301610aac565b6020840152606082013567ffffffffffffffff8111610071576108f29060203691850101610acd565b6040840152608082013567ffffffffffffffff81116100715761091b9060203691850101610acd565b606084015260a0820135608084015260c08201359267ffffffffffffffff84116100715760e0602094936109558695863691840101610acd565b60a0840152013560c08201528152019301926104ec565b833567ffffffffffffffff8111610071576020916109938392836004369288010101610acd565b81520193019261044b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b359067ffffffffffffffff8216820361007157565b60e0810190811067ffffffffffffffff82111761099e57604052565b6080810190811067ffffffffffffffff82111761099e57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761099e57604052565b919082608091031261007157604051610a73816109fe565b6060610aa781839580358552610a8b602082016109cd565b6020860152610a9c604082016109cd565b6040860152016109cd565b910152565b359073ffffffffffffffffffffffffffffffffffffffff8216820361007157565b81601f820112156100715780359067ffffffffffffffff821161099e5760405192610b2060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8601160185610a1a565b8284526020838301011161007157816000926020809301838601378301015290565b67ffffffffffffffff811161099e5760051b60200190565b9080601f8301121561007157813591610b7283610b42565b92610b806040519485610a1a565b80845260208085019160051b830101918383116100715760208101915b838310610bac57505050505090565b823567ffffffffffffffff81116100715782019060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08388030112610071576040519060a0820182811067ffffffffffffffff82111761099e5760405260208301358252604083013567ffffffffffffffff811161007157876020610c3492860101610acd565b602083015260608301356040830152610c4f608084016109cd565b606083015260a08301359167ffffffffffffffff831161007157610c7b88602080969581960101610acd565b6080820152815201920191610b9d565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126100715760043567ffffffffffffffff8111610071576101407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82840301126100715760405191610d01836109e2565b610d0e8183600401610a5b565b8352608482013567ffffffffffffffff811161007157816004610d3392850101610acd565b602084015260a482013567ffffffffffffffff811161007157816004610d5b92850101610acd565b6040840152610d6c60c48301610aac565b606084015260e482013563ffffffff8116810361007157608084015261010482013567ffffffffffffffff811161007157820181602382011215610071576004810135610db881610b42565b91610dc66040519384610a1a565b818352602060048185019360051b83010101908482116100715760248101925b828410610e21575050505060a08401526101248201359167ffffffffffffffff831161007157610e199201600401610b5a565b60c082015290565b833567ffffffffffffffff81116100715760049083010160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082890301126100715760405191610e71836109fe565b602082013567ffffffffffffffff811161007157886020610e9492850101610acd565b8352610ea260408301610aac565b602084015260608201359267ffffffffffffffff841161007157608083610ed08b6020809881980101610acd565b604084015201356060820152815201930192610de6565b919082519283825260005b848110610f315750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610ef2565b9080602083519182815201916020808360051b8301019401926000915b838310610f7257505050505090565b9091929394602080610ff6837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080610fc78584015160a08785015260a0840190610ee7565b926040810151604084015267ffffffffffffffff60608201511660608401520151906080818403910152610ee7565b97019301930191939290610f63565b67ffffffffffffffff1667ffffffffffffffff81146110245760010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91906020835261109360208401825167ffffffffffffffff6060809280518552826020820151166020860152826040820151166040860152015116910152565b6110e56110b1602083015161014060a0870152610160860190610ee7565b60408301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c0870152610ee7565b9273ffffffffffffffffffffffffffffffffffffffff60608301511660e082015263ffffffff60808301511661010082015260a0820151917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301610120830152825180865260208601906020808260051b8901019501916000905b8282106111ac57505050506111a993945060c00151906101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610f46565b90565b90919295602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08c600195030185528951906060806112296111f98551608086526080860190610ee7565b73ffffffffffffffffffffffffffffffffffffffff87870151168786015260408601518582036040870152610ee7565b930151910152980192019201909291611163565b60405161127e81611252602082019485611053565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610a1a565b5190209056fea164736f6c634300081a000a",
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

func (_VerifierEvents *VerifierEventsCaller) SMessageExecuted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _VerifierEvents.contract.Call(opts, &out, "s_messageExecuted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VerifierEvents *VerifierEventsSession) SMessageExecuted(arg0 [32]byte) (bool, error) {
	return _VerifierEvents.Contract.SMessageExecuted(&_VerifierEvents.CallOpts, arg0)
}

func (_VerifierEvents *VerifierEventsCallerSession) SMessageExecuted(arg0 [32]byte) (bool, error) {
	return _VerifierEvents.Contract.SMessageExecuted(&_VerifierEvents.CallOpts, arg0)
}

func (_VerifierEvents *VerifierEventsCaller) SNumMessagesExecuted(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VerifierEvents.contract.Call(opts, &out, "s_numMessagesExecuted")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VerifierEvents *VerifierEventsSession) SNumMessagesExecuted() (uint64, error) {
	return _VerifierEvents.Contract.SNumMessagesExecuted(&_VerifierEvents.CallOpts)
}

func (_VerifierEvents *VerifierEventsCallerSession) SNumMessagesExecuted() (uint64, error) {
	return _VerifierEvents.Contract.SNumMessagesExecuted(&_VerifierEvents.CallOpts)
}

func (_VerifierEvents *VerifierEventsCaller) SNumMessagesReExecuted(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VerifierEvents.contract.Call(opts, &out, "s_numMessagesReExecuted")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VerifierEvents *VerifierEventsSession) SNumMessagesReExecuted() (uint64, error) {
	return _VerifierEvents.Contract.SNumMessagesReExecuted(&_VerifierEvents.CallOpts)
}

func (_VerifierEvents *VerifierEventsCallerSession) SNumMessagesReExecuted() (uint64, error) {
	return _VerifierEvents.Contract.SNumMessagesReExecuted(&_VerifierEvents.CallOpts)
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

func (_VerifierEvents *VerifierEventsTransactor) ExecuteMessage(opts *bind.TransactOpts, message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _VerifierEvents.contract.Transact(opts, "executeMessage", message)
}

func (_VerifierEvents *VerifierEventsSession) ExecuteMessage(message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _VerifierEvents.Contract.ExecuteMessage(&_VerifierEvents.TransactOpts, message)
}

func (_VerifierEvents *VerifierEventsTransactorSession) ExecuteMessage(message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _VerifierEvents.Contract.ExecuteMessage(&_VerifierEvents.TransactOpts, message)
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

type VerifierEventsMessageExecutedIterator struct {
	Event *VerifierEventsMessageExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierEventsMessageExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierEventsMessageExecuted)
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
		it.Event = new(VerifierEventsMessageExecuted)
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

func (it *VerifierEventsMessageExecutedIterator) Error() error {
	return it.fail
}

func (it *VerifierEventsMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierEventsMessageExecuted struct {
	Message InternalAny2EVMMultiProofMessage
	Raw     types.Log
}

func (_VerifierEvents *VerifierEventsFilterer) FilterMessageExecuted(opts *bind.FilterOpts) (*VerifierEventsMessageExecutedIterator, error) {

	logs, sub, err := _VerifierEvents.contract.FilterLogs(opts, "MessageExecuted")
	if err != nil {
		return nil, err
	}
	return &VerifierEventsMessageExecutedIterator{contract: _VerifierEvents.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

func (_VerifierEvents *VerifierEventsFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *VerifierEventsMessageExecuted) (event.Subscription, error) {

	logs, sub, err := _VerifierEvents.contract.WatchLogs(opts, "MessageExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierEventsMessageExecuted)
				if err := _VerifierEvents.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
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

func (_VerifierEvents *VerifierEventsFilterer) ParseMessageExecuted(log types.Log) (*VerifierEventsMessageExecuted, error) {
	event := new(VerifierEventsMessageExecuted)
	if err := _VerifierEvents.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VerifierEvents *VerifierEvents) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierEvents.abi.Events["CCIPMessageSent"].ID:
		return _VerifierEvents.ParseCCIPMessageSent(log)
	case _VerifierEvents.abi.Events["MessageExecuted"].ID:
		return _VerifierEvents.ParseMessageExecuted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierEventsCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x7d6fb821cf54c871623cf9ddb80288c52a51263d358a7125cf8f7a1d9d4ee561")
}

func (VerifierEventsMessageExecuted) Topic() common.Hash {
	return common.HexToHash("0x035593d2600da9bcde7a7f86eef93f32fdd79bbba1a76061d8a0f5095d54329e")
}

func (_VerifierEvents *VerifierEvents) Address() common.Address {
	return _VerifierEvents.address
}

type VerifierEventsInterface interface {
	ExposeAny2EVMMessage(opts *bind.CallOpts, message InternalAny2EVMMultiProofMessage) error

	SMessageExecuted(opts *bind.CallOpts, arg0 [32]byte) (bool, error)

	SNumMessagesExecuted(opts *bind.CallOpts) (uint64, error)

	SNumMessagesReExecuted(opts *bind.CallOpts) (uint64, error)

	EmitCCIPMessageSent(opts *bind.TransactOpts, destChainSelector uint64, sequenceNumber uint64, message InternalEVM2AnyCommitVerifierMessage) (*types.Transaction, error)

	ExecuteMessage(opts *bind.TransactOpts, message InternalAny2EVMMultiProofMessage) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierEventsCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierEventsCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*VerifierEventsCCIPMessageSent, error)

	FilterMessageExecuted(opts *bind.FilterOpts) (*VerifierEventsMessageExecutedIterator, error)

	WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *VerifierEventsMessageExecuted) (event.Subscription, error)

	ParseMessageExecuted(log types.Log) (*VerifierEventsMessageExecuted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
