// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_token_pool

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
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	FeeUSDCents       uint32
	IsEnabled         bool
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

var BurnMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103be5761618e803803809161001d828561043d565b833981019060a0818303126103be5780516001600160a01b038116908190036103be5761004c60208301610460565b60408301519091906001600160401b0381116103be5783019380601f860112156103be578451946001600160401b038611610427578560051b906020820196610098604051988961043d565b87526020808801928201019283116103be57602001905b82821061040f575050506100d160806100ca6060860161046e565b940161046e565b9233156103fe57600180546001600160a01b03191633179055811580156103ed575b80156103dc575b6103cb578160209160049360805260c0526040519283809263313ce56760e01b82525afa6000918161038a575b5061035f575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610241575b604051615b6b9081610623823960805181818161142e015281816116a50152818161181301528181611d9e01528181611ff6015281816120fb015281816129cc01528181612bee01528181612d5c01528181612eb0015281816130620152818161369a015281816136ef01528181613808015281816140dd0152614f32015260a0518181816136560152818161490c0152818161498f0152614fbf015260c051818181610bb9015281816114bc01528181611e2c01528181612a5a0152612f40015260e051818181610b730152818161150101528181611e7101526127bc0152f35b6040516020610250818361043d565b60008252600036813760e0511561034e5760005b82518110156102cb576001906001600160a01b036102828286610482565b51168361028e826104c4565b61029b575b505001610264565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13883610293565b50905060005b8251811015610345576001906001600160a01b036102ef8286610482565b5116801561033f5783610301826105c2565b61030f575b50505b016102d1565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13883610306565b50610309565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff8216818103610373575061012d565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103c3575b816103a66020938361043d565b810103126103be576103b790610460565b9038610127565b600080fd5b3d9150610399565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100fa565b506001600160a01b038416156100f3565b639b15e16f60e01b60005260046000fd5b6020809161041c8461046e565b8152019101906100af565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761042757604052565b519060ff821682036103be57565b51906001600160a01b03821682036103be57565b80518210156104965760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104965760005260206000200190600090565b60008181526003602052604090205480156105bb5760001981018181116105a5576002546000198101919082116105a557818103610554575b505050600254801561053e57600019016105188160026104ac565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61058d6105656105769360026104ac565b90549060031b1c92839260026104ac565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806104fd565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461061c57600254680100000000000000008110156104275761060361057682600185940160025560026104ac565b9055600254906000526003602052604060002055600190565b5060009056fe60a080604052600436101561001357600080fd5b60006080526080513560e01c90816301ffc9a71461399357508063164e68de14613775578063181f5a771461371357806321df0da7146136ce578063240028e81461367a57806324f65ee71461363b5780632a10097b146133e15780632c286daf146132d757806337b19247146131c65780633907753714612e29578063489a68f2146129335780634c5ef0ed146128ee57806354c8a4f31461278a5780635df45a371461276657806362ddd3c4146126e1578063698c2c661461262f5780636d3d1a58146126075780637437ff9f146125d357806379ba5097146125165780637d54534e14612492578063804ba5a91461242a5780638926f54f146123e557806389720a621461237a5780638da5cb5b14612352578063962d40201461220f5780639a4575b914611d4c578063a42a7b8b14611bf5578063a7cd63b714611b87578063acfecf9114611a5a578063af58d59f14611a0e578063b1c71c65146113b1578063b794658014611379578063c4bffe2b1461125d578063c75eea9c146111b2578063cf7401f314611025578063d966866b14610bdd578063dc0bd97114610b98578063e0351e1314610b5a578063e8a1da17146102a25763f2fde38b146101dd57600080fd5b3461029c57602060031936011261029c576001600160a01b036101fe613ac4565b610206614517565b163381146102705760805180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360805180f35b7fdad89dca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b60805180fd5b3461029c576102b036613d5e565b9190926102bb614517565b608051905b828210610998575050506080519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b8181101561099257600581901b8301358581121561029c5783016101208136031261029c576040519461032f86613b71565b813567ffffffffffffffff8116810361098d578652602082013567ffffffffffffffff811161029c5782019436601f8701121561029c57853561037181614035565b9661037f6040519889613ba9565b81885260208089019260051b8201019036821161029c5760208101925b82841061095e575050505060208701958652604083013567ffffffffffffffff811161029c576103cf9036908501613d40565b91604088019283526103f96103e73660608701613eb9565b9460608a0195865260c0369101613eb9565b956080890196875261040b8551615128565b6104158751615128565b835151156109325761043167ffffffffffffffff8a51166159e3565b156108f75767ffffffffffffffff89511660805152600860205260406080512061057586516fffffffffffffffffffffffffffffffff604082015116906105306fffffffffffffffffffffffffffffffff6020830151169151151583608060405161049b81613b71565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b61069b88516fffffffffffffffffffffffffffffffff604082015116906106566fffffffffffffffffffffffffffffffff602083015116915115158360806040516105bf81613b71565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff82116108c6576106be83546141ff565b601f8111610887575b506020906001601f84111461081d579180916106fa9360805192610812575b50506000198260011b9260031b1c19161790565b90555b6080515b8851805182101561073657906107306001926107298367ffffffffffffffff8f5116926141eb565b5190614a93565b01610701565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561080467ffffffffffffffff60019796949851169251935191516107d061079b60405196879687526101006020880152610100870190613be8565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a10193929091936102fd565b015190508e806106e6565b90601f1983169184608051528160805120926080515b81811061086f5750908460019594939210610856575b505050811b0190556106fd565b015160001960f88460031b161c191690558d8080610849565b92936020600181928786015181550195019301610833565b6108b6908460805152602060805120601f850160051c810191602086106108bc575b601f0160051c01906144be565b8d6106c7565b90915081906108a9565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b67ffffffffffffffff8951167f1d5ad3c500000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f14c880ca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b833567ffffffffffffffff811161029c576020916109828392833691870101613d40565b81520193019261039c565b600080fd5b60805180f35b9092919367ffffffffffffffff6109b86109b3868886613f83565b613f3f565b16926109c384615814565b15610b2a57836080515260086020526109e36005604060805120016156b1565b926080515b8451811015610a225760019086608051526008602052610a1b600560406080512001610a1483896141eb565b51906158c7565b50016109e8565b5093909491959250806080515260086020526005604060805120608051815560805160018201556080516002820155608051600382015560048101610a6781546141ff565b80610ada575b505001805490608051815581610ab6575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102c0565b60805152602060805120908101905b81811015610a7e576080518155600101610ac5565b601f8111600114610af3575060805190555b8880610a6d565b610b139082608051526001601f6020608051209201861c820191016144be565b608080518290525160208120918190559055610aec565b837f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b3461029c5760805160031936011261029c5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461029c5760805160031936011261029c5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461029c57602060031936011261029c5760043567ffffffffffffffff811161029c57610c0e903690600401613c29565b610c16614517565b608051915b818310610c285760805180f35b610c366109b38484846143e8565b610c4e610c448585856143e8565b6020810190614428565b9091610c68610c5e8787876143e8565b6040810190614428565b90610c81610c778989896143e8565b6060810190614428565b9091610c9b610c918b8b8b6143e8565b6080810190614428565b949097610cb1610cac368a8461404d565b615085565b610cbf610cac36848661404d565b610ccd610cac36868861404d565b610cdb610cac36888c61404d565b604051610ce781613aee565b610cf2368a8461404d565b8152610cff36848661404d565b6020820152610d0f36868861404d565b6040820152610d1f36888c61404d565b606082015267ffffffffffffffff881660805152600e602052604060805120815180519067ffffffffffffffff82116108c6576801000000000000000082116108c6576020908354838555808410611006575b500182608051526020608051206080515b838110610fe95750505050602082015180519067ffffffffffffffff82116108c6576801000000000000000082116108c6576020906001840154836001860155808410610fc7575b500160018301608051526020608051206080515b838110610faa5750505050604082015180519067ffffffffffffffff82116108c6576801000000000000000082116108c6576020906002840154836002860155808410610f88575b500160028301608051526020608051206080515b838110610f6b575050505060036060919e9c9d9e019101519081519167ffffffffffffffff83116108c6576801000000000000000083116108c6576020908254848455808510610f4c575b500190608051526020608051206080515b838110610f2f5750505050610f14608095610f2495610f067fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a96610ef867ffffffffffffffff976040519d8d8f9e8f90815201916144d5565b918b830360208d01526144d5565b9188830360408a01526144d5565b92858403606087015216966144d5565b0390a2019190610c1b565b60019060206001600160a01b038551169401938184015501610e97565b610f6590846080515285846080512091820191016144be565b38610e86565b60019060206001600160a01b038551169401938184015501610e3b565b610fa490600286016080515284846080512091820191016144be565b38610e27565b60019060206001600160a01b038551169401938184015501610ddf565b610fe390600186016080515284846080512091820191016144be565b38610dcb565b60019060206001600160a01b038551169401938184015501610d83565b61101f90856080515284846080512091820191016144be565b38610d72565b3461029c5760e060031936011261029c5761103e613cc4565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc011261029c5760405161107481613b8d565b602435801515810361029c5781526044356fffffffffffffffffffffffffffffffff8116810361029c5760208201526064356fffffffffffffffffffffffffffffffff8116810361029c5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c011261029c57604051906110fb82613b8d565b608435801515810361029c57825260a4356fffffffffffffffffffffffffffffffff8116810361029c57602083015260c4356fffffffffffffffffffffffffffffffff8116810361029c5760408301526001600160a01b03600a54163314158061119d575b61116d5761099292614df8565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b506001600160a01b0360015416331415611160565b3461029c57602060031936011261029c5767ffffffffffffffff6111d4613cc4565b6111dc614335565b50166080515260086020526112596112006111fb604060805120614360565b615000565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b3461029c5760805160031936011261029c57608051506040516006548082528160208101600660805152602060805120926080515b8181106113605750506112a792500382613ba9565b8051906112cc6112b683614035565b926112c46040519485613ba9565b808452614035565b90601f196020840192013683376080515b815181101561130f578067ffffffffffffffff6112fc600193856141eb565b511661130882876141eb565b52016112dd565b505090604051918291602083019060208452518091526040830191906080515b81811061133d575050500390f35b825167ffffffffffffffff1684528594506020938401939092019160010161132f565b8454835260019485019486945060209093019201611292565b3461029c57602060031936011261029c5761125961139d611398613cc4565b6143c6565b604051918291602083526020830190613be8565b3461029c57606060031936011261029c5760043567ffffffffffffffff811161029c5760a0600319823603011261029c576113ea613c5a565b9060443567ffffffffffffffff811161029c5761140b903690600401613d40565b506114146141d2565b50608481019061142382613fe4565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036119cd57602481019077ffffffffffffffff0000000000000000000000000000000061147c83613f3f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118e3576080519161199e575b50611972576114ff60448201613fe4565b7f000000000000000000000000000000000000000000000000000000000000000061191f575b5067ffffffffffffffff61153883613f3f565b16611550816000526007602052604060002054151590565b156118f05760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156118e3576080519061189a575b6001600160a01b03915016330361186a57606461ffff9101359316928315159283809461185b575b156117ba5761ffff600b541694858110611786575061174a94506116216116116115f785613f3f565b67ffffffffffffffff16600052600c602052604060002090565b8361161b84613fe4565b916154a2565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff61165d61165786613f3f565b93613fe4565b604080516001600160a01b03929092168252602082018690529190931692a25b918290611754575b506113986117199161169684614f28565b61169f81613f3f565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2613f3f565b90611722614fb8565b6040519261172f84613b55565b83526020830152604051928392604084526040840190613e65565b9060208301520390f35b61171991925061177e6113989161271061177761ffff600b5460101c168361447c565b0490614ff3565b929150611685565b85907fe08f03ef00000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b5061174a935067ffffffffffffffff6117d283613f3f565b16806080515260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944828061183b6040608051206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916154a2565b604080516001600160a01b039290921682526020820192909252a261167d565b5061ffff600b541615156115ce565b7f728fe07b0000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b506020813d6020116118db575b816118b460209383613ba9565b8101031261029c57516001600160a01b038116810361029c576001600160a01b03906115a6565b3d91506118a7565b6040513d608051823e3d90fd5b7fa9902c7e00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b6001600160a01b031661193f816000526003602052604060002054151590565b611525577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f53ad11d800000000000000000000000000000000000000000000000000000000608051526004608051fd5b6119c0915060203d6020116119c6575b6119b88183613ba9565b810190614a7b565b856114ee565b503d6119ae565b6001600160a01b036119de83613fe4565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b3461029c57602060031936011261029c5767ffffffffffffffff611a30613cc4565b611a38614335565b50166080515260086020526112596112006111fb600260406080512001614360565b3461029c5767ffffffffffffffff611a7136613db0565b929091611a7c614517565b1690611a95826000526007602052604060002054151590565b15611b575781608051526008602052611ac8600560406080512001611abb368685613d09565b60208151910120906158c7565b15611b10577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611b07604051928392602084526020840191614314565b0390a260805180f35b611b53906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614314565b0390fd5b507f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b3461029c5760805160031936011261029c57608051506040516002548082526020820190600260805152602060805120906080515b818110611bdf5761125985611bd381870382613ba9565b60405191829182613df1565b8254845260209093019260019283019201611bbc565b3461029c57602060031936011261029c5767ffffffffffffffff611c17613cc4565b16608051526008602052611c326005604060805120016156b1565b805190601f19611c5a611c4484614035565b93611c526040519586613ba9565b808552614035565b016080515b818110611d3b5750506080515b8151811015611cb65780611c82600192846141eb565b51608051526009602052611c9a604060805120614252565b611ca482866141eb565b52611caf81856141eb565b5001611c6c565b826040518091602082016020835281518091526040830190602060408260051b860101930191608051905b828210611cf057505050500390f35b91936020611d2b827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613be8565b9601920192018594939192611ce1565b806060602080938701015201611c5f565b3461029c57602060031936011261029c5760043567ffffffffffffffff811161029c5760a0600319823603011261029c57611d856141d2565b5060848101611d9381613fe4565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036121fe57602482019177ffffffffffffffff00000000000000000000000000000000611dec84613f3f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118e357608051916121df575b5061197257611e6f60448201613fe4565b7f000000000000000000000000000000000000000000000000000000000000000061218c575b5067ffffffffffffffff611ea884613f3f565b16611ec0816000526007602052604060002054151590565b156118f05760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156118e35760805190612143575b6001600160a01b03915016330361186a57606401359060805160001461209e5761ffff600b5416806120695750826120399261139892611f5e6116116115f761125998613f3f565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff611f9461165786613f3f565b604080516001600160a01b03929092168252602082018690529190931692a25b611fbd81614f28565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff611ff084613f3f565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681523360208201529081019490945216918060608101611711565b612041614fb8565b6040519161204e83613b55565b82526020820152604051918291602083526020830190613e65565b7fe08f03ef00000000000000000000000000000000000000000000000000000000608051526080516004526024526044608051fd5b50611398826120399267ffffffffffffffff6120bc61125996613f3f565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944828061212360406000206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916154a2565b604080516001600160a01b039290921682526020820192909252a2611fb4565b506020813d602011612184575b8161215d60209383613ba9565b8101031261029c57516001600160a01b038116810361029c576001600160a01b0390611f16565b3d9150612150565b6001600160a01b03166121ac816000526003602052604060002054151590565b611e95577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b6121f8915060203d6020116119c6576119b88183613ba9565b84611e5e565b6119de6001600160a01b0391613fe4565b3461029c57606060031936011261029c5760043567ffffffffffffffff811161029c57612240903690600401613c29565b9060243567ffffffffffffffff811161029c57612261903690600401613e34565b9060443567ffffffffffffffff811161029c57612282903690600401613e34565b6001600160a01b03600a54163314158061233d575b61116d57838614801590612333575b612307576080515b8681106122bb5760805180f35b806123016122cf6109b36001948b8b613f83565b6122da8389896141c2565b6122fb6122f36122eb86898b6141c2565b923690613eb9565b913690613eb9565b91614df8565b016122ae565b7f568efce200000000000000000000000000000000000000000000000000000000608051526004608051fd5b50808614156122a6565b506001600160a01b0360015416331415612297565b3461029c5760805160031936011261029c5760206001600160a01b0360015416604051908152f35b3461029c5760c060031936011261029c57612393613ac4565b5061239c613cad565b6123a4613c6b565b5060843567ffffffffffffffff811161029c576123c5903690600401613cdb565b505060a43590600282101561029c5761125991611bd3916044359061414c565b3461029c57602060031936011261029c57602061242067ffffffffffffffff61240c613cc4565b166000526007602052604060002054151590565b6040519015158152f35b3461029c57602060031936011261029c5760043567ffffffffffffffff811161029c5761245b903690600401613c7c565b6001600160a01b03600a54163314158061247d575b61116d5761099291614572565b506001600160a01b0360015416331415612470565b3461029c57602060031936011261029c577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b036124d6613ac4565b6124de614517565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a160805180f35b3461029c5760805160031936011261029c57608051546001600160a01b03811633036125a7577fffffffffffffffffffffffff000000000000000000000000000000000000000060015491338284161760015516608051556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0608051608051a360805180f35b7f02b543c600000000000000000000000000000000000000000000000000000000608051526004608051fd5b3461029c5760805160031936011261029c57600454600554604080516001600160a01b039093168352602083019190915290f35b3461029c5760805160031936011261029c5760206001600160a01b03600a5416604051908152f35b3461029c57604060031936011261029c57612648613ac4565b602435612653614517565b6001600160a01b038216918215610932577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff00000000000000000000000000000000000000006004541617600455816005556126d860405192839283602090939291936001600160a01b0360408201951681520152565b0390a160805180f35b3461029c576126ef36613db0565b6126fa929192614517565b67ffffffffffffffff821661271c816000526007602052604060002054151590565b15612737575061099292612731913691613d09565b90614a93565b7f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b3461029c5760805160031936011261029c5760206127826140a1565b604051908152f35b3461029c576127b26127ba61279e36613d5e565b94916127ab939193614517565b369161404d565b92369161404d565b7f0000000000000000000000000000000000000000000000000000000000000000156128c2576080515b825181101561284a57806001600160a01b03612802600193866141eb565b511661280d81615714565b612819575b50016127e4565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612812565b506080515b815181101561099257806001600160a01b0361286d600193856141eb565b511680156128bc5761287e81615983565b61288b575b505b0161284f565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612883565b50612885565b7f35f4a7b300000000000000000000000000000000000000000000000000000000608051526004608051fd5b3461029c57604060031936011261029c57612907613cc4565b60243567ffffffffffffffff811161029c5760209161292d612420923690600401613d40565b90613ff8565b3461029c57604060031936011261029c5760043567ffffffffffffffff811161029c578060040190610100600319823603011261029c57612972613c5a565b60405190919061298181613b39565b60805190526129b26129a86129a361299c60c4850187613f93565b3691613d09565b614898565b606483013561498c565b9160848201906129c182613fe4565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036119cd57602483019477ffffffffffffffff00000000000000000000000000000000612a1a87613f3f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118e35760805191612e0a575b506119725767ffffffffffffffff612aa387613f3f565b16612abb816000526007602052604060002054151590565b156118f05760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156118e35760805191612deb575b501561186a57612b2786613f3f565b90612b3d60a486019261292d61299c8585613f93565b15612da45750604493929161ffff16159050612d055767ffffffffffffffff612b6586613f3f565b1660805152600d602052612b826040608051208561161b84613fe4565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff767ffffffffffffffff612bb861165788613f3f565b604080516001600160a01b03929092168252602082018890529190931692a25b0191612be383613fe4565b906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561029c576040517f40c10f190000000000000000000000000000000000000000000000000000000081526080516001600160a01b03909216600482015260248101859052908180604481010381608051875af180156118e357612cec575b50608067ffffffffffffffff6020956001600160a01b03612cba612cb47ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096613f3f565b92613fe4565b60405196875233898801521660408601528560608601521692a260405190612ce182613b39565b815260405190518152f35b608051612cf891613ba9565b60805161029c5784612c70565b5067ffffffffffffffff612d1885613f3f565b16806080515260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8480612d846002604060805120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916154a2565b604080516001600160a01b039290921682526020820192909252a2612bd8565b612dae9250613f93565b611b536040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614314565b612e04915060203d6020116119c6576119b88183613ba9565b87612b18565b612e23915060203d6020116119c6576119b88183613ba9565b87612a8c565b3461029c57602060031936011261029c5760043567ffffffffffffffff811161029c578060040190610100600319823603011261029c57604051612e6c81613b39565b6080519052604051612e7d81613b39565b6080519052612e986129a86129a361299c60c4850186613f93565b9060848101612ea681613fe4565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116036121fe5750602481019277ffffffffffffffff00000000000000000000000000000000612f0085613f3f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156118e357608051916131a7575b506119725767ffffffffffffffff612f8985613f3f565b16612fa1816000526007602052604060002054151590565b156118f05760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156118e35760805191613188575b501561186a5761300d84613f3f565b9061302360a484019261292d61299c8585613f93565b15612da4575082916044915067ffffffffffffffff61304186613f3f565b168060805152600860205261308a6002604060805120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169586916154a2565b604080516001600160a01b0386168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a201926130d084613fe4565b823b1561029c576040517f40c10f190000000000000000000000000000000000000000000000000000000081526080516001600160a01b03909216600482015260248101859052908180604481010381608051875af180156118e3576020956001600160a01b03612cba612cb47ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09660809667ffffffffffffffff96613177575b50613f3f565b875161318291613ba9565b8b613171565b6131a1915060203d6020116119c6576119b88183613ba9565b85612ffe565b6131c0915060203d6020116119c6576119b88183613ba9565b85612f72565b3461029c5760a060031936011261029c576131df613ac4565b506131e8613cad565b60443567ffffffffffffffff811161029c5760031960a0913603011261029c57613210613c6b565b506084359067ffffffffffffffff821161029c5761323b67ffffffffffffffff923690600401613cdb565b505060405161324981613aee565b60805181526080516020820152608051604082015260606080519101521660805152600f602052608060408151206040519061328482613aee565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b3461029c57606060031936011261029c5760043561ffff81169081900361029c57613300613c5a565b9060443567ffffffffffffffff811161029c57613321903690600401613c7c565b9061332a614517565b61ffff8416936127108510156133b15783927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba95926133a0926040967fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600b549360101b1692161717600b55614572565b82519182526020820152a160805180f35b847f95f3517a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b3461029c57604060031936011261029c5760043567ffffffffffffffff811161029c573660238201121561029c57806004013567ffffffffffffffff811161029c5760248201916024369160a0840201011161029c5760243567ffffffffffffffff811161029c57613457903690600401613c29565b919092613462614517565b6080515b8281106134de575050506080515b8181106134815760805180f35b8067ffffffffffffffff61349b6109b36001948688613f83565b168060805152600f602052608051604060805120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8608051608051a201613474565b806134ef6109b36001938686613f00565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70608061351d848888613f00565b9261362d6135f363ffffffff6136226135e68261361767ffffffffffffffff60208c0198169a8b8a5152600f60205260408a51208361355b8b613f54565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff0000000061359887613f54565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e613f54565b60401b1696019e8f613f65565b151560601b169516171617171790558261360f6040519a613f72565b168952613f72565b166020870152613f72565b166040840152613e8f565b15156060820152a201613466565b3461029c5760805160031936011261029c57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461029c57602060031936011261029c576020613695613ac4565b6040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081169216919091148152f35b3461029c5760805160031936011261029c5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461029c5760805160031936011261029c5760408051611259916137379082613ba9565b601b81527f4275726e4d696e74546f6b656e506f6f6c20312e372e302d64657600000000006020820152604051918291602083526020830190613be8565b3461029c57602060031936011261029c5761378e613ac4565b613796614517565b61379e6140a1565b90816137aa5760805180f35b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082019081526001600160a01b038316602483015260448083018590528252906138a5906137fe606482613ba9565b60408051909390917f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316919061383c8685613ba9565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260805191608051915190608051855af13d1561398b573d9161388783613bcc565b9261389487519485613ba9565b83526080513d90602085013e615a92565b8051806138ea575b50506001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992602092519485521692a28080610992565b906020806138fc938301019101614a7b565b156139085783806138ad565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615a92565b3461029c57602060031936011261029c57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361029c57817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115613a9a575b8115613a70575b8115613a46575b8115613a1c575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613a15565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613a0e565b7ff00a89320000000000000000000000000000000000000000000000000000000081149150613a07565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613a00565b600435906001600160a01b038216820361098d57565b35906001600160a01b038216820361098d57565b6080810190811067ffffffffffffffff821117613b0a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613b0a57604052565b6040810190811067ffffffffffffffff821117613b0a57604052565b60a0810190811067ffffffffffffffff821117613b0a57604052565b6060810190811067ffffffffffffffff821117613b0a57604052565b90601f601f19910116810190811067ffffffffffffffff821117613b0a57604052565b67ffffffffffffffff8111613b0a57601f01601f191660200190565b919082519283825260005b848110613c14575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201613bf3565b9181601f8401121561098d5782359167ffffffffffffffff831161098d576020808501948460051b01011161098d57565b6024359061ffff8216820361098d57565b6064359061ffff8216820361098d57565b9181601f8401121561098d5782359167ffffffffffffffff831161098d5760208085019460e0850201011161098d57565b6024359067ffffffffffffffff8216820361098d57565b6004359067ffffffffffffffff8216820361098d57565b9181601f8401121561098d5782359167ffffffffffffffff831161098d576020838186019501011161098d57565b929192613d1582613bcc565b91613d236040519384613ba9565b82948184528183011161098d578281602093846000960137010152565b9080601f8301121561098d57816020613d5b93359101613d09565b90565b604060031982011261098d5760043567ffffffffffffffff811161098d5781613d8991600401613c29565b929092916024359067ffffffffffffffff821161098d57613dac91600401613c29565b9091565b90604060031983011261098d5760043567ffffffffffffffff8116810361098d57916024359067ffffffffffffffff821161098d57613dac91600401613cdb565b602060408183019282815284518094520192019060005b818110613e155750505090565b82516001600160a01b0316845260209384019390920191600101613e08565b9181601f8401121561098d5782359167ffffffffffffffff831161098d576020808501946060850201011161098d57565b613d5b916020613e7e8351604084526040840190613be8565b920151906020818403910152613be8565b3590811515820361098d57565b35906fffffffffffffffffffffffffffffffff8216820361098d57565b919082606091031261098d57604051613ed181613b8d565b6040613efb818395613ee281613e8f565b8552613ef060208201613e9c565b602086015201613e9c565b910152565b9190811015613f105760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff8116810361098d5790565b3563ffffffff8116810361098d5790565b35801515810361098d5790565b359063ffffffff8216820361098d57565b9190811015613f105760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561098d570180359067ffffffffffffffff821161098d5760200191813603831361098d57565b356001600160a01b038116810361098d5790565b9067ffffffffffffffff613d5b92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613b0a5760051b60200190565b92919061405981614035565b936140676040519586613ba9565b602085838152019160051b810192831161098d57905b82821061408957505050565b6020809161409684613ada565b81520191019061407d565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561414057600091614111575090565b90506020813d602011614138575b8161412c60209383613ba9565b8101031261098d575190565b3d915061411f565b6040513d6000823e3d90fd5b67ffffffffffffffff16600052600e60205260406000209160028110156141935760011461418257816001613d5b930190614d04565b8160026003613d5b94019101614d04565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9190811015613f10576060020190565b604051906141df82613b55565b60606020838281520152565b8051821015613f105760209160051b010190565b90600182811c92168015614248575b602083101461421957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f169161420e565b9060405191826000825492614266846141ff565b80845293600181169081156142d4575060011461428d575b5061428b92500383613ba9565b565b90506000929192526020600020906000915b8183106142b857505090602061428b928201013861427e565b602091935080600191548385890101520191019091849261429f565b6020935061428b9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b8201013861427e565b601f8260209493601f19938186528686013760008582860101520116010190565b6040519061434282613b71565b60006080838281528260208201528260408201528260608201520152565b9060405161436d81613b71565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526008602052613d5b6004604060002001614252565b9190811015613f105760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561098d570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561098d570180359067ffffffffffffffff821161098d57602001918160051b3603831361098d57565b8181029291811591840414171561448f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106144c9575050565b600081556001016144be565b9160209082815201919060005b8181106144ef5750505090565b9091926020806001926001600160a01b0361450988613ada565b1681520194019291016144e2565b6001600160a01b0360015416330361452b57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b356fffffffffffffffffffffffffffffffff8116810361098d5790565b9160005b828110156148925760e081028401600061458f82613f3f565b9067ffffffffffffffff8216916145b3836000526007602052604060002054151590565b156148665761467c9260408593614627614621946146216145e7602060019c9b01926115f76145e23686613eb9565b615128565b91825463ffffffff8160801c16159081614848575b81614839575b8161481e575b8161480f575b5080614800575b614775575b3690613eb9565b9061526f565b60808501926146396145e23686613eb9565b8152600d6020522092835463ffffffff8160801c16159081614757575b81614748575b8161472d575b8161471e575b508061470f575b614682575b503690613eb9565b01614576565b61469f60a06fffffffffffffffffffffffffffffffff9201614555565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614674565b5061471982613f65565b61466f565b60ff915060a01c161538614668565b6fffffffffffffffffffffffffffffffff8116159150614662565b8589015460801c15915061465c565b858901546fffffffffffffffffffffffffffffffff16159150614656565b6fffffffffffffffffffffffffffffffff614791878b01614555565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835561461a565b5061480a81613f65565b614615565b60ff915060a01c16153861460e565b6fffffffffffffffffffffffffffffffff8116159150614608565b848e015460801c159150614602565b848e01546fffffffffffffffffffffffffffffffff161591506145fc565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b80518015614908576020036148ca57805160208281019183018390031261098d57519060ff82116148ca575060ff1690565b611b53906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613be8565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161448f57565b60ff16604d811161448f57600a0a90565b811561495d570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614a7457828411614a4a57906149d19161492e565b91604d60ff8416118015614a2f575b6149f9575050906149f3613d5b92614942565b9061447c565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614a3983614942565b801561495d576000190484116149e0565b614a539161492e565b91604d60ff8416116149f957505090614a6e613d5b92614942565b90614953565b5050505090565b9081602091031261098d5751801515810361098d5790565b90805115614c795767ffffffffffffffff81516020830120921691826000526008602052614ac8816005604060002001615a3d565b15614c355760005260096020526040600020815167ffffffffffffffff8111613b0a57614af582546141ff565b601f8111614c03575b506020601f8211600114614b795791614b53827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614b6995600091614b6e575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613be8565b0390a2565b905084015138614b40565b601f1982169083600052806000209160005b818110614beb575092614b699492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614bd2575b5050811b01905561139d565b85015160001960f88460031b161c191690553880614bc6565b9192602060018192868a015181550194019201614b8b565b614c2f90836000526020600020601f840160051c810191602085106108bc57601f0160051c01906144be565b38614afe565b5090611b536040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613be8565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b818110614cd557505061428b92500383613ba9565b84546001600160a01b0316835260019485019487945060209093019201614cc0565b9190820180921161448f57565b614d0d90614ca3565b916005548015159182614ded575b5050614d25575090565b614d2e90614ca3565b90815180614d3c5750905090565b614d47908251614cf7565b92601f19614d6d614d5786614035565b95614d656040519788613ba9565b808752614035565b0136602086013760005b8251811015614da857806001600160a01b03614d95600193866141eb565b5116614da182886141eb565b5201614d77565b509160005b8151811015614de857806001600160a01b03614dcb600193856141eb565b5116614de1614ddb838751614cf7565b886141eb565b5201614dad565b505050565b101590503880614d1b565b67ffffffffffffffff166000818152600760205260409020549092919015614efa5791614ef760e092614ec385614e4f7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97615128565b846000526008602052614e6681604060002061526f565b614e6f83615128565b846000526008602052614e8983600260406000200161526f565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561098d57604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af18015614fad57614fa0575050565b81614faa91613ba9565b50565b6040513d84823e3d90fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613d5b604082613ba9565b9190820391821161448f57565b615008614335565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691615065602085019361505f61505263ffffffff87511642614ff3565b856080890151169061447c565b90614cf7565b8082101561507e57505b16825263ffffffff4216905290565b905061506f565b805160005b81811061509657505050565b6001810180821161448f575b8281106150b2575060010161508a565b6001600160a01b036150c483866141eb565b51166001600160a01b036150d883876141eb565b5116146150e7576001016150a2565b6001600160a01b036150f983866141eb565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8051156151c8576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106151655750565b6064906151c6604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590615250575b6151ef5750565b6064906151c6604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208201511615156151e8565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916153a860609280546152ac63ffffffff8260801c1642614ff3565b90816153e7575b50506fffffffffffffffffffffffffffffffff60018160208601511692828154168085106000146153df57508280855b16167fffffffffffffffffffffffffffffffff0000000000000000000000000000000082541617815561535c8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b614ef760405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916152e3565b6fffffffffffffffffffffffffffffffff9161541c8392836154156001880154948286169560801c9061447c565b9116614cf7565b8082101561549b57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff000000000000000000000000000000001617815538806152b3565b9050615426565b9182549060ff8260a01c161580156156a9575b6156a3576fffffffffffffffffffffffffffffffff821691600185019081546154fa63ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642614ff3565b9081615605575b50508481106155c6575083831061555b5750506155306fffffffffffffffffffffffffffffffff928392614ff3565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c9161556a8185614ff3565b9260001981019080821161448f5761558d615592926001600160a01b0396614cf7565b614953565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615679576156209261505f9160801c9061447c565b808410156156745750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615501565b61562b565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156154b5565b906040519182815491828252602082019060005260206000209260005b8181106156e357505061428b92500383613ba9565b84548352600194850194879450602090930192016156ce565b8054821015613f105760005260206000200190600090565b600081815260036020526040902054801561580d57600019810181811161448f5760025490600019820191821161448f578181036157bc575b505050600254801561578d57600019016157688160026156fc565b60001982549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6157f56157cd6157de9360026156fc565b90549060031b1c92839260026156fc565b81939154906000199060031b92831b921b19161790565b9055600052600360205260406000205538808061574d565b5050600090565b600081815260076020526040902054801561580d57600019810181811161448f5760065490600019820191821161448f5781810361588d575b505050600654801561578d57600019016158688160066156fc565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b6158af61589e6157de9360066156fc565b90549060031b1c92839260066156fc565b9055600052600760205260406000205538808061584d565b906001820191816000528260205260406000205480151560001461597a57600019810181811161448f57825490600019820191821161448f57818103615943575b5050508054801561578d57600019019061592282826156fc565b60001982549160031b1b191690555560005260205260006040812055600190565b6159636159536157de93866156fc565b90549060031b1c928392866156fc565b905560005283602052604060002055388080615908565b50505050600090565b806000526003602052604060002054156000146159dd5760025468010000000000000000811015613b0a576159c46157de82600185940160025560026156fc565b9055600254906000526003602052604060002055600190565b50600090565b806000526007602052604060002054156000146159dd5760065468010000000000000000811015613b0a57615a246157de82600185940160065560066156fc565b9055600654906000526007602052604060002055600190565b600082815260018201602052604090205461580d5780549068010000000000000000821015613b0a5782615a7b6157de8460018096018555846156fc565b905580549260005201602052604060002055600190565b91929015615b0d5750815115615aa6575090565b3b15615aaf5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615b205750805190602001fd5b611b53906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613be856fea164736f6c634300081a000a",
}

var BurnMintTokenPoolABI = BurnMintTokenPoolMetaData.ABI

var BurnMintTokenPoolBin = BurnMintTokenPoolMetaData.Bin

func DeployBurnMintTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnMintTokenPool, error) {
	parsed, err := BurnMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintTokenPool{address: address, abi: *parsed, BurnMintTokenPoolCaller: BurnMintTokenPoolCaller{contract: contract}, BurnMintTokenPoolTransactor: BurnMintTokenPoolTransactor{contract: contract}, BurnMintTokenPoolFilterer: BurnMintTokenPoolFilterer{contract: contract}}, nil
}

type BurnMintTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnMintTokenPoolCaller
	BurnMintTokenPoolTransactor
	BurnMintTokenPoolFilterer
}

type BurnMintTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolSession struct {
	Contract     *BurnMintTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintTokenPoolCallerSession struct {
	Contract *BurnMintTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnMintTokenPoolTransactorSession struct {
	Contract     *BurnMintTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintTokenPoolRaw struct {
	Contract *BurnMintTokenPool
}

type BurnMintTokenPoolCallerRaw struct {
	Contract *BurnMintTokenPoolCaller
}

type BurnMintTokenPoolTransactorRaw struct {
	Contract *BurnMintTokenPoolTransactor
}

func NewBurnMintTokenPool(address common.Address, backend bind.ContractBackend) (*BurnMintTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPool{address: address, abi: abi, BurnMintTokenPoolCaller: BurnMintTokenPoolCaller{contract: contract}, BurnMintTokenPoolTransactor: BurnMintTokenPoolTransactor{contract: contract}, BurnMintTokenPoolFilterer: BurnMintTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnMintTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnMintTokenPoolCaller, error) {
	contract, err := bindBurnMintTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCaller{contract: contract}, nil
}

func NewBurnMintTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintTokenPoolTransactor, error) {
	contract, err := bindBurnMintTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolTransactor{contract: contract}, nil
}

func NewBurnMintTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintTokenPoolFilterer, error) {
	contract, err := bindBurnMintTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolFilterer{contract: contract}, nil
}

func bindBurnMintTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintTokenPool.Contract.BurnMintTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintTokenPool *BurnMintTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.BurnMintTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.BurnMintTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _BurnMintTokenPool.Contract.GetAccumulatedFees(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _BurnMintTokenPool.Contract.GetAccumulatedFees(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetAllowList(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetAllowList(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintTokenPool.Contract.GetAllowListEnabled(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintTokenPool.Contract.GetAllowListEnabled(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnMintTokenPool.Contract.GetDynamicConfig(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _BurnMintTokenPool.Contract.GetDynamicConfig(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRateLimitAdmin(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRateLimitAdmin(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintTokenPool.Contract.GetRemotePools(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintTokenPool.Contract.GetRemotePools(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintTokenPool.Contract.GetRemoteToken(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintTokenPool.Contract.GetRemoteToken(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRequiredCCVs(&_BurnMintTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRequiredCCVs(&_BurnMintTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRmnProxy(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetRmnProxy(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintTokenPool.Contract.GetSupportedChains(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintTokenPool.Contract.GetSupportedChains(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetToken(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnMintTokenPool.Contract.GetToken(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintTokenPool.Contract.GetTokenDecimals(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintTokenPool.Contract.GetTokenDecimals(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintTokenPool.Contract.IsRemotePool(&_BurnMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintTokenPool.Contract.IsRemotePool(&_BurnMintTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintTokenPool.Contract.IsSupportedChain(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintTokenPool.Contract.IsSupportedChain(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintTokenPool.Contract.IsSupportedToken(&_BurnMintTokenPool.CallOpts, token)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintTokenPool.Contract.IsSupportedToken(&_BurnMintTokenPool.CallOpts, token)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) Owner() (common.Address, error) {
	return _BurnMintTokenPool.Contract.Owner(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnMintTokenPool.Contract.Owner(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintTokenPool.Contract.SupportsInterface(&_BurnMintTokenPool.CallOpts, interfaceId)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintTokenPool.Contract.SupportsInterface(&_BurnMintTokenPool.CallOpts, interfaceId)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnMintTokenPool.Contract.TypeAndVersion(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnMintTokenPool.Contract.TypeAndVersion(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.AcceptOwnership(&_BurnMintTokenPool.TransactOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.AcceptOwnership(&_BurnMintTokenPool.TransactOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.AddRemotePool(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.AddRemotePool(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnMintTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnMintTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyChainUpdates(&_BurnMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyChainUpdates(&_BurnMintTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnMintTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnMintTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn0(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn0(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint0(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint0(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.RemoveRemotePool(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.RemoveRemotePool(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnMintTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnMintTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetDynamicConfig(&_BurnMintTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetDynamicConfig(&_BurnMintTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetRateLimitAdmin(&_BurnMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetRateLimitAdmin(&_BurnMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.TransferOwnership(&_BurnMintTokenPool.TransactOpts, to)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.TransferOwnership(&_BurnMintTokenPool.TransactOpts, to)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.WithdrawFees(&_BurnMintTokenPool.TransactOpts, recipient)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.WithdrawFees(&_BurnMintTokenPool.TransactOpts, recipient)
}

type BurnMintTokenPoolAllowListAddIterator struct {
	Event *BurnMintTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAllowListAdd)
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
		it.Event = new(BurnMintTokenPoolAllowListAdd)
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

func (it *BurnMintTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAllowListAddIterator{contract: _BurnMintTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAllowListAdd)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*BurnMintTokenPoolAllowListAdd, error) {
	event := new(BurnMintTokenPoolAllowListAdd)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAllowListRemoveIterator struct {
	Event *BurnMintTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAllowListRemove)
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
		it.Event = new(BurnMintTokenPoolAllowListRemove)
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

func (it *BurnMintTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAllowListRemoveIterator{contract: _BurnMintTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAllowListRemove)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*BurnMintTokenPoolAllowListRemove, error) {
	event := new(BurnMintTokenPoolAllowListRemove)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCCVConfigUpdatedIterator struct {
	Event *BurnMintTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCCVConfigUpdated)
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
		it.Event = new(BurnMintTokenPoolCCVConfigUpdated)
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

func (it *BurnMintTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCCVConfigUpdatedIterator{contract: _BurnMintTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCCVConfigUpdated)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*BurnMintTokenPoolCCVConfigUpdated, error) {
	event := new(BurnMintTokenPoolCCVConfigUpdated)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolChainAddedIterator struct {
	Event *BurnMintTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolChainAdded)
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
		it.Event = new(BurnMintTokenPoolChainAdded)
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

func (it *BurnMintTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnMintTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolChainAddedIterator{contract: _BurnMintTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolChainAdded)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnMintTokenPoolChainAdded, error) {
	event := new(BurnMintTokenPoolChainAdded)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolChainConfiguredIterator struct {
	Event *BurnMintTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolChainConfigured)
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
		it.Event = new(BurnMintTokenPoolChainConfigured)
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

func (it *BurnMintTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolChainConfiguredIterator{contract: _BurnMintTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolChainConfigured)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseChainConfigured(log types.Log) (*BurnMintTokenPoolChainConfigured, error) {
	event := new(BurnMintTokenPoolChainConfigured)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolChainRemovedIterator struct {
	Event *BurnMintTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolChainRemoved)
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
		it.Event = new(BurnMintTokenPoolChainRemoved)
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

func (it *BurnMintTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolChainRemovedIterator{contract: _BurnMintTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolChainRemoved)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnMintTokenPoolChainRemoved, error) {
	event := new(BurnMintTokenPoolChainRemoved)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolConfigChangedIterator struct {
	Event *BurnMintTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolConfigChanged)
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
		it.Event = new(BurnMintTokenPoolConfigChanged)
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

func (it *BurnMintTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolConfigChangedIterator{contract: _BurnMintTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolConfigChanged)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseConfigChanged(log types.Log) (*BurnMintTokenPoolConfigChanged, error) {
	event := new(BurnMintTokenPoolConfigChanged)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
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
		it.Event = new(BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
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

func (it *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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
		it.Event = new(BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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

func (it *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolDynamicConfigSetIterator struct {
	Event *BurnMintTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolDynamicConfigSet)
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
		it.Event = new(BurnMintTokenPoolDynamicConfigSet)
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

func (it *BurnMintTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolDynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolDynamicConfigSetIterator{contract: _BurnMintTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolDynamicConfigSet)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*BurnMintTokenPoolDynamicConfigSet, error) {
	event := new(BurnMintTokenPoolDynamicConfigSet)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolFinalityConfigUpdatedIterator struct {
	Event *BurnMintTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolFinalityConfigUpdated)
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
		it.Event = new(BurnMintTokenPoolFinalityConfigUpdated)
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

func (it *BurnMintTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolFinalityConfigUpdatedIterator{contract: _BurnMintTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolFinalityConfigUpdated)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*BurnMintTokenPoolFinalityConfigUpdated, error) {
	event := new(BurnMintTokenPoolFinalityConfigUpdated)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolInboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(BurnMintTokenPoolInboundRateLimitConsumed)
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

func (it *BurnMintTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolInboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolInboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolInboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolInboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolLockedOrBurnedIterator struct {
	Event *BurnMintTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolLockedOrBurned)
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
		it.Event = new(BurnMintTokenPoolLockedOrBurned)
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

func (it *BurnMintTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolLockedOrBurnedIterator{contract: _BurnMintTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolLockedOrBurned)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*BurnMintTokenPoolLockedOrBurned, error) {
	event := new(BurnMintTokenPoolLockedOrBurned)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(BurnMintTokenPoolOutboundRateLimitConsumed)
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

func (it *BurnMintTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolOutboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolOutboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolOutboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolOutboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolOwnershipTransferRequestedIterator struct {
	Event *BurnMintTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolOwnershipTransferRequested)
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
		it.Event = new(BurnMintTokenPoolOwnershipTransferRequested)
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

func (it *BurnMintTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolOwnershipTransferRequestedIterator{contract: _BurnMintTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolOwnershipTransferRequested)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintTokenPoolOwnershipTransferRequested, error) {
	event := new(BurnMintTokenPoolOwnershipTransferRequested)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolOwnershipTransferredIterator struct {
	Event *BurnMintTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolOwnershipTransferred)
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
		it.Event = new(BurnMintTokenPoolOwnershipTransferred)
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

func (it *BurnMintTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolOwnershipTransferredIterator{contract: _BurnMintTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolOwnershipTransferred)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*BurnMintTokenPoolOwnershipTransferred, error) {
	event := new(BurnMintTokenPoolOwnershipTransferred)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolPoolFeeWithdrawnIterator struct {
	Event *BurnMintTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolPoolFeeWithdrawn)
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
		it.Event = new(BurnMintTokenPoolPoolFeeWithdrawn)
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

func (it *BurnMintTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolPoolFeeWithdrawnIterator{contract: _BurnMintTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolPoolFeeWithdrawn)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*BurnMintTokenPoolPoolFeeWithdrawn, error) {
	event := new(BurnMintTokenPoolPoolFeeWithdrawn)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolRateLimitAdminSetIterator struct {
	Event *BurnMintTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolRateLimitAdminSet)
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
		it.Event = new(BurnMintTokenPoolRateLimitAdminSet)
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

func (it *BurnMintTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolRateLimitAdminSetIterator{contract: _BurnMintTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolRateLimitAdminSet)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*BurnMintTokenPoolRateLimitAdminSet, error) {
	event := new(BurnMintTokenPoolRateLimitAdminSet)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolReleasedOrMintedIterator struct {
	Event *BurnMintTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolReleasedOrMinted)
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
		it.Event = new(BurnMintTokenPoolReleasedOrMinted)
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

func (it *BurnMintTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolReleasedOrMintedIterator{contract: _BurnMintTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolReleasedOrMinted)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*BurnMintTokenPoolReleasedOrMinted, error) {
	event := new(BurnMintTokenPoolReleasedOrMinted)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolRemotePoolAddedIterator struct {
	Event *BurnMintTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolRemotePoolAdded)
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
		it.Event = new(BurnMintTokenPoolRemotePoolAdded)
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

func (it *BurnMintTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolRemotePoolAddedIterator{contract: _BurnMintTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolRemotePoolAdded)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*BurnMintTokenPoolRemotePoolAdded, error) {
	event := new(BurnMintTokenPoolRemotePoolAdded)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolRemotePoolRemovedIterator struct {
	Event *BurnMintTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolRemotePoolRemoved)
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
		it.Event = new(BurnMintTokenPoolRemotePoolRemoved)
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

func (it *BurnMintTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolRemotePoolRemovedIterator{contract: _BurnMintTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolRemotePoolRemoved)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*BurnMintTokenPoolRemotePoolRemoved, error) {
	event := new(BurnMintTokenPoolRemotePoolRemoved)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *BurnMintTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(BurnMintTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _BurnMintTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolTokenTransferFeeConfigDeleted)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(BurnMintTokenPoolTokenTransferFeeConfigDeleted)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *BurnMintTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(BurnMintTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _BurnMintTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolTokenTransferFeeConfigUpdated)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(BurnMintTokenPoolTokenTransferFeeConfigUpdated)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDynamicConfig struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}

func (BurnMintTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnMintTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnMintTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (BurnMintTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnMintTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (BurnMintTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnMintTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (BurnMintTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (BurnMintTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
}

func (BurnMintTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnMintTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnMintTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnMintTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (BurnMintTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (BurnMintTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnMintTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnMintTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnMintTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (BurnMintTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
}

func (_BurnMintTokenPool *BurnMintTokenPool) Address() common.Address {
	return _BurnMintTokenPool.address
}

type BurnMintTokenPoolInterface interface {
	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

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

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

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

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnMintTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnMintTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*BurnMintTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnMintTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnMintTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*BurnMintTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnMintTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnMintTokenPoolConfigChanged, error)

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnMintTokenPoolDynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*BurnMintTokenPoolFinalityConfigUpdated, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnMintTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintTokenPoolOwnershipTransferred, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*BurnMintTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnMintTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*BurnMintTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnMintTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnMintTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnMintTokenPoolRemotePoolRemoved, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
