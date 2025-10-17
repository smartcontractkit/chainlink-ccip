// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_with_lock_release_flag_token_pool

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
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
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

var BurnMintWithLockReleaseFlagTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMITER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x61010080604052346103ee57617144803803809161001d8285610483565b833981019060a0818303126103ee5780516001600160a01b038116908190036103ee5761004c602083016104a6565b60408301519091906001600160401b0381116103ee5783019380601f860112156103ee578451946001600160401b03861161046d578560051b9060208201966100986040519889610483565b87526020808801928201019283116103ee57602001905b828210610455575050506100d160806100ca606086016104b4565b94016104b4565b92331561043f57600180546001600160d01b031690556002546001600160a01b03811661042e576001600160a01b03191633908117600255610112906104f2565b508115801561041d575b801561040c575b6103fb578160209160049360805260c0526040519283809263313ce56760e01b82525afa600091816103ba575b5061038f575b5060a052600580546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610271575b604051616a3390816106f18239608051818181611751015281816119c301528181611b5c015281816121f10152818161244c0152818161257d015281816130b9015281816132d301528181613449015281816135950152818161373e01528181613f8501528181613fd20152818161410801528181614c6a0152615bb1015260a051818181611a3f01528181613f09015281816155d0015281816156bb0152615827015260c051818181610c24015281816117df0152818161227f015281816131470152613624015260e051818181610bdf01528181611822015281816122c20152612e7d0152f35b60405160206102808183610483565b60008252600036813760e0511561037e5760005b82518110156102fb576001906001600160a01b036102b282866104c8565b5116836102be82610598565b6102cb575b505001610294565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138836102c3565b50905060005b8251811015610375576001906001600160a01b0361031f82866104c8565b5116801561036f578361033182610696565b61033f575b50505b01610301565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a13883610336565b50610339565b50505038610188565b6335f4a7b360e01b60005260046000fd5b60ff1660ff82168181036103a35750610156565b6332ad3e0760e11b60005260045260245260446000fd5b9091506020813d6020116103f3575b816103d660209383610483565b810103126103ee576103e7906104a6565b9038610150565b600080fd5b3d91506103c9565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610123565b506001600160a01b0384161561011c565b631fe1e13d60e11b60005260046000fd5b636116401160e11b600052600060045260246000fd5b60208091610462846104b4565b8152019101906100af565b634e487b7160e01b600052604160045260246000fd5b601f909101601f19168101906001600160401b0382119082101761046d57604052565b519060ff821682036103ee57565b51906001600160a01b03821682036103ee57565b80518210156104dc5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b0381166000908152600080516020617124833981519152602052604090205460ff1661057a576001600160a01b0316600081815260008051602061712483398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b80548210156104dc5760005260206000200190600090565b600081815260046020526040902054801561068f5760001981018181116106795760035460001981019190821161067957818103610628575b505050600354801561061257600019016105ec816003610580565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b61066161063961064a936003610580565b90549060031b1c9283926003610580565b819391549060031b91821b91600019901b19161790565b905560005260046020526040600020553880806105d1565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260046020526040600020541560001461057a576003546801000000000000000081101561046d576106d761064a8260018594016003556003610580565b905560035490600052600460205260406000205560019056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146143cb57508063022d63fb146143ad5780630aa6220b146142f15780630bd7c46d14614283578063164e68de14614057578063181f5a7714613ff657806321df0da714613fb2578063240028e814613f5b578063248a9ca314613f2d57806324f65ee714613eef5780632a10097b14613c5b5780632c286daf14613b345780632f2ff15d14613aeb57806336568abe146139b057806337b19247146138a9578063390775371461352a578063489a68f2146130245780634c5ef0ed14612fdd5780634f71592c14612fa857806354c8a4f314612e275780635df45a3714612e0457806362ddd3c414612d60578063634e93da14612c40578063649a5ec714612a2c578063791e5a10146129f1578063804ba5a91461298b57806384ef8ffc1461291e5780638926f54f146129455780638da5cb5b1461291e57806391d14854146128d2578063962d40201461271a5780639a4575b91461219d5780639f68f67314612165578063a1eda53c14612102578063a217fddf146120e6578063a42a7b8b14611f9d578063a7cd63b714611f1c578063acfecf9114611db1578063af58d59f14611d68578063b0f479a114611d41578063b1c71c65146116d3578063b79019b51461140f578063b7946580146113d6578063c0d78655146112f5578063c4bffe2b146111e8578063c75eea9c14611140578063cc8463c814611115578063cefc142914611020578063cf6eefb714610fcd578063cf7401f314610dea578063d547741f14610d70578063d602b9fd14610cf4578063da90a9f314610c48578063dc0bd97114610c04578063e0351e1314610bc7578063e58d80c714610b5d5763e8a1da171461029457600080fd5b34610b5a576102a2366148dc565b9290939182805282602052604083206001600160a01b03331660005260205260ff6040600020541615610b325782918593915b80831061099d575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b85821015610999578160051b850135818112156109955785019061012082360312610995576040519561034087614658565b823567ffffffffffffffff81168103610990578752602083013567ffffffffffffffff811161098c5783019536601f8801121561098c5786359661038388614bc2565b97610391604051998a614690565b8089526020808a019160051b830101903682116109885760208301905b828210610955575050505060208801968752604084013567ffffffffffffffff8111610951576103e19036908601614810565b9860408901998a5261040b6103f936606088016149f0565b9560608b0196875260c03691016149f0565b9660808a0197885261041d8651615e48565b6104278851615e48565b8a5151156109295761044367ffffffffffffffff8b51166168ab565b156108f25767ffffffffffffffff8a5116815260086020526040812061058387516fffffffffffffffffffffffffffffffff6040820151169061053e6fffffffffffffffffffffffffffffffff602083015116915115158360806040516104a981614658565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106a989516fffffffffffffffffffffffffffffffff604082015116906106646fffffffffffffffffffffffffffffffff602083015116915115158360806040516105cd81614658565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff82116108c5576106cc8354614d16565b601f811161088a575b50602090601f831160011461082757610705929185918361081c575b50506000198260011b9260031b1c19161790565b90555b805b89518051821015610740579061073a600192610733838f67ffffffffffffffff90511692614d02565b5190615849565b0161070a565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29295939661080e67ffffffffffffffff600197949c51169251935191516107da6107a5604051968796875261010060208801526101008701906146cf565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a101909394929161030e565b015190508f806106f1565b8385528185209190601f198416865b8181106108725750908460019594939210610859575b505050811b019055610708565b015160001960f88460031b161c191690558e808061084c565b92936020600181928786015181550195019301610836565b6108b59084865260208620601f850160051c810191602086106108bb575b601f0160051c0190614ff3565b8e6106d5565b90915081906108a8565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff8111610984576020916109798392833691890101614810565b8152019101906103ae565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff6109bd6109b8878588614aba565b614a76565b16956109c88761662f565b15610b065786845260086020526109e4600560408620016164d3565b94845b8651811015610a1d576001908987526008602052610a1660056040892001610a0f838b614d02565b51906166e2565b50016109e7565b5093945094909580855260086020526005604086208681558660018201558660028201558660038201558660048201610a568154614d16565b80610ac5575b5050500180549086815581610aa7575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946102d5565b865260208620908101905b81811015610a6c57868155600101610ab2565b601f8111600114610adb5750555b868a80610a5c565b81835260208320610af691601f01861c810190600101614ff3565b8082528160208120915555610ad3565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6004837f2b5c74de000000000000000000000000000000000000000000000000000000008152fd5b80fd5b5034610b5a576020600319360112610b5a576001600160a01b036040610b81614595565b927f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af815280602052209116600052602052602060ff604060002054166040519015158152f35b5034610b5a5780600319360112610b5a5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b5034610b5a5780600319360112610b5a5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610b5a576020600319360112610b5a57610c62614595565b81805281602052604082206001600160a01b03331660005260205260ff6040600020541615610ccc57602081610cb87fd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b93615dc2565b506001600160a01b0360405191168152a180f35b6004827f2b5c74de000000000000000000000000000000000000000000000000000000008152fd5b5034610b5a5780600319360112610b5a57610d0d615065565b600180547fffffffffffff0000000000000000000000000000000000000000000000000000811690915560a01c65ffffffffffff16610d495780f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051098180a180f35b5034610b5a576040600319360112610b5a57600435610d8d6145ab565b908015610dc2579081610db9610db4610dbe94600052600060205260016040600020015490565b6150d1565b615dec565b5080f35b6004837f3fc3c27a000000000000000000000000000000000000000000000000000000008152fd5b5034610b5a5760e0600319360112610b5a57610e04614794565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610b5a57604051610e3b81614674565b6024358015158103610fc95781526044356fffffffffffffffffffffffffffffffff81168103610fc95760208201526064356fffffffffffffffffffffffffffffffff81168103610fc957604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c360112610fc55760405190610ec282614674565b608435801515810361099557825260a4356fffffffffffffffffffffffffffffffff8116810361099557602083015260c4356fffffffffffffffffffffffffffffffff811681036109955760408301527f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af835282602052604083206001600160a01b03331660005260205260ff604060002054161580610f9a575b610f6e57610f6b9293615a77565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5082805282602052604083206001600160a01b03331660005260205260ff6040600020541615610f5d565b5080fd5b8280fd5b5034610b5a5780600319360112610b5a57604065ffffffffffff6110076001549065ffffffffffff6001600160a01b0383169260a01c1690565b6001600160a01b03849392935193168352166020820152f35b5034610b5a5780600319360112610b5a576001546001600160a01b031633036110e9576001546001600160a01b0381169060a01c65ffffffffffff16801580156110df575b6110b45750611088906110826001600160a01b0360025416615d6e565b5061515b565b507fffffffffffff00000000000000000000000000000000000000000000000000006001541660015580f35b7f19ca5ebb000000000000000000000000000000000000000000000000000000008352600452602482fd5b5042811015611065565b807fc22c8022000000000000000000000000000000000000000000000000000000006024925233600452fd5b5034610b5a5780600319360112610b5a57602061113061502c565b65ffffffffffff60405191168152f35b5034610b5a576020600319360112610b5a5761118b61118660406111e49367ffffffffffffffff61116f614794565b611177614e4a565b50168152600860205220614e75565b615c39565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b5034610b5a5780600319360112610b5a57604051906006548083528260208101600684526020842092845b8181106112dc57505061122892500383614690565b815161124c61123682614bc2565b916112446040519384614690565b808352614bc2565b91601f19602083019301368437805b845181101561128d578067ffffffffffffffff61127a60019388614d02565b51166112868286614d02565b520161125b565b50925090604051928392602084019060208552518091526040840192915b8181106112b9575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016112ab565b8454835260019485019487945060209093019201611213565b5034610b5a576020600319360112610b5a5761130f614595565b81805281602052604082206001600160a01b03331660005260205260ff6040600020541615610ccc576001600160a01b031680156113ae5760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160055490807fffffffffffffffffffffffff00000000000000000000000000000000000000008316176005556001600160a01b038351921682526020820152a180f35b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b5034610b5a576020600319360112610b5a576111e46113fb6113f6614794565b61500a565b6040519182916020835260208301906146cf565b5034610b5a576020600319360112610b5a5760043567ffffffffffffffff8111610fc557611441903690600401614710565b82805282602052604083206001600160a01b03331660005260205260ff6040600020541615610b3257825b818110611477578380f35b6114856109b8828486614edb565b61149d611493838587614edb565b6020810190614f1b565b907fb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c6114d76114cd86888a614edb565b6040810190614f1b565b9190926114ed6114e8368784614bda565b615ccb565b6114fb6114e8368587614bda565b604051946115088661463c565b611513368284614bda565b865261155d67ffffffffffffffff61152c368789614bda565b9860208901998a521695869561154f604051958695604087526040870191614f6f565b918483036020860152614f6f565b0390a28652600d60205260408620905180519067ffffffffffffffff82116116a6576801000000000000000082116116a657602090835483855580841061168c575b500182885260208820885b83811061166f575050505060010190519081519167ffffffffffffffff831161164257680100000000000000008311611642576020908254848455808510611628575b500190865260208620865b83811061160b575050505060010161146c565b60019060206001600160a01b0385511694019381840155016115f8565b83895282892061163c918101908601614ff3565b386115ed565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60019060206001600160a01b0385511694019381840155016115aa565b848a52828a206116a0918101908501614ff3565b3861159f565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b5034610b5a576060600319360112610b5a5760043567ffffffffffffffff8111610fc55760a06003198236030112610fc55761170d614741565b9060443567ffffffffffffffff81116109955761172e903690600401614810565b50611737614ce9565b50608481019161174683614aca565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603611d0457602482019177ffffffffffffffff0000000000000000000000000000000061179f84614a76565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611c28578691611cd5575b50611cad5761182060448201614aca565b7f0000000000000000000000000000000000000000000000000000000000000000611c5e575b5067ffffffffffffffff61185984614a76565b16611871816000526007602052604060002054151590565b15611c335760206001600160a01b0360055416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611c28578690611bdf575b6001600160a01b039150163303611bb357606461ffff91013591169384151593848095611ba4575b15611b075761ffff600a5416808710611ad75750611a9b955061193f61192f61191586614a76565b67ffffffffffffffff16600052600b602052604060002090565b8461193984614aca565b9161627a565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff61197b61197587614a76565b93614aca565b604080516001600160a01b03929092168252602082018790529190931692a25b508092611aa5575b506113f6611a37916119b484615ba7565b6119bd81614a76565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2614a76565b9060405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152611a73604082614690565b60405192611a808461463c565b8352602083015260405192839260408452604084019061499c565b9060208301520390f35b611a37919250611acf6113f691612710611ac861ffff600a5460101c1683614fb1565b0490615c2c565b9291506119a3565b82604491887fe08f03ef000000000000000000000000000000000000000000000000000000008352600452602452fd5b50611a9b945067ffffffffffffffff611b1f84614a76565b1680825260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448380611b84604086206001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692839161627a565b604080516001600160a01b039290921682526020820192909252a261199b565b5061ffff600a541615156118ed565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611c20575b81611bf960209383614690565b8101031261098c57516001600160a01b038116810361098c576001600160a01b03906118c5565b3d9150611bec565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b6001600160a01b0316611c7e816000526004602052604060002054151590565b611846577fd0d25976000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611cf7915060203d602011611cfd575b611cef8183614690565b81019061579b565b3861180f565b503d611ce5565b6024846001600160a01b03611d1886614aca565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b5034610b5a5780600319360112610b5a5760206001600160a01b0360055416604051908152f35b5034610b5a576020600319360112610b5a5761118b611186600260406111e49467ffffffffffffffff611d99614794565b611da1614e4a565b5016815260086020522001614e75565b5034610b5a57611dc03661492a565b9183805283602052604084206001600160a01b03331660005260205260ff6040600020541615611ef45767ffffffffffffffff1691611e0c836000526007602052604060002054151590565b15611ec8578284526008602052611e3b60056040862001611e2e3684866147d9565b60208151910120906166e2565b15611e8057907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691611e7a604051928392602084526020840191614e29565b0390a280f35b82611ec4836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191614e29565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6004847f2b5c74de000000000000000000000000000000000000000000000000000000008152fd5b5034610b5a5780600319360112610b5a5760405160038054808352908352909160208301917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b915b818110611f87576111e485611f7b81870382614690565b60405191829182614899565b8254845260209093019260019283019201611f64565b5034610b5a576020600319360112610b5a5767ffffffffffffffff611fc0614794565b1681526008602052611fd7600560408320016164d3565b8051601f19611ffe611fe883614bc2565b92611ff66040519485614690565b808452614bc2565b01835b8181106120d5575050825b8251811015612052578061202260019285614d02565b518552600960205261203660408620614d69565b6120408285614d02565b5261204b8184614d02565b500161200c565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b82821061208a57505050500390f35b919360206120c5827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0600195979984950301865288516146cf565b960192019201859493919261207b565b806060602080938601015201612001565b5034610b5a5780600319360112610b5a57602090604051908152f35b5034610b5a5780600319360112610b5a576002548060d01c918215158061215b575b15612152575060a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b91505080612136565b5042831015612124565b5034610b5a57611f7b600160406111e49367ffffffffffffffff6121883661482e565b505050509050168152600d6020522001614b6c565b5034610b5a576020600319360112610b5a576004359067ffffffffffffffff8211610b5a5760a06003198336030112610b5a576121d8614ce9565b50608482016121e681614aca565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361270657602483019277ffffffffffffffff0000000000000000000000000000000061223f85614a76565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561263a5784916126e7575b506126bf576122c060448201614aca565b7f0000000000000000000000000000000000000000000000000000000000000000612670575b5067ffffffffffffffff6122f985614a76565b16612311816000526007602052604060002054151590565b156126455760206001600160a01b0360055416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561263a5784906125f1575b6001600160a01b0391501633036125c5576064013591801561251f5761ffff600a5416806124ef5750508261248f926113f6926123b46123aa6119156111e498614a76565b8361193984614aca565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff6123ea61197586614a76565b604080516001600160a01b03929092168252602082018690529190931692a25b61241381615ba7565b7ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1067ffffffffffffffff61244684614a76565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681523360208201529081019490945216918060608101611a2f565b6040517ffa7c07de000000000000000000000000000000000000000000000000000000006020820152602081526124c7604082614690565b604051916124d48361463c565b8252602082015260405191829160208352602083019061499c565b604492507fe08f03ef00000000000000000000000000000000000000000000000000000000825281600452602452fd5b50506113f68261248f9267ffffffffffffffff61253e6111e496614a76565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894482806125a560406000206001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692839161627a565b604080516001600160a01b039290921682526020820192909252a261240a565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011612632575b8161260b60209383614690565b8101031261099557516001600160a01b0381168103610995576001600160a01b0390612365565b3d91506125fe565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b6001600160a01b0316612690816000526004602052604060002054151590565b6122e6577fd0d25976000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612700915060203d602011611cfd57611cef8183614690565b386122af565b906001600160a01b03611d18602493614aca565b5034610b5a576060600319360112610b5a5760043567ffffffffffffffff8111610fc55761274c903690600401614710565b60243567ffffffffffffffff81116109955761276c90369060040161496b565b60449291923567ffffffffffffffff811161098c5761278f90369060040161496b565b9190927f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af875286602052604087206001600160a01b03331660005260205260ff6040600020541615806128a7575b61287b57818114801590612871575b61284957865b8181106127fd578780f35b806128436128116109b8600194868c614aba565b61281c83878b614cd9565b61283d61283561282d868b8d614cd9565b9236906149f0565b9136906149f0565b91615a77565b016127f2565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b50828114156127ec565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b5086805286602052604087206001600160a01b03331660005260205260ff60406000205416156127dd565b5034610b5a576040600319360112610b5a576001600160a01b0360406128f66145ab565b92600435815280602052209116600052602052602060ff604060002054166040519015158152f35b5034610b5a5780600319360112610b5a5760206001600160a01b0360025416604051908152f35b5034610b5a576020600319360112610b5a57602061298167ffffffffffffffff61296d614794565b166000526007602052604060002054151590565b6040519015158152f35b5034610b5a576020600319360112610b5a5760043567ffffffffffffffff8111610fc5576129bd903690600401614763565b9082805282602052604083206001600160a01b03331660005260205260ff6040600020541615610b325790610f6b9161524a565b5034610b5a5780600319360112610b5a5760206040517f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8152f35b5034610b5a576020600319360112610b5a5760043565ffffffffffff8116808203610fc957612a59615065565b612a6242616489565b9065ffffffffffff612a7261502c565b1680821115612bd857507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9291612abe9162069780811015612bc75765ffffffffffff905b1690615a59565b906002548060d01c80612b44575b5050600280546001600160a01b031660a083901b79ffffffffffff0000000000000000000000000000000000000000161760d084901b7fffffffffffff0000000000000000000000000000000000000000000000000000161790556040805165ffffffffffff9283168152919092166020820152a180f35b421115612b9d5779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b3880612acc565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58480a1612b96565b5065ffffffffffff62069780612ab7565b0365ffffffffffff8111612c13577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9291612abe9190615a59565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b5034610b5a576020600319360112610b5a57612c5a614595565b612c62615065565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020612c9f612c9142616489565b612c9961502c565b90615a59565b65ffffffffffff6001600160a01b03612cce6001549065ffffffffffff6001600160a01b0383169260a01c1690565b9690501694600154867fffffffffffff000000000000000000000000000000000000000000000000000079ffffffffffff00000000000000000000000000000000000000008660a01b169216171760015516612d37575b65ffffffffffff60405191168152a280f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051098580a1612d25565b5034610b5a57612d6f3661492a565b83809392945282602052604083206001600160a01b03331660005260205260ff6040600020541615610b325767ffffffffffffffff8216612dbd816000526007602052604060002054151590565b15612dd95750610f6b9293612dd39136916147d9565b90615849565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b5034610b5a5780600319360112610b5a576020612e1f614c2e565b604051908152f35b5034610b5a57612e36366148dc565b9392909183805283602052604084206001600160a01b03331660005260205260ff6040600020541615611ef457612e7b9291612e73913691614bda565b933691614bda565b7f000000000000000000000000000000000000000000000000000000000000000015612f8057815b8351811015612f0957806001600160a01b03612ec160019387614d02565b5116612ecc81616536565b612ed8575b5001612ea3565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a138612ed1565b5090805b8251811015610dbe57806001600160a01b03612f2b60019386614d02565b51168015612f7a57612f3c8161684b565b612f49575b505b01612f0d565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a184612f41565b50612f43565b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b5034610b5a57611f7b60406111e49267ffffffffffffffff612fc93661482e565b505050509050168152600d60205220614b6c565b5034610b5a576040600319360112610b5a57612ff7614794565b906024359067ffffffffffffffff8211610b5a5760206129818461301e3660048701614810565b90614b2f565b5034610b5a576040600319360112610b5a5760043567ffffffffffffffff8111610fc55780600401916101006003198336030112610b5a57613064614741565b918160405161307281614620565b5261309f61309561309061308960c4850188614ade565b36916147d9565b6157b3565b60648301356156b8565b9260848201906130ae82614aca565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361351657602483019577ffffffffffffffff0000000000000000000000000000000061310788614a76565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611c285786916134f7575b50611cad5767ffffffffffffffff61318e88614a76565b166131a6816000526007602052604060002054151590565b15611c335760206001600160a01b0360055416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611c285786916134d8575b5015611bb35761321087614a76565b9061322660a486019261301e6130898585614ade565b156134915750604493929161ffff161590506133f65767ffffffffffffffff61324e87614a76565b168452600c602052613267604085208661193984614aca565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff767ffffffffffffffff61329d61197589614a76565b604080516001600160a01b03929092168252602082018990529190931692a25b01926132c884614aca565b916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692833b15610fc5576040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b0391909116600482015260248101859052818160448183885af180156133eb576133d6575b505067ffffffffffffffff6020946133bd8561338c6119757ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc096614a76565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b0390a2806040516133cd81614620565b52604051908152f35b6133e1828092614690565b610b5a578061334d565b6040513d84823e3d90fd5b5067ffffffffffffffff61340986614a76565b1680845260086020527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c8580613471600260408920016001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692839161627a565b604080516001600160a01b039290921682526020820192909252a26132bd565b61349b9250614ade565b611ec46040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191614e29565b6134f1915060203d602011611cfd57611cef8183614690565b38613201565b613510915060203d602011611cfd57611cef8183614690565b38613177565b6024846001600160a01b03611d1885614aca565b5034610b5a576020600319360112610b5a5760043567ffffffffffffffff8111610fc557806004016101006003198336030112610fc9578260405161356e81614620565b5261357c60648301356155ce565b916084810161358a81614aca565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036138955750602481019177ffffffffffffffff000000000000000000000000000000006135e484614a76565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115611c28578691613876575b50611cad5767ffffffffffffffff61366b84614a76565b16613683816000526007602052604060002054151590565b15611c335760206001600160a01b0360055416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115611c28578691613857575b5015611bb3576136ed83614a76565b9061370360a484019261301e6130898585614ade565b15613491575050906044839267ffffffffffffffff61372184614a76565b168087526008602052613766600260408920016001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001696879161627a565b604080516001600160a01b0387168152602081018890527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a201906137ac82614aca565b85843b15610b5a576040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b03929092166004830152602482018690528160448183885af18015611c28578561338c61197567ffffffffffffffff9560209a7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc098966133bd96613847575b5050614a76565b8161385191614690565b38613840565b613870915060203d602011611cfd57611cef8183614690565b386136de565b61388f915060203d602011611cfd57611cef8183614690565b38613654565b846001600160a01b03611d18602493614aca565b5034610b5a5760a0600319360112610b5a576138c3614595565b5060243567ffffffffffffffff8116809103610fc55760443567ffffffffffffffff8111610fc95760031960a09136030112610fc557613901614752565b5060843567ffffffffffffffff8111610fc9579160409161392860809436906004016147ab565b50508160608451613938816145d5565b828152826020820152828682015201528152600e602052206040519061395d826145d5565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b5034610b5a576040600319360112610b5a576004356139cd6145ab565b90801580613ace575b613a1d575b336001600160a01b038316036139f55790610dbe91615dec565b6004837f6697b232000000000000000000000000000000000000000000000000000000008152fd5b60015465ffffffffffff60a082901c16906001600160a01b031615801590613abe575b8015613aac575b613a7957507fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff600154166001556139db565b7f19ca5ebb00000000000000000000000000000000000000000000000000000000845265ffffffffffff16600452602483fd5b504265ffffffffffff82161015613a47565b5065ffffffffffff811615613a40565b506001600160a01b03600254166001600160a01b038316146139d6565b5034610b5a576040600319360112610b5a57600435613b086145ab565b908015610dc2579081613b2f610db4610dbe94600052600060205260016040600020015490565b6151d3565b5034610b5a576060600319360112610b5a576004359061ffff8216809203610b5a57613b5e614741565b9160443567ffffffffffffffff8111610fc957613b7f903690600401614763565b9083805283602052604084206001600160a01b033316855260205260ff60408520541615611ef45761ffff851694612710861015613c2f5791613c1b917f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba959693857fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600a549360101b1692161717600a5561524a565b60408051928352602083019190915290a180f35b602485877f95f3517a000000000000000000000000000000000000000000000000000000008252600452fd5b5034610b5a576040600319360112610b5a5760043567ffffffffffffffff8111610fc55736602382011215610fc557806004013567ffffffffffffffff8111610fc95760248201916024369160a08402010111610fc95760243567ffffffffffffffff811161099557613cd2903690600401614710565b91909284805284602052604085206001600160a01b033316865260205260ff60408620541615613ec75784805b838110613d6957509150505b818110613d16578380f35b8067ffffffffffffffff613d306109b86001948688614aba565b16808652600e6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201613d0b565b6001917f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70608085613eb7613e7d63ffffffff613eac613e7082613ea1613dc08f80613dba8f92836109b8918e614a37565b9a614a37565b604067ffffffffffffffff602083019a169c8d8152600e6020522083613de58b614a8b565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff00000000613e2287614a8b565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e614a8b565b60401b1696019e8f614a9c565b151560601b1695161716171717905582613e996040519a614aa9565b168952614aa9565b166020870152614aa9565b1660408401526149c6565b15156060820152a2018590613cff565b6004857f2b5c74de000000000000000000000000000000000000000000000000000000008152fd5b5034610b5a5780600319360112610b5a57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610b5a576020600319360112610b5a576020612e1f600435600052600060205260016040600020015490565b5034610b5a576020600319360112610b5a57602090613f78614595565b90506001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b5034610b5a5780600319360112610b5a5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610b5a5780600319360112610b5a57506111e4604051614019604082614690565b601b81527f4275726e4d696e74546f6b656e506f6f6c20312e362e332d646576000000000060208201526040519182916020835260208301906146cf565b5034610b5a576020600319360112610b5a57614071614595565b81805281602052604082206001600160a01b033316835260205260ff60408320541615610ccc576140a0614c2e565b90816140aa578280f35b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000602082019081526001600160a01b03831660248301526044808301859052825290614196906140fe606482614690565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016868060409586519461413a8887614690565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d1561427b573d9161417b836146b3565b9261418887519485614690565b83523d89602085013e61695a565b8051806141da575b50506001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992602092519485521692a238808280f35b906020806141ec93830101910161579b565b156141f857388061419e565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b60609161695a565b5034610b5a576020600319360112610b5a5761429d614595565b81805281602052604082206001600160a01b033316835260205260ff60408320541615610ccc57602081610cb87ff7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e38993615131565b5034610b5a5780600319360112610b5a5761430a615065565b6002548060d01c8061432a575b826001600160a01b036002541660025580f35b4211156143835779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b3880614317565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58180a161437c565b5034610b5a5780600319360112610b5a576020604051620697808152f35b905034610fc5576020600319360112610fc5576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610fc957602092507ff208a58f00000000000000000000000000000000000000000000000000000000811490811561456b575b8115614541575b8115614517575b81156144ed575b811561445d575b5015158152f35b7f3149878600000000000000000000000000000000000000000000000000000000811491508115614490575b5038614456565b7f7965db0b000000000000000000000000000000000000000000000000000000008114915081156144c3575b5038614489565b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386144bc565b7f01ffc9a7000000000000000000000000000000000000000000000000000000008114915061444f565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614448565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150614441565b7faff2afbf000000000000000000000000000000000000000000000000000000008114915061443a565b600435906001600160a01b038216820361099057565b602435906001600160a01b038216820361099057565b35906001600160a01b038216820361099057565b6080810190811067ffffffffffffffff8211176145f157604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff8211176145f157604052565b6040810190811067ffffffffffffffff8211176145f157604052565b60a0810190811067ffffffffffffffff8211176145f157604052565b6060810190811067ffffffffffffffff8211176145f157604052565b90601f601f19910116810190811067ffffffffffffffff8211176145f157604052565b67ffffffffffffffff81116145f157601f01601f191660200190565b919082519283825260005b8481106146fb575050601f19601f8460006020809697860101520116010190565b806020809284010151828286010152016146da565b9181601f840112156109905782359167ffffffffffffffff8311610990576020808501948460051b01011161099057565b6024359061ffff8216820361099057565b6064359061ffff8216820361099057565b9181601f840112156109905782359167ffffffffffffffff83116109905760208085019460e0850201011161099057565b6004359067ffffffffffffffff8216820361099057565b9181601f840112156109905782359167ffffffffffffffff8311610990576020838186019501011161099057565b9291926147e5826146b3565b916147f36040519384614690565b829481845281830111610990578281602093846000960137010152565b9080601f830112156109905781602061482b933591016147d9565b90565b60a0600319820112610990576004356001600160a01b0381168103610990579160243567ffffffffffffffff8116810361099057916044359160643561ffff8116810361099057916084359067ffffffffffffffff821161099057614895916004016147ab565b9091565b602060408183019282815284518094520192019060005b8181106148bd5750505090565b82516001600160a01b03168452602093840193909201916001016148b0565b60406003198201126109905760043567ffffffffffffffff8111610990578161490791600401614710565b929092916024359067ffffffffffffffff82116109905761489591600401614710565b9060406003198301126109905760043567ffffffffffffffff8116810361099057916024359067ffffffffffffffff821161099057614895916004016147ab565b9181601f840112156109905782359167ffffffffffffffff8311610990576020808501946060850201011161099057565b61482b9160206149b583516040845260408401906146cf565b9201519060208184039101526146cf565b3590811515820361099057565b35906fffffffffffffffffffffffffffffffff8216820361099057565b919082606091031261099057604051614a0881614674565b6040614a32818395614a19816149c6565b8552614a27602082016149d3565b6020860152016149d3565b910152565b9190811015614a475760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff811681036109905790565b3563ffffffff811681036109905790565b3580151581036109905790565b359063ffffffff8216820361099057565b9190811015614a475760051b0190565b356001600160a01b03811681036109905790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610990570180359067ffffffffffffffff82116109905760200191813603831361099057565b9067ffffffffffffffff61482b92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b818110614ba0575050614b9e92500383614690565b565b84546001600160a01b0316835260019485019487945060209093019201614b89565b67ffffffffffffffff81116145f15760051b60200190565b929190614be681614bc2565b93614bf46040519586614690565b602085838152019160051b810192831161099057905b828210614c1657505050565b60208091614c23846145c1565b815201910190614c0a565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115614ccd57600091614c9e575090565b90506020813d602011614cc5575b81614cb960209383614690565b81010312610990575190565b3d9150614cac565b6040513d6000823e3d90fd5b9190811015614a47576060020190565b60405190614cf68261463c565b60606020838281520152565b8051821015614a475760209160051b010190565b90600182811c92168015614d5f575b6020831014614d3057565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614d25565b9060405191826000825492614d7d84614d16565b8084529360018116908115614de95750600114614da2575b50614b9e92500383614690565b90506000929192526020600020906000915b818310614dcd575050906020614b9e9282010138614d95565b6020919350806001915483858901015201910190918492614db4565b60209350614b9e9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138614d95565b601f8260209493601f19938186528686013760008582860101520116010190565b60405190614e5782614658565b60006080838281528260208201528260408201528260608201520152565b90604051614e8281614658565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b9190811015614a475760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa181360301821215610990570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610990570180359067ffffffffffffffff821161099057602001918160051b3603831361099057565b9160209082815201919060005b818110614f895750505090565b9091926020806001926001600160a01b03614fa3886145c1565b168152019401929101614f7c565b81810292918115918404141715614fc457565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110614ffe575050565b60008155600101614ff3565b67ffffffffffffffff16600052600860205261482b6004604060002001614d69565b6002548060d01c801515908161505b575b50156150515760a01c65ffffffffffff1690565b5060015460d01c90565b905042113861503d565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff161561509e57565b7fe2517d3f0000000000000000000000000000000000000000000000000000000060005233600452600060245260446000fd5b80600052600060205260406000206001600160a01b03331660005260205260ff60406000205416156151005750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b61482b907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af615f8f565b600254906001600160a01b0382166151a95761482b917fffffffffffffffffffffffff00000000000000000000000000000000000000006001600160a01b0383169116176002556000615f8f565b7f3fc3c27a0000000000000000000000000000000000000000000000000000000060005260046000fd5b9081156151e4575b61482b91615f8f565b600254916001600160a01b0383166151a9577fffffffffffffffffffffffff00000000000000000000000000000000000000009092166001600160a01b038216176002556151db565b356fffffffffffffffffffffffffffffffff811681036109905790565b9160005b8281101561556a5760e081028401600061526782614a76565b9067ffffffffffffffff82169161528b836000526007602052604060002054151590565b1561553e5761535492604085936152ff6152f9946152f96152bf602060019c9b01926119156152ba36866149f0565b615e48565b91825463ffffffff8160801c16159081615520575b81615511575b816154f6575b816154e7575b50806154d8575b61544d575b36906149f0565b90616047565b60808501926153116152ba36866149f0565b8152600c6020522092835463ffffffff8160801c1615908161542f575b81615420575b81615405575b816153f6575b50806153e7575b61535a575b5036906149f0565b0161524e565b61537760a06fffffffffffffffffffffffffffffffff920161522d565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783553861534c565b506153f182614a9c565b615347565b60ff915060a01c161538615340565b6fffffffffffffffffffffffffffffffff811615915061533a565b8589015460801c159150615334565b858901546fffffffffffffffffffffffffffffffff1615915061532e565b6fffffffffffffffffffffffffffffffff615469878b0161522d565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783556152f2565b506154e281614a9c565b6152ed565b60ff915060a01c1615386152e6565b6fffffffffffffffffffffffffffffffff81161591506152e0565b848e015460801c1591506152da565b848e01546fffffffffffffffffffffffffffffffff161591506152d4565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b9060ff8091169116039060ff8211614fc457565b60ff16604d8111614fc457600a0a90565b811561559f570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f000000000000000000000000000000000000000000000000000000000000000060ff811690816006146156b3578160061161568857600661560f91615570565b90604d60ff831611801561566d575b61563657509061563061482b92615584565b90614fb1565b90507fa9cb113d00000000000000000000000000000000000000000000000000000000600052600660045260245260445260646000fd5b5061567782615584565b801561559f5760001904831161561e565b615693906006615570565b90604d60ff8316116156365750906156ad61482b92615584565b90615595565b505090565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146157945782841161577057906156fd91615570565b91604d60ff8416118015615755575b61571f5750509061563061482b92615584565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061575f83615584565b801561559f5760001904841161570c565b61577991615570565b91604d60ff84161161571f575050906156ad61482b92615584565b5050505090565b90816020910312610990575180151581036109905790565b80518015615823576020036157e557805160208281019183018390031261099057519060ff82116157e5575060ff1690565b611ec4906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484015260248301906146cf565b50507f000000000000000000000000000000000000000000000000000000000000000090565b90805115615a2f5767ffffffffffffffff8151602083012092169182600052600860205261587e816005604060002001616905565b156159eb5760005260096020526040600020815167ffffffffffffffff81116145f1576158ab8254614d16565b601f81116159b9575b506020601f821160011461592f5791615909827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361591f95600091615924575b506000198260011b9260031b1c19161790565b90556040519182916020835260208301906146cf565b0390a2565b9050840151386158f6565b601f1982169083600052806000209160005b8181106159a157509261591f9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610615988575b5050811b0190556113fb565b85015160001960f88460031b161c19169055388061597c565b9192602060018192868a015181550194019201615941565b6159e590836000526020600020601f840160051c810191602085106108bb57601f0160051c0190614ff3565b386158b4565b5090611ec46040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401526040602484015260448301906146cf565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9065ffffffffffff8091169116019065ffffffffffff8211614fc457565b67ffffffffffffffff166000818152600760205260409020549092919015615b795791615b7660e092615b4285615ace7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97615e48565b846000526008602052615ae5816040600020616047565b615aee83615e48565b846000526008602052615b08836002604060002001616047565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561099057604051907f42966c680000000000000000000000000000000000000000000000000000000082528160248160008096819560048401525af180156133eb57615c1f575050565b81615c2991614690565b50565b91908203918211614fc457565b615c41614e4a565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691615c9e6020850193615c98615c8b63ffffffff87511642615c2c565b8560808901511690614fb1565b90615cbe565b80821015615cb757505b16825263ffffffff4216905290565b9050615ca8565b91908201809211614fc457565b805160005b818110615cdc57505050565b60018101808211614fc4575b828110615cf85750600101615cd0565b6001600160a01b03615d0a8386614d02565b51166001600160a01b03615d1e8387614d02565b511614615d2d57600101615ce8565b6001600160a01b03615d3f8386614d02565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b61482b906001600160a01b03600254166001600160a01b03821614615d95575b600061679e565b7fffffffffffffffffffffffff000000000000000000000000000000000000000060025416600255615d8e565b61482b907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af61679e565b9061482b91801580615e2b575b1561679e577fffffffffffffffffffffffff00000000000000000000000000000000000000006002541660025561679e565b506001600160a01b03600254166001600160a01b03831614615df9565b805115615ee8576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff60208301511610615e855750565b606490615ee6604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590615f70575b615f0f5750565b606490615ee6604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515615f08565b80600052600060205260406000206001600160a01b03831660005260205260ff60406000205416156000146160405780600052600060205260406000206001600160a01b038316600052602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1991616180606092805461608463ffffffff8260801c1642615c2c565b90816161bf575b50506fffffffffffffffffffffffffffffffff60018160208601511692828154168085106000146161b757508280855b16167fffffffffffffffffffffffffffffffff000000000000000000000000000000008254161781556161348651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b615b7660405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916160bb565b6fffffffffffffffffffffffffffffffff916161f48392836161ed6001880154948286169560801c90614fb1565b9116615cbe565b8082101561627357505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff0000000000000000000000000000000016178155388061608b565b90506161fe565b9182549060ff8260a01c16158015616481575b61647b576fffffffffffffffffffffffffffffffff821691600185019081546162d263ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615c2c565b90816163dd575b505084811061639e57508383106163335750506163086fffffffffffffffffffffffffffffffff928392615c2c565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916163428185615c2c565b92600019810190808211614fc45761636561636a926001600160a01b0396615cbe565b615595565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611616451576163f892615c989160801c90614fb1565b8084101561644c5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806162d9565b616403565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561628d565b65ffffffffffff81116164a15765ffffffffffff1690565b7f6dfcc65000000000000000000000000000000000000000000000000000000000600052603060045260245260446000fd5b906040519182815491828252602082019060005260206000209260005b818110616505575050614b9e92500383614690565b84548352600194850194879450602090930192016164f0565b8054821015614a475760005260206000200190600090565b6000818152600460205260409020548015616040576000198101818111614fc457600354906000198201918211614fc4578181036165de575b50505060035480156165af576000190161658a81600361651e565b60001982549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6166176165ef61660093600361651e565b90549060031b1c928392600361651e565b81939154906000199060031b92831b921b19161790565b9055600052600460205260406000205538808061656f565b6000818152600760205260409020548015616040576000198101818111614fc457600654906000198201918211614fc4578181036166a8575b50505060065480156165af576000190161668381600661651e565b60001982549160031b1b19169055600655600052600760205260006040812055600190565b6166ca6166b961660093600661651e565b90549060031b1c928392600661651e565b90556000526007602052604060002055388080616668565b9060018201918160005282602052604060002054801515600014616795576000198101818111614fc4578254906000198201918211614fc45781810361675e575b505050805480156165af57600019019061673d828261651e565b60001982549160031b1b191690555560005260205260006040812055600190565b61677e61676e616600938661651e565b90549060031b1c9283928661651e565b905560005283602052604060002055388080616723565b50505050600090565b80600052600060205260406000206001600160a01b03831660005260205260ff604060002054166000146160405780600052600060205260406000206001600160a01b03831660005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b806000526004602052604060002054156000146168a557600354680100000000000000008110156145f15761688c616600826001859401600355600361651e565b9055600354906000526004602052604060002055600190565b50600090565b806000526007602052604060002054156000146168a557600654680100000000000000008110156145f1576168ec616600826001859401600655600661651e565b9055600654906000526007602052604060002055600190565b600082815260018201602052604090205461604057805490680100000000000000008210156145f1578261694361660084600180960185558461651e565b905580549260005201602052604060002055600190565b919290156169d5575081511561696e575090565b3b156169775790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156169e85750805190602001fd5b611ec4906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906146cf56fea164736f6c634300081a000aad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

var BurnMintWithLockReleaseFlagTokenPoolABI = BurnMintWithLockReleaseFlagTokenPoolMetaData.ABI

var BurnMintWithLockReleaseFlagTokenPoolBin = BurnMintWithLockReleaseFlagTokenPoolMetaData.Bin

func DeployBurnMintWithLockReleaseFlagTokenPool(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, localTokenDecimals uint8, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnMintWithLockReleaseFlagTokenPool, error) {
	parsed, err := BurnMintWithLockReleaseFlagTokenPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintWithLockReleaseFlagTokenPoolBin), backend, token, localTokenDecimals, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintWithLockReleaseFlagTokenPool{address: address, abi: *parsed, BurnMintWithLockReleaseFlagTokenPoolCaller: BurnMintWithLockReleaseFlagTokenPoolCaller{contract: contract}, BurnMintWithLockReleaseFlagTokenPoolTransactor: BurnMintWithLockReleaseFlagTokenPoolTransactor{contract: contract}, BurnMintWithLockReleaseFlagTokenPoolFilterer: BurnMintWithLockReleaseFlagTokenPoolFilterer{contract: contract}}, nil
}

type BurnMintWithLockReleaseFlagTokenPool struct {
	address common.Address
	abi     abi.ABI
	BurnMintWithLockReleaseFlagTokenPoolCaller
	BurnMintWithLockReleaseFlagTokenPoolTransactor
	BurnMintWithLockReleaseFlagTokenPoolFilterer
}

type BurnMintWithLockReleaseFlagTokenPoolCaller struct {
	contract *bind.BoundContract
}

type BurnMintWithLockReleaseFlagTokenPoolTransactor struct {
	contract *bind.BoundContract
}

type BurnMintWithLockReleaseFlagTokenPoolFilterer struct {
	contract *bind.BoundContract
}

type BurnMintWithLockReleaseFlagTokenPoolSession struct {
	Contract     *BurnMintWithLockReleaseFlagTokenPool
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintWithLockReleaseFlagTokenPoolCallerSession struct {
	Contract *BurnMintWithLockReleaseFlagTokenPoolCaller
	CallOpts bind.CallOpts
}

type BurnMintWithLockReleaseFlagTokenPoolTransactorSession struct {
	Contract     *BurnMintWithLockReleaseFlagTokenPoolTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintWithLockReleaseFlagTokenPoolRaw struct {
	Contract *BurnMintWithLockReleaseFlagTokenPool
}

type BurnMintWithLockReleaseFlagTokenPoolCallerRaw struct {
	Contract *BurnMintWithLockReleaseFlagTokenPoolCaller
}

type BurnMintWithLockReleaseFlagTokenPoolTransactorRaw struct {
	Contract *BurnMintWithLockReleaseFlagTokenPoolTransactor
}

func NewBurnMintWithLockReleaseFlagTokenPool(address common.Address, backend bind.ContractBackend) (*BurnMintWithLockReleaseFlagTokenPool, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintWithLockReleaseFlagTokenPoolABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintWithLockReleaseFlagTokenPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPool{address: address, abi: abi, BurnMintWithLockReleaseFlagTokenPoolCaller: BurnMintWithLockReleaseFlagTokenPoolCaller{contract: contract}, BurnMintWithLockReleaseFlagTokenPoolTransactor: BurnMintWithLockReleaseFlagTokenPoolTransactor{contract: contract}, BurnMintWithLockReleaseFlagTokenPoolFilterer: BurnMintWithLockReleaseFlagTokenPoolFilterer{contract: contract}}, nil
}

func NewBurnMintWithLockReleaseFlagTokenPoolCaller(address common.Address, caller bind.ContractCaller) (*BurnMintWithLockReleaseFlagTokenPoolCaller, error) {
	contract, err := bindBurnMintWithLockReleaseFlagTokenPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolCaller{contract: contract}, nil
}

func NewBurnMintWithLockReleaseFlagTokenPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintWithLockReleaseFlagTokenPoolTransactor, error) {
	contract, err := bindBurnMintWithLockReleaseFlagTokenPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolTransactor{contract: contract}, nil
}

func NewBurnMintWithLockReleaseFlagTokenPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintWithLockReleaseFlagTokenPoolFilterer, error) {
	contract, err := bindBurnMintWithLockReleaseFlagTokenPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolFilterer{contract: contract}, nil
}

func bindBurnMintWithLockReleaseFlagTokenPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintWithLockReleaseFlagTokenPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.BurnMintWithLockReleaseFlagTokenPoolCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.BurnMintWithLockReleaseFlagTokenPoolTransactor.contract.Transfer(opts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.BurnMintWithLockReleaseFlagTokenPoolTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.contract.Transfer(opts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.DEFAULTADMINROLE(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.DEFAULTADMINROLE(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "RATE_LIMITER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RATELIMITERADMINROLE(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RATELIMITERADMINROLE(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.DefaultAdmin(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) DefaultAdmin() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.DefaultAdmin(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.DefaultAdminDelay(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.DefaultAdminDelay(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.DefaultAdminDelayIncreaseWait(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getAccumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetAccumulatedFees() (*big.Int, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetAccumulatedFees(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetAccumulatedFees() (*big.Int, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetAccumulatedFees(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetAllowList(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetAllowList(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetAllowListEnabled(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetAllowListEnabled(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetCurrentInboundRateLimiterState(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRemotePools", remoteChainSelector)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRemotePools(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRemotePools(remoteChainSelector uint64) ([][]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRemotePools(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRemoteToken", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRemoteToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRemoteToken(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRemoteToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredInboundCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredInboundCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRequiredOutboundCCVs(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRmnProxy(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRmnProxy(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRoleAdmin(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, role)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRoleAdmin(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, role)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetRouter() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRouter(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetRouter(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetSupportedChains(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetSupportedChains(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetToken() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetToken() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetTokenDecimals(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetTokenDecimals() (uint8, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetTokenDecimals(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) GetTokenTransferFeeConfig(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "getTokenTransferFeeConfig", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new(IPoolV2TokenTransferFeeConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolV2TokenTransferFeeConfig)).(*IPoolV2TokenTransferFeeConfig)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) GetTokenTransferFeeConfig(arg0 common.Address, destChainSelector uint64, arg2 ClientEVM2AnyMessage, arg3 uint16, arg4 []byte) (IPoolV2TokenTransferFeeConfig, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GetTokenTransferFeeConfig(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "hasRateLimitAdminRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.HasRateLimitAdminRole(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.HasRateLimitAdminRole(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.HasRole(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.HasRole(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) IsRemotePool(opts *bind.CallOpts, remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "isRemotePool", remoteChainSelector, remotePoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) IsRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsSupportedChain(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsSupportedChain(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, remoteChainSelector)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) IsSupportedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "isSupportedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsSupportedToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, token)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) IsSupportedToken(token common.Address) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.IsSupportedToken(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, token)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) Owner() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.Owner(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) Owner() (common.Address, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.Owner(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.PendingDefaultAdmin(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.PendingDefaultAdmin(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.PendingDefaultAdminDelay(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.PendingDefaultAdminDelay(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SupportsInterface(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, interfaceId)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SupportsInterface(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts, interfaceId)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintWithLockReleaseFlagTokenPool.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) TypeAndVersion() (string, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.TypeAndVersion(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolCallerSession) TypeAndVersion() (string, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.TypeAndVersion(&_BurnMintWithLockReleaseFlagTokenPool.CallOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AcceptDefaultAdminTransfer(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AcceptDefaultAdminTransfer(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) AddRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "addRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AddRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) AddRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.AddRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyAllowListUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, removes, adds)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyAllowListUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, removes, adds)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ApplyCCVConfigUpdates(opts *bind.TransactOpts, ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "applyCCVConfigUpdates", ccvConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ApplyCCVConfigUpdates(ccvConfigArgs []TokenPoolCCVConfigArg) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyCCVConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, ccvConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ApplyChainUpdates(opts *bind.TransactOpts, remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "applyChainUpdates", remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyChainUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ApplyChainUpdates(remoteChainSelectorsToRemove []uint64, chainsToAdd []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyChainUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelectorsToRemove, chainsToAdd)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "applyFinalityConfigUpdates", finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ApplyFinalityConfigUpdates(finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyFinalityConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, finalityThreshold, customFinalityTransferFeeBps, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "applyTokenTransferFeeConfigUpdates", tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ApplyTokenTransferFeeConfigUpdates(tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ApplyTokenTransferFeeConfigUpdates(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, tokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.BeginDefaultAdminTransfer(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.BeginDefaultAdminTransfer(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newAdmin)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.CancelDefaultAdminTransfer(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.CancelDefaultAdminTransfer(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ChangeDefaultAdminDelay(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newDelay)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ChangeDefaultAdminDelay(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newDelay)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "grantRateLimitAdminRole", account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GrantRateLimitAdminRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GrantRateLimitAdminRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GrantRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.GrantRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.LockOrBurn(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.LockOrBurn(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, lockOrBurnIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) LockOrBurn0(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "lockOrBurn0", lockOrBurnIn, finality, arg2)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.LockOrBurn0(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) LockOrBurn0(lockOrBurnIn PoolLockOrBurnInV1, finality uint16, arg2 []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.LockOrBurn0(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, lockOrBurnIn, finality, arg2)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ReleaseOrMint(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ReleaseOrMint(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, releaseOrMintIn)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) ReleaseOrMint0(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "releaseOrMint0", releaseOrMintIn, finality)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ReleaseOrMint0(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) ReleaseOrMint0(releaseOrMintIn PoolReleaseOrMintInV1, finality uint16) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.ReleaseOrMint0(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, releaseOrMintIn, finality)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) RemoveRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "removeRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RemoveRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) RemoveRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RemoveRemotePool(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "renounceRole", role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RenounceRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RenounceRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "revokeRateLimitAdminRole", account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RevokeRateLimitAdminRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RevokeRateLimitAdminRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RevokeRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RevokeRole(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, role, account)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RollbackDefaultAdminDelay(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.RollbackDefaultAdminDelay(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetChainRateLimiterConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetChainRateLimiterConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setChainRateLimiterConfigs", remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetChainRateLimiterConfigs(remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetChainRateLimiterConfigs(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, remoteChainSelectors, outboundConfigs, inboundConfigs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setCustomFinalityRateLimitConfig", rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetCustomFinalityRateLimitConfig(rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetCustomFinalityRateLimitConfig(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, rateLimitConfigArgs)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetRouter(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newRouter)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.SetRouter(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, newRouter)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactor) WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.contract.Transact(opts, "withdrawFees", recipient)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.WithdrawFees(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, recipient)
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolTransactorSession) WithdrawFees(recipient common.Address) (*types.Transaction, error) {
	return _BurnMintWithLockReleaseFlagTokenPool.Contract.WithdrawFees(&_BurnMintWithLockReleaseFlagTokenPool.TransactOpts, recipient)
}

type BurnMintWithLockReleaseFlagTokenPoolAllowListAddIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolAllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolAllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolAllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolAllowListAddIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolAllowListAddIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolAllowListAdd)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseAllowListAdd(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolAllowListAdd, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolAllowListAdd)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolAllowListRemoveIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolAllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolAllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolAllowListRemoveIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolAllowListRemoveIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolAllowListRemove)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseAllowListRemove(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolAllowListRemove, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolAllowListRemove)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdatedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated struct {
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdatedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdatedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "CCVConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "CCVConfigUpdated", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseCCVConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CCVConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainAdded struct {
	RemoteChainSelector       uint64
	RemoteToken               []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolChainAdded)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseChainAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainAdded, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolChainAdded)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainConfiguredIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainConfiguredIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolChainConfiguredIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolChainConfigured)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseChainConfigured(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainConfigured, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolChainConfigured)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolChainRemoved)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseChainRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainRemoved, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolChainRemoved)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolConfigChangedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolConfigChangedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolConfigChangedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolConfigChanged)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseConfigChanged(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolConfigChanged, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolConfigChanged)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "CustomFinalityOutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "CustomFinalityOutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CustomFinalityOutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "CustomFinalityTransferInboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "CustomFinalityTransferInboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "CustomFinalityTransferInboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceledIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceledIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduledIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduledIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceledIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceledIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduledIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduledIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdatedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated struct {
	FinalityConfig               uint16
	CustomFinalityTransferFeeBps uint16
	Raw                          types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdatedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdatedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "FinalityConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "FinalityConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseFinalityConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "FinalityConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "InboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "InboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "InboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "LockedOrBurned", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "LockedOrBurned", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseLockedOrBurned(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "LockedOrBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed struct {
	RemoteChainSelector uint64
	Token               common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "OutboundRateLimitConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "OutboundRateLimitConsumed", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "OutboundRateLimitConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawnIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawnIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawnIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "PoolFeeWithdrawn", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "PoolFeeWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParsePoolFeeWithdrawn(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "PoolFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGrantedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGrantedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGrantedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RateLimitAdminRoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRateLimitAdminRoleGranted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevokedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked struct {
	Account common.Address
	Raw     types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevokedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevokedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RateLimitAdminRoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRateLimitAdminRoleRevoked(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted struct {
	RemoteChainSelector uint64
	Token               common.Address
	Sender              common.Address
	Recipient           common.Address
	Amount              *big.Int
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "ReleasedOrMinted", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "ReleasedOrMinted", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseReleasedOrMinted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "ReleasedOrMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RemotePoolAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RemotePoolAdded", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRemotePoolAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RemotePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RemotePoolRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RemotePoolRemoved", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRemotePoolRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RemotePoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRoleAdminChangedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintWithLockReleaseFlagTokenPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRoleAdminChangedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRoleAdminChanged(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRoleGrantedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRoleGrantedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRoleGranted)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRoleGranted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRoleGranted, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRoleGranted)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRoleRevokedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRoleRevokedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRoleRevoked)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRoleRevoked(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRoleRevoked, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRoleRevoked)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolRouterUpdatedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolRouterUpdatedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolRouterUpdated)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRouterUpdated, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolRouterUpdated)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted struct {
	DestChainSelector uint64
	Raw               types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "TokenTransferFeeConfigDeleted", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigDeleted", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator struct {
	Event *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated struct {
	DestChainSelector      uint64
	TokenTransferFeeConfig IPoolV2TokenTransferFeeConfig
	Raw                    types.Log
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.FilterLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator{contract: _BurnMintWithLockReleaseFlagTokenPool.contract, event: "TokenTransferFeeConfigUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}

	logs, sub, err := _BurnMintWithLockReleaseFlagTokenPool.contract.WatchLogs(opts, "TokenTransferFeeConfigUpdated", destChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated)
				if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPoolFilterer) ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, error) {
	event := new(BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated)
	if err := _BurnMintWithLockReleaseFlagTokenPool.contract.UnpackLog(event, "TokenTransferFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (BurnMintWithLockReleaseFlagTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnMintWithLockReleaseFlagTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c")
}

func (BurnMintWithLockReleaseFlagTokenPoolChainAdded) Topic() common.Hash {
	return common.HexToHash("0x8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c2")
}

func (BurnMintWithLockReleaseFlagTokenPoolChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (BurnMintWithLockReleaseFlagTokenPoolChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnMintWithLockReleaseFlagTokenPoolConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b2")
}

func (BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff7")
}

func (BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba")
}

func (BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0x50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c")
}

func (BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned) Topic() common.Hash {
	return common.HexToHash("0xf33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae10")
}

func (BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed) Topic() common.Hash {
	return common.HexToHash("0xff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da8178944")
}

func (BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted) Topic() common.Hash {
	return common.HexToHash("0xf7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e389")
}

func (BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b")
}

func (BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted) Topic() common.Hash {
	return common.HexToHash("0xfc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0")
}

func (BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded) Topic() common.Hash {
	return common.HexToHash("0x7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea")
}

func (BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved) Topic() common.Hash {
	return common.HexToHash("0x52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76")
}

func (BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintWithLockReleaseFlagTokenPoolRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintWithLockReleaseFlagTokenPoolRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintWithLockReleaseFlagTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted) Topic() common.Hash {
	return common.HexToHash("0x5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee8")
}

func (BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70")
}

func (_BurnMintWithLockReleaseFlagTokenPool *BurnMintWithLockReleaseFlagTokenPool) Address() common.Address {
	return _BurnMintWithLockReleaseFlagTokenPool.address
}

type BurnMintWithLockReleaseFlagTokenPoolInterface interface {
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

	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)

	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

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

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	WithdrawFees(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolAllowListRemove, error)

	FilterCCVConfigUpdated(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdatedIterator, error)

	WatchCCVConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCCVConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCCVConfigUpdated, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolConfigChanged, error)

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolDefaultAdminTransferScheduled, error)

	FilterFinalityConfigUpdated(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdatedIterator, error)

	WatchFinalityConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated) (event.Subscription, error)

	ParseFinalityConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolFinalityConfigUpdated, error)

	FilterInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumedIterator, error)

	WatchInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseInboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolInboundRateLimitConsumed, error)

	FilterLockedOrBurned(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurnedIterator, error)

	WatchLockedOrBurned(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, remoteChainSelector []uint64) (event.Subscription, error)

	ParseLockedOrBurned(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolLockedOrBurned, error)

	FilterOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumedIterator, error)

	WatchOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseOutboundRateLimitConsumed(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolOutboundRateLimitConsumed, error)

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGrantedIterator, error)

	WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error)

	ParseRateLimitAdminRoleGranted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleGranted, error)

	FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevokedIterator, error)

	WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error)

	ParseRateLimitAdminRoleRevoked(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRateLimitAdminRoleRevoked, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRemotePoolRemoved, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintWithLockReleaseFlagTokenPoolRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintWithLockReleaseFlagTokenPoolRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRoleRevoked, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintWithLockReleaseFlagTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolRouterUpdated, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*BurnMintWithLockReleaseFlagTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
