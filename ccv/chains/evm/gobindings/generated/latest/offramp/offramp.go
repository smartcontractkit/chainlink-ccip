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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getmaxGasBufferToUpdateState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setMaxGasBufferToUpdateState\",\"inputs\":[{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxGasBufferToUpdateStateUpdated\",\"inputs\":[{\"name\":\"oldMaxGasBufferToUpdateState\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"newMaxGasBufferToUpdateState\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"GasCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidGasLimitOverride\",\"inputs\":[{\"name\":\"messageGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100604052346102eb57604051615f3038819003601f8101601f191683016001600160401b038111848210176102f05783928291604052833981010360a081126102eb576080136102eb5760405190608082016001600160401b038111838210176102f05760405261007181610306565b825260208101519161ffff831683036102eb57602081019283526040820151916001600160a01b03831683036102eb57604082019283526060810151906001600160a01b03821682036102eb5760806100d1916060850193845201610306565b9233156102da5760015481516001600160a01b03161580156102c8575b6102b75783516001600160401b0316156102a65761ffff865116156102955783516001600160401b0390811660805282516001600160a01b0390811660a05284511660c052865161ffff1660e0528516958615610295577fffffff0000000000000000ff0000000000000000000000000000000000000000821633600160a81b600160e81b031981169190911760a897881b600160a81b600160e81b031617600155604080516001600160a01b031990941690911790961c6001600160401b0316825260208201969096527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e49509560809561ffff927f26b5c2c1b12f2b97d1f68df81ddac06bb274aecc2bbc3194218e1a48912cfaf89190a16040805195516001600160401b03168652905191909116602085015290516001600160a01b03908116918401919091529051166060820152a1604051615c15908161031b82396080518181816101580152611a18015260a0518181816101bb015261193f015260c0518181816101e301528181614b5601526155af015260e05181818161017f015261508a0152f35b632855a4d960e11b60005260046000fd5b63c656089560e01b60005260046000fd5b6342bcdf7f60e11b60005260046000fd5b5082516001600160a01b0316156100ee565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160401b03821682036102eb5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100f7578063181f5a77146100f257806320f81c88146100ed5780632d608aa0146100e85780633b81ab9b146100e357806349d8033e146100de5780635215505b146100d95780635643a782146100d457806361a10e59146100cf57806379ba5097146100ca5780638da5cb5b146100c55780638f4d0559146100c0578063e9d68a8e146100bb5763f2fde38b146100b657600080fd5b6111bf565b610ffb565b610fbf565b610f8b565b610ec0565b610e2d565b610a70565b610959565b6107ba565b6106e1565b6105ab565b6104ed565b6103e3565b61010c565b600091031261010757565b600080fd5b34610107576000600319360112610107576000606060405161012d81610292565b828152826020820152826040820152015261025f60405161014d81610292565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ae57604052565b610263565b60a0810190811067ffffffffffffffff8211176102ae57604052565b6101c0810190811067ffffffffffffffff8211176102ae57604052565b6020810190811067ffffffffffffffff8211176102ae57604052565b90601f601f19910116810190811067ffffffffffffffff8211176102ae57604052565b6040519061033a60c083610308565b565b6040519061033a60a083610308565b6040519061033a61010083610308565b6040519061033a604083610308565b67ffffffffffffffff81116102ae57601f01601f191660200190565b60405190610395602083610308565b60008252565b60005b8381106103ae5750506000910152565b818101518382015260200161039e565b90601f19601f6020936103dc8151809281875287808801910161039b565b0116010190565b346101075760006003193601126101075761025f60408051906104068183610308565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906103be565b9181601f840112156101075782359167ffffffffffffffff8311610107576020838186019501011161010757565b906020808351928381520192019060005b81811061048e5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610481565b916104e660ff916104d8604094979697606087526060870190610470565b908582036020870152610470565b9416910152565b346101075760206003193601126101075760043567ffffffffffffffff81116101075761057f61052d61052761025f933690600401610442565b9061388d565b61053b610140820151611295565b60601c9067ffffffffffffffff81511691610180820151906105798161ffff60a0860151169463ffffffff60806101a0830151519201511690613d9f565b93613f1d565b604093919351938493846104ba565b67ffffffffffffffff81160361010757565b359061033a8261058e565b34610107576020600319360112610107576004356105c88161058e565b6105d0614363565b67ffffffffffffffff8116801561066d577f26b5c2c1b12f2b97d1f68df81ddac06bb274aecc2bbc3194218e1a48912cfaf8916040917cffffffffffffffff0000000000000000000000000000000000000000006001549260a81b167fffffff0000000000000000ffffffffffffffffffffffffffffffffffffffffff83161760015567ffffffffffffffff83519260a81c1682526020820152a1005b7f50ab49b20000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156101075782359167ffffffffffffffff8311610107576020808501948460051b01011161010757565b63ffffffff81160361010757565b359061033a826106c8565b346101075760806003193601126101075760043567ffffffffffffffff811161010757610712903690600401610442565b9060243567ffffffffffffffff811161010757610733903690600401610697565b926044359367ffffffffffffffff85116101075761075861076d953690600401610697565b93909260643595610768876106c8565b611813565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600411156107a857565b61076f565b9060048210156107a85752565b34610107576020600319360112610107576004356000526006602052602060ff604060002054166107ee60405180926107ad565bf35b9080602083519182815201916020808360051b8301019401926000915b83831061081c57505050505090565b909192939460208061083a83601f19866001960301875289516103be565b9701930193019193929061080d565b6108b49173ffffffffffffffffffffffffffffffffffffffff825116815260208201511515602082015260806108a3610891604085015160a0604086015260a08501906107f0565b60608501518482036060860152610470565b920151906080818403910152610470565b90565b6040810160408252825180915260206060830193019060005b818110610939575050506020818303910152815180825260208201916020808360051b8301019401926000915b83831061090c57505050505090565b909192939460208061092a83601f1986600196030187528951610849565b970193019301919392906108fd565b825167ffffffffffffffff168552602094850194909201916001016108d0565b346101075760006003193601126101075760025461097681611f7b565b906109846040519283610308565b808252601f1961099382611f7b565b0160005b818110610a595750506109a981611fbf565b9060005b8181106109c557505061025f604051928392836108b7565b806109fd6109e46109d7600194615862565b67ffffffffffffffff1690565b6109ee8387612031565b9067ffffffffffffffff169052565b610a3d610a38610a1e610a108488612031565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b6120f9565b610a478287612031565b52610a528186612031565b50016109ad565b602090610a64611f93565b82828701015201610997565b346101075760206003193601126101075760043567ffffffffffffffff811161010757610aa1903690600401610697565b90610aaa614363565b6000905b828210610ab757005b610aca610ac583858461229e565b6123cc565b6020810191610ae46109d7845167ffffffffffffffff1690565b15610e0357610b26610b0d610b0d845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610df6575b610b7d5760808201949060005b86518051821015610ba757610b0d610b5683610b7093612031565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610b7d57600101610b3b565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610bdf57610b0d610b5683610bd293612031565b15610b7d57600101610bb7565b505095929491909394610bf586518251906144b7565b610c0a610a1e835167ffffffffffffffff1690565b90610c3a610c20845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610c4487615897565b606085019860005b8a518051821015610c9c5790610c6481602093612031565b518051928391012091158015610c8c575b610b7d57610c856001928b615966565b5001610c4c565b50610c95612481565b8214610c75565b5050976001975093610db3610dec946003610ddf95610dab7f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e996109d7979f610cf960408e610cf2610d3f945160018b0161263a565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610da1610d608d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b516002850161271f565b51910161271f565b610dd0610dcb6109d7835167ffffffffffffffff1690565b6158d5565b505167ffffffffffffffff1690565b92604051918291826127b3565b0390a20190610aae565b5060808201515115610b2e565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346101075760a06003193601126101075760043567ffffffffffffffff8111610107576101c06003198236030112610107576024359060443567ffffffffffffffff811161010757610e83903690600401610697565b926064359367ffffffffffffffff851161010757610ea861076d953690600401610697565b93909260843595610eb8876106c8565b600401612f7e565b346101075760006003193601126101075760005473ffffffffffffffffffffffffffffffffffffffff81163303610f61577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461010757600060031936011261010757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b3461010757600060031936011261010757602067ffffffffffffffff60015460a81c16604051908152f35b9060206108b4928181520190610849565b346101075760206003193601126101075767ffffffffffffffff6004356110218161058e565b611029611f93565b501660005260046020526040600020604051611044816102b3565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c16151560208201526001820180549061107c82611f7b565b9161108a6040519384610308565b80835260208301916000526020600020916000905b8282106110de5761025f866110cd60038a8960408501526110c260028201612098565b606085015201612098565b608082015260405191829182610fea565b604051600085546110ee81612045565b80845290600181169081156111605750600114611128575b506001928261111a85946020940382610308565b81520194019101909261109f565b6000878152602081209092505b81831061114a57505081016020016001611106565b6001816020925483868801015201920191611135565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050611106565b73ffffffffffffffffffffffffffffffffffffffff81160361010757565b346101075760206003193601126101075773ffffffffffffffffffffffffffffffffffffffff6004356111f1816111a1565b6111f9614363565b1633811461126b57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106112cd575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b8015150361010757565b9081602091031261010757516108b4816112ff565b6040513d6000823e3d90fd5b9060206108b49281815201906103be565b60409073ffffffffffffffffffffffffffffffffffffffff6108b4949316815281602082015201906103be565b9291926113748261036a565b916113826040519384610308565b829481845281830111610107578281602093846000960137010152565b9060048110156107a85760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061140257505050505090565b909192939460208061148983601f1986600196030187528951908151815260a06114786114666114546114428887015160c08a88015260c08701906103be565b604087015186820360408801526103be565b606086015185820360608701526103be565b608085015184820360808601526103be565b9201519060a08184039101526103be565b970193019301919392906113f3565b9160209082815201919060005b8181106114b25750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356114db816111a1565b1681520194019291016114a5565b601f8260209493601f19938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561010757016020813591019167ffffffffffffffff821161010757813603831361010757565b90602083828152019260208260051b82010193836000925b8484106115825750505050505090565b9091929394956020806115aa83601f1986600196030188526115a48b8861150a565b906114e9565b9801940194019294939190611572565b979694916117da906080956117e8956117c761033a9a9561164c8e60a081526115f060a08201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660c0820152604083015167ffffffffffffffff1660e0820152606083015163ffffffff16610100820152828c015163ffffffff1661012082015261014060a084015191019061ffff169052565b8d61016060c08301519101528d6101a061179361175d6117276116f16116bb61168760e08901516101c06101808a01526102608901906103be565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6089830301888a01526103be565b6101208801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60888303016101c08901526103be565b6101408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016101e08801526103be565b6101608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868303016102008701526103be565b6101808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60858303016102208601526113d6565b920151906102407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60828503019101526103be565b9260208d01528b830360408d0152611498565b9188830360608a015261155a565b94019063ffffffff169052565b806118066040926108b495946107ad565b81602082015201906103be565b959192939594909461182b60015460ff9060a01c1690565b611f5157611873740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b611885611880878361388d565b6143ae565b9561192660206118cb6118a36109d78b5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611f4c57600091611f1d575b50611ed25761199e61199a611990610a1e8a5167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b611e87576119b7610c20885167ffffffffffffffff1690565b6119e161199a60e08a01928351602081519101209060019160005201602052604060002054151590565b611e5057506101008701516014815114801590611e1f575b611de85750602087015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603611db15750828603611d7d576101408701516014815103611d43575063ffffffff84168015159081611d1d575b50611cd05790611a82913691611368565b6020815191012095611aa8611aa1886000526006602052604060002090565b5460ff1690565b611ab18161079e565b8015908115611cbc575b5015611c4a57611b5692611b5b959492611b4892611b11611ae68b6000526006602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957f61a10e590000000000000000000000000000000000000000000000000000000060208801528b8b602489016115ba565b03601f198101835282610308565b6143ba565b9015611c18577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b611bac84611ba7886000526006602052604060002090565b61139f565b611be8611bd66040611bc6885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b918360405194859416971695836117f5565b0390a461033a7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff600392611b8f565b611cb88787611c766040611c66835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150611cc98161079e565b1438611abb565b611cb884611ce560808a015163ffffffff1690565b7fdf2964df0000000000000000000000000000000000000000000000000000000060005263ffffffff90811660045216602452604490565b9050611d3c611d3360808a015163ffffffff1690565b63ffffffff1690565b1138611a71565b611d79906040519182917f8d666f600000000000000000000000000000000000000000000000000000000083526004830161132a565b0390fd5b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004869052602483905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b611d79906040519182917f55216e31000000000000000000000000000000000000000000000000000000008352306004840161133b565b50611e32611e2c82611295565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff163014156119f9565b611d7990516040519182917fa50bd1470000000000000000000000000000000000000000000000000000000083526004830161132a565b611cb8611e9c885167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611cb8611ee7885167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611f3f915060203d602011611f45575b611f378183610308565b810190611309565b38611970565b503d611f2d565b61131e565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116102ae5760051b60200190565b60405190611fa0826102b3565b6060608083600081526000602082015282604082015282808201520152565b90611fc982611f7b565b611fd66040519182610308565b828152601f19611fe68294611f7b565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80511561202c5760200190565b611ff0565b805182101561202c5760209160051b010190565b90600182811c9216801561208e575b602083101461205f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612054565b906040519182815491828252602082019060005260206000209260005b8181106120ca57505061033a92500383610308565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016120b5565b90604051612106816102b3565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c16151560208301526001810180549061214082611f7b565b9161214e6040519384610308565b80835260208301916000526020600020916000905b8282106121985750505050600360809261219392604086015261218860028201612098565b606086015201612098565b910152565b604051600085546121a881612045565b808452906001811690811561221a57506001146121e2575b50600192826121d485946020940382610308565b815201940191019092612163565b6000878152602081209092505b818310612204575050810160200160016121c0565b60018160209254838688010152019201916121ef565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b84019091019150600190506121c0565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4181360301821215610107570190565b901561202c57806108b49161225b565b9082101561202c576108b49160051b81019061225b565b359061033a826111a1565b359061033a826112ff565b9080601f83011215610107578160206108b493359101611368565b9080601f830112156101075781356122fd81611f7b565b9261230b6040519485610308565b81845260208085019260051b820101918383116101075760208201905b83821061233757505050505090565b813567ffffffffffffffff81116101075760209161235a878480948801016122cb565b815201910190612328565b9080601f8301121561010757813561237c81611f7b565b9261238a6040519485610308565b81845260208085019260051b82010192831161010757602001905b8282106123b25750505090565b6020809183356123c1816111a1565b8152019101906123a5565b60c081360312610107576123de61032b565b906123e8816122b5565b82526123f6602082016105a0565b6020830152612407604082016122c0565b6040830152606081013567ffffffffffffffff81116101075761242d90369083016122e6565b6060830152608081013567ffffffffffffffff8111610107576124539036908301612365565b608083015260a08101359067ffffffffffffffff82116101075761247991369101612365565b60a082015290565b6040516020810190600082526020815261249c604082610308565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106124dc575050565b600081556001016124d1565b9190601f81116124f757505050565b61033a926000526020600020906020601f840160051c83019310612523575b601f0160051c01906124d1565b9091508190612516565b919091825167ffffffffffffffff81116102ae576125558161254f8454612045565b846124e8565b6020601f82116001146125b35781906125a49394956000926125a8575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612572565b601f198216906125c884600052602060002090565b9160005b818110612622575095836001959697106125eb575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806125e1565b9192602060018192868b0151815501940192016125cc565b8151916801000000000000000083116102ae57815483835580841061269c575b50602061266e910191600052602060002090565b6000915b83831061267f5750505050565b60016020826126908394518661252d565b01920192019190612672565b8260005283602060002091820191015b8181106126b9575061265a565b806126c660019254612045565b806126d3575b50016126ac565b601f811183146126e95750600081555b386126cc565b61270d9083601f6126ff85600052602060002090565b920160051c820191016124d1565b600081815260208120818355556126e3565b81519167ffffffffffffffff83116102ae576801000000000000000083116102ae576020908254848455808510612796575b500190600052602060002060005b83811061276c5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161275f565b6127ad9084600052858460002091820191016124d1565b38612751565b906108b4916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a061282c612817606085015160c0608086015260e08501906107f0565b6080850151601f198583030184860152610470565b9201519060c0601f1982850301910152610470565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610107570180359067ffffffffffffffff821161010757602001918160051b3603831361010757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610107570180359067ffffffffffffffff82116101075760200191813603831361010757565b9160206108b49381815201916114e9565b919091357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106112cd575050565b356108b4816106c8565b356108b48161058e565b61ffff81160361010757565b356108b48161293f565b91909160c0818403126101075761296a61032b565b9281358452602082013567ffffffffffffffff8111610107578161298f9184016122cb565b6020850152604082013567ffffffffffffffff811161010757816129b49184016122cb565b6040850152606082013567ffffffffffffffff811161010757816129d99184016122cb565b6060850152608082013567ffffffffffffffff811161010757816129fe9184016122cb565b608085015260a082013567ffffffffffffffff811161010757612a2192016122cb565b60a0830152565b929190612a3481611f7b565b93612a426040519586610308565b602085838152019160051b8101918383116101075781905b838210612a68575050505050565b813567ffffffffffffffff811161010757602091612a898784938701612955565b815201910190612a5a565b9082101561202c57612aab9160051b810190612895565b9091565b9081602091031261010757516108b4816111a1565b60409073ffffffffffffffffffffffffffffffffffffffff6108b4959316815281602082015201916114e9565b359061033a8261293f565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561010757016020813591019167ffffffffffffffff8211610107578160051b3603831361010757565b90602083828152019060208160051b85010193836000915b838310612b775750505050505090565b909192939495601f1982820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4184360301811215610107576020612c5a600193868394019081358152612c4c612c41612c26612c0b612bf0612be08988018861150a565b60c08b89015260c08801916114e9565b612bfd604088018861150a565b9087830360408901526114e9565b612c18606087018761150a565b9086830360608801526114e9565b612c33608086018661150a565b9085830360808701526114e9565b9260a081019061150a565b9160a08185039101526114e9565b980196019493019190612b67565b612ee06108b49593949260608352612c9460608401612c86836105a0565b67ffffffffffffffff169052565b612cb4612ca3602083016105a0565b67ffffffffffffffff166080850152565b612cd4612cc3604083016105a0565b67ffffffffffffffff1660a0850152565b612cf0612ce3606083016106d6565b63ffffffff1660c0850152565b612d0c612cff608083016106d6565b63ffffffff1660e0850152565b612d27612d1b60a08301612af1565b61ffff16610100850152565b60c0810135610120840152612eaf612ea3612e64612e25612de6612da7612d68612d5460e089018961150a565b6101c06101408d01526102208c01916114e9565b612d7661010089018961150a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d01526114e9565b612db561012088018861150a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c01526114e9565b612df461014087018761150a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b01526114e9565b612e3361016086018661150a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a01526114e9565b612e72610180850185612afc565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152612b4f565b916101a081019061150a565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0858403016102008601526114e9565b93602082015260408185039101526114e9565b604051906040820182811067ffffffffffffffff8211176102ae5760405260006020838281520152565b90612f2782611f7b565b612f346040519182610308565b828152601f19612f448294611f7b565b019060005b828110612f5557505050565b602090612f60612ef3565b82828501015201612f49565b91908203918211612f7957565b6124a2565b9391949092943033036134de576000612f9b610180870187612841565b90506133ff575b613024612fbf611e2c612fb96101408a018a612895565b906128f7565b92612fe984612fd26101a08b018b612895565b9050612fe3611d3360808d0161292b565b90613d9f565b98899160a0612ff78b612935565b8761301e8d61301661300d610180830183612841565b9690920161294b565b943691612a28565b91614752565b92909860005b8a518110156132145780602061304d610b0d610b0d8f61309996610b5691612031565b61306261305a848a612031565b518a8c612a94565b91906040518096819482937fc3a7ded6000000000000000000000000000000000000000000000000000000008452600484016128e6565b03915afa918215611f4c576000926131e4575b5073ffffffffffffffffffffffffffffffffffffffff821615613187576130de6130d68288612031565b51888a612a94565b929073ffffffffffffffffffffffffffffffffffffffff82163b15610107576000918b8373ffffffffffffffffffffffffffffffffffffffff8f613151604051998a97889687947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601612c68565b0393165af1918215611f4c5760019261316c575b500161302a565b8061317b600061318193610308565b806100fc565b38613165565b856131af89898f856131a2610b56611d79986131a894612031565b95612031565b5191612a94565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501612ac4565b61320691925060203d811161320d575b6131fe8183610308565b810190612aaf565b90386130ac565b503d6131f4565b5095979350959350965050613237613230610180840184612841565b9050612f1d565b95613246610180840184612841565b90506132ed575b506132e65761033a946132b661326283612935565b6132a36132aa613276610120870187612895565b6132846101a0890189612895565b94909561328f61033c565b9c8d5267ffffffffffffffff1660208d0152565b3691611368565b60408901523691611368565b6060860152608085015263ffffffff8216156132d3575091614ff2565b6132e0915060800161292b565b91614ff2565b5050505050565b613347613307613301610180860186612841565b9061228e565b613315610120860186612895565b61334161332188612935565b9261333961333160a08b0161294b565b953690612955565b923691611368565b90614ac1565b91906133528961201f565b5273ffffffffffffffffffffffffffffffffffffffff61338c611e2c612fb96133826133016101808a018a612841565b6080810190612895565b921673ffffffffffffffffffffffffffffffffffffffff8316036133b1575b5061324d565b6133e56133ea926133df6133c48b61201f565b515173ffffffffffffffffffffffffffffffffffffffff1690565b9061464c565b612f6c565b60206133f58861201f565b51015238806133ab565b506014613420613416613301610180890189612841565b6060810190612895565b9050036134ca57601461343d613382613301610180890189612841565b9050036134805761347b613461611e2c612fb96133826133016101808b018b612841565b6133df611e2c612fb96134166133016101808c018c612841565b612fa2565b613494613382613301610180880188612841565b90611d796040519283927f8d666f60000000000000000000000000000000000000000000000000000000008452600484016128e6565b613494613416613301610180880188612841565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190613515826102cf565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b901561202c5790565b906043101561202c5760430190565b9082101561202c570190565b906009116101075760010190600890565b906011116101075760090190600890565b906019116101075760110190600890565b90601d116101075760190190600490565b9060211161010757601d0190600490565b906023116101075760210190600290565b906043116101075760230190602090565b909291928360441161010757831161010757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b90939293848311610107578411610107578101920390565b919091357fffffffffffffffff00000000000000000000000000000000000000000000000081169260088110613697575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106136fd575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110613763575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b3590602081106137a3575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff8211176102ae57604052606060a0836000815282602082015282604082015282808201528260808201520152565b604080519091906138238382610308565b6001815291601f19018260005b82811061383c57505050565b6020906138476137d0565b82828501015201613830565b60405190613862602083610308565b600080835282815b82811061387657505050565b6020906138816137d0565b8282850101520161386a565b90613896613508565b91604d8210613d87576138db6138d56138af8484613575565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613d5757506139146139066139006138fa8585613599565b90613663565b60c01c90565b67ffffffffffffffff168452565b6139386139276139006138fa85856135aa565b67ffffffffffffffff166020850152565b61395c61394b6139006138fa85856135bb565b67ffffffffffffffff166040850152565b61398861397b61397561396f85856135cc565b906136c9565b60e01c90565b63ffffffff166060850152565b6139a861399b61397561396f85856135dd565b63ffffffff166080850152565b6139d26139c76139c16139bb85856135ee565b9061372f565b60f01c90565b61ffff1660a0850152565b6139e56139df83836135ff565b90613795565b60c08401528160431015613d4157613a0c613a066138d56138af858561357e565b60ff1690565b9081604401838111613d2b57613a266132a3828685613610565b60e086015283811015613d1557613a066138d56138af613a4793878661358d565b8201916045830190848211613d00576132a3826045613a689301878661364b565b61010086015283811015613cea57613a8c613a066138d56138af604594888761358d565b830101916001830190848211613cd4576132a3826046613aae9301878661364b565b61012086015283811015613cbe57613ad2613a066138d56138af600194888761358d565b830101916001830190848211613ca8576132a3826002613af49301878661364b565b6101408601526003830192848411613c9257613b24613b1d6139c16139bb876001968a8961364b565b61ffff1690565b0101916002830190848211613c7c576132a382613b4292878661364b565b6101608601526004830190848211613c66576139c16139bb83613b6693888761364b565b9261ffff8294168015600014613c1557505050613b81613853565b6101808501525b6002820191838311613bff5780613bac613b1d6139c16139bb876002968a8961364b565b010191838311613be957826132a39185613bc59461364b565b6101a084015203613bd35790565b635a102da160e11b600052601260045260246000fd5b635a102da160e11b600052601160045260246000fd5b635a102da160e11b600052601060045260246000fd5b6002919294508190613c38613c28613812565b966101808a01978852888761512f565b9490965196613c47869861201f565b5201010114613b8857635a102da160e11b600052600f60045260246000fd5b635a102da160e11b600052600e60045260246000fd5b635a102da160e11b600052600d60045260246000fd5b635a102da160e11b600052600c60045260246000fd5b635a102da160e11b600052600b60045260246000fd5b635a102da160e11b600052600a60045260246000fd5b635a102da160e11b600052600960045260246000fd5b635a102da160e11b600052600860045260246000fd5b635a102da160e11b6000526004805260246000fd5b635a102da160e11b600052600360045260246000fd5b635a102da160e11b600052600260045260246000fd5b635a102da160e11b600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b635a102da160e11b600052611cb86024906000600452565b15919082613e1d575b508115613e13575b8115613dba575090565b9050613de67f85572ffb0000000000000000000000000000000000000000000000000000000082615ae8565b9081613e01575b81613df757501590565b61199a9150615a88565b9050613e0c816159c2565b1590613ded565b803b159150613db0565b15915038613da8565b60405190613e35602083610308565b6000808352366020840137565b60408051909190613e538382610308565b6001815291601f1901366020840137565b9060018201809211612f7957565b91908201809211612f7957565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114612f795760010190565b805482101561202c5760005260206000200190600090565b8015612f79577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff168015612f79577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b94929391606091600093613f2f613e26565b968351806142e4575b5050156142d0575051156142bf575050613f50613e26565b613f58613e26565b6000935b6002613fa5613f8a6003613f848a67ffffffffffffffff166000526004602052604060002090565b01612098565b9767ffffffffffffffff166000526004602052604060002090565b0195613fbf613fb78551845190613e72565b825190613e72565b90613fd4613fcf89548094613e72565b611fbf565b9788966000805b885182101561402e579061402660019261400b613ffb610b56858e612031565b9161400581613e7f565b9d612031565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018a98613fdb565b91939597505097909294976000905b8751821015614075579061406d60019261400b61405d610b56858d612031565b9161406781613e7f565b9c612031565b01899761403d565b9750509193969092945060005b85518110156140b757806140b161409e610b566001948a612031565b61400b6140aa8b613e7f565b9a8d612031565b01614082565b50929590935093909360005b82811061423b575b50509091929350600090815b8181106141995750508452805160005b85518110156141935760005b828110614104575b506001016140e7565b614111610b568286612031565b73ffffffffffffffffffffffffffffffffffffffff614136610b0d610b56868c612031565b91161461414b5761414690613e7f565b6140f3565b9161415861417091613ec4565b9261400b614169610b568688612031565b9186612031565b60ff841661417f575b386140fb565b9261418b600191613eef565b939050614179565b50815291565b6141a6610b568289612031565b73ffffffffffffffffffffffffffffffffffffffff8116801561423157600090815b8a878210614204575b5050509060019291156141e7575b505b016140d7565b6141fe9061400b6141f787613e7f565b968b612031565b386141df565b614215610b0d610b56848694612031565b14614222576001016141c8565b5060019150819050388a6141d1565b50506001906141e1565b61424b610b0d610b56838b612031565b15614258576001016140c3565b50909192939460005b828110614273578695949392506140cb565b806142b96142a661428660019486613eac565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b61400b6142b288613e7f565b978c612031565b01614261565b6142ca939193613e42565b91613f5c565b9150506142de915084615726565b93613f5c565b9091975060018103614336575061432e9061430d611e2c866143058761201f565b510151611295565b906143178561201f565b51518a60a06143258861201f565b51015193615545565b953880613f38565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361438457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6143b6613508565b5090565b90604051916143ca60c084610308565b60848352602083019060a03683375a60015490919060a81c67ffffffffffffffff16908183111561442e57614403600093928493612f6c565b82602083519301913090f1903d9060848211614425575b6000908286523e9190565b6084915061441a565b611cb86144696144415a63ffffffff1690565b60e01b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b8051916144c5815184613e72565b9283156145f25760005b8481106144dd575050505050565b818110156145d7576144f2610b568286612031565b73ffffffffffffffffffffffffffffffffffffffff81168015610b7d5761451883613e64565b87811061452a575050506001016144cf565b848110156145a75773ffffffffffffffffffffffffffffffffffffffff614554610b56838a612031565b16821461456357600101614518565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6145d2610b566145cc8885612f6c565b89612031565b614554565b6145ed610b566145e78484612f6c565b85612031565b6144f2565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b3d15614647573d9061462d8261036a565b9161463b6040519384610308565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181614701575b506146fd57826146c761461c565b90611d796040519283927f9fe2f95a0000000000000000000000000000000000000000000000000000000084526004840161133b565b9150565b90916020823d602011614730575b8161471c60209383610308565b8101031261472d57505190386146b9565b80fd5b3d915061470f565b919081101561202c5760051b0190565b356108b4816111a1565b90614761949596939291613f1d565b9193909261476e82611fbf565b9261477883611fbf565b94600091825b8851811015614872576000805b8a8884898285106147fb575b5050505050156147a95760010161477e565b6147b9610b56611cb8928b612031565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610b5661482d926131a26148288873ffffffffffffffffffffffffffffffffffffffff97610b0d96614738565b614748565b91161461483c5760010161478b565b6001915061485b614851614828838b8b614738565b61400b888c612031565b6148676141f787613e7f565b52388a888489614797565b509097965094939291909460ff811690816000985b8a518a10156149445760005b8b8782108061493b575b1561492e5773ffffffffffffffffffffffffffffffffffffffff6148cf610b0d610b568f6131a2614828888f8f614738565b9116146148e4576148df90613e7f565b614893565b93996148f460019294939b613ec4565b94614910614906614828838b8b614738565b61400b8d8c612031565b61492361491c8c613e7f565b9b8b612031565b525b01989091614887565b5050919098600190614925565b5085151561489d565b98509250939594975091508161496d575050508151810361496457509190565b80825283529190565b611cb8929161497b91612f6c565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b60405190610395826102ec565b9081602091031261010757604051906149d4826102ec565b51815290565b6108b49160e0614a84614a726149fb855161010086526101008601906103be565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152614a606080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526103be565b60c085015184820360c08601526103be565b9201519060e08184039101526103be565b9060206108b49281815201906149da565b9061ffff6104e66020929594956040855260408501906149da565b92939193614acd612ef3565b50614ade611e2c6080860151611295565b91614aef611e2c6060870151611295565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909690956020878060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa968715611f4c57600097614e1e575b5073ffffffffffffffffffffffffffffffffffffffff8716948515614dda57614bae6149af565b50614bfc825191614bdf60a0602086015195015195614bcb61034b565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff8816608084015260a083015260c0820152614c2f610386565b60e0820152614c3d856153df565b15614d035790614c819260209260006040518096819582947f489a68f200000000000000000000000000000000000000000000000000000000845260048401614aa6565b03925af160009181614cd2575b50614c9c57836146c761461c565b929091925b51614cc9614cad61035b565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201529190565b614cf591925060203d602011614cfc575b614ced8183610308565b8101906149bc565b9038614c8e565b503d614ce3565b9050614d0e84615435565b15614d9657614d516000926020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301614a95565b03925af160009181614d75575b50614d6c57836146c761461c565b92909192614ca1565b614d8f91925060203d602011614cfc57614ced8183610308565b9038614d5e565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff881660045260246000fd5b614e3891975060203d60201161320d576131fe8183610308565b9538614b87565b9091606082840312610107578151614e56816112ff565b92602083015167ffffffffffffffff81116101075783019080601f8301121561010757815191614e858361036a565b91614e936040519384610308565b8383526020848301011161010757604092614eb4916020808501910161039b565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a08401526080614f31614efd604084015160a060c08801526101208701906103be565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103be565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b818110614fba5750505061ffff909516602083015261033a9291606091614f9e9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101614f70565b906000916150b3938361505a73ffffffffffffffffffffffffffffffffffffffff61503f67ffffffffffffffff60208701511667ffffffffffffffff166000526004602052604060002090565b541673ffffffffffffffffffffffffffffffffffffffff1690565b92604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f00000000000000000000000000000000000000000000000000000000000000009060048601614eba565b03925af1908115611f4c57600090600092615108575b50156150d25750565b611d79906040519182917f0a8d6e8c0000000000000000000000000000000000000000000000000000000083526004830161132a565b905061512791503d806000833e61511f8183610308565b810190614e3f565b5090386150c9565b92916151396137d0565b91808210156153c9576151536138d56138af84848961358d565b600160ff821603613d575750602182018181116153b35761517c6139df8260018601858a61364b565b84528181101561539d57613a066138d56138af61519a93858a61358d565b8201916022830190828211615387576132a38260226151bb9301858a61364b565b602085015281811015615371576151de613a066138d56138af602294868b61358d565b83010191600183019082821161535b576132a38260236152009301858a61364b565b60408501528181101561534557615223613a066138d56138af600194868b61358d565b83010191600183019082821161532f576132a38260026152459301858a61364b565b60608501528181101561531957615268613a066138d56138af600194868b61358d565b8301016001810192828411615303576132a38460026152899301858a61364b565b608085015260038101928284116152ed576002916152b4613b1d6139c16139bb88600196898e61364b565b010101948186116152d7576152ce926132a392879261364b565b60a08201529190565b635a102da160e11b600052601e60045260246000fd5b635a102da160e11b600052601d60045260246000fd5b635a102da160e11b600052601c60045260246000fd5b635a102da160e11b600052601b60045260246000fd5b635a102da160e11b600052601a60045260246000fd5b635a102da160e11b600052601960045260246000fd5b635a102da160e11b600052601860045260246000fd5b635a102da160e11b600052601760045260246000fd5b635a102da160e11b600052601660045260246000fd5b635a102da160e11b600052601560045260246000fd5b635a102da160e11b600052601460045260246000fd5b635a102da160e11b600052601360045260246000fd5b6154097f331710310000000000000000000000000000000000000000000000000000000082615ae8565b9081615423575b81615419575090565b6108b49150615a88565b905061542e816159c2565b1590615410565b6154097faff2afbf0000000000000000000000000000000000000000000000000000000082615ae8565b9080601f8301121561010757815161547681611f7b565b926154846040519485610308565b81845260208085019260051b82010192831161010757602001905b8282106154ac5750505090565b6020809183516154bb816111a1565b81520191019061549f565b9060208282031261010757815167ffffffffffffffff8111610107576108b4920161545f565b95949060019460a09467ffffffffffffffff6155409573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906103be565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095909291906020848060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa938415611f4c576000946156a9575b506155ea846153df565b615608575b5050505050508051156155ff5790565b506108b4613e42565b60009596509061565d73ffffffffffffffffffffffffffffffffffffffff92604051988997889687957f89720a62000000000000000000000000000000000000000000000000000000008752600487016154ec565b0392165afa908115611f4c57600091615686575b5061567b81615b4a565b3880808080806155ef565b6156a391503d806000833e61569b8183610308565b8101906154c6565b38615671565b6156c391945060203d60201161320d576131fe8183610308565b92386155e0565b909160608284031261010757815167ffffffffffffffff811161010757836156f391840161545f565b92602083015167ffffffffffffffff81116101075760409161571691850161545f565b92015160ff811681036101075790565b906157517f7909b5490000000000000000000000000000000000000000000000000000000082615ae8565b80615852575b80615843575b61577b575b505061576c613e42565b90615775613e26565b90600090565b6040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9290921660048301526000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015611f4c5760008092819261581d575b506157f081615b4a565b6157f983615b4a565b805115801590615811575b61580e5750615762565b92565b5060ff82161515615804565b915061583b92503d8091833e6158338183610308565b8101906156ca565b9091386157e6565b5061584d81615a88565b61575d565b5061585c816159c2565b15615757565b60025481101561202c5760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b8181106158ab57505060009055565b806158b860019285613eac565b90549060031b1c600052818401602052600060408120550161589c565b60008181526003602052604090205461596057600254680100000000000000008110156102ae576159476159128260018594016002556002613eac565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b60008281526001820160205260409020546159bb57805490680100000000000000008210156102ae57826159a4615912846001809601855584613eac565b905580549260005201602052604060002055600190565b5050600090565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252615a22604483610308565b6179185a10615a5e576020926000925191617530fa6000513d82615a52575b5081615a4b575090565b9050151590565b60201115915038615a41565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252615a22604483610308565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252615a22604483610308565b80519060005b828110615b5c57505050565b60018101808211612f79575b838110615b785750600101615b50565b73ffffffffffffffffffffffffffffffffffffffff615b978385612031565b5116615ba9610b0d610b568487612031565b14615bb657600101615b68565b611cb8615bc6610b568486612031565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
}

var OffRampABI = OffRampMetaData.ABI

var OffRampBin = OffRampMetaData.Bin

func DeployOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OffRampStaticConfig, maxGasBufferToUpdateState uint64) (common.Address, *types.Transaction, *OffRamp, error) {
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

func (_OffRamp *OffRampCaller) GetmaxGasBufferToUpdateState(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OffRamp.contract.Call(opts, &out, "getmaxGasBufferToUpdateState")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OffRamp *OffRampSession) GetmaxGasBufferToUpdateState() (uint64, error) {
	return _OffRamp.Contract.GetmaxGasBufferToUpdateState(&_OffRamp.CallOpts)
}

func (_OffRamp *OffRampCallerSession) GetmaxGasBufferToUpdateState() (uint64, error) {
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

func (_OffRamp *OffRampTransactor) SetMaxGasBufferToUpdateState(opts *bind.TransactOpts, maxGasBufferToUpdateState uint64) (*types.Transaction, error) {
	return _OffRamp.contract.Transact(opts, "setMaxGasBufferToUpdateState", maxGasBufferToUpdateState)
}

func (_OffRamp *OffRampSession) SetMaxGasBufferToUpdateState(maxGasBufferToUpdateState uint64) (*types.Transaction, error) {
	return _OffRamp.Contract.SetMaxGasBufferToUpdateState(&_OffRamp.TransactOpts, maxGasBufferToUpdateState)
}

func (_OffRamp *OffRampTransactorSession) SetMaxGasBufferToUpdateState(maxGasBufferToUpdateState uint64) (*types.Transaction, error) {
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
	OldMaxGasBufferToUpdateState uint64
	NewMaxGasBufferToUpdateState uint64
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
	return common.HexToHash("0x26b5c2c1b12f2b97d1f68df81ddac06bb274aecc2bbc3194218e1a48912cfaf8")
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

	GetmaxGasBufferToUpdateState(opts *bind.CallOpts) (uint64, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []OffRampSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte, gasLimitOverride uint32) (*types.Transaction, error)

	SetMaxGasBufferToUpdateState(opts *bind.TransactOpts, maxGasBufferToUpdateState uint64) (*types.Transaction, error)

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
