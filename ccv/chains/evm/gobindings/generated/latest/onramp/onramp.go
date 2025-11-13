// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package onramp

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

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type OnRampDestChainConfig struct {
	Router               common.Address
	SequenceNumber       uint64
	AddressBytesLength   uint8
	BaseExecutionGasCost uint32
	DefaultExecutor      common.Address
	LaneMandatedCCVs     []common.Address
	DefaultCCVs          []common.Address
	OffRamp              []byte
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector    uint64
	Router               common.Address
	AddressBytesLength   uint8
	BaseExecutionGasCost uint32
	DefaultCCVs          []common.Address
	LaneMandatedCCVs     []common.Address
	DefaultExecutor      common.Address
	OffRamp              []byte
}

type OnRampDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
}

type OnRampReceipt struct {
	Issuer            common.Address
	DestGasLimit      uint32
	DestBytesOverhead uint32
	FeeTokenAmount    *big.Int
	ExtraArgs         []byte
}

type OnRampStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	TokenAdminRegistry common.Address
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateDestChainAddress\",\"inputs\":[{\"name\":\"rawAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"validatedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102fc5760405161628d38819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a1604051615f42908161034b82396080518181816105bf0152818161164f0152611e04015260a0518181816113cb0152611e3d015260c051818181611e7901526127930152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c691461011757806314a8cfa314610112578063181f5a771461010d57806320487ded1461010857806348a98aa4146101035780635cb80c5d146100fe5780636d7fa1ce146100f95780636def4ce7146100f45780637437ff9f146100ef57806379ba5097146100ea5780638da5cb5b146100e55780639041be3d146100e057806390423fa2146100db578063df0aa9e9146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611d61565b611c6d565b6112fb565b611127565b611085565b611033565b610f4a565b610e7e565b610dd8565b610bf6565b610a9d565b610a2f565b61084a565b6107a9565b610217565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576060610150611de4565b610193604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101955760043567ffffffffffffffff81116101955760040160009280601f830112156102135781359367ffffffffffffffff851161021057506020808301928560051b010111610195579190565b80fd5b8380fd5b34610195576102253661019a565b61022d6130c7565b60005b81811061023957005b610244818385611ed0565b9061024e82611f15565b67ffffffffffffffff811690811580156105b3575b801561059d575b8015610584575b610549576102ba906102d4608086019161028b8388611f41565b94906102b460a08a01966102ac6102a2898d611f41565b9490923691611fad565b923691611fad565b906131b0565b67ffffffffffffffff166000526004602052604060002090565b9160208601906103276102e683612003565b859073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61038461033660408901611f1f565b85547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178555565b61039060608801611f37565b6103c96001860191829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6103e06103d6858a611f41565b9060038801612096565b6103f76103ed838a611f41565b9060028801612096565b60c088019161042161040884612003565b73ffffffffffffffffffffffffffffffffffffffff1690565b1561051f57600198610501846105077f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae996104ad6104f9976104656105169a612003565b7fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff0000000083549260201b169116179055565b6104f06104e96104e360e08801936104d26104c8868b612118565b9060048401612201565b5460a01c67ffffffffffffffff1690565b9a612003565b9a86611f41565b97909686611f41565b949093612003565b94612118565b959094604051998a998a6123b9565b0390a201610230565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b5063ffffffff61059660608601611f37565b1615610271565b5060ff6105ac60408601611f1f565b161561026a565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610263565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610100810190811067ffffffffffffffff82111761063257604052565b6105e6565b6060810190811067ffffffffffffffff82111761063257604052565b6040810190811067ffffffffffffffff82111761063257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761063257604052565b604051906106c06101c08361066f565b565b604051906106c060a08361066f565b604051906106c060c08361066f565b67ffffffffffffffff811161063257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6040519061072960208361066f565b60008252565b60005b8381106107425750506000910152565b8181015183820152602001610732565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361078e8151809281875287808801910161072f565b0116010190565b9060206107a6928181520190610752565b90565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761082660408051906107ea818361066f565b601082527f4f6e52616d7020312e372e302d64657600000000000000000000000000000000602083015251918291602083526020830190610752565b0390f35b67ffffffffffffffff81160361019557565b908160a09103126101955790565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576004356108858161082a565b60243567ffffffffffffffff8111610195576108a590369060040161083c565b6108c38267ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff6108fa610408845473ffffffffffffffffffffffffffffffffffffffff1690565b16156109d9579081610826936109296109196080610974960184612118565b9061092386612548565b8461347c565b9261093261264b565b6040840161094081866126b6565b9050610986575b5061096b604086019182516060880194610965600287519201612427565b91613bf1565b9092525261410d565b60405190815292839250602083019150565b6109d3915060206109b56109a36109ae6109a96109a3868b6126b6565b9061270a565b612003565b93886126b6565b01356109c6602088015161ffff1690565b9060e088015192866139b1565b38610947565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610a6960043561082a565b6020610a7f602435610a7a81610a11565b612734565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b3461019557610aab3661019a565b90610acb60035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b818110610ad857005b610ae96104086109a98385876127fa565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa908115610be6576001948891600093610bb6575b5082610b5e575b5050505001610acf565b610b69918391614526565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610b54565b610bd891935060203d8111610bdf575b610bd0818361066f565b81019061280a565b9138610b4d565b503d610bc6565b612728565b60ff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760043567ffffffffffffffff8111610195573660238201121561019557806004013567ffffffffffffffff81116101955736602482840101116101955761082691610c7e916024803592610c7884610beb565b016128e1565b60405191829182610795565b906020808351928381520192019060005b818110610ca85750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610c9b565b906107a69160208152610d0060208201835173ffffffffffffffffffffffffffffffffffffffff169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015163ffffffff166080820152608082015173ffffffffffffffffffffffffffffffffffffffff1660a082015260e0610da4610d7160a085015161010060c0860152610120850190610c8a565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610c8a565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610752565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff600435610e1c8161082a565b606060e0604051610e2c81610615565b600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201520152166000526004602052610826610e726040600020612548565b60405191829182610cd4565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610eb5611dc5565b50604051610ec281610637565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff600354166040820152604051809161082682606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760005473ffffffffffffffffffffffffffffffffffffffff81163303611009577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff6004356110c98161082a565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff811161110d5760405167ffffffffffffffff9091168152602090f35b61200d565b35906106c082610a11565b8015150361019557565b346101955760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557600060405161116481610637565b60043561117081610a11565b815260243561117e8161111d565b602082019081526044359061119282610a11565b604083019182526111a16130c7565b73ffffffffffffffffffffffffffffffffffffffff835116159182156112db575b5081156112d0575b506112a85780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390611293611de4565b6112a260405192839283614626565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b9050511515386111ca565b5173ffffffffffffffffffffffffffffffffffffffff16159150386111c2565b346101955760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761133560043561082a565b60243567ffffffffffffffff81116101955761135590369060040161083c565b604435611363606435610a11565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff00000000000000000000000000000000169082015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610be657600091611c3e575b50611c015760025460a01c60ff16611bd757611452740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61147260043567ffffffffffffffff166000526004602052604060002090565b9173ffffffffffffffffffffffffffffffffffffffff6064351615611bad578254926114b461040873ffffffffffffffffffffffffffffffffffffffff861681565b3303611b83576114c76080830183612118565b6114d083612548565b6114dd929060043561347c565b9360a01c67ffffffffffffffff166114f490612a1d565b81547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617825590845163ffffffff16602086015161ffff166040513060601b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201528060348101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810182526115b3908261066f565b6040517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060643560601b166020820152601481526115f260348261066f565b6115fc8780612118565b865460e01c60ff169061160e926128e1565b60a08a01519161162160408a018a6126b6565b61162b9150612a7e565b9361163960208b018b612118565b9690976116446106b0565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529a67ffffffffffffffff6004351660208d015267ffffffffffffffff1660408c0152600060608c015263ffffffff1660808b015261ffff1660a08a0152600060c08a015260e08901526116c660048801612488565b61010089015261012088015261014087015261016086015261018085015236906116ef926128aa565b6101a08301526116fd61264b565b61170a60408501856126b6565b9050611b2c575b61173061175f9160408801516060890194610965600287519201612427565b8352806040880152611759608088015173ffffffffffffffffffffffffffffffffffffffff1690565b906146ed565b60c083015261176c612aeb565b9161177a868560043561410d565b5063ffffffff166060830152602084019190825261179b60408601866126b6565b9050611a78575b6117ae8195929561515d565b80855260208151910120906117c7604089015151612b5c565b9460408101958652606060009401935b60408a015180518210156119cc5760206118176104086104086117fd8661185b96612b48565b5173ffffffffffffffffffffffffffffffffffffffff1690565b611822848a51612b48565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612bc3565b03915afa8015610be65773ffffffffffffffffffffffffffffffffffffffff9160009161199e575b5016801561193857906000878b938783886118e26118ab886118a48f612003565b9751612b48565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612d2c565b03925af18015610be6578161191091600194600091611917575b508a519061190a8383612b48565b52612b48565b50016117d7565b611932913d8091833e61192a818361066f565b810190612c26565b386118fc565b61058061194c6117fd8460408f0151612b48565b7f83c758a60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660049081523567ffffffffffffffff16602452604490565b6119bf915060203d81116119c5575b6119b7818361066f565b810190612713565b38611883565b503d6119ad565b61082685808a8c7f276d7e038bc94e70aa9c54ac8cf3a3674da9252bdccbf8a0593523768f989c9667ffffffffffffffff89611a38611a1660408e015167ffffffffffffffff1690565b915194519551604051938493169667ffffffffffffffff600435169684612fd6565b0390a4611a687fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b6001611a8760408701876126b6565b905003611b0257611aef611ad9611aa46109a360408901896126b6565b60c08a0151805115611af557905b60208b015161ffff169060e08c01519260643591611ad4600435913690612b0a565b6148db565b61018083015190611ae982612b3b565b52612b3b565b506117a2565b5061014084015190611ab2565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b5061175f611730611b7b611b496109a96109a360408901896126b6565b6020611b5b6109a360408a018a6126b6565b0135611b6c60208b015161ffff1690565b9060e08b0151926004356139b1565b915050611711565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105806004359067ffffffffffffffff60249216600452565b611c60915060203d602011611c66575b611c58818361066f565b810190612a08565b386113fc565b503d611c4e565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955773ffffffffffffffffffffffffffffffffffffffff600435611cbd81610a11565b611cc56130c7565b16338114611d3757807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557611d9b60043561082a565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611dd282610637565b60006040838281528260208201520152565b611dec611dc5565b50604051611df981610637565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611f105760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181360301821215610195570190565b611ea1565b356107a68161082a565b356107a681610beb565b63ffffffff81160361019557565b356107a681611f29565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160051b3603831361019557565b67ffffffffffffffff81116106325760051b60200190565b929190611fb981611f95565b93611fc7604051958661066f565b602085838152019160051b810192831161019557905b828210611fe957505050565b602080918335611ff881610a11565b815201910190611fdd565b356107a681610a11565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81160361110d57565b8181029291811591840414171561110d57565b81811061208a575050565b6000815560010161207f565b9067ffffffffffffffff8311610632576801000000000000000083116106325781548383558084106120fa575b5090600052602060002060005b8381106120dd5750505050565b60019060208435946120ee86610a11565b019381840155016120d0565b6121129083600052846020600020918201910161207f565b386120c3565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff82116101955760200191813603831361019557565b90600182811c921680156121b2575b602083101461218357565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612178565b9190601f81116121cb57505050565b6106c0926000526020600020906020601f840160051c830193106121f7575b601f0160051c019061207f565b90915081906121ea565b90929167ffffffffffffffff811161063257612227816122218454612169565b846121bc565b6000601f821160011461228557819061227693949560009261227a575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b013590503880612244565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216946122b884600052602060002090565b91805b8781106123115750836001959697106122d9575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690553880806122cf565b909260206001819286860135815501940191016122bb565b9160209082815201919060005b8181106123435750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561236c81610a11565b168152019401929101612336565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b94916124139373ffffffffffffffffffffffffffffffffffffffff95866124059367ffffffffffffffff6107a69e9c9d9b96168a5216602089015260c0604089015260c0880191612329565b918583036060870152612329565b9416608082015260a081850391015261237a565b906040519182815491828252602082019060005260206000209260005b8181106124595750506106c09250038361066f565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612444565b906040519182600082549261249c84612169565b808452936001811690811561250857506001146124c1575b506106c09250038361066f565b90506000929192526020600020906000915b8183106124ec5750509060206106c092820101386124b4565b60209193508060019154838589010152019101909184926124d3565b602093506106c09592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386124b4565b9060405161255581610615565b60e0612646600483956125bd6125b3825473ffffffffffffffffffffffffffffffffffffffff8082161688526125aa6125998267ffffffffffffffff9060a01c1690565b67ffffffffffffffff1660208a0152565b60e01c60ff1690565b60ff166040870152565b61261e61260160018301546125e56125d88263ffffffff1690565b63ffffffff1660608a0152565b60201c73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166080870152565b61262a60028201612427565b60a086015261263b60038201612427565b60c086015201612488565b910152565b6040519061265a60208361066f565b6000808352366020840137565b9061267182611f95565b61267e604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06126ac8294611f95565b0190602036910137565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160061b3603831361019557565b9015611f105790565b9081602091031261019557516107a681610a11565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015610be65773ffffffffffffffffffffffffffffffffffffffff916000916127dd57501690565b6127f6915060203d6020116119c5576119b7818361066f565b1690565b9190811015611f105760051b0190565b90816020910312610195575190565b9160206107a693818152019161237a565b60ff166020039060ff821161110d57565b909291928311610195579190565b906004116101955790600490565b90939293848311610195578411610195578101920390565b35906020811061287d575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b9291926128b6826106e0565b916128c4604051938461066f565b829481845281830111610195578281602093846000960137010152565b9160ff8116906020821061293d575b50810361290357906107a69136916128aa565b906129396040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612819565b0390fd5b602083116129d357602083036128f057905061297161296b60ff6129638496959661282a565b16858561283b565b9061286f565b61299d5791612996918161299061298a6107a69661282a565b60ff1690565b91612857565b36916128aa565b506129396040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612819565b6040517f3aeba39000000000000000000000000000000000000000000000000000000000815280612939858760048401612819565b9081602091031261019557516107a68161111d565b67ffffffffffffffff1667ffffffffffffffff811461110d5760010190565b6040519060c0820182811067ffffffffffffffff82111761063257604052606060a0836000815282602082015282604082015282808201528260808201520152565b90612a8882611f95565b612a95604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612ac38294611f95565b019060005b828110612ad457505050565b602090612adf612a3c565b82828501015201612ac8565b60405190612af882610637565b60606040838281528260208201520152565b919082604091031261019557604051612b2281610653565b60208082948035612b3281610a11565b84520135910152565b805115611f105760200190565b8051821015611f105760209160051b010190565b90612b6682611f95565b612b73604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612ba18294611f95565b019060005b828110612bb257505050565b806060602080938501015201612ba6565b60409067ffffffffffffffff6107a694931681528160208201520190610752565b81601f82011215610195578051612bfa816106e0565b92612c08604051948561066f565b81845260208284010111610195576107a6916020808501910161072f565b9060208282031261019557815167ffffffffffffffff8111610195576107a69201612be4565b9080602083519182815201916020808360051b8301019401926000915b838310612c7857505050505090565b9091929394602080612d1d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a0612d0c612cfa612ce8612cd68887015160c08a88015260c0870190610752565b60408701518682036040880152610752565b60608601518582036060870152610752565b60808501518482036080860152610752565b9201519060a0818403910152610752565b97019301930191939290612c69565b9193906107a69593612f28612f4d9260a08652612d5660a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612ef5612ebf612e89612e53612e1d612de98c61026060e08a0151916101c06101808201520190610752565b6101008801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001888f0152610752565b6101208701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101c08e0152610752565b6101408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608c8303016101e08d0152610752565b6101608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8303016102008c0152610752565b6101808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a8303016102208b0152612c4c565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6087830301610240880152610752565b956020850152604084019073ffffffffffffffffffffffffffffffffffffffff169052565b60608201526080818403910152610752565b9080602083519182815201916020808360051b8301019401926000915b838310612f8b57505050505090565b9091929394602080612fc7837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610752565b97019301930191939290612f7c565b939290612feb90606086526060860190610752565b938085036020820152825180865260208601906020808260051b8901019501916000905b82821061302d57505050506107a69394506040818403910152612f5f565b909192956020806130b9837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001960301865260a060808c5173ffffffffffffffffffffffffffffffffffffffff815116845263ffffffff86820151168685015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610752565b98019201920190929161300f565b73ffffffffffffffffffffffffffffffffffffffff6001541633036130e857565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b906001820180921161110d57565b600101908160011161110d57565b906014820180921161110d57565b9190820180921161110d57565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161110d57565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe820191821161110d57565b9190820391821161110d57565b8051916131be81518461313c565b9283156133155760005b8481106131d6575050505050565b818110156132fa576131eb6117fd8286612b48565b73ffffffffffffffffffffffffffffffffffffffff811680156132d05761321183613112565b878110613223575050506001016131c8565b848110156132a05773ffffffffffffffffffffffffffffffffffffffff61324d6117fd838a612b48565b16821461325c57600101613211565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6132cb6117fd6132c588856131a3565b89612b48565b61324d565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6133106117fd61330a84846131a3565b85612b48565b6131eb565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061334c82610615565b606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106133b2575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261019557825167ffffffffffffffff8111610195578261340c918501612be4565b92602081015161341b81611f29565b92604082015167ffffffffffffffff8111610195576107a69201612be4565b60409067ffffffffffffffff6107a69593168152816020820152019161237a565b9060ff613475602092959495604085526040850190610752565b9416910152565b9291909261348861333f565b60048410158061384a575b15613710575050906134a4916154b1565b9060c0820151805161369c575b506040820180515160005b8181106135ed57505080515115613550575b505b6080820173ffffffffffffffffffffffffffffffffffffffff613507825173ffffffffffffffffffffffffffffffffffffffff1690565b161561351257505090565b61353660806107a693015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b9060c0819493940192613564845151612667565b8352613571845151612b5c565b946060810195865260005b855180518210156135df57906135c061359a6117fd83600195612b48565b6135a5838951612b48565b9073ffffffffffffffffffffffffffffffffffffffff169052565b6135d88189516135ce61071a565b61190a8383612b48565b500161357c565b5050935093509050386134ce565b6135f681613112565b82811061360657506001016134bc565b6136146117fd838651612b48565b73ffffffffffffffffffffffffffffffffffffffff61363a6104086117fd858951612b48565b911614613649576001016135f6565b61058061365a6117fd848751612b48565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b60006136e5916136b0604085015160ff1690565b9060405193849283927f6d7fa1ce0000000000000000000000000000000000000000000000000000000084526004840161345b565b0381305afa8015610be657156134b157613709903d806000833e61192a818361066f565b50386134b1565b60c0859692960194613723865151612667565b9460408301958652613736875151612b5c565b976060840198895260005b8851805182101561377f579061376a61375f6117fd83600195612b48565b6135a5838c51612b48565b613778818c516135ce61071a565b5001613741565b505091955091955060009296506137ec936137b561040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b91604051958694859384937f9cc199960000000000000000000000000000000000000000000000000000000085526004850161343a565b03915afa8015610be65760009060009060009061381e575b60a086015263ffffffff16845290505b60c08301526134d0565b505050613840613814913d806000833e613838818361066f565b8101906133e4565b9192508291613804565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006138a061389a8787612849565b9061337e565b1614613493565b6020818303126101955780519067ffffffffffffffff821161019557019080601f830112156101955781516138db81611f95565b926138e9604051948561066f565b81845260208085019260051b82010192831161019557602001905b8282106139115750505090565b60208091835161392081610a11565b815201910190613904565b95949060009460a09467ffffffffffffffff61397f9573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c0860190610752565b930152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461110d5760010190565b9492949390936139e060036139da8367ffffffffffffffff166000526004602052604060002090565b01612427565b9473ffffffffffffffffffffffffffffffffffffffff613a01818316612734565b16926040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180613a5e60048201907fdc0cbd3600000000000000000000000000000000000000000000000000000000602083019252565b0381885afa908115610be657600091613bd2575b5015613bc85790613ab8600095949392604051998a96879586957f89720a620000000000000000000000000000000000000000000000000000000087526004870161392b565b03915afa928315610be657600093613ba3575b50825115613b9e57613ae8613ae3845184519061313c565b612667565b6000918293835b8651811015613b4d57613b056117fd8289612b48565b73ffffffffffffffffffffffffffffffffffffffff811615613b415790613b3b6001926135a5613b3489613984565b9888612b48565b01613aef565b50945060018095613b3b565b509193909450613b5e575b50815290565b60005b8151811015613b965780613b90613b7d6117fd60019486612b48565b6135a5613b8987613984565b9688612b48565b01613b61565b505038613b58565b915090565b613bc19193503d806000833e613bb9818361066f565b8101906138a7565b9138613acb565b5050505050915090565b613beb915060203d602011611c6657611c58818361066f565b38613a72565b93919293613c0d613c05825186519061313c565b86519061313c565b90613c20613c1a83612667565b92612b5c565b94600096875b8351891015613c865788613c7c613c6f600193613c57613c4d6117fd8e9f9d9e9d8b612b48565b6135a5838c612b48565b613c75613c64858c612b48565b519180938491613984565b9c612b48565b528b612b48565b5001979695613c26565b959250929350955060005b8651811015613d2057613ca76117fd8289612b48565b600073ffffffffffffffffffffffffffffffffffffffff8216815b888110613cf4575b5050906001929115613cde575b5001613c91565b613cee906135a5613b3489613984565b38613cd7565b81613d056104086117fd848c612b48565b14613d1257600101613cc2565b506001915081905038613cca565b509390945060005b8551811015613dbe57613d3e6117fd8288612b48565b600073ffffffffffffffffffffffffffffffffffffffff8216815b878110613d92575b5050906001929115613d75575b5001613d28565b613d8c906135a5613d8588613984565b9787612b48565b38613d6e565b81613da36104086117fd848b612b48565b14613db057600101613d59565b506001915081905038613d61565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176106325760405260606080836000815260006020820152600060408201526000838201520152565b90613e1382611f95565b613e20604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613e4e8294611f95565b019060005b828110613e5f57505050565b602090613e6a613dca565b82828501015201613e53565b519061ffff8216820361019557565b9081606091031261019557613e9981613e76565b9160406020830151613eaa81611f29565b9201516107a681611f29565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561019557016020813591019167ffffffffffffffff821161019557813603831361019557565b9160209082815201919060005b818110613f205750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735613f4981610a11565b16815260208781013590820152019401929101613f13565b949391929067ffffffffffffffff16855260806020860152613fd8613f9b613f898580613eb6565b60a060808a015261012089019161237a565b613fa86020860186613eb6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808984030160a08a015261237a565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215610195578401916020833593019167ffffffffffffffff8411610195578360061b36038313610195576106c0956140a9614073836060976140e8978d60c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80826140da9a0301910152613f06565b9161409f614082888301611112565b73ffffffffffffffffffffffffffffffffffffffff1660e08d0152565b6080810190613eb6565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101008c015261237a565b908782036040890152610752565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff821161110d57565b60009160009260408501928351519561414661414161413c60408701996141348b896126b6565b91905061313c565b613112565b613e09565b906000986000975b875180518a10156143d3579a6141bd60206141776104088d9e9f6117fd8f9e9f61040892612b48565b614186606088019c8d51612b48565b519060405180809581947f958021a70000000000000000000000000000000000000000000000000000000083528d60048401612bc3565b03915afa8015610be65773ffffffffffffffffffffffffffffffffffffffff916000916143b5575b50169687156143545760606141fb8d8c51612b48565b51602087015161ffff16998b8a6142416040519d8e95869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613f61565b03915afa978815610be6578b9a8d60008093819c614317575b508161ffff8192169e519061426e91612b48565b5173ffffffffffffffffffffffffffffffffffffffff1692519061429191612b48565b5161429a6106c2565b73ffffffffffffffffffffffffffffffffffffffff909316835263ffffffff8416602084015263ffffffff8c1660408401528d606084015260808301526142e19089612b48565b526142ec8d88612b48565b506142f6916140f3565b96614300916140f3565b9761430a9161313c565b986001019796999861414e565b61ffff9c508294508291506143429060603d811161434d575b61433a818361066f565b810190613e85565b9d509490915061425a565b503d614330565b610580876143666117fd8f8f51612b48565b7f83c758a60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045267ffffffffffffffff16602452604490565b6143cd915060203d81116119c5576119b7818361066f565b386141e5565b5092965096509261445661446f9161442b61440f8761448e978d6144076143fe8d6020810190612118565b9290508d6126b6565b929050615763565b6144198a51613149565b90614424828c612b48565b5289612b48565b50614450602061444461443e8b51613149565b8b612b48565b51015163ffffffff1690565b906140f3565b9861445060406144446144698a51613149565b8a612b48565b50606061448561447f8751613149565b87612b48565b5101519061313c565b9461449981846126b6565b90506144a8575b505050929190565b6144bd6109a96109a360e09361451d966126b6565b9101516144e76144cb6106c2565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b600060208301526000604083015260006060830152608082015261450b8351613176565b906145168285612b48565b5282612b48565b503880806144a0565b9073ffffffffffffffffffffffffffffffffffffffff6145f89392604051938260208601947fa9059cbb00000000000000000000000000000000000000000000000000000000865216602486015260448501526044845261458860648561066f565b1660008060409384519561459c868861066f565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d1561461d573d6145e96145e0826106e0565b9451948561066f565b83523d6000602085013e615e71565b805180614603575050565b81602080614618936106c09501019101612a08565b6159b1565b60609250615e71565b9160606106c09294936146738160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b906146b2826106e0565b6146bf604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06126ac82946106e0565b91825160148102908082046014149015171561110d5761471761471261471c92613120565b61312e565b6146a8565b92601461472885612b3b565b5360009160215b825184101561475a57601460019161474a6117fd8787612b48565b60601b818901520193019261472f565b92509392905060601b908201526020815191012090565b6040519061477e82610653565b60606020838281520152565b919060408382031261019557604051906147a382610653565b8193805167ffffffffffffffff811161019557826147c2918301612be4565b835260208101519167ffffffffffffffff8311610195576020926126469201612be4565b9060208282031261019557815167ffffffffffffffff8111610195576107a6920161478a565b90608073ffffffffffffffffffffffffffffffffffffffff81614838855160a0865260a0860190610752565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b9060206107a692818152019061480c565b919060408382031261019557825167ffffffffffffffff8111610195576020916148a891850161478a565b92015190565b61ffff6148c76107a6959360608452606084019061480c565b931660208201526040818403910152610752565b909192936148e7612a3c565b506020820190815115614e7b5761491b610408610a7a610408865173ffffffffffffffffffffffffffffffffffffffff1690565b9573ffffffffffffffffffffffffffffffffffffffff87169283158015614df0575b614d8d576149c881519161494f614771565b5051865173ffffffffffffffffffffffffffffffffffffffff16906149a66149756106c2565b8b815267ffffffffffffffff8b1660208201529573ffffffffffffffffffffffffffffffffffffffff166040870152565b606085015273ffffffffffffffffffffffffffffffffffffffff166080840152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdc0cbd36000000000000000000000000000000000000000000000000000000006004820152602081602481885afa908115610be657600091614d6e575b5015614c78575091614a7496979160008094604051998a95869485937fb1c71c65000000000000000000000000000000000000000000000000000000008552600485016148ae565b03925af18015610be657600094600091614c2f575b50946000614bce93614b6c614b08614b40956117fd614adc9a965b6040519b8c91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018c528b61066f565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810186528561066f565b614b99614b8f84519267ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b9060405195869283927f6d7fa1ce0000000000000000000000000000000000000000000000000000000084526004840161345b565b0381305afa928315610be657600093614c0f575b506020015193614bf06106d1565b958652602086015260408501526060840152608083015260a082015290565b6020919350614c28903d806000833e61192a818361066f565b9290614be2565b614adc95506117fd969150614bce93614b6c614b08614b4095614c656000953d8088833e614c5d818361066f565b81019061487d565b9b909b969b9a5050955050509350614a89565b979161ffff919597935016614d445751614d1a576000614cc593604051809581927f9a4575b90000000000000000000000000000000000000000000000000000000083526004830161486c565b038183855af1918215610be657614adc95614b6c614b086000936117fd614b4097614bce998791614cf8575b5096614aa4565b614d1491503d8089833e614d0c818361066f565b8101906147e6565b38614cf1565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614d87915060203d602011611c6657611c58818361066f565b38614a2c565b610580614dae865173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818c5afa908115610be657600091614e5c575b501561493d565b614e75915060203d602011611c6657611c58818361066f565b38614e55565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b90614eb86020928281519485920161072f565b0190565b615009615010996107a69f9e9c9799614fd860019f9b9a60019b6150099f9890614f806150099c8f6150109d7fffffffffffffffff000000000000000000000000000000000000000000000000614fac968d60449f948380967f0100000000000000000000000000000000000000000000000000000000000000845260c01b1691015260c01b1660098d015260c01b1660118b015260198a01907fffffffff000000000000000000000000000000000000000000000000000000009060e01b169052565b60e01b7fffffffff0000000000000000000000000000000000000000000000000000000016601d880152565b60f01b7fffff000000000000000000000000000000000000000000000000000000000000166021860152565b602384015260f81b7fff00000000000000000000000000000000000000000000000000000000000000166043830152565b0190614ea5565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff000000000000000000000000000000000000000000000000000000000000006107a69f9e9c9860f81b1681526150ab825180936020898501910161072f565b019160f01b16838201526150c982518093602060038501910161072f565b01019160f01b16838201526150e7825180936020898501910161072f565b01019160f01b1660028201520190614ea5565b6106c0909291926020604051948261511b879451809285808801910161072f565b830161512f8251809385808501910161072f565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184528361066f565b60e081019060ff825151116154825761010081019060ff825151116154535761012081019060ff825151116154245761014081019060ff825151116153f55761016081019061ffff825151116153c657610180810195600187515111615397576101a082019561ffff8751511161536657825167ffffffffffffffff169160208401516151f19067ffffffffffffffff1690565b9360408101516152089067ffffffffffffffff1690565b97606082015161521b9063ffffffff1690565b92608083015161522e9063ffffffff1690565b9460a084015161523f9061ffff1690565b9360c0015191519182516152539060ff1690565b91519485516152629060ff1690565b94519788516152719060ff1690565b976040519e8f9c60208e019c6152869d614ebc565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184526152b6908461066f565b60609451805115159561534593615319926107a69861534a575b5051956152de875160ff1690565b9351916152ed835161ffff1690565b916152fa825161ffff1690565b905193615309855161ffff1690565b936040519a8b9860208a0161503c565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184528361066f565b6150fa565b61535f91925061535990612b3b565b51615b44565b90386152d0565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526105806024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b906154ba61333f565b91601182106156dd5780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160361566a5750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61553981612667565b6040860190815261554982612b5c565b906060870191825260005b83811061561e57505050506155c983836155bf6155b36155a96155a26155826155d398876155dd9c9b615d06565b73ffffffffffffffffffffffffffffffffffffffff90911660808d015290565b8585615ddc565b92919036916128aa565b60a08a01528383615e44565b94919036916128aa565b60c0880152615ddc565b93919036916128aa565b60e084015281036155ec575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b8060019161566361564d61564661563961565d9a8d8d615d06565b91906135a5868a51612b48565b8b8b615ddc565b9391889a919a51949a36916128aa565b92612b48565b5201615554565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b90816020910312610195576107a690613e76565b9261ffff6107a6959367ffffffffffffffff61575594168652166020850152608060408501526080840190610c8a565b916060818403910152610752565b909161576d613dca565b5061578c8267ffffffffffffffff166000526004602052604060002090565b9361579c855460ff9060e01c1690565b9061585661583b615832608084019461581e6157f86157f06157e360016157d78b5173ffffffffffffffffffffffffffffffffffffffff1690565b9e015463ffffffff1690565b885163ffffffff16614450565b9a607561313c565b9761582c61582460ff61581260a08b019c8d51519061313c565b95169461581e8661203c565b9061313c565b93604f61313c565b9061206c565b63ffffffff1690565b925173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811673eba517d200000000000000000000000000000000036158ea57505061ffff92506158dc906158cf6000935b51956158c26158a66106c2565b73ffffffffffffffffffffffffffffffffffffffff9099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b9061590d61040860209373ffffffffffffffffffffffffffffffffffffffff1690565b604061591d8484015161ffff1690565b9201519185519661595d604051988995869485947fe962e69e00000000000000000000000000000000000000000000000000000000865260048601615725565b03915afa908115610be6576158cf6158dc9261ffff95600091615982575b5093615899565b6159a4915060203d6020116159aa575b61599c818361066f565b810190615711565b3861597b565b503d615992565b156159b857565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b966002615b189761500960226107a69f9e9c9799600199859f9b7fff00000000000000000000000000000000000000000000000000000000000000906150099f826150099c6150109c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615ac9825180936020898501910161072f565b019160f81b1683820152615ae782518093602060238501910161072f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff82515111615cd757604081019160ff83515111615ca857606082019160ff83515111615c7957608081019260ff84515111615c4a5760a0820161ffff81515111615c1b576107a694615bef9351945191615ba6835160ff1690565b975191615bb4835160ff1690565b945190615bc2825160ff1690565b905193615bd0855160ff1690565b935196615bdf885161ffff1690565b966040519c8d9b60208d01615a3c565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261066f565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602560045260246000fd5b929190926001820191848311615daa5781013560001a828115615d9f575060148103615d72578201938411615d3e57013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b91906002820191818311615daa578381013560f01c0160020192818411615e1057918391615e0993612857565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b91906001820191818311615daa578381013560001a0160010192818411615e1057918391615e0993612857565b91929015615eec5750815115615e85575090565b3b15615e8e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615eff5750805190602001fd5b612939906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161079556fea164736f6c634300081a000a",
}

var OnRampABI = OnRampMetaData.ABI

var OnRampBin = OnRampMetaData.Bin

func DeployOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig OnRampStaticConfig, dynamicConfig OnRampDynamicConfig) (common.Address, *types.Transaction, *OnRamp, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OnRampBin), backend, staticConfig, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OnRamp{address: address, abi: *parsed, OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
}

type OnRamp struct {
	address common.Address
	abi     abi.ABI
	OnRampCaller
	OnRampTransactor
	OnRampFilterer
}

type OnRampCaller struct {
	contract *bind.BoundContract
}

type OnRampTransactor struct {
	contract *bind.BoundContract
}

type OnRampFilterer struct {
	contract *bind.BoundContract
}

type OnRampSession struct {
	Contract     *OnRamp
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OnRampCallerSession struct {
	Contract *OnRampCaller
	CallOpts bind.CallOpts
}

type OnRampTransactorSession struct {
	Contract     *OnRampTransactor
	TransactOpts bind.TransactOpts
}

type OnRampRaw struct {
	Contract *OnRamp
}

type OnRampCallerRaw struct {
	Contract *OnRampCaller
}

type OnRampTransactorRaw struct {
	Contract *OnRampTransactor
}

func NewOnRamp(address common.Address, backend bind.ContractBackend) (*OnRamp, error) {
	abi, err := abi.JSON(strings.NewReader(OnRampABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OnRamp{address: address, abi: abi, OnRampCaller: OnRampCaller{contract: contract}, OnRampTransactor: OnRampTransactor{contract: contract}, OnRampFilterer: OnRampFilterer{contract: contract}}, nil
}

func NewOnRampCaller(address common.Address, caller bind.ContractCaller) (*OnRampCaller, error) {
	contract, err := bindOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampCaller{contract: contract}, nil
}

func NewOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*OnRampTransactor, error) {
	contract, err := bindOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnRampTransactor{contract: contract}, nil
}

func NewOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*OnRampFilterer, error) {
	contract, err := bindOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnRampFilterer{contract: contract}, nil
}

func bindOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OnRamp *OnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRamp.Contract.OnRampCaller.contract.Call(opts, result, method, params...)
}

func (_OnRamp *OnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.Contract.OnRampTransactor.contract.Transfer(opts)
}

func (_OnRamp *OnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRamp.Contract.OnRampTransactor.contract.Transact(opts, method, params...)
}

func (_OnRamp *OnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OnRamp.Contract.contract.Call(opts, result, method, params...)
}

func (_OnRamp *OnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.Contract.contract.Transfer(opts)
}

func (_OnRamp *OnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OnRamp.Contract.contract.Transact(opts, method, params...)
}

func (_OnRamp *OnRampCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (OnRampDestChainConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	if err != nil {
		return *new(OnRampDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampDestChainConfig)).(*OnRampDestChainConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetDestChainConfig(destChainSelector uint64) (OnRampDestChainConfig, error) {
	return _OnRamp.Contract.GetDestChainConfig(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetDestChainConfig(destChainSelector uint64) (OnRampDestChainConfig, error) {
	return _OnRamp.Contract.GetDestChainConfig(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(OnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampDynamicConfig)).(*OnRampDynamicConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRamp.Contract.GetDynamicConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetDynamicConfig() (OnRampDynamicConfig, error) {
	return _OnRamp.Contract.GetDynamicConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRamp *OnRampSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextSequenceNumber(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextSequenceNumber(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OnRamp *OnRampSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRamp.Contract.GetFee(&_OnRamp.CallOpts, destChainSelector, message)
}

func (_OnRamp *OnRampCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _OnRamp.Contract.GetFee(&_OnRamp.CallOpts, destChainSelector, message)
}

func (_OnRamp *OnRampCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPoolBySourceToken(&_OnRamp.CallOpts, arg0, sourceToken)
}

func (_OnRamp *OnRampCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _OnRamp.Contract.GetPoolBySourceToken(&_OnRamp.CallOpts, arg0, sourceToken)
}

func (_OnRamp *OnRampCaller) GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(OnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(OnRampStaticConfig)).(*OnRampStaticConfig)

	return out0, err

}

func (_OnRamp *OnRampSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRamp.Contract.GetStaticConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetStaticConfig() (OnRampStaticConfig, error) {
	return _OnRamp.Contract.GetStaticConfig(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetSupportedTokens(&_OnRamp.CallOpts, arg0)
}

func (_OnRamp *OnRampCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _OnRamp.Contract.GetSupportedTokens(&_OnRamp.CallOpts, arg0)
}

func (_OnRamp *OnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OnRamp *OnRampSession) Owner() (common.Address, error) {
	return _OnRamp.Contract.Owner(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) Owner() (common.Address, error) {
	return _OnRamp.Contract.Owner(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OnRamp *OnRampSession) TypeAndVersion() (string, error) {
	return _OnRamp.Contract.TypeAndVersion(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) TypeAndVersion() (string, error) {
	return _OnRamp.Contract.TypeAndVersion(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCaller) ValidateDestChainAddress(opts *bind.CallOpts, rawAddress []byte, addressBytesLength uint8) ([]byte, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "validateDestChainAddress", rawAddress, addressBytesLength)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_OnRamp *OnRampSession) ValidateDestChainAddress(rawAddress []byte, addressBytesLength uint8) ([]byte, error) {
	return _OnRamp.Contract.ValidateDestChainAddress(&_OnRamp.CallOpts, rawAddress, addressBytesLength)
}

func (_OnRamp *OnRampCallerSession) ValidateDestChainAddress(rawAddress []byte, addressBytesLength uint8) ([]byte, error) {
	return _OnRamp.Contract.ValidateDestChainAddress(&_OnRamp.CallOpts, rawAddress, addressBytesLength)
}

func (_OnRamp *OnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "acceptOwnership")
}

func (_OnRamp *OnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRamp.Contract.AcceptOwnership(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OnRamp.Contract.AcceptOwnership(&_OnRamp.TransactOpts)
}

func (_OnRamp *OnRampTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_OnRamp *OnRampSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyDestChainConfigUpdates(&_OnRamp.TransactOpts, destChainConfigArgs)
}

func (_OnRamp *OnRampTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error) {
	return _OnRamp.Contract.ApplyDestChainConfigUpdates(&_OnRamp.TransactOpts, destChainConfigArgs)
}

func (_OnRamp *OnRampTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ForwardFromRouter(&_OnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.ForwardFromRouter(&_OnRamp.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_OnRamp *OnRampTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_OnRamp *OnRampSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetDynamicConfig(&_OnRamp.TransactOpts, dynamicConfig)
}

func (_OnRamp *OnRampTransactorSession) SetDynamicConfig(dynamicConfig OnRampDynamicConfig) (*types.Transaction, error) {
	return _OnRamp.Contract.SetDynamicConfig(&_OnRamp.TransactOpts, dynamicConfig)
}

func (_OnRamp *OnRampTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "transferOwnership", to)
}

func (_OnRamp *OnRampSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.TransferOwnership(&_OnRamp.TransactOpts, to)
}

func (_OnRamp *OnRampTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.TransferOwnership(&_OnRamp.TransactOpts, to)
}

func (_OnRamp *OnRampTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRamp.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_OnRamp *OnRampSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawFeeTokens(&_OnRamp.TransactOpts, feeTokens)
}

func (_OnRamp *OnRampTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _OnRamp.Contract.WithdrawFeeTokens(&_OnRamp.TransactOpts, feeTokens)
}

type OnRampCCIPMessageSentIterator struct {
	Event *OnRampCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampCCIPMessageSent)
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
		it.Event = new(OnRampCCIPMessageSent)
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

func (it *OnRampCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *OnRampCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	MessageId         [32]byte
	EncodedMessage    []byte
	Receipts          []OnRampReceipt
	VerifierBlobs     [][]byte
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPMessageSentIterator{contract: _OnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampCCIPMessageSent)
				if err := _OnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseCCIPMessageSent(log types.Log) (*OnRampCCIPMessageSent, error) {
	event := new(OnRampCCIPMessageSent)
	if err := _OnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampConfigSetIterator struct {
	Event *OnRampConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampConfigSet)
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
		it.Event = new(OnRampConfigSet)
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

func (it *OnRampConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampConfigSet struct {
	StaticConfig  OnRampStaticConfig
	DynamicConfig OnRampDynamicConfig
	Raw           types.Log
}

func (_OnRamp *OnRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error) {

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OnRampConfigSetIterator{contract: _OnRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampConfigSet)
				if err := _OnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseConfigSet(log types.Log) (*OnRampConfigSet, error) {
	event := new(OnRampConfigSet)
	if err := _OnRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampDestChainConfigSetIterator struct {
	Event *OnRampDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampDestChainConfigSet)
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
		it.Event = new(OnRampDestChainConfigSet)
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

func (it *OnRampDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *OnRampDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	DefaultCCVs       []common.Address
	LaneMandatedCCVs  []common.Address
	DefaultExecutor   common.Address
	OffRamp           []byte
	Raw               types.Log
}

func (_OnRamp *OnRampFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &OnRampDestChainConfigSetIterator{contract: _OnRamp.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampDestChainConfigSet)
				if err := _OnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error) {
	event := new(OnRampDestChainConfigSet)
	if err := _OnRamp.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampFeeTokenWithdrawnIterator struct {
	Event *OnRampFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampFeeTokenWithdrawn)
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
		it.Event = new(OnRampFeeTokenWithdrawn)
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

func (it *OnRampFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *OnRampFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &OnRampFeeTokenWithdrawnIterator{contract: _OnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampFeeTokenWithdrawn)
				if err := _OnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseFeeTokenWithdrawn(log types.Log) (*OnRampFeeTokenWithdrawn, error) {
	event := new(OnRampFeeTokenWithdrawn)
	if err := _OnRamp.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOwnershipTransferRequestedIterator struct {
	Event *OnRampOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOwnershipTransferRequested)
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
		it.Event = new(OnRampOwnershipTransferRequested)
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

func (it *OnRampOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OnRampOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOwnershipTransferRequestedIterator{contract: _OnRamp.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOwnershipTransferRequested)
				if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error) {
	event := new(OnRampOwnershipTransferRequested)
	if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OnRampOwnershipTransferredIterator struct {
	Event *OnRampOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OnRampOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnRampOwnershipTransferred)
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
		it.Event = new(OnRampOwnershipTransferred)
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

func (it *OnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OnRampOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OnRamp *OnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OnRampOwnershipTransferredIterator{contract: _OnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OnRampOwnershipTransferred)
				if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OnRamp *OnRampFilterer) ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error) {
	event := new(OnRampOwnershipTransferred)
	if err := _OnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (OnRampCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x276d7e038bc94e70aa9c54ac8cf3a3674da9252bdccbf8a0593523768f989c96")
}

func (OnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3")
}

func (OnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae")
}

func (OnRampFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (OnRampOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OnRampOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OnRamp *OnRamp) Address() common.Address {
	return _OnRamp.address
}

type OnRampInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (OnRampDestChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	ValidateDestChainAddress(opts *bind.CallOpts, rawAddress []byte, addressBytesLength uint8) ([]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*OnRampCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*OnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error)

	Address() common.Address
}
