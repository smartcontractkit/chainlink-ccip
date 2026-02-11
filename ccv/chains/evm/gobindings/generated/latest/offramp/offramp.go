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
	LocalChainSelector        uint64
	GasForCallExactCheck      uint16
	RmnRemote                 common.Address
	TokenAdminRegistry        common.Address
	MaxGasBufferToUpdateState uint32
}

var OffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct MessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"executionGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"ccipReceiveGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"ccvAndExecutorHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"struct MessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenReceiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"verifierResults\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OffRamp.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCVsForMessage\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"optionalCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"threshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum Internal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum Internal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxGasBufferToUpdateStateUpdated\",\"inputs\":[{\"name\":\"oldMaxGasBufferToUpdateState\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newMaxGasBufferToUpdateState\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.SourceChainConfigArgs\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamps\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OffRamp.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGasBufferToUpdateState\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"GasCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFound\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"verifierResults\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidGasLimitOverride\",\"inputs\":[{\"name\":\"messageGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"gasLimitOverride\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOffRamp\",\"inputs\":[{\"name\":\"expected\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidOnRamp\",\"inputs\":[{\"name\":\"got\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierResultsLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120604052346102af57604051601f61622938819003918201601f19168301916001600160401b038311848410176102b45780849260a0946040528339810103126102af576040519060009060a083016001600160401b0381118482101761029b5760405280516001600160401b0381168103610297578352602081015161ffff8116810361029757602084019081526040820151906001600160a01b038216820361029357604085019182526060830151926001600160a01b038416840361028f5760608601938452608001519363ffffffff8516850361028c5760808601948552331561027d57600180546001600160a01b0319163317905582516001600160a01b031615801561026b575b61025c5785516001600160401b03161561024d5761ffff8251161561023e5763ffffffff8551161561023e5785516001600160401b03908116608090815284516001600160a01b0390811660a09081528751821660c052855161ffff90811660e052895163ffffffff90811661010052604080518d519097168752885190921660208701528851841691860191909152885190921660608501528851909116918301919091527f6db4162777b6c980e778bb05a3d9e050f3792b091287ff0d4f3d51bdcd7427db91a1604051615f5e90816102cb823960805181818161015c0152611aad015260a0518181816101bf01526119d4015260c0518181816101fb01528181614d7301526158f8015260e05181818161018301526152a701526101005181818161022701526145a90152f35b632855a4d960e11b8152600490fd5b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b03161561010e565b639b15e16f60e01b8152600490fd5b80fd5b8480fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780633b81ab9b146100c857806349d8033e146100c35780635215505b146100be5780635643a782146100b957806361a10e59146100b457806379ba5097146100af5780638da5cb5b146100aa578063e9d68a8e146100a55763f2fde38b146100a057600080fd5b61123e565b61105c565b610fdc565b610ef3565b610e24565b610a49565b610914565b610757565b610660565b610557565b61042f565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000608060405161012b816102fa565b82815282602082015282604082015282606082015201526102c7604051610151816102fa565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015263ffffffff7f000000000000000000000000000000000000000000000000000000000000000016608082015260405191829182919091608063ffffffff8160a084019567ffffffffffffffff815116855261ffff602082015116602086015273ffffffffffffffffffffffffffffffffffffffff604082015116604086015273ffffffffffffffffffffffffffffffffffffffff6060820151166060860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60a0810190811067ffffffffffffffff82111761031657604052565b6102cb565b6101c0810190811067ffffffffffffffff82111761031657604052565b6020810190811067ffffffffffffffff82111761031657604052565b90601f601f19910116810190811067ffffffffffffffff82111761031657604052565b6040519061038660c083610354565b565b6040519061038660a083610354565b6040519061038661010083610354565b60405190610386604083610354565b67ffffffffffffffff811161031657601f01601f191660200190565b604051906103e1602083610354565b60008252565b60005b8381106103fa5750506000910152565b81810151838201526020016103ea565b90601f19601f602093610428815180928187528780880191016103e7565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576102c760408051906104708183610354565b601182527f4f666652616d7020322e302e302d64657600000000000000000000000000000060208301525191829160208352602083019061040a565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104f85750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104eb565b9161055060ff916105426040949796976060875260608701906104da565b9085820360208701526104da565b9416910152565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106076105b56105af6102c79336906004016104ac565b90613922565b6105c3610140820151611332565b60601c9067ffffffffffffffff81511691610180820151906106018161ffff60a0860151169463ffffffff60806101a0830151519201511690613fc4565b93614142565b60409391935193849384610524565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b63ffffffff8116036100e757565b359061038682610647565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576106af9036906004016104ac565b9060243567ffffffffffffffff81116100e7576106d0903690600401610616565b926044359367ffffffffffffffff85116100e7576106f561070a953690600401610616565b9390926064359561070587610647565b6118b0565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561074557565b61070c565b9060048210156107455752565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576004356000526006602052602060ff604060002054166107a9604051809261074a565bf35b9080602083519182815201916020808360051b8301019401926000915b8383106107d757505050505090565b90919293946020806107f583601f198660019603018752895161040a565b970193019301919392906107c8565b61086f9173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061085e61084c604085015160a0604086015260a08501906107ab565b606085015184820360608601526104da565b9201519060808184039101526104da565b90565b6040810160408252825180915260206060830193019060005b8181106108f4575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106108c757505050505090565b90919293946020806108e583601f1986600196030187528951610804565b970193019301919392906108b8565b825167ffffffffffffffff1685526020948501949092019160010161088b565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461094f81612010565b9061095d6040519283610354565b808252601f1961096c82612010565b0160005b818110610a3257505061098281612054565b9060005b81811061099e5750506102c760405192839283610872565b806109d66109bd6109b0600194615bab565b67ffffffffffffffff1690565b6109c783876120c6565b9067ffffffffffffffff169052565b610a16610a116109f76109e984886120c6565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b61218e565b610a2082876120c6565b52610a2b81866120c6565b5001610986565b602090610a3d612028565b82828701015201610970565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e757610a98903690600401610616565b90610aa1614689565b6000905b828210610aae57005b610ac1610abc838584612333565b612461565b6020810191610adb6109b0845167ffffffffffffffff1690565b15610dfa57610b1d610b04610b04845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610ded575b610b745760808201949060005b86518051821015610b9e57610b04610b4d83610b67936120c6565b5173ffffffffffffffffffffffffffffffffffffffff1690565b15610b7457600101610b32565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050919394929060009260a08601935b84518051821015610bd657610b04610b4d83610bc9936120c6565b15610b7457600101610bae565b505095929491909394610bec86518251906146d4565b610c016109f7835167ffffffffffffffff1690565b90610c31610c17845167ffffffffffffffff1690565b67ffffffffffffffff166000526005602052604060002090565b95610c3b87615be0565b606085019860005b8a518051821015610c935790610c5b816020936120c6565b518051928391012091158015610c83575b610b7457610c7c6001928b615caf565b5001610c43565b50610c8c612516565b8214610c6c565b5050976001975093610daa610de3946003610dd695610da27f72ec11bb832a18492cf3aafef578325a1e9fc7105b5ba447ca94596fec79393e996109b0979f610cf060408e610ce9610d36945160018b016126cf565b0151151590565b86547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178655565b610d98610d578d5173ffffffffffffffffffffffffffffffffffffffff1690565b869073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b51600285016127b4565b5191016127b4565b610dc7610dc26109b0835167ffffffffffffffff1690565b615c1e565b505167ffffffffffffffff1690565b9260405191829182612848565b0390a20190610aa5565b5060808201515115610b25565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e7576024359060443567ffffffffffffffff81116100e757610eb6903690600401610616565b926064359367ffffffffffffffff85116100e757610edb61070a953690600401610616565b93909260843595610eeb87610647565b600401613013565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303610fb2577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b67ffffffffffffffff8116036100e757565b35906103868261102e565b90602061086f928181520190610804565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff6004356110a08161102e565b6110a8612028565b5016600052600460205260406000206040516110c3816102fa565b60ff825473ffffffffffffffffffffffffffffffffffffffff8116835260a01c1615156020820152600182018054906110fb82612010565b916111096040519384610354565b80835260208301916000526020600020916000905b82821061115d576102c78661114c60038a8960408501526111416002820161212d565b60608501520161212d565b60808201526040519182918261104b565b6040516000855461116d816120da565b80845290600181169081156111df57506001146111a7575b506001928261119985946020940382610354565b81520194019101909261111e565b6000878152602081209092505b8183106111c957505081016020016001611185565b60018160209254838688010152019201916111b4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050611185565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561128e81611220565b611296614689565b1633811461130857807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116926014811061136a575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b801515036100e757565b908160209103126100e7575161086f8161139c565b6040513d6000823e3d90fd5b90602061086f92818152019061040a565b60409073ffffffffffffffffffffffffffffffffffffffff61086f9493168152816020820152019061040a565b929192611411826103b6565b9161141f6040519384610354565b8294818452818301116100e7578281602093846000960137010152565b9060048110156107455760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061149f57505050505090565b909192939460208061152683601f1986600196030187528951908151815260a06115156115036114f16114df8887015160c08a88015260c087019061040a565b6040870151868203604088015261040a565b6060860151858203606087015261040a565b6080850151848203608086015261040a565b9201519060a081840391015261040a565b97019301930191939290611490565b9160209082815201919060005b81811061154f5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561157881611220565b168152019401929101611542565b601f8260209493601f19938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b84841061161f5750505050505090565b90919293949560208061164783601f1986600196030188526116418b886115a7565b90611586565b980194019401929493919061160f565b9796949161187790608095611885956118646103869a956116e98e60a0815261168d60a08201845167ffffffffffffffff169052565b602083015167ffffffffffffffff1660c0820152604083015167ffffffffffffffff1660e0820152606083015163ffffffff16610100820152828c015163ffffffff1661012082015261014060a084015191019061ffff169052565b8d61016060c08301519101528d6101a06118306117fa6117c461178e61175861172460e08901516101c06101808a015261026089019061040a565b6101008901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6089830301888a015261040a565b6101208801517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60888303016101c089015261040a565b6101408701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878303016101e088015261040a565b6101608601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608683030161020087015261040a565b6101808501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6085830301610220860152611473565b920151906102407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608285030191015261040a565b9260208d01528b830360408d0152611535565b9188830360608a01526115f7565b94019063ffffffff169052565b806118a360409261086f959461074a565b816020820152019061040a565b95919293959490946118c860015460ff9060a01c1690565b611fe657611910740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b61191a8682613922565b956119bb60206119606119386109b08b5167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611fe157600091611fb2575b50611f6757611a33611a2f611a256109f78a5167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b611f1c57611a4c610c17885167ffffffffffffffff1690565b611a76611a2f60e08a01928351602081519101209060019160005201602052604060002054151590565b611ee557506101008701516014815114801590611eb4575b611e7d5750602087015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603611e465750828603611e12576101408701516014815103611dd8575063ffffffff84168015159081611db2575b50611d655790611b17913691611405565b6020815191012095611b3d611b36886000526006602052604060002090565b5460ff1690565b611b468161073b565b8015908115611d51575b5015611cdf57611beb92611bf0959492611bdd92611ba6611b7b8b6000526006602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957f61a10e590000000000000000000000000000000000000000000000000000000060208801528b8b60248901611657565b03601f198101835282610354565b614582565b9015611cad577f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff6002925b611c4184611c3c886000526006602052604060002090565b61143c565b611c7d611c6b6040611c5b885167ffffffffffffffff1690565b97015167ffffffffffffffff1690565b91836040519485941697169583611892565b0390a46103867fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df267ffffffffffffffff600392611c24565b611d4d8787611d0b6040611cfb835167ffffffffffffffff1690565b92015167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150611d5e8161073b565b1438611b50565b611d4d84611d7a60808a015163ffffffff1690565b7fdf2964df0000000000000000000000000000000000000000000000000000000060005263ffffffff90811660045216602452604490565b9050611dd1611dc860808a015163ffffffff1690565b63ffffffff1690565b1138611b06565b611e0e906040519182917f8d666f60000000000000000000000000000000000000000000000000000000008352600483016113c7565b0390fd5b7f88f80aa2000000000000000000000000000000000000000000000000000000006000526004869052602483905260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b611e0e906040519182917f55216e3100000000000000000000000000000000000000000000000000000000835230600484016113d8565b50611ec7611ec182611332565b60601c90565b73ffffffffffffffffffffffffffffffffffffffff16301415611a8e565b611e0e90516040519182917fa50bd147000000000000000000000000000000000000000000000000000000008352600483016113c7565b611d4d611f31885167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611d4d611f7c885167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611fd4915060203d602011611fda575b611fcc8183610354565b8101906113a6565b38611a05565b503d611fc2565b6113bb565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff81116103165760051b60200190565b60405190612035826102fa565b6060608083600081526000602082015282604082015282808201520152565b9061205e82612010565b61206b6040519182610354565b828152601f1961207b8294612010565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156120c15760200190565b612085565b80518210156120c15760209160051b010190565b90600182811c92168015612123575b60208310146120f457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916120e9565b906040519182815491828252602082019060005260206000209260005b81811061215f57505061038692500383610354565b845473ffffffffffffffffffffffffffffffffffffffff1683526001948501948794506020909301920161214a565b9060405161219b816102fa565b809260ff815473ffffffffffffffffffffffffffffffffffffffff8116845260a01c1615156020830152600181018054906121d582612010565b916121e36040519384610354565b80835260208301916000526020600020916000905b82821061222d5750505050600360809261222892604086015261221d6002820161212d565b60608601520161212d565b910152565b6040516000855461223d816120da565b80845290600181169081156122af5750600114612277575b506001928261226985946020940382610354565b8152019401910190926121f8565b6000878152602081209092505b81831061229957505081016020016001612255565b6001816020925483868801015201920191612284565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208581019190915291151560051b8401909101915060019050612255565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b90156120c1578061086f916122f0565b908210156120c15761086f9160051b8101906122f0565b359061038682611220565b35906103868261139c565b9080601f830112156100e75781602061086f93359101611405565b9080601f830112156100e757813561239281612010565b926123a06040519485610354565b81845260208085019260051b820101918383116100e75760208201905b8382106123cc57505050505090565b813567ffffffffffffffff81116100e7576020916123ef87848094880101612360565b8152019101906123bd565b9080601f830112156100e757813561241181612010565b9261241f6040519485610354565b81845260208085019260051b8201019283116100e757602001905b8282106124475750505090565b60208091833561245681611220565b81520191019061243a565b60c0813603126100e757612473610377565b9061247d8161234a565b825261248b60208201611040565b602083015261249c60408201612355565b6040830152606081013567ffffffffffffffff81116100e7576124c2903690830161237b565b6060830152608081013567ffffffffffffffff81116100e7576124e890369083016123fa565b608083015260a08101359067ffffffffffffffff82116100e75761250e913691016123fa565b60a082015290565b60405160208101906000825260208152612531604082610354565b51902090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110612571575050565b60008155600101612566565b9190601f811161258c57505050565b610386926000526020600020906020601f840160051c830193106125b8575b601f0160051c0190612566565b90915081906125ab565b919091825167ffffffffffffffff8111610316576125ea816125e484546120da565b8461257d565b6020601f821160011461264857819061263993949560009261263d575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612607565b601f1982169061265d84600052602060002090565b9160005b8181106126b757509583600195969710612680575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080612676565b9192602060018192868b015181550194019201612661565b815191680100000000000000008311610316578154838355808410612731575b506020612703910191600052602060002090565b6000915b8383106127145750505050565b6001602082612725839451866125c2565b01920192019190612707565b8260005283602060002091820191015b81811061274e57506126ef565b8061275b600192546120da565b80612768575b5001612741565b601f8111831461277e5750600081555b38612761565b6127a29083601f61279485600052602060002090565b920160051c82019101612566565b60008181526020812081835555612778565b81519167ffffffffffffffff83116103165768010000000000000000831161031657602090825484845580851061282b575b500190600052602060002060005b8381106128015750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016127f4565b612842908460005285846000209182019101612566565b386127e6565b9061086f916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff602083015116604082015260408201511515606082015260a06128c16128ac606085015160c0608086015260e08501906107ab565b6080850151601f1985830301848601526104da565b9201519060c0601f19828503019101526104da565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b91602061086f938181520191611586565b919091357fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008116926014811061136a575050565b3561086f81610647565b3561086f8161102e565b61ffff8116036100e757565b3561086f816129d4565b91909160c0818403126100e7576129ff610377565b9281358452602082013567ffffffffffffffff81116100e75781612a24918401612360565b6020850152604082013567ffffffffffffffff81116100e75781612a49918401612360565b6040850152606082013567ffffffffffffffff81116100e75781612a6e918401612360565b6060850152608082013567ffffffffffffffff81116100e75781612a93918401612360565b608085015260a082013567ffffffffffffffff81116100e757612ab69201612360565b60a0830152565b929190612ac981612010565b93612ad76040519586610354565b602085838152019160051b8101918383116100e75781905b838210612afd575050505050565b813567ffffffffffffffff81116100e757602091612b1e87849387016129ea565b815201910190612aef565b908210156120c157612b409160051b81019061292a565b9091565b908160209103126100e7575161086f81611220565b60409073ffffffffffffffffffffffffffffffffffffffff61086f95931681528160208201520191611586565b3590610386826129d4565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b838310612c0c5750505050505090565b909192939495601f1982820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41843603018112156100e7576020612cef600193868394019081358152612ce1612cd6612cbb612ca0612c85612c75898801886115a7565b60c08b89015260c0880191611586565b612c9260408801886115a7565b908783036040890152611586565b612cad60608701876115a7565b908683036060880152611586565b612cc860808601866115a7565b908583036080870152611586565b9260a08101906115a7565b9160a0818503910152611586565b980196019493019190612bfc565b612f7561086f9593949260608352612d2960608401612d1b83611040565b67ffffffffffffffff169052565b612d49612d3860208301611040565b67ffffffffffffffff166080850152565b612d69612d5860408301611040565b67ffffffffffffffff1660a0850152565b612d85612d7860608301610655565b63ffffffff1660c0850152565b612da1612d9460808301610655565b63ffffffff1660e0850152565b612dbc612db060a08301612b86565b61ffff16610100850152565b60c0810135610120840152612f44612f38612ef9612eba612e7b612e3c612dfd612de960e08901896115a7565b6101c06101408d01526102208c0191611586565b612e0b6101008901896115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08c8403016101608d0152611586565b612e4a6101208801886115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08b8403016101808c0152611586565b612e896101408701876115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8403016101a08b0152611586565b612ec86101608601866115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0898403016101c08a0152611586565b612f07610180850185612b91565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888403016101e0890152612be4565b916101a08101906115a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa085840301610200860152611586565b9360208201526040818503910152611586565b604051906040820182811067ffffffffffffffff8211176103165760405260006020838281520152565b90612fbc82612010565b612fc96040519182610354565b828152601f19612fd98294612010565b019060005b828110612fea57505050565b602090612ff5612f88565b82828501015201612fde565b9190820391821161300e57565b612537565b9391949092943033036135735760006130306101808701876128d6565b9050613494575b6130b9613054611ec161304e6101408a018a61292a565b9061298c565b9261307e846130676101a08b018b61292a565b9050613078611dc860808d016129c0565b90613fc4565b98899160a061308c8b6129ca565b876130b38d6130ab6130a26101808301836128d6565b969092016129e0565b943691612abd565b9161496f565b92909860005b8a518110156132a9578060206130e2610b04610b048f61312e96610b4d916120c6565b6130f76130ef848a6120c6565b518a8c612b29565b91906040518096819482937fc3a7ded60000000000000000000000000000000000000000000000000000000084526004840161297b565b03915afa918215611fe157600092613279575b5073ffffffffffffffffffffffffffffffffffffffff82161561321c5761317361316b82886120c6565b51888a612b29565b929073ffffffffffffffffffffffffffffffffffffffff82163b156100e7576000918b8373ffffffffffffffffffffffffffffffffffffffff8f6131e6604051998a97889687947fbff0ec1d00000000000000000000000000000000000000000000000000000000865260048601612cfd565b0393165af1918215611fe157600192613201575b50016130bf565b80613210600061321693610354565b806100dc565b386131fa565b8561324489898f85613237610b4d611e0e9861323d946120c6565b956120c6565b5191612b29565b6040939193519384937f2665cea200000000000000000000000000000000000000000000000000000000855260048501612b59565b61329b91925060203d81116132a2575b6132938183610354565b810190612b44565b9038613141565b503d613289565b50959793509593509650506132cc6132c56101808401846128d6565b9050612fb2565b956132db6101808401846128d6565b9050613382575b5061337b576103869461334b6132f7836129ca565b61333861333f61330b61012087018761292a565b6133196101a089018961292a565b949095613324610388565b9c8d5267ffffffffffffffff1660208d0152565b3691611405565b60408901523691611405565b6060860152608085015263ffffffff82161561336857509161520f565b61337591506080016129c0565b9161520f565b5050505050565b6133dc61339c6133966101808601866128d6565b90612323565b6133aa61012086018661292a565b6133d66133b6886129ca565b926133ce6133c660a08b016129e0565b9536906129ea565b923691611405565b90614cde565b91906133e7896120b4565b5273ffffffffffffffffffffffffffffffffffffffff613421611ec161304e6134176133966101808a018a6128d6565b608081019061292a565b921673ffffffffffffffffffffffffffffffffffffffff831603613446575b506132e2565b61347a61347f926134746134598b6120b4565b515173ffffffffffffffffffffffffffffffffffffffff1690565b90614869565b613001565b602061348a886120b4565b5101523880613440565b5060146134b56134ab6133966101808901896128d6565b606081019061292a565b90500361355f5760146134d26134176133966101808901896128d6565b905003613515576135106134f6611ec161304e6134176133966101808b018b6128d6565b613474611ec161304e6134ab6133966101808c018c6128d6565b613037565b6135296134176133966101808801886128d6565b90611e0e6040519283927f8d666f600000000000000000000000000000000000000000000000000000000084526004840161297b565b6135296134ab6133966101808801886128d6565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906135aa8261031b565b60606101a08360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c08201528260e082015282610100820152826101208201528261014082015282610160820152826101808201520152565b90156120c15790565b90604310156120c15760430190565b908210156120c1570190565b906009116100e75760010190600890565b906011116100e75760090190600890565b906019116100e75760110190600890565b90601d116100e75760190190600490565b906021116100e757601d0190600490565b906023116100e75760210190600290565b906043116100e75760230190602090565b90929192836044116100e75783116100e757604401917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbc0190565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff0000000000000000000000000000000000000000000000008116926008811061372c575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613792575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b919091357fffff000000000000000000000000000000000000000000000000000000000000811692600281106137f8575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b359060208110613838575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b6040519060c0820182811067ffffffffffffffff82111761031657604052606060a0836000815282602082015282604082015282808201528260808201520152565b604080519091906138b88382610354565b6001815291601f19018260005b8281106138d157505050565b6020906138dc613865565b828285010152016138c5565b604051906138f7602083610354565b600080835282815b82811061390b57505050565b602090613916613865565b828285010152016138ff565b9061392b61359d565b91604d8210613f935761397061396a613944848461360a565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613f6357506139a961399b61399561398f858561362e565b906136f8565b60c01c90565b67ffffffffffffffff168452565b6139cd6139bc61399561398f858561363f565b67ffffffffffffffff166020850152565b6139f16139e061399561398f8585613650565b67ffffffffffffffff166040850152565b613a1d613a10613a0a613a048585613661565b9061375e565b60e01c90565b63ffffffff166060850152565b613a3d613a30613a0a613a048585613672565b63ffffffff166080850152565b613a67613a5c613a56613a508585613683565b906137c4565b60f01c90565b61ffff1660a0850152565b613a7a613a748383613694565b9061382a565b60c08401528160431015613f3457613aa1613a9b61396a6139448585613613565b60ff1690565b9081604401838111613f0557613abb6133388286856136a5565b60e086015283811015613ed657613a9b61396a613944613adc938786613622565b8201916045830190848211613ea857613338826045613afd930187866136e0565b61010086015283811015613e7957613b21613a9b61396a6139446045948887613622565b830101916001830190848211613e4a57613338826046613b43930187866136e0565b61012086015283811015613e1b57613b67613a9b61396a6139446001948887613622565b830101916001830190848211613dec57613338826002613b89930187866136e0565b6101408601526003830192848411613dbd57613bb9613bb2613a56613a50876001968a896136e0565b61ffff1690565b0101916002830190848211613d8e5761333882613bd79287866136e0565b6101608601526004830190848211613d5f57613a56613a5083613bfb9388876136e0565b9261ffff8294168015600014613cf557505050613c166138e8565b6101808501525b6002820191838311613cc65780613c41613bb2613a56613a50876002968a896136e0565b010191838311613c9757826133389185613c5a946136e0565b6101a084015203613c685790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b6002919294508190613d18613d086138a7565b966101808a01978852888761534c565b9490965196613d2786986120b4565b5201010114613c1d577fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600860045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000006000526004805260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052611d4d6024906000600452565b15919082614042575b508115614038575b8115613fdf575090565b905061400b7f85572ffb0000000000000000000000000000000000000000000000000000000082615e31565b9081614026575b8161401c57501590565b611a2f9150615dd1565b905061403181615d0b565b1590614012565b803b159150613fd5565b15915038613fcd565b6040519061405a602083610354565b6000808352366020840137565b604080519091906140788382610354565b6001815291601f1901366020840137565b906001820180921161300e57565b9190820180921161300e57565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461300e5760010190565b80548210156120c15760005260206000200190600090565b801561300e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff16801561300e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9493909160609260009261415461404b565b95835180614503575b5050156144e9575051156144d857505061417561404b565b61417d61404b565b936000925b60026141c360036141a78567ffffffffffffffff166000526004602052604060002090565b019367ffffffffffffffff166000526004602052604060002090565b018254958154956141ec6141e7886141e28b6141e28b518a5190614097565b614097565b612054565b9889976000805b8951821015614246579061423e600192614223614213610b4d858f6120c6565b9161421d816140a4565b9e6120c6565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018b996141f3565b919395979a9294969850506000905b885182101561428d5790614285600192614223614275610b4d858e6120c6565b9161427f816140a4565b9d6120c6565b018a98614255565b959893969991929497505060005b8281106144a5575050505060005b828110614421575b50509091929350600090815b81811061437f5750508452805160005b85518110156143795760005b8281106142ea575b506001016142cd565b6142f7610b4d82866120c6565b73ffffffffffffffffffffffffffffffffffffffff61431c610b04610b4d868c6120c6565b9116146143315761432c906140a4565b6142d9565b9161433e614356916140e9565b9261422361434f610b4d86886120c6565b91866120c6565b60ff8416614365575b386142e1565b92614371600191614114565b93905061435f565b50815291565b61438c610b4d82896120c6565b73ffffffffffffffffffffffffffffffffffffffff8116801561441757600090815b8a8782106143ea575b5050509060019291156143cd575b505b016142bd565b6143e4906142236143dd876140a4565b968b6120c6565b386143c5565b6143fb610b04610b4d8486946120c6565b14614408576001016143ae565b5060019150819050388a6143b7565b50506001906143c7565b614431610b04610b4d838b6120c6565b1561443e576001016142a9565b50909192939460005b828110614459578695949392506142b1565b8061449f61448c61446c600194866140d1565b905473ffffffffffffffffffffffffffffffffffffffff9160031b1c1690565b614223614498886140a4565b978c6120c6565b01614447565b6144ce829394966142236144be61446c856001976140d1565b916144c8816140a4565b996120c6565b019089929161429b565b9190936144e3614067565b91614182565b9150506144f99150849294615a6f565b9094909290614182565b9091965060018103614555575061454d9061452c611ec187614524876120b4565b510151611332565b90614536856120b4565b51518a60a0614544886120b4565b5101519361588e565b94388061415d565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b906040519161459260c084610354565b60848352602083019060a03683375a9063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016908183111561460b576145e0600093928493613001565b82602083519301913090f1903d9060848211614602575b6000908286523e9190565b608491506145f7565b611d4d7fffffffff0000000000000000000000000000000000000000000000000000000063ffffffff5a1660e01b167f2882569d00000000000000000000000000000000000000000000000000000000600052907fffffffff0000000000000000000000000000000000000000000000000000000060249216600452565b73ffffffffffffffffffffffffffffffffffffffff6001541633036146aa57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916146e2815184614097565b92831561480f5760005b8481106146fa575050505050565b818110156147f45761470f610b4d82866120c6565b73ffffffffffffffffffffffffffffffffffffffff81168015610b745761473583614089565b878110614747575050506001016146ec565b848110156147c45773ffffffffffffffffffffffffffffffffffffffff614771610b4d838a6120c6565b16821461478057600101614735565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6147ef610b4d6147e98885613001565b896120c6565b614771565b61480a610b4d6148048484613001565b856120c6565b61470f565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b3d15614864573d9061484a826103b6565b916148586040519384610354565b82523d6000602084013e565b606090565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161491e575b5061491a57826148e4614839565b90611e0e6040519283927f9fe2f95a000000000000000000000000000000000000000000000000000000008452600484016113d8565b9150565b90916020823d60201161494d575b8161493960209383610354565b8101031261494a57505190386148d6565b80fd5b3d915061492c565b91908110156120c15760051b0190565b3561086f81611220565b9061497e949596939291614142565b9193909261498b82612054565b9261499583612054565b94600091825b8851811015614a8f576000805b8a888489828510614a18575b5050505050156149c65760010161499b565b6149d6610b4d611d4d928b6120c6565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b610b4d614a4a92613237614a458873ffffffffffffffffffffffffffffffffffffffff97610b0496614955565b614965565b911614614a59576001016149a8565b60019150614a78614a6e614a45838b8b614955565b614223888c6120c6565b614a846143dd876140a4565b52388a8884896149b4565b509097965094939291909460ff811690816000985b8a518a1015614b615760005b8b87821080614b58575b15614b4b5773ffffffffffffffffffffffffffffffffffffffff614aec610b04610b4d8f613237614a45888f8f614955565b911614614b0157614afc906140a4565b614ab0565b9399614b1160019294939b6140e9565b94614b2d614b23614a45838b8b614955565b6142238d8c6120c6565b614b40614b398c6140a4565b9b8b6120c6565b525b01989091614aa4565b5050919098600190614b42565b50851515614aba565b985092509395949750915081614b8a5750505081518103614b8157509190565b80825283529190565b611d4d9291614b9891613001565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906103e182610338565b908160209103126100e75760405190614bf182610338565b51815290565b61086f9160e0614ca1614c8f614c188551610100865261010086019061040a565b60208681015167ffffffffffffffff169086015260408681015173ffffffffffffffffffffffffffffffffffffffff169086015260608601516060860152614c7d6080870151608087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a086015185820360a087015261040a565b60c085015184820360c086015261040a565b9201519060e081840391015261040a565b90602061086f928181520190614bf7565b9061ffff610550602092959495604085526040850190614bf7565b92939193614cea612f88565b50614cfb611ec16080860151611332565b91614d0c611ec16060870151611332565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909690956020878060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa968715611fe15760009761503b575b5073ffffffffffffffffffffffffffffffffffffffff8716948515614ff757614dcb614bcc565b50614e19825191614dfc60a0602086015195015195614de8610397565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff8816608084015260a083015260c0820152614e4c6103d2565b60e0820152614e5a85615728565b15614f205790614e9e9260209260006040518096819582947f489a68f200000000000000000000000000000000000000000000000000000000845260048401614cc3565b03925af160009181614eef575b50614eb957836148e4614839565b929091925b51614ee6614eca6103a7565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60208201529190565b614f1291925060203d602011614f19575b614f0a8183610354565b810190614bd9565b9038614eab565b503d614f00565b9050614f2b8461577e565b15614fb357614f6e6000926020926040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301614cb2565b03925af160009181614f92575b50614f8957836148e4614839565b92909192614ebe565b614fac91925060203d602011614f1957614f0a8183610354565b9038614f7b565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff841660045260246000fd5b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff881660045260246000fd5b61505591975060203d6020116132a2576132938183610354565b9538614da4565b90916060828403126100e75781516150738161139c565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e7578151916150a2836103b6565b916150b06040519384610354565b838352602084830101116100e7576040926150d191602080850191016103e7565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a0840152608061514e61511a604084015160a060c088015261012087019061040a565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e088015261040a565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106151d75750505061ffff909516602083015261038692916060916151bb9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff168552602090810151818601526040909401939092019160010161518d565b906000916152d0938361527773ffffffffffffffffffffffffffffffffffffffff61525c67ffffffffffffffff60208701511667ffffffffffffffff166000526004602052604060002090565b541673ffffffffffffffffffffffffffffffffffffffff1690565b92604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f000000000000000000000000000000000000000000000000000000000000000090600486016150d7565b03925af1908115611fe157600090600092615325575b50156152ef5750565b611e0e906040519182917f0a8d6e8c000000000000000000000000000000000000000000000000000000008352600483016113c7565b905061534491503d806000833e61533c8183610354565b81019061505c565b5090386152e6565b9291615356613865565b91808210156156f95761537061396a613944848489613622565b600160ff821603613f635750602182018181116156ca57615399613a748260018601858a6136e0565b84528181101561569b57613a9b61396a6139446153b793858a613622565b820191602283019082821161566c576133388260226153d89301858a6136e0565b60208501528181101561563d576153fb613a9b61396a613944602294868b613622565b83010191600183019082821161560e5761333882602361541d9301858a6136e0565b6040850152818110156155df57615440613a9b61396a613944600194868b613622565b8301019160018301908282116155b0576133388260026154629301858a6136e0565b60608501528181101561558157615485613a9b61396a613944600194868b613622565b8301016001810192828411615552576133388460026154a69301858a6136e0565b60808501526003810192828411615523576002916154d1613bb2613a56613a5088600196898e6136e0565b010101948186116154f4576154eb926133389287926136e0565b60a08201529190565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601c60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601b60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601a60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b6157527f331710310000000000000000000000000000000000000000000000000000000082615e31565b908161576c575b81615762575090565b61086f9150615dd1565b905061577781615d0b565b1590615759565b6157527faff2afbf0000000000000000000000000000000000000000000000000000000082615e31565b9080601f830112156100e75781516157bf81612010565b926157cd6040519485610354565b81845260208085019260051b8201019283116100e757602001905b8282106157f55750505090565b60208091835161580481611220565b8152019101906157e8565b906020828203126100e757815167ffffffffffffffff81116100e75761086f92016157a8565b95949060019460a09467ffffffffffffffff6158899573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c086019061040a565b930152565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152606095909291906020848060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa938415611fe1576000946159f2575b5061593384615728565b615951575b5050505050508051156159485790565b5061086f614067565b6000959650906159a673ffffffffffffffffffffffffffffffffffffffff92604051988997889687957f89720a6200000000000000000000000000000000000000000000000000000000875260048701615835565b0392165afa908115611fe1576000916159cf575b506159c481615e93565b388080808080615938565b6159ec91503d806000833e6159e48183610354565b81019061580f565b386159ba565b615a0c91945060203d6020116132a2576132938183610354565b9238615929565b90916060828403126100e757815167ffffffffffffffff81116100e75783615a3c9184016157a8565b92602083015167ffffffffffffffff81116100e757604091615a5f9185016157a8565b92015160ff811681036100e75790565b90615a9a7f7909b5490000000000000000000000000000000000000000000000000000000082615e31565b80615b9b575b80615b8c575b615ac4575b5050615ab5614067565b90615abe61404b565b90600090565b6040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9290921660048301526000908290602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015611fe157600080928192615b66575b50615b3981615e93565b615b4283615e93565b805115801590615b5a575b615b575750615aab565b92565b5060ff82161515615b4d565b9150615b8492503d8091833e615b7c8183610354565b810190615a13565b909138615b2f565b50615b9681615dd1565b615aa6565b50615ba581615d0b565b15615aa0565b6002548110156120c15760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b805460005b818110615bf457505060009055565b80615c01600192856140d1565b90549060031b1c6000528184016020526000604081205501615be5565b600081815260036020526040902054615ca9576002546801000000000000000081101561031657615c90615c5b82600185940160025560026140d1565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b6000828152600182016020526040902054615d0457805490680100000000000000008210156103165782615ced615c5b8460018096018555846140d1565b905580549260005201602052604060002055600190565b5050600090565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252615d6b604483610354565b6179185a10615da7576020926000925191617530fa6000513d82615d9b575b5081615d94575090565b9050151590565b60201115915038615d8a565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252615d6b604483610354565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252615d6b604483610354565b80519060005b828110615ea557505050565b6001810180821161300e575b838110615ec15750600101615e99565b73ffffffffffffffffffffffffffffffffffffffff615ee083856120c6565b5116615ef2610b04610b4d84876120c6565b14615eff57600101615eb1565b611d4d615f0f610b4d84866120c6565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260249056fea164736f6c634300081a000a",
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
	return common.HexToHash("0x6db4162777b6c980e778bb05a3d9e050f3792b091287ff0d4f3d51bdcd7427db")
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
