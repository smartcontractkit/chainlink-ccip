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

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type RateLimiterTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

type TokenPoolChainUpdate struct {
	RemoteChainSelector       uint64
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type USDCTokenPoolProxyPoolAddresses struct {
	LegacyCctpV1Pool common.Address
	CctpV1Pool       common.Address
	CctpV2Pool       common.Address
	LockReleasePool  common.Address
}

var USDCTokenPoolProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"legacyCctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346104c257615ed0803803809161001d82856104df565b833981019080820361010081126104c25781516001600160a01b038116939091908483036104c25761005160208501610502565b60408501516001600160401b0381116104c25785019180601f840112156104c2578251926001600160401b0384116104ac578360051b90602082019461009a60405196876104df565b85526020808601928201019283116104c257602001905b8282106104c75750505060806100c960608701610502565b93607f1901126104c25760405192608084016001600160401b038111858210176104ac576040526100fc60808701610502565b845261010a60a08701610502565b946020850195865261013160e061012360c08a01610502565b9860408801998a5201610502565b9760608601988952331561049b57600180546001600160a01b031916331790558015801561048a575b8015610479575b6104685760049260209260805260c0526040519283809263313ce56760e01b82525afa8091600091610425575b5090610401575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526102d8575b5081516001600160a01b03161580156102c6575b80156102b4575b6102a35751600b80546001600160a01b03199081166001600160a01b03938416179091559151600c805484169183169190911790559151600d805483169184169190911790559151600e8054909316911617905560405161581990816106b782396080518181816105de01528181610683015281816113cd015281816149df0152614e90015260a051816106eb015260c0518181816120c801528181613b280152614479015260e051818181610a67015281816121220152614cec0152f35b63ccb986cb60e01b60005260046000fd5b5083516001600160a01b0316156101e4565b5082516001600160a01b0316156101dd565b919290602092604051926102ec85856104df565b60008452600036813760e051156103f05760005b8451811015610367576001906001600160a01b0361031e8288610516565b51168761032a82610558565b610337575b505001610300565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388761032f565b50919490925060005b83518110156103e4576001906001600160a01b0361038e8287610516565b511680156103de57866103a082610656565b6103ae575b50505b01610370565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138866103a5565b506103a8565b509250929050386101c9565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610195576332ad3e0760e11b600052600660045260245260446000fd5b6020813d602011610460575b8161043e602093836104df565b8101031261045c57519060ff8216820361045957503861018e565b80fd5b5080fd5b3d9150610431565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610161565b506001600160a01b0384161561015a565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b602080916104d484610502565b8152019101906100b1565b601f909101601f19168101906001600160401b038211908210176104ac57604052565b51906001600160a01b03821682036104c257565b805182101561052a5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561052a5760005260206000200190600090565b600081815260036020526040902054801561064f57600019810181811161063957600254600019810191908211610639578181036105e8575b50505060025480156105d257600019016105ac816002610540565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6106216105f961060a936002610540565b90549060031b1c9283926002610540565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610591565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146106b057600254680100000000000000008110156104ac5761069761060a8260018594016002556002610540565b9055600254906000526003602052604060002055600190565b5060009056fe6080604052600436101561001257600080fd5b60003560e01c806301ffc9a714610237578063181f5a771461023257806321df0da71461022d578063240028e81461022857806324f65ee714610223578063309292ac1461021e57806339077537146102195780634c5ef0ed1461021457806354c8a4f31461020f57806362ddd3c41461020a578063673a2a1f146102055780636d3d1a581461020057806379ba5097146101fb5780637d54534e146101f65780638926f54f146101f15780638da5cb5b146101ec578063962d4020146101e75780639a4575b9146101e2578063a42a7b8b146101dd578063a7cd63b7146101d8578063aa86a754146101d3578063acfecf91146101ce578063af58d59f146101c9578063b0f479a1146101c4578063b7946580146101bf578063c0d78655146101ba578063c4bffe2b146101b5578063c75eea9c146101b0578063cf7401f3146101ab578063db4c2aef146101a6578063dc0bd971146101a1578063e0351e131461019c578063e8a1da17146101975763f2fde38b1461019257600080fd5b612566565b612147565b6120ec565b61207d565b611f60565b611e44565b611ce1565b611c30565b611af0565b611a7c565b611a2a565b611960565b61182c565b6117c7565b6116d2565b6115b2565b61129c565b61108c565b611009565b610fac565b610efd565b610e14565b610dc2565b610d25565b610ca2565b610a35565b61092d565b610839565b61070f565b6106b3565b61062b565b610593565b610512565b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610327576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361032757807faff2afbf00000000000000000000000000000000000000000000000000000000602092149081156102fd575b81156102d3575b506040519015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386102c8565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506102c1565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6080810190811067ffffffffffffffff82111761037757604052565b61032c565b6060810190811067ffffffffffffffff82111761037757604052565b6020810190811067ffffffffffffffff82111761037757604052565b6040810190811067ffffffffffffffff82111761037757604052565b60a0810190811067ffffffffffffffff82111761037757604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761037757604052565b6040519061043c60a0836103ec565b565b6040519061043c610100836103ec565b6040519061043c610140836103ec565b67ffffffffffffffff811161037757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106104ab5750506000910152565b818101518382015260200161049b565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936104f781518092818752878088019101610498565b0116010190565b90602061050f9281815201906104bb565b90565b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275761058f604080519061055381836103ec565b601c82527f55534443546f6b656e506f6f6c50726f787920312e362e332d646576000000006020830152519182916020835260208301906104bb565b0390f35b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b73ffffffffffffffffffffffffffffffffffffffff81160361032757565b359061043c82610602565b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760206106a960043561066b81610602565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691161490565b6040519015158152f35b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103275760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760006107486139fb565b73ffffffffffffffffffffffffffffffffffffffff60243561076981610602565b16158015610810575b80156107e1575b6107b957610785612664565b7f67d92722109d4170cee5a282ae6387dbf3fba5c7783912975743d4e51ab25aa8604051806107b381612792565b0390a180f35b807fccb986cb0000000000000000000000000000000000000000000000000000000060049252fd5b5061080a6064356107f181610602565b73ffffffffffffffffffffffffffffffffffffffff1690565b15610779565b5073ffffffffffffffffffffffffffffffffffffffff60443561083281610602565b1615610772565b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760043567ffffffffffffffff8111610327576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610327576108b5602091600401612cdc565b60405190518152f35b67ffffffffffffffff81160361032757565b359061043c826108be565b9291926108e78261045e565b916108f560405193846103ec565b829481845281830111610327578281602093846000960137010152565b9080601f830112156103275781602061050f933591016108db565b346103275760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757600435610968816108be565b60243567ffffffffffffffff81116103275760209161098e6106a9923690600401610912565b90612fb8565b9181601f840112156103275782359167ffffffffffffffff8311610327576020808501948460051b01011161032757565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126103275760043567ffffffffffffffff81116103275781610a0e91600401610994565b929092916024359067ffffffffffffffff821161032757610a3191600401610994565b9091565b3461032757610a5d610a65610a49366109c5565b9491610a569391936139fb565b369161300d565b92369161300d565b7f000000000000000000000000000000000000000000000000000000000000000015610bfd5760005b8251811015610b3e5780610ac1610aa76001938661323e565b5173ffffffffffffffffffffffffffffffffffffffff1690565b610ae8610ae373ffffffffffffffffffffffffffffffffffffffff83166107f1565b6150cf565b610af4575b5001610a8e565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a138610aed565b5060005b8151811015610bfb5780610b5b610aa76001938561323e565b73ffffffffffffffffffffffffffffffffffffffff811615610bf557610b9e610b9973ffffffffffffffffffffffffffffffffffffffff83166107f1565b61537e565b610bab575b505b01610b42565b60405173ffffffffffffffffffffffffffffffffffffffff9190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a183610ba3565b50610ba5565b005b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82011261032757600435610c5d816108be565b9160243567ffffffffffffffff811161032757826023820112156103275780600401359267ffffffffffffffff84116103275760248483010111610327576024019190565b3461032757610cb036610c27565b610cbb9291926139fb565b67ffffffffffffffff8216610cdd816000526006602052604060002054151590565b15610cf85750610bfb92610cf29136916108db565b9061423a565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760006060604051610d648161035b565b82815282602082015282604082015201526080610d7f612846565b73ffffffffffffffffffffffffffffffffffffffff6060604051928281511684528260208201511660208501528260408201511660408501520151166060820152f35b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760005473ffffffffffffffffffffffffffffffffffffffff81163303610ed3577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610327577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff600435610f7081610602565b610f786139fb565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760206106a967ffffffffffffffff600435610ff5816108be565b166000526006602052604060002054151590565b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b9181601f840112156103275782359167ffffffffffffffff8311610327576020808501946060850201011161032757565b346103275760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760043567ffffffffffffffff8111610327576110db903690600401610994565b9060243567ffffffffffffffff8111610327576110fc90369060040161105b565b9060443567ffffffffffffffff81116103275761111d90369060040161105b565b61113f6107f160095473ffffffffffffffffffffffffffffffffffffffff1690565b3314158061121e575b6111ec578386148015906111e2575b6111b85760005b86811061116757005b806111b261118061117b6001948b8b6130b3565b6130c8565b61118b8389896130d2565b6111ac6111a461119c86898b6130d2565b923690611dfb565b913690611dfb565b916142ff565b0161115e565b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b5080861415611157565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6000fd5b506112416107f160015473ffffffffffffffffffffffffffffffffffffffff1690565b331415611148565b9061050f91602081526020611269835160408385015260608401906104bb565b9201519060407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526104bb565b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760043567ffffffffffffffff811161032757806004019060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8236030112610327576113166130e2565b506113208261442f565b61135361134c611332602484016130c8565b67ffffffffffffffff16600052600a602052604060002090565b5460ff1690565b9161135d83611793565b8215611507576000916114279173ffffffffffffffffffffffffffffffffffffffff94611388612846565b90859061139481611793565b60028114871461147357505060646113c660406113f193015173ffffffffffffffffffffffffffffffffffffffff1690565b935b0135837f0000000000000000000000000000000000000000000000000000000000000000614505565b836040519586809581947f9a4575b9000000000000000000000000000000000000000000000000000000008352600483016131b6565b0393165af1801561146e5761058f9160009161144b575b5060405191829182611249565b61146891503d806000833e61146081836103ec565b81019061313d565b3861143e565b612b99565b61147c81611793565b6001811487146114b457505060646114ae60206113f193015173ffffffffffffffffffffffffffffffffffffffff1690565b936113c8565b806114c460039296949396611793565b146114d6575b5060646113f1916113c8565b6113f19193506114ff6060606492015173ffffffffffffffffffffffffffffffffffffffff1690565b9391506114ca565b61121a837f31603b12000000000000000000000000000000000000000000000000000000006000526117a2565b602081016020825282518091526040820191602060408360051b8301019401926000915b83831061156757505050505090565b90919293946020806115a3837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0866001960301875289516104bb565b97019301930191939290611558565b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275767ffffffffffffffff6004356115f6816108be565b16600052600760205261160f6005604060002001614fd9565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061165561163f84612ff5565b9361164d60405195866103ec565b808552612ff5565b0160005b8181106116c157505060005b81518110156116b357806116976116926116816001948661323e565b516000526008602052604060002090565b6132a5565b6116a1828661323e565b526116ac818561323e565b5001611665565b6040518061058f8582611534565b806060602080938701015201611659565b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757611709614f43565b60405180916020820160208352815180915260206040840192019060005b818110611735575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611727565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6004111561179d57565b611764565b90602491600481101561179d57600452565b91906020830192600482101561179d5752565b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275767ffffffffffffffff60043561180b816108be565b16600052600a60205261058f60ff60406000205416604051918291826117b4565b346103275761183a36610c27565b6118459291926139fb565b67ffffffffffffffff82169161186f61186b846000526006602052604060002054151590565b1590565b611929576118b261186b60056118998467ffffffffffffffff166000526007602052604060002090565b016118a53686896108db565b6020815191012090615273565b6118ee57507f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691926118e960405192839283613386565b0390a2005b61192584926040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501613365565b0390fd5b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275767ffffffffffffffff6004356119a4816108be565b6119ac613397565b5016600052600760205261058f6119d16119cc60026040600020016133c2565b614644565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275767ffffffffffffffff600435611ac0816108be565b16600052600760205261058f611adc60046040600020016132a5565b6040519182916020835260208301906104bb565b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275773ffffffffffffffffffffffffffffffffffffffff600435611b4081610602565b611b486139fb565b168015611bc25760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a1005b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b602060408183019282815284518094520192019060005b818110611c105750505090565b825167ffffffffffffffff16845260209384019390920191600101611c03565b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757611c67614f8e565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611c9761163f84612ff5565b0136602084013760005b8151811015611cd3578067ffffffffffffffff611cc06001938561323e565b5116611ccc828661323e565b5201611ca1565b6040518061058f8582611bec565b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275767ffffffffffffffff600435611d25816108be565b611d2d613397565b5016600052600760205261058f6119d16119cc60406000206133c2565b8015150361032757565b35906fffffffffffffffffffffffffffffffff8216820361032757565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c60609101126103275760405190611da88261037c565b81608435611db581611d4a565b815260a4356fffffffffffffffffffffffffffffffff8116810361032757602082015260c435906fffffffffffffffffffffffffffffffff821682036103275760400152565b919082606091031261032757604051611e138161037c565b6040611e3f8183958035611e2681611d4a565b8552611e3460208201611d54565b602086015201611d54565b910152565b346103275760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757600435611e7f816108be565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc36011261032757604051611eb58161037c565b602435611ec181611d4a565b81526044356fffffffffffffffffffffffffffffffff811681036103275760208201526064356fffffffffffffffffffffffffffffffff81168103610327576040820152611f0e36611d71565b9073ffffffffffffffffffffffffffffffffffffffff6009541633141580611f3e575b6111ec57610bfb926142ff565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415611f31565b3461032757611f6e366109c5565b611f766139fb565b8083036111b85760005b838110611f8957005b611fa261186b611f9d61117b84888a6130b3565b613063565b6120365780611fd5611fbf611fba60019486886130b3565b61341c565b611fd061133261117b858a8c6130b3565b613429565b611fe361117b8287896130b3565b7f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc9447167ffffffffffffffff61201b611fba85888a6130b3565b9261202d6040519283921694826117b4565b0390a201611f80565b61204761117b61121a9286886130b3565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261032757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103275760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461032757612155366109c5565b9190926121606139fb565b6000915b8083106124125750505060009163ffffffff4216925b82811061218357005b6121966121918285856135dd565b61369c565b90606082016121a58151614721565b60808301936121b48551614721565b604084019081515115611bc2576121ee61186b6121e96121dc885167ffffffffffffffff1690565b67ffffffffffffffff1690565b61540f565b6123c75761232761222761220d879a999a5167ffffffffffffffff1690565b67ffffffffffffffff166000526007602052604060002090565b6122ea896122e487516122cb61225060408301516fffffffffffffffffffffffffffffffff1690565b916122b261227b61227460208401516fffffffffffffffffffffffffffffffff1690565b9251151590565b6122a961228661042d565b6fffffffffffffffffffffffffffffffff851681529763ffffffff166020890152565b15156040870152565b6fffffffffffffffffffffffffffffffff166060850152565b6fffffffffffffffffffffffffffffffff166080830152565b82613733565b61231c896123138a516122cb61225060408301516fffffffffffffffffffffffffffffffff1690565b60028301613733565b60048451910161383f565b602085019660005b8851805182101561236a579061236460019261235d836123578c5167ffffffffffffffff1690565b9261323e565b519061423a565b0161232f565b505097965094906123be7f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293926123ab6001975167ffffffffffffffff1690565b9251935190519060405194859485613966565b0390a10161217a565b61121a6123dc865167ffffffffffffffff1690565b7f1d5ad3c50000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff16600452602490565b90919261242361117b8584866130b3565b9461243a61186b67ffffffffffffffff88166151ac565b61252e5761246760056124618867ffffffffffffffff166000526007602052604060002090565b01614fd9565b9360005b85518110156124b3576001906124ac600561249a8b67ffffffffffffffff166000526007602052604060002090565b016124a5838a61323e565b5190615273565b500161246b565b509350937f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916612520600193976125056125008267ffffffffffffffff166000526007602052604060002090565b61352e565b60405167ffffffffffffffff90911681529081906020820190565b0390a1019190939293612164565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff861660045260246000fd5b346103275760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103275773ffffffffffffffffffffffffffffffffffffffff6004356125b681610602565b6125be6139fb565b1633811461263057807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b3561050f81610602565b73ffffffffffffffffffffffffffffffffffffffff60043561268581610602565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600b541617600b5573ffffffffffffffffffffffffffffffffffffffff6024356126d081610602565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c5573ffffffffffffffffffffffffffffffffffffffff60443561271b81610602565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600d541617600d5573ffffffffffffffffffffffffffffffffffffffff60643561276681610602565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600e541617600e55565b90608082019173ffffffffffffffffffffffffffffffffffffffff6004356127b981610602565b16815273ffffffffffffffffffffffffffffffffffffffff6024356127dd81610602565b16602082015273ffffffffffffffffffffffffffffffffffffffff60443561280481610602565b166040820152606073ffffffffffffffffffffffffffffffffffffffff60643561282d81610602565b16910152565b6040519061284082610398565b60008252565b604051906128538261035b565b8173ffffffffffffffffffffffffffffffffffffffff600b5416815273ffffffffffffffffffffffffffffffffffffffff600c5416602082015273ffffffffffffffffffffffffffffffffffffffff600d54166040820152606073ffffffffffffffffffffffffffffffffffffffff600e5416910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610327570180359067ffffffffffffffff82116103275760200191813603831361032757565b919091357fffffffff0000000000000000000000000000000000000000000000000000000081169260048110612950575050565b7fffffffff00000000000000000000000000000000000000000000000000000000929350829060040360031b1b161690565b90816020910312610327576040519061299a82610398565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561032757016020813591019167ffffffffffffffff821161032757813603831361032757565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9061050f9160208152612b68612b5d612b20612a61612a4e86806129a0565b61010060208801526101208701916129f0565b612a81612a70602088016108d0565b67ffffffffffffffff166040870152565b612aad612a9060408801610620565b73ffffffffffffffffffffffffffffffffffffffff166060870152565b60608601356080860152612ae3612ac660808801610620565b73ffffffffffffffffffffffffffffffffffffffff1660a0870152565b612af060a08701876129a0565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c08801526129f0565b612b2d60c08601866129a0565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e08701526129f0565b9260e08101906129a0565b916101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101526129f0565b6040513d6000823e3d90fd5b9061050f916020815260e0612c9a612c67612bce855161010060208701526101208601906104bb565b602086015167ffffffffffffffff166040860152604086015173ffffffffffffffffffffffffffffffffffffffff16606086015260608601516080860152612c33608087015160a087019073ffffffffffffffffffffffffffffffffffffffff169052565b60a08601517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08683030160c08701526104bb565b60c08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085830301848601526104bb565b920151906101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828503019101526104bb565b906004116103275790600490565b612ce4612833565b50612cf3606082013582613a5b565b612cfb612846565b60c082017ffa7c07de000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000612d54612d4e84876128cb565b9061291c565b1614612f8857612d66604091846128cb565b905014612ee357612d7a60e08301836128cb565b9060048210612eb457612d4e612d9992612d9392612cce565b60e01c90565b63ffffffff811680612e4a575050612e0b600092612dd56107f16107f160208096015173ffffffffffffffffffffffffffffffffffffffff1690565b906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301612a2f565b03925af190811561146e57600091612e21575090565b61050f915060203d602011612e43575b612e3b81836103ec565b810190612982565b503d612e31565b600103612e815750612e0b600092612dd56107f16107f16040602096015173ffffffffffffffffffffffffffffffffffffffff1690565b7f68d2f8d60000000000000000000000000000000000000000000000000000000060005263ffffffff1660045260246000fd5b7f758b22cc00000000000000000000000000000000000000000000000000000000600052600482905260246000fd5b81602091612ef2600094613c97565b8414612f205790612dd56107f16107f1612e0b945173ffffffffffffffffffffffffffffffffffffffff1690565b612f526107f16107f185612f36612e0b96613fcf565b94015173ffffffffffffffffffffffffffffffffffffffff1690565b906040519485809481937f3907753700000000000000000000000000000000000000000000000000000000835260048301612ba5565b50612e0b600092612dd56107f16107f16060602096015173ffffffffffffffffffffffffffffffffffffffff1690565b9067ffffffffffffffff61050f92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116103775760051b60200190565b92919061301981612ff5565b9361302760405195866103ec565b602085838152019160051b810192831161032757905b82821061304957505050565b60208091833561305881610602565b81520191019061303d565b67ffffffffffffffff61050f91166000526006602052604060002054151590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156130c35760051b0190565b613084565b3561050f816108be565b91908110156130c3576060020190565b604051906130ef826103b4565b60606020838281520152565b81601f820112156103275780516131118161045e565b9261311f60405194856103ec565b818452602082840101116103275761050f9160208085019101610498565b6020818303126103275780519067ffffffffffffffff821161032757016040818303126103275760405191613171836103b4565b815167ffffffffffffffff8111610327578161318e9184016130fb565b8352602082015167ffffffffffffffff8111610327576131ae92016130fb565b602082015290565b6020815260a073ffffffffffffffffffffffffffffffffffffffff60806131f06131e086806129a0565b85602088015260c08701916129f0565b9467ffffffffffffffff6020820135613208816108be565b16604086015282604082013561321d81610602565b166060860152606081013582860152013561323781610602565b1691015290565b80518210156130c35760209160051b010190565b90600182811c9216801561329b575b602083101461326c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613261565b90604051918260008254926132b984613252565b808452936001811690811561332557506001146132de575b5061043c925003836103ec565b90506000929192526020600020906000915b81831061330957505090602061043c92820101386132d1565b60209193508060019154838589010152019101909184926132f0565b6020935061043c9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386132d1565b60409067ffffffffffffffff61050f959316815281602082015201916129f0565b91602061050f9381815201916129f0565b604051906133a4826103d0565b60006080838281528260208201528260408201528260608201520152565b906040516133cf816103d0565b60806fffffffffffffffffffffffffffffffff6001839560ff8154848116875263ffffffff81871c16602088015260a01c1615156040860152015481808216166060850152821c16910152565b3560048110156103275790565b90600481101561179d5760ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055565b91613498918354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055565b8181106134a7575050565b6000815560010161349c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818102929181159184041417156134f557565b6134b3565b805490600081558161350a575050565b6000526020600020908101905b818110613522575050565b60008155600101613517565b600561043c91600081556000600182015560006002820155600060038201556004810161355b8154613252565b908161356a575b5050016134fa565b81601f600093116001146135825750555b3880613562565b8183526020832061359d91601f01861c81019060010161349c565b808252602082209081548360011b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8560031b1c19161790555561357b565b91908110156130c35760051b810135907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee181360301821215610327570190565b9080601f8301121561032757813561363481612ff5565b9261364260405194856103ec565b81845260208085019260051b820101918383116103275760208201905b83821061366e57505050505090565b813567ffffffffffffffff81116103275760209161369187848094880101610912565b81520191019061365f565b6101208136031261032757604051906136b4826103d0565b6136bd816108d0565b8252602081013567ffffffffffffffff8111610327576136e0903690830161361d565b602083015260408101359067ffffffffffffffff82116103275761370a61372b9236908301610912565b604084015261371c3660608301611dfb565b606084015260c0369101611dfb565b608082015290565b8151815460208401516040850151608091821b73ffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffff000000000000000000000000000000000000000000909416939093179290921791151560a01b74ff000000000000000000000000000000000000000016919091178355606084015193810151901b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016921691909117600190910155565b9190601f811161380957505050565b61043c926000526020600020906020601f840160051c83019310613835575b601f0160051c019061349c565b9091508190613828565b919091825167ffffffffffffffff811161037757613867816138618454613252565b846137fa565b6020601f82116001146138c15781906134989394956000926138b6575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b015190503880613884565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216906138f484600052602060002090565b9160005b81811061394e57509583600195969710613917575b505050811b019055565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538808061390d565b9192602060018192868b0151815501940192016138f8565b6139ca61399561043c9597969467ffffffffffffffff60a09516845261010060208501526101008401906104bb565b9660408301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b01906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b73ffffffffffffffffffffffffffffffffffffffff600154163303613a1c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90816020910312610327575161050f81611d4a565b60808101613a6e61186b61066b8361265a565b613c4957506020810190613b0f6020613ab4613a8c6121dc866130c8565b60801b7fffffffffffffffffffffffffffffffff000000000000000000000000000000001690565b6040517f2cbc26bb0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116600482015291829081906024820190565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561146e57600091613c1a575b50613bf057613b6f613b6a836130c8565b614862565b613b78826130c8565b90613b9861186b60a083019361098e613b9186866128cb565b36916108db565b613bb057505090613bab61043c926130c8565b614986565b613bba92506128cb565b906119256040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260048401613386565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b613c3c915060203d602011613c42575b613c3481836103ec565b810190613a46565b38613b59565b503d613c2a565b613c5561121a9161265a565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060005273ffffffffffffffffffffffffffffffffffffffff16600452602490565b73ffffffffffffffffffffffffffffffffffffffff600b5416908115613d6c578060e0613cc59201906128cb565b8101906020818303126103275780359067ffffffffffffffff821161032757016040818303126103275760405191613cfc836103b4565b813567ffffffffffffffff81116103275781613d19918401610912565b8352602082013567ffffffffffffffff811161032757613d6893613d476107f1936074936107f19601610912565b602082015251015173ffffffffffffffffffffffffffffffffffffffff1690565b1490565b5050600090565b60405190610100820182811067ffffffffffffffff82111761037757604052606060e0838281526000602082015260006040820152600083820152600060808201528260a08201528260c08201520152565b9190916101008184031261032757613ddb61043e565b92813567ffffffffffffffff81116103275781613df9918401610912565b8452613e07602083016108d0565b6020850152613e1860408301610620565b604085015260608201356060850152613e3360808301610620565b608085015260a082013567ffffffffffffffff81116103275781613e58918401610912565b60a085015260c082013567ffffffffffffffff81116103275781613e7d918401610912565b60c085015260e082013567ffffffffffffffff811161032757613ea09201610912565b60e0830152565b63ffffffff81160361032757565b90816020910312610327575161050f81613ea7565b90816020910312610327575161050f81610602565b91908260409103126103275760208235613ef8816108be565b92013561050f81613ea7565b90600382101561179d5752565b61043c9092919261012080610140830195613f3684825167ffffffffffffffff169052565b60208181015163ffffffff1690850152613f5860408201516040860190613f04565b60608101516060850152613f796080820151608086019063ffffffff169052565b60a081015160a0850152613faa60c082015160c086019073ffffffffffffffffffffffffffffffffffffffff169052565b60e081015160e0850152610100810151610100850152015191019063ffffffff169052565b613fd7613d73565b50613fe23682613dc5565b906140086107f16107f1600c5473ffffffffffffffffffffffffffffffffffffffff1690565b90604051917f6b716b0d000000000000000000000000000000000000000000000000000000008352602083600481845afa90811561146e576004936000926141e7575b50602090604051948580927f98db96430000000000000000000000000000000000000000000000000000000082525afa92831561146e5761417a936141a69373ffffffffffffffffffffffffffffffffffffffff926000926141ae575b50614133816140c86140c060c06141559501836128cb565b810190613edf565b929096606083013561411a6140f160806140ea6107f16107f160408a0161265a565b960161265a565b9561410d6140fd61044e565b67ffffffffffffffff909c168c52565b63ffffffff1660208b0152565b600160408a0152606089015263ffffffff166080880152565b60a086015273ffffffffffffffffffffffffffffffffffffffff1660c0850152565b1660e08201526000610100820152600061012082015260405192839160208301613f11565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103ec565b60c082015290565b6141559192506141d86141339160203d6020116141e0575b6141d081836103ec565b810190613eca565b9291506140a8565b503d6141c6565b602091925061420b90823d8411614212575b61420381836103ec565b810190613eb5565b919061404b565b503d6141f9565b60409067ffffffffffffffff61050f949316815281602082015201906104bb565b90805115611bc2578051602082012067ffffffffffffffff83169283600052600760205261426f826005604060002001615465565b156142c85750816142b77f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea936142b26142c3946000526008602052604060002090565b61383f565b604051918291826104fe565b0390a2565b90506119256040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401614219565b67ffffffffffffffff16600081815260066020526040902054909291901561440157916143fe60e0926143ca856143567f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97614721565b84600052600760205261436d816040600020614a37565b61437683614721565b846000526007602052614390836002604060002001614a37565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6080810161444261186b61066b8361265a565b613c49575060208101906144606020613ab4613a8c6121dc866130c8565b038173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561146e576000916144e6575b50613bf05760606144dd61043c936144cc6144c76040860161265a565b614cea565b61117b6144d8826130c8565b614d6c565b91013590614e3a565b6144ff915060203d602011613c4257613c3481836103ec565b386144aa565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff93841660248301526044808301959095529381526145dc939092909161456c6064856103ec565b1660008060409384519561458086886103ec565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15614601573d6145cd6145c48261045e565b945194856103ec565b83523d6000602085013e615748565b8051806145e7575050565b816020806145fc9361043c9501019101613a46565b614eb8565b60609250615748565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116134f557565b919082039182116134f557565b61464c613397565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff82511690602083019163ffffffff83511642034281116134f5576146b0906fffffffffffffffffffffffffffffffff608087015116906134e2565b81018091116134f5576146d66fffffffffffffffffffffffffffffffff92918392615736565b161682524263ffffffff16905290565b61043c9092919260608101936fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8051156147c55760408101516fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff61478561477060208501516fffffffffffffffffffffffffffffffff1690565b6fffffffffffffffffffffffffffffffff1690565b91161161478f5750565b611925906040519182917f8020d124000000000000000000000000000000000000000000000000000000008352600483016146e6565b6fffffffffffffffffffffffffffffffff6147f360408301516fffffffffffffffffffffffffffffffff1690565b161580159061483a575b6148045750565b611925906040519182917fd68af9cc000000000000000000000000000000000000000000000000000000008352600483016146e6565b5061485b61477060208301516fffffffffffffffffffffffffffffffff1690565b15156147fd565b61486e61186b82613063565b61494f5760206148e79161489a6107f160045473ffffffffffffffffffffffffffffffffffffffff1690565b6040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90921660048301523360248301529092839190829081906044820190565b03915afa90811561146e57600091614930575b501561490257565b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b614949915060203d602011613c4257613c3481836103ec565b386148fa565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005267ffffffffffffffff1660045260246000fd5b67ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c91169182600052600760205280614a07600260406000200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916154f3565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252602082019290925290819081016142c3565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991614c16614c22928054614a88614a82614a798363ffffffff9060801c1690565b63ffffffff1690565b42614637565b9081614c27575b5050614bd06001614ab360208601516fffffffffffffffffffffffffffffffff1690565b92614b3e614b016147706fffffffffffffffffffffffffffffffff614ae885546fffffffffffffffffffffffffffffffff1690565b166fffffffffffffffffffffffffffffffff8816615736565b82906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b614b91614b4b8751151590565b82547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016178255565b019182906fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b604083015181546fffffffffffffffffffffffffffffffff1660809190911b7fffffffffffffffffffffffffffffffff0000000000000000000000000000000016179055565b604051918291826146e6565b0390a1565b614770614b01916fffffffffffffffffffffffffffffffff614c9b614ca29582614c9460018a01549282614c8d614c86614c70876fffffffffffffffffffffffffffffffff1690565b996fffffffffffffffffffffffffffffffff1690565b9560801c90565b16906134e2565b9116615354565b9116615736565b80547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161781553880614a8f565b7f0000000000000000000000000000000000000000000000000000000000000000614d125750565b73ffffffffffffffffffffffffffffffffffffffff1680600052600360205260406000205415614d3f5750565b7fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b614d7861186b82613063565b61494f576020614de991614da46107f160045473ffffffffffffffffffffffffffffffffffffffff1690565b60405180809581947fa8d87a3b0000000000000000000000000000000000000000000000000000000083526004830191909167ffffffffffffffff6020820193169052565b03915afa801561146e5773ffffffffffffffffffffffffffffffffffffffff91600091614e1b575b5016330361490257565b614e34915060203d6020116141e0576141d081836103ec565b38614e11565b67ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894491169182600052600760205280614a07604060002073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916154f3565b15614ebf57565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b604051906002548083528260208101600260005260206000209260005b818110614f7557505061043c925003836103ec565b8454835260019485019487945060209093019201614f60565b604051906005548083528260208101600560005260206000209260005b818110614fc057505061043c925003836103ec565b8454835260019485019487945060209093019201614fab565b906040519182815491828252602082019060005260206000209260005b81811061500b57505061043c925003836103ec565b8454835260019485019487945060209093019201614ff6565b80548210156130c35760005260206000200190600090565b805480156150a0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906150718282615024565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260036020526040902054908115613d6c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116134f557600254927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116134f557838360009561516b9503615171575b50505061515a600261503c565b600390600052602052604060002090565b55600190565b61515a61519d916151936151896151a3956002615024565b90549060031b1c90565b9283916002615024565b90613460565b5538808061514d565b600081815260066020526040902054908115613d6c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201908282116134f557600554927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84019384116134f557838360009561516b9503615248575b505050615237600561503c565b600690600052602052604060002090565b61523761519d9161526061518961526a956005615024565b9283916005615024565b5538808061522a565b600181019180600052826020526040600020549283151560001461534b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116134f5578354937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85019485116134f557600095858361516b976153039503615312575b50505061503c565b90600052602052604060002090565b61533261519d916153296151896153429588615024565b92839187615024565b8590600052602052604060002090565b553880806152fb565b50505050600090565b919082018092116134f557565b9261536c91926134e2565b81018091116134f55761050f91615736565b6000818152600360205260409020546154095760025468010000000000000000811015610377576153f06153bb8260018594016002556002615024565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600254906000526003602052604060002055600190565b50600090565b60008181526006602052604090205461540957600554680100000000000000008110156103775761544c6153bb8260018594016005556005615024565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054613d6c578054906801000000000000000082101561037757826154a36153bb846001809601855584615024565b905580549260005201602052604060002055600190565b81156154c4570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b8054939290919060ff60a086901c1615801561572e575b615727576155296fffffffffffffffffffffffffffffffff8616614770565b9060018401958654615563614a82614a79615556614770856fffffffffffffffffffffffffffffffff1690565b9460801c63ffffffff1690565b80615693575b505083811061564857508282106155c9575061043c93945061558e9161477091614637565b6fffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9061560061121a936155fb6155ec846155e66147708c5460801c90565b93614637565b6155f58361460a565b90615354565b6154ba565b7fd0c8d23a0000000000000000000000000000000000000000000000000000000060005260045260245273ffffffffffffffffffffffffffffffffffffffff16604452606490565b7f1a76572a00000000000000000000000000000000000000000000000000000000600052600452602483905273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b8285929395116156fd576156ad6147706156b49460801c90565b9185615361565b84547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178555913880615569565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b5050509050565b50811561550a565b9080821015615743575090565b905090565b919290156157c3575081511561575c575090565b3b156157655790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156157d65750805190602001fd5b611925906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352600483016104fe56fea164736f6c634300081a000a",
}

var USDCTokenPoolProxyABI = USDCTokenPoolProxyMetaData.ABI

var USDCTokenPoolProxyBin = USDCTokenPoolProxyMetaData.Bin

func DeployUSDCTokenPoolProxy(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, router common.Address, allowlist []common.Address, rmnProxy common.Address, pools USDCTokenPoolProxyPoolAddresses) (common.Address, *types.Transaction, *USDCTokenPoolProxy, error) {
	parsed, err := USDCTokenPoolProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(USDCTokenPoolProxyBin), backend, token, router, allowlist, rmnProxy, pools)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetAllowList() ([]common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetAllowList(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetAllowList() ([]common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetAllowList(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetAllowListEnabled() (bool, error) {
	return _USDCTokenPoolProxy.Contract.GetAllowListEnabled(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetAllowListEnabled() (bool, error) {
	return _USDCTokenPoolProxy.Contract.GetAllowListEnabled(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPoolProxy.Contract.GetCurrentInboundRateLimiterState(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPoolProxy.Contract.GetCurrentInboundRateLimiterState(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPoolProxy.Contract.GetCurrentOutboundRateLimiterState(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _USDCTokenPoolProxy.Contract.GetCurrentOutboundRateLimiterState(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetRateLimitAdmin() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRateLimitAdmin(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRateLimitAdmin(&_USDCTokenPoolProxy.CallOpts)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetRmnProxy() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRmnProxy(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetRmnProxy() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRmnProxy(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetRouter() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRouter(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetRouter() (common.Address, error) {
	return _USDCTokenPoolProxy.Contract.GetRouter(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetSupportedChains() ([]uint64, error) {
	return _USDCTokenPoolProxy.Contract.GetSupportedChains(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetSupportedChains() ([]uint64, error) {
	return _USDCTokenPoolProxy.Contract.GetSupportedChains(&_USDCTokenPoolProxy.CallOpts)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) GetTokenDecimals() (uint8, error) {
	return _USDCTokenPoolProxy.Contract.GetTokenDecimals(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) GetTokenDecimals() (uint8, error) {
	return _USDCTokenPoolProxy.Contract.GetTokenDecimals(&_USDCTokenPoolProxy.CallOpts)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _USDCTokenPoolProxy.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsRemotePool(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _USDCTokenPoolProxy.Contract.IsRemotePool(&_USDCTokenPoolProxy.CallOpts, remoteChainSelector, remotePoolAddress)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.AddRemotePool(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.AddRemotePool(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ApplyAllowListUpdates(&_USDCTokenPoolProxy.TransactOpts, removes, adds)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ApplyAllowListUpdates(&_USDCTokenPoolProxy.TransactOpts, removes, adds)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ApplyChainUpdates(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ApplyChainUpdates(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.ReleaseOrMint(&_USDCTokenPoolProxy.TransactOpts, releaseOrMintIn)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.RemoveRemotePool(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.RemoveRemotePool(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetChainRateLimiterConfig(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetChainRateLimiterConfig(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetChainRateLimiterConfigs(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetChainRateLimiterConfigs(&_USDCTokenPoolProxy.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetRateLimitAdmin(&_USDCTokenPoolProxy.TransactOpts, rateLimitAdmin)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetRateLimitAdmin(&_USDCTokenPoolProxy.TransactOpts, rateLimitAdmin)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.contract.Transact(opts, "setRouter", newRouter)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxySession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetRouter(&_USDCTokenPoolProxy.TransactOpts, newRouter)
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _USDCTokenPoolProxy.Contract.SetRouter(&_USDCTokenPoolProxy.TransactOpts, newRouter)
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

type USDCTokenPoolProxyAllowListAddIterator struct {
	Event *USDCTokenPoolProxyAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyAllowListAdd)
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
		it.Event = new(USDCTokenPoolProxyAllowListAdd)
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

func (it *USDCTokenPoolProxyAllowListAddIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolProxyAllowListAddIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyAllowListAddIterator{contract: _USDCTokenPoolProxy.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyAllowListAdd)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseAllowListAdd(log types.Log) (*USDCTokenPoolProxyAllowListAdd, error) {
	event := new(USDCTokenPoolProxyAllowListAdd)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyAllowListRemoveIterator struct {
	Event *USDCTokenPoolProxyAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyAllowListRemove)
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
		it.Event = new(USDCTokenPoolProxyAllowListRemove)
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

func (it *USDCTokenPoolProxyAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolProxyAllowListRemoveIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyAllowListRemoveIterator{contract: _USDCTokenPoolProxy.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyAllowListRemove)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseAllowListRemove(log types.Log) (*USDCTokenPoolProxyAllowListRemove, error) {
	event := new(USDCTokenPoolProxyAllowListRemove)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyChainAddedIterator struct {
	Event *USDCTokenPoolProxyChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyChainAdded)
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
		it.Event = new(USDCTokenPoolProxyChainAdded)
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

func (it *USDCTokenPoolProxyChainAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterChainAdded(opts *bind.FilterOpts) (*USDCTokenPoolProxyChainAddedIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyChainAddedIterator{contract: _USDCTokenPoolProxy.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyChainAdded) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyChainAdded)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseChainAdded(log types.Log) (*USDCTokenPoolProxyChainAdded, error) {
	event := new(USDCTokenPoolProxyChainAdded)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyChainConfiguredIterator struct {
	Event *USDCTokenPoolProxyChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyChainConfigured)
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
		it.Event = new(USDCTokenPoolProxyChainConfigured)
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

func (it *USDCTokenPoolProxyChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*USDCTokenPoolProxyChainConfiguredIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyChainConfiguredIterator{contract: _USDCTokenPoolProxy.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyChainConfigured) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyChainConfigured)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseChainConfigured(log types.Log) (*USDCTokenPoolProxyChainConfigured, error) {
	event := new(USDCTokenPoolProxyChainConfigured)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyChainRemovedIterator struct {
	Event *USDCTokenPoolProxyChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyChainRemoved)
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
		it.Event = new(USDCTokenPoolProxyChainRemoved)
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

func (it *USDCTokenPoolProxyChainRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*USDCTokenPoolProxyChainRemovedIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyChainRemovedIterator{contract: _USDCTokenPoolProxy.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyChainRemoved) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyChainRemoved)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseChainRemoved(log types.Log) (*USDCTokenPoolProxyChainRemoved, error) {
	event := new(USDCTokenPoolProxyChainRemoved)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyConfigChangedIterator struct {
	Event *USDCTokenPoolProxyConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyConfigChanged)
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
		it.Event = new(USDCTokenPoolProxyConfigChanged)
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

func (it *USDCTokenPoolProxyConfigChangedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*USDCTokenPoolProxyConfigChangedIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyConfigChangedIterator{contract: _USDCTokenPoolProxy.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyConfigChanged) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyConfigChanged)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseConfigChanged(log types.Log) (*USDCTokenPoolProxyConfigChanged, error) {
	event := new(USDCTokenPoolProxyConfigChanged)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyInboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolProxyInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyInboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolProxyInboundRateLimitConsumed)
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

func (it *USDCTokenPoolProxyInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyInboundRateLimitConsumedIterator{contract: _USDCTokenPoolProxy.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyInboundRateLimitConsumed)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolProxyInboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolProxyInboundRateLimitConsumed)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

type USDCTokenPoolProxyLockedOrBurnedIterator struct {
	Event *USDCTokenPoolProxyLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyLockedOrBurned)
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
		it.Event = new(USDCTokenPoolProxyLockedOrBurned)
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

func (it *USDCTokenPoolProxyLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyLockedOrBurnedIterator{contract: _USDCTokenPoolProxy.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyLockedOrBurned)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseLockedOrBurned(log types.Log) (*USDCTokenPoolProxyLockedOrBurned, error) {
	event := new(USDCTokenPoolProxyLockedOrBurned)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyOutboundRateLimitConsumedIterator struct {
	Event *USDCTokenPoolProxyOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyOutboundRateLimitConsumed)
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
		it.Event = new(USDCTokenPoolProxyOutboundRateLimitConsumed)
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

func (it *USDCTokenPoolProxyOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyOutboundRateLimitConsumedIterator{contract: _USDCTokenPoolProxy.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyOutboundRateLimitConsumed)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolProxyOutboundRateLimitConsumed, error) {
	event := new(USDCTokenPoolProxyOutboundRateLimitConsumed)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

type USDCTokenPoolProxyRateLimitAdminSetIterator struct {
	Event *USDCTokenPoolProxyRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyRateLimitAdminSet)
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
		it.Event = new(USDCTokenPoolProxyRateLimitAdminSet)
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

func (it *USDCTokenPoolProxyRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*USDCTokenPoolProxyRateLimitAdminSetIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyRateLimitAdminSetIterator{contract: _USDCTokenPoolProxy.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyRateLimitAdminSet)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseRateLimitAdminSet(log types.Log) (*USDCTokenPoolProxyRateLimitAdminSet, error) {
	event := new(USDCTokenPoolProxyRateLimitAdminSet)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyReleasedOrMintedIterator struct {
	Event *USDCTokenPoolProxyReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyReleasedOrMinted)
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
		it.Event = new(USDCTokenPoolProxyReleasedOrMinted)
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

func (it *USDCTokenPoolProxyReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyReleasedOrMintedIterator{contract: _USDCTokenPoolProxy.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyReleasedOrMinted)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseReleasedOrMinted(log types.Log) (*USDCTokenPoolProxyReleasedOrMinted, error) {
	event := new(USDCTokenPoolProxyReleasedOrMinted)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyRemotePoolAddedIterator struct {
	Event *USDCTokenPoolProxyRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyRemotePoolAdded)
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
		it.Event = new(USDCTokenPoolProxyRemotePoolAdded)
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

func (it *USDCTokenPoolProxyRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyRemotePoolAddedIterator{contract: _USDCTokenPoolProxy.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyRemotePoolAdded)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseRemotePoolAdded(log types.Log) (*USDCTokenPoolProxyRemotePoolAdded, error) {
	event := new(USDCTokenPoolProxyRemotePoolAdded)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyRemotePoolRemovedIterator struct {
	Event *USDCTokenPoolProxyRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyRemotePoolRemoved)
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
		it.Event = new(USDCTokenPoolProxyRemotePoolRemoved)
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

func (it *USDCTokenPoolProxyRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyRemotePoolRemovedIterator{contract: _USDCTokenPoolProxy.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyRemotePoolRemoved)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseRemotePoolRemoved(log types.Log) (*USDCTokenPoolProxyRemotePoolRemoved, error) {
	event := new(USDCTokenPoolProxyRemotePoolRemoved)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type USDCTokenPoolProxyRouterUpdatedIterator struct {
	Event *USDCTokenPoolProxyRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *USDCTokenPoolProxyRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDCTokenPoolProxyRouterUpdated)
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
		it.Event = new(USDCTokenPoolProxyRouterUpdated)
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

func (it *USDCTokenPoolProxyRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *USDCTokenPoolProxyRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type USDCTokenPoolProxyRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*USDCTokenPoolProxyRouterUpdatedIterator, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &USDCTokenPoolProxyRouterUpdatedIterator{contract: _USDCTokenPoolProxy.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _USDCTokenPoolProxy.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(USDCTokenPoolProxyRouterUpdated)
				if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxyFilterer) ParseRouterUpdated(log types.Log) (*USDCTokenPoolProxyRouterUpdated, error) {
	event := new(USDCTokenPoolProxyRouterUpdated)
	if err := _USDCTokenPoolProxy.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (USDCTokenPoolProxyAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (USDCTokenPoolProxyAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (USDCTokenPoolProxyChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (USDCTokenPoolProxyChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (USDCTokenPoolProxyChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (USDCTokenPoolProxyConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (USDCTokenPoolProxyInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (USDCTokenPoolProxyLockOrBurnMechanismUpdated) Topic() common.Hash {
	return common.HexToHash("0x2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc94471")
}

func (USDCTokenPoolProxyLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (USDCTokenPoolProxyOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
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

func (USDCTokenPoolProxyRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (USDCTokenPoolProxyReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (USDCTokenPoolProxyRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (USDCTokenPoolProxyRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (USDCTokenPoolProxyRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (_USDCTokenPoolProxy *USDCTokenPoolProxy) Address() common.Address {
	return _USDCTokenPoolProxy.address
}

type USDCTokenPoolProxyInterface interface {
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetLockOrBurnMechanism(opts *bind.CallOpts, remoteChainSelector uint64) (uint8, error)

	GetPools(opts *bind.CallOpts) (USDCTokenPoolProxyPoolAddresses, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateLockOrBurnMechanisms(opts *bind.TransactOpts, remoteChainSelectors []uint64, mechanisms []uint8) (*types.Transaction, error)

	UpdatePoolAddresses(opts *bind.TransactOpts, pools USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*USDCTokenPoolProxyAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*USDCTokenPoolProxyAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*USDCTokenPoolProxyAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*USDCTokenPoolProxyAllowListRemove, error)

	FilterChainAdded(opts *bind.FilterOpts) (*USDCTokenPoolProxyChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*USDCTokenPoolProxyChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*USDCTokenPoolProxyChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*USDCTokenPoolProxyChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*USDCTokenPoolProxyChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*USDCTokenPoolProxyChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*USDCTokenPoolProxyConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*USDCTokenPoolProxyConfigChanged, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*USDCTokenPoolProxyInboundRateLimitConsumed, error)

	FilterLockOrBurnMechanismUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockOrBurnMechanismUpdatedIterator, error)

	WatchLockOrBurnMechanismUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockOrBurnMechanismUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockOrBurnMechanismUpdated(log types.Log) (*USDCTokenPoolProxyLockOrBurnMechanismUpdated, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*USDCTokenPoolProxyLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*USDCTokenPoolProxyOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*USDCTokenPoolProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDCTokenPoolProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*USDCTokenPoolProxyOwnershipTransferred, error)

	FilterPoolAddressesUpdated(opts *bind.FilterOpts) (*USDCTokenPoolProxyPoolAddressesUpdatedIterator, error)

	WatchPoolAddressesUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyPoolAddressesUpdated) (event.Subscription, error)

	ParsePoolAddressesUpdated(log types.Log) (*USDCTokenPoolProxyPoolAddressesUpdated, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*USDCTokenPoolProxyRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*USDCTokenPoolProxyRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*USDCTokenPoolProxyReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*USDCTokenPoolProxyRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*USDCTokenPoolProxyRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*USDCTokenPoolProxyRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*USDCTokenPoolProxyRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *USDCTokenPoolProxyRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*USDCTokenPoolProxyRouterUpdated, error)

	Address() common.Address
}
