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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f6162b138819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a160405161605d9081610254823960805181818161015601526110ff015260a0518181816101b90152611026015260c0518181816101e101528181613fde01526159f7015260e05181818161017d015261298f0152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063176133d1146100d2578063181f5a77146100cd57806320f81c88146100c857806349d8033e146100c35780635215505b146100be5780635643a782146100b95780636b8be52c146100b457806379ba5097146100af5780638da5cb5b146100aa578063e9d68a8e146100a55763f2fde38b146100a057600080fd5b611941565b61175f565b6116df565b6115f6565b610e88565b610aad565b61095a565b610761565b610657565b61052f565b610292565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610384565b828152826020820152826040820152015261025d60405161014b81610384565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760243560443567ffffffffffffffff81116100e757610323903690600401610261565b606435939167ffffffffffffffff85116100e757610348610353953690600401610261565b9490936004016125d7565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176103a057604052565b610355565b60a0810190811067ffffffffffffffff8211176103a057604052565b6020810190811067ffffffffffffffff8211176103a057604052565b6101c0810190811067ffffffffffffffff8211176103a057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103a057604052565b6040519061044a60c0836103fa565b565b6040519061044a6040836103fa565b6040519061044a60a0836103fa565b6040519061044a610100836103fa565b67ffffffffffffffff81116103a057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906104c36020836103fa565b60008252565b60005b8381106104dc5750506000910152565b81810151838201526020016104cc565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610528815180928187528780880191016104c9565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061057081836103fa565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906104ec565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106105f85750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016105eb565b9161065060ff916106426040949796976060875260608701906105da565b9085820360208701526105da565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576107076106b56106af61025d9336906004016105ac565b906146b6565b6106c3610140820151612c7b565b60601c9067ffffffffffffffff81511691610180820151906107018161ffff60a0860151169463ffffffff60806101a0830151519201511690613a63565b93614d8b565b60409391935193849384610624565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561074f57565b610716565b90600482101561074f5752565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356000526007602052602060ff604060002054166107b36040518092610754565bf35b9080602083519182815201916020808360051b8301019401926000915b8383106107e157505050505090565b909192939460208061081d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516104ec565b970193019301919392906107d2565b6108979173ffffffffffffffffffffffffffffffffffffffff82511681526020820151151560208201526080610886610874604085015160a0604086015260a08501906107b5565b606085015184820360608601526105da565b9201519060808184039101526105da565b90565b6040810160408252825180915260206060830193019060005b81811061093a575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106108ef57505050505090565b909192939460208061092b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161082c565b970193019301919392906108e0565b825167ffffffffffffffff168552602094850194909201916001016108b3565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461099581611c56565b906109a360405192836103fa565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06109d082611c56565b0160005b818110610a965750506109e681612cdf565b9060005b818110610a0257505061025d6040519283928361089a565b80610a3a610a21610a14600194615caa565b67ffffffffffffffff1690565b610a2b8387611e0c565b9067ffffffffffffffff169052565b610a7a610a75610a5b610a4d8488611e0c565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b612de2565b610a848287611e0c565b52610a8f8186611e0c565b50016109ea565b602090610aa1612cb3565b828287010152016109d4565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610afc903690600401610261565b90610b056151d3565b6000905b828210610b1257005b610b25610b20838584611b00565b613040565b6020810191610b3f610a14845167ffffffffffffffff1690565b15610e5e57610b81610b68610b68845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610e51575b610bd85760808201949060005b86518051821015610c0257610b68610bb183610bcb93611e0c565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610bd857600101610b96565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610c3a57610b68610bb183610c2d93611e0c565b15610bd857600101610c12565b505095929491909394610c50865182519061521e565b610c65610a5b835167ffffffffffffffff1690565b90610c95610c7b845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610c9f87615cdf565b606085019860005b8a518051821015610cf75790610cbf81602093611e0c565b518051928391012091158015610ce7575b610bd857610ce06001928b615dae565b5001610ca7565b50610cf06130f5565b8214610cd0565b5050976001975093610e0e610e47946003610e3a95610e067f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e99610a14979f610d5460408e610d4d610d9a945160018b0161329d565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610dfc610dbb8d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5160028501613382565b519101613382565b610e2b610e26610a14835167ffffffffffffffff1690565b615d1d565b505167ffffffffffffffff1690565b9260405191829182613416565b0390a20190610b09565b5060808201515115610b89565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610ed79036906004016105ac565b60243567ffffffffffffffff81116100e757610ef7903690600401610261565b9060443567ffffffffffffffff81116100e757610f18903690600401610261565b9094610f2660065460ff1690565b6115cc57610f5a60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006006541617600655565b610f6c610f6786836146b6565b615383565b9461100d6020610fb2610f8a610a148a5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156115c757600091611598575b5061154d57611085611081611077610a5b895167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b6115025761109e610c7b875167ffffffffffffffff1690565b6110c861108160e08901928351602081519101209060019160005201602052604060002054151590565b6114cb5750610100860151601481511480159061149a575b6114635750602086015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff82160361142c57508285036113f85761014086015160148151036113be575090611153913691611c6e565b6020815191012094611179611172876000526007602052604060002090565b5460ff1690565b61118281610745565b80159081156113aa575b501561133857916112186112449261124995946111e16111b68a6000526007602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519586947f176133d10000000000000000000000000000000000000000000000000000000060208701528a8a60248801613708565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103fa565b61538f565b9015611306577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b61129a84611295886000526007602052604060002090565b613522565b6112d66112c460406112b4885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b91836040519485941697169583613929565b0390a46103537fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060065416600655565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff60039261127d565b6113a686866113646040611354835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b600391506113b781610745565b143861118c565b6113f4906040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352600483016125c6565b0390fd5b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004859052602483905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6113f4906040519182917f55216e3100000000000000000000000000000000000000000000000000000000835230600484016134f5565b506114ad6114a782612c7b565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff163014156110e0565b6113f490516040519182917fa50bd147000000000000000000000000000000000000000000000000000000008352600483016125c6565b6113a6611517875167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b6113a6611562875167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b6115ba915060203d6020116115c0575b6115b281836103fa565b8101906134e0565b38611057565b503d6115a8565b611e50565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff811633036116b5577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8116036100e757565b359061044a82611731565b90602061089792818152019061082c565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff6004356117a381611731565b6117ab612cb3565b5016600052600460205260406000206040516117c6816103a5565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c1615156020820152600182018054906117fe82611c56565b9161180c60405193846103fa565b80835260208301916000526020600020916000905b8282106118605761025d8661184f60038a89604085015261184460028201612d81565b606085015201612d81565b60808201526040519182918261174e565b6040516000855461187081612d2e565b80845290600181169081156118e257506001146118aa575b506001928261189c859460209403826103fa565b815201940191019092611821565b6000878152602081209092505b8183106118cc57505081016020016001611888565b60018160209254838688010152019201916118b7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050611888565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561199181611923565b6119996151d3565b16338114611a0b57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b9015611afb578061089791611ab8565b611a89565b90821015611afb576108979160051b810190611ab8565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916020610897938181520191611b68565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611bec575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b63ffffffff8116036100e757565b3561089781611c1e565b3561089781611731565b61ffff8116036100e757565b3561089781611c40565b67ffffffffffffffff81116103a05760051b60200190565b929192611c7a8261047a565b91611c8860405193846103fa565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e75781602061089793359101611c6e565b91909160c0818403126100e757611cd561043b565b9281358452602082013567ffffffffffffffff81116100e75781611cfa918401611ca5565b6020850152604082013567ffffffffffffffff81116100e75781611d1f918401611ca5565b6040850152606082013567ffffffffffffffff81116100e75781611d44918401611ca5565b6060850152608082013567ffffffffffffffff81116100e75781611d69918401611ca5565b608085015260a082013567ffffffffffffffff81116100e757611d8c9201611ca5565b60a0830152565b929190611d9f81611c56565b93611dad60405195866103fa565b602085838152019160051b8101918383116100e75781905b838210611dd3575050505050565b813567ffffffffffffffff81116100e757602091611df48784938701611cc0565b815201910190611dc5565b805115611afb5760200190565b8051821015611afb5760209160051b010190565b90821015611afb57611e379160051b810190611b17565b9091565b908160209103126100e7575161089781611923565b6040513d6000823e3d90fd5b60409073ffffffffffffffffffffffffffffffffffffffff61089795931681528160208201520191611b68565b359061044a82611c1e565b359061044a82611c40565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310611f6a5750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e757602061206b60019386839401908135815261205d61205261203761201c612001611ff189880188611e9f565b60c08b89015260c0880191611b68565b61200e6040880188611e9f565b908783036040890152611b68565b6120296060870187611e9f565b908683036060880152611b68565b6120446080860186611e9f565b908583036080870152611b68565b9260a0810190611e9f565b9160a0818503910152611b68565b980196019493019190611f5a565b6122f161089795939492606083526120a56060840161209783611743565b67ffffffffffffffff169052565b6120c56120b460208301611743565b67ffffffffffffffff166080850152565b6120e56120d460408301611743565b67ffffffffffffffff1660a0850152565b6121016120f460608301611e89565b63ffffffff1660c0850152565b61211d61211060808301611e89565b63ffffffff1660e0850152565b61213861212c60a08301611e94565b61ffff16610100850152565b60c08101356101208401526122c06122b46122756122366121f76121b861217961216560e0890189611e9f565b6101c06101408d01526102208c0191611b68565b612187610100890189611e9f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d0152611b68565b6121c6610120880188611e9f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c0152611b68565b612205610140870187611e9f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b0152611b68565b612244610160860186611e9f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a0152611b68565b612283610180850185611eef565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152611f42565b916101a0810190611e9f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa085840301610200860152611b68565b9360208201526040818503910152611b68565b604051906040820182811067ffffffffffffffff8211176103a05760405260006020838281520152565b9061233882611c56565b61234560405191826103fa565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06123738294611c56565b019060005b82811061238457505050565b60209061238f612304565b82828501015201612378565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd12082019182116123f757565b61239b565b919082039182116123f757565b801515036100e757565b90916060828403126100e757815161242a81612409565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e7578151916124598361047a565b9161246760405193846103fa565b838352602084830101116100e75760409261248891602080850191016104c9565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806125056124d1604084015160a060c08801526101208701906104ec565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526104ec565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b81811061258e5750505061ffff909516602083015261044a92916060916125729063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101612544565b9060206108979281815201906104ec565b9293909194303303612c51576000956125f4610180860186611a35565b9050612ba1575b6126866126186114a7612612610140890189611b17565b90611bb8565b9161264b8361262b6101a08a018a611b17565b905061264561263c60808c01611c2c565b63ffffffff1690565b90613a63565b97889160a06126598a611c36565b866126808c61267861266f610180830183611a35565b96909201611c4c565b943691611d93565b91613bb8565b91909760005b8951811015612876578060206126af610b68610b68610bb18f966126fb97611e0c565b6126c46126bc8489611e0c565b51898b611e20565b91906040518096819482937fc3a7ded600000000000000000000000000000000000000000000000000000000845260048401611ba7565b03915afa9182156115c757600092612846575b5073ffffffffffffffffffffffffffffffffffffffff8216156127e9576127406127388287611e0c565b518789611e20565b929073ffffffffffffffffffffffffffffffffffffffff82163b156100e7576000918a8373ffffffffffffffffffffffffffffffffffffffff8e6127b3604051998a97889687947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601612079565b0393165af19182156115c7576001926127ce575b500161268c565b806127dd60006127e3936103fa565b806100dc565b386127c7565b84612811888861280a8f95612804610bb1826113f499611e0c565b95611e0c565b5191611e20565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501611e5c565b61286891925060203d811161286f575b61286081836103fa565b810190611e3b565b903861270e565b503d612856565b50949750949592505050612898612891610180860186611a35565b905061232e565b916128a7610180860186611a35565b9050612a3a575b50612a34576129b893600061295f6080866128ec610b686128d2610a5b879b611c36565b5473ffffffffffffffffffffffffffffffffffffffff1690565b956129506128f983611c36565b61293d61294461290d610120870187611b17565b61291b6101a0890189611b17565b949095602061292861045b565b9e8f908152019067ffffffffffffffff169052565b3691611c6e565b60408a01523691611c6e565b60608701528286015201611c2c565b91604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f0000000000000000000000000000000000000000000000000000000000000000906004860161248e565b03925af19081156115c757600090600092612a0d575b50156129d75750565b6113f4906040519182917f0a8d6e8c000000000000000000000000000000000000000000000000000000008352600483016125c6565b9050612a2c91503d806000833e612a2481836103fa565b810190612413565b5090386129ce565b50505050565b612a94612a54612a4e610180880188611a35565b90611aeb565b612a62610120880188611b17565b612a8e612a6e8a611c36565b92612a86612a7e60a08d01611c4c565b953690611cc0565b923691611c6e565b90613f49565b909173ffffffffffffffffffffffffffffffffffffffff612acf6114a7612612612ac5612a4e6101808d018d611a35565b6080810190611b17565b921673ffffffffffffffffffffffffffffffffffffffff831614600014612b0d575050612afb83611dff565b52612b0582611dff565b505b386128ae565b612b56612b3b612b5c93612b35865173ffffffffffffffffffffffffffffffffffffffff1690565b90613977565b935173ffffffffffffffffffffffffffffffffffffffff1690565b926123fc565b612b83612b6761044c565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b6020820152612b9183611dff565b52612b9b82611dff565b50612b07565b95506014612bc3612bb9612a4e610180880188611a35565b6060810190611b17565b905003612c0757612c01612be76114a7612612612ac5612a4e6101808a018a611a35565b612b356114a7612612612bb9612a4e6101808b018b611a35565b956125fb565b612c1b612bb9612a4e610180870187611a35565b906113f46040519283927f8d666f6000000000000000000000000000000000000000000000000000000000845260048401611ba7565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611bec575050565b60405190612cc0826103a5565b6060608083600081526000602082015282604082015282808201520152565b90612ce982611c56565b612cf660405191826103fa565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612d248294611c56565b0190602036910137565b90600182811c92168015612d77575b6020831014612d4857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612d3d565b906040519182815491828252602082019060005260206000209260005b818110612db357505061044a925003836103fa565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612d9e565b90604051612def816103a5565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c161515602083015260018101805490612e2982611c56565b91612e3760405193846103fa565b80835260208301916000526020600020916000905b828210612e8157505050506003608092612e7c926040860152612e7160028201612d81565b606086015201612d81565b910152565b60405160008554612e9181612d2e565b8084529060018116908115612f035750600114612ecb575b5060019282612ebd859460209403826103fa565b815201940191019092612e4c565b6000878152602081209092505b818310612eed57505081016020016001612ea9565b6001816020925483868801015201920191612ed8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050612ea9565b359061044a82611923565b359061044a82612409565b9080601f830112156100e7578135612f7181611c56565b92612f7f60405194856103fa565b81845260208085019260051b820101918383116100e75760208201905b838210612fab57505050505090565b813567ffffffffffffffff81116100e757602091612fce87848094880101611ca5565b815201910190612f9c565b9080601f830112156100e7578135612ff081611c56565b92612ffe60405194856103fa565b81845260208085019260051b8201019283116100e757602001905b8282106130265750505090565b60208091833561303581611923565b815201910190613019565b60c0813603126100e75761305261043b565b9061305c81612f44565b825261306a60208201611743565b602083015261307b60408201612f4f565b6040830152606081013567ffffffffffffffff81116100e7576130a19036908301612f5a565b6060830152608081013567ffffffffffffffff81116100e7576130c79036908301612fd9565b608083015260a08101359067ffffffffffffffff82116100e7576130ed91369101612fd9565b60a082015290565b604051602081019060008252602081526131106040826103fa565b51902090565b818110613121575050565b60008155600101613116565b9190601f811161313c57505050565b61044a926000526020600020906020601f840160051c83019310613168575b601f0160051c0190613116565b909150819061315b565b919091825167ffffffffffffffff81116103a05761319a816131948454612d2e565b8461312d565b6020601f82116001146131f85781906131e99394956000926131ed575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0151905038806131b7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061322b84600052602060002090565b9160005b8181106132855750958360019596971061324e575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613244565b9192602060018192868b01518155019401920161322f565b8151916801000000000000000083116103a05781548383558084106132ff575b5060206132d1910191600052602060002090565b6000915b8383106132e25750505050565b60016020826132f383945186613172565b019201920191906132d5565b8260005283602060002091820191015b81811061331c57506132bd565b8061332960019254612d2e565b80613336575b500161330f565b601f8111831461334c5750600081555b3861332f565b6133709083601f61336285600052602060002090565b920160051c82019101613116565b60008181526020812081835555613346565b81519167ffffffffffffffff83116103a0576801000000000000000083116103a05760209082548484558085106133f9575b500190600052602060002060005b8381106133cf5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016133c2565b613410908460005285846000209182019101613116565b386133b4565b90610897916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a06134ad61347a606085015160c0608086015260e08501906107b5565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526105da565b9201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526105da565b908160209103126100e7575161089781612409565b60409073ffffffffffffffffffffffffffffffffffffffff610897949316815281602082015201906104ec565b90600481101561074f5760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061358557505050505090565b909192939460208061362a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a06136196136076135f56135e38887015160c08a88015260c08701906104ec565b604087015186820360408801526104ec565b606086015185820360608701526104ec565b608085015184820360808601526104ec565b9201519060a08184039101526104ec565b97019301930191939290613576565b9160209082815201919060005b8181106136535750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561367c81611923565b168152019401929101613646565b90602083828152019260208260051b82010193836000925b8484106136b25750505050505090565b9091929394956020806136f8837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030188526136f28b88611e9f565b90611b68565b98019401940192949391906136a2565b949293610897969461390861391b94936080895261373360808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a0152606081015163ffffffff1660e08a0152608081015163ffffffff166101008a015260a081015161ffff166101208a015260c08101516101408a01526101a06138d561389f6138696138318d6137fb6137c560e08901516101c06101608501526102408401906104ec565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101808501526104ec565b9061012088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526104ec565b8d610140870151906101c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526104ec565b6101608501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101e08f01526104ec565b6101808401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016102008e0152613559565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016102208b01526104ec565b9260208801528683036040880152613639565b92606081850391015261368a565b8061393a6040926108979594610754565b81602082015201906104ec565b3d15613972573d906139588261047a565b9161396660405193846103fa565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181613a2c575b50613a2857826139f2613947565b906113f46040519283927f9fe2f95a000000000000000000000000000000000000000000000000000000008452600484016134f5565b9150565b90916020823d602011613a5b575b81613a47602093836103fa565b81010312613a5857505190386139e4565b80fd5b3d9150613a3a565b15919082613ae1575b508115613ad7575b8115613a7e575090565b9050613aaa7f85572ffb0000000000000000000000000000000000000000000000000000000082615f30565b9081613ac5575b81613abb57501590565b6110819150615ed0565b9050613ad081615e0a565b1590613ab1565b803b159150613a74565b15915038613a6c565b60405190613af96020836103fa565b6000808352366020840137565b60408051909190613b1783826103fa565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001366020840137565b9190811015611afb5760051b0190565b3561089781611923565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146123f75760010190565b80156123f7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b90613bc7949596939291614d8b565b91939092613bd482612cdf565b92613bde83612cdf565b94600091825b8851811015613cfa576000805b8a888489828510613c61575b505050505015613c0f57600101613be4565b613c1f610bb16113a6928b611e0c565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610bb1613c9392612804613c8e8873ffffffffffffffffffffffffffffffffffffffff97610b6896613b46565b613b56565b911614613ca257600101613bf1565b60019150613cdc613cb7613c8e838b8b613b46565b613cc1888c611e0c565b9073ffffffffffffffffffffffffffffffffffffffff169052565b613cef613ce887613b60565b968b611e0c565b52388a888489613bfd565b509097965094939291909460ff811690816000985b8a518a1015613dcc5760005b8b87821080613dc3575b15613db65773ffffffffffffffffffffffffffffffffffffffff613d57610b68610bb18f612804613c8e888f8f613b46565b911614613d6c57613d6790613b60565b613d1b565b9399613d7c60019294939b613b8d565b94613d98613d8e613c8e838b8b613b46565b613cc18d8c611e0c565b613dab613da48c613b60565b9b8b611e0c565b525b01989091613d0f565b5050919098600190613dad565b50851515613d25565b985092509395949750915081613df55750505081518103613dec57509190565b80825283529190565b6113a69291613e03916123fc565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906104c3826103c1565b908160209103126100e75760405190613e5c826103c1565b51815290565b6108979160e0613f0c613efa613e83855161010086526101008601906104ec565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152613ee86080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526104ec565b60c085015184820360c08601526104ec565b9201519060e08184039101526104ec565b906020610897928181520190613e62565b9061ffff610650602092959495604085526040850190613e62565b92939193613f55612304565b50613f666114a76080860151612c7b565b91613f776114a76060870151612c7b565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909690956020878060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9687156115c7576000976142f2575b5073ffffffffffffffffffffffffffffffffffffffff87169485156142ae57906140ec6140849261403e613e37565b506140cf60a08451946140b061405a6114a76020840151612c7b565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290988991820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018952886103fa565b0151956140bb61046a565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff8816608084015260a083015260c082015261411f6104b4565b60e082015261412d8561544b565b156141d757906141719260209260006040518096819582947f489a68f200000000000000000000000000000000000000000000000000000000845260048401613f2e565b03925af1600091816141a6575b5061418c57836139f2613947565b929091925b5161419d612b6761044c565b60208201529190565b6141c991925060203d6020116141d0575b6141c181836103fa565b810190613e44565b903861417e565b503d6141b7565b90506141e2846154a1565b1561426a576142256000926020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301613f1d565b03925af160009181614249575b5061424057836139f2613947565b92909192614191565b61426391925060203d6020116141d0576141c181836103fa565b9038614232565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff881660045260246000fd5b61430c91975060203d60201161286f5761286081836103fa565b953861400f565b60405190614320826103dd565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b9015611afb5790565b9060431015611afb5760430190565b90821015611afb570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906023116100e75760210190600290565b906043116100e75760230190602090565b90929192836044116100e75783116100e757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff000000000000000000000000000000000000000000000000811692600881106144a2575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110614508575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff0000000000000000000000000000000000000000000000000000000000008116926002811061456e575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b3590602081106145ae575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff8211176103a057604052606060a0836000815282602082015282604082015282808201528260808201520152565b6040805190919061462e83826103fa565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b82811061466557505050565b6020906146706145db565b82828501015201614659565b6040519061468b6020836103fa565b600080835282815b82811061469f57505050565b6020906146aa6145db565b82828501015201614693565b906146bf614313565b91604d8210614cf9576147046146fe6146d88484614380565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603614cc9575061473d61472f61472961472385856143a4565b9061446e565b60c01c90565b67ffffffffffffffff168452565b61476161475061472961472385856143b5565b67ffffffffffffffff166020850152565b61478561477461472961472385856143c6565b67ffffffffffffffff166040850152565b6147b16147a461479e61479885856143d7565b906144d4565b60e01c90565b63ffffffff166060850152565b6147d16147c461479e61479885856143e8565b63ffffffff166080850152565b6147fb6147f06147ea6147e485856143f9565b9061453a565b60f01c90565b61ffff1660a0850152565b61480e614808838361440a565b906145a0565b60c08401528160431015614c9a5761483561482f6146fe6146d88585614389565b60ff1690565b9081604401838111614c9a5761484f61293d82868561441b565b60e086015283811015614c6b5761482f6146fe6146d8614870938786614398565b8201916045830190848211614c3c5761293d82604561489193018786614456565b61010086015283811015614c0d576148b561482f6146fe6146d86045948887614398565b830101916001830190848211614bde5761293d8260466148d793018786614456565b61012086015283811015614baf576148fb61482f6146fe6146d86001948887614398565b830101916001830190848211614b805761293d82600261491d93018786614456565b6101408601526003830192848411614b515761494d6149466147ea6147e4876001968a89614456565b61ffff1690565b0101916002830190848211614b225761293d8261496b928786614456565b6101608601526004830190848211614af3576147ea6147e48361498f938887614456565b9261ffff8294168015600014614a89575050506149aa61467c565b6101808501525b6002820191838311614a5a57806149d56149466147ea6147e4876002968a89614456565b010191838311614a2b578261293d91856149ee94614456565b6101a0840152036149fc5790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b6002919294508190614aac614a9c61461d565b966101808a0197885288876154cb565b9490965196614abb8698611dff565b52010101146149b1577fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000006000526113a66024906000600452565b90600182018092116123f757565b919082018092116123f757565b8054821015611afb5760005260206000200190600090565b60ff1680156123f7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b94929391606091600093614d9d613aea565b96835180615154575b50501561513d5750511561512b575050614dbe613aea565b614dc6613aea565b916000935b6002614e14614df96003614df38a67ffffffffffffffff166000526004602052604060002090565b01612d81565b9767ffffffffffffffff166000526004602052604060002090565b0195614e2e614e268551845190614d38565b825190614d38565b90614e43614e3e89548094614d38565b612cdf565b95600098895b8a8851821015614e815790614e7b600192613cc1614e74614e6d610bb1868f611e0c565b929f613b60565b9e8d611e0c565b01614e49565b50509193955091939660005b898751821015614ec55790614ebf600192613cc1614eb8614eb1610bb1868e611e0c565b929e613b60565b9d8c611e0c565b01614e8d565b50509193969092945060005b8551811015614f065780614f00614eed610bb16001948a611e0c565b613cc1614ef98d613b60565b9c8b611e0c565b01614ed1565b509350939195909460005b8281106150a7575b50509091929350600090815b8181106150085750508152835160005b818110614f4457508452929190565b9391949260005b8651811015614ffb57614f61610bb18787611e0c565b73ffffffffffffffffffffffffffffffffffffffff614f86610b68610bb1858c611e0c565b911614614f9e57614f9690613b60565b905b90614f4b565b5090614fa990613b8d565b90614fc4614fba610bb18487611e0c565b613cc18787611e0c565b60ff8316614feb575b84821115614fdd57600090614f98565b929491936001905b01614f35565b91614ff590614d5d565b91614fcd565b5092949193600190614fe5565b615015610bb18286611e0c565b73ffffffffffffffffffffffffffffffffffffffff8116801561509d57600090815b868110615071575b5050906001929115615054575b505b01614f25565b61506b90613cc161506487613b60565b9688611e0c565b3861504c565b81615082610b68610bb1848c611e0c565b1461508f57600101615037565b50600191508190503861503f565b505060019061504e565b6150b7610b68610bb18387611e0c565b156150c457600101614f11565b50909192939460005b8281106150df57869594939250614f19565b806151256151126150f260019486614d45565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b613cc161511e88613b60565b9789611e0c565b016150cd565b91615137939193613b06565b91614dcb565b91505061514b915084615b6e565b93909290614dcb565b90919750600181036151a6575061519e9061517d6114a78661517587611dff565b510151612c7b565b9061518785611dff565b51518a60a061519588611dff565b5101519361598d565b953880614da6565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6001541633036151f457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80519161522c815184614d38565b9283156153595760005b848110615244575050505050565b8181101561533e57615259610bb18286611e0c565b73ffffffffffffffffffffffffffffffffffffffff81168015610bd85761527f83614d2a565b87811061529157505050600101615236565b8481101561530e5773ffffffffffffffffffffffffffffffffffffffff6152bb610bb1838a611e0c565b1682146152ca5760010161527f565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff615339610bb161533388856123fc565b89611e0c565b6152bb565b615354610bb161534e84846123fc565b85611e0c565b615259565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b61538b614313565b5090565b906040519161539f60c0846103fa565b60848352602083019060a03683375a612ee08111156153f0576000916153c583926123ca565b82602083519301913090f1903d90608482116153e7575b6000908286523e9190565b608491506153dc565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6154757f331710310000000000000000000000000000000000000000000000000000000082615f30565b908161548f575b81615485575090565b6108979150615ed0565b905061549a81615e0a565b159061547c565b6154757faff2afbf0000000000000000000000000000000000000000000000000000000082615f30565b92916154d56145db565b9180821015615878576154ef6146fe6146d8848489614398565b600160ff821603614cc9575060218201818111615849576155186148088260018601858a614456565b84528181101561581a5761482f6146fe6146d861553693858a614398565b82019160228301908282116157eb5761293d8260226155579301858a614456565b6020850152818110156157bc5761557a61482f6146fe6146d8602294868b614398565b83010191600183019082821161578d5761293d82602361559c9301858a614456565b60408501528181101561575e576155bf61482f6146fe6146d8600194868b614398565b83010191600183019082821161572f5761293d8260026155e19301858a614456565b6060850152818110156157005761560461482f6146fe6146d8600194868b614398565b83010160018101928284116156d15761293d8460026156259301858a614456565b608085015260038101928284116156a2576002916156506149466147ea6147e488600196898e614456565b010101948186116156735761566a9261293d928792614456565b60a08201529190565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9080601f830112156100e75781516158be81611c56565b926158cc60405194856103fa565b81845260208085019260051b8201019283116100e757602001905b8282106158f45750505090565b60208091835161590381611923565b8152019101906158e7565b906020828203126100e757815167ffffffffffffffff81116100e75761089792016158a7565b95949060019460a09467ffffffffffffffff6159889573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906104ec565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095909291906020848060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9384156115c757600094615af1575b50615a328461544b565b615a50575b505050505050805115615a475790565b50610897613b06565b600095965090615aa573ffffffffffffffffffffffffffffffffffffffff92604051988997889687957f89720a6200000000000000000000000000000000000000000000000000000000875260048701615934565b0392165afa9081156115c757600091615ace575b50615ac381615f92565b388080808080615a37565b615aeb91503d806000833e615ae381836103fa565b81019061590e565b38615ab9565b615b0b91945060203d60201161286f5761286081836103fa565b9238615a28565b90916060828403126100e757815167ffffffffffffffff81116100e75783615b3b9184016158a7565b92602083015167ffffffffffffffff81116100e757604091615b5e9185016158a7565b92015160ff811681036100e75790565b90615b997f7909b5490000000000000000000000000000000000000000000000000000000082615f30565b80615c9a575b80615c8b575b615bc3575b5050615bb4613b06565b90615bbd613aea565b90600090565b6040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9290921660048301526000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa80156115c757600080928192615c65575b50615c3881615f92565b615c4183615f92565b805115801590615c59575b615c565750615baa565b92565b5060ff82161515615c4c565b9150615c8392503d8091833e615c7b81836103fa565b810190615b12565b909138615c2e565b50615c9581615ed0565b615ba5565b50615ca481615e0a565b15615b9f565b600254811015611afb5760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b818110615cf357505060009055565b80615d0060019285614d45565b90549060031b1c6000528184016020526000604081205501615ce4565b600081815260036020526040902054615da857600254680100000000000000008110156103a057615d8f615d5a8260018594016002556002614d45565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615e0357805490680100000000000000008210156103a05782615dec615d5a846001809601855584614d45565b905580549260005201602052604060002055600190565b5050600090565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252615e6a6044836103fa565b6179185a10615ea6576020926000925191617530fa6000513d82615e9a575b5081615e93575090565b9050151590565b60201115915038615e89565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252615e6a6044836103fa565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252615e6a6044836103fa565b80519060005b828110615fa457505050565b600181018082116123f7575b838110615fc05750600101615f98565b73ffffffffffffffffffffffffffffffffffffffff615fdf8385611e0c565b5116615ff1610b68610bb18487611e0c565b14615ffe57600101615fb0565b6113a661600e610bb18486611e0c565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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
