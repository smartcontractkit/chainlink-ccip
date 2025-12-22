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
	LocalChainSelector   uint64
	GasForCallExactCheck uint16
	RmnRemote            common.Address
	TokenAdminRegistry   common.Address
}

var OffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidGasLimitOverride\",\"inputs\":[{\"name\":\"messageGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f61635f38819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a160405161610b908161025482396080518181816101560152611b5d015260a0518181816101b90152611a84015260c0518181816101e101528181614f200152615aa5015260e05181818161017d01526154540152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780633b81ab9b146100c857806349d8033e146100c35780635215505b146100be5780635643a782146100b957806361a10e59146100b457806379ba5097146100af5780638da5cb5b146100aa578063e9d68a8e146100a55763f2fde38b146100a057600080fd5b6112a4565b6110c2565b611042565b610f59565b610e8a565b610aaf565b61095c565b610763565b61066c565b610563565b61043b565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610290565b828152826020820152826040820152015261025d60405161014b81610290565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ac57604052565b610261565b60a0810190811067ffffffffffffffff8211176102ac57604052565b6101c0810190811067ffffffffffffffff8211176102ac57604052565b6020810190811067ffffffffffffffff8211176102ac57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ac57604052565b6040519061035660c083610306565b565b6040519061035660a083610306565b6040519061035661010083610306565b60405190610356604083610306565b67ffffffffffffffff81116102ac57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906103cf602083610306565b60008252565b60005b8381106103e85750506000910152565b81810151838201526020016103d8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610434815180928187528780880191016103d5565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061047c8183610306565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906103f8565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106105045750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104f7565b9161055c60ff9161054e6040949796976060875260608701906104e6565b9085820360208701526104e6565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106136105c16105bb61025d9336906004016104b8565b90613aef565b6105cf610140820151611398565b60601c9067ffffffffffffffff815116916101808201519061060d8161ffff60a0860151169463ffffffff60806101a0830151519201511690614191565b9361432d565b60409391935193849384610530565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b63ffffffff8116036100e757565b359061035682610653565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106bb9036906004016104b8565b9060243567ffffffffffffffff81116100e7576106dc903690600401610622565b926044359367ffffffffffffffff85116100e757610701610716953690600401610622565b9390926064359561071187610653565b611970565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561075157565b610718565b9060048210156107515752565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356000526007602052602060ff604060002054166107b56040518092610756565bf35b9080602083519182815201916020808360051b8301019401926000915b8383106107e357505050505090565b909192939460208061081f837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516103f8565b970193019301919392906107d4565b6108999173ffffffffffffffffffffffffffffffffffffffff82511681526020820151151560208201526080610888610876604085015160a0604086015260a08501906107b7565b606085015184820360608601526104e6565b9201519060808184039101526104e6565b90565b6040810160408252825180915260206060830193019060005b81811061093c575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106108f157505050505090565b909192939460208061092d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161082e565b970193019301919392906108e2565b825167ffffffffffffffff168552602094850194909201916001016108b5565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757600254610997816120de565b906109a56040519283610306565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06109d2826120de565b0160005b818110610a985750506109e881612122565b9060005b818110610a0457505061025d6040519283928361089c565b80610a3c610a23610a16600194615d58565b67ffffffffffffffff1690565b610a2d83876121b2565b9067ffffffffffffffff169052565b610a7c610a77610a5d610a4f84886121b2565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b61227a565b610a8682876121b2565b52610a9181866121b2565b50016109ec565b602090610aa36120f6565b828287010152016109d6565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610afe903690600401610622565b90610b0761482f565b6000905b828210610b1457005b610b27610b2283858461241f565b61254d565b6020810191610b41610a16845167ffffffffffffffff1690565b15610e6057610b83610b6a610b6a845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610e53575b610bda5760808201949060005b86518051821015610c0457610b6a610bb383610bcd936121b2565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610bda57600101610b98565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610c3c57610b6a610bb383610c2f936121b2565b15610bda57600101610c14565b505095929491909394610c52865182519061487a565b610c67610a5d835167ffffffffffffffff1690565b90610c97610c7d845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610ca187615d8d565b606085019860005b8a518051821015610cf95790610cc1816020936121b2565b518051928391012091158015610ce9575b610bda57610ce26001928b615e5c565b5001610ca9565b50610cf2612602565b8214610cd2565b5050976001975093610e10610e49946003610e3c95610e087f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e99610a16979f610d5660408e610d4f610d9c945160018b016127d9565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610dfe610dbd8d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b51600285016128be565b5191016128be565b610e2d610e28610a16835167ffffffffffffffff1690565b615dcb565b505167ffffffffffffffff1690565b9260405191829182612952565b0390a20190610b0b565b5060808201515115610b8b565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e7576024359060443567ffffffffffffffff81116100e757610f1c903690600401610622565b926064359367ffffffffffffffff85116100e757610f41610716953690600401610622565b93909260843595610f5187610653565b6004016131c2565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303611018577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8116036100e757565b359061035682611094565b90602061089992818152019061082e565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff60043561110681611094565b61110e6120f6565b501660005260046020526040600020604051611129816102b1565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015260018201805490611161826120de565b9161116f6040519384610306565b80835260208301916000526020600020916000905b8282106111c35761025d866111b260038a8960408501526111a760028201612219565b606085015201612219565b6080820152604051918291826110b1565b604051600085546111d3816121c6565b8084529060018116908115611245575060011461120d575b50600192826111ff85946020940382610306565b815201940191019092611184565b6000878152602081209092505b81831061122f575050810160200160016111eb565b600181602092548386880101520192019161121a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506111eb565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff6004356112f481611286565b6112fc61482f565b1633811461136e57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106113d0575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b801515036100e757565b908160209103126100e7575161089981611402565b6040513d6000823e3d90fd5b9060206108999281815201906103f8565b60409073ffffffffffffffffffffffffffffffffffffffff610899949316815281602082015201906103f8565b92919261147782610386565b916114856040519384610306565b8294818452818301116100e7578281602093846000960137010152565b9060048110156107515760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061150557505050505090565b90919293946020806115aa837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a06115996115876115756115638887015160c08a88015260c08701906103f8565b604087015186820360408801526103f8565b606086015185820360608701526103f8565b608085015184820360808601526103f8565b9201519060a08184039101526103f8565b970193019301919392906114f6565b9160209082815201919060005b8181106115d35750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356115fc81611286565b1681520194019291016115c6565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b8484106116c15750505050505090565b909192939495602080611707837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030188526117018b88611649565b9061160a565b98019401940192949391906116b1565b9796949161193790608095611945956119246103569a956117a98e60a0815261174d60a08201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660c0820152604083015167ffffffffffffffff1660e0820152606083015163ffffffff16610100820152828c015163ffffffff1661012082015261014060a084015191019061ffff169052565b8d61016060c08301519101528d6101a06118f06118ba61188461184e6118186117e460e08901516101c06101808a01526102608901906103f8565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6089830301888a01526103f8565b6101208801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60888303016101c08901526103f8565b6101408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016101e08801526103f8565b6101608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303016102008701526103f8565b6101808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60858303016102208601526114d9565b920151906102407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60828503019101526103f8565b9260208d01528b830360408d01526115b9565b9188830360608a0152611699565b94019063ffffffff169052565b806119636040926108999594610756565b81602082015201906103f8565b959192939594909461198460065460ff1690565b6120b4576119b860017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006006541617600655565b6119ca6119c58783613aef565b614767565b95611a6b6020611a106119e8610a168b5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156120af57600091612080575b5061203557611ae3611adf611ad5610a5d8a5167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b611fea57611afc610c7d885167ffffffffffffffff1690565b611b26611adf60e08a01928351602081519101209060019160005201602052604060002054151590565b611fb357506101008701516014815114801590611f82575b611f4b5750602087015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603611f145750828603611ee0576101408701516014815103611ea6575063ffffffff84168015159081611e80575b50611e335790611bc791369161146b565b6020815191012095611bed611be6886000526007602052604060002090565b5460ff1690565b611bf681610747565b8015908115611e1f575b5015611dad57611cb992611cbe959492611c8d92611c56611c2b8b6000526007602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957f61a10e590000000000000000000000000000000000000000000000000000000060208801528b8b60248901611717565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610306565b614773565b9015611d7b577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b611d0f84611d0a886000526007602052604060002090565b6114a2565b611d4b611d396040611d29885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b91836040519485941697169583611952565b0390a46103567fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060065416600655565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff600392611cf2565b611e1b8787611dd96040611dc9835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150611e2c81610747565b1438611c00565b611e1b84611e4860808a015163ffffffff1690565b7fdf2964df0000000000000000000000000000000000000000000000000000000060005263ffffffff90811660045216602452604490565b9050611e9f611e9660808a015163ffffffff1690565b63ffffffff1690565b1138611bb6565b611edc906040519182917f8d666f600000000000000000000000000000000000000000000000000000000083526004830161142d565b0390fd5b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004869052602483905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b611edc906040519182917f55216e31000000000000000000000000000000000000000000000000000000008352306004840161143e565b50611f95611f8f82611398565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff16301415611b3e565b611edc90516040519182917fa50bd1470000000000000000000000000000000000000000000000000000000083526004830161142d565b611e1b611fff885167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611e1b61204a885167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b6120a2915060203d6020116120a8575b61209a8183610306565b81019061140c565b38611ab5565b503d612090565b611421565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116102ac5760051b60200190565b60405190612103826102b1565b6060608083600081526000602082015282604082015282808201520152565b9061212c826120de565b6121396040519182610306565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061216782946120de565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156121ad5760200190565b612171565b80518210156121ad5760209160051b010190565b90600182811c9216801561220f575b60208310146121e057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916121d5565b906040519182815491828252602082019060005260206000209260005b81811061224b57505061035692500383610306565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612236565b90604051612287816102b1565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c1615156020830152600181018054906122c1826120de565b916122cf6040519384610306565b80835260208301916000526020600020916000905b8282106123195750505050600360809261231492604086015261230960028201612219565b606086015201612219565b910152565b60405160008554612329816121c6565b808452906001811690811561239b5750600114612363575b506001928261235585946020940382610306565b8152019401910190926122e4565b6000878152602081209092505b81831061238557505081016020016001612341565b6001816020925483868801015201920191612370565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050612341565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b90156121ad5780610899916123dc565b908210156121ad576108999160051b8101906123dc565b359061035682611286565b359061035682611402565b9080601f830112156100e7578160206108999335910161146b565b9080601f830112156100e757813561247e816120de565b9261248c6040519485610306565b81845260208085019260051b820101918383116100e75760208201905b8382106124b857505050505090565b813567ffffffffffffffff81116100e7576020916124db8784809488010161244c565b8152019101906124a9565b9080601f830112156100e75781356124fd816120de565b9261250b6040519485610306565b81845260208085019260051b8201019283116100e757602001905b8282106125335750505090565b60208091833561254281611286565b815201910190612526565b60c0813603126100e75761255f610347565b9061256981612436565b8252612577602082016110a6565b602083015261258860408201612441565b6040830152606081013567ffffffffffffffff81116100e7576125ae9036908301612467565b6060830152608081013567ffffffffffffffff81116100e7576125d490369083016124e6565b608083015260a08101359067ffffffffffffffff82116100e7576125fa913691016124e6565b60a082015290565b6040516020810190600082526020815261261d604082610306565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81811061265d575050565b60008155600101612652565b9190601f811161267857505050565b610356926000526020600020906020601f840160051c830193106126a4575b601f0160051c0190612652565b9091508190612697565b919091825167ffffffffffffffff81116102ac576126d6816126d084546121c6565b84612669565b6020601f8211600114612734578190612725939495600092612729575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0151905038806126f3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061276784600052602060002090565b9160005b8181106127c15750958360019596971061278a575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080612780565b9192602060018192868b01518155019401920161276b565b8151916801000000000000000083116102ac57815483835580841061283b575b50602061280d910191600052602060002090565b6000915b83831061281e5750505050565b600160208261282f839451866126ae565b01920192019190612811565b8260005283602060002091820191015b81811061285857506127f9565b80612865600192546121c6565b80612872575b500161284b565b601f811183146128885750600081555b3861286b565b6128ac9083601f61289e85600052602060002090565b920160051c82019101612652565b60008181526020812081835555612882565b81519167ffffffffffffffff83116102ac576801000000000000000083116102ac576020908254848455808510612935575b500190600052602060002060005b83811061290b5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016128fe565b61294c908460005285846000209182019101612652565b386128f0565b90610899916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a06129e96129b6606085015160c0608086015260e08501906107b7565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526104e6565b9201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526104e6565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b91602061089993818152019161160a565b919091357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106113d0575050565b3561089981610653565b3561089981611094565b61ffff8116036100e757565b3561089981612b1a565b91909160c0818403126100e757612b45610347565b9281358452602082013567ffffffffffffffff81116100e75781612b6a91840161244c565b6020850152604082013567ffffffffffffffff81116100e75781612b8f91840161244c565b6040850152606082013567ffffffffffffffff81116100e75781612bb491840161244c565b6060850152608082013567ffffffffffffffff81116100e75781612bd991840161244c565b608085015260a082013567ffffffffffffffff81116100e757612bfc920161244c565b60a0830152565b929190612c0f816120de565b93612c1d6040519586610306565b602085838152019160051b8101918383116100e75781905b838210612c43575050505050565b813567ffffffffffffffff81116100e757602091612c648784938701612b30565b815201910190612c35565b908210156121ad57612c869160051b810190612a70565b9091565b908160209103126100e7575161089981611286565b60409073ffffffffffffffffffffffffffffffffffffffff6108999593168152816020820152019161160a565b359061035682612b1a565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310612d525750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e7576020612e53600193868394019081358152612e45612e3a612e1f612e04612de9612dd989880188611649565b60c08b89015260c088019161160a565b612df66040880188611649565b90878303604089015261160a565b612e116060870187611649565b90868303606088015261160a565b612e2c6080860186611649565b90858303608087015261160a565b9260a0810190611649565b9160a081850391015261160a565b980196019493019190612d42565b6130d96108999593949260608352612e8d60608401612e7f836110a6565b67ffffffffffffffff169052565b612ead612e9c602083016110a6565b67ffffffffffffffff166080850152565b612ecd612ebc604083016110a6565b67ffffffffffffffff1660a0850152565b612ee9612edc60608301610661565b63ffffffff1660c0850152565b612f05612ef860808301610661565b63ffffffff1660e0850152565b612f20612f1460a08301612ccc565b61ffff16610100850152565b60c08101356101208401526130a861309c61305d61301e612fdf612fa0612f61612f4d60e0890189611649565b6101c06101408d01526102208c019161160a565b612f6f610100890189611649565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d015261160a565b612fae610120880188611649565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c015261160a565b612fed610140870187611649565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b015261160a565b61302c610160860186611649565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a015261160a565b61306b610180850185612cd7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152612d2a565b916101a0810190611649565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08584030161020086015261160a565b936020820152604081850391015261160a565b604051906040820182811067ffffffffffffffff8211176102ac5760405260006020838281520152565b90613120826120de565b61312d6040519182610306565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061315b82946120de565b019060005b82811061316c57505050565b6020906131776130ec565b82828501015201613160565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd12082019182116131b057565b612623565b919082039182116131b057565b9391949092943033036137225760006131df610180870187612a1c565b9050613643575b613268613203611f8f6131fd6101408a018a612a70565b90612ad2565b9261322d846132166101a08b018b612a70565b9050613227611e9660808d01612b06565b90614191565b98899160a061323b8b612b10565b876132628d61325a613251610180830183612a1c565b96909201612b26565b943691612c03565b91614b15565b92909860005b8a5181101561345857806020613291610b6a610b6a8f6132dd96610bb3916121b2565b6132a661329e848a6121b2565b518a8c612c6f565b91906040518096819482937fc3a7ded600000000000000000000000000000000000000000000000000000000845260048401612ac1565b03915afa9182156120af57600092613428575b5073ffffffffffffffffffffffffffffffffffffffff8216156133cb5761332261331a82886121b2565b51888a612c6f565b929073ffffffffffffffffffffffffffffffffffffffff82163b156100e7576000918b8373ffffffffffffffffffffffffffffffffffffffff8f613395604051998a97889687947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601612e61565b0393165af19182156120af576001926133b0575b500161326e565b806133bf60006133c593610306565b806100dc565b386133a9565b856133f389898f856133e6610bb3611edc986133ec946121b2565b956121b2565b5191612c6f565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501612c9f565b61344a91925060203d8111613451575b6134428183610306565b810190612c8a565b90386132f0565b503d613438565b509597935095935096505061347b613474610180840184612a1c565b9050613116565b9561348a610180840184612a1c565b9050613531575b5061352a57610356946134fa6134a683612b10565b6134e76134ee6134ba610120870187612a70565b6134c86101a0890189612a70565b9490956134d3610358565b9c8d5267ffffffffffffffff1660208d0152565b369161146b565b6040890152369161146b565b6060860152608085015263ffffffff8216156135175750916153bc565b6135249150608001612b06565b916153bc565b5050505050565b61358b61354b613545610180860186612a1c565b9061240f565b613559610120860186612a70565b61358561356588612b10565b9261357d61357560a08b01612b26565b953690612b30565b92369161146b565b90614e8b565b9190613596896121a0565b5273ffffffffffffffffffffffffffffffffffffffff6135d0611f8f6131fd6135c66135456101808a018a612a1c565b6080810190612a70565b921673ffffffffffffffffffffffffffffffffffffffff8316036135f5575b50613491565b61362961362e926136236136088b6121a0565b515173ffffffffffffffffffffffffffffffffffffffff1690565b90614a0f565b6131b5565b6020613639886121a0565b51015238806135ef565b50601461366461365a613545610180890189612a1c565b6060810190612a70565b90500361370e5760146136816135c6613545610180890189612a1c565b9050036136c4576136bf6136a5611f8f6131fd6135c66135456101808b018b612a1c565b613623611f8f6131fd61365a6135456101808c018c612a1c565b6131e6565b6136d86135c6613545610180880188612a1c565b90611edc6040519283927f8d666f6000000000000000000000000000000000000000000000000000000000845260048401612ac1565b6136d861365a613545610180880188612a1c565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190613759826102cd565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b90156121ad5790565b90604310156121ad5760430190565b908210156121ad570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906023116100e75760210190600290565b906043116100e75760230190602090565b90929192836044116100e75783116100e757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff000000000000000000000000000000000000000000000000811692600881106138db575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613941575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff000000000000000000000000000000000000000000000000000000000000811692600281106139a7575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b3590602081106139e7575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff8211176102ac57604052606060a0836000815282602082015282604082015282808201528260808201520152565b60408051909190613a678382610306565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b828110613a9e57505050565b602090613aa9613a14565b82828501015201613a92565b60405190613ac4602083610306565b600080835282815b828110613ad857505050565b602090613ae3613a14565b82828501015201613acc565b90613af861374c565b91604d821061416057613b3d613b37613b1184846137b9565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff8216036141305750613b76613b68613b62613b5c85856137dd565b906138a7565b60c01c90565b67ffffffffffffffff168452565b613b9a613b89613b62613b5c85856137ee565b67ffffffffffffffff166020850152565b613bbe613bad613b62613b5c85856137ff565b67ffffffffffffffff166040850152565b613bea613bdd613bd7613bd18585613810565b9061390d565b60e01c90565b63ffffffff166060850152565b613c0a613bfd613bd7613bd18585613821565b63ffffffff166080850152565b613c34613c29613c23613c1d8585613832565b90613973565b60f01c90565b61ffff1660a0850152565b613c47613c418383613843565b906139d9565b60c0840152816043101561410157613c6e613c68613b37613b1185856137c2565b60ff1690565b90816044018381116140d257613c886134e7828685613854565b60e0860152838110156140a357613c68613b37613b11613ca99387866137d1565b8201916045830190848211614075576134e7826045613cca9301878661388f565b6101008601528381101561404657613cee613c68613b37613b1160459488876137d1565b830101916001830190848211614017576134e7826046613d109301878661388f565b61012086015283811015613fe857613d34613c68613b37613b1160019488876137d1565b830101916001830190848211613fb9576134e7826002613d569301878661388f565b6101408601526003830192848411613f8a57613d86613d7f613c23613c1d876001968a8961388f565b61ffff1690565b0101916002830190848211613f5b576134e782613da492878661388f565b6101608601526004830190848211613f2c57613c23613c1d83613dc893888761388f565b9261ffff8294168015600014613ec257505050613de3613ab5565b6101808501525b6002820191838311613e935780613e0e613d7f613c23613c1d876002968a8961388f565b010191838311613e6457826134e79185613e279461388f565b6101a084015203613e355790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b6002919294508190613ee5613ed5613a56565b966101808a0197885288876154f9565b9490965196613ef486986121a0565b5201010114613dea577fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600860045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000006000526004805260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052611e1b6024906000600452565b1591908261420f575b508115614205575b81156141ac575090565b90506141d87f85572ffb0000000000000000000000000000000000000000000000000000000082615fde565b90816141f3575b816141e957501590565b611adf9150615f7e565b90506141fe81615eb8565b15906141df565b803b1591506141a2565b1591503861419a565b60405190614227602083610306565b6000808352366020840137565b604080519091906142458382610306565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001366020840137565b90600182018092116131b057565b919082018092116131b057565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146131b05760010190565b80548210156121ad5760005260206000200190600090565b80156131b0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff1680156131b0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9294939160609160009361433f614218565b978351806146e8575b5050156146d4575051156146c3575050614360614218565b614368614218565b6000945b60026143b561439a60036143948867ffffffffffffffff166000526004602052604060002090565b01612219565b9567ffffffffffffffff166000526004602052604060002090565b01936143cf6143c78551845190614282565b825190614282565b906143e46143df87548094614282565b612122565b9788966000805b885182101561443e579061443660019261441b61440b610bb3858e6121b2565b916144158161428f565b9d6121b2565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018a986143eb565b91939597505097909294976000905b8751821015614485579061447d60019261441b61446d610bb3858d6121b2565b916144778161428f565b9c6121b2565b01899761444d565b9750509193969092945060005b85518110156144c757806144c16144ae610bb36001948a6121b2565b61441b6144ba8b61428f565b9a8d6121b2565b01614492565b509296935093909460005b82811061463f575b50509091929350600090815b8181106145a05750508252805160005b83518110156145975760005b82811061451257506001016144f6565b61451f610bb382866121b2565b73ffffffffffffffffffffffffffffffffffffffff614544610b6a610bb3868a6121b2565b911614614559576145549061428f565b614502565b91614563906142d4565b9161457e614574610bb385876121b2565b61441b83876121b2565b60ff8716156145025795614591906142ff565b95614502565b50815290929091565b6145ad610bb382876121b2565b73ffffffffffffffffffffffffffffffffffffffff8116801561463557600090815b868110614609575b50509060019291156145ec575b505b016144e6565b6146039061441b6145fc8761428f565b96896121b2565b386145e4565b8161461a610b6a610bb3848d6121b2565b14614627576001016145cf565b5060019150819050386145d7565b50506001906145e6565b61464f610b6a610bb383886121b2565b1561465c576001016144d2565b50909192939460005b828110614677578695949392506144da565b806146bd6146aa61468a600194866142bc565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b61441b6146b68861428f565b978a6121b2565b01614665565b6146ce949194614234565b9161436c565b9150506146e2915082615c1c565b9461436c565b909198506001810361473a575061473290614711611f8f86614709876121a0565b510151611398565b9061471b856121a0565b51518860a0614729886121a0565b51015193615a3b565b963880614348565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61476f61374c565b5090565b906040519161478360c084610306565b60848352602083019060a03683375a612ee08111156147d4576000916147a98392613183565b82602083519301913090f1903d90608482116147cb575b6000908286523e9190565b608491506147c0565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361485057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614888815184614282565b9283156149b55760005b8481106148a0575050505050565b8181101561499a576148b5610bb382866121b2565b73ffffffffffffffffffffffffffffffffffffffff81168015610bda576148db83614274565b8781106148ed57505050600101614892565b8481101561496a5773ffffffffffffffffffffffffffffffffffffffff614917610bb3838a6121b2565b168214614926576001016148db565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614995610bb361498f88856131b5565b896121b2565b614917565b6149b0610bb36149aa84846131b5565b856121b2565b6148b5565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b3d15614a0a573d906149f082610386565b916149fe6040519384610306565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181614ac4575b50614ac05782614a8a6149df565b90611edc6040519283927f9fe2f95a0000000000000000000000000000000000000000000000000000000084526004840161143e565b9150565b90916020823d602011614af3575b81614adf60209383610306565b81010312614af05750519038614a7c565b80fd5b3d9150614ad2565b91908110156121ad5760051b0190565b3561089981611286565b90614b2494959693929161432d565b91939092614b3182612122565b92614b3b83612122565b94600091825b8851811015614c3c576000805b8a888489828510614bbe575b505050505015614b6c57600101614b41565b614b7c610bb3611e1b928b6121b2565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610bb3614bf0926133e6614beb8873ffffffffffffffffffffffffffffffffffffffff97610b6a96614afb565b614b0b565b911614614bff57600101614b4e565b60019150614c1e614c14614beb838b8b614afb565b61441b888c6121b2565b614c31614c2a8761428f565b968b6121b2565b52388a888489614b5a565b509097965094939291909460ff811690816000985b8a518a1015614d0e5760005b8b87821080614d05575b15614cf85773ffffffffffffffffffffffffffffffffffffffff614c99610b6a610bb38f6133e6614beb888f8f614afb565b911614614cae57614ca99061428f565b614c5d565b9399614cbe60019294939b6142d4565b94614cda614cd0614beb838b8b614afb565b61441b8d8c6121b2565b614ced614ce68c61428f565b9b8b6121b2565b525b01989091614c51565b5050919098600190614cef565b50851515614c67565b985092509395949750915081614d375750505081518103614d2e57509190565b80825283529190565b611e1b9291614d45916131b5565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906103cf826102ea565b908160209103126100e75760405190614d9e826102ea565b51815290565b6108999160e0614e4e614e3c614dc5855161010086526101008601906103f8565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152614e2a6080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526103f8565b60c085015184820360c08601526103f8565b9201519060e08184039101526103f8565b906020610899928181520190614da4565b9061ffff61055c602092959495604085526040850190614da4565b92939193614e976130ec565b50614ea8611f8f6080860151611398565b91614eb9611f8f6060870151611398565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909690956020878060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9687156120af576000976151e8575b5073ffffffffffffffffffffffffffffffffffffffff87169485156151a457614f78614d79565b50614fc6825191614fa960a0602086015195015195614f95610367565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff8816608084015260a083015260c0820152614ff96103c0565b60e0820152615007856158d5565b156150cd579061504b9260209260006040518096819582947f489a68f200000000000000000000000000000000000000000000000000000000845260048401614e70565b03925af16000918161509c575b506150665783614a8a6149df565b929091925b51615093615077610377565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201529190565b6150bf91925060203d6020116150c6575b6150b78183610306565b810190614d86565b9038615058565b503d6150ad565b90506150d88461592b565b156151605761511b6000926020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301614e5f565b03925af16000918161513f575b506151365783614a8a6149df565b9290919261506b565b61515991925060203d6020116150c6576150b78183610306565b9038615128565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff881660045260246000fd5b61520291975060203d602011613451576134428183610306565b9538614f51565b90916060828403126100e757815161522081611402565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e75781519161524f83610386565b9161525d6040519384610306565b838352602084830101116100e75760409261527e91602080850191016103d5565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806152fb6152c7604084015160a060c08801526101208701906103f8565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103f8565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106153845750505061ffff909516602083015261035692916060916153689063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff168552602090810151818601526040909401939092019160010161533a565b9060009161547d938361542473ffffffffffffffffffffffffffffffffffffffff61540967ffffffffffffffff60208701511667ffffffffffffffff166000526004602052604060002090565b541673ffffffffffffffffffffffffffffffffffffffff1690565b92604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601615284565b03925af19081156120af576000906000926154d2575b501561549c5750565b611edc906040519182917f0a8d6e8c0000000000000000000000000000000000000000000000000000000083526004830161142d565b90506154f191503d806000833e6154e98183610306565b810190615209565b509038615493565b9291615503613a14565b91808210156158a65761551d613b37613b118484896137d1565b600160ff82160361413057506021820181811161587757615546613c418260018601858a61388f565b84528181101561584857613c68613b37613b1161556493858a6137d1565b8201916022830190828211615819576134e78260226155859301858a61388f565b6020850152818110156157ea576155a8613c68613b37613b11602294868b6137d1565b8301019160018301908282116157bb576134e78260236155ca9301858a61388f565b60408501528181101561578c576155ed613c68613b37613b11600194868b6137d1565b83010191600183019082821161575d576134e782600261560f9301858a61388f565b60608501528181101561572e57615632613c68613b37613b11600194868b6137d1565b83010160018101928284116156ff576134e78460026156539301858a61388f565b608085015260038101928284116156d05760029161567e613d7f613c23613c1d88600196898e61388f565b010101948186116156a157615698926134e792879261388f565b60a08201529190565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b6158ff7f331710310000000000000000000000000000000000000000000000000000000082615fde565b9081615919575b8161590f575090565b6108999150615f7e565b905061592481615eb8565b1590615906565b6158ff7faff2afbf0000000000000000000000000000000000000000000000000000000082615fde565b9080601f830112156100e757815161596c816120de565b9261597a6040519485610306565b81845260208085019260051b8201019283116100e757602001905b8282106159a25750505090565b6020809183516159b181611286565b815201910190615995565b906020828203126100e757815167ffffffffffffffff81116100e7576108999201615955565b95949060019460a09467ffffffffffffffff615a369573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906103f8565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095909291906020848060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9384156120af57600094615b9f575b50615ae0846158d5565b615afe575b505050505050805115615af55790565b50610899614234565b600095965090615b5373ffffffffffffffffffffffffffffffffffffffff92604051988997889687957f89720a62000000000000000000000000000000000000000000000000000000008752600487016159e2565b0392165afa9081156120af57600091615b7c575b50615b7181616040565b388080808080615ae5565b615b9991503d806000833e615b918183610306565b8101906159bc565b38615b67565b615bb991945060203d602011613451576134428183610306565b9238615ad6565b90916060828403126100e757815167ffffffffffffffff81116100e75783615be9918401615955565b92602083015167ffffffffffffffff81116100e757604091615c0c918501615955565b92015160ff811681036100e75790565b90615c477f7909b5490000000000000000000000000000000000000000000000000000000082615fde565b80615d48575b80615d39575b615c71575b5050615c62614234565b90615c6b614218565b90600090565b6040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9290921660048301526000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa80156120af57600080928192615d13575b50615ce681616040565b615cef83616040565b805115801590615d07575b615d045750615c58565b92565b5060ff82161515615cfa565b9150615d3192503d8091833e615d298183610306565b810190615bc0565b909138615cdc565b50615d4381615f7e565b615c53565b50615d5281615eb8565b15615c4d565b6002548110156121ad5760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b818110615da157505060009055565b80615dae600192856142bc565b90549060031b1c6000528184016020526000604081205501615d92565b600081815260036020526040902054615e5657600254680100000000000000008110156102ac57615e3d615e0882600185940160025560026142bc565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615eb157805490680100000000000000008210156102ac5782615e9a615e088460018096018555846142bc565b905580549260005201602052604060002055600190565b5050600090565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252615f18604483610306565b6179185a10615f54576020926000925191617530fa6000513d82615f48575b5081615f41575090565b9050151590565b60201115915038615f37565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252615f18604483610306565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252615f18604483610306565b80519060005b82811061605257505050565b600181018082116131b0575b83811061606e5750600101616046565b73ffffffffffffffffffffffffffffffffffffffff61608d83856121b2565b511661609f610b6a610bb384876121b2565b146160ac5760010161605e565b611e1b6160bc610bb384866121b2565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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
	return common.HexToHash("0x4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e4950")
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
