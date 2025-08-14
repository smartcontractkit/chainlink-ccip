// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccv_aggregator

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

type CCVAggregatorAggregatedReport struct {
	Message     InternalAny2EVMMultiProofMessage
	CcvBlobs    [][]byte
	Proofs      [][]byte
	GasOverride uint32
}

type CCVAggregatorDynamicConfig struct {
	ReentrancyGuardEntered bool
}

type CCVAggregatorSourceChainConfig struct {
	Router    common.Address
	IsEnabled bool
}

type CCVAggregatorSourceChainConfigArgs struct {
	Router              common.Address
	SourceChainSelector uint64
	IsEnabled           bool
	OnRamp              []byte
}

type CCVAggregatorStaticConfig struct {
	LocalChainSelector   uint64
	GasForCallExactCheck uint16
	RmnRemote            common.Address
	TokenAdminRegistry   common.Address
}

type InternalAny2EVMMultiProofMessage struct {
	Header            InternalHeader
	Sender            []byte
	Data              []byte
	Receiver          common.Address
	GasLimit          uint32
	TokenAmounts      []InternalAny2EVMMultiProofTokenTransfer
	RequiredVerifiers []common.Address
}

type InternalAny2EVMMultiProofTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  common.Address
	ExtraData         []byte
	Amount            *big.Int
}

type InternalHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

var CCVAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.AggregatedReport\",\"components\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"ccvBlobs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"proofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461052057613e0d803803809161001d8285610556565b833981019080820360c08112610520576080811261052057604051916100428361053b565b61004b81610579565b8352602081015161ffff8116810361052057602084019081526040820151906001600160a01b038216820361052057604085019182526060830151936001600160a01b03851685036105205760209060608701958652607f1901126105205760405192602084016001600160401b03811185821017610525576040526100d36080820161058d565b845260a0810151906001600160401b038211610520570186601f82011215610520578051966001600160401b038811610525578760051b916040519861011c602085018b610556565b89526020808a0193820101908282116105205760208101935b828510610420575050505050331561040f57600180546001600160a01b0319163317905581516001600160a01b03161580156103fd575b6103ec5784516001600160401b0316156103db5784516001600160401b03908116608090815283516001600160a01b0390811660a0528651811660c052835161ffff90811660e05260408051995190941689529351909316602088810191909152935183169187019190915293511660608501527f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b369390927f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a151151560ff196002541660ff821617600255604051908152a16000905b80518210156103855760009160208160051b8301015160018060401b036020820151169081156103765780516001600160a01b031615610367578185526005602052604085209460608201518051801591821561033b575b505061032c5750604081810151865492516001600160a81b031990931690151560a01b60ff60a01b16176001600160a01b039290921691909117855592936001937f734ce292835471880fd796297a6bef2e23db07c84d09cad27f67cd7163f44035916103088461059a565b5060ff82519154878060a01b038116835260a01c1615156020820152a20190610244565b6342bcdf7f60e11b8152600490fd5b6020919250012060405160208101908382526020815261035c604082610556565b51902014388061029c565b6342bcdf7f60e11b8552600485fd5b63c656089560e01b8552600485fd5b6040516137df908161062e823960805181818161047d0152611230015260a0518181816104e00152611194015260c05181818161051c0152612128015260e0518181816104a401528181611c88015261304f0152f35b63c656089560e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b5083516001600160a01b03161561016c565b639b15e16f60e01b60005260046000fd5b84516001600160401b0381116105205782016080818603601f190112610520576040519061044d8261053b565b60208101516001600160a01b038116810361052057825261047060408201610579565b60208301526104816060820161058d565b604083015260808101516001600160401b03811161052057602091010185601f820112156105205780516001600160401b03811161052557604051916104d1601f8301601f191660200184610556565b81835287602083830101116105205760005b82811061050b5750509181600060208096949581960101526060820152815201940193610135565b806020809284010151828287010152016104e3565b600080fd5b634e487b7160e01b600052604160045260246000fd5b608081019081106001600160401b0382111761052557604052565b601f909101601f19168101906001600160401b0382119082101761052557604052565b51906001600160401b038216820361052057565b5190811515820361052057565b806000526004602052604060002054156000146106275760035468010000000000000000811015610525576001810180600355811015610611577fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0181905560035460009182526004602052604090912055600190565b634e487b7160e01b600052603260045260246000fd5b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806304666f9c146100e757806306285c69146100e25780630cea69ee146100dd578063181f5a77146100d85780634708209d146100d35780635215505b146100ce5780635e36480c146100c95780637437ff9f146100c457806379ba5097146100bf5780638da5cb5b146100ba578063e9d68a8e146100b5578063f2fde38b146100b05763f556a188146100ab57600080fd5b61107f565b610f8b565b610ec3565b610e71565b610d88565b610d1f565b610cbe565b610b25565b6109bd565b610940565b61086e565b61042b565b6102cc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761013757604052565b6100ec565b6020810190811067ffffffffffffffff82111761013757604052565b6040810190811067ffffffffffffffff82111761013757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761013757604052565b604051906101c460e083610174565b565b604051906101c460a083610174565b604051906101c461010083610174565b604051906101c4604083610174565b67ffffffffffffffff81116101375760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361022a57565b600080fd5b67ffffffffffffffff81160361022a57565b8015150361022a57565b67ffffffffffffffff811161013757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561022a5780359061029c8261024b565b926102aa6040519485610174565b8284526020838301011161022a57816000926020809301838601378301015290565b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760043567ffffffffffffffff811161022a573660238201121561022a57806004013590610327826101f4565b906103356040519283610174565b8282526024602083019360051b8201019036821161022a5760248101935b8285106103655761036384611649565b005b843567ffffffffffffffff811161022a57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc823603011261022a57604051916103b18361011b565b60248201356103bf8161020c565b835260448201356103cf8161022f565b602084015260648201356103e281610241565b604084015260848201359267ffffffffffffffff841161022a57610410602094936024869536920101610285565b6060820152815201940193610353565b600091031261022a57565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5761046261189c565b506105986040516104728161011b565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b919082608091031261022a576040516105b48161011b565b60608082948035845260208101356105cb8161022f565b602085015260408101356105de8161022f565b60408501520135916105ef8361022f565b0152565b35906101c48261020c565b63ffffffff81160361022a57565b35906101c4826105fe565b81601f8201121561022a5780359061062e826101f4565b9261063c6040519485610174565b82845260208085019360051b8301019181831161022a5760208101935b83851061066857505050505090565b843567ffffffffffffffff811161022a57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603011261022a57604051916106b48361011b565b602082013567ffffffffffffffff811161022a578560206106d792850101610285565b835260408201356106e78161020c565b602084015260608201359267ffffffffffffffff841161022a57608083610715886020809881980101610285565b604084015201356060820152815201940193610659565b9080601f8301121561022a578135610743816101f4565b926107516040519485610174565b81845260208085019260051b82010192831161022a57602001905b8282106107795750505090565b6020809183356107888161020c565b81520191019061076c565b9190916101408184031261022a576107a96101b5565b926107b4818361059c565b8452608082013567ffffffffffffffff811161022a57816107d6918401610285565b602085015260a082013567ffffffffffffffff811161022a57816107fb918401610285565b604085015261080c60c083016105f3565b606085015261081d60e0830161060c565b608085015261010082013567ffffffffffffffff811161022a5781610843918401610617565b60a085015261012082013567ffffffffffffffff811161022a57610867920161072c565b60c0830152565b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760043567ffffffffffffffff811161022a576108c0610363913690600401610793565b611b54565b604051906108d4602083610174565b60008252565b60005b8381106108ed5750506000910152565b81810151838201526020016108dd565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610939815180928187528780880191016108da565b0116010190565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5761059860408051906109818183610174565b601782527f43435641676772656761746f7220312e372e302d6465760000000000000000006020830152519182916020835260208301906108fd565b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a577f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b36610a76604051610a1c8161013c565b600435610a2881610241565b8152610a32612093565b51600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151560ff81169290921790556040519081529081906020820190565b0390a1005b6040810160408252825180915260206060830193019060005b818110610b055750505060208183039101526020808351928381520192019060005b818110610ac35750505090565b9091926020604082610afa60019488516020809173ffffffffffffffffffffffffffffffffffffffff815116845201511515910152565b019401929101610ab6565b825167ffffffffffffffff16855260209485019490920191600101610a94565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a57600354610b60816101f4565b90610b6e6040519283610174565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610b9b826101f4565b0160005b818110610c61575050610bb181611e1b565b9060005b818110610bcd57505061059860405192839283610a7b565b80610c05610bec610bdf60019461352e565b67ffffffffffffffff1690565b610bf68387611e99565b9067ffffffffffffffff169052565b610c45610c40610c26610c188488611e99565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b611eb2565b610c4f8287611e99565b52610c5a8186611e99565b5001610bb5565b602090610c6c6118c1565b82828701015201610b9f565b60041115610c8257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b906004821015610c825752565b3461022a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a576020610d10600435610cfe8161022f565b60243590610d0b8261022f565b611f2a565b610d1d6040518092610cb1565bf35b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a576000604051610d5c8161013c565b52610598604051610d6c8161013c565b60025460ff161515908190526040519081529081906020820190565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760005473ffffffffffffffffffffffffffffffffffffffff81163303610e47577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5767ffffffffffffffff600435610f078161022f565b610f0f6118c1565b50166000526005602052610598604060002060ff60405191610f3083610158565b5473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015260405191829182815173ffffffffffffffffffffffffffffffffffffffff16815260209182015115159181019190915260400190565b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5773ffffffffffffffffffffffffffffffffffffffff600435610fdb8161020c565b610fe3612093565b1633811461105557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760043567ffffffffffffffff811161022a578060040160807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261022a5760025460ff1661161f5761112a60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006002541617600255565b61113f60206111398380611f80565b01611fb3565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa90811561161a576000916115eb575b506115b35761121161120d6112038467ffffffffffffffff166000526005602052604060002090565b5460a01c60ff1690565b1590565b61157b5761122460406111398380611f80565b67ffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036115335761126f61126960606111398480611f80565b83611f2a565b6112a08161127d8480611f80565b61128a6024880186611fd2565b9061129860448a0188611fd2565b949093612a77565b6112ba6112b56112b08480611f80565b612026565b612bef565b906112c481610c78565b80159081159182611520575b156114c957606461135496019363ffffffff6112eb86612031565b166114b45761132b61130b60e06113058461135d95611f80565b01612031565b8551606001516113259067ffffffffffffffff1689612c40565b85613008565b97909561134f87611349606089510167ffffffffffffffff90511690565b8a612ccb565b612031565b63ffffffff1690565b611452575b505061136d82610c78565b600282036113fe575b67ffffffffffffffff7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df291516113ce6113ba606083015167ffffffffffffffff1690565b915196836040519485941697169583612075565b0390a46103637fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060025416600255565b61140782610c78565b600382031561137657516060015161144e929067ffffffffffffffff16907f926c5a3e00000000000000000000000000000000000000000000000000000000600052612052565b6000fd5b61145b84610c78565b600384036113625761146c90610c78565b611477573880611362565b839051516114b06040519283927f2b11b8d90000000000000000000000000000000000000000000000000000000084526004840161203b565b0390fd5b5061135d61132b6114c486612031565b61130b565b61144e856114e4606086510167ffffffffffffffff90511690565b7f3b5754190000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff91821660045216602452604490565b5061152a81610c78565b600381146112d0565b61154560406111398361144e94611f80565b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b61160d915060203d602011611613575b6116058183610174565b810190611fbd565b386111da565b503d6115fb565b611b34565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b611651612093565b60005b8151811015611898576116678183611e99565b519061167e602083015167ffffffffffffffff1690565b67ffffffffffffffff811690811561186e576116cd6116b46116b4865173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b1561182e576116f09067ffffffffffffffff166000526005602052604060002090565b606084015180518015918215611858575b505061182e576118257f734ce292835471880fd796297a6bef2e23db07c84d09cad27f67cd7163f44035916117e66117a58761178b611745604060019b0151151590565b85547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178555565b5173ffffffffffffffffffffffffffffffffffffffff1690565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6117ef84613689565b5060408051915473ffffffffffffffffffffffffffffffffffffffff8116835260a01c60ff161515602083015290918291820190565b0390a201611654565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6020012090506118666120de565b143880611701565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b604051906118a98261011b565b60006060838281528260208201528260408201520152565b604051906118ce82610158565b60006020838281520152565b604051906118e9602083610174565b600080835282815b8281106118fd57505050565b6020906119086118c1565b828285010152016118f1565b9061191e826101f4565b61192b6040519182610174565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061195982946101f4565b019060005b82811061196a57505050565b6020906119756118c1565b8282850101520161195e565b909160608284031261022a57815161199881610241565b92602083015167ffffffffffffffff811161022a5783019080601f8301121561022a578151916119c78361024b565b916119d56040519384610174565b8383526020848301011161022a576040926119f691602080850191016108da565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a08401526080611a73611a3f604084015160a060c08801526101208701906108fd565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526108fd565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b818110611afc5750505061ffff90951660208301526101c49291606091611ae09063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101611ab2565b6040513d6000823e3d90fd5b906020611b519281815201906108fd565b90565b303303611df157611b636118da565b9060a08101518051611da0575b50805191611b8b6020845194015167ffffffffffffffff1690565b906020830151916040840192611bb9845192611ba56101c6565b97885267ffffffffffffffff166020880152565b60408601526060850152608084015251511580611d82575b8015611d5f575b8015611d2d575b61189857611cb19181611c246116b4611c0a610c266020600097510167ffffffffffffffff90511690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9083611c586060611c3c608085015163ffffffff1690565b93015173ffffffffffffffffffffffffffffffffffffffff1690565b93604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f000000000000000000000000000000000000000000000000000000000000000090600486016119fc565b03925af190811561161a57600090600092611d06575b5015611cd05750565b6114b0906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611b40565b9050611d2591503d806000833e611d1d8183610174565b810190611981565b509038611cc7565b50611d5a61120d611d55606084015173ffffffffffffffffffffffffffffffffffffffff1690565b6124ef565b611bdf565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b15611bd8565b5063ffffffff611d99608083015163ffffffff1690565b1615611bd1565b8192506020611dea920151611dcc606085015173ffffffffffffffffffffffffffffffffffffffff1690565b90611de4602086510167ffffffffffffffff90511690565b926120ff565b9038611b70565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b90611e25826101f4565b611e326040519182610174565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611e6082946101f4565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051821015611ead5760209160051b010190565b611e6a565b90604051611ebf81610158565b915473ffffffffffffffffffffffffffffffffffffffff8116835260a01c60ff1615156020830152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211611f2557565b611ee9565b611f3682607f92612599565b9116906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611f25576003911c166004811015610c825790565b604051906108d48261013c565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec18136030182121561022a570190565b35611b518161022f565b9081602091031261022a5751611b5181610241565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561022a570180359067ffffffffffffffff821161022a57602001918160051b3603831361022a57565b611b51903690610793565b35611b51816105fe565b604090611b519392815281602082015201906108fd565b929160449067ffffffffffffffff6101c493816064971660045216602452610cb1565b80612086604092611b519594610cb1565b81602082015201906108fd565b73ffffffffffffffffffffffffffffffffffffffff6001541633036120b457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b604051602081019060008252602081526120f9604082610174565b51902090565b9390919361210d8151611914565b9260009573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016925b80518810156124e45761215d8882611e99565b51976121676118c1565b506121e061218c60208b015173ffffffffffffffffffffffffffffffffffffffff1690565b9960208b604051809481927fbbe4f6db0000000000000000000000000000000000000000000000000000000083526004830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b03818a5afa91821561161a576000926124b4575b5073ffffffffffffffffffffffffffffffffffffffff82169182156124715761221c81612545565b6124715761222c61120d8261256f565b61247157506123046020878c8e946122ba612247878c61371c565b96612250611f73565b50606083015161227c60408551950151956122696101d5565b97885267ffffffffffffffff1688880152565b73ffffffffffffffffffffffffffffffffffffffff8d166040870152606086015273ffffffffffffffffffffffffffffffffffffffff166080850152565b60a083015260c08201526122cc6108c5565b60e0820152604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301613208565b03816000875af160009181612441575b50612358578b612322613331565b906114b06040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401613361565b9a92939495969798999a9173ffffffffffffffffffffffffffffffffffffffff8716036123d6575b509060019291516123ae6123926101e5565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201526123bd828a611e99565b526123c88189611e99565b50019695949392919061214a565b6123e0838761371c565b90808210801561242d575b6123f55750612380565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b506124388183611f18565b835114156123eb565b61246391925060203d811161246a575b61245b8183610174565b8101906131ea565b9038612314565b503d612451565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6124d691925060203d81116124dd575b6124ce8183610174565b8101906131d5565b90386121f4565b503d6124c4565b505050509250905090565b6125197f85572ffb00000000000000000000000000000000000000000000000000000000826134b4565b9081612533575b81612529575090565b611b519150613454565b905061253e8161338e565b1590612520565b6125197ff208a58f00000000000000000000000000000000000000000000000000000000826134b4565b6125197faff2afbf00000000000000000000000000000000000000000000000000000000826134b4565b9067ffffffffffffffff6125db921660005260066020526701ffffffffffffff60406000209160071c1667ffffffffffffffff16600052602052604060002090565b5490565b9190811015611ead5760051b0190565b35611b518161020c565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561022a57016020813591019167ffffffffffffffff821161022a57813603831361022a57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561022a57016020813591019167ffffffffffffffff821161022a578160051b3603831361022a57565b90602083828152019160208260051b8501019381936000915b8483106127045750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818636030182121561022a57602080918760019401906060806127cd61278b61277d86806125f9565b608087526080870191612649565b73ffffffffffffffffffffffffffffffffffffffff878701356127ad8161020c565b16878601526127bf60408701876125f9565b908683036040880152612649565b93013591015298019301930191949392906126f4565b9160209082815201919060005b8181106127fd5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356128268161020c565b1681520194019291016127f0565b90611b5191602081528135602082015267ffffffffffffffff602083013561285b8161022f565b16604082015267ffffffffffffffff60408301356128788161022f565b16606082015267ffffffffffffffff60608301356128958161022f565b16608082015261299861298c6129046128c76128b460808701876125f9565b61014060a0880152610160870191612649565b6128d460a08701876125f9565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c0880152612649565b61293061291360c087016105f3565b73ffffffffffffffffffffffffffffffffffffffff1660e0860152565b61294d61293f60e0870161060c565b63ffffffff16610100860152565b61295b610100860186612688565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403016101208701526126db565b92610120810190612688565b916101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101526127e3565b9190811015611ead5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561022a57019081359167ffffffffffffffff831161022a57602001823603811361022a579190565b9294612a5c612a6a936101c4989a9997612a4e6080989560a0895260a08901906108fd565b918783036020890152612649565b918483036040860152612649565b9560608201520190610cb1565b9195949295939093610120830184612a8f8286611fd2565b905003612bb15760005b612aa38286611fd2565b9050811015612ba657612ace6116b46116b4612ac984612ac3878b611fd2565b906125df565b6125ef565b90604051612b0f81612ae38960208301612834565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610174565b612b1a828c876129c9565b9093612b27848b8a6129c9565b829691963b1561022a57600094612b718e8888946040519b8c998a9889977fb9429a6a00000000000000000000000000000000000000000000000000000000895260048901612a29565b03925af191821561161a57600192612b8b575b5001612a99565b80612b9a6000612ba093610174565b80610420565b38612b84565b505050505050509050565b84612bbf61144e9286611fd2565b7f301b70fa0000000000000000000000000000000000000000000000000000000060005260045250602452604490565b60405160e0810181811067ffffffffffffffff8211176101375760609160c091604052612c1a61189c565b8152826020820152826040820152600083820152600060808201528260a0820152015290565b607f8216906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611f2557612cc89167ffffffffffffffff612c858584612599565b921660005260066020526701ffffffffffffff60406000209460071c169160036001831b921b191617929067ffffffffffffffff16600052602052604060002090565b55565b9091607f8316916801fffffffffffffffe67ffffffffffffffff84169360011b169280840460021490151715611f2557612d058482612599565b926004831015610c8257612cc89367ffffffffffffffff612d4a931660005260066020526003604060002094831b921b191617936701ffffffffffffff9060071c1690565b67ffffffffffffffff16600052602052604060002090565b9080602083519182815201916020808360051b8301019401926000915b838310612d8e57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08560019503018652885190606080612e0c612ddc85516080865260808601906108fd565b73ffffffffffffffffffffffffffffffffffffffff878701511687860152604086015185820360408701526108fd565b93015191015297019301930191939290612d7f565b906020808351928381520192019060005b818110612e3f5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612e32565b90611b519160208152612eae60208201835167ffffffffffffffff6060809280518552826020820151166020860152826040820151166040860152015116910152565b60c0612f6a612f04612ed1602086015161014060a08701526101608601906108fd565b60408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086830301858701526108fd565b606085015173ffffffffffffffffffffffffffffffffffffffff1660e0850152608085015163ffffffff1661010085015260a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301610120860152612d62565b920151906101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152612e21565b90602082519201517fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612fd6575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90612ae3613048613078936040519283917f0cea69ee00000000000000000000000000000000000000000000000000000000602084015260248301612e6b565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000009216903090613563565b50901561308b5750600290611b516108c5565b9072c11c11c11c11c11c11c11c11c11c11c11c11c133146130ad575b60039190565b6130de6130b983612f9e565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f37c3be29000000000000000000000000000000000000000000000000000000001480156131a1575b801561316d575b156130a75761144e61311f83612f9e565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b5061317a6130b983612f9e565b7fea7f4b12000000000000000000000000000000000000000000000000000000001461310e565b506131ae6130b983612f9e565b7fafa32a2c0000000000000000000000000000000000000000000000000000000014613107565b9081602091031261022a5751611b518161020c565b9081602091031261022a57604051906132028261013c565b51815290565b90611b51916020815260e06132fd6132ca613231855161010060208701526101208601906108fd565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff16606086015260608601516080860152613296608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526108fd565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526108fd565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526108fd565b3d1561335c573d906133428261024b565b916133506040519384610174565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff611b51949316815281602082015201906108fd565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526133ee604483610174565b6179185a1061342a576020926000925191617530fa6000513d8261341e575b5081613417575090565b9050151590565b6020111591503861340d565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526133ee604483610174565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526133ee604483610174565b8054821015611ead5760005260206000200190600090565b600354811015611ead5760036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b015490565b939193613570608461024b565b9461357e6040519687610174565b6084865261358c608461024b565b947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602088019601368737833b1561365f575a90808210613635578291038060061c9003111561360b576000918291825a9560208451940192f1905a9003923d9060848211613602575b6000908287523e929190565b608491506135f6565b7f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b806000526004602052604060002054156000146137165760035468010000000000000000811015610137578060016136c692016003556003613516565b81549060031b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84831b921b191617905561371060035491600490600052602052604060002090565b55600190565b50600090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161379b575b506137975782612322613331565b9150565b90916020823d6020116137ca575b816137b660209383610174565b810103126137c75750519038613789565b80fd5b3d91506137a956fea164736f6c634300081a000a",
}

var CCVAggregatorABI = CCVAggregatorMetaData.ABI

var CCVAggregatorBin = CCVAggregatorMetaData.Bin

func DeployCCVAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig CCVAggregatorStaticConfig, dynamicConfig CCVAggregatorDynamicConfig, sourceChainConfigs []CCVAggregatorSourceChainConfigArgs) (common.Address, *types.Transaction, *CCVAggregator, error) {
	parsed, err := CCVAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCVAggregatorBin), backend, staticConfig, dynamicConfig, sourceChainConfigs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCVAggregator{address: address, abi: *parsed, CCVAggregatorCaller: CCVAggregatorCaller{contract: contract}, CCVAggregatorTransactor: CCVAggregatorTransactor{contract: contract}, CCVAggregatorFilterer: CCVAggregatorFilterer{contract: contract}}, nil
}

type CCVAggregator struct {
	address common.Address
	abi     abi.ABI
	CCVAggregatorCaller
	CCVAggregatorTransactor
	CCVAggregatorFilterer
}

type CCVAggregatorCaller struct {
	contract *bind.BoundContract
}

type CCVAggregatorTransactor struct {
	contract *bind.BoundContract
}

type CCVAggregatorFilterer struct {
	contract *bind.BoundContract
}

type CCVAggregatorSession struct {
	Contract     *CCVAggregator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCVAggregatorCallerSession struct {
	Contract *CCVAggregatorCaller
	CallOpts bind.CallOpts
}

type CCVAggregatorTransactorSession struct {
	Contract     *CCVAggregatorTransactor
	TransactOpts bind.TransactOpts
}

type CCVAggregatorRaw struct {
	Contract *CCVAggregator
}

type CCVAggregatorCallerRaw struct {
	Contract *CCVAggregatorCaller
}

type CCVAggregatorTransactorRaw struct {
	Contract *CCVAggregatorTransactor
}

func NewCCVAggregator(address common.Address, backend bind.ContractBackend) (*CCVAggregator, error) {
	abi, err := abi.JSON(strings.NewReader(CCVAggregatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCVAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCVAggregator{address: address, abi: abi, CCVAggregatorCaller: CCVAggregatorCaller{contract: contract}, CCVAggregatorTransactor: CCVAggregatorTransactor{contract: contract}, CCVAggregatorFilterer: CCVAggregatorFilterer{contract: contract}}, nil
}

func NewCCVAggregatorCaller(address common.Address, caller bind.ContractCaller) (*CCVAggregatorCaller, error) {
	contract, err := bindCCVAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorCaller{contract: contract}, nil
}

func NewCCVAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CCVAggregatorTransactor, error) {
	contract, err := bindCCVAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorTransactor{contract: contract}, nil
}

func NewCCVAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CCVAggregatorFilterer, error) {
	contract, err := bindCCVAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorFilterer{contract: contract}, nil
}

func bindCCVAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCVAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCVAggregator *CCVAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCVAggregator.Contract.CCVAggregatorCaller.contract.Call(opts, result, method, params...)
}

func (_CCVAggregator *CCVAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVAggregator.Contract.CCVAggregatorTransactor.contract.Transfer(opts)
}

func (_CCVAggregator *CCVAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCVAggregator.Contract.CCVAggregatorTransactor.contract.Transact(opts, method, params...)
}

func (_CCVAggregator *CCVAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCVAggregator.Contract.contract.Call(opts, result, method, params...)
}

func (_CCVAggregator *CCVAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVAggregator.Contract.contract.Transfer(opts)
}

func (_CCVAggregator *CCVAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCVAggregator.Contract.contract.Transact(opts, method, params...)
}

func (_CCVAggregator *CCVAggregatorCaller) GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []CCVAggregatorSourceChainConfig, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "getAllSourceChainConfigs")

	if err != nil {
		return *new([]uint64), *new([]CCVAggregatorSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]CCVAggregatorSourceChainConfig)).(*[]CCVAggregatorSourceChainConfig)

	return out0, out1, err

}

func (_CCVAggregator *CCVAggregatorSession) GetAllSourceChainConfigs() ([]uint64, []CCVAggregatorSourceChainConfig, error) {
	return _CCVAggregator.Contract.GetAllSourceChainConfigs(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCallerSession) GetAllSourceChainConfigs() ([]uint64, []CCVAggregatorSourceChainConfig, error) {
	return _CCVAggregator.Contract.GetAllSourceChainConfigs(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCaller) GetDynamicConfig(opts *bind.CallOpts) (CCVAggregatorDynamicConfig, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CCVAggregatorDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCVAggregatorDynamicConfig)).(*CCVAggregatorDynamicConfig)

	return out0, err

}

func (_CCVAggregator *CCVAggregatorSession) GetDynamicConfig() (CCVAggregatorDynamicConfig, error) {
	return _CCVAggregator.Contract.GetDynamicConfig(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCallerSession) GetDynamicConfig() (CCVAggregatorDynamicConfig, error) {
	return _CCVAggregator.Contract.GetDynamicConfig(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_CCVAggregator *CCVAggregatorSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _CCVAggregator.Contract.GetExecutionState(&_CCVAggregator.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_CCVAggregator *CCVAggregatorCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _CCVAggregator.Contract.GetExecutionState(&_CCVAggregator.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_CCVAggregator *CCVAggregatorCaller) GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (CCVAggregatorSourceChainConfig, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "getSourceChainConfig", sourceChainSelector)

	if err != nil {
		return *new(CCVAggregatorSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCVAggregatorSourceChainConfig)).(*CCVAggregatorSourceChainConfig)

	return out0, err

}

func (_CCVAggregator *CCVAggregatorSession) GetSourceChainConfig(sourceChainSelector uint64) (CCVAggregatorSourceChainConfig, error) {
	return _CCVAggregator.Contract.GetSourceChainConfig(&_CCVAggregator.CallOpts, sourceChainSelector)
}

func (_CCVAggregator *CCVAggregatorCallerSession) GetSourceChainConfig(sourceChainSelector uint64) (CCVAggregatorSourceChainConfig, error) {
	return _CCVAggregator.Contract.GetSourceChainConfig(&_CCVAggregator.CallOpts, sourceChainSelector)
}

func (_CCVAggregator *CCVAggregatorCaller) GetStaticConfig(opts *bind.CallOpts) (CCVAggregatorStaticConfig, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(CCVAggregatorStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCVAggregatorStaticConfig)).(*CCVAggregatorStaticConfig)

	return out0, err

}

func (_CCVAggregator *CCVAggregatorSession) GetStaticConfig() (CCVAggregatorStaticConfig, error) {
	return _CCVAggregator.Contract.GetStaticConfig(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCallerSession) GetStaticConfig() (CCVAggregatorStaticConfig, error) {
	return _CCVAggregator.Contract.GetStaticConfig(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCVAggregator *CCVAggregatorSession) Owner() (common.Address, error) {
	return _CCVAggregator.Contract.Owner(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCallerSession) Owner() (common.Address, error) {
	return _CCVAggregator.Contract.Owner(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCVAggregator *CCVAggregatorSession) TypeAndVersion() (string, error) {
	return _CCVAggregator.Contract.TypeAndVersion(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorCallerSession) TypeAndVersion() (string, error) {
	return _CCVAggregator.Contract.TypeAndVersion(&_CCVAggregator.CallOpts)
}

func (_CCVAggregator *CCVAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "acceptOwnership")
}

func (_CCVAggregator *CCVAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCVAggregator.Contract.AcceptOwnership(&_CCVAggregator.TransactOpts)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCVAggregator.Contract.AcceptOwnership(&_CCVAggregator.TransactOpts)
}

func (_CCVAggregator *CCVAggregatorTransactor) ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []CCVAggregatorSourceChainConfigArgs) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "applySourceChainConfigUpdates", sourceChainConfigUpdates)
}

func (_CCVAggregator *CCVAggregatorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []CCVAggregatorSourceChainConfigArgs) (*types.Transaction, error) {
	return _CCVAggregator.Contract.ApplySourceChainConfigUpdates(&_CCVAggregator.TransactOpts, sourceChainConfigUpdates)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []CCVAggregatorSourceChainConfigArgs) (*types.Transaction, error) {
	return _CCVAggregator.Contract.ApplySourceChainConfigUpdates(&_CCVAggregator.TransactOpts, sourceChainConfigUpdates)
}

func (_CCVAggregator *CCVAggregatorTransactor) Execute(opts *bind.TransactOpts, report CCVAggregatorAggregatedReport) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "execute", report)
}

func (_CCVAggregator *CCVAggregatorSession) Execute(report CCVAggregatorAggregatedReport) (*types.Transaction, error) {
	return _CCVAggregator.Contract.Execute(&_CCVAggregator.TransactOpts, report)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) Execute(report CCVAggregatorAggregatedReport) (*types.Transaction, error) {
	return _CCVAggregator.Contract.Execute(&_CCVAggregator.TransactOpts, report)
}

func (_CCVAggregator *CCVAggregatorTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "executeSingleMessage", message)
}

func (_CCVAggregator *CCVAggregatorSession) ExecuteSingleMessage(message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _CCVAggregator.Contract.ExecuteSingleMessage(&_CCVAggregator.TransactOpts, message)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) ExecuteSingleMessage(message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _CCVAggregator.Contract.ExecuteSingleMessage(&_CCVAggregator.TransactOpts, message)
}

func (_CCVAggregator *CCVAggregatorTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCVAggregatorDynamicConfig) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CCVAggregator *CCVAggregatorSession) SetDynamicConfig(dynamicConfig CCVAggregatorDynamicConfig) (*types.Transaction, error) {
	return _CCVAggregator.Contract.SetDynamicConfig(&_CCVAggregator.TransactOpts, dynamicConfig)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) SetDynamicConfig(dynamicConfig CCVAggregatorDynamicConfig) (*types.Transaction, error) {
	return _CCVAggregator.Contract.SetDynamicConfig(&_CCVAggregator.TransactOpts, dynamicConfig)
}

func (_CCVAggregator *CCVAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "transferOwnership", to)
}

func (_CCVAggregator *CCVAggregatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCVAggregator.Contract.TransferOwnership(&_CCVAggregator.TransactOpts, to)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCVAggregator.Contract.TransferOwnership(&_CCVAggregator.TransactOpts, to)
}

type CCVAggregatorDynamicConfigSetIterator struct {
	Event *CCVAggregatorDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVAggregatorDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVAggregatorDynamicConfigSet)
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
		it.Event = new(CCVAggregatorDynamicConfigSet)
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

func (it *CCVAggregatorDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCVAggregatorDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVAggregatorDynamicConfigSet struct {
	DynamicConfig CCVAggregatorDynamicConfig
	Raw           types.Log
}

func (_CCVAggregator *CCVAggregatorFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCVAggregatorDynamicConfigSetIterator, error) {

	logs, sub, err := _CCVAggregator.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorDynamicConfigSetIterator{contract: _CCVAggregator.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_CCVAggregator *CCVAggregatorFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCVAggregatorDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCVAggregator.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVAggregatorDynamicConfigSet)
				if err := _CCVAggregator.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_CCVAggregator *CCVAggregatorFilterer) ParseDynamicConfigSet(log types.Log) (*CCVAggregatorDynamicConfigSet, error) {
	event := new(CCVAggregatorDynamicConfigSet)
	if err := _CCVAggregator.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVAggregatorExecutionStateChangedIterator struct {
	Event *CCVAggregatorExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVAggregatorExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVAggregatorExecutionStateChanged)
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
		it.Event = new(CCVAggregatorExecutionStateChanged)
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

func (it *CCVAggregatorExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *CCVAggregatorExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVAggregatorExecutionStateChanged struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	MessageId           [32]byte
	State               uint8
	ReturnData          []byte
	Raw                 types.Log
}

func (_CCVAggregator *CCVAggregatorFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*CCVAggregatorExecutionStateChangedIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCVAggregator.contract.FilterLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorExecutionStateChangedIterator{contract: _CCVAggregator.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_CCVAggregator *CCVAggregatorFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *CCVAggregatorExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _CCVAggregator.contract.WatchLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVAggregatorExecutionStateChanged)
				if err := _CCVAggregator.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_CCVAggregator *CCVAggregatorFilterer) ParseExecutionStateChanged(log types.Log) (*CCVAggregatorExecutionStateChanged, error) {
	event := new(CCVAggregatorExecutionStateChanged)
	if err := _CCVAggregator.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVAggregatorOwnershipTransferRequestedIterator struct {
	Event *CCVAggregatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVAggregatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVAggregatorOwnershipTransferRequested)
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
		it.Event = new(CCVAggregatorOwnershipTransferRequested)
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

func (it *CCVAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCVAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCVAggregator *CCVAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorOwnershipTransferRequestedIterator{contract: _CCVAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCVAggregator *CCVAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCVAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVAggregatorOwnershipTransferRequested)
				if err := _CCVAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCVAggregator *CCVAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCVAggregatorOwnershipTransferRequested, error) {
	event := new(CCVAggregatorOwnershipTransferRequested)
	if err := _CCVAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVAggregatorOwnershipTransferredIterator struct {
	Event *CCVAggregatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVAggregatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVAggregatorOwnershipTransferred)
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
		it.Event = new(CCVAggregatorOwnershipTransferred)
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

func (it *CCVAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCVAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCVAggregator *CCVAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorOwnershipTransferredIterator{contract: _CCVAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCVAggregator *CCVAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCVAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVAggregatorOwnershipTransferred)
				if err := _CCVAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCVAggregator *CCVAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*CCVAggregatorOwnershipTransferred, error) {
	event := new(CCVAggregatorOwnershipTransferred)
	if err := _CCVAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVAggregatorSourceChainConfigSetIterator struct {
	Event *CCVAggregatorSourceChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVAggregatorSourceChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVAggregatorSourceChainConfigSet)
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
		it.Event = new(CCVAggregatorSourceChainConfigSet)
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

func (it *CCVAggregatorSourceChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCVAggregatorSourceChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVAggregatorSourceChainConfigSet struct {
	SourceChainSelector uint64
	SourceConfig        CCVAggregatorSourceChainConfig
	Raw                 types.Log
}

func (_CCVAggregator *CCVAggregatorFilterer) FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*CCVAggregatorSourceChainConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _CCVAggregator.contract.FilterLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorSourceChainConfigSetIterator{contract: _CCVAggregator.contract, event: "SourceChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CCVAggregator *CCVAggregatorFilterer) WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCVAggregatorSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _CCVAggregator.contract.WatchLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVAggregatorSourceChainConfigSet)
				if err := _CCVAggregator.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
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

func (_CCVAggregator *CCVAggregatorFilterer) ParseSourceChainConfigSet(log types.Log) (*CCVAggregatorSourceChainConfigSet, error) {
	event := new(CCVAggregatorSourceChainConfigSet)
	if err := _CCVAggregator.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVAggregatorStaticConfigSetIterator struct {
	Event *CCVAggregatorStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVAggregatorStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVAggregatorStaticConfigSet)
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
		it.Event = new(CCVAggregatorStaticConfigSet)
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

func (it *CCVAggregatorStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCVAggregatorStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVAggregatorStaticConfigSet struct {
	StaticConfig CCVAggregatorStaticConfig
	Raw          types.Log
}

func (_CCVAggregator *CCVAggregatorFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*CCVAggregatorStaticConfigSetIterator, error) {

	logs, sub, err := _CCVAggregator.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCVAggregatorStaticConfigSetIterator{contract: _CCVAggregator.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_CCVAggregator *CCVAggregatorFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *CCVAggregatorStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCVAggregator.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVAggregatorStaticConfigSet)
				if err := _CCVAggregator.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_CCVAggregator *CCVAggregatorFilterer) ParseStaticConfigSet(log types.Log) (*CCVAggregatorStaticConfigSet, error) {
	event := new(CCVAggregatorStaticConfigSet)
	if err := _CCVAggregator.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_CCVAggregator *CCVAggregator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CCVAggregator.abi.Events["DynamicConfigSet"].ID:
		return _CCVAggregator.ParseDynamicConfigSet(log)
	case _CCVAggregator.abi.Events["ExecutionStateChanged"].ID:
		return _CCVAggregator.ParseExecutionStateChanged(log)
	case _CCVAggregator.abi.Events["OwnershipTransferRequested"].ID:
		return _CCVAggregator.ParseOwnershipTransferRequested(log)
	case _CCVAggregator.abi.Events["OwnershipTransferred"].ID:
		return _CCVAggregator.ParseOwnershipTransferred(log)
	case _CCVAggregator.abi.Events["SourceChainConfigSet"].ID:
		return _CCVAggregator.ParseSourceChainConfigSet(log)
	case _CCVAggregator.abi.Events["StaticConfigSet"].ID:
		return _CCVAggregator.ParseStaticConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CCVAggregatorDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b36")
}

func (CCVAggregatorExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2")
}

func (CCVAggregatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCVAggregatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CCVAggregatorSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x734ce292835471880fd796297a6bef2e23db07c84d09cad27f67cd7163f44035")
}

func (CCVAggregatorStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e4950")
}

func (_CCVAggregator *CCVAggregator) Address() common.Address {
	return _CCVAggregator.address
}

type CCVAggregatorInterface interface {
	GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []CCVAggregatorSourceChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (CCVAggregatorDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (CCVAggregatorSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (CCVAggregatorStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []CCVAggregatorSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report CCVAggregatorAggregatedReport) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMultiProofMessage) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCVAggregatorDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCVAggregatorDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCVAggregatorDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CCVAggregatorDynamicConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*CCVAggregatorExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *CCVAggregatorExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*CCVAggregatorExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVAggregatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCVAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCVAggregatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVAggregatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCVAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCVAggregatorOwnershipTransferred, error)

	FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*CCVAggregatorSourceChainConfigSetIterator, error)

	WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCVAggregatorSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSourceChainConfigSet(log types.Log) (*CCVAggregatorSourceChainConfigSet, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*CCVAggregatorStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *CCVAggregatorStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*CCVAggregatorStaticConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
