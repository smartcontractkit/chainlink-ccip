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

var BurnMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentCustomBlockConfirmationRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"blockConfirmationConfigured\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationRateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultFinalityRateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103b757616453803803809161001d8285610436565b833981019060a0818303126103b75780516001600160a01b038116908190036103b75761004c60208301610459565b60408301519091906001600160401b0381116103b75783019380601f860112156103b7578451946001600160401b038611610420578560051b9060208201966100986040519889610436565b87526020808801928201019283116103b757602001905b828210610408575050506100d160806100ca60608601610467565b9401610467565b9233156103f757600180546001600160a01b03191633179055811580156103e6575b80156103d5575b6103c4578160209160049360805260c0526040519283809263313ce56760e01b82525afa60009181610383575b50610358575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261023a575b604051615e37908161061c8239608051818181611c2c01528181611e0701528181611f2101528181612001015281816125240152818161270f015281816130c20152818161329c015281816132f601528181613463015281816135b70152818161376901528181613b3301528181613b880152615256015260a051818181613aef015281816148020152818161486c01526152e3015260c051818181610c7f01528181611cbb015281816125b2015281816131510152613647015260e051818181610b2301528181611d00015281816125f70152612eb30152f35b60405160206102498183610436565b60008252600036813760e051156103475760005b82518110156102c4576001906001600160a01b0361027b828661047b565b511683610287826104bd565b610294575b50500161025d565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388361028c565b50905060005b825181101561033e576001906001600160a01b036102e8828661047b565b5116801561033857836102fa826105bb565b610308575b50505b016102ca565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138836102ff565b50610302565b5050503861015f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff821681810361036c575061012d565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103bc575b8161039f60209383610436565b810103126103b7576103b090610459565b9038610127565b600080fd5b3d9150610392565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100fa565b506001600160a01b038416156100f3565b639b15e16f60e01b60005260046000fd5b6020809161041584610467565b8152019101906100af565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761042057604052565b519060ff821682036103b757565b51906001600160a01b03821682036103b757565b805182101561048f5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561048f5760005260206000200190600090565b60008181526003602052604090205480156105b457600019810181811161059e5760025460001981019190821161059e5781810361054d575b505050600254801561053757600019016105118160026104a5565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61058661055e61056f9360026104a5565b90549060031b1c92839260026104a5565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806104f6565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526003602052604060002054156000146106155760025468010000000000000000811015610420576105fc61056f82600185940160025560026104a5565b9055600254906000526003602052604060002055600190565b5060009056fe60c080604052600436101561001357600080fd5b600060a05260a0513560e01c90816301ffc9a714613c0e57508063181f5a7714613bac57806321df0da714613b67578063240028e814613b1357806324f65ee714613ad45780632c06340414613a3c57806337b19247146138ce5780633907753714613530578063489a68f21461302a5780634c5ef0ed14612fe557806354c8a4f314612e8157806359152aad14612dd457806362ddd3c414612d4f578063698c2c6614612ca85780636d3d1a5814612c805780637437ff9f14612c4c57806379ba509714612b9a5780637d54534e14612b215780638926f54f14612adc57806389720a6214612a715780638da5cb5b14612a49578063962d4020146129065780639751f884146128a15780639a4575b9146124d2578063a42a7b8b1461237b578063a7cd63b71461230d578063acfecf91146121e0578063b1c71c6514611bb0578063b794658014611b74578063bb6bb5a714611b0c578063c4bffe2b146119f0578063c7230a60146116d4578063cf7401f31461156b578063d966866b1461111e578063dc04fa1f14610ca3578063dc0bd97114610c5e578063ded8d95614610b48578063e0351e1314610b0a578063e8a1da17146102c5578063f2fde38b146102115763fa41d79c146101e857600080fd5b3461020b5760a05160031936011261020b57602061ffff600b5416604051908152f35b60a05180fd5b3461020b57602060031936011261020b576001600160a01b03610232613e30565b61023a614970565b163381146102995760a051805473ffffffffffffffffffffffffffffffffffffffff1916821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360a05180f35b7fdad89dca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461020b576102d336613f83565b9190926102de614970565b60a051905b8282106109485750505060a0519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b8181101561094257600581901b8301358581121561020b5783016101208136031261020b576040519461035286613d78565b813567ffffffffffffffff8116810361093d578652602082013567ffffffffffffffff811161020b5782019436601f8701121561020b57853561039481614340565b966103a26040519889613db0565b81885260208089019260051b8201019036821161020b5760208101925b82841061090e575050505060208701958652604083013567ffffffffffffffff811161020b576103f29036908501613f34565b916040880192835261041c61040a3660608701614106565b9460608a0195865260c0369101614106565b956080890196875261042e855161548d565b610438875161548d565b835151156108e25761045467ffffffffffffffff8a5116615abb565b156108a75767ffffffffffffffff89511660a051526008602052604060a0512061057486516001600160801b03604082015116906105386001600160801b03602083015116915115158360806040516104ac81613d78565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160801b0384161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176001830155565b61067688516001600160801b036040820151169061063a6001600160801b03602083015116915115158360806040516105ac81613d78565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166001600160801b0385161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176003830155565b6004855191019080519067ffffffffffffffff821161088f5761069983546144eb565b601f8111610850575b506020906001601f8411146107e6579180916106d59360a051926107db575b50506000198260011b9260031b1c19161790565b90555b60a0515b88518051821015610711579061070b6001926107048367ffffffffffffffff8f5116926144d7565b5190614d49565b016106dc565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29391999750956107cd67ffffffffffffffff60019796949851169251935191516107a261077660405196879687526101006020880152610100870190613def565b9360408601906001600160801b0360408092805115158552826020820151166020860152015116910152565b60a08401906001600160801b0360408092805115158552826020820151166020860152015116910152565b0390a1019392909193610320565b015190508e806106c1565b90601f198316918460a051528160a051209260a0515b818110610838575090846001959493921061081f575b505050811b0190556106d8565b015160001960f88460031b161c191690558d8080610812565b929360206001819287860151815501950193016107fc565b61087f908460a05152602060a05120601f850160051c81019160208610610885575b601f0160051c01906146e7565b8d6106a2565b9091508190610872565b634e487b7160e01b60a051526041600452602460a051fd5b67ffffffffffffffff8951167f1d5ad3c50000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b833567ffffffffffffffff811161020b576020916109328392833691870101613f34565b8152019301926103bf565b600080fd5b60a05180f35b9092919367ffffffffffffffff610968610963868886614409565b6142ee565b1692610973846158ec565b15610ada578360a0515260086020526109936005604060a05120016157a2565b9260a0515b84518110156109d2576001908660a0515260086020526109cb6005604060a05120016109c483896144d7565b519061599f565b5001610998565b50939094919592508060a0515260086020526005604060a0512060a051815560a051600182015560a051600282015560a051600382015560048101610a1781546144eb565b80610a8a575b50500180549060a051815581610a66575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102e3565b60a05152602060a05120908101905b81811015610a2e5760a0518155600101610a75565b601f8111600114610aa3575060a05190555b8880610a1d565b610ac3908260a051526001601f602060a051209201861c820191016146e7565b60a080518290525160208120918190559055610a9c565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461020b5760a05160031936011261020b5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461020b57602060031936011261020b57610140610b64613e87565b610b6c61443f565b50610b7561443f565b50610c5c610bcb610bab610ba6610bb0610bab610ba68767ffffffffffffffff16600052600c602052604060002090565b61446a565b6151d9565b9467ffffffffffffffff16600052600d602052604060002090565b610c1560405180946001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b3461020b5760a05160031936011261020b5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020b57604060031936011261020b5760043567ffffffffffffffff811161020b573660238201121561020b57806004013567ffffffffffffffff811161020b576024820191602436918360081b01011161020b5760243567ffffffffffffffff811161020b57610d19903690600401613f52565b919092610d24614970565b60a0515b828110610da05750505060a0515b818110610d435760a05180f35b8067ffffffffffffffff610d5d6109636001948688614409565b168060a05152600f60205260a051604060a05120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee860a05160a051a201610d36565b610dae610963828585614740565b610db9828585614740565b90602082019060e0830190610dcd82614750565b156110835760a0840161271061ffff610de58361475d565b1610156111125760c085019161271061ffff610e008561475d565b1610156110d65763ffffffff610e158661476c565b161580156110bd575b6110835767ffffffffffffffff16948560a05152600f60205260a05160409020610e478661476c565b63ffffffff16908054906040840191610e5f8361476c565b60201b67ffffffff0000000016936060860194610e7b8661476c565b60401b6bffffffff0000000000000000169660800196610e9a8861476c565b60601b6fffffffff0000000000000000000000001691610eb98a61475d565b60801b71ffff000000000000000000000000000000001693610eda8c61475d565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610f8d87614750565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610fde9061477d565b63ffffffff168752610fef9061477d565b63ffffffff1660208701526110039061477d565b63ffffffff1660408601526110179061477d565b63ffffffff16606085015261102b90613ec0565b61ffff16608084015261103d90613ec0565b61ffff1660a083015261104f906140e5565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610d28565b67ffffffffffffffff907f123322650000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b5063ffffffff6110cf6040880161476c565b1615610e1e565b61ffff6110e28461475d565b7f95f3517a0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b6110e261ffff9161475d565b3461020b57602060031936011261020b5760043567ffffffffffffffff811161020b5761114f903690600401613f52565b90611158614970565b60a0516080525b816080511061116e5760a05180f35b61117e610963608051848461462a565b61119861118e608051858561462a565b602081019061466a565b6111b26111a8608051878761462a565b604081019061466a565b906111cd6111c3608051898961462a565b606081019061466a565b6111e76111dd6080518b8b61462a565b608081019061466a565b9390946111fd6111f836898b614358565b6153ea565b61120b6111f8368385614358565b6112196111f8368587614358565b6112276111f8368789614358565b6040519860808a01908a821067ffffffffffffffff83111761088f5767ffffffffffffffff9160405261125b368a8c614358565b8b52611268368486614358565b60208c0152611278368688614358565b60408c015261128836888a614358565b60608c015216988960a05152600e602052604060a05120815180519067ffffffffffffffff821161088f5768010000000000000000821161088f57602090835483855580841061154c575b50018260a05152602060a0512060a0515b83811061152f57505050506001810160208301519081519167ffffffffffffffff831161088f5768010000000000000000831161088f576020908254848455808510611510575b50019060a05152602060a0512060a0515b8381106114f357505050506002810160408301519081519167ffffffffffffffff831161088f5768010000000000000000831161088f5760209082548484558085106114d4575b50019060a05152602060a0512060a0515b8381106114b7575050505060036060910191015180519067ffffffffffffffff821161088f5768010000000000000000821161088f576020908354838555808410611498575b50019160a05152602060a051209160a0515b82811061147b57505050509261146a949261144e61145c937fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9998966114406040519b8c9b60808d5260808d01916146fe565b918a830360208c01526146fe565b9187830360408901526146fe565b9184830360608601526146fe565b0390a260016080510160805261115f565b60019060206001600160a01b0384511693019281860155016113ec565b6114b1908560a05152848460a0512091820191016146e7565b8f6113da565b60019060206001600160a01b038551169401938184015501611394565b6114ed908460a05152858460a0512091820191016146e7565b38611383565b60019060206001600160a01b03855116940193818401550161133c565b611529908460a05152858460a0512091820191016146e7565b3861132b565b60019060206001600160a01b0385511694019381840155016112e4565b611565908560a05152848460a0512091820191016146e7565b386112d3565b3461020b5760e060031936011261020b57611584613e87565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc011261020b576040516115ba81613d94565b602435801515810361020b5781526044356001600160801b038116810361020b5760208201526064356001600160801b038116810361020b5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c011261020b576040519061162f82613d94565b608435801515810361020b57825260a4356001600160801b038116810361020b57602083015260c4356001600160801b038116810361020b5760408301526001600160a01b03600a5416331415806116bf575b61168f57610942926150ae565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506001600160a01b0360015416331415611682565b3461020b57604060031936011261020b5760043567ffffffffffffffff811161020b57611705903690600401613f52565b6024356001600160a01b03811680820361020b57611721614970565b60a0515b8381106117325760a05180f35b8060206001600160a01b0361175261174d602495898b614409565b6142da565b16604051938480927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa80156119e357838591879460a051916119a4575b50806117ad575b5050506001915001611725565b886118956001600160a01b036117c761174d888a86614409565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082019081526001600160a01b0398909816602482015260448082018790528152911661181a606483613db0565b604080519097909261182c8985613db0565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260a0519160a05191519060a051855af13d1561199c573d9161187783613dd3565b926118848a519485613db0565b835260a0513d90602085013e615d5e565b8051806118ed575b50507f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e916001600160a01b036118db61174d8860019a602096614409565b169451908152a38491508383886117a0565b816020935083929496979850611907955001019101614958565b15611919579083869392888a8061189d565b608482517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615d5e565b93945050505060203d81116119dc575b6119be8183613db0565b602082600092810103126119d9575090838386935189611799565b80fd5b503d6119b4565b6040513d60a051823e3d90fd5b3461020b5760a05160031936011261020b5760a051506040516006548082528160208101600660a05152602060a051209260a0515b818110611af3575050611a3a92500382613db0565b805190611a5f611a4983614340565b92611a576040519485613db0565b808452614340565b90601f1960208401920136833760a0515b8151811015611aa2578067ffffffffffffffff611a8f600193856144d7565b5116611a9b82876144d7565b5201611a70565b5050906040519182916020830190602084525180915260408301919060a0515b818110611ad0575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611ac2565b8454835260019485019486945060209093019201611a25565b3461020b57602060031936011261020b5760043567ffffffffffffffff811161020b57611b3d903690600401613fd5565b6001600160a01b03600a541633141580611b5f575b61168f57610942916149fd565b506001600160a01b0360015416331415611b52565b3461020b57602060031936011261020b57611bac611b98611b93613e87565b614608565b604051918291602083526020830190613def565b0390f35b3461020b57606060031936011261020b5760043567ffffffffffffffff811161020b5760a0600319823603011261020b57611be9613eaf565b9060443567ffffffffffffffff811161020b57611c0a903690600401613f34565b50611c136144be565b5060848101611c21816142da565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361219f5750602481019077ffffffffffffffff00000000000000000000000000000000611c7b836142ee565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119e35760a05191612170575b5061214457611cfe604482016142da565b7f00000000000000000000000000000000000000000000000000000000000000006120f1575b5067ffffffffffffffff611d37836142ee565b16611d4f816000526007602052604060002054151590565b156120c25760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156119e35760a05190612079575b6001600160a01b03915016330361204957606481013561ffff84168015611f9e5761ffff600b54169081611eb6575b505050611b93611dee611eac94611e7b935b600401615317565b92611df88461524c565b611e01816142ee565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a26142ee565b90611e846152dc565b60405192611e9184613d5c565b835260208301526040519283926040845260408401906140bb565b9060208301520390f35b818110611f6c575050611dee611eac94611e7b93611b93937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611f01896142ee565b16918260a05152600c60205280611f49604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615b6a565b604080516001600160a01b039290921682526020820192909252a2935094611dd4565b7f7911d95b0000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b50611dee611eac94611e7b93611b93937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611fe1896142ee565b16918260a05152600860205280612029604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615b6a565b604080516001600160a01b039290921682526020820192909252a2611de6565b7f728fe07b0000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506020813d6020116120ba575b8161209360209383613db0565b8101031261020b57516001600160a01b038116810361020b576001600160a01b0390611da5565b3d9150612086565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b6001600160a01b0316612111816000526003602052604060002054151590565b611d24577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b612192915060203d602011612198575b61218a8183613db0565b810190614958565b84611ced565b503d612180565b6121b06001600160a01b03916142da565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b3461020b5767ffffffffffffffff6121f736614006565b929091612202614970565b169061221b826000526007602052604060002054151590565b156122dd578160a05152600860205261224e6005604060a0512001612241368685613efd565b602081519101209061599f565b15612296577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76919261228d6040519283926020845260208401916145e7565b0390a260a05180f35b6122d9906040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916145e7565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461020b5760a05160031936011261020b5760a051506040516002548082526020820190600260a05152602060a051209060a0515b81811061236557611bac8561235981870382613db0565b60405191829182614047565b8254845260209093019260019283019201612342565b3461020b57602060031936011261020b5767ffffffffffffffff61239d613e87565b1660a0515260086020526123b86005604060a05120016157a2565b805190601f196123e06123ca84614340565b936123d86040519586613db0565b808552614340565b0160a0515b8181106124c157505060a0515b815181101561243c5780612408600192846144d7565b5160a051526009602052612420604060a05120614525565b61242a82866144d7565b5261243581856144d7565b50016123f2565b826040518091602082016020835281518091526040830190602060408260051b86010193019160a051905b82821061247657505050500390f35b919360206124b1827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613def565b9601920192018594939192612467565b8060606020809387010152016123e5565b3461020b57602060031936011261020b5760043567ffffffffffffffff811161020b5760a0600319823603011261020b5761250b6144be565b5060848101612519816142da565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361219f57506024810177ffffffffffffffff00000000000000000000000000000000612572826142ee565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119e35760a05191612882575b50612144576125f5604483016142da565b7f000000000000000000000000000000000000000000000000000000000000000061282f575b5067ffffffffffffffff61262e826142ee565b16612646816000526007602052604060002054151590565b156120c25760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156119e35760a051906127e6575b6001600160a01b03915016330361204957611bac916127b691611b9391606401357ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108167ffffffffffffffff6126f1856142ee565b168060a051526008602052612737604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958691615b6a565b604080516001600160a01b0386168152602081018490527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a261277b8161524c565b67ffffffffffffffff61278d856142ee565b604080516001600160a01b03909616865233602087015285019290925216918060608101611e73565b6127be6152dc565b604051916127cb83613d5c565b825260208201526040519182916020835260208301906140bb565b506020813d602011612827575b8161280060209383613db0565b8101031261020b57516001600160a01b038116810361020b576001600160a01b039061269c565b3d91506127f3565b6001600160a01b031661284f816000526003602052604060002054151590565b61261b577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b61289b915060203d6020116121985761218a8183613db0565b836125e4565b3461020b57602060031936011261020b5767ffffffffffffffff6128c3613e87565b6128cb61443f565b506128d461443f565b501660a051526008602052610140604060a05120610c5c610bcb610bab60026128ff610bab8661446a565b940161446a565b3461020b57606060031936011261020b5760043567ffffffffffffffff811161020b57612937903690600401613f52565b9060243567ffffffffffffffff811161020b5761295890369060040161408a565b9060443567ffffffffffffffff811161020b5761297990369060040161408a565b6001600160a01b03600a541633141580612a34575b61168f57838614801590612a2a575b6129fe5760a0515b8681106129b25760a05180f35b806129f86129c66109636001948b8b614409565b6129d183898961442f565b6129f26129ea6129e286898b61442f565b923690614106565b913690614106565b916150ae565b016129a5565b7f568efce20000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b508086141561299d565b506001600160a01b036001541633141561298e565b3461020b5760a05160031936011261020b5760206001600160a01b0360015416604051908152f35b3461020b5760c060031936011261020b57612a8a613e30565b50612a93613e70565b612a9b613e9e565b5060843567ffffffffffffffff811161020b57612abc903690600401613ecf565b505060a43590600282101561020b57611bac9161235991604435906143ac565b3461020b57602060031936011261020b576020612b1767ffffffffffffffff612b03613e87565b166000526007602052604060002054151590565b6040519015158152f35b3461020b57602060031936011261020b577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b03612b65613e30565b612b6d614970565b168073ffffffffffffffffffffffffffffffffffffffff19600a541617600a55604051908152a160a05180f35b3461020b5760a05160031936011261020b5760a051546001600160a01b0381163303612c205773ffffffffffffffffffffffffffffffffffffffff196001549133828416176001551660a051556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060a05160a051a360a05180f35b7f02b543c60000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461020b5760a05160031936011261020b57600454600554604080516001600160a01b039093168352602083019190915290f35b3461020b5760a05160031936011261020b5760206001600160a01b03600a5416604051908152f35b3461020b57604060031936011261020b57612cc1613e30565b602435612ccc614970565b6001600160a01b0382169182156108e2577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f02389273ffffffffffffffffffffffffffffffffffffffff19600454161760045581600555612d4660405192839283602090939291936001600160a01b0360408201951681520152565b0390a160a05180f35b3461020b57612d5d36614006565b612d68929192614970565b67ffffffffffffffff8216612d8a816000526007602052604060002054151590565b15612da5575061094292612d9f913691613efd565b90614d49565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461020b57604060031936011261020b5760043561ffff81169081900361020b5760243567ffffffffffffffff811161020b577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e91612e74612e3c6020933690600401613fd5565b90612e45614970565b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b556149fd565b604051908152a160a05180f35b3461020b57612ea9612eb1612e9536613f83565b9491612ea2939193614970565b3691614358565b923691614358565b7f000000000000000000000000000000000000000000000000000000000000000015612fb95760a0515b8251811015612f4157806001600160a01b03612ef9600193866144d7565b5116612f0481615805565b612f10575b5001612edb565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612f09565b5060a0515b815181101561094257806001600160a01b03612f64600193856144d7565b51168015612fb357612f7581615a5b565b612f82575b505b01612f46565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612f7a565b50612f7c565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461020b57604060031936011261020b57612ffe613e87565b60243567ffffffffffffffff811161020b57602091613024612b17923690600401613f34565b90614303565b3461020b57604060031936011261020b5760043567ffffffffffffffff811161020b578060040190610100600319823603011261020b57613069613eaf565b60405190919061307881613d40565b60a05190526130a961309f61309a61309360c4850187614289565b3691613efd565b61478e565b6064830135614869565b91608482016130b7816142da565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361219f5750602482019377ffffffffffffffff00000000000000000000000000000000613111866142ee565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119e35760a05191613511575b506121445767ffffffffffffffff61319a866142ee565b166131b2816000526007602052604060002054151590565b156120c25760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156119e35760a051916134f2575b50156120495761321e856142ee565b9061323460a48501926130246130938585614289565b156134ab57506044929161ffff1615905061340d5767ffffffffffffffff61325b856142ee565b168060a05152600d6020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61584806132c4604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615b6a565b604080516001600160a01b039290921682526020820192909252a25b01916132eb836142da565b906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561020b576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390921660048201526024810185905290818060448101038160a051875af180156119e3576133f4575b50608067ffffffffffffffff6020956001600160a01b036133c26133bc7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966142ee565b926142da565b60405196875233898801521660408601528560608601521692a2604051906133e982613d40565b815260405190518152f35b60a05161340091613db0565b60a05161020b5784613378565b67ffffffffffffffff61341f856142ee565b168060a0515260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c848061348b6002604060a05120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615b6a565b604080516001600160a01b039290921682526020820192909252a26132e0565b6134b59250614289565b6122d96040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916145e7565b61350b915060203d6020116121985761218a8183613db0565b8661320f565b61352a915060203d6020116121985761218a8183613db0565b86613183565b3461020b57602060031936011261020b5760043567ffffffffffffffff811161020b578060040190610100600319823603011261020b5760405161357381613d40565b60a051905260405161358481613d40565b60a051905261359f61309f61309a61309360c4850186614289565b90608481016135ad816142da565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691160361219f5750602481019277ffffffffffffffff00000000000000000000000000000000613607856142ee565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119e35760a051916138af575b506121445767ffffffffffffffff613690856142ee565b166136a8816000526007602052604060002054151590565b156120c25760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156119e35760a05191613890575b501561204957613714846142ee565b9061372a60a48401926130246130938585614289565b156134ab575082916044915067ffffffffffffffff613748866142ee565b168060a0515260086020526137916002604060a05120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958691615b6a565b604080516001600160a01b0386168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a201926137d7846142da565b823b1561020b576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390921660048201526024810185905290818060448101038160a051875af180156119e3576020956001600160a01b036133c26133bc7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09660809667ffffffffffffffff9661387e575b506142ee565b60a05161388a91613db0565b8b613878565b6138a9915060203d6020116121985761218a8183613db0565b85613705565b6138c8915060203d6020116121985761218a8183613db0565b85613679565b3461020b5760a060031936011261020b576138e7613e30565b506138f0613e70565b60443567ffffffffffffffff811161020b5760031960a0913603011261020b57613918613e9e565b506084359067ffffffffffffffff821161020b5761394367ffffffffffffffff923690600401613ecf565b505060405161395181613d0e565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260c060a0519101521660a05152600f60205260e0604060a05120604051906139a582613d0e565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b3461020b5760c060031936011261020b57613a55613e30565b50613a5e613e70565b613a66613e46565b5060843561ffff8116810361020b5760a4359067ffffffffffffffff821161020b5763ffffffff61ffff613aad829386613aa660a0973690600401613ecf565b505061414d565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b3461020b5760a05160031936011261020b57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020b57602060031936011261020b576020613b2e613e30565b6040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081169216919091148152f35b3461020b5760a05160031936011261020b5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020b5760a05160031936011261020b5760408051611bac91613bd09082613db0565b601b81527f4275726e4d696e74546f6b656e506f6f6c20312e372e302d64657600000000006020820152604051918291602083526020830190613def565b3461020b57602060031936011261020b57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361020b57817faff2afbf0000000000000000000000000000000000000000000000000000000060209314908115613ce4575b8115613cba575b8115613c90575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613c89565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613c82565b7fdc0cbd360000000000000000000000000000000000000000000000000000000081149150613c7b565b60e0810190811067ffffffffffffffff821117613d2a57604052565b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613d2a57604052565b6040810190811067ffffffffffffffff821117613d2a57604052565b60a0810190811067ffffffffffffffff821117613d2a57604052565b6060810190811067ffffffffffffffff821117613d2a57604052565b90601f601f19910116810190811067ffffffffffffffff821117613d2a57604052565b67ffffffffffffffff8111613d2a57601f01601f191660200190565b919082519283825260005b848110613e1b575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201613dfa565b600435906001600160a01b038216820361093d57565b606435906001600160a01b038216820361093d57565b35906001600160a01b038216820361093d57565b6024359067ffffffffffffffff8216820361093d57565b6004359067ffffffffffffffff8216820361093d57565b6064359061ffff8216820361093d57565b6024359061ffff8216820361093d57565b359061ffff8216820361093d57565b9181601f8401121561093d5782359167ffffffffffffffff831161093d576020838186019501011161093d57565b929192613f0982613dd3565b91613f176040519384613db0565b82948184528183011161093d578281602093846000960137010152565b9080601f8301121561093d57816020613f4f93359101613efd565b90565b9181601f8401121561093d5782359167ffffffffffffffff831161093d576020808501948460051b01011161093d57565b604060031982011261093d5760043567ffffffffffffffff811161093d5781613fae91600401613f52565b929092916024359067ffffffffffffffff821161093d57613fd191600401613f52565b9091565b9181601f8401121561093d5782359167ffffffffffffffff831161093d5760208085019460e0850201011161093d57565b90604060031983011261093d5760043567ffffffffffffffff8116810361093d57916024359067ffffffffffffffff821161093d57613fd191600401613ecf565b602060408183019282815284518094520192019060005b81811061406b5750505090565b82516001600160a01b031684526020938401939092019160010161405e565b9181601f8401121561093d5782359167ffffffffffffffff831161093d576020808501946060850201011161093d57565b613f4f9160206140d48351604084526040840190613def565b920151906020818403910152613def565b3590811515820361093d57565b35906001600160801b038216820361093d57565b919082606091031261093d5760405161411e81613d94565b604061414881839561412f816140e5565b855261413d602082016140f2565b6020860152016140f2565b910152565b67ffffffffffffffff16600052600f60205260406000206040519061417182613d0e565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c168152608085019561ffff8360801c16875260ff60a087019361ffff8160901c16855260a01c1615801560c08801526142705761ffff168061420a5750505063ffffffff808061ffff9351169451169551169351169193929190600190565b919550915061ffff600b54169081811061424057505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050505092505050600090600090600090600090600090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561093d570180359067ffffffffffffffff821161093d5760200191813603831361093d57565b356001600160a01b038116810361093d5790565b3567ffffffffffffffff8116810361093d5790565b9067ffffffffffffffff613f4f92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613d2a5760051b60200190565b92919061436481614340565b936143726040519586613db0565b602085838152019160051b810192831161093d57905b82821061439457505050565b602080916143a184613e5c565b815201910190614388565b67ffffffffffffffff16600052600e60205260406000209160028110156143f3576001146143e257816001613f4f930190614fba565b8160026003613f4f94019101614fba565b634e487b7160e01b600052602160045260246000fd5b91908110156144195760051b0190565b634e487b7160e01b600052603260045260246000fd5b9190811015614419576060020190565b6040519061444c82613d78565b60006080838281528260208201528260408201528260608201520152565b9060405161447781613d78565b60806001829460ff81546001600160801b038116865263ffffffff81861c16602087015260a01c161515604085015201546001600160801b0381166060840152811c910152565b604051906144cb82613d5c565b60606020838281520152565b80518210156144195760209160051b010190565b90600182811c9216801561451b575b602083101461450557565b634e487b7160e01b600052602260045260246000fd5b91607f16916144fa565b9060405191826000825492614539846144eb565b80845293600181169081156145a75750600114614560575b5061455e92500383613db0565b565b90506000929192526020600020906000915b81831061458b57505090602061455e9282010138614551565b6020919350806001915483858901015201910190918492614572565b6020935061455e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614551565b601f8260209493601f19938186528686013760008582860101520116010190565b67ffffffffffffffff166000526008602052613f4f6004604060002001614525565b91908110156144195760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561093d570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561093d570180359067ffffffffffffffff821161093d57602001918160051b3603831361093d57565b818102929181159184041417156146d157565b634e487b7160e01b600052601160045260246000fd5b8181106146f2575050565b600081556001016146e7565b9160209082815201919060005b8181106147185750505090565b9091926020806001926001600160a01b0361473288613e5c565b16815201940192910161470b565b91908110156144195760081b0190565b35801515810361093d5790565b3561ffff8116810361093d5790565b3563ffffffff8116810361093d5790565b359063ffffffff8216820361093d57565b805180156147fe576020036147c057805160208281019183018390031261093d57519060ff82116147c0575060ff1690565b6122d9906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613def565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116146d157565b60ff16604d81116146d157600a0a90565b8115614853570490565b634e487b7160e01b600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146149515782841161492757906148ae91614824565b91604d60ff841611801561490c575b6148d6575050906148d0613f4f92614838565b906146be565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061491683614838565b8015614853576000190484116148bd565b61493091614824565b91604d60ff8416116148d65750509061494b613f4f92614838565b90614849565b5050505090565b9081602091031261093d5751801515810361093d5790565b6001600160a01b0360015416330361498457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b356001600160801b038116810361093d5790565b6001600160801b036149f7604080936149da816140e5565b15158652836149eb602083016140f2565b166020870152016140f2565b16910152565b9160005b82811015614d4357600060e082028501614a1a816142ee565b9067ffffffffffffffff821692614a3e846000526007602052604060002054151590565b15614d17575060019392614b777f20ae59542ddd78610f62f9d2c9dcd658f8b6a5b45a0f03e71e16614e89dda83693614b6d84614b5d602060e0970191614a8d614a883685614106565b61548d565b614af0614aae8667ffffffffffffffff16600052600c602052604060002090565b805463ffffffff8160801c16159081614d02575b81614cf3575b81614ce1575b81614cd2575b5080614cc3575b614c4b575b614aea3686614106565b9061559e565b614b1f6080820195614b05614a883689614106565b67ffffffffffffffff16600052600d602052604060002090565b90815463ffffffff8160801c16159081614c36575b81614c27575b81614c15575b81614c06575b5080614bf7575b614b7e575b50614aea3686614106565b60405194855260208501906149c2565b60808301906149c2565ba101614a01565b614b9260a06001600160801b0392016149ae565b825473ffffffff000000000000000000000000000000004260801b1673ffffffffffffffffffffffffffffffffffffffff19909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717815538614b52565b50614c0186614750565b614b4d565b60ff915060a01c161538614b46565b6001600160801b038116159150614b40565b838e015460801c159150614b3a565b838e01546001600160801b0316159150614b34565b6001600160801b03614c5f604085016149ae565b825473ffffffff000000000000000000000000000000004260801b1673ffffffffffffffffffffffffffffffffffffffff19909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178155614ae0565b50614ccd85614750565b614adb565b60ff915060a01c161538614ad4565b6001600160801b038116159150614ace565b828f015460801c159150614ac8565b828f01546001600160801b0316159150614ac2565b80847f1e670e4b0000000000000000000000000000000000000000000000000000000060249352600452fd5b50915050565b90805115614f2f5767ffffffffffffffff81516020830120921691826000526008602052614d7e816005604060002001615b15565b15614eeb5760005260096020526040600020815167ffffffffffffffff8111613d2a57614dab82546144eb565b601f8111614eb9575b506020601f8211600114614e2f5791614e09827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614e1f95600091614e24575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613def565b0390a2565b905084015138614df6565b601f1982169083600052806000209160005b818110614ea1575092614e1f9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614e88575b5050811b019055611b98565b85015160001960f88460031b161c191690553880614e7c565b9192602060018192868a015181550194019201614e41565b614ee590836000526020600020601f840160051c8101916020851061088557601f0160051c01906146e7565b38614db4565b50906122d96040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613def565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b818110614f8b57505061455e92500383613db0565b84546001600160a01b0316835260019485019487945060209093019201614f76565b919082018092116146d157565b614fc390614f59565b9160055480151591826150a3575b5050614fdb575090565b614fe490614f59565b90815180614ff25750905090565b614ffd908251614fad565b92601f1961502361500d86614340565b9561501b6040519788613db0565b808752614340565b0136602086013760005b825181101561505e57806001600160a01b0361504b600193866144d7565b511661505782886144d7565b520161502d565b509160005b815181101561509e57806001600160a01b03615081600193856144d7565b5116615097615091838751614fad565b886144d7565b5201615063565b505050565b101590503880614fd1565b67ffffffffffffffff16600081815260076020526040902054909291901561519e579161519b60e092615170856151057f73d6dce40db73cbddae4b9ce52576043a1644e08c2702236273d71077435fa169761548d565b84600052600860205261511c81604060002061559e565b6151258361548d565b84600052600860205261513f83600260406000200161559e565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b919082039182116146d157565b6151e161443f565b506001600160801b036060820151166001600160801b03808351169161522c602085019361522661521963ffffffff875116426151cc565b85608089015116906146be565b90614fad565b8082101561524557505b16825263ffffffff4216905290565b9050615236565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561093d57604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af180156152d1576152c4575050565b816152ce91613db0565b50565b6040513d84823e3d90fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613f4f604082613db0565b9061ffff9067ffffffffffffffff615331602085016142ee565b16600052600f6020526040600020826040519161534d83613d0e565b549263ffffffff8416835263ffffffff8460201c16602084015263ffffffff8460401c16604084015263ffffffff8460601c166060840152818460801c169283608082015260c060ff848760901c16968760a085015260a01c1615159101521615156000146153e357505b1680156153db576127106153d46060613f4f94013592836146be565b04906151cc565b506060013590565b90506153b8565b805160005b8181106153fb57505050565b600181018082116146d1575b82811061541757506001016153ef565b6001600160a01b0361542983866144d7565b51166001600160a01b0361543d83876144d7565b51161461544c57600101615407565b6001600160a01b0361545e83866144d7565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115615512576001600160801b036040820151166001600160801b03602083015116106154b85750565b606490615510604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565bfd5b6001600160801b0360408201511615801590615588575b6155305750565b606490615510604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565b506001600160801b036020820151161515615529565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916156c560609280546155db63ffffffff8260801c16426151cc565b90816156fb575b50506001600160801b0360018160208601511692828154168085106000146156f357508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556156828651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166001600160801b031692909217910155565b61519b60405180926001600160801b0360408092805115158552826020820151166020860152015116910152565b838091615609565b6001600160801b03916157278392836157206001880154948286169560801c906146be565b9116614fad565b8082101561579b57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff92909116929092161673ffffffffffffffffffffffffffffffffffffffff19909116174260801b73ffffffff000000000000000000000000000000001617815538806155e2565b9050615731565b906040519182815491828252602082019060005260206000209260005b8181106157d457505061455e92500383613db0565b84548352600194850194879450602090930192016157bf565b80548210156144195760005260206000200190600090565b60008181526003602052604090205480156158e55760001981018181116146d1576002549060001982019182116146d157818103615894575b505050600254801561587e57600019016158598160026157ed565b60001982549160031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6158cd6158a56158b69360026157ed565b90549060031b1c92839260026157ed565b81939154906000199060031b92831b921b19161790565b9055600052600360205260406000205538808061583e565b5050600090565b60008181526007602052604090205480156158e55760001981018181116146d1576006549060001982019182116146d157818103615965575b505050600654801561587e57600019016159408160066157ed565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b6159876159766158b69360066157ed565b90549060031b1c92839260066157ed565b90556000526007602052604060002055388080615925565b9060018201918160005282602052604060002054801515600014615a525760001981018181116146d15782549060001982019182116146d157818103615a1b575b5050508054801561587e5760001901906159fa82826157ed565b60001982549160031b1b191690555560005260205260006040812055600190565b615a3b615a2b6158b693866157ed565b90549060031b1c928392866157ed565b9055600052836020526040600020553880806159e0565b50505050600090565b80600052600360205260406000205415600014615ab55760025468010000000000000000811015613d2a57615a9c6158b682600185940160025560026157ed565b9055600254906000526003602052604060002055600190565b50600090565b80600052600760205260406000205415600014615ab55760065468010000000000000000811015613d2a57615afc6158b682600185940160065560066157ed565b9055600654906000526007602052604060002055600190565b60008281526001820160205260409020546158e55780549068010000000000000000821015613d2a5782615b536158b68460018096018555846157ed565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615d56575b615d50576001600160801b0382169160018501908154615bb063ffffffff6001600160801b0383169360801c16426151cc565b9081615cb2575b5050848110615c735750838310615c08575050615bdd6001600160801b039283926151cc565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615c1781856151cc565b926000198101908082116146d157615c3a615c3f926001600160a01b0396614fad565b614849565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615d2657615ccd926152269160801c906146be565b80841015615d215750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615bb7565b615cd8565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615b7d565b91929015615dd95750815115615d72575090565b3b15615d7b5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615dec5750805190602001fd5b6122d9906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613def56fea164736f6c634300081a000a",
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

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getCurrentCustomBlockConfirmationRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentCustomBlockConfirmationRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _BurnMintTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _BurnMintTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _BurnMintTokenPool.Contract.GetCurrentRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _BurnMintTokenPool.Contract.GetCurrentRateLimiterState(&_BurnMintTokenPool.CallOpts, remoteChainSelector)
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

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetFee(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getFee", arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnMintTokenPool.Contract.GetFee(&_BurnMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetFee(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnMintTokenPool.Contract.GetFee(&_BurnMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BurnMintTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _BurnMintTokenPool.Contract.GetMinBlockConfirmation(&_BurnMintTokenPool.CallOpts)
}

func (_BurnMintTokenPool *BurnMintTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _BurnMintTokenPool.Contract.GetMinBlockConfirmation(&_BurnMintTokenPool.CallOpts)
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

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyCustomBlockConfirmationConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_BurnMintTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_BurnMintTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintTokenPool.TransactOpts, tokenTransferFeeConfigArgs, disableTokenTransferFeeConfigs)
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

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn0(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.LockOrBurn0(&_BurnMintTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, arg2)
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

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint0(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, blockConfirmationRequested uint16) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.ReleaseOrMint0(&_BurnMintTokenPool.TransactOpts, releaseOrMintIn, blockConfirmationRequested)
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

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "setCustomBlockConfirmationRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_BurnMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_BurnMintTokenPool.TransactOpts, rateLimitConfigArgs)
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

func (_BurnMintTokenPool *BurnMintTokenPoolTransactor) WithdrawFeeTokens(opts *bind.TransactOpts, feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.contract.Transact(opts, "withdrawFeeTokens", feeTokens, recipient)
}

func (_BurnMintTokenPool *BurnMintTokenPoolSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.WithdrawFeeTokens(&_BurnMintTokenPool.TransactOpts, feeTokens, recipient)
}

func (_BurnMintTokenPool *BurnMintTokenPoolTransactorSession) WithdrawFeeTokens(feeTokens []common.Address, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPool.Contract.WithdrawFeeTokens(&_BurnMintTokenPool.TransactOpts, feeTokens, recipient)
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

type BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _BurnMintTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator struct {
	Event *BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured)
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
		it.Event = new(BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured)
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

func (it *BurnMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCustomBlockConfirmationRateLimitConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator{contract: _BurnMintTokenPool.contract, event: "CustomBlockConfirmationRateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCustomBlockConfirmationRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationRateLimitConfigured", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCustomBlockConfirmationRateLimitConfigured(log types.Log) (*BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured, error) {
	event := new(BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationRateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolCustomBlockConfirmationUpdatedIterator struct {
	Event *BurnMintTokenPoolCustomBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolCustomBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolCustomBlockConfirmationUpdated)
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
		it.Event = new(BurnMintTokenPoolCustomBlockConfirmationUpdated)
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

func (it *BurnMintTokenPoolCustomBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolCustomBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolCustomBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolCustomBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolCustomBlockConfirmationUpdatedIterator{contract: _BurnMintTokenPool.contract, event: "CustomBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolCustomBlockConfirmationUpdated)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseCustomBlockConfirmationUpdated(log types.Log) (*BurnMintTokenPoolCustomBlockConfirmationUpdated, error) {
	event := new(BurnMintTokenPoolCustomBlockConfirmationUpdated)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolDefaultFinalityRateLimitConfiguredIterator struct {
	Event *BurnMintTokenPoolDefaultFinalityRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolDefaultFinalityRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolDefaultFinalityRateLimitConfigured)
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
		it.Event = new(BurnMintTokenPoolDefaultFinalityRateLimitConfigured)
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

func (it *BurnMintTokenPoolDefaultFinalityRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolDefaultFinalityRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolDefaultFinalityRateLimitConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterDefaultFinalityRateLimitConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolDefaultFinalityRateLimitConfiguredIterator, error) {

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "DefaultFinalityRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolDefaultFinalityRateLimitConfiguredIterator{contract: _BurnMintTokenPool.contract, event: "DefaultFinalityRateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchDefaultFinalityRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolDefaultFinalityRateLimitConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "DefaultFinalityRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolDefaultFinalityRateLimitConfigured)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "DefaultFinalityRateLimitConfigured", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseDefaultFinalityRateLimitConfigured(log types.Log) (*BurnMintTokenPoolDefaultFinalityRateLimitConfigured, error) {
	event := new(BurnMintTokenPoolDefaultFinalityRateLimitConfigured)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "DefaultFinalityRateLimitConfigured", log); err != nil {
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

type BurnMintTokenPoolFeeTokenWithdrawnIterator struct {
	Event *BurnMintTokenPoolFeeTokenWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolFeeTokenWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolFeeTokenWithdrawn)
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
		it.Event = new(BurnMintTokenPoolFeeTokenWithdrawn)
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

func (it *BurnMintTokenPoolFeeTokenWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolFeeTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolFeeTokenWithdrawn struct {
	Recipient common.Address
	FeeToken  common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*BurnMintTokenPoolFeeTokenWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.FilterLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolFeeTokenWithdrawnIterator{contract: _BurnMintTokenPool.contract, event: "FeeTokenWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var feeTokenRule []interface{}
	for _, feeTokenItem := range feeToken {
		feeTokenRule = append(feeTokenRule, feeTokenItem)
	}

	logs, sub, err := _BurnMintTokenPool.contract.WatchLogs(opts, "FeeTokenWithdrawn", recipientRule, feeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolFeeTokenWithdrawn)
				if err := _BurnMintTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (_BurnMintTokenPool *BurnMintTokenPoolFilterer) ParseFeeTokenWithdrawn(log types.Log) (*BurnMintTokenPoolFeeTokenWithdrawn, error) {
	event := new(BurnMintTokenPoolFeeTokenWithdrawn)
	if err := _BurnMintTokenPool.contract.UnpackLog(event, "FeeTokenWithdrawn", log); err != nil {
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

func (BurnMintTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnMintTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x20ae59542ddd78610f62f9d2c9dcd658f8b6a5b45a0f03e71e16614e89dda836")
}

func (BurnMintTokenPoolCustomBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e")
}

func (BurnMintTokenPoolDefaultFinalityRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x73d6dce40db73cbddae4b9ce52576043a1644e08c2702236273d71077435fa16")
}

func (BurnMintTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (BurnMintTokenPoolFeeTokenWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e")
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
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (_BurnMintTokenPool *BurnMintTokenPool) Address() common.Address {
	return _BurnMintTokenPool.address
}

type BurnMintTokenPoolInterface interface {
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

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

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnMintTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnMintTokenPoolConfigChanged, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationRateLimitConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator, error)

	WatchCustomBlockConfirmationRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured) (event.Subscription, error)

	ParseCustomBlockConfirmationRateLimitConfigured(log types.Log) (*BurnMintTokenPoolCustomBlockConfirmationRateLimitConfigured, error)

	FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolCustomBlockConfirmationUpdatedIterator, error)

	WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomBlockConfirmationUpdated(log types.Log) (*BurnMintTokenPoolCustomBlockConfirmationUpdated, error)

	FilterDefaultFinalityRateLimitConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolDefaultFinalityRateLimitConfiguredIterator, error)

	WatchDefaultFinalityRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolDefaultFinalityRateLimitConfigured) (event.Subscription, error)

	ParseDefaultFinalityRateLimitConfigured(log types.Log) (*BurnMintTokenPoolDefaultFinalityRateLimitConfigured, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*BurnMintTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*BurnMintTokenPoolDynamicConfigSet, error)

	FilterFeeTokenWithdrawn(opts *bind.FilterOpts, recipient []common.Address, feeToken []common.Address) (*BurnMintTokenPoolFeeTokenWithdrawnIterator, error)

	WatchFeeTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolFeeTokenWithdrawn, recipient []common.Address, feeToken []common.Address) (event.Subscription, error)

	ParseFeeTokenWithdrawn(log types.Log) (*BurnMintTokenPoolFeeTokenWithdrawn, error)

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
