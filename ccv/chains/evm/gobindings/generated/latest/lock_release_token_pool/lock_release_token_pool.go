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
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var LockReleaseTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferLiquidity\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103775760c081616694803803809161002082856104f3565b8339810103126103775780516001600160a01b038116918282036103775761004a60208201610516565b61005660408301610524565b9061006360608401610524565b61007b60a061007460808701610524565b9501610524565b9433156104e257600180546001600160a01b03191633179055861580156104d1575b80156104c0575b6104485760805260c05260405163313ce56760e01b8152602081600481895afa60009181610484575b50610459575b5060a052600380546001600160a01b03199081166001600160a01b0393841617909155600280549091169282169290921790915516801561044857604051636eb1769f60e11b815230600482015260248101829052602081604481865afa90811561043c5760009161040a575b5061039f5760405191602083019263095ea7b360e01b84528260248201526000196044820152604481526101756064826104f3565b60008060409586519361018888866104f3565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610392573d906001600160401b03821161037c5785516101f99490926101ea601f8201601f1916602001856104f3565b83523d6000602085013e610538565b8051806102fb575b505060e0525161608b908161060982396080518181816104720152818161191f01528181611b2801528181611bc601528181611e4d015281816122c40152818161247b01528181612c3501528181612ed4015281816130cd015281816131a101528181613508015281816137610152818161393301528181613eca01528181613f24015281816140590152615451015260a051818181613d9301528181614f7b01528181614ffe01526154b0015260c051818181610eb3015281816119ba0152818161235e01528181612f6f01526137fc015260e0518181816104be0152818161316a015281816139ee01528181613ff501526153ec0152f35b8160209181010312610377576020015180159081150361037757610320573880610201565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916101f992606091610538565b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d602011610434575b81610425602093836104f3565b81010312610377575138610140565b3d9150610418565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b60ff1660ff821681810361046d57506100d3565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116104b8575b816104a0602093836104f3565b81010312610377576104b190610516565b90386100cd565b3d9150610493565b506001600160a01b038216156100a4565b506001600160a01b0385161561009d565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761037c57604052565b519060ff8216820361037757565b51906001600160a01b038216820361037757565b9192901561059a575081511561054c575090565b3b156105555790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156105ad5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b8381106105f05750508160006044809484010152601f80199101168101030190fd5b602082820181015160448784010152859350016105ce56fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461410a575080630a861f2a14613fa9578063181f5a7714613f4857806321df0da714613ef7578063240028e814613e965780632422ac4514613db757806324f65ee714613d795780632c06340414613ce057806337a3210d14613cac57806339077537146136c1578063432a6ba31461368d578063489a68f214612e2f5780634c5ef0ed14612de85780634e921c3014612d4957806362ddd3c414612cc25780636632008714612ae25780636cfd155314612a275780637437ff9f146129e657806379ba50971461291f5780638926f54f146128d957806389720a62146128125780638da5cb5b146127de5780639a4575b91461224a578063a42a7b8b146120e3578063acfecf9114611feb578063b1c71c6514611879578063b79465801461183c578063bfeffd3f14611790578063c4bffe2b14611665578063c7230a6014611462578063d8aa3f4014611328578063dc04fa1f14610ed7578063dc0bd97114610e86578063dcbd41bc14610c82578063e8a1da17146105c6578063eb521a4c146103fb578063f2fde38b1461032c578063fa41d79c146103075763ff8e03f3146101ce57600080fd5b34610304576040600319360112610304576101e7614385565b906101f06143cb565b6101f8615120565b73ffffffffffffffffffffffffffffffffffffffff83169283156102dc577f22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e616644797092937fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a556102d66040519283928390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b0390a180f35b6004837f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b5034610304578060031936011261030457602061ffff60035460a01c16604051908152f35b50346103045760206003193601126103045773ffffffffffffffffffffffffffffffffffffffff61035b614385565b610363615120565b163381146103d357807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346103045760206003193601126103045760043573ffffffffffffffffffffffffffffffffffffffff600c5416330361059a576040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064808201839052815282907f0000000000000000000000000000000000000000000000000000000000000000906104a7906104a16084826142ab565b82615b01565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610596576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9290921660048301526024820184905282908290604490829084905af1801561058b57610572575b5050337fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb3120888380a380f35b8161057c916142ab565b610587578138610548565b5080fd5b6040513d84823e3d90fd5b8280fd5b6024827f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b50346103045760406003193601126103045760043567ffffffffffffffff8111610587576105f8903690600401614582565b9060243567ffffffffffffffff8111610c7e579061061b84923690600401614582565b939091610626615120565b83905b828210610ae35750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610adf578060051b83013585811215610adb57830161012081360312610adb576040519461068d8661428f565b6106968261443d565b8652602082013567ffffffffffffffff81116105875782019436601f87011215610587578535956106c687614941565b966106d460405198896142ab565b80885260208089019160051b83010190368211610adb5760208301905b828210610aa8575050505060208701958652604083013567ffffffffffffffff81116105965761072490369085016144f5565b916040880192835261074e61073c3660608701614dbc565b9460608a0195865260c0369101614dbc565b956080890196875283515115610a805761077267ffffffffffffffff8a5116615c41565b15610a495767ffffffffffffffff8951168252600860205260408220610799865182615547565b6107a7885160028301615547565b6004855191019080519067ffffffffffffffff8211610a1c576107ca8354614bce565b601f81116109e1575b50602090601f8311600114610942576108219291869183610937575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b8851805182101561085b579061085560019261084e8367ffffffffffffffff8f511692614b8b565b519061516b565b01610826565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561092967ffffffffffffffff60019796949851169251935191516108f56108c060405196879687526101006020880152610100870190614326565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101939290919361065c565b015190508e806107ef565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106109c95750908460019594939210610992575b505050811b019055610824565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080610985565b9293602060018192878601518155019501930161096f565b610a0c9084875260208720601f850160051c81019160208610610a12575b601f0160051c0190614e58565b8d6107d3565b90915081906109ff565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff8111610ad757602091610acc83928336918901016144f5565b8152019101906106f1565b8680fd5b8480fd5b8380f35b9267ffffffffffffffff610b05610b008486889a9699979a614d41565b6148ef565b1691610b1083615837565b15610c52578284526008602052610b2c600560408620016157d4565b94845b8651811015610b65576001908587526008602052610b5e60056040892001610b57838b614b8b565b51906159cd565b5001610b2f565b5093969290945094909480875260086020526005604088208881558860018201558860028201558860038201558860048201610ba18154614bce565b80610c11575b5050500180549088815581610bf3575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101909194939294610629565b885260208820908101905b81811015610bb757888155600101610bfe565b601f8111600114610c275750555b888a80610ba7565b81835260208320610c4291601f01861c810190600101614e58565b8082528160208120915555610c1f565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346103045760206003193601126103045760043567ffffffffffffffff811161058757610cb49036906004016145b3565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580610e64575b610e3857825b818110610ce7578380f35b610cf2818385614d51565b67ffffffffffffffff610d04826148ef565b1690610d1d826000526007602052604060002054151590565b15610e0c57907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610dcc610da6602060019897018b610d5e82614d61565b15610dd3578790526004602052610d8560408d20610d7f3660408801614dbc565b90615547565b868c526005602052610da160408d20610d7f3660a08801614dbc565b614d61565b916040519215158352610dbf6020840160408301614e14565b60a0608084019101614e14565ba201610cdc565b60026040828a610da194526008602052610df5828220610d7f36858c01614dbc565b8a815260086020522001610d7f3660a08801614dbc565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610cd6565b5034610304578060031936011261030457602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346103045760406003193601126103045760043567ffffffffffffffff811161058757610f099036906004016145b3565b60243567ffffffffffffffff8111610c7e57610f29903690600401614582565b919092610f34615120565b845b828110610fa057505050825b818110610f4d578380f35b8067ffffffffffffffff610f67610b006001948688614d41565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610f42565b610fae610b00828585614d51565b610fb9828585614d51565b90602082019060e0830190610fcd82614d61565b156112f35760a0840161271061ffff610fe583614d6e565b1610156112e45760c085019161271061ffff61100085614d6e565b1610156112ac5763ffffffff61101586614d7d565b16156112775767ffffffffffffffff1694858c52600b60205260408c2061103b86614d7d565b63ffffffff1690805490604084019161105383614d7d565b60201b67ffffffff000000001693606086019461106f86614d7d565b60401b6bffffffff000000000000000016966080019661108e88614d7d565b60601b6fffffffff00000000000000000000000016916110ad8a614d6e565b60801b71ffff0000000000000000000000000000000016936110ce8c614d6e565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717815561118187614d61565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055604051966111d290614d8e565b63ffffffff1687526111e390614d8e565b63ffffffff1660208701526111f790614d8e565b63ffffffff16604086015261120b90614d8e565b63ffffffff16606085015261121f90614481565b61ffff16608084015261123190614481565b61ffff1660a083015261124390614452565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610f36565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff6112bb86614d6e565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff6112bb602493614d6e565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b503461030457608060031936011261030457611342614385565b5061134b614426565b611353614470565b5060643567ffffffffffffffff8111610596579167ffffffffffffffff60409261138360e0953690600401614490565b50508260c0855161139381614273565b82815282602082015282878201528260608201528260808201528260a08201520152168152600b60205220604051906113cb82614273565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346103045760406003193601126103045760043567ffffffffffffffff811161058757611494903690600401614582565b9061149d6143cb565b6114a5615120565b835b8381106114b2578480f35b80602073ffffffffffffffffffffffffffffffffffffffff6114df6114da6024958989614d41565b6148ce565b16604051938480927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa801561165a578690611624575b6001925080611530575b50016114a7565b6115b773ffffffffffffffffffffffffffffffffffffffff6115566114da858a8a614d41565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff881660248201526044808201869052815291166115b26064836142ab565b615b01565b73ffffffffffffffffffffffffffffffffffffffff6115da6114da848989614d41565b60405192835216907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602073ffffffffffffffffffffffffffffffffffffffff871692a338611529565b509060203d8111611653575b61163a81836142ab565b602082600092810103126103045750906001915161151f565b503d611630565b6040513d88823e3d90fd5b5034610304578060031936011261030457604051906006548083528260208101600684526020842092845b8181106117775750506116a5925003836142ab565b81516116c96116b382614941565b916116c160405193846142ab565b808352614941565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015611728578067ffffffffffffffff61171560019388614b8b565b51166117218286614b8b565b52016116f6565b50925090604051928392602084019060208552518091526040840192915b818110611754575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611746565b8454835260019485019487945060209093019201611690565b50346103045760206003193601126103045760043573ffffffffffffffffffffffffffffffffffffffff8116809103610587576117cb615120565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a1161760035580f35b50346103045760206003193601126103045761187561186161185c61440f565b614d1f565b604051918291602083526020830190614326565b0390f35b50346103045760606003193601126103045760043567ffffffffffffffff8111610587578060040160a06003198336030112610596576118b761445f565b926044359367ffffffffffffffff8511610587576118dc6118f9953690600401614490565b95906118e6614b72565b506118f183866154e4565b9636916144be565b6084850190611907826148ce565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611fa157602486019577ffffffffffffffff0000000000000000000000000000000061196d886148ef565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561165a578691611f72575b50611f4a5767ffffffffffffffff611a01886148ef565b16611a19816000526007602052604060002054151590565b15611f1f57602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561165a578690611ece575b73ffffffffffffffffffffffffffffffffffffffff9150163303611ea25760648101359561ffff611aab8a89614ce3565b9516948515611df15761ffff60035460a01c168015611dc957808710611d9957507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611aff8b6148ef565b1691828952600460205280611b5060408b2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615cf6565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff600354169283611c82575b611c7889611c4761185c611bad8e8d614ce3565b92611bb7846153d5565b611bc0816148ef565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a26148ef565b90611c506154a9565b60405192611c5d84614257565b83526020830152604051928392604084526040840190614558565b9060208301520390f35b833b15610ad757869493929185918a604051988997889687957f5c3af7ca000000000000000000000000000000000000000000000000000000008752600487016060905280611cd091615784565b6064880160a09052610104880190611ce79261497a565b93611cf19061443d565b67ffffffffffffffff166084870152604401611d0c906143ee565b73ffffffffffffffffffffffffffffffffffffffff1660a48601528c60c4860152611d36906143ee565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526024840152828103600319016044840152611d6b91614326565b03925af1801561058b57611d84575b8080808080611b99565b611d8f8280926142ab565b6103045780611d7a565b87604491887f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b6004887f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611e248b6148ef565b1691828952600860205280611e7560408b2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615cf6565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611b79565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611f17575b81611ee8602093836142ab565b81010312611f1357611f0e73ffffffffffffffffffffffffffffffffffffffff91614959565b611a7a565b8580fd5b3d9150611edb565b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611f94915060203d602011611f9a575b611f8c81836142ab565b810190615108565b386119ea565b503d611f82565b60248473ffffffffffffffffffffffffffffffffffffffff611fc2856148ce565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346103045767ffffffffffffffff61200336614513565b92909161200e615120565b1691612027836000526007602052604060002054151590565b15610c52578284526008602052612056600560408620016120493684866144be565b60208151910120906159cd565b1561209b57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769161209560405192839260208452602084019161497a565b0390a280f35b826120df836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161497a565b0390fd5b50346103045760206003193601126103045767ffffffffffffffff61210661440f565b168152600860205261211d600560408320016157d4565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061216261214c83614941565b9261215a60405194856142ab565b808452614941565b01835b818110612239575050825b82518110156121b6578061218660019285614b8b565b518552600960205261219a60408620614c21565b6121a48285614b8b565b526121af8184614b8b565b5001612170565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106121ee57505050500390f35b91936020612229827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614326565b96019201920185949391926121df565b806060602080938601015201612165565b5034610304576020600319360112610304576004359067ffffffffffffffff821161030457816004019160a060031982360301126105875761228a614b72565b506020926040519261229c85856142ab565b808452608483016122ac816148ce565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036127bd57602484019477ffffffffffffffff00000000000000000000000000000000612312876148ef565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152878160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156127425784916127a0575b506127785767ffffffffffffffff6123a5876148ef565b166123bd816000526007602052604060002054151590565b1561274d578773ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156127425784906126fa575b73ffffffffffffffffffffffffffffffffffffffff91501633036126ce57606485013594859467ffffffffffffffff612455896148ef565b1680865260088a526124a36040872073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016998a91615cf6565b6040805173ffffffffffffffffffffffffffffffffffffffff8a168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a273ffffffffffffffffffffffffffffffffffffffff6003541692836125b7575b8961258761185c8b8b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108c61253f816153d5565b67ffffffffffffffff612551856148ef565b6040805173ffffffffffffffffffffffffffffffffffffffff909616865233602087015285019290925216918060608101611c3f565b906125906154a9565b6040519261259d84614257565b835281830152611875604051928284938452830190614558565b833b15611f135791858094928a9694604051978896879586947f5c3af7ca00000000000000000000000000000000000000000000000000000000865260048601606090528061260591615784565b6064870160a0905261010487019061261c9261497a565b926126269061443d565b67ffffffffffffffff166084860152604401612641906143ee565b73ffffffffffffffffffffffffffffffffffffffff1660a48501528b60c485015261266b906143ee565b73ffffffffffffffffffffffffffffffffffffffff1660e48401528360248401528281036003190160448401526126a191614326565b03925af1801561058b576126b9575b8080808061250b565b6126c48280926142ab565b61030457806126b0565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d831161273b575b61271081836142ab565b81010312610c7e5761273673ffffffffffffffffffffffffffffffffffffffff91614959565b61241d565b503d612706565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6127b79150883d8a11611f9a57611f8c81836142ab565b3861238e565b9073ffffffffffffffffffffffffffffffffffffffff611fc26024936148ce565b5034610304578060031936011261030457602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346103045760c06003193601126103045761282c614385565b612834614426565b9060643561ffff81168103610c7e5760843567ffffffffffffffff8111610adb57612863903690600401614490565b9160a435936002851015610ad75761287e95604435916149b9565b90604051918291602083016020845282518091526020604085019301915b8181106128aa575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff1684528594506020938401939092019160010161289c565b503461030457602060031936011261030457602061291567ffffffffffffffff61290161440f565b166000526007602052604060002054151590565b6040519015158152f35b5034610304578060031936011261030457805473ffffffffffffffffffffffffffffffffffffffff811633036129be577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b5034610304578060031936011261030457600254600a546040805173ffffffffffffffffffffffffffffffffffffffff938416815292909116602083015290f35b5034610304576020600319360112610304577f64187bd7b97e66658c91904f3021d7c28de967281d18b1a20742348afdd6a6b373ffffffffffffffffffffffffffffffffffffffff612a77614385565b612a7f615120565b6102d6600c54918381167fffffffffffffffffffffffff0000000000000000000000000000000000000000841617600c55604051938493168390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b503461030457604060031936011261030457612afc614385565b60243590612b08615120565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214612bd4575b73ffffffffffffffffffffffffffffffffffffffff1690813b1561059657826040517f0a861f2a000000000000000000000000000000000000000000000000000000008152826004820152818160248183885af1801561058b57612bbf575b505060207f6fa7abcf1345d1d478e5ea0da6b5f26a90eadb0546ef15ed3833944fbfd1db6291604051908152a280f35b81612bc9916142ab565b610596578238612b8f565b90506040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015612cb7578390612c6c575b919050612b30565b506020813d602011612caf575b81612c86602093836142ab565b81010312612caa5773ffffffffffffffffffffffffffffffffffffffff9051612c64565b600080fd5b3d9150612c79565b6040513d85823e3d90fd5b503461030457612cd136614513565b612cdd93929193615120565b67ffffffffffffffff8216612cff816000526007602052604060002054151590565b15612d1e5750612d1b9293612d159136916144be565b9061516b565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346103045760206003193601126103045760043561ffff811690818103610596577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb91602091612d98615120565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006003549260a01b16911617600355604051908152a180f35b503461030457604060031936011261030457612e0261440f565b906024359067ffffffffffffffff821161030457602061291584612e2936600487016144f5565b90614904565b50346103045760406003193601126103045760043567ffffffffffffffff8111610587578060040191610100600319833603011261030457612e6f61445f565b9181604051612e7d8161420c565b5260c48101926064820135612ead612ea7612ea2612e9b888a61487d565b36916144be565b614f07565b82614ffb565b946084840191612ebc836148ce565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361366c57602485019777ffffffffffffffff00000000000000000000000000000000612f228a6148ef565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156135ef57889161364d575b506136255767ffffffffffffffff612fb68a6148ef565b16612fce816000526007602052604060002054151590565b156135fa57602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156135ef5788916135d0575b50156135a457613045896148ef565b9461305b60a4880196612e29612e9b898661487d565b1561355d5761ffff1690878a8a84156134aa575067ffffffffffffffff9150613083906148ef565b1680895260056020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a806130f560408d2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615cf6565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff6003541693846132d6575b5050505050505060440192613152846148ce565b9173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001693813b15610596576040517f69328dec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80871660048301526024820188905290911660448201529082908290818381606481015b03925af1801561058b576132c1575b5050608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff61328f6132897ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966148ef565b926148ce565b60405196875233898801521660408601528560608601521692a2604051906132b68261420c565b815260405190518152f35b6132cc8280926142ab565b6103045780613237565b843b156134a6578895949392869289928d6040519a8b998a9889977f5eff3bf700000000000000000000000000000000000000000000000000000000895260048901606090526133268780615784565b60648b0161010090526101648b019061333e9261497a565b946133489061443d565b67ffffffffffffffff1660848a0152604401613363906143ee565b73ffffffffffffffffffffffffffffffffffffffff1660a489015260c488015261338c906143ee565b73ffffffffffffffffffffffffffffffffffffffff1660e48701526133b19084615784565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c878403016101048801526133e6929161497a565b906133f19083615784565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c86840301610124870152613426929161497a565b9060e48b0161343491615784565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610144860152613469929161497a565b908b6024840152604483015203925af18015612cb757908391613491575b808080808061313e565b8161349b916142ab565b610587578138613487565b8880fd5b806135306002604067ffffffffffffffff6134e57f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c976148ef565b16968781526008602052200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615cf6565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a261311e565b613567868361487d565b6120df6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161497a565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b6135e9915060203d602011611f9a57611f8c81836142ab565b38613036565b6040513d8a823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b613666915060203d602011611f9a57611f8c81836142ab565b38612f9f565b60248673ffffffffffffffffffffffffffffffffffffffff611fc2866148ce565b5034610304578060031936011261030457602073ffffffffffffffffffffffffffffffffffffffff600c5416604051908152f35b5034610304576020600319360112610304576004359067ffffffffffffffff8211610304578160040191610100600319823603011261058757816040516137078161420c565b52816040516137158161420c565b5260648101359060c481019061373a613734612ea2612e9b858961487d565b84614ffb565b926084820192613749846148ce565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603613c8b57602483019377ffffffffffffffff000000000000000000000000000000006137af866148ef565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156135ef578891613c6c575b506136255767ffffffffffffffff613843866148ef565b1661385b816000526007602052604060002054151590565b156135fa57602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156135ef578891613c4d575b50156135a4576138d2856148ef565b926138e860a4860194612e29612e9b878d61487d565b15613c43579087988894939288995067ffffffffffffffff613909896148ef565b16808752600860205261395b6002604089200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169b8c91615cf6565b6040805173ffffffffffffffffffffffffffffffffffffffff8c168152602081018d90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a273ffffffffffffffffffffffffffffffffffffffff600354169182613a79575b50505050505050604401936139d7856148ce565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15610596576040517f69328dec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8087166004830152602482018890529091166044820152908290829081838160648101613228565b823b15610ad7578787959186928b604051998a98899788967f5eff3bf70000000000000000000000000000000000000000000000000000000088526004880160609052613ac68780615784565b60648a0161010090526101648a0190613ade9261497a565b94613ae89061443d565b67ffffffffffffffff166084890152604401613b03906143ee565b73ffffffffffffffffffffffffffffffffffffffff1660a488015260c4870152613b2c906143ee565b73ffffffffffffffffffffffffffffffffffffffff1660e4860152613b519084615784565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c86840301610104870152613b86929161497a565b90613b919083615784565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610124860152613bc6929161497a565b9060e48b01613bd491615784565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610144850152613c09929161497a565b8c602483015282604483015203925af1801561058b57613c2e575b80808080806139c3565b81613c38916142ab565b610adb578438613c24565b613567848a61487d565b613c66915060203d602011611f9a57611f8c81836142ab565b386138c3565b613c85915060203d602011611f9a57611f8c81836142ab565b3861382c565b60248673ffffffffffffffffffffffffffffffffffffffff611fc2876148ce565b5034610304578060031936011261030457602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b50346103045760c060031936011261030457613cfa614385565b50613d03614426565b613d0b6143a8565b506084359161ffff831683036103045760a4359067ffffffffffffffff82116103045760a063ffffffff8061ffff613d528888613d4b3660048b01614490565b50506146f8565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b5034610304578060031936011261030457602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461030457604060031936011261030457613dd161440f565b60243591821515830361030457610140613e94613dee8585614675565b613e4460409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b5034610304576020600319360112610304576020613eb2614385565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b5034610304578060031936011261030457602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461030457806003193601126103045750611875604051613f6b6040826142ab565b601e81527f4c6f636b52656c65617365546f6b656e506f6f6c20312e372e302d64657600006020820152604051918291602083526020830190614326565b5034610304576020600319360112610304576004359073ffffffffffffffffffffffffffffffffffffffff600c541633036140de5773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b15610587576040517f69328dec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166004820152602481018490523360448201529082908290606490829084905af1801561058b576140ce575b5090337fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf98401717198380a380f35b816140d8916142ab565b386140a4565b807f8e4a23d6000000000000000000000000000000000000000000000000000000006024925233600452fd5b905034610587576020600319360112610587576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361059657602092507faff2afbf0000000000000000000000000000000000000000000000000000000081149081156141e2575b81156141b8575b811561418e575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438614187565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614180565b7f331710310000000000000000000000000000000000000000000000000000000081149150614179565b6020810190811067ffffffffffffffff82111761422857604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761422857604052565b60e0810190811067ffffffffffffffff82111761422857604052565b60a0810190811067ffffffffffffffff82111761422857604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761422857604052565b67ffffffffffffffff811161422857601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106143705750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201614331565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203612caa57565b6064359073ffffffffffffffffffffffffffffffffffffffff82168203612caa57565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203612caa57565b359073ffffffffffffffffffffffffffffffffffffffff82168203612caa57565b6004359067ffffffffffffffff82168203612caa57565b6024359067ffffffffffffffff82168203612caa57565b359067ffffffffffffffff82168203612caa57565b35908115158203612caa57565b6024359061ffff82168203612caa57565b6044359061ffff82168203612caa57565b359061ffff82168203612caa57565b9181601f84011215612caa5782359167ffffffffffffffff8311612caa5760208381860195010111612caa57565b9291926144ca826142ec565b916144d860405193846142ab565b829481845281830111612caa578281602093846000960137010152565b9080601f83011215612caa57816020614510933591016144be565b90565b906040600319830112612caa5760043567ffffffffffffffff81168103612caa57916024359067ffffffffffffffff8211612caa5761455491600401614490565b9091565b6145109160206145718351604084526040840190614326565b920151906020818403910152614326565b9181601f84011215612caa5782359167ffffffffffffffff8311612caa576020808501948460051b010111612caa57565b9181601f84011215612caa5782359167ffffffffffffffff8311612caa576020808501948460081b010111612caa57565b604051906145f18261428f565b60006080838281528260208201528260408201528260608201520152565b9060405161461c8161428f565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916146876145e4565b506146906145e4565b506146c4571660005260086020526040600020906145106146b860026146bd6146b88661460f565b614e82565b940161460f565b16908160005260046020526146df6146b8604060002061460f565b9160005260056020526145106146b8604060002061460f565b9061ffff8060035460a01c1691169283151592838094614875575b61484b5767ffffffffffffffff16600052600b6020526040600020916040519261473c84614273565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a0152614830576147d1575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b81939750809294501061480057505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b508215614713565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215612caa570180359067ffffffffffffffff8211612caa57602001918136038313612caa57565b3573ffffffffffffffffffffffffffffffffffffffff81168103612caa5790565b3567ffffffffffffffff81168103612caa5790565b9067ffffffffffffffff61451092166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116142285760051b60200190565b519073ffffffffffffffffffffffffffffffffffffffff82168203612caa57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff60035416958615614b5057614a549467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c486019161497a565b916002821015614b21578380600094819460a483015203915afa908115614b1557600091614a80575090565b3d8083833e614a8f81836142ab565b8101906020818303126105965780519067ffffffffffffffff8211610c7e570181601f8201121561059657805190614ac682614941565b93614ad460405195866142ab565b82855260208086019360051b8301019384116103045750602001905b828210614afd5750505090565b60208091614b0a84614959565b815201910190614af0565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b5050505050505050604051614b666020826142ab565b60008152600036813790565b60405190614b7f82614257565b60606020838281520152565b8051821015614b9f5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015614c17575b6020831014614be857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614bdd565b9060405191826000825492614c3584614bce565b8084529360018116908115614ca35750600114614c5c575b50614c5a925003836142ab565b565b90506000929192526020600020906000915b818310614c87575050906020614c5a9282010138614c4d565b6020919350806001915483858901015201910190918492614c6e565b60209350614c5a9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614c4d565b91908203918211614cf057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff1660005260086020526145106004604060002001614c21565b9190811015614b9f5760051b0190565b9190811015614b9f5760081b0190565b358015158103612caa5790565b3561ffff81168103612caa5790565b3563ffffffff81168103612caa5790565b359063ffffffff82168203612caa57565b35906fffffffffffffffffffffffffffffffff82168203612caa57565b9190826060910312612caa576040516060810181811067ffffffffffffffff821117614228576040526040614e0f818395614df681614452565b8552614e0460208201614d9f565b602086015201614d9f565b910152565b6fffffffffffffffffffffffffffffffff614e5260408093614e3581614452565b1515865283614e4660208301614d9f565b16602087015201614d9f565b16910152565b818110614e63575050565b60008155600101614e58565b81810292918115918404141715614cf057565b614e8a6145e4565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614ee76020850193614ee1614ed463ffffffff87511642614ce3565b8560808901511690614e6f565b90615777565b80821015614f0057505b16825263ffffffff4216905290565b9050614ef1565b80518015614f7757602003614f39578051602082810191830183900312612caa57519060ff8211614f39575060ff1690565b6120df906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614326565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211614cf057565b60ff16604d8111614cf057600a0a90565b8115614fcc570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414615101578284116150d7579061504091614f9d565b91604d60ff841611801561509e575b6150685750509061506261451092614fb1565b90614e6f565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506150a883614fb1565b8015614fcc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04841161504f565b6150e091614f9d565b91604d60ff841611615068575050906150fb61451092614fb1565b90614fc2565b5050505090565b90816020910312612caa57518015158103612caa5790565b73ffffffffffffffffffffffffffffffffffffffff60015416330361514157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156153ab5767ffffffffffffffff815160208301209216918260005260086020526151a0816005604060002001615ca1565b156153675760005260096020526040600020815167ffffffffffffffff8111614228576151cd8254614bce565b601f8111615335575b506020601f821160011461526f5791615249827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361525f95600091615264575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614326565b0390a2565b905084015138615218565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b81811061531d57509261525f9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106152e6575b5050811b019055611861565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538806152da565b9192602060018192868a01518155019401920161529f565b61536190836000526020600020601f840160051c81019160208510610a1257601f0160051c0190614e58565b386151d6565b50906120df6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614326565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15612caa576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166004820152602481019190915260009182908290604490829084905af1801561058b5761549c575050565b816154a6916142ab565b50565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526145106040826142ab565b906127109167ffffffffffffffff6154fe602083016148ef565b166000908152600b602052604090209161ffff161561553157606061ffff61552d935460901c16910135614e6f565b0490565b606061ffff61552d935460801c16910135614e6f565b8151919291156156c9576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061566657614c5a91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836156c7604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408401511615801590615758575b6156f757614c5a919261558a565b6064836156c7604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208401511615156156e9565b91908201809211614cf057565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215612caa57016020813591019167ffffffffffffffff8211612caa578136038313612caa57565b906040519182815491828252602082019060005260206000209260005b818110615806575050614c5a925003836142ab565b84548352600194850194879450602090930192016157f1565b8054821015614b9f5760005260206000200190600090565b60008181526007602052604090205480156159c6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614cf057600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614cf057818103615957575b5050506006548015615928577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016158e581600661581f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6159ae61596861597993600661581f565b90549060031b1c928392600661581f565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260076020526040600020553880806158ac565b5050600090565b9060018201918160005282602052604060002054801515600014615af8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614cf0578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614cf057818103615ac1575b50505080548015615928577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615a82828261581f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b615ae1615ad1615979938661581f565b90549060031b1c9283928661581f565b905560005283602052604060002055388080615a4a565b50505050600090565b73ffffffffffffffffffffffffffffffffffffffff615b90911691604092600080855193615b2f87866142ab565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d15615c39573d91615b74836142ec565b92615b81875194856142ab565b83523d6000602085013e615fb2565b80519081615b9d57505050565b602080615bae938301019101615108565b15615bb65750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615fb2565b80600052600760205260406000205415600014615c9b576006546801000000000000000081101561422857615c82615979826001859401600655600661581f565b9055600654906000526007602052604060002055600190565b50600090565b60008281526001820160205260409020546159c657805490680100000000000000008210156142285782615cdf61597984600180960185558461581f565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615faa575b615fa4576fffffffffffffffffffffffffffffffff82169160018501908154615d4e63ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614ce3565b9081615f06575b5050848110615eba5750838310615daf575050615d846fffffffffffffffffffffffffffffffff928392614ce3565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615e4e5781615dc791614ce3565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211614cf057615e15615e1a9273ffffffffffffffffffffffffffffffffffffffff96615777565b614fc2565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615f7a57615f2192614ee19160801c90614e6f565b80841015615f755750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615d55565b615f2c565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615d09565b9192901561602d5750815115615fc6575090565b3b15615fcf5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156160405750805190602001fd5b6120df906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061432656fea164736f6c634300081a000a",
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _LockReleaseTokenPool.Contract.GetCurrentRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _LockReleaseTokenPool.Contract.GetCurrentRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _LockReleaseTokenPool.Contract.GetFee(&_LockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _LockReleaseTokenPool.Contract.GetFee(&_LockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _LockReleaseTokenPool.Contract.GetMinBlockConfirmation(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _LockReleaseTokenPool.Contract.GetMinBlockConfirmation(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRebalancer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRebalancer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRebalancer() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRebalancer(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRebalancer() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRebalancer(&_LockReleaseTokenPool.CallOpts)
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRequiredCCVs(&_LockReleaseTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRequiredCCVs(&_LockReleaseTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn0(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn0(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "provideLiquidity", amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ProvideLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ProvideLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint0(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint0(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetDynamicConfig(&_LockReleaseTokenPool.TransactOpts, router, rateLimitAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetDynamicConfig(&_LockReleaseTokenPool.TransactOpts, router, rateLimitAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setMinBlockConfirmation", minBlockConfirmation)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetMinBlockConfirmation(&_LockReleaseTokenPool.TransactOpts, minBlockConfirmation)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetMinBlockConfirmation(&_LockReleaseTokenPool.TransactOpts, minBlockConfirmation)
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetRebalancer(opts *bind.TransactOpts, rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setRebalancer", rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetRebalancer(rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRebalancer(&_LockReleaseTokenPool.TransactOpts, rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetRebalancer(rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRebalancer(&_LockReleaseTokenPool.TransactOpts, rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) TransferLiquidity(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "transferLiquidity", from, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) TransferLiquidity(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferLiquidity(&_LockReleaseTokenPool.TransactOpts, from, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) TransferLiquidity(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferLiquidity(&_LockReleaseTokenPool.TransactOpts, from, amount)
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
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

type LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*LockReleaseTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolFeeTokenWithdrawnIterator{contract: _LockReleaseTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
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

type LockReleaseTokenPoolLiquidityAddedIterator struct {
	Event *LockReleaseTokenPoolLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityAdded)
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
		it.Event = new(LockReleaseTokenPoolLiquidityAdded)
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

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityAdded struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityAddedIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityAdded)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityAdded(log types.Log) (*LockReleaseTokenPoolLiquidityAdded, error) {
	event := new(LockReleaseTokenPoolLiquidityAdded)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLiquidityRemovedIterator struct {
	Event *LockReleaseTokenPoolLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityRemoved)
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
		it.Event = new(LockReleaseTokenPoolLiquidityRemoved)
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

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityRemoved struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityRemovedIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityRemoved)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityRemoved(log types.Log) (*LockReleaseTokenPoolLiquidityRemoved, error) {
	event := new(LockReleaseTokenPoolLiquidityRemoved)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLiquidityTransferredIterator struct {
	Event *LockReleaseTokenPoolLiquidityTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityTransferred)
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
		it.Event = new(LockReleaseTokenPoolLiquidityTransferred)
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

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityTransferred struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityTransferred(opts *bind.FilterOpts, from []common.Address) (*LockReleaseTokenPoolLiquidityTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityTransferred", fromRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityTransferredIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityTransferred", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityTransferred, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityTransferred", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityTransferred)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityTransferred", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityTransferred(log types.Log) (*LockReleaseTokenPoolLiquidityTransferred, error) {
	event := new(LockReleaseTokenPoolLiquidityTransferred)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityTransferred", log); err != nil {
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

type LockReleaseTokenPoolMinBlockConfirmationSetIterator struct {
	Event *LockReleaseTokenPoolMinBlockConfirmationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolMinBlockConfirmationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolMinBlockConfirmationSet)
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
		it.Event = new(LockReleaseTokenPoolMinBlockConfirmationSet)
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

func (it *LockReleaseTokenPoolMinBlockConfirmationSetIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolMinBlockConfirmationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolMinBlockConfirmationSet struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolMinBlockConfirmationSetIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolMinBlockConfirmationSetIterator{contract: _LockReleaseTokenPool.contract, event: "MinBlockConfirmationSet", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolMinBlockConfirmationSet) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolMinBlockConfirmationSet)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseMinBlockConfirmationSet(log types.Log) (*LockReleaseTokenPoolMinBlockConfirmationSet, error) {
	event := new(LockReleaseTokenPoolMinBlockConfirmationSet)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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
	CustomBlockConfirmation   bool
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

type LockReleaseTokenPoolRebalancerSetIterator struct {
	Event *LockReleaseTokenPoolRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRebalancerSet)
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
		it.Event = new(LockReleaseTokenPoolRebalancerSet)
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

func (it *LockReleaseTokenPoolRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRebalancerSet struct {
	OldRebalancer common.Address
	NewRebalancer common.Address
	Raw           types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRebalancerSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolRebalancerSetIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RebalancerSet")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRebalancerSetIterator{contract: _LockReleaseTokenPool.contract, event: "RebalancerSet", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRebalancerSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRebalancerSet) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RebalancerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRebalancerSet)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RebalancerSet", log); err != nil {
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

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRebalancerSet(log types.Log) (*LockReleaseTokenPoolRebalancerSet, error) {
	event := new(LockReleaseTokenPoolRebalancerSet)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RebalancerSet", log); err != nil {
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

func (LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (LockReleaseTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e6166447970")
}

func (LockReleaseTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (LockReleaseTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (LockReleaseTokenPoolLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088")
}

func (LockReleaseTokenPoolLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719")
}

func (LockReleaseTokenPoolLiquidityTransferred) Topic() common.Hash {
	return common.HexToHash("0x6fa7abcf1345d1d478e5ea0da6b5f26a90eadb0546ef15ed3833944fbfd1db62")
}

func (LockReleaseTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (LockReleaseTokenPoolMinBlockConfirmationSet) Topic() common.Hash {
	return common.HexToHash("0xa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb")
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

func (LockReleaseTokenPoolRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x64187bd7b97e66658c91904f3021d7c28de967281d18b1a20742348afdd6a6b3")
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

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

	GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error)

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

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, rebalancer common.Address) (*types.Transaction, error)

	TransferLiquidity(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolAdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*LockReleaseTokenPoolAdvancedPoolHooksUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*LockReleaseTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolChainRemoved, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*LockReleaseTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*LockReleaseTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*LockReleaseTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolInboundRateLimitConsumed, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*LockReleaseTokenPoolLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*LockReleaseTokenPoolLiquidityRemoved, error)

	FilterLiquidityTransferred(opts *bind.FilterOpts, from []common.Address) (*LockReleaseTokenPoolLiquidityTransferredIterator, error)

	WatchLiquidityTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityTransferred, from []common.Address) (event.Subscription, error)

	ParseLiquidityTransferred(log types.Log) (*LockReleaseTokenPoolLiquidityTransferred, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*LockReleaseTokenPoolLockedOrBurned, error)

	FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolMinBlockConfirmationSetIterator, error)

	WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolMinBlockConfirmationSet) (event.Subscription, error)

	ParseMinBlockConfirmationSet(log types.Log) (*LockReleaseTokenPoolMinBlockConfirmationSet, error)

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

	FilterRebalancerSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolRebalancerSetIterator, error)

	WatchRebalancerSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRebalancerSet) (event.Subscription, error)

	ParseRebalancerSet(log types.Log) (*LockReleaseTokenPoolRebalancerSet, error)

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
