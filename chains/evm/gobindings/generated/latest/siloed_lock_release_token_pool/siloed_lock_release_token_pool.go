// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package siloed_lock_release_token_pool

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

type SiloedLockReleaseTokenPoolSiloConfigUpdate struct {
	RemoteChainSelector uint64
	Rebalancer          common.Address
}

type TokenPoolChainUpdate struct {
	RemoteChainSelector       uint64
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

var SiloedLockReleaseTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"structSiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103b957615f86803803809161001d828561066f565b8339810160c0828203126103b95781516001600160a01b038116929091908383036103b95761004e60208201610692565b60408201516001600160401b0381116103b95782019280601f850112156103b9578351936001600160401b0385116103be578460051b906020820195610097604051978861066f565b86526020808701928201019283116103b957602001905b828210610657575050506100c4606083016106a0565b6100dc60a06100d5608086016106a0565b94016106a0565b94331561064657600180546001600160a01b0319163317905586158015610635575b8015610624575b61048a5760805260c05260405163313ce56760e01b8152602081600481895afa600091816105e8575b506105bd575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261049b575b506001600160a01b0316801561048a57604051636eb1769f60e11b815230600482015260248101829052602081604481865afa90811561047e5760009161044c575b506103e15760405191602083019263095ea7b360e01b84528260248201526000196044820152604481526101dd60648261066f565b6000806040958651936101f0888661066f565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d156103d4573d906001600160401b0382116103be578551610261949092610252601f8201601f19166020018561066f565b83523d6000602085013e610854565b80518061033d575b50506101005251615661908161092582396080518181816103e6015281816112ca01528181611b8001528181611d7001528181612ed5015281816130b40152818161354001528181613727015281816137810152613923015260a051818181611e59015281816136cf0152818161445101526144d4015260c051818181610fce01528181611c1c0152612f71015260e051818181610f7c01528181611c5f0152612c790152610100518181816104980152818161126601528181611f1d015281816131940152818161359401526138bf0152f35b81602091810103126103b957602001518015908115036103b957610362573880610269565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b9161026192606091610854565b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d602011610476575b816104676020938361066f565b810103126103b95751386101a8565b3d915061045a565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b90602090604051906104ad838361066f565b60008252600036813760e051156105ac5760005b8251811015610528576001906001600160a01b036104df82866106b4565b5116856104eb826106f6565b6104f8575b5050016104c1565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138856104f0565b5092905060005b81518110156105a3576001906001600160a01b0361054d82856106b4565b5116801561059d578461055f826107f4565b61056d575b50505b0161052f565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13884610564565b50610567565b50505038610166565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036105d15750610134565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d60201161061c575b816106046020938361066f565b810103126103b95761061590610692565b903861012e565b3d91506105f7565b506001600160a01b03821615610105565b506001600160a01b038416156100fe565b639b15e16f60e01b60005260046000fd5b60208091610664846106a0565b8152019101906100ae565b601f909101601f19168101906001600160401b038211908210176103be57604052565b519060ff821682036103b957565b51906001600160a01b03821682036103b957565b80518210156106c85760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156106c85760005260206000200190600090565b60008181526003602052604090205480156107ed5760001981018181116107d7576002546000198101919082116107d757818103610786575b5050506002548015610770576000190161074a8160026106de565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6107bf6107976107a89360026106de565b90549060031b1c92839260026106de565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061072f565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461084e57600254680100000000000000008110156103be576108356107a882600185940160025560026106de565b9055600254906000526003602052604060002055600190565b50600090565b919290156108b65750815115610868575090565b3b156108715790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156108c95750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b83811061090c5750508160006044809484010152601f80199101168101030190fd5b602082820181015160448784010152859350016108ea56fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613a56575080630a861f2a1461382c578063181f5a77146137a557806321df0da714613754578063240028e8146136f357806324f65ee7146136b55780632d4a148f1461343a57806331238ffc146133f65780633907753714612e3d578063432a6ba314612e095780634c5ef0ed14612dc257806354c8a4f314612c4557806362ddd3c414612bc15780636600f92c14612aab5780636cfd1553146129f05780636d3d1a58146129bc5780636d9d216c1461257857806379ba5097146124b15780637d54534e146124225780638632d5cc146123e05780638926f54f1461239a5780638da5cb5b14612366578063962d4020146121e05780639a4575b914611b16578063a42a7b8b146119af578063a7cd63b714611901578063acfecf91146117e3578063af0e58b9146117c5578063af58d59f1461177c578063b0f479a114611748578063b79465801461170f578063c0d7865514611627578063c4bffe2b146114fc578063c75eea9c14611454578063ce3c75281461119f578063cf7401f314610ff2578063dc0bd97114610fa1578063e0351e1314610f64578063e8a1da171461068f578063eb521a4c146102ec578063f1e73399146102c05763f2fde38b146101ef57600080fd5b346102bd5760206003193601126102bd5773ffffffffffffffffffffffffffffffffffffffff61021d613b86565b6102256145f6565b1633811461029557807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102bd5760206003193601126102bd5760206102e46102df613ba9565b614357565b604051908152f35b50346102bd5760206003193601126102bd5760043580156106675773ffffffffffffffffffffffffffffffffffffffff6103258361408d565b16330361063b57818052600c60205260408220600181015460a01c60ff161561062657610353828254614070565b90555b6040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201908152336024830152306044830152606482018390529083906103cc81608481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282613c5f565b61047773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169183806040968751936104198986613c5f565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082875af13d1561061e573d9061045a82613ca0565b9161046788519384613c5f565b82523d86602084013e5b84615588565b805180610579575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b1561056b5783517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9290921660048301526024820185905282908290604490829084905af1801561056f57610556575b5050807f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf91519284845260208401523392a280f35b8161056091613c5f565b61056b578238610521565b8280fd5b83513d84823e3d90fd5b61058d9293506020809183010191016145de565b1561059b578390388061047f565b608482517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606090610471565b5061063381600a54614070565b600a55610356565b6024827f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b6004827fa90c0d19000000000000000000000000000000000000000000000000000000008152fd5b50346102bd5761069e36613d60565b939190926106aa6145f6565b82915b808310610dcf575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b85821015610dcb578160051b85013581811215610dc75785019061012082360312610dc7576040519561071887613c43565b823567ffffffffffffffff81168103610dc2578752602083013567ffffffffffffffff8111610dbe5783019536601f88011215610dbe5786359661075b88613fb1565b97610769604051998a613c5f565b8089526020808a019160051b83010190368211610dba5760208301905b828210610d87575050505060208801968752604084013567ffffffffffffffff8111610d83576107b99036908601613d11565b9860408901998a526107e36107d13660608801613e65565b9560608b0196875260c0369101613e65565b9660808a019788526107f58651614a60565b6107ff8851614a60565b8a515115610d5b5761081b67ffffffffffffffff8b5116615292565b15610d245767ffffffffffffffff8a5116815260076020526040812061095b87516fffffffffffffffffffffffffffffffff604082015116906109166fffffffffffffffffffffffffffffffff6020830151169151151583608060405161088181613c43565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b610a8189516fffffffffffffffffffffffffffffffff60408201511690610a3c6fffffffffffffffffffffffffffffffff602083015116915115158360806040516109a581613c43565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff8211610cf757610aa48354614126565b601f8111610cbc575b50602090601f8311600114610c1d57610afb9291859183610c12575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b805b89518051821015610b365790610b30600192610b29838f67ffffffffffffffff90511692614112565b5190614641565b01610b00565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c292959396610c0467ffffffffffffffff600197949c5116925193519151610bd0610b9b60405196879687526101006020880152610100870190613b27565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a10190939492916106e6565b015190503880610ac9565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b818110610ca45750908460019594939210610c6d575b505050811b019055610afe565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388080610c60565b92936020600181928786015181550195019301610c4a565b610ce79084865260208620601f850160051c81019160208610610ced575b601f0160051c019061432d565b38610aad565b9091508190610cda565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff8111610db657602091610dab8392833691890101613d11565b815201910190610786565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff610def610dea878588614031565b613f02565b1695610dfa87614fd3565b15610f38578684526007602052610e1660056040862001614dda565b94845b8651811015610e4f576001908987526007602052610e4860056040892001610e41838b614112565b51906150fe565b5001610e19565b5093945094909580855260076020526005604086208681558660018201558660028201558660038201558660048201610e888154614126565b80610ef7575b5050500180549086815581610ed9575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946106ad565b865260208620908101905b81811015610e9e57868155600101610ee4565b601f8111600114610f0d5750555b863880610e8e565b81835260208320610f2891601f01861c81019060010161432d565b8082528160208120915555610f05565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102bd57806003193601126102bd5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b50346102bd57806003193601126102bd57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102bd5760e06003193601126102bd5761100c613ba9565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc3601126102bd5760405161104381613c0b565b602435801515810361056b5781526044356fffffffffffffffffffffffffffffffff8116810361056b5760208201526064356fffffffffffffffffffffffffffffffff8116810361056b57604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c36011261119b57604051906110ca82613c0b565b6084358015158103610dc757825260a4356fffffffffffffffffffffffffffffffff81168103610dc757602083015260c4356fffffffffffffffffffffffffffffffff81168103610dc757604083015273ffffffffffffffffffffffffffffffffffffffff6009541633141580611179575b61114d5761114a92936148ab565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561113c565b5080fd5b50346102bd5760406003193601126102bd576111b9613ba9565b60243567ffffffffffffffff8216808452600c60205260ff600160408620015460a01c1615801561144c575b6114215781156113f95773ffffffffffffffffffffffffffffffffffffffff61120d8461408d565b1633036113cd578352600c60205260408320600181015460a01c60ff1680156113c55781545b8084116113955750156113805761124b828254613f17565b90555b8273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b1561119b576040517f69328dec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166004820152602481018490523360448201529082908290606490829084905af1801561137557611360575b50506040805167ffffffffffffffff9093168352602083019190915233917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e91819081015b0390a280f35b8161136a91613c5f565b61056b578238611315565b6040513d84823e3d90fd5b5061138d81600a54613f17565b600a5561124e565b85846044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600a54611233565b6024847f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b6004847fa90c0d19000000000000000000000000000000000000000000000000000000008152fd5b7f46f5f12b000000000000000000000000000000000000000000000000000000008452600452602483fd5b5080156111e5565b50346102bd5760206003193601126102bd5761149f61149a60406114f89367ffffffffffffffff611483613ba9565b61148b61427a565b501681526007602052206142a5565b6149db565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b50346102bd57806003193601126102bd57604051906005548083528260208101600584526020842092845b81811061160e57505061153c92500383613c5f565b815161156061154a82613fb1565b916115586040519384613c5f565b808352613fb1565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b84518110156115bf578067ffffffffffffffff6115ac60019388614112565b51166115b88286614112565b520161158d565b50925090604051928392602084019060208552518091526040840192915b8181106115eb575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016115dd565b8454835260019485019487945060209093019201611527565b50346102bd5760206003193601126102bd57611641613b86565b6116496145f6565b73ffffffffffffffffffffffffffffffffffffffff81169081156116e757600480547fffffffffffffffffffffffff0000000000000000000000000000000000000000811690931790556040805173ffffffffffffffffffffffffffffffffffffffff93841681529190921660208201527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f168491819081015b0390a180f35b6004837f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346102bd5760206003193601126102bd576114f861173461172f613ba9565b61430b565b604051918291602083526020830190613b27565b50346102bd57806003193601126102bd57602073ffffffffffffffffffffffffffffffffffffffff60045416604051908152f35b50346102bd5760206003193601126102bd5761149f61149a600260406114f89467ffffffffffffffff6117ad613ba9565b6117b561427a565b50168152600760205220016142a5565b50346102bd57806003193601126102bd576020600a54604051908152f35b50346102bd5767ffffffffffffffff6117fb36613db2565b9290916118066145f6565b169161181f836000526006602052604060002054151590565b156118d557828452600760205261184e60056040862001611841368486613cda565b60208151910120906150fe565b1561188d57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769161135a60405192839260208452602084019161423b565b826118d1836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161423b565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102bd57806003193601126102bd57604051600254808252602082018091600285526020852090855b8181106119995750505082611942910383613c5f565b604051928392602084019060208552518091526040840192915b81811061196a575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff1684528594506020938401939092019160010161195c565b825484526020909301926001928301920161192c565b50346102bd5760206003193601126102bd5767ffffffffffffffff6119d2613ba9565b16815260076020526119e960056040832001614dda565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611a2e611a1883613fb1565b92611a266040519485613c5f565b808452613fb1565b01835b818110611b05575050825b8251811015611a825780611a5260019285614112565b5185526008602052611a6660408620614179565b611a708285614112565b52611a7b8184614112565b5001611a3c565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b828210611aba57505050500390f35b91936020611af5827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613b27565b9601920192018594939192611aab565b806060602080938601015201611a31565b50346102bd5760206003193601126102bd576004359067ffffffffffffffff82116102bd5760a060031983360301126102bd57611b516140f9565b50611b5a6140f9565b5060848201611b6881613f53565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036121965750602482019177ffffffffffffffff00000000000000000000000000000000611bcf84613f02565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156120b7578391612167575b5061213f57611c5d60448201613f53565b7f00000000000000000000000000000000000000000000000000000000000000006120ed575b5067ffffffffffffffff611c9684613f02565b16611cae816000526006602052604060002054151590565b156120c257602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156120b7578390612054575b73ffffffffffffffffffffffffffffffffffffffff9150163303612028579067ffffffffffffffff916064611d4385613f02565b910135928391168083526007602052611d986040842073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016958691615341565b6040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018490527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a267ffffffffffffffff611df285613f02565b6040805173ffffffffffffffffffffffffffffffffffffffff871681523360208201529081018490529116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2611e5061172f85613f02565b926040519460ff7f000000000000000000000000000000000000000000000000000000000000000016602087015260208652611e8d604087613c5f565b60405194611e9a86613c27565b85526020850195865267ffffffffffffffff611eb582613f02565b168452600c602052604084206001015460a01c60ff1615612013578067ffffffffffffffff611ee6611f0493613f02565b168552600c60205260408520611efd858254614070565b9055613f02565b505b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15610dc7576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff919091166004820152602481019290925282908290604490829084905af1801561137557611ffe575b611fcd836114f8866040519384936020855251604060208601526060850190613b27565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0848303016040850152613b27565b612009828092613c5f565b6102bd5780611fa9565b5061202082600a54614070565b600a55611f06565b6024827f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d6020116120af575b8161206e60209383613c5f565b8101031261056b575173ffffffffffffffffffffffffffffffffffffffff8116810361056b5773ffffffffffffffffffffffffffffffffffffffff90611d0f565b3d9150612061565b6040513d85823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008352600452602482fd5b73ffffffffffffffffffffffffffffffffffffffff1680835260036020526040832054611c83577fd0d25976000000000000000000000000000000000000000000000000000000008352600452602482fd5b6004827f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612189915060203d60201161218f575b6121818183613c5f565b8101906145de565b38611c4c565b503d612177565b9073ffffffffffffffffffffffffffffffffffffffff6121b7602493613f53565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102bd5760606003193601126102bd5760043567ffffffffffffffff811161119b57612212903690600401613d2f565b60243567ffffffffffffffff8111610dc757612232903690600401613e17565b60449291923567ffffffffffffffff8111610dbe57612255903690600401613e17565b91909273ffffffffffffffffffffffffffffffffffffffff6009541633141580612344575b6123185781811480159061230e575b6122e657865b81811061229a578780f35b806122e06122ae610dea600194868c614031565b6122b983878b6140e9565b6122da6122d26122ca868b8d6140e9565b923690613e65565b913690613e65565b916148ab565b0161228f565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b5082811415612289565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561227a565b50346102bd57806003193601126102bd57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102bd5760206003193601126102bd5760206123d667ffffffffffffffff6123c2613ba9565b166000526006602052604060002054151590565b6040519015158152f35b50346102bd5760206003193601126102bd5760206124046123ff613ba9565b61408d565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b50346102bd5760206003193601126102bd577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff612474613b86565b61247c6145f6565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006009541617600955604051908152a180f35b50346102bd57806003193601126102bd57805473ffffffffffffffffffffffffffffffffffffffff81163303612550577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102bd5760406003193601126102bd5760043567ffffffffffffffff811161119b576125aa903690600401613d2f565b906024359067ffffffffffffffff8211610dc75736602383011215610dc75781600401359267ffffffffffffffff84116129b8576024830192602436918660061b0101116129b8576125fa6145f6565b845b81811061289257505050825b828110612613578380f35b67ffffffffffffffff61262a610dea83868661407d565b1615801561285f575b801561283e575b6127fb5773ffffffffffffffffffffffffffffffffffffffff612669602061266384878761407d565b01613f53565b16156127d3578061275b848473ffffffffffffffffffffffffffffffffffffffff8060018a67ffffffffffffffff6126d5610dea8a6126af6020612663889f8d8d61407d565b99604051956126bd87613c0b565b865287602087019b168b526040860199878b5261407d565b168c52600c60205260408c209051815501935116167fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb61278a610dea83878761407d565b61279a602061266385898961407d565b6040805167ffffffffffffffff93909316835273ffffffffffffffffffffffffffffffffffffffff91909116602083015290a101612608565b6004847f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b610dea612815916024959467ffffffffffffffff9461407d565b7fd9a9cd6800000000000000000000000000000000000000000000000000000000835216600452fd5b5061285967ffffffffffffffff6123c2610dea84878761407d565b1561263a565b5067ffffffffffffffff612877610dea83868661407d565b168452600c60205260ff600160408620015460a01c16612633565b67ffffffffffffffff6128a9610dea838587614031565b168652600c60205260ff600160408820015460a01c1615612975578067ffffffffffffffff6128de610dea6001948688614031565b168752600c6020527f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93604088205461291881600a54614070565b600a5567ffffffffffffffff612932610dea85888a614031565b168952600c6020528884604082208281550155612953610dea848789614031565b6040805167ffffffffffffffff9290921682526020820192909252a1016125fc565b61298f610dea67ffffffffffffffff928894602496614031565b7f46f5f12b00000000000000000000000000000000000000000000000000000000835216600452fd5b8480fd5b50346102bd57806003193601126102bd57602073ffffffffffffffffffffffffffffffffffffffff60095416604051908152f35b50346102bd5760206003193601126102bd577f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b23473ffffffffffffffffffffffffffffffffffffffff612a40613b86565b612a486145f6565b6116e1600b54918381167fffffffffffffffffffffffff0000000000000000000000000000000000000000841617600b55604051938493168390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b50346102bd5760406003193601126102bd57612ac5613ba9565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361056b5767ffffffffffffffff90612af86145f6565b1690818352600c602052600160408420019081549160ff8360a01c1615612b955780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9283169081179091556040805192909316825260208201527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f918190810161135a565b602485857f46f5f12b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102bd57612bd036613db2565b612bdc939291936145f6565b67ffffffffffffffff8216612bfe816000526006602052604060002054151590565b15612c1a575061114a9293612c14913691613cda565b90614641565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102bd57612c6f90612c77612c5b36613d60565b9591612c689391936145f6565b3691613fc9565b933691613fc9565b7f000000000000000000000000000000000000000000000000000000000000000015612d9a57815b8351811015612d12578073ffffffffffffffffffffffffffffffffffffffff612cca60019387614112565b5116612cd581614e3d565b612ce1575b5001612c9f565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138612cda565b5090805b8251811015612d96578073ffffffffffffffffffffffffffffffffffffffff612d4160019386614112565b51168015612d9057612d5281615232565b612d5f575b505b01612d16565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a184612d57565b50612d59565b5080f35b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b50346102bd5760406003193601126102bd57612ddc613ba9565b906024359067ffffffffffffffff82116102bd5760206123d684612e033660048701613d11565b90613f74565b50346102bd57806003193601126102bd57602073ffffffffffffffffffffffffffffffffffffffff600b5416604051908152f35b50346102bd5760206003193601126102bd5760043567ffffffffffffffff811161119b57806004019161010060031983360301126102bd5780604051612e8281613bc0565b52612eaf612ea5612ea0612e9960c4860187613eb1565b3691613cda565b6143de565b60648401356144d1565b9060848301612ebd81613f53565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036121965750602483019077ffffffffffffffff00000000000000000000000000000000612f2483613f02565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156113755782916133d7575b506133af57612faf82613f02565b67ffffffffffffffff8116612fd1816000526006602052604060002054151590565b156120c25750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff93909316918301919091523360248301526020908290604490829073ffffffffffffffffffffffffffffffffffffffff165afa908115611375578291613390575b50156133645761305a82613f02565b61306f60a4860191612e03612e99848a613eb1565b1561331d575090829167ffffffffffffffff61308a83613f02565b1680825260076020526130dc6002604084200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016958691615341565b6040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a267ffffffffffffffff61313683613f02565b168152600c60205260408120600181015460a01c60ff1680156133155781545b8087116132e55750156132d05761316e858254613f17565b905561317982613f02565b505b604473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169501946131c086613f53565b813b1561056b576040517f69328dec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018890529190911660448201529082908290606490829084905af18015611375576132bb575b5050608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff61328b6132857ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096613f02565b92613f53565b60405196875233898801521660408601528560608601521692a2806040516132b281613bc0565b52604051908152f35b6132c6828092613c5f565b6102bd5780613233565b506132dd84600a54613f17565b600a5561317b565b83876044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600a54613156565b6133279086613eb1565b6118d16040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161423b565b807f728fe07b000000000000000000000000000000000000000000000000000000006024925233600452fd5b6133a9915060203d60201161218f576121818183613c5f565b3861304b565b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b6133f0915060203d60201161218f576121818183613c5f565b38612fa1565b50346102bd5760206003193601126102bd5760ff6001604060209367ffffffffffffffff613422613ba9565b168152600c855220015460a01c166040519015158152f35b50346102bd5760406003193601126102bd57613454613ba9565b60243567ffffffffffffffff8216808452600c60205260ff600160408620015460a01c161580156136ad575b6114215781156113f95773ffffffffffffffffffffffffffffffffffffffff6134a88461408d565b1633036113cd578352600c60205260408320600181015460a01c60ff1615613698576134d5828254614070565b90555b6040517f23b872dd00000000000000000000000000000000000000000000000000000000602082019081523360248301523060448301526064820183905290849061352681608481016103a0565b61357373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169183806040968751936104198986613c5f565b805180613676575b505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b1561056b5783517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9290921660048301526024820185905282908290604490829084905af1801561056f57613661575b50505167ffffffffffffffff9092168252602082015233907f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf90806040810161135a565b8161366b91613c5f565b610dc757833861361d565b61368a9293506020809183010191016145de565b1561059b578490388061357b565b506136a581600a54614070565b600a556134d8565b508015613480565b50346102bd57806003193601126102bd57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102bd5760206003193601126102bd57602061370f613b86565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102bd57806003193601126102bd57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102bd57806003193601126102bd57506114f86040516137c8606082613c5f565b602481527f53696c6f65644c6f636b52656c65617365546f6b656e506f6f6c20312e362e3360208201527f2d646576000000000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190613b27565b50346102bd5760206003193601126102bd57600435908115613a2e5773ffffffffffffffffffffffffffffffffffffffff6138668261408d565b163303613a0257808052600c60205260408120600181015460a01c60ff1680156139fa5781545b8085116139ca5750156139b5576138a5838254613f17565b90555b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b1561119b576040517f69328dec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166004820152602481018490523360448201529082908290606490829084905af18015611375576139a5575b50906040519082825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a280f35b816139af91613c5f565b3861396e565b506139c282600a54613f17565b600a556138a8565b83856044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600a5461388d565b807f8e4a23d6000000000000000000000000000000000000000000000000000000006024925233600452fd5b807fa90c0d190000000000000000000000000000000000000000000000000000000060049252fd5b90503461119b57602060031936011261119b576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361056b57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115613afd575b8115613ad3575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613acc565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613ac5565b919082519283825260005b848110613b715750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613b32565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610dc257565b6004359067ffffffffffffffff82168203610dc257565b6020810190811067ffffffffffffffff821117613bdc57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6060810190811067ffffffffffffffff821117613bdc57604052565b6040810190811067ffffffffffffffff821117613bdc57604052565b60a0810190811067ffffffffffffffff821117613bdc57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613bdc57604052565b67ffffffffffffffff8111613bdc57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192613ce682613ca0565b91613cf46040519384613c5f565b829481845281830111610dc2578281602093846000960137010152565b9080601f83011215610dc257816020613d2c93359101613cda565b90565b9181601f84011215610dc25782359167ffffffffffffffff8311610dc2576020808501948460051b010111610dc257565b6040600319820112610dc25760043567ffffffffffffffff8111610dc25781613d8b91600401613d2f565b929092916024359067ffffffffffffffff8211610dc257613dae91600401613d2f565b9091565b6040600319820112610dc25760043567ffffffffffffffff81168103610dc2579160243567ffffffffffffffff8111610dc25782602382011215610dc25780600401359267ffffffffffffffff8411610dc25760248483010111610dc2576024019190565b9181601f84011215610dc25782359167ffffffffffffffff8311610dc25760208085019460608502010111610dc257565b35906fffffffffffffffffffffffffffffffff82168203610dc257565b9190826060910312610dc257604051613e7d81613c0b565b80928035908115158203610dc2576040613eac9181938552613ea160208201613e48565b602086015201613e48565b910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610dc2570180359067ffffffffffffffff8211610dc257602001918136038313610dc257565b3567ffffffffffffffff81168103610dc25790565b91908203918211613f2457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff81168103610dc25790565b9067ffffffffffffffff613d2c92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613bdc5760051b60200190565b9291613fd482613fb1565b93613fe26040519586613c5f565b602085848152019260051b8101918211610dc257915b81831061400457505050565b823573ffffffffffffffffffffffffffffffffffffffff81168103610dc257815260209283019201613ff8565b91908110156140415760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908201809211613f2457565b91908110156140415760061b0190565b67ffffffffffffffff16600052600c60205260016040600020015460ff8160a01c166140d0575073ffffffffffffffffffffffffffffffffffffffff600b541690565b73ffffffffffffffffffffffffffffffffffffffff1690565b9190811015614041576060020190565b6040519061410682613c27565b60606020838281520152565b80518210156140415760209160051b010190565b90600182811c9216801561416f575b602083101461414057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614135565b906040519182600082549261418d84614126565b80845293600181169081156141fb57506001146141b4575b506141b292500383613c5f565b565b90506000929192526020600020906000915b8183106141df5750509060206141b292820101386141a5565b60209193508060019154838589010152019101909184926141c6565b602093506141b29592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386141a5565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b6040519061428782613c43565b60006080838281528260208201528260408201528260608201520152565b906040516142b281613c43565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526007602052613d2c6004604060002001614179565b818110614338575050565b6000815560010161432d565b81810292918115918404141715613f2457565b67ffffffffffffffff16614378816000526006602052604060002054151590565b156143b15780600052600c60205260ff60016040600020015460a01c166143a05750600a5490565b600052600c60205260406000205490565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8051801561444d5760200361440f57602081805181010312610dc25760208101519060ff821161440f575060ff1690565b6118d1906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613b27565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211613f2457565b60ff16604d8111613f2457600a0a90565b81156144a2570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146145d7578284116145ad579061451691614473565b91604d60ff8416118015614574575b61453e57505090614538613d2c92614487565b90614344565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061457e83614487565b80156144a2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411614525565b6145b691614473565b91604d60ff84161161453e575050906145d1613d2c92614487565b90614498565b5050505090565b90816020910312610dc257518015158103610dc25790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361461757565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156148815767ffffffffffffffff815160208301209216918260005260076020526146768160056040600020016152ec565b1561483d5760005260086020526040600020815167ffffffffffffffff8111613bdc576146a38254614126565b601f811161480b575b506020601f8211600114614745579161471f827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936147359560009161473a575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613b27565b0390a2565b9050840151386146ee565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106147f35750926147359492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106147bc575b5050811b019055611734565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538806147b0565b9192602060018192868a015181550194019201614775565b61483790836000526020600020601f840160051c81019160208510610ced57601f0160051c019061432d565b386146ac565b50906118d16040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613b27565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff1660008181526006602052604090205490929190156149ad57916149aa60e092614976856149027f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97614a60565b846000526007602052614919816040600020614ba7565b61492283614a60565b84600052600760205261493c836002604060002001614ba7565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6149e361427a565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614a406020850193614a3a614a2d63ffffffff87511642613f17565b8560808901511690614344565b90614070565b80821015614a5957505b16825263ffffffff4216905290565b9050614a4a565b805115614b00576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff60208301511610614a9d5750565b606490614afe604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590614b88575b614b275750565b606490614afe604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515614b20565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991614ce06060928054614be463ffffffff8260801c1642613f17565b9081614d1f575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014614d1757508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155614c948651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6149aa60405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091614c1b565b6fffffffffffffffffffffffffffffffff91614d54839283614d4d6001880154948286169560801c90614344565b9116614070565b80821015614dd357505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880614beb565b9050614d5e565b906040519182815491828252602082019060005260206000209260005b818110614e0c5750506141b292500383613c5f565b8454835260019485019487945060209093019201614df7565b80548210156140415760005260206000200190600090565b6000818152600360205260409020548015614fcc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613f2457600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613f2457818103614f5d575b5050506002548015614f2e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01614eeb816002614e25565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b614fb4614f6e614f7f936002614e25565b90549060031b1c9283926002614e25565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080614eb2565b5050600090565b6000818152600660205260409020548015614fcc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613f2457600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613f24578181036150c4575b5050506005548015614f2e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615081816005614e25565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b6150e66150d5614f7f936005614e25565b90549060031b1c9283926005614e25565b90556000526006602052604060002055388080615048565b9060018201918160005282602052604060002054801515600014615229577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613f24578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613f24578181036151f2575b50505080548015614f2e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906151b38282614e25565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b615212615202614f7f9386614e25565b90549060031b1c92839286614e25565b90556000528360205260406000205538808061517b565b50505050600090565b8060005260036020526040600020541560001461528c5760025468010000000000000000811015613bdc57615273614f7f8260018594016002556002614e25565b9055600254906000526003602052604060002055600190565b50600090565b8060005260066020526040600020541560001461528c5760055468010000000000000000811015613bdc576152d3614f7f8260018594016005556005614e25565b9055600554906000526006602052604060002055600190565b6000828152600182016020526040902054614fcc5780549068010000000000000000821015613bdc578261532a614f7f846001809601855584614e25565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615580575b61557a576fffffffffffffffffffffffffffffffff8216916001850190815461539963ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642613f17565b90816154dc575b505084811061549057508383106153fa5750506153cf6fffffffffffffffffffffffffffffffff928392613f17565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916154098185613f17565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211613f245761545761545c9273ffffffffffffffffffffffffffffffffffffffff96614070565b614498565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615550576154f792614a3a9160801c90614344565b8084101561554b5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806153a0565b615502565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615354565b91929015615603575081511561559c575090565b3b156155a55790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156156165750805190602001fd5b6118d1906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613b2756fea164736f6c634300081a000a",
}

var SiloedLockReleaseTokenPoolABI = SiloedLockReleaseTokenPoolMetaData.ABI

var SiloedLockReleaseTokenPoolBin = SiloedLockReleaseTokenPoolMetaData.Bin

func DeploySiloedLockReleaseTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address, lockBox common.Address) (common.Address, *types.Transaction, *SiloedLockReleaseTokenPool, error) {
	parsed, err := SiloedLockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SiloedLockReleaseTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router, lockBox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SiloedLockReleaseTokenPool{address: address, abi: *parsed, SiloedLockReleaseTokenPoolCaller: SiloedLockReleaseTokenPoolCaller{contract: contract}, SiloedLockReleaseTokenPoolTransactor: SiloedLockReleaseTokenPoolTransactor{contract: contract}, SiloedLockReleaseTokenPoolFilterer: SiloedLockReleaseTokenPoolFilterer{contract: contract}}, nil
}

type SiloedLockReleaseTokenPool struct {
	address common.Address
	abi     abi.ABI
	SiloedLockReleaseTokenPoolCaller
	SiloedLockReleaseTokenPoolTransactor
	SiloedLockReleaseTokenPoolFilterer
}

type SiloedLockReleaseTokenPoolCaller struct {
	contract *bind.BoundContract
}

type SiloedLockReleaseTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type SiloedLockReleaseTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type SiloedLockReleaseTokenPoolSession struct {
	Contract     *SiloedLockReleaseTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SiloedLockReleaseTokenPoolCallerSession struct {
	Contract *SiloedLockReleaseTokenPoolCaller
	CallOpts bind.CallOpts
}

type SiloedLockReleaseTokenPoolTransactorSession struct {
	Contract     *SiloedLockReleaseTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type SiloedLockReleaseTokenPoolRaw struct {
	Contract *SiloedLockReleaseTokenPool
}

type SiloedLockReleaseTokenPoolCallerRaw struct {
	Contract *SiloedLockReleaseTokenPoolCaller
}

type SiloedLockReleaseTokenPoolTransactorRaw struct {
	Contract *SiloedLockReleaseTokenPoolTransactor
}

func NewSiloedLockReleaseTokenPool(address common.Address, backend bind.ContractBackend) (*SiloedLockReleaseTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(SiloedLockReleaseTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSiloedLockReleaseTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPool{address: address, abi: abi, SiloedLockReleaseTokenPoolCaller: SiloedLockReleaseTokenPoolCaller{contract: contract}, SiloedLockReleaseTokenPoolTransactor: SiloedLockReleaseTokenPoolTransactor{contract: contract}, SiloedLockReleaseTokenPoolFilterer: SiloedLockReleaseTokenPoolFilterer{contract: contract}}, nil
}

func NewSiloedLockReleaseTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*SiloedLockReleaseTokenPoolCaller, error) {
	contract, err := bindSiloedLockReleaseTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCaller{contract: contract}, nil
}

func NewSiloedLockReleaseTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloedLockReleaseTokenPoolTransactor, error) {
	contract, err := bindSiloedLockReleaseTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolTransactor{contract: contract}, nil
}

func NewSiloedLockReleaseTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloedLockReleaseTokenPoolFilterer, error) {
	contract, err := bindSiloedLockReleaseTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolFilterer{contract: contract}, nil
}

func bindSiloedLockReleaseTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloedLockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloedLockReleaseTokenPool.Contract.SiloedLockReleaseTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SiloedLockReleaseTokenPoolTransactor.contract.Transfer(opts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SiloedLockReleaseTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloedLockReleaseTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.contract.Transfer(opts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAllowList(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAllowList(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAllowListEnabled(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAllowListEnabled(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getAvailableTokens", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetAvailableTokens(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAvailableTokens(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetAvailableTokens(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAvailableTokens(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getChainRebalancer", remoteChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetChainRebalancer(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetChainRebalancer(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetChainRebalancer(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetChainRebalancer(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentInboundRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentInboundRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRateLimitAdmin(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRateLimitAdmin(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRebalancer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRebalancer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRebalancer() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRebalancer(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRebalancer() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRebalancer(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRemotePools(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRemotePools(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRemoteToken(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRemoteToken(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRmnProxy(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRmnProxy(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRouter() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRouter(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRouter(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetSupportedChains(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetSupportedChains(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetToken() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetToken(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetToken(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetTokenDecimals(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetTokenDecimals(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetUnsiloedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getUnsiloedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetUnsiloedLiquidity() (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetUnsiloedLiquidity(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetUnsiloedLiquidity() (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetUnsiloedLiquidity(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsRemotePool(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsRemotePool(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) IsSiloed(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "isSiloed", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) IsSiloed(remoteChainSelector uint64) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSiloed(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) IsSiloed(remoteChainSelector uint64) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSiloed(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSupportedChain(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSupportedChain(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSupportedToken(&_SiloedLockReleaseTokenPool.CallOpts, token)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSupportedToken(&_SiloedLockReleaseTokenPool.CallOpts, token)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) Owner() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.Owner(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) Owner() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.Owner(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.SupportsInterface(&_SiloedLockReleaseTokenPool.CallOpts, interfaceId)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.SupportsInterface(&_SiloedLockReleaseTokenPool.CallOpts, interfaceId)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) TypeAndVersion() (string, error) {
	return _SiloedLockReleaseTokenPool.Contract.TypeAndVersion(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _SiloedLockReleaseTokenPool.Contract.TypeAndVersion(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.AcceptOwnership(&_SiloedLockReleaseTokenPool.TransactOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.AcceptOwnership(&_SiloedLockReleaseTokenPool.TransactOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.AddRemotePool(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.AddRemotePool(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyAllowListUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyAllowListUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyChainUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyChainUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "provideLiquidity", amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ProvideLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ProvideLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "provideSiloedLiquidity", remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ProvideSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ProvideSiloedLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ProvideSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ProvideSiloedLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.RemoveRemotePool(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.RemoveRemotePool(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetChainRateLimiterConfig(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetChainRateLimiterConfig(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetChainRateLimiterConfigs(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetChainRateLimiterConfigs(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRateLimitAdmin(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRateLimitAdmin(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setRebalancer", newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRebalancer(&_SiloedLockReleaseTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRebalancer(&_SiloedLockReleaseTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRouter(&_SiloedLockReleaseTokenPool.TransactOpts, newRouter)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRouter(&_SiloedLockReleaseTokenPool.TransactOpts, newRouter)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setSiloRebalancer", remoteChainSelector, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetSiloRebalancer(remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetSiloRebalancer(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetSiloRebalancer(remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetSiloRebalancer(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.TransferOwnership(&_SiloedLockReleaseTokenPool.TransactOpts, to)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.TransferOwnership(&_SiloedLockReleaseTokenPool.TransactOpts, to)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) UpdateSiloDesignations(opts *bind.TransactOpts, removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "updateSiloDesignations", removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) UpdateSiloDesignations(removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.UpdateSiloDesignations(&_SiloedLockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) UpdateSiloDesignations(removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.UpdateSiloDesignations(&_SiloedLockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) WithdrawSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "withdrawSiloedLiquidity", remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) WithdrawSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawSiloedLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) WithdrawSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawSiloedLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, amount)
}

type SiloedLockReleaseTokenPoolAllowListAddIterator struct {
	Event *SiloedLockReleaseTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolAllowListAdd)
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
		it.Event = new(SiloedLockReleaseTokenPoolAllowListAdd)
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

func (it *SiloedLockReleaseTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolAllowListAddIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolAllowListAdd)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*SiloedLockReleaseTokenPoolAllowListAdd, error) {
	event := new(SiloedLockReleaseTokenPoolAllowListAdd)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolAllowListRemoveIterator struct {
	Event *SiloedLockReleaseTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolAllowListRemove)
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
		it.Event = new(SiloedLockReleaseTokenPoolAllowListRemove)
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

func (it *SiloedLockReleaseTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolAllowListRemoveIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolAllowListRemove)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*SiloedLockReleaseTokenPoolAllowListRemove, error) {
	event := new(SiloedLockReleaseTokenPoolAllowListRemove)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainAddedIterator struct {
	Event *SiloedLockReleaseTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainAdded)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainAdded)
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

func (it *SiloedLockReleaseTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainAddedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainAddedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainAdded)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainAdded(log types.Log) (*SiloedLockReleaseTokenPoolChainAdded, error) {
	event := new(SiloedLockReleaseTokenPoolChainAdded)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainConfiguredIterator struct {
	Event *SiloedLockReleaseTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainConfigured)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainConfigured)
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

func (it *SiloedLockReleaseTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainConfiguredIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainConfigured)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainConfigured(log types.Log) (*SiloedLockReleaseTokenPoolChainConfigured, error) {
	event := new(SiloedLockReleaseTokenPoolChainConfigured)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainRemovedIterator struct {
	Event *SiloedLockReleaseTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainRemoved)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainRemoved)
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

func (it *SiloedLockReleaseTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainRemovedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainRemoved)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainRemoved(log types.Log) (*SiloedLockReleaseTokenPoolChainRemoved, error) {
	event := new(SiloedLockReleaseTokenPoolChainRemoved)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainSiloedIterator struct {
	Event *SiloedLockReleaseTokenPoolChainSiloed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainSiloedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainSiloed)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainSiloed)
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

func (it *SiloedLockReleaseTokenPoolChainSiloedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainSiloedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainSiloed struct {
	RemoteChainSelector uint64
	Rebalancer          common.Address
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainSiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainSiloedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainSiloed")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainSiloedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainSiloed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainSiloed) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainSiloed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainSiloed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainSiloed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainSiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainSiloed, error) {
	event := new(SiloedLockReleaseTokenPoolChainSiloed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainSiloed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainUnsiloedIterator struct {
	Event *SiloedLockReleaseTokenPoolChainUnsiloed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainUnsiloedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainUnsiloed)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainUnsiloed)
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

func (it *SiloedLockReleaseTokenPoolChainUnsiloedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainUnsiloedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainUnsiloed struct {
	RemoteChainSelector uint64
	AmountUnsiloed      *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainUnsiloedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainUnsiloed")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainUnsiloedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainUnsiloed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainUnsiloed) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainUnsiloed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainUnsiloed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainUnsiloed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainUnsiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainUnsiloed, error) {
	event := new(SiloedLockReleaseTokenPoolChainUnsiloed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainUnsiloed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolConfigChangedIterator struct {
	Event *SiloedLockReleaseTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolConfigChanged)
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
		it.Event = new(SiloedLockReleaseTokenPoolConfigChanged)
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

func (it *SiloedLockReleaseTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolConfigChangedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolConfigChanged)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseConfigChanged(log types.Log) (*SiloedLockReleaseTokenPoolConfigChanged, error) {
	event := new(SiloedLockReleaseTokenPoolConfigChanged)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(SiloedLockReleaseTokenPoolInboundRateLimitConsumed)
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

func (it *SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolInboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolInboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolInboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolLiquidityAddedIterator struct {
	Event *SiloedLockReleaseTokenPoolLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolLiquidityAdded)
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
		it.Event = new(SiloedLockReleaseTokenPoolLiquidityAdded)
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

func (it *SiloedLockReleaseTokenPoolLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolLiquidityAdded struct {
	RemoteChainSelector uint64
	Provider            common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*SiloedLockReleaseTokenPoolLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolLiquidityAddedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLiquidityAdded, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolLiquidityAdded)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseLiquidityAdded(log types.Log) (*SiloedLockReleaseTokenPoolLiquidityAdded, error) {
	event := new(SiloedLockReleaseTokenPoolLiquidityAdded)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolLiquidityRemovedIterator struct {
	Event *SiloedLockReleaseTokenPoolLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolLiquidityRemoved)
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
		it.Event = new(SiloedLockReleaseTokenPoolLiquidityRemoved)
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

func (it *SiloedLockReleaseTokenPoolLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolLiquidityRemoved struct {
	RemoteChainSelector uint64
	Remover             common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, remover []common.Address) (*SiloedLockReleaseTokenPoolLiquidityRemovedIterator, error) {

	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityRemoved", removerRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolLiquidityRemovedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLiquidityRemoved, remover []common.Address) (event.Subscription, error) {

	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityRemoved", removerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolLiquidityRemoved)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseLiquidityRemoved(log types.Log) (*SiloedLockReleaseTokenPoolLiquidityRemoved, error) {
	event := new(SiloedLockReleaseTokenPoolLiquidityRemoved)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolLockedOrBurnedIterator struct {
	Event *SiloedLockReleaseTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolLockedOrBurned)
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
		it.Event = new(SiloedLockReleaseTokenPoolLockedOrBurned)
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

func (it *SiloedLockReleaseTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolLockedOrBurnedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolLockedOrBurned)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*SiloedLockReleaseTokenPoolLockedOrBurned, error) {
	event := new(SiloedLockReleaseTokenPoolLockedOrBurned)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(SiloedLockReleaseTokenPoolOutboundRateLimitConsumed)
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

func (it *SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolOutboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolOutboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolOutboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator struct {
	Event *SiloedLockReleaseTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolOwnershipTransferRequested)
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
		it.Event = new(SiloedLockReleaseTokenPoolOwnershipTransferRequested)
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

func (it *SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolOwnershipTransferRequested)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*SiloedLockReleaseTokenPoolOwnershipTransferRequested, error) {
	event := new(SiloedLockReleaseTokenPoolOwnershipTransferRequested)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolOwnershipTransferredIterator struct {
	Event *SiloedLockReleaseTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolOwnershipTransferred)
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
		it.Event = new(SiloedLockReleaseTokenPoolOwnershipTransferred)
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

func (it *SiloedLockReleaseTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedLockReleaseTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolOwnershipTransferredIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolOwnershipTransferred)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*SiloedLockReleaseTokenPoolOwnershipTransferred, error) {
	event := new(SiloedLockReleaseTokenPoolOwnershipTransferred)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolRateLimitAdminSetIterator struct {
	Event *SiloedLockReleaseTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolRateLimitAdminSet)
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
		it.Event = new(SiloedLockReleaseTokenPoolRateLimitAdminSet)
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

func (it *SiloedLockReleaseTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolRateLimitAdminSetIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolRateLimitAdminSet)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*SiloedLockReleaseTokenPoolRateLimitAdminSet, error) {
	event := new(SiloedLockReleaseTokenPoolRateLimitAdminSet)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolReleasedOrMintedIterator struct {
	Event *SiloedLockReleaseTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolReleasedOrMinted)
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
		it.Event = new(SiloedLockReleaseTokenPoolReleasedOrMinted)
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

func (it *SiloedLockReleaseTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolReleasedOrMintedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolReleasedOrMinted)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*SiloedLockReleaseTokenPoolReleasedOrMinted, error) {
	event := new(SiloedLockReleaseTokenPoolReleasedOrMinted)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolRemotePoolAddedIterator struct {
	Event *SiloedLockReleaseTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolRemotePoolAdded)
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
		it.Event = new(SiloedLockReleaseTokenPoolRemotePoolAdded)
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

func (it *SiloedLockReleaseTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolRemotePoolAddedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolRemotePoolAdded)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*SiloedLockReleaseTokenPoolRemotePoolAdded, error) {
	event := new(SiloedLockReleaseTokenPoolRemotePoolAdded)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolRemotePoolRemovedIterator struct {
	Event *SiloedLockReleaseTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolRemotePoolRemoved)
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
		it.Event = new(SiloedLockReleaseTokenPoolRemotePoolRemoved)
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

func (it *SiloedLockReleaseTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolRemotePoolRemovedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolRemotePoolRemoved)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*SiloedLockReleaseTokenPoolRemotePoolRemoved, error) {
	event := new(SiloedLockReleaseTokenPoolRemotePoolRemoved)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolRouterUpdatedIterator struct {
	Event *SiloedLockReleaseTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolRouterUpdated)
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
		it.Event = new(SiloedLockReleaseTokenPoolRouterUpdated)
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

func (it *SiloedLockReleaseTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolRouterUpdatedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolRouterUpdated)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*SiloedLockReleaseTokenPoolRouterUpdated, error) {
	event := new(SiloedLockReleaseTokenPoolRouterUpdated)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolSiloRebalancerSetIterator struct {
	Event *SiloedLockReleaseTokenPoolSiloRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolSiloRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolSiloRebalancerSet)
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
		it.Event = new(SiloedLockReleaseTokenPoolSiloRebalancerSet)
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

func (it *SiloedLockReleaseTokenPoolSiloRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolSiloRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolSiloRebalancerSet struct {
	RemoteChainSelector uint64
	OldRebalancer       common.Address
	NewRebalancer       common.Address
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterSiloRebalancerSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolSiloRebalancerSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "SiloRebalancerSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolSiloRebalancerSetIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "SiloRebalancerSet", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchSiloRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolSiloRebalancerSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "SiloRebalancerSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolSiloRebalancerSet)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "SiloRebalancerSet", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseSiloRebalancerSet(log types.Log) (*SiloedLockReleaseTokenPoolSiloRebalancerSet, error) {
	event := new(SiloedLockReleaseTokenPoolSiloRebalancerSet)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "SiloRebalancerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator struct {
	Event *SiloedLockReleaseTokenPoolUnsiloedRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolUnsiloedRebalancerSet)
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
		it.Event = new(SiloedLockReleaseTokenPoolUnsiloedRebalancerSet)
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

func (it *SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolUnsiloedRebalancerSet struct {
	OldRebalancer common.Address
	NewRebalancer common.Address
	Raw           types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterUnsiloedRebalancerSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "UnsiloedRebalancerSet")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "UnsiloedRebalancerSet", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchUnsiloedRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolUnsiloedRebalancerSet) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "UnsiloedRebalancerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolUnsiloedRebalancerSet)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "UnsiloedRebalancerSet", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseUnsiloedRebalancerSet(log types.Log) (*SiloedLockReleaseTokenPoolUnsiloedRebalancerSet, error) {
	event := new(SiloedLockReleaseTokenPoolUnsiloedRebalancerSet)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "UnsiloedRebalancerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (SiloedLockReleaseTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (SiloedLockReleaseTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (SiloedLockReleaseTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (SiloedLockReleaseTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (SiloedLockReleaseTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (SiloedLockReleaseTokenPoolChainSiloed) Topic() common.Hash {
	return common.HexToHash("0x180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb")
}

func (SiloedLockReleaseTokenPoolChainUnsiloed) Topic() common.Hash {
	return common.HexToHash("0x7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93")
}

func (SiloedLockReleaseTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (SiloedLockReleaseTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (SiloedLockReleaseTokenPoolLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0x569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf")
}

func (SiloedLockReleaseTokenPoolLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0x58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e")
}

func (SiloedLockReleaseTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (SiloedLockReleaseTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (SiloedLockReleaseTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SiloedLockReleaseTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (SiloedLockReleaseTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (SiloedLockReleaseTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (SiloedLockReleaseTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (SiloedLockReleaseTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (SiloedLockReleaseTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (SiloedLockReleaseTokenPoolSiloRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f")
}

func (SiloedLockReleaseTokenPoolUnsiloedRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b234")
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPool) Address() common.Address {
	return _SiloedLockReleaseTokenPool.address
}

type SiloedLockReleaseTokenPoolInterface interface {
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetUnsiloedLiquidity(opts *bind.CallOpts) (*big.Int, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSiloed(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

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

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateSiloDesignations(opts *bind.TransactOpts, removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	WithdrawSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*SiloedLockReleaseTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*SiloedLockReleaseTokenPoolAllowListRemove, error)

	FilterChainAdded(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*SiloedLockReleaseTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*SiloedLockReleaseTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*SiloedLockReleaseTokenPoolChainRemoved, error)

	FilterChainSiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainSiloedIterator, error)

	WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainSiloed) (event.Subscription, error)

	ParseChainSiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainSiloed, error)

	FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainUnsiloedIterator, error)

	WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainUnsiloed) (event.Subscription, error)

	ParseChainUnsiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainUnsiloed, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*SiloedLockReleaseTokenPoolConfigChanged, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolInboundRateLimitConsumed, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*SiloedLockReleaseTokenPoolLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLiquidityAdded, provider []common.Address) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*SiloedLockReleaseTokenPoolLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, remover []common.Address) (*SiloedLockReleaseTokenPoolLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLiquidityRemoved, remover []common.Address) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*SiloedLockReleaseTokenPoolLiquidityRemoved, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*SiloedLockReleaseTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SiloedLockReleaseTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedLockReleaseTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SiloedLockReleaseTokenPoolOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*SiloedLockReleaseTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*SiloedLockReleaseTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*SiloedLockReleaseTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*SiloedLockReleaseTokenPoolRemotePoolRemoved, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*SiloedLockReleaseTokenPoolRouterUpdated, error)

	FilterSiloRebalancerSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolSiloRebalancerSetIterator, error)

	WatchSiloRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolSiloRebalancerSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseSiloRebalancerSet(log types.Log) (*SiloedLockReleaseTokenPoolSiloRebalancerSet, error)

	FilterUnsiloedRebalancerSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator, error)

	WatchUnsiloedRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolUnsiloedRebalancerSet) (event.Subscription, error)

	ParseUnsiloedRebalancerSet(log types.Log) (*SiloedLockReleaseTokenPoolUnsiloedRebalancerSet, error)

	Address() common.Address
}
