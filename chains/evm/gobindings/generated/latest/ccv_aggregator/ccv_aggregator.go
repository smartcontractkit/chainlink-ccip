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
	Message  InternalAny2EVMMultiProofMessage
	Ccvs     []common.Address
	CcvBlobs [][]byte
	Proofs   [][]byte
}

type CCVAggregatorDynamicConfig struct {
	ReentrancyGuardEntered bool
}

type CCVAggregatorSourceChainConfig struct {
	Router      common.Address
	IsEnabled   bool
	DefaultCCV  common.Address
	RequiredCCV common.Address
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
	Header       InternalHeader
	Sender       []byte
	Data         []byte
	Receiver     common.Address
	GasLimit     uint32
	TokenAmounts []InternalAny2EVMMultiProofTokenTransfer
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.AggregatedReport\",\"components\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvBlobs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"proofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoCCVQuorumReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461053e576140d8803803809161001d8285610574565b833981019080820360c0811261053e576080811261053e576040519161004283610559565b61004b81610597565b8352602081015161ffff8116810361053e57602084019081526040820151906001600160a01b038216820361053e57604085019182526060830151936001600160a01b038516850361053e5760209060608701958652607f19011261053e5760405192602084016001600160401b03811185821017610543576040526100d3608082016105ab565b845260a0810151906001600160401b03821161053e570186601f8201121561053e578051966001600160401b038811610543578760051b916040519861011c602085018b610574565b89526020808a01938201019082821161053e5760208101935b82851061043e575050505050331561042d57600180546001600160a01b0319163317905581516001600160a01b031615801561041b575b61040a5784516001600160401b0316156103f95784516001600160401b03908116608090815283516001600160a01b0390811660a0528651811660c052835161ffff90811660e05260408051995190941689529351909316602088810191909152935183169187019190915293511660608501527f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b369390927f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a151151560ff196002541660ff821617600255604051908152a16000905b80518210156103a35760009160208160051b8301015160018060401b036020820151169081156103945780516001600160a01b0316156103855781855260056020526040852094606082015180518015918215610359575b505061034a57506040810151855491516001600160a81b031990921690151560a01b60ff60a01b16176001600160a01b03919091161784559192600192907f58a20cdf97a4562295fa419a74c9bdf2683d21773d052231dc4da284a495bfb090608090610308846105b8565b5060408051825460a089811b8a9003808316845291901c60ff161515602083015288840154811692820192909252600290920154166060820152a20190610244565b6342bcdf7f60e11b8152600490fd5b6020919250012060405160208101908382526020815261037a604082610574565b51902014388061029c565b6342bcdf7f60e11b8552600485fd5b63c656089560e01b8552600485fd5b604051613a8c908161064c82396080518181816104610152610f89015260a0518181816104c40152610eed015260c05181818161050001526127ea015260e05181818161048801528181611d2c01526131930152f35b63c656089560e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b5083516001600160a01b03161561016c565b639b15e16f60e01b60005260046000fd5b84516001600160401b03811161053e5782016080818603601f19011261053e576040519061046b82610559565b60208101516001600160a01b038116810361053e57825261048e60408201610597565b602083015261049f606082016105ab565b604083015260808101516001600160401b03811161053e57602091010185601f8201121561053e5780516001600160401b03811161054357604051916104ef601f8301601f191660200184610574565b818352876020838301011161053e5760005b8281106105295750509181600060208096949581960101526060820152815201940193610135565b80602080928401015182828701015201610501565b600080fd5b634e487b7160e01b600052604160045260246000fd5b608081019081106001600160401b0382111761054357604052565b601f909101601f19168101906001600160401b0382119082101761054357604052565b51906001600160401b038216820361053e57565b5190811515820361053e57565b80600052600460205260406000205415600014610645576003546801000000000000000081101561054357600181018060035581101561062f577fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0181905560035460009182526004602052604090912055600190565b634e487b7160e01b600052603260045260246000fd5b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806304666f9c146100e757806306285c69146100e2578063181f5a77146100dd57806345ec4b5f146100d85780634708209d146100d35780635215505b146100ce5780635e36480c146100c95780637437ff9f146100c457806379ba5097146100bf5780638cdf6d78146100ba5780638da5cb5b146100b5578063e9d68a8e146100b05763f2fde38b146100ab57600080fd5b6115b7565b611501565b611464565b610dd8565b610cef565b610c86565b610c25565b610a8c565b61090c565b6108b5565b6105fb565b61040f565b6102b0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761013757604052565b6100ec565b6020810190811067ffffffffffffffff82111761013757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761013757604052565b604051906101a860c083610158565b565b604051906101a860a083610158565b604051906101a861010083610158565b604051906101a8604083610158565b67ffffffffffffffff81116101375760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361020e57565b600080fd5b67ffffffffffffffff81160361020e57565b8015150361020e57565b67ffffffffffffffff811161013757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561020e578035906102808261022f565b9261028e6040519485610158565b8284526020838301011161020e57816000926020809301838601378301015290565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5760043567ffffffffffffffff811161020e573660238201121561020e5780600401359061030b826101d8565b906103196040519283610158565b8282526024602083019360051b8201019036821161020e5760248101935b82851061034957610347846116ab565b005b843567ffffffffffffffff811161020e57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc823603011261020e57604051916103958361011b565b60248201356103a3816101f0565b835260448201356103b381610213565b602084015260648201356103c681610225565b604084015260848201359267ffffffffffffffff841161020e576103f4602094936024869536920101610269565b6060820152815201940193610337565b600091031261020e57565b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e57610446611900565b5061057c6040516104568161011b565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b6040519061058f602083610158565b60008252565b60005b8381106105a85750506000910152565b8181015183820152602001610598565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936105f481518092818752878088019101610595565b0116010190565b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5761057c604080519061063c8183610158565b601782527f43435641676772656761746f7220312e372e302d6465760000000000000000006020830152519182916020835260208301906105b8565b919082608091031261020e576040516106908161011b565b60608082948035845260208101356106a781610213565b602085015260408101356106ba81610213565b60408501520135916106cb83610213565b0152565b35906101a8826101f0565b359063ffffffff8216820361020e57565b81601f8201121561020e57803590610702826101d8565b926107106040519485610158565b82845260208085019360051b8301019181831161020e5760208101935b83851061073c57505050505090565b843567ffffffffffffffff811161020e57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603011261020e57604051916107888361011b565b602082013567ffffffffffffffff811161020e578560206107ab92850101610269565b835260408201356107bb816101f0565b602084015260608201359267ffffffffffffffff841161020e576080836107e9886020809881980101610269565b60408401520135606082015281520194019361072d565b9190916101208184031261020e57610816610199565b926108218183610678565b8452608082013567ffffffffffffffff811161020e5781610843918401610269565b602085015260a082013567ffffffffffffffff811161020e5781610868918401610269565b604085015261087960c083016106cf565b606085015261088a60e083016106da565b608085015261010082013567ffffffffffffffff811161020e576108ae92016106eb565b60a0830152565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5760043567ffffffffffffffff811161020e57610907610347913690600401610800565b611bf4565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e577f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b366109c560405161096b8161013c565b60043561097781610225565b8152610981612537565b51600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151560ff81169290921790556040519081529081906020820190565b0390a1005b6040810160408252825180915260206060830193019060005b818110610a6c5750505060208183039101526020808351928381520192019060005b818110610a125750505090565b9091926020608082610a61600194885173ffffffffffffffffffffffffffffffffffffffff60608092828151168552602081015115156020860152826040820151166040860152015116910152565b019401929101610a05565b825167ffffffffffffffff168552602094850194909201916001016109e3565b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e57600354610ac7816101d8565b90610ad56040519283610158565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610b02826101d8565b0160005b818110610bc8575050610b1881611ed7565b9060005b818110610b3457505061057c604051928392836109ca565b80610b6c610b53610b4660019461356d565b67ffffffffffffffff1690565b610b5d8387611a0d565b9067ffffffffffffffff169052565b610bac610ba7610b8d610b7f8488611a0d565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b611f26565b610bb68287611a0d565b52610bc18186611a0d565b5001610b1c565b602090610bd3611900565b82828701015201610b06565b60041115610be957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b906004821015610be95752565b3461020e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e576020610c77600435610c6581610213565b60243590610c7282610213565b611fb7565b610c846040518092610c18565bf35b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e576000604051610cc38161013c565b5261057c604051610cd38161013c565b60025460ff161515908190526040519081529081906020820190565b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5760005473ffffffffffffffffffffffffffffffffffffffff81163303610dae577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5760043567ffffffffffffffff811161020e578060040160807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261020e5760025460ff1661143a57610e8360017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006002541617600255565b610e986020610e92838061200d565b01612040565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b1660048201529092906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa90811561117b5760009161140b575b506113d357610f6a610f66610f5c8567ffffffffffffffff166000526005602052604060002090565b5460a01c60ff1690565b1590565b61139b57610f7d6040610e92848061200d565b67ffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036113535781610fcd610fc76060610e9284611050989761200d565b84611fb7565b91610fdd6020610e92848061200d565b91611034610ff660c0610ff0848061200d565b0161205f565b61101c6020610ff061101661100b878061200d565b610100810190612069565b906120bd565b94602485019561102c8786612069565b939092612c4d565b61107c611041828061200d565b604051978891602083016122e0565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101885287610158565b6000926064604484019301935b6110938284612069565b9050811015611180576110d76110be6110be6110b9846110b38789612069565b90612433565b61205f565b73ffffffffffffffffffffffffffffffffffffffff1690565b906110ec816110e68787612069565b90612443565b6110fd836110e68a89979597612069565b929094813b1561020e5760008d6111468d83976040519a8b98899788967ffd810751000000000000000000000000000000000000000000000000000000008852600488016124a3565b03925af191821561117b57600192611160575b5001611089565b8061116f600061117593610158565b80610404565b38611159565b611bd4565b858761119c611197611192878061200d565b6124eb565b612d9b565b916111a681610bdf565b801590811561133f575b50156112e9578151606001516111d09067ffffffffffffffff1682612de6565b6111d982613125565b926111fb826111f5606084510167ffffffffffffffff90511690565b85612e71565b61120482610bdf565b60028203611295575b67ffffffffffffffff7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df29151611265611251606083015167ffffffffffffffff1690565b915196836040519485941697169583612519565b0390a46103477fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060025416600255565b61129e82610bdf565b600382031561120d5751606001516112e5929067ffffffffffffffff16907f926c5a3e000000000000000000000000000000000000000000000000000000006000526124f6565b6000fd5b61130360606112e593510167ffffffffffffffff90511690565b7f3b5754190000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff91821660045216602452604490565b6003915061134c81610bdf565b14836111b0565b6112e56113656040610e92858061200d565b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b61142d915060203d602011611433575b6114258183610158565b81019061204a565b38610f33565b503d61141b565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b6101a890929192608081019373ffffffffffffffffffffffffffffffffffffffff60608092828151168552602081015115156020860152826040820151166040860152015116910152565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5767ffffffffffffffff60043561154581610213565b61154d611900565b5016600052600560205261057c604060002073ffffffffffffffffffffffffffffffffffffffff6002604051926115838461011b565b60ff8154848116865260a01c16151560208501528260018201541660408501520154166060820152604051918291826114b6565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5773ffffffffffffffffffffffffffffffffffffffff600435611607816101f0565b61160f612537565b1633811461168157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b6116b3612537565b60005b81518110156118fc576116c98183611a0d565b51906116e0602083015167ffffffffffffffff1690565b67ffffffffffffffff81169081156118d2576117166110be6110be865173ffffffffffffffffffffffffffffffffffffffff1690565b15611892576117399067ffffffffffffffff166000526005602052604060002090565b6060840151805180159182156118bc575b5050611892576118897f58a20cdf97a4562295fa419a74c9bdf2683d21773d052231dc4da284a495bfb09161182f6117ee876117d461178e604060019b0151151590565b85547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178555565b5173ffffffffffffffffffffffffffffffffffffffff1690565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b611838846139ec565b5060405191829182919091606073ffffffffffffffffffffffffffffffffffffffff6002608084019560ff8154848116875260a01c1615156020860152826001820154166040860152015416910152565b0390a2016116b6565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6020012090506118ca612582565b14388061174a565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b6040519061190d8261011b565b60006060838281528260208201528260408201520152565b604051906040820182811067ffffffffffffffff8211176101375760405260006020838281520152565b90611959826101d8565b6119666040519182610158565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061199482946101d8565b019060005b8281106119a557505050565b6020906119b0611925565b82828501015201611999565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156119f85760200190565b6119bc565b8051600110156119f85760400190565b80518210156119f85760209160051b010190565b909160608284031261020e578151611a3881610225565b92602083015167ffffffffffffffff811161020e5783019080601f8301121561020e57815191611a678361022f565b91611a756040519384610158565b8383526020848301011161020e57604092611a969160208085019101610595565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a08401526080611b13611adf604084015160a060c08801526101208701906105b8565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526105b8565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b818110611b9c5750505061ffff90951660208301526101a89291606091611b809063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101611b52565b6040513d6000823e3d90fd5b906020611bf19281815201906105b8565b90565b303303611ead5760a0810190611c0b82515161194f565b91518051611e48575b50805191611c2f6020845194015167ffffffffffffffff1690565b906020830151916040840192611c5d845192611c496101aa565b97885267ffffffffffffffff166020880152565b60408601526060850152608084015251511580611e2a575b8015611e07575b8015611dd5575b6118fc57611d559181611cc86110be611cae610b8d6020600097510167ffffffffffffffff90511690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9083611cfc6060611ce0608085015163ffffffff1690565b93015173ffffffffffffffffffffffffffffffffffffffff1690565b93604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601611a9c565b03925af190811561117b57600090600092611dae575b5015611d745750565b611daa906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611be0565b0390fd5b9050611dcd91503d806000833e611dc58183610158565b810190611a21565b509038611d6b565b50611e02610f66611dfd606084015173ffffffffffffffffffffffffffffffffffffffff1690565b612adb565b611c83565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b15611c7c565b5063ffffffff611e41608083015163ffffffff1690565b1615611c75565b611e54611e93916119eb565b516020830151606084015173ffffffffffffffffffffffffffffffffffffffff1690611e8d602086510167ffffffffffffffff90511690565b9261275c565b611e9c836119eb565b52611ea6826119eb565b5038611c14565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b90611ee1826101d8565b611eee6040519182610158565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611f1c82946101d8565b0190602036910137565b90604051611f338161011b565b606073ffffffffffffffffffffffffffffffffffffffff6002839560ff8154848116875260a01c1615156020860152826001820154166040860152015416910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211611fb257565b611f76565b611fc382607f92612baf565b9116906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611fb2576003911c166004811015610be95790565b6040519061058f8261013c565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee18136030182121561020e570190565b35611bf181610213565b9081602091031261020e5751611bf181610225565b35611bf1816101f0565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561020e570180359067ffffffffffffffff821161020e57602001918160051b3603831361020e57565b90156119f8578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818136030182121561020e570190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561020e57016020813591019167ffffffffffffffff821161020e57813603831361020e57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561020e57016020813591019167ffffffffffffffff821161020e578160051b3603831361020e57565b90602083828152019160208260051b8501019381936000915b8483106122015750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818636030182121561020e57602080918760019401906060806122ca61228861227a86806120f6565b608087526080870191612146565b73ffffffffffffffffffffffffffffffffffffffff878701356122aa816101f0565b16878601526122bc60408701876120f6565b908683036040880152612146565b93013591015298019301930191949392906121f1565b90611bf191602081528135602082015267ffffffffffffffff602083013561230781610213565b16604082015267ffffffffffffffff604083013561232481610213565b16606082015267ffffffffffffffff606083013561234181610213565b1660808201526124026123ad61237061235d60808601866120f6565b61012060a0870152610140860191612146565b61237d60a08601866120f6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160c0870152612146565b926123da6123bd60c083016106cf565b73ffffffffffffffffffffffffffffffffffffffff1660e0850152565b6123f76123e960e083016106da565b63ffffffff16610100850152565b610100810190612185565b916101207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101526121d8565b91908110156119f85760051b0190565b91908110156119f85760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561020e57019081359167ffffffffffffffff831161020e57602001823603811361020e579190565b96959390946124d56124e3936060956124c76101a89960808d5260808d01906105b8565b918b830360208d0152612146565b9188830360408a0152612146565b940190610c18565b611bf1903690610800565b929160449067ffffffffffffffff6101a893816064971660045216602452610c18565b8061252a604092611bf19594610c18565b81602082015201906105b8565b73ffffffffffffffffffffffffffffffffffffffff60015416330361255857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040516020810190600082526020815261259d604082610158565b51902090565b9081602091031261020e5751611bf1816101f0565b9081602091031261020e57604051906125d08261013c565b51815290565b90611bf1916020815260e06126cb6126986125ff855161010060208701526101208601906105b8565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff16606086015260608601516080860152612664608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526105b8565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526105b8565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526105b8565b3d1561272a573d906127108261022f565b9161271e6040519384610158565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff611bf1949316815281602082015201906105b8565b909291612767611925565b50602082015173ffffffffffffffffffffffffffffffffffffffff166040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa94851561117b57600095612aaa575b5073ffffffffffffffffffffffffffffffffffffffff8516948515612a675761284381612b31565b612a6757612853610f6682612b5b565b612a675750612925916020916128698886613317565b95612872612000565b50606081015161289e604083519301519361288b6101b9565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c08201526128ed610580565b60e0820152604051809381927f39077537000000000000000000000000000000000000000000000000000000008352600483016125d6565b03816000885af160009181612a36575b5061297957846129436126ff565b90611daa6040519283927f9fe2f95a0000000000000000000000000000000000000000000000000000000084526004840161272f565b84909373ffffffffffffffffffffffffffffffffffffffff8316036129cc575b505050516129c46129a86101c9565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b6129d591613317565b908082108015612a22575b6129ea5783612999565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b50612a2d8183611fa5565b835114156129e0565b612a5991925060203d602011612a60575b612a518183610158565b8101906125b8565b9038612935565b503d612a47565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b612acd91955060203d602011612ad4575b612ac58183610158565b8101906125a3565b933861281b565b503d612abb565b612b057f85572ffb00000000000000000000000000000000000000000000000000000000826134f3565b9081612b1f575b81612b15575090565b611bf19150613493565b9050612b2a816133cd565b1590612b0c565b612b057ff208a58f00000000000000000000000000000000000000000000000000000000826134f3565b612b057faff2afbf00000000000000000000000000000000000000000000000000000000826134f3565b612b057f7909b54900000000000000000000000000000000000000000000000000000000826134f3565b9067ffffffffffffffff612bf1921660005260066020526701ffffffffffffff60406000209160071c1667ffffffffffffffff16600052602052604060002090565b5490565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611fb25760010190565b8015611fb2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b612c579250613701565b949060005b8351811015612cf5576000805b878110612ca9575b5015612c7f57600101612c5c565b7f8429b8bd0000000000000000000000000000000000000000000000000000000060005260046000fd5b612cb76110b9828a8a612433565b73ffffffffffffffffffffffffffffffffffffffff612cdc6110be6117d4878b611a0d565b911614612ceb57600101612c69565b5050600138612c71565b50915092919360ff16916000935b8251851015612d8f5760005b82811080612d86575b15612d7c57612d2b6110b982858a612433565b73ffffffffffffffffffffffffffffffffffffffff612d506110be6117d48a89611a0d565b911614612d6557612d6090612bf5565b612d0f565b509392612d73600191612c22565b935b0193612d03565b5093600190612d75565b50841515612d18565b945050509050612c7f57565b60405160c0810181811067ffffffffffffffff8211176101375760609160a091604052612dc6611900565b815282602082015282604082015260008382015260006080820152015290565b607f8216906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611fb257612e6e9167ffffffffffffffff612e2b8584612baf565b921660005260066020526701ffffffffffffff60406000209460071c169160036001831b921b191617929067ffffffffffffffff16600052602052604060002090565b55565b9091607f8316916801fffffffffffffffe67ffffffffffffffff84169360011b169280840460021490151715611fb257612eab8482612baf565b926004831015610be957612e6e9367ffffffffffffffff612ef0931660005260066020526003604060002094831b921b191617936701ffffffffffffff9060071c1690565b67ffffffffffffffff16600052602052604060002090565b9080602083519182815201916020808360051b8301019401926000915b838310612f3457505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08560019503018652885190606080612fb2612f8285516080865260808601906105b8565b73ffffffffffffffffffffffffffffffffffffffff878701511687860152604086015185820360408701526105b8565b93015191015297019301930191939290612f25565b90611bf1916020815267ffffffffffffffff60608351805160208501528260208201511660408501528260408201511682850152015116608082015260a06130566130226020850151610120848601526101408501906105b8565b60408501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030160c08601526105b8565b606084015173ffffffffffffffffffffffffffffffffffffffff1660e084015292608081015163ffffffff166101008401520151906101207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152612f08565b90602082519201517fffffffff00000000000000000000000000000000000000000000000000000000811692600481106130f3575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b6131646131906131ba926040519283917f45ec4b5f00000000000000000000000000000000000000000000000000000000602084015260248301612fc7565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610158565b5a7f00000000000000000000000000000000000000000000000000000000000000009130906138c6565b5090156131cd5750600290611bf1610580565b9072c11c11c11c11c11c11c11c11c11c11c11c11c133146131ef575b60039190565b6132206131fb836130bb565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f37c3be29000000000000000000000000000000000000000000000000000000001480156132e3575b80156132af575b156131e9576112e5613261836130bb565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b506132bc6131fb836130bb565b7fea7f4b120000000000000000000000000000000000000000000000000000000014613250565b506132f06131fb836130bb565b7fafa32a2c0000000000000000000000000000000000000000000000000000000014613249565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181613396575b5061339257826129436126ff565b9150565b90916020823d6020116133c5575b816133b160209383610158565b810103126133c25750519038613384565b80fd5b3d91506133a4565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff0000000000000000000000000000000000000000000000000000000060248301526024825261342d604483610158565b6179185a10613469576020926000925191617530fa6000513d8261345d575b5081613456575090565b9050151590565b6020111591503861344c565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a70000000000000000000000000000000000000000000000000000000060248301526024825261342d604483610158565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a70000000000000000000000000000000000000000000000000000000085521660248301526024825261342d604483610158565b80548210156119f85760005260206000200190600090565b6003548110156119f85760036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b015490565b604080519091906135b38382610158565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001366020840137565b604051906135f1602083610158565b6000808352366020840137565b6040516060919061360f8382610158565b60028152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001366020840137565b9080601f8301121561020e578151613655816101d8565b926136636040519485610158565b81845260208085019260051b82010192831161020e57602001905b82821061368b5750505090565b60208091835161369a816101f0565b81520191019061367e565b909160608284031261020e57815167ffffffffffffffff811161020e57836136ce91840161363e565b92602083015167ffffffffffffffff811161020e576040916136f191850161363e565b92015160ff8116810361020e5790565b9190803b1580156138b4575b6137af576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9390931660048401526000908390602490829073ffffffffffffffffffffffffffffffffffffffff165afa90811561117b5760008093600093613784575b50929190565b919350506137a591503d806000833e61379d8183610158565b8101906136a5565b929092913861377e565b5090610ba76137d29167ffffffffffffffff166000526005602052604060002090565b906060820173ffffffffffffffffffffffffffffffffffffffff61380a825173ffffffffffffffffffffffffffffffffffffffff1690565b161561387e57613867613870916117d461384360406138276135fe565b97015173ffffffffffffffffffffffffffffffffffffffff1690565b61384c876119eb565b9073ffffffffffffffffffffffffffffffffffffffff169052565b61384c846119fd565b6138786135e2565b90600090565b506138706138ab604061388f6135a2565b94015173ffffffffffffffffffffffffffffffffffffffff1690565b61384c846119eb565b506138c1610f6682612b85565b61370d565b9391936138d3608461022f565b946138e16040519687610158565b608486526138ef608461022f565b947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602088019601368737833b156139c2575a90808210613998578291038060061c9003111561396e576000918291825a9560208451940192f1905a9003923d9060848211613965575b6000908287523e929190565b60849150613959565b7f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b80600052600460205260406000205415600014613a79576003546801000000000000000081101561013757806001613a2992016003556003613555565b81549060031b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84831b921b1916179055613a7360035491600490600052602052604060002090565b55600190565b5060009056fea164736f6c634300081a000a",
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
	return common.HexToHash("0x58a20cdf97a4562295fa419a74c9bdf2683d21773d052231dc4da284a495bfb0")
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
