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
	Router                    common.Address
	MessageNumber             uint64
	AddressBytesLength        uint8
	TokenReceiverAllowed      bool
	MessageNetworkFeeUSDCents uint16
	TokenNetworkFeeUSDCents   uint16
	BaseExecutionGasCost      uint32
	DefaultExecutor           common.Address
	LaneMandatedCCVs          []common.Address
	DefaultCCVs               []common.Address
	OffRamp                   []byte
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector         uint64
	Router                    common.Address
	AddressBytesLength        uint8
	TokenReceiverAllowed      bool
	MessageNetworkFeeUSDCents uint16
	TokenNetworkFeeUSDCents   uint16
	BaseExecutionGasCost      uint32
	DefaultCCVs               []common.Address
	LaneMandatedCCVs          []common.Address
	DefaultExecutor           common.Address
	OffRamp                   []byte
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
	ChainSelector         uint64
	RmnRemote             common.Address
	MaxUSDCentsPerMessage uint32
	TokenAdminRegistry    common.Address
}

var OnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllDestChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAmountBeforeTokenPoolFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100604052346103675760405161607138819003601f8101601f191683016001600160401b0381118482101761036c578392829160405283398101039060e08212610367576080821261036757610055610382565b81519092906001600160401b03811681036103675783526020820151926001600160a01b03841684036103675760208101938452604083015163ffffffff81168103610367576040820190815260606100af8186016103a1565b83820190815293607f1901126103675760405192606084016001600160401b0381118582101761036c576040526100e8608086016103a1565b845260a08501519485151586036103675760c061010c9160208701978852016103a1565b9560408501968752331561035657600180546001600160a01b0319163317905583516001600160401b0316158015610344575b8015610332575b8015610323575b6103085792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e052815116158015610319575b6103085780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e09260006060610206610382565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff92831691166060610251610382565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a1604051615cbb90816103b68239608051818181610a34015281816114c50152611d33015260a0518181816112a20152611d5f015260c051818181611dba015261275b015260e051818181611d8b0152613eaa0152f35b6306b7c75960e31b60005260046000fd5b508151151561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761036c57604052565b51906001600160a01b03821682036103675756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610117578063181f5a771461011257806320487ded1461010d5780632490769e1461010857806348a98aa4146101035780635cb80c5d146100fe5780636def4ce7146100f95780637437ff9f146100f457806379ba5097146100ef57806389933a51146100ea5780638da5cb5b146100e557806390423fa2146100e0578063df0aa9e9146100db578063e8d80861146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611cb5565b611c04565b611b95565b6111fc565b611095565b61104e565b610f51565b610e09565b610d96565b610d08565b610ab1565b610a6c565b6105e3565b610357565b6102c9565b3461017a57600036600319011261017a576080610132611cfb565b61017860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b634e487b7160e01b600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176101b157604052565b61017f565b6080810190811067ffffffffffffffff8211176101b157604052565b6040810190811067ffffffffffffffff8211176101b157604052565b90601f8019910116810190811067ffffffffffffffff8211176101b157604052565b604051906102206101c0836101ee565b565b60405190610220610160836101ee565b6040519061022060a0836101ee565b6040519061022060c0836101ee565b67ffffffffffffffff81116101b157601f01601f191660200190565b6040519061027b6020836101ee565b60008252565b60005b8381106102945750506000910152565b8181015183820152602001610284565b906020916102bd81518092818552858086019101610281565b601f01601f1916010190565b3461017a57600036600319011261017a5761032860408051906102ec81836101ee565b601082527f4f6e52616d7020322e302e302d646576000000000000000000000000000000006020830152519182916020835260208301906102a4565b0390f35b67ffffffffffffffff81160361017a57565b35906102208261032c565b908160a091031261017a5790565b3461017a57604036600319011261017a576004356103748161032c565b60243567ffffffffffffffff811161017a57610394903690600401610349565b6103b28267ffffffffffffffff166000526006602052604060002090565b918254916001600160a01b036103dd6103d1856001600160a01b031690565b6001600160a01b031690565b161561054857604081019160016103f48484611de2565b90501161051e57610328946104979461048361043c6104166080870187611e18565b6104236020890189611e18565b9050159182610508575b61043687611f7b565b88612f21565b95610445612095565b61044f8288611de2565b90506104b7575b6040880161047981519260608b0193845161047360028b01611e4b565b9161349b565b9092525285611de2565b151590506104a95760f01c90505b906139c3565b60405190815292839250602083019150565b506001015461ffff16610491565b506105036104d66104d16104cb848a611de2565b906120f9565b612107565b60206104e56104cb858b611de2565b01356104f660208b015161ffff1690565b9060e08b0151928961326c565b610456565b91506105148989611de2565b905015159161042d565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6000fd5b90602060031983011261017a5760043567ffffffffffffffff811161017a5760040160009280601f830112156105df5781359367ffffffffffffffff85116105dc57506020808301928560051b01011161017a579190565b80fd5b8380fd5b3461017a576105f136610584565b906105fa6142ce565b6000915b80831061060757005b610612838284612111565b9261061c84612134565b67ffffffffffffffff81169081158015610a28575b8015610a12575b80156109f9575b6109c257856108609161087a6108708361086a6106b160e083019561069761069161066a898761216b565b61068961067f6101008a95949501809a61216b565b94909236916121a1565b9236916121a1565b9061430c565b67ffffffffffffffff166000526006602052604060002090565b9687956106ec6106c360208a01612107565b88906001600160a01b031673ffffffffffffffffffffffffffffffffffffffff19825416179055565b61085a61083c60c060408b019a6107536107058d612149565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b6107b161076260808301612203565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b6107d760016107c260a08401612203565b9c019b8c9061ffff1661ffff19825416179055565b6108368d6107e76060840161220d565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b01612161565b885465ffffffff0000191660109190911b65ffffffff000016178855565b8d61216b565b90600389016122ff565b8a61216b565b90600286016122ff565b61012088019061088c6103d183612107565b156109b15761089d6108e792612107565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b61014087019061090b6109056108fd848b611e18565b939050612149565b60ff1690565b03610986579561096c7f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb2692610951610947600198999a85611e18565b90600484016123f8565b61095a8561581b565b505460a01c67ffffffffffffffff1690565b61097b60405192839283612593565b0390a20191906105fe565b6109909087611e18565b906109ad6040519283926303aeba3960e41b8452600484016123a2565b0390fd5b6306b7c75960e31b60005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a0b60c08801612161565b161561063f565b5060ff610a2160408801612149565b1615610638565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610631565b6001600160a01b0381160361017a57565b3461017a57604036600319011261017a57610a8860043561032c565b610a9c602435610a9781610a5b565b612716565b6040516001600160a01b039091168152602090f35b3461017a57610abf36610584565b906001600160a01b0360035416918215610bd85760005b818110610adf57005b610af06103d16104d1838587614424565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610bd3576001948892600092610ba3575b5081610b57575b5050505001610ad6565b81610b877f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610b979461589f565b6040519081529081906020820190565b0390a338858180610b4d565b610bc591925060203d8111610bcc575b610bbd81836101ee565b810190614434565b9038610b46565b503d610bb3565b61270a565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b906020808351928381520192019060005b818110610c205750505090565b82516001600160a01b0316845260209384019390920191600101610c13565b80516001600160a01b03168252610d059160208281015167ffffffffffffffff169082015260408281015160ff169082015260608281015115159082015260808281015161ffff169082015260a08281015161ffff169082015260c08281015163ffffffff169082015260e0828101516001600160a01b031690820152610140610cf3610cdf610100850151610160610100860152610160850190610c02565b610120850151848203610120860152610c02565b920151906101408184039101526102a4565b90565b3461017a57602036600319011261017a5767ffffffffffffffff600435610d2e8161032c565b610d366127b5565b50166000526006602052610328610d506040600020611f7b565b604051918291602083526020830190610c3f565b6102209092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461017a57600036600319011261017a57600060408051610db681610195565b8281528260208201520152610328604051610dd081610195565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610d64565b3461017a57600036600319011261017a576000546001600160a01b0381163303610e855773ffffffffffffffffffffffffffffffffffffffff19600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b6040810160408252825180915260206060830193019060005b818110610f31575050506020818303910152815180825260208201916020808360051b8301019401926000915b838310610f0457505050505090565b9091929394602080610f22600193601f198682030187528951610c3f565b97019301930191939290610ef5565b825167ffffffffffffffff16855260209485019490920191600101610ec8565b3461017a57600036600319011261017a57600454610f6e8161207d565b90610f7c60405192836101ee565b808252601f19610f8b8261207d565b0160005b818110611037575050610fa1816120b1565b9060005b818110610fbd57505061032860405192839283610eaf565b80610ff5610fdc610fcf60019461595b565b67ffffffffffffffff1690565b610fe6838761282d565b9067ffffffffffffffff169052565b61101b611016610697611008848861282d565b5167ffffffffffffffff1690565b611f7b565b611025828761282d565b52611030818661282d565b5001610fa5565b6020906110426127b5565b82828701015201610f8f565b3461017a57600036600319011261017a5760206001600160a01b0360015416604051908152f35b359061022082610a5b565b8015150361017a57565b359061022082611080565b3461017a57606036600319011261017a5760006040516110b481610195565b6004356110c081610a5b565b81526024356110ce81611080565b60208201908152604435906110e282610a5b565b604083019182526110f16142ce565b6001600160a01b038351161580156111f2575b6111e3576001600160a01b03839261119b6111c69361116a847f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9851166001600160a01b031673ffffffffffffffffffffffffffffffffffffffff196002541617600255565b51151560ff60a01b1974ff000000000000000000000000000000000000000060025492151560a01b16911617600255565b51166001600160a01b031673ffffffffffffffffffffffffffffffffffffffff196003541617600355565b6111ce611cfb565b6111dd60405192839283614443565b0390a180f35b6004846306b7c75960e31b8152fd5b5080511515611104565b3461017a57608036600319011261017a5761121860043561032c565b60243567ffffffffffffffff811161017a57611238903690600401610349565b60443590611247606435610a5b565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610bd357600091611b66575b50611b295760025460a01c60ff16611aff5761130e7401000000000000000000000000000000000000000060ff60a01b196002541617600255565b61132e60043567ffffffffffffffff166000526006602052604060002090565b906001600160a01b036064351615611ad55781546113556103d16001600160a01b03831681565b3303611aab578161136a608086940182611e18565b6113776020840184611e18565b9050159081611a92575b61138a87611f7b565b6113979390600435612f21565b9160a01c67ffffffffffffffff166113ae90612856565b84547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff000000000000000000000000000000000000000016178555825163ffffffff169460208401516114119061ffff1690565b6040805130602080830191909152815297919061142e90896101ee565b604080516064356001600160a01b031660208083019190915281529890611455908a6101ee565b61145f8680611e18565b85549a9160e08c901c60ff1691369061147792612875565b9060ff16611484916144f4565b60a08901519161149760408a018a611de2565b6114a191506128ee565b936114af60208b018b611e18565b9690976114ba610210565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529a67ffffffffffffffff6004351660208d015267ffffffffffffffff1660408c0152600060608c015263ffffffff1660808b015261ffff1660a08a0152600060c08a015260e089015261153c60048801611ed9565b610100890152610120880152610140870152610160860152610180850152369061156592612875565b6101a0830152611573612095565b6115806040850185611de2565b9050611a18575b83611600926115cf6115aa88946040860151606087015161047360028701611e4b565b60608601528060408601526115c960808601516001600160a01b031690565b90614565565b60c08601526115dc61293e565b986115ea6040840184611de2565b15159050611a0a5760f01c90505b6004356139c3565b63ffffffff90911660608401526020870195918652116119e0576116258451836145f3565b6116326040830183611de2565b90506118c0575b61164581959295614ed3565b808352602081519101209061165e6040850151516129c2565b6040840190815294606087019360005b604087015180518210156118115760206116a16103d16103d1611694866116cf9661282d565b516001600160a01b031690565b6116af8460608c015161282d565b5190604051808095819463958021a760e01b835260043560048401612a0c565b03915afa8015610bd3576001600160a01b03916000916117e3575b501680156117a357906000888c9388838961174d611716888f61170e606091612107565b98015161282d565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612b57565b03925af18015610bd3578161177b91600194600091611782575b508b5190611775838361282d565b5261282d565b500161166e565b61179d913d8091833e61179581836101ee565b810190612a6f565b38611767565b6105806117b76116948460408c015161282d565b6341e3ac5360e11b6000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611804915060203d811161180a575b6117fc81836101ee565b8101906126f5565b386116ea565b503d6117f2565b6103288680858d7f371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc8e8a6118448f612107565b936118526040820182611de2565b1590506118b457602061186f6104cb83604061189f950190611de2565b0135955b51915192516040519384936001600160a01b03606435169867ffffffffffffffff600435169886612d22565b0390a4610b8760ff60a01b1960025416600255565b5061189f600095611873565b61192261190c6118d66104cb6040860186611de2565b60c08601805151156119c55751905b602087015161ffff169060e0880151926064359161190760043591369061295d565b61488e565b6101808301519061191c82612820565b52612820565b506119446040611938865182870151519061282d565b51015163ffffffff1690565b60a0611954610180840151612820565b5101515163ffffffff8216811161196c575050611639565b61058092506119846104d16104cb6040870187611de2565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b506119da6119d38680611e18565b3691612875565b906118e5565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff166115f8565b5093506001611a2a6040840184611de2565b90500361051e57611600838388966115cf6115aa611a86611a546104d16104cb6040880188611de2565b6020611a666104cb6040890189611de2565b0135611a77602089015161ffff1690565b9060e08901519260043561326c565b94505050925050611587565b9050611aa16040840184611de2565b9050151590611381565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105806004359067ffffffffffffffff60249216600452565b611b88915060203d602011611b8e575b611b8081836101ee565b810190612841565b386112d3565b503d611b76565b3461017a57602036600319011261017a5767ffffffffffffffff600435611bbb8161032c565b166000526006602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611bff5760405167ffffffffffffffff9091168152602090f35b612217565b3461017a57602036600319011261017a576001600160a01b03600435611c2981610a5b565b611c316142ce565b16338114611c8b578073ffffffffffffffffffffffffffffffffffffffff1960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461017a57602036600319011261017a57611cd160043561032c565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611d0b816101b6565b8281528260208201528260408201520152604051611d28816101b6565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a57602001918160061b3603831361017a57565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a5760200191813603831361017a57565b906040519182815491828252602082019060005260206000209260005b818110611e7d575050610220925003836101ee565b84546001600160a01b0316835260019485019487945060209093019201611e68565b90600182811c92168015611ecf575b6020831014611eb957565b634e487b7160e01b600052602260045260246000fd5b91607f1691611eae565b9060405191826000825492611eed84611e9f565b8084529360018116908115611f595750600114611f12575b50610220925003836101ee565b90506000929192526020600020906000915b818310611f3d5750509060206102209282010138611f05565b6020919350806001915483858901015201910190918492611f24565b90506020925061022094915060ff191682840152151560051b82010138611f05565b906120756004611f89610222565b93611ff8611fed8254611fb2611fa5826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c166040890152611fe760e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b61204b61203b600183015461201c6120118261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b61205760028201611e4b565b61010086015261206960038201611e4b565b61012086015201611ed9565b610140830152565b67ffffffffffffffff81116101b15760051b60200190565b604051906120a46020836101ee565b6000808352366020840137565b906120bb8261207d565b6120c860405191826101ee565b82815280926120d9601f199161207d565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b90156121025790565b6120e3565b35610d0581610a5b565b91908110156121025760051b8101359061015e198136030182121561017a570190565b35610d058161032c565b60ff81160361017a57565b35610d058161213e565b63ffffffff81160361017a57565b35610d0581612153565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a57602001918160051b3603831361017a57565b9291906121ad8161207d565b936121bb60405195866101ee565b602085838152019160051b810192831161017a57905b8282106121dd57505050565b6020809183356121ec81610a5b565b8152019101906121d1565b61ffff81160361017a57565b35610d05816121f7565b35610d0581611080565b634e487b7160e01b600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611bff57565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611bff57565b908160031b9180830460081490151715611bff57565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611bff57565b81810292918115918404141715611bff57565b8181106122f3575050565b600081556001016122e8565b9067ffffffffffffffff83116101b1576801000000000000000083116101b1578154838355808410612363575b5090600052602060002060005b8381106123465750505050565b600190602084359461235786610a5b565b01938184015501612339565b61237b908360005284602060002091820191016122e8565b3861232c565b908060209392818452848401376000828201840152601f01601f1916010190565b916020610d05938181520191612381565b9190601f81116123c257505050565b610220926000526020600020906020601f840160051c830193106123ee575b601f0160051c01906122e8565b90915081906123e1565b90929167ffffffffffffffff81116101b15761241e816124188454611e9f565b846123b3565b6000601f821160011461245f578190612450939495600092612454575b50508160011b916000199060031b1c19161790565b9055565b01359050388061243b565b601f1982169461247484600052602060002090565b91805b8781106124af575083600195969710612495575b505050811b019055565b0135600019600384901b60f8161c1916905538808061248b565b90926020600181928686013581550194019101612477565b35906102208261213e565b3590610220826121f7565b359061022082612153565b9035601e198236030181121561017a57016020813591019167ffffffffffffffff821161017a578160051b3603831361017a57565b9160209082815201919060005b8181106125375750505090565b9091926020806001926001600160a01b03873561255381610a5b565b16815201940192910161252a565b9035601e198236030181121561017a57016020813591019167ffffffffffffffff821161017a57813603831361017a57565b67ffffffffffffffff610d059392168152604060208201526125c9604082016125bb8461033e565b67ffffffffffffffff169052565b6125e86125d860208401611075565b6001600160a01b03166060830152565b6126016125f7604084016124c7565b60ff166080830152565b6126196126106060840161108a565b151560a0830152565b612633612628608084016124d2565b61ffff1660c0830152565b61264d61264260a084016124d2565b61ffff1660e0830152565b61266a61265c60c084016124dd565b63ffffffff16610100830152565b6126e26126b561269461268060e08601866124e8565b6101606101208701526101a086019161251d565b6126a26101008601866124e8565b858303603f19016101408701529061251d565b926126d76126c66101208301611075565b6001600160a01b0316610160850152565b610140810190612561565b91610180603f1982860301910152612381565b9081602091031261017a5751610d0581610a5b565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610bd3576001600160a01b039160009161279857501690565b6127b1915060203d60201161180a576117fc81836101ee565b1690565b60405190610160820182811067ffffffffffffffff8211176101b15760405260606101408360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e082015282610100820152826101208201520152565b8051156121025760200190565b80518210156121025760209160051b010190565b9081602091031261017a5751610d0581611080565b67ffffffffffffffff1667ffffffffffffffff8114611bff5760010190565b92919261288182610250565b9161288f60405193846101ee565b82948184528183011161017a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101b157604052606060a0836000815282602082015282604082015282808201528260808201520152565b906128f88261207d565b61290560405191826101ee565b8281528092612916601f199161207d565b019060005b82811061292757505050565b6020906129326128ac565b8282850101520161291b565b6040519061294b82610195565b60606040838281528260208201520152565b919082604091031261017a57604051612975816101d2565b6020808294803561298581610a5b565b84520135910152565b6040519061299d6020836101ee565b600080835282815b8281106129b157505050565b8060606020809385010152016129a5565b906129cc8261207d565b6129d960405191826101ee565b82815280926129ea601f199161207d565b019060005b8281106129fb57505050565b8060606020809385010152016129ef565b60409067ffffffffffffffff610d05949316815281602082015201906102a4565b81601f8201121561017a578051612a4381610250565b92612a5160405194856101ee565b8184526020828401011161017a57610d059160208085019101610281565b9060208282031261017a57815167ffffffffffffffff811161017a57610d059201612a2d565b9080602083519182815201916020808360051b8301019401926000915b838310612ac157505050505090565b9091929394602080612b48600193601f198682030187528951908151815260a0612b37612b25612b13612b018887015160c08a88015260c08701906102a4565b604087015186820360408801526102a4565b606086015185820360608701526102a4565b608085015184820360808601526102a4565b9201519060a08184039101526102a4565b97019301930191939290612ab2565b919390610d059593612c9f612cb79260a08652612b8160a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612c8a612c72612c5a612c42612c2a612c148c61026060e08a0151916101c061018082015201906102a4565b6101008801518d8203609f1901888f01526102a4565b6101208701518c8203609f19016101c08e01526102a4565b6101408601518b8203609f19016101e08d01526102a4565b6101608501518a8203609f19016102008c01526102a4565b610180840151898203609f19016102208b0152612a95565b910151868203609f19016102408801526102a4565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102a4565b9080602083519182815201916020808360051b8301019401926000915b838310612cf557505050505090565b9091929394602080612d13600193601f1986820301875289516102a4565b97019301930191939290612ce6565b95949290916001600160a01b03612d4c93168752602087015260a0604087015260a08601906102a4565b938085036060820152825180865260208601906020808260051b8901019501916000905b828210612d8e5750505050610d059394506080818403910152612cc9565b90919295602080612def600193601f198d820301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102a4565b980192019201909291612d70565b60405190610100820182811067ffffffffffffffff8211176101b157604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161017a5790600490565b9093929384831161017a57841161017a578101920390565b919091356001600160e01b031981169260048110612e90575050565b6001600160e01b0319929350829060040360031b1b161690565b9160608383031261017a57825167ffffffffffffffff811161017a5782612ed2918501612a2d565b926020810151612ee181612153565b92604082015167ffffffffffffffff811161017a57610d059201612a2d565b60409067ffffffffffffffff610d0595931681528160208201520191612381565b91929092612f2d612dfd565b6004831015806130f6575b15613040575090612f4891615190565b92612f56604085015161534f565b80613025575b60408401612f7981519260608701938451610120880151916153da565b9092525260c083015151612fd5575b50608082016001600160a01b03612fa682516001600160a01b031690565b1615612fb157505090565b612fc860e0610d059301516001600160a01b031690565b6001600160a01b03169052565b612fe9612fe56060840151151590565b1590565b15612f88577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff613039845163ffffffff1690565b1615612f5c565b9491600090613098926130616103d16103d16002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612f00565b03915afa8015610bd3576000906000906000906130ca575b60a088015263ffffffff16865290505b60c0850152612f56565b5050506130ec6130c0913d806000833e6130e481836101ee565b810190612eaa565b91925082916130b0565b5063534eea5560e11b6001600160e01b031961311b6131158686612e4e565b90612e74565b1614612f38565b60208183031261017a5780519067ffffffffffffffff821161017a57019080601f8301121561017a5781516131568161207d565b9261316460405194856101ee565b81845260208085019260051b82010192831161017a57602001905b82821061318c5750505090565b60208091835161319b81610a5b565b81520191019061317f565b95949060009460a09467ffffffffffffffff6131ed956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102a4565b930152565b9060028201809211611bff57565b9060018201809211611bff57565b6001019081600111611bff57565b9060148201809211611bff57565b90600c8201809211611bff57565b91908201809211611bff57565b6000198114611bff5760010190565b80548210156121025760005260206000200190600090565b929394919060036132918567ffffffffffffffff166000526006602052604060002090565b01936001600160a01b036132a6818416612716565b1691821561347e576040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bd35760009161345f575b501561344f57613325600095969798604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016131a6565b03915afa928315610bd35760009361342a575b5082511561341f5782519061335761335284548094613238565b6120b1565b906000928394845b87518110156133be57613375611694828a61282d565b6001600160a01b038116156133b257906133ac60019261339e6133978a613245565b998961282d565b906001600160a01b03169052565b0161335f565b509550600180966133ac565b5091955091936133d0575b5050815290565b60005b8281106133e057506133c9565b806134196134066133f360019486613254565b90546001600160a01b039160031b1c1690565b61339e61341288613245565b978961282d565b016133d3565b9150610d0590611e4b565b6134489193503d806000833e61344081836101ee565b810190613122565b9138613338565b505050509250610d059150611e4b565b613478915060203d602011611b8e57611b8081836101ee565b386132e0565b635f8b555b60e11b6000526001600160a01b031660045260246000fd5b939192936134b76134af8251865190613238565b865190613238565b906134ca6134c4836120b1565b926129c2565b94600096875b835189101561353057886135266135196001936135016134f76116948e9f9d9e9d8b61282d565b61339e838c61282d565b61351f61350e858c61282d565b519180938491613245565b9c61282d565b528b61282d565b50019796956134d0565b959250929350955060005b86518110156135c457613551611694828961282d565b60006001600160a01b038216815b888110613598575b505090600192911561357b575b500161353b565b6135929061339e61358b89613245565b988861282d565b38613574565b816135a96103d1611694848c61282d565b146135b65760010161355f565b506001915081905038613567565b509390945060005b8551811015613655576135e2611694828861282d565b60006001600160a01b038216815b878110613629575b505090600192911561360c575b50016135cc565b6136239061339e61361c88613245565b978761282d565b38613605565b8161363a6103d1611694848b61282d565b14613647576001016135f0565b5060019150819050386135f8565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101b15760405260606080836000815260006020820152600060408201526000838201520152565b906136aa8261207d565b6136b760405191826101ee565b82815280926136c8601f199161207d565b019060005b8281106136d957505050565b6020906136e4613661565b828285010152016136cd565b9081606091031261017a578051613706816121f7565b916040602083015161371781612153565b920151610d0581612153565b9160209082815201919060005b81811061373d5750505090565b9091926040806001926001600160a01b03873561375981610a5b565b16815260208781013590820152019401929101613730565b949391929067ffffffffffffffff168552608060208601526137ca6137ab6137998580612561565b60a060808a0152610120890191612381565b6137b86020860186612561565b888303607f190160a08a015290612381565b6040840135601e198536030181121561017a578401916020833593019167ffffffffffffffff841161017a578360061b3603831361017a576102209561385261382961386593606097613873978d60c0607f1982860301910152613723565b91613848613838888301611075565b6001600160a01b031660e08d0152565b6080810190612561565b8a8303607f19016101008c015290612381565b9087820360408901526102a4565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611bff57565b908160a091031261017a5780519160208201516138b481612153565b9160408101516138c381612153565b91608060608301516138d4816121f7565b920151610d0581611080565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610d059b9a9616885216602087015260408601521660608401521660808201528160a082015201906102a4565b9081606091031261017a57805161370681612153565b600119810191908211611bff57565b600019810191908211611bff57565b600219810191908211611bff57565b91908203918211611bff57565b919082608091031261017a57815161398f81612153565b916020810151916060604083015192015190565b81156139ad570490565b634e487b7160e01b600052601260045260246000fd5b93829460009060009560408101946139fd6139f86139f38851516139eb60408c01809c611de2565b919050613238565b6131f2565b6136a0565b9660009586955b88518051881015613c2f576103d16103d16116948a613a229461282d565b613a5660206060880192613a378b855161282d565b51908a60405180958194829363958021a760e01b845260048401612a0c565b03915afa8015610bd3576001600160a01b0391600091613c11575b50168015613bd6579060608e9392613a8a8b845161282d565b5190613a9b60208b015161ffff1690565b958b613ad6604051988995869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613771565b03915afa8015610bd357600193613b73938b8f8f95600080958197613b7c575b509083929161ffff613b1e85613b17611694613b6799613b6d9d9e5161282d565b945161282d565b5191613b3a613b2b610232565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b1660408501521660608301526080820152611775838361282d565b5061387e565b9961387e565b96019596613a04565b613b6d9750611694965084939291509361ffff613b1e82613b17613bb9613b679960603d8111613bcf575b613bb181836101ee565b8101906136f0565b9c9196909c9d5050505050505090919293613af6565b503d613ba7565b61058088613be86116948c8f5161282d565b6341e3ac5360e11b6000526001600160a01b031660045267ffffffffffffffff16602452604490565b613c29915060203d811161180a576117fc81836101ee565b38613a71565b50919a9496929395509897968a613c468187611de2565b9050613f5a575b50508651613c5a9061393e565b99613c686020860186611e18565b91613c74915086611de2565b9560609150019486613c8587612107565b91613c90938a615608565b613c9a8b8961282d565b52613ca58a8861282d565b50613cb08a8861282d565b516020015163ffffffff16613cc49161387e565b90613ccf8a8861282d565b516040015163ffffffff16613ce39161387e565b91613cec610232565b33815290600060208301819052604083015261ffff166060820152613d0f61026c565b60808201528651613d1f9061394d565b90613d2a828961282d565b52613d35908761282d565b506002546001600160a01b031692613d4c90612107565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610bd357600096600097600094600092613f1d575b506000965b8651881015613e9d57613e26600191613dee88613de98761222d565b6139a3565b613e076060613dfd8d8d61282d565b51019182516122d5565b9052858a14613e2e575b6060613e1d8b8b61282d565b51015190613238565b970196613dcd565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613e6160808c01516001600160a01b031690565b1603613e6f575b5050613e11565b613de9613e7b9261225d565b613e946060613e8a8d8d61282d565b5101918251613238565b90528b88613e68565b9796509750505050613ed97f000000000000000000000000000000000000000000000000000000000000000091613de963ffffffff841661225d565b8411613ee55750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b9298505050613f4591925060803d608011613f53575b613f3d81836101ee565b810190613978565b919790939290919038613dc8565b503d613f33565b610a976103d16104d16104cb613f73948a989698611de2565b926001600160a01b03600091515194169060e08801908151613f93610232565b6001600160a01b0385168152908260208301528260408301528260608301526080820152613fc1878d61282d565b52613fcc868c61282d565b506040516301ffc9a760e01b8152633317103160e01b600482015291602083602481875afa8015610bd3578f948c89968f96948d948f9688916142af575b506141a8575b5050505050501561404f575b6119386140406140479561403a602061193861403a9760409761282d565b9061387e565b958b61282d565b90388a613c4d565b50506140d59160608c6140806104d16104cb6140796103d16103d16002546001600160a01b031690565b938b611de2565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610bd35761193861404060409261403a60206119388f8b906140479b61403a9a60008060009261416a575b63ffffffff92935061415790606061411f888861282d565b51019461414c8a6141308a8a61282d565b510191604061413f8b8b61282d565b51019063ffffffff169052565b9063ffffffff169052565b169052975097505050509550505061401c565b50505063ffffffff6141966141579260603d6060116141a1575b61418e81836101ee565b810190613928565b909350915082614107565b503d614184565b8495985060a09697506141f260206141e86060826141df6104cb6141d86104d16104cb8b61422b9c9d9e9f611de2565b998d611de2565b01359901612107565b9a015161ffff1690565b905190604051998a97889687967f2c063404000000000000000000000000000000000000000000000000000000008852600488016138e0565b03915afa8015610bd3578592828c939181908294614273575b50614267906060614255888861282d565b51019261414c60206141308a8a61282d565b5288888f8c8138614010565b915050614267925061429d915060a03d60a0116142a8575b61429581836101ee565b810190613898565b949192919050614244565b503d61428b565b6142c8915060203d602011611b8e57611b8081836101ee565b3861400a565b6001600160a01b036001541633036142e257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80519161431a815184613238565b9283156143fa5760005b848110614332575050505050565b818110156143df57614347611694828661282d565b6001600160a01b0381168015610bd85761436083613200565b87811061437257505050600101614324565b848110156143bc576001600160a01b0361438f611694838a61282d565b16821461439e57600101614360565b630285c9b960e61b6000526001600160a01b03831660045260246000fd5b6001600160a01b036143da6116946143d4888561396b565b8961282d565b61438f565b6143f56116946143ef848461396b565b8561282d565b614347565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156121025760051b0190565b9081602091031261017a575190565b9160806102209294936144938160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610d059281815201906102a4565b906144d682610250565b6144e360405191826101ee565b82815280926120d9601f1991610250565b908151602082108061455b575b61452a570361450d5790565b6109ad906040519182916303aeba3960e41b8352600483016144bb565b50906020810151908161453c8461228f565b1c61450d575061454b826144cc565b9160200360031b1b602082015290565b5060208114614501565b918251601481029080820460141490151715611bff5761458761458c9161320e565b61321c565b9061459e6145998361322a565b6144cc565b9060146145aa83612820565b5360009260215b86518510156145dc5760146001916145cc611694888b61282d565b60601b81870152019401936145b1565b919550936020935060601b90820152828152012090565b906146036103d160608401612107565b614614600019936040810190611de2565b905061468c575b614625825161394d565b9260005b848110614637575050505050565b80826001921461468757606061464d828761282d565b51015180156146815761467b90614675614667848961282d565b51516001600160a01b031690565b8661589f565b01614629565b5061467b565b61467b565b9150614698815161395c565b916146a6614667848461282d565b6040516301ffc9a760e01b8152633317103160e01b60048201526020816024816001600160a01b0386165afa908115610bd35760009161470d575b506146ed575b5061461b565b6147079060606146fd868661282d565b510151908361589f565b386146e7565b614726915060203d602011611b8e57611b8081836101ee565b386146e1565b60405190614739826101d2565b60606020838281520152565b919060408382031261017a576040519061475e826101d2565b8193805167ffffffffffffffff811161017a578261477d918301612a2d565b835260208101519167ffffffffffffffff831161017a576020926147a19201612a2d565b910152565b9060208282031261017a57815167ffffffffffffffff811161017a57610d059201614745565b9060806001600160a01b03816147eb855160a0865260a08601906102a4565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610d059281815201906147cc565b919060408382031261017a57825167ffffffffffffffff811161017a5760209161485b918501614745565b92015190565b61ffff61487a610d0595936060845260608401906147cc565b9316602082015260408184039101526102a4565b90919392959461489c6128ac565b5060208201805115614cd2576148bf610a976103d185516001600160a01b031690565b946001600160a01b038616916040516301ffc9a760e01b81526020818061490d60048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610bd357600091614cb3575b5015614c835761499988999a82519261493861472c565b505161498461494e89516001600160a01b031690565b926040614959610232565b9e8f9081526149758d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bd357600091614c64575b5015614b6c5750906000929183614a139899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614861565b03925af18015610bd357600094600091614b24575b50614af9614a629596614a9c614a70614a8e956116946020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b6101ee565b6040519586918683019190916001600160a01b036020820193169052565b03601f1981018652856101ee565b614ad8610905614ace614ade8951614ad8610905614ace8c67ffffffffffffffff166000526006602052604060002090565b5460e01c60ff1690565b906144f4565b9767ffffffffffffffff166000526006602052604060002090565b93015193614b05610241565b958652602086015260408501526060840152608083015260a082015290565b614a6295506020915061169496614a9c614a70614a8e95614b5a614af9953d806000833e614b5281836101ee565b810190614830565b9b909b96505095505050969550614a28565b9793929061ffff16614c3a5751614c1057614bbb6000939184926040519586809481937f9a4575b90000000000000000000000000000000000000000000000000000000083526004830161481f565b03925af1908115610bd357614a6295614a9c614a70614af993611694602096614a8e98600091614bed575b5099614a44565b614c0a91503d806000833e614c0281836101ee565b8101906147a6565b38614be6565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7f6cae288c0000000000000000000000000000000000000000000000000000000060005260046000fd5b614c7d915060203d602011611b8e57611b8081836101ee565b386149cb565b610580614c9786516001600160a01b031690565b635f8b555b60e11b6000526001600160a01b0316600452602490565b614ccc915060203d602011611b8e57611b8081836101ee565b38614921565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592614d9a947fffffffffffffffff0000000000000000000000000000000000000000000000006001600160e01b031994928186948160439d9b97600160f81b8e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b90614db660209282815194859201610281565b0190565b9360019694936001600160f81b03198094899896828a9660f81b168152614dea8251809360208985019101610281565b019160f81b1683820152614e08825180936020600285019101610281565b01019160f81b1683820152614e27825180936020600285019101610281565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681956001600160f81b0319610d059f9e9c9860f81b168152614e848251809360208985019101610281565b019160f01b1683820152614ea2825180936020600385019101610281565b01019160f01b1683820152614ec08251809360208985019101610281565b01019160f01b1660028201520190614da3565b60e081019060ff8251511161517a57610100810160ff815151116151645761012082019260ff8451511161514e57610140830160ff815151116151385761016084019061ffff825151116151225761018085019360018551511161510a576101a086019361ffff855151116150f4576060955180516150b8575b50865167ffffffffffffffff16602088015167ffffffffffffffff16906040890151614f809067ffffffffffffffff1690565b986060810151614f939063ffffffff1690565b906080810151614fa69063ffffffff1690565b60a082015161ffff169160c00151926040519c8d966020880196614fc997614cfc565b03601f1981018852614fdb90886101ee565b51908151614fe99060ff1690565b9051805160ff169851908151614fff9060ff1690565b906040519a8b95602087019561501496614dba565b03601f198101875261502690876101ee565b519182516150349060ff1690565b9151805161ffff1694805161504a9061ffff1690565b925192835161505a9061ffff1690565b92604051978897602089019761506f98614e2d565b03601f198101825261508190826101ee565b6040519283926020840161509491614da3565b61509d91614da3565b6150a691614da3565b03601f1981018252610d0590826101ee565b6150cd9196506150c790612820565b51615a67565b9461ffff8651116150de5738614f4d565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b600052602660045260246000fd5b635a102da160e11b6000526105806024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b90615199612dfd565b916011821061533457803563534eea5560e11b6001600160e01b03198216036152f25750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a6151e7816120b1565b604086019081526151f7826129c2565b906060870191825260005b8381106152a6575050505061526a838361526061525461524a615243615230615274988761527e9c9b615b8e565b6001600160a01b0390911660808d015290565b8585615c32565b9291903691612875565b60a08a01528383615c81565b9491903691612875565b60c0880152615c32565b9391903691612875565b60e0840152810361528d575090565b63d9437f9d60e01b600052600360045260245260446000fd5b806001916152eb6152d56152ce6152c16152e59a8d8d615b8e565b919061339e868a5161282d565b8b8b615c32565b9391889a919a51949a3691612875565b9261282d565b5201615202565b7f55a0e02c0000000000000000000000000000000000000000000000000000000060005263534eea5560e11b6004526001600160e01b03191660245260446000fd5b63d9437f9d60e01b6000526002600452602482905260446000fd5b80519060005b82811061536157505050565b60018101808211611bff575b83811061537d5750600101615355565b6001600160a01b0361538f838561282d565b51166153a16103d1611694848761282d565b146153ae5760010161536d565b6105806153be611694848661282d565b630285c9b960e61b6000526001600160a01b0316600452602490565b9093919281511561556657506153ef8161534f565b80519260005b84811061540457509093925050565b6154146103d1611694838661282d565b15615421576001016153f5565b93919461543b6133526154338561394d565b845190613238565b9261545861545361544b8361394d565b855190613238565b6129c2565b968792600097885b8481106155085750505050505060005b81518110156154fb576000805b8681106154bc575b50906001916154b7576154b161549e611694838661282d565b61339e6154aa89613245565b988761282d565b01615470565b6154b1565b6154c9611694828761282d565b6001600160a01b036154e16103d1611694878961282d565b9116146154f05760010161547d565b506001905080615485565b5050909180825283529190565b90919293949882821461555c579061554e615541836155348b61339e6001976152e5611694898e61282d565b61554761350e858961282d565b9e61282d565b528c61282d565b505b01908994939291615460565b9850600190615550565b91935050156155815750615578612095565b90610d0561298e565b90610d0582516129c2565b60011b906101fe60fe831692168203611bff57565b9081602091031261017a5751610d05816121f7565b936155f360809461ffff6001600160a01b039567ffffffffffffffff615601969b9a9b16895216602088015260a0604088015260a0870190610c02565b9085820360608701526102a4565b9416910152565b92919092615614613661565b506156338167ffffffffffffffff166000526006602052604060002090565b805490959060e01c60ff169160808501928351615656906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff166156769161387e565b9661568290608d613238565b9460a087019586515161569491613238565b6156a060ff84166122a5565b6156a991613238565b916156b39061558c565b60ff166156c1906067613238565b6156ca916122d5565b6156d391613238565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b038716036157605750505061ffff9250615752906157456000935b5195615738615729610232565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b6157766103d1602094976001600160a01b031690565b9060406157878583015161ffff1690565b910151926157c7875198604051998a96879586957ff2388958000000000000000000000000000000000000000000000000000000008752600487016155b6565b03915afa908115610bd3576157456157529261ffff956000916157ec575b509361571c565b61580e915060203d602011615814575b61580681836101ee565b8101906155a1565b386157e5565b503d6157fc565b8060005260056020526040600020541560001461589957600454680100000000000000008110156101b15760018101600455600060045482101561210257600490527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01819055600454906000526005602052604060002055600190565b50600090565b91602091600091604051906001600160a01b03858301937fa9059cbb0000000000000000000000000000000000000000000000000000000085521660248301526044820152604481526158f36064826101ee565b519082855af11561270a576000513d61595257506001600160a01b0381163b155b61591b5750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615914565b6004548110156121025760046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b015490565b966002615a3b97615a206022610d059f9e9c9799600199859f9b6001600160f81b031990615a209f82615a209c615a279c600160f81b8452600184015260f81b1660218201526159e98251809360208985019101610281565b019160f81b1683820152615a07825180936020602385019101610281565b010191888301906001600160f81b03199060f81b169052565b0190614da3565b80926001600160f81b03199060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff82515111615b7857604081019160ff83515111615b6257606082019160ff83515111615b4c57608081019260ff84515111615b365760a0820161ffff81515111615b2057610d0594615b129351945191615ac9835160ff1690565b975191615ad7835160ff1690565b945190615ae5825160ff1690565b905193615af3855160ff1690565b935196615b02885161ffff1690565b966040519c8d9b60208d01615990565b03601f1981018352826101ee565b635a102da160e11b600052602b60045260246000fd5b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b929190926001820191848311615c195781013560001a828115615c0e575060148103615be1578201938411615bc657013560601c9190565b63d9437f9d60e01b6000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b63d9437f9d60e01b600052600060045260245260446000fd5b91906002820191818311615c19578381013560f01c0160020192818411615c6657918391615c5f93612e5c565b9290929190565b63d9437f9d60e01b6000526001600452602483905260446000fd5b91906001820191818311615c19578381013560001a0160010192818411615c6657918391615c5f93612e5c56fea164736f6c634300081a000a",
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

func (_OnRamp *OnRampCaller) GetAllDestChainConfigs(opts *bind.CallOpts) ([]uint64, []OnRampDestChainConfig, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getAllDestChainConfigs")

	if err != nil {
		return *new([]uint64), *new([]OnRampDestChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]OnRampDestChainConfig)).(*[]OnRampDestChainConfig)

	return out0, out1, err

}

func (_OnRamp *OnRampSession) GetAllDestChainConfigs() ([]uint64, []OnRampDestChainConfig, error) {
	return _OnRamp.Contract.GetAllDestChainConfigs(&_OnRamp.CallOpts)
}

func (_OnRamp *OnRampCallerSession) GetAllDestChainConfigs() ([]uint64, []OnRampDestChainConfig, error) {
	return _OnRamp.Contract.GetAllDestChainConfigs(&_OnRamp.CallOpts)
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

func (_OnRamp *OnRampCaller) GetExpectedNextMessageNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _OnRamp.contract.Call(opts, &out, "getExpectedNextMessageNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_OnRamp *OnRampSession) GetExpectedNextMessageNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextMessageNumber(&_OnRamp.CallOpts, destChainSelector)
}

func (_OnRamp *OnRampCallerSession) GetExpectedNextMessageNumber(destChainSelector uint64) (uint64, error) {
	return _OnRamp.Contract.GetExpectedNextMessageNumber(&_OnRamp.CallOpts, destChainSelector)
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
	DestChainSelector              uint64
	Sender                         common.Address
	MessageId                      [32]byte
	FeeToken                       common.Address
	TokenAmountBeforeTokenPoolFees *big.Int
	EncodedMessage                 []byte
	Receipts                       []OnRampReceipt
	VerifierBlobs                  [][]byte
	Raw                            types.Log
}

func (_OnRamp *OnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sender []common.Address, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, senderRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &OnRampCCIPMessageSentIterator{contract: _OnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sender []common.Address, messageId [][32]byte) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, senderRule, messageIdRule)
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
	MessageNumber     uint64
	Config            OnRampDestChainConfigArgs
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
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_OnRamp *OnRampFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &OnRampFeeTokenWithdrawnIterator{contract: _OnRamp.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_OnRamp *OnRampFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _OnRamp.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
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
	return common.HexToHash("0x371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc")
}

func (OnRampConfigSet) Topic() common.Hash {
	return common.HexToHash("0x0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c")
}

func (OnRampDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0x99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb26")
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
	GetAllDestChainConfigs(opts *bind.CallOpts) ([]uint64, []OnRampDestChainConfig, error)

	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (OnRampDestChainConfig, error)

	GetDynamicConfig(opts *bind.CallOpts) (OnRampDynamicConfig, error)

	GetExpectedNextMessageNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (OnRampStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []OnRampDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig OnRampDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sender []common.Address, messageId [][32]byte) (*OnRampCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *OnRampCCIPMessageSent, destChainSelector []uint64, sender []common.Address, messageId [][32]byte) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*OnRampCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*OnRampConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*OnRampConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*OnRampDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *OnRampDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*OnRampDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*OnRampFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *OnRampFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*OnRampFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OnRampOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OnRampOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnRampOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OnRampOwnershipTransferred, error)

	Address() common.Address
}
