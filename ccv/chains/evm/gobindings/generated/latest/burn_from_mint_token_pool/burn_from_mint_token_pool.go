// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_from_mint_token_pool

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

var BurnFromMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103485760a081615dda80380380916100208285610471565b8339810103126103485780516001600160a01b03811691908290036103485761004b60208201610494565b610057604083016104a2565b906100706080610069606086016104a2565b94016104a2565b92331561046057600180546001600160a01b031916331790558415801561044f575b801561043e575b61042d57608085905260c05260405163313ce56760e01b8152602081600481885afa600091816103f1575b506103c6575b5060a0526001600160a01b0390811660e052600280546001600160a01b03191692909116919091179055604051636eb1769f60e11b81523060048201819052602482015290602082604481845afa9182156103ba57600092610386575b50600019820180921161037057604051602081019263095ea7b360e01b8452306024830152604482015260448152610160606482610471565b6000806040948551936101738786610471565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610363573d906001600160401b03821161034d5784516101e49490926101d5601f8201601f191660200185610471565b83523d6000602085013e6104b6565b8051806102cd575b8251615853908161058782396080518181816117a6015281816119b901528181611bfb01528181611cde015281816121600152818161231701528181612a6401528181612c5d01528181612d180152818161307f015281816132a301528181613475015281816139b301528181613a0d0152614d2e015260a051818181613879015281816148bd015281816149400152614dbc015260c051818181610c9d01528181611841015281816121fa01528181612aff015261333e015260e0518181816119650152818161239e01528181612cc5015281816134fc01526142fa0152f35b81602091810103126103485760200151801590811503610348576102f25738806101ec565b5162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b916101e4926060916104b6565b634e487b7160e01b600052601160045260246000fd5b9091506020813d6020116103b2575b816103a260209383610471565b8101031261034857519038610127565b3d9150610395565b6040513d6000823e3d90fd5b60ff1660ff82168181036103da57506100ca565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610425575b8161040d60209383610471565b810103126103485761041e90610494565b90386100c4565b3d9150610400565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610099565b506001600160a01b03841615610092565b639b15e16f60e01b60005260046000fd5b601f909101601f19168101906001600160401b0382119082101761034d57604052565b519060ff8216820361034857565b51906001600160a01b038216820361034857565b9192901561051857508151156104ca575090565b3b156104d35790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561052b5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b83811061056e5750508160006044809484010152601f80199101168101030190fd5b6020828201810151604487840101528593500161054c56fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613a9257508063181f5a7714613a3157806321df0da7146139e0578063240028e81461397c5780632422ac451461389d57806324f65ee71461385f5780632c063404146137c65780633907753714613204578063489a68f2146129bf5780634c5ef0ed1461297857806362ddd3c4146128f15780637437ff9f146128a057806379ba5097146127d95780638926f54f1461279357806389720a62146126cc5780638da5cb5b146126985780639a4575b9146120e6578063a42a7b8b14611f7f578063acfecf9114611e87578063b1c71c6514611700578063b7946580146116c3578063c4bffe2b14611598578063c7230a601461124c578063d8aa3f4014611112578063dc04fa1f14610cc1578063dc0bd97114610c70578063dcbd41bc14610a6c578063e8a1da17146103ac578063f2fde38b146102dd5763fdf168751461016b57600080fd5b346102da5760606003193601126102da57610184613d0d565b9061018d613dc9565b6044359273ffffffffffffffffffffffffffffffffffffffff84168085036102d6576101b7614a62565b73ffffffffffffffffffffffffffffffffffffffff821680156102ae57946102a8917fba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a1179596600254907fffffffffffffffffffff0000000000000000000000000000000000000000000075ffff00000000000000000000000000000000000000008860a01b16921617176002557fffffffffffffffffffffffff000000000000000000000000000000000000000060095416176009556040519384938491604091949361ffff73ffffffffffffffffffffffffffffffffffffffff9283606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b80fd5b50346102da5760206003193601126102da5773ffffffffffffffffffffffffffffffffffffffff61030c613d0d565b610314614a62565b1633811461038457807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346102da5760406003193601126102da5760043567ffffffffffffffff81116108c9576103de903690600401613eec565b9060243567ffffffffffffffff81116102d6579061040184923690600401613eec565b93909161040c614a62565b83905b8282106108d15750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156108cd578060051b830135858112156108c5578301610120813603126108c5576040519461047386613c17565b61047c82613da7565b8652602082013567ffffffffffffffff81116108c95782019436601f870112156108c9578535956104ac87614265565b966104ba6040519889613c33565b80885260208089019160051b830101903682116108c55760208301905b828210610892575050505060208701958652604083013567ffffffffffffffff811161088e5761050a9036908501613e5f565b916040880192835261053461052236606087016146fe565b9460608a0195865260c03691016146fe565b9560808901968752835151156108665761055867ffffffffffffffff8a5116615409565b1561082f5767ffffffffffffffff895116825260076020526040822061057f865182614e56565b61058d885160028301614e56565b6004855191019080519067ffffffffffffffff8211610802576105b08354614510565b601f81116107c7575b50602090601f831160011461072857610607929186918361071d575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b88518051821015610641579061063b6001926106348367ffffffffffffffff8f5116926144cd565b5190614aad565b0161060c565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561070f67ffffffffffffffff60019796949851169251935191516106db6106a660405196879687526101006020880152610100870190613cae565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610442565b015190508e806105d5565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b8181106107af5750908460019594939210610778575b505050811b01905561060a565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d808061076b565b92936020600181928786015181550195019301610755565b6107f29084875260208720601f850160051c810191602086106107f8575b601f0160051c019061479a565b8d6105b9565b90915081906107e5565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116108c1576020916108b68392833691890101613e5f565b8152019101906104d7565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff6108f36108ee8486889a9699979a614683565b614213565b16916108fe83615146565b15610a4057828452600760205261091a600560408620016150e3565b94845b865181101561095357600190858752600760205261094c60056040892001610945838b6144cd565b51906152d5565b500161091d565b509396929094509490948087526007602052600560408820888155886001820155886002820155886003820155886004820161098f8154614510565b806109ff575b50505001805490888155816109e1575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190919493929461040f565b885260208820908101905b818110156109a5578881556001016109ec565b601f8111600114610a155750555b888a80610995565b81835260208320610a3091601f01861c81019060010161479a565b8082528160208120915555610a0d565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346102da5760206003193601126102da5760043567ffffffffffffffff81116108c957610a9e903690600401613f1d565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580610c4e575b610c2257825b818110610ad1578380f35b610adc818385614693565b67ffffffffffffffff610aee82614213565b1690610b07826000526006602052604060002054151590565b15610bf657907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610bb6610b90602060019897018b610b48826146a3565b15610bbd578790526003602052610b6f60408d20610b6936604088016146fe565b90614e56565b868c526004602052610b8b60408d20610b693660a088016146fe565b6146a3565b916040519215158352610ba96020840160408301614756565b60a0608084019101614756565ba201610ac6565b60026040828a610b8b94526007602052610bdf828220610b6936858c016146fe565b8a815260076020522001610b693660a088016146fe565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610ac0565b50346102da57806003193601126102da57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102da5760406003193601126102da5760043567ffffffffffffffff81116108c957610cf3903690600401613f1d565b60243567ffffffffffffffff81116102d657610d13903690600401613eec565b919092610d1e614a62565b845b828110610d8a57505050825b818110610d37578380f35b8067ffffffffffffffff610d516108ee6001948688614683565b16808652600a6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610d2c565b610d986108ee828585614693565b610da3828585614693565b90602082019060e0830190610db7826146a3565b156110dd5760a0840161271061ffff610dcf836146b0565b1610156110ce5760c085019161271061ffff610dea856146b0565b1610156110965763ffffffff610dff866146bf565b16156110615767ffffffffffffffff1694858c52600a60205260408c20610e25866146bf565b63ffffffff16908054906040840191610e3d836146bf565b60201b67ffffffff0000000016936060860194610e59866146bf565b60401b6bffffffff0000000000000000169660800196610e78886146bf565b60601b6fffffffff0000000000000000000000001691610e978a6146b0565b60801b71ffff000000000000000000000000000000001693610eb88c6146b0565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610f6b876146a3565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610fbc906146d0565b63ffffffff168752610fcd906146d0565b63ffffffff166020870152610fe1906146d0565b63ffffffff166040860152610ff5906146d0565b63ffffffff16606085015261100990613deb565b61ffff16608084015261101b90613deb565b61ffff1660a083015261102d90613dbc565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610d20565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff6110a5866146b0565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff6110a56024936146b0565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346102da5760806003193601126102da5761112c613d0d565b50611135613d90565b61113d613dda565b5060643567ffffffffffffffff811161088e579167ffffffffffffffff60409261116d60e0953690600401613dfa565b50508260c0855161117d81613bfb565b82815282602082015282878201528260608201528260808201528260a08201520152168152600a60205220604051906111b582613bfb565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346102da5760406003193601126102da5760043567ffffffffffffffff81116108c95761127e903690600401613eec565b6024359073ffffffffffffffffffffffffffffffffffffffff8216908183036108c5576112a9614a62565b845b8181106112b6578580f35b80602073ffffffffffffffffffffffffffffffffffffffff6112e36112de602495878b614683565b6141f2565b16604051938480927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa801561158d5784869185948a91611551575b508061133c575b50505060019150016112ab565b8861143473ffffffffffffffffffffffffffffffffffffffff6113636112de888a86614683565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff98909816602482015260448082018790528152918e91166113c5606484613c33565b81806040998a51956113d78c88613c33565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15611548573d61141881613c74565b906114258b519283613c33565b8152809360203d92013e61577a565b805180611499575b50507f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9173ffffffffffffffffffffffffffffffffffffffff6114876112de8860019a602096614683565b169451908152a382915084843861132f565b8160209350839294969798506114b3955001019101614a4a565b156114c557908484939288388061143c565b608482517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060925061577a565b93945050505060203d8111611586575b61156b8183613c33565b602082600092810103126102da575090848484935138611328565b503d611561565b6040513d89823e3d90fd5b50346102da57806003193601126102da57604051906005548083528260208101600584526020842092845b8181106116aa5750506115d892500383613c33565b81516115fc6115e682614265565b916115f46040519384613c33565b808352614265565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b845181101561165b578067ffffffffffffffff611648600193886144cd565b511661165482866144cd565b5201611629565b50925090604051928392602084019060208552518091526040840192915b818110611687575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611679565b84548352600194850194879450602090930192016115c3565b50346102da5760206003193601126102da576116fc6116e86116e3613d79565b614661565b604051918291602083526020830190613cae565b0390f35b50346102da5760606003193601126102da5760043567ffffffffffffffff81116108c9578060040160a0600319833603011261088e5761173e613dc9565b926044359367ffffffffffffffff85116108c957611763611780953690600401613dfa565b959061176d6144b4565b506117788386614df0565b963691613e28565b608485019061178e826141f2565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611e3d57602486019577ffffffffffffffff000000000000000000000000000000006117f488614213565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611db0578691611e0e575b50611de65767ffffffffffffffff61188888614213565b166118a0816000526006602052604060002054151590565b15611dbb57602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611db0578690611d5f575b73ffffffffffffffffffffffffffffffffffffffff9150163303611d335760648101359561ffff6119328a89614625565b9516948515611c825761ffff60025460a01c1680611b97575b50505b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283611a75575b611a6b89611a3a6116e36119a08e8d614625565b926119aa84614d17565b6119b381614213565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2614213565b90611a43614db5565b60405192611a5084613bdf565b83526020830152604051928392604084526040840190613ec2565b9060208301520390f35b833b156108c157869493929185918a604051988997889687957f5c3af7ca000000000000000000000000000000000000000000000000000000008752600487016060905280611ac391615093565b6064880160a09052610104880190611ada9261429e565b93611ae490613da7565b67ffffffffffffffff166084870152604401611aff90613d58565b73ffffffffffffffffffffffffffffffffffffffff1660a48601528c60c4860152611b2990613d58565b73ffffffffffffffffffffffffffffffffffffffff1660e48501526024840152828103600319016044840152611b5e91613cae565b03925af18015611b8c57611b77575b808080808061198c565b611b82828092613c33565b6102da5780611b6d565b6040513d84823e3d90fd5b808710611c5257507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611bd28b614213565b1691828952600360205280611c2360408b2073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916154be565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2388061194b565b87604491887f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611cb58b614213565b1691828952600760205280611d0660408b2073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916154be565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a261194e565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611da8575b81611d7960209383613c33565b81010312611da457611d9f73ffffffffffffffffffffffffffffffffffffffff9161427d565b611901565b8580fd5b3d9150611d6c565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611e30915060203d602011611e36575b611e288183613c33565b810190614a4a565b38611871565b503d611e1e565b60248473ffffffffffffffffffffffffffffffffffffffff611e5e856141f2565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346102da5767ffffffffffffffff611e9f36613e7d565b929091611eaa614a62565b1691611ec3836000526006602052604060002054151590565b15610a40578284526007602052611ef260056040862001611ee5368486613e28565b60208151910120906152d5565b15611f3757907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611f3160405192839260208452602084019161429e565b0390a280f35b82611f7b836040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161429e565b0390fd5b50346102da5760206003193601126102da5767ffffffffffffffff611fa2613d79565b1681526007602052611fb9600560408320016150e3565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611ffe611fe883614265565b92611ff66040519485613c33565b808452614265565b01835b8181106120d5575050825b82518110156120525780612022600192856144cd565b518552600860205261203660408620614563565b61204082856144cd565b5261204b81846144cd565b500161200c565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061208a57505050500390f35b919360206120c5827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613cae565b960192019201859493919261207b565b806060602080938601015201612001565b50346102da5760206003193601126102da576004359067ffffffffffffffff82116102da57816004019160a060031982360301126108c9576121266144b4565b50602092604051926121388585613c33565b80845260848301612148816141f2565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361267757602484019477ffffffffffffffff000000000000000000000000000000006121ae87614213565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152878160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156125fc57849161265a575b506126325767ffffffffffffffff61224187614213565b16612259816000526006602052604060002054151590565b15612607578773ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156125fc5784906125b4575b73ffffffffffffffffffffffffffffffffffffffff915016330361258857606485013594859467ffffffffffffffff6122f189614213565b1680865260078a5261233f6040872073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016998a916154be565b6040805173ffffffffffffffffffffffffffffffffffffffff8a168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283612471575b896124416116e38b8b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108c6123f981614d17565b67ffffffffffffffff61240b85614213565b6040805173ffffffffffffffffffffffffffffffffffffffff909616865233602087015285019290925216918060608101611a32565b9061244a614db5565b6040519261245784613bdf565b8352818301526116fc604051928284938452830190613ec2565b833b15611da45791858094928a9694604051978896879586947f5c3af7ca0000000000000000000000000000000000000000000000000000000086526004860160609052806124bf91615093565b6064870160a090526101048701906124d69261429e565b926124e090613da7565b67ffffffffffffffff1660848601526044016124fb90613d58565b73ffffffffffffffffffffffffffffffffffffffff1660a48501528b60c485015261252590613d58565b73ffffffffffffffffffffffffffffffffffffffff1660e484015283602484015282810360031901604484015261255b91613cae565b03925af18015611b8c57612573575b808080806123c5565b61257e828092613c33565b6102da578061256a565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d83116125f5575b6125ca8183613c33565b810103126102d6576125f073ffffffffffffffffffffffffffffffffffffffff9161427d565b6122b9565b503d6125c0565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6126719150883d8a11611e3657611e288183613c33565b3861222a565b9073ffffffffffffffffffffffffffffffffffffffff611e5e6024936141f2565b50346102da57806003193601126102da57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346102da5760c06003193601126102da576126e6613d0d565b6126ee613d90565b9060643561ffff811681036102d65760843567ffffffffffffffff81116108c55761271d903690600401613dfa565b9160a4359360028510156108c15761273895604435916142dd565b90604051918291602083016020845282518091526020604085019301915b818110612764575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101612756565b50346102da5760206003193601126102da5760206127cf67ffffffffffffffff6127bb613d79565b166000526006602052604060002054151590565b6040519015158152f35b50346102da57806003193601126102da57805473ffffffffffffffffffffffffffffffffffffffff81163303612878577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b50346102da57806003193601126102da576002546009546040805173ffffffffffffffffffffffffffffffffffffffff808516825260a09490941c61ffff1660208201529290911690820152606090f35b50346102da5761290036613e7d565b61290c93929193614a62565b67ffffffffffffffff821661292e816000526006602052604060002054151590565b1561294d575061294a9293612944913691613e28565b90614aad565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346102da5760406003193601126102da57612992613d79565b906024359067ffffffffffffffff82116102da5760206127cf846129b93660048701613e5f565b90614228565b50346102da5760406003193601126102da5760043567ffffffffffffffff81116108c957806004019161010060031983360301126102da576129ff613dc9565b9181604051612a0d81613b94565b5260c48101926064820135612a3d612a37612a32612a2b888a6141a1565b3691613e28565b614849565b8261493d565b946084840191612a4c836141f2565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036131e357602485019777ffffffffffffffff00000000000000000000000000000000612ab28a614213565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156131665788916131c4575b5061319c5767ffffffffffffffff612b468a614213565b16612b5e816000526006602052604060002054151590565b1561317157602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115613166578891613147575b501561311b57612bd589614213565b94612beb60a48801966129b9612a2b89866141a1565b156130d45761ffff1690878a8a8415613021575067ffffffffffffffff9150612c1390614213565b1680895260046020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a80612c8560408d2073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916154be565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169384612e42575b5050505050505060440192612d00846141f2565b9173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692833b156108c9576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101859052818180604481015b038183885af18015611b8c57612e2d575b5050608067ffffffffffffffff60209573ffffffffffffffffffffffffffffffffffffffff612dfb612df57ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096614213565b926141f2565b60405196875233898801521660408601528560608601521692a260405190612e2282613b94565b815260405190518152f35b612e38828092613c33565b6102da5780612da3565b843b1561301d578895949392869289928d6040519a8b998a9889977f5eff3bf70000000000000000000000000000000000000000000000000000000089526004890160609052612e928780615093565b60648b0161010090526101648b0190612eaa9261429e565b94612eb490613da7565b67ffffffffffffffff1660848a0152604401612ecf90613d58565b73ffffffffffffffffffffffffffffffffffffffff1660a489015260c4880152612ef890613d58565b73ffffffffffffffffffffffffffffffffffffffff1660e4870152612f1d9084615093565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c87840301610104880152612f52929161429e565b90612f5d9083615093565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c86840301610124870152612f92929161429e565b9060e48b01612fa091615093565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c85840301610144860152612fd5929161429e565b908b6024840152604483015203925af1801561301257908391612ffd575b8080808080612cec565b8161300791613c33565b6108c9578138612ff3565b6040513d85823e3d90fd5b8880fd5b806130a76002604067ffffffffffffffff61305c7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c97614213565b16968781526007602052200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283916154be565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612cae565b6130de86836141a1565b611f7b6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161429e565b6024877f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b613160915060203d602011611e3657611e288183613c33565b38612bc6565b6040513d8a823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008852600452602487fd5b6004877f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6131dd915060203d602011611e3657611e288183613c33565b38612b2f565b60248673ffffffffffffffffffffffffffffffffffffffff611e5e866141f2565b50346102da5760206003193601126102da576004359067ffffffffffffffff82116102da57816004019161010060031982360301126108c9578160405161324a81613b94565b528160405161325881613b94565b52606481013560c482019161327c613276612a32612a2b86896141a1565b8361493d565b92608482019261328b846141f2565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036137a557602483019377ffffffffffffffff000000000000000000000000000000006132f186614213565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115613166578891613786575b5061319c5767ffffffffffffffff61338586614213565b1661339d816000526006602052604060002054151590565b1561317157602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115613166578891613767575b501561311b5761341485614213565b9261342a60a48601946129b9612a2b878d6141a1565b1561375d579187989391889388995067ffffffffffffffff61344b89614213565b16808652600760205261349d6002604088200173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169b8c916154be565b6040805173ffffffffffffffffffffffffffffffffffffffff8c168152602081018d90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283613594575b5050505050505060440193613537856141f2565b833b156108c9576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024810185905281818060448101612d92565b833b15611da45788968692604051988997889687957f5eff3bf700000000000000000000000000000000000000000000000000000000875260048701606090528d6135df8780615093565b60648a0161010090526101648a01906135f79261429e565b9461360190613da7565b67ffffffffffffffff16608489015260440161361c90613d58565b73ffffffffffffffffffffffffffffffffffffffff1660a488015260c487015261364590613d58565b73ffffffffffffffffffffffffffffffffffffffff1660e486015261366a9084615093565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c8684030161010487015261369f929161429e565b906136aa9083615093565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101248601526136df929161429e565b9060e48a016136ed91615093565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610144850152613722929161429e565b8b602483015282604483015203925af18015611db057613748575b858180808080613523565b946137568160449397613c33565b949061373d565b6130de848a6141a1565b613780915060203d602011611e3657611e288183613c33565b38613405565b61379f915060203d602011611e3657611e288183613c33565b3861336e565b60248673ffffffffffffffffffffffffffffffffffffffff611e5e876141f2565b50346102da5760c06003193601126102da576137e0613d0d565b506137e9613d90565b6137f1613d35565b506084359161ffff831683036102da5760a4359067ffffffffffffffff82116102da5760a063ffffffff8061ffff61383888886138313660048b01613dfa565b5050614062565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b50346102da57806003193601126102da57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102da5760406003193601126102da576138b7613d79565b6024359182151583036102da5761014061397a6138d48585613fdf565b61392a60409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346102da5760206003193601126102da57602090613999613d0d565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b50346102da57806003193601126102da57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102da57806003193601126102da57506116fc604051613a54604082613c33565b601f81527f4275726e46726f6d4d696e74546f6b656e506f6f6c20312e372e302d646576006020820152604051918291602083526020830190613cae565b9050346108c95760206003193601126108c9576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361088e57602092507faff2afbf000000000000000000000000000000000000000000000000000000008114908115613b6a575b8115613b40575b8115613b16575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438613b0f565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613b08565b7f331710310000000000000000000000000000000000000000000000000000000081149150613b01565b6020810190811067ffffffffffffffff821117613bb057604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff821117613bb057604052565b60e0810190811067ffffffffffffffff821117613bb057604052565b60a0810190811067ffffffffffffffff821117613bb057604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117613bb057604052565b67ffffffffffffffff8111613bb057601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b919082519283825260005b848110613cf85750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201613cb9565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203613d3057565b600080fd5b6064359073ffffffffffffffffffffffffffffffffffffffff82168203613d3057565b359073ffffffffffffffffffffffffffffffffffffffff82168203613d3057565b6004359067ffffffffffffffff82168203613d3057565b6024359067ffffffffffffffff82168203613d3057565b359067ffffffffffffffff82168203613d3057565b35908115158203613d3057565b6024359061ffff82168203613d3057565b6044359061ffff82168203613d3057565b359061ffff82168203613d3057565b9181601f84011215613d305782359167ffffffffffffffff8311613d305760208381860195010111613d3057565b929192613e3482613c74565b91613e426040519384613c33565b829481845281830111613d30578281602093846000960137010152565b9080601f83011215613d3057816020613e7a93359101613e28565b90565b906040600319830112613d305760043567ffffffffffffffff81168103613d3057916024359067ffffffffffffffff8211613d3057613ebe91600401613dfa565b9091565b613e7a916020613edb8351604084526040840190613cae565b920151906020818403910152613cae565b9181601f84011215613d305782359167ffffffffffffffff8311613d30576020808501948460051b010111613d3057565b9181601f84011215613d305782359167ffffffffffffffff8311613d30576020808501948460081b010111613d3057565b60405190613f5b82613c17565b60006080838281528260208201528260408201528260608201520152565b90604051613f8681613c17565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff91613ff1613f4e565b50613ffa613f4e565b5061402e57166000526007602052604060002090613e7a614022600261402761402286613f79565b6147c4565b9401613f79565b16908160005260036020526140496140226040600020613f79565b916000526004602052613e7a6140226040600020613f79565b67ffffffffffffffff16600052600a60205260406000206040519061408682613bfb565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c168152608085019561ffff8360801c16875260ff60a087019361ffff8160901c16855260a01c1615801560c08801526141885761ffff168061411f5750505063ffffffff808061ffff9351169451169551169351169193929190600190565b919550915061ffff60025460a01c169081811061415857505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050505092505050600090600090600090600090600090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215613d30570180359067ffffffffffffffff8211613d3057602001918136038313613d3057565b3573ffffffffffffffffffffffffffffffffffffffff81168103613d305790565b3567ffffffffffffffff81168103613d305790565b9067ffffffffffffffff613e7a92166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613bb05760051b60200190565b519073ffffffffffffffffffffffffffffffffffffffff82168203613d3057565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016958615614492576143969467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c486019161429e565b916002821015614463578380600094819460a483015203915afa908115614457576000916143c2575090565b3d8083833e6143d18183613c33565b81019060208183031261088e5780519067ffffffffffffffff82116102d6570181601f8201121561088e5780519061440882614265565b936144166040519586613c33565b82855260208086019360051b8301019384116102da5750602001905b82821061443f5750505090565b6020809161444c8461427d565b815201910190614432565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b50505050505050506040516144a8602082613c33565b60008152600036813790565b604051906144c182613bdf565b60606020838281520152565b80518210156144e15760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600182811c92168015614559575b602083101461452a57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161451f565b906040519182600082549261457784614510565b80845293600181169081156145e5575060011461459e575b5061459c92500383613c33565b565b90506000929192526020600020906000915b8183106145c957505090602061459c928201013861458f565b60209193508060019154838589010152019101909184926145b0565b6020935061459c9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201013861458f565b9190820391821161463257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff166000526007602052613e7a6004604060002001614563565b91908110156144e15760051b0190565b91908110156144e15760081b0190565b358015158103613d305790565b3561ffff81168103613d305790565b3563ffffffff81168103613d305790565b359063ffffffff82168203613d3057565b35906fffffffffffffffffffffffffffffffff82168203613d3057565b9190826060910312613d30576040516060810181811067ffffffffffffffff821117613bb057604052604061475181839561473881613dbc565b8552614746602082016146e1565b6020860152016146e1565b910152565b6fffffffffffffffffffffffffffffffff6147946040809361477781613dbc565b1515865283614788602083016146e1565b166020870152016146e1565b16910152565b8181106147a5575050565b6000815560010161479a565b8181029291811591840414171561463257565b6147cc613f4e565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614829602085019361482361481663ffffffff87511642614625565b85608089015116906147b1565b90615086565b8082101561484257505b16825263ffffffff4216905290565b9050614833565b805180156148b95760200361487b578051602082810191830183900312613d3057519060ff821161487b575060ff1690565b611f7b906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613cae565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161463257565b60ff16604d811161463257600a0a90565b811561490e570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614a4357828411614a195790614982916148df565b91604d60ff84161180156149e0575b6149aa575050906149a4613e7a926148f3565b906147b1565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506149ea836148f3565b801561490e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411614991565b614a22916148df565b91604d60ff8416116149aa57505090614a3d613e7a926148f3565b90614904565b5050505090565b90816020910312613d3057518015158103613d305790565b73ffffffffffffffffffffffffffffffffffffffff600154163303614a8357565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90805115614ced5767ffffffffffffffff81516020830120921691826000526007602052614ae2816005604060002001615469565b15614ca95760005260086020526040600020815167ffffffffffffffff8111613bb057614b0f8254614510565b601f8111614c77575b506020601f8211600114614bb15791614b8b827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614ba195600091614ba6575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190613cae565b0390a2565b905084015138614b5a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b818110614c5f575092614ba19492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614c28575b5050811b0190556116e8565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880614c1c565b9192602060018192868a015181550194019201614be1565b614ca390836000526020600020601f840160051c810191602085106107f857601f0160051c019061479a565b38614b18565b5090611f7b6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613cae565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b15613d30576040517f79cc6790000000000000000000000000000000000000000000000000000000008152306004820152602481019190915260009182908290604490829084905af18015611b8c57614da8575050565b81614db291613c33565b50565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613e7a604082613c33565b9061ffff809167ffffffffffffffff614e0b60208601614213565b166000908152600a6020526040902090821615614e4c575460901c165b16908115614e4557612710916060614e419201356147b1565b0490565b5050600090565b5460801c16614e28565b815191929115614fd8576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff60208501511610614f755761459c91925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b606483614fd6604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408401511615801590615067575b6150065761459c9192614e99565b606483614fd6604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020840151161515614ff8565b9190820180921161463257565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811215613d3057016020813591019167ffffffffffffffff8211613d30578136038313613d3057565b906040519182815491828252602082019060005260206000209260005b81811061511557505061459c92500383613c33565b8454835260019485019487945060209093019201615100565b80548210156144e15760005260206000200190600090565b6000818152600660205260409020548015614e45577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161463257600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161463257818103615266575b5050506005548015615237577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016151f481600561512e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6152bd61527761528893600561512e565b90549060031b1c928392600561512e565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b905560005260066020526040600020553880806151bb565b9060018201918160005282602052604060002054801515600014615400577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614632578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614632578181036153c9575b50505080548015615237577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061538a828261512e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6153e96153d9615288938661512e565b90549060031b1c9283928661512e565b905560005283602052604060002055388080615352565b50505050600090565b806000526006602052604060002054156000146154635760055468010000000000000000811015613bb05761544a615288826001859401600555600561512e565b9055600554906000526006602052604060002055600190565b50600090565b6000828152600182016020526040902054614e455780549068010000000000000000821015613bb057826154a761528884600180960185558461512e565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615772575b61576c576fffffffffffffffffffffffffffffffff8216916001850190815461551663ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614625565b90816156ce575b5050848110615682575083831061557757505061554c6fffffffffffffffffffffffffffffffff928392614625565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615616578161558f91614625565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211614632576155dd6155e29273ffffffffffffffffffffffffffffffffffffffff96615086565b614904565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615742576156e9926148239160801c906147b1565b8084101561573d5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617865592388061551d565b6156f4565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156154d1565b919290156157f5575081511561578e575090565b3b156157975790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156158085750805190602001fd5b611f7b906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613cae56fea164736f6c634300081a000a",
}

var BurnFromMintTokenPoolABI = BurnFromMintTokenPoolMetaData.ABI

var BurnFromMintTokenPoolBin = BurnFromMintTokenPoolMetaData.Bin

func DeployBurnFromMintTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnFromMintTokenPool, error) {
	parsed, err := BurnFromMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnFromMintTokenPoolBin), backend, token, localTokenDecimals, advancedPoolHooks, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnFromMintTokenPool{address: address, abi: *parsed, BurnFromMintTokenPoolCaller: BurnFromMintTokenPoolCaller{contract: contract}, BurnFromMintTokenPoolTransactor: BurnFromMintTokenPoolTransactor{contract: contract}, BurnFromMintTokenPoolFilterer: BurnFromMintTokenPoolFilterer{contract: contract}}, nil
}

type BurnFromMintTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnFromMintTokenPoolCaller
	BurnFromMintTokenPoolTransactor
	BurnFromMintTokenPoolFilterer
}

type BurnFromMintTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnFromMintTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnFromMintTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnFromMintTokenPoolSession struct {
	Contract     *BurnFromMintTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnFromMintTokenPoolCallerSession struct {
	Contract *BurnFromMintTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnFromMintTokenPoolTransactorSession struct {
	Contract     *BurnFromMintTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnFromMintTokenPoolRaw struct {
	Contract *BurnFromMintTokenPool
}

type BurnFromMintTokenPoolCallerRaw struct {
	Contract *BurnFromMintTokenPoolCaller
}

type BurnFromMintTokenPoolTransactorRaw struct {
	Contract *BurnFromMintTokenPoolTransactor
}

func NewBurnFromMintTokenPool(address common.Address, backend bind.ContractBackend) (*BurnFromMintTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnFromMintTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnFromMintTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPool{address: address, abi: abi, BurnFromMintTokenPoolCaller: BurnFromMintTokenPoolCaller{contract: contract}, BurnFromMintTokenPoolTransactor: BurnFromMintTokenPoolTransactor{contract: contract}, BurnFromMintTokenPoolFilterer: BurnFromMintTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnFromMintTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnFromMintTokenPoolCaller, error) {
	contract, err := bindBurnFromMintTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolCaller{contract: contract}, nil
}

func NewBurnFromMintTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnFromMintTokenPoolTransactor, error) {
	contract, err := bindBurnFromMintTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolTransactor{contract: contract}, nil
}

func NewBurnFromMintTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnFromMintTokenPoolFilterer, error) {
	contract, err := bindBurnFromMintTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolFilterer{contract: contract}, nil
}

func bindBurnFromMintTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnFromMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnFromMintTokenPool.Contract.BurnFromMintTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.BurnFromMintTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.BurnFromMintTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnFromMintTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MinBlockConfirmations = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnFromMintTokenPool.Contract.GetDynamicConfig(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnFromMintTokenPool.Contract.GetDynamicConfig(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnFromMintTokenPool.Contract.GetFee(&_BurnFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnFromMintTokenPool.Contract.GetFee(&_BurnFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnFromMintTokenPool.Contract.GetRemotePools(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnFromMintTokenPool.Contract.GetRemotePools(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnFromMintTokenPool.Contract.GetRemoteToken(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnFromMintTokenPool.Contract.GetRemoteToken(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRequiredCCVs(&_BurnFromMintTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRequiredCCVs(&_BurnFromMintTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRmnProxy(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRmnProxy(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnFromMintTokenPool.Contract.GetSupportedChains(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnFromMintTokenPool.Contract.GetSupportedChains(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetToken(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetToken(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _BurnFromMintTokenPool.Contract.GetTokenDecimals(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnFromMintTokenPool.Contract.GetTokenDecimals(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnFromMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnFromMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsRemotePool(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsRemotePool(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsSupportedChain(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsSupportedChain(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsSupportedToken(&_BurnFromMintTokenPool.CallOpts, token)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnFromMintTokenPool.Contract.IsSupportedToken(&_BurnFromMintTokenPool.CallOpts, token)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) Owner() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.Owner(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.Owner(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnFromMintTokenPool.Contract.SupportsInterface(&_BurnFromMintTokenPool.CallOpts, interfaceId)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnFromMintTokenPool.Contract.SupportsInterface(&_BurnFromMintTokenPool.CallOpts, interfaceId)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnFromMintTokenPool.Contract.TypeAndVersion(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnFromMintTokenPool.Contract.TypeAndVersion(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.AcceptOwnership(&_BurnFromMintTokenPool.TransactOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.AcceptOwnership(&_BurnFromMintTokenPool.TransactOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.AddRemotePool(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.AddRemotePool(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyChainUpdates(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyChainUpdates(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnFromMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnFromMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.LockOrBurn(&_BurnFromMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.LockOrBurn(&_BurnFromMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.LockOrBurn0(&_BurnFromMintTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.LockOrBurn0(&_BurnFromMintTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ReleaseOrMint(&_BurnFromMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ReleaseOrMint(&_BurnFromMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ReleaseOrMint0(&_BurnFromMintTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ReleaseOrMint0(&_BurnFromMintTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.RemoveRemotePool(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.RemoveRemotePool(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setDynamicConfig", router, minBlockConfirmations, rateLimitAdmin)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetDynamicConfig(&_BurnFromMintTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetDynamicConfig(&_BurnFromMintTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetRateLimitConfig(&_BurnFromMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetRateLimitConfig(&_BurnFromMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.TransferOwnership(&_BurnFromMintTokenPool.TransactOpts, to)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.TransferOwnership(&_BurnFromMintTokenPool.TransactOpts, to)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.WithdrawFeeTokens(&_BurnFromMintTokenPool.TransactOpts, feeTokens, recipient)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.WithdrawFeeTokens(&_BurnFromMintTokenPool.TransactOpts, feeTokens, recipient)
}

type BurnFromMintTokenPoolChainAddedIterator struct {
	Event *BurnFromMintTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolChainAdded)
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
		it.Event = new(BurnFromMintTokenPoolChainAdded)
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

func (it *BurnFromMintTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolChainAddedIterator{contract: _BurnFromMintTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolChainAdded)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnFromMintTokenPoolChainAdded, error) {
	event := new(BurnFromMintTokenPoolChainAdded)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolChainRemovedIterator struct {
	Event *BurnFromMintTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolChainRemoved)
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
		it.Event = new(BurnFromMintTokenPoolChainRemoved)
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

func (it *BurnFromMintTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolChainRemovedIterator{contract: _BurnFromMintTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolChainRemoved)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnFromMintTokenPoolChainRemoved, error) {
	event := new(BurnFromMintTokenPoolChainRemoved)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _BurnFromMintTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _BurnFromMintTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolDynamicConfigSetIterator struct {
	Event *BurnFromMintTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolDynamicConfigSet)
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
		it.Event = new(BurnFromMintTokenPoolDynamicConfigSet)
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

func (it *BurnFromMintTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolDynamicConfigSet struct {
	Router                common.Address
	MinBlockConfirmations uint16
	RateLimitAdmin        common.Address
	Raw                   types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnFromMintTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolDynamicConfigSetIterator{contract: _BurnFromMintTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolDynamicConfigSet)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*BurnFromMintTokenPoolDynamicConfigSet, error) {
	event := new(BurnFromMintTokenPoolDynamicConfigSet)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolFeeTokenWithdrawnIterator struct {
	Event *BurnFromMintTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(BurnFromMintTokenPoolFeeTokenWithdrawn)
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

func (it *BurnFromMintTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolFeeTokenWithdrawn struct {
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*BurnFromMintTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolFeeTokenWithdrawnIterator{contract: _BurnFromMintTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolFeeTokenWithdrawn)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*BurnFromMintTokenPoolFeeTokenWithdrawn, error) {
	event := new(BurnFromMintTokenPoolFeeTokenWithdrawn)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolInboundRateLimitConsumedIterator struct {
	Event *BurnFromMintTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(BurnFromMintTokenPoolInboundRateLimitConsumed)
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

func (it *BurnFromMintTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolInboundRateLimitConsumedIterator{contract: _BurnFromMintTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolInboundRateLimitConsumed)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolInboundRateLimitConsumed, error) {
	event := new(BurnFromMintTokenPoolInboundRateLimitConsumed)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolLockedOrBurnedIterator struct {
	Event *BurnFromMintTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolLockedOrBurned)
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
		it.Event = new(BurnFromMintTokenPoolLockedOrBurned)
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

func (it *BurnFromMintTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolLockedOrBurnedIterator{contract: _BurnFromMintTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolLockedOrBurned)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*BurnFromMintTokenPoolLockedOrBurned, error) {
	event := new(BurnFromMintTokenPoolLockedOrBurned)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *BurnFromMintTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(BurnFromMintTokenPoolOutboundRateLimitConsumed)
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

func (it *BurnFromMintTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolOutboundRateLimitConsumedIterator{contract: _BurnFromMintTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolOutboundRateLimitConsumed)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolOutboundRateLimitConsumed, error) {
	event := new(BurnFromMintTokenPoolOutboundRateLimitConsumed)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnFromMintTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolOwnershipTransferRequested)
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
		it.Event = new(BurnFromMintTokenPoolOwnershipTransferRequested)
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

func (it *BurnFromMintTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnFromMintTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolOwnershipTransferRequestedIterator{contract: _BurnFromMintTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolOwnershipTransferRequested)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnFromMintTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnFromMintTokenPoolOwnershipTransferRequested)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolOwnershipTransferredIterator struct {
	Event *BurnFromMintTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolOwnershipTransferred)
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
		it.Event = new(BurnFromMintTokenPoolOwnershipTransferred)
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

func (it *BurnFromMintTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnFromMintTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolOwnershipTransferredIterator{contract: _BurnFromMintTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolOwnershipTransferred)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnFromMintTokenPoolOwnershipTransferred, error) {
	event := new(BurnFromMintTokenPoolOwnershipTransferred)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolRateLimitConfiguredIterator struct {
	Event *BurnFromMintTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolRateLimitConfigured)
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
		it.Event = new(BurnFromMintTokenPoolRateLimitConfigured)
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

func (it *BurnFromMintTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolRateLimitConfiguredIterator{contract: _BurnFromMintTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolRateLimitConfigured)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*BurnFromMintTokenPoolRateLimitConfigured, error) {
	event := new(BurnFromMintTokenPoolRateLimitConfigured)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolReleasedOrMintedIterator struct {
	Event *BurnFromMintTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolReleasedOrMinted)
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
		it.Event = new(BurnFromMintTokenPoolReleasedOrMinted)
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

func (it *BurnFromMintTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolReleasedOrMintedIterator{contract: _BurnFromMintTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolReleasedOrMinted)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*BurnFromMintTokenPoolReleasedOrMinted, error) {
	event := new(BurnFromMintTokenPoolReleasedOrMinted)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolRemotePoolAddedIterator struct {
	Event *BurnFromMintTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolRemotePoolAdded)
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
		it.Event = new(BurnFromMintTokenPoolRemotePoolAdded)
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

func (it *BurnFromMintTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolRemotePoolAddedIterator{contract: _BurnFromMintTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolRemotePoolAdded)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*BurnFromMintTokenPoolRemotePoolAdded, error) {
	event := new(BurnFromMintTokenPoolRemotePoolAdded)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolRemotePoolRemovedIterator struct {
	Event *BurnFromMintTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolRemotePoolRemoved)
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
		it.Event = new(BurnFromMintTokenPoolRemotePoolRemoved)
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

func (it *BurnFromMintTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolRemotePoolRemovedIterator{contract: _BurnFromMintTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolRemotePoolRemoved)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*BurnFromMintTokenPoolRemotePoolRemoved, error) {
	event := new(BurnFromMintTokenPoolRemotePoolRemoved)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *BurnFromMintTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(BurnFromMintTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *BurnFromMintTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnFromMintTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _BurnFromMintTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolTokenTransferFeeConfigDeleted)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnFromMintTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(BurnFromMintTokenPoolTokenTransferFeeConfigDeleted)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *BurnFromMintTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(BurnFromMintTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *BurnFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _BurnFromMintTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolTokenTransferFeeConfigUpdated)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnFromMintTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(BurnFromMintTokenPoolTokenTransferFeeConfigUpdated)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (BurnFromMintTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnFromMintTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (BurnFromMintTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a117")
}

func (BurnFromMintTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (BurnFromMintTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnFromMintTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnFromMintTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnFromMintTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnFromMintTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnFromMintTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (BurnFromMintTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnFromMintTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnFromMintTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnFromMintTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (BurnFromMintTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPool) Address() common.Address {
	return _BurnFromMintTokenPool.address
}

type BurnFromMintTokenPoolInterface interface {
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

	FilterChainAdded(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnFromMintTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnFromMintTokenPoolChainRemoved, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnFromMintTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnFromMintTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*BurnFromMintTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*BurnFromMintTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnFromMintTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnFromMintTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnFromMintTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnFromMintTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnFromMintTokenPoolOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*BurnFromMintTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnFromMintTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnFromMintTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnFromMintTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnFromMintTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnFromMintTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnFromMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnFromMintTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
