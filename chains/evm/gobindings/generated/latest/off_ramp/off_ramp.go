// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package off_ramp

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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enumMessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f615a7638819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615822908161025482396080518181816101560152610ae0015260a0518181816101b90152610a63015260c0518181816101e1015281816146400152615118015260e05181818161017d0152612bf80152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780635215505b146100c85780636b8be52c146100c357806379ba5097146100be5780637ce1552a146100b95780638da5cb5b146100b4578063d2b33733146100af578063e9d68a8e146100aa578063f054ac57146100a55763f2fde38b146100a057600080fd5b61170a565b6113b6565b6112d6565b611203565b6111b1565b611112565b610f41565b610897565b610713565b610546565b61041e565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610290565b828152826020820152826040820152015261025d60405161014b81610290565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ac57604052565b610261565b60a0810190811067ffffffffffffffff8211176102ac57604052565b6020810190811067ffffffffffffffff8211176102ac57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ac57604052565b6040519061033960a0836102e9565b565b6040519061033960c0836102e9565b60405190610339610100836102e9565b604051906103396040836102e9565b67ffffffffffffffff81116102ac57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906103b26020836102e9565b60008252565b60005b8381106103cb5750506000910152565b81810151838201526020016103bb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610417815180928187528780880191016103b8565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061045f81836102e9565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906103db565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104e75750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104da565b9161053f60ff916105316040949796976060875260608701906104c9565b9085820360208701526104c9565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576105d66105a461059e61025d93369060040161049b565b90613466565b67ffffffffffffffff815116906105be60e08201516117fe565b60601c61ffff60a06101208401519301511692613bd8565b60409391935193849384610513565b6106509173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061063f61062d604085015160a0604086015260a08501906103db565b606085015184820360608601526104c9565b9201519060808184039101526104c9565b90565b6040810160408252825180915260206060830193019060005b8181106106f3575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106106a857505050505090565b90919293946020806106e4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516105e5565b97019301930191939290610699565b825167ffffffffffffffff1685526020948501949092019160010161066c565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461074e81611868565b9061075c60405192836102e9565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061078982611868565b0160005b81811061084f57505061079f816118ac565b9060005b8181106107bb57505061025d60405192839283610653565b806107f36107da6107cd6001946153cb565b67ffffffffffffffff1690565b6107e4838761193c565b9067ffffffffffffffff169052565b61083361082e610814610806848861193c565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b611ab0565b61083d828761193c565b52610848818661193c565b50016107a3565b60209061085a611880565b8282870101520161078d565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576108e690369060040161049b565b60243567ffffffffffffffff81116100e757610906903690600401610866565b92909160443567ffffffffffffffff81116100e757610929903690600401610866565b93909261093c60015460ff9060a01c1690565b610f17576109a290610988740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b61099a6109958583613466565b614019565b933691611047565b6020815191012093610a4a60206109ef6109c76107cd875167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610f1257600091610ee3575b50610e9857610ac2610abe610ab4610814865167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b610e4d57602083015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603610e165750808603610de45760e083019384516014815103610daa5750610b73610b3e855167ffffffffffffffff1690565b6040860196610b55885167ffffffffffffffff1690565b610b6d610b6760c08a015193516117fe565b60601c90565b92614025565b96610b92610b8b896000526005602052604060002090565b5460ff1690565b610b9b816110e6565b8015908115610d96575b5015610d2757610cce610cbf610cba96610806610c9e7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df298610c998d8f9967ffffffffffffffff9b610c6d91610ce09b610c37610c0c8f6000526005602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957fd2b337330000000000000000000000000000000000000000000000000000000060208801528b60248801611dd7565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b6140a6565b969015610d12576002998a916000526005602052604060002090565b611b77565b965167ffffffffffffffff1690565b91836040519485941697169583611fc8565b0390a4610d107fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b005b6003998a916000526005602052604060002090565b610d928787610d50610d41895167ffffffffffffffff1690565b915167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150610da3816110e6565b1438610ba5565b610de0906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611b66565b0390fd5b7fb5ace4f300000000000000000000000000000000000000000000000000000000600052600486905260245260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b610d92610e62845167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610d92610ead845167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610f05915060203d602011610f0b575b610efd81836102e9565b810190611b45565b38610a94565b503d610ef3565b611b5a565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303611000577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff8116036100e757565b35906103398261102a565b92919261105382610369565b9161106160405193846102e9565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e75781602061065093359101611047565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600411156110f057565b6110b7565b9060048210156110f05752565b60208101929161033991906110f5565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043561114d8161102a565b6024359061115a8261102a565b6044359167ffffffffffffffff83116100e75761117e61119193369060040161107e565b906064359261118c84611099565b614025565b600052600560205261025d60ff6040600020541660405191829182611102565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760443567ffffffffffffffff81116100e757611291903690600401610866565b916064359267ffffffffffffffff84116100e7576112b6610d10943690600401610866565b93909260243590600401612967565b9060206106509281815201906105e5565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff60043561131a8161102a565b611322611880565b5016600052600460205261025d60406000206113a5600360405192611346846102b1565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c16151560208501526040516113898161138281600186016119a3565b03826102e9565b604085015261139a60028201611a95565b606085015201611a95565b6080820152604051918291826112c5565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757611405903690600401610866565b9061140e614a2e565b6000905b82821061141b57005b61142e611429838584612d6f565b612e2c565b60208101916114486107cd845167ffffffffffffffff1690565b156116e05761148a611471611471845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b1580156116d3575b6114e15760808201949060005b8651805182101561150b576114716114ba836114d49361193c565b5173ffffffffffffffffffffffffffffffffffffffff1690565b156114e15760010161149f565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050939194929060009460a08701955b86518051821015611543576114716114ba836115369361193c565b156114e15760010161151b565b505094959095606082018051805180159182156116bd575b50506114e1576107cd6114ba9461167a6116b3946116708a6116666116a69761165d61161c60019f9c6115b27f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259e51895190614a79565b6115c76108148b5167ffffffffffffffff1690565b9e8f6115d66040840151151590565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b8d9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518d8c01612f5e565b5160028a01613089565b5160038801613089565b6116976116926107cd835167ffffffffffffffff1690565b6156d3565b505167ffffffffffffffff1690565b926040519182918261311d565b0390a20190611412565b6020012090506116cb612ee1565b14388061155b565b5060808201515115611492565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561175a81611099565b611762614a2e565b163381146117d457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611836575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b67ffffffffffffffff81116102ac5760051b60200190565b6040519061188d826102b1565b6060608083600081526000602082015282604082015282808201520152565b906118b682611868565b6118c360405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06118f18294611868565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156119375760200190565b6118fb565b80518210156119375760209160051b010190565b90600182811c92168015611999575b602083101461196a57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161195f565b600092918154916119b383611950565b8083529260018116908115611a0957506001146119cf57505050565b60009081526020812093945091925b8383106119ef575060209250010190565b6001816020929493945483858701015201910191906119de565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b906020825491828152019160005260206000209060005b818110611a695750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201611a5c565b90610339611aa99260405193848092611a45565b03836102e9565b9060036080604051611ac1816102b1565b611b37819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c1615156020840152604051611b028161138281600186016119a3565b6040840152604051611b1b816113828160028601611a45565b6060840152611b306040518096819301611a45565b03846102e9565b0152565b801515036100e757565b908160209103126100e7575161065081611b3b565b6040513d6000823e3d90fd5b9060206106509281815201906103db565b9060048110156110f05760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b838310611bda57505050505090565b9091929394602080611c6a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080611c59611c47611c358786015160a08987015260a08601906103db565b604086015185820360408701526103db565b606085015184820360608601526103db565b9201519060808184039101526103db565b97019301930191939290611bcb565b9160209082815201919060005b818110611c935750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff8735611cbc81611099565b168152019401929101611c86565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b848410611d815750505050505090565b909192939495602080611dc7837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018852611dc18b88611d09565b90611cca565b9801940194019294939190611d71565b9492936106509694611fa7611fba949360808952611e0260808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a0152610140611f74611f3e611f08611ed18d611e8d611e58606089015161016060e08501526101e08401906103db565b60808901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101008501526103db565b60a088015161ffff166101208301529060c088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b8d60e0870151906101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b6101008501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101808f01526103db565b6101208401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101a08e0152611bae565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016101c08b01526103db565b9260208801528683036040880152611c79565b926060818503910152611d59565b80611fd960409261065095946110f5565b81602082015201906103db565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611836575050565b356106508161102a565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b61ffff8116036100e757565b35610650816120c9565b91909160a0818403126100e7576120f461032a565b9281358452602082013567ffffffffffffffff81116100e7578161211991840161107e565b6020850152604082013567ffffffffffffffff81116100e7578161213e91840161107e565b6040850152606082013567ffffffffffffffff81116100e7578161216391840161107e565b6060850152608082013567ffffffffffffffff81116100e757612186920161107e565b6080830152565b92919061219981611868565b936121a760405195866102e9565b602085838152019160051b8101918383116100e75781905b8382106121cd575050505050565b813567ffffffffffffffff81116100e7576020916121ee87849387016120df565b8152019101906121bf565b90821015611937576122109160051b810190611fe6565b9091565b3590610339826120c9565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b83831061229a5750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61843603018112156100e757602061237d60019386839401908135815261236f61236461234961232e61231e88870187611d09565b60a08a88015260a0870191611cca565b61233b6040870187611d09565b908683036040880152611cca565b6123566060860186611d09565b908583036060870152611cca565b926080810190611d09565b916080818503910152611cca565b98019601949301919061228a565b906106509593949273ffffffffffffffffffffffffffffffffffffffff6125d992168352608060208401526123d4608084016123c68361103c565b67ffffffffffffffff169052565b6123f46123e36020830161103c565b67ffffffffffffffff1660a0850152565b6124146124036040830161103c565b67ffffffffffffffff1660c0850152565b6125a861259c61255d61251e6124e06124876124496124366060890189611d09565b61016060e08d01526101e08c0191611cca565b6124566080890189611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808c8403016101008d0152611cca565b6124a261249660a08901612214565b61ffff166101208b0152565b6124af60c0880188611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101408c0152611cca565b6124ed60e0870187611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8403016101608b0152611cca565b61252c610100860186611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80898403016101808a0152611cca565b61256b61012085018561221f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888403016101a0890152612272565b91610140810190611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80858403016101c0860152611cca565b9360408201526060818503910152611cca565b604051906040820182811067ffffffffffffffff8211176102ac5760405260006020838281520152565b9061262082611868565b61262d60405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061265b8294611868565b019060005b82811061266c57505050565b6020906126776125ec565b82828501015201612660565b91908110156119375760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100e7570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60011b906201fffe61fffe83169216820361270957565b6126c3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd120820191821161270957565b9190820391821161270957565b908160041b917f0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81160361270957565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81160361270957565b90610f81820291808304610f81149015171561270957565b90916060828403126100e75781516127d781611b3b565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e75781519161280683610369565b9161281460405193846102e9565b838352602084830101116100e75760409261283591602080850191016103b8565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806128b261287e604084015160a060c08801526101208701906103db565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103db565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b81811061292f5750505061ffff909516602083015261033992916060916040820152019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff16855260209081015181860152604090940193909201916001016128f1565b939194909294303303612d45576129cc6129b596612994610b6761298e60e08a018a611fe6565b90612037565b9260a06129a08961206b565b856129c68b6129be61012082019e8f83612075565b969092016120d5565b94369161218d565b9161417c565b97909160005b8351811015612a82576129ee6114716114716114ba848861193c565b90612a046129fc828d61193c565b5187896121f9565b9290813b156100e7576000918a838d612a4d604051988996879586947f58bfa40a000000000000000000000000000000000000000000000000000000008652306004870161238b565b03925af1918215610f1257600192612a67575b50016129d2565b80612a766000612a7c936102e9565b806100dc565b38612a60565b50949693509496505050612aa0612a998287612075565b9050612616565b9260005b612aae8388612075565b9050811015612b235780612b07612ad1600193612acb878c612075565b90612683565b86612af1612b018c612af9612ae960c0830183611fe6565b94909261206b565b9536906120df565b923691611047565b906145a9565b612b11828861193c565b52612b1c818761193c565b5001612aa4565b509490509290926101408101612b398183611fe6565b90501580612d3d575b8015612d34575b8015612d22575b612d1b57612b5d8261206b565b612b7b9067ffffffffffffffff166000526004602052604060002090565b5473ffffffffffffffffffffffffffffffffffffffff1693612b9c8361206b565b612ba960c0850185611fe6565b90612bb48587611fe6565b929093612bbf61032a565b98895267ffffffffffffffff1660208901523690612bdc92611047565b60408701523690612bec92611047565b606085015260808401527f0000000000000000000000000000000000000000000000000000000000000000915a612c22846126f2565b61ffff16612c2f9161273b565b91612c3991611fe6565b612c439150612748565b612c4c90612778565b612c559161273b565b612c5e906127a8565b600c1c6040517f3cf979830000000000000000000000000000000000000000000000000000000081529485938493612c9a93916004860161283b565b03815a6000948591f1908115610f1257600090600092612cf4575b5015612cbe5750565b610de0906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611b66565b9050612d1391503d806000833e612d0b81836102e9565b8101906127c0565b509038612cb5565b5050505050565b50612d2f610abe86614930565b612b50565b50843b15612b49565b506000612b42565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156119375760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b359061033982611099565b359061033982611b3b565b9080601f830112156100e7578135612ddc81611868565b92612dea60405194856102e9565b81845260208085019260051b8201019283116100e757602001905b828210612e125750505090565b602080918335612e2181611099565b815201910190612e05565b60c0813603126100e757612e3e61033b565b90612e4881612daf565b8252612e566020820161103c565b6020830152612e6760408201612dba565b6040830152606081013567ffffffffffffffff81116100e757612e8d903690830161107e565b6060830152608081013567ffffffffffffffff81116100e757612eb39036908301612dc5565b608083015260a08101359067ffffffffffffffff82116100e757612ed991369101612dc5565b60a082015290565b60405160208101906000825260208152612efc6040826102e9565b51902090565b818110612f0d575050565b60008155600101612f02565b9190601f8111612f2857505050565b610339926000526020600020906020601f840160051c83019310612f54575b601f0160051c0190612f02565b9091508190612f47565b919091825167ffffffffffffffff81116102ac57612f8681612f808454611950565b84612f19565b6020601f8211600114612fe4578190612fd5939495600092612fd9575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612fa3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061301784600052602060002090565b9160005b8181106130715750958360019596971061303a575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613030565b9192602060018192868b01518155019401920161301b565b81519167ffffffffffffffff83116102ac576801000000000000000083116102ac576020908254848455808510613100575b500190600052602060002060005b8381106130d65750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016130c9565b613117908460005285846000209182019101612f02565b386130bb565b6003610650926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a0606084015261319f61316c60c08501600184016119a3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301611a45565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201611a45565b60405190610160820182811067ffffffffffffffff8211176102ac576040526060610140836000815260006020820152600060408201528280820152826080820152600060a08201528260c08201528260e082015282610100820152826101208201520152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146127095760010190565b90156119375790565b90821015611937570190565b906008820180921161270957565b906002820180921161270957565b906001820180921161270957565b906020820180921161270957565b9190820180921161270957565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff0000000000000000000000000000000000000000000000008116926008811061330a575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110613370575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b604051906133af826102b1565b60606080836000815282602082015282604082015282808201520152565b604080519091906133de83826102e9565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b82811061341557505050565b6020906134206133a2565b82828501015201613409565b6040519061343b6020836102e9565b600080835282815b82811061344f57505050565b60209061345a6133a2565b82828501015201613443565b9061346f6131d0565b5060258110613b32576134806131d0565b9160006134bf6134b96134938585613264565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613b0457506135046134f66134f06134ea6134e16001613279565b600188886132be565b906132d6565b60c01c90565b67ffffffffffffffff168552565b6135756135476135146001613279565b6135426135316134f06134ea61352985613279565b858b8b6132be565b67ffffffffffffffff166020890152565b613279565b6135426135646134f06134ea61355c85613279565b858a8a6132be565b67ffffffffffffffff166040880152565b6135986135926134b961349361358a85613237565b94888861326d565b60ff1690565b90846135a483836132b1565b11613ad75790816135cd6135c66135be846135d7966132b1565b8389896132be565b3691611047565b60608801526132b1565b83811015613aa9576135f46135926134b961349361358a85613237565b908461360083836132b1565b11613a7c57908161361a6135c66135be84613624966132b1565b60808801526132b1565b8361362e82613287565b11613a4f578061366361365861365261364c61355c61366896613287565b9061333c565b60f01c90565b61ffff1660a0880152565b613287565b83811015613a21576136856135926134b961349361358a85613237565b908461369183836132b1565b116139f45790816136ab6135c66135be846136b5966132b1565b60c08801526132b1565b838110156139c6576136d26135926134b961349361358a85613237565b90846136de83836132b1565b116139995790816136f86135c66135be84613702966132b1565b60e08801526132b1565b8361370c82613287565b1161396b5761ffff61373761373161365261364c61372986613287565b868a8a6132be565b92613287565b9116908461374583836132b1565b1161393e57908161375f6135c66135be8461376a966132b1565b6101008801526132b1565b908361377583613287565b11613911575061ffff61379b61373161365261364c61379386613287565b8689896132be565b9116806138a857506137ab61342c565b6101208501525b826137bc82613287565b116138795761ffff6137d961373161365261364c61379386613287565b911690836137e783836132b1565b1161384a576138086135c661381394838761380287836132b1565b926132be565b6101408601526132b1565b0361381b5790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b906138ce6138c66138b76133cd565b936101208801948552836132b1565b918585614c19565b909251926138dc829461192a565b52146137b2577fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008152600b600452602490fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008352600a600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526009600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526008600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526007600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526006600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526005600452602490fd5b507fb4205b4200000000000000000000000000000000000000000000000000000000815260048052602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526003600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526002600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526001600452602483fd5b7f789d326300000000000000000000000000000000000000000000000000000000825260ff16600452602490fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052610d926024906000600452565b60405190613b726020836102e9565b6000808352366020840137565b8015612709577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff168015612709577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9391613be2613b63565b93815180613f3a575b505050613bf890846152ad565b92909293613c486002613c27613c2d6003613c278b67ffffffffffffffff166000526004602052604060002090565b01611a95565b9867ffffffffffffffff166000526004602052604060002090565b613c73613c6e613c66613c5e87518651906132b1565b8a51906132b1565b8351906132b1565b6118ac565b9687956000805b8751821015613ccd5790613cc5600192613caa613c9a6114ba858d61193c565b91613ca481613237565b9c61193c565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018997613c7a565b9750509193969092945060005b8551811015613d0f5780613d09613cf66114ba6001948a61193c565b613caa613d028b613237565b9a8d61193c565b01613cda565b50919350919460005b8451811015613d4d5780613d47613d346114ba6001948961193c565b613caa613d408a613237565b998c61193c565b01613d18565b5093919592509360005b828110613ed1575b5050600090815b818110613e325750508152835160005b818110613d8557508452929190565b926000959495915b8351831015613e2357613da36114ba868861193c565b73ffffffffffffffffffffffffffffffffffffffff613dc86114716114ba878961193c565b911603613e1257613dd890613b7f565b90613df3613de96114ba848961193c565b613caa878961193c565b60ff8716613e02575b90613d8d565b95613e0c90613baa565b95613dfc565b9091613e1d90613237565b91613dfc565b91509260019095949501613d76565b613e3f6114ba828661193c565b73ffffffffffffffffffffffffffffffffffffffff81168015613ec757600090815b868110613e9b575b5050906001929115613e7e575b505b01613d66565b613e9590613caa613e8e87613237565b968861193c565b38613e76565b81613eac6114716114ba848c61193c565b14613eb957600101613e61565b506001915081905038613e69565b5050600190613e78565b613ee16114716114ba838761193c565b15613eee57600101613d57565b5091929360009591955b8351811015613f2d5780613f27613f146114ba6001948861193c565b613caa613f208b613237565b9a8961193c565b01613ef8565b5093929150933880613d5f565b909192945060018103613fec575060146060613f558461192a565b5101515103613fa95781613fa091613f7f610b676060613f77613bf89761192a565b5101516117fe565b91876080613f97613f8f8461192a565b51519361192a565b510151936150ab565b92903880613beb565b610de06060613fb78461192a565b5101516040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611b66565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6140216131d0565b5090565b9290612efc9173ffffffffffffffffffffffffffffffffffffffff61407367ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906103db565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b90604051916140b660c0846102e9565b60848352602083019060a03683375a612ee0811115614107576000916140dc839261270e565b82602083519301913090f1903d90608482116140fe575b6000908286523e9190565b608491506140f3565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b91908110156119375760051b0190565b3561065081611099565b90614188939291613bd8565b91939092614195826118ac565b9261419f836118ac565b94600091825b88518110156142a6576000805b8a888489828510614222575b5050505050156141d0576001016141a5565b6141e06114ba610d92928b61193c565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6114ba61425a9261425461424f8873ffffffffffffffffffffffffffffffffffffffff9761147196614162565b614172565b9561193c565b911614614269576001016141b2565b6001915061428861427e61424f838b8b614162565b613caa888c61193c565b61429b61429487613237565b968b61193c565b52388a8884896141be565b509097965094939291909460ff811690816000985b8a518a10156143785760005b8b8782108061436f575b156143625773ffffffffffffffffffffffffffffffffffffffff6143036114716114ba8f61425461424f888f8f614162565b9116146143185761431390613237565b6142c7565b939961432860019294939b613b7f565b9461434461433a61424f838b8b614162565b613caa8d8c61193c565b6143576143508c613237565b9b8b61193c565b525b019890916142bb565b5050919098600190614359565b508515156142d1565b9850925093959497509150816143a1575050508151810361439857509190565b80825283529190565b610d9292916143af9161273b565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b908160209103126100e7575161065081611099565b604051906103b2826102cd565b908160209103126100e7576040519061441d826102cd565b51815290565b90610650916020815260e06145186144e561444c855161010060208701526101208601906103db565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff166060860152606086015160808601526144b1608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103db565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103db565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103db565b3d15614577573d9061455d82610369565b9161456b60405193846102e9565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff610650949316815281602082015201906103db565b9092916145b46125ec565b506145d9611471606084016145c9815161540f565b51602080825183010191016143e3565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa948515610f12576000956148ff575b5073ffffffffffffffffffffffffffffffffffffffff85169485156148bc5761469981614986565b6148bc576146a9610abe826149b0565b6148bc575061477a916020916146bf888661549b565b956146c86143f8565b5080516146f3608086840151930151936146e061034a565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c08201526147426103a3565b60e0820152604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301614423565b03816000885af16000918161488b575b506147ce578461479861454c565b90610de06040519283927f9fe2f95a0000000000000000000000000000000000000000000000000000000084526004840161457c565b84909373ffffffffffffffffffffffffffffffffffffffff831603614821575b505050516148196147fd61035a565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b61482a9161549b565b908082108015614877575b61483f57836147ee565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b50614882818361273b565b83511415614835565b6148ae91925060203d6020116148b5575b6148a681836102e9565b810190614405565b903861478a565b503d61489c565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b61492291955060203d602011614929575b61491a81836102e9565b8101906143e3565b9338614671565b503d614910565b61495a7f85572ffb0000000000000000000000000000000000000000000000000000000082615671565b9081614974575b8161496a575090565b6106509150615611565b905061497f8161554b565b1590614961565b61495a7ff208a58f0000000000000000000000000000000000000000000000000000000082615671565b61495a7faff2afbf0000000000000000000000000000000000000000000000000000000082615671565b61495a7f1ef5498f0000000000000000000000000000000000000000000000000000000082615671565b61495a7f7909b5490000000000000000000000000000000000000000000000000000000082615671565b73ffffffffffffffffffffffffffffffffffffffff600154163303614a4f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614a878151846132b1565b928315614bb45760005b848110614a9f575050505050565b81811015614b9957614ab46114ba828661193c565b73ffffffffffffffffffffffffffffffffffffffff811680156114e157614ada83613295565b878110614aec57505050600101614a91565b84811015614b695773ffffffffffffffffffffffffffffffffffffffff614b166114ba838a61193c565b168214614b2557600101614ada565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614b946114ba614b8e888561273b565b8961193c565b614b16565b614baf6114ba614ba9848461273b565b8561193c565b614ab4565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b359060208110614bec575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b929192614c246133a2565b9382811015614fa257614c476134b9613493614c3f84613237565b93868661326d565b600160ff821603614f72575082614c5d826132a3565b11614f435780614c83614c7d614c75614c8a946132a3565b8387876132be565b90614bde565b86526132a3565b82811015614f1457614caf6135926134b9613493614ca785613237565b94878761326d565b83614cba82846132b1565b11614ee55781614cdb6135c6614cd384614ce5966132b1565b8388886132be565b60208801526132b1565b82811015614eb657614d026135926134b9613493614ca785613237565b83614d0d82846132b1565b11614e875781614d266135c6614cd384614d30966132b1565b60408801526132b1565b82811015614e5857614d4d6135926134b9613493614ca785613237565b83614d5882846132b1565b11614e2957816135cd6135c6614cd384614d71966132b1565b82614d7b82613287565b11614dfa5761ffff614d9861373161365261364c61379386613287565b91169183614da684846132b1565b11614dcb576135c6614dc191836106509661380287836132b1565b60808601526132b1565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b9080601f830112156100e7578151614fe881611868565b92614ff660405194856102e9565b81845260208085019260051b8201019283116100e757602001905b82821061501e5750505090565b60208091835161502d81611099565b815201910190615011565b906020828203126100e757815167ffffffffffffffff81116100e7576106509201614fd1565b919360a09367ffffffffffffffff610650979673ffffffffffffffffffffffffffffffffffffffff61ffff95168652166020850152604084015216606082015281608082015201906103db565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095929493929091906020838060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa928315610f1257600093615230575b50615153836149da565b61518d575b5050505050815115615168575090565b6106509150613c2760029167ffffffffffffffff166000526004602052604060002090565b600094965085916151e373ffffffffffffffffffffffffffffffffffffffff92604051998a97889687957f9f68f6730000000000000000000000000000000000000000000000000000000087526004870161505e565b0392165afa918215610f125760009261520b575b5061520182615757565b3880808080615158565b6152299192503d806000833e61522181836102e9565b810190615038565b90386151f7565b61524a91935060203d6020116149295761491a81836102e9565b9138615149565b90916060828403126100e757815167ffffffffffffffff81116100e7578361527a918401614fd1565b92602083015167ffffffffffffffff81116100e75760409161529d918501614fd1565b92015160ff811681036100e75790565b9190916152d161082e8267ffffffffffffffff166000526004602052604060002090565b90833b6152ee575b506060015191506152e8613b63565b90600090565b6152f784614a04565b156152d9576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015610f12576000809481926153a3575b5061537281615757565b61537b85615757565b805115801590615397575b61539057506152d9565b9392909150565b5060ff82161515615386565b91506153c29294503d8091833e6153ba81836102e9565b810190615251565b90939138615368565b6002548110156119375760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b908160209103126100e7575190565b60208151036154525761542b6020825183010160208301615400565b73ffffffffffffffffffffffffffffffffffffffff811190811561548f575b506154525750565b610de0906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906103db565b6104009150103861544a565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161551a575b50615516578261479861454c565b9150565b61553d91925060203d602011615544575b61553581836102e9565b810190615400565b9038615508565b503d61552b565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526155ab6044836102e9565b6179185a106155e7576020926000925191617530fa6000513d826155db575b50816155d4575090565b9050151590565b602011159150386155ca565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526155ab6044836102e9565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526155ab6044836102e9565b8060005260036020526040600020541560001461575157600254680100000000000000008110156102ac5760018101600255600060025482101561193757600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b50600090565b80519060005b82811061576957505050565b60018101808211612709575b838110615785575060010161575d565b73ffffffffffffffffffffffffffffffffffffffff6157a4838561193c565b51166157b66114716114ba848761193c565b146157c357600101615775565b610d926157d36114ba848661193c565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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
