// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offramp

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

type MessageV1CodecMessageV1 struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
	MessageNumber       uint64
	ExecutionGasLimit   uint32
	CcipReceiveGasLimit uint32
	Finality            [4]byte
	CcvAndExecutorHash  [32]byte
	OnRampAddress       []byte
	OffRampAddress      []byte
	Sender              []byte
	Receiver            []byte
	DestBlob            []byte
	TokenTransfer       []MessageV1CodecTokenTransferV1
	Data                []byte
}

type MessageV1CodecTokenTransferV1 struct {
	Amount             *big.Int
	SourcePoolAddress  []byte
	SourceTokenAddress []byte
	DestTokenAddress   []byte
	TokenReceiver      []byte
	ExtraData          []byte
}

type OffRampSourceChainConfig struct {
	Router           common.Address
	IsEnabled        bool
	OnRamps          [][]byte
	DefaultCCVs      []common.Address
	LaneMandatedCCVs []common.Address
}

type OffRampSourceChainConfigArgs struct {
	Router              common.Address
	SourceChainSelector uint64
	IsEnabled           bool
	OnRamps             [][]byte
	DefaultCCVs         []common.Address
	LaneMandatedCCVs    []common.Address
}

type OffRampStaticConfig struct {
	LocalChainSelector        uint64
	GasForCallExactCheck      uint16
	RmnRemote                 common.Address
	TokenAdminRegistry        common.Address
	MaxGasBufferToUpdateState uint32
}

var OffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"sourceChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"sourceChainConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GasCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidGasLimitOverride\",\"inputs\":[{\"name\":\"messageGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalThreshold\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoStateProgressMade\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120604052346102af57604051601f61618838819003918201601f19168301916001600160401b038311848410176102b45780849260a0946040528339810103126102af576040519060009060a083016001600160401b0381118482101761029b5760405280516001600160401b0381168103610297578352602081015161ffff8116810361029757602084019081526040820151906001600160a01b038216820361029357604085019182526060830151926001600160a01b038416840361028f5760608601938452608001519363ffffffff8516850361028c5760808601948552331561027d57600180546001600160a01b0319163317905582516001600160a01b031615801561026b575b61025c5785516001600160401b03161561024d5761ffff8251161561023e5763ffffffff8551161561023e5785516001600160401b03908116608090815284516001600160a01b0390811660a09081528751821660c052855161ffff90811660e052895163ffffffff90811661010052604080518d519097168752885190921660208701528851841691860191909152885190921660608501528851909116918301919091527f6db4162777b6c980e778bb05a3d9e050f3792b091287ff0d4f3d51bdcd7427db91a1604051615ebd90816102cb823960805181818161013e01526119da015260a0518181816101a10152611901015260c0518181816101dd01528181614b860152615627015260e05181818161016501526150ba01526101005181818161020901526142e70152f35b632855a4d960e11b8152600490fd5b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b03161561010e565b639b15e16f60e01b8152600490fd5b80fd5b8480fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780633b81ab9b146100c857806349d8033e146100c35780635215505b146100be5780635643a782146100b957806379ba5097146100b45780638da5cb5b146100af578063d6162963146100aa578063e9d68a8e146100a55763f2fde38b146100a057600080fd5b61115f565b610f9b565b610eda565b610ea6565b610ddb565b610a1e565b610907565b610768565b61068f565b61051b565b610411565b6100ec565b60009103126100e757565b600080fd5b346100e75760006003193601126100e7576000608060405161010d816102dc565b82815282602082015282604082015282606082015201526102a9604051610133816102dc565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016608082015260405191829182919091608063ffffffff8160a084019567ffffffffffffffff815116855261ffff602082015116602086015273ffffffffffffffffffffffffffffffffffffffff604082015116604086015273ffffffffffffffffffffffffffffffffffffffff6060820151166060860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff8211176102f857604052565b6102ad565b6101c0810190811067ffffffffffffffff8211176102f857604052565b6020810190811067ffffffffffffffff8211176102f857604052565b90601f601f19910116810190811067ffffffffffffffff8211176102f857604052565b6040519061036860c083610336565b565b6040519061036860a083610336565b6040519061036861010083610336565b60405190610368604083610336565b67ffffffffffffffff81116102f857601f01601f191660200190565b604051906103c3602083610336565b60008252565b60005b8381106103dc5750506000910152565b81810151838201526020016103cc565b90601f19601f60209361040a815180928187528780880191016103c9565b0116010190565b346100e75760006003193601126100e7576102a960408051906104348183610336565b600d82527f4f666652616d7020322e302e30000000000000000000000000000000000000006020830152519182916020835260208301906103ec565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104bc5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104af565b9161051460ff9161050660409497969760608752606087019061049e565b90858203602087015261049e565b9416910152565b346100e75760206003193601126100e75760043567ffffffffffffffff81116100e75761054f610555913690600401610470565b906137d2565b61014081018051601481510361060b576102a96105fc846105768551611246565b60601c9061058c815167ffffffffffffffff1690565b91610120820151610180830151916105f6816105cb60a08701517fffffffff000000000000000000000000000000000000000000000000000000001690565b956105f06105e760806101a08401515193015163ffffffff1690565b63ffffffff1690565b90613cff565b94613e7d565b604093919351938493846104e8565b610641906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611235565b0390fd5b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b63ffffffff8116036100e757565b359061036882610676565b346100e75760806003193601126100e75760043567ffffffffffffffff81116100e7576106c0903690600401610470565b9060243567ffffffffffffffff81116100e7576106e1903690600401610645565b926044359367ffffffffffffffff85116100e75761070661071b953690600401610645565b9390926064359561071687610676565b6117e8565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561075657565b61071d565b9060048210156107565752565b346100e75760206003193601126100e7576004356000526006602052602060ff6040600020541661079c604051809261075b565bf35b9080602083519182815201916020808360051b8301019401926000915b8383106107ca57505050505090565b90919293946020806107e883601f19866001960301875289516103ec565b970193019301919392906107bb565b6108629173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061085161083f604085015160a0604086015260a085019061079e565b6060850151848203606086015261049e565b92015190608081840391015261049e565b90565b6040810160408252825180915260206060830193019060005b8181106108e7575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106108ba57505050505090565b90919293946020806108d883601f19866001960301875289516107f7565b970193019301919392906108ab565b825167ffffffffffffffff1685526020948501949092019160010161087e565b346100e75760006003193601126100e75760025461092481611f51565b906109326040519283610336565b808252601f1961094182611f51565b0160005b818110610a0757505061095781611f95565b9060005b8181106109735750506102a960405192839283610865565b806109ab610992610985600194615923565b67ffffffffffffffff1690565b61099c8387612007565b9067ffffffffffffffff169052565b6109eb6109e66109cc6109be8488612007565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b6120cf565b6109f58287612007565b52610a008186612007565b500161095b565b602090610a12611f69565b82828701015201610945565b346100e75760206003193601126100e75760043567ffffffffffffffff81116100e757610a4f903690600401610645565b90610a586143c7565b6000905b828210610a6557005b610a78610a73838584612274565b6123a2565b6020810191610a92610985845167ffffffffffffffff1690565b15610db157610ad4610abb610abb845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610da4575b610b2b5760808201949060005b86518051821015610b5557610abb610b0483610b1e93612007565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610b2b57600101610ae9565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610b8d57610abb610b0483610b8093612007565b15610b2b57600101610b65565b505095929491909394610ba38651825190614412565b610bb86109cc835167ffffffffffffffff1690565b90610be8610bce845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610bf287615958565b606085019860005b8a518051821015610c4a5790610c1281602093612007565b518051928391012091158015610c3a575b610b2b57610c336001928b615a27565b5001610bfa565b50610c43612457565b8214610c23565b5050976001975093610d61610d9a946003610d8d95610d597f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e99610985979f610ca760408e610ca0610ced945160018b01612610565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610d4f610d0e8d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b51600285016126f5565b5191016126f5565b610d7e610d79610985835167ffffffffffffffff1690565b615996565b505167ffffffffffffffff1690565b9260405191829182612789565b0390a20190610a5c565b5060808201515115610adc565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760006003193601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303610e7c577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760006003193601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100e75760a06003193601126100e75760043567ffffffffffffffff81116100e7576101c060031982360301126100e7576024359060443567ffffffffffffffff81116100e757610f30903690600401610645565b926064359367ffffffffffffffff85116100e757610f5561071b953690600401610645565b93909260843595610f6587610676565b600401612f24565b67ffffffffffffffff8116036100e757565b359061036882610f6d565b9060206108629281815201906107f7565b346100e75760206003193601126100e75767ffffffffffffffff600435610fc181610f6d565b610fc9611f69565b501660005260046020526040600020604051610fe4816102dc565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c16151560208201526001820180549061101c82611f51565b9161102a6040519384610336565b80835260208301916000526020600020916000905b82821061107e576102a98661106d60038a8960408501526110626002820161206e565b60608501520161206e565b608082015260405191829182610f8a565b6040516000855461108e8161201b565b808452906001811690811561110057506001146110c8575b50600192826110ba85946020940382610336565b81520194019101909261103f565b6000878152602081209092505b8183106110ea575050810160200160016110a6565b60018160209254838688010152019201916110d5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506110a6565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b346100e75760206003193601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561119181611141565b6111996143c7565b1633811461120b57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060206108629281815201906103ec565b90602082519201517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116926014811061127e575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b801515036100e757565b908160209103126100e75751610862816112b0565b6040513d6000823e3d90fd5b60409073ffffffffffffffffffffffffffffffffffffffff610862949316815281602082015201906103ec565b92919261131482610398565b916113226040519384610336565b8294818452818301116100e7578281602093846000960137010152565b9060048110156107565760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b8383106113a257505050505090565b909192939460208061142983601f1986600196030187528951908151815260a06114186114066113f46113e28887015160c08a88015260c08701906103ec565b604087015186820360408801526103ec565b606086015185820360608701526103ec565b608085015184820360808601526103ec565b9201519060a08184039101526103ec565b97019301930191939290611393565b9160209082815201919060005b8181106114525750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561147b81611141565b168152019401929101611445565b601f8260209493601f19938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b8484106115225750505050505090565b90919293949560208061154a83601f1986600196030188526115448b886114aa565b90611489565b9801940194019294939190611512565b97969491611798906080956117a6956117856103689a9561160a8e60a0815261159060a08201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660c0820152604083015167ffffffffffffffff1660e0820152606083015163ffffffff16610100820152828c015163ffffffff1661012082015261014060a08401519101907fffffffff00000000000000000000000000000000000000000000000000000000169052565b8d61016060c08301519101528d6101a061175161171b6116e56116af61167961164560e08901516101c06101808a01526102608901906103ec565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6089830301888a01526103ec565b6101208801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60888303016101c08901526103ec565b6101408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016101e08801526103ec565b6101608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303016102008701526103ec565b6101808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6085830301610220860152611376565b920151906102407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60828503019101526103ec565b9260208d01528b830360408d0152611438565b9188830360608a01526114fa565b94019063ffffffff169052565b6040906108629392815281602082015201906103ec565b806117db604092610862959461075b565b81602082015201906103ec565b6001549596919560a01c60ff16611f275761183d740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b61184786826137d2565b956118e8602061188d6118656109858b5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611f2257600091611ef3575b50611ea85761196061195c6119526109cc8a5167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b611e5d57611979610bce885167ffffffffffffffff1690565b6119a361195c60e08a01928351602081519101209060019160005201602052604060002054151590565b611e2657506101008701516014815114801590611df5575b611dbe5750602087015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603611d875750878503611d5357610140870151601481510361060b575063ffffffff83168015159081611d36575b50611ce95790611a44913691611308565b6020815191012095611a6a611a63886000526006602052604060002090565b5460ff1690565b94611a748661074c565b85158015611cd6575b15611c6457611b1792611b1c959492611b0992611ad2611aa78c6000526006602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957fd61629630000000000000000000000000000000000000000000000000000000060208801528c8c6024890161155a565b03601f198101835282610336565b6142c0565b9181159081611c50575b50611c195715611be7577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b611b7b84611b76886000526006602052604060002090565b61133f565b611bb7611ba56040611b95885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b918360405194859416971695836117ca565b0390a46103687fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff600392611b5e565b50826106416040519283927f10afb5b6000000000000000000000000000000000000000000000000000000008452600484016117b3565b60039150611c5d8161074c565b1438611b26565b611cd28888611c906040611c80835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b50611ce08661074c565b60038614611a7d565b611cd283611cfe60808a015163ffffffff1690565b7fdf2964df0000000000000000000000000000000000000000000000000000000060005263ffffffff90811660045216602452604490565b9050611d4c6105e760808a015163ffffffff1690565b1138611a33565b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004859052602488905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b610641906040519182917f55216e3100000000000000000000000000000000000000000000000000000000835230600484016112db565b50611e08611e0282611246565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff163014156119bb565b61064190516040519182917fa50bd14700000000000000000000000000000000000000000000000000000000835260048301611235565b611cd2611e72885167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611cd2611ebd885167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611f15915060203d602011611f1b575b611f0d8183610336565b8101906112ba565b38611932565b503d611f03565b6112cf565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116102f85760051b60200190565b60405190611f76826102dc565b6060608083600081526000602082015282604082015282808201520152565b90611f9f82611f51565b611fac6040519182610336565b828152601f19611fbc8294611f51565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156120025760200190565b611fc6565b80518210156120025760209160051b010190565b90600182811c92168015612064575b602083101461203557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161202a565b906040519182815491828252602082019060005260206000209260005b8181106120a057505061036892500383610336565b845473ffffffffffffffffffffffffffffffffffffffff1683526001948501948794506020909301920161208b565b906040516120dc816102dc565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c16151560208301526001810180549061211682611f51565b916121246040519384610336565b80835260208301916000526020600020916000905b82821061216e5750505050600360809261216992604086015261215e6002820161206e565b60608601520161206e565b910152565b6040516000855461217e8161201b565b80845290600181169081156121f057506001146121b8575b50600192826121aa85946020940382610336565b815201940191019092612139565b6000878152602081209092505b8183106121da57505081016020016001612196565b60018160209254838688010152019201916121c5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050612196565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b9015612002578061086291612231565b90821015612002576108629160051b810190612231565b359061036882611141565b3590610368826112b0565b9080601f830112156100e75781602061086293359101611308565b9080601f830112156100e75781356122d381611f51565b926122e16040519485610336565b81845260208085019260051b820101918383116100e75760208201905b83821061230d57505050505090565b813567ffffffffffffffff81116100e757602091612330878480948801016122a1565b8152019101906122fe565b9080601f830112156100e757813561235281611f51565b926123606040519485610336565b81845260208085019260051b8201019283116100e757602001905b8282106123885750505090565b60208091833561239781611141565b81520191019061237b565b60c0813603126100e7576123b4610359565b906123be8161228b565b82526123cc60208201610f7f565b60208301526123dd60408201612296565b6040830152606081013567ffffffffffffffff81116100e75761240390369083016122bc565b6060830152608081013567ffffffffffffffff81116100e757612429903690830161233b565b608083015260a08101359067ffffffffffffffff82116100e75761244f9136910161233b565b60a082015290565b60405160208101906000825260208152612472604082610336565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106124b2575050565b600081556001016124a7565b9190601f81116124cd57505050565b610368926000526020600020906020601f840160051c830193106124f9575b601f0160051c01906124a7565b90915081906124ec565b919091825167ffffffffffffffff81116102f85761252b81612525845461201b565b846124be565b6020601f821160011461258957819061257a93949560009261257e575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612548565b601f1982169061259e84600052602060002090565b9160005b8181106125f8575095836001959697106125c1575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806125b7565b9192602060018192868b0151815501940192016125a2565b8151916801000000000000000083116102f8578154838355808410612672575b506020612644910191600052602060002090565b6000915b8383106126555750505050565b600160208261266683945186612503565b01920192019190612648565b8260005283602060002091820191015b81811061268f5750612630565b8061269c6001925461201b565b806126a9575b5001612682565b601f811183146126bf5750600081555b386126a2565b6126e39083601f6126d585600052602060002090565b920160051c820191016124a7565b600081815260208120818355556126b9565b81519167ffffffffffffffff83116102f8576801000000000000000083116102f857602090825484845580851061276c575b500190600052602060002060005b8381106127425750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501612735565b6127839084600052858460002091820191016124a7565b38612727565b90610862916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a06128026127ed606085015160c0608086015260e085019061079e565b6080850151601f19858303018486015261049e565b9201519060c0601f198285030191015261049e565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b916020610862938181520191611489565b919091357fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116926014811061127e575050565b3561086281610676565b90821015612002576129229160051b81019061286b565b9091565b908160209103126100e7575161086281611141565b60409073ffffffffffffffffffffffffffffffffffffffff61086295931681528160208201520191611489565b7fffffffff000000000000000000000000000000000000000000000000000000008116036100e757565b359061036882612968565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310612a185750505050505090565b909192939495601f1982820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e7576020612afb600193868394019081358152612aed612ae2612ac7612aac612a91612a81898801886114aa565b60c08b89015260c0880191611489565b612a9e60408801886114aa565b908783036040890152611489565b612ab960608701876114aa565b908683036060880152611489565b612ad460808601866114aa565b908583036080870152611489565b9260a08101906114aa565b9160a0818503910152611489565b980196019493019190612a08565b612d9f6108629593949260608352612b3560608401612b2783610f7f565b67ffffffffffffffff169052565b612b55612b4460208301610f7f565b67ffffffffffffffff166080850152565b612b75612b6460408301610f7f565b67ffffffffffffffff1660a0850152565b612b91612b8460608301610684565b63ffffffff1660c0850152565b612bad612ba060808301610684565b63ffffffff1660e0850152565b612be6612bbc60a08301612992565b7fffffffff0000000000000000000000000000000000000000000000000000000016610100850152565b60c0810135610120840152612d6e612d62612d23612ce4612ca5612c66612c27612c1360e08901896114aa565b6101c06101408d01526102208c0191611489565b612c356101008901896114aa565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d0152611489565b612c746101208801886114aa565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c0152611489565b612cb36101408701876114aa565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b0152611489565b612cf26101608601866114aa565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a0152611489565b612d3161018085018561299d565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e08901526129f0565b916101a08101906114aa565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa085840301610200860152611489565b9360208201526040818503910152611489565b604051906040820182811067ffffffffffffffff8211176102f85760405260006020838281520152565b90612de682611f51565b612df36040519182610336565b828152601f19612e038294611f51565b019060005b828110612e1457505050565b602090612e1f612db2565b82828501015201612e08565b3561086281610f6d565b3561086281612968565b91909160c0818403126100e757612e54610359565b9281358452602082013567ffffffffffffffff81116100e75781612e799184016122a1565b6020850152604082013567ffffffffffffffff81116100e75781612e9e9184016122a1565b6040850152606082013567ffffffffffffffff81116100e75781612ec39184016122a1565b6060850152608082013567ffffffffffffffff81116100e75781612ee89184016122a1565b608085015260a082013567ffffffffffffffff81116100e757612f0b92016122a1565b60a0830152565b91908203918211612f1f57565b612478565b9391959694909294303303613423576000612f43610180870187612817565b9050613344575b612f95612f67611e02612f616101408a018a61286b565b906128cd565b98612f8b8a612f7a6101a08b018b61286b565b90506105f06105e760808d01612901565b9889918b8a614719565b98909960005b8b518110156131585761300a60208c612fd38f85612fc5610abb610abb610b0484612fcb96612007565b93612007565b518a8c61290b565b91906040518095819482937fc3a7ded6000000000000000000000000000000000000000000000000000000008452600484016128bc565b03915afa8015611f225773ffffffffffffffffffffffffffffffffffffffff9160009161312a575b50169081156130cd57613050613048828e612007565b51888a61290b565b9290813b156100e7576000918b838e613098604051988996879586947f89e364c700000000000000000000000000000000000000000000000000000000865260048601612b09565b03925af1918215611f22576001926130b2575b5001612f9b565b806130c160006130c793610336565b806100dc565b386130ab565b8b6130f5898f6130ee856130e8610b04610641988f95612007565b95612007565b519161290b565b6040939193519384937f2665cea20000000000000000000000000000000000000000000000000000000085526004850161293b565b61314b915060203d8111613151575b6131438183610336565b810190612926565b38613032565b503d613139565b50959793509593509650965061317c613175610180840184612817565b9050612ddc565b9561318b610180840184612817565b9050613232575b5061322b57610368946131fb6131a783612e2b565b6131e86131ef6131bb61012087018761286b565b6131c96101a089018961286b565b9490956131d461036a565b9c8d5267ffffffffffffffff1660208d0152565b3691611308565b60408901523691611308565b6060860152608085015263ffffffff821615613218575091615022565b6132259150608001612901565b91615022565b5050505050565b61328c61324c613246610180860186612817565b90612264565b61325a61012086018661286b565b61328661326688612e2b565b9261327e61327660a08b01612e35565b953690612e3f565b923691611308565b90614af1565b919061329789611ff5565b5273ffffffffffffffffffffffffffffffffffffffff6132d1611e02612f616132c76132466101808a018a612817565b608081019061286b565b921673ffffffffffffffffffffffffffffffffffffffff8316036132f6575b50613192565b61332a61332f926133246133098b611ff5565b515173ffffffffffffffffffffffffffffffffffffffff1690565b906145a7565b612f12565b602061333a88611ff5565b51015238806132f0565b50601461336561335b613246610180890189612817565b606081019061286b565b90500361340f5760146133826132c7613246610180890189612817565b9050036133c5576133c06133a6611e02612f616132c76132466101808b018b612817565b613324611e02612f6161335b6132466101808c018c612817565b612f4a565b6133d96132c7613246610180880188612817565b906106416040519283927f8d666f60000000000000000000000000000000000000000000000000000000008452600484016128bc565b6133d961335b613246610180880188612817565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061345a826102fd565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b90156120025790565b90604510156120025760450190565b90821015612002570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906025116100e75760210190600490565b906045116100e75760250190602090565b90929192836046116100e75783116100e757604601917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffba0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff000000000000000000000000000000000000000000000000811692600881106135dc575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613642575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b359060208110613682575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b919091357fffff000000000000000000000000000000000000000000000000000000000000811692600281106136e3575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b6040519060c0820182811067ffffffffffffffff8211176102f857604052606060a0836000815282602082015282604082015282808201528260808201520152565b604080519091906137688382610336565b6001815291601f19018260005b82811061378157505050565b60209061378c613715565b82828501015201613775565b604051906137a7602083610336565b600080835282815b8281106137bb57505050565b6020906137c6613715565b828285010152016137af565b906137db61344d565b91604f8210613ce75761382061381a6137f484846134ba565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613cb7575061385961384b61384561383f85856134de565b906135a8565b60c01c90565b67ffffffffffffffff168452565b61387d61386c61384561383f85856134ef565b67ffffffffffffffff166020850152565b6138a161389061384561383f8585613500565b67ffffffffffffffff166040850152565b6138cd6138c06138ba6138b48585613511565b9061360e565b60e01c90565b63ffffffff166060850152565b6138ed6138e06138ba6138b48585613522565b63ffffffff166080850152565b6139266138fd6138b48484613533565b7fffffffff000000000000000000000000000000000000000000000000000000001660a0850152565b6139396139338383613544565b90613674565b60c08401528160451015613ca15761396061395a61381a6137f485856134c3565b60ff1690565b9081604601838111613c8b5761397a6131e8828685613555565b60e086015283811015613c755761395a61381a6137f461399b9387866134d2565b8201916047830190848211613c60576131e88260476139bc93018786613590565b61010086015283811015613c4a576139e061395a61381a6137f460479488876134d2565b830101916001830190848211613c34576131e8826048613a0293018786613590565b61012086015283811015613c1e57613a2661395a61381a6137f460019488876134d2565b830101916001830190848211613c08576131e8826002613a4893018786613590565b6101408601526003830192848411613bf257613a84613a7d613a77613a71876001968a89613590565b906136af565b60f01c90565b61ffff1690565b0101916002830190848211613bdc576131e882613aa2928786613590565b6101608601526004830190848211613bc657613a77613a7183613ac6938887613590565b9261ffff8294168015600014613b7557505050613ae1613798565b6101808501525b6002820191838311613b5f5780613b0c613a7d613a77613a71876002968a89613590565b010191838311613b4957826131e89185613b2594613590565b6101a084015203613b335790565b635a102da160e11b600052600f60045260246000fd5b635a102da160e11b600052600e60045260246000fd5b635a102da160e11b600052600d60045260246000fd5b6002919294508190613b98613b88613757565b966101808a01978852888761515f565b9490965196613ba78698611ff5565b5201010114613ae857635a102da160e11b600052600c60045260246000fd5b635a102da160e11b600052600b60045260246000fd5b635a102da160e11b600052600a60045260246000fd5b635a102da160e11b600052600960045260246000fd5b635a102da160e11b600052600860045260246000fd5b635a102da160e11b600052600760045260246000fd5b635a102da160e11b600052600660045260246000fd5b635a102da160e11b600052600560045260246000fd5b635a102da160e11b6000526004805260246000fd5b635a102da160e11b600052600360045260246000fd5b635a102da160e11b600052600260045260246000fd5b635a102da160e11b600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b635a102da160e11b600052611cd26024906000600452565b15919082613d7d575b508115613d73575b8115613d1a575090565b9050613d467f85572ffb0000000000000000000000000000000000000000000000000000000082615ba9565b9081613d61575b81613d5757501590565b61195c9150615b49565b9050613d6c81615a83565b1590613d4d565b803b159150613d10565b15915038613d08565b60405190613d95602083610336565b6000808352366020840137565b60408051909190613db38382610336565b6001815291601f1901366020840137565b9060018201809211612f1f57565b91908201809211612f1f57565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114612f1f5760010190565b80548210156120025760005260206000200190600090565b8015612f1f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff168015612f1f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9594936060936000939092613e90613d86565b96845180614244575b50156142285750509051159050614217575050613eb4613d86565b613ebc613d86565b936000925b6002613f026003613ee68567ffffffffffffffff166000526004602052604060002090565b019367ffffffffffffffff166000526004602052604060002090565b01825495815495613f2b613f2688613f218b613f218b518a5190613dd2565b613dd2565b611f95565b9889976000805b8951821015613f855790613f7d600192613f62613f52610b04858f612007565b91613f5c81613ddf565b9e612007565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018b99613f32565b919395979a9294969850506000905b8851821015613fcc5790613fc4600192613f62613fb4610b04858e612007565b91613fbe81613ddf565b9d612007565b018a98613f94565b959893969991929497505060005b8281106141e4575050505060005b828110614160575b50509091929350600090815b8181106140be5750508452805160005b85518110156140b85760005b828110614029575b5060010161400c565b614036610b048286612007565b73ffffffffffffffffffffffffffffffffffffffff61405b610abb610b04868c612007565b9116146140705761406b90613ddf565b614018565b9161407d61409591613e24565b92613f6261408e610b048688612007565b9186612007565b60ff84166140a4575b38614020565b926140b0600191613e4f565b93905061409e565b50815291565b6140cb610b048289612007565b73ffffffffffffffffffffffffffffffffffffffff8116801561415657600090815b8a878210614129575b50505090600192911561410c575b505b01613ffc565b61412390613f6261411c87613ddf565b968b612007565b38614104565b61413a610abb610b04848694612007565b14614147576001016140ed565b5060019150819050388a6140f6565b5050600190614106565b614170610abb610b04838b612007565b1561417d57600101613fe8565b50909192939460005b82811061419857869594939250613ff0565b806141de6141cb6141ab60019486613e0c565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b613f626141d788613ddf565b978c612007565b01614186565b61420d82939496613f626141fd6141ab85600197613e0c565b9161420781613ddf565b99612007565b0190899291613fda565b919093614222613da2565b91613ec1565b919350915061423a93508694966157cd565b9094909290613ec1565b90975060018103614293575061428c61426b611e028861426388611ff5565b510151611246565b8461427587611ff5565b51518c60a06142838a611ff5565b510151936155bd565b9638613e99565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90604051916142d060c084610336565b60848352602083019060a03683375a9063ffffffff7f00000000000000000000000000000000000000000000000000000000000000001690818311156143495761431e600093928493612f12565b82602083519301913090f1903d9060848211614340575b6000908286523e9190565b60849150614335565b611cd27fffffffff0000000000000000000000000000000000000000000000000000000063ffffffff5a1660e01b167f2882569d00000000000000000000000000000000000000000000000000000000600052907fffffffff0000000000000000000000000000000000000000000000000000000060249216600452565b73ffffffffffffffffffffffffffffffffffffffff6001541633036143e857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614420815184613dd2565b92831561454d5760005b848110614438575050505050565b818110156145325761444d610b048286612007565b73ffffffffffffffffffffffffffffffffffffffff81168015610b2b5761447383613dc4565b8781106144855750505060010161442a565b848110156145025773ffffffffffffffffffffffffffffffffffffffff6144af610b04838a612007565b1682146144be57600101614473565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff61452d610b046145278885612f12565b89612007565b6144af565b614548610b046145428484612f12565b85612007565b61444d565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b3d156145a2573d9061458882610398565b916145966040519384610336565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161465c575b506146585782614622614577565b906106416040519283927f9fe2f95a000000000000000000000000000000000000000000000000000000008452600484016112db565b9150565b90916020823d60201161468b575b8161467760209383610336565b810103126146885750519038614614565b80fd5b3d915061466a565b92919061469f81611f51565b936146ad6040519586610336565b602085838152019160051b8101918383116100e75781905b8382106146d3575050505050565b813567ffffffffffffffff81116100e7576020916146f48784938701612e3f565b8152019101906146c5565b91908110156120025760051b0190565b3561086281611141565b6147659060a06147739495969361476d61473284612e2b565b9361474161012082018261286b565b969061475d614754610180850185612817565b97909401612e35565b973691611308565b933691614693565b92613e7d565b9193909261478082611f95565b9261478a83611f95565b94600091825b8851811015614884576000805b8a88848982851061480d575b5050505050156147bb57600101614790565b6147cb610b04611cd2928b612007565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610b0461483f926130e861483a8873ffffffffffffffffffffffffffffffffffffffff97610abb966146ff565b61470f565b91161461484e5760010161479d565b6001915061486d61486361483a838b8b6146ff565b613f62888c612007565b61487961411c87613ddf565b52388a8884896147a9565b509097965094939291909460ff811690816000985b8a518a10156149565760005b8b8782108061494d575b156149405773ffffffffffffffffffffffffffffffffffffffff6148e1610abb610b048f6130e861483a888f8f6146ff565b9116146148f6576148f190613ddf565b6148a5565b939961490660019294939b613e24565b9461492261491861483a838b8b6146ff565b613f628d8c612007565b61493561492e8c613ddf565b9b8b612007565b525b01989091614899565b5050919098600190614937565b508515156148af565b98509250939594975091508161497f575050508151810361497657509190565b80825283529190565b611cd2929161498d91612f12565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906103c38261031a565b908160209103126100e757604051906149e68261031a565b51815290565b6108629160e0614a96614a84614a0d855161010086526101008601906103ec565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152614a726080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526103ec565b60c085015184820360c08601526103ec565b9201519060e08184039101526103ec565b9060206108629281815201906149ec565b907fffffffff000000000000000000000000000000000000000000000000000000006105146020929594956040855260408501906149ec565b92939193614afd612db2565b50614b0e611e026080860151611246565b91614b1f611e026060870151611246565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909690956020878060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa968715611f2257600097614e4e575b5073ffffffffffffffffffffffffffffffffffffffff8716948515614e0a57614bde6149c1565b50614c2c825191614c0f60a0602086015195015195614bfb610379565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff8816608084015260a083015260c0820152614c5f6103b4565b60e0820152614c6d8561540f565b15614d335790614cb19260209260006040518096819582947f2cab0fb600000000000000000000000000000000000000000000000000000000845260048401614ab8565b03925af160009181614d02575b50614ccc5783614622614577565b929091925b51614cf9614cdd610389565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201529190565b614d2591925060203d602011614d2c575b614d1d8183610336565b8101906149ce565b9038614cbe565b503d614d13565b9050614d3e84615465565b15614dc657614d816000926020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301614aa7565b03925af160009181614da5575b50614d9c5783614622614577565b92909192614cd1565b614dbf91925060203d602011614d2c57614d1d8183610336565b9038614d8e565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff881660045260246000fd5b614e6891975060203d602011613151576131438183610336565b9538614bb7565b90916060828403126100e7578151614e86816112b0565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e757815191614eb583610398565b91614ec36040519384610336565b838352602084830101116100e757604092614ee491602080850191016103c9565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a08401526080614f61614f2d604084015160a060c08801526101208701906103ec565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103ec565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b818110614fea5750505061ffff90951660208301526103689291606091614fce9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101614fa0565b906000916150e3938361508a73ffffffffffffffffffffffffffffffffffffffff61506f67ffffffffffffffff60208701511667ffffffffffffffff166000526004602052604060002090565b541673ffffffffffffffffffffffffffffffffffffffff1690565b92604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601614eea565b03925af1908115611f2257600090600092615138575b50156151025750565b610641906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611235565b905061515791503d806000833e61514f8183610336565b810190614e6f565b5090386150f9565b9291615169613715565b91808210156153f95761518361381a6137f48484896134d2565b600160ff821603613cb75750602182018181116153e3576151ac6139338260018601858a613590565b8452818110156153cd5761395a61381a6137f46151ca93858a6134d2565b82019160228301908282116153b7576131e88260226151eb9301858a613590565b6020850152818110156153a15761520e61395a61381a6137f4602294868b6134d2565b83010191600183019082821161538b576131e88260236152309301858a613590565b6040850152818110156153755761525361395a61381a6137f4600194868b6134d2565b83010191600183019082821161535f576131e88260026152759301858a613590565b6060850152818110156153495761529861395a61381a6137f4600194868b6134d2565b8301016001810192828411615333576131e88460026152b99301858a613590565b6080850152600381019282841161531d576002916152e4613a7d613a77613a7188600196898e613590565b01010194818611615307576152fe926131e8928792613590565b60a08201529190565b635a102da160e11b600052601b60045260246000fd5b635a102da160e11b600052601a60045260246000fd5b635a102da160e11b600052601960045260246000fd5b635a102da160e11b600052601860045260246000fd5b635a102da160e11b600052601760045260246000fd5b635a102da160e11b600052601660045260246000fd5b635a102da160e11b600052601560045260246000fd5b635a102da160e11b600052601460045260246000fd5b635a102da160e11b600052601360045260246000fd5b635a102da160e11b600052601260045260246000fd5b635a102da160e11b600052601160045260246000fd5b635a102da160e11b600052601060045260246000fd5b6154397f940a15420000000000000000000000000000000000000000000000000000000082615ba9565b9081615453575b81615449575090565b6108629150615b49565b905061545e81615a83565b1590615440565b6154397faff2afbf0000000000000000000000000000000000000000000000000000000082615ba9565b6154397f1bfc84d00000000000000000000000000000000000000000000000000000000082615ba9565b9080601f830112156100e75781516154d081611f51565b926154de6040519485610336565b81845260208085019260051b8201019283116100e757602001905b8282106155065750505090565b60208091835161551581611141565b8152019101906154f9565b906020828203126100e757815167ffffffffffffffff81116100e75761086292016154b9565b95949060019460a09467ffffffffffffffff6155b89573ffffffffffffffffffffffffffffffffffffffff7fffffffff0000000000000000000000000000000000000000000000000000000095168b521660208a0152604089015216606087015260c0608087015260c08601906103ec565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095909291906020848060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa938415611f2257600094615721575b506156628461540f565b615680575b5050505050508051156156775790565b50610862613da2565b6000959650906156d573ffffffffffffffffffffffffffffffffffffffff92604051988997889687957f06b859ef00000000000000000000000000000000000000000000000000000000875260048701615546565b0392165afa908115611f22576000916156fe575b506156f381615c0b565b388080808080615667565b61571b91503d806000833e6157138183610336565b810190615520565b386156e9565b61573b91945060203d602011613151576131438183610336565b9238615658565b9190916080818403126100e757805167ffffffffffffffff81116100e7578361576c9183016154b9565b9260208201519067ffffffffffffffff82116100e75761578d9183016154b9565b91604082015160ff811681036100e75760609092015161086281612968565b60409067ffffffffffffffff610862949316815281602082015201906103ec565b606093600093859385936157e08261548f565b61582a575b505050906157f291615cc9565b80511580159061581e575b61581b5750505061580c613da2565b90615815613d86565b90600090565b92565b5060ff821615156157fd565b9091929450615881965083955073ffffffffffffffffffffffffffffffffffffffff6040518098819582947f1bfc84d0000000000000000000000000000000000000000000000000000000008452600484016157ac565b0392165afa908115611f2257829183948480926158f5575b50509083856158a785615c0b565b6158b081615c0b565b51908160ff8216116158c257806157e5565b7f3d9055a70000000000000000000000000000000000000000000000000000000060005260ff1660045260245260446000fd5b92955092505061591792503d8091833e61590f8183610336565b810190615742565b91939092913880615899565b6002548110156120025760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b81811061596c57505060009055565b8061597960019285613e0c565b90549060031b1c600052818401602052600060408120550161595d565b600081815260036020526040902054615a2157600254680100000000000000008110156102f857615a086159d38260018594016002556002613e0c565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615a7c57805490680100000000000000008210156102f85782615a656159d3846001809601855584613e0c565b905580549260005201602052604060002055600190565b5050600090565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252615ae3604483610336565b6179185a10615b1f576020926000925191617530fa6000513d82615b13575b5081615b0c575090565b9050151590565b60201115915038615b02565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252615ae3604483610336565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252615ae3604483610336565b80519060005b828110615c1d57505050565b60018101808211612f1f575b838110615c395750600101615c11565b73ffffffffffffffffffffffffffffffffffffffff615c588385612007565b5116615c6a610abb610b048487612007565b14615c7757600101615c29565b611cd2615c87610b048486612007565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b7fffffffff00000000000000000000000000000000000000000000000000000000811615615da957615cfa81615dad565b601082811c9082901c167dffff0000000000000000000000000000000000000000000000000000000016615da95761ffff8260e01c168015908115615d98575b50615d43575050565b7fdf63778f000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000009081166004521660245260446000fd5b905061ffff8260e01c161038615d3a565b5050565b7fffffffff00000000000000000000000000000000000000000000000000000000811615615ead577dffff00000000000000000000000000000000000000000000000000000000811615615ea45760ff60015b1660f082901c80615e66575b50600103615e175750565b7fc512f96c000000000000000000000000000000000000000000000000000000006000527fffffffff000000000000000000000000000000000000000000000000000000001660045260246000fd5b60005b60108110615e775750615e0c565b63ffffffff6001821b831616615e90575b600101615e69565b91615e9c600191613dc4565b929050615e88565b60ff6000615e00565b5056fea164736f6c634300081a000a",
}

var OffRampABI = OffRampMetaData.ABI

var OffRampBin = OffRampMetaData.Bin

func DeployOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OffRampStaticConfig) (common.Address, *types.Transaction, *OffRamp, error) {
	parsed, err := OffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffRampBin), backend, staticConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OffRamp{address: address, abi: *parsed, OffRampCaller: OffRampCaller{contract: contract}, OffRampTransactor: OffRampTransactor{contract: contract}, OffRampFilterer: OffRampFilterer{contract: contract}}, nil
}

type OffRamp struct {
	address common.Address
	abi     abi.ABI
	OffRampCaller
	OffRampTransactor
	OffRampFilterer
}

type OffRampCaller struct {
	contract *bind.BoundContract
}

type OffRampTransactor struct {
	contract *bind.BoundContract
}

type OffRampFilterer struct {
	contract *bind.BoundContract
}

type OffRampSession struct {
	Contract     *OffRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OffRampCallerSession struct {
	Contract *OffRampCaller
	CallOpts bind.CallOpts
}

type OffRampTransactorSession struct {
	Contract     *OffRampTransactor
	TransactOpts bind.TransactOpts
}

type OffRampRaw struct {
	Contract *OffRamp
}

type OffRampCallerRaw struct {
	Contract *OffRampCaller
}

type OffRampTransactorRaw struct {
	Contract *OffRampTransactor
}

func NewOffRamp(address common.Address, backend bind.ContractBackend) (*OffRamp, error) {
	abi, err := abi.JSON(strings.NewReader(OffRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OffRamp{address: address, abi: abi, OffRampCaller: OffRampCaller{contract: contract}, OffRampTransactor: OffRampTransactor{contract: contract}, OffRampFilterer: OffRampFilterer{contract: contract}}, nil
}

func NewOffRampCaller(address common.Address, caller bind.ContractCaller) (*OffRampCaller, error) {
	contract, err := bindOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampCaller{contract: contract}, nil
}

func NewOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*OffRampTransactor, error) {
	contract, err := bindOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OffRampTransactor{contract: contract}, nil
}

func NewOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*OffRampFilterer, error) {
	contract, err := bindOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OffRampFilterer{contract: contract}, nil
}

func bindOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OffRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OffRamp *OffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRamp.Contract.OffRampCaller.contract.Call(opts, result, method, params...)
}

func (_OffRamp *OffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.Contract.OffRampTransactor.contract.Transfer(opts)
}

func (_OffRamp *OffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRamp.Contract.OffRampTransactor.contract.Transact(opts, method, params...)
}

func (_OffRamp *OffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OffRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_OffRamp *OffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.Contract.contract.Transfer(opts)
}

func (_OffRamp *OffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OffRamp.Contract.contract.Transact(opts, method, params...)
}

func (_OffRamp *OffRampCaller) GetAllSourceChainConfigs(opts *bind.CallOpts) (GetAllSourceChainConfigs,

	error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getAllSourceChainConfigs")

	outstruct := new(GetAllSourceChainConfigs)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceChainSelectors = *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	outstruct.SourceChainConfigs = *abi.ConvertType(out[1], new([]OffRampSourceChainConfig)).(*[]OffRampSourceChainConfig)

	return *outstruct, err

}

func (_OffRamp *OffRampSession) GetAllSourceChainConfigs() (GetAllSourceChainConfigs,

	error) {
	return _OffRamp.Contract.GetAllSourceChainConfigs(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetAllSourceChainConfigs() (GetAllSourceChainConfigs,

	error) {
	return _OffRamp.Contract.GetAllSourceChainConfigs(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) GetCCVsForMessage(opts *bind.CallOpts, encodedMessage []byte) (GetCCVsForMessage,

	error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getCCVsForMessage", encodedMessage)

	outstruct := new(GetCCVsForMessage)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Threshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_OffRamp *OffRampSession) GetCCVsForMessage(encodedMessage []byte) (GetCCVsForMessage,

	error) {
	return _OffRamp.Contract.GetCCVsForMessage(&_OffRamp.CallOpts, encodedMessage)
}

func (_OffRamp *OffRampCallerSession) GetCCVsForMessage(encodedMessage []byte) (GetCCVsForMessage,

	error) {
	return _OffRamp.Contract.GetCCVsForMessage(&_OffRamp.CallOpts, encodedMessage)
}

func (_OffRamp *OffRampCaller) GetExecutionState(opts *bind.CallOpts, messageId [32]byte) (uint8, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getExecutionState", messageId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_OffRamp *OffRampSession) GetExecutionState(messageId [32]byte) (uint8, error) {
	return _OffRamp.Contract.GetExecutionState(&_OffRamp.CallOpts, messageId)
}

func (_OffRamp *OffRampCallerSession) GetExecutionState(messageId [32]byte) (uint8, error) {
	return _OffRamp.Contract.GetExecutionState(&_OffRamp.CallOpts, messageId)
}

func (_OffRamp *OffRampCaller) GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getSourceChainConfig", sourceChainSelector)

	if err != nil {
		return *new(OffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampSourceChainConfig)).(*OffRampSourceChainConfig)

	return out0, err

}

func (_OffRamp *OffRampSession) GetSourceChainConfig(sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	return _OffRamp.Contract.GetSourceChainConfig(&_OffRamp.CallOpts, sourceChainSelector)
}

func (_OffRamp *OffRampCallerSession) GetSourceChainConfig(sourceChainSelector uint64) (OffRampSourceChainConfig, error) {
	return _OffRamp.Contract.GetSourceChainConfig(&_OffRamp.CallOpts, sourceChainSelector)
}

func (_OffRamp *OffRampCaller) GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OffRampStaticConfig)).(*OffRampStaticConfig)

	return out0, err

}

func (_OffRamp *OffRampSession) GetStaticConfig() (OffRampStaticConfig, error) {
	return _OffRamp.Contract.GetStaticConfig(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetStaticConfig() (OffRampStaticConfig, error) {
	return _OffRamp.Contract.GetStaticConfig(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OffRamp *OffRampSession) Owner() (common.Address, error) {
	return _OffRamp.Contract.Owner(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) Owner() (common.Address, error) {
	return _OffRamp.Contract.Owner(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OffRamp *OffRampSession) TypeAndVersion() (string, error) {
	return _OffRamp.Contract.TypeAndVersion(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) TypeAndVersion() (string, error) {
	return _OffRamp.Contract.TypeAndVersion(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "acceptOwnership")
}

func (_OffRamp *OffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRamp.Contract.AcceptOwnership(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OffRamp.Contract.AcceptOwnership(&_OffRamp.TransactOpts)
}

func (_OffRamp *OffRampTransactor) ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "applySourceChainConfigUpdates", sourceChainConfigUpdates)
}

func (_OffRamp *OffRampSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRamp.Contract.ApplySourceChainConfigUpdates(&_OffRamp.TransactOpts, sourceChainConfigUpdates)
}

func (_OffRamp *OffRampTransactorSession) ApplySourceChainConfigUpdates(sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error) {
	return _OffRamp.Contract.ApplySourceChainConfigUpdates(&_OffRamp.TransactOpts, sourceChainConfigUpdates)
}

func (_OffRamp *OffRampTransactor) Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "execute", encodedMessage, ccvs, verifierResults, gasLimitOverride)
}

func (_OffRamp *OffRampSession) Execute(encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error) {
	return _OffRamp.Contract.Execute(&_OffRamp.TransactOpts, encodedMessage, ccvs, verifierResults, gasLimitOverride)
}

func (_OffRamp *OffRampTransactorSession) Execute(encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error) {
	return _OffRamp.Contract.Execute(&_OffRamp.TransactOpts, encodedMessage, ccvs, verifierResults, gasLimitOverride)
}

func (_OffRamp *OffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "executeSingleMessage", message, messageId, ccvs, verifierResults, gasLimitOverride)
}

func (_OffRamp *OffRampSession) ExecuteSingleMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteSingleMessage(&_OffRamp.TransactOpts, message, messageId, ccvs, verifierResults, gasLimitOverride)
}

func (_OffRamp *OffRampTransactorSession) ExecuteSingleMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteSingleMessage(&_OffRamp.TransactOpts, message, messageId, ccvs, verifierResults, gasLimitOverride)
}

func (_OffRamp *OffRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_OffRamp *OffRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.TransferOwnership(&_OffRamp.TransactOpts, to)
}

func (_OffRamp *OffRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OffRamp.Contract.TransferOwnership(&_OffRamp.TransactOpts, to)
}

type OffRampExecutionStateChangedIterator struct {
	Event *OffRampExecutionStateChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampExecutionStateChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampExecutionStateChanged)
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
		it.Event = new(OffRampExecutionStateChanged)
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

func (it *OffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

func (it *OffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampExecutionStateChanged struct {
	SourceChainSelector uint64
	MessageNumber       uint64
	MessageId           [32]byte
	State               uint8
	ReturnData          []byte
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, messageNumber []uint64, messageId [][32]byte) (*OffRampExecutionStateChangedIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}
	var messageNumberRule []interface{}
	for _, messageNumberItem := range messageNumber {
		messageNumberRule = append(messageNumberRule, messageNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, messageNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OffRampExecutionStateChangedIterator{contract: _OffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampExecutionStateChanged, sourceChainSelector []uint64, messageNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}
	var messageNumberRule []interface{}
	for _, messageNumberItem := range messageNumber {
		messageNumberRule = append(messageNumberRule, messageNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, messageNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampExecutionStateChanged)
				if err := _OffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseExecutionStateChanged(log types.Log) (*OffRampExecutionStateChanged, error) {
	event := new(OffRampExecutionStateChanged)
	if err := _OffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOwnershipTransferRequestedIterator struct {
	Event *OffRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOwnershipTransferRequested)
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
		it.Event = new(OffRampOwnershipTransferRequested)
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

func (it *OffRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OffRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRamp *OffRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOwnershipTransferRequestedIterator{contract: _OffRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOwnershipTransferRequested)
				if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*OffRampOwnershipTransferRequested, error) {
	event := new(OffRampOwnershipTransferRequested)
	if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampOwnershipTransferredIterator struct {
	Event *OffRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampOwnershipTransferred)
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
		it.Event = new(OffRampOwnershipTransferred)
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

func (it *OffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OffRamp *OffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OffRampOwnershipTransferredIterator{contract: _OffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampOwnershipTransferred)
				if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseOwnershipTransferred(log types.Log) (*OffRampOwnershipTransferred, error) {
	event := new(OffRampOwnershipTransferred)
	if err := _OffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampSourceChainConfigSetIterator struct {
	Event *OffRampSourceChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampSourceChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampSourceChainConfigSet)
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
		it.Event = new(OffRampSourceChainConfigSet)
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

func (it *OffRampSourceChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampSourceChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampSourceChainConfigSet struct {
	SourceChainSelector uint64
	SourceConfig        OffRampSourceChainConfigArgs
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*OffRampSourceChainConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OffRampSourceChainConfigSetIterator{contract: _OffRamp.contract, event: "SourceChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampSourceChainConfigSet)
				if err := _OffRamp.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseSourceChainConfigSet(log types.Log) (*OffRampSourceChainConfigSet, error) {
	event := new(OffRampSourceChainConfigSet)
	if err := _OffRamp.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OffRampStaticConfigSetIterator struct {
	Event *OffRampStaticConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampStaticConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampStaticConfigSet)
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
		it.Event = new(OffRampStaticConfigSet)
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

func (it *OffRampStaticConfigSetIterator) Error() error {
	return it.fail
}

func (it *OffRampStaticConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampStaticConfigSet struct {
	StaticConfig OffRampStaticConfig
	Raw          types.Log
}

func (_OffRamp *OffRampFilterer) FilterStaticConfigSet(opts *bind.FilterOpts) (*OffRampStaticConfigSetIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return &OffRampStaticConfigSetIterator{contract: _OffRamp.contract, event: "StaticConfigSet", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampStaticConfigSet) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "StaticConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampStaticConfigSet)
				if err := _OffRamp.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseStaticConfigSet(log types.Log) (*OffRampStaticConfigSet, error) {
	event := new(OffRampStaticConfigSet)
	if err := _OffRamp.contract.UnpackLog(event, "StaticConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetAllSourceChainConfigs struct {
	SourceChainSelectors []uint64
	SourceChainConfigs   []OffRampSourceChainConfig
}
type GetCCVsForMessage struct {
	RequiredCCVs []common.Address
	OptionalCCVs []common.Address
	Threshold    uint8
}

func (OffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2")
}

func (OffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OffRampSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e")
}

func (OffRampStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x6db4162777b6c980e778bb05a3d9e050f3792b091287ff0d4f3d51bdcd7427db")
}

func (_OffRamp *OffRamp) Address() common.Address {
	return _OffRamp.address
}

type OffRampInterface interface {
	GetAllSourceChainConfigs(opts *bind.CallOpts) (GetAllSourceChainConfigs,

		error)

	GetCCVsForMessage(opts *bind.CallOpts, encodedMessage []byte) (GetCCVsForMessage,

		error)

	GetExecutionState(opts *bind.CallOpts, messageId [32]byte) (uint8, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, messageNumber []uint64, messageId [][32]byte) (*OffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampExecutionStateChanged, sourceChainSelector []uint64, messageNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseExecutionStateChanged(log types.Log) (*OffRampExecutionStateChanged, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OffRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OffRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OffRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OffRampOwnershipTransferred, error)

	FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*OffRampSourceChainConfigSetIterator, error)

	WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error)

	ParseSourceChainConfigSet(log types.Log) (*OffRampSourceChainConfigSet, error)

	FilterStaticConfigSet(opts *bind.FilterOpts) (*OffRampStaticConfigSetIterator, error)

	WatchStaticConfigSet(opts *bind.WatchOpts, sink chan<- *OffRampStaticConfigSet) (event.Subscription, error)

	ParseStaticConfigSet(log types.Log) (*OffRampStaticConfigSet, error)

	Address() common.Address
}
