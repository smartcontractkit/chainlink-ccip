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

var CCTPTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCTPVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentCustomBlockConfirmationRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"blockConfirmationConfigured\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationRateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultFinalityRateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InboundImplementationNotFoundForVerifier\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidCCTPVerifier\",\"inputs\":[{\"name\":\"cctpVerifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610120604052346101b8576159d2803803809161001b826101d3565b6101203960c0816101200191126101b8576101205161003981610222565b610044610140610233565b610160516001600160401b0381116101b8578361013f820112156101b8578061012001519361007285610241565b9161008060405193846101ff565b858352610140602084019660051b8201019182116101b85761014001945b81861061019e5750506100d593506100b7610180610258565b906100c36101a0610258565b926100cf6101c0610258565b94610265565b60405161511390816108bf82396080518181816118ee01528181611a7e01528181611bab01528181611c6d015281816120a80152818161224f01528181612bab01528181612d6b01528181612dd501528181612e6e01528181612fc80152818161315e0152818161348b01526134d8015260a05181613441015260c051818181610b49015281816119570152818161211101528181612c140152613031015260e0518181816109f00152818161199b0152818161215401526129e5015261010051816129050152f35b6020809187516101ad81610222565b81520195019461009e565b600080fd5b634e487b7160e01b600052604160045260246000fd5b610120601f91909101601f19168101906001600160401b038211908210176101fa57604052565b6101bd565b601f909101601f19168101906001600160401b038211908210176101fa57604052565b6001600160a01b038116036101b857565b519060ff821682036101b857565b6001600160401b0381116101fa5760051b60200190565b519061026382610222565b565b9392909193610272610400565b6001600160a01b038116801580156103bc575b80156103ab575b61039a5760049260209260805260c0526040519283809263313ce56760e01b82525afa60009181610369575b5061033e575b5060a052600480546001600160a01b0319166001600160a01b039092169190911790558051151560e0819052610327575b506103006102fc8261042b565b1590565b61030a5761010052565b630d9758df60e01b6000526001600160a01b031660045260246000fd5b610338906103336103e4565b610547565b386102ef565b60ff811660ff831603156102be576332ad3e0760e11b60005260ff9182166004521660245260446000fd5b61038c91925060203d602011610393575b61038481836101ff565b8101906103cd565b90386102b8565b503d61037a565b630a64406560e11b60005260046000fd5b506001600160a01b0383161561028c565b506001600160a01b03851615610285565b908160209103126101b8576103e190610233565b90565b604051906103f36020836101ff565b6000808352366020840137565b331561041a57600180546001600160a01b03191633179055565b639b15e16f60e01b60005260046000fd5b60206000604051828101906301ffc9a760e01b82526301ffc9a760e01b60248201526024815261045c6044826101ff565b519084617530fa903d600051908361050c575b5082610502575b5081610499575b81610486575090565b6103e19150635627ff7160e01b9061069b565b905060206000604051828101906301ffc9a760e01b825263ffffffff60e01b6024820152602481526104cc6044826101ff565b519084617530fa6000513d826104f6575b50816104ec575b50159061047d565b90501515386104e4565b602011159150386104dd565b1515915038610476565b6020111592503861046f565b634e487b7160e01b600052603260045260246000fd5b80518210156105425760209160051b010190565b610518565b91906105576102fc60e051151590565b61068a5760005b83518110156105ef57806105846105776001938761052e565b516001600160a01b031690565b6105a66105a16001600160a01b0383165b6001600160a01b031690565b610796565b6105b2575b500161055e565b6040516001600160a01b039190911681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756690602090a1386105ab565b50915060005b8251811015610685578061060e6105776001938661052e565b828060a01b0381161561067f576106356106306001600160a01b038316610595565b61083a565b610642575b505b016105f5565b6040516001600160a01b039190911681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d890602090a13861063a565b5061063c565b509050565b6335f4a7b360e01b60005260046000fd5b600090602092604051848101916301ffc9a760e01b835263ffffffff60e01b166024820152602481526106cf6044826101ff565b5191617530fa6000513d826106f0575b50816106e9575090565b9050151590565b602011159150386106df565b60001981019190821161070b57565b634e487b7160e01b600052601160045260246000fd5b80548210156105425760005260206000200190600090565b916107539183549060031b91821b91600019901b19161790565b9055565b8054801561078057600019019061076e8282610721565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152600360205260409020549081156108335760001982019082821161070b576000926107f2926107cb6002546106fc565b908181036107f8575b5050506107e16002610757565b600390600052602052604060002090565b55600190565b6107e16108249161081a61081061082a956002610721565b90549060031b1c90565b9283916002610721565b90610739565b553880806107d4565b5050600090565b806000526003602052604060002054156000146108b857600254680100000000000000008110156101fa5760018101600255600060025482101561054257600290527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01819055600254906000526003602052604060002055600190565b5060009056fe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461355b57508063181f5a77146134fc57806321df0da7146134b8578063240028e81461346557806324f65ee7146134275780632c0634041461339057806337b19247146132345780633907753714612f55578063489a68f214612b3e5780634c5ef0ed14612afa57806354c8a4f3146129b357806359152aad14612929578063615521a7146128e557806362ddd3c41461287c578063698c2c66146127d55780636d3d1a58146127ae5780637437ff9f1461277b57806379ba5097146126f35780637d54534e1461268b5780638926f54f1461264757806389720a62146125dd5780638da5cb5b146125b6578063962d4020146124965780639751f884146124345780639a4575b914612055578063a42a7b8b14611f21578063a7cd63b714611ebb578063acfecf9114611dc9578063b1c71c6514611874578063b794658014611838578063bb6bb5a7146117d1578063c4bffe2b146116c1578063c7230a6014611417578063cf7401f314611305578063d966866b14610ecf578063dc04fa1f14610b6d578063dc0bd97114610b29578063ded8d95614610a15578063e0351e13146109d8578063e8a1da171461029b578063f2fde38b146102145763fa41d79c146101ed57600080fd5b3461020f57600036600319011261020f57602061ffff600b5416604051908152f35b600080fd5b3461020f57602036600319011261020f576001600160a01b036102356136e2565b61023d613f9e565b1633811461028a57806001600160a01b031960005416176000556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b636d6c4ee560e11b60005260046000fd5b3461020f576102a936613831565b9190926102b4613f9e565b6000905b828210610849575050503682900361011e1901904263ffffffff169060005b81811015610847576000918160051b860135858112156108435786019061012082360312610843576040519561030c8761362f565b82356001600160401b038116810361083f57875260208301356001600160401b03811161083f5783019536601f8801121561083f5786359661034d88613b64565b9761035b604051998a613665565b8089526020808a019160051b8301019036821161083b5760208301905b82821061080957505050506020880196875260408401356001600160401b038111610805576103aa90369086016137e3565b92604089019384526103d46103c236606088016139ae565b9560608b0196875260c03691016139ae565b9660808a019788526103e686516148e1565b6103f088516148e1565b845151156107f65761040b6001600160401b038b5116614e48565b156107d9576001600160401b038a511681526008602052604081206104ef87516001600160801b03604082015116906104c26001600160801b036020830151169151151583608060405161045e8161362f565b858152602081018c905260408101849052606081018690520152855460ff60a01b91151560a01b9190911674ffffffffffffffffffffffffffffffffffffffffff199091166001600160801b0384161763ffffffff60801b60808b901b1617178555565b60809190911b6fffffffffffffffffffffffffffffffff19166001600160801b0391909116176001830155565b6105bf89516001600160801b03604082015116906105926001600160801b03602083015116915115158360806040516105278161362f565b858152602081018c90526040810184905260608101869052015260028601805460ff60a01b92151560a01b9290921674ffffffffffffffffffffffffffffffffffffffffff199092166001600160801b0385161763ffffffff60801b60808c901b1617919091179055565b60809190911b6fffffffffffffffffffffffffffffffff19166001600160801b0391909116176003830155565b600486519101908051906001600160401b0382116107c5576105e18354613d0d565b601f811161078a575b50602090601f83116001146107275761061b929185918361071c575b50508160011b916000199060031b1c19161790565b90555b88518051821015610652579061064c600192610645836001600160401b038f511692613cf9565b51906142df565b0161061e565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561070d6001600160401b0360019796949851169251935191516106e26106b6604051968796875261010060208801526101008701906136a1565b9360408601906001600160801b0360408092805115158552826020820151166020860152015116910152565b60a08401906001600160801b0360408092805115158552826020820151166020860152015116910152565b0390a1019391939290926102d7565b015190508f80610606565b8385528185209190601f198416865b8181106107725750908460019594939210610759575b505050811b01905561061e565b015160001960f88460031b161c191690558e808061074c565b92936020600181928786015181550195019301610736565b6107b59084865260208620601f850160051c810191602086106107bb575b601f0160051c0190613ead565b8e6105ea565b90915081906107a8565b634e487b7160e01b84526041600452602484fd5b8951631d5ad3c560e01b82526001600160401b0316600452602490fd5b630a64406560e11b8152600490fd5b8680fd5b81356001600160401b0381116108375760209161082c83928336918901016137e3565b815201910190610378565b8a80fd5b8880fd5b8580fd5b8380fd5b005b90926001600160401b03610869610864868686999799613c2b565b613b00565b169261087484614c7c565b156109c3578360005260086020526108926005604060002001614b31565b9260005b84518110156108ce576001908660005260086020526108c760056040600020016108c08389613cf9565b5190614d30565b5001610896565b509390949195925080600052600860205260056040600020600081556000600182015560006002820155600060038201556004810161090d8154613d0d565b9081610980575b505001805490600081558161095f575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190919392936102b8565b6000526020600020908101905b81811015610924576000815560010161096c565b81601f600093116001146109985750555b8880610914565b818352602083206109b391601f01861c810190600101613ead565b8082528160208120915555610991565b83631e670e4b60e01b60005260045260246000fd5b3461020f57600036600319011261020f5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461020f57602036600319011261020f57610140610a31613738565b610a39613c61565b50610a42613c61565b50610b27610a96610a77610a72610a7c610a77610a72876001600160401b0316600052600c602052604060002090565b613c8c565b614712565b946001600160401b0316600052600d602052604060002090565b610ae060405180946001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b3461020f57600036600319011261020f5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020f57604036600319011261020f576004356001600160401b03811161020f573660238201121561020f5780600401356001600160401b03811161020f576024820191602436918360081b01011161020f576024356001600160401b03811161020f57610be0903690600401613801565b919092610beb613f9e565b60005b828110610c595750505060005b818110610c0457005b806001600160401b03610c1d6108646001948688613c2b565b1680600052600f602052600060408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8600080a201610bfb565b610c67610864828585613f06565b610c72828585613f06565b90602082019060e0830190610c8682613f16565b15610e845760a0840161271061ffff610c9e83613f23565b161015610ec35760c085019161271061ffff610cb985613f23565b161015610ea25763ffffffff610cce86613f32565b1615610e84576001600160401b03169485600052600f602052604060002090610cf686613f32565b63ffffffff168254926040830191610d0d83613f32565b60201b67ffffffff0000000016906060850194610d2986613f32565b60401b6bffffffff0000000000000000169060800196610d4888613f32565b60601b63ffffffff60601b1691610d5e8a613f23565b60801b61ffff60801b1693610d728c613f23565b60901b61ffff60901b169561ffff60901b199361ffff60801b199263ffffffff60601b19916bffffffffffffffffffffffff19161716171617161717178155610dba87613f16565b815460ff60a01b191690151560a01b60ff60a01b1617905560405196610ddf90613f43565b63ffffffff168752610df090613f43565b63ffffffff166020870152610e0490613f43565b63ffffffff166040860152610e1890613f43565b63ffffffff166060850152610e2c90613770565b61ffff166080840152610e3e90613770565b61ffff1660a0830152610e509061398d565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610bee565b6001600160401b0390631233226560e01b6000521660045260246000fd5b61ffff610eae84613f23565b634af9a8bd60e11b6000521660045260246000fd5b610eae61ffff91613f23565b3461020f57602036600319011261020f576004356001600160401b03811161020f57610eff903690600401613801565b610f07613f9e565b6000905b808210610f1457005b610f22610864838386613e2d565b90610f3b610f31848387613e2d565b6020810190613e4f565b9490610f55610f4b868585613e2d565b6040810190613e4f565b9690610f6f610f65888787613e2d565b6060810190613e4f565b90610f88610f7e8a8989613e2d565b6080810190613e4f565b929093610f9e610f9936888a613b7b565b614857565b610fac610f99368e84613b7b565b610fba610f99368486613b7b565b610fc8610f99368688613b7b565b6040519b60808d018d81106001600160401b0382111761124157604052610ff036888a613b7b565b8d528c610ffe368385613b7b565b9b602082019c8d526001600160401b03611019368789613b7b565b916040840192835261102c368a8c613b7b565b6060850152169c8d600052600e602052604060002092518051906001600160401b03821161124157600160401b82116112415760209085548387558084106112e8575b500184600052602060002060005b8381106112cb57505050506001839e9c9d9e019051908151916001600160401b03831161124157600160401b83116112415760209082548484558085106112ae575b500190600052602060002060005b838110611291575050505060029d9e92939495969798999a9b9c9d82019051908151916001600160401b03831161124157600160401b8311611241576020908254848455808510611274575b500190600052602060002060005b838110611257575050505060036060919e9c9d9e01910151908151916001600160401b03831161124157600160401b8311611241576020908254848455808510611224575b500190600052602060002060005b83811061120757505050506111fc9360019a999896936111e07fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9997946111ee946111d26040519b8c9b60808d5260808d0191613ec4565b918a830360208c0152613ec4565b918783036040890152613ec4565b918483036060860152613ec4565b0390a2019091610f0b565b60019060206001600160a01b03855116940193818401550161117a565b61123b908460005285846000209182019101613ead565b3861116c565b634e487b7160e01b600052604160045260246000fd5b60019060206001600160a01b038551169401938184015501611127565b61128b908460005285846000209182019101613ead565b38611119565b60019060206001600160a01b0385511694019381840155016110cd565b6112c5908460005285846000209182019101613ead565b386110bf565b60019060206001600160a01b03855116940193818401550161107d565b6112ff908760005284846000209182019101613ead565b3861106f565b3461020f5760e036600319011261020f5761131e613738565b606036602319011261020f576040516113368161364a565b602435801515810361020f5781526044356001600160801b038116810361020f5760208201526064356001600160801b038116810361020f576040820152606036608319011261020f576040519061138d8261364a565b608435801515810361020f57825260a4356001600160801b038116810361020f57602083015260c4356001600160801b038116810361020f5760408301526001600160a01b03600a541633141580611402575b6113ed5761084792614601565b63472511eb60e11b6000523360045260246000fd5b506001600160a01b03600154163314156113e0565b3461020f57604036600319011261020f576004356001600160401b03811161020f57611447903690600401613801565b6024356001600160a01b03811680820361020f57611463613f9e565b60005b83811061146f57005b8060206001600160a01b0361148f61148a602495898b613c2b565b613b14565b16604051938480926370a0823160e01b82523060048301525afa80156116b5578385918794600091611677575b50806114d0575b5050506001915001611466565b886115946001600160a01b036114ea61148a888a86613c2b565b60405163a9059cbb60e01b602082019081526001600160a01b03989098166024820152604480820187905281529116611524606483613665565b6000806040988951946115378b87613665565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1561166f573d9161157883613686565b926115858a519485613665565b83523d6000602085013e61506d565b8051806115ec575b50507f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e916001600160a01b036115da61148a8860019a602096613c2b565b169451908152a38491508383886114c3565b816020935083929496979850611606955001019101613f54565b15611618579083869392888a8061159c565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b60609161506d565b9394509150506020823d82116116ad575b8161169560209383613665565b810103126116aa5750908383869351896114bc565b80fd5b3d9150611688565b6040513d6000823e3d90fd5b3461020f57600036600319011261020f576040516006548082528160208101600660005260206000209260005b8181106117b857505061170392500382613665565b80519061172861171283613b64565b926117206040519485613665565b808452613b64565b602083019190601f190136833760005b815181101561176957806001600160401b0361175660019385613cf9565b51166117628287613cf9565b5201611738565b5050906040519182916020830190602084525180915260408301919060005b818110611796575050500390f35b82516001600160401b0316845285945060209384019390920191600101611788565b84548352600194850194869450602090930192016116ee565b3461020f57602036600319011261020f576004356001600160401b03811161020f57611801903690600401613881565b6001600160a01b03600a541633141580611823575b6113ed5761084791614012565b506001600160a01b0360015416331415611816565b3461020f57602036600319011261020f5761187061185c611857613738565b613e0c565b6040519182916020835260208301906136a1565b0390f35b3461020f57606036600319011261020f576004356001600160401b03811161020f5760a0600319823603011261020f576118ac61375f565b906044356001600160401b03811161020f576118cc9036906004016137e3565b506118d5613ce0565b50608481016118e381613b14565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611da35750602481019067ffffffffffffffff60801b61193083613b00565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156116b557600091611d74575b50611d635761199960448201613b14565b7f0000000000000000000000000000000000000000000000000000000000000000611d2b575b506001600160401b036119d183613b00565b166119e9816000526007602052604060002054151590565b15611d175760206001600160a01b03600454169160246040518094819363a8d87a3b60e01b835260048301525afa9081156116b557600091611cca575b506001600160a01b03163303611cb557606481013561ffff84168015611c0d5761ffff600b54169081611b43575b505050611857611a6e611b3994611af1935b600401614785565b92611a7881613b00565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681523360208201529081018690526001600160401b0391909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2613b00565b90604051630c11d61f60e21b602082015260048152611b11602482613665565b60405192611b1e84613614565b83526020830152604051928392604084526040840190613963565b9060208301520390f35b818110611bf6575050611a6e611b3994611af193611857937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac209793286001600160401b03611b8d89613b00565b169182600052600c60205280611bd360406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391614eed565b604080516001600160a01b039290921682526020820192909252a2935094611a54565b637911d95b60e01b60005260045260245260446000fd5b50611a6e611b3994611af193611857937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789446001600160401b03611c4f89613b00565b169182600052600860205280611c9560406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391614eed565b604080516001600160a01b039290921682526020820192909252a2611a66565b63728fe07b60e01b6000523360045260246000fd5b6020813d602011611d0f575b81611ce360209383613665565b81010312611d0b5751906001600160a01b03821682036116aa57506001600160a01b03611a26565b5080fd5b3d9150611cd6565b6354c8163f60e11b60005260045260246000fd5b6001600160a01b0316611d4b816000526003602052604060002054151590565b6119bf576368692cbb60e11b60005260045260246000fd5b630a75a23b60e31b60005260046000fd5b611d96915060203d602011611d9c575b611d8e8183613665565b810190613f54565b84611988565b503d611d84565b611db46001600160a01b0391613b14565b63961c9a4f60e01b6000521660045260246000fd5b3461020f576001600160401b03611ddf366138b1565b929091611dea613f9e565b1690611e03826000526007602052604060002054151590565b15611ea657816000526008602052611e346005604060002001611e273686856137ac565b6020815191012090614d30565b15611e78577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611e73604051928392602084526020840191613deb565b0390a2005b611ea290604051938493631d3c8f1f60e21b85526004850152604060248501526044840191613deb565b0390fd5b50631e670e4b60e01b60005260045260246000fd5b3461020f57600036600319011261020f576040516002548082526020820190600260005260206000209060005b818110611f0b5761187085611eff81870382613665565b604051918291826138f0565b8254845260209093019260019283019201611ee8565b3461020f57602036600319011261020f576001600160401b03611f42613738565b166000526008602052611f5b6005604060002001614b31565b805190611f6782613b64565b91611f756040519384613665565b808352611f84601f1991613b64565b0160005b81811061204457505060005b8151811015611fdc5780611faa60019284613cf9565b516000526009602052611fc06040600020613d47565b611fca8286613cf9565b52611fd58185613cf9565b5001611f94565b826040518091602082016020835281518091526040830190602060408260051b8601019301916000905b82821061201557505050500390f35b919360019193955060206120348192603f198a820301865288516136a1565b9601920192018594939192612006565b806060602080938701015201611f88565b3461020f57602036600319011261020f576004356001600160401b03811161020f5760a0600319823603011261020f5761208d613ce0565b5060006084820161209d81613b14565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036124105750602482019167ffffffffffffffff60801b6120ea84613b00565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561238f5783916123f1575b506123e25761215260448201613b14565b7f00000000000000000000000000000000000000000000000000000000000000006123ac575b506001600160401b0361218a84613b00565b166121a2816000526007602052604060002054151590565b1561239a5760206001600160a01b03600454169160246040518094819363a8d87a3b60e01b835260048301525afa801561238f578390612342575b6001600160a01b03915016330361232f57611870926122e8925061185791606401357ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae106001600160401b03828161223386613b00565b1680600052600860205261227760406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016968791614eed565b604080516001600160a01b0387168152602081018490527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a2611ae96122be86613b00565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b604051630c11d61f60e21b602082015260048152612307602482613665565b6040519161231483613614565b82526020820152604051918291602083526020830190613963565b63728fe07b60e01b825233600452602482fd5b506020813d602011612387575b8161235c60209383613665565b8101031261238357516001600160a01b0381168103612383576001600160a01b03906121dd565b8280fd5b3d915061234f565b6040513d85823e3d90fd5b6354c8163f60e11b8352600452602482fd5b6001600160a01b03166123cc816000526003602052604060002054151590565b612178576368692cbb60e11b8352600452602482fd5b630a75a23b60e31b8252600482fd5b61240a915060203d602011611d9c57611d8e8183613665565b84612141565b906001600160a01b03612424602493613b14565b63961c9a4f60e01b835216600452fd5b3461020f57602036600319011261020f576001600160401b03612455613738565b61245d613c61565b50612466613c61565b501660005260086020526101406040600020610b27610a96610a77600261248f610a7786613c8c565b9401613c8c565b3461020f57606036600319011261020f576004356001600160401b03811161020f576124c6903690600401613801565b906024356001600160401b03811161020f576124e6903690600401613933565b906044356001600160401b03811161020f57612506903690600401613933565b6001600160a01b03600a5416331415806125a1575b6113ed57838614801590612597575b6125865760005b86811061253a57005b8061258061254e6108646001948b8b613c2b565b612559838989613c51565b61257a61257261256a86898b613c51565b9236906139ae565b9136906139ae565b91614601565b01612531565b632b477e7160e11b60005260046000fd5b508086141561252a565b506001600160a01b036001541633141561251b565b3461020f57600036600319011261020f5760206001600160a01b0360015416604051908152f35b3461020f5760c036600319011261020f576125f66136e2565b506125ff613722565b61260761374e565b506084356001600160401b03811161020f5761262790369060040161377f565b505060a43590600282101561020f5761187091611eff9160443590613bcf565b3461020f57602036600319011261020f5760206126816001600160401b0361266d613738565b166000526007602052604060002054151590565b6040519015158152f35b3461020f57602036600319011261020f577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b036126cf6136e2565b6126d7613f9e565b16806001600160a01b0319600a541617600a55604051908152a1005b3461020f57600036600319011261020f576000546001600160a01b038116330361276a5760015490336001600160a01b03198316176001556001600160a01b0319166000556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b63015aa1e360e11b60005260046000fd5b3461020f57600036600319011261020f57600454600554604080516001600160a01b039093168352602083019190915290f35b3461020f57600036600319011261020f5760206001600160a01b03600a5416604051908152f35b3461020f57604036600319011261020f576127ee6136e2565b6024356127f9613f9e565b6001600160a01b03821691821561286b577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238926001600160a01b031960045416176004558160055561286660405192839283602090939291936001600160a01b0360408201951681520152565b0390a1005b630a64406560e11b60005260046000fd5b3461020f5761288a366138b1565b612895929192613f9e565b6001600160401b0382166128b6816000526007602052604060002054151590565b156128d15750610847926128cb9136916137ac565b906142df565b631e670e4b60e01b60005260045260246000fd5b3461020f57600036600319011261020f5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020f57604036600319011261020f5760043561ffff811680910361020f576024356001600160401b03811161020f577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e916129aa61298f6020933690600401613881565b90612998613f9e565b8361ffff19600b541617600b55614012565b604051908152a1005b3461020f576129db6129e36129c736613831565b94916129d4939193613f9e565b3691613b7b565b923691613b7b565b7f000000000000000000000000000000000000000000000000000000000000000015612ae95760005b8251811015612a7257806001600160a01b03612a2a60019386613cf9565b5116612a3581614b94565b612a41575b5001612a0c565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612a3a565b5060005b815181101561084757806001600160a01b03612a9460019385613cf9565b51168015612ae357612aa581614ded565b612ab2575b505b01612a76565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612aaa565b50612aac565b6335f4a7b360e01b60005260046000fd5b3461020f57604036600319011261020f57612b13613738565b6024356001600160401b03811161020f57602091612b386126819236906004016137e3565b90613b28565b3461020f57604036600319011261020f576004356001600160401b03811161020f578060040190610100600319823603011261020f57612b7c61375f565b90600080604051612b8c816135f9565b5260648201359260848301612ba081613b14565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612f415750602483019467ffffffffffffffff60801b612bed87613b00565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115612ef6578491612f22575b50612f13576001600160401b03612c5a87613b00565b16612c72816000526007602052604060002054151590565b15612f015760206001600160a01b0360045416916044604051809481936383826b2b60e01b835260048301523360248301525afa908115612ef6578491612ed7575b5015612ec457612cc386613b00565b90612ce060a4860192612b38612cd98585613f6c565b36916137ac565b15612e965750506001600160401b03612dc56044612dbe6020987ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09661ffff608097161515600014612e23577f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158a80612d9b60408a612d5e88613b00565b16808752600d60205295207f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316928391614eed565b604080516001600160a01b03909216825260208201929092529081908101611ae9565b9501613b14565b936001600160a01b0360405195817f000000000000000000000000000000000000000000000000000000000000000016875233898801521660408601528560608601521692a260405190612e18826135f9565b815260405190518152f35b7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8a80612d9b600260408b612e5789613b00565b169687815260206008905220016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391614eed565b612ea09250613f6c565b611ea26040519283926324eb47e560e01b8452602060048501526024840191613deb565b63728fe07b60e01b835233600452602483fd5b612ef0915060203d602011611d9c57611d8e8183613665565b87612cb4565b6040513d86823e3d90fd5b6354c8163f60e11b8452600452602483fd5b630a75a23b60e31b8352600483fd5b612f3b915060203d602011611d9c57611d8e8183613665565b87612c44565b826001600160a01b03612424602493613b14565b3461020f57602036600319011261020f576004356001600160401b03811161020f5780600401610100600319833603011261020f576000604051612f98816135f9565b5260009081604051612fa9816135f9565b5260648301359160848401612fbd81613b14565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036124105750602484019167ffffffffffffffff60801b61300a84613b00565b60801b1660405190632cbc26bb60e01b825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561238f578391613215575b506123e2576001600160401b0361307784613b00565b1661308f816000526007602052604060002054151590565b1561239a5760206001600160a01b0360045416916044604051809481936383826b2b60e01b835260048301523360248301525afa90811561238f5783916131f6575b501561232f576130e083613b00565b906130f660a4870192612b38612cd98585613f6c565b15612e9657602085807ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc060806001600160401b038b6001600160a01b036131cf60446131c88e8e613186600260408a61314e86613b00565b16938481526020600890522001877f0000000000000000000000000000000000000000000000000000000000000000169c8d91614eed565b604080516001600160a01b038d168152602081018e90527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9181908101611ae9565b9301613b14565b60405196875233898801521660408601528560608601521692a260405190612e18826135f9565b61320f915060203d602011611d9c57611d8e8183613665565b866130d1565b61322e915060203d602011611d9c57611d8e8183613665565b86613061565b3461020f5760a036600319011261020f5761324d6136e2565b50613256613722565b6044356001600160401b03811161020f5760a090600319903603011261020f5761327e61374e565b50608435906001600160401b03821161020f576132a76001600160401b0392369060040161377f565b5050600060c06040516132b9816135de565b8281528260208201528260408201528260608201528260808201528260a0820152015216600052600f60205260e06040600020604051906132f9826135de565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b3461020f5760c036600319011261020f576133a96136e2565b506133b2613722565b6133ba6136f8565b5060843561ffff8116810361020f5760a435906001600160401b03821161020f5763ffffffff61ffff6134008293866133f960a097369060040161377f565b50506139f5565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b3461020f57600036600319011261020f57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020f57602036600319011261020f5760206134806136e2565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b3461020f57600036600319011261020f5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020f57600036600319011261020f57611870604080519061351f8183613665565b601782527f43435450546f6b656e506f6f6c20312e372e302d6465760000000000000000006020830152519182916020835260208301906136a1565b3461020f57602036600319011261020f576004359063ffffffff60e01b821680920361020f5760209163aff2afbf60e01b81149081156135cd575b81156135bc575b81156135ab575b5015158152f35b6301ffc9a760e01b149050836135a4565b630e64dd2960e01b8114915061359d565b636e065e9b60e11b81149150613596565b60e081019081106001600160401b0382111761124157604052565b602081019081106001600160401b0382111761124157604052565b604081019081106001600160401b0382111761124157604052565b60a081019081106001600160401b0382111761124157604052565b606081019081106001600160401b0382111761124157604052565b90601f801991011681019081106001600160401b0382111761124157604052565b6001600160401b03811161124157601f01601f191660200190565b919082519283825260005b8481106136cd575050826000602080949584010152601f8019910116010190565b806020809284010151828286010152016136ac565b600435906001600160a01b038216820361020f57565b606435906001600160a01b038216820361020f57565b35906001600160a01b038216820361020f57565b602435906001600160401b038216820361020f57565b600435906001600160401b038216820361020f57565b6064359061ffff8216820361020f57565b6024359061ffff8216820361020f57565b359061ffff8216820361020f57565b9181601f8401121561020f578235916001600160401b03831161020f576020838186019501011161020f57565b9291926137b882613686565b916137c66040519384613665565b82948184528183011161020f578281602093846000960137010152565b9080601f8301121561020f578160206137fe933591016137ac565b90565b9181601f8401121561020f578235916001600160401b03831161020f576020808501948460051b01011161020f57565b604060031982011261020f576004356001600160401b03811161020f578161385b91600401613801565b92909291602435906001600160401b03821161020f5761387d91600401613801565b9091565b9181601f8401121561020f578235916001600160401b03831161020f5760208085019460e0850201011161020f57565b90604060031983011261020f576004356001600160401b038116810361020f5791602435906001600160401b03821161020f5761387d9160040161377f565b602060408183019282815284518094520192019060005b8181106139145750505090565b82516001600160a01b0316845260209384019390920191600101613907565b9181601f8401121561020f578235916001600160401b03831161020f576020808501946060850201011161020f57565b6137fe91602061397c83516040845260408401906136a1565b9201519060208184039101526136a1565b3590811515820361020f57565b35906001600160801b038216820361020f57565b919082606091031261020f576040516139c68161364a565b60406139f08183956139d78161398d565b85526139e56020820161399a565b60208601520161399a565b910152565b6001600160401b0316600052600f602052604060002060405190613a18826135de565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c168152608085019561ffff8360801c16875260ff60a087019361ffff8160901c16855260a01c1615801560c0880152613ae75761ffff1680613ab15750505063ffffffff808061ffff9351169451169551169351169193929190600190565b919550915061ffff600b541690818110611bf657505063ffffffff808061ffff9351169451169551169351169193929190600190565b5050505092505050600090600090600090600090600090565b356001600160401b038116810361020f5790565b356001600160a01b038116810361020f5790565b906001600160401b036137fe92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b6001600160401b0381116112415760051b60200190565b929190613b8781613b64565b93613b956040519586613665565b602085838152019160051b810192831161020f57905b828210613bb757505050565b60208091613bc48461370e565b815201910190613bab565b6001600160401b0316600052600e6020526040600020916002811015613c1557600114613c04578160016137fe93019061450c565b81600260036137fe9401910161450c565b634e487b7160e01b600052602160045260246000fd5b9190811015613c3b5760051b0190565b634e487b7160e01b600052603260045260246000fd5b9190811015613c3b576060020190565b60405190613c6e8261362f565b60006080838281528260208201528260408201528260608201520152565b90604051613c998161362f565b60806001829460ff81546001600160801b038116865263ffffffff81861c16602087015260a01c161515604085015201546001600160801b0381166060840152811c910152565b60405190613ced82613614565b60606020838281520152565b8051821015613c3b5760209160051b010190565b90600182811c92168015613d3d575b6020831014613d2757565b634e487b7160e01b600052602260045260246000fd5b91607f1691613d1c565b9060405191826000825492613d5b84613d0d565b8084529360018116908115613dc95750600114613d82575b50613d8092500383613665565b565b90506000929192526020600020906000915b818310613dad575050906020613d809282010138613d73565b6020919350806001915483858901015201910190918492613d94565b905060209250613d8094915060ff191682840152151560051b82010138613d73565b908060209392818452848401376000828201840152601f01601f1916010190565b6001600160401b031660005260086020526137fe6004604060002001613d47565b9190811015613c3b5760051b81013590609e198136030182121561020f570190565b903590601e198136030182121561020f57018035906001600160401b03821161020f57602001918160051b3603831361020f57565b81810292918115918404141715613e9757565b634e487b7160e01b600052601160045260246000fd5b818110613eb8575050565b60008155600101613ead565b9160209082815201919060005b818110613ede5750505090565b9091926020806001926001600160a01b03613ef88861370e565b168152019401929101613ed1565b9190811015613c3b5760081b0190565b35801515810361020f5790565b3561ffff8116810361020f5790565b3563ffffffff8116810361020f5790565b359063ffffffff8216820361020f57565b9081602091031261020f5751801515810361020f5790565b903590601e198136030182121561020f57018035906001600160401b03821161020f5760200191813603831361020f57565b6001600160a01b03600154163303613fb257565b6315ae3a6f60e11b60005260046000fd5b356001600160801b038116810361020f5790565b6001600160801b0361400c60408093613fef8161398d565b15158652836140006020830161399a565b1660208701520161399a565b16910152565b9160005b828110156142d957600060e08202850161402f81613b00565b906001600160401b03821692614052846000526007602052604060002054151590565b156142c55750600193926141897f20ae59542ddd78610f62f9d2c9dcd658f8b6a5b45a0f03e71e16614e89dda8369361417f8461416f602060e09701916140a161409c36856139ae565b6148e1565b6141036140c1866001600160401b0316600052600c602052604060002090565b805463ffffffff8160801c161590816142b0575b816142a1575b8161428f575b81614280575b5080614271575b61422b575b6140fd36866139ae565b906149b8565b614131608082019561411861409c36896139ae565b6001600160401b0316600052600d602052604060002090565b90815463ffffffff8160801c16159081614216575b81614207575b816141f5575b816141e6575b50806141d7575b614190575b506140fd36866139ae565b6040519485526020850190613fd7565b6080830190613fd7565ba101614016565b6141a460a06001600160801b039201613fc3565b825463ffffffff60801b4260801b166001600160a01b03199091169190921663ffffffff60801b19161717815538614164565b506141e186613f16565b61415f565b60ff915060a01c161538614158565b6001600160801b038116159150614152565b838e015460801c15915061414c565b838e01546001600160801b0316159150614146565b6001600160801b0361423f60408501613fc3565b825463ffffffff60801b4260801b166001600160a01b03199091169190921663ffffffff60801b1916171781556140f3565b5061427b85613f16565b6140ee565b60ff915060a01c1615386140e7565b6001600160801b0381161591506140e1565b828f015460801c1591506140db565b828f01546001600160801b03161591506140d5565b631e670e4b60e01b81526004849052602490fd5b50915050565b9080511561286b576001600160401b0381516020830120921691826000526008602052614313816005604060002001614e9d565b15614480576000526009602052604060002081516001600160401b0381116112415761433f8254613d0d565b601f811161444e575b506020601f82116001146143c4579161439e827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936143b4956000916143b9575b508160011b916000199060031b1c19161790565b90556040519182916020835260208301906136a1565b0390a2565b90508401513861438a565b601f1982169083600052806000209160005b8181106144365750926143b49492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061441d575b5050811b01905561185c565b85015160001960f88460031b161c191690553880614411565b9192602060018192868a0151815501940192016143d6565b61447a90836000526020600020601f840160051c810191602085106107bb57601f0160051c0190613ead565b38614348565b5090611ea2604051928392631c9dc56960e11b845260048401526040602484015260448301906136a1565b906040519182815491828252602082019060005260206000209260005b8181106144dd575050613d8092500383613665565b84546001600160a01b03168352600194850194879450602090930192016144c8565b91908201809211613e9757565b614515906144ab565b9160055480151591826145f6575b505061452d575090565b614536906144ab565b908151806145445750905090565b61454f9082516144ff565b9261455984613b64565b936145676040519586613665565b808552614576601f1991613b64565b0136602086013760005b82518110156145b157806001600160a01b0361459e60019386613cf9565b51166145aa8288613cf9565b5201614580565b509160005b81518110156145f157806001600160a01b036145d460019385613cf9565b51166145ea6145e48387516144ff565b88613cf9565b52016145b6565b505050565b101590503880614523565b6001600160401b031660008181526007602052604090205490929190156146f057916146ed60e0926146c2856146577f73d6dce40db73cbddae4b9ce52576043a1644e08c2702236273d71077435fa16976148e1565b84600052600860205261466e8160406000206149b8565b614677836148e1565b8460005260086020526146918360026040600020016149b8565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565ba1565b82631e670e4b60e01b60005260045260246000fd5b91908203918211613e9757565b61471a613c61565b506001600160801b036060820151166001600160801b038083511691614765602085019361475f61475263ffffffff87511642614705565b8560808901511690613e84565b906144ff565b8082101561477e57505b16825263ffffffff4216905290565b905061476f565b9061ffff906001600160401b0361479e60208501613b00565b16600052600f602052604060002082604051916147ba836135de565b549263ffffffff8416835263ffffffff8460201c16602084015263ffffffff8460401c16604084015263ffffffff8460601c166060840152818460801c169283608082015260c060ff848760901c16968760a085015260a01c16151591015216151560001461485057505b1680156148485761271061484160606137fe9401359283613e84565b0490614705565b506060013590565b9050614825565b805160005b81811061486857505050565b60018101808211613e97575b828110614884575060010161485c565b6001600160a01b036148968386613cf9565b51166001600160a01b036148aa8387613cf9565b5116146148b957600101614874565b6001600160a01b036148cb8386613cf9565b5116630285c9b960e61b60005260045260246000fd5b805115614948576001600160801b036040820151166001600160801b036020830151161061490c5750565b60408051632008344960e21b815282511515600482015260208301516001600160801b0390811660248301529190920151166044820152606490fd5b6001600160801b03604082015116158015906149a2575b6149665750565b604080516335a2be7360e21b815282511515600482015260208301516001600160801b0390811660248301529190920151166044820152606490fd5b506001600160801b03602082015116151561495f565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991614a8660609280546149f563ffffffff8260801c1642614705565b9081614abc575b50506001600160801b036001816020860151169282815416808510600014614ab457508280855b16168319825416178155614a5286511515829081549060ff60a01b90151560a01b169060ff60a01b1916179055565b60408601516fffffffffffffffffffffffffffffffff1960809190911b16939092166001600160801b031692909217910155565b6146ed60405180926001600160801b0360408092805115158552826020820151166020860152015116910152565b838091614a23565b6001600160801b0391614ae8839283614ae16001880154948286169560801c90613e84565b91166144ff565b80821015614b2a57505b835463ffffffff60801b199290911692909216166001600160a01b0319909116174260801b63ffffffff60801b1617815538806149fc565b9050614af2565b906040519182815491828252602082019060005260206000209260005b818110614b63575050613d8092500383613665565b8454835260019485019487945060209093019201614b4e565b8054821015613c3b5760005260206000200190600090565b6000818152600360205260409020548015614c75576000198101818111613e9757600254600019810191908211613e9757818103614c24575b5050506002548015614c0e5760001901614be8816002614b7c565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b614c5d614c35614c46936002614b7c565b90549060031b1c9283926002614b7c565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080614bcd565b5050600090565b6000818152600760205260409020548015614c75576000198101818111613e9757600654600019810191908211613e9757818103614cf6575b5050506006548015614c0e5760001901614cd0816006614b7c565b8154906000199060031b1b19169055600655600052600760205260006040812055600190565b614d18614d07614c46936006614b7c565b90549060031b1c9283926006614b7c565b90556000526007602052604060002055388080614cb5565b9060018201918160005282602052604060002054801515600014614de4576000198101818111613e97578254600019810191908211613e9757818103614dad575b50505080548015614c0e576000190190614d8b8282614b7c565b8154906000199060031b1b191690555560005260205260006040812055600190565b614dcd614dbd614c469386614b7c565b90549060031b1c92839286614b7c565b905560005283602052604060002055388080614d71565b50505050600090565b80600052600360205260406000205415600014614e4257600254600160401b81101561124157614e29614c468260018594016002556002614b7c565b9055600254906000526003602052604060002055600190565b50600090565b80600052600760205260406000205415600014614e4257600654600160401b81101561124157614e84614c468260018594016006556006614b7c565b9055600654906000526007602052604060002055600190565b6000828152600182016020526040902054614c7557805490600160401b8210156112415782614ed6614c46846001809601855584614b7c565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615065575b61505f576001600160801b0382169160018501908154614f3363ffffffff6001600160801b0383169360801c1642614705565b9081614fff575b5050848110614fd95750838310614f73575050614f606001600160801b03928392614705565b16166001600160801b0319825416179055565b5460801c91614f828185614705565b6000198401848111613e9757614f97916144ff565b928015614fc3576001600160a01b039304636864691d60e11b6000526004526024521660445260646000fd5b634e487b7160e01b600052601260045260246000fd5b82856001600160a01b0392630d3b2b9560e11b6000526004526024521660445260646000fd5b82869293961161504e5761501a9261475f9160801c90613e84565b808410156150495750825b855463ffffffff60801b19164260801b63ffffffff60801b16178655923880614f3a565b615025565b634b92ca1560e11b60005260046000fd5b50505050565b508215614f00565b919290156150cf5750815115615081575090565b3b1561508a5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156150e25750805190602001fd5b60405162461bcd60e51b815260206004820152908190611ea29060248301906136a156fea164736f6c634300081a000a",
}

var CCTPTokenPoolABI = CCTPTokenPoolMetaData.ABI

var CCTPTokenPoolBin = CCTPTokenPoolMetaData.Bin

func DeployCCTPTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address, cctpVerifier common.Address) (common.Address, *types.Transaction, *CCTPTokenPool, error) {
	parsed, err := CCTPTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router, cctpVerifier)
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

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _CCTPTokenPool.Contract.GetAllowList(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _CCTPTokenPool.Contract.GetAllowList(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _CCTPTokenPool.Contract.GetAllowListEnabled(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _CCTPTokenPool.Contract.GetAllowListEnabled(&_CCTPTokenPool.CallOpts)
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
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

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

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetRateLimitAdmin(&_CCTPTokenPool.CallOpts)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _CCTPTokenPool.Contract.GetRateLimitAdmin(&_CCTPTokenPool.CallOpts)
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

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _CCTPTokenPool.Contract.GetRequiredCCVs(&_CCTPTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _CCTPTokenPool.Contract.GetRequiredCCVs(&_CCTPTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
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

func (_CCTPTokenPool *CCTPTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _CCTPTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_CCTPTokenPool *CCTPTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CCTPTokenPool.Contract.GetTokenTransferFeeConfig(&_CCTPTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_CCTPTokenPool *CCTPTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _CCTPTokenPool.Contract.GetTokenTransferFeeConfig(&_CCTPTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyAllowListUpdates(&_CCTPTokenPool.TransactOpts, removes, adds)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyAllowListUpdates(&_CCTPTokenPool.TransactOpts, removes, adds)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyCCVConfigUpdates(&_CCTPTokenPool.TransactOpts, ccvConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyCCVConfigUpdates(&_CCTPTokenPool.TransactOpts, ccvConfigArgs)
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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "applyCustomBlockConfirmationConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_CCTPTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_CCTPTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.LockOrBurn0(&_CCTPTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.LockOrBurn0(&_CCTPTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, arg2)
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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetChainRateLimiterConfig(&_CCTPTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetChainRateLimiterConfig(&_CCTPTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
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

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetDynamicConfig(&_CCTPTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetDynamicConfig(&_CCTPTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_CCTPTokenPool *CCTPTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetRateLimitAdmin(&_CCTPTokenPool.TransactOpts, rateLimitAdmin)
}

func (_CCTPTokenPool *CCTPTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _CCTPTokenPool.Contract.SetRateLimitAdmin(&_CCTPTokenPool.TransactOpts, rateLimitAdmin)
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

type CCTPTokenPoolAllowListAddIterator struct {
	Event *CCTPTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolAllowListAdd)
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
		it.Event = new(CCTPTokenPoolAllowListAdd)
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

func (it *CCTPTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*CCTPTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolAllowListAddIterator{contract: _CCTPTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolAllowListAdd)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*CCTPTokenPoolAllowListAdd, error) {
	event := new(CCTPTokenPoolAllowListAdd)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolAllowListRemoveIterator struct {
	Event *CCTPTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolAllowListRemove)
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
		it.Event = new(CCTPTokenPoolAllowListRemove)
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

func (it *CCTPTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*CCTPTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolAllowListRemoveIterator{contract: _CCTPTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolAllowListRemove)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*CCTPTokenPoolAllowListRemove, error) {
	event := new(CCTPTokenPoolAllowListRemove)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CCTPTokenPoolCCVConfigUpdatedIterator struct {
	Event *CCTPTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolCCVConfigUpdated)
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
		it.Event = new(CCTPTokenPoolCCVConfigUpdated)
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

func (it *CCTPTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolCCVConfigUpdatedIterator{contract: _CCTPTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolCCVConfigUpdated)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*CCTPTokenPoolCCVConfigUpdated, error) {
	event := new(CCTPTokenPoolCCVConfigUpdated)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

type CCTPTokenPoolCustomBlockConfirmationUpdatedIterator struct {
	Event *CCTPTokenPoolCustomBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolCustomBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolCustomBlockConfirmationUpdated)
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
		it.Event = new(CCTPTokenPoolCustomBlockConfirmationUpdated)
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

func (it *CCTPTokenPoolCustomBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolCustomBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolCustomBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*CCTPTokenPoolCustomBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolCustomBlockConfirmationUpdatedIterator{contract: _CCTPTokenPool.contract, event: "CustomBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolCustomBlockConfirmationUpdated)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseCustomBlockConfirmationUpdated(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationUpdated, error) {
	event := new(CCTPTokenPoolCustomBlockConfirmationUpdated)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
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

type CCTPTokenPoolRateLimitAdminSetIterator struct {
	Event *CCTPTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *CCTPTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPTokenPoolRateLimitAdminSet)
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
		it.Event = new(CCTPTokenPoolRateLimitAdminSet)
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

func (it *CCTPTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *CCTPTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type CCTPTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*CCTPTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _CCTPTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &CCTPTokenPoolRateLimitAdminSetIterator{contract: _CCTPTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_CCTPTokenPool *CCTPTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _CCTPTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(CCTPTokenPoolRateLimitAdminSet)
				if err := _CCTPTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_CCTPTokenPool *CCTPTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*CCTPTokenPoolRateLimitAdminSet, error) {
	event := new(CCTPTokenPoolRateLimitAdminSet)
	if err := _CCTPTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
	IsEnabled         bool
}

func (CCTPTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (CCTPTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (CCTPTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
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

func (CCTPTokenPoolCustomBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e")
}

func (CCTPTokenPoolDefaultFinalityRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x73d6dce40db73cbddae4b9ce52576043a1644e08c2702236273d71077435fa16")
}

func (CCTPTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
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

func (CCTPTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
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
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCCTPVerifier(opts *bind.CallOpts) (common.Address, error)

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

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

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

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*CCTPTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*CCTPTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*CCTPTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*CCTPTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*CCTPTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*CCTPTokenPoolCCVConfigUpdated, error)

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

	FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*CCTPTokenPoolCustomBlockConfirmationUpdatedIterator, error)

	WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomBlockConfirmationUpdated(log types.Log) (*CCTPTokenPoolCustomBlockConfirmationUpdated, error)

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

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*CCTPTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *CCTPTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*CCTPTokenPoolRateLimitAdminSet, error)

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
