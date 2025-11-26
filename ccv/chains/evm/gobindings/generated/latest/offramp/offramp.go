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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f615df238819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615b9e908161025482396080518181816101560152610bea015260a0518181816101b90152610b2c015260c0518181816101e10152818161393d0152615767015260e05181818161017d01526127890152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063176133d1146100d2578063181f5a77146100cd57806320f81c88146100c85780635215505b146100c35780636b8be52c146100be57806379ba5097146100b95780637ce1552a146100b45780638da5cb5b146100af578063e9d68a8e146100aa578063f054ac57146100a55763f2fde38b146100a057600080fd5b6117eb565b611497565b6113b7565b611354565b6112b5565b6110e4565b610978565b610825565b610657565b61052f565b610292565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610384565b828152826020820152826040820152015261025d60405161014b81610384565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760243560443567ffffffffffffffff81116100e757610323903690600401610261565b606435939167ffffffffffffffff85116100e757610348610353953690600401610261565b9490936004016123a7565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176103a057604052565b610355565b60a0810190811067ffffffffffffffff8211176103a057604052565b6020810190811067ffffffffffffffff8211176103a057604052565b6101c0810190811067ffffffffffffffff8211176103a057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103a057604052565b6040519061044a60c0836103fa565b565b6040519061044a60a0836103fa565b6040519061044a610100836103fa565b6040519061044a6040836103fa565b67ffffffffffffffff81116103a057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906104c36020836103fa565b60008252565b60005b8381106104dc5750506000910152565b81810151838201526020016104cc565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610528815180928187528780880191016104c9565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061057081836103fa565b601182527f4f666652616d7020312e372e302d6465760000000000000000000000000000006020830152519182916020835260208301906104ec565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106105f85750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016105eb565b9161065060ff916106426040949796976060875260608701906105da565b9085820360208701526105da565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106e86106b56106af61025d9336906004016105ac565b9061424b565b67ffffffffffffffff815116906106d0610140820151612893565b60601c61ffff60a06101808401519301511692614908565b60409391935193849384610624565b6107629173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061075161073f604085015160a0604086015260a08501906104ec565b606085015184820360608601526105da565b9201519060808184039101526105da565b90565b6040810160408252825180915260206060830193019060005b818110610805575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106107ba57505050505090565b90919293946020806107f6837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516106f7565b970193019301919392906107ab565b825167ffffffffffffffff1685526020948501949092019160010161077e565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461086081611a0a565b9061086e60405192836103fa565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061089b82611a0a565b0160005b8181106109615750506108b1816128f7565b9060005b8181106108cd57505061025d60405192839283610765565b806109056108ec6108df600194615a1a565b67ffffffffffffffff1690565b6108f68387611ba2565b9067ffffffffffffffff169052565b6109456109406109266109188488611ba2565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b612aba565b61094f8287611ba2565b5261095a8186611ba2565b50016108b5565b60209061096c6128cb565b8282870101520161089f565b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576109c79036906004016105ac565b60243567ffffffffffffffff81116100e7576109e7903690600401610261565b92909160443567ffffffffffffffff81116100e757610a0a903690600401610261565b939092610a1960055460ff1690565b6110ba57610a6b90610a5160017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff006005541617600555565b610a63610a5e858361424b565b614d2e565b9336916111ea565b6020815191012093610b136020610ab8610a906108df875167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156110b557600091611086575b5061103b57610b77610926845167ffffffffffffffff1690565b8054610b8b9060a01c60ff161590565b1590565b610ff05760e084016001815160208151910120920191610baa83612a3b565b6020815191012003610fb95750506101008301516014815114801590610f8e575b610f575750602083015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603610f205750808603610eee5761014083019384516014815103610eb45750610c7f610c49855167ffffffffffffffff1690565b6040860196610c60885167ffffffffffffffff1690565b610c79610c736101208a01519351612893565b60601c90565b92614d3a565b96610c9e610c97896000526006602052604060002090565b5460ff1690565b610ca781611289565b8015908115610ea0575b5015610e3157610dda610dcb610dc696610918610daa7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df298610da58d8f9967ffffffffffffffff9b610d7991610dec9b610d43610d188f6000526006602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957f176133d10000000000000000000000000000000000000000000000000000000060208801528b60248801612d92565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103fa565b614dbb565b969015610e1c576002998a916000526006602052604060002090565b612bac565b965167ffffffffffffffff1690565b91836040519485941697169583612fb3565b0390a46103537fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060055416600555565b6003998a916000526006602052604060002090565b610e9c8787610e5a610e4b895167ffffffffffffffff1690565b915167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150610ead81611289565b1438610cb1565b610eea906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301612396565b0390fd5b7f88f80aa200000000000000000000000000000000000000000000000000000000600052600486905260245260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b610eea906040519182917f55216e310000000000000000000000000000000000000000000000000000000083523060048401612b7f565b50610f9b610c7382612893565b73ffffffffffffffffffffffffffffffffffffffff16301415610bcb565b5190610eea6040519283927f2130095800000000000000000000000000000000000000000000000000000000845260048401612b5a565b610e9c611005855167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610e9c611050845167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b6110a8915060203d6020116110ae575b6110a081836103fa565b810190612b45565b38610b5d565b503d611096565b611c36565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff811633036111a3577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff8116036100e757565b359061044a826111cd565b9291926111f68261047a565b9161120460405193846103fa565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e757816020610762933591016111ea565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561129357565b61125a565b9060048210156112935752565b60208101929161044a9190611298565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356112f0816111cd565b602435906112fd826111cd565b6044359167ffffffffffffffff83116100e757611321611334933690600401611221565b906064359261132f8461123c565b614d3a565b600052600660205261025d60ff60406000205416604051918291826112a5565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b9060206107629281815201906106f7565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff6004356113fb816111cd565b6114036128cb565b5016600052600460205261025d6040600020611486600360405192611427846103a5565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c161515602085015260405161146a816114638160018601612999565b03826103fa565b604085015261147b60028201612aa6565b606085015201612aa6565b6080820152604051918291826113a6565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576114e6903690600401610261565b906114ef614e77565b6000905b8282106114fc57005b61150f61150a83858461218f565b61304e565b60208101916115296108df845167ffffffffffffffff1690565b156117c15761156b611552611552845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b1580156117b4575b6115c25760808201949060005b865180518210156115ec5761155261159b836115b593611ba2565b5173ffffffffffffffffffffffffffffffffffffffff1690565b156115c257600101611580565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050939194929060009460a08701955b865180518210156116245761155261159b8361161793611ba2565b156115c2576001016115fc565b5050949590956060820180518051801591821561179e575b50506115c2576108df61159b9461175b611794946117518a6117476117879761173e6116fd60019f9c6116937f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259e51895190614ec2565b6116a86109268b5167ffffffffffffffff1690565b9e8f6116b76040840151151590565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b8d9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518d8c01613180565b5160028a016132da565b51600388016132da565b6117786117736108df835167ffffffffffffffff1690565b615a4f565b505167ffffffffffffffff1690565b926040519182918261336e565b0390a201906114f3565b6020012090506117ac613103565b14388061163c565b5060808201515115611573565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561183b8161123c565b611843614e77565b163381146118b557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611964575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b35610762816111cd565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b61ffff8116036100e757565b35610762816119f4565b67ffffffffffffffff81116103a05760051b60200190565b91909160c0818403126100e757611a3761043b565b9281358452602082013567ffffffffffffffff81116100e75781611a5c918401611221565b6020850152604082013567ffffffffffffffff81116100e75781611a81918401611221565b6040850152606082013567ffffffffffffffff81116100e75781611aa6918401611221565b6060850152608082013567ffffffffffffffff81116100e75781611acb918401611221565b608085015260a082013567ffffffffffffffff81116100e757611aee9201611221565b60a0830152565b929190611b0181611a0a565b93611b0f60405195866103fa565b602085838152019160051b8101918383116100e75781905b838210611b35575050505050565b813567ffffffffffffffff81116100e757602091611b568784938701611a22565b815201910190611b27565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805115611b9d5760200190565b611b61565b8051821015611b9d5760209160051b010190565b90821015611b9d57611bcd9160051b8101906118df565b9091565b908160209103126100e757516107628161123c565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916020610762938181520191611be6565b6040513d6000823e3d90fd5b60409073ffffffffffffffffffffffffffffffffffffffff61076295931681528160208201520191611be6565b63ffffffff8116036100e757565b359061044a82611c6f565b359061044a826119f4565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310611d5e5750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e7576020611e5f600193868394019081358152611e51611e46611e2b611e10611df5611de589880188611c93565b60c08b89015260c0880191611be6565b611e026040880188611c93565b908783036040890152611be6565b611e1d6060870187611c93565b908683036060880152611be6565b611e386080860186611c93565b908583036080870152611be6565b9260a0810190611c93565b9160a0818503910152611be6565b980196019493019190611d4e565b6120e56107629593949260608352611e9960608401611e8b836111df565b67ffffffffffffffff169052565b611eb9611ea8602083016111df565b67ffffffffffffffff166080850152565b611ed9611ec8604083016111df565b67ffffffffffffffff1660a0850152565b611ef5611ee860608301611c7d565b63ffffffff1660c0850152565b611f11611f0460808301611c7d565b63ffffffff1660e0850152565b611f2c611f2060a08301611c88565b61ffff16610100850152565b60c08101356101208401526120b46120a861206961202a611feb611fac611f6d611f5960e0890189611c93565b6101c06101408d01526102208c0191611be6565b611f7b610100890189611c93565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d0152611be6565b611fba610120880188611c93565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c0152611be6565b611ff9610140870187611c93565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b0152611be6565b612038610160860186611c93565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a0152611be6565b612077610180850185611ce3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152611d36565b916101a0810190611c93565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa085840301610200860152611be6565b9360208201526040818503910152611be6565b604051906040820182811067ffffffffffffffff8211176103a05760405260006020838281520152565b9061212c82611a0a565b61213960405191826103fa565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06121678294611a0a565b019060005b82811061217857505050565b6020906121836120f8565b8282850101520161216c565b9190811015611b9d5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b3561076281611c6f565b801515036100e757565b90916060828403126100e75781516121fa816121d9565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e7578151916122298361047a565b9161223760405193846103fa565b838352602084830101116100e75760409261225891602080850191016104c9565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806122d56122a1604084015160a060c08801526101208701906104ec565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526104ec565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b81811061235e5750505061ffff909516602083015261044a92916060916123429063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101612314565b9060206107629281815201906104ec565b92949593909195303303612869576123f596612410612402976123da610c736123d46101408a018a6118df565b90611930565b9261240a6123e789611996565b856101808b019d8e8c6119a0565b939060a08d019e8f611a00565b943691611af5565b916134ee565b97909860005b8a518110156125d35761248560208b61244e6124468f61244061155261155261159b8a8095611ba2565b93611ba2565b51898b611bb6565b91906040518095819482937fc3a7ded600000000000000000000000000000000000000000000000000000000845260048401611c25565b03915afa80156110b55773ffffffffffffffffffffffffffffffffffffffff916000916125a5575b5016908115612548576124cb6124c3828d611ba2565b518789611bb6565b9290813b156100e7576000918a838d612513604051988996879586947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601611e6d565b03925af19182156110b55760019261252d575b5001612416565b8061253c6000612542936103fa565b806100dc565b38612526565b8a61257088888f8561256361159b610eea9861256994611ba2565b95611ba2565b5191611bb6565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501611c42565b6125c6915060203d81116125cc575b6125be81836103fa565b810190611bd1565b386124ad565b503d6125b4565b50949750949592509650506125f26125eb83866119a0565b9050612122565b9560005b61260084876119a0565b905081101561267c578061266061262360019361261d888b6119a0565b9061218f565b6126316101208a018a6118df565b61265a61263d8c611996565b9261265261264a8d611a00565b953690611a22565b9236916111ea565b906138ac565b61266a828b611ba2565b52612675818a611ba2565b50016125f6565b509250949390506101a083019061269382856118df565b90501580612850575b8015612847575b8015612835575b61282e5760006127596080866127b29861274a6126ea6115526126d0610926899d611996565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9761273761273e6126fa86611996565b9261271561270c6101208901896118df565b919092896118df565b949095602061272261044c565b9e8f908152019067ffffffffffffffff169052565b36916111ea565b60408a015236916111ea565b606087015282860152016121cf565b91604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f0000000000000000000000000000000000000000000000000000000000000000906004860161225e565b03925af19081156110b557600090600092612807575b50156127d15750565b610eea906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301612396565b905061282691503d806000833e61281e81836103fa565b8101906121e3565b5090386127c8565b5050505050565b50612842610b8784613d4c565b6126aa565b50823b156126a3565b5063ffffffff612862608086016121cf565b161561269c565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611964575050565b604051906128d8826103a5565b6060608083600081526000602082015282604082015282808201520152565b9061290182611a0a565b61290e60405191826103fa565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061293c8294611a0a565b0190602036910137565b90600182811c9216801561298f575b602083101461296057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612955565b600092918154916129a983612946565b80835292600181169081156129ff57506001146129c557505050565b60009081526020812093945091925b8383106129e5575060209250010190565b6001816020929493945483858701015201910191906129d4565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b9061044a612a4f9260405193848092612999565b03836103fa565b906020825491828152019160005260206000209060005b818110612a7a5750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201612a6d565b9061044a612a4f9260405193848092612a56565b9060036080604051612acb816103a5565b612b41819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c1615156020840152604051612b0c816114638160018601612999565b6040840152604051612b25816114638160028601612a56565b6060840152612b3a6040518096819301612a56565b03846103fa565b0152565b908160209103126100e75751610762816121d9565b9091612b7161076293604084526040840190612999565b9160208184039101526104ec565b60409073ffffffffffffffffffffffffffffffffffffffff610762949316815281602082015201906104ec565b9060048110156112935760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b838310612c0f57505050505090565b9091929394602080612cb4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a0612ca3612c91612c7f612c6d8887015160c08a88015260c08701906104ec565b604087015186820360408801526104ec565b606086015185820360608701526104ec565b608085015184820360808601526104ec565b9201519060a08184039101526104ec565b97019301930191939290612c00565b9160209082815201919060005b818110612cdd5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff8735612d068161123c565b168152019401929101612cd0565b90602083828152019260208260051b82010193836000925b848410612d3c5750505050505090565b909192939495602080612d82837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018852612d7c8b88611c93565b90611be6565b9801940194019294939190612d2c565b9492936107629694612f92612fa5949360808952612dbd60808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a0152606081015163ffffffff1660e08a0152608081015163ffffffff166101008a015260a081015161ffff166101208a015260c08101516101408a01526101a0612f5f612f29612ef3612ebb8d612e85612e4f60e08901516101c06101608501526102408401906104ec565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101808501526104ec565b9061012088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526104ec565b8d610140870151906101c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526104ec565b6101608501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101e08f01526104ec565b6101808401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016102008e0152612be3565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016102208b01526104ec565b9260208801528683036040880152612cc3565b926060818503910152612d14565b80612fc46040926107629594611298565b81602082015201906104ec565b359061044a8261123c565b359061044a826121d9565b9080601f830112156100e7578135612ffe81611a0a565b9261300c60405194856103fa565b81845260208085019260051b8201019283116100e757602001905b8282106130345750505090565b6020809183356130438161123c565b815201910190613027565b60c0813603126100e75761306061043b565b9061306a81612fd1565b8252613078602082016111df565b602083015261308960408201612fdc565b6040830152606081013567ffffffffffffffff81116100e7576130af9036908301611221565b6060830152608081013567ffffffffffffffff81116100e7576130d59036908301612fe7565b608083015260a08101359067ffffffffffffffff82116100e7576130fb91369101612fe7565b60a082015290565b6040516020810190600082526020815261311e6040826103fa565b51902090565b81811061312f575050565b60008155600101613124565b9190601f811161314a57505050565b61044a926000526020600020906020601f840160051c83019310613176575b601f0160051c0190613124565b9091508190613169565b919091825167ffffffffffffffff81116103a0576131a8816131a28454612946565b8461313b565b6020601f82116001146132065781906131f79394956000926131fb575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0151905038806131c5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061323984600052602060002090565b9160005b8181106132935750958360019596971061325c575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080613252565b9192602060018192868b01518155019401920161323d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81519167ffffffffffffffff83116103a0576801000000000000000083116103a0576020908254848455808510613351575b500190600052602060002060005b8381106133275750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161331a565b613368908460005285846000209182019101613124565b3861330c565b6003610762926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a060608401526133f06133bd60c0850160018401612999565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301612a56565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201612a56565b604051906134306020836103fa565b6000808352366020840137565b9190811015611b9d5760051b0190565b356107628161123c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146134845760010190565b6132ab565b8015613484577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd120820191821161348457565b9190820391821161348457565b906134fa939291614908565b91939092613507826128f7565b92613511836128f7565b94600091825b885181101561362d576000805b8a888489828510613594575b50505050501561354257600101613517565b61355261159b610e9c928b611ba2565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b61159b6135c6926125636135c18873ffffffffffffffffffffffffffffffffffffffff976115529661343d565b61344d565b9116146135d557600101613524565b6001915061360f6135ea6135c1838b8b61343d565b6135f4888c611ba2565b9073ffffffffffffffffffffffffffffffffffffffff169052565b61362261361b87613457565b968b611ba2565b52388a888489613530565b509097965094939291909460ff811690816000985b8a518a10156136ff5760005b8b878210806136f6575b156136e95773ffffffffffffffffffffffffffffffffffffffff61368a61155261159b8f6125636135c1888f8f61343d565b91161461369f5761369a90613457565b61364e565b93996136af60019294939b613489565b946136cb6136c16135c1838b8b61343d565b6135f48d8c611ba2565b6136de6136d78c613457565b9b8b611ba2565b525b01989091613642565b50509190986001906136e0565b50851515613658565b985092509395949750915081613728575050508151810361371f57509190565b80825283529190565b610e9c9291613736916134e1565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906104c3826103c1565b908160209103126100e7576040519061378f826103c1565b51815290565b6107629160e061383f61382d6137b6855161010086526101008601906104ec565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff16908601526060860151606086015261381b6080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a08701526104ec565b60c085015184820360c08601526104ec565b9201519060e08184039101526104ec565b906020610762928181520190613795565b3d1561388c573d906138728261047a565b9161388060405193846103fa565b82523d6000602084013e565b606090565b9061ffff610650602092959495604085526040850190613795565b906138b56120f8565b506138c6610c736080840151612893565b906138d7610c736060850151612893565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201529095906020818060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156110b557600091613d2d575b5073ffffffffffffffffffffffffffffffffffffffff8116958615613ce9576139e961399a8987615027565b966139a361376a565b50613a3460a0825192613a156139bf610c736020840151612893565b6040805173ffffffffffffffffffffffffffffffffffffffff909216602083015290968791820190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752866103fa565b015193613a2061045b565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff87166040870152606086015273ffffffffffffffffffffffffffffffffffffffff8916608086015260a085015260c0840152613a836104b4565b60e0840152613a9181613da3565b15613c155750602090613ad29260405193849283927f489a68f200000000000000000000000000000000000000000000000000000000845260048401613891565b03816000885af160009181613be4575b50613b265784613af0613861565b90610eea6040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401612b7f565b8490935b73ffffffffffffffffffffffffffffffffffffffff831603613b7a575b50505051613b72613b5661046b565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b613b8391615027565b908082108015613bd0575b613b985783613b47565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b50613bdb81836134e1565b83511415613b8e565b613c0791925060203d602011613c0e575b613bff81836103fa565b810190613777565b9038613ae2565b503d613bf5565b9050613c2081613dfa565b15613ca657506020613c5f91604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301613850565b03816000885af160009181613c85575b50613c7d5784613af0613861565b849093613b2a565b613c9f91925060203d602011613c0e57613bff81836103fa565b9038613c6f565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff821660045260246000fd5b613d46915060203d6020116125cc576125be81836103fa565b3861396e565b613d55816150dd565b9081613d91575b81613d65575090565b61076291507f85572ffb00000000000000000000000000000000000000000000000000000000906151d2565b9050613d9c8161516e565b1590613d5c565b613dac816150dd565b9081613de8575b81613dbc575090565b61076291507f3317103100000000000000000000000000000000000000000000000000000000906151d2565b9050613df38161516e565b1590613db3565b613e03816150dd565b9081613e3f575b81613e13575090565b61076291507faff2afbf00000000000000000000000000000000000000000000000000000000906151d2565b9050613e4a8161516e565b1590613e0a565b613e5a816150dd565b9081613e96575b81613e6a575090565b61076291507f7909b54900000000000000000000000000000000000000000000000000000000906151d2565b9050613ea18161516e565b1590613e61565b60405190613eb5826103dd565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b9015611b9d5790565b9060431015611b9d5760430190565b90821015611b9d570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906023116100e75760210190600290565b906043116100e75760230190602090565b90929192836044116100e75783116100e757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff00000000000000000000000000000000000000000000000081169260088110614037575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff000000000000000000000000000000000000000000000000000000008116926004811061409d575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110614103575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b359060208110614143575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff8211176103a057604052606060a0836000815282602082015282604082015282808201528260808201520152565b604080519091906141c383826103fa565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b8281106141fa57505050565b602090614205614170565b828285010152016141ee565b604051906142206020836103fa565b600080835282815b82811061423457505050565b60209061423f614170565b82828501015201614228565b90614254613ea8565b91604d821061488e5761429961429361426d8484613f15565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff82160361485e57506142d26142c46142be6142b88585613f39565b90614003565b60c01c90565b67ffffffffffffffff168452565b6142f66142e56142be6142b88585613f4a565b67ffffffffffffffff166020850152565b61431a6143096142be6142b88585613f5b565b67ffffffffffffffff166040850152565b61434661433961433361432d8585613f6c565b90614069565b60e01c90565b63ffffffff166060850152565b61436661435961433361432d8585613f7d565b63ffffffff166080850152565b61439061438561437f6143798585613f8e565b906140cf565b60f01c90565b61ffff1660a0850152565b6143a361439d8383613f9f565b90614135565b60c0840152816043101561482f576143ca6143c461429361426d8585613f1e565b60ff1690565b908160440183811161482f576143e4612737828685613fb0565b60e086015283811015614800576143c461429361426d614405938786613f2d565b82019160458301908482116147d15761273782604561442693018786613feb565b610100860152838110156147a25761444a6143c461429361426d6045948887613f2d565b8301019160018301908482116147735761273782604661446c93018786613feb565b61012086015283811015614744576144906143c461429361426d6001948887613f2d565b830101916001830190848211614715576127378260026144b293018786613feb565b61014086015260038301928484116146e6576144e26144db61437f614379876001968a89613feb565b61ffff1690565b01019160028301908482116146b75761273782614500928786613feb565b61016086015260048301908482116146885761437f61437983614524938887613feb565b9261ffff829416801560001461461e5750505061453f614211565b6101808501525b60028201918383116145ef578061456a6144db61437f614379876002968a89613feb565b0101918383116145c05782612737918561458394613feb565b6101a0840152036145915790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b60029192945081906146416146316141b2565b966101808a019788528887615238565b94909651966146508698611b90565b5201010114614546577fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052610e9c6024906000600452565b906001820180921161348457565b9190820180921161348457565b60ff168015613484577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9391614912613421565b93815180614c4f575b50505061492890846158fc565b92909293614978600261495761495d60036149578b67ffffffffffffffff166000526004602052604060002090565b01612aa6565b9867ffffffffffffffff166000526004602052604060002090565b6149a361499e61499661498e87518651906148cd565b8a51906148cd565b8351906148cd565b6128f7565b9687956000805b87518210156149e257906149da6001926135f46149ca61159b858d611ba2565b916149d481613457565b9c611ba2565b0189976149aa565b9750509193969092945060005b8551811015614a245780614a1e614a0b61159b6001948a611ba2565b6135f4614a178b613457565b9a8d611ba2565b016149ef565b50919350919460005b8451811015614a625780614a5c614a4961159b60019489611ba2565b6135f4614a558a613457565b998c611ba2565b01614a2d565b5093919592509360005b828110614be6575b5050600090815b818110614b475750508152835160005b818110614a9a57508452929190565b926000959495915b8351831015614b3857614ab861159b8688611ba2565b73ffffffffffffffffffffffffffffffffffffffff614add61155261159b8789611ba2565b911603614b2757614aed90613489565b90614b08614afe61159b8489611ba2565b6135f48789611ba2565b60ff8716614b17575b90614aa2565b95614b21906148da565b95614b11565b9091614b3290613457565b91614b11565b91509260019095949501614a8b565b614b5461159b8286611ba2565b73ffffffffffffffffffffffffffffffffffffffff81168015614bdc57600090815b868110614bb0575b5050906001929115614b93575b505b01614a7b565b614baa906135f4614ba387613457565b9688611ba2565b38614b8b565b81614bc161155261159b848c611ba2565b14614bce57600101614b76565b506001915081905038614b7e565b5050600190614b8d565b614bf661155261159b8387611ba2565b15614c0357600101614a6c565b5091929360009591955b8351811015614c425780614c3c614c2961159b60019488611ba2565b6135f4614c358b613457565b9a89611ba2565b01614c0d565b5093929150933880614a74565b909192945060018103614d01575060146060614c6a84611b90565b5101515103614cbe5781614cb591614c94610c736060614c8c61492897611b90565b510151612893565b918760a0614cac614ca484611b90565b515193611b90565b510151936156fa565b9290388061491b565b610eea6060614ccc84611b90565b5101516040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301612396565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b614d36613ea8565b5090565b929061311e9173ffffffffffffffffffffffffffffffffffffffff614d8867ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906104ec565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103fa565b9060405191614dcb60c0846103fa565b60848352602083019060a03683375a612ee0811115614e1c57600091614df183926134b4565b82602083519301913090f1903d9060848211614e13575b6000908286523e9190565b60849150614e08565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff600154163303614e9857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614ed08151846148cd565b928315614ffd5760005b848110614ee8575050505050565b81811015614fe257614efd61159b8286611ba2565b73ffffffffffffffffffffffffffffffffffffffff811680156115c257614f23836148bf565b878110614f3557505050600101614eda565b84811015614fb25773ffffffffffffffffffffffffffffffffffffffff614f5f61159b838a611ba2565b168214614f6e57600101614f23565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614fdd61159b614fd788856134e1565b89611ba2565b614f5f565b614ff861159b614ff284846134e1565b85611ba2565b614efd565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa600091816150a6575b506150a25782613af0613861565b9150565b90916020823d6020116150d5575b816150c1602093836103fa565b810103126150d25750519038615094565b80fd5b3d91506150b4565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a7000000000000000000000000000000000000000000000000000000006024820152602481526151416044826103fa565b5191617530fa6000513d82615162575b508161515b575090565b9050151590565b60201115915038615151565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff000000000000000000000000000000000000000000000000000000006024820152602481526151416044826103fa565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a7000000000000000000000000000000000000000000000000000000008452166024820152602481526151416044826103fa565b9291615242614170565b91808210156155e55761525c61429361426d848489613f2d565b600160ff82160361485e5750602182018181116155b65761528561439d8260018601858a613feb565b845281811015615587576143c461429361426d6152a393858a613f2d565b8201916022830190828211615558576127378260226152c49301858a613feb565b602085015281811015615529576152e76143c461429361426d602294868b613f2d565b8301019160018301908282116154fa576127378260236153099301858a613feb565b6040850152818110156154cb5761532c6143c461429361426d600194868b613f2d565b83010191600183019082821161549c5761273782600261534e9301858a613feb565b60608501528181101561546d576153716143c461429361426d600194868b613f2d565b830101600181019282841161543e576127378460026153929301858a613feb565b6080850152600381019282841161540f576002916153bd6144db61437f61437988600196898e613feb565b010101948186116153e0576153d792612737928792613feb565b60a08201529190565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9080601f830112156100e757815161562b81611a0a565b9261563960405194856103fa565b81845260208085019260051b8201019283116100e757602001905b8282106156615750505090565b6020809183516156708161123c565b815201910190615654565b906020828203126100e757815167ffffffffffffffff81116100e7576107629201615614565b95949060019460a09467ffffffffffffffff6156f59573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906104ec565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095929493929091906020838060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9283156110b55760009361587f575b506157a283613da3565b6157dc575b50505050508151156157b7575090565b610762915061495760029167ffffffffffffffff166000526004602052604060002090565b6000949650859161583273ffffffffffffffffffffffffffffffffffffffff92604051998a97889687957f89720a62000000000000000000000000000000000000000000000000000000008752600487016156a1565b0392165afa9182156110b55760009261585a575b5061585082615ad3565b38808080806157a7565b6158789192503d806000833e61587081836103fa565b81019061567b565b9038615846565b61589991935060203d6020116125cc576125be81836103fa565b9138615798565b90916060828403126100e757815167ffffffffffffffff81116100e757836158c9918401615614565b92602083015167ffffffffffffffff81116100e7576040916158ec918501615614565b92015160ff811681036100e75790565b9190916159206109408267ffffffffffffffff166000526004602052604060002090565b90833b61593d575b50606001519150615937613421565b90600090565b61594684613e51565b15615928576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa80156110b5576000809481926159f2575b506159c181615ad3565b6159ca85615ad3565b8051158015906159e6575b6159df5750615928565b9392909150565b5060ff821615156159d5565b9150615a119294503d8091833e615a0981836103fa565b8101906158a0565b909391386159b7565b600254811015611b9d5760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b80600052600360205260406000205415600014615acd57600254680100000000000000008110156103a057600181016002556000600254821015611b9d57600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b50600090565b80519060005b828110615ae557505050565b60018101808211613484575b838110615b015750600101615ad9565b73ffffffffffffffffffffffffffffffffffffffff615b208385611ba2565b5116615b3261155261159b8487611ba2565b14615b3f57600101615af1565b610e9c615b4f61159b8486611ba2565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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

	Execute(opts *bind.TransactOpts, encodedMessage []byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message MessageV1CodecMessageV1, messageId [32]byte, ccvs []common.Address, verifierResults [][]byte) (*types.Transaction, error)

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
