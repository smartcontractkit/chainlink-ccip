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
	SequenceNumber      uint64
	OnRampAddress       []byte
	OffRampAddress      []byte
	Finality            uint16
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
	ExtraData          []byte
}

type OffRampSourceChainConfig struct {
	Router           common.Address
	IsEnabled        bool
	OnRamp           []byte
	DefaultCCVs      []common.Address
	LaneMandatedCCVs []common.Address
}

type OffRampSourceChainConfigArgs struct {
	Router              common.Address
	SourceChainSelector uint64
	IsEnabled           bool
	OnRamp              []byte
	DefaultCCV          []common.Address
	LaneMandatedCCVs    []common.Address
}

type OffRampStaticConfig struct {
	LocalChainSelector   uint64
	GasForCallExactCheck uint16
	RmnRemote            common.Address
	TokenAdminRegistry   common.Address
}

var OffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f615a7a38819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615826908161025482396080518181816101560152610b09015260a0518181816101b90152610a63015260c0518181816101e10152818161466801526151b1015260e05181818161017d0152612c600152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780635215505b146100c85780636b8be52c146100c357806379ba5097146100be5780637ce1552a146100b95780638da5cb5b146100b4578063d2b33733146100af578063e9d68a8e146100aa578063f054ac57146100a55763f2fde38b146100a057600080fd5b61176a565b611416565b611336565b611263565b611211565b611172565b610fa1565b610897565b610713565b610546565b61041e565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610290565b828152826020820152826040820152015261025d60405161014b81610290565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ac57604052565b610261565b60a0810190811067ffffffffffffffff8211176102ac57604052565b6020810190811067ffffffffffffffff8211176102ac57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ac57604052565b6040519061033960a0836102e9565b565b6040519061033960c0836102e9565b60405190610339610100836102e9565b604051906103396040836102e9565b67ffffffffffffffff81116102ac57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906103b26020836102e9565b60008252565b60005b8381106103cb5750506000910152565b81810151838201526020016103bb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610417815180928187528780880191016103b8565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061045f81836102e9565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906103db565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104e75750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104da565b9161053f60ff916105316040949796976060875260608701906104c9565b9085820360208701526104c9565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576105d66105a461059e61025d93369060040161049b565b906134a9565b67ffffffffffffffff815116906105be60e082015161185e565b60601c61ffff60a06101208401519301511692613c14565b60409391935193849384610513565b6106509173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061063f61062d604085015160a0604086015260a08501906103db565b606085015184820360608601526104c9565b9201519060808184039101526104c9565b90565b6040810160408252825180915260206060830193019060005b8181106106f3575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106106a857505050505090565b90919293946020806106e4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516105e5565b97019301930191939290610699565b825167ffffffffffffffff1685526020948501949092019160010161066c565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461074e816118c8565b9061075c60405192836102e9565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610789826118c8565b0160005b81811061084f57505061079f8161190c565b9060005b8181106107bb57505061025d60405192839283610653565b806107f36107da6107cd600194615464565b67ffffffffffffffff1690565b6107e4838761199c565b9067ffffffffffffffff169052565b61083361082e610814610806848861199c565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b611b24565b61083d828761199c565b52610848818661199c565b50016107a3565b60209061085a6118e0565b8282870101520161078d565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576108e690369060040161049b565b60243567ffffffffffffffff81116100e757610906903690600401610866565b92909160443567ffffffffffffffff81116100e757610929903690600401610866565b93909261093c60015460ff9060a01c1690565b610f77576109a290610988740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b61099a61099585836134a9565b614055565b9336916110a7565b6020815191012093610a4a60206109ef6109c76107cd875167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610f7257600091610f43575b50610ef857610aae610814845167ffffffffffffffff1690565b8054610ac29060a01c60ff161590565b1590565b610ead57606084016001815160208151910120920191610ae183611aa5565b6020815191012003610e76575050602083015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603610e3f5750808603610e0d5760e083019384516014815103610dd35750610b9c610b67855167ffffffffffffffff1690565b6040860196610b7e885167ffffffffffffffff1690565b610b96610b9060c08a0151935161185e565b60601c90565b92614061565b96610bbb610bb4896000526005602052604060002090565b5460ff1690565b610bc481611146565b8015908115610dbf575b5015610d5057610cf7610ce8610ce396610806610cc77f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df298610cc28d8f9967ffffffffffffffff9b610c9691610d099b610c60610c358f6000526005602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957fd2b337330000000000000000000000000000000000000000000000000000000060208801528b60248801611e70565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b6140e2565b969015610d3b576002998a916000526005602052604060002090565b611c10565b965167ffffffffffffffff1690565b91836040519485941697169583612061565b0390a4610d397fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b005b6003998a916000526005602052604060002090565b610dbb8787610d79610d6a895167ffffffffffffffff1690565b915167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150610dcc81611146565b1438610bce565b610e09906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611bff565b0390fd5b7fb5ace4f300000000000000000000000000000000000000000000000000000000600052600486905260245260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5190610e096040519283927f2130095800000000000000000000000000000000000000000000000000000000845260048401611bda565b610dbb610ec2855167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610dbb610f0d845167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610f65915060203d602011610f6b575b610f5d81836102e9565b810190611bb9565b38610a94565b503d610f53565b611bce565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303611060577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff8116036100e757565b35906103398261108a565b9291926110b382610369565b916110c160405193846102e9565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e757816020610650933591016110a7565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561115057565b611117565b9060048210156111505752565b6020810192916103399190611155565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356111ad8161108a565b602435906111ba8261108a565b6044359167ffffffffffffffff83116100e7576111de6111f19336906004016110de565b90606435926111ec846110f9565b614061565b600052600560205261025d60ff6040600020541660405191829182611162565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760443567ffffffffffffffff81116100e7576112f1903690600401610866565b916064359267ffffffffffffffff84116100e757611316610d39943690600401610866565b939092602435906004016129ce565b9060206106509281815201906105e5565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff60043561137a8161108a565b6113826118e0565b5016600052600460205261025d60406000206114056003604051926113a6846102b1565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c16151560208501526040516113e9816113e28160018601611a03565b03826102e9565b60408501526113fa60028201611b10565b606085015201611b10565b608082015260405191829182611325565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757611465903690600401610866565b9061146e614abb565b6000905b82821061147b57005b61148e611489838584612db2565b612e6f565b60208101916114a86107cd845167ffffffffffffffff1690565b15611740576114ea6114d16114d1845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015611733575b6115415760808201949060005b8651805182101561156b576114d161151a836115349361199c565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15611541576001016114ff565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050939194929060009460a08701955b865180518210156115a3576114d161151a836115969361199c565b156115415760010161157b565b5050949590956060820180518051801591821561171d575b5050611541576107cd61151a946116da611713946116d08a6116c6611706976116bd61167c60019f9c6116127f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259e51895190614b06565b6116276108148b5167ffffffffffffffff1690565b9e8f6116366040840151151590565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b8d9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518d8c01612fa1565b5160028a016130cc565b51600388016130cc565b6116f76116f26107cd835167ffffffffffffffff1690565b6156d7565b505167ffffffffffffffff1690565b9260405191829182613160565b0390a20190611472565b60200120905061172b612f24565b1438806115bb565b50608082015151156114f2565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff6004356117ba816110f9565b6117c2614abb565b1633811461183457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611896575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b67ffffffffffffffff81116102ac5760051b60200190565b604051906118ed826102b1565b6060608083600081526000602082015282604082015282808201520152565b90611916826118c8565b61192360405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061195182946118c8565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156119975760200190565b61195b565b80518210156119975760209160051b010190565b90600182811c921680156119f9575b60208310146119ca57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916119bf565b60009291815491611a13836119b0565b8083529260018116908115611a695750600114611a2f57505050565b60009081526020812093945091925b838310611a4f575060209250010190565b600181602092949394548385870101520191019190611a3e565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b90610339611ab99260405193848092611a03565b03836102e9565b906020825491828152019160005260206000209060005b818110611ae45750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201611ad7565b90610339611ab99260405193848092611ac0565b9060036080604051611b35816102b1565b611bab819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c1615156020840152604051611b76816113e28160018601611a03565b6040840152604051611b8f816113e28160028601611ac0565b6060840152611ba46040518096819301611ac0565b03846102e9565b0152565b801515036100e757565b908160209103126100e7575161065081611baf565b6040513d6000823e3d90fd5b9091611bf161065093604084526040840190611a03565b9160208184039101526103db565b9060206106509281815201906103db565b9060048110156111505760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b838310611c7357505050505090565b9091929394602080611d03837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080611cf2611ce0611cce8786015160a08987015260a08601906103db565b604086015185820360408701526103db565b606085015184820360608601526103db565b9201519060808184039101526103db565b97019301930191939290611c64565b9160209082815201919060005b818110611d2c5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff8735611d55816110f9565b168152019401929101611d1f565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b848410611e1a5750505050505090565b909192939495602080611e60837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018852611e5a8b88611da2565b90611d63565b9801940194019294939190611e0a565b9492936106509694612040612053949360808952611e9b60808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a015261014061200d611fd7611fa1611f6a8d611f26611ef1606089015161016060e08501526101e08401906103db565b60808901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101008501526103db565b60a088015161ffff166101208301529060c088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b8d60e0870151906101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b6101008501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101808f01526103db565b6101208401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101a08e0152611c47565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016101c08b01526103db565b9260208801528683036040880152611d12565b926060818503910152611df2565b806120726040926106509594611155565b81602082015201906103db565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611896575050565b356106508161108a565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b61ffff8116036100e757565b3561065081612162565b91909160a0818403126100e75761218d61032a565b9281358452602082013567ffffffffffffffff81116100e757816121b29184016110de565b6020850152604082013567ffffffffffffffff81116100e757816121d79184016110de565b6040850152606082013567ffffffffffffffff81116100e757816121fc9184016110de565b6060850152608082013567ffffffffffffffff81116100e75761221f92016110de565b6080830152565b929190612232816118c8565b9361224060405195866102e9565b602085838152019160051b8101918383116100e75781905b838210612266575050505050565b813567ffffffffffffffff81116100e7576020916122878784938701612178565b815201910190612258565b90821015611997576122a99160051b81019061207f565b9091565b359061033982612162565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b8383106123335750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61843603018112156100e75760206124166001938683940190813581526124086123fd6123e26123c76123b788870187611da2565b60a08a88015260a0870191611d63565b6123d46040870187611da2565b908683036040880152611d63565b6123ef6060860186611da2565b908583036060870152611d63565b926080810190611da2565b916080818503910152611d63565b980196019493019190612323565b906106509593949273ffffffffffffffffffffffffffffffffffffffff612672921683526080602084015261246d6080840161245f8361109c565b67ffffffffffffffff169052565b61248d61247c6020830161109c565b67ffffffffffffffff1660a0850152565b6124ad61249c6040830161109c565b67ffffffffffffffff1660c0850152565b6126416126356125f66125b76125796125206124e26124cf6060890189611da2565b61016060e08d01526101e08c0191611d63565b6124ef6080890189611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808c8403016101008d0152611d63565b61253b61252f60a089016122ad565b61ffff166101208b0152565b61254860c0880188611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101408c0152611d63565b61258660e0870187611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8403016101608b0152611d63565b6125c5610100860186611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80898403016101808a0152611d63565b6126046101208501856122b8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888403016101a089015261230b565b91610140810190611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80858403016101c0860152611d63565b9360408201526060818503910152611d63565b604051906040820182811067ffffffffffffffff8211176102ac5760405260006020838281520152565b906126b9826118c8565b6126c660405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06126f482946118c8565b019060005b82811061270557505050565b602090612710612685565b828285010152016126f9565b91908110156119975760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100e7570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116036127bb57565b61275c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8f082019182116127bb57565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd12082019182116127bb57565b919082039182116127bb57565b90916060828403126100e757815161283e81611baf565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e75781519161286d83610369565b9161287b60405193846102e9565b838352602084830101116100e75760409261289c91602080850191016103b8565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806129196128e5604084015160a060c08801526101208701906103db565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103db565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106129965750505061ffff909516602083015261033992916060916040820152019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101612958565b939194909294303303612d8857612a33612a1c966129fb610b906129f560e08a018a61207f565b906120d0565b9260a0612a0789612104565b85612a2d8b612a2561012082019e8f8361210e565b9690920161216e565b943691612226565b916141b8565b97909160005b8351811015612ae957612a556114d16114d161151a848861199c565b90612a6b612a63828d61199c565b518789612292565b9290813b156100e7576000918a838d612ab4604051988996879586947f58bfa40a0000000000000000000000000000000000000000000000000000000086523060048701612424565b03925af1918215610f7257600192612ace575b5001612a39565b80612add6000612ae3936102e9565b806100dc565b38612ac7565b50949693509496505050612b07612b00828761210e565b90506126af565b9260005b612b15838861210e565b9050811015612b8a5780612b6e612b38600193612b32878c61210e565b9061271c565b86612b58612b688c612b60612b5060c083018361207f565b949092612104565b953690612178565b9236916110a7565b906145e5565b612b78828861199c565b52612b83818761199c565b5001612b0b565b509490509290926101408101612ba0818361207f565b90501580612d80575b8015612d77575b8015612d65575b612d5e57600093612c4a5a92612c37612c3e612bf56114d1612bdb6108148a612104565b5473ffffffffffffffffffffffffffffffffffffffff1690565b96612bff81612104565b93612c18612c1060c084018461207f565b92909361207f565b949095612c2361032a565b9b8c5267ffffffffffffffff1660208c0152565b36916110a7565b604088015236916110a7565b6060850152608084015283612ca8612ca3612c987f000000000000000000000000000000000000000000000000000000000000000094612c92612c8d8260061c90565b61278b565b9061281a565b61ffff85169061281a565b6127c0565b93612ce2604051978896879586947f3cf97983000000000000000000000000000000000000000000000000000000008652600486016128a2565b03925af1908115610f7257600090600092612d37575b5015612d015750565b610e09906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611bff565b9050612d5691503d806000833e612d4e81836102e9565b810190612827565b509038612cf8565b5050505050565b50612d72610abe866149bd565b612bb7565b50843b15612bb0565b506000612ba9565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156119975760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b3590610339826110f9565b359061033982611baf565b9080601f830112156100e7578135612e1f816118c8565b92612e2d60405194856102e9565b81845260208085019260051b8201019283116100e757602001905b828210612e555750505090565b602080918335612e64816110f9565b815201910190612e48565b60c0813603126100e757612e8161033b565b90612e8b81612df2565b8252612e996020820161109c565b6020830152612eaa60408201612dfd565b6040830152606081013567ffffffffffffffff81116100e757612ed090369083016110de565b6060830152608081013567ffffffffffffffff81116100e757612ef69036908301612e08565b608083015260a08101359067ffffffffffffffff82116100e757612f1c91369101612e08565b60a082015290565b60405160208101906000825260208152612f3f6040826102e9565b51902090565b818110612f50575050565b60008155600101612f45565b9190601f8111612f6b57505050565b610339926000526020600020906020601f840160051c83019310612f97575b601f0160051c0190612f45565b9091508190612f8a565b919091825167ffffffffffffffff81116102ac57612fc981612fc384546119b0565b84612f5c565b6020601f821160011461302757819061301893949560009261301c575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612fe6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061305a84600052602060002090565b9160005b8181106130b45750958360019596971061307d575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613073565b9192602060018192868b01518155019401920161305e565b81519167ffffffffffffffff83116102ac576801000000000000000083116102ac576020908254848455808510613143575b500190600052602060002060005b8381106131195750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161310c565b61315a908460005285846000209182019101612f45565b386130fe565b6003610650926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a060608401526131e26131af60c0850160018401611a03565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301611ac0565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201611ac0565b60405190610160820182811067ffffffffffffffff8211176102ac576040526060610140836000815260006020820152600060408201528280820152826080820152600060a08201528260c08201528260e082015282610100820152826101208201520152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146127bb5760010190565b90156119975790565b90821015611997570190565b90600882018092116127bb57565b90600282018092116127bb57565b90600182018092116127bb57565b90602082018092116127bb57565b919082018092116127bb57565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff0000000000000000000000000000000000000000000000008116926008811061334d575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffff000000000000000000000000000000000000000000000000000000000000811692600281106133b3575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b604051906133f2826102b1565b60606080836000815282602082015282604082015282808201520152565b6040805190919061342183826102e9565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b82811061345857505050565b6020906134636133e5565b8282850101520161344c565b6040519061347e6020836102e9565b600080835282815b82811061349257505050565b60209061349d6133e5565b82828501015201613486565b906134b2613213565b5060258110613b6e576134c3613213565b9160006135026134fc6134d685856132a7565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613b40575061354761353961353361352d61352460016132bc565b60018888613301565b90613319565b60c01c90565b67ffffffffffffffff168552565b6135b861358a61355760016132bc565b61358561357461353361352d61356c856132bc565b858b8b613301565b67ffffffffffffffff166020890152565b6132bc565b6135856135a761353361352d61359f856132bc565b858a8a613301565b67ffffffffffffffff166040880152565b6135db6135d56134fc6134d66135cd8561327a565b9488886132b0565b60ff1690565b90846135e783836132f4565b11613b13579081613609612c3761360184613613966132f4565b838989613301565b60608801526132f4565b83811015613ae5576136306135d56134fc6134d66135cd8561327a565b908461363c83836132f4565b11613ab8579081613656612c3761360184613660966132f4565b60808801526132f4565b8361366a826132ca565b11613a8b578061369f61369461368e61368861359f6136a4966132ca565b9061337f565b60f01c90565b61ffff1660a0880152565b6132ca565b83811015613a5d576136c16135d56134fc6134d66135cd8561327a565b90846136cd83836132f4565b11613a305790816136e7612c37613601846136f1966132f4565b60c08801526132f4565b83811015613a025761370e6135d56134fc6134d66135cd8561327a565b908461371a83836132f4565b116139d5579081613734612c376136018461373e966132f4565b60e08801526132f4565b83613748826132ca565b116139a75761ffff61377361376d61368e613688613765866132ca565b868a8a613301565b926132ca565b9116908461378183836132f4565b1161397a57908161379b612c37613601846137a6966132f4565b6101008801526132f4565b90836137b1836132ca565b1161394d575061ffff6137d761376d61368e6136886137cf866132ca565b868989613301565b9116806138e457506137e761346f565b6101208501525b826137f8826132ca565b116138b55761ffff61381561376d61368e6136886137cf866132ca565b9116908361382383836132f4565b1161388657613844612c3761384f94838761383e87836132f4565b92613301565b6101408601526132f4565b036138575790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b9061390a6139026138f3613410565b936101208801948552836132f4565b918585614ca6565b90925192613918829461198a565b52146137ee577fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008152600b600452602490fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008352600a600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526009600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526008600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526007600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526006600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526005600452602490fd5b507fb4205b4200000000000000000000000000000000000000000000000000000000815260048052602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526003600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526002600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526001600452602483fd5b7f789d326300000000000000000000000000000000000000000000000000000000825260ff16600452602490fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052610dbb6024906000600452565b60405190613bae6020836102e9565b6000808352366020840137565b80156127bb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff1680156127bb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9391613c1e613b9f565b93815180613f76575b505050613c349084615346565b92909293613c846002613c63613c696003613c638b67ffffffffffffffff166000526004602052604060002090565b01611b10565b9867ffffffffffffffff166000526004602052604060002090565b613caf613caa613ca2613c9a87518651906132f4565b8a51906132f4565b8351906132f4565b61190c565b9687956000805b8751821015613d095790613d01600192613ce6613cd661151a858d61199c565b91613ce08161327a565b9c61199c565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018997613cb6565b9750509193969092945060005b8551811015613d4b5780613d45613d3261151a6001948a61199c565b613ce6613d3e8b61327a565b9a8d61199c565b01613d16565b50919350919460005b8451811015613d895780613d83613d7061151a6001948961199c565b613ce6613d7c8a61327a565b998c61199c565b01613d54565b5093919592509360005b828110613f0d575b5050600090815b818110613e6e5750508152835160005b818110613dc157508452929190565b926000959495915b8351831015613e5f57613ddf61151a868861199c565b73ffffffffffffffffffffffffffffffffffffffff613e046114d161151a878961199c565b911603613e4e57613e1490613bbb565b90613e2f613e2561151a848961199c565b613ce6878961199c565b60ff8716613e3e575b90613dc9565b95613e4890613be6565b95613e38565b9091613e599061327a565b91613e38565b91509260019095949501613db2565b613e7b61151a828661199c565b73ffffffffffffffffffffffffffffffffffffffff81168015613f0357600090815b868110613ed7575b5050906001929115613eba575b505b01613da2565b613ed190613ce6613eca8761327a565b968861199c565b38613eb2565b81613ee86114d161151a848c61199c565b14613ef557600101613e9d565b506001915081905038613ea5565b5050600190613eb4565b613f1d6114d161151a838761199c565b15613f2a57600101613d93565b5091929360009591955b8351811015613f695780613f63613f5061151a6001948861199c565b613ce6613f5c8b61327a565b9a8961199c565b01613f34565b5093929150933880613d9b565b909192945060018103614028575060146060613f918461198a565b5101515103613fe55781613fdc91613fbb610b906060613fb3613c349761198a565b51015161185e565b91876080613fd3613fcb8461198a565b51519361198a565b51015193615144565b92903880613c27565b610e096060613ff38461198a565b5101516040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611bff565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61405d613213565b5090565b9290612f3f9173ffffffffffffffffffffffffffffffffffffffff6140af67ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906103db565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b90604051916140f260c0846102e9565b60848352602083019060a03683375a612ee08111156141435760009161411883926127ed565b82602083519301913090f1903d906084821161413a575b6000908286523e9190565b6084915061412f565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b91908110156119975760051b0190565b35610650816110f9565b906141c4939291613c14565b919390926141d18261190c565b926141db8361190c565b94600091825b88518110156142e2576000805b8a88848982851061425e575b50505050501561420c576001016141e1565b61421c61151a610dbb928b61199c565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b61151a6142969261429061428b8873ffffffffffffffffffffffffffffffffffffffff976114d19661419e565b6141ae565b9561199c565b9116146142a5576001016141ee565b600191506142c46142ba61428b838b8b61419e565b613ce6888c61199c565b6142d76142d08761327a565b968b61199c565b52388a8884896141fa565b509097965094939291909460ff811690816000985b8a518a10156143b45760005b8b878210806143ab575b1561439e5773ffffffffffffffffffffffffffffffffffffffff61433f6114d161151a8f61429061428b888f8f61419e565b9116146143545761434f9061327a565b614303565b939961436460019294939b613bbb565b9461438061437661428b838b8b61419e565b613ce68d8c61199c565b61439361438c8c61327a565b9b8b61199c565b525b019890916142f7565b5050919098600190614395565b5085151561430d565b9850925093959497509150816143dd57505050815181036143d457509190565b80825283529190565b610dbb92916143eb9161281a565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b908160209103126100e75751610650816110f9565b604051906103b2826102cd565b908160209103126100e75760405190614459826102cd565b51815290565b90610650916020815260e0614554614521614488855161010060208701526101208601906103db565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff166060860152606086015160808601526144ed608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103db565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103db565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103db565b3d156145b3573d9061459982610369565b916145a760405193846102e9565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff610650949316815281602082015201906103db565b9092916145f0612685565b50614601610b90606084015161185e565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa948515610f725760009561498c575b5073ffffffffffffffffffffffffffffffffffffffff8516948515614949576146c181614a13565b506146ce610abe82614a3d565b6149495750614807916020916147366146e78987615499565b966146f0614434565b50614780608082519261476261470b610b908a84015161185e565b6040519687918b830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752866102e9565b01519361476d61034a565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c08201526147cf6103a3565b60e0820152604051809381927f390775370000000000000000000000000000000000000000000000000000000083526004830161445f565b03816000885af160009181614918575b5061485b5784614825614588565b90610e096040519283927f9fe2f95a000000000000000000000000000000000000000000000000000000008452600484016145b8565b84909373ffffffffffffffffffffffffffffffffffffffff8316036148ae575b505050516148a661488a61035a565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b6148b791615499565b908082108015614904575b6148cc578361487b565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b5061490f818361281a565b835114156148c2565b61493b91925060203d602011614942575b61493381836102e9565b810190614441565b9038614817565b503d614929565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6149af91955060203d6020116149b6575b6149a781836102e9565b81019061441f565b9338614699565b503d61499d565b6149e77f85572ffb0000000000000000000000000000000000000000000000000000000082615675565b9081614a01575b816149f7575090565b6106509150615615565b9050614a0c8161554f565b15906149ee565b6149e77ff208a58f0000000000000000000000000000000000000000000000000000000082615675565b6149e77faff2afbf0000000000000000000000000000000000000000000000000000000082615675565b6149e77f479eecb20000000000000000000000000000000000000000000000000000000082615675565b6149e77f7909b5490000000000000000000000000000000000000000000000000000000082615675565b73ffffffffffffffffffffffffffffffffffffffff600154163303614adc57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614b148151846132f4565b928315614c415760005b848110614b2c575050505050565b81811015614c2657614b4161151a828661199c565b73ffffffffffffffffffffffffffffffffffffffff8116801561154157614b67836132d8565b878110614b7957505050600101614b1e565b84811015614bf65773ffffffffffffffffffffffffffffffffffffffff614ba361151a838a61199c565b168214614bb257600101614b67565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614c2161151a614c1b888561281a565b8961199c565b614ba3565b614c3c61151a614c36848461281a565b8561199c565b614b41565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b359060208110614c79575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b929192614cb16133e5565b938281101561502f57614cd46134fc6134d6614ccc8461327a565b9386866132b0565b600160ff821603614fff575082614cea826132e6565b11614fd05780614d10614d0a614d02614d17946132e6565b838787613301565b90614c6b565b86526132e6565b82811015614fa157614d3c6135d56134fc6134d6614d348561327a565b9487876132b0565b83614d4782846132f4565b11614f725781614d68612c37614d6084614d72966132f4565b838888613301565b60208801526132f4565b82811015614f4357614d8f6135d56134fc6134d6614d348561327a565b83614d9a82846132f4565b11614f145781614db3612c37614d6084614dbd966132f4565b60408801526132f4565b82811015614ee557614dda6135d56134fc6134d6614d348561327a565b83614de582846132f4565b11614eb65781613609612c37614d6084614dfe966132f4565b82614e08826132ca565b11614e875761ffff614e2561376d61368e6136886137cf866132ca565b91169183614e3384846132f4565b11614e5857612c37614e4e91836106509661383e87836132f4565b60808601526132f4565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b9080601f830112156100e7578151615075816118c8565b9261508360405194856102e9565b81845260208085019260051b8201019283116100e757602001905b8282106150ab5750505090565b6020809183516150ba816110f9565b81520191019061509e565b906020828203126100e757815167ffffffffffffffff81116100e757610650920161505e565b95949060019460a09467ffffffffffffffff61513f9573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906103db565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095929493929091906020838060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa928315610f72576000936152c9575b506151ec83614a67565b615226575b5050505050815115615201575090565b6106509150613c6360029167ffffffffffffffff166000526004602052604060002090565b6000949650859161527c73ffffffffffffffffffffffffffffffffffffffff92604051998a97889687957f89720a62000000000000000000000000000000000000000000000000000000008752600487016150eb565b0392165afa918215610f72576000926152a4575b5061529a8261575b565b38808080806151f1565b6152c29192503d806000833e6152ba81836102e9565b8101906150c5565b9038615290565b6152e391935060203d6020116149b6576149a781836102e9565b91386151e2565b90916060828403126100e757815167ffffffffffffffff81116100e7578361531391840161505e565b92602083015167ffffffffffffffff81116100e75760409161533691850161505e565b92015160ff811681036100e75790565b91909161536a61082e8267ffffffffffffffff166000526004602052604060002090565b90833b615387575b50606001519150615381613b9f565b90600090565b61539084614a91565b15615372576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015610f725760008094819261543c575b5061540b8161575b565b6154148561575b565b805115801590615430575b6154295750615372565b9392909150565b5060ff8216151561541f565b915061545b9294503d8091833e61545381836102e9565b8101906152ea565b90939138615401565b6002548110156119975760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181615518575b506155145782614825614588565b9150565b90916020823d602011615547575b81615533602093836102e9565b810103126155445750519038615506565b80fd5b3d9150615526565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526155af6044836102e9565b6179185a106155eb576020926000925191617530fa6000513d826155df575b50816155d8575090565b9050151590565b602011159150386155ce565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526155af6044836102e9565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526155af6044836102e9565b8060005260036020526040600020541560001461575557600254680100000000000000008110156102ac5760018101600255600060025482101561199757600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b50600090565b80519060005b82811061576d57505050565b600181018082116127bb575b8381106157895750600101615761565b73ffffffffffffffffffffffffffffffffffffffff6157a8838561199c565b51166157ba6114d161151a848761199c565b146157c757600101615779565b610dbb6157d761151a848661199c565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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

func (_OffRamp *OffRampCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber, sender, receiver)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_OffRamp *OffRampSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error) {
	return _OffRamp.Contract.GetExecutionState(&_OffRamp.CallOpts, sourceChainSelector, sequenceNumber, sender, receiver)
}

func (_OffRamp *OffRampCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error) {
	return _OffRamp.Contract.GetExecutionState(&_OffRamp.CallOpts, sourceChainSelector, sequenceNumber, sender, receiver)
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

func (_OffRamp *OffRampTransactor) Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "execute", encodedMessage, ccvs, ccvData)
}

func (_OffRamp *OffRampSession) Execute(encodedMessage []byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Execute(&_OffRamp.TransactOpts, encodedMessage, ccvs, ccvData)
}

func (_OffRamp *OffRampTransactorSession) Execute(encodedMessage []byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Execute(&_OffRamp.TransactOpts, encodedMessage, ccvs, ccvData)
}

func (_OffRamp *OffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "executeSingleMessage", message, messageId, ccvs, ccvData)
}

func (_OffRamp *OffRampSession) ExecuteSingleMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteSingleMessage(&_OffRamp.TransactOpts, message, messageId, ccvs, ccvData)
}

func (_OffRamp *OffRampTransactorSession) ExecuteSingleMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteSingleMessage(&_OffRamp.TransactOpts, message, messageId, ccvs, ccvData)
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
	SequenceNumber      uint64
	MessageId           [32]byte
	State               uint8
	ReturnData          []byte
	Raw                 types.Log
}

func (_OffRamp *OffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OffRampExecutionStateChangedIterator, error) {

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

	logs, sub, err := _OffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OffRampExecutionStateChangedIterator{contract: _OffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

func (_OffRamp *OffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sourceChainSelectorRule, sequenceNumberRule, messageIdRule)
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
	SourceConfig        OffRampSourceChainConfig
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

func (OffRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OffRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OffRampSourceChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e25")
}

func (OffRampStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e4950")
}

func (_OffRamp *OffRamp) Address() common.Address {
	return _OffRamp.address
}

type OffRampInterface interface {
	GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []OffRampSourceChainConfig, error)

	GetCCVsForMessage(opts *bind.CallOpts, encodedMessage []byte) (GetCCVsForMessage,

		error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (OffRampSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (OffRampStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterExecutionStateChanged(opts *bind.FilterOpts, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OffRampExecutionStateChangedIterator, error)

	WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *OffRampExecutionStateChanged, sourceChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

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
