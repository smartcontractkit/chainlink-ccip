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
	Finality            uint16
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxGasBufferToUpdateStateUpdated\",\"inputs\":[{\"name\":\"oldMaxGasBufferToUpdateState\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newMaxGasBufferToUpdateState\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"GasCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidGasLimitOverride\",\"inputs\":[{\"name\":\"messageGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120604052346102af57604051601f615e0138819003918201601f19168301916001600160401b038311848410176102b45780849260a0946040528339810103126102af576040519060009060a083016001600160401b0381118482101761029b5760405280516001600160401b0381168103610297578352602081015161ffff8116810361029757602084019081526040820151906001600160a01b038216820361029357604085019182526060830151926001600160a01b038416840361028f5760608601938452608001519363ffffffff8516850361028c5760808601948552331561027d57600180546001600160a01b0319163317905582516001600160a01b031615801561026b575b61025c5785516001600160401b03161561024d5761ffff8251161561023e5763ffffffff8551161561023e5785516001600160401b03908116608090815284516001600160a01b0390811660a09081528751821660c052855161ffff90811660e052895163ffffffff90811661010052604080518d519097168752885190921660208701528851841691860191909152885190921660608501528851909116918301919091527f6db4162777b6c980e778bb05a3d9e050f3792b091287ff0d4f3d51bdcd7427db91a1604051615b3690816102cb823960805181818161013e015261192f015260a0518181816101a10152611856015260c0518181816101dd01528181614a7701526154d0015260e0518181816101650152614fab01526101005181818161020901526142ad0152f35b632855a4d960e11b8152600490fd5b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b03161561010e565b639b15e16f60e01b8152600490fd5b80fd5b8480fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780633b81ab9b146100c857806349d8033e146100c35780635215505b146100be5780635643a782146100b957806361a10e59146100b457806379ba5097146100af5780638da5cb5b146100aa578063e9d68a8e146100a55763f2fde38b146100a057600080fd5b6110d6565b610f12565b610eb0565b610de5565b610d52565b610995565b61087e565b6106df565b610606565b61051b565b610411565b6100ec565b60009103126100e757565b600080fd5b346100e75760006003193601126100e7576000608060405161010d816102dc565b82815282602082015282604082015282606082015201526102a9604051610133816102dc565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016608082015260405191829182919091608063ffffffff8160a084019567ffffffffffffffff815116855261ffff602082015116602086015273ffffffffffffffffffffffffffffffffffffffff604082015116604086015273ffffffffffffffffffffffffffffffffffffffff6060820151166060860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff8211176102f857604052565b6102ad565b6101c0810190811067ffffffffffffffff8211176102f857604052565b6020810190811067ffffffffffffffff8211176102f857604052565b90601f601f19910116810190811067ffffffffffffffff8211176102f857604052565b6040519061036860c083610336565b565b6040519061036860a083610336565b6040519061036861010083610336565b60405190610368604083610336565b67ffffffffffffffff81116102f857601f01601f191660200190565b604051906103c3602083610336565b60008252565b60005b8381106103dc5750506000910152565b81810151838201526020016103cc565b90601f19601f60209361040a815180928187528780880191016103c9565b0116010190565b346100e75760006003193601126100e7576102a960408051906104348183610336565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906103ec565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104bc5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104af565b9161051460ff9161050660409497969760608752606087019061049e565b90858203602087015261049e565b9416910152565b346100e75760206003193601126100e75760043567ffffffffffffffff81116100e7576105ad61055b6105556102a9933690600401610470565b906137a4565b6105696101408201516111ac565b60601c9067ffffffffffffffff81511691610180820151906105a78161ffff60a0860151169463ffffffff60806101a0830151519201511690613cb6565b93613e34565b604093919351938493846104e8565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b63ffffffff8116036100e757565b3590610368826105ed565b346100e75760806003193601126100e75760043567ffffffffffffffff81116100e757610637903690600401610470565b9060243567ffffffffffffffff81116100e7576106589036906004016105bc565b926044359367ffffffffffffffff85116100e75761067d6106929536906004016105bc565b9390926064359561068d876105ed565b61172a565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600411156106cd57565b610694565b9060048210156106cd5752565b346100e75760206003193601126100e7576004356000526006602052602060ff6040600020541661071360405180926106d2565bf35b9080602083519182815201916020808360051b8301019401926000915b83831061074157505050505090565b909192939460208061075f83601f19866001960301875289516103ec565b97019301930191939290610732565b6107d99173ffffffffffffffffffffffffffffffffffffffff825116815260208201511515602082015260806107c86107b6604085015160a0604086015260a0850190610715565b6060850151848203606086015261049e565b92015190608081840391015261049e565b90565b6040810160408252825180915260206060830193019060005b81811061085e575050506020818303910152815180825260208201916020808360051b8301019401926000915b83831061083157505050505090565b909192939460208061084f83601f198660019603018752895161076e565b97019301930191939290610822565b825167ffffffffffffffff168552602094850194909201916001016107f5565b346100e75760006003193601126100e75760025461089b81611e92565b906108a96040519283610336565b808252601f196108b882611e92565b0160005b81811061097e5750506108ce81611ed6565b9060005b8181106108ea5750506102a9604051928392836107dc565b806109226109096108fc600194615783565b67ffffffffffffffff1690565b6109138387611f48565b9067ffffffffffffffff169052565b61096261095d6109436109358488611f48565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b612010565b61096c8287611f48565b526109778186611f48565b50016108d2565b602090610989611eaa565b828287010152016108bc565b346100e75760206003193601126100e75760043567ffffffffffffffff81116100e7576109c69036906004016105bc565b906109cf61438d565b6000905b8282106109dc57005b6109ef6109ea8385846121b5565b6122e3565b6020810191610a096108fc845167ffffffffffffffff1690565b15610d2857610a4b610a32610a32845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610d1b575b610aa25760808201949060005b86518051821015610acc57610a32610a7b83610a9593611f48565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610aa257600101610a60565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610b0457610a32610a7b83610af793611f48565b15610aa257600101610adc565b505095929491909394610b1a86518251906143d8565b610b2f610943835167ffffffffffffffff1690565b90610b5f610b45845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610b69876157b8565b606085019860005b8a518051821015610bc15790610b8981602093611f48565b518051928391012091158015610bb1575b610aa257610baa6001928b615887565b5001610b71565b50610bba612398565b8214610b9a565b5050976001975093610cd8610d11946003610d0495610cd07f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e996108fc979f610c1e60408e610c17610c64945160018b01612551565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610cc6610c858d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5160028501612636565b519101612636565b610cf5610cf06108fc835167ffffffffffffffff1690565b6157f6565b505167ffffffffffffffff1690565b92604051918291826126ca565b0390a201906109d3565b5060808201515115610a53565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760a06003193601126100e75760043567ffffffffffffffff81116100e7576101c060031982360301126100e7576024359060443567ffffffffffffffff81116100e757610da89036906004016105bc565b926064359367ffffffffffffffff85116100e757610dcd6106929536906004016105bc565b93909260843595610ddd876105ed565b600401612e95565b346100e75760006003193601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303610e86577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760006003193601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8116036100e757565b359061036882610ee4565b9060206107d992818152019061076e565b346100e75760206003193601126100e75767ffffffffffffffff600435610f3881610ee4565b610f40611eaa565b501660005260046020526040600020604051610f5b816102dc565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015260018201805490610f9382611e92565b91610fa16040519384610336565b80835260208301916000526020600020916000905b828210610ff5576102a986610fe460038a896040850152610fd960028201611faf565b606085015201611faf565b608082015260405191829182610f01565b6040516000855461100581611f5c565b8084529060018116908115611077575060011461103f575b506001928261103185946020940382610336565b815201940191019092610fb6565b6000878152602081209092505b8183106110615750508101602001600161101d565b600181602092548386880101520192019161104c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b840190910191506001905061101d565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b346100e75760206003193601126100e75773ffffffffffffffffffffffffffffffffffffffff600435611108816110b8565b61111061438d565b1633811461118257807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106111e4575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b801515036100e757565b908160209103126100e757516107d981611216565b6040513d6000823e3d90fd5b9060206107d99281815201906103ec565b60409073ffffffffffffffffffffffffffffffffffffffff6107d9949316815281602082015201906103ec565b92919261128b82610398565b916112996040519384610336565b8294818452818301116100e7578281602093846000960137010152565b9060048110156106cd5760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061131957505050505090565b90919293946020806113a083601f1986600196030187528951908151815260a061138f61137d61136b6113598887015160c08a88015260c08701906103ec565b604087015186820360408801526103ec565b606086015185820360608701526103ec565b608085015184820360808601526103ec565b9201519060a08184039101526103ec565b9701930193019193929061130a565b9160209082815201919060005b8181106113c95750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356113f2816110b8565b1681520194019291016113bc565b601f8260209493601f19938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b8484106114995750505050505090565b9091929394956020806114c183601f1986600196030188526114bb8b88611421565b90611400565b9801940194019294939190611489565b979694916116f1906080956116ff956116de6103689a956115638e60a0815261150760a08201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660c0820152604083015167ffffffffffffffff1660e0820152606083015163ffffffff16610100820152828c015163ffffffff1661012082015261014060a084015191019061ffff169052565b8d61016060c08301519101528d6101a06116aa61167461163e6116086115d261159e60e08901516101c06101808a01526102608901906103ec565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6089830301888a01526103ec565b6101208801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60888303016101c08901526103ec565b6101408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016101e08801526103ec565b6101608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303016102008701526103ec565b6101808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60858303016102208601526112ed565b920151906102407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60828503019101526103ec565b9260208d01528b830360408d01526113af565b9188830360608a0152611471565b94019063ffffffff169052565b8061171d6040926107d995946106d2565b81602082015201906103ec565b959192939594909461174260015460ff9060a01c1690565b611e685761178a740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b61179c61179787836137a4565b61427a565b9561183d60206117e26117ba6108fc8b5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e6357600091611e34575b50611de9576118b56118b16118a76109438a5167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b611d9e576118ce610b45885167ffffffffffffffff1690565b6118f86118b160e08a01928351602081519101209060019160005201602052604060002054151590565b611d6757506101008701516014815114801590611d36575b611cff5750602087015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603611cc85750828603611c94576101408701516014815103611c5a575063ffffffff84168015159081611c34575b50611be7579061199991369161127f565b60208151910120956119bf6119b8886000526006602052604060002090565b5460ff1690565b6119c8816106c3565b8015908115611bd3575b5015611b6157611a6d92611a72959492611a5f92611a286119fd8b6000526006602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957f61a10e590000000000000000000000000000000000000000000000000000000060208801528b8b602489016114d1565b03601f198101835282610336565b614286565b9015611b2f577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b611ac384611abe886000526006602052604060002090565b6112b6565b611aff611aed6040611add885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b9183604051948594169716958361170c565b0390a46103687fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff600392611aa6565b611bcf8787611b8d6040611b7d835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150611be0816106c3565b14386119d2565b611bcf84611bfc60808a015163ffffffff1690565b7fdf2964df0000000000000000000000000000000000000000000000000000000060005263ffffffff90811660045216602452604490565b9050611c53611c4a60808a015163ffffffff1690565b63ffffffff1690565b1138611988565b611c90906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611241565b0390fd5b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004869052602483905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b611c90906040519182917f55216e310000000000000000000000000000000000000000000000000000000083523060048401611252565b50611d49611d43826111ac565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff16301415611910565b611c9090516040519182917fa50bd14700000000000000000000000000000000000000000000000000000000835260048301611241565b611bcf611db3885167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611bcf611dfe885167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611e56915060203d602011611e5c575b611e4e8183610336565b810190611220565b38611887565b503d611e44565b611235565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116102f85760051b60200190565b60405190611eb7826102dc565b6060608083600081526000602082015282604082015282808201520152565b90611ee082611e92565b611eed6040519182610336565b828152601f19611efd8294611e92565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805115611f435760200190565b611f07565b8051821015611f435760209160051b010190565b90600182811c92168015611fa5575b6020831014611f7657565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611f6b565b906040519182815491828252602082019060005260206000209260005b818110611fe157505061036892500383610336565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201611fcc565b9060405161201d816102dc565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c16151560208301526001810180549061205782611e92565b916120656040519384610336565b80835260208301916000526020600020916000905b8282106120af575050505060036080926120aa92604086015261209f60028201611faf565b606086015201611faf565b910152565b604051600085546120bf81611f5c565b808452906001811690811561213157506001146120f9575b50600192826120eb85946020940382610336565b81520194019101909261207a565b6000878152602081209092505b81831061211b575050810160200160016120d7565b6001816020925483868801015201920191612106565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506120d7565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b9015611f4357806107d991612172565b90821015611f43576107d99160051b810190612172565b3590610368826110b8565b359061036882611216565b9080601f830112156100e7578160206107d99335910161127f565b9080601f830112156100e757813561221481611e92565b926122226040519485610336565b81845260208085019260051b820101918383116100e75760208201905b83821061224e57505050505090565b813567ffffffffffffffff81116100e757602091612271878480948801016121e2565b81520191019061223f565b9080601f830112156100e757813561229381611e92565b926122a16040519485610336565b81845260208085019260051b8201019283116100e757602001905b8282106122c95750505090565b6020809183356122d8816110b8565b8152019101906122bc565b60c0813603126100e7576122f5610359565b906122ff816121cc565b825261230d60208201610ef6565b602083015261231e604082016121d7565b6040830152606081013567ffffffffffffffff81116100e75761234490369083016121fd565b6060830152608081013567ffffffffffffffff81116100e75761236a903690830161227c565b608083015260a08101359067ffffffffffffffff82116100e7576123909136910161227c565b60a082015290565b604051602081019060008252602081526123b3604082610336565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106123f3575050565b600081556001016123e8565b9190601f811161240e57505050565b610368926000526020600020906020601f840160051c8301931061243a575b601f0160051c01906123e8565b909150819061242d565b919091825167ffffffffffffffff81116102f85761246c816124668454611f5c565b846123ff565b6020601f82116001146124ca5781906124bb9394956000926124bf575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612489565b601f198216906124df84600052602060002090565b9160005b81811061253957509583600195969710612502575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806124f8565b9192602060018192868b0151815501940192016124e3565b8151916801000000000000000083116102f85781548383558084106125b3575b506020612585910191600052602060002090565b6000915b8383106125965750505050565b60016020826125a783945186612444565b01920192019190612589565b8260005283602060002091820191015b8181106125d05750612571565b806125dd60019254611f5c565b806125ea575b50016125c3565b601f811183146126005750600081555b386125e3565b6126249083601f61261685600052602060002090565b920160051c820191016123e8565b600081815260208120818355556125fa565b81519167ffffffffffffffff83116102f8576801000000000000000083116102f85760209082548484558085106126ad575b500190600052602060002060005b8381106126835750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501612676565b6126c49084600052858460002091820191016123e8565b38612668565b906107d9916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a061274361272e606085015160c0608086015260e0850190610715565b6080850151601f19858303018486015261049e565b9201519060c0601f198285030191015261049e565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b9160206107d9938181520191611400565b919091357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106111e4575050565b356107d9816105ed565b356107d981610ee4565b61ffff8116036100e757565b356107d981612856565b91909160c0818403126100e757612881610359565b9281358452602082013567ffffffffffffffff81116100e757816128a69184016121e2565b6020850152604082013567ffffffffffffffff81116100e757816128cb9184016121e2565b6040850152606082013567ffffffffffffffff81116100e757816128f09184016121e2565b6060850152608082013567ffffffffffffffff81116100e757816129159184016121e2565b608085015260a082013567ffffffffffffffff81116100e75761293892016121e2565b60a0830152565b92919061294b81611e92565b936129596040519586610336565b602085838152019160051b8101918383116100e75781905b83821061297f575050505050565b813567ffffffffffffffff81116100e7576020916129a0878493870161286c565b815201910190612971565b90821015611f43576129c29160051b8101906127ac565b9091565b908160209103126100e757516107d9816110b8565b60409073ffffffffffffffffffffffffffffffffffffffff6107d995931681528160208201520191611400565b359061036882612856565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310612a8e5750505050505090565b909192939495601f1982820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e7576020612b71600193868394019081358152612b63612b58612b3d612b22612b07612af789880188611421565b60c08b89015260c0880191611400565b612b146040880188611421565b908783036040890152611400565b612b2f6060870187611421565b908683036060880152611400565b612b4a6080860186611421565b908583036080870152611400565b9260a0810190611421565b9160a0818503910152611400565b980196019493019190612a7e565b612df76107d99593949260608352612bab60608401612b9d83610ef6565b67ffffffffffffffff169052565b612bcb612bba60208301610ef6565b67ffffffffffffffff166080850152565b612beb612bda60408301610ef6565b67ffffffffffffffff1660a0850152565b612c07612bfa606083016105fb565b63ffffffff1660c0850152565b612c23612c16608083016105fb565b63ffffffff1660e0850152565b612c3e612c3260a08301612a08565b61ffff16610100850152565b60c0810135610120840152612dc6612dba612d7b612d3c612cfd612cbe612c7f612c6b60e0890189611421565b6101c06101408d01526102208c0191611400565b612c8d610100890189611421565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d0152611400565b612ccc610120880188611421565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c0152611400565b612d0b610140870187611421565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b0152611400565b612d4a610160860186611421565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a0152611400565b612d89610180850185612a13565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152612a66565b916101a0810190611421565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa085840301610200860152611400565b9360208201526040818503910152611400565b604051906040820182811067ffffffffffffffff8211176102f85760405260006020838281520152565b90612e3e82611e92565b612e4b6040519182610336565b828152601f19612e5b8294611e92565b019060005b828110612e6c57505050565b602090612e77612e0a565b82828501015201612e60565b91908203918211612e9057565b6123b9565b9391949092943033036133f5576000612eb2610180870187612758565b9050613316575b612f3b612ed6611d43612ed06101408a018a6127ac565b9061280e565b92612f0084612ee96101a08b018b6127ac565b9050612efa611c4a60808d01612842565b90613cb6565b98899160a0612f0e8b61284c565b87612f358d612f2d612f24610180830183612758565b96909201612862565b94369161293f565b91614673565b92909860005b8a5181101561312b57806020612f64610a32610a328f612fb096610a7b91611f48565b612f79612f71848a611f48565b518a8c6129ab565b91906040518096819482937fc3a7ded6000000000000000000000000000000000000000000000000000000008452600484016127fd565b03915afa918215611e63576000926130fb575b5073ffffffffffffffffffffffffffffffffffffffff82161561309e57612ff5612fed8288611f48565b51888a6129ab565b929073ffffffffffffffffffffffffffffffffffffffff82163b156100e7576000918b8373ffffffffffffffffffffffffffffffffffffffff8f613068604051998a97889687947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601612b7f565b0393165af1918215611e6357600192613083575b5001612f41565b80613092600061309893610336565b806100dc565b3861307c565b856130c689898f856130b9610a7b611c90986130bf94611f48565b95611f48565b51916129ab565b6040939193519384937f2665cea2000000000000000000000000000000000000000000000000000000008552600485016129db565b61311d91925060203d8111613124575b6131158183610336565b8101906129c6565b9038612fc3565b503d61310b565b509597935095935096505061314e613147610180840184612758565b9050612e34565b9561315d610180840184612758565b9050613204575b506131fd57610368946131cd6131798361284c565b6131ba6131c161318d6101208701876127ac565b61319b6101a08901896127ac565b9490956131a661036a565b9c8d5267ffffffffffffffff1660208d0152565b369161127f565b6040890152369161127f565b6060860152608085015263ffffffff8216156131ea575091614f13565b6131f79150608001612842565b91614f13565b5050505050565b61325e61321e613218610180860186612758565b906121a5565b61322c6101208601866127ac565b6132586132388861284c565b9261325061324860a08b01612862565b95369061286c565b92369161127f565b906149e2565b919061326989611f36565b5273ffffffffffffffffffffffffffffffffffffffff6132a3611d43612ed06132996132186101808a018a612758565b60808101906127ac565b921673ffffffffffffffffffffffffffffffffffffffff8316036132c8575b50613164565b6132fc613301926132f66132db8b611f36565b515173ffffffffffffffffffffffffffffffffffffffff1690565b9061456d565b612e83565b602061330c88611f36565b51015238806132c2565b50601461333761332d613218610180890189612758565b60608101906127ac565b9050036133e1576014613354613299613218610180890189612758565b90500361339757613392613378611d43612ed06132996132186101808b018b612758565b6132f6611d43612ed061332d6132186101808c018c612758565b612eb9565b6133ab613299613218610180880188612758565b90611c906040519283927f8d666f60000000000000000000000000000000000000000000000000000000008452600484016127fd565b6133ab61332d613218610180880188612758565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061342c826102fd565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b9015611f435790565b9060431015611f435760430190565b90821015611f43570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906023116100e75760210190600290565b906043116100e75760230190602090565b90929192836044116100e75783116100e757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff000000000000000000000000000000000000000000000000811692600881106135ae575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613614575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff0000000000000000000000000000000000000000000000000000000000008116926002811061367a575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b3590602081106136ba575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff8211176102f857604052606060a0836000815282602082015282604082015282808201528260808201520152565b6040805190919061373a8382610336565b6001815291601f19018260005b82811061375357505050565b60209061375e6136e7565b82828501015201613747565b60405190613779602083610336565b600080835282815b82811061378d57505050565b6020906137986136e7565b82828501015201613781565b906137ad61341f565b91604d8210613c9e576137f26137ec6137c6848461348c565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613c6e575061382b61381d61381761381185856134b0565b9061357a565b60c01c90565b67ffffffffffffffff168452565b61384f61383e61381761381185856134c1565b67ffffffffffffffff166020850152565b61387361386261381761381185856134d2565b67ffffffffffffffff166040850152565b61389f61389261388c61388685856134e3565b906135e0565b60e01c90565b63ffffffff166060850152565b6138bf6138b261388c61388685856134f4565b63ffffffff166080850152565b6138e96138de6138d86138d28585613505565b90613646565b60f01c90565b61ffff1660a0850152565b6138fc6138f68383613516565b906136ac565b60c08401528160431015613c585761392361391d6137ec6137c68585613495565b60ff1690565b9081604401838111613c425761393d6131ba828685613527565b60e086015283811015613c2c5761391d6137ec6137c661395e9387866134a4565b8201916045830190848211613c17576131ba82604561397f93018786613562565b61010086015283811015613c01576139a361391d6137ec6137c660459488876134a4565b830101916001830190848211613beb576131ba8260466139c593018786613562565b61012086015283811015613bd5576139e961391d6137ec6137c660019488876134a4565b830101916001830190848211613bbf576131ba826002613a0b93018786613562565b6101408601526003830192848411613ba957613a3b613a346138d86138d2876001968a89613562565b61ffff1690565b0101916002830190848211613b93576131ba82613a59928786613562565b6101608601526004830190848211613b7d576138d86138d283613a7d938887613562565b9261ffff8294168015600014613b2c57505050613a9861376a565b6101808501525b6002820191838311613b165780613ac3613a346138d86138d2876002968a89613562565b010191838311613b0057826131ba9185613adc94613562565b6101a084015203613aea5790565b635a102da160e11b600052601260045260246000fd5b635a102da160e11b600052601160045260246000fd5b635a102da160e11b600052601060045260246000fd5b6002919294508190613b4f613b3f613729565b966101808a019788528887615050565b9490965196613b5e8698611f36565b5201010114613a9f57635a102da160e11b600052600f60045260246000fd5b635a102da160e11b600052600e60045260246000fd5b635a102da160e11b600052600d60045260246000fd5b635a102da160e11b600052600c60045260246000fd5b635a102da160e11b600052600b60045260246000fd5b635a102da160e11b600052600a60045260246000fd5b635a102da160e11b600052600960045260246000fd5b635a102da160e11b600052600860045260246000fd5b635a102da160e11b6000526004805260246000fd5b635a102da160e11b600052600360045260246000fd5b635a102da160e11b600052600260045260246000fd5b635a102da160e11b600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b635a102da160e11b600052611bcf6024906000600452565b15919082613d34575b508115613d2a575b8115613cd1575090565b9050613cfd7f85572ffb0000000000000000000000000000000000000000000000000000000082615a09565b9081613d18575b81613d0e57501590565b6118b191506159a9565b9050613d23816158e3565b1590613d04565b803b159150613cc7565b15915038613cbf565b60405190613d4c602083610336565b6000808352366020840137565b60408051909190613d6a8382610336565b6001815291601f1901366020840137565b9060018201809211612e9057565b91908201809211612e9057565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114612e905760010190565b8054821015611f435760005260206000200190600090565b8015612e90577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff168015612e90577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b94929391606091600093613e46613d3d565b968351806141fb575b5050156141e7575051156141d6575050613e67613d3d565b613e6f613d3d565b6000935b6002613ebc613ea16003613e9b8a67ffffffffffffffff166000526004602052604060002090565b01611faf565b9767ffffffffffffffff166000526004602052604060002090565b0195613ed6613ece8551845190613d89565b825190613d89565b90613eeb613ee689548094613d89565b611ed6565b9788966000805b8851821015613f455790613f3d600192613f22613f12610a7b858e611f48565b91613f1c81613d96565b9d611f48565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018a98613ef2565b91939597505097909294976000905b8751821015613f8c5790613f84600192613f22613f74610a7b858d611f48565b91613f7e81613d96565b9c611f48565b018997613f54565b9750509193969092945060005b8551811015613fce5780613fc8613fb5610a7b6001948a611f48565b613f22613fc18b613d96565b9a8d611f48565b01613f99565b50929590935093909360005b828110614152575b50509091929350600090815b8181106140b05750508452805160005b85518110156140aa5760005b82811061401b575b50600101613ffe565b614028610a7b8286611f48565b73ffffffffffffffffffffffffffffffffffffffff61404d610a32610a7b868c611f48565b9116146140625761405d90613d96565b61400a565b9161406f61408791613ddb565b92613f22614080610a7b8688611f48565b9186611f48565b60ff8416614096575b38614012565b926140a2600191613e06565b939050614090565b50815291565b6140bd610a7b8289611f48565b73ffffffffffffffffffffffffffffffffffffffff8116801561414857600090815b8a87821061411b575b5050509060019291156140fe575b505b01613fee565b61411590613f2261410e87613d96565b968b611f48565b386140f6565b61412c610a32610a7b848694611f48565b14614139576001016140df565b5060019150819050388a6140e8565b50506001906140f8565b614162610a32610a7b838b611f48565b1561416f57600101613fda565b50909192939460005b82811061418a57869594939250613fe2565b806141d06141bd61419d60019486613dc3565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b613f226141c988613d96565b978c611f48565b01614178565b6141e1939193613d59565b91613e73565b9150506141f5915084615647565b93613e73565b909197506001810361424d575061424590614224611d438661421c87611f36565b5101516111ac565b9061422e85611f36565b51518a60a061423c88611f36565b51015193615466565b953880613e4f565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61428261341f565b5090565b906040519161429660c084610336565b60848352602083019060a03683375a9063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016908183111561430f576142e4600093928493612e83565b82602083519301913090f1903d9060848211614306575b6000908286523e9190565b608491506142fb565b611bcf7fffffffff0000000000000000000000000000000000000000000000000000000063ffffffff5a1660e01b167f2882569d00000000000000000000000000000000000000000000000000000000600052907fffffffff0000000000000000000000000000000000000000000000000000000060249216600452565b73ffffffffffffffffffffffffffffffffffffffff6001541633036143ae57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916143e6815184613d89565b9283156145135760005b8481106143fe575050505050565b818110156144f857614413610a7b8286611f48565b73ffffffffffffffffffffffffffffffffffffffff81168015610aa25761443983613d7b565b87811061444b575050506001016143f0565b848110156144c85773ffffffffffffffffffffffffffffffffffffffff614475610a7b838a611f48565b16821461448457600101614439565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6144f3610a7b6144ed8885612e83565b89611f48565b614475565b61450e610a7b6145088484612e83565b85611f48565b614413565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b3d15614568573d9061454e82610398565b9161455c6040519384610336565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181614622575b5061461e57826145e861453d565b90611c906040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401611252565b9150565b90916020823d602011614651575b8161463d60209383610336565b8101031261464e57505190386145da565b80fd5b3d9150614630565b9190811015611f435760051b0190565b356107d9816110b8565b90614682949596939291613e34565b9193909261468f82611ed6565b9261469983611ed6565b94600091825b8851811015614793576000805b8a88848982851061471c575b5050505050156146ca5760010161469f565b6146da610a7b611bcf928b611f48565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610a7b61474e926130b96147498873ffffffffffffffffffffffffffffffffffffffff97610a3296614659565b614669565b91161461475d576001016146ac565b6001915061477c614772614749838b8b614659565b613f22888c611f48565b61478861410e87613d96565b52388a8884896146b8565b509097965094939291909460ff811690816000985b8a518a10156148655760005b8b8782108061485c575b1561484f5773ffffffffffffffffffffffffffffffffffffffff6147f0610a32610a7b8f6130b9614749888f8f614659565b9116146148055761480090613d96565b6147b4565b939961481560019294939b613ddb565b94614831614827614749838b8b614659565b613f228d8c611f48565b61484461483d8c613d96565b9b8b611f48565b525b019890916147a8565b5050919098600190614846565b508515156147be565b98509250939594975091508161488e575050508151810361488557509190565b80825283529190565b611bcf929161489c91612e83565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906103c38261031a565b908160209103126100e757604051906148f58261031a565b51815290565b6107d99160e06149a561499361491c855161010086526101008601906103ec565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff1690860152606086015160608601526149816080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526103ec565b60c085015184820360c08601526103ec565b9201519060e08184039101526103ec565b9060206107d99281815201906148fb565b9061ffff6105146020929594956040855260408501906148fb565b929391936149ee612e0a565b506149ff611d4360808601516111ac565b91614a10611d4360608701516111ac565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909690956020878060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa968715611e6357600097614d3f575b5073ffffffffffffffffffffffffffffffffffffffff8716948515614cfb57614acf6148d0565b50614b1d825191614b0060a0602086015195015195614aec610379565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff8816608084015260a083015260c0820152614b506103b4565b60e0820152614b5e85615300565b15614c245790614ba29260209260006040518096819582947f489a68f2000000000000000000000000000000000000000000000000000000008452600484016149c7565b03925af160009181614bf3575b50614bbd57836145e861453d565b929091925b51614bea614bce610389565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201529190565b614c1691925060203d602011614c1d575b614c0e8183610336565b8101906148dd565b9038614baf565b503d614c04565b9050614c2f84615356565b15614cb757614c726000926020926040519485809481937f39077537000000000000000000000000000000000000000000000000000000008352600483016149b6565b03925af160009181614c96575b50614c8d57836145e861453d565b92909192614bc2565b614cb091925060203d602011614c1d57614c0e8183610336565b9038614c7f565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff881660045260246000fd5b614d5991975060203d602011613124576131158183610336565b9538614aa8565b90916060828403126100e7578151614d7781611216565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e757815191614da683610398565b91614db46040519384610336565b838352602084830101116100e757604092614dd591602080850191016103c9565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a08401526080614e52614e1e604084015160a060c08801526101208701906103ec565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103ec565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b818110614edb5750505061ffff90951660208301526103689291606091614ebf9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101614e91565b90600091614fd49383614f7b73ffffffffffffffffffffffffffffffffffffffff614f6067ffffffffffffffff60208701511667ffffffffffffffff166000526004602052604060002090565b541673ffffffffffffffffffffffffffffffffffffffff1690565b92604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601614ddb565b03925af1908115611e6357600090600092615029575b5015614ff35750565b611c90906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611241565b905061504891503d806000833e6150408183610336565b810190614d60565b509038614fea565b929161505a6136e7565b91808210156152ea576150746137ec6137c68484896134a4565b600160ff821603613c6e5750602182018181116152d45761509d6138f68260018601858a613562565b8452818110156152be5761391d6137ec6137c66150bb93858a6134a4565b82019160228301908282116152a8576131ba8260226150dc9301858a613562565b602085015281811015615292576150ff61391d6137ec6137c6602294868b6134a4565b83010191600183019082821161527c576131ba8260236151219301858a613562565b6040850152818110156152665761514461391d6137ec6137c6600194868b6134a4565b830101916001830190828211615250576131ba8260026151669301858a613562565b60608501528181101561523a5761518961391d6137ec6137c6600194868b6134a4565b8301016001810192828411615224576131ba8460026151aa9301858a613562565b6080850152600381019282841161520e576002916151d5613a346138d86138d288600196898e613562565b010101948186116151f8576151ef926131ba928792613562565b60a08201529190565b635a102da160e11b600052601e60045260246000fd5b635a102da160e11b600052601d60045260246000fd5b635a102da160e11b600052601c60045260246000fd5b635a102da160e11b600052601b60045260246000fd5b635a102da160e11b600052601a60045260246000fd5b635a102da160e11b600052601960045260246000fd5b635a102da160e11b600052601860045260246000fd5b635a102da160e11b600052601760045260246000fd5b635a102da160e11b600052601660045260246000fd5b635a102da160e11b600052601560045260246000fd5b635a102da160e11b600052601460045260246000fd5b635a102da160e11b600052601360045260246000fd5b61532a7f331710310000000000000000000000000000000000000000000000000000000082615a09565b9081615344575b8161533a575090565b6107d991506159a9565b905061534f816158e3565b1590615331565b61532a7faff2afbf0000000000000000000000000000000000000000000000000000000082615a09565b9080601f830112156100e757815161539781611e92565b926153a56040519485610336565b81845260208085019260051b8201019283116100e757602001905b8282106153cd5750505090565b6020809183516153dc816110b8565b8152019101906153c0565b906020828203126100e757815167ffffffffffffffff81116100e7576107d99201615380565b95949060019460a09467ffffffffffffffff6154619573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906103ec565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095909291906020848060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa938415611e63576000946155ca575b5061550b84615300565b615529575b5050505050508051156155205790565b506107d9613d59565b60009596509061557e73ffffffffffffffffffffffffffffffffffffffff92604051988997889687957f89720a620000000000000000000000000000000000000000000000000000000087526004870161540d565b0392165afa908115611e63576000916155a7575b5061559c81615a6b565b388080808080615510565b6155c491503d806000833e6155bc8183610336565b8101906153e7565b38615592565b6155e491945060203d602011613124576131158183610336565b9238615501565b90916060828403126100e757815167ffffffffffffffff81116100e75783615614918401615380565b92602083015167ffffffffffffffff81116100e757604091615637918501615380565b92015160ff811681036100e75790565b906156727f7909b5490000000000000000000000000000000000000000000000000000000082615a09565b80615773575b80615764575b61569c575b505061568d613d59565b90615696613d3d565b90600090565b6040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9290921660048301526000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015611e635760008092819261573e575b5061571181615a6b565b61571a83615a6b565b805115801590615732575b61572f5750615683565b92565b5060ff82161515615725565b915061575c92503d8091833e6157548183610336565b8101906155eb565b909138615707565b5061576e816159a9565b61567e565b5061577d816158e3565b15615678565b600254811015611f435760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b8181106157cc57505060009055565b806157d960019285613dc3565b90549060031b1c60005281840160205260006040812055016157bd565b60008181526003602052604090205461588157600254680100000000000000008110156102f8576158686158338260018594016002556002613dc3565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b60008281526001820160205260409020546158dc57805490680100000000000000008210156102f857826158c5615833846001809601855584613dc3565b905580549260005201602052604060002055600190565b5050600090565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252615943604483610336565b6179185a1061597f576020926000925191617530fa6000513d82615973575b508161596c575090565b9050151590565b60201115915038615962565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252615943604483610336565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252615943604483610336565b80519060005b828110615a7d57505050565b60018101808211612e90575b838110615a995750600101615a71565b73ffffffffffffffffffffffffffffffffffffffff615ab88385611f48565b5116615aca610a32610a7b8487611f48565b14615ad757600101615a89565b611bcf615ae7610a7b8486611f48565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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

func (_OffRamp *OffRampCaller) GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []OffRampSourceChainConfig, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getAllSourceChainConfigs")

	if err != nil {
		return *new([]uint64), *new([]OffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]OffRampSourceChainConfig)).(*[]OffRampSourceChainConfig)

	return out0, out1, err

}

func (_OffRamp *OffRampSession) GetAllSourceChainConfigs() ([]uint64, []OffRampSourceChainConfig, error) {
	return _OffRamp.Contract.GetAllSourceChainConfigs(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetAllSourceChainConfigs() ([]uint64, []OffRampSourceChainConfig, error) {
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

type OffRampMaxGasBufferToUpdateStateUpdatedIterator struct {
	Event *OffRampMaxGasBufferToUpdateStateUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OffRampMaxGasBufferToUpdateStateUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OffRampMaxGasBufferToUpdateStateUpdated)
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
		it.Event = new(OffRampMaxGasBufferToUpdateStateUpdated)
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

func (it *OffRampMaxGasBufferToUpdateStateUpdatedIterator) Error() error {
	return it.fail
}

func (it *OffRampMaxGasBufferToUpdateStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OffRampMaxGasBufferToUpdateStateUpdated struct {
	OldMaxGasBufferToUpdateState uint32
	NewMaxGasBufferToUpdateState uint32
	Raw                          types.Log
}

func (_OffRamp *OffRampFilterer) FilterMaxGasBufferToUpdateStateUpdated(opts *bind.FilterOpts) (*OffRampMaxGasBufferToUpdateStateUpdatedIterator, error) {

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "MaxGasBufferToUpdateStateUpdated")
	if err != nil {
		return nil, err
	}
	return &OffRampMaxGasBufferToUpdateStateUpdatedIterator{contract: _OffRamp.contract, event: "MaxGasBufferToUpdateStateUpdated", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchMaxGasBufferToUpdateStateUpdated(opts *bind.WatchOpts, sink chan<- *OffRampMaxGasBufferToUpdateStateUpdated) (event.Subscription, error) {

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "MaxGasBufferToUpdateStateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OffRampMaxGasBufferToUpdateStateUpdated)
				if err := _OffRamp.contract.UnpackLog(event, "MaxGasBufferToUpdateStateUpdated", log); err != nil {
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

func (_OffRamp *OffRampFilterer) ParseMaxGasBufferToUpdateStateUpdated(log types.Log) (*OffRampMaxGasBufferToUpdateStateUpdated, error) {
	event := new(OffRampMaxGasBufferToUpdateStateUpdated)
	if err := _OffRamp.contract.UnpackLog(event, "MaxGasBufferToUpdateStateUpdated", log); err != nil {
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

type GetCCVsForMessage struct {
	RequiredCCVs []common.Address
	OptionalCCVs []common.Address
	Threshold    uint8
}

func (OffRampExecutionStateChanged) Topic() common.Hash {
	return common.HexToHash("0x8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2")
}

func (OffRampMaxGasBufferToUpdateStateUpdated) Topic() common.Hash {
	return common.HexToHash("0x7266121371af537e246f0b727f08d4a221cdcdb38ff862bb874bc4b55a6642dd")
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
	GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []OffRampSourceChainConfig, error)

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

	FilterMaxGasBufferToUpdateStateUpdated(opts *bind.FilterOpts) (*OffRampMaxGasBufferToUpdateStateUpdatedIterator, error)

	WatchMaxGasBufferToUpdateStateUpdated(opts *bind.WatchOpts, sink chan<- *OffRampMaxGasBufferToUpdateStateUpdated) (event.Subscription, error)

	ParseMaxGasBufferToUpdateStateUpdated(log types.Log) (*OffRampMaxGasBufferToUpdateStateUpdated, error)

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
