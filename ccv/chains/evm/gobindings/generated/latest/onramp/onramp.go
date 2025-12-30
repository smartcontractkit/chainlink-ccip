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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61010060405234610380576040516160ec38819003601f8101601f191683016001600160401b03811184821017610385578392829160405283398101039060e0821261038057608082126103805761005561039b565b81519092906001600160401b03811681036103805783526020820151926001600160a01b03841684036103805760208101938452604083015163ffffffff81168103610380576040820190815260606100af8186016103ba565b83820190815293607f1901126103805760405192606084016001600160401b03811185821017610385576040526100e8608086016103ba565b845260a08501519485151586036103805760c061010c9160208701978852016103ba565b9560408501968752331561036f57600180546001600160a01b0319163317905583516001600160401b031615801561035d575b801561034b575b801561033c575b61030f5792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e05281511615801561032a575b8015610320575b61030f5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e0926000606061020d61039b565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff9283169116606061025861039b565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a1604051615d1d90816103cf8239608051818181610a5c015281816113b30152611c55015260a0518181816111710152611c81015260c051818181611cdc01526126b6015260e051818181611cad0152613e5f0152f35b6306b7c75960e31b60005260046000fd5b5081511515610192565b5082516001600160a01b03161561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761038557604052565b51906001600160a01b03821682036103805756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd5780632490769e146100f857806348a98aa4146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da57806390423fa2146100d5578063df0aa9e9146100d0578063e8d80861146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611bd7565b611b26565b611ab7565b6110cb565b610f47565b610f00565b610e5a565b610de7565b610d15565b610ad9565b610a94565b6105ad565b610365565b6102d7565b3461016a57600060031936011261016a576080610122611c1d565b61016860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b634e487b7160e01b600052604160045260246000fd5b610160810190811067ffffffffffffffff8211176101a257604052565b61016f565b6060810190811067ffffffffffffffff8211176101a257604052565b6080810190811067ffffffffffffffff8211176101a257604052565b6040810190811067ffffffffffffffff8211176101a257604052565b90601f601f19910116810190811067ffffffffffffffff8211176101a257604052565b6040519061022e6101c0836101fb565b565b6040519061022e610160836101fb565b6040519061022e60a0836101fb565b6040519061022e60c0836101fb565b67ffffffffffffffff81116101a257601f01601f191660200190565b604051906102896020836101fb565b60008252565b60005b8381106102a25750506000910152565b8181015183820152602001610292565b90601f19601f6020936102d08151809281875287808801910161028f565b0116010190565b3461016a57600060031936011261016a5761033660408051906102fa81836101fb565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102b2565b0390f35b67ffffffffffffffff81160361016a57565b359061022e8261033a565b908160a091031261016a5790565b3461016a57604060031936011261016a576004356103828161033a565b60243567ffffffffffffffff811161016a576103a2903690600401610357565b6103c08267ffffffffffffffff166000526004602052604060002090565b805491906001600160a01b036103df8185165b6001600160a01b031690565b161561051257906103369361048893926104256103ff6080850185611d04565b61040c6020870187611d04565b90501591826104f9575b61041f85611ebb565b86612e38565b93610474610431611fd5565b60408601906104408288611d37565b90506104a8575b6040880161046a81519260608b0193845161046460028b01611d6d565b916133e3565b9092525285611d37565b1515905061049a5760f01c90505b90613946565b60405190815292839250602083019150565b506001015461ffff16610482565b506104f46104c76104c26104bc848a611d37565b90612038565b612046565b60206104d66104bc858b611d37565b01356104e760208b015161ffff1690565b9060e08b0151928961319b565b610447565b91506105086040870187611d37565b9050151591610416565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156105a95781359367ffffffffffffffff85116105a657506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a576105bb3661054e565b906105c4614283565b6000915b8083106105d157005b6105dc838284612050565b926105e684612090565b67ffffffffffffffff81169081158015610a50575b8015610a3a575b8015610a21575b6109ea57856108609161087a6108708361086a61067b60e083019561066161065b61063489876120c7565b6106536106496101008a95949501809a6120c7565b94909236916120fd565b9236916120fd565b906142c1565b67ffffffffffffffff166000526004602052604060002090565b9687956106b661068d60208a01612046565b88906001600160a01b031673ffffffffffffffffffffffffffffffffffffffff19825416179055565b61085a61082360c060408b019a61071d6106cf8d6120a5565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b61077b61072c6080830161215f565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b6107be600161078c60a0840161215f565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b61081d8d6107ce60608401612169565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b016120bd565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d6120c7565b906003890161225b565b8a6120c7565b906002860161225b565b61012088019061088c6103d383612046565b156109c05761089d6108e792612046565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b61014087019061090b6109056108fd848b611d04565b9390506120a5565b60ff1690565b0361097c57956109627f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb2692610951610947600198999a85611d04565b9060048401612354565b5460a01c67ffffffffffffffff1690565b610971604051928392836124ee565b0390a20191906105c8565b6109869087611d04565b906109bc6040519283927f3aeba390000000000000000000000000000000000000000000000000000000008452600484016122fe565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a3360c088016120bd565b1615610609565b5060ff610a49604088016120a5565b1615610602565b5067ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001682146105fb565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610ab060043561033a565b610ac4602435610abf81610a83565b612671565b6040516001600160a01b039091168152602090f35b3461016a57610ae73661054e565b906001600160a01b03600354169160005b818110610b0157005b610b126103d36104c283858761441c565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610bf5576001948892600092610bc5575b5081610b79575b5050505001610af8565b81610ba97f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610bb9946158d2565b6040519081529081906020820190565b0390a338858180610b6f565b610be791925060203d8111610bee575b610bdf81836101fb565b81019061442c565b9038610b68565b503d610bd5565b612665565b906020808351928381520192019060005b818110610c185750505090565b82516001600160a01b0316845260209384019390920191600101610c0b565b90610d129160208152610c566020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015115156080820152608082015161ffff1660a082015260a082015161ffff1660c082015260c082015163ffffffff1660e082015260e08201516001600160a01b0316610100820152610140610cfc610ce6610100850151610160610120860152610180850190610bfa565b610120850151601f198583030184860152610bfa565b92015190610160601f19828503019101526102b2565b90565b3461016a57602060031936011261016a5767ffffffffffffffff600435610d3b8161033a565b6060610140604051610d4c81610185565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e082015282610100820152826101208201520152166000526004602052610336610da96040600020611ebb565b60405191829182610c37565b61022e9092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a57600060408051610e07816101a7565b8281528260208201520152610336604051610e21816101a7565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610db5565b3461016a57600060031936011261016a576000546001600160a01b0381163303610ed65773ffffffffffffffffffffffffffffffffffffffff19600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061022e82610a83565b8015150361016a57565b359061022e82610f32565b3461016a57606060031936011261016a576000604051610f66816101a7565b600435610f7281610a83565b8152602435610f8081610f32565b6020820190815260443590610f9482610a83565b60408301918252610fa3614283565b6001600160a01b03835116159182156110b8575b5081156110ad575b506110855780516002805460208401517fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160a01b039384161790151560a01b74ff00000000000000000000000000000000000000001617905560408201516003805473ffffffffffffffffffffffffffffffffffffffff1916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c90611070611c1d565b61107f6040519283928361443b565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610fbf565b516001600160a01b031615915038610fb7565b3461016a57608060031936011261016a576110e760043561033a565b60243567ffffffffffffffff811161016a57611107903690600401610357565b60443590611116606435610a83565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610bf557600091611a88575b50611a4b5760025460a01c60ff16611a21576111f8740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61121860043567ffffffffffffffff166000526004602052604060002090565b6001600160a01b0360643516156119f75780549061123f6103d36001600160a01b03841681565b33036119cd578383916112556080840184611d04565b949060208501956112668787611d04565b90501590816119b4575b61127985611ebb565b6112869390600435612e38565b93849160a01c67ffffffffffffffff1661129f90612725565b83547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617845592825163ffffffff16602084015161ffff166040805130602082015299908a90810103601f1981018b5261131d908b6101fb565b604080516064356001600160a01b031660208083019190915281529a90611344908c6101fb565b61134e8680611d04565b86549c9160e08e901c60ff1691369061136692612744565b9060ff16611373916144eb565b9060a08901519260408901611388908a611d37565b61139291506127bd565b9461139d908a611d04565b9690976113a861021e565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252600435811660208301529d909d1660408e0152600060608e015263ffffffff1660808d015261ffff1660a08c0152600060c08c015260e08b015261141c60048801611dfb565b6101008b01526101208a0152610140890152610160880152610180870152369061144592612744565b6101a0850152611453611fd5565b6114606040840184611d37565b905061190e575b906114af61148a85949360406114e0970151606087015161046460028701611d6d565b60608601528060408601526114a960808601516001600160a01b031690565b90614575565b60c08601526114bc61280c565b976114ca6040840184611d37565b151590506119005760f01c90505b600435613946565b63ffffffff90911660608401526020860193918452116118d657611505825186614603565b6115126040860186611d37565b90506117b6575b61152581959295614f60565b808552602081519101209061153e6040850151516128b1565b9460408101958652606060009401935b604086015180518210156117235760206115816103d36103d3611574866115c896612869565b516001600160a01b031690565b61158f8460608b0151612869565b519060405180809581947f958021a7000000000000000000000000000000000000000000000000000000008352600435600484016128fa565b03915afa8015610bf5576001600160a01b03916000916116f5575b5016801561169c57906000878b9387838861164661160f8860608f61160790612046565b980151612869565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612a45565b03925af18015610bf557816116749160019460009161167b575b508a519061166e8383612869565b52612869565b500161154e565b611696913d8091833e61168e81836101fb565b81019061295d565b38611660565b61054a6116b06115748460408b0151612869565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611716915060203d811161171c575b61170e81836101fb565b810190612650565b386115e3565b503d611704565b61033685808b867fb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd8d6117558d612046565b925193519051906117866040519283926001600160a01b03606435169767ffffffffffffffff600435169785612c10565b0390a4610ba97fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6118186118026117cc6104bc6040890189611d37565b60c08601805151156118bb5751905b602087015161ffff169060e088015192606435916117fd60043591369061282b565b61489e565b610180830151906118128261285c565b5261285c565b5061183a604061182e8451828701515190612869565b51015163ffffffff1690565b60a061184a61018084015161285c565b5101515163ffffffff82168111611862575050611519565b61054a925061187a6104c26104bc60408a018a611d37565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b506118d06118c98980611d04565b3691612744565b906117db565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff166114d8565b509350600191506040611922910187611d37565b90500361198a576114e0838688946114af61148a61197e61194c6104c26104bc6040880188611d37565b602061195e6104bc6040890189611d37565b013561196f602089015161ffff1690565b9060e08901519260043561319b565b92939495505050611467565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b90506119c36040870187611d37565b9050151590611270565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005261054a6004359067ffffffffffffffff60249216600452565b611aaa915060203d602011611ab0575b611aa281836101fb565b810190612710565b386111a2565b503d611a98565b3461016a57602060031936011261016a5767ffffffffffffffff600435611add8161033a565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611b215760405167ffffffffffffffff9091168152602090f35b612173565b3461016a57602060031936011261016a576001600160a01b03600435611b4b81610a83565b611b53614283565b16338114611bad578073ffffffffffffffffffffffffffffffffffffffff1960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611bf360043561033a565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611c2d816101c3565b8281528260208201528260408201520152604051611c4a816101c3565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611d9f57505061022e925003836101fb565b84546001600160a01b0316835260019485019487945060209093019201611d8a565b90600182811c92168015611df1575b6020831014611ddb57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611dd0565b9060405191826000825492611e0f84611dc1565b8084529360018116908115611e7b5750600114611e34575b5061022e925003836101fb565b90506000929192526020600020906000915b818310611e5f57505090602061022e9282010138611e27565b6020919350806001915483858901015201910190918492611e46565b6020935061022e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611e27565b90611fb56004611ec9610230565b93611f38611f2d8254611ef2611ee5826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c166040890152611f2760e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b611f8b611f7b6001830154611f5c611f518261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b611f9760028201611d6d565b610100860152611fa960038201611d6d565b61012086015201611dfb565b610140830152565b67ffffffffffffffff81116101a25760051b60200190565b60405190611fe46020836101fb565b6000808352366020840137565b90611ffb82611fbd565b61200860405191826101fb565b828152601f196120188294611fbd565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b90156120415790565b612022565b35610d1281610a83565b91908110156120415760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561016a570190565b35610d128161033a565b60ff81160361016a57565b35610d128161209a565b63ffffffff81160361016a57565b35610d12816120af565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b92919061210981611fbd565b9361211760405195866101fb565b602085838152019160051b810192831161016a57905b82821061213957505050565b60208091833561214881610a83565b81520191019061212d565b61ffff81160361016a57565b35610d1281612153565b35610d1281610f32565b634e487b7160e01b600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611b2157565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611b2157565b908160031b9180830460081490151715611b2157565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611b2157565b81810292918115918404141715611b2157565b81811061224f575050565b60008155600101612244565b9067ffffffffffffffff83116101a2576801000000000000000083116101a25781548383558084106122bf575b5090600052602060002060005b8381106122a25750505050565b60019060208435946122b386610a83565b01938184015501612295565b6122d790836000528460206000209182019101612244565b38612288565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610d129381815201916122dd565b9190601f811161231e57505050565b61022e926000526020600020906020601f840160051c8301931061234a575b601f0160051c0190612244565b909150819061233d565b90929167ffffffffffffffff81116101a25761237a816123748454611dc1565b8461230f565b6000601f82116001146123ba5781906123ab9394956000926123af575b50506000198260011b9260031b1c19161790565b9055565b013590503880612397565b601f198216946123cf84600052602060002090565b91805b87811061240a5750836001959697106123f0575b505050811b019055565b60001960f88560031b161c199101351690553880806123e6565b909260206001819286860135815501940191016123d2565b359061022e8261209a565b359061022e82612153565b359061022e826120af565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b8181106124925750505090565b9091926020806001926001600160a01b0387356124ae81610a83565b168152019401929101612485565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b67ffffffffffffffff610d12939216815260406020820152612524604082016125168461034c565b67ffffffffffffffff169052565b61254361253360208401610f27565b6001600160a01b03166060830152565b61255c61255260408401612422565b60ff166080830152565b61257461256b60608401610f3c565b151560a0830152565b61258e6125836080840161242d565b61ffff1660c0830152565b6125a861259d60a0840161242d565b61ffff1660e0830152565b6125c56125b760c08401612438565b63ffffffff16610100830152565b61263d6126106125ef6125db60e0860186612443565b6101606101208701526101a0860191612478565b6125fd610100860186612443565b90603f1986840301610140870152612478565b926126326126216101208301610f27565b6001600160a01b0316610160850152565b6101408101906124bc565b91610180603f19828603019101526122dd565b9081602091031261016a5751610d1281610a83565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610bf5576001600160a01b03916000916126f357501690565b61270c915060203d60201161171c5761170e81836101fb565b1690565b9081602091031261016a5751610d1281610f32565b67ffffffffffffffff1667ffffffffffffffff8114611b215760010190565b9291926127508261025e565b9161275e60405193846101fb565b82948184528183011161016a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101a257604052606060a0836000815282602082015282604082015282808201528260808201520152565b906127c782611fbd565b6127d460405191826101fb565b828152601f196127e48294611fbd565b019060005b8281106127f557505050565b60209061280061277b565b828285010152016127e9565b60405190612819826101a7565b60606040838281528260208201520152565b919082604091031261016a57604051612843816101df565b6020808294803561285381610a83565b84520135910152565b8051156120415760200190565b80518210156120415760209160051b010190565b6040519061288c6020836101fb565b600080835282815b8281106128a057505050565b806060602080938501015201612894565b906128bb82611fbd565b6128c860405191826101fb565b828152601f196128d88294611fbd565b019060005b8281106128e957505050565b8060606020809385010152016128dd565b60409067ffffffffffffffff610d12949316815281602082015201906102b2565b81601f8201121561016a5780516129318161025e565b9261293f60405194856101fb565b8184526020828401011161016a57610d12916020808501910161028f565b9060208282031261016a57815167ffffffffffffffff811161016a57610d12920161291b565b9080602083519182815201916020808360051b8301019401926000915b8383106129af57505050505090565b9091929394602080612a3683601f1986600196030187528951908151815260a0612a25612a13612a016129ef8887015160c08a88015260c08701906102b2565b604087015186820360408801526102b2565b606086015185820360608701526102b2565b608085015184820360808601526102b2565b9201519060a08184039101526102b2565b970193019301919392906129a0565b919390610d129593612b8d612ba59260a08652612a6f60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612b78612b60612b48612b30612b18612b028c61026060e08a0151916101c061018082015201906102b2565b6101008801518d8203609f1901888f01526102b2565b6101208701518c8203609f19016101c08e01526102b2565b610140860151609f198c8303016101e08d01526102b2565b610160850151609f198b8303016102008c01526102b2565b610180840151609f198a8303016102208b0152612983565b910151609f19878303016102408801526102b2565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102b2565b9080602083519182815201916020808360051b8301019401926000915b838310612be357505050505090565b9091929394602080612c0183601f19866001960301875289516102b2565b97019301930191939290612bd4565b9493916001600160a01b03612c33921686526080602087015260808601906102b2565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612c755750505050610d129394506060818403910152612bb7565b90919295602080612cd683601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102b2565b980192019201909291612c57565b60405190610100820182811067ffffffffffffffff8211176101a257604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612d8f575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a5782612de991850161291b565b926020810151612df8816120af565b92604082015167ffffffffffffffff811161016a57610d12920161291b565b60409067ffffffffffffffff610d12959316815281602082015201916122dd565b91929092612e44612ce4565b60048310158061300d575b15612f57575090612e5f9161521d565b92612e6d604085015161540c565b80612f3c575b60408401612e9081519260608701938451610120880151916154b0565b9092525260c083015151612eec575b50608082016001600160a01b03612ebd82516001600160a01b031690565b1615612ec857505090565b612edf60e0610d129301516001600160a01b031690565b6001600160a01b03169052565b612f00612efc6060840151151590565b1590565b15612e9f577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff612f50845163ffffffff1690565b1615612e73565b9491600090612faf92612f786103d36103d36002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612e17565b03915afa8015610bf557600090600090600090612fe1575b60a088015263ffffffff16865290505b60c0850152612e6d565b505050613003612fd7913d806000833e612ffb81836101fb565b810190612dc1565b9192508291612fc7565b5063302326cb60e01b7fffffffff0000000000000000000000000000000000000000000000000000000061304a6130448686612d35565b90612d5b565b1614612e4f565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a57815161308581611fbd565b9261309360405194856101fb565b81845260208085019260051b82010192831161016a57602001905b8282106130bb5750505090565b6020809183516130ca81610a83565b8152019101906130ae565b95949060009460a09467ffffffffffffffff61311c956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102b2565b930152565b9060028201809211611b2157565b9060018201809211611b2157565b6001019081600111611b2157565b9060148201809211611b2157565b90600c8201809211611b2157565b91908201809211611b2157565b6000198114611b215760010190565b80548210156120415760005260206000200190600090565b929394919060036131c08567ffffffffffffffff166000526004602052604060002090565b01936001600160a01b036131d5818416612671565b169182156133ad576040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bf55760009161338e575b501561337e57613254600095969798604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016130d5565b03915afa928315610bf557600093613359575b5082511561334e5782519061328661328184548094613167565b611ff1565b906000928394845b87518110156132ed576132a4611574828a612869565b6001600160a01b038116156132e157906132db6001926132cd6132c68a613174565b9989612869565b906001600160a01b03169052565b0161328e565b509550600180966132db565b5091955091936132ff575b5050815290565b60005b82811061330f57506132f8565b8061334861333561332260019486613183565b90546001600160a01b039160031b1c1690565b6132cd61334188613174565b9789612869565b01613302565b9150610d1290611d6d565b6133779193503d806000833e61336f81836101fb565b810190613051565b9138613267565b505050509250610d129150611d6d565b6133a7915060203d602011611ab057611aa281836101fb565b3861320f565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b939192936133ff6133f78251865190613167565b865190613167565b9061341261340c83611ff1565b926128b1565b94600096875b8351891015613478578861346e61346160019361344961343f6115748e9f9d9e9d8b612869565b6132cd838c612869565b613467613456858c612869565b519180938491613174565b9c612869565b528b612869565b5001979695613418565b959250929350955060005b865181101561350c576134996115748289612869565b60006001600160a01b038216815b8881106134e0575b50509060019291156134c3575b5001613483565b6134da906132cd6134d389613174565b9888612869565b386134bc565b816134f16103d3611574848c612869565b146134fe576001016134a7565b5060019150819050386134af565b509390945060005b855181101561359d5761352a6115748288612869565b60006001600160a01b038216815b878110613571575b5050906001929115613554575b5001613514565b61356b906132cd61356488613174565b9787612869565b3861354d565b816135826103d3611574848b612869565b1461358f57600101613538565b506001915081905038613540565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101a25760405260606080836000815260006020820152600060408201526000838201520152565b906135f282611fbd565b6135ff60405191826101fb565b828152601f1961360f8294611fbd565b019060005b82811061362057505050565b60209061362b6135a9565b82828501015201613614565b9081606091031261016a57805161364d81612153565b916040602083015161365e816120af565b920151610d12816120af565b9160209082815201919060005b8181106136845750505090565b9091926040806001926001600160a01b0387356136a081610a83565b16815260208781013590820152019401929101613677565b949391929067ffffffffffffffff168552608060208601526137116136f26136e085806124bc565b60a060808a01526101208901916122dd565b6136ff60208601866124bc565b90607f198984030160a08a01526122dd565b6040840135601e198536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a5761022e95613799613770836060976137ba978d60c0607f19826137ac9a030191015261366a565b9161378f61377f888301610f27565b6001600160a01b031660e08d0152565b60808101906124bc565b90607f198b8403016101008c01526122dd565b9087820360408901526102b2565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611b2157565b908160a091031261016a5780519160208201516137fb816120af565b91604081015161380a816120af565b916080606083015161381b81612153565b920151610d1281610f32565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610d129b9a9616885216602087015260408601521660608401521660808201528160a082015201906102b2565b9081606091031261016a57805161364d816120af565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611b2157565b906000198201918211611b2157565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611b2157565b91908203918211611b2157565b919082608091031261016a578151613912816120af565b916020810151916060604083015192015190565b8115613930570490565b634e487b7160e01b600052601260045260246000fd5b938294600090600095604081019461398061397b61397688515161396e60408c01809c611d37565b919050613167565b613121565b6135e8565b9660009586955b88518051881015613be4576103d36103d36115748a6139a594612869565b6139f2602060608801926139ba8b8551612869565b51908a6040518095819482937f958021a7000000000000000000000000000000000000000000000000000000008452600484016128fa565b03915afa8015610bf5576001600160a01b0391600091613bc6575b50168015613b72579060608e9392613a268b8451612869565b5190613a3760208b015161ffff1690565b958b613a72604051988995869485947f80485e25000000000000000000000000000000000000000000000000000000008652600486016136b8565b03915afa8015610bf557600193613b0f938b8f8f95600080958197613b18575b509083929161ffff613aba85613ab3611574613b0399613b099d9e51612869565b9451612869565b5191613ad6613ac7610240565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261166e8383612869565b506137c5565b996137c5565b96019596613987565b613b099750611574965084939291509361ffff613aba82613ab3613b55613b039960603d8111613b6b575b613b4d81836101fb565b810190613637565b9c9196909c9d5050505050505090919293613a92565b503d613b43565b61054a88613b846115748c8f51612869565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613bde915060203d811161171c5761170e81836101fb565b38613a0d565b50919a9496929395509897968a613bfb8187611d37565b9050613f0f575b50508651613c0f90613885565b99613c1d6020860186611d04565b91613c29915086611d37565b9560609150019486613c3a87612046565b91613c45938a6156c9565b613c4f8b89612869565b52613c5a8a88612869565b50613c658a88612869565b516020015163ffffffff16613c79916137c5565b90613c848a88612869565b516040015163ffffffff16613c98916137c5565b91613ca1610240565b33815290600060208301819052604083015261ffff166060820152613cc461027a565b60808201528651613cd4906138b2565b90613cdf8289612869565b52613cea9087612869565b506002546001600160a01b031692613d0190612046565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610bf557600096600097600094600092613ed2575b506000965b8651881015613e5257613ddb600191613da388613d9e87612189565b613926565b613dbc6060613db28d8d612869565b5101918251612231565b9052858a14613de3575b6060613dd28b8b612869565b51015190613167565b970196613d82565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613e1660808c01516001600160a01b031690565b1603613e24575b5050613dc6565b613d9e613e30926121b9565b613e496060613e3f8d8d612869565b5101918251613167565b90528b88613e1d565b9796509750505050613e8e7f000000000000000000000000000000000000000000000000000000000000000091613d9e63ffffffff84166121b9565b8411613e9a5750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b9298505050613efa91925060803d608011613f08575b613ef281836101fb565b8101906138fb565b919790939290919038613d7d565b503d613ee8565b610abf6103d36104c26104bc613f28948a989698611d37565b926001600160a01b03600091515194169060e08801908151613f48610240565b6001600160a01b0385168152908260208301528260408301528260608301526080820152613f76878d612869565b52613f81868c612869565b506040516301ffc9a760e01b8152633317103160e01b600482015291602083602481875afa8015610bf5578f948c89968f96948d948f968891614264575b5061415d575b50505050505015614004575b61182e613ff5613ffc95613fef602061182e613fef97604097612869565b906137c5565b958b612869565b90388a613c02565b505061408a9160608c6140356104c26104bc61402e6103d36103d36002546001600160a01b031690565b938b611d37565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610bf55761182e613ff5604092613fef602061182e8f8b90613ffc9b613fef9a60008060009261411f575b63ffffffff92935061410c9060606140d48888612869565b5101946141018a6140e58a8a612869565b51019160406140f48b8b612869565b51019063ffffffff169052565b9063ffffffff169052565b1690529750975050505095505050613fd1565b50505063ffffffff61414b61410c9260603d606011614156575b61414381836101fb565b81019061386f565b9093509150826140bc565b503d614139565b8495985060a09697506141a7602061419d6060826141946104bc61418d6104c26104bc8b6141e09c9d9e9f611d37565b998d611d37565b01359901612046565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613827565b03915afa8015610bf5578592828c939181908294614228575b5061421c90606061420a8888612869565b51019261410160206140e58a8a612869565b5288888f8c8138613fc5565b91505061421c9250614252915060a03d60a01161425d575b61424a81836101fb565b8101906137df565b9491929190506141f9565b503d614240565b61427d915060203d602011611ab057611aa281836101fb565b38613fbf565b6001600160a01b0360015416330361429757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916142cf815184613167565b9283156143f25760005b8481106142e7575050505050565b818110156143d7576142fc6115748286612869565b6001600160a01b03811680156143ad576143158361312f565b878110614327575050506001016142d9565b8481101561438a576001600160a01b03614344611574838a612869565b16821461435357600101614315565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036143a86115746143a288856138ee565b89612869565b614344565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6143ed6115746143e784846138ee565b85612869565b6142fc565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156120415760051b0190565b9081602091031261016a575190565b91608061022e92949361448b8160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610d129281815201906102b2565b906144ce8261025e565b6144db60405191826101fb565b828152601f19612018829461025e565b908151602082108061456b575b61453a57036145045790565b6109bc906040519182917f3aeba390000000000000000000000000000000000000000000000000000000008352600483016144b3565b50906020810151908161454c846121eb565b1c614504575061455b826144c4565b9160200360031b1b602082015290565b50602081146144f8565b918251601481029080820460141490151715611b215761459761459c9161313d565b61314b565b906145ae6145a983613159565b6144c4565b9060146145ba8361285c565b5360009260215b86518510156145ec5760146001916145dc611574888b612869565b60601b81870152019401936145c1565b919550936020935060601b90820152828152012090565b906146136103d360608401612046565b614624600019936040810190611d37565b905061469c575b61463582516138b2565b9260005b848110614647575050505050565b80826001921461469757606061465d8287612869565b51015180156146915761468b906146856146778489612869565b51516001600160a01b031690565b866158d2565b01614639565b5061468b565b61468b565b91506146a881516138c1565b916146b66146778484612869565b6040516301ffc9a760e01b8152633317103160e01b60048201526020816024816001600160a01b0386165afa908115610bf55760009161471d575b506146fd575b5061462b565b61471790606061470d8686612869565b51015190836158d2565b386146f7565b614736915060203d602011611ab057611aa281836101fb565b386146f1565b60405190614749826101df565b60606020838281520152565b919060408382031261016a576040519061476e826101df565b8193805167ffffffffffffffff811161016a578261478d91830161291b565b835260208101519167ffffffffffffffff831161016a576020926147b1920161291b565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a57610d129201614755565b9060806001600160a01b03816147fb855160a0865260a08601906102b2565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610d129281815201906147dc565b919060408382031261016a57825167ffffffffffffffff811161016a5760209161486b918501614755565b92015190565b61ffff61488a610d1295936060845260608401906147dc565b9316602082015260408184039101526102b2565b9091939295946148ac61277b565b5060208201805115614cfb576148cf610abf6103d385516001600160a01b031690565b946001600160a01b038616916040516301ffc9a760e01b81526020818061491d60048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610bf557600091614cdc575b5015614c93576149a988999a82519261494861473c565b505161499461495e89516001600160a01b031690565b926040614969610240565b9e8f9081526149858d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bf557600091614c74575b5015614b7c5750906000929183614a239899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614871565b03925af18015610bf557600094600091614b34575b50614b09614a729596614aac614a80614a9e956115746020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b6101fb565b6040519586918683019190916001600160a01b036020820193169052565b03601f1981018652856101fb565b614ae8610905614ade614aee8951614ae8610905614ade8c67ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b906144eb565b9767ffffffffffffffff166000526004602052604060002090565b93015193614b1561024f565b958652602086015260408501526060840152608083015260a082015290565b614a7295506020915061157496614aac614a80614a9e95614b6a614b09953d806000833e614b6281836101fb565b810190614840565b9b909b96505095505050969550614a38565b9793929061ffff16614c4a5751614c2057614bcb6000939184926040519586809481937f9a4575b90000000000000000000000000000000000000000000000000000000083526004830161482f565b03925af1908115610bf557614a7295614aac614a80614b0993611574602096614a9e98600091614bfd575b5099614a54565b614c1a91503d806000833e614c1281836101fb565b8101906147b6565b38614bf6565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614c8d915060203d602011611ab057611aa281836101fb565b386149db565b61054a614ca786516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b614cf5915060203d602011611ab057611aa281836101fb565b38614931565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592614df7947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b90614e136020928281519485920161028f565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b168152614e5f825180936020898501910161028f565b019160f81b1683820152614e7d82518093602060028501910161028f565b01019160f81b1683820152614e9c82518093602060028501910161028f565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff00000000000000000000000000000000000000000000000000000000000000610d129f9e9c9860f81b168152614f11825180936020898501910161028f565b019160f01b1683820152614f2f82518093602060038501910161028f565b01019160f01b1683820152614f4d825180936020898501910161028f565b01019160f01b1660028201520190614e00565b60e081019060ff8251511161520757610100810160ff815151116151f15761012082019260ff845151116151db57610140830160ff815151116151c55761016084019061ffff825151116151af57610180850193600185515111615197576101a086019361ffff8551511161518157606095518051615145575b50865167ffffffffffffffff16602088015167ffffffffffffffff1690604089015161500d9067ffffffffffffffff1690565b9860608101516150209063ffffffff1690565b9060808101516150339063ffffffff1690565b60a082015161ffff169160c00151926040519c8d96602088019661505697614d25565b03601f198101885261506890886101fb565b519081516150769060ff1690565b9051805160ff16985190815161508c9060ff1690565b906040519a8b9560208701956150a196614e17565b03601f19810187526150b390876101fb565b519182516150c19060ff1690565b9151805161ffff169480516150d79061ffff1690565b92519283516150e79061ffff1690565b9260405197889760208901976150fc98614ea2565b03601f198101825261510e90826101fb565b6040519283926020840161512191614e00565b61512a91614e00565b61513391614e00565b03601f1981018252610d1290826101fb565b61515a9196506151549061285c565b51615ac9565b9461ffff86511161516b5738614fda565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b600052602660045260246000fd5b635a102da160e11b60005261054a6024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b90615226612ce4565b91601182106153f157803563302326cb60e01b7fffffffff000000000000000000000000000000000000000000000000000000008216036153975750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61528c81611ff1565b6040860190815261529c826128b1565b906060870191825260005b83811061534b575050505061530f83836153056152f96152ef6152e86152d561531998876153239c9b615bf0565b6001600160a01b0390911660808d015290565b8585615c94565b9291903691612744565b60a08a01528383615ce3565b9491903691612744565b60c0880152615c94565b9391903691612744565b60e08401528103615332575090565b63d9437f9d60e01b600052600360045260245260446000fd5b8060019161539061537a61537361536661538a9a8d8d615bf0565b91906132cd868a51612869565b8b8b615c94565b9391889a919a51949a3691612744565b92612869565b52016152a7565b7f55a0e02c0000000000000000000000000000000000000000000000000000000060005263302326cb60e01b6004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b63d9437f9d60e01b6000526002600452602482905260446000fd5b80519060005b82811061541e57505050565b60018101808211611b21575b83811061543a5750600101615412565b6001600160a01b0361544c8385612869565b511661545e6103d36115748487612869565b1461546b5760010161542a565b61054a61547b6115748486612869565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b9093919281511561563c57506154c58161540c565b80519260005b8481106154da57509093925050565b6154ea6103d36115748386612869565b156154f7576001016154cb565b939194615511613281615509856138b2565b845190613167565b9261552e615529615521836138b2565b855190613167565b6128b1565b968792600097885b8481106155de5750505050505060005b81518110156155d1576000805b868110615592575b509060019161558d576155876155746115748386612869565b6132cd61558089613174565b9887612869565b01615546565b615587565b61559f6115748287612869565b6001600160a01b036155b76103d36115748789612869565b9116146155c657600101615553565b50600190508061555b565b5050909180825283529190565b90919293949882821461563257906156246156178361560a8b6132cd60019761538a611574898e612869565b61561d6134568589612869565b9e612869565b528c612869565b505b01908994939291615536565b9850600190615626565b9193505015615657575061564e611fd5565b90610d1261287d565b90610d1282516128b1565b9081602091031261016a5751610d1281612153565b936156b460809461ffff6001600160a01b039567ffffffffffffffff6156c2969b9a9b16895216602088015260a0604088015260a0870190610bfa565b9085820360608701526102b2565b9416910152565b929190926156d56135a9565b506156f48167ffffffffffffffff166000526004602052604060002090565b805490959060e01c60ff169160808501928351615717906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff16615737916137c5565b9661574390608d613167565b9460a087019586515161575591613167565b9160ff169161576383612201565b61576c91613167565b91615778906067613167565b61578191612231565b61578a91613167565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b038716036158175750505061ffff9250615809906157fc6000935b51956157ef6157e0610240565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b61582d6103d3602094976001600160a01b031690565b90604061583e8583015161ffff1690565b9101519261587e875198604051998a96879586957ff238895800000000000000000000000000000000000000000000000000000000875260048701615677565b03915afa908115610bf5576157fc6158099261ffff956000916158a3575b50936157d3565b6158c5915060203d6020116158cb575b6158bd81836101fb565b810190615662565b3861589c565b503d6158b3565b91602091600091604051906001600160a01b03858301937fa9059cbb0000000000000000000000000000000000000000000000000000000085521660248301526044820152604481526159266064826101fb565b519082855af115612665576000513d61598557506001600160a01b0381163b155b61594e5750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615947565b966002615a9d97615a6a6022610d129f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090615a6a9f82615a6a9c615a719c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615a1b825180936020898501910161028f565b019160f81b1683820152615a3982518093602060238501910161028f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190614e00565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff82515111615bda57604081019160ff83515111615bc457606082019160ff83515111615bae57608081019260ff84515111615b985760a0820161ffff81515111615b8257610d1294615b749351945191615b2b835160ff1690565b975191615b39835160ff1690565b945190615b47825160ff1690565b905193615b55855160ff1690565b935196615b64885161ffff1690565b966040519c8d9b60208d0161598e565b03601f1981018352826101fb565b635a102da160e11b600052602b60045260246000fd5b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b929190926001820191848311615c7b5781013560001a828115615c70575060148103615c43578201938411615c2857013560601c9190565b63d9437f9d60e01b6000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b63d9437f9d60e01b600052600060045260245260446000fd5b91906002820191818311615c7b578381013560f01c0160020192818411615cc857918391615cc193612d43565b9290929190565b63d9437f9d60e01b6000526001600452602483905260446000fd5b91906001820191818311615c7b578381013560001a0160010192818411615cc857918391615cc193612d4356fea164736f6c634300081a000a",
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
