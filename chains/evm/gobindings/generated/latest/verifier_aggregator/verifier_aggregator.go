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
	RequiredVerifiers [][]byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.AggregatedReport\",\"components\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"proofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"requiredVerifiers\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461052057613f03803803809161001d8285610556565b833981019080820360c08112610520576080811261052057604051916100428361053b565b61004b81610579565b8352602081015161ffff8116810361052057602084019081526040820151906001600160a01b038216820361052057604085019182526060830151936001600160a01b03851685036105205760209060608701958652607f1901126105205760405192602084016001600160401b03811185821017610525576040526100d36080820161058d565b845260a0810151906001600160401b038211610520570186601f82011215610520578051966001600160401b038811610525578760051b916040519861011c602085018b610556565b89526020808a0193820101908282116105205760208101935b828510610420575050505050331561040f57600180546001600160a01b0319163317905581516001600160a01b03161580156103fd575b6103ec5784516001600160401b0316156103db5784516001600160401b03908116608090815283516001600160a01b0390811660a0528651811660c052835161ffff90811660e05260408051995190941689529351909316602088810191909152935183169187019190915293511660608501527f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b369390927f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a151151560ff196002541660ff821617600255604051908152a16000905b80518210156103855760009160208160051b8301015160018060401b036020820151169081156103765780516001600160a01b031615610367578185526005602052604085209460608201518051801591821561033b575b505061032c5750604081810151865492516001600160a81b031990931690151560a01b60ff60a01b16176001600160a01b039290921691909117855592936001937f734ce292835471880fd796297a6bef2e23db07c84d09cad27f67cd7163f44035916103088461059a565b5060ff82519154878060a01b038116835260a01c1615156020820152a20190610244565b6342bcdf7f60e11b8152600490fd5b6020919250012060405160208101908382526020815261035c604082610556565b51902014388061029c565b6342bcdf7f60e11b8552600485fd5b63c656089560e01b8552600485fd5b6040516138d5908161062e823960805181818161047d015261108c015260a0518181816104e00152610ff0015260c05181818161051c0152612178015260e0518181816104a401528181611df701526131450152f35b63c656089560e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b5083516001600160a01b03161561016c565b639b15e16f60e01b60005260046000fd5b84516001600160401b0381116105205782016080818603601f190112610520576040519061044d8261053b565b60208101516001600160a01b038116810361052057825261047060408201610579565b60208301526104816060820161058d565b604083015260808101516001600160401b03811161052057602091010185601f820112156105205780516001600160401b03811161052557604051916104d1601f8301601f191660200184610556565b81835287602083830101116105205760005b82811061050b5750509181600060208096949581960101526060820152815201940193610135565b806020809284010151828287010152016104e3565b600080fd5b634e487b7160e01b600052604160045260246000fd5b608081019081106001600160401b0382111761052557604052565b601f909101601f19168101906001600160401b0382119082101761052557604052565b51906001600160401b038216820361052057565b5190811515820361052057565b806000526004602052604060002054156000146106275760035468010000000000000000811015610525576001810180600355811015610611577fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0181905560035460009182526004602052604090912055600190565b634e487b7160e01b600052603260045260246000fd5b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806304666f9c146100e757806306285c69146100e2578063181f5a77146100dd5780634708209d146100d85780635215505b146100d35780635e36480c146100ce5780637437ff9f146100c957806379ba5097146100c4578063855cbf2a146100bf5780638da5cb5b146100ba578063a078d392146100b5578063e9d68a8e146100b05763f2fde38b146100ab57600080fd5b61155f565b611497565b610edb565b610e89565b610e32565b610a5f565b6109f6565b610995565b6107fc565b610694565b610617565b61042b565b6102cc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761013757604052565b6100ec565b6020810190811067ffffffffffffffff82111761013757604052565b6040810190811067ffffffffffffffff82111761013757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761013757604052565b604051906101c460e083610174565b565b604051906101c460a083610174565b604051906101c461010083610174565b604051906101c4604083610174565b67ffffffffffffffff81116101375760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361022a57565b600080fd5b67ffffffffffffffff81160361022a57565b8015150361022a57565b67ffffffffffffffff811161013757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561022a5780359061029c8261024b565b926102aa6040519485610174565b8284526020838301011161022a57816000926020809301838601378301015290565b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760043567ffffffffffffffff811161022a573660238201121561022a57806004013590610327826101f4565b906103356040519283610174565b8282526024602083019360051b8201019036821161022a5760248101935b8285106103655761036384611653565b005b843567ffffffffffffffff811161022a57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc823603011261022a57604051916103b18361011b565b60248201356103bf8161020c565b835260448201356103cf8161022f565b602084015260648201356103e281610241565b604084015260848201359267ffffffffffffffff841161022a57610410602094936024869536920101610285565b6060820152815201940193610353565b600091031261022a57565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a576104626118a6565b506105986040516104728161011b565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b604051906105ab602083610174565b60008252565b60005b8381106105c45750506000910152565b81810151838201526020016105b4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610610815180928187528780880191016105b1565b0116010190565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5761059860408051906106588183610174565b601c82527f566572696669657241676772656761746f7220312e372e302d646576000000006020830152519182916020835260208301906105d4565b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a577f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b3661074d6040516106f38161013c565b6004356106ff81610241565b815261070961209d565b51600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151560ff81169290921790556040519081529081906020820190565b0390a1005b6040810160408252825180915260206060830193019060005b8181106107dc5750505060208183039101526020808351928381520192019060005b81811061079a5750505090565b90919260206040826107d160019488516020809173ffffffffffffffffffffffffffffffffffffffff815116845201511515910152565b01940192910161078d565b825167ffffffffffffffff1685526020948501949092019160010161076b565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a57600354610837816101f4565b906108456040519283610174565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610872826101f4565b0160005b818110610938575050610888816118e4565b9060005b8181106108a457505061059860405192839283610752565b806108dc6108c36108b66001946132e3565b67ffffffffffffffff1690565b6108cd8387611962565b9067ffffffffffffffff169052565b61091c6109176108fd6108ef8488611962565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b61197b565b6109268287611962565b526109318186611962565b500161088c565b6020906109436118cb565b82828701015201610876565b6004111561095957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9060048210156109595752565b3461022a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760206109e76004356109d58161022f565b602435906109e28261022f565b6119f3565b6109f46040518092610988565bf35b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a576000604051610a338161013c565b52610598604051610a438161013c565b60025460ff161515908190526040519081529081906020820190565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760005473ffffffffffffffffffffffffffffffffffffffff81163303610b1e577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b919082608091031261022a57604051610b608161011b565b6060808294803584526020810135610b778161022f565b60208501526040810135610b8a8161022f565b6040850152013591610b9b8361022f565b0152565b35906101c48261020c565b63ffffffff81160361022a57565b35906101c482610baa565b81601f8201121561022a57803590610bda826101f4565b92610be86040519485610174565b82845260208085019360051b8301019181831161022a5760208101935b838510610c1457505050505090565b843567ffffffffffffffff811161022a57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603011261022a5760405191610c608361011b565b602082013567ffffffffffffffff811161022a57856020610c8392850101610285565b83526040820135610c938161020c565b602084015260608201359267ffffffffffffffff841161022a57608083610cc1886020809881980101610285565b604084015201356060820152815201940193610c05565b9080601f8301121561022a578135610cef816101f4565b92610cfd6040519485610174565b81845260208085019260051b8201019183831161022a5760208201905b838210610d2957505050505090565b813567ffffffffffffffff811161022a57602091610d4c87848094880101610285565b815201910190610d1a565b9190916101408184031261022a57610d6d6101b5565b92610d788183610b48565b8452608082013567ffffffffffffffff811161022a5781610d9a918401610285565b602085015260a082013567ffffffffffffffff811161022a5781610dbf918401610285565b6040850152610dd060c08301610b9f565b6060850152610de160e08301610bb8565b608085015261010082013567ffffffffffffffff811161022a5781610e07918401610bc3565b60a085015261012082013567ffffffffffffffff811161022a57610e2b9201610cd8565b60c0830152565b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760043567ffffffffffffffff811161022a57610e84610363913690600401610d57565b611cc3565b3461022a5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5760043567ffffffffffffffff811161022a578060040160607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261022a5760025460ff1661146d57610f8660017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006002541617600255565b610f9b6020610f958380611f8a565b01611fbd565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa90811561146857600091611439575b506114015761106d61106961105f8467ffffffffffffffff166000526005602052604060002090565b5460a01c60ff1690565b1590565b6113c9576110806040610f958380611f8a565b67ffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611381576110cb6110c56060610f958480611f8a565b836119f3565b6110ee816110d98480611f8a565b6110e66024880186611fdc565b929091612ae7565b6111086111036110fe8480611f8a565b612030565b612cea565b906111128161094f565b8015908115918261136e575b156113175760446111a296019363ffffffff6111398661203b565b166113025761117961115960e0611153846111ab95611f8a565b0161203b565b8551606001516111739067ffffffffffffffff1689612d3b565b856130fe565b97909561119d87611197606089510167ffffffffffffffff90511690565b8a612dc6565b61203b565b63ffffffff1690565b6112a0575b50506111bb8261094f565b6002820361124c575b67ffffffffffffffff7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2915161121c611208606083015167ffffffffffffffff1690565b91519683604051948594169716958361207f565b0390a46103637fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060025416600255565b6112558261094f565b60038203156111c457516060015161129c929067ffffffffffffffff16907f926c5a3e0000000000000000000000000000000000000000000000000000000060005261205c565b6000fd5b6112a98461094f565b600384036111b0576112ba9061094f565b6112c55738806111b0565b839051516112fe6040519283927f2b11b8d900000000000000000000000000000000000000000000000000000000845260048401612045565b0390fd5b506111ab6111796113128661203b565b611159565b61129c85611332606086510167ffffffffffffffff90511690565b7f3b5754190000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff91821660045216602452604490565b506113788161094f565b6003811461111e565b6113936040610f958361129c94611f8a565b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b61145b915060203d602011611461575b6114538183610174565b810190611fc7565b38611036565b503d611449565b611ca3565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5767ffffffffffffffff6004356114db8161022f565b6114e36118cb565b50166000526005602052610598604060002060ff6040519161150483610158565b5473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015260405191829182815173ffffffffffffffffffffffffffffffffffffffff16815260209182015115159181019190915260400190565b3461022a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022a5773ffffffffffffffffffffffffffffffffffffffff6004356115af8161020c565b6115b761209d565b1633811461162957807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b61165b61209d565b60005b81518110156118a2576116718183611962565b5190611688602083015167ffffffffffffffff1690565b67ffffffffffffffff8116908115611878576116d76116be6116be865173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b15611838576116fa9067ffffffffffffffff166000526005602052604060002090565b606084015180518015918215611862575b50506118385761182f7f734ce292835471880fd796297a6bef2e23db07c84d09cad27f67cd7163f44035916117f06117af8761179561174f604060019b0151151590565b85547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178555565b5173ffffffffffffffffffffffffffffffffffffffff1690565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6117f98461377f565b5060408051915473ffffffffffffffffffffffffffffffffffffffff8116835260a01c60ff161515602083015290918291820190565b0390a20161165e565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6020012090506118706120e8565b14388061170b565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b604051906118b38261011b565b60006060838281528260208201528260408201520152565b604051906118d882610158565b60006020838281520152565b906118ee826101f4565b6118fb6040519182610174565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061192982946101f4565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156119765760209160051b010190565b611933565b9060405161198881610158565b915473ffffffffffffffffffffffffffffffffffffffff8116835260a01c60ff1615156020830152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b919082039182116119ee57565b6119b2565b6119ff82607f92612109565b9116906801fffffffffffffffe67ffffffffffffffff83169260011b1691808304600214901517156119ee576003911c1660048110156109595790565b604051906105ab8261013c565b60405190611a58602083610174565b600080835282815b828110611a6c57505050565b602090611a776118cb565b82828501015201611a60565b90611a8d826101f4565b611a9a6040519182610174565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611ac882946101f4565b019060005b828110611ad957505050565b602090611ae46118cb565b82828501015201611acd565b909160608284031261022a578151611b0781610241565b92602083015167ffffffffffffffff811161022a5783019080601f8301121561022a57815191611b368361024b565b91611b446040519384610174565b8383526020848301011161022a57604092611b6591602080850191016105b1565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a08401526080611be2611bae604084015160a060c08801526101208701906105d4565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526105d4565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b818110611c6b5750505061ffff90951660208301526101c49291606091611c4f9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101611c21565b6040513d6000823e3d90fd5b906020611cc09281815201906105d4565b90565b303303611f6057611cd2611a49565b9060a08101518051611f0f575b50805191611cfa6020845194015167ffffffffffffffff1690565b906020830151916040840192611d28845192611d146101c6565b97885267ffffffffffffffff166020880152565b60408601526060850152608084015251511580611ef1575b8015611ece575b8015611e9c575b6118a257611e209181611d936116be611d796108fd6020600097510167ffffffffffffffff90511690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9083611dc76060611dab608085015163ffffffff1690565b93015173ffffffffffffffffffffffffffffffffffffffff1690565b93604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601611b6b565b03925af190811561146857600090600092611e75575b5015611e3f5750565b6112fe906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611caf565b9050611e9491503d806000833e611e8c8183610174565b810190611af0565b509038611e36565b50611ec9611069611ec4606084015173ffffffffffffffffffffffffffffffffffffffff1690565b61253f565b611d4e565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b15611d47565b5063ffffffff611f08608083015163ffffffff1690565b1615611d40565b8192506020611f59920151611f3b606085015173ffffffffffffffffffffffffffffffffffffffff1690565b90611f53602086510167ffffffffffffffff90511690565b9261214f565b9038611cdf565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec18136030182121561022a570190565b35611cc08161022f565b9081602091031261022a5751611cc081610241565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561022a570180359067ffffffffffffffff821161022a57602001918160051b3603831361022a57565b611cc0903690610d57565b35611cc081610baa565b604090611cc09392815281602082015201906105d4565b929160449067ffffffffffffffff6101c493816064971660045216602452610988565b80612090604092611cc09594610988565b81602082015201906105d4565b73ffffffffffffffffffffffffffffffffffffffff6001541633036120be57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101906000825260208152612103604082610174565b51902090565b9067ffffffffffffffff61214b921660005260076020526701ffffffffffffff60406000209160071c1667ffffffffffffffff16600052602052604060002090565b5490565b9390919361215d8151611a83565b9260009573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016925b8051881015612534576121ad8882611962565b51976121b76118cb565b506122306121dc60208b015173ffffffffffffffffffffffffffffffffffffffff1690565b9960208b604051809481927fbbe4f6db0000000000000000000000000000000000000000000000000000000083526004830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b03818a5afa91821561146857600092612504575b5073ffffffffffffffffffffffffffffffffffffffff82169182156124c15761226c81612595565b6124c15761227c611069826125bf565b6124c157506123546020878c8e9461230a612297878c613812565b966122a0611a3c565b5060608301516122cc60408551950151956122b96101d5565b97885267ffffffffffffffff1688880152565b73ffffffffffffffffffffffffffffffffffffffff8d166040870152606086015273ffffffffffffffffffffffffffffffffffffffff166080850152565b60a083015260c082015261231c61059c565b60e0820152604051809381927f390775370000000000000000000000000000000000000000000000000000000083526004830161334b565b03816000875af160009181612491575b506123a8578b612372613474565b906112fe6040519283927f9fe2f95a000000000000000000000000000000000000000000000000000000008452600484016134a4565b9a92939495969798999a9173ffffffffffffffffffffffffffffffffffffffff871603612426575b509060019291516123fe6123e26101e5565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015261240d828a611962565b526124188189611962565b50019695949392919061219a565b6124308387613812565b90808210801561247d575b61244557506123d0565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b5061248881836119e1565b8351141561243b565b6124b391925060203d81116124ba575b6124ab8183610174565b81019061332d565b9038612364565b503d6124a1565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b61252691925060203d811161252d575b61251e8183610174565b810190613318565b9038612244565b503d612514565b505050509250905090565b6125697f85572ffb00000000000000000000000000000000000000000000000000000000826135f7565b9081612583575b81612579575090565b611cc09150613597565b905061258e816134d1565b1590612570565b6125697ff208a58f00000000000000000000000000000000000000000000000000000000826135f7565b6125697faff2afbf00000000000000000000000000000000000000000000000000000000826135f7565b91908110156119765760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561022a57019081359167ffffffffffffffff831161022a57602001823603811361022a579190565b919091357fffffffff000000000000000000000000000000000000000000000000000000008116926004811061267d575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561022a57016020813591019167ffffffffffffffff821161022a57813603831361022a57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561022a57016020813591019167ffffffffffffffff821161022a578160051b3603831361022a57565b90602083828152019160208260051b8501019381936000915b8483106127ba5750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818636030182121561022a576020809187600194019060608061288361284161283386806126af565b6080875260808701916126ff565b73ffffffffffffffffffffffffffffffffffffffff878701356128638161020c565b168786015261287560408701876126af565b9086830360408801526126ff565b93013591015298019301930191949392906127aa565b90602083828152019260208260051b82010193836000925b8484106128c15750505050505090565b909192939495602080612907837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030188526129018b886126af565b906126ff565b98019401940192949391906128b1565b90611cc091602081528135602082015267ffffffffffffffff602083013561293e8161022f565b16604082015267ffffffffffffffff604083013561295b8161022f565b16606082015267ffffffffffffffff60608301356129788161022f565b166080820152612a7b612a6f6129e76129aa61299760808701876126af565b61014060a08801526101608701916126ff565b6129b760a08701876126af565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c08801526126ff565b612a136129f660c08701610b9f565b73ffffffffffffffffffffffffffffffffffffffff1660e0860152565b612a30612a2260e08701610bb8565b63ffffffff16610100860152565b612a3e61010086018661273e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086840301610120870152612791565b9261012081019061273e565b916101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301910152612899565b906060926101c495979694612acc612ada936080865260808601906105d4565b9184830360208601526126ff565b9560408201520190610988565b91929092610120830182612afb8286611fdc565b905003612cac5760005b612b0f8286611fdc565b9050811015612ca457612b35612b2f82612b298589611fdc565b906125e9565b90612649565b73ffffffffffffffffffffffffffffffffffffffff612b83611d79837fffffffff00000000000000000000000000000000000000000000000000000000166000526006602052604060002090565b16908115612c555750604051612bcc81612ba08960208301612917565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610174565b612bd78387876125e9565b9092803b1561022a57849360008094612c208d604051998a97889687957fcba4c71a00000000000000000000000000000000000000000000000000000000875260048701612aac565b03925af191821561146857600192612c3a575b5001612b05565b80612c496000612c4f93610174565b80610420565b38612c33565b7f29391c08000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b505050505050565b612cba915061129c93611fdc565b7f301b70fa0000000000000000000000000000000000000000000000000000000060005260045250602452604490565b60405160e0810181811067ffffffffffffffff8211176101375760609160c091604052612d156118a6565b8152826020820152826040820152600083820152600060808201528260a0820152015290565b607f8216906801fffffffffffffffe67ffffffffffffffff83169260011b1691808304600214901517156119ee57612dc39167ffffffffffffffff612d808584612109565b921660005260076020526701ffffffffffffff60406000209460071c169160036001831b921b191617929067ffffffffffffffff16600052602052604060002090565b55565b9091607f8316916801fffffffffffffffe67ffffffffffffffff84169360011b1692808404600214901517156119ee57612e008482612109565b92600483101561095957612dc39367ffffffffffffffff612e45931660005260076020526003604060002094831b921b191617936701ffffffffffffff9060071c1690565b67ffffffffffffffff16600052602052604060002090565b9080602083519182815201916020808360051b8301019401926000915b838310612e8957505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08560019503018652885190606080612f07612ed785516080865260808601906105d4565b73ffffffffffffffffffffffffffffffffffffffff878701511687860152604086015185820360408701526105d4565b93015191015297019301930191939290612e7a565b9080602083519182815201916020808360051b8301019401926000915b838310612f4857505050505090565b9091929394602080612f84837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516105d4565b97019301930191939290612f39565b90611cc09160208152612fd660208201835167ffffffffffffffff6060809280518552826020820151166020860152826040820151166040860152015116910152565b60c061309261302c612ff9602086015161014060a08701526101608601906105d4565b60408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086830301858701526105d4565b606085015173ffffffffffffffffffffffffffffffffffffffff1660e0850152608085015163ffffffff1661010085015260a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301610120860152612e5d565b920151906101407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152612f1c565b90602082519201517fffffffff000000000000000000000000000000000000000000000000000000008116926004811061267d575050565b90612ba061313e61316e936040519283917f855cbf2a00000000000000000000000000000000000000000000000000000000602084015260248301612f93565b63ffffffff7f00000000000000000000000000000000000000000000000000000000000000009216903090613659565b5090156131815750600290611cc061059c565b9072c11c11c11c11c11c11c11c11c11c11c11c11c133146131a3575b60039190565b6131d46131af836130c6565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f37c3be2900000000000000000000000000000000000000000000000000000000148015613297575b8015613263575b1561319d5761129c613215836130c6565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b506132706131af836130c6565b7fea7f4b120000000000000000000000000000000000000000000000000000000014613204565b506132a46131af836130c6565b7fafa32a2c00000000000000000000000000000000000000000000000000000000146131fd565b80548210156119765760005260206000200190600090565b6003548110156119765760036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b015490565b9081602091031261022a5751611cc08161020c565b9081602091031261022a57604051906133458261013c565b51815290565b90611cc0916020815260e061344061340d613374855161010060208701526101208601906105d4565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff166060860152606086015160808601526133d9608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526105d4565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526105d4565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526105d4565b3d1561349f573d906134858261024b565b916134936040519384610174565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff611cc0949316815281602082015201906105d4565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252613531604483610174565b6179185a1061356d576020926000925191617530fa6000513d82613561575b508161355a575090565b9050151590565b60201115915038613550565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252613531604483610174565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252613531604483610174565b939193613666608461024b565b946136746040519687610174565b60848652613682608461024b565b947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602088019601368737833b15613755575a9080821061372b578291038060061c90031115613701576000918291825a9560208451940192f1905a9003923d90608482116136f8575b6000908287523e929190565b608491506136ec565b7f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b8060005260046020526040600020541560001461380c5760035468010000000000000000811015610137578060016137bc920160035560036132cb565b81549060031b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84831b921b191617905561380660035491600490600052602052604060002090565b55600190565b50600090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181613891575b5061388d5782612372613474565b9150565b90916020823d6020116138c0575b816138ac60209383610174565b810103126138bd575051903861387f565b80fd5b3d915061389f56fea164736f6c634300081a000a",
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
