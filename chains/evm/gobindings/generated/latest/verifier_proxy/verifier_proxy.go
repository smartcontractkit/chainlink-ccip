// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_proxy

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
	FeeTokenAmount    *big.Int
	DestGasLimit      uint64
	DestBytesOverhead uint32
	ExtraArgs         []byte
}

type VerifierProxyDestChainConfigArgs struct {
	DestChainSelector uint64
	Router            common.Address
	DefaultVerifier   common.Address
	RequiredVerifier  common.Address
	DefaultExecutor   common.Address
}

type VerifierProxyDynamicConfig struct {
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	FeeAggregator          common.Address
}

type VerifierProxyStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	TokenAdminRegistry common.Address
}

var VerifierProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyDestChainConfigUpdates\",\"inputs\":[{\"name\":\"destChainConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structVerifierProxy.DestChainConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"contractIRouter\"},{\"name\":\"defaultVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultExecutor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardFromRouter\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExpectedNextSequenceNumber\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolBySourceToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoolV1\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.EVM2AnyVerifierMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structInternal.Header\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenTransfer\",\"type\":\"tuple\",\"internalType\":\"structInternal.EVMTokenTransfer\",\"components\":[{\"name\":\"sourceTokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"verifierReceipts\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.Receipt[]\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"executorReceipt\",\"type\":\"tuple\",\"internalType\":\"structInternal.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"tokenReceipt\",\"type\":\"tuple\",\"internalType\":\"structInternal.Receipt\",\"components\":[{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"receiptBlobs\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierProxy.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"contractIRMNRemote\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structVerifierProxy.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DestChainConfigSet\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIRouter\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CanOnlySendOneTokenPerMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotSendZeroTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"GetSupportedTokensFunctionalityRemovedCheckAdminRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDestChainConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidVerifierId\",\"inputs\":[{\"name\":\"verifierId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MalformedVerifierExtraArgs\",\"inputs\":[{\"name\":\"verifierExtraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MustBeCalledByRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RouterMustSetOriginalSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60e0604052346102fc57604051613a1738819003601f8101601f191683016001600160401b03811184821017610301578392829160405283398101039060c082126102fc57606082126102fc57610054610317565b81519092906001600160401b03811681036102fc5783526020820151906001600160a01b03821682036102fc5760208401918252606061009660408501610336565b6040860190815291605f1901126102fc576100af610317565b916100bc60608501610336565b835260808401519384151585036102fc5760a06100e0916020860196875201610336565b946040840195865233156102eb57600180546001600160a01b0319163317905580516001600160401b03161580156102d9575b80156102c7575b61029a57516001600160401b0316608052516001600160a01b0390811660a0529051811660c0528151161580156102b5575b80156102ab575b61029a5780516002805484516001600160a81b03199091166001600160a01b039384161790151560a01b60ff60a01b161790558351600380546001600160a01b031916919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a39260c092600060406101ce610317565b82815260208101839052015260805160a05185516001600160401b0390921694916001600160a01b0390811691166040610206610317565b878152602080820193845291019283526040805197885291516001600160a01b0390811691880191909152915182169086015290518116606085015290511515608084015290511660a0820152a16040516136cc908161034b8239608051818181610f150152818161112c0152611688015260a0518181816104bc01526116c1015260c0518181816116fd0152611a080152f35b6306b7c75960e31b60005260046000fd5b5081511515610153565b5082516001600160a01b03161561014c565b5082516001600160a01b03161561011a565b5081516001600160a01b031615610113565b639b15e16f60e01b60005260046000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b60405190606082016001600160401b0381118382101761030157604052565b51906001600160a01b03821682036102fc5756fe6080604052600436101561001257600080fd5b60003560e01c806306285c6914610107578063181f5a771461010257806320487ded146100fd57806348a98aa4146100f85780635147b35a146100f35780635cb80c5d146100ee5780636def4ce7146100e95780637437ff9f146100e457806379ba5097146100df5780638da5cb5b146100da5780639041be3d146100d557806390423fa2146100d0578063df0aa9e9146100cb578063f2fde38b146100c65763fbca3b74146100c157600080fd5b6115e7565b6114f3565b610cbd565b610ae9565b610a4f565b6109fd565b610914565b610848565b6107c9565b61074e565b6106d1565b610665565b610410565b61036a565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576060610140611668565b610183604051809273ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565bf35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff8211176101d557604052565b61018a565b60a0810190811067ffffffffffffffff8211176101d557604052565b6040810190811067ffffffffffffffff8211176101d557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101d557604052565b60405190610262608083610212565b565b6040519061026261016083610212565b6040519061026260a083610212565b6040519061026260c083610212565b60405190610262604083610212565b67ffffffffffffffff81116101d557601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b604051906102ea602083610212565b60008252565b60005b8381106103035750506000910152565b81810151838201526020016102f3565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361034f815180928187528780880191016102f0565b0116010190565b906020610367928181520190610313565b90565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576103e760408051906103ab8183610212565b601782527f566572696669657250726f787920312e372e302d646576000000000000000000602083015251918291602083526020830190610313565b0390f35b6004359067ffffffffffffffff8216820361018557565b908160a09103126101855790565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576104476103eb565b60243567ffffffffffffffff811161018557610467903690600401610402565b6040517f2cbc26bb00000000000000000000000000000000000000000000000000000000815277ffffffffffffffff00000000000000000000000000000000608084901b1660048201529091906020816024817f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff165afa9081156105d857600091610618575b506105dd576105829160209161054c61053361053360025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b906040518095819482937fd8694ccd0000000000000000000000000000000000000000000000000000000084526004840161183f565b03915afa80156105d8576103e7916000916105a9575b506040519081529081906020820190565b6105cb915060203d6020116105d1575b6105c38183610212565b810190611746565b38610598565b503d6105b9565b61173a565b7ffdbd6a720000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b6000fd5b61063a915060203d602011610640575b6106328183610212565b810190611725565b38610502565b503d610628565b73ffffffffffffffffffffffffffffffffffffffff81160361018557565b346101855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855761069c6103eb565b5060206106b36024356106ae81610647565b6119a9565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760043567ffffffffffffffff8111610185573660238201121561018557806004013567ffffffffffffffff81116101855736602460a083028401011161018557602461074c9201611aa2565b005b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760043567ffffffffffffffff8111610185573660238201121561018557806004013567ffffffffffffffff8111610185573660248260051b8401011161018557602461074c9201611e21565b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff6108096103eb565b1660005260046020526040806000205473ffffffffffffffffffffffffffffffffffffffff82519167ffffffffffffffff81168352831c166020820152f35b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855761087f611649565b5060405161088c816101b9565b60ff60025473ffffffffffffffffffffffffffffffffffffffff8116835260a01c161515602082015273ffffffffffffffffffffffffffffffffffffffff60035416604082015260405180916103e782606081019273ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855760005473ffffffffffffffffffffffffffffffffffffffff811633036109d3577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855767ffffffffffffffff610a8f6103eb565b166000526004602052600167ffffffffffffffff604060002054160167ffffffffffffffff8111610acf5760209067ffffffffffffffff60405191168152f35b611f4e565b359061026282610647565b8015150361018557565b346101855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610185576000604051610b26816101b9565b600435610b3281610647565b8152602435610b4081610adf565b6020820190815260443590610b5482610647565b60408301918252610b6361262d565b73ffffffffffffffffffffffffffffffffffffffff83511615918215610c9d575b508115610c92575b50610c6a5780516002805460208401517fffffffffffffffffffffff00000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff9384161790151560a01b74ff0000000000000000000000000000000000000000161790556040820151600380547fffffffffffffffffffffffff000000000000000000000000000000000000000016919092161790557f1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a390610c55611668565b610c6460405192839283612778565b0390a180f35b6004827f35be3ac8000000000000000000000000000000000000000000000000000000008152fd5b905051151538610b8c565b5173ffffffffffffffffffffffffffffffffffffffff1615915038610b84565b346101855760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261018557610cf46103eb565b60243567ffffffffffffffff811161018557610d14903690600401610402565b9060443591610d24606435610647565b60025460a01c60ff166114c957610d75740100000000000000000000000000000000000000007fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff6002541617600255565b610d9b610d968367ffffffffffffffff166000526004602052604060002090565b611f7d565b9073ffffffffffffffffffffffffffffffffffffffff606435161561149f57610de1610533610533602085015173ffffffffffffffffffffffffffffffffffffffff1690565b330361147557610e6491610e02610dfb6080840184611fe2565b9083612b88565b936000610e2a61053361053360025473ffffffffffffffffffffffffffffffffffffffff1690565b60208701519060405180809881947f9b1115e400000000000000000000000000000000000000000000000000000000835260048301610356565b03915afa9384156105d857600094611458575b5083511561142f575b610e886120c4565b95604086019485519660608101988998895192608001928351610eab9060ff1690565b91610eb593612df7565b60ff169092529088528087525195610ecb6120e0565b610ed361210e565b938651610ee79067ffffffffffffffff1690565b610ef090612144565b67ffffffffffffffff8116909752610f06610253565b6000815267ffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660208201529667ffffffffffffffff8716604089015267ffffffffffffffff166060880152610f676020890189611fe2565b9a90610f738a80611fe2565b90918c610f8260608e01611e17565b945151610f8e91612163565b610f9790612170565b9d610fa0610264565b9b8c5273ffffffffffffffffffffffffffffffffffffffff6064351660208d01523690610fcc926121dd565b60408b01523690610fdc926121dd565b606089015273ffffffffffffffffffffffffffffffffffffffff16608088015260a087015260c086016000905260e0860193845261010086019889528061012087015261014086015260005b8781106113c057505060005b885180518210156110d857906110d161104f82600194612221565b516020611070825173ffffffffffffffffffffffffffffffffffffffff1690565b91015161109a61107e610274565b73ffffffffffffffffffffffffffffffffffffffff9093168352565b60006020830152600060408301526000606083015260808201528a516110c08b85612163565b916110cb8383612221565b52612221565b5001611034565b505086939460408101916110ec8383612235565b905061134d575b5050604080517f130ac867e79e2789f923760a88743d292acdf7002139a588206e2260f73f73216020820190815267ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081169383019390935291851660608201523060808201526111a8935090915061119f8160a081015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610212565b51902084613446565b835152604051926111ec846111c08360208301612508565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610212565b6111f7835151612519565b9160005b845180518210156112af579060008161123b6105336105336112208461127298612221565b515173ffffffffffffffffffffffffffffffffffffffff1690565b89836040518097819582947f1234eab000000000000000000000000000000000000000000000000000000000845260048401612580565b03925af19182156105d85760019261128c575b50016111fb565b6112a8903d806000833e6112a08183610212565b810190612075565b5087611285565b50506103e7927f30d29bd09b10f6be582d3e2b98d565492453aa73d1892b2e10934ade0d9283ce67ffffffffffffffff6112f6606086510167ffffffffffffffff90511690565b61130b8260405193849316961694878361259c565b0390a361133b7fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff60025416600255565b51516040519081529081906020820190565b60016113598484612235565b905003611396578461137161138b9461138694612235565b9390611381606435953692612289565b612292565b613159565b9052838080806110f3565b7f68c2514e0000000000000000000000000000000000000000000000000000000060005260046000fd5b80611428816113d26001948651612221565b5160206113f3825173ffffffffffffffffffffffffffffffffffffffff1690565b91015161140161107e610274565b60006020830152600060408301526000606083015260808201528c51906110cb8383612221565b5001611028565b925061145261143e8380611fe2565b91906111736040519384926020840161209b565b92610e80565b61146e9194503d806000833e6112a08183610212565b9238610e77565b7f1c0a35290000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa4ec74790000000000000000000000000000000000000000000000000000000060005260046000fd5b7f3ee5aeb50000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855773ffffffffffffffffffffffffffffffffffffffff60043561154381610647565b61154b61262d565b163381146115bd57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101855761161e6103eb565b507f9e7177c80000000000000000000000000000000000000000000000000000000060005260046000fd5b60405190611656826101b9565b60006040838281528260208201520152565b611670611649565b5060405161167d816101b9565b67ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016604082015290565b90816020910312610185575161036781610adf565b6040513d6000823e3d90fd5b90816020910312610185575190565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561018557016020813591019167ffffffffffffffff821161018557813603831361018557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9160209082815201919060005b8181106117fe5750505090565b90919260408060019273ffffffffffffffffffffffffffffffffffffffff873561182781610647565b168152602087810135908201520194019291016117f1565b919067ffffffffffffffff168252604060208301526118b26118756118648380611755565b60a0604087015260e08601916117a5565b6118826020840184611755565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08684030160608701526117a5565b9160408201357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18336030181121561018557820160208135910167ffffffffffffffff8211610185578160061b36038113610185578461197992611942927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08661036799030160808701526117e4565b9261196f61195260608301610ad4565b73ffffffffffffffffffffffffffffffffffffffff1660a0850152565b6080810190611755565b9160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0828603019101526117a5565b73ffffffffffffffffffffffffffffffffffffffff604051917fbbe4f6db00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156105d857600090611a53575b73ffffffffffffffffffffffffffffffffffffffff91501690565b506020813d602011611a9a575b81611a6d60209383610212565b810103126101855773ffffffffffffffffffffffffffffffffffffffff9051611a9581610647565b611a38565b3d9150611a60565b90611aab61262d565b60005b818110611aba57505050565b611acd611ac8828486611d76565b611d8b565b805167ffffffffffffffff1690818015611d10576001939291611cd8611b287ff2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed09367ffffffffffffffff166000526004602052604060002090565b611cca611c866080611b51602087015173ffffffffffffffffffffffffffffffffffffffff1690565b84547fffffffff0000000000000000000000000000000000000000ffffffffffffffff16604082901b7bffffffffffffffffffffffffffffffffffffffff00000000000000001617855595611c03611bc0604083015173ffffffffffffffffffffffffffffffffffffffff1690565b8c87019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b611c6b611c27606083015173ffffffffffffffffffffffffffffffffffffffff1690565b600287019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b015173ffffffffffffffffffffffffffffffffffffffff1690565b600383019073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5467ffffffffffffffff1690565b6040805167ffffffffffffffff92909216825273ffffffffffffffffffffffffffffffffffffffff929092166020820152a201611aae565b7fc35aa79d0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015611d865760a0020190565b611d47565b60a0813603126101855760405190611da2826101da565b80359067ffffffffffffffff821682036101855760809183526020810135611dc981610647565b60208401526040810135611ddc81610647565b60408401526060810135611def81610647565b60608401520135611dff81610647565b608082015290565b9190811015611d865760051b0190565b3561036781610647565b60035473ffffffffffffffffffffffffffffffffffffffff169160005b818110611e4b5750505050565b611e61610533611e5c838587611e07565b611e17565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290919073ffffffffffffffffffffffffffffffffffffffff831690602081602481855afa9081156105d8576001948891600093611f2e575b5082611ed6575b5050505001611e3e565b611ee1918391612678565b60405190815273ffffffffffffffffffffffffffffffffffffffff8716907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e90602090a338808681611ecc565b611f4791935060203d81116105d1576105c38183610212565b9138611ec5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90604051611f8a816101da565b608073ffffffffffffffffffffffffffffffffffffffff806003849682815467ffffffffffffffff8116885260401c166020870152826001820154166040870152828060028301541616606087015201541616910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff82116101855760200191813603831361018557565b81601f82011215610185578051612049816102a1565b926120576040519485610212565b818452602082840101116101855761036791602080850191016102f0565b9060208282031261018557815167ffffffffffffffff8111610185576103679201612033565b9160206103679381815201916117a5565b67ffffffffffffffff81116101d55760051b60200190565b604051906120d3602083610212565b6000808352366020840137565b604051906120ed826101da565b60606080836000815260006020820152600060408201526000838201520152565b604051906080820182811067ffffffffffffffff8211176101d55760405260006060838281528160208201528160408201520152565b67ffffffffffffffff1667ffffffffffffffff8114610acf5760010190565b91908201809211610acf57565b9061217a826120ac565b6121876040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06121b582946120ac565b019060005b8281106121c657505050565b6020906121d16120e0565b828285010152016121ba565b9291926121e9826102a1565b916121f76040519384610212565b829481845281830111610185578281602093846000960137010152565b805115611d865760200190565b8051821015611d865760209160051b010190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610185570180359067ffffffffffffffff821161018557602001918160061b3603831361018557565b9015611d865790565b9190826040910312610185576040516122aa816101f6565b602080829480356122ba81610647565b84520135910152565b9073ffffffffffffffffffffffffffffffffffffffff825116815260608061230f6122fd6020860151608060208701526080860190610313565b60408601518582036040870152610313565b93015191015290565b9060a060806103679373ffffffffffffffffffffffffffffffffffffffff81511684526020810151602085015267ffffffffffffffff604082015116604085015263ffffffff60608201511660608501520151918160808201520190610313565b9080602083519182815201916020808360051b8301019401926000915b8383106123a557505050505090565b90919293946020806123e1837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951612318565b97019301930191939290612396565b6103679161242b81835167ffffffffffffffff6060809280518552826020820151166020860152826040820151166040860152015116910152565b602082015173ffffffffffffffffffffffffffffffffffffffff1660808201526101406124f66124e26124ce61248761247560408801516101c060a08901526101c0880190610313565b606088015187820360c0890152610313565b608087015173ffffffffffffffffffffffffffffffffffffffff1660e087015260a087015161010087015260c087015161012087015260e0870151868203868801526122c3565b610100860151858203610160870152612379565b610120850151848203610180860152612318565b920151906101a0818403910152612318565b9060206103679281815201906123f0565b90612523826120ac565b6125306040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061255e82946120ac565b019060005b82811061256f57505050565b806060602080938501015201612563565b929190612597602091604086526040860190610313565b930152565b906125af906040835260408301906123f0565b906020818303910152815180825260208201916020808360051b8301019401926000915b8383106125e257505050505090565b909192939460208061261e837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe086600196030187528951610313565b970193019301919392906125d3565b73ffffffffffffffffffffffffffffffffffffffff60015416330361264e57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9073ffffffffffffffffffffffffffffffffffffffff61274a9392604051938260208601947fa9059cbb0000000000000000000000000000000000000000000000000000000086521660248601526044850152604484526126da606485610212565b166000806040938451956126ee8688610212565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d1561276f573d61273b612732826102a1565b94519485610212565b83523d6000602085013e6135f7565b805180612755575050565b8160208061276a936102629501019101611725565b61356c565b606092506135f7565b9160606102629294936127c58160c081019773ffffffffffffffffffffffffffffffffffffffff6040809267ffffffffffffffff8151168552826020820151166020860152015116910152565b019073ffffffffffffffffffffffffffffffffffffffff60408092828151168552602081015115156020860152015116910152565b6040519060c0820182811067ffffffffffffffff8211176101d557604052606060a083600081528260208201528260408201528280820152600060808201520152565b906004116101855790600490565b919091357fffffffff000000000000000000000000000000000000000000000000000000008116926004811061287f575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b604080519091906128c28382610212565b60018152917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018260005b8281106128f957505050565b602090604051612908816101f6565b60008152606083820152828285010152016128ed565b90612928826120ac565b6129356040519182610212565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061296382946120ac565b019060005b82811061297457505050565b602090604051612983816101f6565b6000815260608382015282828501015201612968565b9080601f8301121561018557816020610367933591016121dd565b9080601f83011215610185578135916129cc836120ac565b926129da6040519485610212565b80845260208085019160051b830101918383116101855760208101915b838310612a0657505050505090565b823567ffffffffffffffff81116101855782019060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083880301126101855760405190612a53826101f6565b6020830135612a6181610647565b825260408301359167ffffffffffffffff831161018557612a8a88602080969581960101612999565b838201528152019201916129f7565b359060ff8216820361018557565b6020818303126101855780359067ffffffffffffffff8211610185570160c08183031261018557612ad6610283565b91612ae082610ad4565b8352602082013567ffffffffffffffff81116101855781612b02918401612999565b6020840152604082013567ffffffffffffffff81116101855781612b279184016129b4565b6040840152606082013567ffffffffffffffff81116101855781612b4c9184016129b4565b6060840152612b5d60808301612a99565b608084015260a082013567ffffffffffffffff811161018557612b809201612999565b60a082015290565b90612b916127fa565b927f302326cb000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000612be7612be1848661283d565b9061284b565b1603612cf257612bfb929350810190612aa7565b9060608201515180612cd6575b612c186040840191825151612163565b15612c7c575b505b73ffffffffffffffffffffffffffffffffffffffff612c53835173ffffffffffffffffffffffffffffffffffffffff1690565b1615612c5d575090565b6080015173ffffffffffffffffffffffffffffffffffffffff16815290565b612ccf90612c886128b1565b8152604083015173ffffffffffffffffffffffffffffffffffffffff1690612cb161107e610292565b612cb96102db565b60208301525190612cc982612214565b52612214565b5038612c1e565b608083015160ff16808211156101855760ff16612c0857600080fd5b612d6991612d013683836121dd565b6020860152612cb9612d116128b1565b9160408701928352612d3a604087015173ffffffffffffffffffffffffffffffffffffffff1690565b93612d62612d46610292565b73ffffffffffffffffffffffffffffffffffffffff9096168652565b36916121dd565b50612c20565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610acf5760010190565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211610acf57565b60ff168015610acf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b919092612e04835161291e565b91600080935b8551851015612fcb57612e3c612e2286889a98612221565b5173ffffffffffffffffffffffffffffffffffffffff1690565b60009773ffffffffffffffffffffffffffffffffffffffff821697895b8251811015612fbb57612e726105336112208386612221565b8a14612e8057600101612e59565b50969297919098506001999394995b15612ea4575b50506001019396919094612e0a565b612ee190612ec5612ebe879c9899949b959a969c97612d6f565b9686612221565b519073ffffffffffffffffffffffffffffffffffffffff169052565b60005b8551811015612fab57612efd6105336112208389612221565b8914612f0b57600101612ee4565b859993979298506020612f2382612f5c949998612221565b5101516020612f3a612f3488612d9c565b8b612221565b510152612f4a612f348b51612d9c565b51612f55828c612221565b5289612221565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff885101885260ff8416612f97575b6001905b9038612e95565b92612fa3600191612dc9565b939050612f8c565b5094939791956001919750612f90565b5096929791909899939499612e8f565b919694909295935082612fe057505050929190565b612ff583612ffa929894969793959851612163565b61291e565b9260005b85811061305157505060005b865181101561304757806130406130236001938a612221565b5161302e8389612163565b906130398289612221565b5286612221565b500161300a565b5091945092909150565b8061305e60019284612221565b516130698288612221565b526130748187612221565b5001612ffe565b6020818303126101855780519067ffffffffffffffff8211610185570160408183031261018557604051916130af836101f6565b815167ffffffffffffffff811161018557816130cc918401612033565b8352602082015167ffffffffffffffff8111610185576130ec9201612033565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff6080613126855184602087015260c0860190610313565b9467ffffffffffffffff602082015116604086015282604082015116606086015260608101518286015201511691015290565b9192909261316561210e565b50602083019384511561340b576131996105336106ae610533875173ffffffffffffffffffffffffffffffffffffffff1690565b9273ffffffffffffffffffffffffffffffffffffffff8416158015613380575b61331d5792600094939261324f61328493879561322d8a51916132106131f38b5173ffffffffffffffffffffffffffffffffffffffff1690565b946131fc610274565b97885267ffffffffffffffff166020880152565b73ffffffffffffffffffffffffffffffffffffffff166040860152565b606084015273ffffffffffffffffffffffffffffffffffffffff166080830152565b6040519586809481937f9a4575b9000000000000000000000000000000000000000000000000000000008352600483016130f4565b03925af19182156105d8576000926132f4575b505173ffffffffffffffffffffffffffffffffffffffff165b91602082519201519051916132e26132c6610253565b73ffffffffffffffffffffffffffffffffffffffff9095168552565b60208401526040830152606082015290565b6132b0919250613316903d806000833e61330e8183610212565b81019061307b565b9190613297565b61061461333e865173ffffffffffffffffffffffffffffffffffffffff1690565b7fbf16aab60000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527faff2afbf000000000000000000000000000000000000000000000000000000006004820152602081602481885afa9081156105d8576000916133ec575b50156131b9565b613405915060203d602011610640576106328183610212565b386133e5565b7f5cf044490000000000000000000000000000000000000000000000000000000060005260046000fd5b9060206103679281815201906122c3565b61356661346a602083015173ffffffffffffffffffffffffffffffffffffffff1690565b82516060015167ffffffffffffffff166134f261349e608086015173ffffffffffffffffffffffffffffffffffffffff1690565b9161117360a08701516040519485936020850197889094939273ffffffffffffffffffffffffffffffffffffffff9067ffffffffffffffff6060948360808601991685521660208401521660408201520152565b5190206111736060840151602081519101209360e060408201516020815191012091015160405161352b81611173602082019485613435565b51902090604051958694602086019889919260a093969594919660c08401976000855260208501526040840152606083015260808201520152565b51902090565b1561357357565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b91929015613672575081511561360b575090565b3b156136145790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156136855750805190602001fd5b6136bb906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260048301610356565b0390fdfea164736f6c634300081a000a",
}

var VerifierProxyABI = VerifierProxyMetaData.ABI

var VerifierProxyBin = VerifierProxyMetaData.Bin

func DeployVerifierProxy(auth *bind.TransactOpts, backend bind.ContractBackend, staticConfig VerifierProxyStaticConfig, dynamicConfig VerifierProxyDynamicConfig) (common.Address, *types.Transaction, *VerifierProxy, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierProxyBin), backend, staticConfig, dynamicConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierProxy{address: address, abi: *parsed, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

type VerifierProxy struct {
	address common.Address
	abi     abi.ABI
	VerifierProxyCaller
	VerifierProxyTransactor
	VerifierProxyFilterer
}

type VerifierProxyCaller struct {
	contract *bind.BoundContract
}

type VerifierProxyTransactor struct {
	contract *bind.BoundContract
}

type VerifierProxyFilterer struct {
	contract *bind.BoundContract
}

type VerifierProxySession struct {
	Contract     *VerifierProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierProxyCallerSession struct {
	Contract *VerifierProxyCaller
	CallOpts bind.CallOpts
}

type VerifierProxyTransactorSession struct {
	Contract     *VerifierProxyTransactor
	TransactOpts bind.TransactOpts
}

type VerifierProxyRaw struct {
	Contract *VerifierProxy
}

type VerifierProxyCallerRaw struct {
	Contract *VerifierProxyCaller
}

type VerifierProxyTransactorRaw struct {
	Contract *VerifierProxyTransactor
}

func NewVerifierProxy(address common.Address, backend bind.ContractBackend) (*VerifierProxy, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierProxy{address: address, abi: abi, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

func NewVerifierProxyCaller(address common.Address, caller bind.ContractCaller) (*VerifierProxyCaller, error) {
	contract, err := bindVerifierProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCaller{contract: contract}, nil
}

func NewVerifierProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierProxyTransactor, error) {
	contract, err := bindVerifierProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyTransactor{contract: contract}, nil
}

func NewVerifierProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierProxyFilterer, error) {
	contract, err := bindVerifierProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFilterer{contract: contract}, nil
}

func bindVerifierProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierProxy *VerifierProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.VerifierProxyCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCaller) GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

	error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getDestChainConfig", destChainSelector)

	outstruct := new(GetDestChainConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_VerifierProxy *VerifierProxySession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _VerifierProxy.Contract.GetDestChainConfig(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetDestChainConfig(destChainSelector uint64) (GetDestChainConfig,

	error) {
	return _VerifierProxy.Contract.GetDestChainConfig(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCaller) GetDynamicConfig(opts *bind.CallOpts) (VerifierProxyDynamicConfig, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(VerifierProxyDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierProxyDynamicConfig)).(*VerifierProxyDynamicConfig)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetDynamicConfig() (VerifierProxyDynamicConfig, error) {
	return _VerifierProxy.Contract.GetDynamicConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetDynamicConfig() (VerifierProxyDynamicConfig, error) {
	return _VerifierProxy.Contract.GetDynamicConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getExpectedNextSequenceNumber", destChainSelector)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _VerifierProxy.Contract.GetExpectedNextSequenceNumber(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetExpectedNextSequenceNumber(destChainSelector uint64) (uint64, error) {
	return _VerifierProxy.Contract.GetExpectedNextSequenceNumber(&_VerifierProxy.CallOpts, destChainSelector)
}

func (_VerifierProxy *VerifierProxyCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getFee", destChainSelector, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _VerifierProxy.Contract.GetFee(&_VerifierProxy.CallOpts, destChainSelector, message)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetFee(destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error) {
	return _VerifierProxy.Contract.GetFee(&_VerifierProxy.CallOpts, destChainSelector, message)
}

func (_VerifierProxy *VerifierProxyCaller) GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getPoolBySourceToken", arg0, sourceToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _VerifierProxy.Contract.GetPoolBySourceToken(&_VerifierProxy.CallOpts, arg0, sourceToken)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetPoolBySourceToken(arg0 uint64, sourceToken common.Address) (common.Address, error) {
	return _VerifierProxy.Contract.GetPoolBySourceToken(&_VerifierProxy.CallOpts, arg0, sourceToken)
}

func (_VerifierProxy *VerifierProxyCaller) GetStaticConfig(opts *bind.CallOpts) (VerifierProxyStaticConfig, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(VerifierProxyStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierProxyStaticConfig)).(*VerifierProxyStaticConfig)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetStaticConfig() (VerifierProxyStaticConfig, error) {
	return _VerifierProxy.Contract.GetStaticConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetStaticConfig() (VerifierProxyStaticConfig, error) {
	return _VerifierProxy.Contract.GetStaticConfig(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getSupportedTokens", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _VerifierProxy.Contract.GetSupportedTokens(&_VerifierProxy.CallOpts, arg0)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetSupportedTokens(arg0 uint64) ([]common.Address, error) {
	return _VerifierProxy.Contract.GetSupportedTokens(&_VerifierProxy.CallOpts, arg0)
}

func (_VerifierProxy *VerifierProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierProxy *VerifierProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "applyDestChainConfigUpdates", destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxySession) ApplyDestChainConfigUpdates(destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ApplyDestChainConfigUpdates(&_VerifierProxy.TransactOpts, destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxyTransactorSession) ApplyDestChainConfigUpdates(destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ApplyDestChainConfigUpdates(&_VerifierProxy.TransactOpts, destChainConfigArgs)
}

func (_VerifierProxy *VerifierProxyTransactor) ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "forwardFromRouter", destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxySession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ForwardFromRouter(&_VerifierProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxyTransactorSession) ForwardFromRouter(destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.ForwardFromRouter(&_VerifierProxy.TransactOpts, destChainSelector, message, feeTokenAmount, originalSender)
}

func (_VerifierProxy *VerifierProxyTransactor) SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setDynamicConfig", dynamicConfig)
}

func (_VerifierProxy *VerifierProxySession) SetDynamicConfig(dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetDynamicConfig(&_VerifierProxy.TransactOpts, dynamicConfig)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetDynamicConfig(dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetDynamicConfig(&_VerifierProxy.TransactOpts, dynamicConfig)
}

func (_VerifierProxy *VerifierProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierProxy *VerifierProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_VerifierProxy *VerifierProxySession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.WithdrawFeeTokens(&_VerifierProxy.TransactOpts, feeTokens)
}

func (_VerifierProxy *VerifierProxyTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.WithdrawFeeTokens(&_VerifierProxy.TransactOpts, feeTokens)
}

type VerifierProxyCCIPMessageSentIterator struct {
	Event *VerifierProxyCCIPMessageSent

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyCCIPMessageSentIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyCCIPMessageSent)
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
		it.Event = new(VerifierProxyCCIPMessageSent)
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

func (it *VerifierProxyCCIPMessageSentIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           InternalEVM2AnyVerifierMessage
	ReceiptBlobs      [][]byte
	Raw               types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierProxyCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCCIPMessageSentIterator{contract: _VerifierProxy.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyCCIPMessageSent)
				if err := _VerifierProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseCCIPMessageSent(log types.Log) (*VerifierProxyCCIPMessageSent, error) {
	event := new(VerifierProxyCCIPMessageSent)
	if err := _VerifierProxy.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyConfigSetIterator struct {
	Event *VerifierProxyConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyConfigSet)
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
		it.Event = new(VerifierProxyConfigSet)
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

func (it *VerifierProxyConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyConfigSet struct {
	StaticConfig  VerifierProxyStaticConfig
	DynamicConfig VerifierProxyDynamicConfig
	Raw           types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VerifierProxyConfigSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyConfigSetIterator{contract: _VerifierProxy.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyConfigSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyConfigSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseConfigSet(log types.Log) (*VerifierProxyConfigSet, error) {
	event := new(VerifierProxyConfigSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyDestChainConfigSetIterator struct {
	Event *VerifierProxyDestChainConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyDestChainConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyDestChainConfigSet)
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
		it.Event = new(VerifierProxyDestChainConfigSet)
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

func (it *VerifierProxyDestChainConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyDestChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyDestChainConfigSet struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Router            common.Address
	Raw               types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierProxyDestChainConfigSetIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyDestChainConfigSetIterator{contract: _VerifierProxy.contract, event: "DestChainConfigSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "DestChainConfigSet", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyDestChainConfigSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseDestChainConfigSet(log types.Log) (*VerifierProxyDestChainConfigSet, error) {
	event := new(VerifierProxyDestChainConfigSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "DestChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyFeeTokenWithdrawnIterator struct {
	Event *VerifierProxyFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyFeeTokenWithdrawn)
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
		it.Event = new(VerifierProxyFeeTokenWithdrawn)
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

func (it *VerifierProxyFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyFeeTokenWithdrawn struct {
	FeeAggregator common.Address
	FeeToken      common.Address
	Amount        *big.Int
	Raw           types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*VerifierProxyFeeTokenWithdrawnIterator, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFeeTokenWithdrawnIterator{contract: _VerifierProxy.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var feeAggregatorRule []interface{}
	for _, feeAggregatorItem := range feeAggregator {
		feeAggregatorRule = append(feeAggregatorRule, feeAggregatorItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "FeeTokenWithdrawn", feeAggregatorRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyFeeTokenWithdrawn)
				if err := _VerifierProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseFeeTokenWithdrawn(log types.Log) (*VerifierProxyFeeTokenWithdrawn, error) {
	event := new(VerifierProxyFeeTokenWithdrawn)
	if err := _VerifierProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferRequestedIterator struct {
	Event *VerifierProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferRequested)
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
		it.Event = new(VerifierProxyOwnershipTransferRequested)
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

func (it *VerifierProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferRequestedIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferRequested)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error) {
	event := new(VerifierProxyOwnershipTransferRequested)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferredIterator struct {
	Event *VerifierProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferred)
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
		it.Event = new(VerifierProxyOwnershipTransferred)
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

func (it *VerifierProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferredIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferred)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error) {
	event := new(VerifierProxyOwnershipTransferred)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDestChainConfig struct {
	SequenceNumber uint64
	Router         common.Address
}

func (_VerifierProxy *VerifierProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierProxy.abi.Events["CCIPMessageSent"].ID:
		return _VerifierProxy.ParseCCIPMessageSent(log)
	case _VerifierProxy.abi.Events["ConfigSet"].ID:
		return _VerifierProxy.ParseConfigSet(log)
	case _VerifierProxy.abi.Events["DestChainConfigSet"].ID:
		return _VerifierProxy.ParseDestChainConfigSet(log)
	case _VerifierProxy.abi.Events["FeeTokenWithdrawn"].ID:
		return _VerifierProxy.ParseFeeTokenWithdrawn(log)
	case _VerifierProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _VerifierProxy.ParseOwnershipTransferRequested(log)
	case _VerifierProxy.abi.Events["OwnershipTransferred"].ID:
		return _VerifierProxy.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierProxyCCIPMessageSent) Topic() common.Hash {
	return common.HexToHash("0x30d29bd09b10f6be582d3e2b98d565492453aa73d1892b2e10934ade0d9283ce")
}

func (VerifierProxyConfigSet) Topic() common.Hash {
	return common.HexToHash("0x1266079276a6f57589aa41ba2b2485823d246a0de19b10bf77d954f2a83745a3")
}

func (VerifierProxyDestChainConfigSet) Topic() common.Hash {
	return common.HexToHash("0xf2c5b50c4521263fb32bdc393c35317cfdea4fa5bcc42315b241f7e00841bed0")
}

func (VerifierProxyFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (VerifierProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_VerifierProxy *VerifierProxy) Address() common.Address {
	return _VerifierProxy.address
}

type VerifierProxyInterface interface {
	GetDestChainConfig(opts *bind.CallOpts, destChainSelector uint64) (GetDestChainConfig,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (VerifierProxyDynamicConfig, error)

	GetExpectedNextSequenceNumber(opts *bind.CallOpts, destChainSelector uint64) (uint64, error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, message ClientEVM2AnyMessage) (*big.Int, error)

	GetPoolBySourceToken(opts *bind.CallOpts, arg0 uint64, sourceToken common.Address) (common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (VerifierProxyStaticConfig, error)

	GetSupportedTokens(opts *bind.CallOpts, arg0 uint64) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyDestChainConfigUpdates(opts *bind.TransactOpts, destChainConfigArgs []VerifierProxyDestChainConfigArgs) (*types.Transaction, error)

	ForwardFromRouter(opts *bind.TransactOpts, destChainSelector uint64, message ClientEVM2AnyMessage, feeTokenAmount *big.Int, originalSender common.Address) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, dynamicConfig VerifierProxyDynamicConfig) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*VerifierProxyCCIPMessageSentIterator, error)

	WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *VerifierProxyCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error)

	ParseCCIPMessageSent(log types.Log) (*VerifierProxyCCIPMessageSent, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VerifierProxyConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VerifierProxyConfigSet, error)

	FilterDestChainConfigSet(opts *bind.FilterOpts, destChainSelector []uint64) (*VerifierProxyDestChainConfigSetIterator, error)

	WatchDestChainConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyDestChainConfigSet, destChainSelector []uint64) (event.Subscription, error)

	ParseDestChainConfigSet(log types.Log) (*VerifierProxyDestChainConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, feeAggregator []common.Address, feeToken []common.Address) (*VerifierProxyFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeTokenWithdrawn, feeAggregator []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*VerifierProxyFeeTokenWithdrawn, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
