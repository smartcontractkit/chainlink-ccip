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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllDestChainConfigs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfig[]\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100604052346103675760405161660a38819003601f8101601f191683016001600160401b0381118482101761036c578392829160405283398101039060e08212610367576080821261036757610055610382565b81519092906001600160401b03811681036103675783526020820151926001600160a01b03841684036103675760208101938452604083015163ffffffff81168103610367576040820190815260606100af8186016103a1565b83820190815293607f1901126103675760405192606084016001600160401b0381118582101761036c576040526100e8608086016103a1565b845260a08501519485151586036103675760c061010c9160208701978852016103a1565b9560408501968752331561035657600180546001600160a01b0319163317905583516001600160401b0316158015610344575b8015610332575b8015610323575b6103085792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e052815116158015610319575b6103085780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e09260006060610206610382565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff92831691166060610251610382565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a160405161625490816103b68239608051818181610ac1015281816115c60152611e49015260a0518181816113840152611e75015260c051818181611ed00152612931015260e051818181611ea101526141a90152f35b6306b7c75960e31b60005260046000fd5b508151151561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761036c57604052565b51906001600160a01b03821682036103675756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610117578063181f5a771461011257806320487ded1461010d5780632490769e1461010857806348a98aa4146101035780635cb80c5d146100fe5780636def4ce7146100f95780637437ff9f146100f457806379ba5097146100ef57806389933a51146100ea5780638da5cb5b146100e557806390423fa2146100e0578063df0aa9e9146100db578063e8d80861146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611dcb565b611d0f565b611ca0565b6112de565b61112d565b6110e6565b610fe9565b610e96565b610e23565b610d95565b610b3e565b610af9565b6105fd565b610371565b6102e3565b3461017a57600060031936011261017a576080610132611e11565b61017860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176101ca57604052565b61017f565b6080810190811067ffffffffffffffff8211176101ca57604052565b6040810190811067ffffffffffffffff8211176101ca57604052565b90601f601f19910116810190811067ffffffffffffffff8211176101ca57604052565b6040519061023a6101c083610207565b565b6040519061023a61016083610207565b6040519061023a60a083610207565b6040519061023a60c083610207565b67ffffffffffffffff81116101ca57601f01601f191660200190565b60405190610295602083610207565b60008252565b60005b8381106102ae5750506000910152565b818101518382015260200161029e565b90601f19601f6020936102dc8151809281875287808801910161029b565b0116010190565b3461017a57600060031936011261017a5761034260408051906103068183610207565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102be565b0390f35b67ffffffffffffffff81160361017a57565b359061023a82610346565b908160a091031261017a5790565b3461017a57604060031936011261017a5760043561038e81610346565b60243567ffffffffffffffff811161017a576103ae903690600401610363565b6103cc8267ffffffffffffffff166000526006602052604060002090565b918254916001600160a01b036103f76103eb856001600160a01b031690565b6001600160a01b031690565b1615610562576040810191600161040e8484611ef8565b90501161053857610342946104b19461049d6104566104306080870187611f2e565b61043d6020890189611f2e565b9050159182610522575b610450876120c8565b8861311e565b9561045f6121e2565b6104698288611ef8565b90506104d1575b6040880161049381519260608b0193845161048d60028b01611f61565b91613714565b9092525285611ef8565b151590506104c35760f01c90505b90613c90565b60405190815292839250602083019150565b506001015461ffff166104ab565b5061051d6104f06104eb6104e5848a611ef8565b9061225e565b61226c565b60206104ff6104e5858b611ef8565b013561051060208b015161ffff1690565b9060e08b0151928961349a565b610470565b915061052e8989611ef8565b9050151591610447565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6000fd5b90602060031983011261017a5760043567ffffffffffffffff811161017a5760040160009280601f830112156105f95781359367ffffffffffffffff85116105f657506020808301928560051b01011161017a579190565b80fd5b8380fd5b3461017a5761060b3661059e565b906106146145ff565b6000915b80831061062157005b61062c838284612276565b92610636846122b6565b67ffffffffffffffff81169081158015610ab5575b8015610a9f575b8015610a86575b610a4f57856108bb916108d56108cb836108c56106cb60e08301956106b16106ab61068489876122ed565b6106a36106996101008a95949501809a6122ed565b9490923691612323565b923691612323565b9061463d565b67ffffffffffffffff166000526006602052604060002090565b9687956107116106dd60208a0161226c565b88906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6108b561087e60c060408b019a61077861072a8d6122cb565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b6107d661078760808301612385565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b61081960016107e760a08401612385565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6108788d6108296060840161238f565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b016122e3565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d6122ed565b906003890161249a565b8a6122ed565b906002860161249a565b6101208801906108e76103eb8361226c565b15610a25576108f86109429261226c565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b610140870190610966610960610958848b611f2e565b9390506122cb565b60ff1690565b036109e157956109c77f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb26926109ac6109a2600198999a85611f2e565b9060048401612593565b6109b585615d05565b505460a01c67ffffffffffffffff1690565b6109d66040519283928361272d565b0390a2019190610618565b6109eb9087611f2e565b90610a216040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161253d565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a9860c088016122e3565b1615610659565b5060ff610aae604088016122cb565b1615610652565b5067ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016821461064b565b6001600160a01b0381160361017a57565b3461017a57604060031936011261017a57610b15600435610346565b610b29602435610b2481610ae8565b6128ec565b6040516001600160a01b039091168152602090f35b3461017a57610b4c3661059e565b906001600160a01b0360035416918215610c655760005b818110610b6c57005b610b7d6103eb6104eb83858761476e565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610c60576001948892600092610c30575b5081610be4575b5050505001610b63565b81610c147f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610c2494615d89565b6040519081529081906020820190565b0390a338858180610bda565b610c5291925060203d8111610c59575b610c4a8183610207565b81019061477e565b9038610bd3565b503d610c40565b6128e0565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b906020808351928381520192019060005b818110610cad5750505090565b82516001600160a01b0316845260209384019390920191600101610ca0565b80516001600160a01b03168252610d929160208281015167ffffffffffffffff169082015260408281015160ff169082015260608281015115159082015260808281015161ffff169082015260a08281015161ffff169082015260c08281015163ffffffff169082015260e0828101516001600160a01b031690820152610140610d80610d6c610100850151610160610100860152610160850190610c8f565b610120850151848203610120860152610c8f565b920151906101408184039101526102be565b90565b3461017a57602060031936011261017a5767ffffffffffffffff600435610dbb81610346565b610dc361298b565b50166000526006602052610342610ddd60406000206120c8565b604051918291602083526020830190610ccc565b61023a9092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461017a57600060031936011261017a57600060408051610e43816101ae565b8281528260208201520152610342604051610e5d816101ae565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610df1565b3461017a57600060031936011261017a576000546001600160a01b0381163303610f1d577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b6040810160408252825180915260206060830193019060005b818110610fc9575050506020818303910152815180825260208201916020808360051b8301019401926000915b838310610f9c57505050505090565b9091929394602080610fba83601f1986600196030187528951610ccc565b97019301930191939290610f8d565b825167ffffffffffffffff16855260209485019490920191600101610f60565b3461017a57600060031936011261017a57600454611006816121ca565b906110146040519283610207565b808252601f19611023826121ca565b0160005b8181106110cf575050611039816121fe565b9060005b81811061105557505061034260405192839283610f47565b8061108d611074611067600194615e45565b67ffffffffffffffff1690565b61107e8387612a03565b9067ffffffffffffffff169052565b6110b36110ae6106b16110a08488612a03565b5167ffffffffffffffff1690565b6120c8565b6110bd8287612a03565b526110c88186612a03565b500161103d565b6020906110da61298b565b82828701015201611027565b3461017a57600060031936011261017a5760206001600160a01b0360015416604051908152f35b359061023a82610ae8565b8015150361017a57565b359061023a82611118565b3461017a57606060031936011261017a57600060405161114c816101ae565b60043561115881610ae8565b815260243561116681611118565b602082019081526044359061117a82610ae8565b604083019182526111896145ff565b6001600160a01b038351161580156112d4575b6112ac576001600160a01b03839261125961128f9361120d847f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9851166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006002541617600255565b5115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000060025492151560a01b16911617600255565b51166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006003541617600355565b611297611e11565b6112a66040519283928361478d565b0390a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b508051151561119c565b3461017a57608060031936011261017a576112fa600435610346565b60243567ffffffffffffffff811161017a5761131a903690600401610363565b60443590611329606435610ae8565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610c6057600091611c71575b50611c345760025460a01c60ff16611c0a5761140b740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61142b60043567ffffffffffffffff166000526006602052604060002090565b6001600160a01b036064351615611be0578054906114526103eb6001600160a01b03841681565b3303611bb6578383916114686080840184611f2e565b949060208501956114798787611f2e565b9050159081611b9d575b61148c856120c8565b611499939060043561311e565b93849160a01c67ffffffffffffffff166114b290612a2c565b83547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617845592825163ffffffff16602084015161ffff166040805130602082015299908a90810103601f1981018b52611530908b610207565b604080516064356001600160a01b031660208083019190915281529a90611557908c610207565b6115618680611f2e565b86549c9160e08e901c60ff1691369061157992612a4b565b9060ff166115869161483d565b9060a0890151926040890161159b908a611ef8565b6115a59150612ac4565b946115b0908a611f2e565b9690976115bb61022a565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252600435811660208301529d909d1660408e0152600060608e015263ffffffff1660808d015261ffff1660a08c0152600060c08c015260e08b015261162f60048801612008565b6101008b01526101208a0152610140890152610160880152610180870152369061165892612a4b565b6101a08501526116666121e2565b6116736040840184611ef8565b9050611b21575b906116c261169d85949360406116f3970151606087015161048d60028701611f61565b60608601528060408601526116bc60808601516001600160a01b031690565b906148c7565b60c08601526116cf612b13565b976116dd6040840184611ef8565b15159050611b135760f01c90505b600435613c90565b63ffffffff9091166060840152602086019391845211611ae957611718825186614955565b6117256040860186611ef8565b90506119c9575b6117388195929561532f565b8085526020815191012090611751604085015151612b97565b9460408101958652606060009401935b604086015180518210156119365760206117946103eb6103eb611787866117db96612a03565b516001600160a01b031690565b6117a28460608b0151612a03565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612be0565b03915afa8015610c60576001600160a01b0391600091611908575b501680156118af57906000878b938783886118596118228860608f61181a9061226c565b980151612a03565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612d2b565b03925af18015610c6057816118879160019460009161188e575b508a51906118818383612a03565b52612a03565b5001611761565b6118a9913d8091833e6118a18183610207565b810190612c43565b38611873565b61059a6118c36117878460408b0151612a03565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611929915060203d811161192f575b6119218183610207565b8101906128cb565b386117f6565b503d611917565b61034285808b867fb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd8d6119688d61226c565b925193519051906119996040519283926001600160a01b03606435169767ffffffffffffffff600435169785612ef6565b0390a4610c147fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b611a2b611a156119df6104e56040890189611ef8565b60c0860180515115611ace5751905b602087015161ffff169060e08801519260643591611a10600435913690612b32565b614c22565b61018083015190611a25826129f6565b526129f6565b50611a4d6040611a418451828701515190612a03565b51015163ffffffff1690565b60a0611a5d6101808401516129f6565b5101515163ffffffff82168111611a7557505061172c565b61059a9250611a8d6104eb6104e560408a018a611ef8565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b50611ae3611adc8980611f2e565b3691612a4b565b906119ee565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff166116eb565b509350600191506040611b35910187611ef8565b905003610538576116f3838688946116c261169d611b91611b5f6104eb6104e56040880188611ef8565b6020611b716104e56040890189611ef8565b0135611b82602089015161ffff1690565b9060e08901519260043561349a565b9293949550505061167a565b9050611bac6040870187611ef8565b9050151590611483565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005261059a6004359067ffffffffffffffff60249216600452565b611c93915060203d602011611c99575b611c8b8183610207565b810190612a17565b386113b5565b503d611c81565b3461017a57602060031936011261017a5767ffffffffffffffff600435611cc681610346565b166000526006602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611d0a5760405167ffffffffffffffff9091168152602090f35b612399565b3461017a57602060031936011261017a576001600160a01b03600435611d3481610ae8565b611d3c6145ff565b16338114611da157807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461017a57602060031936011261017a57611de7600435610346565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611e21816101cf565b8281528260208201528260408201520152604051611e3e816101cf565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a57602001918160061b3603831361017a57565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a5760200191813603831361017a57565b906040519182815491828252602082019060005260206000209260005b818110611f9357505061023a92500383610207565b84546001600160a01b0316835260019485019487945060209093019201611f7e565b90600182811c92168015611ffe575b6020831014611fcf57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611fc4565b906040519182600082549261201c84611fb5565b80845293600181169081156120885750600114612041575b5061023a92500383610207565b90506000929192526020600020906000915b81831061206c57505090602061023a9282010138612034565b6020919350806001915483858901015201910190918492612053565b6020935061023a9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138612034565b906121c260046120d661023c565b9361214561213a82546120ff6120f2826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c16604089015261213460e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b612198612188600183015461216961215e8261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b6121a460028201611f61565b6101008601526121b660038201611f61565b61012086015201612008565b610140830152565b67ffffffffffffffff81116101ca5760051b60200190565b604051906121f1602083610207565b6000808352366020840137565b90612208826121ca565b6122156040519182610207565b828152601f1961222582946121ca565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90156122675790565b61222f565b35610d9281610ae8565b91908110156122675760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561017a570190565b35610d9281610346565b60ff81160361017a57565b35610d92816122c0565b63ffffffff81160361017a57565b35610d92816122d5565b903590601e198136030182121561017a570180359067ffffffffffffffff821161017a57602001918160051b3603831361017a57565b92919061232f816121ca565b9361233d6040519586610207565b602085838152019160051b810192831161017a57905b82821061235f57505050565b60208091833561236e81610ae8565b815201910190612353565b61ffff81160361017a57565b35610d9281612379565b35610d9281611118565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611d0a57565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611d0a57565b908160031b9180830460081490151715611d0a57565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611d0a57565b81810292918115918404141715611d0a57565b81811061248e575050565b60008155600101612483565b9067ffffffffffffffff83116101ca576801000000000000000083116101ca5781548383558084106124fe575b5090600052602060002060005b8381106124e15750505050565b60019060208435946124f286610ae8565b019381840155016124d4565b61251690836000528460206000209182019101612483565b386124c7565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610d9293818152019161251c565b9190601f811161255d57505050565b61023a926000526020600020906020601f840160051c83019310612589575b601f0160051c0190612483565b909150819061257c565b90929167ffffffffffffffff81116101ca576125b9816125b38454611fb5565b8461254e565b6000601f82116001146125f95781906125ea9394956000926125ee575b50506000198260011b9260031b1c19161790565b9055565b0135905038806125d6565b601f1982169461260e84600052602060002090565b91805b87811061264957508360019596971061262f575b505050811b019055565b60001960f88560031b161c19910135169055388080612625565b90926020600181928686013581550194019101612611565b359061023a826122c0565b359061023a82612379565b359061023a826122d5565b9035601e198236030181121561017a57016020813591019167ffffffffffffffff821161017a578160051b3603831361017a57565b9160209082815201919060005b8181106126d15750505090565b9091926020806001926001600160a01b0387356126ed81610ae8565b1681520194019291016126c4565b9035601e198236030181121561017a57016020813591019167ffffffffffffffff821161017a57813603831361017a57565b67ffffffffffffffff610d929392168152604060208201526127636040820161275584610358565b67ffffffffffffffff169052565b6127826127726020840161110d565b6001600160a01b03166060830152565b61279b61279160408401612661565b60ff166080830152565b6127b36127aa60608401611122565b151560a0830152565b6127cd6127c26080840161266c565b61ffff1660c0830152565b6127e76127dc60a0840161266c565b61ffff1660e0830152565b6128046127f660c08401612677565b63ffffffff16610100830152565b61289a61286d61282e61281a60e0860186612682565b6101606101208701526101a08601916126b7565b61283c610100860186612682565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0868403016101408701526126b7565b9261288f61287e610120830161110d565b6001600160a01b0316610160850152565b6101408101906126fb565b916101807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08286030191015261251c565b9081602091031261017a5751610d9281610ae8565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610c60576001600160a01b039160009161296e57501690565b612987915060203d60201161192f576119218183610207565b1690565b60405190610160820182811067ffffffffffffffff8211176101ca5760405260606101408360008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e082015282610100820152826101208201520152565b8051156122675760200190565b80518210156122675760209160051b010190565b9081602091031261017a5751610d9281611118565b67ffffffffffffffff1667ffffffffffffffff8114611d0a5760010190565b929192612a578261026a565b91612a656040519384610207565b82948184528183011161017a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101ca57604052606060a0836000815282602082015282604082015282808201528260808201520152565b90612ace826121ca565b612adb6040519182610207565b828152601f19612aeb82946121ca565b019060005b828110612afc57505050565b602090612b07612a82565b82828501015201612af0565b60405190612b20826101ae565b60606040838281528260208201520152565b919082604091031261017a57604051612b4a816101eb565b60208082948035612b5a81610ae8565b84520135910152565b60405190612b72602083610207565b600080835282815b828110612b8657505050565b806060602080938501015201612b7a565b90612ba1826121ca565b612bae6040519182610207565b828152601f19612bbe82946121ca565b019060005b828110612bcf57505050565b806060602080938501015201612bc3565b60409067ffffffffffffffff610d92949316815281602082015201906102be565b81601f8201121561017a578051612c178161026a565b92612c256040519485610207565b8184526020828401011161017a57610d92916020808501910161029b565b9060208282031261017a57815167ffffffffffffffff811161017a57610d929201612c01565b9080602083519182815201916020808360051b8301019401926000915b838310612c9557505050505090565b9091929394602080612d1c83601f1986600196030187528951908151815260a0612d0b612cf9612ce7612cd58887015160c08a88015260c08701906102be565b604087015186820360408801526102be565b606086015185820360608701526102be565b608085015184820360808601526102be565b9201519060a08184039101526102be565b97019301930191939290612c86565b919390610d929593612e73612e8b9260a08652612d5560a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612e5e612e46612e2e612e16612dfe612de88c61026060e08a0151916101c061018082015201906102be565b6101008801518d8203609f1901888f01526102be565b6101208701518c8203609f19016101c08e01526102be565b610140860151609f198c8303016101e08d01526102be565b610160850151609f198b8303016102008c01526102be565b610180840151609f198a8303016102208b0152612c69565b910151609f19878303016102408801526102be565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102be565b9080602083519182815201916020808360051b8301019401926000915b838310612ec957505050505090565b9091929394602080612ee783601f19866001960301875289516102be565b97019301930191939290612eba565b9493916001600160a01b03612f19921686526080602087015260808601906102be565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612f5b5750505050610d929394506060818403910152612e9d565b90919295602080612fbc83601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102be565b980192019201909291612f3d565b60405190610100820182811067ffffffffffffffff8211176101ca57604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161017a5790600490565b9093929384831161017a57841161017a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110613075575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261017a57825167ffffffffffffffff811161017a57826130cf918501612c01565b9260208101516130de816122d5565b92604082015167ffffffffffffffff811161017a57610d929201612c01565b60409067ffffffffffffffff610d929593168152816020820152019161251c565b9192909261312a612fca565b6004831015806132f3575b1561323d575090613145916155ec565b92613153604085015161583f565b80613222575b6040840161317681519260608701938451610120880151916158e3565b9092525260c0830151516131d2575b50608082016001600160a01b036131a382516001600160a01b031690565b16156131ae57505090565b6131c560e0610d929301516001600160a01b031690565b6001600160a01b03169052565b6131e66131e26060840151151590565b1590565b15613185577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff613236845163ffffffff1690565b1615613159565b94916000906132959261325e6103eb6103eb6002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a600485016130fd565b03915afa8015610c60576000906000906000906132c7575b60a088015263ffffffff16865290505b60c0850152613153565b5050506132e96132bd913d806000833e6132e18183610207565b8101906130a7565b91925082916132ad565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613349613343868661301b565b90613041565b1614613135565b60208183031261017a5780519067ffffffffffffffff821161017a57019080601f8301121561017a578151613384816121ca565b926133926040519485610207565b81845260208085019260051b82010192831161017a57602001905b8282106133ba5750505090565b6020809183516133c981610ae8565b8152019101906133ad565b95949060009460a09467ffffffffffffffff61341b956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102be565b930152565b9060028201809211611d0a57565b9060018201809211611d0a57565b6001019081600111611d0a57565b9060148201809211611d0a57565b90600c8201809211611d0a57565b91908201809211611d0a57565b6000198114611d0a5760010190565b80548210156122675760005260206000200190600090565b929394919060036134bf8567ffffffffffffffff166000526006602052604060002090565b01936001600160a01b036134d48184166128ec565b169182156136de576040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c60576000916136bf575b50156136af57613585600095969798604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016133d4565b03915afa928315610c605760009361368a575b5082511561367f578251906135b76135b284548094613466565b6121fe565b906000928394845b875181101561361e576135d5611787828a612a03565b6001600160a01b03811615613612579061360c6001926135fe6135f78a613473565b9989612a03565b906001600160a01b03169052565b016135bf565b5095506001809661360c565b509195509193613630575b5050815290565b60005b8281106136405750613629565b8061367961366661365360019486613482565b90546001600160a01b039160031b1c1690565b6135fe61367288613473565b9789612a03565b01613633565b9150610d9290611f61565b6136a89193503d806000833e6136a08183610207565b810190613350565b9138613598565b505050509250610d929150611f61565b6136d8915060203d602011611c9957611c8b8183610207565b38613540565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b939192936137306137288251865190613466565b865190613466565b9061374361373d836121fe565b92612b97565b94600096875b83518910156137a9578861379f61379260019361377a6137706117878e9f9d9e9d8b612a03565b6135fe838c612a03565b613798613787858c612a03565b519180938491613473565b9c612a03565b528b612a03565b5001979695613749565b959250929350955060005b865181101561383d576137ca6117878289612a03565b60006001600160a01b038216815b888110613811575b50509060019291156137f4575b50016137b4565b61380b906135fe61380489613473565b9888612a03565b386137ed565b816138226103eb611787848c612a03565b1461382f576001016137d8565b5060019150819050386137e0565b509390945060005b85518110156138ce5761385b6117878288612a03565b60006001600160a01b038216815b8781106138a2575b5050906001929115613885575b5001613845565b61389c906135fe61389588613473565b9787612a03565b3861387e565b816138b36103eb611787848b612a03565b146138c057600101613869565b506001915081905038613871565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101ca5760405260606080836000815260006020820152600060408201526000838201520152565b90613923826121ca565b6139306040519182610207565b828152601f1961394082946121ca565b019060005b82811061395157505050565b60209061395c6138da565b82828501015201613945565b9081606091031261017a57805161397e81612379565b916040602083015161398f816122d5565b920151610d92816122d5565b9160209082815201919060005b8181106139b55750505090565b9091926040806001926001600160a01b0387356139d181610ae8565b168152602087810135908201520194019291016139a8565b949391929067ffffffffffffffff16855260806020860152613a42613a23613a1185806126fb565b60a060808a015261012089019161251c565b613a3060208601866126fb565b90607f198984030160a08a015261251c565b6040840135601e198536030181121561017a578401916020833593019167ffffffffffffffff841161017a578360061b3603831361017a5761023a95613aca613aa183606097613aeb978d60c0607f1982613add9a030191015261399b565b91613ac0613ab088830161110d565b6001600160a01b031660e08d0152565b60808101906126fb565b90607f198b8403016101008c015261251c565b9087820360408901526102be565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611d0a57565b908160a091031261017a578051916020820151613b2c816122d5565b916040810151613b3b816122d5565b9160806060830151613b4c81612379565b920151610d9281611118565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610d929b9a9616885216602087015260408601521660608401521660808201528160a082015201906102be565b9081606091031261017a57805161397e816122d5565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611d0a57565b906000198201918211611d0a57565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611d0a57565b91908203918211611d0a57565b919082608091031261017a578151613c43816122d5565b916020810151916060604083015192015190565b8115613c61570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9382946000906000956040810194613cca613cc5613cc0885151613cb860408c01809c611ef8565b919050613466565b613420565b613919565b9660009586955b88518051881015613f2e576103eb6103eb6117878a613cef94612a03565b613d3c60206060880192613d048b8551612a03565b51908a6040518095819482937f958021a700000000000000000000000000000000000000000000000000000000845260048401612be0565b03915afa8015610c60576001600160a01b0391600091613f10575b50168015613ebc579060608e9392613d708b8451612a03565b5190613d8160208b015161ffff1690565b958b613dbc604051988995869485947f80485e25000000000000000000000000000000000000000000000000000000008652600486016139e9565b03915afa8015610c6057600193613e59938b8f8f95600080958197613e62575b509083929161ffff613e0485613dfd611787613e4d99613e539d9e51612a03565b9451612a03565b5191613e20613e1161024c565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b16604085015216606083015260808201526118818383612a03565b50613af6565b99613af6565b96019596613cd1565b613e539750611787965084939291509361ffff613e0482613dfd613e9f613e4d9960603d8111613eb5575b613e978183610207565b810190613968565b9c9196909c9d5050505050505090919293613ddc565b503d613e8d565b61059a88613ece6117878c8f51612a03565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613f28915060203d811161192f576119218183610207565b38613d57565b50919a9496929395509897968a613f458187611ef8565b9050614259575b50508651613f5990613bb6565b99613f676020860186611f2e565b91613f73915086611ef8565b9560609150019486613f848761226c565b91613f8f938a615afc565b613f998b89612a03565b52613fa48a88612a03565b50613faf8a88612a03565b516020015163ffffffff16613fc391613af6565b90613fce8a88612a03565b516040015163ffffffff16613fe291613af6565b91613feb61024c565b33815290600060208301819052604083015261ffff16606082015261400e610286565b6080820152865161401e90613be3565b906140298289612a03565b526140349087612a03565b506002546001600160a01b03169261404b9061226c565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610c605760009660009760009460009261421c575b506000965b865188101561419c576141256001916140ed886140e8876123c8565b613c57565b61410660606140fc8d8d612a03565b5101918251612470565b9052858a1461412d575b606061411c8b8b612a03565b51015190613466565b9701966140cc565b8b8873eba517d2000000000000000000000000000000006001600160a01b0361416060808c01516001600160a01b031690565b160361416e575b5050614110565b6140e861417a926123f8565b61419360606141898d8d612a03565b5101918251613466565b90528b88614167565b97965097505050506141d87f0000000000000000000000000000000000000000000000000000000000000000916140e863ffffffff84166123f8565b84116141e45750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b929850505061424491925060803d608011614252575b61423c8183610207565b810190613c2c565b9197909392909190386140c7565b503d614232565b610b246103eb6104eb6104e5614272948a989698611ef8565b926001600160a01b03600091515194169060e0880190815161429261024c565b6001600160a01b03851681529082602083015282604083015282606083015260808201526142c0878d612a03565b526142cb868c612a03565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3317103100000000000000000000000000000000000000000000000000000000600482015291602083602481875afa8015610c60578f948c89968f96948d948f9688916145e0575b506144d9575b50505050505015614380575b611a416143716143789561436b6020611a4161436b97604097612a03565b90613af6565b958b612a03565b90388a613f4c565b50506144069160608c6143b16104eb6104e56143aa6103eb6103eb6002546001600160a01b031690565b938b611ef8565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610c6057611a4161437160409261436b6020611a418f8b906143789b61436b9a60008060009261449b575b63ffffffff9293506144889060606144508888612a03565b51019461447d8a6144618a8a612a03565b51019160406144708b8b612a03565b51019063ffffffff169052565b9063ffffffff169052565b169052975097505050509550505061434d565b50505063ffffffff6144c76144889260603d6060116144d2575b6144bf8183610207565b810190613ba0565b909350915082614438565b503d6144b5565b8495985060a096975061452360206145196060826145106104e56145096104eb6104e58b61455c9c9d9e9f611ef8565b998d611ef8565b0135990161226c565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613b58565b03915afa8015610c60578592828c9391819082946145a4575b506145989060606145868888612a03565b51019261447d60206144618a8a612a03565b5288888f8c8138614341565b91505061459892506145ce915060a03d60a0116145d9575b6145c68183610207565b810190613b10565b949192919050614575565b503d6145bc565b6145f9915060203d602011611c9957611c8b8183610207565b3861433b565b6001600160a01b0360015416330361461357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b80519161464b815184613466565b9283156147445760005b848110614663575050505050565b81811015614729576146786117878286612a03565b6001600160a01b0381168015610c65576146918361342e565b8781106146a357505050600101614655565b84811015614706576001600160a01b036146c0611787838a612a03565b1682146146cf57600101614691565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b0361472461178761471e8885613c1f565b89612a03565b6146c0565b61473f6117876147398484613c1f565b85612a03565b614678565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156122675760051b0190565b9081602091031261017a575190565b91608061023a9294936147dd8160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610d929281815201906102be565b906148208261026a565b61482d6040519182610207565b828152601f19612225829461026a565b90815160208210806148bd575b61488c57036148565790565b610a21906040519182917f3aeba39000000000000000000000000000000000000000000000000000000000835260048301614805565b50906020810151908161489e8461242a565b1c61485657506148ad82614816565b9160200360031b1b602082015290565b506020811461484a565b918251601481029080820460141490151715611d0a576148e96148ee9161343c565b61344a565b906149006148fb83613458565b614816565b90601461490c836129f6565b5360009260215b865185101561493e57601460019161492e611787888b612a03565b60601b8187015201940193614913565b919550936020935060601b90820152828152012090565b906149656103eb6060840161226c565b614976600019936040810190611ef8565b90506149ee575b6149878251613be3565b9260005b848110614999575050505050565b8082600192146149e95760606149af8287612a03565b51015180156149e3576149dd906149d76149c98489612a03565b51516001600160a01b031690565b86615d89565b0161498b565b506149dd565b6149dd565b91506149fa8151613bf2565b91614a086149c98484612a03565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f331710310000000000000000000000000000000000000000000000000000000060048201526020816024816001600160a01b0386165afa908115610c6057600091614aa1575b50614a81575b5061497d565b614a9b906060614a918686612a03565b5101519083615d89565b38614a7b565b614aba915060203d602011611c9957611c8b8183610207565b38614a75565b60405190614acd826101eb565b60606020838281520152565b919060408382031261017a5760405190614af2826101eb565b8193805167ffffffffffffffff811161017a5782614b11918301612c01565b835260208101519167ffffffffffffffff831161017a57602092614b359201612c01565b910152565b9060208282031261017a57815167ffffffffffffffff811161017a57610d929201614ad9565b9060806001600160a01b0381614b7f855160a0865260a08601906102be565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610d92928181520190614b60565b919060408382031261017a57825167ffffffffffffffff811161017a57602091614bef918501614ad9565b92015190565b61ffff614c0e610d929593606084526060840190614b60565b9316602082015260408184039101526102be565b909193929594614c30612a82565b50602082018051156150ca57614c53610b246103eb85516001600160a01b031690565b946001600160a01b038616916040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180614cba60048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610c60576000916150ab575b501561506257614d4688999a825192614ce5614ac0565b5051614d31614cfb89516001600160a01b031690565b926040614d0661024c565b9e8f908152614d228d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c6057600091615043575b5015614f4b5750906000929183614df29899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614bf5565b03925af18015610c6057600094600091614f03575b50614ed8614e419596614e7b614e4f614e6d956117876020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b610207565b6040519586918683019190916001600160a01b036020820193169052565b03601f198101865285610207565b614eb7610960614ead614ebd8951614eb7610960614ead8c67ffffffffffffffff166000526006602052604060002090565b5460e01c60ff1690565b9061483d565b9767ffffffffffffffff166000526006602052604060002090565b93015193614ee461025b565b958652602086015260408501526060840152608083015260a082015290565b614e4195506020915061178796614e7b614e4f614e6d95614f39614ed8953d806000833e614f318183610207565b810190614bc4565b9b909b96505095505050969550614e07565b9793929061ffff166150195751614fef57614f9a6000939184926040519586809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614bb3565b03925af1908115610c6057614e4195614e7b614e4f614ed893611787602096614e6d98600091614fcc575b5099614e23565b614fe991503d806000833e614fe18183610207565b810190614b3a565b38614fc5565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b61505c915060203d602011611c9957611c8b8183610207565b38614daa565b61059a61507686516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b6150c4915060203d602011611c9957611c8b8183610207565b38614cce565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b95926151c6947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b906151e26020928281519485920161029b565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b16815261522e825180936020898501910161029b565b019160f81b168382015261524c82518093602060028501910161029b565b01019160f81b168382015261526b82518093602060028501910161029b565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff00000000000000000000000000000000000000000000000000000000000000610d929f9e9c9860f81b1681526152e0825180936020898501910161029b565b019160f01b16838201526152fe82518093602060038501910161029b565b01019160f01b168382015261531c825180936020898501910161029b565b01019160f01b16600282015201906151cf565b60e081019060ff825151116155d657610100810160ff815151116155c05761012082019260ff845151116155aa57610140830160ff815151116155945761016084019061ffff8251511161557e57610180850193600185515111615566576101a086019361ffff8551511161555057606095518051615514575b50865167ffffffffffffffff16602088015167ffffffffffffffff169060408901516153dc9067ffffffffffffffff1690565b9860608101516153ef9063ffffffff1690565b9060808101516154029063ffffffff1690565b60a082015161ffff169160c00151926040519c8d966020880196615425976150f4565b03601f19810188526154379088610207565b519081516154459060ff1690565b9051805160ff16985190815161545b9060ff1690565b906040519a8b956020870195615470966151e6565b03601f19810187526154829087610207565b519182516154909060ff1690565b9151805161ffff169480516154a69061ffff1690565b92519283516154b69061ffff1690565b9260405197889760208901976154cb98615271565b03601f19810182526154dd9082610207565b604051928392602084016154f0916151cf565b6154f9916151cf565b615502916151cf565b03601f1981018252610d929082610207565b615529919650615523906129f6565b51615fb5565b9461ffff86511161553a57386153a9565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b600052602660045260246000fd5b635a102da160e11b60005261059a6024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b906155f5612fca565b916011821061580b5780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216036157985750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a615674816121fe565b6040860190815261568482612b97565b906060870191825260005b83811061574c57505050506156f783836156ed6156e16156d76156d06156bd615701988761570b9c9b6160dc565b6001600160a01b0390911660808d015290565b85856161b2565b9291903691612a4b565b60a08a0152838361621a565b9491903691612a4b565b60c08801526161b2565b9391903691612a4b565b60e0840152810361571a575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b8060019161579161577b61577461576761578b9a8d8d6160dc565b91906135fe868a51612a03565b8b8b6161b2565b9391889a919a51949a3691612a4b565b92612a03565b520161568f565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b80519060005b82811061585157505050565b60018101808211611d0a575b83811061586d5750600101615845565b6001600160a01b0361587f8385612a03565b51166158916103eb6117878487612a03565b1461589e5760010161585d565b61059a6158ae6117878486612a03565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b90939192815115615a6f57506158f88161583f565b80519260005b84811061590d57509093925050565b61591d6103eb6117878386612a03565b1561592a576001016158fe565b9391946159446135b261593c85613be3565b845190613466565b9261596161595c61595483613be3565b855190613466565b612b97565b968792600097885b848110615a115750505050505060005b8151811015615a04576000805b8681106159c5575b50906001916159c0576159ba6159a76117878386612a03565b6135fe6159b389613473565b9887612a03565b01615979565b6159ba565b6159d26117878287612a03565b6001600160a01b036159ea6103eb6117878789612a03565b9116146159f957600101615986565b50600190508061598e565b5050909180825283529190565b909192939498828214615a655790615a57615a4a83615a3d8b6135fe60019761578b611787898e612a03565b615a506137878589612a03565b9e612a03565b528c612a03565b505b01908994939291615969565b9850600190615a59565b9193505015615a8a5750615a816121e2565b90610d92612b63565b90610d928251612b97565b9081602091031261017a5751610d9281612379565b93615ae760809461ffff6001600160a01b039567ffffffffffffffff615af5969b9a9b16895216602088015260a0604088015260a0870190610c8f565b9085820360608701526102be565b9416910152565b92919092615b086138da565b50615b278167ffffffffffffffff166000526006602052604060002090565b805490959060e01c60ff169160808501928351615b4a906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff16615b6a91613af6565b96615b7690608d613466565b9460a0870195865151615b8891613466565b9160ff1691615b9683612440565b615b9f91613466565b91615bab906067613466565b615bb491612470565b615bbd91613466565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b03871603615c4a5750505061ffff9250615c3c90615c2f6000935b5195615c22615c1361024c565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b615c606103eb602094976001600160a01b031690565b906040615c718583015161ffff1690565b91015192615cb1875198604051998a96879586957ff238895800000000000000000000000000000000000000000000000000000000875260048701615aaa565b03915afa908115610c6057615c2f615c3c9261ffff95600091615cd6575b5093615c06565b615cf8915060203d602011615cfe575b615cf08183610207565b810190615a95565b38615ccf565b503d615ce6565b80600052600560205260406000205415600014615d8357600454680100000000000000008110156101ca5760018101600455600060045482101561226757600490527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01819055600454906000526005602052604060002055600190565b50600090565b91602091600091604051906001600160a01b03858301937fa9059cbb000000000000000000000000000000000000000000000000000000008552166024830152604482015260448152615ddd606482610207565b519082855af1156128e0576000513d615e3c57506001600160a01b0381163b155b615e055750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615dfe565b6004548110156122675760046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b015490565b966002615f8997615f566022610d929f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090615f569f82615f569c615f5d9c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615f07825180936020898501910161029b565b019160f81b1683820152615f2582518093602060238501910161029b565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b01906151cf565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff825151116160c657604081019160ff835151116160b057606082019160ff8351511161609a57608081019260ff845151116160845760a0820161ffff8151511161606e57610d92946160609351945191616017835160ff1690565b975191616025835160ff1690565b945190616033825160ff1690565b905193616041855160ff1690565b935196616050885161ffff1690565b966040519c8d9b60208d01615e7a565b03601f198101835282610207565b635a102da160e11b600052602b60045260246000fd5b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b9291909260018201918483116161805781013560001a82811561617557506014810361614857820193841161611457013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b91906002820191818311616180578381013560f01c01600201928184116161e6579183916161df93613029565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b91906001820191818311616180578381013560001a01600101928184116161e6579183916161df9361302956fea164736f6c634300081a000a",
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
	DestChainSelector uint64
	Sender            common.Address
	MessageId         [32]byte
	FeeToken          common.Address
	EncodedMessage    []byte
	Receipts          []OnRampReceipt
	VerifierBlobs     [][]byte
	Raw               types.Log
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
	return common.HexToHash("0xb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd")
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
