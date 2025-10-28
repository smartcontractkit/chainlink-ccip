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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"struct SiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmationConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610120806040523461040d57617a53803803809161001d82856106c3565b8339810160c08282031261040d5781516001600160a01b0381169290919083830361040d5761004e602082016106e6565b60408201516001600160401b03811161040d5782019280601f8501121561040d578351936001600160401b038511610412578460051b90602082019561009760405197886106c3565b865260208087019282010192831161040d57602001905b8282106106ab575050506100c4606083016106f4565b6100dc60a06100d5608086016106f4565b94016106f4565b94331561069a57600180546001600160a01b0319163317905586158015610689575b8015610678575b6104de5760805260c05260405163313ce56760e01b8152602081600481895afa6000918161063c575b50610611575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e08190526104ef575b506001600160a01b031680156104de57604051636eb1769f60e11b815230600482015260248101829052602081604481865afa9081156104d2576000916104a0575b506104355760405191602083019263095ea7b360e01b84528260248201526000196044820152604481526101dd6064826106c3565b6000806040958651936101f088866106c3565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610428573d906001600160401b038211610412578551610261949092610252601f8201601f1916602001856106c3565b83523d6000602085013e6108a8565b805180610391575b505061010052516170da908161097982396080518181816104000152818161199b01528181611db90152818161203d01528181612175015281816127120152818161294301528181612b9501528181613a8201528181613cdc01528181613d8c01528181613f02015281816140c4015281816145cc0152818161476c015281816147ba0152818161491c01528181614a290152615399015260a051818181614722015281816159c901528181615a4c015261648d015260c051818181610eba01528181611e47015281816127a001528181613b100152613f91015260e051818181610e7401528181611e8c015281816127e5015261387501526101005181818161044b0152818161194001528181612a450152818161418d0152818161460701526149ce0152f35b816020918101031261040d576020015180159081150361040d576103b6573880610269565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b91610261926060916108a8565b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d6020116104ca575b816104bb602093836106c3565b8101031261040d5751386101a8565b3d91506104ae565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b906020906040519061050183836106c3565b60008252600036813760e051156106005760005b825181101561057c576001906001600160a01b036105338286610708565b51168561053f8261074a565b61054c575b505001610515565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13885610544565b5092905060005b81518110156105f7576001906001600160a01b036105a18285610708565b511680156105f157846105b382610848565b6105c1575b50505b01610583565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138846105b8565b506105bb565b50505038610166565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036106255750610134565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d602011610670575b81610658602093836106c3565b8101031261040d57610669906106e6565b903861012e565b3d915061064b565b506001600160a01b03821615610105565b506001600160a01b038416156100fe565b639b15e16f60e01b60005260046000fd5b602080916106b8846106f4565b8152019101906100ae565b601f909101601f19168101906001600160401b0382119082101761041257604052565b519060ff8216820361040d57565b51906001600160a01b038216820361040d57565b805182101561071c5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561071c5760005260206000200190600090565b600081815260036020526040902054801561084157600019810181811161082b5760025460001981019190821161082b578181036107da575b50505060025480156107c4576000190161079e816002610732565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6108136107eb6107fc936002610732565b90549060031b1c9283926002610732565b819391549060031b91821b91600019901b19161790565b90556000526003602052604060002055388080610783565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146108a25760025468010000000000000000811015610412576108896107fc8260018594016002556002610732565b9055600254906000526003602052604060002055600190565b50600090565b9192901561090a57508151156108bc575090565b3b156108c55790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561091d5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b8381106109605750508160006044809484010152601f80199101168101030190fd5b6020828201810151604487840101528593500161093e56fe60c080604052600436101561001357600080fd5b600060a05260a0513560e01c90816301ffc9a714614b73575080630574d6a614614ae35780630a861f2a14614950578063164e68de14614866578063181f5a77146147de57806321df0da714614799578063240028e81461474657806324f65ee7146147075780632d4a148f146144f457806331238ffc146144ac57806337b192471461433e5780633907753714613e81578063432a6ba314613e59578063489a68f2146139ec5780634c5ef0ed146139a757806354c8a4f31461384357806359152aad146137965780635df45a371461377a57806362ddd3c4146136f55780636600f92c146135fb578063698c2c66146135525780636cfd1553146134a95780636d3d1a58146134815780636d9d216c1461306d5780637437ff9f1461303957806379ba509714612f7c5780637d54534e14612ef85780638632d5cc14612ec45780638926f54f14612e7f57806389720a6214612e145780638da5cb5b14612dec578063962d402014612ca95780639a4575b9146126b7578063a42a7b8b14612560578063a7cd63b7146124f2578063acfecf91146123ce578063af0e58b9146123af578063af58d59f14612363578063b1c71c6514611d3c578063b794658014611d04578063bb6bb5a714611c9c578063c4bffe2b14611b80578063c75eea9c14611ad5578063ce3c752814611887578063cf7401f31461172a578063d966866b146112dd578063dc04fa1f14610ede578063dc0bd97114610e99578063e0351e1314610e5b578063e8a1da17146105a3578063eb521a4c1461035c578063f1e73399146103315763f2fde38b1461026c57600080fd5b3461032b57602060031936011261032b576001600160a01b0361028d614cfe565b610295615917565b163381146102ff5760a05180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360a05180f35b7fdad89dca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b60a05180fd5b3461032b57602060031936011261032b57602061035461034f614ca4565b615890565b604051908152f35b3461032b57602060031936011261032b576004358015610577576001600160a01b0361038960a051615464565b1633036105475760a05160a051526012602052604060a0512060ff600182015460a01c16600014610532576103bf828254615447565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f0000000000000000000000000000000000000000000000000000000000000000906104419061043b81608481015b03601f198101835282614e83565b826166ab565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561032b576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a080516001600160a01b039390931660048301526024820185905251909283916044918391905af180156105255761050c575b506040519060a0515060a051825260208201527f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf60403392a260a05180f35b60a05161051891614e83565b60a05161032b57816104cd565b6040513d60a051823e3d90fd5b5061053f81601054615447565b6010556103c2565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b7fa90c0d190000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461032b576105b136614f48565b9190926105bc615917565b60a051905b828210610c995750505060a0519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610c9357600581901b8301358581121561032b5783016101208136031261032b576040519461063086614e67565b813567ffffffffffffffff81168103610c8e578652602082013567ffffffffffffffff811161032b5782019436601f8701121561032b578535610672816152f1565b966106806040519889614e83565b81885260208089019260051b8201019036821161032b5760208101925b828410610c5f575050505060208701958652604083013567ffffffffffffffff811161032b576106d09036908501614ef9565b91604088019283526106fa6106e836606087016150d4565b9460608a0195865260c03691016150d4565b956080890196875261070c8551616564565b6107168751616564565b83515115610c335761073267ffffffffffffffff8a5116616f52565b15610bf85767ffffffffffffffff89511660a051526008602052604060a0512061087686516fffffffffffffffffffffffffffffffff604082015116906108316fffffffffffffffffffffffffffffffff6020830151169151151583608060405161079c81614e67565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b61099c88516fffffffffffffffffffffffffffffffff604082015116906109576fffffffffffffffffffffffffffffffff602083015116915115158360806040516108c081614e67565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff8211610bc7576109bf8354615559565b601f8111610b88575b506020906001601f841114610b1e579180916109fb9360a05192610b13575b50506000198260011b9260031b1c19161790565b90555b60a0515b88518051821015610a375790610a31600192610a2a8367ffffffffffffffff8f511692615545565b5190615e93565b01610a02565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610b0567ffffffffffffffff6001979694985116925193519151610ad1610a9c60405196879687526101006020880152610100870190614d87565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a10193929091936105fe565b015190508e806109e7565b90601f198316918460a051528160a051209260a0515b818110610b705750908460019594939210610b57575b505050811b0190556109fe565b015160001960f88460031b161c191690558d8080610b4a565b92936020600181928786015181550195019301610b34565b610bb7908460a05152602060a05120601f850160051c81019160208610610bbd575b601f0160051c01906157e9565b8d6109c8565b9091508190610baa565b7f4e487b710000000000000000000000000000000000000000000000000000000060a051526041600452602460a051fd5b67ffffffffffffffff8951167f1d5ad3c50000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b833567ffffffffffffffff811161032b57602091610c838392833691870101614ef9565b81520193019261069d565b600080fd5b60a05180f35b9092919367ffffffffffffffff610cb9610cb4868886615408565b61524f565b1692610cc484616d83565b15610e2b578360a051526008602052610ce46005604060a0512001616c20565b9260a0515b8451811015610d23576001908660a051526008602052610d1c6005604060a0512001610d158389615545565b5190616e36565b5001610ce9565b50939094919592508060a0515260086020526005604060a0512060a051815560a051600182015560a051600282015560a051600382015560048101610d688154615559565b80610ddb575b50500180549060a051815581610db7575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916105c1565b60a05152602060a05120908101905b81811015610d7f5760a0518155600101610dc6565b601f8111600114610df4575060a05190555b8880610d6e565b610e14908260a051526001601f602060a051209201861c820191016157e9565b60a080518290525160208120918190559055610ded565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461032b5760a05160031936011261032b5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461032b5760a05160031936011261032b5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461032b57604060031936011261032b5760043567ffffffffffffffff811161032b573660238201121561032b57806004013567ffffffffffffffff811161032b576024820191602436918360081b01011161032b5760243567ffffffffffffffff811161032b57610f54903690600401614f17565b919092610f5f615917565b60a0515b828110610fdb5750505060a0515b818110610f7e5760a05180f35b8067ffffffffffffffff610f98610cb46001948688615408565b168060a05152600f60205260a051604060a05120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee860a05160a051a201610f71565b610fe9610cb4828585615842565b610ff4828585615842565b602081019060a081019261271061ffff61100d86615852565b1610156112d15760c082019061271061ffff61102884615852565b1610156112955767ffffffffffffffff16938460a05152600f60205260a051604090209261105585615861565b63ffffffff16845491604081019061106c82615861565b60201b67ffffffff000000001693606082019361108885615861565b60401b6bffffffff0000000000000000169560808401966110a888615861565b60601b6fffffffff00000000000000000000000016916110c78a615852565b60801b71ffff0000000000000000000000000000000016936110e88c615852565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717875560e0019561119f87615872565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff000000000000000000000000000000000000000016179055604051966111f09061587f565b63ffffffff1687526112019061587f565b63ffffffff1660208701526112159061587f565b63ffffffff1660408601526112299061587f565b63ffffffff16606085015261123d90614d4a565b61ffff16608084015261124f90614d4a565b61ffff1660a0830152611261906150aa565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610f63565b61ffff6112a183615852565b7f95f3517a0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b61ffff6112a185615852565b3461032b57602060031936011261032b5760043567ffffffffffffffff811161032b5761130e903690600401614f17565b90611317615917565b60a0516080525b816080511061132d5760a05180f35b61133d610cb46080518484615742565b61135761134d6080518585615742565b6020810190615782565b6113716113676080518787615742565b6040810190615782565b9061138c6113826080518989615742565b6060810190615782565b6113a661139c6080518b8b615742565b6080810190615782565b9390946113bc6113b736898b615309565b6164c1565b6113ca6113b7368385615309565b6113d86113b7368587615309565b6113e66113b7368789615309565b6040519860808a01908a821067ffffffffffffffff831117610bc75767ffffffffffffffff9160405261141a368a8c615309565b8b52611427368486615309565b60208c0152611437368688615309565b60408c015261144736888a615309565b60608c015216988960a05152600e602052604060a05120815180519067ffffffffffffffff8211610bc757680100000000000000008211610bc757602090835483855580841061170b575b50018260a05152602060a0512060a0515b8381106116ee57505050506001810160208301519081519167ffffffffffffffff8311610bc757680100000000000000008311610bc75760209082548484558085106116cf575b50019060a05152602060a0512060a0515b8381106116b257505050506002810160408301519081519167ffffffffffffffff8311610bc757680100000000000000008311610bc7576020908254848455808510611693575b50019060a05152602060a0512060a0515b838110611676575050505060036060910191015180519067ffffffffffffffff8211610bc757680100000000000000008211610bc7576020908354838555808410611657575b50019160a05152602060a051209160a0515b82811061163a575050505092611629949261160d61161b937fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9998966115ff6040519b8c9b60808d5260808d0191615800565b918a830360208c0152615800565b918783036040890152615800565b918483036060860152615800565b0390a260016080510160805261131e565b60019060206001600160a01b0384511693019281860155016115ab565b611670908560a05152848460a0512091820191016157e9565b8f611599565b60019060206001600160a01b038551169401938184015501611553565b6116ac908460a05152858460a0512091820191016157e9565b38611542565b60019060206001600160a01b0385511694019381840155016114fb565b6116e8908460a05152858460a0512091820191016157e9565b386114ea565b60019060206001600160a01b0385511694019381840155016114a3565b611724908560a05152848460a0512091820191016157e9565b38611492565b3461032b5760e060031936011261032b57611743614ca4565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc011261032b5760405161177981614e2f565b602435801515810361032b5781526044356fffffffffffffffffffffffffffffffff8116810361032b5760208201526064356fffffffffffffffffffffffffffffffff8116810361032b5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c011261032b576040519061180082614e2f565b608435801515810361032b57825260a4356fffffffffffffffffffffffffffffffff8116810361032b57602083015260c4356fffffffffffffffffffffffffffffffff8116810361032b5760408301526001600160a01b03600a541633141580611872575b61054757610c93926161eb565b506001600160a01b0360015416331415611865565b3461032b57604060031936011261032b576118a0614ca4565b60243567ffffffffffffffff82168060a05152601260205260ff6001604060a05120015460a01c16158015611acd575b611a9e578115610577576001600160a01b036118eb84615464565b1633036105475760a051526012602052604060a0512060ff600182015460a01c1660a0515080600014611a965781545b808411611a62575015611a4d57611933828254615264565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1561032b576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af1801561052557611a34575b506040805167ffffffffffffffff9093168352602083019190915233917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e91819081015b0390a260a05180f35b60a051611a4091614e83565b60a05161032b57826119e7565b50611a5a81601054615264565b601055611936565b83907fa17e11d50000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b60105461191b565b7f46f5f12b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b5080156118d0565b3461032b57602060031936011261032b5767ffffffffffffffff611af7614ca4565b611aff61568f565b501660a051526008602052611b7c611b23611b1e604060a051206156ba565b61631b565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b3461032b5760a05160031936011261032b5760a051506040516006548082528160208101600660a05152602060a051209260a0515b818110611c83575050611bca92500382614e83565b805190611bef611bd9836152f1565b92611be76040519485614e83565b8084526152f1565b90601f1960208401920136833760a0515b8151811015611c32578067ffffffffffffffff611c1f60019385615545565b5116611c2b8287615545565b5201611c00565b5050906040519182916020830190602084525180915260408301919060a0515b818110611c60575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611c52565b8454835260019485019486945060209093019201611bb5565b3461032b57602060031936011261032b5760043567ffffffffffffffff811161032b57611ccd903690600401614f9a565b6001600160a01b03600a541633141580611cef575b61054757610c9391615b6d565b506001600160a01b0360015416331415611ce2565b3461032b57602060031936011261032b57611b7c611d28611d23614ca4565b615720565b604051918291602083526020830190614d87565b3461032b57606060031936011261032b5760043567ffffffffffffffff811161032b5760a0600319823603011261032b57611d75614d39565b9060443567ffffffffffffffff811161032b57611d96903690600401614ef9565b50611d9f61552c565b506084810190611dae826152a0565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361232257602481019177ffffffffffffffff00000000000000000000000000000000611e078461524f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105255760a051916122f3575b506122c757611e8a604483016152a0565b7f0000000000000000000000000000000000000000000000000000000000000000612274575b5067ffffffffffffffff611ec38461524f565b16611edb816000526007602052604060002054151590565b156122455760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156105255760a051906121fc575b6001600160a01b0391501633036121cc5760648201359061ffff8516801515806121bd575b156121115761ffff600b5416908181106120df5750506120d5946120a493611d23937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff61202595611fdf611fcf611fb58c61524f565b67ffffffffffffffff16600052600c602052604060002090565b85611fd9846152a0565b916167de565b612019611ff4611fee8c61524f565b926152a0565b94604051938493169583602090939291936001600160a01b0360408201951681520152565b0390a25b6004016163a0565b9261202f8161524f565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a261524f565b906120ad616486565b604051926120ba84614e4b565b83526020830152604051928392604084526040840190615080565b9060208301520390f35b7f7911d95b0000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b50506120256120d5946120a493611d23937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff6121558961524f565b16918260a0515260086020528061219d604060a051206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916167de565b604080516001600160a01b039290921682526020820192909252a261201d565b5061ffff600b54161515611f56565b7f728fe07b0000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506020813d60201161223d575b8161221660209383614e83565b8101031261032b57516001600160a01b038116810361032b576001600160a01b0390611f31565b3d9150612209565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b6001600160a01b0316612294816000526003602052604060002054151590565b611eb0577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b612315915060203d60201161231b575b61230d8183614e83565b810190615b38565b85611e79565b503d612303565b6001600160a01b03612333836152a0565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b3461032b57602060031936011261032b5767ffffffffffffffff612385614ca4565b61238d61568f565b501660a051526008602052611b7c611b23611b1e6002604060a05120016156ba565b3461032b5760a05160031936011261032b576020601054604051908152f35b3461032b5767ffffffffffffffff6123e536614fcb565b9290916123f0615917565b1690612409826000526007602052604060002054151590565b156124c2578160a05152600860205261243c6005604060a051200161242f368685614ec2565b6020815191012090616e36565b1561247b577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611a2b60405192839260208452602084019161566e565b6124be906040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161566e565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461032b5760a05160031936011261032b5760a051506040516002548082526020820190600260a05152602060a051209060a0515b81811061254a57611b7c8561253e81870382614e83565b6040519182918261500c565b8254845260209093019260019283019201612527565b3461032b57602060031936011261032b5767ffffffffffffffff612582614ca4565b1660a05152600860205261259d6005604060a0512001616c20565b805190601f196125c56125af846152f1565b936125bd6040519586614e83565b8085526152f1565b0160a0515b8181106126a657505060a0515b815181101561262157806125ed60019284615545565b5160a051526009602052612605604060a051206155ac565b61260f8286615545565b5261261a8185615545565b50016125d7565b826040518091602082016020835281518091526040830190602060408260051b86010193019160a051905b82821061265b57505050500390f35b91936020612696827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614d87565b960192019201859493919261264c565b8060606020809387010152016125ca565b3461032b57602060031936011261032b5760043567ffffffffffffffff811161032b5760a0600319823603011261032b576126f061552c565b506126f961552c565b5060848101612707816152a0565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612c9857602482019177ffffffffffffffff000000000000000000000000000000006127608461524f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105255760a05191612c79575b506122c7576127e3604482016152a0565b7f0000000000000000000000000000000000000000000000000000000000000000612c26575b5067ffffffffffffffff61281c8461524f565b16612834816000526007602052604060002054151590565b156122455760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156105255760a05190612bdd575b6001600160a01b0391501633036121cc57606401359060a051600014612b435761ffff600b541680612b0e57506128d06128c6611fb58561524f565b83611fd9846152a0565b7f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff61290c6129068661524f565b936152a0565b604080516001600160a01b03929092168252602082018690529190931692a25b6129358261524f565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016808252336020830152918101849052909167ffffffffffffffff16907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a26129b1611d238461524f565b926129ba616486565b604051946129c786614e4b565b8552602085015267ffffffffffffffff6129e08261524f565b1660a05152601260205260ff6001604060a05120015460a01c16600014612af9578067ffffffffffffffff612a17612a399361524f565b1660a051526012602052604060a05120612a32858254615447565b905561524f565b505b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561032b576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390931660048201526024810191909152918290818060448101039160a051905af1801561052557612ae0575b60405160208082528190611b7c90820185615080565b60a051612aec91614e83565b60a05161032b5781612aca565b50612b0682601054615447565b601055612a3b565b7f7911d95b0000000000000000000000000000000000000000000000000000000060a0515260a051600452602452604460a051fd5b5067ffffffffffffffff612b568361524f565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280612bbd60406000206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916167de565b604080516001600160a01b039290921682526020820192909252a261292c565b506020813d602011612c1e575b81612bf760209383614e83565b8101031261032b57516001600160a01b038116810361032b576001600160a01b039061288a565b3d9150612bea565b6001600160a01b0316612c46816000526003602052604060002054151590565b612809577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b612c92915060203d60201161231b5761230d8183614e83565b846127d2565b6123336001600160a01b03916152a0565b3461032b57606060031936011261032b5760043567ffffffffffffffff811161032b57612cda903690600401614f17565b9060243567ffffffffffffffff811161032b57612cfb90369060040161504f565b9060443567ffffffffffffffff811161032b57612d1c90369060040161504f565b6001600160a01b03600a541633141580612dd7575b61054757838614801590612dcd575b612da15760a0515b868110612d555760a05180f35b80612d9b612d69610cb46001948b8b615408565b612d7483898961551c565b612d95612d8d612d8586898b61551c565b9236906150d4565b9136906150d4565b916161eb565b01612d48565b7f568efce20000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b5080861415612d40565b506001600160a01b0360015416331415612d31565b3461032b5760a05160031936011261032b5760206001600160a01b0360015416604051908152f35b3461032b5760c060031936011261032b57612e2d614cfe565b50612e36614cbb565b612e3e614d28565b5060843567ffffffffffffffff811161032b57612e5f903690600401614d59565b505060a43590600282101561032b57611b7c9161253e91604435906154a6565b3461032b57602060031936011261032b576020612eba67ffffffffffffffff612ea6614ca4565b166000526007602052604060002054151590565b6040519015158152f35b3461032b57602060031936011261032b576020612ee7612ee2614ca4565b615464565b6001600160a01b0360405191168152f35b3461032b57602060031936011261032b577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b03612f3c614cfe565b612f44615917565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a160a05180f35b3461032b5760a05160031936011261032b5760a051546001600160a01b038116330361300d577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660a051556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060a05160a051a360a05180f35b7f02b543c60000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461032b5760a05160031936011261032b57600454600554604080516001600160a01b039093168352602083019190915290f35b3461032b57604060031936011261032b5760043567ffffffffffffffff811161032b5761309e903690600401614f17565b6024359167ffffffffffffffff831161032b573660238401121561032b5782600401359167ffffffffffffffff831161032b576024840193602436918560061b01011161032b576130ed615917565b60a0515b8181106133485750505060a0515b81811061310c5760a05180f35b67ffffffffffffffff613123610cb4838587615454565b16158015613311575b80156132f0575b6132a9576001600160a01b03613155602061314f848688615454565b016152a0565b1615610c33578061323e613171602061314f6001958789615454565b856001600160a01b03808660405161318881614e2f565b60a051815282602082019616865267ffffffffffffffff6131b4610cb48a8d6040860199878b52615454565b1660a051526012602052604060a051209051815501935116167fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb61326d610cb4838688615454565b61327d602061314f85888a615454565b6040805167ffffffffffffffff9390931683526001600160a01b0391909116602083015290a1016130ff565b610cb4906132c09267ffffffffffffffff94615454565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b5061330b67ffffffffffffffff612ea6610cb4848688615454565b15613133565b5067ffffffffffffffff613329610cb4838587615454565b1660a05152601260205260ff6001604060a05120015460a01c1661312c565b67ffffffffffffffff61335f610cb4838587615408565b1660a05152601260205260ff6001604060a05120015460a01c161561343a578067ffffffffffffffff613398610cb46001948688615408565b1660a0515260126020527f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93604060a05120546133d681601054615447565b60105567ffffffffffffffff6133f0610cb485888a615408565b60a0805191909216905260126020525160408120818155850155613418610cb4848789615408565b6040805167ffffffffffffffff9290921682526020820192909252a1016130f1565b610cb4906134519267ffffffffffffffff94615408565b7f46f5f12b0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b3461032b5760a05160031936011261032b5760206001600160a01b03600a5416604051908152f35b3461032b57602060031936011261032b577f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b2346001600160a01b036134eb614cfe565b6134f3615917565b613549601154918381167fffffffffffffffffffffffff000000000000000000000000000000000000000084161760115560405193849316839092916001600160a01b0360209181604085019616845216910152565b0390a160a05180f35b3461032b57604060031936011261032b5761356b614cfe565b602435613576615917565b6001600160a01b038216918215610c33577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff000000000000000000000000000000000000000060045416176004558160055561354960405192839283602090939291936001600160a01b0360408201951681520152565b3461032b57604060031936011261032b57613614614ca4565b67ffffffffffffffff613625614cd2565b9161362e615917565b16908160a0515260126020526001604060a051200190815460ff8160a01c16156136c55782547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179093556040805191909216815260208101929092527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f919081908101611a2b565b837f46f5f12b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461032b5761370336614fcb565b61370e929192615917565b67ffffffffffffffff8216613730816000526007602052604060002054151590565b1561374b5750610c9392613745913691614ec2565b90615e93565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461032b5760a05160031936011261032b57602061035461535d565b3461032b57604060031936011261032b5760043561ffff81169081900361032b5760243567ffffffffffffffff811161032b577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e916138366137fe6020933690600401614f9a565b90613807615917565b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b55615b6d565b604051908152a160a05180f35b3461032b5761386b61387361385736614f48565b9491613864939193615917565b3691615309565b923691615309565b7f00000000000000000000000000000000000000000000000000000000000000001561397b5760a0515b825181101561390357806001600160a01b036138bb60019386615545565b51166138c681616c83565b6138d2575b500161389d565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1846138cb565b5060a0515b8151811015610c9357806001600160a01b0361392660019385615545565b511680156139755761393781616ef2565b613944575b505b01613908565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a18361393c565b5061393e565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461032b57604060031936011261032b576139c0614ca4565b60243567ffffffffffffffff811161032b576020916139e6612eba923690600401614ef9565b906152b4565b3461032b57604060031936011261032b5760043567ffffffffffffffff811161032b5780600401610100600319833603011261032b57613a2a614d39565b90604051613a3781614e13565b60a0519052613a68613a5e613a59613a5260c48701856151fe565b3691614ec2565b615955565b6064850135615a49565b916084840190613a77826152a0565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361232257602485019277ffffffffffffffff00000000000000000000000000000000613ad08561524f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105255760a05191613e3a575b506122c757613b508461524f565b67ffffffffffffffff8116613b72816000526007602052604060002054151590565b156122455750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9390931691830191909152336024830152602090829060449082906001600160a01b03165afa9081156105255760a05191613e1b575b50156121cc57613bf08461524f565b90613c0660a48801926139e6613a5285856151fe565b15613dd4575050613cd661290660446020977ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09561ffff67ffffffffffffffff96161515600014613d3e5785613c5b8961524f565b1660a05152600d8a52613c77604060a051208a611fd9846152a0565b7f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61586613ca56129068b61524f565b604080516001600160a01b03929092168252602082018d90529190931692a25b0194613cd0866152a0565b5061524f565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081168252336020830152909216908201526060810185905292169180608081015b0390a280604051613d3581614e13565b52604051908152f35b5084613d498861524f565b168060a0515260088a527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8980613db46002604060a05120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916167de565b604080516001600160a01b039290921682526020820192909252a2613cc5565b613dde92506151fe565b6124be6040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161566e565b613e34915060203d60201161231b5761230d8183614e83565b87613be1565b613e53915060203d60201161231b5761230d8183614e83565b87613b42565b3461032b5760a05160031936011261032b5760206001600160a01b0360115416604051908152f35b3461032b57602060031936011261032b5760043567ffffffffffffffff811161032b578060040190610100600319823603011261032b57604051613ec481614e13565b60a0519052613ee9613edf613a59613a5260c48501866151fe565b6064830135615a49565b9060848101613ef7816152a0565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612c985750602481019277ffffffffffffffff00000000000000000000000000000000613f518561524f565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105255760a0519161431f575b506122c757613fd18461524f565b67ffffffffffffffff8116613ff3816000526007602052604060002054151590565b156122455750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9390931691830191909152336024830152602090829060449082906001600160a01b03165afa9081156105255760a05191614300575b50156121cc576140718461524f565b9061408760a48401926139e6613a5285856151fe565b15613dd457508291905067ffffffffffffffff6140a38561524f565b168060a0515260086020526140ec6002604060a05120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169485916167de565b604080516001600160a01b0385168152602081018690527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a267ffffffffffffffff6141398561524f565b1660a051526012602052604060a0512060ff600182015460a01c1660a05150806000146142f85781545b8086116142c45750156142af5761417b848254615264565b90556141868461524f565b505b6044017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166141be826152a0565b813b1561032b576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b0387811660048501526024840189905293909316604483015251909283916064918391905af1801561052557614296575b5067ffffffffffffffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc091613d258561426561290660209961524f565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b60a0516142a291614e83565b60a05161032b5784614227565b506142bc83601054615264565b601055614188565b85907fa17e11d50000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b601054614163565b614319915060203d60201161231b5761230d8183614e83565b85614062565b614338915060203d60201161231b5761230d8183614e83565b85613fc3565b3461032b5760a060031936011261032b57614357614cfe565b50614360614cbb565b60443567ffffffffffffffff811161032b5760031960a0913603011261032b57614388614d28565b506084359067ffffffffffffffff821161032b576143b367ffffffffffffffff923690600401614d59565b50506040516143c181614dc8565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260c060a0519101521660a05152600f60205260e0604060a051206040519061441582614dc8565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b3461032b57602060031936011261032b5767ffffffffffffffff6144ce614ca4565b1660a051526012602052602060ff6001604060a05120015460a01c166040519015158152f35b3461032b57604060031936011261032b5761450d614ca4565b60243567ffffffffffffffff82168060a05152601260205260ff6001604060a05120015460a01c161580156146ff575b611a9e578115610577576001600160a01b0361455884615464565b1633036105475760a051526012602052604060a0512060ff600182015460a01c166000146146ea5761458b828254615447565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f0000000000000000000000000000000000000000000000000000000000000000906145fd9061043b816084810161042d565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561032b576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a080516001600160a01b039390931660048301526024820185905251909283916044918391905af18015610525576146d1575b506040805167ffffffffffffffff9093168352602083019190915233917f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf9181908101611a2b565b60a0516146dd91614e83565b60a05161032b5782614689565b506146f781601054615447565b60105561458e565b50801561453d565b3461032b5760a05160031936011261032b57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461032b57602060031936011261032b576020614761614cfe565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b3461032b5760a05160031936011261032b5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461032b5760a05160031936011261032b57604051611b7c90614802606082614e83565b602481527f53696c6f65644c6f636b52656c65617365546f6b656e506f6f6c20312e372e3060208201527f2d646576000000000000000000000000000000000000000000000000000000006040820152604051918291602083526020830190614d87565b3461032b57602060031936011261032b5761487f614cfe565b614887615917565b61488f61535d565b908161489b5760a05180f35b60206001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599926149406040517fa9059cbb000000000000000000000000000000000000000000000000000000008582015261491a8161042d898660248401602090939291936001600160a01b0360408201951681520152565b7f00000000000000000000000000000000000000000000000000000000000000006166ab565b6040519485521692a28080610c93565b3461032b57602060031936011261032b576004358015610577576001600160a01b0361497d60a051615464565b1633036105475760a0805180526012602052805160409020600181015490911c60ff168015614adb5781545b808411611a62575015614ac6576149c1828254615264565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1561032b576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af1801561052557614ab4575b506040519060a0515060a051825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a260a05180f35b60a051614ac091614e83565b81614a75565b50614ad381601054615264565b6010556149c4565b6010546149a9565b3461032b5760c060031936011261032b57614afc614ca4565b614b04614cd2565b50614b0d614ce8565b5060843561ffff8116810361032b5760a4359067ffffffffffffffff821161032b5763ffffffff614b5461ffff92608095614b4d84963690600401614d59565b505061511b565b9392959091604051968752166020860152166040840152166060820152f35b3461032b57602060031936011261032b57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361032b57817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115614c7a575b8115614c50575b8115614c26575b8115614bfc575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483614bf5565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614bee565b7ff57e5f940000000000000000000000000000000000000000000000000000000081149150614be7565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150614be0565b6004359067ffffffffffffffff82168203610c8e57565b6024359067ffffffffffffffff82168203610c8e57565b602435906001600160a01b0382168203610c8e57565b606435906001600160a01b0382168203610c8e57565b600435906001600160a01b0382168203610c8e57565b35906001600160a01b0382168203610c8e57565b6064359061ffff82168203610c8e57565b6024359061ffff82168203610c8e57565b359061ffff82168203610c8e57565b9181601f84011215610c8e5782359167ffffffffffffffff8311610c8e5760208381860195010111610c8e57565b919082519283825260005b848110614db3575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201614d92565b60e0810190811067ffffffffffffffff821117614de457604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117614de457604052565b6060810190811067ffffffffffffffff821117614de457604052565b6040810190811067ffffffffffffffff821117614de457604052565b60a0810190811067ffffffffffffffff821117614de457604052565b90601f601f19910116810190811067ffffffffffffffff821117614de457604052565b67ffffffffffffffff8111614de457601f01601f191660200190565b929192614ece82614ea6565b91614edc6040519384614e83565b829481845281830111610c8e578281602093846000960137010152565b9080601f83011215610c8e57816020614f1493359101614ec2565b90565b9181601f84011215610c8e5782359167ffffffffffffffff8311610c8e576020808501948460051b010111610c8e57565b6040600319820112610c8e5760043567ffffffffffffffff8111610c8e5781614f7391600401614f17565b929092916024359067ffffffffffffffff8211610c8e57614f9691600401614f17565b9091565b9181601f84011215610c8e5782359167ffffffffffffffff8311610c8e5760208085019460e08502010111610c8e57565b906040600319830112610c8e5760043567ffffffffffffffff81168103610c8e57916024359067ffffffffffffffff8211610c8e57614f9691600401614d59565b602060408183019282815284518094520192019060005b8181106150305750505090565b82516001600160a01b0316845260209384019390920191600101615023565b9181601f84011215610c8e5782359167ffffffffffffffff8311610c8e5760208085019460608502010111610c8e57565b614f149160206150998351604084526040840190614d87565b920151906020818403910152614d87565b35908115158203610c8e57565b35906fffffffffffffffffffffffffffffffff82168203610c8e57565b9190826060910312610c8e576040516150ec81614e2f565b60406151168183956150fd816150aa565b855261510b602082016150b7565b6020860152016150b7565b910152565b67ffffffffffffffff16600052600f6020526040600020906040519161514083614dc8565b549263ffffffff84169384845263ffffffff8160201c169384602082015263ffffffff8260401c169384604083015263ffffffff8360601c169081606084015261ffff8460801c169283608082015260ff61ffff8660901c16958660a084015260a01c16159060c082159101526151e85761ffff161580159563ffffffff949392916151e05750945b156151d85750925b1693929190565b9050926151d1565b9050946151c9565b5050505092505050600090600090600090600090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610c8e570180359067ffffffffffffffff8211610c8e57602001918136038313610c8e57565b3567ffffffffffffffff81168103610c8e5790565b9190820391821161527157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b356001600160a01b0381168103610c8e5790565b9067ffffffffffffffff614f1492166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111614de45760051b60200190565b929190615315816152f1565b936153236040519586614e83565b602085838152019160051b8101928311610c8e57905b82821061534557505050565b6020809161535284614d14565b815201910190615339565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156153fc576000916153cd575090565b90506020813d6020116153f4575b816153e860209383614e83565b81010312610c8e575190565b3d91506153db565b6040513d6000823e3d90fd5b91908110156154185760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190820180921161527157565b91908110156154185760061b0190565b67ffffffffffffffff16600052601260205260016040600020015460ff8160a01c1661549a57506001600160a01b036011541690565b6001600160a01b031690565b67ffffffffffffffff16600052600e60205260406000209160028110156154ed576001146154dc57816001614f149301906160f7565b8160026003614f14940191016160f7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9190811015615418576060020190565b6040519061553982614e4b565b60606020838281520152565b80518210156154185760209160051b010190565b90600182811c921680156155a2575b602083101461557357565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691615568565b90604051918260008254926155c084615559565b808452936001811690811561562e57506001146155e7575b506155e592500383614e83565b565b90506000929192526020600020906000915b8183106156125750509060206155e592820101386155d8565b60209193508060019154838589010152019101909184926155f9565b602093506155e59592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386155d8565b601f8260209493601f19938186528686013760008582860101520116010190565b6040519061569c82614e67565b60006080838281528260208201528260408201528260608201520152565b906040516156c781614e67565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526008602052614f1460046040600020016155ac565b91908110156154185760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610c8e570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610c8e570180359067ffffffffffffffff8211610c8e57602001918160051b36038313610c8e57565b8181029291811591840414171561527157565b8181106157f4575050565b600081556001016157e9565b9160209082815201919060005b81811061581a5750505090565b9091926020806001926001600160a01b0361583488614d14565b16815201940192910161580d565b91908110156154185760081b0190565b3561ffff81168103610c8e5790565b3563ffffffff81168103610c8e5790565b358015158103610c8e5790565b359063ffffffff82168203610c8e57565b67ffffffffffffffff166158b1816000526007602052604060002054151590565b156158ea5780600052601260205260ff60016040600020015460a01c166158d9575060105490565b600052601260205260406000205490565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6001600160a01b0360015416330361592b57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b805180156159c557602003615987578051602082810191830183900312610c8e57519060ff8211615987575060ff1690565b6124be906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614d87565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161527157565b60ff16604d811161527157600a0a90565b8115615a1a570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414615b3157828411615b075790615a8e916159eb565b91604d60ff8416118015615aec575b615ab657505090615ab0614f14926159ff565b906157d6565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50615af6836159ff565b8015615a1a57600019048411615a9d565b615b10916159eb565b91604d60ff841611615ab657505090615b2b614f14926159ff565b90615a10565b5050505090565b90816020910312610c8e57518015158103610c8e5790565b356fffffffffffffffffffffffffffffffff81168103610c8e5790565b9160005b82811015615e8d5760e0810284016000615b8a8261524f565b9067ffffffffffffffff821691615bae836000526007602052604060002054151590565b15615e6157615c779260408593615c22615c1c94615c1c615be2602060019c9b0192611fb5615bdd36866150d4565b616564565b91825463ffffffff8160801c16159081615e43575b81615e34575b81615e19575b81615e0a575b5080615dfb575b615d70575b36906150d4565b906169ed565b6080850192615c34615bdd36866150d4565b8152600d6020522092835463ffffffff8160801c16159081615d52575b81615d43575b81615d28575b81615d19575b5080615d0a575b615c7d575b5036906150d4565b01615b71565b615c9a60a06fffffffffffffffffffffffffffffffff9201615b50565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538615c6f565b50615d1482615872565b615c6a565b60ff915060a01c161538615c63565b6fffffffffffffffffffffffffffffffff8116159150615c5d565b8589015460801c159150615c57565b858901546fffffffffffffffffffffffffffffffff16159150615c51565b6fffffffffffffffffffffffffffffffff615d8c878b01615b50565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355615c15565b50615e0581615872565b615c10565b60ff915060a01c161538615c09565b6fffffffffffffffffffffffffffffffff8116159150615c03565b848e015460801c159150615bfd565b848e01546fffffffffffffffffffffffffffffffff16159150615bf7565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b908051156160795767ffffffffffffffff81516020830120921691826000526008602052615ec8816005604060002001616fac565b156160355760005260096020526040600020815167ffffffffffffffff8111614de457615ef58254615559565b601f8111616003575b506020601f8211600114615f795791615f53827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593615f6995600091615f6e575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190614d87565b0390a2565b905084015138615f40565b601f1982169083600052806000209160005b818110615feb575092615f699492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615fd2575b5050811b019055611d28565b85015160001960f88460031b161c191690553880615fc6565b9192602060018192868a015181550194019201615f8b565b61602f90836000526020600020601f840160051c81019160208510610bbd57601f0160051c01906157e9565b38615efe565b50906124be6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614d87565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b8181106160d55750506155e592500383614e83565b84546001600160a01b03168352600194850194879450602090930192016160c0565b616100906160a3565b9160055480151591826161e0575b5050616118575090565b616121906160a3565b9081518061612f5750905090565b61613a908251615447565b92601f1961616061614a866152f1565b956161586040519788614e83565b8087526152f1565b0136602086013760005b825181101561619b57806001600160a01b0361618860019386615545565b51166161948288615545565b520161616a565b509160005b81518110156161db57806001600160a01b036161be60019385615545565b51166161d46161ce838751615447565b88615545565b52016161a0565b505050565b10159050388061610e565b67ffffffffffffffff1660008181526007602052604090205490929190156162ed57916162ea60e0926162b6856162427f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97616564565b8460005260086020526162598160406000206169ed565b61626283616564565b84600052600860205261627c8360026040600020016169ed565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61632361568f565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691616380602085019361637a61636d63ffffffff87511642615264565b85608089015116906157d6565b90615447565b8082101561639957505b16825263ffffffff4216905290565b905061638a565b67ffffffffffffffff6163b56020830161524f565b16600052600f6020526040600020604051906163d082614dc8565b549063ffffffff8216815263ffffffff8260201c16602082015263ffffffff8260401c16604082015263ffffffff8260601c16606082015261ffff8260801c169081608082015260ff61ffff8460901c16938460a084015260a01c16159060c0821591015261647b57919261ffff929083161561647457505b16801561646c576127106164656060614f1494013592836157d6565b0490615264565b506060013590565b9050616449565b505060609150013590565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152614f14604082614e83565b805160005b8181106164d257505050565b60018101808211615271575b8281106164ee57506001016164c6565b6001600160a01b036165008386615545565b51166001600160a01b036165148387615545565b511614616523576001016164de565b6001600160a01b036165358386615545565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115616604576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106165a15750565b606490616602604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040820151161580159061668c575b61662b5750565b606490616602604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515616624565b6001600160a01b0361672d9116916040926000808551936166cc8786614e83565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d156167d6573d9161671183614ea6565b9261671e87519485614e83565b83523d6000602085013e617001565b8051908161673a57505050565b60208061674b938301019101615b38565b156167535750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091617001565b9182549060ff8260a01c161580156169e5575b6169df576fffffffffffffffffffffffffffffffff8216916001850190815461683663ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615264565b9081616941575b5050848110616902575083831061689757505061686c6fffffffffffffffffffffffffffffffff928392615264565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916168a68185615264565b92600019810190808211615271576168c96168ce926001600160a01b0396615447565b615a10565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b8286929396116169b55761695c9261637a9160801c906157d6565b808410156169b05750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617865592388061683d565b616967565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156167f1565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991616b266060928054616a2a63ffffffff8260801c1642615264565b9081616b65575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014616b5d57508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155616ada8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6162ea60405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091616a61565b6fffffffffffffffffffffffffffffffff91616b9a839283616b936001880154948286169560801c906157d6565b9116615447565b80821015616c1957505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880616a31565b9050616ba4565b906040519182815491828252602082019060005260206000209260005b818110616c525750506155e592500383614e83565b8454835260019485019487945060209093019201616c3d565b80548210156154185760005260206000200190600090565b6000818152600360205260409020548015616d7c5760001981018181116152715760025490600019820191821161527157818103616d2b575b5050506002548015616cfc5760001901616cd7816002616c6b565b60001982549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b616d64616d3c616d4d936002616c6b565b90549060031b1c9283926002616c6b565b81939154906000199060031b92831b921b19161790565b90556000526003602052604060002055388080616cbc565b5050600090565b6000818152600760205260409020548015616d7c5760001981018181116152715760065490600019820191821161527157818103616dfc575b5050506006548015616cfc5760001901616dd7816006616c6b565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b616e1e616e0d616d4d936006616c6b565b90549060031b1c9283926006616c6b565b90556000526007602052604060002055388080616dbc565b9060018201918160005282602052604060002054801515600014616ee957600019810181811161527157825490600019820191821161527157818103616eb2575b50505080548015616cfc576000190190616e918282616c6b565b60001982549160031b1b191690555560005260205260006040812055600190565b616ed2616ec2616d4d9386616c6b565b90549060031b1c92839286616c6b565b905560005283602052604060002055388080616e77565b50505050600090565b80600052600360205260406000205415600014616f4c5760025468010000000000000000811015614de457616f33616d4d8260018594016002556002616c6b565b9055600254906000526003602052604060002055600190565b50600090565b80600052600760205260406000205415600014616f4c5760065468010000000000000000811015614de457616f93616d4d8260018594016006556006616c6b565b9055600654906000526007602052604060002055600190565b6000828152600182016020526040902054616d7c5780549068010000000000000000821015614de45782616fea616d4d846001809601855584616c6b565b905580549260005201602052604060002055600190565b9192901561707c5750815115617015575090565b3b1561701e5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561708f5750805190602001fd5b6124be906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190614d8756fea164736f6c634300081a000a",
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAccumulatedFees(&_SiloedLockReleaseTokenPool.CallOpts)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetAccumulatedFees(&_SiloedLockReleaseTokenPool.CallOpts)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentInboundRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentInboundRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedLockReleaseTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_SiloedLockReleaseTokenPool.CallOpts, remoteChainSelector)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _SiloedLockReleaseTokenPool.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) GetFee(destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetFee(&_SiloedLockReleaseTokenPool.CallOpts, destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolCallerSession) GetFee(destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedLockReleaseTokenPool.Contract.GetFee(&_SiloedLockReleaseTokenPool.CallOpts, destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn0(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.LockOrBurn0(&_SiloedLockReleaseTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint0(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.ReleaseOrMint0(&_SiloedLockReleaseTokenPool.TransactOpts, releaseOrMintIn, finality)
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawFees(&_SiloedLockReleaseTokenPool.TransactOpts, recipient)
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _SiloedLockReleaseTokenPool.Contract.WithdrawFees(&_SiloedLockReleaseTokenPool.TransactOpts, recipient)
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

type SiloedLockReleaseTokenPoolPoolFeeWithdrawnIterator struct {
	Event *SiloedLockReleaseTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedLockReleaseTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedLockReleaseTokenPoolPoolFeeWithdrawn)
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
		it.Event = new(SiloedLockReleaseTokenPoolPoolFeeWithdrawn)
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

func (it *SiloedLockReleaseTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *SiloedLockReleaseTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedLockReleaseTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*SiloedLockReleaseTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &SiloedLockReleaseTokenPoolPoolFeeWithdrawnIterator{contract: _SiloedLockReleaseTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SiloedLockReleaseTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedLockReleaseTokenPoolPoolFeeWithdrawn)
				if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*SiloedLockReleaseTokenPoolPoolFeeWithdrawn, error) {
	event := new(SiloedLockReleaseTokenPoolPoolFeeWithdrawn)
	if err := _SiloedLockReleaseTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (SiloedLockReleaseTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
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
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (SiloedLockReleaseTokenPoolUnsiloedRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b234")
}

func (_SiloedLockReleaseTokenPool *SiloedLockReleaseTokenPool) Address() common.Address {
	return _SiloedLockReleaseTokenPool.address
}

type SiloedLockReleaseTokenPoolInterface interface {
	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

		error)

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

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

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

	WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

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

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*SiloedLockReleaseTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedLockReleaseTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*SiloedLockReleaseTokenPoolPoolFeeWithdrawn, error)

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
