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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f61612538819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615ed19081610254823960805181818161015601526110e0015260a0518181816101b90152611007015260c0518181816101e101528181613c3301526159dc015260e05181818161017d01526129120152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063176133d1146100d2578063181f5a77146100cd57806320f81c88146100c857806349d8033e146100c35780635215505b146100be5780635643a782146100b95780636b8be52c146100b457806379ba5097146100af5780638da5cb5b146100aa578063e9d68a8e146100a55763f2fde38b146100a057600080fd5b611922565b611740565b6116c0565b6115d7565b610e69565b610a8e565b61093b565b610742565b610657565b61052f565b610292565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610384565b828152826020820152826040820152015261025d60405161014b81610384565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760243560443567ffffffffffffffff81116100e757610323903690600401610261565b606435939167ffffffffffffffff85116100e757610348610353953690600401610261565b949093600401612530565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176103a057604052565b610355565b60a0810190811067ffffffffffffffff8211176103a057604052565b6020810190811067ffffffffffffffff8211176103a057604052565b6101c0810190811067ffffffffffffffff8211176103a057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103a057604052565b6040519061044a60c0836103fa565b565b6040519061044a60a0836103fa565b6040519061044a610100836103fa565b6040519061044a6040836103fa565b67ffffffffffffffff81116103a057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906104c36020836103fa565b60008252565b60005b8381106104dc5750506000910152565b81810151838201526020016104cc565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610528815180928187528780880191016104c9565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061057081836103fa565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906104ec565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106105f85750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016105eb565b9161065060ff916106426040949796976060875260608701906105da565b9085820360208701526105da565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106e86106b56106af61025d9336906004016105ac565b90614541565b67ffffffffffffffff815116906106d0610140820151612a1c565b60601c61ffff60a06101808401519301511692614bfe565b60409391935193849384610624565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561073057565b6106f7565b9060048210156107305752565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356000526007602052602060ff604060002054166107946040518092610735565bf35b9080602083519182815201916020808360051b8301019401926000915b8383106107c257505050505090565b90919293946020806107fe837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516104ec565b970193019301919392906107b3565b6108789173ffffffffffffffffffffffffffffffffffffffff82511681526020820151151560208201526080610867610855604085015160a0604086015260a0850190610796565b606085015184820360608601526105da565b9201519060808184039101526105da565b90565b6040810160408252825180915260206060830193019060005b81811061091b575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106108d057505050505090565b909192939460208061090c837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161080d565b970193019301919392906108c1565b825167ffffffffffffffff16855260209485019490920191600101610894565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461097681611b41565b9061098460405192836103fa565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06109b182611b41565b0160005b818110610a775750506109c781612a80565b9060005b8181106109e357505061025d6040519283928361087b565b80610a1b610a026109f5600194615ca6565b67ffffffffffffffff1690565b610a0c8387611d2b565b9067ffffffffffffffff169052565b610a5b610a56610a3c610a2e8488611d2b565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b612b83565b610a658287611d2b565b52610a708186611d2b565b50016109cb565b602090610a82612a54565b828287010152016109b5565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610add903690600401610261565b90610ae6615024565b6000905b828210610af357005b610b06610b01838584612318565b612de1565b6020810191610b206109f5845167ffffffffffffffff1690565b15610e3f57610b62610b49610b49845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610e32575b610bb95760808201949060005b86518051821015610be357610b49610b9283610bac93611d2b565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610bb957600101610b77565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610c1b57610b49610b9283610c0e93611d2b565b15610bb957600101610bf3565b505095929491909394610c31865182519061506f565b610c46610a3c835167ffffffffffffffff1690565b90610c76610c5c845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610c8087615cdb565b606085019860005b8a518051821015610cd85790610ca081602093611d2b565b518051928391012091158015610cc8575b610bb957610cc16001928b615daa565b5001610c88565b50610cd1612e96565b8214610cb1565b5050976001975093610def610e28946003610e1b95610de77f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e996109f5979f610d3560408e610d2e610d7b945160018b0161306d565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610ddd610d9c8d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5160028501613152565b519101613152565b610e0c610e076109f5835167ffffffffffffffff1690565b615d19565b505167ffffffffffffffff1690565b92604051918291826131e6565b0390a20190610aea565b5060808201515115610b6a565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610eb89036906004016105ac565b60243567ffffffffffffffff81116100e757610ed8903690600401610261565b9060443567ffffffffffffffff81116100e757610ef9903690600401610261565b9094610f0760065460ff1690565b6115ad57610f3b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006006541617600655565b610f4d610f488683614541565b6151d4565b94610fee6020610f93610f6b6109f58a5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156115a857600091611579575b5061152e57611066611062611058610a3c895167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b6114e35761107f610c5c875167ffffffffffffffff1690565b6110a961106260e08901928351602081519101209060019160005201602052604060002054151590565b6114ac5750610100860151601481511480159061147b575b6114445750602086015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff82160361140d57508285036113d957610140860151601481510361139f575090611134913691611b59565b602081519101209461115a611153876000526007602052604060002090565b5460ff1690565b61116381610726565b801590811561138b575b501561131957916111f96112259261122a95946111c26111978a6000526007602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519586947f176133d10000000000000000000000000000000000000000000000000000000060208701528a8a602488016134d8565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103fa565b6151e0565b90156112e7577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b61127b84611276886000526007602052604060002090565b6132f2565b6112b76112a56040611295885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b918360405194859416971695836136f9565b0390a46103537fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060065416600655565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff60039261125e565b61138786866113456040611335835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b6003915061139881610726565b143861116d565b6113d5906040519182917f8d666f600000000000000000000000000000000000000000000000000000000083526004830161251f565b0390fd5b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004859052602483905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6113d5906040519182917f55216e3100000000000000000000000000000000000000000000000000000000835230600484016132c5565b5061148e61148882612a1c565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff163014156110c1565b6113d590516040519182917fa50bd1470000000000000000000000000000000000000000000000000000000083526004830161251f565b6113876114f8875167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611387611543875167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b61159b915060203d6020116115a1575b61159381836103fa565b8101906132b0565b38611038565b503d611589565b611dbf565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303611696577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8116036100e757565b359061044a82611712565b90602061087892818152019061080d565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff60043561178481611712565b61178c612a54565b5016600052600460205260406000206040516117a7816103a5565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c1615156020820152600182018054906117df82611b41565b916117ed60405193846103fa565b80835260208301916000526020600020916000905b8282106118415761025d8661183060038a89604085015261182560028201612b22565b606085015201612b22565b60808201526040519182918261172f565b6040516000855461185181612acf565b80845290600181169081156118c3575060011461188b575b506001928261187d859460209403826103fa565b815201940191019092611802565b6000878152602081209092505b8183106118ad57505081016020016001611869565b6001816020925483868801015201920191611898565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050611869565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561197281611904565b61197a615024565b163381146119ec57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611a9b575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b3561087881611712565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b61ffff8116036100e757565b3561087881611b2b565b67ffffffffffffffff81116103a05760051b60200190565b929192611b658261047a565b91611b7360405193846103fa565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e75781602061087893359101611b59565b91909160c0818403126100e757611bc061043b565b9281358452602082013567ffffffffffffffff81116100e75781611be5918401611b90565b6020850152604082013567ffffffffffffffff81116100e75781611c0a918401611b90565b6040850152606082013567ffffffffffffffff81116100e75781611c2f918401611b90565b6060850152608082013567ffffffffffffffff81116100e75781611c54918401611b90565b608085015260a082013567ffffffffffffffff81116100e757611c779201611b90565b60a0830152565b929190611c8a81611b41565b93611c9860405195866103fa565b602085838152019160051b8101918383116100e75781905b838210611cbe575050505050565b813567ffffffffffffffff81116100e757602091611cdf8784938701611bab565b815201910190611cb0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805115611d265760200190565b611cea565b8051821015611d265760209160051b010190565b90821015611d2657611d569160051b810190611a16565b9091565b908160209103126100e7575161087881611904565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916020610878938181520191611d6f565b6040513d6000823e3d90fd5b60409073ffffffffffffffffffffffffffffffffffffffff61087895931681528160208201520191611d6f565b63ffffffff8116036100e757565b359061044a82611df8565b359061044a82611b2b565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310611ee75750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e7576020611fe8600193868394019081358152611fda611fcf611fb4611f99611f7e611f6e89880188611e1c565b60c08b89015260c0880191611d6f565b611f8b6040880188611e1c565b908783036040890152611d6f565b611fa66060870187611e1c565b908683036060880152611d6f565b611fc16080860186611e1c565b908583036080870152611d6f565b9260a0810190611e1c565b9160a0818503910152611d6f565b980196019493019190611ed7565b61226e61087895939492606083526120226060840161201483611724565b67ffffffffffffffff169052565b61204261203160208301611724565b67ffffffffffffffff166080850152565b61206261205160408301611724565b67ffffffffffffffff1660a0850152565b61207e61207160608301611e06565b63ffffffff1660c0850152565b61209a61208d60808301611e06565b63ffffffff1660e0850152565b6120b56120a960a08301611e11565b61ffff16610100850152565b60c081013561012084015261223d6122316121f26121b36121746121356120f66120e260e0890189611e1c565b6101c06101408d01526102208c0191611d6f565b612104610100890189611e1c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d0152611d6f565b612143610120880188611e1c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c0152611d6f565b612182610140870187611e1c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b0152611d6f565b6121c1610160860186611e1c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a0152611d6f565b612200610180850185611e6c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152611ebf565b916101a0810190611e1c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa085840301610200860152611d6f565b9360208201526040818503910152611d6f565b604051906040820182811067ffffffffffffffff8211176103a05760405260006020838281520152565b906122b582611b41565b6122c260405191826103fa565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06122f08294611b41565b019060005b82811061230157505050565b60209061230c612281565b828285010152016122f5565b9190811015611d265760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b3561087881611df8565b801515036100e757565b90916060828403126100e757815161238381612362565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e7578151916123b28361047a565b916123c060405193846103fa565b838352602084830101116100e7576040926123e191602080850191016104c9565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a0840152608061245e61242a604084015160a060c08801526101208701906104ec565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526104ec565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106124e75750505061ffff909516602083015261044a92916060916124cb9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff168552602090810151818601526040909401939092019160010161249d565b9060206108789281815201906104ec565b929495939091953033036129f25761257e9661259961258b9761256361148861255d6101408a018a611a16565b90611a67565b9261259361257089611acd565b856101808b019d8e8c611ad7565b939060a08d019e8f611b37565b943691611c7e565b916137e4565b97909860005b8a5181101561275c5761260e60208b6125d76125cf8f6125c9610b49610b49610b928a8095611d2b565b93611d2b565b51898b611d3f565b91906040518095819482937fc3a7ded600000000000000000000000000000000000000000000000000000000845260048401611dae565b03915afa80156115a85773ffffffffffffffffffffffffffffffffffffffff9160009161272e575b50169081156126d15761265461264c828d611d2b565b518789611d3f565b9290813b156100e7576000918a838d61269c604051988996879586947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601611ff6565b03925af19182156115a8576001926126b6575b500161259f565b806126c560006126cb936103fa565b806100dc565b386126af565b8a6126f988888f856126ec610b926113d5986126f294611d2b565b95611d2b565b5191611d3f565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501611dcb565b61274f915060203d8111612755575b61274781836103fa565b810190611d5a565b38612636565b503d61273d565b509497509495925096505061277b6127748386611ad7565b90506122ab565b9560005b6127898487611ad7565b905081101561280557806127e96127ac6001936127a6888b611ad7565b90612318565b6127ba6101208a018a611a16565b6127e36127c68c611acd565b926127db6127d38d611b37565b953690611bab565b923691611b59565b90613ba2565b6127f3828b611d2b565b526127fe818a611d2b565b500161277f565b509250949390506101a083019061281c8285611a16565b905015806129d9575b80156129d0575b80156129be575b6129b75760006128e260808661293b986128d3612873610b49612859610a3c899d611acd565b5473ffffffffffffffffffffffffffffffffffffffff1690565b976128c06128c761288386611acd565b9261289e612895610120890189611a16565b91909289611a16565b94909560206128ab61044c565b9e8f908152019067ffffffffffffffff169052565b3691611b59565b60408a01523691611b59565b60608701528286015201612358565b91604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f000000000000000000000000000000000000000000000000000000000000000090600486016123e7565b03925af19081156115a857600090600092612990575b501561295a5750565b6113d5906040519182917f0a8d6e8c0000000000000000000000000000000000000000000000000000000083526004830161251f565b90506129af91503d806000833e6129a781836103fa565b81019061236c565b509038612951565b5050505050565b506129cb61106284614042565b612833565b50823b1561282c565b5063ffffffff6129eb60808601612358565b1615612825565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611a9b575050565b60405190612a61826103a5565b6060608083600081526000602082015282604082015282808201520152565b90612a8a82611b41565b612a9760405191826103fa565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612ac58294611b41565b0190602036910137565b90600182811c92168015612b18575b6020831014612ae957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612ade565b906040519182815491828252602082019060005260206000209260005b818110612b5457505061044a925003836103fa565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612b3f565b90604051612b90816103a5565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c161515602083015260018101805490612bca82611b41565b91612bd860405193846103fa565b80835260208301916000526020600020916000905b828210612c2257505050506003608092612c1d926040860152612c1260028201612b22565b606086015201612b22565b910152565b60405160008554612c3281612acf565b8084529060018116908115612ca45750600114612c6c575b5060019282612c5e859460209403826103fa565b815201940191019092612bed565b6000878152602081209092505b818310612c8e57505081016020016001612c4a565b6001816020925483868801015201920191612c79565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050612c4a565b359061044a82611904565b359061044a82612362565b9080601f830112156100e7578135612d1281611b41565b92612d2060405194856103fa565b81845260208085019260051b820101918383116100e75760208201905b838210612d4c57505050505090565b813567ffffffffffffffff81116100e757602091612d6f87848094880101611b90565b815201910190612d3d565b9080601f830112156100e7578135612d9181611b41565b92612d9f60405194856103fa565b81845260208085019260051b8201019283116100e757602001905b828210612dc75750505090565b602080918335612dd681611904565b815201910190612dba565b60c0813603126100e757612df361043b565b90612dfd81612ce5565b8252612e0b60208201611724565b6020830152612e1c60408201612cf0565b6040830152606081013567ffffffffffffffff81116100e757612e429036908301612cfb565b6060830152608081013567ffffffffffffffff81116100e757612e689036908301612d7a565b608083015260a08101359067ffffffffffffffff82116100e757612e8e91369101612d7a565b60a082015290565b60405160208101906000825260208152612eb16040826103fa565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110612ef1575050565b60008155600101612ee6565b9190601f8111612f0c57505050565b61044a926000526020600020906020601f840160051c83019310612f38575b601f0160051c0190612ee6565b9091508190612f2b565b919091825167ffffffffffffffff81116103a057612f6a81612f648454612acf565b84612efd565b6020601f8211600114612fc8578190612fb9939495600092612fbd575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612f87565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690612ffb84600052602060002090565b9160005b8181106130555750958360019596971061301e575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613014565b9192602060018192868b015181550194019201612fff565b8151916801000000000000000083116103a05781548383558084106130cf575b5060206130a1910191600052602060002090565b6000915b8383106130b25750505050565b60016020826130c383945186612f42565b019201920191906130a5565b8260005283602060002091820191015b8181106130ec575061308d565b806130f960019254612acf565b80613106575b50016130df565b601f8111831461311c5750600081555b386130ff565b6131409083601f61313285600052602060002090565b920160051c82019101612ee6565b60008181526020812081835555613116565b81519167ffffffffffffffff83116103a0576801000000000000000083116103a05760209082548484558085106131c9575b500190600052602060002060005b83811061319f5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501613192565b6131e0908460005285846000209182019101612ee6565b38613184565b90610878916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a061327d61324a606085015160c0608086015260e0850190610796565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526105da565b9201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526105da565b908160209103126100e7575161087881612362565b60409073ffffffffffffffffffffffffffffffffffffffff610878949316815281602082015201906104ec565b9060048110156107305760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061335557505050505090565b90919293946020806133fa837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a06133e96133d76133c56133b38887015160c08a88015260c08701906104ec565b604087015186820360408801526104ec565b606086015185820360608701526104ec565b608085015184820360808601526104ec565b9201519060a08184039101526104ec565b97019301930191939290613346565b9160209082815201919060005b8181106134235750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561344c81611904565b168152019401929101613416565b90602083828152019260208260051b82010193836000925b8484106134825750505050505090565b9091929394956020806134c8837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030188526134c28b88611e1c565b90611d6f565b9801940194019294939190613472565b94929361087896946136d86136eb94936080895261350360808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a0152606081015163ffffffff1660e08a0152608081015163ffffffff166101008a015260a081015161ffff166101208a015260c08101516101408a01526101a06136a561366f6136396136018d6135cb61359560e08901516101c06101608501526102408401906104ec565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101808501526104ec565b9061012088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526104ec565b8d610140870151906101c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526104ec565b6101608501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101e08f01526104ec565b6101808401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016102008e0152613329565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016102208b01526104ec565b9260208801528683036040880152613409565b92606081850391015261345a565b8061370a6040926108789594610735565b81602082015201906104ec565b604051906137266020836103fa565b6000808352366020840137565b9190811015611d265760051b0190565b3561087881611904565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461377a5760010190565b612eb7565b801561377a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd120820191821161377a57565b9190820391821161377a57565b906137f0939291614bfe565b919390926137fd82612a80565b9261380783612a80565b94600091825b8851811015613923576000805b8a88848982851061388a575b5050505050156138385760010161380d565b613848610b92611387928b611d2b565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610b926138bc926126ec6138b78873ffffffffffffffffffffffffffffffffffffffff97610b4996613733565b613743565b9116146138cb5760010161381a565b600191506139056138e06138b7838b8b613733565b6138ea888c611d2b565b9073ffffffffffffffffffffffffffffffffffffffff169052565b6139186139118761374d565b968b611d2b565b52388a888489613826565b509097965094939291909460ff811690816000985b8a518a10156139f55760005b8b878210806139ec575b156139df5773ffffffffffffffffffffffffffffffffffffffff613980610b49610b928f6126ec6138b7888f8f613733565b911614613995576139909061374d565b613944565b93996139a560019294939b61377f565b946139c16139b76138b7838b8b613733565b6138ea8d8c611d2b565b6139d46139cd8c61374d565b9b8b611d2b565b525b01989091613938565b50509190986001906139d6565b5085151561394e565b985092509395949750915081613a1e5750505081518103613a1557509190565b80825283529190565b6113879291613a2c916137d7565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906104c3826103c1565b908160209103126100e75760405190613a85826103c1565b51815290565b6108789160e0613b35613b23613aac855161010086526101008601906104ec565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152613b116080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526104ec565b60c085015184820360c08601526104ec565b9201519060e08184039101526104ec565b906020610878928181520190613a8b565b3d15613b82573d90613b688261047a565b91613b7660405193846103fa565b82523d6000602084013e565b606090565b9061ffff610650602092959495604085526040850190613a8b565b90613bab612281565b50613bbc6114886080840151612a1c565b90613bcd6114886060850151612a1c565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201529095906020818060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156115a857600091614023575b5073ffffffffffffffffffffffffffffffffffffffff8116958615613fdf57613cdf613c90898761529c565b96613c99613a60565b50613d2a60a0825192613d0b613cb56114886020840151612a1c565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290968791820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752866103fa565b015193613d1661045b565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff87166040870152606086015273ffffffffffffffffffffffffffffffffffffffff8916608086015260a085015260c0840152613d796104b4565b60e0840152613d8781614099565b15613f0b5750602090613dc89260405193849283927f489a68f200000000000000000000000000000000000000000000000000000000845260048401613b87565b03816000885af160009181613eda575b50613e1c5784613de6613b57565b906113d56040519283927f9fe2f95a000000000000000000000000000000000000000000000000000000008452600484016132c5565b8490935b73ffffffffffffffffffffffffffffffffffffffff831603613e70575b50505051613e68613e4c61046b565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b613e799161529c565b908082108015613ec6575b613e8e5783613e3d565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b50613ed181836137d7565b83511415613e84565b613efd91925060203d602011613f04575b613ef581836103fa565b810190613a6d565b9038613dd8565b503d613eeb565b9050613f16816140f0565b15613f9c57506020613f5591604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301613b46565b03816000885af160009181613f7b575b50613f735784613de6613b57565b849093613e20565b613f9591925060203d602011613f0457613ef581836103fa565b9038613f65565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff821660045260246000fd5b61403c915060203d6020116127555761274781836103fa565b38613c64565b61404b81615352565b9081614087575b8161405b575090565b61087891507f85572ffb0000000000000000000000000000000000000000000000000000000090615447565b9050614092816153e3565b1590614052565b6140a281615352565b90816140de575b816140b2575090565b61087891507f331710310000000000000000000000000000000000000000000000000000000090615447565b90506140e9816153e3565b15906140a9565b6140f981615352565b9081614135575b81614109575090565b61087891507faff2afbf0000000000000000000000000000000000000000000000000000000090615447565b9050614140816153e3565b1590614100565b61415081615352565b908161418c575b81614160575090565b61087891507f7909b5490000000000000000000000000000000000000000000000000000000090615447565b9050614197816153e3565b1590614157565b604051906141ab826103dd565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b9015611d265790565b9060431015611d265760430190565b90821015611d26570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906023116100e75760210190600290565b906043116100e75760230190602090565b90929192836044116100e75783116100e757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff0000000000000000000000000000000000000000000000008116926008811061432d575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110614393575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff000000000000000000000000000000000000000000000000000000000000811692600281106143f9575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b359060208110614439575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff8211176103a057604052606060a0836000815282602082015282604082015282808201528260808201520152565b604080519091906144b983826103fa565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b8281106144f057505050565b6020906144fb614466565b828285010152016144e4565b604051906145166020836103fa565b600080835282815b82811061452a57505050565b602090614535614466565b8282850101520161451e565b9061454a61419e565b91604d8210614b845761458f614589614563848461420b565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603614b5457506145c86145ba6145b46145ae858561422f565b906142f9565b60c01c90565b67ffffffffffffffff168452565b6145ec6145db6145b46145ae8585614240565b67ffffffffffffffff166020850152565b6146106145ff6145b46145ae8585614251565b67ffffffffffffffff166040850152565b61463c61462f6146296146238585614262565b9061435f565b60e01c90565b63ffffffff166060850152565b61465c61464f6146296146238585614273565b63ffffffff166080850152565b61468661467b61467561466f8585614284565b906143c5565b60f01c90565b61ffff1660a0850152565b6146996146938383614295565b9061442b565b60c08401528160431015614b25576146c06146ba6145896145638585614214565b60ff1690565b9081604401838111614b25576146da6128c08286856142a6565b60e086015283811015614af6576146ba6145896145636146fb938786614223565b8201916045830190848211614ac7576128c082604561471c930187866142e1565b61010086015283811015614a98576147406146ba6145896145636045948887614223565b830101916001830190848211614a69576128c0826046614762930187866142e1565b61012086015283811015614a3a576147866146ba6145896145636001948887614223565b830101916001830190848211614a0b576128c08260026147a8930187866142e1565b61014086015260038301928484116149dc576147d86147d161467561466f876001968a896142e1565b61ffff1690565b01019160028301908482116149ad576128c0826147f69287866142e1565b610160860152600483019084821161497e5761467561466f8361481a9388876142e1565b9261ffff829416801560001461491457505050614835614507565b6101808501525b60028201918383116148e557806148606147d161467561466f876002968a896142e1565b0101918383116148b657826128c09185614879946142e1565b6101a0840152036148875790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b60029192945081906149376149276144a8565b966101808a0197885288876154ad565b94909651966149468698611d19565b520101011461483c577fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000006000526113876024906000600452565b906001820180921161377a57565b9190820180921161377a57565b60ff16801561377a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9391614c08613717565b93815180614f45575b505050614c1e9084615b71565b92909293614c6e6002614c4d614c536003614c4d8b67ffffffffffffffff166000526004602052604060002090565b01612b22565b9867ffffffffffffffff166000526004602052604060002090565b614c99614c94614c8c614c848751865190614bc3565b8a5190614bc3565b835190614bc3565b612a80565b9687956000805b8751821015614cd85790614cd06001926138ea614cc0610b92858d611d2b565b91614cca8161374d565b9c611d2b565b018997614ca0565b9750509193969092945060005b8551811015614d1a5780614d14614d01610b926001948a611d2b565b6138ea614d0d8b61374d565b9a8d611d2b565b01614ce5565b50919350919460005b8451811015614d585780614d52614d3f610b9260019489611d2b565b6138ea614d4b8a61374d565b998c611d2b565b01614d23565b5093919592509360005b828110614edc575b5050600090815b818110614e3d5750508152835160005b818110614d9057508452929190565b926000959495915b8351831015614e2e57614dae610b928688611d2b565b73ffffffffffffffffffffffffffffffffffffffff614dd3610b49610b928789611d2b565b911603614e1d57614de39061377f565b90614dfe614df4610b928489611d2b565b6138ea8789611d2b565b60ff8716614e0d575b90614d98565b95614e1790614bd0565b95614e07565b9091614e289061374d565b91614e07565b91509260019095949501614d81565b614e4a610b928286611d2b565b73ffffffffffffffffffffffffffffffffffffffff81168015614ed257600090815b868110614ea6575b5050906001929115614e89575b505b01614d71565b614ea0906138ea614e998761374d565b9688611d2b565b38614e81565b81614eb7610b49610b92848c611d2b565b14614ec457600101614e6c565b506001915081905038614e74565b5050600190614e83565b614eec610b49610b928387611d2b565b15614ef957600101614d62565b5091929360009591955b8351811015614f385780614f32614f1f610b9260019488611d2b565b6138ea614f2b8b61374d565b9a89611d2b565b01614f03565b5093929150933880614d6a565b909192945060018103614ff7575060146060614f6084611d19565b5101515103614fb45781614fab91614f8a6114886060614f82614c1e97611d19565b510151612a1c565b918760a0614fa2614f9a84611d19565b515193611d19565b5101519361596f565b92903880614c11565b6113d56060614fc284611d19565b5101516040519182917f8d666f600000000000000000000000000000000000000000000000000000000083526004830161251f565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361504557565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80519161507d815184614bc3565b9283156151aa5760005b848110615095575050505050565b8181101561518f576150aa610b928286611d2b565b73ffffffffffffffffffffffffffffffffffffffff81168015610bb9576150d083614bb5565b8781106150e257505050600101615087565b8481101561515f5773ffffffffffffffffffffffffffffffffffffffff61510c610b92838a611d2b565b16821461511b576001016150d0565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff61518a610b9261518488856137d7565b89611d2b565b61510c565b6151a5610b9261519f84846137d7565b85611d2b565b6150aa565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b6151dc61419e565b5090565b90604051916151f060c0846103fa565b60848352602083019060a03683375a612ee08111156152415760009161521683926137aa565b82602083519301913090f1903d9060848211615238575b6000908286523e9190565b6084915061522d565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161531b575b506153175782613de6613b57565b9150565b90916020823d60201161534a575b81615336602093836103fa565b810103126153475750519038615309565b80fd5b3d9150615329565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a7000000000000000000000000000000000000000000000000000000006024820152602481526153b66044826103fa565b5191617530fa6000513d826153d7575b50816153d0575090565b9050151590565b602011159150386153c6565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff000000000000000000000000000000000000000000000000000000006024820152602481526153b66044826103fa565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a7000000000000000000000000000000000000000000000000000000008452166024820152602481526153b66044826103fa565b92916154b7614466565b918082101561585a576154d1614589614563848489614223565b600160ff821603614b5457506021820181811161582b576154fa6146938260018601858a6142e1565b8452818110156157fc576146ba61458961456361551893858a614223565b82019160228301908282116157cd576128c08260226155399301858a6142e1565b60208501528181101561579e5761555c6146ba614589614563602294868b614223565b83010191600183019082821161576f576128c082602361557e9301858a6142e1565b604085015281811015615740576155a16146ba614589614563600194868b614223565b830101916001830190828211615711576128c08260026155c39301858a6142e1565b6060850152818110156156e2576155e66146ba614589614563600194868b614223565b83010160018101928284116156b3576128c08460026156079301858a6142e1565b60808501526003810192828411615684576002916156326147d161467561466f88600196898e6142e1565b010101948186116156555761564c926128c09287926142e1565b60a08201529190565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9080601f830112156100e75781516158a081611b41565b926158ae60405194856103fa565b81845260208085019260051b8201019283116100e757602001905b8282106158d65750505090565b6020809183516158e581611904565b8152019101906158c9565b906020828203126100e757815167ffffffffffffffff81116100e7576108789201615889565b95949060019460a09467ffffffffffffffff61596a9573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906104ec565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095929493929091906020838060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9283156115a857600093615af4575b50615a1783614099565b615a51575b5050505050815115615a2c575090565b6108789150614c4d60029167ffffffffffffffff166000526004602052604060002090565b60009496508591615aa773ffffffffffffffffffffffffffffffffffffffff92604051998a97889687957f89720a6200000000000000000000000000000000000000000000000000000000875260048701615916565b0392165afa9182156115a857600092615acf575b50615ac582615e06565b3880808080615a1c565b615aed9192503d806000833e615ae581836103fa565b8101906158f0565b9038615abb565b615b0e91935060203d6020116127555761274781836103fa565b9138615a0d565b90916060828403126100e757815167ffffffffffffffff81116100e75783615b3e918401615889565b92602083015167ffffffffffffffff81116100e757604091615b61918501615889565b92015160ff811681036100e75790565b90803b615bb4575b50615b9a60029167ffffffffffffffff166000526004602052604060002090565b0190615bad615ba7613717565b92612b22565b9190600090565b615bbd81614147565b15615b79576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152906000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa80156115a857600080928192615c68575b50615c3681615e06565b615c3f83615e06565b805115801590615c5c575b615c55575050615b79565b9391925090565b5060ff82161515615c4a565b9150615c8692503d8091833e615c7e81836103fa565b810190615b15565b909138615c2c565b8054821015611d265760005260206000200190600090565b600254811015611d265760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b818110615cef57505060009055565b80615cfc60019285615c8e565b90549060031b1c6000528184016020526000604081205501615ce0565b600081815260036020526040902054615da457600254680100000000000000008110156103a057615d8b615d568260018594016002556002615c8e565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615dff57805490680100000000000000008210156103a05782615de8615d56846001809601855584615c8e565b905580549260005201602052604060002055600190565b5050600090565b80519060005b828110615e1857505050565b6001810180821161377a575b838110615e345750600101615e0c565b73ffffffffffffffffffffffffffffffffffffffff615e538385611d2b565b5116615e65610b49610b928487611d2b565b14615e7257600101615e24565b611387615e82610b928486611d2b565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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

func (_OffRamp *OffRampTransactor) Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "execute", encodedMessage, ccvs, verifierResults)
}

func (_OffRamp *OffRampSession) Execute(encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Execute(&_OffRamp.TransactOpts, encodedMessage, ccvs, verifierResults)
}

func (_OffRamp *OffRampTransactorSession) Execute(encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.Execute(&_OffRamp.TransactOpts, encodedMessage, ccvs, verifierResults)
}

func (_OffRamp *OffRampTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "executeSingleMessage", message, messageId, ccvs, verifierResults)
}

func (_OffRamp *OffRampSession) ExecuteSingleMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteSingleMessage(&_OffRamp.TransactOpts, message, messageId, ccvs, verifierResults)
}

func (_OffRamp *OffRampTransactorSession) ExecuteSingleMessage(message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error) {
	return _OffRamp.Contract.ExecuteSingleMessage(&_OffRamp.TransactOpts, message, messageId, ccvs, verifierResults)
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

	Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error)

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
