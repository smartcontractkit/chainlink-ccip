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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applySourceChainConfigUpdates\",\"inputs\":[{\"name\":\"sourceChainConfigUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfigArgs[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCV\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSingleMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structMessageV1Codec.MessageV1\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destBlob\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structMessageV1Codec.TokenTransferV1[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ccvs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ccvData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllSourceChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCCVAggregator.SourceChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumInternal.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVAggregator.StaticConfig\",\"components\":[{\"name\":\"localChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySelfCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExecutionError\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasToCompleteTx\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidCCVDataLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enumMessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidEVMAddress\",\"inputs\":[{\"name\":\"encodedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEncodingVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageDestChainSelector\",\"inputs\":[{\"name\":\"messageDestChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNewState\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newState\",\"type\":\"uint8\",\"internalType\":\"enumInternal.MessageExecutionState\"}]},{\"type\":\"error\",\"name\":\"InvalidNumberOfTokens\",\"inputs\":[{\"name\":\"numTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierSelector\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotACompatiblePool\",\"inputs\":[{\"name\":\"notPool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OptionalCCVQuorumNotReached\",\"inputs\":[{\"name\":\"wanted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReceiverError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReleaseOrMintBalanceMismatch\",\"inputs\":[{\"name\":\"amountReleased\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePre\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balancePost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RequiredCCVMissing\",\"inputs\":[{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isPoolCCV\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"SkippedAlreadyExecutedMessage\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"SourceChainNotEnabled\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenHandlingError\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainSelectorNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101006040523461023857604051601f6156d938819003918201601f19168301916001600160401b0383118484101761023d578084926080946040528339810103126102385760405190600090608083016001600160401b038111848210176102245760405280516001600160401b0381168103610220578352602081015161ffff8116810361022057602084019081526040820151916001600160a01b038316830361021c576040850192835260600151926001600160a01b03841684036102195760608501938452331561020a57600180546001600160a01b0319163317905582516001600160a01b03161580156101f8575b6101e95784516001600160401b0316156101da5784516001600160401b03908116608090815284516001600160a01b0390811660a0528651811660c052845161ffff90811660e052604080518a51909516855286519091166020850152865182169084015286511660608301527f4811b8f4a862be218e56fe7f80b1142a234944c26d69ce7010eb3519622e495091a1604051615485908161025482396080518181816101460152611a35015260a0518181816101a901526119b8015260c0518181816101d101528181613cfb0152614603015260e05181818161016d0152612a160152f35b63c656089560e01b8152600490fd5b6342bcdf7f60e11b8152600490fd5b5083516001600160a01b0316156100f4565b639b15e16f60e01b8152600490fd5b80fd5b8380fd5b8280fd5b634e487b7160e01b83526041600452602483fd5b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001257600080fd5b60003560e01c806306285c69146100c7578063181f5a77146100c25780635215505b146100bd5780636b8be52c146100b857806379ba5097146100b35780637ce1552a146100ae5780638da5cb5b146100a9578063d2b33733146100a4578063e9d68a8e1461009f578063f054ac571461009a5763f2fde38b1461009557600080fd5b610fa9565b610cbc565b610bdc565b610b09565b610ab7565b610a18565b610847565b610787565b610603565b61040e565b6100dc565b60009103126100d757565b600080fd5b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d7576000606060405161011b81610280565b828152826020820152826040820152015261024d60405161013b81610280565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815261ffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000811660408301527f000000000000000000000000000000000000000000000000000000000000000016606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff81608084019567ffffffffffffffff815116855261ffff6020820151166020860152826040820151166040860152015116910152565b0390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761029c57604052565b610251565b60a0810190811067ffffffffffffffff82111761029c57604052565b6020810190811067ffffffffffffffff82111761029c57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761029c57604052565b6040519061032960a0836102d9565b565b6040519061032960c0836102d9565b60405190610329610100836102d9565b604051906103296040836102d9565b67ffffffffffffffff811161029c57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906103a26020836102d9565b60008252565b60005b8381106103bb5750506000910152565b81810151838201526020016103ab565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093610407815180928187528780880191016103a8565b0116010190565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75761024d604080519061044f81836102d9565b601782527f43435641676772656761746f7220312e372e302d6465760000000000000000006020830152519182916020835260208301906103cb565b906020808351928381520192019060005b8181106104a95750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161049c565b6105409173ffffffffffffffffffffffffffffffffffffffff8251168152602082015115156020820152608061052f61051d604085015160a0604086015260a08501906103cb565b6060850151848203606086015261048b565b92015190608081840391015261048b565b90565b6040810160408252825180915260206060830193019060005b8181106105e3575050506020818303910152815180825260208201916020808360051b8301019401926000915b83831061059857505050505090565b90919293946020806105d4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516104d5565b97019301930191939290610589565b825167ffffffffffffffff1685526020948501949092019160010161055c565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760025461063e8161109d565b9061064c60405192836102d9565b8082527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06106798261109d565b0160005b81811061073f57505061068f816110e1565b9060005b8181106106ab57505061024d60405192839283610543565b806106e36106ca6106bd600194614b91565b67ffffffffffffffff1690565b6106d48387611171565b9067ffffffffffffffff169052565b61072361071e6107046106f68488611171565b5167ffffffffffffffff1690565b67ffffffffffffffff166000526004602052604060002090565b6112e5565b61072d8287611171565b526107388186611171565b5001610693565b60209061074a6110b5565b8282870101520161067d565b9181601f840112156100d75782359167ffffffffffffffff83116100d7576020808501948460051b0101116100d757565b346100d75760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d757366023820112156100d75780600401359067ffffffffffffffff82116100d75736602483830101116100d75760243567ffffffffffffffff81116100d757610814903690600401610756565b604435929167ffffffffffffffff84116100d7576108459461083c6024953690600401610756565b95909401611885565b005b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760005473ffffffffffffffffffffffffffffffffffffffff81163303610906577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff8116036100d757565b359061032982610930565b92919261095982610359565b9161096760405193846102d9565b8294818452818301116100d7578281602093846000960137010152565b9080601f830112156100d7578160206105409335910161094d565b73ffffffffffffffffffffffffffffffffffffffff8116036100d757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600411156109f657565b6109bd565b9060048210156109f65752565b60208101929161032991906109fb565b346100d75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d757600435610a5381610930565b60243590610a6082610930565b6044359167ffffffffffffffff83116100d757610a84610a97933690600401610984565b9060643592610a928461099f565b613a70565b600052600560205261024d60ff6040600020541660405191829182610a08565b346100d75760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346100d75760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d7576101607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126100d75760443567ffffffffffffffff81116100d757610b97903690600401610756565b916064359267ffffffffffffffff84116100d757610bbc610845943690600401610756565b9390926024359060040161279b565b9060206105409281815201906104d5565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75767ffffffffffffffff600435610c2081610930565b610c286110b5565b5016600052600460205261024d6040600020610cab600360405192610c4c846102a1565b60ff815473ffffffffffffffffffffffffffffffffffffffff8116865260a01c1615156020850152604051610c8f81610c8881600186016111d8565b03826102d9565b6040850152610ca0600282016112ca565b6060850152016112ca565b608082015260405191829182610bcb565b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75760043567ffffffffffffffff81116100d757610d0b903690600401610756565b90610d146149e1565b60005b828110610d2057005b610d33610d2e828585612c77565b612d34565b906020820191610d4e6106bd845167ffffffffffffffff1690565b15610f7f57610d90610d77610d77835173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b158015610f72575b610f32576060810190815180518015918215610f5c575b5050610f3257610f29610f1c6106bd600196610ef085610ee6610e7898610edc60807f04a080dee5faf023415dfb59e1b260d185fcfa4b5a56ce9d24f42312927e4e259a0191610ed3610e92845196610e0e60a0820198895190614a2c565b610e236107048b5167ffffffffffffffff1690565b9e8f610e326040840151151590565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055565b5173ffffffffffffffffffffffffffffffffffffffff1690565b8d9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b518d8c01612e66565b5160028a01612f91565b5160038801612f91565b610f0d610f086106bd835167ffffffffffffffff1690565b6153f4565b505167ffffffffffffffff1690565b9260405191829182613025565b0390a201610d17565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b602001209050610f6a612de9565b143880610daf565b5060808101515115610d98565b7fc65608950000000000000000000000000000000000000000000000000000000060005260046000fd5b346100d75760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100d75773ffffffffffffffffffffffffffffffffffffffff600435610ff98161099f565b6110016149e1565b1633811461107357807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff811161029c5760051b60200190565b604051906110c2826102a1565b6060608083600081526000602082015282604082015282808201520152565b906110eb8261109d565b6110f860405191826102d9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611126829461109d565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80511561116c5760200190565b611130565b805182101561116c5760209160051b010190565b90600182811c921680156111ce575b602083101461119f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611194565b600092918154916111e883611185565b808352926001811690811561123e575060011461120457505050565b60009081526020812093945091925b838310611224575060209250010190565b600181602092949394548385870101520191019190611213565b905060209495507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091509291921683830152151560051b010190565b906020825491828152019160005260206000209060005b81811061129e5750505090565b825473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019283019201611291565b906103296112de926040519384809261127a565b03836102d9565b90600360806040516112f6816102a1565b61136c819560ff815473ffffffffffffffffffffffffffffffffffffffff8116855260a01c161515602084015260405161133781610c8881600186016111d8565b604084015260405161135081610c88816002860161127a565b6060840152611365604051809681930161127a565b03846102d9565b0152565b801515036100d757565b908160209103126100d7575161054081611370565b6040513d6000823e3d90fd5b9060206105409281815201906103cb565b90602082519201517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106113e4575050565b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000929350829060140360031b1b161690565b9060048110156109f65760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b9080602083519182815201916020808360051b8301019401926000915b83831061147957505050505090565b9091929394602080611509837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260806114f86114e66114d48786015160a08987015260a08601906103cb565b604086015185820360408701526103cb565b606085015184820360608601526103cb565b9201519060808184039101526103cb565b9701930193019193929061146a565b9160209082815201919060005b8181106115325750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561155b8161099f565b168152019401929101611525565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100d757016020813591019167ffffffffffffffff82116100d75781360383136100d757565b90602083828152019260208260051b82010193836000925b8484106116205750505050505090565b909192939495602080611666837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030188526116608b886115a8565b90611569565b9801940194019294939190611610565b94929361054096946118466118599493608089526116a160808a01825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660a08a0152604081015167ffffffffffffffff1660c08a01526101406118136117dd6117a76117708d61172c6116f7606089015161016060e08501526101e08401906103cb565b60808901517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848303016101008501526103cb565b60a088015161ffff166101208301529060c088015190877fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103cb565b8d60e0870151906101607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80828503019101526103cb565b6101008501518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101808f01526103cb565b6101208401518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80016101a08e015261144d565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8303016101c08b01526103cb565b9260208801528683036040880152611518565b9260608185039101526115f8565b8061187860409261054095946109fb565b81602082015201906103cb565b60015491929160a01c60ff16611e6a576118f7906118dd740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6001541617600155565b6118ef6118ea858361336e565b613a64565b93369161094d565b602081519101209461199f602061194461191c6106bd875167ffffffffffffffff1690565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e6557600091611e36575b50611deb57611a17611a13611a09610704865167ffffffffffffffff1690565b5460a01c60ff1690565b1590565b611da057602083015167ffffffffffffffff1667ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001667ffffffffffffffff821603611d695750808403611d375760e083019485516014815103611cfd5750611ac8611a93855167ffffffffffffffff1690565b6040860197611aaa895167ffffffffffffffff1690565b611ac2611abc60c08a015193516113ac565b60601c90565b92613a70565b94611ae7611ae0876000526005602052604060002090565b5460ff1690565b611af0816109ec565b8015908115611ce9575b5015611c7a57611c147f8c324ce1367b83031769f6a813e3bb4c117aba2185789d66b98b791405be6df2956106f6611bf367ffffffffffffffff97611bee8d611bc28d9a611c359a611c0f9f9a611b8c611b61611c239d6000526005602052604060002090565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00825416179055565b6040519687957fd2b337330000000000000000000000000000000000000000000000000000000060208801528b60248801611676565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102d9565b613af1565b969015611c65576002998a916000526005602052604060002090565b611416565b965167ffffffffffffffff1690565b91836040519485941697169583611867565b0390a46103297fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60015416600155565b6003998a916000526005602052604060002090565b611ce58888611ca3611c94895167ffffffffffffffff1690565b915167ffffffffffffffff1690565b917f5e570e5100000000000000000000000000000000000000000000000000000000600052929167ffffffffffffffff80926064956004521660245216604452565b6000fd5b60039150611cf6816109ec565b1438611afa565b611d33906040519182917f8d666f600000000000000000000000000000000000000000000000000000000083526004830161139b565b0390fd5b7fb5ace4f300000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b7f38432a220000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b611ce5611db5845167ffffffffffffffff1690565b7fed053c590000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611ce5611e00845167ffffffffffffffff1690565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b611e58915060203d602011611e5e575b611e5081836102d9565b81019061137a565b386119e9565b503d611e46565b61138f565b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100d7570180359067ffffffffffffffff82116100d7576020019181360383136100d757565b919091357fffffffffffffffffffffffffffffffffffffffff000000000000000000000000811692601481106113e4575050565b60405190611f286020836102d9565b6000808352366020840137565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156100d7570180359067ffffffffffffffff82116100d757602001918160051b360383136100d757565b901561116c578035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100d7570190565b919081101561116c5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156100d7570190565b916020610540938181520191611569565b3561054081610930565b9082101561116c576120349160051b810190611e94565b9091565b359061ffff821682036100d757565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156100d757016020813591019167ffffffffffffffff82116100d7578160051b360383136100d757565b90602083828152019060208160051b85010193836000915b8383106120c25750505050505090565b9091929394957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082820301865286357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61843603018112156100d75760206121a560019386839401908135815261219761218c612171612156612146888701876115a8565b60a08a88015260a0870191611569565b61216360408701876115a8565b908683036040880152611569565b61217e60608601866115a8565b908583036060870152611569565b9260808101906115a8565b916080818503910152611569565b9801960194930191906120b2565b906105409593949273ffffffffffffffffffffffffffffffffffffffff61240192168352608060208401526121fc608084016121ee83610942565b67ffffffffffffffff169052565b61221c61220b60208301610942565b67ffffffffffffffff1660a0850152565b61223c61222b60408301610942565b67ffffffffffffffff1660c0850152565b6123d06123c46123856123466123086122af61227161225e60608901896115a8565b61016060e08d01526101e08c0191611569565b61227e60808901896115a8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808c8403016101008d0152611569565b6122ca6122be60a08901612038565b61ffff166101208b0152565b6122d760c08801886115a8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101408c0152611569565b61231560e08701876115a8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8403016101608b0152611569565b6123546101008601866115a8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80898403016101808a0152611569565b612393610120850185612047565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80888403016101a089015261209a565b916101408101906115a8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80858403016101c0860152611569565b9360408201526060818503910152611569565b604051906040820182811067ffffffffffffffff82111761029c5760405260006020838281520152565b906124488261109d565b61245560405191826102d9565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612483829461109d565b019060005b82811061249457505050565b60209061249f612414565b82828501015201612488565b91909160a0818403126100d7576124c061031a565b9281358452602082013567ffffffffffffffff81116100d757816124e5918401610984565b6020850152604082013567ffffffffffffffff81116100d7578161250a918401610984565b6040850152606082013567ffffffffffffffff81116100d7578161252f918401610984565b6060850152608082013567ffffffffffffffff81116100d7576125529201610984565b6080830152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8f082019182116125b557565b612559565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd12082019182116125b557565b919082039182116125b557565b90916060828403126100d757815161260b81611370565b92602083015167ffffffffffffffff81116100d75783019080601f830112156100d75781519161263a83610359565b9161264860405193846102d9565b838352602084830101116100d75760409261266991602080850191016103a8565b92015190565b9194939290608083528051608084015267ffffffffffffffff60208201511660a084015260806126e66126b2604084015160a060c08801526101208701906103cb565b60608401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808783030160e08801526103cb565b910151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80848203016101008501526020808351928381520192019060005b8181106127635750505061ffff909516602083015261032992916060916040820152019073ffffffffffffffffffffffffffffffffffffffff169052565b8251805173ffffffffffffffffffffffffffffffffffffffff1685526020908101518186015260409094019390920191600101612725565b939194909294303303612c4d576127c1611abc6127bb60e0880188611e94565b90611ee5565b906127ca611f19565b6101208701976127da8989611f35565b9050612b2d575b6127f492846127ef8a612013565b613e83565b97909160005b83518110156128aa57612816610d77610d77610e788488611171565b9061282c612824828d611171565b51878961201d565b9290813b156100d7576000918a838d612875604051988996879586947f58bfa40a00000000000000000000000000000000000000000000000000000000865230600487016121b3565b03925af1918215611e655760019261288f575b50016127fa565b8061289e60006128a4936102d9565b806100cc565b38612888565b509496935094965050506128c86128c18287611f35565b905061243e565b9260005b6128d68388611f35565b905081101561294b578061292f6128f96001936128f3878c611f35565b90611fc2565b866129196129298c61292161291160c0830183611e94565b949092612013565b9536906124ab565b92369161094d565b9061456c565b6129398288611171565b526129448187611171565b50016128cc565b509490509290926101408101926129628483611e94565b90501580612b25575b8015612b1c575b8015612b0a575b612b0357612a0a6000946129f76129fe6129b5610d7761299b61070489612013565b5473ffffffffffffffffffffffffffffffffffffffff1690565b956129bf81612013565b936129d86129d060c0840184611e94565b929093611e94565b9490956129e361031a565b9a8b5267ffffffffffffffff1660208b0152565b369161094d565b6040870152369161094d565b606084015260808301527f000000000000000000000000000000000000000000000000000000000000000083612a4d612a485a61ffff8516906125e7565b612588565b93612a87604051978896879586947f3cf979830000000000000000000000000000000000000000000000000000000086526004860161266f565b03925af1908115611e6557600090600092612adc575b5015612aa65750565b611d33906040519182917f0a8d6e8c0000000000000000000000000000000000000000000000000000000083526004830161139b565b9050612afb91503d806000833e612af381836102d9565b8101906125f4565b509038612a9d565b5050505050565b50612b17611a13866148e3565b612979565b50843b15612972565b50600061296b565b90506001612b3b8989611f35565b905003612c13576014612b64612b5a612b548b8b611f35565b90611f89565b6060810190611e94565b905003612bcd57906127f491612bc5612b89611abc6127bb612b5a612b548e8e611f35565b898b612bbf6129f7612bb5612b54612ba086612013565b94612bae612b548289611f35565b3596611f35565b6080810190611e94565b92613c92565b9192506127e1565b612bdd612b5a612b548a8a611f35565b90611d336040519283927f8d666f6000000000000000000000000000000000000000000000000000000000845260048401612002565b611ce5612c208989611f35565b7f83d526690000000000000000000000000000000000000000000000000000000060005260045250602490565b7f371a73280000000000000000000000000000000000000000000000000000000060005260046000fd5b919081101561116c5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41813603018212156100d7570190565b35906103298261099f565b359061032982611370565b9080601f830112156100d7578135612ce48161109d565b92612cf260405194856102d9565b81845260208085019260051b8201019283116100d757602001905b828210612d1a5750505090565b602080918335612d298161099f565b815201910190612d0d565b60c0813603126100d757612d4661032b565b90612d5081612cb7565b8252612d5e60208201610942565b6020830152612d6f60408201612cc2565b6040830152606081013567ffffffffffffffff81116100d757612d959036908301610984565b6060830152608081013567ffffffffffffffff81116100d757612dbb9036908301612ccd565b608083015260a08101359067ffffffffffffffff82116100d757612de191369101612ccd565b60a082015290565b60405160208101906000825260208152612e046040826102d9565b51902090565b818110612e15575050565b60008155600101612e0a565b9190601f8111612e3057505050565b610329926000526020600020906020601f840160051c83019310612e5c575b601f0160051c0190612e0a565b9091508190612e4f565b919091825167ffffffffffffffff811161029c57612e8e81612e888454611185565b84612e21565b6020601f8211600114612eec578190612edd939495600092612ee1575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b015190503880612eab565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0821690612f1f84600052602060002090565b9160005b818110612f7957509583600195969710612f42575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080612f38565b9192602060018192868b015181550194019201612f23565b81519167ffffffffffffffff831161029c5768010000000000000000831161029c576020908254848455808510613008575b500190600052602060002060005b838110612fde5750505050565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501612fd1565b61301f908460005285846000209182019101612e0a565b38612fc3565b6003610540926020835260ff815473ffffffffffffffffffffffffffffffffffffffff8116602086015260a01c161515604084015260a060608401526130a761307460c08501600184016111d8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08582030160808601526002830161127a565b9260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101520161127a565b60405190610160820182811067ffffffffffffffff82111761029c576040526060610140836000815260006020820152600060408201528280820152826080820152600060a08201528260c08201528260e082015282610100820152826101208201520152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146125b55760010190565b901561116c5790565b9082101561116c570190565b90600882018092116125b557565b90600282018092116125b557565b90600182018092116125b557565b90602082018092116125b557565b919082018092116125b557565b909392938483116100d75784116100d7578101920390565b919091357fffffffffffffffff00000000000000000000000000000000000000000000000081169260088110613212575050565b7fffffffffffffffff000000000000000000000000000000000000000000000000929350829060080360031b1b161690565b919091357fffff00000000000000000000000000000000000000000000000000000000000081169260028110613278575050565b7fffff000000000000000000000000000000000000000000000000000000000000929350829060020360031b1b161690565b604051906132b7826102a1565b60606080836000815282602082015282604082015282808201520152565b604080519091906132e683826102d9565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b82811061331d57505050565b6020906133286132aa565b82828501015201613311565b604051906133436020836102d9565b600080835282815b82811061335757505050565b6020906133626132aa565b8282850101520161334b565b906133776130d8565b5060258110613a33576133886130d8565b9160006133c76133c161339b858561316c565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b600160ff821603613a05575061340c6133fe6133f86133f26133e96001613181565b600188886131c6565b906131de565b60c01c90565b67ffffffffffffffff168552565b61347d61344f61341c6001613181565b61344a6134396133f86133f261343185613181565b858b8b6131c6565b67ffffffffffffffff166020890152565b613181565b61344a61346c6133f86133f261346485613181565b858a8a6131c6565b67ffffffffffffffff166040880152565b6134a061349a6133c161339b6134928561313f565b948888613175565b60ff1690565b90846134ac83836131b9565b116139d85790816134ce6129f76134c6846134d8966131b9565b8389896131c6565b60608801526131b9565b838110156139aa576134f561349a6133c161339b6134928561313f565b908461350183836131b9565b1161397d57908161351b6129f76134c684613525966131b9565b60808801526131b9565b8361352f8261318f565b11613950578061356461355961355361354d6134646135699661318f565b90613244565b60f01c90565b61ffff1660a0880152565b61318f565b838110156139225761358661349a6133c161339b6134928561313f565b908461359283836131b9565b116138f55790816135ac6129f76134c6846135b6966131b9565b60c08801526131b9565b838110156138c7576135d361349a6133c161339b6134928561313f565b90846135df83836131b9565b1161389a5790816135f96129f76134c684613603966131b9565b60e08801526131b9565b8361360d8261318f565b1161386c5761ffff61363861363261355361354d61362a8661318f565b868a8a6131c6565b9261318f565b9116908461364683836131b9565b1161383f5790816136606129f76134c68461366b966131b9565b6101008801526131b9565b90836136768361318f565b11613812575061ffff61369c61363261355361354d6136948661318f565b8689896131c6565b9116806137a957506136ac613334565b6101208501525b826136bd8261318f565b1161377a5761ffff6136da61363261355361354d6136948661318f565b911690836136e883836131b9565b1161374b576137096129f761371494838761370387836131b9565b926131c6565b6101408601526131b9565b0361371c5790565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052600d60045260246000fd5b906137cf6137c76137b86132d5565b936101208801948552836131b9565b918585614c01565b909251926137dd829461115f565b52146136b3577fb4205b4200000000000000000000000000000000000000000000000000000000600052600c60045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008152600b600452602490fd5b7fb4205b42000000000000000000000000000000000000000000000000000000008352600a600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526009600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526008600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526007600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526006600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526005600452602490fd5b507fb4205b4200000000000000000000000000000000000000000000000000000000815260048052602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526003600452602483fd5b507fb4205b420000000000000000000000000000000000000000000000000000000081526002600452602490fd5b7fb4205b420000000000000000000000000000000000000000000000000000000083526001600452602483fd5b7f789d326300000000000000000000000000000000000000000000000000000000825260ff16600452602490fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052611ce56024906000600452565b613a6c6130d8565b5090565b9290612e049173ffffffffffffffffffffffffffffffffffffffff613abe67ffffffffffffffff9560405196879581602088019a168a521660408601526080606086015260a08501906103cb565b91166080830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102d9565b9060405191613b0160c0846102d9565b60848352602083019060a03683375a612ee0811115613b5257600091613b2783926125ba565b82602083519301913090f1903d9060848211613b49575b6000908286523e9190565b60849150613b3e565b7fffffffff000000000000000000000000000000000000000000000000000000008063ffffffff5a1660e01b167f2882569d000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b908160209103126100d757516105408161099f565b9080601f830112156100d7578151613bd98161109d565b92613be760405194856102d9565b81845260208085019260051b8201019283116100d757602001905b828210613c0f5750505090565b602080918351613c1e8161099f565b815201910190613c02565b906020828203126100d757815167ffffffffffffffff81116100d7576105409201613bc2565b909267ffffffffffffffff60809373ffffffffffffffffffffffffffffffffffffffff6105409796168452166020830152604082015281606082015201906103cb565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526060949293916020828060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa918215611e6557600092613e0d575b50613d3682614939565b613d75575b50505050815115613d4a575090565b6105409150613d6f60029167ffffffffffffffff166000526004602052604060002090565b016112ca565b8495509373ffffffffffffffffffffffffffffffffffffffff60009495613dcb604051978896879586947f0ba375f900000000000000000000000000000000000000000000000000000000865260048601613c4f565b0392165afa908115611e6557600091613dea575b509038808080613d3b565b613e0791503d806000833e613dff81836102d9565b810190613c29565b38613ddf565b613e3091925060203d602011613e37575b613e2881836102d9565b810190613bad565b9038613d2c565b503d613e1e565b919081101561116c5760051b0190565b356105408161099f565b80156125b5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b9491600390613eb9613e9a613d6f97969489615015565b949198909967ffffffffffffffff166000526004602052604060002090565b90613ec3846110e1565b92613ecd856110e1565b96613ed7866110e1565b96600094855b8b51811015614042576000805b8d8b848a828510613f60575b505050505015613f0857600101613edd565b613f18610e78611ce5928e611171565b7fbd76195f000000000000000000000000000000000000000000000000000000006000529073ffffffffffffffffffffffffffffffffffffffff604492166004526000602452565b610e78613f9892613f92613f8d8873ffffffffffffffffffffffffffffffffffffffff97610d7796613e3e565b613e4e565b95611171565b911614613fa757600101613eea565b60019891508b8d8b8d613fc3613fbd8686611171565b51151590565b6140325792846140168761401c958f8496613ff684613ff0613f8d6140279f9e6140229e61401197613e3e565b92611171565b9073ffffffffffffffffffffffffffffffffffffffff169052565b611171565b52611171565b60019052565b61313f565b96388d8b848a613ef6565b9150508f9150848a9b949b613ef6565b50919950919960005b8a51811015614170576000805b8a8d8b858a8286106140cf575b505050505050156140785760010161404b565b614088610e78611ce5928d611171565b7fbd76195f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff166004526001602452604490565b610e786140fc92613f92613f8d8973ffffffffffffffffffffffffffffffffffffffff97610d7796613e3e565b91161461410c5750600101614058565b9791508b898b614121613fbd8560019d611171565b61416057928061414d866140229561401161401c968f84613ff0613f8d896141549f9e613ff695613e3e565b528d611171565b95388a8d8b858a614065565b8f91508d9250858a9b959b614065565b50919990985060005b8951811015614258576000805b898c8a85898286106141b6575b505050505050156141a657600101614179565b613f18610e78611ce5928c611171565b610e786141e392613f92613f8d8973ffffffffffffffffffffffffffffffffffffffff97610d7796613e3e565b9116146141f35750600101614186565b9691508a888a614208613fbd8560019c611171565b614248579280614235866140229561401161401c96613ff68f9b613ff0613f8d8961423c9f9e849f613e3e565b528c611171565b9438898c8a8589614193565b8c92508e915085899a959a614193565b509098975095929594919493909360ff811690816000995b8b518b10156143515760005b8c86821080614348575b1561433b5773ffffffffffffffffffffffffffffffffffffffff6142ba610d778f610e7890613f928f888e613f8d92613e3e565b9116146142cf576142ca9061313f565b61427c565b936142df909b999193929b613e58565b938a6142ee613fbd838a611171565b614330579161402261401c838061431e858f8f9060019a8f84613ff0613f8d6143259e613ff69461401197613e3e565b528a611171565b985b01999091614270565b505097600190614327565b5050919099600190614327565b50851515614286565b9950935094509297505081614379575050508151810361437057509190565b80825283529190565b611ce59291614387916125e7565b7f403b06ae0000000000000000000000000000000000000000000000000000000060005260ff909116600452602452604490565b604051906103a2826102bd565b908160209103126100d757604051906143e0826102bd565b51815290565b90610540916020815260e06144db6144a861440f855161010060208701526101208601906103cb565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff16606086015260608601516080860152614474608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103cb565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103cb565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103cb565b3d1561453a573d9061452082610359565b9161452e60405193846102d9565b82523d6000602084013e565b606090565b60409073ffffffffffffffffffffffffffffffffffffffff610540949316815281602082015201906103cb565b909291614577612414565b5061459c610d776060840161458c8151615130565b5160208082518301019101613bad565b6040517fbbe4f6db00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152909490936020858060248101038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa948515611e65576000956148c2575b5073ffffffffffffffffffffffffffffffffffffffff851694851561487f5761465c81614963565b61487f5761466c611a138261498d565b61487f575061473d9160209161468288866151bc565b9561468b6143bb565b5080516146b6608086840151930151936146a361033a565b95865267ffffffffffffffff1686860152565b73ffffffffffffffffffffffffffffffffffffffff87166040850152606084015273ffffffffffffffffffffffffffffffffffffffff8916608084015260a083015260c0820152614705610393565b60e0820152604051809381927f39077537000000000000000000000000000000000000000000000000000000008352600483016143e6565b03816000885af16000918161484e575b50614791578461475b61450f565b90611d336040519283927f9fe2f95a0000000000000000000000000000000000000000000000000000000084526004840161453f565b84909373ffffffffffffffffffffffffffffffffffffffff8316036147e4575b505050516147dc6147c061034a565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b602082015290565b6147ed916151bc565b90808210801561483a575b61480257836147b1565b91517fa966e21f0000000000000000000000000000000000000000000000000000000060005260045260249190915260445260646000fd5b5061484581836125e7565b835114156147f8565b61487191925060203d602011614878575b61486981836102d9565b8101906143c8565b903861474d565b503d61485f565b7fae9b4ce90000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045260246000fd5b6148dc91955060203d602011613e3757613e2881836102d9565b9338614634565b61490d7f85572ffb0000000000000000000000000000000000000000000000000000000082615392565b9081614927575b8161491d575090565b6105409150615332565b90506149328161526c565b1590614914565b61490d7f05c7a8d00000000000000000000000000000000000000000000000000000000082615392565b61490d7ff208a58f0000000000000000000000000000000000000000000000000000000082615392565b61490d7faff2afbf0000000000000000000000000000000000000000000000000000000082615392565b61490d7f7909b5490000000000000000000000000000000000000000000000000000000082615392565b73ffffffffffffffffffffffffffffffffffffffff600154163303614a0257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614a3a8151846131b9565b928315614b675760005b848110614a52575050505050565b81811015614b4c57614a67610e788286611171565b73ffffffffffffffffffffffffffffffffffffffff81168015610f3257614a8d8361319d565b878110614a9f57505050600101614a44565b84811015614b1c5773ffffffffffffffffffffffffffffffffffffffff614ac9610e78838a611171565b168214614ad857600101614a8d565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff614b47610e78614b4188856125e7565b89611171565b614ac9565b614b62610e78614b5c84846125e7565b85611171565b614a67565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b60025481101561116c5760026000527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015490565b359060208110614bd4575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b929192614c0c6132aa565b9382811015614f8a57614c2f6133c161339b614c278461313f565b938686613175565b600160ff821603614f5a575082614c45826131ab565b11614f2b5780614c6b614c65614c5d614c72946131ab565b8387876131c6565b90614bc6565b86526131ab565b82811015614efc57614c9761349a6133c161339b614c8f8561313f565b948787613175565b83614ca282846131b9565b11614ecd5781614cc36129f7614cbb84614ccd966131b9565b8388886131c6565b60208801526131b9565b82811015614e9e57614cea61349a6133c161339b614c8f8561313f565b83614cf582846131b9565b11614e6f5781614d0e6129f7614cbb84614d18966131b9565b60408801526131b9565b82811015614e4057614d3561349a6133c161339b614c8f8561313f565b83614d4082846131b9565b11614e1157816134ce6129f7614cbb84614d59966131b9565b82614d638261318f565b11614de25761ffff614d8061363261355361354d6136948661318f565b91169183614d8e84846131b9565b11614db3576129f7614da991836105409661370387836131b9565b60808601526131b9565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601560045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601460045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f789d32630000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601060045260246000fd5b90916060828403126100d757815167ffffffffffffffff81116100d75783614fe2918401613bc2565b92602083015167ffffffffffffffff81116100d757604091615005918501613bc2565b92015160ff811681036100d75790565b91909161503961071e8267ffffffffffffffff166000526004602052604060002090565b90833b615056575b50606001519150615050611f19565b90600090565b61505f846149b7565b15615041576040517f7909b54900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff919091166004820152926000908490602490829073ffffffffffffffffffffffffffffffffffffffff165afa8015611e65576000809481926150f9575b508051158015906150ed575b6150e65750615041565b9392909150565b5060ff821615156150dc565b91506151189294503d8091833e61511081836102d9565b810190614fb9565b909391386150d0565b908160209103126100d7575190565b60208151036151735761514c6020825183010160208301615121565b73ffffffffffffffffffffffffffffffffffffffff81119081156151b0575b506151735750565b611d33906040519182917f8d666f6000000000000000000000000000000000000000000000000000000000835260206004840181815201906103cb565b6104009150103861516b565b91909173ffffffffffffffffffffffffffffffffffffffff604051917f70a0823100000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa6000918161523b575b50615237578261475b61450f565b9150565b61525e91925060203d602011615265575b61525681836102d9565b810190615121565b9038615229565b503d61524c565b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527fffffffff000000000000000000000000000000000000000000000000000000006024830152602482526152cc6044836102d9565b6179185a10615308576020926000925191617530fa6000513d826152fc575b50816152f5575090565b9050151590565b602011159150386152eb565b7fea7f4b120000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160208101917f01ffc9a70000000000000000000000000000000000000000000000000000000083527f01ffc9a7000000000000000000000000000000000000000000000000000000006024830152602482526152cc6044836102d9565b604051907fffffffff0000000000000000000000000000000000000000000000000000000060208301937f01ffc9a7000000000000000000000000000000000000000000000000000000008552166024830152602482526152cc6044836102d9565b80600052600360205260406000205415600014615472576002546801000000000000000081101561029c5760018101600255600060025482101561116c57600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b5060009056fea164736f6c634300081a000a",
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
