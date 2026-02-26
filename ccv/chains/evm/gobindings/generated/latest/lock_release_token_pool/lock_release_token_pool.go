// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lock_release_token_pool

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

type TokenPoolRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmations  bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var LockReleaseTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockBox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmations\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmations\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationsInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationsOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationsSet\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmations\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461036e5760c081615dd8803803809161002082856103ae565b83398101031261036e5780516001600160a01b0381169182820361036e5761004a602082016103e7565b90610057604082016103f5565b610063606083016103f5565b9261007c60a0610075608086016103f5565b94016103f5565b93331561039d57600180546001600160a01b031916331790558615801561038c575b801561037b575b6102fe57608086905260c05260405163313ce56760e01b81526020816004818a5afa6000918161033a575b5061030f575b5060a052600380546001600160a01b03199081166001600160a01b03938416179091556002805490911692821692909217909155169182156102fe576040516375151b6360e01b815260048101829052602081602481875afa9081156102f2576000916102b0575b501561029c57604051906020600081840163095ea7b360e01b815286602486015281196044860152604485526101756064866103ae565b84519082875af1903d600051908361027d575b50505015610238575b8260e052604051615973908161046582396080518181816117490152818161222401528181612b4101528181612e260152818161332a01528181613a2b01528181613a850152614f50015260a0518181816138f101528181614a9d01528181614ae70152614fa4015260c051818181610b68015281816117e4015281816122be01528181612bdc01526133c5015260e05181818161278601528181612dad0152614ed50152f35b6102769161027160405163095ea7b360e01b6020820152856024820152600060448201526044815261026b6064826103ae565b82610409565b610409565b3880610191565b9192509061029257503b15155b388080610188565b600191501461028a565b63961c9a4f60e01b60005260045260246000fd5b6020813d6020116102ea575b816102c9602093836103ae565b810103126102e657519081151582036102e357503861013e565b80fd5b5080fd5b3d91506102bc565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b60ff1660ff821681810361032357506100d6565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610373575b81610356602093836103ae565b8101031261036e57610367906103e7565b90386100d0565b600080fd5b3d9150610349565b506001600160a01b038116156100a5565b506001600160a01b0384161561009e565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b038211908210176103d157604052565b634e487b7160e01b600052604160045260246000fd5b519060ff8216820361036e57565b51906001600160a01b038216820361036e57565b906000602091828151910182855af1156102f2576000513d61045b57506001600160a01b0381163b155b61043a5750565b635274afe760e01b60009081526001600160a01b0391909116600452602490fd5b6001141561043356fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613b0a57508063181f5a7714613aa957806321df0da714613a58578063240028e8146139f45780632422ac451461391557806324f65ee7146138d75780632c0634041461383e57806337a3210d1461380a578063390775371461328b578063489a68f214612a9a5780634c5ef0ed14612a5357806362ddd3c4146129cc5780637437ff9f1461297e57806379ba5097146128b75780638926f54f1461287157806389720a62146127aa5780638c6894fb146127595780638da5cb5b146127255780639a4575b9146121ab578063a42a7b8b14612044578063acfecf9114611f4c578063ae39a25714611dc1578063b1c71c65146116a0578063b794658014611663578063b8d5005e1461163e578063bfeffd3f14611592578063c4bffe2b14611467578063c7230a60146111b6578063d4d6de2314611117578063d8aa3f4014610fdd578063dc04fa1f14610b8c578063dc0bd97114610b3b578063dcbd41bc14610937578063e8a1da17146102735763f2fde38b146101a257600080fd5b346102705760206003193601126102705773ffffffffffffffffffffffffffffffffffffffff6101d0613c6b565b6101d8614c09565b1633811461024857807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102705760406003193601126102705760043567ffffffffffffffff8111610790576102a5903690600401613f7b565b9060243567ffffffffffffffff811161093357906102c884923690600401613f7b565b9390916102d3614c09565b83905b8282106107985750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610794578060051b8301358581121561078c5783016101208136031261078c576040519461033a86613e2c565b61034382613d28565b8652602082013567ffffffffffffffff81116107905782019436601f87011215610790578535956103738761433a565b966103816040519889613e48565b80885260208089019160051b8301019036821161078c5760208301905b828210610759575050505060208701958652604083013567ffffffffffffffff8111610755576103d19036908501613eee565b91604088019283526103fb6103e936606087016148f1565b9460608a0195865260c03691016148f1565b95608089019687528351511561072d5761041f67ffffffffffffffff8a51166155f5565b156106f65767ffffffffffffffff895116825260086020526040822061044686518261503b565b61045488516002830161503b565b6004855191019080519067ffffffffffffffff82116106c957610477835461473f565b601f811161068e575b50602090601f83116001146105ef576104ce92918691836105e4575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b8851805182101561050857906105026001926104fb8367ffffffffffffffff8f5116926146fc565b5190614c54565b016104d3565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29391999750956105d667ffffffffffffffff60019796949851169251935191516105a261056d60405196879687526101006020880152610100870190613c0c565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610309565b015190508e8061049c565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b818110610676575090846001959493921061063f575b505050811b0190556104d1565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080610632565b9293602060018192878601518155019501930161061c565b6106b99084875260208720601f850160051c810191602086106106bf575b601f0160051c019061498d565b8d610480565b90915081906106ac565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116107885760209161077d8392833691890101613eee565b81520191019061039e565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff6107ba6107b58486889a9699979a6148c4565b6142e8565b16916107c58361532b565b156109075782845260086020526107e1600560408620016152c8565b94845b865181101561081a5760019085875260086020526108136005604089200161080c838b6146fc565b51906154c1565b50016107e4565b5093969290945094909480875260086020526005604088208881558860018201558860028201558860038201558860048201610856815461473f565b806108c6575b50505001805490888155816108a8575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019091949392946102d6565b885260208820908101905b8181101561086c578881556001016108b3565b601f81116001146108dc5750555b888a8061085c565b818352602083206108f791601f01861c81019060010161498d565b80825281602081209155556108d4565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346102705760206003193601126102705760043567ffffffffffffffff811161079057610969903690600401613fac565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580610b19575b610aed57825b81811061099c578380f35b6109a7818385614876565b67ffffffffffffffff6109b9826142e8565b16906109d2826000526007602052604060002054151590565b15610ac157907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610a81610a5b602060019897018b610a1382614886565b15610a88578790526004602052610a3a60408d20610a3436604088016148f1565b9061503b565b868c526005602052610a5660408d20610a343660a088016148f1565b614886565b916040519215158352610a746020840160408301614949565b60a0608084019101614949565ba201610991565b60026040828a610a5694526008602052610aaa828220610a3436858c016148f1565b8a815260086020522001610a343660a088016148f1565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561098b565b5034610270578060031936011261027057602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102705760406003193601126102705760043567ffffffffffffffff811161079057610bbe903690600401613fac565b60243567ffffffffffffffff811161093357610bde903690600401613f7b565b919092610be9614c09565b845b828110610c5557505050825b818110610c02578380f35b8067ffffffffffffffff610c1c6107b560019486886148c4565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610bf7565b610c636107b5828585614876565b610c6e828585614876565b90602082019060e0830190610c8282614886565b15610fa85760a0840161271061ffff610c9a83614893565b161015610f995760c085019161271061ffff610cb585614893565b161015610f615763ffffffff610cca866148a2565b1615610f2c5767ffffffffffffffff1694858c52600b60205260408c20610cf0866148a2565b63ffffffff16908054906040840191610d08836148a2565b60201b67ffffffff0000000016936060860194610d24866148a2565b60401b6bffffffff0000000000000000169660800196610d43886148a2565b60601b6fffffffff0000000000000000000000001691610d628a614893565b60801b71ffff000000000000000000000000000000001693610d838c614893565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610e3687614886565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610e87906148b3565b63ffffffff168752610e98906148b3565b63ffffffff166020870152610eac906148b3565b63ffffffff166040860152610ec0906148b3565b63ffffffff166060850152610ed490613d6c565b61ffff166080840152610ee690613d6c565b61ffff1660a0830152610ef890613d3d565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610beb565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff610f7086614893565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff610f70602493614893565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b503461027057608060031936011261027057610ff7613c6b565b50611000613d11565b611008613d5b565b5060643567ffffffffffffffff8111610755579167ffffffffffffffff60409261103860e0953690600401613d7b565b50508260c0855161104881613e10565b82815282602082015282878201528260608201528260808201528260a08201520152168152600b602052206040519061108082613e10565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346102705760206003193601126102705760043561ffff811690818103610755577f46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a0191602091611166614c09565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006002549260a01b16911617600255604051908152a180f35b50346102705760406003193601126102705760043567ffffffffffffffff8111610790576111e8903690600401613f7b565b906111f1613cb6565b9173ffffffffffffffffffffffffffffffffffffffff6001541633141580611445575b6114195773ffffffffffffffffffffffffffffffffffffffff83169081156113f157845b818110611243578580f35b73ffffffffffffffffffffffffffffffffffffffff61126b6112668385886148c4565b6142c7565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156113e65788916113b3575b50806112c0575b5050600101611238565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff8a16602484015260448084018590528352918a9190611321606482613e48565b519082865af1156113a85787513d61139f5750813b155b6113735790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a390386112b6565b602488837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b60011415611338565b6040513d89823e3d90fd5b905060203d81116113df575b6113c98183613e48565b60208260009281010312610270575051386112af565b503d6113bf565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600c5416331415611214565b5034610270578060031936011261027057604051906006548083528260208101600684526020842092845b8181106115795750506114a792500383613e48565b81516114cb6114b58261433a565b916114c36040519384613e48565b80835261433a565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b845181101561152a578067ffffffffffffffff611517600193886146fc565b511661152382866146fc565b52016114f8565b50925090604051928392602084019060208552518091526040840192915b818110611556575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611548565b8454835260019485019487945060209093019201611492565b50346102705760206003193601126102705760043573ffffffffffffffffffffffffffffffffffffffff8116809103610790576115cd614c09565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a1161760035580f35b5034610270578060031936011261027057602061ffff60025460a01c16604051908152f35b50346102705760206003193601126102705761169c611688611683613cfa565b614854565b604051918291602083526020830190613c0c565b0390f35b5034610270576060600319360112610270576004359067ffffffffffffffff8211610270578160040160a06003198436030112610790576116df613d4a565b9160443567ffffffffffffffff8111610790579061170561172293923690600401613d7b565b939061170f6146e3565b5061171a8685614fd8565b943691613e89565b936084860194611731866142c7565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611d7757602487019677ffffffffffffffff00000000000000000000000000000000611797896142e8565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611cea578591611d48575b50611d205767ffffffffffffffff61182b896142e8565b16611843816000526007602052604060002054151590565b15611cf557602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611cea578590611c9d575b73ffffffffffffffffffffffffffffffffffffffff9150163303611c715760648101359461ffff6118d588886143cd565b9416938989828715611bf65750505061ffff60025460a01c168015611bce57808610611b9e5750897f996b829383cc7e136842d4c4c175083bcf4e20807c7432105c1b794ba885e7768267ffffffffffffffff61193a6119348e6142c7565b946142e8565b1692838a526004602052611952818360408d206156aa565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff600354169384611a74575b611a6a8a611a396116838e6119b08e8e6143cd565b936119c3856119be846142e8565b614ebe565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6119ff6119f9856142e8565b936142c7565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252336020830152810188905292169180606081015b0390a26142e8565b90611a42614f9d565b60405192611a4f84613df4565b83526020830152604051928392604084526040840190613f51565b9060208301520390f35b843b15610788578694928a949286928d604051998a98899788967f1ff7703e000000000000000000000000000000000000000000000000000000008852600488016080905280611ac391615278565b6084890160a09052610124890190611ada926143fb565b93611ae490613d28565b67ffffffffffffffff1660a4880152604401611aff90613cd9565b73ffffffffffffffffffffffffffffffffffffffff1660c48701528d60e4870152611b2990613cd9565b73ffffffffffffffffffffffffffffffffffffffff166101048601526024850152838103600319016044850152611b5f91613c0c565b90606483015203925af18015611b9357611b7e575b808080808061199b565b611b89828092613e48565b6102705780611b74565b6040513d84823e3d90fd5b86604491877f1f5b9f77000000000000000000000000000000000000000000000000000000008352600452602452fd5b6004877f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b67ffffffffffffffff611c2c6119347fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944946142c7565b1692838a526008602052611c44818360408d206156aa565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a261197b565b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611ce2575b81611cb760209383613e48565b8101031261078c57611cdd73ffffffffffffffffffffffffffffffffffffffff916143da565b6118a4565b3d9150611caa565b6040513d87823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611d6a915060203d602011611d70575b611d628183613e48565b810190614bf1565b38611814565b503d611d58565b60248373ffffffffffffffffffffffffffffffffffffffff611d98896142c7565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b503461027057606060031936011261027057611ddb613c6b565b90611de4613cb6565b6044359273ffffffffffffffffffffffffffffffffffffffff841680850361093357611e0e614c09565b73ffffffffffffffffffffffffffffffffffffffff82168015611f245794611f1e917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c556040519384938491604091949373ffffffffffffffffffffffffffffffffffffffff809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346102705767ffffffffffffffff611f6436613f0c565b929091611f6f614c09565b1691611f88836000526007602052604060002054151590565b15610907578284526008602052611fb760056040862001611faa368486613e89565b60208151910120906154c1565b15611ffc57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611ff66040519283926020845260208401916143fb565b0390a280f35b82612040836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916143fb565b0390fd5b50346102705760206003193601126102705767ffffffffffffffff612067613cfa565b168152600860205261207e600560408320016152c8565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120c36120ad8361433a565b926120bb6040519485613e48565b80845261433a565b01835b81811061219a575050825b825181101561211757806120e7600192856146fc565b51855260096020526120fb60408620614792565b61210582856146fc565b5261211081846146fc565b50016120d1565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061214f57505050500390f35b9193602061218a827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613c0c565b9601920192018594939192612140565b8060606020809386010152016120c6565b50346102705760206003193601126102705760043567ffffffffffffffff811161079057806004019060a06003198236030112610755576121ea6146e3565b506040516020936121fb8583613e48565b808252608483019161220c836142c7565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361270457602484019477ffffffffffffffff00000000000000000000000000000000612272876142e8565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152878160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156126895784916126e7575b506126bf5767ffffffffffffffff612305876142e8565b1661231d816000526007602052604060002054151590565b15612694578773ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015612689578490612641575b73ffffffffffffffffffffffffffffffffffffffff9150163303612615576064850135946123aa856142c7565b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448767ffffffffffffffff6123de8b6142e8565b169283885260088c526123f5818360408b206156aa565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a273ffffffffffffffffffffffffffffffffffffffff6003541691826124fa575b886124ca6116838a8a7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff8c61247e846119be876142e8565b611a3161249361248d876142e8565b926142c7565b6040805173ffffffffffffffffffffffffffffffffffffffff90921682523360208301528101959095529116929081906060820190565b906124d3614f9d565b604051926124e084613df4565b83528183015261169c604051928284938452830190613f51565b823b1561078c57918791858094604051968795869485937f1ff7703e00000000000000000000000000000000000000000000000000000000855260048501608090528061254691615278565b6084860160a0905261012486019061255d926143fb565b9161256790613d28565b67ffffffffffffffff1660a485015260440161258290613cd9565b73ffffffffffffffffffffffffffffffffffffffff1660c48401528b60e48401526125ac8b613cd9565b73ffffffffffffffffffffffffffffffffffffffff166101048401528360248401528281036003190160448401526125e391613c0c565b8a606483015203925af18015611b9357612600575b80808061243d565b61260b828092613e48565b61027057806125f8565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d8311612682575b6126578183613e48565b810103126109335761267d73ffffffffffffffffffffffffffffffffffffffff916143da565b61237d565b503d61264d565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6126fe9150883d8a11611d7057611d628183613e48565b386122ee565b5073ffffffffffffffffffffffffffffffffffffffff611d986024936142c7565b5034610270578060031936011261027057602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b5034610270578060031936011261027057602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102705760c0600319360112610270576127c4613c6b565b6127cc613d11565b9060643561ffff811681036109335760843567ffffffffffffffff811161078c576127fb903690600401613d7b565b9160a43593600285101561078857612816956044359161443a565b90604051918291602083016020845282518091526020604085019301915b818110612842575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101612834565b50346102705760206003193601126102705760206128ad67ffffffffffffffff612899613cfa565b166000526007602052604060002054151590565b6040519015158152f35b5034610270578060031936011261027057805473ffffffffffffffffffffffffffffffffffffffff81163303612956577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b5034610270578060031936011261027057600254600a54600c546040805173ffffffffffffffffffffffffffffffffffffffff94851681529284166020840152921691810191909152606090f35b5034610270576129db36613f0c565b6129e793929193614c09565b67ffffffffffffffff8216612a09816000526007602052604060002054151590565b15612a285750612a259293612a1f913691613e89565b90614c54565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b503461027057604060031936011261027057612a6d613cfa565b906024359067ffffffffffffffff82116102705760206128ad84612a943660048701613eee565b906142fd565b5034610270576040600319360112610270576004359067ffffffffffffffff8211610270578160040190610100600319843603011261027057612adb613d4a565b9181604051612ae981613da9565b5260648401359360c4810193612b1a612b14612b0f612b088887614276565b3691613e89565b614a29565b87614ae4565b946084830196612b29886142c7565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361326a57602484019477ffffffffffffffff00000000000000000000000000000000612b8f876142e8565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156113e657889161324b575b506132235767ffffffffffffffff612c23876142e8565b16612c3b816000526007602052604060002054151590565b156131f857602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156113e65788916131d9575b50156131ad57612cb2866142e8565b93612cc860a4870195612a94612b088886614276565b156131665761ffff169087878a8c85156130e157612d317f63335ad9e238acd0e9e6c1c20f529ffbea4cda73578c329a7aa7e9d61e5cdcc59383604067ffffffffffffffff612d19612d21966142c7565b9586946142e8565b16978881526005602052206156aa565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff600354169384612f0e575b50505050505060440191612d8d836142c7565b612d96836142e8565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b1561093357608484928367ffffffffffffffff9373ffffffffffffffffffffffffffffffffffffffff60405197889687957f74fd18ac000000000000000000000000000000000000000000000000000000008752837f00000000000000000000000000000000000000000000000000000000000000001660048801521660248601528c60448601521660648401525af18015611b9357612ef9575b5050608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff612ec5612ebf6119f97ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0976142e8565b966142c7565b816040519716875233898801521660408601528560608601521692a260405190612eee82613da9565b815260405190518152f35b612f04828092613e48565b6102705780612e6a565b843b156130dd57868995938c959387938b6040519a8b998a9889977f1abfe46e0000000000000000000000000000000000000000000000000000000089526004890160609052612f5e8780615278565b60648b0161010090526101648b0190612f76926143fb565b94612f8090613d28565b67ffffffffffffffff1660848a0152604401612f9b90613cd9565b73ffffffffffffffffffffffffffffffffffffffff1660a489015260c4880152612fc490613cd9565b73ffffffffffffffffffffffffffffffffffffffff1660e4870152612fe99084615278565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8784030161010488015261301e92916143fb565b906130299083615278565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161012487015261305e92916143fb565b9060e48b0161306c91615278565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101448601526130a192916143fb565b908c6024840152604483015203925af18015612689579084916130c8575b80808080612d7a565b816130d291613e48565b6107555782386130bf565b8880fd5b6131397f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c93836002604067ffffffffffffffff613120613128976142c7565b9687956142e8565b1698898152600860205220016156aa565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612d5a565b6131708583614276565b6120406040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916143fb565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b6131f2915060203d602011611d7057611d628183613e48565b38612ca3565b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b613264915060203d602011611d7057611d628183613e48565b38612c0c565b60248673ffffffffffffffffffffffffffffffffffffffff611d988b6142c7565b5034610270576020600319360112610270576004359067ffffffffffffffff8211610270578160040190610100600319843603011261027057806040516132d181613da9565b52806040516132df81613da9565b52606483013560c48401936133036132fd612b0f612b088888614276565b83614ae4565b936084820195613312876142c7565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036137e957602483019377ffffffffffffffff00000000000000000000000000000000613378866142e8565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156113a85787916137ca575b506137a25767ffffffffffffffff61340c866142e8565b16613424816000526007602052604060002054151590565b1561377757602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156113a8578791613758575b501561372c5761349b856142e8565b926134b160a4860194612a94612b088785614276565b1561372257867f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c896131286134f88a838f604067ffffffffffffffff6131206002936142c7565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a273ffffffffffffffffffffffffffffffffffffffff600354169283613552575b505050505060440191612d8d836142c7565b833b1561371e57878795938195938c93604051988997889687957f1abfe46e00000000000000000000000000000000000000000000000000000000875260048701606090528d6135a28780615278565b60648a0161010090526101648a01906135ba926143fb565b946135c490613d28565b67ffffffffffffffff1660848901526044016135df90613cd9565b73ffffffffffffffffffffffffffffffffffffffff1660a488015260c487015261360890613cd9565b73ffffffffffffffffffffffffffffffffffffffff1660e486015261362d9084615278565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161010487015261366292916143fb565b9061366d9083615278565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101248601526136a292916143fb565b9060e48a016136b091615278565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c848403016101448501526136e592916143fb565b8b602483015282604483015203925af1801561268957613709575b80808080613540565b926137178160449395613e48565b9290613700565b8780fd5b8361317091614276565b6024867f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b613771915060203d602011611d7057611d628183613e48565b3861348c565b7fa9902c7e000000000000000000000000000000000000000000000000000000008752600452602486fd5b6004867f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6137e3915060203d602011611d7057611d628183613e48565b386133f5565b60248573ffffffffffffffffffffffffffffffffffffffff611d988a6142c7565b5034610270578060031936011261027057602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b50346102705760c060031936011261027057613858613c6b565b50613861613d11565b613869613c93565b506084359161ffff831683036102705760a4359067ffffffffffffffff82116102705760a063ffffffff8061ffff6138b088886138a93660048b01613d7b565b50506140f1565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b5034610270578060031936011261027057602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102705760406003193601126102705761392f613cfa565b602435918215158303610270576101406139f261394c858561406e565b6139a260409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b503461027057602060031936011261027057602090613a11613c6b565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b5034610270578060031936011261027057602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102705780600319360112610270575061169c604051613acc604082613e48565b601e81527f4c6f636b52656c65617365546f6b656e506f6f6c20322e302e302d64657600006020820152604051918291602083526020830190613c0c565b905034610790576020600319360112610790576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361075557602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115613be2575b8115613bb8575b8115613b8e575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613b87565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613b80565b7f331710310000000000000000000000000000000000000000000000000000000081149150613b79565b919082519283825260005b848110613c565750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613c17565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203613c8e57565b600080fd5b6064359073ffffffffffffffffffffffffffffffffffffffff82168203613c8e57565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203613c8e57565b359073ffffffffffffffffffffffffffffffffffffffff82168203613c8e57565b6004359067ffffffffffffffff82168203613c8e57565b6024359067ffffffffffffffff82168203613c8e57565b359067ffffffffffffffff82168203613c8e57565b35908115158203613c8e57565b6024359061ffff82168203613c8e57565b6044359061ffff82168203613c8e57565b359061ffff82168203613c8e57565b9181601f84011215613c8e5782359167ffffffffffffffff8311613c8e5760208381860195010111613c8e57565b6020810190811067ffffffffffffffff821117613dc557604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117613dc557604052565b60e0810190811067ffffffffffffffff821117613dc557604052565b60a0810190811067ffffffffffffffff821117613dc557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613dc557604052565b92919267ffffffffffffffff8211613dc55760405191613ed1601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184613e48565b829481845281830111613c8e578281602093846000960137010152565b9080601f83011215613c8e57816020613f0993359101613e89565b90565b906040600319830112613c8e5760043567ffffffffffffffff81168103613c8e57916024359067ffffffffffffffff8211613c8e57613f4d91600401613d7b565b9091565b613f09916020613f6a8351604084526040840190613c0c565b920151906020818403910152613c0c565b9181601f84011215613c8e5782359167ffffffffffffffff8311613c8e576020808501948460051b010111613c8e57565b9181601f84011215613c8e5782359167ffffffffffffffff8311613c8e576020808501948460081b010111613c8e57565b60405190613fea82613e2c565b60006080838281528260208201528260408201528260608201520152565b9060405161401581613e2c565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff91614080613fdd565b50614089613fdd565b506140bd57166000526008602052604060002090613f096140b160026140b66140b186614008565b6149a4565b9401614008565b16908160005260046020526140d86140b16040600020614008565b916000526005602052613f096140b16040600020614008565b9061ffff8060025460a01c169116928315159283809461426e575b6142445767ffffffffffffffff16600052600b6020526040600020916040519261413584613e10565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a0152614229576141ca575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b8193975080929450106141f957505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f1f5b9f770000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b50821561410c565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613c8e570180359067ffffffffffffffff8211613c8e57602001918136038313613c8e57565b3573ffffffffffffffffffffffffffffffffffffffff81168103613c8e5790565b3567ffffffffffffffff81168103613c8e5790565b9067ffffffffffffffff613f0992166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613dc55760051b60200190565b8181029291811591840414171561436557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b811561439e570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b9190820391821161436557565b519073ffffffffffffffffffffffffffffffffffffffff82168203613c8e57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b929593909473ffffffffffffffffffffffffffffffffffffffff600354169586156146c15760028610156146925773ffffffffffffffffffffffffffffffffffffffff976144e89461ffff9288156145b0575b67ffffffffffffffff906040519b8c997f89720a62000000000000000000000000000000000000000000000000000000008b521660048a0152166024880152604487015216606485015260c0608485015260c48401916143fb565b928180600095869560a483015203915afa9182156145a357819261450b57505090565b9091503d8083833e61451d8183613e48565b8101906020818303126107555780519067ffffffffffffffff8211610933570181601f82011215610755578051906145548261433a565b936145626040519586613e48565b82855260208086019360051b8301019384116102705750602001905b82821061458b5750505090565b60208091614598846143da565b81520191019061457e565b50604051903d90823e3d90fd5b67ffffffffffffffff8116600052600b6020526040600020604051906145d582613e10565b549063ffffffff8216815263ffffffff8260201c16602082015263ffffffff8260401c16604082015263ffffffff8260601c16606082015260c0868360801c169182608082015260ff60a0820194898160901c16865260a01c1615159182910152614642575b505061448d565b919267ffffffffffffffff928587161561467a575061271061466a8761467194511683614352565b04906143cd565b915b903861463b565b61468c925061466a6127109183614352565b91614673565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b50505050505050506040516146d7602082613e48565b60008152600036813790565b604051906146f082613df4565b60606020838281520152565b80518210156147105760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015614788575b602083101461475957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161474e565b90604051918260008254926147a68461473f565b808452936001811690811561481457506001146147cd575b506147cb92500383613e48565b565b90506000929192526020600020906000915b8183106147f85750509060206147cb92820101386147be565b60209193508060019154838589010152019101909184926147df565b602093506147cb9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386147be565b67ffffffffffffffff166000526008602052613f096004604060002001614792565b91908110156147105760081b0190565b358015158103613c8e5790565b3561ffff81168103613c8e5790565b3563ffffffff81168103613c8e5790565b359063ffffffff82168203613c8e57565b91908110156147105760051b0190565b35906fffffffffffffffffffffffffffffffff82168203613c8e57565b9190826060910312613c8e576040516060810181811067ffffffffffffffff821117613dc557604052604061494481839561492b81613d3d565b8552614939602082016148d4565b6020860152016148d4565b910152565b6fffffffffffffffffffffffffffffffff6149876040809361496a81613d3d565b151586528361497b602083016148d4565b166020870152016148d4565b16910152565b818110614998575050565b6000815560010161498d565b6149ac613fdd565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614a096020850193614a036149f663ffffffff875116426143cd565b8560808901511690614352565b9061526b565b80821015614a2257505b16825263ffffffff4216905290565b9050614a13565b80518015614a9957602003614a5b578051602082810191830183900312613c8e57519060ff8211614a5b575060ff1690565b612040906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613c0c565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161436557565b60ff16604d811161436557600a0a90565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614bea57828411614bc05790614b2991614abf565b91604d60ff8416118015614b87575b614b5157505090614b4b613f0992614ad3565b90614352565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614b9183614ad3565b801561439e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411614b38565b614bc991614abf565b91604d60ff841611614b5157505090614be4613f0992614ad3565b90614394565b5050505090565b90816020910312613c8e57518015158103613c8e5790565b73ffffffffffffffffffffffffffffffffffffffff600154163303614c2a57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115614e945767ffffffffffffffff81516020830120921691826000526008602052614c89816005604060002001615655565b15614e505760005260096020526040600020815167ffffffffffffffff8111613dc557614cb6825461473f565b601f8111614e1e575b506020601f8211600114614d585791614d32827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614d4895600091614d4d575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613c0c565b0390a2565b905084015138614d01565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110614e06575092614d489492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614dcf575b5050811b019055611688565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880614dc3565b9192602060018192868a015181550194019201614d88565b614e4a90836000526020600020601f840160051c810191602085106106bf57601f0160051c019061498d565b38614cbf565b50906120406040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613c0c565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15613c8e5767ffffffffffffffff906064604051809481937fa36a7fee0000000000000000000000000000000000000000000000000000000083526000978896879373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016600487015216602485015260448401525af18015611b9357614f90575050565b81614f9a91613e48565b50565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613f09604082613e48565b906127109167ffffffffffffffff614ff2602083016142e8565b166000908152600b602052604090209161ffff161561502557606061ffff615021935460901c16910135614352565b0490565b606061ffff615021935460801c16910135614352565b8151919291156151bd576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061515a576147cb91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836151bb604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040840151161580159061524c575b6151eb576147cb919261507e565b6064836151bb604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208401511615156151dd565b9190820180921161436557565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215613c8e57016020813591019167ffffffffffffffff8211613c8e578136038313613c8e57565b906040519182815491828252602082019060005260206000209260005b8181106152fa5750506147cb92500383613e48565b84548352600194850194879450602090930192016152e5565b80548210156147105760005260206000200190600090565b60008181526007602052604090205480156154ba577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161436557600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116143655781810361544b575b505050600654801561541c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016153d9816006615313565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6154a261545c61546d936006615313565b90549060031b1c9283926006615313565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260076020526040600020553880806153a0565b5050600090565b90600182019181600052826020526040600020548015156000146155ec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614365578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614365578181036155b5575b5050508054801561541c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906155768282615313565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6155d56155c561546d9386615313565b90549060031b1c92839286615313565b90556000528360205260406000205538808061553e565b50505050600090565b8060005260076020526040600020541560001461564f5760065468010000000000000000811015613dc55761563661546d8260018594016006556006615313565b9055600654906000526007602052604060002055600190565b50600090565b60008281526001820160205260409020546154ba5780549068010000000000000000821015613dc5578261569361546d846001809601855584615313565b905580549260005201602052604060002055600190565b9182549060ff8260a01c1615801561595e575b615958576fffffffffffffffffffffffffffffffff8216916001850190815461570263ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426143cd565b90816158ba575b505084811061586e57508383106157635750506157386fffffffffffffffffffffffffffffffff9283926143cd565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615802578161577b916143cd565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211614365576157c96157ce9273ffffffffffffffffffffffffffffffffffffffff9661526b565b614394565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82869293961161592e576158d592614a039160801c90614352565b808410156159295750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615709565b6158e0565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156156bd56fea164736f6c634300081a000a",
}

var LockReleaseTokenPoolABI = LockReleaseTokenPoolMetaData.ABI

var LockReleaseTokenPoolBin = LockReleaseTokenPoolMetaData.Bin

func DeployLockReleaseTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address, lockBox common.Address) (common.Address, *types.Transaction, *LockReleaseTokenPool, error) {
	parsed, err := LockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LockReleaseTokenPoolBin), backend, token, localTokenDecimals, advancedPoolHooks, rmnProxy, router, lockBox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockReleaseTokenPool{address: address, abi: *parsed, LockReleaseTokenPoolCaller: LockReleaseTokenPoolCaller{contract: contract}, LockReleaseTokenPoolTransactor: LockReleaseTokenPoolTransactor{contract: contract}, LockReleaseTokenPoolFilterer: LockReleaseTokenPoolFilterer{contract: contract}}, nil
}

type LockReleaseTokenPool struct {
	address common.Address
	abi     abi.ABI
	LockReleaseTokenPoolCaller
	LockReleaseTokenPoolTransactor
	LockReleaseTokenPoolFilterer
}

type LockReleaseTokenPoolCaller struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolSession struct {
	Contract     *LockReleaseTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LockReleaseTokenPoolCallerSession struct {
	Contract *LockReleaseTokenPoolCaller
	CallOpts bind.CallOpts
}

type LockReleaseTokenPoolTransactorSession struct {
	Contract     *LockReleaseTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type LockReleaseTokenPoolRaw struct {
	Contract *LockReleaseTokenPool
}

type LockReleaseTokenPoolCallerRaw struct {
	Contract *LockReleaseTokenPoolCaller
}

type LockReleaseTokenPoolTransactorRaw struct {
	Contract *LockReleaseTokenPoolTransactor
}

func NewLockReleaseTokenPool(address common.Address, backend bind.ContractBackend) (*LockReleaseTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(LockReleaseTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLockReleaseTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPool{address: address, abi: abi, LockReleaseTokenPoolCaller: LockReleaseTokenPoolCaller{contract: contract}, LockReleaseTokenPoolTransactor: LockReleaseTokenPoolTransactor{contract: contract}, LockReleaseTokenPoolFilterer: LockReleaseTokenPoolFilterer{contract: contract}}, nil
}

func NewLockReleaseTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*LockReleaseTokenPoolCaller, error) {
	contract, err := bindLockReleaseTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCaller{contract: contract}, nil
}

func NewLockReleaseTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LockReleaseTokenPoolTransactor, error) {
	contract, err := bindLockReleaseTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolTransactor{contract: contract}, nil
}

func NewLockReleaseTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LockReleaseTokenPoolFilterer, error) {
	contract, err := bindLockReleaseTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolFilterer{contract: contract}, nil
}

func bindLockReleaseTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolTransactor.contract.Transfer(opts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockReleaseTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.contract.Transfer(opts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getAdvancedPoolHooks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetAdvancedPoolHooks(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetAdvancedPoolHooks(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmations)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	return _LockReleaseTokenPool.Contract.GetCurrentRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector, customBlockConfirmations)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	return _LockReleaseTokenPool.Contract.GetCurrentRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector, customBlockConfirmations)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _LockReleaseTokenPool.Contract.GetDynamicConfig(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _LockReleaseTokenPool.Contract.GetDynamicConfig(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)

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

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _LockReleaseTokenPool.Contract.GetFee(&_LockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _LockReleaseTokenPool.Contract.GetFee(&_LockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetLockBox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getLockBox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetLockBox() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetLockBox(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetLockBox() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetLockBox(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getMinBlockConfirmations")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetMinBlockConfirmations() (uint16, error) {
	return _LockReleaseTokenPool.Contract.GetMinBlockConfirmations(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetMinBlockConfirmations() (uint16, error) {
	return _LockReleaseTokenPool.Contract.GetMinBlockConfirmations(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemotePools(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemotePools(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemoteToken(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemoteToken(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationsRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRequiredCCVs(&_LockReleaseTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationsRequested, extraData, direction)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRequiredCCVs(&_LockReleaseTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationsRequested, extraData, direction)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRmnProxy(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRmnProxy(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _LockReleaseTokenPool.Contract.GetSupportedChains(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _LockReleaseTokenPool.Contract.GetSupportedChains(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetToken() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetToken(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetToken(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _LockReleaseTokenPool.Contract.GetTokenDecimals(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _LockReleaseTokenPool.Contract.GetTokenDecimals(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _LockReleaseTokenPool.Contract.GetTokenTransferFeeConfig(&_LockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _LockReleaseTokenPool.Contract.GetTokenTransferFeeConfig(&_LockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsRemotePool(&_LockReleaseTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsRemotePool(&_LockReleaseTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedChain(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedChain(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedToken(&_LockReleaseTokenPool.CallOpts, token)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedToken(&_LockReleaseTokenPool.CallOpts, token)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) Owner() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.Owner(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) Owner() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.Owner(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.SupportsInterface(&_LockReleaseTokenPool.CallOpts, interfaceId)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.SupportsInterface(&_LockReleaseTokenPool.CallOpts, interfaceId)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) TypeAndVersion() (string, error) {
	return _LockReleaseTokenPool.Contract.TypeAndVersion(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _LockReleaseTokenPool.Contract.TypeAndVersion(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AcceptOwnership(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AcceptOwnership(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AddRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AddRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyChainUpdates(&_LockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyChainUpdates(&_LockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_LockReleaseTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_LockReleaseTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn0(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn0(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationsRequested)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint0(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint0(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RemoveRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RemoveRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetDynamicConfig(&_LockReleaseTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetDynamicConfig(&_LockReleaseTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setMinBlockConfirmations", minBlockConfirmations)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetMinBlockConfirmations(&_LockReleaseTokenPool.TransactOpts, minBlockConfirmations)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetMinBlockConfirmations(&_LockReleaseTokenPool.TransactOpts, minBlockConfirmations)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRateLimitConfig(&_LockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRateLimitConfig(&_LockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferOwnership(&_LockReleaseTokenPool.TransactOpts, to)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferOwnership(&_LockReleaseTokenPool.TransactOpts, to)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "updateAdvancedPoolHooks", newHook)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.UpdateAdvancedPoolHooks(&_LockReleaseTokenPool.TransactOpts, newHook)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.UpdateAdvancedPoolHooks(&_LockReleaseTokenPool.TransactOpts, newHook)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawFeeTokens(&_LockReleaseTokenPool.TransactOpts, feeTokens, recipient)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawFeeTokens(&_LockReleaseTokenPool.TransactOpts, feeTokens, recipient)
}

type LockReleaseTokenPoolAdvancedPoolHooksUpdatedIterator struct {
	Event *LockReleaseTokenPoolAdvancedPoolHooksUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAdvancedPoolHooksUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAdvancedPoolHooksUpdated)
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
		it.Event = new(LockReleaseTokenPoolAdvancedPoolHooksUpdated)
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

func (it *LockReleaseTokenPoolAdvancedPoolHooksUpdatedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAdvancedPoolHooksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAdvancedPoolHooksUpdated struct {
	OldHook common.Address
	NewHook common.Address
	Raw     types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolAdvancedPoolHooksUpdatedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAdvancedPoolHooksUpdatedIterator{contract: _LockReleaseTokenPool.contract, event: "AdvancedPoolHooksUpdated", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAdvancedPoolHooksUpdated)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseAdvancedPoolHooksUpdated(log types.Log) (*LockReleaseTokenPoolAdvancedPoolHooksUpdated, error) {
	event := new(LockReleaseTokenPoolAdvancedPoolHooksUpdated)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolChainAddedIterator struct {
	Event *LockReleaseTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolChainAdded)
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
		it.Event = new(LockReleaseTokenPoolChainAdded)
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

func (it *LockReleaseTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainAddedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolChainAddedIterator{contract: _LockReleaseTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolChainAdded)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseChainAdded(log types.Log) (*LockReleaseTokenPoolChainAdded, error) {
	event := new(LockReleaseTokenPoolChainAdded)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolChainRemovedIterator struct {
	Event *LockReleaseTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolChainRemoved)
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
		it.Event = new(LockReleaseTokenPoolChainRemoved)
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

func (it *LockReleaseTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolChainRemovedIterator{contract: _LockReleaseTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolChainRemoved)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolChainRemoved, error) {
	event := new(LockReleaseTokenPoolChainRemoved)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
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
		it.Event = new(LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
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

func (it *LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationsInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "CustomBlockConfirmationsInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationsInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsInboundRateLimitConsumed", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationsInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
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
		it.Event = new(LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
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

func (it *LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationsOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "CustomBlockConfirmationsOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationsOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsOutboundRateLimitConsumed", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationsOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolDynamicConfigSetIterator struct {
	Event *LockReleaseTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolDynamicConfigSet)
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
		it.Event = new(LockReleaseTokenPoolDynamicConfigSet)
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

func (it *LockReleaseTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolDynamicConfigSet struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAdmin       common.Address
	Raw            types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolDynamicConfigSetIterator{contract: _LockReleaseTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolDynamicConfigSet)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*LockReleaseTokenPoolDynamicConfigSet, error) {
	event := new(LockReleaseTokenPoolDynamicConfigSet)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolFeeTokenWithdrawnIterator struct {
	Event *LockReleaseTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(LockReleaseTokenPoolFeeTokenWithdrawn)
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

func (it *LockReleaseTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LockReleaseTokenPoolFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolFeeTokenWithdrawnIterator{contract: _LockReleaseTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolFeeTokenWithdrawn)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*LockReleaseTokenPoolFeeTokenWithdrawn, error) {
	event := new(LockReleaseTokenPoolFeeTokenWithdrawn)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolInboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(LockReleaseTokenPoolInboundRateLimitConsumed)
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

func (it *LockReleaseTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolInboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolInboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolInboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolInboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLockedOrBurnedIterator struct {
	Event *LockReleaseTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLockedOrBurned)
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
		it.Event = new(LockReleaseTokenPoolLockedOrBurned)
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

func (it *LockReleaseTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLockedOrBurnedIterator{contract: _LockReleaseTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLockedOrBurned)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*LockReleaseTokenPoolLockedOrBurned, error) {
	event := new(LockReleaseTokenPoolLockedOrBurned)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolMinBlockConfirmationsSetIterator struct {
	Event *LockReleaseTokenPoolMinBlockConfirmationsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolMinBlockConfirmationsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolMinBlockConfirmationsSet)
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
		it.Event = new(LockReleaseTokenPoolMinBlockConfirmationsSet)
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

func (it *LockReleaseTokenPoolMinBlockConfirmationsSetIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolMinBlockConfirmationsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolMinBlockConfirmationsSet struct {
	MinBlockConfirmations uint16
	Raw                   types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolMinBlockConfirmationsSetIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolMinBlockConfirmationsSetIterator{contract: _LockReleaseTokenPool.contract, event: "MinBlockConfirmationsSet", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolMinBlockConfirmationsSet) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolMinBlockConfirmationsSet)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseMinBlockConfirmationsSet(log types.Log) (*LockReleaseTokenPoolMinBlockConfirmationsSet, error) {
	event := new(LockReleaseTokenPoolMinBlockConfirmationsSet)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(LockReleaseTokenPoolOutboundRateLimitConsumed)
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

func (it *LockReleaseTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolOutboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolOutboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolOutboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolOutboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolOwnershipTransferRequestedIterator struct {
	Event *LockReleaseTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolOwnershipTransferRequested)
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
		it.Event = new(LockReleaseTokenPoolOwnershipTransferRequested)
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

func (it *LockReleaseTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolOwnershipTransferRequestedIterator{contract: _LockReleaseTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolOwnershipTransferRequested)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*LockReleaseTokenPoolOwnershipTransferRequested, error) {
	event := new(LockReleaseTokenPoolOwnershipTransferRequested)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolOwnershipTransferredIterator struct {
	Event *LockReleaseTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolOwnershipTransferred)
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
		it.Event = new(LockReleaseTokenPoolOwnershipTransferred)
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

func (it *LockReleaseTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolOwnershipTransferredIterator{contract: _LockReleaseTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolOwnershipTransferred)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LockReleaseTokenPoolOwnershipTransferred, error) {
	event := new(LockReleaseTokenPoolOwnershipTransferred)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRateLimitConfiguredIterator struct {
	Event *LockReleaseTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRateLimitConfigured)
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
		it.Event = new(LockReleaseTokenPoolRateLimitConfigured)
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

func (it *LockReleaseTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmations  bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRateLimitConfiguredIterator{contract: _LockReleaseTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRateLimitConfigured)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*LockReleaseTokenPoolRateLimitConfigured, error) {
	event := new(LockReleaseTokenPoolRateLimitConfigured)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolReleasedOrMintedIterator struct {
	Event *LockReleaseTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolReleasedOrMinted)
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
		it.Event = new(LockReleaseTokenPoolReleasedOrMinted)
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

func (it *LockReleaseTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolReleasedOrMintedIterator{contract: _LockReleaseTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolReleasedOrMinted)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*LockReleaseTokenPoolReleasedOrMinted, error) {
	event := new(LockReleaseTokenPoolReleasedOrMinted)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRemotePoolAddedIterator struct {
	Event *LockReleaseTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRemotePoolAdded)
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
		it.Event = new(LockReleaseTokenPoolRemotePoolAdded)
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

func (it *LockReleaseTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRemotePoolAddedIterator{contract: _LockReleaseTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRemotePoolAdded)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*LockReleaseTokenPoolRemotePoolAdded, error) {
	event := new(LockReleaseTokenPoolRemotePoolAdded)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRemotePoolRemovedIterator struct {
	Event *LockReleaseTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRemotePoolRemoved)
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
		it.Event = new(LockReleaseTokenPoolRemotePoolRemoved)
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

func (it *LockReleaseTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRemotePoolRemovedIterator{contract: _LockReleaseTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRemotePoolRemoved)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*LockReleaseTokenPoolRemotePoolRemoved, error) {
	event := new(LockReleaseTokenPoolRemotePoolRemoved)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *LockReleaseTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(LockReleaseTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _LockReleaseTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolTokenTransferFeeConfigDeleted)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*LockReleaseTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(LockReleaseTokenPoolTokenTransferFeeConfigDeleted)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *LockReleaseTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(LockReleaseTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _LockReleaseTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolTokenTransferFeeConfigUpdated)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*LockReleaseTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(LockReleaseTokenPoolTokenTransferFeeConfigUpdated)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetCurrentRateLimiterState struct {
	OutboundRateLimiterState RateLimiterTokenBucket
	InboundRateLimiterState  RateLimiterTokenBucket
}
type GetDynamicConfig struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAdmin       common.Address
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}

func (LockReleaseTokenPoolAdvancedPoolHooksUpdated) Topic() common.Hash {
	return common.HexToHash("0xbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209")
}

func (LockReleaseTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (LockReleaseTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x63335ad9e238acd0e9e6c1c20f529ffbea4cda73578c329a7aa7e9d61e5cdcc5")
}

func (LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x996b829383cc7e136842d4c4c175083bcf4e20807c7432105c1b794ba885e776")
}

func (LockReleaseTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (LockReleaseTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (LockReleaseTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (LockReleaseTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (LockReleaseTokenPoolMinBlockConfirmationsSet) Topic() common.Hash {
	return common.HexToHash("0x46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a01")
}

func (LockReleaseTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (LockReleaseTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (LockReleaseTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (LockReleaseTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (LockReleaseTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (LockReleaseTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (LockReleaseTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (LockReleaseTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (LockReleaseTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_LockReleaseTokenPool *LockReleaseTokenPool) Address() common.Address {
	return _LockReleaseTokenPool.address
}

type LockReleaseTokenPoolInterface interface {
	GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

		error)

	GetLockBox(opts *bind.CallOpts) (common.Address, error)

	GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolAdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*LockReleaseTokenPoolAdvancedPoolHooksUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*LockReleaseTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolChainRemoved, error)

	FilterCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationsInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationsOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*LockReleaseTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LockReleaseTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*LockReleaseTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*LockReleaseTokenPoolLockedOrBurned, error)

	FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolMinBlockConfirmationsSetIterator, error)

	WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolMinBlockConfirmationsSet) (event.Subscription, error)

	ParseMinBlockConfirmationsSet(log types.Log) (*LockReleaseTokenPoolMinBlockConfirmationsSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*LockReleaseTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LockReleaseTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*LockReleaseTokenPoolOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*LockReleaseTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*LockReleaseTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*LockReleaseTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*LockReleaseTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*LockReleaseTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*LockReleaseTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
