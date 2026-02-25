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

var CCTPThroughCCVTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCTPVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmations\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationsRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmations\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationsInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationsOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationsSet\",\"inputs\":[{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmations\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationsFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationsTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CCVNotSetOnResolver\",\"inputs\":[{\"name\":\"resolver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"IPoolV1NotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmations\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103745761500c803803809161001d82856103dd565b833981019060c0818303126103745780516001600160a01b038116908181036103745761004c60208401610400565b906100596040850161040e565b906100666060860161040e565b936100736080870161040e565b60a087015190966001600160401b03821161037457019680601f89011215610374578751976001600160401b0389116102ff578860051b9060208201996100bd6040519b8c6103dd565b8a526020808b019282010192831161037457602001905b8282106103c55750505033156103b457600180546001600160a01b03191633179055801580156103a3575b8015610392575b6103815760049260209260805260c0526040519283809263313ce56760e01b82525afa60009181610340575b50610315575b5060a052600380546001600160a01b0319908116909155600280549091166001600160a01b039290921691909117905560405160209061017882826103dd565b60008152600036813760408051949085016001600160401b038111868210176102ff576040528452808285015260005b815181101561020f576001906001600160a01b036101c68285610422565b5116846101d282610464565b6101df575b5050016101a8565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138846101d7565b5050915160005b8151811015610287576001600160a01b036102318284610422565b5116908115610276577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8583610268600195610562565b50604051908152a101610216565b6342bcdf7f60e11b60005260046000fd5b6001600160a01b03831660e052604051614a4990816105c3823960805181818161197001528181611b570152818161276f0152818161295c01528181612d5b0152612db5015260a05181612bac015260c051818181610b8d01528181611a0b015261280a015260e05181818161112901526126880152f35b634e487b7160e01b600052604160045260246000fd5b60ff1660ff82168181036103295750610138565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610379575b8161035c602093836103dd565b810103126103745761036d90610400565b9038610132565b600080fd5b3d915061034f565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610106565b506001600160a01b038516156100ff565b639b15e16f60e01b60005260046000fd5b602080916103d28461040e565b8152019101906100d4565b601f909101601f19168101906001600160401b038211908210176102ff57604052565b519060ff8216820361037457565b51906001600160a01b038216820361037457565b80518210156104365760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104365760005260206000200190600090565b6000818152600e6020526040902054801561055b57600019810181811161054557600d54600019810191908211610545578082036104f4575b505050600d5480156104de57600019016104b881600d61044c565b8154906000199060031b1b19169055600d55600052600e60205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61052d61050561051693600d61044c565b90549060031b1c928392600d61044c565b819391549060031b91821b91600019901b19161790565b9055600052600e60205260406000205538808061049d565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600e602052604060002054156000146105bc57600d54680100000000000000008110156102ff576105a3610516826001859401600d55600d61044c565b9055600d5490600052600e602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714612e6057508063181f5a7714612dd957806321df0da714612d88578063240028e814612d245780632422ac4514612c455780632451a62714612bd057806324f65ee714612b925780632c06340414612af957806337a3210d14612ac55780633907753714612a5c578063489a68f2146126f35780634c5ef0ed146126ac578063615521a71461265b57806362ddd3c4146125d45780637437ff9f1461258657806379ba5097146124bf5780638926f54f1461247957806389720a62146123f95780638da5cb5b146123c557806391a2749a146122175780639a4575b9146121b4578063a42a7b8b1461204d578063acfecf9114611f55578063ae39a25714611dca578063b1c71c65146118db578063b79465801461189e578063b8d5005e14611879578063bfeffd3f146117cd578063c4bffe2b146116a2578063c7230a60146113ef578063d4d6de2314611350578063d8aa3f4014611002578063dc04fa1f14610bb1578063dc0bd97114610b60578063dcbd41bc1461095c578063e8a1da17146102895763f2fde38b146101b857600080fd5b346102865760206003193601126102865773ffffffffffffffffffffffffffffffffffffffff6101e66130bd565b6101ee613d75565b1633811461025e57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102865760406003193601126102865760043567ffffffffffffffff81116107b0576102bb90369060040161336b565b9060243567ffffffffffffffff811161095857906102de8492369060040161336b565b9390916102e9613d75565b83905b8282106107bd5750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156107b9578060051b830135858112156107ac578301610120813603126107ac576040519461035086613001565b813567ffffffffffffffff811681036107b4578652602082013567ffffffffffffffff81116107b05782019436601f870112156107b057853595610393876132d8565b966103a1604051988961301d565b80885260208089019160051b830101903682116107ac5760208301905b828210610779575050505060208701958652604083013567ffffffffffffffff8111610775576103f19036908501613275565b916040880192835261041b6104093660608701613be5565b9460608a0195865260c0369101613be5565b95608089019687528351511561074d5761043f67ffffffffffffffff8a5116614647565b156107165767ffffffffffffffff895116825260086020526040822061046686518261402a565b61047488516002830161402a565b6004855191019080519067ffffffffffffffff82116106e9576104978354613a33565b601f81116106ae575b50602090601f831160011461060f576104ee9291869183610604575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b88518051821015610528579061052260019261051b8367ffffffffffffffff8f5116926139f0565b5190613dc0565b016104f3565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29391999750956105f667ffffffffffffffff60019796949851169251935191516105c261058d6040519687968752610100602088015261010087019061305e565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101939290919361031f565b015190508e806104bc565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b818110610696575090846001959493921061065f575b505050811b0190556104f1565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080610652565b9293602060018192878601518155019501930161063c565b6106d99084875260208720601f850160051c810191602086106106df575b601f0160051c0190613c70565b8d6104a0565b90915081906106cc565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116107a85760209161079d8392833691890101613275565b8152019101906103be565b8680fd5b8480fd5b5080fd5b600080fd5b8380f35b9267ffffffffffffffff6107df6107da8486889a9699979a613bb8565b61360c565b16916107ea836143f5565b1561092c57828452600860205261080660056040862001614267565b94845b865181101561083f57600190858752600860205261083860056040892001610831838b6139f0565b5190614501565b5001610809565b509396929094509490948087526008602052600560408820888155886001820155886002820155886003820155886004820161087b8154613a33565b806108eb575b50505001805490888155816108cd575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a1019091949392946102ec565b885260208820908101905b81811015610891578881556001016108d8565b601f81116001146109015750555b888a80610881565b8183526020832061091c91601f01861c810190600101613c70565b80825281602081209155556108f9565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346102865760206003193601126102865760043567ffffffffffffffff81116107b05761098e90369060040161339c565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580610b3e575b610b1257825b8181106109c1578380f35b6109cc818385613b6a565b67ffffffffffffffff6109de8261360c565b16906109f7826000526007602052604060002054151590565b15610ae657907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610aa6610a80602060019897018b610a3882613b7a565b15610aad578790526004602052610a5f60408d20610a593660408801613be5565b9061402a565b868c526005602052610a7b60408d20610a593660a08801613be5565b613b7a565b916040519215158352610a996020840160408301613c2c565b60a0608084019101613c2c565ba2016109b6565b60026040828a610a7b94526008602052610acf828220610a5936858c01613be5565b8a815260086020522001610a593660a08801613be5565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600154163314156109b0565b5034610286578060031936011261028657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102865760406003193601126102865760043567ffffffffffffffff81116107b057610be390369060040161339c565b60243567ffffffffffffffff811161095857610c0390369060040161336b565b919092610c0e613d75565b845b828110610c7a57505050825b818110610c27578380f35b8067ffffffffffffffff610c416107da6001948688613bb8565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610c1c565b610c886107da828585613b6a565b610c93828585613b6a565b90602082019060e0830190610ca782613b7a565b15610fcd5760a0840161271061ffff610cbf83613b87565b161015610fbe5760c085019161271061ffff610cda85613b87565b161015610f865763ffffffff610cef86613b96565b1615610f515767ffffffffffffffff1694858c52600b60205260408c20610d1586613b96565b63ffffffff16908054906040840191610d2d83613b96565b60201b67ffffffff0000000016936060860194610d4986613b96565b60401b6bffffffff0000000000000000169660800196610d6888613b96565b60601b6fffffffff0000000000000000000000001691610d878a613b87565b60801b71ffff000000000000000000000000000000001693610da88c613b87565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610e5b87613b7a565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610eac90613ba7565b63ffffffff168752610ebd90613ba7565b63ffffffff166020870152610ed190613ba7565b63ffffffff166040860152610ee590613ba7565b63ffffffff166060850152610ef9906131d3565b61ffff166080840152610f0b906131d3565b61ffff1660a0830152610f1d90613154565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c10565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff610f9586613b87565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff610f95602493613b87565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346102865760806003193601126102865761101c6130bd565b5061102561313d565b61102d6131c2565b506064359067ffffffffffffffff82116107755761105867ffffffffffffffff9236906004016131e2565b50508260c060405161106981612fc9565b8281528260208201528260408201528260608201528260808201528260a0820152015216808252600b6020526040822090604051916110a783612fc9565b549063ffffffff82168352602083019363ffffffff8360201c168552604084019463ffffffff8460401c168652606085019063ffffffff8560601c168252608086019261ffff8660801c16845260a087019461ffff8760901c16865260ff60c089019760a01c161515875273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690604051907f958021a7000000000000000000000000000000000000000000000000000000008252600482015260406024820152826044820152602081606481855afa80156113455783906112f8575b73ffffffffffffffffffffffffffffffffffffffff9150169081156112cd57506060600491604051928380927f7437ff9f0000000000000000000000000000000000000000000000000000000082525afa9182156112c157809261124c575b505061ffff9493919263ffffffff60e09981889687604083970151168952816040519c51168c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b9091506060823d6060116112b9575b816112686060938361301d565b8101031261028657604080519261127e84612fe5565b611287816136ce565b8452611295602082016136ce565b602085015201519061ffff82168203610286575060408201528263ffffffff6111f8565b3d915061125b565b604051903d90823e3d90fd5b7f4172d660000000000000000000000000000000000000000000000000000000008352600452602482fd5b506020813d60201161133d575b816113126020938361301d565b810103126107755761133873ffffffffffffffffffffffffffffffffffffffff916136ce565b611199565b3d9150611305565b6040513d85823e3d90fd5b50346102865760206003193601126102865760043561ffff811690818103610775577f46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a019160209161139f613d75565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006002549260a01b16911617600255604051908152a180f35b50346102865760406003193601126102865760043567ffffffffffffffff81116107b05761142190369060040161336b565b9061142a613103565b9173ffffffffffffffffffffffffffffffffffffffff6001541633141580611680575b6116545773ffffffffffffffffffffffffffffffffffffffff831690811561162c57845b81811061147c578580f35b73ffffffffffffffffffffffffffffffffffffffff6114a461149f838588613bb8565b613621565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156116215788916115ec575b50806114f9575b5050600101611471565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff8a16602484015260448084018590528352918a919061155a60648261301d565b519082865af1156115e15787513d6115d85750813b155b6115ac5790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a390386114ef565b602488837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b60011415611571565b6040513d89823e3d90fd5b90506020813d8211611619575b816116066020938361301d565b810103126116155751386114e8565b8780fd5b3d91506115f9565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600c541633141561144d565b5034610286578060031936011261028657604051906006548083528260208101600684526020842092845b8181106117b45750506116e29250038361301d565b81516117066116f0826132d8565b916116fe604051938461301d565b8083526132d8565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015611765578067ffffffffffffffff611752600193886139f0565b511661175e82866139f0565b5201611733565b50925090604051928392602084019060208552518091526040840192915b818110611791575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611783565b84548352600194850194879450602090930192016116cd565b50346102865760206003193601126102865760043573ffffffffffffffffffffffffffffffffffffffff81168091036107b057611808613d75565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a1161760035580f35b5034610286578060031936011261028657602061ffff60025460a01c16604051908152f35b5034610286576020600319360112610286576118d76118c36118be613126565b613b48565b60405191829160208352602083019061305e565b0390f35b5034610286576060600319360112610286576004359067ffffffffffffffff82116102865760a06003198336030112610286576119166131b1565b60443567ffffffffffffffff8111610775576119396119499136906004016131e2565b6119416139d7565b503691613210565b50608483019061195882613621565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611d8057602484019377ffffffffffffffff000000000000000000000000000000006119be8661360c565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d75578591611d46575b50611d1e5790606461ffff92611a58611a538861360c565b6142b2565b01359350168015611c8f5761ffff60025460a01c16908115611c6557818110611c355750506118be611bfa93611a90611b8d93613621565b7f996b829383cc7e136842d4c4c175083bcf4e20807c7432105c1b794ba885e7768567ffffffffffffffff611ac48561360c565b1692836000526004602052611ade818360406000206146b7565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b611b108161360c565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10606067ffffffffffffffff6040519373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001685523360208601528860408601521692a261360c565b90611c2b6040517f3047587c00000000000000000000000000000000000000000000000000000000602082015260048152611bc960248261301d565b60405193611bd685612fad565b8452602084019081526040519485946040865251604080870152608086019061305e565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc085830301606086015261305e565b9060208301520390f35b7f1f5b9f770000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b506118be611bfa93611ca3611b8d93613621565b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448567ffffffffffffffff611cd78561360c565b1692836000526008602052611cf1818360406000206146b7565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611b07565b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611d68915060203d602011611d6e575b611d60818361301d565b810190613d0c565b38611a3b565b503d611d56565b6040513d87823e3d90fd5b60248373ffffffffffffffffffffffffffffffffffffffff611da185613621565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b503461028657606060031936011261028657611de46130bd565b90611ded613103565b6044359273ffffffffffffffffffffffffffffffffffffffff841680850361095857611e17613d75565b73ffffffffffffffffffffffffffffffffffffffff82168015611f2d5794611f27917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c556040519384938491604091949373ffffffffffffffffffffffffffffffffffffffff809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346102865767ffffffffffffffff611f6d36613293565b929091611f78613d75565b1691611f91836000526007602052604060002054151590565b1561092c578284526008602052611fc060056040862001611fb3368486613210565b6020815191012090614501565b1561200557907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611fff6040519283926020845260208401916136ef565b0390a280f35b82612049836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916136ef565b0390fd5b50346102865760206003193601126102865767ffffffffffffffff612070613126565b168152600860205261208760056040832001614267565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120cc6120b6836132d8565b926120c4604051948561301d565b8084526132d8565b01835b8181106121a3575050825b825181101561212057806120f0600192856139f0565b518552600960205261210460408620613a86565b61210e82856139f0565b5261211981846139f0565b50016120da565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061215857505050500390f35b91936020612193827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652885161305e565b9601920192018594939192612149565b8060606020809386010152016120cf565b50346102865760206003193601126102865760043567ffffffffffffffff81116107b05760031960a09136030112610286576004906121f16139d7565b507f690a7a40000000000000000000000000000000000000000000000000000000008152fd5b50346102865760206003193601126102865760043567ffffffffffffffff81116107b057604060031982360301126107b0576040519061225682612fad565b806004013567ffffffffffffffff81116109585761227a90600436918401016132f0565b825260248101359067ffffffffffffffff82116109585760046122a092369201016132f0565b602082019081526122af613d75565b5191805b8351811015612326578073ffffffffffffffffffffffffffffffffffffffff6122de600193876139f0565b51166122e98161496c565b6122f5575b50016122b3565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1386122ee565b509051815b81518110156123c15773ffffffffffffffffffffffffffffffffffffffff61235382846139f0565b5116801561239957907feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef60208361238b600195614608565b50604051908152a10161232b565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8280f35b5034610286578060031936011261028657602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102865760c0600319360112610286576124136130bd565b61241b61313d565b60643561ffff811681036109585760843567ffffffffffffffff81116107ac576124499036906004016131e2565b93909260a435956002871015610286576118d761246d88888888604435888a61372e565b60405191829182613161565b50346102865760206003193601126102865760206124b567ffffffffffffffff6124a1613126565b166000526007602052604060002054151590565b6040519015158152f35b5034610286578060031936011261028657805473ffffffffffffffffffffffffffffffffffffffff8116330361255e577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b5034610286578060031936011261028657600254600a54600c546040805173ffffffffffffffffffffffffffffffffffffffff94851681529284166020840152921691810191909152606090f35b5034610286576125e336613293565b6125ef93929193613d75565b67ffffffffffffffff8216612611816000526007602052604060002054151590565b15612630575061262d9293612627913691613210565b90613dc0565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b5034610286578060031936011261028657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610286576040600319360112610286576126c6613126565b906024359067ffffffffffffffff82116102865760206124b5846126ed3660048701613275565b90613642565b5034610286576040600319360112610286576004359067ffffffffffffffff821161028657816004019161010060031982360301126107b0576127346131b1565b918060405161274281612f62565b52606482013592608483019061275782613621565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611d8057602484019577ffffffffffffffff000000000000000000000000000000006127bd8861360c565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d75578591612a3d575b50611d1e5761284b611a538861360c565b6128548761360c565b9061287160a48701926126ed61286a8585613d24565b3691613210565b156129f657505061293f60446129386020987ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966080968a61ffff67ffffffffffffffff981615156000146129a8576128ea7f63335ad9e238acd0e9e6c1c20f529ffbea4cda73578c329a7aa7e9d61e5cdcc592613621565b61290b818360408c6128fb8a61360c565b16808952600560205297206146b7565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a261360c565b9501613621565b9373ffffffffffffffffffffffffffffffffffffffff60405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a28060405161299f81612f62565b52604051908152f35b6129d27f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c92613621565b61290b8183600260408d6129e58b61360c565b16808a5260086020529820016146b7565b612a009250613d24565b6120496040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916136ef565b612a56915060203d602011611d6e57611d60818361301d565b3861283a565b50346102865760206003193601126102865760043567ffffffffffffffff81116107b05760031961010091360301126102865780600491604051612a9f81612f62565b527f690a7a40000000000000000000000000000000000000000000000000000000008152fd5b5034610286578060031936011261028657602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b50346102865760c060031936011261028657612b136130bd565b50612b1c61313d565b612b246130e0565b506084359161ffff831683036102865760a4359067ffffffffffffffff82116102865760a063ffffffff8061ffff612b6b8888612b643660048b016131e2565b50506134e1565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b5034610286578060031936011261028657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610286578060031936011261028657604051600d8054808352908352909160208301917fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5915b818110612c2f576118d78561246d8187038261301d565b8254845260209093019260019283019201612c18565b503461028657604060031936011261028657612c5f613126565b60243591821515830361028657610140612d22612c7c858561345e565b612cd260409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b503461028657602060031936011261028657602090612d416130bd565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b5034610286578060031936011261028657602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610286578060031936011261028657506118d7604051612dfc60608261301d565b602181527f434354505468726f756768434356546f6b656e506f6f6c20322e302e302d646560208201527f7600000000000000000000000000000000000000000000000000000000000000604082015260405191829160208352602083019061305e565b9050346107b05760206003193601126107b0576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361077557602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115612f38575b8115612f0e575b8115612ee4575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438612edd565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150612ed6565b7f331710310000000000000000000000000000000000000000000000000000000081149150612ecf565b6020810190811067ffffffffffffffff821117612f7e57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117612f7e57604052565b60e0810190811067ffffffffffffffff821117612f7e57604052565b6060810190811067ffffffffffffffff821117612f7e57604052565b60a0810190811067ffffffffffffffff821117612f7e57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612f7e57604052565b919082519283825260005b8481106130a85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613069565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036107b457565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036107b457565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036107b457565b6004359067ffffffffffffffff821682036107b457565b6024359067ffffffffffffffff821682036107b457565b359081151582036107b457565b602060408183019282815284518094520192019060005b8181106131855750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613178565b6024359061ffff821682036107b457565b6044359061ffff821682036107b457565b359061ffff821682036107b457565b9181601f840112156107b45782359167ffffffffffffffff83116107b457602083818601950101116107b457565b92919267ffffffffffffffff8211612f7e5760405191613258601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0166020018461301d565b8294818452818301116107b4578281602093846000960137010152565b9080601f830112156107b45781602061329093359101613210565b90565b9060406003198301126107b45760043567ffffffffffffffff811681036107b457916024359067ffffffffffffffff82116107b4576132d4916004016131e2565b9091565b67ffffffffffffffff8111612f7e5760051b60200190565b9080601f830112156107b457813590613308826132d8565b92613316604051948561301d565b82845260208085019360051b8201019182116107b457602001915b81831061333e5750505090565b823573ffffffffffffffffffffffffffffffffffffffff811681036107b457815260209283019201613331565b9181601f840112156107b45782359167ffffffffffffffff83116107b4576020808501948460051b0101116107b457565b9181601f840112156107b45782359167ffffffffffffffff83116107b4576020808501948460081b0101116107b457565b604051906133da82613001565b60006080838281528260208201528260408201528260608201520152565b9060405161340581613001565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916134706133cd565b506134796133cd565b506134ad571660005260086020526040600020906132906134a160026134a66134a1866133f8565b613c87565b94016133f8565b16908160005260046020526134c86134a160406000206133f8565b9160005260056020526132906134a160406000206133f8565b9061ffff8060025460a01c1691169283151592838094613604575b611c655767ffffffffffffffff16600052600b6020526040600020916040519261352584612fc9565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a01526135e9576135ba575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b819397508092945010611c3557505063ffffffff808061ffff9351169451169551169351169193929190600190565b50505050505092505050600090600090600090600090600090565b5082156134fc565b3567ffffffffffffffff811681036107b45790565b3573ffffffffffffffffffffffffffffffffffffffff811681036107b45790565b9067ffffffffffffffff61329092166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b8181029291811591840414171561369257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9190820391821161369257565b519073ffffffffffffffffffffffffffffffffffffffff821682036107b457565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b929593909473ffffffffffffffffffffffffffffffffffffffff600354169586156139b55760028610156139865773ffffffffffffffffffffffffffffffffffffffff976137dc9461ffff9288156138a4575b67ffffffffffffffff906040519b8c997f89720a62000000000000000000000000000000000000000000000000000000008b521660048a0152166024880152604487015216606485015260c0608485015260c48401916136ef565b928180600095869560a483015203915afa9182156138975781926137ff57505090565b9091503d8083833e613811818361301d565b8101906020818303126107755780519067ffffffffffffffff8211610958570181601f8201121561077557805190613848826132d8565b93613856604051958661301d565b82855260208086019360051b8301019384116102865750602001905b82821061387f5750505090565b6020809161388c846136ce565b815201910190613872565b50604051903d90823e3d90fd5b67ffffffffffffffff8116600052600b6020526040600020604051906138c982612fc9565b549063ffffffff8216815263ffffffff8260201c16602082015263ffffffff8260401c16604082015263ffffffff8260601c16606082015260c0868360801c169182608082015260ff60a0820194898160901c16865260a01c1615159182910152613936575b5050613781565b919267ffffffffffffffff928587161561396e575061271061395e876139659451168361367f565b04906136c1565b915b903861392f565b613980925061395e612710918361367f565b91613967565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b50505050505050506040516139cb60208261301d565b60008152600036813790565b604051906139e482612fad565b60606020838281520152565b8051821015613a045760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015613a7c575b6020831014613a4d57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613a42565b9060405191826000825492613a9a84613a33565b8084529360018116908115613b085750600114613ac1575b50613abf9250038361301d565b565b90506000929192526020600020906000915b818310613aec575050906020613abf9282010138613ab2565b6020919350806001915483858901015201910190918492613ad3565b60209350613abf9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613ab2565b67ffffffffffffffff1660005260086020526132906004604060002001613a86565b9190811015613a045760081b0190565b3580151581036107b45790565b3561ffff811681036107b45790565b3563ffffffff811681036107b45790565b359063ffffffff821682036107b457565b9190811015613a045760051b0190565b35906fffffffffffffffffffffffffffffffff821682036107b457565b91908260609103126107b457604051613bfd81612fe5565b6040613c27818395613c0e81613154565b8552613c1c60208201613bc8565b602086015201613bc8565b910152565b6fffffffffffffffffffffffffffffffff613c6a60408093613c4d81613154565b1515865283613c5e60208301613bc8565b16602087015201613bc8565b16910152565b818110613c7b575050565b60008155600101613c70565b613c8f6133cd565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691613cec6020850193613ce6613cd963ffffffff875116426136c1565b856080890151169061367f565b9061425a565b80821015613d0557505b16825263ffffffff4216905290565b9050613cf6565b908160209103126107b4575180151581036107b45790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156107b4570180359067ffffffffffffffff82116107b4576020019181360383136107b457565b73ffffffffffffffffffffffffffffffffffffffff600154163303613d9657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156140005767ffffffffffffffff81516020830120921691826000526008602052613df5816005604060002001614680565b15613fbc5760005260096020526040600020815167ffffffffffffffff8111612f7e57613e228254613a33565b601f8111613f8a575b506020601f8211600114613ec45791613e9e827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593613eb495600091613eb9575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b905560405191829160208352602083019061305e565b0390a2565b905084015138613e6d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110613f72575092613eb49492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610613f3b575b5050811b0190556118c3565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880613f2f565b9192602060018192868a015181550194019201613ef4565b613fb690836000526020600020601f840160051c810191602085106106df57601f0160051c0190613c70565b38613e2b565b50906120496040519283927f393b8ad2000000000000000000000000000000000000000000000000000000008452600484015260406024840152604483019061305e565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b8151919291156141ac576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061414957613abf91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836141aa604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040840151161580159061423b575b6141da57613abf919261406d565b6064836141aa604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208401511615156141cc565b9190820180921161369257565b906040519182815491828252602082019060005260206000209260005b818110614299575050613abf9250038361301d565b8454835260019485019487945060209093019201614284565b67ffffffffffffffff166142d3816000526007602052604060002054151590565b1561431d575033600052600e602052604060002054156142ef57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8054821015613a045760005260206000200190600090565b805480156143c6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190614397828261434a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008181526007602052604090205480156144fa577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161369257600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116136925781810361448b575b5050506144776006614362565b600052600760205260006040812055600190565b6144e261449c6144ad93600661434a565b90549060031b1c928392600661434a565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600760205260406000205538808061446a565b5050600090565b9060018201918160005282602052604060002054908115156000146145d4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918083116136925781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811161369257838161458b950361459d575b505050614362565b60005260205260006040812055600190565b6145bd6145ad6144ad938661434a565b90549060031b1c9283928661434a565b905560005284602052604060002055388080614583565b50505050600090565b80549068010000000000000000821015612f7e57816144ad9160016146049401815561434a565b9055565b80600052600e602052604060002054156000146146415761462a81600d6145dd565b600d5490600052600e602052604060002055600190565b50600090565b80600052600760205260406000205415600014614641576146698160066145dd565b600654906000526007602052604060002055600190565b60008281526001820160205260409020546144fa57806146a2836001936145dd565b80549260005201602052604060002055600190565b9182549060ff8260a01c16158015614964575b61495e576fffffffffffffffffffffffffffffffff8216916001850190815461470f63ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426136c1565b90816148c0575b505084811061487457508383106147705750506147456fffffffffffffffffffffffffffffffff9283926136c1565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c9283156148085781614788916136c1565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101938185116136925773ffffffffffffffffffffffffffffffffffffffff946147d39161425a565b047fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611614934576148db92613ce69160801c9061367f565b8084101561492f5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880614716565b6148e6565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156146ca565b6000818152600e602052604090205480156144fa577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161369257600d54907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161369257808203614a02575b5050506149ee600d614362565b600052600e60205260006040812055600190565b614a24614a136144ad93600d61434a565b90549060031b1c928392600d61434a565b9055600052600e6020526040600020553880806149e156fea164736f6c634300081a000a",
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmations)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector, customBlockConfirmations)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector, customBlockConfirmations)
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
	outstruct.FeeAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)

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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetFee(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetFee(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationsRequested, arg5)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetMinBlockConfirmations(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getMinBlockConfirmations")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetMinBlockConfirmations() (uint16, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetMinBlockConfirmations(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetMinBlockConfirmations() (uint16, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetMinBlockConfirmations(&_CCTPThroughCCVTokenPool.CallOpts)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationsRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRequiredCCVs(&_CCTPThroughCCVTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationsRequested, extraData, direction)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationsRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRequiredCCVs(&_CCTPThroughCCVTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationsRequested, extraData, direction)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.LockOrBurn0(&_CCTPThroughCCVTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.LockOrBurn0(&_CCTPThroughCCVTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationsRequested, tokenArgs)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationsRequested)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint0(&_CCTPThroughCCVTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint0(&_CCTPThroughCCVTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationsRequested)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAdmin)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetDynamicConfig(&_CCTPThroughCCVTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetDynamicConfig(&_CCTPThroughCCVTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "setMinBlockConfirmations", minBlockConfirmations)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetMinBlockConfirmations(&_CCTPThroughCCVTokenPool.TransactOpts, minBlockConfirmations)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) SetMinBlockConfirmations(minBlockConfirmations uint16) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetMinBlockConfirmations(&_CCTPThroughCCVTokenPool.TransactOpts, minBlockConfirmations)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.WithdrawFeeTokens(&_CCTPThroughCCVTokenPool.TransactOpts, feeTokens, recipient)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.WithdrawFeeTokens(&_CCTPThroughCCVTokenPool.TransactOpts, feeTokens, recipient)
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

type CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator struct {
	Event *CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
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
		it.Event = new(CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
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

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationsInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "CustomBlockConfirmationsInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationsInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsInboundRateLimitConsumed", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseCustomBlockConfirmationsInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, error) {
	event := new(CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator struct {
	Event *CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
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
		it.Event = new(CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
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

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationsOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "CustomBlockConfirmationsOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationsOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsOutboundRateLimitConsumed", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseCustomBlockConfirmationsOutboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, error) {
	event := new(CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationsOutboundRateLimitConsumed", log); err != nil {
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
	FeeAdmin       common.Address
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

type CCTPThroughCCVTokenPoolMinBlockConfirmationsSetIterator struct {
	Event *CCTPThroughCCVTokenPoolMinBlockConfirmationsSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolMinBlockConfirmationsSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolMinBlockConfirmationsSet)
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
		it.Event = new(CCTPThroughCCVTokenPoolMinBlockConfirmationsSet)
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

func (it *CCTPThroughCCVTokenPoolMinBlockConfirmationsSetIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolMinBlockConfirmationsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolMinBlockConfirmationsSet struct {
	MinBlockConfirmations uint16
	Raw                   types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolMinBlockConfirmationsSetIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolMinBlockConfirmationsSetIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "MinBlockConfirmationsSet", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolMinBlockConfirmationsSet) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolMinBlockConfirmationsSet)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseMinBlockConfirmationsSet(log types.Log) (*CCTPThroughCCVTokenPoolMinBlockConfirmationsSet, error) {
	event := new(CCTPThroughCCVTokenPoolMinBlockConfirmationsSet)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "MinBlockConfirmationsSet", log); err != nil {
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
	CustomBlockConfirmations  bool
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
	FeeAdmin       common.Address
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

func (CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x63335ad9e238acd0e9e6c1c20f529ffbea4cda73578c329a7aa7e9d61e5cdcc5")
}

func (CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x996b829383cc7e136842d4c4c175083bcf4e20807c7432105c1b794ba885e776")
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

func (CCTPThroughCCVTokenPoolMinBlockConfirmationsSet) Topic() common.Hash {
	return common.HexToHash("0x46c9c0585a955b2702c7ea47fec541db623565d20827a0edda42864e6b859a01")
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

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmations bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationsRequested uint16, arg5 []byte) (GetFee,

		error)

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

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, arg0 PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationsRequested uint16, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, arg0 PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationsRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmations(opts *bind.TransactOpts, minBlockConfirmations uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

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

	FilterCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationsInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationsInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationsInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationsOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationsOutboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolCustomBlockConfirmationsOutboundRateLimitConsumed, error)

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

	FilterMinBlockConfirmationsSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolMinBlockConfirmationsSetIterator, error)

	WatchMinBlockConfirmationsSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolMinBlockConfirmationsSet) (event.Subscription, error)

	ParseMinBlockConfirmationsSet(log types.Log) (*CCTPThroughCCVTokenPoolMinBlockConfirmationsSet, error)

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
