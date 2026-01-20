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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllDestChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAmountBeforeTokenPoolFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100604052346103675760405161665038819003601f8101601f191683016001600160401b0381118482101761036c578392829160405283398101039060e08212610367576080821261036757610055610382565b81519092906001600160401b03811681036103675783526020820151926001600160a01b03841684036103675760208101938452604083015163ffffffff81168103610367576040820190815260606100af8186016103a1565b83820190815293607f1901126103675760405192606084016001600160401b0381118582101761036c576040526100e8608086016103a1565b845260a08501519485151586036103675760c061010c9160208701978852016103a1565b9560408501968752331561035657600180546001600160a01b0319163317905583516001600160401b0316158015610344575b8015610332575b8015610323575b6103085792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e052815116158015610319575b6103085780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e09260006060610206610382565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff92831691166060610251610382565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a160405161629a90816103b68239608051818181610ac1015281816115c20152611e88015260a0518181816113840152611eb4015260c051818181611f0f0152612970015260e051818181611ee001526141ef0152f35b6306b7c75960e31b60005260046000fd5b508151151561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761036c57604052565b51906001600160a01b03821682036103675756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610117578063181f5a771461011257806320487ded1461010d5780632490769e1461010857806348a98aa4146101035780635cb80c5d146100fe5780636def4ce7146100f95780637437ff9f146100f457806379ba5097146100ef57806389933a51146100ea5780638da5cb5b146100e557806390423fa2146100e0578063df0aa9e9146100db578063e8d80861146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611e0a565b611d4e565b611cdf565b6112de565b61112d565b6110e6565b610fe9565b610e96565b610e23565b610d95565b610b3e565b610af9565b6105fd565b610371565b6102e3565b3461017a57600060031936011261017a576080610132611e50565b61017860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176101ca57604052565b61017f565b6080810190811067ffffffffffffffff8211176101ca57604052565b6040810190811067ffffffffffffffff8211176101ca57604052565b90601f601f19910116810190811067ffffffffffffffff8211176101ca57604052565b6040519061023a6101c083610207565b565b6040519061023a61016083610207565b6040519061023a60a083610207565b6040519061023a60c083610207565b67ffffffffffffffff81116101ca57601f01601f191660200190565b60405190610295602083610207565b60008252565b60005b8381106102ae5750506000910152565b818101518382015260200161029e565b90601f19601f6020936102dc8151809281875287808801910161029b565b0116010190565b3461017a57600060031936011261017a5761034260408051906103068183610207565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102be565b0390f35b67ffffffffffffffff81160361017a57565b359061023a82610346565b908160a091031261017a5790565b3461017a57604060031936011261017a5760043561038e81610346565b60243567ffffffffffffffff811161017a576103ae903690600401610363565b6103cc8267ffffffffffffffff166000526006602052604060002090565b918254916001600160a01b036103f76103eb856001600160a01b031690565b6001600160a01b031690565b1615610562576040810191600161040e8484611f37565b90501161053857610342946104b19461049d6104566104306080870187611f6d565b61043d6020890189611f6d565b9050159182610522575b61045087612107565b88613164565b9561045f612221565b6104698288611f37565b90506104d1575b6040880161049381519260608b0193845161048d60028b01611fa0565b9161375a565b9092525285611f37565b151590506104c35760f01c90505b90613cd6565b60405190815292839250602083019150565b506001015461ffff166104ab565b5061051d6104f06104eb6104e5848a611f37565b9061229d565b6122ab565b60206104ff6104e5858b611f37565b013561051060208b015161ffff1690565b9060e08b015192896134e0565b610470565b915061052e8989611f37565b9050151591610447565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6000fd5b90602060031983011261017a5760043567ffffffffffffffff811161017a5760040160009280601f830112156105f95781359367ffffffffffffffff85116105f657506020808301928560051b01011161017a579190565b80fd5b8380fd5b3461017a5761060b3661059e565b90610614614645565b6000915b80831061062157005b61062c8382846122b5565b92610636846122f5565b67ffffffffffffffff81169081158015610ab5575b8015610a9f575b8015610a86575b610a4f57856108bb916108d56108cb836108c56106cb60e08301956106b16106ab610684898761232c565b6106a36106996101008a95949501809a61232c565b9490923691612362565b923691612362565b90614683565b67ffffffffffffffff166000526006602052604060002090565b9687956107116106dd60208a016122ab565b88906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6108b561087e60c060408b019a61077861072a8d61230a565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b6107d6610787608083016123c4565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b61081960016107e760a084016123c4565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6108788d610829606084016123ce565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b01612322565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d61232c565b90600389016124d9565b8a61232c565b90600286016124d9565b6101208801906108e76103eb836122ab565b15610a25576108f8610942926122ab565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b610140870190610966610960610958848b611f6d565b93905061230a565b60ff1690565b036109e157956109c77f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb26926109ac6109a2600198999a85611f6d565b90600484016125d2565b6109b585615d4b565b505460a01c67ffffffffffffffff1690565b6109d66040519283928361276c565b0390a2019190610618565b6109eb9087611f6d565b90610a216040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161257c565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a9860c08801612322565b1615610659565b5060ff610aae6040880161230a565b1615610652565b5067ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016821461064b565b6001600160a01b0381160361017a57565b3461017a57604060031936011261017a57610b15600435610346565b610b29602435610b2481610ae8565b61292b565b6040516001600160a01b039091168152602090f35b3461017a57610b4c3661059e565b906001600160a01b0360035416918215610c655760005b818110610b6c57005b610b7d6103eb6104eb8385876147b4565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610c60576001948892600092610c30575b5081610be4575b5050505001610b63565b81610c147f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610c2494615dcf565b6040519081529081906020820190565b0390a338858180610bda565b610c5291925060203d8111610c59575b610c4a8183610207565b8101906147c4565b9038610bd3565b503d610c40565b61291f565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b906020808351928381520192019060005b818110610cad5750505090565b82516001600160a01b0316845260209384019390920191600101610ca0565b80516001600160a01b03168252610d929160208281015167ffffffffffffffff169082015260408281015160ff169082015260608281015115159082015260808281015161ffff169082015260a08281015161ffff169082015260c08281015163ffffffff169082015260e0828101516001600160a01b031690820152610140610d80610d6c610100850151610160610100860152610160850190610c8f565b610120850151848203610120860152610c8f565b920151906101408184039101526102be565b90565b3461017a57602060031936011261017a5767ffffffffffffffff600435610dbb81610346565b610dc36129ca565b50166000526006602052610342610ddd6040600020612107565b604051918291602083526020830190610ccc565b61023a9092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461017a57600060031936011261017a57600060408051610e43816101ae565b8281528260208201520152610342604051610e5d816101ae565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610df1565b3461017a57600060031936011261017a576000546001600160a01b0381163303610f1d577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b6040810160408252825180915260206060830193019060005b818110610fc9575050506020818303910152815180825260208201916020808360051b8301019401926000915b838310610f9c57505050505090565b9091929394602080610fba83601f1986600196030187528951610ccc565b97019301930191939290610f8d565b825167ffffffffffffffff16855260209485019490920191600101610f60565b3461017a57600060031936011261017a5760045461100681612209565b906110146040519283610207565b808252601f1961102382612209565b0160005b8181106110cf5750506110398161223d565b9060005b81811061105557505061034260405192839283610f47565b8061108d611074611067600194615e8b565b67ffffffffffffffff1690565b61107e8387612a42565b9067ffffffffffffffff169052565b6110b36110ae6106b16110a08488612a42565b5167ffffffffffffffff1690565b612107565b6110bd8287612a42565b526110c88186612a42565b500161103d565b6020906110da6129ca565b82828701015201611027565b3461017a57600060031936011261017a5760206001600160a01b0360015416604051908152f35b359061023a82610ae8565b8015150361017a57565b359061023a82611118565b3461017a57606060031936011261017a57600060405161114c816101ae565b60043561115881610ae8565b815260243561116681611118565b602082019081526044359061117a82610ae8565b60408301918252611189614645565b6001600160a01b038351161580156112d4575b6112ac576001600160a01b03839261125961128f9361120d847f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9851166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006002541617600255565b5115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000060025492151560a01b16911617600255565b51166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006003541617600355565b611297611e50565b6112a6604051928392836147d3565b0390a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b508051151561119c565b3461017a57608060031936011261017a576112fa600435610346565b60243567ffffffffffffffff811161017a5761131a903690600401610363565b60443590611329606435610ae8565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610c6057600091611cb0575b50611c735760025460a01c60ff16611c495761140b740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61142b60043567ffffffffffffffff166000526006602052604060002090565b906001600160a01b036064351615611c1f5781546114526103eb6001600160a01b03831681565b3303611bf55781611467608086940182611f6d565b6114746020840184611f6d565b9050159081611bdc575b61148787612107565b6114949390600435613164565b9160a01c67ffffffffffffffff166114ab90612a6b565b84547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff000000000000000000000000000000000000000016178555825163ffffffff1694602084015161150e9061ffff1690565b6040805130602080830191909152815297919061152b9089610207565b604080516064356001600160a01b031660208083019190915281529890611552908a610207565b61155c8680611f6d565b85549a9160e08c901c60ff1691369061157492612a8a565b9060ff1661158191614883565b60a08901519161159460408a018a611f37565b61159e9150612b03565b936115ac60208b018b611f6d565b9690976115b761022a565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529a67ffffffffffffffff6004351660208d015267ffffffffffffffff1660408c0152600060608c015263ffffffff1660808b015261ffff1660a08a0152600060c08a015260e089015261163960048801612047565b610100890152610120880152610140870152610160860152610180850152369061166292612a8a565b6101a0830152611670612221565b61167d6040850185611f37565b9050611b62575b836116fd926116cc6116a788946040860151606087015161048d60028701611fa0565b60608601528060408601526116c660808601516001600160a01b031690565b9061490d565b60c08601526116d9612b52565b986116e76040840184611f37565b15159050611b545760f01c90505b600435613cd6565b63ffffffff9091166060840152602087019591865211611b2a5761172284518361499b565b61172f6040830183611f37565b9050611a0a575b61174281959295615375565b808352602081519101209061175b604085015151612bd6565b6040840190815294606087019360005b6040870151805182101561194057602061179e6103eb6103eb611791866117e596612a42565b516001600160a01b031690565b6117ac8460608c0151612a42565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612c1f565b03915afa8015610c60576001600160a01b0391600091611912575b501680156118b957906000888c9388838961186361182c888f6118246060916122ab565b980151612a42565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612d6a565b03925af18015610c60578161189191600194600091611898575b508b519061188b8383612a42565b52612a42565b500161176b565b6118b3913d8091833e6118ab8183610207565b810190612c82565b3861187d565b61059a6118cd6117918460408c0151612a42565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611933915060203d8111611939575b61192b8183610207565b81019061290a565b38611800565b503d611921565b6103428680858d7f371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc8e8a6119738f6122ab565b936119816040820182611f37565b1590506119fe57602061199e6104e58360406119ce950190611f37565b0135955b51915192516040519384936001600160a01b03606435169867ffffffffffffffff600435169886612f35565b0390a4610c147fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b506119ce6000956119a2565b611a6c611a56611a206104e56040860186611f37565b60c0860180515115611b0f5751905b602087015161ffff169060e08801519260643591611a51600435913690612b71565b614c68565b61018083015190611a6682612a35565b52612a35565b50611a8e6040611a828651828701515190612a42565b51015163ffffffff1690565b60a0611a9e610180840151612a35565b5101515163ffffffff82168111611ab6575050611736565b61059a9250611ace6104eb6104e56040870187611f37565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b50611b24611b1d8680611f6d565b3691612a8a565b90611a2f565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff166116f5565b5093506001611b746040840184611f37565b905003610538576116fd838388966116cc6116a7611bd0611b9e6104eb6104e56040880188611f37565b6020611bb06104e56040890189611f37565b0135611bc1602089015161ffff1690565b9060e0890151926004356134e0565b94505050925050611684565b9050611beb6040840184611f37565b905015159061147e565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005261059a6004359067ffffffffffffffff60249216600452565b611cd2915060203d602011611cd8575b611cca8183610207565b810190612a56565b386113b5565b503d611cc0565b3461017a57602060031936011261017a5767ffffffffffffffff600435611d0581610346565b166000526006602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611d495760405167ffffffffffffffff9091168152602090f35b6123d8565b3461017a57602060031936011261017a576001600160a01b03600435611d7381610ae8565b611d7b614645565b16338114611de057807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461017a57602060031936011261017a57611e26600435610346565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611e60816101cf565b8281528260208201528260408201520152604051611e7d816101cf565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a57602001918160061b3603831361017a57565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a5760200191813603831361017a57565b906040519182815491828252602082019060005260206000209260005b818110611fd257505061023a92500383610207565b84546001600160a01b0316835260019485019487945060209093019201611fbd565b90600182811c9216801561203d575b602083101461200e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691612003565b906040519182600082549261205b84611ff4565b80845293600181169081156120c75750600114612080575b5061023a92500383610207565b90506000929192526020600020906000915b8183106120ab57505090602061023a9282010138612073565b6020919350806001915483858901015201910190918492612092565b6020935061023a9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138612073565b90612201600461211561023c565b93612184612179825461213e612131826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c16604089015261217360e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b6121d76121c760018301546121a861219d8261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b6121e360028201611fa0565b6101008601526121f560038201611fa0565b61012086015201612047565b610140830152565b67ffffffffffffffff81116101ca5760051b60200190565b60405190612230602083610207565b6000808352366020840137565b9061224782612209565b6122546040519182610207565b828152601f196122648294612209565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90156122a65790565b61226e565b35610d9281610ae8565b91908110156122a65760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561017a570190565b35610d9281610346565b60ff81160361017a57565b35610d92816122ff565b63ffffffff81160361017a57565b35610d9281612314565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a57602001918160051b3603831361017a57565b92919061236e81612209565b9361237c6040519586610207565b602085838152019160051b810192831161017a57905b82821061239e57505050565b6020809183356123ad81610ae8565b815201910190612392565b61ffff81160361017a57565b35610d92816123b8565b35610d9281611118565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611d4957565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611d4957565b908160031b9180830460081490151715611d4957565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611d4957565b81810292918115918404141715611d4957565b8181106124cd575050565b600081556001016124c2565b9067ffffffffffffffff83116101ca576801000000000000000083116101ca57815483835580841061253d575b5090600052602060002060005b8381106125205750505050565b600190602084359461253186610ae8565b01938184015501612513565b612555908360005284602060002091820191016124c2565b38612506565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610d9293818152019161255b565b9190601f811161259c57505050565b61023a926000526020600020906020601f840160051c830193106125c8575b601f0160051c01906124c2565b90915081906125bb565b90929167ffffffffffffffff81116101ca576125f8816125f28454611ff4565b8461258d565b6000601f821160011461263857819061262993949560009261262d575b50506000198260011b9260031b1c19161790565b9055565b013590503880612615565b601f1982169461264d84600052602060002090565b91805b87811061268857508360019596971061266e575b505050811b019055565b60001960f88560031b161c19910135169055388080612664565b90926020600181928686013581550194019101612650565b359061023a826122ff565b359061023a826123b8565b359061023a82612314565b9035601e198236030181121561017a57016020813591019167ffffffffffffffff821161017a578160051b3603831361017a57565b9160209082815201919060005b8181106127105750505090565b9091926020806001926001600160a01b03873561272c81610ae8565b168152019401929101612703565b9035601e198236030181121561017a57016020813591019167ffffffffffffffff821161017a57813603831361017a57565b67ffffffffffffffff610d929392168152604060208201526127a26040820161279484610358565b67ffffffffffffffff169052565b6127c16127b16020840161110d565b6001600160a01b03166060830152565b6127da6127d0604084016126a0565b60ff166080830152565b6127f26127e960608401611122565b151560a0830152565b61280c612801608084016126ab565b61ffff1660c0830152565b61282661281b60a084016126ab565b61ffff1660e0830152565b61284361283560c084016126b6565b63ffffffff16610100830152565b6128d96128ac61286d61285960e08601866126c1565b6101606101208701526101a08601916126f6565b61287b6101008601866126c1565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0868403016101408701526126f6565b926128ce6128bd610120830161110d565b6001600160a01b0316610160850152565b61014081019061273a565b916101807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08286030191015261255b565b9081602091031261017a5751610d9281610ae8565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610c60576001600160a01b03916000916129ad57501690565b6129c6915060203d6020116119395761192b8183610207565b1690565b60405190610160820182811067ffffffffffffffff8211176101ca5760405260606101408360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e082015282610100820152826101208201520152565b8051156122a65760200190565b80518210156122a65760209160051b010190565b9081602091031261017a5751610d9281611118565b67ffffffffffffffff1667ffffffffffffffff8114611d495760010190565b929192612a968261026a565b91612aa46040519384610207565b82948184528183011161017a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101ca57604052606060a0836000815282602082015282604082015282808201528260808201520152565b90612b0d82612209565b612b1a6040519182610207565b828152601f19612b2a8294612209565b019060005b828110612b3b57505050565b602090612b46612ac1565b82828501015201612b2f565b60405190612b5f826101ae565b60606040838281528260208201520152565b919082604091031261017a57604051612b89816101eb565b60208082948035612b9981610ae8565b84520135910152565b60405190612bb1602083610207565b600080835282815b828110612bc557505050565b806060602080938501015201612bb9565b90612be082612209565b612bed6040519182610207565b828152601f19612bfd8294612209565b019060005b828110612c0e57505050565b806060602080938501015201612c02565b60409067ffffffffffffffff610d92949316815281602082015201906102be565b81601f8201121561017a578051612c568161026a565b92612c646040519485610207565b8184526020828401011161017a57610d92916020808501910161029b565b9060208282031261017a57815167ffffffffffffffff811161017a57610d929201612c40565b9080602083519182815201916020808360051b8301019401926000915b838310612cd457505050505090565b9091929394602080612d5b83601f1986600196030187528951908151815260a0612d4a612d38612d26612d148887015160c08a88015260c08701906102be565b604087015186820360408801526102be565b606086015185820360608701526102be565b608085015184820360808601526102be565b9201519060a08184039101526102be565b97019301930191939290612cc5565b919390610d929593612eb2612eca9260a08652612d9460a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612e9d612e85612e6d612e55612e3d612e278c61026060e08a0151916101c061018082015201906102be565b6101008801518d8203609f1901888f01526102be565b6101208701518c8203609f19016101c08e01526102be565b610140860151609f198c8303016101e08d01526102be565b610160850151609f198b8303016102008c01526102be565b610180840151609f198a8303016102208b0152612ca8565b910151609f19878303016102408801526102be565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102be565b9080602083519182815201916020808360051b8301019401926000915b838310612f0857505050505090565b9091929394602080612f2683601f19866001960301875289516102be565b97019301930191939290612ef9565b95949290916001600160a01b03612f5f93168752602087015260a0604087015260a08601906102be565b938085036060820152825180865260208601906020808260051b8901019501916000905b828210612fa15750505050610d929394506080818403910152612edc565b9091929560208061300283601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102be565b980192019201909291612f83565b60405190610100820182811067ffffffffffffffff8211176101ca57604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161017a5790600490565b9093929384831161017a57841161017a578101920390565b919091357fffffffff00000000000000000000000000000000000000000000000000000000811692600481106130bb575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261017a57825167ffffffffffffffff811161017a5782613115918501612c40565b92602081015161312481612314565b92604082015167ffffffffffffffff811161017a57610d929201612c40565b60409067ffffffffffffffff610d929593168152816020820152019161255b565b91929092613170613010565b600483101580613339575b1561328357509061318b91615632565b926131996040850151615885565b80613268575b604084016131bc8151926060870193845161012088015191615929565b9092525260c083015151613218575b50608082016001600160a01b036131e982516001600160a01b031690565b16156131f457505090565b61320b60e0610d929301516001600160a01b031690565b6001600160a01b03169052565b61322c6132286060840151151590565b1590565b156131cb577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff61327c845163ffffffff1690565b161561319f565b94916000906132db926132a46103eb6103eb6002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501613143565b03915afa8015610c605760009060009060009061330d575b60a088015263ffffffff16865290505b60c0850152613199565b50505061332f613303913d806000833e6133278183610207565b8101906130ed565b91925082916132f3565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061338f6133898686613061565b90613087565b161461317b565b60208183031261017a5780519067ffffffffffffffff821161017a57019080601f8301121561017a5781516133ca81612209565b926133d86040519485610207565b81845260208085019260051b82010192831161017a57602001905b8282106134005750505090565b60208091835161340f81610ae8565b8152019101906133f3565b95949060009460a09467ffffffffffffffff613461956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102be565b930152565b9060028201809211611d4957565b9060018201809211611d4957565b6001019081600111611d4957565b9060148201809211611d4957565b90600c8201809211611d4957565b91908201809211611d4957565b6000198114611d495760010190565b80548210156122a65760005260206000200190600090565b929394919060036135058567ffffffffffffffff166000526006602052604060002090565b01936001600160a01b0361351a81841661292b565b16918215613724576040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c6057600091613705575b50156136f5576135cb600095969798604051998a96879586957f89720a620000000000000000000000000000000000000000000000000000000087526004870161341a565b03915afa928315610c60576000936136d0575b508251156136c5578251906135fd6135f8845480946134ac565b61223d565b906000928394845b87518110156136645761361b611791828a612a42565b6001600160a01b03811615613658579061365260019261364461363d8a6134b9565b9989612a42565b906001600160a01b03169052565b01613605565b50955060018096613652565b509195509193613676575b5050815290565b60005b828110613686575061366f565b806136bf6136ac613699600194866134c8565b90546001600160a01b039160031b1c1690565b6136446136b8886134b9565b9789612a42565b01613679565b9150610d9290611fa0565b6136ee9193503d806000833e6136e68183610207565b810190613396565b91386135de565b505050509250610d929150611fa0565b61371e915060203d602011611cd857611cca8183610207565b38613586565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b9391929361377661376e82518651906134ac565b8651906134ac565b906137896137838361223d565b92612bd6565b94600096875b83518910156137ef57886137e56137d86001936137c06137b66117918e9f9d9e9d8b612a42565b613644838c612a42565b6137de6137cd858c612a42565b5191809384916134b9565b9c612a42565b528b612a42565b500197969561378f565b959250929350955060005b8651811015613883576138106117918289612a42565b60006001600160a01b038216815b888110613857575b505090600192911561383a575b50016137fa565b6138519061364461384a896134b9565b9888612a42565b38613833565b816138686103eb611791848c612a42565b146138755760010161381e565b506001915081905038613826565b509390945060005b8551811015613914576138a16117918288612a42565b60006001600160a01b038216815b8781106138e8575b50509060019291156138cb575b500161388b565b6138e2906136446138db886134b9565b9787612a42565b386138c4565b816138f96103eb611791848b612a42565b14613906576001016138af565b5060019150819050386138b7565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101ca5760405260606080836000815260006020820152600060408201526000838201520152565b9061396982612209565b6139766040519182610207565b828152601f196139868294612209565b019060005b82811061399757505050565b6020906139a2613920565b8282850101520161398b565b9081606091031261017a5780516139c4816123b8565b91604060208301516139d581612314565b920151610d9281612314565b9160209082815201919060005b8181106139fb5750505090565b9091926040806001926001600160a01b038735613a1781610ae8565b168152602087810135908201520194019291016139ee565b949391929067ffffffffffffffff16855260806020860152613a88613a69613a57858061273a565b60a060808a015261012089019161255b565b613a76602086018661273a565b90607f198984030160a08a015261255b565b6040840135601e198536030181121561017a578401916020833593019167ffffffffffffffff841161017a578360061b3603831361017a5761023a95613b10613ae783606097613b31978d60c0607f1982613b239a03019101526139e1565b91613b06613af688830161110d565b6001600160a01b031660e08d0152565b608081019061273a565b90607f198b8403016101008c015261255b565b9087820360408901526102be565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611d4957565b908160a091031261017a578051916020820151613b7281612314565b916040810151613b8181612314565b9160806060830151613b92816123b8565b920151610d9281611118565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610d929b9a9616885216602087015260408601521660608401521660808201528160a082015201906102be565b9081606091031261017a5780516139c481612314565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611d4957565b906000198201918211611d4957565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611d4957565b91908203918211611d4957565b919082608091031261017a578151613c8981612314565b916020810151916060604083015192015190565b8115613ca7570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9382946000906000956040810194613d10613d0b613d06885151613cfe60408c01809c611f37565b9190506134ac565b613466565b61395f565b9660009586955b88518051881015613f74576103eb6103eb6117918a613d3594612a42565b613d8260206060880192613d4a8b8551612a42565b51908a6040518095819482937f958021a700000000000000000000000000000000000000000000000000000000845260048401612c1f565b03915afa8015610c60576001600160a01b0391600091613f56575b50168015613f02579060608e9392613db68b8451612a42565b5190613dc760208b015161ffff1690565b958b613e02604051988995869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613a2f565b03915afa8015610c6057600193613e9f938b8f8f95600080958197613ea8575b509083929161ffff613e4a85613e43611791613e9399613e999d9e51612a42565b9451612a42565b5191613e66613e5761024c565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261188b8383612a42565b50613b3c565b99613b3c565b96019596613d17565b613e999750611791965084939291509361ffff613e4a82613e43613ee5613e939960603d8111613efb575b613edd8183610207565b8101906139ae565b9c9196909c9d5050505050505090919293613e22565b503d613ed3565b61059a88613f146117918c8f51612a42565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613f6e915060203d81116119395761192b8183610207565b38613d9d565b50919a9496929395509897968a613f8b8187611f37565b905061429f575b50508651613f9f90613bfc565b99613fad6020860186611f6d565b91613fb9915086611f37565b9560609150019486613fca876122ab565b91613fd5938a615b42565b613fdf8b89612a42565b52613fea8a88612a42565b50613ff58a88612a42565b516020015163ffffffff1661400991613b3c565b906140148a88612a42565b516040015163ffffffff1661402891613b3c565b9161403161024c565b33815290600060208301819052604083015261ffff166060820152614054610286565b6080820152865161406490613c29565b9061406f8289612a42565b5261407a9087612a42565b506002546001600160a01b031692614091906122ab565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610c6057600096600097600094600092614262575b506000965b86518810156141e25761416b6001916141338861412e87612407565b613c9d565b61414c60606141428d8d612a42565b51019182516124af565b9052858a14614173575b60606141628b8b612a42565b510151906134ac565b970196614112565b8b8873eba517d2000000000000000000000000000000006001600160a01b036141a660808c01516001600160a01b031690565b16036141b4575b5050614156565b61412e6141c092612437565b6141d960606141cf8d8d612a42565b51019182516134ac565b90528b886141ad565b979650975050505061421e7f00000000000000000000000000000000000000000000000000000000000000009161412e63ffffffff8416612437565b841161422a5750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b929850505061428a91925060803d608011614298575b6142828183610207565b810190613c72565b91979093929091903861410d565b503d614278565b610b246103eb6104eb6104e56142b8948a989698611f37565b926001600160a01b03600091515194169060e088019081516142d861024c565b6001600160a01b0385168152908260208301528260408301528260608301526080820152614306878d612a42565b52614311868c612a42565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3317103100000000000000000000000000000000000000000000000000000000600482015291602083602481875afa8015610c60578f948c89968f96948d948f968891614626575b5061451f575b505050505050156143c6575b611a826143b76143be956143b16020611a826143b197604097612a42565b90613b3c565b958b612a42565b90388a613f92565b505061444c9160608c6143f76104eb6104e56143f06103eb6103eb6002546001600160a01b031690565b938b611f37565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610c6057611a826143b76040926143b16020611a828f8b906143be9b6143b19a6000806000926144e1575b63ffffffff9293506144ce9060606144968888612a42565b5101946144c38a6144a78a8a612a42565b51019160406144b68b8b612a42565b51019063ffffffff169052565b9063ffffffff169052565b1690529750975050505095505050614393565b50505063ffffffff61450d6144ce9260603d606011614518575b6145058183610207565b810190613be6565b90935091508261447e565b503d6144fb565b8495985060a0969750614569602061455f6060826145566104e561454f6104eb6104e58b6145a29c9d9e9f611f37565b998d611f37565b013599016122ab565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613b9e565b03915afa8015610c60578592828c9391819082946145ea575b506145de9060606145cc8888612a42565b5101926144c360206144a78a8a612a42565b5288888f8c8138614387565b9150506145de9250614614915060a03d60a01161461f575b61460c8183610207565b810190613b56565b9491929190506145bb565b503d614602565b61463f915060203d602011611cd857611cca8183610207565b38614381565b6001600160a01b0360015416330361465957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916146918151846134ac565b92831561478a5760005b8481106146a9575050505050565b8181101561476f576146be6117918286612a42565b6001600160a01b0381168015610c65576146d783613474565b8781106146e95750505060010161469b565b8481101561474c576001600160a01b03614706611791838a612a42565b168214614715576001016146d7565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b0361476a6117916147648885613c65565b89612a42565b614706565b61478561179161477f8484613c65565b85612a42565b6146be565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156122a65760051b0190565b9081602091031261017a575190565b91608061023a9294936148238160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610d929281815201906102be565b906148668261026a565b6148736040519182610207565b828152601f19612264829461026a565b9081516020821080614903575b6148d2570361489c5790565b610a21906040519182917f3aeba3900000000000000000000000000000000000000000000000000000000083526004830161484b565b5090602081015190816148e484612469565b1c61489c57506148f38261485c565b9160200360031b1b602082015290565b5060208114614890565b918251601481029080820460141490151715611d495761492f61493491613482565b613490565b906149466149418361349e565b61485c565b90601461495283612a35565b5360009260215b8651851015614984576014600191614974611791888b612a42565b60601b8187015201940193614959565b919550936020935060601b90820152828152012090565b906149ab6103eb606084016122ab565b6149bc600019936040810190611f37565b9050614a34575b6149cd8251613c29565b9260005b8481106149df575050505050565b808260019214614a2f5760606149f58287612a42565b5101518015614a2957614a2390614a1d614a0f8489612a42565b51516001600160a01b031690565b86615dcf565b016149d1565b50614a23565b614a23565b9150614a408151613c38565b91614a4e614a0f8484612a42565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f331710310000000000000000000000000000000000000000000000000000000060048201526020816024816001600160a01b0386165afa908115610c6057600091614ae7575b50614ac7575b506149c3565b614ae1906060614ad78686612a42565b5101519083615dcf565b38614ac1565b614b00915060203d602011611cd857611cca8183610207565b38614abb565b60405190614b13826101eb565b60606020838281520152565b919060408382031261017a5760405190614b38826101eb565b8193805167ffffffffffffffff811161017a5782614b57918301612c40565b835260208101519167ffffffffffffffff831161017a57602092614b7b9201612c40565b910152565b9060208282031261017a57815167ffffffffffffffff811161017a57610d929201614b1f565b9060806001600160a01b0381614bc5855160a0865260a08601906102be565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610d92928181520190614ba6565b919060408382031261017a57825167ffffffffffffffff811161017a57602091614c35918501614b1f565b92015190565b61ffff614c54610d929593606084526060840190614ba6565b9316602082015260408184039101526102be565b909193929594614c76612ac1565b506020820180511561511057614c99610b246103eb85516001600160a01b031690565b946001600160a01b038616916040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180614d0060048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610c60576000916150f1575b50156150a857614d8c88999a825192614d2b614b06565b5051614d77614d4189516001600160a01b031690565b926040614d4c61024c565b9e8f908152614d688d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c6057600091615089575b5015614f915750906000929183614e389899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614c3b565b03925af18015610c6057600094600091614f49575b50614f1e614e879596614ec1614e95614eb3956117916020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b610207565b6040519586918683019190916001600160a01b036020820193169052565b03601f198101865285610207565b614efd610960614ef3614f038951614efd610960614ef38c67ffffffffffffffff166000526006602052604060002090565b5460e01c60ff1690565b90614883565b9767ffffffffffffffff166000526006602052604060002090565b93015193614f2a61025b565b958652602086015260408501526060840152608083015260a082015290565b614e8795506020915061179196614ec1614e95614eb395614f7f614f1e953d806000833e614f778183610207565b810190614c0a565b9b909b96505095505050969550614e4d565b9793929061ffff1661505f575161503557614fe06000939184926040519586809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614bf9565b03925af1908115610c6057614e8795614ec1614e95614f1e93611791602096614eb398600091615012575b5099614e69565b61502f91503d806000833e6150278183610207565b810190614b80565b3861500b565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b6150a2915060203d602011611cd857611cca8183610207565b38614df0565b61059a6150bc86516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b61510a915060203d602011611cd857611cca8183610207565b38614d14565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b959261520c947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b906152286020928281519485920161029b565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b168152615274825180936020898501910161029b565b019160f81b168382015261529282518093602060028501910161029b565b01019160f81b16838201526152b182518093602060028501910161029b565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff00000000000000000000000000000000000000000000000000000000000000610d929f9e9c9860f81b168152615326825180936020898501910161029b565b019160f01b168382015261534482518093602060038501910161029b565b01019160f01b1683820152615362825180936020898501910161029b565b01019160f01b1660028201520190615215565b60e081019060ff8251511161561c57610100810160ff815151116156065761012082019260ff845151116155f057610140830160ff815151116155da5761016084019061ffff825151116155c4576101808501936001855151116155ac576101a086019361ffff855151116155965760609551805161555a575b50865167ffffffffffffffff16602088015167ffffffffffffffff169060408901516154229067ffffffffffffffff1690565b9860608101516154359063ffffffff1690565b9060808101516154489063ffffffff1690565b60a082015161ffff169160c00151926040519c8d96602088019661546b9761513a565b03601f198101885261547d9088610207565b5190815161548b9060ff1690565b9051805160ff1698519081516154a19060ff1690565b906040519a8b9560208701956154b69661522c565b03601f19810187526154c89087610207565b519182516154d69060ff1690565b9151805161ffff169480516154ec9061ffff1690565b92519283516154fc9061ffff1690565b926040519788976020890197615511986152b7565b03601f19810182526155239082610207565b6040519283926020840161553691615215565b61553f91615215565b61554891615215565b03601f1981018252610d929082610207565b61556f91965061556990612a35565b51615ffb565b9461ffff86511161558057386153ef565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b600052602660045260246000fd5b635a102da160e11b60005261059a6024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b9061563b613010565b91601182106158515780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216036157de5750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a6156ba8161223d565b604086019081526156ca82612bd6565b906060870191825260005b838110615792575050505061573d838361573361572761571d61571661570361574798876157519c9b616122565b6001600160a01b0390911660808d015290565b85856161f8565b9291903691612a8a565b60a08a01528383616260565b9491903691612a8a565b60c08801526161f8565b9391903691612a8a565b60e08401528103615760575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b806001916157d76157c16157ba6157ad6157d19a8d8d616122565b9190613644868a51612a42565b8b8b6161f8565b9391889a919a51949a3691612a8a565b92612a42565b52016156d5565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b80519060005b82811061589757505050565b60018101808211611d49575b8381106158b3575060010161588b565b6001600160a01b036158c58385612a42565b51166158d76103eb6117918487612a42565b146158e4576001016158a3565b61059a6158f46117918486612a42565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b90939192815115615ab5575061593e81615885565b80519260005b84811061595357509093925050565b6159636103eb6117918386612a42565b1561597057600101615944565b93919461598a6135f861598285613c29565b8451906134ac565b926159a76159a261599a83613c29565b8551906134ac565b612bd6565b968792600097885b848110615a575750505050505060005b8151811015615a4a576000805b868110615a0b575b5090600191615a0657615a006159ed6117918386612a42565b6136446159f9896134b9565b9887612a42565b016159bf565b615a00565b615a186117918287612a42565b6001600160a01b03615a306103eb6117918789612a42565b911614615a3f576001016159cc565b5060019050806159d4565b5050909180825283529190565b909192939498828214615aab5790615a9d615a9083615a838b6136446001976157d1611791898e612a42565b615a966137cd8589612a42565b9e612a42565b528c612a42565b505b019089949392916159af565b9850600190615a9f565b9193505015615ad05750615ac7612221565b90610d92612ba2565b90610d928251612bd6565b9081602091031261017a5751610d92816123b8565b93615b2d60809461ffff6001600160a01b039567ffffffffffffffff615b3b969b9a9b16895216602088015260a0604088015260a0870190610c8f565b9085820360608701526102be565b9416910152565b92919092615b4e613920565b50615b6d8167ffffffffffffffff166000526006602052604060002090565b805490959060e01c60ff169160808501928351615b90906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff16615bb091613b3c565b96615bbc90608d6134ac565b9460a0870195865151615bce916134ac565b9160ff1691615bdc8361247f565b615be5916134ac565b91615bf19060676134ac565b615bfa916124af565b615c03916134ac565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b03871603615c905750505061ffff9250615c8290615c756000935b5195615c68615c5961024c565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b615ca66103eb602094976001600160a01b031690565b906040615cb78583015161ffff1690565b91015192615cf7875198604051998a96879586957ff238895800000000000000000000000000000000000000000000000000000000875260048701615af0565b03915afa908115610c6057615c75615c829261ffff95600091615d1c575b5093615c4c565b615d3e915060203d602011615d44575b615d368183610207565b810190615adb565b38615d15565b503d615d2c565b80600052600560205260406000205415600014615dc957600454680100000000000000008110156101ca576001810160045560006004548210156122a657600490527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01819055600454906000526005602052604060002055600190565b50600090565b91602091600091604051906001600160a01b03858301937fa9059cbb000000000000000000000000000000000000000000000000000000008552166024830152604482015260448152615e23606482610207565b519082855af11561291f576000513d615e8257506001600160a01b0381163b155b615e4b5750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615e44565b6004548110156122a65760046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b015490565b966002615fcf97615f9c6022610d929f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090615f9c9f82615f9c9c615fa39c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615f4d825180936020898501910161029b565b019160f81b1683820152615f6b82518093602060238501910161029b565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190615215565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff8251511161610c57604081019160ff835151116160f657606082019160ff835151116160e057608081019260ff845151116160ca5760a0820161ffff815151116160b457610d92946160a6935194519161605d835160ff1690565b97519161606b835160ff1690565b945190616079825160ff1690565b905193616087855160ff1690565b935196616096885161ffff1690565b966040519c8d9b60208d01615ec0565b03601f198101835282610207565b635a102da160e11b600052602b60045260246000fd5b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b9291909260018201918483116161c65781013560001a8281156161bb57506014810361618e57820193841161615a57013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b919060028201918183116161c6578381013560f01c016002019281841161622c579183916162259361306f565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b919060018201918183116161c6578381013560001a016001019281841161622c579183916162259361306f56fea164736f6c634300081a000a",
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
