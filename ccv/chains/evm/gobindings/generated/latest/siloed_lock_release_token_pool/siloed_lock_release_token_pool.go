// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package siloed_lock_release_token_pool

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

type ClientEVM2AnyMessage struct {
	Receiver     []byte
	Data         []byte
	TokenAmounts []ClientEVMTokenAmount
	FeeToken     common.Address
	ExtraArgs    []byte
}

type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

type IPoolV2TokenTransferFeeConfig struct {
	DestGasOverhead                        uint32
	DestBytesOverhead                      uint32
	DefaultBlockConfirmationFeeUSDCents    uint32
	CustomBlockConfirmationFeeUSDCents     uint32
	DefaultBlockConfirmationTransferFeeBps uint16
	CustomBlockConfirmationTransferFeeBps  uint16
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

type SiloedLockReleaseTokenPoolSiloConfigUpdate struct {
	RemoteChainSelector uint64
	Rebalancer          common.Address
}

type TokenPoolCCVConfigArg struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
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

var SiloedLockReleaseTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentCustomBlockConfirmationRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"blockConfirmationConfigured\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"struct SiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x6101208060405234610406576176ff803803809161001d82856106bc565b8339810160c0828203126104065781516001600160a01b038116929091908383036104065761004e602082016106df565b60408201516001600160401b0381116104065782019280601f85011215610406578351936001600160401b03851161040b578460051b90602082019561009760405197886106bc565b865260208087019282010192831161040657602001905b8282106106a4575050506100c4606083016106ed565b6100dc60a06100d5608086016106ed565b94016106ed565b94331561069357600180546001600160a01b0319163317905586158015610682575b8015610671575b6104d75760805260c05260405163313ce56760e01b8152602081600481895afa60009181610635575b5061060a575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526104e8575b506001600160a01b031680156104d757604051636eb1769f60e11b815230600482015260248101829052602081604481865afa9081156104cb57600091610499575b5061042e5760405191602083019263095ea7b360e01b84528260248201526000196044820152604481526101dd6064826106bc565b6000806040958651936101f088866106bc565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610421573d906001600160401b03821161040b578551610261949092610252601f8201601f1916602001856106bc565b83523d6000602085013e6108a1565b80518061038a575b50506101005251616d8d908161097282396080518181816104180152818161163301528181611b6801528181611d4201528181611e4f01528181611f2f015281816124710152818161263001528181613a6501528181613c7e01528181613cdf01528181613d8e01528181613f04015281816140c6015281816145ad015281816147dd0152818161482b01526149b0015260a0518181816147930152818161574d015281816157b701526161d9015260c051818181610f7501528181611bf70152818161250001528181613af40152613f93015260e051818181610e1901528181611c3c015281816125450152613859015261010051818181610463015281816115d80152818161278a0152818161418f015281816145e801526149550152f35b81602091810103126104065760200151801590811503610406576103af573880610269565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b91610261926060916108a1565b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d6020116104c3575b816104b4602093836106bc565b810103126104065751386101a8565b3d91506104a7565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b90602090604051906104fa83836106bc565b60008252600036813760e051156105f95760005b8251811015610575576001906001600160a01b0361052c8286610701565b51168561053882610743565b610545575b50500161050e565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388561053d565b5092905060005b81518110156105f0576001906001600160a01b0361059a8285610701565b511680156105ea57846105ac82610841565b6105ba575b50505b0161057c565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138846105b1565b506105b4565b50505038610166565b6335f4a7b360e01b60005260046000fd5b60ff1660ff821681810361061e5750610134565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610669575b81610651602093836106bc565b8101031261040657610662906106df565b903861012e565b3d9150610644565b506001600160a01b03821615610105565b506001600160a01b038416156100fe565b639b15e16f60e01b60005260046000fd5b602080916106b1846106ed565b8152019101906100ae565b601f909101601f19168101906001600160401b0382119082101761040b57604052565b519060ff8216820361040657565b51906001600160a01b038216820361040657565b80518210156107155760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156107155760005260206000200190600090565b600081815260036020526040902054801561083a57600019810181811161082457600254600019810191908211610824578181036107d3575b50505060025480156107bd576000190161079781600261072b565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61080c6107e46107f593600261072b565b90549060031b1c928392600261072b565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061077c565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461089b576002546801000000000000000081101561040b576108826107f5826001859401600255600261072b565b9055600254906000526003602052604060002055600190565b50600090565b9192901561090357508151156108b5575090565b3b156108be5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156109165750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b8381106109595750508160006044809484010152601f80199101168101030190fd5b6020828201810151604487840101528593500161093756fe60c080604052600436101561001357600080fd5b600060a05260a0513560e01c90816301ffc9a714614a6a575080630a861f2a146148d7578063181f5a771461484f57806321df0da71461480a578063240028e8146147b757806324f65ee7146147785780632c063404146146e85780632d4a148f146144d557806331238ffc1461448d57806337b19247146143405780633907753714613e83578063432a6ba314613e5b578063489a68f2146139d05780634c5ef0ed1461398b57806354c8a4f31461382757806359152aad1461377a5780635f789b401461340257806362ddd3c41461337d5780636600f92c1461328e578063698c2c66146131f05780636cfd1553146131525780636d3d1a581461312a5780636d9d216c14612d215780637437ff9f14612ced57806379ba509714612c3b5780637d54534e14612bc25780638632d5cc14612b8e5780638926f54f14612b4957806389720a6214612ade5780638da5cb5b14612ab6578063962d4020146129735780639751f8841461290e5780639a4575b914612416578063a42a7b8b146122bf578063a7cd63b714612251578063acfecf911461212d578063af0e58b91461210e578063b1c71c6514611aec578063b794658014611ab0578063bb6bb5a714611a48578063c4bffe2b1461192c578063c7230a601461176d578063ce3c75281461151f578063cf7401f3146113e6578063d966866b14610f99578063dc0bd97114610f54578063ded8d95614610e3e578063e0351e1314610e00578063e8a1da17146105bb578063eb521a4c14610374578063f1e7339914610349578063f2fde38b146102955763fa41d79c1461026c57600080fd5b3461028f5760a05160031936011261028f57602061ffff600b5416604051908152f35b60a05180fd5b3461028f57602060031936011261028f576001600160a01b036102b6614bdc565b6102be6158bb565b1633811461031d5760a051805473ffffffffffffffffffffffffffffffffffffffff1916821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360a05180f35b7fdad89dca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461028f57602060031936011261028f57602061036c610367614c49565b615652565b604051908152f35b3461028f57602060031936011261028f57600435801561058f576001600160a01b036103a160a0516152b8565b16330361055f5760a05160a051526012602052604060a0512060ff600182015460a01c1660001461054a576103d782825461529b565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f0000000000000000000000000000000000000000000000000000000000000000906104599061045381608481015b03601f198101835282614d61565b8261687e565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561028f576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a080516001600160a01b039390931660048301526024820185905251909283916044918391905af1801561053d57610524575b506040519060a0515060a051825260208201527f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf60403392a260a05180f35b60a05161053091614d61565b60a05161028f57816104e5565b6040513d60a051823e3d90fd5b506105578160105461529b565b6010556103da565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b7fa90c0d190000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461028f576105c936614e26565b9190926105d46158bb565b60a051905b828210610c3e5750505060a0519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610c3857600581901b8301358581121561028f5783016101208136031261028f576040519461064886614d45565b813567ffffffffffffffff81168103610c33578652602082013567ffffffffffffffff811161028f5782019436601f8701121561028f57853561068a816151c8565b966106986040519889614d61565b81885260208089019260051b8201019036821161028f5760208101925b828410610c04575050505060208701958652604083013567ffffffffffffffff811161028f576106e89036908501614dd7565b91604088019283526107126107003660608701614f9c565b9460608a0195865260c0369101614f9c565b956080890196875261072485516162b0565b61072e87516162b0565b83515115610bd85761074a67ffffffffffffffff8a5116616a11565b15610b9d5767ffffffffffffffff89511660a051526008602052604060a0512061086a86516001600160801b036040820151169061082e6001600160801b03602083015116915115158360806040516107a281614d45565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160801b0384161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176001830155565b61096c88516001600160801b03604082015116906109306001600160801b03602083015116915115158360806040516108a281614d45565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166001600160801b0385161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176003830155565b6004855191019080519067ffffffffffffffff8211610b855761098f8354615413565b601f8111610b46575b506020906001601f841114610adc579180916109cb9360a05192610ad1575b50506000198260011b9260031b1c19161790565b90555b60a0515b88518051821015610a075790610a016001926109fa8367ffffffffffffffff8f5116926153ff565b5190615c23565b016109d2565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610ac367ffffffffffffffff6001979694985116925193519151610a98610a6c60405196879687526101006020880152610100870190614b9b565b9360408601906001600160801b0360408092805115158552826020820151166020860152015116910152565b60a08401906001600160801b0360408092805115158552826020820151166020860152015116910152565b0390a1019392909193610616565b015190508e806109b7565b90601f198316918460a051528160a051209260a0515b818110610b2e5750908460019594939210610b15575b505050811b0190556109ce565b015160001960f88460031b161c191690558d8080610b08565b92936020600181928786015181550195019301610af2565b610b75908460a05152602060a05120601f850160051c81019160208610610b7b575b601f0160051c01906155f9565b8d610998565b9091508190610b68565b634e487b7160e01b60a051526041600452602460a051fd5b67ffffffffffffffff8951167f1d5ad3c50000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b833567ffffffffffffffff811161028f57602091610c288392833691870101614dd7565b8152019301926106b5565b600080fd5b60a05180f35b9092919367ffffffffffffffff610c5e610c5986888661528b565b61513f565b1692610c698461670f565b15610dd0578360a051526008602052610c896005604060a05120016165c5565b9260a0515b8451811015610cc8576001908660a051526008602052610cc16005604060a0512001610cba83896153ff565b51906167c2565b5001610c8e565b50939094919592508060a0515260086020526005604060a0512060a051815560a051600182015560a051600282015560a051600382015560048101610d0d8154615413565b80610d80575b50500180549060a051815581610d5c575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916105d9565b60a05152602060a05120908101905b81811015610d245760a0518155600101610d6b565b601f8111600114610d99575060a05190555b8880610d13565b610db9908260a051526001601f602060a051209201861c820191016155f9565b60a080518290525160208120918190559055610d92565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461028f5760a05160031936011261028f5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461028f57602060031936011261028f57610140610e5a614c49565b610e62615367565b50610e6b615367565b50610f52610ec1610ea1610e9c610ea6610ea1610e9c8767ffffffffffffffff16600052600c602052604060002090565b615392565b616099565b9467ffffffffffffffff16600052600d602052604060002090565b610f0b60405180946001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b3461028f5760a05160031936011261028f5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461028f57602060031936011261028f5760043567ffffffffffffffff811161028f57610fca903690600401614df5565b90610fd36158bb565b60a0516080525b8160805110610fe95760a05180f35b610ff9610c596080518484615552565b6110136110096080518585615552565b6020810190615592565b61102d6110236080518787615552565b6040810190615592565b9061104861103e6080518989615552565b6060810190615592565b6110626110586080518b8b615552565b6080810190615592565b93909461107861107336898b6151e0565b61620d565b6110866110733683856151e0565b6110946110733685876151e0565b6110a26110733687896151e0565b6040519860808a01908a821067ffffffffffffffff831117610b855767ffffffffffffffff916040526110d6368a8c6151e0565b8b526110e33684866151e0565b60208c01526110f33686886151e0565b60408c015261110336888a6151e0565b60608c015216988960a05152600e602052604060a05120815180519067ffffffffffffffff8211610b8557680100000000000000008211610b855760209083548385558084106113c7575b50018260a05152602060a0512060a0515b8381106113aa57505050506001810160208301519081519167ffffffffffffffff8311610b8557680100000000000000008311610b8557602090825484845580851061138b575b50019060a05152602060a0512060a0515b83811061136e57505050506002810160408301519081519167ffffffffffffffff8311610b8557680100000000000000008311610b8557602090825484845580851061134f575b50019060a05152602060a0512060a0515b838110611332575050505060036060910191015180519067ffffffffffffffff8211610b8557680100000000000000008211610b85576020908354838555808410611313575b50019160a05152602060a051209160a0515b8281106112f65750505050926112e594926112c96112d7937fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9998966112bb6040519b8c9b60808d5260808d0191615610565b918a830360208c0152615610565b918783036040890152615610565b918483036060860152615610565b0390a2600160805101608052610fda565b60019060206001600160a01b038451169301928186015501611267565b61132c908560a05152848460a0512091820191016155f9565b8f611255565b60019060206001600160a01b03855116940193818401550161120f565b611368908460a05152858460a0512091820191016155f9565b386111fe565b60019060206001600160a01b0385511694019381840155016111b7565b6113a4908460a05152858460a0512091820191016155f9565b386111a6565b60019060206001600160a01b03855116940193818401550161115f565b6113e0908560a05152848460a0512091820191016155f9565b3861114e565b3461028f5760e060031936011261028f576113ff614c49565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc011261028f5760405161143581614d0d565b602435801515810361028f5781526044356001600160801b038116810361028f5760208201526064356001600160801b038116810361028f5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c011261028f57604051906114aa82614d0d565b608435801515810361028f57825260a4356001600160801b038116810361028f57602083015260c4356001600160801b038116810361028f5760408301526001600160a01b03600a54163314158061150a575b61055f57610c3892615f7b565b506001600160a01b03600154163314156114fd565b3461028f57604060031936011261028f57611538614c49565b60243567ffffffffffffffff82168060a05152601260205260ff6001604060a05120015460a01c16158015611765575b61173657811561058f576001600160a01b03611583846152b8565b16330361055f5760a051526012602052604060a0512060ff600182015460a01c1660a051508060001461172e5781545b8084116116fa5750156116e5576115cb828254615154565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1561028f576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af1801561053d576116cc575b506040805167ffffffffffffffff9093168352602083019190915233917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e91819081015b0390a260a05180f35b60a0516116d891614d61565b60a05161028f578261167f565b506116f281601054615154565b6010556115ce565b83907fa17e11d50000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b6010546115b3565b7f46f5f12b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b508015611568565b3461028f57604060031936011261028f5760043567ffffffffffffffff811161028f5761179e903690600401614df5565b906117a7614c08565b6117af6158bb565b60a0515b8381106117c05760a05180f35b8060206001600160a01b036117e06117db602495898961528b565b615177565b16604051938480927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa801561053d5760a051906118f3575b6001925080611833575b50016117b3565b6118a06001600160a01b0361184c6117db858a8a61528b565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208201526001600160a01b038816602482015260448082018690528152911661189b606483614d61565b61687e565b6001600160a01b036118b66117db84898961528b565b60405192835216907f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e60206001600160a01b03871692a38561182c565b509060203d8111611925575b6119098183614d61565b6020826000928101031261192257509060019151611822565b80fd5b503d6118ff565b3461028f5760a05160031936011261028f5760a051506040516006548082528160208101600660a05152602060a051209260a0515b818110611a2f57505061197692500382614d61565b80519061199b611985836151c8565b926119936040519485614d61565b8084526151c8565b90601f1960208401920136833760a0515b81518110156119de578067ffffffffffffffff6119cb600193856153ff565b51166119d782876153ff565b52016119ac565b5050906040519182916020830190602084525180915260408301919060a0515b818110611a0c575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016119fe565b8454835260019485019486945060209093019201611961565b3461028f57602060031936011261028f5760043567ffffffffffffffff811161028f57611a79903690600401614e78565b6001600160a01b03600a541633141580611a9b575b61055f57610c389161591a565b506001600160a01b0360015416331415611a8e565b3461028f57602060031936011261028f57611ae8611ad4611acf614c49565b615530565b604051918291602083526020830190614b9b565b0390f35b3461028f57606060031936011261028f5760043567ffffffffffffffff811161028f5760a0600319823603011261028f57611b25614c71565b9060443567ffffffffffffffff811161028f57611b46903690600401614dd7565b50611b4f6153e6565b5060848101611b5d81615177565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036120cd5750602481019077ffffffffffffffff00000000000000000000000000000000611bb78361513f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561053d5760a0519161209e575b5061207257611c3a60448201615177565b7f000000000000000000000000000000000000000000000000000000000000000061201f575b5067ffffffffffffffff611c738361513f565b16611c8b816000526007602052604060002054151590565b15611ff05760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561053d5760a05190611fa7575b6001600160a01b039150163303611f7757606481013561ffff84168015611ecc5761ffff600b54169081611de4575b505050611acf611d2a611dda94611da9935b60040161610c565b92611d348161513f565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a261513f565b90611db26161d2565b60405192611dbf84614d29565b83526020830152604051928392604084526040840190614f5e565b9060208301520390f35b818110611e9a575050611d2a611dda94611da993611acf937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611e2f8961513f565b16918260a05152600c60205280611e77604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616ac0565b604080516001600160a01b039290921682526020820192909252a2935094611d10565b7f7911d95b0000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b50611d2a611dda94611da993611acf937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611f0f8961513f565b16918260a05152600860205280611f57604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616ac0565b604080516001600160a01b039290921682526020820192909252a2611d22565b7f728fe07b0000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506020813d602011611fe8575b81611fc160209383614d61565b8101031261028f57516001600160a01b038116810361028f576001600160a01b0390611ce1565b3d9150611fb4565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b6001600160a01b031661203f816000526003602052604060002054151590565b611c60577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b6120c0915060203d6020116120c6575b6120b88183614d61565b8101906158a3565b84611c29565b503d6120ae565b6120de6001600160a01b0391615177565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b3461028f5760a05160031936011261028f576020601054604051908152f35b3461028f5767ffffffffffffffff61214436614ea9565b92909161214f6158bb565b1690612168826000526007602052604060002054151590565b15612221578160a05152600860205261219b6005604060a051200161218e368685614da0565b60208151910120906167c2565b156121da577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691926116c360405192839260208452602084019161550f565b61221d906040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161550f565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461028f5760a05160031936011261028f5760a051506040516002548082526020820190600260a05152602060a051209060a0515b8181106122a957611ae88561229d81870382614d61565b60405191829182614eea565b8254845260209093019260019283019201612286565b3461028f57602060031936011261028f5767ffffffffffffffff6122e1614c49565b1660a0515260086020526122fc6005604060a05120016165c5565b805190601f1961232461230e846151c8565b9361231c6040519586614d61565b8085526151c8565b0160a0515b81811061240557505060a0515b8151811015612380578061234c600192846153ff565b5160a051526009602052612364604060a0512061544d565b61236e82866153ff565b5261237981856153ff565b5001612336565b826040518091602082016020835281518091526040830190602060408260051b86010193019160a051905b8282106123ba57505050500390f35b919360206123f5827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614b9b565b96019201920185949391926123ab565b806060602080938701015201612329565b3461028f57602060031936011261028f5760043567ffffffffffffffff811161028f5760a0600319823603011261028f5761244f6153e6565b506124586153e6565b506084810161246681615177565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036120cd5750602481019077ffffffffffffffff000000000000000000000000000000006124c08361513f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561053d5760a051916128ef575b506120725761254360448201615177565b7f000000000000000000000000000000000000000000000000000000000000000061289c575b5067ffffffffffffffff61257c8361513f565b16612594816000526007602052604060002054151590565b15611ff05760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561053d5760a05190612853575b6001600160a01b039150163303611f7757606401358067ffffffffffffffff6126128461513f565b168060a051526008602052612658604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016938491616ac0565b604080516001600160a01b0384168152602081018590527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a267ffffffffffffffff6126a58461513f565b604080516001600160a01b03851681523360208201529081018590529116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a26126f6611acf8461513f565b926126ff6161d2565b6040519461270c86614d29565b8552602085015267ffffffffffffffff6127258261513f565b1660a05152601260205260ff6001604060a05120015460a01c1660001461283e578067ffffffffffffffff61275c61277e9361513f565b1660a051526012602052604060a0512061277785825461529b565b905561513f565b505b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561028f576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390931660048201526024810191909152918290818060448101039160a051905af1801561053d57612825575b60405160208082528190611ae890820185614f5e565b60a05161283191614d61565b60a05161028f578161280f565b5061284b8260105461529b565b601055612780565b506020813d602011612894575b8161286d60209383614d61565b8101031261028f57516001600160a01b038116810361028f576001600160a01b03906125ea565b3d9150612860565b6001600160a01b03166128bc816000526003602052604060002054151590565b612569577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b612908915060203d6020116120c6576120b88183614d61565b83612532565b3461028f57602060031936011261028f5767ffffffffffffffff612930614c49565b612938615367565b50612941615367565b501660a051526008602052610140604060a05120610f52610ec1610ea1600261296c610ea186615392565b9401615392565b3461028f57606060031936011261028f5760043567ffffffffffffffff811161028f576129a4903690600401614df5565b9060243567ffffffffffffffff811161028f576129c5903690600401614f2d565b9060443567ffffffffffffffff811161028f576129e6903690600401614f2d565b6001600160a01b03600a541633141580612aa1575b61055f57838614801590612a97575b612a6b5760a0515b868110612a1f5760a05180f35b80612a65612a33610c596001948b8b61528b565b612a3e838989615357565b612a5f612a57612a4f86898b615357565b923690614f9c565b913690614f9c565b91615f7b565b01612a12565b7f568efce20000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b5080861415612a0a565b506001600160a01b03600154163314156129fb565b3461028f5760a05160031936011261028f5760206001600160a01b0360015416604051908152f35b3461028f5760c060031936011261028f57612af7614bdc565b50612b00614c32565b612b08614c60565b5060843567ffffffffffffffff811161028f57612b29903690600401614c91565b505060a43590600282101561028f57611ae89161229d91604435906152fa565b3461028f57602060031936011261028f576020612b8467ffffffffffffffff612b70614c49565b166000526007602052604060002054151590565b6040519015158152f35b3461028f57602060031936011261028f576020612bb1612bac614c49565b6152b8565b6001600160a01b0360405191168152f35b3461028f57602060031936011261028f577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b03612c06614bdc565b612c0e6158bb565b168073ffffffffffffffffffffffffffffffffffffffff19600a541617600a55604051908152a160a05180f35b3461028f5760a05160031936011261028f5760a051546001600160a01b0381163303612cc15773ffffffffffffffffffffffffffffffffffffffff196001549133828416176001551660a051556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060a05160a051a360a05180f35b7f02b543c60000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461028f5760a05160031936011261028f57600454600554604080516001600160a01b039093168352602083019190915290f35b3461028f57604060031936011261028f5760043567ffffffffffffffff811161028f57612d52903690600401614df5565b6024359167ffffffffffffffff831161028f573660238401121561028f5782600401359167ffffffffffffffff831161028f576024840193602436918560061b01011161028f57612da16158bb565b60a0515b818110612ff15750505060a0515b818110612dc05760a05180f35b67ffffffffffffffff612dd7610c598385876152a8565b16158015612fba575b8015612f99575b612f52576001600160a01b03612e096020612e038486886152a8565b01615177565b1615610bd85780612ee7612e256020612e0360019587896152a8565b856001600160a01b038086604051612e3c81614d0d565b60a051815282602082019616865267ffffffffffffffff612e68610c598a8d6040860199878b526152a8565b1660a051526012602052604060a0512090518155019351161673ffffffffffffffffffffffffffffffffffffffff198354161782555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb612f16610c598386886152a8565b612f266020612e0385888a6152a8565b6040805167ffffffffffffffff9390931683526001600160a01b0391909116602083015290a101612db3565b610c5990612f699267ffffffffffffffff946152a8565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b50612fb467ffffffffffffffff612b70610c598486886152a8565b15612de7565b5067ffffffffffffffff612fd2610c598385876152a8565b1660a05152601260205260ff6001604060a05120015460a01c16612de0565b67ffffffffffffffff613008610c5983858761528b565b1660a05152601260205260ff6001604060a05120015460a01c16156130e3578067ffffffffffffffff613041610c59600194868861528b565b1660a0515260126020527f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93604060a051205461307f8160105461529b565b60105567ffffffffffffffff613099610c5985888a61528b565b60a08051919092169052601260205251604081208181558501556130c1610c5984878961528b565b6040805167ffffffffffffffff9290921682526020820192909252a101612da5565b610c59906130fa9267ffffffffffffffff9461528b565b7f46f5f12b0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b3461028f5760a05160031936011261028f5760206001600160a01b03600a5416604051908152f35b3461028f57602060031936011261028f577f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b2346001600160a01b03613194614bdc565b61319c6158bb565b6131e76011549183811673ffffffffffffffffffffffffffffffffffffffff1984161760115560405193849316839092916001600160a01b0360209181604085019616845216910152565b0390a160a05180f35b3461028f57604060031936011261028f57613209614bdc565b6024356132146158bb565b6001600160a01b038216918215610bd8577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f02389273ffffffffffffffffffffffffffffffffffffffff196004541617600455816005556131e760405192839283602090939291936001600160a01b0360408201951681520152565b3461028f57604060031936011261028f576132a7614c49565b67ffffffffffffffff6132b8614c08565b916132c16158bb565b16908160a0515260126020526001604060a051200190815460ff8160a01c161561334d57825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039283169081179093556040805191909216815260208101929092527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f9190819081016116c3565b837f46f5f12b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461028f5761338b36614ea9565b6133969291926158bb565b67ffffffffffffffff82166133b8816000526007602052604060002054151590565b156133d35750610c38926133cd913691614da0565b90615c23565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461028f57604060031936011261028f5760043567ffffffffffffffff811161028f57613433903690600401614e78565b60243567ffffffffffffffff811161028f57613453903690600401614df5565b91909261345e6158bb565b60a0515b8281106134da5750505060a0515b81811061347d5760a05180f35b8067ffffffffffffffff613497610c59600194868861528b565b168060a05152600f60205260a051604060a05120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee860a05160a051a201613470565b6134e8610c59828585615234565b6134f3828585615234565b90602082019060a0830161271061ffff61350c8361525a565b16101561376e5760c084019161271061ffff6135278561525a565b1610156137325767ffffffffffffffff16938460a05152600f60205260a0516040902061355385615269565b63ffffffff1690805490604084019161356b83615269565b60201b67ffffffff000000001693606086019461358786615269565b60401b6bffffffff00000000000000001696608001966135a688615269565b60601b6fffffffff00000000000000000000000016916135c58a61525a565b60801b71ffff0000000000000000000000000000000016936135e68c61525a565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff1617171790556040519561369d9061527a565b63ffffffff1686526136ae9061527a565b63ffffffff1660208601526136c29061527a565b63ffffffff1660408501526136d69061527a565b63ffffffff1660608401526136ea90614c82565b61ffff1660808301526136fc90614c82565b61ffff1660a082015260c07f8657880088af4e012a270ce48ac455fdcd83841e09edec744816e9b77d2d73d491a2600101613462565b61ffff61373e8461525a565b7f95f3517a0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b61373e61ffff9161525a565b3461028f57604060031936011261028f5760043561ffff81169081900361028f5760243567ffffffffffffffff811161028f577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e9161381a6137e26020933690600401614e78565b906137eb6158bb565b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b5561591a565b604051908152a160a05180f35b3461028f5761384f61385761383b36614e26565b94916138489391936158bb565b36916151e0565b9236916151e0565b7f00000000000000000000000000000000000000000000000000000000000000001561395f5760a0515b82518110156138e757806001600160a01b0361389f600193866153ff565b51166138aa81616628565b6138b6575b5001613881565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1846138af565b5060a0515b8151811015610c3857806001600160a01b0361390a600193856153ff565b511680156139595761391b816169b1565b613928575b505b016138ec565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183613920565b50613922565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461028f57604060031936011261028f576139a4614c49565b60243567ffffffffffffffff811161028f576020916139ca612b84923690600401614dd7565b9061518b565b3461028f57604060031936011261028f5760043567ffffffffffffffff811161028f5780600401610100600319833603011261028f57613a0e614c71565b90604051613a1b81614cf1565b60a0519052613a4c613a42613a3d613a3660c48701856150ee565b3691614da0565b6156d9565b60648501356157b4565b9160848401613a5a81615177565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036120cd5750602484019177ffffffffffffffff00000000000000000000000000000000613ab48461513f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561053d5760a05191613e3c575b5061207257613b348361513f565b67ffffffffffffffff8116613b56816000526007602052604060002054151590565b15611ff05750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9390931691830191909152336024830152602090829060449082906001600160a01b03165afa90811561053d5760a05191613e1d575b5015611f7757613bd48361513f565b90613bea60a48701926139ca613a3685856150ee565b15613dd657505067ffffffffffffffff613cd9613cd3604460209761ffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096161515600014613d415784613c3e8861513f565b168060a05152600d8a527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158980613ca6604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616ac0565b604080516001600160a01b039290921682526020820192909252a25b0194613ccd86615177565b5061513f565b93615177565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081168252336020830152909216908201526060810185905292169180608081015b0390a280604051613d3881614cf1565b52604051908152f35b84613d4b8861513f565b168060a0515260088a527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8980613db66002604060a05120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616ac0565b604080516001600160a01b039290921682526020820192909252a2613cc2565b613de092506150ee565b61221d6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161550f565b613e36915060203d6020116120c6576120b88183614d61565b86613bc5565b613e55915060203d6020116120c6576120b88183614d61565b86613b26565b3461028f5760a05160031936011261028f5760206001600160a01b0360115416604051908152f35b3461028f57602060031936011261028f5760043567ffffffffffffffff811161028f578060040190610100600319823603011261028f57604051613ec681614cf1565b60a0519052613eeb613ee1613a3d613a3660c48501866150ee565b60648301356157b4565b9060848101613ef981615177565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036120cd5750602481019277ffffffffffffffff00000000000000000000000000000000613f538561513f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561053d5760a05191614321575b5061207257613fd38461513f565b67ffffffffffffffff8116613ff5816000526007602052604060002054151590565b15611ff05750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9390931691830191909152336024830152602090829060449082906001600160a01b03165afa90811561053d5760a05191614302575b5015611f77576140738461513f565b9061408960a48401926139ca613a3685856150ee565b15613dd657508291905067ffffffffffffffff6140a58561513f565b168060a0515260086020526140ee6002604060a05120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016948591616ac0565b604080516001600160a01b0385168152602081018690527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a267ffffffffffffffff61413b8561513f565b1660a051526012602052604060a0512060ff600182015460a01c1660a05150806000146142fa5781545b8086116142c65750156142b15761417d848254615154565b90556141888461513f565b505b6044017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166141c082615177565b813b1561028f576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b0387811660048501526024840189905293909316604483015251909283916064918391905af1801561053d57614298575b5067ffffffffffffffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc091613d2885614267613cd360209961513f565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b60a0516142a491614d61565b60a05161028f5784614229565b506142be83601054615154565b60105561418a565b85907fa17e11d50000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b601054614165565b61431b915060203d6020116120c6576120b88183614d61565b85614064565b61433a915060203d6020116120c6576120b88183614d61565b85613fc5565b3461028f5760a060031936011261028f57614359614bdc565b50614362614c32565b60443567ffffffffffffffff811161028f5760031960a0913603011261028f5761438a614c60565b506084359067ffffffffffffffff821161028f576143b567ffffffffffffffff923690600401614c91565b50506040516143c381614cbf565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a080519101521660a05152600f60205260c0604060a0512061ffff6040519161441183614cbf565b548163ffffffff82169384815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c1685528760a06080890198828c60801c168a52019960901c1689526040519a8b52511660208a0152511660408801525116606086015251166080840152511660a0820152f35b3461028f57602060031936011261028f5767ffffffffffffffff6144af614c49565b1660a051526012602052602060ff6001604060a05120015460a01c166040519015158152f35b3461028f57604060031936011261028f576144ee614c49565b60243567ffffffffffffffff82168060a05152601260205260ff6001604060a05120015460a01c161580156146e0575b61173657811561058f576001600160a01b03614539846152b8565b16330361055f5760a051526012602052604060a0512060ff600182015460a01c166000146146cb5761456c82825461529b565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f0000000000000000000000000000000000000000000000000000000000000000906145de906104538160848101610445565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561028f576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a080516001600160a01b039390931660048301526024820185905251909283916044918391905af1801561053d576146b2575b506040805167ffffffffffffffff9093168352602083019190915233917f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf91819081016116c3565b60a0516146be91614d61565b60a05161028f578261466a565b506146d88160105461529b565b60105561456f565b50801561451e565b3461028f5760c060031936011261028f57614701614bdc565b5061470a614c32565b614712614bf2565b5060843561ffff8116810361028f5760a4359067ffffffffffffffff821161028f5763ffffffff61475961ffff9260809561475284963690600401614c91565b5050614fe8565b9392959091604051968752166020860152166040840152166060820152f35b3461028f5760a05160031936011261028f57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461028f57602060031936011261028f5760206147d2614bdc565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b3461028f5760a05160031936011261028f5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461028f5760a05160031936011261028f57604051611ae890614873606082614d61565b602481527f53696c6f65644c6f636b52656c65617365546f6b656e506f6f6c20312e372e3060208201527f2d646576000000000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190614b9b565b3461028f57602060031936011261028f57600435801561058f576001600160a01b0361490460a0516152b8565b16330361055f5760a0805180526012602052805160409020600181015490911c60ff168015614a625781545b8084116116fa575015614a4d57614948828254615154565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1561028f576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af1801561053d57614a3b575b506040519060a0515060a051825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a260a05180f35b60a051614a4791614d61565b816149fc565b50614a5a81601054615154565b60105561494b565b601054614930565b3461028f57602060031936011261028f57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361028f57817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115614b71575b8115614b47575b8115614b1d575b8115614af3575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483614aec565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614ae5565b7fdc0cbd360000000000000000000000000000000000000000000000000000000081149150614ade565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150614ad7565b919082519283825260005b848110614bc7575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201614ba6565b600435906001600160a01b0382168203610c3357565b606435906001600160a01b0382168203610c3357565b602435906001600160a01b0382168203610c3357565b35906001600160a01b0382168203610c3357565b6024359067ffffffffffffffff82168203610c3357565b6004359067ffffffffffffffff82168203610c3357565b6064359061ffff82168203610c3357565b6024359061ffff82168203610c3357565b359061ffff82168203610c3357565b9181601f84011215610c335782359167ffffffffffffffff8311610c335760208381860195010111610c3357565b60c0810190811067ffffffffffffffff821117614cdb57604052565b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff821117614cdb57604052565b6060810190811067ffffffffffffffff821117614cdb57604052565b6040810190811067ffffffffffffffff821117614cdb57604052565b60a0810190811067ffffffffffffffff821117614cdb57604052565b90601f601f19910116810190811067ffffffffffffffff821117614cdb57604052565b67ffffffffffffffff8111614cdb57601f01601f191660200190565b929192614dac82614d84565b91614dba6040519384614d61565b829481845281830111610c33578281602093846000960137010152565b9080601f83011215610c3357816020614df293359101614da0565b90565b9181601f84011215610c335782359167ffffffffffffffff8311610c33576020808501948460051b010111610c3357565b6040600319820112610c335760043567ffffffffffffffff8111610c335781614e5191600401614df5565b929092916024359067ffffffffffffffff8211610c3357614e7491600401614df5565b9091565b9181601f84011215610c335782359167ffffffffffffffff8311610c335760208085019460e08502010111610c3357565b906040600319830112610c335760043567ffffffffffffffff81168103610c3357916024359067ffffffffffffffff8211610c3357614e7491600401614c91565b602060408183019282815284518094520192019060005b818110614f0e5750505090565b82516001600160a01b0316845260209384019390920191600101614f01565b9181601f84011215610c335782359167ffffffffffffffff8311610c335760208085019460608502010111610c3357565b614df2916020614f778351604084526040840190614b9b565b920151906020818403910152614b9b565b35906001600160801b0382168203610c3357565b9190826060910312610c3357604051614fb481614d0d565b80928035908115158203610c33576040614fe39181938552614fd860208201614f88565b602086015201614f88565b910152565b67ffffffffffffffff16600052600f60205260406000206040519061500c82614cbf565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c16815261ffff6080860196818460801c1688528160a088019460901c168452168061508b5750505063ffffffff808061ffff9351169451169551169351169193929190565b919550915061ffff600b5416908181106150be57505063ffffffff808061ffff9351169451169551169351169193929190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610c33570180359067ffffffffffffffff8211610c3357602001918136038313610c3357565b3567ffffffffffffffff81168103610c335790565b9190820391821161516157565b634e487b7160e01b600052601160045260246000fd5b356001600160a01b0381168103610c335790565b9067ffffffffffffffff614df292166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111614cdb5760051b60200190565b9291906151ec816151c8565b936151fa6040519586614d61565b602085838152019160051b8101928311610c3357905b82821061521c57505050565b6020809161522984614c1e565b815201910190615210565b91908110156152445760e0020190565b634e487b7160e01b600052603260045260246000fd5b3561ffff81168103610c335790565b3563ffffffff81168103610c335790565b359063ffffffff82168203610c3357565b91908110156152445760051b0190565b9190820180921161516157565b91908110156152445760061b0190565b67ffffffffffffffff16600052601260205260016040600020015460ff8160a01c166152ee57506001600160a01b036011541690565b6001600160a01b031690565b67ffffffffffffffff16600052600e60205260406000209160028110156153415760011461533057816001614df2930190615e87565b8160026003614df294019101615e87565b634e487b7160e01b600052602160045260246000fd5b9190811015615244576060020190565b6040519061537482614d45565b60006080838281528260208201528260408201528260608201520152565b9060405161539f81614d45565b60806001829460ff81546001600160801b038116865263ffffffff81861c16602087015260a01c161515604085015201546001600160801b0381166060840152811c910152565b604051906153f382614d29565b60606020838281520152565b80518210156152445760209160051b010190565b90600182811c92168015615443575b602083101461542d57565b634e487b7160e01b600052602260045260246000fd5b91607f1691615422565b906040519182600082549261546184615413565b80845293600181169081156154cf5750600114615488575b5061548692500383614d61565b565b90506000929192526020600020906000915b8183106154b35750509060206154869282010138615479565b602091935080600191548385890101520191019091849261549a565b602093506154869592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138615479565b601f8260209493601f19938186528686013760008582860101520116010190565b67ffffffffffffffff166000526008602052614df2600460406000200161544d565b91908110156152445760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610c33570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610c33570180359067ffffffffffffffff8211610c3357602001918160051b36038313610c3357565b8181029291811591840414171561516157565b818110615604575050565b600081556001016155f9565b9160209082815201919060005b81811061562a5750505090565b9091926020806001926001600160a01b0361564488614c1e565b16815201940192910161561d565b67ffffffffffffffff16615673816000526007602052604060002054151590565b156156ac5780600052601260205260ff60016040600020015460a01c1661569b575060105490565b600052601260205260406000205490565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805180156157495760200361570b578051602082810191830183900312610c3357519060ff821161570b575060ff1690565b61221d906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614b9b565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161516157565b60ff16604d811161516157600a0a90565b811561579e570490565b634e487b7160e01b600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff81169282841461589c5782841161587257906157f99161576f565b91604d60ff8416118015615857575b6158215750509061581b614df292615783565b906155e6565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061586183615783565b801561579e57600019048411615808565b61587b9161576f565b91604d60ff84161161582157505090615896614df292615783565b90615794565b5050505090565b90816020910312610c3357518015158103610c335790565b6001600160a01b036001541633036158cf57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b358015158103610c335790565b356001600160801b0381168103610c335790565b919060005b81811061592c5750509050565b615937818386615234565b6159408161513f565b67ffffffffffffffff8116615962816000526007602052604060002054151590565b15615bf65750816159e7615a57926159ed6020600197960161598c6159873683614f9c565b6162b0565b6159e76159ad8467ffffffffffffffff16600052600c602052604060002090565b91825463ffffffff8160801c16159081615be1575b81615bd2575b81615bc0575b81615bb1575b5080615ba2575b615b2a575b3690614f9c565b906163c1565b615a1c6080840191615a026159873685614f9c565b67ffffffffffffffff16600052600d602052604060002090565b92835463ffffffff8160801c16159081615b15575b81615b06575b81615af4575b81615ae5575b5080615ad6575b615a5d575b503690614f9c565b0161591f565b615a7160a06001600160801b039201615906565b845473ffffffff000000000000000000000000000000004260801b1673ffffffffffffffffffffffffffffffffffffffff19909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538615a4f565b50615ae0826158f9565b615a4a565b60ff915060a01c161538615a43565b6001600160801b038116159150615a3d565b8589015460801c159150615a37565b858901546001600160801b0316159150615a31565b6001600160801b03615b3e60408901615906565b845473ffffffff000000000000000000000000000000004260801b1673ffffffffffffffffffffffffffffffffffffffff19909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783556159e0565b50615bac816158f9565b6159db565b60ff915060a01c1615386159d4565b6001600160801b0381161591506159ce565b848c015460801c1591506159c8565b848c01546001600160801b03161591506159c2565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b90805115615e095767ffffffffffffffff81516020830120921691826000526008602052615c58816005604060002001616a6b565b15615dc55760005260096020526040600020815167ffffffffffffffff8111614cdb57615c858254615413565b601f8111615d93575b506020601f8211600114615d095791615ce3827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593615cf995600091615cfe575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190614b9b565b0390a2565b905084015138615cd0565b601f1982169083600052806000209160005b818110615d7b575092615cf99492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615d62575b5050811b019055611ad4565b85015160001960f88460031b161c191690553880615d56565b9192602060018192868a015181550194019201615d1b565b615dbf90836000526020600020601f840160051c81019160208510610b7b57601f0160051c01906155f9565b38615c8e565b509061221d6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614b9b565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b818110615e6557505061548692500383614d61565b84546001600160a01b0316835260019485019487945060209093019201615e50565b615e9090615e33565b916005548015159182615f70575b5050615ea8575090565b615eb190615e33565b90815180615ebf5750905090565b615eca90825161529b565b92601f19615ef0615eda866151c8565b95615ee86040519788614d61565b8087526151c8565b0136602086013760005b8251811015615f2b57806001600160a01b03615f18600193866153ff565b5116615f2482886153ff565b5201615efa565b509160005b8151811015615f6b57806001600160a01b03615f4e600193856153ff565b5116615f64615f5e83875161529b565b886153ff565b5201615f30565b505050565b101590503880615e9e565b67ffffffffffffffff16600081815260076020526040902054909291901561606b579161606860e09261603d85615fd27f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b976162b0565b846000526008602052615fe98160406000206163c1565b615ff2836162b0565b84600052600860205261600c8360026040600020016163c1565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6160a1615367565b506001600160801b036060820151166001600160801b0380835116916160ec60208501936160e66160d963ffffffff87511642615154565b85608089015116906155e6565b9061529b565b8082101561610557505b16825263ffffffff4216905290565b90506160f6565b9060a061ffff9167ffffffffffffffff6161286020860161513f565b16600052600f602052826040600020916040519261614584614cbf565b549263ffffffff8416815263ffffffff8460201c16602082015263ffffffff8460401c16604082015263ffffffff8460601c16606082015282808560801c169485608084015260901c1694859101521615156000146161cb57505b1680156161c3576127106161bc6060614df294013592836155e6565b0490615154565b506060013590565b90506161a0565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152614df2604082614d61565b805160005b81811061621e57505050565b60018101808211615161575b82811061623a5750600101616212565b6001600160a01b0361624c83866153ff565b51166001600160a01b0361626083876153ff565b51161461626f5760010161622a565b6001600160a01b0361628183866153ff565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115616335576001600160801b036040820151166001600160801b03602083015116106162db5750565b606490616333604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565bfd5b6001600160801b03604082015116158015906163ab575b6163535750565b606490616333604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565b506001600160801b03602082015116151561634c565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916164e860609280546163fe63ffffffff8260801c1642615154565b908161651e575b50506001600160801b03600181602086015116928281541680851060001461651657508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556164a58651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166001600160801b031692909217910155565b61606860405180926001600160801b0360408092805115158552826020820151166020860152015116910152565b83809161642c565b6001600160801b039161654a8392836165436001880154948286169560801c906155e6565b911661529b565b808210156165be57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff92909116929092161673ffffffffffffffffffffffffffffffffffffffff19909116174260801b73ffffffff00000000000000000000000000000000161781553880616405565b9050616554565b906040519182815491828252602082019060005260206000209260005b8181106165f757505061548692500383614d61565b84548352600194850194879450602090930192016165e2565b80548210156152445760005260206000200190600090565b600081815260036020526040902054801561670857600019810181811161516157600254906000198201918211615161578181036166b7575b50505060025480156166a1576000190161667c816002616610565b60001982549160031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6166f06166c86166d9936002616610565b90549060031b1c9283926002616610565b81939154906000199060031b92831b921b19161790565b90556000526003602052604060002055388080616661565b5050600090565b60008181526007602052604090205480156167085760001981018181116151615760065490600019820191821161516157818103616788575b50505060065480156166a15760001901616763816006616610565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b6167aa6167996166d9936006616610565b90549060031b1c9283926006616610565b90556000526007602052604060002055388080616748565b90600182019181600052826020526040600020548015156000146168755760001981018181116151615782549060001982019182116151615781810361683e575b505050805480156166a157600019019061681d8282616610565b60001982549160031b1b191690555560005260205260006040812055600190565b61685e61684e6166d99386616610565b90549060031b1c92839286616610565b905560005283602052604060002055388080616803565b50505050600090565b6001600160a01b0361690091169160409260008085519361689f8786614d61565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d156169a9573d916168e483614d84565b926168f187519485614d61565b83523d6000602085013e616cb4565b8051908161690d57505050565b60208061691e9383010191016158a3565b156169265750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091616cb4565b80600052600360205260406000205415600014616a0b5760025468010000000000000000811015614cdb576169f26166d98260018594016002556002616610565b9055600254906000526003602052604060002055600190565b50600090565b80600052600760205260406000205415600014616a0b5760065468010000000000000000811015614cdb57616a526166d98260018594016006556006616610565b9055600654906000526007602052604060002055600190565b60008281526001820160205260409020546167085780549068010000000000000000821015614cdb5782616aa96166d9846001809601855584616610565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015616cac575b616ca6576001600160801b0382169160018501908154616b0663ffffffff6001600160801b0383169360801c1642615154565b9081616c08575b5050848110616bc95750838310616b5e575050616b336001600160801b03928392615154565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91616b6d8185615154565b9260001981019080821161516157616b90616b95926001600160a01b039661529b565b615794565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611616c7c57616c23926160e69160801c906155e6565b80841015616c775750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880616b0d565b616c2e565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215616ad3565b91929015616d2f5750815115616cc8575090565b3b15616cd15790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616d425750805190602001fd5b61221d906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190614b9b56fea164736f6c634300081a000a",
}

var SiloedLockReleaseTokenPoolABI = SiloedLockReleaseTokenPoolMetaData.ABI

var SiloedLockReleaseTokenPoolBin = SiloedLockReleaseTokenPoolMetaData.Bin

func DeploySiloedLockReleaseTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address, lockBox common.Address) (common.Address, *types.Transaction, *SiloedLockReleaseTokenPool, error) {
	parsed, err := SiloedLockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SiloedLockReleaseTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router, lockBox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SiloedLockReleaseTokenPool{address: address, abi: *parsed, SiloedLockReleaseTokenPoolCaller: SiloedLockReleaseTokenPoolCaller{contract: contract}, SiloedLockReleaseTokenPoolTransactor: SiloedLockReleaseTokenPoolTransactor{contract: contract}, SiloedLockReleaseTokenPoolFilterer: SiloedLockReleaseTokenPoolFilterer{contract: contract}}, nil
}

type SiloedLockReleaseTokenPool struct {
	address common.Address
	abi     abi.ABI
	SiloedLockReleaseTokenPoolCaller
	SiloedLockReleaseTokenPoolTransactor
	SiloedLockReleaseTokenPoolFilterer
}

type SiloedLockReleaseTokenPoolCaller struct {
	contract *bind.BoundContract
}

type SiloedLockReleaseTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type SiloedLockReleaseTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type SiloedLockReleaseTokenPoolSession struct {
	Contract     *SiloedLockReleaseTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SiloedLockReleaseTokenPoolCallerSession struct {
	Contract *SiloedLockReleaseTokenPoolCaller
	CallOpts bind.CallOpts
}

type SiloedLockReleaseTokenPoolTransactorSession struct {
	Contract     *SiloedLockReleaseTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type SiloedLockReleaseTokenPoolRaw struct {
	Contract *SiloedLockReleaseTokenPool
}

type SiloedLockReleaseTokenPoolCallerRaw struct {
	Contract *SiloedLockReleaseTokenPoolCaller
}

type SiloedLockReleaseTokenPoolTransactorRaw struct {
	Contract *SiloedLockReleaseTokenPoolTransactor
}

func NewSiloedLockReleaseTokenPool(address common.Address, backend bind.ContractBackend) (*SiloedLockReleaseTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(SiloedLockReleaseTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSiloedLockReleaseTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPool{address: address, abi: abi, SiloedLockReleaseTokenPoolCaller: SiloedLockReleaseTokenPoolCaller{contract: contract}, SiloedLockReleaseTokenPoolTransactor: SiloedLockReleaseTokenPoolTransactor{contract: contract}, SiloedLockReleaseTokenPoolFilterer: SiloedLockReleaseTokenPoolFilterer{contract: contract}}, nil
}

func NewSiloedLockReleaseTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*SiloedLockReleaseTokenPoolCaller, error) {
	contract, err := bindSiloedLockReleaseTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCaller{contract: contract}, nil
}

func NewSiloedLockReleaseTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloedLockReleaseTokenPoolTransactor, error) {
	contract, err := bindSiloedLockReleaseTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolTransactor{contract: contract}, nil
}

func NewSiloedLockReleaseTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloedLockReleaseTokenPoolFilterer, error) {
	contract, err := bindSiloedLockReleaseTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolFilterer{contract: contract}, nil
}

func bindSiloedLockReleaseTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloedLockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloedLockReleaseTokenPool.Contract.SiloedLockReleaseTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SiloedLockReleaseTokenPoolTransactor.contract.Transfer(opts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SiloedLockReleaseTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloedLockReleaseTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.contract.Transfer(opts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAllowList(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAllowList(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAllowListEnabled(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAllowListEnabled(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getAvailableTokens", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetAvailableTokens(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAvailableTokens(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetAvailableTokens(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAvailableTokens(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getChainRebalancer", remoteChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetChainRebalancer(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetChainRebalancer(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetChainRebalancer(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetChainRebalancer(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getCurrentCustomBlockConfirmationRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentCustomBlockConfirmationRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetDynamicConfig(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetDynamicConfig(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

	outstruct := new(GetFee)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeUSDCents = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DestGasOverhead = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.DestBytesOverhead = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TokenFeeBps = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetFee(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetFee(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetMinBlockConfirmation(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetMinBlockConfirmation(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRateLimitAdmin(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRateLimitAdmin(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRebalancer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRebalancer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRebalancer() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRebalancer(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRebalancer() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRebalancer(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRemotePools(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRemotePools(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRemoteToken(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRemoteToken(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRequiredCCVs(&_SiloedLockReleaseTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRequiredCCVs(&_SiloedLockReleaseTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRmnProxy(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetRmnProxy(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetSupportedChains(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetSupportedChains(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetToken() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetToken(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetToken(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetTokenDecimals(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetTokenDecimals(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetTokenTransferFeeConfig(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetTokenTransferFeeConfig(&_SiloedLockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetUnsiloedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getUnsiloedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetUnsiloedLiquidity() (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetUnsiloedLiquidity(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetUnsiloedLiquidity() (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetUnsiloedLiquidity(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsRemotePool(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsRemotePool(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) IsSiloed(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "isSiloed", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) IsSiloed(remoteChainSelector uint64) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSiloed(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) IsSiloed(remoteChainSelector uint64) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSiloed(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSupportedChain(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSupportedChain(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSupportedToken(&_SiloedLockReleaseTokenPool.CallOpts, token)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.IsSupportedToken(&_SiloedLockReleaseTokenPool.CallOpts, token)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) Owner() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.Owner(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) Owner() (common.Address, error) {
	return _SiloedLockReleaseTokenPool.Contract.Owner(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.SupportsInterface(&_SiloedLockReleaseTokenPool.CallOpts, interfaceId)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SiloedLockReleaseTokenPool.Contract.SupportsInterface(&_SiloedLockReleaseTokenPool.CallOpts, interfaceId)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) TypeAndVersion() (string, error) {
	return _SiloedLockReleaseTokenPool.Contract.TypeAndVersion(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _SiloedLockReleaseTokenPool.Contract.TypeAndVersion(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.AcceptOwnership(&_SiloedLockReleaseTokenPool.TransactOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.AcceptOwnership(&_SiloedLockReleaseTokenPool.TransactOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.AddRemotePool(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.AddRemotePool(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyAllowListUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyAllowListUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyCCVConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, ccvConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyCCVConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, ccvConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyChainUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyChainUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyCustomBlockConfirmationConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_SiloedLockReleaseTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn0(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn0(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "provideLiquidity", amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ProvideLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ProvideLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "provideSiloedLiquidity", remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ProvideSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ProvideSiloedLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ProvideSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ProvideSiloedLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint0(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint0(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.RemoveRemotePool(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.RemoveRemotePool(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetChainRateLimiterConfig(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetChainRateLimiterConfig(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetChainRateLimiterConfigs(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetChainRateLimiterConfigs(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setCustomBlockConfirmationRateLimitConfig", rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetDynamicConfig(&_SiloedLockReleaseTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetDynamicConfig(&_SiloedLockReleaseTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRateLimitAdmin(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRateLimitAdmin(&_SiloedLockReleaseTokenPool.TransactOpts, rateLimitAdmin)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setRebalancer", newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRebalancer(&_SiloedLockReleaseTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetRebalancer(&_SiloedLockReleaseTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "setSiloRebalancer", remoteChainSelector, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) SetSiloRebalancer(remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetSiloRebalancer(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) SetSiloRebalancer(remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.SetSiloRebalancer(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, newRebalancer)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.TransferOwnership(&_SiloedLockReleaseTokenPool.TransactOpts, to)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.TransferOwnership(&_SiloedLockReleaseTokenPool.TransactOpts, to)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) UpdateSiloDesignations(opts *bind.TransactOpts, removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "updateSiloDesignations", removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) UpdateSiloDesignations(removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.UpdateSiloDesignations(&_SiloedLockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) UpdateSiloDesignations(removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.UpdateSiloDesignations(&_SiloedLockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawFeeTokens(&_SiloedLockReleaseTokenPool.TransactOpts, feeTokens, recipient)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawFeeTokens(&_SiloedLockReleaseTokenPool.TransactOpts, feeTokens, recipient)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) WithdrawSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "withdrawSiloedLiquidity", remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) WithdrawSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawSiloedLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) WithdrawSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawSiloedLiquidity(&_SiloedLockReleaseTokenPool.TransactOpts, remoteChainSelector, amount)
}

type SiloedLockReleaseTokenPoolAllowListAddIterator struct {
	Event *SiloedLockReleaseTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolAllowListAdd)
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
		it.Event = new(SiloedLockReleaseTokenPoolAllowListAdd)
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

func (it *SiloedLockReleaseTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolAllowListAddIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolAllowListAdd)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*SiloedLockReleaseTokenPoolAllowListAdd, error) {
	event := new(SiloedLockReleaseTokenPoolAllowListAdd)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolAllowListRemoveIterator struct {
	Event *SiloedLockReleaseTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolAllowListRemove)
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
		it.Event = new(SiloedLockReleaseTokenPoolAllowListRemove)
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

func (it *SiloedLockReleaseTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolAllowListRemoveIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolAllowListRemove)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*SiloedLockReleaseTokenPoolAllowListRemove, error) {
	event := new(SiloedLockReleaseTokenPoolAllowListRemove)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolCCVConfigUpdatedIterator struct {
	Event *SiloedLockReleaseTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolCCVConfigUpdated)
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
		it.Event = new(SiloedLockReleaseTokenPoolCCVConfigUpdated)
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

func (it *SiloedLockReleaseTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCCVConfigUpdatedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolCCVConfigUpdated)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*SiloedLockReleaseTokenPoolCCVConfigUpdated, error) {
	event := new(SiloedLockReleaseTokenPoolCCVConfigUpdated)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainAddedIterator struct {
	Event *SiloedLockReleaseTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainAdded)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainAdded)
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

func (it *SiloedLockReleaseTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainAddedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainAddedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainAdded)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainAdded(log types.Log) (*SiloedLockReleaseTokenPoolChainAdded, error) {
	event := new(SiloedLockReleaseTokenPoolChainAdded)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainConfiguredIterator struct {
	Event *SiloedLockReleaseTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainConfigured)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainConfigured)
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

func (it *SiloedLockReleaseTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainConfiguredIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainConfigured)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainConfigured(log types.Log) (*SiloedLockReleaseTokenPoolChainConfigured, error) {
	event := new(SiloedLockReleaseTokenPoolChainConfigured)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainRemovedIterator struct {
	Event *SiloedLockReleaseTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainRemoved)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainRemoved)
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

func (it *SiloedLockReleaseTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainRemovedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainRemoved)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainRemoved(log types.Log) (*SiloedLockReleaseTokenPoolChainRemoved, error) {
	event := new(SiloedLockReleaseTokenPoolChainRemoved)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainSiloedIterator struct {
	Event *SiloedLockReleaseTokenPoolChainSiloed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainSiloedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainSiloed)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainSiloed)
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

func (it *SiloedLockReleaseTokenPoolChainSiloedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainSiloedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainSiloed struct {
	RemoteChainSelector uint64
	Rebalancer          common.Address
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainSiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainSiloedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainSiloed")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainSiloedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainSiloed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainSiloed) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainSiloed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainSiloed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainSiloed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainSiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainSiloed, error) {
	event := new(SiloedLockReleaseTokenPoolChainSiloed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainSiloed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolChainUnsiloedIterator struct {
	Event *SiloedLockReleaseTokenPoolChainUnsiloed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolChainUnsiloedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolChainUnsiloed)
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
		it.Event = new(SiloedLockReleaseTokenPoolChainUnsiloed)
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

func (it *SiloedLockReleaseTokenPoolChainUnsiloedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolChainUnsiloedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolChainUnsiloed struct {
	RemoteChainSelector uint64
	AmountUnsiloed      *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainUnsiloedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ChainUnsiloed")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolChainUnsiloedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ChainUnsiloed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainUnsiloed) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ChainUnsiloed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolChainUnsiloed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainUnsiloed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseChainUnsiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainUnsiloed, error) {
	event := new(SiloedLockReleaseTokenPoolChainUnsiloed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ChainUnsiloed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolConfigChangedIterator struct {
	Event *SiloedLockReleaseTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolConfigChanged)
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
		it.Event = new(SiloedLockReleaseTokenPoolConfigChanged)
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

func (it *SiloedLockReleaseTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolConfigChangedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolConfigChanged)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseConfigChanged(log types.Log) (*SiloedLockReleaseTokenPoolConfigChanged, error) {
	event := new(SiloedLockReleaseTokenPoolConfigChanged)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdatedIterator struct {
	Event *SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated)
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
		it.Event = new(SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated)
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

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdatedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "CustomBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseCustomBlockConfirmationUpdated(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated, error) {
	event := new(SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolDynamicConfigSetIterator struct {
	Event *SiloedLockReleaseTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolDynamicConfigSet)
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
		it.Event = new(SiloedLockReleaseTokenPoolDynamicConfigSet)
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

func (it *SiloedLockReleaseTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolDynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolDynamicConfigSetIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolDynamicConfigSet)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*SiloedLockReleaseTokenPoolDynamicConfigSet, error) {
	event := new(SiloedLockReleaseTokenPoolDynamicConfigSet)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator struct {
	Event *SiloedLockReleaseTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(SiloedLockReleaseTokenPoolFeeTokenWithdrawn)
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

func (it *SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolFeeTokenWithdrawn struct {
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolFeeTokenWithdrawn)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*SiloedLockReleaseTokenPoolFeeTokenWithdrawn, error) {
	event := new(SiloedLockReleaseTokenPoolFeeTokenWithdrawn)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(SiloedLockReleaseTokenPoolInboundRateLimitConsumed)
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

func (it *SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolInboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolInboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolInboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolLiquidityAddedIterator struct {
	Event *SiloedLockReleaseTokenPoolLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolLiquidityAdded)
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
		it.Event = new(SiloedLockReleaseTokenPoolLiquidityAdded)
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

func (it *SiloedLockReleaseTokenPoolLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolLiquidityAdded struct {
	RemoteChainSelector uint64
	Provider            common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*SiloedLockReleaseTokenPoolLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolLiquidityAddedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLiquidityAdded, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolLiquidityAdded)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseLiquidityAdded(log types.Log) (*SiloedLockReleaseTokenPoolLiquidityAdded, error) {
	event := new(SiloedLockReleaseTokenPoolLiquidityAdded)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolLiquidityRemovedIterator struct {
	Event *SiloedLockReleaseTokenPoolLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolLiquidityRemoved)
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
		it.Event = new(SiloedLockReleaseTokenPoolLiquidityRemoved)
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

func (it *SiloedLockReleaseTokenPoolLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolLiquidityRemoved struct {
	RemoteChainSelector uint64
	Remover             common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, remover []common.Address) (*SiloedLockReleaseTokenPoolLiquidityRemovedIterator, error) {

	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityRemoved", removerRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolLiquidityRemovedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLiquidityRemoved, remover []common.Address) (event.Subscription, error) {

	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityRemoved", removerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolLiquidityRemoved)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseLiquidityRemoved(log types.Log) (*SiloedLockReleaseTokenPoolLiquidityRemoved, error) {
	event := new(SiloedLockReleaseTokenPoolLiquidityRemoved)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolLockedOrBurnedIterator struct {
	Event *SiloedLockReleaseTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolLockedOrBurned)
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
		it.Event = new(SiloedLockReleaseTokenPoolLockedOrBurned)
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

func (it *SiloedLockReleaseTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolLockedOrBurnedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolLockedOrBurned)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*SiloedLockReleaseTokenPoolLockedOrBurned, error) {
	event := new(SiloedLockReleaseTokenPoolLockedOrBurned)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *SiloedLockReleaseTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(SiloedLockReleaseTokenPoolOutboundRateLimitConsumed)
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

func (it *SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolOutboundRateLimitConsumed)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolOutboundRateLimitConsumed, error) {
	event := new(SiloedLockReleaseTokenPoolOutboundRateLimitConsumed)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator struct {
	Event *SiloedLockReleaseTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolOwnershipTransferRequested)
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
		it.Event = new(SiloedLockReleaseTokenPoolOwnershipTransferRequested)
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

func (it *SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolOwnershipTransferRequested)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*SiloedLockReleaseTokenPoolOwnershipTransferRequested, error) {
	event := new(SiloedLockReleaseTokenPoolOwnershipTransferRequested)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolOwnershipTransferredIterator struct {
	Event *SiloedLockReleaseTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolOwnershipTransferred)
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
		it.Event = new(SiloedLockReleaseTokenPoolOwnershipTransferred)
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

func (it *SiloedLockReleaseTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedLockReleaseTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolOwnershipTransferredIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolOwnershipTransferred)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*SiloedLockReleaseTokenPoolOwnershipTransferred, error) {
	event := new(SiloedLockReleaseTokenPoolOwnershipTransferred)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolRateLimitAdminSetIterator struct {
	Event *SiloedLockReleaseTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolRateLimitAdminSet)
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
		it.Event = new(SiloedLockReleaseTokenPoolRateLimitAdminSet)
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

func (it *SiloedLockReleaseTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolRateLimitAdminSetIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolRateLimitAdminSet)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*SiloedLockReleaseTokenPoolRateLimitAdminSet, error) {
	event := new(SiloedLockReleaseTokenPoolRateLimitAdminSet)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolReleasedOrMintedIterator struct {
	Event *SiloedLockReleaseTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolReleasedOrMinted)
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
		it.Event = new(SiloedLockReleaseTokenPoolReleasedOrMinted)
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

func (it *SiloedLockReleaseTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolReleasedOrMintedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolReleasedOrMinted)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*SiloedLockReleaseTokenPoolReleasedOrMinted, error) {
	event := new(SiloedLockReleaseTokenPoolReleasedOrMinted)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolRemotePoolAddedIterator struct {
	Event *SiloedLockReleaseTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolRemotePoolAdded)
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
		it.Event = new(SiloedLockReleaseTokenPoolRemotePoolAdded)
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

func (it *SiloedLockReleaseTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolRemotePoolAddedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolRemotePoolAdded)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*SiloedLockReleaseTokenPoolRemotePoolAdded, error) {
	event := new(SiloedLockReleaseTokenPoolRemotePoolAdded)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolRemotePoolRemovedIterator struct {
	Event *SiloedLockReleaseTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolRemotePoolRemoved)
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
		it.Event = new(SiloedLockReleaseTokenPoolRemotePoolRemoved)
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

func (it *SiloedLockReleaseTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolRemotePoolRemovedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolRemotePoolRemoved)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*SiloedLockReleaseTokenPoolRemotePoolRemoved, error) {
	event := new(SiloedLockReleaseTokenPoolRemotePoolRemoved)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolSiloRebalancerSetIterator struct {
	Event *SiloedLockReleaseTokenPoolSiloRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolSiloRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolSiloRebalancerSet)
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
		it.Event = new(SiloedLockReleaseTokenPoolSiloRebalancerSet)
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

func (it *SiloedLockReleaseTokenPoolSiloRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolSiloRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolSiloRebalancerSet struct {
	RemoteChainSelector uint64
	OldRebalancer       common.Address
	NewRebalancer       common.Address
	Raw                 types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterSiloRebalancerSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolSiloRebalancerSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "SiloRebalancerSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolSiloRebalancerSetIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "SiloRebalancerSet", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchSiloRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolSiloRebalancerSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "SiloRebalancerSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolSiloRebalancerSet)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "SiloRebalancerSet", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseSiloRebalancerSet(log types.Log) (*SiloedLockReleaseTokenPoolSiloRebalancerSet, error) {
	event := new(SiloedLockReleaseTokenPoolSiloRebalancerSet)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "SiloRebalancerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator struct {
	Event *SiloedLockReleaseTokenPoolUnsiloedRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolUnsiloedRebalancerSet)
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
		it.Event = new(SiloedLockReleaseTokenPoolUnsiloedRebalancerSet)
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

func (it *SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolUnsiloedRebalancerSet struct {
	OldRebalancer common.Address
	NewRebalancer common.Address
	Raw           types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterUnsiloedRebalancerSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "UnsiloedRebalancerSet")
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "UnsiloedRebalancerSet", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchUnsiloedRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolUnsiloedRebalancerSet) (event.Subscription, error) {

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "UnsiloedRebalancerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolUnsiloedRebalancerSet)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "UnsiloedRebalancerSet", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParseUnsiloedRebalancerSet(log types.Log) (*SiloedLockReleaseTokenPoolUnsiloedRebalancerSet, error) {
	event := new(SiloedLockReleaseTokenPoolUnsiloedRebalancerSet)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "UnsiloedRebalancerSet", log); err != nil {
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
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
}

func (SiloedLockReleaseTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (SiloedLockReleaseTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (SiloedLockReleaseTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (SiloedLockReleaseTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (SiloedLockReleaseTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (SiloedLockReleaseTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (SiloedLockReleaseTokenPoolChainSiloed) Topic() common.Hash {
	return common.HexToHash("0x180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb")
}

func (SiloedLockReleaseTokenPoolChainUnsiloed) Topic() common.Hash {
	return common.HexToHash("0x7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93")
}

func (SiloedLockReleaseTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e")
}

func (SiloedLockReleaseTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (SiloedLockReleaseTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
}

func (SiloedLockReleaseTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (SiloedLockReleaseTokenPoolLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0x569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf")
}

func (SiloedLockReleaseTokenPoolLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0x58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e")
}

func (SiloedLockReleaseTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (SiloedLockReleaseTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (SiloedLockReleaseTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SiloedLockReleaseTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (SiloedLockReleaseTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (SiloedLockReleaseTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (SiloedLockReleaseTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (SiloedLockReleaseTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (SiloedLockReleaseTokenPoolSiloRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f")
}

func (SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x8657880088af4e012a270ce48ac455fdcd83841e09edec744816e9b77d2d73d4")
}

func (SiloedLockReleaseTokenPoolUnsiloedRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b234")
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPool) Address() common.Address {
	return _SiloedLockReleaseTokenPool.address
}

type SiloedLockReleaseTokenPoolInterface interface {
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

		error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

	GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error)

	GetUnsiloedLiquidity(opts *bind.CallOpts) (*big.Int, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSiloed(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error)

	SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateSiloDesignations(opts *bind.TransactOpts, removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	WithdrawSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*SiloedLockReleaseTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*SiloedLockReleaseTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*SiloedLockReleaseTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*SiloedLockReleaseTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*SiloedLockReleaseTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*SiloedLockReleaseTokenPoolChainRemoved, error)

	FilterChainSiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainSiloedIterator, error)

	WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainSiloed) (event.Subscription, error)

	ParseChainSiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainSiloed, error)

	FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolChainUnsiloedIterator, error)

	WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolChainUnsiloed) (event.Subscription, error)

	ParseChainUnsiloed(log types.Log) (*SiloedLockReleaseTokenPoolChainUnsiloed, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*SiloedLockReleaseTokenPoolConfigChanged, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdatedIterator, error)

	WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomBlockConfirmationUpdated(log types.Log) (*SiloedLockReleaseTokenPoolCustomBlockConfirmationUpdated, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*SiloedLockReleaseTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*SiloedLockReleaseTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*SiloedLockReleaseTokenPoolFeeTokenWithdrawn, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolInboundRateLimitConsumed, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*SiloedLockReleaseTokenPoolLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLiquidityAdded, provider []common.Address) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*SiloedLockReleaseTokenPoolLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, remover []common.Address) (*SiloedLockReleaseTokenPoolLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLiquidityRemoved, remover []common.Address) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*SiloedLockReleaseTokenPoolLiquidityRemoved, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*SiloedLockReleaseTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*SiloedLockReleaseTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedLockReleaseTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SiloedLockReleaseTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedLockReleaseTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SiloedLockReleaseTokenPoolOwnershipTransferred, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*SiloedLockReleaseTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*SiloedLockReleaseTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*SiloedLockReleaseTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*SiloedLockReleaseTokenPoolRemotePoolRemoved, error)

	FilterSiloRebalancerSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedLockReleaseTokenPoolSiloRebalancerSetIterator, error)

	WatchSiloRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolSiloRebalancerSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseSiloRebalancerSet(log types.Log) (*SiloedLockReleaseTokenPoolSiloRebalancerSet, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*SiloedLockReleaseTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*SiloedLockReleaseTokenPoolTokenTransferFeeConfigUpdated, error)

	FilterUnsiloedRebalancerSet(opts *bind.FilterOpts) (*SiloedLockReleaseTokenPoolUnsiloedRebalancerSetIterator, error)

	WatchUnsiloedRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolUnsiloedRebalancerSet) (event.Subscription, error)

	ParseUnsiloedRebalancerSet(log types.Log) (*SiloedLockReleaseTokenPoolUnsiloedRebalancerSet, error)

	Address() common.Address
}
