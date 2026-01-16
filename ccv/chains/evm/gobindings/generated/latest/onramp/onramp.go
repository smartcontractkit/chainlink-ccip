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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextMessageNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAmountBeforeTokenPoolFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DestChainConfigArgs\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tokenReceiverAllowed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenNetworkFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseExecutionGasCost\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupportedByCCV\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxAllowed\",\"inputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUSDCentsPerMessage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFeeTokenAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum ExtraArgsCodec.EncodingErrorLocation\"},{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidExtraArgsTag\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"actual\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SourceTokenDataTooLarge\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"actualLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TokenArgsNotSupportedOnPoolV1\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenReceiverNotAllowed\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61010060405234610367576040516163e938819003601f8101601f191683016001600160401b0381118482101761036c578392829160405283398101039060e08212610367576080821261036757610055610382565b81519092906001600160401b03811681036103675783526020820151926001600160a01b03841684036103675760208101938452604083015163ffffffff81168103610367576040820190815260606100af8186016103a1565b83820190815293607f1901126103675760405192606084016001600160401b0381118582101761036c576040526100e8608086016103a1565b845260a08501519485151586036103675760c061010c9160208701978852016103a1565b9560408501968752331561035657600180546001600160a01b0319163317905583516001600160401b0316158015610344575b8015610332575b8015610323575b6103085792516001600160401b031660805291516001600160a01b0390811660a0529151821660c0525163ffffffff1660e052815116158015610319575b6103085780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9260e09260006060610206610382565b8281526020810183905260408101839052015260805160a051855160c0516001600160401b0390931695926001600160a01b039081169263ffffffff92831691166060610251610382565b89815260208082019384526040808301958652929091019586528151998a5291516001600160a01b03908116928a0192909252915192909216908701529051811660608601529051811660808501529051151560a084015290511660c0820152a160405161603390816103b68239608051818181610ac40152818161147f0152611d45015260a0518181816112410152611d71015260c051818181611dcc015261282d015260e051818181611d9d01526140410152f35b6306b7c75960e31b60005260046000fd5b508151151561018b565b5063ffffffff8351161561014d565b5081516001600160a01b031615610146565b5080516001600160a01b03161561013f565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190608082016001600160401b0381118382101761036c57604052565b51906001600160a01b03821682036103675756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd5780632490769e146100f857806348a98aa4146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da57806390423fa2146100d5578063df0aa9e9146100d0578063e8d80861146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b611cc7565b611c0b565b611b9c565b61119b565b610fea565b610fa3565b610ef2565b610e7f565b610dad565b610b41565b610afc565b61060a565b61037e565b6102f0565b3461016a57600060031936011261016a576080610122611d0d565b61016860405180926001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610160810190811067ffffffffffffffff8211176101bb57604052565b61016f565b6060810190811067ffffffffffffffff8211176101bb57604052565b6080810190811067ffffffffffffffff8211176101bb57604052565b6040810190811067ffffffffffffffff8211176101bb57604052565b90601f601f19910116810190811067ffffffffffffffff8211176101bb57604052565b604051906102476101c083610214565b565b6040519061024761016083610214565b6040519061024760a083610214565b6040519061024760c083610214565b67ffffffffffffffff81116101bb57601f01601f191660200190565b604051906102a2602083610214565b60008252565b60005b8381106102bb5750506000910152565b81810151838201526020016102ab565b90601f19601f6020936102e9815180928187528780880191016102a8565b0116010190565b3461016a57600060031936011261016a5761034f60408051906103138183610214565b601082527f4f6e52616d7020312e372e302d646576000000000000000000000000000000006020830152519182916020835260208301906102cb565b0390f35b67ffffffffffffffff81160361016a57565b359061024782610353565b908160a091031261016a5790565b3461016a57604060031936011261016a5760043561039b81610353565b60243567ffffffffffffffff811161016a576103bb903690600401610370565b6103d98267ffffffffffffffff166000526004602052604060002090565b918254916001600160a01b036104046103f8856001600160a01b031690565b6001600160a01b031690565b161561056f576040810191600161041b8484611df4565b9050116105455761034f946104be946104aa61046361043d6080870187611e2a565b61044a6020890189611e2a565b905015918261052f575b61045d87611fc4565b88612fb6565b9561046c6120de565b6104768288611df4565b90506104de575b604088016104a081519260608b0193845161049a60028b01611e5d565b916135ac565b9092525285611df4565b151590506104d05760f01c90505b90613b28565b60405190815292839250602083019150565b506001015461ffff166104b8565b5061052a6104fd6104f86104f2848a611df4565b9061215a565b612168565b602061050c6104f2858b611df4565b013561051d60208b015161ffff1690565b9060e08b01519289613332565b61047d565b915061053b8989611df4565b9050151591610454565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff821660045260246000fd5b6000fd5b90602060031983011261016a5760043567ffffffffffffffff811161016a5760040160009280601f830112156106065781359367ffffffffffffffff851161060357506020808301928560051b01011161016a579190565b80fd5b8380fd5b3461016a57610618366105ab565b90610621614497565b6000915b80831061062e57005b610639838284612172565b92610643846121b2565b67ffffffffffffffff81169081158015610ab8575b8015610aa2575b8015610a89575b610a5257856108c8916108e26108d8836108d26106d860e08301956106be6106b861069189876121e9565b6106b06106a66101008a95949501809a6121e9565b949092369161221f565b92369161221f565b906144d5565b67ffffffffffffffff166000526004602052604060002090565b96879561071e6106ea60208a01612168565b88906001600160a01b03167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b6108c261088b60c060408b019a6107856107378d6121c7565b8c547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178c55565b6107e361079460808301612281565b8c547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016178c55565b61082660016107f460a08401612281565b9c019b8c9061ffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000825416179055565b6108858d6108366060840161228b565b81547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690151560e81b7dff000000000000000000000000000000000000000000000000000000000016179055565b016121df565b88547fffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ffff1660109190911b65ffffffff000016178855565b8d6121e9565b9060038901612396565b8a6121e9565b9060028601612396565b6101208801906108f46103f883612168565b15610a285761090561094f92612168565b7fffffffffffff0000000000000000000000000000000000000000ffffffffffff79ffffffffffffffffffffffffffffffffffffffff00000000000083549260301b169116179055565b61014087019061097361096d610965848b611e2a565b9390506121c7565b60ff1690565b036109e457956109ca7f99415f1fd5d7f97dec3730fd98d0161792f21251c2e963782304b609b288cb26926109b96109af600198999a85611e2a565b906004840161248f565b5460a01c67ffffffffffffffff1690565b6109d960405192839283612629565b0390a2019190610625565b6109ee9087611e2a565b90610a246040519283927f3aeba39000000000000000000000000000000000000000000000000000000000845260048401612439565b0390fd5b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff610a9b60c088016121df565b1615610666565b5060ff610ab1604088016121c7565b161561065f565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610658565b6001600160a01b0381160361016a57565b3461016a57604060031936011261016a57610b18600435610353565b610b2c602435610b2781610aeb565b6127e8565b6040516001600160a01b039091168152602090f35b3461016a57610b4f366105ab565b906001600160a01b0360035416918215610c685760005b818110610b6f57005b610b806103f86104f8838587614606565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529091906001600160a01b03831690602081602481855afa8015610c63576001948892600092610c33575b5081610be7575b5050505001610b66565b81610c177f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385610c2794615b9d565b6040519081529081906020820190565b0390a338858180610bdd565b610c5591925060203d8111610c5c575b610c4d8183610214565b810190614616565b9038610bd6565b503d610c43565b6127dc565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b906020808351928381520192019060005b818110610cb05750505090565b82516001600160a01b0316845260209384019390920191600101610ca3565b90610daa9160208152610cee6020820183516001600160a01b03169052565b602082015167ffffffffffffffff166040820152604082015160ff166060820152606082015115156080820152608082015161ffff1660a082015260a082015161ffff1660c082015260c082015163ffffffff1660e082015260e08201516001600160a01b0316610100820152610140610d94610d7e610100850151610160610120860152610180850190610c92565b610120850151601f198583030184860152610c92565b92015190610160601f19828503019101526102cb565b90565b3461016a57602060031936011261016a5767ffffffffffffffff600435610dd381610353565b6060610140604051610de48161019e565b60008152600060208201526000604082015260008382015260006080820152600060a0820152600060c0820152600060e08201528261010082015282610120820152015216600052600460205261034f610e416040600020611fc4565b60405191829182610ccf565b6102479092919260608101936001600160a01b0360408092828151168552602081015115156020860152015116910152565b3461016a57600060031936011261016a57600060408051610e9f816101c0565b828152826020820152015261034f604051610eb9816101c0565b60ff6002546001600160a01b038116835260a01c16151560208201526001600160a01b0360035416604082015260405191829182610e4d565b3461016a57600060031936011261016a576000546001600160a01b0381163303610f79577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57600060031936011261016a5760206001600160a01b0360015416604051908152f35b359061024782610aeb565b8015150361016a57565b359061024782610fd5565b3461016a57606060031936011261016a576000604051611009816101c0565b60043561101581610aeb565b815260243561102381610fd5565b602082019081526044359061103782610aeb565b60408301918252611046614497565b6001600160a01b03835116158015611191575b611169576001600160a01b03839261111661114c936110ca847f0a7df5db91fe3762aa97fa5fb05e9c7adfb1fb97fa4c95ec9cfc0d755e6ef85c9851166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006002541617600255565b5115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff000000000000000000000000000000000000000060025492151560a01b16911617600255565b51166001600160a01b03167fffffffffffffffffffffffff00000000000000000000000000000000000000006003541617600355565b611154611d0d565b61116360405192839283614625565b0390a180f35b6004847f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b5080511515611059565b3461016a57608060031936011261016a576111b7600435610353565b60243567ffffffffffffffff811161016a576111d7903690600401610370565b604435906111e6606435610aeb565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081526004803560801b77ffffffffffffffff0000000000000000000000000000000016908201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610c6357600091611b6d575b50611b305760025460a01c60ff16611b06576112c8740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b6112e860043567ffffffffffffffff166000526004602052604060002090565b906001600160a01b036064351615611adc57815461130f6103f86001600160a01b03831681565b3303611ab25781611324608086940182611e2a565b6113316020840184611e2a565b9050159081611a99575b61134487611fc4565b6113519390600435612fb6565b9160a01c67ffffffffffffffff166113689061289c565b84547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff000000000000000000000000000000000000000016178555825163ffffffff169460208401516113cb9061ffff1690565b604080513060208083019190915281529791906113e89089610214565b604080516064356001600160a01b03166020808301919091528152989061140f908a610214565b6114198680611e2a565b85549a9160e08c901c60ff16913690611431926128bb565b9060ff1661143e916146d5565b60a08901519161145160408a018a611df4565b61145b9150612934565b9361146960208b018b611e2a565b969097611474610237565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529a67ffffffffffffffff6004351660208d015267ffffffffffffffff1660408c0152600060608c015263ffffffff1660808b015261ffff1660a08a0152600060c08a015260e08901526114f660048801611f04565b610100890152610120880152610140870152610160860152610180850152369061151f926128bb565b6101a083015261152d6120de565b61153a6040850185611df4565b9050611a1f575b836115ba9261158961156488946040860151606087015161049a60028701611e5d565b606086015280604086015261158360808601516001600160a01b031690565b9061475f565b60c0860152611596612983565b986115a46040840184611df4565b15159050611a115760f01c90505b600435613b28565b63ffffffff90911660608401526020870195918652116119e7576115df8451836147ed565b6115ec6040830183611df4565b90506118c7575b6115ff819592956151c7565b8083526020815191012090611618604085015151612a28565b6040840190815294606087019360005b604087015180518210156117fd57602061165b6103f86103f861164e866116a2966129e0565b516001600160a01b031690565b6116698460608c01516129e0565b519060405180809581947f958021a700000000000000000000000000000000000000000000000000000000835260043560048401612a71565b03915afa8015610c63576001600160a01b03916000916117cf575b5016801561177657906000888c938883896117206116e9888f6116e1606091612168565b9801516129e0565b51604051998a97889687957f3bbbed4b00000000000000000000000000000000000000000000000000000000875260048701612bbc565b03925af18015610c63578161174e91600194600091611755575b508b519061174883836129e0565b526129e0565b5001611628565b611770913d8091833e6117688183610214565b810190612ad4565b3861173a565b6105a761178a61164e8460408c01516129e0565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660049081523567ffffffffffffffff16602452604490565b6117f0915060203d81116117f6575b6117e88183610214565b8101906127c7565b386116bd565b503d6117de565b61034f8680858d7f371bc2ff0a006f4ef863b1d27a065d4e9f938b6d883eb154572b4aea593b32cc8e8a6118308f612168565b9361183e6040820182611df4565b1590506118bb57602061185b6104f283604061188b950190611df4565b0135955b51915192516040519384936001600160a01b03606435169867ffffffffffffffff600435169886612d87565b0390a4610c177fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b5061188b60009561185f565b6119296119136118dd6104f26040860186611df4565b60c08601805151156119cc5751905b602087015161ffff169060e0880151926064359161190e6004359136906129a2565b614aba565b61018083015190611923826129d3565b526129d3565b5061194b604061193f86518287015151906129e0565b51015163ffffffff1690565b60a061195b6101808401516129d3565b5101515163ffffffff821681116119735750506115f3565b6105a7925061198b6104f86104f26040870187611df4565b7f06cf7cbc000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260245263ffffffff16604452606490565b506119e16119da8680611e2a565b36916128bb565b906118ec565b7f07da6ee60000000000000000000000000000000000000000000000000000000060005260046000fd5b506001015461ffff166115b2565b5093506001611a316040840184611df4565b905003610545576115ba83838896611589611564611a8d611a5b6104f86104f26040880188611df4565b6020611a6d6104f26040890189611df4565b0135611a7e602089015161ffff1690565b9060e089015192600435613332565b94505050925050611541565b9050611aa86040840184611df4565b905015159061133b565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b7ffdbd6a72000000000000000000000000000000000000000000000000000000006000526105a76004359067ffffffffffffffff60249216600452565b611b8f915060203d602011611b95575b611b878183610214565b810190612887565b38611272565b503d611b7d565b3461016a57602060031936011261016a5767ffffffffffffffff600435611bc281610353565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111611c065760405167ffffffffffffffff9091168152602090f35b612295565b3461016a57602060031936011261016a576001600160a01b03600435611c3081610aeb565b611c38614497565b16338114611c9d57807fffffffffffffffffffffffff000000000000000000000000000000000000000060005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3461016a57602060031936011261016a57611ce3600435610353565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60006060604051611d1d816101dc565b8281528260208201528260408201520152604051611d3a816101dc565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016602082015263ffffffff7f00000000000000000000000000000000000000000000000000000000000000001660408201526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016606082015290565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160061b3603831361016a57565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a5760200191813603831361016a57565b906040519182815491828252602082019060005260206000209260005b818110611e8f57505061024792500383610214565b84546001600160a01b0316835260019485019487945060209093019201611e7a565b90600182811c92168015611efa575b6020831014611ecb57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611ec0565b9060405191826000825492611f1884611eb1565b8084529360018116908115611f845750600114611f3d575b5061024792500383610214565b90506000929192526020600020906000915b818310611f685750509060206102479282010138611f30565b6020919350806001915483858901015201910190918492611f4f565b602093506102479592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611f30565b906120be6004611fd2610249565b936120416120368254611ffb611fee826001600160a01b031690565b6001600160a01b03168952565b67ffffffffffffffff60a082901c16602089015260ff60e082901c16604089015261203060e882901c60ff16151560608a0152565b60f01c90565b61ffff166080870152565b612094612084600183015461206561205a8261ffff1690565b61ffff1660a08a0152565b63ffffffff601082901c1660c089015260301c6001600160a01b031690565b6001600160a01b031660e0870152565b6120a060028201611e5d565b6101008601526120b260038201611e5d565b61012086015201611f04565b610140830152565b67ffffffffffffffff81116101bb5760051b60200190565b604051906120ed602083610214565b6000808352366020840137565b90612104826120c6565b6121116040519182610214565b828152601f1961212182946120c6565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90156121635790565b61212b565b35610daa81610aeb565b91908110156121635760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea18136030182121561016a570190565b35610daa81610353565b60ff81160361016a57565b35610daa816121bc565b63ffffffff81160361016a57565b35610daa816121d1565b903590601e198136030182121561016a570180359067ffffffffffffffff821161016a57602001918160051b3603831361016a57565b92919061222b816120c6565b936122396040519586610214565b602085838152019160051b810192831161016a57905b82821061225b57505050565b60208091833561226a81610aeb565b81520191019061224f565b61ffff81160361016a57565b35610daa81612275565b35610daa81610fd5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906d04ee2d6d415b85acef81000000008202918083046d04ee2d6d415b85acef81000000001490151715611c0657565b906e01ed09bead87c0378d8e64000000008202918083046e01ed09bead87c0378d8e64000000001490151715611c0657565b908160031b9180830460081490151715611c0657565b908160011b917f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811603611c0657565b81810292918115918404141715611c0657565b81811061238a575050565b6000815560010161237f565b9067ffffffffffffffff83116101bb576801000000000000000083116101bb5781548383558084106123fa575b5090600052602060002060005b8381106123dd5750505050565b60019060208435946123ee86610aeb565b019381840155016123d0565b6124129083600052846020600020918201910161237f565b386123c3565b601f8260209493601f19938186528686013760008582860101520116010190565b916020610daa938181520191612418565b9190601f811161245957505050565b610247926000526020600020906020601f840160051c83019310612485575b601f0160051c019061237f565b9091508190612478565b90929167ffffffffffffffff81116101bb576124b5816124af8454611eb1565b8461244a565b6000601f82116001146124f55781906124e69394956000926124ea575b50506000198260011b9260031b1c19161790565b9055565b0135905038806124d2565b601f1982169461250a84600052602060002090565b91805b87811061254557508360019596971061252b575b505050811b019055565b60001960f88560031b161c19910135169055388080612521565b9092602060018192868601358155019401910161250d565b3590610247826121bc565b359061024782612275565b3590610247826121d1565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a578160051b3603831361016a57565b9160209082815201919060005b8181106125cd5750505090565b9091926020806001926001600160a01b0387356125e981610aeb565b1681520194019291016125c0565b9035601e198236030181121561016a57016020813591019167ffffffffffffffff821161016a57813603831361016a57565b67ffffffffffffffff610daa93921681526040602082015261265f6040820161265184610365565b67ffffffffffffffff169052565b61267e61266e60208401610fca565b6001600160a01b03166060830152565b61269761268d6040840161255d565b60ff166080830152565b6126af6126a660608401610fdf565b151560a0830152565b6126c96126be60808401612568565b61ffff1660c0830152565b6126e36126d860a08401612568565b61ffff1660e0830152565b6127006126f260c08401612573565b63ffffffff16610100830152565b61279661276961272a61271660e086018661257e565b6101606101208701526101a08601916125b3565b61273861010086018661257e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0868403016101408701526125b3565b9261278b61277a6101208301610fca565b6001600160a01b0316610160850152565b6101408101906125f7565b916101807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082860301910152612418565b9081602091031261016a5751610daa81610aeb565b6040513d6000823e3d90fd5b6001600160a01b03604051917fbbe4f6db0000000000000000000000000000000000000000000000000000000083521660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015610c63576001600160a01b039160009161286a57501690565b612883915060203d6020116117f6576117e88183610214565b1690565b9081602091031261016a5751610daa81610fd5565b67ffffffffffffffff1667ffffffffffffffff8114611c065760010190565b9291926128c782610277565b916128d56040519384610214565b82948184528183011161016a578281602093846000960137010152565b6040519060c0820182811067ffffffffffffffff8211176101bb57604052606060a0836000815282602082015282604082015282808201528260808201520152565b9061293e826120c6565b61294b6040519182610214565b828152601f1961295b82946120c6565b019060005b82811061296c57505050565b6020906129776128f2565b82828501015201612960565b60405190612990826101c0565b60606040838281528260208201520152565b919082604091031261016a576040516129ba816101f8565b602080829480356129ca81610aeb565b84520135910152565b8051156121635760200190565b80518210156121635760209160051b010190565b60405190612a03602083610214565b600080835282815b828110612a1757505050565b806060602080938501015201612a0b565b90612a32826120c6565b612a3f6040519182610214565b828152601f19612a4f82946120c6565b019060005b828110612a6057505050565b806060602080938501015201612a54565b60409067ffffffffffffffff610daa949316815281602082015201906102cb565b81601f8201121561016a578051612aa881610277565b92612ab66040519485610214565b8184526020828401011161016a57610daa91602080850191016102a8565b9060208282031261016a57815167ffffffffffffffff811161016a57610daa9201612a92565b9080602083519182815201916020808360051b8301019401926000915b838310612b2657505050505090565b9091929394602080612bad83601f1986600196030187528951908151815260a0612b9c612b8a612b78612b668887015160c08a88015260c08701906102cb565b604087015186820360408801526102cb565b606086015185820360608701526102cb565b608085015184820360808601526102cb565b9201519060a08184039101526102cb565b97019301930191939290612b17565b919390610daa9593612d04612d1c9260a08652612be660a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152606081015163ffffffff16610100870152608081015163ffffffff1661012087015260a081015161ffff1661014087015260c08101516101608701526101a0612cef612cd7612cbf612ca7612c8f612c798c61026060e08a0151916101c061018082015201906102cb565b6101008801518d8203609f1901888f01526102cb565b6101208701518c8203609f19016101c08e01526102cb565b610140860151609f198c8303016101e08d01526102cb565b610160850151609f198b8303016102008c01526102cb565b610180840151609f198a8303016102208b0152612afa565b910151609f19878303016102408801526102cb565b95602085015260408401906001600160a01b03169052565b606082015260808184039101526102cb565b9080602083519182815201916020808360051b8301019401926000915b838310612d5a57505050505090565b9091929394602080612d7883601f19866001960301875289516102cb565b97019301930191939290612d4b565b95949290916001600160a01b03612db193168752602087015260a0604087015260a08601906102cb565b938085036060820152825180865260208601906020808260051b8901019501916000905b828210612df35750505050610daa9394506080818403910152612d2e565b90919295602080612e5483601f198d6001960301865260a060808c516001600160a01b03815116845263ffffffff86820151168685015263ffffffff60408201511660408501526060810151606085015201519181608082015201906102cb565b980192019201909291612dd5565b60405190610100820182811067ffffffffffffffff8211176101bb57604052606060e08360008152600060208201528260408201528280820152600060808201528260a08201528260c08201520152565b9060041161016a5790600490565b9093929384831161016a57841161016a578101920390565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612f0d575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b9160608383031261016a57825167ffffffffffffffff811161016a5782612f67918501612a92565b926020810151612f76816121d1565b92604082015167ffffffffffffffff811161016a57610daa9201612a92565b60409067ffffffffffffffff610daa95931681528160208201520191612418565b91929092612fc2612e62565b60048310158061318b575b156130d5575090612fdd91615484565b92612feb60408501516156d7565b806130ba575b6040840161300e815192606087019384516101208801519161577b565b9092525260c08301515161306a575b50608082016001600160a01b0361303b82516001600160a01b031690565b161561304657505090565b61305d60e0610daa9301516001600160a01b031690565b6001600160a01b03169052565b61307e61307a6060840151151590565b1590565b1561301d577f68a8cf4a0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b5063ffffffff6130ce845163ffffffff1690565b1615612ff1565b949160009061312d926130f66103f86103f86002546001600160a01b031690565b906040518095819482937f9cc199960000000000000000000000000000000000000000000000000000000084528a60048501612f95565b03915afa8015610c635760009060009060009061315f575b60a088015263ffffffff16865290505b60c0850152612feb565b505050613181613155913d806000833e6131798183610214565b810190612f3f565b9192508291613145565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000006131e16131db8686612eb3565b90612ed9565b1614612fcd565b60208183031261016a5780519067ffffffffffffffff821161016a57019080601f8301121561016a57815161321c816120c6565b9261322a6040519485610214565b81845260208085019260051b82010192831161016a57602001905b8282106132525750505090565b60208091835161326181610aeb565b815201910190613245565b95949060009460a09467ffffffffffffffff6132b3956001600160a01b0361ffff95168b521660208a0152604089015216606087015260c0608087015260c08601906102cb565b930152565b9060028201809211611c0657565b9060018201809211611c0657565b6001019081600111611c0657565b9060148201809211611c0657565b90600c8201809211611c0657565b91908201809211611c0657565b6000198114611c065760010190565b80548210156121635760005260206000200190600090565b929394919060036133578567ffffffffffffffff166000526004602052604060002090565b01936001600160a01b0361336c8184166127e8565b16918215613576576040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c6357600091613557575b50156135475761341d600095969798604051998a96879586957f89720a620000000000000000000000000000000000000000000000000000000087526004870161326c565b03915afa928315610c6357600093613522575b508251156135175782519061344f61344a845480946132fe565b6120fa565b906000928394845b87518110156134b65761346d61164e828a6129e0565b6001600160a01b038116156134aa57906134a460019261349661348f8a61330b565b99896129e0565b906001600160a01b03169052565b01613457565b509550600180966134a4565b5091955091936134c8575b5050815290565b60005b8281106134d857506134c1565b806135116134fe6134eb6001948661331a565b90546001600160a01b039160031b1c1690565b61349661350a8861330b565b97896129e0565b016134cb565b9150610daa90611e5d565b6135409193503d806000833e6135388183610214565b8101906131e8565b9138613430565b505050509250610daa9150611e5d565b613570915060203d602011611b9557611b878183610214565b386133d8565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045260246000fd5b939192936135c86135c082518651906132fe565b8651906132fe565b906135db6135d5836120fa565b92612a28565b94600096875b8351891015613641578861363761362a60019361361261360861164e8e9f9d9e9d8b6129e0565b613496838c6129e0565b61363061361f858c6129e0565b51918093849161330b565b9c6129e0565b528b6129e0565b50019796956135e1565b959250929350955060005b86518110156136d55761366261164e82896129e0565b60006001600160a01b038216815b8881106136a9575b505090600192911561368c575b500161364c565b6136a39061349661369c8961330b565b98886129e0565b38613685565b816136ba6103f861164e848c6129e0565b146136c757600101613670565b506001915081905038613678565b509390945060005b8551811015613766576136f361164e82886129e0565b60006001600160a01b038216815b87811061373a575b505090600192911561371d575b50016136dd565b6137349061349661372d8861330b565b97876129e0565b38613716565b8161374b6103f861164e848b6129e0565b1461375857600101613701565b506001915081905038613709565b50828252918252925090565b6040519060a0820182811067ffffffffffffffff8211176101bb5760405260606080836000815260006020820152600060408201526000838201520152565b906137bb826120c6565b6137c86040519182610214565b828152601f196137d882946120c6565b019060005b8281106137e957505050565b6020906137f4613772565b828285010152016137dd565b9081606091031261016a57805161381681612275565b9160406020830151613827816121d1565b920151610daa816121d1565b9160209082815201919060005b81811061384d5750505090565b9091926040806001926001600160a01b03873561386981610aeb565b16815260208781013590820152019401929101613840565b949391929067ffffffffffffffff168552608060208601526138da6138bb6138a985806125f7565b60a060808a0152610120890191612418565b6138c860208601866125f7565b90607f198984030160a08a0152612418565b6040840135601e198536030181121561016a578401916020833593019167ffffffffffffffff841161016a578360061b3603831361016a576102479561396261393983606097613983978d60c0607f19826139759a0301910152613833565b91613958613948888301610fca565b6001600160a01b031660e08d0152565b60808101906125f7565b90607f198b8403016101008c0152612418565b9087820360408901526102cb565b94019061ffff169052565b9063ffffffff8091169116019063ffffffff8211611c0657565b908160a091031261016a5780519160208201516139c4816121d1565b9160408101516139d3816121d1565b91608060608301516139e481612275565b920151610daa81610fd5565b9260c0946001600160a01b039167ffffffffffffffff61ffff9584610daa9b9a9616885216602087015260408601521660608401521660808201528160a082015201906102cb565b9081606091031261016a578051613816816121d1565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8201918211611c0657565b906000198201918211611c0657565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201918211611c0657565b91908203918211611c0657565b919082608091031261016a578151613adb816121d1565b916020810151916060604083015192015190565b8115613af9570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9382946000906000956040810194613b62613b5d613b58885151613b5060408c01809c611df4565b9190506132fe565b6132b8565b6137b1565b9660009586955b88518051881015613dc6576103f86103f861164e8a613b87946129e0565b613bd460206060880192613b9c8b85516129e0565b51908a6040518095819482937f958021a700000000000000000000000000000000000000000000000000000000845260048401612a71565b03915afa8015610c63576001600160a01b0391600091613da8575b50168015613d54579060608e9392613c088b84516129e0565b5190613c1960208b015161ffff1690565b958b613c54604051988995869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613881565b03915afa8015610c6357600193613cf1938b8f8f95600080958197613cfa575b509083929161ffff613c9c85613c9561164e613ce599613ceb9d9e516129e0565b94516129e0565b5191613cb8613ca9610259565b6001600160a01b039095168552565b63ffffffff8916602085015263ffffffff8b166040850152166060830152608082015261174883836129e0565b5061398e565b9961398e565b96019596613b69565b613ceb975061164e965084939291509361ffff613c9c82613c95613d37613ce59960603d8111613d4d575b613d2f8183610214565b810190613800565b9c9196909c9d5050505050505090919293613c74565b503d613d25565b6105a788613d6661164e8c8f516129e0565b7f83c758a6000000000000000000000000000000000000000000000000000000006000526001600160a01b031660045267ffffffffffffffff16602452604490565b613dc0915060203d81116117f6576117e88183610214565b38613bef565b50919a9496929395509897968a613ddd8187611df4565b90506140f1575b50508651613df190613a4e565b99613dff6020860186611e2a565b91613e0b915086611df4565b9560609150019486613e1c87612168565b91613e27938a615994565b613e318b896129e0565b52613e3c8a886129e0565b50613e478a886129e0565b516020015163ffffffff16613e5b9161398e565b90613e668a886129e0565b516040015163ffffffff16613e7a9161398e565b91613e83610259565b33815290600060208301819052604083015261ffff166060820152613ea6610293565b60808201528651613eb690613a7b565b90613ec182896129e0565b52613ecc90876129e0565b506002546001600160a01b031692613ee390612168565b6040517f910d8f5900000000000000000000000000000000000000000000000000000000815267ffffffffffffffff96909616600487015263ffffffff9182166024870152911660448501526001600160a01b03166064840152826084815a93608094fa958615610c63576000966000976000946000926140b4575b506000965b865188101561403457613fbd600191613f8588613f80876122c4565b613aef565b613f9e6060613f948d8d6129e0565b510191825161236c565b9052858a14613fc5575b6060613fb48b8b6129e0565b510151906132fe565b970196613f64565b8b8873eba517d2000000000000000000000000000000006001600160a01b03613ff860808c01516001600160a01b031690565b1603614006575b5050613fa8565b613f80614012926122f4565b61402b60606140218d8d6129e0565b51019182516132fe565b90528b88613fff565b97965097505050506140707f000000000000000000000000000000000000000000000000000000000000000091613f8063ffffffff84166122f4565b841161407c5750929190565b7f25c2df0a00000000000000000000000000000000000000000000000000000000600052600484905263ffffffff1660245260446000fd5b92985050506140dc91925060803d6080116140ea575b6140d48183610214565b810190613ac4565b919790939290919038613f5f565b503d6140ca565b610b276103f86104f86104f261410a948a989698611df4565b926001600160a01b03600091515194169060e0880190815161412a610259565b6001600160a01b0385168152908260208301528260408301528260608301526080820152614158878d6129e0565b52614163868c6129e0565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3317103100000000000000000000000000000000000000000000000000000000600482015291602083602481875afa8015610c63578f948c89968f96948d948f968891614478575b50614371575b50505050505015614218575b61193f61420961421095614203602061193f614203976040976129e0565b9061398e565b958b6129e0565b90388a613de4565b505061429e9160608c6142496104f86104f26142426103f86103f86002546001600160a01b031690565b938b611df4565b6040517f947f821700000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8c1660048201526001600160a01b03909116602482015294859190829081906044820190565b03915afa908115610c635761193f614209604092614203602061193f8f8b906142109b6142039a600080600092614333575b63ffffffff9293506143209060606142e888886129e0565b5101946143158a6142f98a8a6129e0565b51019160406143088b8b6129e0565b51019063ffffffff169052565b9063ffffffff169052565b16905297509750505050955050506141e5565b50505063ffffffff61435f6143209260603d60601161436a575b6143578183610214565b810190613a38565b9093509150826142d0565b503d61434d565b8495985060a09697506143bb60206143b16060826143a86104f26143a16104f86104f28b6143f49c9d9e9f611df4565b998d611df4565b01359901612168565b9a015161ffff1690565b905190604051998a97889687967f2c063404000000000000000000000000000000000000000000000000000000008852600488016139f0565b03915afa8015610c63578592828c93918190829461443c575b5061443090606061441e88886129e0565b51019261431560206142f98a8a6129e0565b5288888f8c81386141d9565b9150506144309250614466915060a03d60a011614471575b61445e8183610214565b8101906139a8565b94919291905061440d565b503d614454565b614491915060203d602011611b9557611b878183610214565b386141d3565b6001600160a01b036001541633036144ab57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b8051916144e38151846132fe565b9283156145dc5760005b8481106144fb575050505050565b818110156145c15761451061164e82866129e0565b6001600160a01b0381168015610c6857614529836132c6565b87811061453b575050506001016144ed565b8481101561459e576001600160a01b0361455861164e838a6129e0565b16821461456757600101614529565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b03831660045260246000fd5b6001600160a01b036145bc61164e6145b68885613ab7565b896129e0565b614558565b6145d761164e6145d18484613ab7565b856129e0565b614510565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b91908110156121635760051b0190565b9081602091031261016a575190565b9160806102479294936146758160e08101976001600160a01b036060809267ffffffffffffffff815116855282602082015116602086015263ffffffff6040820151166040860152015116910152565b01906001600160a01b0360408092828151168552602081015115156020860152015116910152565b906020610daa9281815201906102cb565b906146b882610277565b6146c56040519182610214565b828152601f196121218294610277565b9081516020821080614755575b61472457036146ee5790565b610a24906040519182917f3aeba3900000000000000000000000000000000000000000000000000000000083526004830161469d565b50906020810151908161473684612326565b1c6146ee5750614745826146ae565b9160200360031b1b602082015290565b50602081146146e2565b918251601481029080820460141490151715611c0657614781614786916132d4565b6132e2565b90614798614793836132f0565b6146ae565b9060146147a4836129d3565b5360009260215b86518510156147d65760146001916147c661164e888b6129e0565b60601b81870152019401936147ab565b919550936020935060601b90820152828152012090565b906147fd6103f860608401612168565b61480e600019936040810190611df4565b9050614886575b61481f8251613a7b565b9260005b848110614831575050505050565b80826001921461488157606061484782876129e0565b510151801561487b576148759061486f61486184896129e0565b51516001600160a01b031690565b86615b9d565b01614823565b50614875565b614875565b91506148928151613a8a565b916148a061486184846129e0565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f331710310000000000000000000000000000000000000000000000000000000060048201526020816024816001600160a01b0386165afa908115610c6357600091614939575b50614919575b50614815565b61493390606061492986866129e0565b5101519083615b9d565b38614913565b614952915060203d602011611b9557611b878183610214565b3861490d565b60405190614965826101f8565b60606020838281520152565b919060408382031261016a576040519061498a826101f8565b8193805167ffffffffffffffff811161016a57826149a9918301612a92565b835260208101519167ffffffffffffffff831161016a576020926149cd9201612a92565b910152565b9060208282031261016a57815167ffffffffffffffff811161016a57610daa9201614971565b9060806001600160a01b0381614a17855160a0865260a08601906102cb565b9467ffffffffffffffff60208201511660208601528260408201511660408601526060810151606086015201511691015290565b906020610daa9281815201906149f8565b919060408382031261016a57825167ffffffffffffffff811161016a57602091614a87918501614971565b92015190565b61ffff614aa6610daa95936060845260608401906149f8565b9316602082015260408184039101526102cb565b909193929594614ac86128f2565b5060208201805115614f6257614aeb610b276103f885516001600160a01b031690565b946001600160a01b038616916040517f01ffc9a700000000000000000000000000000000000000000000000000000000815260208180614b5260048201907faff2afbf00000000000000000000000000000000000000000000000000000000602083019252565b0381875afa908115610c6357600091614f43575b5015614efa57614bde88999a825192614b7d614958565b5051614bc9614b9389516001600160a01b031690565b926040614b9e610259565b9e8f908152614bba8d602083019067ffffffffffffffff169052565b01906001600160a01b03169052565b60608c01526001600160a01b031660808b0152565b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f33171031000000000000000000000000000000000000000000000000000000006004820152602081602481875afa908115610c6357600091614edb575b5015614de35750906000929183614c8a9899604051998a95869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501614a8d565b03925af18015610c6357600094600091614d9b575b50614d70614cd99596614d13614ce7614d059561164e6020969b995b6040519b8c918983019190916001600160a01b036020820193169052565b03601f1981018c528b610214565b6040519586918683019190916001600160a01b036020820193169052565b03601f198101865285610214565b614d4f61096d614d45614d558951614d4f61096d614d458c67ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b906146d5565b9767ffffffffffffffff166000526004602052604060002090565b93015193614d7c610268565b958652602086015260408501526060840152608083015260a082015290565b614cd995506020915061164e96614d13614ce7614d0595614dd1614d70953d806000833e614dc98183610214565b810190614a5c565b9b909b96505095505050969550614c9f565b9793929061ffff16614eb15751614e8757614e326000939184926040519586809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260048301614a4b565b03925af1908115610c6357614cd995614d13614ce7614d709361164e602096614d0598600091614e64575b5099614cbb565b614e8191503d806000833e614e798183610214565b8101906149d2565b38614e5d565b7f9218ad0a0000000000000000000000000000000000000000000000000000000060005260046000fd5b7fbf2937150000000000000000000000000000000000000000000000000000000060005260046000fd5b614ef4915060203d602011611b9557611b878183610214565b38614c42565b6105a7614f0e86516001600160a01b031690565b7fbf16aab6000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b614f5c915060203d602011611b9557611b878183610214565b38614b66565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b959261505e947fffffffffffffffff0000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000094928186948160439d9b977f01000000000000000000000000000000000000000000000000000000000000008e5260c01b1660018d015260c01b1660098b015260c01b16601189015260e01b16601987015260e01b16601d85015260218401907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60238201520190565b9061507a602092828151948592016102a8565b0190565b9360019694937fff000000000000000000000000000000000000000000000000000000000000008094899896828a9660f81b1681526150c682518093602089850191016102a8565b019160f81b16838201526150e48251809360206002850191016102a8565b01019160f81b16838201526151038251809360206002850191016102a8565b01010190565b60017fffff000000000000000000000000000000000000000000000000000000000000956002958760049a9681957fff00000000000000000000000000000000000000000000000000000000000000610daa9f9e9c9860f81b16815261517882518093602089850191016102a8565b019160f01b16838201526151968251809360206003850191016102a8565b01019160f01b16838201526151b482518093602089850191016102a8565b01019160f01b1660028201520190615067565b60e081019060ff8251511161546e57610100810160ff815151116154585761012082019260ff8451511161544257610140830160ff8151511161542c5761016084019061ffff82515111615416576101808501936001855151116153fe576101a086019361ffff855151116153e8576060955180516153ac575b50865167ffffffffffffffff16602088015167ffffffffffffffff169060408901516152749067ffffffffffffffff1690565b9860608101516152879063ffffffff1690565b90608081015161529a9063ffffffff1690565b60a082015161ffff169160c00151926040519c8d9660208801966152bd97614f8c565b03601f19810188526152cf9088610214565b519081516152dd9060ff1690565b9051805160ff1698519081516152f39060ff1690565b906040519a8b9560208701956153089661507e565b03601f198101875261531a9087610214565b519182516153289060ff1690565b9151805161ffff1694805161533e9061ffff1690565b925192835161534e9061ffff1690565b92604051978897602089019761536398615109565b03601f19810182526153759082610214565b6040519283926020840161538891615067565b61539191615067565b61539a91615067565b03601f1981018252610daa9082610214565b6153c19196506153bb906129d3565b51615d94565b9461ffff8651116153d25738615241565b635a102da160e11b600052602560045260246000fd5b635a102da160e11b600052602660045260246000fd5b635a102da160e11b6000526105a76024906024600452565b635a102da160e11b600052602360045260246000fd5b635a102da160e11b600052602260045260246000fd5b635a102da160e11b600052602160045260246000fd5b635a102da160e11b600052602060045260246000fd5b635a102da160e11b600052601f60045260246000fd5b9061548d612e62565b91601182106156a35780357f302326cb000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216036156305750600481013560e01c8352600881013560f01c6020840152600b600a82013560001a61550c816120fa565b6040860190815261551c82612a28565b906060870191825260005b8381106155e4575050505061558f838361558561557961556f61556861555561559998876155a39c9b615ebb565b6001600160a01b0390911660808d015290565b8585615f91565b92919036916128bb565b60a08a01528383615ff9565b94919036916128bb565b60c0880152615f91565b93919036916128bb565b60e084015281036155b2575090565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600360045260245260446000fd5b8060019161562961561361560c6155ff6156239a8d8d615ebb565b9190613496868a516129e0565b8b8b615f91565b9391889a919a51949a36916128bb565b926129e0565b5201615527565b7f55a0e02c000000000000000000000000000000000000000000000000000000006000527f302326cb000000000000000000000000000000000000000000000000000000006004527fffffffff000000000000000000000000000000000000000000000000000000001660245260446000fd5b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526002600452602482905260446000fd5b80519060005b8281106156e957505050565b60018101808211611c06575b83811061570557506001016156dd565b6001600160a01b0361571783856129e0565b51166157296103f861164e84876129e0565b14615736576001016156f5565b6105a761574661164e84866129e0565b7fa1726e40000000000000000000000000000000000000000000000000000000006000526001600160a01b0316600452602490565b909391928151156159075750615790816156d7565b80519260005b8481106157a557509093925050565b6157b56103f861164e83866129e0565b156157c257600101615796565b9391946157dc61344a6157d485613a7b565b8451906132fe565b926157f96157f46157ec83613a7b565b8551906132fe565b612a28565b968792600097885b8481106158a95750505050505060005b815181101561589c576000805b86811061585d575b50906001916158585761585261583f61164e83866129e0565b61349661584b8961330b565b98876129e0565b01615811565b615852565b61586a61164e82876129e0565b6001600160a01b036158826103f861164e87896129e0565b9116146158915760010161581e565b506001905080615826565b5050909180825283529190565b9091929394988282146158fd57906158ef6158e2836158d58b61349660019761562361164e898e6129e0565b6158e861361f85896129e0565b9e6129e0565b528c6129e0565b505b01908994939291615801565b98506001906158f1565b919350501561592257506159196120de565b90610daa6129f4565b90610daa8251612a28565b9081602091031261016a5751610daa81612275565b9361597f60809461ffff6001600160a01b039567ffffffffffffffff61598d969b9a9b16895216602088015260a0604088015260a0870190610c92565b9085820360608701526102cb565b9416910152565b929190926159a0613772565b506159bf8167ffffffffffffffff166000526004602052604060002090565b805490959060e01c60ff1691608085019283516159e2906001600160a01b031690565b60019098015460101c63ffffffff16865163ffffffff16615a029161398e565b96615a0e90608d6132fe565b9460a0870195865151615a20916132fe565b9160ff1691615a2e8361233c565b615a37916132fe565b91615a439060676132fe565b615a4c9161236c565b615a55916132fe565b63ffffffff1692516001600160a01b03169473eba517d2000000000000000000000000000000006001600160a01b03871603615ae25750505061ffff9250615ad490615ac76000935b5195615aba615aab610259565b6001600160a01b039099168952565b63ffffffff166020880152565b63ffffffff166040860152565b166060830152608082015290565b615af86103f8602094976001600160a01b031690565b906040615b098583015161ffff1690565b91015192615b49875198604051998a96879586957ff238895800000000000000000000000000000000000000000000000000000000875260048701615942565b03915afa908115610c6357615ac7615ad49261ffff95600091615b6e575b5093615a9e565b615b90915060203d602011615b96575b615b888183610214565b81019061592d565b38615b67565b503d615b7e565b91602091600091604051906001600160a01b03858301937fa9059cbb000000000000000000000000000000000000000000000000000000008552166024830152604482015260448152615bf1606482610214565b519082855af1156127dc576000513d615c5057506001600160a01b0381163b155b615c195750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415615c12565b966002615d6897615d356022610daa9f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090615d359f82615d359c615d3c9c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b166021820152615ce682518093602089850191016102a8565b019160f81b1683820152615d048251809360206023850191016102a8565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190615067565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b602081019060ff82515111615ea557604081019160ff83515111615e8f57606082019160ff83515111615e7957608081019260ff84515111615e635760a0820161ffff81515111615e4d57610daa94615e3f9351945191615df6835160ff1690565b975191615e04835160ff1690565b945190615e12825160ff1690565b905193615e20855160ff1690565b935196615e2f885161ffff1690565b966040519c8d9b60208d01615c59565b03601f198101835282610214565b635a102da160e11b600052602b60045260246000fd5b635a102da160e11b600052602a60045260246000fd5b635a102da160e11b600052602960045260246000fd5b635a102da160e11b600052602860045260246000fd5b635a102da160e11b600052602760045260246000fd5b929190926001820191848311615f5f5781013560001a828115615f54575060148103615f27578201938411615ef357013560601c9190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602482905260446000fd5b7f6d1eca280000000000000000000000000000000000000000000000000000000060005260045260246000fd5b945050505060009190565b7fd9437f9d00000000000000000000000000000000000000000000000000000000600052600060045260245260446000fd5b91906002820191818311615f5f578381013560f01c0160020192818411615fc557918391615fbe93612ec1565b9290929190565b7fd9437f9d000000000000000000000000000000000000000000000000000000006000526001600452602483905260446000fd5b91906001820191818311615f5f578381013560001a0160010192818411615fc557918391615fbe93612ec156fea164736f6c634300081a000a",
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
