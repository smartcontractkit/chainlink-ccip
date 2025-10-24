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
	Bin: "0x6101006040523461023857604051601f615a5b38819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615807908161025482396080518181816101560152610b09015260a0518181816101b90152610a63015260c0518181816101e1015281816146490152615192015260e05181818161017d0152612c410152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780635215505b146100c85780636b8be52c146100c357806379ba5097146100be5780637ce1552a146100b95780638da5cb5b146100b4578063d2b33733146100af578063e9d68a8e146100aa578063f054ac57146100a55763f2fde38b146100a057600080fd5b61176a565b611416565b611336565b611263565b611211565b611172565b610fa1565b610897565b610713565b610546565b61041e565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610290565b828152826020820152826040820152015261025d60405161014b81610290565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ac57604052565b610261565b60a0810190811067ffffffffffffffff8211176102ac57604052565b6020810190811067ffffffffffffffff8211176102ac57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ac57604052565b6040519061033960a0836102e9565b565b6040519061033960c0836102e9565b60405190610339610100836102e9565b604051906103396040836102e9565b67ffffffffffffffff81116102ac57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906103b26020836102e9565b60008252565b60005b8381106103cb5750506000910152565b81810151838201526020016103bb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610417815180928187528780880191016103b8565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061045f81836102e9565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906103db565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104e75750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104da565b9161053f60ff916105316040949796976060875260608701906104c9565b9085820360208701526104c9565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576105d66105a461059e61025d93369060040161049b565b9061348a565b67ffffffffffffffff815116906105be60e082015161185e565b60601c61ffff60a06101208401519301511692613bf5565b60409391935193849384610513565b6106509173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061063f61062d604085015160a0604086015260a08501906103db565b606085015184820360608601526104c9565b9201519060808184039101526104c9565b90565b6040810160408252825180915260206060830193019060005b8181106106f3575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106106a857505050505090565b90919293946020806106e4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516105e5565b97019301930191939290610699565b825167ffffffffffffffff1685526020948501949092019160010161066c565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461074e816118c8565b9061075c60405192836102e9565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610789826118c8565b0160005b81811061084f57505061079f8161190c565b9060005b8181106107bb57505061025d60405192839283610653565b806107f36107da6107cd600194615445565b67ffffffffffffffff1690565b6107e4838761199c565b9067ffffffffffffffff169052565b61083361082e610814610806848861199c565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b611b24565b61083d828761199c565b52610848818661199c565b50016107a3565b60209061085a6118e0565b8282870101520161078d565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576108e690369060040161049b565b60243567ffffffffffffffff81116100e757610906903690600401610866565b92909160443567ffffffffffffffff81116100e757610929903690600401610866565b93909261093c60015460ff9060a01c1690565b610f77576109a290610988740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b61099a610995858361348a565b614036565b9336916110a7565b6020815191012093610a4a60206109ef6109c76107cd875167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610f7257600091610f43575b50610ef857610aae610814845167ffffffffffffffff1690565b8054610ac29060a01c60ff161590565b1590565b610ead57606084016001815160208151910120920191610ae183611aa5565b6020815191012003610e76575050602083015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603610e3f5750808603610e0d5760e083019384516014815103610dd35750610b9c610b67855167ffffffffffffffff1690565b6040860196610b7e885167ffffffffffffffff1690565b610b96610b9060c08a0151935161185e565b60601c90565b92614042565b96610bbb610bb4896000526005602052604060002090565b5460ff1690565b610bc481611146565b8015908115610dbf575b5015610d5057610cf7610ce8610ce396610806610cc77f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df298610cc28d8f9967ffffffffffffffff9b610c9691610d099b610c60610c358f6000526005602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957fd2b337330000000000000000000000000000000000000000000000000000000060208801528b60248801611e70565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b6140c3565b969015610d3b576002998a916000526005602052604060002090565b611c10565b965167ffffffffffffffff1690565b91836040519485941697169583612061565b0390a4610d397fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b005b6003998a916000526005602052604060002090565b610dbb8787610d79610d6a895167ffffffffffffffff1690565b915167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150610dcc81611146565b1438610bce565b610e09906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611bff565b0390fd5b7fb5ace4f300000000000000000000000000000000000000000000000000000000600052600486905260245260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5190610e096040519283927f2130095800000000000000000000000000000000000000000000000000000000845260048401611bda565b610dbb610ec2855167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610dbb610f0d845167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610f65915060203d602011610f6b575b610f5d81836102e9565b810190611bb9565b38610a94565b503d610f53565b611bce565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303611060577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff8116036100e757565b35906103398261108a565b9291926110b382610369565b916110c160405193846102e9565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e757816020610650933591016110a7565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561115057565b611117565b9060048210156111505752565b6020810192916103399190611155565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356111ad8161108a565b602435906111ba8261108a565b6044359167ffffffffffffffff83116100e7576111de6111f19336906004016110de565b90606435926111ec846110f9565b614042565b600052600560205261025d60ff6040600020541660405191829182611162565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760443567ffffffffffffffff81116100e7576112f1903690600401610866565b916064359267ffffffffffffffff84116100e757611316610d39943690600401610866565b939092602435906004016129b0565b9060206106509281815201906105e5565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff60043561137a8161108a565b6113826118e0565b5016600052600460205261025d60406000206114056003604051926113a6846102b1565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c16151560208501526040516113e9816113e28160018601611a03565b03826102e9565b60408501526113fa60028201611b10565b606085015201611b10565b608082015260405191829182611325565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757611465903690600401610866565b9061146e614a9c565b6000905b82821061147b57005b61148e611489838584612d93565b612e50565b60208101916114a86107cd845167ffffffffffffffff1690565b15611740576114ea6114d16114d1845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015611733575b6115415760808201949060005b8651805182101561156b576114d161151a836115349361199c565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15611541576001016114ff565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050939194929060009460a08701955b865180518210156115a3576114d161151a836115969361199c565b156115415760010161157b565b5050949590956060820180518051801591821561171d575b5050611541576107cd61151a946116da611713946116d08a6116c6611706976116bd61167c60019f9c6116127f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259e51895190614ae7565b6116276108148b5167ffffffffffffffff1690565b9e8f6116366040840151151590565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b8d9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518d8c01612f82565b5160028a016130ad565b51600388016130ad565b6116f76116f26107cd835167ffffffffffffffff1690565b6156b8565b505167ffffffffffffffff1690565b9260405191829182613141565b0390a20190611472565b60200120905061172b612f05565b1438806115bb565b50608082015151156114f2565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff6004356117ba816110f9565b6117c2614a9c565b1633811461183457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611896575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b67ffffffffffffffff81116102ac5760051b60200190565b604051906118ed826102b1565b6060608083600081526000602082015282604082015282808201520152565b90611916826118c8565b61192360405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061195182946118c8565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156119975760200190565b61195b565b80518210156119975760209160051b010190565b90600182811c921680156119f9575b60208310146119ca57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916119bf565b60009291815491611a13836119b0565b8083529260018116908115611a695750600114611a2f57505050565b60009081526020812093945091925b838310611a4f575060209250010190565b600181602092949394548385870101520191019190611a3e565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b90610339611ab99260405193848092611a03565b03836102e9565b906020825491828152019160005260206000209060005b818110611ae45750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201611ad7565b90610339611ab99260405193848092611ac0565b9060036080604051611b35816102b1565b611bab819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c1615156020840152604051611b76816113e28160018601611a03565b6040840152604051611b8f816113e28160028601611ac0565b6060840152611ba46040518096819301611ac0565b03846102e9565b0152565b801515036100e757565b908160209103126100e7575161065081611baf565b6040513d6000823e3d90fd5b9091611bf161065093604084526040840190611a03565b9160208184039101526103db565b9060206106509281815201906103db565b9060048110156111505760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b838310611c7357505050505090565b9091929394602080611d03837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080611cf2611ce0611cce8786015160a08987015260a08601906103db565b604086015185820360408701526103db565b606085015184820360608601526103db565b9201519060808184039101526103db565b97019301930191939290611c64565b9160209082815201919060005b818110611d2c5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff8735611d55816110f9565b168152019401929101611d1f565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b848410611e1a5750505050505090565b909192939495602080611e60837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018852611e5a8b88611da2565b90611d63565b9801940194019294939190611e0a565b9492936106509694612040612053949360808952611e9b60808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a015261014061200d611fd7611fa1611f6a8d611f26611ef1606089015161016060e08501526101e08401906103db565b60808901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101008501526103db565b60a088015161ffff166101208301529060c088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b8d60e0870151906101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b6101008501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101808f01526103db565b6101208401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101a08e0152611c47565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016101c08b01526103db565b9260208801528683036040880152611d12565b926060818503910152611df2565b806120726040926106509594611155565b81602082015201906103db565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611896575050565b356106508161108a565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b61ffff8116036100e757565b3561065081612162565b91909160a0818403126100e75761218d61032a565b9281358452602082013567ffffffffffffffff81116100e757816121b29184016110de565b6020850152604082013567ffffffffffffffff81116100e757816121d79184016110de565b6040850152606082013567ffffffffffffffff81116100e757816121fc9184016110de565b6060850152608082013567ffffffffffffffff81116100e75761221f92016110de565b6080830152565b929190612232816118c8565b9361224060405195866102e9565b602085838152019160051b8101918383116100e75781905b838210612266575050505050565b813567ffffffffffffffff81116100e7576020916122878784938701612178565b815201910190612258565b90821015611997576122a99160051b81019061207f565b9091565b359061033982612162565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b8383106123335750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61843603018112156100e75760206124166001938683940190813581526124086123fd6123e26123c76123b788870187611da2565b60a08a88015260a0870191611d63565b6123d46040870187611da2565b908683036040880152611d63565b6123ef6060860186611da2565b908583036060870152611d63565b926080810190611da2565b916080818503910152611d63565b980196019493019190612323565b6126546106509593949260608352612450606084016124428361109c565b67ffffffffffffffff169052565b61247061245f6020830161109c565b67ffffffffffffffff166080850152565b61249061247f6040830161109c565b67ffffffffffffffff1660a0850152565b6126236126176125d861259961255b6125026124c56124b26060890189611da2565b61016060c08d01526101c08c0191611d63565b6124d26080890189611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c84030160e08d0152611d63565b61251d61251160a089016122ad565b61ffff166101008b0152565b61252a60c0880188611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101208c0152611d63565b61256860e0870187611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101408b0152611d63565b6125a7610100860186611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101608a0152611d63565b6125e66101208501856122b8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08884030161018089015261230b565b91610140810190611da2565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0858403016101a0860152611d63565b9360208201526040818503910152611d63565b604051906040820182811067ffffffffffffffff8211176102ac5760405260006020838281520152565b9061269b826118c8565b6126a860405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06126d682946118c8565b019060005b8281106126e757505050565b6020906126f2612667565b828285010152016126db565b91908110156119975760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100e7570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81160361279d57565b61273e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8f0820191821161279d57565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd120820191821161279d57565b9190820391821161279d57565b90916060828403126100e757815161282081611baf565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e75781519161284f83610369565b9161285d60405193846102e9565b838352602084830101116100e75760409261287e91602080850191016103b8565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806128fb6128c7604084015160a060c08801526101208701906103db565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103db565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106129785750505061ffff909516602083015261033992916060916040820152019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff168552602090810151818601526040909401939092019160010161293a565b939194909294303303612d6957612a156129fe966129dd610b906129d760e08a018a61207f565b906120d0565b9260a06129e989612104565b85612a0f8b612a0761012082019e8f8361210e565b9690920161216e565b943691612226565b91614199565b97909160005b8351811015612aca57612a376114d16114d161151a848861199c565b90612a4d612a45828d61199c565b518789612292565b9290813b156100e7576000918a838d612a95604051988996879586947fe8aa10be00000000000000000000000000000000000000000000000000000000865260048601612424565b03925af1918215610f7257600192612aaf575b5001612a1b565b80612abe6000612ac4936102e9565b806100dc565b38612aa8565b50949693509496505050612ae8612ae1828761210e565b9050612691565b9260005b612af6838861210e565b9050811015612b6b5780612b4f612b19600193612b13878c61210e565b906126fe565b86612b39612b498c612b41612b3160c083018361207f565b949092612104565b953690612178565b9236916110a7565b906145c6565b612b59828861199c565b52612b64818761199c565b5001612aec565b509490509290926101408101612b81818361207f565b90501580612d61575b8015612d58575b8015612d46575b612d3f57600093612c2b5a92612c18612c1f612bd66114d1612bbc6108148a612104565b5473ffffffffffffffffffffffffffffffffffffffff1690565b96612be081612104565b93612bf9612bf160c084018461207f565b92909361207f565b949095612c0461032a565b9b8c5267ffffffffffffffff1660208c0152565b36916110a7565b604088015236916110a7565b6060850152608084015283612c89612c84612c797f000000000000000000000000000000000000000000000000000000000000000094612c73612c6e8260061c90565b61276d565b906127fc565b61ffff8516906127fc565b6127a2565b93612cc3604051978896879586947f3cf9798300000000000000000000000000000000000000000000000000000000865260048601612884565b03925af1908115610f7257600090600092612d18575b5015612ce25750565b610e09906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611bff565b9050612d3791503d806000833e612d2f81836102e9565b810190612809565b509038612cd9565b5050505050565b50612d53610abe8661499e565b612b98565b50843b15612b91565b506000612b8a565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156119975760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b3590610339826110f9565b359061033982611baf565b9080601f830112156100e7578135612e00816118c8565b92612e0e60405194856102e9565b81845260208085019260051b8201019283116100e757602001905b828210612e365750505090565b602080918335612e45816110f9565b815201910190612e29565b60c0813603126100e757612e6261033b565b90612e6c81612dd3565b8252612e7a6020820161109c565b6020830152612e8b60408201612dde565b6040830152606081013567ffffffffffffffff81116100e757612eb190369083016110de565b6060830152608081013567ffffffffffffffff81116100e757612ed79036908301612de9565b608083015260a08101359067ffffffffffffffff82116100e757612efd91369101612de9565b60a082015290565b60405160208101906000825260208152612f206040826102e9565b51902090565b818110612f31575050565b60008155600101612f26565b9190601f8111612f4c57505050565b610339926000526020600020906020601f840160051c83019310612f78575b601f0160051c0190612f26565b9091508190612f6b565b919091825167ffffffffffffffff81116102ac57612faa81612fa484546119b0565b84612f3d565b6020601f8211600114613008578190612ff9939495600092612ffd575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612fc7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061303b84600052602060002090565b9160005b8181106130955750958360019596971061305e575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613054565b9192602060018192868b01518155019401920161303f565b81519167ffffffffffffffff83116102ac576801000000000000000083116102ac576020908254848455808510613124575b500190600052602060002060005b8381106130fa5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016130ed565b61313b908460005285846000209182019101612f26565b386130df565b6003610650926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a060608401526131c361319060c0850160018401611a03565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301611ac0565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201611ac0565b60405190610160820182811067ffffffffffffffff8211176102ac576040526060610140836000815260006020820152600060408201528280820152826080820152600060a08201528260c08201528260e082015282610100820152826101208201520152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461279d5760010190565b90156119975790565b90821015611997570190565b906008820180921161279d57565b906002820180921161279d57565b906001820180921161279d57565b906020820180921161279d57565b9190820180921161279d57565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff0000000000000000000000000000000000000000000000008116926008811061332e575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110613394575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b604051906133d3826102b1565b60606080836000815282602082015282604082015282808201520152565b6040805190919061340283826102e9565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b82811061343957505050565b6020906134446133c6565b8282850101520161342d565b6040519061345f6020836102e9565b600080835282815b82811061347357505050565b60209061347e6133c6565b82828501015201613467565b906134936131f4565b5060258110613b4f576134a46131f4565b9160006134e36134dd6134b78585613288565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613b21575061352861351a61351461350e613505600161329d565b600188886132e2565b906132fa565b60c01c90565b67ffffffffffffffff168552565b61359961356b613538600161329d565b61356661355561351461350e61354d8561329d565b858b8b6132e2565b67ffffffffffffffff166020890152565b61329d565b61356661358861351461350e6135808561329d565b858a8a6132e2565b67ffffffffffffffff166040880152565b6135bc6135b66134dd6134b76135ae8561325b565b948888613291565b60ff1690565b90846135c883836132d5565b11613af45790816135ea612c186135e2846135f4966132d5565b8389896132e2565b60608801526132d5565b83811015613ac6576136116135b66134dd6134b76135ae8561325b565b908461361d83836132d5565b11613a99579081613637612c186135e284613641966132d5565b60808801526132d5565b8361364b826132ab565b11613a6c578061368061367561366f613669613580613685966132ab565b90613360565b60f01c90565b61ffff1660a0880152565b6132ab565b83811015613a3e576136a26135b66134dd6134b76135ae8561325b565b90846136ae83836132d5565b11613a115790816136c8612c186135e2846136d2966132d5565b60c08801526132d5565b838110156139e3576136ef6135b66134dd6134b76135ae8561325b565b90846136fb83836132d5565b116139b6579081613715612c186135e28461371f966132d5565b60e08801526132d5565b83613729826132ab565b116139885761ffff61375461374e61366f613669613746866132ab565b868a8a6132e2565b926132ab565b9116908461376283836132d5565b1161395b57908161377c612c186135e284613787966132d5565b6101008801526132d5565b9083613792836132ab565b1161392e575061ffff6137b861374e61366f6136696137b0866132ab565b8689896132e2565b9116806138c557506137c8613450565b6101208501525b826137d9826132ab565b116138965761ffff6137f661374e61366f6136696137b0866132ab565b9116908361380483836132d5565b1161386757613825612c1861383094838761381f87836132d5565b926132e2565b6101408601526132d5565b036138385790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b906138eb6138e36138d46133f1565b936101208801948552836132d5565b918585614c87565b909251926138f9829461198a565b52146137cf577fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008152600b600452602490fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008352600a600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526009600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526008600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526007600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526006600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526005600452602490fd5b507fb4205b4200000000000000000000000000000000000000000000000000000000815260048052602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526003600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526002600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526001600452602483fd5b7f789d326300000000000000000000000000000000000000000000000000000000825260ff16600452602490fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052610dbb6024906000600452565b60405190613b8f6020836102e9565b6000808352366020840137565b801561279d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff16801561279d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9391613bff613b80565b93815180613f57575b505050613c159084615327565b92909293613c656002613c44613c4a6003613c448b67ffffffffffffffff166000526004602052604060002090565b01611b10565b9867ffffffffffffffff166000526004602052604060002090565b613c90613c8b613c83613c7b87518651906132d5565b8a51906132d5565b8351906132d5565b61190c565b9687956000805b8751821015613cea5790613ce2600192613cc7613cb761151a858d61199c565b91613cc18161325b565b9c61199c565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018997613c97565b9750509193969092945060005b8551811015613d2c5780613d26613d1361151a6001948a61199c565b613cc7613d1f8b61325b565b9a8d61199c565b01613cf7565b50919350919460005b8451811015613d6a5780613d64613d5161151a6001948961199c565b613cc7613d5d8a61325b565b998c61199c565b01613d35565b5093919592509360005b828110613eee575b5050600090815b818110613e4f5750508152835160005b818110613da257508452929190565b926000959495915b8351831015613e4057613dc061151a868861199c565b73ffffffffffffffffffffffffffffffffffffffff613de56114d161151a878961199c565b911603613e2f57613df590613b9c565b90613e10613e0661151a848961199c565b613cc7878961199c565b60ff8716613e1f575b90613daa565b95613e2990613bc7565b95613e19565b9091613e3a9061325b565b91613e19565b91509260019095949501613d93565b613e5c61151a828661199c565b73ffffffffffffffffffffffffffffffffffffffff81168015613ee457600090815b868110613eb8575b5050906001929115613e9b575b505b01613d83565b613eb290613cc7613eab8761325b565b968861199c565b38613e93565b81613ec96114d161151a848c61199c565b14613ed657600101613e7e565b506001915081905038613e86565b5050600190613e95565b613efe6114d161151a838761199c565b15613f0b57600101613d74565b5091929360009591955b8351811015613f4a5780613f44613f3161151a6001948861199c565b613cc7613f3d8b61325b565b9a8961199c565b01613f15565b5093929150933880613d7c565b909192945060018103614009575060146060613f728461198a565b5101515103613fc65781613fbd91613f9c610b906060613f94613c159761198a565b51015161185e565b91876080613fb4613fac8461198a565b51519361198a565b51015193615125565b92903880613c08565b610e096060613fd48461198a565b5101516040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611bff565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61403e6131f4565b5090565b9290612f209173ffffffffffffffffffffffffffffffffffffffff61409067ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906103db565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b90604051916140d360c0846102e9565b60848352602083019060a03683375a612ee0811115614124576000916140f983926127cf565b82602083519301913090f1903d906084821161411b575b6000908286523e9190565b60849150614110565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b91908110156119975760051b0190565b35610650816110f9565b906141a5939291613bf5565b919390926141b28261190c565b926141bc8361190c565b94600091825b88518110156142c3576000805b8a88848982851061423f575b5050505050156141ed576001016141c2565b6141fd61151a610dbb928b61199c565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b61151a6142779261427161426c8873ffffffffffffffffffffffffffffffffffffffff976114d19661417f565b61418f565b9561199c565b911614614286576001016141cf565b600191506142a561429b61426c838b8b61417f565b613cc7888c61199c565b6142b86142b18761325b565b968b61199c565b52388a8884896141db565b509097965094939291909460ff811690816000985b8a518a10156143955760005b8b8782108061438c575b1561437f5773ffffffffffffffffffffffffffffffffffffffff6143206114d161151a8f61427161426c888f8f61417f565b911614614335576143309061325b565b6142e4565b939961434560019294939b613b9c565b9461436161435761426c838b8b61417f565b613cc78d8c61199c565b61437461436d8c61325b565b9b8b61199c565b525b019890916142d8565b5050919098600190614376565b508515156142ee565b9850925093959497509150816143be57505050815181036143b557509190565b80825283529190565b610dbb92916143cc916127fc565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b908160209103126100e75751610650816110f9565b604051906103b2826102cd565b908160209103126100e7576040519061443a826102cd565b51815290565b90610650916020815260e0614535614502614469855161010060208701526101208601906103db565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff166060860152606086015160808601526144ce608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103db565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103db565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103db565b3d15614594573d9061457a82610369565b9161458860405193846102e9565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff610650949316815281602082015201906103db565b9092916145d1612667565b506145e2610b90606084015161185e565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa948515610f725760009561496d575b5073ffffffffffffffffffffffffffffffffffffffff851694851561492a576146a2816149f4565b506146af610abe82614a1e565b61492a57506147e8916020916147176146c8898761547a565b966146d1614415565b5061476160808251926147436146ec610b908a84015161185e565b6040519687918b830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752866102e9565b01519361474e61034a565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c08201526147b06103a3565b60e0820152604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301614440565b03816000885af1600091816148f9575b5061483c5784614806614569565b90610e096040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401614599565b84909373ffffffffffffffffffffffffffffffffffffffff83160361488f575b5050505161488761486b61035a565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b6148989161547a565b9080821080156148e5575b6148ad578361485c565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b506148f081836127fc565b835114156148a3565b61491c91925060203d602011614923575b61491481836102e9565b810190614422565b90386147f8565b503d61490a565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b61499091955060203d602011614997575b61498881836102e9565b810190614400565b933861467a565b503d61497e565b6149c87f85572ffb0000000000000000000000000000000000000000000000000000000082615656565b90816149e2575b816149d8575090565b61065091506155f6565b90506149ed81615530565b15906149cf565b6149c87ff208a58f0000000000000000000000000000000000000000000000000000000082615656565b6149c87faff2afbf0000000000000000000000000000000000000000000000000000000082615656565b6149c87ff00a89320000000000000000000000000000000000000000000000000000000082615656565b6149c87f7909b5490000000000000000000000000000000000000000000000000000000082615656565b73ffffffffffffffffffffffffffffffffffffffff600154163303614abd57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614af58151846132d5565b928315614c225760005b848110614b0d575050505050565b81811015614c0757614b2261151a828661199c565b73ffffffffffffffffffffffffffffffffffffffff8116801561154157614b48836132b9565b878110614b5a57505050600101614aff565b84811015614bd75773ffffffffffffffffffffffffffffffffffffffff614b8461151a838a61199c565b168214614b9357600101614b48565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614c0261151a614bfc88856127fc565b8961199c565b614b84565b614c1d61151a614c1784846127fc565b8561199c565b614b22565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b359060208110614c5a575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b929192614c926133c6565b938281101561501057614cb56134dd6134b7614cad8461325b565b938686613291565b600160ff821603614fe0575082614ccb826132c7565b11614fb15780614cf1614ceb614ce3614cf8946132c7565b8387876132e2565b90614c4c565b86526132c7565b82811015614f8257614d1d6135b66134dd6134b7614d158561325b565b948787613291565b83614d2882846132d5565b11614f535781614d49612c18614d4184614d53966132d5565b8388886132e2565b60208801526132d5565b82811015614f2457614d706135b66134dd6134b7614d158561325b565b83614d7b82846132d5565b11614ef55781614d94612c18614d4184614d9e966132d5565b60408801526132d5565b82811015614ec657614dbb6135b66134dd6134b7614d158561325b565b83614dc682846132d5565b11614e9757816135ea612c18614d4184614ddf966132d5565b82614de9826132ab565b11614e685761ffff614e0661374e61366f6136696137b0866132ab565b91169183614e1484846132d5565b11614e3957612c18614e2f91836106509661381f87836132d5565b60808601526132d5565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b9080601f830112156100e7578151615056816118c8565b9261506460405194856102e9565b81845260208085019260051b8201019283116100e757602001905b82821061508c5750505090565b60208091835161509b816110f9565b81520191019061507f565b906020828203126100e757815167ffffffffffffffff81116100e757610650920161503f565b95949060019460a09467ffffffffffffffff6151209573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906103db565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095929493929091906020838060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa928315610f72576000936152aa575b506151cd83614a48565b615207575b50505050508151156151e2575090565b6106509150613c4460029167ffffffffffffffff166000526004602052604060002090565b6000949650859161525d73ffffffffffffffffffffffffffffffffffffffff92604051998a97889687957f89720a62000000000000000000000000000000000000000000000000000000008752600487016150cc565b0392165afa918215610f7257600092615285575b5061527b8261573c565b38808080806151d2565b6152a39192503d806000833e61529b81836102e9565b8101906150a6565b9038615271565b6152c491935060203d6020116149975761498881836102e9565b91386151c3565b90916060828403126100e757815167ffffffffffffffff81116100e757836152f491840161503f565b92602083015167ffffffffffffffff81116100e75760409161531791850161503f565b92015160ff811681036100e75790565b91909161534b61082e8267ffffffffffffffff166000526004602052604060002090565b90833b615368575b50606001519150615362613b80565b90600090565b61537184614a72565b15615353576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015610f725760008094819261541d575b506153ec8161573c565b6153f58561573c565b805115801590615411575b61540a5750615353565b9392909150565b5060ff82161515615400565b915061543c9294503d8091833e61543481836102e9565b8101906152cb565b909391386153e2565b6002548110156119975760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa600091816154f9575b506154f55782614806614569565b9150565b90916020823d602011615528575b81615514602093836102e9565b8101031261552557505190386154e7565b80fd5b3d9150615507565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526155906044836102e9565b6179185a106155cc576020926000925191617530fa6000513d826155c0575b50816155b9575090565b9050151590565b602011159150386155af565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526155906044836102e9565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526155906044836102e9565b8060005260036020526040600020541560001461573657600254680100000000000000008110156102ac5760018101600255600060025482101561199757600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b50600090565b80519060005b82811061574e57505050565b6001810180821161279d575b83811061576a5750600101615742565b73ffffffffffffffffffffffffffffffffffffffff615789838561199c565b511661579b6114d161151a848761199c565b146157a85760010161575a565b610dbb6157b861151a848661199c565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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
