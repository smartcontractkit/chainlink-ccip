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
	Bin: "0x60e0604052346102fc57604051616a8b38819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a1604051616740908161034b82396080518181816105bf015281816116500152611e09015260a0518181816113cb0152611e42015260c051818181611e7e01526127f80152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c691461011757806314a8cfa314610112578063181f5a771461010d57806320487ded1461010857806348a98aa4146101035780635cb80c5d146100fe5780636d7fa1ce146100f95780636def4ce7146100f45780637437ff9f146100ef57806379ba5097146100ea5780638da5cb5b146100e55780639041be3d146100e057806390423fa2146100db578063df0aa9e9146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611d66565b611c72565b6112fb565b611127565b611085565b611033565b610f4a565b610e7e565b610dd8565b610bf6565b610a9d565b610a2f565b61084a565b6107a9565b610217565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576060610150611de9565b610193604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101955760043567ffffffffffffffff81116101955760040160009280601f830112156102135781359367ffffffffffffffff851161021057506020808301928560051b010111610195579190565b80fd5b8380fd5b34610195576102253661019a565b61022d61312c565b60005b81811061023957005b610244818385611ed5565b9061024e82611f1a565b67ffffffffffffffff811690811580156105b3575b801561059d575b8015610584575b610549576102ba906102d4608086019161028b8388611f46565b94906102b460a08a01966102ac6102a2898d611f46565b9490923691611fb2565b923691611fb2565b90613223565b67ffffffffffffffff166000526004602052604060002090565b9160208601906103276102e683612008565b859073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61038461033660408901611f24565b85547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178555565b61039060608801611f3c565b6103c96001860191829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6103e06103d6858a611f46565b90600388016120fb565b6103f76103ed838a611f46565b90600288016120fb565b60c088019161042161040884612008565b73ffffffffffffffffffffffffffffffffffffffff1690565b1561051f57600198610501846105077f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae996104ad6104f9976104656105169a612008565b7fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff0000000083549260201b169116179055565b6104f06104e96104e360e08801936104d26104c8868b61217d565b9060048401612266565b5460a01c67ffffffffffffffff1690565b9a612008565b9a86611f46565b97909686611f46565b949093612008565b9461217d565b959094604051998a998a61241e565b0390a201610230565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b5063ffffffff61059660608601611f3c565b1615610271565b5060ff6105ac60408601611f24565b161561026a565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610263565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610100810190811067ffffffffffffffff82111761063257604052565b6105e6565b6060810190811067ffffffffffffffff82111761063257604052565b6040810190811067ffffffffffffffff82111761063257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761063257604052565b604051906106c06101c08361066f565b565b604051906106c060a08361066f565b604051906106c060c08361066f565b67ffffffffffffffff811161063257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6040519061072960208361066f565b60008252565b60005b8381106107425750506000910152565b8181015183820152602001610732565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361078e8151809281875287808801910161072f565b0116010190565b9060206107a6928181520190610752565b90565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761082660408051906107ea818361066f565b601082527f4f6e52616d7020312e372e302d64657600000000000000000000000000000000602083015251918291602083526020830190610752565b0390f35b67ffffffffffffffff81160361019557565b908160a09103126101955790565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576004356108858161082a565b60243567ffffffffffffffff8111610195576108a590369060040161083c565b6108c38267ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff6108fa610408845473ffffffffffffffffffffffffffffffffffffffff1690565b16156109d957908161082693610929610919608061097496018461217d565b90610923866125ad565b846134ef565b926109326126b0565b60408401610940818661271b565b9050610986575b5061096b60408601918251606088019461096560028751920161248c565b91613c64565b90925252614295565b60405190815292839250602083019150565b6109d3915060206109b56109a36109ae6109a96109a3868b61271b565b9061276f565b612008565b938861271b565b01356109c6602088015161ffff1690565b9060e08801519286613a24565b38610947565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610a6960043561082a565b6020610a7f602435610a7a81610a11565b612799565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b3461019557610aab3661019a565b90610acb60035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b818110610ad857005b610ae96104086109a983858761285f565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa908115610be6576001948891600093610bb6575b5082610b5e575b5050505001610acf565b610b69918391614b7f565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610b54565b610bd891935060203d8111610bdf575b610bd0818361066f565b81019061286f565b9138610b4d565b503d610bc6565b61278d565b60ff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760043567ffffffffffffffff8111610195573660238201121561019557806004013567ffffffffffffffff81116101955736602482840101116101955761082691610c7e916024803592610c7884610beb565b01612946565b60405191829182610795565b906020808351928381520192019060005b818110610ca85750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610c9b565b906107a69160208152610d0060208201835173ffffffffffffffffffffffffffffffffffffffff169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015163ffffffff166080820152608082015173ffffffffffffffffffffffffffffffffffffffff1660a082015260e0610da4610d7160a085015161010060c0860152610120850190610c8a565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610c8a565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610752565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff600435610e1c8161082a565b606060e0604051610e2c81610615565b600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201520152166000526004602052610826610e7260406000206125ad565b60405191829182610cd4565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610eb5611dca565b50604051610ec281610637565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff600354166040820152604051809161082682606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760005473ffffffffffffffffffffffffffffffffffffffff81163303611009577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff6004356110c98161082a565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff811161110d5760405167ffffffffffffffff9091168152602090f35b612012565b35906106c082610a11565b8015150361019557565b346101955760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557600060405161116481610637565b60043561117081610a11565b815260243561117e8161111d565b602082019081526044359061119282610a11565b604083019182526111a161312c565b73ffffffffffffffffffffffffffffffffffffffff835116159182156112db575b5081156112d0575b506112a85780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390611293611de9565b6112a260405192839283614c7f565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b9050511515386111ca565b5173ffffffffffffffffffffffffffffffffffffffff16159150386111c2565b346101955760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761133560043561082a565b60243567ffffffffffffffff81116101955761135590369060040161083c565b604435611363606435610a11565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff00000000000000000000000000000000169082015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610be657600091611c43575b50611c065760025460a01c60ff16611bdc57611452740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61147260043567ffffffffffffffff166000526004602052604060002090565b9173ffffffffffffffffffffffffffffffffffffffff6064351615611bb25782546114b361040873ffffffffffffffffffffffffffffffffffffffff831681565b3303611b88576114c6608083018361217d565b6114cf866125ad565b6114dc92906004356134ef565b9060a01c67ffffffffffffffff166114f390612a82565b84547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617855590805163ffffffff16602082015161ffff166040513060601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602082015291908260348101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184526115b4908461066f565b6040517fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060643560601b166020820152601481526115f360348261066f565b6115fd878061217d565b8a5460e01c60ff169061160f92612946565b60a08601519061162260408a018a61271b565b61162c9150612ae3565b9261163a60208b018b61217d565b9590966116456106b0565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529a67ffffffffffffffff6004351660208d015267ffffffffffffffff1660408c0152600060608c015263ffffffff1660808b015261ffff1660a08a015260c08901966000885260e08a015260048c016116ca906124ed565b6101008a015261012089015261014088015261016087015261018086015236906116f39261290f565b6101a08401526117016126b0565b61170e604086018661271b565b9050611b31575b61173561175e91604085019889519061096560026060890151920161248c565b6060850152808852608084015173ffffffffffffffffffffffffffffffffffffffff1690614d46565b9052611768612b50565b61179661178a61177b8487600435614295565b5063ffffffff16606087015290565b80602084015285614dd4565b6117a3604085018561271b565b9050611a7d575b926117b4836158df565b808552602081519101206117c9875151612bc1565b9360408601948552606060009301925b885180518210156119cf5760206118166104086104086117fc8661185d96612bad565b5173ffffffffffffffffffffffffffffffffffffffff1690565b6118248460608a0151612bad565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612c28565b03915afa8015610be65773ffffffffffffffffffffffffffffffffffffffff916000916119a1575b5016801561193e579060006118e89261189d87612008565b908b836118ae8660608d0151612bad565b5193604051978895869485937f3bbbed4b0000000000000000000000000000000000000000000000000000000085528d8d60048701612d91565b03925af18015610be657816119169160019460009161191d575b508951906119108383612bad565b52612bad565b50016117d9565b611938913d8091833e611930818361066f565b810190612c8b565b38611902565b61058061194f6117fc848d51612bad565b7f83c758a60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660049081523567ffffffffffffffff16602452604490565b6119c2915060203d81116119c8575b6119ba818361066f565b810190612778565b38611885565b503d6119b0565b6108268480898b7f276d7e038bc94e70aa9c54ac8cf3a3674da9252bdccbf8a0593523768f989c9667ffffffffffffffff611a1560408b015167ffffffffffffffff1690565b611a3d602085519501519551604051938493169667ffffffffffffffff60043516968461303b565b0390a4611a6d7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b6001611a8c604086018661271b565b905003611b0757611af4611ade611aa96109a3604088018861271b565b60c0850151805115611afa57905b602086015161ffff169060e08701519260643591611ad9600435913690612b6f565b6150da565b61018085015190611aee82612ba0565b52612ba0565b506117aa565b5061014086015190611ab7565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b5061175e611735611b80611b4e6109a96109a360408a018a61271b565b6020611b606109a360408b018b61271b565b0135611b71602088015161ffff1690565b9060e088015192600435613a24565b915050611715565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105806004359067ffffffffffffffff60249216600452565b611c65915060203d602011611c6b575b611c5d818361066f565b810190612a6d565b386113fc565b503d611c53565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955773ffffffffffffffffffffffffffffffffffffffff600435611cc281610a11565b611cca61312c565b16338114611d3c57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557611da060043561082a565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611dd782610637565b60006040838281528260208201520152565b611df1611dca565b50604051611dfe81610637565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611f155760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181360301821215610195570190565b611ea6565b356107a68161082a565b356107a681610beb565b63ffffffff81160361019557565b356107a681611f2e565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160051b3603831361019557565b67ffffffffffffffff81116106325760051b60200190565b929190611fbe81611f9a565b93611fcc604051958661066f565b602085838152019160051b810192831161019557905b828210611fee57505050565b602080918335611ffd81610a11565b815201910190611fe2565b356107a681610a11565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906c0c9f2c9cd04674edea400000008202918083046c0c9f2c9cd04674edea40000000149015171561110d57565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e6400000000149015171561110d57565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81160361110d57565b8181029291811591840414171561110d57565b8181106120ef575050565b600081556001016120e4565b9067ffffffffffffffff83116106325768010000000000000000831161063257815483835580841061215f575b5090600052602060002060005b8381106121425750505050565b600190602084359461215386610a11565b01938184015501612135565b612177908360005284602060002091820191016120e4565b38612128565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff82116101955760200191813603831361019557565b90600182811c92168015612217575b60208310146121e857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916121dd565b9190601f811161223057505050565b6106c0926000526020600020906020601f840160051c8301931061225c575b601f0160051c01906120e4565b909150819061224f565b90929167ffffffffffffffff81116106325761228c8161228684546121ce565b84612221565b6000601f82116001146122ea5781906122db9394956000926122df575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b0135905038806122a9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169461231d84600052602060002090565b91805b87811061237657508360019596971061233e575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c19910135169055388080612334565b90926020600181928686013581550194019101612320565b9160209082815201919060005b8181106123a85750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff87356123d181610a11565b16815201940192910161239b565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b94916124789373ffffffffffffffffffffffffffffffffffffffff958661246a9367ffffffffffffffff6107a69e9c9d9b96168a5216602089015260c0604089015260c088019161238e565b91858303606087015261238e565b9416608082015260a08185039101526123df565b906040519182815491828252602082019060005260206000209260005b8181106124be5750506106c09250038361066f565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016124a9565b9060405191826000825492612501846121ce565b808452936001811690811561256d5750600114612526575b506106c09250038361066f565b90506000929192526020600020906000915b8183106125515750509060206106c09282010138612519565b6020919350806001915483858901015201910190918492612538565b602093506106c09592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138612519565b906040516125ba81610615565b60e06126ab60048395612622612618825473ffffffffffffffffffffffffffffffffffffffff80821616885261260f6125fe8267ffffffffffffffff9060a01c1690565b67ffffffffffffffff1660208a0152565b60e01c60ff1690565b60ff166040870152565b612683612666600183015461264a61263d8263ffffffff1690565b63ffffffff1660608a0152565b60201c73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166080870152565b61268f6002820161248c565b60a08601526126a06003820161248c565b60c0860152016124ed565b910152565b604051906126bf60208361066f565b6000808352366020840137565b906126d682611f9a565b6126e3604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06127118294611f9a565b0190602036910137565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160061b3603831361019557565b9015611f155790565b9081602091031261019557516107a681610a11565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015610be65773ffffffffffffffffffffffffffffffffffffffff9160009161284257501690565b61285b915060203d6020116119c8576119ba818361066f565b1690565b9190811015611f155760051b0190565b90816020910312610195575190565b9160206107a69381815201916123df565b60ff166020039060ff821161110d57565b909291928311610195579190565b906004116101955790600490565b90939293848311610195578411610195578101920390565b3590602081106128e2575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b92919261291b826106e0565b91612929604051938461066f565b829481845281830111610195578281602093846000960137010152565b9160ff811690602082106129a2575b50810361296857906107a691369161290f565b9061299e6040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161287e565b0390fd5b60208311612a3857602083036129555790506129d66129d060ff6129c88496959661288f565b1685856128a0565b906128d4565b612a0257916129fb91816129f56129ef6107a69661288f565b60ff1690565b916128bc565b369161290f565b5061299e6040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161287e565b6040517f3aeba3900000000000000000000000000000000000000000000000000000000081528061299e85876004840161287e565b9081602091031261019557516107a68161111d565b67ffffffffffffffff1667ffffffffffffffff811461110d5760010190565b6040519060c0820182811067ffffffffffffffff82111761063257604052606060a0836000815282602082015282604082015282808201528260808201520152565b90612aed82611f9a565b612afa604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612b288294611f9a565b019060005b828110612b3957505050565b602090612b44612aa1565b82828501015201612b2d565b60405190612b5d82610637565b60606040838281528260208201520152565b919082604091031261019557604051612b8781610653565b60208082948035612b9781610a11565b84520135910152565b805115611f155760200190565b8051821015611f155760209160051b010190565b90612bcb82611f9a565b612bd8604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612c068294611f9a565b019060005b828110612c1757505050565b806060602080938501015201612c0b565b60409067ffffffffffffffff6107a694931681528160208201520190610752565b81601f82011215610195578051612c5f816106e0565b92612c6d604051948561066f565b81845260208284010111610195576107a6916020808501910161072f565b9060208282031261019557815167ffffffffffffffff8111610195576107a69201612c49565b9080602083519182815201916020808360051b8301019401926000915b838310612cdd57505050505090565b9091929394602080612d82837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a0612d71612d5f612d4d612d3b8887015160c08a88015260c0870190610752565b60408701518682036040880152610752565b60608601518582036060870152610752565b60808501518482036080860152610752565b9201519060a0818403910152610752565b97019301930191939290612cce565b9193906107a69593612f8d612fb29260a08652612dbb60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612f5a612f24612eee612eb8612e82612e4e8c61026060e08a0151916101c06101808201520190610752565b6101008801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001888f0152610752565b6101208701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101c08e0152610752565b6101408601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608c8303016101e08d0152610752565b6101608501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8303016102008c0152610752565b6101808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a8303016102208b0152612cb1565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6087830301610240880152610752565b956020850152604084019073ffffffffffffffffffffffffffffffffffffffff169052565b60608201526080818403910152610752565b9080602083519182815201916020808360051b8301019401926000915b838310612ff057505050505090565b909192939460208061302c837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610752565b97019301930191939290612fe1565b93929061305090606086526060860190610752565b938085036020820152825180865260208601906020808260051b8901019501916000905b82821061309257505050506107a69394506040818403910152612fc4565b9091929560208061311e837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001960301865260a060808c5173ffffffffffffffffffffffffffffffffffffffff815116845263ffffffff86820151168685015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610752565b980192019201909291613074565b73ffffffffffffffffffffffffffffffffffffffff60015416330361314d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b906001820180921161110d57565b600101908160011161110d57565b906014820180921161110d57565b90600c820180921161110d57565b9190820180921161110d57565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161110d57565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe820191821161110d57565b9190820391821161110d57565b8051916132318151846131af565b9283156133885760005b848110613249575050505050565b8181101561336d5761325e6117fc8286612bad565b73ffffffffffffffffffffffffffffffffffffffff811680156133435761328483613177565b8781106132965750505060010161323b565b848110156133135773ffffffffffffffffffffffffffffffffffffffff6132c06117fc838a612bad565b1682146132cf57600101613284565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff61333e6117fc6133388885613216565b89612bad565b6132c0565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6133836117fc61337d8484613216565b85612bad565b61325e565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b604051906133bf82610615565b606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613425575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261019557825167ffffffffffffffff8111610195578261347f918501612c49565b92602081015161348e81611f2e565b92604082015167ffffffffffffffff8111610195576107a69201612c49565b60409067ffffffffffffffff6107a6959316815281602082015201916123df565b9060ff6134e8602092959495604085526040850190610752565b9416910152565b929190926134fb6133b2565b6004841015806138bd575b156137835750509061351791615ca7565b9060c0820151805161370f575b506040820180515160005b818110613660575050805151156135c3575b505b6080820173ffffffffffffffffffffffffffffffffffffffff61357a825173ffffffffffffffffffffffffffffffffffffffff1690565b161561358557505090565b6135a960806107a693015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b9060c08194939401926135d78451516126cc565b83526135e4845151612bc1565b946060810195865260005b85518051821015613652579061363361360d6117fc83600195612bad565b613618838951612bad565b9073ffffffffffffffffffffffffffffffffffffffff169052565b61364b81895161364161071a565b6119108383612bad565b50016135ef565b505093509350905038613541565b61366981613177565b828110613679575060010161352f565b6136876117fc838651612bad565b73ffffffffffffffffffffffffffffffffffffffff6136ad6104086117fc858951612bad565b9116146136bc57600101613669565b6105806136cd6117fc848751612bad565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b600061375891613723604085015160ff1690565b9060405193849283927f6d7fa1ce000000000000000000000000000000000000000000000000000000008452600484016134ce565b0381305afa8015610be657156135245761377c903d806000833e611930818361066f565b5038613524565b60c08596929601946137968651516126cc565b94604083019586526137a9875151612bc1565b976060840198895260005b885180518210156137f257906137dd6137d26117fc83600195612bad565b613618838c51612bad565b6137eb818c5161364161071a565b50016137b4565b5050919550919550600092965061385f9361382861040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b91604051958694859384937f9cc19996000000000000000000000000000000000000000000000000000000008552600485016134ad565b03915afa8015610be657600090600090600090613891575b60a086015263ffffffff16845290505b60c0830152613543565b5050506138b3613887913d806000833e6138ab818361066f565b810190613457565b9192508291613877565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061391361390d87876128ae565b906133f1565b1614613506565b6020818303126101955780519067ffffffffffffffff821161019557019080601f8301121561019557815161394e81611f9a565b9261395c604051948561066f565b81845260208085019260051b82010192831161019557602001905b8282106139845750505090565b60208091835161399381610a11565b815201910190613977565b95949060009460a09467ffffffffffffffff6139f29573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c0860190610752565b930152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461110d5760010190565b949294939093613a536003613a4d8367ffffffffffffffff166000526004602052604060002090565b0161248c565b9473ffffffffffffffffffffffffffffffffffffffff613a74818316612799565b16926040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180613ad160048201907fdc0cbd3600000000000000000000000000000000000000000000000000000000602083019252565b0381885afa908115610be657600091613c45575b5015613c3b5790613b2b600095949392604051998a96879586957f89720a620000000000000000000000000000000000000000000000000000000087526004870161399e565b03915afa928315610be657600093613c16575b50825115613c1157613b5b613b5684518451906131af565b6126cc565b6000918293835b8651811015613bc057613b786117fc8289612bad565b73ffffffffffffffffffffffffffffffffffffffff811615613bb45790613bae600192613618613ba7896139f7565b9888612bad565b01613b62565b50945060018095613bae565b509193909450613bd1575b50815290565b60005b8151811015613c095780613c03613bf06117fc60019486612bad565b613618613bfc876139f7565b9688612bad565b01613bd4565b505038613bcb565b915090565b613c349193503d806000833e613c2c818361066f565b81019061391a565b9138613b3e565b5050505050915090565b613c5e915060203d602011611c6b57611c5d818361066f565b38613ae5565b93919293613c80613c7882518651906131af565b8651906131af565b90613c93613c8d836126cc565b92612bc1565b94600096875b8351891015613cf95788613cef613ce2600193613cca613cc06117fc8e9f9d9e9d8b612bad565b613618838c612bad565b613ce8613cd7858c612bad565b5191809384916139f7565b9c612bad565b528b612bad565b5001979695613c99565b959250929350955060005b8651811015613d9357613d1a6117fc8289612bad565b600073ffffffffffffffffffffffffffffffffffffffff8216815b888110613d67575b5050906001929115613d51575b5001613d04565b613d6190613618613ba7896139f7565b38613d4a565b81613d786104086117fc848c612bad565b14613d8557600101613d35565b506001915081905038613d3d565b509390945060005b8551811015613e3157613db16117fc8288612bad565b600073ffffffffffffffffffffffffffffffffffffffff8216815b878110613e05575b5050906001929115613de8575b5001613d9b565b613dff90613618613df8886139f7565b9787612bad565b38613de1565b81613e166104086117fc848b612bad565b14613e2357600101613dcc565b506001915081905038613dd4565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176106325760405260606080836000815260006020820152600060408201526000838201520152565b90613e8682611f9a565b613e93604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613ec18294611f9a565b019060005b828110613ed257505050565b602090613edd613e3d565b82828501015201613ec6565b519061ffff8216820361019557565b9081606091031261019557613f0c81613ee9565b9160406020830151613f1d81611f2e565b9201516107a681611f2e565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561019557016020813591019167ffffffffffffffff821161019557813603831361019557565b9160209082815201919060005b818110613f935750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735613fbc81610a11565b16815260208781013590820152019401929101613f86565b949391929067ffffffffffffffff1685526080602086015261404b61400e613ffc8580613f29565b60a060808a01526101208901916123df565b61401b6020860186613f29565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808984030160a08a01526123df565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215610195578401916020833593019167ffffffffffffffff8411610195578360061b36038313610195576106c09561411c6140e68360609761415b978d60c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808261414d9a0301910152613f79565b916141126140f5888301611112565b73ffffffffffffffffffffffffffffffffffffffff1660e08d0152565b6080810190613f29565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101008c01526123df565b908782036040890152610752565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff821161110d57565b908160a091031261019557805191602082015161419c81611f2e565b9160408101516141ab81611f2e565b9160806141ba60608401613ee9565b9201516107a68161111d565b9260c09473ffffffffffffffffffffffffffffffffffffffff9167ffffffffffffffff61ffff95846107a69b9a9616885216602087015260408601521660608401521660808201528160a08201520190610752565b90816060910312610195578051613f0c81611f2e565b919082608091031261019557815161424881611f2e565b916020810151916060604083015192015190565b8115614266570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b929190926000916000926040820192835151926142d26142cd6142c860408b01966142c0888d61271b565b9190506131af565b613177565b613e7c565b956000946000945b8751805187101561455d576104086104086117fc896142f894612bad565b6143446020606087019261430d8a8551612bad565b519060405180809581947f958021a70000000000000000000000000000000000000000000000000000000083528d60048401612c28565b03915afa8015610be65773ffffffffffffffffffffffffffffffffffffffff9160009161453f575b501680156144de579060608d93926143858a8451612bad565b519061439660208a015161ffff1690565b958a6143d1604051988995869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613fd4565b03915afa8015610be65760019361447b938a8e8e95600080958197614484575b509083929161ffff614419856144126117fc61446f996144759d9e51612bad565b9451612bad565b51916144426144266106c2565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b63ffffffff8916602085015263ffffffff8b16604085015216606083015260808201526119108383612bad565b50614166565b98614166565b950194956142da565b61447597506117fc965084939291509361ffff614419826144126144c161446f9960603d81116144d7575b6144b9818361066f565b810190613ef8565b9c9196909c9d50505050505050909192936143f1565b503d6144af565b610580876144f06117fc8b8e51612bad565b7f83c758a60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff1660045267ffffffffffffffff16602452604490565b614557915060203d81116119c8576119ba818361066f565b3861436c565b509197969099939592945089614573818661271b565b9050614807575b50508551614587906131bc565b9884614596602086018661217d565b926145a291508661271b565b90506145ae9288615f59565b6145b88a88612bad565b526145c38987612bad565b506145ce8987612bad565b516020015163ffffffff166145e291614166565b906145ed8987612bad565b516040015163ffffffff1661460191614166565b906002546146229073ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169261464590606001612008565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff91821660248701529116604485015273ffffffffffffffffffffffffffffffffffffffff166064840152826084815a93608094fa958615610be6576000966000976000946000926147ca575b506000965b86518810156147bd5761472c6001916146f4886146ef87612041565b61425c565b61470d60606147038d8d612bad565b51019182516120d1565b9052858a14614734575b60606147238b8b612bad565b510151906131af565b9701966146d3565b8b8873eba517d20000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff61478160808c015173ffffffffffffffffffffffffffffffffffffffff1690565b160361478f575b5050614717565b6146ef61479b9261206f565b6147b460606147aa8d8d612bad565b51019182516131af565b90528b88614788565b9598509694955050505050565b92985050506147f291925060803d608011614800575b6147ea818361066f565b810190614231565b9197909392909190386146ce565b503d6147e0565b610a7a6104086109a96109a361481d948961271b565b9073ffffffffffffffffffffffffffffffffffffffff6000915151921660e0870180516148486106c2565b73ffffffffffffffffffffffffffffffffffffffff84168152908460208301528460408301528460608301526080820152614883858c612bad565b5261488e848b612bad565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdc0cbd3600000000000000000000000000000000000000000000000000000000600482015291602083602481845afa908115610be6578a8f8a938c968891614b60575b50614a5d575b50505050505015614911575b8961457a565b6149ae60608b61494c6109a96109a361494561040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b938a61271b565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8b16600482015273ffffffffffffffffffffffffffffffffffffffff909116602482015292839190829081906044820190565b03915afa8015610be65763ffffffff9189600092600092600091614a20575b5090614a0c614a17939260406149ff60606149e88b87612bad565b51019960206149f78288612bad565b510195612bad565b51019063ffffffff169052565b9063ffffffff169052565b1690523861490b565b614a0c9450614a179350614a4c915060603d606011614a56575b614a44818361066f565b81019061421b565b91949093506149cd565b503d614a3a565b614ad860a095614a9f6020614a95606082614a8c6109a3614a856109a96109a38f8d9061271b565b998d61271b565b01359901612008565b99015161ffff1690565b905190604051988997889687967f2c063404000000000000000000000000000000000000000000000000000000008852600488016141c6565b03915afa918215610be657809181908294614b24575b50614b18908b614a0c6060614b038984612bad565b51019460406149ff8a60206149f78288612bad565b52853886818a8f6148ff565b915050614b189250614b4e915060a03d60a011614b59575b614b46818361066f565b810190614180565b949192919050614aee565b503d614b3c565b614b79915060203d602011611c6b57611c5d818361066f565b386148f9565b9073ffffffffffffffffffffffffffffffffffffffff614c519392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452614be160648561066f565b16600080604093845195614bf5868861066f565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15614c76573d614c42614c39826106e0565b9451948561066f565b83523d6000602085013e61666f565b805180614c5c575050565b81602080614c71936106c09501019101612a6d565b6161a8565b6060925061666f565b9160606106c0929493614ccc8160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b90614d0b826106e0565b614d18604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061271182946106e0565b91825160148102908082046014149015171561110d57614d68614d6d91613185565b613193565b90614d7f614d7a836131a1565b614d01565b906014614d8b83612ba0565b5360009260215b8651851015614dbd576014600191614dad6117fc888b612bad565b60601b8187015201940193614d92565b919550936020935060601b90820152828152012090565b90614de461040860608401612008565b91614e147fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff91604081019061271b565b9050614e93575b60005b8251811015614e8d57806060614e3660019386612bad565b51015180158015614e84575b614e7e57614e7890614e72614e578488612bad565b515173ffffffffffffffffffffffffffffffffffffffff1690565b87614b7f565b01614e1e565b50614e78565b50838214614e42565b50505050565b50614e9e81516131e9565b614eab614e578284612bad565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdc0cbd3600000000000000000000000000000000000000000000000000000000600482015260208160248173ffffffffffffffffffffffffffffffffffffffff86165afa908115610be657600091614f51575b50614f31575b50614e1b565b614f4b906060614f418486612bad565b5101519085614b7f565b38614f2b565b614f6a915060203d602011611c6b57611c5d818361066f565b38614f25565b60405190614f7d82610653565b60606020838281520152565b91906040838203126101955760405190614fa282610653565b8193805167ffffffffffffffff81116101955782614fc1918301612c49565b835260208101519167ffffffffffffffff8311610195576020926126ab9201612c49565b9060208282031261019557815167ffffffffffffffff8111610195576107a69201614f89565b90608073ffffffffffffffffffffffffffffffffffffffff81615037855160a0865260a0860190610752565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b9060206107a692818152019061500b565b919060408382031261019557825167ffffffffffffffff8111610195576020916150a7918501614f89565b92015190565b61ffff6150c66107a6959360608452606084019061500b565b931660208201526040818403910152610752565b909192936150e6612aa1565b50602082019081511561567a5761511a610408610a7a610408865173ffffffffffffffffffffffffffffffffffffffff1690565b9573ffffffffffffffffffffffffffffffffffffffff871692831580156155ef575b61558c576151c781519161514e614f70565b5051865173ffffffffffffffffffffffffffffffffffffffff16906151a56151746106c2565b8b815267ffffffffffffffff8b1660208201529573ffffffffffffffffffffffffffffffffffffffff166040870152565b606085015273ffffffffffffffffffffffffffffffffffffffff166080840152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdc0cbd36000000000000000000000000000000000000000000000000000000006004820152602081602481885afa908115610be65760009161556d575b501561547757509161527396979160008094604051998a95869485937fb1c71c65000000000000000000000000000000000000000000000000000000008552600485016150ad565b03925af18015610be65760009460009161542e575b509460006153cd9361536b61530761533f956117fc6152db9a965b6040519b8c91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018c528b61066f565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810186528561066f565b61539861538e84519267ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b9060405195869283927f6d7fa1ce000000000000000000000000000000000000000000000000000000008452600484016134ce565b0381305afa928315610be65760009361540e575b5060200151936153ef6106d1565b958652602086015260408501526060840152608083015260a082015290565b6020919350615427903d806000833e611930818361066f565b92906153e1565b6152db95506117fc9691506153cd9361536b61530761533f956154646000953d8088833e61545c818361066f565b81019061507c565b9b909b969b9a5050955050509350615288565b979161ffff91959793501661554357516155195760006154c493604051809581927f9a4575b90000000000000000000000000000000000000000000000000000000083526004830161506b565b038183855af1918215610be6576152db9561536b6153076000936117fc61533f976153cd9987916154f7575b50966152a3565b61551391503d8089833e61550b818361066f565b810190614fe5565b386154f0565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b615586915060203d602011611c6b57611c5d818361066f565b3861522b565b6105806155ad865173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818c5afa908115610be65760009161565b575b501561513c565b615674915060203d602011611c6b57611c5d818361066f565b38615654565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592615776947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b906157926020928281519485920161072f565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b1681526157de825180936020898501910161072f565b019160f81b16838201526157fc82518093602060028501910161072f565b01019160f81b168382015261581b82518093602060028501910161072f565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff000000000000000000000000000000000000000000000000000000000000006107a69f9e9c9860f81b168152615890825180936020898501910161072f565b019160f01b16838201526158ae82518093602060038501910161072f565b01019160f01b16838201526158cc825180936020898501910161072f565b01019160f01b166002820152019061577f565b60e081019060ff82515111615c785761010081019060ff82515111615c495761012081019260ff84515111615c1a5761014082019060ff82515111615beb5761016083019461ffff86515111615bbc57610180840194600186515111615b8d576101a085019261ffff84515111615b5c57855167ffffffffffffffff16602087015167ffffffffffffffff169060408801516159829067ffffffffffffffff1690565b9760608101516159959063ffffffff1690565b9060808101516159a89063ffffffff1690565b60a082015161ffff169160c00151926040519b8c9660208801966159cb976156a4565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810187526159fb908761066f565b51908151615a099060ff1690565b9051805160ff169351908151615a1f9060ff1690565b906040519586956020870195615a3496615796565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252615a64908261066f565b606094519081511515615b0e97615ad26107a698615b1497615b0e96615afe95615b40575b505196615a97885160ff1690565b935191615aa6835161ffff1690565b91615ab3825161ffff1690565b905193615ac2855161ffff1690565b936040519b8c9860208a01615821565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810185528461066f565b604051968795602087019061577f565b9061577f565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261066f565b615b55919250615b4f90612ba0565b5161636e565b9038615a89565b7fb4205b42000000000000000000000000000000000000000000000000000000006000526105806024906024600452565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b90615cb06133b2565b9160118210615ed35780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000821603615e605750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a615d2f816126cc565b60408601908152615d3f82612bc1565b906060870191825260005b838110615e145750505050615dbf8383615db5615da9615d9f615d98615d78615dc99887615dd39c9b616504565b73ffffffffffffffffffffffffffffffffffffffff90911660808d015290565b85856165da565b929190369161290f565b60a08a01528383616642565b949190369161290f565b60c08801526165da565b939190369161290f565b60e08401528103615de2575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b80600191615e59615e43615e3c615e2f615e539a8d8d616504565b9190613618868a51612bad565b8b8b6165da565b9391889a919a51949a369161290f565b92612bad565b5201615d4a565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b90816020910312610195576107a690613ee9565b9261ffff6107a6959367ffffffffffffffff615f4b94168652166020850152608060408501526080840190610c8a565b916060818403910152610752565b9091615f63613e3d565b50615f828267ffffffffffffffff166000526004602052604060002090565b93615f92855460ff9060e01c1690565b9061604d6160326160296080840194616015615fef615fe7615fd96001615fcd8b5173ffffffffffffffffffffffffffffffffffffffff1690565b9e015463ffffffff1690565b885163ffffffff1690614166565b9a60756131af565b9761602361601b60ff61600960a08b019c8d5151906131af565b951694616015866120a1565b906131af565b93604f6131af565b906120d1565b63ffffffff1690565b925173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811673eba517d200000000000000000000000000000000036160e157505061ffff92506160d3906160c66000935b51956160b961609d6106c2565b73ffffffffffffffffffffffffffffffffffffffff9099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b9061610461040860209373ffffffffffffffffffffffffffffffffffffffff1690565b60406161148484015161ffff1690565b92015191855196616154604051988995869485947fe962e69e00000000000000000000000000000000000000000000000000000000865260048601615f1b565b03915afa908115610be6576160c66160d39261ffff95600091616179575b5093616090565b61619b915060203d6020116161a1575b616193818361066f565b810190615f07565b38616172565b503d616189565b156161af57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b9660026163429761630f60226107a69f9e9c9799600199859f9b7fff000000000000000000000000000000000000000000000000000000000000009061630f9f8261630f9c6163169c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b1660218201526162c0825180936020898501910161072f565b019160f81b16838201526162de82518093602060238501910161072f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b019061577f565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff825151116164d557604081019160ff835151116164a657606082019160ff8351511161647757608081019260ff845151116164485760a0820161ffff81515111616419576107a694615b1493519451916163d0835160ff1690565b9751916163de835160ff1690565b9451906163ec825160ff1690565b9051936163fa855160ff1690565b935196616409885161ffff1690565b966040519c8d9b60208d01616233565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602960045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602560045260246000fd5b9291909260018201918483116165a85781013560001a82811561659d57506014810361657057820193841161653c57013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b919060028201918183116165a8578381013560f01c016002019281841161660e57918391616607936128bc565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b919060018201918183116165a8578381013560001a016001019281841161660e57918391616607936128bc565b919290156166ea5750815115616683575090565b3b1561668c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156166fd5750805190602001fd5b61299e906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161079556fea164736f6c634300081a000a",
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
