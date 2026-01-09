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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeAggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeAggregator\",\"inputs\":[{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"siloedLockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainNotSupportedByVerifier\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustSetPoolForMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"NoLockOrBurnMechanismSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenPoolUnsupported\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e080604052346103d057806139c5803803809161001d82856103d5565b833981010360e081126103d0578151916001600160a01b038316918284036103d057608090601f1901126103d057604051608081016001600160401b038111828210176103ba57604052610073602083016103f8565b8152610081604083016103f8565b9260208201938452610095606084016103f8565b92604083019384526100a9608082016103f8565b95606084019687526100c960c06100c260a085016103f8565b93016103f8565b9233156103a957600180546001600160a01b03191633179055158015610398575b8015610387575b610376576080526001600160a01b0390811660a05290811660c0528151168015159081610365575b506103445782516001600160a01b03168015159081610333575b506103115781516001600160a01b031680151590816102c6575b506102a45783516001600160a01b03168015159081610293575b506102715751600380546001600160a01b03199081166001600160a01b039384169081179092558451600480548316918516919091179055835160058054831691851691909117905585516006805490921690841617905560408051918252935182166020820152915181169282019290925291511660608201527f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa890608090a16040516134b29081610513823960805181818161125c015281816126bd0152818161292b01528181612c580152612d7c015260a051818181611b6701528181611fda015281816125d60152612b86015260c051818181612384015281816128c50152612d140152f35b835163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61029d915061040c565b1538610167565b505163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b90506102d18161044d565b90816102ff575b816102e6575b50153861014d565b6102f99150633317103160e01b906104de565b386102de565b905061030a816104ac565b15906102d8565b825163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61033d915061040c565b1538610133565b5163be676d1960e01b60009081526001600160a01b03909116600452602490fd5b61036f915061040c565b1538610119565b6303988b8160e61b60005260046000fd5b506001600160a01b038316156100f1565b506001600160a01b038216156100ea565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b601f909101601f19168101906001600160401b038211908210176103ba57604052565b51906001600160a01b03821682036103d057565b6104158161044d565b908161043b575b81610425575090565b6104389150630e64dd2960e01b906104de565b90565b9050610446816104ac565b159061041c565b6000602091604051838101906301ffc9a760e01b82526301ffc9a760e01b60248201526024815261047f6044826103d5565b5191617530fa6000513d826104a0575b5081610499575090565b9050151590565b6020111591503861048f565b6000602091604051838101906301ffc9a760e01b825263ffffffff60e01b60248201526024815261047f6044826103d5565b600090602092604051848101916301ffc9a760e01b835263ffffffff60e01b1660248201526024815261047f6044826103d556fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a7146113455750806315b358e0146112e0578063181f5a7714611281578063240028e8146112105780632c06340414611173578063309292ac14610d605780633907753714610d11578063489a68f214610cbf5780635cb80c5d14610b3f578063673a2a1f14610a3a57806379ba50971461096f5780638926f54f146108f657806389720a621461083a5780638da5cb5b146108065780639a4575b9146107965780639cb406c914610762578063aa86a7541461071d578063b1c71c651461065b578063b794658014610561578063d8aa3f401461049a578063db4c2aef146101e85763f2fde38b1461011157600080fd5b346101e35760206003193601126101e35773ffffffffffffffffffffffffffffffffffffffff61013f611445565b610147613041565b163381146101b957807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101e35760406003193601126101e35760043567ffffffffffffffff81116101e3576102199036906004016116bf565b9060243567ffffffffffffffff81116101e35761023a9036906004016116bf565b610245929192613041565b8084036104705760005b84811061025857005b610263818386613031565b3560058110156101e35767ffffffffffffffff610289610284848988613031565b611931565b16600052600260205260406000206000907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff8416179055506003811480610450575b6103c35760006001821480610430575b61035857506002811480610410575b6103c357600060048214806103a3575b6103585750906001917f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc94471602067ffffffffffffffff610342610284868c8b613031565b16926103516040518092611739565ba20161024f565b6103a160449267ffffffffffffffff610375610284878c8b613031565b7f87d77d330000000000000000000000000000000000000000000000000000000085521660045261171d565bfd5b5073ffffffffffffffffffffffffffffffffffffffff60055416156102ff565b67ffffffffffffffff6103dd61028461040a948988613031565b7f87d77d33000000000000000000000000000000000000000000000000000000006000521660045261171d565b60446000fd5b5073ffffffffffffffffffffffffffffffffffffffff60045416156102ef565b5073ffffffffffffffffffffffffffffffffffffffff60035416156102e0565b5073ffffffffffffffffffffffffffffffffffffffff60065416156102d0565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b346101e35760806003193601126101e3576104b3611445565b6104bb61162c565b906044359161ffff831683036101e3576064359167ffffffffffffffff83116101e35760e0936104f26104fa943690600401611691565b939092612e98565b60c06040519163ffffffff815116835263ffffffff602082015116602084015263ffffffff604082015116604084015263ffffffff606082015116606084015261ffff608082015116608084015261ffff60a08201511660a08401520151151560c0820152f35b346101e35760206003193601126101e35761057a611643565b600067ffffffffffffffff60246105908461308c565b509373ffffffffffffffffffffffffffffffffffffffff60405195869485937fb7946580000000000000000000000000000000000000000000000000000000008552166004840152165afa801561064f57600090610605575b610601906040519182916020835260208301906115e9565b0390f35b3d8082843e610614818461154b565b82019160208184031261064b5780519167ffffffffffffffff831161064857509161064391610601930161244a565b6105e9565b80fd5b5080fd5b6040513d6000823e3d90fd5b346101e35760606003193601126101e35760043567ffffffffffffffff81116101e35760a060031982360301126101e35761069461166f565b906044359067ffffffffffffffff82116101e357366023830112156101e3578160040135926106c28461158c565b916106d0604051938461154b565b84835236602486860101116101e35760006020866107139760246106ff98018388013785010152600401612b19565b6040519283926040845260408401906116f0565b9060208301520390f35b346101e35760206003193601126101e35767ffffffffffffffff61073f611643565b166000526002602052602060ff604060002054166107606040518092611739565bf35b346101e35760006003193601126101e357602073ffffffffffffffffffffffffffffffffffffffff60075416604051908152f35b346101e35760206003193601126101e35760043567ffffffffffffffff81116101e35760a060031982360301126101e3576107cf612405565b506107f0602091604051906107e4848361154b565b6000825260040161256a565b50906106016040519282849384528301906116f0565b346101e35760006003193601126101e357602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101e35760c06003193601126101e357610853611445565b5061085c61162c565b610864611680565b5060843567ffffffffffffffff81116101e357610885903690600401611691565b5050600260a43510156101e35761089b906122da565b60405180916020820160208352815180915260206040840192019060005b8181106108c7575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016108b9565b346101e35760206003193601126101e35767ffffffffffffffff610918611643565b16600052600260205260ff604060002054166005811015610940576020906040519015158152f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b346101e35760006003193601126101e35760005473ffffffffffffffffffffffffffffffffffffffff81163303610a10577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101e35760006003193601126101e35760006060604051610a5b816114ac565b828152826020820152826040820152015261060173ffffffffffffffffffffffffffffffffffffffff6003541673ffffffffffffffffffffffffffffffffffffffff6004541673ffffffffffffffffffffffffffffffffffffffff6005541673ffffffffffffffffffffffffffffffffffffffff600654169160405193610ae1856114ac565b845260208401526040830152606082015260405191829182919091606073ffffffffffffffffffffffffffffffffffffffff816080840195828151168552826020820151166020860152826040820151166040860152015116910152565b346101e35760206003193601126101e35760043567ffffffffffffffff81116101e357610b709036906004016116bf565b9073ffffffffffffffffffffffffffffffffffffffff60075416918215610c955760005b818110610b9d57005b610ba8818385613031565b3573ffffffffffffffffffffffffffffffffffffffff81168091036101e3576040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa801561064f578691600091610c60575b5090816001949392610c22575b50505001610b94565b602081610c517f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e938587613210565b604051908152a3858581610c19565b91506020823d8211610c8d575b81610c7a6020938361154b565b8101031261064857505185906001610c0c565b3d9150610c6d565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101e35760406003193601126101e35760043567ffffffffffffffff81116101e35761010060031982360301126101e357610d08602091610cff61166f565b90600401611f62565b60405190518152f35b346101e35760206003193601126101e35760043567ffffffffffffffff81116101e35761010060031982360301126101e357610d086020916000604051610d57816114f7565b52600401611af1565b346101e35760806003193601126101e357610d79613041565b604051610d85816114ac565b610d8d611445565b80825260243573ffffffffffffffffffffffffffffffffffffffff811681036101e357602083019081526044359073ffffffffffffffffffffffffffffffffffffffff821682036101e3576040840191825273ffffffffffffffffffffffffffffffffffffffff610dfc611468565b9360608601948552168015159081611162575b5061111d5773ffffffffffffffffffffffffffffffffffffffff815116801515908161110c575b506110c75773ffffffffffffffffffffffffffffffffffffffff8251168015159081611063575b5061101e5773ffffffffffffffffffffffffffffffffffffffff835116801515908161100d575b50610fc857610fc39273ffffffffffffffffffffffffffffffffffffffff8593818094817f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa89951167fffffffffffffffffffffffff0000000000000000000000000000000000000000600354161760035551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055551167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065560405191829182919091606073ffffffffffffffffffffffffffffffffffffffff816080840195828151168552826020820151166020860152826040820151166040860152015116910152565b0390a1005b73ffffffffffffffffffffffffffffffffffffffff8351167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61101791506132f3565b1585610e84565b73ffffffffffffffffffffffffffffffffffffffff8251167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b905061106e8161334a565b90816110b5575b81611083575b501585610e5d565b6110af91507f33171031000000000000000000000000000000000000000000000000000000009061343f565b8561107b565b90506110c0816133db565b1590611075565b73ffffffffffffffffffffffffffffffffffffffff9051167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61111691506132f3565b1585610e36565b73ffffffffffffffffffffffffffffffffffffffff8451167fbe676d190000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61116c91506132f3565b1585610e0f565b346101e35760c06003193601126101e35761118c611445565b61119461162c565b61119c611468565b6084359061ffff821682036101e35760a43567ffffffffffffffff81116101e35760a09463ffffffff94859461ffff946111dd6111e9953690600401611691565b949093604435916117b2565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b346101e35760206003193601126101e357602061122b611445565b73ffffffffffffffffffffffffffffffffffffffff604051911673ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148152f35b346101e35760006003193601126101e35761060160408051906112a4818361154b565b601c82527f55534443546f6b656e506f6f6c50726f787920312e372e302d646576000000006020830152519182916020835260208301906115e9565b346101e35760206003193601126101e35773ffffffffffffffffffffffffffffffffffffffff61130e611445565b611316613041565b167fffffffffffffffffffffffff00000000000000000000000000000000000000006007541617600755600080f35b346101e35760206003193601126101e357600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101e357817f33171031000000000000000000000000000000000000000000000000000000006020931490811561141b575b81156113f1575b81156113c7575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836113c0565b7faff2afbf00000000000000000000000000000000000000000000000000000000811491506113b9565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506113b2565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101e357565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036101e357565b359073ffffffffffffffffffffffffffffffffffffffff821682036101e357565b6080810190811067ffffffffffffffff8211176114c857604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff8211176114c857604052565b6040810190811067ffffffffffffffff8211176114c857604052565b60e0810190811067ffffffffffffffff8211176114c857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176114c857604052565b67ffffffffffffffff81116114c857601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106115d95750506000910152565b81810151838201526020016115c9565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093611625815180928187528780880191016115c6565b0116010190565b6024359067ffffffffffffffff821682036101e357565b6004359067ffffffffffffffff821682036101e357565b359067ffffffffffffffff821682036101e357565b6024359061ffff821682036101e357565b6064359061ffff821682036101e357565b9181601f840112156101e35782359167ffffffffffffffff83116101e357602083818601950101116101e357565b9181601f840112156101e35782359167ffffffffffffffff83116101e3576020808501948460051b0101116101e357565b61171a91602061170983516040845260408401906115e9565b9201519060208184039101526115e9565b90565b600581101561094057602452565b600581101561094057600452565b9060058210156109405752565b519063ffffffff821682036101e357565b519061ffff821682036101e357565b519081151582036101e357565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9390959194926117c18761308c565b9690966118165767ffffffffffffffff8816600052600260205261181060ff604060002054167f31603b120000000000000000000000000000000000000000000000000000000060005261172b565b60246000fd5b6040517f2c06340400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff968716600482015267ffffffffffffffff98909816602489015260448801528416606487015261ffff16608486015260c060a486015260a09385939092849283916118a09160c4840191611773565b0392165afa90811561064f57600091829383809381936118c3575b509493929190565b9450925093505060a0823d60a011611929575b816118e360a0938361154b565b81010312610648575080516118fa60208301611746565b61190660408401611746565b9361191f608061191860608701611757565b9501611766565b91949392386118bb565b3d91506118d6565b3567ffffffffffffffff811681036101e35790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101e3570180359067ffffffffffffffff82116101e3576020019181360383136101e357565b908160209103126101e357604051906119af826114f7565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156101e357016020813591019167ffffffffffffffff82116101e35781360383136101e357565b61171a91611ae3611ad8611abd611a2f611a1f86806119b5565b6101008752610100870191611773565b67ffffffffffffffff611a446020880161165a565b16602086015273ffffffffffffffffffffffffffffffffffffffff611a6b6040880161148b565b1660408601526060860135606086015273ffffffffffffffffffffffffffffffffffffffff611a9c6080880161148b565b166080860152611aaf60a08701876119b5565b9086830360a0880152611773565b611aca60c08601866119b5565b9085830360c0870152611773565b9260e08101906119b5565b9160e0818503910152611773565b90604051611afe816114f7565b60009052611b0e60208301611931565b67ffffffffffffffff604051917f83826b2b00000000000000000000000000000000000000000000000000000000835216600482015233602482015260208160448173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561064f57600091611f28575b5015611efa5760c0820191611bad8382611946565b6004116101e357357fffffffff0000000000000000000000000000000000000000000000000000000016927ffa7c07de000000000000000000000000000000000000000000000000000000008414611e98577ff3567d18000000000000000000000000000000000000000000000000000000008414611e36577fb148ea5f000000000000000000000000000000000000000000000000000000008414611dd4577f3047587c000000000000000000000000000000000000000000000000000000008414611d5557611c8060409183611946565b905014611cb557827fcacdaf2b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6000919250611d1660209173ffffffffffffffffffffffffffffffffffffffff60035416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611a05565b03925af190811561064f57600091611d2c575090565b61171a915060203d602011611d4e575b611d46818361154b565b810190611997565b503d611d3c565b506000919250611db860209173ffffffffffffffffffffffffffffffffffffffff60055416906040519485809481937f489a68f2000000000000000000000000000000000000000000000000000000008352604060048401526044830190611a05565b82602483015203925af190811561064f57600091611d2c575090565b506000919250611d1660209173ffffffffffffffffffffffffffffffffffffffff60045416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611a05565b506000919250611d1660209173ffffffffffffffffffffffffffffffffffffffff60035416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611a05565b506000919250611d1660209173ffffffffffffffffffffffffffffffffffffffff60065416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611a05565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b90506020813d602011611f5a575b81611f436020938361154b565b810103126101e357611f5490611766565b38611b98565b3d9150611f36565b919091604051611f71816114f7565b60009052611f8160208201611931565b67ffffffffffffffff604051917f83826b2b00000000000000000000000000000000000000000000000000000000835216600482015233602482015260208160448173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561064f576000916122a0575b5015611efa5760c0810161201f8183611946565b6004116101e357357fffffffff0000000000000000000000000000000000000000000000000000000016937ffa7c07de00000000000000000000000000000000000000000000000000000000851461223d577ff3567d180000000000000000000000000000000000000000000000000000000085146121da577fb148ea5f000000000000000000000000000000000000000000000000000000008514612177577f3047587c0000000000000000000000000000000000000000000000000000000085146120f35750611c8060409183611946565b9293505061ffff60209161215a600073ffffffffffffffffffffffffffffffffffffffff6005541692604051968795869485937f489a68f2000000000000000000000000000000000000000000000000000000008552604060048601526044850190611a05565b9116602483015203925af190811561064f57600091611d2c575090565b50506000919250611d1660209173ffffffffffffffffffffffffffffffffffffffff60045416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611a05565b50506000919250611d1660209173ffffffffffffffffffffffffffffffffffffffff60035416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611a05565b50506000919250611d1660209173ffffffffffffffffffffffffffffffffffffffff60065416906040519485809481937f390775370000000000000000000000000000000000000000000000000000000083528760048401526024830190611a05565b90506020813d6020116122d2575b816122bb6020938361154b565b810103126101e3576122cc90611766565b3861200b565b3d91506122ae565b67ffffffffffffffff1680600052600260205260ff60406000205416600581101561094057156123d85760409060ff825192612316818561154b565b6001845260208401927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0820136853760005260026020526000205416600581101561094057600414612366575090565b8151156123a95773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f28c4f25e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6040519061241282611513565b60606020838281520152565b908160209103126101e3575173ffffffffffffffffffffffffffffffffffffffff811681036101e35790565b81601f820112156101e35780516124608161158c565b9261246e604051948561154b565b818452602082840101116101e35761171a91602080850191016115c6565b91906040838203126101e357604051906124a582611513565b8193805167ffffffffffffffff81116101e357826124c491830161244a565b835260208101519167ffffffffffffffff83116101e3576020926124e8920161244a565b910152565b90608073ffffffffffffffffffffffffffffffffffffffff6125638261252461251687806119b5565b60a0885260a0880191611773565b9567ffffffffffffffff61253a6020830161165a565b1660208701528361254d6040830161148b565b166040870152606081013560608701520161148b565b1691015290565b9091612574612405565b50602082019261258384611931565b67ffffffffffffffff604051917fa8d87a3b00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561064f57600091612afa575b5073ffffffffffffffffffffffffffffffffffffffff33911603611efa5767ffffffffffffffff61263785611931565b16600052600260205260ff604060002054166005811015610940578015612acd57600091600482146128305750600281036127ba57505073ffffffffffffffffffffffffffffffffffffffff60045416915b73ffffffffffffffffffffffffffffffffffffffff831690811561277a57600094508491816126e1606061271e94013580977f0000000000000000000000000000000000000000000000000000000000000000613210565b6040519687809481937f9a4575b90000000000000000000000000000000000000000000000000000000083526020600484015260248301906124ed565b03925af192831561064f576000936127335750565b3d8085833e612742818361154b565b8101906020818303126127765780519067ffffffffffffffff82116127725761276e939495500161248c565b9190565b8580fd5b8480fd5b67ffffffffffffffff61278c86611931565b7f28c4f25e000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6000600182036127e65750505073ffffffffffffffffffffffffffffffffffffffff6003541691612689565b509092907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016126895760065473ffffffffffffffffffffffffffffffffffffffff169250612689565b919293905073ffffffffffffffffffffffffffffffffffffffff60055416908115612a8f5761285e86611931565b67ffffffffffffffff604051917f958021a700000000000000000000000000000000000000000000000000000000835216600482015260406024820152602081806128ac60448201886115e9565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115612a84578591612a55575b5073ffffffffffffffffffffffffffffffffffffffff811615612a17576129a59650848261294f61298d936060849897960135907f0000000000000000000000000000000000000000000000000000000000000000613210565b604051988995869485937fb1c71c650000000000000000000000000000000000000000000000000000000085526060600486015260648501906124ed565b908460248501526003198483030160448501526115e9565b03925af1928315612a0a578180946129be575b50509190565b915092503d8082853e6129d1818561154b565b8301906040848303126106485783519067ffffffffffffffff821161064857506020916129ff91850161248c565b9201519138806129b8565b50604051903d90823e3d90fd5b60248567ffffffffffffffff612a2c8a611931565b7fe86656fb00000000000000000000000000000000000000000000000000000000835216600452fd5b612a77915060203d602011612a7d575b612a6f818361154b565b81019061241e565b386128f5565b503d612a65565b6040513d87823e3d90fd5b60248467ffffffffffffffff612aa489611931565b7f28c4f25e00000000000000000000000000000000000000000000000000000000835216600452fd5b611810907f31603b120000000000000000000000000000000000000000000000000000000060005261172b565b612b13915060203d602011612a7d57612a6f818361154b565b38612607565b919290612b24612405565b506020830193612b3385611931565b67ffffffffffffffff604051917fa8d87a3b00000000000000000000000000000000000000000000000000000000835216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561064f57600091612e79575b5073ffffffffffffffffffffffffffffffffffffffff33911603611efa5767ffffffffffffffff612be786611931565b16600052600260205260ff60406000205416906005821015610940578115612e4c5760009260048314612c7c575050600281036127ba57505060045473ffffffffffffffffffffffffffffffffffffffff16918290811561277a57600094508491816126e1606061271e94013580977f0000000000000000000000000000000000000000000000000000000000000000613210565b9095929394915073ffffffffffffffffffffffffffffffffffffffff60055416928315612e3757612cac81611931565b9067ffffffffffffffff604051927f958021a70000000000000000000000000000000000000000000000000000000084521660048301526040602483015260208280612cfb604482018c6115e9565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa918215612e2c578692612e0b575b5073ffffffffffffffffffffffffffffffffffffffff821615612df6575084836129a5829694612da0612dde95606061ffff990135907f0000000000000000000000000000000000000000000000000000000000000000613210565b6040519a8b97889687957fb1c71c650000000000000000000000000000000000000000000000000000000087526060600488015260648701906124ed565b921660248501526003198483030160448501526115e9565b8567ffffffffffffffff612a2c602493611931565b612e2591925060203d602011612a7d57612a6f818361154b565b9038612d44565b6040513d88823e3d90fd5b8467ffffffffffffffff612aa4602493611931565b611810827f31603b120000000000000000000000000000000000000000000000000000000060005261172b565b612e92915060203d602011612a7d57612a6f818361154b565b38612bb7565b9392919091604051612ea98161152f565b6000815260006020820152600060408201526000606082015260006080820152600060a0820152600060c082015293612ee18461308c565b939093612ef2575050505050905090565b60e0955095612f6f73ffffffffffffffffffffffffffffffffffffffff93849361ffff67ffffffffffffffff9a6040519b8c9a8b998a987fd8aa3f40000000000000000000000000000000000000000000000000000000008a52166004890152166024870152166044850152608060648501526084840191611773565b0392165afa90811561064f57600091612f86575090565b905060e0813d60e011613029575b81612fa160e0938361154b565b810103126101e35761302160c060405192612fbb8461152f565b612fc481611746565b8452612fd260208201611746565b6020850152612fe360408201611746565b6040850152612ff460608201611746565b606085015261300560808201611757565b608085015261301660a08201611757565b60a085015201611766565b60c082015290565b3d9150612f94565b91908110156123a95760051b0190565b73ffffffffffffffffffffffffffffffffffffffff60015416330361306257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90600067ffffffffffffffff600093169283600052600260205260ff604060002054169160058310156109405782156131e35760006004840361313d5750505073ffffffffffffffffffffffffffffffffffffffff60055416906001935b73ffffffffffffffffffffffffffffffffffffffff83161561310c5750509190565b61040a92507f87d77d330000000000000000000000000000000000000000000000000000000060005260045261171d565b50909391906002820361316b575073ffffffffffffffffffffffffffffffffffffffff60045416915b6130ea565b60006001830361319657505073ffffffffffffffffffffffffffffffffffffffff60035416916130ea565b50917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd8201613166576006546001955073ffffffffffffffffffffffffffffffffffffffff1692506130ea565b611810837f31603b120000000000000000000000000000000000000000000000000000000060005261172b565b916020916000916040519073ffffffffffffffffffffffffffffffffffffffff858301937fa9059cbb00000000000000000000000000000000000000000000000000000000855216602483015260448201526044815261327160648261154b565b519082855af11561064f576000513d6132ea575073ffffffffffffffffffffffffffffffffffffffff81163b155b6132a65750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6001141561329f565b6132fc8161334a565b9081613338575b8161330c575090565b61171a91507f0e64dd29000000000000000000000000000000000000000000000000000000009061343f565b9050613343816133db565b1590613303565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a7000000000000000000000000000000000000000000000000000000006024820152602481526133ae60448261154b565b5191617530fa6000513d826133cf575b50816133c8575090565b9050151590565b602011159150386133be565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff000000000000000000000000000000000000000000000000000000006024820152602481526133ae60448261154b565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a7000000000000000000000000000000000000000000000000000000008452166024820152602481526133ae60448261154b56fea164736f6c634300081a000a",
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
