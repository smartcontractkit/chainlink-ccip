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

type TokenPoolCustomFinalityRateLimitConfigArgs struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

type TokenPoolTokenTransferFeeConfigArgs struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
}

var SiloedUSDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AUTHORIZED_CALLER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMITER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enumIPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"structSiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyMigrated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCircle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenLockingNotAllowedAfterMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610120806040523461044b57618eba803803809161001d8285610717565b8339810160c08282031261044b5781516001600160a01b0381169290919083830361044b5761004e6020820161073a565b60408201516001600160401b03811161044b5782019280601f8501121561044b578351936001600160401b038511610450578460051b9060208201956100976040519788610717565b865260208087019282010192831161044b57602001905b8282106106ff575050506100c460608301610748565b6100dc60a06100d560808601610748565b9401610748565b9433156106e957600180546001600160d01b031690556002546001600160a01b0381166106d8576001600160a01b0319163390811760025561011d90610786565b50861580156106c7575b80156106b6575b61051c5760805260c05260405163313ce56760e01b8152602081600481895afa6000918161067a575b5061064f575b5060a052600580546001600160a01b0319166001600160a01b03929092169190911790558051151560e081905261052d575b506001600160a01b0316801561051c57604051636eb1769f60e11b815230600482015260248101829052602081604481865afa908115610510576000916104de575b506104735760405191602083019263095ea7b360e01b8452826024820152600019604482015260448152610206606482610717565b6000806040958651936102198886610717565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020860152519082865af13d15610466573d906001600160401b03821161045057855161028a94909261027b601f8201601f191660200185610717565b83523d6000602085013e610984565b8051806103cf575b505061010052516184459081610a5582396080518181816104c901528181611bd401528181611fe2015281816121d801528181612339015281816128d801528181612a7501528181612cc70152818161303f0152818161442d015281816145f3015281816146a3015281816147e70152818161491501528181614f680152818161551a01528181615568015281816156cd015281816159150152616221015260a0518181816154a301528181616cb601528181616da101528181616f0d0152617461015260c0518181816111880152818161207001528181612966015281816144bb0152614876015260e051818181611035015281816120b5015281816129ab0152613fc701526101005181818161051401528181611b7901528181612b7701528181613015015281816149e201528181614fa301526158ba0152f35b816020918101031261044b576020015180159081150361044b576103f4573880610292565b815162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b9161028a92606091610984565b60405162461bcd60e51b815260206004820152603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90506020813d602011610508575b816104f960209383610717565b8101031261044b5751386101d1565b3d91506104ec565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b906020906040519061053f8383610717565b60008252600036813760e0511561063e5760005b82518110156105ba576001906001600160a01b03610571828661075c565b51168561057d8261082c565b61058a575b505001610553565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13885610582565b5092905060005b8151811015610635576001906001600160a01b036105df828561075c565b5116801561062f57846105f18261092a565b6105ff575b50505b016105c1565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138846105f6565b506105f9565b5050503861018f565b6335f4a7b360e01b60005260046000fd5b60ff1660ff8216818103610663575061015d565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116106ae575b8161069660209383610717565b8101031261044b576106a79061073a565b9038610157565b3d9150610689565b506001600160a01b0382161561012e565b506001600160a01b03841615610127565b631fe1e13d60e11b60005260046000fd5b636116401160e11b600052600060045260246000fd5b6020809161070c84610748565b8152019101906100ae565b601f909101601f19168101906001600160401b0382119082101761045057604052565b519060ff8216820361044b57565b51906001600160a01b038216820361044b57565b80518210156107705760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b0381166000908152600080516020618e9a833981519152602052604090205460ff1661080e576001600160a01b03166000818152600080516020618e9a83398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b80548210156107705760005260206000200190600090565b600081815260046020526040902054801561092357600019810181811161090d5760035460001981019190821161090d578181036108bc575b50505060035480156108a65760001901610880816003610814565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6108f56108cd6108de936003610814565b90549060031b1c9283926003610814565b819391549060031b91821b91600019901b19161790565b90556000526004602052604060002055388080610865565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260046020526040600020541560001461080e57600354680100000000000000008110156104505761096b6108de8260018594016003556003610814565b9055600354906000526004602052604060002055600190565b919290156109e65750815115610998575090565b3b156109a15790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b8251909150156109f95750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610a3c5750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610a1a56fe60a080604052600436101561001357600080fd5b60006080526080513560e01c90816301ffc9a7146159ee57508063022d63fb146159cf5780630a861f2a1461583d5780630aa6220b146157785780630bd7c46d14615701578063164e68de146155ee578063181f5a771461558c57806321df0da714615547578063240028e8146154f4578063248a9ca3146154c757806324f65ee7146154885780632a10097b146152055780632c286daf146150d25780632d4a148f14614e755780632f2ff15d14614e2f57806331238ffc14614de757806336568abe14614ca657806337b1924714614b955780633907753714614779578063432a6ba314614751578063489a68f2146143975780634ad01f0b146142d15780634c5ef0ed1461428c57806350d1a35a146140f957806354c8a4f314613f6d5780635df45a3714613f5157806362ddd3c414613ea3578063634e93da14613d7d578063649a5ec714613b5d5780636600f92c14613a3a578063698c2c66146139685780636cfd1553146138a95780636d9d216c14613470578063714bf907146133c25780637437ff9f1461338e578063791e5a1014613352578063804ba5a9146132e757806384ef8ffc14612f805780638632d5cc146132b35780638926f54f1461326e57806389720a62146132035780638a5e52bb14612fa85780638da5cb5b14612f8057806391d1485414612f2f578063962d402014612d925780639a4575b91461287d578063a1eda53c14612818578063a217fddf146127fa578063a42a7b8b146126a3578063a7cd63b714612635578063acfecf91146124ea578063af0e58b9146124cb578063af58d59f1461247f578063b1c71c6514611f65578063b794658014611f2d578063c4bffe2b14611e11578063c75eea9c14611d66578063cc8463c814611d3a578063cd306a6c14611d0e578063ce3c752814611ac0578063cefc1429146119bc578063cf6eefb714611968578063cf7401f3146117b1578063d547741f14611733578063d602b9fd146116ad578063d966866b1461123a578063da90a9f3146111ac578063dc0bd97114611167578063de814c571461105a578063e0351e131461101c578063e58d80c714610fad578063e8a1da171461069e578063eb521a4c14610413578063f1e73399146103e8578063f573388e146103ac5763f65a88861461036957600080fd5b346103a65760206003193601126103a65767ffffffffffffffff61038b615cbb565b16608051526014602052602060406080512054604051908152f35b60805180fd5b346103a6576080516003193601126103a65760206040517ff12fb6eaf1f045883c82d7d192627f7a36a50ce00c45e305919895908135a8a88152f35b346103a65760206003193601126103a657602061040b610406615cbb565b6166c4565b604051908152f35b346103a65760206003193601126103a65760808051805260166020525160409020546004359061066c578015610640576001600160a01b036104566080516162ad565b1633036106105760808051805260126020525160409020600181015460a01c60ff16156105fb57610488828254616290565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f00000000000000000000000000000000000000000000000000000000000000009061050a9061050481608481015b03601f198101835282615dd2565b82617811565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b156103a6576040517f47e7ef24000000000000000000000000000000000000000000000000000000008152608080516001600160a01b039390931660048301526024820185905251909283916044918391905af180156105ee576105d5575b506040519060805150608051825260208201527f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf60403392a260805180f35b6080516105e191615dd2565b6080516103a65781610596565b6040513d608051823e3d90fd5b5061060881601054616290565b60105561048b565b7f8e4a23d60000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b7fa90c0d1900000000000000000000000000000000000000000000000000000000608051526004608051fd5b7f6469724600000000000000000000000000000000000000000000000000000000608051526080516004526024608051fd5b346103a6576106ac36615e66565b91608093919351608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815790608051905b828210610dbf575050506080519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610db957600581901b830135858112156103a6578301610120813603126103a6576040519461075686615db6565b813567ffffffffffffffff81168103610db4578652602082013567ffffffffffffffff81116103a65782019436601f870112156103a657853561079881616179565b966107a66040519889615dd2565b81885260208089019260051b820101903682116103a65760208101925b828410610d85575050505060208701958652604083013567ffffffffffffffff81116103a6576107f69036908501615e48565b916040880192835261082061080e3660608701615fc1565b9460608a0195865260c0369101615fc1565b95608089019687526108328551617612565b61083c8751617612565b83515115610d595761085867ffffffffffffffff8a5116617f74565b15610d1e5767ffffffffffffffff89511660805152600960205260406080512061099c86516fffffffffffffffffffffffffffffffff604082015116906109576fffffffffffffffffffffffffffffffff602083015116915115158360806040516108c281615db6565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b610ac288516fffffffffffffffffffffffffffffffff60408201511690610a7d6fffffffffffffffffffffffffffffffff602083015116915115158360806040516109e681615db6565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff8211610ced57610ae583546163a2565b601f8111610cae575b506020906001601f841114610c4457918091610b219360805192610c39575b50506000198260011b9260031b1c19161790565b90555b6080515b88518051821015610b5d5790610b57600192610b508367ffffffffffffffff8f51169261638e565b5190616f2f565b01610b28565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2939199975095610c2b67ffffffffffffffff6001979694985116925193519151610bf7610bc260405196879687526101006020880152610100870190615bf6565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610724565b015190508e80610b0d565b90601f1983169184608051528160805120926080515b818110610c965750908460019594939210610c7d575b505050811b019055610b24565b015160001960f88460031b161c191690558d8080610c70565b92936020600181928786015181550195019301610c5a565b610cdd908460805152602060805120601f850160051c81019160208610610ce3575b601f0160051c019061666b565b8d610aee565b9091508190610cd0565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b67ffffffffffffffff8951167f1d5ad3c500000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f14c880ca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b833567ffffffffffffffff81116103a657602091610da98392833691870101615e48565b8152019301926107c3565b600080fd5b60805180f35b9092919367ffffffffffffffff610ddf610dda86888661608b565b616047565b1692610dea84618150565b15610f515783608051526009602052610e0a600560406080512001618023565b926080515b8451811015610e495760019086608051526009602052610e42600560406080512001610e3b838961638e565b5190618203565b5001610e0f565b5093909491959250806080515260096020526005604060805120608051815560805160018201556080516002820155608051600382015560048101610e8e81546163a2565b80610f01575b505001805490608051815581610edd575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10190916106e7565b60805152602060805120908101905b81811015610ea5576080518155600101610eec565b601f8111600114610f1a575060805190555b8880610e94565b610f3a9082608051526001601f6020608051209201861c8201910161666b565b608080518290525160208120918190559055610f13565b837f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f2b5c74de00000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103a65760206003193601126103a657610fc6615bb6565b7f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af608051526080516020526001600160a01b036040608051209116600052602052602060ff604060002054166040519015158152f35b346103a6576080516003193601126103a65760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346103a65760406003193601126103a657611073615cbb565b602435608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815767ffffffffffffffff8060135460a01c16921680920361113b5760407fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220918360805152601460205281608051206110ff828254616290565b90558360805152601260205261112a826080512054856080515260146020528360805120549061609b565b82519182526020820152a260805180f35b7fa94cb98800000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103a6576080516003193601126103a65760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103a65760206003193601126103a6576111c5615bb6565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f81576020816112237fd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b9361758c565b506001600160a01b0360405191168152a160805180f35b346103a65760206003193601126103a65760043567ffffffffffffffff81116103a65761126b903690600401615c37565b90608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815790608051915b8183106112b05760805180f35b6112be610dda8484846165c4565b6112d66112cc8585856165c4565b6020810190616604565b90916112f06112e68787876165c4565b6040810190616604565b906113096112ff8989896165c4565b6060810190616604565b90916113236113198b8b8b6165c4565b6080810190616604565b949097611339611334368a84616191565b617495565b611347611334368486616191565b611355611334368688616191565b61136361133436888c616191565b60405161136f81615d17565b61137a368a84616191565b8152611387368486616191565b6020820152611397368688616191565b60408201526113a736888c616191565b606082015267ffffffffffffffff881660805152600e602052604060805120815180519067ffffffffffffffff8211610ced57680100000000000000008211610ced57602090835483855580841061168e575b500182608051526020608051206080515b8381106116715750505050602082015180519067ffffffffffffffff8211610ced57680100000000000000008211610ced57602090600184015483600186015580841061164f575b500160018301608051526020608051206080515b8381106116325750505050604082015180519067ffffffffffffffff8211610ced57680100000000000000008211610ced576020906002840154836002860155808410611610575b500160028301608051526020608051206080515b8381106115f3575050505060036060919e9c9d9e019101519081519167ffffffffffffffff8311610ced57680100000000000000008311610ced5760209082548484558085106115d4575b500190608051526020608051206080515b8381106115b7575050505061159c6080956115ac9561158e7fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a9661158067ffffffffffffffff976040519d8d8f9e8f9081520191616682565b918b830360208d0152616682565b9188830360408a0152616682565b9285840360608701521696616682565b0390a20191906112a3565b60019060206001600160a01b03855116940193818401550161151f565b6115ed908460805152858460805120918201910161666b565b3861150e565b60019060206001600160a01b0385511694019381840155016114c3565b61162c906002860160805152848460805120918201910161666b565b386114af565b60019060206001600160a01b038551169401938184015501611467565b61166b906001860160805152848460805120918201910161666b565b38611453565b60019060206001600160a01b03855116940193818401550161140b565b6116a7908560805152848460805120918201910161666b565b386113fa565b346103a6576080516003193601126103a6576116c761674b565b600180547fffffffffffff0000000000000000000000000000000000000000000000000000811690915560a01c65ffffffffffff166117065760805180f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109608051608051a1610db9565b346103a65760406003193601126103a65760043561174f615bcc565b8115611785578161177961177461177e94600052600060205260016040600020015490565b6167b7565b6175b6565b5060805180f35b7f3fc3c27a00000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103a65760e06003193601126103a6576117ca615cbb565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01126103a65760405161180081615d7e565b60243580151581036103a65781526044356fffffffffffffffffffffffffffffffff811681036103a65760208201526064356fffffffffffffffffffffffffffffffff811681036103a65760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c01126103a6576040519061188782615d7e565b60843580151581036103a657825260a4356fffffffffffffffffffffffffffffffff811681036103a657602083015260c4356fffffffffffffffffffffffffffffffff811681036103a65760408301527f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af608051526080516020526040608051206001600160a01b03331660005260205260ff604060002054161580611935575b61061057610db9926172a5565b50608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615611928565b346103a6576080516003193601126103a657604065ffffffffffff6119a36001549065ffffffffffff6001600160a01b0383169260a01c1690565b6001600160a01b03849392935193168352166020820152f35b346103a6576080516003193601126103a6576001546001600160a01b03163303611a905760015460a081901c65ffffffffffff16906001600160a01b031681158015611a86575b611a5657611a2590611a1f6001600160a01b0360025416617538565b50616841565b507fffffffffffff000000000000000000000000000000000000000000000000000060015416600155608051608051f35b507f19ca5ebb00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b5042821015611a03565b7fc22c80220000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b346103a65760406003193601126103a657611ad9615cbb565b60243567ffffffffffffffff82168060805152601260205260ff6001604060805120015460a01c16158015611d06575b611cd7578115610640576001600160a01b03611b24846162ad565b1633036106105760805152601260205260406080512060ff600182015460a01c166080515080600014611ccf5781545b808411611c9b575015611c8657611b6c82825461609b565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b156103a6576040517f69328dec000000000000000000000000000000000000000000000000000000008152608080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af180156105ee57611c6d575b506040805167ffffffffffffffff9093168352602083019190915233917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e91819081015b0390a260805180f35b608051611c7991615dd2565b6080516103a65782611c20565b50611c938160105461609b565b601055611b6f565b83907fa17e11d500000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b601054611b54565b7f46f5f12b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b508015611b09565b346103a6576080516003193601126103a657602067ffffffffffffffff60135460a01c16604051908152f35b346103a6576080516003193601126103a6576020611d5661658b565b65ffffffffffff60405191168152f35b346103a65760206003193601126103a65767ffffffffffffffff611d88615cbb565b611d906164d8565b5016608051526009602052611e0d611db4611daf604060805120616503565b6173d5565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346103a6576080516003193601126103a657608051506040516007548082528160208101600760805152602060805120926080515b818110611f14575050611e5b92500382615dd2565b805190611e80611e6a83616179565b92611e786040519485615dd2565b808452616179565b90601f196020840192013683376080515b8151811015611ec3578067ffffffffffffffff611eb06001938561638e565b5116611ebc828761638e565b5201611e91565b505090604051918291602083019060208452518091526040830191906080515b818110611ef1575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611ee3565b8454835260019485019486945060209093019201611e46565b346103a65760206003193601126103a657611e0d611f51611f4c615cbb565b616569565b604051918291602083526020830190615bf6565b346103a65760606003193601126103a65760043567ffffffffffffffff81116103a65760a060031982360301126103a657611f9e615c68565b9060443567ffffffffffffffff81116103a657611fbf903690600401615e48565b50611fc8616375565b506084810190611fd7826160d7565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361243e57602481019077ffffffffffffffff0000000000000000000000000000000061203083616047565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105ee576080519161240f575b506123e3576120b3604482016160d7565b7f0000000000000000000000000000000000000000000000000000000000000000612390575b50606461ffff916120f16120ec85616047565b617b77565b013593169283151592838094612381575b156122e05761ffff600b5416948581106122ac5750612270945061215561214561212b85616047565b67ffffffffffffffff16600052600c602052604060002090565b8361214f846160d7565b91617c32565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff61219161218b86616047565b936160d7565b604080516001600160a01b03929092168252602082018690529190931692a25b91829061227a575b50611f4c816121ca61223f93616047565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2616047565b9061224861745a565b6040519261225584615d9a565b83526020830152604051928392604084526040840190615f6d565b9060208301520390f35b61223f9192506122a4611f4c9161271061229d61ffff600b5460101c1683616658565b049061609b565b9291506121b9565b85907fe08f03ef00000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b50612270935067ffffffffffffffff6122f883616047565b16806080515260096020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894482806123616040608051206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391617c32565b604080516001600160a01b039290921682526020820192909252a26121b1565b5061ffff600b54161515612102565b6001600160a01b03166123b0816000526004602052604060002054151590565b6120d9577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f53ad11d800000000000000000000000000000000000000000000000000000000608051526004608051fd5b612431915060203d602011612437575b6124298183615dd2565b810190616e81565b856120a2565b503d61241f565b6001600160a01b0361244f836160d7565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b346103a65760206003193601126103a65767ffffffffffffffff6124a1615cbb565b6124a96164d8565b5016608051526009602052611e0d611db4611daf600260406080512001616503565b346103a6576080516003193601126103a6576020601054604051908152f35b346103a6576124f836615eb8565b91608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815767ffffffffffffffff169061254c826000526008602052604060002054151590565b15612605578160805152600960205261257f600560406080512001612572368685615e11565b6020815191012090618203565b156125be577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d769192611c646040519283926020845260208401916164b7565b612601906040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916164b7565b0390fd5b507f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103a6576080516003193601126103a657608051506040516003548082526020820190600360805152602060805120906080515b81811061268d57611e0d8561268181870382615dd2565b60405191829182615ef9565b825484526020909301926001928301920161266a565b346103a65760206003193601126103a65767ffffffffffffffff6126c5615cbb565b166080515260096020526126e0600560406080512001618023565b805190601f196127086126f284616179565b936127006040519586615dd2565b808552616179565b016080515b8181106127e95750506080515b815181101561276457806127306001928461638e565b5160805152600a6020526127486040608051206163f5565b612752828661638e565b5261275d818561638e565b500161271a565b826040518091602082016020835281518091526040830190602060408260051b860101930191608051905b82821061279e57505050500390f35b919360206127d9827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851615bf6565b960192019201859493919261278f565b80606060208093870101520161270d565b346103a6576080516003193601126103a65760206040516080518152f35b346103a6576080516003193601126103a6576002548060d01c9081151580612873575b156128685760a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b50506080518061284c565b504282101561283b565b346103a65760206003193601126103a65760043567ffffffffffffffff81116103a65760a060031982360301126103a6576128b6616375565b506128bf616375565b50608481016128cd816160d7565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612d8157602482019177ffffffffffffffff0000000000000000000000000000000061292684616047565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105ee5760805191612d62575b506123e3576129a9604482016160d7565b7f0000000000000000000000000000000000000000000000000000000000000000612d0f575b506064906129df6120ec85616047565b013590608051600014612c755761ffff600b541680612c405750612a0861214561212b85616047565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff612a3e61218b86616047565b604080516001600160a01b03929092168252602082018690529190931692a25b612a6782616047565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016808252336020830152918101849052909167ffffffffffffffff16907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a2612ae3611f4c84616047565b92612aec61745a565b60405194612af986615d9a565b8552602085015267ffffffffffffffff612b1282616047565b1660805152601260205260ff6001604060805120015460a01c16600014612c2b578067ffffffffffffffff612b49612b6b93616047565b16608051526012602052604060805120612b64858254616290565b9055616047565b505b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b156103a6576040517f47e7ef240000000000000000000000000000000000000000000000000000000081526080516001600160a01b03909316600482015260248101919091529182908180604481010391608051905af180156105ee57612c12575b60405160208082528190611e0d90820185615f6d565b608051612c1e91615dd2565b6080516103a65781612bfc565b50612c3882601054616290565b601055612b6d565b7fe08f03ef00000000000000000000000000000000000000000000000000000000608051526080516004526024526044608051fd5b5067ffffffffffffffff612c8883616047565b168060005260096020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280612cef60406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391617c32565b604080516001600160a01b039290921682526020820192909252a2612a5e565b6001600160a01b0316612d2f816000526004602052604060002054151590565b6129cf577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b612d7b915060203d602011612437576124298183615dd2565b84612998565b61244f6001600160a01b03916160d7565b346103a65760606003193601126103a65760043567ffffffffffffffff81116103a657612dc3903690600401615c37565b9060243567ffffffffffffffff81116103a657612de4903690600401615f3c565b9060443567ffffffffffffffff81116103a657612e05903690600401615f3c565b7f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af608051526080516020526040608051206001600160a01b03331660005260205260ff604060002054161580612efc575b61061057838614801590612ef2575b612ec6576080515b868110612e7a5760805180f35b80612ec0612e8e610dda6001948b8b61608b565b612e99838989616365565b612eba612eb2612eaa86898b616365565b923690615fc1565b913690615fc1565b916172a5565b01612e6d565b7f568efce200000000000000000000000000000000000000000000000000000000608051526004608051fd5b5080861415612e65565b50608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615612e56565b346103a65760406003193601126103a657612f48615bcc565b600435608051526080516020526001600160a01b036040608051209116600052602052602060ff604060002054166040519015158152f35b346103a6576080516003193601126103a65760206001600160a01b0360025416604051908152f35b346103a6576080516003193601126103a6576013546001600160a01b03811633036131d75760a01c67ffffffffffffffff16801561113b578060805152601260205261300b6040608051205482608051526014602052604060805120549061609b565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690803b156103a6576040517f69328dec000000000000000000000000000000000000000000000000000000008152608080516001600160a01b038516600484015260248301869052306044840152905191929091839160649183915af180156105ee576131be575b50608080518490526012602052516040812055601380547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff169055803b156103a657604051907f42966c680000000000000000000000000000000000000000000000000000000082528260048301528160248160805193608051905af180156105ee576131a5575b508161317b7fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e493617f1a565b506040805167ffffffffffffffff9092168252602082019290925290819081015b0390a160805180f35b6080516131b191615dd2565b6080516103a6578261314f565b6080516131ca91615dd2565b6080516103a657836130c7565b7f5fff6eee00000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103a65760c06003193601126103a65761321c615bb6565b50613225615cd2565b61322d615c79565b5060843567ffffffffffffffff81116103a65761324e903690600401615ce9565b505060a4359060028210156103a657611e0d9161268191604435906162ef565b346103a65760206003193601126103a65760206132a967ffffffffffffffff613295615cbb565b166000526008602052604060002054151590565b6040519015158152f35b346103a65760206003193601126103a65760206132d66132d1615cbb565b6162ad565b6001600160a01b0360405191168152f35b346103a65760206003193601126103a65760043567ffffffffffffffff81116103a657613318903690600401615c8a565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f8157610db991616930565b346103a6576080516003193601126103a65760206040517f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8152f35b346103a6576080516003193601126103a657600554600654604080516001600160a01b039093168352602083019190915290f35b346103a65760206003193601126103a6576133db615bb6565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815760206001600160a01b037f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b71689216807fffffffffffffffffffffffff00000000000000000000000000000000000000006013541617601355604051908152a160805180f35b346103a65760406003193601126103a65760043567ffffffffffffffff81116103a6576134a1903690600401615c37565b906024359067ffffffffffffffff82116103a657366023830112156103a65781600401359267ffffffffffffffff84116103a6576024830192602436918660061b0101116103a657608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f81576080515b818110613770575050506080515b8281106135395760805180f35b67ffffffffffffffff613550610dda83868661629d565b16158015613739575b8015613718575b6136d1576001600160a01b03613582602061357c84878761629d565b016160d7565b1615610d59578061366c61359e602061357c600195888861629d565b846001600160a01b0380868967ffffffffffffffff6135e2610dda8a604051946135c786615d7e565b608051865287602087019b168b526040860199878b5261629d565b166080515260126020526040608051209051815501935116167fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb604061369d610dda84888861629d565b6001600160a01b036136b5602061357c878b8b61629d565b67ffffffffffffffff845193168352166020820152a10161352c565b610dda9067ffffffffffffffff936136e89361629d565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b5061373367ffffffffffffffff613295610dda84878761629d565b15613560565b5067ffffffffffffffff613751610dda83868661629d565b1660805152601260205260ff6001604060805120015460a01c16613559565b67ffffffffffffffff613787610dda83858761608b565b1660805152601260205260ff6001604060805120015460a01c1615613862578067ffffffffffffffff6137c0610dda600194868861608b565b166080515260126020527f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93604060805120546137fe81601054616290565b60105567ffffffffffffffff613818610dda85888a61608b565b6080805191909216905260126020525160408120818155850155613840610dda84878961608b565b6040805167ffffffffffffffff9290921682526020820192909252a10161351e565b610dda906138799267ffffffffffffffff9461608b565b7f46f5f12b0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b346103a65760206003193601126103a6576138c2615bb6565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f8157601180547fffffffffffffffffffffffff000000000000000000000000000000000000000081166001600160a01b039384169081179092556040805191909316815260208101919091527f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b234918190810161319c565b346103a65760406003193601126103a657613981615bb6565b602435608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f81576001600160a01b038216918215610d59577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff000000000000000000000000000000000000000060055416176005558160065561319c60405192839283602090939291936001600160a01b0360408201951681520152565b346103a65760406003193601126103a657613a53615cbb565b613a5b615bcc565b90608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815767ffffffffffffffff16908160805152601260205260016040608051200190815460ff8160a01c1615613b2d5782547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179093556040805191909216815260208101929092527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f919081908101611c64565b837f46f5f12b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103a65760206003193601126103a65760043565ffffffffffff8116908181036103a657613b8a61674b565b613b9342617e41565b9165ffffffffffff613ba361658b565b1680821115613d1257507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92613bee9162069780811015613d015765ffffffffffff905b169061713f565b906002548060d01c80613c7a575b5050600280546001600160a01b031660a083901b79ffffffffffff0000000000000000000000000000000000000000161760d084901b7fffffffffffff0000000000000000000000000000000000000000000000000000161790556040805165ffffffffffff9283168152919092166020820152908190810161319c565b421115613cd35779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b8380613bfc565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5608051608051a1613ccc565b5065ffffffffffff62069780613be7565b0365ffffffffffff8111613d4c577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92613bee919061713f565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260116004526024608051fd5b346103a65760206003193601126103a657613d96615bb6565b613d9e61674b565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020613ddb613dcd42617e41565b613dd561658b565b9061713f565b65ffffffffffff6001600160a01b03613e0a6001549065ffffffffffff6001600160a01b0383169260a01c1690565b9690501694600154867fffffffffffff000000000000000000000000000000000000000000000000000079ffffffffffff00000000000000000000000000000000000000008660a01b169216171760015516613e76575b65ffffffffffff60405191168152a260805180f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109608051608051a1613e61565b346103a657613eb136615eb8565b608092919251608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815767ffffffffffffffff8216613f07816000526008602052604060002054151590565b15613f225750610db992613f1c913691615e11565b90616f2f565b7f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103a6576080516003193601126103a657602061040b6161e5565b346103a657613f7b36615e66565b9291608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f8157613fc592613fbd913691616191565b923691616191565b7f0000000000000000000000000000000000000000000000000000000000000000156140cd576080515b825181101561405557806001600160a01b0361400d6001938661638e565b51166140188161806e565b614024575b5001613fef565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a18461401d565b506080515b8151811015610db957806001600160a01b036140786001938561638e565b511680156140c75761408981617ea3565b614096575b505b0161405a565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a18361408e565b50614090565b7f35f4a7b300000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103a65760206003193601126103a657614112615cbb565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f81576013549067ffffffffffffffff8260a01c166142605767ffffffffffffffff811661417b816000526016602052604060002054151590565b6142315780156141ff577f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef49927fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff000000000000000000000000000000000000000060209460a01b16911617601355604051908152a160805180f35b7fd9a9cd6800000000000000000000000000000000000000000000000000000000608051526080516004526024608051fd5b7f1c49a87b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f692bc13100000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103a65760406003193601126103a6576142a5615cbb565b60243567ffffffffffffffff81116103a6576020916142cb6132a9923690600401615e48565b9061613c565b346103a6576080516003193601126103a657608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815760135467ffffffffffffffff8160a01c1690811561113b577fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff1660135560808051829052601460209081529051604080822091909155519182527f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda51891a160805180f35b346103a65760406003193601126103a65760043567ffffffffffffffff81116103a6578060040161010060031983360301126103a6576143d5615c68565b906040516143e281615d62565b60805190526144136144096144046143fd60c48701856160eb565b3691615e11565b616e99565b6064850135616d9e565b916084840190614422826160d7565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361243e57602485019277ffffffffffffffff0000000000000000000000000000000061447b85616047565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105ee5760805191614732575b506123e3576144fe6120ec85616047565b61450784616047565b9061451d60a48801926142cb6143fd85856160eb565b156146eb5750506145ed61218b60446020977ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09561ffff67ffffffffffffffff96161515600014614655578561457289616047565b1660805152600d8a5261458e6040608051208a61214f846160d7565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7866145bc61218b8b616047565b604080516001600160a01b03929092168252602082018d90529190931692a25b01946145e7866160d7565b50616047565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081168252336020830152909216908201526060810185905292169180608081015b0390a28060405161464c81615d62565b52604051908152f35b508461466088616047565b16806080515260098a527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c89806146cb6002604060805120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391617c32565b604080516001600160a01b039290921682526020820192909252a26145dc565b6146f592506160eb565b6126016040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916164b7565b61474b915060203d602011612437576124298183615dd2565b876144ed565b346103a6576080516003193601126103a65760206001600160a01b0360115416604051908152f35b346103a65760206003193601126103a65760043567ffffffffffffffff81116103a657806004019061010060031982360301126103a6576040516147bc81615d62565b60805190526147ce6064820135616cb4565b90608481016147dc816160d7565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603612d815750602481019277ffffffffffffffff0000000000000000000000000000000061483685616047565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105ee5760805191614b76575b506123e3576148b96120ec85616047565b6148c284616047565b906148d860a48401926142cb6143fd85856160eb565b156146eb57508291905067ffffffffffffffff6148f485616047565b168060805152600960205261493d6002604060805120016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016948591617c32565b604080516001600160a01b0385168152602081018690527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a267ffffffffffffffff61498a85616047565b1660805152601260205260406080512067ffffffffffffffff6149ac86616047565b16608051526014602052604060805120548015600014614b3857508054808511614b0457846149da9161609b565b90555b6044017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316614a13826160d7565b813b156103a6576040517f69328dec000000000000000000000000000000000000000000000000000000008152608080516001600160a01b0387811660048501526024840189905293909316604483015251909283916064918391905af180156105ee57614aeb575b5067ffffffffffffffff7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09161463c85614aba61218b602099616047565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b608051614af791615dd2565b6080516103a65784614a7c565b84907fa17e11d500000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b8091508411611c9b575067ffffffffffffffff614b5485616047565b16608051526014602052604060805120614b6f84825461609b565b90556149dd565b614b8f915060203d602011612437576124298183615dd2565b856148a8565b346103a65760a06003193601126103a657614bae615bb6565b50614bb7615cd2565b60443567ffffffffffffffff81116103a65760031960a091360301126103a657614bdf615c79565b506084359067ffffffffffffffff82116103a657614c0a67ffffffffffffffff923690600401615ce9565b5050604051614c1881615d17565b60805181526080516020820152608051604082015260606080519101521660805152600f6020526080604081512060405190614c5382615d17565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b346103a65760406003193601126103a657600435614cc2615bcc565b811580614dca575b614d14575b336001600160a01b03821603614ce85761177e916175b6565b7f6697b23200000000000000000000000000000000000000000000000000000000608051526004608051fd5b60015465ffffffffffff60a082901c16906001600160a01b031615801590614dba575b8015614da8575b614d7057507fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff60015416600155614ccf565b65ffffffffffff907f19ca5ebb0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b504265ffffffffffff82161015614d3e565b5065ffffffffffff811615614d37565b506001600160a01b03600254166001600160a01b03821614614cca565b346103a65760206003193601126103a65767ffffffffffffffff614e09615cbb565b16608051526012602052602060ff6001604060805120015460a01c166040519015158152f35b346103a65760406003193601126103a657600435614e4b615bcc565b81156117855781614e7061177461177e94600052600060205260016040600020015490565b6168b9565b346103a65760406003193601126103a657614e8e615cbb565b60243567ffffffffffffffff82168060805152601260205260ff6001604060805120015460a01c161580156150ca575b611cd757614ed9816000526016602052604060002054151590565b61509b578115610640576001600160a01b03614ef4846162ad565b1633036106105760805152601260205260406080512060ff600182015460a01c1660001461508657614f27828254616290565b90555b6040517f23b872dd000000000000000000000000000000000000000000000000000000006020820152336024820152306044820152606481018290527f000000000000000000000000000000000000000000000000000000000000000090614f999061050481608481016104f6565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b156103a6576040517f47e7ef24000000000000000000000000000000000000000000000000000000008152608080516001600160a01b039390931660048301526024820185905251909283916044918391905af180156105ee5761506d575b506040805167ffffffffffffffff9093168352602083019190915233917f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf9181908101611c64565b60805161507991615dd2565b6080516103a65782615025565b5061509381601054616290565b601055614f2a565b7f6469724600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b508015614ebe565b346103a65760606003193601126103a65760043561ffff8116908190036103a6576150fb615c68565b9060443567ffffffffffffffff81116103a65761511c903690600401615c8a565b90608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f815761ffff8416936127108510156151d55783927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba95926151c4926040967fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600b549360101b1692161717600b55616930565b82519182526020820152a160805180f35b847f95f3517a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103a65760406003193601126103a65760043567ffffffffffffffff81116103a657366023820112156103a657806004013567ffffffffffffffff81116103a65760248201916024369160a084020101116103a65760243567ffffffffffffffff81116103a65761527b903690600401615c37565b919092608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f81576080515b82811061532b575050506080515b8181106152ce5760805180f35b8067ffffffffffffffff6152e8610dda600194868861608b565b168060805152600f602052608051604060805120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8608051608051a2016152c1565b8061533c610dda6001938686616008565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70608061536a848888616008565b9261547a61544063ffffffff61546f6154338261546467ffffffffffffffff60208c0198169a8b8a5152600f60205260408a5120836153a88b61605c565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff000000006153e58761605c565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e61605c565b60401b1696019e8f61606d565b151560601b169516171617171790558261545c6040519a61607a565b16895261607a565b16602087015261607a565b166040840152615f97565b15156060820152a2016152b3565b346103a6576080516003193601126103a657602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103a65760206003193601126103a657602061040b600435600052600060205260016040600020015490565b346103a65760206003193601126103a657602061550f615bb6565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b346103a6576080516003193601126103a65760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103a6576080516003193601126103a65760408051611e0d916155b09082615dd2565b601d81527f53696c6f656455534443546f6b656e506f6f6c20312e362e332d6465760000006020820152604051918291602083526020830190615bf6565b346103a65760206003193601126103a657615607615bb6565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f81576156406161e5565b908161564c5760805180f35b60206001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599926156f16040517fa9059cbb00000000000000000000000000000000000000000000000000000000858201526156cb816104f6898660248401602090939291936001600160a01b0360408201951681520152565b7f0000000000000000000000000000000000000000000000000000000000000000617811565b6040519485521692a28080610db9565b346103a65760206003193601126103a65761571a615bb6565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610f81576020816112237ff7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e38993616817565b346103a6576080516003193601126103a65761579261674b565b6002548060d01c806157b6575b6001600160a01b0360025416600255608051608051f35b42111561580f5779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b808061579f565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5608051608051a1615808565b346103a65760206003193601126103a6576004358015610640576001600160a01b0361586a6080516162ad565b1633036106105760808051805260126020525160409020600181015460a01c60ff1680156159c75781545b808411611c9b5750156159b2576158ad82825461609b565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b156103a6576040517f69328dec000000000000000000000000000000000000000000000000000000008152608080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016600484015260248301859052336044840152905191929091839160649183915af180156105ee576159a0575b506040519060805150608051825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a260805180f35b6080516159ac91615dd2565b81615961565b506159bf8160105461609b565b6010556158b0565b601054615895565b346103a6576080516003193601126103a6576020604051620697808152f35b346103a65760206003193601126103a657600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036103a657817ff208a58f0000000000000000000000000000000000000000000000000000000060209314908115615b8c575b8115615b62575b8115615b38575b8115615b0e575b8115615a7e575b5015158152f35b7f3149878600000000000000000000000000000000000000000000000000000000811491508115615ab1575b5083615a77565b7f7965db0b00000000000000000000000000000000000000000000000000000000811491508115615ae4575b5083615aaa565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483615add565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150615a70565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150615a69565b7f479eecb20000000000000000000000000000000000000000000000000000000081149150615a62565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150615a5b565b600435906001600160a01b0382168203610db457565b602435906001600160a01b0382168203610db457565b35906001600160a01b0382168203610db457565b919082519283825260005b848110615c22575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201615c01565b9181601f84011215610db45782359167ffffffffffffffff8311610db4576020808501948460051b010111610db457565b6024359061ffff82168203610db457565b6064359061ffff82168203610db457565b9181601f84011215610db45782359167ffffffffffffffff8311610db45760208085019460e08502010111610db457565b6004359067ffffffffffffffff82168203610db457565b6024359067ffffffffffffffff82168203610db457565b9181601f84011215610db45782359167ffffffffffffffff8311610db45760208381860195010111610db457565b6080810190811067ffffffffffffffff821117615d3357604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117615d3357604052565b6060810190811067ffffffffffffffff821117615d3357604052565b6040810190811067ffffffffffffffff821117615d3357604052565b60a0810190811067ffffffffffffffff821117615d3357604052565b90601f601f19910116810190811067ffffffffffffffff821117615d3357604052565b67ffffffffffffffff8111615d3357601f01601f191660200190565b929192615e1d82615df5565b91615e2b6040519384615dd2565b829481845281830111610db4578281602093846000960137010152565b9080601f83011215610db457816020615e6393359101615e11565b90565b6040600319820112610db45760043567ffffffffffffffff8111610db45781615e9191600401615c37565b929092916024359067ffffffffffffffff8211610db457615eb491600401615c37565b9091565b906040600319830112610db45760043567ffffffffffffffff81168103610db457916024359067ffffffffffffffff8211610db457615eb491600401615ce9565b602060408183019282815284518094520192019060005b818110615f1d5750505090565b82516001600160a01b0316845260209384019390920191600101615f10565b9181601f84011215610db45782359167ffffffffffffffff8311610db45760208085019460608502010111610db457565b615e63916020615f868351604084526040840190615bf6565b920151906020818403910152615bf6565b35908115158203610db457565b35906fffffffffffffffffffffffffffffffff82168203610db457565b9190826060910312610db457604051615fd981615d7e565b6040616003818395615fea81615f97565b8552615ff860208201615fa4565b602086015201615fa4565b910152565b91908110156160185760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff81168103610db45790565b3563ffffffff81168103610db45790565b358015158103610db45790565b359063ffffffff82168203610db457565b91908110156160185760051b0190565b919082039182116160a857565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b356001600160a01b0381168103610db45790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610db4570180359067ffffffffffffffff8211610db457602001918136038313610db457565b9067ffffffffffffffff615e6392166000526009602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff8111615d335760051b60200190565b92919061619d81616179565b936161ab6040519586615dd2565b602085838152019160051b8101928311610db457905b8282106161cd57505050565b602080916161da84615be2565b8152019101906161c1565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561628457600091616255575090565b90506020813d60201161627c575b8161627060209383615dd2565b81010312610db4575190565b3d9150616263565b6040513d6000823e3d90fd5b919082018092116160a857565b91908110156160185760061b0190565b67ffffffffffffffff16600052601260205260016040600020015460ff8160a01c166162e357506001600160a01b036011541690565b6001600160a01b031690565b67ffffffffffffffff16600052600e60205260406000209160028110156163365760011461632557816001615e639301906171b1565b8160026003615e63940191016171b1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9190811015616018576060020190565b6040519061638282615d9a565b60606020838281520152565b80518210156160185760209160051b010190565b90600182811c921680156163eb575b60208310146163bc57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916163b1565b9060405191826000825492616409846163a2565b80845293600181169081156164775750600114616430575b5061642e92500383615dd2565b565b90506000929192526020600020906000915b81831061645b57505090602061642e9282010138616421565b6020919350806001915483858901015201910190918492616442565b6020935061642e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138616421565b601f8260209493601f19938186528686013760008582860101520116010190565b604051906164e582615db6565b60006080838281528260208201528260408201528260608201520152565b9060405161651081615db6565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526009602052615e6360046040600020016163f5565b6002548060d01c80151590816165ba575b50156165b05760a01c65ffffffffffff1690565b5060015460d01c90565b905042113861659c565b91908110156160185760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610db4570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610db4570180359067ffffffffffffffff8211610db457602001918160051b36038313610db457565b818102929181159184041417156160a857565b818110616676575050565b6000815560010161666b565b9160209082815201919060005b81811061669c5750505090565b9091926020806001926001600160a01b036166b688615be2565b16815201940192910161668f565b67ffffffffffffffff166166e5816000526008602052604060002054151590565b1561671e5780600052601260205260ff60016040600020015460a01c1661670d575060105490565b600052601260205260406000205490565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060005260045260246000fd5b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff161561678457565b7fe2517d3f0000000000000000000000000000000000000000000000000000000060005233600452600060245260446000fd5b80600052600060205260406000206001600160a01b03331660005260205260ff60406000205416156167e65750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b615e63907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af617759565b600254906001600160a01b03821661688f57615e63917fffffffffffffffffffffffff00000000000000000000000000000000000000006001600160a01b0383169116176002556000617759565b7f3fc3c27a0000000000000000000000000000000000000000000000000000000060005260046000fd5b9081156168ca575b615e6391617759565b600254916001600160a01b03831661688f577fffffffffffffffffffffffff00000000000000000000000000000000000000009092166001600160a01b038216176002556168c1565b356fffffffffffffffffffffffffffffffff81168103610db45790565b9160005b82811015616c505760e081028401600061694d82616047565b9067ffffffffffffffff821691616971836000526008602052604060002054151590565b15616c2457616a3a92604085936169e56169df946169df6169a5602060019c9b019261212b6169a03686615fc1565b617612565b91825463ffffffff8160801c16159081616c06575b81616bf7575b81616bdc575b81616bcd575b5080616bbe575b616b33575b3690615fc1565b90617944565b60808501926169f76169a03686615fc1565b8152600d6020522092835463ffffffff8160801c16159081616b15575b81616b06575b81616aeb575b81616adc575b5080616acd575b616a40575b503690615fc1565b01616934565b616a5d60a06fffffffffffffffffffffffffffffffff9201616913565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538616a32565b50616ad78261606d565b616a2d565b60ff915060a01c161538616a26565b6fffffffffffffffffffffffffffffffff8116159150616a20565b8589015460801c159150616a1a565b858901546fffffffffffffffffffffffffffffffff16159150616a14565b6fffffffffffffffffffffffffffffffff616b4f878b01616913565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783556169d8565b50616bc88161606d565b6169d3565b60ff915060a01c1615386169cc565b6fffffffffffffffffffffffffffffffff81161591506169c6565b848e015460801c1591506169c0565b848e01546fffffffffffffffffffffffffffffffff161591506169ba565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b9060ff8091169116039060ff82116160a857565b60ff16604d81116160a857600a0a90565b8115616c85570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f000000000000000000000000000000000000000000000000000000000000000060ff81169081600614616d995781600611616d6e576006616cf591616c56565b90604d60ff8316118015616d53575b616d1c575090616d16615e6392616c6a565b90616658565b90507fa9cb113d00000000000000000000000000000000000000000000000000000000600052600660045260245260445260646000fd5b50616d5d82616c6a565b8015616c8557600019048311616d04565b616d79906006616c56565b90604d60ff831611616d1c575090616d93615e6392616c6a565b90616c7b565b505090565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414616e7a57828411616e565790616de391616c56565b91604d60ff8416118015616e3b575b616e0557505090616d16615e6392616c6a565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50616e4583616c6a565b8015616c8557600019048411616df2565b616e5f91616c56565b91604d60ff841611616e0557505090616d93615e6392616c6a565b5050505090565b90816020910312610db457518015158103610db45790565b80518015616f0957602003616ecb578051602082810191830183900312610db457519060ff8211616ecb575060ff1690565b612601906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190615bf6565b50507f000000000000000000000000000000000000000000000000000000000000000090565b908051156171155767ffffffffffffffff81516020830120921691826000526009602052616f64816005604060002001617fce565b156170d157600052600a6020526040600020815167ffffffffffffffff8111615d3357616f9182546163a2565b601f811161709f575b506020601f82116001146170155791616fef827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea95936170059560009161700a575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190615bf6565b0390a2565b905084015138616fdc565b601f1982169083600052806000209160005b8181106170875750926170059492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea98961061706e575b5050811b019055611f51565b85015160001960f88460031b161c191690553880617062565b9192602060018192868a015181550194019201617027565b6170cb90836000526020600020601f840160051c81019160208510610ce357601f0160051c019061666b565b38616f9a565b50906126016040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190615bf6565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9065ffffffffffff8091169116019065ffffffffffff82116160a857565b906040519182815491828252602082019060005260206000209260005b81811061718f57505061642e92500383615dd2565b84546001600160a01b031683526001948501948794506020909301920161717a565b6171ba9061715d565b91600654801515918261729a575b50506171d2575090565b6171db9061715d565b908151806171e95750905090565b6171f4908251616290565b92601f1961721a61720486616179565b956172126040519788615dd2565b808752616179565b0136602086013760005b825181101561725557806001600160a01b036172426001938661638e565b511661724e828861638e565b5201617224565b509160005b815181101561729557806001600160a01b036172786001938561638e565b511661728e617288838751616290565b8861638e565b520161725a565b505050565b1015905038806171c8565b67ffffffffffffffff1660008181526008602052604090205490929190156173a757916173a460e092617370856172fc7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97617612565b846000526009602052617313816040600020617944565b61731c83617612565b846000526009602052617336836002604060002001617944565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6173dd6164d8565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff808351169161743a602085019361743461742763ffffffff8751164261609b565b8560808901511690616658565b90616290565b8082101561745357505b16825263ffffffff4216905290565b9050617444565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152615e63604082615dd2565b805160005b8181106174a657505050565b600181018082116160a8575b8281106174c2575060010161749a565b6001600160a01b036174d4838661638e565b51166001600160a01b036174e8838761638e565b5116146174f7576001016174b2565b6001600160a01b03617509838661638e565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b615e63906001600160a01b03600254166001600160a01b0382161461755f575b60006182bf565b7fffffffffffffffffffffffff000000000000000000000000000000000000000060025416600255617558565b615e63907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af6182bf565b90615e63918015806175f5575b156182bf577fffffffffffffffffffffffff0000000000000000000000000000000000000000600254166002556182bf565b506001600160a01b03600254166001600160a01b038316146175c3565b8051156176b2576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff6020830151161061764f5750565b6064906176b0604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040820151161580159061773a575b6176d95750565b6064906176b0604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff60208201511615156176d2565b80600052600060205260406000206001600160a01b03831660005260205260ff604060002054161560001461780a5780600052600060205260406000206001600160a01b038316600052602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b6001600160a01b036178939116916040926000808551936178328786615dd2565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d1561793c573d9161787783615df5565b9261788487519485615dd2565b83523d6000602085013e61836c565b805190816178a057505050565b6020806178b1938301019101616e81565b156178b95750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60609161836c565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991617a7d606092805461798163ffffffff8260801c164261609b565b9081617abc575b50506fffffffffffffffffffffffffffffffff6001816020860151169282815416808510600014617ab457508280855b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416178155617a318651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6173a460405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916179b8565b6fffffffffffffffffffffffffffffffff91617af1839283617aea6001880154948286169560801c90616658565b9116616290565b80821015617b7057505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880617988565b9050617afb565b67ffffffffffffffff16617b98816000526008602052604060002054151590565b15617c0557503360009081527f04c57a7d2bd5d0e733fe996f5b5aecc738999f0a2f9ddc4137bc3e1665bdf893602052604090205460ff1615617bd757565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9182549060ff8260a01c16158015617e39575b617e33576fffffffffffffffffffffffffffffffff82169160018501908154617c8a63ffffffff6fffffffffffffffffffffffffffffffff83169360801c164261609b565b9081617d95575b5050848110617d565750838310617ceb575050617cc06fffffffffffffffffffffffffffffffff92839261609b565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91617cfa818561609b565b926000198101908082116160a857617d1d617d22926001600160a01b0396616290565b616c7b565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611617e0957617db0926174349160801c90616658565b80841015617e045750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880617c91565b617dbb565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215617c45565b65ffffffffffff8111617e595765ffffffffffff1690565b7f6dfcc65000000000000000000000000000000000000000000000000000000000600052603060045260245260446000fd5b80548210156160185760005260206000200190600090565b80600052600460205260406000205415600014617f145760035468010000000000000000811015615d3357617efb617ee48260018594016003556003617e8b565b81939154906000199060031b92831b921b19161790565b9055600354906000526004602052604060002055600190565b50600090565b80600052601660205260406000205415600014617f145760155468010000000000000000811015615d3357617f5b617ee48260018594016015556015617e8b565b9055601554906000526016602052604060002055600190565b80600052600860205260406000205415600014617f145760075468010000000000000000811015615d3357617fb5617ee48260018594016007556007617e8b565b9055600754906000526008602052604060002055600190565b600082815260018201602052604090205461780a5780549068010000000000000000821015615d33578261800c617ee4846001809601855584617e8b565b905580549260005201602052604060002055600190565b906040519182815491828252602082019060005260206000209260005b81811061805557505061642e92500383615dd2565b8454835260019485019487945060209093019201618040565b600081815260046020526040902054801561780a5760001981018181116160a8576003549060001982019182116160a857818103618116575b50505060035480156180e757600019016180c2816003617e8b565b60001982549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b618138618127617ee4936003617e8b565b90549060031b1c9283926003617e8b565b905560005260046020526040600020553880806180a7565b600081815260086020526040902054801561780a5760001981018181116160a8576007549060001982019182116160a8578181036181c9575b50505060075480156180e757600019016181a4816007617e8b565b60001982549160031b1b19169055600755600052600860205260006040812055600190565b6181eb6181da617ee4936007617e8b565b90549060031b1c9283926007617e8b565b90556000526008602052604060002055388080618189565b90600182019181600052826020526040600020548015156000146182b65760001981018181116160a85782549060001982019182116160a85781810361827f575b505050805480156180e757600019019061825e8282617e8b565b60001982549160031b1b191690555560005260205260006040812055600190565b61829f61828f617ee49386617e8b565b90549060031b1c92839286617e8b565b905560005283602052604060002055388080618244565b50505050600090565b80600052600060205260406000206001600160a01b03831660005260205260ff6040600020541660001461780a5780600052600060205260406000206001600160a01b03831660005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b919290156183e75750815115618380575090565b3b156183895790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156183fa5750805190602001fd5b612601906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190615bf656fea164736f6c634300081a000aad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) AUTHORIZEDCALLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "AUTHORIZED_CALLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) AUTHORIZEDCALLERROLE() ([32]byte, error) {
	return _SiloedUSDCTokenPool.Contract.AUTHORIZEDCALLERROLE(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) AUTHORIZEDCALLERROLE() ([32]byte, error) {
	return _SiloedUSDCTokenPool.Contract.AUTHORIZEDCALLERROLE(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SiloedUSDCTokenPool.Contract.DEFAULTADMINROLE(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SiloedUSDCTokenPool.Contract.DEFAULTADMINROLE(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "RATE_LIMITER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _SiloedUSDCTokenPool.Contract.RATELIMITERADMINROLE(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _SiloedUSDCTokenPool.Contract.RATELIMITERADMINROLE(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) DefaultAdmin() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.DefaultAdmin(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) DefaultAdmin() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.DefaultAdmin(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) DefaultAdminDelay() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.DefaultAdminDelay(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.DefaultAdminDelay(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.DefaultAdminDelayIncreaseWait(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _SiloedUSDCTokenPool.Contract.DefaultAdminDelayIncreaseWait(&_SiloedUSDCTokenPool.CallOpts)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentInboundRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _SiloedUSDCTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_SiloedUSDCTokenPool.CallOpts, remoteChainSelector)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRoleAdmin(&_SiloedUSDCTokenPool.CallOpts, role)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SiloedUSDCTokenPool.Contract.GetRoleAdmin(&_SiloedUSDCTokenPool.CallOpts, role)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "hasRateLimitAdminRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.HasRateLimitAdminRole(&_SiloedUSDCTokenPool.CallOpts, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.HasRateLimitAdminRole(&_SiloedUSDCTokenPool.CallOpts, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.HasRole(&_SiloedUSDCTokenPool.CallOpts, role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SiloedUSDCTokenPool.Contract.HasRole(&_SiloedUSDCTokenPool.CallOpts, role, account)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _SiloedUSDCTokenPool.Contract.PendingDefaultAdmin(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _SiloedUSDCTokenPool.Contract.PendingDefaultAdmin(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _SiloedUSDCTokenPool.Contract.PendingDefaultAdminDelay(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _SiloedUSDCTokenPool.Contract.PendingDefaultAdminDelay(&_SiloedUSDCTokenPool.CallOpts)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AcceptDefaultAdminTransfer(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.AcceptDefaultAdminTransfer(&_SiloedUSDCTokenPool.TransactOpts)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyFinalityConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ApplyFinalityConfigUpdates(&_SiloedUSDCTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.BeginDefaultAdminTransfer(&_SiloedUSDCTokenPool.TransactOpts, newAdmin)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.BeginDefaultAdminTransfer(&_SiloedUSDCTokenPool.TransactOpts, newAdmin)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.CancelDefaultAdminTransfer(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.CancelDefaultAdminTransfer(&_SiloedUSDCTokenPool.TransactOpts)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ChangeDefaultAdminDelay(&_SiloedUSDCTokenPool.TransactOpts, newDelay)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.ChangeDefaultAdminDelay(&_SiloedUSDCTokenPool.TransactOpts, newDelay)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "grantRateLimitAdminRole", account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.GrantRateLimitAdminRole(&_SiloedUSDCTokenPool.TransactOpts, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.GrantRateLimitAdminRole(&_SiloedUSDCTokenPool.TransactOpts, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "grantRole", role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.GrantRole(&_SiloedUSDCTokenPool.TransactOpts, role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.GrantRole(&_SiloedUSDCTokenPool.TransactOpts, role, account)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "renounceRole", role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RenounceRole(&_SiloedUSDCTokenPool.TransactOpts, role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RenounceRole(&_SiloedUSDCTokenPool.TransactOpts, role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "revokeRateLimitAdminRole", account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RevokeRateLimitAdminRole(&_SiloedUSDCTokenPool.TransactOpts, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RevokeRateLimitAdminRole(&_SiloedUSDCTokenPool.TransactOpts, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "revokeRole", role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RevokeRole(&_SiloedUSDCTokenPool.TransactOpts, role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RevokeRole(&_SiloedUSDCTokenPool.TransactOpts, role, account)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RollbackDefaultAdminDelay(&_SiloedUSDCTokenPool.TransactOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RollbackDefaultAdminDelay(&_SiloedUSDCTokenPool.TransactOpts)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_SiloedUSDCTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_SiloedUSDCTokenPool.TransactOpts, rateLimitConfigArgs)
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

type SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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
		it.Event = new(SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
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

func (it *SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _SiloedUSDCTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceledIterator struct {
	Event *SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled)
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
		it.Event = new(SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled)
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

func (it *SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceledIterator{contract: _SiloedUSDCTokenPool.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled, error) {
	event := new(SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduledIterator struct {
	Event *SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled)
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
		it.Event = new(SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled)
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

func (it *SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduledIterator{contract: _SiloedUSDCTokenPool.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled, error) {
	event := new(SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolDefaultAdminTransferCanceledIterator struct {
	Event *SiloedUSDCTokenPoolDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolDefaultAdminTransferCanceled)
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
		it.Event = new(SiloedUSDCTokenPoolDefaultAdminTransferCanceled)
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

func (it *SiloedUSDCTokenPoolDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolDefaultAdminTransferCanceledIterator{contract: _SiloedUSDCTokenPool.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolDefaultAdminTransferCanceled)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*SiloedUSDCTokenPoolDefaultAdminTransferCanceled, error) {
	event := new(SiloedUSDCTokenPoolDefaultAdminTransferCanceled)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolDefaultAdminTransferScheduledIterator struct {
	Event *SiloedUSDCTokenPoolDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolDefaultAdminTransferScheduled)
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
		it.Event = new(SiloedUSDCTokenPoolDefaultAdminTransferScheduled)
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

func (it *SiloedUSDCTokenPoolDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*SiloedUSDCTokenPoolDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolDefaultAdminTransferScheduledIterator{contract: _SiloedUSDCTokenPool.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolDefaultAdminTransferScheduled)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*SiloedUSDCTokenPoolDefaultAdminTransferScheduled, error) {
	event := new(SiloedUSDCTokenPoolDefaultAdminTransferScheduled)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

type SiloedUSDCTokenPoolFinalityConfigUpdatedIterator struct {
	Event *SiloedUSDCTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolFinalityConfigUpdated)
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
		it.Event = new(SiloedUSDCTokenPoolFinalityConfigUpdated)
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

func (it *SiloedUSDCTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolFinalityConfigUpdatedIterator{contract: _SiloedUSDCTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolFinalityConfigUpdated)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolFinalityConfigUpdated, error) {
	event := new(SiloedUSDCTokenPoolFinalityConfigUpdated)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
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

type SiloedUSDCTokenPoolRateLimitAdminRoleGrantedIterator struct {
	Event *SiloedUSDCTokenPoolRateLimitAdminRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRateLimitAdminRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRateLimitAdminRoleGranted)
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
		it.Event = new(SiloedUSDCTokenPoolRateLimitAdminRoleGranted)
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

func (it *SiloedUSDCTokenPoolRateLimitAdminRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRateLimitAdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRateLimitAdminRoleGranted struct {
	Account common.Address
	Raw     types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolRateLimitAdminRoleGrantedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRateLimitAdminRoleGrantedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RateLimitAdminRoleGranted", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRateLimitAdminRoleGranted)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRateLimitAdminRoleGranted(log types.Log) (*SiloedUSDCTokenPoolRateLimitAdminRoleGranted, error) {
	event := new(SiloedUSDCTokenPoolRateLimitAdminRoleGranted)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRateLimitAdminRoleRevokedIterator struct {
	Event *SiloedUSDCTokenPoolRateLimitAdminRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRateLimitAdminRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRateLimitAdminRoleRevoked)
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
		it.Event = new(SiloedUSDCTokenPoolRateLimitAdminRoleRevoked)
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

func (it *SiloedUSDCTokenPoolRateLimitAdminRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRateLimitAdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRateLimitAdminRoleRevoked struct {
	Account common.Address
	Raw     types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolRateLimitAdminRoleRevokedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRateLimitAdminRoleRevokedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RateLimitAdminRoleRevoked", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRateLimitAdminRoleRevoked)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRateLimitAdminRoleRevoked(log types.Log) (*SiloedUSDCTokenPoolRateLimitAdminRoleRevoked, error) {
	event := new(SiloedUSDCTokenPoolRateLimitAdminRoleRevoked)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
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

type SiloedUSDCTokenPoolRoleAdminChangedIterator struct {
	Event *SiloedUSDCTokenPoolRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRoleAdminChanged)
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
		it.Event = new(SiloedUSDCTokenPoolRoleAdminChanged)
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

func (it *SiloedUSDCTokenPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SiloedUSDCTokenPoolRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRoleAdminChangedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRoleAdminChanged)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRoleAdminChanged(log types.Log) (*SiloedUSDCTokenPoolRoleAdminChanged, error) {
	event := new(SiloedUSDCTokenPoolRoleAdminChanged)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRoleGrantedIterator struct {
	Event *SiloedUSDCTokenPoolRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRoleGranted)
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
		it.Event = new(SiloedUSDCTokenPoolRoleGranted)
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

func (it *SiloedUSDCTokenPoolRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SiloedUSDCTokenPoolRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRoleGrantedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRoleGranted)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRoleGranted(log types.Log) (*SiloedUSDCTokenPoolRoleGranted, error) {
	event := new(SiloedUSDCTokenPoolRoleGranted)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type SiloedUSDCTokenPoolRoleRevokedIterator struct {
	Event *SiloedUSDCTokenPoolRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRoleRevoked)
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
		it.Event = new(SiloedUSDCTokenPoolRoleRevoked)
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

func (it *SiloedUSDCTokenPoolRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SiloedUSDCTokenPoolRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRoleRevokedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRoleRevoked)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRoleRevoked(log types.Log) (*SiloedUSDCTokenPoolRoleRevoked, error) {
	event := new(SiloedUSDCTokenPoolRoleRevoked)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

type GetDynamicConfig struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}
type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (SiloedUSDCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (SiloedUSDCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
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

func (SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (SiloedUSDCTokenPoolDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (SiloedUSDCTokenPoolDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (SiloedUSDCTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (SiloedUSDCTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
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

func (SiloedUSDCTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (SiloedUSDCTokenPoolRateLimitAdminRoleGranted) Topic() common.Hash {
	return common.HexToHash("0xf7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e389")
}

func (SiloedUSDCTokenPoolRateLimitAdminRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b")
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

func (SiloedUSDCTokenPoolRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (SiloedUSDCTokenPoolRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (SiloedUSDCTokenPoolRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (SiloedUSDCTokenPoolSiloRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f")
}

func (SiloedUSDCTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (SiloedUSDCTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
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
	AUTHORIZEDCALLERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	DefaultAdmin(opts *bind.CallOpts) (common.Address, error)

	DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error)

	DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error)

	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

	GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)

	GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error)

	GetUnsiloedLiquidity(opts *bind.CallOpts) (*big.Int, error)

	HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

	IsSiloed(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

		error)

	PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

		error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error)

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error)

	ExcludeTokensFromBurn(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ProposeCCTPMigration(opts *bind.TransactOpts, remoteChainSelector uint64) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ProvideSiloedLiquidity(opts *bind.TransactOpts, remoteChainSelector uint64, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error)

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error)

	SetSiloRebalancer(opts *bind.TransactOpts, remoteChainSelector uint64, newRebalancer common.Address) (*types.Transaction, error)

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

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*SiloedUSDCTokenPoolDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*SiloedUSDCTokenPoolDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*SiloedUSDCTokenPoolDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*SiloedUSDCTokenPoolDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*SiloedUSDCTokenPoolDefaultAdminTransferScheduled, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*SiloedUSDCTokenPoolDynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*SiloedUSDCTokenPoolFinalityConfigUpdated, error)

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

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*SiloedUSDCTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*SiloedUSDCTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolRateLimitAdminRoleGrantedIterator, error)

	WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error)

	ParseRateLimitAdminRoleGranted(log types.Log) (*SiloedUSDCTokenPoolRateLimitAdminRoleGranted, error)

	FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolRateLimitAdminRoleRevokedIterator, error)

	WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error)

	ParseRateLimitAdminRoleRevoked(log types.Log) (*SiloedUSDCTokenPoolRateLimitAdminRoleRevoked, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*SiloedUSDCTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*SiloedUSDCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*SiloedUSDCTokenPoolRemotePoolRemoved, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SiloedUSDCTokenPoolRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*SiloedUSDCTokenPoolRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SiloedUSDCTokenPoolRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*SiloedUSDCTokenPoolRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SiloedUSDCTokenPoolRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*SiloedUSDCTokenPoolRoleRevoked, error)

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
