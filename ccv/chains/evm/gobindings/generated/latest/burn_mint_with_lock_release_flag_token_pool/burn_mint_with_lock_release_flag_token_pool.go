// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_with_lock_release_flag_token_pool

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

var BurnMintWithLockReleaseFlagTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60e080604052346102045760a0816159de803803809161001f8285610255565b8339810103126102045780516001600160a01b03811690819003610204576100496020830161028e565b6100556040840161029c565b9161006e60806100676060870161029c565b950161029c565b93331561024457600180546001600160a01b0319163317905581158015610233575b8015610222575b610211578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816101d0575b506101a5575b5060a052600380546001600160a01b039283166001600160a01b0319918216179091556002805493909216921691909117905560405161572d90816102b182396080518181816116a8015281816118b10152818161195401528181611ca5015281816122ac0152818161246301528181612c6501528181612e5e01528181612f03015281816132670152818161346d0152818161363b01528181613b960152613bf0015260a051818181613a5c0152614ff2015260c051818181610b82015281816117430152818161234601528181612d0001526135080152f35b60ff1660ff82168181036101b957506100ca565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610209575b816101ec60209383610255565b81010312610204576101fd9061028e565b90386100c4565b600080fd5b3d91506101df565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610097565b506001600160a01b03851615610090565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761027857604052565b634e487b7160e01b600052604160045260246000fd5b519060ff8216820361020457565b51906001600160a01b03821682036102045756fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613c9b57508063181f5a7714613c1457806321df0da714613bc3578063240028e814613b5f5780632422ac4514613a8057806324f65ee714613a425780632c063404146139a957806337a3210d1461397557806339077537146133ec578063489a68f214612bea5780634c5ef0ed14612ba35780634e921c3014612b0457806362ddd3c414612a7d5780637437ff9f14612a2f57806379ba5097146129685780638926f54f1461292257806389720a621461285b5780638da5cb5b146128275780639a4575b914612229578063a42a7b8b146120c2578063acfecf9114611fca578063ae39a25714611e3f578063b1c71c65146115f8578063b7946580146115bb578063bfeffd3f1461150f578063c4bffe2b146113e4578063c7230a6014611131578063d8aa3f4014610ff7578063dc04fa1f14610ba6578063dc0bd97114610b55578063dcbd41bc14610951578063e8a1da171461028d578063f2fde38b146101be5763fa41d79c1461019757600080fd5b346101bb57806003193601126101bb57602061ffff60025460a01c16604051908152f35b80fd5b50346101bb5760206003193601126101bb5773ffffffffffffffffffffffffffffffffffffffff6101ed613dfc565b6101f5614aa9565b1633811461026557807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101bb5760406003193601126101bb5760043567ffffffffffffffff81116107aa576102bf90369060040161410c565b9060243567ffffffffffffffff811161094d57906102e28492369060040161410c565b9390916102ed614aa9565b83905b8282106107b25750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156107ae578060051b830135858112156107a6578301610120813603126107a6576040519461035486613fbd565b61035d82613eb9565b8652602082013567ffffffffffffffff81116107aa5782019436601f870112156107aa5785359561038d8761447a565b9661039b6040519889613fd9565b80885260208089019160051b830101903682116107a65760208301905b828210610773575050505060208701958652604083013567ffffffffffffffff811161076f576103eb903690850161407f565b916040880192835261041561040336606087016148b9565b9460608a0195865260c03691016148b9565b9560808901968752835151156107475761043967ffffffffffffffff8a51166153b6565b156107105767ffffffffffffffff8951168252600860205260408220610460865182614d5e565b61046e885160028301614d5e565b6004855191019080519067ffffffffffffffff82116106e3576104918354614707565b601f81116106a8575b50602090601f8311600114610609576104e892918691836105fe575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b88518051821015610522579061051c6001926105158367ffffffffffffffff8f5116926146c4565b5190614af4565b016104ed565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29391999750956105f067ffffffffffffffff60019796949851169251935191516105bc61058760405196879687526101006020880152610100870190613d9d565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610323565b015190508e806104b6565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106106905750908460019594939210610659575b505050811b0190556104eb565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d808061064c565b92936020600181928786015181550195019301610636565b6106d39084875260208720601f850160051c810191602086106106d9575b601f0160051c0190614955565b8d61049a565b90915081906106c6565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116107a257602091610797839283369189010161407f565b8152019101906103b8565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff6107d46107cf8486889a9699979a61488c565b614428565b16916107df83615089565b156109215782845260086020526107fb60056040862001615026565b94845b865181101561083457600190858752600860205261082d60056040892001610826838b6146c4565b519061521f565b50016107fe565b50939692909450949094808752600860205260056040882088815588600182015588600282015588600382015588600482016108708154614707565b806108e0575b50505001805490888155816108c2575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019091949392946102f0565b885260208820908101905b81811015610886578881556001016108cd565b601f81116001146108f65750555b888a80610876565b8183526020832061091191601f01861c810190600101614955565b80825281602081209155556108ee565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346101bb5760206003193601126101bb5760043567ffffffffffffffff81116107aa5761098390369060040161413d565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580610b33575b610b0757825b8181106109b6578380f35b6109c181838561483e565b67ffffffffffffffff6109d382614428565b16906109ec826000526007602052604060002054151590565b15610adb57907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610a9b610a75602060019897018b610a2d8261484e565b15610aa2578790526004602052610a5460408d20610a4e36604088016148b9565b90614d5e565b868c526005602052610a7060408d20610a4e3660a088016148b9565b61484e565b916040519215158352610a8e6020840160408301614911565b60a0608084019101614911565ba2016109ab565b60026040828a610a7094526008602052610ac4828220610a4e36858c016148b9565b8a815260086020522001610a4e3660a088016148b9565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600154163314156109a5565b50346101bb57806003193601126101bb57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101bb5760406003193601126101bb5760043567ffffffffffffffff81116107aa57610bd890369060040161413d565b60243567ffffffffffffffff811161094d57610bf890369060040161410c565b919092610c03614aa9565b845b828110610c6f57505050825b818110610c1c578380f35b8067ffffffffffffffff610c366107cf600194868861488c565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610c11565b610c7d6107cf82858561483e565b610c8882858561483e565b90602082019060e0830190610c9c8261484e565b15610fc25760a0840161271061ffff610cb48361485b565b161015610fb35760c085019161271061ffff610ccf8561485b565b161015610f7b5763ffffffff610ce48661486a565b1615610f465767ffffffffffffffff1694858c52600b60205260408c20610d0a8661486a565b63ffffffff16908054906040840191610d228361486a565b60201b67ffffffff0000000016936060860194610d3e8661486a565b60401b6bffffffff0000000000000000169660800196610d5d8861486a565b60601b6fffffffff0000000000000000000000001691610d7c8a61485b565b60801b71ffff000000000000000000000000000000001693610d9d8c61485b565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610e508761484e565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610ea19061487b565b63ffffffff168752610eb29061487b565b63ffffffff166020870152610ec69061487b565b63ffffffff166040860152610eda9061487b565b63ffffffff166060850152610eee90613efd565b61ffff166080840152610f0090613efd565b61ffff1660a0830152610f1290613ece565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c05565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff610f8a8661485b565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff610f8a60249361485b565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346101bb5760806003193601126101bb57611011613dfc565b5061101a613ea2565b611022613eec565b5060643567ffffffffffffffff811161076f579167ffffffffffffffff60409261105260e0953690600401613f0c565b50508260c0855161106281613fa1565b82815282602082015282878201528260608201528260808201528260a08201520152168152600b602052206040519061109a82613fa1565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346101bb5760406003193601126101bb5760043567ffffffffffffffff81116107aa5761116390369060040161410c565b9061116c613e47565b9173ffffffffffffffffffffffffffffffffffffffff60015416331415806113c2575b6113965773ffffffffffffffffffffffffffffffffffffffff831690811561136e57845b8181106111be578580f35b73ffffffffffffffffffffffffffffffffffffffff6111e66111e183858861488c565b614407565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa90811561136357889161132e575b508061123b575b50506001016111b3565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff8a16602484015260448084018590528352918a919061129c606482613fd9565b519082865af1156113235787513d61131a5750813b155b6112ee5790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a39038611231565b602488837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b600114156112b3565b6040513d89823e3d90fd5b90506020813d821161135b575b8161134860209383613fd9565b8101031261135757513861122a565b8780fd5b3d915061133b565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600c541633141561118f565b50346101bb57806003193601126101bb57604051906006548083528260208101600684526020842092845b8181106114f657505061142492500383613fd9565b81516114486114328261447a565b916114406040519384613fd9565b80835261447a565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b84518110156114a7578067ffffffffffffffff611494600193886146c4565b51166114a082866146c4565b5201611475565b50925090604051928392602084019060208552518091526040840192915b8181106114d3575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016114c5565b845483526001948501948794506020909301920161140f565b50346101bb5760206003193601126101bb5760043573ffffffffffffffffffffffffffffffffffffffff81168091036107aa5761154a614aa9565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a1161760035580f35b50346101bb5760206003193601126101bb576115f46115e06115db613e8b565b61481c565b604051918291602083526020830190613d9d565b0390f35b50346101bb5760606003193601126101bb576004359067ffffffffffffffff82116101bb578160040160a060031984360301126107aa57611637613edb565b60443567ffffffffffffffff811161094d579061165c61168293923690600401613f0c565b93906116666146ab565b5061166f6146ab565b5061167a8385615353565b94369161401a565b608486019061169082614407565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611df557602487019677ffffffffffffffff000000000000000000000000000000006116f689614428565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611363578891611dc6575b50611d9e5767ffffffffffffffff61178a89614428565b166117a2816000526007602052604060002054151590565b15611d7357602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611363578890611d26575b73ffffffffffffffffffffffffffffffffffffffff9150163303611cfa5760648101359461ffff61183488886149ae565b9516948515611c495761ffff60025460a01c168015611c2157808710611bf157507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff6118888c614428565b1691828b526004602052806118d960408d2073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161546b565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff600354169283611ac9575b50505050505090611932916149ae565b9061193c83614428565b5073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b156101bb576040517f42966c68000000000000000000000000000000000000000000000000000000008152836004820152818160248183875af18015611abe57611aa9575b611a9f84611a326115db88877ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108567ffffffffffffffff6119fc85614428565b6040805173ffffffffffffffffffffffffffffffffffffffff9690961686523360208701528501929092521691606090a2614428565b90611a3b614feb565b60405192611a4884613f85565b835260208301908152604051907ffa7c07de00000000000000000000000000000000000000000000000000000000602083015260208252611a8a604083613fd9565b526040519283926040845260408401906140e2565b9060208301520390f35b611ab4828092613fd9565b6101bb57806119bc565b6040513d84823e3d90fd5b833b15611bed57889493929185918b604051988997889687957f5c3af7ca000000000000000000000000000000000000000000000000000000008752600487016060905280611b1791614f9b565b6064880160a09052610104880190611b2e926144b3565b93611b3890613eb9565b67ffffffffffffffff166084870152604401611b5390613e6a565b73ffffffffffffffffffffffffffffffffffffffff1660a48601528b60c4860152611b7d90613e6a565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526024840152828103600319016044840152611bb291613d9d565b03925af18015611be257908491611bcd575b80808080611922565b81611bd791613fd9565b61076f578238611bc4565b6040513d86823e3d90fd5b8880fd5b89604491887f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b60048a7f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611c7c8c614428565b1691828b52600860205280611ccd60408d2073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161546b565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611902565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611d6b575b81611d4060209383613fd9565b8101031261135757611d6673ffffffffffffffffffffffffffffffffffffffff91614492565b611803565b3d9150611d33565b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611de8915060203d602011611dee575b611de08183613fd9565b810190614a40565b38611773565b503d611dd6565b60248673ffffffffffffffffffffffffffffffffffffffff611e1685614407565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346101bb5760606003193601126101bb57611e59613dfc565b90611e62613e47565b6044359273ffffffffffffffffffffffffffffffffffffffff841680850361094d57611e8c614aa9565b73ffffffffffffffffffffffffffffffffffffffff82168015611fa25794611f9c917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c556040519384938491604091949373ffffffffffffffffffffffffffffffffffffffff809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346101bb5767ffffffffffffffff611fe23661409d565b929091611fed614aa9565b1691612006836000526007602052604060002054151590565b156109215782845260086020526120356005604086200161202836848661401a565b602081519101209061521f565b1561207a57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76916120746040519283926020845260208401916144b3565b0390a280f35b826120be836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916144b3565b0390fd5b50346101bb5760206003193601126101bb5767ffffffffffffffff6120e5613e8b565b16815260086020526120fc60056040832001615026565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061214161212b8361447a565b926121396040519485613fd9565b80845261447a565b01835b818110612218575050825b82518110156121955780612165600192856146c4565b51855260096020526121796040862061475a565b61218382856146c4565b5261218e81846146c4565b500161214f565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106121cd57505050500390f35b91936020612208827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613d9d565b96019201920185949391926121be565b806060602080938601015201612144565b50346101bb5760206003193601126101bb576004359067ffffffffffffffff82116101bb57816004019160a060031982360301126107aa576122696146ab565b506122726146ab565b50602092604051926122848585613fd9565b8084526084830161229481614407565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361280657602484019477ffffffffffffffff000000000000000000000000000000006122fa87614428565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152878160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611be25784916127e9575b506127c15767ffffffffffffffff61238d87614428565b166123a5816000526007602052604060002054151590565b15612796578773ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611be257849061274e575b73ffffffffffffffffffffffffffffffffffffffff915016330361272257606485013594859467ffffffffffffffff61243d89614428565b1680865260088a5261248b6040872073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016998a9161546b565b6040805173ffffffffffffffffffffffffffffffffffffffff8a168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a273ffffffffffffffffffffffffffffffffffffffff600354169283612605575b505050505061250184614428565b50823b156101bb576040517f42966c68000000000000000000000000000000000000000000000000000000008152826004820152818160248183885af18015611abe576125f0575b856125866115db87877ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108867ffffffffffffffff6119fc85614428565b9061258f614feb565b6040519261259c84613f85565b8352818301908152604051907ffa7c07de00000000000000000000000000000000000000000000000000000000838301528282526125db604083613fd9565b526115f46040519282849384528301906140e2565b6125fb828092613fd9565b6101bb5780612549565b833b1561271e5791858094928a9694604051978896879586947f5c3af7ca00000000000000000000000000000000000000000000000000000000865260048601606090528061265391614f9b565b6064870160a0905261010487019061266a926144b3565b9261267490613eb9565b67ffffffffffffffff16608486015260440161268f90613e6a565b73ffffffffffffffffffffffffffffffffffffffff1660a48501528b60c48501526126b990613e6a565b73ffffffffffffffffffffffffffffffffffffffff1660e48401528360248401528281036003190160448401526126ef91613d9d565b03925af18015611abe57908291612709575b8080806124f3565b8161271391613fd9565b6101bb578038612701565b8580fd5b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d831161278f575b6127648183613fd9565b8101031261094d5761278a73ffffffffffffffffffffffffffffffffffffffff91614492565b612405565b503d61275a565b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6128009150883d8a11611dee57611de08183613fd9565b38612376565b9073ffffffffffffffffffffffffffffffffffffffff611e16602493614407565b50346101bb57806003193601126101bb57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346101bb5760c06003193601126101bb57612875613dfc565b61287d613ea2565b9060643561ffff8116810361094d5760843567ffffffffffffffff81116107a6576128ac903690600401613f0c565b9160a4359360028510156107a2576128c795604435916144f2565b90604051918291602083016020845282518091526020604085019301915b8181106128f3575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016128e5565b50346101bb5760206003193601126101bb57602061295e67ffffffffffffffff61294a613e8b565b166000526007602052604060002054151590565b6040519015158152f35b50346101bb57806003193601126101bb57805473ffffffffffffffffffffffffffffffffffffffff81163303612a07577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346101bb57806003193601126101bb57600254600a54600c546040805173ffffffffffffffffffffffffffffffffffffffff94851681529284166020840152921691810191909152606090f35b50346101bb57612a8c3661409d565b612a9893929193614aa9565b67ffffffffffffffff8216612aba816000526007602052604060002054151590565b15612ad95750612ad69293612ad091369161401a565b90614af4565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346101bb5760206003193601126101bb5760043561ffff81169081810361076f577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb91602091612b53614aa9565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006002549260a01b16911617600255604051908152a180f35b50346101bb5760406003193601126101bb57612bbd613e8b565b906024359067ffffffffffffffff82116101bb57602061295e84612be4366004870161407f565b9061443d565b50346101bb5760406003193601126101bb5760043567ffffffffffffffff81116107aa57806004019161010060031983360301126101bb57612c2a613edb565b9181604051612c3881613f3a565b526064810135926084820194612c4d86614407565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036133cb57602483019577ffffffffffffffff00000000000000000000000000000000612cb388614428565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561334e5786916133ac575b506133845767ffffffffffffffff612d4788614428565b16612d5f816000526007602052604060002054151590565b1561335957602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561334e57869161332f575b501561330357612dd687614428565b92612df360a4860194612be4612dec8787614a58565b369161401a565b156132bc5761ffff1680156132085767ffffffffffffffff612e1489614428565b1680875260056020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158880612e8660408b2073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161546b565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff60035416928361302d575b505050505060440192612ee184614407565b91612eeb82614428565b5073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b156107aa576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101859052818180604481015b038183885af18015611abe57613018575b5050608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff612fe6612fe07ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096614428565b92614407565b60405196875233898801521660408601528560608601521692a26040519061300d82613f3a565b815260405190518152f35b613023828092613fd9565b6101bb5780612f8e565b833b156107a2579186888195938b9795604051988997889687957f5eff3bf700000000000000000000000000000000000000000000000000000000875260048701606090528d61307d8680614f9b565b60648a0161010090526101648a0190613095926144b3565b9461309f90613eb9565b67ffffffffffffffff1660848901526044016130ba90613e6a565b73ffffffffffffffffffffffffffffffffffffffff1660a488015260c48701526130e390613e6a565b73ffffffffffffffffffffffffffffffffffffffff1660e48601526131089083614f9b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161010487015261313d92916144b3565b61314a60c48c0183614f9b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161012487015261317f92916144b3565b9060e48b0161318d91614f9b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101448601526131c292916144b3565b908b6024840152604483015203925af180156131fd579083916131e8575b808080612ecf565b816131f291613fd9565b6107aa5781386131e0565b6040513d85823e3d90fd5b67ffffffffffffffff61321a89614428565b1680875260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c888061328f600260408c200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161546b565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612eaf565b6132c68484614a58565b6120be6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916144b3565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b613348915060203d602011611dee57611de08183613fd9565b38612dc7565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6133c5915060203d602011611dee57611de08183613fd9565b38612d30565b60248473ffffffffffffffffffffffffffffffffffffffff611e1689614407565b50346101bb5760206003193601126101bb576004359067ffffffffffffffff82116101bb57816004019161010060031982360301126107aa578160405161343281613f3a565b528160405161344081613f3a565b52606481013590608481019061345582614407565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361395457602481019177ffffffffffffffff000000000000000000000000000000006134bb84614428565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561334e578691613935575b506133845767ffffffffffffffff61354f84614428565b16613567816000526007602052604060002054151590565b1561335957602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561334e578691613916575b5015613303576135de83614428565b906135f460a4840192612be4612dec858b614a58565b1561390c5794809691859667ffffffffffffffff61361187614428565b1680845260086020526136636002604086200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016998a9161546b565b6040805173ffffffffffffffffffffffffffffffffffffffff8a168152602081018b90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a273ffffffffffffffffffffffffffffffffffffffff600354169081613744575b5050505050604401936136dd85614407565b6136e683614428565b50833b156107aa576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024810185905281818060448101612f7d565b813b1561094d5783898896604051978896879586947f5eff3bf7000000000000000000000000000000000000000000000000000000008652600486016060905261378e8580614f9b565b6064880161010090526101648801906137a6926144b3565b936137b090613eb9565b67ffffffffffffffff1660848701526137cb60448e01613e6a565b73ffffffffffffffffffffffffffffffffffffffff1660a487015260c48601526137f490613e6a565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526138199083614f9b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8584030161010486015261384e92916144b3565b61385b60c48b0183614f9b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8584030161012486015261389092916144b3565b9060e48a0161389e91614f9b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c848403016101448501526138d392916144b3565b8b602483015282604483015203925af1801561334e576138f7575b858180806136cb565b946139058160449397613fd9565b94906138ee565b6132c68288614a58565b61392f915060203d602011611dee57611de08183613fd9565b386135cf565b61394e915060203d602011611dee57611de08183613fd9565b38613538565b60248473ffffffffffffffffffffffffffffffffffffffff611e1685614407565b50346101bb57806003193601126101bb57602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b50346101bb5760c06003193601126101bb576139c3613dfc565b506139cc613ea2565b6139d4613e24565b506084359161ffff831683036101bb5760a4359067ffffffffffffffff82116101bb5760a063ffffffff8061ffff613a1b8888613a143660048b01613f0c565b5050614282565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346101bb57806003193601126101bb57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101bb5760406003193601126101bb57613a9a613e8b565b6024359182151583036101bb57610140613b5d613ab785856141ff565b613b0d60409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346101bb5760206003193601126101bb57602090613b7c613dfc565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346101bb57806003193601126101bb57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101bb57806003193601126101bb57506115f4604051613c37606082613fd9565b602e81527f4275726e4d696e74576974684c6f636b52656c65617365466c6167546f6b656e60208201527f506f6f6c20312e372e302d6465760000000000000000000000000000000000006040820152604051918291602083526020830190613d9d565b9050346107aa5760206003193601126107aa576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361076f57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115613d73575b8115613d49575b8115613d1f575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613d18565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613d11565b7f331710310000000000000000000000000000000000000000000000000000000081149150613d0a565b919082519283825260005b848110613de75750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613da8565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203613e1f57565b600080fd5b6064359073ffffffffffffffffffffffffffffffffffffffff82168203613e1f57565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203613e1f57565b359073ffffffffffffffffffffffffffffffffffffffff82168203613e1f57565b6004359067ffffffffffffffff82168203613e1f57565b6024359067ffffffffffffffff82168203613e1f57565b359067ffffffffffffffff82168203613e1f57565b35908115158203613e1f57565b6024359061ffff82168203613e1f57565b6044359061ffff82168203613e1f57565b359061ffff82168203613e1f57565b9181601f84011215613e1f5782359167ffffffffffffffff8311613e1f5760208381860195010111613e1f57565b6020810190811067ffffffffffffffff821117613f5657604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117613f5657604052565b60e0810190811067ffffffffffffffff821117613f5657604052565b60a0810190811067ffffffffffffffff821117613f5657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613f5657604052565b92919267ffffffffffffffff8211613f565760405191614062601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184613fd9565b829481845281830111613e1f578281602093846000960137010152565b9080601f83011215613e1f5781602061409a9335910161401a565b90565b906040600319830112613e1f5760043567ffffffffffffffff81168103613e1f57916024359067ffffffffffffffff8211613e1f576140de91600401613f0c565b9091565b61409a9160206140fb8351604084526040840190613d9d565b920151906020818403910152613d9d565b9181601f84011215613e1f5782359167ffffffffffffffff8311613e1f576020808501948460051b010111613e1f57565b9181601f84011215613e1f5782359167ffffffffffffffff8311613e1f576020808501948460081b010111613e1f57565b6040519061417b82613fbd565b60006080838281528260208201528260408201528260608201520152565b906040516141a681613fbd565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff9161421161416e565b5061421a61416e565b5061424e5716600052600860205260406000209061409a614242600261424761424286614199565b6149bb565b9401614199565b16908160005260046020526142696142426040600020614199565b91600052600560205261409a6142426040600020614199565b9061ffff8060025460a01c16911692831515928380946143ff575b6143d55767ffffffffffffffff16600052600b602052604060002091604051926142c684613fa1565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a01526143ba5761435b575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b81939750809294501061438a57505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b50821561429d565b3573ffffffffffffffffffffffffffffffffffffffff81168103613e1f5790565b3567ffffffffffffffff81168103613e1f5790565b9067ffffffffffffffff61409a92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613f565760051b60200190565b519073ffffffffffffffffffffffffffffffffffffffff82168203613e1f57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff600354169586156146895761458d9467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c48601916144b3565b91600282101561465a578380600094819460a483015203915afa90811561464e576000916145b9575090565b3d8083833e6145c88183613fd9565b81019060208183031261076f5780519067ffffffffffffffff821161094d570181601f8201121561076f578051906145ff8261447a565b9361460d6040519586613fd9565b82855260208086019360051b8301019384116101bb5750602001905b8282106146365750505090565b6020809161464384614492565b815201910190614629565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b505050505050505060405161469f602082613fd9565b60008152600036813790565b604051906146b882613f85565b60606020838281520152565b80518210156146d85760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015614750575b602083101461472157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614716565b906040519182600082549261476e84614707565b80845293600181169081156147dc5750600114614795575b5061479392500383613fd9565b565b90506000929192526020600020906000915b8183106147c05750509060206147939282010138614786565b60209193508060019154838589010152019101909184926147a7565b602093506147939592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614786565b67ffffffffffffffff16600052600860205261409a600460406000200161475a565b91908110156146d85760081b0190565b358015158103613e1f5790565b3561ffff81168103613e1f5790565b3563ffffffff81168103613e1f5790565b359063ffffffff82168203613e1f57565b91908110156146d85760051b0190565b35906fffffffffffffffffffffffffffffffff82168203613e1f57565b9190826060910312613e1f576040516060810181811067ffffffffffffffff821117613f5657604052604061490c8183956148f381613ece565b85526149016020820161489c565b60208601520161489c565b910152565b6fffffffffffffffffffffffffffffffff61494f6040809361493281613ece565b15158652836149436020830161489c565b1660208701520161489c565b16910152565b818110614960575050565b60008155600101614955565b8181029291811591840414171561497f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9190820391821161497f57565b6149c361416e565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614a206020850193614a1a614a0d63ffffffff875116426149ae565b856080890151169061496c565b90614f8e565b80821015614a3957505b16825263ffffffff4216905290565b9050614a2a565b90816020910312613e1f57518015158103613e1f5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613e1f570180359067ffffffffffffffff8211613e1f57602001918136038313613e1f57565b73ffffffffffffffffffffffffffffffffffffffff600154163303614aca57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115614d345767ffffffffffffffff81516020830120921691826000526008602052614b29816005604060002001615416565b15614cf05760005260096020526040600020815167ffffffffffffffff8111613f5657614b568254614707565b601f8111614cbe575b506020601f8211600114614bf85791614bd2827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614be895600091614bed575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613d9d565b0390a2565b905084015138614ba1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110614ca6575092614be89492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614c6f575b5050811b0190556115e0565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880614c63565b9192602060018192868a015181550194019201614c28565b614cea90836000526020600020601f840160051c810191602085106106d957601f0160051c0190614955565b38614b5f565b50906120be6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613d9d565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b815191929115614ee0576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff60208501511610614e7d5761479391925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b606483614ede604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408401511615801590614f6f575b614f0e576147939192614da1565b606483614ede604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020840151161515614f00565b9190820180921161497f57565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215613e1f57016020813591019167ffffffffffffffff8211613e1f578136038313613e1f57565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001660208201526020815261409a604082613fd9565b906040519182815491828252602082019060005260206000209260005b81811061505857505061479392500383613fd9565b8454835260019485019487945060209093019201615043565b80548210156146d85760005260206000200190600090565b6000818152600760205260409020548015615218577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161497f57600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161497f578181036151a9575b505050600654801561517a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615137816006615071565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6152006151ba6151cb936006615071565b90549060031b1c9283926006615071565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260076020526040600020553880806150fe565b5050600090565b906001820191816000528260205260406000205480151560001461534a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161497f578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161497f57818103615313575b5050508054801561517a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906152d48282615071565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6153336153236151cb9386615071565b90549060031b1c92839286615071565b90556000528360205260406000205538808061529c565b50505050600090565b906127109167ffffffffffffffff61536d60208301614428565b166000908152600b602052604090209161ffff16156153a057606061ffff61539c935460901c1691013561496c565b0490565b606061ffff61539c935460801c1691013561496c565b806000526007602052604060002054156000146154105760065468010000000000000000811015613f56576153f76151cb8260018594016006556006615071565b9055600654906000526007602052604060002055600190565b50600090565b60008281526001820160205260409020546152185780549068010000000000000000821015613f5657826154546151cb846001809601855584615071565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615718575b615712576fffffffffffffffffffffffffffffffff821691600185019081546154c363ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426149ae565b9081615674575b505084811061562857508383106155245750506154f96fffffffffffffffffffffffffffffffff9283926149ae565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c9283156155bc578161553c916149ae565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019381851161497f5773ffffffffffffffffffffffffffffffffffffffff9461558791614f8e565b047fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b8286929396116156e85761568f92614a1a9160801c9061496c565b808410156156e35750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806154ca565b61569a565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561547e56fea164736f6c634300081a000a",
}

var BurnMintWithLockReleaseFlagTokenPoolABI = BurnMintWithLockReleaseFlagTokenPoolMetaData.ABI

var BurnMintWithLockReleaseFlagTokenPoolBin = BurnMintWithLockReleaseFlagTokenPoolMetaData.Bin

func DeployBurnMintWithLockReleaseFlagTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnMintWithLockReleaseFlagTokenPool, error) {
	parsed, err := BurnMintWithLockReleaseFlagTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintWithLockReleaseFlagTokenPoolBin), backend, token, localTokenDecimals, advancedPoolHooks, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintWithLockReleaseFlagTokenPool{address: address, abi: *parsed, BurnMintWithLockReleaseFlagTokenPoolCaller: BurnMintWithLockReleaseFlagTokenPoolCaller{contract: contract}, BurnMintWithLockReleaseFlagTokenPoolTransactor: BurnMintWithLockReleaseFlagTokenPoolTransactor{contract: contract}, BurnMintWithLockReleaseFlagTokenPoolFilterer: BurnMintWithLockReleaseFlagTokenPoolFilterer{contract: contract}}, nil
}

type BurnMintWithLockReleaseFlagTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnMintWithLockReleaseFlagTokenPoolCaller
	BurnMintWithLockReleaseFlagTokenPoolTransactor
	BurnMintWithLockReleaseFlagTokenPoolFilterer
}

type BurnMintWithLockReleaseFlagTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnMintWithLockReleaseFlagTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnMintWithLockReleaseFlagTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnMintWithLockReleaseFlagTokenPoolSession struct {
	Contract     *BurnMintWithLockReleaseFlagTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintWithLockReleaseFlagTokenPoolCallerSession struct {
	Contract *BurnMintWithLockReleaseFlagTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnMintWithLockReleaseFlagTokenPoolTransactorSession struct {
	Contract     *BurnMintWithLockReleaseFlagTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintWithLockReleaseFlagTokenPoolRaw struct {
	Contract *BurnMintWithLockReleaseFlagTokenPool
}

type BurnMintWithLockReleaseFlagTokenPoolCallerRaw struct {
	Contract *BurnMintWithLockReleaseFlagTokenPoolCaller
}

type BurnMintWithLockReleaseFlagTokenPoolTransactorRaw struct {
	Contract *BurnMintWithLockReleaseFlagTokenPoolTransactor
}

func NewBurnMintWithLockReleaseFlagTokenPool(address common.Address, backend bind.ContractBackend) (*BurnMintWithLockReleaseFlagTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintWithLockReleaseFlagTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintWithLockReleaseFlagTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPool{address: address, abi: abi, BurnMintWithLockReleaseFlagTokenPoolCaller: BurnMintWithLockReleaseFlagTokenPoolCaller{contract: contract}, BurnMintWithLockReleaseFlagTokenPoolTransactor: BurnMintWithLockReleaseFlagTokenPoolTransactor{contract: contract}, BurnMintWithLockReleaseFlagTokenPoolFilterer: BurnMintWithLockReleaseFlagTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnMintWithLockReleaseFlagTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnMintWithLockReleaseFlagTokenPoolCaller, error) {
	contract, err := bindBurnMintWithLockReleaseFlagTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolCaller{contract: contract}, nil
}

func NewBurnMintWithLockReleaseFlagTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintWithLockReleaseFlagTokenPoolTransactor, error) {
	contract, err := bindBurnMintWithLockReleaseFlagTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolTransactor{contract: contract}, nil
}

func NewBurnMintWithLockReleaseFlagTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintWithLockReleaseFlagTokenPoolFilterer, error) {
	contract, err := bindBurnMintWithLockReleaseFlagTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolFilterer{contract: contract}, nil
}

func bindBurnMintWithLockReleaseFlagTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintWithLockReleaseFlagTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.BurnMintWithLockReleaseFlagTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.BurnMintWithLockReleaseFlagTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.BurnMintWithLockReleaseFlagTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getAdvancedPoolHooks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetAdvancedPoolHooks(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetAdvancedPoolHooks(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetCurrentRateLimiterState(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetCurrentRateLimiterState(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetFee(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetFee(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetMinBlockConfirmation(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetMinBlockConfirmation(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRemotePools(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRemotePools(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRemoteToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRemoteToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRmnProxy(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRmnProxy(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetSupportedChains(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetSupportedChains(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetTokenDecimals(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetTokenDecimals(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsSupportedChain(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsSupportedChain(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsSupportedToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, token)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsSupportedToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, token)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) Owner() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.Owner(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.Owner(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SupportsInterface(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, interfaceId)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SupportsInterface(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, interfaceId)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.TypeAndVersion(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.TypeAndVersion(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AcceptOwnership(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AcceptOwnership(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AddRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AddRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyChainUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyChainUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.LockOrBurn(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.LockOrBurn(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.LockOrBurn0(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.LockOrBurn0(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ReleaseOrMint(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ReleaseOrMint(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ReleaseOrMint0(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ReleaseOrMint0(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RemoveRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RemoveRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setMinBlockConfirmation", minBlockConfirmation)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetMinBlockConfirmation(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, minBlockConfirmation)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetMinBlockConfirmation(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, minBlockConfirmation)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetRateLimitConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetRateLimitConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.TransferOwnership(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, to)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.TransferOwnership(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, to)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "updateAdvancedPoolHooks", newHook)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.UpdateAdvancedPoolHooks(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newHook)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.UpdateAdvancedPoolHooks(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newHook)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.WithdrawFeeTokens(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, feeTokens, recipient)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.WithdrawFeeTokens(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, feeTokens, recipient)
}

type BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdatedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated struct {
	OldHook common.Address
	NewHook common.Address
	Raw     types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdatedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdatedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "AdvancedPoolHooksUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseAdvancedPoolHooksUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainAdded)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainAdded)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolChainAdded)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainAdded, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolChainAdded)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainRemoved)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainRemoved)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolChainRemoved)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainRemoved, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolChainRemoved)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAdmin       common.Address
	Raw            types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSetIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSetIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSetIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "MinBlockConfirmationSet", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseMinBlockConfirmationSet(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRateLimitConfiguredIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRateLimitConfiguredIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated) Topic() common.Hash {
	return common.HexToHash("0xbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209")
}

func (BurnMintWithLockReleaseFlagTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnMintWithLockReleaseFlagTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet) Topic() common.Hash {
	return common.HexToHash("0xa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb")
}

func (BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPool) Address() common.Address {
	return _BurnMintWithLockReleaseFlagTokenPool.address
}

type BurnMintWithLockReleaseFlagTokenPoolInterface interface {
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

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolAdvancedPoolHooksUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainRemoved, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, error)

	FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSetIterator, error)

	WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet) (event.Subscription, error)

	ParseMinBlockConfirmationSet(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolMinBlockConfirmationSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
