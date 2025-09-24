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
	Bin: "0x6101006040523461023857604051601f6158a638819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615652908161025482396080518181816101560152610acb015260a0518181816101b90152610a4e015260c0518181816101e10152818161455a0152615024015260e05181818161017d0152612b540152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100d7578063181f5a77146100d257806320f81c88146100cd5780635215505b146100c85780636b8be52c146100c357806379ba5097146100be5780637ce1552a146100b95780638da5cb5b146100b4578063d2b33733146100af578063e9d68a8e146100aa578063f054ac57146100a55763f2fde38b146100a057600080fd5b6116f5565b6113a1565b6112c1565b6111ee565b61119c565b6110fd565b610f2c565b610882565b6106fe565b610513565b61041e565b6100ec565b60009103126100e757565b600080fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e7576000606060405161012b81610290565b828152826020820152826040820152015261025d60405161014b81610290565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff8211176102ac57604052565b610261565b60a0810190811067ffffffffffffffff8211176102ac57604052565b6020810190811067ffffffffffffffff8211176102ac57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ac57604052565b6040519061033960a0836102e9565b565b6040519061033960c0836102e9565b60405190610339610100836102e9565b604051906103396040836102e9565b67ffffffffffffffff81116102ac57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906103b26020836102e9565b60008252565b60005b8381106103cb5750506000910152565b81810151838201526020016103bb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610417815180928187528780880191016103b8565b0116010190565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75761025d604080519061045f81836102e9565b601782527f43435641676772656761746f7220312e372e302d6465760000000000000000006020830152519182916020835260208301906103db565b9181601f840112156100e75782359167ffffffffffffffff83116100e757602083818601950101116100e757565b906020808351928381520192019060005b8181106104e75750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016104da565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576105c560ff61059f6105766105706105b795369060040161049b565b9061338c565b67ffffffffffffffff8151169061012061059360e08301516117e9565b60601c91015191613af7565b929193906040519586956060875260608701906104c9565b9085820360208701526104c9565b911660408301520390f35b61063b9173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061062a610618604085015160a0604086015260a08501906103db565b606085015184820360608601526104c9565b9201519060808184039101526104c9565b90565b6040810160408252825180915260206060830193019060005b8181106106de575050506020818303910152815180825260208201916020808360051b8301019401926000915b83831061069357505050505090565b90919293946020806106cf837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516105d0565b97019301930191939290610684565b825167ffffffffffffffff16855260209485019490920191600101610657565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760025461073981611853565b9061074760405192836102e9565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061077482611853565b0160005b81811061083a57505061078a81611897565b9060005b8181106107a657505061025d6040519283928361063e565b806107de6107c56107b86001946152b9565b67ffffffffffffffff1690565b6107cf8387611927565b9067ffffffffffffffff169052565b61081e6108196107ff6107f18488611927565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b611a9b565b6108288287611927565b526108338186611927565b500161078e565b60209061084561186b565b82828701015201610778565b9181601f840112156100e75782359167ffffffffffffffff83116100e7576020808501948460051b0101116100e757565b346100e75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576108d190369060040161049b565b60243567ffffffffffffffff81116100e7576108f1903690600401610851565b92909160443567ffffffffffffffff81116100e757610914903690600401610851565b93909261092760015460ff9060a01c1690565b610f025761098d90610973740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b610985610980858361338c565b613f34565b933691611032565b6020815191012093610a3560206109da6109b26107b8875167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610efd57600091610ece575b50610e8357610aad610aa9610a9f6107ff865167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b610e3857602083015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603610e015750808603610dcf5760e083019384516014815103610d955750610b5e610b29855167ffffffffffffffff1690565b6040860196610b40885167ffffffffffffffff1690565b610b58610b5260c08a015193516117e9565b60601c90565b92613f40565b96610b7d610b76896000526005602052604060002090565b5460ff1690565b610b86816110d1565b8015908115610d81575b5015610d1257610cb9610caa610ca5966107f1610c897f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df298610c848d8f9967ffffffffffffffff9b610c5891610ccb9b610c22610bf78f6000526005602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957fd2b337330000000000000000000000000000000000000000000000000000000060208801528b60248801611dc2565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b613fc1565b969015610cfd576002998a916000526005602052604060002090565b611b62565b965167ffffffffffffffff1690565b91836040519485941697169583611fb3565b0390a4610cfb7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b005b6003998a916000526005602052604060002090565b610d7d8787610d3b610d2c895167ffffffffffffffff1690565b915167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150610d8e816110d1565b1438610b90565b610dcb906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611b51565b0390fd5b7fb5ace4f300000000000000000000000000000000000000000000000000000000600052600486905260245260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b610d7d610e4d845167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610d7d610e98845167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b610ef0915060203d602011610ef6575b610ee881836102e9565b810190611b30565b38610a7f565b503d610ede565b611b45565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760005473ffffffffffffffffffffffffffffffffffffffff81163303610feb577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff8116036100e757565b359061033982611015565b92919261103e82610369565b9161104c60405193846102e9565b8294818452818301116100e7578281602093846000960137010152565b9080601f830112156100e75781602061063b93359101611032565b73ffffffffffffffffffffffffffffffffffffffff8116036100e757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600411156110db57565b6110a2565b9060048210156110db5752565b60208101929161033991906110e0565b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043561113881611015565b6024359061114582611015565b6044359167ffffffffffffffff83116100e75761116961117c933690600401611069565b906064359261117784611084565b613f40565b600052600560205261025d60ff60406000205416604051918291826110ed565b346100e75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100e75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100e75760443567ffffffffffffffff81116100e75761127c903690600401610851565b916064359267ffffffffffffffff84116100e7576112a1610cfb943690600401610851565b939092602435906004016128de565b90602061063b9281815201906105d0565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75767ffffffffffffffff60043561130581611015565b61130d61186b565b5016600052600460205261025d6040600020611390600360405192611331846102b1565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c16151560208501526040516113748161136d816001860161198e565b03826102e9565b604085015261138560028201611a80565b606085015201611a80565b6080820152604051918291826112b0565b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75760043567ffffffffffffffff81116100e7576113f0903690600401610851565b906113f9614948565b6000905b82821061140657005b611419611414838584612c95565b612d52565b60208101916114336107b8845167ffffffffffffffff1690565b156116cb5761147561145c61145c845173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b1580156116be575b6114cc5760808201949060005b865180518210156114f65761145c6114a5836114bf93611927565b5173ffffffffffffffffffffffffffffffffffffffff1690565b156114cc5760010161148a565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050939194929060009460a08701955b8651805182101561152e5761145c6114a58361152193611927565b156114cc57600101611506565b505094959095606082018051805180159182156116a8575b50506114cc576107b86114a59461166561169e9461165b8a6116516116919761164861160760019f9c61159d7f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259e51895190614993565b6115b26107ff8b5167ffffffffffffffff1690565b9e8f6115c16040840151151590565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b8d9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518d8c01612e84565b5160028a01612faf565b5160038801612faf565b61168261167d6107b8835167ffffffffffffffff1690565b6155c1565b505167ffffffffffffffff1690565b9260405191829182613043565b0390a201906113fd565b6020012090506116b6612e07565b143880611546565b506080820151511561147d565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100e75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100e75773ffffffffffffffffffffffffffffffffffffffff60043561174581611084565b61174d614948565b163381146117bf57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b90602082519201517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611821575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b67ffffffffffffffff81116102ac5760051b60200190565b60405190611878826102b1565b6060608083600081526000602082015282604082015282808201520152565b906118a182611853565b6118ae60405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06118dc8294611853565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156119225760200190565b6118e6565b80518210156119225760209160051b010190565b90600182811c92168015611984575b602083101461195557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161194a565b6000929181549161199e8361193b565b80835292600181169081156119f457506001146119ba57505050565b60009081526020812093945091925b8383106119da575060209250010190565b6001816020929493945483858701015201910191906119c9565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b906020825491828152019160005260206000209060005b818110611a545750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201611a47565b90610339611a949260405193848092611a30565b03836102e9565b9060036080604051611aac816102b1565b611b22819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c1615156020840152604051611aed8161136d816001860161198e565b6040840152604051611b068161136d8160028601611a30565b6060840152611b1b6040518096819301611a30565b03846102e9565b0152565b801515036100e757565b908160209103126100e7575161063b81611b26565b6040513d6000823e3d90fd5b90602061063b9281815201906103db565b9060048110156110db5760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b838310611bc557505050505090565b9091929394602080611c55837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895190815181526080611c44611c32611c208786015160a08987015260a08601906103db565b604086015185820360408701526103db565b606085015184820360608601526103db565b9201519060808184039101526103db565b97019301930191939290611bb6565b9160209082815201919060005b818110611c7e5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff8735611ca781611084565b168152019401929101611c71565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e75781360383136100e757565b90602083828152019260208260051b82010193836000925b848410611d6c5750505050505090565b909192939495602080611db2837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018852611dac8b88611cf4565b90611cb5565b9801940194019294939190611d5c565b94929361063b9694611f92611fa5949360808952611ded60808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a0152610140611f5f611f29611ef3611ebc8d611e78611e43606089015161016060e08501526101e08401906103db565b60808901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101008501526103db565b60a088015161ffff166101208301529060c088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b8d60e0870151906101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103db565b6101008501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101808f01526103db565b6101208401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101a08e0152611b99565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016101c08b01526103db565b9260208801528683036040880152611c64565b926060818503910152611d44565b80611fc460409261063b95946110e0565b81602082015201906103db565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e7576020019181360383136100e757565b919091357fffffffffffffffffffffffffffffffffffffffff00000000000000000000000081169260148110611821575050565b3561063b81611015565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100e7570180359067ffffffffffffffff82116100e757602001918160051b360383136100e757565b91909160a0818403126100e7576120c961032a565b9281358452602082013567ffffffffffffffff81116100e757816120ee918401611069565b6020850152604082013567ffffffffffffffff81116100e75781612113918401611069565b6040850152606082013567ffffffffffffffff81116100e75781612138918401611069565b6060850152608082013567ffffffffffffffff81116100e75761215b9201611069565b6080830152565b92919061216e81611853565b9361217c60405195866102e9565b602085838152019160051b8101918383116100e75781905b8382106121a2575050505050565b813567ffffffffffffffff81116100e7576020916121c387849387016120b4565b815201910190612194565b90821015611922576121e59160051b810190611fd1565b9091565b359061ffff821682036100e757565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100e757016020813591019167ffffffffffffffff82116100e7578160051b360383136100e757565b90602083828152019060208160051b85010193836000915b8383106122735750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61843603018112156100e757602061235660019386839401908135815261234861233d6123226123076122f788870187611cf4565b60a08a88015260a0870191611cb5565b6123146040870187611cf4565b908683036040880152611cb5565b61232f6060860186611cf4565b908583036060870152611cb5565b926080810190611cf4565b916080818503910152611cb5565b980196019493019190612263565b9061063b9593949273ffffffffffffffffffffffffffffffffffffffff6125b292168352608060208401526123ad6080840161239f83611027565b67ffffffffffffffff169052565b6123cd6123bc60208301611027565b67ffffffffffffffff1660a0850152565b6123ed6123dc60408301611027565b67ffffffffffffffff1660c0850152565b6125816125756125366124f76124b961246061242261240f6060890189611cf4565b61016060e08d01526101e08c0191611cb5565b61242f6080890189611cf4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808c8403016101008d0152611cb5565b61247b61246f60a089016121e9565b61ffff166101208b0152565b61248860c0880188611cf4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101408c0152611cb5565b6124c660e0870187611cf4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8403016101608b0152611cb5565b612505610100860186611cf4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80898403016101808a0152611cb5565b6125446101208501856121f8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888403016101a089015261224b565b91610140810190611cf4565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80858403016101c0860152611cb5565b9360408201526060818503910152611cb5565b604051906040820182811067ffffffffffffffff8211176102ac5760405260006020838281520152565b906125f982611853565b61260660405191826102e9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06126348294611853565b019060005b82811061264557505050565b6020906126506125c5565b82828501015201612639565b91908110156119225760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100e7570190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8f082019182116126f857565b61269c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd12082019182116126f857565b919082039182116126f857565b90916060828403126100e757815161274e81611b26565b92602083015167ffffffffffffffff81116100e75783019080601f830112156100e75781519161277d83610369565b9161278b60405193846102e9565b838352602084830101116100e7576040926127ac91602080850191016103b8565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806128296127f5604084015160a060c08801526101208701906103db565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103db565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106128a65750505061ffff909516602083015261033992916060916040820152019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101612868565b9391949092303303612c6b57612932612906610b5261290060e0890189611fd1565b90612022565b9161291087612056565b976101208801988461292c6129258c8c612060565b3691612162565b91614097565b97909160005b83518110156129e85761295461145c61145c6114a58488611927565b9061296a612962828d611927565b5187896121ce565b9290813b156100e7576000918a838d6129b3604051988996879586947f58bfa40a0000000000000000000000000000000000000000000000000000000086523060048701612364565b03925af1918215610efd576001926129cd575b5001612938565b806129dc60006129e2936102e9565b806100dc565b386129c6565b50949693509496505050612a066129ff8287612060565b90506125ef565b9260005b612a148388612060565b9050811015612a895780612a6d612a37600193612a31878c612060565b9061265c565b86612a57612a678c612a5f612a4f60c0830183611fd1565b949092612056565b9536906120b4565b923691611032565b906144c3565b612a778288611927565b52612a828187611927565b5001612a0a565b50949050929092610140810192612aa08483611fd1565b90501580612c63575b8015612c5a575b8015612c48575b612c4157612b48600094612b35612b3c612af361145c612ad96107ff89612056565b5473ffffffffffffffffffffffffffffffffffffffff1690565b95612afd81612056565b93612b16612b0e60c0840184611fd1565b929093611fd1565b949095612b2161032a565b9a8b5267ffffffffffffffff1660208b0152565b3691611032565b60408701523691611032565b606084015260808301527f000000000000000000000000000000000000000000000000000000000000000083612b8b612b865a61ffff85169061272a565b6126cb565b93612bc5604051978896879586947f3cf97983000000000000000000000000000000000000000000000000000000008652600486016127b2565b03925af1908115610efd57600090600092612c1a575b5015612be45750565b610dcb906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611b51565b9050612c3991503d806000833e612c3181836102e9565b810190612737565b509038612bdb565b5050505050565b50612c55610aa98661484a565b612ab7565b50843b15612ab0565b506000612aa9565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156119225760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100e7570190565b359061033982611084565b359061033982611b26565b9080601f830112156100e7578135612d0281611853565b92612d1060405194856102e9565b81845260208085019260051b8201019283116100e757602001905b828210612d385750505090565b602080918335612d4781611084565b815201910190612d2b565b60c0813603126100e757612d6461033b565b90612d6e81612cd5565b8252612d7c60208201611027565b6020830152612d8d60408201612ce0565b6040830152606081013567ffffffffffffffff81116100e757612db39036908301611069565b6060830152608081013567ffffffffffffffff81116100e757612dd99036908301612ceb565b608083015260a08101359067ffffffffffffffff82116100e757612dff91369101612ceb565b60a082015290565b60405160208101906000825260208152612e226040826102e9565b51902090565b818110612e33575050565b60008155600101612e28565b9190601f8111612e4e57505050565b610339926000526020600020906020601f840160051c83019310612e7a575b601f0160051c0190612e28565b9091508190612e6d565b919091825167ffffffffffffffff81116102ac57612eac81612ea6845461193b565b84612e3f565b6020601f8211600114612f0a578190612efb939495600092612eff575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612ec9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690612f3d84600052602060002090565b9160005b818110612f9757509583600195969710612f60575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080612f56565b9192602060018192868b015181550194019201612f41565b81519167ffffffffffffffff83116102ac576801000000000000000083116102ac576020908254848455808510613026575b500190600052602060002060005b838110612ffc5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501612fef565b61303d908460005285846000209182019101612e28565b38612fe1565b600361063b926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a060608401526130c561309260c085016001840161198e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301611a30565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201611a30565b60405190610160820182811067ffffffffffffffff8211176102ac576040526060610140836000815260006020820152600060408201528280820152826080820152600060a08201528260c08201528260e082015282610100820152826101208201520152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146126f85760010190565b90156119225790565b90821015611922570190565b90600882018092116126f857565b90600282018092116126f857565b90600182018092116126f857565b90602082018092116126f857565b919082018092116126f857565b909392938483116100e75784116100e7578101920390565b919091357fffffffffffffffff00000000000000000000000000000000000000000000000081169260088110613230575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110613296575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b604051906132d5826102b1565b60606080836000815282602082015282604082015282808201520152565b6040805190919061330483826102e9565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b82811061333b57505050565b6020906133466132c8565b8282850101520161332f565b604051906133616020836102e9565b600080835282815b82811061337557505050565b6020906133806132c8565b82828501015201613369565b906133956130f6565b5060258110613a51576133a66130f6565b9160006133e56133df6133b9858561318a565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613a23575061342a61341c613416613410613407600161319f565b600188886131e4565b906131fc565b60c01c90565b67ffffffffffffffff168552565b61349b61346d61343a600161319f565b61346861345761341661341061344f8561319f565b858b8b6131e4565b67ffffffffffffffff166020890152565b61319f565b61346861348a6134166134106134828561319f565b858a8a6131e4565b67ffffffffffffffff166040880152565b6134be6134b86133df6133b96134b08561315d565b948888613193565b60ff1690565b90846134ca83836131d7565b116139f65790816134ec612b356134e4846134f6966131d7565b8389896131e4565b60608801526131d7565b838110156139c8576135136134b86133df6133b96134b08561315d565b908461351f83836131d7565b1161399b579081613539612b356134e484613543966131d7565b60808801526131d7565b8361354d826131ad565b1161396e578061358261357761357161356b613482613587966131ad565b90613262565b60f01c90565b61ffff1660a0880152565b6131ad565b83811015613940576135a46134b86133df6133b96134b08561315d565b90846135b083836131d7565b116139135790816135ca612b356134e4846135d4966131d7565b60c08801526131d7565b838110156138e5576135f16134b86133df6133b96134b08561315d565b90846135fd83836131d7565b116138b8579081613617612b356134e484613621966131d7565b60e08801526131d7565b8361362b826131ad565b1161388a5761ffff61365661365061357161356b613648866131ad565b868a8a6131e4565b926131ad565b9116908461366483836131d7565b1161385d57908161367e612b356134e484613689966131d7565b6101008801526131d7565b9083613694836131ad565b11613830575061ffff6136ba61365061357161356b6136b2866131ad565b8689896131e4565b9116806137c757506136ca613352565b6101208501525b826136db826131ad565b116137985761ffff6136f861365061357161356b6136b2866131ad565b9116908361370683836131d7565b1161376957613727612b3561373294838761372187836131d7565b926131e4565b6101408601526131d7565b0361373a5790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b906137ed6137e56137d66132f3565b936101208801948552836131d7565b918585614b33565b909251926137fb8294611915565b52146136d1577fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008152600b600452602490fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008352600a600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526009600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526008600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526007600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526006600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526005600452602490fd5b507fb4205b4200000000000000000000000000000000000000000000000000000000815260048052602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526003600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526002600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526001600452602483fd5b7f789d326300000000000000000000000000000000000000000000000000000000825260ff16600452602490fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052610d7d6024906000600452565b60405190613a916020836102e9565b6000808352366020840137565b80156126f8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60ff1680156126f8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9290613b01613a82565b92805180613e58575b5050613b1690846151ad565b92909293613b666002613b45613b4b6003613b458b67ffffffffffffffff166000526004602052604060002090565b01611a80565b9867ffffffffffffffff166000526004602052604060002090565b613b91613b8c613b84613b7c87518651906131d7565b8a51906131d7565b8351906131d7565b611897565b9687956000805b8751821015613beb5790613be3600192613bc8613bb86114a5858d611927565b91613bc28161315d565b9c611927565b9073ffffffffffffffffffffffffffffffffffffffff169052565b018997613b98565b9750509193969092945060005b8551811015613c2d5780613c27613c146114a56001948a611927565b613bc8613c208b61315d565b9a8d611927565b01613bf8565b50919350919460005b8451811015613c6b5780613c65613c526114a560019489611927565b613bc8613c5e8a61315d565b998c611927565b01613c36565b5093919592509360005b828110613def575b5050600090815b818110613d505750508152835160005b818110613ca357508452929190565b926000959495915b8351831015613d4157613cc16114a58688611927565b73ffffffffffffffffffffffffffffffffffffffff613ce661145c6114a58789611927565b911603613d3057613cf690613a9e565b90613d11613d076114a58489611927565b613bc88789611927565b60ff8716613d20575b90613cab565b95613d2a90613ac9565b95613d1a565b9091613d3b9061315d565b91613d1a565b91509260019095949501613c94565b613d5d6114a58286611927565b73ffffffffffffffffffffffffffffffffffffffff81168015613de557600090815b868110613db9575b5050906001929115613d9c575b505b01613c84565b613db390613bc8613dac8761315d565b9688611927565b38613d94565b81613dca61145c6114a5848c611927565b14613dd757600101613d7f565b506001915081905038613d87565b5050600190613d96565b613dff61145c6114a58387611927565b15613e0c57600101613c75565b5091929360009591955b8351811015613e4b5780613e45613e326114a560019488611927565b613bc8613e3e8b61315d565b9a89611927565b01613e16565b5093929150933880613c7d565b9091935060018103613f07575060146060613e7283611915565b5101515103613ec457613ebc81613e9b610b526060613e93613b1696611915565b5101516117e9565b90866080613eb3613eab84611915565b515193611915565b51015192614fbb565b929038613b0a565b6060613ed2610dcb92611915565b5101516040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260048301611b51565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b613f3c6130f6565b5090565b9290612e229173ffffffffffffffffffffffffffffffffffffffff613f8e67ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906103db565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102e9565b9060405191613fd160c0846102e9565b60848352602083019060a03683375a612ee081111561402257600091613ff783926126fd565b82602083519301913090f1903d9060848211614019575b6000908286523e9190565b6084915061400e565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b91908110156119225760051b0190565b3561063b81611084565b906140a29291613af7565b919390926140af82611897565b926140b983611897565b94600091825b88518110156141c0576000805b8a88848982851061413c575b5050505050156140ea576001016140bf565b6140fa6114a5610d7d928b611927565b7f518d2ac50000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6114a56141749261416e6141698873ffffffffffffffffffffffffffffffffffffffff9761145c9661407d565b61408d565b95611927565b911614614183576001016140cc565b600191506141a2614198614169838b8b61407d565b613bc8888c611927565b6141b56141ae8761315d565b968b611927565b52388a8884896140d8565b509097965094939291909460ff811690816000985b8a518a10156142925760005b8b87821080614289575b1561427c5773ffffffffffffffffffffffffffffffffffffffff61421d61145c6114a58f61416e614169888f8f61407d565b9116146142325761422d9061315d565b6141e1565b939961424260019294939b613a9e565b9461425e614254614169838b8b61407d565b613bc88d8c611927565b61427161426a8c61315d565b9b8b611927565b525b019890916141d5565b5050919098600190614273565b508515156141eb565b9850925093959497509150816142bb57505050815181036142b257509190565b80825283529190565b610d7d92916142c99161272a565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b908160209103126100e7575161063b81611084565b604051906103b2826102cd565b908160209103126100e75760405190614337826102cd565b51815290565b9061063b916020815260e06144326143ff614366855161010060208701526101208601906103db565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff166060860152606086015160808601526143cb608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103db565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103db565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103db565b3d15614491573d9061447782610369565b9161448560405193846102e9565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff61063b949316815281602082015201906103db565b9092916144ce6125c5565b506144f361145c606084016144e381516152fd565b51602080825183010191016142fd565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa948515610efd57600095614819575b5073ffffffffffffffffffffffffffffffffffffffff85169485156147d6576145b3816148a0565b6147d6576145c3610aa9826148ca565b6147d65750614694916020916145d98886615389565b956145e2614312565b50805161460d608086840151930151936145fa61034a565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c082015261465c6103a3565b60e0820152604051809381927f390775370000000000000000000000000000000000000000000000000000000083526004830161433d565b03816000885af1600091816147a5575b506146e857846146b2614466565b90610dcb6040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401614496565b84909373ffffffffffffffffffffffffffffffffffffffff83160361473b575b5050505161473361471761035a565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b61474491615389565b908082108015614791575b6147595783614708565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b5061479c818361272a565b8351141561474f565b6147c891925060203d6020116147cf575b6147c081836102e9565b81019061431f565b90386146a4565b503d6147b6565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b61483c91955060203d602011614843575b61483481836102e9565b8101906142fd565b933861458b565b503d61482a565b6148747f85572ffb000000000000000000000000000000000000000000000000000000008261555f565b908161488e575b81614884575090565b61063b91506154ff565b905061489981615439565b159061487b565b6148747ff208a58f000000000000000000000000000000000000000000000000000000008261555f565b6148747faff2afbf000000000000000000000000000000000000000000000000000000008261555f565b6148747f05c7a8d0000000000000000000000000000000000000000000000000000000008261555f565b6148747f7909b549000000000000000000000000000000000000000000000000000000008261555f565b73ffffffffffffffffffffffffffffffffffffffff60015416330361496957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916149a18151846131d7565b928315614ace5760005b8481106149b9575050505050565b81811015614ab3576149ce6114a58286611927565b73ffffffffffffffffffffffffffffffffffffffff811680156114cc576149f4836131bb565b878110614a06575050506001016149ab565b84811015614a835773ffffffffffffffffffffffffffffffffffffffff614a306114a5838a611927565b168214614a3f576001016149f4565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614aae6114a5614aa8888561272a565b89611927565b614a30565b614ac96114a5614ac3848461272a565b85611927565b6149ce565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b359060208110614b06575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b929192614b3e6132c8565b9382811015614ebc57614b616133df6133b9614b598461315d565b938686613193565b600160ff821603614e8c575082614b77826131c9565b11614e5d5780614b9d614b97614b8f614ba4946131c9565b8387876131e4565b90614af8565b86526131c9565b82811015614e2e57614bc96134b86133df6133b9614bc18561315d565b948787613193565b83614bd482846131d7565b11614dff5781614bf5612b35614bed84614bff966131d7565b8388886131e4565b60208801526131d7565b82811015614dd057614c1c6134b86133df6133b9614bc18561315d565b83614c2782846131d7565b11614da15781614c40612b35614bed84614c4a966131d7565b60408801526131d7565b82811015614d7257614c676134b86133df6133b9614bc18561315d565b83614c7282846131d7565b11614d4357816134ec612b35614bed84614c8b966131d7565b82614c95826131ad565b11614d145761ffff614cb261365061357161356b6136b2866131ad565b91169183614cc084846131d7565b11614ce557612b35614cdb918361063b9661372187836131d7565b60808601526131d7565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b9080601f830112156100e7578151614f0281611853565b92614f1060405194856102e9565b81845260208085019260051b8201019283116100e757602001905b828210614f385750505090565b602080918351614f4781611084565b815201910190614f2b565b906020828203126100e757815167ffffffffffffffff81116100e75761063b9201614eeb565b909267ffffffffffffffff60809373ffffffffffffffffffffffffffffffffffffffff61063b9796168452166020830152604082015281606082015201906103db565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526060949293916020828060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa918215610efd57600092615130575b5061505f826148f4565b615098575b50505050815115615073575090565b61063b9150613b4560029167ffffffffffffffff166000526004602052604060002090565b8495509373ffffffffffffffffffffffffffffffffffffffff600094956150ee604051978896879586947f0ba375f900000000000000000000000000000000000000000000000000000000865260048601614f78565b0392165afa908115610efd5760009161510d575b509038808080615064565b61512a91503d806000833e61512281836102e9565b810190614f52565b38615102565b61514a91925060203d6020116148435761483481836102e9565b9038615055565b90916060828403126100e757815167ffffffffffffffff81116100e7578361517a918401614eeb565b92602083015167ffffffffffffffff81116100e75760409161519d918501614eeb565b92015160ff811681036100e75790565b9190916151d16108198267ffffffffffffffff166000526004602052604060002090565b90833b6151ee575b506060015191506151e8613a82565b90600090565b6151f78461491e565b156151d9576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015610efd57600080948192615291575b50805115801590615285575b61527e57506151d9565b9392909150565b5060ff82161515615274565b91506152b09294503d8091833e6152a881836102e9565b810190615151565b90939138615268565b6002548110156119225760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b908160209103126100e7575190565b60208151036153405761531960208251830101602083016152ee565b73ffffffffffffffffffffffffffffffffffffffff811190811561537d575b506153405750565b610dcb906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906103db565b61040091501038615338565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181615408575b5061540457826146b2614466565b9150565b61542b91925060203d602011615432575b61542381836102e9565b8101906152ee565b90386153f6565b503d615419565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526154996044836102e9565b6179185a106154d5576020926000925191617530fa6000513d826154c9575b50816154c2575090565b9050151590565b602011159150386154b8565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526154996044836102e9565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526154996044836102e9565b8060005260036020526040600020541560001461563f57600254680100000000000000008110156102ac5760018101600255600060025482101561192257600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
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
