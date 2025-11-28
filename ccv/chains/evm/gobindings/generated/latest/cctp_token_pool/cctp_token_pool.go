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

type TokenPoolCustomBlockConfirmationRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var CCTPTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCTPVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentCustomBlockConfirmationRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationRateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultFinalityRateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFoundForVerifier\",\"inputs\":[{\"name\":\"verifierVersion\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"InvalidCCTPVerifier\",\"inputs\":[{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutboundImplementationNotFoundForVerifier\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103925760c081614e53803803809161002082856103e3565b8339810103126103925780516001600160a01b03811691908281036103925761004b6020830161041c565b6100576040840161042a565b916100646060850161042a565b9061007d60a06100766080880161042a565b960161042a565b9533156103d257600180546001600160a01b03191633179055801580156103c1575b80156103b0575b61039f5760049260209260805260c0526040519283809263313ce56760e01b82525afa6000918161035e575b50610333575b5060a0526001600160a01b0390811660e052600280546001600160a01b031916929091169190911790556040516301ffc9a760e01b602082810182815260248085019390935291835291600091906101316044826103e3565b519084617530fa6000513d82610327575b508161031d575b50806102b6575b80610250575b1561022f5761010052604051614a14908161043f82396080518181816117e30152818161197a01528181611cea01528181611da2015281816121440152818161228801528181612bdb01528181612d9c01528181612e0a01528181612ea301528181612ff001528181613187015281816133af01526133fc015260a05181613361015260c051818181610b8d0152818161184b015281816121ab01528181612c440152613059015260e05181818161192b015281816122f501526139ef015261010051818181611a2a0152818161239e0152612b050152f35b630d9758df60e01b60009081526001600160a01b0391909116600452602490fd5b5060206000604051828101906301ffc9a760e01b8252635627ff7160e01b6024820152602481526102826044826103e3565b519084617530fa6000513d826102aa575b50816102a0575b50610156565b905015153861029a565b60201115915038610293565b5060206000604051828101906301ffc9a760e01b825263ffffffff60e01b6024820152602481526102e86044826103e3565b519084617530fa6000513d82610311575b5081610307575b5015610150565b9050151538610300565b602011159150386102f9565b9050151538610149565b60201115915038610142565b60ff1660ff821681810361034757506100d8565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610397575b8161037a602093836103e3565b810103126103925761038b9061041c565b90386100d2565b600080fd5b3d915061036d565b630a64406560e11b60005260046000fd5b506001600160a01b038316156100a6565b506001600160a01b0386161561009f565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761040657604052565b634e487b7160e01b600052604160045260246000fd5b519060ff8216820361039257565b51906001600160a01b03821682036103925756fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461348157508063181f5a771461342057806321df0da7146133dc578063240028e81461338557806324f65ee7146133475780632c063404146132af5780633907753714612f7f578063489a68f214612b6f5780634c5ef0ed14612b29578063615521a714612ae557806362ddd3c414612a785780637437ff9f14612a3457806379ba5097146129b05780638926f54f1461296b57806389720a62146128b25780638da5cb5b1461288b578063962d4020146126925780639751f884146126315780639a4575b9146120db578063a42a7b8b14611fae578063acfecf9114611ebc578063b1c71c6514611756578063b794658014611719578063bb6bb5a71461141b578063c4bffe2b14611310578063c7230a6014611064578063d8aa3f4014610f2c578063dc04fa1f14610bb1578063dc0bd97114610b6d578063ded8d95614610a6c578063e8a1da171461030b578063f2fde38b146102875763fdf168751461018c57600080fd5b34610284576060366003190112610284576101a5613605565b906101ae61368a565b604435926001600160a01b038416808503610280576101cb613f08565b6001600160a01b0382168015610271576002805475ffffffffffffffffffffffffffffffffffffffffffff191690911760a085901b61ffff60a01b16179055600980546001600160a01b0319169091179055604080516001600160a01b03928316815261ffff939093166020840152931692810192909252907fba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a1179080606081010390a180f35b630a64406560e11b8552600485fd5b8380fd5b80fd5b5034610284576020366003190112610284576001600160a01b036102a9613605565b6102b1613f08565b163381146102fc5781546001600160a01b031916811782556001546001600160a01b03167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b636d6c4ee560e11b8252600482fd5b5034610284576040366003190112610284576004356001600160401b038111610a685761033c903690600401613780565b916024356001600160401b038111610a685761035c903690600401613780565b939092610367613f08565b82915b8083106108ec575050508063ffffffff42169161011e19843603015b858210156108e8578160051b85013581811215610280578501906101208236031261028057604051956103b88761356d565b6103c183613676565b875260208301356001600160401b0381116108e45783019536601f880112156108e4578635966103f088613993565b976103fe604051998a613588565b8089526020808a019160051b830101903682116108e05760208301905b8282106108ae57505050506020880196875260408401356001600160401b0381116108aa5761044d903690860161371f565b9860408901998a526104776104653660608801613bbf565b9560608b0196875260c0369101613bbf565b9660808a0197885261048986516142d1565b61049388516142d1565b8a51511561089b576104ae6001600160401b038b5116614739565b1561087e576001600160401b038a5116815260076020526040812061059287516001600160801b03604082015116906105656001600160801b03602083015116915115158360806040516105018161356d565b858152602081018c905260408101849052606081018690520152855460ff60a01b91151560a01b9190911674ffffffffffffffffffffffffffffffffffffffffff199091166001600160801b0384161763ffffffff60801b60808b901b1617178555565b60809190911b6fffffffffffffffffffffffffffffffff19166001600160801b0391909116176001830155565b61066289516001600160801b03604082015116906106356001600160801b03602083015116915115158360806040516105ca8161356d565b858152602081018c90526040810184905260608101869052015260028601805460ff60a01b92151560a01b9290921674ffffffffffffffffffffffffffffffffffffffffff199092166001600160801b0385161763ffffffff60801b60808c901b1617919091179055565b60809190911b6fffffffffffffffffffffffffffffffff19166001600160801b0391909116176003830155565b60048c519101908051906001600160401b03821161086a576106848354613ce2565b601f811161082f575b50602090601f83116001146107cc576106be92918591836107c1575b50508160011b916000199060031b1c19161790565b90555b805b895180518210156106f857906106f26001926106eb838f6001600160401b0390511692613cce565b5190613f2d565b016106c3565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2929593966107b36001600160401b03600197949c511692519351915161078861075c604051968796875261010060208801526101008701906135c4565b9360408601906001600160801b0360408092805115158552826020820151166020860152015116910152565b60a08401906001600160801b0360408092805115158552826020820151166020860152015116910152565b0390a1019093949291610386565b0151905038806106a9565b8385528185209190601f198416865b81811061081757509084600195949392106107fe575b505050811b0190556106c1565b015160001960f88460031b161c191690553880806107f1565b929360206001819287860151815501950193016107db565b61085a9084865260208620601f850160051c81019160208610610860575b601f0160051c0190613e7e565b3861068d565b909150819061084d565b634e487b7160e01b84526041600452602484fd5b8951631d5ad3c560e01b82526001600160401b0316600452602490fd5b630a64406560e11b8152600490fd5b8680fd5b81356001600160401b0381116108dc576020916108d1839283369189010161371f565b81520191019061041b565b8a80fd5b8880fd5b8580fd5b8280f35b909291936001600160401b0361090b610906878588613b68565b61392f565b169561091687614594565b15610a5457868452600760205261093260056040862001614531565b94845b865181101561096b5760019089875260076020526109646005604089200161095d838b613cce565b519061467c565b5001610935565b50939450949095808552600760205260056040862086815586600182015586600282015586600382015586600482016109a48154613ce2565b80610a13575b50505001805490868155816109f5575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101919094939461036a565b865260208620908101905b818110156109ba57868155600101610a00565b601f8111600114610a295750555b8638806109aa565b81835260208320610a4491601f01861c810190600101613e7e565b8082528160208120915555610a21565b631e670e4b60e01b84526004879052602484fd5b5080fd5b50346102845760203660031901126102845761014090610b6b610ada610ac760406001600160401b03610a9d613660565b610aa5613c16565b50610aae613c16565b5016948581526003602052610acc610ac7838320613c41565b61415c565b958152600460205220613c41565b610b2460405180946001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b503461028457806003193601126102845760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610284576040366003190112610284576004356001600160401b038111610a685736602382011215610a685780600401356001600160401b038111610f28576024820191602436918360081b010111610f28576024356001600160401b03811161028057610c25903690600401613780565b919092610c30613f08565b845b828110610c9b57505050825b818110610c49578380f35b806001600160401b03610c626109066001948688613b68565b16808652600a6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610c3e565b610ca9610906828585613e3d565b610cb4828585613e3d565b90602082019060e0830190610cc882613de1565b15610f0d5760a0840161271061ffff610ce083613e4d565b161015610efe5760c085019161271061ffff610cfb85613e4d565b161015610edf5763ffffffff610d1086613e5c565b1615610ec4576001600160401b031694858c52600a60205260408c2090610d3686613e5c565b63ffffffff168254926040830191610d4d83613e5c565b60201b67ffffffff0000000016906060850194610d6986613e5c565b60401b6bffffffff0000000000000000169060800196610d8888613e5c565b60601b63ffffffff60601b1691610d9e8a613e4d565b60801b61ffff60801b1693610db28c613e4d565b60901b61ffff60901b169561ffff60901b199361ffff60801b199263ffffffff60601b19916bffffffffffffffffffffffff19161716171617161717178155610dfa87613de1565b815460ff60a01b191690151560a01b60ff60a01b1617905560405196610e1f90613e6d565b63ffffffff168752610e3090613e6d565b63ffffffff166020870152610e4490613e6d565b63ffffffff166040860152610e5890613e6d565b63ffffffff166060850152610e6c906136ac565b61ffff166080840152610e7e906136ac565b61ffff1660a0830152610e9090613b9e565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c32565b631233226560e01b8c526001600160401b031660045260248bfd5b60248c61ffff610eee86613e4d565b634af9a8bd60e11b835216600452fd5b8a61ffff610eee602493613e4d565b631233226560e01b8a526001600160401b0316600452602489fd5b8280fd5b503461028457608036600319011261028457610f46613605565b50610f4f61364a565b610f5761369b565b506064356001600160401b038111610f2857916001600160401b03604092610f8560e09536906004016136bb565b50508260c08551610f9581613552565b82815282602082015282878201528260608201528260808201528260a08201520152168152600a6020522060405190610fcd82613552565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b5034610284576040366003190112610284576004356001600160401b038111610a6857611095903690600401613780565b602435906001600160a01b0382169081830361130c576110b3613f08565b845b8181106110c0578580f35b8060206001600160a01b036110e06110db602495878b613b68565b613943565b16604051938480926370a0823160e01b82523060048301525afa80156113015784869185948a916112c9575b5080611120575b50505060019150016110b5565b886111e56001600160a01b0361113a6110db888a86613b68565b60405163a9059cbb60e01b602082019081526001600160a01b0398909816602482015260448082018790528152918e9116611176606484613588565b81806040998a51956111888c88613588565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d156112c0573d6111c9816135a9565b906111d68b519283613588565b8152809360203d92013e61496e565b80518061123d575b50507f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e916001600160a01b0361122b6110db8860019a602096613b68565b169451908152a3829150848438611113565b816020935083929496979850611257955001019101613ebe565b156112695790848493928838806111ed565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b6060925061496e565b94505050506020823d82116112f9575b816112e660209383613588565b810103126108aa5784848493513861110c565b3d91506112d9565b6040513d89823e3d90fd5b8480fd5b5034610284578060031936011261028457604051906005548083528260208101600584526020842092845b81811061140257505061135092500383613588565b815161137461135e82613993565b9161136c6040519384613588565b808352613993565b602082019290601f1901368437805b84518110156113b457806001600160401b036113a160019388613cce565b51166113ad8286613cce565b5201611383565b50925090604051928392602084019060208552518091526040840192915b8181106113e0575050500390f35b82516001600160401b03168452859450602093840193909201916001016113d2565b845483526001948501948794506020909301920161133b565b5034610284576020366003190112610284576004356001600160401b038111610a685736602382011215610a68578060040135906001600160401b038211610f285736602460e0840283010111610f285761147461410a565b825b828110156117155760e0810282016001600160401b036114986024830161392f565b16906114b1826000526006602052604060002054151590565b156117015760e0600193926115c5886115bb856115ab60447f20ae59542ddd78610f62f9d2c9dcd658f8b6a5b45a0f03e71e16614e89dda8369801916114ff6114fa3685613bbf565b6142d1565b868552600360205261155060408620805463ffffffff8160801c161590816116ec575b816116dd575b816116cb575b816116bc575b50806116ad575b611667575b61154a3686613bbf565b906143a8565b604060a48201956115646114fa3689613bbf565b88815260046020522090815463ffffffff8160801c16159081611652575b81611643575b81611631575b81611622575b5080611613575b6115cc575b5061154a3686613bbf565b6040519485526020850190613e02565b6080830190613e02565ba101611476565b6115e060c46001600160801b039201613dee565b825463ffffffff60801b4260801b166001600160a01b03199091169190921663ffffffff60801b191617178155386115a0565b5061161d86613de1565b61159b565b60ff915060a01c161538611594565b6001600160801b03811615915061158e565b838e015460801c159150611588565b838e01546001600160801b0316159150611582565b6001600160801b0361167b60648501613dee565b825463ffffffff60801b4260801b166001600160a01b03199091169190921663ffffffff60801b191617178155611540565b506116b785613de1565b61153b565b60ff915060a01c161538611534565b6001600160801b03811615915061152e565b828f015460801c159150611528565b828f01546001600160801b0316159150611522565b631e670e4b60e01b86526004829052602486fd5b8380f35b50346102845760203660031901126102845761175261173e611739613660565b613dc0565b6040519182916020835260208301906135c4565b0390f35b503461028457606036600319011261028457600435906001600160401b038211610284573682900360031901916004810160a08412610f285761179761368a565b6044356001600160401b03811161130c576117b96117c99136906004016136bb565b6117c1613c95565b5036916136e8565b9460848401956117d887613943565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611e9857602485019467ffffffffffffffff60801b6118248761392f565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e3d578891611e69575b50611e5a576001600160401b036118918761392f565b166118a9816000526006602052604060002054151590565b15611e485760206001600160a01b03600254169160246040518094819363a8d87a3b60e01b835260048301525afa8015611e3d578890611dfd575b6001600160a01b039150163303611dea5761ffff841691606482013591908315611d4e5761ffff60025460a01c1680611c8e575b50889998959697985b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283611b7b575b5050505050505050611963916141ff565b61196c8261392f565b604080516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001681523360208201529081018390526001600160401b0391909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2611a1e60206119e98461392f565b6040518093819263958021a760e01b8352600483016001600160401b0360609216815260406020820152600060408201520190565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015611b0c578490611b3b575b6001600160a01b03915016918215611b17576020611a7b61173960049361392f565b936040519283809263fe163eed60e01b82525afa908115611b0c5790611aab91611ad39591611add575b506141cf565b60405192611ab884613537565b835260208301526040519283926040845260408401906137e0565b9060208301520390f35b611aff915060203d602011611b05575b611af78183613588565b810190613cae565b38611aa5565b503d611aed565b6040513d86823e3d90fd5b836001600160401b03611b2b60249361392f565b6395c6c70560e01b835216600452fd5b506020813d602011611b73575b81611b5560209383613588565b8101031261028057611b6e6001600160a01b03916139aa565b611a59565b3d9150611b48565b833b15611c8a57604051632e1d7be560e11b815260606004820152968a3590601e19018112156108e05781600491010194602086359601936001600160401b038711611c86578636038513611c865789976001600160a01b03611c308f94968c9a9783611c206044611c4c9a8f6084819f829f611c106001600160401b0392611c169260a060648801526101048701916139be565b9e613676565b1691015201613636565b1660a48a015260c4890152613636565b1660e486015260248501528382036003190160448501526135c4565b03915afa8015611c7b57611c66575b808080808080611952565b81611c7091613588565b610280578338611c5b565b6040513d84823e3d90fd5b8980fd5b8780fd5b808510611d37575088996001600160401b03611cad8a9b98999a61392f565b1680885260036020527f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac209793288580611d1260408c206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916147ee565b604080516001600160a01b039290921682526020820192909252a29998979695611918565b637911d95b60e01b8a526004859052602452604489fd5b88996001600160401b03611d658a9b98999a61392f565b1680885260076020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448580611dca60408c206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916147ee565b604080516001600160a01b039290921682526020820192909252a2611921565b63728fe07b60e01b875233600452602487fd5b506020813d602011611e35575b81611e1760209383613588565b81010312611c8a57611e306001600160a01b03916139aa565b6118e4565b3d9150611e0a565b6040513d8a823e3d90fd5b6354c8163f60e11b8852600452602487fd5b630a75a23b60e31b8752600487fd5b611e8b915060203d602011611e91575b611e838183613588565b810190613ebe565b3861187b565b503d611e79565b6024866001600160a01b03611eac8a613943565b63961c9a4f60e01b835216600452fd5b5034610284576001600160401b03611ed33661373d565b929091611ede613f08565b1691611ef7836000526006602052604060002054151590565b15611f9a578284526007602052611f2660056040862001611f193684866136e8565b602081519101209061467c565b15611f6b57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f656040519283926020845260208401916139be565b0390a280f35b82611f9683604051938493631d3c8f1f60e21b855260048501526040602485015260448401916139be565b0390fd5b631e670e4b60e01b84526004839052602484fd5b5034610284576020366003190112610284576001600160401b03611fd0613660565b1681526007602052611fe760056040832001614531565b8051611ff281613993565b906120006040519283613588565b80825261200f601f1991613993565b01835b8181106120ca575050825b8251811015612063578061203360019285613cce565b518552600860205261204760408620613d1c565b6120518285613cce565b5261205c8184613cce565b500161201d565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061209b57505050500390f35b919360019193955060206120ba8192603f198a820301865288516135c4565b960192019201859493919261208c565b806060602080938601015201612012565b5034610284576020366003190112610284576004356001600160401b038111610a6857600319813603019060a08212610f2857612116613c95565b50602090604051926121288385613588565b848452608482019061213982613943565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361261d57602483019467ffffffffffffffff60801b6121858761392f565b60801b1660405190632cbc26bb60e01b8252600482015285816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e3d578891612600575b50611e5a576001600160401b036121f18761392f565b16612209816000526006602052604060002054151590565b15611e4857856001600160a01b03600254169160246040518094819363a8d87a3b60e01b835260048301525afa8015611e3d5788906125c5575b6001600160a01b039150163303611dea57869060648501359485946001600160401b0361226f8a61392f565b16808552600789526122b0604086206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016998a916147ee565b604080516001600160a01b038a168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a26001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169182612493575b5050505050506001600160401b037ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10916123826123588761392f565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b0390a2612392816119e98461392f565b03816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015611b0c578490612458575b6001600160a01b03915016918215611b1757816123ee61173960049361392f565b936040519283809263fe163eed60e01b82525afa908115611b0c5761241a9293949161244157506141cf565b6040519261242784613537565b8352818301526117526040519282849384528301906137e0565b611aff9150833d8511611b0557611af78183613588565b508181813d831161248c575b61246e8183613588565b81010312610280576124876001600160a01b03916139aa565b6123cd565b503d612464565b823b1561130c5760405195632e1d7be560e11b875260606004880152816004013590601e19018112156108e457816004910101918983359301916001600160401b0384116108aa5783360383136108aa5787956001600160a01b03612547612564948f94612536604485926001600160401b0361252a8f9d8e9d612524918f8060a0606461010493015201916139be565b9a613676565b1660848c015201613636565b1660a48801528c60c4880152613636565b1660e48501528760248501526003198483030160448501526135c4565b03915afa80156125ba5761257d575b858180808061231c565b856125b27ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1093976001600160401b0393613588565b959150612573565b6040513d88823e3d90fd5b508581813d83116125f9575b6125db8183613588565b81010312611c8a576125f46001600160a01b03916139aa565b612243565b503d6125d1565b6126179150863d8811611e9157611e838183613588565b386121db565b6024866001600160a01b03611eac85613943565b5034610284576020366003190112610284576040610140916001600160401b03612659613660565b612661613c16565b5061266a613c16565b50168152600760205220610b6b610ada610ac7600261268b610ac786613c41565b9401613c41565b5034610284576060366003190112610284576004356001600160401b038111610a68576126c3903690600401613780565b906024356001600160401b038111610280576126e39036906004016137b0565b6044929192356001600160401b0381116108e4576127059036906004016137b0565b909261270f61410a565b828614801590612881575b61287257865b86811061272b578780f35b612739610906828985613b68565b612744828689613b8e565b906001600160401b0361276e61276661275e86898c613b8e565b943690613bbf565b933690613bbf565b911691612788836000526006602052604060002054151590565b1561285e57600193926128578361282c848f7f73d6dce40db73cbddae4b9ce52576043a1644e08c2702236273d71077435fa1697600260406127fb936127cf60e09b6142d1565b89815260076020526127e3868383206143a8565b6127ec846142d1565b898152600760205220016143a8565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565ba101612720565b631e670e4b60e01b8b52600483905260248bfd5b632b477e7160e11b8752600487fd5b508186141561271a565b503461028457806003193601126102845760206001600160a01b0360015416604051908152f35b50346102845760c0366003190112610284576128cc613605565b6128d461364a565b9060643561ffff81168103610280576084356001600160401b03811161130c576129029036906004016136bb565b9160a4359360028510156108aa5761291d95604435916139df565b90604051918291602083016020845282518091526020604085019301915b818110612949575050500390f35b82516001600160a01b031684528594506020938401939092019160010161293b565b50346102845760203660031901126102845760206129a66001600160401b03612992613660565b166000526006602052604060002054151590565b6040519015158152f35b503461028457806003193601126102845780546001600160a01b0381163303612a255760015490336001600160a01b03198316176001556001600160a01b03191682556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b63015aa1e360e11b8252600482fd5b5034610284578060031936011261028457600254600954604080516001600160a01b03808516825260a09490941c61ffff1660208201529290911690820152606090f35b503461028457612a873661373d565b612a9393929193613f08565b6001600160401b038216612ab4816000526006602052604060002054151590565b15612ad35750612ad09293612aca9136916136e8565b90613f2d565b80f35b631e670e4b60e01b8452600452602483fd5b503461028457806003193601126102845760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461028457604036600319011261028457612b43613660565b90602435906001600160401b0382116102845760206129a684612b69366004870161371f565b90613957565b503461028457604036600319011261028457600435906001600160401b03821161028457816004016101006003198436030112610a6857612bae61368a565b9082604051612bbc81613506565b5260648401359160848501612bd081613943565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612f6b5750602485019167ffffffffffffffff60801b612c1d8461392f565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156125ba578691612f4c575b50612f3d576001600160401b03612c8a8461392f565b16612ca2816000526006602052604060002054151590565b15612f2b5760206001600160a01b0360025416916044604051809481936383826b2b60e01b835260048301523360248301525afa9081156125ba578691612f0c575b5015612ef957612cf38361392f565b90612d1060a4880192612b69612d098585613ed6565b36916136e8565b15612ecb5750506001600160401b03602095612dfa6044612df37ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09661ffff608097161515600014612e5757897f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a80612dcc60408b612d8f8861392f565b16808752600460205295207f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169283916147ee565b604080516001600160a01b039092168252602082019290925290819081015b0390a261392f565b9201613943565b946001600160a01b0360405196817f0000000000000000000000000000000000000000000000000000000000000000168852338a89015216604087015260608601521692a260405190612e4c82613506565b815260405190518152f35b897f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8a80612dcc600260408c612e8c8961392f565b169687815260206007905220016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916147ee565b612ed59250613ed6565b611f966040519283926324eb47e560e01b84526020600485015260248401916139be565b63728fe07b60e01b855233600452602485fd5b612f25915060203d602011611e9157611e838183613588565b38612ce4565b6354c8163f60e11b8652600452602485fd5b630a75a23b60e31b8552600485fd5b612f65915060203d602011611e9157611e838183613588565b38612c74565b846001600160a01b03611eac602493613943565b503461028457602036600319011261028457600435906001600160401b03821161028457816004016101006003198436030112610a685781604051612fc381613506565b5281604051612fd181613506565b5260648301359060848401612fe581613943565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361329b5750602484019067ffffffffffffffff60801b6130328361392f565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561325057859161327c575b5061326d576001600160401b0361309f8361392f565b166130b7816000526006602052604060002054151590565b1561325b5760206001600160a01b0360025416916044604051809481936383826b2b60e01b835260048301523360248301525afa908115613250578591613231575b501561321e576131088261392f565b9061311e60a4870192612b69612d098585613ed6565b15612ecb57505060806001600160401b03602095846001600160a01b036131f860446131f17ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0988b6131af600260408b6131778661392f565b16938481526020600790522001877f0000000000000000000000000000000000000000000000000000000000000000169d8e916147ee565b604080516001600160a01b038e168152602081018990527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9181908101612deb565b9401613943565b604051978852338a89015216604087015260608601521692a260405190612e4c82613506565b63728fe07b60e01b845233600452602484fd5b61324a915060203d602011611e9157611e838183613588565b386130f9565b6040513d87823e3d90fd5b6354c8163f60e11b8552600452602484fd5b630a75a23b60e31b8452600484fd5b613295915060203d602011611e9157611e838183613588565b38613089565b836001600160a01b03611eac602493613943565b50346102845760c0366003190112610284576132c9613605565b506132d261364a565b6132da613620565b506084359161ffff831683036102845760a435906001600160401b0382116102845760a063ffffffff8061ffff61332088886133193660048b016136bb565b505061380a565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b5034610284578060031936011261028457602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610284576020366003190112610284576020906133a2613605565b90506001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461028457806003193601126102845760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461028457806003193601126102845750611752604051613443604082613588565b601781527f43435450546f6b656e506f6f6c20312e372e302d64657600000000000000000060208201526040519182916020835260208301906135c4565b905034610a68576020366003190112610a685760043563ffffffff60e01b8116809103610f28576020925063aff2afbf60e01b81149081156134f5575b81156134e4575b81156134d3575b5015158152f35b6301ffc9a760e01b149050386134cc565b630e64dd2960e01b811491506134c5565b633317103160e01b811491506134be565b602081019081106001600160401b0382111761352157604052565b634e487b7160e01b600052604160045260246000fd5b604081019081106001600160401b0382111761352157604052565b60e081019081106001600160401b0382111761352157604052565b60a081019081106001600160401b0382111761352157604052565b90601f801991011681019081106001600160401b0382111761352157604052565b6001600160401b03811161352157601f01601f191660200190565b919082519283825260005b8481106135f0575050826000602080949584010152601f8019910116010190565b806020809284010151828286010152016135cf565b600435906001600160a01b038216820361361b57565b600080fd5b606435906001600160a01b038216820361361b57565b35906001600160a01b038216820361361b57565b602435906001600160401b038216820361361b57565b600435906001600160401b038216820361361b57565b35906001600160401b038216820361361b57565b6024359061ffff8216820361361b57565b6044359061ffff8216820361361b57565b359061ffff8216820361361b57565b9181601f8401121561361b578235916001600160401b03831161361b576020838186019501011161361b57565b9291926136f4826135a9565b916137026040519384613588565b82948184528183011161361b578281602093846000960137010152565b9080601f8301121561361b5781602061373a933591016136e8565b90565b90604060031983011261361b576004356001600160401b038116810361361b5791602435906001600160401b03821161361b5761377c916004016136bb565b9091565b9181601f8401121561361b578235916001600160401b03831161361b576020808501948460051b01011161361b57565b9181601f8401121561361b578235916001600160401b03831161361b576020808501946060850201011161361b57565b61373a9160206137f983516040845260408401906135c4565b9201519060208184039101526135c4565b6001600160401b0316600052600a60205260406000206040519061382d82613552565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c168152608085019561ffff8360801c16875260ff60a087019361ffff8160901c16855260a01c1615801560c08801526139165761ffff16806138c65750505063ffffffff808061ffff9351169451169551169351169193929190600190565b919550915061ffff60025460a01c16908181106138ff57505063ffffffff808061ffff9351169451169551169351169193929190600190565b637911d95b60e01b60005260045260245260446000fd5b5050505092505050600090600090600090600090600090565b356001600160401b038116810361361b5790565b356001600160a01b038116810361361b5790565b906001600160401b0361373a92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b6001600160401b0381116135215760051b60200190565b51906001600160a01b038216820361361b57565b908060209392818452848401376000828201840152601f01601f1916010190565b9593919294906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958615613b4657613a64946001600160401b0361ffff936001600160a01b036040519b6344b9053160e11b8d521660048c01521660248a0152604489015216606487015260c0608487015260c48601916139be565b916002821015613b30578380600094819460a483015203915afa908115613b2457600091613a90575090565b3d8083833e613a9f8183613588565b810190602081830312610f28578051906001600160401b038211610280570181601f82011215610f2857805190613ad582613993565b93613ae36040519586613588565b82855260208086019360051b8301019384116102845750602001905b828210613b0c5750505090565b60208091613b19846139aa565b815201910190613aff565b6040513d6000823e3d90fd5b634e487b7160e01b600052602160045260246000fd5b5050505050505050604051613b5c602082613588565b60008152600036813790565b9190811015613b785760051b0190565b634e487b7160e01b600052603260045260246000fd5b9190811015613b78576060020190565b3590811515820361361b57565b35906001600160801b038216820361361b57565b919082606091031261361b57604051606081018181106001600160401b03821117613521576040526040613c11818395613bf881613b9e565b8552613c0660208201613bab565b602086015201613bab565b910152565b60405190613c238261356d565b60006080838281528260208201528260408201528260608201520152565b90604051613c4e8161356d565b60806001829460ff81546001600160801b038116865263ffffffff81861c16602087015260a01c161515604085015201546001600160801b0381166060840152811c910152565b60405190613ca282613537565b60606020838281520152565b9081602091031261361b57516001600160e01b03198116810361361b5790565b8051821015613b785760209160051b010190565b90600182811c92168015613d12575b6020831014613cfc57565b634e487b7160e01b600052602260045260246000fd5b91607f1691613cf1565b9060405191826000825492613d3084613ce2565b8084529360018116908115613d9e5750600114613d57575b50613d5592500383613588565b565b90506000929192526020600020906000915b818310613d82575050906020613d559282010138613d48565b6020919350806001915483858901015201910190918492613d69565b905060209250613d5594915060ff191682840152151560051b82010138613d48565b6001600160401b0316600052600760205261373a6004604060002001613d1c565b35801515810361361b5790565b356001600160801b038116810361361b5790565b6001600160801b03613e3760408093613e1a81613b9e565b1515865283613e2b60208301613bab565b16602087015201613bab565b16910152565b9190811015613b785760081b0190565b3561ffff8116810361361b5790565b3563ffffffff8116810361361b5790565b359063ffffffff8216820361361b57565b818110613e89575050565b60008155600101613e7e565b81810292918115918404141715613ea857565b634e487b7160e01b600052601160045260246000fd5b9081602091031261361b5751801515810361361b5790565b903590601e198136030182121561361b57018035906001600160401b03821161361b5760200191813603831361361b57565b6001600160a01b03600154163303613f1c57565b6315ae3a6f60e11b60005260046000fd5b908051156140f9576001600160401b0381516020830120921691826000526007602052613f61816005604060002001614799565b156140ce576000526008602052604060002081516001600160401b03811161352157613f8d8254613ce2565b601f811161409c575b506020601f82116001146140125791613fec827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361400295600091614007575b508160011b916000199060031b1c19161790565b90556040519182916020835260208301906135c4565b0390a2565b905084015138613fd8565b601f1982169083600052806000209160005b8181106140845750926140029492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061406b575b5050811b01905561173e565b85015160001960f88460031b161c19169055388061405f565b9192602060018192868a015181550194019201614024565b6140c890836000526020600020601f840160051c8101916020851061086057601f0160051c0190613e7e565b38613f96565b5090611f96604051928392631c9dc56960e11b845260048401526040602484015260448301906135c4565b630a64406560e11b60005260046000fd5b6001600160a01b03600954163314158061413a575b61412557565b63472511eb60e11b6000523360045260246000fd5b506001600160a01b036001541633141561411f565b91908203918211613ea857565b614164613c16565b506001600160801b036060820151166001600160801b0380835116916141af60208501936141a961419c63ffffffff8751164261414f565b8560808901511690613e95565b90614524565b808210156141c857505b16825263ffffffff4216905290565b90506141b9565b604051630c11d61f60e21b60208201526001600160e01b031990911660248201526008815261373a602882613588565b9061ffff906001600160401b036142186020850161392f565b16600052600a6020526040600020826040519161423483613552565b549263ffffffff8416835263ffffffff8460201c16602084015263ffffffff8460401c16604084015263ffffffff8460601c166060840152818460801c169283608082015260c060ff848760901c16968760a085015260a01c1615159101521615156000146142ca57505b1680156142c2576127106142bb606061373a9401359283613e95565b049061414f565b506060013590565b905061429f565b805115614338576001600160801b036040820151166001600160801b03602083015116106142fc5750565b60408051632008344960e21b815282511515600482015260208301516001600160801b0390811660248301529190920151166044820152606490fd5b6001600160801b0360408201511615801590614392575b6143565750565b604080516335a2be7360e21b815282511515600482015260208301516001600160801b0390811660248301529190920151166044820152606490fd5b506001600160801b03602082015116151561434f565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c199161447660609280546143e563ffffffff8260801c164261414f565b90816144af575b50506001600160801b0360018160208601511692828154168085106000146144a757508280855b1616831982541617815561444286511515829081549060ff60a01b90151560a01b169060ff60a01b1916179055565b60408601516fffffffffffffffffffffffffffffffff1960809190911b16939092166001600160801b031692909217910155565b6144a460405180926001600160801b0360408092805115158552826020820151166020860152015116910152565ba1565b838091614413565b6001600160801b03916144db8392836144d46001880154948286169560801c90613e95565b9116614524565b8082101561451d57505b835463ffffffff60801b199290911692909216166001600160a01b0319909116174260801b63ffffffff60801b1617815538806143ec565b90506144e5565b91908201809211613ea857565b906040519182815491828252602082019060005260206000209260005b818110614563575050613d5592500383613588565b845483526001948501948794506020909301920161454e565b8054821015613b785760005260206000200190600090565b6000818152600660205260409020548015614675576000198101818111613ea857600554600019810191908211613ea857818103614624575b505050600554801561460e57600019016145e881600561457c565b8154906000199060031b1b19169055600555600052600660205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61465d61463561464693600561457c565b90549060031b1c928392600561457c565b819391549060031b91821b91600019901b19161790565b905560005260066020526040600020553880806145cd565b5050600090565b9060018201918160005282602052604060002054801515600014614730576000198101818111613ea8578254600019810191908211613ea8578181036146f9575b5050508054801561460e5760001901906146d7828261457c565b8154906000199060031b1b191690555560005260205260006040812055600190565b614719614709614646938661457c565b90549060031b1c9283928661457c565b9055600052836020526040600020553880806146bd565b50505050600090565b8060005260066020526040600020541560001461479357600554680100000000000000008110156135215761477a614646826001859401600555600561457c565b9055600554906000526006602052604060002055600190565b50600090565b6000828152600182016020526040902054614675578054906801000000000000000082101561352157826147d761464684600180960185558461457c565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015614966575b614960576001600160801b038216916001850190815461483463ffffffff6001600160801b0383169360801c164261414f565b9081614900575b50508481106148da57508383106148745750506148616001600160801b0392839261414f565b16166001600160801b0319825416179055565b5460801c91614883818561414f565b6000198401848111613ea85761489891614524565b9280156148c4576001600160a01b039304636864691d60e11b6000526004526024521660445260646000fd5b634e487b7160e01b600052601260045260246000fd5b82856001600160a01b0392630d3b2b9560e11b6000526004526024521660445260646000fd5b82869293961161494f5761491b926141a99160801c90613e95565b8084101561494a5750825b855463ffffffff60801b19164260801b63ffffffff60801b1617865592388061483b565b614926565b634b92ca1560e11b60005260046000fd5b50505050565b508215614801565b919290156149d05750815115614982575090565b3b1561498b5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156149e35750805190602001fd5b60405162461bcd60e51b815260206004820152908190611f969060248301906135c456fea164736f6c634300081a000a",
}

var CCTPTokenPoolABI = CCTPTokenPoolMetaData.ABI

var CCTPTokenPoolBin = CCTPTokenPoolMetaData.Bin

func DeployCCTPTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address, cctpVerifier common.Address) (common.Address, *types.Transaction, *CCTPTokenPool, error) {
	parsed, err := CCTPTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPTokenPoolBin), backend, token, localTokenDecimals, advancedPoolHooks, rmnProxy, router, cctpVerifier)
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

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getCurrentCustomBlockConfirmationRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentCustomBlockConfirmationRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _CCTPTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _CCTPTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _CCTPTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPTokenPool.CallOpts, remoteChainSelector)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _CCTPTokenPool.Contract.GetCurrentRateLimiterState(&_CCTPTokenPool.CallOpts, remoteChainSelector)
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
	outstruct.MinBlockConfirmations = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.LockOrBurn(&_CCTPTokenPool.TransactOpts, lockOrBurnIn)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.LockOrBurn(&_CCTPTokenPool.TransactOpts, lockOrBurnIn)
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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ReleaseOrMint(&_CCTPTokenPool.TransactOpts, releaseOrMintIn)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ReleaseOrMint(&_CCTPTokenPool.TransactOpts, releaseOrMintIn)
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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetChainRateLimiterConfigs(&_CCTPTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetChainRateLimiterConfigs(&_CCTPTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setCustomBlockConfirmationRateLimitConfig", rateLimitConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_CCTPTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_CCTPTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setDynamicConfig", router, minBlockConfirmations, rateLimitAdmin)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetDynamicConfig(&_CCTPTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetDynamicConfig(&_CCTPTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.WithdrawFeeTokens(&_CCTPTokenPool.TransactOpts, feeTokens, recipient)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.WithdrawFeeTokens(&_CCTPTokenPool.TransactOpts, feeTokens, recipient)
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

type CCTPTokenPoolConfigChangedIterator struct {
	Event *CCTPTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolConfigChanged)
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
		it.Event = new(CCTPTokenPoolConfigChanged)
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

func (it *CCTPTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*CCTPTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolConfigChangedIterator{contract: _CCTPTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolConfigChanged)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseConfigChanged(log types.Log) (*CCTPTokenPoolConfigChanged, error) {
	event := new(CCTPTokenPoolConfigChanged)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

type CCTPTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator struct {
	Event *CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured)
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
		it.Event = new(CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured)
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

func (it *CCTPTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterCustomBlockConfirmationRateLimitConfigured(opts *bind.FilterOpts) (*CCTPTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator{contract: _CCTPTokenPool.contract, event: "CustomBlockConfirmationRateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchCustomBlockConfirmationRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationRateLimitConfigured", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseCustomBlockConfirmationRateLimitConfigured(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured, error) {
	event := new(CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationRateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolDefaultFinalityRateLimitConfiguredIterator struct {
	Event *CCTPTokenPoolDefaultFinalityRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolDefaultFinalityRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolDefaultFinalityRateLimitConfigured)
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
		it.Event = new(CCTPTokenPoolDefaultFinalityRateLimitConfigured)
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

func (it *CCTPTokenPoolDefaultFinalityRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolDefaultFinalityRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolDefaultFinalityRateLimitConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterDefaultFinalityRateLimitConfigured(opts *bind.FilterOpts) (*CCTPTokenPoolDefaultFinalityRateLimitConfiguredIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "DefaultFinalityRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolDefaultFinalityRateLimitConfiguredIterator{contract: _CCTPTokenPool.contract, event: "DefaultFinalityRateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchDefaultFinalityRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolDefaultFinalityRateLimitConfigured) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "DefaultFinalityRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolDefaultFinalityRateLimitConfigured)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "DefaultFinalityRateLimitConfigured", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseDefaultFinalityRateLimitConfigured(log types.Log) (*CCTPTokenPoolDefaultFinalityRateLimitConfigured, error) {
	event := new(CCTPTokenPoolDefaultFinalityRateLimitConfigured)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "DefaultFinalityRateLimitConfigured", log); err != nil {
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
	Router                common.Address
	MinBlockConfirmations uint16
	RateLimitAdmin        common.Address
	Raw                   types.Log
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
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*CCTPTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolFeeTokenWithdrawnIterator{contract: _CCTPTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
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

type GetCurrentCustomBlockConfirmationRateLimiterState struct {
	OutboundRateLimiterState RateLimiterTokenBucket
	InboundRateLimiterState  RateLimiterTokenBucket
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

func (CCTPTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (CCTPTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (CCTPTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x20ae59542ddd78610f62f9d2c9dcd658f8b6a5b45a0f03e71e16614e89dda836")
}

func (CCTPTokenPoolDefaultFinalityRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x73d6dce40db73cbddae4b9ce52576043a1644e08c2702236273d71077435fa16")
}

func (CCTPTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a117")
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

func (CCTPTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (CCTPTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (CCTPTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
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
	GetCCTPVerifier(opts *bind.CallOpts) (common.Address, error)

	GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

		error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

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

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterChainAdded(opts *bind.FilterOpts) (*CCTPTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*CCTPTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*CCTPTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*CCTPTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*CCTPTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*CCTPTokenPoolConfigChanged, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationRateLimitConfigured(opts *bind.FilterOpts) (*CCTPTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator, error)

	WatchCustomBlockConfirmationRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured) (event.Subscription, error)

	ParseCustomBlockConfirmationRateLimitConfigured(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationRateLimitConfigured, error)

	FilterDefaultFinalityRateLimitConfigured(opts *bind.FilterOpts) (*CCTPTokenPoolDefaultFinalityRateLimitConfiguredIterator, error)

	WatchDefaultFinalityRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolDefaultFinalityRateLimitConfigured) (event.Subscription, error)

	ParseDefaultFinalityRateLimitConfigured(log types.Log) (*CCTPTokenPoolDefaultFinalityRateLimitConfigured, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*CCTPTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*CCTPTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*CCTPTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*CCTPTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*CCTPTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*CCTPTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*CCTPTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CCTPTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*CCTPTokenPoolOwnershipTransferred, error)

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
