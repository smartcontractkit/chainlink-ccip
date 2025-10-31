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
	GasLimit            uint32
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f615c0138819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a16040516159ad908161025482396080518181816101560152610bb6015260a0518181816101b90152610b10015260c0518181816101e1015281816146070152615338015260e05181818161017d015261258f0152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780634272068d146100c85780635215505b146100c35780636b8be52c146100be57806379ba5097146100b95780637ce1552a146100b45780638da5cb5b146100af578063e9d68a8e146100aa578063f054ac57146100a55763f2fde38b146100a057600080fd5b611754565b611400565b611320565b6112bd565b61121e565b61104d565b61095c565b610809565b610617565b610546565b61041e565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610290565b828152826020820152826040820152015261025d60405161014b81610290565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ac57604052565b610261565b60a0810190811067ffffffffffffffff8211176102ac57604052565b6020810190811067ffffffffffffffff8211176102ac57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ac57604052565b6040519061033960c0836102e9565b565b6040519061033960a0836102e9565b60405190610339610100836102e9565b604051906103396040836102e9565b67ffffffffffffffff81116102ac57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906103b26020836102e9565b60008252565b60005b8381106103cb5750506000910152565b81810151838201526020016103bb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610417815180928187528780880191016103b8565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061045f81836102e9565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906103db565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104e75750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104da565b9161053f60ff916105316040949796976060875260608701906104c9565b9085820360208701526104c9565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576105d76105a461059e61025d93369060040161049b565b906134d2565b67ffffffffffffffff815116906105bf610100820151611848565b60601c61ffff60a06101408401519301511692613cb2565b60409391935193849384610513565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760443567ffffffffffffffff81116100e7576106a59036906004016105e6565b916064359267ffffffffffffffff84116100e7576106ca6106d99436906004016105e6565b939092602435906004016122cb565b005b6107469173ffffffffffffffffffffffffffffffffffffffff82511681526020820151151560208201526080610735610723604085015160a0604086015260a08501906103db565b606085015184820360608601526104c9565b9201519060808184039101526104c9565b90565b6040810160408252825180915260206060830193019060005b8181106107e9575050506020818303910152815180825260208201916020808360051b8301019401926000915b83831061079e57505050505090565b90919293946020806107da837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516106db565b9701930193019193929061078f565b825167ffffffffffffffff16855260209485019490920191600101610762565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757600254610844816119ab565b9061085260405192836102e9565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061087f826119ab565b0160005b818110610945575050610895816126c5565b9060005b8181106108b157505061025d60405192839283610749565b806108e96108d06108c3600194615829565b67ffffffffffffffff1690565b6108da8387611b43565b9067ffffffffffffffff169052565b61092961092461090a6108fc8488611b43565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b612888565b6109338287611b43565b5261093e8186611b43565b5001610899565b602090610950612699565b82828701015201610883565b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576109ab90369060040161049b565b60243567ffffffffffffffff81116100e7576109cb9036906004016105e6565b92909160443567ffffffffffffffff81116100e7576109ee9036906004016105e6565b9390926109fd60055460ff1690565b61102357610a4f90610a3560017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006005541617600555565b610a47610a4285836134d2565b614a5a565b933691611153565b6020815191012093610af76020610a9c610a746108c3875167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561101e57600091610fef575b50610fa457610b5b61090a845167ffffffffffffffff1690565b8054610b6f9060a01c60ff161590565b1590565b610f5957606084016001815160208151910120920191610b8e83612809565b6020815191012003610f22575050602083015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603610eeb5750808603610eb95761010083019384516014815103610e7f5750610c4a610c15855167ffffffffffffffff1690565b6040860196610c2c885167ffffffffffffffff1690565b610c44610c3e60e08a01519351611848565b60601c90565b92614a66565b96610c69610c62896000526006602052604060002090565b5460ff1690565b610c72816111f2565b8015908115610e6b575b5015610dfc57610da5610d96610d91966108fc610d757f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df298610d708d8f9967ffffffffffffffff9b610d4491610db79b610d0e610ce38f6000526006602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957f4272068d0000000000000000000000000000000000000000000000000000000060208801528b60248801612b33565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b614ae7565b969015610de7576002998a916000526006602052604060002090565b61294d565b965167ffffffffffffffff1690565b91836040519485941697169583612d36565b0390a46106d97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060055416600555565b6003998a916000526006602052604060002090565b610e678787610e25610e16895167ffffffffffffffff1690565b915167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150610e78816111f2565b1438610c7c565b610eb5906040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352600483016122ba565b0390fd5b7fb5ace4f300000000000000000000000000000000000000000000000000000000600052600486905260245260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5190610eb56040519283927f2130095800000000000000000000000000000000000000000000000000000000845260048401612928565b610e67610f6e855167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610e67610fb9845167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611011915060203d602011611017575b61100981836102e9565b810190612913565b38610b41565b503d610fff565b612010565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff8116330361110c577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff8116036100e757565b359061033982611136565b92919261115f82610369565b9161116d60405193846102e9565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e75781602061074693359101611153565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600411156111fc57565b6111c3565b9060048210156111fc5752565b6020810192916103399190611201565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043561125981611136565b6024359061126682611136565b6044359167ffffffffffffffff83116100e75761128a61129d93369060040161118a565b9060643592611298846111a5565b614a66565b600052600660205261025d60ff604060002054166040519182918261120e565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b9060206107469281815201906106db565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff60043561136481611136565b61136c612699565b5016600052600460205261025d60406000206113ef600360405192611390846102b1565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c16151560208501526040516113d3816113cc8160018601612767565b03826102e9565b60408501526113e460028201612874565b606085015201612874565b60808201526040519182918261130f565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e75761144f9036906004016105e6565b90611458614ba3565b6000905b82821061146557005b6114786114738385846120b3565b612dd1565b60208101916114926108c3845167ffffffffffffffff1690565b1561172a576114d46114bb6114bb845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b15801561171d575b61152b5760808201949060005b86518051821015611555576114bb6115048361151e93611b43565b5173ffffffffffffffffffffffffffffffffffffffff1690565b1561152b576001016114e9565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050939194929060009460a08701955b8651805182101561158d576114bb6115048361158093611b43565b1561152b57600101611565565b50509495909560608201805180518015918215611707575b505061152b576108c3611504946116c46116fd946116ba8a6116b06116f0976116a761166660019f9c6115fc7f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259e51895190614bee565b61161161090a8b5167ffffffffffffffff1690565b9e8f6116206040840151151590565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b8d9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518d8c01612f03565b5160028a0161305d565b516003880161305d565b6116e16116dc6108c3835167ffffffffffffffff1690565b61585e565b505167ffffffffffffffff1690565b92604051918291826130f1565b0390a2019061145c565b602001209050611715612e86565b1438806115a5565b50608082015151156114dc565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff6004356117a4816111a5565b6117ac614ba3565b1633811461181e57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611880575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611880575050565b3561074681611136565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b61ffff8116036100e757565b3561074681611995565b67ffffffffffffffff81116102ac5760051b60200190565b91909160c0818403126100e7576119d861032a565b9281358452602082013567ffffffffffffffff81116100e757816119fd91840161118a565b6020850152604082013567ffffffffffffffff81116100e75781611a2291840161118a565b6040850152606082013567ffffffffffffffff81116100e75781611a4791840161118a565b6060850152608082013567ffffffffffffffff81116100e75781611a6c91840161118a565b608085015260a082013567ffffffffffffffff81116100e757611a8f920161118a565b60a0830152565b929190611aa2816119ab565b93611ab060405195866102e9565b602085838152019160051b8101918383116100e75781905b838210611ad6575050505050565b813567ffffffffffffffff81116100e757602091611af787849387016119c3565b815201910190611ac8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805115611b3e5760200190565b611b02565b8051821015611b3e5760209160051b010190565b90821015611b3e57611b6e9160051b8101906118b2565b9091565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b359061033982611995565b63ffffffff8116036100e757565b359061033982611c0c565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310611ca05750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e7576020611da1600193868394019081358152611d93611d88611d6d611d52611d37611d2789880188611b72565b60c08b89015260c0880191611bc2565b611d446040880188611b72565b908783036040890152611bc2565b611d5f6060870187611b72565b908683036060880152611bc2565b611d7a6080860186611b72565b908583036080870152611bc2565b9260a0810190611b72565b9160a0818503910152611bc2565b980196019493019190611c90565b611ffd6107469593949260608352611ddb60608401611dcd83611148565b67ffffffffffffffff169052565b611dfb611dea60208301611148565b67ffffffffffffffff166080850152565b611e1b611e0a60408301611148565b67ffffffffffffffff1660a0850152565b611fcc611fc0611f81611f42611f03611e8d611e50611e3d6060890189611b72565b61018060c08d01526101e08c0191611bc2565b611e5d6080890189611b72565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c84030160e08d0152611bc2565b611ea8611e9c60a08901611c01565b61ffff166101008b0152565b611ec5611eb760c08901611c1a565b63ffffffff166101208b0152565b611ed260e0880188611b72565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101408c0152611bc2565b611f11610100870187611b72565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101608b0152611bc2565b611f50610120860186611b72565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101808a0152611bc2565b611f8f610140850185611c25565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101a0890152611c78565b91610160810190611b72565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0858403016101c0860152611bc2565b9360208201526040818503910152611bc2565b6040513d6000823e3d90fd5b604051906040820182811067ffffffffffffffff8211176102ac5760405260006020838281520152565b90612050826119ab565b61205d60405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061208b82946119ab565b019060005b82811061209c57505050565b6020906120a761201c565b82828501015201612090565b9190811015611b3e5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b3561074681611c0c565b801515036100e757565b90916060828403126100e757815161211e816120fd565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e75781519161214d83610369565b9161215b60405193846102e9565b838352602084830101116100e75760409261217c91602080850191016103b8565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806121f96121c5604084015160a060c08801526101208701906103db565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103db565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106122825750505061ffff909516602083015261033992916060916122669063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101612238565b9060206107469281815201906103db565b9391959490929530330361266f576123326122f6610c3e6122f06101008901896118b2565b90611903565b9761230087611937565b9260a06101408901948b61232c8b61232461231b8a83611941565b969092016119a1565b943691611a96565b91614147565b97909160005b83518110156123e7576123546114bb6114bb6115048488611b43565b9061236a612362828d611b43565b518789611b57565b9290813b156100e7576000918a838d6123b2604051988996879586947f7d61129300000000000000000000000000000000000000000000000000000000865260048601611daf565b03925af191821561101e576001926123cc575b5001612338565b806123db60006123e1936102e9565b806100dc565b386123c5565b5094975095925050506124046123fd8385611941565b9050612046565b9360005b6124128486611941565b9050811015612485578061246961243560019361242f888a611941565b906120b3565b61246361244560e08a018a6118b2565b919061245b6124538c611937565b9436906119c3565b923691611153565b90614574565b6124738289611b43565b5261247e8188611b43565b5001612408565b5091509361016083019061249982856118b2565b90501580612656575b801561264d575b801561263b575b61263457600061255f60c0866125b89861254f6124f06114bb6124d661090a899d611937565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9761253c61254361250086611937565b9261251a61251160e08901896118b2565b919092896118b2565b949095602061252761033b565b9e8f908152019067ffffffffffffffff169052565b3691611153565b60408a01523691611153565b60608701526080860152016120f3565b91604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601612182565b03925af190811561101e5760009060009261260d575b50156125d75750565b610eb5906040519182917f0a8d6e8c000000000000000000000000000000000000000000000000000000008352600483016122ba565b905061262c91503d806000833e61262481836102e9565b810190612107565b5090386125ce565b5050505050565b50612648610b6b8461495c565b6124b0565b50823b156124a9565b5063ffffffff61266860c086016120f3565b16156124a2565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906126a6826102b1565b6060608083600081526000602082015282604082015282808201520152565b906126cf826119ab565b6126dc60405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061270a82946119ab565b0190602036910137565b90600182811c9216801561275d575b602083101461272e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612723565b6000929181549161277783612714565b80835292600181169081156127cd575060011461279357505050565b60009081526020812093945091925b8383106127b3575060209250010190565b6001816020929493945483858701015201910191906127a2565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b9061033961281d9260405193848092612767565b03836102e9565b906020825491828152019160005260206000209060005b8181106128485750505090565b825473ffffffffffffffffffffffffffffffffffffffff1684526020909301926001928301920161283b565b9061033961281d9260405193848092612824565b9060036080604051612899816102b1565b61290f819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c16151560208401526040516128da816113cc8160018601612767565b60408401526040516128f3816113cc8160028601612824565b60608401526129086040518096819301612824565b03846102e9565b0152565b908160209103126100e75751610746816120fd565b909161293f61074693604084526040840190612767565b9160208184039101526103db565b9060048110156111fc5760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b8383106129b057505050505090565b9091929394602080612a55837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a0612a44612a32612a20612a0e8887015160c08a88015260c08701906103db565b604087015186820360408801526103db565b606086015185820360608701526103db565b608085015184820360808601526103db565b9201519060a08184039101526103db565b970193019301919392906129a1565b9160209082815201919060005b818110612a7e5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff8735612aa7816111a5565b168152019401929101612a71565b90602083828152019260208260051b82010193836000925b848410612add5750505050505090565b909192939495602080612b23837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018852612b1d8b88611b72565b90611bc2565b9801940194019294939190612acd565b9492936107469694612d15612d28949360808952612b5e60808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a0152610160612ce2612cac612c76612c3e8d612be9612bb4606089015161018060e08501526102008401906103db565b60808901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101008501526103db565b60a088015161ffff166101208301529060c088015163ffffffff1661014082015260e088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b8d610100870151906101807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b6101208501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101a08f01526103db565b6101408401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101c08e0152612984565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016101e08b01526103db565b9260208801528683036040880152612a64565b926060818503910152612ab5565b80612d476040926107469594611201565b81602082015201906103db565b3590610339826111a5565b3590610339826120fd565b9080601f830112156100e7578135612d81816119ab565b92612d8f60405194856102e9565b81845260208085019260051b8201019283116100e757602001905b828210612db75750505090565b602080918335612dc6816111a5565b815201910190612daa565b60c0813603126100e757612de361032a565b90612ded81612d54565b8252612dfb60208201611148565b6020830152612e0c60408201612d5f565b6040830152606081013567ffffffffffffffff81116100e757612e32903690830161118a565b6060830152608081013567ffffffffffffffff81116100e757612e589036908301612d6a565b608083015260a08101359067ffffffffffffffff82116100e757612e7e91369101612d6a565b60a082015290565b60405160208101906000825260208152612ea16040826102e9565b51902090565b818110612eb2575050565b60008155600101612ea7565b9190601f8111612ecd57505050565b610339926000526020600020906020601f840160051c83019310612ef9575b601f0160051c0190612ea7565b9091508190612eec565b919091825167ffffffffffffffff81116102ac57612f2b81612f258454612714565b84612ebe565b6020601f8211600114612f89578190612f7a939495600092612f7e575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612f48565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690612fbc84600052602060002090565b9160005b81811061301657509583600195969710612fdf575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080612fd5565b9192602060018192868b015181550194019201612fc0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81519167ffffffffffffffff83116102ac576801000000000000000083116102ac5760209082548484558085106130d4575b500190600052602060002060005b8381106130aa5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161309d565b6130eb908460005285846000209182019101612ea7565b3861308f565b6003610746926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a0606084015261317361314060c0850160018401612767565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301612824565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201612824565b60405190610180820182811067ffffffffffffffff8211176102ac576040526060610160836000815260006020820152600060408201528280820152826080820152600060a0820152600060c08201528260e08201528261010082015282610120820152826101408201520152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146132405760010190565b61302e565b9015611b3e5790565b90821015611b3e570190565b906008820180921161324057565b906002820180921161324057565b906004820180921161324057565b906001820180921161324057565b906020820180921161324057565b9190820180921161324057565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff000000000000000000000000000000000000000000000000811692600881106132f9575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffff0000000000000000000000000000000000000000000000000000000000008116926002811061335f575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106133c5575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b6040519060c0820182811067ffffffffffffffff8211176102ac57604052606060a0836000815282602082015282604082015282808201528260808201520152565b6040805190919061344a83826102e9565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b82811061348157505050565b60209061348c6133f7565b82828501015201613475565b604051906134a76020836102e9565b600080835282815b8281106134bb57505050565b6020906134c66133f7565b828285010152016134af565b906134db6131a4565b5060258110613c0c576134ec6131a4565b91600061352b6135256134ff8585613245565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613bde575061357061356261355c61355661354d600161325a565b600188886132ad565b906132c5565b60c01c90565b67ffffffffffffffff168552565b6135e16135b3613580600161325a565b6135ae61359d61355c6135566135958561325a565b858b8b6132ad565b67ffffffffffffffff166020890152565b61325a565b6135ae6135d061355c6135566135c88561325a565b858a8a6132ad565b67ffffffffffffffff166040880152565b6136046135fe6135256134ff6135f685613213565b94888861324e565b60ff1690565b908461361083836132a0565b11613bb157908161363261253c61362a8461363c966132a0565b8389896132ad565b60608801526132a0565b83811015613b83576136596135fe6135256134ff6135f685613213565b908461366583836132a0565b11613b5657908161367f61253c61362a84613689966132a0565b60808801526132a0565b8361369382613268565b11613b2957806136c86136bd6136b76136b16135c86136cd96613268565b9061332b565b60f01c90565b61ffff1660a0880152565b613268565b836136d782613276565b11613afb578061370e6137016136fb6136f56135c861371396613276565b90613391565b60e01c90565b63ffffffff1660c0880152565b613276565b83811015613acd576137306135fe6135256134ff6135f685613213565b908461373c83836132a0565b11613aa057908161375661253c61362a84613760966132a0565b60e08801526132a0565b83811015613a725761377d6135fe6135256134ff6135f685613213565b908461378983836132a0565b11613a455790816137a361253c61362a846137ae966132a0565b6101008801526132a0565b836137b882613268565b11613a175761ffff6137e36137dd6136b76136b16137d586613268565b868a8a6132ad565b92613268565b911690846137f183836132a0565b116139ea57908161380b61253c61362a84613816966132a0565b6101208801526132a0565b908361382183613268565b116139bd575061ffff6138476137dd6136b76136b161383f86613268565b8689896132ad565b9116806139545750613857613498565b6101408501525b8261386882613268565b116139255761ffff6138856137dd6136b76136b161383f86613268565b9116908361389383836132a0565b116138f6576138b461253c6138bf9483876138ae87836132a0565b926132ad565b6101608601526132a0565b036138c75790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b9061397a613972613963613439565b936101408801948552836132a0565b918585614d8e565b909251926139888294611b31565b521461385e577fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008152600c600452602490fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008352600b600452602483fd5b507fb4205b42000000000000000000000000000000000000000000000000000000008152600a600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526009600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526008600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526007600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526006600452602490fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526005600452602490fd5b507fb4205b4200000000000000000000000000000000000000000000000000000000815260048052602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526003600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526002600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526001600452602483fd5b7f789d326300000000000000000000000000000000000000000000000000000000825260ff16600452602490fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052610e676024906000600452565b60405190613c4c6020836102e9565b6000808352366020840137565b8015613240577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff168015613240577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9391613cbc613c3d565b93815180614014575b505050613cd290846154cd565b92909293613d226002613d01613d076003613d018b67ffffffffffffffff166000526004602052604060002090565b01612874565b9867ffffffffffffffff166000526004602052604060002090565b613d4d613d48613d40613d3887518651906132a0565b8a51906132a0565b8351906132a0565b6126c5565b9687956000805b8751821015613da75790613d9f600192613d84613d74611504858d611b43565b91613d7e81613213565b9c611b43565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018997613d54565b9750509193969092945060005b8551811015613de95780613de3613dd06115046001948a611b43565b613d84613ddc8b613213565b9a8d611b43565b01613db4565b50919350919460005b8451811015613e275780613e21613e0e61150460019489611b43565b613d84613e1a8a613213565b998c611b43565b01613df2565b5093919592509360005b828110613fab575b5050600090815b818110613f0c5750508152835160005b818110613e5f57508452929190565b926000959495915b8351831015613efd57613e7d6115048688611b43565b73ffffffffffffffffffffffffffffffffffffffff613ea26114bb6115048789611b43565b911603613eec57613eb290613c59565b90613ecd613ec36115048489611b43565b613d848789611b43565b60ff8716613edc575b90613e67565b95613ee690613c84565b95613ed6565b9091613ef790613213565b91613ed6565b91509260019095949501613e50565b613f196115048286611b43565b73ffffffffffffffffffffffffffffffffffffffff81168015613fa157600090815b868110613f75575b5050906001929115613f58575b505b01613e40565b613f6f90613d84613f6887613213565b9688611b43565b38613f50565b81613f866114bb611504848c611b43565b14613f9357600101613f3b565b506001915081905038613f43565b5050600190613f52565b613fbb6114bb6115048387611b43565b15613fc857600101613e31565b5091929360009591955b83518110156140075780614001613fee61150460019488611b43565b613d84613ffa8b613213565b9a89611b43565b01613fd2565b5093929150933880613e39565b9091929450600181036140c657506014606061402f84611b31565b5101515103614083578161407a91614059610c3e6060614051613cd297611b31565b510151611848565b918760a061407161406984611b31565b515193611b31565b510151936152cb565b92903880613cc5565b610eb5606061409184611b31565b5101516040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352600483016122ba565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9190811015611b3e5760051b0190565b35610746816111a5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd120820191821161324057565b9190820391821161324057565b90614153939291613cb2565b91939092614160826126c5565b9261416a836126c5565b94600091825b8851811015614271576000805b8a8884898285106141ed575b50505050501561419b57600101614170565b6141ab611504610e67928b611b43565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6115046142259261421f61421a8873ffffffffffffffffffffffffffffffffffffffff976114bb966140f3565b614103565b95611b43565b9116146142345760010161417d565b6001915061425361424961421a838b8b6140f3565b613d84888c611b43565b61426661425f87613213565b968b611b43565b52388a888489614189565b509097965094939291909460ff811690816000985b8a518a10156143435760005b8b8782108061433a575b1561432d5773ffffffffffffffffffffffffffffffffffffffff6142ce6114bb6115048f61421f61421a888f8f6140f3565b9116146142e3576142de90613213565b614292565b93996142f360019294939b613c59565b9461430f61430561421a838b8b6140f3565b613d848d8c611b43565b61432261431b8c613213565b9b8b611b43565b525b01989091614286565b5050919098600190614324565b5085151561429c565b98509250939594975091508161436c575050508151810361436357509190565b80825283529190565b610e67929161437a9161413a565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b908160209103126100e75751610746816111a5565b604051906103b2826102cd565b908160209103126100e757604051906143e8826102cd565b51815290565b90610746916020815260e06144e36144b0614417855161010060208701526101208601906103db565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff1660608601526060860151608086015261447c608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103db565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103db565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103db565b3d15614542573d9061452882610369565b9161453660405193846102e9565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff610746949316815281602082015201906103db565b91909161457f61201c565b50614590610c3e6080830151611848565b6145a0610c3e6060840151611848565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa94851561101e5760009561492b575b5073ffffffffffffffffffffffffffffffffffffffff85169485156148e857614660816149b2565b5061466d610b6b826149dc565b6148e857506147a6916020916146d561468689876155eb565b9661468f6143c3565b5061471f60a08251926147016146aa610c3e8a840151611848565b6040519687918b830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752866102e9565b01519361470c61034a565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c082015261476e6103a3565b60e0820152604051809381927f39077537000000000000000000000000000000000000000000000000000000008352600483016143ee565b03816000885af1600091816148b7575b506147fa57846147c4614517565b90610eb56040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401614547565b84909373ffffffffffffffffffffffffffffffffffffffff83160361484d575b5050505161484561482961035a565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b614856916155eb565b9080821080156148a3575b61486b578361481a565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b506148ae818361413a565b83511415614861565b6148da91925060203d6020116148e1575b6148d281836102e9565b8101906143d0565b90386147b6565b503d6148c8565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b61494e91955060203d602011614955575b61494681836102e9565b8101906143ae565b9338614638565b503d61493c565b6149867f85572ffb00000000000000000000000000000000000000000000000000000000826157c7565b90816149a0575b81614996575090565b6107469150615767565b90506149ab816156a1565b159061498d565b6149867ff208a58f00000000000000000000000000000000000000000000000000000000826157c7565b6149867faff2afbf00000000000000000000000000000000000000000000000000000000826157c7565b6149867fdc0cbd3600000000000000000000000000000000000000000000000000000000826157c7565b6149867f7909b54900000000000000000000000000000000000000000000000000000000826157c7565b614a626131a4565b5090565b9290612ea19173ffffffffffffffffffffffffffffffffffffffff614ab467ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906103db565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b9060405191614af760c0846102e9565b60848352602083019060a03683375a612ee0811115614b4857600091614b1d839261410d565b82602083519301913090f1903d9060848211614b3f575b6000908286523e9190565b60849150614b34565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff600154163303614bc457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614bfc8151846132a0565b928315614d295760005b848110614c14575050505050565b81811015614d0e57614c296115048286611b43565b73ffffffffffffffffffffffffffffffffffffffff8116801561152b57614c4f83613284565b878110614c6157505050600101614c06565b84811015614cde5773ffffffffffffffffffffffffffffffffffffffff614c8b611504838a611b43565b168214614c9a57600101614c4f565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614d09611504614d03888561413a565b89611b43565b614c8b565b614d24611504614d1e848461413a565b85611b43565b614c29565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b359060208110614d61575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b929192614d996133f7565b93828110156151b657614dbc6135256134ff614db484613213565b93868661324e565b600160ff821603615186575082614dd282613292565b116151575780614df8614df2614dea614dff94613292565b8387876132ad565b90614d53565b8652613292565b8281101561512857614e246135fe6135256134ff614e1c85613213565b94878761324e565b83614e2f82846132a0565b116150f95781614e5061253c614e4884614e5a966132a0565b8388886132ad565b60208801526132a0565b828110156150ca57614e776135fe6135256134ff614e1c85613213565b83614e8282846132a0565b1161509b5781614e9b61253c614e4884614ea5966132a0565b60408801526132a0565b8281101561506c57614ec26135fe6135256134ff614e1c85613213565b83614ecd82846132a0565b1161503d578161363261253c614e4884614ee6966132a0565b8281101561500e57614f036135fe6135256134ff614e1c85613213565b83614f0e82846132a0565b11614fdf578161367f61253c614e4884614f27966132a0565b82614f3182613268565b11614fb05761ffff614f4e6137dd6136b76136b161383f86613268565b91169183614f5c84846132a0565b11614f815761253c614f779183610746966138ae87836132a0565b60a08601526132a0565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9080601f830112156100e75781516151fc816119ab565b9261520a60405194856102e9565b81845260208085019260051b8201019283116100e757602001905b8282106152325750505090565b602080918351615241816111a5565b815201910190615225565b906020828203126100e757815167ffffffffffffffff81116100e75761074692016151e5565b95949060019460a09467ffffffffffffffff6152c69573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906103db565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095929493929091906020838060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa92831561101e57600093615450575b5061537383614a06565b6153ad575b5050505050815115615388575090565b6107469150613d0160029167ffffffffffffffff166000526004602052604060002090565b6000949650859161540373ffffffffffffffffffffffffffffffffffffffff92604051998a97889687957f89720a6200000000000000000000000000000000000000000000000000000000875260048701615272565b0392165afa91821561101e5760009261542b575b50615421826158e2565b3880808080615378565b6154499192503d806000833e61544181836102e9565b81019061524c565b9038615417565b61546a91935060203d6020116149555761494681836102e9565b9138615369565b90916060828403126100e757815167ffffffffffffffff81116100e7578361549a9184016151e5565b92602083015167ffffffffffffffff81116100e7576040916154bd9185016151e5565b92015160ff811681036100e75790565b9190916154f16109248267ffffffffffffffff166000526004602052604060002090565b90833b61550e575b50606001519150615508613c3d565b90600090565b61551784614a30565b156154f9576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa801561101e576000809481926155c3575b50615592816158e2565b61559b856158e2565b8051158015906155b7575b6155b057506154f9565b9392909150565b5060ff821615156155a6565b91506155e29294503d8091833e6155da81836102e9565b810190615471565b90939138615588565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161566a575b5061566657826147c4614517565b9150565b90916020823d602011615699575b81615685602093836102e9565b810103126156965750519038615658565b80fd5b3d9150615678565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526157016044836102e9565b6179185a1061573d576020926000925191617530fa6000513d82615731575b508161572a575090565b9050151590565b60201115915038615720565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526157016044836102e9565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526157016044836102e9565b600254811015611b3e5760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b806000526003602052604060002054156000146158dc57600254680100000000000000008110156102ac57600181016002556000600254821015611b3e57600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b50600090565b80519060005b8281106158f457505050565b60018101808211613240575b83811061591057506001016158e8565b73ffffffffffffffffffffffffffffffffffffffff61592f8385611b43565b51166159416114bb6115048487611b43565b1461594e57600101615900565b610e6761595e6115048486611b43565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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
