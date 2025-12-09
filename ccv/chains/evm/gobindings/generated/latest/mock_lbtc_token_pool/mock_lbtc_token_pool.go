// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_lbtc_token_pool

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

var MockE2ELBTCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"advancedPoolHooks\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"tokenArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_destPoolData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.RateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"minBlockConfirmations\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"customBlockConfirmation\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"TokenDataMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x6101008060405234610463576162c9803803809161001d8285610468565b833981019060a0818303126104635780516001600160a01b038116908190036104635761004c6020830161048b565b906100596040840161048b565b6100656060850161048b565b608085015190946001600160401b038211610463570185601f82011215610463578051906001600160401b03821161038d57604051966100af601f8401601f191660200189610468565b828852602083830101116104635760005b82811061044e575050602060009187010152331561043d57600180546001600160a01b031916331790558115801561042c575b801561041b575b61040a578160209160049360805260c0526040519283809263313ce56760e01b82525afa80916000916103c7575b50906103a3575b50600860a0526001600160a01b0390811660e052600280546001600160a01b0319169290911691909117905580516001600160401b03811161038d57600b54600181811c91168015610383575b602082101461036d57601f8111610308575b50602091601f82116001146102a457918192600092610299575b50508160011b916000199060031b1c191617600b555b604051615e2990816104a0823960805181818161191901528181611b2301528181611db601528181611e99015281816123750152818161252c01528181612de101528181612fd4015281816130b601528181613357015281816135b90152818161378701528181613ed30152613f4b015260a051818181611bac01528181613d5d01528181614f840152615007015260c051818181610d3e015281816119b40152818161240f01528181612e7c0152613654015260e051818181611ad8015281816125b30152818161303c0152818161380e0152614abe0152f35b0151905038806101a8565b601f19821692600b600052806000209160005b8581106102f0575083600195106102d7575b505050811b01600b556101be565b015160001960f88460031b161c191690553880806102c9565b919260206001819286850151815501940192016102b7565b600b6000527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9601f830160051c81019160208410610363575b601f0160051c01905b818110610357575061018e565b6000815560010161034a565b9091508190610341565b634e487b7160e01b600052602260045260246000fd5b90607f169061017c565b634e487b7160e01b600052604160045260246000fd5b60ff166008811461012f576332ad3e0760e11b600052600860045260245260446000fd5b6020813d602011610402575b816103e060209383610468565b810103126103fe57519060ff821682036103fb575038610128565b80fd5b5080fd5b3d91506103d3565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100fa565b506001600160a01b038416156100f3565b639b15e16f60e01b60005260046000fd5b80602080928401015182828b010152016100c0565b600080fd5b601f909101601f19168101906001600160401b0382119082101761038d57604052565b51906001600160a01b03821682036104635756fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714613fee57508063181f5a7714613f6f57806321df0da714613f00578063240028e814613e7e5780632422ac4514613d8157806324f65ee714613d255780632c06340414613c6e578063390775371461350b5780633e5db5d1146134d1578063489a68f214612d005780634c5ef0ed14612c9b57806362ddd3c414612c145780637437ff9f14612ba557806379ba509714612ac05780638926f54f14612a5c57806389720a62146129775780638da5cb5b146129255780639a4575b9146122bf578063a42a7b8b1461213a578063acfecf9114612042578063b1c71c6514611837578063b7946580146117dc578063c4bffe2b14611693578063c7230a6014611329578063d8aa3f40146111d1578063dc04fa1f14610d62578063dc0bd97114610cf3578063dcbd41bc14610ad1578063e8a1da17146103f3578063f2fde38b146103065763fdf168751461017657600080fd5b346103035760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610303576101ad61428e565b906101b661434a565b6044359273ffffffffffffffffffffffffffffffffffffffff84168085036102ff576101e0615111565b73ffffffffffffffffffffffffffffffffffffffff821680156102d757946102d1917fba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a1179596600254907fffffffffffffffffffff0000000000000000000000000000000000000000000075ffff00000000000000000000000000000000000000008860a01b16921617176002557fffffffffffffffffffffffff000000000000000000000000000000000000000060095416176009556040519384938491604091949361ffff73ffffffffffffffffffffffffffffffffffffffff9283606087019816865216602085015216910152565b0390a180f35b6004857f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b80fd5b50346103035760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035773ffffffffffffffffffffffffffffffffffffffff61035361428e565b61035b615111565b163381146103cb57807fffffffffffffffffffffffff000000000000000000000000000000000000000083541617825573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b50346103035760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035760043567ffffffffffffffff811161092e57610443903690600401614675565b9060243567ffffffffffffffff81116102ff579061046684923690600401614675565b939091610471615111565b83905b8282106109365750505081927ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610932578060051b8301358581121561092a5783016101208136031261092a57604051946104d886614191565b6104e182614328565b8652602082013567ffffffffffffffff811161092e5782019436601f8701121561092e5785359561051187614a29565b9661051f60405198896141ad565b80885260208089019160051b8301019036821161092a5760208301905b8282106108f7575050505060208701958652604083013567ffffffffffffffff81116108f35761056f90369085016145ca565b91604088019283526105996105873660608701614dad565b9460608a0195865260c0369101614dad565b9560808901968752835151156108cb576105bd67ffffffffffffffff8a51166159df565b156108945767ffffffffffffffff89511682526007602052604082206105e486518261542c565b6105f288516002830161542c565b6004855191019080519067ffffffffffffffff82116108675761061583546143a9565b601f811161082c575b50602090601f831160011461078d5761066c9291869183610782575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b815b885180518210156106a657906106a06001926106998367ffffffffffffffff8f511692614c91565b519061515c565b01610671565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561077467ffffffffffffffff600197969498511692519351915161074061070b6040519687968752610100602088015261010087019061424b565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a10193929091936104a7565b015190508e8061063a565b83865281862091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416875b81811061081457509084600195949392106107dd575b505050811b01905561066f565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806107d0565b929360206001819287860151815501950193016107ba565b6108579084875260208720601f850160051c8101916020861061085d575b601f0160051c0190614e49565b8d61061e565b909150819061084a565b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60248267ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b8280fd5b813567ffffffffffffffff81116109265760209161091b83928336918901016145ca565b81520191019061053c565b8680fd5b8480fd5b5080fd5b8380f35b9267ffffffffffffffff6109586109538486889a9699979a614d32565b6149d7565b16916109638361571c565b15610aa557828452600760205261097f600560408620016156b9565b94845b86518110156109b85760019085875260076020526109b1600560408920016109aa838b614c91565b51906158ab565b5001610982565b50939692909450949094808752600760205260056040882088815588600182015588600282015588600382015588600482016109f481546143a9565b80610a64575b5050500180549088815581610a46575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101909194939294610474565b885260208820908101905b81811015610a0a57888155600101610a51565b601f8111600114610a7a5750555b888a806109fa565b81835260208320610a9591601f01861c810190600101614e49565b8082528160208120915555610a72565b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50346103035760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035760043567ffffffffffffffff811161092e57610b219036906004016146a6565b73ffffffffffffffffffffffffffffffffffffffff6009541633141580610cd1575b610ca557825b818110610b54578380f35b610b5f818385614d42565b67ffffffffffffffff610b71826149d7565b1690610b8a826000526006602052604060002054151590565b15610c7957907f41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb60e083610c39610c13602060019897018b610bcb82614d52565b15610c40578790526003602052610bf260408d20610bec3660408801614dad565b9061542c565b868c526004602052610c0e60408d20610bec3660a08801614dad565b614d52565b916040519215158352610c2c6020840160408301614e05565b60a0608084019101614e05565ba201610b49565b60026040828a610c0e94526007602052610c62828220610bec36858c01614dad565b8a815260076020522001610bec3660a08801614dad565b602486837f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5073ffffffffffffffffffffffffffffffffffffffff60015416331415610b43565b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346103035760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035760043567ffffffffffffffff811161092e57610db29036906004016146a6565b60243567ffffffffffffffff81116102ff57610dd2903690600401614675565b919092610ddd615111565b845b828110610e4957505050825b818110610df6578380f35b8067ffffffffffffffff610e106109536001948688614d32565b16808652600a6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201610deb565b610e57610953828585614d42565b610e62828585614d42565b90602082019060e0830190610e7682614d52565b1561119c5760a0840161271061ffff610e8e83614d5f565b16101561118d5760c085019161271061ffff610ea985614d5f565b1610156111555763ffffffff610ebe86614d6e565b16156111205767ffffffffffffffff1694858c52600a60205260408c20610ee486614d6e565b63ffffffff16908054906040840191610efc83614d6e565b60201b67ffffffff0000000016936060860194610f1886614d6e565b60401b6bffffffff0000000000000000169660800196610f3788614d6e565b60601b6fffffffff0000000000000000000000001691610f568a614d5f565b60801b71ffff000000000000000000000000000000001693610f778c614d5f565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717815561102a87614d52565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161790556040519661107b90614d7f565b63ffffffff16875261108c90614d7f565b63ffffffff1660208701526110a090614d7f565b63ffffffff1660408601526110b490614d7f565b63ffffffff1660608501526110c89061436c565b61ffff1660808401526110da9061436c565b61ffff1660a08301526110ec9061433d565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610ddf565b7f12332265000000000000000000000000000000000000000000000000000000008c5267ffffffffffffffff1660045260248bfd5b60248c61ffff61116486614d5f565b7f95f3517a00000000000000000000000000000000000000000000000000000000835216600452fd5b8a61ffff611164602493614d5f565b7f12332265000000000000000000000000000000000000000000000000000000008a5267ffffffffffffffff16600452602489fd5b50346103035760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035761120961428e565b50611212614311565b61121a61435b565b5060643567ffffffffffffffff81116108f3579167ffffffffffffffff60409261124a60e095369060040161437b565b50508260c0855161125a81614175565b82815282602082015282878201528260608201528260808201528260a08201520152168152600a602052206040519061129282614175565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b50346103035760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035760043567ffffffffffffffff811161092e57611379903690600401614675565b6024359073ffffffffffffffffffffffffffffffffffffffff82169081830361092a576113a4615111565b845b8181106113b1578580f35b80602073ffffffffffffffffffffffffffffffffffffffff6113de6113d9602495878b614d32565b6149b6565b16604051938480927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa80156116885784869185948a9161164c575b5080611437575b50505060019150016113a6565b8861152f73ffffffffffffffffffffffffffffffffffffffff61145e6113d9888a86614d32565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff98909816602482015260448082018790528152918e91166114c06064846141ad565b81806040998a51956114d28c886141ad565b602087527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020880152519082855af1903d15611643573d611513816141ee565b906115208b5192836141ad565b8152809360203d92013e615d50565b805180611594575b50507f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e9173ffffffffffffffffffffffffffffffffffffffff6115826113d98860019a602096614d32565b169451908152a382915084843861142a565b8160209350839294969798506115ae955001019101614ef8565b156115c0579084849392883880611537565b608482517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60609250615d50565b93945050505060203d8111611681575b61166681836141ad565b60208260009281010312610303575090848484935138611423565b503d61165c565b6040513d89823e3d90fd5b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357604051906005548083528260208101600584526020842092845b8181106117c35750506116f1925003836141ad565b81516117156116ff82614a29565b9161170d60405193846141ad565b808352614a29565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b8451811015611774578067ffffffffffffffff61176160019388614c91565b511661176d8286614c91565b5201611742565b50925090604051928392602084019060208552518091526040840192915b8181106117a0575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611792565b84548352600194850194879450602090930192016116dc565b50346103035760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035761183361181f61181a6142fa565b614d10565b60405191829160208352602083019061424b565b0390f35b50346103035760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035760043567ffffffffffffffff811161092e578060040160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83360301126108f3576118b161434a565b926044359367ffffffffffffffff851161092e576118d66118f395369060040161437b565b95906118e0614c78565b506118eb83866153c6565b963691614593565b6084850190611901826149b6565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611ff857602486019577ffffffffffffffff00000000000000000000000000000000611967886149d7565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611f6b578691611fc9575b50611fa15767ffffffffffffffff6119fb886149d7565b16611a13816000526006602052604060002054151590565b15611f7657602073ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611f6b578690611f1a575b73ffffffffffffffffffffffffffffffffffffffff9150163303611eee5760648101359561ffff611aa58a89614cd4565b9516948515611e3d5761ffff60025460a01c1680611d52575b50505b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169283611c12575b611c0889611ba461181a611b138e8d614cd4565b92611b1d816149d7565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a26149d7565b9060405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152611be06040826141ad565b60405192611bed84614159565b8352602083015260405192839260408452604084019061464b565b9060208301520390f35b833b1561092657869493929185918a604051988997889687957f5c3af7ca000000000000000000000000000000000000000000000000000000008752600487016060905280611c6091615669565b6064880160a09052610104880190611c7792614a62565b93611c8190614328565b67ffffffffffffffff166084870152604401611c9c906142d9565b73ffffffffffffffffffffffffffffffffffffffff1660a48601528c60c4860152611cc6906142d9565b73ffffffffffffffffffffffffffffffffffffffff1660e485015260248401528281037ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc016044840152611d199161424b565b03925af18015611d4757611d32575b8080808080611aff565b611d3d8280926141ad565b6103035780611d28565b6040513d84823e3d90fd5b808710611e0d57507f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611d8d8b6149d7565b1691828952600360205280611dde60408b2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615a94565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a23880611abe565b87604491887f7911d95b000000000000000000000000000000000000000000000000000000008352600452602452fd5b7fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611e708b6149d7565b1691828952600760205280611ec160408b2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615a94565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611ac1565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611f63575b81611f34602093836141ad565b81010312611f5f57611f5a73ffffffffffffffffffffffffffffffffffffffff91614a41565b611a74565b8580fd5b3d9150611f27565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611feb915060203d602011611ff1575b611fe381836141ad565b810190614ef8565b386119e4565b503d611fd9565b60248473ffffffffffffffffffffffffffffffffffffffff612019856149b6565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b50346103035767ffffffffffffffff61205a366145e8565b929091612065615111565b169161207e836000526006602052604060002054151590565b15610aa55782845260076020526120ad600560408620016120a0368486614593565b60208151910120906158ab565b156120f257907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76916120ec604051928392602084526020840191614a62565b0390a280f35b82612136836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614a62565b0390fd5b50346103035760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035767ffffffffffffffff61217b6142fa565b1681526007602052612192600560408320016156b9565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06121d76121c183614a29565b926121cf60405194856141ad565b808452614a29565b01835b8181106122ae575050825b825181101561222b57806121fb60019285614c91565b518552600860205261220f604086206144d3565b6122198285614c91565b526122248184614c91565b50016121e5565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061226357505050500390f35b9193602061229e827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06001959799849503018652885161424b565b9601920192018594939192612254565b8060606020809386010152016121da565b50346103035760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610303576004359067ffffffffffffffff821161030357816004019160a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261092e5761233b614c78565b506020926040519261234d85856141ad565b8084526084830161235d816149b6565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361290457602484019477ffffffffffffffff000000000000000000000000000000006123c3876149d7565b60801b16604051907f2cbc26bb0000000000000000000000000000000000000000000000000000000082526004820152878160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156128895784916128e7575b506128bf5767ffffffffffffffff612456876149d7565b1661246e816000526006602052604060002054151590565b15612894578773ffffffffffffffffffffffffffffffffffffffff60025416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015612889578490612841575b73ffffffffffffffffffffffffffffffffffffffff915016330361281557606485013594859467ffffffffffffffff612506896149d7565b1680865260078a526125546040872073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016998a91615a94565b6040805173ffffffffffffffffffffffffffffffffffffffff8a168152602081018990527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692836126de575b5050505050823b15610303576040517f42966c68000000000000000000000000000000000000000000000000000000008152826004820152818160248183885af18015611d47576126c9575b8561269961181a87877ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108867ffffffffffffffff612663856149d7565b6040805173ffffffffffffffffffffffffffffffffffffffff909616865233602087015285019290925216918060608101611b9c565b90604051916126a783614159565b82526126b16143fc565b8183015261183360405192828493845283019061464b565b6126d48280926141ad565b6103035780612626565b833b15611f5f5791858094928a9694604051978896879586947f5c3af7ca00000000000000000000000000000000000000000000000000000000865260048601606090528061272c91615669565b6064870160a0905261010487019061274392614a62565b9261274d90614328565b67ffffffffffffffff166084860152604401612768906142d9565b73ffffffffffffffffffffffffffffffffffffffff1660a48501528b60c4850152612792906142d9565b73ffffffffffffffffffffffffffffffffffffffff1660e48401528360248401528281037ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0160448401526127e69161424b565b03925af18015611d4757908291612800575b8080806125da565b8161280a916141ad565b6103035780386127f8565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b508781813d8311612882575b61285781836141ad565b810103126102ff5761287d73ffffffffffffffffffffffffffffffffffffffff91614a41565b6124ce565b503d61284d565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6128fe9150883d8a11611ff157611fe381836141ad565b3861243f565b9073ffffffffffffffffffffffffffffffffffffffff6120196024936149b6565b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b50346103035760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610303576129af61428e565b6129b7614311565b9060643561ffff811681036102ff5760843567ffffffffffffffff811161092a576129e690369060040161437b565b9160a43593600285101561092657612a019560443591614aa1565b90604051918291602083016020845282518091526020604085019301915b818110612a2d575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101612a1f565b50346103035760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610303576020612ab667ffffffffffffffff612aa26142fa565b166000526006602052604060002054151590565b6040519015158152f35b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357805473ffffffffffffffffffffffffffffffffffffffff81163303612b7d577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516825573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610303576002546009546040805173ffffffffffffffffffffffffffffffffffffffff808516825260a09490941c61ffff1660208201529290911690820152606090f35b503461030357612c23366145e8565b612c2f93929193615111565b67ffffffffffffffff8216612c51816000526006602052604060002054151590565b15612c705750612c6d9293612c67913691614593565b9061515c565b80f35b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b50346103035760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357612cd36142fa565b906024359067ffffffffffffffff8211610303576020612ab684612cfa36600487016145ca565b906149ec565b50346103035760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610303576004359067ffffffffffffffff821161030357816004016101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc843603011261092e57612d7c61434a565b9180604051612d8a8161410e565b5260648401359260c48501612dba612db4612daf612da8848861492a565b3691614593565b614f10565b86615004565b946084870190612dc9826149b6565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036134b057602488019577ffffffffffffffff00000000000000000000000000000000612e2f886149d7565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611688578791613491575b506134695767ffffffffffffffff612ec3886149d7565b16612edb816000526006602052604060002054151590565b1561343e57602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561168857879161341f575b50156133f357612f52876149d7565b94612f6860a48b0196612cfa612da8898661492a565b156133ac5761ffff169081156132f85767ffffffffffffffff612f8a896149d7565b1680885260046020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a80612ffc60408c2073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615a94565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169485613125575b60208a60448d017ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc067ffffffffffffffff6130b06130aa8f6130a4866149b6565b506149d7565b936149b6565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081168252336020830152909216908201526060810185905292169180608081015b0390a28060405161311c8161410e565b52604051908152f35b853b156132f45792889694928b8997959388946040519a8b998a9889977f5eff3bf700000000000000000000000000000000000000000000000000000000895260048901606090526131778680615669565b60648b0161010090526101648b019061318f92614a62565b9361319990614328565b67ffffffffffffffff1660848a01526131b4604488016142d9565b73ffffffffffffffffffffffffffffffffffffffff1660a48a015260c48901526131dd906142d9565b73ffffffffffffffffffffffffffffffffffffffff1660e48801526132029084615669565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c888403016101048901526132379291614a62565b906132429083615669565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c878403016101248801526132779291614a62565b9160e40161328491615669565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c858403016101448601526132b99291614a62565b908b6024840152604483015203925af18015611d47576132df575b808080808080613063565b6132ea8280926141ad565b61030357806132d4565b8780fd5b67ffffffffffffffff61330a896149d7565b1680885260076020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8a8061337f600260408d200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615a94565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2613025565b6133b6868361492a565b6121366040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614a62565b6024867f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b613438915060203d602011611ff157611fe381836141ad565b38612f43565b7fa9902c7e000000000000000000000000000000000000000000000000000000008752600452602486fd5b6004867f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b6134aa915060203d602011611ff157611fe381836141ad565b38612eac565b60248573ffffffffffffffffffffffffffffffffffffffff612019856149b6565b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035761183361181f6143fc565b50346103035760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610303576004359067ffffffffffffffff821161030357816004016101007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc843603011261092e578160405161358c8161410e565b5260648301359060848401936135a1856149b6565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603613c4d57602481019177ffffffffffffffff00000000000000000000000000000000613607846149d7565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611f6b578691613c2e575b50611fa15767ffffffffffffffff61369b846149d7565b166136b3816000526006602052604060002054151590565b15611f7657602073ffffffffffffffffffffffffffffffffffffffff60025416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611f6b578691613c0f575b5015611eee5761372a836149d7565b9561374060a4840197612cfa612da88a8661492a565b15613c055785809786975067ffffffffffffffff61375d876149d7565b1680835260076020526137af6002604085200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016998a91615a94565b6040805173ffffffffffffffffffffffffffffffffffffffff8a168152602081018b90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680613a40575b505050506020613845600b546143a9565b14613951575b5060440193613859856149b6565b833b1561092e576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91909116600482015260248101859052818160448183885af18015611d475761393c575b505067ffffffffffffffff60209461310c856138fe6130aa7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966149d7565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b6139478280926141ad565b61030357806138bf565b8561395f60e484018361492a565b8101919060408184031261092e57803567ffffffffffffffff81116108f3578361398a9183016145ca565b60208201359167ffffffffffffffff83116102ff576020946139c4936139b092016145ca565b508360405192828480945193849201614228565b8101039060025afa15613a355785519060c48301906139ec6139e6838361492a565b9061497b565b83036139f957505061384b565b91613a0a6139e6899360449561492a565b7f7f249311000000000000000000000000000000000000000000000000000000008352600452602452fd5b6040513d87823e3d90fd5b803b156108f3578693838a604051968795869485937f5eff3bf70000000000000000000000000000000000000000000000000000000085526004850160609052613a8a8c80615669565b606487016101009052610164870190613aa292614a62565b93613aac90614328565b67ffffffffffffffff166084860152613ac760448e016142d9565b73ffffffffffffffffffffffffffffffffffffffff1660a486015260c4850152613af0906142d9565b73ffffffffffffffffffffffffffffffffffffffff1660e4840152613b15908a615669565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610104850152613b4a9291614a62565b613b5760c48b018a615669565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610124850152613b8c9291614a62565b613b9960e48b018a615669565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9c84840301610144850152613bce9291614a62565b8c602483015282604483015203925af1801561168857613bf1575b868180613834565b86613bfe919792976141ad565b9438613be9565b6133b6878361492a565b613c28915060203d602011611ff157611fe381836141ad565b3861371b565b613c47915060203d602011611ff157611fe381836141ad565b38613684565b60248473ffffffffffffffffffffffffffffffffffffffff612019886149b6565b50346103035760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357613ca661428e565b50613caf614311565b613cb76142b6565b506084359161ffff831683036103035760a4359067ffffffffffffffff82116103035760a063ffffffff8061ffff613cfe8888613cf73660048b0161437b565b50506147eb565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346103035760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357613db96142fa565b60243591821515830361030357610140613e7c613dd68585614768565b613e2c60409392935180946fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906fffffffffffffffffffffffffffffffff6080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b50346103035760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357602090613eb961428e565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261030357602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461030357807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126103035750611833604051613fb06040826141ad565b601a81527f4d6f636b4532454c425443546f6b656e506f6f6c20312e352e31000000000000602082015260405191829160208352602083019061424b565b90503461092e5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261092e576004357fffffffff0000000000000000000000000000000000000000000000000000000081168091036108f357602092507faff2afbf0000000000000000000000000000000000000000000000000000000081149081156140e4575b81156140ba575b8115614090575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438614089565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614082565b7f33171031000000000000000000000000000000000000000000000000000000008114915061407b565b6020810190811067ffffffffffffffff82111761412a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040810190811067ffffffffffffffff82111761412a57604052565b60e0810190811067ffffffffffffffff82111761412a57604052565b60a0810190811067ffffffffffffffff82111761412a57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761412a57604052565b67ffffffffffffffff811161412a57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b83811061423b5750506000910152565b818101518382015260200161422b565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361428781518092818752878088019101614228565b0116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036142b157565b600080fd5b6064359073ffffffffffffffffffffffffffffffffffffffff821682036142b157565b359073ffffffffffffffffffffffffffffffffffffffff821682036142b157565b6004359067ffffffffffffffff821682036142b157565b6024359067ffffffffffffffff821682036142b157565b359067ffffffffffffffff821682036142b157565b359081151582036142b157565b6024359061ffff821682036142b157565b6044359061ffff821682036142b157565b359061ffff821682036142b157565b9181601f840112156142b15782359167ffffffffffffffff83116142b157602083818601950101116142b157565b90600182811c921680156143f2575b60208310146143c357565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916143b8565b60405190600082600b5491614410836143a9565b80835292600181169081156144965750600114614436575b614434925003836141ad565b565b50600b600090815290917f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db95b81831061447a57505090602061443492820101614428565b6020919350806001915483858901015201910190918492614462565b602092506144349491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101614428565b90604051918260008254926144e7846143a9565b8084529360018116908115614553575060011461450c575b50614434925003836141ad565b90506000929192526020600020906000915b81831061453757505090602061443492820101386144ff565b602091935080600191548385890101520191019091849261451e565b602093506144349592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386144ff565b92919261459f826141ee565b916145ad60405193846141ad565b8294818452818301116142b1578281602093846000960137010152565b9080601f830112156142b1578160206145e593359101614593565b90565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126142b15760043567ffffffffffffffff811681036142b157916024359067ffffffffffffffff82116142b1576146479160040161437b565b9091565b6145e5916020614664835160408452604084019061424b565b92015190602081840391015261424b565b9181601f840112156142b15782359167ffffffffffffffff83116142b1576020808501948460051b0101116142b157565b9181601f840112156142b15782359167ffffffffffffffff83116142b1576020808501948460081b0101116142b157565b604051906146e482614191565b60006080838281528260208201528260408201528260608201520152565b9060405161470f81614191565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff9161477a6146d7565b506147836146d7565b506147b7571660005260076020526040600020906145e56147ab60026147b06147ab86614702565b614e73565b9401614702565b16908160005260036020526147d26147ab6040600020614702565b9160005260046020526145e56147ab6040600020614702565b67ffffffffffffffff16600052600a60205260406000206040519061480f82614175565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c168152608085019561ffff8360801c16875260ff60a087019361ffff8160901c16855260a01c1615801560c08801526149115761ffff16806148a85750505063ffffffff808061ffff9351169451169551169351169193929190600190565b919550915061ffff60025460a01c16908181106148e157505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050505092505050600090600090600090600090600090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156142b1570180359067ffffffffffffffff82116142b1576020019181360383136142b157565b359060208110614989575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b3573ffffffffffffffffffffffffffffffffffffffff811681036142b15790565b3567ffffffffffffffff811681036142b15790565b9067ffffffffffffffff6145e592166000526007602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff811161412a5760051b60200190565b519073ffffffffffffffffffffffffffffffffffffffff821682036142b157565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b95939192949073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016958615614c5657614b5a9467ffffffffffffffff61ffff9373ffffffffffffffffffffffffffffffffffffffff6040519b7f89720a62000000000000000000000000000000000000000000000000000000008d521660048c01521660248a0152604489015216606487015260c0608487015260c4860191614a62565b916002821015614c27578380600094819460a483015203915afa908115614c1b57600091614b86575090565b3d8083833e614b9581836141ad565b8101906020818303126108f35780519067ffffffffffffffff82116102ff570181601f820112156108f357805190614bcc82614a29565b93614bda60405195866141ad565b82855260208086019360051b8301019384116103035750602001905b828210614c035750505090565b60208091614c1084614a41565b815201910190614bf6565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b5050505050505050604051614c6c6020826141ad565b60008152600036813790565b60405190614c8582614159565b60606020838281520152565b8051821015614ca55760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908203918211614ce157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff1660005260076020526145e560046040600020016144d3565b9190811015614ca55760051b0190565b9190811015614ca55760081b0190565b3580151581036142b15790565b3561ffff811681036142b15790565b3563ffffffff811681036142b15790565b359063ffffffff821682036142b157565b35906fffffffffffffffffffffffffffffffff821682036142b157565b91908260609103126142b1576040516060810181811067ffffffffffffffff82111761412a576040526040614e00818395614de78161433d565b8552614df560208201614d90565b602086015201614d90565b910152565b6fffffffffffffffffffffffffffffffff614e4360408093614e268161433d565b1515865283614e3760208301614d90565b16602087015201614d90565b16910152565b818110614e54575050565b60008155600101614e49565b81810292918115918404141715614ce157565b614e7b6146d7565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691614ed86020850193614ed2614ec563ffffffff87511642614cd4565b8560808901511690614e60565b9061565c565b80821015614ef157505b16825263ffffffff4216905290565b9050614ee2565b908160209103126142b1575180151581036142b15790565b80518015614f8057602003614f425780516020828101918301839003126142b157519060ff8211614f42575060ff1690565b612136906040519182917f953576f700000000000000000000000000000000000000000000000000000000835260206004840152602483019061424b565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211614ce157565b60ff16604d8111614ce157600a0a90565b8115614fd5570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff81169282841461510a578284116150e0579061504991614fa6565b91604d60ff84161180156150a7575b6150715750509061506b6145e592614fba565b90614e60565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506150b183614fba565b8015614fd5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411615058565b6150e991614fa6565b91604d60ff841611615071575050906151046145e592614fba565b90614fcb565b5050505090565b73ffffffffffffffffffffffffffffffffffffffff60015416330361513257565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9080511561539c5767ffffffffffffffff81516020830120921691826000526007602052615191816005604060002001615a3f565b156153585760005260086020526040600020815167ffffffffffffffff811161412a576151be82546143a9565b601f8111615326575b506020601f8211600114615260579161523a827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361525095600091615255575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b905560405191829160208352602083019061424b565b0390a2565b905084015138615209565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b81811061530e5750926152509492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106152d7575b5050811b01905561181f565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538806152cb565b9192602060018192868a015181550194019201615290565b61535290836000526020600020601f840160051c8101916020851061085d57601f0160051c0190614e49565b386151c7565b50906121366040519283927f393b8ad2000000000000000000000000000000000000000000000000000000008452600484015260406024840152604483019061424b565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9061ffff809167ffffffffffffffff6153e1602086016149d7565b166000908152600a6020526040902090821615615422575460901c165b1690811561541b57612710916060615417920135614e60565b0490565b5050600090565b5460801c166153fe565b8151919291156155ae576fffffffffffffffffffffffffffffffff6040840151166fffffffffffffffffffffffffffffffff6020850151161061554b5761443491925b805182547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161782556020810151825460409290920151608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff9290921691821760018501557fffffffffffffffffffffffff0000000000000000000000000000000000000000909216174290911b73ffffffff0000000000000000000000000000000016179055565b6064836155ac604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040840151161580159061563d575b6155dc57614434919261546f565b6064836155ac604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208401511615156155ce565b91908201809211614ce157565b90357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1823603018112156142b157016020813591019167ffffffffffffffff82116142b15781360383136142b157565b906040519182815491828252602082019060005260206000209260005b8181106156eb575050614434925003836141ad565b84548352600194850194879450602090930192016156d6565b8054821015614ca55760005260206000200190600090565b600081815260066020526040902054801561541b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614ce157600554907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614ce15781810361583c575b505050600554801561580d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016157ca816005615704565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600555600052600660205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61589361584d61585e936005615704565b90549060031b1c9283926005615704565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526006602052604060002055388080615791565b90600182019181600052826020526040600020548015156000146159d6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614ce1578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614ce15781810361599f575b5050508054801561580d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906159608282615704565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6159bf6159af61585e9386615704565b90549060031b1c92839286615704565b905560005283602052604060002055388080615928565b50505050600090565b80600052600660205260406000205415600014615a39576005546801000000000000000081101561412a57615a2061585e8260018594016005556005615704565b9055600554906000526006602052604060002055600190565b50600090565b600082815260018201602052604090205461541b578054906801000000000000000082101561412a5782615a7d61585e846001809601855584615704565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615d48575b615d42576fffffffffffffffffffffffffffffffff82169160018501908154615aec63ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614cd4565b9081615ca4575b5050848110615c585750838310615b4d575050615b226fffffffffffffffffffffffffffffffff928392614cd4565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b9190915460801c928315615bec5781615b6591614cd4565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211614ce157615bb3615bb89273ffffffffffffffffffffffffffffffffffffffff9661565c565b614fcb565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff83837fd0c8d23a000000000000000000000000000000000000000000000000000000006000527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615d1857615cbf92614ed29160801c90614e60565b80841015615d135750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615af3565b615cca565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615aa7565b91929015615dcb5750815115615d64575090565b3b15615d6d5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615dde5750805190602001fd5b612136906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061424b56fea164736f6c634300081a000a",
}

var MockE2ELBTCTokenPoolABI = MockE2ELBTCTokenPoolMetaData.ABI

var MockE2ELBTCTokenPoolBin = MockE2ELBTCTokenPoolMetaData.Bin

func DeployMockE2ELBTCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, advancedPoolHooks common.Address, rmnProxy common.Address, router common.Address, destPoolData []byte) (common.Address, *types.Transaction, *MockE2ELBTCTokenPool, error) {
	parsed, err := MockE2ELBTCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockE2ELBTCTokenPoolBin), backend, token, advancedPoolHooks, rmnProxy, router, destPoolData)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockE2ELBTCTokenPool{address: address, abi: *parsed, MockE2ELBTCTokenPoolCaller: MockE2ELBTCTokenPoolCaller{contract: contract}, MockE2ELBTCTokenPoolTransactor: MockE2ELBTCTokenPoolTransactor{contract: contract}, MockE2ELBTCTokenPoolFilterer: MockE2ELBTCTokenPoolFilterer{contract: contract}}, nil
}

type MockE2ELBTCTokenPool struct {
	address common.Address
	abi     abi.ABI
	MockE2ELBTCTokenPoolCaller
	MockE2ELBTCTokenPoolTransactor
	MockE2ELBTCTokenPoolFilterer
}

type MockE2ELBTCTokenPoolCaller struct {
	contract *bind.BoundContract
}

type MockE2ELBTCTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type MockE2ELBTCTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type MockE2ELBTCTokenPoolSession struct {
	Contract     *MockE2ELBTCTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockE2ELBTCTokenPoolCallerSession struct {
	Contract *MockE2ELBTCTokenPoolCaller
	CallOpts bind.CallOpts
}

type MockE2ELBTCTokenPoolTransactorSession struct {
	Contract     *MockE2ELBTCTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type MockE2ELBTCTokenPoolRaw struct {
	Contract *MockE2ELBTCTokenPool
}

type MockE2ELBTCTokenPoolCallerRaw struct {
	Contract *MockE2ELBTCTokenPoolCaller
}

type MockE2ELBTCTokenPoolTransactorRaw struct {
	Contract *MockE2ELBTCTokenPoolTransactor
}

func NewMockE2ELBTCTokenPool(address common.Address, backend bind.ContractBackend) (*MockE2ELBTCTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(MockE2ELBTCTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockE2ELBTCTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPool{address: address, abi: abi, MockE2ELBTCTokenPoolCaller: MockE2ELBTCTokenPoolCaller{contract: contract}, MockE2ELBTCTokenPoolTransactor: MockE2ELBTCTokenPoolTransactor{contract: contract}, MockE2ELBTCTokenPoolFilterer: MockE2ELBTCTokenPoolFilterer{contract: contract}}, nil
}

func NewMockE2ELBTCTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*MockE2ELBTCTokenPoolCaller, error) {
	contract, err := bindMockE2ELBTCTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolCaller{contract: contract}, nil
}

func NewMockE2ELBTCTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*MockE2ELBTCTokenPoolTransactor, error) {
	contract, err := bindMockE2ELBTCTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolTransactor{contract: contract}, nil
}

func NewMockE2ELBTCTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*MockE2ELBTCTokenPoolFilterer, error) {
	contract, err := bindMockE2ELBTCTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolFilterer{contract: contract}, nil
}

func bindMockE2ELBTCTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockE2ELBTCTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2ELBTCTokenPool.Contract.MockE2ELBTCTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.MockE2ELBTCTokenPoolTransactor.contract.Transfer(opts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.MockE2ELBTCTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockE2ELBTCTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.contract.Transfer(opts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector, customBlockConfirmation)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64, customBlockConfirmation bool) (GetCurrentRateLimiterState,

	error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector, customBlockConfirmation)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MinBlockConfirmations = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.RateLimitAdmin = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _MockE2ELBTCTokenPool.Contract.GetDynamicConfig(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _MockE2ELBTCTokenPool.Contract.GetDynamicConfig(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _MockE2ELBTCTokenPool.Contract.GetFee(&_MockE2ELBTCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _MockE2ELBTCTokenPool.Contract.GetFee(&_MockE2ELBTCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRemotePools(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRemotePools(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRemoteToken(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRemoteToken(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRequiredCCVs", localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRequiredCCVs(&_MockE2ELBTCTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRequiredCCVs(localToken common.Address, remoteChainSelector uint64, amount *big.Int, blockConfirmationRequested uint16, extraData []byte, direction uint8) ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRequiredCCVs(&_MockE2ELBTCTokenPool.CallOpts, localToken, remoteChainSelector, amount, blockConfirmationRequested, extraData, direction)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRmnProxy(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRmnProxy(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _MockE2ELBTCTokenPool.Contract.GetSupportedChains(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _MockE2ELBTCTokenPool.Contract.GetSupportedChains(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetToken() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetToken(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetToken(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _MockE2ELBTCTokenPool.Contract.GetTokenDecimals(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _MockE2ELBTCTokenPool.Contract.GetTokenDecimals(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _MockE2ELBTCTokenPool.Contract.GetTokenTransferFeeConfig(&_MockE2ELBTCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 uint16, arg3 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _MockE2ELBTCTokenPool.Contract.GetTokenTransferFeeConfig(&_MockE2ELBTCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsRemotePool(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsRemotePool(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsSupportedChain(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsSupportedChain(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsSupportedToken(&_MockE2ELBTCTokenPool.CallOpts, token)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.IsSupportedToken(&_MockE2ELBTCTokenPool.CallOpts, token)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) Owner() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.Owner(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) Owner() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.Owner(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) SDestPoolData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "s_destPoolData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SDestPoolData() ([]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.SDestPoolData(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) SDestPoolData() ([]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.SDestPoolData(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.SupportsInterface(&_MockE2ELBTCTokenPool.CallOpts, interfaceId)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.SupportsInterface(&_MockE2ELBTCTokenPool.CallOpts, interfaceId)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) TypeAndVersion() (string, error) {
	return _MockE2ELBTCTokenPool.Contract.TypeAndVersion(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _MockE2ELBTCTokenPool.Contract.TypeAndVersion(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AcceptOwnership(&_MockE2ELBTCTokenPool.TransactOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AcceptOwnership(&_MockE2ELBTCTokenPool.TransactOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AddRemotePool(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AddRemotePool(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyChainUpdates(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyChainUpdates(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.LockOrBurn(&_MockE2ELBTCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.LockOrBurn(&_MockE2ELBTCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.LockOrBurn0(&_MockE2ELBTCTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, tokenArgs []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.LockOrBurn0(&_MockE2ELBTCTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, tokenArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ReleaseOrMint(&_MockE2ELBTCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ReleaseOrMint(&_MockE2ELBTCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ReleaseOrMint0(&_MockE2ELBTCTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ReleaseOrMint0(&_MockE2ELBTCTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RemoveRemotePool(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RemoveRemotePool(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setDynamicConfig", router, minBlockConfirmations, rateLimitAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetDynamicConfig(&_MockE2ELBTCTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetDynamicConfig(router common.Address, minBlockConfirmations uint16, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetDynamicConfig(&_MockE2ELBTCTokenPool.TransactOpts, router, minBlockConfirmations, rateLimitAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setRateLimitConfig", rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRateLimitConfig(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetRateLimitConfig(rateLimitConfigArgs []TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRateLimitConfig(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.TransferOwnership(&_MockE2ELBTCTokenPool.TransactOpts, to)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.TransferOwnership(&_MockE2ELBTCTokenPool.TransactOpts, to)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.WithdrawFeeTokens(&_MockE2ELBTCTokenPool.TransactOpts, feeTokens, recipient)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.WithdrawFeeTokens(&_MockE2ELBTCTokenPool.TransactOpts, feeTokens, recipient)
}

type MockE2ELBTCTokenPoolChainAddedIterator struct {
	Event *MockE2ELBTCTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolChainAdded)
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
		it.Event = new(MockE2ELBTCTokenPoolChainAdded)
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

func (it *MockE2ELBTCTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainAddedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolChainAddedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolChainAdded)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseChainAdded(log types.Log) (*MockE2ELBTCTokenPoolChainAdded, error) {
	event := new(MockE2ELBTCTokenPoolChainAdded)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolChainRemovedIterator struct {
	Event *MockE2ELBTCTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolChainRemoved)
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
		it.Event = new(MockE2ELBTCTokenPoolChainRemoved)
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

func (it *MockE2ELBTCTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolChainRemovedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolChainRemoved)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseChainRemoved(log types.Log) (*MockE2ELBTCTokenPoolChainRemoved, error) {
	event := new(MockE2ELBTCTokenPoolChainRemoved)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolDynamicConfigSetIterator struct {
	Event *MockE2ELBTCTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolDynamicConfigSet)
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
		it.Event = new(MockE2ELBTCTokenPoolDynamicConfigSet)
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

func (it *MockE2ELBTCTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolDynamicConfigSet struct {
	Router                common.Address
	MinBlockConfirmations uint16
	RateLimitAdmin        common.Address
	Raw                   types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolDynamicConfigSetIterator{contract: _MockE2ELBTCTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolDynamicConfigSet)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*MockE2ELBTCTokenPoolDynamicConfigSet, error) {
	event := new(MockE2ELBTCTokenPoolDynamicConfigSet)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolFeeTokenWithdrawnIterator struct {
	Event *MockE2ELBTCTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(MockE2ELBTCTokenPoolFeeTokenWithdrawn)
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

func (it *MockE2ELBTCTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolFeeTokenWithdrawn struct {
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*MockE2ELBTCTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolFeeTokenWithdrawnIterator{contract: _MockE2ELBTCTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolFeeTokenWithdrawn)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*MockE2ELBTCTokenPoolFeeTokenWithdrawn, error) {
	event := new(MockE2ELBTCTokenPoolFeeTokenWithdrawn)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator struct {
	Event *MockE2ELBTCTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(MockE2ELBTCTokenPoolInboundRateLimitConsumed)
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

func (it *MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolInboundRateLimitConsumed)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolInboundRateLimitConsumed, error) {
	event := new(MockE2ELBTCTokenPoolInboundRateLimitConsumed)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolLockedOrBurnedIterator struct {
	Event *MockE2ELBTCTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolLockedOrBurned)
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
		it.Event = new(MockE2ELBTCTokenPoolLockedOrBurned)
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

func (it *MockE2ELBTCTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolLockedOrBurnedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolLockedOrBurned)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*MockE2ELBTCTokenPoolLockedOrBurned, error) {
	event := new(MockE2ELBTCTokenPoolLockedOrBurned)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *MockE2ELBTCTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(MockE2ELBTCTokenPoolOutboundRateLimitConsumed)
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

func (it *MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolOutboundRateLimitConsumed)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolOutboundRateLimitConsumed, error) {
	event := new(MockE2ELBTCTokenPoolOutboundRateLimitConsumed)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator struct {
	Event *MockE2ELBTCTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolOwnershipTransferRequested)
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
		it.Event = new(MockE2ELBTCTokenPoolOwnershipTransferRequested)
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

func (it *MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolOwnershipTransferRequested)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*MockE2ELBTCTokenPoolOwnershipTransferRequested, error) {
	event := new(MockE2ELBTCTokenPoolOwnershipTransferRequested)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolOwnershipTransferredIterator struct {
	Event *MockE2ELBTCTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolOwnershipTransferred)
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
		it.Event = new(MockE2ELBTCTokenPoolOwnershipTransferred)
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

func (it *MockE2ELBTCTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockE2ELBTCTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolOwnershipTransferredIterator{contract: _MockE2ELBTCTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolOwnershipTransferred)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*MockE2ELBTCTokenPoolOwnershipTransferred, error) {
	event := new(MockE2ELBTCTokenPoolOwnershipTransferred)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRateLimitConfiguredIterator struct {
	Event *MockE2ELBTCTokenPoolRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRateLimitConfigured)
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
		it.Event = new(MockE2ELBTCTokenPoolRateLimitConfigured)
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

func (it *MockE2ELBTCTokenPoolRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRateLimitConfigured struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRateLimitConfiguredIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRateLimitConfiguredIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RateLimitConfigured", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRateLimitConfigured)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRateLimitConfigured(log types.Log) (*MockE2ELBTCTokenPoolRateLimitConfigured, error) {
	event := new(MockE2ELBTCTokenPoolRateLimitConfigured)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolReleasedOrMintedIterator struct {
	Event *MockE2ELBTCTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolReleasedOrMinted)
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
		it.Event = new(MockE2ELBTCTokenPoolReleasedOrMinted)
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

func (it *MockE2ELBTCTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolReleasedOrMintedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolReleasedOrMinted)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*MockE2ELBTCTokenPoolReleasedOrMinted, error) {
	event := new(MockE2ELBTCTokenPoolReleasedOrMinted)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRemotePoolAddedIterator struct {
	Event *MockE2ELBTCTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRemotePoolAdded)
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
		it.Event = new(MockE2ELBTCTokenPoolRemotePoolAdded)
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

func (it *MockE2ELBTCTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRemotePoolAddedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRemotePoolAdded)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolAdded, error) {
	event := new(MockE2ELBTCTokenPoolRemotePoolAdded)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRemotePoolRemovedIterator struct {
	Event *MockE2ELBTCTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRemotePoolRemoved)
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
		it.Event = new(MockE2ELBTCTokenPoolRemotePoolRemoved)
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

func (it *MockE2ELBTCTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRemotePoolRemovedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRemotePoolRemoved)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolRemoved, error) {
	event := new(MockE2ELBTCTokenPoolRemotePoolRemoved)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *MockE2ELBTCTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (MockE2ELBTCTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (MockE2ELBTCTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (MockE2ELBTCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0xba9213054b14c2e884f779120bb196f0735cef27140498a9d26117eeab77a117")
}

func (MockE2ELBTCTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (MockE2ELBTCTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (MockE2ELBTCTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (MockE2ELBTCTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (MockE2ELBTCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MockE2ELBTCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MockE2ELBTCTokenPoolRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x41f7c8f7cfdad9350aa495e6c54cbbf750a07ab38a9098aed1256e30dd1682bb")
}

func (MockE2ELBTCTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (MockE2ELBTCTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (MockE2ELBTCTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPool) Address() common.Address {
	return _MockE2ELBTCTokenPool.address
}

type MockE2ELBTCTokenPoolInterface interface {
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

	SDestPoolData(opts *bind.CallOpts) ([]byte, error)

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

	FilterChainAdded(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*MockE2ELBTCTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*MockE2ELBTCTokenPoolChainRemoved, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*MockE2ELBTCTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*MockE2ELBTCTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*MockE2ELBTCTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*MockE2ELBTCTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockE2ELBTCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MockE2ELBTCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockE2ELBTCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MockE2ELBTCTokenPoolOwnershipTransferred, error)

	FilterRateLimitConfigured(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRateLimitConfiguredIterator, error)

	WatchRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitConfigured, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRateLimitConfigured(log types.Log) (*MockE2ELBTCTokenPoolRateLimitConfigured, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*MockE2ELBTCTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
