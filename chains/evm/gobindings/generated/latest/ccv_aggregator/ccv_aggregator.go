// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccv_aggregator

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

type CCVAggregatorSourceChainConfig struct {
	Router           common.Address
	IsEnabled        bool
	OnRamp           []byte
	DefaultCCVs      []common.Address
	LaneMandatedCCVs []common.Address
}

type CCVAggregatorSourceChainConfigArgs struct {
	Router              common.Address
	SourceChainSelector uint64
	IsEnabled           bool
	OnRamp              []byte
	DefaultCCV          []common.Address
	LaneMandatedCCVs    []common.Address
}

type CCVAggregatorStaticConfig struct {
	LocalChainSelector   uint64
	GasForCallExactCheck uint16
	RmnRemote            common.Address
	TokenAdminRegistry   common.Address
}

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

var CCVAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enumMessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f61598938819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615735908161025482396080518181816101560152610ae0015260a0518181816101b90152610a63015260c0518181816101e101528181614553015261502b015260e05181818161017d0152612b220152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780635215505b146100c85780636b8be52c146100c357806379ba5097146100be5780637ce1552a146100b95780638da5cb5b146100b4578063d2b33733146100af578063e9d68a8e146100aa578063f054ac57146100a55763f2fde38b146100a057600080fd5b61170a565b6113b6565b6112d6565b611203565b6111b1565b611112565b610f41565b610897565b610713565b610546565b61041e565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610290565b828152826020820152826040820152015261025d60405161014b81610290565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ac57604052565b610261565b60a0810190811067ffffffffffffffff8211176102ac57604052565b6020810190811067ffffffffffffffff8211176102ac57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ac57604052565b6040519061033960a0836102e9565b565b6040519061033960c0836102e9565b60405190610339610100836102e9565b604051906103396040836102e9565b67ffffffffffffffff81116102ac57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906103b26020836102e9565b60008252565b60005b8381106103cb5750506000910152565b81810151838201526020016103bb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610417815180928187528780880191016103b8565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061045f81836102e9565b601782527f43435641676772656761746f7220312e372e302d6465760000000000000000006020830152519182916020835260208301906103db565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104e75750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104da565b9161053f60ff916105316040949796976060875260608701906104c9565b9085820360208701526104c9565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576105d66105a461059e61025d93369060040161049b565b90613346565b67ffffffffffffffff815116906105be60e08201516117fe565b60601c61ffff60a06101208401519301511692613ab1565b60409391935193849384610513565b6106509173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061063f61062d604085015160a0604086015260a08501906103db565b606085015184820360608601526104c9565b9201519060808184039101526104c9565b90565b6040810160408252825180915260206060830193019060005b8181106106f3575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106106a857505050505090565b90919293946020806106e4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516105e5565b97019301930191939290610699565b825167ffffffffffffffff1685526020948501949092019160010161066c565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461074e81611868565b9061075c60405192836102e9565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061078982611868565b0160005b81811061084f57505061079f816118ac565b9060005b8181106107bb57505061025d60405192839283610653565b806107f36107da6107cd6001946152de565b67ffffffffffffffff1690565b6107e4838761193c565b9067ffffffffffffffff169052565b61083361082e610814610806848861193c565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b611ab0565b61083d828761193c565b52610848818661193c565b50016107a3565b60209061085a611880565b8282870101520161078d565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576108e690369060040161049b565b60243567ffffffffffffffff81116100e757610906903690600401610866565b92909160443567ffffffffffffffff81116100e757610929903690600401610866565b93909261093c60015460ff9060a01c1690565b610f17576109a290610988740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b61099a6109958583613346565b613ef2565b933691611047565b6020815191012093610a4a60206109ef6109c76107cd875167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610f1257600091610ee3575b50610e9857610ac2610abe610ab4610814865167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b610e4d57602083015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603610e165750808603610de45760e083019384516014815103610daa5750610b73610b3e855167ffffffffffffffff1690565b6040860196610b55885167ffffffffffffffff1690565b610b6d610b6760c08a015193516117fe565b60601c90565b92613efe565b96610b92610b8b896000526005602052604060002090565b5460ff1690565b610b9b816110e6565b8015908115610d96575b5015610d2757610cce610cbf610cba96610806610c9e7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df298610c998d8f9967ffffffffffffffff9b610c6d91610ce09b610c37610c0c8f6000526005602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957fd2b337330000000000000000000000000000000000000000000000000000000060208801528b60248801611dd7565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b613fb9565b969015610d12576002998a916000526005602052604060002090565b611b77565b965167ffffffffffffffff1690565b91836040519485941697169583611fc8565b0390a4610d107fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b005b6003998a916000526005602052604060002090565b610d928787610d50610d41895167ffffffffffffffff1690565b915167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150610da3816110e6565b1438610ba5565b610de0906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611b66565b0390fd5b7fb5ace4f300000000000000000000000000000000000000000000000000000000600052600486905260245260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b610d92610e62845167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610d92610ead845167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610f05915060203d602011610f0b575b610efd81836102e9565b810190611b45565b38610a94565b503d610ef3565b611b5a565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303611000577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff8116036100e757565b35906103398261102a565b92919261105382610369565b9161106160405193846102e9565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e75781602061065093359101611047565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600411156110f057565b6110b7565b9060048210156110f05752565b60208101929161033991906110f5565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043561114d8161102a565b6024359061115a8261102a565b6044359167ffffffffffffffff83116100e75761117e61119193369060040161107e565b906064359261118c84611099565b613efe565b600052600560205261025d60ff6040600020541660405191829182611102565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760443567ffffffffffffffff81116100e757611291903690600401610866565b916064359267ffffffffffffffff84116100e7576112b6610d10943690600401610866565b9390926024359060040161286e565b9060206106509281815201906105e5565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff60043561131a8161102a565b611322611880565b5016600052600460205261025d60406000206113a5600360405192611346846102b1565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c16151560208501526040516113898161138281600186016119a3565b03826102e9565b604085015261139a60028201611a95565b606085015201611a95565b6080820152604051918291826112c5565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757611405903690600401610866565b9061140e614941565b6000905b82821061141b57005b61142e611429838584612c1b565b612cd8565b60208101916114486107cd845167ffffffffffffffff1690565b156116e05761148a611471611471845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b1580156116d3575b6114e15760808201949060005b8651805182101561150b576114716114ba836114d49361193c565b5173ffffffffffffffffffffffffffffffffffffffff1690565b156114e15760010161149f565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050939194929060009460a08701955b86518051821015611543576114716114ba836115369361193c565b156114e15760010161151b565b505094959095606082018051805180159182156116bd575b50506114e1576107cd6114ba9461167a6116b3946116708a6116666116a69761165d61161c60019f9c6115b27f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259e5189519061498c565b6115c76108148b5167ffffffffffffffff1690565b9e8f6115d66040840151151590565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b8d9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518d8c01612e0a565b5160028a01612f64565b5160038801612f64565b6116976116926107cd835167ffffffffffffffff1690565b6155e6565b505167ffffffffffffffff1690565b9260405191829182612ff8565b0390a20190611412565b6020012090506116cb612d8d565b14388061155b565b5060808201515115611492565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561175a81611099565b611762614941565b163381146117d457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611836575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b67ffffffffffffffff81116102ac5760051b60200190565b6040519061188d826102b1565b6060608083600081526000602082015282604082015282808201520152565b906118b682611868565b6118c360405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06118f18294611868565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156119375760200190565b6118fb565b80518210156119375760209160051b010190565b90600182811c92168015611999575b602083101461196a57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161195f565b600092918154916119b383611950565b8083529260018116908115611a0957506001146119cf57505050565b60009081526020812093945091925b8383106119ef575060209250010190565b6001816020929493945483858701015201910191906119de565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b906020825491828152019160005260206000209060005b818110611a695750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201611a5c565b90610339611aa99260405193848092611a45565b03836102e9565b9060036080604051611ac1816102b1565b611b37819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c1615156020840152604051611b028161138281600186016119a3565b6040840152604051611b1b816113828160028601611a45565b6060840152611b306040518096819301611a45565b03846102e9565b0152565b801515036100e757565b908160209103126100e7575161065081611b3b565b6040513d6000823e3d90fd5b9060206106509281815201906103db565b9060048110156110f05760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b838310611bda57505050505090565b9091929394602080611c6a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080611c59611c47611c358786015160a08987015260a08601906103db565b604086015185820360408701526103db565b606085015184820360608601526103db565b9201519060808184039101526103db565b97019301930191939290611bcb565b9160209082815201919060005b818110611c935750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff8735611cbc81611099565b168152019401929101611c86565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b848410611d815750505050505090565b909192939495602080611dc7837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018852611dc18b88611d09565b90611cca565b9801940194019294939190611d71565b9492936106509694611fa7611fba949360808952611e0260808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a0152610140611f74611f3e611f08611ed18d611e8d611e58606089015161016060e08501526101e08401906103db565b60808901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101008501526103db565b60a088015161ffff166101208301529060c088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b8d60e0870151906101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b6101008501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101808f01526103db565b6101208401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101a08e0152611bae565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016101c08b01526103db565b9260208801528683036040880152611c79565b926060818503910152611d59565b80611fd960409261065095946110f5565b81602082015201906103db565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611836575050565b356106508161102a565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b61ffff8116036100e757565b35610650816120c9565b91909160a0818403126100e7576120f461032a565b9281358452602082013567ffffffffffffffff81116100e7578161211991840161107e565b6020850152604082013567ffffffffffffffff81116100e7578161213e91840161107e565b6040850152606082013567ffffffffffffffff81116100e7578161216391840161107e565b6060850152608082013567ffffffffffffffff81116100e757612186920161107e565b6080830152565b92919061219981611868565b936121a760405195866102e9565b602085838152019160051b8101918383116100e75781905b8382106121cd575050505050565b813567ffffffffffffffff81116100e7576020916121ee87849387016120df565b8152019101906121bf565b90821015611937576122109160051b810190611fe6565b9091565b3590610339826120c9565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b83831061229a5750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61843603018112156100e757602061237d60019386839401908135815261236f61236461234961232e61231e88870187611d09565b60a08a88015260a0870191611cca565b61233b6040870187611d09565b908683036040880152611cca565b6123566060860186611d09565b908583036060870152611cca565b926080810190611d09565b916080818503910152611cca565b98019601949301919061228a565b906106509593949273ffffffffffffffffffffffffffffffffffffffff6125d992168352608060208401526123d4608084016123c68361103c565b67ffffffffffffffff169052565b6123f46123e36020830161103c565b67ffffffffffffffff1660a0850152565b6124146124036040830161103c565b67ffffffffffffffff1660c0850152565b6125a861259c61255d61251e6124e06124876124496124366060890189611d09565b61016060e08d01526101e08c0191611cca565b6124566080890189611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808c8403016101008d0152611cca565b6124a261249660a08901612214565b61ffff166101208b0152565b6124af60c0880188611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101408c0152611cca565b6124ed60e0870187611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8403016101608b0152611cca565b61252c610100860186611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80898403016101808a0152611cca565b61256b61012085018561221f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888403016101a0890152612272565b91610140810190611d09565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80858403016101c0860152611cca565b9360408201526060818503910152611cca565b604051906040820182811067ffffffffffffffff8211176102ac5760405260006020838281520152565b9061262082611868565b61262d60405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061265b8294611868565b019060005b82811061266c57505050565b6020906126776125ec565b82828501015201612660565b91908110156119375760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100e7570190565b90916060828403126100e75781516126da81611b3b565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e75781519161270983610369565b9161271760405193846102e9565b838352602084830101116100e75760409261273891602080850191016103b8565b92015190565b9093929193608082528051608083015267ffffffffffffffff60208201511660a083015260806127b5612781604084015160a060c08701526101208601906103db565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808683030160e08701526103db565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80838203016101008401526020808351928381520192019060005b8181106128365750505061ffff9094166020820152610339919060609062030d406040820152019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff16855260209081015181860152604090940193909201916001016127f4565b939194909294303303612bf1576128d36128bc9661289b610b6761289560e08a018a611fe6565b90612037565b9260a06128a78961206b565b856128cd8b6128c561012082019e8f83612075565b969092016120d5565b94369161218d565b9161408f565b97909160005b8351811015612989576128f56114716114716114ba848861193c565b9061290b612903828d61193c565b5187896121f9565b9290813b156100e7576000918a838d612954604051988996879586947f58bfa40a000000000000000000000000000000000000000000000000000000008652306004870161238b565b03925af1918215610f125760019261296e575b50016128d9565b8061297d6000612983936102e9565b806100dc565b38612967565b5094969350965050506129a661299f8285612075565b9050612616565b9460005b6129b48386612075565b9050811015612a285780612a0c6129d76001936129d1878a612075565b90612683565b86612a066129e860c08b018b611fe6565b91906129fe6129f68d61206b565b9536906120df565b923691611047565b906144bc565b612a16828a61193c565b52612a21818961193c565b50016129aa565b5094936101408401939150612a3d8483611fe6565b90501580612be9575b8015612be0575b8015612bce575b612bc757612b4b94612ae9600095612ad6612add612a94611471612a7a6108148a61206b565b5473ffffffffffffffffffffffffffffffffffffffff1690565b96612a9e8161206b565b93612ab7612aaf60c0840184611fe6565b929093611fe6565b949095612ac261032a565b998a5267ffffffffffffffff1660208a0152565b3691611047565b60408601523691611047565b60608301526080820152836040518096819582947f3cf979830000000000000000000000000000000000000000000000000000000084527f0000000000000000000000000000000000000000000000000000000000000000906004850161273e565b03925af1908115610f1257600090600092612ba0575b5015612b6a5750565b610de0906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611b66565b9050612bbf91503d806000833e612bb781836102e9565b8101906126c3565b509038612b61565b5050505050565b50612bdb610abe84614843565b612a54565b50823b15612a4d565b506000612a46565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156119375760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b359061033982611099565b359061033982611b3b565b9080601f830112156100e7578135612c8881611868565b92612c9660405194856102e9565b81845260208085019260051b8201019283116100e757602001905b828210612cbe5750505090565b602080918335612ccd81611099565b815201910190612cb1565b60c0813603126100e757612cea61033b565b90612cf481612c5b565b8252612d026020820161103c565b6020830152612d1360408201612c66565b6040830152606081013567ffffffffffffffff81116100e757612d39903690830161107e565b6060830152608081013567ffffffffffffffff81116100e757612d5f9036908301612c71565b608083015260a08101359067ffffffffffffffff82116100e757612d8591369101612c71565b60a082015290565b60405160208101906000825260208152612da86040826102e9565b51902090565b818110612db9575050565b60008155600101612dae565b9190601f8111612dd457505050565b610339926000526020600020906020601f840160051c83019310612e00575b601f0160051c0190612dae565b9091508190612df3565b919091825167ffffffffffffffff81116102ac57612e3281612e2c8454611950565b84612dc5565b6020601f8211600114612e90578190612e81939495600092612e85575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612e4f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690612ec384600052602060002090565b9160005b818110612f1d57509583600195969710612ee6575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080612edc565b9192602060018192868b015181550194019201612ec7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81519167ffffffffffffffff83116102ac576801000000000000000083116102ac576020908254848455808510612fdb575b500190600052602060002060005b838110612fb15750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501612fa4565b612ff2908460005285846000209182019101612dae565b38612f96565b6003610650926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a0606084015261307a61304760c08501600184016119a3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301611a45565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201611a45565b60405190610160820182811067ffffffffffffffff8211176102ac576040526060610140836000815260006020820152600060408201528280820152826080820152600060a08201528260c08201528260e082015282610100820152826101208201520152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461313f5760010190565b612f35565b90156119375790565b90821015611937570190565b906008820180921161313f57565b906002820180921161313f57565b906001820180921161313f57565b906020820180921161313f57565b9190820180921161313f57565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff000000000000000000000000000000000000000000000000811692600881106131ea575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110613250575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b6040519061328f826102b1565b60606080836000815282602082015282604082015282808201520152565b604080519091906132be83826102e9565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b8281106132f557505050565b602090613300613282565b828285010152016132e9565b6040519061331b6020836102e9565b600080835282815b82811061332f57505050565b60209061333a613282565b82828501015201613323565b9061334f6130ab565b5060258110613a0b576133606130ab565b91600061339f6133996133738585613144565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff8216036139dd57506133e46133d66133d06133ca6133c16001613159565b6001888861319e565b906131b6565b60c01c90565b67ffffffffffffffff168552565b6134556134276133f46001613159565b6134226134116133d06133ca61340985613159565b858b8b61319e565b67ffffffffffffffff166020890152565b613159565b6134226134446133d06133ca61343c85613159565b858a8a61319e565b67ffffffffffffffff166040880152565b61347861347261339961337361346a85613112565b94888861314d565b60ff1690565b90846134848383613191565b116139b05790816134a6612ad661349e846134b096613191565b83898961319e565b6060880152613191565b83811015613982576134cd61347261339961337361346a85613112565b90846134d98383613191565b116139555790816134f3612ad661349e846134fd96613191565b6080880152613191565b8361350782613167565b11613928578061353c61353161352b61352561343c61354196613167565b9061321c565b60f01c90565b61ffff1660a0880152565b613167565b838110156138fa5761355e61347261339961337361346a85613112565b908461356a8383613191565b116138cd579081613584612ad661349e8461358e96613191565b60c0880152613191565b8381101561389f576135ab61347261339961337361346a85613112565b90846135b78383613191565b116138725790816135d1612ad661349e846135db96613191565b60e0880152613191565b836135e582613167565b116138445761ffff61361061360a61352b61352561360286613167565b868a8a61319e565b92613167565b9116908461361e8383613191565b11613817579081613638612ad661349e8461364396613191565b610100880152613191565b908361364e83613167565b116137ea575061ffff61367461360a61352b61352561366c86613167565b86898961319e565b911680613781575061368461330c565b6101208501525b8261369582613167565b116137525761ffff6136b261360a61352b61352561366c86613167565b911690836136c08383613191565b11613723576136e1612ad66136ec9483876136db8783613191565b9261319e565b610140860152613191565b036136f45790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b906137a761379f6137906132ad565b93610120880194855283613191565b918585614b2c565b909251926137b5829461192a565b521461368b577fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008152600b600452602490fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008352600a600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526009600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526008600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526007600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526006600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526005600452602490fd5b507fb4205b4200000000000000000000000000000000000000000000000000000000815260048052602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526003600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526002600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526001600452602483fd5b7f789d326300000000000000000000000000000000000000000000000000000000825260ff16600452602490fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052610d926024906000600452565b60405190613a4b6020836102e9565b6000808352366020840137565b801561313f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff16801561313f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9391613abb613a3c565b93815180613e13575b505050613ad190846151c0565b92909293613b216002613b00613b066003613b008b67ffffffffffffffff166000526004602052604060002090565b01611a95565b9867ffffffffffffffff166000526004602052604060002090565b613b4c613b47613b3f613b378751865190613191565b8a5190613191565b835190613191565b6118ac565b9687956000805b8751821015613ba65790613b9e600192613b83613b736114ba858d61193c565b91613b7d81613112565b9c61193c565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018997613b53565b9750509193969092945060005b8551811015613be85780613be2613bcf6114ba6001948a61193c565b613b83613bdb8b613112565b9a8d61193c565b01613bb3565b50919350919460005b8451811015613c265780613c20613c0d6114ba6001948961193c565b613b83613c198a613112565b998c61193c565b01613bf1565b5093919592509360005b828110613daa575b5050600090815b818110613d0b5750508152835160005b818110613c5e57508452929190565b926000959495915b8351831015613cfc57613c7c6114ba868861193c565b73ffffffffffffffffffffffffffffffffffffffff613ca16114716114ba878961193c565b911603613ceb57613cb190613a58565b90613ccc613cc26114ba848961193c565b613b83878961193c565b60ff8716613cdb575b90613c66565b95613ce590613a83565b95613cd5565b9091613cf690613112565b91613cd5565b91509260019095949501613c4f565b613d186114ba828661193c565b73ffffffffffffffffffffffffffffffffffffffff81168015613da057600090815b868110613d74575b5050906001929115613d57575b505b01613c3f565b613d6e90613b83613d6787613112565b968861193c565b38613d4f565b81613d856114716114ba848c61193c565b14613d9257600101613d3a565b506001915081905038613d42565b5050600190613d51565b613dba6114716114ba838761193c565b15613dc757600101613c30565b5091929360009591955b8351811015613e065780613e00613ded6114ba6001948861193c565b613b83613df98b613112565b9a8961193c565b01613dd1565b5093929150933880613c38565b909192945060018103613ec5575060146060613e2e8461192a565b5101515103613e825781613e7991613e58610b676060613e50613ad19761192a565b5101516117fe565b91876080613e70613e688461192a565b51519361192a565b51015193614fbe565b92903880613ac4565b610de06060613e908461192a565b5101516040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611b66565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b613efa6130ab565b5090565b9290612da89173ffffffffffffffffffffffffffffffffffffffff613f4c67ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906103db565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd120820191821161313f57565b9190820391821161313f57565b9060405191613fc960c0846102e9565b60848352602083019060a03683375a612ee081111561401a57600091613fef8392613f7f565b82602083519301913090f1903d9060848211614011575b6000908286523e9190565b60849150614006565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b91908110156119375760051b0190565b3561065081611099565b9061409b939291613ab1565b919390926140a8826118ac565b926140b2836118ac565b94600091825b88518110156141b9576000805b8a888489828510614135575b5050505050156140e3576001016140b8565b6140f36114ba610d92928b61193c565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6114ba61416d926141676141628873ffffffffffffffffffffffffffffffffffffffff9761147196614075565b614085565b9561193c565b91161461417c576001016140c5565b6001915061419b614191614162838b8b614075565b613b83888c61193c565b6141ae6141a787613112565b968b61193c565b52388a8884896140d1565b509097965094939291909460ff811690816000985b8a518a101561428b5760005b8b87821080614282575b156142755773ffffffffffffffffffffffffffffffffffffffff6142166114716114ba8f614167614162888f8f614075565b91161461422b5761422690613112565b6141da565b939961423b60019294939b613a58565b9461425761424d614162838b8b614075565b613b838d8c61193c565b61426a6142638c613112565b9b8b61193c565b525b019890916141ce565b505091909860019061426c565b508515156141e4565b9850925093959497509150816142b457505050815181036142ab57509190565b80825283529190565b610d9292916142c291613fac565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b908160209103126100e7575161065081611099565b604051906103b2826102cd565b908160209103126100e75760405190614330826102cd565b51815290565b90610650916020815260e061442b6143f861435f855161010060208701526101208601906103db565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff166060860152606086015160808601526143c4608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103db565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103db565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103db565b3d1561448a573d9061447082610369565b9161447e60405193846102e9565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff610650949316815281602082015201906103db565b9092916144c76125ec565b506144ec611471606084016144dc8151615322565b51602080825183010191016142f6565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa948515610f1257600095614812575b5073ffffffffffffffffffffffffffffffffffffffff85169485156147cf576145ac81614899565b6147cf576145bc610abe826148c3565b6147cf575061468d916020916145d288866153ae565b956145db61430b565b508051614606608086840151930151936145f361034a565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c08201526146556103a3565b60e0820152604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301614336565b03816000885af16000918161479e575b506146e157846146ab61445f565b90610de06040519283927f9fe2f95a0000000000000000000000000000000000000000000000000000000084526004840161448f565b84909373ffffffffffffffffffffffffffffffffffffffff831603614734575b5050505161472c61471061035a565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b61473d916153ae565b90808210801561478a575b6147525783614701565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b506147958183613fac565b83511415614748565b6147c191925060203d6020116147c8575b6147b981836102e9565b810190614318565b903861469d565b503d6147af565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b61483591955060203d60201161483c575b61482d81836102e9565b8101906142f6565b9338614584565b503d614823565b61486d7f85572ffb0000000000000000000000000000000000000000000000000000000082615584565b9081614887575b8161487d575090565b6106509150615524565b90506148928161545e565b1590614874565b61486d7ff208a58f0000000000000000000000000000000000000000000000000000000082615584565b61486d7faff2afbf0000000000000000000000000000000000000000000000000000000082615584565b61486d7f0e8b773f0000000000000000000000000000000000000000000000000000000082615584565b61486d7f7909b5490000000000000000000000000000000000000000000000000000000082615584565b73ffffffffffffffffffffffffffffffffffffffff60015416330361496257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80519161499a815184613191565b928315614ac75760005b8481106149b2575050505050565b81811015614aac576149c76114ba828661193c565b73ffffffffffffffffffffffffffffffffffffffff811680156114e1576149ed83613175565b8781106149ff575050506001016149a4565b84811015614a7c5773ffffffffffffffffffffffffffffffffffffffff614a296114ba838a61193c565b168214614a38576001016149ed565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614aa76114ba614aa18885613fac565b8961193c565b614a29565b614ac26114ba614abc8484613fac565b8561193c565b6149c7565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b359060208110614aff575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b929192614b37613282565b9382811015614eb557614b5a613399613373614b5284613112565b93868661314d565b600160ff821603614e85575082614b7082613183565b11614e565780614b96614b90614b88614b9d94613183565b83878761319e565b90614af1565b8652613183565b82811015614e2757614bc2613472613399613373614bba85613112565b94878761314d565b83614bcd8284613191565b11614df85781614bee612ad6614be684614bf896613191565b83888861319e565b6020880152613191565b82811015614dc957614c15613472613399613373614bba85613112565b83614c208284613191565b11614d9a5781614c39612ad6614be684614c4396613191565b6040880152613191565b82811015614d6b57614c60613472613399613373614bba85613112565b83614c6b8284613191565b11614d3c57816134a6612ad6614be684614c8496613191565b82614c8e82613167565b11614d0d5761ffff614cab61360a61352b61352561366c86613167565b91169183614cb98484613191565b11614cde57612ad6614cd49183610650966136db8783613191565b6080860152613191565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b9080601f830112156100e7578151614efb81611868565b92614f0960405194856102e9565b81845260208085019260051b8201019283116100e757602001905b828210614f315750505090565b602080918351614f4081611099565b815201910190614f24565b906020828203126100e757815167ffffffffffffffff81116100e7576106509201614ee4565b919360a09367ffffffffffffffff610650979673ffffffffffffffffffffffffffffffffffffffff61ffff95168652166020850152604084015216606082015281608082015201906103db565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095929493929091906020838060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa928315610f1257600093615143575b50615066836148ed565b6150a0575b505050505081511561507b575090565b6106509150613b0060029167ffffffffffffffff166000526004602052604060002090565b600094965085916150f673ffffffffffffffffffffffffffffffffffffffff92604051998a97889687957f9f68f67300000000000000000000000000000000000000000000000000000000875260048701614f71565b0392165afa918215610f125760009261511e575b506151148261566a565b388080808061506b565b61513c9192503d806000833e61513481836102e9565b810190614f4b565b903861510a565b61515d91935060203d60201161483c5761482d81836102e9565b913861505c565b90916060828403126100e757815167ffffffffffffffff81116100e7578361518d918401614ee4565b92602083015167ffffffffffffffff81116100e7576040916151b0918501614ee4565b92015160ff811681036100e75790565b9190916151e461082e8267ffffffffffffffff166000526004602052604060002090565b90833b615201575b506060015191506151fb613a3c565b90600090565b61520a84614917565b156151ec576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015610f12576000809481926152b6575b506152858161566a565b61528e8561566a565b8051158015906152aa575b6152a357506151ec565b9392909150565b5060ff82161515615299565b91506152d59294503d8091833e6152cd81836102e9565b810190615164565b9093913861527b565b6002548110156119375760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b908160209103126100e7575190565b60208151036153655761533e6020825183010160208301615313565b73ffffffffffffffffffffffffffffffffffffffff81119081156153a2575b506153655750565b610de0906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906103db565b6104009150103861535d565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161542d575b5061542957826146ab61445f565b9150565b61545091925060203d602011615457575b61544881836102e9565b810190615313565b903861541b565b503d61543e565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526154be6044836102e9565b6179185a106154fa576020926000925191617530fa6000513d826154ee575b50816154e7575090565b9050151590565b602011159150386154dd565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526154be6044836102e9565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526154be6044836102e9565b8060005260036020526040600020541560001461566457600254680100000000000000008110156102ac5760018101600255600060025482101561193757600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b50600090565b80519060005b82811061567c57505050565b6001810180821161313f575b8381106156985750600101615670565b73ffffffffffffffffffffffffffffffffffffffff6156b7838561193c565b51166156c96114716114ba848761193c565b146156d657600101615688565b610d926156e66114ba848661193c565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
}

var CCVAggregatorABI = CCVAggregatorMetaData.ABI

var CCVAggregatorBin = CCVAggregatorMetaData.Bin

func DeployCCVAggregator(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig CCVAggregatorStaticConfig) (common.Address, *types.Transaction, *CCVAggregator, error) {
	parsed, err := CCVAggregatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCVAggregatorBin), backend, staticConfig)
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

func (_CCVAggregator *CCVAggregatorCaller) GetCCVsForMessage(opts *bind.CallOpts, encodedMessage []byte) (GetCCVsForMessage,

	error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "getCCVsForMessage", encodedMessage)

	outstruct := new(GetCCVsForMessage)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequiredCCVs = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OptionalCCVs = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Threshold = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

func (_CCVAggregator *CCVAggregatorSession) GetCCVsForMessage(encodedMessage []byte) (GetCCVsForMessage,

	error) {
	return _CCVAggregator.Contract.GetCCVsForMessage(&_CCVAggregator.CallOpts, encodedMessage)
}

func (_CCVAggregator *CCVAggregatorCallerSession) GetCCVsForMessage(encodedMessage []byte) (GetCCVsForMessage,

	error) {
	return _CCVAggregator.Contract.GetCCVsForMessage(&_CCVAggregator.CallOpts, encodedMessage)
}

func (_CCVAggregator *CCVAggregatorCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber, sender, receiver)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_CCVAggregator *CCVAggregatorSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error) {
	return _CCVAggregator.Contract.GetExecutionState(&_CCVAggregator.CallOpts, sourceChainSelector, sequenceNumber, sender, receiver)
}

func (_CCVAggregator *CCVAggregatorCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error) {
	return _CCVAggregator.Contract.GetExecutionState(&_CCVAggregator.CallOpts, sourceChainSelector, sequenceNumber, sender, receiver)
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

func (_CCVAggregator *CCVAggregatorTransactor) Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "execute", encodedMessage, ccvs, ccvData)
}

func (_CCVAggregator *CCVAggregatorSession) Execute(encodedMessage []byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _CCVAggregator.Contract.Execute(&_CCVAggregator.TransactOpts, encodedMessage, ccvs, ccvData)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) Execute(encodedMessage []byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _CCVAggregator.Contract.Execute(&_CCVAggregator.TransactOpts, encodedMessage, ccvs, ccvData)
}

func (_CCVAggregator *CCVAggregatorTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "executeSingleMessage", message, messageId, ccvs, ccvData)
}

func (_CCVAggregator *CCVAggregatorSession) ExecuteSingleMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _CCVAggregator.Contract.ExecuteSingleMessage(&_CCVAggregator.TransactOpts, message, messageId, ccvs, ccvData)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) ExecuteSingleMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error) {
	return _CCVAggregator.Contract.ExecuteSingleMessage(&_CCVAggregator.TransactOpts, message, messageId, ccvs, ccvData)
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

type GetCCVsForMessage struct {
	RequiredCCVs []common.Address
	OptionalCCVs []common.Address
	Threshold    uint8
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
	return common.HexToHash("0x04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e25")
}

func (CCVAggregatorStaticConfigSet) Topic() common.Hash {
	return common.HexToHash("0x4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e4950")
}

func (_CCVAggregator *CCVAggregator) Address() common.Address {
	return _CCVAggregator.address
}

type CCVAggregatorInterface interface {
	GetAllSourceChainConfigs(opts *bind.CallOpts) ([]uint64, []CCVAggregatorSourceChainConfig, error)

	GetCCVsForMessage(opts *bind.CallOpts, encodedMessage []byte) (GetCCVsForMessage,

		error)

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (CCVAggregatorSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (CCVAggregatorStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []CCVAggregatorSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, ccvData [][]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

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

	Address() common.Address
}
