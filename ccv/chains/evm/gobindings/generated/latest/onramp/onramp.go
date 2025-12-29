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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61010060405234610380576040516160f238819003601f8101601f191683016001600160401b03811184821017610385578392829160405283398101039060e0821261038057608082126103805761005561039b565b81519092906001600160401b03811681036103805783526020820151926001600160a01b03841684036103805760208101938452604083015163ffffffff81168103610380576040820190815260606100af8186016103ba565b83820190815293607f1901126103805760405192606084016001600160401b03811185821017610385576040526100e8608086016103ba565b845260a08501519485151586036103805760c061010c9160208701978852016103ba565b9560408501968752331561036f57600180546001600160a01b0319163317905583516001600160401b031615801561035d575b801561034b575b801561033c575b61030f5792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e05281511615801561032a575b8015610320575b61030f5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e0926000606061020d61039b565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff9283169116606061025861039b565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a1604051615d2390816103cf8239608051818181610a5c015281816113e30152611c85015260a0518181816111a10152611cb1015260c051818181611d0c01526126e6015260e051818181611cdd0152613e8f0152f35b6306b7c75960e31b60005260046000fd5b5081511515610192565b5082516001600160a01b03161561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761038557604052565b51906001600160a01b03821682036103805756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd5780632490769e146100f857806348a98aa4146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da57806390423fa2146100d5578063df0aa9e9146100d0578063e8d80861146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611c07565b611b56565b611ae7565b6110fb565b610f77565b610f30565b610e8a565b610e17565b610d45565b610ad9565b610a94565b6105ad565b610365565b6102d7565b3461016a57600060031936011261016a576080610122611c4d565b61016860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b634e487b7160e01b600052604160045260246000fd5b610160810190811067ffffffffffffffff8211176101a257604052565b61016f565b6060810190811067ffffffffffffffff8211176101a257604052565b6080810190811067ffffffffffffffff8211176101a257604052565b6040810190811067ffffffffffffffff8211176101a257604052565b90601f601f19910116810190811067ffffffffffffffff8211176101a257604052565b6040519061022e6101c0836101fb565b565b6040519061022e610160836101fb565b6040519061022e60a0836101fb565b6040519061022e60c0836101fb565b67ffffffffffffffff81116101a257601f01601f191660200190565b604051906102896020836101fb565b60008252565b60005b8381106102a25750506000910152565b8181015183820152602001610292565b90601f19601f6020936102d08151809281875287808801910161028f565b0116010190565b3461016a57600060031936011261016a5761033660408051906102fa81836101fb565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102b2565b0390f35b67ffffffffffffffff81160361016a57565b359061022e8261033a565b908160a091031261016a5790565b3461016a57604060031936011261016a576004356103828161033a565b60243567ffffffffffffffff811161016a576103a2903690600401610357565b6103c08267ffffffffffffffff166000526004602052604060002090565b805491906001600160a01b036103df8185165b6001600160a01b031690565b161561051257906103369361048893926104256103ff6080850185611d34565b61040c6020870187611d34565b90501591826104f9575b61041f85611eeb565b86612e68565b93610474610431612005565b60408601906104408288611d67565b90506104a8575b6040880161046a81519260608b0193845161046460028b01611d9d565b91613413565b9092525285611d67565b1515905061049a5760f01c90505b90613976565b60405190815292839250602083019150565b506001015461ffff16610482565b506104f46104c76104c26104bc848a611d67565b90612068565b612076565b60206104d66104bc858b611d67565b01356104e760208b015161ffff1690565b9060e08b015192896131cb565b610447565b91506105086040870187611d67565b9050151591610416565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156105a95781359367ffffffffffffffff85116105a657506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a576105bb3661054e565b906105c46142b3565b6000915b8083106105d157005b6105dc838284612080565b926105e6846120c0565b67ffffffffffffffff81169081158015610a50575b8015610a3a575b8015610a21575b6109ea57856108609161087a6108708361086a61067b60e083019561066161065b61063489876120f7565b6106536106496101008a95949501809a6120f7565b949092369161212d565b92369161212d565b906142f1565b67ffffffffffffffff166000526004602052604060002090565b9687956106b661068d60208a01612076565b88906001600160a01b031673ffffffffffffffffffffffffffffffffffffffff19825416179055565b61085a61082360c060408b019a61071d6106cf8d6120d5565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b61077b61072c6080830161218f565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b6107be600161078c60a0840161218f565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b61081d8d6107ce60608401612199565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b016120ed565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d6120f7565b906003890161228b565b8a6120f7565b906002860161228b565b61012088019061088c6103d383612076565b156109c05761089d6108e792612076565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b61014087019061090b6109056108fd848b611d34565b9390506120d5565b60ff1690565b0361097c57956109627f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb2692610951610947600198999a85611d34565b9060048401612384565b5460a01c67ffffffffffffffff1690565b6109716040519283928361251e565b0390a20191906105c8565b6109869087611d34565b906109bc6040519283927f3aeba3900000000000000000000000000000000000000000000000000000000084526004840161232e565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a3360c088016120ed565b1615610609565b5060ff610a49604088016120d5565b1615610602565b5067ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001682146105fb565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610ab060043561033a565b610ac4602435610abf81610a83565b6126a1565b6040516001600160a01b039091168152602090f35b3461016a57610ae73661054e565b906001600160a01b0360035416918215610c005760005b818110610b0757005b610b186103d36104c2838587614422565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610bfb576001948892600092610bcb575b5081610b7f575b5050505001610afe565b81610baf7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610bbf946158d8565b6040519081529081906020820190565b0390a338858180610b75565b610bed91925060203d8111610bf4575b610be581836101fb565b810190614432565b9038610b6e565b503d610bdb565b612695565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b906020808351928381520192019060005b818110610c485750505090565b82516001600160a01b0316845260209384019390920191600101610c3b565b90610d429160208152610c866020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015115156080820152608082015161ffff1660a082015260a082015161ffff1660c082015260c082015163ffffffff1660e082015260e08201516001600160a01b0316610100820152610140610d2c610d16610100850151610160610120860152610180850190610c2a565b610120850151601f198583030184860152610c2a565b92015190610160601f19828503019101526102b2565b90565b3461016a57602060031936011261016a5767ffffffffffffffff600435610d6b8161033a565b6060610140604051610d7c81610185565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e082015282610100820152826101208201520152166000526004602052610336610dd96040600020611eeb565b60405191829182610c67565b61022e9092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a57600060408051610e37816101a7565b8281528260208201520152610336604051610e51816101a7565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610de5565b3461016a57600060031936011261016a576000546001600160a01b0381163303610f065773ffffffffffffffffffffffffffffffffffffffff19600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061022e82610a83565b8015150361016a57565b359061022e82610f62565b3461016a57606060031936011261016a576000604051610f96816101a7565b600435610fa281610a83565b8152602435610fb081610f62565b6020820190815260443590610fc482610a83565b60408301918252610fd36142b3565b6001600160a01b03835116159182156110e8575b5081156110dd575b506110b55780516002805460208401517fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160a01b039384161790151560a01b74ff00000000000000000000000000000000000000001617905560408201516003805473ffffffffffffffffffffffffffffffffffffffff1916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c906110a0611c4d565b6110af60405192839283614441565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610fef565b516001600160a01b031615915038610fe7565b3461016a57608060031936011261016a5761111760043561033a565b60243567ffffffffffffffff811161016a57611137903690600401610357565b60443590611146606435610a83565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610bfb57600091611ab8575b50611a7b5760025460a01c60ff16611a5157611228740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b61124860043567ffffffffffffffff166000526004602052604060002090565b6001600160a01b036064351615611a275780549061126f6103d36001600160a01b03841681565b33036119fd578383916112856080840184611d34565b949060208501956112968787611d34565b90501590816119e4575b6112a985611eeb565b6112b69390600435612e68565b93849160a01c67ffffffffffffffff166112cf90612755565b83547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff00000000000000000000000000000000000000001617845592825163ffffffff16602084015161ffff166040805130602082015299908a90810103601f1981018b5261134d908b6101fb565b604080516064356001600160a01b031660208083019190915281529a90611374908c6101fb565b61137e8680611d34565b86549c9160e08e901c60ff1691369061139692612774565b9060ff166113a3916144f1565b9060a089015192604089016113b8908a611d67565b6113c291506127ed565b946113cd908a611d34565b9690976113d861021e565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081168252600435811660208301529d909d1660408e0152600060608e015263ffffffff1660808d015261ffff1660a08c0152600060c08c015260e08b015261144c60048801611e2b565b6101008b01526101208a0152610140890152610160880152610180870152369061147592612774565b6101a0850152611483612005565b6114906040840184611d67565b905061193e575b906114df6114ba8594936040611510970151606087015161046460028701611d9d565b60608601528060408601526114d960808601516001600160a01b031690565b9061457b565b60c08601526114ec61283c565b976114fa6040840184611d67565b151590506119305760f01c90505b600435613976565b63ffffffff909116606084015260208601939184521161190657611535825186614609565b6115426040860186611d67565b90506117e6575b61155581959295614f66565b808552602081519101209061156e6040850151516128e1565b9460408101958652606060009401935b604086015180518210156117535760206115b16103d36103d36115a4866115f896612899565b516001600160a01b031690565b6115bf8460608b0151612899565b519060405180809581947f958021a70000000000000000000000000000000000000000000000000000000083526004356004840161292a565b03915afa8015610bfb576001600160a01b0391600091611725575b501680156116cc57906000878b9387838861167661163f8860608f61163790612076565b980151612899565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612a75565b03925af18015610bfb57816116a4916001946000916116ab575b508a519061169e8383612899565b52612899565b500161157e565b6116c6913d8091833e6116be81836101fb565b81019061298d565b38611690565b61054a6116e06115a48460408b0151612899565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b611746915060203d811161174c575b61173e81836101fb565b810190612680565b38611613565b503d611734565b61033685808b867fb3005a72901faa1df7bde1059ea556c28eaf46c0259e643959f68398dbf411fd8d6117858d612076565b925193519051906117b66040519283926001600160a01b03606435169767ffffffffffffffff600435169785612c40565b0390a4610baf7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6118486118326117fc6104bc6040890189611d67565b60c08601805151156118eb5751905b602087015161ffff169060e0880151926064359161182d60043591369061285b565b6148a4565b610180830151906118428261288c565b5261288c565b5061186a604061185e8451828701515190612899565b51015163ffffffff1690565b60a061187a61018084015161288c565b5101515163ffffffff82168111611892575050611549565b61054a92506118aa6104c26104bc60408a018a611d67565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b506119006118f98980611d34565b3691612774565b9061180b565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff16611508565b509350600191506040611952910187611d67565b9050036119ba57611510838688946114df6114ba6119ae61197c6104c26104bc6040880188611d67565b602061198e6104bc6040890189611d67565b013561199f602089015161ffff1690565b9060e0890151926004356131cb565b92939495505050611497565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b90506119f36040870187611d67565b90501515906112a0565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005261054a6004359067ffffffffffffffff60249216600452565b611ada915060203d602011611ae0575b611ad281836101fb565b810190612740565b386111d2565b503d611ac8565b3461016a57602060031936011261016a5767ffffffffffffffff600435611b0d8161033a565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611b515760405167ffffffffffffffff9091168152602090f35b6121a3565b3461016a57602060031936011261016a576001600160a01b03600435611b7b81610a83565b611b836142b3565b16338114611bdd578073ffffffffffffffffffffffffffffffffffffffff1960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611c2360043561033a565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611c5d816101c3565b8281528260208201528260408201520152604051611c7a816101c3565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611dcf57505061022e925003836101fb565b84546001600160a01b0316835260019485019487945060209093019201611dba565b90600182811c92168015611e21575b6020831014611e0b57565b634e487b7160e01b600052602260045260246000fd5b91607f1691611e00565b9060405191826000825492611e3f84611df1565b8084529360018116908115611eab5750600114611e64575b5061022e925003836101fb565b90506000929192526020600020906000915b818310611e8f57505090602061022e9282010138611e57565b6020919350806001915483858901015201910190918492611e76565b6020935061022e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611e57565b90611fe56004611ef9610230565b93611f68611f5d8254611f22611f15826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c166040890152611f5760e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b611fbb611fab6001830154611f8c611f818261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b611fc760028201611d9d565b610100860152611fd960038201611d9d565b61012086015201611e2b565b610140830152565b67ffffffffffffffff81116101a25760051b60200190565b604051906120146020836101fb565b6000808352366020840137565b9061202b82611fed565b61203860405191826101fb565b828152601f196120488294611fed565b0190602036910137565b634e487b7160e01b600052603260045260246000fd5b90156120715790565b612052565b35610d4281610a83565b91908110156120715760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561016a570190565b35610d428161033a565b60ff81160361016a57565b35610d42816120ca565b63ffffffff81160361016a57565b35610d42816120df565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b92919061213981611fed565b9361214760405195866101fb565b602085838152019160051b810192831161016a57905b82821061216957505050565b60208091833561217881610a83565b81520191019061215d565b61ffff81160361016a57565b35610d4281612183565b35610d4281610f62565b634e487b7160e01b600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611b5157565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611b5157565b908160031b9180830460081490151715611b5157565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611b5157565b81810292918115918404141715611b5157565b81811061227f575050565b60008155600101612274565b9067ffffffffffffffff83116101a2576801000000000000000083116101a25781548383558084106122ef575b5090600052602060002060005b8381106122d25750505050565b60019060208435946122e386610a83565b019381840155016122c5565b61230790836000528460206000209182019101612274565b386122b8565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610d4293818152019161230d565b9190601f811161234e57505050565b61022e926000526020600020906020601f840160051c8301931061237a575b601f0160051c0190612274565b909150819061236d565b90929167ffffffffffffffff81116101a2576123aa816123a48454611df1565b8461233f565b6000601f82116001146123ea5781906123db9394956000926123df575b50506000198260011b9260031b1c19161790565b9055565b0135905038806123c7565b601f198216946123ff84600052602060002090565b91805b87811061243a575083600195969710612420575b505050811b019055565b60001960f88560031b161c19910135169055388080612416565b90926020600181928686013581550194019101612402565b359061022e826120ca565b359061022e82612183565b359061022e826120df565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b8181106124c25750505090565b9091926020806001926001600160a01b0387356124de81610a83565b1681520194019291016124b5565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b67ffffffffffffffff610d42939216815260406020820152612554604082016125468461034c565b67ffffffffffffffff169052565b61257361256360208401610f57565b6001600160a01b03166060830152565b61258c61258260408401612452565b60ff166080830152565b6125a461259b60608401610f6c565b151560a0830152565b6125be6125b36080840161245d565b61ffff1660c0830152565b6125d86125cd60a0840161245d565b61ffff1660e0830152565b6125f56125e760c08401612468565b63ffffffff16610100830152565b61266d61264061261f61260b60e0860186612473565b6101606101208701526101a08601916124a8565b61262d610100860186612473565b90603f19868403016101408701526124a8565b926126626126516101208301610f57565b6001600160a01b0316610160850152565b6101408101906124ec565b91610180603f198286030191015261230d565b9081602091031261016a5751610d4281610a83565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610bfb576001600160a01b039160009161272357501690565b61273c915060203d60201161174c5761173e81836101fb565b1690565b9081602091031261016a5751610d4281610f62565b67ffffffffffffffff1667ffffffffffffffff8114611b515760010190565b9291926127808261025e565b9161278e60405193846101fb565b82948184528183011161016a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101a257604052606060a0836000815282602082015282604082015282808201528260808201520152565b906127f782611fed565b61280460405191826101fb565b828152601f196128148294611fed565b019060005b82811061282557505050565b6020906128306127ab565b82828501015201612819565b60405190612849826101a7565b60606040838281528260208201520152565b919082604091031261016a57604051612873816101df565b6020808294803561288381610a83565b84520135910152565b8051156120715760200190565b80518210156120715760209160051b010190565b604051906128bc6020836101fb565b600080835282815b8281106128d057505050565b8060606020809385010152016128c4565b906128eb82611fed565b6128f860405191826101fb565b828152601f196129088294611fed565b019060005b82811061291957505050565b80606060208093850101520161290d565b60409067ffffffffffffffff610d42949316815281602082015201906102b2565b81601f8201121561016a5780516129618161025e565b9261296f60405194856101fb565b8184526020828401011161016a57610d42916020808501910161028f565b9060208282031261016a57815167ffffffffffffffff811161016a57610d42920161294b565b9080602083519182815201916020808360051b8301019401926000915b8383106129df57505050505090565b9091929394602080612a6683601f1986600196030187528951908151815260a0612a55612a43612a31612a1f8887015160c08a88015260c08701906102b2565b604087015186820360408801526102b2565b606086015185820360608701526102b2565b608085015184820360808601526102b2565b9201519060a08184039101526102b2565b970193019301919392906129d0565b919390610d429593612bbd612bd59260a08652612a9f60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612ba8612b90612b78612b60612b48612b328c61026060e08a0151916101c061018082015201906102b2565b6101008801518d8203609f1901888f01526102b2565b6101208701518c8203609f19016101c08e01526102b2565b610140860151609f198c8303016101e08d01526102b2565b610160850151609f198b8303016102008c01526102b2565b610180840151609f198a8303016102208b01526129b3565b910151609f19878303016102408801526102b2565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102b2565b9080602083519182815201916020808360051b8301019401926000915b838310612c1357505050505090565b9091929394602080612c3183601f19866001960301875289516102b2565b97019301930191939290612c04565b9493916001600160a01b03612c63921686526080602087015260808601906102b2565b938085036040820152825180865260208601906020808260051b8901019501916000905b828210612ca55750505050610d429394506060818403910152612be7565b90919295602080612d0683601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102b2565b980192019201909291612c87565b60405190610100820182811067ffffffffffffffff8211176101a257604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612dbf575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a5782612e1991850161294b565b926020810151612e28816120df565b92604082015167ffffffffffffffff811161016a57610d42920161294b565b60409067ffffffffffffffff610d429593168152816020820152019161230d565b91929092612e74612d14565b60048310158061303d575b15612f87575090612e8f91615223565b92612e9d6040850151615412565b80612f6c575b60408401612ec081519260608701938451610120880151916154b6565b9092525260c083015151612f1c575b50608082016001600160a01b03612eed82516001600160a01b031690565b1615612ef857505090565b612f0f60e0610d429301516001600160a01b031690565b6001600160a01b03169052565b612f30612f2c6060840151151590565b1590565b15612ecf577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff612f80845163ffffffff1690565b1615612ea3565b9491600090612fdf92612fa86103d36103d36002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612e47565b03915afa8015610bfb57600090600090600090613011575b60a088015263ffffffff16865290505b60c0850152612e9d565b505050613033613007913d806000833e61302b81836101fb565b810190612df1565b9192508291612ff7565b5063302326cb60e01b7fffffffff0000000000000000000000000000000000000000000000000000000061307a6130748686612d65565b90612d8b565b1614612e7f565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a5781516130b581611fed565b926130c360405194856101fb565b81845260208085019260051b82010192831161016a57602001905b8282106130eb5750505090565b6020809183516130fa81610a83565b8152019101906130de565b95949060009460a09467ffffffffffffffff61314c956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102b2565b930152565b9060028201809211611b5157565b9060018201809211611b5157565b6001019081600111611b5157565b9060148201809211611b5157565b90600c8201809211611b5157565b91908201809211611b5157565b6000198114611b515760010190565b80548210156120715760005260206000200190600090565b929394919060036131f08567ffffffffffffffff166000526004602052604060002090565b01936001600160a01b036132058184166126a1565b169182156133dd576040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bfb576000916133be575b50156133ae57613284600095969798604051998a96879586957f89720a6200000000000000000000000000000000000000000000000000000000875260048701613105565b03915afa928315610bfb57600093613389575b5082511561337e578251906132b66132b184548094613197565b612021565b906000928394845b875181101561331d576132d46115a4828a612899565b6001600160a01b03811615613311579061330b6001926132fd6132f68a6131a4565b9989612899565b906001600160a01b03169052565b016132be565b5095506001809661330b565b50919550919361332f575b5050815290565b60005b82811061333f5750613328565b80613378613365613352600194866131b3565b90546001600160a01b039160031b1c1690565b6132fd613371886131a4565b9789612899565b01613332565b9150610d4290611d9d565b6133a79193503d806000833e61339f81836101fb565b810190613081565b9138613297565b505050509250610d429150611d9d565b6133d7915060203d602011611ae057611ad281836101fb565b3861323f565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b9391929361342f6134278251865190613197565b865190613197565b9061344261343c83612021565b926128e1565b94600096875b83518910156134a8578861349e61349160019361347961346f6115a48e9f9d9e9d8b612899565b6132fd838c612899565b613497613486858c612899565b5191809384916131a4565b9c612899565b528b612899565b5001979695613448565b959250929350955060005b865181101561353c576134c96115a48289612899565b60006001600160a01b038216815b888110613510575b50509060019291156134f3575b50016134b3565b61350a906132fd613503896131a4565b9888612899565b386134ec565b816135216103d36115a4848c612899565b1461352e576001016134d7565b5060019150819050386134df565b509390945060005b85518110156135cd5761355a6115a48288612899565b60006001600160a01b038216815b8781106135a1575b5050906001929115613584575b5001613544565b61359b906132fd613594886131a4565b9787612899565b3861357d565b816135b26103d36115a4848b612899565b146135bf57600101613568565b506001915081905038613570565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101a25760405260606080836000815260006020820152600060408201526000838201520152565b9061362282611fed565b61362f60405191826101fb565b828152601f1961363f8294611fed565b019060005b82811061365057505050565b60209061365b6135d9565b82828501015201613644565b9081606091031261016a57805161367d81612183565b916040602083015161368e816120df565b920151610d42816120df565b9160209082815201919060005b8181106136b45750505090565b9091926040806001926001600160a01b0387356136d081610a83565b168152602087810135908201520194019291016136a7565b949391929067ffffffffffffffff1685526080602086015261374161372261371085806124ec565b60a060808a015261012089019161230d565b61372f60208601866124ec565b90607f198984030160a08a015261230d565b6040840135601e198536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a5761022e956137c96137a0836060976137ea978d60c0607f19826137dc9a030191015261369a565b916137bf6137af888301610f57565b6001600160a01b031660e08d0152565b60808101906124ec565b90607f198b8403016101008c015261230d565b9087820360408901526102b2565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611b5157565b908160a091031261016a57805191602082015161382b816120df565b91604081015161383a816120df565b916080606083015161384b81612183565b920151610d4281610f62565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610d429b9a9616885216602087015260408601521660608401521660808201528160a082015201906102b2565b9081606091031261016a57805161367d816120df565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611b5157565b906000198201918211611b5157565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611b5157565b91908203918211611b5157565b919082608091031261016a578151613942816120df565b916020810151916060604083015192015190565b8115613960570490565b634e487b7160e01b600052601260045260246000fd5b93829460009060009560408101946139b06139ab6139a688515161399e60408c01809c611d67565b919050613197565b613151565b613618565b9660009586955b88518051881015613c14576103d36103d36115a48a6139d594612899565b613a22602060608801926139ea8b8551612899565b51908a6040518095819482937f958021a70000000000000000000000000000000000000000000000000000000084526004840161292a565b03915afa8015610bfb576001600160a01b0391600091613bf6575b50168015613ba2579060608e9392613a568b8451612899565b5190613a6760208b015161ffff1690565b958b613aa2604051988995869485947f80485e25000000000000000000000000000000000000000000000000000000008652600486016136e8565b03915afa8015610bfb57600193613b3f938b8f8f95600080958197613b48575b509083929161ffff613aea85613ae36115a4613b3399613b399d9e51612899565b9451612899565b5191613b06613af7610240565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261169e8383612899565b506137f5565b996137f5565b960195966139b7565b613b3997506115a4965084939291509361ffff613aea82613ae3613b85613b339960603d8111613b9b575b613b7d81836101fb565b810190613667565b9c9196909c9d5050505050505090919293613ac2565b503d613b73565b61054a88613bb46115a48c8f51612899565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613c0e915060203d811161174c5761173e81836101fb565b38613a3d565b50919a9496929395509897968a613c2b8187611d67565b9050613f3f575b50508651613c3f906138b5565b99613c4d6020860186611d34565b91613c59915086611d67565b9560609150019486613c6a87612076565b91613c75938a6156cf565b613c7f8b89612899565b52613c8a8a88612899565b50613c958a88612899565b516020015163ffffffff16613ca9916137f5565b90613cb48a88612899565b516040015163ffffffff16613cc8916137f5565b91613cd1610240565b33815290600060208301819052604083015261ffff166060820152613cf461027a565b60808201528651613d04906138e2565b90613d0f8289612899565b52613d1a9087612899565b506002546001600160a01b031692613d3190612076565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610bfb57600096600097600094600092613f02575b506000965b8651881015613e8257613e0b600191613dd388613dce876121b9565b613956565b613dec6060613de28d8d612899565b5101918251612261565b9052858a14613e13575b6060613e028b8b612899565b51015190613197565b970196613db2565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613e4660808c01516001600160a01b031690565b1603613e54575b5050613df6565b613dce613e60926121e9565b613e796060613e6f8d8d612899565b5101918251613197565b90528b88613e4d565b9796509750505050613ebe7f000000000000000000000000000000000000000000000000000000000000000091613dce63ffffffff84166121e9565b8411613eca5750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b9298505050613f2a91925060803d608011613f38575b613f2281836101fb565b81019061392b565b919790939290919038613dad565b503d613f18565b610abf6103d36104c26104bc613f58948a989698611d67565b926001600160a01b03600091515194169060e08801908151613f78610240565b6001600160a01b0385168152908260208301528260408301528260608301526080820152613fa6878d612899565b52613fb1868c612899565b506040516301ffc9a760e01b8152633317103160e01b600482015291602083602481875afa8015610bfb578f948c89968f96948d948f968891614294575b5061418d575b50505050505015614034575b61185e61402561402c9561401f602061185e61401f97604097612899565b906137f5565b958b612899565b90388a613c32565b50506140ba9160608c6140656104c26104bc61405e6103d36103d36002546001600160a01b031690565b938b611d67565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610bfb5761185e61402560409261401f602061185e8f8b9061402c9b61401f9a60008060009261414f575b63ffffffff92935061413c9060606141048888612899565b5101946141318a6141158a8a612899565b51019160406141248b8b612899565b51019063ffffffff169052565b9063ffffffff169052565b1690529750975050505095505050614001565b50505063ffffffff61417b61413c9260603d606011614186575b61417381836101fb565b81019061389f565b9093509150826140ec565b503d614169565b8495985060a09697506141d760206141cd6060826141c46104bc6141bd6104c26104bc8b6142109c9d9e9f611d67565b998d611d67565b01359901612076565b9a015161ffff1690565b905190604051998a97889687967f2c06340400000000000000000000000000000000000000000000000000000000885260048801613857565b03915afa8015610bfb578592828c939181908294614258575b5061424c90606061423a8888612899565b51019261413160206141158a8a612899565b5288888f8c8138613ff5565b91505061424c9250614282915060a03d60a01161428d575b61427a81836101fb565b81019061380f565b949192919050614229565b503d614270565b6142ad915060203d602011611ae057611ad281836101fb565b38613fef565b6001600160a01b036001541633036142c757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916142ff815184613197565b9283156143f85760005b848110614317575050505050565b818110156143dd5761432c6115a48286612899565b6001600160a01b0381168015610c00576143458361315f565b87811061435757505050600101614309565b848110156143ba576001600160a01b036143746115a4838a612899565b16821461438357600101614345565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036143d86115a46143d2888561391e565b89612899565b614374565b6143f36115a46143ed848461391e565b85612899565b61432c565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156120715760051b0190565b9081602091031261016a575190565b91608061022e9294936144918160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610d429281815201906102b2565b906144d48261025e565b6144e160405191826101fb565b828152601f19612048829461025e565b9081516020821080614571575b614540570361450a5790565b6109bc906040519182917f3aeba390000000000000000000000000000000000000000000000000000000008352600483016144b9565b5090602081015190816145528461221b565b1c61450a5750614561826144ca565b9160200360031b1b602082015290565b50602081146144fe565b918251601481029080820460141490151715611b515761459d6145a29161316d565b61317b565b906145b46145af83613189565b6144ca565b9060146145c08361288c565b5360009260215b86518510156145f25760146001916145e26115a4888b612899565b60601b81870152019401936145c7565b919550936020935060601b90820152828152012090565b906146196103d360608401612076565b61462a600019936040810190611d67565b90506146a2575b61463b82516138e2565b9260005b84811061464d575050505050565b80826001921461469d5760606146638287612899565b5101518015614697576146919061468b61467d8489612899565b51516001600160a01b031690565b866158d8565b0161463f565b50614691565b614691565b91506146ae81516138f1565b916146bc61467d8484612899565b6040516301ffc9a760e01b8152633317103160e01b60048201526020816024816001600160a01b0386165afa908115610bfb57600091614723575b50614703575b50614631565b61471d9060606147138686612899565b51015190836158d8565b386146fd565b61473c915060203d602011611ae057611ad281836101fb565b386146f7565b6040519061474f826101df565b60606020838281520152565b919060408382031261016a5760405190614774826101df565b8193805167ffffffffffffffff811161016a578261479391830161294b565b835260208101519167ffffffffffffffff831161016a576020926147b7920161294b565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a57610d42920161475b565b9060806001600160a01b0381614801855160a0865260a08601906102b2565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610d429281815201906147e2565b919060408382031261016a57825167ffffffffffffffff811161016a5760209161487191850161475b565b92015190565b61ffff614890610d4295936060845260608401906147e2565b9316602082015260408184039101526102b2565b9091939295946148b26127ab565b5060208201805115614d01576148d5610abf6103d385516001600160a01b031690565b946001600160a01b038616916040516301ffc9a760e01b81526020818061492360048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610bfb57600091614ce2575b5015614c99576149af88999a82519261494e614742565b505161499a61496489516001600160a01b031690565b92604061496f610240565b9e8f90815261498b8d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040516301ffc9a760e01b8152633317103160e01b6004820152602081602481875afa908115610bfb57600091614c7a575b5015614b825750906000929183614a299899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614877565b03925af18015610bfb57600094600091614b3a575b50614b0f614a789596614ab2614a86614aa4956115a46020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b6101fb565b6040519586918683019190916001600160a01b036020820193169052565b03601f1981018652856101fb565b614aee610905614ae4614af48951614aee610905614ae48c67ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b906144f1565b9767ffffffffffffffff166000526004602052604060002090565b93015193614b1b61024f565b958652602086015260408501526060840152608083015260a082015290565b614a789550602091506115a496614ab2614a86614aa495614b70614b0f953d806000833e614b6881836101fb565b810190614846565b9b909b96505095505050969550614a3e565b9793929061ffff16614c505751614c2657614bd16000939184926040519586809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614835565b03925af1908115610bfb57614a7895614ab2614a86614b0f936115a4602096614aa498600091614c03575b5099614a5a565b614c2091503d806000833e614c1881836101fb565b8101906147bc565b38614bfc565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614c93915060203d602011611ae057611ad281836101fb565b386149e1565b61054a614cad86516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b614cfb915060203d602011611ae057611ad281836101fb565b38614937565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9592614dfd947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b90614e196020928281519485920161028f565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b168152614e65825180936020898501910161028f565b019160f81b1683820152614e8382518093602060028501910161028f565b01019160f81b1683820152614ea282518093602060028501910161028f565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff00000000000000000000000000000000000000000000000000000000000000610d429f9e9c9860f81b168152614f17825180936020898501910161028f565b019160f01b1683820152614f3582518093602060038501910161028f565b01019160f01b1683820152614f53825180936020898501910161028f565b01019160f01b1660028201520190614e06565b60e081019060ff8251511161520d57610100810160ff815151116151f75761012082019260ff845151116151e157610140830160ff815151116151cb5761016084019061ffff825151116151b55761018085019360018551511161519d576101a086019361ffff855151116151875760609551805161514b575b50865167ffffffffffffffff16602088015167ffffffffffffffff169060408901516150139067ffffffffffffffff1690565b9860608101516150269063ffffffff1690565b9060808101516150399063ffffffff1690565b60a082015161ffff169160c00151926040519c8d96602088019661505c97614d2b565b03601f198101885261506e90886101fb565b5190815161507c9060ff1690565b9051805160ff1698519081516150929060ff1690565b906040519a8b9560208701956150a796614e1d565b03601f19810187526150b990876101fb565b519182516150c79060ff1690565b9151805161ffff169480516150dd9061ffff1690565b92519283516150ed9061ffff1690565b92604051978897602089019761510298614ea8565b03601f198101825261511490826101fb565b6040519283926020840161512791614e06565b61513091614e06565b61513991614e06565b03601f1981018252610d4290826101fb565b61516091965061515a9061288c565b51615acf565b9461ffff8651116151715738614fe0565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b600052602660045260246000fd5b635a102da160e11b60005261054a6024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b9061522c612d14565b91601182106153f757803563302326cb60e01b7fffffffff0000000000000000000000000000000000000000000000000000000082160361539d5750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61529281612021565b604086019081526152a2826128e1565b906060870191825260005b8381106153515750505050615315838361530b6152ff6152f56152ee6152db61531f98876153299c9b615bf6565b6001600160a01b0390911660808d015290565b8585615c9a565b9291903691612774565b60a08a01528383615ce9565b9491903691612774565b60c0880152615c9a565b9391903691612774565b60e08401528103615338575090565b63d9437f9d60e01b600052600360045260245260446000fd5b8060019161539661538061537961536c6153909a8d8d615bf6565b91906132fd868a51612899565b8b8b615c9a565b9391889a919a51949a3691612774565b92612899565b52016152ad565b7f55a0e02c0000000000000000000000000000000000000000000000000000000060005263302326cb60e01b6004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b63d9437f9d60e01b6000526002600452602482905260446000fd5b80519060005b82811061542457505050565b60018101808211611b51575b8381106154405750600101615418565b6001600160a01b036154528385612899565b51166154646103d36115a48487612899565b1461547157600101615430565b61054a6154816115a48486612899565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b9093919281511561564257506154cb81615412565b80519260005b8481106154e057509093925050565b6154f06103d36115a48386612899565b156154fd576001016154d1565b9391946155176132b161550f856138e2565b845190613197565b9261553461552f615527836138e2565b855190613197565b6128e1565b968792600097885b8481106155e45750505050505060005b81518110156155d7576000805b868110615598575b50906001916155935761558d61557a6115a48386612899565b6132fd615586896131a4565b9887612899565b0161554c565b61558d565b6155a56115a48287612899565b6001600160a01b036155bd6103d36115a48789612899565b9116146155cc57600101615559565b506001905080615561565b5050909180825283529190565b909192939498828214615638579061562a61561d836156108b6132fd6001976153906115a4898e612899565b6156236134868589612899565b9e612899565b528c612899565b505b0190899493929161553c565b985060019061562c565b919350501561565d5750615654612005565b90610d426128ad565b90610d4282516128e1565b9081602091031261016a5751610d4281612183565b936156ba60809461ffff6001600160a01b039567ffffffffffffffff6156c8969b9a9b16895216602088015260a0604088015260a0870190610c2a565b9085820360608701526102b2565b9416910152565b929190926156db6135d9565b506156fa8167ffffffffffffffff166000526004602052604060002090565b805490959060e01c60ff16916080850192835161571d906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff1661573d916137f5565b9661574990608d613197565b9460a087019586515161575b91613197565b9160ff169161576983612231565b61577291613197565b9161577e906067613197565b61578791612261565b61579091613197565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b0387160361581d5750505061ffff925061580f906158026000935b51956157f56157e6610240565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b6158336103d3602094976001600160a01b031690565b9060406158448583015161ffff1690565b91015192615884875198604051998a96879586957ff23889580000000000000000000000000000000000000000000000000000000087526004870161567d565b03915afa908115610bfb5761580261580f9261ffff956000916158a9575b50936157d9565b6158cb915060203d6020116158d1575b6158c381836101fb565b810190615668565b386158a2565b503d6158b9565b91602091600091604051906001600160a01b03858301937fa9059cbb00000000000000000000000000000000000000000000000000000000855216602483015260448201526044815261592c6064826101fb565b519082855af115612695576000513d61598b57506001600160a01b0381163b155b6159545750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6001141561594d565b966002615aa397615a706022610d429f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090615a709f82615a709c615a779c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615a21825180936020898501910161028f565b019160f81b1683820152615a3f82518093602060238501910161028f565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190614e06565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff82515111615be057604081019160ff83515111615bca57606082019160ff83515111615bb457608081019260ff84515111615b9e5760a0820161ffff81515111615b8857610d4294615b7a9351945191615b31835160ff1690565b975191615b3f835160ff1690565b945190615b4d825160ff1690565b905193615b5b855160ff1690565b935196615b6a885161ffff1690565b966040519c8d9b60208d01615994565b03601f1981018352826101fb565b635a102da160e11b600052602b60045260246000fd5b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b929190926001820191848311615c815781013560001a828115615c76575060148103615c49578201938411615c2e57013560601c9190565b63d9437f9d60e01b6000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b63d9437f9d60e01b600052600060045260245260446000fd5b91906002820191818311615c81578381013560f01c0160020192818411615cce57918391615cc793612d73565b9290929190565b63d9437f9d60e01b6000526001600452602483905260446000fd5b91906001820191818311615c81578381013560001a0160010192818411615cce57918391615cc793612d7356fea164736f6c634300081a000a",
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
