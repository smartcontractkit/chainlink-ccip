// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccv_aggregator

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type CCVAggregatorAggregatedReport struct {
	Message InternalAny2EVMMessage
	Ccvs    []common.Address
	CcvData [][]byte
}

type CCVAggregatorSourceChainConfig struct {
	Router           common.Address
	IsEnabled        bool
	OnRamp           []byte
	DefaultCCV       []common.Address
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

type InternalAny2EVMMessage struct {
	Header       InternalHeader
	Sender       []byte
	Data         []byte
	Receiver     common.Address
	GasLimit     uint32
	TokenAmounts []InternalTokenTransfer
}

type InternalHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

type InternalTokenTransfer struct {
	SourcePoolAddress []byte
	DestTokenAddress  common.Address
	ExtraData         []byte
	Amount            *big.Int
}

var CCVAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.AggregatedReport\",\"components\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isPoolCCV\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023f57604051601f61473c38819003918201601f19168301916001600160401b038311848410176102445780849260809460405283398101031261023f5760405190600090608083016001600160401b0381118482101761022b5760405280516001600160401b0381168103610227578352602081015161ffff8116810361022757602084019081526040820151916001600160a01b0383168303610223576040850192835260600151926001600160a01b03841684036102205760608501938452331561021157600180546001600160a01b0319163317905582516001600160a01b03161580156101ff575b6101f05784516001600160401b0316156101e15784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a16040516144e1908161025b823960805181818161012e01526110b3015260a0518181816101910152611017015260c0518181816101cd01528181612d730152613377015260e05181818161015501528181611b3c0152613dcb0152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100c7578063181f5a77146100c257806345ec4b5f146100bd5780635215505b146100b85780635e36480c146100b357806379ba5097146100ae5780638da5cb5b146100a9578063e9d68a8e146100a4578063f054ac571461009f578063f2fde38b1461009a5763f734ef0e1461009557600080fd5b610ecf565b610ddb565b610d60565b610c80565b610c1d565b610b34565b610ad3565b61093a565b61076c565b61040a565b6100dc565b60009103126100d757565b600080fd5b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d7576101136116b1565b506102496040516101238161027c565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761029857604052565b61024d565b60a0810190811067ffffffffffffffff82111761029857604052565b6020810190811067ffffffffffffffff82111761029857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761029857604052565b6040519061032560c0836102d5565b565b6040519061032560a0836102d5565b60405190610325610100836102d5565b604051906103256040836102d5565b67ffffffffffffffff811161029857601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6040519061039e6020836102d5565b60008252565b60005b8381106103b75750506000910152565b81810151838201526020016103a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610403815180928187528780880191016103a4565b0116010190565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d757610249604080519061044b81836102d5565b601782527f43435641676772656761746f7220312e372e302d6465760000000000000000006020830152519182916020835260208301906103c7565b67ffffffffffffffff8116036100d757565b359061032582610487565b91908260809103126100d7576040516104bc8161027c565b60608082948035845260208101356104d381610487565b602085015260408101356104e681610487565b60408501520135916104f783610487565b0152565b92919261050782610355565b9161051560405193846102d5565b8294818452818301116100d7578281602093846000960137010152565b9080601f830112156100d75781602061054d933591016104fb565b90565b73ffffffffffffffffffffffffffffffffffffffff8116036100d757565b359061032582610550565b359063ffffffff821682036100d757565b67ffffffffffffffff81116102985760051b60200190565b81601f820112156100d7578035906105b98261058a565b926105c760405194856102d5565b82845260208085019360051b830101918183116100d75760208101935b8385106105f357505050505090565b843567ffffffffffffffff81116100d757820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126100d7576040519161063f8361027c565b602082013567ffffffffffffffff81116100d75785602061066292850101610532565b8352604082013561067281610550565b602084015260608201359267ffffffffffffffff84116100d7576080836106a0886020809881980101610532565b6040840152013560608201528152019401936105e4565b919091610120818403126100d7576106cd610316565b926106d881836104a4565b8452608082013567ffffffffffffffff81116100d757816106fa918401610532565b602085015260a082013567ffffffffffffffff81116100d7578161071f918401610532565b604085015261073060c0830161056e565b606085015261074160e08301610579565b608085015261010082013567ffffffffffffffff81116100d75761076592016105a2565b60a0830152565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d7576107be6107c39136906004016106b7565b61198f565b005b906020808351928381520192019060005b8181106107e35750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016107d6565b61054d9173ffffffffffffffffffffffffffffffffffffffff82511681526020820151151560208201526080610869610857604085015160a0604086015260a08501906103c7565b606085015184820360608601526107c5565b9201519060808184039101526107c5565b6040810160408252825180915260206060830193019060005b81811061091a575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106108cf57505050505090565b909192939460208061090b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161080f565b970193019301919392906108c0565b825167ffffffffffffffff16855260209485019490920191600101610893565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d7576002546109758161058a565b9061098360405192836102d5565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06109b08261058a565b0160005b818110610a765750506109c681611cb2565b9060005b8181106109e25750506102496040519283928361087a565b80610a1a610a016109f460019461418d565b67ffffffffffffffff1690565b610a0b838761179c565b9067ffffffffffffffff169052565b610a5a610a55610a3b610a2d848861179c565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b611e61565b610a64828761179c565b52610a6f818661179c565b50016109ca565b602090610a81611c86565b828287010152016109b4565b60041115610a9757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b906004821015610a975752565b346100d75760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d7576020610b25600435610b1381610487565b60243590610b2082610487565b611f29565b610b326040518092610ac6565bf35b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760005473ffffffffffffffffffffffffffffffffffffffff81163303610bf3577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b90602061054d92818152019061080f565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75767ffffffffffffffff600435610cc481610487565b610ccc611c86565b501660005260046020526102496040600020610d4f600360405192610cf08461029d565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c1615156020850152604051610d3381610d2c8160018601611d54565b03826102d5565b6040850152610d4460028201611e46565b606085015201611e46565b608082015260405191829182610c6f565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d757366023820112156100d757806004013567ffffffffffffffff81116100d7573660248260051b840101116100d75760246107c39201611f72565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75773ffffffffffffffffffffffffffffffffffffffff600435610e2b81610550565b610e336131a8565b16338114610ea557807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d757806004019060607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100d75760015460a01c60ff1661168757610f92740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b610fac610fa7610fa28480612626565b612659565b6131f3565b91610fc26020610fbc8380612626565b01612664565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b1660048201529092906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa9081156112da57600091611658575b50611620576110946110906110868567ffffffffffffffff166000526004602052604060002090565b5460a01c60ff1690565b1590565b6115e8576110a76040610fbc8480612626565b67ffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036115a057602481019060446110eb8385612683565b92905001906110fa8285612683565b919050036115535761111a6111146060610fbc8680612626565b85611f29565b9161112483610a8d565b82158015611540575b156114e95761113a6126d7565b6111526111478680612626565b610100810190612683565b905061141a575b6111939161116c6020610fbc8880612626565b61118b61118460c061117e8a80612626565b0161272c565b9288612683565b929091613512565b9061119e8580612626565b6040516111df816111b3602082019485612a77565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102d5565b5190209660005b82518110156112df5761123561121c61121c611202848761179c565b5173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906112408880612626565b61125e61124d888b612683565b611257858a61179c565b5191612a88565b93803b156100d7578c600080946112a58d604051998a97889687957f01e0fdea00000000000000000000000000000000000000000000000000000000875260048701612aa3565b03925af19182156112da576001926112bf575b50016111e6565b806112ce60006112d4936102d5565b806100cc565b386112b8565b611972565b50866113016112fb606084510167ffffffffffffffff90511690565b82613a4a565b61130a82613d89565b9261132c82611326606084510167ffffffffffffffff90511690565b85613ad5565b61133582610a8d565b600282036113c6575b67ffffffffffffffff7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df29151611396611382606083015167ffffffffffffffff1690565b915196836040519485941697169583612b01565b0390a46107c37fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b6113cf82610a8d565b600382031561133e575160600151611416929067ffffffffffffffff16907f926c5a3e00000000000000000000000000000000000000000000000000000000600052612ade565b6000fd5b50600161142a6111478680612626565b9050036114ac57611193906114a4611453602061117e61144d6111478a80612626565b906126f3565b6114626020610fbc8980612626565b606061147461144d6111478b80612626565b01359061149e61149761148d61144d6111478d80612626565b6040810190612736565b36916104fb565b9261330e565b909150611159565b6114166114bc6111478680612626565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045250602490565b61141685611504606089510167ffffffffffffffff90511690565b7f3b5754190000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff91821660045216602452604490565b5061154a83610a8d565b6003831461112d565b611416935061156561156d9284612683565b939050612683565b7fb5ace4f30000000000000000000000000000000000000000000000000000000060005260049290925250602452604490565b6114166115b26040610fbc8580612626565b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b61167a915060203d602011611680575b61167281836102d5565b81019061266e565b3861105d565b503d611668565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906116be8261027c565b60006060838281528260208201528260408201520152565b604051906040820182811067ffffffffffffffff8211176102985760405260006020838281520152565b9061170a8261058a565b61171760405191826102d5565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611745829461058a565b019060005b82811061175657505050565b6020906117616116d6565b8282850101520161174a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156117b05760209160051b010190565b61176d565b801515036100d757565b90916060828403126100d75781516117d6816117b5565b92602083015167ffffffffffffffff81116100d75783019080601f830112156100d75781519161180583610355565b9161181360405193846102d5565b838352602084830101116100d75760409261183491602080850191016103a4565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806118b161187d604084015160a060c08801526101208701906103c7565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103c7565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b81811061193a5750505061ffff9095166020830152610325929160609161191e9063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff16855260209081015181860152604090940193909201916001016118f0565b6040513d6000823e3d90fd5b90602061054d9281815201906103c7565b90303303611c5c5760a08201916119a7835151611700565b9160005b84518051821015611a215790611a056119c68260019461179c565b516020860151606087015173ffffffffffffffffffffffffffffffffffffffff16906119ff602089510167ffffffffffffffff90511690565b92612ce5565b611a0f828761179c565b52611a1a818661179c565b50016119ab565b50509250805191611a3f6020845194015167ffffffffffffffff1690565b906020830151916040840192611a6d845192611a59610327565b97885267ffffffffffffffff166020880152565b60408601526060850152608084015251511580611c3e575b8015611c1b575b8015611be9575b611be557611b659181611ad861121c611abe610a3b6020600097510167ffffffffffffffff90511690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9083611b0c6060611af0608085015163ffffffff1690565b93015173ffffffffffffffffffffffffffffffffffffffff1690565b93604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f0000000000000000000000000000000000000000000000000000000000000000906004860161183a565b03925af19081156112da57600090600092611bbe575b5015611b845750565b611bba906040519182917f0a8d6e8c0000000000000000000000000000000000000000000000000000000083526004830161197e565b0390fd5b9050611bdd91503d806000833e611bd581836102d5565b8101906117bf565b509038611b7b565b5050565b50611c16611090611c11606084015173ffffffffffffffffffffffffffffffffffffffff1690565b613064565b611a93565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b15611a8c565b5063ffffffff611c55608083015163ffffffff1690565b1615611a85565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611c938261029d565b6060608083600081526000602082015282604082015282808201520152565b90611cbc8261058a565b611cc960405191826102d5565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611cf7829461058a565b0190602036910137565b90600182811c92168015611d4a575b6020831014611d1b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611d10565b60009291815491611d6483611d01565b8083529260018116908115611dba5750600114611d8057505050565b60009081526020812093945091925b838310611da0575060209250010190565b600181602092949394548385870101520191019190611d8f565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b906020825491828152019160005260206000209060005b818110611e1a5750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201611e0d565b90610325611e5a9260405193848092611df6565b03836102d5565b9060036080604051611e728161029d565b6104f7819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c1615156020840152604051611eb381610d2c8160018601611d54565b6040840152604051611ecc81610d2c8160028601611df6565b6060840152611ee16040518096819301611df6565b03846102d5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211611f2457565b611ee8565b611f3582607f92613162565b9116906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611f24576003911c166004811015610a975790565b90611f7b6131a8565b60005b818110611f8a57505050565b611f9d611f988284866121d0565b612282565b90611fb3602083015167ffffffffffffffff1690565b67ffffffffffffffff81169081156121a657611fe961121c61121c865173ffffffffffffffffffffffffffffffffffffffff1690565b158015612199575b612159576120139067ffffffffffffffff166000526004602052604060002090565b9260608101805180518015918215612183575b50506121595760019461213a60a08461211e612150956120b361206d60407f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259a0151151590565b87547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178755565b6121156120d4845173ffffffffffffffffffffffffffffffffffffffff1690565b879073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518a86016123b4565b61212f6080820151600286016124df565b0151600383016124df565b612143846141c2565b5060405191829182612573565b0390a201611f7e565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b602001209050612191612337565b143880612026565b5060808401515115611ff1565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156117b05760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100d7570190565b3590610325826117b5565b9080601f830112156100d75781356122328161058a565b9261224060405194856102d5565b81845260208085019260051b8201019283116100d757602001905b8282106122685750505090565b60208091833561227781610550565b81520191019061225b565b60c0813603126100d757612294610316565b9061229e8161056e565b82526122ac60208201610499565b60208301526122bd60408201612210565b6040830152606081013567ffffffffffffffff81116100d7576122e39036908301610532565b6060830152608081013567ffffffffffffffff81116100d757612309903690830161221b565b608083015260a08101359067ffffffffffffffff82116100d75761232f9136910161221b565b60a082015290565b604051602081019060008252602081526123526040826102d5565b51902090565b818110612363575050565b60008155600101612358565b9190601f811161237e57505050565b610325926000526020600020906020601f840160051c830193106123aa575b601f0160051c0190612358565b909150819061239d565b919091825167ffffffffffffffff8111610298576123dc816123d68454611d01565b8461236f565b6020601f821160011461243a57819061242b93949560009261242f575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0151905038806123f9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061246d84600052602060002090565b9160005b8181106124c757509583600195969710612490575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080612486565b9192602060018192868b015181550194019201612471565b81519167ffffffffffffffff831161029857680100000000000000008311610298576020908254848455808510612556575b500190600052602060002060005b83811061252c5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161251f565b61256d908460005285846000209182019101612358565b38612511565b600361054d926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a060608401526125f56125c260c0850160018401611d54565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301611df6565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201611df6565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1813603018212156100d7570190565b61054d9036906106b7565b3561054d81610487565b908160209103126100d7575161054d816117b5565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100d7570180359067ffffffffffffffff82116100d757602001918160051b360383136100d757565b604051906126e66020836102d5565b6000808352366020840137565b90156117b0578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81813603018212156100d7570190565b3561054d81610550565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100d7570180359067ffffffffffffffff82116100d7576020019181360383136100d757565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100d757016020813591019167ffffffffffffffff82116100d75781360383136100d757565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100d757016020813591019167ffffffffffffffff82116100d7578160051b360383136100d757565b90602083828152019160208260051b8501019381936000915b8483106128925750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08282030183528635907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81863603018212156100d7576020809187600194019060608061295b61291961290b8680612787565b6080875260808701916127d7565b73ffffffffffffffffffffffffffffffffffffffff8787013561293b81610550565b168786015261294d6040870187612787565b9086830360408801526127d7565b9301359101529801930193019194939290612882565b61054d918135815267ffffffffffffffff602083013561299081610487565b16602082015267ffffffffffffffff60408301356129ad81610487565b16604082015267ffffffffffffffff60608301356129ca81610487565b166060820152612a68612a146129f96129e66080860186612787565b61012060808701526101208601916127d7565b612a0660a0860186612787565b9085830360a08701526127d7565b92612a41612a2460c0830161056e565b73ffffffffffffffffffffffffffffffffffffffff1660c0850152565b612a5d612a5060e08301610579565b63ffffffff1660e0850152565b610100810190612816565b91610100818503910152612869565b90602061054d928181520190612971565b908210156117b057612a9f9160051b810190612736565b9091565b9594929161032594606093612ac3612ad69460808b5260808b0190612971565b9260208a015288830360408a01526127d7565b940190610ac6565b929160449067ffffffffffffffff61032593816064971660045216602452610ac6565b80612b1260409261054d9594610ac6565b81602082015201906103c7565b908160209103126100d7575161054d81610550565b6040519061039e826102b9565b908160209103126100d75760405190612b59826102b9565b51815290565b9061054d916020815260e0612c54612c21612b88855161010060208701526101208601906103c7565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff16606086015260608601516080860152612bed608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103c7565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103c7565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103c7565b3d15612cb3573d90612c9982610355565b91612ca760405193846102d5565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff61054d949316815281602082015201906103c7565b909291612cf06116d6565b50602082015173ffffffffffffffffffffffffffffffffffffffff166040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9485156112da57600095613033575b5073ffffffffffffffffffffffffffffffffffffffff8516948515612ff057612dcc816130ba565b612ff057612ddc611090826130e4565b612ff05750612eae91602091612df28886613f4f565b95612dfb612b34565b506060810151612e276040835193015193612e14610336565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c0820152612e7661038f565b60e0820152604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301612b5f565b03816000885af160009181612fbf575b50612f025784612ecc612c88565b90611bba6040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401612cb8565b84909373ffffffffffffffffffffffffffffffffffffffff831603612f55575b50505051612f4d612f31610346565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b612f5e91613f4f565b908082108015612fab575b612f735783612f22565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b50612fb68183611f17565b83511415612f69565b612fe291925060203d602011612fe9575b612fda81836102d5565b810190612b41565b9038612ebe565b503d612fd0565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b61305691955060203d60201161305d575b61304e81836102d5565b810190612b1f565b9338612da4565b503d613044565b61308e7f85572ffb000000000000000000000000000000000000000000000000000000008261412b565b90816130a8575b8161309e575090565b61054d91506140cb565b90506130b381614005565b1590613095565b61308e7ff208a58f000000000000000000000000000000000000000000000000000000008261412b565b61308e7faff2afbf000000000000000000000000000000000000000000000000000000008261412b565b61308e7f05c7a8d0000000000000000000000000000000000000000000000000000000008261412b565b61308e7f7909b549000000000000000000000000000000000000000000000000000000008261412b565b9067ffffffffffffffff6131a4921660005260056020526701ffffffffffffff60406000209160071c1667ffffffffffffffff16600052602052604060002090565b5490565b73ffffffffffffffffffffffffffffffffffffffff6001541633036131c957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160c0810181811067ffffffffffffffff8211176102985760609160a09160405261321e6116b1565b815282602082015282604082015260008382015260006080820152015290565b9080601f830112156100d75781516132558161058a565b9261326360405194856102d5565b81845260208085019260051b8201019283116100d757602001905b82821061328b5750505090565b60208091835161329a81610550565b81520191019061327e565b906020828203126100d757815167ffffffffffffffff81116100d75761054d920161323e565b909267ffffffffffffffff60809373ffffffffffffffffffffffffffffffffffffffff61054d9796168452166020830152604082015281606082015201906103c7565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526060949293916020828060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9182156112da57600092613489575b506133b28261310e565b6133f1575b505050508151156133c6575090565b61054d91506133eb60029167ffffffffffffffff166000526004602052604060002090565b01611e46565b8495509373ffffffffffffffffffffffffffffffffffffffff60009495613447604051978896879586947f0ba375f9000000000000000000000000000000000000000000000000000000008652600486016132cb565b0392165afa9081156112da57600091613466575b5090388080806133b7565b61348391503d806000833e61347b81836102d5565b8101906132a5565b3861345b565b6134a391925060203d60201161305d5761304e81836102d5565b90386133a8565b91908110156117b05760051b0190565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611f245760010190565b8015611f24577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b94916003906135486135296133eb979694896142a2565b949198909967ffffffffffffffff166000526004602052604060002090565b9061355284611cb2565b9261355c85611cb2565b9661356686611cb2565b96600094855b8b518110156136d1576000805b8d8b848a8285106135ef575b5050505050156135975760010161356c565b6135a7611202611416928e61179c565b7fbd76195f000000000000000000000000000000000000000000000000000000006000529073ffffffffffffffffffffffffffffffffffffffff604492166004526000602452565b6112026136279261362161361c8873ffffffffffffffffffffffffffffffffffffffff9761121c966134aa565b61272c565b9561179c565b91161461363657600101613579565b60019891508b8d8b8d61365261364c868661179c565b51151590565b6136c15792846136a5876136ab958f84966136858461367f61361c6136b69f9e6136b19e6136a0976134aa565b9261179c565b9073ffffffffffffffffffffffffffffffffffffffff169052565b61179c565b5261179c565b60019052565b6134ba565b96388d8b848a613585565b9150508f9150848a9b949b613585565b50919950919960005b8a518110156137ff576000805b8a8d8b858a82861061375e575b50505050505015613707576001016136da565b613717611202611416928d61179c565b7fbd76195f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff166004526001602452604490565b61120261378b9261362161361c8973ffffffffffffffffffffffffffffffffffffffff9761121c966134aa565b91161461379b57506001016136e7565b9791508b898b6137b061364c8560019d61179c565b6137ef5792806137dc866136b1956136a06136ab968f8461367f61361c896137e39f9e613685956134aa565b528d61179c565b95388a8d8b858a6136f4565b8f91508d9250858a9b959b6136f4565b50919990985060005b89518110156138e7576000805b898c8a8589828610613845575b5050505050501561383557600101613808565b613717611202611416928c61179c565b6112026138729261362161361c8973ffffffffffffffffffffffffffffffffffffffff9761121c966134aa565b9116146138825750600101613815565b9691508a888a61389761364c8560019c61179c565b6138d75792806138c4866136b1956136a06136ab966136858f9b61367f61361c896138cb9f9e849f6134aa565b528c61179c565b9438898c8a8589613822565b8c92508e915085899a959a613822565b509098975095929594919493909360ff811690816000995b8b518b10156139e05760005b8c868210806139d7575b156139ca5773ffffffffffffffffffffffffffffffffffffffff61394961121c8f611202906136218f888e61361c926134aa565b91161461395e57613959906134ba565b61390b565b9361396e909b999193929b6134e7565b938a61397d61364c838a61179c565b6139bf57916136b16136ab83806139ad858f8f9060019a8f8461367f61361c6139b49e613685946136a0976134aa565b528a61179c565b985b019990916138ff565b5050976001906139b6565b50509190996001906139b6565b50851515613915565b9950935094509297505081613a0857505050815181036139ff57509190565b80825283529190565b6114169291613a1691611f17565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b607f8216906801fffffffffffffffe67ffffffffffffffff83169260011b169180830460021490151715611f2457613ad29167ffffffffffffffff613a8f8584613162565b921660005260056020526701ffffffffffffff60406000209460071c169160036001831b921b191617929067ffffffffffffffff16600052602052604060002090565b55565b9091607f8316916801fffffffffffffffe67ffffffffffffffff84169360011b169280840460021490151715611f2457613b0f8482613162565b926004831015610a9757613ad29367ffffffffffffffff613b54931660005260056020526003604060002094831b921b191617936701ffffffffffffff9060071c1690565b67ffffffffffffffff16600052602052604060002090565b9080602083519182815201916020808360051b8301019401926000915b838310613b9857505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08560019503018652885190606080613c16613be685516080865260808601906103c7565b73ffffffffffffffffffffffffffffffffffffffff878701511687860152604086015185820360408701526103c7565b93015191015297019301930191939290613b89565b9061054d916020815267ffffffffffffffff60608351805160208501528260208201511660408501528260408201511682850152015116608082015260a0613cba613c866020850151610120848601526101408501906103c7565b60408501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030160c08601526103c7565b606084015173ffffffffffffffffffffffffffffffffffffffff1660e084015292608081015163ffffffff166101008401520151906101207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152613b6c565b90602082519201517fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613d57575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b6111b3613dc8613df2926040519283917f45ec4b5f00000000000000000000000000000000000000000000000000000000602084015260248301613c2b565b5a7f00000000000000000000000000000000000000000000000000000000000000009130906143ae565b509015613e05575060029061054d61038f565b9072c11c11c11c11c11c11c11c11c11c11c11c11c13314613e27575b60039190565b613e58613e3383613d1f565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f37c3be2900000000000000000000000000000000000000000000000000000000148015613f1b575b8015613ee7575b15613e2157611416613e9983613d1f565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b50613ef4613e3383613d1f565b7fea7f4b120000000000000000000000000000000000000000000000000000000014613e88565b50613f28613e3383613d1f565b7fafa32a2c0000000000000000000000000000000000000000000000000000000014613e81565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181613fce575b50613fca5782612ecc612c88565b9150565b90916020823d602011613ffd575b81613fe9602093836102d5565b81010312613ffa5750519038613fbc565b80fd5b3d9150613fdc565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526140656044836102d5565b6179185a106140a1576020926000925191617530fa6000513d82614095575b508161408e575090565b9050151590565b60201115915038614084565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526140656044836102d5565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526140656044836102d5565b6002548110156117b05760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b806000526003602052604060002054156000146142405760025468010000000000000000811015610298576001810160025560006002548210156117b057600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b50600090565b90916060828403126100d757815167ffffffffffffffff81116100d7578361426f91840161323e565b92602083015167ffffffffffffffff81116100d75760409161429291850161323e565b92015160ff811681036100d75790565b9190916142c6610a558267ffffffffffffffff166000526004602052604060002090565b90833b6142e3575b506060015191506142dd6126d7565b90600090565b6142ec84613138565b156142ce576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa80156112da57600080948192614386575b5080511580159061437a575b61437357506142ce565b9392909150565b5060ff82161515614369565b91506143a59294503d8091833e61439d81836102d5565b810190614246565b9093913861435d565b9391936143bb6084610355565b946143c960405196876102d5565b608486526143d76084610355565b947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602088019601368737833b156144aa575a90808210614480578291038060061c90031115614456576000918291825a9560208451940192f1905a9003923d906084821161444d575b6000908287523e929190565b60849150614441565b7f37c3be290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fafa32a2c0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f0c3b563c0000000000000000000000000000000000000000000000000000000060005260046000fdfea164736f6c634300081a000a",
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

func (_CCVAggregator *CCVAggregatorCaller) GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	var out []interface{}
	err := _CCVAggregator.contract.Call(opts, &out, "getExecutionState", sourceChainSelector, sequenceNumber)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_CCVAggregator *CCVAggregatorSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _CCVAggregator.Contract.GetExecutionState(&_CCVAggregator.CallOpts, sourceChainSelector, sequenceNumber)
}

func (_CCVAggregator *CCVAggregatorCallerSession) GetExecutionState(sourceChainSelector uint64, sequenceNumber uint64) (uint8, error) {
	return _CCVAggregator.Contract.GetExecutionState(&_CCVAggregator.CallOpts, sourceChainSelector, sequenceNumber)
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

func (_CCVAggregator *CCVAggregatorTransactor) Execute(opts *bind.TransactOpts, report CCVAggregatorAggregatedReport) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "execute", report)
}

func (_CCVAggregator *CCVAggregatorSession) Execute(report CCVAggregatorAggregatedReport) (*types.Transaction, error) {
	return _CCVAggregator.Contract.Execute(&_CCVAggregator.TransactOpts, report)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) Execute(report CCVAggregatorAggregatedReport) (*types.Transaction, error) {
	return _CCVAggregator.Contract.Execute(&_CCVAggregator.TransactOpts, report)
}

func (_CCVAggregator *CCVAggregatorTransactor) ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessage) (*types.Transaction, error) {
	return _CCVAggregator.contract.Transact(opts, "executeSingleMessage", message)
}

func (_CCVAggregator *CCVAggregatorSession) ExecuteSingleMessage(message InternalAny2EVMMessage) (*types.Transaction, error) {
	return _CCVAggregator.Contract.ExecuteSingleMessage(&_CCVAggregator.TransactOpts, message)
}

func (_CCVAggregator *CCVAggregatorTransactorSession) ExecuteSingleMessage(message InternalAny2EVMMessage) (*types.Transaction, error) {
	return _CCVAggregator.Contract.ExecuteSingleMessage(&_CCVAggregator.TransactOpts, message)
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

func (_CCVAggregator *CCVAggregator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CCVAggregator.abi.Events["ExecutionStateChanged"].ID:
		return _CCVAggregator.ParseExecutionStateChanged(log)
	case _CCVAggregator.abi.Events["OwnershipTransferRequested"].ID:
		return _CCVAggregator.ParseOwnershipTransferRequested(log)
	case _CCVAggregator.abi.Events["OwnershipTransferred"].ID:
		return _CCVAggregator.ParseOwnershipTransferred(log)
	case _CCVAggregator.abi.Events["SourceChainConfigSet"].ID:
		return _CCVAggregator.ParseSourceChainConfigSet(log)
	case _CCVAggregator.abi.Events["StaticConfigSet"].ID:
		return _CCVAggregator.ParseStaticConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
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

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64) (uint8, error)

	GetSourceChainConfig(opts *bind.CallOpts, sourceChainSelector uint64) (CCVAggregatorSourceChainConfig, error)

	GetStaticConfig(opts *bind.CallOpts) (CCVAggregatorStaticConfig, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplySourceChainConfigUpdates(opts *bind.TransactOpts, sourceChainConfigUpdates []CCVAggregatorSourceChainConfigArgs) (*types.Transaction, error)

	Execute(opts *bind.TransactOpts, report CCVAggregatorAggregatedReport) (*types.Transaction, error)

	ExecuteSingleMessage(opts *bind.TransactOpts, message InternalAny2EVMMessage) (*types.Transaction, error)

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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
