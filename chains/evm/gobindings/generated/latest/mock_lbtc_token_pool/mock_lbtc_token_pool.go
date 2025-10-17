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

var MockE2ELBTCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_destPoolData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenDataMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346105e757616be8803803809161001d8285610604565b833981019060a0818303126105e75780516001600160a01b03811691908290036105e75760208101516001600160401b0381116105e75781019183601f840112156105e7578251926001600160401b0384116103f4578360051b60208101946100896040519687610604565b8552602080860191830101918683116105e757602001905b8282106105ec575050506100b760408301610627565b6100c360608401610627565b608084015190936001600160401b0382116105e7570185601f820112156105e7578051906001600160401b0382116103f4576040519661010d601f8401601f191660200189610604565b828852602083830101116105e75760005b8281106105d257505060206000918701015233156105c157600180546001600160a01b03191633179055811580156105b0575b801561059f575b61058e578160209160049360805260c0526040519283809263313ce56760e01b82525afa809160009161054b575b5090610527575b50600860a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261040a575b5080516001600160401b0381116103f457601054600181811c911680156103ea575b60208210146103d457601f811161036f575b50602091601f821160011461030b57918192600092610300575b50508160011b916000199060031b1c1916176010555b60405161640c90816107dc8239608051818181611533015281816117d50152818161199001528181611f87015281816121dc0152818161239a01528181612d2c01528181612f9a01528181613064015281816131cc015281816133a501528181613b3901528181613ba801528181613cce0152614884015260a05181818161185e01528181613af501528181614fee0152615071015260c051818181610c45015281816115ce0152818161202201528181612dc70152613269015260e051818181610bf201528181611613015281816120670152612af80152f35b01519050388061020f565b601f198216926010600052806000209160005b8581106103575750836001951061033e575b505050811b01601055610225565b015160001960f88460031b161c19169055388080610330565b9192602060018192868501518155019401920161031e565b60106000527f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae672601f830160051c810191602084106103ca575b601f0160051c01905b8181106103be57506101f5565b600081556001016103b1565b90915081906103a8565b634e487b7160e01b600052602260045260246000fd5b90607f16906101e3565b634e487b7160e01b600052604160045260246000fd5b60206040516104198282610604565b60008152600036813760e051156105165760005b8151811015610494576001906001600160a01b0361044b828561063b565b5116846104578261067d565b610464575b50500161042d565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388461045c565b505060005b825181101561050d576001906001600160a01b036104b7828661063b565b5116801561050757836104c98261077b565b6104d7575b50505b01610499565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138836104ce565b506104d1565b505050386101c1565b6335f4a7b360e01b60005260046000fd5b60ff166008811461018d576332ad3e0760e11b600052600860045260245260446000fd5b6020813d602011610586575b8161056460209383610604565b8101031261058257519060ff8216820361057f575038610186565b80fd5b5080fd5b3d9150610557565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610158565b506001600160a01b03831615610151565b639b15e16f60e01b60005260046000fd5b80602080928401015182828b0101520161011e565b600080fd5b602080916105f984610627565b8152019101906100a1565b601f909101601f19168101906001600160401b038211908210176103f457604052565b51906001600160a01b03821682036105e757565b805182101561064f5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561064f5760005260206000200190600090565b600081815260036020526040902054801561077457600019810181811161075e5760025460001981019190821161075e5781810361070d575b50505060025480156106f757600019016106d1816002610665565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61074661071e61072f936002610665565b90549060031b1c9283926002610665565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806106b6565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146107d557600254680100000000000000008110156103f4576107bc61072f8260018594016002556002610665565b9055600254906000526003602052604060002055600190565b5060009056fe60a080604052600436101561001357600080fd5b60006080526080513560e01c90816301ffc9a714613e7357508063164e68de14613c2e578063181f5a7714613bcc57806321df0da714613b7a578063240028e814613b1957806324f65ee714613ada5780632a10097b146138805780632c286daf1461377657806337b1924714613665578063390775371461315b5780633e5db5d11461313e578063489a68f214612c895780634c5ef0ed14612c4457806354c8a4f314612ac65780635df45a3714612aa257806362ddd3c414612a1d578063698c2c66146129515780636d3d1a581461291c5780637437ff9f146128db57806379ba5097146128045780637d54534e14612773578063804ba5a9146127315780638926f54f146126ec57806389720a62146126815780638da5cb5b1461264c578063962d4020146124ef5780639a4575b914611f28578063a42a7b8b14611db3578063a7cd63b714611d45578063acfecf9114611c18578063af58d59f14611bcc578063b1c71c65146114a9578063b794658014611471578063c4bffe2b14611337578063c75eea9c1461128c578063cf7401f3146110e5578063d966866b14610c69578063dc0bd97114610c17578063e0351e1314610bd9578063e8a1da17146102c75763f2fde38b146101e857600080fd5b346102c15760206003193601126102c15773ffffffffffffffffffffffffffffffffffffffff610216613fa4565b61021e614bd4565b163381146102955760805180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168217815560015473ffffffffffffffffffffffffffffffffffffffff16907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360805180f35b7fdad89dca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b60805180fd5b346102c1576102d5366144a3565b9190926102e0614bd4565b608051905b828210610a17575050506080519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610a1157600581901b830135858112156102c1578301610120813603126102c157604051946103548661406b565b813567ffffffffffffffff81168103610a0c578652602082013567ffffffffffffffff81116102c15782019436601f870112156102c1578535610396816147cf565b966103a460405198896140a3565b81885260208089019260051b820101903682116102c15760208101925b8284106109dd575050505060208701958652604083013567ffffffffffffffff81116102c1576103f49036908501614485565b916040880192835261041e61040c366060870161460b565b9460608a0195865260c036910161460b565b9560808901968752610430855161580b565b61043a875161580b565b835151156109b15761045667ffffffffffffffff8a5116616284565b156109765767ffffffffffffffff89511660805152600860205260406080512061059a86516fffffffffffffffffffffffffffffffff604082015116906105556fffffffffffffffffffffffffffffffff602083015116915115158360806040516104c08161406b565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106c088516fffffffffffffffffffffffffffffffff6040820151169061067b6fffffffffffffffffffffffffffffffff602083015116915115158360806040516105e48161406b565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff8211610945576106e38354614264565b601f8111610906575b506020906001601f8411146108605791809161073d9360805192610855575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b6080515b88518051821015610779579061077360019261076c8367ffffffffffffffff8f511692614992565b519061517b565b01610744565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561084767ffffffffffffffff60019796949851169251935191516108136107de60405196879687526101006020880152610100870190614141565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610322565b015190508e8061070b565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe083169184608051528160805120926080515b8181106108ee57509084600195949392106108b7575b505050811b019055610740565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806108aa565b92936020600181928786015181550195019301610894565b610935908460805152602060805120601f850160051c8101916020861061093b575b601f0160051c0190614b6e565b8d6106ec565b9091508190610928565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b67ffffffffffffffff8951167f1d5ad3c500000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f14c880ca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b833567ffffffffffffffff81116102c157602091610a018392833691870101614485565b8152019301926103c1565b600080fd5b60805180f35b9092919367ffffffffffffffff610a37610a328688866146d5565b614691565b1692610a4284615fc5565b15610ba95783608051526008602052610a62600560406080512001615dcc565b926080515b8451811015610aa15760019086608051526008602052610a9a600560406080512001610a938389614992565b51906160f0565b5001610a67565b5093909491959250806080515260086020526005604060805120608051815560805160018201556080516002820155608051600382015560048101610ae68154614264565b80610b59575b505001805490608051815581610b35575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102e5565b60805152602060805120908101905b81811015610afd576080518155600101610b44565b601f8111600114610b72575060805190555b8880610aec565b610b929082608051526001601f6020608051209201861c82019101614b6e565b608080518290525160208120918190559055610b6b565b837f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c15760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102c1576080516003193601126102c157602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c157610c9a903690600401614184565b610ca2614bd4565b608051915b818310610cb45760805180f35b610cc2610a32848484614a98565b610cda610cd0858585614a98565b6020810190614ad8565b9091610cf4610cea878787614a98565b6040810190614ad8565b90610d0d610d03898989614a98565b6060810190614ad8565b9091610d27610d1d8b8b8b614a98565b6080810190614ad8565b949097610d3d610d38368a846147e7565b615741565b610d4b610d383684866147e7565b610d59610d383686886147e7565b610d67610d3836888c6147e7565b604051610d7381613fe8565b610d7e368a846147e7565b8152610d8b3684866147e7565b6020820152610d9b3686886147e7565b6040820152610dab36888c6147e7565b606082015267ffffffffffffffff881660805152600e602052604060805120815180519067ffffffffffffffff8211610945576801000000000000000082116109455760209083548385558084106110c6575b500182608051526020608051206080515b83811061109c5750505050602082015180519067ffffffffffffffff82116109455768010000000000000000821161094557602090600184015483600186015580841061107a575b500160018301608051526020608051206080515b8381106110505750505050604082015180519067ffffffffffffffff82116109455768010000000000000000821161094557602090600284015483600286015580841061102e575b500160028301608051526020608051206080515b838110611004575050505060036060919e9c9d9e019101519081519167ffffffffffffffff831161094557680100000000000000008311610945576020908254848455808510610fe5575b500190608051526020608051206080515b838110610fbb5750505050610fa0608095610fb095610f927fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a96610f8467ffffffffffffffff976040519d8d8f9e8f9081520191614b85565b918b830360208d0152614b85565b9188830360408a0152614b85565b9285840360608701521696614b85565b0390a2019190610ca7565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610f23565b610ffe9084608051528584608051209182019101614b6e565b38610f12565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610ec7565b61104a9060028601608051528484608051209182019101614b6e565b38610eb3565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e6b565b6110969060018601608051528484608051209182019101614b6e565b38610e57565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e0f565b6110df9085608051528484608051209182019101614b6e565b38610dfe565b346102c15760e06003193601126102c1576110fe61421f565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01126102c15760405161113481614087565b60243580151581036102c15781526044356fffffffffffffffffffffffffffffffff811681036102c15760208201526064356fffffffffffffffffffffffffffffffff811681036102c15760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c01126102c157604051906111bb82614087565b60843580151581036102c157825260a4356fffffffffffffffffffffffffffffffff811681036102c157602083015260c4356fffffffffffffffffffffffffffffffff811681036102c157604083015273ffffffffffffffffffffffffffffffffffffffff600a54163314158061126a575b61123a57610a119261557f565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561122d565b346102c15760206003193601126102c15767ffffffffffffffff6112ae61421f565b6112b66149e5565b50166080515260086020526113336112da6112d5604060805120614a10565b6156bc565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102c1576080516003193601126102c157608051506040516006548082528160208101600660805152602060805120926080515b818110611458575050611381925003826140a3565b8051906113a6611390836147cf565b9261139e60405194856140a3565b8084526147cf565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020840192013683376080515b8151811015611407578067ffffffffffffffff6113f460019385614992565b51166114008287614992565b52016113d5565b505090604051918291602083019060208452518091526040830191906080515b818110611435575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611427565b845483526001948501948694506020909301920161136c565b346102c15760206003193601126102c15761133361149561149061421f565b614a76565b604051918291602083526020830190614141565b346102c15760606003193601126102c15760043567ffffffffffffffff81116102c15760a060031982360301126102c1576114e26141b5565b9060443567ffffffffffffffff81116102c157611503903690600401614485565b5061150c614979565b50608481019061151b82614771565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611b7e57602481019077ffffffffffffffff0000000000000000000000000000000061158183614691565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a875760805191611b4f575b50611b235761161160448201614771565b7f0000000000000000000000000000000000000000000000000000000000000000611ac3575b5067ffffffffffffffff61164a83614691565b16611662816000526007602052604060002054151590565b15611a9457602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611a875760805190611a24575b73ffffffffffffffffffffffffffffffffffffffff91501633036119f457606461ffff910135931692831515928380946119e5575b1561192a5761ffff600b5416948581106118f657506118ba945061174d61173d61172385614691565b67ffffffffffffffff16600052600c602052604060002090565b8361174784614771565b91615b85565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff61178961178386614691565b93614771565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b9182906118c4575b50611490816117cf61185693614691565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2614691565b9060405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526118926040826140a3565b6040519261189f8461404f565b835260208301526040519283926040845260408401906145b7565b9060208301520390f35b6118569192506118ee611490916127106118e761ffff600b5460101c1683614b2c565b04906156af565b9291506117be565b85907fe08f03ef00000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b506118ba935067ffffffffffffffff61194283614691565b16806080515260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894482806119b860406080512073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615b85565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26117b6565b5061ffff600b541615156116fa565b7f728fe07b0000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b506020813d602011611a7f575b81611a3e602093836140a3565b810103126102c1575173ffffffffffffffffffffffffffffffffffffffff811681036102c15773ffffffffffffffffffffffffffffffffffffffff906116c5565b3d9150611a31565b6040513d608051823e3d90fd5b7fa9902c7e00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b73ffffffffffffffffffffffffffffffffffffffff16611af0816000526003602052604060002054151590565b611637577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f53ad11d800000000000000000000000000000000000000000000000000000000608051526004608051fd5b611b71915060203d602011611b77575b611b6981836140a3565b810190614f62565b85611600565b503d611b5f565b73ffffffffffffffffffffffffffffffffffffffff611b9c83614771565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b346102c15760206003193601126102c15767ffffffffffffffff611bee61421f565b611bf66149e5565b50166080515260086020526113336112da6112d5600260406080512001614a10565b346102c15767ffffffffffffffff611c2f366144f5565b929091611c3a614bd4565b1690611c53826000526007602052604060002054151590565b15611d155781608051526008602052611c86600560406080512001611c7936868561444e565b60208151910120906160f0565b15611cce577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611cc56040519283926020845260208401916149a6565b0390a260805180f35b611d11906040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916149a6565b0390fd5b507f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c157608051506040516002548082526020820190600260805152602060805120906080515b818110611d9d5761133385611d91818703826140a3565b60405191829182614536565b8254845260209093019260019283019201611d7a565b346102c15760206003193601126102c15767ffffffffffffffff611dd561421f565b16608051526008602052611df0600560406080512001615dcc565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611e36611e20846147cf565b93611e2e60405195866140a3565b8085526147cf565b016080515b818110611f175750506080515b8151811015611e925780611e5e60019284614992565b51608051526009602052611e7660406080512061438e565b611e808286614992565b52611e8b8185614992565b5001611e48565b826040518091602082016020835281518091526040830190602060408260051b860101930191608051905b828210611ecc57505050500390f35b91936020611f07827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614141565b9601920192018594939192611ebd565b806060602080938701015201611e3b565b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c15760a060031982360301126102c157611f61614979565b5060848101611f6f81614771565b73ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116036124d157602482019177ffffffffffffffff00000000000000000000000000000000611fd584614691565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a8757608051916124b2575b50611b235761206560448201614771565b7f0000000000000000000000000000000000000000000000000000000000000000612452575b5067ffffffffffffffff61209e84614691565b166120b6816000526007602052604060002054151590565b15611a9457602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611a8757608051906123ef575b73ffffffffffffffffffffffffffffffffffffffff91501633036119f457606401359060805160001461233b5761ffff600b541680612306575061216261173d61172385614691565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff61219861178386614691565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b156102c1576040517f42966c6800000000000000000000000000000000000000000000000000000000815281600482015260805181602481608051875af18015611a87576122ed575b6113336122bd61149086867ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108767ffffffffffffffff61228785614691565b6040805173ffffffffffffffffffffffffffffffffffffffff90961686523360208701528501929092521691806060810161184e565b604051906122ca8261404f565b81526122d46142b7565b60208201526040519182916020835260208301906145b7565b6080516122f9916140a3565b6080516102c15783612248565b7fe08f03ef00000000000000000000000000000000000000000000000000000000608051526080516004526024526044608051fd5b5067ffffffffffffffff61234e83614691565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894482806123c2604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615b85565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26121c5565b506020813d60201161244a575b81612409602093836140a3565b810103126102c1575173ffffffffffffffffffffffffffffffffffffffff811681036102c15773ffffffffffffffffffffffffffffffffffffffff90612119565b3d91506123fc565b73ffffffffffffffffffffffffffffffffffffffff1661247f816000526003602052604060002054151590565b61208b577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b6124cb915060203d602011611b7757611b6981836140a3565b84612054565b611b9c73ffffffffffffffffffffffffffffffffffffffff91614771565b346102c15760606003193601126102c15760043567ffffffffffffffff81116102c157612520903690600401614184565b9060243567ffffffffffffffff81116102c157612541903690600401614586565b9060443567ffffffffffffffff81116102c157612562903690600401614586565b73ffffffffffffffffffffffffffffffffffffffff600a54163314158061262a575b61123a57838614801590612620575b6125f4576080515b8681106125a85760805180f35b806125ee6125bc610a326001948b8b6146d5565b6125c7838989614969565b6125e86125e06125d886898b614969565b92369061460b565b91369061460b565b9161557f565b0161259b565b7f568efce200000000000000000000000000000000000000000000000000000000608051526004608051fd5b5080861415612593565b5073ffffffffffffffffffffffffffffffffffffffff60015416331415612584565b346102c1576080516003193601126102c157602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102c15760c06003193601126102c15761269a613fa4565b506126a3614208565b6126ab6141c6565b5060843567ffffffffffffffff81116102c1576126cc903690600401614236565b505060a4359060028210156102c15761133391611d9191604435906148f3565b346102c15760206003193601126102c157602061272767ffffffffffffffff61271361421f565b166000526007602052604060002054151590565b6040519015158152f35b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c157612765610a119136906004016141d7565b9061276e614bd4565b614c3c565b346102c15760206003193601126102c1577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6127c4613fa4565b6127cc614bd4565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a160805180f35b346102c1576080516003193601126102c1576080515473ffffffffffffffffffffffffffffffffffffffff811633036128af577fffffffffffffffffffffffff0000000000000000000000000000000000000000600154913382841617600155166080515573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0608051608051a360805180f35b7f02b543c600000000000000000000000000000000000000000000000000000000608051526004608051fd5b346102c1576080516003193601126102c1576004546005546040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915290f35b346102c1576080516003193601126102c157602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b346102c15760406003193601126102c15761296a613fa4565b602435612975614bd4565b73ffffffffffffffffffffffffffffffffffffffff82169182156109b1577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045581600555612a14604051928392836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0390a160805180f35b346102c157612a2b366144f5565b612a36929192614bd4565b67ffffffffffffffff8216612a58816000526007602052604060002054151590565b15612a735750610a1192612a6d91369161444e565b9061517b565b7f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c1576080516003193601126102c1576020612abe61483b565b604051908152f35b346102c157612aee612af6612ada366144a3565b9491612ae7939193614bd4565b36916147e7565b9236916147e7565b7f000000000000000000000000000000000000000000000000000000000000000015612c18576080515b8251811015612b93578073ffffffffffffffffffffffffffffffffffffffff612b4b60019386614992565b5116612b5681615e2f565b612b62575b5001612b20565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612b5b565b506080515b8151811015610a11578073ffffffffffffffffffffffffffffffffffffffff612bc360019385614992565b51168015612c1257612bd481616224565b612be1575b505b01612b98565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612bd9565b50612bdb565b7f35f4a7b300000000000000000000000000000000000000000000000000000000608051526004608051fd5b346102c15760406003193601126102c157612c5d61421f565b60243567ffffffffffffffff81116102c157602091612c83612727923690600401614485565b90614792565b346102c15760406003193601126102c15760043567ffffffffffffffff81116102c1578060040161010060031983360301126102c157612cc76141b5565b90604051612cd481614033565b6080519052612d05612cfb612cf6612cef60c48701856146e5565b369161444e565b614f7a565b606485013561506e565b916084840190612d1482614771565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611b7e57602485019277ffffffffffffffff00000000000000000000000000000000612d7a85614691565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a87576080519161311f575b50611b235767ffffffffffffffff612e1085614691565b16612e28816000526007602052604060002054151590565b15611a9457602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611a875760805191613100575b50156119f457612ea184614691565b90612eb760a4880192612c83612cef85856146e5565b156130b9575050612f9461178360446020977ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09561ffff67ffffffffffffffff961615156000146130095785612f0c89614691565b1660805152600d8a52612f286040608051208a61174784614771565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff786612f566117838b614691565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018d90529190931692a25b0194612f8e86614771565b50614691565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081168252336020830152909216908201526060810185905292169180608081015b0390a28060405161300081614033565b52604051908152f35b508461301488614691565b16806080515260088a527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c898061308c60026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391615b85565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612f83565b6130c392506146e5565b611d116040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916149a6565b613119915060203d602011611b7757611b6981836140a3565b87612e92565b613138915060203d602011611b7757611b6981836140a3565b87612df9565b346102c1576080516003193601126102c1576113336114956142b7565b346102c15760206003193601126102c15760043567ffffffffffffffff81116102c157806004019061010060031982360301126102c15760405161319e81614033565b60805190526064810135608482016131b581614771565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081169116036124d15750602482019177ffffffffffffffff0000000000000000000000000000000061321c84614691565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611a875760805191613646575b50611b235767ffffffffffffffff6132b284614691565b166132ca816000526007602052604060002054151590565b15611a9457602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611a875760805191613627575b50156119f45761334383614691565b61335860a4830191612c83612cef84896146e5565b1561361d57508192936080515067ffffffffffffffff61337786614691565b16806080515260086020526133cd60026040608051200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016958691615b85565b6040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a26020613422601054614264565b14613533575b506044019261343684614771565b823b156102c1576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260805173ffffffffffffffffffffffffffffffffffffffff909216600482015260248101859052908180604481010381608051875af1948515611a8757612ff0856134e36117837ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09667ffffffffffffffff9660209b6135215750614691565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b60805161352d916140a3565b8b612f8e565b61354060e48301826146e5565b8101906040818303126102c157803567ffffffffffffffff81116102c1578261356a918301614485565b60208201359167ffffffffffffffff83116102c157602093849361358e9201614485565b506135a2604051918281519485920161411e565b80608051928101039060025afa15611a8757608051519060c48301906135d16135cb83836146e5565b90614736565b83036135de575050613428565b6135eb916135cb916146e5565b7f7f24931100000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b6130c390856146e5565b613640915060203d602011611b7757611b6981836140a3565b85613334565b61365f915060203d602011611b7757611b6981836140a3565b8561329b565b346102c15760a06003193601126102c15761367e613fa4565b50613687614208565b60443567ffffffffffffffff81116102c15760031960a091360301126102c1576136af6141c6565b506084359067ffffffffffffffff82116102c1576136da67ffffffffffffffff923690600401614236565b50506040516136e881613fe8565b60805181526080516020820152608051604082015260606080519101521660805152600f602052608060408151206040519061372382613fe8565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b346102c15760606003193601126102c15760043561ffff8116908190036102c15761379f6141b5565b9060443567ffffffffffffffff81116102c1576137c09036906004016141d7565b906137c9614bd4565b61ffff8416936127108510156138505783927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba959261383f926040967fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600b549360101b1692161717600b55614c3c565b82519182526020820152a160805180f35b847f95f3517a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346102c15760406003193601126102c15760043567ffffffffffffffff81116102c157366023820112156102c157806004013567ffffffffffffffff81116102c15760248201916024369160a084020101116102c15760243567ffffffffffffffff81116102c1576138f6903690600401614184565b919092613901614bd4565b6080515b82811061397d575050506080515b8181106139205760805180f35b8067ffffffffffffffff61393a610a3260019486886146d5565b168060805152600f602052608051604060805120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8608051608051a201613913565b8061398e610a326001938686614652565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7060806139bc848888614652565b92613acc613a9263ffffffff613ac1613a8582613ab667ffffffffffffffff60208c0198169a8b8a5152600f60205260408a5120836139fa8b6146a6565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff00000000613a37876146a6565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e6146a6565b60401b1696019e8f6146b7565b151560601b1695161716171717905582613aae6040519a6146c4565b1689526146c4565b1660208701526146c4565b1660408401526145e1565b15156060820152a201613905565b346102c1576080516003193601126102c157602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c15760206003193601126102c1576020613b34613fa4565b6040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081169216919091148152f35b346102c1576080516003193601126102c157602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102c1576080516003193601126102c1576040805161133391613bf090826140a3565b601a81527f4d6f636b4532454c425443546f6b656e506f6f6c20312e352e310000000000006020820152604051918291602083526020830190614141565b346102c15760206003193601126102c157613c47613fa4565b613c4f614bd4565b613c5761483b565b9081613c635760805180f35b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff831660248301526044808301859052825290613d7890613cc46064826140a3565b60408051909390917f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169190613d0f86856140a3565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260805191608051915190608051855af13d15613e6b573d91613d5a836140e4565b92613d67875194856140a3565b83526080513d90602085013e616333565b805180613dca575b505073ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992602092519485521692a28080610a11565b90602080613ddc938301019101614f62565b15613de8578380613d80565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091616333565b346102c15760206003193601126102c157600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102c157817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115613f7a575b8115613f50575b8115613f26575b8115613efc575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613ef5565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613eee565b7f479eecb20000000000000000000000000000000000000000000000000000000081149150613ee7565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613ee0565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610a0c57565b359073ffffffffffffffffffffffffffffffffffffffff82168203610a0c57565b6080810190811067ffffffffffffffff82111761400457604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761400457604052565b6040810190811067ffffffffffffffff82111761400457604052565b60a0810190811067ffffffffffffffff82111761400457604052565b6060810190811067ffffffffffffffff82111761400457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761400457604052565b67ffffffffffffffff811161400457601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106141315750506000910152565b8181015183820152602001614121565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f60209361417d8151809281875287808801910161411e565b0116010190565b9181601f84011215610a0c5782359167ffffffffffffffff8311610a0c576020808501948460051b010111610a0c57565b6024359061ffff82168203610a0c57565b6064359061ffff82168203610a0c57565b9181601f84011215610a0c5782359167ffffffffffffffff8311610a0c5760208085019460e08502010111610a0c57565b6024359067ffffffffffffffff82168203610a0c57565b6004359067ffffffffffffffff82168203610a0c57565b9181601f84011215610a0c5782359167ffffffffffffffff8311610a0c5760208381860195010111610a0c57565b90600182811c921680156142ad575b602083101461427e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614273565b60405190600082601054916142cb83614264565b808352926001811690811561435157506001146142f1575b6142ef925003836140a3565b565b506010600090815290917f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae6725b8183106143355750509060206142ef928201016142e3565b602091935080600191548385890101520191019091849261431d565b602092506142ef9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b8201016142e3565b90604051918260008254926143a284614264565b808452936001811690811561440e57506001146143c7575b506142ef925003836140a3565b90506000929192526020600020906000915b8183106143f25750509060206142ef92820101386143ba565b60209193508060019154838589010152019101909184926143d9565b602093506142ef9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386143ba565b92919261445a826140e4565b9161446860405193846140a3565b829481845281830111610a0c578281602093846000960137010152565b9080601f83011215610a0c578160206144a09335910161444e565b90565b6040600319820112610a0c5760043567ffffffffffffffff8111610a0c57816144ce91600401614184565b929092916024359067ffffffffffffffff8211610a0c576144f191600401614184565b9091565b906040600319830112610a0c5760043567ffffffffffffffff81168103610a0c57916024359067ffffffffffffffff8211610a0c576144f191600401614236565b602060408183019282815284518094520192019060005b81811061455a5750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161454d565b9181601f84011215610a0c5782359167ffffffffffffffff8311610a0c5760208085019460608502010111610a0c57565b6144a09160206145d08351604084526040840190614141565b920151906020818403910152614141565b35908115158203610a0c57565b35906fffffffffffffffffffffffffffffffff82168203610a0c57565b9190826060910312610a0c5760405161462381614087565b604061464d818395614634816145e1565b8552614642602082016145ee565b6020860152016145ee565b910152565b91908110156146625760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff81168103610a0c5790565b3563ffffffff81168103610a0c5790565b358015158103610a0c5790565b359063ffffffff82168203610a0c57565b91908110156146625760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610a0c570180359067ffffffffffffffff8211610a0c57602001918136038313610a0c57565b359060208110614744575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b3573ffffffffffffffffffffffffffffffffffffffff81168103610a0c5790565b9067ffffffffffffffff6144a092166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116140045760051b60200190565b9291906147f3816147cf565b9361480160405195866140a3565b602085838152019160051b8101928311610a0c57905b82821061482357505050565b6020809161483084613fc7565b815201910190614817565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa9081156148e7576000916148b8575090565b90506020813d6020116148df575b816148d3602093836140a3565b81010312610a0c575190565b3d91506148c6565b6040513d6000823e3d90fd5b67ffffffffffffffff16600052600e602052604060002091600281101561493a57600114614929578160016144a0930190615453565b81600260036144a094019101615453565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9190811015614662576060020190565b604051906149868261404f565b60606020838281520152565b80518210156146625760209160051b010190565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b604051906149f28261406b565b60006080838281528260208201528260408201528260608201520152565b90604051614a1d8161406b565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff1660005260086020526144a0600460406000200161438e565b91908110156146625760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610a0c570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610a0c570180359067ffffffffffffffff8211610a0c57602001918160051b36038313610a0c57565b81810292918115918404141715614b3f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110614b79575050565b60008155600101614b6e565b9160209082815201919060005b818110614b9f5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff614bc688613fc7565b168152019401929101614b92565b73ffffffffffffffffffffffffffffffffffffffff600154163303614bf557565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b356fffffffffffffffffffffffffffffffff81168103610a0c5790565b9160005b82811015614f5c5760e0810284016000614c5982614691565b9067ffffffffffffffff821691614c7d836000526007602052604060002054151590565b15614f3057614d469260408593614cf1614ceb94614ceb614cb1602060019c9b0192611723614cac368661460b565b61580b565b91825463ffffffff8160801c16159081614f12575b81614f03575b81614ee8575b81614ed9575b5080614eca575b614e3f575b369061460b565b90615952565b6080850192614d03614cac368661460b565b8152600d6020522092835463ffffffff8160801c16159081614e21575b81614e12575b81614df7575b81614de8575b5080614dd9575b614d4c575b50369061460b565b01614c40565b614d6960a06fffffffffffffffffffffffffffffffff9201614c1f565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614d3e565b50614de3826146b7565b614d39565b60ff915060a01c161538614d32565b6fffffffffffffffffffffffffffffffff8116159150614d2c565b8589015460801c159150614d26565b858901546fffffffffffffffffffffffffffffffff16159150614d20565b6fffffffffffffffffffffffffffffffff614e5b878b01614c1f565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355614ce4565b50614ed4816146b7565b614cdf565b60ff915060a01c161538614cd8565b6fffffffffffffffffffffffffffffffff8116159150614cd2565b848e015460801c159150614ccc565b848e01546fffffffffffffffffffffffffffffffff16159150614cc6565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b90816020910312610a0c57518015158103610a0c5790565b80518015614fea57602003614fac578051602082810191830183900312610a0c57519060ff8211614fac575060ff1690565b611d11906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614141565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211614b3f57565b60ff16604d8111614b3f57600a0a90565b811561503f570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146151745782841161514a57906150b391615010565b91604d60ff8416118015615111575b6150db575050906150d56144a092615024565b90614b2c565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061511b83615024565b801561503f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0484116150c2565b61515391615010565b91604d60ff8416116150db5750509061516e6144a092615024565b90615035565b5050505090565b908051156153bb5767ffffffffffffffff815160208301209216918260005260086020526151b08160056040600020016162de565b156153775760005260096020526040600020815167ffffffffffffffff8111614004576151dd8254614264565b601f8111615345575b506020601f821160011461527f5791615259827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361526f95600091615274575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614141565b0390a2565b905084015138615228565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b81811061532d57509261526f9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106152f6575b5050811b019055611495565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905538806152ea565b9192602060018192868a0151815501940192016152af565b61537190836000526020600020601f840160051c8101916020851061093b57601f0160051c0190614b6e565b386151e6565b5090611d116040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614141565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b8181106154175750506142ef925003836140a3565b845473ffffffffffffffffffffffffffffffffffffffff16835260019485019487945060209093019201615402565b91908201809211614b3f57565b61545c906153e5565b916005548015159182615574575b5050615474575090565b61547d906153e5565b9081518061548b5750905090565b615496908251615446565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06154da6154c4866147cf565b956154d260405197886140a3565b8087526147cf565b0136602086013760005b8251811015615522578073ffffffffffffffffffffffffffffffffffffffff61550f60019386614992565b511661551b8288614992565b52016154e4565b509160005b815181101561556f578073ffffffffffffffffffffffffffffffffffffffff61555260019385614992565b5116615568615562838751615446565b88614992565b5201615527565b505050565b10159050388061546a565b67ffffffffffffffff166000818152600760205260409020549092919015615681579161567e60e09261564a856155d67f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b9761580b565b8460005260086020526155ed816040600020615952565b6155f68361580b565b846000526008602052615610836002604060002001615952565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b91908203918211614b3f57565b6156c46149e5565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691615721602085019361571b61570e63ffffffff875116426156af565b8560808901511690614b2c565b90615446565b8082101561573a57505b16825263ffffffff4216905290565b905061572b565b805160005b81811061575257505050565b60018101808211614b3f575b82811061576e5750600101615746565b73ffffffffffffffffffffffffffffffffffffffff61578d8386614992565b511673ffffffffffffffffffffffffffffffffffffffff6157ae8387614992565b5116146157bd5760010161575e565b73ffffffffffffffffffffffffffffffffffffffff6157dc8386614992565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8051156158ab576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106158485750565b6064906158a9604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590615933575b6158d25750565b6064906158a9604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208201511615156158cb565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615a8b606092805461598f63ffffffff8260801c16426156af565b9081615aca575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014615ac257508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155615a3f8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b61567e60405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916159c6565b6fffffffffffffffffffffffffffffffff91615aff839283615af86001880154948286169560801c90614b2c565b9116615446565b80821015615b7e57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880615996565b9050615b09565b9182549060ff8260a01c16158015615dc4575b615dbe576fffffffffffffffffffffffffffffffff82169160018501908154615bdd63ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426156af565b9081615d20575b5050848110615cd45750838310615c3e575050615c136fffffffffffffffffffffffffffffffff9283926156af565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615c4d81856156af565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211614b3f57615c9b615ca09273ffffffffffffffffffffffffffffffffffffffff96615446565b615035565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615d9457615d3b9261571b9160801c90614b2c565b80841015615d8f5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615be4565b615d46565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615b98565b906040519182815491828252602082019060005260206000209260005b818110615dfe5750506142ef925003836140a3565b8454835260019485019487945060209093019201615de9565b80548210156146625760005260206000200190600090565b6000818152600360205260409020548015615fbe577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614b3f57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614b3f57818103615f4f575b5050506002548015615f20577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01615edd816002615e17565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615fa6615f60615f71936002615e17565b90549060031b1c9283926002615e17565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526003602052604060002055388080615ea4565b5050600090565b6000818152600760205260409020548015615fbe577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614b3f57600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614b3f578181036160b6575b5050506006548015615f20577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01616073816006615e17565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b6160d86160c7615f71936006615e17565b90549060031b1c9283926006615e17565b9055600052600760205260406000205538808061603a565b906001820191816000528260205260406000205480151560001461621b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614b3f578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614b3f578181036161e4575b50505080548015615f20577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906161a58282615e17565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6162046161f4615f719386615e17565b90549060031b1c92839286615e17565b90556000528360205260406000205538808061616d565b50505050600090565b8060005260036020526040600020541560001461627e576002546801000000000000000081101561400457616265615f718260018594016002556002615e17565b9055600254906000526003602052604060002055600190565b50600090565b8060005260076020526040600020541560001461627e5760065468010000000000000000811015614004576162c5615f718260018594016006556006615e17565b9055600654906000526007602052604060002055600190565b6000828152600182016020526040902054615fbe5780549068010000000000000000821015614004578261631c615f71846001809601855584615e17565b905580549260005201602052604060002055600190565b919290156163ae5750815115616347575090565b3b156163505790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156163c15750805190602001fd5b611d11906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061414156fea164736f6c634300081a000a",
}

var MockE2ELBTCTokenPoolABI = MockE2ELBTCTokenPoolMetaData.ABI

var MockE2ELBTCTokenPoolBin = MockE2ELBTCTokenPoolMetaData.Bin

func DeployMockE2ELBTCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address, destPoolData []byte) (common.Address, *types.Transaction, *MockE2ELBTCTokenPool, error) {
	parsed, err := MockE2ELBTCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockE2ELBTCTokenPoolBin), backend, token, allowlist, rmnProxy, router, destPoolData)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAccumulatedFees(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAccumulatedFees(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAllowList(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAllowList(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAllowListEnabled(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.GetAllowListEnabled(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _MockE2ELBTCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_MockE2ELBTCTokenPool.CallOpts, remoteChainSelector)
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
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRateLimitAdmin(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRateLimitAdmin(&_MockE2ELBTCTokenPool.CallOpts)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRequiredCCVs(&_MockE2ELBTCTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRequiredCCVs(&_MockE2ELBTCTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _MockE2ELBTCTokenPool.Contract.GetTokenTransferFeeConfig(&_MockE2ELBTCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _MockE2ELBTCTokenPool.Contract.GetTokenTransferFeeConfig(&_MockE2ELBTCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyAllowListUpdates(&_MockE2ELBTCTokenPool.TransactOpts, removes, adds)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyAllowListUpdates(&_MockE2ELBTCTokenPool.TransactOpts, removes, adds)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyCCVConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, ccvConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyCCVConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, ccvConfigArgs)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyFinalityConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyFinalityConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.LockOrBurn0(&_MockE2ELBTCTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.LockOrBurn0(&_MockE2ELBTCTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ReleaseOrMint0(&_MockE2ELBTCTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ReleaseOrMint0(&_MockE2ELBTCTokenPool.TransactOpts, releaseOrMintIn, finality)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetChainRateLimiterConfig(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetChainRateLimiterConfig(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetChainRateLimiterConfigs(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetChainRateLimiterConfigs(&_MockE2ELBTCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetDynamicConfig(&_MockE2ELBTCTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetDynamicConfig(&_MockE2ELBTCTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRateLimitAdmin(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRateLimitAdmin(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitAdmin)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.WithdrawFees(&_MockE2ELBTCTokenPool.TransactOpts, recipient)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.WithdrawFees(&_MockE2ELBTCTokenPool.TransactOpts, recipient)
}

type MockE2ELBTCTokenPoolAllowListAddIterator struct {
	Event *MockE2ELBTCTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolAllowListAdd)
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
		it.Event = new(MockE2ELBTCTokenPoolAllowListAdd)
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

func (it *MockE2ELBTCTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolAllowListAddIterator{contract: _MockE2ELBTCTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolAllowListAdd)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*MockE2ELBTCTokenPoolAllowListAdd, error) {
	event := new(MockE2ELBTCTokenPoolAllowListAdd)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolAllowListRemoveIterator struct {
	Event *MockE2ELBTCTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolAllowListRemove)
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
		it.Event = new(MockE2ELBTCTokenPoolAllowListRemove)
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

func (it *MockE2ELBTCTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolAllowListRemoveIterator{contract: _MockE2ELBTCTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolAllowListRemove)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*MockE2ELBTCTokenPoolAllowListRemove, error) {
	event := new(MockE2ELBTCTokenPoolAllowListRemove)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolCCVConfigUpdatedIterator struct {
	Event *MockE2ELBTCTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolCCVConfigUpdated)
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
		it.Event = new(MockE2ELBTCTokenPoolCCVConfigUpdated)
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

func (it *MockE2ELBTCTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolCCVConfigUpdatedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolCCVConfigUpdated)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*MockE2ELBTCTokenPoolCCVConfigUpdated, error) {
	event := new(MockE2ELBTCTokenPoolCCVConfigUpdated)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

type MockE2ELBTCTokenPoolChainConfiguredIterator struct {
	Event *MockE2ELBTCTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolChainConfigured)
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
		it.Event = new(MockE2ELBTCTokenPoolChainConfigured)
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

func (it *MockE2ELBTCTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolChainConfiguredIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolChainConfigured)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseChainConfigured(log types.Log) (*MockE2ELBTCTokenPoolChainConfigured, error) {
	event := new(MockE2ELBTCTokenPoolChainConfigured)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

type MockE2ELBTCTokenPoolConfigChangedIterator struct {
	Event *MockE2ELBTCTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolConfigChanged)
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
		it.Event = new(MockE2ELBTCTokenPoolConfigChanged)
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

func (it *MockE2ELBTCTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolConfigChangedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolConfigChanged)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseConfigChanged(log types.Log) (*MockE2ELBTCTokenPoolConfigChanged, error) {
	event := new(MockE2ELBTCTokenPoolConfigChanged)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed)
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
		it.Event = new(MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed)
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

func (it *MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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
		it.Event = new(MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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

func (it *MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
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

type MockE2ELBTCTokenPoolFinalityConfigUpdatedIterator struct {
	Event *MockE2ELBTCTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolFinalityConfigUpdated)
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
		it.Event = new(MockE2ELBTCTokenPoolFinalityConfigUpdated)
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

func (it *MockE2ELBTCTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolFinalityConfigUpdatedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolFinalityConfigUpdated)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*MockE2ELBTCTokenPoolFinalityConfigUpdated, error) {
	event := new(MockE2ELBTCTokenPoolFinalityConfigUpdated)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

type MockE2ELBTCTokenPoolPoolFeeWithdrawnIterator struct {
	Event *MockE2ELBTCTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolPoolFeeWithdrawn)
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
		it.Event = new(MockE2ELBTCTokenPoolPoolFeeWithdrawn)
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

func (it *MockE2ELBTCTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*MockE2ELBTCTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolPoolFeeWithdrawnIterator{contract: _MockE2ELBTCTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolPoolFeeWithdrawn)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*MockE2ELBTCTokenPoolPoolFeeWithdrawn, error) {
	event := new(MockE2ELBTCTokenPoolPoolFeeWithdrawn)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRateLimitAdminSetIterator struct {
	Event *MockE2ELBTCTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRateLimitAdminSet)
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
		it.Event = new(MockE2ELBTCTokenPoolRateLimitAdminSet)
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

func (it *MockE2ELBTCTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRateLimitAdminSetIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRateLimitAdminSet)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*MockE2ELBTCTokenPoolRateLimitAdminSet, error) {
	event := new(MockE2ELBTCTokenPoolRateLimitAdminSet)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

type GetDynamicConfig struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}

func (MockE2ELBTCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (MockE2ELBTCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (MockE2ELBTCTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (MockE2ELBTCTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (MockE2ELBTCTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (MockE2ELBTCTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (MockE2ELBTCTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (MockE2ELBTCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (MockE2ELBTCTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
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

func (MockE2ELBTCTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (MockE2ELBTCTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
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
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPool) Address() common.Address {
	return _MockE2ELBTCTokenPool.address
}

type MockE2ELBTCTokenPoolInterface interface {
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

	SDestPoolData(opts *bind.CallOpts) ([]byte, error)

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

	FilterAllowListAdd(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*MockE2ELBTCTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*MockE2ELBTCTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*MockE2ELBTCTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*MockE2ELBTCTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*MockE2ELBTCTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*MockE2ELBTCTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*MockE2ELBTCTokenPoolConfigChanged, error)

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*MockE2ELBTCTokenPoolDynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*MockE2ELBTCTokenPoolFinalityConfigUpdated, error)

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

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*MockE2ELBTCTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*MockE2ELBTCTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*MockE2ELBTCTokenPoolRateLimitAdminSet, error)

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
