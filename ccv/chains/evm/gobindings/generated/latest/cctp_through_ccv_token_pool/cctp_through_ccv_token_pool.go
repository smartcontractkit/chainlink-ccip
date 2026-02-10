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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCTPVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CCVNotSetOnResolver\",\"inputs\":[{\"name\":\"resolver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"IPoolV1NotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461037457614f1c803803809161001d82856103dd565b833981019060c0818303126103745780516001600160a01b038116908181036103745761004c60208401610400565b906100596040850161040e565b906100666060860161040e565b936100736080870161040e565b60a087015190966001600160401b03821161037457019680601f89011215610374578751976001600160401b0389116102ff578860051b9060208201996100bd6040519b8c6103dd565b8a526020808b019282010192831161037457602001905b8282106103c55750505033156103b457600180546001600160a01b03191633179055801580156103a3575b8015610392575b6103815760049260209260805260c0526040519283809263313ce56760e01b82525afa60009181610340575b50610315575b5060a052600380546001600160a01b0319908116909155600280549091166001600160a01b039290921691909117905560405160209061017882826103dd565b60008152600036813760408051949085016001600160401b038111868210176102ff576040528452808285015260005b815181101561020f576001906001600160a01b036101c68285610422565b5116846101d282610464565b6101df575b5050016101a8565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138846101d7565b5050915160005b8151811015610287576001600160a01b036102318284610422565b5116908115610276577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8583610268600195610562565b50604051908152a101610216565b6342bcdf7f60e11b60005260046000fd5b6001600160a01b03831660e05260405161495990816105c382396080518181816118d101528181611ab80152818161276f0152818161295c01528181612d5b0152612db5015260a05181612bac015260c051818181610bb20152818161196c015261280a015260e05181818161114e01526125e90152f35b634e487b7160e01b600052604160045260246000fd5b60ff1660ff82168181036103295750610138565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610379575b8161035c602093836103dd565b810103126103745761036d90610400565b9038610132565b600080fd5b3d915061034f565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610106565b506001600160a01b038516156100ff565b639b15e16f60e01b60005260046000fd5b602080916103d28461040e565b8152019101906100d4565b601f909101601f19168101906001600160401b038211908210176102ff57604052565b519060ff8216820361037457565b51906001600160a01b038216820361037457565b80518210156104365760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104365760005260206000200190600090565b6000818152600e6020526040902054801561055b57600019810181811161054557600d54600019810191908211610545578082036104f4575b505050600d5480156104de57600019016104b881600d61044c565b8154906000199060031b1b19169055600d55600052600e60205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61052d61050561051693600d61044c565b90549060031b1c928392600d61044c565b819391549060031b91821b91600019901b19161790565b9055600052600e60205260406000205538808061049d565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600e602052604060002054156000146105bc57600d54680100000000000000008110156102ff576105a3610516826001859401600d55600d61044c565b9055600d5490600052600e602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714612e6057508063181f5a7714612dd957806321df0da714612d88578063240028e814612d245780632422ac4514612c455780632451a62714612bd057806324f65ee714612b925780632c06340414612af957806337a3210d14612ac55780633907753714612a5c578063489a68f2146126f35780634c5ef0ed146126ac5780634e921c301461260d578063615521a7146125bc57806362ddd3c4146125355780637437ff9f146124e757806379ba5097146124205780638926f54f146123da57806389720a621461235a5780638da5cb5b1461232657806391a2749a146121785780639a4575b914612115578063a42a7b8b14611fae578063acfecf9114611eb6578063ae39a25714611d2b578063b1c71c651461183c578063b7946580146117ff578063bfeffd3f14611753578063c4bffe2b14611628578063c7230a6014611375578063d8aa3f4014611027578063dc04fa1f14610bd6578063dc0bd97114610b85578063dcbd41bc14610981578063e8a1da17146102ae578063f2fde38b146101df5763fa41d79c146101b857600080fd5b346101dc57806003193601126101dc57602061ffff60025460a01c16604051908152f35b80fd5b50346101dc5760206003193601126101dc5773ffffffffffffffffffffffffffffffffffffffff61020e6130bd565b610216613c85565b1633811461028657807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346101dc5760406003193601126101dc5760043567ffffffffffffffff81116107d5576102e090369060040161336b565b9060243567ffffffffffffffff811161097d57906103038492369060040161336b565b93909161030e613c85565b83905b8282106107e25750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156107de578060051b830135858112156107d1578301610120813603126107d1576040519461037586613001565b813567ffffffffffffffff811681036107d9578652602082013567ffffffffffffffff81116107d55782019436601f870112156107d5578535956103b8876132d8565b966103c6604051988961301d565b80885260208089019160051b830101903682116107d15760208301905b82821061079e575050505060208701958652604083013567ffffffffffffffff811161079a576104169036908501613275565b916040880192835261044061042e3660608701613aa6565b9460608a0195865260c0369101613aa6565b9560808901968752835151156107725761046467ffffffffffffffff8a5116614557565b1561073b5767ffffffffffffffff895116825260086020526040822061048b865182613f3a565b610499885160028301613f3a565b6004855191019080519067ffffffffffffffff821161070e576104bc83546138f4565b601f81116106d3575b50602090601f8311600114610634576105139291869183610629575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b8851805182101561054d57906105476001926105408367ffffffffffffffff8f5116926138b1565b5190613cd0565b01610518565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561061b67ffffffffffffffff60019796949851169251935191516105e76105b26040519687968752610100602088015261010087019061305e565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610344565b015190508e806104e1565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106106bb5750908460019594939210610684575b505050811b019055610516565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080610677565b92936020600181928786015181550195019301610661565b6106fe9084875260208720601f850160051c81019160208610610704575b601f0160051c0190613b31565b8d6104c5565b90915081906106f1565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116107cd576020916107c28392833691890101613275565b8152019101906103e3565b8680fd5b8480fd5b5080fd5b600080fd5b8380f35b9267ffffffffffffffff6108046107ff8486889a9699979a613a79565b61360c565b169161080f83614305565b1561095157828452600860205261082b60056040862001614177565b94845b865181101561086457600190858752600860205261085d60056040892001610856838b6138b1565b5190614411565b500161082e565b50939692909450949094808752600860205260056040882088815588600182015588600282015588600382015588600482016108a081546138f4565b80610910575b50505001805490888155816108f2575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101909194939294610311565b885260208820908101905b818110156108b6578881556001016108fd565b601f81116001146109265750555b888a806108a6565b8183526020832061094191601f01861c810190600101613b31565b808252816020812091555561091e565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346101dc5760206003193601126101dc5760043567ffffffffffffffff81116107d5576109b390369060040161339c565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580610b63575b610b3757825b8181106109e6578380f35b6109f1818385613a2b565b67ffffffffffffffff610a038261360c565b1690610a1c826000526007602052604060002054151590565b15610b0b57907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610acb610aa5602060019897018b610a5d82613a3b565b15610ad2578790526004602052610a8460408d20610a7e3660408801613aa6565b90613f3a565b868c526005602052610aa060408d20610a7e3660a08801613aa6565b613a3b565b916040519215158352610abe6020840160408301613aed565b60a0608084019101613aed565ba2016109db565b60026040828a610aa094526008602052610af4828220610a7e36858c01613aa6565b8a815260086020522001610a7e3660a08801613aa6565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600154163314156109d5565b50346101dc57806003193601126101dc57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101dc5760406003193601126101dc5760043567ffffffffffffffff81116107d557610c0890369060040161339c565b60243567ffffffffffffffff811161097d57610c2890369060040161336b565b919092610c33613c85565b845b828110610c9f57505050825b818110610c4c578380f35b8067ffffffffffffffff610c666107ff6001948688613a79565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610c41565b610cad6107ff828585613a2b565b610cb8828585613a2b565b90602082019060e0830190610ccc82613a3b565b15610ff25760a0840161271061ffff610ce483613a48565b161015610fe35760c085019161271061ffff610cff85613a48565b161015610fab5763ffffffff610d1486613a57565b1615610f765767ffffffffffffffff1694858c52600b60205260408c20610d3a86613a57565b63ffffffff16908054906040840191610d5283613a57565b60201b67ffffffff0000000016936060860194610d6e86613a57565b60401b6bffffffff0000000000000000169660800196610d8d88613a57565b60601b6fffffffff0000000000000000000000001691610dac8a613a48565b60801b71ffff000000000000000000000000000000001693610dcd8c613a48565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610e8087613a3b565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610ed190613a68565b63ffffffff168752610ee290613a68565b63ffffffff166020870152610ef690613a68565b63ffffffff166040860152610f0a90613a68565b63ffffffff166060850152610f1e906131d3565b61ffff166080840152610f30906131d3565b61ffff1660a0830152610f4290613154565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c35565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff610fba86613a48565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff610fba602493613a48565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346101dc5760806003193601126101dc576110416130bd565b5061104a61313d565b6110526131c2565b506064359067ffffffffffffffff821161079a5761107d67ffffffffffffffff9236906004016131e2565b50508260c060405161108e81612fc9565b8281528260208201528260408201528260608201528260808201528260a0820152015216808252600b6020526040822090604051916110cc83612fc9565b549063ffffffff82168352602083019363ffffffff8360201c168552604084019463ffffffff8460401c168652606085019063ffffffff8560601c168252608086019261ffff8660801c16845260a087019461ffff8760901c16865260ff60c089019760a01c161515875273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690604051907f958021a7000000000000000000000000000000000000000000000000000000008252600482015260406024820152826044820152602081606481855afa801561136a57839061131d575b73ffffffffffffffffffffffffffffffffffffffff9150169081156112f257506060600491604051928380927f7437ff9f0000000000000000000000000000000000000000000000000000000082525afa9182156112e6578092611271575b505061ffff9493919263ffffffff60e09981889687604083970151168952816040519c51168c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b9091506060823d6060116112de575b8161128d6060938361301d565b810103126101dc5760408051926112a384612fe5565b6112ac8161367f565b84526112ba6020820161367f565b602085015201519061ffff821682036101dc575060408201528263ffffffff61121d565b3d9150611280565b604051903d90823e3d90fd5b7f4172d660000000000000000000000000000000000000000000000000000000008352600452602482fd5b506020813d602011611362575b816113376020938361301d565b8101031261079a5761135d73ffffffffffffffffffffffffffffffffffffffff9161367f565b6111be565b3d915061132a565b6040513d85823e3d90fd5b50346101dc5760406003193601126101dc5760043567ffffffffffffffff81116107d5576113a790369060040161336b565b906113b0613103565b9173ffffffffffffffffffffffffffffffffffffffff6001541633141580611606575b6115da5773ffffffffffffffffffffffffffffffffffffffff83169081156115b257845b818110611402578580f35b73ffffffffffffffffffffffffffffffffffffffff61142a611425838588613a79565b613621565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa9081156115a7578891611572575b508061147f575b50506001016113f7565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff8a16602484015260448084018590528352918a91906114e060648261301d565b519082865af1156115675787513d61155e5750813b155b6115325790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a39038611475565b602488837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b600114156114f7565b6040513d89823e3d90fd5b90506020813d821161159f575b8161158c6020938361301d565b8101031261159b57513861146e565b8780fd5b3d915061157f565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600c54163314156113d3565b50346101dc57806003193601126101dc57604051906006548083528260208101600684526020842092845b81811061173a5750506116689250038361301d565b815161168c611676826132d8565b91611684604051938461301d565b8083526132d8565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b84518110156116eb578067ffffffffffffffff6116d8600193886138b1565b51166116e482866138b1565b52016116b9565b50925090604051928392602084019060208552518091526040840192915b818110611717575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611709565b8454835260019485019487945060209093019201611653565b50346101dc5760206003193601126101dc5760043573ffffffffffffffffffffffffffffffffffffffff81168091036107d55761178e613c85565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a1161760035580f35b50346101dc5760206003193601126101dc5761183861182461181f613126565b613a09565b60405191829160208352602083019061305e565b0390f35b50346101dc5760606003193601126101dc576004359067ffffffffffffffff82116101dc5760a060031983360301126101dc576118776131b1565b60443567ffffffffffffffff811161079a5761189a6118aa9136906004016131e2565b6118a2613898565b503691613210565b5060848301906118b982613621565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611ce157602484019377ffffffffffffffff0000000000000000000000000000000061191f8661360c565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611cd6578591611ca7575b50611c7f5790606461ffff926119b96119b48861360c565b6141c2565b01359350168015611bf05761ffff60025460a01c16908115611bc657818110611b9657505061181f611b5b936119f1611aee93613621565b7f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac209793288567ffffffffffffffff611a258561360c565b1692836000526004602052611a3f818360406000206145c7565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b611a718161360c565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10606067ffffffffffffffff6040519373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001685523360208601528860408601521692a261360c565b90611b8c6040517f3047587c00000000000000000000000000000000000000000000000000000000602082015260048152611b2a60248261301d565b60405193611b3785612fad565b8452602084019081526040519485946040865251604080870152608086019061305e565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc085830301606086015261305e565b9060208301520390f35b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b5061181f611b5b93611c04611aee93613621565b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448567ffffffffffffffff611c388561360c565b1692836000526008602052611c52818360406000206145c7565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611a68565b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611cc9915060203d602011611ccf575b611cc1818361301d565b810190613c1c565b3861199c565b503d611cb7565b6040513d87823e3d90fd5b60248373ffffffffffffffffffffffffffffffffffffffff611d0285613621565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346101dc5760606003193601126101dc57611d456130bd565b90611d4e613103565b6044359273ffffffffffffffffffffffffffffffffffffffff841680850361097d57611d78613c85565b73ffffffffffffffffffffffffffffffffffffffff82168015611e8e5794611e88917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c556040519384938491604091949373ffffffffffffffffffffffffffffffffffffffff809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346101dc5767ffffffffffffffff611ece36613293565b929091611ed9613c85565b1691611ef2836000526007602052604060002054151590565b15610951578284526008602052611f2160056040862001611f14368486613210565b6020815191012090614411565b15611f6657907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f606040519283926020845260208401916136a0565b0390a280f35b82611faa836040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916136a0565b0390fd5b50346101dc5760206003193601126101dc5767ffffffffffffffff611fd1613126565b1681526008602052611fe860056040832001614177565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061202d612017836132d8565b92612025604051948561301d565b8084526132d8565b01835b818110612104575050825b82518110156120815780612051600192856138b1565b518552600960205261206560408620613947565b61206f82856138b1565b5261207a81846138b1565b500161203b565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106120b957505050500390f35b919360206120f4827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652885161305e565b96019201920185949391926120aa565b806060602080938601015201612030565b50346101dc5760206003193601126101dc5760043567ffffffffffffffff81116107d55760031960a091360301126101dc57600490612152613898565b507f690a7a40000000000000000000000000000000000000000000000000000000008152fd5b50346101dc5760206003193601126101dc5760043567ffffffffffffffff81116107d557604060031982360301126107d557604051906121b782612fad565b806004013567ffffffffffffffff811161097d576121db90600436918401016132f0565b825260248101359067ffffffffffffffff821161097d57600461220192369201016132f0565b60208201908152612210613c85565b5191805b8351811015612287578073ffffffffffffffffffffffffffffffffffffffff61223f600193876138b1565b511661224a8161487c565b612256575b5001612214565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a13861224f565b509051815b81518110156123225773ffffffffffffffffffffffffffffffffffffffff6122b482846138b1565b511680156122fa57907feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef6020836122ec600195614518565b50604051908152a10161228c565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8280f35b50346101dc57806003193601126101dc57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346101dc5760c06003193601126101dc576123746130bd565b61237c61313d565b60643561ffff8116810361097d5760843567ffffffffffffffff81116107d1576123aa9036906004016131e2565b93909260a4359560028710156101dc576118386123ce88888888604435888a6136df565b60405191829182613161565b50346101dc5760206003193601126101dc57602061241667ffffffffffffffff612402613126565b166000526007602052604060002054151590565b6040519015158152f35b50346101dc57806003193601126101dc57805473ffffffffffffffffffffffffffffffffffffffff811633036124bf577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346101dc57806003193601126101dc57600254600a54600c546040805173ffffffffffffffffffffffffffffffffffffffff94851681529284166020840152921691810191909152606090f35b50346101dc5761254436613293565b61255093929193613c85565b67ffffffffffffffff8216612572816000526007602052604060002054151590565b15612591575061258e9293612588913691613210565b90613cd0565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346101dc57806003193601126101dc57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101dc5760206003193601126101dc5760043561ffff81169081810361079a577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb9160209161265c613c85565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006002549260a01b16911617600255604051908152a180f35b50346101dc5760406003193601126101dc576126c6613126565b906024359067ffffffffffffffff82116101dc576020612416846126ed3660048701613275565b90613642565b50346101dc5760406003193601126101dc576004359067ffffffffffffffff82116101dc57816004019161010060031982360301126107d5576127346131b1565b918060405161274281612f62565b52606482013592608483019061275782613621565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611ce157602484019577ffffffffffffffff000000000000000000000000000000006127bd8861360c565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611cd6578591612a3d575b50611c7f5761284b6119b48861360c565b6128548761360c565b9061287160a48701926126ed61286a8585613c34565b3691613210565b156129f657505061293f60446129386020987ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966080968a61ffff67ffffffffffffffff981615156000146129a8576128ea7f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61592613621565b61290b818360408c6128fb8a61360c565b16808952600560205297206145c7565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a261360c565b9501613621565b9373ffffffffffffffffffffffffffffffffffffffff60405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a28060405161299f81612f62565b52604051908152f35b6129d27f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c92613621565b61290b8183600260408d6129e58b61360c565b16808a5260086020529820016145c7565b612a009250613c34565b611faa6040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916136a0565b612a56915060203d602011611ccf57611cc1818361301d565b3861283a565b50346101dc5760206003193601126101dc5760043567ffffffffffffffff81116107d55760031961010091360301126101dc5780600491604051612a9f81612f62565b527f690a7a40000000000000000000000000000000000000000000000000000000008152fd5b50346101dc57806003193601126101dc57602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b50346101dc5760c06003193601126101dc57612b136130bd565b50612b1c61313d565b612b246130e0565b506084359161ffff831683036101dc5760a4359067ffffffffffffffff82116101dc5760a063ffffffff8061ffff612b6b8888612b643660048b016131e2565b50506134e1565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346101dc57806003193601126101dc57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101dc57806003193601126101dc57604051600d8054808352908352909160208301917fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5915b818110612c2f57611838856123ce8187038261301d565b8254845260209093019260019283019201612c18565b50346101dc5760406003193601126101dc57612c5f613126565b6024359182151583036101dc57610140612d22612c7c858561345e565b612cd260409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346101dc5760206003193601126101dc57602090612d416130bd565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346101dc57806003193601126101dc57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346101dc57806003193601126101dc5750611838604051612dfc60608261301d565b602181527f434354505468726f756768434356546f6b656e506f6f6c20322e302e302d646560208201527f7600000000000000000000000000000000000000000000000000000000000000604082015260405191829160208352602083019061305e565b9050346107d55760206003193601126107d5576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361079a57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115612f38575b8115612f0e575b8115612ee4575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438612edd565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150612ed6565b7f331710310000000000000000000000000000000000000000000000000000000081149150612ecf565b6020810190811067ffffffffffffffff821117612f7e57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117612f7e57604052565b60e0810190811067ffffffffffffffff821117612f7e57604052565b6060810190811067ffffffffffffffff821117612f7e57604052565b60a0810190811067ffffffffffffffff821117612f7e57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117612f7e57604052565b919082519283825260005b8481106130a85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613069565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036107d957565b6064359073ffffffffffffffffffffffffffffffffffffffff821682036107d957565b6024359073ffffffffffffffffffffffffffffffffffffffff821682036107d957565b6004359067ffffffffffffffff821682036107d957565b6024359067ffffffffffffffff821682036107d957565b359081151582036107d957565b602060408183019282815284518094520192019060005b8181106131855750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613178565b6024359061ffff821682036107d957565b6044359061ffff821682036107d957565b359061ffff821682036107d957565b9181601f840112156107d95782359167ffffffffffffffff83116107d957602083818601950101116107d957565b92919267ffffffffffffffff8211612f7e5760405191613258601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0166020018461301d565b8294818452818301116107d9578281602093846000960137010152565b9080601f830112156107d95781602061329093359101613210565b90565b9060406003198301126107d95760043567ffffffffffffffff811681036107d957916024359067ffffffffffffffff82116107d9576132d4916004016131e2565b9091565b67ffffffffffffffff8111612f7e5760051b60200190565b9080601f830112156107d957813590613308826132d8565b92613316604051948561301d565b82845260208085019360051b8201019182116107d957602001915b81831061333e5750505090565b823573ffffffffffffffffffffffffffffffffffffffff811681036107d957815260209283019201613331565b9181601f840112156107d95782359167ffffffffffffffff83116107d9576020808501948460051b0101116107d957565b9181601f840112156107d95782359167ffffffffffffffff83116107d9576020808501948460081b0101116107d957565b604051906133da82613001565b60006080838281528260208201528260408201528260608201520152565b9060405161340581613001565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916134706133cd565b506134796133cd565b506134ad571660005260086020526040600020906132906134a160026134a66134a1866133f8565b613b97565b94016133f8565b16908160005260046020526134c86134a160406000206133f8565b9160005260056020526132906134a160406000206133f8565b9061ffff8060025460a01c1691169283151592838094613604575b611bc65767ffffffffffffffff16600052600b6020526040600020916040519261352584612fc9565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a01526135e9576135ba575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b819397508092945010611b9657505063ffffffff808061ffff9351169451169551169351169193929190600190565b50505050505092505050600090600090600090600090600090565b5082156134fc565b3567ffffffffffffffff811681036107d95790565b3573ffffffffffffffffffffffffffffffffffffffff811681036107d95790565b9067ffffffffffffffff61329092166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b519073ffffffffffffffffffffffffffffffffffffffff821682036107d957565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff600354169586156138765761377a9467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c48601916136a0565b916002821015613847578380600094819460a483015203915afa90811561383b576000916137a6575090565b3d8083833e6137b5818361301d565b81019060208183031261079a5780519067ffffffffffffffff821161097d570181601f8201121561079a578051906137ec826132d8565b936137fa604051958661301d565b82855260208086019360051b8301019384116101dc5750602001905b8282106138235750505090565b602080916138308461367f565b815201910190613816565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b505050505050505060405161388c60208261301d565b60008152600036813790565b604051906138a582612fad565b60606020838281520152565b80518210156138c55760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c9216801561393d575b602083101461390e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613903565b906040519182600082549261395b846138f4565b80845293600181169081156139c95750600114613982575b506139809250038361301d565b565b90506000929192526020600020906000915b8183106139ad5750509060206139809282010138613973565b6020919350806001915483858901015201910190918492613994565b602093506139809592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613973565b67ffffffffffffffff1660005260086020526132906004604060002001613947565b91908110156138c55760081b0190565b3580151581036107d95790565b3561ffff811681036107d95790565b3563ffffffff811681036107d95790565b359063ffffffff821682036107d957565b91908110156138c55760051b0190565b35906fffffffffffffffffffffffffffffffff821682036107d957565b91908260609103126107d957604051613abe81612fe5565b6040613ae8818395613acf81613154565b8552613add60208201613a89565b602086015201613a89565b910152565b6fffffffffffffffffffffffffffffffff613b2b60408093613b0e81613154565b1515865283613b1f60208301613a89565b16602087015201613a89565b16910152565b818110613b3c575050565b60008155600101613b31565b81810292918115918404141715613b5b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211613b5b57565b613b9f6133cd565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691613bfc6020850193613bf6613be963ffffffff87511642613b8a565b8560808901511690613b48565b9061416a565b80821015613c1557505b16825263ffffffff4216905290565b9050613c06565b908160209103126107d9575180151581036107d95790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156107d9570180359067ffffffffffffffff82116107d9576020019181360383136107d957565b73ffffffffffffffffffffffffffffffffffffffff600154163303613ca657565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115613f105767ffffffffffffffff81516020830120921691826000526008602052613d05816005604060002001614590565b15613ecc5760005260096020526040600020815167ffffffffffffffff8111612f7e57613d3282546138f4565b601f8111613e9a575b506020601f8211600114613dd45791613dae827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593613dc495600091613dc9575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b905560405191829160208352602083019061305e565b0390a2565b905084015138613d7d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110613e82575092613dc49492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610613e4b575b5050811b019055611824565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880613e3f565b9192602060018192868a015181550194019201613e04565b613ec690836000526020600020601f840160051c8101916020851061070457601f0160051c0190613b31565b38613d3b565b5090611faa6040519283927f393b8ad2000000000000000000000000000000000000000000000000000000008452600484015260406024840152604483019061305e565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b8151919291156140bc576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff602085015116106140595761398091925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836140ba604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040840151161580159061414b575b6140ea576139809192613f7d565b6064836140ba604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208401511615156140dc565b91908201809211613b5b57565b906040519182815491828252602082019060005260206000209260005b8181106141a95750506139809250038361301d565b8454835260019485019487945060209093019201614194565b67ffffffffffffffff166141e3816000526007602052604060002054151590565b1561422d575033600052600e602052604060002054156141ff57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80548210156138c55760005260206000200190600090565b805480156142d6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906142a7828261425a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260076020526040902054801561440a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613b5b57600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613b5b5781810361439b575b5050506143876006614272565b600052600760205260006040812055600190565b6143f26143ac6143bd93600661425a565b90549060031b1c928392600661425a565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600760205260406000205538808061437a565b5050600090565b9060018201918160005282602052604060002054908115156000146144e4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191808311613b5b5781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111613b5b57838161449b95036144ad575b505050614272565b60005260205260006040812055600190565b6144cd6144bd6143bd938661425a565b90549060031b1c9283928661425a565b905560005284602052604060002055388080614493565b50505050600090565b80549068010000000000000000821015612f7e57816143bd9160016145149401815561425a565b9055565b80600052600e602052604060002054156000146145515761453a81600d6144ed565b600d5490600052600e602052604060002055600190565b50600090565b80600052600760205260406000205415600014614551576145798160066144ed565b600654906000526007602052604060002055600190565b600082815260018201602052604090205461440a57806145b2836001936144ed565b80549260005201602052604060002055600190565b9182549060ff8260a01c16158015614874575b61486e576fffffffffffffffffffffffffffffffff8216916001850190815461461f63ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642613b8a565b90816147d0575b505084811061478457508383106146805750506146556fffffffffffffffffffffffffffffffff928392613b8a565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315614718578161469891613b8a565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810193818511613b5b5773ffffffffffffffffffffffffffffffffffffffff946146e39161416a565b7fd0c8d23a00000000000000000000000000000000000000000000000000000000600052046004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611614844576147eb92613bf69160801c90613b48565b8084101561483f5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880614626565b6147f6565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156145da565b6000818152600e6020526040902054801561440a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613b5b57600d54907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613b5b57808203614912575b5050506148fe600d614272565b600052600e60205260006040812055600190565b6149346149236143bd93600d61425a565b90549060031b1c928392600d61425a565b9055600052600e6020526040600020553880806148f156fea164736f6c634300081a000a",
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin, feeAdmin)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetDynamicConfig(&_CCTPThroughCCVTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetDynamicConfig(&_CCTPThroughCCVTokenPool.TransactOpts, router, rateLimitAdmin, feeAdmin)
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

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

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
