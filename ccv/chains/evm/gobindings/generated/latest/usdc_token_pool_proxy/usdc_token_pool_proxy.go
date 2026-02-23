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
	DestGasOverhead                        uint32
	DestBytesOverhead                      uint32
	DefaultBlockConfirmationFeeUSDCents    uint32
	CustomBlockConfirmationFeeUSDCents     uint32
	DefaultBlockConfirmationTransferFeeBps uint16
	CustomBlockConfirmationTransferFeeBps  uint16
	IsEnabled                              bool
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeAggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeAggregator\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainNotSupportedByVerifier\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSetPoolForMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"NoLockOrBurnMechanismSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenPoolUnsupported\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e080604052346103a95780613750803803809161001d82856103ae565b833981010360e081126103a9578151916001600160a01b038316918284036103a957608090601f1901126103a957604051608081016001600160401b0381118282101761039357604052610073602083016103d1565b8152610081604083016103d1565b9260208201938452610095606084016103d1565b92604083019384526100a9608082016103d1565b95606084019687526100c960c06100c260a085016103d1565b93016103d1565b92331561038257600180546001600160a01b03191633179055158015610371575b8015610360575b61034f576080526001600160a01b0390811660a05290811660c052815116801515908161033e575b5061031d5782516001600160a01b0316801515908161030c575b506102ea5781516001600160a01b031680151590816102d9575b506102b75783516001600160a01b031680151590816102a6575b508061028b575b6102695751600380546001600160a01b03199081166001600160a01b039384169081179092558451600480548316918516919091179055835160058054831691851691909117905585516006805490921690841617905560408051918252935182166020820152915181169282019290925291511660608201527f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa890608090a1604051613226908161052a823960805181818161090301528181610b0d01528181610ff90152611b3d015260a05181818161076401528181610f270152818161244f015261274a015260c05181818161088401526129cf0152f35b835163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b5083516102a0906001600160a01b0316610426565b1561016e565b6102b091506103e5565b1538610167565b505163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b6102e39150610426565b153861014d565b825163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61031691506103e5565b1538610133565b5163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61034891506103e5565b1538610119565b6303988b8160e61b60005260046000fd5b506001600160a01b038316156100f1565b506001600160a01b038216156100ea565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761039357604052565b51906001600160a01b03821682036103a957565b6103ee81610464565b9081610414575b816103fe575090565b6104119150630e64dd2960e01b906104f5565b90565b905061041f816104c3565b15906103f5565b61042f81610464565b9081610452575b8161043f575090565b6104119150633317103160e01b906104f5565b905061045d816104c3565b1590610436565b6000602091604051838101906301ffc9a760e01b82526301ffc9a760e01b6024820152602481526104966044826103ae565b5191617530fa6000513d826104b7575b50816104b0575090565b9050151590565b602011159150386104a6565b6000602091604051838101906301ffc9a760e01b825263ffffffff60e01b6024820152602481526104966044826103ae565b600090602092604051848101916301ffc9a760e01b835263ffffffff60e01b166024820152602481526104966044826103ae56fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714611c265750806315b358e014611bc1578063181f5a7714611b62578063240028e814611af15780632c06340414611a54578063309292ac146116675780633907753714611627578063489a68f2146115d55780635cb80c5d14611455578063673a2a1f1461135057806379ba5097146112855780638926f54f1461123b57806389720a621461117f5780638da5cb5b1461114b5780639a4575b914610e8c5780639cb406c914610e58578063a42a7b8b14610c59578063aa86a75414610c14578063b1c71c6514610666578063b79465801461056c578063d8aa3f40146104a5578063db4c2aef146101f35763f2fde38b1461011c57600080fd5b346101ee5760206003193601126101ee5773ffffffffffffffffffffffffffffffffffffffff61014a611d26565b610152612d5e565b163381146101c457807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101ee5760406003193601126101ee5760043567ffffffffffffffff81116101ee57610224903690600401611fa0565b9060243567ffffffffffffffff81116101ee57610245903690600401611fa0565b610250929192612d5e565b80840361047b5760005b84811061026357005b61026e818386612d4e565b3560058110156101ee5767ffffffffffffffff61029461028f848988612d4e565b61220c565b16600052600260205260406000206000907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff841617905550600381148061045b575b6103ce576000600182148061043b575b6103635750600281148061041b575b6103ce57600060048214806103ae575b6103635750906001917f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc94471602067ffffffffffffffff61034d61028f868c8b612d4e565b169261035c604051809261201a565ba20161025a565b6103ac60449267ffffffffffffffff61038061028f878c8b612d4e565b7f87d77d330000000000000000000000000000000000000000000000000000000085521660045261200c565bfd5b5073ffffffffffffffffffffffffffffffffffffffff600554161561030a565b67ffffffffffffffff6103e861028f610415948988612d4e565b7f87d77d33000000000000000000000000000000000000000000000000000000006000521660045261200c565b60446000fd5b5073ffffffffffffffffffffffffffffffffffffffff60045416156102fa565b5073ffffffffffffffffffffffffffffffffffffffff60035416156102eb565b5073ffffffffffffffffffffffffffffffffffffffff60065416156102db565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b346101ee5760806003193601126101ee576104be611d26565b6104c6611f0d565b906044359161ffff831683036101ee576064359167ffffffffffffffff83116101ee5760e0936104fd610505943690600401611f72565b939092612bb5565b60c06040519163ffffffff815116835263ffffffff602082015116602084015263ffffffff604082015116604084015263ffffffff606082015116606084015261ffff608082015116608084015261ffff60a08201511660a08401520151151560c0820152f35b346101ee5760206003193601126101ee57610585611f24565b600067ffffffffffffffff602461059b84612da9565b509373ffffffffffffffffffffffffffffffffffffffff60405195869485937fb7946580000000000000000000000000000000000000000000000000000000008552166004840152165afa801561065a57600090610610575b61060c90604051918291602083526020830190611eca565b0390f35b3d8082843e61061f8184611e2c565b8201916020818403126106565780519167ffffffffffffffff831161065357509161064e9161060c9301612a95565b6105f4565b80fd5b5080fd5b6040513d6000823e3d90fd5b346101ee5760606003193601126101ee5760043567ffffffffffffffff81116101ee5760a060031982360301126101ee5761069f611f50565b6044359167ffffffffffffffff83116101ee57366023840112156101ee5782600401356106cb81611e6d565b936106d96040519586611e2c565b81855236602483830101116101ee57816000926024602093018388013785010152610702612a50565b5060248101916107118361220c565b67ffffffffffffffff604051917fa8d87a3b00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561065a57600091610bf5575b5073ffffffffffffffffffffffffffffffffffffffff33911603610bc75767ffffffffffffffff6107c58461220c565b16600052600260205260ff604060002054166005811015610b98578015610b655760009060048103610ab357505073ffffffffffffffffffffffffffffffffffffffff60055416928315610a735761081c8161220c565b9067ffffffffffffffff604051927f958021a7000000000000000000000000000000000000000000000000000000008452166004830152604060248301526020828061086b604482018a611eca565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa91821561065a57600092610a42575b5073ffffffffffffffffffffffffffffffffffffffff821615610a025750600073ffffffffffffffffffffffffffffffffffffffff8461098083979561092761096896606461ffff9a0135907f0000000000000000000000000000000000000000000000000000000000000000612f2d565b604051998a98899788957fb1c71c65000000000000000000000000000000000000000000000000000000008752606060048801526064870190600401612b38565b92166024850152600319848303016044850152611eca565b0393165af1801561065a5760009081906109b8575b6109ae9250604051928392604084526040840190611fd1565b9060208301520390f35b50903d8083833e6109c98183611e2c565b8101916040828403126106535781519067ffffffffffffffff821161065357506109fa6109ae936020928401612ad7565b910151610995565b610a1467ffffffffffffffff9161220c565b7fe86656fb000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b610a6591925060203d602011610a6c575b610a5d8183611e2c565b810190612a69565b90866108b5565b503d610a53565b610a8567ffffffffffffffff9161220c565b7f28c4f25e000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60038103610b3657505073ffffffffffffffffffffffffffffffffffffffff60065416928315610a73575091610968600073ffffffffffffffffffffffffffffffffffffffff846109808397610b31606461ffff990135887f0000000000000000000000000000000000000000000000000000000000000000612f2d565b610927565b906103ac6024927f31603b12000000000000000000000000000000000000000000000000000000008352611ffe565b610b92907f31603b1200000000000000000000000000000000000000000000000000000000600052611ffe565b60246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b610c0e915060203d602011610a6c57610a5d8183611e2c565b85610795565b346101ee5760206003193601126101ee5767ffffffffffffffff610c36611f24565b166000526002602052602060ff60406000205416610c57604051809261201a565bf35b346101ee5760206003193601126101ee57610c72611f24565b600067ffffffffffffffff6024610c8884612da9565b509373ffffffffffffffffffffffffffffffffffffffff60405195869485937fa42a7b8b000000000000000000000000000000000000000000000000000000008552166004840152165afa90811561065a57600091610d66575b506040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210610d1b57505050500390f35b91936020610d56827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851611eca565b9601920192018594939192610d0c565b3d8083833e610d758183611e2c565b810190602081830312610e545780519067ffffffffffffffff8211610e23570181601f82011215610e545780519267ffffffffffffffff8411610e27578360051b906020820194610dc96040519687611e2c565b8552602080860192840101928484116106565760208101925b848410610df55750505050505081610ce2565b835167ffffffffffffffff8111610e2357602091610e1888848094870101612a95565b815201930192610de2565b8380fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526041600452fd5b8280fd5b346101ee5760006003193601126101ee57602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b346101ee5760206003193601126101ee5760043567ffffffffffffffff81116101ee5760a060031982360301126101ee57610ec5612a50565b506024810190610ed48261220c565b67ffffffffffffffff604051917fa8d87a3b00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561065a5760009161112c575b5073ffffffffffffffffffffffffffffffffffffffff33911603610bc75767ffffffffffffffff610f888361220c565b16600052600260205260ff604060002054166005811015610b9857600281036110dc575073ffffffffffffffffffffffffffffffffffffffff60045416905b73ffffffffffffffffffffffffffffffffffffffff82169182156110ca576000928261101d61105d93606487960135907f0000000000000000000000000000000000000000000000000000000000000000612f2d565b6040519485809481937f9a4575b9000000000000000000000000000000000000000000000000000000008352602060048401526024830190600401612b38565b03925af1801561065a57600090611087575b61060c90604051918291602083526020830190611fd1565b3d8082843e6110968184611e2c565b8201916020818403126106565780519167ffffffffffffffff83116106535750916110c59161060c9301612ad7565b61106f565b67ffffffffffffffff610a858561220c565b60018103611104575073ffffffffffffffffffffffffffffffffffffffff6003541690610fc7565b60038103610b65575073ffffffffffffffffffffffffffffffffffffffff6006541690610fc7565b611145915060203d602011610a6c57610a5d8183611e2c565b83610f58565b346101ee5760006003193601126101ee57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101ee5760c06003193601126101ee57611198611d26565b506111a1611f0d565b6111a9611f61565b5060843567ffffffffffffffff81116101ee576111ca903690600401611f72565b5050600260a43510156101ee576111e090612925565b60405180916020820160208352815180915260206040840192019060005b81811061120c575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016111fe565b346101ee5760206003193601126101ee5767ffffffffffffffff61125d611f24565b16600052600260205260ff604060002054166005811015610b98576020906040519015158152f35b346101ee5760006003193601126101ee5760005473ffffffffffffffffffffffffffffffffffffffff81163303611326577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101ee5760006003193601126101ee576000606060405161137181611d8d565b828152826020820152826040820152015261060c73ffffffffffffffffffffffffffffffffffffffff6003541673ffffffffffffffffffffffffffffffffffffffff6004541673ffffffffffffffffffffffffffffffffffffffff6005541673ffffffffffffffffffffffffffffffffffffffff6006541691604051936113f785611d8d565b845260208401526040830152606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff816080840195828151168552826020820151166020860152826040820151166040860152015116910152565b346101ee5760206003193601126101ee5760043567ffffffffffffffff81116101ee57611486903690600401611fa0565b9073ffffffffffffffffffffffffffffffffffffffff600754169182156115ab5760005b8181106114b357005b6114be818385612d4e565b3573ffffffffffffffffffffffffffffffffffffffff81168091036101ee576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa801561065a578691600091611576575b5090816001949392611538575b505050016114aa565b6020816115677f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e938587612f2d565b604051908152a385858161152f565b91506020823d82116115a3575b8161159060209383611e2c565b8101031261065357505185906001611522565b3d9150611583565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101ee5760406003193601126101ee5760043567ffffffffffffffff81116101ee5761010060031982360301126101ee5761161e602091611615611f50565b9060040161270f565b60405190518152f35b346101ee5760206003193601126101ee5760043567ffffffffffffffff81116101ee5761010060031982360301126101ee5761161e6020916004016123cc565b346101ee5760806003193601126101ee57611680612d5e565b60405161168c81611d8d565b611694611d26565b80825260243573ffffffffffffffffffffffffffffffffffffffff811681036101ee57602083019081526044359073ffffffffffffffffffffffffffffffffffffffff821682036101ee576040840191825273ffffffffffffffffffffffffffffffffffffffff611703611d49565b9360608601948552168015159081611a43575b506119fe5773ffffffffffffffffffffffffffffffffffffffff81511680151590816119ed575b506119a85773ffffffffffffffffffffffffffffffffffffffff8251168015159081611997575b506119525773ffffffffffffffffffffffffffffffffffffffff8351168015159081611941575b508061191a575b6118d5576118d09273ffffffffffffffffffffffffffffffffffffffff8593818094817f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa89951167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065560405191829182919091606073ffffffffffffffffffffffffffffffffffffffff816080840195828151168552826020820151166020860152826040820151166040860152015116910152565b0390a1005b73ffffffffffffffffffffffffffffffffffffffff8351167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b5061193b73ffffffffffffffffffffffffffffffffffffffff845116613067565b15611792565b61194b9150613010565b158561178b565b73ffffffffffffffffffffffffffffffffffffffff8251167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6119a19150613067565b1585611764565b73ffffffffffffffffffffffffffffffffffffffff9051167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6119f79150613010565b158561173d565b73ffffffffffffffffffffffffffffffffffffffff8451167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b611a4d9150613010565b1585611716565b346101ee5760c06003193601126101ee57611a6d611d26565b611a75611f0d565b611a7d611d49565b6084359061ffff821682036101ee5760a43567ffffffffffffffff81116101ee5760a09463ffffffff94859461ffff94611abe611aca953690600401611f72565b94909360443591612093565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b346101ee5760206003193601126101ee576020611b0c611d26565b73ffffffffffffffffffffffffffffffffffffffff604051911673ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148152f35b346101ee5760006003193601126101ee5761060c6040805190611b858183611e2c565b601c82527f55534443546f6b656e506f6f6c50726f787920322e302e302d64657600000000602083015251918291602083526020830190611eca565b346101ee5760206003193601126101ee5773ffffffffffffffffffffffffffffffffffffffff611bef611d26565b611bf7612d5e565b167fffffffffffffffffffffffff00000000000000000000000000000000000000006007541617600755600080f35b346101ee5760206003193601126101ee57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101ee57817f331710310000000000000000000000000000000000000000000000000000000060209314908115611cfc575b8115611cd2575b8115611ca8575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611ca1565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150611c9a565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150611c93565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101ee57565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036101ee57565b359073ffffffffffffffffffffffffffffffffffffffff821682036101ee57565b6080810190811067ffffffffffffffff821117611da957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117611da957604052565b6040810190811067ffffffffffffffff821117611da957604052565b60e0810190811067ffffffffffffffff821117611da957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117611da957604052565b67ffffffffffffffff8111611da957601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b838110611eba5750506000910152565b8181015183820152602001611eaa565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093611f0681518092818752878088019101611ea7565b0116010190565b6024359067ffffffffffffffff821682036101ee57565b6004359067ffffffffffffffff821682036101ee57565b359067ffffffffffffffff821682036101ee57565b6024359061ffff821682036101ee57565b6064359061ffff821682036101ee57565b9181601f840112156101ee5782359167ffffffffffffffff83116101ee57602083818601950101116101ee57565b9181601f840112156101ee5782359167ffffffffffffffff83116101ee576020808501948460051b0101116101ee57565b611ffb916020611fea8351604084526040840190611eca565b920151906020818403910152611eca565b90565b6005811015610b9857600452565b6005811015610b9857602452565b906005821015610b985752565b519063ffffffff821682036101ee57565b519061ffff821682036101ee57565b519081151582036101ee57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9390959194926120a287612da9565b9690966120f15767ffffffffffffffff88166000526002602052610b9260ff604060002054167f31603b1200000000000000000000000000000000000000000000000000000000600052611ffe565b6040517f2c06340400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff968716600482015267ffffffffffffffff98909816602489015260448801528416606487015261ffff16608486015260c060a486015260a093859390928492839161217b9160c4840191612054565b0392165afa90811561065a576000918293838093819361219e575b509493929190565b9450925093505060a0823d60a011612204575b816121be60a09383611e2c565b81010312610653575080516121d560208301612027565b6121e160408401612027565b936121fa60806121f360608701612038565b9501612047565b9194939238612196565b3d91506121b1565b3567ffffffffffffffff811681036101ee5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101ee570180359067ffffffffffffffff82116101ee576020019181360383136101ee57565b908160209103126101ee576040519061228a82611dd8565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101ee57016020813591019167ffffffffffffffff82116101ee5781360383136101ee57565b611ffb916123be6123b361239861230a6122fa8680612290565b6101008752610100870191612054565b67ffffffffffffffff61231f60208801611f3b565b16602086015273ffffffffffffffffffffffffffffffffffffffff61234660408801611d6c565b1660408601526060860135606086015273ffffffffffffffffffffffffffffffffffffffff61237760808801611d6c565b16608086015261238a60a0870187612290565b9086830360a0880152612054565b6123a560c0860186612290565b9085830360c0870152612054565b9260e0810190612290565b9160e0818503910152612054565b906040516123d981611dd8565b6000905261243660206123ed84820161220c565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015233602482015291829081906044820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561065a576000916126b3575b5015610bc75760c08201916124958382612221565b6004116101ee57357fffffffff0000000000000000000000000000000000000000000000000000000016927ffa7c07de000000000000000000000000000000000000000000000000000000008414612651577fb148ea5f0000000000000000000000000000000000000000000000000000000084146125ef5761251a60409183612221565b90501461254f57827fcacdaf2b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60009192506125b060209173ffffffffffffffffffffffffffffffffffffffff60035416906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835287600484015260248301906122e0565b03925af190811561065a576000916125c6575090565b611ffb915060203d6020116125e8575b6125e08183611e2c565b810190612272565b503d6125d6565b5060009192506125b060209173ffffffffffffffffffffffffffffffffffffffff60045416906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835287600484015260248301906122e0565b5060009192506125b060209173ffffffffffffffffffffffffffffffffffffffff60065416906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835287600484015260248301906122e0565b90506020813d6020116126e5575b816126ce60209383611e2c565b810103126101ee576126df90612047565b38612480565b3d91506126c1565b9061ffff6127086020929594956040855260408501906122e0565b9416910152565b919060405161271d81611dd8565b6000905261273160206123ed85820161220c565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561065a576000916128eb575b5015610bc75761278e60c0840184612221565b6004116101ee57357fffffffff0000000000000000000000000000000000000000000000000000000016927ffa7c07de000000000000000000000000000000000000000000000000000000008414612890577f3047587c00000000000000000000000000000000000000000000000000000000841461283557837fcacdaf2b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6125b09293509060209173ffffffffffffffffffffffffffffffffffffffff600554169060006040518096819582947f489a68f2000000000000000000000000000000000000000000000000000000008452600484016126ed565b6125b09293509060209173ffffffffffffffffffffffffffffffffffffffff600654169060006040518096819582947f489a68f2000000000000000000000000000000000000000000000000000000008452600484016126ed565b90506020813d60201161291d575b8161290660209383611e2c565b810103126101ee5761291790612047565b3861277b565b3d91506128f9565b67ffffffffffffffff1680600052600260205260ff604060002054166005811015610b985715612a235760409060ff8251926129618185611e2c565b6001845260208401927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08201368537600052600260205260002054166005811015610b98576004146129b1575090565b8151156129f45773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f28c4f25e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60405190612a5d82611df4565b60606020838281520152565b908160209103126101ee575173ffffffffffffffffffffffffffffffffffffffff811681036101ee5790565b81601f820112156101ee578051612aab81611e6d565b92612ab96040519485611e2c565b818452602082840101116101ee57611ffb9160208085019101611ea7565b91906040838203126101ee5760405190612af082611df4565b8193805167ffffffffffffffff81116101ee5782612b0f918301612a95565b835260208101519167ffffffffffffffff83116101ee57602092612b339201612a95565b910152565b90608073ffffffffffffffffffffffffffffffffffffffff612bae82612b6f612b618780612290565b60a0885260a0880191612054565b9567ffffffffffffffff612b8560208301611f3b565b16602087015283612b9860408301611d6c565b1660408701526060810135606087015201611d6c565b1691015290565b9392919091604051612bc681611e10565b6000815260006020820152600060408201526000606082015260006080820152600060a0820152600060c082015293612bfe84612da9565b939093612c0f575050505050905090565b60e0955095612c8c73ffffffffffffffffffffffffffffffffffffffff93849361ffff67ffffffffffffffff9a6040519b8c9a8b998a987fd8aa3f40000000000000000000000000000000000000000000000000000000008a52166004890152166024870152166044850152608060648501526084840191612054565b0392165afa90811561065a57600091612ca3575090565b905060e0813d60e011612d46575b81612cbe60e09383611e2c565b810103126101ee57612d3e60c060405192612cd884611e10565b612ce181612027565b8452612cef60208201612027565b6020850152612d0060408201612027565b6040850152612d1160608201612027565b6060850152612d2260808201612038565b6080850152612d3360a08201612038565b60a085015201612047565b60c082015290565b3d9150612cb1565b91908110156129f45760051b0190565b73ffffffffffffffffffffffffffffffffffffffff600154163303612d7f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90600067ffffffffffffffff600093169283600052600260205260ff60406000205416916005831015610b98578215612f0057600060048403612e5a5750505073ffffffffffffffffffffffffffffffffffffffff60055416906001935b73ffffffffffffffffffffffffffffffffffffffff831615612e295750509190565b61041592507f87d77d330000000000000000000000000000000000000000000000000000000060005260045261200c565b509093919060028203612e88575073ffffffffffffffffffffffffffffffffffffffff60045416915b612e07565b600060018303612eb357505073ffffffffffffffffffffffffffffffffffffffff6003541691612e07565b50917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201612e83576006546001955073ffffffffffffffffffffffffffffffffffffffff169250612e07565b610b92837f31603b1200000000000000000000000000000000000000000000000000000000600052611ffe565b916020916000916040519073ffffffffffffffffffffffffffffffffffffffff858301937fa9059cbb000000000000000000000000000000000000000000000000000000008552166024830152604482015260448152612f8e606482611e2c565b519082855af11561065a576000513d613007575073ffffffffffffffffffffffffffffffffffffffff81163b155b612fc35750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415612fbc565b613019816130be565b9081613055575b81613029575090565b611ffb91507f0e64dd2900000000000000000000000000000000000000000000000000000000906131b3565b90506130608161314f565b1590613020565b613070816130be565b90816130ac575b81613080575090565b611ffb91507f3317103100000000000000000000000000000000000000000000000000000000906131b3565b90506130b78161314f565b1590613077565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a700000000000000000000000000000000000000000000000000000000602482015260248152613122604482611e2c565b5191617530fa6000513d82613143575b508161313c575090565b9050151590565b60201115915038613132565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff00000000000000000000000000000000000000000000000000000000602482015260248152613122604482611e2c565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a700000000000000000000000000000000000000000000000000000000845216602482015260248152613122604482611e2c56fea164736f6c634300081a000a",
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetFee(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, amount *big.Int, feeToken common.Address, blockConfirmationRequested uint16, tokenArgs []byte) (GetFee,

	error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getFee", localToken, destChainSelector, amount, feeToken, blockConfirmationRequested, tokenArgs)

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

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetFee(localToken common.Address, destChainSelector uint64, amount *big.Int, feeToken common.Address, blockConfirmationRequested uint16, tokenArgs []byte) (GetFee,

	error) {
	return _USDCTokenPoolProxy.Contract.GetFee(&_USDCTokenPoolProxy.CallOpts, localToken, destChainSelector, amount, feeToken, blockConfirmationRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetFee(localToken common.Address, destChainSelector uint64, amount *big.Int, feeToken common.Address, blockConfirmationRequested uint16, tokenArgs []byte) (GetFee,

	error) {
	return _USDCTokenPoolProxy.Contract.GetFee(&_USDCTokenPoolProxy.CallOpts, localToken, destChainSelector, amount, feeToken, blockConfirmationRequested, tokenArgs)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte, arg5 uint8) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte, arg5 uint8) ([]common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRequiredCCVs(&_USDCTokenPoolProxy.CallOpts, arg0, remoteChainSelector, arg2, arg3, arg4, arg5)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte, arg5 uint8) ([]common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRequiredCCVs(&_USDCTokenPoolProxy.CallOpts, arg0, remoteChainSelector, arg2, arg3, arg4, arg5)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, blockConfirmationRequested uint16, tokenArgs []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getTokenTransferFeeConfig", localToken, destChainSelector, blockConfirmationRequested, tokenArgs)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetTokenTransferFeeConfig(localToken common.Address, destChainSelector uint64, blockConfirmationRequested uint16, tokenArgs []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPoolProxy.Contract.GetTokenTransferFeeConfig(&_USDCTokenPoolProxy.CallOpts, localToken, destChainSelector, blockConfirmationRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetTokenTransferFeeConfig(localToken common.Address, destChainSelector uint64, blockConfirmationRequested uint16, tokenArgs []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _USDCTokenPoolProxy.Contract.GetTokenTransferFeeConfig(&_USDCTokenPoolProxy.CallOpts, localToken, destChainSelector, blockConfirmationRequested, tokenArgs)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.LockOrBurn0(&_USDCTokenPoolProxy.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.LockOrBurn0(&_USDCTokenPoolProxy.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint0(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint0(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
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
	GetFee(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, amount *big.Int, feeToken common.Address, blockConfirmationRequested uint16, tokenArgs []byte) (GetFee,

		error)

	GetFeeAggregator(opts *bind.CallOpts) (common.Address, error)

	GetLockOrBurnMechanism(opts *bind.CallOpts, remoteChainSelector uint64) (uint8, error)

	GetPools(opts *bind.CallOpts) (USDCTokenPoolProxyPoolAddresses, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte, arg5 uint8) ([]common.Address, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, localToken common.Address, destChainSelector uint64, blockConfirmationRequested uint16, tokenArgs []byte) (IPoolV2TokenTransferFeeConfig, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

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
