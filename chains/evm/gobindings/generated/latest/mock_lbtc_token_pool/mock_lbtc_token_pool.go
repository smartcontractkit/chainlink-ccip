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

var MockE2ELBTCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIBurnMintERC20\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMITER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_destPoolData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenDataMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]}]",
	Bin: "0x610100806040523461062657617d15803803809161001d8285610643565b833981019060a0818303126106265780516001600160a01b03811691908290036106265760208101516001600160401b0381116106265781019183601f84011215610626578251926001600160401b03841161041d578360051b60208101946100896040519687610643565b85526020808601918301019186831161062657602001905b82821061062b575050506100b760408301610666565b6100c360608401610666565b608084015190936001600160401b038211610626570185601f82011215610626578051906001600160401b03821161041d576040519661010d601f8401601f191660200189610643565b828852602083830101116106265760005b82811061061157505060206000918701015233156105fb57600180546001600160d01b031690556002546001600160a01b0381166105ea576001600160a01b03191633908117600255610170906106a4565b50811580156105d9575b80156105c8575b6105b7578160209160049360805260c0526040519283809263313ce56760e01b82525afa8091600091610574575b5090610550575b50600860a052600580546001600160a01b0319166001600160a01b03929092169190911790558051151560e0819052610433575b5080516001600160401b03811161041d57600f54600181811c91168015610413575b60208210146103fd57601f8111610398575b50602091601f821160011461033457918192600092610329575b50508160011b916000199060031b1c191617600f555b60405161745290816108a382396080518181816118cb01528181611b6801528181611d1b015281816124360152818161268d015281816128490152818161346e015281816136d9015281816137a20152818161392b01528181613af901528181614588015281816145e20152818161473f01526155b4015260a051818181611bf1015281816144ff01528181615e9f0152615f22015260c051818181610cb001528181611966015281816124d10152818161350901526139c7015260e051818181610c5e015281816119a901528181612514015261320c0152f35b015190503880610238565b601f19821692600f600052806000209160005b85811061038057508360019510610367575b505050811b01600f5561024e565b015160001960f88460031b161c19169055388080610359565b91926020600181928685015181550194019201610347565b600f6000527f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac802601f830160051c810191602084106103f3575b601f0160051c01905b8181106103e7575061021e565b600081556001016103da565b90915081906103d1565b634e487b7160e01b600052602260045260246000fd5b90607f169061020c565b634e487b7160e01b600052604160045260246000fd5b60206040516104428282610643565b60008152600036813760e0511561053f5760005b81518110156104bd576001906001600160a01b03610474828561067a565b5116846104808261074a565b61048d575b505001610456565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13884610485565b505060005b8251811015610536576001906001600160a01b036104e0828661067a565b5116801561053057836104f282610848565b610500575b50505b016104c2565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138836104f7565b506104fa565b505050386101ea565b6335f4a7b360e01b60005260046000fd5b60ff16600881146101b6576332ad3e0760e11b600052600860045260245260446000fd5b6020813d6020116105af575b8161058d60209383610643565b810103126105ab57519060ff821682036105a85750386101af565b80fd5b5080fd5b3d9150610580565b630a64406560e11b60005260046000fd5b506001600160a01b03811615610181565b506001600160a01b0383161561017a565b631fe1e13d60e11b60005260046000fd5b636116401160e11b600052600060045260246000fd5b80602080928401015182828b0101520161011e565b600080fd5b6020809161063884610666565b8152019101906100a1565b601f909101601f19168101906001600160401b0382119082101761041d57604052565b51906001600160a01b038216820361062657565b805182101561068e5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b0381166000908152600080516020617cf5833981519152602052604090205460ff1661072c576001600160a01b03166000818152600080516020617cf583398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b805482101561068e5760005260206000200190600090565b600081815260046020526040902054801561084157600019810181811161082b5760035460001981019190821161082b578181036107da575b50505060035480156107c4576000190161079e816003610732565b8154906000199060031b1b19169055600355600052600460205260006040812055600190565b634e487b7160e01b600052603160045260246000fd5b6108136107eb6107fc936003610732565b90549060031b1c9283926003610732565b819391549060031b91821b91600019901b19161790565b90556000526004602052604060002055388080610783565b634e487b7160e01b600052601160045260246000fd5b5050600090565b8060005260046020526040600020541560001461072c576003546801000000000000000081101561041d576108896107fc8260018594016003556003610732565b905560035490600052600460205260406000205560019056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714614a2957508063022d63fb14614a0b5780630aa6220b146149425780630bd7c46d146148c7578063164e68de14614667578063181f5a771461460657806321df0da7146145b5578063240028e814614551578063248a9ca31461452357806324f65ee7146144e55780632a10097b1461426c5780632c286daf146141165780632f2ff15d146140cd57806336568abe14613f5e57806337b1924714613e5757806339077537146138b95780633e5db5d11461389d578063489a68f2146133cd5780634c5ef0ed146133865780634f71592c1461335157806354c8a4f3146131a95780635df45a371461318657806362ddd3c4146130d5578063634e93da14612f9b578063649a5ec714612d7a578063791e5a1014612d3f578063804ba5a914612ccc57806384ef8ffc14612c525780638926f54f14612c865780638da5cb5b14612c5257806391d1485414612bf9578063962d402014612a275780639a4575b9146123d55780639f68f6731461239d578063a1eda53c1461233a578063a217fddf1461231e578063a42a7b8b146121b7578063a7cd63b714612136578063acfecf9114611fbe578063af58d59f14611f75578063b0f479a114611f41578063b1c71c6514611840578063b79019b514611555578063b79465801461151c578063c0d7865514611414578063c4bffe2b146112e9578063c75eea9c14611241578063cc8463c814611216578063cefc1429146110fa578063cf6eefb71461108d578063cf7401f314610e90578063d547741f14610e16578063d602b9fd14610d9a578063da90a9f314610cd4578063dc0bd97114610c83578063e0351e1314610c46578063e58d80c714610bcf5763e8a1da171461029f57600080fd5b34610bcc576102ad366151c6565b92909391828052826020526040832073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615610ba45782918593915b808310610a0f575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b85821015610a0b578160051b85013581811215610a075785019061012082360312610a07576040519561035887614cdd565b823567ffffffffffffffff81168103610a02578752602083013567ffffffffffffffff81116109fe5783019536601f880112156109fe5786359661039b886154ff565b976103a9604051998a614d15565b8089526020808a019160051b830101903682116109fa5760208301905b8282106109c7575050505060208801968752604084013567ffffffffffffffff81116109c3576103f990369086016150e0565b9860408901998a5261042361041136606088016152da565b9560608b0196875260c03691016152da565b9660808a01978852610435865161665b565b61043f885161665b565b8a51511561099b5761045b67ffffffffffffffff8b51166172ca565b156109645767ffffffffffffffff8a5116815260086020526040812061059b87516fffffffffffffffffffffffffffffffff604082015116906105566fffffffffffffffffffffffffffffffff602083015116915115158360806040516104c181614cdd565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b6106c189516fffffffffffffffffffffffffffffffff6040820151169061067c6fffffffffffffffffffffffffffffffff602083015116915115158360806040516105e581614cdd565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff8211610937576106e48354614ebf565b601f81116108fc575b50602090601f831160011461085d5761073b9291859183610852575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b90555b805b895180518210156107765790610770600192610769838f67ffffffffffffffff9051169261564c565b519061602c565b01610740565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c29295939661084467ffffffffffffffff600197949c51169251935191516108106107db60405196879687526101006020880152610100870190614db3565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019093949291610326565b015190508f80610709565b83855281852091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08416865b8181106108e457509084600195949392106108ad575b505050811b01905561073e565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558e80806108a0565b9293602060018192878601518155019501930161088a565b6109279084865260208620601f850160051c8101916020861061092d575b601f0160051c0190615855565b8e6106ed565b909150819061091a565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff81116109f6576020916109eb83928336918901016150e0565b8152019101906103c6565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff610a2f610a2a8785886153a4565b615360565b1695610a3a87616f37565b15610b78578684526008602052610a5660056040862001616d45565b94845b8651811015610a8f576001908987526008602052610a8860056040892001610a81838b61564c565b5190617062565b5001610a59565b5093945094909580855260086020526005604086208681558660018201558660028201558660038201558660048201610ac88154614ebf565b80610b37575b5050500180549086815581610b19575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946102ed565b865260208620908101905b81811015610ade57868155600101610b24565b601f8111600114610b4d5750555b868a80610ace565b81835260208320610b6891601f01861c810190600101615855565b8082528160208120915555610b45565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6004837f2b5c74de000000000000000000000000000000000000000000000000000000008152fd5b80fd5b5034610bcc576020600319360112610bcc5773ffffffffffffffffffffffffffffffffffffffff6040610c00614bf3565b927f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af815280602052209116600052602052602060ff604060002054166040519015158152f35b5034610bcc5780600319360112610bcc5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b5034610bcc5780600319360112610bcc57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610bcc576020600319360112610bcc57610cee614bf3565b818052816020526040822073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615610d7257602081610d517fd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b936165bb565b5073ffffffffffffffffffffffffffffffffffffffff60405191168152a180f35b6004827f2b5c74de000000000000000000000000000000000000000000000000000000008152fd5b5034610bcc5780600319360112610bcc57610db36158c7565b600180547fffffffffffff0000000000000000000000000000000000000000000000000000811690915560a01c65ffffffffffff16610def5780f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051098180a180f35b5034610bcc576040600319360112610bcc57600435610e33614c16565b908015610e68579081610e5f610e5a610e6494600052600060205260016040600020015490565b615933565b6165e5565b5080f35b6004837f3fc3c27a000000000000000000000000000000000000000000000000000000008152fd5b5034610bcc5760e0600319360112610bcc57610eaa614e7a565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc360112610bcc57604051610ee181614cf9565b60243580151581036110895781526044356fffffffffffffffffffffffffffffffff811681036110895760208201526064356fffffffffffffffffffffffffffffffff8116810361108957604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c3601126110855760405190610f6882614cf9565b6084358015158103610a0757825260a4356fffffffffffffffffffffffffffffffff81168103610a0757602083015260c4356fffffffffffffffffffffffffffffffff81168103610a075760408301527f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8352826020526040832073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff60406000205416158061104d575b6110215761101e92936162b4565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b50828052826020526040832073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615611010565b5080fd5b8280fd5b5034610bcc5780600319360112610bcc57604065ffffffffffff6110d46001549065ffffffffffff73ffffffffffffffffffffffffffffffffffffffff83169260a01c1690565b73ffffffffffffffffffffffffffffffffffffffff849392935193168352166020820152f35b5034610bcc5780600319360112610bcc5760015473ffffffffffffffffffffffffffffffffffffffff1633036111ea5760015473ffffffffffffffffffffffffffffffffffffffff81169060a01c65ffffffffffff16801580156111e0575b6111b557506111899061118373ffffffffffffffffffffffffffffffffffffffff6002541661654d565b506159ca565b507fffffffffffff00000000000000000000000000000000000000000000000000006001541660015580f35b7f19ca5ebb000000000000000000000000000000000000000000000000000000008352600452602482fd5b5042811015611159565b807fc22c8022000000000000000000000000000000000000000000000000000000006024925233600452fd5b5034610bcc5780600319360112610bcc57602061123161588e565b65ffffffffffff60405191168152f35b5034610bcc576020600319360112610bcc5761128c61128760406112e59367ffffffffffffffff611270614e7a565b61127861569f565b501681526008602052206156ca565b6163f1565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b5034610bcc5780600319360112610bcc57604051906006548083528260208101600684526020842092845b8181106113fb57505061132992500383614d15565b815161134d611337826154ff565b916113456040519384614d15565b8083526154ff565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0602083019301368437805b84518110156113ac578067ffffffffffffffff6113996001938861564c565b51166113a5828661564c565b520161137a565b50925090604051928392602084019060208552518091526040840192915b8181106113d8575050500390f35b825167ffffffffffffffff168452859450602093840193909201916001016113ca565b8454835260019485019487945060209093019201611314565b5034610bcc576020600319360112610bcc5761142e614bf3565b818052816020526040822073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615610d725773ffffffffffffffffffffffffffffffffffffffff1680156114f45760407f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f16849160055490807fffffffffffffffffffffffff000000000000000000000000000000000000000083161760055573ffffffffffffffffffffffffffffffffffffffff8351921682526020820152a180f35b6004827f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b5034610bcc576020600319360112610bcc576112e561154161153c614e7a565b61586c565b604051918291602083526020830190614db3565b5034610bcc576020600319360112610bcc5760043567ffffffffffffffff811161108557611587903690600401614df6565b828052826020526040832073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615610ba457825b8181106115ca578380f35b6115d8610a2a828486615730565b6115f06115e6838587615730565b6020810190615770565b907fb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c61162a61162086888a615730565b6040810190615770565b91909261164061163b368784615517565b616483565b61164e61163b368587615517565b6040519461165b86614cc1565b611666368284615517565b86526116b067ffffffffffffffff61167f368789615517565b9860208901998a52169586956116a26040519586956040875260408701916157c4565b9184830360208601526157c4565b0390a28652600d60205260408620905180519067ffffffffffffffff8211611813576801000000000000000082116118135760209083548385558084106117f9575b500182885260208820885b8381106117cf575050505060010190519081519167ffffffffffffffff83116117a2576801000000000000000083116117a2576020908254848455808510611788575b500190865260208620865b83811061175e57505050506001016115bf565b600190602073ffffffffffffffffffffffffffffffffffffffff855116940193818401550161174b565b83895282892061179c918101908601615855565b38611740565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b600190602073ffffffffffffffffffffffffffffffffffffffff85511694019381840155016116fd565b848a52828a2061180d918101908501615855565b386116f2565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b5034610bcc576060600319360112610bcc5760043567ffffffffffffffff81116110855760a060031982360301126110855761187a614e27565b9060443567ffffffffffffffff8111610a075761189b9036906004016150e0565b506118a4615633565b5060848101916118b383615440565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603611ef757602482019177ffffffffffffffff0000000000000000000000000000000061191984615360565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115611e0e578691611ec8575b50611ea0576119a760448201615440565b7f0000000000000000000000000000000000000000000000000000000000000000611e44575b5067ffffffffffffffff6119e084615360565b166119f8816000526007602052604060002054151590565b15611e1957602073ffffffffffffffffffffffffffffffffffffffff60055416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa8015611e0e578690611dab575b73ffffffffffffffffffffffffffffffffffffffff9150163303611d7f57606461ffff91013591169384151593848095611d70575b15611cb95761ffff600a5416808710611c895750611c4d9550611ae0611ad0611ab686615360565b67ffffffffffffffff16600052600b602052604060002090565b84611ada84615440565b91616ab4565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff611b1c611b1687615360565b93615440565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018790529190931692a25b508092611c57575b5061153c81611b62611be993615360565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae109080606081015b0390a2615360565b9060405160ff7f000000000000000000000000000000000000000000000000000000000000000016602082015260208152611c25604082614d15565b60405192611c3284614cc1565b83526020830152604051928392604084526040840190615286565b9060208301520390f35b611be9919250611c8161153c91612710611c7a61ffff600a5460101c1683615813565b04906163e4565b929150611b51565b82604491887fe08f03ef000000000000000000000000000000000000000000000000000000008352600452602452fd5b50611c4d945067ffffffffffffffff611cd184615360565b1680825260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448380611d436040862073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616ab4565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2611b49565b5061ffff600a54161515611a8e565b6024857f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011611e06575b81611dc560209383614d15565b810103126109fe575173ffffffffffffffffffffffffffffffffffffffff811681036109fe5773ffffffffffffffffffffffffffffffffffffffff90611a59565b3d9150611db8565b6040513d88823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008652600452602485fd5b73ffffffffffffffffffffffffffffffffffffffff16611e71816000526004602052604060002054151590565b6119cd577fd0d25976000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b611eea915060203d602011611ef0575b611ee28183614d15565b810190615e13565b38611996565b503d611ed8565b60248473ffffffffffffffffffffffffffffffffffffffff611f1886615440565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b5034610bcc5780600319360112610bcc57602073ffffffffffffffffffffffffffffffffffffffff60055416604051908152f35b5034610bcc576020600319360112610bcc5761128c611287600260406112e59467ffffffffffffffff611fa6614e7a565b611fae61569f565b50168152600860205220016156ca565b5034610bcc57611fcd36615214565b91838052836020526040842073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff604060002054161561210e5767ffffffffffffffff1691612026836000526007602052604060002054151590565b156120e2578284526008602052612055600560408620016120483684866150a9565b6020815191012090617062565b1561209a57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d7691612094604051928392602084526020840191615660565b0390a280f35b826120de836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191615660565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b6004847f2b5c74de000000000000000000000000000000000000000000000000000000008152fd5b5034610bcc5780600319360112610bcc5760405160038054808352908352909160208301917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b915b8181106121a1576112e58561219581870382614d15565b60405191829182615176565b825484526020909301926001928301920161217e565b5034610bcc576020600319360112610bcc5767ffffffffffffffff6121da614e7a565b16815260086020526121f160056040832001616d45565b80517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612236612220836154ff565b9261222e6040519485614d15565b8084526154ff565b01835b81811061230d575050825b825181101561228a578061225a6001928561564c565b518552600960205261226e60408620614fe9565b612278828561564c565b52612283818461564c565b5001612244565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106122c257505050500390f35b919360206122fd827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc060019597998495030186528851614db3565b96019201920185949391926122b3565b806060602080938601015201612239565b5034610bcc5780600319360112610bcc57602090604051908152f35b5034610bcc5780600319360112610bcc576002548060d01c9182151580612393575b1561238a575060a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b9150508061236e565b504283101561235c565b5034610bcc57612195600160406112e59367ffffffffffffffff6123c0366150fe565b505050509050168152600d602052200161549e565b5034610bcc576020600319360112610bcc576004359067ffffffffffffffff8211610bcc5760a06003198336030112610bcc57612410615633565b506084820161241e81615440565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603612a0657602483019277ffffffffffffffff0000000000000000000000000000000061248485615360565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561292d5784916129e7575b506129bf5761251260448201615440565b7f0000000000000000000000000000000000000000000000000000000000000000612963575b5067ffffffffffffffff61254b85615360565b16612563816000526007602052604060002054151590565b1561293857602073ffffffffffffffffffffffffffffffffffffffff60055416916024604051809481937fa8d87a3b00000000000000000000000000000000000000000000000000000000835260048301525afa801561292d5784906128ca575b73ffffffffffffffffffffffffffffffffffffffff915016330361289e57606401359082156127ea5761ffff600a5416806127ba5750612613612609611ab686615360565b83611ada84615440565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff612649611b1687615360565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018690529190931692a25b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691823b15610bcc576040517f42966c68000000000000000000000000000000000000000000000000000000008152826004820152818160248183885af180156127af5761279a575b6112e561276a61153c87877ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae108867ffffffffffffffff61273485615360565b6040805173ffffffffffffffffffffffffffffffffffffffff909616865233602087015285019290925216918060608101611be1565b6040519061277782614cc1565b8152612781614f12565b6020820152604051918291602083526020830190615286565b6127a5828092614d15565b610bcc57806126f5565b6040513d84823e3d90fd5b836044917fe08f03ef00000000000000000000000000000000000000000000000000000000825281600452602452fd5b5067ffffffffffffffff6127fd84615360565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448280612871604060002073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616ab4565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a2612676565b6024837f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b506020813d602011612925575b816128e460209383614d15565b81010312610a07575173ffffffffffffffffffffffffffffffffffffffff81168103610a075773ffffffffffffffffffffffffffffffffffffffff906125c4565b3d91506128d7565b6040513d86823e3d90fd5b7fa9902c7e000000000000000000000000000000000000000000000000000000008452600452602483fd5b73ffffffffffffffffffffffffffffffffffffffff16612990816000526004602052604060002054151590565b612538577fd0d25976000000000000000000000000000000000000000000000000000000008452600452602483fd5b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612a00915060203d602011611ef057611ee28183614d15565b38612501565b9073ffffffffffffffffffffffffffffffffffffffff611f18602493615440565b5034610bcc576060600319360112610bcc5760043567ffffffffffffffff811161108557612a59903690600401614df6565b60243567ffffffffffffffff8111610a0757612a79903690600401615255565b60449291923567ffffffffffffffff81116109fe57612a9c903690600401615255565b9190927f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8752866020526040872073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff604060002054161580612bc1575b612b9557818114801590612b8b575b612b6357865b818110612b17578780f35b80612b5d612b2b610a2a600194868c6153a4565b612b3683878b615623565b612b57612b4f612b47868b8d615623565b9236906152da565b9136906152da565b916162b4565b01612b0c565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b5082811415612b06565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b50868052866020526040872073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615612af7565b5034610bcc576040600319360112610bcc5773ffffffffffffffffffffffffffffffffffffffff6040612c2a614c16565b92600435815280602052209116600052602052602060ff604060002054166040519015158152f35b5034610bcc5780600319360112610bcc57602073ffffffffffffffffffffffffffffffffffffffff60025416604051908152f35b5034610bcc576020600319360112610bcc576020612cc267ffffffffffffffff612cae614e7a565b166000526007602052604060002054151590565b6040519015158152f35b5034610bcc576020600319360112610bcc5760043567ffffffffffffffff811161108557612cfe903690600401614e49565b90828052826020526040832073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615610ba4579061101e91615aed565b5034610bcc5780600319360112610bcc5760206040517f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8152f35b5034610bcc576020600319360112610bcc5760043565ffffffffffff811680820361108957612da76158c7565b612db042616cfb565b9065ffffffffffff612dc061588e565b1680821115612f3357507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9291612e0c9162069780811015612f225765ffffffffffff905b1690616296565b906002548060d01c80612e9f575b50506002805473ffffffffffffffffffffffffffffffffffffffff1660a083901b79ffffffffffff0000000000000000000000000000000000000000161760d084901b7fffffffffffff0000000000000000000000000000000000000000000000000000161790556040805165ffffffffffff9283168152919092166020820152a180f35b421115612ef85779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b3880612e1a565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58480a1612ef1565b5065ffffffffffff62069780612e05565b0365ffffffffffff8111612f6e577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b9291612e0c9190616296565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b5034610bcc576020600319360112610bcc57612fb5614bf3565b612fbd6158c7565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed66020612ffa612fec42616cfb565b612ff461588e565b90616296565b65ffffffffffff73ffffffffffffffffffffffffffffffffffffffff6130436001549065ffffffffffff73ffffffffffffffffffffffffffffffffffffffff83169260a01c1690565b9690501694600154867fffffffffffff000000000000000000000000000000000000000000000000000079ffffffffffff00000000000000000000000000000000000000008660a01b1692161717600155166130ac575b65ffffffffffff60405191168152a280f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051098580a161309a565b5034610bcc576130e436615214565b838093929452826020526040832073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff6040600020541615610ba45767ffffffffffffffff821661313f816000526007602052604060002054151590565b1561315b575061101e92936131559136916150a9565b9061602c565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b5034610bcc5780600319360112610bcc5760206131a161556b565b604051908152f35b5034610bcc576131b8366151c6565b93929091838052836020526040842073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff604060002054161561210e5761320a9291613202913691615517565b933691615517565b7f00000000000000000000000000000000000000000000000000000000000000001561332957815b83518110156132a5578073ffffffffffffffffffffffffffffffffffffffff61325d6001938761564c565b511661326881616da8565b613274575b5001613232565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13861326d565b5090805b8251811015610e64578073ffffffffffffffffffffffffffffffffffffffff6132d46001938661564c565b51168015613323576132e58161726a565b6132f2575b505b016132a9565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a1846132ea565b506132ec565b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b5034610bcc5761219560406112e59267ffffffffffffffff613372366150fe565b505050509050168152600d6020522061549e565b5034610bcc576040600319360112610bcc576133a0614e7a565b906024359067ffffffffffffffff8211610bcc576020612cc2846133c736600487016150e0565b90615461565b5034610bcc576040600319360112610bcc576004359067ffffffffffffffff8211610bcc578160040161010060031984360301126110855761340d614e27565b8260405161341a81614ca5565b5261344761343d61343861343160c48801866153b4565b36916150a9565b615e2b565b6064860135615f1f565b92608485019161345683615440565b73ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001691160361387c57602486019377ffffffffffffffff000000000000000000000000000000006134bc86615360565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa90811561292d57849161385d575b506129bf5767ffffffffffffffff61355086615360565b16613568816000526007602052604060002054151590565b1561293857602073ffffffffffffffffffffffffffffffffffffffff60055416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa90811561292d57849161383e575b501561289e576135df85615360565b906135f560a48901926133c761343185856153b4565b156137f7575050611b1660446020977ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09567ffffffffffffffff95898961ffff6136d3981615156000146137485760409150918861365561366794615360565b168152600c8d52208a611ada84615440565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff786613695611b168b615360565b6040805173ffffffffffffffffffffffffffffffffffffffff929092168252602082018d90529190931692a25b01946136cd86615440565b50615360565b604080517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff9081168252336020830152909216908201526060810185905292169180608081015b0390a28060405161373f81614ca5565b52604051908152f35b7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9293506137ca600260408b61377e8695615360565b1696878152602060089052200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016928391616ab4565b6040805173ffffffffffffffffffffffffffffffffffffffff9290921682526020820192909252a26136c2565b61380192506153b4565b6120de6040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191615660565b613857915060203d602011611ef057611ee28183614d15565b386135d0565b613876915060203d602011611ef057611ee28183614d15565b38613539565b5073ffffffffffffffffffffffffffffffffffffffff611f18602493615440565b5034610bcc5780600319360112610bcc576112e5611541614f12565b5034610bcc576020600319360112610bcc576004359067ffffffffffffffff8211610bcc578160040191610100600319823603011261108557816040516138ff81614ca5565b526064810135906084810161391381615440565b73ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016911603613e365750602481019077ffffffffffffffff0000000000000000000000000000000061397a83615360565b60801b16604051907f2cbc26bb000000000000000000000000000000000000000000000000000000008252600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115613d64578591613e17575b50613def5767ffffffffffffffff613a0e83615360565b16613a26816000526007602052604060002054151590565b15613dc457602073ffffffffffffffffffffffffffffffffffffffff60055416916044604051809481937f83826b2b00000000000000000000000000000000000000000000000000000000835260048301523360248301525afa908115613d64578591613da5575b5015613d7957613a9d82615360565b613ab260a48301916133c7613431848a6153b4565b15613d6f5750839483945067ffffffffffffffff613acf84615360565b168087526008602052613b216002604089200173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016968791616ab4565b6040805173ffffffffffffffffffffffffffffffffffffffff87168152602081018890527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a26020613b76600f54614ebf565b14613c80575b5060440190613b8a82615440565b85843b15610bcc576040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff929092166004830152602482018690528160448183885af18015611e0e5785613c32611b1667ffffffffffffffff9560209a7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc0989661372f96613c70575b5050615360565b6040805173ffffffffffffffffffffffffffffffffffffffff9889168152336020820152979091169087015260608601529116929081906080820190565b81613c7a91614d15565b38613c2b565b85613c8e60e48401836153b4565b8101919060408184031261108557803567ffffffffffffffff81116110895783613cb99183016150e0565b60208201359167ffffffffffffffff8311610a0757602094613cf393613cdf92016150e0565b508360405192828480945193849201614d90565b8101039060025afa15613d645785519060c4830190613d1b613d1583836153b4565b90615405565b8303613d28575050613b7c565b91613d39613d1589936044956153b4565b7f7f249311000000000000000000000000000000000000000000000000000000008352600452602452fd5b6040513d87823e3d90fd5b61380190866153b4565b6024847f728fe07b00000000000000000000000000000000000000000000000000000000815233600452fd5b613dbe915060203d602011611ef057611ee28183614d15565b38613a8e565b7fa9902c7e000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b613e30915060203d602011611ef057611ee28183614d15565b386139f7565b8373ffffffffffffffffffffffffffffffffffffffff611f18602493615440565b5034610bcc5760a0600319360112610bcc57613e71614bf3565b5060243567ffffffffffffffff81168091036110855760443567ffffffffffffffff81116110895760031960a0913603011261108557613eaf614e38565b5060843567ffffffffffffffff81116110895791604091613ed66080943690600401614e91565b50508160608451613ee681614c5a565b828152826020820152828682015201528152600e6020522060405190613f0b82614c5a565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b5034610bcc576040600319360112610bcc57600435613f7b614c16565b90801580614096575b613fd8575b3373ffffffffffffffffffffffffffffffffffffffff831603613fb05790610e64916165e5565b6004837f6697b232000000000000000000000000000000000000000000000000000000008152fd5b60015465ffffffffffff60a082901c169073ffffffffffffffffffffffffffffffffffffffff1615801590614086575b8015614074575b61404157507fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff60015416600155613f89565b7f19ca5ebb00000000000000000000000000000000000000000000000000000000845265ffffffffffff16600452602483fd5b504265ffffffffffff8216101561400f565b5065ffffffffffff811615614008565b5073ffffffffffffffffffffffffffffffffffffffff6002541673ffffffffffffffffffffffffffffffffffffffff831614613f84565b5034610bcc576040600319360112610bcc576004356140ea614c16565b908015610e68579081614111610e5a610e6494600052600060205260016040600020015490565b615a5c565b5034610bcc576060600319360112610bcc5760043561ffff81168091036110855761413f614e27565b60443567ffffffffffffffff8111610a075761415f903690600401614e49565b90848052846020526040852073ffffffffffffffffffffffffffffffffffffffff3316865260205260ff604086205416156142445761ffff8316926127108410156142185784926040949261420a927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba977fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600a549360101b1692161717600a55615aed565b82519182526020820152a180f35b602486857f95f3517a000000000000000000000000000000000000000000000000000000008252600452fd5b6004857f2b5c74de000000000000000000000000000000000000000000000000000000008152fd5b5034610bcc576040600319360112610bcc5760043567ffffffffffffffff8111611085573660238201121561108557806004013567ffffffffffffffff81116110895760248201916024369160a084020101116110895760243567ffffffffffffffff8111610a07576142e3903690600401614df6565b919092848052846020526040852073ffffffffffffffffffffffffffffffffffffffff3316865260205260ff604086205416156142445784805b83811061438757509150505b818110614334578380f35b8067ffffffffffffffff61434e610a2a60019486886153a4565b16808652600e6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201614329565b6001917f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d706080856144d561449b63ffffffff6144ca61448e826144bf6143de8f806143d88f9283610a2a918e615321565b9a615321565b604067ffffffffffffffff602083019a169c8d8152600e60205220836144038b615375565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff0000000061444087615375565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e615375565b60401b1696019e8f615386565b151560601b16951617161717179055826144b76040519a615393565b168952615393565b166020870152615393565b1660408401526152b0565b15156060820152a201859061431d565b5034610bcc5780600319360112610bcc57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610bcc576020600319360112610bcc5760206131a1600435600052600060205260016040600020015490565b5034610bcc576020600319360112610bcc5760209061456e614bf3565b905073ffffffffffffffffffffffffffffffffffffffff807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b5034610bcc5780600319360112610bcc57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b5034610bcc5780600319360112610bcc57506112e5604051614629604082614d15565b601a81527f4d6f636b4532454c425443546f6b656e506f6f6c20312e352e310000000000006020820152604051918291602083526020830190614db3565b5034610bcc576020600319360112610bcc57614681614bf3565b818052816020526040822073ffffffffffffffffffffffffffffffffffffffff3316835260205260ff60408320541615610d72576146bd61556b565b90816146c7578280f35b6040517fa9059cbb000000000000000000000000000000000000000000000000000000006020820190815273ffffffffffffffffffffffffffffffffffffffff8316602483015260448083018590528252906147cd90614728606482614d15565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001686806040958651946147718887614d15565b602086527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646020870152519082855af13d156148bf573d916147b283614d56565b926147bf87519485614d15565b83523d89602085013e617379565b80518061481e575b505073ffffffffffffffffffffffffffffffffffffffff7f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992602092519485521692a238808280f35b90602080614830938301019101615e13565b1561483c5738806147d5565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b606091617379565b5034610bcc576020600319360112610bcc576148e1614bf3565b818052816020526040822073ffffffffffffffffffffffffffffffffffffffff3316835260205260ff60408320541615610d7257602081610d517ff7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e389936159a0565b5034610bcc5780600319360112610bcc5761495b6158c7565b6002548060d01c80614988575b8273ffffffffffffffffffffffffffffffffffffffff6002541660025580f35b4211156149e15779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b3880614968565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58180a16149da565b5034610bcc5780600319360112610bcc576020604051620697808152f35b905034611085576020600319360112611085576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361108957602092507ff208a58f000000000000000000000000000000000000000000000000000000008114908115614bc9575b8115614b9f575b8115614b75575b8115614b4b575b8115614abb575b5015158152f35b7f3149878600000000000000000000000000000000000000000000000000000000811491508115614aee575b5038614ab4565b7f7965db0b00000000000000000000000000000000000000000000000000000000811491508115614b21575b5038614ae7565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438614b1a565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150614aad565b7f0e64dd290000000000000000000000000000000000000000000000000000000081149150614aa6565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150614a9f565b7faff2afbf0000000000000000000000000000000000000000000000000000000081149150614a98565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610a0257565b6024359073ffffffffffffffffffffffffffffffffffffffff82168203610a0257565b359073ffffffffffffffffffffffffffffffffffffffff82168203610a0257565b6080810190811067ffffffffffffffff821117614c7657604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff821117614c7657604052565b6040810190811067ffffffffffffffff821117614c7657604052565b60a0810190811067ffffffffffffffff821117614c7657604052565b6060810190811067ffffffffffffffff821117614c7657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117614c7657604052565b67ffffffffffffffff8111614c7657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60005b838110614da35750506000910152565b8181015183820152602001614d93565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602093614def81518092818752878088019101614d90565b0116010190565b9181601f84011215610a025782359167ffffffffffffffff8311610a02576020808501948460051b010111610a0257565b6024359061ffff82168203610a0257565b6064359061ffff82168203610a0257565b9181601f84011215610a025782359167ffffffffffffffff8311610a025760208085019460e08502010111610a0257565b6004359067ffffffffffffffff82168203610a0257565b9181601f84011215610a025782359167ffffffffffffffff8311610a025760208381860195010111610a0257565b90600182811c92168015614f08575b6020831014614ed957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691614ece565b60405190600082600f5491614f2683614ebf565b8083529260018116908115614fac5750600114614f4c575b614f4a92500383614d15565b565b50600f600090815290917f8d1108e10bcb7c27dddfc02ed9d693a074039d026cf4ea4240b40f7d581ac8025b818310614f90575050906020614f4a92820101614f3e565b6020919350806001915483858901015201910190918492614f78565b60209250614f4a9491507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001682840152151560051b820101614f3e565b9060405191826000825492614ffd84614ebf565b80845293600181169081156150695750600114615022575b50614f4a92500383614d15565b90506000929192526020600020906000915b81831061504d575050906020614f4a9282010138615015565b6020919350806001915483858901015201910190918492615034565b60209350614f4a9592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138615015565b9291926150b582614d56565b916150c36040519384614d15565b829481845281830111610a02578281602093846000960137010152565b9080601f83011215610a02578160206150fb933591016150a9565b90565b60a0600319820112610a025760043573ffffffffffffffffffffffffffffffffffffffff81168103610a02579160243567ffffffffffffffff81168103610a0257916044359160643561ffff81168103610a0257916084359067ffffffffffffffff8211610a025761517291600401614e91565b9091565b602060408183019282815284518094520192019060005b81811061519a5750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161518d565b6040600319820112610a025760043567ffffffffffffffff8111610a0257816151f191600401614df6565b929092916024359067ffffffffffffffff8211610a025761517291600401614df6565b906040600319830112610a025760043567ffffffffffffffff81168103610a0257916024359067ffffffffffffffff8211610a025761517291600401614e91565b9181601f84011215610a025782359167ffffffffffffffff8311610a025760208085019460608502010111610a0257565b6150fb91602061529f8351604084526040840190614db3565b920151906020818403910152614db3565b35908115158203610a0257565b35906fffffffffffffffffffffffffffffffff82168203610a0257565b9190826060910312610a02576040516152f281614cf9565b604061531c818395615303816152b0565b8552615311602082016152bd565b6020860152016152bd565b910152565b91908110156153315760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff81168103610a025790565b3563ffffffff81168103610a025790565b358015158103610a025790565b359063ffffffff82168203610a0257565b91908110156153315760051b0190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610a02570180359067ffffffffffffffff8211610a0257602001918136038313610a0257565b359060208110615413575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b3573ffffffffffffffffffffffffffffffffffffffff81168103610a025790565b9067ffffffffffffffff6150fb92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b8181106154d0575050614f4a92500383614d15565b845473ffffffffffffffffffffffffffffffffffffffff168352600194850194879450602090930192016154bb565b67ffffffffffffffff8111614c765760051b60200190565b929190615523816154ff565b936155316040519586614d15565b602085838152019160051b8101928311610a0257905b82821061555357505050565b6020809161556084614c39565b815201910190615547565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260208160248173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa908115615617576000916155e8575090565b90506020813d60201161560f575b8161560360209383614d15565b81010312610a02575190565b3d91506155f6565b6040513d6000823e3d90fd5b9190811015615331576060020190565b6040519061564082614cc1565b60606020838281520152565b80518210156153315760209160051b010190565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b604051906156ac82614cdd565b60006080838281528260208201528260408201528260608201520152565b906040516156d781614cdd565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b91908110156153315760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa181360301821215610a02570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610a02570180359067ffffffffffffffff8211610a0257602001918160051b36038313610a0257565b9160209082815201919060005b8181106157de5750505090565b90919260208060019273ffffffffffffffffffffffffffffffffffffffff61580588614c39565b1681520194019291016157d1565b8181029291811591840414171561582657565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818110615860575050565b60008155600101615855565b67ffffffffffffffff1660005260086020526150fb6004604060002001614fe9565b6002548060d01c80151590816158bd575b50156158b35760a01c65ffffffffffff1690565b5060015460d01c90565b905042113861589f565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff161561590057565b7fe2517d3f0000000000000000000000000000000000000000000000000000000060005233600452600060245260446000fd5b806000526000602052604060002073ffffffffffffffffffffffffffffffffffffffff331660005260205260ff604060002054161561596f5750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b6150fb907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af6167a2565b6002549073ffffffffffffffffffffffffffffffffffffffff8216615a32576150fb917fffffffffffffffffffffffff000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff831691161760025560006167a2565b7f3fc3c27a0000000000000000000000000000000000000000000000000000000060005260046000fd5b908115615a6d575b6150fb916167a2565b6002549173ffffffffffffffffffffffffffffffffffffffff8316615a32577fffffffffffffffffffffffff000000000000000000000000000000000000000090921673ffffffffffffffffffffffffffffffffffffffff821617600255615a64565b356fffffffffffffffffffffffffffffffff81168103610a025790565b9160005b82811015615e0d5760e0810284016000615b0a82615360565b9067ffffffffffffffff821691615b2e836000526007602052604060002054151590565b15615de157615bf79260408593615ba2615b9c94615b9c615b62602060019c9b0192611ab6615b5d36866152da565b61665b565b91825463ffffffff8160801c16159081615dc3575b81615db4575b81615d99575b81615d8a575b5080615d7b575b615cf0575b36906152da565b90616881565b6080850192615bb4615b5d36866152da565b8152600c6020522092835463ffffffff8160801c16159081615cd2575b81615cc3575b81615ca8575b81615c99575b5080615c8a575b615bfd575b5036906152da565b01615af1565b615c1a60a06fffffffffffffffffffffffffffffffff9201615ad0565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538615bef565b50615c9482615386565b615bea565b60ff915060a01c161538615be3565b6fffffffffffffffffffffffffffffffff8116159150615bdd565b8589015460801c159150615bd7565b858901546fffffffffffffffffffffffffffffffff16159150615bd1565b6fffffffffffffffffffffffffffffffff615d0c878b01615ad0565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff1617178355615b95565b50615d8581615386565b615b90565b60ff915060a01c161538615b89565b6fffffffffffffffffffffffffffffffff8116159150615b83565b848e015460801c159150615b7d565b848e01546fffffffffffffffffffffffffffffffff16159150615b77565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b90816020910312610a0257518015158103610a025790565b80518015615e9b57602003615e5d578051602082810191830183900312610a0257519060ff8211615e5d575060ff1690565b6120de906040519182917f953576f7000000000000000000000000000000000000000000000000000000008352602060048401526024830190614db3565b50507f000000000000000000000000000000000000000000000000000000000000000090565b9060ff8091169116039060ff821161582657565b60ff16604d811161582657600a0a90565b8115615ef0570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff81169282841461602557828411615ffb5790615f6491615ec1565b91604d60ff8416118015615fc2575b615f8c57505090615f866150fb92615ed5565b90615813565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b50615fcc83615ed5565b8015615ef0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048411615f73565b61600491615ec1565b91604d60ff841611615f8c5750509061601f6150fb92615ed5565b90615ee6565b5050505090565b9080511561626c5767ffffffffffffffff81516020830120921691826000526008602052616061816005604060002001617324565b156162285760005260096020526040600020815167ffffffffffffffff8111614c765761608e8254614ebf565b601f81116161f6575b506020601f8211600114616130579161610a827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361612095600091616125575b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c19161790565b9055604051918291602083526020830190614db3565b0390a2565b9050840151386160d9565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082169083600052806000209160005b8181106161de5750926161209492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea9896106161a7575b5050811b019055611541565b8501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055388061619b565b9192602060018192868a015181550194019201616160565b61622290836000526020600020601f840160051c8101916020851061092d57601f0160051c0190615855565b38616097565b50906120de6040519283927f393b8ad20000000000000000000000000000000000000000000000000000000084526004840152604060248401526044830190614db3565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b9065ffffffffffff8091169116019065ffffffffffff821161582657565b67ffffffffffffffff1660008181526007602052604090205490929190156163b657916163b360e09261637f8561630b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b9761665b565b846000526008602052616322816040600020616881565b61632b8361665b565b846000526008602052616345836002604060002001616881565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9190820391821161582657565b6163f961569f565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691616456602085019361645061644363ffffffff875116426163e4565b8560808901511690615813565b90616476565b8082101561646f57505b16825263ffffffff4216905290565b9050616460565b9190820180921161582657565b805160005b81811061649457505050565b60018101808211615826575b8281106164b05750600101616488565b73ffffffffffffffffffffffffffffffffffffffff6164cf838661564c565b511673ffffffffffffffffffffffffffffffffffffffff6164f0838761564c565b5116146164ff576001016164a0565b73ffffffffffffffffffffffffffffffffffffffff61651e838661564c565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6150fb9073ffffffffffffffffffffffffffffffffffffffff6002541673ffffffffffffffffffffffffffffffffffffffff82161461658e575b6000617196565b7fffffffffffffffffffffffff000000000000000000000000000000000000000060025416600255616587565b6150fb907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af617196565b906150fb91801580616624575b15617196577fffffffffffffffffffffffff000000000000000000000000000000000000000060025416600255617196565b5073ffffffffffffffffffffffffffffffffffffffff6002541673ffffffffffffffffffffffffffffffffffffffff8316146165f2565b8051156166fb576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff602083015116106166985750565b6064906166f9604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590616783575b6167225750565b6064906166f9604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff602082015116151561671b565b806000526000602052604060002073ffffffffffffffffffffffffffffffffffffffff831660005260205260ff604060002054161560001461687a57806000526000602052604060002073ffffffffffffffffffffffffffffffffffffffff8316600052602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0082541617905573ffffffffffffffffffffffffffffffffffffffff339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916169ba60609280546168be63ffffffff8260801c16426163e4565b90816169f9575b50506fffffffffffffffffffffffffffffffff60018160208601511692828154168085106000146169f157508280855b16167fffffffffffffffffffffffffffffffff0000000000000000000000000000000082541617815561696e8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b6163b360405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b8380916168f5565b6fffffffffffffffffffffffffffffffff91616a2e839283616a276001880154948286169560801c90615813565b9116616476565b80821015616aad57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff000000000000000000000000000000001617815538806168c5565b9050616a38565b9182549060ff8260a01c16158015616cf3575b616ced576fffffffffffffffffffffffffffffffff82169160018501908154616b0c63ffffffff6fffffffffffffffffffffffffffffffff83169360801c16426163e4565b9081616c4f575b5050848110616c035750838310616b6d575050616b426fffffffffffffffffffffffffffffffff9283926163e4565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c91616b7c81856163e4565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019080821161582657616bca616bcf9273ffffffffffffffffffffffffffffffffffffffff96616476565b615ee6565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828573ffffffffffffffffffffffffffffffffffffffff927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611616cc357616c6a926164509160801c90615813565b80841015616cbe5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff0000000000000000000000000000000016178655923880616b13565b616c75565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b508215616ac7565b65ffffffffffff8111616d135765ffffffffffff1690565b7f6dfcc65000000000000000000000000000000000000000000000000000000000600052603060045260245260446000fd5b906040519182815491828252602082019060005260206000209260005b818110616d77575050614f4a92500383614d15565b8454835260019485019487945060209093019201616d62565b80548210156153315760005260206000200190600090565b600081815260046020526040902054801561687a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161582657600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161582657818103616ec8575b5050506003548015616e99577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01616e56816003616d90565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b616f1f616ed9616eea936003616d90565b90549060031b1c9283926003616d90565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526004602052604060002055388080616e1d565b600081815260076020526040902054801561687a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161582657600654907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161582657818103617028575b5050506006548015616e99577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01616fe5816006616d90565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600655600052600760205260006040812055600190565b61704a617039616eea936006616d90565b90549060031b1c9283926006616d90565b90556000526007602052604060002055388080616fac565b906001820191816000528260205260406000205480151560001461718d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111615826578254907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820191821161582657818103617156575b50505080548015616e99577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906171178282616d90565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b191690555560005260205260006040812055600190565b617176617166616eea9386616d90565b90549060031b1c92839286616d90565b9055600052836020526040600020553880806170df565b50505050600090565b806000526000602052604060002073ffffffffffffffffffffffffffffffffffffffff831660005260205260ff6040600020541660001461687a57806000526000602052604060002073ffffffffffffffffffffffffffffffffffffffff831660005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815416905573ffffffffffffffffffffffffffffffffffffffff339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b806000526004602052604060002054156000146172c45760035468010000000000000000811015614c76576172ab616eea8260018594016003556003616d90565b9055600354906000526004602052604060002055600190565b50600090565b806000526007602052604060002054156000146172c45760065468010000000000000000811015614c765761730b616eea8260018594016006556006616d90565b9055600654906000526007602052604060002055600190565b600082815260018201602052604090205461687a5780549068010000000000000000821015614c765782617362616eea846001809601855584616d90565b905580549260005201602052604060002055600190565b919290156173f4575081511561738d575090565b3b156173965790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b8251909150156174075750805190602001fd5b6120de906040519182917f08c379a0000000000000000000000000000000000000000000000000000000008352602060048401526024830190614db356fea164736f6c634300081a000aad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.DEFAULTADMINROLE(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.DEFAULTADMINROLE(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "RATE_LIMITER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.RATELIMITERADMINROLE(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) RATELIMITERADMINROLE() ([32]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.RATELIMITERADMINROLE(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) DefaultAdmin() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.DefaultAdmin(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) DefaultAdmin() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.DefaultAdmin(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) DefaultAdminDelay() (*big.Int, error) {
	return _MockE2ELBTCTokenPool.Contract.DefaultAdminDelay(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _MockE2ELBTCTokenPool.Contract.DefaultAdminDelay(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _MockE2ELBTCTokenPool.Contract.DefaultAdminDelayIncreaseWait(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _MockE2ELBTCTokenPool.Contract.DefaultAdminDelayIncreaseWait(&_MockE2ELBTCTokenPool.CallOpts)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRequiredInboundCCVs(&_MockE2ELBTCTokenPool.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRequiredInboundCCVs(&_MockE2ELBTCTokenPool.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRequiredOutboundCCVs(&_MockE2ELBTCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRequiredOutboundCCVs(&_MockE2ELBTCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRoleAdmin(&_MockE2ELBTCTokenPool.CallOpts, role)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRoleAdmin(&_MockE2ELBTCTokenPool.CallOpts, role)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GetRouter() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRouter(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _MockE2ELBTCTokenPool.Contract.GetRouter(&_MockE2ELBTCTokenPool.CallOpts)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "hasRateLimitAdminRole", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.HasRateLimitAdminRole(&_MockE2ELBTCTokenPool.CallOpts, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) HasRateLimitAdminRole(account common.Address) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.HasRateLimitAdminRole(&_MockE2ELBTCTokenPool.CallOpts, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.HasRole(&_MockE2ELBTCTokenPool.CallOpts, role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MockE2ELBTCTokenPool.Contract.HasRole(&_MockE2ELBTCTokenPool.CallOpts, role, account)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _MockE2ELBTCTokenPool.Contract.PendingDefaultAdmin(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _MockE2ELBTCTokenPool.Contract.PendingDefaultAdmin(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _MockE2ELBTCTokenPool.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _MockE2ELBTCTokenPool.Contract.PendingDefaultAdminDelay(&_MockE2ELBTCTokenPool.CallOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _MockE2ELBTCTokenPool.Contract.PendingDefaultAdminDelay(&_MockE2ELBTCTokenPool.CallOpts)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AcceptDefaultAdminTransfer(&_MockE2ELBTCTokenPool.TransactOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.AcceptDefaultAdminTransfer(&_MockE2ELBTCTokenPool.TransactOpts)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.BeginDefaultAdminTransfer(&_MockE2ELBTCTokenPool.TransactOpts, newAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.BeginDefaultAdminTransfer(&_MockE2ELBTCTokenPool.TransactOpts, newAdmin)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.CancelDefaultAdminTransfer(&_MockE2ELBTCTokenPool.TransactOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.CancelDefaultAdminTransfer(&_MockE2ELBTCTokenPool.TransactOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ChangeDefaultAdminDelay(&_MockE2ELBTCTokenPool.TransactOpts, newDelay)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.ChangeDefaultAdminDelay(&_MockE2ELBTCTokenPool.TransactOpts, newDelay)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) GrantRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "grantRateLimitAdminRole", account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.GrantRateLimitAdminRole(&_MockE2ELBTCTokenPool.TransactOpts, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) GrantRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.GrantRateLimitAdminRole(&_MockE2ELBTCTokenPool.TransactOpts, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "grantRole", role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.GrantRole(&_MockE2ELBTCTokenPool.TransactOpts, role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.GrantRole(&_MockE2ELBTCTokenPool.TransactOpts, role, account)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "renounceRole", role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RenounceRole(&_MockE2ELBTCTokenPool.TransactOpts, role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RenounceRole(&_MockE2ELBTCTokenPool.TransactOpts, role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "revokeRateLimitAdminRole", account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RevokeRateLimitAdminRole(&_MockE2ELBTCTokenPool.TransactOpts, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) RevokeRateLimitAdminRole(account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RevokeRateLimitAdminRole(&_MockE2ELBTCTokenPool.TransactOpts, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "revokeRole", role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RevokeRole(&_MockE2ELBTCTokenPool.TransactOpts, role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RevokeRole(&_MockE2ELBTCTokenPool.TransactOpts, role, account)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RollbackDefaultAdminDelay(&_MockE2ELBTCTokenPool.TransactOpts)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.RollbackDefaultAdminDelay(&_MockE2ELBTCTokenPool.TransactOpts)
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

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRouter(&_MockE2ELBTCTokenPool.TransactOpts, newRouter)
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _MockE2ELBTCTokenPool.Contract.SetRouter(&_MockE2ELBTCTokenPool.TransactOpts, newRouter)
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
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
	Raw                 types.Log
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

type MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceledIterator struct {
	Event *MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceledIterator{contract: _MockE2ELBTCTokenPool.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled, error) {
	event := new(MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduledIterator struct {
	Event *MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduledIterator{contract: _MockE2ELBTCTokenPool.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled, error) {
	event := new(MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolDefaultAdminTransferCanceledIterator struct {
	Event *MockE2ELBTCTokenPoolDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolDefaultAdminTransferCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolDefaultAdminTransferCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolDefaultAdminTransferCanceledIterator{contract: _MockE2ELBTCTokenPool.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolDefaultAdminTransferCanceled)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*MockE2ELBTCTokenPoolDefaultAdminTransferCanceled, error) {
	event := new(MockE2ELBTCTokenPoolDefaultAdminTransferCanceled)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolDefaultAdminTransferScheduledIterator struct {
	Event *MockE2ELBTCTokenPoolDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolDefaultAdminTransferScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolDefaultAdminTransferScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*MockE2ELBTCTokenPoolDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolDefaultAdminTransferScheduledIterator{contract: _MockE2ELBTCTokenPool.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolDefaultAdminTransferScheduled)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*MockE2ELBTCTokenPoolDefaultAdminTransferScheduled, error) {
	event := new(MockE2ELBTCTokenPoolDefaultAdminTransferScheduled)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
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

type MockE2ELBTCTokenPoolRateLimitAdminRoleGrantedIterator struct {
	Event *MockE2ELBTCTokenPoolRateLimitAdminRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRateLimitAdminRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRateLimitAdminRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRateLimitAdminRoleGranted struct {
	Account common.Address
	Raw     types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRateLimitAdminRoleGrantedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRateLimitAdminRoleGrantedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RateLimitAdminRoleGranted", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRateLimitAdminRoleGranted)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRateLimitAdminRoleGranted(log types.Log) (*MockE2ELBTCTokenPoolRateLimitAdminRoleGranted, error) {
	event := new(MockE2ELBTCTokenPoolRateLimitAdminRoleGranted)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRateLimitAdminRoleRevokedIterator struct {
	Event *MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRateLimitAdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked struct {
	Account common.Address
	Raw     types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRateLimitAdminRoleRevokedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRateLimitAdminRoleRevokedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RateLimitAdminRoleRevoked", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RateLimitAdminRoleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRateLimitAdminRoleRevoked(log types.Log) (*MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked, error) {
	event := new(MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RateLimitAdminRoleRevoked", log); err != nil {
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

type MockE2ELBTCTokenPoolRoleAdminChangedIterator struct {
	Event *MockE2ELBTCTokenPoolRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MockE2ELBTCTokenPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRoleAdminChangedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRoleAdminChanged)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRoleAdminChanged(log types.Log) (*MockE2ELBTCTokenPoolRoleAdminChanged, error) {
	event := new(MockE2ELBTCTokenPoolRoleAdminChanged)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRoleGrantedIterator struct {
	Event *MockE2ELBTCTokenPoolRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MockE2ELBTCTokenPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRoleGrantedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRoleGranted)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRoleGranted(log types.Log) (*MockE2ELBTCTokenPoolRoleGranted, error) {
	event := new(MockE2ELBTCTokenPoolRoleGranted)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRoleRevokedIterator struct {
	Event *MockE2ELBTCTokenPoolRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MockE2ELBTCTokenPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRoleRevokedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRoleRevoked)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRoleRevoked(log types.Log) (*MockE2ELBTCTokenPoolRoleRevoked, error) {
	event := new(MockE2ELBTCTokenPoolRoleRevoked)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockE2ELBTCTokenPoolRouterUpdatedIterator struct {
	Event *MockE2ELBTCTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockE2ELBTCTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockE2ELBTCTokenPoolRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockE2ELBTCTokenPoolRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockE2ELBTCTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *MockE2ELBTCTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockE2ELBTCTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &MockE2ELBTCTokenPoolRouterUpdatedIterator{contract: _MockE2ELBTCTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _MockE2ELBTCTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockE2ELBTCTokenPoolRouterUpdated)
				if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockE2ELBTCTokenPool *MockE2ELBTCTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*MockE2ELBTCTokenPoolRouterUpdated, error) {
	event := new(MockE2ELBTCTokenPoolRouterUpdated)
	if err := _MockE2ELBTCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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

type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (MockE2ELBTCTokenPoolAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (MockE2ELBTCTokenPoolAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (MockE2ELBTCTokenPoolCCVConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c")
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

func (MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (MockE2ELBTCTokenPoolDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (MockE2ELBTCTokenPoolDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
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

func (MockE2ELBTCTokenPoolPoolFeeWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b859599")
}

func (MockE2ELBTCTokenPoolRateLimitAdminRoleGranted) Topic() common.Hash {
	return common.HexToHash("0xf7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e389")
}

func (MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b")
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

func (MockE2ELBTCTokenPoolRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (MockE2ELBTCTokenPoolRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (MockE2ELBTCTokenPoolRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (MockE2ELBTCTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
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

	SDestPoolData(opts *bind.CallOpts) ([]byte, error)

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

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*MockE2ELBTCTokenPoolDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*MockE2ELBTCTokenPoolDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*MockE2ELBTCTokenPoolDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*MockE2ELBTCTokenPoolDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*MockE2ELBTCTokenPoolDefaultAdminTransferScheduled, error)

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

	FilterPoolFeeWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*MockE2ELBTCTokenPoolPoolFeeWithdrawnIterator, error)

	WatchPoolFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolPoolFeeWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParsePoolFeeWithdrawn(log types.Log) (*MockE2ELBTCTokenPoolPoolFeeWithdrawn, error)

	FilterRateLimitAdminRoleGranted(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRateLimitAdminRoleGrantedIterator, error)

	WatchRateLimitAdminRoleGranted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitAdminRoleGranted) (event.Subscription, error)

	ParseRateLimitAdminRoleGranted(log types.Log) (*MockE2ELBTCTokenPoolRateLimitAdminRoleGranted, error)

	FilterRateLimitAdminRoleRevoked(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRateLimitAdminRoleRevokedIterator, error)

	WatchRateLimitAdminRoleRevoked(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked) (event.Subscription, error)

	ParseRateLimitAdminRoleRevoked(log types.Log) (*MockE2ELBTCTokenPoolRateLimitAdminRoleRevoked, error)

	FilterReleasedOrMinted(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolReleasedOrMintedIterator, error)

	WatchReleasedOrMinted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolReleasedOrMinted, remoteChainSelector []uint64) (event.Subscription, error)

	ParseReleasedOrMinted(log types.Log) (*MockE2ELBTCTokenPoolReleasedOrMinted, error)

	FilterRemotePoolAdded(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolAddedIterator, error)

	WatchRemotePoolAdded(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolAdded, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolAdded(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolAdded, error)

	FilterRemotePoolRemoved(opts *bind.FilterOpts, remoteChainSelector []uint64) (*MockE2ELBTCTokenPoolRemotePoolRemovedIterator, error)

	WatchRemotePoolRemoved(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRemotePoolRemoved, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolRemoved(log types.Log) (*MockE2ELBTCTokenPoolRemotePoolRemoved, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MockE2ELBTCTokenPoolRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*MockE2ELBTCTokenPoolRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MockE2ELBTCTokenPoolRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*MockE2ELBTCTokenPoolRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MockE2ELBTCTokenPoolRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*MockE2ELBTCTokenPoolRoleRevoked, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*MockE2ELBTCTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*MockE2ELBTCTokenPoolRouterUpdated, error)

	FilterTokenTransferFeeConfigDeleted(opts *bind.FilterOpts, destChainSelector []uint64) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigDeletedIterator, error)

	WatchTokenTransferFeeConfigDeleted(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigDeleted(log types.Log) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigDeleted, error)

	FilterTokenTransferFeeConfigUpdated(opts *bind.FilterOpts, destChainSelector []uint64) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdatedIterator, error)

	WatchTokenTransferFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated, destChainSelector []uint64) (event.Subscription, error)

	ParseTokenTransferFeeConfigUpdated(log types.Log) (*MockE2ELBTCTokenPoolTokenTransferFeeConfigUpdated, error)

	Address() common.Address
}
