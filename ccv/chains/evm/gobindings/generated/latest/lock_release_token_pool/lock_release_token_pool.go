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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferLiquidity\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103a75760c081616811803803809161002082856105e2565b8339810103126103a75761003381610605565b9061004060208201610619565b9061004d60408201610605565b61005960608301610605565b9261007260a061006b60808601610605565b9401610605565b9333156105d157600180546001600160a01b031916331790556001600160a01b03861695861580156105c0575b80156105af575b6105375760805260c05260405163313ce56760e01b8152602081600481895afa60009181610573575b50610548575b5060a052600380546001600160a01b03199081166001600160a01b03938416179091556002805490911692821692909217909155168015610537576040516321df0da760e01b8152602081600481855afa801561046c5783916000916104fb575b506001600160a01b03160361047857604051636eb1769f60e11b815230600482015260248101829052602081604481865afa90811561046c5760009161043a575b506103cf5760405191602083019263095ea7b360e01b84528260248201526000196044820152604481526101ac6064826105e2565b6000806040958651936101bf88866105e2565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d156103c2573d906001600160401b0382116103ac578551610230949092610221601f8201601f1916602001856105e2565b83523d6000602085013e610627565b80518061032b575b505060e0525161611990816106f8823960805181818161046f0152818161190e01528181611b1d01528181611bc201528181611f1c015281816123a60152818161255d01528181612de901528181613095015281816132980152818161333e015281816136e80152818161394601528181613b160152818161402a0152818161408401526141c2015260a051818181613ef3015281816155cf01528181615652015261577b015260c051818181610edc015281816119a9015281816124400152818161313001526139e1015260e0518181816104e501528181611c6901528181612698015281816133f301526141550152f35b81602091810103126103a757602001518015908115036103a757610350573880610238565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b9161023092606091610627565b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d602011610464575b81610455602093836105e2565b810103126103a7575138610177565b3d9150610448565b6040513d6000823e3d90fd5b6020600491604051928380926321df0da760e01b82525afa90811561046c576000916104c1575b5063961c9a4f60e01b60009081526001600160a01b0391909116600452602490fd5b90506020813d6020116104f3575b816104dc602093836105e2565b810103126103a7576104ed90610605565b8161049f565b3d91506104cf565b9150506020813d60201161052f575b81610517602093836105e2565b810103126103a7576105298391610605565b38610136565b3d915061050a565b630a64406560e11b60005260046000fd5b60ff1660ff821681810361055c57506100d5565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116105a7575b8161058f602093836105e2565b810103126103a7576105a090610619565b90386100cf565b3d9150610582565b506001600160a01b038216156100a6565b506001600160a01b0385161561009f565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b038211908210176103ac57604052565b51906001600160a01b03821682036103a757565b519060ff821682036103a757565b91929015610689575081511561063b575090565b3b156106445790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561069c5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b8381106106df5750508160006044809484010152601f80199101168101030190fd5b602082820181015160448784010152859350016106bd56fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461426c575080630a861f2a14614109578063181f5a77146140a857806321df0da714614057578063240028e814613ff65780632422ac4514613f1757806324f65ee714613ed95780632c06340414613e4057806337a3210d14613e0c57806339077537146138a1578063432a6ba31461386d578063489a68f214612fe35780634c5ef0ed14612f9c5780634e921c3014612efd57806362ddd3c414612e765780636632008714612c965780636cfd155314612bdb5780637437ff9f14612b9a57806379ba509714612ad35780638926f54f14612a8d57806389720a62146129c65780638da5cb5b146129925780639a4575b914612324578063a42a7b8b146121bd578063acfecf91146120c5578063b1c71c651461185e578063b794658014611821578063bfeffd3f14611775578063c4bffe2b1461164a578063c7230a601461148b578063d8aa3f4014611351578063dc04fa1f14610f00578063dc0bd97114610eaf578063dcbd41bc14610cab578063e8a1da17146105ef578063eb521a4c146103fb578063f2fde38b1461032c578063fa41d79c146103075763ff8e03f3146101ce57600080fd5b34610304576040600319360112610304576101e76144e7565b906101f061452d565b6101f8615069565b73ffffffffffffffffffffffffffffffffffffffff83169283156102dc577f22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e616644797092937fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a556102d66040519283928390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b0390a180f35b6004837f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b5034610304578060031936011261030457602061ffff60035460a01c16604051908152f35b50346103045760206003193601126103045773ffffffffffffffffffffffffffffffffffffffff61035b6144e7565b610363615069565b163381146103d357807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346103045760206003193601126103045760043573ffffffffffffffffffffffffffffffffffffffff600c541633036105c3576040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064810182905282907f0000000000000000000000000000000000000000000000000000000000000000906104ce906104c881608481015b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261440d565b82615bf4565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b156105bf5773ffffffffffffffffffffffffffffffffffffffff606484928360405195869485937fa36a7fee0000000000000000000000000000000000000000000000000000000085521660048401528160248401528860448401525af180156105b45761059b575b5050337fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb3120888380a380f35b816105a59161440d565b6105b0578138610571565b5080fd5b6040513d84823e3d90fd5b8280fd5b6024827f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b50346103045760406003193601126103045760043567ffffffffffffffff81116105b0576106219036906004016146e4565b9060243567ffffffffffffffff8111610ca75790610644849236906004016146e4565b93909161064f615069565b83905b828210610b0c5750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610b08578060051b83013585811215610b0457830161012081360312610b0457604051946106b6866143f1565b6106bf8261459f565b8652602082013567ffffffffffffffff81116105b05782019436601f870112156105b0578535956106ef87614a52565b966106fd604051988961440d565b80885260208089019160051b83010190368211610b045760208301905b828210610ad1575050505060208701958652604083013567ffffffffffffffff81116105bf5761074d9036908501614657565b91604088019283526107776107653660608701614e91565b9460608a0195865260c0369101614e91565b956080890196875283515115610aa95761079b67ffffffffffffffff8a5116615b3f565b15610a725767ffffffffffffffff89511682526008602052604082206107c286518261531e565b6107d088516002830161531e565b6004855191019080519067ffffffffffffffff8211610a45576107f38354614cdf565b601f8111610a0a575b50602090601f831160011461096b5761084a9291869183610960575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b88518051821015610884579061087e6001926108778367ffffffffffffffff8f511692614c9c565b51906150b4565b0161084f565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561095267ffffffffffffffff600197969498511692519351915161091e6108e960405196879687526101006020880152610100870190614488565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610685565b015190508e80610818565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106109f257509084600195949392106109bb575b505050811b01905561084d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806109ae565b92936020600181928786015181550195019301610998565b610a359084875260208720601f850160051c81019160208610610a3b575b601f0160051c0190614f2d565b8d6107fc565b9091508190610a28565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff8111610b0057602091610af58392833691890101614657565b81520191019061071a565b8680fd5b8480fd5b8380f35b9267ffffffffffffffff610b2e610b298486889a9699979a614e64565b6149df565b1691610b3983615812565b15610c7b578284526008602052610b55600560408620016157af565b94845b8651811015610b8e576001908587526008602052610b8760056040892001610b80838b614c9c565b51906159a8565b5001610b58565b5093969290945094909480875260086020526005604088208881558860018201558860028201558860038201558860048201610bca8154614cdf565b80610c3a575b5050500180549088815581610c1c575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101909194939294610652565b885260208820908101905b81811015610be057888155600101610c27565b601f8111600114610c505750555b888a80610bd0565b81835260208320610c6b91601f01861c810190600101614f2d565b8082528160208120915555610c48565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346103045760206003193601126103045760043567ffffffffffffffff81116105b057610cdd903690600401614715565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580610e8d575b610e6157825b818110610d10578380f35b610d1b818385614e16565b67ffffffffffffffff610d2d826149df565b1690610d46826000526007602052604060002054151590565b15610e3557907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610df5610dcf602060019897018b610d8782614e26565b15610dfc578790526004602052610dae60408d20610da83660408801614e91565b9061531e565b868c526005602052610dca60408d20610da83660a08801614e91565b614e26565b916040519215158352610de86020840160408301614ee9565b60a0608084019101614ee9565ba201610d05565b60026040828a610dca94526008602052610e1e828220610da836858c01614e91565b8a815260086020522001610da83660a08801614e91565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610cff565b5034610304578060031936011261030457602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346103045760406003193601126103045760043567ffffffffffffffff81116105b057610f32903690600401614715565b60243567ffffffffffffffff8111610ca757610f529036906004016146e4565b919092610f5d615069565b845b828110610fc957505050825b818110610f76578380f35b8067ffffffffffffffff610f90610b296001948688614e64565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610f6b565b610fd7610b29828585614e16565b610fe2828585614e16565b90602082019060e0830190610ff682614e26565b1561131c5760a0840161271061ffff61100e83614e33565b16101561130d5760c085019161271061ffff61102985614e33565b1610156112d55763ffffffff61103e86614e42565b16156112a05767ffffffffffffffff1694858c52600b60205260408c2061106486614e42565b63ffffffff1690805490604084019161107c83614e42565b60201b67ffffffff000000001693606086019461109886614e42565b60401b6bffffffff00000000000000001696608001966110b788614e42565b60601b6fffffffff00000000000000000000000016916110d68a614e33565b60801b71ffff0000000000000000000000000000000016936110f78c614e33565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1617171781556111aa87614e26565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055604051966111fb90614e53565b63ffffffff16875261120c90614e53565b63ffffffff16602087015261122090614e53565b63ffffffff16604086015261123490614e53565b63ffffffff166060850152611248906145e3565b61ffff16608084015261125a906145e3565b61ffff1660a083015261126c906145b4565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610f5f565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff6112e486614e33565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff6112e4602493614e33565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346103045760806003193601126103045761136b6144e7565b50611374614588565b61137c6145d2565b5060643567ffffffffffffffff81116105bf579167ffffffffffffffff6040926113ac60e09536906004016145f2565b50508260c085516113bc816143d5565b82815282602082015282878201528260608201528260808201528260a08201520152168152600b60205220604051906113f4826143d5565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346103045760406003193601126103045760043567ffffffffffffffff81116105b0576114bd9036906004016146e4565b906114c661452d565b916114cf615069565b835b8181106114dc578480f35b73ffffffffffffffffffffffffffffffffffffffff6115046114ff838587614e64565b6149f4565b1690604051917f70a08231000000000000000000000000000000000000000000000000000000008352306004840152602083602481845afa801561163f578790611607575b600193508061155b575b5050016114d1565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff88166024820152604481018290526115be906115b8816064810161049c565b83615bf4565b6040519081527f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602073ffffffffffffffffffffffffffffffffffffffff891692a33880611553565b50909160203d8111611638575b61161e818361440d565b602082600092810103126103045750906001929151611549565b503d611614565b6040513d89823e3d90fd5b5034610304578060031936011261030457604051906006548083528260208101600684526020842092845b81811061175c57505061168a9250038361440d565b81516116ae61169882614a52565b916116a6604051938461440d565b808352614a52565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b845181101561170d578067ffffffffffffffff6116fa60019388614c9c565b51166117068286614c9c565b52016116db565b50925090604051928392602084019060208552518091526040840192915b818110611739575050500390f35b825167ffffffffffffffff1684528594506020938401939092019160010161172b565b8454835260019485019487945060209093019201611675565b50346103045760206003193601126103045760043573ffffffffffffffffffffffffffffffffffffffff81168091036105b0576117b0615069565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a1161760035580f35b50346103045760206003193601126103045761185a611846611841614571565b614df4565b604051918291602083526020830190614488565b0390f35b5034610304576060600319360112610304576004359067ffffffffffffffff8211610304578160040160a060031984360301126105b05761189d6145c1565b9260443567ffffffffffffffff8111610ca7576118c16118e79136906004016145f2565b91906118cb614c83565b506118d4614c83565b506118df8786615adc565b923691614620565b9160848101936118f6856149f4565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361207b57602482019477ffffffffffffffff0000000000000000000000000000000061195c876149df565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611fee57889161204c575b506120245767ffffffffffffffff6119f0876149df565b16611a08816000526007602052604060002054151590565b15611ff957602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611fee578890611f9d575b73ffffffffffffffffffffffffffffffffffffffff9150163303611f715760648301359461ffff611a9a8688614f86565b9916938415611eba5761ffff60035460a01c168015611e9257808610611e62575088997f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611af48b9c999a9b6149df565b1691828952600460205280611b4560408b2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615d84565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff600354169182611d4e575b50505050505050611b9e91614f86565b611ba7826149df565b6040805173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680825233602083015291810184905290939167ffffffffffffffff16907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a283611c32611841836149df565b93611c3b615774565b60405195611c48876143b9565b86526020860152611c8f73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016936149df565b92803b156105bf576040517fa36a7fee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909216600483015267ffffffffffffffff9093166024820152604481018490529182908183816064810103925af18015611d4357611d2e575b611d2483836040519283926040845260408401906146ba565b9060208301520390f35b611d3984809261440d565b6105bf5782611d0b565b6040513d86823e3d90fd5b823b15610b0057869485918b604051988997889687957f5c3af7ca000000000000000000000000000000000000000000000000000000008752600487016060905280611d9991615d34565b6064880160a09052610104880190611db092614a8b565b93611dba9061459f565b67ffffffffffffffff166084870152604401611dd590614550565b73ffffffffffffffffffffffffffffffffffffffff1660a48601528d60c4860152611dff90614550565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526024840152828103600319016044840152611e3491614488565b03925af180156105b457611e4d575b8080808080611b8e565b81611e579161440d565b610ca7578338611e43565b89604491877f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b60048a7f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b88997fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611ef38b9c999a9b6149df565b1691828952600860205280611f4460408b2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615d84565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611b6e565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611fe6575b81611fb76020938361440d565b81010312611fe257611fdd73ffffffffffffffffffffffffffffffffffffffff91614a6a565b611a69565b8780fd5b3d9150611faa565b6040513d8a823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b61206e915060203d602011612074575b612066818361440d565b81019061575c565b386119d9565b503d61205c565b60248673ffffffffffffffffffffffffffffffffffffffff61209c886149f4565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346103045767ffffffffffffffff6120dd36614675565b9290916120e8615069565b1691612101836000526007602052604060002054151590565b15610c7b57828452600860205261213060056040862001612123368486614620565b60208151910120906159a8565b1561217557907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769161216f604051928392602084526020840191614a8b565b0390a280f35b826121b9836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614a8b565b0390fd5b50346103045760206003193601126103045767ffffffffffffffff6121e0614571565b16815260086020526121f7600560408320016157af565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061223c61222683614a52565b92612234604051948561440d565b808452614a52565b01835b818110612313575050825b8251811015612290578061226060019285614c9c565b518552600960205261227460408620614d32565b61227e8285614c9c565b526122898184614c9c565b500161224a565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106122c857505050500390f35b91936020612303827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614488565b96019201920185949391926122b9565b80606060208093860101520161223f565b50346103045760206003193601126103045760043567ffffffffffffffff81116105b057806004019160a0600319833603011261030457612363614c83565b5061236c614c83565b5060405160209261237d848361440d565b828252608481019161238e836149f4565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361297157602482019177ffffffffffffffff000000000000000000000000000000006123f4846149df565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152868160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156128f6578691612954575b5061292c5767ffffffffffffffff612487846149df565b1661249f816000526007602052604060002054151590565b15612901578673ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156128f65786906128aa575b73ffffffffffffffffffffffffffffffffffffffff915016330361287e57606481013596879467ffffffffffffffff612537866149df565b16808852600889526125856040892073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169b8c91615d84565b6040805173ffffffffffffffffffffffffffffffffffffffff8c168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a273ffffffffffffffffffffffffffffffffffffffff600354169182612767575b505050505067ffffffffffffffff612604826149df565b6040805173ffffffffffffffffffffffffffffffffffffffff891681523360208201529081018590529116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2612662611841826149df565b9461266b615774565b60405196612678886143b9565b8752858701526126be73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016926149df565b90823b15610b04576040517fa36a7fee00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015267ffffffffffffffff90911660248201526044810192909252829082908183816064810103925af180156105b457612752575b505061185a6040519282849384528301906146ba565b61275d82809261440d565b610304578061273c565b823b15611fe257908780949392604051968795869485937f5c3af7ca0000000000000000000000000000000000000000000000000000000085526004850160609052806127b391615d34565b6064860160a090526101048601906127ca92614a8b565b916127d48c61459f565b67ffffffffffffffff1660848601526044016127ef90614550565b73ffffffffffffffffffffffffffffffffffffffff1660a48501528b60c485015261281990614550565b73ffffffffffffffffffffffffffffffffffffffff1660e484015283602484015282810360031901604484015261284f91614488565b03925af18015611d4357908491612869575b8080806125ed565b816128739161440d565b6105bf578238612861565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508681813d83116128ef575b6128c0818361440d565b810103126128eb576128e673ffffffffffffffffffffffffffffffffffffffff91614a6a565b6124ff565b8580fd5b503d6128b6565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b61296b9150873d891161207457612066818361440d565b38612470565b60248473ffffffffffffffffffffffffffffffffffffffff61209c866149f4565b5034610304578060031936011261030457602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346103045760c0600319360112610304576129e06144e7565b6129e8614588565b9060643561ffff81168103610ca75760843567ffffffffffffffff8111610b0457612a179036906004016145f2565b9160a435936002851015610b0057612a329560443591614aca565b90604051918291602083016020845282518091526020604085019301915b818110612a5e575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101612a50565b5034610304576020600319360112610304576020612ac967ffffffffffffffff612ab5614571565b166000526007602052604060002054151590565b6040519015158152f35b5034610304578060031936011261030457805473ffffffffffffffffffffffffffffffffffffffff81163303612b72577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b5034610304578060031936011261030457600254600a546040805173ffffffffffffffffffffffffffffffffffffffff938416815292909116602083015290f35b5034610304576020600319360112610304577f64187bd7b97e66658c91904f3021d7c28de967281d18b1a20742348afdd6a6b373ffffffffffffffffffffffffffffffffffffffff612c2b6144e7565b612c33615069565b6102d6600c54918381167fffffffffffffffffffffffff0000000000000000000000000000000000000000841617600c55604051938493168390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b503461030457604060031936011261030457612cb06144e7565b60243590612cbc615069565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214612d88575b73ffffffffffffffffffffffffffffffffffffffff1690813b156105bf57826040517f0a861f2a000000000000000000000000000000000000000000000000000000008152826004820152818160248183885af180156105b457612d73575b505060207f6fa7abcf1345d1d478e5ea0da6b5f26a90eadb0546ef15ed3833944fbfd1db6291604051908152a280f35b81612d7d9161440d565b6105bf578238612d43565b90506040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa8015612e6b578390612e20575b919050612ce4565b506020813d602011612e63575b81612e3a6020938361440d565b81010312612e5e5773ffffffffffffffffffffffffffffffffffffffff9051612e18565b600080fd5b3d9150612e2d565b6040513d85823e3d90fd5b503461030457612e8536614675565b612e9193929193615069565b67ffffffffffffffff8216612eb3816000526007602052604060002054151590565b15612ed25750612ecf9293612ec9913691614620565b906150b4565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346103045760206003193601126103045760043561ffff8116908181036105bf577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb91602091612f4c615069565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006003549260a01b16911617600355604051908152a180f35b503461030457604060031936011261030457612fb6614571565b906024359067ffffffffffffffff8211610304576020612ac984612fdd3660048701614657565b90614a15565b5034610304576040600319360112610304576004359067ffffffffffffffff8211610304578160040161010060031984360301126105b0576130236145c1565b92826040516130318161436e565b528260405161303f8161436e565b5260c4810190606481013561306f61306961306461305d8688615018565b3691614620565b61555b565b8261564f565b936084830161307d816149f4565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361384c57602484019477ffffffffffffffff000000000000000000000000000000006130e3876149df565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156137cf57899161382d575b506138055767ffffffffffffffff613177876149df565b1661318f816000526007602052604060002054151590565b156137da57602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156137cf5789916137b0575b501561378457613206866149df565b9361321c60a4870195612fdd61305d8888615018565b1561373d5761ffff899a99979899169387898b87151560001461368a575067ffffffffffffffff915061324e906149df565b1680895260056020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158b806132c060408d2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615d84565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff6003541692836134b9575b5050505050505060440161331c816149f4565b50613326836149df565b9373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016947ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0608067ffffffffffffffff613394866149f4565b9373ffffffffffffffffffffffffffffffffffffffff604051958b87523360208801521660408601528560608601521692a2604051946133d38661436e565b81865261341f61341973ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016966149df565b936149f4565b853b15610b045767ffffffffffffffff9573ffffffffffffffffffffffffffffffffffffffff8692608495604051998a9889977f74fd18ac000000000000000000000000000000000000000000000000000000008952600489015216602487015260448601521660648401525af18015612e6b576134a4575b60208260405190518152f35b6134af83809261440d565b6105b05781613498565b833b15613686578895869289928c6040519a8b998a9889977f5eff3bf700000000000000000000000000000000000000000000000000000000895260048901606090526135068780615d34565b60648b0161010090526101648b019061351e92614a8b565b946135289061459f565b67ffffffffffffffff1660848a015260440161354390614550565b73ffffffffffffffffffffffffffffffffffffffff1660a489015260c488015261356c90614550565b73ffffffffffffffffffffffffffffffffffffffff1660e48701526135919084615d34565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c878403016101048801526135c69291614a8b565b906135d19083615d34565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c868403016101248701526136069291614a8b565b9060e48b0161361491615d34565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101448601526136499291614a8b565b908c6024840152604483015203925af1908115612e6b578391613671575b8080808080613309565b8161367b9161440d565b6105b0578138613667565b8880fd5b806137106002604067ffffffffffffffff6136c57f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c976149df565b16968781526008602052200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615d84565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26132e9565b6137478585615018565b6121b96040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614a8b565b6024887f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b6137c9915060203d60201161207457612066818361440d565b386131f7565b6040513d8b823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008952600452602488fd5b6004887f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b613846915060203d60201161207457612066818361440d565b38613160565b8673ffffffffffffffffffffffffffffffffffffffff61209c6024936149f4565b5034610304578060031936011261030457602073ffffffffffffffffffffffffffffffffffffffff600c5416604051908152f35b50346103045760206003193601126103045760043567ffffffffffffffff81116105b057806004019061010060031982360301126105bf57826040516138e68161436e565b52826040516138f48161436e565b52826040516139028161436e565b5260c4810190606481013561392061306961306461305d8688615018565b936084830161392e816149f4565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361384c57602484019477ffffffffffffffff00000000000000000000000000000000613994876149df565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156137cf578991613ded575b506138055767ffffffffffffffff613a28876149df565b16613a40816000526007602052604060002054151590565b156137da57602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156137cf578991613dce575b501561378457613ab7866149df565b91613acd60a4870193612fdd61305d8688615018565b15613dc457918895939188959367ffffffffffffffff613aec8a6149df565b168089526008602052613b3e600260408b200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169c8d91615d84565b6040805173ffffffffffffffffffffffffffffffffffffffff8d168152602081018a90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a273ffffffffffffffffffffffffffffffffffffffff600354169182613bf8575b50505050505060440190613bb9826149f4565b50613bc3846149df565b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0608067ffffffffffffffff613394866149f4565b823b1561368657899589958692604051988997889687957f5eff3bf700000000000000000000000000000000000000000000000000000000875260048701606090528d613c458780615d34565b60648a0161010090526101648a0190613c5d92614a8b565b94613c679061459f565b67ffffffffffffffff166084890152604401613c8290614550565b73ffffffffffffffffffffffffffffffffffffffff1660a488015260c4870152613cab90614550565b73ffffffffffffffffffffffffffffffffffffffff1660e4860152613cd09084615d34565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c86840301610104870152613d059291614a8b565b90613d109083615d34565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610124860152613d459291614a8b565b9060e48a01613d5391615d34565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610144850152613d889291614a8b565b89602483015282604483015203925af1908115611d43578491613daf575b80808080613ba6565b81613db99161440d565b6105bf578238613da6565b6137478385615018565b613de7915060203d60201161207457612066818361440d565b38613aa8565b613e06915060203d60201161207457612066818361440d565b38613a11565b5034610304578060031936011261030457602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b50346103045760c060031936011261030457613e5a6144e7565b50613e63614588565b613e6b61450a565b506084359161ffff831683036103045760a4359067ffffffffffffffff82116103045760a063ffffffff8061ffff613eb28888613eab3660048b016145f2565b505061485a565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b5034610304578060031936011261030457602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461030457604060031936011261030457613f31614571565b60243591821515830361030457610140613ff4613f4e85856147d7565b613fa460409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346103045760206003193601126103045760206140126144e7565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b5034610304578060031936011261030457602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346103045780600319360112610304575061185a6040516140cb60408261440d565b601e81527f4c6f636b52656c65617365546f6b656e506f6f6c20312e372e302d64657600006020820152604051918291602083526020830190614488565b5034610304576020600319360112610304576004359073ffffffffffffffffffffffffffffffffffffffff600c541633036142405773ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803b156105b0578180916084604051809481937f74fd18ac00000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660048401528160248401528860448401523360648401525af180156105b457614230575b5090337fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf98401717198380a380f35b8161423a9161440d565b38614206565b807f8e4a23d6000000000000000000000000000000000000000000000000000000006024925233600452fd5b9050346105b05760206003193601126105b0576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036105bf57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115614344575b811561431a575b81156142f0575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386142e9565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506142e2565b7f3317103100000000000000000000000000000000000000000000000000000000811491506142db565b6020810190811067ffffffffffffffff82111761438a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761438a57604052565b60e0810190811067ffffffffffffffff82111761438a57604052565b60a0810190811067ffffffffffffffff82111761438a57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761438a57604052565b67ffffffffffffffff811161438a57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b8481106144d25750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201614493565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203612e5e57565b6064359073ffffffffffffffffffffffffffffffffffffffff82168203612e5e57565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203612e5e57565b359073ffffffffffffffffffffffffffffffffffffffff82168203612e5e57565b6004359067ffffffffffffffff82168203612e5e57565b6024359067ffffffffffffffff82168203612e5e57565b359067ffffffffffffffff82168203612e5e57565b35908115158203612e5e57565b6024359061ffff82168203612e5e57565b6044359061ffff82168203612e5e57565b359061ffff82168203612e5e57565b9181601f84011215612e5e5782359167ffffffffffffffff8311612e5e5760208381860195010111612e5e57565b92919261462c8261444e565b9161463a604051938461440d565b829481845281830111612e5e578281602093846000960137010152565b9080601f83011215612e5e5781602061467293359101614620565b90565b906040600319830112612e5e5760043567ffffffffffffffff81168103612e5e57916024359067ffffffffffffffff8211612e5e576146b6916004016145f2565b9091565b6146729160206146d38351604084526040840190614488565b920151906020818403910152614488565b9181601f84011215612e5e5782359167ffffffffffffffff8311612e5e576020808501948460051b010111612e5e57565b9181601f84011215612e5e5782359167ffffffffffffffff8311612e5e576020808501948460081b010111612e5e57565b60405190614753826143f1565b60006080838281528260208201528260408201528260608201520152565b9060405161477e816143f1565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916147e9614746565b506147f2614746565b506148265716600052600860205260406000209061467261481a600261481f61481a86614771565b614f93565b9401614771565b169081600052600460205261484161481a6040600020614771565b91600052600560205261467261481a6040600020614771565b9061ffff8060035460a01c16911692831515928380946149d7575b6149ad5767ffffffffffffffff16600052600b6020526040600020916040519261489e846143d5565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a015261499257614933575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b81939750809294501061496257505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b508215614875565b3567ffffffffffffffff81168103612e5e5790565b3573ffffffffffffffffffffffffffffffffffffffff81168103612e5e5790565b9067ffffffffffffffff61467292166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff811161438a5760051b60200190565b519073ffffffffffffffffffffffffffffffffffffffff82168203612e5e57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff60035416958615614c6157614b659467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c4860191614a8b565b916002821015614c32578380600094819460a483015203915afa908115614c2657600091614b91575090565b3d8083833e614ba0818361440d565b8101906020818303126105bf5780519067ffffffffffffffff8211610ca7570181601f820112156105bf57805190614bd782614a52565b93614be5604051958661440d565b82855260208086019360051b8301019384116103045750602001905b828210614c0e5750505090565b60208091614c1b84614a6a565b815201910190614c01565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b5050505050505050604051614c7760208261440d565b60008152600036813790565b60405190614c90826143b9565b60606020838281520152565b8051821015614cb05760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015614d28575b6020831014614cf957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614cee565b9060405191826000825492614d4684614cdf565b8084529360018116908115614db45750600114614d6d575b50614d6b9250038361440d565b565b90506000929192526020600020906000915b818310614d98575050906020614d6b9282010138614d5e565b6020919350806001915483858901015201910190918492614d7f565b60209350614d6b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614d5e565b67ffffffffffffffff1660005260086020526146726004604060002001614d32565b9190811015614cb05760081b0190565b358015158103612e5e5790565b3561ffff81168103612e5e5790565b3563ffffffff81168103612e5e5790565b359063ffffffff82168203612e5e57565b9190811015614cb05760051b0190565b35906fffffffffffffffffffffffffffffffff82168203612e5e57565b9190826060910312612e5e576040516060810181811067ffffffffffffffff82111761438a576040526040614ee4818395614ecb816145b4565b8552614ed960208201614e74565b602086015201614e74565b910152565b6fffffffffffffffffffffffffffffffff614f2760408093614f0a816145b4565b1515865283614f1b60208301614e74565b16602087015201614e74565b16910152565b818110614f38575050565b60008155600101614f2d565b81810292918115918404141715614f5757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211614f5757565b614f9b614746565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614ff86020850193614ff2614fe563ffffffff87511642614f86565b8560808901511690614f44565b9061554e565b8082101561501157505b16825263ffffffff4216905290565b9050615002565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215612e5e570180359067ffffffffffffffff8211612e5e57602001918136038313612e5e57565b73ffffffffffffffffffffffffffffffffffffffff60015416330361508a57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156152f45767ffffffffffffffff815160208301209216918260005260086020526150e9816005604060002001615b9f565b156152b05760005260096020526040600020815167ffffffffffffffff811161438a576151168254614cdf565b601f811161527e575b506020601f82116001146151b85791615192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936151a8956000916151ad575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614488565b0390a2565b905084015138615161565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106152665750926151a89492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061522f575b5050811b019055611846565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880615223565b9192602060018192868a0151815501940192016151e8565b6152aa90836000526020600020601f840160051c81019160208510610a3b57601f0160051c0190614f2d565b3861511f565b50906121b96040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614488565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b8151919291156154a0576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061543d57614d6b91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b60648361549e604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040840151161580159061552f575b6154ce57614d6b9192615361565b60648361549e604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208401511615156154c0565b91908201809211614f5757565b805180156155cb5760200361558d578051602082810191830183900312612e5e57519060ff821161558d575060ff1690565b6121b9906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614488565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211614f5757565b60ff16604d8111614f5757600a0a90565b8115615620570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146157555782841161572b5790615694916155f1565b91604d60ff84161180156156f2575b6156bc575050906156b661467292615605565b90614f44565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506156fc83615605565b8015615620577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0484116156a3565b615734916155f1565b91604d60ff8416116156bc5750509061574f61467292615605565b90615616565b5050505090565b90816020910312612e5e57518015158103612e5e5790565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001660208201526020815261467260408261440d565b906040519182815491828252602082019060005260206000209260005b8181106157e1575050614d6b9250038361440d565b84548352600194850194879450602090930192016157cc565b8054821015614cb05760005260206000200190600090565b60008181526007602052604090205480156159a1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614f5757600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614f5757818103615932575b5050506006548015615903577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016158c08160066157fa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6159896159436159549360066157fa565b90549060031b1c92839260066157fa565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526007602052604060002055388080615887565b5050600090565b9060018201918160005282602052604060002054801515600014615ad3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614f57578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614f5757818103615a9c575b50505080548015615903577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190615a5d82826157fa565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b615abc615aac61595493866157fa565b90549060031b1c928392866157fa565b905560005283602052604060002055388080615a25565b50505050600090565b906127109167ffffffffffffffff615af6602083016149df565b166000908152600b602052604090209161ffff1615615b2957606061ffff615b25935460901c16910135614f44565b0490565b606061ffff615b25935460801c16910135614f44565b80600052600760205260406000205415600014615b99576006546801000000000000000081101561438a57615b8061595482600185940160065560066157fa565b9055600654906000526007602052604060002055600190565b50600090565b60008281526001820160205260409020546159a1578054906801000000000000000082101561438a5782615bdd6159548460018096018555846157fa565b905580549260005201602052604060002055600190565b73ffffffffffffffffffffffffffffffffffffffff615c83911691604092600080855193615c22878661440d565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d15615d2c573d91615c678361444e565b92615c748751948561440d565b83523d6000602085013e616040565b80519081615c9057505050565b602080615ca193830101910161575c565b15615ca95750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091616040565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215612e5e57016020813591019167ffffffffffffffff8211612e5e578136038313612e5e57565b9182549060ff8260a01c16158015616038575b616032576fffffffffffffffffffffffffffffffff82169160018501908154615ddc63ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614f86565b9081615f94575b5050848110615f485750838310615e3d575050615e126fffffffffffffffffffffffffffffffff928392614f86565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615edc5781615e5591614f86565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211614f5757615ea3615ea89273ffffffffffffffffffffffffffffffffffffffff9661554e565b615616565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82869293961161600857615faf92614ff29160801c90614f44565b808410156160035750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615de3565b615fba565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615d97565b919290156160bb5750815115616054575090565b3b1561605d5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156160ce5750805190602001fd5b6121b9906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061448856fea164736f6c634300081a000a",
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

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

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

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*LockReleaseTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*LockReleaseTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

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
