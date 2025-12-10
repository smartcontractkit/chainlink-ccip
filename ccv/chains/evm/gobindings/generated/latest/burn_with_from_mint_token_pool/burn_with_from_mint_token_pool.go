// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_with_from_mint_token_pool

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

var BurnWithFromMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103485760a081615ec580380380916100208285610471565b8339810103126103485780516001600160a01b03811691908290036103485761004b60208201610494565b610057604083016104a2565b906100706080610069606086016104a2565b94016104a2565b92331561046057600180546001600160a01b031916331790558415801561044f575b801561043e575b61042d57608085905260c05260405163313ce56760e01b8152602081600481885afa600091816103f1575b506103c6575b5060a0526001600160a01b0390811660e052600280546001600160a01b03191692909116919091179055604051636eb1769f60e11b81523060048201819052602482015290602082604481845afa9182156103ba57600092610386575b50600019820180921161037057604051602081019263095ea7b360e01b8452306024830152604482015260448152610160606482610471565b6000806040948551936101738786610471565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610363573d906001600160401b03821161034d5784516101e49490926101d5601f8201601f191660200185610471565b83523d6000602085013e6104b6565b8051806102cd575b825161593e908161058782396080518181816117c9015281816119c701528181611a8301528181611d15015281816121b00152818161236701528181612b4001528181612d3901528181612df40152818161315b0152818161337f0152818161355101528181613a6e01528181613ac80152614e78015260a05181818161393401528181614a0701528181614a8a0152614f06015260c051818181610ca30152818161185f0152818161224a01528181612bdb015261341a015260e051818181611a2f015281816123ee01528181612da1015281816135d801526144440152f35b81602091810103126103485760200151801590811503610348576102f25738806101ec565b5162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916101e4926060916104b6565b634e487b7160e01b600052601160045260246000fd5b9091506020813d6020116103b2575b816103a260209383610471565b8101031261034857519038610127565b3d9150610395565b6040513d6000823e3d90fd5b60ff1660ff82168181036103da57506100ca565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610425575b8161040d60209383610471565b810103126103485761041e90610494565b90386100c4565b3d9150610400565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610099565b506001600160a01b03841615610092565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761034d57604052565b519060ff8216820361034857565b51906001600160a01b038216820361034857565b9192901561051857508151156104ca575090565b3b156104d35790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561052b5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b83811061056e5750508160006044809484010152601f80199101168101030190fd5b6020828201810151604487840101528593500161054c56fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613b7357508063181f5a7714613aec57806321df0da714613a9b578063240028e814613a375780632422ac451461395857806324f65ee71461391a5780632c0634041461388157806339077537146132e0578063489a68f214612aa25780634c5ef0ed14612a5b5780634e921c30146129bc57806362ddd3c4146129355780637437ff9f146128f457806379ba50971461282d5780638926f54f146127e757806389720a62146127205780638da5cb5b146126ec5780639a4575b914612136578063a42a7b8b14611fcf578063acfecf9114611ed7578063b1c71c65146116dd578063b7946580146116a0578063c4bffe2b14611575578063c7230a6014611252578063d8aa3f4014611118578063dc04fa1f14610cc7578063dc0bd97114610c76578063dcbd41bc14610a72578063e8a1da17146103ae578063f2fde38b146102df578063fa41d79c146102ba5763ff8e03f31461018157600080fd5b346102b75760406003193601126102b75761019a613dee565b906101a3613e39565b6101ab614bac565b73ffffffffffffffffffffffffffffffffffffffff831692831561028f577f22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e616644797092937fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff000000000000000000000000000000000000000060095416176009556102896040519283928390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b0390a180f35b6004837f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102b757806003193601126102b757602061ffff60025460a01c16604051908152f35b50346102b75760206003193601126102b75773ffffffffffffffffffffffffffffffffffffffff61030e613dee565b610316614bac565b1633811461038657807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346102b75760406003193601126102b75760043567ffffffffffffffff81116108cb576103e0903690600401613ff0565b9060243567ffffffffffffffff8111610a6e579061040384923690600401613ff0565b93909161040e614bac565b83905b8282106108d35750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156108cf578060051b830135858112156108c7578301610120813603126108c7576040519461047586613cf8565b61047e82613eab565b8652602082013567ffffffffffffffff81116108cb5782019436601f870112156108cb578535956104ae876143af565b966104bc6040519889613d14565b80885260208089019160051b830101903682116108c75760208301905b828210610894575050505060208701958652604083013567ffffffffffffffff81116108905761050c9036908501613f63565b91604088019283526105366105243660608701614848565b9460608a0195865260c0369101614848565b9560808901968752835151156108685761055a67ffffffffffffffff8a51166154f4565b156108315767ffffffffffffffff8951168252600760205260408220610581865182614f3a565b61058f885160028301614f3a565b6004855191019080519067ffffffffffffffff8211610804576105b2835461465a565b601f81116107c9575b50602090601f831160011461072a57610609929186918361071f575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b88518051821015610643579061063d6001926106368367ffffffffffffffff8f511692614617565b5190614bf7565b0161060e565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561071167ffffffffffffffff60019796949851169251935191516106dd6106a860405196879687526101006020880152610100870190613d8f565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610444565b015190508e806105d7565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106107b1575090846001959493921061077a575b505050811b01905561060c565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d808061076d565b92936020600181928786015181550195019301610757565b6107f49084875260208720601f850160051c810191602086106107fa575b601f0160051c01906148e4565b8d6105bb565b90915081906107e7565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116108c3576020916108b88392833691890101613f63565b8152019101906104d9565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff6108f56108f08486889a9699979a6147cd565b61435d565b16916109008361522a565b15610a4257828452600760205261091c600560408620016151c7565b94845b865181101561095557600190858752600760205261094e60056040892001610947838b614617565b51906153c0565b500161091f565b5093969290945094909480875260076020526005604088208881558860018201558860028201558860038201558860048201610991815461465a565b80610a01575b50505001805490888155816109e3575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101909194939294610411565b885260208820908101905b818110156109a7578881556001016109ee565b601f8111600114610a175750555b888a80610997565b81835260208320610a3291601f01861c8101906001016148e4565b8082528160208120915555610a0f565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346102b75760206003193601126102b75760043567ffffffffffffffff81116108cb57610aa4903690600401614021565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580610c54575b610c2857825b818110610ad7578380f35b610ae28183856147dd565b67ffffffffffffffff610af48261435d565b1690610b0d826000526006602052604060002054151590565b15610bfc57907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610bbc610b96602060019897018b610b4e826147ed565b15610bc3578790526003602052610b7560408d20610b6f3660408801614848565b90614f3a565b868c526004602052610b9160408d20610b6f3660a08801614848565b6147ed565b916040519215158352610baf60208401604083016148a0565b60a06080840191016148a0565ba201610acc565b60026040828a610b9194526007602052610be5828220610b6f36858c01614848565b8a815260076020522001610b6f3660a08801614848565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610ac6565b50346102b757806003193601126102b757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102b75760406003193601126102b75760043567ffffffffffffffff81116108cb57610cf9903690600401614021565b60243567ffffffffffffffff8111610a6e57610d19903690600401613ff0565b919092610d24614bac565b845b828110610d9057505050825b818110610d3d578380f35b8067ffffffffffffffff610d576108f060019486886147cd565b16808652600a6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610d32565b610d9e6108f08285856147dd565b610da98285856147dd565b90602082019060e0830190610dbd826147ed565b156110e35760a0840161271061ffff610dd5836147fa565b1610156110d45760c085019161271061ffff610df0856147fa565b16101561109c5763ffffffff610e0586614809565b16156110675767ffffffffffffffff1694858c52600a60205260408c20610e2b86614809565b63ffffffff16908054906040840191610e4383614809565b60201b67ffffffff0000000016936060860194610e5f86614809565b60401b6bffffffff0000000000000000169660800196610e7e88614809565b60601b6fffffffff0000000000000000000000001691610e9d8a6147fa565b60801b71ffff000000000000000000000000000000001693610ebe8c6147fa565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610f71876147ed565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610fc29061481a565b63ffffffff168752610fd39061481a565b63ffffffff166020870152610fe79061481a565b63ffffffff166040860152610ffb9061481a565b63ffffffff16606085015261100f90613eef565b61ffff16608084015261102190613eef565b61ffff1660a083015261103390613ec0565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610d26565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff6110ab866147fa565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff6110ab6024936147fa565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346102b75760806003193601126102b757611132613dee565b5061113b613e94565b611143613ede565b5060643567ffffffffffffffff8111610890579167ffffffffffffffff60409261117360e0953690600401613efe565b50508260c0855161118381613cdc565b82815282602082015282878201528260608201528260808201528260a08201520152168152600a60205220604051906111bb82613cdc565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346102b75760406003193601126102b75760043567ffffffffffffffff81116108cb57611284903690600401613ff0565b9061128d613e39565b90611296614bac565b835b8381106112a3578480f35b6024602073ffffffffffffffffffffffffffffffffffffffff6112cf6112ca8589886147cd565b61433c565b16604051928380927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa90811561156a578691611537575b5080611320575b50600101611298565b73ffffffffffffffffffffffffffffffffffffffff6113436112ca8488876147cd565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff88166024830152604480830186905282529261141392166113a5606483613d14565b89806040958651946113b78887613d14565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1561152f573d916113f883613d55565b9261140587519485613d14565b83523d8c602085013e615865565b80518061148e575b50509073ffffffffffffffffffffffffffffffffffffffff83926114456112ca6001968a896147cd565b905192835216907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602073ffffffffffffffffffffffffffffffffffffffff881692a390611317565b906020806114a0938301019101614b94565b156114ac57388061141b565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615865565b905060203d8111611563575b61154d8183613d14565b602082600092810103126102b757505138611310565b503d611543565b6040513d88823e3d90fd5b50346102b757806003193601126102b757604051906005548083528260208101600584526020842092845b8181106116875750506115b592500383613d14565b81516115d96115c3826143af565b916115d16040519384613d14565b8083526143af565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015611638578067ffffffffffffffff61162560019388614617565b51166116318286614617565b5201611606565b50925090604051928392602084019060208552518091526040840192915b818110611664575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611656565b84548352600194850194879450602090930192016115a0565b50346102b75760206003193601126102b7576116d96116c56116c0613e7d565b6147ab565b604051918291602083526020830190613d8f565b0390f35b50346102b75760606003193601126102b75760043567ffffffffffffffff81116108cb57806004019160a060031983360301126102b75761171c613ecd565b9060443567ffffffffffffffff81116108cb5761173d903690600401613efe565b946117466145fe565b50602485019467ffffffffffffffff61175e8761435d565b168452600a6020526117a261ffff604086209616936127108515159788600014611eba5761ffff611798915460901c1660648601356148fb565b04985b3691613f2c565b9060848101926117b18461433c565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611e705777ffffffffffffffff000000000000000000000000000000006118128961435d565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611de3578791611e41575b50611e195767ffffffffffffffff6118a68961435d565b166118be816000526006602052604060002054151590565b15611dee57602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611de3578790611d96575b73ffffffffffffffffffffffffffffffffffffffff9150163303611d6a5760648201359661194d8a8961476f565b9015611cb95761ffff60025460a01c168015611c9157808710611c6157507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff61199e8b61435d565b16918289526003602052806119ef60408b2073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916155a9565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283611b3f575b611b3589611b046116c0611a6a8e8d61476f565b92611a7484614e61565b611a7d8161435d565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a261435d565b90611b0d614eff565b60405192611b1a84613cc0565b83526020830152604051928392604084526040840190613fc6565b9060208301520390f35b833b156108c357869493929185918a604051988997889687957f5c3af7ca000000000000000000000000000000000000000000000000000000008752600487016060905280611b8d91615177565b6064880160a09052610104880190611ba4926143e8565b93611bae90613eab565b67ffffffffffffffff166084870152604401611bc990613e5c565b73ffffffffffffffffffffffffffffffffffffffff1660a48601528c60c4860152611bf390613e5c565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526024840152828103600319016044840152611c2891613d8f565b03925af18015611c5657611c41575b8080808080611a56565b611c4c828092613d14565b6102b75780611c37565b6040513d84823e3d90fd5b87604491887f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b6004887f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611cec8b61435d565b1691828952600760205280611d3d60408b2073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916155a9565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611a18565b6024867f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611ddb575b81611db060209383613d14565b810103126108c357611dd673ffffffffffffffffffffffffffffffffffffffff916143c7565b61191f565b3d9150611da3565b6040513d89823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008752600452602486fd5b6004867f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611e63915060203d602011611e69575b611e5b8183613d14565b810190614b94565b3861188f565b503d611e51565b60248673ffffffffffffffffffffffffffffffffffffffff611e918761433c565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b61ffff611ed0915460801c1660648601356148fb565b049861179b565b50346102b75767ffffffffffffffff611eef36613f81565b929091611efa614bac565b1691611f13836000526006602052604060002054151590565b15610a42578284526007602052611f4260056040862001611f35368486613f2c565b60208151910120906153c0565b15611f8757907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f816040519283926020845260208401916143e8565b0390a280f35b82611fcb836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916143e8565b0390fd5b50346102b75760206003193601126102b75767ffffffffffffffff611ff2613e7d565b1681526007602052612009600560408320016151c7565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061204e612038836143af565b926120466040519485613d14565b8084526143af565b01835b818110612125575050825b82518110156120a2578061207260019285614617565b5185526008602052612086604086206146ad565b6120908285614617565b5261209b8184614617565b500161205c565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106120da57505050500390f35b91936020612115827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613d8f565b96019201920185949391926120cb565b806060602080938601015201612051565b50346102b75760206003193601126102b7576004359067ffffffffffffffff82116102b757816004019160a060031982360301126108cb576121766145fe565b50602092604051926121888585613d14565b808452608483016121988161433c565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036126cb57602484019477ffffffffffffffff000000000000000000000000000000006121fe8761435d565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152878160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156126505784916126ae575b506126865767ffffffffffffffff6122918761435d565b166122a9816000526006602052604060002054151590565b1561265b578773ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015612650578490612608575b73ffffffffffffffffffffffffffffffffffffffff91501633036125dc57606485013594859467ffffffffffffffff6123418961435d565b1680865260078a5261238f6040872073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016998a916155a9565b6040805173ffffffffffffffffffffffffffffffffffffffff8a168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692836124c1575b896124916116c08b8b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108c61244981614e61565b67ffffffffffffffff61245b8561435d565b6040805173ffffffffffffffffffffffffffffffffffffffff909616865233602087015285019290925216918060608101611afc565b9061249a614eff565b604051926124a784613cc0565b8352818301526116d9604051928284938452830190613fc6565b833b156125d85791858094928a9694604051978896879586947f5c3af7ca00000000000000000000000000000000000000000000000000000000865260048601606090528061250f91615177565b6064870160a09052610104870190612526926143e8565b9261253090613eab565b67ffffffffffffffff16608486015260440161254b90613e5c565b73ffffffffffffffffffffffffffffffffffffffff1660a48501528b60c485015261257590613e5c565b73ffffffffffffffffffffffffffffffffffffffff1660e48401528360248401528281036003190160448401526125ab91613d8f565b03925af18015611c56576125c3575b80808080612415565b6125ce828092613d14565b6102b757806125ba565b8580fd5b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d8311612649575b61261e8183613d14565b81010312610a6e5761264473ffffffffffffffffffffffffffffffffffffffff916143c7565b612309565b503d612614565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6126c59150883d8a11611e6957611e5b8183613d14565b3861227a565b9073ffffffffffffffffffffffffffffffffffffffff611e9160249361433c565b50346102b757806003193601126102b757602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102b75760c06003193601126102b75761273a613dee565b612742613e94565b9060643561ffff81168103610a6e5760843567ffffffffffffffff81116108c757612771903690600401613efe565b9160a4359360028510156108c35761278c9560443591614427565b90604051918291602083016020845282518091526020604085019301915b8181106127b8575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff168452859450602093840193909201916001016127aa565b50346102b75760206003193601126102b757602061282367ffffffffffffffff61280f613e7d565b166000526006602052604060002054151590565b6040519015158152f35b50346102b757806003193601126102b757805473ffffffffffffffffffffffffffffffffffffffff811633036128cc577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102b757806003193601126102b7576002546009546040805173ffffffffffffffffffffffffffffffffffffffff938416815292909116602083015290f35b50346102b75761294436613f81565b61295093929193614bac565b67ffffffffffffffff8216612972816000526006602052604060002054151590565b15612991575061298e9293612988913691613f2c565b90614bf7565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102b75760206003193601126102b75760043561ffff811690818103610890577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb91602091612a0b614bac565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006002549260a01b16911617600255604051908152a180f35b50346102b75760406003193601126102b757612a75613e7d565b906024359067ffffffffffffffff82116102b757602061282384612a9c3660048701613f63565b90614372565b50346102b75760406003193601126102b75760043567ffffffffffffffff81116108cb57806004019161010060031983360301126102b757612ae2613ecd565b9181604051612af081613c75565b5260c48101926064820135612b19612b13612b0e61179b888a6142eb565b614993565b82614a87565b946084840191612b288361433c565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036132bf57602485019777ffffffffffffffff00000000000000000000000000000000612b8e8a61435d565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156132425788916132a0575b506132785767ffffffffffffffff612c228a61435d565b16612c3a816000526006602052604060002054151590565b1561324d57602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115613242578891613223575b50156131f757612cb18961435d565b94612cc760a4880196612a9c61179b89866142eb565b156131b05761ffff1690878a8a84156130fd575067ffffffffffffffff9150612cef9061435d565b1680895260046020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a80612d6160408d2073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916155a9565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169384612f1e575b5050505050505060440192612ddc8461433c565b9173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b156108cb576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101859052818180604481015b038183885af18015611c5657612f09575b5050608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff612ed7612ed17ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09661435d565b9261433c565b60405196875233898801521660408601528560608601521692a260405190612efe82613c75565b815260405190518152f35b612f14828092613d14565b6102b75780612e7f565b843b156130f9578895949392869289928d6040519a8b998a9889977f5eff3bf70000000000000000000000000000000000000000000000000000000089526004890160609052612f6e8780615177565b60648b0161010090526101648b0190612f86926143e8565b94612f9090613eab565b67ffffffffffffffff1660848a0152604401612fab90613e5c565b73ffffffffffffffffffffffffffffffffffffffff1660a489015260c4880152612fd490613e5c565b73ffffffffffffffffffffffffffffffffffffffff1660e4870152612ff99084615177565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8784030161010488015261302e92916143e8565b906130399083615177565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161012487015261306e92916143e8565b9060e48b0161307c91615177565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101448601526130b192916143e8565b908b6024840152604483015203925af180156130ee579083916130d9575b8080808080612dc8565b816130e391613d14565b6108cb5781386130cf565b6040513d85823e3d90fd5b8880fd5b806131836002604067ffffffffffffffff6131387f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9761435d565b16968781526007602052200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916155a9565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612d8a565b6131ba86836142eb565b611fcb6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916143e8565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b61323c915060203d602011611e6957611e5b8183613d14565b38612ca2565b6040513d8a823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6132b9915060203d602011611e6957611e5b8183613d14565b38612c0b565b60248673ffffffffffffffffffffffffffffffffffffffff611e918661433c565b50346102b75760206003193601126102b7576004359067ffffffffffffffff82116102b757816004019161010060031982360301126108cb578160405161332681613c75565b528160405161333481613c75565b52606481013560c4820191613358613352612b0e61179b86896142eb565b83614a87565b9260848201926133678461433c565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611e7057602483019377ffffffffffffffff000000000000000000000000000000006133cd8661435d565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115613242578891613862575b506132785767ffffffffffffffff6134618661435d565b16613479816000526006602052604060002054151590565b1561324d57602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115613242578891613843575b50156131f7576134f08561435d565b9261350660a4860194612a9c61179b878d6142eb565b15613839579187989391889388995067ffffffffffffffff6135278961435d565b1680865260076020526135796002604088200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169b8c916155a9565b6040805173ffffffffffffffffffffffffffffffffffffffff8c168152602081018d90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283613670575b50505050505050604401936136138561433c565b833b156108cb576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024810185905281818060448101612e6e565b833b156125d85788968692604051988997889687957f5eff3bf700000000000000000000000000000000000000000000000000000000875260048701606090528d6136bb8780615177565b60648a0161010090526101648a01906136d3926143e8565b946136dd90613eab565b67ffffffffffffffff1660848901526044016136f890613e5c565b73ffffffffffffffffffffffffffffffffffffffff1660a488015260c487015261372190613e5c565b73ffffffffffffffffffffffffffffffffffffffff1660e48601526137469084615177565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161010487015261377b92916143e8565b906137869083615177565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101248601526137bb92916143e8565b9060e48a016137c991615177565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c848403016101448501526137fe92916143e8565b8b602483015282604483015203925af1801561156a57613824575b8581808080806135ff565b946138328160449397613d14565b9490613819565b6131ba848a6142eb565b61385c915060203d602011611e6957611e5b8183613d14565b386134e1565b61387b915060203d602011611e6957611e5b8183613d14565b3861344a565b50346102b75760c06003193601126102b75761389b613dee565b506138a4613e94565b6138ac613e16565b506084359161ffff831683036102b75760a4359067ffffffffffffffff82116102b75760a063ffffffff8061ffff6138f388886138ec3660048b01613efe565b5050614166565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346102b757806003193601126102b757602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102b75760406003193601126102b757613972613e7d565b6024359182151583036102b757610140613a3561398f85856140e3565b6139e560409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346102b75760206003193601126102b757602090613a54613dee565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102b757806003193601126102b757602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102b757806003193601126102b757506116d9604051613b0f606082613d14565b602381527f4275726e5769746846726f6d4d696e74546f6b656e506f6f6c20312e372e302d60208201527f64657600000000000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190613d8f565b9050346108cb5760206003193601126108cb576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361089057602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115613c4b575b8115613c21575b8115613bf7575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613bf0565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613be9565b7f331710310000000000000000000000000000000000000000000000000000000081149150613be2565b6020810190811067ffffffffffffffff821117613c9157604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117613c9157604052565b60e0810190811067ffffffffffffffff821117613c9157604052565b60a0810190811067ffffffffffffffff821117613c9157604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613c9157604052565b67ffffffffffffffff8111613c9157601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110613dd95750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613d9a565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203613e1157565b600080fd5b6064359073ffffffffffffffffffffffffffffffffffffffff82168203613e1157565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203613e1157565b359073ffffffffffffffffffffffffffffffffffffffff82168203613e1157565b6004359067ffffffffffffffff82168203613e1157565b6024359067ffffffffffffffff82168203613e1157565b359067ffffffffffffffff82168203613e1157565b35908115158203613e1157565b6024359061ffff82168203613e1157565b6044359061ffff82168203613e1157565b359061ffff82168203613e1157565b9181601f84011215613e115782359167ffffffffffffffff8311613e115760208381860195010111613e1157565b929192613f3882613d55565b91613f466040519384613d14565b829481845281830111613e11578281602093846000960137010152565b9080601f83011215613e1157816020613f7e93359101613f2c565b90565b906040600319830112613e115760043567ffffffffffffffff81168103613e1157916024359067ffffffffffffffff8211613e1157613fc291600401613efe565b9091565b613f7e916020613fdf8351604084526040840190613d8f565b920151906020818403910152613d8f565b9181601f84011215613e115782359167ffffffffffffffff8311613e11576020808501948460051b010111613e1157565b9181601f84011215613e115782359167ffffffffffffffff8311613e11576020808501948460081b010111613e1157565b6040519061405f82613cf8565b60006080838281528260208201528260408201528260608201520152565b9060405161408a81613cf8565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916140f5614052565b506140fe614052565b5061413257166000526007602052604060002090613f7e614126600261412b6141268661407d565b61490e565b940161407d565b169081600052600360205261414d614126604060002061407d565b916000526004602052613f7e614126604060002061407d565b9061ffff8060025460a01c16911692831515928380946142e3575b6142b95767ffffffffffffffff16600052600a602052604060002091604051926141aa84613cdc565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a015261429e5761423f575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b81939750809294501061426e57505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b508215614181565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613e11570180359067ffffffffffffffff8211613e1157602001918136038313613e1157565b3573ffffffffffffffffffffffffffffffffffffffff81168103613e115790565b3567ffffffffffffffff81168103613e115790565b9067ffffffffffffffff613f7e92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613c915760051b60200190565b519073ffffffffffffffffffffffffffffffffffffffff82168203613e1157565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169586156145dc576144e09467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c48601916143e8565b9160028210156145ad578380600094819460a483015203915afa9081156145a15760009161450c575090565b3d8083833e61451b8183613d14565b8101906020818303126108905780519067ffffffffffffffff8211610a6e570181601f8201121561089057805190614552826143af565b936145606040519586613d14565b82855260208086019360051b8301019384116102b75750602001905b8282106145895750505090565b60208091614596846143c7565b81520191019061457c565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b50505050505050506040516145f2602082613d14565b60008152600036813790565b6040519061460b82613cc0565b60606020838281520152565b805182101561462b5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c921680156146a3575b602083101461467457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614669565b90604051918260008254926146c18461465a565b808452936001811690811561472f57506001146146e8575b506146e692500383613d14565b565b90506000929192526020600020906000915b8183106147135750509060206146e692820101386146d9565b60209193508060019154838589010152019101909184926146fa565b602093506146e69592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386146d9565b9190820391821161477c57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff166000526007602052613f7e60046040600020016146ad565b919081101561462b5760051b0190565b919081101561462b5760081b0190565b358015158103613e115790565b3561ffff81168103613e115790565b3563ffffffff81168103613e115790565b359063ffffffff82168203613e1157565b35906fffffffffffffffffffffffffffffffff82168203613e1157565b9190826060910312613e11576040516060810181811067ffffffffffffffff821117613c9157604052604061489b81839561488281613ec0565b85526148906020820161482b565b60208601520161482b565b910152565b6fffffffffffffffffffffffffffffffff6148de604080936148c181613ec0565b15158652836148d26020830161482b565b1660208701520161482b565b16910152565b8181106148ef575050565b600081556001016148e4565b8181029291811591840414171561477c57565b614916614052565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614973602085019361496d61496063ffffffff8751164261476f565b85608089015116906148fb565b9061516a565b8082101561498c57505b16825263ffffffff4216905290565b905061497d565b80518015614a03576020036149c5578051602082810191830183900312613e1157519060ff82116149c5575060ff1690565b611fcb906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613d8f565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161477c57565b60ff16604d811161477c57600a0a90565b8115614a58570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614b8d57828411614b635790614acc91614a29565b91604d60ff8416118015614b2a575b614af457505090614aee613f7e92614a3d565b906148fb565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614b3483614a3d565b8015614a58577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411614adb565b614b6c91614a29565b91604d60ff841611614af457505090614b87613f7e92614a3d565b90614a4e565b5050505090565b90816020910312613e1157518015158103613e115790565b73ffffffffffffffffffffffffffffffffffffffff600154163303614bcd57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115614e375767ffffffffffffffff81516020830120921691826000526007602052614c2c816005604060002001615554565b15614df35760005260086020526040600020815167ffffffffffffffff8111613c9157614c59825461465a565b601f8111614dc1575b506020601f8211600114614cfb5791614cd5827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614ceb95600091614cf0575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613d8f565b0390a2565b905084015138614ca4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110614da9575092614ceb9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614d72575b5050811b0190556116c5565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880614d66565b9192602060018192868a015181550194019201614d2b565b614ded90836000526020600020601f840160051c810191602085106107fa57601f0160051c01906148e4565b38614c62565b5090611fcb6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613d8f565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15613e11576040517f9dc29fac000000000000000000000000000000000000000000000000000000008152306004820152602481019190915260009182908290604490829084905af18015611c5657614ef2575050565b81614efc91613d14565b50565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613f7e604082613d14565b8151919291156150bc576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff60208501511610615059576146e691925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836150ba604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040840151161580159061514b575b6150ea576146e69192614f7d565b6064836150ba604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208401511615156150dc565b9190820180921161477c57565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215613e1157016020813591019167ffffffffffffffff8211613e11578136038313613e1157565b906040519182815491828252602082019060005260206000209260005b8181106151f95750506146e692500383613d14565b84548352600194850194879450602090930192016151e4565b805482101561462b5760005260206000200190600090565b60008181526006602052604090205480156153b9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161477c57600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161477c5781810361534a575b505050600554801561531b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016152d8816005615212565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6153a161535b61536c936005615212565b90549060031b1c9283926005615212565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600660205260406000205538808061529f565b5050600090565b90600182019181600052826020526040600020548015156000146154eb577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161477c578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161477c578181036154b4575b5050508054801561531b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906154758282615212565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6154d46154c461536c9386615212565b90549060031b1c92839286615212565b90556000528360205260406000205538808061543d565b50505050600090565b8060005260066020526040600020541560001461554e5760055468010000000000000000811015613c915761553561536c8260018594016005556005615212565b9055600554906000526006602052604060002055600190565b50600090565b60008281526001820160205260409020546153b95780549068010000000000000000821015613c91578261559261536c846001809601855584615212565b905580549260005201602052604060002055600190565b9182549060ff8260a01c1615801561585d575b615857576fffffffffffffffffffffffffffffffff8216916001850190815461560163ffffffff6fffffffffffffffffffffffffffffffff83169360801c164261476f565b90816157b9575b505084811061576d57508383106156625750506156376fffffffffffffffffffffffffffffffff92839261476f565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615701578161567a9161476f565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019080821161477c576156c86156cd9273ffffffffffffffffffffffffffffffffffffffff9661516a565b614a4e565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82869293961161582d576157d49261496d9160801c906148fb565b808410156158285750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615608565b6157df565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156155bc565b919290156158e05750815115615879575090565b3b156158825790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156158f35750805190602001fd5b611fcb906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613d8f56fea164736f6c634300081a000a",
}

var BurnWithFromMintTokenPoolABI = BurnWithFromMintTokenPoolMetaData.ABI

var BurnWithFromMintTokenPoolBin = BurnWithFromMintTokenPoolMetaData.Bin

func DeployBurnWithFromMintTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnWithFromMintTokenPool, error) {
	parsed, err := BurnWithFromMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnWithFromMintTokenPoolBin), backend, token, localTokenDecimals, advancedPoolHooks, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnWithFromMintTokenPool{address: address, abi: *parsed, BurnWithFromMintTokenPoolCaller: BurnWithFromMintTokenPoolCaller{contract: contract}, BurnWithFromMintTokenPoolTransactor: BurnWithFromMintTokenPoolTransactor{contract: contract}, BurnWithFromMintTokenPoolFilterer: BurnWithFromMintTokenPoolFilterer{contract: contract}}, nil
}

type BurnWithFromMintTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnWithFromMintTokenPoolCaller
	BurnWithFromMintTokenPoolTransactor
	BurnWithFromMintTokenPoolFilterer
}

type BurnWithFromMintTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnWithFromMintTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnWithFromMintTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnWithFromMintTokenPoolSession struct {
	Contract     *BurnWithFromMintTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnWithFromMintTokenPoolCallerSession struct {
	Contract *BurnWithFromMintTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnWithFromMintTokenPoolTransactorSession struct {
	Contract     *BurnWithFromMintTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnWithFromMintTokenPoolRaw struct {
	Contract *BurnWithFromMintTokenPool
}

type BurnWithFromMintTokenPoolCallerRaw struct {
	Contract *BurnWithFromMintTokenPoolCaller
}

type BurnWithFromMintTokenPoolTransactorRaw struct {
	Contract *BurnWithFromMintTokenPoolTransactor
}

func NewBurnWithFromMintTokenPool(address common.Address, backend bind.ContractBackend) (*BurnWithFromMintTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnWithFromMintTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnWithFromMintTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPool{address: address, abi: abi, BurnWithFromMintTokenPoolCaller: BurnWithFromMintTokenPoolCaller{contract: contract}, BurnWithFromMintTokenPoolTransactor: BurnWithFromMintTokenPoolTransactor{contract: contract}, BurnWithFromMintTokenPoolFilterer: BurnWithFromMintTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnWithFromMintTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnWithFromMintTokenPoolCaller, error) {
	contract, err := bindBurnWithFromMintTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolCaller{contract: contract}, nil
}

func NewBurnWithFromMintTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnWithFromMintTokenPoolTransactor, error) {
	contract, err := bindBurnWithFromMintTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolTransactor{contract: contract}, nil
}

func NewBurnWithFromMintTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnWithFromMintTokenPoolFilterer, error) {
	contract, err := bindBurnWithFromMintTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolFilterer{contract: contract}, nil
}

func bindBurnWithFromMintTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnWithFromMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnWithFromMintTokenPool.Contract.BurnWithFromMintTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.BurnWithFromMintTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.BurnWithFromMintTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnWithFromMintTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _BurnWithFromMintTokenPool.Contract.GetCurrentRateLimiterState(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _BurnWithFromMintTokenPool.Contract.GetCurrentRateLimiterState(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnWithFromMintTokenPool.Contract.GetDynamicConfig(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnWithFromMintTokenPool.Contract.GetDynamicConfig(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnWithFromMintTokenPool.Contract.GetFee(&_BurnWithFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnWithFromMintTokenPool.Contract.GetFee(&_BurnWithFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _BurnWithFromMintTokenPool.Contract.GetMinBlockConfirmation(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _BurnWithFromMintTokenPool.Contract.GetMinBlockConfirmation(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnWithFromMintTokenPool.Contract.GetRemotePools(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnWithFromMintTokenPool.Contract.GetRemotePools(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnWithFromMintTokenPool.Contract.GetRemoteToken(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnWithFromMintTokenPool.Contract.GetRemoteToken(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _BurnWithFromMintTokenPool.Contract.GetRequiredCCVs(&_BurnWithFromMintTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _BurnWithFromMintTokenPool.Contract.GetRequiredCCVs(&_BurnWithFromMintTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnWithFromMintTokenPool.Contract.GetRmnProxy(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnWithFromMintTokenPool.Contract.GetRmnProxy(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnWithFromMintTokenPool.Contract.GetSupportedChains(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnWithFromMintTokenPool.Contract.GetSupportedChains(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnWithFromMintTokenPool.Contract.GetToken(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnWithFromMintTokenPool.Contract.GetToken(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _BurnWithFromMintTokenPool.Contract.GetTokenDecimals(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnWithFromMintTokenPool.Contract.GetTokenDecimals(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnWithFromMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnWithFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnWithFromMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnWithFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnWithFromMintTokenPool.Contract.IsRemotePool(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnWithFromMintTokenPool.Contract.IsRemotePool(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnWithFromMintTokenPool.Contract.IsSupportedChain(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnWithFromMintTokenPool.Contract.IsSupportedChain(&_BurnWithFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnWithFromMintTokenPool.Contract.IsSupportedToken(&_BurnWithFromMintTokenPool.CallOpts, token)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnWithFromMintTokenPool.Contract.IsSupportedToken(&_BurnWithFromMintTokenPool.CallOpts, token)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) Owner() (common.Address, error) {
	return _BurnWithFromMintTokenPool.Contract.Owner(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnWithFromMintTokenPool.Contract.Owner(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnWithFromMintTokenPool.Contract.SupportsInterface(&_BurnWithFromMintTokenPool.CallOpts, interfaceId)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnWithFromMintTokenPool.Contract.SupportsInterface(&_BurnWithFromMintTokenPool.CallOpts, interfaceId)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnWithFromMintTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnWithFromMintTokenPool.Contract.TypeAndVersion(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnWithFromMintTokenPool.Contract.TypeAndVersion(&_BurnWithFromMintTokenPool.CallOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.AcceptOwnership(&_BurnWithFromMintTokenPool.TransactOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.AcceptOwnership(&_BurnWithFromMintTokenPool.TransactOpts)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.AddRemotePool(&_BurnWithFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.AddRemotePool(&_BurnWithFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.ApplyChainUpdates(&_BurnWithFromMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.ApplyChainUpdates(&_BurnWithFromMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnWithFromMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnWithFromMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.LockOrBurn(&_BurnWithFromMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.LockOrBurn(&_BurnWithFromMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.LockOrBurn0(&_BurnWithFromMintTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.LockOrBurn0(&_BurnWithFromMintTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.ReleaseOrMint(&_BurnWithFromMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.ReleaseOrMint(&_BurnWithFromMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.ReleaseOrMint0(&_BurnWithFromMintTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.ReleaseOrMint0(&_BurnWithFromMintTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.RemoveRemotePool(&_BurnWithFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.RemoveRemotePool(&_BurnWithFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.SetDynamicConfig(&_BurnWithFromMintTokenPool.TransactOpts, router, rateLimitAdmin)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.SetDynamicConfig(&_BurnWithFromMintTokenPool.TransactOpts, router, rateLimitAdmin)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "setMinBlockConfirmation", minBlockConfirmation)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.SetMinBlockConfirmation(&_BurnWithFromMintTokenPool.TransactOpts, minBlockConfirmation)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.SetMinBlockConfirmation(&_BurnWithFromMintTokenPool.TransactOpts, minBlockConfirmation)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.SetRateLimitConfig(&_BurnWithFromMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.SetRateLimitConfig(&_BurnWithFromMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.TransferOwnership(&_BurnWithFromMintTokenPool.TransactOpts, to)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.TransferOwnership(&_BurnWithFromMintTokenPool.TransactOpts, to)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.WithdrawFeeTokens(&_BurnWithFromMintTokenPool.TransactOpts, feeTokens, recipient)
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnWithFromMintTokenPool.Contract.WithdrawFeeTokens(&_BurnWithFromMintTokenPool.TransactOpts, feeTokens, recipient)
}

type BurnWithFromMintTokenPoolChainAddedIterator struct {
	Event *BurnWithFromMintTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolChainAdded)
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
		it.Event = new(BurnWithFromMintTokenPoolChainAdded)
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

func (it *BurnWithFromMintTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnWithFromMintTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolChainAddedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolChainAdded)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnWithFromMintTokenPoolChainAdded, error) {
	event := new(BurnWithFromMintTokenPoolChainAdded)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolChainRemovedIterator struct {
	Event *BurnWithFromMintTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolChainRemoved)
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
		it.Event = new(BurnWithFromMintTokenPoolChainRemoved)
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

func (it *BurnWithFromMintTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnWithFromMintTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolChainRemovedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolChainRemoved)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnWithFromMintTokenPoolChainRemoved, error) {
	event := new(BurnWithFromMintTokenPoolChainRemoved)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolDynamicConfigSetIterator struct {
	Event *BurnWithFromMintTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolDynamicConfigSet)
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
		it.Event = new(BurnWithFromMintTokenPoolDynamicConfigSet)
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

func (it *BurnWithFromMintTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolDynamicConfigSet struct {
	Router         common.Address
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnWithFromMintTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolDynamicConfigSetIterator{contract: _BurnWithFromMintTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolDynamicConfigSet)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*BurnWithFromMintTokenPoolDynamicConfigSet, error) {
	event := new(BurnWithFromMintTokenPoolDynamicConfigSet)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolFeeTokenWithdrawnIterator struct {
	Event *BurnWithFromMintTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(BurnWithFromMintTokenPoolFeeTokenWithdrawn)
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

func (it *BurnWithFromMintTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolFeeTokenWithdrawn struct {
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*BurnWithFromMintTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolFeeTokenWithdrawnIterator{contract: _BurnWithFromMintTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolFeeTokenWithdrawn)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*BurnWithFromMintTokenPoolFeeTokenWithdrawn, error) {
	event := new(BurnWithFromMintTokenPoolFeeTokenWithdrawn)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolInboundRateLimitConsumedIterator struct {
	Event *BurnWithFromMintTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(BurnWithFromMintTokenPoolInboundRateLimitConsumed)
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

func (it *BurnWithFromMintTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolInboundRateLimitConsumedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolInboundRateLimitConsumed)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnWithFromMintTokenPoolInboundRateLimitConsumed, error) {
	event := new(BurnWithFromMintTokenPoolInboundRateLimitConsumed)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolLockedOrBurnedIterator struct {
	Event *BurnWithFromMintTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolLockedOrBurned)
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
		it.Event = new(BurnWithFromMintTokenPoolLockedOrBurned)
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

func (it *BurnWithFromMintTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolLockedOrBurnedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolLockedOrBurned)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*BurnWithFromMintTokenPoolLockedOrBurned, error) {
	event := new(BurnWithFromMintTokenPoolLockedOrBurned)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolMinBlockConfirmationSetIterator struct {
	Event *BurnWithFromMintTokenPoolMinBlockConfirmationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolMinBlockConfirmationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolMinBlockConfirmationSet)
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
		it.Event = new(BurnWithFromMintTokenPoolMinBlockConfirmationSet)
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

func (it *BurnWithFromMintTokenPoolMinBlockConfirmationSetIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolMinBlockConfirmationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolMinBlockConfirmationSet struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*BurnWithFromMintTokenPoolMinBlockConfirmationSetIterator, error) {

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolMinBlockConfirmationSetIterator{contract: _BurnWithFromMintTokenPool.contract, event: "MinBlockConfirmationSet", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolMinBlockConfirmationSet) (event.Subscription, error) {

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolMinBlockConfirmationSet)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseMinBlockConfirmationSet(log types.Log) (*BurnWithFromMintTokenPoolMinBlockConfirmationSet, error) {
	event := new(BurnWithFromMintTokenPoolMinBlockConfirmationSet)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *BurnWithFromMintTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(BurnWithFromMintTokenPoolOutboundRateLimitConsumed)
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

func (it *BurnWithFromMintTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolOutboundRateLimitConsumedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolOutboundRateLimitConsumed)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnWithFromMintTokenPoolOutboundRateLimitConsumed, error) {
	event := new(BurnWithFromMintTokenPoolOutboundRateLimitConsumed)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnWithFromMintTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolOwnershipTransferRequested)
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
		it.Event = new(BurnWithFromMintTokenPoolOwnershipTransferRequested)
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

func (it *BurnWithFromMintTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnWithFromMintTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolOwnershipTransferRequestedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolOwnershipTransferRequested)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnWithFromMintTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnWithFromMintTokenPoolOwnershipTransferRequested)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolOwnershipTransferredIterator struct {
	Event *BurnWithFromMintTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolOwnershipTransferred)
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
		it.Event = new(BurnWithFromMintTokenPoolOwnershipTransferred)
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

func (it *BurnWithFromMintTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnWithFromMintTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolOwnershipTransferredIterator{contract: _BurnWithFromMintTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolOwnershipTransferred)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnWithFromMintTokenPoolOwnershipTransferred, error) {
	event := new(BurnWithFromMintTokenPoolOwnershipTransferred)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolRateLimitConfiguredIterator struct {
	Event *BurnWithFromMintTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolRateLimitConfigured)
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
		it.Event = new(BurnWithFromMintTokenPoolRateLimitConfigured)
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

func (it *BurnWithFromMintTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolRateLimitConfiguredIterator{contract: _BurnWithFromMintTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolRateLimitConfigured)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*BurnWithFromMintTokenPoolRateLimitConfigured, error) {
	event := new(BurnWithFromMintTokenPoolRateLimitConfigured)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolReleasedOrMintedIterator struct {
	Event *BurnWithFromMintTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolReleasedOrMinted)
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
		it.Event = new(BurnWithFromMintTokenPoolReleasedOrMinted)
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

func (it *BurnWithFromMintTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolReleasedOrMintedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolReleasedOrMinted)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*BurnWithFromMintTokenPoolReleasedOrMinted, error) {
	event := new(BurnWithFromMintTokenPoolReleasedOrMinted)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolRemotePoolAddedIterator struct {
	Event *BurnWithFromMintTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolRemotePoolAdded)
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
		it.Event = new(BurnWithFromMintTokenPoolRemotePoolAdded)
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

func (it *BurnWithFromMintTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolRemotePoolAddedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolRemotePoolAdded)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*BurnWithFromMintTokenPoolRemotePoolAdded, error) {
	event := new(BurnWithFromMintTokenPoolRemotePoolAdded)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolRemotePoolRemovedIterator struct {
	Event *BurnWithFromMintTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolRemotePoolRemoved)
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
		it.Event = new(BurnWithFromMintTokenPoolRemotePoolRemoved)
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

func (it *BurnWithFromMintTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolRemotePoolRemovedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolRemotePoolRemoved)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*BurnWithFromMintTokenPoolRemotePoolRemoved, error) {
	event := new(BurnWithFromMintTokenPoolRemotePoolRemoved)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *BurnWithFromMintTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnWithFromMintTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _BurnWithFromMintTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnWithFromMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated)
				if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated)
	if err := _BurnWithFromMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (BurnWithFromMintTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnWithFromMintTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (BurnWithFromMintTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e6166447970")
}

func (BurnWithFromMintTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (BurnWithFromMintTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnWithFromMintTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnWithFromMintTokenPoolMinBlockConfirmationSet) Topic() common.Hash {
	return common.HexToHash("0xa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb")
}

func (BurnWithFromMintTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnWithFromMintTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnWithFromMintTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnWithFromMintTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (BurnWithFromMintTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnWithFromMintTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnWithFromMintTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_BurnWithFromMintTokenPool *BurnWithFromMintTokenPool) Address() common.Address {
	return _BurnWithFromMintTokenPool.address
}

type BurnWithFromMintTokenPoolInterface interface {
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

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnWithFromMintTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnWithFromMintTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnWithFromMintTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnWithFromMintTokenPoolChainRemoved, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnWithFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnWithFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnWithFromMintTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnWithFromMintTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*BurnWithFromMintTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*BurnWithFromMintTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnWithFromMintTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnWithFromMintTokenPoolLockedOrBurned, error)

	FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*BurnWithFromMintTokenPoolMinBlockConfirmationSetIterator, error)

	WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolMinBlockConfirmationSet) (event.Subscription, error)

	ParseMinBlockConfirmationSet(log types.Log) (*BurnWithFromMintTokenPoolMinBlockConfirmationSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnWithFromMintTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnWithFromMintTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnWithFromMintTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnWithFromMintTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnWithFromMintTokenPoolOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*BurnWithFromMintTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnWithFromMintTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnWithFromMintTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnWithFromMintTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnWithFromMintTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnWithFromMintTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnWithFromMintTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnWithFromMintTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
