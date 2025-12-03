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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f6161da38819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615f86908161025482396080518181816101560152611163015260a0518181816101b9015261108a015260c0518181816101e101528181613ea60152615a91015260e05181818161017d0152612ae30152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063176133d1146100d2578063181f5a77146100cd57806320f81c88146100c857806349d8033e146100c35780635215505b146100be5780635643a782146100b95780636b8be52c146100b457806379ba5097146100af5780638da5cb5b146100aa578063e9d68a8e146100a55763f2fde38b146100a057600080fd5b61199b565b6117b9565b611739565b611650565b610eec565b610b11565b6109be565b6107c5565b610657565b61052f565b610292565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610384565b828152826020820152826040820152015261025d60405161014b81610384565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760243560443567ffffffffffffffff81116100e757610323903690600401610261565b606435939167ffffffffffffffff85116100e757610348610353953690600401610261565b9490936004016125f8565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176103a057604052565b610355565b60a0810190811067ffffffffffffffff8211176103a057604052565b6020810190811067ffffffffffffffff8211176103a057604052565b6101c0810190811067ffffffffffffffff8211176103a057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103a057604052565b6040519061044a60c0836103fa565b565b6040519061044a60a0836103fa565b6040519061044a610100836103fa565b6040519061044a6040836103fa565b67ffffffffffffffff81116103a057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906104c36020836103fa565b60008252565b60005b8381106104dc5750506000910152565b81810151838201526020016104cc565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610528815180928187528780880191016104c9565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061057081836103fa565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906104ec565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106105f85750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016105eb565b9161065060ff916106426040949796976060875260608701906105da565b9085820360208701526105da565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106a96106af9136906004016105ac565b90614705565b61018081019081518051610716575b61025d61070784846106d8815167ffffffffffffffff1690565b9161070160a06106f56106ef610140860151612bed565b60601c90565b925193015161ffff1690565b92614dc2565b60409391935193849384610624565b6060610723601492611cc5565b510151510361073257386106be565b61077660606107418451611cc5565b5101516040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352600483016125e7565b0390fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600411156107b357565b61077a565b9060048210156107b35752565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356000526007602052602060ff6040600020541661081760405180926107b8565bf35b9080602083519182815201916020808360051b8301019401926000915b83831061084557505050505090565b9091929394602080610881837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516104ec565b97019301930191939290610836565b6108fb9173ffffffffffffffffffffffffffffffffffffffff825116815260208201511515602082015260806108ea6108d8604085015160a0604086015260a0850190610819565b606085015184820360608601526105da565b9201519060808184039101526105da565b90565b6040810160408252825180915260206060830193019060005b81811061099e575050506020818303910152815180825260208201916020808360051b8301019401926000915b83831061095357505050505090565b909192939460208061098f837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610890565b97019301930191939290610944565b825167ffffffffffffffff16855260209485019490920191600101610917565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576002546109f981611b9a565b90610a0760405192836103fa565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610a3482611b9a565b0160005b818110610afa575050610a4a81611bb2565b9060005b818110610a6657505061025d604051928392836108fe565b80610a9e610a85610a78600194615d5b565b67ffffffffffffffff1690565b610a8f8387611cd2565b9067ffffffffffffffff169052565b610ade610ad9610abf610ab18488611cd2565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b612d05565b610ae88287611cd2565b52610af38186611cd2565b5001610a4e565b602090610b05612c25565b82828701015201610a38565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610b60903690600401610261565b90610b6961518f565b6000905b828210610b7657005b610b89610b84838584611c30565b612f63565b6020810191610ba3610a78845167ffffffffffffffff1690565b15610ec257610be5610bcc610bcc845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610eb5575b610c3c5760808201949060005b86518051821015610c6657610bcc610c1583610c2f93611cd2565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610c3c57600101610bfa565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610c9e57610bcc610c1583610c9193611cd2565b15610c3c57600101610c76565b505095929491909394610cb486518251906151da565b610cc9610abf835167ffffffffffffffff1690565b90610cf9610cdf845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610d0387615d90565b606085019860005b8a518051821015610d5b5790610d2381602093611cd2565b518051928391012091158015610d4b575b610c3c57610d446001928b615e5f565b5001610d0b565b50610d54613018565b8214610d34565b5050976001975093610e72610eab946003610e9e95610e6a7f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e99610a78979f610db860408e610db1610dfe945160018b016131ef565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610e60610e1f8d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b51600285016132d4565b5191016132d4565b610e8f610e8a610a78835167ffffffffffffffff1690565b615dce565b505167ffffffffffffffff1690565b9260405191829182613368565b0390a20190610b6d565b5060808201515115610bed565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610f3b9036906004016105ac565b60243567ffffffffffffffff81116100e757610f5b903690600401610261565b9060443567ffffffffffffffff81116100e757610f7c903690600401610261565b9094610f8a60065460ff1690565b61162657610fbe60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006006541617600655565b610fd0610fcb8683614705565b61533f565b946110716020611016610fee610a788a5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611621576000916115f2575b506115a7576110e96110e56110db610abf895167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b61155c57611102610cdf875167ffffffffffffffff1690565b61112c6110e560e08901928351602081519101209060019160005201602052604060002054151590565b611525575061010086015160148151148015906114fa575b6114c35750602086015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff82160361148c57508285036114585761014086015160148151036114225750906111b7913691611d06565b60208151910120946111dd6111d6876000526007602052604060002090565b5460ff1690565b6111e6816107a9565b801590811561140e575b501561139c579161127c6112a8926112ad959461124561121a8a6000526007602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519586947f176133d10000000000000000000000000000000000000000000000000000000060208701528a8a6024880161365a565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103fa565b61534b565b901561136a577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b6112fe846112f9886000526007602052604060002090565b613474565b61133a6113286040611318885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b9183604051948594169716958361387b565b0390a46103537fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060065416600655565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6003926112e1565b61140a86866113c860406113b8835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b6003915061141b816107a9565b14386111f0565b610776906040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352600483016125e7565b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004859052602483905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b610776906040519182917f55216e310000000000000000000000000000000000000000000000000000000083523060048401613447565b506115076106ef82612bed565b73ffffffffffffffffffffffffffffffffffffffff16301415611144565b61077690516040519182917fa50bd147000000000000000000000000000000000000000000000000000000008352600483016125e7565b61140a611571875167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b61140a6115bc875167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611614915060203d60201161161a575b61160c81836103fa565b810190613432565b386110bb565b503d611602565b611ec7565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff8116330361170f577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8116036100e757565b359061044a8261178b565b9060206108fb928181520190610890565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff6004356117fd8161178b565b611805612c25565b501660005260046020526040600020604051611820816103a5565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c16151560208201526001820180549061185882611b9a565b9161186660405193846103fa565b80835260208301916000526020600020916000905b8282106118ba5761025d866118a960038a89604085015261189e60028201612ca4565b606085015201612ca4565b6080820152604051918291826117a8565b604051600085546118ca81612c51565b808452906001811690811561193c5750600114611904575b50600192826118f6859460209403826103fa565b81520194019101909261187b565b6000878152602081209092505b818310611926575050810160200160016118e2565b6001816020925483868801015201920191611911565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506118e2565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff6004356119eb8161197d565b6119f361518f565b16338114611a6557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611b14575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b67ffffffffffffffff81116103a05760051b60200190565b90611bbc82611b9a565b611bc960405191826103fa565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611bf78294611b9a565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611c705760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b611c01565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9160206108fb938181520191611c75565b805115611c705760200190565b8051821015611c705760209160051b010190565b356108fb8161178b565b61ffff8116036100e757565b356108fb81611cf0565b929192611d128261047a565b91611d2060405193846103fa565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e7578160206108fb93359101611d06565b91909160c0818403126100e757611d6d61043b565b9281358452602082013567ffffffffffffffff81116100e75781611d92918401611d3d565b6020850152604082013567ffffffffffffffff81116100e75781611db7918401611d3d565b6040850152606082013567ffffffffffffffff81116100e75781611ddc918401611d3d565b6060850152608082013567ffffffffffffffff81116100e75781611e01918401611d3d565b608085015260a082013567ffffffffffffffff81116100e757611e249201611d3d565b60a0830152565b929190611e3781611b9a565b93611e4560405195866103fa565b602085838152019160051b8101918383116100e75781905b838210611e6b575050505050565b813567ffffffffffffffff81116100e757602091611e8c8784938701611d58565b815201910190611e5d565b90821015611c7057611eae9160051b810190611a8f565b9091565b908160209103126100e757516108fb8161197d565b6040513d6000823e3d90fd5b60409073ffffffffffffffffffffffffffffffffffffffff6108fb95931681528160208201520191611c75565b63ffffffff8116036100e757565b359061044a82611f00565b359061044a82611cf0565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310611fef5750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e75760206120f06001938683940190813581526120e26120d76120bc6120a161208661207689880188611f24565b60c08b89015260c0880191611c75565b6120936040880188611f24565b908783036040890152611c75565b6120ae6060870187611f24565b908683036060880152611c75565b6120c96080860186611f24565b908583036080870152611c75565b9260a0810190611f24565b9160a0818503910152611c75565b980196019493019190611fdf565b6123766108fb959394926060835261212a6060840161211c8361179d565b67ffffffffffffffff169052565b61214a6121396020830161179d565b67ffffffffffffffff166080850152565b61216a6121596040830161179d565b67ffffffffffffffff1660a0850152565b61218661217960608301611f0e565b63ffffffff1660c0850152565b6121a261219560808301611f0e565b63ffffffff1660e0850152565b6121bd6121b160a08301611f19565b61ffff16610100850152565b60c08101356101208401526123456123396122fa6122bb61227c61223d6121fe6121ea60e0890189611f24565b6101c06101408d01526102208c0191611c75565b61220c610100890189611f24565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d0152611c75565b61224b610120880188611f24565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c0152611c75565b61228a610140870187611f24565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b0152611c75565b6122c9610160860186611f24565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a0152611c75565b612308610180850185611f74565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152611fc7565b916101a0810190611f24565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa085840301610200860152611c75565b9360208201526040818503910152611c75565b604051906040820182811067ffffffffffffffff8211176103a05760405260006020838281520152565b906123bd82611b9a565b6123ca60405191826103fa565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06123f88294611b9a565b019060005b82811061240957505050565b602090612414612389565b828285010152016123fd565b356108fb81611f00565b801515036100e757565b90916060828403126100e757815161244b8161242a565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e75781519161247a8361047a565b9161248860405193846103fa565b838352602084830101116100e7576040926124a991602080850191016104c9565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806125266124f2604084015160a060c08801526101208701906104ec565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526104ec565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106125af5750505061ffff909516602083015261044a92916060916125939063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101612565565b9060206108fb9281815201906104ec565b9590949294303303612bc35761261e6106ef6126186101408a018a611a8f565b90611ae0565b94610180880194612639612632878b611b46565b9050611bb2565b9660005b612647888c611b46565b905081101561271957601461267461266a836126648f8d90611b46565b90611c30565b6060810190611a8f565b9050036126d257806126c18c6126bb6106ef61261861266a868f806126b56106ef6126186126ab60019e612664612664978e611b46565b6080810190611a8f565b97611b46565b906138c9565b6126cb828c611cd2565b520161263d565b61266a6126e3916126648a8e611b46565b906107766040519283927f8d666f6000000000000000000000000000000000000000000000000000000000845260048401611cb4565b509193969092949798959861272d87611ce6565b906127388b89611b46565b60a08a019b9186906127498e611cfc565b92369061275592611e2b565b9061275f95613a82565b98909960005b8b51811015612922576127d460208c61279d8f8561278f610bcc610bcc610c158461279596611cd2565b93611cd2565b518a8c611e97565b91906040518095819482937fc3a7ded600000000000000000000000000000000000000000000000000000000845260048401611cb4565b03915afa80156116215773ffffffffffffffffffffffffffffffffffffffff916000916128f4575b50169081156128975761281a612812828e611cd2565b51888a611e97565b9290813b156100e7576000918b838e612862604051988996879586947fbff0ec1d000000000000000000000000000000000000000000000000000000008652600486016120fe565b03925af19182156116215760019261287c575b5001612765565b8061288b6000612891936103fa565b806100dc565b38612875565b8b6128bf898f6128b8856128b2610c15610776988f95611cd2565b95611cd2565b5191611e97565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501611ed3565b612915915060203d811161291b575b61290d81836103fa565b810190611eb2565b386127fc565b503d612903565b50945094975094959150965061294261293b8387611b46565b90506123b3565b9660005b6129508488611b46565b90508110156129d557806129b961296d600193612664888c611b46565b88886129966129b38d6129ab6129a28961299c61298e610120860186611a8f565b979095611ce6565b97611cfc565b97611cd2565b51963690611d58565b923691611d06565b90613e10565b6129c3828c611cd2565b526129ce818b611cd2565b5001612946565b50935095949150506101a08301906129ed8285611a8f565b90501580612baa575b8015612ba1575b8015612b8f575b612b88576000612ab3608086612b0c98612aa4612a44610bcc612a2a610abf899d611ce6565b5473ffffffffffffffffffffffffffffffffffffffff1690565b97612a91612a98612a5486611ce6565b92612a6f612a66610120890189611a8f565b91909289611a8f565b9490956020612a7c61044c565b9e8f908152019067ffffffffffffffff169052565b3691611d06565b60408a01523691611d06565b60608701528286015201612420565b91604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f000000000000000000000000000000000000000000000000000000000000000090600486016124af565b03925af190811561162157600090600092612b61575b5015612b2b5750565b610776906040519182917f0a8d6e8c000000000000000000000000000000000000000000000000000000008352600483016125e7565b9050612b8091503d806000833e612b7881836103fa565b810190612434565b509038612b22565b5050505050565b50612b9c6110e584614206565b612a04565b50823b156129fd565b5063ffffffff612bbc60808601612420565b16156129f6565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611b14575050565b60405190612c32826103a5565b6060608083600081526000602082015282604082015282808201520152565b90600182811c92168015612c9a575b6020831014612c6b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612c60565b906040519182815491828252602082019060005260206000209260005b818110612cd657505061044a925003836103fa565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612cc1565b90604051612d12816103a5565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c161515602083015260018101805490612d4c82611b9a565b91612d5a60405193846103fa565b80835260208301916000526020600020916000905b828210612da457505050506003608092612d9f926040860152612d9460028201612ca4565b606086015201612ca4565b910152565b60405160008554612db481612c51565b8084529060018116908115612e265750600114612dee575b5060019282612de0859460209403826103fa565b815201940191019092612d6f565b6000878152602081209092505b818310612e1057505081016020016001612dcc565b6001816020925483868801015201920191612dfb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050612dcc565b359061044a8261197d565b359061044a8261242a565b9080601f830112156100e7578135612e9481611b9a565b92612ea260405194856103fa565b81845260208085019260051b820101918383116100e75760208201905b838210612ece57505050505090565b813567ffffffffffffffff81116100e757602091612ef187848094880101611d3d565b815201910190612ebf565b9080601f830112156100e7578135612f1381611b9a565b92612f2160405194856103fa565b81845260208085019260051b8201019283116100e757602001905b828210612f495750505090565b602080918335612f588161197d565b815201910190612f3c565b60c0813603126100e757612f7561043b565b90612f7f81612e67565b8252612f8d6020820161179d565b6020830152612f9e60408201612e72565b6040830152606081013567ffffffffffffffff81116100e757612fc49036908301612e7d565b6060830152608081013567ffffffffffffffff81116100e757612fea9036908301612efc565b608083015260a08101359067ffffffffffffffff82116100e75761301091369101612efc565b60a082015290565b604051602081019060008252602081526130336040826103fa565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110613073575050565b60008155600101613068565b9190601f811161308e57505050565b61044a926000526020600020906020601f840160051c830193106130ba575b601f0160051c0190613068565b90915081906130ad565b919091825167ffffffffffffffff81116103a0576130ec816130e68454612c51565b8461307f565b6020601f821160011461314a57819061313b93949560009261313f575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880613109565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061317d84600052602060002090565b9160005b8181106131d7575095836001959697106131a0575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613196565b9192602060018192868b015181550194019201613181565b8151916801000000000000000083116103a0578154838355808410613251575b506020613223910191600052602060002090565b6000915b8383106132345750505050565b6001602082613245839451866130c4565b01920192019190613227565b8260005283602060002091820191015b81811061326e575061320f565b8061327b60019254612c51565b80613288575b5001613261565b601f8111831461329e5750600081555b38613281565b6132c29083601f6132b485600052602060002090565b920160051c82019101613068565b60008181526020812081835555613298565b81519167ffffffffffffffff83116103a0576801000000000000000083116103a057602090825484845580851061334b575b500190600052602060002060005b8381106133215750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501613314565b613362908460005285846000209182019101613068565b38613306565b906108fb916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a06133ff6133cc606085015160c0608086015260e0850190610819565b60808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526105da565b9201519060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526105da565b908160209103126100e757516108fb8161242a565b60409073ffffffffffffffffffffffffffffffffffffffff6108fb949316815281602082015201906104ec565b9060048110156107b35760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b8383106134d757505050505090565b909192939460208061357c837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a061356b6135596135476135358887015160c08a88015260c08701906104ec565b604087015186820360408801526104ec565b606086015185820360608701526104ec565b608085015184820360808601526104ec565b9201519060a08184039101526104ec565b970193019301919392906134c8565b9160209082815201919060005b8181106135a55750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356135ce8161197d565b168152019401929101613598565b90602083828152019260208260051b82010193836000925b8484106136045750505050505090565b90919293949560208061364a837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030188526136448b88611f24565b90611c75565b98019401940192949391906135f4565b9492936108fb969461385a61386d94936080895261368560808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a0152606081015163ffffffff1660e08a0152608081015163ffffffff166101008a015260a081015161ffff166101208a015260c08101516101408a01526101a06138276137f16137bb6137838d61374d61371760e08901516101c06101608501526102408401906104ec565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101808501526104ec565b9061012088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526104ec565b8d610140870151906101c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526104ec565b6101608501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101e08f01526104ec565b6101808401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016102008e01526134ab565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016102208b01526104ec565b926020880152868303604088015261358b565b9260608185039101526135dc565b8061388c6040926108fb95946107b8565b81602082015201906104ec565b3d156138c4573d906138aa8261047a565b916138b860405193846103fa565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161397e575b5061397a5782613944613899565b906107766040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401613447565b9150565b90916020823d6020116139ad575b81613999602093836103fa565b810103126139aa5750519038613936565b80fd5b3d915061398c565b604051906139c46020836103fa565b6000808352366020840137565b9190811015611c705760051b0190565b356108fb8161197d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114613a185760010190565b613039565b8015613a18577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd1208201918211613a1857565b91908203918211613a1857565b90613a8e939291614dc2565b91939092613a9b82611bb2565b92613aa583611bb2565b94600091825b8851811015613bc1576000805b8a888489828510613b28575b505050505015613ad657600101613aab565b613ae6610c1561140a928b611cd2565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610c15613b5a926128b2613b558873ffffffffffffffffffffffffffffffffffffffff97610bcc966139d1565b6139e1565b911614613b6957600101613ab8565b60019150613ba3613b7e613b55838b8b6139d1565b613b88888c611cd2565b9073ffffffffffffffffffffffffffffffffffffffff169052565b613bb6613baf876139eb565b968b611cd2565b52388a888489613ac4565b509097965094939291909460ff811690816000985b8a518a1015613c935760005b8b87821080613c8a575b15613c7d5773ffffffffffffffffffffffffffffffffffffffff613c1e610bcc610c158f6128b2613b55888f8f6139d1565b911614613c3357613c2e906139eb565b613be2565b9399613c4360019294939b613a1d565b94613c5f613c55613b55838b8b6139d1565b613b888d8c611cd2565b613c72613c6b8c6139eb565b9b8b611cd2565b525b01989091613bd6565b5050919098600190613c74565b50851515613bec565b985092509395949750915081613cbc5750505081518103613cb357509190565b80825283529190565b61140a9291613cca91613a75565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906104c3826103c1565b908160209103126100e75760405190613d23826103c1565b51815290565b6108fb9160e0613dd3613dc1613d4a855161010086526101008601906104ec565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152613daf6080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526104ec565b60c085015184820360c08601526104ec565b9201519060e08184039101526104ec565b9060206108fb928181520190613d29565b9061ffff610650602092959495604085526040850190613d29565b919492939094613e1e612389565b50613e2f6106ef6080850151612bed565b90613e406106ef6060860151612bed565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201529096906020818060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611621576000916141e7575b5073ffffffffffffffffffffffffffffffffffffffff81169283156141a357613f0089866138c9565b96613f09613cfe565b5080516020820151613f1a90612bed565b6040805160609290921c6020808401919091528252909290613f3c90846103fa565b60a0015192613f4961045b565b9c8d5267ffffffffffffffff1660208d015273ffffffffffffffffffffffffffffffffffffffff871660408d015260608c015273ffffffffffffffffffffffffffffffffffffffff8a1660808c015260a08b015260c08a0152613faa6104b4565b60e08a0152613fb88161425d565b156140ce5750602090613ff998604051998a9283927f489a68f200000000000000000000000000000000000000000000000000000000845260048401613df5565b03816000855af16000978161409d575b506140175785613944613899565b859496905b73ffffffffffffffffffffffffffffffffffffffff83161461408c575061405593614046916138c9565b90828214614084575b50613a75565b61407c61406061046b565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b91503861404f565b93505050505161407c61406061046b565b6140c091985060203d6020116140c7575b6140b881836103fa565b810190613d0b565b9638614009565b503d6140ae565b90506140d9816142b4565b156141605750602061411897604051809981927f3907753700000000000000000000000000000000000000000000000000000000835260048301613de4565b03816000855af16000978161413f575b506141365785613944613899565b8594969061401c565b61415991985060203d6020116140c7576140b881836103fa565b9638614128565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff821660045260246000fd5b614200915060203d60201161291b5761290d81836103fa565b38613ed7565b61420f81615407565b908161424b575b8161421f575090565b6108fb91507f85572ffb00000000000000000000000000000000000000000000000000000000906154fc565b905061425681615498565b1590614216565b61426681615407565b90816142a2575b81614276575090565b6108fb91507f3317103100000000000000000000000000000000000000000000000000000000906154fc565b90506142ad81615498565b159061426d565b6142bd81615407565b90816142f9575b816142cd575090565b6108fb91507faff2afbf00000000000000000000000000000000000000000000000000000000906154fc565b905061430481615498565b15906142c4565b61431481615407565b9081614350575b81614324575090565b6108fb91507f7909b54900000000000000000000000000000000000000000000000000000000906154fc565b905061435b81615498565b159061431b565b6040519061436f826103dd565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b9015611c705790565b9060431015611c705760430190565b90821015611c70570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906023116100e75760210190600290565b906043116100e75760230190602090565b90929192836044116100e75783116100e757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff000000000000000000000000000000000000000000000000811692600881106144f1575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110614557575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff000000000000000000000000000000000000000000000000000000000000811692600281106145bd575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b3590602081106145fd575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff8211176103a057604052606060a0836000815282602082015282604082015282808201528260808201520152565b6040805190919061467d83826103fa565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b8281106146b457505050565b6020906146bf61462a565b828285010152016146a8565b604051906146da6020836103fa565b600080835282815b8281106146ee57505050565b6020906146f961462a565b828285010152016146e2565b9061470e614362565b91604d8210614d485761475361474d61472784846143cf565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603614d18575061478c61477e61477861477285856143f3565b906144bd565b60c01c90565b67ffffffffffffffff168452565b6147b061479f6147786147728585614404565b67ffffffffffffffff166020850152565b6147d46147c36147786147728585614415565b67ffffffffffffffff166040850152565b6148006147f36147ed6147e78585614426565b90614523565b60e01c90565b63ffffffff166060850152565b6148206148136147ed6147e78585614437565b63ffffffff166080850152565b61484a61483f6148396148338585614448565b90614589565b60f01c90565b61ffff1660a0850152565b61485d6148578383614459565b906145ef565b60c08401528160431015614ce95761488461487e61474d61472785856143d8565b60ff1690565b9081604401838111614ce95761489e612a9182868561446a565b60e086015283811015614cba5761487e61474d6147276148bf9387866143e7565b8201916045830190848211614c8b57612a918260456148e0930187866144a5565b61010086015283811015614c5c5761490461487e61474d61472760459488876143e7565b830101916001830190848211614c2d57612a91826046614926930187866144a5565b61012086015283811015614bfe5761494a61487e61474d61472760019488876143e7565b830101916001830190848211614bcf57612a9182600261496c930187866144a5565b6101408601526003830192848411614ba05761499c614995614839614833876001968a896144a5565b61ffff1690565b0101916002830190848211614b7157612a91826149ba9287866144a5565b6101608601526004830190848211614b4257614839614833836149de9388876144a5565b9261ffff8294168015600014614ad8575050506149f96146cb565b6101808501525b6002820191838311614aa95780614a24614995614839614833876002968a896144a5565b010191838311614a7a5782612a919185614a3d946144a5565b6101a084015203614a4b5790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b6002919294508190614afb614aeb61466c565b966101808a019788528887615562565b9490965196614b0a8698611cc5565b5201010114614a00577fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b420000000000000000000000000000000000000000000000000000000060005261140a6024906000600452565b9060018201809211613a1857565b91908201809211613a1857565b60ff168015613a18577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9391614dcc6139b5565b93815180615109575b505050614de29084615c26565b92909293614e326002614e11614e176003614e118b67ffffffffffffffff166000526004602052604060002090565b01612ca4565b9867ffffffffffffffff166000526004602052604060002090565b614e5d614e58614e50614e488751865190614d87565b8a5190614d87565b835190614d87565b611bb2565b9687956000805b8751821015614e9c5790614e94600192613b88614e84610c15858d611cd2565b91614e8e816139eb565b9c611cd2565b018997614e64565b9750509193969092945060005b8551811015614ede5780614ed8614ec5610c156001948a611cd2565b613b88614ed18b6139eb565b9a8d611cd2565b01614ea9565b50919350919460005b8451811015614f1c5780614f16614f03610c1560019489611cd2565b613b88614f0f8a6139eb565b998c611cd2565b01614ee7565b5093919592509360005b8281106150a0575b5050600090815b8181106150015750508152835160005b818110614f5457508452929190565b926000959495915b8351831015614ff257614f72610c158688611cd2565b73ffffffffffffffffffffffffffffffffffffffff614f97610bcc610c158789611cd2565b911603614fe157614fa790613a1d565b90614fc2614fb8610c158489611cd2565b613b888789611cd2565b60ff8716614fd1575b90614f5c565b95614fdb90614d94565b95614fcb565b9091614fec906139eb565b91614fcb565b91509260019095949501614f45565b61500e610c158286611cd2565b73ffffffffffffffffffffffffffffffffffffffff8116801561509657600090815b86811061506a575b505090600192911561504d575b505b01614f35565b61506490613b8861505d876139eb565b9688611cd2565b38615045565b8161507b610bcc610c15848c611cd2565b1461508857600101615030565b506001915081905038615038565b5050600190615047565b6150b0610bcc610c158387611cd2565b156150bd57600101614f26565b5091929360009591955b83518110156150fc57806150f66150e3610c1560019488611cd2565b613b886150ef8b6139eb565b9a89611cd2565b016150c7565b5093929150933880614f2e565b909192945060018103615162575081615159916151386106ef6060615130614de297611cc5565b510151612bed565b918760a061515061514884611cc5565b515193611cc5565b51015193615a24565b92903880614dd5565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6001541633036151b057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916151e8815184614d87565b9283156153155760005b848110615200575050505050565b818110156152fa57615215610c158286611cd2565b73ffffffffffffffffffffffffffffffffffffffff81168015610c3c5761523b83614d79565b87811061524d575050506001016151f2565b848110156152ca5773ffffffffffffffffffffffffffffffffffffffff615277610c15838a611cd2565b1682146152865760010161523b565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6152f5610c156152ef8885613a75565b89611cd2565b615277565b615310610c1561530a8484613a75565b85611cd2565b615215565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b615347614362565b5090565b906040519161535b60c0846103fa565b60848352602083019060a03683375a612ee08111156153ac576000916153818392613a48565b82602083519301913090f1903d90608482116153a3575b6000908286523e9190565b60849150615398565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a70000000000000000000000000000000000000000000000000000000060248201526024815261546b6044826103fa565b5191617530fa6000513d8261548c575b5081615485575090565b9050151590565b6020111591503861547b565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff0000000000000000000000000000000000000000000000000000000060248201526024815261546b6044826103fa565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a70000000000000000000000000000000000000000000000000000000084521660248201526024815261546b6044826103fa565b929161556c61462a565b918082101561590f5761558661474d6147278484896143e7565b600160ff821603614d185750602182018181116158e0576155af6148578260018601858a6144a5565b8452818110156158b15761487e61474d6147276155cd93858a6143e7565b820191602283019082821161588257612a918260226155ee9301858a6144a5565b6020850152818110156158535761561161487e61474d614727602294868b6143e7565b83010191600183019082821161582457612a918260236156339301858a6144a5565b6040850152818110156157f55761565661487e61474d614727600194868b6143e7565b8301019160018301908282116157c657612a918260026156789301858a6144a5565b6060850152818110156157975761569b61487e61474d614727600194868b6143e7565b830101600181019282841161576857612a918460026156bc9301858a6144a5565b60808501526003810192828411615739576002916156e761499561483961483388600196898e6144a5565b0101019481861161570a5761570192612a919287926144a5565b60a08201529190565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9080601f830112156100e757815161595581611b9a565b9261596360405194856103fa565b81845260208085019260051b8201019283116100e757602001905b82821061598b5750505090565b60208091835161599a8161197d565b81520191019061597e565b906020828203126100e757815167ffffffffffffffff81116100e7576108fb920161593e565b95949060019460a09467ffffffffffffffff615a1f9573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906104ec565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095929493929091906020838060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa92831561162157600093615ba9575b50615acc8361425d565b615b06575b5050505050815115615ae1575090565b6108fb9150614e1160029167ffffffffffffffff166000526004602052604060002090565b60009496508591615b5c73ffffffffffffffffffffffffffffffffffffffff92604051998a97889687957f89720a62000000000000000000000000000000000000000000000000000000008752600487016159cb565b0392165afa91821561162157600092615b84575b50615b7a82615ebb565b3880808080615ad1565b615ba29192503d806000833e615b9a81836103fa565b8101906159a5565b9038615b70565b615bc391935060203d60201161291b5761290d81836103fa565b9138615ac2565b90916060828403126100e757815167ffffffffffffffff81116100e75783615bf391840161593e565b92602083015167ffffffffffffffff81116100e757604091615c1691850161593e565b92015160ff811681036100e75790565b90803b615c69575b50615c4f60029167ffffffffffffffff166000526004602052604060002090565b0190615c62615c5c6139b5565b92612ca4565b9190600090565b615c728161430b565b15615c2e576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff83166004820152906000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa801561162157600080928192615d1d575b50615ceb81615ebb565b615cf483615ebb565b805115801590615d11575b615d0a575050615c2e565b9391925090565b5060ff82161515615cff565b9150615d3b92503d8091833e615d3381836103fa565b810190615bca565b909138615ce1565b8054821015611c705760005260206000200190600090565b600254811015611c705760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b818110615da457505060009055565b80615db160019285615d43565b90549060031b1c6000528184016020526000604081205501615d95565b600081815260036020526040902054615e5957600254680100000000000000008110156103a057615e40615e0b8260018594016002556002615d43565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615eb457805490680100000000000000008210156103a05782615e9d615e0b846001809601855584615d43565b905580549260005201602052604060002055600190565b5050600090565b80519060005b828110615ecd57505050565b60018101808211613a18575b838110615ee95750600101615ec1565b73ffffffffffffffffffffffffffffffffffffffff615f088385611cd2565b5116615f1a610bcc610c158487611cd2565b14615f2757600101615ed9565b61140a615f37610c158486611cd2565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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
