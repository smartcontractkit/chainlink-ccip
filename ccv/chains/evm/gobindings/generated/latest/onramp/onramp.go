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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateDestChainAddress\",\"inputs\":[{\"name\":\"rawAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"validatedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102fc57604051616b5d38819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a1604051616812908161034b82396080518181816105bf015281816116500152611e0c015260a0518181816113cb0152611e45015260c051818181611e8101526127cd0152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c691461011757806314a8cfa314610112578063181f5a771461010d57806320487ded1461010857806348a98aa4146101035780635cb80c5d146100fe5780636d7fa1ce146100f95780636def4ce7146100f45780637437ff9f146100ef57806379ba5097146100ea5780638da5cb5b146100e55780639041be3d146100e057806390423fa2146100db578063df0aa9e9146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611d69565b611c75565b6112fb565b611127565b611085565b611033565b610f4a565b610e7e565b610dd8565b610bf6565b610a9d565b610a2f565b61084a565b6107a9565b610217565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576060610150611dec565b610193604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101955760043567ffffffffffffffff81116101955760040160009280601f830112156102135781359367ffffffffffffffff851161021057506020808301928560051b010111610195579190565b80fd5b8380fd5b34610195576102253661019a565b61022d613101565b60005b81811061023957005b610244818385611ed8565b9061024e82611f1d565b67ffffffffffffffff811690811580156105b3575b801561059d575b8015610584575b610549576102ba906102d4608086019161028b8388611f49565b94906102b460a08a01966102ac6102a2898d611f49565b9490923691611fb5565b923691611fb5565b906131f8565b67ffffffffffffffff166000526004602052604060002090565b9160208601906103276102e68361200b565b859073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61038461033660408901611f27565b85547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178555565b61039060608801611f3f565b6103c96001860191829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6103e06103d6858a611f49565b90600388016120d0565b6103f76103ed838a611f49565b90600288016120d0565b60c08801916104216104088461200b565b73ffffffffffffffffffffffffffffffffffffffff1690565b1561051f57600198610501846105077f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae996104ad6104f9976104656105169a61200b565b7fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff0000000083549260201b169116179055565b6104f06104e96104e360e08801936104d26104c8868b612152565b906004840161223b565b5460a01c67ffffffffffffffff1690565b9a61200b565b9a86611f49565b97909686611f49565b94909361200b565b94612152565b959094604051998a998a6123f3565b0390a201610230565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b5063ffffffff61059660608601611f3f565b1615610271565b5060ff6105ac60408601611f27565b161561026a565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610263565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610100810190811067ffffffffffffffff82111761063257604052565b6105e6565b6060810190811067ffffffffffffffff82111761063257604052565b6040810190811067ffffffffffffffff82111761063257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761063257604052565b604051906106c06101c08361066f565b565b604051906106c060a08361066f565b604051906106c060c08361066f565b67ffffffffffffffff811161063257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6040519061072960208361066f565b60008252565b60005b8381106107425750506000910152565b8181015183820152602001610732565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361078e8151809281875287808801910161072f565b0116010190565b9060206107a6928181520190610752565b90565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761082660408051906107ea818361066f565b601082527f4f6e52616d7020312e372e302d64657600000000000000000000000000000000602083015251918291602083526020830190610752565b0390f35b67ffffffffffffffff81160361019557565b908160a09103126101955790565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576004356108858161082a565b60243567ffffffffffffffff8111610195576108a590369060040161083c565b6108c38267ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff6108fa610408845473ffffffffffffffffffffffffffffffffffffffff1690565b16156109d9579081610826936109296109196080610974960184612152565b9061092386612582565b846134c4565b92610932612685565b6040840161094081866126f0565b9050610986575b5061096b604086019182516060880194610965600287519201612461565b91613c39565b90925252614292565b60405190815292839250602083019150565b6109d3915060206109b56109a36109ae6109a96109a3868b6126f0565b90612744565b61200b565b93886126f0565b01356109c6602088015161ffff1690565b9060e088015192866139f9565b38610947565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610a6960043561082a565b6020610a7f602435610a7a81610a11565b61276e565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b3461019557610aab3661019a565b90610acb60035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b818110610ad857005b610ae96104086109a9838587612834565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa908115610be6576001948891600093610bb6575b5082610b5e575b5050505001610acf565b610b69918391614c46565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610b54565b610bd891935060203d8111610bdf575b610bd0818361066f565b810190612844565b9138610b4d565b503d610bc6565b612762565b60ff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760043567ffffffffffffffff8111610195573660238201121561019557806004013567ffffffffffffffff81116101955736602482840101116101955761082691610c7e916024803592610c7884610beb565b0161291b565b60405191829182610795565b906020808351928381520192019060005b818110610ca85750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610c9b565b906107a69160208152610d0060208201835173ffffffffffffffffffffffffffffffffffffffff169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015163ffffffff166080820152608082015173ffffffffffffffffffffffffffffffffffffffff1660a082015260e0610da4610d7160a085015161010060c0860152610120850190610c8a565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610c8a565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610752565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff600435610e1c8161082a565b606060e0604051610e2c81610615565b600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201520152166000526004602052610826610e726040600020612582565b60405191829182610cd4565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610eb5611dcd565b50604051610ec281610637565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff600354166040820152604051809161082682606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760005473ffffffffffffffffffffffffffffffffffffffff81163303611009577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff6004356110c98161082a565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff811161110d5760405167ffffffffffffffff9091168152602090f35b612015565b35906106c082610a11565b8015150361019557565b346101955760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557600060405161116481610637565b60043561117081610a11565b815260243561117e8161111d565b602082019081526044359061119282610a11565b604083019182526111a1613101565b73ffffffffffffffffffffffffffffffffffffffff835116159182156112db575b5081156112d0575b506112a85780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390611293611dec565b6112a260405192839283614d46565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b9050511515386111ca565b5173ffffffffffffffffffffffffffffffffffffffff16159150386111c2565b346101955760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761133560043561082a565b60243567ffffffffffffffff81116101955761135590369060040161083c565b604435611363606435610a11565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff00000000000000000000000000000000169082015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610be657600091611c46575b50611c095760025460a01c60ff16611bdf57611452740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61147260043567ffffffffffffffff166000526004602052604060002090565b9173ffffffffffffffffffffffffffffffffffffffff6064351615611bb55782546114b361040873ffffffffffffffffffffffffffffffffffffffff831681565b3303611b8b576114c66080830183612152565b6114cf86612582565b6114dc92906004356134c4565b9060a01c67ffffffffffffffff166114f390612a57565b84547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617855590805163ffffffff16602082015161ffff166040513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015291908260348101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184526115b4908461066f565b6040517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060643560601b166020820152601481526115f360348261066f565b6115fd8780612152565b8a5460e01c60ff169061160f9261291b565b60a08601519061162260408a018a6126f0565b61162c9150612ab8565b9261163a60208b018b612152565b9590966116456106b0565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529a67ffffffffffffffff6004351660208d015267ffffffffffffffff1660408c0152600060608c015263ffffffff1660808b015261ffff1660a08a015260c08901966000885260e08a015260048c016116ca906124c2565b6101008a015261012089015261014088015261016087015261018086015236906116f3926128e4565b6101a0840152611701612685565b61170e60408601866126f0565b9050611b34575b61173561175e916040850198895190610965600260608901519201612461565b6060850152808852608084015173ffffffffffffffffffffffffffffffffffffffff1690614e0d565b9052611768612b25565b61179961178a61177b8487600435614292565b5063ffffffff16606087015290565b80602084015285600435614e9b565b6117a660408501856126f0565b9050611a80575b926117b7836159b2565b808552602081519101206117cc875151612b96565b9360408601948552606060009301925b885180518210156119d25760206118196104086104086117ff8661186096612b82565b5173ffffffffffffffffffffffffffffffffffffffff1690565b6118278460608a0151612b82565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612bfd565b03915afa8015610be65773ffffffffffffffffffffffffffffffffffffffff916000916119a4575b50168015611941579060006118eb926118a08761200b565b908b836118b18660608d0151612b82565b5193604051978895869485937f3bbbed4b0000000000000000000000000000000000000000000000000000000085528d8d60048701612d66565b03925af18015610be6578161191991600194600091611920575b508951906119138383612b82565b52612b82565b50016117dc565b61193b913d8091833e611933818361066f565b810190612c60565b38611905565b6105806119526117ff848d51612b82565b7f83c758a60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660049081523567ffffffffffffffff16602452604490565b6119c5915060203d81116119cb575b6119bd818361066f565b81019061274d565b38611888565b503d6119b3565b6108268480898b7f276d7e038bc94e70aa9c54ac8cf3a3674da9252bdccbf8a0593523768f989c9667ffffffffffffffff611a1860408b015167ffffffffffffffff1690565b611a40602085519501519551604051938493169667ffffffffffffffff600435169684613010565b0390a4611a707fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b6001611a8f60408601866126f0565b905003611b0a57611af7611ae1611aac6109a360408801886126f0565b60c0850151805115611afd57905b602086015161ffff169060e08701519260643591611adc600435913690612b44565b6151ad565b61018085015190611af182612b75565b52612b75565b506117ad565b5061014086015190611aba565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b5061175e611735611b83611b516109a96109a360408a018a6126f0565b6020611b636109a360408b018b6126f0565b0135611b74602088015161ffff1690565b9060e0880151926004356139f9565b915050611715565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105806004359067ffffffffffffffff60249216600452565b611c68915060203d602011611c6e575b611c60818361066f565b810190612a42565b386113fc565b503d611c56565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955773ffffffffffffffffffffffffffffffffffffffff600435611cc581610a11565b611ccd613101565b16338114611d3f57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557611da360043561082a565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611dda82610637565b60006040838281528260208201520152565b611df4611dcd565b50604051611e0181610637565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611f185760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181360301821215610195570190565b611ea9565b356107a68161082a565b356107a681610beb565b63ffffffff81160361019557565b356107a681611f31565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160051b3603831361019557565b67ffffffffffffffff81116106325760051b60200190565b929190611fc181611f9d565b93611fcf604051958661066f565b602085838152019160051b810192831161019557905b828210611ff157505050565b60208091833561200081610a11565b815201910190611fe5565b356107a681610a11565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e6400000000149015171561110d57565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81160361110d57565b8181029291811591840414171561110d57565b8181106120c4575050565b600081556001016120b9565b9067ffffffffffffffff831161063257680100000000000000008311610632578154838355808410612134575b5090600052602060002060005b8381106121175750505050565b600190602084359461212886610a11565b0193818401550161210a565b61214c908360005284602060002091820191016120b9565b386120fd565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff82116101955760200191813603831361019557565b90600182811c921680156121ec575b60208310146121bd57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916121b2565b9190601f811161220557505050565b6106c0926000526020600020906020601f840160051c83019310612231575b601f0160051c01906120b9565b9091508190612224565b90929167ffffffffffffffff8111610632576122618161225b84546121a3565b846121f6565b6000601f82116001146122bf5781906122b09394956000926122b4575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b01359050388061227e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216946122f284600052602060002090565b91805b87811061234b575083600195969710612313575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612309565b909260206001819286860135815501940191016122f5565b9160209082815201919060005b81811061237d5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356123a681610a11565b168152019401929101612370565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b949161244d9373ffffffffffffffffffffffffffffffffffffffff958661243f9367ffffffffffffffff6107a69e9c9d9b96168a5216602089015260c0604089015260c0880191612363565b918583036060870152612363565b9416608082015260a08185039101526123b4565b906040519182815491828252602082019060005260206000209260005b8181106124935750506106c09250038361066f565b845473ffffffffffffffffffffffffffffffffffffffff1683526001948501948794506020909301920161247e565b90604051918260008254926124d6846121a3565b808452936001811690811561254257506001146124fb575b506106c09250038361066f565b90506000929192526020600020906000915b8183106125265750509060206106c092820101386124ee565b602091935080600191548385890101520191019091849261250d565b602093506106c09592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386124ee565b9060405161258f81610615565b60e0612680600483956125f76125ed825473ffffffffffffffffffffffffffffffffffffffff8082161688526125e46125d38267ffffffffffffffff9060a01c1690565b67ffffffffffffffff1660208a0152565b60e01c60ff1690565b60ff166040870152565b61265861263b600183015461261f6126128263ffffffff1690565b63ffffffff1660608a0152565b60201c73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166080870152565b61266460028201612461565b60a086015261267560038201612461565b60c0860152016124c2565b910152565b6040519061269460208361066f565b6000808352366020840137565b906126ab82611f9d565b6126b8604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06126e68294611f9d565b0190602036910137565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160061b3603831361019557565b9015611f185790565b9081602091031261019557516107a681610a11565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015610be65773ffffffffffffffffffffffffffffffffffffffff9160009161281757501690565b612830915060203d6020116119cb576119bd818361066f565b1690565b9190811015611f185760051b0190565b90816020910312610195575190565b9160206107a69381815201916123b4565b60ff166020039060ff821161110d57565b909291928311610195579190565b906004116101955790600490565b90939293848311610195578411610195578101920390565b3590602081106128b7575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b9291926128f0826106e0565b916128fe604051938461066f565b829481845281830111610195578281602093846000960137010152565b9160ff81169060208210612977575b50810361293d57906107a69136916128e4565b906129736040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612853565b0390fd5b60208311612a0d576020830361292a5790506129ab6129a560ff61299d84969596612864565b168585612875565b906128a9565b6129d757916129d091816129ca6129c46107a696612864565b60ff1690565b91612891565b36916128e4565b506129736040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612853565b6040517f3aeba39000000000000000000000000000000000000000000000000000000000815280612973858760048401612853565b9081602091031261019557516107a68161111d565b67ffffffffffffffff1667ffffffffffffffff811461110d5760010190565b6040519060c0820182811067ffffffffffffffff82111761063257604052606060a0836000815282602082015282604082015282808201528260808201520152565b90612ac282611f9d565b612acf604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612afd8294611f9d565b019060005b828110612b0e57505050565b602090612b19612a76565b82828501015201612b02565b60405190612b3282610637565b60606040838281528260208201520152565b919082604091031261019557604051612b5c81610653565b60208082948035612b6c81610a11565b84520135910152565b805115611f185760200190565b8051821015611f185760209160051b010190565b90612ba082611f9d565b612bad604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612bdb8294611f9d565b019060005b828110612bec57505050565b806060602080938501015201612be0565b60409067ffffffffffffffff6107a694931681528160208201520190610752565b81601f82011215610195578051612c34816106e0565b92612c42604051948561066f565b81845260208284010111610195576107a6916020808501910161072f565b9060208282031261019557815167ffffffffffffffff8111610195576107a69201612c1e565b9080602083519182815201916020808360051b8301019401926000915b838310612cb257505050505090565b9091929394602080612d57837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a0612d46612d34612d22612d108887015160c08a88015260c0870190610752565b60408701518682036040880152610752565b60608601518582036060870152610752565b60808501518482036080860152610752565b9201519060a0818403910152610752565b97019301930191939290612ca3565b9193906107a69593612f62612f879260a08652612d9060a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612f2f612ef9612ec3612e8d612e57612e238c61026060e08a0151916101c06101808201520190610752565b6101008801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001888f0152610752565b6101208701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101c08e0152610752565b6101408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608c8303016101e08d0152610752565b6101608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8303016102008c0152610752565b6101808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a8303016102208b0152612c86565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6087830301610240880152610752565b956020850152604084019073ffffffffffffffffffffffffffffffffffffffff169052565b60608201526080818403910152610752565b9080602083519182815201916020808360051b8301019401926000915b838310612fc557505050505090565b9091929394602080613001837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610752565b97019301930191939290612fb6565b93929061302590606086526060860190610752565b938085036020820152825180865260208601906020808260051b8901019501916000905b82821061306757505050506107a69394506040818403910152612f99565b909192956020806130f3837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001960301865260a060808c5173ffffffffffffffffffffffffffffffffffffffff815116845263ffffffff86820151168685015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610752565b980192019201909291613049565b73ffffffffffffffffffffffffffffffffffffffff60015416330361312257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b906001820180921161110d57565b600101908160011161110d57565b906014820180921161110d57565b90600c820180921161110d57565b9190820180921161110d57565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161110d57565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe820191821161110d57565b9190820391821161110d57565b805191613206815184613184565b92831561335d5760005b84811061321e575050505050565b81811015613342576132336117ff8286612b82565b73ffffffffffffffffffffffffffffffffffffffff81168015613318576132598361314c565b87811061326b57505050600101613210565b848110156132e85773ffffffffffffffffffffffffffffffffffffffff6132956117ff838a612b82565b1682146132a457600101613259565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6133136117ff61330d88856131eb565b89612b82565b613295565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6133586117ff61335284846131eb565b85612b82565b613233565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061339482610615565b606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106133fa575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261019557825167ffffffffffffffff81116101955782613454918501612c1e565b92602081015161346381611f31565b92604082015167ffffffffffffffff8111610195576107a69201612c1e565b60409067ffffffffffffffff6107a6959316815281602082015201916123b4565b9060ff6134bd602092959495604085526040850190610752565b9416910152565b929190926134d0613387565b600484101580613892575b15613758575050906134ec91615d7a565b9060c082015180516136e4575b506040820180515160005b81811061363557505080515115613598575b505b6080820173ffffffffffffffffffffffffffffffffffffffff61354f825173ffffffffffffffffffffffffffffffffffffffff1690565b161561355a57505090565b61357e60806107a693015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b9060c08194939401926135ac8451516126a1565b83526135b9845151612b96565b946060810195865260005b8551805182101561362757906136086135e26117ff83600195612b82565b6135ed838951612b82565b9073ffffffffffffffffffffffffffffffffffffffff169052565b61362081895161361661071a565b6119138383612b82565b50016135c4565b505093509350905038613516565b61363e8161314c565b82811061364e5750600101613504565b61365c6117ff838651612b82565b73ffffffffffffffffffffffffffffffffffffffff6136826104086117ff858951612b82565b9116146136915760010161363e565b6105806136a26117ff848751612b82565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b600061372d916136f8604085015160ff1690565b9060405193849283927f6d7fa1ce000000000000000000000000000000000000000000000000000000008452600484016134a3565b0381305afa8015610be657156134f957613751903d806000833e611933818361066f565b50386134f9565b60c085969296019461376b8651516126a1565b946040830195865261377e875151612b96565b976060840198895260005b885180518210156137c757906137b26137a76117ff83600195612b82565b6135ed838c51612b82565b6137c0818c5161361661071a565b5001613789565b50509195509195506000929650613834936137fd61040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b91604051958694859384937f9cc1999600000000000000000000000000000000000000000000000000000000855260048501613482565b03915afa8015610be657600090600090600090613866575b60a086015263ffffffff16845290505b60c0830152613518565b50505061388861385c913d806000833e613880818361066f565b81019061342c565b919250829161384c565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006138e86138e28787612883565b906133c6565b16146134db565b6020818303126101955780519067ffffffffffffffff821161019557019080601f8301121561019557815161392381611f9d565b92613931604051948561066f565b81845260208085019260051b82010192831161019557602001905b8282106139595750505090565b60208091835161396881610a11565b81520191019061394c565b95949060009460a09467ffffffffffffffff6139c79573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c0860190610752565b930152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461110d5760010190565b949294939093613a286003613a228367ffffffffffffffff166000526004602052604060002090565b01612461565b9473ffffffffffffffffffffffffffffffffffffffff613a4981831661276e565b16926040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180613aa660048201907fdc0cbd3600000000000000000000000000000000000000000000000000000000602083019252565b0381885afa908115610be657600091613c1a575b5015613c105790613b00600095949392604051998a96879586957f89720a6200000000000000000000000000000000000000000000000000000000875260048701613973565b03915afa928315610be657600093613beb575b50825115613be657613b30613b2b8451845190613184565b6126a1565b6000918293835b8651811015613b9557613b4d6117ff8289612b82565b73ffffffffffffffffffffffffffffffffffffffff811615613b895790613b836001926135ed613b7c896139cc565b9888612b82565b01613b37565b50945060018095613b83565b509193909450613ba6575b50815290565b60005b8151811015613bde5780613bd8613bc56117ff60019486612b82565b6135ed613bd1876139cc565b9688612b82565b01613ba9565b505038613ba0565b915090565b613c099193503d806000833e613c01818361066f565b8101906138ef565b9138613b13565b5050505050915090565b613c33915060203d602011611c6e57611c60818361066f565b38613aba565b93919293613c55613c4d8251865190613184565b865190613184565b90613c68613c62836126a1565b92612b96565b94600096875b8351891015613cce5788613cc4613cb7600193613c9f613c956117ff8e9f9d9e9d8b612b82565b6135ed838c612b82565b613cbd613cac858c612b82565b5191809384916139cc565b9c612b82565b528b612b82565b5001979695613c6e565b959250929350955060005b8651811015613d6857613cef6117ff8289612b82565b600073ffffffffffffffffffffffffffffffffffffffff8216815b888110613d3c575b5050906001929115613d26575b5001613cd9565b613d36906135ed613b7c896139cc565b38613d1f565b81613d4d6104086117ff848c612b82565b14613d5a57600101613d0a565b506001915081905038613d12565b509390945060005b8551811015613e0657613d866117ff8288612b82565b600073ffffffffffffffffffffffffffffffffffffffff8216815b878110613dda575b5050906001929115613dbd575b5001613d70565b613dd4906135ed613dcd886139cc565b9787612b82565b38613db6565b81613deb6104086117ff848b612b82565b14613df857600101613da1565b506001915081905038613da9565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176106325760405260606080836000815260006020820152600060408201526000838201520152565b90613e5b82611f9d565b613e68604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613e968294611f9d565b019060005b828110613ea757505050565b602090613eb2613e12565b82828501015201613e9b565b519061ffff8216820361019557565b9081606091031261019557613ee181613ebe565b9160406020830151613ef281611f31565b9201516107a681611f31565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561019557016020813591019167ffffffffffffffff821161019557813603831361019557565b9160209082815201919060005b818110613f685750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735613f9181610a11565b16815260208781013590820152019401929101613f5b565b949391929067ffffffffffffffff16855260806020860152614020613fe3613fd18580613efe565b60a060808a01526101208901916123b4565b613ff06020860186613efe565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808984030160a08a01526123b4565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215610195578401916020833593019167ffffffffffffffff8411610195578360061b36038313610195576106c0956140f16140bb83606097614130978d60c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80826141229a0301910152613f4e565b916140e76140ca888301611112565b73ffffffffffffffffffffffffffffffffffffffff1660e08d0152565b6080810190613efe565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101008c01526123b4565b908782036040890152610752565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff821161110d57565b908160a091031261019557805191602082015161417181611f31565b91604081015161418081611f31565b91608061418f60608401613ebe565b9201516107a68161111d565b9260c09473ffffffffffffffffffffffffffffffffffffffff9167ffffffffffffffff61ffff95846107a69b9a9616885216602087015260408601521660608401521660808201528160a08201520190610752565b90816060910312610195578051613ee181611f31565b9190826040910312610195576020825161421f81611f31565b92015190565b9081602091031261019557517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff811681036101955790565b8115614263570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b929190926000926000936040840192835151946142cf6142ca6142c560408b01986142bd8a8d6126f0565b919050613184565b61314c565b613e51565b956000946000945b8751805187101561455a576104086104086117ff896142f594612b82565b6143416020606087019261430a8a8551612b82565b519060405180809581947f958021a70000000000000000000000000000000000000000000000000000000083528d60048401612bfd565b03915afa8015610be65773ffffffffffffffffffffffffffffffffffffffff9160009161453c575b501680156144db579060608d93926143828a8451612b82565b519061439360208a015161ffff1690565b958a6143ce604051988995869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613fa9565b03915afa8015610be657600193614478938a8e8e95600080958197614481575b509083929161ffff6144168561440f6117ff61446c996144729d9e51612b82565b9451612b82565b519161443f6144236106c2565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b63ffffffff8916602085015263ffffffff8b16604085015216606083015260808201526119138383612b82565b5061413b565b9861413b565b950194956142d7565b61447297506117ff965084939291509361ffff6144168261440f6144be61446c9960603d81116144d4575b6144b6818361066f565b810190613ecd565b9c9196909c9d50505050505050909192936143ee565b503d6144ac565b610580876144ed6117ff8b8e51612b82565b7f83c758a60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045267ffffffffffffffff16602452604490565b614554915060203d81116119cb576119bd818361066f565b38614369565b509192939599945097969561456f83866126f0565b90506148a7575b506145f76145e86040926145b18661458e8b51613191565b966145a861459f60208c018c612152565b9290508b6126f0565b9190508b61602c565b6145bb868b612b82565b526145c6858a612b82565b506145e260206145d6878c612b82565b51015163ffffffff1690565b9061413b565b996145e2836145d6868b612b82565b9861461d61040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b82517f1f4ea29d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff97909716600488015263ffffffff91821660248801529916604486015293979388606481875afa918215610be65760009860009361482a575b5060209361473f95936146ed9373eba517d20000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff6146de6080606097015173ffffffffffffffffffffffffffffffffffffffff1690565b1603614809575b50500161200b565b9060405180809581947f4ab35b0b0000000000000000000000000000000000000000000000000000000083526004830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b03915afa8015610be6577bffffffffffffffffffffffffffffffffffffffffffffffffffffffff916000916147da575b5016936000925b82518410156147d1576147c96001916147a5886147a060606147988a8a612b82565b510151612044565b614259565b60606147b18888612b82565b51015260606147c08787612b82565b51015190613184565b930192614776565b91945092909150565b6147fc915060203d602011614802575b6147f4818361066f565b810190614225565b3861476f565b503d6147ea565b83614817614821928b612b82565b5101918251613184565b905238806146e5565b6080995061473f959350936060916146ed9373eba517d20000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff6146de61488c60209a60403d6040116148a0575b614884818361066f565b810190614206565b9f909f999b50505050509350915093614684565b503d61487a565b6148bd610a7a6104086109a96109a3878a6126f0565b9060009051519173ffffffffffffffffffffffffffffffffffffffff6148e96109a96109a3888b6126f0565b9160e088019283516149186148fc6106c2565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b8560208301528560408301528560608301526080820152614939868d612b82565b52614944858c612b82565b5016906040517f01ffc9a7000000000000000000000000000000000000000000000000000000008152602081806149a260048201907fdc0cbd3600000000000000000000000000000000000000000000000000000000602083019252565b0381865afa928315610be6578a888a958c948891614c27575b50614b24575b5050505050506145765790614a699060606149f761040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b614a076109a96109a3888b6126f0565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b16600482015273ffffffffffffffffffffffffffffffffffffffff909116602482015293849190829081906044820190565b03915afa908115610be65760409363ffffffff6145e8936145f7958c600092600092600091614ae7575b5090614ad1614adc93928b614ac46060614aad8b87612b82565b5101996020614abc8288612b82565b510195612b82565b51019063ffffffff169052565b9063ffffffff169052565b169052925050614576565b614ad19450614adc9350614b13915060603d606011614b1d575b614b0b818361066f565b8101906141f0565b9194909350614a93565b503d614b01565b614b9f60a095614b666020614b5c606082614b536109a3614b4c6109a96109a38f8d906126f0565b998d6126f0565b0135990161200b565b99015161ffff1690565b905190604051988997889687967f2c0634040000000000000000000000000000000000000000000000000000000088526004880161419b565b03915afa918215610be657809181908294614beb575b50614bdf908b614ad16060614bca8984612b82565b5101946040614ac48a6020614abc8288612b82565b52843887818a886149c1565b915050614bdf9250614c15915060a03d60a011614c20575b614c0d818361066f565b810190614155565b949192919050614bb5565b503d614c03565b614c40915060203d602011611c6e57611c60818361066f565b386149bb565b9073ffffffffffffffffffffffffffffffffffffffff614d189392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452614ca860648561066f565b16600080604093845195614cbc868861066f565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15614d3d573d614d09614d00826106e0565b9451948561066f565b83523d6000602085013e616741565b805180614d23575050565b81602080614d38936106c09501019101612a42565b61627a565b60609250616741565b9160606106c0929493614d938160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b90614dd2826106e0565b614ddf604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06126e682946106e0565b91825160148102908082046014149015171561110d57614e2f614e349161315a565b613168565b90614e46614e4183613176565b614dc8565b906014614e5283612b75565b5360009260215b8651851015614e84576014600191614e746117ff888b612b82565b60601b8187015201940193614e59565b919550936020935060601b90820152828152012090565b5090614eac6104086060840161200b565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060408101614edd81836126f0565b9050614f5e575b505060005b8251811015614f5857806060614f0160019386612b82565b51015180158015614f4f575b614f4957614f4390614f3d614f228488612b82565b515173ffffffffffffffffffffffffffffffffffffffff1690565b87614c46565b01614ee9565b50614f43565b50838214614f0d565b50505050565b614f84919250610a7a6104086109a96109a361040894614f7e89516131be565b976126f0565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdc0cbd3600000000000000000000000000000000000000000000000000000000600482015260208160248173ffffffffffffffffffffffffffffffffffffffff86165afa908115610be65760009161502a575b5061500a575b80614ee4565b61502490606061501a8486612b82565b5101519085614c46565b38615004565b615043915060203d602011611c6e57611c60818361066f565b38614ffe565b6040519061505682610653565b60606020838281520152565b9190604083820312610195576040519061507b82610653565b8193805167ffffffffffffffff8111610195578261509a918301612c1e565b835260208101519167ffffffffffffffff8311610195576020926126809201612c1e565b9060208282031261019557815167ffffffffffffffff8111610195576107a69201615062565b90608073ffffffffffffffffffffffffffffffffffffffff81615110855160a0865260a0860190610752565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b9060206107a69281815201906150e4565b919060408382031261019557825167ffffffffffffffff81116101955760209161421f918501615062565b61ffff6151996107a695936060845260608401906150e4565b931660208201526040818403910152610752565b909192936151b9612a76565b50602082019081511561574d576151ed610408610a7a610408865173ffffffffffffffffffffffffffffffffffffffff1690565b9573ffffffffffffffffffffffffffffffffffffffff871692831580156156c2575b61565f5761529a815191615221615049565b5051865173ffffffffffffffffffffffffffffffffffffffff16906152786152476106c2565b8b815267ffffffffffffffff8b1660208201529573ffffffffffffffffffffffffffffffffffffffff166040870152565b606085015273ffffffffffffffffffffffffffffffffffffffff166080840152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdc0cbd36000000000000000000000000000000000000000000000000000000006004820152602081602481885afa908115610be657600091615640575b501561554a57509161534696979160008094604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501615180565b03925af18015610be657600094600091615501575b509460006154a09361543e6153da615412956117ff6153ae9a965b6040519b8c91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018c528b61066f565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810186528561066f565b61546b61546184519267ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b9060405195869283927f6d7fa1ce000000000000000000000000000000000000000000000000000000008452600484016134a3565b0381305afa928315610be6576000936154e1575b5060200151936154c26106d1565b958652602086015260408501526060840152608083015260a082015290565b60209193506154fa903d806000833e611933818361066f565b92906154b4565b6153ae95506117ff9691506154a09361543e6153da615412956155376000953d8088833e61552f818361066f565b810190615155565b9b909b969b9a505095505050935061535b565b979161ffff91959793501661561657516155ec57600061559793604051809581927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301615144565b038183855af1918215610be6576153ae9561543e6153da6000936117ff615412976154a09987916155ca575b5096615376565b6155e691503d8089833e6155de818361066f565b8101906150be565b386155c3565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b615659915060203d602011611c6e57611c60818361066f565b386152fe565b610580615680865173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818c5afa908115610be65760009161572e575b501561520f565b615747915060203d602011611c6e57611c60818361066f565b38615727565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592615849947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b906158656020928281519485920161072f565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b1681526158b1825180936020898501910161072f565b019160f81b16838201526158cf82518093602060028501910161072f565b01019160f81b16838201526158ee82518093602060028501910161072f565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff000000000000000000000000000000000000000000000000000000000000006107a69f9e9c9860f81b168152615963825180936020898501910161072f565b019160f01b168382015261598182518093602060038501910161072f565b01019160f01b168382015261599f825180936020898501910161072f565b01019160f01b1660028201520190615852565b60e081019060ff82515111615d4b5761010081019060ff82515111615d1c5761012081019260ff84515111615ced5761014082019060ff82515111615cbe5761016083019461ffff86515111615c8f57610180840194600186515111615c60576101a085019261ffff84515111615c2f57855167ffffffffffffffff16602087015167ffffffffffffffff16906040880151615a559067ffffffffffffffff1690565b976060810151615a689063ffffffff1690565b906080810151615a7b9063ffffffff1690565b60a082015161ffff169160c00151926040519b8c966020880196615a9e97615777565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752615ace908761066f565b51908151615adc9060ff1690565b9051805160ff169351908151615af29060ff1690565b906040519586956020870195615b0796615869565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252615b37908261066f565b606094519081511515615be197615ba56107a698615be797615be196615bd195615c13575b505196615b6a885160ff1690565b935191615b79835161ffff1690565b91615b86825161ffff1690565b905193615b95855161ffff1690565b936040519b8c9860208a016158f4565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810185528461066f565b6040519687956020870190615852565b90615852565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261066f565b615c28919250615c2290612b75565b51616440565b9038615b5c565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526105806024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b90615d83613387565b9160118210615fa65780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821603615f335750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a615e02816126a1565b60408601908152615e1282612b96565b906060870191825260005b838110615ee75750505050615e928383615e88615e7c615e72615e6b615e4b615e9c9887615ea69c9b6165d6565b73ffffffffffffffffffffffffffffffffffffffff90911660808d015290565b85856166ac565b92919036916128e4565b60a08a01528383616714565b94919036916128e4565b60c08801526166ac565b93919036916128e4565b60e08401528103615eb5575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b80600191615f2c615f16615f0f615f02615f269a8d8d6165d6565b91906135ed868a51612b82565b8b8b6166ac565b9391889a919a51949a36916128e4565b92612b82565b5201615e1d565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b90816020910312610195576107a690613ebe565b9261ffff6107a6959367ffffffffffffffff61601e94168652166020850152608060408501526080840190610c8a565b916060818403910152610752565b9091616036613e12565b506160558267ffffffffffffffff166000526004602052604060002090565b93616065855460ff9060e01c1690565b9061611f6161046160fb60808401946160e76160c16160b96160ac60016160a08b5173ffffffffffffffffffffffffffffffffffffffff1690565b9e015463ffffffff1690565b885163ffffffff166145e2565b9a6075613184565b976160f56160ed60ff6160db60a08b019c8d515190613184565b9516946160e786612076565b90613184565b93604f613184565b906120a6565b63ffffffff1690565b925173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811673eba517d200000000000000000000000000000000036161b357505061ffff92506161a5906161986000935b519561618b61616f6106c2565b73ffffffffffffffffffffffffffffffffffffffff9099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b906161d661040860209373ffffffffffffffffffffffffffffffffffffffff1690565b60406161e68484015161ffff1690565b92015191855196616226604051988995869485947fe962e69e00000000000000000000000000000000000000000000000000000000865260048601615fee565b03915afa908115610be6576161986161a59261ffff9560009161624b575b5093616162565b61626d915060203d602011616273575b616265818361066f565b810190615fda565b38616244565b503d61625b565b1561628157565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b966002616414976163e160226107a69f9e9c9799600199859f9b7fff00000000000000000000000000000000000000000000000000000000000000906163e19f826163e19c6163e89c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152616392825180936020898501910161072f565b019160f81b16838201526163b082518093602060238501910161072f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190615852565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff825151116165a757604081019160ff8351511161657857606082019160ff8351511161654957608081019260ff8451511161651a5760a0820161ffff815151116164eb576107a694615be793519451916164a2835160ff1690565b9751916164b0835160ff1690565b9451906164be825160ff1690565b9051936164cc855160ff1690565b9351966164db885161ffff1690565b966040519c8d9b60208d01616305565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602560045260246000fd5b92919092600182019184831161667a5781013560001a82811561666f57506014810361664257820193841161660e57013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b9190600282019181831161667a578381013560f01c01600201928184116166e0579183916166d993612891565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b9190600182019181831161667a578381013560001a01600101928184116166e0579183916166d993612891565b919290156167bc5750815115616755575090565b3b1561675e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156167cf5750805190602001fd5b612973906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161079556fea164736f6c634300081a000a",
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
