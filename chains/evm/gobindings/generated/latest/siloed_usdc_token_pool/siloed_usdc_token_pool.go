// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package siloed_usdc_token_pool

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

type AuthorizedCallersAuthorizedCallerArgs struct {
	AddedCallers   []common.Address
	RemovedCallers []common.Address
}

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

var SiloedUSDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"struct AuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCustomBlockConfirmationConfigUpdates\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfiguredMinBlockConfirmation\",\"inputs\":[],\"outputs\":[{\"name\":\"blockConfirmationConfigured\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentCustomBlockConfirmationRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"outboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterState\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFee\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockConfirmationRequested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeUSDCents\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tokenFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRateLimitAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enum IPoolV2.MessageDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contract IERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Client.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"struct Client.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct Pool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"struct RateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomBlockConfirmationRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"struct TokenPool.CustomBlockConfirmationRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRateLimitAdmin\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"struct SiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomBlockConfirmationUpdated\",\"inputs\":[{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminSet\",\"inputs\":[{\"name\":\"rateLimitAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct IPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"customBlockConfirmationFeeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customBlockConfirmationTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyMigrated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmation\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minBlockConfirmation\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidMinBlockConfirmationConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"struct RateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCircle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenLockingNotAllowedAfterMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120806040523461056557618270803803809161001d82856107fb565b8339810160c082820312610565578151906001600160a01b038216908183036105655761004c6020850161081e565b60408501519091906001600160401b0381116105655785019080601f83011215610565578151916001600160401b0383116104d5578260051b90602082019361009860405195866107fb565b845260208085019282010192831161056557602001905b8282106107e3575050506100c56060860161082c565b946100de60a06100d76080840161082c565b920161082c565b92602096604051966100f089896107fb565b60008852600036813733156107d257600180546001600160a01b03191633179055861580156107c1575b80156107b0575b61061b5760805260c05260405163313ce56760e01b81528781600481895afa60009181610779575b5061074e575b5060a052600480546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261062c575b506001600160a01b0316801561061b57604051636eb1769f60e11b8152306004820152602481018290528481604481865afa90811561060f576000916105e2575b5061057757604051918483019263095ea7b360e01b84528260248201526000196044820152604481526101f66064826107fb565b60008060409586519361020988866107fb565b8985527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648a860152519082865af13d1561056a573d906001600160401b0382116104d5578551610276949092610268601f8201601f19168b01856107fb565b83523d60008a85013e610ac0565b848151806104eb575b50505061010052805161029284826107fb565b6000815260003681378151928383016001600160401b038111858210176104d55783528352808484015260005b8151811015610325576001906001600160a01b036102dd8285610840565b5116866102e9826108c1565b6102f6575b5050016102bf565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580918651908152a138866102ee565b505090519060005b825181101561039d576001600160a01b036103488285610840565b511690811561038c577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef858361037f6001956109f3565b508551908152a10161032d565b6342bcdf7f60e11b60005260046000fd5b50516176df9081610b9182396080518181816104d901528181611c0901528181611fac015281816121b0015281816122e8015281816127910152818161293e01528181612b9001528181613058015281816140bb015281816142810152818161433101528181614475015281816145a301528181614acc01528181614cfd01528181614d4b01528181614e8701528181614f9401526158e4015260a051818181614c5101528181615e8a01528181615f75015281816160e101526169f4015260c05181818161114c0152818161203a0152818161281f015281816141490152614504015260e051818181610f0c0152818161207f015281816128640152613ca701526101005181818161052401528181611bae01528181612a400152818161302e0152818161467001528181614b070152614f390152f35b634e487b7160e01b600052604160045260246000fd5b829081010312610565578401518015908115036105655761050e5738848161027f565b815162461bcd60e51b815260048101859052602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b9161027692606091610ac0565b60405162461bcd60e51b815260048101859052603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90508481813d8311610608575b6105f981836107fb565b810103126105655751386101c2565b503d6105ef565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b604051929361063b86856107fb565b60008452600036813760e0511561073d5760005b84518110156106b6576001906001600160a01b0361066d8288610840565b51168861067982610a2c565b610686575b50500161064f565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1388861067e565b50919390925060005b8351811015610733576001906001600160a01b036106dd8287610840565b5116801561072d57876106ef826109b4565b6106fd575b50505b016106bf565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138876106f4565b506106f7565b5091509138610181565b6335f4a7b360e01b60005260046000fd5b60ff1660ff8216818103610762575061014f565b6332ad3e0760e11b60005260045260245260446000fd5b9091508881813d83116107a9575b61079181836107fb565b81010312610565576107a29061081e565b9038610149565b503d610787565b506001600160a01b03821615610121565b506001600160a01b0384161561011a565b639b15e16f60e01b60005260046000fd5b602080916107f08461082c565b8152019101906100af565b601f909101601f19168101906001600160401b038211908210176104d557604052565b519060ff8216820361056557565b51906001600160a01b038216820361056557565b80518210156108545760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b80548210156108545760005260206000200190600090565b805480156108ab576000190190610899828261086a565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b600081815260146020526040902054801561098257600019810181811161096c5760135460001981019190821161096c5780820361091b575b5050506109076013610882565b600052601460205260006040812055600190565b61095461092c61093d93601361086a565b90549060031b1c928392601361086a565b819391549060031b91821b91600019901b19161790565b905560005260146020526040600020553880806108fa565b634e487b7160e01b600052601160045260246000fd5b5050600090565b805490680100000000000000008210156104d5578161093d9160016109b09401815561086a565b9055565b806000526003602052604060002054156000146109ed576109d6816002610989565b600254906000526003602052604060002055600190565b50600090565b806000526014602052604060002054156000146109ed57610a15816013610989565b601354906000526014602052604060002055600190565b600081815260036020526040902054801561098257600019810181811161096c5760025460001981019190821161096c57818103610a86575b505050610a726002610882565b600052600360205260006040812055600190565b610aa8610a9761093d93600261086a565b90549060031b1c928392600261086a565b90556000526003602052604060002055388080610a65565b91929015610b225750815115610ad4575090565b3b15610add5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610b355750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610b785750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610b5656fe60c080604052600436101561001357600080fd5b600060a05260a0513560e01c90816301ffc9a7146150de575080630574d6a61461504e5780630a861f2a14614ebb578063164e68de14614dd1578063181f5a7714614d6f57806321df0da714614d2a578063240028e814614cd75780632451a62714614c7557806324f65ee714614c365780632d4a148f146149d957806331238ffc1461499157806337b19247146148235780633907753714614407578063432a6ba3146143df578063489a68f2146140255780634ad01f0b14613f885780634c5ef0ed14613f4357806350d1a35a14613dd957806354c8a4f314613c7557806359152aad14613bc85780635df45a3714613bac57806362ddd3c414613b275780636600f92c14613a38578063698c2c661461399a5780636cfd1553146139055780636d3d1a58146138dd5780636d9d216c146134d8578063714bf9071461345f5780637437ff9f1461342b57806379ba5097146133795780637d54534e146133005780638632d5cc146132cc5780638926f54f1461328757806389720a621461321c5780638a5e52bb14612fc15780638da5cb5b14612f9957806391a2749a14612e03578063962d402014612cc05780639751f88414612c5b5780639a4575b914612736578063a42a7b8b146125df578063a7cd63b714612571578063acfecf911461244d578063af0e58b91461242e578063b1c71c6514611f2f578063b794658014611ef3578063bb6bb5a714611e8b578063c4bffe2b14611d6f578063cd306a6c14611d43578063ce3c752814611af5578063cf7401f3146119bc578063d966866b1461156f578063dc04fa1f14611170578063dc0bd9711461112b578063de814c5714611047578063ded8d95614610f31578063e0351e1314610ef3578063e8a1da17146106ae578063eb521a4c1461041f578063f1e73399146103f4578063f2fde38b14610340578063f37c77fe1461031d5763f65a8886146102da57600080fd5b346103175760206003193601126103175767ffffffffffffffff6102fc61520f565b1660a0515260166020526020604060a0512054604051908152f35b60a05180fd5b346103175760a05160031936011261031757602061ffff600b5416604051908152f35b34610317576020600319360112610317576001600160a01b03610361615269565b610369615e05565b163381146103c85760a051805473ffffffffffffffffffffffffffffffffffffffff1916821781556001546001600160a01b0316907fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789080a360a05180f35b7fdad89dca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461031757602060031936011261031757602061041761041261520f565b615d7e565b604051908152f35b346103175760206003193601126103175760a08051805260186020525160409020546004359061067c578015610650576001600160a01b0361046260a051615996565b1633036106205760a05160a051526012602052604060a0512060ff600182015460a01c1660001461060b57610498828254615979565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f00000000000000000000000000000000000000000000000000000000000000009061051a9061051481608481015b03601f198101835282615418565b82616bdc565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b15610317576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a080516001600160a01b039390931660048301526024820185905251909283916044918391905af180156105fe576105e5575b506040519060a0515060a051825260208201527f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf60403392a260a05180f35b60a0516105f191615418565b60a05161031757816105a6565b6040513d60a051823e3d90fd5b5061061881601054615979565b60105561049b565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060a0515233600452602460a051fd5b7fa90c0d190000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b7f646972460000000000000000000000000000000000000000000000000000000060a0515260a051600452602460a051fd5b34610317576106bc366154dd565b9190926106c7615e05565b60a051905b828210610d315750505060a0519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610d2b57600581901b8301358581121561031757830161012081360312610317576040519461073b866153fc565b813567ffffffffffffffff81168103610d26578652602082013567ffffffffffffffff81116103175782019436601f8701121561031757853561077d816155a1565b9661078b6040519889615418565b81885260208089019260051b820101903682116103175760208101925b828410610cf7575050505060208701958652604083013567ffffffffffffffff8111610317576107db903690850161548e565b91604088019283526108056107f336606087016156a4565b9460608a0195865260c03691016156a4565b95608089019687526108178551616acb565b6108218751616acb565b83515115610ccb5761083d67ffffffffffffffff8a51166172f5565b15610c905767ffffffffffffffff89511660a051526008602052604060a0512061095d86516001600160801b03604082015116906109216001600160801b0360208301511691511515836080604051610895816153fc565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166001600160801b0384161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176001830155565b610a5f88516001600160801b0360408201511690610a236001600160801b0360208301511691511515836080604051610995816153fc565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166001600160801b0385161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166001600160801b0391909116176003830155565b6004855191019080519067ffffffffffffffff8211610c7857610a828354615af1565b601f8111610c39575b506020906001601f841114610bcf57918091610abe9360a05192610bc4575b50506000198260011b9260031b1c19161790565b90555b60a0515b88518051821015610afa5790610af4600192610aed8367ffffffffffffffff8f511692615add565b519061641e565b01610ac5565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610bb667ffffffffffffffff6001979694985116925193519151610b8b610b5f604051968796875261010060208801526101008701906152f2565b9360408601906001600160801b0360408092805115158552826020820151166020860152015116910152565b60a08401906001600160801b0360408092805115158552826020820151166020860152015116910152565b0390a1019392909193610709565b015190508e80610aaa565b90601f198316918460a051528160a051209260a0515b818110610c215750908460019594939210610c08575b505050811b019055610ac1565b015160001960f88460031b161c191690558d8080610bfb565b92936020600181928786015181550195019301610be5565b610c68908460a05152602060a05120601f850160051c81019160208610610c6e575b601f0160051c0190615cd7565b8d610a8b565b9091508190610c5b565b634e487b7160e01b60a051526041600452602460a051fd5b67ffffffffffffffff8951167f1d5ad3c50000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f14c880ca0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b833567ffffffffffffffff811161031757602091610d1b839283369187010161548e565b8152019301926107a8565b600080fd5b60a05180f35b9092919367ffffffffffffffff610d51610d4c868886615953565b6157ce565b1692610d5c8461743e565b15610ec3578360a051526008602052610d7c6005604060a0512001616d0f565b9260a0515b8451811015610dbb576001908660a051526008602052610db46005604060a0512001610dad8389615add565b51906174d2565b5001610d81565b50939094919592508060a0515260086020526005604060a0512060a051815560a051600182015560a051600282015560a051600382015560048101610e008154615af1565b80610e73575b50500180549060a051815581610e4f575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916106cc565b60a05152602060a05120908101905b81811015610e175760a0518155600101610e5e565b601f8111600114610e8c575060a05190555b8880610e06565b610eac908260a051526001601f602060a051209201861c82019101615cd7565b60a080518290525160208120918190559055610e85565b837f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346103175760a0516003193601126103175760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b3461031757602060031936011261031757610140610f4d61520f565b610f55615a45565b50610f5e615a45565b50611045610fb4610f94610f8f610f99610f94610f8f8767ffffffffffffffff16600052600c602052604060002090565b615a70565b616894565b9467ffffffffffffffff16600052600d602052604060002090565b610ffe60405180946001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b60a08301906001600160801b036080809282815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565bf35b346103175760406003193601126103175761106061520f565b60243561106b615e05565b67ffffffffffffffff8060155460a01c1692168092036110ff5760407fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220918360a0515260166020528160a051206110c3828254615979565b90558360a0515260126020526110ee8260a05120548560a0515260166020528360a0512054906157e3565b82519182526020820152a260a05180f35b7fa94cb9880000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346103175760a0516003193601126103175760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103175760406003193601126103175760043567ffffffffffffffff8111610317573660238201121561031757806004013567ffffffffffffffff8111610317576024820191602436918360081b0101116103175760243567ffffffffffffffff8111610317576111e69036906004016154ac565b9190926111f1615e05565b60a0515b82811061126d5750505060a0515b8181106112105760a05180f35b8067ffffffffffffffff61122a610d4c6001948688615953565b168060a05152600f60205260a051604060a05120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee860a05160a051a201611203565b61127b610d4c828585615d30565b611286828585615d30565b602081019060a081019261271061ffff61129f86615d40565b1610156115635760c082019061271061ffff6112ba84615d40565b1610156115275767ffffffffffffffff16938460a05152600f60205260a05160409020926112e785615d4f565b63ffffffff1684549160408101906112fe82615d4f565b60201b67ffffffff000000001693606082019361131a85615d4f565b60401b6bffffffff00000000000000001695608084019661133a88615d4f565b60601b6fffffffff00000000000000000000000016916113598a615d40565b60801b71ffff00000000000000000000000000000000169361137a8c615d40565b60901b73ffff00000000000000000000000000000000000016957fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016177fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff16177fffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffff16171717875560e0019561143187615d60565b81547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690151560a01b74ff0000000000000000000000000000000000000000161790556040519661148290615d6d565b63ffffffff16875261149390615d6d565b63ffffffff1660208701526114a790615d6d565b63ffffffff1660408601526114bb90615d6d565b63ffffffff1660608501526114cf906152b5565b61ffff1660808401526114e1906152b5565b61ffff1660a08301526114f390615683565b151560c082015260e07ffae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa104191a26001016111f5565b61ffff61153383615d40565b7f95f3517a0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b61ffff61153385615d40565b346103175760206003193601126103175760043567ffffffffffffffff8111610317576115a09036906004016154ac565b906115a9615e05565b60a0516080525b81608051106115bf5760a05180f35b6115cf610d4c6080518484615c30565b6115e96115df6080518585615c30565b6020810190615c70565b6116036115f96080518787615c30565b6040810190615c70565b9061161e6116146080518989615c30565b6060810190615c70565b61163861162e6080518b8b615c30565b6080810190615c70565b93909461164e61164936898b6155b9565b616a28565b61165c6116493683856155b9565b61166a6116493685876155b9565b6116786116493687896155b9565b6040519860808a01908a821067ffffffffffffffff831117610c785767ffffffffffffffff916040526116ac368a8c6155b9565b8b526116b93684866155b9565b60208c01526116c93686886155b9565b60408c01526116d936888a6155b9565b60608c015216988960a05152600e602052604060a05120815180519067ffffffffffffffff8211610c7857680100000000000000008211610c7857602090835483855580841061199d575b50018260a05152602060a0512060a0515b83811061198057505050506001810160208301519081519167ffffffffffffffff8311610c7857680100000000000000008311610c78576020908254848455808510611961575b50019060a05152602060a0512060a0515b83811061194457505050506002810160408301519081519167ffffffffffffffff8311610c7857680100000000000000008311610c78576020908254848455808510611925575b50019060a05152602060a0512060a0515b838110611908575050505060036060910191015180519067ffffffffffffffff8211610c7857680100000000000000008211610c785760209083548385558084106118e9575b50019160a05152602060a051209160a0515b8281106118cc5750505050926118bb949261189f6118ad937fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9998966118916040519b8c9b60808d5260808d0191615cee565b918a830360208c0152615cee565b918783036040890152615cee565b918483036060860152615cee565b0390a26001608051016080526115b0565b60019060206001600160a01b03845116930192818601550161183d565b611902908560a05152848460a051209182019101615cd7565b8f61182b565b60019060206001600160a01b0385511694019381840155016117e5565b61193e908460a05152858460a051209182019101615cd7565b386117d4565b60019060206001600160a01b03855116940193818401550161178d565b61197a908460a05152858460a051209182019101615cd7565b3861177c565b60019060206001600160a01b038551169401938184015501611735565b6119b6908560a05152848460a051209182019101615cd7565b38611724565b346103175760e0600319360112610317576119d561520f565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc011261031757604051611a0b816153c4565b60243580151581036103175781526044356001600160801b03811681036103175760208201526064356001600160801b03811681036103175760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c01126103175760405190611a80826153c4565b608435801515810361031757825260a4356001600160801b038116810361031757602083015260c4356001600160801b03811681036103175760408301526001600160a01b03600a541633141580611ae0575b61062057610d2b92616776565b506001600160a01b0360015416331415611ad3565b3461031757604060031936011261031757611b0e61520f565b60243567ffffffffffffffff82168060a05152601260205260ff6001604060a05120015460a01c16158015611d3b575b611d0c578115610650576001600160a01b03611b5984615996565b1633036106205760a051526012602052604060a0512060ff600182015460a01c1660a0515080600014611d045781545b808411611cd0575015611cbb57611ba18282546157e3565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610317576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af180156105fe57611ca2575b506040805167ffffffffffffffff9093168352602083019190915233917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e91819081015b0390a260a05180f35b60a051611cae91615418565b60a0516103175782611c55565b50611cc8816010546157e3565b601055611ba4565b83907fa17e11d50000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b601054611b89565b7f46f5f12b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b508015611b3e565b346103175760a05160031936011261031757602067ffffffffffffffff60155460a01c16604051908152f35b346103175760a0516003193601126103175760a051506040516006548082528160208101600660a05152602060a051209260a0515b818110611e72575050611db992500382615418565b805190611dde611dc8836155a1565b92611dd66040519485615418565b8084526155a1565b90601f1960208401920136833760a0515b8151811015611e21578067ffffffffffffffff611e0e60019385615add565b5116611e1a8287615add565b5201611def565b5050906040519182916020830190602084525180915260408301919060a0515b818110611e4f575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611e41565b8454835260019485019486945060209093019201611da4565b346103175760206003193601126103175760043567ffffffffffffffff811161031757611ebc90369060040161552f565b6001600160a01b03600a541633141580611ede575b61062057610d2b91616117565b506001600160a01b0360015416331415611ed1565b3461031757602060031936011261031757611f2b611f17611f1261520f565b615c0e565b6040519182916020835260208301906152f2565b0390f35b346103175760606003193601126103175760043567ffffffffffffffff81116103175760a0600319823603011261031757611f686152a4565b9060443567ffffffffffffffff811161031757611f8990369060040161548e565b50611f92615ac4565b506084810190611fa182615806565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036123ed57602481019177ffffffffffffffff00000000000000000000000000000000611ffa846157ce565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105fe5760a051916123be575b506123925761207d60448301615806565b7f000000000000000000000000000000000000000000000000000000000000000061233f575b506120b56120b0846157ce565b616d5a565b60648201359061ffff851680151580612330575b156122845761ffff600b5416908181106122525750506122489461221793611f12937f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff612198956121526121426121288c6157ce565b67ffffffffffffffff16600052600c602052604060002090565b8561214c84615806565b91616df2565b61218c6121676121618c6157ce565b92615806565b94604051938493169583602090939291936001600160a01b0360408201951681520152565b0390a25b600401616907565b926121a2816157ce565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a26157ce565b906122206169ed565b6040519261222d846153e0565b83526020830152604051928392604084526040840190615659565b9060208301520390f35b7f7911d95b0000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b50506121986122489461221793611f12937fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894467ffffffffffffffff6122c8896157ce565b16918260a05152600860205280612310604060a051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616df2565b604080516001600160a01b039290921682526020820192909252a2612190565b5061ffff600b541615156120c9565b6001600160a01b031661235f816000526003602052604060002054151590565b6120a3577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f53ad11d80000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b6123e0915060203d6020116123e6575b6123d88183615418565b810190616055565b8561206c565b503d6123ce565b6001600160a01b036123fe83615806565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b346103175760a051600319360112610317576020601054604051908152f35b346103175767ffffffffffffffff61246436615560565b92909161246f615e05565b1690612488826000526007602052604060002054151590565b15612541578160a0515260086020526124bb6005604060a05120016124ae368685615457565b60208151910120906174d2565b156124fa577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611c99604051928392602084526020840191615bed565b61253d906040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191615bed565b0390fd5b507f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346103175760a0516003193601126103175760a051506040516002548082526020820190600260a05152602060a051209060a0515b8181106125c957611f2b856125bd81870382615418565b60405191829182615333565b82548452602090930192600192830192016125a6565b346103175760206003193601126103175767ffffffffffffffff61260161520f565b1660a05152600860205261261c6005604060a0512001616d0f565b805190601f1961264461262e846155a1565b9361263c6040519586615418565b8085526155a1565b0160a0515b81811061272557505060a0515b81518110156126a0578061266c60019284615add565b5160a051526009602052612684604060a05120615b2b565b61268e8286615add565b526126998185615add565b5001612656565b826040518091602082016020835281518091526040830190602060408260051b86010193019160a051905b8282106126da57505050500390f35b91936020612715827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0600195979984950301865288516152f2565b96019201920185949391926126cb565b806060602080938701015201612649565b346103175760206003193601126103175760043567ffffffffffffffff81116103175760a060031982360301126103175761276f615ac4565b50612778615ac4565b506084810161278681615806565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612c4a57602482019177ffffffffffffffff000000000000000000000000000000006127df846157ce565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105fe5760a05191612c2b575b506123925761286260448201615806565b7f0000000000000000000000000000000000000000000000000000000000000000612bd8575b506064906128986120b0856157ce565b01359060a051600014612b3e5761ffff600b541680612b0957506128cb6128c1612128856157ce565b8361214c84615806565b7f61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac2097932867ffffffffffffffff612907612901866157ce565b93615806565b604080516001600160a01b03929092168252602082018690529190931692a25b612930826157ce565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016808252336020830152918101849052909167ffffffffffffffff16907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a26129ac611f12846157ce565b926129b56169ed565b604051946129c2866153e0565b8552602085015267ffffffffffffffff6129db826157ce565b1660a05152601260205260ff6001604060a05120015460a01c16600014612af4578067ffffffffffffffff612a12612a34936157ce565b1660a051526012602052604060a05120612a2d858254615979565b90556157ce565b505b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b15610317576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a0516001600160a01b0390931660048201526024810191909152918290818060448101039160a051905af180156105fe57612adb575b60405160208082528190611f2b90820185615659565b60a051612ae791615418565b60a0516103175781612ac5565b50612b0182601054615979565b601055612a36565b7f7911d95b0000000000000000000000000000000000000000000000000000000060a0515260a051600452602452604460a051fd5b5067ffffffffffffffff612b51836157ce565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280612bb860406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616df2565b604080516001600160a01b039290921682526020820192909252a2612927565b6001600160a01b0316612bf8816000526003602052604060002054151590565b612888577fd0d259760000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b612c44915060203d6020116123e6576123d88183615418565b84612851565b6123fe6001600160a01b0391615806565b346103175760206003193601126103175767ffffffffffffffff612c7d61520f565b612c85615a45565b50612c8e615a45565b501660a051526008602052610140604060a05120611045610fb4610f946002612cb9610f9486615a70565b9401615a70565b346103175760606003193601126103175760043567ffffffffffffffff811161031757612cf19036906004016154ac565b9060243567ffffffffffffffff811161031757612d12903690600401615628565b9060443567ffffffffffffffff811161031757612d33903690600401615628565b6001600160a01b03600a541633141580612dee575b61062057838614801590612de4575b612db85760a0515b868110612d6c5760a05180f35b80612db2612d80610d4c6001948b8b615953565b612d8b838989615a35565b612dac612da4612d9c86898b615a35565b9236906156a4565b9136906156a4565b91616776565b01612d5f565b7f568efce20000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b5080861415612d57565b506001600160a01b0360015416331415612d48565b346103175760206003193601126103175760043567ffffffffffffffff811161031757604060031982360301126103175760405190612e41826153e0565b806004013567ffffffffffffffff811161031757612e65906004369184010161560d565b825260248101359067ffffffffffffffff8211610317576004612e8b923692010161560d565b60208201908152612e9a615e05565b519060a0515b8251811015612f0657806001600160a01b03612ebe60019386615add565b5116612ec981617572565b612ed5575b5001612ea0565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a184612ece565b505160a0515b8151811015610d2b576001600160a01b03612f278284615add565b5116908115612f6d577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef602083612f5f6001956172bc565b50604051908152a101612f0c565b7f8579befe0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346103175760a0516003193601126103175760206001600160a01b0360015416604051908152f35b346103175760a051600319360112610317576015546001600160a01b03811633036131f05760a01c67ffffffffffffffff1680156110ff578060a051526012602052613024604060a05120548260a051526016602052604060a0512054906157e3565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690803b15610317576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b038516600484015260248301869052306044840152905191929091839160649183915af180156105fe576131d7575b5060a080518490526012602052516040812055601580547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff169055803b1561031757604051907f42966c680000000000000000000000000000000000000000000000000000000082528260048301528160248160a0519360a051905af180156105fe576131be575b50816131947fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e493617283565b506040805167ffffffffffffffff9092168252602082019290925290819081015b0390a160a05180f35b60a0516131ca91615418565b60a0516103175782613168565b60a0516131e391615418565b60a05161031757836130e0565b7f5fff6eee0000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346103175760c060031936011261031757613235615269565b5061323e615226565b613246615293565b5060843567ffffffffffffffff8111610317576132679036906004016152c4565b505060a43590600282101561031757611f2b916125bd91604435906159d8565b346103175760206003193601126103175760206132c267ffffffffffffffff6132ae61520f565b166000526007602052604060002054151590565b6040519015158152f35b346103175760206003193601126103175760206132ef6132ea61520f565b615996565b6001600160a01b0360405191168152f35b34610317576020600319360112610317577f44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d0917460206001600160a01b03613344615269565b61334c615e05565b168073ffffffffffffffffffffffffffffffffffffffff19600a541617600a55604051908152a160a05180f35b346103175760a0516003193601126103175760a051546001600160a01b03811633036133ff5773ffffffffffffffffffffffffffffffffffffffff196001549133828416176001551660a051556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060a05160a051a360a05180f35b7f02b543c60000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b346103175760a05160031936011261031757600454600554604080516001600160a01b039093168352602083019190915290f35b34610317576020600319360112610317577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b716860206001600160a01b036134a3615269565b6134ab615e05565b168073ffffffffffffffffffffffffffffffffffffffff196015541617601555604051908152a160a05180f35b346103175760406003193601126103175760043567ffffffffffffffff8111610317576135099036906004016154ac565b906024359067ffffffffffffffff821161031757366023830112156103175781600401359267ffffffffffffffff8411610317576024830192602436918660061b01011161031757613559615e05565b60a0515b8181106137a45750505060a0515b8281106135785760a05180f35b67ffffffffffffffff61358f610d4c838686615986565b1615801561376d575b801561374c575b613705576001600160a01b036135c160206135bb848787615986565b01615806565b1615610ccb57806136a06135dd60206135bb6001958888615986565b846001600160a01b0380868967ffffffffffffffff613621610d4c8a60405194613606866153c4565b60a051865287602087019b168b526040860199878b52615986565b1660a051526012602052604060a0512090518155019351161673ffffffffffffffffffffffffffffffffffffffff198354161782555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb60406136d1610d4c848888615986565b6001600160a01b036136e960206135bb878b8b615986565b67ffffffffffffffff845193168352166020820152a10161356b565b610d4c9067ffffffffffffffff9361371c93615986565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b5061376767ffffffffffffffff6132ae610d4c848787615986565b1561359f565b5067ffffffffffffffff613785610d4c838686615986565b1660a05152601260205260ff6001604060a05120015460a01c16613598565b67ffffffffffffffff6137bb610d4c838587615953565b1660a05152601260205260ff6001604060a05120015460a01c1615613896578067ffffffffffffffff6137f4610d4c6001948688615953565b1660a0515260126020527f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93604060a051205461383281601054615979565b60105567ffffffffffffffff61384c610d4c85888a615953565b60a0805191909216905260126020525160408120818155850155613874610d4c848789615953565b6040805167ffffffffffffffff9290921682526020820192909252a10161355d565b610d4c906138ad9267ffffffffffffffff94615953565b7f46f5f12b0000000000000000000000000000000000000000000000000000000060a0515216600452602460a051fd5b346103175760a0516003193601126103175760206001600160a01b03600a5416604051908152f35b34610317576020600319360112610317577f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b2346001600160a01b03613947615269565b61394f615e05565b6131b56011549183811673ffffffffffffffffffffffffffffffffffffffff1984161760115560405193849316839092916001600160a01b0360209181604085019616845216910152565b34610317576040600319360112610317576139b3615269565b6024356139be615e05565b6001600160a01b038216918215610ccb577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f02389273ffffffffffffffffffffffffffffffffffffffff196004541617600455816005556131b560405192839283602090939291936001600160a01b0360408201951681520152565b3461031757604060031936011261031757613a5161520f565b67ffffffffffffffff613a6261523d565b91613a6b615e05565b16908160a0515260126020526001604060a051200190815460ff8160a01c1615613af757825473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039283169081179093556040805191909216815260208101929092527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f919081908101611c99565b837f46f5f12b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b3461031757613b3536615560565b613b40929192615e05565b67ffffffffffffffff8216613b62816000526007602052604060002054151590565b15613b7d5750610d2b92613b77913691615457565b9061641e565b7f1e670e4b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b346103175760a0516003193601126103175760206104176158a8565b346103175760406003193601126103175760043561ffff8116908190036103175760243567ffffffffffffffff8111610317577f303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e91613c68613c30602093369060040161552f565b90613c39615e05565b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000600b541617600b55616117565b604051908152a160a05180f35b3461031757613c9d613ca5613c89366154dd565b9491613c96939193615e05565b36916155b9565b9236916155b9565b7f000000000000000000000000000000000000000000000000000000000000000015613dad5760a0515b8251811015613d3557806001600160a01b03613ced60019386615add565b5116613cf8816173aa565b613d04575b5001613ccf565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184613cfd565b5060a0515b8151811015610d2b57806001600160a01b03613d5860019385615add565b51168015613da757613d6981617244565b613d76575b505b01613d3a565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a183613d6e565b50613d70565b7f35f4a7b30000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461031757602060031936011261031757613df261520f565b613dfa615e05565b6015549067ffffffffffffffff8260a01c16613f175767ffffffffffffffff8116613e32816000526018602052604060002054151590565b613ee8578015613eb6577f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49927fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000060209460a01b16911617601555604051908152a160a05180f35b7fd9a9cd680000000000000000000000000000000000000000000000000000000060a0515260a051600452602460a051fd5b7f1c49a87b0000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b7f692bc1310000000000000000000000000000000000000000000000000000000060a05152600460a051fd5b3461031757604060031936011261031757613f5c61520f565b60243567ffffffffffffffff811161031757602091613f826132c292369060040161548e565b9061586b565b346103175760a05160031936011261031757613fa2615e05565b60155467ffffffffffffffff8160a01c169081156110ff577fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660155560a08051829052601660209081529051604080822091909155519182527f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda51891a160a05180f35b346103175760406003193601126103175760043567ffffffffffffffff811161031757806004016101006003198336030112610317576140636152a4565b90604051614070816153a8565b60a05190526140a161409761409261408b60c487018561581a565b3691615457565b61606d565b6064850135615f72565b9160848401906140b082615806565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036123ed57602485019277ffffffffffffffff00000000000000000000000000000000614109856157ce565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105fe5760a051916143c0575b506123925761418c6120b0856157ce565b614195846157ce565b906141ab60a4880192613f8261408b858561581a565b1561437957505061427b61290160446020977ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09561ffff67ffffffffffffffff961615156000146142e35785614200896157ce565b1660a05152600d8a5261421c604060a051208a61214c84615806565b7f0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce6158661424a6129018b6157ce565b604080516001600160a01b03929092168252602082018d90529190931692a25b019461427586615806565b506157ce565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081168252336020830152909216908201526060810185905292169180608081015b0390a2806040516142da816153a8565b52604051908152f35b50846142ee886157ce565b168060a0515260088a527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c89806143596002604060a05120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391616df2565b604080516001600160a01b039290921682526020820192909252a261426a565b614383925061581a565b61253d6040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191615bed565b6143d9915060203d6020116123e6576123d88183615418565b8761417b565b346103175760a0516003193601126103175760206001600160a01b0360115416604051908152f35b346103175760206003193601126103175760043567ffffffffffffffff811161031757806004019061010060031982360301126103175760405161444a816153a8565b60a051905261445c6064820135615e88565b906084810161446a81615806565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612c4a5750602481019277ffffffffffffffff000000000000000000000000000000006144c4856157ce565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105fe5760a05191614804575b50612392576145476120b0856157ce565b614550846157ce565b9061456660a4840192613f8261408b858561581a565b1561437957508291905067ffffffffffffffff614582856157ce565b168060a0515260086020526145cb6002604060a05120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016948591616df2565b604080516001600160a01b0385168152602081018690527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a267ffffffffffffffff614618856157ce565b1660a051526012602052604060a0512067ffffffffffffffff61463a866157ce565b1660a051526016602052604060a051205480156000146147c6575080548085116147925784614668916157e3565b90555b6044017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166146a182615806565b813b15610317576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b0387811660048501526024840189905293909316604483015251909283916064918391905af180156105fe57614779575b5067ffffffffffffffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0916142ca856147486129016020996157ce565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b60a05161478591615418565b60a051610317578461470a565b84907fa17e11d50000000000000000000000000000000000000000000000000000000060a05152600452602452604460a051fd5b8091508411611cd0575067ffffffffffffffff6147e2856157ce565b1660a051526016602052604060a051206147fd8482546157e3565b905561466b565b61481d915060203d6020116123e6576123d88183615418565b85614536565b346103175760a06003193601126103175761483c615269565b50614845615226565b60443567ffffffffffffffff81116103175760031960a091360301126103175761486d615293565b506084359067ffffffffffffffff82116103175761489867ffffffffffffffff9236906004016152c4565b50506040516148a681615376565b60a051815260a051602082015260a051604082015260a051606082015260a051608082015260a05160a082015260c060a0519101521660a05152600f60205260e0604060a05120604051906148fa82615376565b5461ffff818163ffffffff82169485815263ffffffff60208201818560201c1681528160408401818760401c168152816060860193818960601c16855260ff60c060808901988a8c60801c168a528a60a082019c60901c168c52019b60a01c1615158b526040519b8c52511660208b0152511660408901525116606087015251166080850152511660a083015251151560c0820152f35b346103175760206003193601126103175767ffffffffffffffff6149b361520f565b1660a051526012602052602060ff6001604060a05120015460a01c166040519015158152f35b34610317576040600319360112610317576149f261520f565b60243567ffffffffffffffff82168060a05152601260205260ff6001604060a05120015460a01c16158015614c2e575b611d0c57614a3d816000526018602052604060002054151590565b614bff578115610650576001600160a01b03614a5884615996565b1633036106205760a051526012602052604060a0512060ff600182015460a01c16600014614bea57614a8b828254615979565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f000000000000000000000000000000000000000000000000000000000000000090614afd906105148160848101610506565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b15610317576040517f47e7ef2400000000000000000000000000000000000000000000000000000000815260a080516001600160a01b039390931660048301526024820185905251909283916044918391905af180156105fe57614bd1575b506040805167ffffffffffffffff9093168352602083019190915233917f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf9181908101611c99565b60a051614bdd91615418565b60a0516103175782614b89565b50614bf781601054615979565b601055614a8e565b7f646972460000000000000000000000000000000000000000000000000000000060a05152600452602460a051fd5b508015614a22565b346103175760a05160031936011261031757602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103175760a0516003193601126103175760a051506040516013548082526020820190601360a05152602060a051209060a0515b818110614cc157611f2b856125bd81870382615418565b8254845260209093019260019283019201614caa565b34610317576020600319360112610317576020614cf2615269565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b346103175760a0516003193601126103175760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103175760a0516003193601126103175760408051611f2b91614d939082615418565b601d81527f53696c6f656455534443546f6b656e506f6f6c20312e372e302d64657600000060208201526040519182916020835260208301906152f2565b3461031757602060031936011261031757614dea615269565b614df2615e05565b614dfa6158a8565b9081614e065760a05180f35b60206001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992614eab6040517fa9059cbb0000000000000000000000000000000000000000000000000000000085820152614e8581610506898660248401602090939291936001600160a01b0360408201951681520152565b7f0000000000000000000000000000000000000000000000000000000000000000616bdc565b6040519485521692a28080610d2b565b34610317576020600319360112610317576004358015610650576001600160a01b03614ee860a051615996565b1633036106205760a0805180526012602052805160409020600181015490911c60ff1680156150465781545b808411611cd057501561503157614f2c8282546157e3565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610317576040517f69328dec00000000000000000000000000000000000000000000000000000000815260a080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af180156105fe5761501f575b506040519060a0515060a051825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a260a05180f35b60a05161502b91615418565b81614fe0565b5061503e816010546157e3565b601055614f2f565b601054614f14565b346103175760c06003193601126103175761506761520f565b61506f61523d565b50615078615253565b5060843561ffff811681036103175760a4359067ffffffffffffffff82116103175763ffffffff6150bf61ffff926080956150b8849636906004016152c4565b50506156eb565b9392959091604051968752166020860152166040840152166060820152f35b3461031757602060031936011261031757600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361031757817ff208a58f00000000000000000000000000000000000000000000000000000000602093149081156151e5575b81156151bb575b8115615191575b8115615167575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483615160565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150615159565b7ff57e5f940000000000000000000000000000000000000000000000000000000081149150615152565b7faff2afbf000000000000000000000000000000000000000000000000000000008114915061514b565b6004359067ffffffffffffffff82168203610d2657565b6024359067ffffffffffffffff82168203610d2657565b602435906001600160a01b0382168203610d2657565b606435906001600160a01b0382168203610d2657565b600435906001600160a01b0382168203610d2657565b35906001600160a01b0382168203610d2657565b6064359061ffff82168203610d2657565b6024359061ffff82168203610d2657565b359061ffff82168203610d2657565b9181601f84011215610d265782359167ffffffffffffffff8311610d265760208381860195010111610d2657565b919082519283825260005b84811061531e575050601f19601f8460006020809697860101520116010190565b806020809284010151828286010152016152fd565b602060408183019282815284518094520192019060005b8181106153575750505090565b82516001600160a01b031684526020938401939092019160010161534a565b60e0810190811067ffffffffffffffff82111761539257604052565b634e487b7160e01b600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761539257604052565b6060810190811067ffffffffffffffff82111761539257604052565b6040810190811067ffffffffffffffff82111761539257604052565b60a0810190811067ffffffffffffffff82111761539257604052565b90601f601f19910116810190811067ffffffffffffffff82111761539257604052565b67ffffffffffffffff811161539257601f01601f191660200190565b9291926154638261543b565b916154716040519384615418565b829481845281830111610d26578281602093846000960137010152565b9080601f83011215610d26578160206154a993359101615457565b90565b9181601f84011215610d265782359167ffffffffffffffff8311610d26576020808501948460051b010111610d2657565b6040600319820112610d265760043567ffffffffffffffff8111610d265781615508916004016154ac565b929092916024359067ffffffffffffffff8211610d265761552b916004016154ac565b9091565b9181601f84011215610d265782359167ffffffffffffffff8311610d265760208085019460e08502010111610d2657565b906040600319830112610d265760043567ffffffffffffffff81168103610d2657916024359067ffffffffffffffff8211610d265761552b916004016152c4565b67ffffffffffffffff81116153925760051b60200190565b9291906155c5816155a1565b936155d36040519586615418565b602085838152019160051b8101928311610d2657905b8282106155f557505050565b602080916156028461527f565b8152019101906155e9565b9080601f83011215610d26578160206154a9933591016155b9565b9181601f84011215610d265782359167ffffffffffffffff8311610d265760208085019460608502010111610d2657565b6154a991602061567283516040845260408401906152f2565b9201519060208184039101526152f2565b35908115158203610d2657565b35906001600160801b0382168203610d2657565b9190826060910312610d26576040516156bc816153c4565b60406156e68183956156cd81615683565b85526156db60208201615690565b602086015201615690565b910152565b67ffffffffffffffff16600052600f6020526040600020906040519161571083615376565b549263ffffffff84169384845263ffffffff8160201c169384602082015263ffffffff8260401c169384604083015263ffffffff8360601c169081606084015261ffff8460801c169283608082015260ff61ffff8660901c16958660a084015260a01c16159060c082159101526157b85761ffff161580159563ffffffff949392916157b05750945b156157a85750925b1693929190565b9050926157a1565b905094615799565b5050505092505050600090600090600090600090565b3567ffffffffffffffff81168103610d265790565b919082039182116157f057565b634e487b7160e01b600052601160045260246000fd5b356001600160a01b0381168103610d265790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610d26570180359067ffffffffffffffff8211610d2657602001918136038313610d2657565b9067ffffffffffffffff6154a992166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561594757600091615918575090565b90506020813d60201161593f575b8161593360209383615418565b81010312610d26575190565b3d9150615926565b6040513d6000823e3d90fd5b91908110156159635760051b0190565b634e487b7160e01b600052603260045260246000fd5b919082018092116157f057565b91908110156159635760061b0190565b67ffffffffffffffff16600052601260205260016040600020015460ff8160a01c166159cc57506001600160a01b036011541690565b6001600160a01b031690565b67ffffffffffffffff16600052600e6020526040600020916002811015615a1f57600114615a0e578160016154a9930190616682565b81600260036154a994019101616682565b634e487b7160e01b600052602160045260246000fd5b9190811015615963576060020190565b60405190615a52826153fc565b60006080838281528260208201528260408201528260608201520152565b90604051615a7d816153fc565b60806001829460ff81546001600160801b038116865263ffffffff81861c16602087015260a01c161515604085015201546001600160801b0381166060840152811c910152565b60405190615ad1826153e0565b60606020838281520152565b80518210156159635760209160051b010190565b90600182811c92168015615b21575b6020831014615b0b57565b634e487b7160e01b600052602260045260246000fd5b91607f1691615b00565b9060405191826000825492615b3f84615af1565b8084529360018116908115615bad5750600114615b66575b50615b6492500383615418565b565b90506000929192526020600020906000915b818310615b91575050906020615b649282010138615b57565b6020919350806001915483858901015201910190918492615b78565b60209350615b649592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138615b57565b601f8260209493601f19938186528686013760008582860101520116010190565b67ffffffffffffffff1660005260086020526154a96004604060002001615b2b565b91908110156159635760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610d26570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610d26570180359067ffffffffffffffff8211610d2657602001918160051b36038313610d2657565b818102929181159184041417156157f057565b818110615ce2575050565b60008155600101615cd7565b9160209082815201919060005b818110615d085750505090565b9091926020806001926001600160a01b03615d228861527f565b168152019401929101615cfb565b91908110156159635760081b0190565b3561ffff81168103610d265790565b3563ffffffff81168103610d265790565b358015158103610d265790565b359063ffffffff82168203610d2657565b67ffffffffffffffff16615d9f816000526007602052604060002054151590565b15615dd85780600052601260205260ff60016040600020015460a01c16615dc7575060105490565b600052601260205260406000205490565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6001600160a01b03600154163303615e1957565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b9060ff8091169116039060ff82116157f057565b60ff16604d81116157f057600a0a90565b8115615e72570490565b634e487b7160e01b600052601260045260246000fd5b7f000000000000000000000000000000000000000000000000000000000000000060ff81169081600614615f6d5781600611615f42576006615ec991615e43565b90604d60ff8316118015615f27575b615ef0575090615eea6154a992615e57565b90615cc4565b90507fa9cb113d00000000000000000000000000000000000000000000000000000000600052600660045260245260445260646000fd5b50615f3182615e57565b8015615e7257600019048311615ed8565b615f4d906006615e43565b90604d60ff831611615ef0575090615f676154a992615e57565b90615e68565b505090565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff81169282841461604e5782841161602a5790615fb791615e43565b91604d60ff841611801561600f575b615fd957505090615eea6154a992615e57565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061601983615e57565b8015615e7257600019048411615fc6565b61603391615e43565b91604d60ff841611615fd957505090615f676154a992615e57565b5050505090565b90816020910312610d2657518015158103610d265790565b805180156160dd5760200361609f578051602082810191830183900312610d2657519060ff821161609f575060ff1690565b61253d906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484015260248301906152f2565b50507f000000000000000000000000000000000000000000000000000000000000000090565b356001600160801b0381168103610d265790565b9160005b828110156164185760e0810284016000616134826157ce565b9067ffffffffffffffff8216616157816000526007602052604060002054151590565b156163ec575050816161dd61624d926161e36020600197960161618261617d36836156a4565b616acb565b6161dd6161a38467ffffffffffffffff16600052600c602052604060002090565b91825463ffffffff8160801c161590816163d7575b816163c8575b816163b6575b816163a7575b5080616398575b616320575b36906156a4565b90616fe6565b61621260808401916161f861617d36856156a4565b67ffffffffffffffff16600052600d602052604060002090565b92835463ffffffff8160801c1615908161630b575b816162fc575b816162ea575b816162db575b50806162cc575b616253575b5036906156a4565b0161611b565b61626760a06001600160801b039201616103565b845473ffffffff000000000000000000000000000000004260801b1673ffffffffffffffffffffffffffffffffffffffff19909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538616245565b506162d682615d60565b616240565b60ff915060a01c161538616239565b6001600160801b038116159150616233565b8589015460801c15915061622d565b858901546001600160801b0316159150616227565b6001600160801b0361633460408901616103565b845473ffffffff000000000000000000000000000000004260801b1673ffffffffffffffffffffffffffffffffffffffff19909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783556161d6565b506163a281615d60565b6161d1565b60ff915060a01c1615386161ca565b6001600160801b0381161591506161c4565b848c015460801c1591506161be565b848c01546001600160801b03161591506161b8565b602492507f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b908051156166045767ffffffffffffffff8151602083012092169182600052600860205261645381600560406000200161732e565b156165c05760005260096020526040600020815167ffffffffffffffff8111615392576164808254615af1565b601f811161658e575b506020601f821160011461650457916164de827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936164f4956000916164f9575b506000198260011b9260031b1c19161790565b90556040519182916020835260208301906152f2565b0390a2565b9050840151386164cb565b601f1982169083600052806000209160005b8181106165765750926164f49492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061655d575b5050811b019055611f17565b85015160001960f88460031b161c191690553880616551565b9192602060018192868a015181550194019201616516565b6165ba90836000526020600020601f840160051c81019160208510610c6e57601f0160051c0190615cd7565b38616489565b509061253d6040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401526040602484015260448301906152f2565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b906040519182815491828252602082019060005260206000209260005b818110616660575050615b6492500383615418565b84546001600160a01b031683526001948501948794506020909301920161664b565b61668b9061662e565b91600554801515918261676b575b50506166a3575090565b6166ac9061662e565b908151806166ba5750905090565b6166c5908251615979565b92601f196166eb6166d5866155a1565b956166e36040519788615418565b8087526155a1565b0136602086013760005b825181101561672657806001600160a01b0361671360019386615add565b511661671f8288615add565b52016166f5565b509160005b815181101561676657806001600160a01b0361674960019385615add565b511661675f616759838751615979565b88615add565b520161672b565b505050565b101590503880616699565b67ffffffffffffffff166000818152600760205260409020549092919015616866579161686360e092616838856167cd7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97616acb565b8460005260086020526167e4816040600020616fe6565b6167ed83616acb565b846000526008602052616807836002604060002001616fe6565b60405194855260208501906001600160801b0360408092805115158552826020820151166020860152015116910152565b60808301906001600160801b0360408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61689c615a45565b506001600160801b036060820151166001600160801b0380835116916168e760208501936168e16168d463ffffffff875116426157e3565b8560808901511690615cc4565b90615979565b8082101561690057505b16825263ffffffff4216905290565b90506168f1565b67ffffffffffffffff61691c602083016157ce565b16600052600f60205260406000206040519061693782615376565b549063ffffffff8216815263ffffffff8260201c16602082015263ffffffff8260401c16604082015263ffffffff8260601c16606082015261ffff8260801c169081608082015260ff61ffff8460901c16938460a084015260a01c16159060c082159101526169e257919261ffff92908316156169db57505b1680156169d3576127106169cc60606154a99401359283615cc4565b04906157e3565b506060013590565b90506169b0565b505060609150013590565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000166020820152602081526154a9604082615418565b805160005b818110616a3957505050565b600181018082116157f0575b828110616a555750600101616a2d565b6001600160a01b03616a678386615add565b51166001600160a01b03616a7b8387615add565b511614616a8a57600101616a45565b6001600160a01b03616a9c8386615add565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115616b50576001600160801b036040820151166001600160801b0360208301511610616af65750565b606490616b4e604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565bfd5b6001600160801b0360408201511615801590616bc6575b616b6e5750565b606490616b4e604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906001600160801b0360408092805115158552826020820151166020860152015116910152565b506001600160801b036020820151161515616b67565b6001600160a01b03616c5e911691604092600080855193616bfd8786615418565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d15616d07573d91616c428361543b565b92616c4f87519485615418565b83523d6000602085013e617606565b80519081616c6b57505050565b602080616c7c938301019101616055565b15616c845750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091617606565b906040519182815491828252602082019060005260206000209260005b818110616d41575050615b6492500383615418565b8454835260019485019487945060209093019201616d2c565b67ffffffffffffffff16616d7b816000526007602052604060002054151590565b15616dc5575033600052601460205260406000205415616d9757565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9182549060ff8260a01c16158015616fde575b616fd8576001600160801b0382169160018501908154616e3863ffffffff6001600160801b0383169360801c16426157e3565b9081616f3a575b5050848110616efb5750838310616e90575050616e656001600160801b039283926157e3565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91616e9f81856157e3565b926000198101908082116157f057616ec2616ec7926001600160a01b0396615979565b615e68565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611616fae57616f55926168e19160801c90615cc4565b80841015616fa95750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880616e3f565b616f60565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215616e05565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c199161710d606092805461702363ffffffff8260801c16426157e3565b9081617143575b50506001600160801b03600181602086015116928281541680851060001461713b57508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556170ca8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166001600160801b031692909217910155565b61686360405180926001600160801b0360408092805115158552826020820151166020860152015116910152565b838091617051565b6001600160801b039161716f8392836171686001880154948286169560801c90615cc4565b9116615979565b808210156171e357505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff92909116929092161673ffffffffffffffffffffffffffffffffffffffff19909116174260801b73ffffffff0000000000000000000000000000000016178155388061702a565b9050617179565b80548210156159635760005260206000200190600090565b805490680100000000000000008210156153925781617229916001617240940181556171ea565b81939154906000199060031b92831b921b19161790565b9055565b8060005260036020526040600020541560001461727d57617266816002617202565b600254906000526003602052604060002055600190565b50600090565b8060005260186020526040600020541560001461727d576172a5816017617202565b601754906000526018602052604060002055600190565b8060005260146020526040600020541560001461727d576172de816013617202565b601354906000526014602052604060002055600190565b8060005260076020526040600020541560001461727d57617317816006617202565b600654906000526007602052604060002055600190565b6000828152600182016020526040902054617365578061735083600193617202565b80549260005201602052604060002055600190565b5050600090565b8054801561739457600019019061738382826171ea565b60001982549160031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b60008181526003602052604090205480156173655760001981018181116157f0576002549060001982019182116157f057818103617404575b5050506173f0600261736c565b600052600360205260006040812055600190565b6174266174156172299360026171ea565b90549060031b1c92839260026171ea565b905560005260036020526040600020553880806173e3565b60008181526007602052604090205480156173655760001981018181116157f0576006549060001982019182116157f057818103617498575b505050617484600661736c565b600052600760205260006040812055600190565b6174ba6174a96172299360066171ea565b90549060031b1c92839260066171ea565b90556000526007602052604060002055388080617477565b906001820191816000528260205260406000205490811515600014617569576000198201918083116157f057815460001981019081116157f05783816175209503617532575b50505061736c565b60005260205260006040812055600190565b61755261754261722993866171ea565b90549060031b1c928392866171ea565b905560005284602052604060002055388080617518565b50505050600090565b60008181526014602052604090205480156173655760001981018181116157f0576013549060001982019182116157f0578082036175cc575b5050506175b8601361736c565b600052601460205260006040812055600190565b6175ee6175dd6172299360136171ea565b90549060031b1c92839260136171ea565b905560005260146020526040600020553880806175ab565b91929015617681575081511561761a575090565b3b156176235790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156176945750805190602001fd5b61253d906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906152f256fea164736f6c634300081a000a",
}

var SiloedUSDCTokenPoolABI = SiloedUSDCTokenPoolMetaData.ABI

var SiloedUSDCTokenPoolBin = SiloedUSDCTokenPoolMetaData.Bin

func DeploySiloedUSDCTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address, lockBox common.Address) (common.Address, *types.Transaction, *SiloedUSDCTokenPool, error) {
	parsed, err := SiloedUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SiloedUSDCTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router, lockBox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SiloedUSDCTokenPool{address: address, abi: *parsed, SiloedUSDCTokenPoolCaller: SiloedUSDCTokenPoolCaller{contract: contract}, SiloedUSDCTokenPoolTransactor: SiloedUSDCTokenPoolTransactor{contract: contract}, SiloedUSDCTokenPoolFilterer: SiloedUSDCTokenPoolFilterer{contract: contract}}, nil
}

type SiloedUSDCTokenPool struct {
	address common.Address
	abi     abi.ABI
	SiloedUSDCTokenPoolCaller
	SiloedUSDCTokenPoolTransactor
	SiloedUSDCTokenPoolFilterer
}

type SiloedUSDCTokenPoolCaller struct {
	contract *bind.BoundContract
}

type SiloedUSDCTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type SiloedUSDCTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type SiloedUSDCTokenPoolSession struct {
	Contract     *SiloedUSDCTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type SiloedUSDCTokenPoolCallerSession struct {
	Contract *SiloedUSDCTokenPoolCaller
	CallOpts bind.CallOpts
}

type SiloedUSDCTokenPoolTransactorSession struct {
	Contract     *SiloedUSDCTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type SiloedUSDCTokenPoolRaw struct {
	Contract *SiloedUSDCTokenPool
}

type SiloedUSDCTokenPoolCallerRaw struct {
	Contract *SiloedUSDCTokenPoolCaller
}

type SiloedUSDCTokenPoolTransactorRaw struct {
	Contract *SiloedUSDCTokenPoolTransactor
}

func NewSiloedUSDCTokenPool(address common.Address, backend bind.ContractBackend) (*SiloedUSDCTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(SiloedUSDCTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindSiloedUSDCTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPool{address: address, abi: abi, SiloedUSDCTokenPoolCaller: SiloedUSDCTokenPoolCaller{contract: contract}, SiloedUSDCTokenPoolTransactor: SiloedUSDCTokenPoolTransactor{contract: contract}, SiloedUSDCTokenPoolFilterer: SiloedUSDCTokenPoolFilterer{contract: contract}}, nil
}

func NewSiloedUSDCTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*SiloedUSDCTokenPoolCaller, error) {
	contract, err := bindSiloedUSDCTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCaller{contract: contract}, nil
}

func NewSiloedUSDCTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SiloedUSDCTokenPoolTransactor, error) {
	contract, err := bindSiloedUSDCTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolTransactor{contract: contract}, nil
}

func NewSiloedUSDCTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SiloedUSDCTokenPoolFilterer, error) {
	contract, err := bindSiloedUSDCTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolFilterer{contract: contract}, nil
}

func bindSiloedUSDCTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SiloedUSDCTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloedUSDCTokenPool.Contract.SiloedUSDCTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SiloedUSDCTokenPoolTransactor.contract.Transfer(opts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SiloedUSDCTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiloedUSDCTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.contract.Transfer(opts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetAccumulatedFees(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetAccumulatedFees(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAllAuthorizedCallers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllAuthorizedCallers(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAllAuthorizedCallers() ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllAuthorizedCallers(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllowList(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllowList(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllowListEnabled(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _SiloedUSDCTokenPool.Contract.GetAllowListEnabled(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getAvailableTokens", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetAvailableTokens(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetAvailableTokens(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetAvailableTokens(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetAvailableTokens(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getChainRebalancer", remoteChainSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetChainRebalancer(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetChainRebalancer(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetChainRebalancer(remoteChainSelector uint64) (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetChainRebalancer(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetConfiguredMinBlockConfirmation(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getConfiguredMinBlockConfirmation")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetConfiguredMinBlockConfirmation() (uint16, error) {
	return _SiloedUSDCTokenPool.Contract.GetConfiguredMinBlockConfirmation(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetConfiguredMinBlockConfirmation() (uint16, error) {
	return _SiloedUSDCTokenPool.Contract.GetConfiguredMinBlockConfirmation(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getCurrentCustomBlockConfirmationRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentCustomBlockConfirmationRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetCurrentCustomBlockConfirmationRateLimiterState(remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentCustomBlockConfirmationRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getCurrentProposedCCTPChainMigration")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetCurrentProposedCCTPChainMigration() (uint64, error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentProposedCCTPChainMigration(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetCurrentProposedCCTPChainMigration() (uint64, error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentProposedCCTPChainMigration(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getCurrentRateLimiterState", remoteChainSelector)

	outstruct := new(GetCurrentRateLimiterState)
	if err != nil {
		return *outstruct, err
	}

	outstruct.OutboundRateLimiterState = *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)
	outstruct.InboundRateLimiterState = *abi.ConvertType(out[1], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return *outstruct, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetCurrentRateLimiterState(remoteChainSelector uint64) (GetCurrentRateLimiterState,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetDynamicConfig(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetDynamicConfig(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getExcludedTokensByChain", remoteChainSelector)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetExcludedTokensByChain(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetExcludedTokensByChain(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetExcludedTokensByChain(remoteChainSelector uint64) (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetExcludedTokensByChain(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetFee(opts *bind.CallOpts, destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getFee", destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)

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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetFee(destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetFee(&_SiloedUSDCTokenPool.CallOpts, destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetFee(destChainSelector uint64, arg1 common.Address, arg2 *big.Int, arg3 common.Address, blockConfirmationRequested uint16, arg5 []byte) (GetFee,

	error) {
	return _SiloedUSDCTokenPool.Contract.GetFee(&_SiloedUSDCTokenPool.CallOpts, destChainSelector, arg1, arg2, arg3, blockConfirmationRequested, arg5)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRateLimitAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRateLimitAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRateLimitAdmin() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRateLimitAdmin(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRateLimitAdmin() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRateLimitAdmin(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRebalancer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRebalancer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRebalancer() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRebalancer(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRebalancer() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRebalancer(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRemotePools(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRemotePools(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRemoteToken(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRemoteToken(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredCCVs(&_SiloedUSDCTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredCCVs(&_SiloedUSDCTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRmnProxy(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRmnProxy(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _SiloedUSDCTokenPool.Contract.GetSupportedChains(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _SiloedUSDCTokenPool.Contract.GetSupportedChains(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetToken() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetToken(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetToken(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _SiloedUSDCTokenPool.Contract.GetTokenDecimals(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _SiloedUSDCTokenPool.Contract.GetTokenDecimals(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _SiloedUSDCTokenPool.Contract.GetTokenTransferFeeConfig(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _SiloedUSDCTokenPool.Contract.GetTokenTransferFeeConfig(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetUnsiloedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getUnsiloedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetUnsiloedLiquidity() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetUnsiloedLiquidity(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetUnsiloedLiquidity() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.GetUnsiloedLiquidity(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsRemotePool(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsRemotePool(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) IsSiloed(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "isSiloed", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) IsSiloed(remoteChainSelector uint64) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSiloed(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) IsSiloed(remoteChainSelector uint64) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSiloed(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSupportedChain(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSupportedChain(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSupportedToken(&_SiloedUSDCTokenPool.CallOpts, token)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.IsSupportedToken(&_SiloedUSDCTokenPool.CallOpts, token)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) Owner() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.Owner(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) Owner() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.Owner(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.SupportsInterface(&_SiloedUSDCTokenPool.CallOpts, interfaceId)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.SupportsInterface(&_SiloedUSDCTokenPool.CallOpts, interfaceId)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) TypeAndVersion() (string, error) {
	return _SiloedUSDCTokenPool.Contract.TypeAndVersion(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _SiloedUSDCTokenPool.Contract.TypeAndVersion(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "acceptOwnership")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AcceptOwnership(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AcceptOwnership(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AddRemotePool(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AddRemotePool(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyAllowListUpdates(&_SiloedUSDCTokenPool.TransactOpts, removes, adds)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyAllowListUpdates(&_SiloedUSDCTokenPool.TransactOpts, removes, adds)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyAuthorizedCallerUpdates", authorizedCallerArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_SiloedUSDCTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyAuthorizedCallerUpdates(authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyAuthorizedCallerUpdates(&_SiloedUSDCTokenPool.TransactOpts, authorizedCallerArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyCCVConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, ccvConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyCCVConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, ccvConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyChainUpdates(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyChainUpdates(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyCustomBlockConfirmationConfigUpdates", minBlockConfirmation, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyCustomBlockConfirmationConfigUpdates(minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyCustomBlockConfirmationConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, minBlockConfirmation, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "burnLockedUSDC")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) BurnLockedUSDC() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.BurnLockedUSDC(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) BurnLockedUSDC() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.BurnLockedUSDC(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "cancelExistingCCTPMigrationProposal")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) CancelExistingCCTPMigrationProposal() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.CancelExistingCCTPMigrationProposal(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) CancelExistingCCTPMigrationProposal() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.CancelExistingCCTPMigrationProposal(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ExcludeTokensFromBurn(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "excludeTokensFromBurn", remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ExcludeTokensFromBurn(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ExcludeTokensFromBurn(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ExcludeTokensFromBurn(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ExcludeTokensFromBurn(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn0(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.LockOrBurn0(&_SiloedUSDCTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ProposeCCTPMigration(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "proposeCCTPMigration", remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ProposeCCTPMigration(remoteChainSelector uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProposeCCTPMigration(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ProposeCCTPMigration(remoteChainSelector uint64) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProposeCCTPMigration(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "provideLiquidity", amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProvideLiquidity(&_SiloedUSDCTokenPool.TransactOpts, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProvideLiquidity(&_SiloedUSDCTokenPool.TransactOpts, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "provideSiloedLiquidity", remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ProvideSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProvideSiloedLiquidity(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ProvideSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ProvideSiloedLiquidity(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint0(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ReleaseOrMint0(&_SiloedUSDCTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RemoveRemotePool(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RemoveRemotePool(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetChainRateLimiterConfig(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetChainRateLimiterConfig(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetChainRateLimiterConfigs(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetChainRateLimiterConfigs(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setCircleMigratorAddress", migrator)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetCircleMigratorAddress(migrator common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetCircleMigratorAddress(&_SiloedUSDCTokenPool.TransactOpts, migrator)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetCircleMigratorAddress(migrator common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetCircleMigratorAddress(&_SiloedUSDCTokenPool.TransactOpts, migrator)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetCustomBlockConfirmationRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setCustomBlockConfirmationRateLimitConfig", rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_SiloedUSDCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetCustomBlockConfirmationRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetCustomBlockConfirmationRateLimitConfig(&_SiloedUSDCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetDynamicConfig(&_SiloedUSDCTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetDynamicConfig(&_SiloedUSDCTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetRateLimitAdmin(opts *bind.TransactOpts, rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setRateLimitAdmin", rateLimitAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRateLimitAdmin(&_SiloedUSDCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetRateLimitAdmin(rateLimitAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRateLimitAdmin(&_SiloedUSDCTokenPool.TransactOpts, rateLimitAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setRebalancer", newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRebalancer(&_SiloedUSDCTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRebalancer(&_SiloedUSDCTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setSiloRebalancer", remoteChainSelector, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetSiloRebalancer(remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetSiloRebalancer(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetSiloRebalancer(remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetSiloRebalancer(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "transferOwnership", to)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.TransferOwnership(&_SiloedUSDCTokenPool.TransactOpts, to)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.TransferOwnership(&_SiloedUSDCTokenPool.TransactOpts, to)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) UpdateSiloDesignations(opts *bind.TransactOpts, removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "updateSiloDesignations", removes, adds)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) UpdateSiloDesignations(removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.UpdateSiloDesignations(&_SiloedUSDCTokenPool.TransactOpts, removes, adds)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) UpdateSiloDesignations(removes []uint64, adds []SiloedLockReleaseTokenPoolSiloConfigUpdate) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.UpdateSiloDesignations(&_SiloedUSDCTokenPool.TransactOpts, removes, adds)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawFees(&_SiloedUSDCTokenPool.TransactOpts, recipient)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawFees(&_SiloedUSDCTokenPool.TransactOpts, recipient)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawLiquidity(&_SiloedUSDCTokenPool.TransactOpts, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawLiquidity(&_SiloedUSDCTokenPool.TransactOpts, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) WithdrawSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "withdrawSiloedLiquidity", remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) WithdrawSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawSiloedLiquidity(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) WithdrawSiloedLiquidity(remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.WithdrawSiloedLiquidity(&_SiloedUSDCTokenPool.TransactOpts, remoteChainSelector, amount)
}

type SiloedUSDCTokenPoolAllowListAddIterator struct {
	Event *SiloedUSDCTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolAllowListAdd)
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
		it.Event = new(SiloedUSDCTokenPoolAllowListAdd)
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

func (it *SiloedUSDCTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolAllowListAddIterator{contract: _SiloedUSDCTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolAllowListAdd)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*SiloedUSDCTokenPoolAllowListAdd, error) {
	event := new(SiloedUSDCTokenPoolAllowListAdd)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolAllowListRemoveIterator struct {
	Event *SiloedUSDCTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolAllowListRemove)
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
		it.Event = new(SiloedUSDCTokenPoolAllowListRemove)
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

func (it *SiloedUSDCTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolAllowListRemoveIterator{contract: _SiloedUSDCTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolAllowListRemove)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*SiloedUSDCTokenPoolAllowListRemove, error) {
	event := new(SiloedUSDCTokenPoolAllowListRemove)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolAuthorizedCallerAddedIterator struct {
	Event *SiloedUSDCTokenPoolAuthorizedCallerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolAuthorizedCallerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolAuthorizedCallerAdded)
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
		it.Event = new(SiloedUSDCTokenPoolAuthorizedCallerAdded)
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

func (it *SiloedUSDCTokenPoolAuthorizedCallerAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolAuthorizedCallerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolAuthorizedCallerAdded struct {
	Caller common.Address
	Raw    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAuthorizedCallerAddedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolAuthorizedCallerAddedIterator{contract: _SiloedUSDCTokenPool.contract, event: "AuthorizedCallerAdded", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAuthorizedCallerAdded) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "AuthorizedCallerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolAuthorizedCallerAdded)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseAuthorizedCallerAdded(log types.Log) (*SiloedUSDCTokenPoolAuthorizedCallerAdded, error) {
	event := new(SiloedUSDCTokenPoolAuthorizedCallerAdded)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator struct {
	Event *SiloedUSDCTokenPoolAuthorizedCallerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolAuthorizedCallerRemoved)
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
		it.Event = new(SiloedUSDCTokenPoolAuthorizedCallerRemoved)
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

func (it *SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolAuthorizedCallerRemoved struct {
	Caller common.Address
	Raw    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator{contract: _SiloedUSDCTokenPool.contract, event: "AuthorizedCallerRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAuthorizedCallerRemoved) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "AuthorizedCallerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolAuthorizedCallerRemoved)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseAuthorizedCallerRemoved(log types.Log) (*SiloedUSDCTokenPoolAuthorizedCallerRemoved, error) {
	event := new(SiloedUSDCTokenPoolAuthorizedCallerRemoved)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "AuthorizedCallerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCCTPMigrationCancelledIterator struct {
	Event *SiloedUSDCTokenPoolCCTPMigrationCancelled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCCTPMigrationCancelledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCCTPMigrationCancelled)
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
		it.Event = new(SiloedUSDCTokenPoolCCTPMigrationCancelled)
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

func (it *SiloedUSDCTokenPoolCCTPMigrationCancelledIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCCTPMigrationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCCTPMigrationCancelled struct {
	ExistingProposalSelector uint64
	Raw                      types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCCTPMigrationCancelled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationCancelledIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationCancelled")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCCTPMigrationCancelledIterator{contract: _SiloedUSDCTokenPool.contract, event: "CCTPMigrationCancelled", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCCTPMigrationCancelled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationCancelled) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCCTPMigrationCancelled)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationCancelled", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCCTPMigrationCancelled(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationCancelled, error) {
	event := new(SiloedUSDCTokenPoolCCTPMigrationCancelled)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCCTPMigrationExecutedIterator struct {
	Event *SiloedUSDCTokenPoolCCTPMigrationExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCCTPMigrationExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCCTPMigrationExecuted)
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
		it.Event = new(SiloedUSDCTokenPoolCCTPMigrationExecuted)
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

func (it *SiloedUSDCTokenPoolCCTPMigrationExecutedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCCTPMigrationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCCTPMigrationExecuted struct {
	RemoteChainSelector uint64
	USDCBurned          *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCCTPMigrationExecuted(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationExecutedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationExecuted")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCCTPMigrationExecutedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CCTPMigrationExecuted", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCCTPMigrationExecuted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationExecuted) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCCTPMigrationExecuted)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationExecuted", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCCTPMigrationExecuted(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationExecuted, error) {
	event := new(SiloedUSDCTokenPoolCCTPMigrationExecuted)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCCTPMigrationProposedIterator struct {
	Event *SiloedUSDCTokenPoolCCTPMigrationProposed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCCTPMigrationProposedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCCTPMigrationProposed)
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
		it.Event = new(SiloedUSDCTokenPoolCCTPMigrationProposed)
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

func (it *SiloedUSDCTokenPoolCCTPMigrationProposedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCCTPMigrationProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCCTPMigrationProposed struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCCTPMigrationProposed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationProposedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CCTPMigrationProposed")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCCTPMigrationProposedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CCTPMigrationProposed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCCTPMigrationProposed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationProposed) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CCTPMigrationProposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCCTPMigrationProposed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationProposed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCCTPMigrationProposed(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationProposed, error) {
	event := new(SiloedUSDCTokenPoolCCTPMigrationProposed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCTPMigrationProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCCVConfigUpdatedIterator struct {
	Event *SiloedUSDCTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCCVConfigUpdated)
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
		it.Event = new(SiloedUSDCTokenPoolCCVConfigUpdated)
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

func (it *SiloedUSDCTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCCVConfigUpdatedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCCVConfigUpdated)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolCCVConfigUpdated, error) {
	event := new(SiloedUSDCTokenPoolCCVConfigUpdated)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolChainAddedIterator struct {
	Event *SiloedUSDCTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainAdded)
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
		it.Event = new(SiloedUSDCTokenPoolChainAdded)
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

func (it *SiloedUSDCTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainAddedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainAddedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainAdded)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainAdded(log types.Log) (*SiloedUSDCTokenPoolChainAdded, error) {
	event := new(SiloedUSDCTokenPoolChainAdded)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolChainConfiguredIterator struct {
	Event *SiloedUSDCTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainConfigured)
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
		it.Event = new(SiloedUSDCTokenPoolChainConfigured)
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

func (it *SiloedUSDCTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainConfiguredIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainConfigured)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainConfigured(log types.Log) (*SiloedUSDCTokenPoolChainConfigured, error) {
	event := new(SiloedUSDCTokenPoolChainConfigured)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolChainRemovedIterator struct {
	Event *SiloedUSDCTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainRemoved)
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
		it.Event = new(SiloedUSDCTokenPoolChainRemoved)
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

func (it *SiloedUSDCTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainRemovedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainRemoved)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainRemoved(log types.Log) (*SiloedUSDCTokenPoolChainRemoved, error) {
	event := new(SiloedUSDCTokenPoolChainRemoved)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolChainSiloedIterator struct {
	Event *SiloedUSDCTokenPoolChainSiloed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainSiloedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainSiloed)
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
		it.Event = new(SiloedUSDCTokenPoolChainSiloed)
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

func (it *SiloedUSDCTokenPoolChainSiloedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainSiloedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainSiloed struct {
	RemoteChainSelector uint64
	Rebalancer          common.Address
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainSiloed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainSiloedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainSiloed")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainSiloedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainSiloed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainSiloed) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainSiloed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainSiloed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainSiloed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainSiloed(log types.Log) (*SiloedUSDCTokenPoolChainSiloed, error) {
	event := new(SiloedUSDCTokenPoolChainSiloed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainSiloed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolChainUnsiloedIterator struct {
	Event *SiloedUSDCTokenPoolChainUnsiloed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolChainUnsiloedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolChainUnsiloed)
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
		it.Event = new(SiloedUSDCTokenPoolChainUnsiloed)
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

func (it *SiloedUSDCTokenPoolChainUnsiloedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolChainUnsiloedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolChainUnsiloed struct {
	RemoteChainSelector uint64
	AmountUnsiloed      *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainUnsiloedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ChainUnsiloed")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolChainUnsiloedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ChainUnsiloed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainUnsiloed) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ChainUnsiloed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolChainUnsiloed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainUnsiloed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseChainUnsiloed(log types.Log) (*SiloedUSDCTokenPoolChainUnsiloed, error) {
	event := new(SiloedUSDCTokenPoolChainUnsiloed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ChainUnsiloed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCircleMigratorAddressSetIterator struct {
	Event *SiloedUSDCTokenPoolCircleMigratorAddressSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCircleMigratorAddressSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCircleMigratorAddressSet)
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
		it.Event = new(SiloedUSDCTokenPoolCircleMigratorAddressSet)
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

func (it *SiloedUSDCTokenPoolCircleMigratorAddressSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCircleMigratorAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCircleMigratorAddressSet struct {
	MigratorAddress common.Address
	Raw             types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCircleMigratorAddressSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCircleMigratorAddressSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CircleMigratorAddressSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCircleMigratorAddressSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "CircleMigratorAddressSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCircleMigratorAddressSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCircleMigratorAddressSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CircleMigratorAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCircleMigratorAddressSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CircleMigratorAddressSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCircleMigratorAddressSet(log types.Log) (*SiloedUSDCTokenPoolCircleMigratorAddressSet, error) {
	event := new(SiloedUSDCTokenPoolCircleMigratorAddressSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CircleMigratorAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolConfigChangedIterator struct {
	Event *SiloedUSDCTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolConfigChanged)
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
		it.Event = new(SiloedUSDCTokenPoolConfigChanged)
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

func (it *SiloedUSDCTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolConfigChangedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolConfigChanged)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseConfigChanged(log types.Log) (*SiloedUSDCTokenPoolConfigChanged, error) {
	event := new(SiloedUSDCTokenPoolConfigChanged)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomBlockConfirmationInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomBlockConfirmationOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationUpdatedIterator struct {
	Event *SiloedUSDCTokenPoolCustomBlockConfirmationUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationUpdated)
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
		it.Event = new(SiloedUSDCTokenPoolCustomBlockConfirmationUpdated)
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

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomBlockConfirmationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomBlockConfirmationUpdated struct {
	MinBlockConfirmation uint16
	Raw                  types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCustomBlockConfirmationUpdatedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomBlockConfirmationUpdatedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomBlockConfirmationUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomBlockConfirmationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomBlockConfirmationUpdated)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomBlockConfirmationUpdated(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationUpdated, error) {
	event := new(SiloedUSDCTokenPoolCustomBlockConfirmationUpdated)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomBlockConfirmationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolDynamicConfigSetIterator struct {
	Event *SiloedUSDCTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolDynamicConfigSet)
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
		it.Event = new(SiloedUSDCTokenPoolDynamicConfigSet)
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

func (it *SiloedUSDCTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolDynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolDynamicConfigSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolDynamicConfigSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*SiloedUSDCTokenPoolDynamicConfigSet, error) {
	event := new(SiloedUSDCTokenPoolDynamicConfigSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolInboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolInboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolInboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolInboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolInboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolInboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolInboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolLiquidityAddedIterator struct {
	Event *SiloedUSDCTokenPoolLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolLiquidityAdded)
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
		it.Event = new(SiloedUSDCTokenPoolLiquidityAdded)
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

func (it *SiloedUSDCTokenPoolLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolLiquidityAdded struct {
	RemoteChainSelector uint64
	Provider            common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*SiloedUSDCTokenPoolLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolLiquidityAddedIterator{contract: _SiloedUSDCTokenPool.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLiquidityAdded, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolLiquidityAdded)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseLiquidityAdded(log types.Log) (*SiloedUSDCTokenPoolLiquidityAdded, error) {
	event := new(SiloedUSDCTokenPoolLiquidityAdded)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolLiquidityRemovedIterator struct {
	Event *SiloedUSDCTokenPoolLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolLiquidityRemoved)
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
		it.Event = new(SiloedUSDCTokenPoolLiquidityRemoved)
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

func (it *SiloedUSDCTokenPoolLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolLiquidityRemoved struct {
	RemoteChainSelector uint64
	Remover             common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, remover []common.Address) (*SiloedUSDCTokenPoolLiquidityRemovedIterator, error) {

	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "LiquidityRemoved", removerRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolLiquidityRemovedIterator{contract: _SiloedUSDCTokenPool.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLiquidityRemoved, remover []common.Address) (event.Subscription, error) {

	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "LiquidityRemoved", removerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolLiquidityRemoved)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseLiquidityRemoved(log types.Log) (*SiloedUSDCTokenPoolLiquidityRemoved, error) {
	event := new(SiloedUSDCTokenPoolLiquidityRemoved)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolLockedOrBurnedIterator struct {
	Event *SiloedUSDCTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolLockedOrBurned)
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
		it.Event = new(SiloedUSDCTokenPoolLockedOrBurned)
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

func (it *SiloedUSDCTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolLockedOrBurnedIterator{contract: _SiloedUSDCTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolLockedOrBurned)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*SiloedUSDCTokenPoolLockedOrBurned, error) {
	event := new(SiloedUSDCTokenPoolLockedOrBurned)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolOutboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolOutboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolOutboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolOutboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolOutboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolOwnershipTransferRequestedIterator struct {
	Event *SiloedUSDCTokenPoolOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolOwnershipTransferRequested)
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
		it.Event = new(SiloedUSDCTokenPoolOwnershipTransferRequested)
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

func (it *SiloedUSDCTokenPoolOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolOwnershipTransferRequestedIterator{contract: _SiloedUSDCTokenPool.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolOwnershipTransferRequested)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseOwnershipTransferRequested(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferRequested, error) {
	event := new(SiloedUSDCTokenPoolOwnershipTransferRequested)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolOwnershipTransferredIterator struct {
	Event *SiloedUSDCTokenPoolOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolOwnershipTransferred)
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
		it.Event = new(SiloedUSDCTokenPoolOwnershipTransferred)
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

func (it *SiloedUSDCTokenPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolOwnershipTransferredIterator{contract: _SiloedUSDCTokenPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolOwnershipTransferred)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseOwnershipTransferred(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferred, error) {
	event := new(SiloedUSDCTokenPoolOwnershipTransferred)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolPoolFeeWithdrawnIterator struct {
	Event *SiloedUSDCTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolPoolFeeWithdrawn)
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
		it.Event = new(SiloedUSDCTokenPoolPoolFeeWithdrawn)
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

func (it *SiloedUSDCTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*SiloedUSDCTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolPoolFeeWithdrawnIterator{contract: _SiloedUSDCTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolPoolFeeWithdrawn)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*SiloedUSDCTokenPoolPoolFeeWithdrawn, error) {
	event := new(SiloedUSDCTokenPoolPoolFeeWithdrawn)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRateLimitAdminSetIterator struct {
	Event *SiloedUSDCTokenPoolRateLimitAdminSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRateLimitAdminSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRateLimitAdminSet)
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
		it.Event = new(SiloedUSDCTokenPoolRateLimitAdminSet)
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

func (it *SiloedUSDCTokenPoolRateLimitAdminSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRateLimitAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRateLimitAdminSet struct {
	RateLimitAdmin common.Address
	Raw            types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRateLimitAdminSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolRateLimitAdminSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRateLimitAdminSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "RateLimitAdminSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRateLimitAdminSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RateLimitAdminSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRateLimitAdminSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRateLimitAdminSet(log types.Log) (*SiloedUSDCTokenPoolRateLimitAdminSet, error) {
	event := new(SiloedUSDCTokenPoolRateLimitAdminSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RateLimitAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolReleasedOrMintedIterator struct {
	Event *SiloedUSDCTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolReleasedOrMinted)
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
		it.Event = new(SiloedUSDCTokenPoolReleasedOrMinted)
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

func (it *SiloedUSDCTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolReleasedOrMintedIterator{contract: _SiloedUSDCTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolReleasedOrMinted)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*SiloedUSDCTokenPoolReleasedOrMinted, error) {
	event := new(SiloedUSDCTokenPoolReleasedOrMinted)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRemotePoolAddedIterator struct {
	Event *SiloedUSDCTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRemotePoolAdded)
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
		it.Event = new(SiloedUSDCTokenPoolRemotePoolAdded)
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

func (it *SiloedUSDCTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRemotePoolAddedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRemotePoolAdded)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*SiloedUSDCTokenPoolRemotePoolAdded, error) {
	event := new(SiloedUSDCTokenPoolRemotePoolAdded)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRemotePoolRemovedIterator struct {
	Event *SiloedUSDCTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRemotePoolRemoved)
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
		it.Event = new(SiloedUSDCTokenPoolRemotePoolRemoved)
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

func (it *SiloedUSDCTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRemotePoolRemovedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRemotePoolRemoved)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*SiloedUSDCTokenPoolRemotePoolRemoved, error) {
	event := new(SiloedUSDCTokenPoolRemotePoolRemoved)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolSiloRebalancerSetIterator struct {
	Event *SiloedUSDCTokenPoolSiloRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolSiloRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolSiloRebalancerSet)
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
		it.Event = new(SiloedUSDCTokenPoolSiloRebalancerSet)
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

func (it *SiloedUSDCTokenPoolSiloRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolSiloRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolSiloRebalancerSet struct {
	RemoteChainSelector uint64
	OldRebalancer       common.Address
	NewRebalancer       common.Address
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterSiloRebalancerSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolSiloRebalancerSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "SiloRebalancerSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolSiloRebalancerSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "SiloRebalancerSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchSiloRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolSiloRebalancerSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "SiloRebalancerSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolSiloRebalancerSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "SiloRebalancerSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseSiloRebalancerSet(log types.Log) (*SiloedUSDCTokenPoolSiloRebalancerSet, error) {
	event := new(SiloedUSDCTokenPoolSiloRebalancerSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "SiloRebalancerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted)
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
		it.Event = new(SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted)
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

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _SiloedUSDCTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated)
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
		it.Event = new(SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated)
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

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _SiloedUSDCTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolTokensExcludedFromBurnIterator struct {
	Event *SiloedUSDCTokenPoolTokensExcludedFromBurn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolTokensExcludedFromBurnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolTokensExcludedFromBurn)
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
		it.Event = new(SiloedUSDCTokenPoolTokensExcludedFromBurn)
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

func (it *SiloedUSDCTokenPoolTokensExcludedFromBurnIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolTokensExcludedFromBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolTokensExcludedFromBurn struct {
	RemoteChainSelector          uint64
	Amount                       *big.Int
	BurnableAmountAfterExclusion *big.Int
	Raw                          types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterTokensExcludedFromBurn(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolTokensExcludedFromBurnIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "TokensExcludedFromBurn", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolTokensExcludedFromBurnIterator{contract: _SiloedUSDCTokenPool.contract, event: "TokensExcludedFromBurn", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchTokensExcludedFromBurn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokensExcludedFromBurn, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "TokensExcludedFromBurn", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolTokensExcludedFromBurn)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokensExcludedFromBurn", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseTokensExcludedFromBurn(log types.Log) (*SiloedUSDCTokenPoolTokensExcludedFromBurn, error) {
	event := new(SiloedUSDCTokenPoolTokensExcludedFromBurn)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "TokensExcludedFromBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator struct {
	Event *SiloedUSDCTokenPoolUnsiloedRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolUnsiloedRebalancerSet)
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
		it.Event = new(SiloedUSDCTokenPoolUnsiloedRebalancerSet)
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

func (it *SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolUnsiloedRebalancerSet struct {
	OldRebalancer common.Address
	NewRebalancer common.Address
	Raw           types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterUnsiloedRebalancerSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "UnsiloedRebalancerSet")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator{contract: _SiloedUSDCTokenPool.contract, event: "UnsiloedRebalancerSet", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchUnsiloedRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolUnsiloedRebalancerSet) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "UnsiloedRebalancerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolUnsiloedRebalancerSet)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "UnsiloedRebalancerSet", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseUnsiloedRebalancerSet(log types.Log) (*SiloedUSDCTokenPoolUnsiloedRebalancerSet, error) {
	event := new(SiloedUSDCTokenPoolUnsiloedRebalancerSet)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "UnsiloedRebalancerSet", log); err != nil {
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

func (SiloedUSDCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (SiloedUSDCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (SiloedUSDCTokenPoolAuthorizedCallerAdded) Topic() common.Hash {
	return common.HexToHash("0xeb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef")
}

func (SiloedUSDCTokenPoolAuthorizedCallerRemoved) Topic() common.Hash {
	return common.HexToHash("0xc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580")
}

func (SiloedUSDCTokenPoolCCTPMigrationCancelled) Topic() common.Hash {
	return common.HexToHash("0x375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda518")
}

func (SiloedUSDCTokenPoolCCTPMigrationExecuted) Topic() common.Hash {
	return common.HexToHash("0xdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e4")
}

func (SiloedUSDCTokenPoolCCTPMigrationProposed) Topic() common.Hash {
	return common.HexToHash("0x20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49")
}

func (SiloedUSDCTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (SiloedUSDCTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (SiloedUSDCTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (SiloedUSDCTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (SiloedUSDCTokenPoolChainSiloed) Topic() common.Hash {
	return common.HexToHash("0x180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb")
}

func (SiloedUSDCTokenPoolChainUnsiloed) Topic() common.Hash {
	return common.HexToHash("0x7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93")
}

func (SiloedUSDCTokenPoolCircleMigratorAddressSet) Topic() common.Hash {
	return common.HexToHash("0x084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b7168")
}

func (SiloedUSDCTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x0ca6b6e55c811717f7cba5bdd97184f0a2d5aa61e5ea64818392da5cef6ce615")
}

func (SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x61b04c91cb711a3783cc99d0de694ecd7ead5c96a0a87f5c89d436ac20979328")
}

func (SiloedUSDCTokenPoolCustomBlockConfirmationUpdated) Topic() common.Hash {
	return common.HexToHash("0x303439e67d1363a21c3ecd1158164e797c51eced31b0351ec0e1f7afaf97779e")
}

func (SiloedUSDCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (SiloedUSDCTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (SiloedUSDCTokenPoolLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0x569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf")
}

func (SiloedUSDCTokenPoolLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0x58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e")
}

func (SiloedUSDCTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (SiloedUSDCTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (SiloedUSDCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SiloedUSDCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (SiloedUSDCTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (SiloedUSDCTokenPoolRateLimitAdminSet) Topic() common.Hash {
	return common.HexToHash("0x44676b5284b809a22248eba0da87391d79098be38bb03154be88a58bf4d09174")
}

func (SiloedUSDCTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (SiloedUSDCTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (SiloedUSDCTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (SiloedUSDCTokenPoolSiloRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f")
}

func (SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xfae1e296719dac5269c3886fb5002bb29bf17ae403060c6eb063a55abaaa1041")
}

func (SiloedUSDCTokenPoolTokensExcludedFromBurn) Topic() common.Hash {
	return common.HexToHash("0xe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220")
}

func (SiloedUSDCTokenPoolUnsiloedRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b234")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPool) Address() common.Address {
	return _SiloedUSDCTokenPool.address
}

type SiloedUSDCTokenPoolInterface interface {
	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetConfiguredMinBlockConfirmation(opts *bind.CallOpts) (uint16, error)

	GetCurrentCustomBlockConfirmationRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentCustomBlockConfirmationRateLimiterState,

		error)

	GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error)

	GetCurrentRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (GetCurrentRateLimiterState,

		error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

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

	ApplyAuthorizedCallerUpdates(opts *bind.TransactOpts, authorizedCallerArgs AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error)

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyCustomBlockConfirmationConfigUpdates(opts *bind.TransactOpts, minBlockConfirmation uint16, rateLimitConfigArgs []TokenPoolCustomBlockConfirmationRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error)

	ExcludeTokensFromBurn(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ProposeCCTPMigration(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error)

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

	FilterAllowListAdd(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*SiloedUSDCTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*SiloedUSDCTokenPoolAllowListRemove, error)

	FilterAuthorizedCallerAdded(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAuthorizedCallerAddedIterator, error)

	WatchAuthorizedCallerAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAuthorizedCallerAdded) (event.Subscription, error)

	ParseAuthorizedCallerAdded(log types.Log) (*SiloedUSDCTokenPoolAuthorizedCallerAdded, error)

	FilterAuthorizedCallerRemoved(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolAuthorizedCallerRemovedIterator, error)

	WatchAuthorizedCallerRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolAuthorizedCallerRemoved) (event.Subscription, error)

	ParseAuthorizedCallerRemoved(log types.Log) (*SiloedUSDCTokenPoolAuthorizedCallerRemoved, error)

	FilterCCTPMigrationCancelled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationCancelledIterator, error)

	WatchCCTPMigrationCancelled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationCancelled) (event.Subscription, error)

	ParseCCTPMigrationCancelled(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationCancelled, error)

	FilterCCTPMigrationExecuted(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationExecutedIterator, error)

	WatchCCTPMigrationExecuted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationExecuted) (event.Subscription, error)

	ParseCCTPMigrationExecuted(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationExecuted, error)

	FilterCCTPMigrationProposed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCCTPMigrationProposedIterator, error)

	WatchCCTPMigrationProposed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCTPMigrationProposed) (event.Subscription, error)

	ParseCCTPMigrationProposed(log types.Log) (*SiloedUSDCTokenPoolCCTPMigrationProposed, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*SiloedUSDCTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*SiloedUSDCTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*SiloedUSDCTokenPoolChainRemoved, error)

	FilterChainSiloed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainSiloedIterator, error)

	WatchChainSiloed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainSiloed) (event.Subscription, error)

	ParseChainSiloed(log types.Log) (*SiloedUSDCTokenPoolChainSiloed, error)

	FilterChainUnsiloed(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolChainUnsiloedIterator, error)

	WatchChainUnsiloed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolChainUnsiloed) (event.Subscription, error)

	ParseChainUnsiloed(log types.Log) (*SiloedUSDCTokenPoolChainUnsiloed, error)

	FilterCircleMigratorAddressSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCircleMigratorAddressSetIterator, error)

	WatchCircleMigratorAddressSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCircleMigratorAddressSet) (event.Subscription, error)

	ParseCircleMigratorAddressSet(log types.Log) (*SiloedUSDCTokenPoolCircleMigratorAddressSet, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*SiloedUSDCTokenPoolConfigChanged, error)

	FilterCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationInboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumedIterator, error)

	WatchCustomBlockConfirmationOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomBlockConfirmationOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationOutboundRateLimitConsumed, error)

	FilterCustomBlockConfirmationUpdated(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolCustomBlockConfirmationUpdatedIterator, error)

	WatchCustomBlockConfirmationUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomBlockConfirmationUpdated) (event.Subscription, error)

	ParseCustomBlockConfirmationUpdated(log types.Log) (*SiloedUSDCTokenPoolCustomBlockConfirmationUpdated, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*SiloedUSDCTokenPoolDynamicConfigSet, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolInboundRateLimitConsumed, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*SiloedUSDCTokenPoolLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLiquidityAdded, provider []common.Address) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*SiloedUSDCTokenPoolLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, remover []common.Address) (*SiloedUSDCTokenPoolLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLiquidityRemoved, remover []common.Address) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*SiloedUSDCTokenPoolLiquidityRemoved, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*SiloedUSDCTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolOutboundRateLimitConsumed, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferred, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*SiloedUSDCTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*SiloedUSDCTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolRateLimitAdminSetIterator, error)

	WatchRateLimitAdminSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRateLimitAdminSet) (event.Subscription, error)

	ParseRateLimitAdminSet(log types.Log) (*SiloedUSDCTokenPoolRateLimitAdminSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*SiloedUSDCTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*SiloedUSDCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*SiloedUSDCTokenPoolRemotePoolRemoved, error)

	FilterSiloRebalancerSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolSiloRebalancerSetIterator, error)

	WatchSiloRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolSiloRebalancerSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseSiloRebalancerSet(log types.Log) (*SiloedUSDCTokenPoolSiloRebalancerSet, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated, error)

	FilterTokensExcludedFromBurn(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolTokensExcludedFromBurnIterator, error)

	WatchTokensExcludedFromBurn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolTokensExcludedFromBurn, remoteChainSelector []uint64) (event.Subscription, error)

	ParseTokensExcludedFromBurn(log types.Log) (*SiloedUSDCTokenPoolTokensExcludedFromBurn, error)

	FilterUnsiloedRebalancerSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolUnsiloedRebalancerSetIterator, error)

	WatchUnsiloedRebalancerSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolUnsiloedRebalancerSet) (event.Subscription, error)

	ParseUnsiloedRebalancerSet(log types.Log) (*SiloedUSDCTokenPoolUnsiloedRebalancerSet, error)

	Address() common.Address
}
