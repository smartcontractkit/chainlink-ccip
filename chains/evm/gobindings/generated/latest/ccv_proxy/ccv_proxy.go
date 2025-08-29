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
	TokenTransfer    []InternalEVMTokenTransfer
	VerifierReceipts []InternalReceipt
	ExecutorReceipt  InternalReceipt
}

type InternalEVMTokenTransfer struct {
	SourceTokenAddress common.Address
	DestTokenAddress   []byte
	Amount             *big.Int
	ExtraData          []byte
	Receipt            InternalReceipt
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structCCVProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"defaultCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.EVMTokenTransfer[]\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receipt\",\"type\":\"tuple\",\"internalType\":\"structInternal.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"verifierReceipts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"executorReceipt\",\"type\":\"tuple\",\"internalType\":\"structInternal.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"receiptBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCCVProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidOptionalCCVThreshold\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60e0604052346102fc57604051613fc138819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a1604051613c76908161034b8239608051818181610f14015281816111ea015261178d015260a0518181816104bc01526117c6015260c0518181816118020152611ab00152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd57806348a98aa4146100f85780635147b35a146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da5780639041be3d146100d557806390423fa2146100d0578063df0aa9e9146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b6116ec565b6115f8565b610cc1565b610aed565b610a50565b6109fe565b610915565b610849565b6107c9565b61074e565b6106d1565b610665565b610410565b61036a565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557606061014061176d565b610183604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176101d557604052565b61018a565b60a0810190811067ffffffffffffffff8211176101d557604052565b6040810190811067ffffffffffffffff8211176101d557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101d557604052565b60405190610262608083610212565b565b6040519061026261014083610212565b6040519061026260a083610212565b6040519061026260e083610212565b60405190610262604083610212565b67ffffffffffffffff81116101d557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906102ea602083610212565b60008252565b60005b8381106103035750506000910152565b81810151838201526020016102f3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361034f815180928187528780880191016102f0565b0116010190565b906020610367928181520190610313565b90565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576103e760408051906103ab8183610212565b601282527f43435650726f787920312e372e302d6465760000000000000000000000000000602083015251918291602083526020830190610313565b0390f35b6004359067ffffffffffffffff8216820361018557565b908160a09103126101855790565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576104476103eb565b60243567ffffffffffffffff811161018557610467903690600401610402565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa9081156105d857600091610618575b506105dd576105829160209161054c61053361053360025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906040518095819482937fd8694ccd00000000000000000000000000000000000000000000000000000000845260048401611a30565b03915afa80156105d8576103e7916000916105a9575b506040519081529081906020820190565b6105cb915060203d6020116105d1575b6105c38183610212565b81019061184b565b38610598565b503d6105b9565b61183f565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b61063a915060203d602011610640575b6106328183610212565b81019061182a565b38610502565b503d610628565b73ffffffffffffffffffffffffffffffffffffffff81160361018557565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855761069c6103eb565b5060206106b36024356106ae81610647565b611a51565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760043567ffffffffffffffff8111610185573660238201121561018557806004013567ffffffffffffffff81116101855736602460a083028401011161018557602461074c9201611b4a565b005b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760043567ffffffffffffffff8111610185573660238201121561018557806004013567ffffffffffffffff8111610185573660248260051b8401011161018557602461074c9201611ec0565b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff6108096103eb565b1660005260046020526040806000205473ffffffffffffffffffffffffffffffffffffffff82519167ffffffffffffffff8160a01c168352166020820152f35b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855761088061174e565b5060405161088d816101b9565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015260405180916103e782606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760005473ffffffffffffffffffffffffffffffffffffffff811633036109d4577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff610a906103eb565b166000526004602052600167ffffffffffffffff60406000205460a01c160167ffffffffffffffff8111610ad35760209067ffffffffffffffff60405191168152f35b611fed565b359061026282610647565b8015150361018557565b346101855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576000604051610b2a816101b9565b600435610b3681610647565b8152602435610b4481610ae3565b6020820190815260443590610b5882610647565b60408301918252610b6761298e565b73ffffffffffffffffffffffffffffffffffffffff83511615918215610ca1575b508115610c96575b50610c6e5780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390610c5961176d565b610c6860405192839283612ad9565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610b90565b5173ffffffffffffffffffffffffffffffffffffffff1615915038610b88565b346101855760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557610cf86103eb565b60243567ffffffffffffffff811161018557610d18903690600401610402565b9060443590610d28606435610647565b60025460a01c60ff166115ce57610d79740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b610d9f610d9a8267ffffffffffffffff166000526004602052604060002090565b61201c565b9073ffffffffffffffffffffffffffffffffffffffff60643516156115a457610de2610533610533845173ffffffffffffffffffffffffffffffffffffffff1690565b330361157a57610e6591610e03610dfc6080870187612086565b9083612f4d565b916000610e2b61053361053360025473ffffffffffffffffffffffffffffffffffffffff1690565b60a08501519060405180809881947f9b1115e400000000000000000000000000000000000000000000000000000000835260048301610356565b03915afa9384156105d85760009461155d575b50835115611534575b610e89612150565b610e91612196565b968451602086019889519160408801928351610ead9060ff1690565b91610eb793613304565b60ff16909252908952808652519588610ece612150565b95602001958651610ee69067ffffffffffffffff1690565b610eef906121b2565b67ffffffffffffffff8116909752610f05610253565b6000815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208201529667ffffffffffffffff8716604089015267ffffffffffffffff166060880152610f666020850185612086565b9a90610f728680612086565b610f7e60608901611eb6565b928d610f8d60408b018b6121d1565b610f97915061225a565b975151610fa3916122d5565b610fac906122e2565b9e610fb5610264565b9c8d5273ffffffffffffffffffffffffffffffffffffffff6064351660208e01523690610fe19261234f565b60408c01523690610ff19261234f565b60608a015273ffffffffffffffffffffffffffffffffffffffff16608089015260a088015260c087016000905260e087019182526101008701998a5261012087015260005b8881106114c5575060005b8a5180518210156110e557906110de61105c82600194612393565b51602061107d825173ffffffffffffffffffffffffffffffffffffffff1690565b9101516110a761108b610274565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60006020830152600060408301526000606083015260808201528c516110cd8d856122d5565b916110d88383612393565b52612393565b5001611041565b505061118d602061112b8b989961115761111c610533610533608085015173ffffffffffffffffffffffffffffffffffffffff1690565b91604051938491868301612445565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101845283610212565b60405180809581947f4a7597b50000000000000000000000000000000000000000000000000000000083528a8d6004850161254b565b03915afa80156105d8576114a8575b506111aa60408401846121d1565b905061140b575b5050604080517f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f73216020820190815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116938301939093529185166060820152306080820152611266935090915061125d8160a081015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610212565b519020846139f0565b835152604051926112aa8461127e8360208301612869565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610212565b6112b583515161287a565b9160005b8451805182101561136d57906000816112f96105336105336112de8461133098612393565b515173ffffffffffffffffffffffffffffffffffffffff1690565b89836040518097819582947f1234eab0000000000000000000000000000000000000000000000000000000008452600484016128e1565b03925af19182156105d85760019261134a575b50016112b9565b611366903d806000833e61135e8183610212565b810190612119565b5087611343565b50506103e7927fa816f7e08da08b1aa0143155f28f728327e40df7f707f612cb3566ab9122982067ffffffffffffffff6113b4606086510167ffffffffffffffff90511690565b6113c9826040519384931696169487836128fd565b0390a36113f97fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b51516040519081529081906020820190565b600161141a60408501856121d1565b90500361147e5761146b611459611454611472948861143f88604060809a01906121d1565b939061144f606435953692612580565b612589565b613657565b82519061146582612386565b52612386565b5051612386565b510152838080806111b1565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b6114c09060203d6020116105d1576105c38183610212565b61119c565b8061152d816114d76001948c51612393565b5160206114f8825173ffffffffffffffffffffffffffffffffffffffff1690565b91015161150661108b610274565b60006020830152600060408301526000606083015260808201528d51906110d88383612393565b5001611036565b92506115576115438680612086565b91906112316040519384926020840161213f565b92610e81565b6115739194503d806000833e61135e8183610212565b9238610e78565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855773ffffffffffffffffffffffffffffffffffffffff60043561164881610647565b61165061298e565b163381146116c257807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576117236103eb565b507f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b6040519061175b826101b9565b60006040838281528260208201520152565b61177561174e565b50604051611782816101b9565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b90816020910312610185575161036781610ae3565b6040513d6000823e3d90fd5b90816020910312610185575190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561018557016020813591019167ffffffffffffffff821161018557813603831361018557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9160209082815201919060005b8181106119035750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff873561192c81610647565b168152602087810135908201520194019291016118f6565b61197d611962611954838061185a565b60a0865260a08601916118aa565b61196f602084018461185a565b9085830360208701526118aa565b9160408201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181121561018557820160208135910167ffffffffffffffff8211610185578160061b360381136101855784611a22926119eb9285610367980360408701526118e9565b92611a186119fb60608301610ad8565b73ffffffffffffffffffffffffffffffffffffffff166060850152565b608081019061185a565b9160808185039101526118aa565b60409067ffffffffffffffff61036794931681528160208201520190611944565b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156105d857600090611afb575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011611b42575b81611b1560209383610212565b810103126101855773ffffffffffffffffffffffffffffffffffffffff9051611b3d81610647565b611ae0565b3d9150611b08565b90611b5361298e565b60005b818110611b6257505050565b611b75611b70828486611e15565b611e2a565b805167ffffffffffffffff1690818015611daf576001939291611d77611bd07ff2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed09367ffffffffffffffff166000526004602052604060002090565b611d66611d226080611bf9602087015173ffffffffffffffffffffffffffffffffffffffff1690565b84547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff821617855595611ca0611c5c604083015173ffffffffffffffffffffffffffffffffffffffff1690565b600287019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b611d07611cc4606083015173ffffffffffffffffffffffffffffffffffffffff1690565b8c87019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b015173ffffffffffffffffffffffffffffffffffffffff1690565b600383019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5460a01c67ffffffffffffffff1690565b6040805167ffffffffffffffff92909216825273ffffffffffffffffffffffffffffffffffffffff929092166020820152a201611b56565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611e255760a0020190565b611de6565b60a0813603126101855760405190611e41826101da565b80359067ffffffffffffffff821682036101855760809183526020810135611e6881610647565b60208401526040810135611e7b81610647565b60408401526060810135611e8e81610647565b60608401520135611e9e81610647565b608082015290565b9190811015611e255760051b0190565b3561036781610647565b60035473ffffffffffffffffffffffffffffffffffffffff169160005b818110611eea5750505050565b611f00610533611efb838587611ea6565b611eb6565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa9081156105d8576001948891600093611fcd575b5082611f75575b5050505001611edd565b611f809183916129d9565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681611f6b565b611fe691935060203d81116105d1576105c38183610212565b9138611f64565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9061026260405161202c816101da565b835473ffffffffffffffffffffffffffffffffffffffff818116835260a09190911c67ffffffffffffffff166020830152600185015481166040830152600285015481166060830152600394909401549093166080840152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff82116101855760200191813603831361018557565b81601f820112156101855780516120ed816102a1565b926120fb6040519485610212565b818452602082840101116101855761036791602080850191016102f0565b9060208282031261018557815167ffffffffffffffff81116101855761036792016120d7565b9160206103679381815201916118aa565b6040519061215d826101da565b60606080836000815260006020820152600060408201526000838201520152565b67ffffffffffffffff81116101d55760051b60200190565b604051906121a5602083610212565b6000808352366020840137565b67ffffffffffffffff1667ffffffffffffffff8114610ad35760010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160061b3603831361018557565b60405190612232826101da565b816000815260606020820152600060408201526060808201526080612255612150565b910152565b906122648261217e565b6122716040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061229f829461217e565b019060005b8281106122b057505050565b6020906122bb612225565b828285010152016122a4565b9060018201809211610ad357565b91908201809211610ad357565b906122ec8261217e565b6122f96040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612327829461217e565b019060005b82811061233857505050565b602090612343612150565b8282850101520161232c565b92919261235b826102a1565b916123696040519384610212565b829481845281830111610185578281602093846000960137010152565b805115611e255760200190565b8051821015611e255760209160051b010190565b9080602083519182815201916020808360051b8301019401926000915b8383106123d357505050505090565b9091929394602080612436837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187526040838b5173ffffffffffffffffffffffffffffffffffffffff815116845201519181858201520190610313565b970193019301919392906123c4565b90610367916020815260c06125186124a161246d855160e060208701526101008601906123a7565b60208601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160408701526123a7565b60ff60408601511660608501526124c56060860151608086019063ffffffff169052565b608085015173ffffffffffffffffffffffffffffffffffffffff1660a085015260a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08583030184860152610313565b9201519060e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082850301910152610313565b916125729067ffffffffffffffff6103679593168452606060208501526060840190611944565b916040818403910152610313565b9015611e255790565b9190826040910312610185576040516125a1816101f6565b602080829480356125b181610647565b84520135910152565b9060a060806103679373ffffffffffffffffffffffffffffffffffffffff815116845267ffffffffffffffff602082015116602085015263ffffffff6040820151166040850152606081015160608501520151918160808201520190610313565b9080602083519182815201916020808360051b8301019401926000915b83831061264757505050505090565b90919293946020806126e2837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289519073ffffffffffffffffffffffffffffffffffffffff825116815260806126d16126b58685015160a08886015260a0850190610313565b6040850151604085015260608501518482036060860152610313565b9201519060808184039101526125ba565b97019301930191939290612638565b9080602083519182815201916020808360051b8301019401926000915b83831061271d57505050505090565b9091929394602080612759837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0866001960301875289516125ba565b9701930193019193929061270e565b610367916127a381835167ffffffffffffffff6060809280518552826020820151166020860152826040820151166040860152015116910152565b602082015173ffffffffffffffffffffffffffffffffffffffff1660808201526101206128576128436127fc6127ea60408701516101a060a08801526101a0870190610313565b606087015186820360c0880152610313565b608086015173ffffffffffffffffffffffffffffffffffffffff1660e086015260a086015161010086015260c08601518486015260e086015185820361014087015261261b565b6101008501518482036101608601526126f1565b920151906101808184039101526125ba565b906020610367928181520190612768565b906128848261217e565b6128916040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06128bf829461217e565b019060005b8281106128d057505050565b8060606020809385010152016128c4565b9291906128f8602091604086526040860190610313565b930152565b9061291090604083526040830190612768565b906020818303910152815180825260208201916020808360051b8301019401926000915b83831061294357505050505090565b909192939460208061297f837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610313565b97019301930191939290612934565b73ffffffffffffffffffffffffffffffffffffffff6001541633036129af57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9073ffffffffffffffffffffffffffffffffffffffff612aab9392604051938260208601947fa9059cbb000000000000000000000000000000000000000000000000000000008652166024860152604485015260448452612a3b606485610212565b16600080604093845195612a4f8688610212565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15612ad0573d612a9c612a93826102a1565b94519485610212565b83523d6000602085013e613ba1565b805180612ab6575050565b81602080612acb93610262950101910161182a565b613b16565b60609250613ba1565b916060610262929493612b268160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b6040519060e0820182811067ffffffffffffffff8211176101d557604052606060c08382815282602082015260006040820152600083820152600060808201528260a08201520152565b906004116101855790600490565b909291928360041161018557831161018557600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612c22575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b60408051909190612c658382610212565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b828110612c9c57505050565b602090604051612cab816101f6565b6000815260608382015282828501015201612c90565b90612ccb8261217e565b612cd86040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612d06829461217e565b019060005b828110612d1757505050565b602090604051612d26816101f6565b6000815260608382015282828501015201612d0b565b9080601f83011215610185578160206103679335910161234f565b9080601f8301121561018557813591612d6f8361217e565b92612d7d6040519485610212565b80845260208085019160051b830101918383116101855760208101915b838310612da957505050505090565b823567ffffffffffffffff81116101855782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083880301126101855760405190612df6826101f6565b6020830135612e0481610647565b825260408301359167ffffffffffffffff831161018557612e2d88602080969581960101612d3c565b83820152815201920191612d9a565b359060ff8216820361018557565b359063ffffffff8216820361018557565b6020818303126101855780359067ffffffffffffffff8211610185570160e08183031261018557612e8a610283565b91813567ffffffffffffffff81116101855781612ea8918401612d57565b8352602082013567ffffffffffffffff81116101855781612eca918401612d57565b6020840152612edb60408301612e3c565b6040840152612eec60608301612e4a565b6060840152612efd60808301610ad8565b608084015260a082013567ffffffffffffffff81116101855781612f22918401612d3c565b60a084015260c082013567ffffffffffffffff811161018557612f459201612d3c565b60c082015290565b929190612f58612b5b565b7f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000612fad612fa78686612ba5565b90612bee565b160361320c575081612fca92612fc292612bb3565b810190612e5b565b916020830190815151806131b7575b50612fec604092855151905151906122d5565b1561315f575b6080840173ffffffffffffffffffffffffffffffffffffffff613029825173ffffffffffffffffffffffffffffffffffffffff1690565b161561311a575b5001613053610533825173ffffffffffffffffffffffffffffffffffffffff1690565b61305a5750565b61308b61307061306b8551516122c7565b612cc1565b915173ffffffffffffffffffffffffffffffffffffffff1690565b6130b2613096610292565b73ffffffffffffffffffffffffffffffffffffffff9092168252565b6130ba6102db565b60208201526130c882612386565b526130d281612386565b5060005b83518051821015613114579061310d6130f182600194612393565b516130fb836122c7565b906131068287612393565b5284612393565b50016130d6565b50508252565b6131599061313f608084015173ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff169052565b38613030565b613167612c54565b84526131b161318d606083015173ffffffffffffffffffffffffffffffffffffffff1690565b613198613096610292565b6131a06102db565b602082015285519061146582612386565b50612ff2565b604085015160ff1690811080159190613200575b506131d65738612fd9565b7fb273a0e40000000000000000000000000000000000000000000000000000000060005260046000fd5b60ff91501615386131cb565b93906040926131a06131b19261322336848361234f565b60a0890152613230612c54565b8852606085015173ffffffffffffffffffffffffffffffffffffffff1692613275613259610292565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b369161234f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610ad35760010190565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610ad357565b60ff168015610ad3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b939190926133128551612cc1565b91600080935b87518510156134d25761334a613330868a989a612393565b5173ffffffffffffffffffffffffffffffffffffffff1690565b60009773ffffffffffffffffffffffffffffffffffffffff821697895b82518110156134c5576133806105336112de8386612393565b8a1461338e57600101613367565b50999692979190985060015b156133ac575b50506001019394613318565b6133ea906133ce6133c7869798999c959a969b949b9761327c565b9686612393565b519073ffffffffffffffffffffffffffffffffffffffff169052565b60005b85518110156134b4576134066105336112de8389612393565b8914613414576001016133ed565b613465919850602061342d828897969a959c9998612393565b510151602061344461343e876132a9565b8b612393565b51015261345a61345486516132a9565b86612393565b516131068287612393565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835101835260ff84166134a0575b6001905b90386133a0565b926134ac6001916132d6565b939050613495565b509650969392919094600190613499565b509996929791909861339a565b91965091949250816134e5575050929190565b91949092936134f861306b8588516122d5565b9260005b85811061354f57505060005b8651811015613545578061353e6135216001938a612393565b5161352c83896122d5565b906135378289612393565b5286612393565b5001613508565b5091945092909150565b8061355c60019284612393565b516135678288612393565b526135728187612393565b50016134fc565b6020818303126101855780519067ffffffffffffffff8211610185570160408183031261018557604051916135ad836101f6565b815167ffffffffffffffff811161018557816135ca9184016120d7565b8352602082015167ffffffffffffffff8111610185576135ea92016120d7565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080613624855184602087015260c0860190610313565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b919392613662612225565b506020830191825115613916576136966105336106ae610533875173ffffffffffffffffffffffffffffffffffffffff1690565b9173ffffffffffffffffffffffffffffffffffffffff831615801561388b575b61382857916000959661374b61378093889561372988519161370c6136ef8c5173ffffffffffffffffffffffffffffffffffffffff1690565b946136f8610274565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b6040519687809481937f9a4575b9000000000000000000000000000000000000000000000000000000008352600483016135f2565b03925af19283156105d857600093613803575b506137ba61379f612150565b925173ffffffffffffffffffffffffffffffffffffffff1690565b92602081519101519151906137ec6137d0610274565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b602085015260408401526060830152608082015290565b6138219193503d806000833e6138198183610212565b810190613579565b9138613793565b610614613849865173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481875afa9081156105d8576000916138f7575b50156136b6565b613910915060203d602011610640576106328183610212565b386138f0565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061397357505050505090565b90919293946020806139e1837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289519073ffffffffffffffffffffffffffffffffffffffff825116815260806126d16126b58685015160a08886015260a0850190610313565b97019301930191939290613964565b613b10613a14602083015173ffffffffffffffffffffffffffffffffffffffff1690565b82516060015167ffffffffffffffff16613a9c613a48608086015173ffffffffffffffffffffffffffffffffffffffff1690565b9161123160a08701516040519485936020850197889094939273ffffffffffffffffffffffffffffffffffffffff9067ffffffffffffffff6060948360808601991685521660208401521660408201520152565b5190206112316060840151602081519101209360e0604082015160208151910120910151604051613ad581611231602082019485613940565b51902090604051958694602086019889919260a093969594919660c08401976000855260208501526040840152606083015260808201520152565b51902090565b15613b1d57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91929015613c1c5750815115613bb5575090565b3b15613bbe5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015613c2f5750805190602001fd5b613c65906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610356565b0390fdfea164736f6c634300081a000a",
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
	return common.HexToHash("0xa816f7e08da08b1aa0143155f28f728327e40df7f707f612cb3566ab91229820")
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
