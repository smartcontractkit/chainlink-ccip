// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lock_release_token_pool

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

var LockReleaseTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMITER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"direction\",\"type\":\"uint8\",\"internalType\":\"enumIPoolV2.CCVDirection\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"lockOrBurnOutV1\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDynamicConfig\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferLiquidity\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"outboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVsToAddAboveThreshold\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DynamicConfigSet\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"thresholdAmountForAdditionalCCVs\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103f45761768b803803809161001d8285610489565b833981019060a0818303126103f45780516001600160a01b038116918282036103f45761004c602082016104ac565b60408201519092906001600160401b0381116103f45782019480601f870112156103f4578551956001600160401b038711610473578660051b906020820197610098604051998a610489565b88526020808901928201019283116103f457602001905b82821061045b575050506100d160806100ca606085016104ba565b93016104ba565b93331561044557600180546001600160d01b031690556002546001600160a01b038116610434576001600160a01b03191633908117600255610112906104f8565b5080158015610423575b8015610412575b6104015760049260209260805260c0526040519283809263313ce56760e01b82525afa600091816103c0575b50610395575b5060a052600580546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610277575b604051616f7490816106f7823960805181818161033a0152818161198001528181611bee01528181611d5c01528181612391015281816125e0015281816126e501528181612e620152818161353e0152818161378801528181613834015281816139ba01528181613b6101528181614374015281816143c9015281816144d4015281816146600152614f86015260a051818181614303015281816159d201528181615a550152616013015260c051818181610da001528181611a0e0152818161241f015281816135cd0152613a49015260e051818181610d5a01528181611a530152818161246401526133300152f35b60405160206102868183610489565b60008252600036813760e051156103845760005b8251811015610301576001906001600160a01b036102b882866104ce565b5116836102c48261059e565b6102d1575b50500161029a565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138836102c9565b50905060005b825181101561037b576001906001600160a01b0361032582866104ce565b5116801561037557836103378261069c565b610345575b50505b01610307565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1388361033c565b5061033f565b50505038610187565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036103a95750610155565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103f9575b816103dc60209383610489565b810103126103f4576103ed906104ac565b903861014f565b600080fd5b3d91506103cf565b630a64406560e11b60005260046000fd5b506001600160a01b03831615610123565b506001600160a01b0385161561011c565b631fe1e13d60e11b60005260046000fd5b636116401160e11b600052600060045260246000fd5b60208091610468846104ba565b8152019101906100af565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761047357604052565b519060ff821682036103f457565b51906001600160a01b03821682036103f457565b80518210156104e25760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b038116600090815260008051602061766b833981519152602052604090205460ff16610580576001600160a01b0316600081815260008051602061766b83398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b80548210156104e25760005260206000200190600090565b600081815260046020526040902054801561069557600019810181811161067f5760035460001981019190821161067f5781810361062e575b505050600354801561061857600019016105f2816003610586565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61066761063f610650936003610586565b90549060031b1c9283926003610586565b819391549060031b91821b91600019901b19161790565b905560005260046020526040600020553880806105d7565b634e487b7160e01b600052601160045260246000fd5b5050600090565b806000526004602052604060002054156000146105805760035468010000000000000000811015610473576106dd6106508260018594016003556003610586565b905560035490600052600460205260406000205560019056fe60a080604052600436101561001357600080fd5b60006080526080513560e01c90816301ffc9a71461478f57508063022d63fb146147705780630a861f2a146146385780630aa6220b146145735780630bd7c46d14614508578063164e68de1461444f578063181f5a77146143ed57806321df0da7146143a8578063240028e814614354578063248a9ca31461432757806324f65ee7146142e85780632a10097b1461406f5780632c286daf14613f465780632f2ff15d14613f0057806336568abe14613dbf57806337b1924714613cae5780633907753714613929578063432a6ba314613901578063489a68f2146134a75780634c5ef0ed1461346257806354c8a4f3146132e05780635df45a37146132bc57806362ddd3c41461321a578063634e93da146130f4578063649a5ec714612ed45780636632008714612d32578063698c2c6614612c575780636cfd155314612b9a5780637437ff9f14612b66578063791e5a1014612b2a578063804ba5a914612abf57806384ef8ffc146129e75780638926f54f14612a7a57806389720a6214612a0f5780638da5cb5b146129e757806391d1485414612996578063962d4020146127f95780639a4575b91461233f578063a1eda53c146122da578063a217fddf146122bc578063a42a7b8b14612165578063a7cd63b7146120f7578063acfecf9114611fa3578063af58d59f14611f57578063b1c71c6514611903578063b7946580146118cb578063c4bffe2b146117af578063c75eea9c14611704578063cc8463c8146116d8578063cefc1429146115d4578063cf6eefb714611580578063cf7401f3146113c9578063d547741f1461134b578063d602b9fd146112c5578063d966866b14610e52578063da90a9f314610dc4578063dc0bd97114610d7f578063e0351e1314610d41578063e58d80c714610cd2578063e8a1da17146103c35763eb521a4c146102c457600080fd5b346103bd5760206003193601126103bd576004356001600160a01b0360105416330361038d576040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064808201839052815261035e90610338608482614a52565b7f000000000000000000000000000000000000000000000000000000000000000061639d565b337fc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088608051608051a360805180f35b7f8e4a23d60000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b60805180fd5b346103bd576103d136614c07565b91608093919351608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610ca65790608051905b828210610ae4575050506080519163ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee182360301935b81811015610ade57600581901b830135858112156103bd578301610120813603126103bd576040519461047b86614a1a565b813567ffffffffffffffff81168103610ad9578652602082013567ffffffffffffffff81116103bd5782019436601f870112156103bd5785356104bd81614ede565b966104cb6040519889614a52565b81885260208089019260051b820101903682116103bd5760208101925b828410610aaa575050505060208701958652604083013567ffffffffffffffff81116103bd5761051b9036908501614be9565b91604088019283526105456105333660608701614d62565b9460608a0195865260c0369101614d62565b95608089019687526105578551616256565b6105618751616256565b83515115610a7e5761057d67ffffffffffffffff8a5116616dec565b15610a435767ffffffffffffffff8951166080515260096020526040608051206106c186516fffffffffffffffffffffffffffffffff6040820151169061067c6fffffffffffffffffffffffffffffffff602083015116915115158360806040516105e781614a1a565b858152602081018b905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808a901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6107e788516fffffffffffffffffffffffffffffffff604082015116906107a26fffffffffffffffffffffffffffffffff6020830151169151151583608060405161070b81614a1a565b858152602081018b9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808b901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b6004855191019080519067ffffffffffffffff8211610a125761080a83546150a8565b601f81116109d3575b506020906001601f84111461096957918091610846936080519261095e575b50506000198260011b9260031b1c19161790565b90555b6080515b88518051821015610882579061087c6001926108758367ffffffffffffffff8f511692615094565b5190615b59565b0161084d565b5050977f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c293919997509561095067ffffffffffffffff600197969498511692519351915161091c6108e760405196879687526101006020880152610100870190614a91565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019392909193610449565b015190508e80610832565b90601f1983169184608051528160805120926080515b8181106109bb57509084600195949392106109a2575b505050811b019055610849565b015160001960f88460031b161c191690558d8080610995565b9293602060018192878601518155019501930161097f565b610a02908460805152602060805120601f850160051c81019160208610610a08575b601f0160051c01906153a0565b8d610813565b90915081906109f5565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260416004526024608051fd5b67ffffffffffffffff8951167f1d5ad3c500000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f14c880ca00000000000000000000000000000000000000000000000000000000608051526004608051fd5b833567ffffffffffffffff81116103bd57602091610ace8392833691870101614be9565b8152019301926104e8565b600080fd5b60805180f35b9092919367ffffffffffffffff610b04610aff868886614e2c565b614de8565b1692610b0f84616b70565b15610c765783608051526009602052610b2f600560406080512001616a14565b926080515b8451811015610b6e5760019086608051526009602052610b67600560406080512001610b608389615094565b5190616c23565b5001610b34565b5093909491959250806080515260096020526005604060805120608051815560805160018201556080516002820155608051600382015560048101610bb381546150a8565b80610c26575b505001805490608051815581610c02575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a101909161040c565b60805152602060805120908101905b81811015610bca576080518155600101610c11565b601f8111600114610c3f575060805190555b8880610bb9565b610c5f9082608051526001601f6020608051209201861c820191016153a0565b608080518290525160208120918190559055610c38565b837f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f2b5c74de00000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103bd5760206003193601126103bd57610ceb614957565b7f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af608051526080516020526001600160a01b036040608051209116600052602052602060ff604060002054166040519015158152f35b346103bd576080516003193601126103bd5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b346103bd576080516003193601126103bd5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103bd5760206003193601126103bd57610ddd614957565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610ca657602081610e3b7fd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b936161d0565b506001600160a01b0360405191168152a160805180f35b346103bd5760206003193601126103bd5760043567ffffffffffffffff81116103bd57610e83903690600401614ad2565b90608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610ca65790608051915b818310610ec85760805180f35b610ed6610aff8484846152ca565b610eee610ee48585856152ca565b602081019061530a565b9091610f08610efe8787876152ca565b604081019061530a565b90610f21610f178989896152ca565b606081019061530a565b9091610f3b610f318b8b8b6152ca565b608081019061530a565b949097610f51610f4c368a84614ef6565b6160d9565b610f5f610f4c368486614ef6565b610f6d610f4c368688614ef6565b610f7b610f4c36888c614ef6565b604051610f8781614997565b610f92368a84614ef6565b8152610f9f368486614ef6565b6020820152610faf368688614ef6565b6040820152610fbf36888c614ef6565b606082015267ffffffffffffffff881660805152600e602052604060805120815180519067ffffffffffffffff8211610a1257680100000000000000008211610a125760209083548385558084106112a6575b500182608051526020608051206080515b8381106112895750505050602082015180519067ffffffffffffffff8211610a1257680100000000000000008211610a12576020906001840154836001860155808410611267575b500160018301608051526020608051206080515b83811061124a5750505050604082015180519067ffffffffffffffff8211610a1257680100000000000000008211610a12576020906002840154836002860155808410611228575b500160028301608051526020608051206080515b83811061120b575050505060036060919e9c9d9e019101519081519167ffffffffffffffff8311610a1257680100000000000000008311610a125760209082548484558085106111ec575b500190608051526020608051206080515b8381106111cf57505050506111b46080956111c4956111a67fece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d9a9660019e9d9c9a9661119867ffffffffffffffff976040519d8d8f9e8f90815201916153b7565b918b830360208d01526153b7565b9188830360408a01526153b7565b92858403606087015216966153b7565b0390a2019190610ebb565b60019060206001600160a01b038551169401938184015501611137565b61120590846080515285846080512091820191016153a0565b38611126565b60019060206001600160a01b0385511694019381840155016110db565b61124490600286016080515284846080512091820191016153a0565b386110c7565b60019060206001600160a01b03855116940193818401550161107f565b61128390600186016080515284846080512091820191016153a0565b3861106b565b60019060206001600160a01b038551169401938184015501611023565b6112bf90856080515284846080512091820191016153a0565b38611012565b346103bd576080516003193601126103bd576112df615453565b600180547fffffffffffff0000000000000000000000000000000000000000000000000000811690915560a01c65ffffffffffff1661131e5760805180f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109608051608051a1610ade565b346103bd5760406003193601126103bd5760043561136761496d565b811561139d578161139161138c61139694600052600060205260016040600020015490565b6154bf565b6161fa565b5060805180f35b7f3fc3c27a00000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103bd5760e06003193601126103bd576113e2614b6d565b6060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc01126103bd5760405161141881614a36565b60243580151581036103bd5781526044356fffffffffffffffffffffffffffffffff811681036103bd5760208201526064356fffffffffffffffffffffffffffffffff811681036103bd5760408201526060367fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c01126103bd576040519061149f82614a36565b60843580151581036103bd57825260a4356fffffffffffffffffffffffffffffffff811681036103bd57602083015260c4356fffffffffffffffffffffffffffffffff811681036103bd5760408301527f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af608051526080516020526040608051206001600160a01b03331660005260205260ff60406000205416158061154d575b61038d57610ade92615edc565b50608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615611540565b346103bd576080516003193601126103bd57604065ffffffffffff6115bb6001549065ffffffffffff6001600160a01b0383169260a01c1690565b6001600160a01b03849392935193168352166020820152f35b346103bd576080516003193601126103bd576001546001600160a01b031633036116a85760015460a081901c65ffffffffffff16906001600160a01b03168115801561169e575b61166e5761163d906116376001600160a01b036002541661617c565b50615549565b507fffffffffffff000000000000000000000000000000000000000000000000000060015416600155608051608051f35b507f19ca5ebb00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b504282101561161b565b7fc22c80220000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b346103bd576080516003193601126103bd5760206116f4615291565b65ffffffffffff60405191168152f35b346103bd5760206003193601126103bd5767ffffffffffffffff611726614b6d565b61172e6151de565b50166080515260096020526117ab61175261174d604060805120615209565b616054565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b346103bd576080516003193601126103bd57608051506040516007548082528160208101600760805152602060805120926080515b8181106118b25750506117f992500382614a52565b80519061181e61180883614ede565b926118166040519485614a52565b808452614ede565b90601f196020840192013683376080515b8151811015611861578067ffffffffffffffff61184e60019385615094565b511661185a8287615094565b520161182f565b505090604051918291602083019060208452518091526040830191906080515b81811061188f575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611881565b84548352600194850194869450602090930192016117e4565b346103bd5760206003193601126103bd576117ab6118ef6118ea614b6d565b61526f565b604051918291602083526020830190614a91565b346103bd5760606003193601126103bd5760043567ffffffffffffffff81116103bd5760a060031982360301126103bd5761193c614b03565b9060443567ffffffffffffffff81116103bd5761195d903690600401614be9565b5061196661507b565b50608481019061197582614e8d565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611f1657602481019077ffffffffffffffff000000000000000000000000000000006119ce83614de8565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e2c5760805191611ee7575b50611ebb57611a5160448201614e8d565b7f0000000000000000000000000000000000000000000000000000000000000000611e68575b5067ffffffffffffffff611a8a83614de8565b16611aa2816000526008602052604060002054151590565b15611e395760206001600160a01b0360055416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611e2c5760805190611de3575b6001600160a01b039150163303611db357606461ffff91013593169283151592838094611da4575b15611d035761ffff600b541694858110611ccf5750611c939450611b73611b63611b4985614de8565b67ffffffffffffffff16600052600c602052604060002090565b83611b6d84614e8d565b916167bb565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff611baf611ba986614de8565b93614e8d565b604080516001600160a01b03929092168252602082018690529190931692a25b918290611c9d575b506118ea81611be8611c6293614de8565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2614de8565b90611c6b61600c565b60405192611c78846149fe565b83526020830152604051928392604084526040840190614d0e565b9060208301520390f35b611c62919250611cc76118ea91612710611cc061ffff600b5460101c168361535e565b0490616047565b929150611bd7565b85907fe08f03ef00000000000000000000000000000000000000000000000000000000608051526004526024526044608051fd5b50611c93935067ffffffffffffffff611d1b83614de8565b16806080515260096020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280611d846040608051206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916167bb565b604080516001600160a01b039290921682526020820192909252a2611bcf565b5061ffff600b54161515611b20565b7f728fe07b0000000000000000000000000000000000000000000000000000000060805152336004526024608051fd5b506020813d602011611e24575b81611dfd60209383614a52565b810103126103bd57516001600160a01b03811681036103bd576001600160a01b0390611af8565b3d9150611df0565b6040513d608051823e3d90fd5b7fa9902c7e00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b6001600160a01b0316611e88816000526004602052604060002054151590565b611a77577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b7f53ad11d800000000000000000000000000000000000000000000000000000000608051526004608051fd5b611f09915060203d602011611f0f575b611f018183614a52565b810190615b41565b85611a40565b503d611ef7565b6001600160a01b03611f2783614e8d565b7f961c9a4f0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b346103bd5760206003193601126103bd5767ffffffffffffffff611f79614b6d565b611f816151de565b50166080515260096020526117ab61175261174d600260406080512001615209565b346103bd57611fb136614c59565b91608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610ca65767ffffffffffffffff1690612005826000526008602052604060002054151590565b156120c7578160805152600960205261203860056040608051200161202b368685614bb2565b6020815191012090616c23565b15612080577f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691926120776040519283926020845260208401916151bd565b0390a260805180f35b6120c3906040519384937f74f23c7c00000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916151bd565b0390fd5b507f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103bd576080516003193601126103bd57608051506040516003548082526020820190600360805152602060805120906080515b81811061214f576117ab8561214381870382614a52565b60405191829182614c9a565b825484526020909301926001928301920161212c565b346103bd5760206003193601126103bd5767ffffffffffffffff612187614b6d565b166080515260096020526121a2600560406080512001616a14565b805190601f196121ca6121b484614ede565b936121c26040519586614a52565b808552614ede565b016080515b8181106122ab5750506080515b815181101561222657806121f260019284615094565b5160805152600a60205261220a6040608051206150fb565b6122148286615094565b5261221f8185615094565b50016121dc565b826040518091602082016020835281518091526040830190602060408260051b860101930191608051905b82821061226057505050500390f35b9193602061229b827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614a91565b9601920192018594939192612251565b8060606020809387010152016121cf565b346103bd576080516003193601126103bd5760206040516080518152f35b346103bd576080516003193601126103bd576002548060d01c9081151580612335575b1561232a5760a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b50506080518061230e565b50428210156122fd565b346103bd5760206003193601126103bd5760043567ffffffffffffffff81116103bd5760a060031982360301126103bd5761237861507b565b506084810161238681614e8d565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036127e857602482019177ffffffffffffffff000000000000000000000000000000006123df84614de8565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e2c57608051916127c9575b50611ebb5761246260448201614e8d565b7f0000000000000000000000000000000000000000000000000000000000000000612776575b5067ffffffffffffffff61249b84614de8565b166124b3816000526008602052604060002054151590565b15611e395760206001600160a01b0360055416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611e2c576080519061272d575b6001600160a01b039150163303611db35760640135906080516000146126885761ffff600b541680612653575082612623926118ea92612551611b63611b496117ab98614de8565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff612587611ba986614de8565b604080516001600160a01b03929092168252602082018690529190931692a25b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff6125da84614de8565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681523360208201529081019490945216918060608101611c5a565b61262b61600c565b60405191612638836149fe565b82526020820152604051918291602083526020830190614d0e565b7fe08f03ef00000000000000000000000000000000000000000000000000000000608051526080516004526024526044608051fd5b506118ea826126239267ffffffffffffffff6126a66117ab96614de8565b168060005260096020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944828061270d60406000206001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916167bb565b604080516001600160a01b039290921682526020820192909252a26125a7565b506020813d60201161276e575b8161274760209383614a52565b810103126103bd57516001600160a01b03811681036103bd576001600160a01b0390612509565b3d915061273a565b6001600160a01b0316612796816000526004602052604060002054151590565b612488577fd0d2597600000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b6127e2915060203d602011611f0f57611f018183614a52565b84612451565b611f276001600160a01b0391614e8d565b346103bd5760606003193601126103bd5760043567ffffffffffffffff81116103bd5761282a903690600401614ad2565b9060243567ffffffffffffffff81116103bd5761284b903690600401614cdd565b9060443567ffffffffffffffff81116103bd5761286c903690600401614cdd565b7f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af608051526080516020526040608051206001600160a01b03331660005260205260ff604060002054161580612963575b61038d57838614801590612959575b61292d576080515b8681106128e15760805180f35b806129276128f5610aff6001948b8b614e2c565b61290083898961506b565b61292161291961291186898b61506b565b923690614d62565b913690614d62565b91615edc565b016128d4565b7f568efce200000000000000000000000000000000000000000000000000000000608051526004608051fd5b50808614156128cc565b50608051608051526080516020526040608051206001600160a01b03331660005260205260ff60406000205416156128bd565b346103bd5760406003193601126103bd576129af61496d565b600435608051526080516020526001600160a01b036040608051209116600052602052602060ff604060002054166040519015158152f35b346103bd576080516003193601126103bd5760206001600160a01b0360025416604051908152f35b346103bd5760c06003193601126103bd57612a28614957565b50612a31614b56565b612a39614b14565b5060843567ffffffffffffffff81116103bd57612a5a903690600401614b84565b505060a4359060028210156103bd576117ab916121439160443590614ff5565b346103bd5760206003193601126103bd576020612ab567ffffffffffffffff612aa1614b6d565b166000526008602052604060002054151590565b6040519015158152f35b346103bd5760206003193601126103bd5760043567ffffffffffffffff81116103bd57612af0903690600401614b25565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610ca657610ade91615638565b346103bd576080516003193601126103bd5760206040517f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8152f35b346103bd576080516003193601126103bd57600554600654604080516001600160a01b039093168352602083019190915290f35b346103bd5760206003193601126103bd57612bb3614957565b608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610ca65760407f64187bd7b97e66658c91904f3021d7c28de967281d18b1a20742348afdd6a6b3916001600160a01b036010549116807fffffffffffffffffffffffff00000000000000000000000000000000000000008316176010556001600160a01b038351921682526020820152a160805180f35b346103bd5760406003193601126103bd57612c70614957565b602435608051608051526080516020526040608051206001600160a01b03331660005260205260ff6040600020541615610ca6576001600160a01b038216918215610a7e577f78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238927fffffffffffffffffffffffff0000000000000000000000000000000000000000600554161760055581600655612d2960405192839283602090939291936001600160a01b0360408201951681520152565b0390a160805180f35b346103bd5760406003193601126103bd57612d4b614957565b6080805180528051602081815260408083203390935291905290512054602435919060ff1615610ca6576000198214612e1b575b6001600160a01b031690813b156103bd57604051907f0a861f2a00000000000000000000000000000000000000000000000000000000825280600483015260805182602481608051875af1908115611e2c577f6fa7abcf1345d1d478e5ea0da6b5f26a90eadb0546ef15ed3833944fbfd1db6292602092612e09575b50604051908152a260805180f35b608051612e1591614a52565b84612dfb565b90506040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa8015611e2c5760805190612e9b575b919050612d7f565b506020813d602011612ecc575b81612eb560209383614a52565b810103126103bd576001600160a01b039051612e93565b3d9150612ea8565b346103bd5760206003193601126103bd5760043565ffffffffffff8116908181036103bd57612f01615453565b612f0a426169ca565b9165ffffffffffff612f1a615291565b168082111561308957507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92612f6591620697808110156130785765ffffffffffff905b1690615d69565b906002548060d01c80612ff1575b5050600280546001600160a01b031660a083901b79ffffffffffff0000000000000000000000000000000000000000161760d084901b7fffffffffffff0000000000000000000000000000000000000000000000000000161790556040805165ffffffffffff92831681529190921660208201529081908101612d29565b42111561304a5779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b8380612f73565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5608051608051a1613043565b5065ffffffffffff62069780612f5e565b0365ffffffffffff81116130c3577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92612f659190615d69565b7f4e487b71000000000000000000000000000000000000000000000000000000006080515260116004526024608051fd5b346103bd5760206003193601126103bd5761310d614957565b613115615453565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020613152613144426169ca565b61314c615291565b90615d69565b65ffffffffffff6001600160a01b036131816001549065ffffffffffff6001600160a01b0383169260a01c1690565b9690501694600154867fffffffffffff000000000000000000000000000000000000000000000000000079ffffffffffff00000000000000000000000000000000000000008660a01b1692161717600155166131ed575b65ffffffffffff60405191168152a260805180f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109608051608051a16131d8565b346103bd5761322836614c59565b608080518052805160208181526040808320339093529190529051205491929160ff1615610ca65767ffffffffffffffff8216613272816000526008602052604060002054151590565b1561328d5750610ade92613287913691614bb2565b90615b59565b7f1e670e4b00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103bd576080516003193601126103bd5760206132d8614f4a565b604051908152f35b346103bd576132ee36614c07565b60808051805280516020818152604080832033909352919052905120549093929060ff1615610ca65761332e92613326913691614ef6565b923691614ef6565b7f000000000000000000000000000000000000000000000000000000000000000015613436576080515b82518110156133be57806001600160a01b0361337660019386615094565b511661338181616a77565b61338d575b5001613358565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a184613386565b506080515b8151811015610ade57806001600160a01b036133e160019385615094565b51168015613430576133f281616d8c565b6133ff575b505b016133c3565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1836133f7565b506133f9565b7f35f4a7b300000000000000000000000000000000000000000000000000000000608051526004608051fd5b346103bd5760406003193601126103bd5761347b614b6d565b60243567ffffffffffffffff81116103bd576020916134a1612ab5923690600401614be9565b90614ea1565b346103bd5760406003193601126103bd5760043567ffffffffffffffff81116103bd578060040161010060031983360301126103bd576134e5614b03565b6040519091906134f4816149e2565b608051905261352561351b61351661350f60c4870185614e3c565b3691614bb2565b61595e565b6064850135615a52565b91608484019061353482614e8d565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116911603611f1657602485019277ffffffffffffffff0000000000000000000000000000000061358d85614de8565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e2c57608051916138e2575b50611ebb5767ffffffffffffffff61361685614de8565b1661362e816000526008602052604060002054151590565b15611e395760206001600160a01b0360055416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611e2c57608051916138c3575b5015611db35761369a84614de8565b906136b060a48801926134a161350f8585614e3c565b1561387c57505067ffffffffffffffff60446020967ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09461ffff6080951615156000146137e8578361370188614de8565b16855152600d895261371b604086512089611b6d84614e8d565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff784613749611ba98a614de8565b604080516001600160a01b03929092168252602082018c90529190931692a25b016001600160a01b036137b46137ae61378184614e8d565b97610aff8a7f00000000000000000000000000000000000000000000000000000000000000009a8b6153f9565b92614e8d565b816040519716875233898801521660408601528560608601521692a2604051906137dd826149e2565b815260405190518152f35b50826137f387614de8565b1680855152600989527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c888061385c600260408a5120016001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169283916167bb565b604080516001600160a01b039290921682526020820192909252a2613769565b6138869250614e3c565b6120c36040519283927f24eb47e50000000000000000000000000000000000000000000000000000000084526020600485015260248401916151bd565b6138dc915060203d602011611f0f57611f018183614a52565b8761368b565b6138fb915060203d602011611f0f57611f018183614a52565b876135ff565b346103bd576080516003193601126103bd5760206001600160a01b0360105416604051908152f35b346103bd5760206003193601126103bd5760043567ffffffffffffffff81116103bd57806004019061010060031982360301126103bd5760405161396c816149e2565b608051905260405161397d816149e2565b60805190526139a261399861351661350f60c4850186614e3c565b6064830135615a52565b90608481016139b081614e8d565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169116036127e857506024810177ffffffffffffffff00000000000000000000000000000000613a0982614de8565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e2c5760805191613c8f575b50611ebb5767ffffffffffffffff613a9282614de8565b16613aaa816000526008602052604060002054151590565b15611e395760206001600160a01b0360055416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611e2c5760805191613c70575b5015611db357613b1681614de8565b613b2b60a48401916134a161350f8489614e3c565b15613c6657509192508291613b3f81614de8565b67ffffffffffffffff16918260805152600960205260805160409020600201927f0000000000000000000000000000000000000000000000000000000000000000936001600160a01b0385168096613b96926167bb565b604080516001600160a01b0387168152602081018890527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a26044019184613bdf84614e8d565b613be8926153f9565b613bf190614de8565b90613bfb90614e8d565b6040519283523360208401526001600160a01b0316604083015282606083015267ffffffffffffffff169060807ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc091a260405190613c58826149e2565b815260405190518152602090f35b6138869085614e3c565b613c89915060203d602011611f0f57611f018183614a52565b85613b07565b613ca8915060203d602011611f0f57611f018183614a52565b85613a7b565b346103bd5760a06003193601126103bd57613cc7614957565b50613cd0614b56565b60443567ffffffffffffffff81116103bd5760031960a091360301126103bd57613cf8614b14565b506084359067ffffffffffffffff82116103bd57613d2367ffffffffffffffff923690600401614b84565b5050604051613d3181614997565b60805181526080516020820152608051604082015260606080519101521660805152600f6020526080604081512060405190613d6c82614997565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b346103bd5760406003193601126103bd57600435613ddb61496d565b811580613ee3575b613e2d575b336001600160a01b03821603613e0157611396916161fa565b7f6697b23200000000000000000000000000000000000000000000000000000000608051526004608051fd5b60015465ffffffffffff60a082901c16906001600160a01b031615801590613ed3575b8015613ec1575b613e8957507fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff60015416600155613de8565b65ffffffffffff907f19ca5ebb0000000000000000000000000000000000000000000000000000000060805152166004526024608051fd5b504265ffffffffffff82161015613e57565b5065ffffffffffff811615613e50565b506001600160a01b03600254166001600160a01b03821614613de3565b346103bd5760406003193601126103bd57600435613f1c61496d565b811561139d5781613f4161138c61139694600052600060205260016040600020015490565b6155c1565b346103bd5760606003193601126103bd5760043561ffff8116908190036103bd57613f6f614b03565b9060443567ffffffffffffffff81116103bd57613f90903690600401614b25565b608080518052805160208181526040808320339093529190529051205490919060ff1615610ca65761ffff84169361271085101561403f5783927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba959261402e926040967fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600b549360101b1692161717600b55615638565b82519182526020820152a160805180f35b847f95f3517a00000000000000000000000000000000000000000000000000000000608051526004526024608051fd5b346103bd5760406003193601126103bd5760043567ffffffffffffffff81116103bd57366023820112156103bd57806004013567ffffffffffffffff81116103bd5760248201916024369160a084020101116103bd5760243567ffffffffffffffff81116103bd576140e5903690600401614ad2565b6080805180528051602081815260408083203390935291905290512054919390929160ff1615610ca6576080515b82811061418b575050506080515b81811061412e5760805180f35b8067ffffffffffffffff614148610aff6001948688614e2c565b168060805152600f602052608051604060805120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8608051608051a201614121565b8061419c610aff6001938686614da9565b7f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d7060806141ca848888614da9565b926142da6142a063ffffffff6142cf614293826142c467ffffffffffffffff60208c0198169a8b8a5152600f60205260408a5120836142088b614dfd565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff0000000061424587614dfd565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e614dfd565b60401b1696019e8f614e0e565b151560601b16951617161717179055826142bc6040519a614e1b565b168952614e1b565b166020870152614e1b565b166040840152614d38565b15156060820152a201614113565b346103bd576080516003193601126103bd57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103bd5760206003193601126103bd5760206132d8600435600052600060205260016040600020015490565b346103bd5760206003193601126103bd57602061436f614957565b6040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081169216919091148152f35b346103bd576080516003193601126103bd5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b346103bd576080516003193601126103bd57604080516117ab916144119082614a52565b601e81527f4c6f636b52656c65617365546f6b656e506f6f6c20312e362e332d64657600006020820152604051918291602083526020830190614a91565b346103bd5760206003193601126103bd57614468614957565b608080518052805160208181526040808320339093529190529051205460ff1615610ca657614495614f4a565b90816144a15760805180f35b60206001600160a01b03826144f8857f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599957f00000000000000000000000000000000000000000000000000000000000000006153f9565b6040519485521692a28080610ade565b346103bd5760206003193601126103bd57614521614957565b608080518052805160208181526040808320339093529190529051205460ff1615610ca657602081610e3b7ff7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e3899361551f565b346103bd576080516003193601126103bd5761458d615453565b6002548060d01c806145b1575b6001600160a01b0360025416600255608051608051f35b42111561460a5779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b808061459a565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5608051608051a1614603565b346103bd5760206003193601126103bd576004356001600160a01b0360105416330361038d577f00000000000000000000000000000000000000000000000000000000000000006040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b0386165afa8015611e2c5783916080519161473b575b501061470f57816146e09133906153f9565b337fc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719608051608051a360805180f35b7fbb55fd2700000000000000000000000000000000000000000000000000000000608051526004608051fd5b9150506020813d602011614768575b8161475760209383614a52565b810103126103bd57829051846146ce565b3d915061474a565b346103bd576080516003193601126103bd576020604051620697808152f35b346103bd5760206003193601126103bd57600435907fffffffff0000000000000000000000000000000000000000000000000000000082168092036103bd57817ff208a58f000000000000000000000000000000000000000000000000000000006020931490811561492d575b8115614903575b81156148d9575b81156148af575b811561481f575b5015158152f35b7f3149878600000000000000000000000000000000000000000000000000000000811491508115614852575b5083614818565b7f7965db0b00000000000000000000000000000000000000000000000000000000811491508115614885575b508361484b565b7f01ffc9a7000000000000000000000000000000000000000000000000000000009150148361487e565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150614811565b7f0e64dd29000000000000000000000000000000000000000000000000000000008114915061480a565b7f479eecb20000000000000000000000000000000000000000000000000000000081149150614803565b7faff2afbf00000000000000000000000000000000000000000000000000000000811491506147fc565b600435906001600160a01b0382168203610ad957565b602435906001600160a01b0382168203610ad957565b35906001600160a01b0382168203610ad957565b6080810190811067ffffffffffffffff8211176149b357604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff8211176149b357604052565b6040810190811067ffffffffffffffff8211176149b357604052565b60a0810190811067ffffffffffffffff8211176149b357604052565b6060810190811067ffffffffffffffff8211176149b357604052565b90601f601f19910116810190811067ffffffffffffffff8211176149b357604052565b67ffffffffffffffff81116149b357601f01601f191660200190565b919082519283825260005b848110614abd575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201614a9c565b9181601f84011215610ad95782359167ffffffffffffffff8311610ad9576020808501948460051b010111610ad957565b6024359061ffff82168203610ad957565b6064359061ffff82168203610ad957565b9181601f84011215610ad95782359167ffffffffffffffff8311610ad95760208085019460e08502010111610ad957565b6024359067ffffffffffffffff82168203610ad957565b6004359067ffffffffffffffff82168203610ad957565b9181601f84011215610ad95782359167ffffffffffffffff8311610ad95760208381860195010111610ad957565b929192614bbe82614a75565b91614bcc6040519384614a52565b829481845281830111610ad9578281602093846000960137010152565b9080601f83011215610ad957816020614c0493359101614bb2565b90565b6040600319820112610ad95760043567ffffffffffffffff8111610ad95781614c3291600401614ad2565b929092916024359067ffffffffffffffff8211610ad957614c5591600401614ad2565b9091565b906040600319830112610ad95760043567ffffffffffffffff81168103610ad957916024359067ffffffffffffffff8211610ad957614c5591600401614b84565b602060408183019282815284518094520192019060005b818110614cbe5750505090565b82516001600160a01b0316845260209384019390920191600101614cb1565b9181601f84011215610ad95782359167ffffffffffffffff8311610ad95760208085019460608502010111610ad957565b614c04916020614d278351604084526040840190614a91565b920151906020818403910152614a91565b35908115158203610ad957565b35906fffffffffffffffffffffffffffffffff82168203610ad957565b9190826060910312610ad957604051614d7a81614a36565b6040614da4818395614d8b81614d38565b8552614d9960208201614d45565b602086015201614d45565b910152565b9190811015614db95760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff81168103610ad95790565b3563ffffffff81168103610ad95790565b358015158103610ad95790565b359063ffffffff82168203610ad957565b9190811015614db95760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610ad9570180359067ffffffffffffffff8211610ad957602001918136038313610ad957565b356001600160a01b0381168103610ad95790565b9067ffffffffffffffff614c0492166000526009602052600560406000200190602081519101209060019160005201602052604060002054151590565b67ffffffffffffffff81116149b35760051b60200190565b929190614f0281614ede565b93614f106040519586614a52565b602085838152019160051b8101928311610ad957905b828210614f3257505050565b60208091614f3f84614983565b815201910190614f26565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115614fe957600091614fba575090565b90506020813d602011614fe1575b81614fd560209383614a52565b81010312610ad9575190565b3d9150614fc8565b6040513d6000823e3d90fd5b67ffffffffffffffff16600052600e602052604060002091600281101561503c5760011461502b57816001614c04930190615de8565b8160026003614c0494019101615de8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9190811015614db9576060020190565b60405190615088826149fe565b60606020838281520152565b8051821015614db95760209160051b010190565b90600182811c921680156150f1575b60208310146150c257565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f16916150b7565b906040519182600082549261510f846150a8565b808452936001811690811561517d5750600114615136575b5061513492500383614a52565b565b90506000929192526020600020906000915b8183106151615750509060206151349282010138615127565b6020919350806001915483858901015201910190918492615148565b602093506151349592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138615127565b601f8260209493601f19938186528686013760008582860101520116010190565b604051906151eb82614a1a565b60006080838281528260208201528260408201528260608201520152565b9060405161521681614a1a565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b67ffffffffffffffff166000526009602052614c0460046040600020016150fb565b6002548060d01c80151590816152c0575b50156152b65760a01c65ffffffffffff1690565b5060015460d01c90565b90504211386152a2565b9190811015614db95760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6181360301821215610ad9570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610ad9570180359067ffffffffffffffff8211610ad957602001918160051b36038313610ad957565b8181029291811591840414171561537157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181106153ab575050565b600081556001016153a0565b9160209082815201919060005b8181106153d15750505090565b9091926020806001926001600160a01b036153eb88614983565b1681520194019291016153c4565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000060208201526001600160a01b039290921660248301526044808301939093529181526151349161544e606483614a52565b61639d565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff161561548c57565b7fe2517d3f0000000000000000000000000000000000000000000000000000000060005233600452600060245260446000fd5b80600052600060205260406000206001600160a01b03331660005260205260ff60406000205416156154ee5750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b614c04907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af6164d0565b600254906001600160a01b03821661559757614c04917fffffffffffffffffffffffff00000000000000000000000000000000000000006001600160a01b03831691161760025560006164d0565b7f3fc3c27a0000000000000000000000000000000000000000000000000000000060005260046000fd5b9081156155d2575b614c04916164d0565b600254916001600160a01b038316615597577fffffffffffffffffffffffff00000000000000000000000000000000000000009092166001600160a01b038216176002556155c9565b356fffffffffffffffffffffffffffffffff81168103610ad95790565b9160005b828110156159585760e081028401600061565582614de8565b9067ffffffffffffffff821691615679836000526008602052604060002054151590565b1561592c5761574292604085936156ed6156e7946156e76156ad602060019c9b0192611b496156a83686614d62565b616256565b91825463ffffffff8160801c1615908161590e575b816158ff575b816158e4575b816158d5575b50806158c6575b61583b575b3690614d62565b90616588565b60808501926156ff6156a83686614d62565b8152600d6020522092835463ffffffff8160801c1615908161581d575b8161580e575b816157f3575b816157e4575b50806157d5575b615748575b503690614d62565b0161563c565b61576560a06fffffffffffffffffffffffffffffffff920161561b565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783553861573a565b506157df82614e0e565b615735565b60ff915060a01c16153861572e565b6fffffffffffffffffffffffffffffffff8116159150615728565b8589015460801c159150615722565b858901546fffffffffffffffffffffffffffffffff1615915061571c565b6fffffffffffffffffffffffffffffffff615857878b0161561b565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783556156e0565b506158d081614e0e565b6156db565b60ff915060a01c1615386156d4565b6fffffffffffffffffffffffffffffffff81161591506156ce565b848e015460801c1591506156c8565b848e01546fffffffffffffffffffffffffffffffff161591506156c2565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b805180156159ce57602003615990578051602082810191830183900312610ad957519060ff8211615990575060ff1690565b6120c3906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614a91565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161537157565b60ff16604d811161537157600a0a90565b8115615a23570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff811692828414615b3a57828411615b105790615a97916159f4565b91604d60ff8416118015615af5575b615abf57505090615ab9614c0492615a08565b9061535e565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50615aff83615a08565b8015615a2357600019048411615aa6565b615b19916159f4565b91604d60ff841611615abf57505090615b34614c0492615a08565b90615a19565b5050505090565b90816020910312610ad957518015158103610ad95790565b90805115615d3f5767ffffffffffffffff81516020830120921691826000526009602052615b8e816005604060002001616e46565b15615cfb57600052600a6020526040600020815167ffffffffffffffff81116149b357615bbb82546150a8565b601f8111615cc9575b506020601f8211600114615c3f5791615c19827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9593615c2f95600091615c34575b506000198260011b9260031b1c19161790565b9055604051918291602083526020830190614a91565b0390a2565b905084015138615c06565b601f1982169083600052806000209160005b818110615cb1575092615c2f9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615c98575b5050811b0190556118ef565b85015160001960f88460031b161c191690553880615c8c565b9192602060018192868a015181550194019201615c51565b615cf590836000526020600020601f840160051c81019160208510610a0857601f0160051c01906153a0565b38615bc4565b50906120c36040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614a91565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9065ffffffffffff8091169116019065ffffffffffff821161537157565b906040519182815491828252602082019060005260206000209260005b818110615db957505061513492500383614a52565b84546001600160a01b0316835260019485019487945060209093019201615da4565b9190820180921161537157565b615df190615d87565b916006548015159182615ed1575b5050615e09575090565b615e1290615d87565b90815180615e205750905090565b615e2b908251615ddb565b92601f19615e51615e3b86614ede565b95615e496040519788614a52565b808752614ede565b0136602086013760005b8251811015615e8c57806001600160a01b03615e7960019386615094565b5116615e858288615094565b5201615e5b565b509160005b8151811015615ecc57806001600160a01b03615eaf60019385615094565b5116615ec5615ebf838751615ddb565b88615094565b5201615e91565b505050565b101590503880615dff565b67ffffffffffffffff166000818152600860205260409020549092919015615fde5791615fdb60e092615fa785615f337f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97616256565b846000526009602052615f4a816040600020616588565b615f5383616256565b846000526009602052615f6d836002604060002001616588565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152614c04604082614a52565b9190820391821161537157565b61605c6151de565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff80835116916160b960208501936160b36160a663ffffffff87511642616047565b856080890151169061535e565b90615ddb565b808210156160d257505b16825263ffffffff4216905290565b90506160c3565b805160005b8181106160ea57505050565b60018101808211615371575b82811061610657506001016160de565b6001600160a01b036161188386615094565b51166001600160a01b0361612c8387615094565b51161461613b576001016160f6565b6001600160a01b0361614d8386615094565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b614c04906001600160a01b03600254166001600160a01b038216146161a3575b6000616cdf565b7fffffffffffffffffffffffff00000000000000000000000000000000000000006002541660025561619c565b614c04907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af616cdf565b90614c0491801580616239575b15616cdf577fffffffffffffffffffffffff000000000000000000000000000000000000000060025416600255616cdf565b506001600160a01b03600254166001600160a01b03831614616207565b8051156162f6576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106162935750565b6064906162f4604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff6040820151161580159061637e575b61631d5750565b6064906162f4604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515616316565b6001600160a01b0361641f9116916040926000808551936163be8786614a52565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d156164c8573d9161640383614a75565b9261641087519485614a52565b83523d6000602085013e616e9b565b8051908161642c57505050565b60208061643d938301019101615b41565b156164455750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091616e9b565b80600052600060205260406000206001600160a01b03831660005260205260ff60406000205416156000146165815780600052600060205260406000206001600160a01b038316600052602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916166c160609280546165c563ffffffff8260801c1642616047565b9081616700575b50506fffffffffffffffffffffffffffffffff60018160208601511692828154168085106000146166f857508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556166758651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b615fdb60405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916165fc565b6fffffffffffffffffffffffffffffffff9161673583928361672e6001880154948286169560801c9061535e565b9116615ddb565b808210156167b457505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff000000000000000000000000000000001617815538806165cc565b905061673f565b9182549060ff8260a01c161580156169c2575b6169bc576fffffffffffffffffffffffffffffffff8216916001850190815461681363ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642616047565b908161691e575b50508481106168df57508383106168745750506168496fffffffffffffffffffffffffffffffff928392616047565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916168838185616047565b92600019810190808211615371576168a66168ab926001600160a01b0396615ddb565b615a19565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82869293961161699257616939926160b39160801c9061535e565b8084101561698d5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff000000000000000000000000000000001617865592388061681a565b616944565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b5082156167ce565b65ffffffffffff81116169e25765ffffffffffff1690565b7f6dfcc65000000000000000000000000000000000000000000000000000000000600052603060045260245260446000fd5b906040519182815491828252602082019060005260206000209260005b818110616a4657505061513492500383614a52565b8454835260019485019487945060209093019201616a31565b8054821015614db95760005260206000200190600090565b60008181526004602052604090205480156165815760001981018181116153715760035490600019820191821161537157818103616b1f575b5050506003548015616af05760001901616acb816003616a5f565b60001982549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b616b58616b30616b41936003616a5f565b90549060031b1c9283926003616a5f565b81939154906000199060031b92831b921b19161790565b90556000526004602052604060002055388080616ab0565b60008181526008602052604090205480156165815760001981018181116153715760075490600019820191821161537157818103616be9575b5050506007548015616af05760001901616bc4816007616a5f565b60001982549160031b1b19169055600755600052600860205260006040812055600190565b616c0b616bfa616b41936007616a5f565b90549060031b1c9283926007616a5f565b90556000526008602052604060002055388080616ba9565b9060018201918160005282602052604060002054801515600014616cd657600019810181811161537157825490600019820191821161537157818103616c9f575b50505080548015616af0576000190190616c7e8282616a5f565b60001982549160031b1b191690555560005260205260006040812055600190565b616cbf616caf616b419386616a5f565b90549060031b1c92839286616a5f565b905560005283602052604060002055388080616c64565b50505050600090565b80600052600060205260406000206001600160a01b03831660005260205260ff604060002054166000146165815780600052600060205260406000206001600160a01b03831660005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b80600052600460205260406000205415600014616de657600354680100000000000000008110156149b357616dcd616b418260018594016003556003616a5f565b9055600354906000526004602052604060002055600190565b50600090565b80600052600860205260406000205415600014616de657600754680100000000000000008110156149b357616e2d616b418260018594016007556007616a5f565b9055600754906000526008602052604060002055600190565b600082815260018201602052604090205461658157805490680100000000000000008210156149b35782616e84616b41846001809601855584616a5f565b905580549260005201602052604060002055600190565b91929015616f165750815115616eaf575090565b3b15616eb85790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b825190915015616f295750805190602001fd5b6120c3906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190614a9156fea164736f6c634300081a000aad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

var LockReleaseTokenPoolABI = LockReleaseTokenPoolMetaData.ABI

var LockReleaseTokenPoolBin = LockReleaseTokenPoolMetaData.Bin

func DeployLockReleaseTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *LockReleaseTokenPool, error) {
	parsed, err := LockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LockReleaseTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockReleaseTokenPool{address: address, abi: *parsed, LockReleaseTokenPoolCaller: LockReleaseTokenPoolCaller{contract: contract}, LockReleaseTokenPoolTransactor: LockReleaseTokenPoolTransactor{contract: contract}, LockReleaseTokenPoolFilterer: LockReleaseTokenPoolFilterer{contract: contract}}, nil
}

type LockReleaseTokenPool struct {
	address common.Address
	abi     abi.ABI
	LockReleaseTokenPoolCaller
	LockReleaseTokenPoolTransactor
	LockReleaseTokenPoolFilterer
}

type LockReleaseTokenPoolCaller struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type LockReleaseTokenPoolSession struct {
	Contract     *LockReleaseTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type LockReleaseTokenPoolCallerSession struct {
	Contract *LockReleaseTokenPoolCaller
	CallOpts bind.CallOpts
}

type LockReleaseTokenPoolTransactorSession struct {
	Contract     *LockReleaseTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type LockReleaseTokenPoolRaw struct {
	Contract *LockReleaseTokenPool
}

type LockReleaseTokenPoolCallerRaw struct {
	Contract *LockReleaseTokenPoolCaller
}

type LockReleaseTokenPoolTransactorRaw struct {
	Contract *LockReleaseTokenPoolTransactor
}

func NewLockReleaseTokenPool(address common.Address, backend bind.ContractBackend) (*LockReleaseTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(LockReleaseTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindLockReleaseTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPool{address: address, abi: abi, LockReleaseTokenPoolCaller: LockReleaseTokenPoolCaller{contract: contract}, LockReleaseTokenPoolTransactor: LockReleaseTokenPoolTransactor{contract: contract}, LockReleaseTokenPoolFilterer: LockReleaseTokenPoolFilterer{contract: contract}}, nil
}

func NewLockReleaseTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*LockReleaseTokenPoolCaller, error) {
	contract, err := bindLockReleaseTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCaller{contract: contract}, nil
}

func NewLockReleaseTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LockReleaseTokenPoolTransactor, error) {
	contract, err := bindLockReleaseTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolTransactor{contract: contract}, nil
}

func NewLockReleaseTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LockReleaseTokenPoolFilterer, error) {
	contract, err := bindLockReleaseTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolFilterer{contract: contract}, nil
}

func bindLockReleaseTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LockReleaseTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolTransactor.contract.Transfer(opts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockReleaseTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LockReleaseTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.contract.Transfer(opts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LockReleaseTokenPool.Contract.DEFAULTADMINROLE(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LockReleaseTokenPool.Contract.DEFAULTADMINROLE(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "RATE_LIMITER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _LockReleaseTokenPool.Contract.RATELIMITERADMINROLE(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _LockReleaseTokenPool.Contract.RATELIMITERADMINROLE(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) DefaultAdmin() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.DefaultAdmin(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) DefaultAdmin() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.DefaultAdmin(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) DefaultAdminDelay() (*big.Int, error) {
	return _LockReleaseTokenPool.Contract.DefaultAdminDelay(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _LockReleaseTokenPool.Contract.DefaultAdminDelay(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _LockReleaseTokenPool.Contract.DefaultAdminDelayIncreaseWait(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _LockReleaseTokenPool.Contract.DefaultAdminDelayIncreaseWait(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _LockReleaseTokenPool.Contract.GetAccumulatedFees(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _LockReleaseTokenPool.Contract.GetAccumulatedFees(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetAllowList(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetAllowList(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _LockReleaseTokenPool.Contract.GetAllowListEnabled(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _LockReleaseTokenPool.Contract.GetAllowListEnabled(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPool.Contract.GetCurrentInboundRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPool.Contract.GetCurrentInboundRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _LockReleaseTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

	error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getDynamicConfig")

	outstruct := new(GetDynamicConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ThresholdAmountForAdditionalCCVs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _LockReleaseTokenPool.Contract.GetDynamicConfig(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetDynamicConfig() (GetDynamicConfig,

	error) {
	return _LockReleaseTokenPool.Contract.GetDynamicConfig(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRebalancer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRebalancer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRebalancer() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRebalancer(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRebalancer() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRebalancer(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemotePools(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemotePools(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemoteToken(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRemoteToken(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRequiredCCVs(opts *bind.CallOpts, arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRequiredCCVs", arg0, remoteChainSelector, amount, arg3, arg4, direction)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRequiredCCVs(&_LockReleaseTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRequiredCCVs(arg0 common.Address, remoteChainSelector uint64, amount *big.Int, arg3 uint16, arg4 []byte, direction uint8) ([]common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRequiredCCVs(&_LockReleaseTokenPool.CallOpts, arg0, remoteChainSelector, amount, arg3, arg4, direction)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRmnProxy(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetRmnProxy(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRoleAdmin(&_LockReleaseTokenPool.CallOpts, role)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LockReleaseTokenPool.Contract.GetRoleAdmin(&_LockReleaseTokenPool.CallOpts, role)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _LockReleaseTokenPool.Contract.GetSupportedChains(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _LockReleaseTokenPool.Contract.GetSupportedChains(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetToken() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetToken(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.GetToken(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _LockReleaseTokenPool.Contract.GetTokenDecimals(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _LockReleaseTokenPool.Contract.GetTokenDecimals(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _LockReleaseTokenPool.Contract.GetTokenTransferFeeConfig(&_LockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _LockReleaseTokenPool.Contract.GetTokenTransferFeeConfig(&_LockReleaseTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "hasRateLimitAdminRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.HasRateLimitAdminRole(&_LockReleaseTokenPool.CallOpts, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.HasRateLimitAdminRole(&_LockReleaseTokenPool.CallOpts, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.HasRole(&_LockReleaseTokenPool.CallOpts, role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.HasRole(&_LockReleaseTokenPool.CallOpts, role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsRemotePool(&_LockReleaseTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsRemotePool(&_LockReleaseTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedChain(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedChain(&_LockReleaseTokenPool.CallOpts, remoteChainSelector)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedToken(&_LockReleaseTokenPool.CallOpts, token)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _LockReleaseTokenPool.Contract.IsSupportedToken(&_LockReleaseTokenPool.CallOpts, token)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) Owner() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.Owner(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) Owner() (common.Address, error) {
	return _LockReleaseTokenPool.Contract.Owner(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _LockReleaseTokenPool.Contract.PendingDefaultAdmin(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _LockReleaseTokenPool.Contract.PendingDefaultAdmin(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _LockReleaseTokenPool.Contract.PendingDefaultAdminDelay(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _LockReleaseTokenPool.Contract.PendingDefaultAdminDelay(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.SupportsInterface(&_LockReleaseTokenPool.CallOpts, interfaceId)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LockReleaseTokenPool.Contract.SupportsInterface(&_LockReleaseTokenPool.CallOpts, interfaceId)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LockReleaseTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) TypeAndVersion() (string, error) {
	return _LockReleaseTokenPool.Contract.TypeAndVersion(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _LockReleaseTokenPool.Contract.TypeAndVersion(&_LockReleaseTokenPool.CallOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AcceptDefaultAdminTransfer(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AcceptDefaultAdminTransfer(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AddRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.AddRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyAllowListUpdates(&_LockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyAllowListUpdates(&_LockReleaseTokenPool.TransactOpts, removes, adds)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyCCVConfigUpdates(&_LockReleaseTokenPool.TransactOpts, ccvConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyCCVConfigUpdates(&_LockReleaseTokenPool.TransactOpts, ccvConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyChainUpdates(&_LockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyChainUpdates(&_LockReleaseTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyFinalityConfigUpdates(&_LockReleaseTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyFinalityConfigUpdates(&_LockReleaseTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_LockReleaseTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_LockReleaseTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.BeginDefaultAdminTransfer(&_LockReleaseTokenPool.TransactOpts, newAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.BeginDefaultAdminTransfer(&_LockReleaseTokenPool.TransactOpts, newAdmin)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.CancelDefaultAdminTransfer(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.CancelDefaultAdminTransfer(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ChangeDefaultAdminDelay(&_LockReleaseTokenPool.TransactOpts, newDelay)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ChangeDefaultAdminDelay(&_LockReleaseTokenPool.TransactOpts, newDelay)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "grantRateLimitAdminRole", account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.GrantRateLimitAdminRole(&_LockReleaseTokenPool.TransactOpts, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.GrantRateLimitAdminRole(&_LockReleaseTokenPool.TransactOpts, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "grantRole", role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.GrantRole(&_LockReleaseTokenPool.TransactOpts, role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.GrantRole(&_LockReleaseTokenPool.TransactOpts, role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn0(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.LockOrBurn0(&_LockReleaseTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "provideLiquidity", amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ProvideLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ProvideLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ProvideLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint0(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.ReleaseOrMint0(&_LockReleaseTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RemoveRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RemoveRemotePool(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "renounceRole", role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RenounceRole(&_LockReleaseTokenPool.TransactOpts, role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RenounceRole(&_LockReleaseTokenPool.TransactOpts, role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "revokeRateLimitAdminRole", account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RevokeRateLimitAdminRole(&_LockReleaseTokenPool.TransactOpts, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RevokeRateLimitAdminRole(&_LockReleaseTokenPool.TransactOpts, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "revokeRole", role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RevokeRole(&_LockReleaseTokenPool.TransactOpts, role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RevokeRole(&_LockReleaseTokenPool.TransactOpts, role, account)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RollbackDefaultAdminDelay(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.RollbackDefaultAdminDelay(&_LockReleaseTokenPool.TransactOpts)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetChainRateLimiterConfig(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetChainRateLimiterConfig(&_LockReleaseTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetChainRateLimiterConfigs(&_LockReleaseTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetChainRateLimiterConfigs(&_LockReleaseTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_LockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_LockReleaseTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setDynamicConfig", router, thresholdAmountForAdditionalCCVs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetDynamicConfig(&_LockReleaseTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetDynamicConfig(router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetDynamicConfig(&_LockReleaseTokenPool.TransactOpts, router, thresholdAmountForAdditionalCCVs)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) SetRebalancer(opts *bind.TransactOpts, rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "setRebalancer", rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) SetRebalancer(rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRebalancer(&_LockReleaseTokenPool.TransactOpts, rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) SetRebalancer(rebalancer common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.SetRebalancer(&_LockReleaseTokenPool.TransactOpts, rebalancer)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) TransferLiquidity(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "transferLiquidity", from, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) TransferLiquidity(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferLiquidity(&_LockReleaseTokenPool.TransactOpts, from, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) TransferLiquidity(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.TransferLiquidity(&_LockReleaseTokenPool.TransactOpts, from, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawFees(&_LockReleaseTokenPool.TransactOpts, recipient)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawFees(&_LockReleaseTokenPool.TransactOpts, recipient)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.contract.Transact(opts, "withdrawLiquidity", amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolTransactorSession) WithdrawLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _LockReleaseTokenPool.Contract.WithdrawLiquidity(&_LockReleaseTokenPool.TransactOpts, amount)
}

type LockReleaseTokenPoolAllowListAddIterator struct {
	Event *LockReleaseTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolAllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*LockReleaseTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAllowListAddIterator{contract: _LockReleaseTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAllowListAdd)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*LockReleaseTokenPoolAllowListAdd, error) {
	event := new(LockReleaseTokenPoolAllowListAdd)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolAllowListRemoveIterator struct {
	Event *LockReleaseTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolAllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolAllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*LockReleaseTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolAllowListRemoveIterator{contract: _LockReleaseTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolAllowListRemove)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*LockReleaseTokenPoolAllowListRemove, error) {
	event := new(LockReleaseTokenPoolAllowListRemove)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolCCVConfigUpdatedIterator struct {
	Event *LockReleaseTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolCCVConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolCCVConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector             uint64
	OutboundCCVs                    []common.Address
	OutboundCCVsToAddAboveThreshold []common.Address
	InboundCCVs                     []common.Address
	InboundCCVsToAddAboveThreshold  []common.Address
	Raw                             types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCCVConfigUpdatedIterator{contract: _LockReleaseTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolCCVConfigUpdated)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*LockReleaseTokenPoolCCVConfigUpdated, error) {
	event := new(LockReleaseTokenPoolCCVConfigUpdated)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolChainAddedIterator struct {
	Event *LockReleaseTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainAddedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolChainAddedIterator{contract: _LockReleaseTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolChainAdded)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseChainAdded(log types.Log) (*LockReleaseTokenPoolChainAdded, error) {
	event := new(LockReleaseTokenPoolChainAdded)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolChainConfiguredIterator struct {
	Event *LockReleaseTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolChainConfiguredIterator{contract: _LockReleaseTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolChainConfigured)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseChainConfigured(log types.Log) (*LockReleaseTokenPoolChainConfigured, error) {
	event := new(LockReleaseTokenPoolChainConfigured)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolChainRemovedIterator struct {
	Event *LockReleaseTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolChainRemovedIterator{contract: _LockReleaseTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolChainRemoved)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolChainRemoved, error) {
	event := new(LockReleaseTokenPoolChainRemoved)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolConfigChangedIterator struct {
	Event *LockReleaseTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*LockReleaseTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolConfigChangedIterator{contract: _LockReleaseTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolConfigChanged)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseConfigChanged(log types.Log) (*LockReleaseTokenPoolConfigChanged, error) {
	event := new(LockReleaseTokenPoolConfigChanged)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolDefaultAdminDelayChangeCanceledIterator struct {
	Event *LockReleaseTokenPoolDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolDefaultAdminDelayChangeCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolDefaultAdminDelayChangeCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*LockReleaseTokenPoolDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolDefaultAdminDelayChangeCanceledIterator{contract: _LockReleaseTokenPool.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolDefaultAdminDelayChangeCanceled)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*LockReleaseTokenPoolDefaultAdminDelayChangeCanceled, error) {
	event := new(LockReleaseTokenPoolDefaultAdminDelayChangeCanceled)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolDefaultAdminDelayChangeScheduledIterator struct {
	Event *LockReleaseTokenPoolDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolDefaultAdminDelayChangeScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolDefaultAdminDelayChangeScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*LockReleaseTokenPoolDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolDefaultAdminDelayChangeScheduledIterator{contract: _LockReleaseTokenPool.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolDefaultAdminDelayChangeScheduled)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*LockReleaseTokenPoolDefaultAdminDelayChangeScheduled, error) {
	event := new(LockReleaseTokenPoolDefaultAdminDelayChangeScheduled)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolDefaultAdminTransferCanceledIterator struct {
	Event *LockReleaseTokenPoolDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolDefaultAdminTransferCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolDefaultAdminTransferCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*LockReleaseTokenPoolDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolDefaultAdminTransferCanceledIterator{contract: _LockReleaseTokenPool.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolDefaultAdminTransferCanceled)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*LockReleaseTokenPoolDefaultAdminTransferCanceled, error) {
	event := new(LockReleaseTokenPoolDefaultAdminTransferCanceled)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolDefaultAdminTransferScheduledIterator struct {
	Event *LockReleaseTokenPoolDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolDefaultAdminTransferScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolDefaultAdminTransferScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*LockReleaseTokenPoolDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolDefaultAdminTransferScheduledIterator{contract: _LockReleaseTokenPool.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolDefaultAdminTransferScheduled)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*LockReleaseTokenPoolDefaultAdminTransferScheduled, error) {
	event := new(LockReleaseTokenPoolDefaultAdminTransferScheduled)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolDynamicConfigSetIterator struct {
	Event *LockReleaseTokenPoolDynamicConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolDynamicConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolDynamicConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolDynamicConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolDynamicConfigSetIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolDynamicConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolDynamicConfigSet struct {
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	Raw                              types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterDynamicConfigSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolDynamicConfigSetIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolDynamicConfigSetIterator{contract: _LockReleaseTokenPool.contract, event: "DynamicConfigSet", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "DynamicConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolDynamicConfigSet)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseDynamicConfigSet(log types.Log) (*LockReleaseTokenPoolDynamicConfigSet, error) {
	event := new(LockReleaseTokenPoolDynamicConfigSet)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "DynamicConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolFinalityConfigUpdatedIterator struct {
	Event *LockReleaseTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolFinalityConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolFinalityConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolFinalityConfigUpdatedIterator{contract: _LockReleaseTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolFinalityConfigUpdated)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*LockReleaseTokenPoolFinalityConfigUpdated, error) {
	event := new(LockReleaseTokenPoolFinalityConfigUpdated)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolInboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolInboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolInboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolInboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolInboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLiquidityAddedIterator struct {
	Event *LockReleaseTokenPoolLiquidityAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolLiquidityAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityAdded struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityAddedIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityAdded", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityAdded)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityAdded(log types.Log) (*LockReleaseTokenPoolLiquidityAdded, error) {
	event := new(LockReleaseTokenPoolLiquidityAdded)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLiquidityRemovedIterator struct {
	Event *LockReleaseTokenPoolLiquidityRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolLiquidityRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityRemoved struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityRemovedIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityRemoved", providerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityRemoved)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityRemoved(log types.Log) (*LockReleaseTokenPoolLiquidityRemoved, error) {
	event := new(LockReleaseTokenPoolLiquidityRemoved)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLiquidityTransferredIterator struct {
	Event *LockReleaseTokenPoolLiquidityTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLiquidityTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolLiquidityTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLiquidityTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLiquidityTransferred struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLiquidityTransferred(opts *bind.FilterOpts, from []common.Address) (*LockReleaseTokenPoolLiquidityTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LiquidityTransferred", fromRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLiquidityTransferredIterator{contract: _LockReleaseTokenPool.contract, event: "LiquidityTransferred", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLiquidityTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityTransferred, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LiquidityTransferred", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLiquidityTransferred)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLiquidityTransferred(log types.Log) (*LockReleaseTokenPoolLiquidityTransferred, error) {
	event := new(LockReleaseTokenPoolLiquidityTransferred)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LiquidityTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolLockedOrBurnedIterator struct {
	Event *LockReleaseTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolLockedOrBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolLockedOrBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolLockedOrBurnedIterator{contract: _LockReleaseTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolLockedOrBurned)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*LockReleaseTokenPoolLockedOrBurned, error) {
	event := new(LockReleaseTokenPoolLockedOrBurned)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *LockReleaseTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolOutboundRateLimitConsumedIterator{contract: _LockReleaseTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolOutboundRateLimitConsumed)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolOutboundRateLimitConsumed, error) {
	event := new(LockReleaseTokenPoolOutboundRateLimitConsumed)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolPoolFeeWithdrawnIterator struct {
	Event *LockReleaseTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolPoolFeeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolPoolFeeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*LockReleaseTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolPoolFeeWithdrawnIterator{contract: _LockReleaseTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolPoolFeeWithdrawn)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*LockReleaseTokenPoolPoolFeeWithdrawn, error) {
	event := new(LockReleaseTokenPoolPoolFeeWithdrawn)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRateLimitAdminRoleGrantedIterator struct {
	Event *LockReleaseTokenPoolRateLimitAdminRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRateLimitAdminRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRateLimitAdminRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolRateLimitAdminRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolRateLimitAdminRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRateLimitAdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRateLimitAdminRoleGranted struct {
	Account common.Address
	Raw     types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*LockReleaseTokenPoolRateLimitAdminRoleGrantedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRateLimitAdminRoleGrantedIterator{contract: _LockReleaseTokenPool.contract, event: "RateLimitAdminRoleGranted", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRateLimitAdminRoleGranted)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRateLimitAdminRoleGranted(log types.Log) (*LockReleaseTokenPoolRateLimitAdminRoleGranted, error) {
	event := new(LockReleaseTokenPoolRateLimitAdminRoleGranted)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRateLimitAdminRoleRevokedIterator struct {
	Event *LockReleaseTokenPoolRateLimitAdminRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRateLimitAdminRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRateLimitAdminRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolRateLimitAdminRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolRateLimitAdminRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRateLimitAdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRateLimitAdminRoleRevoked struct {
	Account common.Address
	Raw     types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*LockReleaseTokenPoolRateLimitAdminRoleRevokedIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRateLimitAdminRoleRevokedIterator{contract: _LockReleaseTokenPool.contract, event: "RateLimitAdminRoleRevoked", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRateLimitAdminRoleRevoked)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRateLimitAdminRoleRevoked(log types.Log) (*LockReleaseTokenPoolRateLimitAdminRoleRevoked, error) {
	event := new(LockReleaseTokenPoolRateLimitAdminRoleRevoked)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRebalancerSetIterator struct {
	Event *LockReleaseTokenPoolRebalancerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRebalancerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRebalancerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolRebalancerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolRebalancerSetIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRebalancerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRebalancerSet struct {
	OldRebalancer common.Address
	NewRebalancer common.Address
	Raw           types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRebalancerSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolRebalancerSetIterator, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RebalancerSet")
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRebalancerSetIterator{contract: _LockReleaseTokenPool.contract, event: "RebalancerSet", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRebalancerSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRebalancerSet) (event.Subscription, error) {

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RebalancerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRebalancerSet)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RebalancerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRebalancerSet(log types.Log) (*LockReleaseTokenPoolRebalancerSet, error) {
	event := new(LockReleaseTokenPoolRebalancerSet)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RebalancerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolReleasedOrMintedIterator struct {
	Event *LockReleaseTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolReleasedOrMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolReleasedOrMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolReleasedOrMintedIterator{contract: _LockReleaseTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolReleasedOrMinted)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*LockReleaseTokenPoolReleasedOrMinted, error) {
	event := new(LockReleaseTokenPoolReleasedOrMinted)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRemotePoolAddedIterator struct {
	Event *LockReleaseTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolRemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRemotePoolAddedIterator{contract: _LockReleaseTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRemotePoolAdded)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*LockReleaseTokenPoolRemotePoolAdded, error) {
	event := new(LockReleaseTokenPoolRemotePoolAdded)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRemotePoolRemovedIterator struct {
	Event *LockReleaseTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolRemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRemotePoolRemovedIterator{contract: _LockReleaseTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRemotePoolRemoved)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*LockReleaseTokenPoolRemotePoolRemoved, error) {
	event := new(LockReleaseTokenPoolRemotePoolRemoved)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRoleAdminChangedIterator struct {
	Event *LockReleaseTokenPoolRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LockReleaseTokenPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRoleAdminChangedIterator{contract: _LockReleaseTokenPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRoleAdminChanged)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRoleAdminChanged(log types.Log) (*LockReleaseTokenPoolRoleAdminChanged, error) {
	event := new(LockReleaseTokenPoolRoleAdminChanged)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRoleGrantedIterator struct {
	Event *LockReleaseTokenPoolRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LockReleaseTokenPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRoleGrantedIterator{contract: _LockReleaseTokenPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRoleGranted)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRoleGranted(log types.Log) (*LockReleaseTokenPoolRoleGranted, error) {
	event := new(LockReleaseTokenPoolRoleGranted)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolRoleRevokedIterator struct {
	Event *LockReleaseTokenPoolRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LockReleaseTokenPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolRoleRevokedIterator{contract: _LockReleaseTokenPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolRoleRevoked)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseRoleRevoked(log types.Log) (*LockReleaseTokenPoolRoleRevoked, error) {
	event := new(LockReleaseTokenPoolRoleRevoked)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *LockReleaseTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolTokenTransferFeeConfigDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolTokenTransferFeeConfigDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _LockReleaseTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolTokenTransferFeeConfigDeleted)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*LockReleaseTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(LockReleaseTokenPoolTokenTransferFeeConfigDeleted)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *LockReleaseTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockReleaseTokenPoolTokenTransferFeeConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(LockReleaseTokenPoolTokenTransferFeeConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type LockReleaseTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _LockReleaseTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _LockReleaseTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(LockReleaseTokenPoolTokenTransferFeeConfigUpdated)
				if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_LockReleaseTokenPool *LockReleaseTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*LockReleaseTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(LockReleaseTokenPoolTokenTransferFeeConfigUpdated)
	if err := _LockReleaseTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
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

func (LockReleaseTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (LockReleaseTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (LockReleaseTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xece8a336aec3d0587372c99a62c7158c83d7419e28f8c519094cf44763b00e7d")
}

func (LockReleaseTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (LockReleaseTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (LockReleaseTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (LockReleaseTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (LockReleaseTokenPoolDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (LockReleaseTokenPoolDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (LockReleaseTokenPoolDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (LockReleaseTokenPoolDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (LockReleaseTokenPoolDynamicConfigSet) Topic() common.Hash {
	return common.HexToHash("0x78c5af2c6ab8d53b1850f16dd49fb61b0c1fef46835b922a40e3ce1f623f0238")
}

func (LockReleaseTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
}

func (LockReleaseTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (LockReleaseTokenPoolLiquidityAdded) Topic() common.Hash {
	return common.HexToHash("0xc17cea59c2955cb181b03393209566960365771dbba9dc3d510180e7cb312088")
}

func (LockReleaseTokenPoolLiquidityRemoved) Topic() common.Hash {
	return common.HexToHash("0xc2c3f06e49b9f15e7b4af9055e183b0d73362e033ad82a07dec9bf9840171719")
}

func (LockReleaseTokenPoolLiquidityTransferred) Topic() common.Hash {
	return common.HexToHash("0x6fa7abcf1345d1d478e5ea0da6b5f26a90eadb0546ef15ed3833944fbfd1db62")
}

func (LockReleaseTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (LockReleaseTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (LockReleaseTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (LockReleaseTokenPoolRateLimitAdminRoleGranted) Topic() common.Hash {
	return common.HexToHash("0xf7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e389")
}

func (LockReleaseTokenPoolRateLimitAdminRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b")
}

func (LockReleaseTokenPoolRebalancerSet) Topic() common.Hash {
	return common.HexToHash("0x64187bd7b97e66658c91904f3021d7c28de967281d18b1a20742348afdd6a6b3")
}

func (LockReleaseTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (LockReleaseTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (LockReleaseTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (LockReleaseTokenPoolRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (LockReleaseTokenPoolRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (LockReleaseTokenPoolRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (LockReleaseTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (LockReleaseTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
}

func (_LockReleaseTokenPool *LockReleaseTokenPool) Address() common.Address {
	return _LockReleaseTokenPool.address
}

type LockReleaseTokenPoolInterface interface {
	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	DefaultAdmin(opts *bind.CallOpts) (common.Address, error)

	DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error)

	DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error)

	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetDynamicConfig(opts *bind.CallOpts) (GetDynamicConfig,

		error)

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

	HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error)

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

	CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error)

	GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error)

	ProvideLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error)

	RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	SetDynamicConfig(opts *bind.TransactOpts, router common.Address, thresholdAmountForAdditionalCCVs *big.Int) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, rebalancer common.Address) (*types.Transaction, error)

	TransferLiquidity(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error)

	WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*LockReleaseTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*LockReleaseTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*LockReleaseTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*LockReleaseTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*LockReleaseTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*LockReleaseTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*LockReleaseTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*LockReleaseTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*LockReleaseTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*LockReleaseTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*LockReleaseTokenPoolConfigChanged, error)

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*LockReleaseTokenPoolDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*LockReleaseTokenPoolDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*LockReleaseTokenPoolDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*LockReleaseTokenPoolDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*LockReleaseTokenPoolDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*LockReleaseTokenPoolDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*LockReleaseTokenPoolDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*LockReleaseTokenPoolDefaultAdminTransferScheduled, error)

	FilterDynamicConfigSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolDynamicConfigSetIterator, error)

	WatchDynamicConfigSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolDynamicConfigSet) (event.Subscription, error)

	ParseDynamicConfigSet(log types.Log) (*LockReleaseTokenPoolDynamicConfigSet, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*LockReleaseTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*LockReleaseTokenPoolFinalityConfigUpdated, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolInboundRateLimitConsumed, error)

	FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityAddedIterator, error)

	WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityAdded, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityAdded(log types.Log) (*LockReleaseTokenPoolLiquidityAdded, error)

	FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address, amount []*big.Int) (*LockReleaseTokenPoolLiquidityRemovedIterator, error)

	WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityRemoved, provider []common.Address, amount []*big.Int) (event.Subscription, error)

	ParseLiquidityRemoved(log types.Log) (*LockReleaseTokenPoolLiquidityRemoved, error)

	FilterLiquidityTransferred(opts *bind.FilterOpts, from []common.Address) (*LockReleaseTokenPoolLiquidityTransferredIterator, error)

	WatchLiquidityTransferred(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLiquidityTransferred, from []common.Address) (event.Subscription, error)

	ParseLiquidityTransferred(log types.Log) (*LockReleaseTokenPoolLiquidityTransferred, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*LockReleaseTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*LockReleaseTokenPoolOutboundRateLimitConsumed, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*LockReleaseTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*LockReleaseTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*LockReleaseTokenPoolRateLimitAdminRoleGrantedIterator, error)

	WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error)

	ParseRateLimitAdminRoleGranted(log types.Log) (*LockReleaseTokenPoolRateLimitAdminRoleGranted, error)

	FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*LockReleaseTokenPoolRateLimitAdminRoleRevokedIterator, error)

	WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error)

	ParseRateLimitAdminRoleRevoked(log types.Log) (*LockReleaseTokenPoolRateLimitAdminRoleRevoked, error)

	FilterRebalancerSet(opts *bind.FilterOpts) (*LockReleaseTokenPoolRebalancerSetIterator, error)

	WatchRebalancerSet(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRebalancerSet) (event.Subscription, error)

	ParseRebalancerSet(log types.Log) (*LockReleaseTokenPoolRebalancerSet, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*LockReleaseTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*LockReleaseTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*LockReleaseTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*LockReleaseTokenPoolRemotePoolRemoved, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LockReleaseTokenPoolRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*LockReleaseTokenPoolRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LockReleaseTokenPoolRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*LockReleaseTokenPoolRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LockReleaseTokenPoolRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*LockReleaseTokenPoolRoleRevoked, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*LockReleaseTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*LockReleaseTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*LockReleaseTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *LockReleaseTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*LockReleaseTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
