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

var BurnFromMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"disableTokenTransferFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentCustomBlockConfirmationRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"blockConfirmationConfigured\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFeeTokens\",\"inputs\":[{\"name\":\"feeTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationRateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultFinalityRateLimitConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeTokenWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feeToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103a8576166c9803803809161001d8285610606565b8339810160a0828203126103a85781516001600160a01b03811692908390036103a85761004c60208201610629565b60408201516001600160401b0381116103a85782019280601f850112156103a8578351936001600160401b0385116103ad578460051b9060208201956100956040519788610606565b86526020808701928201019283116103a857602001905b8282106105ee575050506100ce60806100c760608501610637565b9301610637565b9133156105dd57600180546001600160a01b03191633179055841580156105cc575b80156105bb575b6105aa57608085905260c05260405163313ce56760e01b8152602081600481885afa6000918161056e575b50610543575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610426575b50604051636eb1769f60e11b81523060048201819052602482015290602082604481845afa91821561041a576000926103e6575b5060001982018092116103d057604051602081019263095ea7b360e01b84523060248301526044820152604481526101c7606482610606565b6000806040948551936101da8786610606565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d156103c3573d906001600160401b0382116103ad57845161024b94909261023c601f8201601f191660200185610606565b83523d6000602085013e6107d5565b80518061032d575b8251615e2390816108a68239608051818181611c0c01528181611de701528181611f0101528181611fe101528181612504015281816126ef015281816130a20152818161327c015281816132d601528181613443015281816135970152818161374901528181613b1301528181613b680152615236015260a051818181613acf015281816147e20152818161484c01526152cf015260c051818181610c7f01528181611c9b01528181612592015281816131310152613627015260e051818181610b2301528181611ce0015281816125d70152612e930152f35b81602091810103126103a857602001518015908115036103a857610352573880610253565b5162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b9161024b926060916107d5565b634e487b7160e01b600052601160045260246000fd5b9091506020813d602011610412575b8161040260209383610606565b810103126103a85751903861018e565b3d91506103f5565b6040513d6000823e3d90fd5b60206040516104358282610606565b60008152600036813760e051156105325760005b81518110156104b0576001906001600160a01b03610467828561064b565b5116846104738261068d565b610480575b505001610449565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13884610478565b505060005b8251811015610529576001906001600160a01b036104d3828661064b565b5116801561052357836104e582610775565b6104f3575b50505b016104b5565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138836104ea565b506104ed565b5050503861015a565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036105575750610128565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116105a2575b8161058a60209383610606565b810103126103a85761059b90610629565b9038610122565b3d915061057d565b630a64406560e11b60005260046000fd5b506001600160a01b038116156100f7565b506001600160a01b038316156100f0565b639b15e16f60e01b60005260046000fd5b602080916105fb84610637565b8152019101906100ac565b601f909101601f19168101906001600160401b038211908210176103ad57604052565b519060ff821682036103a857565b51906001600160a01b03821682036103a857565b805182101561065f5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b805482101561065f5760005260206000200190600090565b600081815260036020526040902054801561076e5760001981018181116103d0576002546000198101919082116103d05781810361071d575b505050600254801561070757600019016106e1816002610675565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61075661072e61073f936002610675565b90549060031b1c9283926002610675565b819391549060031b91821b91600019901b19161790565b905560005260036020526040600020553880806106c6565b5050600090565b806000526003602052604060002054156000146107cf57600254680100000000000000008110156103ad576107b661073f8260018594016002556002610675565b9055600254906000526003602052604060002055600190565b50600090565b9192901561083757508151156107e9575090565b3b156107f25790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b82519091501561084a5750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b83811061088d5750508160006044809484010152601f80199101168101030190fd5b6020828201810151604487840101528593500161086b56fe60c080604052600436101561001357600080fd5b600060a05260a0513560e01c90816301ffc9a714613bee57508063181f5a7714613b8c57806321df0da714613b47578063240028e814613af357806324f65ee714613ab45780632c06340414613a1c57806337b19247146138ae5780633907753714613510578063489a68f21461300a5780634c5ef0ed14612fc557806354c8a4f314612e6157806359152aad14612db457806362ddd3c414612d2f578063698c2c6614612c885780636d3d1a5814612c605780637437ff9f14612c2c57806379ba509714612b7a5780637d54534e14612b015780638926f54f14612abc57806389720a6214612a515780638da5cb5b14612a29578063962d4020146128e65780639751f884146128815780639a4575b9146124b2578063a42a7b8b1461235b578063a7cd63b7146122ed578063acfecf91146121c0578063b1c71c6514611b90578063b794658014611b54578063bb6bb5a714611aec578063c4bffe2b146119d0578063c7230a60146116b4578063cf7401f31461154b578063d966866b146110fe578063dc04fa1f14610ca3578063dc0bd97114610c5e578063ded8d95614610b48578063e0351e1314610b0a578063e8a1da17146102c5578063f2fde38b146102115763fa41d79c146101e857600080fd5b3461020b5760a05160031936011261020b57602061ffff600b5416604051908152f35b60a05180fd5b3461020b57602060031936011261020b576001600160a01b03610232613e10565b61023a614950565b163381146102995760a051805473ffffffffffffffffffffffffffffffffffffffff1916821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360a05180f35b7fdad89dca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461020b576102d336613f63565b9190926102de614950565b60a051905b8282106109485750505060a0519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b8181101561094257600581901b8301358581121561020b5783016101208136031261020b576040519461035286613d58565b813567ffffffffffffffff8116810361093d578652602082013567ffffffffffffffff811161020b5782019436601f8701121561020b57853561039481614320565b966103a26040519889613d90565b81885260208089019260051b8201019036821161020b5760208101925b82841061090e575050505060208701958652604083013567ffffffffffffffff811161020b576103f29036908501613f14565b916040880192835261041c61040a36606087016140e6565b9460608a0195865260c03691016140e6565b956080890196875261042e8551615479565b6104388751615479565b835151156108e25761045467ffffffffffffffff8a5116615aa7565b156108a75767ffffffffffffffff89511660a051526008602052604060a0512061057486516001600160801b03604082015116906105386001600160801b03602083015116915115158360806040516104ac81613d58565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160801b0384161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176001830155565b61067688516001600160801b036040820151169061063a6001600160801b03602083015116915115158360806040516105ac81613d58565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166001600160801b0385161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176003830155565b6004855191019080519067ffffffffffffffff821161088f5761069983546144cb565b601f8111610850575b506020906001601f8411146107e6579180916106d59360a051926107db575b50506000198260011b9260031b1c19161790565b90555b60a0515b88518051821015610711579061070b6001926107048367ffffffffffffffff8f5116926144b7565b5190614d29565b016106dc565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29391999750956107cd67ffffffffffffffff60019796949851169251935191516107a261077660405196879687526101006020880152610100870190613dcf565b9360408601906001600160801b0360408092805115158552826020820151166020860152015116910152565b60a08401906001600160801b0360408092805115158552826020820151166020860152015116910152565b0390a1019392909193610320565b015190508e806106c1565b90601f198316918460a051528160a051209260a0515b818110610838575090846001959493921061081f575b505050811b0190556106d8565b015160001960f88460031b161c191690558d8080610812565b929360206001819287860151815501950193016107fc565b61087f908460a05152602060a05120601f850160051c81019160208610610885575b601f0160051c01906146c7565b8d6106a2565b9091508190610872565b634e487b7160e01b60a051526041600452602460a051fd5b67ffffffffffffffff8951167f1d5ad3c50000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b833567ffffffffffffffff811161020b576020916109328392833691870101613f14565b8152019301926103bf565b600080fd5b60a05180f35b9092919367ffffffffffffffff6109686109638688866143e9565b6142ce565b1692610973846158d8565b15610ada578360a0515260086020526109936005604060a051200161578e565b9260a0515b84518110156109d2576001908660a0515260086020526109cb6005604060a05120016109c483896144b7565b519061598b565b5001610998565b50939094919592508060a0515260086020526005604060a0512060a051815560a051600182015560a051600282015560a051600382015560048101610a1781546144cb565b80610a8a575b50500180549060a051815581610a66575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102e3565b60a05152602060a05120908101905b81811015610a2e5760a0518155600101610a75565b601f8111600114610aa3575060a05190555b8880610a1d565b610ac3908260a051526001601f602060a051209201861c820191016146c7565b60a080518290525160208120918190559055610a9c565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461020b5760a05160031936011261020b5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461020b57602060031936011261020b57610140610b64613e67565b610b6c61441f565b50610b7561441f565b50610c5c610bcb610bab610ba6610bb0610bab610ba68767ffffffffffffffff16600052600c602052604060002090565b61444a565b6151b9565b9467ffffffffffffffff16600052600d602052604060002090565b610c1560405180946001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b3461020b5760a05160031936011261020b5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020b57604060031936011261020b5760043567ffffffffffffffff811161020b573660238201121561020b57806004013567ffffffffffffffff811161020b576024820191602436918360081b01011161020b5760243567ffffffffffffffff811161020b57610d19903690600401613f32565b919092610d24614950565b60a0515b828110610da05750505060a0515b818110610d435760a05180f35b8067ffffffffffffffff610d5d61096360019486886143e9565b168060a05152600f60205260a051604060a05120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee860a05160a051a201610d36565b610dae610963828585614720565b610db9828585614720565b90602082019060e0830190610dcd82614730565b1561107c5760a0840161271061ffff610de58361473d565b1610156110f25760c085019161271061ffff610e008561473d565b1610156110b65763ffffffff610e158661474c565b161561107c5767ffffffffffffffff16948560a05152600f60205260a05160409020610e408661474c565b63ffffffff16908054906040840191610e588361474c565b60201b67ffffffff0000000016936060860194610e748661474c565b60401b6bffffffff0000000000000000169660800196610e938861474c565b60601b6fffffffff0000000000000000000000001691610eb28a61473d565b60801b71ffff000000000000000000000000000000001693610ed38c61473d565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff161717178155610f8687614730565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610fd79061475d565b63ffffffff168752610fe89061475d565b63ffffffff166020870152610ffc9061475d565b63ffffffff1660408601526110109061475d565b63ffffffff16606085015261102490613ea0565b61ffff16608084015261103690613ea0565b61ffff1660a0830152611048906140c5565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610d28565b67ffffffffffffffff907f123322650000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b61ffff6110c28461473d565b7f95f3517a0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b6110c261ffff9161473d565b3461020b57602060031936011261020b5760043567ffffffffffffffff811161020b5761112f903690600401613f32565b90611138614950565b60a0516080525b816080511061114e5760a05180f35b61115e610963608051848461460a565b61117861116e608051858561460a565b602081019061464a565b611192611188608051878761460a565b604081019061464a565b906111ad6111a3608051898961460a565b606081019061464a565b6111c76111bd6080518b8b61460a565b608081019061464a565b9390946111dd6111d836898b614338565b6153d6565b6111eb6111d8368385614338565b6111f96111d8368587614338565b6112076111d8368789614338565b6040519860808a01908a821067ffffffffffffffff83111761088f5767ffffffffffffffff9160405261123b368a8c614338565b8b52611248368486614338565b60208c0152611258368688614338565b60408c015261126836888a614338565b60608c015216988960a05152600e602052604060a05120815180519067ffffffffffffffff821161088f5768010000000000000000821161088f57602090835483855580841061152c575b50018260a05152602060a0512060a0515b83811061150f57505050506001810160208301519081519167ffffffffffffffff831161088f5768010000000000000000831161088f5760209082548484558085106114f0575b50019060a05152602060a0512060a0515b8381106114d357505050506002810160408301519081519167ffffffffffffffff831161088f5768010000000000000000831161088f5760209082548484558085106114b4575b50019060a05152602060a0512060a0515b838110611497575050505060036060910191015180519067ffffffffffffffff821161088f5768010000000000000000821161088f576020908354838555808410611478575b50019160a05152602060a051209160a0515b82811061145b57505050509261144a949261142e61143c937fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9998966114206040519b8c9b60808d5260808d01916146de565b918a830360208c01526146de565b9187830360408901526146de565b9184830360608601526146de565b0390a260016080510160805261113f565b60019060206001600160a01b0384511693019281860155016113cc565b611491908560a05152848460a0512091820191016146c7565b8f6113ba565b60019060206001600160a01b038551169401938184015501611374565b6114cd908460a05152858460a0512091820191016146c7565b38611363565b60019060206001600160a01b03855116940193818401550161131c565b611509908460a05152858460a0512091820191016146c7565b3861130b565b60019060206001600160a01b0385511694019381840155016112c4565b611545908560a05152848460a0512091820191016146c7565b386112b3565b3461020b5760e060031936011261020b57611564613e67565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc011261020b5760405161159a81613d74565b602435801515810361020b5781526044356001600160801b038116810361020b5760208201526064356001600160801b038116810361020b5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c011261020b576040519061160f82613d74565b608435801515810361020b57825260a4356001600160801b038116810361020b57602083015260c4356001600160801b038116810361020b5760408301526001600160a01b03600a54163314158061169f575b61166f576109429261508e565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506001600160a01b0360015416331415611662565b3461020b57604060031936011261020b5760043567ffffffffffffffff811161020b576116e5903690600401613f32565b6024356001600160a01b03811680820361020b57611701614950565b60a0515b8381106117125760a05180f35b8060206001600160a01b0361173261172d602495898b6143e9565b6142ba565b16604051938480927f70a082310000000000000000000000000000000000000000000000000000000082523060048301525afa80156119c357838591879460a05191611984575b508061178d575b5050506001915001611705565b886118756001600160a01b036117a761172d888a866143e9565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082019081526001600160a01b039890981660248201526044808201879052815291166117fa606483613d90565b604080519097909261180c8985613d90565b602084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602085015260a0519160a05191519060a051855af13d1561197c573d9161185783613db3565b926118648a519485613d90565b835260a0513d90602085013e615d4a565b8051806118cd575b50507f508d7d183612c18fc339b42618912b9fa3239f631dd7ec0671f950200a0fa66e916001600160a01b036118bb61172d8860019a6020966143e9565b169451908152a3849150838388611780565b8160209350839294969798506118e7955001019101614938565b156118f9579083869392888a8061187d565b608482517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615d4a565b93945050505060203d81116119bc575b61199e8183613d90565b602082600092810103126119b9575090838386935189611779565b80fd5b503d611994565b6040513d60a051823e3d90fd5b3461020b5760a05160031936011261020b5760a051506040516006548082528160208101600660a05152602060a051209260a0515b818110611ad3575050611a1a92500382613d90565b805190611a3f611a2983614320565b92611a376040519485613d90565b808452614320565b90601f1960208401920136833760a0515b8151811015611a82578067ffffffffffffffff611a6f600193856144b7565b5116611a7b82876144b7565b5201611a50565b5050906040519182916020830190602084525180915260408301919060a0515b818110611ab0575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611aa2565b8454835260019485019486945060209093019201611a05565b3461020b57602060031936011261020b5760043567ffffffffffffffff811161020b57611b1d903690600401613fb5565b6001600160a01b03600a541633141580611b3f575b61166f57610942916149dd565b506001600160a01b0360015416331415611b32565b3461020b57602060031936011261020b57611b8c611b78611b73613e67565b6145e8565b604051918291602083526020830190613dcf565b0390f35b3461020b57606060031936011261020b5760043567ffffffffffffffff811161020b5760a0600319823603011261020b57611bc9613e8f565b9060443567ffffffffffffffff811161020b57611bea903690600401613f14565b50611bf361449e565b5060848101611c01816142ba565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361217f5750602481019077ffffffffffffffff00000000000000000000000000000000611c5b836142ce565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119c35760a05191612150575b5061212457611cde604482016142ba565b7f00000000000000000000000000000000000000000000000000000000000000006120d1575b5067ffffffffffffffff611d17836142ce565b16611d2f816000526007602052604060002054151590565b156120a25760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156119c35760a05190612059575b6001600160a01b03915016330361202957606481013561ffff84168015611f7e5761ffff600b54169081611e96575b505050611b73611dce611e8c94611e5b935b600401615303565b92611dd88461522c565b611de1816142ce565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a26142ce565b90611e646152c8565b60405192611e7184613d3c565b8352602083015260405192839260408452604084019061409b565b9060208301520390f35b818110611f4c575050611dce611e8c94611e5b93611b73937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611ee1896142ce565b16918260a05152600c60205280611f29604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615b56565b604080516001600160a01b039290921682526020820192909252a2935094611db4565b7f7911d95b0000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b50611dce611e8c94611e5b93611b73937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611fc1896142ce565b16918260a05152600860205280612009604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615b56565b604080516001600160a01b039290921682526020820192909252a2611dc6565b7f728fe07b0000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506020813d60201161209a575b8161207360209383613d90565b8101031261020b57516001600160a01b038116810361020b576001600160a01b0390611d85565b3d9150612066565b7fa9902c7e0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b6001600160a01b03166120f1816000526003602052604060002054151590565b611d04577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b612172915060203d602011612178575b61216a8183613d90565b810190614938565b84611ccd565b503d612160565b6121906001600160a01b03916142ba565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b3461020b5767ffffffffffffffff6121d736613fe6565b9290916121e2614950565b16906121fb826000526007602052604060002054151590565b156122bd578160a05152600860205261222e6005604060a0512001612221368685613edd565b602081519101209061598b565b15612276577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76919261226d6040519283926020845260208401916145c7565b0390a260a05180f35b6122b9906040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916145c7565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461020b5760a05160031936011261020b5760a051506040516002548082526020820190600260a05152602060a051209060a0515b81811061234557611b8c8561233981870382613d90565b60405191829182614027565b8254845260209093019260019283019201612322565b3461020b57602060031936011261020b5767ffffffffffffffff61237d613e67565b1660a0515260086020526123986005604060a051200161578e565b805190601f196123c06123aa84614320565b936123b86040519586613d90565b808552614320565b0160a0515b8181106124a157505060a0515b815181101561241c57806123e8600192846144b7565b5160a051526009602052612400604060a05120614505565b61240a82866144b7565b5261241581856144b7565b50016123d2565b826040518091602082016020835281518091526040830190602060408260051b86010193019160a051905b82821061245657505050500390f35b91936020612491827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613dcf565b9601920192018594939192612447565b8060606020809387010152016123c5565b3461020b57602060031936011261020b5760043567ffffffffffffffff811161020b5760a0600319823603011261020b576124eb61449e565b50608481016124f9816142ba565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361217f57506024810177ffffffffffffffff00000000000000000000000000000000612552826142ce565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119c35760a05191612862575b50612124576125d5604483016142ba565b7f000000000000000000000000000000000000000000000000000000000000000061280f575b5067ffffffffffffffff61260e826142ce565b16612626816000526007602052604060002054151590565b156120a25760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa80156119c35760a051906127c6575b6001600160a01b03915016330361202957611b8c9161279691611b7391606401357ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108167ffffffffffffffff6126d1856142ce565b168060a051526008602052612717604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958691615b56565b604080516001600160a01b0386168152602081018490527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789449190a261275b8161522c565b67ffffffffffffffff61276d856142ce565b604080516001600160a01b03909616865233602087015285019290925216918060608101611e53565b61279e6152c8565b604051916127ab83613d3c565b8252602082015260405191829160208352602083019061409b565b506020813d602011612807575b816127e060209383613d90565b8101031261020b57516001600160a01b038116810361020b576001600160a01b039061267c565b3d91506127d3565b6001600160a01b031661282f816000526003602052604060002054151590565b6125fb577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b61287b915060203d6020116121785761216a8183613d90565b836125c4565b3461020b57602060031936011261020b5767ffffffffffffffff6128a3613e67565b6128ab61441f565b506128b461441f565b501660a051526008602052610140604060a05120610c5c610bcb610bab60026128df610bab8661444a565b940161444a565b3461020b57606060031936011261020b5760043567ffffffffffffffff811161020b57612917903690600401613f32565b9060243567ffffffffffffffff811161020b5761293890369060040161406a565b9060443567ffffffffffffffff811161020b5761295990369060040161406a565b6001600160a01b03600a541633141580612a14575b61166f57838614801590612a0a575b6129de5760a0515b8681106129925760a05180f35b806129d86129a66109636001948b8b6143e9565b6129b183898961440f565b6129d26129ca6129c286898b61440f565b9236906140e6565b9136906140e6565b9161508e565b01612985565b7f568efce20000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b508086141561297d565b506001600160a01b036001541633141561296e565b3461020b5760a05160031936011261020b5760206001600160a01b0360015416604051908152f35b3461020b5760c060031936011261020b57612a6a613e10565b50612a73613e50565b612a7b613e7e565b5060843567ffffffffffffffff811161020b57612a9c903690600401613eaf565b505060a43590600282101561020b57611b8c91612339916044359061438c565b3461020b57602060031936011261020b576020612af767ffffffffffffffff612ae3613e67565b166000526007602052604060002054151590565b6040519015158152f35b3461020b57602060031936011261020b577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b03612b45613e10565b612b4d614950565b168073ffffffffffffffffffffffffffffffffffffffff19600a541617600a55604051908152a160a05180f35b3461020b5760a05160031936011261020b5760a051546001600160a01b0381163303612c005773ffffffffffffffffffffffffffffffffffffffff196001549133828416176001551660a051556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060a05160a051a360a05180f35b7f02b543c60000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461020b5760a05160031936011261020b57600454600554604080516001600160a01b039093168352602083019190915290f35b3461020b5760a05160031936011261020b5760206001600160a01b03600a5416604051908152f35b3461020b57604060031936011261020b57612ca1613e10565b602435612cac614950565b6001600160a01b0382169182156108e2577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f02389273ffffffffffffffffffffffffffffffffffffffff19600454161760045581600555612d2660405192839283602090939291936001600160a01b0360408201951681520152565b0390a160a05180f35b3461020b57612d3d36613fe6565b612d48929192614950565b67ffffffffffffffff8216612d6a816000526007602052604060002054151590565b15612d85575061094292612d7f913691613edd565b90614d29565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461020b57604060031936011261020b5760043561ffff81169081900361020b5760243567ffffffffffffffff811161020b577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e91612e54612e1c6020933690600401613fb5565b90612e25614950565b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b556149dd565b604051908152a160a05180f35b3461020b57612e89612e91612e7536613f63565b9491612e82939193614950565b3691614338565b923691614338565b7f000000000000000000000000000000000000000000000000000000000000000015612f995760a0515b8251811015612f2157806001600160a01b03612ed9600193866144b7565b5116612ee4816157f1565b612ef0575b5001612ebb565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612ee9565b5060a0515b815181101561094257806001600160a01b03612f44600193856144b7565b51168015612f9357612f5581615a47565b612f62575b505b01612f26565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612f5a565b50612f5c565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461020b57604060031936011261020b57612fde613e67565b60243567ffffffffffffffff811161020b57602091613004612af7923690600401613f14565b906142e3565b3461020b57604060031936011261020b5760043567ffffffffffffffff811161020b578060040190610100600319823603011261020b57613049613e8f565b60405190919061305881613d20565b60a051905261308961307f61307a61307360c4850187614269565b3691613edd565b61476e565b6064830135614849565b9160848201613097816142ba565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361217f5750602482019377ffffffffffffffff000000000000000000000000000000006130f1866142ce565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119c35760a051916134f1575b506121245767ffffffffffffffff61317a866142ce565b16613192816000526007602052604060002054151590565b156120a25760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156119c35760a051916134d2575b5015612029576131fe856142ce565b9061321460a48501926130046130738585614269565b1561348b57506044929161ffff161590506133ed5767ffffffffffffffff61323b856142ce565b168060a05152600d6020527f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61584806132a4604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615b56565b604080516001600160a01b039290921682526020820192909252a25b01916132cb836142ba565b906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561020b576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390921660048201526024810185905290818060448101038160a051875af180156119c3576133d4575b50608067ffffffffffffffff6020956001600160a01b036133a261339c7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966142ce565b926142ba565b60405196875233898801521660408601528560608601521692a2604051906133c982613d20565b815260405190518152f35b60a0516133e091613d90565b60a05161020b5784613358565b67ffffffffffffffff6133ff856142ce565b168060a0515260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c848061346b6002604060a05120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391615b56565b604080516001600160a01b039290921682526020820192909252a26132c0565b6134959250614269565b6122b96040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916145c7565b6134eb915060203d6020116121785761216a8183613d90565b866131ef565b61350a915060203d6020116121785761216a8183613d90565b86613163565b3461020b57602060031936011261020b5760043567ffffffffffffffff811161020b578060040190610100600319823603011261020b5760405161355381613d20565b60a051905260405161356481613d20565b60a051905261357f61307f61307a61307360c4850186614269565b906084810161358d816142ba565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811691160361217f5750602481019277ffffffffffffffff000000000000000000000000000000006135e7856142ce565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156119c35760a0519161388f575b506121245767ffffffffffffffff613670856142ce565b16613688816000526007602052604060002054151590565b156120a25760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa9081156119c35760a05191613870575b5015612029576136f4846142ce565b9061370a60a48401926130046130738585614269565b1561348b575082916044915067ffffffffffffffff613728866142ce565b168060a0515260086020526137716002604060a05120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958691615b56565b604080516001600160a01b0386168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a201926137b7846142ba565b823b1561020b576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390921660048201526024810185905290818060448101038160a051875af180156119c3576020956001600160a01b036133a261339c7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09660809667ffffffffffffffff9661385e575b506142ce565b60a05161386a91613d90565b8b613858565b613889915060203d6020116121785761216a8183613d90565b856136e5565b6138a8915060203d6020116121785761216a8183613d90565b85613659565b3461020b5760a060031936011261020b576138c7613e10565b506138d0613e50565b60443567ffffffffffffffff811161020b5760031960a0913603011261020b576138f8613e7e565b506084359067ffffffffffffffff821161020b5761392367ffffffffffffffff923690600401613eaf565b505060405161393181613cee565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260c060a0519101521660a05152600f60205260e0604060a051206040519061398582613cee565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b3461020b5760c060031936011261020b57613a35613e10565b50613a3e613e50565b613a46613e26565b5060843561ffff8116810361020b5760a4359067ffffffffffffffff821161020b5763ffffffff61ffff613a8d829386613a8660a0973690600401613eaf565b505061412d565b95926040979194975197885216602087015216604085015216606083015215156080820152f35b3461020b5760a05160031936011261020b57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020b57602060031936011261020b576020613b0e613e10565b6040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081169216919091148152f35b3461020b5760a05160031936011261020b5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b3461020b5760a05160031936011261020b5760408051611b8c91613bb09082613d90565b601f81527f4275726e46726f6d4d696e74546f6b656e506f6f6c20312e372e302d646576006020820152604051918291602083526020830190613dcf565b3461020b57602060031936011261020b57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361020b57817faff2afbf0000000000000000000000000000000000000000000000000000000060209314908115613cc4575b8115613c9a575b8115613c70575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613c69565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613c62565b7fdc0cbd360000000000000000000000000000000000000000000000000000000081149150613c5b565b60e0810190811067ffffffffffffffff821117613d0a57604052565b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613d0a57604052565b6040810190811067ffffffffffffffff821117613d0a57604052565b60a0810190811067ffffffffffffffff821117613d0a57604052565b6060810190811067ffffffffffffffff821117613d0a57604052565b90601f601f19910116810190811067ffffffffffffffff821117613d0a57604052565b67ffffffffffffffff8111613d0a57601f01601f191660200190565b919082519283825260005b848110613dfb575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201613dda565b600435906001600160a01b038216820361093d57565b606435906001600160a01b038216820361093d57565b35906001600160a01b038216820361093d57565b6024359067ffffffffffffffff8216820361093d57565b6004359067ffffffffffffffff8216820361093d57565b6064359061ffff8216820361093d57565b6024359061ffff8216820361093d57565b359061ffff8216820361093d57565b9181601f8401121561093d5782359167ffffffffffffffff831161093d576020838186019501011161093d57565b929192613ee982613db3565b91613ef76040519384613d90565b82948184528183011161093d578281602093846000960137010152565b9080601f8301121561093d57816020613f2f93359101613edd565b90565b9181601f8401121561093d5782359167ffffffffffffffff831161093d576020808501948460051b01011161093d57565b604060031982011261093d5760043567ffffffffffffffff811161093d5781613f8e91600401613f32565b929092916024359067ffffffffffffffff821161093d57613fb191600401613f32565b9091565b9181601f8401121561093d5782359167ffffffffffffffff831161093d5760208085019460e0850201011161093d57565b90604060031983011261093d5760043567ffffffffffffffff8116810361093d57916024359067ffffffffffffffff821161093d57613fb191600401613eaf565b602060408183019282815284518094520192019060005b81811061404b5750505090565b82516001600160a01b031684526020938401939092019160010161403e565b9181601f8401121561093d5782359167ffffffffffffffff831161093d576020808501946060850201011161093d57565b613f2f9160206140b48351604084526040840190613dcf565b920151906020818403910152613dcf565b3590811515820361093d57565b35906001600160801b038216820361093d57565b919082606091031261093d576040516140fe81613d74565b604061412881839561410f816140c5565b855261411d602082016140d2565b6020860152016140d2565b910152565b67ffffffffffffffff16600052600f60205260406000206040519061415182613cee565b549263ffffffff84168252602082019363ffffffff8160201c168552604083019063ffffffff8160401c1682526060840163ffffffff8260601c168152608085019561ffff8360801c16875260ff60a087019361ffff8160901c16855260a01c1615801560c08801526142505761ffff16806141ea5750505063ffffffff808061ffff9351169451169551169351169193929190600190565b919550915061ffff600b54169081811061422057505063ffffffff808061ffff9351169451169551169351169193929190600190565b7f7911d95b0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050505092505050600090600090600090600090600090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561093d570180359067ffffffffffffffff821161093d5760200191813603831361093d57565b356001600160a01b038116810361093d5790565b3567ffffffffffffffff8116810361093d5790565b9067ffffffffffffffff613f2f92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613d0a5760051b60200190565b92919061434481614320565b936143526040519586613d90565b602085838152019160051b810192831161093d57905b82821061437457505050565b6020809161438184613e3c565b815201910190614368565b67ffffffffffffffff16600052600e60205260406000209160028110156143d3576001146143c257816001613f2f930190614f9a565b8160026003613f2f94019101614f9a565b634e487b7160e01b600052602160045260246000fd5b91908110156143f95760051b0190565b634e487b7160e01b600052603260045260246000fd5b91908110156143f9576060020190565b6040519061442c82613d58565b60006080838281528260208201528260408201528260608201520152565b9060405161445781613d58565b60806001829460ff81546001600160801b038116865263ffffffff81861c16602087015260a01c161515604085015201546001600160801b0381166060840152811c910152565b604051906144ab82613d3c565b60606020838281520152565b80518210156143f95760209160051b010190565b90600182811c921680156144fb575b60208310146144e557565b634e487b7160e01b600052602260045260246000fd5b91607f16916144da565b9060405191826000825492614519846144cb565b80845293600181169081156145875750600114614540575b5061453e92500383613d90565b565b90506000929192526020600020906000915b81831061456b57505090602061453e9282010138614531565b6020919350806001915483858901015201910190918492614552565b6020935061453e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614531565b601f8260209493601f19938186528686013760008582860101520116010190565b67ffffffffffffffff166000526008602052613f2f6004604060002001614505565b91908110156143f95760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff618136030182121561093d570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561093d570180359067ffffffffffffffff821161093d57602001918160051b3603831361093d57565b818102929181159184041417156146b157565b634e487b7160e01b600052601160045260246000fd5b8181106146d2575050565b600081556001016146c7565b9160209082815201919060005b8181106146f85750505090565b9091926020806001926001600160a01b0361471288613e3c565b1681520194019291016146eb565b91908110156143f95760081b0190565b35801515810361093d5790565b3561ffff8116810361093d5790565b3563ffffffff8116810361093d5790565b359063ffffffff8216820361093d57565b805180156147de576020036147a057805160208281019183018390031261093d57519060ff82116147a0575060ff1690565b6122b9906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613dcf565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116146b157565b60ff16604d81116146b157600a0a90565b8115614833570490565b634e487b7160e01b600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff81169282841461493157828411614907579061488e91614804565b91604d60ff84161180156148ec575b6148b6575050906148b0613f2f92614818565b9061469e565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b506148f683614818565b80156148335760001904841161489d565b61491091614804565b91604d60ff8416116148b65750509061492b613f2f92614818565b90614829565b5050505090565b9081602091031261093d5751801515810361093d5790565b6001600160a01b0360015416330361496457565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b356001600160801b038116810361093d5790565b6001600160801b036149d7604080936149ba816140c5565b15158652836149cb602083016140d2565b166020870152016140d2565b16910152565b9160005b82811015614d2357600060e0820285016149fa816142ce565b9067ffffffffffffffff821692614a1e846000526007602052604060002054151590565b15614cf7575060019392614b577f20ae59542ddd78610f62f9d2c9dcd658f8b6a5b45a0f03e71e16614e89dda83693614b4d84614b3d602060e0970191614a6d614a6836856140e6565b615479565b614ad0614a8e8667ffffffffffffffff16600052600c602052604060002090565b805463ffffffff8160801c16159081614ce2575b81614cd3575b81614cc1575b81614cb2575b5080614ca3575b614c2b575b614aca36866140e6565b9061558a565b614aff6080820195614ae5614a6836896140e6565b67ffffffffffffffff16600052600d602052604060002090565b90815463ffffffff8160801c16159081614c16575b81614c07575b81614bf5575b81614be6575b5080614bd7575b614b5e575b50614aca36866140e6565b60405194855260208501906149a2565b60808301906149a2565ba1016149e1565b614b7260a06001600160801b03920161498e565b825473ffffffff000000000000000000000000000000004260801b1673ffffffffffffffffffffffffffffffffffffffff19909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717815538614b32565b50614be186614730565b614b2d565b60ff915060a01c161538614b26565b6001600160801b038116159150614b20565b838e015460801c159150614b1a565b838e01546001600160801b0316159150614b14565b6001600160801b03614c3f6040850161498e565b825473ffffffff000000000000000000000000000000004260801b1673ffffffffffffffffffffffffffffffffffffffff19909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178155614ac0565b50614cad85614730565b614abb565b60ff915060a01c161538614ab4565b6001600160801b038116159150614aae565b828f015460801c159150614aa8565b828f01546001600160801b0316159150614aa2565b80847f1e670e4b0000000000000000000000000000000000000000000000000000000060249352600452fd5b50915050565b90805115614f0f5767ffffffffffffffff81516020830120921691826000526008602052614d5e816005604060002001615b01565b15614ecb5760005260096020526040600020815167ffffffffffffffff8111613d0a57614d8b82546144cb565b601f8111614e99575b506020601f8211600114614e0f5791614de9827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614dff95600091614e04575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613dcf565b0390a2565b905084015138614dd6565b601f1982169083600052806000209160005b818110614e81575092614dff9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614e68575b5050811b019055611b78565b85015160001960f88460031b161c191690553880614e5c565b9192602060018192868a015181550194019201614e21565b614ec590836000526020600020601f840160051c8101916020851061088557601f0160051c01906146c7565b38614d94565b50906122b96040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613dcf565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b818110614f6b57505061453e92500383613d90565b84546001600160a01b0316835260019485019487945060209093019201614f56565b919082018092116146b157565b614fa390614f39565b916005548015159182615083575b5050614fbb575090565b614fc490614f39565b90815180614fd25750905090565b614fdd908251614f8d565b92601f19615003614fed86614320565b95614ffb6040519788613d90565b808752614320565b0136602086013760005b825181101561503e57806001600160a01b0361502b600193866144b7565b511661503782886144b7565b520161500d565b509160005b815181101561507e57806001600160a01b03615061600193856144b7565b5116615077615071838751614f8d565b886144b7565b5201615043565b505050565b101590503880614fb1565b67ffffffffffffffff16600081815260076020526040902054909291901561517e579161517b60e092615150856150e57f73d6dce40db73cbddae4b9ce52576043a1644e08c2702236273d71077435fa1697615479565b8460005260086020526150fc81604060002061558a565b61510583615479565b84600052600860205261511f83600260406000200161558a565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b919082039182116146b157565b6151c161441f565b506001600160801b036060820151166001600160801b03808351169161520c60208501936152066151f963ffffffff875116426151ac565b856080890151169061469e565b90614f8d565b8082101561522557505b16825263ffffffff4216905290565b9050615216565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561093d576040517f79cc6790000000000000000000000000000000000000000000000000000000008152306004820152602481019190915260009182908290604490829084905af180156152bd576152b0575050565b816152ba91613d90565b50565b6040513d84823e3d90fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613f2f604082613d90565b9061ffff9067ffffffffffffffff61531d602085016142ce565b16600052600f6020526040600020826040519161533983613cee565b549263ffffffff8416835263ffffffff8460201c16602084015263ffffffff8460401c16604084015263ffffffff8460601c166060840152818460801c169283608082015260c060ff848760901c16968760a085015260a01c1615159101521615156000146153cf57505b1680156153c7576127106153c06060613f2f940135928361469e565b04906151ac565b506060013590565b90506153a4565b805160005b8181106153e757505050565b600181018082116146b1575b82811061540357506001016153db565b6001600160a01b0361541583866144b7565b51166001600160a01b0361542983876144b7565b511614615438576001016153f3565b6001600160a01b0361544a83866144b7565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8051156154fe576001600160801b036040820151166001600160801b03602083015116106154a45750565b6064906154fc604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565bfd5b6001600160801b0360408201511615801590615574575b61551c5750565b6064906154fc604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565b506001600160801b036020820151161515615515565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916156b160609280546155c763ffffffff8260801c16426151ac565b90816156e7575b50506001600160801b0360018160208601511692828154168085106000146156df57508280855b16167fffffffffffffffffffffffffffffffff0000000000000000000000000000000082541617815561566e8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166001600160801b031692909217910155565b61517b60405180926001600160801b0360408092805115158552826020820151166020860152015116910152565b8380916155f5565b6001600160801b039161571383928361570c6001880154948286169560801c9061469e565b9116614f8d565b8082101561578757505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff92909116929092161673ffffffffffffffffffffffffffffffffffffffff19909116174260801b73ffffffff000000000000000000000000000000001617815538806155ce565b905061571d565b906040519182815491828252602082019060005260206000209260005b8181106157c057505061453e92500383613d90565b84548352600194850194879450602090930192016157ab565b80548210156143f95760005260206000200190600090565b60008181526003602052604090205480156158d15760001981018181116146b1576002549060001982019182116146b157818103615880575b505050600254801561586a57600019016158458160026157d9565b60001982549160031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6158b96158916158a29360026157d9565b90549060031b1c92839260026157d9565b81939154906000199060031b92831b921b19161790565b9055600052600360205260406000205538808061582a565b5050600090565b60008181526007602052604090205480156158d15760001981018181116146b1576006549060001982019182116146b157818103615951575b505050600654801561586a576000190161592c8160066157d9565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b6159736159626158a29360066157d9565b90549060031b1c92839260066157d9565b90556000526007602052604060002055388080615911565b9060018201918160005282602052604060002054801515600014615a3e5760001981018181116146b15782549060001982019182116146b157818103615a07575b5050508054801561586a5760001901906159e682826157d9565b60001982549160031b1b191690555560005260205260006040812055600190565b615a27615a176158a293866157d9565b90549060031b1c928392866157d9565b9055600052836020526040600020553880806159cc565b50505050600090565b80600052600360205260406000205415600014615aa15760025468010000000000000000811015613d0a57615a886158a282600185940160025560026157d9565b9055600254906000526003602052604060002055600190565b50600090565b80600052600760205260406000205415600014615aa15760065468010000000000000000811015613d0a57615ae86158a282600185940160065560066157d9565b9055600654906000526007602052604060002055600190565b60008281526001820160205260409020546158d15780549068010000000000000000821015613d0a5782615b3f6158a28460018096018555846157d9565b905580549260005201602052604060002055600190565b9182549060ff8260a01c16158015615d42575b615d3c576001600160801b0382169160018501908154615b9c63ffffffff6001600160801b0383169360801c16426151ac565b9081615c9e575b5050848110615c5f5750838310615bf4575050615bc96001600160801b039283926151ac565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91615c0381856151ac565b926000198101908082116146b157615c26615c2b926001600160a01b0396614f8d565b614829565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611615d1257615cb9926152069160801c9061469e565b80841015615d0d5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615ba3565b615cc4565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215615b69565b91929015615dc55750815115615d5e575090565b3b15615d675790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615dd85750805190602001fd5b6122b9906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613dcf56fea164736f6c634300081a000a",
}

var BurnFromMintTokenPoolABI = BurnFromMintTokenPoolMetaData.ABI

var BurnFromMintTokenPoolBin = BurnFromMintTokenPoolMetaData.Bin

func DeployBurnFromMintTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnFromMintTokenPool, error) {
	parsed, err := BurnFromMintTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnFromMintTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router)
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetAllowList(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetAllowList(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _BurnFromMintTokenPool.Contract.GetAllowListEnabled(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnFromMintTokenPool.Contract.GetAllowListEnabled(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getCurrentCustomBlockConfirmationRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentCustomBlockConfirmationRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _BurnFromMintTokenPool.Contract.GetCurrentRateLimiterState(&_BurnFromMintTokenPool.CallOpts, remoteChainSelector)
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
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetMinBlockConfirmation() (uint16, error) {
	return _BurnFromMintTokenPool.Contract.GetMinBlockConfirmation(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetMinBlockConfirmation() (uint16, error) {
	return _BurnFromMintTokenPool.Contract.GetMinBlockConfirmation(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRateLimitAdmin(&_BurnFromMintTokenPool.CallOpts)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRateLimitAdmin(&_BurnFromMintTokenPool.CallOpts)
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRequiredCCVs(&_BurnFromMintTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _BurnFromMintTokenPool.Contract.GetRequiredCCVs(&_BurnFromMintTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _BurnFromMintTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnFromMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnFromMintTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnFromMintTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnFromMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyAllowListUpdates(&_BurnFromMintTokenPool.TransactOpts, removes, adds)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnFromMintTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnFromMintTokenPool.TransactOpts, ccvConfigArgs)
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "applyCustomBlockConfirmationConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_BurnFromMintTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_BurnFromMintTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.LockOrBurn0(&_BurnFromMintTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, arg2)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, blockConfirmationRequested uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.LockOrBurn0(&_BurnFromMintTokenPool.TransactOpts, lockOrBurnIn, blockConfirmationRequested, arg2)
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetChainRateLimiterConfig(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnFromMintTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setCustomBlockConfirmationRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_BurnFromMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_BurnFromMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetDynamicConfig(&_BurnFromMintTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetDynamicConfig(&_BurnFromMintTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetRateLimitAdmin(&_BurnFromMintTokenPool.TransactOpts, rateLimitAdmin)
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _BurnFromMintTokenPool.Contract.SetRateLimitAdmin(&_BurnFromMintTokenPool.TransactOpts, rateLimitAdmin)
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

type BurnFromMintTokenPoolAllowListAddIterator struct {
	Event *BurnFromMintTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolAllowListAdd)
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
		it.Event = new(BurnFromMintTokenPoolAllowListAdd)
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

func (it *BurnFromMintTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnFromMintTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolAllowListAddIterator{contract: _BurnFromMintTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolAllowListAdd)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*BurnFromMintTokenPoolAllowListAdd, error) {
	event := new(BurnFromMintTokenPoolAllowListAdd)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolAllowListRemoveIterator struct {
	Event *BurnFromMintTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolAllowListRemove)
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
		it.Event = new(BurnFromMintTokenPoolAllowListRemove)
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

func (it *BurnFromMintTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnFromMintTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolAllowListRemoveIterator{contract: _BurnFromMintTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolAllowListRemove)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*BurnFromMintTokenPoolAllowListRemove, error) {
	event := new(BurnFromMintTokenPoolAllowListRemove)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolCCVConfigUpdatedIterator struct {
	Event *BurnFromMintTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolCCVConfigUpdated)
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
		it.Event = new(BurnFromMintTokenPoolCCVConfigUpdated)
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

func (it *BurnFromMintTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolCCVConfigUpdatedIterator{contract: _BurnFromMintTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolCCVConfigUpdated)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*BurnFromMintTokenPoolCCVConfigUpdated, error) {
	event := new(BurnFromMintTokenPoolCCVConfigUpdated)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

type BurnFromMintTokenPoolConfigChangedIterator struct {
	Event *BurnFromMintTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolConfigChanged)
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
		it.Event = new(BurnFromMintTokenPoolConfigChanged)
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

func (it *BurnFromMintTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnFromMintTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolConfigChangedIterator{contract: _BurnFromMintTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolConfigChanged)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseConfigChanged(log types.Log) (*BurnFromMintTokenPoolConfigChanged, error) {
	event := new(BurnFromMintTokenPoolConfigChanged)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

type BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator struct {
	Event *BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured)
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
		it.Event = new(BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured)
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

func (it *BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterCustomBlockConfirmationRateLimitConfigured(opts *bind.FilterOpts) (*BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator{contract: _BurnFromMintTokenPool.contract, event: "CustomBlockConfirmationRateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchCustomBlockConfirmationRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationRateLimitConfigured", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseCustomBlockConfirmationRateLimitConfigured(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured, error) {
	event := new(BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationRateLimitConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolCustomBlockConfirmationUpdatedIterator struct {
	Event *BurnFromMintTokenPoolCustomBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolCustomBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolCustomBlockConfirmationUpdated)
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
		it.Event = new(BurnFromMintTokenPoolCustomBlockConfirmationUpdated)
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

func (it *BurnFromMintTokenPoolCustomBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolCustomBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolCustomBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*BurnFromMintTokenPoolCustomBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolCustomBlockConfirmationUpdatedIterator{contract: _BurnFromMintTokenPool.contract, event: "CustomBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolCustomBlockConfirmationUpdated)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseCustomBlockConfirmationUpdated(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationUpdated, error) {
	event := new(BurnFromMintTokenPoolCustomBlockConfirmationUpdated)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnFromMintTokenPoolDefaultFinalityRateLimitConfiguredIterator struct {
	Event *BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolDefaultFinalityRateLimitConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured)
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
		it.Event = new(BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured)
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

func (it *BurnFromMintTokenPoolDefaultFinalityRateLimitConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolDefaultFinalityRateLimitConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterDefaultFinalityRateLimitConfigured(opts *bind.FilterOpts) (*BurnFromMintTokenPoolDefaultFinalityRateLimitConfiguredIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "DefaultFinalityRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolDefaultFinalityRateLimitConfiguredIterator{contract: _BurnFromMintTokenPool.contract, event: "DefaultFinalityRateLimitConfigured", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchDefaultFinalityRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "DefaultFinalityRateLimitConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "DefaultFinalityRateLimitConfigured", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseDefaultFinalityRateLimitConfigured(log types.Log) (*BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured, error) {
	event := new(BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "DefaultFinalityRateLimitConfigured", log); err != nil {
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
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
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

type BurnFromMintTokenPoolRateLimitAdminSetIterator struct {
	Event *BurnFromMintTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnFromMintTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnFromMintTokenPoolRateLimitAdminSet)
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
		it.Event = new(BurnFromMintTokenPoolRateLimitAdminSet)
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

func (it *BurnFromMintTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *BurnFromMintTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnFromMintTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnFromMintTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &BurnFromMintTokenPoolRateLimitAdminSetIterator{contract: _BurnFromMintTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _BurnFromMintTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnFromMintTokenPoolRateLimitAdminSet)
				if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_BurnFromMintTokenPool *BurnFromMintTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*BurnFromMintTokenPoolRateLimitAdminSet, error) {
	event := new(BurnFromMintTokenPoolRateLimitAdminSet)
	if err := _BurnFromMintTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (BurnFromMintTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnFromMintTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnFromMintTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (BurnFromMintTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnFromMintTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnFromMintTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x20ae59542ddd78610f62f9d2c9dcd658f8b6a5b45a0f03e71e16614e89dda836")
}

func (BurnFromMintTokenPoolCustomBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e")
}

func (BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured) Topic() common.Hash {
	return common.HexToHash("0x73d6dce40db73cbddae4b9ce52576043a1644e08c2702236273d71077435fa16")
}

func (BurnFromMintTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
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

func (BurnFromMintTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
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

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnFromMintTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnFromMintTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnFromMintTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnFromMintTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*BurnFromMintTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnFromMintTokenPoolChainAdded, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnFromMintTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnFromMintTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnFromMintTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnFromMintTokenPoolConfigChanged, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationRateLimitConfigured(opts *bind.FilterOpts) (*BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfiguredIterator, error)

	WatchCustomBlockConfirmationRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured) (event.Subscription, error)

	ParseCustomBlockConfirmationRateLimitConfigured(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationRateLimitConfigured, error)

	FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*BurnFromMintTokenPoolCustomBlockConfirmationUpdatedIterator, error)

	WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomBlockConfirmationUpdated(log types.Log) (*BurnFromMintTokenPoolCustomBlockConfirmationUpdated, error)

	FilterDefaultFinalityRateLimitConfigured(opts *bind.FilterOpts) (*BurnFromMintTokenPoolDefaultFinalityRateLimitConfiguredIterator, error)

	WatchDefaultFinalityRateLimitConfigured(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured) (event.Subscription, error)

	ParseDefaultFinalityRateLimitConfigured(log types.Log) (*BurnFromMintTokenPoolDefaultFinalityRateLimitConfigured, error)

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

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*BurnFromMintTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *BurnFromMintTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*BurnFromMintTokenPoolRateLimitAdminSet, error)

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
