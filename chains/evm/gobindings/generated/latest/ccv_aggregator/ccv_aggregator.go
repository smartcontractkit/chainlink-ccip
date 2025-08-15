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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.AggregatedReport\",\"components\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvBlobs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"proofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMultiProofMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Any2EVMMultiProofTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.DynamicConfig\",\"components\":[{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"defaultCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461053e57613bd0803803809161001d8285610574565b833981019080820360c0811261053e576080811261053e576040519161004283610559565b61004b81610597565b8352602081015161ffff8116810361053e57602084019081526040820151906001600160a01b038216820361053e57604085019182526060830151936001600160a01b038516850361053e5760209060608701958652607f19011261053e5760405192602084016001600160401b03811185821017610543576040526100d3608082016105ab565b845260a0810151906001600160401b03821161053e570186601f8201121561053e578051966001600160401b038811610543578760051b916040519861011c602085018b610574565b89526020808a01938201019082821161053e5760208101935b82851061043e575050505050331561042d57600180546001600160a01b0319163317905581516001600160a01b031615801561041b575b61040a5784516001600160401b0316156103f95784516001600160401b03908116608090815283516001600160a01b0390811660a0528651811660c052835161ffff90811660e05260408051995190941689529351909316602088810191909152935183169187019190915293511660608501527f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b369390927f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a151151560ff196002541660ff821617600255604051908152a16000905b80518210156103a35760009160208160051b8301015160018060401b036020820151169081156103945780516001600160a01b0316156103855781855260056020526040852094606082015180518015918215610359575b505061034a57506040810151855491516001600160a81b031990921690151560a01b60ff60a01b16176001600160a01b03919091161784559192600192907f58a20cdf97a4562295fa419a74c9bdf2683d21773d052231dc4da284a495bfb090608090610308846105b8565b5060408051825460a089811b8a9003808316845291901c60ff161515602083015288840154811692820192909252600290920154166060820152a20190610244565b6342bcdf7f60e11b8152600490fd5b6020919250012060405160208101908382526020815261037a604082610574565b51902014388061029c565b6342bcdf7f60e11b8552600485fd5b63c656089560e01b8552600485fd5b604051613584908161064c82396080518181816104610152610f86015260a0518181816104c40152610eea015260c0518181816105000152611fae015260e05181818161048801528181611b660152612df60152f35b63c656089560e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b5083516001600160a01b03161561016c565b639b15e16f60e01b60005260046000fd5b84516001600160401b03811161053e5782016080818603601f19011261053e576040519061046b82610559565b60208101516001600160a01b038116810361053e57825261048e60408201610597565b602083015261049f606082016105ab565b604083015260808101516001600160401b03811161053e57602091010185601f8201121561053e5780516001600160401b03811161054357604051916104ef601f8301601f191660200184610574565b818352876020838301011161053e5760005b8281106105295750509181600060208096949581960101526060820152815201940193610135565b80602080928401015182828701015201610501565b600080fd5b634e487b7160e01b600052604160045260246000fd5b608081019081106001600160401b0382111761054357604052565b601f909101601f19168101906001600160401b0382119082101761054357604052565b51906001600160401b038216820361053e57565b5190811515820361053e57565b80600052600460205260406000205415600014610645576003546801000000000000000081101561054357600181018060035581101561062f577fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0181905560035460009182526004602052604090912055600190565b634e487b7160e01b600052603260045260246000fd5b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806304666f9c146100e757806306285c69146100e2578063181f5a77146100dd57806345ec4b5f146100d85780634708209d146100d35780635215505b146100ce5780635e36480c146100c95780637437ff9f146100c457806379ba5097146100bf5780638cdf6d78146100ba5780638da5cb5b146100b5578063e9d68a8e146100b05763f2fde38b146100ab57600080fd5b611407565b611351565b6112b4565b610dd8565b610cef565b610c86565b610c25565b610a8c565b61090c565b6108b5565b6105fb565b61040f565b6102b0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761013757604052565b6100ec565b6020810190811067ffffffffffffffff82111761013757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761013757604052565b604051906101a860c083610158565b565b604051906101a860a083610158565b604051906101a861010083610158565b604051906101a8604083610158565b67ffffffffffffffff81116101375760051b60200190565b73ffffffffffffffffffffffffffffffffffffffff81160361020e57565b600080fd5b67ffffffffffffffff81160361020e57565b8015150361020e57565b67ffffffffffffffff811161013757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561020e578035906102808261022f565b9261028e6040519485610158565b8284526020838301011161020e57816000926020809301838601378301015290565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5760043567ffffffffffffffff811161020e573660238201121561020e5780600401359061030b826101d8565b906103196040519283610158565b8282526024602083019360051b8201019036821161020e5760248101935b82851061034957610347846114fb565b005b843567ffffffffffffffff811161020e57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc823603011261020e57604051916103958361011b565b60248201356103a3816101f0565b835260448201356103b381610213565b602084015260648201356103c681610225565b604084015260848201359267ffffffffffffffff841161020e576103f4602094936024869536920101610269565b6060820152815201940193610337565b600091031261020e57565b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e57610446611769565b5061057c6040516104568161011b565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b6040519061058f602083610158565b60008252565b60005b8381106105a85750506000910152565b8181015183820152602001610598565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936105f481518092818752878088019101610595565b0116010190565b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5761057c604080519061063c8183610158565b601782527f43435641676772656761746f7220312e372e302d6465760000000000000000006020830152519182916020835260208301906105b8565b919082608091031261020e576040516106908161011b565b60608082948035845260208101356106a781610213565b602085015260408101356106ba81610213565b60408501520135916106cb83610213565b0152565b35906101a8826101f0565b359063ffffffff8216820361020e57565b81601f8201121561020e57803590610702826101d8565b926107106040519485610158565b82845260208085019360051b8301019181831161020e5760208101935b83851061073c57505050505090565b843567ffffffffffffffff811161020e57820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603011261020e57604051916107888361011b565b602082013567ffffffffffffffff811161020e578560206107ab92850101610269565b835260408201356107bb816101f0565b602084015260608201359267ffffffffffffffff841161020e576080836107e9886020809881980101610269565b60408401520135606082015281520194019361072d565b9190916101208184031261020e57610816610199565b926108218183610678565b8452608082013567ffffffffffffffff811161020e5781610843918401610269565b602085015260a082013567ffffffffffffffff811161020e5781610868918401610269565b604085015261087960c083016106cf565b606085015261088a60e083016106da565b608085015261010082013567ffffffffffffffff811161020e576108ae92016106eb565b60a0830152565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5760043567ffffffffffffffff811161020e57610907610347913690600401610800565b611a32565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e577f38c8130879f9e78081031dd00c0318d4d754b05105d2090bdc1f43e5cbe20b366109c560405161096b8161013c565b60043561097781610225565b8152610981611f19565b51600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151560ff81169290921790556040519081529081906020820190565b0390a1005b6040810160408252825180915260206060830193019060005b818110610a6c5750505060208183039101526020808351928381520192019060005b818110610a125750505090565b9091926020608082610a61600194885173ffffffffffffffffffffffffffffffffffffffff60608092828151168552602081015115156020860152826040820151166040860152015116910152565b019401929101610a05565b825167ffffffffffffffff168552602094850194909201916001016109e3565b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e57600354610ac7816101d8565b90610ad56040519283610158565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610b02826101d8565b0160005b818110610bc8575050610b1881611cfd565b9060005b818110610b3457505061057c604051928392836109ca565b80610b6c610b53610b466001946132d3565b67ffffffffffffffff1690565b610b5d8387611d7b565b9067ffffffffffffffff169052565b610bac610ba7610b8d610b7f8488611d7b565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b611d94565b610bb68287611d7b565b52610bc18186611d7b565b5001610b1c565b602090610bd3611769565b82828701015201610b06565b60041115610be957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b906004821015610be95752565b3461020e5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e576020610c77600435610c6581610213565b60243590610c7282610213565b611e25565b610c846040518092610c18565bf35b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e576000604051610cc38161013c565b5261057c604051610cd38161013c565b60025460ff161515908190526040519081529081906020820190565b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5760005473ffffffffffffffffffffffffffffffffffffffff81163303610dae577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5760043567ffffffffffffffff811161020e5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82600401923603011261020e5760025460ff1661128a57610e8360017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006002541617600255565b610e986020610e928380611e7b565b01611eae565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b1660048201526020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa90811561128557600091611256575b5061121f57610f67610f63610f598367ffffffffffffffff166000526005602052604060002090565b5460a01c60ff1690565b1590565b6111e857610f7a6040610e928480611e7b565b67ffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036111a057610fe9610fe4610fdf610fce610fc86060610e928880611e7b565b85611e25565b94610fd986826128be565b80611e7b565b611ecd565b612a2a565b91610ff381610bdf565b801590811561118c575b50156111365781516060015161101d9067ffffffffffffffff1682612a75565b61102682612db4565b9261104882611042606084510167ffffffffffffffff90511690565b85612b00565b61105182610bdf565b600282036110e2575b67ffffffffffffffff7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df291516110b261109e606083015167ffffffffffffffff1690565b915196836040519485941697169583611efb565b0390a46103477fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060025416600255565b6110eb82610bdf565b600382031561105a575160600151611132929067ffffffffffffffff16907f926c5a3e00000000000000000000000000000000000000000000000000000000600052611ed8565b6000fd5b611150606061113293510167ffffffffffffffff90511690565b7f3b5754190000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff91821660045216602452604490565b6003915061119981610bdf565b1438610ffd565b6111326111b26040610e928580611e7b565b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b611278915060203d60201161127e575b6112708183610158565b810190611eb8565b38610f30565b503d611266565b611a12565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b3461020e5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b6101a890929192608081019373ffffffffffffffffffffffffffffffffffffffff60608092828151168552602081015115156020860152826040820151166040860152015116910152565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5767ffffffffffffffff60043561139581610213565b61139d611769565b5016600052600560205261057c604060002073ffffffffffffffffffffffffffffffffffffffff6002604051926113d38461011b565b60ff8154848116865260a01c1615156020850152826001820154166040850152015416606082015260405191829182611306565b3461020e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261020e5773ffffffffffffffffffffffffffffffffffffffff600435611457816101f0565b61145f611f19565b163381146114d157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b611503611f19565b60005b8151811015611765576115198183611d7b565b5190611530602083015167ffffffffffffffff1690565b67ffffffffffffffff811690811561173b5761157f611566611566865173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b156116fb576115a29067ffffffffffffffff166000526005602052604060002090565b606084015180518015918215611725575b50506116fb576116f27f58a20cdf97a4562295fa419a74c9bdf2683d21773d052231dc4da284a495bfb0916116986116578761163d6115f7604060019b0151151590565b85547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178555565b5173ffffffffffffffffffffffffffffffffffffffff1690565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6116a18461342e565b5060405191829182919091606073ffffffffffffffffffffffffffffffffffffffff6002608084019560ff8154848116875260a01c1615156020860152826001820154166040860152015416910152565b0390a201611506565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b602001209050611733611f64565b1438806115b3565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b5050565b604051906117768261011b565b60006060838281528260208201528260408201520152565b604051906040820182811067ffffffffffffffff8211176101375760405260006020838281520152565b604051906117c7602083610158565b600080835282815b8281106117db57505050565b6020906117e661178e565b828285010152016117cf565b906117fc826101d8565b6118096040519182610158565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061183782946101d8565b019060005b82811061184857505050565b60209061185361178e565b8282850101520161183c565b909160608284031261020e57815161187681610225565b92602083015167ffffffffffffffff811161020e5783019080601f8301121561020e578151916118a58361022f565b916118b36040519384610158565b8383526020848301011161020e576040926118d49160208085019101610595565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a0840152608061195161191d604084015160a060c08801526101208701906105b8565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526105b8565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106119da5750505061ffff90951660208301526101a892916060916119be9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101611990565b6040513d6000823e3d90fd5b906020611a2f9281815201906105b8565b90565b303303611cd357611a416117b8565b9060a08101518051611c82575b50805191611a696020845194015167ffffffffffffffff1690565b906020830151916040840192611a97845192611a836101aa565b97885267ffffffffffffffff166020880152565b60408601526060850152608084015251511580611c64575b8015611c41575b8015611c0f575b61176557611b8f9181611b02611566611ae8610b8d6020600097510167ffffffffffffffff90511690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9083611b366060611b1a608085015163ffffffff1690565b93015173ffffffffffffffffffffffffffffffffffffffff1690565b93604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f000000000000000000000000000000000000000000000000000000000000000090600486016118da565b03925af190811561128557600090600092611be8575b5015611bae5750565b611be4906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611a1e565b0390fd5b9050611c0791503d806000833e611bff8183610158565b81019061185f565b509038611ba5565b50611c3c610f63611c37606084015173ffffffffffffffffffffffffffffffffffffffff1690565b612375565b611abd565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b15611ab6565b5063ffffffff611c7b608083015163ffffffff1690565b1615611aaf565b8192506020611ccc920151611cae606085015173ffffffffffffffffffffffffffffffffffffffff1690565b90611cc6602086510167ffffffffffffffff90511690565b92611f85565b9038611a4e565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b90611d07826101d8565b611d146040519182610158565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611d4282946101d8565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051821015611d8f5760209160051b010190565b611d4c565b90604051611da18161011b565b606073ffffffffffffffffffffffffffffffffffffffff6002839560ff8154848116875260a01c1615156020860152826001820154166040860152015416910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211611e2057565b611de4565b611e3182607f9261241f565b9116906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611e20576003911c166004811015610be95790565b6040519061058f8261013c565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee18136030182121561020e570190565b35611a2f81610213565b9081602091031261020e5751611a2f81610225565b611a2f903690610800565b929160449067ffffffffffffffff6101a893816064971660045216602452610c18565b80611f0c604092611a2f9594610c18565b81602082015201906105b8565b73ffffffffffffffffffffffffffffffffffffffff600154163303611f3a57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101906000825260208152611f7f604082610158565b51902090565b93909193611f9381516117f2565b9260009573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016925b805188101561236a57611fe38882611d7b565b5197611fed61178e565b5061206661201260208b015173ffffffffffffffffffffffffffffffffffffffff1690565b9960208b604051809481927fbbe4f6db0000000000000000000000000000000000000000000000000000000083526004830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b03818a5afa9182156112855760009261233a575b5073ffffffffffffffffffffffffffffffffffffffff82169182156122f7576120a2816123cb565b6122f7576120b2610f63826123f5565b6122f7575061218a6020878c8e946121406120cd878c6134c1565b966120d6611e6e565b50606083015161210260408551950151956120ef6101b9565b97885267ffffffffffffffff1688880152565b73ffffffffffffffffffffffffffffffffffffffff8d166040870152606086015273ffffffffffffffffffffffffffffffffffffffff166080850152565b60a083015260c0820152612152610580565b60e0820152604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301612fad565b03816000875af1600091816122c7575b506121de578b6121a86130d6565b90611be46040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401613106565b9a92939495969798999a9173ffffffffffffffffffffffffffffffffffffffff87160361225c575b509060019291516122346122186101c9565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b6020820152612243828a611d7b565b5261224e8189611d7b565b500196959493929190611fd0565b61226683876134c1565b9080821080156122b3575b61227b5750612206565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b506122be8183611e13565b83511415612271565b6122e991925060203d81116122f0575b6122e18183610158565b810190612f8f565b903861219a565b503d6122d7565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b61235c91925060203d8111612363575b6123548183610158565b810190612f7a565b903861207a565b503d61234a565b505050509250905090565b61239f7f85572ffb0000000000000000000000000000000000000000000000000000000082613259565b90816123b9575b816123af575090565b611a2f91506131f9565b90506123c481613133565b15906123a6565b61239f7ff208a58f0000000000000000000000000000000000000000000000000000000082613259565b61239f7faff2afbf0000000000000000000000000000000000000000000000000000000082613259565b9067ffffffffffffffff612461921660005260066020526701ffffffffffffff60406000209160071c1667ffffffffffffffff16600052602052604060002090565b5490565b35611a2f816101f0565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561020e570180359067ffffffffffffffff821161020e57602001918160051b3603831361020e57565b9190811015611d8f5760051b0190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561020e57016020813591019167ffffffffffffffff821161020e57813603831361020e57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561020e57016020813591019167ffffffffffffffff821161020e578160051b3603831361020e57565b90602083828152019160208260051b8501019381936000915b8483106125de5750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818636030182121561020e57602080918760019401906060806126a761266561265786806124d3565b608087526080870191612523565b73ffffffffffffffffffffffffffffffffffffffff87870135612687816101f0565b168786015261269960408701876124d3565b908683036040880152612523565b93013591015298019301930191949392906125ce565b90611a2f91602081528135602082015267ffffffffffffffff60208301356126e481610213565b16604082015267ffffffffffffffff604083013561270181610213565b16606082015267ffffffffffffffff606083013561271e81610213565b1660808201526127df61278a61274d61273a60808601866124d3565b61012060a0870152610140860191612523565b61275a60a08601866124d3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160c0870152612523565b926127b761279a60c083016106cf565b73ffffffffffffffffffffffffffffffffffffffff1660e0850152565b6127d46127c660e083016106da565b63ffffffff16610100850152565b610100810190612562565b916101207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101526125b5565b9190811015611d8f5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561020e57019081359167ffffffffffffffff831161020e57602001823603811361020e579190565b92946128a36128b1936101a8989a99976128956080989560a0895260a08901906105b8565b918783036020890152612523565b918483036040860152612523565b9560608201520190610c18565b91906128cf6020610e928580611e7b565b506128e560c06128df8580611e7b565b01612465565b50602083016128f4818561246f565b505060005b612903828661246f565b9050811015612a235761292e61156661156661292984612923878b61246f565b906124c3565b612465565b9061294e61297a61293f8880611e7b565b604051928391602083016126bd565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610158565b6129918261298b60408a018a61246f565b90612810565b90936129a48461298b60608c018c61246f565b829691963b1561020e576000946129ee8a8888946040519b8c998a9889977fb9429a6a00000000000000000000000000000000000000000000000000000000895260048901612870565b03925af191821561128557600192612a08575b50016128f9565b80612a176000612a1d93610158565b80610404565b38612a01565b5050509050565b60405160c0810181811067ffffffffffffffff8211176101375760609160a091604052612a55611769565b815282602082015282604082015260008382015260006080820152015290565b607f8216906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611e2057612afd9167ffffffffffffffff612aba858461241f565b921660005260066020526701ffffffffffffff60406000209460071c169160036001831b921b191617929067ffffffffffffffff16600052602052604060002090565b55565b9091607f8316916801fffffffffffffffe67ffffffffffffffff84169360011b169280840460021490151715611e2057612b3a848261241f565b926004831015610be957612afd9367ffffffffffffffff612b7f931660005260066020526003604060002094831b921b191617936701ffffffffffffff9060071c1690565b67ffffffffffffffff16600052602052604060002090565b9080602083519182815201916020808360051b8301019401926000915b838310612bc357505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08560019503018652885190606080612c41612c1185516080865260808601906105b8565b73ffffffffffffffffffffffffffffffffffffffff878701511687860152604086015185820360408701526105b8565b93015191015297019301930191939290612bb4565b90611a2f916020815267ffffffffffffffff60608351805160208501528260208201511660408501528260408201511682850152015116608082015260a0612ce5612cb16020850151610120848601526101408501906105b8565b60408501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030160c08601526105b8565b606084015173ffffffffffffffffffffffffffffffffffffffff1660e084015292608081015163ffffffff166101008401520151906101207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152612b97565b90602082519201517fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612d82575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b61294e612df3612e1d926040519283917f45ec4b5f00000000000000000000000000000000000000000000000000000000602084015260248301612c56565b5a7f0000000000000000000000000000000000000000000000000000000000000000913090613308565b509015612e305750600290611a2f610580565b9072c11c11c11c11c11c11c11c11c11c11c11c11c13314612e52575b60039190565b612e83612e5e83612d4a565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f37c3be2900000000000000000000000000000000000000000000000000000000148015612f46575b8015612f12575b15612e4c57611132612ec483612d4a565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b50612f1f612e5e83612d4a565b7fea7f4b120000000000000000000000000000000000000000000000000000000014612eb3565b50612f53612e5e83612d4a565b7fafa32a2c0000000000000000000000000000000000000000000000000000000014612eac565b9081602091031261020e5751611a2f816101f0565b9081602091031261020e5760405190612fa78261013c565b51815290565b90611a2f916020815260e06130a261306f612fd6855161010060208701526101208601906105b8565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff1660608601526060860151608086015261303b608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526105b8565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526105b8565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526105b8565b3d15613101573d906130e78261022f565b916130f56040519384610158565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff611a2f949316815281602082015201906105b8565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252613193604483610158565b6179185a106131cf576020926000925191617530fa6000513d826131c3575b50816131bc575090565b9050151590565b602011159150386131b2565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252613193604483610158565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252613193604483610158565b8054821015611d8f5760005260206000200190600090565b600354811015611d8f5760036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b015490565b939193613315608461022f565b946133236040519687610158565b60848652613331608461022f565b947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602088019601368737833b15613404575a908082106133da578291038060061c900311156133b0576000918291825a9560208451940192f1905a9003923d90608482116133a7575b6000908287523e929190565b6084915061339b565b7f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fd5b806000526004602052604060002054156000146134bb57600354680100000000000000008110156101375780600161346b920160035560036132bb565b81549060031b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84831b921b19161790556134b560035491600490600052602052604060002090565b55600190565b50600090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181613540575b5061353c57826121a86130d6565b9150565b90916020823d60201161356f575b8161355b60209383610158565b8101031261356c575051903861352e565b80fd5b3d915061354e56fea164736f6c634300081a000a",
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
