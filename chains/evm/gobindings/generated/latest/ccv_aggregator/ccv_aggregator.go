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

type CCVAggregatorAggregatedReport struct {
	Message InternalAny2EVMMessage
	Ccvs    []common.Address
	CcvData [][]byte
}

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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.AggregatedReport\",\"components\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structInternal.Any2EVMMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isPoolCCV\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f61425e38819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a160405161400a9081610254823960805181818161012e01526110fa015260a051818181610191015261105e015260c0518181816101cd01528181612c540152613293015260e0518181816101550152611bc70152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100c7578063181f5a77146100c257806345ec4b5f146100bd5780635215505b146100b857806379ba5097146100b35780637ce1552a146100ae5780638da5cb5b146100a9578063e9d68a8e146100a4578063f054ac571461009f578063f2fde38b1461009a5763f734ef0e1461009557600080fd5b610f1b565b610e27565b610dac565b610cd0565b610c6d565b610bca565b610a8b565b610938565b610767565b61040a565b6100dc565b60009103126100d757565b600080fd5b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75761011361172f565b506102496040516101238161027c565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761029857604052565b61024d565b60a0810190811067ffffffffffffffff82111761029857604052565b6020810190811067ffffffffffffffff82111761029857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761029857604052565b6040519061032560c0836102d5565b565b6040519061032560a0836102d5565b60405190610325610100836102d5565b604051906103256040836102d5565b67ffffffffffffffff811161029857601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6040519061039e6020836102d5565b60008252565b60005b8381106103b75750506000910152565b81810151838201526020016103a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610403815180928187528780880191016103a4565b0116010190565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d757610249604080519061044b81836102d5565b601782527f43435641676772656761746f7220312e372e302d6465760000000000000000006020830152519182916020835260208301906103c7565b6004359067ffffffffffffffff821682036100d757565b359067ffffffffffffffff821682036100d757565b91908260809103126100d7576040516104cb8161027c565b60606104ff818395803585526104e36020820161049e565b60208601526104f46040820161049e565b60408601520161049e565b910152565b81601f820112156100d75780359061051b82610355565b9261052960405194856102d5565b828452602083830101116100d757816000926020809301838601378301015290565b73ffffffffffffffffffffffffffffffffffffffff8116036100d757565b35906103258261054b565b359063ffffffff821682036100d757565b67ffffffffffffffff81116102985760051b60200190565b81601f820112156100d7578035906105b482610585565b926105c260405194856102d5565b82845260208085019360051b830101918183116100d75760208101935b8385106105ee57505050505090565b843567ffffffffffffffff81116100d757820160807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126100d7576040519161063a8361027c565b602082013567ffffffffffffffff81116100d75785602061065d92850101610504565b8352604082013561066d8161054b565b602084015260608201359267ffffffffffffffff84116100d75760808361069b886020809881980101610504565b6040840152013560608201528152019401936105df565b919091610120818403126100d7576106c8610316565b926106d381836104b3565b8452608082013567ffffffffffffffff81116100d757816106f5918401610504565b602085015260a082013567ffffffffffffffff81116100d7578161071a918401610504565b604085015261072b60c08301610569565b606085015261073c60e08301610574565b608085015261010082013567ffffffffffffffff81116100d757610760920161059d565b60a0830152565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d7576107b96107be9136906004016106b2565b611a1a565b005b906020808351928381520192019060005b8181106107de5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016107d1565b6108759173ffffffffffffffffffffffffffffffffffffffff82511681526020820151151560208201526080610864610852604085015160a0604086015260a08501906103c7565b606085015184820360608601526107c0565b9201519060808184039101526107c0565b90565b6040810160408252825180915260206060830193019060005b818110610918575050506020818303910152815180825260208201916020808360051b8301019401926000915b8383106108cd57505050505090565b9091929394602080610909837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08660019603018752895161080a565b970193019301919392906108be565b825167ffffffffffffffff16855260209485019490920191600101610891565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760025461097381610585565b9061098160405192836102d5565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06109ae82610585565b0160005b818110610a745750506109c481611d3d565b9060005b8181106109e057505061024960405192839283610878565b80610a186109ff6109f2600194613ddc565b67ffffffffffffffff1690565b610a09838761182c565b9067ffffffffffffffff169052565b610a58610a53610a39610a2b848861182c565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b611eec565b610a62828761182c565b52610a6d818661182c565b50016109c8565b602090610a7f611d11565b828287010152016109b2565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760005473ffffffffffffffffffffffffffffffffffffffff81163303610b4a577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b60041115610b7e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b906004821015610b7e5752565b6020810192916103259190610bad565b346100d75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d757610c01610487565b6024359067ffffffffffffffff821682036100d7576044359167ffffffffffffffff83116100d757610c3a610c4d933690600401610504565b9060643592610c488461054b565b613043565b600052600560205261024960ff6040600020541660405191829182610bba565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b90602061087592818152019061080a565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75767ffffffffffffffff610d10610487565b610d18611d11565b501660005260046020526102496040600020610d9b600360405192610d3c8461029d565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c1615156020850152604051610d7f81610d788160018601611ddf565b03826102d5565b6040850152610d9060028201611ed1565b606085015201611ed1565b608082015260405191829182610cbf565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d757366023820112156100d757806004013567ffffffffffffffff81116100d7573660248260051b840101116100d75760246107be9201611f77565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75773ffffffffffffffffffffffffffffffffffffffff600435610e778161054b565b610e7f6130c4565b16338114610ef157807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d757806004019060607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100d75760015460a01c60ff1661170557610fde740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b610ff8610ff3610fee8480612671565b6126a4565b61310f565b80516020908101516040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608083901b16600482015267ffffffffffffffff9091169491816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa908115611368576000916116d6575b5061169e576110db6110d76110cd8667ffffffffffffffff166000526004602052604060002090565b5460a01c60ff1690565b1590565b6116665781516040015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff82160361162f57506024830190604461113c83836126c4565b959050019361114b85836126c4565b919050036115e35782516060015167ffffffffffffffff16936111976020850151956060860196611190885173ffffffffffffffffffffffffffffffffffffffff1690565b9189613043565b946111b66111af876000526005602052604060002090565b5460ff1690565b936111c085610b74565b841580156115d0575b15611579576111d6612718565b9060a08701805151806114c7575b50506112319261122961122261120760208b510167ffffffffffffffff90511690565b925173ffffffffffffffffffffffffffffffffffffffff1690565b92876126c4565b929091613438565b91906040516020810190611275816112498a8561289b565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102d5565b5190209760005b825181101561136d576112cb6112b26112b2611298848761182c565b5173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906112ea6112d986896126c4565b6112e3848a61182c565b51916128ac565b9290813b156100d7576000918d838d6113338e604051998a97889687957f01e0fdea0000000000000000000000000000000000000000000000000000000087526004870161290c565b03925af19182156113685760019261134d575b500161127c565b8061135c6000611362936102d5565b806100cc565b38611346565b6119fd565b6113d4828a8a6113b561138a836000526005602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6113d96113c1826139da565b9581946000526005602052604060002090565b612976565b6113e282610b74565b60028203611473575b67ffffffffffffffff7f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2915161144361142f606083015167ffffffffffffffff1690565b9151968360405194859416971695836129d0565b0390a46107be7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b61147c82610b74565b60038203156113eb5751606001516114c3929067ffffffffffffffff16907f926c5a3e000000000000000000000000000000000000000000000000000000006000526129ad565b6000fd5b909192506001810361154c5750906115448261150860206114ec61123197965161181a565b51015173ffffffffffffffffffffffffffffffffffffffff1690565b9061152060208b510167ffffffffffffffff90511690565b604061153b6060611531855161181a565b510151935161181a565b5101519261322a565b9192386111e4565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6114c388611594606089510167ffffffffffffffff90511690565b7f3b5754190000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff91821660045216602452604490565b506115da85610b74565b600385146111c9565b6115fc846115f46114c394846126c4565b9390506126c4565b7fb5ace4f30000000000000000000000000000000000000000000000000000000060005260049290925250602452604490565b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6116f8915060203d6020116116fe575b6116f081836102d5565b8101906126af565b386110a4565b503d6116e6565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061173c8261027c565b60006060838281528260208201528260408201520152565b604051906040820182811067ffffffffffffffff8211176102985760405260006020838281520152565b9061178882610585565b61179560405191826102d5565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06117c38294610585565b019060005b8281106117d457505050565b6020906117df611754565b828285010152016117c8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8051156118275760200190565b6117eb565b80518210156118275760209160051b010190565b801515036100d757565b90916060828403126100d757815161186181611840565b92602083015167ffffffffffffffff81116100d75783019080601f830112156100d75781519161189083610355565b9161189e60405193846102d5565b838352602084830101116100d7576040926118bf91602080850191016103a4565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a0840152608061193c611908604084015160a060c08801526101208701906103c7565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103c7565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106119c55750505061ffff909516602083015261032592916060916119a99063ffffffff166040830152565b019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff168552602090810151818601526040909401939092019160010161197b565b6040513d6000823e3d90fd5b9060206108759281815201906103c7565b90303303611ce75760a0820191611a3283515161177e565b9160005b84518051821015611aac5790611a90611a518260019461182c565b516020860151606087015173ffffffffffffffffffffffffffffffffffffffff1690611a8a602089510167ffffffffffffffff90511690565b92612bc6565b611a9a828761182c565b52611aa5818661182c565b5001611a36565b50509250805191611aca6020845194015167ffffffffffffffff1690565b906020830151916040840192611af8845192611ae4610327565b97885267ffffffffffffffff166020880152565b60408601526060850152608084015251511580611cc9575b8015611ca6575b8015611c74575b611c7057611bf09181611b636112b2611b49610a396020600097510167ffffffffffffffff90511690565b5473ffffffffffffffffffffffffffffffffffffffff1690565b9083611b976060611b7b608085015163ffffffff1690565b93015173ffffffffffffffffffffffffffffffffffffffff1690565b93604051968795869485937f3cf979830000000000000000000000000000000000000000000000000000000085527f000000000000000000000000000000000000000000000000000000000000000090600486016118c5565b03925af190811561136857600090600092611c49575b5015611c0f5750565b611c45906040519182917f0a8d6e8c00000000000000000000000000000000000000000000000000000000835260048301611a09565b0390fd5b9050611c6891503d806000833e611c6081836102d5565b81019061184a565b509038611c06565b5050565b50611ca16110d7611c9c606084015173ffffffffffffffffffffffffffffffffffffffff1690565b612f45565b611b1e565b50606081015173ffffffffffffffffffffffffffffffffffffffff163b15611b17565b5063ffffffff611ce0608083015163ffffffff1690565b1615611b10565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611d1e8261029d565b6060608083600081526000602082015282604082015282808201520152565b90611d4782610585565b611d5460405191826102d5565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611d828294610585565b0190602036910137565b90600182811c92168015611dd5575b6020831014611da657565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611d9b565b60009291815491611def83611d8c565b8083529260018116908115611e455750600114611e0b57505050565b60009081526020812093945091925b838310611e2b575060209250010190565b600181602092949394548385870101520191019190611e1a565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b906020825491828152019160005260206000209060005b818110611ea55750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201611e98565b90610325611ee59260405193848092611e81565b03836102d5565b9060036080604051611efd8161029d565b611f73819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c1615156020840152604051611f3e81610d788160018601611ddf565b6040840152604051611f5781610d788160028601611e81565b6060840152611f6c6040518096819301611e81565b03846102d5565b0152565b90611f806130c4565b60005b818110611f8f57505050565b611fa2611f9d8284866121ec565b61229e565b9060208201611fbc6109f2825167ffffffffffffffff1690565b156121c257611fe56112b26112b2855173ffffffffffffffffffffffffffffffffffffffff1690565b1580156121b5575b61217557606083019081518051801591821561219f575b50506121755761216c61215f6109f28361213360a0600199612117612055610a397f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259a5167ffffffffffffffff1690565b996120ac6120666040850151151590565b8c547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178c55565b61210e6120cd845173ffffffffffffffffffffffffffffffffffffffff1690565b8c9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518c8b016123d0565b612128608082015160028b0161252a565b01516003880161252a565b61215061214b6109f2835167ffffffffffffffff1690565b613e11565b505167ffffffffffffffff1690565b92604051918291826125be565b0390a201611f83565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6020012090506121ad612353565b143880612004565b5060808301515115611fed565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156118275760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100d7570190565b359061032582611840565b9080601f830112156100d757813561224e81610585565b9261225c60405194856102d5565b81845260208085019260051b8201019283116100d757602001905b8282106122845750505090565b6020809183356122938161054b565b815201910190612277565b60c0813603126100d7576122b0610316565b906122ba81610569565b82526122c86020820161049e565b60208301526122d96040820161222c565b6040830152606081013567ffffffffffffffff81116100d7576122ff9036908301610504565b6060830152608081013567ffffffffffffffff81116100d7576123259036908301612237565b608083015260a08101359067ffffffffffffffff82116100d75761234b91369101612237565b60a082015290565b6040516020810190600082526020815261236e6040826102d5565b51902090565b81811061237f575050565b60008155600101612374565b9190601f811161239a57505050565b610325926000526020600020906020601f840160051c830193106123c6575b601f0160051c0190612374565b90915081906123b9565b919091825167ffffffffffffffff8111610298576123f8816123f28454611d8c565b8461238b565b6020601f821160011461245657819061244793949560009261244b575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612415565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169061248984600052602060002090565b9160005b8181106124e3575095836001959697106124ac575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880806124a2565b9192602060018192868b01518155019401920161248d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81519167ffffffffffffffff8311610298576801000000000000000083116102985760209082548484558085106125a1575b500190600052602060002060005b8381106125775750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161256a565b6125b8908460005285846000209182019101612374565b3861255c565b6003610875926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a0606084015261264061260d60c0850160018401611ddf565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301608086015260028301611e81565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08286030191015201611e81565b9035907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1813603018212156100d7570190565b6108759036906106b2565b908160209103126100d7575161087581611840565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100d7570180359067ffffffffffffffff82116100d757602001918160051b360383136100d757565b604051906127276020836102d5565b6000808352366020840137565b9080602083519182815201916020808360051b8301019401926000915b83831061276057505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085600195030186528851906060806127de6127ae85516080865260808601906103c7565b73ffffffffffffffffffffffffffffffffffffffff878701511687860152604086015185820360408701526103c7565b93015191015297019301930191939290612751565b6108759167ffffffffffffffff6060835180518452826020820151166020850152826040820151166040850152015116606082015260a0612859612848602085015161012060808601526101208501906103c7565b6040850151848203848601526103c7565b606084015173ffffffffffffffffffffffffffffffffffffffff1660c084015292608081015163ffffffff1660e0840152015190610100818403910152612734565b9060206108759281815201906127f3565b91908110156118275760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100d757019081359167ffffffffffffffff83116100d75760200182360381136100d7579190565b959492917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f866103259760609661294e60209760808e81815201906127f3565b95878d01528b860360408d015281865286860137600085828601015201160101940190610bad565b906004811015610b7e5760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b929160449067ffffffffffffffff61032593816064971660045216602452610bad565b806129e16040926108759594610bad565b81602082015201906103c7565b908160209103126100d757516108758161054b565b6040519061039e826102b9565b908160209103126100d75760405190612a28826102b9565b51815290565b90610875916020815260e0612b23612af0612a57855161010060208701526101208601906103c7565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff16606086015260608601516080860152612abc608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103c7565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103c7565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103c7565b3d15612b82573d90612b6882610355565b91612b7660405193846102d5565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff610875949316815281602082015201906103c7565b91908203918211612bc157565b6124fb565b909291612bd1611754565b50602082015173ffffffffffffffffffffffffffffffffffffffff166040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa94851561136857600095612f14575b5073ffffffffffffffffffffffffffffffffffffffff8516948515612ed157612cad81612f9b565b612ed157612cbd6110d782612fc5565b612ed15750612d8f91602091612cd38886613b9e565b95612cdc612a03565b506060810151612d086040835193015193612cf5610336565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c0820152612d5761038f565b60e0820152604051809381927f3907753700000000000000000000000000000000000000000000000000000000835260048301612a2e565b03816000885af160009181612ea0575b50612de35784612dad612b57565b90611c456040519283927f9fe2f95a00000000000000000000000000000000000000000000000000000000845260048401612b87565b84909373ffffffffffffffffffffffffffffffffffffffff831603612e36575b50505051612e2e612e12610346565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b612e3f91613b9e565b908082108015612e8c575b612e545783612e03565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b50612e978183612bb4565b83511415612e4a565b612ec391925060203d602011612eca575b612ebb81836102d5565b810190612a10565b9038612d9f565b503d612eb1565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b612f3791955060203d602011612f3e575b612f2f81836102d5565b8101906129ee565b9338612c85565b503d612f25565b612f6f7f85572ffb0000000000000000000000000000000000000000000000000000000082613d7a565b9081612f89575b81612f7f575090565b6108759150613d1a565b9050612f9481613c54565b1590612f76565b612f6f7ff208a58f0000000000000000000000000000000000000000000000000000000082613d7a565b612f6f7faff2afbf0000000000000000000000000000000000000000000000000000000082613d7a565b612f6f7f05c7a8d00000000000000000000000000000000000000000000000000000000082613d7a565b612f6f7f7909b5490000000000000000000000000000000000000000000000000000000082613d7a565b929061236e9173ffffffffffffffffffffffffffffffffffffffff61309167ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906103c7565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102d5565b73ffffffffffffffffffffffffffffffffffffffff6001541633036130e557565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160c0810181811067ffffffffffffffff8211176102985760609160a09160405261313a61172f565b815282602082015282604082015260008382015260006080820152015290565b9080601f830112156100d757815161317181610585565b9261317f60405194856102d5565b81845260208085019260051b8201019283116100d757602001905b8282106131a75750505090565b6020809183516131b68161054b565b81520191019061319a565b906020828203126100d757815167ffffffffffffffff81116100d757610875920161315a565b909267ffffffffffffffff60809373ffffffffffffffffffffffffffffffffffffffff6108759796168452166020830152604082015281606082015201906103c7565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526060949293916020828060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa918215611368576000926133a5575b506132ce82612fef565b61330d575b505050508151156132e2575090565b610875915061330760029167ffffffffffffffff166000526004602052604060002090565b01611ed1565b8495509373ffffffffffffffffffffffffffffffffffffffff60009495613363604051978896879586947f0ba375f9000000000000000000000000000000000000000000000000000000008652600486016131e7565b0392165afa90811561136857600091613382575b5090388080806132d3565b61339f91503d806000833e61339781836102d5565b8101906131c1565b38613377565b6133bf91925060203d602011612f3e57612f2f81836102d5565b90386132c4565b91908110156118275760051b0190565b356108758161054b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114612bc15760010190565b8015612bc1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b949160039061346e61344f61330797969489613ef1565b949198909967ffffffffffffffff166000526004602052604060002090565b9061347884611d3d565b9261348285611d3d565b9661348c86611d3d565b96600094855b8b518110156135f7576000805b8d8b848a828510613515575b5050505050156134bd57600101613492565b6134cd6112986114c3928e61182c565b7fbd76195f000000000000000000000000000000000000000000000000000000006000529073ffffffffffffffffffffffffffffffffffffffff604492166004526000602452565b61129861354d926135476135428873ffffffffffffffffffffffffffffffffffffffff976112b2966133c6565b6133d6565b9561182c565b91161461355c5760010161349f565b60019891508b8d8b8d613578613572868661182c565b51151590565b6135e75792846135cb876135d1958f84966135ab846135a56135426135dc9f9e6135d79e6135c6976133c6565b9261182c565b9073ffffffffffffffffffffffffffffffffffffffff169052565b61182c565b5261182c565b60019052565b6133e0565b96388d8b848a6134ab565b9150508f9150848a9b949b6134ab565b50919950919960005b8a51811015613725576000805b8a8d8b858a828610613684575b5050505050501561362d57600101613600565b61363d6112986114c3928d61182c565b7fbd76195f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff166004526001602452604490565b6112986136b1926135476135428973ffffffffffffffffffffffffffffffffffffffff976112b2966133c6565b9116146136c1575060010161360d565b9791508b898b6136d66135728560019d61182c565b613715579280613702866135d7956135c66135d1968f846135a5613542896137099f9e6135ab956133c6565b528d61182c565b95388a8d8b858a61361a565b8f91508d9250858a9b959b61361a565b50919990985060005b895181101561380d576000805b898c8a858982861061376b575b5050505050501561375b5760010161372e565b61363d6112986114c3928c61182c565b611298613798926135476135428973ffffffffffffffffffffffffffffffffffffffff976112b2966133c6565b9116146137a8575060010161373b565b9691508a888a6137bd6135728560019c61182c565b6137fd5792806137ea866135d7956135c66135d1966135ab8f9b6135a5613542896137f19f9e849f6133c6565b528c61182c565b9438898c8a8589613748565b8c92508e915085899a959a613748565b509098975095929594919493909360ff811690816000995b8b518b10156139065760005b8c868210806138fd575b156138f05773ffffffffffffffffffffffffffffffffffffffff61386f6112b28f611298906135478f888e613542926133c6565b9116146138845761387f906133e0565b613831565b93613894909b999193929b61340d565b938a6138a3613572838a61182c565b6138e557916135d76135d183806138d3858f8f9060019a8f846135a56135426138da9e6135ab946135c6976133c6565b528a61182c565b985b01999091613825565b5050976001906138dc565b50509190996001906138dc565b5085151561383b565b995093509450929750508161392e575050508151810361392557509190565b80825283529190565b6114c3929161393c91612bb4565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b90602082519201517fffffffff00000000000000000000000000000000000000000000000000000000811692600481106139a8575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b303b156100d7576000613a1a91604051809381927f45ec4b5f0000000000000000000000000000000000000000000000000000000083526004830161289b565b038183305af19081613b89575b50613b7e57613a34612b57565b9072c11c11c11c11c11c11c11c11c11c11c11c11c13314613a56575b60039190565b613a87613a6283613970565b7fffffffff000000000000000000000000000000000000000000000000000000001690565b7f37c3be2900000000000000000000000000000000000000000000000000000000148015613b4a575b8015613b16575b15613a50576114c3613ac883613970565b7f2882569d000000000000000000000000000000000000000000000000000000006000527fffffffff0000000000000000000000000000000000000000000000000000000016600452602490565b50613b23613a6283613970565b7fea7f4b120000000000000000000000000000000000000000000000000000000014613ab7565b50613b57613a6283613970565b7fafa32a2c0000000000000000000000000000000000000000000000000000000014613ab0565b60029061087561038f565b8061135c6000613b98936102d5565b38613a27565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa60009181613c1d575b50613c195782612dad612b57565b9150565b90916020823d602011613c4c575b81613c38602093836102d5565b81010312613c495750519038613c0b565b80fd5b3d9150613c2b565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff00000000000000000000000000000000000000000000000000000000602483015260248252613cb46044836102d5565b6179185a10613cf0576020926000925191617530fa6000513d82613ce4575b5081613cdd575090565b9050151590565b60201115915038613cd3565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a700000000000000000000000000000000000000000000000000000000602483015260248252613cb46044836102d5565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a700000000000000000000000000000000000000000000000000000000855216602483015260248252613cb46044836102d5565b6002548110156118275760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b80600052600360205260406000205415600014613e8f57600254680100000000000000008110156102985760018101600255600060025482101561182757600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b50600090565b90916060828403126100d757815167ffffffffffffffff81116100d75783613ebe91840161315a565b92602083015167ffffffffffffffff81116100d757604091613ee191850161315a565b92015160ff811681036100d75790565b919091613f15610a538267ffffffffffffffff166000526004602052604060002090565b90833b613f32575b50606001519150613f2c612718565b90600090565b613f3b84613019565b15613f1d576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa801561136857600080948192613fd5575b50805115801590613fc9575b613fc25750613f1d565b9392909150565b5060ff82161515613fb8565b9150613ff49294503d8091833e613fec81836102d5565b810190613e95565b90939138613fac56fea164736f6c634300081a000a",
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

	GetExecutionState(opts *bind.CallOpts, sourceChainSelector uint64, sequenceNumber uint64, sender []byte, receiver common.Address) (uint8, error)

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

	Address() common.Address
}
