// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ccv_proxy

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
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

type CCVProxyDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
	DefaultCCV        common.Address
	RequiredCCV       common.Address
	DefaultExecutor   common.Address
}

type CCVProxyDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
}

type CCVProxyStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	TokenAdminRegistry common.Address
}

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

type InternalEVM2AnyVerifierMessage struct {
	Header           InternalHeader
	Sender           common.Address
	Data             []byte
	Receiver         []byte
	FeeToken         common.Address
	FeeTokenAmount   *big.Int
	FeeValueJuels    *big.Int
	TokenTransfer    InternalEVMTokenTransfer
	VerifierReceipts []InternalReceipt
	ExecutorReceipt  InternalReceipt
	TokenReceipt     InternalReceipt
}

type InternalEVMTokenTransfer struct {
	SourceTokenAddress common.Address
	DestTokenAddress   []byte
	ExtraData          []byte
	Amount             *big.Int
}

type InternalHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

type InternalReceipt struct {
	Issuer            common.Address
	DestGasLimit      uint64
	DestBytesOverhead uint32
	FeeTokenAmount    *big.Int
	ExtraArgs         []byte
}

var CCVProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCCVProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"defaultCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVMTokenTransfer\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"verifierReceipts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"executorReceipt\",\"type\":\"tuple\",\"internalType\":\"structInternal.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"tokenReceipt\",\"type\":\"tuple\",\"internalType\":\"structInternal.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"receiptBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60e0604052346102f557604051613a6a38819003601f8101601f191683016001600160401b038111848210176102fa578392829160405283398101039060c082126102f557606082126102f557610054610310565b81519092906001600160401b03811681036102f55783526020820151906001600160a01b03821682036102f5576020840191825260606100966040850161032f565b6040860190815291605f1901126102f5576100af610310565b916100bc6060850161032f565b835260808401519384151585036102f55760a06100e091602086019687520161032f565b946040840195865233156102e457600180546001600160a01b0319163317905580516001600160401b03161580156102d2575b80156102c0575b61029357516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102ae575b80156102a4575b6102935780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610310565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610310565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a160405161372690816103448239608051818181610ed2015261166f015260a0518181816104bc01526116a8015260c0518181816116e401526119ef0152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b038111838210176102fa57604052565b51906001600160a01b03821682036102f55756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd57806348a98aa4146100f85780635147b35a146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da5780639041be3d146100d557806390423fa2146100d0578063df0aa9e9146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b6115ce565b6114da565b610cc1565b610aed565b610a50565b6109fe565b610915565b610849565b6107c9565b61074e565b6106d1565b610665565b610410565b61036a565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557606061014061164f565b610183604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176101d557604052565b61018a565b60a0810190811067ffffffffffffffff8211176101d557604052565b6040810190811067ffffffffffffffff8211176101d557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101d557604052565b60405190610262608083610212565b565b6040519061026261016083610212565b6040519061026260a083610212565b6040519061026260e083610212565b60405190610262604083610212565b67ffffffffffffffff81116101d557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906102ea602083610212565b60008252565b60005b8381106103035750506000910152565b81810151838201526020016102f3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361034f815180928187528780880191016102f0565b0116010190565b906020610367928181520190610313565b90565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576103e760408051906103ab8183610212565b601282527f43435650726f787920312e372e302d6465760000000000000000000000000000602083015251918291602083526020830190610313565b0390f35b6004359067ffffffffffffffff8216820361018557565b908160a09103126101855790565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576104476103eb565b60243567ffffffffffffffff811161018557610467903690600401610402565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa9081156105d857600091610618575b506105dd576105829160209161054c61053361053360025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906040518095819482937fd8694ccd00000000000000000000000000000000000000000000000000000000845260048401611826565b03915afa80156105d8576103e7916000916105a9575b506040519081529081906020820190565b6105cb915060203d6020116105d1575b6105c38183610212565b81019061172d565b38610598565b503d6105b9565b611721565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b61063a915060203d602011610640575b6106328183610212565b81019061170c565b38610502565b503d610628565b73ffffffffffffffffffffffffffffffffffffffff81160361018557565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855761069c6103eb565b5060206106b36024356106ae81610647565b611990565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760043567ffffffffffffffff8111610185573660238201121561018557806004013567ffffffffffffffff81116101855736602460a083028401011161018557602461074c9201611a89565b005b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760043567ffffffffffffffff8111610185573660238201121561018557806004013567ffffffffffffffff8111610185573660248260051b8401011161018557602461074c9201611dff565b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff6108096103eb565b1660005260046020526040806000205473ffffffffffffffffffffffffffffffffffffffff82519167ffffffffffffffff8160a01c168352166020820152f35b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557610880611630565b5060405161088d816101b9565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015260405180916103e782606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760005473ffffffffffffffffffffffffffffffffffffffff811633036109d4577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff610a906103eb565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111610ad35760209067ffffffffffffffff60405191168152f35b611f2c565b359061026282610647565b8015150361018557565b346101855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576000604051610b2a816101b9565b600435610b3681610647565b8152602435610b4481610ae3565b6020820190815260443590610b5882610647565b60408301918252610b67612610565b73ffffffffffffffffffffffffffffffffffffffff83511615918215610ca1575b508115610c96575b50610c6e5780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390610c5961164f565b610c686040519283928361275b565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610b90565b5173ffffffffffffffffffffffffffffffffffffffff1615915038610b88565b346101855760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557610cf86103eb565b60243567ffffffffffffffff811161018557610d18903690600401610402565b60443591610d27606435610647565b60025460a01c60ff166114b057610d78740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b610d9e610d998267ffffffffffffffff166000526004602052604060002090565b611f5b565b9173ffffffffffffffffffffffffffffffffffffffff606435161561148657610de1610533610533855173ffffffffffffffffffffffffffffffffffffffff1690565b330361145c57610e6493610e02610dfb6080840184611fc5565b9086612b94565b926000610e2a61053361053360025473ffffffffffffffffffffffffffffffffffffffff1690565b60a08601519060405180809a81947f9b1115e400000000000000000000000000000000000000000000000000000000835260048301610356565b03915afa9586156105d85760009661143f575b50855115611416575b610e886120a7565b958451602086019788519160408801928351610ea49060ff1690565b91610eae93612e6d565b60ff1690925290885280865251958695610ec66120c3565b90610ecf6120f1565b917f000000000000000000000000000000000000000000000000000000000000000093602001968751610f099067ffffffffffffffff1690565b610f1290612127565b67ffffffffffffffff8116909852610f28610253565b6000815267ffffffffffffffff861660208201529767ffffffffffffffff881660408a015267ffffffffffffffff1660608901528b89610f6b6020820182611fc5565b9d90610f778380611fc5565b929093606001610f8690611df5565b945151610f9291612146565b610f9b90612153565b9e610fa4610264565b9c8d5273ffffffffffffffffffffffffffffffffffffffff6064351660208e01523690610fd0926121c0565b60408c01523690610fe0926121c0565b60608a015273ffffffffffffffffffffffffffffffffffffffff16608089015260a088015260c087016000905260e087019283526101008701998a528061012088015261014087015260005b8881106113a757505060005b895180518210156110dc57906110d561105382600194612204565b516020611074825173ffffffffffffffffffffffffffffffffffffffff1690565b91015161109e611082610274565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60006020830152600060408301526000606083015260808201528b516110c48c85612146565b916110cf8383612204565b52612204565b5001611038565b505087949560408101906110f08282612218565b9050611331575b5050604080517f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f73216020820190815267ffffffffffffffff94851692820192909252928516606084015230608084015261118c93509190506111838160a081015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610212565b519020846134a0565b835152604051926111d0846111a483602083016124eb565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610212565b6111db8351516124fc565b9160005b84518051821015611293579060008161121f6105336105336112048461125698612204565b515173ffffffffffffffffffffffffffffffffffffffff1690565b89836040518097819582947f1234eab000000000000000000000000000000000000000000000000000000000845260048401612563565b03925af19182156105d857600192611270575b50016111df565b61128c903d806000833e6112848183610212565b810190612058565b5087611269565b50506103e7927f6442264de2e65ad3f5adb302a516a0ed3ae067e3e7d189a1ea66a9b09fe6b7ad67ffffffffffffffff6112da606086510167ffffffffffffffff90511690565b6112ef8260405193849316961694878361257f565b0390a361131f7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b51516040519081529081906020820190565b600161133d8383612218565b90500361137d5761118c94866113596113739461136e94612218565b939061136960643595369261226c565b612275565b6131cf565b90528580806110f7565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b8061140f816113b96001948651612204565b5160206113da825173ffffffffffffffffffffffffffffffffffffffff1690565b9101516113e8611082610274565b60006020830152600060408301526000606083015260808201528d51906110cf8383612204565b500161102c565b94506114396114258380611fc5565b91906111576040519384926020840161207e565b94610e80565b6114559196503d806000833e6112848183610212565b9438610e77565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855773ffffffffffffffffffffffffffffffffffffffff60043561152a81610647565b611532612610565b163381146115a457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576116056103eb565b507f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061163d826101b9565b60006040838281528260208201520152565b611657611630565b50604051611664816101b9565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b90816020910312610185575161036781610ae3565b6040513d6000823e3d90fd5b90816020910312610185575190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561018557016020813591019167ffffffffffffffff821161018557813603831361018557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9160209082815201919060005b8181106117e55750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff873561180e81610647565b168152602087810135908201520194019291016117d8565b919067ffffffffffffffff1682526040602083015261189961185c61184b838061173c565b60a0604087015260e086019161178c565b611869602084018461173c565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc086840301606087015261178c565b9160408201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181121561018557820160208135910167ffffffffffffffff8211610185578160061b36038113610185578461196092611929927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08661036799030160808701526117cb565b9261195661193960608301610ad8565b73ffffffffffffffffffffffffffffffffffffffff1660a0850152565b608081019061173c565b9160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08286030191015261178c565b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156105d857600090611a3a575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011611a81575b81611a5460209383610212565b810103126101855773ffffffffffffffffffffffffffffffffffffffff9051611a7c81610647565b611a1f565b3d9150611a47565b90611a92612610565b60005b818110611aa157505050565b611ab4611aaf828486611d54565b611d69565b805167ffffffffffffffff1690818015611cee576001939291611cb6611b0f7ff2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed09367ffffffffffffffff166000526004602052604060002090565b611ca5611c616080611b38602087015173ffffffffffffffffffffffffffffffffffffffff1690565b84547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff821617855595611bdf611b9b604083015173ffffffffffffffffffffffffffffffffffffffff1690565b600287019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b611c46611c03606083015173ffffffffffffffffffffffffffffffffffffffff1690565b8c87019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b015173ffffffffffffffffffffffffffffffffffffffff1690565b600383019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5460a01c67ffffffffffffffff1690565b6040805167ffffffffffffffff92909216825273ffffffffffffffffffffffffffffffffffffffff929092166020820152a201611a95565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611d645760a0020190565b611d25565b60a0813603126101855760405190611d80826101da565b80359067ffffffffffffffff821682036101855760809183526020810135611da781610647565b60208401526040810135611dba81610647565b60408401526060810135611dcd81610647565b60608401520135611ddd81610647565b608082015290565b9190811015611d645760051b0190565b3561036781610647565b60035473ffffffffffffffffffffffffffffffffffffffff169160005b818110611e295750505050565b611e3f610533611e3a838587611de5565b611df5565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa9081156105d8576001948891600093611f0c575b5082611eb4575b5050505001611e1c565b611ebf91839161265b565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681611eaa565b611f2591935060203d81116105d1576105c38183610212565b9138611ea3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90610262604051611f6b816101da565b835473ffffffffffffffffffffffffffffffffffffffff818116835260a09190911c67ffffffffffffffff166020830152600185015481166040830152600285015481166060830152600394909401549093166080840152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff82116101855760200191813603831361018557565b81601f8201121561018557805161202c816102a1565b9261203a6040519485610212565b818452602082840101116101855761036791602080850191016102f0565b9060208282031261018557815167ffffffffffffffff8111610185576103679201612016565b91602061036793818152019161178c565b67ffffffffffffffff81116101d55760051b60200190565b604051906120b6602083610212565b6000808352366020840137565b604051906120d0826101da565b60606080836000815260006020820152600060408201526000838201520152565b604051906080820182811067ffffffffffffffff8211176101d55760405260006060838281528160208201528160408201520152565b67ffffffffffffffff1667ffffffffffffffff8114610ad35760010190565b91908201809211610ad357565b9061215d8261208f565b61216a6040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612198829461208f565b019060005b8281106121a957505050565b6020906121b46120c3565b8282850101520161219d565b9291926121cc826102a1565b916121da6040519384610212565b829481845281830111610185578281602093846000960137010152565b805115611d645760200190565b8051821015611d645760209160051b010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160061b3603831361018557565b9015611d645790565b91908260409103126101855760405161228d816101f6565b6020808294803561229d81610647565b84520135910152565b9073ffffffffffffffffffffffffffffffffffffffff82511681526060806122f26122e06020860151608060208701526080860190610313565b60408601518582036040870152610313565b93015191015290565b9060a060806103679373ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff602082015116602085015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610313565b9080602083519182815201916020808360051b8301019401926000915b83831061238857505050505090565b90919293946020806123c4837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516122fb565b97019301930191939290612379565b6103679161240e81835167ffffffffffffffff6060809280518552826020820151166020860152826040820151166040860152015116910152565b602082015173ffffffffffffffffffffffffffffffffffffffff1660808201526101406124d96124c56124b161246a61245860408801516101c060a08901526101c0880190610313565b606088015187820360c0890152610313565b608087015173ffffffffffffffffffffffffffffffffffffffff1660e087015260a087015161010087015260c087015161012087015260e0870151868203868801526122a6565b61010086015185820361016087015261235c565b6101208501518482036101808601526122fb565b920151906101a08184039101526122fb565b9060206103679281815201906123d3565b906125068261208f565b6125136040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612541829461208f565b019060005b82811061255257505050565b806060602080938501015201612546565b92919061257a602091604086526040860190610313565b930152565b90612592906040835260408301906123d3565b906020818303910152815180825260208201916020808360051b8301019401926000915b8383106125c557505050505090565b9091929394602080612601837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610313565b970193019301919392906125b6565b73ffffffffffffffffffffffffffffffffffffffff60015416330361263157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9073ffffffffffffffffffffffffffffffffffffffff61272d9392604051938260208601947fa9059cbb0000000000000000000000000000000000000000000000000000000086521660248601526044850152604484526126bd606485610212565b166000806040938451956126d18688610212565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15612752573d61271e612715826102a1565b94519485610212565b83523d6000602085013e613651565b805180612738575050565b8160208061274d93610262950101910161170c565b6135c6565b60609250613651565b9160606102629294936127a88160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b6040519060e0820182811067ffffffffffffffff8211176101d557604052606060c08382815282602082015260006040820152600083820152600060808201528260a08201520152565b906004116101855790600490565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612869575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b604080519091906128ac8382610212565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b8281106128e357505050565b6020906040516128f2816101f6565b60008152606083820152828285010152016128d7565b906129128261208f565b61291f6040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061294d829461208f565b019060005b82811061295e57505050565b60209060405161296d816101f6565b6000815260608382015282828501015201612952565b9080601f8301121561018557816020610367933591016121c0565b9080601f83011215610185578135916129b68361208f565b926129c46040519485610212565b80845260208085019160051b830101918383116101855760208101915b8383106129f057505050505090565b823567ffffffffffffffff81116101855782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083880301126101855760405190612a3d826101f6565b6020830135612a4b81610647565b825260408301359167ffffffffffffffff831161018557612a7488602080969581960101612983565b838201528152019201916129e1565b359060ff8216820361018557565b359063ffffffff8216820361018557565b6020818303126101855780359067ffffffffffffffff8211610185570160e08183031261018557612ad1610283565b91813567ffffffffffffffff81116101855781612aef91840161299e565b8352602082013567ffffffffffffffff81116101855781612b1191840161299e565b6020840152612b2260408301612a83565b6040840152612b3360608301612a91565b6060840152612b4460808301610ad8565b608084015260a082013567ffffffffffffffff81116101855781612b69918401612983565b60a084015260c082013567ffffffffffffffff811161018557612b8c9201612983565b60c082015290565b90612b9d6127dd565b927f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000612bf3612bed8486612827565b90612835565b1603612d7a57612c07929350810190612aa2565b906020820180515180612d25575b50825151905151612c2591612146565b15612cab575b6080820173ffffffffffffffffffffffffffffffffffffffff612c62825173ffffffffffffffffffffffffffffffffffffffff1690565b1615612c6d57505090565b612c91608061036793015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b612cb361289b565b8252612d1f612cd9606083015173ffffffffffffffffffffffffffffffffffffffff1690565b612d00612ce4610292565b73ffffffffffffffffffffffffffffffffffffffff9092168252565b612d086102db565b6020820152835190612d19826121f7565b526121f7565b50612c2b565b604084015160ff1690811080159190612d6e575b50612d445738612c15565b7fb273a0e40000000000000000000000000000000000000000000000000000000060005260046000fd5b60ff9150161538612d39565b612d08612d1f92612d8c3684836121c0565b60a0870152612d9961289b565b8652606085015173ffffffffffffffffffffffffffffffffffffffff1692612dde612dc2610292565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b36916121c0565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610ad35760010190565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610ad357565b60ff168015610ad3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b919092612e7a8351612908565b91600080935b855185101561304157612eb2612e9886889a98612204565b5173ffffffffffffffffffffffffffffffffffffffff1690565b60009773ffffffffffffffffffffffffffffffffffffffff821697895b825181101561303157612ee86105336112048386612204565b8a14612ef657600101612ecf565b50969297919098506001999394995b15612f1a575b50506001019396919094612e80565b612f5790612f3b612f34879c9899949b959a969c97612de5565b9686612204565b519073ffffffffffffffffffffffffffffffffffffffff169052565b60005b855181101561302157612f736105336112048389612204565b8914612f8157600101612f5a565b859993979298506020612f9982612fd2949998612204565b5101516020612fb0612faa88612e12565b8b612204565b510152612fc0612faa8b51612e12565b51612fcb828c612204565b5289612204565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff885101885260ff841661300d575b6001905b9038612f0b565b92613019600191612e3f565b939050613002565b5094939791956001919750613006565b5096929791909899939499612f05565b91969490929593508261305657505050929190565b61306b83613070929894969793959851612146565b612908565b9260005b8581106130c757505060005b86518110156130bd57806130b66130996001938a612204565b516130a48389612146565b906130af8289612204565b5286612204565b5001613080565b5091945092909150565b806130d460019284612204565b516130df8288612204565b526130ea8187612204565b5001613074565b6020818303126101855780519067ffffffffffffffff821161018557016040818303126101855760405191613125836101f6565b815167ffffffffffffffff81116101855781613142918401612016565b8352602082015167ffffffffffffffff8111610185576131629201612016565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff608061319c855184602087015260c0860190610313565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b919290926131db6120f1565b5060208301938451156134655761320f6105336106ae610533875173ffffffffffffffffffffffffffffffffffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff84161580156133da575b613377579260009493926132c56132fa9387956132a38a51916132866132698b5173ffffffffffffffffffffffffffffffffffffffff1690565b94613272610274565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b6040519586809481937f9a4575b90000000000000000000000000000000000000000000000000000000083526004830161316a565b03925af19182156105d85760009261334e575b505173ffffffffffffffffffffffffffffffffffffffff165b916020825192015190519161333c612dc2610253565b60208401526040830152606082015290565b613326919250613370903d806000833e6133688183610212565b8101906130f1565b919061330d565b610614613398865173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481885afa9081156105d857600091613446575b501561322f565b61345f915060203d602011610640576106328183610212565b3861343f565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9060206103679281815201906122a6565b6135c06134c4602083015173ffffffffffffffffffffffffffffffffffffffff1690565b82516060015167ffffffffffffffff1661354c6134f8608086015173ffffffffffffffffffffffffffffffffffffffff1690565b9161115760a08701516040519485936020850197889094939273ffffffffffffffffffffffffffffffffffffffff9067ffffffffffffffff6060948360808601991685521660208401521660408201520152565b5190206111576060840151602081519101209360e06040820151602081519101209101516040516135858161115760208201948561348f565b51902090604051958694602086019889919260a093969594919660c08401976000855260208501526040840152606083015260808201520152565b51902090565b156135cd57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b919290156136cc5750815115613665575090565b3b1561366e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156136df5750805190602001fd5b613715906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610356565b0390fdfea164736f6c634300081a000a",
}

var CCVProxyABI = CCVProxyMetaData.ABI

var CCVProxyBin = CCVProxyMetaData.Bin

func DeployCCVProxy(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig CCVProxyStaticConfig, dynamicConfig CCVProxyDynamicConfig) (common.Address, *types.Transaction, *CCVProxy, error) {
	parsed, err := CCVProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCVProxyBin), backend, staticConfig, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCVProxy{address: address, abi: *parsed, CCVProxyCaller: CCVProxyCaller{contract: contract}, CCVProxyTransactor: CCVProxyTransactor{contract: contract}, CCVProxyFilterer: CCVProxyFilterer{contract: contract}}, nil
}

type CCVProxy struct {
	address common.Address
	abi     abi.ABI
	CCVProxyCaller
	CCVProxyTransactor
	CCVProxyFilterer
}

type CCVProxyCaller struct {
	contract *bind.BoundContract
}

type CCVProxyTransactor struct {
	contract *bind.BoundContract
}

type CCVProxyFilterer struct {
	contract *bind.BoundContract
}

type CCVProxySession struct {
	Contract     *CCVProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCVProxyCallerSession struct {
	Contract *CCVProxyCaller
	CallOpts bind.CallOpts
}

type CCVProxyTransactorSession struct {
	Contract     *CCVProxyTransactor
	TransactOpts bind.TransactOpts
}

type CCVProxyRaw struct {
	Contract *CCVProxy
}

type CCVProxyCallerRaw struct {
	Contract *CCVProxyCaller
}

type CCVProxyTransactorRaw struct {
	Contract *CCVProxyTransactor
}

func NewCCVProxy(address common.Address, backend bind.ContractBackend) (*CCVProxy, error) {
	abi, err := abi.JSON(strings.NewReader(CCVProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCVProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCVProxy{address: address, abi: abi, CCVProxyCaller: CCVProxyCaller{contract: contract}, CCVProxyTransactor: CCVProxyTransactor{contract: contract}, CCVProxyFilterer: CCVProxyFilterer{contract: contract}}, nil
}

func NewCCVProxyCaller(address common.Address, caller bind.ContractCaller) (*CCVProxyCaller, error) {
	contract, err := bindCCVProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCVProxyCaller{contract: contract}, nil
}

func NewCCVProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CCVProxyTransactor, error) {
	contract, err := bindCCVProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCVProxyTransactor{contract: contract}, nil
}

func NewCCVProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CCVProxyFilterer, error) {
	contract, err := bindCCVProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCVProxyFilterer{contract: contract}, nil
}

func bindCCVProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCVProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCVProxy *CCVProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCVProxy.Contract.CCVProxyCaller.contract.Call(opts, result, method, params...)
}

func (_CCVProxy *CCVProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVProxy.Contract.CCVProxyTransactor.contract.Transfer(opts)
}

func (_CCVProxy *CCVProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCVProxy.Contract.CCVProxyTransactor.contract.Transact(opts, method, params...)
}

func (_CCVProxy *CCVProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCVProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_CCVProxy *CCVProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVProxy.Contract.contract.Transfer(opts)
}

func (_CCVProxy *CCVProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCVProxy.Contract.contract.Transact(opts, method, params...)
}

func (_CCVProxy *CCVProxyCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_CCVProxy *CCVProxySession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CCVProxy.Contract.GetDestChainConfig(&_CCVProxy.CallOpts, destChainSelector)
}

func (_CCVProxy *CCVProxyCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _CCVProxy.Contract.GetDestChainConfig(&_CCVProxy.CallOpts, destChainSelector)
}

func (_CCVProxy *CCVProxyCaller) GetDynamicConfig(opts *bind.CallOpts) (CCVProxyDynamicConfig, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(CCVProxyDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCVProxyDynamicConfig)).(*CCVProxyDynamicConfig)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetDynamicConfig() (CCVProxyDynamicConfig, error) {
	return _CCVProxy.Contract.GetDynamicConfig(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCallerSession) GetDynamicConfig() (CCVProxyDynamicConfig, error) {
	return _CCVProxy.Contract.GetDynamicConfig(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _CCVProxy.Contract.GetExpectedNextSequenceNumber(&_CCVProxy.CallOpts, destChainSelector)
}

func (_CCVProxy *CCVProxyCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _CCVProxy.Contract.GetExpectedNextSequenceNumber(&_CCVProxy.CallOpts, destChainSelector)
}

func (_CCVProxy *CCVProxyCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _CCVProxy.Contract.GetFee(&_CCVProxy.CallOpts, destChainSelector, message)
}

func (_CCVProxy *CCVProxyCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _CCVProxy.Contract.GetFee(&_CCVProxy.CallOpts, destChainSelector, message)
}

func (_CCVProxy *CCVProxyCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _CCVProxy.Contract.GetPoolBySourceToken(&_CCVProxy.CallOpts, arg0, sourceToken)
}

func (_CCVProxy *CCVProxyCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _CCVProxy.Contract.GetPoolBySourceToken(&_CCVProxy.CallOpts, arg0, sourceToken)
}

func (_CCVProxy *CCVProxyCaller) GetStaticConfig(opts *bind.CallOpts) (CCVProxyStaticConfig, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(CCVProxyStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CCVProxyStaticConfig)).(*CCVProxyStaticConfig)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetStaticConfig() (CCVProxyStaticConfig, error) {
	return _CCVProxy.Contract.GetStaticConfig(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCallerSession) GetStaticConfig() (CCVProxyStaticConfig, error) {
	return _CCVProxy.Contract.GetStaticConfig(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCVProxy *CCVProxySession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _CCVProxy.Contract.GetSupportedTokens(&_CCVProxy.CallOpts, arg0)
}

func (_CCVProxy *CCVProxyCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _CCVProxy.Contract.GetSupportedTokens(&_CCVProxy.CallOpts, arg0)
}

func (_CCVProxy *CCVProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCVProxy *CCVProxySession) Owner() (common.Address, error) {
	return _CCVProxy.Contract.Owner(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCallerSession) Owner() (common.Address, error) {
	return _CCVProxy.Contract.Owner(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCVProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCVProxy *CCVProxySession) TypeAndVersion() (string, error) {
	return _CCVProxy.Contract.TypeAndVersion(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyCallerSession) TypeAndVersion() (string, error) {
	return _CCVProxy.Contract.TypeAndVersion(&_CCVProxy.CallOpts)
}

func (_CCVProxy *CCVProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "acceptOwnership")
}

func (_CCVProxy *CCVProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _CCVProxy.Contract.AcceptOwnership(&_CCVProxy.TransactOpts)
}

func (_CCVProxy *CCVProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCVProxy.Contract.AcceptOwnership(&_CCVProxy.TransactOpts)
}

func (_CCVProxy *CCVProxyTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []CCVProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_CCVProxy *CCVProxySession) ApplyDestChainConfigUpdates(destChainConfigArgs []CCVProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _CCVProxy.Contract.ApplyDestChainConfigUpdates(&_CCVProxy.TransactOpts, destChainConfigArgs)
}

func (_CCVProxy *CCVProxyTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []CCVProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _CCVProxy.Contract.ApplyDestChainConfigUpdates(&_CCVProxy.TransactOpts, destChainConfigArgs)
}

func (_CCVProxy *CCVProxyTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_CCVProxy *CCVProxySession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.ForwardFromRouter(&_CCVProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_CCVProxy *CCVProxyTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.ForwardFromRouter(&_CCVProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_CCVProxy *CCVProxyTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCVProxyDynamicConfig) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_CCVProxy *CCVProxySession) SetDynamicConfig(dynamicConfig CCVProxyDynamicConfig) (*types.Transaction, error) {
	return _CCVProxy.Contract.SetDynamicConfig(&_CCVProxy.TransactOpts, dynamicConfig)
}

func (_CCVProxy *CCVProxyTransactorSession) SetDynamicConfig(dynamicConfig CCVProxyDynamicConfig) (*types.Transaction, error) {
	return _CCVProxy.Contract.SetDynamicConfig(&_CCVProxy.TransactOpts, dynamicConfig)
}

func (_CCVProxy *CCVProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_CCVProxy *CCVProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.TransferOwnership(&_CCVProxy.TransactOpts, to)
}

func (_CCVProxy *CCVProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.TransferOwnership(&_CCVProxy.TransactOpts, to)
}

func (_CCVProxy *CCVProxyTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CCVProxy.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CCVProxy *CCVProxySession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.WithdrawFeeTokens(&_CCVProxy.TransactOpts, feeTokens)
}

func (_CCVProxy *CCVProxyTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCVProxy.Contract.WithdrawFeeTokens(&_CCVProxy.TransactOpts, feeTokens)
}

type CCVProxyCCIPMessageSentIterator struct {
	Event *CCVProxyCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyCCIPMessageSent)
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
		it.Event = new(CCVProxyCCIPMessageSent)
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

func (it *CCVProxyCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *CCVProxyCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalEVM2AnyVerifierMessage
	ReceiptBlobs      [][]byte
	Raw               types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*CCVProxyCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyCCIPMessageSentIterator{contract: _CCVProxy.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *CCVProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyCCIPMessageSent)
				if err := _CCVProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseCCIPMessageSent(log types.Log) (*CCVProxyCCIPMessageSent, error) {
	event := new(CCVProxyCCIPMessageSent)
	if err := _CCVProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyConfigSetIterator struct {
	Event *CCVProxyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyConfigSet)
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
		it.Event = new(CCVProxyConfigSet)
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

func (it *CCVProxyConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCVProxyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyConfigSet struct {
	StaticConfig  CCVProxyStaticConfig
	DynamicConfig CCVProxyDynamicConfig
	Raw           types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterConfigSet(opts *bind.FilterOpts) (*CCVProxyConfigSetIterator, error) {

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCVProxyConfigSetIterator{contract: _CCVProxy.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CCVProxyConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyConfigSet)
				if err := _CCVProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseConfigSet(log types.Log) (*CCVProxyConfigSet, error) {
	event := new(CCVProxyConfigSet)
	if err := _CCVProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyDestChainConfigSetIterator struct {
	Event *CCVProxyDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyDestChainConfigSet)
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
		it.Event = new(CCVProxyDestChainConfigSet)
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

func (it *CCVProxyDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCVProxyDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	Raw               types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CCVProxyDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyDestChainConfigSetIterator{contract: _CCVProxy.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCVProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyDestChainConfigSet)
				if err := _CCVProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseDestChainConfigSet(log types.Log) (*CCVProxyDestChainConfigSet, error) {
	event := new(CCVProxyDestChainConfigSet)
	if err := _CCVProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyFeeTokenWithdrawnIterator struct {
	Event *CCVProxyFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyFeeTokenWithdrawn)
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
		it.Event = new(CCVProxyFeeTokenWithdrawn)
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

func (it *CCVProxyFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CCVProxyFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*CCVProxyFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyFeeTokenWithdrawnIterator{contract: _CCVProxy.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCVProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyFeeTokenWithdrawn)
				if err := _CCVProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CCVProxyFeeTokenWithdrawn, error) {
	event := new(CCVProxyFeeTokenWithdrawn)
	if err := _CCVProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyOwnershipTransferRequestedIterator struct {
	Event *CCVProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyOwnershipTransferRequested)
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
		it.Event = new(CCVProxyOwnershipTransferRequested)
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

func (it *CCVProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCVProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyOwnershipTransferRequestedIterator{contract: _CCVProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCVProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyOwnershipTransferRequested)
				if err := _CCVProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCVProxyOwnershipTransferRequested, error) {
	event := new(CCVProxyOwnershipTransferRequested)
	if err := _CCVProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCVProxyOwnershipTransferredIterator struct {
	Event *CCVProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCVProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCVProxyOwnershipTransferred)
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
		it.Event = new(CCVProxyOwnershipTransferred)
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

func (it *CCVProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCVProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCVProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCVProxy *CCVProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCVProxyOwnershipTransferredIterator{contract: _CCVProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCVProxy *CCVProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCVProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCVProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCVProxyOwnershipTransferred)
				if err := _CCVProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCVProxy *CCVProxyFilterer) ParseOwnershipTransferred(log types.Log) (*CCVProxyOwnershipTransferred, error) {
	event := new(CCVProxyOwnershipTransferred)
	if err := _CCVProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDestChainConfig struct {
	SequenceNumber uint64
	Router         common.Address
}

func (_CCVProxy *CCVProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _CCVProxy.abi.Events["CCIPMessageSent"].ID:
		return _CCVProxy.ParseCCIPMessageSent(log)
	case _CCVProxy.abi.Events["ConfigSet"].ID:
		return _CCVProxy.ParseConfigSet(log)
	case _CCVProxy.abi.Events["DestChainConfigSet"].ID:
		return _CCVProxy.ParseDestChainConfigSet(log)
	case _CCVProxy.abi.Events["FeeTokenWithdrawn"].ID:
		return _CCVProxy.ParseFeeTokenWithdrawn(log)
	case _CCVProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _CCVProxy.ParseOwnershipTransferRequested(log)
	case _CCVProxy.abi.Events["OwnershipTransferred"].ID:
		return _CCVProxy.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (CCVProxyCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x6442264de2e65ad3f5adb302a516a0ed3ae067e3e7d189a1ea66a9b09fe6b7ad")
}

func (CCVProxyConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3")
}

func (CCVProxyDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed0")
}

func (CCVProxyFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CCVProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCVProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_CCVProxy *CCVProxy) Address() common.Address {
	return _CCVProxy.address
}

type CCVProxyInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (CCVProxyDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (CCVProxyStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []CCVProxyDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig CCVProxyDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*CCVProxyCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *CCVProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*CCVProxyCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*CCVProxyConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *CCVProxyConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*CCVProxyConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*CCVProxyDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *CCVProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*CCVProxyDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*CCVProxyFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCVProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCVProxyFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCVProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCVProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCVProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCVProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCVProxyOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
