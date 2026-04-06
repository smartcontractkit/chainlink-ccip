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
	DestGasOverhead            uint32
	DestBytesOverhead          uint32
	FinalityFeeUSDCents        uint32
	FastFinalityFeeUSDCents    uint32
	FinalityTransferFeeBps     uint16
	FastFinalityTransferFeeBps uint16
	IsEnabled                  bool
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
	FastFinality              bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var CCTPThroughCCVTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowedFinalityConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCTPVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"requestedFinalityConfig\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllowedFinalityConfig\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFinalityInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FastFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigSet\",\"inputs\":[{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"fastFinality\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fastFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"finalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"fastFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CCVNotSetOnResolver\",\"inputs\":[{\"name\":\"resolver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotOwnerOrFeeAdmin\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"IPoolV1NotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRequestedFinality\",\"inputs\":[{\"name\":\"requestedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"RequestedFinalityCanOnlyHaveOneMode\",\"inputs\":[{\"name\":\"encodedFinality\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610100806040523461039e576155fa803803809161001d82856103f6565b833981019060c08183031261039e5780516001600160a01b03811680820361039e5761004b60208401610419565b9061005860408501610427565b61006460608601610427565b9361007160808701610427565b60a087015190966001600160401b03821161039e57019680601f8901121561039e578751976001600160401b038911610302578860051b9060208201996100bb6040519b8c6103f6565b8a526020808b019282010192831161039e57602001905b8282106103de5750505033156103cd57600180546001600160a01b03191633179055821580156103bc575b80156103ab575b6102f15760805260c052308103610318575b5060a052600380546001600160a01b0319908116909155600280549091166001600160a01b039290921691909117905560405160209061015682826103f6565b60008152600036813760408051949085016001600160401b03811186821017610302576040528452808285015260005b81518110156101ed576001906001600160a01b036101a4828561043b565b5116846101b08261047d565b6101bd575b505001610186565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138846101b5565b5050915160005b8151811015610265576001600160a01b0361020f828461043b565b5116908115610254577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef858361024660019561057b565b50604051908152a1016101f4565b6342bcdf7f60e11b60005260046000fd5b6001600160a01b03831680156102f15760e05260405161501e90816105dc823960805181818161024b015281816104060152818161271d015281816128f901528181612bf00152612c4a015260a051818181612a3501528181613da00152613dea015260c0518181816102e60152818161133c01526127b8015260e0518181816107fa015261259a0152f35b630a64406560e11b60005260046000fd5b634e487b7160e01b600052604160045260246000fd5b60206004916040519283809263313ce56760e01b82525afa6000918161036a575b50156101165760ff1660ff82168181036103535750610116565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103a3575b81610386602093836103f6565b8101031261039e5761039790610419565b9038610339565b600080fd5b3d9150610379565b506001600160a01b03821615610104565b506001600160a01b038516156100fd565b639b15e16f60e01b60005260046000fd5b602080916103eb84610427565b8152019101906100d2565b601f909101601f19168101906001600160401b0382119082101761030257604052565b519060ff8216820361039e57565b51906001600160a01b038216820361039e57565b805182101561044f5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561044f5760005260206000200190600090565b6000818152600e6020526040902054801561057457600019810181811161055e57600d5460001981019190821161055e5780820361050d575b505050600d5480156104f757600019016104d181600d610465565b8154906000199060031b1b19169055600d55600052600e60205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61054661051e61052f93600d610465565b90549060031b1c928392600d610465565b819391549060031b91821b91600019901b19161790565b9055600052600e6020526040600020553880806104b6565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600e602052604060002054156000146105d557600d5468010000000000000000811015610302576105bc61052f826001859401600d55600d610465565b9055600d5490600052600e602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714612e185750806306b859ef14612d86578063181f5a7714612d255780631826b1e714612c6e57806321df0da714612c1d578063240028e814612bb95780632422ac4514612ada5780632451a62714612a5957806324f65ee714612a1b5780632cab0fb6146126a257806337a3210d1461266e57806339077537146126055780634c5ef0ed146125be578063615521a71461256d57806362ddd3c4146124e65780637437ff9f1461249857806379ba5097146123d15780638926f54f1461238b5780638da5cb5b1461235757806391a2749a146121a95780639a4575b914612146578063a42a7b8b14611fdf578063acfecf9114611ee7578063ae39a25714611d5c578063b6cfa3b714611ca1578063b794658014611c69578063bfeffd3f14611bbd578063c4bffe2b14611a92578063c7230a60146117e4578063dc04fa1f14611360578063dc0bd9711461130f578063dcbd41bc1461110b578063e8a1da1714610a25578063ea6396db146106d3578063ec6ae7a714610690578063f2fde38b146105c15763fbc801a7146101b857600080fd5b346105be5760606003193601126105be5760043567ffffffffffffffff81116105ba5760a060031982360301126105ba576101f1612f4a565b9060443567ffffffffffffffff81116105b65761021561022591369060040161303f565b61021d613a6d565b503691613225565b506084810161023381613a0f565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361056c57602482019177ffffffffffffffff00000000000000000000000000000000610299846139fa565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115610561578691610532575b5061050a579161043c916104ae9560646104419561033b610336866139fa565b6146e1565b0135958691507fffffffff000000000000000000000000000000000000000000000000000000008116156104e9576103b6926103a26103a7927fffffffff0000000000000000000000000000000000000000000000000000000060025460401b1690613ef4565b613a0f565b6103b0846139fa565b90614c15565b6103bf816139fa565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10606067ffffffffffffffff6040519373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001685523360208601528860408601521692a26139fa565b613bde565b906104df6040517f3047587c0000000000000000000000000000000000000000000000000000000060208201526004815261047d602482613178565b6040519361048a85613108565b845260208401908152604051948594604086525160408087015260808601906131b9565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08583030160608601526131b9565b9060208301520390f35b506104f661050592613a0f565b6104ff846139fa565b90614bcf565b6103b6565b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b610554915060203d60201161055a575b61054c8183613178565b810190614060565b38610316565b503d610542565b6040513d88823e3d90fd5b8373ffffffffffffffffffffffffffffffffffffffff61058d602493613a0f565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b8380fd5b5080fd5b80fd5b50346105be5760206003193601126105be5773ffffffffffffffffffffffffffffffffffffffff6105f0612fa8565b6105f86140c9565b1633811461066857807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346105be57806003193601126105be5760207fffffffff0000000000000000000000000000000000000000000000000000000060025460401b16604051908152f35b50346105be5760806003193601126105be576106ed612fa8565b506106f6613011565b6106fe612f79565b506064359067ffffffffffffffff8211610a0e5761072967ffffffffffffffff92369060040161303f565b50508260c060405161073a81613140565b8281528260208201528260408201528260608201528260808201528260a0820152015216808252600b60205260408220906040519161077883613140565b549063ffffffff82168352602083019363ffffffff8360201c168552604084019463ffffffff8460401c168652606085019063ffffffff8560601c168252608086019261ffff8660801c16845260a087019461ffff8760901c16865260ff60c089019760a01c161515875273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690604051907f958021a7000000000000000000000000000000000000000000000000000000008252600482015260406024820152826044820152602081606481855afa8015610a1a5783906109c9575b73ffffffffffffffffffffffffffffffffffffffff91501690811561099e57506060600491604051928380927f7437ff9f0000000000000000000000000000000000000000000000000000000082525afa91821561099257809261091d575b505061ffff9493919263ffffffff60e09981889687604083970151168952816040519c51168c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b9091506060823d60601161098a575b8161093960609383613178565b810103126105be57604080519261094f8461315c565b6109588161346a565b84526109666020820161346a565b602085015201519061ffff821682036105be575060408201528263ffffffff6108c9565b3d915061092c565b604051903d90823e3d90fd5b7f4172d660000000000000000000000000000000000000000000000000000000008352600452602482fd5b506020813d602011610a12575b816109e360209383613178565b81010312610a0e57610a0973ffffffffffffffffffffffffffffffffffffffff9161346a565b61086a565b8280fd5b3d91506109d6565b6040513d85823e3d90fd5b50346105be5760406003193601126105be5760043567ffffffffffffffff81116105ba57610a57903690600401613380565b9060243567ffffffffffffffff81116105b65790610a7a84923690600401613380565b939091610a856140c9565b83905b828210610f4c5750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610f48578060051b83013585811215610f4457830161012081360312610f445760405194610aec86613124565b813567ffffffffffffffff811681036105ba578652602082013567ffffffffffffffff81116105ba5782019436601f870112156105ba57853595610b2f876132ed565b96610b3d6040519889613178565b80885260208089019160051b83010190368211610f445760208301905b828210610f11575050505060208701958652604083013567ffffffffffffffff8111610a0e57610b8d903690850161328a565b9160408801928352610bb7610ba53660608701613c8a565b9460608a0195865260c0369101613c8a565b956080890196875283515115610ee957610bdb67ffffffffffffffff8a5116614b5f565b15610eb25767ffffffffffffffff8951168252600860205260408220610c0286518261437e565b610c1088516002830161437e565b6004855191019080519067ffffffffffffffff8211610e8557610c338354613ac9565b601f8111610e4a575b50602090601f8311600114610dab57610c8a9291869183610da0575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b88518051821015610cc45790610cbe600192610cb78367ffffffffffffffff8f511692613a86565b5190614114565b01610c8f565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610d9267ffffffffffffffff6001979694985116925193519151610d5e610d29604051968796875261010060208801526101008701906131b9565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610abb565b015190508e80610c58565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b818110610e325750908460019594939210610dfb575b505050811b019055610c8d565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d8080610dee565b92936020600181928786015181550195019301610dd8565b610e759084875260208720601f850160051c81019160208610610e7b575b601f0160051c0190613d15565b8d610c3c565b9091508190610e68565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b813567ffffffffffffffff8111610f4057602091610f35839283369189010161328a565b815201910190610b5a565b8680fd5b8480fd5b8380f35b9267ffffffffffffffff610f6e610f698486889a9699979a613c5d565b6139fa565b1691610f798361490d565b156110df578284526008602052610f9560056040862001614696565b94845b8651811015610fce576001908587526008602052610fc760056040892001610fc0838b613a86565b5190614a19565b5001610f98565b509396929094509490948087526008602052600560408820888155886001820155886002820155886003820155886004820161100a8154613ac9565b8061109e575b5050500180549088815581611080575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020836001948a52600482528985604082208281550155808a52600582528985604082208281550155604051908152a101909194939294610a88565b885260208820908101905b818110156110205788815560010161108b565b601f81116001146110b45750555b888a80611010565b818352602083206110cf91601f01861c810190600101613d15565b80825281602081209155556110ac565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346105be5760206003193601126105be5760043567ffffffffffffffff81116105ba5761113d9036906004016133b1565b73ffffffffffffffffffffffffffffffffffffffff600a5416331415806112ed575b6112c157825b818110611170578380f35b61117b818385613c00565b67ffffffffffffffff61118d826139fa565b16906111a6826000526007602052604060002054151590565b1561129557907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e08361125561122f602060019897018b6111e782613c10565b1561125c57879052600460205261120e60408d206112083660408801613c8a565b9061437e565b868c52600560205261122a60408d206112083660a08801613c8a565b613c10565b9160405192151583526112486020840160408301613cd1565b60a0608084019101613cd1565ba201611165565b60026040828a61122a9452600860205261127e82822061120836858c01613c8a565b8a8152600860205220016112083660a08801613c8a565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561115f565b50346105be57806003193601126105be57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346105be5760406003193601126105be5760043567ffffffffffffffff81116105ba576113929036906004016133b1565b60243567ffffffffffffffff81116105b6576113b2903690600401613380565b9190926113bd6140c9565b845b82811061142957505050825b8181106113d6578380f35b8067ffffffffffffffff6113f0610f696001948688613c5d565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a2016113cb565b67ffffffffffffffff611440610f69838686613c00565b16611458816000526007602052604060002054151590565b156117b957611468828585613c00565b602081019060e081019061147b82613c10565b1561178d5760a0810161271061ffff61149383613c1d565b16101561177e5760c082019161271061ffff6114ae85613c1d565b1610156117465763ffffffff6114c386613c2c565b161561171a57858c52600b60205260408c206114de86613c2c565b63ffffffff169080549060408401916114f683613c2c565b60201b67ffffffff000000001693606086019461151286613c2c565b60401b6bffffffff000000000000000016966080019661153188613c2c565b60601b6fffffffff00000000000000000000000016916115508a613c1d565b60801b71ffff0000000000000000000000000000000016936115718c613c1d565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717815561162487613c10565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161790556040519661167590613c3d565b63ffffffff16875261168690613c3d565b63ffffffff16602087015261169a90613c3d565b63ffffffff1660408601526116ae90613c3d565b63ffffffff1660608501526116c290613c4e565b61ffff1660808401526116d490613c4e565b61ffff1660a08301526116e690613218565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a26001016113bf565b60248c877f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b60248c61ffff61175586613c1d565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff611755602493613c1d565b60248a857f12332265000000000000000000000000000000000000000000000000000000008252600452fd5b7f1e670e4b000000000000000000000000000000000000000000000000000000008752600452602486fd5b50346105be5760406003193601126105be5760043567ffffffffffffffff81116105ba57611816903690600401613380565b9061181f612fee565b9173ffffffffffffffffffffffffffffffffffffffff6001541633141580611a70575b611a445773ffffffffffffffffffffffffffffffffffffffff8316908115611a1c57845b818110611871578580f35b73ffffffffffffffffffffffffffffffffffffffff6118946103a2838588613c5d565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115611a115788916119dc575b50806118e9575b5050600101611866565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff8a16602484015260448084018590528352918a919061194a606482613178565b519082865af1156119d15787513d6119c85750813b155b61199c5790847f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602060019594604051908152a390386118df565b602488837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b60011415611961565b6040513d89823e3d90fd5b90506020813d8211611a09575b816119f660209383613178565b81010312611a055751386118d8565b8780fd5b3d91506119e9565b6040513d8a823e3d90fd5b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b6024847fcb1afbd700000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff600c5416331415611842565b50346105be57806003193601126105be57604051906006548083528260208101600684526020842092845b818110611ba4575050611ad292500383613178565b8151611af6611ae0826132ed565b91611aee6040519384613178565b8083526132ed565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015611b55578067ffffffffffffffff611b4260019388613a86565b5116611b4e8286613a86565b5201611b23565b50925090604051928392602084019060208552518091526040840192915b818110611b81575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611b73565b8454835260019485019487945060209093019201611abd565b50346105be5760206003193601126105be5760043573ffffffffffffffffffffffffffffffffffffffff81168091036105ba57611bf86140c9565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a1161760035580f35b50346105be5760206003193601126105be57611c9d611c8961043c613028565b6040519182916020835260208301906131b9565b0390f35b50346105be5760206003193601126105be577f307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a6020611cde612f16565b611ce66140c9565b6002547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff77ffffffff0000000000000000000000000000000000000000808460401c16169116176002557fffffffff0000000000000000000000000000000000000000000000000000000060405191168152a180f35b50346105be5760606003193601126105be57611d76612fa8565b90611d7f612fee565b6044359273ffffffffffffffffffffffffffffffffffffffff84168085036105b657611da96140c9565b73ffffffffffffffffffffffffffffffffffffffff82168015611ebf5794611eb9917f3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe70195967fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff85167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a557fffffffffffffffffffffffff0000000000000000000000000000000000000000600c541617600c556040519384938491604091949373ffffffffffffffffffffffffffffffffffffffff809281606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b50346105be5767ffffffffffffffff611eff366132a8565b929091611f0a6140c9565b1691611f23836000526007602052604060002054151590565b156110df578284526008602052611f5260056040862001611f45368486613225565b6020815191012090614a19565b15611f9757907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f9160405192839260208452602084019161348b565b0390a280f35b82611fdb836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161348b565b0390fd5b50346105be5760206003193601126105be5767ffffffffffffffff612002613028565b168152600860205261201960056040832001614696565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061205e612048836132ed565b926120566040519485613178565b8084526132ed565b01835b818110612135575050825b82518110156120b2578061208260019285613a86565b518552600960205261209660408620613b1c565b6120a08285613a86565b526120ab8184613a86565b500161206c565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106120ea57505050500390f35b91936020612125827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0600195979984950301865288516131b9565b96019201920185949391926120db565b806060602080938601015201612061565b50346105be5760206003193601126105be5760043567ffffffffffffffff81116105ba5760031960a091360301126105be57600490612183613a6d565b507f690a7a40000000000000000000000000000000000000000000000000000000008152fd5b50346105be5760206003193601126105be5760043567ffffffffffffffff81116105ba57604060031982360301126105ba57604051906121e882613108565b806004013567ffffffffffffffff81116105b65761220c9060043691840101613305565b825260248101359067ffffffffffffffff82116105b65760046122329236920101613305565b602082019081526122416140c9565b5191805b83518110156122b8578073ffffffffffffffffffffffffffffffffffffffff61227060019387613a86565b511661227b81614f41565b612287575b5001612245565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138612280565b509051815b81518110156123535773ffffffffffffffffffffffffffffffffffffffff6122e58284613a86565b5116801561232b57907feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef60208361231d600195614b20565b50604051908152a1016122bd565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8280f35b50346105be57806003193601126105be57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346105be5760206003193601126105be5760206123c767ffffffffffffffff6123b3613028565b166000526007602052604060002054151590565b6040519015158152f35b50346105be57806003193601126105be57805473ffffffffffffffffffffffffffffffffffffffff81163303612470577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346105be57806003193601126105be57600254600a54600c546040805173ffffffffffffffffffffffffffffffffffffffff94851681529284166020840152921691810191909152606090f35b50346105be576124f5366132a8565b612501939291936140c9565b67ffffffffffffffff8216612523816000526007602052604060002054151590565b15612542575061253f9293612539913691613225565b90614114565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346105be57806003193601126105be57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346105be5760406003193601126105be576125d8613028565b906024359067ffffffffffffffff82116105be5760206123c7846125ff366004870161328a565b90613a30565b50346105be5760206003193601126105be5760043567ffffffffffffffff81116105ba5760031961010091360301126105be5780600491604051612648816130bd565b527f690a7a40000000000000000000000000000000000000000000000000000000008152fd5b50346105be57806003193601126105be57602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b50346105be5760406003193601126105be5760043567ffffffffffffffff81116105ba5780600401906101006003198236030112610a0e576126e2612f4a565b91836040516126f0816130bd565b52606482013592608483019161270583613a0f565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036129fa57602484019577ffffffffffffffff0000000000000000000000000000000061276b886139fa565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156129ef5782916129d0575b506129a857506127fa610336876139fa565b612803866139fa565b9061282060a48601926125ff6128198585614078565b3691613225565b1561296157505067ffffffffffffffff6128dc60446128d56020987ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09689877fffffffff0000000000000000000000000000000000000000000000000000000060809916151560001461294557610f69926103a26128c6927fffffffff0000000000000000000000000000000000000000000000000000000060025460401b1690613ef4565b6128cf846139fa565b906147f2565b9501613a0f565b9373ffffffffffffffffffffffffffffffffffffffff60405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a28060405161293c816130bd565b52604051908152f35b50612952610f6992613a0f565b61295b846139fa565b90614779565b61296b9250614078565b611fdb6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161348b565b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b6129e9915060203d60201161055a5761054c8183613178565b386127e8565b6040513d84823e3d90fd5b60248673ffffffffffffffffffffffffffffffffffffffff61058d86613a0f565b50346105be57806003193601126105be57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346105be57806003193601126105be57604051600d8054808352908352909160208301917fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb5915b818110612ac457611c9d85612ab881870382613178565b6040519182918261306d565b8254845260209093019260019283019201612aa1565b50346105be5760406003193601126105be57612af4613028565b6024359182151583036105be57610140612bb7612b118585613977565b612b6760409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346105be5760206003193601126105be57602090612bd6612fa8565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346105be57806003193601126105be57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346105be5760c06003193601126105be57612c88612fa8565b50612c91613011565b612c99612fcb565b50608435917fffffffff00000000000000000000000000000000000000000000000000000000831683036105be5760a4359067ffffffffffffffff82116105be5760a063ffffffff8061ffff612cfe8888612cf73660048b0161303f565b50506137c7565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346105be57806003193601126105be5750611c9d604051612d48604082613178565b601d81527f434354505468726f756768434356546f6b656e506f6f6c20322e302e3000000060208201526040519182916020835260208301906131b9565b50346105be5760c06003193601126105be57612da0612fa8565b612da8613011565b6064357fffffffff00000000000000000000000000000000000000000000000000000000811681036105b65760843567ffffffffffffffff8111610f4457612df490369060040161303f565b93909260a4359560028710156105be57611c9d612ab888888888604435888a6134ca565b9050346105ba5760206003193601126105ba576020907fffffffff00000000000000000000000000000000000000000000000000000000612e57612f16565b167faff2afbf000000000000000000000000000000000000000000000000000000008114908115612eec575b8115612ec2575b8115612e98575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483612e91565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150612e8a565b7f940a15420000000000000000000000000000000000000000000000000000000081149150612e83565b600435907fffffffff0000000000000000000000000000000000000000000000000000000082168203612f4557565b600080fd5b602435907fffffffff0000000000000000000000000000000000000000000000000000000082168203612f4557565b604435907fffffffff0000000000000000000000000000000000000000000000000000000082168203612f4557565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203612f4557565b6064359073ffffffffffffffffffffffffffffffffffffffff82168203612f4557565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203612f4557565b6024359067ffffffffffffffff82168203612f4557565b6004359067ffffffffffffffff82168203612f4557565b9181601f84011215612f455782359167ffffffffffffffff8311612f455760208381860195010111612f4557565b602060408183019282815284518094520192019060005b8181106130915750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613084565b6020810190811067ffffffffffffffff8211176130d957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176130d957604052565b60a0810190811067ffffffffffffffff8211176130d957604052565b60e0810190811067ffffffffffffffff8211176130d957604052565b6060810190811067ffffffffffffffff8211176130d957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176130d957604052565b919082519283825260005b8481106132035750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016131c4565b35908115158203612f4557565b92919267ffffffffffffffff82116130d9576040519161326d601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184613178565b829481845281830111612f45578281602093846000960137010152565b9080601f83011215612f45578160206132a593359101613225565b90565b906040600319830112612f455760043567ffffffffffffffff81168103612f4557916024359067ffffffffffffffff8211612f45576132e99160040161303f565b9091565b67ffffffffffffffff81116130d95760051b60200190565b9080601f83011215612f455781359061331d826132ed565b9261332b6040519485613178565b82845260208085019360051b820101918211612f4557602001915b8183106133535750505090565b823573ffffffffffffffffffffffffffffffffffffffff81168103612f4557815260209283019201613346565b9181601f84011215612f455782359167ffffffffffffffff8311612f45576020808501948460051b010111612f4557565b9181601f84011215612f455782359167ffffffffffffffff8311612f45576020808501948460081b010111612f4557565b818102929181159184041417156133f557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b811561342e570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b919082039182116133f557565b519073ffffffffffffffffffffffffffffffffffffffff82168203612f4557565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b92959390919473ffffffffffffffffffffffffffffffffffffffff600354169586156137a557809760028710156137765773ffffffffffffffffffffffffffffffffffffffff9861362b957fffffffff0000000000000000000000000000000000000000000000000000000093896137475767ffffffffffffffff8216600052600b6020526040600020906040519161356283613140565b549163ffffffff8316815263ffffffff8360201c16602082015263ffffffff8360401c16604082015263ffffffff8360601c16606082015260c061ffff8460801c169182608082015260ff60a082019561ffff8160901c16875260a01c16151591829101526136f3575b50505067ffffffffffffffff905b6040519b8c997f06b859ef000000000000000000000000000000000000000000000000000000008b521660048a0152166024880152604487015216606485015260c0608485015260c484019161348b565b928180600095869560a483015203915afa9182156136e657819261364e57505090565b9091503d8083833e6136608183613178565b810190602081830312610a0e5780519067ffffffffffffffff82116105b6570181601f82011215610a0e57805190613697826132ed565b936136a56040519586613178565b82855260208086019360051b8301019384116105be5750602001905b8282106136ce5750505090565b602080916136db8461346a565b8152019101906136c1565b50604051903d90823e3d90fd5b92935067ffffffffffffffff928587161561372f575061271061371e61ffff613725945116836133e2565b049061345d565b915b9038806135cc565b613741925061371e61271091836133e2565b91613727565b67ffffffffffffffff9192506137709061376a61376536898b613225565b613d2c565b90613de7565b916135da565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b50505050505050506040516137bb602082613178565b60008152600036813790565b67ffffffffffffffff909291926138057fffffffff0000000000000000000000000000000000000000000000000000000060025460401b1685613ef4565b16600052600b60205260406000206040519061382082613140565b549163ffffffff83169384835263ffffffff8460201c169384602085015263ffffffff8160401c169182604086015263ffffffff8260601c169081606087015261ffff8360801c169586608082015260ff61ffff8560901c16948560a084015260a01c16159060c082159101526138cd577fffffffff00000000000000000000000000000000000000000000000000000000166138c257505093929190600190565b959493509160019150565b5050505092505050600090600090600090600090600090565b604051906138f382613124565b60006080838281528260208201528260408201528260608201520152565b9060405161391e81613124565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916139896138e6565b506139926138e6565b506139c6571660005260086020526040600020906132a56139ba60026139bf6139ba86613911565b613fdb565b9401613911565b16908160005260046020526139e16139ba6040600020613911565b9160005260056020526132a56139ba6040600020613911565b3567ffffffffffffffff81168103612f455790565b3573ffffffffffffffffffffffffffffffffffffffff81168103612f455790565b9067ffffffffffffffff6132a592166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b60405190613a7a82613108565b60606020838281520152565b8051821015613a9a5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015613b12575b6020831014613ae357565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613ad8565b9060405191826000825492613b3084613ac9565b8084529360018116908115613b9e5750600114613b57575b50613b5592500383613178565b565b90506000929192526020600020906000915b818310613b82575050906020613b559282010138613b48565b6020919350806001915483858901015201910190918492613b69565b60209350613b559592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613b48565b67ffffffffffffffff1660005260086020526132a56004604060002001613b1c565b9190811015613a9a5760081b0190565b358015158103612f455790565b3561ffff81168103612f455790565b3563ffffffff81168103612f455790565b359063ffffffff82168203612f4557565b359061ffff82168203612f4557565b9190811015613a9a5760051b0190565b35906fffffffffffffffffffffffffffffffff82168203612f4557565b9190826060910312612f4557604051613ca28161315c565b6040613ccc818395613cb381613218565b8552613cc160208201613c6d565b602086015201613c6d565b910152565b6fffffffffffffffffffffffffffffffff613d0f60408093613cf281613218565b1515865283613d0360208301613c6d565b16602087015201613c6d565b16910152565b818110613d20575050565b60008155600101613d15565b80518015613d9c57602003613d5e578051602082810191830183900312612f4557519060ff8211613d5e575060ff1690565b611fdb906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484015260248301906131b9565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116133f557565b60ff16604d81116133f557600a0a90565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414613eed57828411613ec35790613e2c91613dc2565b91604d60ff8416118015613e8a575b613e5457505090613e4e6132a592613dd6565b906133e2565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50613e9483613dd6565b801561342e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411613e3b565b613ecc91613dc2565b91604d60ff841611613e5457505090613ee76132a592613dd6565b90613424565b5050505090565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115613fd657613f27816145bb565b7dffff00000000000000000000000000000000000000000000000000000000601082811c9085901c1616613fd65761ffff8360e01c168015918215613fc5575b5050613f71575050565b7fffffffff0000000000000000000000000000000000000000000000000000000092507fdf63778f000000000000000000000000000000000000000000000000000000006000526004521660245260446000fd5b60e01c61ffff161090503880613f67565b505050565b613fe36138e6565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614040602085019361403a61402d63ffffffff8751164261345d565b85608089015116906133e2565b906145ae565b8082101561405957505b16825263ffffffff4216905290565b905061404a565b90816020910312612f4557518015158103612f455790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215612f45570180359067ffffffffffffffff8211612f4557602001918136038313612f4557565b73ffffffffffffffffffffffffffffffffffffffff6001541633036140ea57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156143545767ffffffffffffffff81516020830120921691826000526008602052614149816005604060002001614b98565b156143105760005260096020526040600020815167ffffffffffffffff81116130d9576141768254613ac9565b601f81116142de575b506020601f821160011461421857916141f2827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936142089560009161420d575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90556040519182916020835260208301906131b9565b0390a2565b9050840151386141c1565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106142c65750926142089492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061428f575b5050811b019055611c89565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880614283565b9192602060018192868a015181550194019201614248565b61430a90836000526020600020601f840160051c81019160208510610e7b57601f0160051c0190613d15565b3861417f565b5090611fdb6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401526040602484015260448301906131b9565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b815191929115614500576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061449d57613b5591925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836144fe604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040840151161580159061458f575b61452e57613b5591926143c1565b6064836144fe604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020840151161515614520565b919082018092116133f557565b7fffffffff000000000000000000000000000000000000000000000000000000008116908115614692577dffff000000000000000000000000000000000000000000000000000000008116156146895760ff60015b169060f01c80614653575b506001036146265750565b7fc512f96c0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60005b60108110614664575061461b565b6001811b8216614677575b600101614656565b91600181018091116133f5579161466f565b60ff6000614610565b5050565b906040519182815491828252602082019060005260206000209260005b8181106146c8575050613b5592500383613178565b84548352600194850194879450602090930192016146b3565b67ffffffffffffffff16614702816000526007602052604060002054151590565b1561474c575033600052600e6020526040600020541561471e57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9167ffffffffffffffff7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9216928360005260086020526147c281836002604060002001614c85565b6040805173ffffffffffffffffffffffffffffffffffffffff909216825260208201929092529081908101614208565b91909167ffffffffffffffff83169283600052600560205260ff60406000205460a01c16156148575750907fc6735cd4fa2bbe7b203b1682936e6ee61bc1702464bbbd12abb6630229d9a5f9918360005260056020526147c281836040600020614c85565b90613b559350614779565b8054821015613a9a5760005260206000200190600090565b805480156148de577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906148af8282614862565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000818152600760205260409020548015614a12577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116133f557600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116133f5578181036149a3575b50505061498f600661487a565b600052600760205260006040812055600190565b6149fa6149b46149c5936006614862565b90549060031b1c9283926006614862565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526007602052604060002055388080614982565b5050600090565b906001820191816000528260205260406000205490811515600014614aec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918083116133f55781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019081116133f5578381614aa39503614ab5575b50505061487a565b60005260205260006040812055600190565b614ad5614ac56149c59386614862565b90549060031b1c92839286614862565b905560005284602052604060002055388080614a9b565b50505050600090565b805490680100000000000000008210156130d957816149c5916001614b1c94018155614862565b9055565b80600052600e60205260406000205415600014614b5957614b4281600d614af5565b600d5490600052600e602052604060002055600190565b50600090565b80600052600760205260406000205415600014614b5957614b81816006614af5565b600654906000526007602052604060002055600190565b6000828152600182016020526040902054614a125780614bba83600193614af5565b80549260005201602052604060002055600190565b9167ffffffffffffffff7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449216928360005260086020526147c281836040600020614c85565b91909167ffffffffffffffff83169283600052600460205260ff60406000205460a01c1615614c7a5750907f28d6c52e2b0b7587b0d195539fbe6af984b28791aca4d2cc0844244e38bce29e918360005260046020526147c281836040600020614c85565b90613b559350614bcf565b9182549060ff8260a01c16158015614f39575b614f33576fffffffffffffffffffffffffffffffff82169160018501908154614cdd63ffffffff6fffffffffffffffffffffffffffffffff83169360801c164261345d565b9081614e95575b5050848110614e495750838310614d3e575050614d136fffffffffffffffffffffffffffffffff92839261345d565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315614ddd5781614d569161345d565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908082116133f557614da4614da99273ffffffffffffffffffffffffffffffffffffffff966145ae565b613424565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611614f0957614eb09261403a9160801c906133e2565b80841015614f045750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880614ce4565b614ebb565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215614c98565b6000818152600e60205260409020548015614a12577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018181116133f557600d54907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116133f557808203614fd7575b505050614fc3600d61487a565b600052600e60205260006040812055600190565b614ff9614fe86149c593600d614862565b90549060031b1c928392600d614862565b9055600052600e602052604060002055388080614fb656fea164736f6c634300081a000a",
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getAllowedFinalityConfig")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetAllowedFinalityConfig(&_CCTPThroughCCVTokenPool.CallOpts)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetAllowedFinalityConfig() ([4]byte, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetAllowedFinalityConfig(&_CCTPThroughCCVTokenPool.CallOpts)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, fastFinality)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector, fastFinality)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPThroughCCVTokenPool.CallOpts, remoteChainSelector, fastFinality)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)

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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetFee(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

	error) {
	return _CCTPThroughCCVTokenPool.Contract.GetFee(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, requestedFinalityConfig, arg5)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, sourceDenominatedAmount, requestedFinalityConfig, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRequiredCCVs(&_CCTPThroughCCVTokenPool.CallOpts, localToken, remoteChainSelector, sourceDenominatedAmount, requestedFinalityConfig, extraData, direction)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetRequiredCCVs(&_CCTPThroughCCVTokenPool.CallOpts, localToken, remoteChainSelector, sourceDenominatedAmount, requestedFinalityConfig, extraData, direction)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _CCTPThroughCCVTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CCTPThroughCCVTokenPool.Contract.GetTokenTransferFeeConfig(&_CCTPThroughCCVTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.LockOrBurn0(&_CCTPThroughCCVTokenPool.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.LockOrBurn0(&_CCTPThroughCCVTokenPool.TransactOpts, lockOrBurnIn, requestedFinalityConfig, tokenArgs)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn, requestedFinalityConfig)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint(&_CCTPThroughCCVTokenPool.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint(&_CCTPThroughCCVTokenPool.TransactOpts, releaseOrMintIn, requestedFinalityConfig)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "releaseOrMint0", arg0)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) ReleaseOrMint0(arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint0(&_CCTPThroughCCVTokenPool.TransactOpts, arg0)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) ReleaseOrMint0(arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.ReleaseOrMint0(&_CCTPThroughCCVTokenPool.TransactOpts, arg0)
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactor) SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.contract.Transact(opts, "setAllowedFinalityConfig", allowedFinality)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetAllowedFinalityConfig(&_CCTPThroughCCVTokenPool.TransactOpts, allowedFinality)
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolTransactorSession) SetAllowedFinalityConfig(allowedFinality [4]byte) (*types.Transaction, error) {
	return _CCTPThroughCCVTokenPool.Contract.SetAllowedFinalityConfig(&_CCTPThroughCCVTokenPool.TransactOpts, allowedFinality)
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

type CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumedIterator struct {
	Event *CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed)
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
		it.Event = new(CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed)
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

func (it *CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterFastFinalityInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "FastFinalityInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "FastFinalityInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchFastFinalityInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "FastFinalityInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "FastFinalityInboundRateLimitConsumed", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseFastFinalityInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed, error) {
	event := new(CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "FastFinalityInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumedIterator struct {
	Event *CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed)
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
		it.Event = new(CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed)
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

func (it *CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterFastFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "FastFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumedIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "FastFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchFastFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "FastFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "FastFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseFastFinalityOutboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed, error) {
	event := new(CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "FastFinalityOutboundRateLimitConsumed", log); err != nil {
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

type CCTPThroughCCVTokenPoolFinalityConfigSetIterator struct {
	Event *CCTPThroughCCVTokenPoolFinalityConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPThroughCCVTokenPoolFinalityConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPThroughCCVTokenPoolFinalityConfigSet)
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
		it.Event = new(CCTPThroughCCVTokenPoolFinalityConfigSet)
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

func (it *CCTPThroughCCVTokenPoolFinalityConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPThroughCCVTokenPoolFinalityConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPThroughCCVTokenPoolFinalityConfigSet struct {
	AllowedFinality [4]byte
	Raw             types.Log
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) FilterFinalityConfigSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolFinalityConfigSetIterator, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.FilterLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPThroughCCVTokenPoolFinalityConfigSetIterator{contract: _CCTPThroughCCVTokenPool.contract, event: "FinalityConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFinalityConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPThroughCCVTokenPool.contract.WatchLogs(opts, "FinalityConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPThroughCCVTokenPoolFinalityConfigSet)
				if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
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

func (_CCTPThroughCCVTokenPool *CCTPThroughCCVTokenPoolFilterer) ParseFinalityConfigSet(log types.Log) (*CCTPThroughCCVTokenPoolFinalityConfigSet, error) {
	event := new(CCTPThroughCCVTokenPoolFinalityConfigSet)
	if err := _CCTPThroughCCVTokenPool.contract.UnpackLog(event, "FinalityConfigSet", log); err != nil {
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
	FastFinality              bool
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

func (CCTPThroughCCVTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x3f1036e85d016a93254a0b1415844f79b85424959d90ae5ad51ce8f4533fe701")
}

func (CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xc6735cd4fa2bbe7b203b1682936e6ee61bc1702464bbbd12abb6630229d9a5f9")
}

func (CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x28d6c52e2b0b7587b0d195539fbe6af984b28791aca4d2cc0844244e38bce29e")
}

func (CCTPThroughCCVTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CCTPThroughCCVTokenPoolFinalityConfigSet) Topic() common.Hash {
	return common.HexToHash("0x307cf716eade81675bea3ccb6917b0f91baa2160056765d9a83d76f819caf06a")
}

func (CCTPThroughCCVTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (CCTPThroughCCVTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
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

	GetAllowedFinalityConfig(opts *bind.CallOpts) ([4]byte, error)

	GetCCTPVerifier(opts *bind.CallOpts) (common.Address, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, fastFinality bool) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, requestedFinalityConfig [4]byte, arg5 []byte) (GetFee,

		error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, sourceDenominatedAmount *big.Int, requestedFinalityConfig [4]byte, extraData []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 [4]byte, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error)

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

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, requestedFinalityConfig [4]byte, tokenArgs []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, requestedFinalityConfig [4]byte) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, arg0 PoolReleaseOrMintInV1) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetAllowedFinalityConfig(opts *bind.TransactOpts, allowedFinality [4]byte) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address, feeAdmin common.Address) (*types.Transaction, error)

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

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CCTPThroughCCVTokenPoolDynamicConfigSet, error)

	FilterFastFinalityInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumedIterator, error)

	WatchFastFinalityInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastFinalityInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolFastFinalityInboundRateLimitConsumed, error)

	FilterFastFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumedIterator, error)

	WatchFastFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseFastFinalityOutboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolFastFinalityOutboundRateLimitConsumed, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPThroughCCVTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCTPThroughCCVTokenPoolFeeTokenWithdrawn, error)

	FilterFinalityConfigSet(opts *bind.FilterOpts) (*CCTPThroughCCVTokenPoolFinalityConfigSetIterator, error)

	WatchFinalityConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolFinalityConfigSet) (event.Subscription, error)

	ParseFinalityConfigSet(log types.Log) (*CCTPThroughCCVTokenPoolFinalityConfigSet, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*CCTPThroughCCVTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPThroughCCVTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CCTPThroughCCVTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*CCTPThroughCCVTokenPoolLockedOrBurned, error)

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
