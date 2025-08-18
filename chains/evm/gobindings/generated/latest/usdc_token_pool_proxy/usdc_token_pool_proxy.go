// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package usdc_token_pool_proxy

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
	CctpV1Pool      common.Address
	CctpV2Pool      common.Address
	LockReleasePool common.Address
}

var USDCTokenPoolProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockOrBurnMechanism\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPools\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateLockOrBurnMechanisms\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"mechanisms\",\"type\":\"uint8[]\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updatePoolAddresses\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockOrBurnMechanismUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"mechanism\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolAddressesUpdated\",\"inputs\":[{\"name\":\"pools\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structUSDCTokenPoolProxy.PoolAddresses\",\"components\":[{\"name\":\"cctpV1Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpV2Pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockReleasePool\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidLockOrBurnMechanism\",\"inputs\":[{\"name\":\"mechanism\",\"type\":\"uint8\",\"internalType\":\"enumUSDCTokenPoolProxy.LockOrBurnMechanism\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMessageVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346104a257615a6f803803809161001d82856104bf565b83398101908082039160e083126104a25781516001600160a01b038116939091908483036104a257610051602085016104e2565b60408501519091906001600160401b0381116104a25785019280601f850112156104a2578351936001600160401b03851161048c578460051b90602082019561009d60405197886104bf565b86526020808701928201019283116104a257602001905b8282106104a75750505060606100cb8187016104e2565b91607f1901126104a25760405193606085016001600160401b0381118682101761048c576040526100fe608087016104e2565b855261011f60c061011160a089016104e2565b9760208801988952016104e2565b9660408601978852331561047b57600180546001600160a01b031916331790558015801561046a575b8015610459575b6104485760049260209260805260c0526040519283809263313ce56760e01b82525afa8091600091610405575b50906103e1575b50600660a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526102be575b5080516001600160a01b03161580156102ac575b801561029a575b6102895751600a80546001600160a01b039283166001600160a01b0319918216179091559151600b80549183169184169190911790559151600c80549190931691161790556040516153d890816106978239608051818181611aac01528181611ca001528181612b4f01528181612bc7015281816134820152613677015260a05181612ada015260c051818181610c0b01528181611b47015261351c015260e051818181610b9b01528181611b8d01526128720152f35b63ccb986cb60e01b60005260046000fd5b5082516001600160a01b0316156101d2565b5081516001600160a01b0316156101cb565b602091604051916102cf84846104bf565b60008352600036813760e051156103d05760005b835181101561034a576001906001600160a01b0361030182876104f6565b51168661030d82610538565b61031a575b5050016102e3565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13886610312565b5091509260005b82518110156103c5576001906001600160a01b0361036f82866104f6565b511680156103bf578561038182610636565b61038f575b50505b01610351565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13885610386565b50610389565b5092915050386101b7565b6335f4a7b360e01b60005260046000fd5b60ff1660068114610183576332ad3e0760e11b600052600660045260245260446000fd5b6020813d602011610440575b8161041e602093836104bf565b8101031261043c57519060ff8216820361043957503861017c565b80fd5b5080fd5b3d9150610411565b630a64406560e11b60005260046000fd5b506001600160a01b0383161561014f565b506001600160a01b03841615610148565b639b15e16f60e01b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b600080fd5b602080916104b4846104e2565b8152019101906100b4565b601f909101601f19168101906001600160401b0382119082101761048c57604052565b51906001600160a01b03821682036104a257565b805182101561050a5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561050a5760005260206000200190600090565b600081815260036020526040902054801561062f57600019810181811161061957600254600019810191908211610619578181036105c8575b50505060025480156105b2576000190161058c816002610520565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6106016105d96105ea936002610520565b90549060031b1c9283926002610520565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610571565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600360205260406000205415600014610690576002546801000000000000000081101561048c576106776105ea8260018594016002556002610520565b9055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714612c6857508063181f5a7714612beb57806321df0da714612b7c578063240028e814612afe57806324f65ee714612aa25780633907753714612a1d5780634c5ef0ed146129ba57806354c8a4f31461284057806362ddd3c4146127bd578063673a2a1f146127345780636d3d1a58146126e257806379ba5097146125f95780637d54534e1461254e5780638926f54f146124ff5780638da5cb5b146124ad578063962d4020146123635780639a4575b914611a00578063a42a7b8b14611874578063a7cd63b7146117a2578063aa86a75414611741578063acfecf911461161c578063af58d59f146115b4578063b0f479a114611562578063b7946580146114f2578063c0d78655146113fa578063c4bffe2b146112ac578063c527a6c21461105f578063c75eea9c14610f98578063cf7401f314610dd5578063db4c2aef14610c2f578063dc0bd97114610bc0578063e0351e1314610b65578063e8a1da171461028a5763f2fde38b1461019557600080fd5b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855773ffffffffffffffffffffffffffffffffffffffff6101e1612ef2565b6101e9614350565b1633811461025b57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346102855761029836612fe8565b9190926102a3614350565b6000905b8282106109bc5750505060009063ffffffff4216907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee184360301925b818110156109ba576000918160051b860135858112156109b657860190610120823603126109b6576040519561031887612dbc565b61032183612f4d565b8752602083013567ffffffffffffffff81116109b25783019536601f880112156109b25786359661035188613f89565b9761035f604051998a612e11565b8089526020808a019160051b830101903682116109ae5760208301905b82821061097b575050505060208801968752604084013567ffffffffffffffff8111610977576103af9036908601612f99565b92604089019384526103d96103c7366060880161313c565b9560608b0196875260c036910161313c565b9660808a019788526103eb865161479d565b6103f5885161479d565b8451511561094f5761041167ffffffffffffffff8b5116614fdc565b156109185767ffffffffffffffff8a5116815260076020526040812061055187516fffffffffffffffffffffffffffffffff6040820151169061050c6fffffffffffffffffffffffffffffffff6020830151169151151583608060405161047781612dbc565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b61067789516fffffffffffffffffffffffffffffffff604082015116906106326fffffffffffffffffffffffffffffffff6020830151169151151583608060405161059b81612dbc565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004865191019080519067ffffffffffffffff82116108eb5761069a83546140af565b601f81116108b0575b50602090601f8311600114610811576106f19291859183610806575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b88518051821015610729579061072360019261071c8367ffffffffffffffff8f51169261409b565b519061439b565b016106f4565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29391999750956107f767ffffffffffffffff60019796949851169251935191516107c361078e60405196879687526101006020880152610100870190612eaf565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019391939290926102e3565b015190508f806106bf565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b8181106108985750908460019594939210610861575b505050811b0190556106f4565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558e8080610854565b9293602060018192878601518155019501930161083e565b6108db9084865260208620601f850160051c810191602086106108e1575b601f0160051c01906142df565b8e6106a3565b90915081906108ce565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff81116109aa5760209161099f8392833691890101612f99565b81520191019061037c565b8a80fd5b8880fd5b8580fd5b8380fd5b005b909267ffffffffffffffff6109dd6109d8868686999799613ff5565b614034565b16926109e884614d10565b15610b3757836000526007602052610a066005604060002001614b17565b9260005b8451811015610a4257600190866000526007602052610a3b6005604060002001610a34838961409b565b5190614e3b565b5001610a0a565b5093909491959250806000526007602052600560406000206000815560006001820155600060028201556000600382015560048101610a8181546140af565b9081610af4575b5050018054906000815581610ad3575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190919392936102a7565b6000526020600020908101905b81811015610a985760008155600101610ae0565b81601f60009311600114610b0c5750555b8880610a88565b81835260208320610b2791601f01861c8101906001016142df565b8082528160208120915555610b05565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461028557610c3d36612fe8565b9092610c47614350565b818303610dab5760005b838110610c5a57005b610c8867ffffffffffffffff610c746109d8848887613ff5565b166000526006602052604060002054151590565b15610d6657610c98818487613ff5565b3560048110156102855767ffffffffffffffff610cb96109d8848887613ff5565b16600052600d602052604060002060009160ff7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008354169116179055610d036109d8838786613ff5565b610d0e838689613ff5565b35916004831015610d635750907f2e89b8ad2616113d66baef8b897282a99a93ee74fc684282392d6a725bc9447167ffffffffffffffff60019493610d5a60405192839216948261310c565b0390a201610c51565b80fd5b6109d8610d7d918567ffffffffffffffff94613ff5565b7f1e670e4b000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b7f568efce20000000000000000000000000000000000000000000000000000000060005260046000fd5b346102855760e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557610e0c612f36565b60607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc36011261028557604051610e4281612d55565b60243580151581036102855781526044356fffffffffffffffffffffffffffffffff811681036102855760208201526064356fffffffffffffffffffffffffffffffff8116810361028557604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126102855760405190610ec982612d55565b608435801515810361028557825260a4356fffffffffffffffffffffffffffffffff8116810361028557602083015260c4356fffffffffffffffffffffffffffffffff8116810361028557604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580610f76575b610f48576109ba926145db565b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610f3b565b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855767ffffffffffffffff610fd8612f36565b610fe06141c4565b5016600052600760205261105b611002610ffd60406000206141ef565b614718565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610285576000611098614350565b73ffffffffffffffffffffffffffffffffffffffff6110b5614255565b16158015611287575b8015611262575b61123a5773ffffffffffffffffffffffffffffffffffffffff6110e6614255565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a5573ffffffffffffffffffffffffffffffffffffffff61112d614278565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600b541617600b5573ffffffffffffffffffffffffffffffffffffffff61117461429b565b167fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c5560405173ffffffffffffffffffffffffffffffffffffffff6111be612ef2565b16815260243573ffffffffffffffffffffffffffffffffffffffff811680910361123657602082015260443573ffffffffffffffffffffffffffffffffffffffff8116809103611236578160609160407fa33a66dd32843dc04c77f6aa048e4557f085fd2d793a098e06d6e9f89c715aec940152a180f35b8280fd5b807fccb986cb0000000000000000000000000000000000000000000000000000000060049252fd5b5073ffffffffffffffffffffffffffffffffffffffff61128061429b565b16156110c5565b5073ffffffffffffffffffffffffffffffffffffffff6112a5614278565b16156110be565b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610285576040516005548082528160208101600560005260206000209260005b8181106113e157505061130c92500382612e11565b80519061133161131b83613f89565b926113296040519485612e11565b808452613f89565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208401920136833760005b8151811015611391578067ffffffffffffffff61137e6001938561409b565b511661138a828761409b565b520161135f565b5050906040519182916020830190602084525180915260408301919060005b8181106113be575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016113b0565b84548352600194850194869450602090930192016112f7565b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855773ffffffffffffffffffffffffffffffffffffffff611446612ef2565b61144e614350565b1680156114c85760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160045490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760045573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a1005b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855767ffffffffffffffff611532612f36565b16600052600760205261105b61154e6004604060002001614102565b604051918291602083526020830190612eaf565b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855767ffffffffffffffff6115f4612f36565b6115fc6141c4565b5016600052600760205261105b611002610ffd60026040600020016141ef565b346102855767ffffffffffffffff61163336613058565b92909161163e614350565b1690611657826000526006602052604060002054151590565b1561171357816000526007602052611688600560406000200161167b368685612f62565b6020815191012090614e3b565b156116cc577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691926116c76040519283926020845260208401916132ae565b0390a2005b61170f906040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916132ae565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855767ffffffffffffffff611781612f36565b16600052600d60205261105b60ff604060002054166040519182918261310c565b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855760405160025490818152602081018092600260005260206000209060005b81811061185e5750505081611805910382612e11565b6040519182916020830190602084525180915260408301919060005b81811061182f575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101611821565b82548452602090930192600192830192016117ef565b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855767ffffffffffffffff6118b4612f36565b1660005260076020526118cd6005604060002001614b17565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06119136118fd84613f89565b9361190b6040519586612e11565b808552613f89565b0160005b8181106119ef57505060005b815181101561196b57806119396001928461409b565b51600052600860205261194f6040600020614102565b611959828661409b565b52611964818561409b565b5001611923565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b8282106119a457505050500390f35b919360206119df827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851612eaf565b9601920192018594939192611995565b806060602080938701015201611917565b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855760043567ffffffffffffffff811161028557806004019060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126102855760606020604051611a8281612da0565b828152015260848101611a94816142be565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361231757602482019277ffffffffffffffff00000000000000000000000000000000611afa85614034565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561222f576000916122e8575b506122be576044830190611b8b826142be565b7f0000000000000000000000000000000000000000000000000000000000000000612268575b5067ffffffffffffffff611bc486614034565b16611bdc816000526006602052604060002054151590565b1561223b57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa90811561222f576000916121c8575b5073ffffffffffffffffffffffffffffffffffffffff16330361219a57846064611c7167ffffffffffffffff97614034565b95013595869516806000526007602052611cc8604060002073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001697889161508b565b6040805173ffffffffffffffffffffffffffffffffffffffff88168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a267ffffffffffffffff611d2282614034565b16600052600d60205260ff6040600020541660048110158061213c57811561216b57611d4c613188565b9160009161213c57600381036120dc5750506040015173ffffffffffffffffffffffffffffffffffffffff16935b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff8716602483015260448083018a9052825296611e4c9190611ddc606483612e11565b60008060409a8b5194611def8d87612e11565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d156120d4573d91611e3083612e52565b92611e3d8c519485612e11565b83523d6000602085013e6152ff565b80519081612034575b505085517f9a4575b9000000000000000000000000000000000000000000000000000000008152602060048201529687948594611e92908061325e565b6024870160a0905260c4870190611ea8926132ae565b93611eb290612f4d565b67ffffffffffffffff166044860152611eca90612f15565b73ffffffffffffffffffffffffffffffffffffffff1660648501526084840152611ef390612f15565b73ffffffffffffffffffffffffffffffffffffffff1660a4830152039173ffffffffffffffffffffffffffffffffffffffff1691815a6000948591f191821561202957600092611f93575b5061105b8151928392602084526020611f61835183838801526060870190612eaf565b920151907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030190850152612eaf565b3d8084833e611fa28183612e11565b8101906020818303126109b65780519067ffffffffffffffff821161202557019082828203126109b657825191611fd883612da0565b805167ffffffffffffffff81116109b25782611ff5918301614059565b835260208101519067ffffffffffffffff82116109b2576120199394955001614059565b60208201529082611f3e565b8480fd5b50513d6000823e3d90fd5b602080612045938301019101614338565b15612051578780611e55565b608486517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060916152ff565b6000600182036121075750505073ffffffffffffffffffffffffffffffffffffffff90511693611d7a565b509095919060021461211a575b50611d7a565b6020015173ffffffffffffffffffffffffffffffffffffffff16945087612114565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6024827f31603b1200000000000000000000000000000000000000000000000000000000600052600090600452fd5b7f728fe07b000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b6020813d602011612227575b816121e160209383612e11565b8101031261222357519073ffffffffffffffffffffffffffffffffffffffff82168203610d63575073ffffffffffffffffffffffffffffffffffffffff611c3f565b5080fd5b3d91506121d4565b6040513d6000823e3d90fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff16806000526003602052604060002054611bb1577fd0d259760000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b61230a915060203d602011612310575b6123028183612e11565b810190614338565b85611b78565b503d6122f8565b61233573ffffffffffffffffffffffffffffffffffffffff916142be565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346102855760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855760043567ffffffffffffffff8111610285576123b2903690600401612fb7565b9060243567ffffffffffffffff8111610285576123d39036906004016130db565b9060443567ffffffffffffffff8111610285576123f49036906004016130db565b73ffffffffffffffffffffffffffffffffffffffff600954163314158061248b575b610f4857838614801590612481575b610dab5760005b86811061243557005b8061247b6124496109d86001948b8b613ff5565b612454838989614049565b61247561246d61246586898b614049565b92369061313c565b91369061313c565b916145db565b0161242c565b5080861415612425565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612416565b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557602061254467ffffffffffffffff610c74612f36565b6040519015158152f35b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610285577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6125bd612ef2565b6125c5614350565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a1005b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855760005473ffffffffffffffffffffffffffffffffffffffff811633036126b8577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855760006040805161277281612d55565b82815282602082015201526060612787613188565b73ffffffffffffffffffffffffffffffffffffffff60408051928281511684528260208201511660208501520151166040820152f35b34610285576127cb36613058565b6127d6929192614350565b67ffffffffffffffff82166127f8816000526006602052604060002054151590565b1561281357506109ba9261280d913691612f62565b9061439b565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346102855761286861287061285436612fe8565b9491612861939193614350565b3691613fa1565b923691613fa1565b7f0000000000000000000000000000000000000000000000000000000000000000156129905760005b825181101561290c578073ffffffffffffffffffffffffffffffffffffffff6128c46001938661409b565b51166128cf81614b7a565b6128db575b5001612899565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1846128d4565b5060005b81518110156109ba578073ffffffffffffffffffffffffffffffffffffffff61293b6001938561409b565b5116801561298a5761294c81614f7c565b612959575b505b01612910565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612951565b50612953565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060005260046000fd5b346102855760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610285576129f1612f36565b60243567ffffffffffffffff811161028557602091612a17612544923690600401612f99565b90613f4c565b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855760043567ffffffffffffffff8111610285576101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261028557612a99602091600401613448565b60405190518152f35b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610285576020612b37612ef2565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102855760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102855761105b6040805190612c2c8183612e11565b601c82527f55534443546f6b656e506f6f6c50726f787920312e362e322d64657600000000602083015251918291602083526020830190612eaf565b346102855760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261028557600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361028557817faff2afbf0000000000000000000000000000000000000000000000000000000060209314908115612d2b575b8115612d01575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483612cfa565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150612cf3565b6060810190811067ffffffffffffffff821117612d7157604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117612d7157604052565b60a0810190811067ffffffffffffffff821117612d7157604052565b6020810190811067ffffffffffffffff821117612d7157604052565b610100810190811067ffffffffffffffff821117612d7157604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612d7157604052565b67ffffffffffffffff8111612d7157601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b838110612e9f5750506000910152565b8181015183820152602001612e8f565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093612eeb81518092818752878088019101612e8c565b0116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361028557565b359073ffffffffffffffffffffffffffffffffffffffff8216820361028557565b6004359067ffffffffffffffff8216820361028557565b359067ffffffffffffffff8216820361028557565b929192612f6e82612e52565b91612f7c6040519384612e11565b829481845281830111610285578281602093846000960137010152565b9080601f8301121561028557816020612fb493359101612f62565b90565b9181601f840112156102855782359167ffffffffffffffff8311610285576020808501948460051b01011161028557565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102855760043567ffffffffffffffff8111610285578161303191600401612fb7565b929092916024359067ffffffffffffffff82116102855761305491600401612fb7565b9091565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126102855760043567ffffffffffffffff81168103610285579160243567ffffffffffffffff811161028557826023820112156102855780600401359267ffffffffffffffff84116102855760248483010111610285576024019190565b9181601f840112156102855782359167ffffffffffffffff8311610285576020808501946060850201011161028557565b91906020830192600482101561213c5752565b35906fffffffffffffffffffffffffffffffff8216820361028557565b91908260609103126102855760405161315481612d55565b8092803590811515820361028557604061318391819385526131786020820161311f565b60208601520161311f565b910152565b6040519061319582612d55565b8173ffffffffffffffffffffffffffffffffffffffff600a5416815273ffffffffffffffffffffffffffffffffffffffff600b54166020820152604073ffffffffffffffffffffffffffffffffffffffff600c5416910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610285570180359067ffffffffffffffff82116102855760200191813603831361028557565b90816020910312610285576040519061325882612dd8565b51815290565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561028557016020813591019167ffffffffffffffff821161028557813603831361028557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b90612fb4916020815261341761340c6133cf61331f61330c868061325e565b61010060208801526101208701916132ae565b67ffffffffffffffff61333460208801612f4d565b16604086015273ffffffffffffffffffffffffffffffffffffffff61335b60408801612f15565b1660608601526060860135608086015273ffffffffffffffffffffffffffffffffffffffff61338c60808801612f15565b1660a086015261339f60a087018761325e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08784030160c08801526132ae565b6133dc60c086018661325e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08684030160e08701526132ae565b9260e081019061325e565b916101007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828603019101526132ae565b60405161345481612dd8565b6000905260808101606082013561346a826142be565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603613f2e576020830177ffffffffffffffff000000000000000000000000000000006134cf82614034565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561222f57600091613f0f575b506122be5767ffffffffffffffff61356482614034565b1661357c816000526006602052604060002054151590565b1561223b57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561222f57600091613ef0575b501561219a576135f481614034565b9061361160a0860192612a1761360a85896131ef565b3691612f62565b15613ea95767ffffffffffffffff61362882614034565b168060005260076020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c848061369f600260406000200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161508b565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26136cf613188565b9360c086017ffa7c07de000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000613720838a6131ef565b90358281169160048110613e94575b50501614613e3657604061374382896131ef565b9050146138c557505050505061375c60e08301836131ef565b6004811061389857600411610285573560e01c8061380a57506137cb60009273ffffffffffffffffffffffffffffffffffffffff6020935116906040519485809481937f39077537000000000000000000000000000000000000000000000000000000008352600483016132ed565b03925af190811561222f576000916137e1575090565b612fb4915060203d602011613803575b6137fb8183612e11565b810190613240565b503d6137f1565b6001810361386b57506137cb60009273ffffffffffffffffffffffffffffffffffffffff60208094015116906040519485809481937f39077537000000000000000000000000000000000000000000000000000000008352600483016132ed565b7f68d2f8d60000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f758b22cc0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b606060e0604098939495979698516138dc81612df4565b8281526000602082015260006040820152600083820152600060808201528260a08201528260c0820152015261010082360312610285576040519361392085612df4565b82359067ffffffffffffffff8211610285576139426139499236908601612f99565b8652612f4d565b9060208501918252604083019561395f87612f15565b96604087019788526060870199808b5261397887612f15565b99608089019a8b523567ffffffffffffffff81116102855761399d9036908801612f99565b9360a08901948552803567ffffffffffffffff8111610285576139c39036908901612f99565b9660c08a0197885260e081013567ffffffffffffffff8111610285576139ec9036908301612f99565b9860e08b01998a5273ffffffffffffffffffffffffffffffffffffffff600a541692604051927f6b716b0d000000000000000000000000000000000000000000000000000000008452602084600481885afa93841561222f57600094613dee575b506020600495604051968780927f98db96430000000000000000000000000000000000000000000000000000000082525afa94851561222f57600095613d99575b506040969594939291613aa0916131ef565b9080979181010312610285576020613ab787612f4d565b9601359463ffffffff86168096036102855773ffffffffffffffffffffffffffffffffffffffff613aea613af1926142be565b16916142be565b916040519661014088019088821067ffffffffffffffff831117612d715767ffffffffffffffff91604052168752602087019586526040870194600186526060880190815263ffffffff608089019216825260a0880192835273ffffffffffffffffffffffffffffffffffffffff60c089019416845273ffffffffffffffffffffffffffffffffffffffff60e08901951685526101008801956000875263ffffffff6101208a019860008a5267ffffffffffffffff6040519b511660208c0152511660408a015251600381101561213c5760608901525160808801525163ffffffff1660a08701525160c08601525173ffffffffffffffffffffffffffffffffffffffff1660e085015251610100840152516101208301525163ffffffff166101408201526101408152613c2761016082612e11565b84525173ffffffffffffffffffffffffffffffffffffffff169660405198899788977f390775370000000000000000000000000000000000000000000000000000000089526004890160209052516024890161010090526101248901613c8c91612eaf565b945167ffffffffffffffff1660448901525173ffffffffffffffffffffffffffffffffffffffff1660648801525160848701525173ffffffffffffffffffffffffffffffffffffffff1660a486015251908481037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc0160c4860152613d1091612eaf565b9051908381037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc0160e4850152613d4691612eaf565b9051908281037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01610104840152613d7d91612eaf565b03815a602094600091f190811561222f576000916137e1575090565b6020959195813d602011613de6575b81613db560209383612e11565b8101031261222357519073ffffffffffffffffffffffffffffffffffffffff82168203610d63575093613aa0613a8e565b3d9150613da8565b60209491929394813d602011613e2e575b81613e0c60209383612e11565b8101031261222357519063ffffffff82168203610d6357509291906020613a4d565b3d9150613dff565b50505050506137cb60009273ffffffffffffffffffffffffffffffffffffffff6040602094015116906040519485809481937f39077537000000000000000000000000000000000000000000000000000000008352600483016132ed565b839250829060040360031b1b1616388061372f565b613eb382866131ef565b61170f6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916132ae565b613f09915060203d602011612310576123028183612e11565b386135e5565b613f28915060203d602011612310576123028183612e11565b3861354d565b73ffffffffffffffffffffffffffffffffffffffff612335836142be565b9067ffffffffffffffff612fb492166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111612d715760051b60200190565b929190613fad81613f89565b93613fbb6040519586612e11565b602085838152019160051b810192831161028557905b828210613fdd57505050565b60208091613fea84612f15565b815201910190613fd1565b91908110156140055760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff811681036102855790565b9190811015614005576060020190565b81601f8201121561028557805161406f81612e52565b9261407d6040519485612e11565b8184526020828401011161028557612fb49160208085019101612e8c565b80518210156140055760209160051b010190565b90600182811c921680156140f8575b60208310146140c957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916140be565b9060405191826000825492614116846140af565b8084529360018116908115614184575060011461413d575b5061413b92500383612e11565b565b90506000929192526020600020906000915b81831061416857505090602061413b928201013861412e565b602091935080600191548385890101520191019091849261414f565b6020935061413b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201013861412e565b604051906141d182612dbc565b60006080838281528260208201528260408201528260608201520152565b906040516141fc81612dbc565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b60043573ffffffffffffffffffffffffffffffffffffffff811681036102855790565b60243573ffffffffffffffffffffffffffffffffffffffff811681036102855790565b60443573ffffffffffffffffffffffffffffffffffffffff811681036102855790565b3573ffffffffffffffffffffffffffffffffffffffff811681036102855790565b8181106142ea575050565b600081556001016142df565b8181029291811591840414171561430957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90816020910312610285575180151581036102855790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361437157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156114c85767ffffffffffffffff815160208301209216918260005260076020526143d0816005604060002001615036565b156145975760005260086020526040600020815167ffffffffffffffff8111612d71576143fd82546140af565b601f8111614565575b506020601f821160011461449f5791614479827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361448f95600091614494575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190612eaf565b0390a2565b905084015138614448565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b81811061454d57509261448f9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614516575b5050811b01905561154e565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388061450a565b9192602060018192868a0151815501940192016144cf565b61459190836000526020600020601f840160051c810191602085106108e157601f0160051c01906142df565b38614406565b509061170f6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190612eaf565b67ffffffffffffffff1660008181526006602052604090205490929190156146dd57916146da60e0926146a6856146327f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b9761479d565b8460005260076020526146498160406000206148e4565b6146528361479d565b84600052600760205261466c8360026040600020016148e4565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9190820391821161430957565b6147206141c4565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff808351169161477d602085019361477761476a63ffffffff8751164261470b565b85608089015116906142f6565b90614f6f565b8082101561479657505b16825263ffffffff4216905290565b9050614787565b80511561483d576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106147da5750565b60649061483b604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604082015116158015906148c5575b6148645750565b60649061483b604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff602082015116151561485d565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991614a1d606092805461492163ffffffff8260801c164261470b565b9081614a5c575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014614a5457508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556149d18651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6146da60405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091614958565b6fffffffffffffffffffffffffffffffff91614a91839283614a8a6001880154948286169560801c906142f6565b9116614f6f565b80821015614b1057505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880614928565b9050614a9b565b906040519182815491828252602082019060005260206000209260005b818110614b4957505061413b92500383612e11565b8454835260019485019487945060209093019201614b34565b80548210156140055760005260206000200190600090565b6000818152600360205260409020548015614d09577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161430957600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161430957818103614c9a575b5050506002548015614c6b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01614c28816002614b62565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b614cf1614cab614cbc936002614b62565b90549060031b1c9283926002614b62565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080614bef565b5050600090565b6000818152600660205260409020548015614d09577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161430957600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161430957818103614e01575b5050506005548015614c6b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01614dbe816005614b62565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b614e23614e12614cbc936005614b62565b90549060031b1c9283926005614b62565b90556000526006602052604060002055388080614d85565b9060018201918160005282602052604060002054801515600014614f66577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614309578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161430957818103614f2f575b50505080548015614c6b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190614ef08282614b62565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b614f4f614f3f614cbc9386614b62565b90549060031b1c92839286614b62565b905560005283602052604060002055388080614eb8565b50505050600090565b9190820180921161430957565b80600052600360205260406000205415600014614fd65760025468010000000000000000811015612d7157614fbd614cbc8260018594016002556002614b62565b9055600254906000526003602052604060002055600190565b50600090565b80600052600660205260406000205415600014614fd65760055468010000000000000000811015612d715761501d614cbc8260018594016005556005614b62565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054614d095780549068010000000000000000821015612d715782615074614cbc846001809601855584614b62565b905580549260005201602052604060002055600190565b9182549060ff8260a01c161580156152f7575b6152f1576fffffffffffffffffffffffffffffffff821691600185019081546150e363ffffffff6fffffffffffffffffffffffffffffffff83169360801c164261470b565b9081615253575b505084811061520757508383106151445750506151196fffffffffffffffffffffffffffffffff92839261470b565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615153818561470b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018481116143095761518691614f6f565b9280156151d85773ffffffffffffffffffffffffffffffffffffffff937fd0c8d23a00000000000000000000000000000000000000000000000000000000600052046004526024521660445260646000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b8286929396116152c75761526e926147779160801c906142f6565b808410156152c25750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806150ea565b615279565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561509e565b9192901561537a5750815115615313575090565b3b1561531c5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561538d5750805190602001fd5b61170f906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190612eaf56fea164736f6c634300081a000a",
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

func (_USDCTokenPoolProxy *USDCTokenPoolProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _USDCTokenPoolProxy.abi.Events["AllowListAdd"].ID:
		return _USDCTokenPoolProxy.ParseAllowListAdd(log)
	case _USDCTokenPoolProxy.abi.Events["AllowListRemove"].ID:
		return _USDCTokenPoolProxy.ParseAllowListRemove(log)
	case _USDCTokenPoolProxy.abi.Events["ChainAdded"].ID:
		return _USDCTokenPoolProxy.ParseChainAdded(log)
	case _USDCTokenPoolProxy.abi.Events["ChainConfigured"].ID:
		return _USDCTokenPoolProxy.ParseChainConfigured(log)
	case _USDCTokenPoolProxy.abi.Events["ChainRemoved"].ID:
		return _USDCTokenPoolProxy.ParseChainRemoved(log)
	case _USDCTokenPoolProxy.abi.Events["ConfigChanged"].ID:
		return _USDCTokenPoolProxy.ParseConfigChanged(log)
	case _USDCTokenPoolProxy.abi.Events["InboundRateLimitConsumed"].ID:
		return _USDCTokenPoolProxy.ParseInboundRateLimitConsumed(log)
	case _USDCTokenPoolProxy.abi.Events["LockOrBurnMechanismUpdated"].ID:
		return _USDCTokenPoolProxy.ParseLockOrBurnMechanismUpdated(log)
	case _USDCTokenPoolProxy.abi.Events["LockedOrBurned"].ID:
		return _USDCTokenPoolProxy.ParseLockedOrBurned(log)
	case _USDCTokenPoolProxy.abi.Events["OutboundRateLimitConsumed"].ID:
		return _USDCTokenPoolProxy.ParseOutboundRateLimitConsumed(log)
	case _USDCTokenPoolProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _USDCTokenPoolProxy.ParseOwnershipTransferRequested(log)
	case _USDCTokenPoolProxy.abi.Events["OwnershipTransferred"].ID:
		return _USDCTokenPoolProxy.ParseOwnershipTransferred(log)
	case _USDCTokenPoolProxy.abi.Events["PoolAddressesUpdated"].ID:
		return _USDCTokenPoolProxy.ParsePoolAddressesUpdated(log)
	case _USDCTokenPoolProxy.abi.Events["RateLimitAdminSet"].ID:
		return _USDCTokenPoolProxy.ParseRateLimitAdminSet(log)
	case _USDCTokenPoolProxy.abi.Events["ReleasedOrMinted"].ID:
		return _USDCTokenPoolProxy.ParseReleasedOrMinted(log)
	case _USDCTokenPoolProxy.abi.Events["RemotePoolAdded"].ID:
		return _USDCTokenPoolProxy.ParseRemotePoolAdded(log)
	case _USDCTokenPoolProxy.abi.Events["RemotePoolRemoved"].ID:
		return _USDCTokenPoolProxy.ParseRemotePoolRemoved(log)
	case _USDCTokenPoolProxy.abi.Events["RouterUpdated"].ID:
		return _USDCTokenPoolProxy.ParseRouterUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
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
	return common.HexToHash("0xa33a66dd32843dc04c77f6aa048e4557f085fd2d793a098e06d6e9f89c715aec")
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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
