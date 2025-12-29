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
	LegacyCctpV1Pool  common.Address
	CctpV1Pool        common.Address
	CctpV2Pool        common.Address
	CctpV2PoolWithCCV common.Address
}

var USDCTokenPoolProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockReleasePoolAddress\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"lockOrBurnOut\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateLockReleasePoolAddresses\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"lockReleasePools\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockReleasePoolUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct USDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2PoolWithCCV\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CCVCompatiblePoolNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainNotSupportedByVerifier\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enum USDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoLockOrBurnMechanismSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenPoolUnsupported\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60e06040523461022857604051613aa738819003601f8101601f191683016001600160401b03811184821017610212578392829160405283398101039060e082126102285780516001600160a01b0381169283820361022857608090601f19011261022857604051608081016001600160401b038111828210176102125760405261008c6020840161022d565b815261009a6040840161022d565b602082019081526100ad6060850161022d565b91604081019283526100c16080860161022d565b93606082019485526100e160c06100da60a0890161022d565b970161022d565b96331561020157600180546001600160a01b031916331790551580156101f0575b80156101df575b6101ce5760805251600480546001600160a01b03199081166001600160a01b0393841617909155915160058054841691831691909117905591516006805483169184169190911790559151600780549093169082161790915590811660a0521660c052604051613865908161024282396080518181816105e9015281816130570152818161322101526134fb015260a051818181611e170152818161225201528181612eaa015261339c015260c05181818161266301528181612fe8015261348c0152f35b6303988b8160e61b60005260046000fd5b506001600160a01b03871615610109565b506001600160a01b03861615610102565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b51906001600160a01b03821682036102285756fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610167578063181f5a77146101625780631c255a9d1461015d578063240028e8146101585780632c06340414610153578063309292ac1461014e5780633907753714610149578063489a68f214610144578063673a2a1f1461013f57806379ba50971461013a5780638302d09f146101355780638926f54f1461013057806389720a621461012b5780638da5cb5b146101265780639a4575b914610121578063aa86a7541461011c578063b1c71c6514610117578063b794658014610112578063d8aa3f401461010d578063db4c2aef146101085763f2fde38b1461010357600080fd5b61161f565b61154a565b61147f565b6112dc565b611210565b6111ab565b6110f1565b611064565b610fc6565b610edb565b610cfe565b610b74565b610ad7565b610a71565b610a00565b61082e565b610648565b61057b565b6104df565b610441565b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610288576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610288577f3317103100000000000000000000000000000000000000000000000000000000811490811561025e575b8115610234575b811561020a575b506040519015158152602090f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386101fc565b7faff2afbf00000000000000000000000000000000000000000000000000000000811491506101f5565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506101ee565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff8211176102d857604052565b61028d565b6080810190811067ffffffffffffffff8211176102d857604052565b60e0810190811067ffffffffffffffff8211176102d857604052565b6040810190811067ffffffffffffffff8211176102d857604052565b610100810190811067ffffffffffffffff8211176102d857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102d857604052565b6040519061039f6101008361034e565b565b67ffffffffffffffff81116102d857601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106103ee5750506000910152565b81810151838201526020016103de565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361043a815180928187528780880191016103db565b0116010190565b346102885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610288576104be6040805190610482818361034e565b601c82527f55534443546f6b656e506f6f6c50726f787920312e372e302d646576000000006020830152519182916020835260208301906103fe565b0390f35b67ffffffffffffffff81160361028857565b359061039f826104c2565b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885767ffffffffffffffff600435610523816104c2565b166000526003602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b73ffffffffffffffffffffffffffffffffffffffff81160361028857565b359061039f82610552565b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760206004356105b881610552565b73ffffffffffffffffffffffffffffffffffffffff604051911673ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148152f35b61ffff81160361028857565b9181601f840112156102885782359167ffffffffffffffff8311610288576020838186019501011161028857565b346102885760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760043561068381610552565b60243590610690826104c2565b604435606435916106a083610552565b6084356106ac8161060e565b60a43567ffffffffffffffff8111610288576106cc90369060040161061a565b909160009573ffffffffffffffffffffffffffffffffffffffff60075416958615610806579373ffffffffffffffffffffffffffffffffffffffff9361ffff67ffffffffffffffff9a97948661077b9560a09b996040519e8f9c8d9b8c9b7f2c063404000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015216608485015260c060a485015260c484019161178a565b03915afa8015610801578190828081926107c5575b6040805194855263ffffffff958616602086015294169383019390935261ffff9092166060820152901515608082015260a090f35b50505050506107ee6104be9160a03d60a0116107fa575b6107e6818361034e565b810190611744565b92945084939291610790565b503d6107dc565b6117c9565b6004887f0f480b32000000000000000000000000000000000000000000000000000000008152fd5b346102885760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028857600061086761291f565b73ffffffffffffffffffffffffffffffffffffffff6108846117d5565b161515806109dd575b6109d2576108b561089c6117e1565b73ffffffffffffffffffffffffffffffffffffffff1690565b1515806109be575b6109b3576108cc61089c6117ed565b610990575b6108dc61089c6117f9565b151580610973575b610924576108f061180f565b7f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa86040518061091e8161193d565b0390a180f35b61097061092f6117f9565b7fbe676d1900000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff16600452602490565b90fd5b5061098b6109876109826117f9565b61296a565b1590565b6108e4565b6109a361098761099e6117ed565b6129c1565b156108d15761097061092f6117ed565b61097061092f6117e1565b506109cd6109876109826117e1565b6108bd565b61097061092f6117d5565b506109ec6109876109826117d5565b61088d565b90816101009103126102885790565b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760043567ffffffffffffffff811161028857610a68610a5460209236906004016109f1565b6000604051610a62816102bc565b52611d99565b60405190518152f35b346102885760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760043567ffffffffffffffff811161028857610a68610ac560209236906004016109f1565b60243590610ad28261060e565b61221d565b346102885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760006060604051610b16816102dd565b82815282602082015282604082015201526080610b3161247b565b73ffffffffffffffffffffffffffffffffffffffff6060604051928281511684528260208201511660208501528260408201511660408501520151166060820152f35b346102885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760005473ffffffffffffffffffffffffffffffffffffffff81163303610c33577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156102885782359167ffffffffffffffff8311610288576020808501948460051b01011161028857565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102885760043567ffffffffffffffff81116102885781610cd791600401610c5d565b929092916024359067ffffffffffffffff821161028857610cfa91600401610c5d565b9091565b3461028857610d0c36610c8e565b9291610d1661291f565b838103610eb15760005b818110610d2957005b610d3f61089c610d3a83888761252f565b611805565b151580610e97575b610e405780610dcf610d5f610d3a600194898861252f565b610d8f610d75610d7085888b61252f565b6119f1565b67ffffffffffffffff166000526003602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b610ddd610d7082858861252f565b7f2a4ec2a96b51064b74fa8f2157f98cf8fb9fd4dcef4e9fdf3c44c0d74d0e826467ffffffffffffffff610e15610d3a858b8a61252f565b60405173ffffffffffffffffffffffffffffffffffffffff919091168152921691602090a201610d20565b610e51610d3a610e9392878661252f565b7fbe676d190000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b6000fd5b50610eac610987610982610d3a84898861252f565b610d47565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885767ffffffffffffffff600435610f1f816104c2565b16600052600260205260ff604060002054166005811015610f47576040519015158152602090f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b602060408183019282815284518094520192019060005b818110610f9a5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610f8d565b346102885760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760043561100181610552565b6024359061100e826104c2565b60443560643561101d8161060e565b60843567ffffffffffffffff81116102885761103d90369060040161061a565b9160a435936002851015610288576104be9661105896612544565b60405191829182610f76565b346102885760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028857602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b908160a09103126102885790565b6110ee9160206110dd83516040845260408401906103fe565b9201519060208184039101526103fe565b90565b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760043567ffffffffffffffff8111610288576111409036906004016110b6565b6111486126de565b506111666020916040519061115d848361034e565b60008252612e30565b50906104be6040519282849384528301906110c4565b60051115610f4757565b906024916005811015610f4757600452565b919060208301926005821015610f475752565b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885767ffffffffffffffff6004356111ef816104c2565b1660005260026020526104be60ff6040600020541660405191829182611198565b346102885760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885760043567ffffffffffffffff81116102885761125f9036906004016110b6565b6024359061126c8261060e565b6044359067ffffffffffffffff8211610288576112c1926112a76112976112ad94369060040161061a565b61129f6126de565b5036916126f7565b91613363565b6040519283926040845260408401906110c4565b9060208301520390f35b9060206110ee9281815201906103fe565b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028857600435611317816104c2565b73ffffffffffffffffffffffffffffffffffffffff600754169081156113e05767ffffffffffffffff602460009260405194859384927fb79465800000000000000000000000000000000000000000000000000000000084521660048301525afa801561080157600090611396575b6104be90604051918291826112cb565b3d8082843e6113a5818461034e565b8201916020818403126113dc5780519167ffffffffffffffff83116113d95750916113d4916104be930161272e565b611386565b80fd5b5080fd5b7f0f480b320000000000000000000000000000000000000000000000000000000060005260046000fd5b61039f9092919260c08060e083019563ffffffff815116845263ffffffff602082015116602085015263ffffffff604082015116604085015263ffffffff606082015116606085015261ffff608082015116608085015261147660a082015160a086019061ffff169052565b01511515910152565b346102885760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610288576004356114ba81610552565b602435906114c7826104c2565b604435916114d48361060e565b6064359167ffffffffffffffff8311610288576104be936114fc61153e94369060040161061a565b93909260405161150b816102f9565b6000815260006020820152600060408201526000606082015260006080820152600060a0820152600060c0820152612770565b6040519182918261140a565b346102885761155836610c8e565b9261156161291f565b838303610eb15760005b83811061157457005b806115be61158d611588600194898861252f565b6128db565b67ffffffffffffffff6115a184898861252f565b356115ab816104c2565b16600052600260205260406000206128e8565b6115cc610d7082878661252f565b7f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc9447167ffffffffffffffff611604611588858b8a61252f565b92611616604051928392169482611198565b0390a20161156b565b346102885760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102885773ffffffffffffffffffffffffffffffffffffffff60043561166f81610552565b61167761291f565b163381146116e957807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b63ffffffff81160361028857565b519061039f82611713565b519061039f8261060e565b5190811515820361028857565b908160a091031261028857805191602082015161176081611713565b91604081015161176f81611713565b916110ee608060608401516117838161060e565b9301611737565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b6040513d6000823e3d90fd5b6024356110ee81610552565b6044356110ee81610552565b6064356110ee81610552565b6004356110ee81610552565b356110ee81610552565b73ffffffffffffffffffffffffffffffffffffffff60043561183081610552565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045573ffffffffffffffffffffffffffffffffffffffff60243561187b81610552565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055573ffffffffffffffffffffffffffffffffffffffff6044356118c681610552565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065573ffffffffffffffffffffffffffffffffffffffff60643561191181610552565b167fffffffffffffffffffffffff00000000000000000000000000000000000000006007541617600755565b90608082019173ffffffffffffffffffffffffffffffffffffffff60043561196481610552565b16815273ffffffffffffffffffffffffffffffffffffffff60243561198881610552565b16602082015273ffffffffffffffffffffffffffffffffffffffff6044356119af81610552565b166040820152606073ffffffffffffffffffffffffffffffffffffffff6064356119d881610552565b16910152565b604051906119eb826102bc565b60008252565b356110ee816104c2565b90816020910312610288576110ee90611737565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610288570180359067ffffffffffffffff82116102885760200191813603831361028857565b906004116102885790600490565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110611aa2575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b908160209103126102885760405190611aec826102bc565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561028857016020813591019167ffffffffffffffff821161028857813603831361028857565b6110ee91611c2f611c24611c09611b6c611b5c8680611af2565b610100875261010087019161178a565b611b8c611b7b602088016104d4565b67ffffffffffffffff166020870152565b611bb8611b9b60408801610570565b73ffffffffffffffffffffffffffffffffffffffff166040870152565b60608601356060860152611bee611bd160808801610570565b73ffffffffffffffffffffffffffffffffffffffff166080870152565b611bfb60a0870187611af2565b9086830360a088015261178a565b611c1660c0860186611af2565b9085830360c087015261178a565b9260e0810190611af2565b9160e081850391015261178a565b9060206110ee928181520190611b42565b9061ffff611c69602092959495604085526040850190611b42565b9416910152565b906110ee916020815260e0611d65611d32611c99855161010060208701526101208601906103fe565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff16606086015260608601516080860152611cfe608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526103fe565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526103fe565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526103fe565b611da16119de565b506020810190611dfe6020611db5846119f1565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015233602482015291829081906044820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610801576000916121ee575b50156121c05760c08101611e6e611e68611e628385611a0f565b90611a60565b90611a6e565b927fffffffff000000000000000000000000000000000000000000000000000000008416907ffa7c07de00000000000000000000000000000000000000000000000000000000821461218357507ff3567d18000000000000000000000000000000000000000000000000000000008114612152577fb148ea5f000000000000000000000000000000000000000000000000000000008114612121577f3047587c00000000000000000000000000000000000000000000000000000000146120ba57611f3b60409183611a0f565b905014611f92577fcacdaf2b000000000000000000000000000000000000000000000000000000006000527fffffffff00000000000000000000000000000000000000000000000000000000821660045260246000fd5b6000915080611faf611fa960e06020940183611a0f565b90612a33565b83146120535761201490611fde61089c61089c60045473ffffffffffffffffffffffffffffffffffffffff1690565b906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301611c3d565b03925af19081156108015760009161202a575090565b6110ee915060203d60201161204c575b612044818361034e565b810190611ad4565b503d61203a565b6120149061208561207f61089c61089c60055473ffffffffffffffffffffffffffffffffffffffff1690565b91612b3c565b6040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301611c70565b50612014915060006020916120ea61089c61089c60075473ffffffffffffffffffffffffffffffffffffffff1690565b90826040518096819582947f489a68f200000000000000000000000000000000000000000000000000000000845260048401611c4e565b505060009150612014602091611fde61089c61089c60065473ffffffffffffffffffffffffffffffffffffffff1690565b505060009150612014602091611fde61089c61089c60055473ffffffffffffffffffffffffffffffffffffffff1690565b600094506120149250602093915061089c61089c6121a6610d75611fde946119f1565b5473ffffffffffffffffffffffffffffffffffffffff1690565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b612210915060203d602011612216575b612208818361034e565b8101906119fb565b38611e48565b503d6121fe565b6122256119de565b5060208101916122396020611db5856119f1565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156108015760009161245c575b50156121c05760c082019061229e611e68611e628486611a0f565b937fffffffff000000000000000000000000000000000000000000000000000000008516907ffa7c07de00000000000000000000000000000000000000000000000000000000821461243757507ff3567d18000000000000000000000000000000000000000000000000000000008114612405577fb148ea5f0000000000000000000000000000000000000000000000000000000081146123d3577f3047587c000000000000000000000000000000000000000000000000000000001461236c5750611f3b60409183611a0f565b9050612014925060209161239b61089c61089c60075473ffffffffffffffffffffffffffffffffffffffff1690565b9060006040518096819582947f489a68f200000000000000000000000000000000000000000000000000000000845260048401611c4e565b50505060009150612014602091611fde61089c61089c60065473ffffffffffffffffffffffffffffffffffffffff1690565b50505060009150612014602091611fde61089c61089c60055473ffffffffffffffffffffffffffffffffffffffff1690565b6000955061201493506020949250611fde915061089c6121a6610d7561089c936119f1565b612475915060203d60201161221657612208818361034e565b38612283565b60405190612488826102dd565b8173ffffffffffffffffffffffffffffffffffffffff60045416815273ffffffffffffffffffffffffffffffffffffffff60055416602082015273ffffffffffffffffffffffffffffffffffffffff600654166040820152606073ffffffffffffffffffffffffffffffffffffffff60075416910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b919081101561253f5760051b0190565b612500565b5094505050505073ffffffffffffffffffffffffffffffffffffffff60075416156113e0576110ee90612597565b67ffffffffffffffff81116102d85760051b60200190565b80511561253f5760200190565b67ffffffffffffffff8116600052600260205260ff604060002054166125bc8161117c565b156126a75760046126386126316125d36001612572565b936125e1604051958661034e565b600185527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06126106001612572565b0136602087013767ffffffffffffffff166000526002602052604060002090565b5460ff1690565b6126418161117c565b146126495790565b6110ee73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001661268c8361258a565b9073ffffffffffffffffffffffffffffffffffffffff169052565b7f28c4f25e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b604051906126eb82610315565b60606020838281520152565b929192612703826103a1565b91612711604051938461034e565b829481845281830111610288578281602093846000960137010152565b81601f82011215610288578051612744816103a1565b92612752604051948561034e565b81845260208284010111610288576110ee91602080850191016103db565b5090919373ffffffffffffffffffffffffffffffffffffffff600754169182156113e05761280d60e09573ffffffffffffffffffffffffffffffffffffffff9361ffff67ffffffffffffffff996040519a8b998a9889987fd8aa3f40000000000000000000000000000000000000000000000000000000008a5216600489015216602487015216604485015260806064850152608484019161178a565b03915afa90811561080157600091612823575090565b6110ee915060e03d60e011612845575b61283d818361034e565b81019061284c565b503d612833565b908160e0910312610288576128d360c060405192612869846102f9565b805161287481611713565b8452602081015161288481611713565b602085015261289560408201611721565b60408501526128a660608201611721565b60608501526128b76080820161172c565b60808501526128c860a0820161172c565b60a085015201611737565b60c082015290565b3560058110156102885790565b906005811015610f475760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b73ffffffffffffffffffffffffffffffffffffffff60015416330361294057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b61297381613586565b90816129af575b81612983575090565b6110ee91507f0e64dd29000000000000000000000000000000000000000000000000000000009061367b565b90506129ba81613617565b159061297a565b6129ca81613586565b9081612a06575b816129da575090565b6110ee91507f33171031000000000000000000000000000000000000000000000000000000009061367b565b9050612a1181613617565b15906129d1565b9080601f83011215610288578160206110ee933591016126f7565b73ffffffffffffffffffffffffffffffffffffffff60045416918215612afa578101906020818303126102885780359067ffffffffffffffff821161028857016040818303126102885760405191612a8a83610315565b813567ffffffffffffffff81116102885781612aa7918401612a18565b8352602082013567ffffffffffffffff811161028857612af693612ad561089c9360749361089c9601612a18565b602082015251015173ffffffffffffffffffffffffffffffffffffffff1690565b1490565b505050600090565b9081604091031261028857602060405191612b1c83610315565b8035612b27816104c2565b83520135612b3481611713565b602082015290565b606060e0604051612b4c81610331565b8281526000602082015260006040820152600083820152600060808201528260a08201528260c082015201526101008136031261028857612b8b61038f565b90803567ffffffffffffffff811161028857612baa9036908301612a18565b8252612bb8602082016104d4565b6020830152612bc960408201610570565b604083015260608101356060830152612be460808201610570565b608083015260a081013567ffffffffffffffff811161028857612c0a9036908301612a18565b60a083015260c08101803567ffffffffffffffff811161028857612c319036908401612a18565b9160c0840192835260e081013567ffffffffffffffff811161028857612c7c92612c7792612c65612c6f9336908301612a18565b60e0880152611a0f565b810190612b02565b6136e1565b905290565b9081602091031261028857516110ee81610552565b60409067ffffffffffffffff6110ee949316815281602082015201906103fe565b91906040838203126102885760405190612cd082610315565b8193805167ffffffffffffffff81116102885782612cef91830161272e565b835260208101519167ffffffffffffffff831161028857602092612d13920161272e565b910152565b919060408382031261028857825167ffffffffffffffff811161028857602091612d43918501612cb7565b92015190565b90608073ffffffffffffffffffffffffffffffffffffffff81612d7d612d6f8680611af2565b60a0875260a087019161178a565b9467ffffffffffffffff6020820135612d95816104c2565b166020860152826040820135612daa81610552565b166040860152606081013560608601520135612dc581610552565b1691015290565b61ffff612de56110ee9593606084526060840190612d49565b9316602082015260408184039101526103fe565b9060208282031261028857815167ffffffffffffffff8111610288576110ee9201612cb7565b9060206110ee928181520190612d49565b9190612e3a6126de565b506020830192612e916020612e4e866119f1565b6040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561080157600091613344575b5073ffffffffffffffffffffffffffffffffffffffff339116036121c057612f22612631612f08866119f1565b67ffffffffffffffff166000526002602052604060002090565b612f2b8161117c565b801561331757612f3961247b565b90600090612f468161117c565b600481036131b957505060600173ffffffffffffffffffffffffffffffffffffffff612f86825173ffffffffffffffffffffffffffffffffffffffff1690565b161561317757612fcf602084612f9b886119f1565b60405193849283927f958021a700000000000000000000000000000000000000000000000000000000845260048401612c96565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561080157600091613148575b5073ffffffffffffffffffffffffffffffffffffffff811615613106576130cc9550600061309561089c61089c8397969561307b85966060890135907f0000000000000000000000000000000000000000000000000000000000000000613775565b5173ffffffffffffffffffffffffffffffffffffffff1690565b92604051978895869485937fb1c71c6500000000000000000000000000000000000000000000000000000000855260048501612dcc565b03925af1918215610801576000906000936130e657509190565b9050610cfa9192503d806000833e6130fe818361034e565b810190612d18565b610e93613112876119f1565b7fe86656fb0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b61316a915060203d602011613170575b613162818361034e565b810190612c81565b38613019565b503d613158565b610e93613183866119f1565b7f28c4f25e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b9091929593506131c88161117c565b600281036132be5750506040015173ffffffffffffffffffffffffffffffffffffffff16925b73ffffffffffffffffffffffffffffffffffffffff84169182156132b2575061327a6000928392613245606082013580987f0000000000000000000000000000000000000000000000000000000000000000613775565b6040519485809481937f9a4575b900000000000000000000000000000000000000000000000000000000835260048301612e1f565b03925af19081156108015760009161329157509190565b6132ae91503d806000833e6132a6818361034e565b810190612df9565b9190565b613183610e93916119f1565b6132c78161117c565b600181036132f25750506020015173ffffffffffffffffffffffffffffffffffffffff165b926131ee565b6003919592506133018161117c565b036131ee5792506132ec6121a6610d75836119f1565b610e93907f31603b1200000000000000000000000000000000000000000000000000000000600052611186565b61335d915060203d60201161317057613162818361034e565b38612edb565b9291909261336f6126de565b5060208101936133836020612e4e876119f1565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561080157600091613567575b5073ffffffffffffffffffffffffffffffffffffffff339116036121c0576133fa612631612f08876119f1565b6134038161117c565b80156133175761341161247b565b9060009061341e8161117c565b6004810361355657505060600173ffffffffffffffffffffffffffffffffffffffff61345e825173ffffffffffffffffffffffffffffffffffffffff1690565b161561354a57613473602085612f9b896119f1565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156108015760009161352b575b5073ffffffffffffffffffffffffffffffffffffffff81161561351f576130cc965061309561089c61089c60009796959461307b89956060890135907f0000000000000000000000000000000000000000000000000000000000000000613775565b610e93613112886119f1565b613544915060203d60201161317057613162818361034e565b386134bd565b610e93613183876119f1565b8092949795506131c891935061117c565b613580915060203d60201161317057613162818361034e565b386133cd565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527f01ffc9a7000000000000000000000000000000000000000000000000000000006024820152602481526135ea60448261034e565b5191617530fa6000513d8261360b575b5081613604575090565b9050151590565b602011159150386135fa565b6000602091604051838101907f01ffc9a70000000000000000000000000000000000000000000000000000000082527fffffffff000000000000000000000000000000000000000000000000000000006024820152602481526135ea60448261034e565b6000906020926040517fffffffff00000000000000000000000000000000000000000000000000000000858201927f01ffc9a7000000000000000000000000000000000000000000000000000000008452166024820152602481526135ea60448261034e565b7fffffffff00000000000000000000000000000000000000000000000000000000602082519201517fffffffffffffffff000000000000000000000000000000000000000000000000604051937ff3567d1800000000000000000000000000000000000000000000000000000000602086015260c01b16602484015260e01b16602c820152601081526110ee60308261034e565b916020916000916040519073ffffffffffffffffffffffffffffffffffffffff858301937fa9059cbb0000000000000000000000000000000000000000000000000000000085521660248301526044820152604481526137d660648261034e565b519082855af1156117c9576000513d61384f575073ffffffffffffffffffffffffffffffffffffffff81163b155b61380b5750565b73ffffffffffffffffffffffffffffffffffffffff907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b6001141561380456fea164736f6c634300081a000a",
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetLockReleasePoolAddress(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getLockReleasePoolAddress", remoteChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetLockReleasePoolAddress(remoteChainSelector uint64) (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetLockReleasePoolAddress(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetLockReleasePoolAddress(remoteChainSelector uint64) (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetLockReleasePoolAddress(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) UpdateLockReleasePoolAddresses(opts *bind.TransactOpts, remoteChainSelectors []uint64, lockReleasePools []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "updateLockReleasePoolAddresses", remoteChainSelectors, lockReleasePools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) UpdateLockReleasePoolAddresses(remoteChainSelectors []uint64, lockReleasePools []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdateLockReleasePoolAddresses(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, lockReleasePools)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) UpdateLockReleasePoolAddresses(remoteChainSelectors []uint64, lockReleasePools []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.UpdateLockReleasePoolAddresses(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, lockReleasePools)
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

type USDCTokenPoolProxyLockReleasePoolUpdatedIterator struct {
	Event *USDCTokenPoolProxyLockReleasePoolUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyLockReleasePoolUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyLockReleasePoolUpdated)
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
		it.Event = new(USDCTokenPoolProxyLockReleasePoolUpdated)
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

func (it *USDCTokenPoolProxyLockReleasePoolUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyLockReleasePoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyLockReleasePoolUpdated struct {
	RemoteChainSelector uint64
	LockReleasePool     common.Address
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterLockReleasePoolUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockReleasePoolUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "LockReleasePoolUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyLockReleasePoolUpdatedIterator{contract: _USDCTokenPoolProxy.contract, event: "LockReleasePoolUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchLockReleasePoolUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockReleasePoolUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "LockReleasePoolUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyLockReleasePoolUpdated)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockReleasePoolUpdated", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseLockReleasePoolUpdated(log types.Log) (*USDCTokenPoolProxyLockReleasePoolUpdated, error) {
	event := new(USDCTokenPoolProxyLockReleasePoolUpdated)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockReleasePoolUpdated", log); err != nil {
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

func (USDCTokenPoolProxyLockOrBurnMechanismUpdated) Topic() common.Hash {
	return common.HexToHash("0x2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc94471")
}

func (USDCTokenPoolProxyLockReleasePoolUpdated) Topic() common.Hash {
	return common.HexToHash("0x2a4ec2a96b51064b74fa8f2157f98cf8fb9fd4dcef4e9fdf3c44c0d74d0e8264")
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

	GetLockOrBurnMechanism(opts *bind.CallOpts, remoteChainSelector uint64) (uint8, error)

	GetLockReleasePoolAddress(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

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

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateLockOrBurnMechanisms(opts *bind.TransactOpts, remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error)

	UpdateLockReleasePoolAddresses(opts *bind.TransactOpts, remoteChainSelectors []uint64, lockReleasePools []common.Address) (*types.Transaction, error)

	UpdatePoolAddresses(opts *bind.TransactOpts, pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error)

	FilterLockOrBurnMechanismUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator, error)

	WatchLockOrBurnMechanismUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockOrBurnMechanismUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockOrBurnMechanismUpdated(log types.Log) (*USDCTokenPoolProxyLockOrBurnMechanismUpdated, error)

	FilterLockReleasePoolUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockReleasePoolUpdatedIterator, error)

	WatchLockReleasePoolUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockReleasePoolUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockReleasePoolUpdated(log types.Log) (*USDCTokenPoolProxyLockReleasePoolUpdated, error)

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
