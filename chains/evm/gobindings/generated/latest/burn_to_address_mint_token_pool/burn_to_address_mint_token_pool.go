// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_to_address_mint_token_pool

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
	DestGasOverhead               uint32
	DestBytesOverhead             uint32
	DefaultFinalityFeeUSDCents    uint32
	CustomFinalityFeeUSDCents     uint32
	DefaultFinalityTransferFeeBps uint16
	CustomFinalityTransferFeeBps  uint16
	IsEnabled                     bool
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

type TokenPoolCustomFinalityRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var BurnToAddressMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_burnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityMinimumBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customFinalityFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103dd576165ae803803809161001d828561045c565b8339810160c0828203126103dd5781516001600160a01b03811692908390036103dd5761004c6020820161047f565b60408201516001600160401b0381116103dd5782019280601f850112156103dd578351936001600160401b038511610446578460051b906020820195610095604051978861045c565b86526020808701928201019283116103dd57602001905b82821061042e575050506100c26060830161048d565b936100db60a06100d46080860161048d565b940161048d565b94331561041d57600180546001600160a01b031916331790558115801561040c575b80156103fb575b6103ea578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103a9575b5061037e575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610261575b5061010052604051615f6c9081610642823960805181818161185801528181611ac701528181611c41015281816121cc0152818161241e0152818161256101528181612e83015281816130a50152818161320d0152818161340e015281816135c001528181613898015281816138ed015281816139db015261419e015260a05181818161385401528181614cb701528181614d3a01526152da015260c051818181610bda015281816118e60152818161225a01528181612f11015261349e015260e051818181610b940152818161192b0152818161229f0152612c73015261010051818181611aec015281816124440152613dc60152f35b6020604051610270828261045c565b60008152600036813760e0511561036d5760005b81518110156102eb576001906001600160a01b036102a282856104a1565b5116846102ae826104e3565b6102bb575b505001610284565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138846102b3565b505060005b8251811015610364576001906001600160a01b0361030e82866104a1565b5116801561035e5783610320826105e1565b61032e575b50505b016102f0565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13883610325565b50610328565b50505038610169565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036103925750610137565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103e2575b816103c56020938361045c565b810103126103dd576103d69061047f565b9038610131565b600080fd5b3d91506103b8565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610104565b506001600160a01b038416156100fd565b639b15e16f60e01b60005260046000fd5b6020809161043b8461048d565b8152019101906100ac565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761044657604052565b519060ff821682036103dd57565b51906001600160a01b03821682036103dd57565b80518210156104b55760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104b55760005260206000200190600090565b60008181526003602052604090205480156105da5760001981018181116105c4576002546000198101919082116105c457818103610573575b505050600254801561055d57600019016105378160026104cb565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6105ac6105846105959360026104cb565b90549060031b1c92839260026104cb565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061051c565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461063b57600254680100000000000000008110156104465761062261059582600185940160025560026104cb565b9055600254906000526003602052604060002055600190565b5060009056fe60c080604052600436101561001357600080fd5b600060a05260a0513560e01c90816301ffc9a714613a0f57508063164e68de14613973578063181f5a771461391157806321df0da7146138cc578063240028e81461387857806324f65ee71461383957806337b192471461372557806338b39d29146115d757806339077537146133875780633ce85d41146132da578063489a68f214612dea5780634c5ef0ed14612da557806354c8a4f314612c415780635df45a3714612c1d57806362ddd3c414612b98578063698c2c6614612ae65780636d3d1a5814612abe5780637437ff9f14612a8a57806379ba5097146129cd5780637d54534e14612949578063804ba5a9146128e15780638926f54f1461289c57806389720a62146128255780638a86fe14146127e05780638da5cb5b146127b8578063962d4020146126755780639a4575b91461217a578063a42a7b8b14612023578063a7cd63b714611fb5578063acfecf9114611e88578063af58d59f14611e3c578063b1c71c65146117db578063b7946580146117a3578063c4bffe2b14611687578063c75eea9c146115dc578063c8de9fe0146115d7578063cf7401f31461144a578063d966866b14610ffd578063dc04fa1f14610bfe578063dc0bd97114610bb9578063e0351e1314610b7b578063e8a1da17146102c35763f2fde38b146101fe57600080fd5b346102bd5760206003193601126102bd576001600160a01b0361021f613b40565b610227614748565b163381146102915760a05180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360a05180f35b7fdad89dca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b60a05180fd5b346102bd576102d136613ea1565b9190926102dc614748565b60a051905b8282106109b95750505060a0519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156109b357600581901b830135858112156102bd578301610120813603126102bd576040519461035086613bed565b813567ffffffffffffffff811681036109ae578652602082013567ffffffffffffffff81116102bd5782019436601f870112156102bd578535610392816140f6565b966103a06040519889613c25565b81885260208089019260051b820101903682116102bd5760208101925b82841061097f575050505060208701958652604083013567ffffffffffffffff81116102bd576103f09036908501613e52565b916040880192835261041a6104083660608701613ff8565b9460608a0195865260c0369101613ff8565b956080890196875261042c8551615529565b6104368751615529565b835151156109535761045267ffffffffffffffff8a5116615de4565b156109185767ffffffffffffffff89511660a051526008602052604060a0512061059686516fffffffffffffffffffffffffffffffff604082015116906105516fffffffffffffffffffffffffffffffff602083015116915115158360806040516104bc81613bed565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106bc88516fffffffffffffffffffffffffffffffff604082015116906106776fffffffffffffffffffffffffffffffff602083015116915115158360806040516105e081613bed565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff82116108e7576106df83546143e2565b601f81116108a8575b506020906001601f84111461083e5791809161071b9360a05192610833575b50506000198260011b9260031b1c19161790565b90555b60a0515b88518051821015610757579061075160019261074a8367ffffffffffffffff8f5116926143ce565b5190614e3e565b01610722565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561082567ffffffffffffffff60019796949851169251935191516107f16107bc60405196879687526101006020880152610100870190613c64565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101939290919361031e565b015190508e80610707565b90601f198316918460a051528160a051209260a0515b8181106108905750908460019594939210610877575b505050811b01905561071e565b015160001960f88460031b161c191690558d808061086a565b92936020600181928786015181550195019301610854565b6108d7908460a05152602060a05120601f850160051c810191602086106108dd575b601f0160051c01906146a1565b8d6106e8565b90915081906108ca565b7f4e487b710000000000000000000000000000000000000000000000000000000060a051526041600452602460a051fd5b67ffffffffffffffff8951167f1d5ad3c50000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b833567ffffffffffffffff81116102bd576020916109a38392833691870101613e52565b8152019301926103bd565b600080fd5b60a05180f35b9092919367ffffffffffffffff6109d96109d4868886614366565b6140a4565b16926109e484615c15565b15610b4b578360a051526008602052610a046005604060a0512001615ab2565b9260a0515b8451811015610a43576001908660a051526008602052610a3c6005604060a0512001610a3583896143ce565b5190615cc8565b5001610a09565b50939094919592508060a0515260086020526005604060a0512060a051815560a051600182015560a051600282015560a051600382015560048101610a8881546143e2565b80610afb575b50500180549060a051815581610ad7575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102e1565b60a05152602060a05120908101905b81811015610a9f5760a0518155600101610ae6565b601f8111600114610b14575060a05190555b8880610a8e565b610b34908260a051526001601f602060a051209201861c820191016146a1565b60a080518290525160208120918190559055610b0d565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102bd5760a0516003193601126102bd5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102bd5760a0516003193601126102bd5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bd5760406003193601126102bd5760043567ffffffffffffffff81116102bd57366023820112156102bd57806004013567ffffffffffffffff81116102bd576024820191602436918360081b0101116102bd5760243567ffffffffffffffff81116102bd57610c74903690600401613e70565b919092610c7f614748565b60a0515b828110610cfb5750505060a0515b818110610c9e5760a05180f35b8067ffffffffffffffff610cb86109d46001948688614366565b168060a05152600f60205260a051604060a05120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee860a05160a051a201610c91565b610d096109d48285856146fa565b610d148285856146fa565b602081019060a081019261271061ffff610d2d8661470a565b161015610ff15760c082019061271061ffff610d488461470a565b161015610fb55767ffffffffffffffff16938460a05152600f60205260a0516040902092610d7585614719565b63ffffffff168454916040810190610d8c82614719565b60201b67ffffffff0000000016936060820193610da885614719565b60401b6bffffffff000000000000000016956080840196610dc888614719565b60601b6fffffffff0000000000000000000000001691610de78a61470a565b60801b71ffff000000000000000000000000000000001693610e088c61470a565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717875560e00195610ebf8761472a565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610f1090614737565b63ffffffff168752610f2190614737565b63ffffffff166020870152610f3590614737565b63ffffffff166040860152610f4990614737565b63ffffffff166060850152610f5d90613cde565b61ffff166080840152610f6f90613cde565b61ffff1660a0830152610f8190613fce565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c83565b61ffff610fc18361470a565b7f95f3517a0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b61ffff610fc18561470a565b346102bd5760206003193601126102bd5760043567ffffffffffffffff81116102bd5761102e903690600401613e70565b90611037614748565b60a0516080525b816080511061104d5760a05180f35b61105d6109d460805184846145cb565b61107761106d60805185856145cb565b602081019061460b565b61109161108760805187876145cb565b604081019061460b565b906110ac6110a260805189896145cb565b606081019061460b565b6110c66110bc6080518b8b6145cb565b608081019061460b565b9390946110dc6110d736898b61410e565b615486565b6110ea6110d736838561410e565b6110f86110d736858761410e565b6111066110d736878961410e565b6040519860808a01908a821067ffffffffffffffff8311176108e75767ffffffffffffffff9160405261113a368a8c61410e565b8b5261114736848661410e565b60208c015261115736868861410e565b60408c015261116736888a61410e565b60608c015216988960a05152600e602052604060a05120815180519067ffffffffffffffff82116108e7576801000000000000000082116108e757602090835483855580841061142b575b50018260a05152602060a0512060a0515b83811061140e57505050506001810160208301519081519167ffffffffffffffff83116108e7576801000000000000000083116108e75760209082548484558085106113ef575b50019060a05152602060a0512060a0515b8381106113d257505050506002810160408301519081519167ffffffffffffffff83116108e7576801000000000000000083116108e75760209082548484558085106113b3575b50019060a05152602060a0512060a0515b838110611396575050505060036060910191015180519067ffffffffffffffff82116108e7576801000000000000000082116108e7576020908354838555808410611377575b50019160a05152602060a051209160a0515b82811061135a575050505092611349949261132d61133b937fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a99989661131f6040519b8c9b60808d5260808d01916146b8565b918a830360208c01526146b8565b9187830360408901526146b8565b9184830360608601526146b8565b0390a260016080510160805261103e565b60019060206001600160a01b0384511693019281860155016112cb565b611390908560a05152848460a0512091820191016146a1565b8f6112b9565b60019060206001600160a01b038551169401938184015501611273565b6113cc908460a05152858460a0512091820191016146a1565b38611262565b60019060206001600160a01b03855116940193818401550161121b565b611408908460a05152858460a0512091820191016146a1565b3861120a565b60019060206001600160a01b0385511694019381840155016111c3565b611444908560a05152848460a0512091820191016146a1565b386111b2565b346102bd5760e06003193601126102bd57611463613ca5565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01126102bd5760405161149981613c09565b60243580151581036102bd5781526044356fffffffffffffffffffffffffffffffff811681036102bd5760208201526064356fffffffffffffffffffffffffffffffff811681036102bd5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c01126102bd576040519061152082613c09565b60843580151581036102bd57825260a4356fffffffffffffffffffffffffffffffff811681036102bd57602083015260c4356fffffffffffffffffffffffffffffffff811681036102bd5760408301526001600160a01b03600a5416331415806115c2575b611592576109b3926151a3565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506001600160a01b0360015416331415611585565b613da6565b346102bd5760206003193601126102bd5767ffffffffffffffff6115fe613ca5565b611606614518565b501660a05152600860205261168361162a611625604060a05120614543565b61531b565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102bd5760a0516003193601126102bd5760a051506040516006548082528160208101600660a05152602060a051209260a0515b81811061178a5750506116d192500382613c25565b8051906116f66116e0836140f6565b926116ee6040519485613c25565b8084526140f6565b90601f1960208401920136833760a0515b8151811015611739578067ffffffffffffffff611726600193856143ce565b511661173282876143ce565b5201611707565b5050906040519182916020830190602084525180915260408301919060a0515b818110611767575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611759565b84548352600194850194869450602090930192016116bc565b346102bd5760206003193601126102bd576116836117c76117c2613ca5565b6145a9565b604051918291602083526020830190613c64565b346102bd5760606003193601126102bd5760043567ffffffffffffffff81116102bd5760a060031982360301126102bd57611814613cbc565b9060443567ffffffffffffffff81116102bd57611835903690600401613e52565b5061183e6143b5565b50608481019061184d82614090565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611dfb57602481019177ffffffffffffffff000000000000000000000000000000006118a6846140a4565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d115760a05191611dcc575b50611da05761192960448301614090565b7f0000000000000000000000000000000000000000000000000000000000000000611d4d575b5067ffffffffffffffff611962846140a4565b1661197a816000526007602052604060002054151590565b15611d1e5760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611d115760a05190611cc8575b6001600160a01b039150163303611c985760648201359061ffff851680151580611c89575b15611bdd5761ffff600b541690818110611bab575050611ba194611b70936117c2937f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff611ac495611a7e611a6e611a548c6140a4565b67ffffffffffffffff16600052600c602052604060002090565b85611a7884614090565b916158a3565b611ab8611a93611a8d8c6140a4565b92614090565b94604051938493169583602090939291936001600160a01b0360408201951681520152565b0390a25b6004016153a0565b927f0000000000000000000000000000000000000000000000000000000000000000611b11857f000000000000000000000000000000000000000000000000000000000000000083614786565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff611b44846140a4565b604080516001600160a01b0390951685523360208601528401889052169180606081015b0390a26140a4565b90611b796152d3565b60405192611b8684613bd1565b83526020830152604051928392604084526040840190613fa4565b9060208301520390f35b7f7911d95b0000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b5050611ac4611ba194611b70936117c2937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611c21896140a4565b16918260a05152600860205280611c69604060a051206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916158a3565b604080516001600160a01b039290921682526020820192909252a2611abc565b5061ffff600b541615156119f5565b7f728fe07b0000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506020813d602011611d09575b81611ce260209383613c25565b810103126102bd57516001600160a01b03811681036102bd576001600160a01b03906119d0565b3d9150611cd5565b6040513d60a051823e3d90fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b6001600160a01b0316611d6d816000526003602052604060002054151590565b61194f577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b611dee915060203d602011611df4575b611de68183613c25565b810190614e26565b85611918565b503d611ddc565b6001600160a01b03611e0c83614090565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b346102bd5760206003193601126102bd5767ffffffffffffffff611e5e613ca5565b611e66614518565b501660a05152600860205261168361162a6116256002604060a0512001614543565b346102bd5767ffffffffffffffff611e9f36613eef565b929091611eaa614748565b1690611ec3826000526007602052604060002054151590565b15611f85578160a051526008602052611ef66005604060a0512001611ee9368685613e1b565b6020815191012090615cc8565b15611f3e577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611f356040519283926020845260208401916144f7565b0390a260a05180f35b611f81906040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916144f7565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102bd5760a0516003193601126102bd5760a051506040516002548082526020820190600260a05152602060a051209060a0515b81811061200d576116838561200181870382613c25565b60405191829182613f30565b8254845260209093019260019283019201611fea565b346102bd5760206003193601126102bd5767ffffffffffffffff612045613ca5565b1660a0515260086020526120606005604060a0512001615ab2565b805190601f19612088612072846140f6565b936120806040519586613c25565b8085526140f6565b0160a0515b81811061216957505060a0515b81518110156120e457806120b0600192846143ce565b5160a0515260096020526120c8604060a05120614435565b6120d282866143ce565b526120dd81856143ce565b500161209a565b826040518091602082016020835281518091526040830190602060408260051b86010193019160a051905b82821061211e57505050500390f35b91936020612159827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613c64565b960192019201859493919261210f565b80606060208093870101520161208d565b346102bd5760206003193601126102bd5760043567ffffffffffffffff81116102bd5760a060031982360301126102bd576121b36143b5565b50608481016121c181614090565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361266457602482019177ffffffffffffffff0000000000000000000000000000000061221a846140a4565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d115760a05191612645575b50611da05761229d60448201614090565b7f00000000000000000000000000000000000000000000000000000000000000006125f2575b5067ffffffffffffffff6122d6846140a4565b166122ee816000526007602052604060002054151590565b15611d1e5760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611d115760a051906125a9575b6001600160a01b039150163303611c9857606401359060a0516000146125045761ffff600b5416806124cf57508261249f926117c29261239661238c611a54611683986140a4565b83611a7884614090565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff6123d26123cc866140a4565b93614090565b604080516001600160a01b03929092168252602082018690529190931692a25b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000092612469817f000000000000000000000000000000000000000000000000000000000000000086614786565b611b68612475866140a4565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b6124a76152d3565b604051916124b483613bd1565b82526020820152604051918291602083526020830190613fa4565b7f7911d95b0000000000000000000000000000000000000000000000000000000060a0515260a051600452602452604460a051fd5b506117c28261249f9267ffffffffffffffff612522611683966140a4565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944828061258960406000206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916158a3565b604080516001600160a01b039290921682526020820192909252a26123f2565b506020813d6020116125ea575b816125c360209383613c25565b810103126102bd57516001600160a01b03811681036102bd576001600160a01b0390612344565b3d91506125b6565b6001600160a01b0316612612816000526003602052604060002054151590565b6122c3577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b61265e915060203d602011611df457611de68183613c25565b8461228c565b611e0c6001600160a01b0391614090565b346102bd5760606003193601126102bd5760043567ffffffffffffffff81116102bd576126a6903690600401613e70565b9060243567ffffffffffffffff81116102bd576126c7903690600401613f73565b9060443567ffffffffffffffff81116102bd576126e8903690600401613f73565b6001600160a01b03600a5416331415806127a3575b61159257838614801590612799575b61276d5760a0515b8681106127215760a05180f35b806127676127356109d46001948b8b614366565b6127408389896143a5565b61276161275961275186898b6143a5565b923690613ff8565b913690613ff8565b916151a3565b01612714565b7f568efce20000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b508086141561270c565b506001600160a01b03600154163314156126fd565b346102bd5760a0516003193601126102bd5760206001600160a01b0360015416604051908152f35b346102bd57608063ffffffff61ffff816128066127fc36613d1b565b5050925050614283565b9392959091604051968752166020860152166040840152166060820152f35b346102bd5760c06003193601126102bd5761283e613b40565b5060243567ffffffffffffffff811681036102bd5761285b613ccd565b5060843567ffffffffffffffff81116102bd5761287c903690600401613ced565b505060a4359060028210156102bd5761168391612001916044359061420d565b346102bd5760206003193601126102bd5760206128d767ffffffffffffffff6128c3613ca5565b166000526007602052604060002054151590565b6040519015158152f35b346102bd5760206003193601126102bd5760043567ffffffffffffffff81116102bd57612912903690600401613dea565b6001600160a01b03600a541633141580612934575b611592576109b39161491d565b506001600160a01b0360015416331415612927565b346102bd5760206003193601126102bd577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b0361298d613b40565b612995614748565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a160a05180f35b346102bd5760a0516003193601126102bd5760a051546001600160a01b0381163303612a5e577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660a051556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060a05160a051a360a05180f35b7f02b543c60000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346102bd5760a0516003193601126102bd57600454600554604080516001600160a01b039093168352602083019190915290f35b346102bd5760a0516003193601126102bd5760206001600160a01b03600a5416604051908152f35b346102bd5760406003193601126102bd57612aff613b40565b602435612b0a614748565b6001600160a01b038216918215610953577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045581600555612b8f60405192839283602090939291936001600160a01b0360408201951681520152565b0390a160a05180f35b346102bd57612ba636613eef565b612bb1929192614748565b67ffffffffffffffff8216612bd3816000526007602052604060002054151590565b15612bee57506109b392612be8913691613e1b565b90614e3e565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102bd5760a0516003193601126102bd576020612c39614162565b604051908152f35b346102bd57612c69612c71612c5536613ea1565b9491612c62939193614748565b369161410e565b92369161410e565b7f000000000000000000000000000000000000000000000000000000000000000015612d795760a0515b8251811015612d0157806001600160a01b03612cb9600193866143ce565b5116612cc481615b15565b612cd0575b5001612c9b565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612cc9565b5060a0515b81518110156109b357806001600160a01b03612d24600193856143ce565b51168015612d7357612d3581615d84565b612d42575b505b01612d06565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612d3a565b50612d3c565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346102bd5760406003193601126102bd57612dbe613ca5565b60243567ffffffffffffffff81116102bd57602091612de46128d7923690600401613e52565b906140b9565b346102bd5760406003193601126102bd5760043567ffffffffffffffff81116102bd57806004019061010060031982360301126102bd57612e29613cbc565b604051909190612e3881613bb5565b60a0519052612e69612e5f612e5a612e5360c485018761403f565b3691613e1b565b614c43565b6064830135614d37565b916084820190612e7882614090565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611dfb57602483019477ffffffffffffffff00000000000000000000000000000000612ed1876140a4565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d115760a051916132bb575b50611da05767ffffffffffffffff612f5a876140a4565b16612f72816000526007602052604060002054151590565b15611d1e5760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611d115760a0519161329c575b5015611c9857612fde866140a4565b90612ff460a4860192612de4612e53858561403f565b156132555750604493929161ffff161590506131b65767ffffffffffffffff61301c866140a4565b1660a05152600d602052613039604060a0512085611a7884614090565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff767ffffffffffffffff61306f6123cc886140a4565b604080516001600160a01b03929092168252602082018890529190931692a25b019161309a83614090565b906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102bd576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390921660048201526024810185905290818060448101038160a051875af18015611d115761319d575b50608067ffffffffffffffff6020956001600160a01b0361316b611a8d7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966140a4565b60405196875233898801521660408601528560608601521692a26040519061319282613bb5565b815260405190518152f35b60a0516131a991613c25565b60a0516102bd5784613127565b5067ffffffffffffffff6131c9856140a4565b168060a0515260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c84806132356002604060a05120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916158a3565b604080516001600160a01b039290921682526020820192909252a261308f565b61325f925061403f565b611f816040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916144f7565b6132b5915060203d602011611df457611de68183613c25565b87612fcf565b6132d4915060203d602011611df457611de68183613c25565b87612f43565b346102bd5760406003193601126102bd5760043561ffff8116908190036102bd5760243567ffffffffffffffff81116102bd577f401d6d5b7a131da25eadbfe5bc7b2c80d7a0094d8d026d2d2c452f81e9b2af7e9161337a6133426020933690600401613dea565b9061334b614748565b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b5561491d565b604051908152a160a05180f35b346102bd5760206003193601126102bd5760043567ffffffffffffffff81116102bd57806004019061010060031982360301126102bd576040516133ca81613bb5565b60a05190526040516133db81613bb5565b60a05190526133f6612e5f612e5a612e5360c485018661403f565b906084810161340481614090565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116036126645750602481019277ffffffffffffffff0000000000000000000000000000000061345e856140a4565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d115760a05191613706575b50611da05767ffffffffffffffff6134e7856140a4565b166134ff816000526007602052604060002054151590565b15611d1e5760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611d115760a051916136e7575b5015611c985761356b846140a4565b9061358160a4840192612de4612e53858561403f565b15613255575082916044915067ffffffffffffffff61359f866140a4565b168060a0515260086020526135e86002604060a05120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169586916158a3565b604080516001600160a01b0386168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a2019261362e84614090565b823b156102bd576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390921660048201526024810185905290818060448101038160a051875af18015611d11576020956001600160a01b0361316b611a8d7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09660809667ffffffffffffffff966136d5575b506140a4565b60a0516136e191613c25565b8b6136cf565b613700915060203d602011611df457611de68183613c25565b8561355c565b61371f915060203d602011611df457611de68183613c25565b856134d0565b346102bd5767ffffffffffffffff61373c36613d1b565b50505050905060405161374e81613b6a565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260c060a0519101521660a05152600f60205260e0604060a05120604051906137a282613b6a565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b346102bd5760a0516003193601126102bd57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bd5760206003193601126102bd576020613893613b40565b6040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081169216919091148152f35b346102bd5760a0516003193601126102bd5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bd5760a0516003193601126102bd5760408051611683916139359082613c25565b602081527f4275726e546f41646472657373546f6b656e506f6f6c20312e372e302d6465766020820152604051918291602083526020830190613c64565b346102bd5760206003193601126102bd5761398c613b40565b613994614748565b61399c614162565b90816139a85760a05180f35b60206001600160a01b03826139ff857f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599957f0000000000000000000000000000000000000000000000000000000000000000614786565b6040519485521692a280806109b3565b346102bd5760206003193601126102bd57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102bd57817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115613b16575b8115613aec575b8115613ac2575b8115613a98575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613a91565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613a8a565b7f7a8c77260000000000000000000000000000000000000000000000000000000081149150613a83565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613a7c565b600435906001600160a01b03821682036109ae57565b35906001600160a01b03821682036109ae57565b60e0810190811067ffffffffffffffff821117613b8657604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613b8657604052565b6040810190811067ffffffffffffffff821117613b8657604052565b60a0810190811067ffffffffffffffff821117613b8657604052565b6060810190811067ffffffffffffffff821117613b8657604052565b90601f601f19910116810190811067ffffffffffffffff821117613b8657604052565b67ffffffffffffffff8111613b8657601f01601f191660200190565b919082519283825260005b848110613c90575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201613c6f565b6004359067ffffffffffffffff821682036109ae57565b6024359061ffff821682036109ae57565b6064359061ffff821682036109ae57565b359061ffff821682036109ae57565b9181601f840112156109ae5782359167ffffffffffffffff83116109ae57602083818601950101116109ae57565b60a06003198201126109ae576004356001600160a01b03811681036109ae579160243567ffffffffffffffff811681036109ae579160443567ffffffffffffffff81116109ae5760a060031982840301126109ae576004019160643561ffff811681036109ae57916084359067ffffffffffffffff82116109ae57613da291600401613ced565b9091565b346109ae5760006003193601126109ae5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b9181601f840112156109ae5782359167ffffffffffffffff83116109ae5760208085019460e085020101116109ae57565b929192613e2782613c48565b91613e356040519384613c25565b8294818452818301116109ae578281602093846000960137010152565b9080601f830112156109ae57816020613e6d93359101613e1b565b90565b9181601f840112156109ae5782359167ffffffffffffffff83116109ae576020808501948460051b0101116109ae57565b60406003198201126109ae5760043567ffffffffffffffff81116109ae5781613ecc91600401613e70565b929092916024359067ffffffffffffffff82116109ae57613da291600401613e70565b9060406003198301126109ae5760043567ffffffffffffffff811681036109ae57916024359067ffffffffffffffff82116109ae57613da291600401613ced565b602060408183019282815284518094520192019060005b818110613f545750505090565b82516001600160a01b0316845260209384019390920191600101613f47565b9181601f840112156109ae5782359167ffffffffffffffff83116109ae57602080850194606085020101116109ae57565b613e6d916020613fbd8351604084526040840190613c64565b920151906020818403910152613c64565b359081151582036109ae57565b35906fffffffffffffffffffffffffffffffff821682036109ae57565b91908260609103126109ae5760405161401081613c09565b604061403a81839561402181613fce565b855261402f60208201613fdb565b602086015201613fdb565b910152565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109ae570180359067ffffffffffffffff82116109ae576020019181360383136109ae57565b356001600160a01b03811681036109ae5790565b3567ffffffffffffffff811681036109ae5790565b9067ffffffffffffffff613e6d92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613b865760051b60200190565b92919061411a816140f6565b936141286040519586613c25565b602085838152019160051b81019283116109ae57905b82821061414a57505050565b6020809161415784613b56565b81520191019061413e565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115614201576000916141d2575090565b90506020813d6020116141f9575b816141ed60209383613c25565b810103126109ae575190565b3d91506141e0565b6040513d6000823e3d90fd5b67ffffffffffffffff16600052600e60205260406000209160028110156142545760011461424357816001613e6d9301906150af565b8160026003613e6d940191016150af565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b67ffffffffffffffff16600052600f602052604060002090604051916142a883613b6a565b549263ffffffff84169384845263ffffffff8160201c169384602082015263ffffffff8260401c169384604083015263ffffffff8360601c169081606084015261ffff8460801c169283608082015260ff61ffff8660901c16958660a084015260a01c16159060c082159101526143505761ffff161580159563ffffffff949392916143485750945b156143405750925b1693929190565b905092614339565b905094614331565b5050505092505050600090600090600090600090565b91908110156143765760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190811015614376576060020190565b604051906143c282613bd1565b60606020838281520152565b80518210156143765760209160051b010190565b90600182811c9216801561442b575b60208310146143fc57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916143f1565b9060405191826000825492614449846143e2565b80845293600181169081156144b75750600114614470575b5061446e92500383613c25565b565b90506000929192526020600020906000915b81831061449b57505090602061446e9282010138614461565b6020919350806001915483858901015201910190918492614482565b6020935061446e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614461565b601f8260209493601f19938186528686013760008582860101520116010190565b6040519061452582613bed565b60006080838281528260208201528260408201528260608201520152565b9060405161455081613bed565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526008602052613e6d6004604060002001614435565b91908110156143765760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156109ae570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109ae570180359067ffffffffffffffff82116109ae57602001918160051b360383136109ae57565b8181029291811591840414171561467257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106146ac575050565b600081556001016146a1565b9160209082815201919060005b8181106146d25750505090565b9091926020806001926001600160a01b036146ec88613b56565b1681520194019291016146c5565b91908110156143765760081b0190565b3561ffff811681036109ae5790565b3563ffffffff811681036109ae5790565b3580151581036109ae5790565b359063ffffffff821682036109ae57565b6001600160a01b0360015416330361475c57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082019081526001600160a01b03938416602483015260448083019590955293815261484f9290916147de606484613c25565b166000806040958651946147f28887613c25565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d156148f8573d9161483383613c48565b9261484087519485613c25565b83523d6000602085013e615e93565b8051908161485c57505050565b60208061486d938301019101614e26565b156148755750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615e93565b356fffffffffffffffffffffffffffffffff811681036109ae5790565b9160005b82811015614c3d5760e081028401600061493a826140a4565b9067ffffffffffffffff82169161495e836000526007602052604060002054151590565b15614c1157614a2792604085936149d26149cc946149cc614992602060019c9b0192611a5461498d3686613ff8565b615529565b91825463ffffffff8160801c16159081614bf3575b81614be4575b81614bc9575b81614bba575b5080614bab575b614b20575b3690613ff8565b90615670565b60808501926149e461498d3686613ff8565b8152600d6020522092835463ffffffff8160801c16159081614b02575b81614af3575b81614ad8575b81614ac9575b5080614aba575b614a2d575b503690613ff8565b01614921565b614a4a60a06fffffffffffffffffffffffffffffffff9201614900565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614a1f565b50614ac48261472a565b614a1a565b60ff915060a01c161538614a13565b6fffffffffffffffffffffffffffffffff8116159150614a0d565b8589015460801c159150614a07565b858901546fffffffffffffffffffffffffffffffff16159150614a01565b6fffffffffffffffffffffffffffffffff614b3c878b01614900565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783556149c5565b50614bb58161472a565b6149c0565b60ff915060a01c1615386149b9565b6fffffffffffffffffffffffffffffffff81161591506149b3565b848e015460801c1591506149ad565b848e01546fffffffffffffffffffffffffffffffff161591506149a7565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b80518015614cb357602003614c755780516020828101918301839003126109ae57519060ff8211614c75575060ff1690565b611f81906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613c64565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161467257565b60ff16604d811161467257600a0a90565b8115614d08570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614e1f57828411614df55790614d7c91614cd9565b91604d60ff8416118015614dda575b614da457505090614d9e613e6d92614ced565b9061465f565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614de483614ced565b8015614d0857600019048411614d8b565b614dfe91614cd9565b91604d60ff841611614da457505090614e19613e6d92614ced565b90614cfe565b5050505090565b908160209103126109ae575180151581036109ae5790565b908051156150245767ffffffffffffffff81516020830120921691826000526008602052614e73816005604060002001615e3e565b15614fe05760005260096020526040600020815167ffffffffffffffff8111613b8657614ea082546143e2565b601f8111614fae575b506020601f8211600114614f245791614efe827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614f1495600091614f19575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613c64565b0390a2565b905084015138614eeb565b601f1982169083600052806000209160005b818110614f96575092614f149492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614f7d575b5050811b0190556117c7565b85015160001960f88460031b161c191690553880614f71565b9192602060018192868a015181550194019201614f36565b614fda90836000526020600020601f840160051c810191602085106108dd57601f0160051c01906146a1565b38614ea9565b5090611f816040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613c64565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b81811061508057505061446e92500383613c25565b84546001600160a01b031683526001948501948794506020909301920161506b565b9190820180921161467257565b6150b89061504e565b916005548015159182615198575b50506150d0575090565b6150d99061504e565b908151806150e75750905090565b6150f29082516150a2565b92601f19615118615102866140f6565b956151106040519788613c25565b8087526140f6565b0136602086013760005b825181101561515357806001600160a01b03615140600193866143ce565b511661514c82886143ce565b5201615122565b509160005b815181101561519357806001600160a01b03615176600193856143ce565b511661518c6151868387516150a2565b886143ce565b5201615158565b505050565b1015905038806150c6565b67ffffffffffffffff1660008181526007602052604090205490929190156152a557916152a260e09261526e856151fa7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97615529565b846000526008602052615211816040600020615670565b61521a83615529565b846000526008602052615234836002604060002001615670565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613e6d604082613c25565b9190820391821161467257565b615323614518565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691615380602085019361537a61536d63ffffffff8751164261530e565b856080890151169061465f565b906150a2565b8082101561539957505b16825263ffffffff4216905290565b905061538a565b67ffffffffffffffff6153b5602083016140a4565b16600052600f6020526040600020604051906153d082613b6a565b549063ffffffff8216815263ffffffff8260201c16602082015263ffffffff8260401c16604082015263ffffffff8260601c16606082015261ffff8260801c169081608082015260ff61ffff8460901c16938460a084015260a01c16159060c0821591015261547b57919261ffff929083161561547457505b16801561546c576127106154656060613e6d940135928361465f565b049061530e565b506060013590565b9050615449565b505060609150013590565b805160005b81811061549757505050565b60018101808211614672575b8281106154b3575060010161548b565b6001600160a01b036154c583866143ce565b51166001600160a01b036154d983876143ce565b5116146154e8576001016154a3565b6001600160a01b036154fa83866143ce565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8051156155c9576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106155665750565b6064906155c7604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590615651575b6155f05750565b6064906155c7604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208201511615156155e9565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916157a960609280546156ad63ffffffff8260801c164261530e565b90816157e8575b50506fffffffffffffffffffffffffffffffff60018160208601511692828154168085106000146157e057508280855b16167fffffffffffffffffffffffffffffffff0000000000000000000000000000000082541617815561575d8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6152a260405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916156e4565b6fffffffffffffffffffffffffffffffff9161581d8392836158166001880154948286169560801c9061465f565b91166150a2565b8082101561589c57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff000000000000000000000000000000001617815538806156b4565b9050615827565b9182549060ff8260a01c16158015615aaa575b615aa4576fffffffffffffffffffffffffffffffff821691600185019081546158fb63ffffffff6fffffffffffffffffffffffffffffffff83169360801c164261530e565b9081615a06575b50508481106159c7575083831061595c5750506159316fffffffffffffffffffffffffffffffff92839261530e565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c9161596b818561530e565b926000198101908082116146725761598e615993926001600160a01b03966150a2565b614cfe565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615a7a57615a219261537a9160801c9061465f565b80841015615a755750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615902565b615a2c565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156158b6565b906040519182815491828252602082019060005260206000209260005b818110615ae457505061446e92500383613c25565b8454835260019485019487945060209093019201615acf565b80548210156143765760005260206000200190600090565b6000818152600360205260409020548015615c0e5760001981018181116146725760025490600019820191821161467257818103615bbd575b5050506002548015615b8e5760001901615b69816002615afd565b60001982549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615bf6615bce615bdf936002615afd565b90549060031b1c9283926002615afd565b81939154906000199060031b92831b921b19161790565b90556000526003602052604060002055388080615b4e565b5050600090565b6000818152600760205260409020548015615c0e5760001981018181116146725760065490600019820191821161467257818103615c8e575b5050506006548015615b8e5760001901615c69816006615afd565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b615cb0615c9f615bdf936006615afd565b90549060031b1c9283926006615afd565b90556000526007602052604060002055388080615c4e565b9060018201918160005282602052604060002054801515600014615d7b57600019810181811161467257825490600019820191821161467257818103615d44575b50505080548015615b8e576000190190615d238282615afd565b60001982549160031b1b191690555560005260205260006040812055600190565b615d64615d54615bdf9386615afd565b90549060031b1c92839286615afd565b905560005283602052604060002055388080615d09565b50505050600090565b80600052600360205260406000205415600014615dde5760025468010000000000000000811015613b8657615dc5615bdf8260018594016002556002615afd565b9055600254906000526003602052604060002055600190565b50600090565b80600052600760205260406000205415600014615dde5760065468010000000000000000811015613b8657615e25615bdf8260018594016006556006615afd565b9055600654906000526007602052604060002055600190565b6000828152600182016020526040902054615c0e5780549068010000000000000000821015613b865782615e7c615bdf846001809601855584615afd565b905580549260005201602052604060002055600190565b91929015615f0e5750815115615ea7575090565b3b15615eb05790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615f215750805190602001fd5b611f81906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613c6456fea164736f6c634300081a000a",
}

var BurnToAddressMintTokenPoolABI = BurnToAddressMintTokenPoolMetaData.ABI

var BurnToAddressMintTokenPoolBin = BurnToAddressMintTokenPoolMetaData.Bin

func DeployBurnToAddressMintTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address, burnAddress common.Address) (common.Address, *types.Transaction, *BurnToAddressMintTokenPool, error) {
	parsed, err := BurnToAddressMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnToAddressMintTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router, burnAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnToAddressMintTokenPool{address: address, abi: *parsed, BurnToAddressMintTokenPoolCaller: BurnToAddressMintTokenPoolCaller{contract: contract}, BurnToAddressMintTokenPoolTransactor: BurnToAddressMintTokenPoolTransactor{contract: contract}, BurnToAddressMintTokenPoolFilterer: BurnToAddressMintTokenPoolFilterer{contract: contract}}, nil
}

type BurnToAddressMintTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnToAddressMintTokenPoolCaller
	BurnToAddressMintTokenPoolTransactor
	BurnToAddressMintTokenPoolFilterer
}

type BurnToAddressMintTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnToAddressMintTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnToAddressMintTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnToAddressMintTokenPoolSession struct {
	Contract     *BurnToAddressMintTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnToAddressMintTokenPoolCallerSession struct {
	Contract *BurnToAddressMintTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnToAddressMintTokenPoolTransactorSession struct {
	Contract     *BurnToAddressMintTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnToAddressMintTokenPoolRaw struct {
	Contract *BurnToAddressMintTokenPool
}

type BurnToAddressMintTokenPoolCallerRaw struct {
	Contract *BurnToAddressMintTokenPoolCaller
}

type BurnToAddressMintTokenPoolTransactorRaw struct {
	Contract *BurnToAddressMintTokenPoolTransactor
}

func NewBurnToAddressMintTokenPool(address common.Address, backend bind.ContractBackend) (*BurnToAddressMintTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnToAddressMintTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnToAddressMintTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPool{address: address, abi: abi, BurnToAddressMintTokenPoolCaller: BurnToAddressMintTokenPoolCaller{contract: contract}, BurnToAddressMintTokenPoolTransactor: BurnToAddressMintTokenPoolTransactor{contract: contract}, BurnToAddressMintTokenPoolFilterer: BurnToAddressMintTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnToAddressMintTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnToAddressMintTokenPoolCaller, error) {
	contract, err := bindBurnToAddressMintTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCaller{contract: contract}, nil
}

func NewBurnToAddressMintTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnToAddressMintTokenPoolTransactor, error) {
	contract, err := bindBurnToAddressMintTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolTransactor{contract: contract}, nil
}

func NewBurnToAddressMintTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnToAddressMintTokenPoolFilterer, error) {
	contract, err := bindBurnToAddressMintTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolFilterer{contract: contract}, nil
}

func bindBurnToAddressMintTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnToAddressMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnToAddressMintTokenPool.Contract.BurnToAddressMintTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.BurnToAddressMintTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.BurnToAddressMintTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnToAddressMintTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _BurnToAddressMintTokenPool.Contract.GetAccumulatedFees(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _BurnToAddressMintTokenPool.Contract.GetAccumulatedFees(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetAllowList(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetAllowList(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.GetAllowListEnabled(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.GetAllowListEnabled(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetBurnAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getBurnAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetBurnAddress() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetBurnAddress(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetBurnAddress() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetBurnAddress(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnToAddressMintTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnToAddressMintTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnToAddressMintTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnToAddressMintTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnToAddressMintTokenPool.Contract.GetDynamicConfig(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnToAddressMintTokenPool.Contract.GetDynamicConfig(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, finality uint16, arg4 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, finality, arg4)

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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, finality uint16, arg4 []byte) (GetFee,

	error) {
	return _BurnToAddressMintTokenPool.Contract.GetFee(&_BurnToAddressMintTokenPool.CallOpts, arg0, destChainSelector, arg2, finality, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, finality uint16, arg4 []byte) (GetFee,

	error) {
	return _BurnToAddressMintTokenPool.Contract.GetFee(&_BurnToAddressMintTokenPool.CallOpts, arg0, destChainSelector, arg2, finality, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRateLimitAdmin(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRateLimitAdmin(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRemotePools(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRemotePools(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRemoteToken(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRemoteToken(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRequiredCCVs(&_BurnToAddressMintTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRmnProxy(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetRmnProxy(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnToAddressMintTokenPool.Contract.GetSupportedChains(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnToAddressMintTokenPool.Contract.GetSupportedChains(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetToken(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.GetToken(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _BurnToAddressMintTokenPool.Contract.GetTokenDecimals(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnToAddressMintTokenPool.Contract.GetTokenDecimals(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnToAddressMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnToAddressMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnToAddressMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnToAddressMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) IBurnAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "i_burnAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) IBurnAddress() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.IBurnAddress(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) IBurnAddress() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.IBurnAddress(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.IsRemotePool(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.IsRemotePool(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.IsSupportedChain(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.IsSupportedChain(&_BurnToAddressMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.IsSupportedToken(&_BurnToAddressMintTokenPool.CallOpts, token)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.IsSupportedToken(&_BurnToAddressMintTokenPool.CallOpts, token)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) Owner() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.Owner(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnToAddressMintTokenPool.Contract.Owner(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.SupportsInterface(&_BurnToAddressMintTokenPool.CallOpts, interfaceId)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnToAddressMintTokenPool.Contract.SupportsInterface(&_BurnToAddressMintTokenPool.CallOpts, interfaceId)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnToAddressMintTokenPool.Contract.TypeAndVersion(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnToAddressMintTokenPool.Contract.TypeAndVersion(&_BurnToAddressMintTokenPool.CallOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.AcceptOwnership(&_BurnToAddressMintTokenPool.TransactOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.AcceptOwnership(&_BurnToAddressMintTokenPool.TransactOpts)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.AddRemotePool(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.AddRemotePool(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnToAddressMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnToAddressMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyChainUpdates(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyChainUpdates(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ApplyFinalityConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ApplyFinalityConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.LockOrBurn(&_BurnToAddressMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.LockOrBurn(&_BurnToAddressMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.LockOrBurn0(&_BurnToAddressMintTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.LockOrBurn0(&_BurnToAddressMintTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ReleaseOrMint(&_BurnToAddressMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ReleaseOrMint(&_BurnToAddressMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ReleaseOrMint0(&_BurnToAddressMintTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ReleaseOrMint0(&_BurnToAddressMintTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.RemoveRemotePool(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.RemoveRemotePool(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnToAddressMintTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetDynamicConfig(&_BurnToAddressMintTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetDynamicConfig(&_BurnToAddressMintTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetRateLimitAdmin(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetRateLimitAdmin(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.TransferOwnership(&_BurnToAddressMintTokenPool.TransactOpts, to)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.TransferOwnership(&_BurnToAddressMintTokenPool.TransactOpts, to)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.WithdrawFees(&_BurnToAddressMintTokenPool.TransactOpts, recipient)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.WithdrawFees(&_BurnToAddressMintTokenPool.TransactOpts, recipient)
}

type BurnToAddressMintTokenPoolAllowListAddIterator struct {
	Event *BurnToAddressMintTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolAllowListAdd)
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
		it.Event = new(BurnToAddressMintTokenPoolAllowListAdd)
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

func (it *BurnToAddressMintTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolAllowListAddIterator{contract: _BurnToAddressMintTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolAllowListAdd)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*BurnToAddressMintTokenPoolAllowListAdd, error) {
	event := new(BurnToAddressMintTokenPoolAllowListAdd)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolAllowListRemoveIterator struct {
	Event *BurnToAddressMintTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolAllowListRemove)
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
		it.Event = new(BurnToAddressMintTokenPoolAllowListRemove)
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

func (it *BurnToAddressMintTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolAllowListRemoveIterator{contract: _BurnToAddressMintTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolAllowListRemove)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*BurnToAddressMintTokenPoolAllowListRemove, error) {
	event := new(BurnToAddressMintTokenPoolAllowListRemove)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolCCVConfigUpdatedIterator struct {
	Event *BurnToAddressMintTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCCVConfigUpdated)
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
		it.Event = new(BurnToAddressMintTokenPoolCCVConfigUpdated)
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

func (it *BurnToAddressMintTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCCVConfigUpdatedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCCVConfigUpdated)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*BurnToAddressMintTokenPoolCCVConfigUpdated, error) {
	event := new(BurnToAddressMintTokenPoolCCVConfigUpdated)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolChainAddedIterator struct {
	Event *BurnToAddressMintTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolChainAdded)
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
		it.Event = new(BurnToAddressMintTokenPoolChainAdded)
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

func (it *BurnToAddressMintTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolChainAddedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolChainAdded)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnToAddressMintTokenPoolChainAdded, error) {
	event := new(BurnToAddressMintTokenPoolChainAdded)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolChainConfiguredIterator struct {
	Event *BurnToAddressMintTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolChainConfigured)
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
		it.Event = new(BurnToAddressMintTokenPoolChainConfigured)
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

func (it *BurnToAddressMintTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolChainConfiguredIterator{contract: _BurnToAddressMintTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolChainConfigured)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseChainConfigured(log types.Log) (*BurnToAddressMintTokenPoolChainConfigured, error) {
	event := new(BurnToAddressMintTokenPoolChainConfigured)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolChainRemovedIterator struct {
	Event *BurnToAddressMintTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolChainRemoved)
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
		it.Event = new(BurnToAddressMintTokenPoolChainRemoved)
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

func (it *BurnToAddressMintTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolChainRemovedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolChainRemoved)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnToAddressMintTokenPoolChainRemoved, error) {
	event := new(BurnToAddressMintTokenPoolChainRemoved)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolConfigChangedIterator struct {
	Event *BurnToAddressMintTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolConfigChanged)
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
		it.Event = new(BurnToAddressMintTokenPoolConfigChanged)
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

func (it *BurnToAddressMintTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolConfigChangedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolConfigChanged)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseConfigChanged(log types.Log) (*BurnToAddressMintTokenPoolConfigChanged, error) {
	event := new(BurnToAddressMintTokenPoolConfigChanged)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdatedIterator struct {
	Event *BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated)
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
		it.Event = new(BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated)
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

func (it *BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCustomFinalityMinimumBlockConfirmationUpdated(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CustomFinalityMinimumBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdatedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CustomFinalityMinimumBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCustomFinalityMinimumBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CustomFinalityMinimumBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityMinimumBlockConfirmationUpdated", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCustomFinalityMinimumBlockConfirmationUpdated(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated, error) {
	event := new(BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityMinimumBlockConfirmationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
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
		it.Event = new(BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
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

func (it *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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
		it.Event = new(BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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

func (it *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolDynamicConfigSetIterator struct {
	Event *BurnToAddressMintTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolDynamicConfigSet)
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
		it.Event = new(BurnToAddressMintTokenPoolDynamicConfigSet)
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

func (it *BurnToAddressMintTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolDynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolDynamicConfigSetIterator{contract: _BurnToAddressMintTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolDynamicConfigSet)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*BurnToAddressMintTokenPoolDynamicConfigSet, error) {
	event := new(BurnToAddressMintTokenPoolDynamicConfigSet)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolInboundRateLimitConsumedIterator struct {
	Event *BurnToAddressMintTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(BurnToAddressMintTokenPoolInboundRateLimitConsumed)
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

func (it *BurnToAddressMintTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolInboundRateLimitConsumedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolInboundRateLimitConsumed)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolInboundRateLimitConsumed, error) {
	event := new(BurnToAddressMintTokenPoolInboundRateLimitConsumed)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolLockedOrBurnedIterator struct {
	Event *BurnToAddressMintTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolLockedOrBurned)
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
		it.Event = new(BurnToAddressMintTokenPoolLockedOrBurned)
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

func (it *BurnToAddressMintTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolLockedOrBurnedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolLockedOrBurned)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*BurnToAddressMintTokenPoolLockedOrBurned, error) {
	event := new(BurnToAddressMintTokenPoolLockedOrBurned)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *BurnToAddressMintTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(BurnToAddressMintTokenPoolOutboundRateLimitConsumed)
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

func (it *BurnToAddressMintTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolOutboundRateLimitConsumedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolOutboundRateLimitConsumed)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolOutboundRateLimitConsumed, error) {
	event := new(BurnToAddressMintTokenPoolOutboundRateLimitConsumed)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnToAddressMintTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolOwnershipTransferRequested)
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
		it.Event = new(BurnToAddressMintTokenPoolOwnershipTransferRequested)
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

func (it *BurnToAddressMintTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnToAddressMintTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolOwnershipTransferRequestedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolOwnershipTransferRequested)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnToAddressMintTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnToAddressMintTokenPoolOwnershipTransferRequested)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolOwnershipTransferredIterator struct {
	Event *BurnToAddressMintTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolOwnershipTransferred)
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
		it.Event = new(BurnToAddressMintTokenPoolOwnershipTransferred)
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

func (it *BurnToAddressMintTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnToAddressMintTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolOwnershipTransferredIterator{contract: _BurnToAddressMintTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolOwnershipTransferred)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnToAddressMintTokenPoolOwnershipTransferred, error) {
	event := new(BurnToAddressMintTokenPoolOwnershipTransferred)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolPoolFeeWithdrawnIterator struct {
	Event *BurnToAddressMintTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolPoolFeeWithdrawn)
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
		it.Event = new(BurnToAddressMintTokenPoolPoolFeeWithdrawn)
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

func (it *BurnToAddressMintTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnToAddressMintTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolPoolFeeWithdrawnIterator{contract: _BurnToAddressMintTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolPoolFeeWithdrawn)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*BurnToAddressMintTokenPoolPoolFeeWithdrawn, error) {
	event := new(BurnToAddressMintTokenPoolPoolFeeWithdrawn)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolRateLimitAdminSetIterator struct {
	Event *BurnToAddressMintTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRateLimitAdminSet)
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
		it.Event = new(BurnToAddressMintTokenPoolRateLimitAdminSet)
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

func (it *BurnToAddressMintTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRateLimitAdminSetIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRateLimitAdminSet)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*BurnToAddressMintTokenPoolRateLimitAdminSet, error) {
	event := new(BurnToAddressMintTokenPoolRateLimitAdminSet)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolReleasedOrMintedIterator struct {
	Event *BurnToAddressMintTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolReleasedOrMinted)
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
		it.Event = new(BurnToAddressMintTokenPoolReleasedOrMinted)
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

func (it *BurnToAddressMintTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolReleasedOrMintedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolReleasedOrMinted)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*BurnToAddressMintTokenPoolReleasedOrMinted, error) {
	event := new(BurnToAddressMintTokenPoolReleasedOrMinted)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolRemotePoolAddedIterator struct {
	Event *BurnToAddressMintTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRemotePoolAdded)
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
		it.Event = new(BurnToAddressMintTokenPoolRemotePoolAdded)
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

func (it *BurnToAddressMintTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRemotePoolAddedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRemotePoolAdded)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*BurnToAddressMintTokenPoolRemotePoolAdded, error) {
	event := new(BurnToAddressMintTokenPoolRemotePoolAdded)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolRemotePoolRemovedIterator struct {
	Event *BurnToAddressMintTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolRemotePoolRemoved)
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
		it.Event = new(BurnToAddressMintTokenPoolRemotePoolRemoved)
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

func (it *BurnToAddressMintTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolRemotePoolRemovedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolRemotePoolRemoved)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*BurnToAddressMintTokenPoolRemotePoolRemoved, error) {
	event := new(BurnToAddressMintTokenPoolRemotePoolRemoved)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *BurnToAddressMintTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

func (BurnToAddressMintTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnToAddressMintTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnToAddressMintTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (BurnToAddressMintTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnToAddressMintTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (BurnToAddressMintTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnToAddressMintTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x401d6d5b7a131da25eadbfe5bc7b2c80d7a0094d8d026d2d2c452f81e9b2af7e")
}

func (BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (BurnToAddressMintTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (BurnToAddressMintTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnToAddressMintTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnToAddressMintTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnToAddressMintTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnToAddressMintTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnToAddressMintTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (BurnToAddressMintTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (BurnToAddressMintTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnToAddressMintTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnToAddressMintTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPool) Address() common.Address {
	return _BurnToAddressMintTokenPool.address
}

type BurnToAddressMintTokenPoolInterface interface {
	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetBurnAddress(opts *bind.CallOpts) (common.Address, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, finality uint16, arg4 []byte) (GetFee,

		error)

	GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error)

	IBurnAddress(opts *bind.CallOpts) (common.Address, error)

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

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnToAddressMintTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnToAddressMintTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*BurnToAddressMintTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnToAddressMintTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*BurnToAddressMintTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnToAddressMintTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnToAddressMintTokenPoolConfigChanged, error)

	FilterCustomFinalityMinimumBlockConfirmationUpdated(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdatedIterator, error)

	WatchCustomFinalityMinimumBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomFinalityMinimumBlockConfirmationUpdated(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityMinimumBlockConfirmationUpdated, error)

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnToAddressMintTokenPoolDynamicConfigSet, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnToAddressMintTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnToAddressMintTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnToAddressMintTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnToAddressMintTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnToAddressMintTokenPoolOwnershipTransferred, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnToAddressMintTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*BurnToAddressMintTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*BurnToAddressMintTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnToAddressMintTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnToAddressMintTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnToAddressMintTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnToAddressMintTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
