// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cctp_through_ccv_token_pool

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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
}

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

var CCTPThroughCCVTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCTPVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"IPoolV1NotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461039057614ede803803809161001d82856103f9565b833981019060c0818303126103905780516001600160a01b038116908181036103905761004c6020840161041c565b906100596040850161042a565b906100666060860161042a565b936100736080870161042a565b60a087015190966001600160401b03821161039057019680601f89011215610390578751976001600160401b03891161031b578860051b9060208201996100bd6040519b8c6103f9565b8a526020808b019282010192831161039057602001905b8282106103e15750505033156103d057600180546001600160a01b03191633179055801580156103bf575b80156103ae575b61039d5760049260209260805260c0526040519283809263313ce56760e01b82525afa6000918161035c575b50610331575b5060a052600380546001600160a01b0319908116909155600280549091166001600160a01b039290921691909117905560405160209061017882826103f9565b60008152600036813760408051949085016001600160401b0381118682101761031b576040528452808285015260005b815181101561020f576001906001600160a01b036101c6828561043e565b5116846101d282610480565b6101df575b5050016101a8565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138846101d7565b5050915160005b8151811015610287576001600160a01b03610231828461043e565b5116908115610276577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef858361026860019561057e565b50604051908152a101610216565b6342bcdf7f60e11b60005260046000fd5b6001600160a01b03831660e0526040516148ff90816105df82396080518181816115d401528181611754015281816117f501528181611997015281816126c70152818161286c015281816128e50152818161298901528181612cf60152612d50015260a05181612b68015260c051818181610bc3015281816116700152612763015260e05181818161116f015261234e0152f35b634e487b7160e01b600052604160045260246000fd5b60ff1660ff82168181036103455750610138565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610395575b81610378602093836103f9565b81010312610390576103899061041c565b9038610132565b600080fd5b3d915061036b565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610106565b506001600160a01b038516156100ff565b639b15e16f60e01b60005260046000fd5b602080916103ee8461042a565b8152019101906100d4565b601f909101601f19168101906001600160401b0382119082101761031b57604052565b519060ff8216820361039057565b51906001600160a01b038216820361039057565b80518210156104525760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104525760005260206000200190600090565b6000818152600e6020526040902054801561057757600019810181811161056157600d5460001981019190821161056157808203610510575b505050600d5480156104fa57600019016104d481600d610468565b8154906000199060031b1b19169055600d55600052600e60205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61054961052161053293600d610468565b90549060031b1c928392600d610468565b819391549060031b91821b91600019901b19161790565b9055600052600e6020526040600020553880806104b9565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600e602052604060002054156000146105d857600d546801000000000000000081101561031b576105bf610532826001859401600d55600d610468565b9055600d5490600052600e602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a714612dfa57508063181f5a7714612d7457806321df0da714612d23578063240028e814612cc35780632422ac4514612be65780632451a62714612b8c57806324f65ee714612b4e5780632c06340414612ab657806337a3210d14612a825780633907753714612a17578063489a68f21461264d5780634c5ef0ed146126085780634e921c301461256b5780635cb80c5d14612372578063615521a71461232157806362ddd3c41461229e5780637437ff9f1461225057806379ba5097146121855780638926f54f1461214057806389720a62146120c65780638da5cb5b1461209257806391a2749a14611ee65780639a4575b914611e82578063a42a7b8b14611d14578063acfecf9114611c1d578063ae39a25714611a91578063b1c71c6514611531578063b7946580146114f5578063bfeffd3f1461144b578063c4bffe2b1461131b578063d8aa3f4014611009578063dc04fa1f14610be7578063dc0bd97114610b96578063dcbd41bc14610987578063e8a1da17146102b2578063f2fde38b146101e05763fa41d79c146101b657600080fd5b346101db5760006003193601126101db57602061ffff60025460a01c16604051908152f35b600080fd5b346101db5760206003193601126101db5773ffffffffffffffffffffffffffffffffffffffff61020e613055565b610216613bf2565b1633811461028857807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101db5760406003193601126101db5760043567ffffffffffffffff81116101db576102e3903690600401613208565b9060243567ffffffffffffffff81116101db57610304903690600401613208565b92909161030f613bf2565b6000905b8282106107de5750505060007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301905b838110156107dc576000908060051b840135838112156107d8578401610120813603126107d8576040519461037a86612f99565b813567ffffffffffffffff811681036107d4578652602082013567ffffffffffffffff81116107d45782019436601f870112156107d4578535956103bd8761327e565b966103cb6040519889612fb5565b80885260208089019160051b830101903682116107d05760208301905b82821061079d575050505060208701958652604083013567ffffffffffffffff81116107995761041b90369085016131ea565b91604088019283526104456104333660608701613a13565b9460608a0195865260c0369101613a13565b9560808901968752835151156107715761046967ffffffffffffffff8a51166144fd565b1561073a5767ffffffffffffffff8951168152600860205260408120610490865182613ee0565b61049e885160028301613ee0565b6004855191019080519067ffffffffffffffff821161070d576104c18354613861565b601f81116106d2575b50602090601f8311600114610633576105189291859183610628575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b87518051821015610550579061054a6001926105438367ffffffffffffffff8e51169261381e565b5190613c3d565b0161051b565b505097967f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293929196509461061c67ffffffffffffffff60019751169251935191516105e86105b360405196879687526101006020880152610100870190612ff6565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101929192610346565b015190508e806104e6565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b8181106106ba5750908460019594939210610683575b505050811b01905561051b565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080610676565b92936020600181928786015181550195019301610660565b6106fd9084865260208620601f850160051c81019160208610610703575b601f0160051c0190613a9e565b8d6104ca565b90915081906106f0565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8a51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8580fd5b813567ffffffffffffffff81116107cc576020916107c183928336918901016131ea565b8152019101906103e8565b8980fd5b8780fd5b8480fd5b8280fd5b005b909267ffffffffffffffff6107ff6107fa8686869997996139e6565b613581565b169261080a846142ab565b1561095957836000526008602052610828600560406000200161411d565b9260005b84518110156108645760019086600052600860205261085d6005604060002001610856838961381e565b51906143b7565b500161082c565b50939094919592508060005260086020526005604060002060008155600060018201556000600282015560006003820155600481016108a38154613861565b9081610916575b50500180549060008155816108f5575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019091939293610313565b6000526020600020908101905b818110156108ba5760008155600101610902565b81601f6000931160011461092e5750555b88806108aa565b8183526020832061094991601f01861c810190600101613a9e565b8082528160208120915555610927565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101db5760206003193601126101db5760043567ffffffffffffffff81116101db576109b8903690600401613311565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580610b74575b610b465760005b8181106109ea57005b6109f5818385613998565b67ffffffffffffffff610a0782613581565b1690610a20826000526007602052604060002054151590565b15610b1857907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610ad2610aac60206001989701610a60816139a8565b15610ad957866000526004602052610a896040600020610a833660408801613a13565b90613ee0565b866000526005602052610aa76040600020610a833660a08801613a13565b6139a8565b916040519215158352610ac56020840160408301613a5a565b60a0608084019101613a5a565ba2016109e1565b866000526008602052610af76040600020610a833660408801613a13565b866000526008602052610aa76002604060002001610a833660a08801613a13565b507f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b7f8e4a23d6000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b5073ffffffffffffffffffffffffffffffffffffffff600154163314156109da565b346101db5760006003193601126101db57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101db5760406003193601126101db5760043567ffffffffffffffff81116101db57610c18903690600401613311565b60243567ffffffffffffffff81116101db57610c38903690600401613208565b919092610c43613bf2565b60005b828110610cb25750505060005b818110610c5c57005b8067ffffffffffffffff610c766107fa60019486886139e6565b1680600052600b602052600060408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a201610c53565b610cc06107fa828585613998565b610ccb828585613998565b90602082019060e0830190610cdf826139a8565b15610f8b5760a0840161271061ffff610cf7836139b5565b161015610ffd5760c085019161271061ffff610d12856139b5565b161015610fc35763ffffffff610d27866139c4565b1615610f8b5767ffffffffffffffff169485600052600b6020526040600020610d4f866139c4565b63ffffffff16908054906040840191610d67836139c4565b60201b67ffffffff0000000016936060860194610d83866139c4565b60401b6bffffffff0000000000000000169660800196610da2886139c4565b60601b6fffffffff0000000000000000000000001691610dc18a6139b5565b60801b71ffff000000000000000000000000000000001693610de28c6139b5565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610e95876139a8565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610ee6906139d5565b63ffffffff168752610ef7906139d5565b63ffffffff166020870152610f0b906139d5565b63ffffffff166040860152610f1f906139d5565b63ffffffff166060850152610f3390613148565b61ffff166080840152610f4590613148565b61ffff1660a0830152610f57906130c9565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c46565b67ffffffffffffffff907f12332265000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b61ffff610fcf846139b5565b7f95f3517a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b610fcf61ffff916139b5565b346101db5760806003193601126101db57611022613055565b5061102b6130b2565b611033613137565b506064359067ffffffffffffffff82116101db5761105e67ffffffffffffffff923690600401613157565b5050600060c060405161107081612f61565b8281528260208201528260408201528260608201528260808201528260a082015201521680600052600b6020526040600020604051906110af82612f61565b5463ffffffff81168252602082019263ffffffff8260201c168452604083019363ffffffff8360401c1685526060840163ffffffff8460601c168152608085019161ffff8560801c16835260a086019361ffff8660901c16855260ff60c088019660a01c1615158652604051907f958021a70000000000000000000000000000000000000000000000000000000082526004820152604060248201526000604482015260208160648173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112bd576000916112c9575b50606073ffffffffffffffffffffffffffffffffffffffff916004604051809481937f7437ff9f000000000000000000000000000000000000000000000000000000008352165afa9081156112bd5760009161124a575b5061ffff9493919263ffffffff60e09981889687604083970151168952816040519c51168c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b906060823d6060116112b5575b8161126460609383612fb5565b810103126112b257604080519261127a84612f7d565b611283816135f4565b8452611291602082016135f4565b602085015201519061ffff821682036112b25750604082015261ffff6111f7565b80fd5b3d9150611257565b6040513d6000823e3d90fd5b90506020813d602011611313575b816112e460209383612fb5565b810103126101db57606061130c73ffffffffffffffffffffffffffffffffffffffff926135f4565b91506111a0565b3d91506112d7565b346101db5760006003193601126101db576040516006548082528160208101600660005260206000209260005b81811061143257505061135d92500382612fb5565b80519061138261136c8361327e565b9261137a6040519485612fb5565b80845261327e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208401920136833760005b81518110156113e2578067ffffffffffffffff6113cf6001938561381e565b51166113db828761381e565b52016113b0565b5050906040519182916020830190602084525180915260408301919060005b81811061140f575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611401565b8454835260019485019486945060209093019201611348565b346101db5760206003193601126101db5760043573ffffffffffffffffffffffffffffffffffffffff81168091036101db57611485613bf2565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a11617600355005b346101db5760206003193601126101db5761152d61151961151461309b565b613976565b604051918291602083526020830190612ff6565b0390f35b346101db5760606003193601126101db5760043567ffffffffffffffff81116101db5760a060031982360301126101db5761156a613126565b6044359067ffffffffffffffff82116101db5761158e6115ae923690600401613157565b9290611598613805565b506115a68386600401613e7d565b933691613185565b50608483016115bc81613596565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611a455750602483019277ffffffffffffffff0000000000000000000000000000000061162385613581565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112bd57600091611a16575b506119ec576116c9606461ffff926116c06116bb88613581565b614168565b01359384613af7565b9116801561192d5761ffff60025460a01c16908115611903578181106118d35750506115148361182b927f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff61172961189898613581565b16918260005260046020528061177c604060002073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161456d565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b6117ae81613581565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10606067ffffffffffffffff6040519373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001685523360208601528860408601521692a2613581565b906118c96040517f3047587c00000000000000000000000000000000000000000000000000000000602082015260048152611867602482612fb5565b6040519361187485612f45565b84526020840190815260405194859460408652516040808701526080860190612ff6565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0858303016060860152612ff6565b9060208301520390f35b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b506115148361182b927fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff61196c61189898613581565b1691826000526008602052806119bf604060002073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161456d565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26117a5565b7f53ad11d80000000000000000000000000000000000000000000000000000000060005260046000fd5b611a38915060203d602011611a3e575b611a308183612fb5565b810190613b89565b856116a1565b503d611a26565b611a6373ffffffffffffffffffffffffffffffffffffffff91613596565b7f961c9a4f000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346101db5760606003193601126101db57611aaa613055565b60243573ffffffffffffffffffffffffffffffffffffffff8116918282036101db5760443573ffffffffffffffffffffffffffffffffffffffff8116908181036101db57611af6613bf2565b73ffffffffffffffffffffffffffffffffffffffff8316918215611bf3577f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195611bee937fffffffffffffffffffffffff000000000000000000000000000000000000000060025416176002557fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c556040519384938491604091949373ffffffffffffffffffffffffffffffffffffffff809281606087019816865216602085015216910152565b0390a1005b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101db5767ffffffffffffffff611c3436613239565b929091611c3f613bf2565b1690611c58826000526007602052604060002054151590565b15610b1857816000526008602052611c896005604060002001611c7c368685613185565b60208151910120906143b7565b15611ccd577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611cc8604051928392602084526020840191613615565b0390a2005b611d10906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191613615565b0390fd5b346101db5760206003193601126101db5767ffffffffffffffff611d3661309b565b166000526008602052611d4f600560406000200161411d565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611d95611d7f8461327e565b93611d8d6040519586612fb5565b80855261327e565b0160005b818110611e7157505060005b8151811015611ded5780611dbb6001928461381e565b516000526009602052611dd160406000206138b4565b611ddb828661381e565b52611de6818561381e565b5001611da5565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b828210611e2657505050500390f35b91936020611e61827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851612ff6565b9601920192018594939192611e17565b806060602080938701015201611d99565b346101db5760206003193601126101db5760043567ffffffffffffffff81116101db5760031960a091360301126101db57611ebb613805565b507f690a7a400000000000000000000000000000000000000000000000000000000060005260046000fd5b346101db5760206003193601126101db5760043567ffffffffffffffff81116101db57604060031982360301126101db5760405190611f2482612f45565b806004013567ffffffffffffffff81116101db57611f489060043691840101613296565b825260248101359067ffffffffffffffff82116101db576004611f6e9236920101613296565b60208201908152611f7d613bf2565b519060005b8251811015611ff5578073ffffffffffffffffffffffffffffffffffffffff611fad6001938661381e565b5116611fb881614822565b611fc4575b5001611f82565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184611fbd565b505160005b81518110156107dc5773ffffffffffffffffffffffffffffffffffffffff612022828461381e565b5116908115612068577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef60208361205a6001956144be565b50604051908152a101611ffa565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101db5760006003193601126101db57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101db5760c06003193601126101db576120df613055565b6120e76130b2565b60643561ffff811681036101db5760843567ffffffffffffffff81116101db57612115903690600401613157565b9060a4359260028410156101db5761152d956121349560443591613654565b604051918291826130d6565b346101db5760206003193601126101db57602061217b67ffffffffffffffff61216761309b565b166000526007602052604060002054151590565b6040519015158152f35b346101db5760006003193601126101db5760005473ffffffffffffffffffffffffffffffffffffffff81163303612226577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101db5760006003193601126101db57600254600a54600c546040805173ffffffffffffffffffffffffffffffffffffffff94851681529284166020840152921691810191909152606090f35b346101db576122ac36613239565b6122b7929192613bf2565b67ffffffffffffffff82166122d9816000526007602052604060002054151590565b156122f457506107dc926122ee913691613185565b90613c3d565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b346101db5760006003193601126101db57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101db5760206003193601126101db5760043567ffffffffffffffff81116101db576123a3903690600401613208565b9073ffffffffffffffffffffffffffffffffffffffff600c54169182156120685760005b8181106123d057005b73ffffffffffffffffffffffffffffffffffffffff6123f86123f38385876139e6565b613596565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156112bd5760009161253a575b508061244e575b50506001016123c7565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff891660248401526044808401859052835291600091906124b0606482612fb5565b519082865af1156112bd576000513d6125315750813b155b6125035790857f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a39085612444565b507f5274afe70000000000000000000000000000000000000000000000000000000060005260045260246000fd5b600114156124c8565b906020823d8211612563575b8161255360209383612fb5565b810103126112b25750518661243d565b3d9150612546565b346101db5760206003193601126101db5760043561ffff8116908181036101db577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb916020916125b9613bf2565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006002549260a01b16911617600255604051908152a1005b346101db5760406003193601126101db5761262161309b565b60243567ffffffffffffffff81116101db5760209161264761217b9236906004016131ea565b906135b7565b346101db5760406003193601126101db5760043567ffffffffffffffff81116101db57806004019061010060031982360301126101db5761268c613126565b90600060405161269b81612efa565b52606481013591608482016126af81613596565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611a455750602482019377ffffffffffffffff0000000000000000000000000000000061271686613581565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156112bd576000916129f8575b506119ec576127a56116bb86613581565b6127ae85613581565b906127cb60a48501926126476127c48585613ba1565b3691613185565b156129b1575050608067ffffffffffffffff6128c860446128c160209861ffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc097161515600014612931578461282182613581565b168060005260058b527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a80612894604060002073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161456d565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2613581565b9501613596565b9373ffffffffffffffffffffffffffffffffffffffff60405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a28060405161292881612efa565b52604051908152f35b8461293b82613581565b168060005260088b527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8a80612894600260406000200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161456d565b6129bb9250613ba1565b611d106040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191613615565b612a11915060203d602011611a3e57611a308183612fb5565b86612794565b346101db5760206003193601126101db5760043567ffffffffffffffff81116101db5760031961010091360301126101db576000604051612a5781612efa565b527f690a7a400000000000000000000000000000000000000000000000000000000060005260046000fd5b346101db5760006003193601126101db57602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b346101db5760c06003193601126101db57612acf613055565b50612ad86130b2565b612ae0613078565b5060843561ffff811681036101db5760a4359067ffffffffffffffff82116101db5763ffffffff61ffff612b27829386612b2060a0973690600401613157565b5050613456565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b346101db5760006003193601126101db57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101db5760006003193601126101db57604051600d548082526020820190600d60005260206000209060005b818110612bd05761152d8561213481870382612fb5565b8254845260209093019260019283019201612bb9565b346101db5760406003193601126101db57612bff61309b565b60243580151581036101db57612c1b612cc191610140936133d3565b612c7160409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b346101db5760206003193601126101db576020612cde613055565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b346101db5760006003193601126101db57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101db5760006003193601126101db5761152d604051612d96606082612fb5565b602181527f434354505468726f756768434356546f6b656e506f6f6c20312e372e302d646560208201527f76000000000000000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190612ff6565b346101db5760206003193601126101db57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036101db57817faff2afbf0000000000000000000000000000000000000000000000000000000060209314908115612ed0575b8115612ea6575b8115612e7c575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483612e75565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150612e6e565b7f331710310000000000000000000000000000000000000000000000000000000081149150612e67565b6020810190811067ffffffffffffffff821117612f1657604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117612f1657604052565b60e0810190811067ffffffffffffffff821117612f1657604052565b6060810190811067ffffffffffffffff821117612f1657604052565b60a0810190811067ffffffffffffffff821117612f1657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612f1657604052565b919082519283825260005b8481106130405750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613001565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101db57565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036101db57565b6004359067ffffffffffffffff821682036101db57565b6024359067ffffffffffffffff821682036101db57565b359081151582036101db57565b602060408183019282815284518094520192019060005b8181106130fa5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016130ed565b6024359061ffff821682036101db57565b6044359061ffff821682036101db57565b359061ffff821682036101db57565b9181601f840112156101db5782359167ffffffffffffffff83116101db57602083818601950101116101db57565b92919267ffffffffffffffff8211612f1657604051916131cd601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184612fb5565b8294818452818301116101db578281602093846000960137010152565b9080601f830112156101db5781602061320593359101613185565b90565b9181601f840112156101db5782359167ffffffffffffffff83116101db576020808501948460051b0101116101db57565b9060406003198301126101db5760043567ffffffffffffffff811681036101db57916024359067ffffffffffffffff82116101db5761327a91600401613157565b9091565b67ffffffffffffffff8111612f165760051b60200190565b9080601f830112156101db578135906132ae8261327e565b926132bc6040519485612fb5565b82845260208085019360051b8201019182116101db57602001915b8183106132e45750505090565b823573ffffffffffffffffffffffffffffffffffffffff811681036101db578152602092830192016132d7565b9181601f840112156101db5782359167ffffffffffffffff83116101db576020808501948460081b0101116101db57565b6040519061334f82612f99565b60006080838281528260208201528260408201528260608201520152565b9060405161337a81612f99565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916133e5613342565b506133ee613342565b5061342257166000526008602052604060002090613205613416600261341b6134168661336d565b613b04565b940161336d565b169081600052600460205261343d613416604060002061336d565b916000526005602052613205613416604060002061336d565b9061ffff8060025460a01c1691169283151592838094613579575b6119035767ffffffffffffffff16600052600b6020526040600020916040519261349a84612f61565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a015261355e5761352f575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b8193975080929450106118d357505063ffffffff808061ffff9351169451169551169351169193929190600190565b50505050505092505050600090600090600090600090600090565b508215613471565b3567ffffffffffffffff811681036101db5790565b3573ffffffffffffffffffffffffffffffffffffffff811681036101db5790565b9067ffffffffffffffff61320592166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b519073ffffffffffffffffffffffffffffffffffffffff821682036101db57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff600354169586156137e3576136ef9467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c4860191613615565b9160028210156137b4578380600094819460a483015203915afa9081156112bd5760009161371b575090565b3d8083833e61372a8183612fb5565b8101906020818303126107d85780519067ffffffffffffffff82116137b0570181601f820112156107d8578051906137618261327e565b9361376f6040519586612fb5565b82855260208086019360051b8301019384116112b25750602001905b8282106137985750505090565b602080916137a5846135f4565b81520191019061378b565b8380fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b50505050505050506040516137f9602082612fb5565b60008152600036813790565b6040519061381282612f45565b60606020838281520152565b80518210156138325760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c921680156138aa575b602083101461387b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613870565b90604051918260008254926138c884613861565b808452936001811690811561393657506001146138ef575b506138ed92500383612fb5565b565b90506000929192526020600020906000915b81831061391a5750509060206138ed92820101386138e0565b6020919350806001915483858901015201910190918492613901565b602093506138ed9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386138e0565b67ffffffffffffffff16600052600860205261320560046040600020016138b4565b91908110156138325760081b0190565b3580151581036101db5790565b3561ffff811681036101db5790565b3563ffffffff811681036101db5790565b359063ffffffff821682036101db57565b91908110156138325760051b0190565b35906fffffffffffffffffffffffffffffffff821682036101db57565b91908260609103126101db57604051613a2b81612f7d565b6040613a55818395613a3c816130c9565b8552613a4a602082016139f6565b6020860152016139f6565b910152565b6fffffffffffffffffffffffffffffffff613a9860408093613a7b816130c9565b1515865283613a8c602083016139f6565b166020870152016139f6565b16910152565b818110613aa9575050565b60008155600101613a9e565b81810292918115918404141715613ac857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211613ac857565b613b0c613342565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691613b696020850193613b63613b5663ffffffff87511642613af7565b8560808901511690613ab5565b90614110565b80821015613b8257505b16825263ffffffff4216905290565b9050613b73565b908160209103126101db575180151581036101db5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101db570180359067ffffffffffffffff82116101db576020019181360383136101db57565b73ffffffffffffffffffffffffffffffffffffffff600154163303613c1357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115611bf35767ffffffffffffffff81516020830120921691826000526008602052613c72816005604060002001614536565b15613e395760005260096020526040600020815167ffffffffffffffff8111612f1657613c9f8254613861565b601f8111613e07575b506020601f8211600114613d415791613d1b827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593613d3195600091613d36575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190612ff6565b0390a2565b905084015138613cea565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110613def575092613d319492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610613db8575b5050811b019055611519565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880613dac565b9192602060018192868a015181550194019201613d71565b613e3390836000526020600020601f840160051c8101916020851061070357601f0160051c0190613a9e565b38613ca8565b5090611d106040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190612ff6565b906127109167ffffffffffffffff613e9760208301613581565b166000908152600b602052604090209161ffff1615613eca57606061ffff613ec6935460901c16910135613ab5565b0490565b606061ffff613ec6935460801c16910135613ab5565b815191929115614062576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff60208501511610613fff576138ed91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b606483614060604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604084015116158015906140f1575b614090576138ed9192613f23565b606483614060604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020840151161515614082565b91908201809211613ac857565b906040519182815491828252602082019060005260206000209260005b81811061414f5750506138ed92500383612fb5565b845483526001948501948794506020909301920161413a565b67ffffffffffffffff16614189816000526007602052604060002054151590565b156141d3575033600052600e602052604060002054156141a557565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80548210156138325760005260206000200190600090565b8054801561427c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061424d8282614200565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008181526007602052604090205480156143b0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613ac857600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613ac857818103614341575b50505061432d6006614218565b600052600760205260006040812055600190565b614398614352614363936006614200565b90549060031b1c9283926006614200565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526007602052604060002055388080614320565b5050600090565b90600182019181600052826020526040600020549081151560001461448a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191808311613ac85781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111613ac85783816144419503614453575b505050614218565b60005260205260006040812055600190565b6144736144636143639386614200565b90549060031b1c92839286614200565b905560005284602052604060002055388080614439565b50505050600090565b80549068010000000000000000821015612f1657816143639160016144ba94018155614200565b9055565b80600052600e602052604060002054156000146144f7576144e081600d614493565b600d5490600052600e602052604060002055600190565b50600090565b806000526007602052604060002054156000146144f75761451f816006614493565b600654906000526007602052604060002055600190565b60008281526001820160205260409020546143b0578061455883600193614493565b80549260005201602052604060002055600190565b9182549060ff8260a01c1615801561481a575b614814576fffffffffffffffffffffffffffffffff821691600185019081546145c563ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642613af7565b9081614776575b505084811061472a57508383106146265750506145fb6fffffffffffffffffffffffffffffffff928392613af7565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c9283156146be578161463e91613af7565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810193818511613ac85773ffffffffffffffffffffffffffffffffffffffff9461468991614110565b047fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b8286929396116147ea5761479192613b639160801c90613ab5565b808410156147e55750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806145cc565b61479c565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215614580565b6000818152600e602052604090205480156143b0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613ac857600d54907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613ac8578082036148b8575b5050506148a4600d614218565b600052600e60205260006040812055600190565b6148da6148c961436393600d614200565b90549060031b1c928392600d614200565b9055600052600e60205260406000205538808061489756fea164736f6c634300081a000a",
}

var CCTPThroughCCVTokenPoolABI = CCTPThroughCCVTokenPoolMetaData.ABI

var CCTPThroughCCVTokenPoolBin = CCTPThroughCCVTokenPoolMetaData.Bin

func DeployCCTPThroughCCVTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, rmnProxy common.Address, router common.Address, cctpVerifier common.Address, allowedCallers []common.Address) (common.Address, *types.Transaction, *CCTPThroughCCVTokenPool, error) {
	parsed, err := CCTPThroughCCVTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPThroughCCVTokenPoolBin), backend, token, localTokenDecimals, rmnProxy, router, cctpVerifier, allowedCallers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCTPThroughCCVTokenPool{address: address, abi: *parsed, CCTPThroughCCVTokenPoolCaller: CCTPThroughCCVTokenPoolCaller{contract: contract}, CCTPThroughCCVTokenPoolTransactor: CCTPThroughCCVTokenPoolTransactor{contract: contract}, CCTPThroughCCVTokenPoolFilterer: CCTPThroughCCVTokenPoolFilterer{contract: contract}}, nil
}

type CCTPThroughCCVTokenPool struct {
	address common.Address
	abi     abi.ABI
	CCTPThroughCCVTokenPoolCaller
	CCTPThroughCCVTokenPoolTransactor
	CCTPThroughCCVTokenPoolFilterer
}

type CCTPThroughCCVTokenPoolCaller struct {
	contract *bind.BoundContract
}

type CCTPThroughCCVTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type CCTPThroughCCVTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type CCTPThroughCCVTokenPoolSession struct {
	Contract     *CCTPThroughCCVTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCTPThroughCCVTokenPoolCallerSession struct {
	Contract *CCTPThroughCCVTokenPoolCaller
	CallOpts bind.CallOpts
}

type CCTPThroughCCVTokenPoolTransactorSession struct {
	Contract     *CCTPThroughCCVTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type CCTPThroughCCVTokenPoolRaw struct {
	Contract *CCTPThroughCCVTokenPool
}

type CCTPThroughCCVTokenPoolCallerRaw struct {
	Contract *CCTPThroughCCVTokenPoolCaller
}

type CCTPThroughCCVTokenPoolTransactorRaw struct {
	Contract *CCTPThroughCCVTokenPoolTransactor
}

func NewCCTPThroughCCVTokenPool(address common.Address, backend bind.ContractBackend) (*CCTPThroughCCVTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(CCTPThroughCCVTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCTPThroughCCVTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPool{address: address, abi: abi, CCTPThroughCCVTokenPoolCaller: CCTPThroughCCVTokenPoolCaller{contract: contract}, CCTPThroughCCVTokenPoolTransactor: CCTPThroughCCVTokenPoolTransactor{contract: contract}, CCTPThroughCCVTokenPoolFilterer: CCTPThroughCCVTokenPoolFilterer{contract: contract}}, nil
}

func NewCCTPThroughCCVTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*CCTPThroughCCVTokenPoolCaller, error) {
	contract, err := bindCCTPThroughCCVTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolCaller{contract: contract}, nil
}

func NewCCTPThroughCCVTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CCTPThroughCCVTokenPoolTransactor, error) {
	contract, err := bindCCTPThroughCCVTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolTransactor{contract: contract}, nil
}

func NewCCTPThroughCCVTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CCTPThroughCCVTokenPoolFilterer, error) {
	contract, err := bindCCTPThroughCCVTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolFilterer{contract: contract}, nil
}

func bindCCTPThroughCCVTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCTPThroughCCVTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPThroughCCVTokenPool.Contract.CCTPThroughCCVTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.CCTPThroughCCVTokenPoolTransactor.contract.Transfer(opts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.CCTPThroughCCVTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPThroughCCVTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.contract.Transfer(opts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getAdvancedPoolHooks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetAdvancedPoolHooks(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetAdvancedPoolHooks(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetAllAuthorizedCallers(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetAllAuthorizedCallers(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetCCTPVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getCCTPVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetCCTPVerifier() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetCCTPVerifier(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetCCTPVerifier() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetCCTPVerifier(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeAggregator = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetDynamicConfig(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetDynamicConfig(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetFee(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetFee(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetMinBlockConfirmation(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetMinBlockConfirmation(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRemotePools(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRemotePools(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRemoteToken(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRemoteToken(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRequiredCCVs(&_CCTPThroughCCVTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRequiredCCVs(&_CCTPThroughCCVTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRmnProxy(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRmnProxy(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetSupportedChains(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetSupportedChains(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetToken() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetToken(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetToken(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetTokenDecimals(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetTokenDecimals(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetTokenTransferFeeConfig(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetTokenTransferFeeConfig(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _CCTPThroughCCVTokenPool.Contract.IsRemotePool(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _CCTPThroughCCVTokenPool.Contract.IsRemotePool(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _CCTPThroughCCVTokenPool.Contract.IsSupportedChain(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _CCTPThroughCCVTokenPool.Contract.IsSupportedChain(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _CCTPThroughCCVTokenPool.Contract.IsSupportedToken(&_CCTPThroughCCVTokenPool.CallOpts, token)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _CCTPThroughCCVTokenPool.Contract.IsSupportedToken(&_CCTPThroughCCVTokenPool.CallOpts, token)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) Owner() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.Owner(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) Owner() (common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.Owner(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPThroughCCVTokenPool.Contract.SupportsInterface(&_CCTPThroughCCVTokenPool.CallOpts, interfaceId)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPThroughCCVTokenPool.Contract.SupportsInterface(&_CCTPThroughCCVTokenPool.CallOpts, interfaceId)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) TypeAndVersion() (string, error) {
	return _CCTPThroughCCVTokenPool.Contract.TypeAndVersion(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _CCTPThroughCCVTokenPool.Contract.TypeAndVersion(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.AcceptOwnership(&_CCTPThroughCCVTokenPool.TransactOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.AcceptOwnership(&_CCTPThroughCCVTokenPool.TransactOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.AddRemotePool(&_CCTPThroughCCVTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.AddRemotePool(&_CCTPThroughCCVTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_CCTPThroughCCVTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_CCTPThroughCCVTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ApplyChainUpdates(&_CCTPThroughCCVTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ApplyChainUpdates(&_CCTPThroughCCVTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_CCTPThroughCCVTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_CCTPThroughCCVTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, arg0 PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "lockOrBurn", arg0)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) LockOrBurn(arg0 PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.LockOrBurn(&_CCTPThroughCCVTokenPool.TransactOpts, arg0)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) LockOrBurn(arg0 PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.LockOrBurn(&_CCTPThroughCCVTokenPool.TransactOpts, arg0)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.LockOrBurn0(&_CCTPThroughCCVTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.LockOrBurn0(&_CCTPThroughCCVTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "releaseOrMint", arg0)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) ReleaseOrMint(arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint(&_CCTPThroughCCVTokenPool.TransactOpts, arg0)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) ReleaseOrMint(arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint(&_CCTPThroughCCVTokenPool.TransactOpts, arg0)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint0(&_CCTPThroughCCVTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint0(&_CCTPThroughCCVTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.RemoveRemotePool(&_CCTPThroughCCVTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.RemoveRemotePool(&_CCTPThroughCCVTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAggregator common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAggregator)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAggregator common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetDynamicConfig(&_CCTPThroughCCVTokenPool.TransactOpts, router, rateLimitAdmin, feeAggregator)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAggregator common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetDynamicConfig(&_CCTPThroughCCVTokenPool.TransactOpts, router, rateLimitAdmin, feeAggregator)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "setMinBlockConfirmation", minBlockConfirmation)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetMinBlockConfirmation(&_CCTPThroughCCVTokenPool.TransactOpts, minBlockConfirmation)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetMinBlockConfirmation(&_CCTPThroughCCVTokenPool.TransactOpts, minBlockConfirmation)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetRateLimitConfig(&_CCTPThroughCCVTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetRateLimitConfig(&_CCTPThroughCCVTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.TransferOwnership(&_CCTPThroughCCVTokenPool.TransactOpts, to)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.TransferOwnership(&_CCTPThroughCCVTokenPool.TransactOpts, to)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "updateAdvancedPoolHooks", newHook)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.UpdateAdvancedPoolHooks(&_CCTPThroughCCVTokenPool.TransactOpts, newHook)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.UpdateAdvancedPoolHooks(&_CCTPThroughCCVTokenPool.TransactOpts, newHook)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.WithdrawFeeTokens(&_CCTPThroughCCVTokenPool.TransactOpts, feeTokens)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.WithdrawFeeTokens(&_CCTPThroughCCVTokenPool.TransactOpts, feeTokens)
}

type CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdatedIterator struct {
	Event *CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated)
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
		it.Event = new(CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated)
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

func (it *CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated struct {
	OldHook common.Address
	NewHook common.Address
	Raw     types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdatedIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdatedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "AdvancedPoolHooksUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseAdvancedPoolHooksUpdated(log types.Log) (*CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated, error) {
	event := new(CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolAuthorizedCallerAddedIterator struct {
	Event *CCTPThroughCCVTokenPoolAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolAuthorizedCallerAdded)
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
		it.Event = new(CCTPThroughCCVTokenPoolAuthorizedCallerAdded)
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

func (it *CCTPThroughCCVTokenPoolAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolAuthorizedCallerAddedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolAuthorizedCallerAdded)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseAuthorizedCallerAdded(log types.Log) (*CCTPThroughCCVTokenPoolAuthorizedCallerAdded, error) {
	event := new(CCTPThroughCCVTokenPoolAuthorizedCallerAdded)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolAuthorizedCallerRemovedIterator struct {
	Event *CCTPThroughCCVTokenPoolAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolAuthorizedCallerRemoved)
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
		it.Event = new(CCTPThroughCCVTokenPoolAuthorizedCallerRemoved)
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

func (it *CCTPThroughCCVTokenPoolAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolAuthorizedCallerRemovedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolAuthorizedCallerRemoved)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*CCTPThroughCCVTokenPoolAuthorizedCallerRemoved, error) {
	event := new(CCTPThroughCCVTokenPoolAuthorizedCallerRemoved)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolChainAddedIterator struct {
	Event *CCTPThroughCCVTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolChainAdded)
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
		it.Event = new(CCTPThroughCCVTokenPoolChainAdded)
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

func (it *CCTPThroughCCVTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolChainAddedIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolChainAddedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolChainAdded)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseChainAdded(log types.Log) (*CCTPThroughCCVTokenPoolChainAdded, error) {
	event := new(CCTPThroughCCVTokenPoolChainAdded)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolChainRemovedIterator struct {
	Event *CCTPThroughCCVTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolChainRemoved)
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
		it.Event = new(CCTPThroughCCVTokenPoolChainRemoved)
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

func (it *CCTPThroughCCVTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolChainRemovedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolChainRemoved)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseChainRemoved(log types.Log) (*CCTPThroughCCVTokenPoolChainRemoved, error) {
	event := new(CCTPThroughCCVTokenPoolChainRemoved)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolDynamicConfigSetIterator struct {
	Event *CCTPThroughCCVTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolDynamicConfigSet)
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
		it.Event = new(CCTPThroughCCVTokenPoolDynamicConfigSet)
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

func (it *CCTPThroughCCVTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolDynamicConfigSet struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAggregator  common.Address
	Raw            types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolDynamicConfigSetIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolDynamicConfigSet)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*CCTPThroughCCVTokenPoolDynamicConfigSet, error) {
	event := new(CCTPThroughCCVTokenPoolDynamicConfigSet)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolFeeTokenWithdrawnIterator struct {
	Event *CCTPThroughCCVTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(CCTPThroughCCVTokenPoolFeeTokenWithdrawn)
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

func (it *CCTPThroughCCVTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPThroughCCVTokenPoolFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolFeeTokenWithdrawnIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolFeeTokenWithdrawn)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CCTPThroughCCVTokenPoolFeeTokenWithdrawn, error) {
	event := new(CCTPThroughCCVTokenPoolFeeTokenWithdrawn)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolInboundRateLimitConsumedIterator struct {
	Event *CCTPThroughCCVTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(CCTPThroughCCVTokenPoolInboundRateLimitConsumed)
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

func (it *CCTPThroughCCVTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolInboundRateLimitConsumedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolInboundRateLimitConsumed)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolInboundRateLimitConsumed, error) {
	event := new(CCTPThroughCCVTokenPoolInboundRateLimitConsumed)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolLockedOrBurnedIterator struct {
	Event *CCTPThroughCCVTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolLockedOrBurned)
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
		it.Event = new(CCTPThroughCCVTokenPoolLockedOrBurned)
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

func (it *CCTPThroughCCVTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolLockedOrBurnedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolLockedOrBurned)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*CCTPThroughCCVTokenPoolLockedOrBurned, error) {
	event := new(CCTPThroughCCVTokenPoolLockedOrBurned)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolMinBlockConfirmationSetIterator struct {
	Event *CCTPThroughCCVTokenPoolMinBlockConfirmationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolMinBlockConfirmationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolMinBlockConfirmationSet)
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
		it.Event = new(CCTPThroughCCVTokenPoolMinBlockConfirmationSet)
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

func (it *CCTPThroughCCVTokenPoolMinBlockConfirmationSetIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolMinBlockConfirmationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolMinBlockConfirmationSet struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolMinBlockConfirmationSetIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolMinBlockConfirmationSetIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "MinBlockConfirmationSet", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolMinBlockConfirmationSet) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolMinBlockConfirmationSet)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseMinBlockConfirmationSet(log types.Log) (*CCTPThroughCCVTokenPoolMinBlockConfirmationSet, error) {
	event := new(CCTPThroughCCVTokenPoolMinBlockConfirmationSet)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *CCTPThroughCCVTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(CCTPThroughCCVTokenPoolOutboundRateLimitConsumed)
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

func (it *CCTPThroughCCVTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolOutboundRateLimitConsumedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolOutboundRateLimitConsumed)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolOutboundRateLimitConsumed, error) {
	event := new(CCTPThroughCCVTokenPoolOutboundRateLimitConsumed)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolOwnershipTransferRequestedIterator struct {
	Event *CCTPThroughCCVTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolOwnershipTransferRequested)
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
		it.Event = new(CCTPThroughCCVTokenPoolOwnershipTransferRequested)
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

func (it *CCTPThroughCCVTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPThroughCCVTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolOwnershipTransferRequestedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolOwnershipTransferRequested)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCTPThroughCCVTokenPoolOwnershipTransferRequested, error) {
	event := new(CCTPThroughCCVTokenPoolOwnershipTransferRequested)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolOwnershipTransferredIterator struct {
	Event *CCTPThroughCCVTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolOwnershipTransferred)
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
		it.Event = new(CCTPThroughCCVTokenPoolOwnershipTransferred)
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

func (it *CCTPThroughCCVTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPThroughCCVTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolOwnershipTransferredIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolOwnershipTransferred)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*CCTPThroughCCVTokenPoolOwnershipTransferred, error) {
	event := new(CCTPThroughCCVTokenPoolOwnershipTransferred)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolRateLimitConfiguredIterator struct {
	Event *CCTPThroughCCVTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolRateLimitConfigured)
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
		it.Event = new(CCTPThroughCCVTokenPoolRateLimitConfigured)
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

func (it *CCTPThroughCCVTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolRateLimitConfiguredIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolRateLimitConfigured)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*CCTPThroughCCVTokenPoolRateLimitConfigured, error) {
	event := new(CCTPThroughCCVTokenPoolRateLimitConfigured)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolReleasedOrMintedIterator struct {
	Event *CCTPThroughCCVTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolReleasedOrMinted)
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
		it.Event = new(CCTPThroughCCVTokenPoolReleasedOrMinted)
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

func (it *CCTPThroughCCVTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolReleasedOrMintedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolReleasedOrMinted)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*CCTPThroughCCVTokenPoolReleasedOrMinted, error) {
	event := new(CCTPThroughCCVTokenPoolReleasedOrMinted)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolRemotePoolAddedIterator struct {
	Event *CCTPThroughCCVTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolRemotePoolAdded)
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
		it.Event = new(CCTPThroughCCVTokenPoolRemotePoolAdded)
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

func (it *CCTPThroughCCVTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolRemotePoolAddedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolRemotePoolAdded)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*CCTPThroughCCVTokenPoolRemotePoolAdded, error) {
	event := new(CCTPThroughCCVTokenPoolRemotePoolAdded)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolRemotePoolRemovedIterator struct {
	Event *CCTPThroughCCVTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolRemotePoolRemoved)
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
		it.Event = new(CCTPThroughCCVTokenPoolRemotePoolRemoved)
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

func (it *CCTPThroughCCVTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolRemotePoolRemovedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolRemotePoolRemoved)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*CCTPThroughCCVTokenPoolRemotePoolRemoved, error) {
	event := new(CCTPThroughCCVTokenPoolRemotePoolRemoved)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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
	FeeAggregator  common.Address
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}

func (CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated) Topic() common.Hash {
	return common.HexToHash("0xbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209")
}

func (CCTPThroughCCVTokenPoolAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (CCTPThroughCCVTokenPoolAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (CCTPThroughCCVTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (CCTPThroughCCVTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (CCTPThroughCCVTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (CCTPThroughCCVTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CCTPThroughCCVTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (CCTPThroughCCVTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (CCTPThroughCCVTokenPoolMinBlockConfirmationSet) Topic() common.Hash {
	return common.HexToHash("0xa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb")
}

func (CCTPThroughCCVTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (CCTPThroughCCVTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCTPThroughCCVTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CCTPThroughCCVTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (CCTPThroughCCVTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (CCTPThroughCCVTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (CCTPThroughCCVTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPool) Address() common.Address {
	return _CCTPThroughCCVTokenPool.address
}

type CCTPThroughCCVTokenPoolInterface interface {
	GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetCCTPVerifier(opts *bind.CallOpts) (common.Address, error)

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

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, arg0 PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, arg0 PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAggregator common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*CCTPThroughCCVTokenPoolAdvancedPoolHooksUpdated, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*CCTPThroughCCVTokenPoolAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*CCTPThroughCCVTokenPoolAuthorizedCallerRemoved, error)

	FilterChainAdded(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*CCTPThroughCCVTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*CCTPThroughCCVTokenPoolChainRemoved, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CCTPThroughCCVTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPThroughCCVTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCTPThroughCCVTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*CCTPThroughCCVTokenPoolLockedOrBurned, error)

	FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolMinBlockConfirmationSetIterator, error)

	WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolMinBlockConfirmationSet) (event.Subscription, error)

	ParseMinBlockConfirmationSet(log types.Log) (*CCTPThroughCCVTokenPoolMinBlockConfirmationSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPThroughCCVTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPThroughCCVTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPThroughCCVTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPThroughCCVTokenPoolOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*CCTPThroughCCVTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*CCTPThroughCCVTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*CCTPThroughCCVTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*CCTPThroughCCVTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*CCTPThroughCCVTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*CCTPThroughCCVTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
