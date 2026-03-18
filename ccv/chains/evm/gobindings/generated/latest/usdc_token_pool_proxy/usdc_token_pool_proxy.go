// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc_token_pool_proxy

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

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead                         uint32
	DestBytesOverhead                       uint32
	DefaultBlockConfirmationsFeeUSDCents    uint32
	CustomBlockConfirmationsFeeUSDCents     uint32
	DefaultBlockConfirmationsTransferFeeBps uint16
	CustomBlockConfirmationsTransferFeeBps  uint16
	IsEnabled                               bool
}

type PoolLockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      common.Address
	Amount              *big.Int
	LocalToken          common.Address
}

type PoolLockOrBurnOutV1 struct {
	DestTokenAddress []byte
	DestPoolData     []byte
}

type PoolReleaseOrMintInV1 struct {
	OriginalSender          []byte
	RemoteChainSelector     uint64
	Receiver                common.Address
	SourceDenominatedAmount *big.Int
	LocalToken              common.Address
	SourcePoolAddress       []byte
	SourcePoolData          []byte
	OffchainTokenData       []byte
}

type PoolReleaseOrMintOutV1 struct {
	DestinationAmount *big.Int
}

type USDCTokenPoolProxyPoolAddresses struct {
	CctpV1Pool            common.Address
	CctpV2Pool            common.Address
	CctpV2PoolWithCCV     common.Address
	SiloedLockReleasePool common.Address
}

var USDCTokenPoolProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeAggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeAggregator\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainNotSupportedByVerifier\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSetPoolForMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"NoLockOrBurnMechanismSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAddressCannotBeSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenPoolUnsupported\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e0806040523461043e5780613b7a803803809161001d8285610443565b833981010360e0811261043e578151906001600160a01b0382169081830361043e57608090601f19011261043e5760405191608083016001600160401b038111848210176104285760405261007460208501610466565b835261008260408501610466565b916020840192835261009660608601610466565b94604085019586526100aa60808201610466565b92606086019384526100ca60c06100c360a08501610466565b9301610466565b92331561041757600180546001600160a01b03191633179055158015610406575b80156103f5575b6103e4576080526001600160a01b0390811660a05290811660c05283511680151590816103d3575b506103b15781516001600160a01b031680151590816103a0575b5061037e5783516001600160a01b0316801515908161036d575b5061034b5780516001600160a01b0316801515908161033a575b508061031f575b6102fe5782516001600160a01b0316301480156102eb575b80156102d8575b80156102c5575b6102b4579151600380546001600160a01b03199081166001600160a01b039384169081179092558351600480548316918516919091179055855160058054831691851691909117905584516006805490921690841617905560408051918252925182166020820152935181169184019190915290511660608201527f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa890608090a16040516135bb90816105bf823960805181818161091901528181610b230152818161100f01528181611c1001528181611c620152611d77015260a05181818161077a01528181610f3d01528181611db00152818161263c0152612938015260c05181818161089a01528181611dec01528181612beb0152612d110152f35b636a5db6d560e11b60005260046000fd5b5080516001600160a01b03163014610195565b5083516001600160a01b0316301461018e565b5081516001600160a01b03163014610187565b5163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b508051610334906001600160a01b03166104bb565b1561016f565b610344915061047a565b1538610168565b835163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61037791506104bb565b153861014e565b505163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b6103aa915061047a565b1538610134565b825163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b6103dd915061047a565b153861011a565b6303988b8160e61b60005260046000fd5b506001600160a01b038316156100f2565b506001600160a01b038216156100eb565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761042857604052565b51906001600160a01b038216820361043e57565b610483816104f9565b90816104a9575b81610493575090565b6104a69150630e64dd2960e01b9061058a565b90565b90506104b481610558565b159061048a565b6104c4816104f9565b90816104e7575b816104d4575090565b6104a69150633317103160e01b9061058a565b90506104f281610558565b15906104cb565b6000602091604051838101906301ffc9a760e01b82526301ffc9a760e01b60248201526024815261052b604482610443565b5191617530fa6000513d8261054c575b5081610545575090565b9050151590565b6020111591503861053b565b6000602091604051838101906301ffc9a760e01b825263ffffffff60e01b60248201526024815261052b604482610443565b600090602092604051848101916301ffc9a760e01b835263ffffffff60e01b1660248201526024815261052b60448261044356fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714611e135750806306285c6914611d4a57806315b358e014611ce5578063181f5a7714611c8657806321df0da714611c35578063240028e814611bc45780632c06340414611b27578063309292ac1461167d578063390775371461163d578063489a68f2146115eb5780635cb80c5d1461146b578063673a2a1f1461136657806379ba50971461129b5780638926f54f1461125157806389720a62146111955780638da5cb5b146111615780639a4575b914610ea25780639cb406c914610e6e578063a42a7b8b14610c6f578063aa86a75414610c2a578063b1c71c651461067c578063b794658014610582578063d8aa3f40146104bb578063db4c2aef146102095763f2fde38b1461013257600080fd5b346102045760206003193601126102045773ffffffffffffffffffffffffffffffffffffffff610160611f13565b610168613048565b163381146101da57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102045760406003193601126102045760043567ffffffffffffffff81116102045761023a90369060040161218d565b9060243567ffffffffffffffff81116102045761025b90369060040161218d565b610266929192613048565b8084036104915760005b84811061027957005b610284818386613038565b3560058110156102045767ffffffffffffffff6102aa6102a5848988613038565b6123f9565b16600052600260205260406000206000907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff8416179055506003811480610471575b6103e45760006001821480610451575b61037957506002811480610431575b6103e457600060048214806103c4575b6103795750906001917f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc94471602067ffffffffffffffff6103636102a5868c8b613038565b16926103726040518092612207565ba201610270565b6103c260449267ffffffffffffffff6103966102a5878c8b613038565b7f87d77d33000000000000000000000000000000000000000000000000000000008552166004526121f9565bfd5b5073ffffffffffffffffffffffffffffffffffffffff6005541615610320565b67ffffffffffffffff6103fe6102a561042b948988613038565b7f87d77d3300000000000000000000000000000000000000000000000000000000600052166004526121f9565b60446000fd5b5073ffffffffffffffffffffffffffffffffffffffff6004541615610310565b5073ffffffffffffffffffffffffffffffffffffffff6003541615610301565b5073ffffffffffffffffffffffffffffffffffffffff60065416156102f1565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b34610204576080600319360112610204576104d4611f13565b6104dc6120fa565b906044359161ffff83168303610204576064359167ffffffffffffffff83116102045760e09361051361051b94369060040161215f565b939092612e9f565b60c06040519163ffffffff815116835263ffffffff602082015116602084015263ffffffff604082015116604084015263ffffffff606082015116606084015261ffff608082015116608084015261ffff60a08201511660a08401520151151560c0820152f35b346102045760206003193601126102045761059b612111565b600067ffffffffffffffff60246105b184613093565b509373ffffffffffffffffffffffffffffffffffffffff60405195869485937fb7946580000000000000000000000000000000000000000000000000000000008552166004840152165afa801561067057600090610626575b610622906040519182916020835260208301906120b7565b0390f35b3d8082843e6106358184612019565b82019160208184031261066c5780519167ffffffffffffffff8311610669575091610664916106229301612d7f565b61060a565b80fd5b5080fd5b6040513d6000823e3d90fd5b346102045760606003193601126102045760043567ffffffffffffffff81116102045760a06003198236030112610204576106b561213d565b6044359167ffffffffffffffff831161020457366023840112156102045782600401356106e18161205a565b936106ef6040519586612019565b818552366024838301011161020457816000926024602093018388013785010152610718612d3a565b506024810191610727836123f9565b67ffffffffffffffff604051917fa8d87a3b00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561067057600091610c0b575b5073ffffffffffffffffffffffffffffffffffffffff33911603610bdd5767ffffffffffffffff6107db846123f9565b16600052600260205260ff604060002054166005811015610bae578015610b7b5760009060048103610ac957505073ffffffffffffffffffffffffffffffffffffffff60055416928315610a8957610832816123f9565b9067ffffffffffffffff604051927f958021a70000000000000000000000000000000000000000000000000000000084521660048301526040602483015260208280610881604482018a6120b7565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa91821561067057600092610a58575b5073ffffffffffffffffffffffffffffffffffffffff821615610a185750600073ffffffffffffffffffffffffffffffffffffffff8461099683979561093d61097e96606461ffff9a0135907f00000000000000000000000000000000000000000000000000000000000000006132c2565b604051998a98899788957fb1c71c65000000000000000000000000000000000000000000000000000000008752606060048801526064870190600401612e22565b921660248501526003198483030160448501526120b7565b0393165af180156106705760009081906109ce575b6109c492506040519283926040845260408401906121be565b9060208301520390f35b50903d8083833e6109df8183612019565b8101916040828403126106695781519067ffffffffffffffff82116106695750610a106109c4936020928401612dc1565b9101516109ab565b610a2a67ffffffffffffffff916123f9565b7fe86656fb000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b610a7b91925060203d602011610a82575b610a738183612019565b810190612d53565b90866108cb565b503d610a69565b610a9b67ffffffffffffffff916123f9565b7f28c4f25e000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60038103610b4c57505073ffffffffffffffffffffffffffffffffffffffff60065416928315610a8957509161097e600073ffffffffffffffffffffffffffffffffffffffff846109968397610b47606461ffff990135887f00000000000000000000000000000000000000000000000000000000000000006132c2565b61093d565b906103c26024927f31603b120000000000000000000000000000000000000000000000000000000083526121eb565b610ba8907f31603b12000000000000000000000000000000000000000000000000000000006000526121eb565b60246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b610c24915060203d602011610a8257610a738183612019565b856107ab565b346102045760206003193601126102045767ffffffffffffffff610c4c612111565b166000526002602052602060ff60406000205416610c6d6040518092612207565bf35b3461020457602060031936011261020457610c88612111565b600067ffffffffffffffff6024610c9e84613093565b509373ffffffffffffffffffffffffffffffffffffffff60405195869485937fa42a7b8b000000000000000000000000000000000000000000000000000000008552166004840152165afa90811561067057600091610d7c575b506040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210610d3157505050500390f35b91936020610d6c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0600195979984950301865288516120b7565b9601920192018594939192610d22565b3d8083833e610d8b8183612019565b810190602081830312610e6a5780519067ffffffffffffffff8211610e39570181601f82011215610e6a5780519267ffffffffffffffff8411610e3d578360051b906020820194610ddf6040519687612019565b85526020808601928401019284841161066c5760208101925b848410610e0b5750505050505081610cf8565b835167ffffffffffffffff8111610e3957602091610e2e88848094870101612d7f565b815201930192610df8565b8380fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b8280fd5b3461020457600060031936011261020457602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b346102045760206003193601126102045760043567ffffffffffffffff81116102045760a0600319823603011261020457610edb612d3a565b506024810190610eea826123f9565b67ffffffffffffffff604051917fa8d87a3b00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561067057600091611142575b5073ffffffffffffffffffffffffffffffffffffffff33911603610bdd5767ffffffffffffffff610f9e836123f9565b16600052600260205260ff604060002054166005811015610bae57600281036110f2575073ffffffffffffffffffffffffffffffffffffffff60045416905b73ffffffffffffffffffffffffffffffffffffffff82169182156110e0576000928261103361107393606487960135907f00000000000000000000000000000000000000000000000000000000000000006132c2565b6040519485809481937f9a4575b9000000000000000000000000000000000000000000000000000000008352602060048401526024830190600401612e22565b03925af180156106705760009061109d575b610622906040519182916020835260208301906121be565b3d8082843e6110ac8184612019565b82019160208184031261066c5780519167ffffffffffffffff83116106695750916110db916106229301612dc1565b611085565b67ffffffffffffffff610a9b856123f9565b6001810361111a575073ffffffffffffffffffffffffffffffffffffffff6003541690610fdd565b60038103610b7b575073ffffffffffffffffffffffffffffffffffffffff6006541690610fdd565b61115b915060203d602011610a8257610a738183612019565b83610f6e565b3461020457600060031936011261020457602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102045760c0600319360112610204576111ae611f13565b506111b76120fa565b6111bf61214e565b5060843567ffffffffffffffff8111610204576111e090369060040161215f565b60a435916002831015610204576111f693612b4a565b60405180916020820160208352815180915260206040840192019060005b818110611222575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611214565b346102045760206003193601126102045767ffffffffffffffff611273612111565b16600052600260205260ff604060002054166005811015610bae576020906040519015158152f35b346102045760006003193601126102045760005473ffffffffffffffffffffffffffffffffffffffff8116330361133c577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b34610204576000600319360112610204576000606060405161138781611f7a565b828152826020820152826040820152015261062273ffffffffffffffffffffffffffffffffffffffff6003541673ffffffffffffffffffffffffffffffffffffffff6004541673ffffffffffffffffffffffffffffffffffffffff6005541673ffffffffffffffffffffffffffffffffffffffff60065416916040519361140d85611f7a565b845260208401526040830152606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff816080840195828151168552826020820151166020860152826040820151166040860152015116910152565b346102045760206003193601126102045760043567ffffffffffffffff81116102045761149c90369060040161218d565b9073ffffffffffffffffffffffffffffffffffffffff600754169182156115c15760005b8181106114c957005b6114d4818385613038565b3573ffffffffffffffffffffffffffffffffffffffff8116809103610204576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa801561067057869160009161158c575b509081600194939261154e575b505050016114c0565b60208161157d7f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385876132c2565b604051908152a3858581611545565b91506020823d82116115b9575b816115a660209383612019565b8101031261066957505185906001611538565b3d9150611599565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102045760406003193601126102045760043567ffffffffffffffff81116102045761010060031982360301126102045761163460209161162b61213d565b906004016128fd565b60405190518152f35b346102045760206003193601126102045760043567ffffffffffffffff8111610204576101006003198236030112610204576116346020916004016125b9565b3461020457608060031936011261020457611696613048565b6040516116a281611f7a565b6116aa611f13565b908181526024359073ffffffffffffffffffffffffffffffffffffffff82168203610204576020810191825260443573ffffffffffffffffffffffffffffffffffffffff81168103610204576040820190815273ffffffffffffffffffffffffffffffffffffffff61171a611f36565b9460608401958652168015159081611b16575b50611ad15773ffffffffffffffffffffffffffffffffffffffff8351168015159081611ac0575b50611a7b5773ffffffffffffffffffffffffffffffffffffffff8151168015159081611a6a575b50611a255773ffffffffffffffffffffffffffffffffffffffff8451168015159081611a14575b50806119ed575b6119a85773ffffffffffffffffffffffffffffffffffffffff82511630148015611988575b8015611968575b8015611948575b61191e577f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa89373ffffffffffffffffffffffffffffffffffffffff80928161191996818751167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065560405191829182919091606073ffffffffffffffffffffffffffffffffffffffff816080840195828151168552826020820151166020860152826040820151166040860152015116910152565b0390a1005b7fd4bb6daa0000000000000000000000000000000000000000000000000000000060005260046000fd5b503073ffffffffffffffffffffffffffffffffffffffff855116146117dc565b503073ffffffffffffffffffffffffffffffffffffffff825116146117d5565b503073ffffffffffffffffffffffffffffffffffffffff845116146117ce565b73ffffffffffffffffffffffffffffffffffffffff8451167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b50611a0e73ffffffffffffffffffffffffffffffffffffffff8551166133fc565b156117a9565b611a1e91506133a5565b15856117a2565b73ffffffffffffffffffffffffffffffffffffffff9051167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b611a7491506133fc565b158561177b565b73ffffffffffffffffffffffffffffffffffffffff8351167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b611aca91506133a5565b1585611754565b73ffffffffffffffffffffffffffffffffffffffff8251167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b611b2091506133a5565b158561172d565b346102045760c060031936011261020457611b40611f13565b611b486120fa565b611b50611f36565b6084359061ffff821682036102045760a43567ffffffffffffffff81116102045760a09463ffffffff94859461ffff94611b91611b9d95369060040161215f565b94909360443591612280565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b34610204576020600319360112610204576020611bdf611f13565b73ffffffffffffffffffffffffffffffffffffffff604051911673ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148152f35b3461020457600060031936011261020457602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b34610204576000600319360112610204576106226040805190611ca98183612019565b601882527f55534443546f6b656e506f6f6c50726f787920322e302e3000000000000000006020830152519182916020835260208301906120b7565b346102045760206003193601126102045773ffffffffffffffffffffffffffffffffffffffff611d13611f13565b611d1b613048565b167fffffffffffffffffffffffff00000000000000000000000000000000000000006007541617600755600080f35b3461020457600060031936011261020457606060405173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602082015273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166040820152f35b3461020457602060031936011261020457600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361020457817f331710310000000000000000000000000000000000000000000000000000000060209314908115611ee9575b8115611ebf575b8115611e95575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611e8e565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150611e87565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150611e80565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361020457565b6064359073ffffffffffffffffffffffffffffffffffffffff8216820361020457565b359073ffffffffffffffffffffffffffffffffffffffff8216820361020457565b6080810190811067ffffffffffffffff821117611f9657604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117611f9657604052565b6040810190811067ffffffffffffffff821117611f9657604052565b60e0810190811067ffffffffffffffff821117611f9657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117611f9657604052565b67ffffffffffffffff8111611f9657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106120a75750506000910152565b8181015183820152602001612097565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936120f381518092818752878088019101612094565b0116010190565b6024359067ffffffffffffffff8216820361020457565b6004359067ffffffffffffffff8216820361020457565b359067ffffffffffffffff8216820361020457565b6024359061ffff8216820361020457565b6064359061ffff8216820361020457565b9181601f840112156102045782359167ffffffffffffffff8311610204576020838186019501011161020457565b9181601f840112156102045782359167ffffffffffffffff8311610204576020808501948460051b01011161020457565b6121e89160206121d783516040845260408401906120b7565b9201519060208184039101526120b7565b90565b6005811015610bae57600452565b6005811015610bae57602452565b906005821015610bae5752565b519063ffffffff8216820361020457565b519061ffff8216820361020457565b5190811515820361020457565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b93909591949261228f87613093565b9690966122de5767ffffffffffffffff88166000526002602052610ba860ff604060002054167f31603b12000000000000000000000000000000000000000000000000000000006000526121eb565b6040517f2c06340400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff968716600482015267ffffffffffffffff98909816602489015260448801528416606487015261ffff16608486015260c060a486015260a09385939092849283916123689160c4840191612241565b0392165afa908115610670576000918293838093819361238b575b509493929190565b9450925093505060a0823d60a0116123f1575b816123ab60a09383612019565b81010312610669575080516123c260208301612214565b6123ce60408401612214565b936123e760806123e060608701612225565b9501612234565b9194939238612383565b3d915061239e565b3567ffffffffffffffff811681036102045790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610204570180359067ffffffffffffffff82116102045760200191813603831361020457565b90816020910312610204576040519061247782611fc5565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561020457016020813591019167ffffffffffffffff821161020457813603831361020457565b6121e8916125ab6125a06125856124f76124e7868061247d565b6101008752610100870191612241565b67ffffffffffffffff61250c60208801612128565b16602086015273ffffffffffffffffffffffffffffffffffffffff61253360408801611f59565b1660408601526060860135606086015273ffffffffffffffffffffffffffffffffffffffff61256460808801611f59565b16608086015261257760a087018761247d565b9086830360a0880152612241565b61259260c086018661247d565b9085830360c0870152612241565b9260e081019061247d565b9160e0818503910152612241565b906040516125c681611fc5565b6000905261262360206125da8482016123f9565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015233602482015291829081906044820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610670576000916128a1575b5015610bdd5760c08201917fffffffff000000000000000000000000000000000000000000000000000000006126ac6126a6858461240e565b90613217565b16927ffa7c07de00000000000000000000000000000000000000000000000000000000841461283f577fb148ea5f0000000000000000000000000000000000000000000000000000000084146127dd576127086040918361240e565b90501461273d57827fcacdaf2b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b600091925061279e60209173ffffffffffffffffffffffffffffffffffffffff60035416906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835287600484015260248301906124cd565b03925af1908115610670576000916127b4575090565b6121e8915060203d6020116127d6575b6127ce8183612019565b81019061245f565b503d6127c4565b50600091925061279e60209173ffffffffffffffffffffffffffffffffffffffff60045416906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835287600484015260248301906124cd565b50600091925061279e60209173ffffffffffffffffffffffffffffffffffffffff60065416906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835287600484015260248301906124cd565b90506020813d6020116128d3575b816128bc60209383612019565b81010312610204576128cd90612234565b3861266d565b3d91506128af565b9061ffff6128f66020929594956040855260408501906124cd565b9416910152565b919060405161290b81611fc5565b6000905261291f60206125da8582016123f9565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561067057600091612ad4575b5015610bdd577fffffffff000000000000000000000000000000000000000000000000000000006129a06126a660c086018661240e565b16927ffa7c07de000000000000000000000000000000000000000000000000000000008414612a79577f3047587c000000000000000000000000000000000000000000000000000000008414612a1e57837fcacdaf2b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61279e9293509060209173ffffffffffffffffffffffffffffffffffffffff600554169060006040518096819582947f489a68f2000000000000000000000000000000000000000000000000000000008452600484016128db565b61279e9293509060209173ffffffffffffffffffffffffffffffffffffffff600654169060006040518096819582947f489a68f2000000000000000000000000000000000000000000000000000000008452600484016128db565b90506020813d602011612b06575b81612aef60209383612019565b8101031261020457612b0090612234565b38612969565b3d9150612ae2565b805115612b1b5760200190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b67ffffffffffffffff90929192169081600052600260205260ff60406000205416926040947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0865196612b9d8189612019565b60018852013660208801376002811015610bae57600114612c455750506005821015610bae578115612c185750600414612bd45790565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016612c1482612b0e565b5290565b7f28c4f25e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7fffffffff000000000000000000000000000000000000000000000000000000009350612c759250949394613217565b16917f3047587c000000000000000000000000000000000000000000000000000000008314612cf7577ffa7c07de000000000000000000000000000000000000000000000000000000008314612cf357827fcacdaf2b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9150565b90915073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016612c1482612b0e565b60405190612d4782611fe1565b60606020838281520152565b90816020910312610204575173ffffffffffffffffffffffffffffffffffffffff811681036102045790565b81601f82011215610204578051612d958161205a565b92612da36040519485612019565b81845260208284010111610204576121e89160208085019101612094565b91906040838203126102045760405190612dda82611fe1565b8193805167ffffffffffffffff81116102045782612df9918301612d7f565b835260208101519167ffffffffffffffff831161020457602092612e1d9201612d7f565b910152565b90608073ffffffffffffffffffffffffffffffffffffffff612e9882612e59612e4b878061247d565b60a0885260a0880191612241565b9567ffffffffffffffff612e6f60208301612128565b16602087015283612e8260408301611f59565b1660408701526060810135606087015201611f59565b1691015290565b9392919091604051612eb081611ffd565b6000815260006020820152600060408201526000606082015260006080820152600060a0820152600060c082015293612ee884613093565b939093612ef9575050505050905090565b60e0955095612f7673ffffffffffffffffffffffffffffffffffffffff93849361ffff67ffffffffffffffff9a6040519b8c9a8b998a987fd8aa3f40000000000000000000000000000000000000000000000000000000008a52166004890152166024870152166044850152608060648501526084840191612241565b0392165afa90811561067057600091612f8d575090565b905060e0813d60e011613030575b81612fa860e09383612019565b810103126102045761302860c060405192612fc284611ffd565b612fcb81612214565b8452612fd960208201612214565b6020850152612fea60408201612214565b6040850152612ffb60608201612214565b606085015261300c60808201612225565b608085015261301d60a08201612225565b60a085015201612234565b60c082015290565b3d9150612f9b565b9190811015612b1b5760051b0190565b73ffffffffffffffffffffffffffffffffffffffff60015416330361306957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90600067ffffffffffffffff600093169283600052600260205260ff60406000205416916005831015610bae5782156131ea576000600484036131445750505073ffffffffffffffffffffffffffffffffffffffff60055416906001935b73ffffffffffffffffffffffffffffffffffffffff8316156131135750509190565b61042b92507f87d77d33000000000000000000000000000000000000000000000000000000006000526004526121f9565b509093919060028203613172575073ffffffffffffffffffffffffffffffffffffffff60045416915b6130f1565b60006001830361319d57505073ffffffffffffffffffffffffffffffffffffffff60035416916130f1565b50917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd820161316d576006546001955073ffffffffffffffffffffffffffffffffffffffff1692506130f1565b610ba8837f31603b12000000000000000000000000000000000000000000000000000000006000526121eb565b90600481108061324f575060041161020457357fffffffff000000000000000000000000000000000000000000000000000000001690565b907fffffffff00000000000000000000000000000000000000000000000000000000923590838216926132ad575b50507fcacdaf2b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b839250829060040360031b1b1616828061327d565b916020916000916040519073ffffffffffffffffffffffffffffffffffffffff858301937fa9059cbb000000000000000000000000000000000000000000000000000000008552166024830152604482015260448152613323606482612019565b519082855af115610670576000513d61339c575073ffffffffffffffffffffffffffffffffffffffff81163b155b6133585750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415613351565b6133ae81613453565b90816133ea575b816133be575090565b6121e891507f0e64dd290000000000000000000000000000000000000000000000000000000090613548565b90506133f5816134e4565b15906133b5565b61340581613453565b9081613441575b81613415575090565b6121e891507f331710310000000000000000000000000000000000000000000000000000000090613548565b905061344c816134e4565b159061340c565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a7000000000000000000000000000000000000000000000000000000006024820152602481526134b7604482612019565b5191617530fa6000513d826134d8575b50816134d1575090565b9050151590565b602011159150386134c7565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff000000000000000000000000000000000000000000000000000000006024820152602481526134b7604482612019565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a7000000000000000000000000000000000000000000000000000000008452166024820152602481526134b760448261201956fea164736f6c634300081a000a",
}

var USDCTokenPoolProxyABI = USDCTokenPoolProxyMetaData.ABI

var USDCTokenPoolProxyBin = USDCTokenPoolProxyMetaData.Bin

func DeployUSDCTokenPoolProxy(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, pools USDCTokenPoolProxyPoolAddresses, router common.Address, cctpVerifier common.Address) (common.Address, *types.Transaction, *USDCTokenPoolProxy, error) {
	parsed, err := USDCTokenPoolProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDCTokenPoolProxyBin), backend, token, pools, router, cctpVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDCTokenPoolProxy{address: address, abi: *parsed, USDCTokenPoolProxyCaller: USDCTokenPoolProxyCaller{contract: contract}, USDCTokenPoolProxyTransactor: USDCTokenPoolProxyTransactor{contract: contract}, USDCTokenPoolProxyFilterer: USDCTokenPoolProxyFilterer{contract: contract}}, nil
}

type USDCTokenPoolProxy struct {
	address common.Address
	abi     abi.ABI
	USDCTokenPoolProxyCaller
	USDCTokenPoolProxyTransactor
	USDCTokenPoolProxyFilterer
}

type USDCTokenPoolProxyCaller struct {
	contract *bind.BoundContract
}

type USDCTokenPoolProxyTransactor struct {
	contract *bind.BoundContract
}

type USDCTokenPoolProxyFilterer struct {
	contract *bind.BoundContract
}

type USDCTokenPoolProxySession struct {
	Contract     *USDCTokenPoolProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolProxyCallerSession struct {
	Contract *USDCTokenPoolProxyCaller
	CallOpts bind.CallOpts
}

type USDCTokenPoolProxyTransactorSession struct {
	Contract     *USDCTokenPoolProxyTransactor
	TransactOpts bind.TransactOpts
}

type USDCTokenPoolProxyRaw struct {
	Contract *USDCTokenPoolProxy
}

type USDCTokenPoolProxyCallerRaw struct {
	Contract *USDCTokenPoolProxyCaller
}

type USDCTokenPoolProxyTransactorRaw struct {
	Contract *USDCTokenPoolProxyTransactor
}

func NewUSDCTokenPoolProxy(address common.Address, backend bind.ContractBackend) (*USDCTokenPoolProxy, error) {
	abi, err := abi.JSON(strings.NewReader(USDCTokenPoolProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindUSDCTokenPoolProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxy{address: address, abi: abi, USDCTokenPoolProxyCaller: USDCTokenPoolProxyCaller{contract: contract}, USDCTokenPoolProxyTransactor: USDCTokenPoolProxyTransactor{contract: contract}, USDCTokenPoolProxyFilterer: USDCTokenPoolProxyFilterer{contract: contract}}, nil
}

func NewUSDCTokenPoolProxyCaller(address common.Address, caller bind.ContractCaller) (*USDCTokenPoolProxyCaller, error) {
	contract, err := bindUSDCTokenPoolProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyCaller{contract: contract}, nil
}

func NewUSDCTokenPoolProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*USDCTokenPoolProxyTransactor, error) {
	contract, err := bindUSDCTokenPoolProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyTransactor{contract: contract}, nil
}

func NewUSDCTokenPoolProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*USDCTokenPoolProxyFilterer, error) {
	contract, err := bindUSDCTokenPoolProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyFilterer{contract: contract}, nil
}

func bindUSDCTokenPoolProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := USDCTokenPoolProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPoolProxy.Contract.USDCTokenPoolProxyCaller.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.USDCTokenPoolProxyTransactor.contract.Transfer(opts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.USDCTokenPoolProxyTransactor.contract.Transact(opts, method, params...)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDCTokenPoolProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.contract.Transfer(opts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.contract.Transact(opts, method, params...)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetFee(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, amount *big.Int, feeToken common.Address, blockConfirmationsRequested uint16, tokenArgs []byte) (GetFee,

	error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getFee", localToken, destChainSelector, amount, feeToken, blockConfirmationsRequested, tokenArgs)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DestGasOverhead = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.DestBytesOverhead = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TokenFeeBps = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.IsEnabled = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetFee(localToken common.Address, destChainSelector uint64, amount *big.Int, feeToken common.Address, blockConfirmationsRequested uint16, tokenArgs []byte) (GetFee,

	error) {
	return _USDCTokenPoolProxy.Contract.GetFee(&_USDCTokenPoolProxy.CallOpts, localToken, destChainSelector, amount, feeToken, blockConfirmationsRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetFee(localToken common.Address, destChainSelector uint64, amount *big.Int, feeToken common.Address, blockConfirmationsRequested uint16, tokenArgs []byte) (GetFee,

	error) {
	return _USDCTokenPoolProxy.Contract.GetFee(&_USDCTokenPoolProxy.CallOpts, localToken, destChainSelector, amount, feeToken, blockConfirmationsRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetFeeAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getFeeAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetFeeAggregator() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetFeeAggregator(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetFeeAggregator() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetFeeAggregator(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetLockOrBurnMechanism(opts *bind.CallOpts, remoteChainSelector uint64) (uint8, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getLockOrBurnMechanism", remoteChainSelector)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetLockOrBurnMechanism(remoteChainSelector uint64) (uint8, error) {
	return _USDCTokenPoolProxy.Contract.GetLockOrBurnMechanism(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetLockOrBurnMechanism(remoteChainSelector uint64) (uint8, error) {
	return _USDCTokenPoolProxy.Contract.GetLockOrBurnMechanism(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetPools(opts *bind.CallOpts) (USDCTokenPoolProxyPoolAddresses, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getPools")

	if err != nil {
		return *new(USDCTokenPoolProxyPoolAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new(USDCTokenPoolProxyPoolAddresses)).(*USDCTokenPoolProxyPoolAddresses)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetPools() (USDCTokenPoolProxyPoolAddresses, error) {
	return _USDCTokenPoolProxy.Contract.GetPools(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetPools() (USDCTokenPoolProxyPoolAddresses, error) {
	return _USDCTokenPoolProxy.Contract.GetPools(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _USDCTokenPoolProxy.Contract.GetRemotePools(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _USDCTokenPoolProxy.Contract.GetRemotePools(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _USDCTokenPoolProxy.Contract.GetRemoteToken(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _USDCTokenPoolProxy.Contract.GetRemoteToken(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, arg2 *big.Int, arg3 uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, arg2, arg3, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, arg2 *big.Int, arg3 uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRequiredCCVs(&_USDCTokenPoolProxy.CallOpts, arg0, remoteChainSelector, arg2, arg3, extraData, direction)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, arg2 *big.Int, arg3 uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRequiredCCVs(&_USDCTokenPoolProxy.CallOpts, arg0, remoteChainSelector, arg2, arg3, extraData, direction)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetStaticConfig(opts *bind.CallOpts) (GetStaticConfig,

	error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getStaticConfig")

	outstruct := new(GetStaticConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Router = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CctpVerifier = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetStaticConfig() (GetStaticConfig,

	error) {
	return _USDCTokenPoolProxy.Contract.GetStaticConfig(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetStaticConfig() (GetStaticConfig,

	error) {
	return _USDCTokenPoolProxy.Contract.GetStaticConfig(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetToken() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetToken(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetToken() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetToken(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, blockConfirmationsRequested uint16, tokenArgs []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getTokenTransferFeeConfig", localToken, destChainSelector, blockConfirmationsRequested, tokenArgs)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetTokenTransferFeeConfig(localToken common.Address, destChainSelector uint64, blockConfirmationsRequested uint16, tokenArgs []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPoolProxy.Contract.GetTokenTransferFeeConfig(&_USDCTokenPoolProxy.CallOpts, localToken, destChainSelector, blockConfirmationsRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetTokenTransferFeeConfig(localToken common.Address, destChainSelector uint64, blockConfirmationsRequested uint16, tokenArgs []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPoolProxy.Contract.GetTokenTransferFeeConfig(&_USDCTokenPoolProxy.CallOpts, localToken, destChainSelector, blockConfirmationsRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsSupportedChain(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsSupportedChain(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) IsSupportedToken(token common.Address) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsSupportedToken(&_USDCTokenPoolProxy.CallOpts, token)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsSupportedToken(&_USDCTokenPoolProxy.CallOpts, token)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) Owner() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.Owner(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) Owner() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.Owner(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPoolProxy.Contract.SupportsInterface(&_USDCTokenPoolProxy.CallOpts, interfaceId)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _USDCTokenPoolProxy.Contract.SupportsInterface(&_USDCTokenPoolProxy.CallOpts, interfaceId)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) TypeAndVersion() (string, error) {
	return _USDCTokenPoolProxy.Contract.TypeAndVersion(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) TypeAndVersion() (string, error) {
	return _USDCTokenPoolProxy.Contract.TypeAndVersion(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "acceptOwnership")
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.AcceptOwnership(&_USDCTokenPoolProxy.TransactOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.AcceptOwnership(&_USDCTokenPoolProxy.TransactOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.LockOrBurn(&_USDCTokenPoolProxy.TransactOpts, lockOrBurnIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.LockOrBurn(&_USDCTokenPoolProxy.TransactOpts, lockOrBurnIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.LockOrBurn0(&_USDCTokenPoolProxy.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.LockOrBurn0(&_USDCTokenPoolProxy.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationsRequested)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint0(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint0(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) SetFeeAggregator(opts *bind.TransactOpts, feeAggregator common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "setFeeAggregator", feeAggregator)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) SetFeeAggregator(feeAggregator common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetFeeAggregator(&_USDCTokenPoolProxy.TransactOpts, feeAggregator)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) SetFeeAggregator(feeAggregator common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetFeeAggregator(&_USDCTokenPoolProxy.TransactOpts, feeAggregator)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.TransferOwnership(&_USDCTokenPoolProxy.TransactOpts, to)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.TransferOwnership(&_USDCTokenPoolProxy.TransactOpts, to)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) UpdateLockOrBurnMechanisms(opts *bind.TransactOpts, remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "updateLockOrBurnMechanisms", remoteChainSelectors, mechanisms)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) UpdateLockOrBurnMechanisms(remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdateLockOrBurnMechanisms(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, mechanisms)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) UpdateLockOrBurnMechanisms(remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdateLockOrBurnMechanisms(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, mechanisms)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) UpdatePoolAddresses(opts *bind.TransactOpts, pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "updatePoolAddresses", pools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) UpdatePoolAddresses(pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdatePoolAddresses(&_USDCTokenPoolProxy.TransactOpts, pools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) UpdatePoolAddresses(pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdatePoolAddresses(&_USDCTokenPoolProxy.TransactOpts, pools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.WithdrawFeeTokens(&_USDCTokenPoolProxy.TransactOpts, feeTokens)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.WithdrawFeeTokens(&_USDCTokenPoolProxy.TransactOpts, feeTokens)
}

type USDCTokenPoolProxyFeeTokenWithdrawnIterator struct {
	Event *USDCTokenPoolProxyFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyFeeTokenWithdrawn)
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
		it.Event = new(USDCTokenPoolProxyFeeTokenWithdrawn)
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

func (it *USDCTokenPoolProxyFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*USDCTokenPoolProxyFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyFeeTokenWithdrawnIterator{contract: _USDCTokenPoolProxy.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyFeeTokenWithdrawn)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseFeeTokenWithdrawn(log types.Log) (*USDCTokenPoolProxyFeeTokenWithdrawn, error) {
	event := new(USDCTokenPoolProxyFeeTokenWithdrawn)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator struct {
	Event *USDCTokenPoolProxyLockOrBurnMechanismUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyLockOrBurnMechanismUpdated)
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
		it.Event = new(USDCTokenPoolProxyLockOrBurnMechanismUpdated)
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

func (it *USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyLockOrBurnMechanismUpdated struct {
	RemoteChainSelector uint64
	Mechanism           uint8
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterLockOrBurnMechanismUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "LockOrBurnMechanismUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator{contract: _USDCTokenPoolProxy.contract, event: "LockOrBurnMechanismUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchLockOrBurnMechanismUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockOrBurnMechanismUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "LockOrBurnMechanismUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyLockOrBurnMechanismUpdated)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockOrBurnMechanismUpdated", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseLockOrBurnMechanismUpdated(log types.Log) (*USDCTokenPoolProxyLockOrBurnMechanismUpdated, error) {
	event := new(USDCTokenPoolProxyLockOrBurnMechanismUpdated)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockOrBurnMechanismUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyOwnershipTransferRequestedIterator struct {
	Event *USDCTokenPoolProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyOwnershipTransferRequested)
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
		it.Event = new(USDCTokenPoolProxyOwnershipTransferRequested)
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

func (it *USDCTokenPoolProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyOwnershipTransferRequestedIterator{contract: _USDCTokenPoolProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyOwnershipTransferRequested)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolProxyOwnershipTransferRequested, error) {
	event := new(USDCTokenPoolProxyOwnershipTransferRequested)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyOwnershipTransferredIterator struct {
	Event *USDCTokenPoolProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyOwnershipTransferred)
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
		it.Event = new(USDCTokenPoolProxyOwnershipTransferred)
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

func (it *USDCTokenPoolProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyOwnershipTransferredIterator{contract: _USDCTokenPoolProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyOwnershipTransferred)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolProxyOwnershipTransferred, error) {
	event := new(USDCTokenPoolProxyOwnershipTransferred)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyPoolAddressesUpdatedIterator struct {
	Event *USDCTokenPoolProxyPoolAddressesUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyPoolAddressesUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyPoolAddressesUpdated)
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
		it.Event = new(USDCTokenPoolProxyPoolAddressesUpdated)
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

func (it *USDCTokenPoolProxyPoolAddressesUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyPoolAddressesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyPoolAddressesUpdated struct {
	Pools USDCTokenPoolProxyPoolAddresses
	Raw   types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterPoolAddressesUpdated(opts *bind.FilterOpts) (*USDCTokenPoolProxyPoolAddressesUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "PoolAddressesUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyPoolAddressesUpdatedIterator{contract: _USDCTokenPoolProxy.contract, event: "PoolAddressesUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchPoolAddressesUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyPoolAddressesUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "PoolAddressesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyPoolAddressesUpdated)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "PoolAddressesUpdated", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParsePoolAddressesUpdated(log types.Log) (*USDCTokenPoolProxyPoolAddressesUpdated, error) {
	event := new(USDCTokenPoolProxyPoolAddressesUpdated)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "PoolAddressesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}
type GetStaticConfig struct {
	Token        common.Address
	Router       common.Address
	CctpVerifier common.Address
}

func (USDCTokenPoolProxyFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (USDCTokenPoolProxyLockOrBurnMechanismUpdated) Topic() common.Hash {
	return common.HexToHash("0x2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc94471")
}

func (USDCTokenPoolProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (USDCTokenPoolProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (USDCTokenPoolProxyPoolAddressesUpdated) Topic() common.Hash {
	return common.HexToHash("0x67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa8")
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxy) Address() common.Address {
	return _USDCTokenPoolProxy.address
}

type USDCTokenPoolProxyInterface interface {
	GetFee(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, amount *big.Int, feeToken common.Address, blockConfirmationsRequested uint16, tokenArgs []byte) (GetFee,

		error)

	GetFeeAggregator(opts *bind.CallOpts) (common.Address, error)

	GetLockOrBurnMechanism(opts *bind.CallOpts, remoteChainSelector uint64) (uint8, error)

	GetPools(opts *bind.CallOpts) (USDCTokenPoolProxyPoolAddresses, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, arg2 *big.Int, arg3 uint16, extraData []byte, direction uint8) ([]common.Address, error)

	GetStaticConfig(opts *bind.CallOpts) (GetStaticConfig,

		error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, blockConfirmationsRequested uint16, tokenArgs []byte) (IPoolV2TokenTransferFeeConfig, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error)

	SetFeeAggregator(opts *bind.TransactOpts, feeAggregator common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateLockOrBurnMechanisms(opts *bind.TransactOpts, remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error)

	UpdatePoolAddresses(opts *bind.TransactOpts, pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*USDCTokenPoolProxyFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*USDCTokenPoolProxyFeeTokenWithdrawn, error)

	FilterLockOrBurnMechanismUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator, error)

	WatchLockOrBurnMechanismUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockOrBurnMechanismUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockOrBurnMechanismUpdated(log types.Log) (*USDCTokenPoolProxyLockOrBurnMechanismUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolProxyOwnershipTransferred, error)

	FilterPoolAddressesUpdated(opts *bind.FilterOpts) (*USDCTokenPoolProxyPoolAddressesUpdatedIterator, error)

	WatchPoolAddressesUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyPoolAddressesUpdated) (event.Subscription, error)

	ParsePoolAddressesUpdated(log types.Log) (*USDCTokenPoolProxyPoolAddressesUpdated, error)

	Address() common.Address
}
