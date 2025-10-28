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

var BurnToAddressMintTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"burnAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"i_burnAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmationConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61012080604052346103dd57616603803803809161001d828561045c565b8339810160c0828203126103dd5781516001600160a01b03811692908390036103dd5761004c6020820161047f565b60408201516001600160401b0381116103dd5782019280601f850112156103dd578351936001600160401b038511610446578460051b906020820195610095604051978861045c565b86526020808701928201019283116103dd57602001905b82821061042e575050506100c26060830161048d565b936100db60a06100d46080860161048d565b940161048d565b94331561041d57600180546001600160a01b031916331790558115801561040c575b80156103fb575b6103ea578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103a9575b5061037e575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610261575b5061010052604051615fc1908161064282396080518181816118c001528181611b2f01528181611ca90152818161223401528181612486015281816125c901528181612edf0152818161310101528181613269015281816133bd0152818161356f015281816138a1015281816138f6015281816139e401526142d6015260a05181818161385d015281816149c901528181614a4c015261532f015260c051818181610bda0152818161194e015281816122c201528181612f6d015261344d015260e051818181610b9401528181611993015281816123070152612ccf015261010051818181611b54015281816124ac0152613e170152f35b6020604051610270828261045c565b60008152600036813760e0511561036d5760005b81518110156102eb576001906001600160a01b036102a282856104a1565b5116846102ae826104e3565b6102bb575b505001610284565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138846102b3565b505060005b8251811015610364576001906001600160a01b0361030e82866104a1565b5116801561035e5783610320826105e1565b61032e575b50505b016102f0565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13883610325565b50610328565b50505038610169565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036103925750610137565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103e2575b816103c56020938361045c565b810103126103dd576103d69061047f565b9038610131565b600080fd5b3d91506103b8565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610104565b506001600160a01b038416156100fd565b639b15e16f60e01b60005260046000fd5b6020809161043b8461048d565b8152019101906100ac565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761044657604052565b519060ff821682036103dd57565b51906001600160a01b03821682036103dd57565b80518210156104b55760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156104b55760005260206000200190600090565b60008181526003602052604090205480156105da5760001981018181116105c4576002546000198101919082116105c457818103610573575b505050600254801561055d57600019016105378160026104cb565b8154906000199060031b1b19169055600255600052600360205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6105ac6105846105959360026104cb565b90549060031b1c92839260026104cb565b819391549060031b91821b91600019901b19161790565b9055600052600360205260406000205538808061051c565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260036020526040600020541560001461063b57600254680100000000000000008110156104465761062261059582600185940160025560026104cb565b9055600254906000526003602052604060002055600190565b5060009056fe60c080604052600436101561001357600080fd5b600060a05260a0513560e01c90816301ffc9a714613aa8575080630574d6a614613a18578063164e68de1461397c578063181f5a771461391a57806321df0da7146138d5578063240028e81461388157806324f65ee71461384257806337b19247146136d457806338b39d29146115d75780633907753714613336578063489a68f214612e465780634c5ef0ed14612e0157806354c8a4f314612c9d57806359152aad14612bf05780635df45a3714612bcc57806362ddd3c414612b47578063698c2c6614612a955780636d3d1a5814612a6d5780637437ff9f14612a3957806379ba50971461297c5780637d54534e146128f85780638926f54f146128b357806389720a62146128485780638da5cb5b14612820578063962d4020146126dd5780639a4575b9146121e2578063a42a7b8b1461208b578063a7cd63b71461201d578063acfecf9114611ef0578063af58d59f14611ea4578063b1c71c6514611843578063b79465801461180b578063bb6bb5a7146117a3578063c4bffe2b14611687578063c75eea9c146115dc578063c8de9fe0146115d7578063cf7401f31461144a578063d966866b14610ffd578063dc04fa1f14610bfe578063dc0bd97114610bb9578063e0351e1314610b7b578063e8a1da17146102c35763f2fde38b146101fe57600080fd5b346102bd5760206003193601126102bd576001600160a01b0361021f613c33565b61022761479d565b163381146102915760a05180547fffffffffffffffffffffffff000000000000000000000000000000000000000016821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360a05180f35b7fdad89dca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b60a05180fd5b346102bd576102d136613ec1565b9190926102dc61479d565b60a051905b8282106109b95750505060a0519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b818110156109b357600581901b830135858112156102bd578301610120813603126102bd576040519461035086613d3f565b813567ffffffffffffffff811681036109ae578652602082013567ffffffffffffffff81116102bd5782019436601f870112156102bd5785356103928161422e565b966103a06040519889613d77565b81885260208089019260051b820101903682116102bd5760208101925b82841061097f575050505060208701958652604083013567ffffffffffffffff81116102bd576103f09036908501613e72565b916040880192835261041a610408366060870161404d565b9460608a0195865260c036910161404d565b956080890196875261042c855161557e565b610436875161557e565b835151156109535761045267ffffffffffffffff8a5116615e39565b156109185767ffffffffffffffff89511660a051526008602052604060a0512061059686516fffffffffffffffffffffffffffffffff604082015116906105516fffffffffffffffffffffffffffffffff602083015116915115158360806040516104bc81613d3f565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106bc88516fffffffffffffffffffffffffffffffff604082015116906106776fffffffffffffffffffffffffffffffff602083015116915115158360806040516105e081613d3f565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff82116108e7576106df8354614437565b601f81116108a8575b506020906001601f84111461083e5791809161071b9360a05192610833575b50506000198260011b9260031b1c19161790565b90555b60a0515b88518051821015610757579061075160019261074a8367ffffffffffffffff8f511692614423565b5190614e93565b01610722565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561082567ffffffffffffffff60019796949851169251935191516107f16107bc60405196879687526101006020880152610100870190613db6565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101939290919361031e565b015190508e80610707565b90601f198316918460a051528160a051209260a0515b8181106108905750908460019594939210610877575b505050811b01905561071e565b015160001960f88460031b161c191690558d808061086a565b92936020600181928786015181550195019301610854565b6108d7908460a05152602060a05120601f850160051c810191602086106108dd575b601f0160051c01906146f6565b8d6106e8565b90915081906108ca565b7f4e487b710000000000000000000000000000000000000000000000000000000060a051526041600452602460a051fd5b67ffffffffffffffff8951167f1d5ad3c50000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b833567ffffffffffffffff81116102bd576020916109a38392833691870101613e72565b8152019301926103bd565b600080fd5b60a05180f35b9092919367ffffffffffffffff6109d96109d48688866143bb565b6141dc565b16926109e484615c6a565b15610b4b578360a051526008602052610a046005604060a0512001615b07565b9260a0515b8451811015610a43576001908660a051526008602052610a3c6005604060a0512001610a358389614423565b5190615d1d565b5001610a09565b50939094919592508060a0515260086020526005604060a0512060a051815560a051600182015560a051600282015560a051600382015560048101610a888154614437565b80610afb575b50500180549060a051815581610ad7575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916102e1565b60a05152602060a05120908101905b81811015610a9f5760a0518155600101610ae6565b601f8111600114610b14575060a05190555b8880610a8e565b610b34908260a051526001601f602060a051209201861c820191016146f6565b60a080518290525160208120918190559055610b0d565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102bd5760a0516003193601126102bd5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346102bd5760a0516003193601126102bd5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bd5760406003193601126102bd5760043567ffffffffffffffff81116102bd57366023820112156102bd57806004013567ffffffffffffffff81116102bd576024820191602436918360081b0101116102bd5760243567ffffffffffffffff81116102bd57610c74903690600401613e90565b919092610c7f61479d565b60a0515b828110610cfb5750505060a0515b818110610c9e5760a05180f35b8067ffffffffffffffff610cb86109d460019486886143bb565b168060a05152600f60205260a051604060a05120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee860a05160a051a201610c91565b610d096109d482858561474f565b610d1482858561474f565b602081019060a081019261271061ffff610d2d8661475f565b161015610ff15760c082019061271061ffff610d488461475f565b161015610fb55767ffffffffffffffff16938460a05152600f60205260a0516040902092610d758561476e565b63ffffffff168454916040810190610d8c8261476e565b60201b67ffffffff0000000016936060820193610da88561476e565b60401b6bffffffff000000000000000016956080840196610dc88861476e565b60601b6fffffffff0000000000000000000000001691610de78a61475f565b60801b71ffff000000000000000000000000000000001693610e088c61475f565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717875560e00195610ebf8761477f565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff00000000000000000000000000000000000000001617905560405196610f109061478c565b63ffffffff168752610f219061478c565b63ffffffff166020870152610f359061478c565b63ffffffff166040860152610f499061478c565b63ffffffff166060850152610f5d90613c7f565b61ffff166080840152610f6f90613c7f565b61ffff1660a0830152610f8190614023565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a2600101610c83565b61ffff610fc18361475f565b7f95f3517a0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b61ffff610fc18561475f565b346102bd5760206003193601126102bd5760043567ffffffffffffffff81116102bd5761102e903690600401613e90565b9061103761479d565b60a0516080525b816080511061104d5760a05180f35b61105d6109d46080518484614620565b61107761106d6080518585614620565b6020810190614660565b6110916110876080518787614620565b6040810190614660565b906110ac6110a26080518989614620565b6060810190614660565b6110c66110bc6080518b8b614620565b6080810190614660565b9390946110dc6110d736898b614246565b6154db565b6110ea6110d7368385614246565b6110f86110d7368587614246565b6111066110d7368789614246565b6040519860808a01908a821067ffffffffffffffff8311176108e75767ffffffffffffffff9160405261113a368a8c614246565b8b52611147368486614246565b60208c0152611157368688614246565b60408c015261116736888a614246565b60608c015216988960a05152600e602052604060a05120815180519067ffffffffffffffff82116108e7576801000000000000000082116108e757602090835483855580841061142b575b50018260a05152602060a0512060a0515b83811061140e57505050506001810160208301519081519167ffffffffffffffff83116108e7576801000000000000000083116108e75760209082548484558085106113ef575b50019060a05152602060a0512060a0515b8381106113d257505050506002810160408301519081519167ffffffffffffffff83116108e7576801000000000000000083116108e75760209082548484558085106113b3575b50019060a05152602060a0512060a0515b838110611396575050505060036060910191015180519067ffffffffffffffff82116108e7576801000000000000000082116108e7576020908354838555808410611377575b50019160a05152602060a051209160a0515b82811061135a575050505092611349949261132d61133b937fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a99989661131f6040519b8c9b60808d5260808d019161470d565b918a830360208c015261470d565b91878303604089015261470d565b91848303606086015261470d565b0390a260016080510160805261103e565b60019060206001600160a01b0384511693019281860155016112cb565b611390908560a05152848460a0512091820191016146f6565b8f6112b9565b60019060206001600160a01b038551169401938184015501611273565b6113cc908460a05152858460a0512091820191016146f6565b38611262565b60019060206001600160a01b03855116940193818401550161121b565b611408908460a05152858460a0512091820191016146f6565b3861120a565b60019060206001600160a01b0385511694019381840155016111c3565b611444908560a05152848460a0512091820191016146f6565b386111b2565b346102bd5760e06003193601126102bd57611463613bd9565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01126102bd5760405161149981613d5b565b60243580151581036102bd5781526044356fffffffffffffffffffffffffffffffff811681036102bd5760208201526064356fffffffffffffffffffffffffffffffff811681036102bd5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c01126102bd576040519061152082613d5b565b60843580151581036102bd57825260a4356fffffffffffffffffffffffffffffffff811681036102bd57602083015260c4356fffffffffffffffffffffffffffffffff811681036102bd5760408301526001600160a01b03600a5416331415806115c2575b611592576109b3926151f8565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506001600160a01b0360015416331415611585565b613df7565b346102bd5760206003193601126102bd5767ffffffffffffffff6115fe613bd9565b61160661456d565b501660a05152600860205261168361162a611625604060a05120614598565b615370565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346102bd5760a0516003193601126102bd5760a051506040516006548082528160208101600660a05152602060a051209260a0515b81811061178a5750506116d192500382613d77565b8051906116f66116e08361422e565b926116ee6040519485613d77565b80845261422e565b90601f1960208401920136833760a0515b8151811015611739578067ffffffffffffffff61172660019385614423565b51166117328287614423565b5201611707565b5050906040519182916020830190602084525180915260408301919060a0515b818110611767575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611759565b84548352600194850194869450602090930192016116bc565b346102bd5760206003193601126102bd5760043567ffffffffffffffff81116102bd576117d4903690600401613f13565b6001600160a01b03600a5416331415806117f6575b611592576109b391614b6d565b506001600160a01b03600154163314156117e9565b346102bd5760206003193601126102bd5761168361182f61182a613bd9565b6145fe565b604051918291602083526020830190613db6565b346102bd5760606003193601126102bd5760043567ffffffffffffffff81116102bd5760a060031982360301126102bd5761187c613c6e565b9060443567ffffffffffffffff81116102bd5761189d903690600401613e72565b506118a661440a565b5060848101906118b5826141c8565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611e6357602481019177ffffffffffffffff0000000000000000000000000000000061190e846141dc565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d795760a05191611e34575b50611e0857611991604483016141c8565b7f0000000000000000000000000000000000000000000000000000000000000000611db5575b5067ffffffffffffffff6119ca846141dc565b166119e2816000526007602052604060002054151590565b15611d865760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611d795760a05190611d30575b6001600160a01b039150163303611d005760648201359061ffff851680151580611cf1575b15611c455761ffff600b541690818110611c13575050611c0994611bd89361182a937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff611b2c95611ae6611ad6611abc8c6141dc565b67ffffffffffffffff16600052600c602052604060002090565b85611ae0846141c8565b916156c5565b611b20611afb611af58c6141dc565b926141c8565b94604051938493169583602090939291936001600160a01b0360408201951681520152565b0390a25b6004016153f5565b927f0000000000000000000000000000000000000000000000000000000000000000611b79857f0000000000000000000000000000000000000000000000000000000000000000836147db565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff611bac846141dc565b604080516001600160a01b0390951685523360208601528401889052169180606081015b0390a26141dc565b90611be1615328565b60405192611bee84613d23565b83526020830152604051928392604084526040840190613ff9565b9060208301520390f35b7f7911d95b0000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b5050611b2c611c0994611bd89361182a937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff611c89896141dc565b16918260a05152600860205280611cd1604060a051206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916156c5565b604080516001600160a01b039290921682526020820192909252a2611b24565b5061ffff600b54161515611a5d565b7f728fe07b0000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b506020813d602011611d71575b81611d4a60209383613d77565b810103126102bd57516001600160a01b03811681036102bd576001600160a01b0390611a38565b3d9150611d3d565b6040513d60a051823e3d90fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b6001600160a01b0316611dd5816000526003602052604060002054151590565b6119b7577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b611e56915060203d602011611e5c575b611e4e8183613d77565b810190614b38565b85611980565b503d611e44565b6001600160a01b03611e74836141c8565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b346102bd5760206003193601126102bd5767ffffffffffffffff611ec6613bd9565b611ece61456d565b501660a05152600860205261168361162a6116256002604060a0512001614598565b346102bd5767ffffffffffffffff611f0736613f44565b929091611f1261479d565b1690611f2b826000526007602052604060002054151590565b15611fed578160a051526008602052611f5e6005604060a0512001611f51368685613e3b565b6020815191012090615d1d565b15611fa6577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611f9d60405192839260208452602084019161454c565b0390a260a05180f35b611fe9906040519384937f74f23c7c000000000000000000000000000000000000000000000000000000008552600485015260406024850152604484019161454c565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102bd5760a0516003193601126102bd5760a051506040516002548082526020820190600260a05152602060a051209060a0515b818110612075576116838561206981870382613d77565b60405191829182613f85565b8254845260209093019260019283019201612052565b346102bd5760206003193601126102bd5767ffffffffffffffff6120ad613bd9565b1660a0515260086020526120c86005604060a0512001615b07565b805190601f196120f06120da8461422e565b936120e86040519586613d77565b80855261422e565b0160a0515b8181106121d157505060a0515b815181101561214c578061211860019284614423565b5160a051526009602052612130604060a0512061448a565b61213a8286614423565b526121458185614423565b5001612102565b826040518091602082016020835281518091526040830190602060408260051b86010193019160a051905b82821061218657505050500390f35b919360206121c1827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851613db6565b9601920192018594939192612177565b8060606020809387010152016120f5565b346102bd5760206003193601126102bd5760043567ffffffffffffffff81116102bd5760a060031982360301126102bd5761221b61440a565b5060848101612229816141c8565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036126cc57602482019177ffffffffffffffff00000000000000000000000000000000612282846141dc565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d795760a051916126ad575b50611e0857612305604482016141c8565b7f000000000000000000000000000000000000000000000000000000000000000061265a575b5067ffffffffffffffff61233e846141dc565b16612356816000526007602052604060002054151590565b15611d865760206001600160a01b0360045416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611d795760a05190612611575b6001600160a01b039150163303611d0057606401359060a05160001461256c5761ffff600b5416806125375750826125079261182a926123fe6123f4611abc611683986141dc565b83611ae0846141c8565b7f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff61243a612434866141dc565b936141c8565b604080516001600160a01b03929092168252602082018690529190931692a25b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000926124d1817f0000000000000000000000000000000000000000000000000000000000000000866147db565b611bd06124dd866141dc565b604080516001600160a01b0390971687523360208801528601929092529116929081906060820190565b61250f615328565b6040519161251c83613d23565b82526020820152604051918291602083526020830190613ff9565b7f7911d95b0000000000000000000000000000000000000000000000000000000060a0515260a051600452602452604460a051fd5b5061182a826125079267ffffffffffffffff61258a611683966141dc565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894482806125f160406000206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916156c5565b604080516001600160a01b039290921682526020820192909252a261245a565b506020813d602011612652575b8161262b60209383613d77565b810103126102bd57516001600160a01b03811681036102bd576001600160a01b03906123ac565b3d915061261e565b6001600160a01b031661267a816000526003602052604060002054151590565b61232b577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b6126c6915060203d602011611e5c57611e4e8183613d77565b846122f4565b611e746001600160a01b03916141c8565b346102bd5760606003193601126102bd5760043567ffffffffffffffff81116102bd5761270e903690600401613e90565b9060243567ffffffffffffffff81116102bd5761272f903690600401613fc8565b9060443567ffffffffffffffff81116102bd57612750903690600401613fc8565b6001600160a01b03600a54163314158061280b575b61159257838614801590612801575b6127d55760a0515b8681106127895760a05180f35b806127cf61279d6109d46001948b8b6143bb565b6127a88389896143fa565b6127c96127c16127b986898b6143fa565b92369061404d565b91369061404d565b916151f8565b0161277c565b7f568efce20000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b5080861415612774565b506001600160a01b0360015416331415612765565b346102bd5760a0516003193601126102bd5760206001600160a01b0360015416604051908152f35b346102bd5760c06003193601126102bd57612861613c33565b5061286a613bf0565b612872613c5d565b5060843567ffffffffffffffff81116102bd57612893903690600401613c8e565b505060a4359060028210156102bd57611683916120699160443590614345565b346102bd5760206003193601126102bd5760206128ee67ffffffffffffffff6128da613bd9565b166000526007602052604060002054151590565b6040519015158152f35b346102bd5760206003193601126102bd577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b0361293c613c33565b61294461479d565b16807fffffffffffffffffffffffff0000000000000000000000000000000000000000600a541617600a55604051908152a160a05180f35b346102bd5760a0516003193601126102bd5760a051546001600160a01b0381163303612a0d577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660a051556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060a05160a051a360a05180f35b7f02b543c60000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346102bd5760a0516003193601126102bd57600454600554604080516001600160a01b039093168352602083019190915290f35b346102bd5760a0516003193601126102bd5760206001600160a01b03600a5416604051908152f35b346102bd5760406003193601126102bd57612aae613c33565b602435612ab961479d565b6001600160a01b038216918215610953577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600454161760045581600555612b3e60405192839283602090939291936001600160a01b0360408201951681520152565b0390a160a05180f35b346102bd57612b5536613f44565b612b6092919261479d565b67ffffffffffffffff8216612b82816000526007602052604060002054151590565b15612b9d57506109b392612b97913691613e3b565b90614e93565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346102bd5760a0516003193601126102bd576020612be861429a565b604051908152f35b346102bd5760406003193601126102bd5760043561ffff8116908190036102bd5760243567ffffffffffffffff81116102bd577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e91612c90612c586020933690600401613f13565b90612c6161479d565b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b55614b6d565b604051908152a160a05180f35b346102bd57612cc5612ccd612cb136613ec1565b9491612cbe93919361479d565b3691614246565b923691614246565b7f000000000000000000000000000000000000000000000000000000000000000015612dd55760a0515b8251811015612d5d57806001600160a01b03612d1560019386614423565b5116612d2081615b6a565b612d2c575b5001612cf7565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184612d25565b5060a0515b81518110156109b357806001600160a01b03612d8060019385614423565b51168015612dcf57612d9181615dd9565b612d9e575b505b01612d62565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183612d96565b50612d98565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346102bd5760406003193601126102bd57612e1a613bd9565b60243567ffffffffffffffff81116102bd57602091612e406128ee923690600401613e72565b906141f1565b346102bd5760406003193601126102bd5760043567ffffffffffffffff81116102bd57806004019061010060031982360301126102bd57612e85613c6e565b604051909190612e9481613d07565b60a0519052612ec5612ebb612eb6612eaf60c4850187614177565b3691613e3b565b614955565b6064830135614a49565b916084820190612ed4826141c8565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611e6357602483019477ffffffffffffffff00000000000000000000000000000000612f2d876141dc565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d795760a05191613317575b50611e085767ffffffffffffffff612fb6876141dc565b16612fce816000526007602052604060002054151590565b15611d865760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611d795760a051916132f8575b5015611d005761303a866141dc565b9061305060a4860192612e40612eaf8585614177565b156132b15750604493929161ffff161590506132125767ffffffffffffffff613078866141dc565b1660a05152600d602052613095604060a0512085611ae0846141c8565b7f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce61567ffffffffffffffff6130cb612434886141dc565b604080516001600160a01b03929092168252602082018890529190931692a25b01916130f6836141c8565b906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156102bd576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390921660048201526024810185905290818060448101038160a051875af18015611d79576131f9575b50608067ffffffffffffffff6020956001600160a01b036131c7611af57ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0966141dc565b60405196875233898801521660408601528560608601521692a2604051906131ee82613d07565b815260405190518152f35b60a05161320591613d77565b60a0516102bd5784613183565b5067ffffffffffffffff613225856141dc565b168060a0515260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c84806132916002604060a05120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916156c5565b604080516001600160a01b039290921682526020820192909252a26130eb565b6132bb9250614177565b611fe96040519283927f24eb47e500000000000000000000000000000000000000000000000000000000845260206004850152602484019161454c565b613311915060203d602011611e5c57611e4e8183613d77565b8761302b565b613330915060203d602011611e5c57611e4e8183613d77565b87612f9f565b346102bd5760206003193601126102bd5760043567ffffffffffffffff81116102bd57806004019061010060031982360301126102bd5760405161337981613d07565b60a051905260405161338a81613d07565b60a05190526133a5612ebb612eb6612eaf60c4850186614177565b90608481016133b3816141c8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116036126cc5750602481019277ffffffffffffffff0000000000000000000000000000000061340d856141dc565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611d795760a051916136b5575b50611e085767ffffffffffffffff613496856141dc565b166134ae816000526007602052604060002054151590565b15611d865760206001600160a01b0360045416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611d795760a05191613696575b5015611d005761351a846141dc565b9061353060a4840192612e40612eaf8585614177565b156132b1575082916044915067ffffffffffffffff61354e866141dc565b168060a0515260086020526135976002604060a05120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169586916156c5565b604080516001600160a01b0386168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a201926135dd846141c8565b823b156102bd576040517f40c10f1900000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390921660048201526024810185905290818060448101038160a051875af18015611d79576020956001600160a01b036131c7611af57ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09660809667ffffffffffffffff96613684575b506141dc565b60a05161369091613d77565b8b61367e565b6136af915060203d602011611e5c57611e4e8183613d77565b8561350b565b6136ce915060203d602011611e5c57611e4e8183613d77565b8561347f565b346102bd5760a06003193601126102bd576136ed613c33565b506136f6613bf0565b60443567ffffffffffffffff81116102bd5760031960a091360301126102bd5761371e613c5d565b506084359067ffffffffffffffff82116102bd5761374967ffffffffffffffff923690600401613c8e565b505060405161375781613cbc565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260c060a0519101521660a05152600f60205260e0604060a05120604051906137ab82613cbc565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b346102bd5760a0516003193601126102bd57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bd5760206003193601126102bd57602061389c613c33565b6040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081169216919091148152f35b346102bd5760a0516003193601126102bd5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346102bd5760a0516003193601126102bd57604080516116839161393e9082613d77565b602081527f4275726e546f41646472657373546f6b656e506f6f6c20312e372e302d6465766020820152604051918291602083526020830190613db6565b346102bd5760206003193601126102bd57613995613c33565b61399d61479d565b6139a561429a565b90816139b15760a05180f35b60206001600160a01b0382613a08857f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599957f00000000000000000000000000000000000000000000000000000000000000006147db565b6040519485521692a280806109b3565b346102bd5760c06003193601126102bd57613a31613bd9565b613a39613c07565b50613a42613c1d565b5060843561ffff811681036102bd5760a4359067ffffffffffffffff82116102bd5763ffffffff613a8961ffff92608095613a8284963690600401613c8e565b5050614094565b9392959091604051968752166020860152166040840152166060820152f35b346102bd5760206003193601126102bd57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036102bd57817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115613baf575b8115613b85575b8115613b5b575b8115613b31575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613b2a565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150613b23565b7ff57e5f940000000000000000000000000000000000000000000000000000000081149150613b1c565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150613b15565b6004359067ffffffffffffffff821682036109ae57565b6024359067ffffffffffffffff821682036109ae57565b602435906001600160a01b03821682036109ae57565b606435906001600160a01b03821682036109ae57565b600435906001600160a01b03821682036109ae57565b35906001600160a01b03821682036109ae57565b6064359061ffff821682036109ae57565b6024359061ffff821682036109ae57565b359061ffff821682036109ae57565b9181601f840112156109ae5782359167ffffffffffffffff83116109ae57602083818601950101116109ae57565b60e0810190811067ffffffffffffffff821117613cd857604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117613cd857604052565b6040810190811067ffffffffffffffff821117613cd857604052565b60a0810190811067ffffffffffffffff821117613cd857604052565b6060810190811067ffffffffffffffff821117613cd857604052565b90601f601f19910116810190811067ffffffffffffffff821117613cd857604052565b67ffffffffffffffff8111613cd857601f01601f191660200190565b919082519283825260005b848110613de2575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201613dc1565b346109ae5760006003193601126109ae5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b929192613e4782613d9a565b91613e556040519384613d77565b8294818452818301116109ae578281602093846000960137010152565b9080601f830112156109ae57816020613e8d93359101613e3b565b90565b9181601f840112156109ae5782359167ffffffffffffffff83116109ae576020808501948460051b0101116109ae57565b60406003198201126109ae5760043567ffffffffffffffff81116109ae5781613eec91600401613e90565b929092916024359067ffffffffffffffff82116109ae57613f0f91600401613e90565b9091565b9181601f840112156109ae5782359167ffffffffffffffff83116109ae5760208085019460e085020101116109ae57565b9060406003198301126109ae5760043567ffffffffffffffff811681036109ae57916024359067ffffffffffffffff82116109ae57613f0f91600401613c8e565b602060408183019282815284518094520192019060005b818110613fa95750505090565b82516001600160a01b0316845260209384019390920191600101613f9c565b9181601f840112156109ae5782359167ffffffffffffffff83116109ae57602080850194606085020101116109ae57565b613e8d9160206140128351604084526040840190613db6565b920151906020818403910152613db6565b359081151582036109ae57565b35906fffffffffffffffffffffffffffffffff821682036109ae57565b91908260609103126109ae5760405161406581613d5b565b604061408f81839561407681614023565b855261408460208201614030565b602086015201614030565b910152565b67ffffffffffffffff16600052600f602052604060002090604051916140b983613cbc565b549263ffffffff84169384845263ffffffff8160201c169384602082015263ffffffff8260401c169384604083015263ffffffff8360601c169081606084015261ffff8460801c169283608082015260ff61ffff8660901c16958660a084015260a01c16159060c082159101526141615761ffff161580159563ffffffff949392916141595750945b156141515750925b1693929190565b90509261414a565b905094614142565b5050505092505050600090600090600090600090565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109ae570180359067ffffffffffffffff82116109ae576020019181360383136109ae57565b356001600160a01b03811681036109ae5790565b3567ffffffffffffffff811681036109ae5790565b9067ffffffffffffffff613e8d92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111613cd85760051b60200190565b9291906142528161422e565b936142606040519586613d77565b602085838152019160051b81019283116109ae57905b82821061428257505050565b6020809161428f84613c49565b815201910190614276565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156143395760009161430a575090565b90506020813d602011614331575b8161432560209383613d77565b810103126109ae575190565b3d9150614318565b6040513d6000823e3d90fd5b67ffffffffffffffff16600052600e602052604060002091600281101561438c5760011461437b57816001613e8d930190615104565b8160026003613e8d94019101615104565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91908110156143cb5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b91908110156143cb576060020190565b6040519061441782613d23565b60606020838281520152565b80518210156143cb5760209160051b010190565b90600182811c92168015614480575b602083101461445157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614446565b906040519182600082549261449e84614437565b808452936001811690811561450c57506001146144c5575b506144c392500383613d77565b565b90506000929192526020600020906000915b8183106144f05750509060206144c392820101386144b6565b60209193508060019154838589010152019101909184926144d7565b602093506144c39592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b820101386144b6565b601f8260209493601f19938186528686013760008582860101520116010190565b6040519061457a82613d3f565b60006080838281528260208201528260408201528260608201520152565b906040516145a581613d3f565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526008602052613e8d600460406000200161448a565b91908110156143cb5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61813603018212156109ae570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156109ae570180359067ffffffffffffffff82116109ae57602001918160051b360383136109ae57565b818102929181159184041417156146c757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110614701575050565b600081556001016146f6565b9160209082815201919060005b8181106147275750505090565b9091926020806001926001600160a01b0361474188613c49565b16815201940192910161471a565b91908110156143cb5760081b0190565b3561ffff811681036109ae5790565b3563ffffffff811681036109ae5790565b3580151581036109ae5790565b359063ffffffff821682036109ae57565b6001600160a01b036001541633036147b157565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082019081526001600160a01b0393841660248301526044808301959095529381526148a4929091614833606484613d77565b166000806040958651946148478887613d77565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1561494d573d9161488883613d9a565b9261489587519485613d77565b83523d6000602085013e615ee8565b805190816148b157505050565b6020806148c2938301019101614b38565b156148ca5750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091615ee8565b805180156149c5576020036149875780516020828101918301839003126109ae57519060ff8211614987575060ff1690565b611fe9906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190613db6565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff82116146c757565b60ff16604d81116146c757600a0a90565b8115614a1a570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414614b3157828411614b075790614a8e916149eb565b91604d60ff8416118015614aec575b614ab657505090614ab0613e8d926149ff565b906146b4565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50614af6836149ff565b8015614a1a57600019048411614a9d565b614b10916149eb565b91604d60ff841611614ab657505090614b2b613e8d926149ff565b90614a10565b5050505090565b908160209103126109ae575180151581036109ae5790565b356fffffffffffffffffffffffffffffffff811681036109ae5790565b9160005b82811015614e8d5760e0810284016000614b8a826141dc565b9067ffffffffffffffff821691614bae836000526007602052604060002054151590565b15614e6157614c779260408593614c22614c1c94614c1c614be2602060019c9b0192611abc614bdd368661404d565b61557e565b91825463ffffffff8160801c16159081614e43575b81614e34575b81614e19575b81614e0a575b5080614dfb575b614d70575b369061404d565b906158d4565b6080850192614c34614bdd368661404d565b8152600d6020522092835463ffffffff8160801c16159081614d52575b81614d43575b81614d28575b81614d19575b5080614d0a575b614c7d575b50369061404d565b01614b71565b614c9a60a06fffffffffffffffffffffffffffffffff9201614b50565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538614c6f565b50614d148261477f565b614c6a565b60ff915060a01c161538614c63565b6fffffffffffffffffffffffffffffffff8116159150614c5d565b8589015460801c159150614c57565b858901546fffffffffffffffffffffffffffffffff16159150614c51565b6fffffffffffffffffffffffffffffffff614d8c878b01614b50565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355614c15565b50614e058161477f565b614c10565b60ff915060a01c161538614c09565b6fffffffffffffffffffffffffffffffff8116159150614c03565b848e015460801c159150614bfd565b848e01546fffffffffffffffffffffffffffffffff16159150614bf7565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b908051156150795767ffffffffffffffff81516020830120921691826000526008602052614ec8816005604060002001615e93565b156150355760005260096020526040600020815167ffffffffffffffff8111613cd857614ef58254614437565b601f8111615003575b506020601f8211600114614f795791614f53827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593614f6995600091614f6e575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190613db6565b0390a2565b905084015138614f40565b601f1982169083600052806000209160005b818110614feb575092614f699492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610614fd2575b5050811b01905561182f565b85015160001960f88460031b161c191690553880614fc6565b9192602060018192868a015181550194019201614f8b565b61502f90836000526020600020601f840160051c810191602085106108dd57601f0160051c01906146f6565b38614efe565b5090611fe96040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190613db6565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b8181106150d55750506144c392500383613d77565b84546001600160a01b03168352600194850194879450602090930192016150c0565b919082018092116146c757565b61510d906150a3565b9160055480151591826151ed575b5050615125575090565b61512e906150a3565b9081518061513c5750905090565b6151479082516150f7565b92601f1961516d6151578661422e565b956151656040519788613d77565b80875261422e565b0136602086013760005b82518110156151a857806001600160a01b0361519560019386614423565b51166151a18288614423565b5201615177565b509160005b81518110156151e857806001600160a01b036151cb60019385614423565b51166151e16151db8387516150f7565b88614423565b52016151ad565b505050565b10159050388061511b565b67ffffffffffffffff1660008181526007602052604090205490929190156152fa57916152f760e0926152c38561524f7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b9761557e565b8460005260086020526152668160406000206158d4565b61526f8361557e565b8460005260086020526152898360026040600020016158d4565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152613e8d604082613d77565b919082039182116146c757565b61537861456d565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916153d560208501936153cf6153c263ffffffff87511642615363565b85608089015116906146b4565b906150f7565b808210156153ee57505b16825263ffffffff4216905290565b90506153df565b67ffffffffffffffff61540a602083016141dc565b16600052600f60205260406000206040519061542582613cbc565b549063ffffffff8216815263ffffffff8260201c16602082015263ffffffff8260401c16604082015263ffffffff8260601c16606082015261ffff8260801c169081608082015260ff61ffff8460901c16938460a084015260a01c16159060c082159101526154d057919261ffff92908316156154c957505b1680156154c1576127106154ba6060613e8d94013592836146b4565b0490615363565b506060013590565b905061549e565b505060609150013590565b805160005b8181106154ec57505050565b600181018082116146c7575b82811061550857506001016154e0565b6001600160a01b0361551a8386614423565b51166001600160a01b0361552e8387614423565b51161461553d576001016154f8565b6001600160a01b0361554f8386614423565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b80511561561e576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106155bb5750565b60649061561c604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff604082015116158015906156a6575b6156455750565b60649061561c604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff602082015116151561563e565b9182549060ff8260a01c161580156158cc575b6158c6576fffffffffffffffffffffffffffffffff8216916001850190815461571d63ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615363565b9081615828575b50508481106157e9575083831061577e5750506157536fffffffffffffffffffffffffffffffff928392615363565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c9161578d8185615363565b926000198101908082116146c7576157b06157b5926001600160a01b03966150f7565b614a10565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82869293961161589c57615843926153cf9160801c906146b4565b808410156158975750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880615724565b61584e565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156156d8565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991615a0d606092805461591163ffffffff8260801c1642615363565b9081615a4c575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014615a4457508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556159c18651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6152f760405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091615948565b6fffffffffffffffffffffffffffffffff91615a81839283615a7a6001880154948286169560801c906146b4565b91166150f7565b80821015615b0057505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880615918565b9050615a8b565b906040519182815491828252602082019060005260206000209260005b818110615b395750506144c392500383613d77565b8454835260019485019487945060209093019201615b24565b80548210156143cb5760005260206000200190600090565b6000818152600360205260409020548015615c635760001981018181116146c7576002549060001982019182116146c757818103615c12575b5050506002548015615be35760001901615bbe816002615b52565b60001982549160031b1b19169055600255600052600360205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b615c4b615c23615c34936002615b52565b90549060031b1c9283926002615b52565b81939154906000199060031b92831b921b19161790565b90556000526003602052604060002055388080615ba3565b5050600090565b6000818152600760205260409020548015615c635760001981018181116146c7576006549060001982019182116146c757818103615ce3575b5050506006548015615be35760001901615cbe816006615b52565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b615d05615cf4615c34936006615b52565b90549060031b1c9283926006615b52565b90556000526007602052604060002055388080615ca3565b9060018201918160005282602052604060002054801515600014615dd05760001981018181116146c75782549060001982019182116146c757818103615d99575b50505080548015615be3576000190190615d788282615b52565b60001982549160031b1b191690555560005260205260006040812055600190565b615db9615da9615c349386615b52565b90549060031b1c92839286615b52565b905560005283602052604060002055388080615d5e565b50505050600090565b80600052600360205260406000205415600014615e335760025468010000000000000000811015613cd857615e1a615c348260018594016002556002615b52565b9055600254906000526003602052604060002055600190565b50600090565b80600052600760205260406000205415600014615e335760065468010000000000000000811015613cd857615e7a615c348260018594016006556006615b52565b9055600654906000526007602052604060002055600190565b6000828152600182016020526040902054615c635780549068010000000000000000821015613cd85782615ed1615c34846001809601855584615b52565b905580549260005201602052604060002055600190565b91929015615f635750815115615efc575090565b3b15615f055790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015615f765750805190602001fd5b611fe9906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190613db656fea164736f6c634300081a000a",
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _BurnToAddressMintTokenPool.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) GetFee(destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnToAddressMintTokenPool.Contract.GetFee(&_BurnToAddressMintTokenPool.CallOpts, destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolCallerSession) GetFee(destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _BurnToAddressMintTokenPool.Contract.GetFee(&_BurnToAddressMintTokenPool.CallOpts, destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "applyCustomBlockConfirmationConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_BurnToAddressMintTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactor) SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.contract.Transact(opts, "setCustomBlockConfirmationRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolTransactorSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnToAddressMintTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_BurnToAddressMintTokenPool.TransactOpts, rateLimitConfigArgs)
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

type BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnToAddressMintTokenPoolCustomBlockConfirmationUpdatedIterator struct {
	Event *BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated)
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
		it.Event = new(BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated)
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

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnToAddressMintTokenPoolCustomBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolCustomBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnToAddressMintTokenPoolCustomBlockConfirmationUpdatedIterator{contract: _BurnToAddressMintTokenPool.contract, event: "CustomBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnToAddressMintTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated)
				if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (_BurnToAddressMintTokenPool *BurnToAddressMintTokenPoolFilterer) ParseCustomBlockConfirmationUpdated(log types.Log) (*BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated, error) {
	event := new(BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated)
	if err := _BurnToAddressMintTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e")
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

	GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

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

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*BurnToAddressMintTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*BurnToAddressMintTokenPoolCustomBlockConfirmationUpdatedIterator, error)

	WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomBlockConfirmationUpdated(log types.Log) (*BurnToAddressMintTokenPoolCustomBlockConfirmationUpdated, error)

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
