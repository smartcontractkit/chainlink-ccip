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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getmaxGasBufferToUpdateState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMaxGasBufferToUpdateState\",\"inputs\":[{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxGasBufferToUpdateStateUpdated\",\"inputs\":[{\"name\":\"oldMaxGasBufferToUpdateState\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newMaxGasBufferToUpdateState\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"GasCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidGasLimitOverride\",\"inputs\":[{\"name\":\"messageGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100604052346102d457604051615f0738819003601f8101601f191683016001600160401b038111848210176102d95783928291604052833981010360a081126102d4576080136102d45760405190608082016001600160401b038111838210176102d95760405280516001600160401b03811681036102d457825260208101519161ffff831683036102d457602081019283526040820151916001600160a01b03831683036102d457604082019283526060810151906001600160a01b03821682036102d4576080906060840192835201519263ffffffff8416948585036102d45733156102c35760015482519091906001600160a01b03161580156102b1575b6102a05784516001600160401b03161561028f5761ffff8151161561027e5784516001600160401b031660805282516001600160a01b0390811660a05284511660c052805161ffff1660e052861561027e577f7266121371af537e246f0b727f08d4a221cdcdb38ff862bb874bc4b55a6642dd604061ffff937f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e49509960809963ffffffff60a81b9060a81b1663ffffffff60a81b1933166bffffffffffffff00000000ff60a01b8416171760015563ffffffff835192339060018060a01b0319161760a81c1682526020820152a16040805195516001600160401b03168652905191909116602085015290516001600160a01b03908116918401919091529051166060820152a1604051615c1790816102f082396080518181816101580152611a21015260a0518181816101bb0152611948015260c0518181816101e301528181614b5801526155b1015260e05181818161017f015261508c0152f35b632855a4d960e11b60005260046000fd5b63c656089560e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b5083516001600160a01b031615610102565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100f7578063181f5a77146100f257806320f81c88146100ed5780633b81ab9b146100e857806349d8033e146100e35780635215505b146100de5780635643a782146100d957806361a10e59146100d457806379ba5097146100cf5780638da5cb5b146100ca5780638f4d0559146100c5578063e9d68a8e146100c0578063ea3fbef9146100bb5763f2fde38b146100b657600080fd5b6111c8565b6110b1565b610f0b565b610eb6565b610e82565b610db7565b610d24565b610967565b610850565b6106b1565b6105d8565b6104ed565b6103e3565b61010c565b600091031261010757565b600080fd5b34610107576000600319360112610107576000606060405161012d81610292565b828152826020820152826040820152015261025f60405161014d81610292565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ae57604052565b610263565b60a0810190811067ffffffffffffffff8211176102ae57604052565b6101c0810190811067ffffffffffffffff8211176102ae57604052565b6020810190811067ffffffffffffffff8211176102ae57604052565b90601f601f19910116810190811067ffffffffffffffff8211176102ae57604052565b6040519061033a60c083610308565b565b6040519061033a60a083610308565b6040519061033a61010083610308565b6040519061033a604083610308565b67ffffffffffffffff81116102ae57601f01601f191660200190565b60405190610395602083610308565b60008252565b60005b8381106103ae5750506000910152565b818101518382015260200161039e565b90601f19601f6020936103dc8151809281875287808801910161039b565b0116010190565b346101075760006003193601126101075761025f60408051906104068183610308565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906103be565b9181601f840112156101075782359167ffffffffffffffff8311610107576020838186019501011161010757565b906020808351928381520192019060005b81811061048e5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610481565b916104e660ff916104d8604094979697606087526060870190610470565b908582036020870152610470565b9416910152565b346101075760206003193601126101075760043567ffffffffffffffff81116101075761057f61052d61052761025f933690600401610442565b90613896565b61053b61014082015161129e565b60601c9067ffffffffffffffff81511691610180820151906105798161ffff60a0860151169463ffffffff60806101a0830151519201511690613da8565b93613f26565b604093919351938493846104ba565b9181601f840112156101075782359167ffffffffffffffff8311610107576020808501948460051b01011161010757565b63ffffffff81160361010757565b359061033a826105bf565b346101075760806003193601126101075760043567ffffffffffffffff811161010757610609903690600401610442565b9060243567ffffffffffffffff81116101075761062a90369060040161058e565b926044359367ffffffffffffffff85116101075761064f61066495369060040161058e565b9390926064359561065f876105bf565b61181c565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561069f57565b610666565b90600482101561069f5752565b34610107576020600319360112610107576004356000526006602052602060ff604060002054166106e560405180926106a4565bf35b9080602083519182815201916020808360051b8301019401926000915b83831061071357505050505090565b909192939460208061073183601f19866001960301875289516103be565b97019301930191939290610704565b6107ab9173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061079a610788604085015160a0604086015260a08501906106e7565b60608501518482036060860152610470565b920151906080818403910152610470565b90565b6040810160408252825180915260206060830193019060005b818110610830575050506020818303910152815180825260208201916020808360051b8301019401926000915b83831061080357505050505090565b909192939460208061082183601f1986600196030187528951610740565b970193019301919392906107f4565b825167ffffffffffffffff168552602094850194909201916001016107c7565b346101075760006003193601126101075760025461086d81611f84565b9061087b6040519283610308565b808252601f1961088a82611f84565b0160005b8181106109505750506108a081611fc8565b9060005b8181106108bc57505061025f604051928392836107ae565b806108f46108db6108ce600194615864565b67ffffffffffffffff1690565b6108e5838761203a565b9067ffffffffffffffff169052565b61093461092f610915610907848861203a565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b612102565b61093e828761203a565b52610949818661203a565b50016108a4565b60209061095b611f9c565b8282870101520161088e565b346101075760206003193601126101075760043567ffffffffffffffff81116101075761099890369060040161058e565b906109a161446e565b6000905b8282106109ae57005b6109c16109bc8385846122a7565b6123d5565b60208101916109db6108ce845167ffffffffffffffff1690565b15610cfa57610a1d610a04610a04845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610ced575b610a745760808201949060005b86518051821015610a9e57610a04610a4d83610a679361203a565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610a7457600101610a32565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610ad657610a04610a4d83610ac99361203a565b15610a7457600101610aae565b505095929491909394610aec86518251906144b9565b610b01610915835167ffffffffffffffff1690565b90610b31610b17845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610b3b87615899565b606085019860005b8a518051821015610b935790610b5b8160209361203a565b518051928391012091158015610b83575b610a7457610b7c6001928b615968565b5001610b43565b50610b8c61248a565b8214610b6c565b5050976001975093610caa610ce3946003610cd695610ca27f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e996108ce979f610bf060408e610be9610c36945160018b01612643565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610c98610c578d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5160028501612728565b519101612728565b610cc7610cc26108ce835167ffffffffffffffff1690565b6158d7565b505167ffffffffffffffff1690565b92604051918291826127bc565b0390a201906109a5565b5060808201515115610a25565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346101075760a06003193601126101075760043567ffffffffffffffff8111610107576101c06003198236030112610107576024359060443567ffffffffffffffff811161010757610d7a90369060040161058e565b926064359367ffffffffffffffff851161010757610d9f61066495369060040161058e565b93909260843595610daf876105bf565b600401612f87565b346101075760006003193601126101075760005473ffffffffffffffffffffffffffffffffffffffff81163303610e58577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461010757600060031936011261010757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461010757600060031936011261010757602063ffffffff60015460a81c16604051908152f35b67ffffffffffffffff81160361010757565b359061033a82610edd565b9060206107ab928181520190610740565b346101075760206003193601126101075767ffffffffffffffff600435610f3181610edd565b610f39611f9c565b501660005260046020526040600020604051610f54816102b3565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015260018201805490610f8c82611f84565b91610f9a6040519384610308565b80835260208301916000526020600020916000905b828210610fee5761025f86610fdd60038a896040850152610fd2600282016120a1565b6060850152016120a1565b608082015260405191829182610efa565b60405160008554610ffe8161204e565b80845290600181169081156110705750600114611038575b506001928261102a85946020940382610308565b815201940191019092610faf565b6000878152602081209092505b81831061105a57505081016020016001611016565b6001816020925483868801015201920191611045565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050611016565b34610107576020600319360112610107576004356110ce816105bf565b6110d661446e565b63ffffffff8116156111805763ffffffff7f7266121371af537e246f0b727f08d4a221cdcdb38ff862bb874bc4b55a6642dd9161117b6001549178ffffffff0000000000000000000000000000000000000000008160a81b167fffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffffff84161760015560405193849360a81c168390929163ffffffff60209181604085019616845216910152565b0390a1005b7f50ab49b20000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361010757565b346101075760206003193601126101075773ffffffffffffffffffffffffffffffffffffffff6004356111fa816111aa565b61120261446e565b1633811461127457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106112d6575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b8015150361010757565b9081602091031261010757516107ab81611308565b6040513d6000823e3d90fd5b9060206107ab9281815201906103be565b60409073ffffffffffffffffffffffffffffffffffffffff6107ab949316815281602082015201906103be565b92919261137d8261036a565b9161138b6040519384610308565b829481845281830111610107578281602093846000960137010152565b90600481101561069f5760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061140b57505050505090565b909192939460208061149283601f1986600196030187528951908151815260a061148161146f61145d61144b8887015160c08a88015260c08701906103be565b604087015186820360408801526103be565b606086015185820360608701526103be565b608085015184820360808601526103be565b9201519060a08184039101526103be565b970193019301919392906113fc565b9160209082815201919060005b8181106114bb5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356114e4816111aa565b1681520194019291016114ae565b601f8260209493601f19938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561010757016020813591019167ffffffffffffffff821161010757813603831361010757565b90602083828152019260208260051b82010193836000925b84841061158b5750505050505090565b9091929394956020806115b383601f1986600196030188526115ad8b88611513565b906114f2565b980194019401929493919061157b565b979694916117e3906080956117f1956117d061033a9a956116558e60a081526115f960a08201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660c0820152604083015167ffffffffffffffff1660e0820152606083015163ffffffff16610100820152828c015163ffffffff1661012082015261014060a084015191019061ffff169052565b8d61016060c08301519101528d6101a061179c6117666117306116fa6116c461169060e08901516101c06101808a01526102608901906103be565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6089830301888a01526103be565b6101208801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60888303016101c08901526103be565b6101408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016101e08801526103be565b6101608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303016102008701526103be565b6101808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60858303016102208601526113df565b920151906102407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60828503019101526103be565b9260208d01528b830360408d01526114a1565b9188830360608a0152611563565b94019063ffffffff169052565b8061180f6040926107ab95946106a4565b81602082015201906103be565b959192939594909461183460015460ff9060a01c1690565b611f5a5761187c740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b61188e6118898783613896565b61436c565b9561192f60206118d46118ac6108ce8b5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611f5557600091611f26575b50611edb576119a76119a36119996109158a5167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b611e90576119c0610b17885167ffffffffffffffff1690565b6119ea6119a360e08a01928351602081519101209060019160005201602052604060002054151590565b611e5957506101008701516014815114801590611e28575b611df15750602087015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603611dba5750828603611d86576101408701516014815103611d4c575063ffffffff84168015159081611d26575b50611cd95790611a8b913691611371565b6020815191012095611ab1611aaa886000526006602052604060002090565b5460ff1690565b611aba81610695565b8015908115611cc5575b5015611c5357611b5f92611b64959492611b5192611b1a611aef8b6000526006602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957f61a10e590000000000000000000000000000000000000000000000000000000060208801528b8b602489016115c3565b03601f198101835282610308565b614378565b9015611c21577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b611bb584611bb0886000526006602052604060002090565b6113a8565b611bf1611bdf6040611bcf885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b918360405194859416971695836117fe565b0390a461033a7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff600392611b98565b611cc18787611c7f6040611c6f835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150611cd281610695565b1438611ac4565b611cc184611cee60808a015163ffffffff1690565b7fdf2964df0000000000000000000000000000000000000000000000000000000060005263ffffffff90811660045216602452604490565b9050611d45611d3c60808a015163ffffffff1690565b63ffffffff1690565b1138611a7a565b611d82906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611333565b0390fd5b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004869052602483905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b611d82906040519182917f55216e310000000000000000000000000000000000000000000000000000000083523060048401611344565b50611e3b611e358261129e565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff16301415611a02565b611d8290516040519182917fa50bd14700000000000000000000000000000000000000000000000000000000835260048301611333565b611cc1611ea5885167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611cc1611ef0885167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611f48915060203d602011611f4e575b611f408183610308565b810190611312565b38611979565b503d611f36565b611327565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116102ae5760051b60200190565b60405190611fa9826102b3565b6060608083600081526000602082015282604082015282808201520152565b90611fd282611f84565b611fdf6040519182610308565b828152601f19611fef8294611f84565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156120355760200190565b611ff9565b80518210156120355760209160051b010190565b90600182811c92168015612097575b602083101461206857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161205d565b906040519182815491828252602082019060005260206000209260005b8181106120d357505061033a92500383610308565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016120be565b9060405161210f816102b3565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c16151560208301526001810180549061214982611f84565b916121576040519384610308565b80835260208301916000526020600020916000905b8282106121a15750505050600360809261219c926040860152612191600282016120a1565b6060860152016120a1565b910152565b604051600085546121b18161204e565b808452906001811690811561222357506001146121eb575b50600192826121dd85946020940382610308565b81520194019101909261216c565b6000878152602081209092505b81831061220d575050810160200160016121c9565b60018160209254838688010152019201916121f8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506121c9565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610107570190565b901561203557806107ab91612264565b90821015612035576107ab9160051b810190612264565b359061033a826111aa565b359061033a82611308565b9080601f83011215610107578160206107ab93359101611371565b9080601f8301121561010757813561230681611f84565b926123146040519485610308565b81845260208085019260051b820101918383116101075760208201905b83821061234057505050505090565b813567ffffffffffffffff811161010757602091612363878480948801016122d4565b815201910190612331565b9080601f8301121561010757813561238581611f84565b926123936040519485610308565b81845260208085019260051b82010192831161010757602001905b8282106123bb5750505090565b6020809183356123ca816111aa565b8152019101906123ae565b60c081360312610107576123e761032b565b906123f1816122be565b82526123ff60208201610eef565b6020830152612410604082016122c9565b6040830152606081013567ffffffffffffffff81116101075761243690369083016122ef565b6060830152608081013567ffffffffffffffff81116101075761245c903690830161236e565b608083015260a08101359067ffffffffffffffff8211610107576124829136910161236e565b60a082015290565b604051602081019060008252602081526124a5604082610308565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106124e5575050565b600081556001016124da565b9190601f811161250057505050565b61033a926000526020600020906020601f840160051c8301931061252c575b601f0160051c01906124da565b909150819061251f565b919091825167ffffffffffffffff81116102ae5761255e81612558845461204e565b846124f1565b6020601f82116001146125bc5781906125ad9394956000926125b1575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b01519050388061257b565b601f198216906125d184600052602060002090565b9160005b81811061262b575095836001959697106125f4575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806125ea565b9192602060018192868b0151815501940192016125d5565b8151916801000000000000000083116102ae5781548383558084106126a5575b506020612677910191600052602060002090565b6000915b8383106126885750505050565b600160208261269983945186612536565b0192019201919061267b565b8260005283602060002091820191015b8181106126c25750612663565b806126cf6001925461204e565b806126dc575b50016126b5565b601f811183146126f25750600081555b386126d5565b6127169083601f61270885600052602060002090565b920160051c820191016124da565b600081815260208120818355556126ec565b81519167ffffffffffffffff83116102ae576801000000000000000083116102ae57602090825484845580851061279f575b500190600052602060002060005b8381106127755750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501612768565b6127b69084600052858460002091820191016124da565b3861275a565b906107ab916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a0612835612820606085015160c0608086015260e08501906106e7565b6080850151601f198583030184860152610470565b9201519060c0601f1982850301910152610470565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610107570180359067ffffffffffffffff821161010757602001918160051b3603831361010757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610107570180359067ffffffffffffffff82116101075760200191813603831361010757565b9160206107ab9381815201916114f2565b919091357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106112d6575050565b356107ab816105bf565b356107ab81610edd565b61ffff81160361010757565b356107ab81612948565b91909160c0818403126101075761297361032b565b9281358452602082013567ffffffffffffffff811161010757816129989184016122d4565b6020850152604082013567ffffffffffffffff811161010757816129bd9184016122d4565b6040850152606082013567ffffffffffffffff811161010757816129e29184016122d4565b6060850152608082013567ffffffffffffffff81116101075781612a079184016122d4565b608085015260a082013567ffffffffffffffff811161010757612a2a92016122d4565b60a0830152565b929190612a3d81611f84565b93612a4b6040519586610308565b602085838152019160051b8101918383116101075781905b838210612a71575050505050565b813567ffffffffffffffff811161010757602091612a92878493870161295e565b815201910190612a63565b9082101561203557612ab49160051b81019061289e565b9091565b9081602091031261010757516107ab816111aa565b60409073ffffffffffffffffffffffffffffffffffffffff6107ab959316815281602082015201916114f2565b359061033a82612948565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561010757016020813591019167ffffffffffffffff8211610107578160051b3603831361010757565b90602083828152019060208160051b85010193836000915b838310612b805750505050505090565b909192939495601f1982820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4184360301811215610107576020612c63600193868394019081358152612c55612c4a612c2f612c14612bf9612be989880188611513565b60c08b89015260c08801916114f2565b612c066040880188611513565b9087830360408901526114f2565b612c216060870187611513565b9086830360608801526114f2565b612c3c6080860186611513565b9085830360808701526114f2565b9260a0810190611513565b9160a08185039101526114f2565b980196019493019190612b70565b612ee96107ab9593949260608352612c9d60608401612c8f83610eef565b67ffffffffffffffff169052565b612cbd612cac60208301610eef565b67ffffffffffffffff166080850152565b612cdd612ccc60408301610eef565b67ffffffffffffffff1660a0850152565b612cf9612cec606083016105cd565b63ffffffff1660c0850152565b612d15612d08608083016105cd565b63ffffffff1660e0850152565b612d30612d2460a08301612afa565b61ffff16610100850152565b60c0810135610120840152612eb8612eac612e6d612e2e612def612db0612d71612d5d60e0890189611513565b6101c06101408d01526102208c01916114f2565b612d7f610100890189611513565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d01526114f2565b612dbe610120880188611513565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c01526114f2565b612dfd610140870187611513565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b01526114f2565b612e3c610160860186611513565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a01526114f2565b612e7b610180850185612b05565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152612b58565b916101a0810190611513565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0858403016102008601526114f2565b93602082015260408185039101526114f2565b604051906040820182811067ffffffffffffffff8211176102ae5760405260006020838281520152565b90612f3082611f84565b612f3d6040519182610308565b828152601f19612f4d8294611f84565b019060005b828110612f5e57505050565b602090612f69612efc565b82828501015201612f52565b91908203918211612f8257565b6124ab565b9391949092943033036134e7576000612fa461018087018761284a565b9050613408575b61302d612fc8611e35612fc26101408a018a61289e565b90612900565b92612ff284612fdb6101a08b018b61289e565b9050612fec611d3c60808d01612934565b90613da8565b98899160a06130008b61293e565b876130278d61301f61301661018083018361284a565b96909201612954565b943691612a31565b91614754565b92909860005b8a5181101561321d57806020613056610a04610a048f6130a296610a4d9161203a565b61306b613063848a61203a565b518a8c612a9d565b91906040518096819482937fc3a7ded6000000000000000000000000000000000000000000000000000000008452600484016128ef565b03915afa918215611f55576000926131ed575b5073ffffffffffffffffffffffffffffffffffffffff821615613190576130e76130df828861203a565b51888a612a9d565b929073ffffffffffffffffffffffffffffffffffffffff82163b15610107576000918b8373ffffffffffffffffffffffffffffffffffffffff8f61315a604051998a97889687947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601612c71565b0393165af1918215611f5557600192613175575b5001613033565b80613184600061318a93610308565b806100fc565b3861316e565b856131b889898f856131ab610a4d611d82986131b19461203a565b9561203a565b5191612a9d565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501612acd565b61320f91925060203d8111613216575b6132078183610308565b810190612ab8565b90386130b5565b503d6131fd565b509597935095935096505061324061323961018084018461284a565b9050612f26565b9561324f61018084018461284a565b90506132f6575b506132ef5761033a946132bf61326b8361293e565b6132ac6132b361327f61012087018761289e565b61328d6101a089018961289e565b94909561329861033c565b9c8d5267ffffffffffffffff1660208d0152565b3691611371565b60408901523691611371565b6060860152608085015263ffffffff8216156132dc575091614ff4565b6132e99150608001612934565b91614ff4565b5050505050565b61335061331061330a61018086018661284a565b90612297565b61331e61012086018661289e565b61334a61332a8861293e565b9261334261333a60a08b01612954565b95369061295e565b923691611371565b90614ac3565b919061335b89612028565b5273ffffffffffffffffffffffffffffffffffffffff613395611e35612fc261338b61330a6101808a018a61284a565b608081019061289e565b921673ffffffffffffffffffffffffffffffffffffffff8316036133ba575b50613256565b6133ee6133f3926133e86133cd8b612028565b515173ffffffffffffffffffffffffffffffffffffffff1690565b9061464e565b612f75565b60206133fe88612028565b51015238806133b4565b50601461342961341f61330a61018089018961284a565b606081019061289e565b9050036134d357601461344661338b61330a61018089018961284a565b9050036134895761348461346a611e35612fc261338b61330a6101808b018b61284a565b6133e8611e35612fc261341f61330a6101808c018c61284a565b612fab565b61349d61338b61330a61018088018861284a565b90611d826040519283927f8d666f60000000000000000000000000000000000000000000000000000000008452600484016128ef565b61349d61341f61330a61018088018861284a565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061351e826102cf565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b90156120355790565b90604310156120355760430190565b90821015612035570190565b906009116101075760010190600890565b906011116101075760090190600890565b906019116101075760110190600890565b90601d116101075760190190600490565b9060211161010757601d0190600490565b906023116101075760210190600290565b906043116101075760230190602090565b909291928360441161010757831161010757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b90939293848311610107578411610107578101920390565b919091357fffffffffffffffff000000000000000000000000000000000000000000000000811692600881106136a0575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613706575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff0000000000000000000000000000000000000000000000000000000000008116926002811061376c575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b3590602081106137ac575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff8211176102ae57604052606060a0836000815282602082015282604082015282808201528260808201520152565b6040805190919061382c8382610308565b6001815291601f19018260005b82811061384557505050565b6020906138506137d9565b82828501015201613839565b6040519061386b602083610308565b600080835282815b82811061387f57505050565b60209061388a6137d9565b82828501015201613873565b9061389f613511565b91604d8210613d90576138e46138de6138b8848461357e565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613d60575061391d61390f61390961390385856135a2565b9061366c565b60c01c90565b67ffffffffffffffff168452565b61394161393061390961390385856135b3565b67ffffffffffffffff166020850152565b61396561395461390961390385856135c4565b67ffffffffffffffff166040850152565b61399161398461397e61397885856135d5565b906136d2565b60e01c90565b63ffffffff166060850152565b6139b16139a461397e61397885856135e6565b63ffffffff166080850152565b6139db6139d06139ca6139c485856135f7565b90613738565b60f01c90565b61ffff1660a0850152565b6139ee6139e88383613608565b9061379e565b60c08401528160431015613d4a57613a15613a0f6138de6138b88585613587565b60ff1690565b9081604401838111613d3457613a2f6132ac828685613619565b60e086015283811015613d1e57613a0f6138de6138b8613a50938786613596565b8201916045830190848211613d09576132ac826045613a7193018786613654565b61010086015283811015613cf357613a95613a0f6138de6138b86045948887613596565b830101916001830190848211613cdd576132ac826046613ab793018786613654565b61012086015283811015613cc757613adb613a0f6138de6138b86001948887613596565b830101916001830190848211613cb1576132ac826002613afd93018786613654565b6101408601526003830192848411613c9b57613b2d613b266139ca6139c4876001968a89613654565b61ffff1690565b0101916002830190848211613c85576132ac82613b4b928786613654565b6101608601526004830190848211613c6f576139ca6139c483613b6f938887613654565b9261ffff8294168015600014613c1e57505050613b8a61385c565b6101808501525b6002820191838311613c085780613bb5613b266139ca6139c4876002968a89613654565b010191838311613bf257826132ac9185613bce94613654565b6101a084015203613bdc5790565b635a102da160e11b600052601260045260246000fd5b635a102da160e11b600052601160045260246000fd5b635a102da160e11b600052601060045260246000fd5b6002919294508190613c41613c3161381b565b966101808a019788528887615131565b9490965196613c508698612028565b5201010114613b9157635a102da160e11b600052600f60045260246000fd5b635a102da160e11b600052600e60045260246000fd5b635a102da160e11b600052600d60045260246000fd5b635a102da160e11b600052600c60045260246000fd5b635a102da160e11b600052600b60045260246000fd5b635a102da160e11b600052600a60045260246000fd5b635a102da160e11b600052600960045260246000fd5b635a102da160e11b600052600860045260246000fd5b635a102da160e11b6000526004805260246000fd5b635a102da160e11b600052600360045260246000fd5b635a102da160e11b600052600260045260246000fd5b635a102da160e11b600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b635a102da160e11b600052611cc16024906000600452565b15919082613e26575b508115613e1c575b8115613dc3575090565b9050613def7f85572ffb0000000000000000000000000000000000000000000000000000000082615aea565b9081613e0a575b81613e0057501590565b6119a39150615a8a565b9050613e15816159c4565b1590613df6565b803b159150613db9565b15915038613db1565b60405190613e3e602083610308565b6000808352366020840137565b60408051909190613e5c8382610308565b6001815291601f1901366020840137565b9060018201809211612f8257565b91908201809211612f8257565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114612f825760010190565b80548210156120355760005260206000200190600090565b8015612f82577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff168015612f82577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b94929391606091600093613f38613e2f565b968351806142ed575b5050156142d9575051156142c8575050613f59613e2f565b613f61613e2f565b6000935b6002613fae613f936003613f8d8a67ffffffffffffffff166000526004602052604060002090565b016120a1565b9767ffffffffffffffff166000526004602052604060002090565b0195613fc8613fc08551845190613e7b565b825190613e7b565b90613fdd613fd889548094613e7b565b611fc8565b9788966000805b8851821015614037579061402f600192614014614004610a4d858e61203a565b9161400e81613e88565b9d61203a565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018a98613fe4565b91939597505097909294976000905b875182101561407e5790614076600192614014614066610a4d858d61203a565b9161407081613e88565b9c61203a565b018997614046565b9750509193969092945060005b85518110156140c057806140ba6140a7610a4d6001948a61203a565b6140146140b38b613e88565b9a8d61203a565b0161408b565b50929590935093909360005b828110614244575b50509091929350600090815b8181106141a25750508452805160005b855181101561419c5760005b82811061410d575b506001016140f0565b61411a610a4d828661203a565b73ffffffffffffffffffffffffffffffffffffffff61413f610a04610a4d868c61203a565b9116146141545761414f90613e88565b6140fc565b9161416161417991613ecd565b92614014614172610a4d868861203a565b918661203a565b60ff8416614188575b38614104565b92614194600191613ef8565b939050614182565b50815291565b6141af610a4d828961203a565b73ffffffffffffffffffffffffffffffffffffffff8116801561423a57600090815b8a87821061420d575b5050509060019291156141f0575b505b016140e0565b6142079061401461420087613e88565b968b61203a565b386141e8565b61421e610a04610a4d84869461203a565b1461422b576001016141d1565b5060019150819050388a6141da565b50506001906141ea565b614254610a04610a4d838b61203a565b15614261576001016140cc565b50909192939460005b82811061427c578695949392506140d4565b806142c26142af61428f60019486613eb5565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b6140146142bb88613e88565b978c61203a565b0161426a565b6142d3939193613e4b565b91613f65565b9150506142e7915084615728565b93613f65565b909197506001810361433f575061433790614316611e358661430e87612028565b51015161129e565b9061432085612028565b51518a60a061432e88612028565b51015193615547565b953880613f41565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b614374613511565b5090565b906040519161438860c084610308565b60848352602083019060a036833760015460a81c63ffffffff165a91818311156143e5576143ba600093928493612f75565b82602083519301913090f1903d90608482116143dc575b6000908286523e9190565b608491506143d1565b611cc16144206143f85a63ffffffff1690565b60e01b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b73ffffffffffffffffffffffffffffffffffffffff60015416330361448f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916144c7815184613e7b565b9283156145f45760005b8481106144df575050505050565b818110156145d9576144f4610a4d828661203a565b73ffffffffffffffffffffffffffffffffffffffff81168015610a745761451a83613e6d565b87811061452c575050506001016144d1565b848110156145a95773ffffffffffffffffffffffffffffffffffffffff614556610a4d838a61203a565b1682146145655760010161451a565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6145d4610a4d6145ce8885612f75565b8961203a565b614556565b6145ef610a4d6145e98484612f75565b8561203a565b6144f4565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b3d15614649573d9061462f8261036a565b9161463d6040519384610308565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181614703575b506146ff57826146c961461e565b90611d826040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401611344565b9150565b90916020823d602011614732575b8161471e60209383610308565b8101031261472f57505190386146bb565b80fd5b3d9150614711565b91908110156120355760051b0190565b356107ab816111aa565b90614763949596939291613f26565b9193909261477082611fc8565b9261477a83611fc8565b94600091825b8851811015614874576000805b8a8884898285106147fd575b5050505050156147ab57600101614780565b6147bb610a4d611cc1928b61203a565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610a4d61482f926131ab61482a8873ffffffffffffffffffffffffffffffffffffffff97610a049661473a565b61474a565b91161461483e5760010161478d565b6001915061485d61485361482a838b8b61473a565b614014888c61203a565b61486961420087613e88565b52388a888489614799565b509097965094939291909460ff811690816000985b8a518a10156149465760005b8b8782108061493d575b156149305773ffffffffffffffffffffffffffffffffffffffff6148d1610a04610a4d8f6131ab61482a888f8f61473a565b9116146148e6576148e190613e88565b614895565b93996148f660019294939b613ecd565b9461491261490861482a838b8b61473a565b6140148d8c61203a565b61492561491e8c613e88565b9b8b61203a565b525b01989091614889565b5050919098600190614927565b5085151561489f565b98509250939594975091508161496f575050508151810361496657509190565b80825283529190565b611cc1929161497d91612f75565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b60405190610395826102ec565b9081602091031261010757604051906149d6826102ec565b51815290565b6107ab9160e0614a86614a746149fd855161010086526101008601906103be565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152614a626080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526103be565b60c085015184820360c08601526103be565b9201519060e08184039101526103be565b9060206107ab9281815201906149dc565b9061ffff6104e66020929594956040855260408501906149dc565b92939193614acf612efc565b50614ae0611e35608086015161129e565b91614af1611e35606087015161129e565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909690956020878060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa968715611f5557600097614e20575b5073ffffffffffffffffffffffffffffffffffffffff8716948515614ddc57614bb06149b1565b50614bfe825191614be160a0602086015195015195614bcd61034b565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff8816608084015260a083015260c0820152614c31610386565b60e0820152614c3f856153e1565b15614d055790614c839260209260006040518096819582947f489a68f200000000000000000000000000000000000000000000000000000000845260048401614aa8565b03925af160009181614cd4575b50614c9e57836146c961461e565b929091925b51614ccb614caf61035b565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201529190565b614cf791925060203d602011614cfe575b614cef8183610308565b8101906149be565b9038614c90565b503d614ce5565b9050614d1084615437565b15614d9857614d536000926020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301614a97565b03925af160009181614d77575b50614d6e57836146c961461e565b92909192614ca3565b614d9191925060203d602011614cfe57614cef8183610308565b9038614d60565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff881660045260246000fd5b614e3a91975060203d602011613216576132078183610308565b9538614b89565b9091606082840312610107578151614e5881611308565b92602083015167ffffffffffffffff81116101075783019080601f8301121561010757815191614e878361036a565b91614e956040519384610308565b8383526020848301011161010757604092614eb6916020808501910161039b565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a08401526080614f33614eff604084015160a060c08801526101208701906103be565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103be565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b818110614fbc5750505061ffff909516602083015261033a9291606091614fa09063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101614f72565b906000916150b5938361505c73ffffffffffffffffffffffffffffffffffffffff61504167ffffffffffffffff60208701511667ffffffffffffffff166000526004602052604060002090565b541673ffffffffffffffffffffffffffffffffffffffff1690565b92604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601614ebc565b03925af1908115611f555760009060009261510a575b50156150d45750565b611d82906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611333565b905061512991503d806000833e6151218183610308565b810190614e41565b5090386150cb565b929161513b6137d9565b91808210156153cb576151556138de6138b8848489613596565b600160ff821603613d605750602182018181116153b55761517e6139e88260018601858a613654565b84528181101561539f57613a0f6138de6138b861519c93858a613596565b8201916022830190828211615389576132ac8260226151bd9301858a613654565b602085015281811015615373576151e0613a0f6138de6138b8602294868b613596565b83010191600183019082821161535d576132ac8260236152029301858a613654565b60408501528181101561534757615225613a0f6138de6138b8600194868b613596565b830101916001830190828211615331576132ac8260026152479301858a613654565b60608501528181101561531b5761526a613a0f6138de6138b8600194868b613596565b8301016001810192828411615305576132ac84600261528b9301858a613654565b608085015260038101928284116152ef576002916152b6613b266139ca6139c488600196898e613654565b010101948186116152d9576152d0926132ac928792613654565b60a08201529190565b635a102da160e11b600052601e60045260246000fd5b635a102da160e11b600052601d60045260246000fd5b635a102da160e11b600052601c60045260246000fd5b635a102da160e11b600052601b60045260246000fd5b635a102da160e11b600052601a60045260246000fd5b635a102da160e11b600052601960045260246000fd5b635a102da160e11b600052601860045260246000fd5b635a102da160e11b600052601760045260246000fd5b635a102da160e11b600052601660045260246000fd5b635a102da160e11b600052601560045260246000fd5b635a102da160e11b600052601460045260246000fd5b635a102da160e11b600052601360045260246000fd5b61540b7f331710310000000000000000000000000000000000000000000000000000000082615aea565b9081615425575b8161541b575090565b6107ab9150615a8a565b9050615430816159c4565b1590615412565b61540b7faff2afbf0000000000000000000000000000000000000000000000000000000082615aea565b9080601f8301121561010757815161547881611f84565b926154866040519485610308565b81845260208085019260051b82010192831161010757602001905b8282106154ae5750505090565b6020809183516154bd816111aa565b8152019101906154a1565b9060208282031261010757815167ffffffffffffffff8111610107576107ab9201615461565b95949060019460a09467ffffffffffffffff6155429573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906103be565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095909291906020848060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa938415611f55576000946156ab575b506155ec846153e1565b61560a575b5050505050508051156156015790565b506107ab613e4b565b60009596509061565f73ffffffffffffffffffffffffffffffffffffffff92604051988997889687957f89720a62000000000000000000000000000000000000000000000000000000008752600487016154ee565b0392165afa908115611f5557600091615688575b5061567d81615b4c565b3880808080806155f1565b6156a591503d806000833e61569d8183610308565b8101906154c8565b38615673565b6156c591945060203d602011613216576132078183610308565b92386155e2565b909160608284031261010757815167ffffffffffffffff811161010757836156f5918401615461565b92602083015167ffffffffffffffff811161010757604091615718918501615461565b92015160ff811681036101075790565b906157537f7909b5490000000000000000000000000000000000000000000000000000000082615aea565b80615854575b80615845575b61577d575b505061576e613e4b565b90615777613e2f565b90600090565b6040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9290921660048301526000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015611f555760008092819261581f575b506157f281615b4c565b6157fb83615b4c565b805115801590615813575b6158105750615764565b92565b5060ff82161515615806565b915061583d92503d8091833e6158358183610308565b8101906156cc565b9091386157e8565b5061584f81615a8a565b61575f565b5061585e816159c4565b15615759565b6002548110156120355760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b8181106158ad57505060009055565b806158ba60019285613eb5565b90549060031b1c600052818401602052600060408120550161589e565b60008181526003602052604090205461596257600254680100000000000000008110156102ae576159496159148260018594016002556002613eb5565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b60008281526001820160205260409020546159bd57805490680100000000000000008210156102ae57826159a6615914846001809601855584613eb5565b905580549260005201602052604060002055600190565b5050600090565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252615a24604483610308565b6179185a10615a60576020926000925191617530fa6000513d82615a54575b5081615a4d575090565b9050151590565b60201115915038615a43565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252615a24604483610308565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252615a24604483610308565b80519060005b828110615b5e57505050565b60018101808211612f82575b838110615b7a5750600101615b52565b73ffffffffffffffffffffffffffffffffffffffff615b99838561203a565b5116615bab610a04610a4d848761203a565b14615bb857600101615b6a565b611cc1615bc8610a4d848661203a565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
}

var OffRampABI = OffRampMetaData.ABI

var OffRampBin = OffRampMetaData.Bin

func DeployOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OffRampStaticConfig, maxGasBufferToUpdateState uint32) (common.Address, *types.Transaction, *OffRamp, error) {
	parsed, err := OffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OffRampBin), backend, staticConfig, maxGasBufferToUpdateState)
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

func (_OffRamp *OffRampCaller) GetmaxGasBufferToUpdateState(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getmaxGasBufferToUpdateState")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_OffRamp *OffRampSession) GetmaxGasBufferToUpdateState() (uint32, error) {
	return _OffRamp.Contract.GetmaxGasBufferToUpdateState(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetmaxGasBufferToUpdateState() (uint32, error) {
	return _OffRamp.Contract.GetmaxGasBufferToUpdateState(&_OffRamp.CallOpts)
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

func (_OffRamp *OffRampTransactor) SetMaxGasBufferToUpdateState(opts *bind.TransactOpts, maxGasBufferToUpdateState uint32) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setMaxGasBufferToUpdateState", maxGasBufferToUpdateState)
}

func (_OffRamp *OffRampSession) SetMaxGasBufferToUpdateState(maxGasBufferToUpdateState uint32) (*types.Transaction, error) {
	return _OffRamp.Contract.SetMaxGasBufferToUpdateState(&_OffRamp.TransactOpts, maxGasBufferToUpdateState)
}

func (_OffRamp *OffRampTransactorSession) SetMaxGasBufferToUpdateState(maxGasBufferToUpdateState uint32) (*types.Transaction, error) {
	return _OffRamp.Contract.SetMaxGasBufferToUpdateState(&_OffRamp.TransactOpts, maxGasBufferToUpdateState)
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

	GetmaxGasBufferToUpdateState(opts *bind.CallOpts) (uint32, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error)

	SetMaxGasBufferToUpdateState(opts *bind.TransactOpts, maxGasBufferToUpdateState uint32) (*types.Transaction, error)

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
