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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610100806040523461023d5760a081615d8b8038038091610020828561028e565b83398101031261023d5780516001600160a01b0381169081900361023d5761004a602083016102c7565b610056604084016102d5565b9161006f6080610068606087016102d5565b95016102d5565b93331561027d57600180546001600160a01b031916331790558115801561026c575b801561025b575b61024a578160209160049360805260c0526040519283809263313ce56760e01b82525afa60009181610209575b506101de575b5060a0526001600160a01b0390811660e052600280546001600160a01b03191692909116919091179055604051615aa190816102ea82396080518181816117990152818161199801528181611a5401528181611d19015281816121970152818161234e01528181612acd01528181612cc601528181612d81015281816130fa015281816132fb015281816134cb01528181613a9201528181613aec0152614f4f015260a051818181611add01528181613958015281816149ce01528181614ad70152614c61015260c051818181610c9d015281816118340152818161223101528181612b680152613396015260e051818181611a00015281816123d501528181612d2e01528181613552015261441f0152f35b60ff1660ff82168181036101f257506100cb565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610242575b816102256020938361028e565b8101031261023d57610236906102c7565b90386100c5565b600080fd5b3d9150610218565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610098565b506001600160a01b03851615610091565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b038211908210176102b157604052565b634e487b7160e01b600052604160045260246000fd5b519060ff8216820361023d57565b51906001600160a01b038216820361023d5756fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613b7157508063181f5a7714613b1057806321df0da714613abf578063240028e814613a5b5780632422ac451461397c57806324f65ee71461393e5780632c063404146138a5578063390775371461327f578063489a68f214612a285780634c5ef0ed146129e157806362ddd3c41461295a5780637437ff9f1461290957806379ba5097146128425780638926f54f146127fc57806389720a62146127355780638da5cb5b146127015780639a4575b91461211d578063a42a7b8b14611fb6578063acfecf9114611ebe578063b1c71c6514611700578063b7946580146116c3578063c4bffe2b14611598578063c7230a601461124c578063d8aa3f4014611112578063dc04fa1f14610cc1578063dc0bd97114610c70578063dcbd41bc14610a6c578063e8a1da17146103ac578063f2fde38b146102dd5763fdf168751461016b57600080fd5b346102da5760606003193601126102da57610184613dec565b9061018d613ea8565b6044359273ffffffffffffffffffffffffffffffffffffffff84168085036102d6576101b7614c83565b73ffffffffffffffffffffffffffffffffffffffff821680156102ae57946102a8917fba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a1179596600254907fffffffffffffffffffff0000000000000000000000000000000000000000000075ffff00000000000000000000000000000000000000008860a01b16921617176002557fffffffffffffffffffffffff000000000000000000000000000000000000000060095416176009556040519384938491604091949361ffff73ffffffffffffffffffffffffffffffffffffffff9283606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b80fd5b50346102da5760206003193601126102da5773ffffffffffffffffffffffffffffffffffffffff61030c613dec565b610314614c83565b1633811461038457807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346102da5760406003193601126102da5760043567ffffffffffffffff81116108c9576103de903690600401613fcb565b9060243567ffffffffffffffff81116102d6579061040184923690600401613fcb565b93909161040c614c83565b83905b8282106108d15750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156108cd578060051b830135858112156108c5578301610120813603126108c5576040519461047386613cf6565b61047c82613e86565b8652602082013567ffffffffffffffff81116108c95782019436601f870112156108c9578535956104ac8761438a565b966104ba6040519889613d12565b80885260208089019160051b830101903682116108c55760208301905b828210610892575050505060208701958652604083013567ffffffffffffffff811161088e5761050a9036908501613f3e565b916040880192835261053461052236606087016147e7565b9460608a0195865260c03691016147e7565b9560808901968752835151156108665761055867ffffffffffffffff8a5116615657565b1561082f5767ffffffffffffffff895116825260076020526040822061057f86518261509d565b61058d88516002830161509d565b6004855191019080519067ffffffffffffffff8211610802576105b08354614635565b601f81116107c7575b50602090601f831160011461072857610607929186918361071d575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b88518051821015610641579061063b6001926106348367ffffffffffffffff8f5116926145f2565b5190614cce565b0161060c565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561070f67ffffffffffffffff60019796949851169251935191516106db6106a660405196879687526101006020880152610100870190613d8d565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610442565b015190508e806105d5565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106107af5750908460019594939210610778575b505050811b01905561060a565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d808061076b565b92936020600181928786015181550195019301610755565b6107f29084875260208720601f850160051c810191602086106107f8575b601f0160051c0190614883565b8d6105b9565b90915081906107e5565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116108c1576020916108b68392833691890101613f3e565b8152019101906104d7565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff6108f36108ee8486889a9699979a61476c565b6142e7565b16916108fe8361538d565b15610a4057828452600760205261091a6005604086200161532a565b94845b865181101561095357600190858752600760205261094c60056040892001610945838b6145f2565b5190615523565b500161091d565b509396929094509490948087526007602052600560408820888155886001820155886002820155886003820155886004820161098f8154614635565b806109ff575b50505001805490888155816109e1575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190919493929461040f565b885260208820908101905b818110156109a5578881556001016109ec565b601f8111600114610a155750555b888a80610995565b81835260208320610a3091601f01861c810190600101614883565b8082528160208120915555610a0d565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102da5760206003193601126102da5760043567ffffffffffffffff81116108c957610a9e903690600401613ffc565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580610c4e575b610c2257825b818110610ad1578380f35b610adc81838561477c565b67ffffffffffffffff610aee826142e7565b1690610b07826000526006602052604060002054151590565b15610bf657907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610bb6610b90602060019897018b610b488261478c565b15610bbd578790526003602052610b6f60408d20610b6936604088016147e7565b9061509d565b868c526004602052610b8b60408d20610b693660a088016147e7565b61478c565b916040519215158352610ba9602084016040830161483f565b60a060808401910161483f565ba201610ac6565b60026040828a610b8b94526007602052610bdf828220610b6936858c016147e7565b8a815260076020522001610b693660a088016147e7565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610ac0565b50346102da57806003193601126102da57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102da5760406003193601126102da5760043567ffffffffffffffff81116108c957610cf3903690600401613ffc565b60243567ffffffffffffffff81116102d657610d13903690600401613fcb565b919092610d1e614c83565b845b828110610d8a57505050825b818110610d37578380f35b8067ffffffffffffffff610d516108ee600194868861476c565b16808652600a6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610d2c565b610d986108ee82858561477c565b610da382858561477c565b90602082019060e0830190610db78261478c565b156110dd5760a0840161271061ffff610dcf83614799565b1610156110ce5760c085019161271061ffff610dea85614799565b1610156110965763ffffffff610dff866147a8565b16156110615767ffffffffffffffff1694858c52600a60205260408c20610e25866147a8565b63ffffffff16908054906040840191610e3d836147a8565b60201b67ffffffff0000000016936060860194610e59866147a8565b60401b6bffffffff0000000000000000169660800196610e78886147a8565b60601b6fffffffff0000000000000000000000001691610e978a614799565b60801b71ffff000000000000000000000000000000001693610eb88c614799565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610f6b8761478c565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610fbc906147b9565b63ffffffff168752610fcd906147b9565b63ffffffff166020870152610fe1906147b9565b63ffffffff166040860152610ff5906147b9565b63ffffffff16606085015261100990613eca565b61ffff16608084015261101b90613eca565b61ffff1660a083015261102d90613e9b565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610d20565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff6110a586614799565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff6110a5602493614799565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346102da5760806003193601126102da5761112c613dec565b50611135613e6f565b61113d613eb9565b5060643567ffffffffffffffff811161088e579167ffffffffffffffff60409261116d60e0953690600401613ed9565b50508260c0855161117d81613cda565b82815282602082015282878201528260608201528260808201528260a08201520152168152600a60205220604051906111b582613cda565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346102da5760406003193601126102da5760043567ffffffffffffffff81116108c95761127e903690600401613fcb565b6024359073ffffffffffffffffffffffffffffffffffffffff8216908183036108c5576112a9614c83565b845b8181106112b6578580f35b80602073ffffffffffffffffffffffffffffffffffffffff6112e36112de602495878b61476c565b6142c6565b16604051938480927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa801561158d5784869185948a91611551575b508061133c575b50505060019150016112ab565b8861143473ffffffffffffffffffffffffffffffffffffffff6113636112de888a8661476c565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff98909816602482015260448082018790528152918e91166113c5606484613d12565b81806040998a51956113d78c88613d12565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15611548573d61141881613d53565b906114258b519283613d12565b8152809360203d92013e6159c8565b805180611499575b50507f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9173ffffffffffffffffffffffffffffffffffffffff6114876112de8860019a60209661476c565b169451908152a382915084843861132f565b8160209350839294969798506114b3955001019101614bd5565b156114c557908484939288388061143c565b608482517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606092506159c8565b93945050505060203d8111611586575b61156b8183613d12565b602082600092810103126102da575090848484935138611328565b503d611561565b6040513d89823e3d90fd5b50346102da57806003193601126102da57604051906005548083528260208101600584526020842092845b8181106116aa5750506115d892500383613d12565b81516115fc6115e68261438a565b916115f46040519384613d12565b80835261438a565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b845181101561165b578067ffffffffffffffff611648600193886145f2565b511661165482866145f2565b5201611629565b50925090604051928392602084019060208552518091526040840192915b818110611687575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611679565b84548352600194850194879450602090930192016115c3565b50346102da5760206003193601126102da576116fc6116e86116e3613e58565b61474a565b604051918291602083526020830190613d8d565b0390f35b50346102da5760606003193601126102da5760043567ffffffffffffffff81116108c9578060040160a0600319833603011261088e5761173e613ea8565b9260443567ffffffffffffffff81116108c957611762611772913690600401613ed9565b61176a6145d9565b503691613f07565b926084810190611781826142c6565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611e7457602481019477ffffffffffffffff000000000000000000000000000000006117e7876142e7565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611de7578591611e45575b50611e1d5767ffffffffffffffff61187b876142e7565b16611893816000526006602052604060002054151590565b15611df257602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611de7578590611d9a575b73ffffffffffffffffffffffffffffffffffffffff9150163303611d6e5760648201359161ffff8816918215611cbd5761ffff60025460a01c168015611c9557808410611c65575067ffffffffffffffff61194e896142e7565b1680875260036020527f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932885806119c060408b2073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161570c565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283611b43575b611b3989611ad56116e3611a3b8e8d614fca565b92611a4584614f38565b611a4e816142e7565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a26142e7565b9060405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152611b11604082613d12565b60405192611b1e84613cbe565b83526020830152604051928392604084526040840190613fa1565b9060208301520390f35b833b156108c1578787959493928a8793604051998a98899788967f5c3af7ca000000000000000000000000000000000000000000000000000000008852600488016060905280611b92916152da565b6064890160a09052610104890190611ba9926143c3565b94611bb390613e86565b67ffffffffffffffff166084880152604401611bce90613e37565b73ffffffffffffffffffffffffffffffffffffffff1660a487015260c4860152611bf790613e37565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526024840152828103600319016044840152611c2c91613d8d565b03925af18015611c5a57611c45575b8080808080611a27565b611c50828092613d12565b6102da5780611c3b565b6040513d84823e3d90fd5b86604491857f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b6004877f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b67ffffffffffffffff611ccf896142e7565b1680875260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448580611d4160408b2073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161570c565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26119e9565b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611ddf575b81611db460209383613d12565b810103126108c557611dda73ffffffffffffffffffffffffffffffffffffffff916143a2565b6118f4565b3d9150611da7565b6040513d87823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611e67915060203d602011611e6d575b611e5f8183613d12565b810190614bd5565b38611864565b503d611e55565b60248373ffffffffffffffffffffffffffffffffffffffff611e95856142c6565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102da5767ffffffffffffffff611ed636613f5c565b929091611ee1614c83565b1691611efa836000526006602052604060002054151590565b15610a40578284526007602052611f2960056040862001611f1c368486613f07565b6020815191012090615523565b15611f6e57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f686040519283926020845260208401916143c3565b0390a280f35b82611fb2836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916143c3565b0390fd5b50346102da5760206003193601126102da5767ffffffffffffffff611fd9613e58565b1681526007602052611ff06005604083200161532a565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061203561201f8361438a565b9261202d6040519485613d12565b80845261438a565b01835b81811061210c575050825b82518110156120895780612059600192856145f2565b518552600860205261206d60408620614688565b61207782856145f2565b5261208281846145f2565b5001612043565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106120c157505050500390f35b919360206120fc827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613d8d565b96019201920185949391926120b2565b806060602080938601015201612038565b50346102da5760206003193601126102da576004359067ffffffffffffffff82116102da57816004019160a060031982360301126108c95761215d6145d9565b506020926040519261216f8585613d12565b8084526084830161217f816142c6565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036126e057602484019477ffffffffffffffff000000000000000000000000000000006121e5876142e7565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152878160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156126655784916126c3575b5061269b5767ffffffffffffffff612278876142e7565b16612290816000526006602052604060002054151590565b15612670578773ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561266557849061261d575b73ffffffffffffffffffffffffffffffffffffffff91501633036125f157606485013594859467ffffffffffffffff612328896142e7565b1680865260078a526123766040872073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016998a9161570c565b6040805173ffffffffffffffffffffffffffffffffffffffff8a168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692836124d6575b896124786116e38b8b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108c61243081614f38565b67ffffffffffffffff612442856142e7565b6040805173ffffffffffffffffffffffffffffffffffffffff909616865233602087015285019290925216918060608101611acd565b906040517ffa7c07de00000000000000000000000000000000000000000000000000000000828201528181526124af604082613d12565b604051926124bc84613cbe565b8352818301526116fc604051928284938452830190613fa1565b833b156125ed5791858094928a9694604051978896879586947f5c3af7ca000000000000000000000000000000000000000000000000000000008652600486016060905280612524916152da565b6064870160a0905261010487019061253b926143c3565b9261254590613e86565b67ffffffffffffffff16608486015260440161256090613e37565b73ffffffffffffffffffffffffffffffffffffffff1660a48501528b60c485015261258a90613e37565b73ffffffffffffffffffffffffffffffffffffffff1660e48401528360248401528281036003190160448401526125c091613d8d565b03925af18015611c5a576125d8575b808080806123fc565b6125e3828092613d12565b6102da57806125cf565b8580fd5b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d831161265e575b6126338183613d12565b810103126102d65761265973ffffffffffffffffffffffffffffffffffffffff916143a2565b6122f0565b503d612629565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6126da9150883d8a11611e6d57611e5f8183613d12565b38612261565b9073ffffffffffffffffffffffffffffffffffffffff611e956024936142c6565b50346102da57806003193601126102da57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102da5760c06003193601126102da5761274f613dec565b612757613e6f565b9060643561ffff811681036102d65760843567ffffffffffffffff81116108c557612786903690600401613ed9565b9160a4359360028510156108c1576127a19560443591614402565b90604051918291602083016020845282518091526020604085019301915b8181106127cd575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016127bf565b50346102da5760206003193601126102da57602061283867ffffffffffffffff612824613e58565b166000526006602052604060002054151590565b6040519015158152f35b50346102da57806003193601126102da57805473ffffffffffffffffffffffffffffffffffffffff811633036128e1577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102da57806003193601126102da576002546009546040805173ffffffffffffffffffffffffffffffffffffffff808516825260a09490941c61ffff1660208201529290911690820152606090f35b50346102da5761296936613f5c565b61297593929193614c83565b67ffffffffffffffff8216612997816000526006602052604060002054151590565b156129b657506129b392936129ad913691613f07565b90614cce565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102da5760406003193601126102da576129fb613e58565b906024359067ffffffffffffffff82116102da57602061283884612a223660048701613f3e565b9061434d565b50346102da5760406003193601126102da5760043567ffffffffffffffff81116108c957806004019161010060031983360301126102da57612a68613ea8565b9181604051612a7681613c73565b5260c48101926064820135612aa6612aa0612a9b612a94888a6142fc565b3691613f07565b614bed565b82614ad4565b946084840191612ab5836142c6565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361325e57602485019777ffffffffffffffff00000000000000000000000000000000612b1b8a6142e7565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156131e157889161323f575b506132175767ffffffffffffffff612baf8a6142e7565b16612bc7816000526006602052604060002054151590565b156131ec57602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156131e15788916131c2575b501561319657612c3e896142e7565b94612c5460a4880196612a22612a9489866142fc565b1561314f5761ffff1690878a8a841561309c575067ffffffffffffffff9150612c7c906142e7565b1680895260046020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a80612cee60408d2073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161570c565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169384612ebd575b5050505050505060440192612d69846142c6565b9173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b156108c9576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101859052818180604481015b038183885af18015611c5a57612ea8575b505067ffffffffffffffff602094612e8f85612e51612e4b7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966142e7565b936142c6565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b0390a280604051612e9f81613c73565b52604051908152f35b612eb3828092613d12565b6102da5780612e0c565b843b15613098578895949392869289928d6040519a8b998a9889977f5eff3bf70000000000000000000000000000000000000000000000000000000089526004890160609052612f0d87806152da565b60648b0161010090526101648b0190612f25926143c3565b94612f2f90613e86565b67ffffffffffffffff1660848a0152604401612f4a90613e37565b73ffffffffffffffffffffffffffffffffffffffff1660a489015260c4880152612f7390613e37565b73ffffffffffffffffffffffffffffffffffffffff1660e4870152612f9890846152da565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c87840301610104880152612fcd92916143c3565b90612fd890836152da565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161012487015261300d92916143c3565b9060e48b0161301b916152da565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8584030161014486015261305092916143c3565b908b6024840152604483015203925af1801561308d57908391613078575b8080808080612d55565b8161308291613d12565b6108c957813861306e565b6040513d85823e3d90fd5b8880fd5b806131226002604067ffffffffffffffff6130d77f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c976142e7565b16968781526007602052200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161570c565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612d17565b61315986836142fc565b611fb26040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916143c3565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b6131db915060203d602011611e6d57611e5f8183613d12565b38612c2f565b6040513d8a823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b613258915060203d602011611e6d57611e5f8183613d12565b38612b98565b60248673ffffffffffffffffffffffffffffffffffffffff611e95866142c6565b50346102da5760206003193601126102da576004359067ffffffffffffffff82116102da57816004019161010060031982360301126108c957816040516132c581613c73565b5260648101356132d4816149cc565b9160848101916132e3836142c6565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361388457602482019277ffffffffffffffff00000000000000000000000000000000613349856142e7565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561158d578791613865575b5061383d5767ffffffffffffffff6133dd856142e7565b166133f5816000526006602052604060002054151590565b1561381257602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561158d5787916137f3575b50156137c75761346c846142e7565b9061348260a4850192612a22612a94858c6142fc565b156137bd57908697879287985067ffffffffffffffff6134a1886142e7565b1680855260076020526134f36002604087200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169a8b9161570c565b6040805173ffffffffffffffffffffffffffffffffffffffff8b168152602081018c90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691826135e9575b5050505050506044019361358c856142c6565b833b156108c9576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024810185905281818060448101612dfb565b823b156108c55787958591604051978896879586947f5eff3bf7000000000000000000000000000000000000000000000000000000008652600486016060905261363385806152da565b60648801610100905261016488019061364b926143c3565b9361365590613e86565b67ffffffffffffffff16608487015261367060448e01613e37565b73ffffffffffffffffffffffffffffffffffffffff1660a487015260c486015261369990613e37565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526136be90836152da565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101048601526136f392916143c3565b61370060c48b01836152da565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8584030161012486015261373592916143c3565b9060e48a01613743916152da565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8484030161014485015261377892916143c3565b8b602483015282604483015203925af180156137b25761379d575b8581808080613579565b946137ab8160449397613d12565b9490613793565b6040513d88823e3d90fd5b61315982896142fc565b6024867f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b61380c915060203d602011611e6d57611e5f8183613d12565b3861345d565b7fa9902c7e000000000000000000000000000000000000000000000000000000008752600452602486fd5b6004867f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b61387e915060203d602011611e6d57611e5f8183613d12565b386133c6565b60248573ffffffffffffffffffffffffffffffffffffffff611e95866142c6565b50346102da5760c06003193601126102da576138bf613dec565b506138c8613e6f565b6138d0613e14565b506084359161ffff831683036102da5760a4359067ffffffffffffffff82116102da5760a063ffffffff8061ffff61391788886139103660048b01613ed9565b5050614141565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346102da57806003193601126102da57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102da5760406003193601126102da57613996613e58565b6024359182151583036102da57610140613a596139b385856140be565b613a0960409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346102da5760206003193601126102da57602090613a78613dec565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102da57806003193601126102da57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102da57806003193601126102da57506116fc604051613b33604082613d12565b601b81527f4275726e4d696e74546f6b656e506f6f6c20312e372e302d64657600000000006020820152604051918291602083526020830190613d8d565b9050346108c95760206003193601126108c9576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361088e57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115613c49575b8115613c1f575b8115613bf5575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613bee565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613be7565b7f331710310000000000000000000000000000000000000000000000000000000081149150613be0565b6020810190811067ffffffffffffffff821117613c8f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117613c8f57604052565b60e0810190811067ffffffffffffffff821117613c8f57604052565b60a0810190811067ffffffffffffffff821117613c8f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613c8f57604052565b67ffffffffffffffff8111613c8f57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110613dd75750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613d98565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203613e0f57565b600080fd5b6064359073ffffffffffffffffffffffffffffffffffffffff82168203613e0f57565b359073ffffffffffffffffffffffffffffffffffffffff82168203613e0f57565b6004359067ffffffffffffffff82168203613e0f57565b6024359067ffffffffffffffff82168203613e0f57565b359067ffffffffffffffff82168203613e0f57565b35908115158203613e0f57565b6024359061ffff82168203613e0f57565b6044359061ffff82168203613e0f57565b359061ffff82168203613e0f57565b9181601f84011215613e0f5782359167ffffffffffffffff8311613e0f5760208381860195010111613e0f57565b929192613f1382613d53565b91613f216040519384613d12565b829481845281830111613e0f578281602093846000960137010152565b9080601f83011215613e0f57816020613f5993359101613f07565b90565b906040600319830112613e0f5760043567ffffffffffffffff81168103613e0f57916024359067ffffffffffffffff8211613e0f57613f9d91600401613ed9565b9091565b613f59916020613fba8351604084526040840190613d8d565b920151906020818403910152613d8d565b9181601f84011215613e0f5782359167ffffffffffffffff8311613e0f576020808501948460051b010111613e0f57565b9181601f84011215613e0f5782359167ffffffffffffffff8311613e0f576020808501948460081b010111613e0f57565b6040519061403a82613cf6565b60006080838281528260208201528260408201528260608201520152565b9060405161406581613cf6565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916140d061402d565b506140d961402d565b5061410d57166000526007602052604060002090613f59614101600261410661410186614058565b6148e9565b9401614058565b16908160005260036020526141286141016040600020614058565b916000526004602052613f596141016040600020614058565b9061ffff8060025460a01c16911692831515928380946142be575b6142945767ffffffffffffffff16600052600a6020526040600020916040519261418584613cda565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a01526142795761421a575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b81939750809294501061424957505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b50821561415c565b3573ffffffffffffffffffffffffffffffffffffffff81168103613e0f5790565b3567ffffffffffffffff81168103613e0f5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613e0f570180359067ffffffffffffffff8211613e0f57602001918136038313613e0f57565b9067ffffffffffffffff613f5992166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613c8f5760051b60200190565b519073ffffffffffffffffffffffffffffffffffffffff82168203613e0f57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169586156145b7576144bb9467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c48601916143c3565b916002821015614588578380600094819460a483015203915afa90811561457c576000916144e7575090565b3d8083833e6144f68183613d12565b81019060208183031261088e5780519067ffffffffffffffff82116102d6570181601f8201121561088e5780519061452d8261438a565b9361453b6040519586613d12565b82855260208086019360051b8301019384116102da5750602001905b8282106145645750505090565b60208091614571846143a2565b815201910190614557565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b50505050505050506040516145cd602082613d12565b60008152600036813790565b604051906145e682613cbe565b60606020838281520152565b80518210156146065760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c9216801561467e575b602083101461464f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614644565b906040519182600082549261469c84614635565b808452936001811690811561470a57506001146146c3575b506146c192500383613d12565b565b90506000929192526020600020906000915b8183106146ee5750509060206146c192820101386146b4565b60209193508060019154838589010152019101909184926146d5565b602093506146c19592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386146b4565b67ffffffffffffffff166000526007602052613f596004604060002001614688565b91908110156146065760051b0190565b91908110156146065760081b0190565b358015158103613e0f5790565b3561ffff81168103613e0f5790565b3563ffffffff81168103613e0f5790565b359063ffffffff82168203613e0f57565b35906fffffffffffffffffffffffffffffffff82168203613e0f57565b9190826060910312613e0f576040516060810181811067ffffffffffffffff821117613c8f57604052604061483a81839561482181613e9b565b855261482f602082016147ca565b6020860152016147ca565b910152565b6fffffffffffffffffffffffffffffffff61487d6040809361486081613e9b565b1515865283614871602083016147ca565b166020870152016147ca565b16910152565b81811061488e575050565b60008155600101614883565b818102929181159184041417156148ad57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b919082039182116148ad57565b6148f161402d565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff808351169161494e602085019361494861493b63ffffffff875116426148dc565b856080890151169061489a565b906152cd565b8082101561496757505b16825263ffffffff4216905290565b9050614958565b9060ff8091169116039060ff82116148ad57565b60ff16604d81116148ad57600a0a90565b811561499d570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f000000000000000000000000000000000000000000000000000000000000000060ff81169081600614614acf5781600611614aa4576006614a0d9161496e565b90604d60ff8316118015614a6b575b614a34575090614a2e613f5992614982565b9061489a565b90507fa9cb113d00000000000000000000000000000000000000000000000000000000600052600660045260245260445260646000fd5b50614a7582614982565b801561499d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311614a1c565b614aaf90600661496e565b90604d60ff831611614a34575090614ac9613f5992614982565b90614993565b505090565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614bce57828411614baa5790614b199161496e565b91604d60ff8416118015614b71575b614b3b57505090614a2e613f5992614982565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614b7b83614982565b801561499d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411614b28565b614bb39161496e565b91604d60ff841611614b3b57505090614ac9613f5992614982565b5050505090565b90816020910312613e0f57518015158103613e0f5790565b80518015614c5d57602003614c1f578051602082810191830183900312613e0f57519060ff8211614c1f575060ff1690565b611fb2906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613d8d565b50507f000000000000000000000000000000000000000000000000000000000000000090565b73ffffffffffffffffffffffffffffffffffffffff600154163303614ca457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115614f0e5767ffffffffffffffff81516020830120921691826000526007602052614d038160056040600020016156b7565b15614eca5760005260086020526040600020815167ffffffffffffffff8111613c8f57614d308254614635565b601f8111614e98575b506020601f8211600114614dd25791614dac827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614dc295600091614dc7575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613d8d565b0390a2565b905084015138614d7b565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110614e80575092614dc29492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614e49575b5050811b0190556116e8565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880614e3d565b9192602060018192868a015181550194019201614e02565b614ec490836000526020600020601f840160051c810191602085106107f857601f0160051c0190614883565b38614d39565b5090611fb26040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613d8d565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15613e0f57604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af18015611c5a57614fbd575050565b81614fc791613d12565b50565b9061ffff9067ffffffffffffffff614fe4602085016142e7565b16600052600a6020526040600020826040519161500083613cda565b549263ffffffff8416835263ffffffff8460201c16602084015263ffffffff8460401c16604084015263ffffffff8460601c166060840152818460801c169283608082015260c060ff848760901c16968760a085015260a01c16151591015216151560001461509657505b16801561508e576127106150876060613f59940135928361489a565b04906148dc565b506060013590565b905061506b565b81519192911561521f576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff602085015116106151bc576146c191925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b60648361521d604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604084015116158015906152ae575b61524d576146c191926150e0565b60648361521d604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff602084015116151561523f565b919082018092116148ad57565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215613e0f57016020813591019167ffffffffffffffff8211613e0f578136038313613e0f57565b906040519182815491828252602082019060005260206000209260005b81811061535c5750506146c192500383613d12565b8454835260019485019487945060209093019201615347565b80548210156146065760005260206000200190600090565b600081815260066020526040902054801561551c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116148ad57600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116148ad578181036154ad575b505050600554801561547e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161543b816005615375565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6155046154be6154cf936005615375565b90549060031b1c9283926005615375565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526006602052604060002055388080615402565b5050600090565b906001820191816000528260205260406000205480151560001461564e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116148ad578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116148ad57818103615617575b5050508054801561547e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906155d88282615375565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6156376156276154cf9386615375565b90549060031b1c92839286615375565b9055600052836020526040600020553880806155a0565b50505050600090565b806000526006602052604060002054156000146156b15760055468010000000000000000811015613c8f576156986154cf8260018594016005556005615375565b9055600554906000526006602052604060002055600190565b50600090565b600082815260018201602052604090205461551c5780549068010000000000000000821015613c8f57826156f56154cf846001809601855584615375565b905580549260005201602052604060002055600190565b9182549060ff8260a01c161580156159c0575b6159ba576fffffffffffffffffffffffffffffffff8216916001850190815461576463ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426148dc565b908161591c575b50508481106158d057508383106157c557505061579a6fffffffffffffffffffffffffffffffff9283926148dc565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c92831561586457816157dd916148dc565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908082116148ad5761582b6158309273ffffffffffffffffffffffffffffffffffffffff966152cd565b614993565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82869293961161599057615937926149489160801c9061489a565b8084101561598b5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617865592388061576b565b615942565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561571f565b91929015615a4357508151156159dc575090565b3b156159e55790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615a565750805190602001fd5b611fb2906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613d8d56fea164736f6c634300081a000a",
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
	outstruct.MinBlockConfirmations = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setDynamicConfig", router, minBlockConfirmations, rateLimitAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetDynamicConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
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

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.WithdrawFeeTokens(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, feeTokens, recipient)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.WithdrawFeeTokens(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, feeTokens, recipient)
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
	Router                common.Address
	MinBlockConfirmations uint16
	RateLimitAdmin        common.Address
	Raw                   types.Log
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
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
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
	Router                common.Address
	MinBlockConfirmations uint16
	RateLimitAdmin        common.Address
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
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
	return common.HexToHash("0xba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a117")
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
	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

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

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

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

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, error)

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
