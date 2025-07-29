// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_aggregator

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

type VerifierAggregatorAggregatedReport struct {
	Message     InternalAny2EVMMultiProofMessage
	Proofs      [][]byte
	GasOverride uint32
}

type VerifierAggregatorDynamicConfig struct {
	ReentrancyGuardEntered bool
}

type VerifierAggregatorSourceChainConfig struct {
	Router    common.Address
	IsEnabled bool
}

type VerifierAggregatorSourceChainConfigArgs struct {
	Router              common.Address
	SourceChainSelector uint64
	IsEnabled           bool
	OnRamp              []byte
}

type VerifierAggregatorStaticConfig struct {
	LocalChainSelector   uint64
	GasForCallExactCheck uint16
	RmnRemote            common.Address
	TokenAdminRegistry   common.Address
}

var VerifierAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.AggregatedReport\",\"components\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"proofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.RequiredVerifier[]\",\"components\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101008060405234610520576140bc803803809161001d8285610556565b833981019080820360c08112610520576080811261052057604051916100428361053b565b61004b81610579565b8352602081015161ffff8116810361052057602084019081526040820151906001600160a01b038216820361052057604085019182526060830151936001600160a01b03851685036105205760209060608701958652607f1901126105205760405192602084016001600160401b03811185821017610525576040526100d36080820161058d565b845260a0810151906001600160401b038211610520570186601f82011215610520578051966001600160401b038811610525578760051b916040519861011c602085018b610556565b89526020808a0193820101908282116105205760208101935b828510610420575050505050331561040f57600180546001600160a01b0319163317905581516001600160a01b03161580156103fd575b6103ec5784516001600160401b0316156103db5784516001600160401b03908116608090815283516001600160a01b0390811660a0528651811660c052835161ffff90811660e05260408051995190941689529351909316602088810191909152935183169187019190915293511660608501527f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b369390927f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a151151560ff196002541660ff821617600255604051908152a16000905b80518210156103855760009160208160051b8301015160018060401b036020820151169081156103765780516001600160a01b031615610367578185526005602052604085209460608201518051801591821561033b575b505061032c5750604081810151865492516001600160a81b031990931690151560a01b60ff60a01b16176001600160a01b039290921691909117855592936001937f734ce292835471880fd796297a6bef2e23db07c84d09cad27f67cd7163f44035916103088461059a565b5060ff82519154878060a01b038116835260a01c1615156020820152a20190610244565b6342bcdf7f60e11b8152600490fd5b6020919250012060405160208101908382526020815261035c604082610556565b51902014388061029c565b6342bcdf7f60e11b8552600485fd5b63c656089560e01b8552600485fd5b604051613a8e908161062e82396080518181816104990152610d15015260a0518181816104fc0152610c79015260c0518181816105380152613013015260e0518181816104c001528181611fc90152612e640152f35b63c656089560e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b5083516001600160a01b03161561016c565b639b15e16f60e01b60005260046000fd5b84516001600160401b0381116105205782016080818603601f190112610520576040519061044d8261053b565b60208101516001600160a01b038116810361052057825261047060408201610579565b60208301526104816060820161058d565b604083015260808101516001600160401b03811161052057602091010185601f820112156105205780516001600160401b03811161052557604051916104d1601f8301601f191660200184610556565b81835287602083830101116105205760005b82811061050b5750509181600060208096949581960101526060820152815201940193610135565b806020809284010151828287010152016104e3565b600080fd5b634e487b7160e01b600052604160045260246000fd5b608081019081106001600160401b0382111761052557604052565b601f909101601f19168101906001600160401b0382119082101761052557604052565b51906001600160401b038216820361052057565b5190811515820361052057565b806000526004602052604060002054156000146106275760035468010000000000000000811015610525576001810180600355811015610611577fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0181905560035460009182526004602052604090912055600190565b634e487b7160e01b600052603260045260246000fd5b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806304666f9c146100e757806306285c69146100e2578063181f5a77146100dd5780634708209d146100d85780635215505b146100d35780635e36480c146100ce5780637437ff9f146100c957806379ba5097146100c45780638840c68d146100bf5780638da5cb5b146100ba57806395e1ee5d146100b5578063e9d68a8e146100b05763f2fde38b146100ab57600080fd5b61161e565b611556565b6114ff565b611120565b610b64565b610a7b565b610a12565b6109b1565b610818565b6106b0565b610633565b610447565b6102e8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761013757604052565b6100ec565b6020810190811067ffffffffffffffff82111761013757604052565b60a0810190811067ffffffffffffffff82111761013757604052565b6040810190811067ffffffffffffffff82111761013757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761013757604052565b604051906101e060e083610190565b565b604051906101e060a083610190565b604051906101e061010083610190565b604051906101e0604083610190565b67ffffffffffffffff81116101375760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361024657565b600080fd5b67ffffffffffffffff81160361024657565b8015150361024657565b67ffffffffffffffff811161013757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f82011215610246578035906102b882610267565b926102c66040519485610190565b8284526020838301011161024657816000926020809301838601378301015290565b346102465760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102465760043567ffffffffffffffff811161024657366023820112156102465780600401359061034382610210565b906103516040519283610190565b8282526024602083019360051b820101903682116102465760248101935b8285106103815761037f84611712565b005b843567ffffffffffffffff811161024657820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc823603011261024657604051916103cd8361011b565b60248201356103db81610228565b835260448201356103eb8161024b565b602084015260648201356103fe8161025d565b604084015260848201359267ffffffffffffffff84116102465761042c6020949360248695369201016102a1565b606082015281520194019361036f565b600091031261024657565b346102465760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102465761047e611965565b506105b460405161048e8161011b565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b604051906105c7602083610190565b60008252565b60005b8381106105e05750506000910152565b81810151838201526020016105d0565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361062c815180928187528780880191016105cd565b0116010190565b346102465760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610246576105b460408051906106748183610190565b601c82527f566572696669657241676772656761746f7220312e372e302d646576000000006020830152519182916020835260208301906105f0565b346102465760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610246577f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b3661076960405161070f8161013c565b60043561071b8161025d565b815261072561215c565b51600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151560ff81169290921790556040519081529081906020820190565b0390a1005b6040810160408252825180915260206060830193019060005b8181106107f85750505060208183039101526020808351928381520192019060005b8181106107b65750505090565b90919260206040826107ed60019488516020809173ffffffffffffffffffffffffffffffffffffffff815116845201511515910152565b0194019291016107a9565b825167ffffffffffffffff16855260209485019490920191600101610787565b346102465760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102465760035461085381610210565b906108616040519283610190565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061088e82610210565b0160005b8181106109545750506108a4816119a3565b9060005b8181106108c05750506105b46040519283928361076e565b806108f86108df6108d260019461349c565b67ffffffffffffffff1690565b6108e98387611a21565b9067ffffffffffffffff169052565b61093861093361091961090b8488611a21565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b611a3a565b6109428287611a21565b5261094d8186611a21565b50016108a8565b60209061095f61198a565b82828701015201610892565b6004111561097557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9060048210156109755752565b346102465760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610246576020610a036004356109f18161024b565b602435906109fe8261024b565b611ab2565b610a1060405180926109a4565bf35b346102465760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610246576000604051610a4f8161013c565b526105b4604051610a5f8161013c565b60025460ff161515908190526040519081529081906020820190565b346102465760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102465760005473ffffffffffffffffffffffffffffffffffffffff81163303610b3a577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102465760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102465760043567ffffffffffffffff8111610246578060040160607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126102465760025460ff166110f657610c0f60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006002541617600255565b610c246020610c1e8380611b08565b01611b3b565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa9081156110f1576000916110c2575b5061108a57610cf6610cf2610ce88467ffffffffffffffff166000526005602052604060002090565b5460a01c60ff1690565b1590565b61105257610d096040610c1e8380611b08565b67ffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361100a57610d54610d4e6060610c1e8480611b08565b83611ab2565b610d7781610d628480611b08565b610d6f6024880186611b69565b92909161276c565b610d91610d8c610d878480611b08565b611bbd565b61298f565b90610d9b8161096b565b80159081159182610ff7575b15610fa0576044610e2b96019363ffffffff610dc286611bc8565b16610f8b57610e02610de260e0610ddc84610e3495611b08565b01611bc8565b855160600151610dfc9067ffffffffffffffff16896129e0565b85612e1d565b979095610e2687610e20606089510167ffffffffffffffff90511690565b8a612a6b565b611bc8565b63ffffffff1690565b610f29575b5050610e448261096b565b60028203610ed5575b67ffffffffffffffff7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df29151610ea5610e91606083015167ffffffffffffffff1690565b915196836040519485941697169583611c0c565b0390a461037f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060025416600255565b610ede8261096b565b6003820315610e4d575160600151610f25929067ffffffffffffffff16907f926c5a3e00000000000000000000000000000000000000000000000000000000600052611be9565b6000fd5b610f328461096b565b60038403610e3957610f439061096b565b610f4e573880610e39565b83905151610f876040519283927f2b11b8d900000000000000000000000000000000000000000000000000000000845260048401611bd2565b0390fd5b50610e34610e02610f9b86611bc8565b610de2565b610f2585610fbb606086510167ffffffffffffffff90511690565b7f3b5754190000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff91821660045216602452604490565b506110018161096b565b60038114610da7565b61101c6040610c1e83610f2594611b08565b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6110e4915060203d6020116110ea575b6110dc8183610190565b810190611b48565b38610cbf565b503d6110d2565b611b5d565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346102465760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261024657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b91908260809103126102465760405161118a8161011b565b60608082948035845260208101356111a18161024b565b602085015260408101356111b48161024b565b60408501520135916111c58361024b565b0152565b35906101e082610228565b63ffffffff81160361024657565b35906101e0826111d4565b81601f820112156102465780359061120482610210565b926112126040519485610190565b82845260208085019360051b830101918183116102465760208101935b83851061123e57505050505090565b843567ffffffffffffffff811161024657820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030112610246576040519161128a8361011b565b602082013567ffffffffffffffff8111610246578560206112ad928501016102a1565b835260408201356112bd81610228565b602084015260608201359267ffffffffffffffff8411610246576080836112eb8860208098819801016102a1565b60408401520135606082015281520194019361122f565b9080601f830112156102465781359161131a83610210565b926113286040519485610190565b80845260208085019160051b830101918383116102465760208101915b83831061135457505050505090565b823567ffffffffffffffff81116102465782019060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838803011261024657604051906113a182610158565b60208301358252604083013567ffffffffffffffff8111610246578760206113cb928601016102a1565b60208301526060830135604083015260808301356113e88161024b565b606083015260a08301359167ffffffffffffffff831161024657611414886020809695819601016102a1565b6080820152815201920191611345565b919091610140818403126102465761143a6101d1565b926114458183611172565b8452608082013567ffffffffffffffff811161024657816114679184016102a1565b602085015260a082013567ffffffffffffffff8111610246578161148c9184016102a1565b604085015261149d60c083016111c9565b60608501526114ae60e083016111e2565b608085015261010082013567ffffffffffffffff811161024657816114d49184016111ed565b60a085015261012082013567ffffffffffffffff8111610246576114f89201611302565b60c0830152565b346102465760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102465760043567ffffffffffffffff81116102465761155161037f913690600401611424565b611e95565b346102465760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102465767ffffffffffffffff60043561159a8161024b565b6115a261198a565b501660005260056020526105b4604060002060ff604051916115c383610174565b5473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015260405191829182815173ffffffffffffffffffffffffffffffffffffffff16815260209182015115159181019190915260400190565b346102465760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102465773ffffffffffffffffffffffffffffffffffffffff60043561166e81610228565b61167661215c565b163381146116e857807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b61171a61215c565b60005b8151811015611961576117308183611a21565b5190611747602083015167ffffffffffffffff1690565b67ffffffffffffffff81169081156119375761179661177d61177d865173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b156118f7576117b99067ffffffffffffffff166000526005602052604060002090565b606084015180518015918215611921575b50506118f7576118ee7f734ce292835471880fd796297a6bef2e23db07c84d09cad27f67cd7163f44035916118af61186e8761185461180e604060019b0151151590565b85547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178555565b5173ffffffffffffffffffffffffffffffffffffffff1690565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6118b884613938565b5060408051915473ffffffffffffffffffffffffffffffffffffffff8116835260a01c60ff161515602083015290918291820190565b0390a20161171d565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b60200120905061192f6121a7565b1438806117ca565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b604051906119728261011b565b60006060838281528260208201528260408201520152565b6040519061199782610174565b60006020838281520152565b906119ad82610210565b6119ba6040519182610190565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06119e88294610210565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051821015611a355760209160051b010190565b6119f2565b90604051611a4781610174565b915473ffffffffffffffffffffffffffffffffffffffff8116835260a01c60ff1615156020830152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211611aad57565b611a71565b611abe82607f926121c8565b9116906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611aad576003911c1660048110156109755790565b604051906105c78261013c565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec181360301821215610246570190565b35611b458161024b565b90565b908160209103126102465751611b458161025d565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610246570180359067ffffffffffffffff821161024657602001918160051b3603831361024657565b611b45903690611424565b35611b45816111d4565b604090611b459392815281602082015201906105f0565b929160449067ffffffffffffffff6101e0938160649716600452166024526109a4565b80611c1d604092611b4595946109a4565b81602082015201906105f0565b60405190611c39602083610190565b600080835282815b828110611c4d57505050565b602090611c5861198a565b82828501015201611c41565b90611c6e82610210565b611c7b6040519182610190565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611ca98294610210565b019060005b828110611cba57505050565b602090611cc561198a565b82828501015201611cae565b9091606082840312610246578151611ce88161025d565b92602083015167ffffffffffffffff81116102465783019080601f8301121561024657815191611d1783610267565b91611d256040519384610190565b8383526020848301011161024657604092611d4691602080850191016105cd565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a08401526080611dc3611d8f604084015160a060c08801526101208701906105f0565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526105f0565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b818110611e4c5750505061ffff90951660208301526101e09291606091611e309063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101611e02565b906020611b459281815201906105f0565b30330361213257611ea4611c2a565b9060a081015180516120e1575b50805191611ecc6020845194015167ffffffffffffffff1690565b906020830151916040840192611efa845192611ee66101e2565b97885267ffffffffffffffff166020880152565b604086015260608501526080840152515115806120c3575b80156120a0575b801561206e575b61196157611ff29181611f6561177d611f4b6109196020600097510167ffffffffffffffff90511690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9083611f996060611f7d608085015163ffffffff1690565b93015173ffffffffffffffffffffffffffffffffffffffff1690565b93604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601611d4c565b03925af19081156110f157600090600092612047575b50156120115750565b610f87906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611e84565b905061206691503d806000833e61205e8183610190565b810190611cd1565b509038612008565b5061209b610cf2612096606084015173ffffffffffffffffffffffffffffffffffffffff1690565b6133da565b611f20565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b15611f19565b5063ffffffff6120da608083015163ffffffff1690565b1615611f12565b819250602061212b92015161210d606085015173ffffffffffffffffffffffffffffffffffffffff1690565b90612125602086510167ffffffffffffffff90511690565b92612fea565b9038611eb1565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361217d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b604051602081019060008252602081526121c2604082610190565b51902090565b9067ffffffffffffffff61220a921660005260076020526701ffffffffffffff60406000209160071c1667ffffffffffffffff16600052602052604060002090565b5490565b9190811015611a355760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610246570190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561024657016020813591019167ffffffffffffffff821161024657813603831361024657565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561024657016020813591019167ffffffffffffffff8211610246578160051b3603831361024657565b90602083828152019160208260051b8501019381936000915b8483106123595750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818636030182121561024657602080918760019401906060806124226123e06123d2868061224e565b60808752608087019161229e565b73ffffffffffffffffffffffffffffffffffffffff8787013561240281610228565b1687860152612414604087018761224e565b90868303604088015261229e565b9301359101529801930193019194939290612349565b90602083828152019060208160051b85010193836000915b8383106124605750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618436030181121561024657602061252e6001938683940190813581526125206124ee6124de8685018561224e565b60a08886015260a085019161229e565b926040810135604084015267ffffffffffffffff60608201356125108161024b565b166060840152608081019061224e565b91608081850391015261229e565b980196019493019190612450565b90611b4591602081528135602082015267ffffffffffffffff60208301356125638161024b565b16604082015267ffffffffffffffff60408301356125808161024b565b16606082015267ffffffffffffffff606083013561259d8161024b565b1660808201526126a061269461260c6125cf6125bc608087018761224e565b61014060a088015261016087019161229e565b6125dc60a087018761224e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c088015261229e565b61263861261b60c087016111c9565b73ffffffffffffffffffffffffffffffffffffffff1660e0860152565b61265561264760e087016111e2565b63ffffffff16610100860152565b6126636101008601866122dd565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086840301610120870152612330565b926101208101906122dd565b916101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301910152612438565b9190811015611a355760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561024657019081359167ffffffffffffffff8311610246576020018236038113610246579190565b906060926101e09597969461275161275f936080865260808601906105f0565b91848303602086015261229e565b95604082015201906109a4565b919290926101208301826127808286611b69565b9050036129515760005b6127948286611b69565b9050811015612949576127da6127b4826127ae8589611b69565b9061220e565b357fffffffff000000000000000000000000000000000000000000000000000000001690565b73ffffffffffffffffffffffffffffffffffffffff612828611f4b837fffffffff00000000000000000000000000000000000000000000000000000000166000526006602052604060002090565b169081156128fa575060405161287181612845896020830161253c565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610190565b61287c8387876126d1565b9092803b15610246578493600080946128c58d604051998a97889687957fcba4c71a00000000000000000000000000000000000000000000000000000000875260048701612731565b03925af19182156110f1576001926128df575b500161278a565b806128ee60006128f493610190565b8061043c565b386128d8565b7f29391c08000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b505050505050565b61295f9150610f2593611b69565b7f301b70fa0000000000000000000000000000000000000000000000000000000060005260045250602452604490565b60405160e0810181811067ffffffffffffffff8211176101375760609160c0916040526129ba611965565b8152826020820152826040820152600083820152600060808201528260a0820152015290565b607f8216906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611aad57612a689167ffffffffffffffff612a2585846121c8565b921660005260076020526701ffffffffffffff60406000209460071c169160036001831b921b191617929067ffffffffffffffff16600052602052604060002090565b55565b9091607f8316916801fffffffffffffffe67ffffffffffffffff84169360011b169280840460021490151715611aad57612aa584826121c8565b92600483101561097557612a689367ffffffffffffffff612aea931660005260076020526003604060002094831b921b191617936701ffffffffffffff9060071c1690565b67ffffffffffffffff16600052602052604060002090565b9080602083519182815201916020808360051b8301019401926000915b838310612b2e57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08560019503018652885190606080612bac612b7c85516080865260808601906105f0565b73ffffffffffffffffffffffffffffffffffffffff878701511687860152604086015185820360408701526105f0565b93015191015297019301930191939290612b1f565b9080602083519182815201916020808360051b8301019401926000915b838310612bed57505050505090565b9091929394602080612c71837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080612c428584015160a08785015260a08401906105f0565b926040810151604084015267ffffffffffffffff606082015116606084015201519060808184039101526105f0565b97019301930191939290612bde565b90611b459160208152612cc360208201835167ffffffffffffffff6060809280518552826020820151166020860152826040820151166040860152015116910152565b60c0612d7f612d19612ce6602086015161014060a08701526101608601906105f0565b60408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086830301858701526105f0565b606085015173ffffffffffffffffffffffffffffffffffffffff1660e0850152608085015163ffffffff1661010085015260a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301610120860152612b02565b920151906101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152612bc1565b90602082519201517fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612deb575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90612845612e5d612e8d936040519283917f95e1ee5d00000000000000000000000000000000000000000000000000000000602084015260248301612c80565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000092169030906134d1565b509015612ea05750600290611b456105b8565b9072c11c11c11c11c11c11c11c11c11c11c11c11c13314612ec2575b60039190565b612ef3612ece83612db3565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f37c3be2900000000000000000000000000000000000000000000000000000000148015612fb6575b8015612f82575b15612ebc57610f25612f3483612db3565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b50612f8f612ece83612db3565b7fea7f4b120000000000000000000000000000000000000000000000000000000014612f23565b50612fc3612ece83612db3565b7fafa32a2c0000000000000000000000000000000000000000000000000000000014612f1c565b93909193612ff88151611c64565b9260009573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016925b80518810156133cf576130488882611a21565b519761305261198a565b506130cb61307760208b015173ffffffffffffffffffffffffffffffffffffffff1690565b9960208b604051809481927fbbe4f6db0000000000000000000000000000000000000000000000000000000083526004830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b03818a5afa9182156110f15760009261339f575b5073ffffffffffffffffffffffffffffffffffffffff821691821561335c5761310781613430565b61335c57613117610cf28261345a565b61335c57506131ef6020878c8e946131a5613132878c6139cb565b9661313b611afb565b50606083015161316760408551950151956131546101f1565b97885267ffffffffffffffff1688880152565b73ffffffffffffffffffffffffffffffffffffffff8d166040870152606086015273ffffffffffffffffffffffffffffffffffffffff166080850152565b60a083015260c08201526131b76105b8565b60e0820152604051809381927f390775370000000000000000000000000000000000000000000000000000000083526004830161362a565b03816000875af16000918161332c575b50613243578b61320d613753565b90610f876040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401613783565b9a92939495969798999a9173ffffffffffffffffffffffffffffffffffffffff8716036132c1575b5090600192915161329961327d610201565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201526132a8828a611a21565b526132b38189611a21565b500196959493929190613035565b6132cb83876139cb565b908082108015613318575b6132e0575061326b565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b506133238183611aa0565b835114156132d6565b61334e91925060203d8111613355575b6133468183610190565b81019061360c565b90386131ff565b503d61333c565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6133c191925060203d81116133c8575b6133b98183610190565b8101906135f7565b90386130df565b503d6133af565b505050509250905090565b6134047f85572ffb00000000000000000000000000000000000000000000000000000000826138d6565b908161341e575b81613414575090565b611b459150613876565b9050613429816137b0565b159061340b565b6134047ff208a58f00000000000000000000000000000000000000000000000000000000826138d6565b6134047faff2afbf00000000000000000000000000000000000000000000000000000000826138d6565b8054821015611a355760005260206000200190600090565b600354811015611a355760036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b015490565b9391936134de6084610267565b946134ec6040519687610190565b608486526134fa6084610267565b947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602088019601368737833b156135cd575a908082106135a3578291038060061c90031115613579576000918291825a9560208451940192f1905a9003923d9060848211613570575b6000908287523e929190565b60849150613564565b7f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b908160209103126102465751611b4581610228565b9081602091031261024657604051906136248261013c565b51815290565b90611b45916020815260e061371f6136ec613653855161010060208701526101208601906105f0565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff166060860152606086015160808601526136b8608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526105f0565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526105f0565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526105f0565b3d1561377e573d9061376482610267565b916137726040519384610190565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff611b45949316815281602082015201906105f0565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252613810604483610190565b6179185a1061384c576020926000925191617530fa6000513d82613840575b5081613839575090565b9050151590565b6020111591503861382f565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252613810604483610190565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252613810604483610190565b806000526004602052604060002054156000146139c557600354680100000000000000008110156101375780600161397592016003556003613484565b81549060031b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84831b921b19161790556139bf60035491600490600052602052604060002090565b55600190565b50600090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181613a4a575b50613a46578261320d613753565b9150565b90916020823d602011613a79575b81613a6560209383610190565b81010312613a765750519038613a38565b80fd5b3d9150613a5856fea164736f6c634300081a000a",
}

var VerifierAggregatorABI = VerifierAggregatorMetaData.ABI

var VerifierAggregatorBin = VerifierAggregatorMetaData.Bin

func DeployVerifierAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig VerifierAggregatorStaticConfig, dynamicConfig VerifierAggregatorDynamicConfig, sourceChainConfigs []VerifierAggregatorSourceChainConfigArgs) (common.Address, *types.Transaction, *VerifierAggregator, error) {
	parsed, err := VerifierAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierAggregatorBin), backend, staticConfig, dynamicConfig, sourceChainConfigs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierAggregator{address: address, abi: *parsed, VerifierAggregatorCaller: VerifierAggregatorCaller{contract: contract}, VerifierAggregatorTransactor: VerifierAggregatorTransactor{contract: contract}, VerifierAggregatorFilterer: VerifierAggregatorFilterer{contract: contract}}, nil
}

type VerifierAggregator struct {
	address common.Address
	abi     abi.ABI
	VerifierAggregatorCaller
	VerifierAggregatorTransactor
	VerifierAggregatorFilterer
}

type VerifierAggregatorCaller struct {
	contract *bind.BoundContract
}

type VerifierAggregatorTransactor struct {
	contract *bind.BoundContract
}

type VerifierAggregatorFilterer struct {
	contract *bind.BoundContract
}

type VerifierAggregatorSession struct {
	Contract     *VerifierAggregator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierAggregatorCallerSession struct {
	Contract *VerifierAggregatorCaller
	CallOpts bind.CallOpts
}

type VerifierAggregatorTransactorSession struct {
	Contract     *VerifierAggregatorTransactor
	TransactOpts bind.TransactOpts
}

type VerifierAggregatorRaw struct {
	Contract *VerifierAggregator
}

type VerifierAggregatorCallerRaw struct {
	Contract *VerifierAggregatorCaller
}

type VerifierAggregatorTransactorRaw struct {
	Contract *VerifierAggregatorTransactor
}

func NewVerifierAggregator(address common.Address, backend bind.ContractBackend) (*VerifierAggregator, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierAggregatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierAggregator{address: address, abi: abi, VerifierAggregatorCaller: VerifierAggregatorCaller{contract: contract}, VerifierAggregatorTransactor: VerifierAggregatorTransactor{contract: contract}, VerifierAggregatorFilterer: VerifierAggregatorFilterer{contract: contract}}, nil
}

func NewVerifierAggregatorCaller(address common.Address, caller bind.ContractCaller) (*VerifierAggregatorCaller, error) {
	contract, err := bindVerifierAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorCaller{contract: contract}, nil
}

func NewVerifierAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierAggregatorTransactor, error) {
	contract, err := bindVerifierAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorTransactor{contract: contract}, nil
}

func NewVerifierAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierAggregatorFilterer, error) {
	contract, err := bindVerifierAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorFilterer{contract: contract}, nil
}

func bindVerifierAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierAggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierAggregator *VerifierAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierAggregator.Contract.VerifierAggregatorCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierAggregator *VerifierAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.VerifierAggregatorTransactor.contract.Transfer(opts)
}

func (_VerifierAggregator *VerifierAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.VerifierAggregatorTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierAggregator *VerifierAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierAggregator.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierAggregator *VerifierAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.contract.Transfer(opts)
}

func (_VerifierAggregator *VerifierAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierAggregator *VerifierAggregatorCaller) GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []VerifierAggregatorSourceChainConfig, error) {
	var out []interface{}
	err := _VerifierAggregator.contract.Call(opts, &out, "getAllSourceChainConfigs")

	if err != nil {
		return *new([]uint64), *new([]VerifierAggregatorSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]VerifierAggregatorSourceChainConfig)).(*[]VerifierAggregatorSourceChainConfig)

	return out0, out1, err

}

func (_VerifierAggregator *VerifierAggregatorSession) GetAllSourceChainConfigs() ([]uint64, []VerifierAggregatorSourceChainConfig, error) {
	return _VerifierAggregator.Contract.GetAllSourceChainConfigs(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCallerSession) GetAllSourceChainConfigs() ([]uint64, []VerifierAggregatorSourceChainConfig, error) {
	return _VerifierAggregator.Contract.GetAllSourceChainConfigs(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCaller) GetDynamicConfig(opts *bind.CallOpts) (VerifierAggregatorDynamicConfig, error) {
	var out []interface{}
	err := _VerifierAggregator.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(VerifierAggregatorDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierAggregatorDynamicConfig)).(*VerifierAggregatorDynamicConfig)

	return out0, err

}

func (_VerifierAggregator *VerifierAggregatorSession) GetDynamicConfig() (VerifierAggregatorDynamicConfig, error) {
	return _VerifierAggregator.Contract.GetDynamicConfig(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCallerSession) GetDynamicConfig() (VerifierAggregatorDynamicConfig, error) {
	return _VerifierAggregator.Contract.GetDynamicConfig(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _VerifierAggregator.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_VerifierAggregator *VerifierAggregatorSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _VerifierAggregator.Contract.GetExecutionState(&_VerifierAggregator.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_VerifierAggregator *VerifierAggregatorCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _VerifierAggregator.Contract.GetExecutionState(&_VerifierAggregator.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_VerifierAggregator *VerifierAggregatorCaller) GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (VerifierAggregatorSourceChainConfig, error) {
	var out []interface{}
	err := _VerifierAggregator.contract.Call(opts, &out, "getSourceChainConfig", sourceChainSelector)

	if err != nil {
		return *new(VerifierAggregatorSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierAggregatorSourceChainConfig)).(*VerifierAggregatorSourceChainConfig)

	return out0, err

}

func (_VerifierAggregator *VerifierAggregatorSession) GetSourceChainConfig(sourceChainSelector uint64) (VerifierAggregatorSourceChainConfig, error) {
	return _VerifierAggregator.Contract.GetSourceChainConfig(&_VerifierAggregator.CallOpts, sourceChainSelector)
}

func (_VerifierAggregator *VerifierAggregatorCallerSession) GetSourceChainConfig(sourceChainSelector uint64) (VerifierAggregatorSourceChainConfig, error) {
	return _VerifierAggregator.Contract.GetSourceChainConfig(&_VerifierAggregator.CallOpts, sourceChainSelector)
}

func (_VerifierAggregator *VerifierAggregatorCaller) GetStaticConfig(opts *bind.CallOpts) (VerifierAggregatorStaticConfig, error) {
	var out []interface{}
	err := _VerifierAggregator.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(VerifierAggregatorStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierAggregatorStaticConfig)).(*VerifierAggregatorStaticConfig)

	return out0, err

}

func (_VerifierAggregator *VerifierAggregatorSession) GetStaticConfig() (VerifierAggregatorStaticConfig, error) {
	return _VerifierAggregator.Contract.GetStaticConfig(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCallerSession) GetStaticConfig() (VerifierAggregatorStaticConfig, error) {
	return _VerifierAggregator.Contract.GetStaticConfig(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierAggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierAggregator *VerifierAggregatorSession) Owner() (common.Address, error) {
	return _VerifierAggregator.Contract.Owner(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCallerSession) Owner() (common.Address, error) {
	return _VerifierAggregator.Contract.Owner(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifierAggregator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VerifierAggregator *VerifierAggregatorSession) TypeAndVersion() (string, error) {
	return _VerifierAggregator.Contract.TypeAndVersion(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorCallerSession) TypeAndVersion() (string, error) {
	return _VerifierAggregator.Contract.TypeAndVersion(&_VerifierAggregator.CallOpts)
}

func (_VerifierAggregator *VerifierAggregatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierAggregator.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierAggregator *VerifierAggregatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierAggregator.Contract.AcceptOwnership(&_VerifierAggregator.TransactOpts)
}

func (_VerifierAggregator *VerifierAggregatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierAggregator.Contract.AcceptOwnership(&_VerifierAggregator.TransactOpts)
}

func (_VerifierAggregator *VerifierAggregatorTransactor) ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []VerifierAggregatorSourceChainConfigArgs) (*types.Transaction, error) {
	return _VerifierAggregator.contract.Transact(opts, "applySourceChainConfigUpdates", sourceChainConfigUpdates)
}

func (_VerifierAggregator *VerifierAggregatorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []VerifierAggregatorSourceChainConfigArgs) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.ApplySourceChainConfigUpdates(&_VerifierAggregator.TransactOpts, sourceChainConfigUpdates)
}

func (_VerifierAggregator *VerifierAggregatorTransactorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []VerifierAggregatorSourceChainConfigArgs) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.ApplySourceChainConfigUpdates(&_VerifierAggregator.TransactOpts, sourceChainConfigUpdates)
}

func (_VerifierAggregator *VerifierAggregatorTransactor) Execute(opts *bind.TransactOpts, report VerifierAggregatorAggregatedReport) (*types.Transaction, error) {
	return _VerifierAggregator.contract.Transact(opts, "execute", report)
}

func (_VerifierAggregator *VerifierAggregatorSession) Execute(report VerifierAggregatorAggregatedReport) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.Execute(&_VerifierAggregator.TransactOpts, report)
}

func (_VerifierAggregator *VerifierAggregatorTransactorSession) Execute(report VerifierAggregatorAggregatedReport) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.Execute(&_VerifierAggregator.TransactOpts, report)
}

func (_VerifierAggregator *VerifierAggregatorTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _VerifierAggregator.contract.Transact(opts, "executeSingleMessage", message)
}

func (_VerifierAggregator *VerifierAggregatorSession) ExecuteSingleMessage(message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.ExecuteSingleMessage(&_VerifierAggregator.TransactOpts, message)
}

func (_VerifierAggregator *VerifierAggregatorTransactorSession) ExecuteSingleMessage(message InternalAny2EVMMultiProofMessage) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.ExecuteSingleMessage(&_VerifierAggregator.TransactOpts, message)
}

func (_VerifierAggregator *VerifierAggregatorTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig VerifierAggregatorDynamicConfig) (*types.Transaction, error) {
	return _VerifierAggregator.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_VerifierAggregator *VerifierAggregatorSession) SetDynamicConfig(dynamicConfig VerifierAggregatorDynamicConfig) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.SetDynamicConfig(&_VerifierAggregator.TransactOpts, dynamicConfig)
}

func (_VerifierAggregator *VerifierAggregatorTransactorSession) SetDynamicConfig(dynamicConfig VerifierAggregatorDynamicConfig) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.SetDynamicConfig(&_VerifierAggregator.TransactOpts, dynamicConfig)
}

func (_VerifierAggregator *VerifierAggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierAggregator.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierAggregator *VerifierAggregatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.TransferOwnership(&_VerifierAggregator.TransactOpts, to)
}

func (_VerifierAggregator *VerifierAggregatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierAggregator.Contract.TransferOwnership(&_VerifierAggregator.TransactOpts, to)
}

type VerifierAggregatorDynamicConfigSetIterator struct {
	Event *VerifierAggregatorDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierAggregatorDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierAggregatorDynamicConfigSet)
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
		it.Event = new(VerifierAggregatorDynamicConfigSet)
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

func (it *VerifierAggregatorDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierAggregatorDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierAggregatorDynamicConfigSet struct {
	DynamicConfig VerifierAggregatorDynamicConfig
	Raw           types.Log
}

func (_VerifierAggregator *VerifierAggregatorFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*VerifierAggregatorDynamicConfigSetIterator, error) {

	logs, sub, err := _VerifierAggregator.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorDynamicConfigSetIterator{contract: _VerifierAggregator.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierAggregator *VerifierAggregatorFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _VerifierAggregator.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierAggregatorDynamicConfigSet)
				if err := _VerifierAggregator.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_VerifierAggregator *VerifierAggregatorFilterer) ParseDynamicConfigSet(log types.Log) (*VerifierAggregatorDynamicConfigSet, error) {
	event := new(VerifierAggregatorDynamicConfigSet)
	if err := _VerifierAggregator.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierAggregatorExecutionStateChangedIterator struct {
	Event *VerifierAggregatorExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierAggregatorExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierAggregatorExecutionStateChanged)
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
		it.Event = new(VerifierAggregatorExecutionStateChanged)
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

func (it *VerifierAggregatorExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *VerifierAggregatorExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierAggregatorExecutionStateChanged struct {
	SourceChainSelector uint64
	SequenceNumber      uint64
	MessageId           [32]byte
	State               uint8
	ReturnData          []byte
	Raw                 types.Log
}

func (_VerifierAggregator *VerifierAggregatorFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*VerifierAggregatorExecutionStateChangedIterator, error) {

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

	logs, sub, err := _VerifierAggregator.contract.FilterLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorExecutionStateChangedIterator{contract: _VerifierAggregator.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_VerifierAggregator *VerifierAggregatorFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _VerifierAggregator.contract.WatchLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierAggregatorExecutionStateChanged)
				if err := _VerifierAggregator.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_VerifierAggregator *VerifierAggregatorFilterer) ParseExecutionStateChanged(log types.Log) (*VerifierAggregatorExecutionStateChanged, error) {
	event := new(VerifierAggregatorExecutionStateChanged)
	if err := _VerifierAggregator.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierAggregatorOwnershipTransferRequestedIterator struct {
	Event *VerifierAggregatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierAggregatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierAggregatorOwnershipTransferRequested)
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
		it.Event = new(VerifierAggregatorOwnershipTransferRequested)
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

func (it *VerifierAggregatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierAggregatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierAggregatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierAggregator *VerifierAggregatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierAggregatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierAggregator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorOwnershipTransferRequestedIterator{contract: _VerifierAggregator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierAggregator *VerifierAggregatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierAggregator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierAggregatorOwnershipTransferRequested)
				if err := _VerifierAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VerifierAggregator *VerifierAggregatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierAggregatorOwnershipTransferRequested, error) {
	event := new(VerifierAggregatorOwnershipTransferRequested)
	if err := _VerifierAggregator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierAggregatorOwnershipTransferredIterator struct {
	Event *VerifierAggregatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierAggregatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierAggregatorOwnershipTransferred)
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
		it.Event = new(VerifierAggregatorOwnershipTransferred)
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

func (it *VerifierAggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierAggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierAggregatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierAggregator *VerifierAggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierAggregatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierAggregator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorOwnershipTransferredIterator{contract: _VerifierAggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierAggregator *VerifierAggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierAggregator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierAggregatorOwnershipTransferred)
				if err := _VerifierAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VerifierAggregator *VerifierAggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierAggregatorOwnershipTransferred, error) {
	event := new(VerifierAggregatorOwnershipTransferred)
	if err := _VerifierAggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierAggregatorSourceChainConfigSetIterator struct {
	Event *VerifierAggregatorSourceChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierAggregatorSourceChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierAggregatorSourceChainConfigSet)
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
		it.Event = new(VerifierAggregatorSourceChainConfigSet)
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

func (it *VerifierAggregatorSourceChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierAggregatorSourceChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierAggregatorSourceChainConfigSet struct {
	SourceChainSelector uint64
	SourceConfig        VerifierAggregatorSourceChainConfig
	Raw                 types.Log
}

func (_VerifierAggregator *VerifierAggregatorFilterer) FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*VerifierAggregatorSourceChainConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _VerifierAggregator.contract.FilterLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorSourceChainConfigSetIterator{contract: _VerifierAggregator.contract, event: "SourceChainConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierAggregator *VerifierAggregatorFilterer) WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _VerifierAggregator.contract.WatchLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierAggregatorSourceChainConfigSet)
				if err := _VerifierAggregator.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
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

func (_VerifierAggregator *VerifierAggregatorFilterer) ParseSourceChainConfigSet(log types.Log) (*VerifierAggregatorSourceChainConfigSet, error) {
	event := new(VerifierAggregatorSourceChainConfigSet)
	if err := _VerifierAggregator.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierAggregatorStaticConfigSetIterator struct {
	Event *VerifierAggregatorStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierAggregatorStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierAggregatorStaticConfigSet)
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
		it.Event = new(VerifierAggregatorStaticConfigSet)
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

func (it *VerifierAggregatorStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierAggregatorStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierAggregatorStaticConfigSet struct {
	StaticConfig VerifierAggregatorStaticConfig
	Raw          types.Log
}

func (_VerifierAggregator *VerifierAggregatorFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*VerifierAggregatorStaticConfigSetIterator, error) {

	logs, sub, err := _VerifierAggregator.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &VerifierAggregatorStaticConfigSetIterator{contract: _VerifierAggregator.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierAggregator *VerifierAggregatorFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _VerifierAggregator.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierAggregatorStaticConfigSet)
				if err := _VerifierAggregator.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_VerifierAggregator *VerifierAggregatorFilterer) ParseStaticConfigSet(log types.Log) (*VerifierAggregatorStaticConfigSet, error) {
	event := new(VerifierAggregatorStaticConfigSet)
	if err := _VerifierAggregator.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VerifierAggregator *VerifierAggregator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierAggregator.abi.Events["DynamicConfigSet"].ID:
		return _VerifierAggregator.ParseDynamicConfigSet(log)
	case _VerifierAggregator.abi.Events["ExecutionStateChanged"].ID:
		return _VerifierAggregator.ParseExecutionStateChanged(log)
	case _VerifierAggregator.abi.Events["OwnershipTransferRequested"].ID:
		return _VerifierAggregator.ParseOwnershipTransferRequested(log)
	case _VerifierAggregator.abi.Events["OwnershipTransferred"].ID:
		return _VerifierAggregator.ParseOwnershipTransferred(log)
	case _VerifierAggregator.abi.Events["SourceChainConfigSet"].ID:
		return _VerifierAggregator.ParseSourceChainConfigSet(log)
	case _VerifierAggregator.abi.Events["StaticConfigSet"].ID:
		return _VerifierAggregator.ParseStaticConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierAggregatorDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b36")
}

func (VerifierAggregatorExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2")
}

func (VerifierAggregatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierAggregatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VerifierAggregatorSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x734ce292835471880fd796297a6bef2e23db07c84d09cad27f67cd7163f44035")
}

func (VerifierAggregatorStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e4950")
}

func (_VerifierAggregator *VerifierAggregator) Address() common.Address {
	return _VerifierAggregator.address
}

type VerifierAggregatorInterface interface {
	GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []VerifierAggregatorSourceChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (VerifierAggregatorDynamicConfig, error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (VerifierAggregatorSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (VerifierAggregatorStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []VerifierAggregatorSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report VerifierAggregatorAggregatedReport) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMultiProofMessage) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig VerifierAggregatorDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*VerifierAggregatorDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*VerifierAggregatorDynamicConfigSet, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*VerifierAggregatorExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*VerifierAggregatorExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierAggregatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierAggregatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierAggregatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierAggregatorOwnershipTransferred, error)

	FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*VerifierAggregatorSourceChainConfigSetIterator, error)

	WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSourceChainConfigSet(log types.Log) (*VerifierAggregatorSourceChainConfigSet, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*VerifierAggregatorStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierAggregatorStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*VerifierAggregatorStaticConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
