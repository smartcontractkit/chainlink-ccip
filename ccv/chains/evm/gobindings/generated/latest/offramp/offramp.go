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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxGasBufferToUpdateStateUpdated\",\"inputs\":[{\"name\":\"oldMaxGasBufferToUpdateState\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newMaxGasBufferToUpdateState\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GasCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidGasLimitOverride\",\"inputs\":[{\"name\":\"messageGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalThreshold\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoStateProgressMade\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120604052346102af57604051601f6162ad38819003918201601f19168301916001600160401b038311848410176102b45780849260a0946040528339810103126102af576040519060009060a083016001600160401b0381118482101761029b5760405280516001600160401b0381168103610297578352602081015161ffff8116810361029757602084019081526040820151906001600160a01b038216820361029357604085019182526060830151926001600160a01b038416840361028f5760608601938452608001519363ffffffff8516850361028c5760808601948552331561027d57600180546001600160a01b0319163317905582516001600160a01b031615801561026b575b61025c5785516001600160401b03161561024d5761ffff8251161561023e5763ffffffff8551161561023e5785516001600160401b03908116608090815284516001600160a01b0390811660a09081528751821660c052855161ffff90811660e052895163ffffffff90811661010052604080518d519097168752885190921660208701528851841691860191909152885190921660608501528851909116918301919091527f6db4162777b6c980e778bb05a3d9e050f3792b091287ff0d4f3d51bdcd7427db91a1604051615fe290816102cb823960805181818161015c0152611ab9015260a0518181816101bf01526119e0015260c0518181816101fb01528181614dd60152615985015260e051818181610183015261530a015261010051818181610227015261460c0152f35b632855a4d960e11b8152600490fd5b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b03161561010e565b639b15e16f60e01b8152600490fd5b80fd5b8480fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780633b81ab9b146100c857806349d8033e146100c35780635215505b146100be5780635643a782146100b957806361a10e59146100b457806379ba5097146100af5780638da5cb5b146100aa578063e9d68a8e146100a55763f2fde38b146100a057600080fd5b61123e565b61105c565b610fdc565b610ef3565b610e24565b610a49565b610914565b610757565b610660565b610557565b61042f565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000608060405161012b816102fa565b82815282602082015282604082015282606082015201526102c7604051610151816102fa565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016608082015260405191829182919091608063ffffffff8160a084019567ffffffffffffffff815116855261ffff602082015116602086015273ffffffffffffffffffffffffffffffffffffffff604082015116604086015273ffffffffffffffffffffffffffffffffffffffff6060820151166060860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff82111761031657604052565b6102cb565b6101c0810190811067ffffffffffffffff82111761031657604052565b6020810190811067ffffffffffffffff82111761031657604052565b90601f601f19910116810190811067ffffffffffffffff82111761031657604052565b6040519061038660c083610354565b565b6040519061038660a083610354565b6040519061038661010083610354565b60405190610386604083610354565b67ffffffffffffffff811161031657601f01601f191660200190565b604051906103e1602083610354565b60008252565b60005b8381106103fa5750506000910152565b81810151838201526020016103ea565b90601f19601f602093610428815180928187528780880191016103e7565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576102c760408051906104708183610354565b601182527f4f666652616d7020322e302e302d64657600000000000000000000000000000060208301525191829160208352602083019061040a565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104f85750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104eb565b9161055060ff916105426040949796976060875260608701906104da565b9085820360208701526104da565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106076105b56105af6102c79336906004016104ac565b90613985565b6105c3610140820151611332565b60601c9067ffffffffffffffff81511691610180820151906106018161ffff60a0860151169463ffffffff60806101a0830151519201511690614027565b936141a5565b60409391935193849384610524565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b63ffffffff8116036100e757565b359061038682610647565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106af9036906004016104ac565b9060243567ffffffffffffffff81116100e7576106d0903690600401610616565b926044359367ffffffffffffffff85116100e7576106f561070a953690600401610616565b9390926064359561070587610647565b6118c7565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561074557565b61070c565b9060048210156107455752565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356000526006602052602060ff604060002054166107a9604051809261074a565bf35b9080602083519182815201916020808360051b8301019401926000915b8383106107d757505050505090565b90919293946020806107f583601f198660019603018752895161040a565b970193019301919392906107c8565b61086f9173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061085e61084c604085015160a0604086015260a08501906107ab565b606085015184820360608601526104da565b9201519060808184039101526104da565b90565b6040810160408252825180915260206060830193019060005b8181106108f4575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106108c757505050505090565b90919293946020806108e583601f1986600196030187528951610804565b970193019301919392906108b8565b825167ffffffffffffffff1685526020948501949092019160010161088b565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461094f81612073565b9061095d6040519283610354565b808252601f1961096c82612073565b0160005b818110610a32575050610982816120b7565b9060005b81811061099e5750506102c760405192839283610872565b806109d66109bd6109b0600194615c2f565b67ffffffffffffffff1690565b6109c78387612129565b9067ffffffffffffffff169052565b610a16610a116109f76109e98488612129565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b6121f1565b610a208287612129565b52610a2b8186612129565b5001610986565b602090610a3d61208b565b82828701015201610970565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610a98903690600401610616565b90610aa16146ec565b6000905b828210610aae57005b610ac1610abc838584612396565b6124c4565b6020810191610adb6109b0845167ffffffffffffffff1690565b15610dfa57610b1d610b04610b04845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610ded575b610b745760808201949060005b86518051821015610b9e57610b04610b4d83610b6793612129565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610b7457600101610b32565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610bd657610b04610b4d83610bc993612129565b15610b7457600101610bae565b505095929491909394610bec8651825190614737565b610c016109f7835167ffffffffffffffff1690565b90610c31610c17845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610c3b87615c64565b606085019860005b8a518051821015610c935790610c5b81602093612129565b518051928391012091158015610c83575b610b7457610c7c6001928b615d33565b5001610c43565b50610c8c612579565b8214610c6c565b5050976001975093610daa610de3946003610dd695610da27f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e996109b0979f610cf060408e610ce9610d36945160018b01612732565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610d98610d578d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5160028501612817565b519101612817565b610dc7610dc26109b0835167ffffffffffffffff1690565b615ca2565b505167ffffffffffffffff1690565b92604051918291826128ab565b0390a20190610aa5565b5060808201515115610b25565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e7576024359060443567ffffffffffffffff81116100e757610eb6903690600401610616565b926064359367ffffffffffffffff85116100e757610edb61070a953690600401610616565b93909260843595610eeb87610647565b600401613076565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303610fb2577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8116036100e757565b35906103868261102e565b90602061086f928181520190610804565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff6004356110a08161102e565b6110a861208b565b5016600052600460205260406000206040516110c3816102fa565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c1615156020820152600182018054906110fb82612073565b916111096040519384610354565b80835260208301916000526020600020916000905b82821061115d576102c78661114c60038a89604085015261114160028201612190565b606085015201612190565b60808201526040519182918261104b565b6040516000855461116d8161213d565b80845290600181169081156111df57506001146111a7575b506001928261119985946020940382610354565b81520194019101909261111e565b6000878152602081209092505b8183106111c957505081016020016001611185565b60018160209254838688010152019201916111b4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050611185565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561128e81611220565b6112966146ec565b1633811461130857807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116926014811061136a575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b801515036100e757565b908160209103126100e7575161086f8161139c565b6040513d6000823e3d90fd5b90602061086f92818152019061040a565b60409073ffffffffffffffffffffffffffffffffffffffff61086f9493168152816020820152019061040a565b929192611411826103b6565b9161141f6040519384610354565b8294818452818301116100e7578281602093846000960137010152565b9060048110156107455760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061149f57505050505090565b909192939460208061152683601f1986600196030187528951908151815260a06115156115036114f16114df8887015160c08a88015260c087019061040a565b6040870151868203604088015261040a565b6060860151858203606087015261040a565b6080850151848203608086015261040a565b9201519060a081840391015261040a565b97019301930191939290611490565b9160209082815201919060005b81811061154f5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561157881611220565b168152019401929101611542565b601f8260209493601f19938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b84841061161f5750505050505090565b90919293949560208061164783601f1986600196030188526116418b886115a7565b90611586565b980194019401929493919061160f565b9796949161187790608095611885956118646103869a956116e98e60a0815261168d60a08201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660c0820152604083015167ffffffffffffffff1660e0820152606083015163ffffffff16610100820152828c015163ffffffff1661012082015261014060a084015191019061ffff169052565b8d61016060c08301519101528d6101a06118306117fa6117c461178e61175861172460e08901516101c06101808a015261026089019061040a565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6089830301888a015261040a565b6101208801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60888303016101c089015261040a565b6101408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016101e088015261040a565b6101608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608683030161020087015261040a565b6101808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6085830301610220860152611473565b920151906102407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608285030191015261040a565b9260208d01528b830360408d0152611535565b9188830360608a01526115f7565b94019063ffffffff169052565b60409061086f93928152816020820152019061040a565b806118ba60409261086f959461074a565b816020820152019061040a565b6001549596919560a01c60ff166120495761191c740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b6119268682613985565b956119c7602061196c6119446109b08b5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561204457600091612015575b50611fca57611a3f611a3b611a316109f78a5167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b611f7f57611a58610c17885167ffffffffffffffff1690565b611a82611a3b60e08a01928351602081519101209060019160005201602052604060002054151590565b611f4857506101008701516014815114801590611f17575b611ee05750602087015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603611ea95750878503611e75576101408701516014815103611e3f575063ffffffff83168015159081611e19575b50611dcc5790611b23913691611405565b6020815191012095611b49611b42886000526006602052604060002090565b5460ff1690565b94611b538661073b565b85158015611db9575b15611d4757611bf692611bfb959492611be892611bb1611b868c6000526006602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957f61a10e590000000000000000000000000000000000000000000000000000000060208801528c8c60248901611657565b03601f198101835282610354565b6145e5565b9181159081611d33575b50611cf85715611cc6577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b611c5a84611c55886000526006602052604060002090565b61143c565b611c96611c846040611c74885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b918360405194859416971695836118a9565b0390a46103867fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff600392611c3d565b5082611d2f6040519283927f10afb5b600000000000000000000000000000000000000000000000000000000845260048401611892565b0390fd5b60039150611d408161073b565b1438611c05565b611db58888611d736040611d63835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b50611dc38661073b565b60038614611b5c565b611db583611de160808a015163ffffffff1690565b7fdf2964df0000000000000000000000000000000000000000000000000000000060005263ffffffff90811660045216602452604490565b9050611e38611e2f60808a015163ffffffff1690565b63ffffffff1690565b1138611b12565b611d2f906040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352600483016113c7565b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004859052602488905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b611d2f906040519182917f55216e3100000000000000000000000000000000000000000000000000000000835230600484016113d8565b50611f2a611f2482611332565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff16301415611a9a565b611d2f90516040519182917fa50bd147000000000000000000000000000000000000000000000000000000008352600483016113c7565b611db5611f94885167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611db5611fdf885167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b612037915060203d60201161203d575b61202f8183610354565b8101906113a6565b38611a11565b503d612025565b6113bb565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116103165760051b60200190565b60405190612098826102fa565b6060608083600081526000602082015282604082015282808201520152565b906120c182612073565b6120ce6040519182610354565b828152601f196120de8294612073565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156121245760200190565b6120e8565b80518210156121245760209160051b010190565b90600182811c92168015612186575b602083101461215757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161214c565b906040519182815491828252602082019060005260206000209260005b8181106121c257505061038692500383610354565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016121ad565b906040516121fe816102fa565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c16151560208301526001810180549061223882612073565b916122466040519384610354565b80835260208301916000526020600020916000905b8282106122905750505050600360809261228b92604086015261228060028201612190565b606086015201612190565b910152565b604051600085546122a08161213d565b808452906001811690811561231257506001146122da575b50600192826122cc85946020940382610354565b81520194019101909261225b565b6000878152602081209092505b8183106122fc575050810160200160016122b8565b60018160209254838688010152019201916122e7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506122b8565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b9015612124578061086f91612353565b908210156121245761086f9160051b810190612353565b359061038682611220565b35906103868261139c565b9080601f830112156100e75781602061086f93359101611405565b9080601f830112156100e75781356123f581612073565b926124036040519485610354565b81845260208085019260051b820101918383116100e75760208201905b83821061242f57505050505090565b813567ffffffffffffffff81116100e757602091612452878480948801016123c3565b815201910190612420565b9080601f830112156100e757813561247481612073565b926124826040519485610354565b81845260208085019260051b8201019283116100e757602001905b8282106124aa5750505090565b6020809183356124b981611220565b81520191019061249d565b60c0813603126100e7576124d6610377565b906124e0816123ad565b82526124ee60208201611040565b60208301526124ff604082016123b8565b6040830152606081013567ffffffffffffffff81116100e75761252590369083016123de565b6060830152608081013567ffffffffffffffff81116100e75761254b903690830161245d565b608083015260a08101359067ffffffffffffffff82116100e7576125719136910161245d565b60a082015290565b60405160208101906000825260208152612594604082610354565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106125d4575050565b600081556001016125c9565b9190601f81116125ef57505050565b610386926000526020600020906020601f840160051c8301931061261b575b601f0160051c01906125c9565b909150819061260e565b919091825167ffffffffffffffff81116103165761264d81612647845461213d565b846125e0565b6020601f82116001146126ab57819061269c9394956000926126a0575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b01519050388061266a565b601f198216906126c084600052602060002090565b9160005b81811061271a575095836001959697106126e3575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806126d9565b9192602060018192868b0151815501940192016126c4565b815191680100000000000000008311610316578154838355808410612794575b506020612766910191600052602060002090565b6000915b8383106127775750505050565b600160208261278883945186612625565b0192019201919061276a565b8260005283602060002091820191015b8181106127b15750612752565b806127be6001925461213d565b806127cb575b50016127a4565b601f811183146127e15750600081555b386127c4565b6128059083601f6127f785600052602060002090565b920160051c820191016125c9565b600081815260208120818355556127db565b81519167ffffffffffffffff83116103165768010000000000000000831161031657602090825484845580851061288e575b500190600052602060002060005b8381106128645750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501612857565b6128a59084600052858460002091820191016125c9565b38612849565b9061086f916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a061292461290f606085015160c0608086015260e08501906107ab565b6080850151601f1985830301848601526104da565b9201519060c0601f19828503019101526104da565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b91602061086f938181520191611586565b919091357fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116926014811061136a575050565b3561086f81610647565b3561086f8161102e565b61ffff8116036100e757565b3561086f81612a37565b91909160c0818403126100e757612a62610377565b9281358452602082013567ffffffffffffffff81116100e75781612a879184016123c3565b6020850152604082013567ffffffffffffffff81116100e75781612aac9184016123c3565b6040850152606082013567ffffffffffffffff81116100e75781612ad19184016123c3565b6060850152608082013567ffffffffffffffff81116100e75781612af69184016123c3565b608085015260a082013567ffffffffffffffff81116100e757612b1992016123c3565b60a0830152565b929190612b2c81612073565b93612b3a6040519586610354565b602085838152019160051b8101918383116100e75781905b838210612b60575050505050565b813567ffffffffffffffff81116100e757602091612b818784938701612a4d565b815201910190612b52565b9082101561212457612ba39160051b81019061298d565b9091565b908160209103126100e7575161086f81611220565b60409073ffffffffffffffffffffffffffffffffffffffff61086f95931681528160208201520191611586565b359061038682612a37565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310612c6f5750505050505090565b909192939495601f1982820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e7576020612d52600193868394019081358152612d44612d39612d1e612d03612ce8612cd8898801886115a7565b60c08b89015260c0880191611586565b612cf560408801886115a7565b908783036040890152611586565b612d1060608701876115a7565b908683036060880152611586565b612d2b60808601866115a7565b908583036080870152611586565b9260a08101906115a7565b9160a0818503910152611586565b980196019493019190612c5f565b612fd861086f9593949260608352612d8c60608401612d7e83611040565b67ffffffffffffffff169052565b612dac612d9b60208301611040565b67ffffffffffffffff166080850152565b612dcc612dbb60408301611040565b67ffffffffffffffff1660a0850152565b612de8612ddb60608301610655565b63ffffffff1660c0850152565b612e04612df760808301610655565b63ffffffff1660e0850152565b612e1f612e1360a08301612be9565b61ffff16610100850152565b60c0810135610120840152612fa7612f9b612f5c612f1d612ede612e9f612e60612e4c60e08901896115a7565b6101c06101408d01526102208c0191611586565b612e6e6101008901896115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d0152611586565b612ead6101208801886115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c0152611586565b612eec6101408701876115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b0152611586565b612f2b6101608601866115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a0152611586565b612f6a610180850185612bf4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152612c47565b916101a08101906115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa085840301610200860152611586565b9360208201526040818503910152611586565b604051906040820182811067ffffffffffffffff8211176103165760405260006020838281520152565b9061301f82612073565b61302c6040519182610354565b828152601f1961303c8294612073565b019060005b82811061304d57505050565b602090613058612feb565b82828501015201613041565b9190820391821161307157565b61259a565b9391949092943033036135d6576000613093610180870187612939565b90506134f7575b61311c6130b7611f246130b16101408a018a61298d565b906129ef565b926130e1846130ca6101a08b018b61298d565b90506130db611e2f60808d01612a23565b90614027565b98899160a06130ef8b612a2d565b876131168d61310e613105610180830183612939565b96909201612a43565b943691612b20565b916149d2565b92909860005b8a5181101561330c57806020613145610b04610b048f61319196610b4d91612129565b61315a613152848a612129565b518a8c612b8c565b91906040518096819482937fc3a7ded6000000000000000000000000000000000000000000000000000000008452600484016129de565b03915afa918215612044576000926132dc575b5073ffffffffffffffffffffffffffffffffffffffff82161561327f576131d66131ce8288612129565b51888a612b8c565b929073ffffffffffffffffffffffffffffffffffffffff82163b156100e7576000918b8373ffffffffffffffffffffffffffffffffffffffff8f613249604051998a97889687947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601612d60565b0393165af191821561204457600192613264575b5001613122565b80613273600061327993610354565b806100dc565b3861325d565b856132a789898f8561329a610b4d611d2f986132a094612129565b95612129565b5191612b8c565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501612bbc565b6132fe91925060203d8111613305575b6132f68183610354565b810190612ba7565b90386131a4565b503d6132ec565b509597935095935096505061332f613328610180840184612939565b9050613015565b9561333e610180840184612939565b90506133e5575b506133de57610386946133ae61335a83612a2d565b61339b6133a261336e61012087018761298d565b61337c6101a089018961298d565b949095613387610388565b9c8d5267ffffffffffffffff1660208d0152565b3691611405565b60408901523691611405565b6060860152608085015263ffffffff8216156133cb575091615272565b6133d89150608001612a23565b91615272565b5050505050565b61343f6133ff6133f9610180860186612939565b90612386565b61340d61012086018661298d565b61343961341988612a2d565b9261343161342960a08b01612a43565b953690612a4d565b923691611405565b90614d41565b919061344a89612117565b5273ffffffffffffffffffffffffffffffffffffffff613484611f246130b161347a6133f96101808a018a612939565b608081019061298d565b921673ffffffffffffffffffffffffffffffffffffffff8316036134a9575b50613345565b6134dd6134e2926134d76134bc8b612117565b515173ffffffffffffffffffffffffffffffffffffffff1690565b906148cc565b613064565b60206134ed88612117565b51015238806134a3565b50601461351861350e6133f9610180890189612939565b606081019061298d565b9050036135c257601461353561347a6133f9610180890189612939565b90500361357857613573613559611f246130b161347a6133f96101808b018b612939565b6134d7611f246130b161350e6133f96101808c018c612939565b61309a565b61358c61347a6133f9610180880188612939565b90611d2f6040519283927f8d666f60000000000000000000000000000000000000000000000000000000008452600484016129de565b61358c61350e6133f9610180880188612939565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061360d8261031b565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b90156121245790565b90604310156121245760430190565b90821015612124570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906023116100e75760210190600290565b906043116100e75760230190602090565b90929192836044116100e75783116100e757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff0000000000000000000000000000000000000000000000008116926008811061378f575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106137f5575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff0000000000000000000000000000000000000000000000000000000000008116926002811061385b575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b35906020811061389b575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff82111761031657604052606060a0836000815282602082015282604082015282808201528260808201520152565b6040805190919061391b8382610354565b6001815291601f19018260005b82811061393457505050565b60209061393f6138c8565b82828501015201613928565b6040519061395a602083610354565b600080835282815b82811061396e57505050565b6020906139796138c8565b82828501015201613962565b9061398e613600565b91604d8210613ff6576139d36139cd6139a7848461366d565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613fc65750613a0c6139fe6139f86139f28585613691565b9061375b565b60c01c90565b67ffffffffffffffff168452565b613a30613a1f6139f86139f285856136a2565b67ffffffffffffffff166020850152565b613a54613a436139f86139f285856136b3565b67ffffffffffffffff166040850152565b613a80613a73613a6d613a6785856136c4565b906137c1565b60e01c90565b63ffffffff166060850152565b613aa0613a93613a6d613a6785856136d5565b63ffffffff166080850152565b613aca613abf613ab9613ab385856136e6565b90613827565b60f01c90565b61ffff1660a0850152565b613add613ad783836136f7565b9061388d565b60c08401528160431015613f9757613b04613afe6139cd6139a78585613676565b60ff1690565b9081604401838111613f6857613b1e61339b828685613708565b60e086015283811015613f3957613afe6139cd6139a7613b3f938786613685565b8201916045830190848211613f0b5761339b826045613b6093018786613743565b61010086015283811015613edc57613b84613afe6139cd6139a76045948887613685565b830101916001830190848211613ead5761339b826046613ba693018786613743565b61012086015283811015613e7e57613bca613afe6139cd6139a76001948887613685565b830101916001830190848211613e4f5761339b826002613bec93018786613743565b6101408601526003830192848411613e2057613c1c613c15613ab9613ab3876001968a89613743565b61ffff1690565b0101916002830190848211613df15761339b82613c3a928786613743565b6101608601526004830190848211613dc257613ab9613ab383613c5e938887613743565b9261ffff8294168015600014613d5857505050613c7961394b565b6101808501525b6002820191838311613d295780613ca4613c15613ab9613ab3876002968a89613743565b010191838311613cfa578261339b9185613cbd94613743565b6101a084015203613ccb5790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b6002919294508190613d7b613d6b61390a565b966101808a0197885288876153af565b9490965196613d8a8698612117565b5201010114613c80577fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600860045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000006000526004805260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052611db56024906000600452565b159190826140a5575b50811561409b575b8115614042575090565b905061406e7f85572ffb0000000000000000000000000000000000000000000000000000000082615eb5565b9081614089575b8161407f57501590565b611a3b9150615e55565b905061409481615d8f565b1590614075565b803b159150614038565b15915038614030565b604051906140bd602083610354565b6000808352366020840137565b604080519091906140db8382610354565b6001815291601f1901366020840137565b906001820180921161307157565b9190820180921161307157565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146130715760010190565b80548210156121245760005260206000200190600090565b8015613071577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff168015613071577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b949390916060926000926141b76140ae565b95835180614566575b50501561454c5750511561453b5750506141d86140ae565b6141e06140ae565b936000925b6002614226600361420a8567ffffffffffffffff166000526004602052604060002090565b019367ffffffffffffffff166000526004602052604060002090565b0182549581549561424f61424a886142458b6142458b518a51906140fa565b6140fa565b6120b7565b9889976000805b89518210156142a957906142a1600192614286614276610b4d858f612129565b9161428081614107565b9e612129565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018b99614256565b919395979a9294969850506000905b88518210156142f057906142e86001926142866142d8610b4d858e612129565b916142e281614107565b9d612129565b018a986142b8565b959893969991929497505060005b828110614508575050505060005b828110614484575b50509091929350600090815b8181106143e25750508452805160005b85518110156143dc5760005b82811061434d575b50600101614330565b61435a610b4d8286612129565b73ffffffffffffffffffffffffffffffffffffffff61437f610b04610b4d868c612129565b9116146143945761438f90614107565b61433c565b916143a16143b99161414c565b926142866143b2610b4d8688612129565b9186612129565b60ff84166143c8575b38614344565b926143d4600191614177565b9390506143c2565b50815291565b6143ef610b4d8289612129565b73ffffffffffffffffffffffffffffffffffffffff8116801561447a57600090815b8a87821061444d575b505050906001929115614430575b505b01614320565b6144479061428661444087614107565b968b612129565b38614428565b61445e610b04610b4d848694612129565b1461446b57600101614411565b5060019150819050388a61441a565b505060019061442a565b614494610b04610b4d838b612129565b156144a15760010161430c565b50909192939460005b8281106144bc57869594939250614314565b806145026144ef6144cf60019486614134565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b6142866144fb88614107565b978c612129565b016144aa565b614531829394966142866145216144cf85600197614134565b9161452b81614107565b99612129565b01908992916142fe565b9190936145466140ca565b916141e5565b91505061455c9150849294615afc565b90949092906141e5565b90919650600181036145b857506145b09061458f611f248761458787612117565b510151611332565b9061459985612117565b51518a60a06145a788612117565b5101519361591b565b9438806141c0565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90604051916145f560c084610354565b60848352602083019060a03683375a9063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016908183111561466e57614643600093928493613064565b82602083519301913090f1903d9060848211614665575b6000908286523e9190565b6084915061465a565b611db57fffffffff0000000000000000000000000000000000000000000000000000000063ffffffff5a1660e01b167f2882569d00000000000000000000000000000000000000000000000000000000600052907fffffffff0000000000000000000000000000000000000000000000000000000060249216600452565b73ffffffffffffffffffffffffffffffffffffffff60015416330361470d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916147458151846140fa565b9283156148725760005b84811061475d575050505050565b8181101561485757614772610b4d8286612129565b73ffffffffffffffffffffffffffffffffffffffff81168015610b7457614798836140ec565b8781106147aa5750505060010161474f565b848110156148275773ffffffffffffffffffffffffffffffffffffffff6147d4610b4d838a612129565b1682146147e357600101614798565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614852610b4d61484c8885613064565b89612129565b6147d4565b61486d610b4d6148678484613064565b85612129565b614772565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b3d156148c7573d906148ad826103b6565b916148bb6040519384610354565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181614981575b5061497d578261494761489c565b90611d2f6040519283927f9fe2f95a000000000000000000000000000000000000000000000000000000008452600484016113d8565b9150565b90916020823d6020116149b0575b8161499c60209383610354565b810103126149ad5750519038614939565b80fd5b3d915061498f565b91908110156121245760051b0190565b3561086f81611220565b906149e19495969392916141a5565b919390926149ee826120b7565b926149f8836120b7565b94600091825b8851811015614af2576000805b8a888489828510614a7b575b505050505015614a29576001016149fe565b614a39610b4d611db5928b612129565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610b4d614aad9261329a614aa88873ffffffffffffffffffffffffffffffffffffffff97610b04966149b8565b6149c8565b911614614abc57600101614a0b565b60019150614adb614ad1614aa8838b8b6149b8565b614286888c612129565b614ae761444087614107565b52388a888489614a17565b509097965094939291909460ff811690816000985b8a518a1015614bc45760005b8b87821080614bbb575b15614bae5773ffffffffffffffffffffffffffffffffffffffff614b4f610b04610b4d8f61329a614aa8888f8f6149b8565b911614614b6457614b5f90614107565b614b13565b9399614b7460019294939b61414c565b94614b90614b86614aa8838b8b6149b8565b6142868d8c612129565b614ba3614b9c8c614107565b9b8b612129565b525b01989091614b07565b5050919098600190614ba5565b50851515614b1d565b985092509395949750915081614bed5750505081518103614be457509190565b80825283529190565b611db59291614bfb91613064565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906103e182610338565b908160209103126100e75760405190614c5482610338565b51815290565b61086f9160e0614d04614cf2614c7b8551610100865261010086019061040a565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152614ce06080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a087015261040a565b60c085015184820360c086015261040a565b9201519060e081840391015261040a565b90602061086f928181520190614c5a565b9061ffff610550602092959495604085526040850190614c5a565b92939193614d4d612feb565b50614d5e611f246080860151611332565b91614d6f611f246060870151611332565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909690956020878060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9687156120445760009761509e575b5073ffffffffffffffffffffffffffffffffffffffff871694851561505a57614e2e614c2f565b50614e7c825191614e5f60a0602086015195015195614e4b610397565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff8816608084015260a083015260c0820152614eaf6103d2565b60e0820152614ebd8561578b565b15614f835790614f019260209260006040518096819582947f489a68f200000000000000000000000000000000000000000000000000000000845260048401614d26565b03925af160009181614f52575b50614f1c578361494761489c565b929091925b51614f49614f2d6103a7565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201529190565b614f7591925060203d602011614f7c575b614f6d8183610354565b810190614c3c565b9038614f0e565b503d614f63565b9050614f8e846157e1565b1561501657614fd16000926020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301614d15565b03925af160009181614ff5575b50614fec578361494761489c565b92909192614f21565b61500f91925060203d602011614f7c57614f6d8183610354565b9038614fde565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff881660045260246000fd5b6150b891975060203d602011613305576132f68183610354565b9538614e07565b90916060828403126100e75781516150d68161139c565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e757815191615105836103b6565b916151136040519384610354565b838352602084830101116100e75760409261513491602080850191016103e7565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806151b161517d604084015160a060c088015261012087019061040a565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e088015261040a565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b81811061523a5750505061ffff9095166020830152610386929160609161521e9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff16855260209081015181860152604090940193909201916001016151f0565b9060009161533393836152da73ffffffffffffffffffffffffffffffffffffffff6152bf67ffffffffffffffff60208701511667ffffffffffffffff166000526004602052604060002090565b541673ffffffffffffffffffffffffffffffffffffffff1690565b92604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f0000000000000000000000000000000000000000000000000000000000000000906004860161513a565b03925af190811561204457600090600092615388575b50156153525750565b611d2f906040519182917f0a8d6e8c000000000000000000000000000000000000000000000000000000008352600483016113c7565b90506153a791503d806000833e61539f8183610354565b8101906150bf565b509038615349565b92916153b96138c8565b918082101561575c576153d36139cd6139a7848489613685565b600160ff821603613fc657506021820181811161572d576153fc613ad78260018601858a613743565b8452818110156156fe57613afe6139cd6139a761541a93858a613685565b82019160228301908282116156cf5761339b82602261543b9301858a613743565b6020850152818110156156a05761545e613afe6139cd6139a7602294868b613685565b8301019160018301908282116156715761339b8260236154809301858a613743565b604085015281811015615642576154a3613afe6139cd6139a7600194868b613685565b8301019160018301908282116156135761339b8260026154c59301858a613743565b6060850152818110156155e4576154e8613afe6139cd6139a7600194868b613685565b83010160018101928284116155b55761339b8460026155099301858a613743565b6080850152600381019282841161558657600291615534613c15613ab9613ab388600196898e613743565b010101948186116155575761554e9261339b928792613743565b60a08201529190565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b6157b57f331710310000000000000000000000000000000000000000000000000000000082615eb5565b90816157cf575b816157c5575090565b61086f9150615e55565b90506157da81615d8f565b15906157bc565b6157b57faff2afbf0000000000000000000000000000000000000000000000000000000082615eb5565b6157b57f7909b5490000000000000000000000000000000000000000000000000000000082615eb5565b9080601f830112156100e757815161584c81612073565b9261585a6040519485610354565b81845260208085019260051b8201019283116100e757602001905b8282106158825750505090565b60208091835161589181611220565b815201910190615875565b906020828203126100e757815167ffffffffffffffff81116100e75761086f9201615835565b95949060019460a09467ffffffffffffffff6159169573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c086019061040a565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095909291906020848060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa93841561204457600094615a7f575b506159c08461578b565b6159de575b5050505050508051156159d55790565b5061086f6140ca565b600095965090615a3373ffffffffffffffffffffffffffffffffffffffff92604051988997889687957f89720a62000000000000000000000000000000000000000000000000000000008752600487016158c2565b0392165afa90811561204457600091615a5c575b50615a5181615f17565b3880808080806159c5565b615a7991503d806000833e615a718183610354565b81019061589c565b38615a47565b615a9991945060203d602011613305576132f68183610354565b92386159b6565b90916060828403126100e757815167ffffffffffffffff81116100e75783615ac9918401615835565b92602083015167ffffffffffffffff81116100e757604091615aec918501615835565b92015160ff811681036100e75790565b90615b068161580b565b615b24575b5050615b156140ca565b90615b1e6140ae565b90600090565b6040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9290921660048301526000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa801561204457600080928192615c09575b50615b9981615f17565b615ba283615f17565b825160ff831690808211615bd5575081511590811591615bcb575b50615bc85750615b0b565b92565b9050151538615bbd565b7f3d9055a70000000000000000000000000000000000000000000000000000000060005260ff841660045260245260446000fd5b9150615c2792503d8091833e615c1f8183610354565b810190615aa0565b909138615b8f565b6002548110156121245760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b818110615c7857505060009055565b80615c8560019285614134565b90549060031b1c6000528184016020526000604081205501615c69565b600081815260036020526040902054615d2d576002546801000000000000000081101561031657615d14615cdf8260018594016002556002614134565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615d8857805490680100000000000000008210156103165782615d71615cdf846001809601855584614134565b905580549260005201602052604060002055600190565b5050600090565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252615def604483610354565b6179185a10615e2b576020926000925191617530fa6000513d82615e1f575b5081615e18575090565b9050151590565b60201115915038615e0e565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252615def604483610354565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252615def604483610354565b80519060005b828110615f2957505050565b60018101808211613071575b838110615f455750600101615f1d565b73ffffffffffffffffffffffffffffffffffffffff615f648385612129565b5116615f76610b04610b4d8487612129565b14615f8357600101615f35565b611db5615f93610b4d8486612129565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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
