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
	DestGasLimit      uint64
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateDestChainAddress\",\"inputs\":[{\"name\":\"rawAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"validatedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102f557604051615b3938819003601f8101601f191683016001600160401b038111848210176102fa578392829160405283398101039060c082126102f557606082126102f557610054610310565b81519092906001600160401b03811681036102f55783526020820151906001600160a01b03821682036102f5576020840191825260606100966040850161032f565b6040860190815291605f1901126102f5576100af610310565b916100bc6060850161032f565b835260808401519384151585036102f55760a06100e091602086019687520161032f565b946040840195865233156102e457600180546001600160a01b0319163317905580516001600160401b03161580156102d2575b80156102c0575b61029357516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102ae575b80156102a4575b6102935780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610310565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610310565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a16040516157f5908161034482396080518181816105bf015281816116630152611c29015260a05181611c62015260c051818181611c9e01526125df0152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b038111838210176102fa57604052565b51906001600160a01b03821682036102f55756fe6080604052600436101561001257600080fd5b60003560e01c806306285c691461011757806314a8cfa314610112578063181f5a771461010d57806320487ded1461010857806348a98aa4146101035780635cb80c5d146100fe5780636d7fa1ce146100f95780636def4ce7146100f45780637437ff9f146100ef57806379ba5097146100ea5780638da5cb5b146100e55780639041be3d146100e057806390423fa2146100db578063df0aa9e9146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611b86565b611a92565b6113a2565b6111ce565b61112c565b6110da565b610ff1565b610f25565b610e7f565b610c9d565b610b44565b610ad6565b610868565b6107c7565b610217565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576060610150611c09565b610193604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101955760043567ffffffffffffffff81116101955760040160009280601f830112156102135781359367ffffffffffffffff851161021057506020808301928560051b010111610195579190565b80fd5b8380fd5b34610195576102253661019a565b61022d612f10565b60005b81811061023957005b610244818385611cf5565b9061024e82611d3a565b67ffffffffffffffff811690811580156105b3575b801561059d575b8015610584575b610549576102ba906102d4608086019161028b8388611d66565b94906102b460a08a01966102ac6102a2898d611d66565b9490923691611dd2565b923691611dd2565b90612fc2565b67ffffffffffffffff166000526004602052604060002090565b9160208601906103276102e683611e28565b859073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b61038461033660408901611d44565b85547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178555565b61039060608801611d5c565b6103c96001860191829063ffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000825416179055565b6103e06103d6858a611d66565b9060038801611ebb565b6103f76103ed838a611d66565b9060028801611ebb565b60c088019161042161040884611e28565b73ffffffffffffffffffffffffffffffffffffffff1690565b1561051f57600198610501846105077f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae996104ad6104f9976104656105169a611e28565b7fffffffffffffffff0000000000000000000000000000000000000000ffffffff77ffffffffffffffffffffffffffffffffffffffff0000000083549260201b169116179055565b6104f06104e96104e360e08801936104d26104c8868b611f3d565b9060048401612026565b5460a01c67ffffffffffffffff1690565b9a611e28565b9a86611d66565b97909686611d66565b949093611e28565b94611f3d565b959094604051998a998a6121de565b0390a201610230565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b5063ffffffff61059660608601611d5c565b1615610271565b5060ff6105ac60408601611d44565b161561026a565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610263565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610100810190811067ffffffffffffffff82111761063257604052565b6105e6565b6060810190811067ffffffffffffffff82111761063257604052565b6040810190811067ffffffffffffffff82111761063257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761063257604052565b604051906106c06101808361066f565b565b604051906106c060e08361066f565b604051906106c060408361066f565b604051906106c060a08361066f565b604051906106c060c08361066f565b67ffffffffffffffff811161063257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6040519061074760208361066f565b60008252565b60005b8381106107605750506000910152565b8181015183820152602001610750565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936107ac8151809281875287808801910161074d565b0116010190565b9060206107c4928181520190610770565b90565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576108446040805190610808818361066f565b601082527f4f6e52616d7020312e372e302d64657600000000000000000000000000000000602083015251918291602083526020830190610770565b0390f35b67ffffffffffffffff81160361019557565b908160a09103126101955790565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576004356108a381610848565b60243567ffffffffffffffff8111610195576108c390369060040161085a565b6000906108e48367ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff61091b610408845473ffffffffffffffffffffffffffffffffffffffff1690565b1615610a8057610984929361097d6109496109396080850185611f3d565b906109438761236d565b85613521565b93610952612470565b906040850161096181876124db565b9050610a2b575b5061097760028751920161224c565b90613c37565b8352614173565b6000916000916000915b8151831015610a2057610a166109ee6109b860019360606109af8888612545565b51015190612567565b966109e86109db60206109cb8989612545565b51015167ffffffffffffffff1690565b67ffffffffffffffff1690565b90612567565b946109e8610a0d6040610a018888612545565b51015163ffffffff1690565b63ffffffff1690565b920191929361098e565b604051908152602090f35b610a799192506020610a5b610a49610a54610a4f610a49868c6124db565b9061252f565b611e28565b93896124db565b0135610a6c602089015161ffff1690565b9060c089015192876139cc565b9038610968565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610b10600435610848565b6020610b26602435610b2181610ab8565b612580565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b3461019557610b523661019a565b90610b7260035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b818110610b7f57005b610b90610408610a4f838587612679565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa908115610c8d576001948891600093610c5d575b5082610c05575b5050505001610b76565b610c1091839161441b565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610bfb565b610c7f91935060203d8111610c86575b610c77818361066f565b810190612689565b9138610bf4565b503d610c6d565b612574565b60ff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760043567ffffffffffffffff8111610195573660238201121561019557806004013567ffffffffffffffff81116101955736602482840101116101955761084491610d25916024803592610d1f84610c92565b0161279b565b604051918291826107b3565b906020808351928381520192019060005b818110610d4f5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610d42565b906107c49160208152610da760208201835173ffffffffffffffffffffffffffffffffffffffff169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015163ffffffff166080820152608082015173ffffffffffffffffffffffffffffffffffffffff1660a082015260e0610e4b610e1860a085015161010060c0860152610120850190610d31565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610d31565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610770565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff600435610ec381610848565b606060e0604051610ed381610615565b600081526000602082015260006040820152600083820152600060808201528260a08201528260c08201520152166000526004602052610844610f19604060002061236d565b60405191829182610d7b565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610f5c611bea565b50604051610f6981610637565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff600354166040820152604051809161084482606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760005473ffffffffffffffffffffffffffffffffffffffff811633036110b0577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff60043561117081610848565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff81116111b45760405167ffffffffffffffff9091168152602090f35b611e32565b35906106c082610ab8565b8015150361019557565b346101955760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557600060405161120b81610637565b60043561121781610ab8565b8152602435611225816111c4565b602082019081526044359061123982610ab8565b60408301918252611248612f10565b73ffffffffffffffffffffffffffffffffffffffff83511615918215611382575b508115611377575b5061134f5780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39061133a611c09565b6113496040519283928361451b565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538611271565b5173ffffffffffffffffffffffffffffffffffffffff1615915038611269565b346101955760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576113dc600435610848565b60243567ffffffffffffffff8111610195576113fc90369060040161085a565b60443561140a606435610ab8565b60025460a01c60ff16611a685761145b740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61147b60043567ffffffffffffffff166000526004602052604060002090565b9173ffffffffffffffffffffffffffffffffffffffff6064351615611a3e578254926114bd61040873ffffffffffffffffffffffffffffffffffffffff861681565b3303611a14576114d06080830183611f3d565b6114d98361236d565b6114e69290600435613521565b9360a01c67ffffffffffffffff166114fd906128c2565b81547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff0000000000000000000000000000000000000000161782556040513060601b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526014815290949061158760348261066f565b602082015161ffff16604083015163ffffffff1660405160643560601b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166020820152909490928360348101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018552611604908561066f565b61160e8780611f3d565b835460e01c60ff16906116209261279b565b9260808601516040890195611635878b6124db565b61163f9150612923565b9861164d60208c018c611f3d565b9490956116586106b0565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529e8f9067ffffffffffffffff60043516602083015267ffffffffffffffff166040820152606001528d6116ba600489016122ad565b608082015260a001906116d0919061ffff169052565b63ffffffff1660c08d015260e08c01526101008b019485526101208b01526101408a01968752369061170192612764565b61016089015261170f612470565b61171984886124db565b6117329391506119ca575b61097760028751920161224c565b835261173c612990565b9361174a8487600435614173565b926020860193845261175c81886124db565b905061192d575b5050509261177086614d97565b808452602081519101206117858351516129e0565b9360408101948552606060009301925b8451805182101561188057906000868686838e61182a8f986117f1896117ea6117e46104086104086117c985602098612545565b515173ffffffffffffffffffffffffffffffffffffffff1690565b98611e28565b9851612545565b510151604051998a97889687957f97048c7100000000000000000000000000000000000000000000000000000000875260048701612b8f565b03925af18015610c8d57816118589160019460009161185f575b508951906118528383612545565b52612545565b5001611795565b61187a913d8091833e611872818361066f565b810190612a89565b38611844565b6108448480898b7f9b8fdf7fa94e7e8692c830c07cc6ce91a34c507d9f8efea07eb71cd64ed4891f8f896118ed6118cb604067ffffffffffffffff94015167ffffffffffffffff1690565b915194519551604051938493169667ffffffffffffffff600435169684612e1b565b0390a461191d7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b600161193982896124db565b9050036119a05761198f92611954610a4961197d938a6124db565b60a08701518051919290911561199857505b606435916119786004359136906129af565b61469c565b90519061198982612538565b52612538565b50388080611763565b905051611966565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b9050611a0e6119df610a4f610a49878b6124db565b60206119ee610a49888c6124db565b01356119ff602089015161ffff1690565b9060c0890151926004356139cc565b90611724565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955773ffffffffffffffffffffffffffffffffffffffff600435611ae281610ab8565b611aea612f10565b16338114611b5c57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557611bc0600435610848565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611bf782610637565b60006040838281528260208201520152565b611c11611bea565b50604051611c1e81610637565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611d355760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181360301821215610195570190565b611cc6565b356107c481610848565b356107c481610c92565b63ffffffff81160361019557565b356107c481611d4e565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160051b3603831361019557565b67ffffffffffffffff81116106325760051b60200190565b929190611dde81611dba565b93611dec604051958661066f565b602085838152019160051b810192831161019557905b828210611e0e57505050565b602080918335611e1d81610ab8565b815201910190611e02565b356107c481610ab8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8116036111b457565b818102929181159184041417156111b457565b818110611eaf575050565b60008155600101611ea4565b9067ffffffffffffffff831161063257680100000000000000008311610632578154838355808410611f1f575b5090600052602060002060005b838110611f025750505050565b6001906020843594611f1386610ab8565b01938184015501611ef5565b611f3790836000528460206000209182019101611ea4565b38611ee8565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff82116101955760200191813603831361019557565b90600182811c92168015611fd7575b6020831014611fa857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611f9d565b9190601f8111611ff057505050565b6106c0926000526020600020906020601f840160051c8301931061201c575b601f0160051c0190611ea4565b909150819061200f565b90929167ffffffffffffffff81116106325761204c816120468454611f8e565b84611fe1565b6000601f82116001146120aa57819061209b93949560009261209f575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b013590503880612069565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216946120dd84600052602060002090565b91805b8781106121365750836001959697106120fe575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c199101351690553880806120f4565b909260206001819286860135815501940191016120e0565b9160209082815201919060005b8181106121685750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561219181610ab8565b16815201940192910161215b565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b94916122389373ffffffffffffffffffffffffffffffffffffffff958661222a9367ffffffffffffffff6107c49e9c9d9b96168a5216602089015260c0604089015260c088019161214e565b91858303606087015261214e565b9416608082015260a081850391015261219f565b906040519182815491828252602082019060005260206000209260005b81811061227e5750506106c09250038361066f565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201612269565b90604051918260008254926122c184611f8e565b808452936001811690811561232d57506001146122e6575b506106c09250038361066f565b90506000929192526020600020906000915b8183106123115750509060206106c092820101386122d9565b60209193508060019154838589010152019101909184926122f8565b602093506106c09592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386122d9565b9060405161237a81610615565b60e061246b600483956123e26123d8825473ffffffffffffffffffffffffffffffffffffffff8082161688526123cf6123be8267ffffffffffffffff9060a01c1690565b67ffffffffffffffff1660208a0152565b60e01c60ff1690565b60ff166040870152565b612443612426600183015461240a6123fd8263ffffffff1690565b63ffffffff1660608a0152565b60201c73ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166080870152565b61244f6002820161224c565b60a08601526124606003820161224c565b60c0860152016122ad565b910152565b6040519061247f60208361066f565b6000808352366020840137565b9061249682611dba565b6124a3604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06124d18294611dba565b0190602036910137565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160061b3603831361019557565b9015611d355790565b805115611d355760200190565b8051821015611d355760209160051b010190565b90600182018092116111b457565b919082018092116111b457565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015610c8d5760009061262a575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011612671575b816126446020938361066f565b810103126101955773ffffffffffffffffffffffffffffffffffffffff905161266c81610ab8565b61260f565b3d9150612637565b9190811015611d355760051b0190565b90816020910312610195575190565b9160206107c493818152019161219f565b60ff166020039060ff82116111b457565b909291928311610195579190565b906004116101955790600490565b909291928360041161019557831161019557600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b90939293848311610195578411610195578101920390565b359060208110612737575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b929192612770826106fe565b9161277e604051938461066f565b829481845281830111610195578281602093846000960137010152565b9160ff811690602082106127f7575b5081036127bd57906107c4913691612764565b906127f36040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612698565b0390fd5b6020831161288d57602083036127aa57905061282b61282560ff61281d849695966126a9565b1685856126ba565b90612729565b6128575791612850918161284a6128446107c4966126a9565b60ff1690565b91612711565b3691612764565b506127f36040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612698565b6040517f3aeba390000000000000000000000000000000000000000000000000000000008152806127f3858760048401612698565b67ffffffffffffffff1667ffffffffffffffff81146111b45760010190565b6040519060c0820182811067ffffffffffffffff82111761063257604052606060a0836000815282602082015282604082015282808201528260808201520152565b9061292d82611dba565b61293a604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06129688294611dba565b019060005b82811061297957505050565b6020906129846128e1565b8282850101520161296d565b6040519061299d82610637565b60606040838281528260208201520152565b9190826040910312610195576040516129c781610653565b602080829480356129d781610ab8565b84520135910152565b906129ea82611dba565b6129f7604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612a258294611dba565b019060005b828110612a3657505050565b806060602080938501015201612a2a565b81601f82011215610195578051612a5d816106fe565b92612a6b604051948561066f565b81845260208284010111610195576107c4916020808501910161074d565b9060208282031261019557815167ffffffffffffffff8111610195576107c49201612a47565b9080602083519182815201916020808360051b8301019401926000915b838310612adb57505050505090565b9091929394602080612b80837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a0612b6f612b5d612b4b612b398887015160c08a88015260c0870190610770565b60408701518682036040880152610770565b60608601518582036060870152610770565b60808501518482036080860152610770565b9201519060a0818403910152610770565b97019301930191939290612acc565b9193906107c49593612d6d612d929260a08652612bb960a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152610160612d3a612d04612cce612c98612c45612c108c61022060608a0151916101806101008201520190610770565b60808801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101208f0152610770565b60a087015161ffff166101408d015260c087015163ffffffff16868d015260e08701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101808e0152610770565b6101008601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608c8303016101a08d0152610770565b6101208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8303016101c08c0152610770565b6101408401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a8303016101e08b0152612aaf565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6087830301610200880152610770565b956020850152604084019073ffffffffffffffffffffffffffffffffffffffff169052565b60608201526080818403910152610770565b9080602083519182815201916020808360051b8301019401926000915b838310612dd057505050505090565b9091929394602080612e0c837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610770565b97019301930191939290612dc1565b939290612e3090606086526060860190610770565b938085036020820152825180865260208601906020808260051b8901019501916000905b828210612e7257505050506107c49394506040818403910152612da4565b90919295602080612f02837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001960301865260a060808c5173ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff86820151168685015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610770565b980192019201909291612e54565b73ffffffffffffffffffffffffffffffffffffffff600154163303612f3157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116111b457565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe82019182116111b457565b919082039182116111b457565b805191612fd0815184612567565b9283156131415760005b848110612fe8575050505050565b8181101561312657613017612ffd8286612545565b5173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811680156130fc5761303d83612559565b87811061304f57505050600101612fda565b848110156130cc5773ffffffffffffffffffffffffffffffffffffffff613079612ffd838a612545565b1682146130885760010161303d565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6130f7612ffd6130f18885612fb5565b89612545565b613079565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b61313c612ffd6131368484612fb5565b85612545565b613017565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519060e0820182811067ffffffffffffffff82111761063257604052606060c08382815260006020820152600060408201526000838201528260808201528260a08201520152565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106131e9575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9061322582611dba565b613232604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06132608294611dba565b019060005b82811061327157505050565b60209060405161328081610653565b6000815260608382015282828501015201613265565b9160608383031261019557825167ffffffffffffffff811161019557826132be918501612a47565b9260208101516132cd81611d4e565b92604082015167ffffffffffffffff8111610195576107c49201612a47565b60409067ffffffffffffffff6107c49593168152816020820152019161219f565b9080601f83011215610195578160206107c493359101612764565b9080601f830112156101955781359161334083611dba565b9261334e604051948561066f565b80845260208085019160051b830101918383116101955760208101915b83831061337a57505050505090565b823567ffffffffffffffff81116101955782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838803011261019557604051906133c782610653565b60208301356133d581610ab8565b825260408301359167ffffffffffffffff8311610195576133fe8860208096958196010161330d565b8382015281520192019161336b565b61ffff81160361019557565b35906106c08261340d565b35906106c082611d4e565b6020818303126101955780359067ffffffffffffffff8211610195570160e0818303126101955761345e6106c2565b91813567ffffffffffffffff8111610195578161347c918401613328565b835261348a60208301613419565b602084015261349b60408301613424565b60408401526134ac606083016111b9565b6060840152608082013567ffffffffffffffff811161019557816134d191840161330d565b608084015260a082013567ffffffffffffffff811161019557816134f691840161330d565b60a084015260c082013567ffffffffffffffff811161019557613519920161330d565b60c082015290565b9291909261352d61316b565b600484101580613850575b15613722575050816135559261354d926126d6565b81019061342f565b9081515160005b818110613673575050815151156135ee575b6060820173ffffffffffffffffffffffffffffffffffffffff6135a5825173ffffffffffffffffffffffffffffffffffffffff1690565b16156135b057505090565b6135d460806107c493015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b60c08101926135fe84515161321b565b835260005b8451805182101561366a579061366381613622612ffd82600196612545565b61364961362d6106d1565b73ffffffffffffffffffffffffffffffffffffffff9092168252565b613651610738565b60208201528751906118528383612545565b5001613603565b5050925061356e565b61367c81612559565b82811061368c575060010161355c565b61369a6117c9838751612545565b73ffffffffffffffffffffffffffffffffffffffff6136c06104086117c9858a51612545565b9116146136cf5760010161367c565b6105806136e06117c9848851612545565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b939094919260c082019561373787515161321b565b865260005b8751805182101561378757906137808161375b612ffd82600196612545565b61376661362d6106d1565b61376e610738565b60208201528a51906118528383612545565b500161373c565b50506137ef939650600092946137b861040861040860025473ffffffffffffffffffffffffffffffffffffffff1690565b91604051958694859384937f9cc19996000000000000000000000000000000000000000000000000000000008552600485016132ec565b03915afa8015610c8d57600090600090600090613824575b608086015263ffffffff16604085015290505b60a083015261356e565b50505061384661381a913d806000833e61383e818361066f565b810190613296565b9192508291613807565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006138a66138a087876126c8565b906131b5565b1614613538565b9081602091031261019557516107c4816111c4565b6020818303126101955780519067ffffffffffffffff821161019557019080601f830112156101955781516138f681611dba565b92613904604051948561066f565b81845260208085019260051b82010192831161019557602001905b82821061392c5750505090565b60208091835161393b81610ab8565b81520191019061391f565b95949060009460a09467ffffffffffffffff61399a9573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c0860190610770565b930152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146111b45760010190565b9492949390936139fb60036139f58367ffffffffffffffff166000526004602052604060002090565b0161224c565b9473ffffffffffffffffffffffffffffffffffffffff613a1c818316612580565b16926040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180613a7960048201907fdc0cbd3600000000000000000000000000000000000000000000000000000000602083019252565b0381885afa908115610c8d57600091613c08575b5015613bfe5790613ad3600095949392604051998a96879586957f89720a6200000000000000000000000000000000000000000000000000000000875260048701613946565b03915afa928315610c8d57600093613bd9575b50825115613bd457613b03613afe8451845190612567565b61248c565b6000918293835b8651811015613b8357613b20612ffd8289612545565b73ffffffffffffffffffffffffffffffffffffffff811615613b775790613b71600192613b56613b4f8961399f565b9888612545565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01613b0a565b50945060018095613b71565b509193909450613b94575b50815290565b60005b8151811015613bcc5780613bc6613bb3612ffd60019486612545565b613b56613bbf8761399f565b9688612545565b01613b97565b505038613b8e565b915090565b613bf79193503d806000833e613bef818361066f565b8101906138c2565b9138613ae6565b5050505050915090565b613c2a915060203d602011613c30575b613c22818361066f565b8101906138ad565b38613a8d565b503d613c18565b613c57613c52613c4a8351855190612567565b855190612567565b61321b565b91600094855b8351811015613c9a5780613c93613c7660019387612545565b5198613c818161399f565b99613c8c828a612545565b5287612545565b5001613c5d565b50939094915060005b8551811015613d7557613cb9612ffd8288612545565b600073ffffffffffffffffffffffffffffffffffffffff8216815b868110613d4b575b505015613ced575b50600101613ca3565b9290613d44600192613d1c613d006106d1565b73ffffffffffffffffffffffffffffffffffffffff9097168752565b613d24610738565b6020870152613d328161399f565b95613d3d8289612545565b5286612545565b5090613ce4565b81613d5c6104086117c9848c612545565b14613d6957600101613cd4565b50505060013880613cdc565b5093506000905b8451821015613e5057613d92612ffd8387612545565b91600073ffffffffffffffffffffffffffffffffffffffff8416815b848110613e26575b505015613dc8575b6001019150613d7c565b613e1d600192613df5613dd96106d1565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b613dfd610738565b6020860152613e0b8161399f565b94613e168288612545565b5285612545565b50829150613dbe565b81613e376104086117c9848b612545565b14613e4457600101613dae565b50505060013880613db6565b8252509150565b6040519060a0820182811067ffffffffffffffff8211176106325760405260606080836000815260006020820152600060408201526000838201520152565b90613ea082611dba565b613ead604051918261066f565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613edb8294611dba565b019060005b828110613eec57505050565b602090613ef7613e57565b82828501015201613ee0565b90816060910312610195578051613f198161340d565b9160406020830151613f2a81611d4e565b9201516107c481611d4e565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561019557016020813591019167ffffffffffffffff821161019557813603831361019557565b9160209082815201919060005b818110613fa05750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735613fc981610ab8565b16815260208781013590820152019401929101613f93565b949391929067ffffffffffffffff1685526080602086015261405861401b6140098580613f36565b60a060808a015261012089019161219f565b6140286020860186613f36565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808984030160a08a015261219f565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215610195578401916020833593019167ffffffffffffffff8411610195578360061b36038313610195576106c0956141296140f383606097614168978d60c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808261415a9a0301910152613f86565b9161411f6141028883016111b9565b73ffffffffffffffffffffffffffffffffffffffff1660e08d0152565b6080810190613f36565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101008c015261219f565b908782036040890152610770565b94019061ffff169052565b9291908151516141a361419e614199604085019361419185876124db565b919050612567565b612559565b613e96565b9060005b8451805182101561433f57816141bc91612545565b51906141e2610408610408845173ffffffffffffffffffffffffffffffffffffffff1690565b6020830190606082516141fa60208b015161ffff1690565b92898d614236604051968795869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613fe1565b03915afa938415610c8d57600194600091829383926142ef575b5061ffff92916142bb63ffffffff61427f6142c8945173ffffffffffffffffffffffffffffffffffffffff1690565b9751966142a961428d6106e0565b73ffffffffffffffffffffffffffffffffffffffff909a168a52565b1667ffffffffffffffff166020880152565b63ffffffff166040860152565b16606083015260808201526142dd8286612545565b526142e88185612545565b50016141a7565b63ffffffff945061ffff935061427f9250906142bb6143276142c89360603d8111614338575b61431f818361066f565b810190613f03565b979195909795965050505090614250565b503d614315565b5050909394614370846143859295949561435c6020880188611f3d565b905061436886896124db565b9290506151bf565b61437a8651612f5b565b90613e168288612545565b5061439081846124db565b905061439d575b50505090565b6143b2610a4f610a4960c093614412966124db565b9101516143dc6143c06106e0565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60006020830152600060408301526000606083015260808201526144008351612f88565b9061440b8285612545565b5282612545565b50388080614397565b9073ffffffffffffffffffffffffffffffffffffffff6144ed9392604051938260208601947fa9059cbb00000000000000000000000000000000000000000000000000000000865216602486015260448501526044845261447d60648561066f565b16600080604093845195614491868861066f565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15614512573d6144de6144d5826106fe565b9451948561066f565b83523d6000602085013e615724565b8051806144f8575050565b8160208061450d936106c095010191016138ad565b6153cd565b60609250615724565b9160606106c09294936145688160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b6020818303126101955780519067ffffffffffffffff8211610195570160408183031261019557604051916145d183610653565b815167ffffffffffffffff811161019557816145ee918401612a47565b8352602082015167ffffffffffffffff81116101955761460e9201612a47565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080614648855184602087015260c0860190610770565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b9060ff614695602092959495604085526040850190610770565b9416910152565b9190916146a76128e1565b5060208101805115614a97576146da610408610b21610408855173ffffffffffffffffffffffffffffffffffffffff1690565b9473ffffffffffffffffffffffffffffffffffffffff8616158015614a0c575b6149a9579060006147ba926147878351614728875173ffffffffffffffffffffffffffffffffffffffff1690565b906147656147346106e0565b8a815267ffffffffffffffff8c1660208201529473ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809481927f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614616565b0381838a5af18015610c8d57614823956149159461488794600093614976575b5061484f6148b391612ffd60009596519b6040519b8c91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018c528b61066f565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810186528561066f565b6148e06148d684519267ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b9060405195869283927f6d7fa1ce0000000000000000000000000000000000000000000000000000000084526004840161467b565b0381305afa928315610c8d57600093614956575b5060200151936149376106ef565b958652602086015260408501526060840152608083015260a082015290565b602091935061496f903d806000833e611872818361066f565b9290614929565b600093506148b391612ffd61499f61484f933d8089833e614997818361066f565b81019061459d565b95505091506147da565b6105806149ca845173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818a5afa908115610c8d57600091614a78575b50156146fa565b614a91915060203d602011613c3057613c22818361066f565b38614a71565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b90614ad46020928281519485920161074d565b0190565b6020614bf296614b92601a614bbf947fff0000000000000000000000000000000000000000000000000000000000000060069f9c979a7fffffffffffffffff000000000000000000000000000000000000000000000000614ad49f9b8160019c614bc69f82907f0100000000000000000000000000000000000000000000000000000000000000895260c01b168e88015260c01b16600986015260c01b16601184015260f81b16601982015201918281519485920161074d565b0180927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190614ac1565b80947fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60e01b7fffffffff00000000000000000000000000000000000000000000000000000000166002830152565b614bbf96614d08967fffff000000000000000000000000000000000000000000000000000000000000600160029c986107c49f9e9c97987fff0000000000000000000000000000000000000000000000000000000000000060049a849882614bbf9b60f81b168152614c99825180936020898501910161074d565b019160f81b1683820152614cb88251600293602082958501910161074d565b01019160f01b1683820152614cd782518093602060038501910161074d565b010191888301907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b6106c09092919260206040519482614d55879451809285808801910161074d565b8301614d698251809385808501910161074d565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810184528361066f565b606081019060ff8251511161509357608081019060ff825151116150645760e081019160ff835151116150355761010082019060ff825151116150065761012083019361ffff85515111614fd757610140840193600185515111614fa85761016081019261ffff84515111614f7957606095518051614f5d575b50815167ffffffffffffffff16906020830151614e359067ffffffffffffffff1690565b926040810151614e4c9067ffffffffffffffff1690565b9951805160ff169251908151614e629060ff1690565b9060a0840151614e739061ffff1690565b60c09094015163ffffffff16946040519d8e9860208a0198614e9499614ad8565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752614ec4908761066f565b51928351614ed29060ff1690565b9251908151614ee19060ff1690565b9551928351614ef19061ffff1690565b938251614eff9061ffff1690565b9151948551614f0f9061ffff1690565b94604051998a9960208b0199614f249a614c1e565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252614f54908261066f565b6107c491614d34565b614f72919650614f6c90612538565b51615560565b9438614e11565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b9063ffffffff8091169116019063ffffffff82116111b457565b9081602091031261019557516107c48161340d565b9094939161ffff9067ffffffffffffffff608084019716835216602082015260806040820152825180955260a0810194602060a08260051b84010194019060005b81811061514f575050506107c49394506060818403910152610770565b9091946020806151b0837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608860019603018c526040838b5173ffffffffffffffffffffffffffffffffffffffff815116845201519181858201520190610770565b97019801910196919096615132565b9290916151ca613e57565b506151e98467ffffffffffffffff166000526004602052604060002090565b936151f9855460ff9060e01c1690565b926152a861528d610a0d60608401966109e861525961525161524060016152348d5173ffffffffffffffffffffffffffffffffffffffff1690565b9e015463ffffffff1690565b604089015163ffffffff16906150c2565b9a6051612567565b9761528761527f60ff61527360808b019c8d515190612567565b9516946109e886611e61565b93604f612567565b90611e91565b945173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff811673eba517d2000000000000000000000000000000000361530a575050506152fc61ffff926142bb63ffffffff6000945b51966142a961428d6106e0565b166060830152608082015290565b9061532d61040860209373ffffffffffffffffffffffffffffffffffffffff1690565b8183015161ffff16915191855194615374604051968795869485947f7cb4e8f2000000000000000000000000000000000000000000000000000000008652600486016150f1565b03915afa928315610c8d576142bb63ffffffff61ffff956152fc9460009161539e575b50946152ef565b6153c0915060203d6020116153c6575b6153b8818361066f565b8101906150dc565b38615397565b503d6153ae565b156153d457565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b966002614d0897614bbf60226107c49f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090614bbf9f82614bbf9c6155349c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b1660218201526154e5825180936020898501910161074d565b019160f81b168382015261550382518093602060238501910161074d565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b602081019060ff825151116156f357604081019160ff835151116156c457606082019160ff8351511161569557608081019260ff845151116156665760a0820161ffff81515111615637576107c49461560b93519451916155c2835160ff1690565b9751916155d0835160ff1690565b9451906155de825160ff1690565b9051936155ec855160ff1690565b9351966155fb885161ffff1690565b966040519c8d9b60208d01615458565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261066f565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602560045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000006000526105806024906024600452565b9192901561579f5750815115615738575090565b3b156157415790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156157b25750805190602001fd5b6127f3906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016107b356fea164736f6c634300081a000a",
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
	return common.HexToHash("0x9b8fdf7fa94e7e8692c830c07cc6ce91a34c507d9f8efea07eb71cd64ed4891f")
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
