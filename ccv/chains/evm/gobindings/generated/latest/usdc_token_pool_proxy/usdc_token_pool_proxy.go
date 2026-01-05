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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeAggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeAggregator\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CCVCompatiblePoolNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainNotSupportedByVerifier\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSetPoolForMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"NoLockOrBurnMechanismSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenPoolUnsupported\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e080604052346103d05780613b5c803803809161001d82856103d5565b833981010360e081126103d0578151916001600160a01b038316918284036103d057608090601f1901126103d057604051608081016001600160401b038111828210176103ba57604052610073602083016103f8565b8152610081604083016103f8565b9260208201938452610095606084016103f8565b92604083019384526100a9608082016103f8565b95606084019687526100c960c06100c260a085016103f8565b93016103f8565b9233156103a957600180546001600160a01b03191633179055158015610398575b8015610387575b610376576080526001600160a01b0390811660a05290811660c0528151168015159081610365575b506103445782516001600160a01b03168015159081610333575b506103115781516001600160a01b031680151590816102c6575b506102a45783516001600160a01b03168015159081610293575b506102715751600380546001600160a01b03199081166001600160a01b039384169081179092558451600480548316918516919091179055835160058054831691851691909117905585516006805490921690841617905560408051918252935182166020820152915181169282019290925291511660608201527f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa890608090a16040516136499081610513823960805181818161151d01528181612b6b01528181612dd90152818161310c0152613230015260a051818181611d0f0152818161247f01528181612a84015261303a015260c05181818161283201528181612d7301526131c80152f35b835163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61029d915061040c565b1538610167565b505163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b90506102d18161044d565b90816102ff575b816102e6575b50153861014d565b6102f99150633317103160e01b906104de565b386102de565b905061030a816104ac565b15906102d8565b825163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61033d915061040c565b1538610133565b5163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61036f915061040c565b1538610119565b6303988b8160e61b60005260046000fd5b506001600160a01b038316156100f1565b506001600160a01b038216156100ea565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b601f909101601f19168101906001600160401b038211908210176103ba57604052565b51906001600160a01b03821682036103d057565b6104158161044d565b908161043b575b81610425575090565b6104389150630e64dd2960e01b906104de565b90565b9050610446816104ac565b159061041c565b6000602091604051838101906301ffc9a760e01b82526301ffc9a760e01b60248201526024815261047f6044826103d5565b5191617530fa6000513d826104a0575b5081610499575090565b9050151590565b6020111591503861048f565b6000602091604051838101906301ffc9a760e01b825263ffffffff60e01b60248201526024815261047f6044826103d5565b600090602092604051848101916301ffc9a760e01b835263ffffffff60e01b1660248201526024815261047f6044826103d556fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146116065750806315b358e0146115a1578063181f5a7714611542578063240028e8146114d15780632c06340414611305578063309292ac14610ef25780633907753714610ea3578063489a68f214610e515780635cb80c5d14610cd1578063673a2a1f14610bcc57806379ba509714610b015780638926f54f14610a8857806389720a62146109ae5780638da5cb5b1461097a5780639a4575b91461090a5780639cb406c9146108d6578063aa86a75414610891578063b1c71c6514610808578063b79465801461071e578063d8aa3f401461049a578063db4c2aef146101e85763f2fde38b1461011157600080fd5b346101e35760206003193601126101e35773ffffffffffffffffffffffffffffffffffffffff61013f611706565b61014761335c565b163381146101b957807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101e35760406003193601126101e35760043567ffffffffffffffff81116101e35761021990369060040161199d565b9060243567ffffffffffffffff81116101e35761023a90369060040161199d565b61024592919261335c565b8084036104705760005b84811061025857005b61026381838661334c565b3560058110156101e35767ffffffffffffffff61028961028484898861334c565b611ad7565b16600052600260205260406000206000907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff8416179055506003811480610450575b6103c35760006001821480610430575b61035857506002811480610410575b6103c357600060048214806103a3575b6103585750906001917f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc94471602067ffffffffffffffff610342610284868c8b61334c565b16926103516040518092611a17565ba20161024f565b6103a160449267ffffffffffffffff610375610284878c8b61334c565b7f87d77d33000000000000000000000000000000000000000000000000000000008552166004526119fb565bfd5b5073ffffffffffffffffffffffffffffffffffffffff60055416156102ff565b67ffffffffffffffff6103dd61028461040a94898861334c565b7f87d77d3300000000000000000000000000000000000000000000000000000000600052166004526119fb565b60446000fd5b5073ffffffffffffffffffffffffffffffffffffffff60045416156102ef565b5073ffffffffffffffffffffffffffffffffffffffff60035416156102e0565b5073ffffffffffffffffffffffffffffffffffffffff60065416156102d0565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b346101e35760806003193601126101e3576104b3611706565b6104bb61190a565b906044359161ffff83168093036101e35760643567ffffffffffffffff81116101e3576104ec90369060040161196f565b600060c06040959395516104ff816117d4565b8281528260208201528260408201528260608201528260808201528260a0820152015273ffffffffffffffffffffffffffffffffffffffff600554169182156106f4576105b760e09573ffffffffffffffffffffffffffffffffffffffff9367ffffffffffffffff98604051998a98899788977fd8aa3f400000000000000000000000000000000000000000000000000000000089521660048801521660248601526044850152608060648501526084840191611a98565b03915afa9081156106e857600091610634575b60e08260c06040519163ffffffff815116835263ffffffff602082015116602084015263ffffffff604082015116604084015263ffffffff606082015116606084015261ffff608082015116608084015261ffff60a08201511660a08401520151151560c0820152f35b60e0813d60e0116106e0575b8161064d60e09383611829565b810103126106dc5760e091506106d160c06040519261066b846117d4565b61067481611a6b565b845261068260208201611a6b565b602085015261069360408201611a6b565b60408501526106a460608201611a6b565b60608501526106b560808201611a7c565b60808501526106c660a08201611a7c565b60a085015201611a8b565b60c0820152906105ca565b5080fd5b3d9150610640565b6040513d6000823e3d90fd5b7f0f480b320000000000000000000000000000000000000000000000000000000060005260046000fd5b346101e35760206003193601126101e357610737611921565b73ffffffffffffffffffffffffffffffffffffffff600554169081156106f45767ffffffffffffffff602460009260405194859384927fb79465800000000000000000000000000000000000000000000000000000000084521660048301525afa80156106e8576000906107c2575b6107be906040519182916020835260208301906118c7565b0390f35b3d8082843e6107d18184611829565b8201916020818403126106dc5780519167ffffffffffffffff8311610805575091610800916107be93016128f8565b6107a6565b80fd5b346101e35760606003193601126101e35760043567ffffffffffffffff81116101e35760a060031982360301126101e35761084161194d565b906044359067ffffffffffffffff82116101e3576108879261086a610873933690600401611a24565b91600401612fcd565b6040519283926040845260408401906119ce565b9060208301520390f35b346101e35760206003193601126101e35767ffffffffffffffff6108b3611921565b166000526002602052602060ff604060002054166108d46040518092611a17565bf35b346101e35760006003193601126101e357602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b346101e35760206003193601126101e35760043567ffffffffffffffff81116101e35760a060031982360301126101e3576109436128b3565b50610964602091604051906109588483611829565b60008252600401612a18565b50906107be6040519282849384528301906119ce565b346101e35760006003193601126101e357602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101e35760c06003193601126101e3576109c7611706565b506109d061190a565b6109d861195e565b5060843567ffffffffffffffff81116101e3576109f990369060040161196f565b5050600260a43510156101e35773ffffffffffffffffffffffffffffffffffffffff60055416156106f457610a2d90612788565b60405180916020820160208352815180915260206040840192019060005b818110610a59575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610a4b565b346101e35760206003193601126101e35767ffffffffffffffff610aaa611921565b16600052600260205260ff604060002054166005811015610ad2576020906040519015158152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b346101e35760006003193601126101e35760005473ffffffffffffffffffffffffffffffffffffffff81163303610ba2577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101e35760006003193601126101e35760006060604051610bed8161176d565b82815282602082015282604082015201526107be73ffffffffffffffffffffffffffffffffffffffff6003541673ffffffffffffffffffffffffffffffffffffffff6004541673ffffffffffffffffffffffffffffffffffffffff6005541673ffffffffffffffffffffffffffffffffffffffff600654169160405193610c738561176d565b845260208401526040830152606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff816080840195828151168552826020820151166020860152826040820151166040860152015116910152565b346101e35760206003193601126101e35760043567ffffffffffffffff81116101e357610d0290369060040161199d565b9073ffffffffffffffffffffffffffffffffffffffff60075416918215610e275760005b818110610d2f57005b610d3a81838561334c565b3573ffffffffffffffffffffffffffffffffffffffff81168091036101e3576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa80156106e8578691600091610df2575b5090816001949392610db4575b50505001610d26565b602081610de37f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9385876133a7565b604051908152a3858581610dab565b91506020823d8211610e1f575b81610e0c60209383611829565b8101031261080557505185906001610d9e565b3d9150610dff565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101e35760406003193601126101e35760043567ffffffffffffffff81116101e35761010060031982360301126101e357610e9a602091610e9161194d565b90600401612405565b60405190518152f35b346101e35760206003193601126101e35760043567ffffffffffffffff81116101e35761010060031982360301126101e357610e9a6020916000604051610ee9816117b8565b52600401611c97565b346101e35760806003193601126101e357610f0b61335c565b604051610f178161176d565b610f1f611706565b80825260243573ffffffffffffffffffffffffffffffffffffffff811681036101e357602083019081526044359073ffffffffffffffffffffffffffffffffffffffff821682036101e3576040840191825273ffffffffffffffffffffffffffffffffffffffff610f8e611729565b93606086019485521680151590816112f4575b506112af5773ffffffffffffffffffffffffffffffffffffffff815116801515908161129e575b506112595773ffffffffffffffffffffffffffffffffffffffff82511680151590816111f5575b506111b05773ffffffffffffffffffffffffffffffffffffffff835116801515908161119f575b5061115a576111559273ffffffffffffffffffffffffffffffffffffffff8593818094817f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa89951167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065560405191829182919091606073ffffffffffffffffffffffffffffffffffffffff816080840195828151168552826020820151166020860152826040820151166040860152015116910152565b0390a1005b73ffffffffffffffffffffffffffffffffffffffff8351167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6111a9915061348a565b1585611016565b73ffffffffffffffffffffffffffffffffffffffff8251167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9050611200816134e1565b9081611247575b81611215575b501585610fef565b61124191507f3317103100000000000000000000000000000000000000000000000000000000906135d6565b8561120d565b905061125281613572565b1590611207565b73ffffffffffffffffffffffffffffffffffffffff9051167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6112a8915061348a565b1585610fc8565b73ffffffffffffffffffffffffffffffffffffffff8451167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6112fe915061348a565b1585610fa1565b346101e35760c06003193601126101e35761131e611706565b61132661190a565b61132e611729565b906084359161ffff83168093036101e35760a43567ffffffffffffffff81116101e35761135f90369060040161196f565b93909473ffffffffffffffffffffffffffffffffffffffff600554169283156106f45767ffffffffffffffff9660a09673ffffffffffffffffffffffffffffffffffffffff9485611407946040519b8c9a8b998a997f2c063404000000000000000000000000000000000000000000000000000000008b521660048a01521660248801526044356044880152166064860152608485015260c060a485015260c4840191611a98565b03915afa80156106e85760009081829083948493611450575b509363ffffffff61ffff928160a09760405197885216602087015216604085015216606083015215156080820152f35b93945050505060a0813d60a0116114c9575b8161146f60a09383611829565b810103126106dc5760a0915080519061ffff61148d60208301611a6b565b9263ffffffff61149f60408501611a6b565b93816114b960806114b260608501611a7c565b9301611a8b565b9197509295909493509150611420565b3d9150611462565b346101e35760206003193601126101e35760206114ec611706565b73ffffffffffffffffffffffffffffffffffffffff604051911673ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148152f35b346101e35760006003193601126101e3576107be60408051906115658183611829565b601c82527f55534443546f6b656e506f6f6c50726f787920312e372e302d646576000000006020830152519182916020835260208301906118c7565b346101e35760206003193601126101e35773ffffffffffffffffffffffffffffffffffffffff6115cf611706565b6115d761335c565b167fffffffffffffffffffffffff00000000000000000000000000000000000000006007541617600755600080f35b346101e35760206003193601126101e357600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101e357817f3317103100000000000000000000000000000000000000000000000000000000602093149081156116dc575b81156116b2575b8115611688575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483611681565b7faff2afbf000000000000000000000000000000000000000000000000000000008114915061167a565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150611673565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101e357565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036101e357565b359073ffffffffffffffffffffffffffffffffffffffff821682036101e357565b6080810190811067ffffffffffffffff82111761178957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761178957604052565b60e0810190811067ffffffffffffffff82111761178957604052565b610100810190811067ffffffffffffffff82111761178957604052565b6040810190811067ffffffffffffffff82111761178957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761178957604052565b67ffffffffffffffff811161178957601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106118b75750506000910152565b81810151838201526020016118a7565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093611903815180928187528780880191016118a4565b0116010190565b6024359067ffffffffffffffff821682036101e357565b6004359067ffffffffffffffff821682036101e357565b359067ffffffffffffffff821682036101e357565b6024359061ffff821682036101e357565b6064359061ffff821682036101e357565b9181601f840112156101e35782359167ffffffffffffffff83116101e357602083818601950101116101e357565b9181601f840112156101e35782359167ffffffffffffffff83116101e3576020808501948460051b0101116101e357565b6119f89160206119e783516040845260408401906118c7565b9201519060208184039101526118c7565b90565b6005811015610ad257602452565b6005811015610ad257600452565b906005821015610ad25752565b81601f820112156101e357803590611a3b8261186a565b92611a496040519485611829565b828452602083830101116101e357816000926020809301838601378301015290565b519063ffffffff821682036101e357565b519061ffff821682036101e357565b519081151582036101e357565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b3567ffffffffffffffff811681036101e35790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101e3570180359067ffffffffffffffff82116101e3576020019181360383136101e357565b908160209103126101e35760405190611b55826117b8565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101e357016020813591019167ffffffffffffffff82116101e35781360383136101e357565b6119f891611c89611c7e611c63611bd5611bc58680611b5b565b6101008752610100870191611a98565b67ffffffffffffffff611bea60208801611938565b16602086015273ffffffffffffffffffffffffffffffffffffffff611c116040880161174c565b1660408601526060860135606086015273ffffffffffffffffffffffffffffffffffffffff611c426080880161174c565b166080860152611c5560a0870187611b5b565b9086830360a0880152611a98565b611c7060c0860186611b5b565b9085830360c0870152611a98565b9260e0810190611b5b565b9160e0818503910152611a98565b90604051611ca4816117b8565b600090526020820191611cb683611ad7565b67ffffffffffffffff604051917f83826b2b00000000000000000000000000000000000000000000000000000000835216600482015233602482015260208160448173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156106e8576000916123cb575b501561239d5760c0810190611d558282611aec565b6004116101e357357fffffffff0000000000000000000000000000000000000000000000000000000016937ffa7c07de000000000000000000000000000000000000000000000000000000008514612339577ff3567d180000000000000000000000000000000000000000000000000000000085146122d5577fb148ea5f00000000000000000000000000000000000000000000000000000000851461225b577f3047587c0000000000000000000000000000000000000000000000000000000085146121da576040611e288484611aec565b905014611e5d57847fcacdaf2b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9091929350610100823603126101e35760405190611e7a826117f0565b82359067ffffffffffffffff82116101e357611e9c611ea39236908601611a24565b8352611938565b9060208101918252611eb76040840161174c565b9060408101918252606081019360608101358552611ed76080820161174c565b936080830194855260a082013567ffffffffffffffff81116101e357611f009036908401611a24565b9060a08401918252873567ffffffffffffffff81116101e357611f269036908501611a24565b9160c0850192835260e084013567ffffffffffffffff81116101e357604099611f55611f629236908801611a24565b9560e08801968752611aec565b90809a91810103126101e357600098604051611f7d8161180d565b6020611f8883611938565b9283835201359063ffffffff821682036121d65760208291015260405191602083017ff3567d1800000000000000000000000000000000000000000000000000000000905260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016602483015260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602c82015260108152612030603082611829565b835260035473ffffffffffffffffffffffffffffffffffffffff169660405198899788977f390775370000000000000000000000000000000000000000000000000000000089526004890160209052516024890161010090526101248901612097916118c7565b945167ffffffffffffffff1660448901525173ffffffffffffffffffffffffffffffffffffffff1660648801525160848701525173ffffffffffffffffffffffffffffffffffffffff1660a486015251908481037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc0160c486015261211b916118c7565b9051908381037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc0160e4850152612151916118c7565b9051908281037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01610104840152612188916118c7565b0381855a94602095f19182156121ca57916121a1575090565b6119f8915060203d6020116121c3575b6121bb8183611829565b810190611b3d565b503d6121b1565b604051903d90823e3d90fd5b8b80fd5b5060009293506020915061223f9073ffffffffffffffffffffffffffffffffffffffff60055416906040519485809481937f489a68f2000000000000000000000000000000000000000000000000000000008352604060048401526044830190611bab565b82602483015203925af19081156106e8576000916121a1575090565b506000929350602091506122bf9073ffffffffffffffffffffffffffffffffffffffff60045416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611bab565b03925af19081156106e8576000916121a1575090565b506000929350602091506122bf9073ffffffffffffffffffffffffffffffffffffffff60035416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611bab565b506000929350602091506122bf9073ffffffffffffffffffffffffffffffffffffffff60065416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611bab565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b90506020813d6020116123fd575b816123e660209383611829565b810103126101e3576123f790611a8b565b38611d40565b3d91506123d9565b919091604051612414816117b8565b60009052602081019061242682611ad7565b67ffffffffffffffff604051917f83826b2b00000000000000000000000000000000000000000000000000000000835216600482015233602482015260208160448173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156106e85760009161274e575b501561239d5760c08101916124c58383611aec565b6004116101e357357fffffffff0000000000000000000000000000000000000000000000000000000016947ffa7c07de0000000000000000000000000000000000000000000000000000000086146126e9577ff3567d18000000000000000000000000000000000000000000000000000000008614612684577fb148ea5f00000000000000000000000000000000000000000000000000000000861461261f577f3047587c00000000000000000000000000000000000000000000000000000000861461259957506040611e288484611aec565b939450506020915061ffff90612602600073ffffffffffffffffffffffffffffffffffffffff6005541692604051968795869485937f489a68f2000000000000000000000000000000000000000000000000000000008552604060048601526044850190611bab565b9116602483015203925af19081156106e8576000916121a1575090565b50506000929350602091506122bf9073ffffffffffffffffffffffffffffffffffffffff60045416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611bab565b50506000929350602091506122bf9073ffffffffffffffffffffffffffffffffffffffff60035416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611bab565b50506000929350602091506122bf9073ffffffffffffffffffffffffffffffffffffffff60065416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611bab565b90506020813d602011612780575b8161276960209383611829565b810103126101e35761277a90611a8b565b386124b0565b3d915061275c565b67ffffffffffffffff1680600052600260205260ff604060002054166005811015610ad257156128865760409060ff8251926127c48185611829565b6001845260208401927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08201368537600052600260205260002054166005811015610ad257600414612814575090565b8151156128575773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f28c4f25e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b604051906128c08261180d565b60606020838281520152565b908160209103126101e3575173ffffffffffffffffffffffffffffffffffffffff811681036101e35790565b81601f820112156101e357805161290e8161186a565b9261291c6040519485611829565b818452602082840101116101e3576119f891602080850191016118a4565b91906040838203126101e357604051906129538261180d565b8193805167ffffffffffffffff81116101e357826129729183016128f8565b835260208101519167ffffffffffffffff83116101e35760209261299692016128f8565b910152565b90608073ffffffffffffffffffffffffffffffffffffffff612a11826129d26129c48780611b5b565b60a0885260a0880191611a98565b9567ffffffffffffffff6129e860208301611938565b166020870152836129fb6040830161174c565b166040870152606081013560608701520161174c565b1691015290565b9091612a226128b3565b506020820192612a3184611ad7565b67ffffffffffffffff604051917fa8d87a3b00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156106e857600091612fae575b5073ffffffffffffffffffffffffffffffffffffffff3391160361239d5767ffffffffffffffff612ae585611ad7565b16600052600260205260ff604060002054166005811015610ad2578015612f7b5760009160048214612cde575060028103612c6857505073ffffffffffffffffffffffffffffffffffffffff60045416915b73ffffffffffffffffffffffffffffffffffffffff8316908115612c285760009450849181612b8f6060612bcc94013580977f00000000000000000000000000000000000000000000000000000000000000006133a7565b6040519687809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260206004840152602483019061299b565b03925af19283156106e857600093612be15750565b3d8085833e612bf08183611829565b810190602081830312612c245780519067ffffffffffffffff8211612c2057612c1c939495500161293a565b9190565b8580fd5b8480fd5b67ffffffffffffffff612c3a86611ad7565b7f28c4f25e000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b600060018203612c945750505073ffffffffffffffffffffffffffffffffffffffff6003541691612b37565b509092907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01612b375760065473ffffffffffffffffffffffffffffffffffffffff169250612b37565b919293905073ffffffffffffffffffffffffffffffffffffffff60055416908115612f3d57612d0c86611ad7565b67ffffffffffffffff604051917f958021a70000000000000000000000000000000000000000000000000000000083521660048201526040602482015260208180612d5a60448201886118c7565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115612f32578591612f03575b5073ffffffffffffffffffffffffffffffffffffffff811615612ec557612e5396508482612dfd612e3b936060849897960135907f00000000000000000000000000000000000000000000000000000000000000006133a7565b604051988995869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260606004860152606485019061299b565b908460248501526003198483030160448501526118c7565b03925af1928315612eb857818094612e6c575b50509190565b915092503d8082853e612e7f8185611829565b8301906040848303126108055783519067ffffffffffffffff82116108055750602091612ead91850161293a565b920151913880612e66565b50604051903d90823e3d90fd5b60248567ffffffffffffffff612eda8a611ad7565b7fe86656fb00000000000000000000000000000000000000000000000000000000835216600452fd5b612f25915060203d602011612f2b575b612f1d8183611829565b8101906128cc565b38612da3565b503d612f13565b6040513d87823e3d90fd5b60248467ffffffffffffffff612f5289611ad7565b7f28c4f25e00000000000000000000000000000000000000000000000000000000835216600452fd5b612fa8907f31603b1200000000000000000000000000000000000000000000000000000000600052611a09565b60246000fd5b612fc7915060203d602011612f2b57612f1d8183611829565b38612ab5565b919290612fd86128b3565b506020830193612fe785611ad7565b67ffffffffffffffff604051917fa8d87a3b00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156106e85760009161332d575b5073ffffffffffffffffffffffffffffffffffffffff3391160361239d5767ffffffffffffffff61309b86611ad7565b16600052600260205260ff60406000205416906005821015610ad2578115613300576000926004831461313057505060028103612c6857505060045473ffffffffffffffffffffffffffffffffffffffff169182908115612c285760009450849181612b8f6060612bcc94013580977f00000000000000000000000000000000000000000000000000000000000000006133a7565b9095929394915073ffffffffffffffffffffffffffffffffffffffff600554169283156132eb5761316081611ad7565b9067ffffffffffffffff604051927f958021a700000000000000000000000000000000000000000000000000000000845216600483015260406024830152602082806131af604482018c6118c7565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9182156132e05786926132bf575b5073ffffffffffffffffffffffffffffffffffffffff8216156132aa57508483612e5382969461325461329295606061ffff990135907f00000000000000000000000000000000000000000000000000000000000000006133a7565b6040519a8b97889687957fb1c71c6500000000000000000000000000000000000000000000000000000000875260606004880152606487019061299b565b921660248501526003198483030160448501526118c7565b8567ffffffffffffffff612eda602493611ad7565b6132d991925060203d602011612f2b57612f1d8183611829565b90386131f8565b6040513d88823e3d90fd5b8467ffffffffffffffff612f52602493611ad7565b612fa8827f31603b1200000000000000000000000000000000000000000000000000000000600052611a09565b613346915060203d602011612f2b57612f1d8183611829565b3861306b565b91908110156128575760051b0190565b73ffffffffffffffffffffffffffffffffffffffff60015416330361337d57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b916020916000916040519073ffffffffffffffffffffffffffffffffffffffff858301937fa9059cbb000000000000000000000000000000000000000000000000000000008552166024830152604482015260448152613408606482611829565b519082855af1156106e8576000513d613481575073ffffffffffffffffffffffffffffffffffffffff81163b155b61343d5750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415613436565b613493816134e1565b90816134cf575b816134a3575090565b6119f891507f0e64dd2900000000000000000000000000000000000000000000000000000000906135d6565b90506134da81613572565b159061349a565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a700000000000000000000000000000000000000000000000000000000602482015260248152613545604482611829565b5191617530fa6000513d82613566575b508161355f575090565b9050151590565b60201115915038613555565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff00000000000000000000000000000000000000000000000000000000602482015260248152613545604482611829565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a70000000000000000000000000000000000000000000000000000000084521660248201526024815261354560448261182956fea164736f6c634300081a000a",
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
