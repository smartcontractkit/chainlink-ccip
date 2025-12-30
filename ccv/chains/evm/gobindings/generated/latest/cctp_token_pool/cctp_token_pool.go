// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cctp_token_pool

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

var CCTPTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdvancedPoolHooks\",\"inputs\":[],\"outputs\":[{\"name\":\"advancedPoolHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCTPVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBlockConfirmation\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateAdvancedPoolHooks\",\"inputs\":[{\"name\":\"newHook\",\"type\":\"address\",\"internalType\":\"contract IAdvancedPoolHooks\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdvancedPoolHooksUpdated\",\"inputs\":[{\"name\":\"oldHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"},{\"name\":\"newHook\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contract IAdvancedPoolHooks\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBlockConfirmationSet\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustomBlockConfirmationsNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"IPoolV1NotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x6101008060405234610390576153b5803803809161001d82856103f9565b833981019060c0818303126103905780516001600160a01b038116908181036103905761004c6020840161041c565b906100596040850161042a565b906100666060860161042a565b936100736080870161042a565b60a087015190966001600160401b03821161039057019680601f89011215610390578751976001600160401b03891161031b578860051b9060208201996100bd6040519b8c6103f9565b8a526020808b019282010192831161039057602001905b8282106103e15750505033156103d057600180546001600160a01b03191633179055801580156103bf575b80156103ae575b61039d5760049260209260805260c0526040519283809263313ce56760e01b82525afa6000918161035c575b50610331575b5060a052600380546001600160a01b0319908116909155600280549091166001600160a01b039290921691909117905560405160209061017882826103f9565b60008152600036813760408051949085016001600160401b0381118682101761031b576040528452808285015260005b815181101561020f576001906001600160a01b036101c6828561043e565b5116846101d282610480565b6101df575b5050016101a8565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138846101d7565b5050915160005b8151811015610287576001600160a01b03610231828461043e565b5116908115610276577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef858361026860019561057e565b50604051908152a101610216565b6342bcdf7f60e11b60005260046000fd5b6001600160a01b03831660e052604051614dd690816105df823960805181818161193d01528181611aad01528181611b7901528181611e22015281816127c20152818161293001528181612a0301528181612cbe015281816130cc0152613126015260a05181612f1d015260c051818181610cda015281816119d8015261285d015260e0518181816112b5015261263e0152f35b634e487b7160e01b600052604160045260246000fd5b60ff1660ff82168181036103455750610138565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610395575b81610378602093836103f9565b81010312610390576103899061041c565b9038610132565b600080fd5b3d915061036b565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610106565b506001600160a01b038516156100ff565b639b15e16f60e01b60005260046000fd5b602080916103ee8461042a565b8152019101906100d4565b601f909101601f19168101906001600160401b0382119082101761031b57604052565b519060ff8216820361039057565b51906001600160a01b038216820361039057565b80518210156104525760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104525760005260206000200190600090565b6000818152600d6020526040902054801561057757600019810181811161056157600c5460001981019190821161056157808203610510575b505050600c5480156104fa57600019016104d481600c610468565b8154906000199060031b1b19169055600c55600052600d60205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61054961052161053293600c610468565b90549060031b1c928392600c610468565b819391549060031b91821b91600019901b19161790565b9055600052600d6020526040600020553880806104b9565b634e487b7160e01b600052601160045260246000fd5b5050600090565b80600052600d602052604060002054156000146105d857600c546801000000000000000081101561031b576105bf610532826001859401600c55600c610468565b9055600c5490600052600d602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146131ab57508063181f5a771461314a57806321df0da7146130f9578063240028e8146130955780632422ac4514612fb65780632451a62714612f4157806324f65ee714612f035780632c06340414612e6a57806337a3210d14612e365780633907753714612dcd578063489a68f2146127485780634c5ef0ed146127015780634e921c3014612662578063615521a71461261157806362ddd3c41461258a5780637437ff9f1461254957806379ba5097146124825780638926f54f1461243c57806389720a62146123bc5780638da5cb5b1461238857806391a2749a146121da5780639a4575b914612177578063a42a7b8b14612010578063acfecf9114611f18578063b1c71c6514611895578063b794658014611858578063bfeffd3f146117ac578063c4bffe2b14611681578063c7230a601461146d578063d8aa3f401461114f578063dc04fa1f14610cfe578063dc0bd97114610cad578063dcbd41bc14610aa9578063e8a1da17146103e5578063f2fde38b14610316578063fa41d79c146102f15763ff8e03f3146101b857600080fd5b346102ee5760406003193601126102ee576101d1613408565b906101da613453565b6101e261404f565b73ffffffffffffffffffffffffffffffffffffffff83169283156102c6577f22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e616644797092937fffffffffffffffffffffffff0000000000000000000000000000000000000000600254161760025573ffffffffffffffffffffffffffffffffffffffff82167fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a556102c06040519283928390929173ffffffffffffffffffffffffffffffffffffffff60209181604085019616845216910152565b0390a180f35b6004837f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b80fd5b50346102ee57806003193601126102ee57602061ffff60035460a01c16604051908152f35b50346102ee5760206003193601126102ee5773ffffffffffffffffffffffffffffffffffffffff610345613408565b61034d61404f565b163381146103bd57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346102ee5760406003193601126102ee5760043567ffffffffffffffff8111610902576104179036906004016136db565b9060243567ffffffffffffffff8111610aa5579061043a849236906004016136db565b93909161044561404f565b83905b82821061090a5750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610906578060051b830135858112156108fe578301610120813603126108fe57604051946104ac8661334c565b6104b5826134c5565b8652602082013567ffffffffffffffff81116109025782019436601f87011215610902578535956104e58761365e565b966104f36040519889613368565b80885260208089019160051b830101903682116108fe5760208301905b8282106108cb575050505060208701958652604083013567ffffffffffffffff81116108c75761054390369085016135fb565b916040880192835261056d61055b3660608701613e70565b9460608a0195865260c0369101613e70565b95608089019687528351511561089f5761059167ffffffffffffffff8a51166149d4565b156108685767ffffffffffffffff89511682526008602052604082206105b8865182614367565b6105c6885160028301614367565b6004855191019080519067ffffffffffffffff821161083b576105e98354613cbe565b601f8111610800575b50602090601f8311600114610761576106409291869183610756575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b8851805182101561067a579061067460019261066d8367ffffffffffffffff8f511692613c7b565b519061409a565b01610645565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561074867ffffffffffffffff60019796949851169251935191516107146106df604051968796875261010060208801526101008701906133a9565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101939290919361047b565b015190508e8061060e565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106107e857509084600195949392106107b1575b505050811b019055610643565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806107a4565b9293602060018192878601518155019501930161078e565b61082b9084875260208720601f850160051c81019160208610610831575b601f0160051c0190613efb565b8d6105f2565b909150819061081e565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116108fa576020916108ef83928336918901016135fb565b815201910190610510565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff61092c6109278486889a9699979a613e43565b6139d6565b169161093783614782565b15610a79578284526008602052610953600560408620016145a4565b94845b865181101561098c5760019085875260086020526109856005604089200161097e838b613c7b565b519061488e565b5001610956565b50939692909450949094808752600860205260056040882088815588600182015588600282015588600382015588600482016109c88154613cbe565b80610a38575b5050500180549088815581610a1a575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101909194939294610448565b885260208820908101905b818110156109de57888155600101610a25565b601f8111600114610a4e5750555b888a806109ce565b81835260208320610a6991601f01861c810190600101613efb565b8082528160208120915555610a46565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b8380fd5b50346102ee5760206003193601126102ee5760043567ffffffffffffffff811161090257610adb90369060040161370c565b73ffffffffffffffffffffffffffffffffffffffff600a541633141580610c8b575b610c5f57825b818110610b0e578380f35b610b19818385613df5565b67ffffffffffffffff610b2b826139d6565b1690610b44826000526007602052604060002054151590565b15610c3357907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610bf3610bcd602060019897018b610b8582613e05565b15610bfa578790526004602052610bac60408d20610ba63660408801613e70565b90614367565b868c526005602052610bc860408d20610ba63660a08801613e70565b613e05565b916040519215158352610be66020840160408301613eb7565b60a0608084019101613eb7565ba201610b03565b60026040828a610bc894526008602052610c1c828220610ba636858c01613e70565b8a815260086020522001610ba63660a08801613e70565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610afd565b50346102ee57806003193601126102ee57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102ee5760406003193601126102ee5760043567ffffffffffffffff811161090257610d3090369060040161370c565b60243567ffffffffffffffff8111610aa557610d509036906004016136db565b919092610d5b61404f565b845b828110610dc757505050825b818110610d74578380f35b8067ffffffffffffffff610d8e6109276001948688613e43565b16808652600b6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610d69565b610dd5610927828585613df5565b610de0828585613df5565b90602082019060e0830190610df482613e05565b1561111a5760a0840161271061ffff610e0c83613e12565b16101561110b5760c085019161271061ffff610e2785613e12565b1610156110d35763ffffffff610e3c86613e21565b161561109e5767ffffffffffffffff1694858c52600b60205260408c20610e6286613e21565b63ffffffff16908054906040840191610e7a83613e21565b60201b67ffffffff0000000016936060860194610e9686613e21565b60401b6bffffffff0000000000000000169660800196610eb588613e21565b60601b6fffffffff0000000000000000000000001691610ed48a613e12565b60801b71ffff000000000000000000000000000000001693610ef58c613e12565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610fa887613e05565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610ff990613e32565b63ffffffff16875261100a90613e32565b63ffffffff16602087015261101e90613e32565b63ffffffff16604086015261103290613e32565b63ffffffff16606085015261104690613559565b61ffff16608084015261105890613559565b61ffff1660a083015261106a906134da565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610d5d565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff6110e286613e12565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff6110e2602493613e12565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346102ee5760806003193601126102ee57611169613408565b506111726134ae565b61117a613548565b506064359067ffffffffffffffff82116108c7576111a567ffffffffffffffff923690600401613568565b50508260c06040516111b681613314565b8281528260208201528260408201528260608201528260808201528260a0820152015216808252600b6020526040822090604051916111f483613314565b549063ffffffff82168352602083019363ffffffff8360201c168552604084019463ffffffff8460401c168652606085019063ffffffff8560601c168252608086019261ffff8660801c16845260a087019461ffff8760901c16865260ff60c089019760a01c1615158752604051907f958021a700000000000000000000000000000000000000000000000000000000825260048201526040602482015281604482015260208160648173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611462578291611410575b50606073ffffffffffffffffffffffffffffffffffffffff916004604051809481937f7437ff9f000000000000000000000000000000000000000000000000000000008352165afa91821561140457809261138f575b505061ffff9493919263ffffffff60e09981889687604083970151168952816040519c51168c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b9091506060823d6060116113fc575b816113ab60609383613368565b810103126102ee5760408051926113c184613330565b6113ca81613a49565b84526113d860208201613a49565b602085015201519061ffff821682036102ee575060408201528263ffffffff61133b565b3d915061139e565b604051903d90823e3d90fd5b90506020813d60201161145a575b8161142b60209383613368565b8101031261090257606061145373ffffffffffffffffffffffffffffffffffffffff92613a49565b91506112e5565b3d915061141e565b6040513d84823e3d90fd5b50346102ee5760406003193601126102ee5760043567ffffffffffffffff81116109025761149f9036906004016136db565b906114a8613453565b916114b161404f565b835b8181106114be578480f35b73ffffffffffffffffffffffffffffffffffffffff6114e66114e1838587613e43565b6139eb565b166040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152602081602481855afa908115611676578791611645575b508061153b575b50506001016114b3565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602080830191825273ffffffffffffffffffffffffffffffffffffffff89166024840152604480840185905283529189919061159c606482613368565b519082865af11561163a5786513d6116315750813b155b6116055790600192916040519081527f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e602073ffffffffffffffffffffffffffffffffffffffff891692a39038611531565b602487837f5274afe7000000000000000000000000000000000000000000000000000000008252600452fd5b600114156115b3565b6040513d88823e3d90fd5b90506020813d821161166e575b8161165f60209383613368565b810103126108fa57513861152a565b3d9150611652565b6040513d89823e3d90fd5b50346102ee57806003193601126102ee57604051906006548083528260208101600684526020842092845b8181106117935750506116c192500383613368565b81516116e56116cf8261365e565b916116dd6040519384613368565b80835261365e565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015611744578067ffffffffffffffff61173160019388613c7b565b511661173d8286613c7b565b5201611712565b50925090604051928392602084019060208552518091526040840192915b818110611770575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611762565b84548352600194850194879450602090930192016116ac565b50346102ee5760206003193601126102ee5760043573ffffffffffffffffffffffffffffffffffffffff8116809103610902576117e761404f565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006003547fbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d5812096040805173ffffffffffffffffffffffffffffffffffffffff84168152856020820152a1161760035580f35b50346102ee5760206003193601126102ee5761189161187d611878613497565b613dd3565b6040519182916020835260208301906133a9565b0390f35b50346102ee5760606003193601126102ee576004359067ffffffffffffffff82116102ee57816004019060a060031984360301126102ee576118d5613537565b6044359367ffffffffffffffff85116108c7576118fa61191794953690600401613568565b9490611904613c62565b5061190f8488614304565b953691613596565b6084820191611925836139eb565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611ece57602481019677ffffffffffffffff0000000000000000000000000000000061198b896139d6565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611676578791611e9f575b50611e7757611a1e611a19896139d6565b6145ef565b61ffff611a3060648401359889613f54565b9516948515611dc65761ffff60035460a01c168015611d9e57808710611d6e57507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611a848b6139d6565b1691828952600460205280611ad560408b2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391614a44565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff600354169283611c57575b611c1c88611baf6118788c611b32816139d6565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10606067ffffffffffffffff6040519373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001685523360208601528860408601521692a26139d6565b90611c4d6040517f3047587c00000000000000000000000000000000000000000000000000000000602082015260048152611beb602482613368565b60405193611bf8856132f8565b845260208401908152604051948594604086525160408087015260808601906133a9565b90517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08583030160608601526133a9565b9060208301520390f35b833b156108fa57869493929185918a604051988997889687957f5c3af7ca000000000000000000000000000000000000000000000000000000008752600487016060905280611ca591614687565b6064880160a09052610104880190611cbc92613a6a565b93611cc6906134c5565b67ffffffffffffffff166084870152604401611ce190613476565b73ffffffffffffffffffffffffffffffffffffffff1660a48601528c60c4860152611d0b90613476565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526024840152828103600319016044840152611d40916133a9565b03925af1801561146257611d59575b8080808080611b1e565b611d64828092613368565b6102ee5780611d4f565b87604491887f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b6004887f98d50fd7000000000000000000000000000000000000000000000000000000008152fd5b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611df98b6139d6565b1691828952600860205280611e4a60408b2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391614a44565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611afe565b6004867f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611ec1915060203d602011611ec7575b611eb98183613368565b810190613fe6565b38611a08565b503d611eaf565b60248573ffffffffffffffffffffffffffffffffffffffff611eef866139eb565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102ee5767ffffffffffffffff611f3036613619565b929091611f3b61404f565b1691611f54836000526007602052604060002054151590565b15610a79578284526008602052611f8360056040862001611f76368486613596565b602081519101209061488e565b15611fc857907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611fc2604051928392602084526020840191613a6a565b0390a280f35b8261200c836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191613a6a565b0390fd5b50346102ee5760206003193601126102ee5767ffffffffffffffff612033613497565b168152600860205261204a600560408320016145a4565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061208f6120798361365e565b926120876040519485613368565b80845261365e565b01835b818110612166575050825b82518110156120e357806120b360019285613c7b565b51855260096020526120c760408620613d11565b6120d18285613c7b565b526120dc8184613c7b565b500161209d565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061211b57505050500390f35b91936020612156827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0600195979984950301865288516133a9565b960192019201859493919261210c565b806060602080938601015201612092565b50346102ee5760206003193601126102ee5760043567ffffffffffffffff81116109025760031960a091360301126102ee576004906121b4613c62565b507f690a7a40000000000000000000000000000000000000000000000000000000008152fd5b50346102ee5760206003193601126102ee5760043567ffffffffffffffff811161090257604060031982360301126109025760405190612219826132f8565b806004013567ffffffffffffffff8111610aa55761223d9060043691840101613676565b825260248101359067ffffffffffffffff8211610aa55760046122639236920101613676565b6020820190815261227261404f565b5191805b83518110156122e9578073ffffffffffffffffffffffffffffffffffffffff6122a160019387613c7b565b51166122ac81614cf9565b6122b8575b5001612276565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a1386122b1565b509051815b81518110156123845773ffffffffffffffffffffffffffffffffffffffff6123168284613c7b565b5116801561235c57907feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef60208361234e600195614995565b50604051908152a1016122ee565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b8280f35b50346102ee57806003193601126102ee57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102ee5760c06003193601126102ee576123d6613408565b6123de6134ae565b60643561ffff81168103610aa55760843567ffffffffffffffff81116108fe5761240c903690600401613568565b93909260a4359560028710156102ee5761189161243088888888604435888a613aa9565b604051918291826134e7565b50346102ee5760206003193601126102ee57602061247867ffffffffffffffff612464613497565b166000526007602052604060002054151590565b6040519015158152f35b50346102ee57806003193601126102ee57805473ffffffffffffffffffffffffffffffffffffffff81163303612521577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102ee57806003193601126102ee57600254600a546040805173ffffffffffffffffffffffffffffffffffffffff938416815292909116602083015290f35b50346102ee5761259936613619565b6125a59392919361404f565b67ffffffffffffffff82166125c7816000526007602052604060002054151590565b156125e657506125e392936125dd913691613596565b9061409a565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102ee57806003193601126102ee57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102ee5760206003193601126102ee5760043561ffff8116908181036108c7577fa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb916020916126b161404f565b7fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff75ffff00000000000000000000000000000000000000006003549260a01b16911617600355604051908152a180f35b50346102ee5760406003193601126102ee5761271b613497565b906024359067ffffffffffffffff82116102ee5760206124788461274236600487016135fb565b90613a0c565b50346102ee5760406003193601126102ee5760043567ffffffffffffffff811161090257806004019061010060031982360301126108c757612788613537565b9183604051612796816132ad565b52606482013592608483016127aa816139eb565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603612dac57602484019277ffffffffffffffff00000000000000000000000000000000612810856139d6565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115612da1578891612d82575b50612d5a5761289e611a19856139d6565b6128a7846139d6565b926128c460a48701946127426128bd8786613ffe565b3691613596565b15612d135761ffff16908115612c5e5767ffffffffffffffff6128e6866139d6565b1680895260056020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615888061295860408d2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391614a44565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff600354169182612a4f575b602088887ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0608067ffffffffffffffff6129e660446129df8e6139d6565b95016139eb565b9373ffffffffffffffffffffffffffffffffffffffff60405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a280604051612a46816132ad565b52604051908152f35b9082918994933b156108fe578785918a9389604051998a98899788967f5eff3bf70000000000000000000000000000000000000000000000000000000088526004880160609052612aa08680614687565b60648a0161010090526101648a0190612ab892613a6a565b94612ac2906134c5565b67ffffffffffffffff166084890152604401612add90613476565b73ffffffffffffffffffffffffffffffffffffffff1660a488015260c4870152612b0690613476565b73ffffffffffffffffffffffffffffffffffffffff1660e4860152612b2b9083614687565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c86840301610104870152612b609291613a6a565b612b6d60c48d0183614687565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c86840301610124870152612ba29291613a6a565b9060e48c01612bb091614687565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610144860152612be59291613a6a565b908b6024840152604483015203925af18015612c53576129e660446129df7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09560209960809667ffffffffffffffff96612c43575b505095506129a1565b81612c4d91613368565b38612c3a565b6040513d87823e3d90fd5b877f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8880612ce66002604067ffffffffffffffff612c9b8d6139d6565b16968781526008602052200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391614a44565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612981565b612d1d8483613ffe565b61200c6040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191613a6a565b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612d9b915060203d602011611ec757611eb98183613368565b3861288d565b6040513d8a823e3d90fd5b8573ffffffffffffffffffffffffffffffffffffffff611eef6024936139eb565b50346102ee5760206003193601126102ee5760043567ffffffffffffffff81116109025760031961010091360301126102ee5780600491604051612e10816132ad565b527f690a7a40000000000000000000000000000000000000000000000000000000008152fd5b50346102ee57806003193601126102ee57602073ffffffffffffffffffffffffffffffffffffffff60035416604051908152f35b50346102ee5760c06003193601126102ee57612e84613408565b50612e8d6134ae565b612e95613430565b506084359161ffff831683036102ee5760a4359067ffffffffffffffff82116102ee5760a063ffffffff8061ffff612edc8888612ed53660048b01613568565b5050613851565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346102ee57806003193601126102ee57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102ee57806003193601126102ee57604051600c8054808352908352909160208301917fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7915b818110612fa0576118918561243081870382613368565b8254845260209093019260019283019201612f89565b50346102ee5760406003193601126102ee57612fd0613497565b6024359182151583036102ee57610140613093612fed85856137ce565b61304360409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346102ee5760206003193601126102ee576020906130b2613408565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102ee57806003193601126102ee57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102ee57806003193601126102ee575061189160405161316d604082613368565b601781527f43435450546f6b656e506f6f6c20312e372e302d64657600000000000000000060208201526040519182916020835260208301906133a9565b905034610902576020600319360112610902576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036108c757602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115613283575b8115613259575b811561322f575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613228565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613221565b7f33171031000000000000000000000000000000000000000000000000000000008114915061321a565b6020810190811067ffffffffffffffff8211176132c957604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff8211176132c957604052565b60e0810190811067ffffffffffffffff8211176132c957604052565b6060810190811067ffffffffffffffff8211176132c957604052565b60a0810190811067ffffffffffffffff8211176132c957604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176132c957604052565b919082519283825260005b8481106133f35750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016133b4565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361342b57565b600080fd5b6064359073ffffffffffffffffffffffffffffffffffffffff8216820361342b57565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361342b57565b359073ffffffffffffffffffffffffffffffffffffffff8216820361342b57565b6004359067ffffffffffffffff8216820361342b57565b6024359067ffffffffffffffff8216820361342b57565b359067ffffffffffffffff8216820361342b57565b3590811515820361342b57565b602060408183019282815284518094520192019060005b81811061350b5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016134fe565b6024359061ffff8216820361342b57565b6044359061ffff8216820361342b57565b359061ffff8216820361342b57565b9181601f8401121561342b5782359167ffffffffffffffff831161342b576020838186019501011161342b57565b92919267ffffffffffffffff82116132c957604051916135de601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200184613368565b82948184528183011161342b578281602093846000960137010152565b9080601f8301121561342b5781602061361693359101613596565b90565b90604060031983011261342b5760043567ffffffffffffffff8116810361342b57916024359067ffffffffffffffff821161342b5761365a91600401613568565b9091565b67ffffffffffffffff81116132c95760051b60200190565b9080601f8301121561342b57813561368d8161365e565b9261369b6040519485613368565b81845260208085019260051b82010192831161342b57602001905b8282106136c35750505090565b602080916136d084613476565b8152019101906136b6565b9181601f8401121561342b5782359167ffffffffffffffff831161342b576020808501948460051b01011161342b57565b9181601f8401121561342b5782359167ffffffffffffffff831161342b576020808501948460081b01011161342b57565b6040519061374a8261334c565b60006080838281528260208201528260408201528260608201520152565b906040516137758161334c565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff916137e061373d565b506137e961373d565b5061381d57166000526008602052604060002090613616613811600261381661381186613768565b613f61565b9401613768565b16908160005260046020526138386138116040600020613768565b9160005260056020526136166138116040600020613768565b9061ffff8060035460a01c16911692831515928380946139ce575b6139a45767ffffffffffffffff16600052600b6020526040600020916040519261389584613314565b5463ffffffff81168452602084019563ffffffff8260201c168752604085019263ffffffff8360401c168452606086019163ffffffff8460601c168352608087019761ffff8560801c16895260ff60a089019561ffff8160901c16875260a01c1615801560c08a01526139895761392a575050505063ffffffff808061ffff9351169451169551169351169193929190600190565b81939750809294501061395957505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b50505050505092505050600090600090600090600090600090565b7f98d50fd70000000000000000000000000000000000000000000000000000000060005260046000fd5b50821561386c565b3567ffffffffffffffff8116810361342b5790565b3573ffffffffffffffffffffffffffffffffffffffff8116810361342b5790565b9067ffffffffffffffff61361692166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b519073ffffffffffffffffffffffffffffffffffffffff8216820361342b57565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff60035416958615613c4057613b449467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c4860191613a6a565b916002821015613c11578380600094819460a483015203915afa908115613c0557600091613b70575090565b3d8083833e613b7f8183613368565b8101906020818303126108c75780519067ffffffffffffffff8211610aa5570181601f820112156108c757805190613bb68261365e565b93613bc46040519586613368565b82855260208086019360051b8301019384116102ee5750602001905b828210613bed5750505090565b60208091613bfa84613a49565b815201910190613be0565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b5050505050505050604051613c56602082613368565b60008152600036813790565b60405190613c6f826132f8565b60606020838281520152565b8051821015613c8f5760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015613d07575b6020831014613cd857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691613ccd565b9060405191826000825492613d2584613cbe565b8084529360018116908115613d935750600114613d4c575b50613d4a92500383613368565b565b90506000929192526020600020906000915b818310613d77575050906020613d4a9282010138613d3d565b6020919350806001915483858901015201910190918492613d5e565b60209350613d4a9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138613d3d565b67ffffffffffffffff1660005260086020526136166004604060002001613d11565b9190811015613c8f5760081b0190565b35801515810361342b5790565b3561ffff8116810361342b5790565b3563ffffffff8116810361342b5790565b359063ffffffff8216820361342b57565b9190811015613c8f5760051b0190565b35906fffffffffffffffffffffffffffffffff8216820361342b57565b919082606091031261342b57604051613e8881613330565b6040613eb2818395613e99816134da565b8552613ea760208201613e53565b602086015201613e53565b910152565b6fffffffffffffffffffffffffffffffff613ef560408093613ed8816134da565b1515865283613ee960208301613e53565b16602087015201613e53565b16910152565b818110613f06575050565b60008155600101613efb565b81810292918115918404141715613f2557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b91908203918211613f2557565b613f6961373d565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691613fc66020850193613fc0613fb363ffffffff87511642613f54565b8560808901511690613f12565b90614597565b80821015613fdf57505b16825263ffffffff4216905290565b9050613fd0565b9081602091031261342b5751801515810361342b5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561342b570180359067ffffffffffffffff821161342b5760200191813603831361342b57565b73ffffffffffffffffffffffffffffffffffffffff60015416330361407057565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b908051156142da5767ffffffffffffffff815160208301209216918260005260086020526140cf816005604060002001614a0d565b156142965760005260096020526040600020815167ffffffffffffffff81116132c9576140fc8254613cbe565b601f8111614264575b506020601f821160011461419e5791614178827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361418e95600091614193575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90556040519182916020835260208301906133a9565b0390a2565b905084015138614147565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b81811061424c57509261418e9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614215575b5050811b01905561187d565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880614209565b9192602060018192868a0151815501940192016141ce565b61429090836000526020600020601f840160051c8101916020851061083157601f0160051c0190613efb565b38614105565b509061200c6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401526040602484015260448301906133a9565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906127109167ffffffffffffffff61431e602083016139d6565b166000908152600b602052604090209161ffff161561435157606061ffff61434d935460901c16910135613f12565b0490565b606061ffff61434d935460801c16910135613f12565b8151919291156144e9576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061448657613d4a91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836144e7604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408401511615801590614578575b61451757613d4a91926143aa565b6064836144e7604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020840151161515614509565b91908201809211613f2557565b906040519182815491828252602082019060005260206000209260005b8181106145d6575050613d4a92500383613368565b84548352600194850194879450602090930192016145c1565b67ffffffffffffffff16614610816000526007602052604060002054151590565b1561465a575033600052600d6020526040600020541561462c57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18236030181121561342b57016020813591019167ffffffffffffffff821161342b57813603831361342b57565b8054821015613c8f5760005260206000200190600090565b80548015614753577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061472482826146d7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000818152600760205260409020548015614887577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613f2557600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613f2557818103614818575b50505061480460066146ef565b600052600760205260006040812055600190565b61486f61482961483a9360066146d7565b90549060031b1c92839260066146d7565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260076020526040600020553880806147f7565b5050600090565b906001820191816000528260205260406000205490811515600014614961577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191808311613f255781547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111613f25578381614918950361492a575b5050506146ef565b60005260205260006040812055600190565b61494a61493a61483a93866146d7565b90549060031b1c928392866146d7565b905560005284602052604060002055388080614910565b50505050600090565b805490680100000000000000008210156132c9578161483a916001614991940181556146d7565b9055565b80600052600d602052604060002054156000146149ce576149b781600c61496a565b600c5490600052600d602052604060002055600190565b50600090565b806000526007602052604060002054156000146149ce576149f681600661496a565b600654906000526007602052604060002055600190565b60008281526001820160205260409020546148875780614a2f8360019361496a565b80549260005201602052604060002055600190565b9182549060ff8260a01c16158015614cf1575b614ceb576fffffffffffffffffffffffffffffffff82169160018501908154614a9c63ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642613f54565b9081614c4d575b5050848110614c015750838310614afd575050614ad26fffffffffffffffffffffffffffffffff928392613f54565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315614b955781614b1591613f54565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810193818511613f255773ffffffffffffffffffffffffffffffffffffffff94614b6091614597565b047fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611614cc157614c6892613fc09160801c90613f12565b80841015614cbc5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880614aa3565b614c73565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215614a57565b6000818152600d60205260409020548015614887577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111613f2557600c54907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211613f2557808203614d8f575b505050614d7b600c6146ef565b600052600d60205260006040812055600190565b614db1614da061483a93600c6146d7565b90549060031b1c928392600c6146d7565b9055600052600d602052604060002055388080614d6e56fea164736f6c634300081a000a",
}

var CCTPTokenPoolABI = CCTPTokenPoolMetaData.ABI

var CCTPTokenPoolBin = CCTPTokenPoolMetaData.Bin

func DeployCCTPTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, rmnProxy common.Address, router common.Address, cctpVerifier common.Address, allowedCallers []common.Address) (common.Address, *types.Transaction, *CCTPTokenPool, error) {
	parsed, err := CCTPTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPTokenPoolBin), backend, token, localTokenDecimals, rmnProxy, router, cctpVerifier, allowedCallers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCTPTokenPool{address: address, abi: *parsed, CCTPTokenPoolCaller: CCTPTokenPoolCaller{contract: contract}, CCTPTokenPoolTransactor: CCTPTokenPoolTransactor{contract: contract}, CCTPTokenPoolFilterer: CCTPTokenPoolFilterer{contract: contract}}, nil
}

type CCTPTokenPool struct {
	address common.Address
	abi     abi.ABI
	CCTPTokenPoolCaller
	CCTPTokenPoolTransactor
	CCTPTokenPoolFilterer
}

type CCTPTokenPoolCaller struct {
	contract *bind.BoundContract
}

type CCTPTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type CCTPTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type CCTPTokenPoolSession struct {
	Contract     *CCTPTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type CCTPTokenPoolCallerSession struct {
	Contract *CCTPTokenPoolCaller
	CallOpts bind.CallOpts
}

type CCTPTokenPoolTransactorSession struct {
	Contract     *CCTPTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type CCTPTokenPoolRaw struct {
	Contract *CCTPTokenPool
}

type CCTPTokenPoolCallerRaw struct {
	Contract *CCTPTokenPoolCaller
}

type CCTPTokenPoolTransactorRaw struct {
	Contract *CCTPTokenPoolTransactor
}

func NewCCTPTokenPool(address common.Address, backend bind.ContractBackend) (*CCTPTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(CCTPTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindCCTPTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPool{address: address, abi: abi, CCTPTokenPoolCaller: CCTPTokenPoolCaller{contract: contract}, CCTPTokenPoolTransactor: CCTPTokenPoolTransactor{contract: contract}, CCTPTokenPoolFilterer: CCTPTokenPoolFilterer{contract: contract}}, nil
}

func NewCCTPTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*CCTPTokenPoolCaller, error) {
	contract, err := bindCCTPTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolCaller{contract: contract}, nil
}

func NewCCTPTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*CCTPTokenPoolTransactor, error) {
	contract, err := bindCCTPTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolTransactor{contract: contract}, nil
}

func NewCCTPTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*CCTPTokenPoolFilterer, error) {
	contract, err := bindCCTPTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolFilterer{contract: contract}, nil
}

func bindCCTPTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCTPTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_CCTPTokenPool *CCTPTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPTokenPool.Contract.CCTPTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_CCTPTokenPool *CCTPTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.CCTPTokenPoolTransactor.contract.Transfer(opts)
}

func (_CCTPTokenPool *CCTPTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.CCTPTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.contract.Transfer(opts)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetAdvancedPoolHooks(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getAdvancedPoolHooks")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetAdvancedPoolHooks(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetAdvancedPoolHooks() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetAdvancedPoolHooks(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _CCTPTokenPool.Contract.GetAllAuthorizedCallers(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _CCTPTokenPool.Contract.GetAllAuthorizedCallers(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetCCTPVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getCCTPVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetCCTPVerifier() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetCCTPVerifier(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetCCTPVerifier() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetCCTPVerifier(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _CCTPTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _CCTPTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _CCTPTokenPool.Contract.GetDynamicConfig(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _CCTPTokenPool.Contract.GetDynamicConfig(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_CCTPTokenPool *CCTPTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _CCTPTokenPool.Contract.GetFee(&_CCTPTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _CCTPTokenPool.Contract.GetFee(&_CCTPTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _CCTPTokenPool.Contract.GetMinBlockConfirmation(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _CCTPTokenPool.Contract.GetMinBlockConfirmation(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _CCTPTokenPool.Contract.GetRemotePools(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _CCTPTokenPool.Contract.GetRemotePools(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _CCTPTokenPool.Contract.GetRemoteToken(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _CCTPTokenPool.Contract.GetRemoteToken(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CCTPTokenPool.Contract.GetRequiredCCVs(&_CCTPTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _CCTPTokenPool.Contract.GetRequiredCCVs(&_CCTPTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetRmnProxy(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetRmnProxy(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _CCTPTokenPool.Contract.GetSupportedChains(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _CCTPTokenPool.Contract.GetSupportedChains(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetToken() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetToken(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetToken(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _CCTPTokenPool.Contract.GetTokenDecimals(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _CCTPTokenPool.Contract.GetTokenDecimals(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CCTPTokenPool.Contract.GetTokenTransferFeeConfig(&_CCTPTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CCTPTokenPool.Contract.GetTokenTransferFeeConfig(&_CCTPTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _CCTPTokenPool.Contract.IsRemotePool(&_CCTPTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _CCTPTokenPool.Contract.IsRemotePool(&_CCTPTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _CCTPTokenPool.Contract.IsSupportedChain(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _CCTPTokenPool.Contract.IsSupportedChain(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _CCTPTokenPool.Contract.IsSupportedToken(&_CCTPTokenPool.CallOpts, token)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _CCTPTokenPool.Contract.IsSupportedToken(&_CCTPTokenPool.CallOpts, token)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) Owner() (common.Address, error) {
	return _CCTPTokenPool.Contract.Owner(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) Owner() (common.Address, error) {
	return _CCTPTokenPool.Contract.Owner(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPTokenPool.Contract.SupportsInterface(&_CCTPTokenPool.CallOpts, interfaceId)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPTokenPool.Contract.SupportsInterface(&_CCTPTokenPool.CallOpts, interfaceId)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) TypeAndVersion() (string, error) {
	return _CCTPTokenPool.Contract.TypeAndVersion(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _CCTPTokenPool.Contract.TypeAndVersion(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_CCTPTokenPool *CCTPTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.AcceptOwnership(&_CCTPTokenPool.TransactOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.AcceptOwnership(&_CCTPTokenPool.TransactOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.AddRemotePool(&_CCTPTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.AddRemotePool(&_CCTPTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_CCTPTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_CCTPTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyChainUpdates(&_CCTPTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyChainUpdates(&_CCTPTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_CCTPTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_CCTPTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, arg0 PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "lockOrBurn", arg0)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) LockOrBurn(arg0 PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.LockOrBurn(&_CCTPTokenPool.TransactOpts, arg0)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) LockOrBurn(arg0 PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.LockOrBurn(&_CCTPTokenPool.TransactOpts, arg0)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.LockOrBurn0(&_CCTPTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.LockOrBurn0(&_CCTPTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "releaseOrMint", arg0)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ReleaseOrMint(arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ReleaseOrMint(&_CCTPTokenPool.TransactOpts, arg0)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ReleaseOrMint(arg0 PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ReleaseOrMint(&_CCTPTokenPool.TransactOpts, arg0)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ReleaseOrMint0(&_CCTPTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ReleaseOrMint0(&_CCTPTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.RemoveRemotePool(&_CCTPTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.RemoveRemotePool(&_CCTPTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setDynamicConfig", router, rateLimitAdmin)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetDynamicConfig(&_CCTPTokenPool.TransactOpts, router, rateLimitAdmin)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetDynamicConfig(router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetDynamicConfig(&_CCTPTokenPool.TransactOpts, router, rateLimitAdmin)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setMinBlockConfirmation", minBlockConfirmation)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetMinBlockConfirmation(&_CCTPTokenPool.TransactOpts, minBlockConfirmation)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetMinBlockConfirmation(minBlockConfirmation uint16) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetMinBlockConfirmation(&_CCTPTokenPool.TransactOpts, minBlockConfirmation)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetRateLimitConfig(&_CCTPTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetRateLimitConfig(&_CCTPTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.TransferOwnership(&_CCTPTokenPool.TransactOpts, to)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.TransferOwnership(&_CCTPTokenPool.TransactOpts, to)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "updateAdvancedPoolHooks", newHook)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.UpdateAdvancedPoolHooks(&_CCTPTokenPool.TransactOpts, newHook)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) UpdateAdvancedPoolHooks(newHook common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.UpdateAdvancedPoolHooks(&_CCTPTokenPool.TransactOpts, newHook)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.WithdrawFeeTokens(&_CCTPTokenPool.TransactOpts, feeTokens, recipient)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.WithdrawFeeTokens(&_CCTPTokenPool.TransactOpts, feeTokens, recipient)
}

type CCTPTokenPoolAdvancedPoolHooksUpdatedIterator struct {
	Event *CCTPTokenPoolAdvancedPoolHooksUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolAdvancedPoolHooksUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolAdvancedPoolHooksUpdated)
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
		it.Event = new(CCTPTokenPoolAdvancedPoolHooksUpdated)
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

func (it *CCTPTokenPoolAdvancedPoolHooksUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolAdvancedPoolHooksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolAdvancedPoolHooksUpdated struct {
	OldHook common.Address
	NewHook common.Address
	Raw     types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*CCTPTokenPoolAdvancedPoolHooksUpdatedIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolAdvancedPoolHooksUpdatedIterator{contract: _CCTPTokenPool.contract, event: "AdvancedPoolHooksUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "AdvancedPoolHooksUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolAdvancedPoolHooksUpdated)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseAdvancedPoolHooksUpdated(log types.Log) (*CCTPTokenPoolAdvancedPoolHooksUpdated, error) {
	event := new(CCTPTokenPoolAdvancedPoolHooksUpdated)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "AdvancedPoolHooksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolAuthorizedCallerAddedIterator struct {
	Event *CCTPTokenPoolAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolAuthorizedCallerAdded)
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
		it.Event = new(CCTPTokenPoolAuthorizedCallerAdded)
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

func (it *CCTPTokenPoolAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*CCTPTokenPoolAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolAuthorizedCallerAddedIterator{contract: _CCTPTokenPool.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolAuthorizedCallerAdded)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseAuthorizedCallerAdded(log types.Log) (*CCTPTokenPoolAuthorizedCallerAdded, error) {
	event := new(CCTPTokenPoolAuthorizedCallerAdded)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolAuthorizedCallerRemovedIterator struct {
	Event *CCTPTokenPoolAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolAuthorizedCallerRemoved)
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
		it.Event = new(CCTPTokenPoolAuthorizedCallerRemoved)
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

func (it *CCTPTokenPoolAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*CCTPTokenPoolAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolAuthorizedCallerRemovedIterator{contract: _CCTPTokenPool.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolAuthorizedCallerRemoved)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*CCTPTokenPoolAuthorizedCallerRemoved, error) {
	event := new(CCTPTokenPoolAuthorizedCallerRemoved)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolChainAddedIterator struct {
	Event *CCTPTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolChainAdded)
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
		it.Event = new(CCTPTokenPoolChainAdded)
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

func (it *CCTPTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*CCTPTokenPoolChainAddedIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolChainAddedIterator{contract: _CCTPTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolChainAdded)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseChainAdded(log types.Log) (*CCTPTokenPoolChainAdded, error) {
	event := new(CCTPTokenPoolChainAdded)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolChainRemovedIterator struct {
	Event *CCTPTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolChainRemoved)
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
		it.Event = new(CCTPTokenPoolChainRemoved)
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

func (it *CCTPTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*CCTPTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolChainRemovedIterator{contract: _CCTPTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolChainRemoved)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseChainRemoved(log types.Log) (*CCTPTokenPoolChainRemoved, error) {
	event := new(CCTPTokenPoolChainRemoved)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _CCTPTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _CCTPTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolDynamicConfigSetIterator struct {
	Event *CCTPTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolDynamicConfigSet)
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
		it.Event = new(CCTPTokenPoolDynamicConfigSet)
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

func (it *CCTPTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolDynamicConfigSet struct {
	Router         common.Address
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolDynamicConfigSetIterator{contract: _CCTPTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolDynamicConfigSet)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*CCTPTokenPoolDynamicConfigSet, error) {
	event := new(CCTPTokenPoolDynamicConfigSet)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolFeeTokenWithdrawnIterator struct {
	Event *CCTPTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(CCTPTokenPoolFeeTokenWithdrawn)
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

func (it *CCTPTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolFeeTokenWithdrawn struct {
	Receiver common.Address
	FeeToken common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPTokenPoolFeeTokenWithdrawnIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolFeeTokenWithdrawnIterator{contract: _CCTPTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", receiverRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolFeeTokenWithdrawn)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*CCTPTokenPoolFeeTokenWithdrawn, error) {
	event := new(CCTPTokenPoolFeeTokenWithdrawn)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolInboundRateLimitConsumedIterator struct {
	Event *CCTPTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(CCTPTokenPoolInboundRateLimitConsumed)
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

func (it *CCTPTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolInboundRateLimitConsumedIterator{contract: _CCTPTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolInboundRateLimitConsumed)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolInboundRateLimitConsumed, error) {
	event := new(CCTPTokenPoolInboundRateLimitConsumed)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolLockedOrBurnedIterator struct {
	Event *CCTPTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolLockedOrBurned)
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
		it.Event = new(CCTPTokenPoolLockedOrBurned)
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

func (it *CCTPTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolLockedOrBurnedIterator{contract: _CCTPTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolLockedOrBurned)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*CCTPTokenPoolLockedOrBurned, error) {
	event := new(CCTPTokenPoolLockedOrBurned)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolMinBlockConfirmationSetIterator struct {
	Event *CCTPTokenPoolMinBlockConfirmationSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolMinBlockConfirmationSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolMinBlockConfirmationSet)
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
		it.Event = new(CCTPTokenPoolMinBlockConfirmationSet)
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

func (it *CCTPTokenPoolMinBlockConfirmationSetIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolMinBlockConfirmationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolMinBlockConfirmationSet struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*CCTPTokenPoolMinBlockConfirmationSetIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolMinBlockConfirmationSetIterator{contract: _CCTPTokenPool.contract, event: "MinBlockConfirmationSet", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolMinBlockConfirmationSet) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "MinBlockConfirmationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolMinBlockConfirmationSet)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseMinBlockConfirmationSet(log types.Log) (*CCTPTokenPoolMinBlockConfirmationSet, error) {
	event := new(CCTPTokenPoolMinBlockConfirmationSet)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "MinBlockConfirmationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *CCTPTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(CCTPTokenPoolOutboundRateLimitConsumed)
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

func (it *CCTPTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolOutboundRateLimitConsumedIterator{contract: _CCTPTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolOutboundRateLimitConsumed)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolOutboundRateLimitConsumed, error) {
	event := new(CCTPTokenPoolOutboundRateLimitConsumed)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolOwnershipTransferRequestedIterator struct {
	Event *CCTPTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolOwnershipTransferRequested)
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
		it.Event = new(CCTPTokenPoolOwnershipTransferRequested)
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

func (it *CCTPTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolOwnershipTransferRequestedIterator{contract: _CCTPTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolOwnershipTransferRequested)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*CCTPTokenPoolOwnershipTransferRequested, error) {
	event := new(CCTPTokenPoolOwnershipTransferRequested)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolOwnershipTransferredIterator struct {
	Event *CCTPTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolOwnershipTransferred)
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
		it.Event = new(CCTPTokenPoolOwnershipTransferred)
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

func (it *CCTPTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolOwnershipTransferredIterator{contract: _CCTPTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolOwnershipTransferred)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*CCTPTokenPoolOwnershipTransferred, error) {
	event := new(CCTPTokenPoolOwnershipTransferred)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolRateLimitConfiguredIterator struct {
	Event *CCTPTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolRateLimitConfigured)
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
		it.Event = new(CCTPTokenPoolRateLimitConfigured)
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

func (it *CCTPTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolRateLimitConfiguredIterator{contract: _CCTPTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolRateLimitConfigured)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*CCTPTokenPoolRateLimitConfigured, error) {
	event := new(CCTPTokenPoolRateLimitConfigured)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolReleasedOrMintedIterator struct {
	Event *CCTPTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolReleasedOrMinted)
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
		it.Event = new(CCTPTokenPoolReleasedOrMinted)
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

func (it *CCTPTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolReleasedOrMintedIterator{contract: _CCTPTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolReleasedOrMinted)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*CCTPTokenPoolReleasedOrMinted, error) {
	event := new(CCTPTokenPoolReleasedOrMinted)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolRemotePoolAddedIterator struct {
	Event *CCTPTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolRemotePoolAdded)
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
		it.Event = new(CCTPTokenPoolRemotePoolAdded)
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

func (it *CCTPTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolRemotePoolAddedIterator{contract: _CCTPTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolRemotePoolAdded)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*CCTPTokenPoolRemotePoolAdded, error) {
	event := new(CCTPTokenPoolRemotePoolAdded)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolRemotePoolRemovedIterator struct {
	Event *CCTPTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolRemotePoolRemoved)
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
		it.Event = new(CCTPTokenPoolRemotePoolRemoved)
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

func (it *CCTPTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolRemotePoolRemovedIterator{contract: _CCTPTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolRemotePoolRemoved)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*CCTPTokenPoolRemotePoolRemoved, error) {
	event := new(CCTPTokenPoolRemotePoolRemoved)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *CCTPTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(CCTPTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *CCTPTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _CCTPTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolTokenTransferFeeConfigDeleted)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*CCTPTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(CCTPTokenPoolTokenTransferFeeConfigDeleted)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *CCTPTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(CCTPTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *CCTPTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _CCTPTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolTokenTransferFeeConfigUpdated)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*CCTPTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(CCTPTokenPoolTokenTransferFeeConfigUpdated)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (CCTPTokenPoolAdvancedPoolHooksUpdated) Topic() common.Hash {
	return common.HexToHash("0xbaff46844acf36d6ee996f489a1a288709c4542bd33cd557770afd267d581209")
}

func (CCTPTokenPoolAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (CCTPTokenPoolAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (CCTPTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (CCTPTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (CCTPTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x22a0dbb8195755fbfc99667a86ae684c568e9dfbb1eccf7f90084e6166447970")
}

func (CCTPTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (CCTPTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (CCTPTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (CCTPTokenPoolMinBlockConfirmationSet) Topic() common.Hash {
	return common.HexToHash("0xa7f8dbba8cdb126ce4a0e7939ec58e0161b70d808b585dd651d68e59d27e11fb")
}

func (CCTPTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (CCTPTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCTPTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (CCTPTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (CCTPTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (CCTPTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (CCTPTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (CCTPTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (CCTPTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_CCTPTokenPool *CCTPTokenPool) Address() common.Address {
	return _CCTPTokenPool.address
}

type CCTPTokenPoolInterface interface {
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

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetMinBlockConfirmation(opts *bind.TransactOpts, minBlockConfirmation uint16) (*types.Transaction, error)

	SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateAdvancedPoolHooks(opts *bind.TransactOpts, newHook common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterAdvancedPoolHooksUpdated(opts *bind.FilterOpts) (*CCTPTokenPoolAdvancedPoolHooksUpdatedIterator, error)

	WatchAdvancedPoolHooksUpdated(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAdvancedPoolHooksUpdated) (event.Subscription, error)

	ParseAdvancedPoolHooksUpdated(log types.Log) (*CCTPTokenPoolAdvancedPoolHooksUpdated, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*CCTPTokenPoolAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*CCTPTokenPoolAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*CCTPTokenPoolAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*CCTPTokenPoolAuthorizedCallerRemoved, error)

	FilterChainAdded(opts *bind.FilterOpts) (*CCTPTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*CCTPTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*CCTPTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*CCTPTokenPoolChainRemoved, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CCTPTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, receiver []common.Address, feeToken []common.Address) (*CCTPTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolFeeTokenWithdrawn, receiver []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCTPTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*CCTPTokenPoolLockedOrBurned, error)

	FilterMinBlockConfirmationSet(opts *bind.FilterOpts) (*CCTPTokenPoolMinBlockConfirmationSetIterator, error)

	WatchMinBlockConfirmationSet(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolMinBlockConfirmationSet) (event.Subscription, error)

	ParseMinBlockConfirmationSet(log types.Log) (*CCTPTokenPoolMinBlockConfirmationSet, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPTokenPoolOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*CCTPTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*CCTPTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*CCTPTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*CCTPTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*CCTPTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*CCTPTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*CCTPTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
