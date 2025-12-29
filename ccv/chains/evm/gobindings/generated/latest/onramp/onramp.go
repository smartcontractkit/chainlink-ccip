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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100604052346103805760405161603b38819003601f8101601f191683016001600160401b03811184821017610385578392829160405283398101039060e0821261038057608082126103805761005561039b565b81519092906001600160401b03811681036103805783526020820151926001600160a01b03841684036103805760208101938452604083015163ffffffff81168103610380576040820190815260606100af8186016103ba565b83820190815293607f1901126103805760405192606084016001600160401b03811185821017610385576040526100e8608086016103ba565b845260a08501519485151586036103805760c061010c9160208701978852016103ba565b9560408501968752331561036f57600180546001600160a01b0319163317905583516001600160401b031615801561035d575b801561034b575b801561033c575b61030f5792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e05281511615801561032a575b8015610320575b61030f5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e0926000606061020d61039b565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff9283169116606061025861039b565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a1604051615c6c90816103cf8239608051818181610a70015281816113c40152611c66015260a0518181816111820152611c92015260c051818181611ced01526126c7015260e051818181611cbe0152613e590152f35b6306b7c75960e31b60005260046000fd5b5081511515610192565b5082516001600160a01b03161561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761038557604052565b51906001600160a01b03821682036103805756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd5780632490769e146100f857806348a98aa4146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da57806390423fa2146100d5578063df0aa9e9146100d0578063e8d80861146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611be8565b611b37565b611ac8565b6110dc565b610f58565b610f11565b610e6b565b610df8565b610d26565b610aed565b610aa8565b6105c1565b610379565b6102eb565b3461016a57600060031936011261016a576080610122611c2e565b61016860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b634e487b7160e01b600052604160045260246000fd5b610160810190811067ffffffffffffffff8211176101a257604052565b61016f565b6060810190811067ffffffffffffffff8211176101a257604052565b6080810190811067ffffffffffffffff8211176101a257604052565b6040810190811067ffffffffffffffff8211176101a257604052565b90601f601f19910116810190811067ffffffffffffffff8211176101a257604052565b6040519061022e6101c0836101fb565b565b6040519061022e610160836101fb565b6040519061022e60a0836101fb565b6040519061022e60c0836101fb565b67ffffffffffffffff81116101a257601f01601f191660200190565b604051906102896020836101fb565b60008252565b60005b8381106102a25750506000910152565b8181015183820152602001610292565b90601f19601f6020936102d08151809281875287808801910161028f565b0116010190565b9060206102e89281815201906102b2565b90565b3461016a57600060031936011261016a5761034a604080519061030e81836101fb565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102b2565b0390f35b67ffffffffffffffff81160361016a57565b359061022e8261034e565b908160a091031261016a5790565b3461016a57604060031936011261016a576004356103968161034e565b60243567ffffffffffffffff811161016a576103b690369060040161036b565b6103d48267ffffffffffffffff166000526004602052604060002090565b805491906001600160a01b036103f38185165b6001600160a01b031690565b1615610526579061034a9361049c93926104396104136080850185611d15565b6104206020870187611d15565b905015918261050d575b61043385611ecc565b86612e15565b93610488610445611fe6565b60408601906104548288611d48565b90506104bc575b6040880161047e81519260608b0193845161047860028b01611d7e565b916133dd565b9092525285611d48565b151590506104ae5760f01c90505b90613940565b60405190815292839250602083019150565b506001015461ffff16610496565b506105086104db6104d66104d0848a611d48565b90612049565b612057565b60206104ea6104d0858b611d48565b01356104fb60208b015161ffff1690565b9060e08b01519289613195565b61045b565b915061051c6040870187611d48565b905015159161042a565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156105bd5781359367ffffffffffffffff85116105ba57506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a576105cf36610562565b906105d861427d565b6000915b8083106105e557005b6105f0838284612061565b926105fa846120a1565b67ffffffffffffffff81169081158015610a64575b8015610a4e575b8015610a35575b6109fe57856108749161088e6108848361087e61068f60e083019561067561066f61064889876120d8565b61066761065d6101008a95949501809a6120d8565b949092369161210e565b92369161210e565b906142bb565b67ffffffffffffffff166000526004602052604060002090565b9687956106ca6106a160208a01612057565b88906001600160a01b031673ffffffffffffffffffffffffffffffffffffffff19825416179055565b61086e61083760c060408b019a6107316106e38d6120b6565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b61078f61074060808301612170565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b6107d260016107a060a08401612170565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6108318d6107e26060840161217a565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b016120ce565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d6120d8565b906003890161226c565b8a6120d8565b906002860161226c565b6101208801906108a06103e783612057565b156109d4576108b16108fb92612057565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b61014087019061091f610919610911848b611d15565b9390506120b6565b60ff1690565b0361099057956109767f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb269261096561095b600198999a85611d15565b9060048401612365565b5460a01c67ffffffffffffffff1690565b610985604051928392836124ff565b0390a20191906105dc565b61099a9087611d15565b906109d06040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161230f565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a4760c088016120ce565b161561061d565b5060ff610a5d604088016120b6565b1615610616565b5067ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016821461060f565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610ac460043561034e565b610ad8602435610ad381610a97565b612682565b6040516001600160a01b039091168152602090f35b3461016a57610afb36610562565b906001600160a01b03600354169160005b818110610b1557005b610b266103e76104d6838587614416565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610c09576001948892600092610bd9575b5081610b8d575b5050505001610b0c565b81610bbd7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610bcd946156e9565b6040519081529081906020820190565b0390a338858180610b83565b610bfb91925060203d8111610c02575b610bf381836101fb565b810190614426565b9038610b7c565b503d610be9565b612676565b906020808351928381520192019060005b818110610c2c5750505090565b82516001600160a01b0316845260209384019390920191600101610c1f565b906102e89160208152610c6a6020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015115156080820152608082015161ffff1660a082015260a082015161ffff1660c082015260c082015163ffffffff1660e082015260e08201516001600160a01b0316610100820152610140610d10610cfa610100850151610160610120860152610180850190610c0e565b610120850151601f198583030184860152610c0e565b92015190610160601f19828503019101526102b2565b3461016a57602060031936011261016a5767ffffffffffffffff600435610d4c8161034e565b6060610140604051610d5d81610185565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e08201528261010082015282610120820152015216600052600460205261034a610dba6040600020611ecc565b60405191829182610c4b565b61022e9092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a57600060408051610e18816101a7565b828152826020820152015261034a604051610e32816101a7565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610dc6565b3461016a57600060031936011261016a576000546001600160a01b0381163303610ee75773ffffffffffffffffffffffffffffffffffffffff19600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061022e82610a97565b8015150361016a57565b359061022e82610f43565b3461016a57606060031936011261016a576000604051610f77816101a7565b600435610f8381610a97565b8152602435610f9181610f43565b6020820190815260443590610fa582610a97565b60408301918252610fb461427d565b6001600160a01b03835116159182156110c9575b5081156110be575b506110965780516002805460208401517fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160a01b039384161790151560a01b74ff00000000000000000000000000000000000000001617905560408201516003805473ffffffffffffffffffffffffffffffffffffffff1916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c90611081611c2e565b61109060405192839283614435565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610fd0565b516001600160a01b031615915038610fc8565b3461016a57608060031936011261016a576110f860043561034e565b60243567ffffffffffffffff811161016a5761111890369060040161036b565b60443590611127606435610a97565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610c0957600091611a99575b50611a5c5760025460a01c60ff16611a3257611209740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61122960043567ffffffffffffffff166000526004602052604060002090565b6001600160a01b036064351615611a08578054906112506103e76001600160a01b03841681565b33036119de578383916112666080840184611d15565b949060208501956112778787611d15565b90501590816119c5575b61128a85611ecc565b6112979390600435612e15565b93849160a01c67ffffffffffffffff166112b090612736565b83547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617845592825163ffffffff16602084015161ffff166040805130602082015299908a90810103601f1981018b5261132e908b6101fb565b604080516064356001600160a01b031660208083019190915281529a90611355908c6101fb565b61135f8680611d15565b86549c9160e08e901c60ff1691369061137792612755565b9060ff16611384916144d4565b9060a08901519260408901611399908a611d48565b6113a391506127ce565b946113ae908a611d15565b9690976113b961021e565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252600435811660208301529d909d1660408e0152600060608e015263ffffffff1660808d015261ffff1660a08c0152600060c08c015260e08b015261142d60048801611e0c565b6101008b01526101208a0152610140890152610160880152610180870152369061145692612755565b6101a0850152611464611fe6565b6114716040840184611d48565b905061191f575b906114c061149b85949360406114f1970151606087015161047860028701611d7e565b60608601528060408601526114ba60808601516001600160a01b031690565b9061455e565b60c08601526114cd61281d565b976114db6040840184611d48565b151590506119115760f01c90505b600435613940565b63ffffffff90911660608401526020860193918452116118e7576115168251866145ec565b6115236040860186611d48565b90506117c7575b61153681959295614f49565b808552602081519101209061154f60408501515161288e565b9460408101958652606060009401935b604086015180518210156117345760206115926103e76103e7611585866115d99661287a565b516001600160a01b031690565b6115a08460608b015161287a565b519060405180809581947f958021a7000000000000000000000000000000000000000000000000000000008352600435600484016128d7565b03915afa8015610c09576001600160a01b0391600091611706575b501680156116ad57906000878b938783886116576116208860608f61161890612057565b98015161287a565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612a22565b03925af18015610c0957816116859160019460009161168c575b508a519061167f838361287a565b5261287a565b500161155f565b6116a7913d8091833e61169f81836101fb565b81019061293a565b38611671565b61055e6116c16115858460408b015161287a565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611727915060203d811161172d575b61171f81836101fb565b810190612661565b386115f4565b503d611715565b61034a85808b867fb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd8d6117668d612057565b925193519051906117976040519283926001600160a01b03606435169767ffffffffffffffff600435169785612bed565b0390a4610bbd7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6118296118136117dd6104d06040890189611d48565b60c08601805151156118cc5751905b602087015161ffff169060e0880151926064359161180e60043591369061283c565b614887565b610180830151906118238261286d565b5261286d565b5061184b604061183f845182870151519061287a565b51015163ffffffff1690565b60a061185b61018084015161286d565b5101515163ffffffff8216811161187357505061152a565b61055e925061188b6104d66104d060408a018a611d48565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b506118e16118da8980611d15565b3691612755565b906117ec565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff166114e9565b509350600191506040611933910187611d48565b90500361199b576114f1838688946114c061149b61198f61195d6104d66104d06040880188611d48565b602061196f6104d06040890189611d48565b0135611980602089015161ffff1690565b9060e089015192600435613195565b92939495505050611478565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b90506119d46040870187611d48565b9050151590611281565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005261055e6004359067ffffffffffffffff60249216600452565b611abb915060203d602011611ac1575b611ab381836101fb565b810190612721565b386111b3565b503d611aa9565b3461016a57602060031936011261016a5767ffffffffffffffff600435611aee8161034e565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611b325760405167ffffffffffffffff9091168152602090f35b612184565b3461016a57602060031936011261016a576001600160a01b03600435611b5c81610a97565b611b6461427d565b16338114611bbe578073ffffffffffffffffffffffffffffffffffffffff1960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611c0460043561034e565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611c3e816101c3565b8281528260208201528260408201520152604051611c5b816101c3565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611db057505061022e925003836101fb565b84546001600160a01b0316835260019485019487945060209093019201611d9b565b90600182811c92168015611e02575b6020831014611dec57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611de1565b9060405191826000825492611e2084611dd2565b8084529360018116908115611e8c5750600114611e45575b5061022e925003836101fb565b90506000929192526020600020906000915b818310611e7057505090602061022e9282010138611e38565b6020919350806001915483858901015201910190918492611e57565b6020935061022e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611e38565b90611fc66004611eda610230565b93611f49611f3e8254611f03611ef6826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c166040890152611f3860e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b611f9c611f8c6001830154611f6d611f628261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b611fa860028201611d7e565b610100860152611fba60038201611d7e565b61012086015201611e0c565b610140830152565b67ffffffffffffffff81116101a25760051b60200190565b60405190611ff56020836101fb565b6000808352366020840137565b9061200c82611fce565b61201960405191826101fb565b828152601f196120298294611fce565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b90156120525790565b612033565b356102e881610a97565b91908110156120525760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561016a570190565b356102e88161034e565b60ff81160361016a57565b356102e8816120ab565b63ffffffff81160361016a57565b356102e8816120c0565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b92919061211a81611fce565b9361212860405195866101fb565b602085838152019160051b810192831161016a57905b82821061214a57505050565b60208091833561215981610a97565b81520191019061213e565b61ffff81160361016a57565b356102e881612164565b356102e881610f43565b634e487b7160e01b600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611b3257565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611b3257565b908160031b9180830460081490151715611b3257565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611b3257565b81810292918115918404141715611b3257565b818110612260575050565b60008155600101612255565b9067ffffffffffffffff83116101a2576801000000000000000083116101a25781548383558084106122d0575b5090600052602060002060005b8381106122b35750505050565b60019060208435946122c486610a97565b019381840155016122a6565b6122e890836000528460206000209182019101612255565b38612299565b601f8260209493601f19938186528686013760008582860101520116010190565b9160206102e89381815201916122ee565b9190601f811161232f57505050565b61022e926000526020600020906020601f840160051c8301931061235b575b601f0160051c0190612255565b909150819061234e565b90929167ffffffffffffffff81116101a25761238b816123858454611dd2565b84612320565b6000601f82116001146123cb5781906123bc9394956000926123c0575b50506000198260011b9260031b1c19161790565b9055565b0135905038806123a8565b601f198216946123e084600052602060002090565b91805b87811061241b575083600195969710612401575b505050811b019055565b60001960f88560031b161c199101351690553880806123f7565b909260206001819286860135815501940191016123e3565b359061022e826120ab565b359061022e82612164565b359061022e826120c0565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b8181106124a35750505090565b9091926020806001926001600160a01b0387356124bf81610a97565b168152019401929101612496565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b67ffffffffffffffff6102e89392168152604060208201526125356040820161252784610360565b67ffffffffffffffff169052565b61255461254460208401610f38565b6001600160a01b03166060830152565b61256d61256360408401612433565b60ff166080830152565b61258561257c60608401610f4d565b151560a0830152565b61259f6125946080840161243e565b61ffff1660c0830152565b6125b96125ae60a0840161243e565b61ffff1660e0830152565b6125d66125c860c08401612449565b63ffffffff16610100830152565b61264e6126216126006125ec60e0860186612454565b6101606101208701526101a0860191612489565b61260e610100860186612454565b90603f1986840301610140870152612489565b926126436126326101208301610f38565b6001600160a01b0316610160850152565b6101408101906124cd565b91610180603f19828603019101526122ee565b9081602091031261016a57516102e881610a97565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610c09576001600160a01b039160009161270457501690565b61271d915060203d60201161172d5761171f81836101fb565b1690565b9081602091031261016a57516102e881610f43565b67ffffffffffffffff1667ffffffffffffffff8114611b325760010190565b9291926127618261025e565b9161276f60405193846101fb565b82948184528183011161016a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101a257604052606060a0836000815282602082015282604082015282808201528260808201520152565b906127d882611fce565b6127e560405191826101fb565b828152601f196127f58294611fce565b019060005b82811061280657505050565b60209061281161278c565b828285010152016127fa565b6040519061282a826101a7565b60606040838281528260208201520152565b919082604091031261016a57604051612854816101df565b6020808294803561286481610a97565b84520135910152565b8051156120525760200190565b80518210156120525760209160051b010190565b9061289882611fce565b6128a560405191826101fb565b828152601f196128b58294611fce565b019060005b8281106128c657505050565b8060606020809385010152016128ba565b60409067ffffffffffffffff6102e8949316815281602082015201906102b2565b81601f8201121561016a57805161290e8161025e565b9261291c60405194856101fb565b8184526020828401011161016a576102e8916020808501910161028f565b9060208282031261016a57815167ffffffffffffffff811161016a576102e892016128f8565b9080602083519182815201916020808360051b8301019401926000915b83831061298c57505050505090565b9091929394602080612a1383601f1986600196030187528951908151815260a0612a026129f06129de6129cc8887015160c08a88015260c08701906102b2565b604087015186820360408801526102b2565b606086015185820360608701526102b2565b608085015184820360808601526102b2565b9201519060a08184039101526102b2565b9701930193019193929061297d565b9193906102e89593612b6a612b829260a08652612a4c60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612b55612b3d612b25612b0d612af5612adf8c61026060e08a0151916101c061018082015201906102b2565b6101008801518d8203609f1901888f01526102b2565b6101208701518c8203609f19016101c08e01526102b2565b610140860151609f198c8303016101e08d01526102b2565b610160850151609f198b8303016102008c01526102b2565b610180840151609f198a8303016102208b0152612960565b910151609f19878303016102408801526102b2565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102b2565b9080602083519182815201916020808360051b8301019401926000915b838310612bc057505050505090565b9091929394602080612bde83601f19866001960301875289516102b2565b97019301930191939290612bb1565b9493916001600160a01b03612c10921686526080602087015260808601906102b2565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612c5257505050506102e89394506060818403910152612b94565b90919295602080612cb383601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102b2565b980192019201909291612c34565b60405190610100820182811067ffffffffffffffff8211176101a257604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612d6c575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a5782612dc69185016128f8565b926020810151612dd5816120c0565b92604082015167ffffffffffffffff811161016a576102e892016128f8565b60409067ffffffffffffffff6102e8959316815281602082015201916122ee565b91929092612e21612cc1565b600483101580613007575b15612f51575090612e3c916151e6565b92612e4a60408501516153d5565b6040840190815151159081612f27575b50612f09575b5060c083015151612eb9575b50608082016001600160a01b03612e8a82516001600160a01b031690565b1615612e9557505090565b612eac60e06102e89301516001600160a01b031690565b6001600160a01b03169052565b612ecd612ec96060840151151590565b1590565b15612e6c577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b612f1c906101208401518091525161288e565b606084015238612e60565b905080612f36575b1538612e5a565b5063ffffffff612f4a855163ffffffff1690565b1615612f2f565b9491600090612fa992612f726103e76103e76002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612df4565b03915afa8015610c0957600090600090600090612fdb575b60a088015263ffffffff16865290505b60c0850152612e4a565b505050612ffd612fd1913d806000833e612ff581836101fb565b810190612d9e565b9192508291612fc1565b5063302326cb60e01b7fffffffff0000000000000000000000000000000000000000000000000000000061304461303e8686612d12565b90612d38565b1614612e2c565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a57815161307f81611fce565b9261308d60405194856101fb565b81845260208085019260051b82010192831161016a57602001905b8282106130b55750505090565b6020809183516130c481610a97565b8152019101906130a8565b95949060009460a09467ffffffffffffffff613116956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102b2565b930152565b9060028201809211611b3257565b9060018201809211611b3257565b6001019081600111611b3257565b9060148201809211611b3257565b90600c8201809211611b3257565b91908201809211611b3257565b6000198114611b325760010190565b80548210156120525760005260206000200190600090565b929394919060036131ba8567ffffffffffffffff166000526004602052604060002090565b01936001600160a01b036131cf818416612682565b169182156133a7576040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610c0957600091613388575b50156133785761324e600095969798604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016130cf565b03915afa928315610c0957600093613353575b508251156133485782519061328061327b84548094613161565b612002565b906000928394845b87518110156132e75761329e611585828a61287a565b6001600160a01b038116156132db57906132d56001926132c76132c08a61316e565b998961287a565b906001600160a01b03169052565b01613288565b509550600180966132d5565b5091955091936132f9575b5050815290565b60005b82811061330957506132f2565b8061334261332f61331c6001948661317d565b90546001600160a01b039160031b1c1690565b6132c761333b8861316e565b978961287a565b016132fc565b91506102e890611d7e565b6133719193503d806000833e61336981836101fb565b81019061304b565b9138613261565b5050505092506102e89150611d7e565b6133a1915060203d602011611ac157611ab381836101fb565b38613209565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b939192936133f96133f18251865190613161565b865190613161565b9061340c61340683612002565b9261288e565b94600096875b8351891015613472578861346861345b6001936134436134396115858e9f9d9e9d8b61287a565b6132c7838c61287a565b613461613450858c61287a565b51918093849161316e565b9c61287a565b528b61287a565b5001979695613412565b959250929350955060005b865181101561350657613493611585828961287a565b60006001600160a01b038216815b8881106134da575b50509060019291156134bd575b500161347d565b6134d4906132c76134cd8961316e565b988861287a565b386134b6565b816134eb6103e7611585848c61287a565b146134f8576001016134a1565b5060019150819050386134a9565b509390945060005b855181101561359757613524611585828861287a565b60006001600160a01b038216815b87811061356b575b505090600192911561354e575b500161350e565b613565906132c761355e8861316e565b978761287a565b38613547565b8161357c6103e7611585848b61287a565b1461358957600101613532565b50600191508190503861353a565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101a25760405260606080836000815260006020820152600060408201526000838201520152565b906135ec82611fce565b6135f960405191826101fb565b828152601f196136098294611fce565b019060005b82811061361a57505050565b6020906136256135a3565b8282850101520161360e565b9081606091031261016a57805161364781612164565b9160406020830151613658816120c0565b9201516102e8816120c0565b9160209082815201919060005b81811061367e5750505090565b9091926040806001926001600160a01b03873561369a81610a97565b16815260208781013590820152019401929101613671565b949391929067ffffffffffffffff1685526080602086015261370b6136ec6136da85806124cd565b60a060808a01526101208901916122ee565b6136f960208601866124cd565b90607f198984030160a08a01526122ee565b6040840135601e198536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a5761022e9561379361376a836060976137b4978d60c0607f19826137a69a0301910152613664565b91613789613779888301610f38565b6001600160a01b031660e08d0152565b60808101906124cd565b90607f198b8403016101008c01526122ee565b9087820360408901526102b2565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611b3257565b908160a091031261016a5780519160208201516137f5816120c0565b916040810151613804816120c0565b916080606083015161381581612164565b9201516102e881610f43565b9260c0946001600160a01b039167ffffffffffffffff61ffff95846102e89b9a9616885216602087015260408601521660608401521660808201528160a082015201906102b2565b9081606091031261016a578051613647816120c0565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611b3257565b906000198201918211611b3257565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611b3257565b91908203918211611b3257565b919082608091031261016a57815161390c816120c0565b916020810151916060604083015192015190565b811561392a570490565b634e487b7160e01b600052601260045260246000fd5b938294600090600095604081019461397a61397561397088515161396860408c01809c611d48565b919050613161565b61311b565b6135e2565b9660009586955b88518051881015613bde576103e76103e76115858a61399f9461287a565b6139ec602060608801926139b48b855161287a565b51908a6040518095819482937f958021a7000000000000000000000000000000000000000000000000000000008452600484016128d7565b03915afa8015610c09576001600160a01b0391600091613bc0575b50168015613b6c579060608e9392613a208b845161287a565b5190613a3160208b015161ffff1690565b958b613a6c604051988995869485947f80485e25000000000000000000000000000000000000000000000000000000008652600486016136b2565b03915afa8015610c0957600193613b09938b8f8f95600080958197613b12575b509083929161ffff613ab485613aad611585613afd99613b039d9e5161287a565b945161287a565b5191613ad0613ac1610240565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261167f838361287a565b506137bf565b996137bf565b96019596613981565b613b039750611585965084939291509361ffff613ab482613aad613b4f613afd9960603d8111613b65575b613b4781836101fb565b810190613631565b9c9196909c9d5050505050505090919293613a8c565b503d613b3d565b61055e88613b7e6115858c8f5161287a565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613bd8915060203d811161172d5761171f81836101fb565b38613a07565b50919a9496929395509897968a613bf58187611d48565b9050613f09575b50508651613c099061387f565b99613c176020860186611d15565b91613c23915086611d48565b9560609150019486613c3487612057565b91613c3f938a6154e0565b613c498b8961287a565b52613c548a8861287a565b50613c5f8a8861287a565b516020015163ffffffff16613c73916137bf565b90613c7e8a8861287a565b516040015163ffffffff16613c92916137bf565b91613c9b610240565b33815290600060208301819052604083015261ffff166060820152613cbe61027a565b60808201528651613cce906138ac565b90613cd9828961287a565b52613ce4908761287a565b506002546001600160a01b031692613cfb90612057565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610c0957600096600097600094600092613ecc575b506000965b8651881015613e4c57613dd5600191613d9d88613d988761219a565b613920565b613db66060613dac8d8d61287a565b5101918251612242565b9052858a14613ddd575b6060613dcc8b8b61287a565b51015190613161565b970196613d7c565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613e1060808c01516001600160a01b031690565b1603613e1e575b5050613dc0565b613d98613e2a926121ca565b613e436060613e398d8d61287a565b5101918251613161565b90528b88613e17565b9796509750505050613e887f000000000000000000000000000000000000000000000000000000000000000091613d9863ffffffff84166121ca565b8411613e945750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b9298505050613ef491925060803d608011613f02575b613eec81836101fb565b8101906138f5565b919790939290919038613d77565b503d613ee2565b610ad36103e76104d66104d0613f22948a989698611d48565b926001600160a01b03600091515194169060e08801908151613f42610240565b6001600160a01b0385168152908260208301528260408301528260608301526080820152613f70878d61287a565b52613f7b868c61287a565b506040516301ffc9a760e01b8152633317103160e01b600482015291602083602481875afa8015610c09578f948c89968f96948d948f96889161425e575b50614157575b50505050505015613ffe575b61183f613fef613ff695613fe9602061183f613fe99760409761287a565b906137bf565b958b61287a565b90388a613bfc565b50506140849160608c61402f6104d66104d06140286103e76103e76002546001600160a01b031690565b938b611d48565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610c095761183f613fef604092613fe9602061183f8f8b90613ff69b613fe99a600080600092614119575b63ffffffff9293506141069060606140ce888861287a565b5101946140fb8a6140df8a8a61287a565b51019160406140ee8b8b61287a565b51019063ffffffff169052565b9063ffffffff169052565b1690529750975050505095505050613fcb565b50505063ffffffff6141456141069260603d606011614150575b61413d81836101fb565b810190613869565b9093509150826140b6565b503d614133565b8495985060a09697506141a1602061419760608261418e6104d06141876104d66104d08b6141da9c9d9e9f611d48565b998d611d48565b01359901612057565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613821565b03915afa8015610c09578592828c939181908294614222575b50614216906060614204888861287a565b5101926140fb60206140df8a8a61287a565b5288888f8c8138613fbf565b915050614216925061424c915060a03d60a011614257575b61424481836101fb565b8101906137d9565b9491929190506141f3565b503d61423a565b614277915060203d602011611ac157611ab381836101fb565b38613fb9565b6001600160a01b0360015416330361429157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916142c9815184613161565b9283156143ec5760005b8481106142e1575050505050565b818110156143d1576142f6611585828661287a565b6001600160a01b03811680156143a75761430f83613129565b878110614321575050506001016142d3565b84811015614384576001600160a01b0361433e611585838a61287a565b16821461434d5760010161430f565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036143a261158561439c88856138e8565b8961287a565b61433e565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6143e76115856143e184846138e8565b8561287a565b6142f6565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156120525760051b0190565b9081602091031261016a575190565b91608061022e9294936144858160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906144b78261025e565b6144c460405191826101fb565b828152601f19612029829461025e565b9081516020821080614554575b61452357036144ed5790565b6109d0906040519182917f3aeba390000000000000000000000000000000000000000000000000000000008352600483016102d7565b509060208101519081614535846121fc565b1c6144ed5750614544826144ad565b9160200360031b1b602082015290565b50602081146144e1565b918251601481029080820460141490151715611b325761458061458591613137565b613145565b9061459761459283613153565b6144ad565b9060146145a38361286d565b5360009260215b86518510156145d55760146001916145c5611585888b61287a565b60601b81870152019401936145aa565b919550936020935060601b90820152828152012090565b906145fc6103e760608401612057565b61460d600019936040810190611d48565b9050614685575b61461e82516138ac565b9260005b848110614630575050505050565b808260019214614680576060614646828761287a565b510151801561467a576146749061466e614660848961287a565b51516001600160a01b031690565b866156e9565b01614622565b50614674565b614674565b915061469181516138bb565b9161469f614660848461287a565b6040516301ffc9a760e01b8152633317103160e01b60048201526020816024816001600160a01b0386165afa908115610c0957600091614706575b506146e6575b50614614565b6147009060606146f6868661287a565b51015190836156e9565b386146e0565b61471f915060203d602011611ac157611ab381836101fb565b386146da565b60405190614732826101df565b60606020838281520152565b919060408382031261016a5760405190614757826101df565b8193805167ffffffffffffffff811161016a57826147769183016128f8565b835260208101519167ffffffffffffffff831161016a5760209261479a92016128f8565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a576102e8920161473e565b9060806001600160a01b03816147e4855160a0865260a08601906102b2565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b9060206102e89281815201906147c5565b919060408382031261016a57825167ffffffffffffffff811161016a5760209161485491850161473e565b92015190565b61ffff6148736102e895936060845260608401906147c5565b9316602082015260408184039101526102b2565b90919392959461489561278c565b5060208201805115614ce4576148b8610ad36103e785516001600160a01b031690565b946001600160a01b038616916040516301ffc9a760e01b81526020818061490660048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610c0957600091614cc5575b5015614c7c5761499288999a825192614931614725565b505161497d61494789516001600160a01b031690565b926040614952610240565b9e8f90815261496e8d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610c0957600091614c5d575b5015614b655750906000929183614a0c9899604051998a95869485937fb1c71c650000000000000000000000000000000000000000000000000000000085526004850161485a565b03925af18015610c0957600094600091614b1d575b50614af2614a5b9596614a95614a69614a87956115856020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b6101fb565b6040519586918683019190916001600160a01b036020820193169052565b03601f1981018652856101fb565b614ad1610919614ac7614ad78951614ad1610919614ac78c67ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b906144d4565b9767ffffffffffffffff166000526004602052604060002090565b93015193614afe61024f565b958652602086015260408501526060840152608083015260a082015290565b614a5b95506020915061158596614a95614a69614a8795614b53614af2953d806000833e614b4b81836101fb565b810190614829565b9b909b96505095505050969550614a21565b9793929061ffff16614c335751614c0957614bb46000939184926040519586809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614818565b03925af1908115610c0957614a5b95614a95614a69614af293611585602096614a8798600091614be6575b5099614a3d565b614c0391503d806000833e614bfb81836101fb565b81019061479f565b38614bdf565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614c76915060203d602011611ac157611ab381836101fb565b386149c4565b61055e614c9086516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b614cde915060203d602011611ac157611ab381836101fb565b3861491a565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592614de0947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b90614dfc6020928281519485920161028f565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b168152614e48825180936020898501910161028f565b019160f81b1683820152614e6682518093602060028501910161028f565b01019160f81b1683820152614e8582518093602060028501910161028f565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff000000000000000000000000000000000000000000000000000000000000006102e89f9e9c9860f81b168152614efa825180936020898501910161028f565b019160f01b1683820152614f1882518093602060038501910161028f565b01019160f01b1683820152614f36825180936020898501910161028f565b01019160f01b1660028201520190614de9565b60e081019060ff825151116151d057610100810160ff815151116151ba5761012082019260ff845151116151a457610140830160ff8151511161518e5761016084019061ffff8251511161517857610180850193600185515111615160576101a086019361ffff8551511161514a5760609551805161512e575b50865167ffffffffffffffff16602088015167ffffffffffffffff16906040890151614ff69067ffffffffffffffff1690565b9860608101516150099063ffffffff1690565b90608081015161501c9063ffffffff1690565b60a082015161ffff169160c00151926040519c8d96602088019661503f97614d0e565b03601f198101885261505190886101fb565b5190815161505f9060ff1690565b9051805160ff1698519081516150759060ff1690565b906040519a8b95602087019561508a96614e00565b03601f198101875261509c90876101fb565b519182516150aa9060ff1690565b9151805161ffff169480516150c09061ffff1690565b92519283516150d09061ffff1690565b9260405197889760208901976150e598614e8b565b03601f19810182526150f790826101fb565b6040519283926020840161510a91614de9565b61511391614de9565b61511c91614de9565b03601f19810182526102e890826101fb565b61514391965061513d9061286d565b51615917565b9438614fc3565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b60005261055e6024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b906151ef612cc1565b91601182106153ba57803563302326cb60e01b7fffffffff000000000000000000000000000000000000000000000000000000008216036153605750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61525581612002565b604086019081526152658261288e565b906060870191825260005b83811061531457505050506152d883836152ce6152c26152b86152b161529e6152e298876152ec9c9b615a3e565b6001600160a01b0390911660808d015290565b8585615ae2565b9291903691612755565b60a08a01528383615b31565b9491903691612755565b60c0880152615ae2565b9391903691612755565b60e084015281036152fb575090565b63d9437f9d60e01b600052600360045260245260446000fd5b8060019161535961534361533c61532f6153539a8d8d615a3e565b91906132c7868a5161287a565b8b8b615ae2565b9391889a919a51949a3691612755565b9261287a565b5201615270565b7f55a0e02c0000000000000000000000000000000000000000000000000000000060005263302326cb60e01b6004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b63d9437f9d60e01b6000526002600452602482905260446000fd5b80519060005b8281106153e757505050565b60018101808211611b32575b83811061540357506001016153db565b6001600160a01b03615415838561287a565b51166154276103e7611585848761287a565b14615434576001016153f3565b61055e615444611585848661287a565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b9081602091031261016a57516102e881612164565b936154cb60809461ffff6001600160a01b039567ffffffffffffffff6154d9969b9a9b16895216602088015260a0604088015260a0870190610c0e565b9085820360608701526102b2565b9416910152565b929190926154ec6135a3565b5061550b8167ffffffffffffffff166000526004602052604060002090565b805490959060e01c60ff16916080850192835161552e906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff1661554e916137bf565b9661555a90608d613161565b9460a087019586515161556c91613161565b9160ff169161557a83612212565b61558391613161565b9161558f906067613161565b61559891612242565b6155a191613161565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b0387160361562e5750505061ffff9250615620906156136000935b51956156066155f7610240565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b6156446103e7602094976001600160a01b031690565b9060406156558583015161ffff1690565b91015192615695875198604051998a96879586957ff23889580000000000000000000000000000000000000000000000000000000087526004870161548e565b03915afa908115610c09576156136156209261ffff956000916156ba575b50936155ea565b6156dc915060203d6020116156e2575b6156d481836101fb565b810190615479565b386156b3565b503d6156ca565b906001600160a01b036157ae9392604051938260208601947fa9059cbb00000000000000000000000000000000000000000000000000000000865216602486015260448501526044845261573e6064856101fb565b1660008060409384519561575286886101fb565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d156157d3573d61579f6157968261025e565b945194856101fb565b83523d6000602085013e615bcf565b8051806157b9575050565b816020806157ce9361022e9501019101612721565b615b5e565b60609250615bcf565b9660026158eb976158b860226102e89f9e9c9799600199859f9b7fff00000000000000000000000000000000000000000000000000000000000000906158b89f826158b89c6158bf9c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615869825180936020898501910161028f565b019160f81b168382015261588782518093602060238501910161028f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190614de9565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff82515111615a2857604081019160ff83515111615a1257606082019160ff835151116159fc57608081019260ff845151116159e65760a0820161ffff815151116159d0576102e8946159c29351945191615979835160ff1690565b975191615987835160ff1690565b945190615995825160ff1690565b9051936159a3855160ff1690565b9351966159b2885161ffff1690565b966040519c8d9b60208d016157dc565b03601f1981018352826101fb565b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b635a102da160e11b600052602660045260246000fd5b929190926001820191848311615ac95781013560001a828115615abe575060148103615a91578201938411615a7657013560601c9190565b63d9437f9d60e01b6000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b63d9437f9d60e01b600052600060045260245260446000fd5b91906002820191818311615ac9578381013560f01c0160020192818411615b1657918391615b0f93612d20565b9290929190565b63d9437f9d60e01b6000526001600452602483905260446000fd5b91906001820191818311615ac9578381013560001a0160010192818411615b1657918391615b0f93612d20565b15615b6557565b608460405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91929015615c305750815115615be3575090565b3b15615bec5790565b606460405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615c435750805190602001fd5b6109d09060405191829162461bcd60e51b8352600483016102d756fea164736f6c634300081a000a",
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
