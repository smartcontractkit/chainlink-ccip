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

var MockE2ELBTCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_destPoolData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmationConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenDataMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346105fc5761701b803803809161001d8285610619565b833981019060a0818303126105fc5780516001600160a01b03811691908290036105fc5760208101516001600160401b0381116105fc5781019183601f840112156105fc578251926001600160401b038411610409578360051b60208101946100896040519687610619565b8552602080860191830101918683116105fc57602001905b828210610601575050506100b76040830161063c565b6100c36060840161063c565b608084015190936001600160401b0382116105fc570185601f820112156105fc578051906001600160401b038211610409576040519661010d601f8401601f191660200189610619565b828852602083830101116105fc5760005b8281106105e757505060206000918701015233156105d657600180546001600160a01b03191633179055811580156105c5575b80156105b4575b6105a3578160209160049360805260c0526040519283809263313ce56760e01b82525afa8091600091610560575b509061053c575b50600860a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261041f575b5080516001600160401b03811161040957601054600181811c911680156103ff575b60208210146103e957601f8111610384575b50602091601f821160011461032057918192600092610315575b50508160011b916000199060031b1c1916176010555b60405161682a90816107f182396080518181816115c40152818161180d0152818161187601528181611a0701528181611ffe0152818161222a015281816122920152818161244f015281816131a5015281816133d20152818161344001528181613509015281816136710152818161384a01528181613d4601528181613db501528181613edb0152614b43015260a0518181816118ff01528181613d0201528181614fea015261506d015260c051818181610c50015281816116600152818161209a01528181613241015261370e015260e051818181610bfd015281816116a5015281816120df0152612f720152f35b01519050388061020f565b601f198216926010600052806000209160005b85811061036c57508360019510610353575b505050811b01601055610225565b015160001960f88460031b161c19169055388080610345565b91926020600181928685015181550194019201610333565b60106000527f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae672601f830160051c810191602084106103df575b601f0160051c01905b8181106103d357506101f5565b600081556001016103c6565b90915081906103bd565b634e487b7160e01b600052602260045260246000fd5b90607f16906101e3565b634e487b7160e01b600052604160045260246000fd5b602060405161042e8282610619565b60008152600036813760e0511561052b5760005b81518110156104a9576001906001600160a01b036104608285610650565b51168461046c82610692565b610479575b505001610442565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13884610471565b505060005b8251811015610522576001906001600160a01b036104cc8286610650565b5116801561051c57836104de82610790565b6104ec575b50505b016104ae565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138836104e3565b506104e6565b505050386101c1565b6335f4a7b360e01b60005260046000fd5b60ff166008811461018d576332ad3e0760e11b600052600860045260245260446000fd5b6020813d60201161059b575b8161057960209383610619565b8101031261059757519060ff82168203610594575038610186565b80fd5b5080fd5b3d915061056c565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610158565b506001600160a01b03831615610151565b639b15e16f60e01b60005260046000fd5b80602080928401015182828b0101520161011e565b600080fd5b6020809161060e8461063c565b8152019101906100a1565b601f909101601f19168101906001600160401b0382119082101761040957604052565b51906001600160a01b03821682036105fc57565b80518210156106645760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156106645760005260206000200190600090565b60008181526003602052604090205480156107895760001981018181116107735760025460001981019190821161077357818103610722575b505050600254801561070c57600019016106e681600261067a565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61075b61073361074493600261067a565b90549060031b1c928392600261067a565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806106cb565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146107ea5760025468010000000000000000811015610409576107d1610744826001859401600255600261067a565b9055600254906000526003602052604060002055600190565b5060009056fe60c080604052600436101561001357600080fd5b600060a05260a0513560e01c90816301ffc9a71461408057508063164e68de14613e3b578063181f5a7714613dd957806321df0da714613d87578063240028e814613d2657806324f65ee714613ce75780632c06340414613c5757806337b1924714613b0a57806339077537146136005780633e5db5d1146135e3578063489a68f2146131035780634c5ef0ed146130be57806354c8a4f314612f4057806359152aad14612e935780635df45a3714612e6f5780635f789b4014612af757806362ddd3c414612a72578063698c2c66146129a65780636d3d1a58146129715780637437ff9f1461293057806379ba5097146128595780637d54534e146127c85780638926f54f1461278357806389720a62146127185780638da5cb5b146126e3578063962d4020146125865780639a4575b914611f9f578063a42a7b8b14611e2a578063a7cd63b714611dbc578063acfecf9114611c8f578063af58d59f14611c43578063b1c71c651461153b578063b794658014611503578063bb6bb5a714611481578063c4bffe2b14611347578063c75eea9c1461129c578063cf7401f3146110f5578063d966866b14610c74578063dc0bd97114610c22578063e0351e1314610be4578063e8a1da17146102d25763f2fde38b146101f357600080fd5b346102cc5760206003193601126102cc5773ffffffffffffffffffffffffffffffffffffffff6102216141b1565b610229614f13565b163381146102a05760a05180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168217815560015473ffffffffffffffffffffffffffffffffffffffff16907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360a05180f35b7fdad89dca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b60a05180fd5b346102cc576102e0366146b1565b9190926102eb614f13565b60a051905b828210610a225750505060a0519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610a1c57600581901b830135858112156102cc578301610120813603126102cc576040519461035f8661429b565b813567ffffffffffffffff81168103610a17578652602082013567ffffffffffffffff81116102cc5782019436601f870112156102cc5785356103a181614a8e565b966103af60405198896142d3565b81885260208089019260051b820101903682116102cc5760208101925b8284106109e8575050505060208701958652604083013567ffffffffffffffff81116102cc576103ff9036908501614662565b9160408801928352610429610417366060870161483d565b9460608a0195865260c036910161483d565b956080890196875261043b8551615c29565b6104458751615c29565b835151156109bc5761046167ffffffffffffffff8a511661645b565b156109815767ffffffffffffffff89511660a051526008602052604060a051206105a586516fffffffffffffffffffffffffffffffff604082015116906105606fffffffffffffffffffffffffffffffff602083015116915115158360806040516104cb8161429b565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106cb88516fffffffffffffffffffffffffffffffff604082015116906106866fffffffffffffffffffffffffffffffff602083015116915115158360806040516105ef8161429b565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff8211610950576106ee8354614441565b601f8111610911575b506020906001601f84111461086b579180916107489360a05192610860575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b60a0515b88518051821015610784579061077e6001926107778367ffffffffffffffff8f511692614cd1565b51906154d3565b0161074f565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561085267ffffffffffffffff600197969498511692519351915161081e6107e960405196879687526101006020880152610100870190614371565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101939290919361032d565b015190508e80610716565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08316918460a051528160a051209260a0515b8181106108f957509084600195949392106108c2575b505050811b01905561074b565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558d80806108b5565b9293602060018192878601518155019501930161089f565b610940908460a05152602060a05120601f850160051c81019160208610610946575b601f0160051c0190614ead565b8d6106f7565b9091508190610933565b7f4e487b710000000000000000000000000000000000000000000000000000000060a051526041600452602460a051fd5b67ffffffffffffffff8951167f1d5ad3c50000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b833567ffffffffffffffff81116102cc57602091610a0c8392833691870101614662565b8152019301926103cc565b600080fd5b60a05180f35b9092919367ffffffffffffffff610a42610a3d868886614c22565b614a3c565b1692610a4d8461619c565b15610bb4578360a051526008602052610a6d6005604060a0512001615fa3565b9260a0515b8451811015610aac576001908660a051526008602052610aa56005604060a0512001610a9e8389614cd1565b51906162c7565b5001610a72565b50939094919592508060a0515260086020526005604060a0512060a051815560a051600182015560a051600282015560a051600382015560048101610af18154614441565b80610b64575b50500180549060a051815581610b40575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102f0565b60a05152602060a05120908101905b81811015610b085760a0518155600101610b4f565b601f8111600114610b7d575060a05190555b8880610af7565b610b9d908260a051526001601f602060a051209201861c82019101614ead565b60a080518290525160208120918190559055610b76565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102cc5760a0516003193601126102cc5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102cc5760a0516003193601126102cc57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102cc5760206003193601126102cc5760043567ffffffffffffffff81116102cc57610ca5903690600401614680565b90610cae614f13565b60a0516080525b8160805110610cc45760a05180f35b610cd4610a3d6080518484614dd7565b610cee610ce46080518585614dd7565b6020810190614e17565b610d08610cfe6080518787614dd7565b6040810190614e17565b90610d23610d196080518989614dd7565b6060810190614e17565b610d3d610d336080518b8b614dd7565b6080810190614e17565b939094610d53610d4e36898b614aa6565b615b5f565b610d61610d4e368385614aa6565b610d6f610d4e368587614aa6565b610d7d610d4e368789614aa6565b6040519860808a01908a821067ffffffffffffffff8311176109505767ffffffffffffffff91604052610db1368a8c614aa6565b8b52610dbe368486614aa6565b60208c0152610dce368688614aa6565b60408c0152610dde36888a614aa6565b60608c015216988960a05152600e602052604060a05120815180519067ffffffffffffffff8211610950576801000000000000000082116109505760209083548385558084106110d6575b50018260a05152602060a0512060a0515b8381106110ac57505050506001810160208301519081519167ffffffffffffffff83116109505768010000000000000000831161095057602090825484845580851061108d575b50019060a05152602060a0512060a0515b83811061106357505050506002810160408301519081519167ffffffffffffffff831161095057680100000000000000008311610950576020908254848455808510611044575b50019060a05152602060a0512060a0515b83811061101a575050505060036060910191015180519067ffffffffffffffff821161095057680100000000000000008211610950576020908354838555808410610ffb575b50019160a05152602060a051209160a0515b828110610fd1575050505092610fc09492610fa4610fb2937fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a999896610f966040519b8c9b60808d5260808d0191614ec4565b918a830360208c0152614ec4565b918783036040890152614ec4565b918483036060860152614ec4565b0390a2600160805101608052610cb5565b600190602073ffffffffffffffffffffffffffffffffffffffff8451169301928186015501610f42565b611014908560a05152848460a051209182019101614ead565b8f610f30565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610eea565b61105d908460a05152858460a051209182019101614ead565b38610ed9565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e92565b6110a6908460a05152858460a051209182019101614ead565b38610e81565b600190602073ffffffffffffffffffffffffffffffffffffffff8551169401938184015501610e3a565b6110ef908560a05152848460a051209182019101614ead565b38610e29565b346102cc5760e06003193601126102cc5761110e6143cb565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01126102cc57604051611144816142b7565b60243580151581036102cc5781526044356fffffffffffffffffffffffffffffffff811681036102cc5760208201526064356fffffffffffffffffffffffffffffffff811681036102cc5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c01126102cc57604051906111cb826142b7565b60843580151581036102cc57825260a4356fffffffffffffffffffffffffffffffff811681036102cc57602083015260c4356fffffffffffffffffffffffffffffffff811681036102cc57604083015273ffffffffffffffffffffffffffffffffffffffff600a54163314158061127a575b61124a57610a1c926158d7565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561123d565b346102cc5760206003193601126102cc5767ffffffffffffffff6112be6143cb565b6112c6614d24565b501660a0515260086020526113436112ea6112e5604060a05120614d4f565b615a14565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102cc5760a0516003193601126102cc5760a051506040516006548082528160208101600660a05152602060a051209260a0515b818110611468575050611391925003826142d3565b8051906113b66113a083614a8e565b926113ae60405194856142d3565b808452614a8e565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe060208401920136833760a0515b8151811015611417578067ffffffffffffffff61140460019385614cd1565b51166114108287614cd1565b52016113e5565b5050906040519182916020830190602084525180915260408301919060a0515b818110611445575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611437565b845483526001948501948694506020909301920161137c565b346102cc5760206003193601126102cc5760043567ffffffffffffffff81116102cc576114b2903690600401614703565b73ffffffffffffffffffffffffffffffffffffffff600a5416331415806114e1575b61124a57610a1c916151a1565b5073ffffffffffffffffffffffffffffffffffffffff600154163314156114d4565b346102cc5760206003193601126102cc576113436115276115226143cb565b614db5565b604051918291602083526020830190614371565b346102cc5760606003193601126102cc5760043567ffffffffffffffff81116102cc5760a060031982360301126102cc576115746143f3565b9060443567ffffffffffffffff81116102cc57611595903690600401614662565b5061159e614cb8565b50608481016115ac81614a1b565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611bf55750602481019077ffffffffffffffff0000000000000000000000000000000061161383614a3c565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611afe5760a05191611bc6575b50611b9a576116a360448201614a1b565b7f0000000000000000000000000000000000000000000000000000000000000000611b3a575b5067ffffffffffffffff6116dc83614a3c565b166116f4816000526007602052604060002054151590565b15611b0b57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611afe5760a05190611a9b575b73ffffffffffffffffffffffffffffffffffffffff9150163303611a6b57606481013561ffff841680151580611a5c575b156119975761ffff600b54169081811061196557505061186661195b946118f793611522937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff6117e089614a3c565b16918260a05152600c60205280611835604060a0512073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161650a565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b600401615a99565b9261187081614a3c565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2614a3c565b9060405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526119336040826142d3565b604051926119408461427f565b835260208301526040519283926040845260408401906147f6565b9060208301520390f35b7f7911d95b0000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b5061186661195b946118f793611522937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff6119da89614a3c565b16918260a05152600860205280611a2f604060a0512073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161650a565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a261185e565b5061ffff600b54161515611788565b7f728fe07b0000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506020813d602011611af6575b81611ab5602093836142d3565b810103126102cc575173ffffffffffffffffffffffffffffffffffffffff811681036102cc5773ffffffffffffffffffffffffffffffffffffffff90611757565b3d9150611aa8565b6040513d60a051823e3d90fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b73ffffffffffffffffffffffffffffffffffffffff16611b67816000526003602052604060002054151590565b6116c9577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b611be8915060203d602011611bee575b611be081836142d3565b810190614f5e565b84611692565b503d611bd6565b611c1373ffffffffffffffffffffffffffffffffffffffff91614a1b565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b346102cc5760206003193601126102cc5767ffffffffffffffff611c656143cb565b611c6d614d24565b501660a0515260086020526113436112ea6112e56002604060a0512001614d4f565b346102cc5767ffffffffffffffff611ca636614734565b929091611cb1614f13565b1690611cca826000526007602052604060002054151590565b15611d8c578160a051526008602052611cfd6005604060a0512001611cf036868561462b565b60208151910120906162c7565b15611d45577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611d3c604051928392602084526020840191614ce5565b0390a260a05180f35b611d88906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614ce5565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102cc5760a0516003193601126102cc5760a051506040516002548082526020820190600260a05152602060a051209060a0515b818110611e145761134385611e08818703826142d3565b60405191829182614775565b8254845260209093019260019283019201611df1565b346102cc5760206003193601126102cc5767ffffffffffffffff611e4c6143cb565b1660a051526008602052611e676005604060a0512001615fa3565b8051907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611ead611e9784614a8e565b93611ea560405195866142d3565b808552614a8e565b0160a0515b818110611f8e57505060a0515b8151811015611f095780611ed560019284614cd1565b5160a051526009602052611eed604060a0512061456b565b611ef78286614cd1565b52611f028185614cd1565b5001611ebf565b826040518091602082016020835281518091526040830190602060408260051b86010193019160a051905b828210611f4357505050500390f35b91936020611f7e827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614371565b9601920192018594939192611f34565b806060602080938701015201611eb2565b346102cc5760206003193601126102cc5760043567ffffffffffffffff81116102cc5760a060031982360301126102cc57611fd8614cb8565b5060848101611fe681614a1b565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611bf55750602481019077ffffffffffffffff0000000000000000000000000000000061204d83614a3c565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611afe5760a05191612567575b50611b9a576120dd60448201614a1b565b7f0000000000000000000000000000000000000000000000000000000000000000612507575b5067ffffffffffffffff61211683614a3c565b1661212e816000526007602052604060002054151590565b15611b0b57602073ffffffffffffffffffffffffffffffffffffffff60045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611afe5760a051906124a4575b73ffffffffffffffffffffffffffffffffffffffff9150163303611a6b576064013560a0516000146123f15761ffff600b5416806123bc575067ffffffffffffffff6121dc83614a3c565b168060a05152600c6020527f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac209793288280612252604060a0512073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161650a565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690813b156102cc576040517f42966c6800000000000000000000000000000000000000000000000000000000815281600482015260a0518160248160a051875af18015611afe576123a3575b61134361237361152286867ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108767ffffffffffffffff61233d85614a3c565b6040805173ffffffffffffffffffffffffffffffffffffffff9096168652336020870152850192909252169180606081016118ef565b604051906123808261427f565b815261238a614494565b60208201526040519182916020835260208301906147f6565b60a0516123af916142d3565b60a0516102cc57836122fe565b7f7911d95b0000000000000000000000000000000000000000000000000000000060a0515260a051600452602452604460a051fd5b67ffffffffffffffff61240383614a3c565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280612477604060002073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161650a565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a261227b565b506020813d6020116124ff575b816124be602093836142d3565b810103126102cc575173ffffffffffffffffffffffffffffffffffffffff811681036102cc5773ffffffffffffffffffffffffffffffffffffffff90612191565b3d91506124b1565b73ffffffffffffffffffffffffffffffffffffffff16612534816000526003602052604060002054151590565b612103577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b612580915060203d602011611bee57611be081836142d3565b836120cc565b346102cc5760606003193601126102cc5760043567ffffffffffffffff81116102cc576125b7903690600401614680565b9060243567ffffffffffffffff81116102cc576125d89036906004016147c5565b9060443567ffffffffffffffff81116102cc576125f99036906004016147c5565b73ffffffffffffffffffffffffffffffffffffffff600a5416331415806126c1575b61124a578386148015906126b7575b61268b5760a0515b86811061263f5760a05180f35b80612685612653610a3d6001948b8b614c22565b61265e838989614ca8565b61267f61267761266f86898b614ca8565b92369061483d565b91369061483d565b916158d7565b01612632565b7f568efce20000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b508086141561262a565b5073ffffffffffffffffffffffffffffffffffffffff6001541633141561261b565b346102cc5760a0516003193601126102cc57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346102cc5760c06003193601126102cc576127316141b1565b5061273a6143b4565b6127426143e2565b5060843567ffffffffffffffff81116102cc57612763903690600401614413565b505060a4359060028210156102cc5761134391611e089160443590614c32565b346102cc5760206003193601126102cc5760206127be67ffffffffffffffff6127aa6143cb565b166000526007602052604060002054151590565b6040519015158152f35b346102cc5760206003193601126102cc577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174602073ffffffffffffffffffffffffffffffffffffffff6128196141b1565b612821614f13565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a160a05180f35b346102cc5760a0516003193601126102cc5760a0515473ffffffffffffffffffffffffffffffffffffffff81163303612904577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660a0515573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060a05160a051a360a05180f35b7f02b543c60000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346102cc5760a0516003193601126102cc576004546005546040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915290f35b346102cc5760a0516003193601126102cc57602073ffffffffffffffffffffffffffffffffffffffff600a5416604051908152f35b346102cc5760406003193601126102cc576129bf6141b1565b6024356129ca614f13565b73ffffffffffffffffffffffffffffffffffffffff82169182156109bc577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045581600555612a69604051928392836020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0390a160a05180f35b346102cc57612a8036614734565b612a8b929192614f13565b67ffffffffffffffff8216612aad816000526007602052604060002054151590565b15612ac85750610a1c92612ac291369161462b565b906154d3565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102cc5760406003193601126102cc5760043567ffffffffffffffff81116102cc57612b28903690600401614703565b60243567ffffffffffffffff81116102cc57612b48903690600401614680565b919092612b53614f13565b60a0515b828110612bcf5750505060a0515b818110612b725760a05180f35b8067ffffffffffffffff612b8c610a3d6001948688614c22565b168060a05152600f60205260a051604060a05120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee860a05160a051a201612b65565b612bdd610a3d828585614bb2565b612be8828585614bb2565b90602082019060a0830161271061ffff612c0183614bf1565b161015612e635760c084019161271061ffff612c1c85614bf1565b161015612e275767ffffffffffffffff16938460a05152600f60205260a05160409020612c4885614c00565b63ffffffff16908054906040840191612c6083614c00565b60201b67ffffffff0000000016936060860194612c7c86614c00565b60401b6bffffffff0000000000000000169660800196612c9b88614c00565b60601b6fffffffff0000000000000000000000001691612cba8a614bf1565b60801b71ffff000000000000000000000000000000001693612cdb8c614bf1565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717905560405195612d9290614c11565b63ffffffff168652612da390614c11565b63ffffffff166020860152612db790614c11565b63ffffffff166040850152612dcb90614c11565b63ffffffff166060840152612ddf90614404565b61ffff166080830152612df190614404565b61ffff1660a082015260c07f8657880088af4e012a270ce48ac455fdcd83841e09edec744816e9b77d2d73d491a2600101612b57565b61ffff612e3384614bf1565b7f95f3517a0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b612e3361ffff91614bf1565b346102cc5760a0516003193601126102cc576020612e8b614afa565b604051908152f35b346102cc5760406003193601126102cc5760043561ffff8116908190036102cc5760243567ffffffffffffffff81116102cc577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e91612f33612efb6020933690600401614703565b90612f04614f13565b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b556151a1565b604051908152a160a05180f35b346102cc57612f68612f70612f54366146b1565b9491612f61939193614f13565b3691614aa6565b923691614aa6565b7f0000000000000000000000000000000000000000000000000000000000000000156130925760a0515b825181101561300d578073ffffffffffffffffffffffffffffffffffffffff612fc560019386614cd1565b5116612fd081616006565b612fdc575b5001612f9a565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612fd5565b5060a0515b8151811015610a1c578073ffffffffffffffffffffffffffffffffffffffff61303d60019385614cd1565b5116801561308c5761304e816163fb565b61305b575b505b01613012565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183613053565b50613055565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346102cc5760406003193601126102cc576130d76143cb565b60243567ffffffffffffffff81116102cc576020916130fd6127be923690600401614662565b90614a51565b346102cc5760406003193601126102cc5760043567ffffffffffffffff81116102cc578060040161010060031983360301126102cc576131416143f3565b9060405161314e81614263565b60a051905261317f61317561317061316960c487018561498f565b369161462b565b614f76565b606485013561506a565b916084840161318d81614a1b565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611bf55750602484019177ffffffffffffffff000000000000000000000000000000006131f484614a3c565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611afe5760a051916135c4575b50611b9a5767ffffffffffffffff61328a84614a3c565b166132a2816000526007602052604060002054151590565b15611b0b57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611afe5760a051916135a5575b5015611a6b5761331b83614a3c565b9061333160a48701926130fd613169858561498f565b1561355e57505067ffffffffffffffff61343a613434604460209761ffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0961615156000146134af578461338588614a3c565b168060a05152600d8a527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61589806133fa604060a0512073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161650a565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a25b019461342e86614a1b565b50614a3c565b93614a1b565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081168252336020830152909216908201526060810185905292169180608081015b0390a2806040516134a681614263565b52604051908152f35b846134b988614a3c565b168060a0515260088a527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c89806135316002604060a051200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001692839161650a565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2613423565b613568925061498f565b611d886040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614ce5565b6135be915060203d602011611bee57611be081836142d3565b8661330c565b6135dd915060203d602011611bee57611be081836142d3565b86613273565b346102cc5760a0516003193601126102cc57611343611527614494565b346102cc5760206003193601126102cc5760043567ffffffffffffffff81116102cc57806004019061010060031982360301126102cc5760405161364381614263565b60a051905260648101356084820161365a81614a1b565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116911603611bf55750602482019177ffffffffffffffff000000000000000000000000000000006136c184614a3c565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611afe5760a05191613aeb575b50611b9a5767ffffffffffffffff61375784614a3c565b1661376f816000526007602052604060002054151590565b15611b0b57602073ffffffffffffffffffffffffffffffffffffffff60045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611afe5760a05191613acc575b5015611a6b576137e883614a3c565b6137fd60a48301916130fd613169848961498f565b15613ac2575081929360a0515067ffffffffffffffff61381c86614a3c565b168060a0515260086020526138726002604060a051200173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001695869161650a565b6040805173ffffffffffffffffffffffffffffffffffffffff86168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a260206138c7601054614441565b146139d8575b50604401926138db84614a1b565b823b156102cc576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a05173ffffffffffffffffffffffffffffffffffffffff90921660048201526024810185905290818060448101038160a051875af1948515611afe57613496856139886134347ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09667ffffffffffffffff9660209b6139c65750614a3c565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b60a0516139d2916142d3565b8b61342e565b6139e560e483018261498f565b8101906040818303126102cc57803567ffffffffffffffff81116102cc5782613a0f918301614662565b60208201359167ffffffffffffffff83116102cc576020938493613a339201614662565b50613a47604051918281519485920161434e565b8060a051928101039060025afa15611afe5760a051519060c4830190613a76613a70838361498f565b906149e0565b8303613a835750506138cd565b613a9091613a709161498f565b7f7f2493110000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b613568908561498f565b613ae5915060203d602011611bee57611be081836142d3565b856137d9565b613b04915060203d602011611bee57611be081836142d3565b85613740565b346102cc5760a06003193601126102cc57613b236141b1565b50613b2c6143b4565b60443567ffffffffffffffff81116102cc5760031960a091360301126102cc57613b546143e2565b506084359067ffffffffffffffff82116102cc57613b7f67ffffffffffffffff923690600401614413565b5050604051613b8d81614218565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a080519101521660a05152600f60205260c0604060a0512061ffff60405191613bdb83614218565b548163ffffffff82169384815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c1685528760a06080890198828c60801c168a52019960901c1689526040519a8b52511660208a0152511660408801525116606086015251166080840152511660a0820152f35b346102cc5760c06003193601126102cc57613c706141b1565b50613c796143b4565b613c816141d4565b5060843561ffff811681036102cc5760a4359067ffffffffffffffff82116102cc5763ffffffff613cc861ffff92608095613cc184963690600401614413565b5050614889565b9392959091604051968752166020860152166040840152166060820152f35b346102cc5760a0516003193601126102cc57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102cc5760206003193601126102cc576020613d416141b1565b6040517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081169216919091148152f35b346102cc5760a0516003193601126102cc57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102cc5760a0516003193601126102cc576040805161134391613dfd90826142d3565b601a81527f4d6f636b4532454c425443546f6b656e506f6f6c20312e352e310000000000006020820152604051918291602083526020830190614371565b346102cc5760206003193601126102cc57613e546141b1565b613e5c614f13565b613e64614afa565b9081613e705760a05180f35b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff831660248301526044808301859052825290613f8590613ed16064826142d3565b60408051909390917f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169190613f1c86856142d3565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260a0519160a05191519060a051855af13d15614078573d91613f6783614314565b92613f74875194856142d3565b835260a0513d90602085013e616751565b805180613fd7575b505073ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992602092519485521692a28080610a1c565b90602080613fe9938301019101614f5e565b15613ff5578380613f8d565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091616751565b346102cc5760206003193601126102cc57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102cc57817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115614187575b811561415d575b8115614133575b8115614109575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483614102565b7f0e64dd2900000000000000000000000000000000000000000000000000000000811491506140fb565b7fdc0cbd3600000000000000000000000000000000000000000000000000000000811491506140f4565b7faff2afbf00000000000000000000000000000000000000000000000000000000811491506140ed565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610a1757565b6064359073ffffffffffffffffffffffffffffffffffffffff82168203610a1757565b359073ffffffffffffffffffffffffffffffffffffffff82168203610a1757565b60c0810190811067ffffffffffffffff82111761423457604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761423457604052565b6040810190811067ffffffffffffffff82111761423457604052565b60a0810190811067ffffffffffffffff82111761423457604052565b6060810190811067ffffffffffffffff82111761423457604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761423457604052565b67ffffffffffffffff811161423457601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b8381106143615750506000910152565b8181015183820152602001614351565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020936143ad8151809281875287808801910161434e565b0116010190565b6024359067ffffffffffffffff82168203610a1757565b6004359067ffffffffffffffff82168203610a1757565b6064359061ffff82168203610a1757565b6024359061ffff82168203610a1757565b359061ffff82168203610a1757565b9181601f84011215610a175782359167ffffffffffffffff8311610a175760208381860195010111610a1757565b90600182811c9216801561448a575b602083101461445b57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614450565b60405190600082601054916144a883614441565b808352926001811690811561452e57506001146144ce575b6144cc925003836142d3565b565b506010600090815290917f1b6847dc741a1b0cd08d278845f9d819d87b734759afb55fe2de5cb82a9ae6725b8183106145125750509060206144cc928201016144c0565b60209193508060019154838589010152019101909184926144fa565b602092506144cc9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b8201016144c0565b906040519182600082549261457f84614441565b80845293600181169081156145eb57506001146145a4575b506144cc925003836142d3565b90506000929192526020600020906000915b8183106145cf5750509060206144cc9282010138614597565b60209193508060019154838589010152019101909184926145b6565b602093506144cc9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614597565b92919261463782614314565b9161464560405193846142d3565b829481845281830111610a17578281602093846000960137010152565b9080601f83011215610a175781602061467d9335910161462b565b90565b9181601f84011215610a175782359167ffffffffffffffff8311610a17576020808501948460051b010111610a1757565b6040600319820112610a175760043567ffffffffffffffff8111610a1757816146dc91600401614680565b929092916024359067ffffffffffffffff8211610a17576146ff91600401614680565b9091565b9181601f84011215610a175782359167ffffffffffffffff8311610a175760208085019460e08502010111610a1757565b906040600319830112610a175760043567ffffffffffffffff81168103610a1757916024359067ffffffffffffffff8211610a17576146ff91600401614413565b602060408183019282815284518094520192019060005b8181106147995750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161478c565b9181601f84011215610a175782359167ffffffffffffffff8311610a175760208085019460608502010111610a1757565b61467d91602061480f8351604084526040840190614371565b920151906020818403910152614371565b35906fffffffffffffffffffffffffffffffff82168203610a1757565b9190826060910312610a1757604051614855816142b7565b80928035908115158203610a17576040614884918193855261487960208201614820565b602086015201614820565b910152565b67ffffffffffffffff16600052600f6020526040600020604051906148ad82614218565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c16815261ffff6080860196818460801c1688528160a088019460901c168452168061492c5750505063ffffffff808061ffff9351169451169551169351169193929190565b919550915061ffff600b54169081811061495f57505063ffffffff808061ffff9351169451169551169351169193929190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610a17570180359067ffffffffffffffff8211610a1757602001918136038313610a1757565b3590602081106149ee575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b3573ffffffffffffffffffffffffffffffffffffffff81168103610a175790565b3567ffffffffffffffff81168103610a175790565b9067ffffffffffffffff61467d92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116142345760051b60200190565b929190614ab281614a8e565b93614ac060405195866142d3565b602085838152019160051b8101928311610a1757905b828210614ae257505050565b60208091614aef846141f7565b815201910190614ad6565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115614ba657600091614b77575090565b90506020813d602011614b9e575b81614b92602093836142d3565b81010312610a17575190565b3d9150614b85565b6040513d6000823e3d90fd5b9190811015614bc25760e0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3561ffff81168103610a175790565b3563ffffffff81168103610a175790565b359063ffffffff82168203610a1757565b9190811015614bc25760051b0190565b67ffffffffffffffff16600052600e6020526040600020916002811015614c7957600114614c685781600161467d9301906157ab565b816002600361467d940191016157ab565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9190811015614bc2576060020190565b60405190614cc58261427f565b60606020838281520152565b8051821015614bc25760209160051b010190565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b60405190614d318261429b565b60006080838281528260208201528260408201528260608201520152565b90604051614d5c8161429b565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff16600052600860205261467d600460406000200161456b565b9190811015614bc25760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610a17570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610a17570180359067ffffffffffffffff8211610a1757602001918160051b36038313610a1757565b81810292918115918404141715614e7e57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110614eb8575050565b60008155600101614ead565b9160209082815201919060005b818110614ede5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff614f05886141f7565b168152019401929101614ed1565b73ffffffffffffffffffffffffffffffffffffffff600154163303614f3457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b90816020910312610a1757518015158103610a175790565b80518015614fe657602003614fa8578051602082810191830183900312610a1757519060ff8211614fa8575060ff1690565b611d88906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614371565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff8211614e7e57565b60ff16604d8111614e7e57600a0a90565b811561503b570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146151705782841161514657906150af9161500c565b91604d60ff841611801561510d575b6150d7575050906150d161467d92615020565b90614e6b565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061511783615020565b801561503b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0484116150be565b61514f9161500c565b91604d60ff8416116150d75750509061516a61467d92615020565b90615031565b5050505090565b358015158103610a175790565b356fffffffffffffffffffffffffffffffff81168103610a175790565b919060005b8181106151b35750509050565b6151be818386614bb2565b67ffffffffffffffff6151d082614a3c565b166151e8816000526007602052604060002054151590565b156154a6578161525b6152bb926152616020600197960161521161520c368361483d565b615c29565b61525b60406000858152600c6020522091825463ffffffff8160801c16159081615488575b81615479575b8161545e575b8161544f575b5080615440575b6153b4575b369061483d565b90615d70565b6040608084019161527561520c368561483d565b6000908152600d6020522092835463ffffffff8160801c16159081615396575b81615387575b8161536c575b8161535d575b508061534e575b6152c1575b50369061483d565b016151a6565b6152de60a06fffffffffffffffffffffffffffffffff9201615184565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355386152b3565b5061535882615177565b6152ae565b60ff915060a01c1615386152a7565b6fffffffffffffffffffffffffffffffff81161591506152a1565b8589015460801c15915061529b565b858901546fffffffffffffffffffffffffffffffff16159150615295565b6fffffffffffffffffffffffffffffffff6153d160408901615184565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355615254565b5061544a81615177565b61524f565b60ff915060a01c161538615248565b6fffffffffffffffffffffffffffffffff8116159150615242565b848c015460801c15915061523c565b848c01546fffffffffffffffffffffffffffffffff16159150615236565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b908051156157135767ffffffffffffffff815160208301209216918260005260086020526155088160056040600020016164b5565b156156cf5760005260096020526040600020815167ffffffffffffffff8111614234576155358254614441565b601f811161569d575b506020601f82116001146155d757916155b1827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936155c7956000916155cc575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614371565b0390a2565b905084015138615580565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106156855750926155c79492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061564e575b5050811b019055611527565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690553880615642565b9192602060018192868a015181550194019201615607565b6156c990836000526020600020601f840160051c8101916020851061094657601f0160051c0190614ead565b3861553e565b5090611d886040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614371565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b81811061576f5750506144cc925003836142d3565b845473ffffffffffffffffffffffffffffffffffffffff1683526001948501948794506020909301920161575a565b91908201809211614e7e57565b6157b49061573d565b9160055480151591826158cc575b50506157cc575090565b6157d59061573d565b908151806157e35750905090565b6157ee90825161579e565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061583261581c86614a8e565b9561582a60405197886142d3565b808752614a8e565b0136602086013760005b825181101561587a578073ffffffffffffffffffffffffffffffffffffffff61586760019386614cd1565b51166158738288614cd1565b520161583c565b509160005b81518110156158c7578073ffffffffffffffffffffffffffffffffffffffff6158aa60019385614cd1565b51166158c06158ba83875161579e565b88614cd1565b520161587f565b505050565b1015905038806157c2565b67ffffffffffffffff1660008181526007602052604090205490929190156159d957916159d660e0926159a28561592e7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97615c29565b846000526008602052615945816040600020615d70565b61594e83615c29565b846000526008602052615968836002604060002001615d70565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b91908203918211614e7e57565b615a1c614d24565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691615a796020850193615a73615a6663ffffffff87511642615a07565b8560808901511690614e6b565b9061579e565b80821015615a9257505b16825263ffffffff4216905290565b9050615a83565b9060a061ffff9167ffffffffffffffff615ab560208601614a3c565b16600052600f6020528260406000209160405192615ad284614218565b549263ffffffff8416815263ffffffff8460201c16602082015263ffffffff8460401c16604082015263ffffffff8460601c16606082015282808560801c169485608084015260901c169485910152161515600014615b5857505b168015615b5057612710615b49606061467d9401359283614e6b565b0490615a07565b506060013590565b9050615b2d565b805160005b818110615b7057505050565b60018101808211614e7e575b828110615b8c5750600101615b64565b73ffffffffffffffffffffffffffffffffffffffff615bab8386614cd1565b511673ffffffffffffffffffffffffffffffffffffffff615bcc8387614cd1565b511614615bdb57600101615b7c565b73ffffffffffffffffffffffffffffffffffffffff615bfa8386614cd1565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115615cc9576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff60208301511610615c665750565b606490615cc7604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590615d51575b615cf05750565b606490615cc7604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515615ce9565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615ea96060928054615dad63ffffffff8260801c1642615a07565b9081615ee8575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014615ee057508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155615e5d8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6159d660405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091615de4565b6fffffffffffffffffffffffffffffffff91615f1d839283615f166001880154948286169560801c90614e6b565b911661579e565b80821015615f9c57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880615db4565b9050615f27565b906040519182815491828252602082019060005260206000209260005b818110615fd55750506144cc925003836142d3565b8454835260019485019487945060209093019201615fc0565b8054821015614bc25760005260206000200190600090565b6000818152600360205260409020548015616195577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614e7e57600254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614e7e57818103616126575b50505060025480156160f7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016160b4816002615fee565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61617d616137616148936002615fee565b90549060031b1c9283926002615fee565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600360205260406000205538808061607b565b5050600090565b6000818152600760205260409020548015616195577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614e7e57600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614e7e5781810361628d575b50505060065480156160f7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161624a816006615fee565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b6162af61629e616148936006615fee565b90549060031b1c9283926006615fee565b90556000526007602052604060002055388080616211565b90600182019181600052826020526040600020548015156000146163f2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111614e7e578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211614e7e578181036163bb575b505050805480156160f7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061637c8282615fee565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b6163db6163cb6161489386615fee565b90549060031b1c92839286615fee565b905560005283602052604060002055388080616344565b50505050600090565b8060005260036020526040600020541560001461645557600254680100000000000000008110156142345761643c6161488260018594016002556002615fee565b9055600254906000526003602052604060002055600190565b50600090565b8060005260076020526040600020541560001461645557600654680100000000000000008110156142345761649c6161488260018594016006556006615fee565b9055600654906000526007602052604060002055600190565b6000828152600182016020526040902054616195578054906801000000000000000082101561423457826164f3616148846001809601855584615fee565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015616749575b616743576fffffffffffffffffffffffffffffffff8216916001850190815461656263ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615a07565b90816166a5575b505084811061665957508383106165c35750506165986fffffffffffffffffffffffffffffffff928392615a07565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916165d28185615a07565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190808211614e7e576166206166259273ffffffffffffffffffffffffffffffffffffffff9661579e565b615031565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611616719576166c092615a739160801c90614e6b565b808410156167145750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880616569565b6166cb565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561651d565b919290156167cc5750815115616765575090565b3b1561676e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156167df5750805190602001fd5b611d88906040519182917f08c379a000000000000000000000000000000000000000000000000000000000835260206004840152602483019061437156fea164736f6c634300081a000a",
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "applyCustomBlockConfirmationConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_MockE2ELBTCTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setCustomBlockConfirmationRateLimitConfig", rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_MockE2ELBTCTokenPool.TransactOpts, rateLimitConfigArgs)
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

type MockE2ELBTCTokenPoolCustomBlockConfirmationUpdatedIterator struct {
	Event *MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated)
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
		it.Event = new(MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated)
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

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolCustomBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolCustomBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolCustomBlockConfirmationUpdatedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "CustomBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseCustomBlockConfirmationUpdated(log types.Log) (*MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated, error) {
	event := new(MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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
type GetFee struct {
	FeeUSDCents       *big.Int
	DestGasOverhead   uint32
	DestBytesOverhead uint32
	TokenFeeBps       uint16
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

func (MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e")
}

func (MockE2ELBTCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
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
	return common.HexToHash("0x8657880088af4e012a270ce48ac455fdcd83841e09edec744816e9b77d2d73d4")
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

	GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

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

	ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

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

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*MockE2ELBTCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolCustomBlockConfirmationUpdatedIterator, error)

	WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomBlockConfirmationUpdated(log types.Log) (*MockE2ELBTCTokenPoolCustomBlockConfirmationUpdated, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*MockE2ELBTCTokenPoolDynamicConfigSet, error)

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
