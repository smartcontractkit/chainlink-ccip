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
	Router             common.Address
	SequenceNumber     uint64
	AddressBytesLength uint8
	DefaultExecutor    common.Address
	LaneMandatedCCVs   []common.Address
	DefaultCCVs        []common.Address
	OffRamp            []byte
}

type OnRampDestChainConfigArgs struct {
	DestChainSelector  uint64
	Router             common.Address
	AddressBytesLength uint8
	DefaultCCVs        []common.Address
	LaneMandatedCCVs   []common.Address
	DefaultExecutor    common.Address
	OffRamp            []byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct OnRamp.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"destChainConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DestChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contract IRouter\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateDestChainAddress\",\"inputs\":[{\"name\":\"rawAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"addressBytesLength\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"validatedAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"encodedMessage\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receipts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"struct OnRamp.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"verifierBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contract IRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct OnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IRouter\"},{\"name\":\"defaultCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"laneMandatedCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"offRamp\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DestinationChainNotSupported\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataLength\",\"inputs\":[{\"name\":\"location\",\"type\":\"uint8\",\"internalType\":\"enum MessageV1Codec.EncodingErrorLocation\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainAddress\",\"inputs\":[{\"name\":\"destChainAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSpecifyDefaultOrRequiredCCVs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0604052346102f55760405161590d38819003601f8101601f191683016001600160401b038111848210176102fa578392829160405283398101039060c082126102f557606082126102f557610054610310565b81519092906001600160401b03811681036102f55783526020820151906001600160a01b03821682036102f5576020840191825260606100966040850161032f565b6040860190815291605f1901126102f5576100af610310565b916100bc6060850161032f565b835260808401519384151585036102f55760a06100e091602086019687520161032f565b946040840195865233156102e457600180546001600160a01b0319163317905580516001600160401b03161580156102d2575b80156102c0575b61029357516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102ae575b80156102a4575b6102935780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610310565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610310565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a16040516155c990816103448239608051818181610f4d015281816115e60152611bac015260a05181611be5015260c051818181611c2101526120d00152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b038111838210176102fa57604052565b51906001600160a01b03821682036102f55756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610117578063181f5a771461011257806320487ded1461010d57806348a98aa4146101085780635cb80c5d146101035780636d7fa1ce146100fe5780636def4ce7146100f95780637437ff9f146100f457806378802634146100ef57806379ba5097146100ea5780638da5cb5b146100e55780639041be3d146100e057806390423fa2146100db578063df0aa9e9146100d6578063f2fde38b146100d15763fbca3b74146100cc57600080fd5b611b09565b611a15565b611325565b611151565b6110af565b61105d565b610f74565b610c26565b610b5a565b610abb565b6108ea565b610791565b6106a6565b61041b565b61037a565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576060610150611b8c565b610193604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176101e557604052565b61019a565b6060810190811067ffffffffffffffff8211176101e557604052565b6040810190811067ffffffffffffffff8211176101e557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101e557604052565b6040519061027361018083610222565b565b6040519061027360e083610222565b60405190610273604083610222565b6040519061027360a083610222565b6040519061027360c083610222565b67ffffffffffffffff81116101e557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906102fa602083610222565b60008252565b60005b8381106103135750506000910152565b8181015183820152602001610303565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361035f81518092818752878088019101610300565b0116010190565b906020610377928181520190610323565b90565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576103f760408051906103bb8183610222565b601082527f4f6e52616d7020312e372e302d64657600000000000000000000000000000000602083015251918291602083526020830190610323565b0390f35b67ffffffffffffffff81160361019557565b908160a09103126101955790565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557600435610456816103fb565b60243567ffffffffffffffff81116101955761047690369060040161040d565b6000906104978367ffffffffffffffff166000526004602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff6104e76104ce845473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b161561064c5761055092936105496105156105056080850185611c49565b9061050f87611e0e565b856131b6565b9361051e611ef4565b906040850161052d8187611f5f565b90506105f7575b50610543600287519201611c9a565b906138e6565b8352613f63565b6000916000916000915b81518310156105ec576105e26105ba610584600193606061057b8888612007565b51015190612058565b966105b46105a760206105978989612007565b51015167ffffffffffffffff1690565b67ffffffffffffffff1690565b90612058565b946105b46105d960406105cd8888612007565b51015163ffffffff1690565b63ffffffff1690565b920191929361055a565b604051908152602090f35b610645919250602061062761061561062061061b610615868c611f5f565b90611fe2565b611ff0565b9389611f5f565b0135610638602089015161ffff1690565b9060c0890151928761367b565b9038610534565b7fbff66cef0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff841660045260246000fd5b6000fd5b73ffffffffffffffffffffffffffffffffffffffff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610195576106e06004356103fb565b60206106f66024356106f181610688565b612071565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b9060207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101955760043567ffffffffffffffff81116101955760040160009280601f8301121561078d5781359367ffffffffffffffff851161078a57506020808301928560051b010111610195579190565b80fd5b8380fd5b346101955761079f36610714565b906107bf60035473ffffffffffffffffffffffffffffffffffffffff1690565b9160005b8181106107cc57005b6107dd6104ce61061b83858761216a565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa9081156108da5760019488916000936108aa575b5082610852575b50505050016107c3565b61085d918391614320565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681610848565b6108cc91935060203d81116108d3575b6108c48183610222565b81019061217a565b9138610841565b503d6108ba565b612065565b60ff81160361019557565b346101955760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760043567ffffffffffffffff8111610195573660238201121561019557806004013567ffffffffffffffff8111610195573660248284010111610195576103f79161097291602480359261096c846108df565b016122cb565b60405191829182610366565b906020808351928381520192019060005b81811061099c5750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161098f565b90610377916020815273ffffffffffffffffffffffffffffffffffffffff825116602082015267ffffffffffffffff6020830151166040820152610a166040830151606083019060ff169052565b606082015173ffffffffffffffffffffffffffffffffffffffff16608082015260c0610a88610a55608085015160e060a086015261010085019061097e565b60a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858303018486015261097e565b9201519060e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610323565b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff600435610aff816103fb565b606060c0604051610b0f816101c9565b6000815260006020820152600060408201526000838201528260808201528260a082015201521660005260046020526103f7610b4e6040600020611e0e565b604051918291826109c8565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557610b91611b6d565b50604051610b9e816101ea565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015260405180916103f782606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b3461019557610c3436610714565b90610c3d614420565b6000915b808310610c4a57005b610c558382846123f2565b92610c5f84612432565b9367ffffffffffffffff85169081158015610f41575b8015610f2b575b610ef357610cc7939495610ce16060830191610c988385612446565b9790610cc16080870199610cb9610caf8c8a612446565b949092369161249a565b92369161249a565b9061446b565b67ffffffffffffffff166000526004602052604060002090565b6020830190610d33610cf283611ff0565b829073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b610d90610d426040860161243c565b82547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660e09190911b7cff0000000000000000000000000000000000000000000000000000000016178255565b610da7610d9d8486612446565b9060038401612507565b610dbe610db48886612446565b9060028401612507565b60a08401610dce6104ce82611ff0565b15610ec9576001977f5ba821cbe35d9e1dec357bb6a26f019c75c549223460f8a23321af7431e5e6ae95610ea983610eaf610ebe96610e55610e12610ea198611ff0565b8f83019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b610e98610e91610e8b60c0880193610e7a610e70868b611c49565b90600484016125ce565b5460a01c67ffffffffffffffff1690565b9a611ff0565b9a86612446565b97909686612446565b949093611ff0565b94611c49565b959094604051998a998a612747565b0390a2019190610c41565b7f35be3ac80000000000000000000000000000000000000000000000000000000060005260046000fd5b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b5060ff610f3a6040830161243c565b1615610c7c565b5067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168214610c75565b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955760005473ffffffffffffffffffffffffffffffffffffffff81163303611033577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955767ffffffffffffffff6004356110f3816103fb565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff81116111375760405167ffffffffffffffff9091168152602090f35b61201b565b359061027382610688565b8015150361019557565b346101955760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557600060405161118e816101ea565b60043561119a81610688565b81526024356111a881611147565b60208201908152604435906111bc82610688565b604083019182526111cb614420565b73ffffffffffffffffffffffffffffffffffffffff83511615918215611305575b5081156112fa575b506112d25780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3906112bd611b8c565b6112cc604051928392836145fa565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b9050511515386111f4565b5173ffffffffffffffffffffffffffffffffffffffff16159150386111ec565b346101955760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955761135f6004356103fb565b60243567ffffffffffffffff81116101955761137f90369060040161040d565b60443561138d606435610688565b60025460a01c60ff166119eb576113de740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b6113fe60043567ffffffffffffffff166000526004602052604060002090565b9173ffffffffffffffffffffffffffffffffffffffff60643516156119c1578254926114406104ce73ffffffffffffffffffffffffffffffffffffffff861681565b3303611997576114536080830183611c49565b61145c83611e0e565b61146992906004356131b6565b9360a01c67ffffffffffffffff16611480906127b5565b81547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660a082901b7bffffffffffffffff0000000000000000000000000000000000000000161782556040513060601b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208201526014815290949061150a603482610222565b602082015161ffff16604083015163ffffffff1660405160643560601b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166020820152909490928360348101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810185526115879085610222565b6115918780611c49565b835460e01c60ff16906115a3926122cb565b92608086015160408901956115b8878b611f5f565b6115c29150612816565b986115d060208c018c611c49565b9490956115db610263565b67ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001681529e8f9067ffffffffffffffff60043516602083015267ffffffffffffffff166040820152606001528d61163d60048901611d4e565b608082015260a00190611653919061ffff169052565b63ffffffff1660c08d015260e08c01526101008b019485526101208b01526101408a01968752369061168492612294565b610160890152611692611ef4565b61169c8488611f5f565b6116b593915061194d575b610543600287519201611c9a565b83526116bf612883565b936116cd8487600435613f63565b92602086019384526116df8188611f5f565b90506118b0575b505050926116f386614e76565b808452602081519101206117088351516128d3565b9360408101948552606060009301925b8451805182101561180357906000868686838e6117ad8f986117748961176d6117676104ce6104ce61174c85602098612007565b515173ffffffffffffffffffffffffffffffffffffffff1690565b98611ff0565b9851612007565b510151604051998a97889687957f97048c7100000000000000000000000000000000000000000000000000000000875260048701612a82565b03925af180156108da57816117db916001946000916117e2575b508951906117d58383612007565b52612007565b5001611718565b6117fd913d8091833e6117f58183610222565b81019061297c565b386117c7565b6103f78480898b7f9b8fdf7fa94e7e8692c830c07cc6ce91a34c507d9f8efea07eb71cd64ed4891f8f8961187061184e604067ffffffffffffffff94015167ffffffffffffffff1690565b915194519551604051938493169667ffffffffffffffff600435169684612d0e565b0390a46118a07fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b6040519081529081906020820190565b60016118bc8289611f5f565b90500361192357611912926118d7610615611900938a611f5f565b60a08701518051919290911561191b57505b606435916118fb6004359136906128a2565b61477b565b90519061190c82611ffa565b52611ffa565b503880806116e6565b9050516118e9565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b905061199161196261061b610615878b611f5f565b6020611971610615888c611f5f565b0135611982602089015161ffff1690565b9060c08901519260043561367b565b906116a7565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101955773ffffffffffffffffffffffffffffffffffffffff600435611a6581610688565b611a6d614420565b16338114611adf57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101955760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019557611b436004356103fb565b7f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611b7a826101ea565b60006040838281528260208201520152565b611b94611b6d565b50604051611ba1816101ea565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff82116101955760200191813603831361019557565b906040519182815491828252602082019060005260206000209260005b818110611ccc57505061027392500383610222565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201611cb7565b90600182811c92168015611d44575b6020831014611d1557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691611d0a565b9060405191826000825492611d6284611cfb565b8084529360018116908115611dce5750600114611d87575b5061027392500383610222565b90506000929192526020600020906000915b818310611db25750509060206102739282010138611d7a565b6020919350806001915483858901015201910190918492611d99565b602093506102739592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138611d7a565b90604051611e1b816101c9565b60c0611ed760048395611e6e611e64825473ffffffffffffffffffffffffffffffffffffffff8116885267ffffffffffffffff808260a01c1616602089015260ff9060e01c1690565b60ff166040870152565b611eaf611e92600183015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff166060870152565b611ebb60028201611c9a565b6080860152611ecc60038201611c9a565b60a086015201611d4e565b910152565b67ffffffffffffffff81116101e55760051b60200190565b60405190611f03602083610222565b6000808352366020840137565b90611f1a82611edc565b611f276040519182610222565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611f558294611edc565b0190602036910137565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160061b3603831361019557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9015611feb5790565b611fb3565b3561037781610688565b805115611feb5760200190565b8051821015611feb5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161113757565b9190820180921161113757565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156108da5760009061211b575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011612162575b8161213560209383610222565b810103126101955773ffffffffffffffffffffffffffffffffffffffff905161215d81610688565b612100565b3d9150612128565b9190811015611feb5760051b0190565b90816020910312610195575190565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916020610377938181520191612189565b60ff166020039060ff821161113757565b909291928311610195579190565b906004116101955790600490565b909291928360041161019557831161019557600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b90939293848311610195578411610195578101920390565b359060208110612267575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b9291926122a0826102b1565b916122ae6040519384610222565b829481845281830111610195578281602093846000960137010152565b9160ff81169060208210612327575b5081036122ed5790610377913691612294565b906123236040519283927f3aeba390000000000000000000000000000000000000000000000000000000008452600484016121c8565b0390fd5b602083116123bd57602083036122da57905061235b61235560ff61234d849695966121d9565b1685856121ea565b90612259565b6123875791612380918161237a612374610377966121d9565b60ff1690565b91612241565b3691612294565b506123236040519283927f3aeba390000000000000000000000000000000000000000000000000000000008452600484016121c8565b6040517f3aeba390000000000000000000000000000000000000000000000000000000008152806123238587600484016121c8565b9190811015611feb5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2181360301821215610195570190565b35610377816103fb565b35610377816108df565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610195570180359067ffffffffffffffff821161019557602001918160051b3603831361019557565b9291906124a681611edc565b936124b46040519586610222565b602085838152019160051b810192831161019557905b8282106124d657505050565b6020809183356124e581610688565b8152019101906124ca565b8181106124fb575050565b600081556001016124f0565b9067ffffffffffffffff83116101e5576801000000000000000083116101e557815483835580841061256b575b5090600052602060002060005b83811061254e5750505050565b600190602084359461255f86610688565b01938184015501612541565b612583908360005284602060002091820191016124f0565b38612534565b9190601f811161259857505050565b610273926000526020600020906020601f840160051c830193106125c4575b601f0160051c01906124f0565b90915081906125b7565b90929167ffffffffffffffff81116101e5576125f4816125ee8454611cfb565b84612589565b6000601f8211600114612652578190612643939495600092612647575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055565b013590503880612611565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169461268584600052602060002090565b91805b8781106126de5750836001959697106126a6575b505050811b019055565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88560031b161c1991013516905538808061269c565b90926020600181928686013581550194019101612688565b9160209082815201919060005b8181106127105750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff873561273981610688565b168152019401929101612703565b94916127a19373ffffffffffffffffffffffffffffffffffffffff95866127939367ffffffffffffffff6103779e9c9d9b96168a5216602089015260c0604089015260c08801916126f6565b9185830360608701526126f6565b9416608082015260a0818503910152612189565b67ffffffffffffffff1667ffffffffffffffff81146111375760010190565b6040519060c0820182811067ffffffffffffffff8211176101e557604052606060a0836000815282602082015282604082015282808201528260808201520152565b9061282082611edc565b61282d6040519182610222565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061285b8294611edc565b019060005b82811061286c57505050565b6020906128776127d4565b82828501015201612860565b60405190612890826101ea565b60606040838281528260208201520152565b9190826040910312610195576040516128ba81610206565b602080829480356128ca81610688565b84520135910152565b906128dd82611edc565b6128ea6040519182610222565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06129188294611edc565b019060005b82811061292957505050565b80606060208093850101520161291d565b81601f82011215610195578051612950816102b1565b9261295e6040519485610222565b81845260208284010111610195576103779160208085019101610300565b9060208282031261019557815167ffffffffffffffff811161019557610377920161293a565b9080602083519182815201916020808360051b8301019401926000915b8383106129ce57505050505090565b9091929394602080612a73837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951908151815260a0612a62612a50612a3e612a2c8887015160c08a88015260c0870190610323565b60408701518682036040880152610323565b60608601518582036060870152610323565b60808501518482036080860152610323565b9201519060a0818403910152610323565b970193019301919392906129bf565b9193906103779593612c60612c859260a08652612aac60a08701825167ffffffffffffffff169052565b602081015167ffffffffffffffff1660c0870152604081015167ffffffffffffffff1660e0870152610160612c2d612bf7612bc1612b8b612b38612b038c61022060608a0151916101806101008201520190610323565b60808801518d82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101208f0152610323565b60a087015161ffff166101408d015260c087015163ffffffff16868d015260e08701518c82037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60016101808e0152610323565b6101008601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608c8303016101a08d0152610323565b6101208501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608b8303016101c08c0152610323565b6101408401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608a8303016101e08b01526129a2565b9101517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6087830301610200880152610323565b956020850152604084019073ffffffffffffffffffffffffffffffffffffffff169052565b60608201526080818403910152610323565b9080602083519182815201916020808360051b8301019401926000915b838310612cc357505050505090565b9091929394602080612cff837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610323565b97019301930191939290612cb4565b939290612d2390606086526060860190610323565b938085036020820152825180865260208601906020808260051b8901019501916000905b828210612d6557505050506103779394506040818403910152612c97565b90919295602080612df5837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08d6001960301865260a060808c5173ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff86820151168685015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610323565b980192019201909291612d47565b60405190612e10826101c9565b606060c08382815260006020820152600060408201526000838201528260808201528260a08201520152565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612e70575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90612eac82611edc565b612eb96040519182610222565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612ee78294611edc565b019060005b828110612ef857505050565b602090604051612f0781610206565b6000815260608382015282828501015201612eec565b63ffffffff81160361019557565b9160608383031261019557825167ffffffffffffffff81116101955782612f5391850161293a565b926020810151612f6281612f1d565b92604082015167ffffffffffffffff811161019557610377920161293a565b60409067ffffffffffffffff61037795931681528160208201520191612189565b9080601f830112156101955781602061037793359101612294565b9080601f8301121561019557813591612fd583611edc565b92612fe36040519485610222565b80845260208085019160051b830101918383116101955760208101915b83831061300f57505050505090565b823567ffffffffffffffff81116101955782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08388030112610195576040519061305c82610206565b602083013561306a81610688565b825260408301359167ffffffffffffffff83116101955761309388602080969581960101612fa2565b83820152815201920191613000565b61ffff81160361019557565b3590610273826130a2565b359061027382612f1d565b6020818303126101955780359067ffffffffffffffff8211610195570160e081830312610195576130f3610275565b91813567ffffffffffffffff81116101955781613111918401612fbd565b835261311f602083016130ae565b6020840152613130604083016130b9565b60408401526131416060830161113c565b6060840152608082013567ffffffffffffffff81116101955781613166918401612fa2565b608084015260a082013567ffffffffffffffff8111610195578161318b918401612fa2565b60a084015260c082013567ffffffffffffffff8111610195576131ae9201612fa2565b60c082015290565b929190926131c2612e03565b6004841015806134ff575b156133d1575050816131ea926131e292612206565b8101906130c4565b9081515160005b81811061332257505081515115613283575b6060820173ffffffffffffffffffffffffffffffffffffffff61323a825173ffffffffffffffffffffffffffffffffffffffff1690565b161561324557505090565b613269606061037793015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b60a0810192613293845151612ea2565b835260005b845180518210156133195790613312816132d16132b782600196612007565b5173ffffffffffffffffffffffffffffffffffffffff1690565b6132f86132dc610284565b73ffffffffffffffffffffffffffffffffffffffff9092168252565b6133006102eb565b60208201528751906117d58383612007565b5001613298565b50509250613203565b61332b8161204a565b82811061333b57506001016131f1565b61334961174c838751612007565b73ffffffffffffffffffffffffffffffffffffffff61336f6104ce61174c858a51612007565b91161461337e5760010161332b565b61068461338f61174c848851612007565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b939094919260a08201956133e6875151612ea2565b865260005b87518051821015613436579061342f8161340a6132b782600196612007565b6134156132dc610284565b61341d6102eb565b60208201528a51906117d58383612007565b50016133eb565b505061349e939650600092946134676104ce6104ce60025473ffffffffffffffffffffffffffffffffffffffff1690565b91604051958694859384937f9cc1999600000000000000000000000000000000000000000000000000000000855260048501612f81565b03915afa80156108da576000906000906000906134d3575b608086015263ffffffff16604085015290505b60a0830152613203565b5050506134f56134c9913d806000833e6134ed8183610222565b810190612f2b565b91925082916134b6565b507f302326cb000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000061355561354f87876121f8565b90612e3c565b16146131cd565b90816020910312610195575161037781611147565b6020818303126101955780519067ffffffffffffffff821161019557019080601f830112156101955781516135a581611edc565b926135b36040519485610222565b81845260208085019260051b82010192831161019557602001905b8282106135db5750505090565b6020809183516135ea81610688565b8152019101906135ce565b95949060009460a09467ffffffffffffffff6136499573ffffffffffffffffffffffffffffffffffffffff61ffff95168b521660208a0152604089015216606087015260c0608087015260c0860190610323565b930152565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146111375760010190565b9492949390936136aa60036136a48367ffffffffffffffff166000526004602052604060002090565b01611c9a565b9473ffffffffffffffffffffffffffffffffffffffff6136cb818316612071565b16926040517f01ffc9a70000000000000000000000000000000000000000000000000000000081526020818061372860048201907fdc0cbd3600000000000000000000000000000000000000000000000000000000602083019252565b0381885afa9081156108da576000916138b7575b50156138ad5790613782600095949392604051998a96879586957f89720a62000000000000000000000000000000000000000000000000000000008752600487016135f5565b03915afa9283156108da57600093613888575b50825115613883576137b26137ad8451845190612058565b611f10565b6000918293835b8651811015613832576137cf6132b78289612007565b73ffffffffffffffffffffffffffffffffffffffff81161561382657906138206001926138056137fe8961364e565b9888612007565b9073ffffffffffffffffffffffffffffffffffffffff169052565b016137b9565b50945060018095613820565b509193909450613843575b50815290565b60005b815181101561387b57806138756138626132b760019486612007565b61380561386e8761364e565b9688612007565b01613846565b50503861383d565b915090565b6138a69193503d806000833e61389e8183610222565b810190613571565b9138613795565b5050505050915090565b6138d9915060203d6020116138df575b6138d18183610222565b81019061355c565b3861373c565b503d6138c7565b6139066139016138f98351855190612058565b855190612058565b612ea2565b91600094855b8351811015613949578061394261392560019387612007565b51986139308161364e565b9961393b828a612007565b5287612007565b500161390c565b50939094915060005b8551811015613a24576139686132b78288612007565b600073ffffffffffffffffffffffffffffffffffffffff8216815b8681106139fa575b50501561399c575b50600101613952565b92906139f36001926139cb6139af610284565b73ffffffffffffffffffffffffffffffffffffffff9097168752565b6139d36102eb565b60208701526139e18161364e565b956139ec8289612007565b5286612007565b5090613993565b81613a0b6104ce61174c848c612007565b14613a1857600101613983565b5050506001388061398b565b5093506000905b8451821015613aff57613a416132b78387612007565b91600073ffffffffffffffffffffffffffffffffffffffff8416815b848110613ad5575b505015613a77575b6001019150613a2b565b613acc600192613aa4613a88610284565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b613aac6102eb565b6020860152613aba8161364e565b94613ac58288612007565b5285612007565b50829150613a6d565b81613ae66104ce61174c848b612007565b14613af357600101613a5d565b50505060013880613a65565b8252509150565b90613b1082611edc565b613b1d6040519182610222565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613b4b8294611edc565b0160005b818110613b5b57505050565b6040519060a082019180831067ffffffffffffffff8411176101e5576020926040526000815260008382015260006040820152600060608201526060608082015282828601015201613b4f565b90816060910312610195578051613bbe816130a2565b9160406020830151613bcf81612f1d565b92015161037781612f1d565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561019557016020813591019167ffffffffffffffff821161019557813603831361019557565b9160209082815201919060005b818110613c455750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff8735613c6e81610688565b16815260208781013590820152019401929101613c38565b949391929067ffffffffffffffff16855260806020860152613cfd613cc0613cae8580613bdb565b60a060808a0152610120890191612189565b613ccd6020860186613bdb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808984030160a08a0152612189565b60408401357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe185360301811215610195578401916020833593019167ffffffffffffffff8411610195578360061b360383136101955761027395613dce613d9883606097613e0d978d60c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8082613dff9a0301910152613c2b565b91613dc4613da788830161113c565b73ffffffffffffffffffffffffffffffffffffffff1660e08d0152565b6080810190613bdb565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808b8403016101008c0152612189565b908782036040890152610323565b94019061ffff169052565b9263ffffffff9061ffff60ff94999896939967ffffffffffffffff60c088019b16875216602086015216604084015216606082015260c06080820152825180955260e0810194602060e08260051b84010194019060005b818110613e8c5750505061037793945060a0818403910152610323565b909194602080613eed837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff208860019603018c526040838b5173ffffffffffffffffffffffffffffffffffffffff815116845201519181858201520190610323565b97019801910196919096613e6f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161113757565b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe820191821161113757565b9190820391821161113757565b929190815151613f93613f8e613f896040850193613f818587611f5f565b919050612058565b61204a565b613b06565b9060005b8451805182101561412f5781613fac91612007565b5190613fd26104ce6104ce845173ffffffffffffffffffffffffffffffffffffffff1690565b602083019060608251613fea60208b015161ffff1690565b92898d614026604051968795869485947f80485e2500000000000000000000000000000000000000000000000000000000865260048601613c86565b03915afa9384156108da57600194600091829383926140df575b5061ffff92916140ab63ffffffff61406f6140b8945173ffffffffffffffffffffffffffffffffffffffff1690565b97519661409961407d610293565b73ffffffffffffffffffffffffffffffffffffffff909a168a52565b1667ffffffffffffffff166020880152565b63ffffffff166040860152565b16606083015260808201526140cd8286612007565b526140d88185612007565b5001613f97565b63ffffffff945061ffff935061406f9250906140ab6141176140b89360603d8111614128575b61410f8183610222565b810190613ba8565b979195909795965050505090614040565b503d614105565b505090939492919260608301906141606104ce6104ce845173ffffffffffffffffffffffffffffffffffffffff1690565b606063ffffffff614176602088015161ffff1690565b9260ff61418660208b018b611c49565b9690506141dd6141968a8d611f5f565b9290508b5160808d01998a51926040519a8b998a9889987f84f369ce000000000000000000000000000000000000000000000000000000008a521693169160048801613e18565b03915afa9283156108da57614249936000916000936000926142df575b5061ffff92916140ab63ffffffff61406f614229945173ffffffffffffffffffffffffffffffffffffffff1690565b166060830152608082015261423e8651613efc565b90613ac58288612007565b506142548184611f5f565b9050614261575b50505090565b61427661061b61061560c0936142d696611f5f565b9101516142a0614284610293565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60006020830152600060408301526000606083015260808201526142c48351613f29565b906142cf8285612007565b5282612007565b5038808061425b565b63ffffffff945061ffff935061406f9250906140ab61430f6142299360603d6060116141285761410f8183610222565b9791959097959650505050906141fa565b9073ffffffffffffffffffffffffffffffffffffffff6143f29392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452614382606485610222565b166000806040938451956143968688610222565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15614417573d6143e36143da826102b1565b94519485610222565b83523d6000602085013e6154f8565b8051806143fd575050565b8160208061441293610273950101910161355c565b6151a1565b606092506154f8565b73ffffffffffffffffffffffffffffffffffffffff60015416330361444157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805191614479815184612058565b9283156145d05760005b848110614491575050505050565b818110156145b5576144a66132b78286612007565b73ffffffffffffffffffffffffffffffffffffffff8116801561458b576144cc8361204a565b8781106144de57505050600101614483565b8481101561455b5773ffffffffffffffffffffffffffffffffffffffff6145086132b7838a612007565b168214614517576001016144cc565b7fa1726e400000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff831660045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff6145866132b76145808885613f56565b89612007565b614508565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b6145cb6132b76145c58484613f56565b85612007565b6144a6565b7f7b6c02970000000000000000000000000000000000000000000000000000000060005260046000fd5b9160606102739294936146478160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b6020818303126101955780519067ffffffffffffffff8211610195570160408183031261019557604051916146b083610206565b815167ffffffffffffffff811161019557816146cd91840161293a565b8352602082015167ffffffffffffffff8111610195576146ed920161293a565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080614727855184602087015260c0860190610323565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b9060ff614774602092959495604085526040850190610323565b9416910152565b9190916147866127d4565b5060208101805115614b76576147b96104ce6106f16104ce855173ffffffffffffffffffffffffffffffffffffffff1690565b9473ffffffffffffffffffffffffffffffffffffffff8616158015614aeb575b614a8857906000614899926148668351614807875173ffffffffffffffffffffffffffffffffffffffff1690565b90614844614813610293565b8a815267ffffffffffffffff8c1660208201529473ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b604051809481927f9a4575b9000000000000000000000000000000000000000000000000000000008352600483016146f5565b0381838a5af180156108da57614902956149f49461496694600093614a55575b5061492e614992916132b760009596519b6040519b8c91602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018c528b610222565b604051958691602083017fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060149260601b1681520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610222565b6149bf6149b584519267ffffffffffffffff166000526004602052604060002090565b5460e01c60ff1690565b9060405195869283927f6d7fa1ce0000000000000000000000000000000000000000000000000000000084526004840161475a565b0381305afa9283156108da57600093614a35575b506020015193614a166102a2565b958652602086015260408501526060840152608083015260a082015290565b6020919350614a4e903d806000833e6117f58183610222565b9290614a08565b60009350614992916132b7614a7e61492e933d8089833e614a768183610222565b81019061467c565b95505091506148b9565b610684614aa9845173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf0000000000000000000000000000000000000000000000000000000060048201526020816024818a5afa9081156108da57600091614b57575b50156147d9565b614b70915060203d6020116138df576138d18183610222565b38614b50565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b90614bb360209282815194859201610300565b0190565b6020614cd196614c71601a614c9e947fff0000000000000000000000000000000000000000000000000000000000000060069f9c979a7fffffffffffffffff000000000000000000000000000000000000000000000000614bb39f9b8160019c614ca59f82907f0100000000000000000000000000000000000000000000000000000000000000895260c01b168e88015260c01b16600986015260c01b16601184015260f81b166019820152019182815194859201610300565b0180927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b0190614ba0565b80947fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b60e01b7fffffffff00000000000000000000000000000000000000000000000000000000166002830152565b614c9e96614de7967fffff000000000000000000000000000000000000000000000000000000000000600160029c986103779f9e9c97987fff0000000000000000000000000000000000000000000000000000000000000060049a849882614c9e9b60f81b168152614d788251809360208985019101610300565b019160f81b1683820152614d9782516002936020829585019101610300565b01019160f01b1683820152614db6825180936020600385019101610300565b010191888301907fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b80927fffff0000000000000000000000000000000000000000000000000000000000009060f01b169052565b6102739092919260206040519482614e348794518092858088019101610300565b8301614e4882518093858085019101610300565b0101037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283610222565b606081019060ff8251511161517257608081019060ff825151116151435760e081019160ff835151116151145761010082019060ff825151116150e55761012083019361ffff855151116150b6576101408401936001855151116150875761016081019261ffff845151116150585760609551805161503c575b50815167ffffffffffffffff16906020830151614f149067ffffffffffffffff1690565b926040810151614f2b9067ffffffffffffffff1690565b9951805160ff169251908151614f419060ff1690565b9060a0840151614f529061ffff1690565b60c09094015163ffffffff16946040519d8e9860208a0198614f7399614bb7565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018752614fa39087610222565b51928351614fb19060ff1690565b9251908151614fc09060ff1690565b9551928351614fd09061ffff1690565b938251614fde9061ffff1690565b9151948551614fee9061ffff1690565b94604051998a9960208b01996150039a614cfd565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810182526150339082610222565b61037791614e13565b61505191965061504b90611ffa565b51615334565b9438614ef0565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602360045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602060045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601f60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601e60045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052601d60045260246000fd5b156151a857565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b966002614de797614c9e60226103779f9e9c9799600199859f9b7fff0000000000000000000000000000000000000000000000000000000000000090614c9e9f82614c9e9c6153089c7f01000000000000000000000000000000000000000000000000000000000000008452600184015260f81b1660218201526152b98251809360208985019101610300565b019160f81b16838201526152d7825180936020602385019101610300565b010191888301907fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b80927fff000000000000000000000000000000000000000000000000000000000000009060f81b169052565b602081019060ff825151116154c757604081019160ff8351511161549857606082019160ff8351511161546957608081019260ff8451511161543a5760a0820161ffff8151511161540b57610377946153df9351945191615396835160ff1690565b9751916153a4835160ff1690565b9451906153b2825160ff1690565b9051936153c0855160ff1690565b9351966153cf885161ffff1690565b966040519c8d9b60208d0161522c565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610222565b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602860045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602760045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602660045260246000fd5b7fb4205b4200000000000000000000000000000000000000000000000000000000600052602560045260246000fd5b7fb4205b42000000000000000000000000000000000000000000000000000000006000526106846024906024600452565b91929015615573575081511561550c575090565b3b156155155790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156155865750805190602001fd5b612323906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526004830161036656fea164736f6c634300081a000a",
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
