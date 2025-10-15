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

var SiloedUSDCTokenPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"localTokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"allowlist\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockBox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RATE_LIMITER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAllowListUpdates\",\"inputs\":[{\"name\":\"removes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"adds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyAuthorizedCallerUpdates\",\"inputs\":[{\"name\":\"authorizedCallerArgs\",\"type\":\"tuple\",\"internalType\":\"structAuthorizedCallers.AuthorizedCallerArgs\",\"components\":[{\"name\":\"addedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"removedCallers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyCCVConfigUpdates\",\"inputs\":[{\"name\":\"ccvConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CCVConfigArg[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyChainUpdates\",\"inputs\":[{\"name\":\"remoteChainSelectorsToRemove\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"chainsToAdd\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.ChainUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddresses\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"remoteTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyFinalityConfigUpdates\",\"inputs\":[{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"applyTokenTransferFeeConfigUpdates\",\"inputs\":[{\"name\":\"tokenTransferFeeConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.TokenTransferFeeConfigArgs[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}]},{\"name\":\"destToUseDefaultFeeConfigs\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnLockedUSDC\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelExistingCCTPMigrationProposal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"excludeTokensFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccumulatedFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllAuthorizedCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllowListEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAvailableTokens\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"lockedTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentProposedCCTPChainMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExcludedTokensByChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRebalancer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemotePools\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemoteToken\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredInboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequiredOutboundCCVs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requiredCCVs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRmnProxy\",\"inputs\":[],\"outputs\":[{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedChains\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDecimals\",\"inputs\":[],\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenTransferFeeConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structClient.EVM2AnyMessage\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"feeConfig\",\"type\":\"tuple\",\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnsiloedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"out\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockOrBurn\",\"inputs\":[{\"name\":\"lockOrBurnIn\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnInV1\",\"components\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"originalSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.LockOrBurnOutV1\",\"components\":[{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destPoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"destTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeCCTPMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provideSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"releaseOrMint\",\"inputs\":[{\"name\":\"releaseOrMintIn\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintInV1\",\"components\":[{\"name\":\"originalSender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceDenominatedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"localToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainTokenData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"finality\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"components\":[{\"name\":\"destinationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeRemotePool\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRateLimitAdminRole\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfig\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChainRateLimiterConfigs\",\"inputs\":[{\"name\":\"remoteChainSelectors\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"outboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundConfigs\",\"type\":\"tuple[]\",\"internalType\":\"structRateLimiter.Config[]\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCircleMigratorAddress\",\"inputs\":[{\"name\":\"migrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomFinalityRateLimitConfig\",\"inputs\":[{\"name\":\"rateLimitConfigArgs\",\"type\":\"tuple[]\",\"internalType\":\"structTokenPool.CustomFinalityRateLimitConfigArgs[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRebalancer\",\"inputs\":[{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRouter\",\"inputs\":[{\"name\":\"newRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSiloRebalancer\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSiloDesignations\",\"inputs\":[{\"name\":\"removes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"adds\",\"type\":\"tuple[]\",\"internalType\":\"structSiloedLockReleaseTokenPool.SiloConfigUpdate[]\",\"components\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawLiquidity\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSiloedLiquidity\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowListAdd\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowListRemove\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerAdded\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuthorizedCallerRemoved\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationCancelled\",\"inputs\":[{\"name\":\"existingProposalSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationExecuted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"USDCBurned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCTPMigrationProposed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCVConfigUpdated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"outboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"inboundCCVs\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remoteToken\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainConfigured\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChainUnsiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"amountUnsiloed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CircleMigratorAddressSet\",\"inputs\":[{\"name\":\"migratorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigChanged\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityOutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CustomFinalityTransferInboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FinalityConfigUpdated\",\"inputs\":[{\"name\":\"finalityConfig\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"customFinalityTransferFeeBps\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidityRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"remover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedOrBurned\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OutboundRateLimitConsumed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolFeeWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleGranted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RateLimitAdminRoleRevoked\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReleasedOrMinted\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemotePoolRemoved\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouterUpdated\",\"inputs\":[{\"name\":\"oldRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRouter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SiloRebalancerSet\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigDeleted\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFeeConfigUpdated\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"tokenTransferFeeConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPoolV2.TokenTransferFeeConfig\",\"components\":[{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeUSDCents\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensExcludedFromBurn\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnableAmountAfterExclusion\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnsiloedRebalancerSet\",\"inputs\":[{\"name\":\"oldRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRebalancer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AllowListNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BucketOverfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotARampOnRouter\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChainAlreadyExists\",\"inputs\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainAlreadyMigrated\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotAllowed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"ChainNotSiloed\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"CursedByRMN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisabledNonZeroRateLimit\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"DuplicateCCVNotAllowed\",\"inputs\":[{\"name\":\"ccvAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ExistingMigrationProposal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLiquidity\",\"inputs\":[{\"name\":\"availableLiquidity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidChainSelector\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidDecimalArgs\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"actual\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidDestBytesOverhead\",\"inputs\":[{\"name\":\"destBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"InvalidFinality\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"finalityThreshold\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidFinalityConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRateLimitRate\",\"inputs\":[{\"name\":\"rateLimiterConfig\",\"type\":\"tuple\",\"internalType\":\"structRateLimiter.Config\",\"components\":[{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}]},{\"type\":\"error\",\"name\":\"InvalidRemoteChainDecimals\",\"inputs\":[{\"name\":\"sourcePoolData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidRemotePoolForChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSourcePoolAddress\",\"inputs\":[{\"name\":\"sourcePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidTransferFeeBps\",\"inputs\":[{\"name\":\"bps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LiquidityAmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMigrationProposalPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonExistentChain\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCircle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OverflowDetected\",\"inputs\":[{\"name\":\"remoteDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"localDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"remoteAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PoolAlreadyAdded\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"remotePoolAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SenderNotAllowed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenLockingNotAllowedAfterMigration\",\"inputs\":[{\"name\":\"remoteChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"TokenMaxCapacityExceeded\",\"inputs\":[{\"name\":\"capacity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenRateLimitReached\",\"inputs\":[{\"name\":\"minWaitInSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokenTransferFeeConfigNotEnabled\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x610120806040523461056e576185cb803803809161001d8285610804565b8339810160c08282031261056e578151906001600160a01b0382169081830361056e5761004c60208501610827565b60408501519091906001600160401b03811161056e5785019080601f8301121561056e578151916001600160401b0383116104de578260051b9060208201936100986040519586610804565b845260208085019282010192831161056e57602001905b8282106107ec575050506100c560608601610835565b946100de60a06100d760808401610835565b9201610835565b92602096604051966100f08989610804565b60008852600036813733156107db57600180546001600160a01b03191633179055861580156107ca575b80156107b9575b6106245760805260c05260405163313ce56760e01b81528781600481895afa60009181610782575b50610757575b5060a052600580546001600160a01b0319166001600160a01b039290921691909117905561017c33610873565b50805115158060e052610635575b506001600160a01b0316801561062457604051636eb1769f60e11b8152306004820152602481018290528481604481865afa908115610618576000916105eb575b5061058057604051918483019263095ea7b360e01b84528260248201526000196044820152604481526101ff606482610804565b6000806040958651936102128886610804565b8985527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648a860152519082865af13d15610573573d906001600160401b0382116104de57855161027f949092610271601f8201601f19168b0185610804565b83523d60008a85013e610b51565b848151806104f4575b50505061010052805161029b8482610804565b6000815260003681378151928383016001600160401b038111858210176104de5783528352808484015260005b815181101561032e576001906001600160a01b036102e68285610849565b5116866102f282610958565b6102ff575b5050016102c8565b7fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda77580918651908152a138866102f7565b505090519060005b82518110156103a6576001600160a01b036103518285610849565b5116908115610395577feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef8583610388600195610a84565b508551908152a101610336565b6342bcdf7f60e11b60005260046000fd5b50516179899081610c2282396080518181816104ed0152818161146f01528181611bf601528181611de901528181611f42015281816124ad0152818161264f0152818161289101528181612ddb01528181613c7901528181613e3e01528181613eed015281816140750152818161419b0152818161474f01528181614d0401528181614d5101528181614e8b01528181614fdc01526159d2015260a051818181614c16015281816165cd015281816166b8015281816168240152616c12015260c0518181816110be01528181611c840152818161253b01528181613d070152614104015260e051818181610fa701528181611cc70152818161257e0152613850015261010051818181610538015281816114180152818161274801528181612db1015281816142630152818161478a0152614f850152f35b634e487b7160e01b600052604160045260246000fd5b82908101031261056e5784015180159081150361056e5761051757388481610288565b815162461bcd60e51b815260048101859052602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b600080fd5b9161027f92606091610b51565b60405162461bcd60e51b815260048101859052603660248201527f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60448201527f20746f206e6f6e2d7a65726f20616c6c6f77616e6365000000000000000000006064820152608490fd5b90508481813d8311610611575b6106028183610804565b8101031261056e5751386101cb565b503d6105f8565b6040513d6000823e3d90fd5b630a64406560e11b60005260046000fd5b60405192936106448685610804565b60008452600036813760e051156107465760005b84518110156106bf576001906001600160a01b036106768288610849565b51168861068282610abd565b61068f575b505001610658565b7f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a13888610687565b50919390925060005b835181101561073c576001906001600160a01b036106e68287610849565b5116801561073657876106f882610a4b565b610706575b50505b016106c8565b7f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a138876106fd565b50610700565b509150913861018a565b6335f4a7b360e01b60005260046000fd5b60ff1660ff821681810361076b575061014f565b6332ad3e0760e11b60005260045260245260446000fd5b9091508881813d83116107b2575b61079a8183610804565b8101031261056e576107ab90610827565b9038610149565b503d610790565b506001600160a01b03821615610121565b506001600160a01b0384161561011a565b639b15e16f60e01b60005260046000fd5b602080916107f984610835565b8152019101906100af565b601f909101601f19168101906001600160401b038211908210176104de57604052565b519060ff8216820361056e57565b51906001600160a01b038216820361056e57565b805182101561085d5760209160051b010190565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b03811660009081526000805160206185ab833981519152602052604090205460ff166108fb576001600160a01b031660008181526000805160206185ab83398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b805482101561085d5760005260206000200190600090565b805480156109425760001901906109308282610901565b8154906000199060031b1b1916905555565b634e487b7160e01b600052603160045260246000fd5b6000818152601360205260409020548015610a19576000198101818111610a0357601254600019810191908211610a03578082036109b2575b50505061099e6012610919565b600052601360205260006040812055600190565b6109eb6109c36109d4936012610901565b90549060031b1c9283926012610901565b819391549060031b91821b91600019901b19161790565b90556000526013602052604060002055388080610991565b634e487b7160e01b600052601160045260246000fd5b5050600090565b805490680100000000000000008210156104de57816109d4916001610a4794018155610901565b9055565b806000526004602052604060002054156000146108fb57610a6d816003610a20565b600354906000526004602052604060002055600190565b806000526013602052604060002054156000146108fb57610aa6816012610a20565b601254906000526013602052604060002055600190565b6000818152600460205260409020548015610a19576000198101818111610a0357600354600019810191908211610a0357818103610b17575b505050610b036003610919565b600052600460205260006040812055600190565b610b39610b286109d4936003610901565b90549060031b1c9283926003610901565b90556000526004602052604060002055388080610af6565b91929015610bb35750815115610b65575090565b3b15610b6e5790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610bc65750805190602001fd5b6040519062461bcd60e51b8252602060048301528181519182602483015260005b838110610c095750508160006044809484010152601f80199101168101030190fd5b60208282018101516044878401015285935001610be756fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a71461510f575080630a861f2a14614f0c5780630bd7c46d14614ebe578063164e68de14614dd6578063181f5a7714614d7557806321df0da714614d31578063240028e814614cdd5780632451a62714614c68578063248a9ca314614c3a57806324f65ee714614bfc5780632a10097b146149af5780632c286daf146148ad5780632d4a148f146146635780632f2ff15d1461462257806331238ffc146145de57806336568abe1461458057806337b19247146144795780633907753714614009578063432a6ba314613fe2578063489a68f214613be55780634ad01f0b14613b535780634c5ef0ed14613b0c5780634f71592c14613ad757806350d1a35a1461397b57806354c8a4f31461381c5780635df45a371461380157806362ddd3c41461377d5780636600f92c1461368d5780636cfd1553146135ec5780636d9d216c146131d5578063714bf90714613153578063791e5a101461311857806379ba50971461306b578063804ba5a9146130285780638632d5cc14612ff35780638926f54f14612fad5780638a5e52bb14612d4c5780638da5cb5b14612d2557806391a2749a14612b9557806391d1485414612b48578063962d40201461298e5780639a4575b9146124515780639f68f67314612419578063a217fddf146123fd578063a42a7b8b146122b4578063a7cd63b714612233578063acfecf9114612115578063af0e58b9146120f7578063af58d59f146120ae578063b0f479a114612087578063b1c71c6514611b78578063b79019b5146118d5578063b79465801461189c578063c0d78655146117ce578063c4bffe2b146116c1578063c75eea9c14611619578063cd306a6c146115ee578063ce3c75281461136b578063cf7401f31461118e578063d547741f14611144578063da90a9f3146110e2578063dc0bd9711461109e578063de814c5714610fcc578063e0351e1314610f8f578063e58d80c714610f24578063e8a1da17146106a9578063eb521a4c1461043f578063f1e7339914610413578063f2fde38b1461035e5763f65a88861461032357600080fd5b3461035b57602060031936011261035b57604060209167ffffffffffffffff61034a6153ee565b168152601583522054604051908152f35b80fd5b503461035b57602060031936011261035b576001600160a01b036103806152a6565b610388615e4b565b163381146103eb57807fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782556001600160a01b03600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b6004827fdad89dca000000000000000000000000000000000000000000000000000000008152fd5b503461035b57602060031936011261035b5760206104376104326153ee565b615dc4565b604051908152f35b503461035b57602060031936011261035b576004358180526017602052604082205461067d578015610655576001600160a01b0361047c83615a5e565b16330361062957818052601160205260408220600181015460a01c60ff1615610614576104aa828254615a41565b90555b6040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064810182905282907f00000000000000000000000000000000000000000000000000000000000000009061052e9061052881608481015b03601f1981018352826154ee565b82616e30565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610610576040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039290921660048301526024820184905282908290604490829084905af18015610605576105ec575b50506040519082825260208201527f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf60403392a280f35b816105f6916154ee565b6106015781386105b5565b5080fd5b6040513d84823e3d90fd5b8280fd5b5061062181600f54615a41565b600f556104ad565b6024827f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b6004827fa90c0d19000000000000000000000000000000000000000000000000000000008152fd5b6024827f6469724600000000000000000000000000000000000000000000000000000000815280600452fd5b503461035b576106b8366155ed565b939190926106c4615e4b565b82915b808310610d8f575050508063ffffffff4216917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1843603015b85821015610d8b578160051b85013581811215610d875785019061012082360312610d875760405195610732876154d2565b823567ffffffffffffffff81168103610d82578752602083013567ffffffffffffffff8111610d7e5783019536601f88011215610d7e578635966107758861567c565b97610783604051998a6154ee565b8089526020808a019160051b83010190368211610d7a5760208301905b828210610d47575050505060208801968752604084013567ffffffffffffffff8111610d43576107d39036908601615564565b9860408901998a526107fd6107eb3660608801615788565b9560608b0196875260c0369101615788565b9660808a0197885261080f8651616ce9565b6108198851616ce9565b8a515115610d1b5761083567ffffffffffffffff8b511661758d565b15610ce45767ffffffffffffffff8a5116815260086020526040812061097587516fffffffffffffffffffffffffffffffff604082015116906109306fffffffffffffffffffffffffffffffff6020830151169151151583608060405161089b816154d2565b858152602081018c905260408101849052606081018690520152855474ff000000000000000000000000000000000000000091151560a01b919091167fffffffffffffffffffffff0000000000000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff84161773ffffffff0000000000000000000000000000000060808b901b1617178555565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176001830155565b610a9b89516fffffffffffffffffffffffffffffffff60408201511690610a566fffffffffffffffffffffffffffffffff602083015116915115158360806040516109bf816154d2565b858152602081018c9052604081018490526060810186905201526002860180547fffffffffffffffffffffff000000000000000000000000000000000000000000166fffffffffffffffffffffffffffffffff85161773ffffffff0000000000000000000000000000000060808c901b161791151560a01b74ff000000000000000000000000000000000000000016919091179055565b60809190911b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff91909116176003830155565b60048c5191019080519067ffffffffffffffff8211610cb757610abe8354615add565b601f8111610c7c575b50602090601f8311600114610c1957610af79291859183610c0e575b50506000198260011b9260031b1c19161790565b90555b805b89518051821015610b325790610b2c600192610b25838f67ffffffffffffffff90511692615ac9565b5190616846565b01610afc565b5050975097987f8d340f17e19058004c20453540862a9c62778504476f6756755cb33bcd6c38c292959396610c0067ffffffffffffffff600197949c5116925193519151610bcc610b97604051968796875261010060208801526101008701906152e6565b9360408601906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60a08401906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b0390a1019093949291610700565b015190503880610ae3565b8385528185209190601f198416865b818110610c645750908460019594939210610c4b575b505050811b019055610afa565b015160001960f88460031b161c19169055388080610c3e565b92936020600181928786015181550195019301610c28565b610ca79084865260208620601f850160051c81019160208610610cad575b601f0160051c0190615d8b565b38610ac7565b9091508190610c9a565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60249067ffffffffffffffff8b51167f1d5ad3c5000000000000000000000000000000000000000000000000000000008252600452fd5b807f14c880ca0000000000000000000000000000000000000000000000000000000060049252fd5b8680fd5b813567ffffffffffffffff8111610d7657602091610d6b8392833691890101615564565b8152019101906107a0565b8a80fd5b8880fd5b8580fd5b600080fd5b8380fd5b8280f35b9092919367ffffffffffffffff610daf610daa878588615852565b61580e565b1695610dba876176e8565b15610ef8578684526008602052610dd660056040862001616f63565b94845b8651811015610e0f576001908987526008602052610e0860056040892001610e01838b615ac9565b519061777c565b5001610dd9565b5093945094909580855260086020526005604086208681558660018201558660028201558660038201558660048201610e488154615add565b80610eb7575b5050500180549086815581610e99575b5050907f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599166020600193604051908152a10191909493946106c7565b865260208620908101905b81811015610e5e57868155600101610ea4565b601f8111600114610ecd5750555b863880610e4e565b81835260208320610ee891601f01861c810190600101615d8b565b8082528160208120915555610ec5565b602484887f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b503461035b57602060031936011261035b576001600160a01b036040610f486152a6565b927f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af81526002602052209116600052602052602060ff604060002054166040519015158152f35b503461035b578060031936011261035b5760206040517f000000000000000000000000000000000000000000000000000000000000000015158152f35b503461035b57604060031936011261035b57610fe66153ee565b602435610ff1615e4b565b67ffffffffffffffff8060145460a01c1692168092036110765760407fe1e6c22ce6b566f66cdb457ec2e7910ff1f9a9e5654ed75303476fa870468220918385526015602052818520611045828254615a41565b905583855260116020526110688286205485875260156020528387205490615862565b82519182526020820152a280f35b6004837fa94cb988000000000000000000000000000000000000000000000000000000008152fd5b503461035b578060031936011261035b5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461035b57602060031936011261035b577fd63806009f622849e3b7cfd82d762420d57574c39f945f678871b2b5f1e8ce4b602061111f6152a6565b611127615e4b565b611130816163d6565b506001600160a01b0360405191168152a180f35b503461035b57604060031936011261035b5761118a6004356111646152bc565b9061118561118082600052600260205260016040600020015490565b616376565b6164c0565b5080f35b503461035b5760e060031936011261035b576111a86153ee565b9060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc36011261035b576040516111df8161549a565b60243580151581036106105781526044356fffffffffffffffffffffffffffffffff811681036106105760208201526064356fffffffffffffffffffffffffffffffff8116810361061057604082015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7c36011261060157604051906112668261549a565b6084358015158103610d8757825260a4356fffffffffffffffffffffffffffffffff81168103610d8757602083015260c4356fffffffffffffffffffffffffffffffff81168103610d875760408301527f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af83526002602052604083206001600160a01b03331660005260205260ff60406000205416158061133f575b611313576113109293616a56565b80f35b6024837f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b508280526002602052604083206001600160a01b03331660005260205260ff6040600020541615611302565b503461035b57604060031936011261035b576113856153ee565b60243567ffffffffffffffff8216808452601160205260ff600160408620015460a01c161580156115e6575b6115bb578115611593576001600160a01b036113cc84615a5e565b163303611567578352601160205260408320600181015460a01c60ff16801561155f5781545b80841161152f57501561151a5761140a828254615862565b90555b826001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610601576040517f69328dec0000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166004820152602481018490523360448201529082908290606490829084905af1801561060557611505575b50506040805167ffffffffffffffff9093168352602083019190915233917f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e91819081015b0390a280f35b8161150f916154ee565b6106105782386114ba565b5061152781600f54615862565b600f5561140d565b85846044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600f546113f2565b6024847f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b6004847fa90c0d19000000000000000000000000000000000000000000000000000000008152fd5b7f46f5f12b000000000000000000000000000000000000000000000000000000008452600452602483fd5b5080156113b1565b503461035b578060031936011261035b57602067ffffffffffffffff60145460a01c16604051908152f35b503461035b57602060031936011261035b5761166461165f60406116bd9367ffffffffffffffff6116486153ee565b611650615c11565b50168152600860205220615c3c565b616b86565b6040519182918291909160806fffffffffffffffffffffffffffffffff8160a084019582815116855263ffffffff6020820151166020860152604081015115156040860152826060820151166060860152015116910152565b0390f35b503461035b578060031936011261035b57604051906006548083528260208101600684526020842092845b8181106117b5575050611701925003836154ee565b815161172561170f8261567c565b9161171d60405193846154ee565b80835261567c565b91601f19602083019301368437805b8451811015611766578067ffffffffffffffff61175360019388615ac9565b511661175f8286615ac9565b5201611734565b50925090604051928392602084019060208552518091526040840192915b818110611792575050500390f35b825167ffffffffffffffff16845285945060209384019390920191600101611784565b84548352600194850194879450602090930192016116ec565b503461035b57602060031936011261035b576117e86152a6565b6117f0615e4b565b6001600160a01b03811690811561187457600580547fffffffffffffffffffffffff000000000000000000000000000000000000000081169093179055604080516001600160a01b0393841681529190921660208201527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f168491819081015b0390a180f35b6004837f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b503461035b57602060031936011261035b576116bd6118c16118bc6153ee565b615da2565b6040519182916020835260208301906152e6565b503461035b57602060031936011261035b5760043567ffffffffffffffff81116106015761190790369060040161536a565b61190f615e4b565b825b81811061191c578380f35b61192a610daa828486615ca2565b611942611938838587615ca2565b6020810190615ce2565b907fb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c61197c61197286888a615ca2565b6040810190615ce2565b91909261199261198d368784615694565b616c46565b6119a061198d368587615694565b604051946119ad866154b6565b6119b8368284615694565b8652611a0267ffffffffffffffff6119d1368789615694565b9860208901998a52169586956119f4604051958695604087526040870191615d36565b918483036020860152615d36565b0390a28652600d60205260408620905180519067ffffffffffffffff8211611b4b57680100000000000000008211611b4b576020908354838555808410611b31575b500182885260208820885b838110611b14575050505060010190519081519167ffffffffffffffff8311611ae757680100000000000000008311611ae7576020908254848455808510611acd575b500190865260208620865b838110611ab05750505050600101611911565b60019060206001600160a01b038551169401938184015501611a9d565b838952828920611ae1918101908601615d8b565b38611a92565b6024877f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b60019060206001600160a01b038551169401938184015501611a4f565b848a52828a20611b45918101908501615d8b565b38611a44565b6024887f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b503461035b57606060031936011261035b5760043567ffffffffffffffff81116106015760a0600319823603011261060157611bb261539b565b9060443567ffffffffffffffff8111610d8757611bd3903690600401615564565b50611bdc615ab0565b506084810191611beb8361589e565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361204a57602482019177ffffffffffffffff00000000000000000000000000000000611c448461580e565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561203f578691612010575b50611fe857611cc56044820161589e565b7f0000000000000000000000000000000000000000000000000000000000000000611f99575b50606461ffff91611d03611cfe8661580e565b6171e1565b013591169384151593848095611f8a575b15611eed5761ffff600a5416808710611ebd5750611e819550611d66611d56611d3c8661580e565b67ffffffffffffffff16600052600b602052604060002090565b84611d608461589e565b91617279565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff611da2611d9c8761580e565b9361589e565b604080516001600160a01b03929092168252602082018790529190931692a25b508092611e8b575b506118bc81611ddb611e509361580e565b604080516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016815233602082015290810186905267ffffffffffffffff91909116907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a261580e565b90611e59616c0b565b60405192611e66846154b6565b83526020830152604051928392604084526040840190615734565b9060208301520390f35b611e50919250611eb56118bc91612710611eae61ffff600a5460101c1683615d78565b0490615862565b929150611dca565b82604491887fe08f03ef000000000000000000000000000000000000000000000000000000008352600452602452fd5b50611e81945067ffffffffffffffff611f058461580e565b1680825260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da81789448380611f6a604086206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391617279565b604080516001600160a01b039290921682526020820192909252a2611dc2565b5061ffff600a54161515611d14565b6001600160a01b0316611fb9816000526004602052604060002054151590565b611ceb577fd0d25976000000000000000000000000000000000000000000000000000000008652600452602485fd5b6004857f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612032915060203d602011612038575b61202a81836154ee565b810190616798565b38611cb4565b503d612020565b6040513d88823e3d90fd5b6024846001600160a01b0361205e8661589e565b7f961c9a4f00000000000000000000000000000000000000000000000000000000835216600452fd5b503461035b578060031936011261035b5760206001600160a01b0360055416604051908152f35b503461035b57602060031936011261035b5761166461165f600260406116bd9467ffffffffffffffff6120df6153ee565b6120e7615c11565b5016815260086020522001615c3c565b503461035b578060031936011261035b576020600f54604051908152f35b503461035b5767ffffffffffffffff61212d3661563b565b929091612138615e4b565b1691612151836000526007602052604060002054151590565b156122075782845260086020526121806005604086200161217336848661552d565b602081519101209061777c565b156121bf57907f52d00ee4d9bd51b40168f2afc5848837288ce258784ad914278791464b3f4d76916114ff604051928392602084526020840191615bf0565b82612203836040519384937f74f23c7c0000000000000000000000000000000000000000000000000000000085526004850152604060248501526044840191615bf0565b0390fd5b602484847f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b503461035b578060031936011261035b5760405160038054808352908352909160208301917fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b915b81811061229e576116bd85612292818703826154ee565b60405191829182615327565b825484526020909301926001928301920161227b565b503461035b57602060031936011261035b5767ffffffffffffffff6122d76153ee565b16815260086020526122ee60056040832001616f63565b8051601f196123156122ff8361567c565b9261230d60405194856154ee565b80845261567c565b01835b8181106123ec575050825b8251811015612369578061233960019285615ac9565b518552600960205261234d60408620615b30565b6123578285615ac9565b526123628184615ac9565b5001612323565b81846040519182916020830160208452825180915260408401602060408360051b870101940192905b8282106123a157505050500390f35b919360206123dc827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0600195979984950301865288516152e6565b9601920192018594939192612392565b806060602080938601015201612318565b503461035b578060031936011261035b57602090604051908152f35b503461035b57612292600160406116bd9367ffffffffffffffff61243c36615582565b505050509050168152600d6020522001615940565b503461035b57602060031936011261035b5760043567ffffffffffffffff81116106015760a060031982360301126106015761248b615ab0565b50612494615ab0565b50608481016124a28161589e565b6001600160a01b03807f00000000000000000000000000000000000000000000000000000000000000001691160361297a57602482019177ffffffffffffffff000000000000000000000000000000006124fb8461580e565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa90811561296f578591612950575b506129285761257c6044820161589e565b7f00000000000000000000000000000000000000000000000000000000000000006128d9575b506064906125b2611cfe8561580e565b013590831561283f5761ffff600a54168061280f57506125e16125d7611d3c8561580e565b83611d608461589e565b7f7c5343c904d7bdd0794d318f4681059f06df378f04bd8aa69d054ac065f300b267ffffffffffffffff612617611d9c8661580e565b604080516001600160a01b03929092168252602082018690529190931692a25b826126418361580e565b604080516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001680825233602083015291810185905290939167ffffffffffffffff16907ff33bc26b4413b0e7f19f1ea739fdf99098c0061f1f87d954b11f5293fad9ae1090606090a26126be6118bc8561580e565b936126c7616c0b565b604051956126d4876154b6565b8652602086015267ffffffffffffffff6126ed8261580e565b1683526011602052604083206001015460a01c60ff16156127fa578067ffffffffffffffff61271e61273c9361580e565b168452601160205260408420612735848254615a41565b905561580e565b505b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610610576040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0394909416600485015260248401919091528290604490829084905af180156127ef576127da575b604051602080825281906116bd90820185615734565b6127e58380926154ee565b61060157816127c4565b6040513d85823e3d90fd5b5061280781600f54615a41565b600f5561273e565b846044917fe08f03ef00000000000000000000000000000000000000000000000000000000825281600452602452fd5b5067ffffffffffffffff6128528361580e565b168060005260086020527fff0133389f9bb82d5b9385826160eaf2328039f6fa950eeb8cf0836da817894482806128b960406000206001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391617279565b604080516001600160a01b039290921682526020820192909252a2612637565b6001600160a01b03166128f9816000526004602052604060002054151590565b6125a2577fd0d25976000000000000000000000000000000000000000000000000000000008552600452602484fd5b6004847f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b612969915060203d6020116120385761202a81836154ee565b3861256b565b6040513d87823e3d90fd5b826001600160a01b0361205e60249361589e565b503461035b57606060031936011261035b5760043567ffffffffffffffff8111610601576129c090369060040161536a565b60243567ffffffffffffffff8111610d87576129e0903690600401615703565b60449291923567ffffffffffffffff8111610d7e57612a03903690600401615703565b9190927f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af87526002602052604087206001600160a01b03331660005260205260ff604060002054161580612b1c575b612af057818114801590612ae6575b612abe57865b818110612a72578780f35b80612ab8612a86610daa600194868c615852565b612a9183878b615aa0565b612ab2612aaa612aa2868b8d615aa0565b923690615788565b913690615788565b91616a56565b01612a67565b6004877f568efce2000000000000000000000000000000000000000000000000000000008152fd5b5082811415612a61565b6024877f8e4a23d600000000000000000000000000000000000000000000000000000000815233600452fd5b508680526002602052604087206001600160a01b03331660005260205260ff6040600020541615612a52565b503461035b57604060031936011261035b576001600160a01b036040612b6c6152bc565b9260043581526002602052209116600052602052602060ff604060002054166040519015158152f35b503461035b57602060031936011261035b5760043567ffffffffffffffff811161060157604060031982360301126106015760405190612bd4826154b6565b806004013567ffffffffffffffff8111610d8757612bf890600436918401016156e8565b825260248101359067ffffffffffffffff8211610d87576004612c1e92369201016156e8565b60208201908152612c2d615e4b565b5191805b8351811015612c9757806001600160a01b03612c4f60019387615ac9565b5116612c5a8161781c565b612c66575b5001612c31565b60207fc3803387881faad271c47728894e3e36fac830ffc8602ca6fc07733cbda7758091604051908152a138612c5f565b509051815b8151811015610d8b576001600160a01b03612cb78284615ac9565b51168015612cfd57907feb1b9b92e50b7f88f9ff25d56765095ac6e91540eee214906f4036a908ffbdef602083612cef600195617554565b50604051908152a101612c9c565b6004847f8579befe000000000000000000000000000000000000000000000000000000008152fd5b503461035b578060031936011261035b5760206001600160a01b0360015416604051908152f35b503461035b578060031936011261035b576014546001600160a01b0381163303612f855760a01c67ffffffffffffffff168015612f5d578082526011602052612da660408320548284526015602052604084205490615862565b826001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690803b15610610576040517f69328dec0000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602481018590523060448201529083908290606490829084905af19081156127ef578391612f48575b5084905260116020528160408120557fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff60145416601455803b15610601578180916024604051809481937f42966c680000000000000000000000000000000000000000000000000000000083528860048401525af1801561060557612f33575b505081612f0e7fdea60ddd4c7ebdab804f5694c70350cca7893ece3efeecb142312eacac5c73e49361751b565b506040805167ffffffffffffffff90921682526020820192909252908190810161186e565b81612f3d916154ee565b610610578238612ee1565b81612f52916154ee565b610601578138612e61565b6004827fa94cb988000000000000000000000000000000000000000000000000000000008152fd5b6004827f5fff6eee000000000000000000000000000000000000000000000000000000008152fd5b503461035b57602060031936011261035b576020612fe967ffffffffffffffff612fd56153ee565b166000526007602052604060002054151590565b6040519015158152f35b503461035b57602060031936011261035b5760206130176130126153ee565b615a5e565b6001600160a01b0360405191168152f35b503461035b57602060031936011261035b5760043567ffffffffffffffff81116106015761305d6113109136906004016153bd565b90613066615e4b565b616050565b503461035b578060031936011261035b5780546001600160a01b03811633036130f0577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551682556001600160a01b033391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6004827f02b543c6000000000000000000000000000000000000000000000000000000008152fd5b503461035b578060031936011261035b5760206040517f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af8152f35b503461035b57602060031936011261035b577f084e6f0e9791c2e56153bd49e6ec6dd63ba9a72c258d71558d74c63fc75b716860206001600160a01b036131986152a6565b6131a0615e4b565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000006014541617601455604051908152a180f35b503461035b57604060031936011261035b5760043567ffffffffffffffff81116106015761320790369060040161536a565b906024359067ffffffffffffffff8211610d875736602383011215610d875781600401359267ffffffffffffffff84116135e8576024830192602436918660061b0101116135e857613257615e4b565b845b8181106134c257505050825b828110613270578380f35b67ffffffffffffffff613287610daa838686615a4e565b1615801561348f575b801561346e575b61342b576001600160a01b036132b960206132b3848787615a4e565b0161589e565b1615613403578061339e84846001600160a01b038060018a67ffffffffffffffff613318610daa8a6132f260206132b3889f8d8d615a4e565b99604051956133008761549a565b865287602087019b168b526040860199878b52615a4e565b168c52601160205260408c209051815501935116167fffffffffffffffffffffffff00000000000000000000000000000000000000008354161782555115157fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b7f180c6940bd64ba8f75679203ca32f8be2f629477a3307b190656e4b14dd5ddeb60406133cf610daa848888615a4e565b6001600160a01b036133e760206132b3878b8b615a4e565b67ffffffffffffffff845193168352166020820152a101613265565b6004847f14c880ca000000000000000000000000000000000000000000000000000000008152fd5b610daa613445916024959467ffffffffffffffff94615a4e565b7fd9a9cd6800000000000000000000000000000000000000000000000000000000835216600452fd5b5061348967ffffffffffffffff612fd5610daa848787615a4e565b15613297565b5067ffffffffffffffff6134a7610daa838686615a4e565b168452601160205260ff600160408620015460a01c16613290565b67ffffffffffffffff6134d9610daa838587615852565b168652601160205260ff600160408820015460a01c16156135a5578067ffffffffffffffff61350e610daa6001948688615852565b16875260116020527f7b5efb3f8090c5cfd24e170b667d0e2b6fdc3db6540d75b86d5b6655ba00eb93604088205461354881600f54615a41565b600f5567ffffffffffffffff613562610daa85888a615852565b16895260116020528884604082208281550155613583610daa848789615852565b6040805167ffffffffffffffff9290921682526020820192909252a101613259565b6135bf610daa67ffffffffffffffff928894602496615852565b7f46f5f12b00000000000000000000000000000000000000000000000000000000835216600452fd5b8480fd5b503461035b57602060031936011261035b577f66b1c1bdec8b60a3442bb25b5b6cd6fff3d0eceb6f5390be8e2f82a8ad39b2346001600160a01b0361362f6152a6565b613637615e4b565b61186e601054918381167fffffffffffffffffffffffff000000000000000000000000000000000000000084161760105560405193849316839092916001600160a01b0360209181604085019616845216910152565b503461035b57604060031936011261035b576136a76153ee565b67ffffffffffffffff6136b86152bc565b916136c1615e4b565b16908183526011602052600160408420019081549160ff8360a01c16156137515780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556040805192909316825260208201527f01efd4cd7dd64263689551000d4359d6559c839f39b773b1df3fd19ff060cf5f91819081016114ff565b602485857f46f5f12b000000000000000000000000000000000000000000000000000000008252600452fd5b503461035b5761378c3661563b565b61379893929193615e4b565b67ffffffffffffffff82166137ba816000526007602052604060002054151590565b156137d6575061131092936137d091369161552d565b90616846565b7f1e670e4b000000000000000000000000000000000000000000000000000000008452600452602483fd5b503461035b578060031936011261035b576020610437615996565b503461035b576138469061384e613832366155ed565b959161383f939193615e4b565b3691615694565b933691615694565b7f00000000000000000000000000000000000000000000000000000000000000001561395357815b83518110156138dc57806001600160a01b0361389460019387615ac9565b511661389f81617654565b6138ab575b5001613876565b60207f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf756691604051908152a1386138a4565b5090805b825181101561118a57806001600160a01b036138fe60019386615ac9565b5116801561394d5761390f816174e2565b61391c575b505b016138e0565b60207f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d891604051908152a184613914565b50613916565b6004827f35f4a7b3000000000000000000000000000000000000000000000000000000008152fd5b503461035b57602060031936011261035b576139956153ee565b61399d615e4b565b60145467ffffffffffffffff8160a01c16613aaf5767ffffffffffffffff8216916139d5836000526017602052604060002054151590565b613a83578215613a5757916020917fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff7bffffffffffffffff00000000000000000000000000000000000000007f20331f191af84dbff48b162aa5a5985e7891ae646297b0a2ac80487f9109ef499560a01b16911617601455604051908152a180f35b6024847fd9a9cd6800000000000000000000000000000000000000000000000000000000815280600452fd5b602484847f1c49a87b000000000000000000000000000000000000000000000000000000008252600452fd5b6004837f692bc131000000000000000000000000000000000000000000000000000000008152fd5b503461035b5761229260406116bd9267ffffffffffffffff613af836615582565b505050509050168152600d60205220615940565b503461035b57604060031936011261035b57613b266153ee565b906024359067ffffffffffffffff821161035b576020612fe984613b4d3660048701615564565b90615903565b503461035b578060031936011261035b57613b6c615e4b565b60145467ffffffffffffffff8160a01c16908115611076577f375f1ad1194a2bec317c5efec05cc63ffa06ddd0c4b276619f6fd47298eda518917fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff6020921660145580845260158252836040812055604051908152a180f35b503461035b57604060031936011261035b576004359067ffffffffffffffff821161035b5781600401610100600319843603011261060157613c2561539b565b82604051613c328161547e565b52613c5f613c55613c50613c4960c48801866158b2565b369161552d565b6167b0565b60648601356166b5565b926084850191613c6e8361589e565b6001600160a01b03807f000000000000000000000000000000000000000000000000000000000000000016911603613fce57602486019377ffffffffffffffff00000000000000000000000000000000613cc78661580e565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115613fc3578491613fa4575b50613f7c57613d48611cfe8661580e565b613d518561580e565b90613d6760a4890192613b4d613c4985856158b2565b15613f35575050611d9c60446020977ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09567ffffffffffffffff95898961ffff613e3898161515600014613ea057604091509188613dc7613dd99461580e565b168152600c8d52208a611d608461589e565b7f41a8aa8df7945f0fb8ac5f7d88279638d9dc2ef9a6bf4ec9a53b80681b34aff786613e07611d9c8b61580e565b604080516001600160a01b03929092168252602082018d90529190931692a25b0194613e328661589e565b5061580e565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081168252336020830152909216908201526060810185905292169180608081015b0390a280604051613e978161547e565b52604051908152f35b7f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c929350613f15600260408b613ed6869561580e565b169687815260206008905220016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016928391617279565b604080516001600160a01b039290921682526020820192909252a2613e27565b613f3f92506158b2565b6122036040519283927f24eb47e5000000000000000000000000000000000000000000000000000000008452602060048501526024840191615bf0565b6004837f53ad11d8000000000000000000000000000000000000000000000000000000008152fd5b613fbd915060203d6020116120385761202a81836154ee565b38613d37565b6040513d86823e3d90fd5b506001600160a01b0361205e60249361589e565b503461035b578060031936011261035b5760206001600160a01b0360105416604051908152f35b503461035b57602060031936011261035b5760043567ffffffffffffffff8111610601578060040191610100600319833603011261035b578060405161404e8161547e565b5261405c60648301356165cb565b906084830161406a8161589e565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116036144655750602483019077ffffffffffffffff000000000000000000000000000000006140c48361580e565b60801b16604051907f2cbc26bb00000000000000000000000000000000000000000000000000000000825260048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115610605578291614446575b5061441e57614145611cfe8361580e565b61414e8261580e565b61416360a4860191613b4d613c49848a6158b2565b15614414575082919067ffffffffffffffff61417e8361580e565b1680825260086020526141c3600260408420016001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016958691617279565b604080516001600160a01b0386168152602081018790527f50f6fbee3ceedce6b7fd7eaef18244487867e6718aec7208187efb6b7908c14c9190a267ffffffffffffffff6142108361580e565b16815260116020526040812067ffffffffffffffff61422e8461580e565b16825260156020526040822054806143aa5750805480861161437a578561425491615862565b90555b60446001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001695019461428f8661589e565b813b15610610576040517f69328dec0000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152602482018890529190911660448201529082908290606490829084905af1801561060557614365575b505067ffffffffffffffff602094613e8785614334611d9c7ffc5e3a5bddc11d92c2dc20fae6f7d5eb989f056be35239f7de7e86150609abc09661580e565b604080516001600160a01b039889168152336020820152979091169087015260608601529116929081906080820190565b6143708280926154ee565b61035b57806142f5565b82866044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b80915085116143e4575067ffffffffffffffff6143c68361580e565b1681526015602052604081206143dd858254615862565b9055614257565b90846044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b613f3f90866158b2565b807f53ad11d80000000000000000000000000000000000000000000000000000000060049252fd5b61445f915060203d6020116120385761202a81836154ee565b38614134565b906001600160a01b0361205e60249361589e565b503461035b5760a060031936011261035b576144936152a6565b5060243567ffffffffffffffff81168091036106015760443567ffffffffffffffff81116106105760031960a09136030112610601576144d16153ac565b5060843567ffffffffffffffff811161061057916040916144f86080943690600401615405565b5050816060845161450881615433565b828152826020820152828682015201528152600e602052206040519061452d82615433565b5463ffffffff808216928381528160208201818560201c16815260ff60606040850194848860401c168652019560601c161515855260405195865251166020850152511660408301525115156060820152f35b503461035b57604060031936011261035b5761459a6152bc565b336001600160a01b038216036145b65761118a906004356164c0565b6004827f6697b232000000000000000000000000000000000000000000000000000000008152fd5b503461035b57602060031936011261035b5760ff6001604060209367ffffffffffffffff61460a6153ee565b1681526011855220015460a01c166040519015158152f35b503461035b57604060031936011261035b5761118a6004356146426152bc565b9061465e61118082600052600260205260016040600020015490565b615f7b565b503461035b57604060031936011261035b5761467d6153ee565b60243567ffffffffffffffff8216808452601160205260ff600160408620015460a01c161580156148a5575b6115bb576146c4816000526017602052604060002054151590565b61487a578115611593576001600160a01b036146df84615a5e565b163303611567578352601160205260408320600181015460a01c60ff16156148655761470c828254615a41565b90555b6040517f23b872dd0000000000000000000000000000000000000000000000000000000060208201523360248201523060448201526064810182905283907f00000000000000000000000000000000000000000000000000000000000000009061478090610528816084810161051a565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610610576040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039290921660048301526024820184905282908290604490829084905af1801561060557614850575b50506040805167ffffffffffffffff9093168352602083019190915233917f569a440e6842b5e5a7ac02286311855f5a0b81b9390909e552e82aaf02c9e9bf91819081016114ff565b8161485a916154ee565b610610578238614807565b5061487281600f54615a41565b600f5561470f565b7f64697246000000000000000000000000000000000000000000000000000000008452600452602483fd5b5080156146a9565b503461035b57606060031936011261035b5760043561ffff8116809103610601576148d661539b565b60443567ffffffffffffffff8111610d87576148f69036906004016153bd565b906148ff615e4b565b61ffff83169261271084101561498357849260409492614975927f52aa194b292c8bfb5aaca8ee2000a965c3a051b306ff841873b16147526a39ba977fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000063ffff0000600a549360101b1692161717600a55616050565b82519182526020820152a180f35b602486857f95f3517a000000000000000000000000000000000000000000000000000000008252600452fd5b503461035b57604060031936011261035b5760043567ffffffffffffffff8111610601573660238201121561060157806004013567ffffffffffffffff81116106105760248201916024369160a084020101116106105760243567ffffffffffffffff8111610d8757614a2690369060040161536a565b919092614a31615e4b565b84805b838110614a9e57509150505b818110614a4b578380f35b8067ffffffffffffffff614a65610daa6001948688615852565b16808652600e6020528560408120557f5479bbc0288b7eaeaf2ace0943b88016cc648964fcd42919a86fd93b15fdbee88680a201614a40565b6001917f56f77aeff2def50c8b5f5a0df3bab7183df09bf36c6feba496bb42551db77d70608085614bec614bb263ffffffff614be1614ba582614bd6614af58f80614aef8f9283610daa918e6157cf565b9a6157cf565b604067ffffffffffffffff602083019a169c8d8152600e6020522083614b1a8b615823565b169181549060408101937fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff67ffffffff00000000614b5787615823565b60201b16918f6cff0000000000000000000000007fffffffffffffffffffffffffffffffffffffffff000000000000000000000000916bffffffff0000000000000000606088019d8e615823565b60401b1696019e8f615834565b151560601b1695161716171717905582614bce6040519a615841565b168952615841565b166020870152615841565b16604084015261575e565b15156060820152a2018590614a34565b503461035b578060031936011261035b57602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461035b57602060031936011261035b576020610437600435600052600260205260016040600020015490565b503461035b578060031936011261035b5760405160128054808352908352909160208301917fbb8a6a4669ba250d26cd7a459eca9d215f8307e33aebe50379bc5a3617ec3444915b818110614cc7576116bd85612292818703826154ee565b8254845260209093019260019283019201614cb0565b503461035b57602060031936011261035b576020614cf96152a6565b6001600160a01b03807f0000000000000000000000000000000000000000000000000000000000000000169116146040519015158152f35b503461035b578060031936011261035b5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461035b578060031936011261035b57506116bd604051614d986040826154ee565b601d81527f53696c6f656455534443546f6b656e506f6f6c20312e362e332d64657600000060208201526040519182916020835260208301906152e6565b503461035b57602060031936011261035b57614df06152a6565b614df8615e4b565b614e00615996565b9081614e0a578280f35b60206001600160a01b037f738b39462909f2593b7546a62adee9bc4e5cadde8e0e0f80686198081b85959992614eaf6040517fa9059cbb0000000000000000000000000000000000000000000000000000000085820152614e898161051a898660248401602090939291936001600160a01b0360408201951681520152565b7f0000000000000000000000000000000000000000000000000000000000000000616e30565b6040519485521692a238808280f35b503461035b57602060031936011261035b577ff7af318a70f367e30346e2704f6ef646b378a7dcb49767beb98a1774cd11e3896020614efb6152a6565b614f03615e4b565b61113081615e89565b503461035b57602060031936011261035b576004359081156150e7576001600160a01b03614f3982615a5e565b1633036150bb57808052601160205260408120600181015460a01c60ff1680156150b35781545b80851161508357501561506e57614f78838254615862565b90555b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15610601576040517f69328dec0000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166004820152602481018490523360448201529082908290606490829084905af180156106055761505e575b50906040519082825260208201527f58fca2457646a9f47422ab9eb9bff90cef88cd8b8725ab52b1d17baa392d784e60403392a280f35b81615068916154ee565b38615027565b5061507b82600f54615862565b600f55614f7b565b83856044927fa17e11d5000000000000000000000000000000000000000000000000000000008352600452602452fd5b600f54614f60565b807f8e4a23d6000000000000000000000000000000000000000000000000000000006024925233600452fd5b807fa90c0d190000000000000000000000000000000000000000000000000000000060049252fd5b905034610601576020600319360112610601576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361061057602092507ff208a58f00000000000000000000000000000000000000000000000000000000811490811561527c575b8115615252575b8115615228575b81156151fe575b81156151a1575b5015158152f35b7f7965db0b000000000000000000000000000000000000000000000000000000008114915081156151d4575b503861519a565b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014386151cd565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081149150615193565b7f0e64dd29000000000000000000000000000000000000000000000000000000008114915061518c565b7f1ef5498f0000000000000000000000000000000000000000000000000000000081149150615185565b7faff2afbf000000000000000000000000000000000000000000000000000000008114915061517e565b600435906001600160a01b0382168203610d8257565b602435906001600160a01b0382168203610d8257565b35906001600160a01b0382168203610d8257565b919082519283825260005b848110615312575050601f19601f8460006020809697860101520116010190565b806020809284010151828286010152016152f1565b602060408183019282815284518094520192019060005b81811061534b5750505090565b82516001600160a01b031684526020938401939092019160010161533e565b9181601f84011215610d825782359167ffffffffffffffff8311610d82576020808501948460051b010111610d8257565b6024359061ffff82168203610d8257565b6064359061ffff82168203610d8257565b9181601f84011215610d825782359167ffffffffffffffff8311610d825760208085019460e08502010111610d8257565b6004359067ffffffffffffffff82168203610d8257565b9181601f84011215610d825782359167ffffffffffffffff8311610d825760208381860195010111610d8257565b6080810190811067ffffffffffffffff82111761544f57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6020810190811067ffffffffffffffff82111761544f57604052565b6060810190811067ffffffffffffffff82111761544f57604052565b6040810190811067ffffffffffffffff82111761544f57604052565b60a0810190811067ffffffffffffffff82111761544f57604052565b90601f601f19910116810190811067ffffffffffffffff82111761544f57604052565b67ffffffffffffffff811161544f57601f01601f191660200190565b92919261553982615511565b9161554760405193846154ee565b829481845281830111610d82578281602093846000960137010152565b9080601f83011215610d825781602061557f9335910161552d565b90565b60a0600319820112610d82576004356001600160a01b0381168103610d82579160243567ffffffffffffffff81168103610d8257916044359160643561ffff81168103610d8257916084359067ffffffffffffffff8211610d82576155e991600401615405565b9091565b6040600319820112610d825760043567ffffffffffffffff8111610d8257816156189160040161536a565b929092916024359067ffffffffffffffff8211610d82576155e99160040161536a565b906040600319830112610d825760043567ffffffffffffffff81168103610d8257916024359067ffffffffffffffff8211610d82576155e991600401615405565b67ffffffffffffffff811161544f5760051b60200190565b9291906156a08161567c565b936156ae60405195866154ee565b602085838152019160051b8101928311610d8257905b8282106156d057505050565b602080916156dd846152d2565b8152019101906156c4565b9080601f83011215610d825781602061557f93359101615694565b9181601f84011215610d825782359167ffffffffffffffff8311610d825760208085019460608502010111610d8257565b61557f91602061574d83516040845260408401906152e6565b9201519060208184039101526152e6565b35908115158203610d8257565b35906fffffffffffffffffffffffffffffffff82168203610d8257565b9190826060910312610d82576040516157a08161549a565b60406157ca8183956157b18161575e565b85526157bf6020820161576b565b60208601520161576b565b910152565b91908110156157df5760a0020190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3567ffffffffffffffff81168103610d825790565b3563ffffffff81168103610d825790565b358015158103610d825790565b359063ffffffff82168203610d8257565b91908110156157df5760051b0190565b9190820391821161586f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b356001600160a01b0381168103610d825790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610d82570180359067ffffffffffffffff8211610d8257602001918136038313610d8257565b9067ffffffffffffffff61557f92166000526008602052600560406000200190602081519101209060019160005201602052604060002054151590565b906040519182815491828252602082019060005260206000209260005b818110615974575050615972925003836154ee565b565b84546001600160a01b031683526001948501948794506020909301920161595d565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa908115615a3557600091615a06575090565b90506020813d602011615a2d575b81615a21602093836154ee565b81010312610d82575190565b3d9150615a14565b6040513d6000823e3d90fd5b9190820180921161586f57565b91908110156157df5760061b0190565b67ffffffffffffffff16600052601160205260016040600020015460ff8160a01c16615a9457506001600160a01b036010541690565b6001600160a01b031690565b91908110156157df576060020190565b60405190615abd826154b6565b60606020838281520152565b80518210156157df5760209160051b010190565b90600182811c92168015615b26575b6020831014615af757565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b91607f1691615aec565b9060405191826000825492615b4484615add565b8084529360018116908115615bb05750600114615b69575b50615972925003836154ee565b90506000929192526020600020906000915b818310615b945750509060206159729282010138615b5c565b6020919350806001915483858901015201910190918492615b7b565b602093506159729592507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0091501682840152151560051b82010138615b5c565b601f8260209493601f19938186528686013760008582860101520116010190565b60405190615c1e826154d2565b60006080838281528260208201528260408201528260608201520152565b90604051615c49816154d2565b60806001829460ff81546fffffffffffffffffffffffffffffffff8116865263ffffffff81861c16602087015260a01c161515604085015201546fffffffffffffffffffffffffffffffff81166060840152811c910152565b91908110156157df5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa181360301821215610d82570190565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301821215610d82570180359067ffffffffffffffff8211610d8257602001918160051b36038313610d8257565b9160209082815201919060005b818110615d505750505090565b9091926020806001926001600160a01b03615d6a886152d2565b168152019401929101615d43565b8181029291811591840414171561586f57565b818110615d96575050565b60008155600101615d8b565b67ffffffffffffffff16600052600860205261557f6004604060002001615b30565b67ffffffffffffffff16615de5816000526007602052604060002054151590565b15615e1e5780600052601160205260ff60016040600020015460a01c16615e0d5750600f5490565b600052601160205260406000205490565b7fd9a9cd680000000000000000000000000000000000000000000000000000000060005260045260246000fd5b6001600160a01b03600154163303615e5f57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b6001600160a01b03811660009081527fa0bc845bafb3af92358952eeb50d6b57502488fa7f04f1920acc69cb8f116b20602052604090205460ff16615f75576001600160a01b031660008181527fa0bc845bafb3af92358952eeb50d6b57502488fa7f04f1920acc69cb8f116b206020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790553391907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9080a4600190565b50600090565b80600052600260205260406000206001600160a01b03831660005260205260ff604060002054161560001461602c5780600052600260205260406000206001600160a01b038316600052602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b356fffffffffffffffffffffffffffffffff81168103610d825790565b9160005b828110156163705760e081028401600061606d8261580e565b9067ffffffffffffffff821691616091836000526007602052604060002054151590565b156163445761615a92604085936161056160ff946160ff6160c5602060019c9b0192611d3c6160c03686615788565b616ce9565b91825463ffffffff8160801c16159081616326575b81616317575b816162fc575b816162ed575b50806162de575b616253575b3690615788565b90616fae565b60808501926161176160c03686615788565b8152600c6020522092835463ffffffff8160801c16159081616235575b81616226575b8161620b575b816161fc575b50806161ed575b616160575b503690615788565b01616054565b61617d60a06fffffffffffffffffffffffffffffffff9201616033565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff161717835538616152565b506161f782615834565b61614d565b60ff915060a01c161538616146565b6fffffffffffffffffffffffffffffffff8116159150616140565b8589015460801c15915061613a565b858901546fffffffffffffffffffffffffffffffff16159150616134565b6fffffffffffffffffffffffffffffffff61626f878b01616033565b845473ffffffff000000000000000000000000000000004260801b167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116919092167fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff16171783556160f8565b506162e881615834565b6160f3565b60ff915060a01c1615386160ec565b6fffffffffffffffffffffffffffffffff81161591506160e6565b848e015460801c1591506160e0565b848e01546fffffffffffffffffffffffffffffffff161591506160da565b506024917f1e670e4b000000000000000000000000000000000000000000000000000000008252600452fd5b50915050565b80600052600260205260406000206001600160a01b03331660005260205260ff60406000205416156163a55750565b7fe2517d3f000000000000000000000000000000000000000000000000000000006000523360045260245260446000fd5b6001600160a01b03811660009081527fa0bc845bafb3af92358952eeb50d6b57502488fa7f04f1920acc69cb8f116b20602052604090205460ff1615615f75576001600160a01b031660008181527fa0bc845bafb3af92358952eeb50d6b57502488fa7f04f1920acc69cb8f116b206020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690553391907f1e2af826b947397cb8f2b6a77511b5c805f9cbc82085d4c1f3e92bd927e9c5af907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9080a4600190565b80600052600260205260406000206001600160a01b03831660005260205260ff6040600020541660001461602c5780600052600260205260406000206001600160a01b03831660005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b9060ff8091169116039060ff821161586f57565b60ff16604d811161586f57600a0a90565b811561659c570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f000000000000000000000000000000000000000000000000000000000000000060ff811690816006146166b0578160061161668557600661660c9161656d565b90604d60ff831611801561666a575b61663357509061662d61557f92616581565b90615d78565b90507fa9cb113d00000000000000000000000000000000000000000000000000000000600052600660045260245260445260646000fd5b5061667482616581565b801561659c5760001904831161661b565b61669090600661656d565b90604d60ff8316116166335750906166aa61557f92616581565b90616592565b505090565b907f00000000000000000000000000000000000000000000000000000000000000009060ff82169060ff8116928284146167915782841161676d57906166fa9161656d565b91604d60ff8416118015616752575b61671c5750509061662d61557f92616581565b9091507fa9cb113d0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b5061675c83616581565b801561659c57600019048411616709565b6167769161656d565b91604d60ff84161161671c575050906166aa61557f92616581565b5050505090565b90816020910312610d8257518015158103610d825790565b80518015616820576020036167e2578051602082810191830183900312610d8257519060ff82116167e2575060ff1690565b612203906040519182917f953576f70000000000000000000000000000000000000000000000000000000083526020600484015260248301906152e6565b50507f000000000000000000000000000000000000000000000000000000000000000090565b90805115616a2c5767ffffffffffffffff8151602083012092169182600052600860205261687b8160056040600020016175c6565b156169e85760005260096020526040600020815167ffffffffffffffff811161544f576168a88254615add565b601f81116169b6575b506020601f821160011461692c5791616906827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea959361691c95600091616921575b506000198260011b9260031b1c19161790565b90556040519182916020835260208301906152e6565b0390a2565b9050840151386168f3565b601f1982169083600052806000209160005b81811061699e57509261691c9492600192827f7d628c9a1796743d365ab521a8b2a4686e419b3269919dc9145ea2ce853b54ea989610616985575b5050811b0190556118c1565b85015160001960f88460031b161c191690553880616979565b9192602060018192868a01518155019401920161693e565b6169e290836000526020600020601f840160051c81019160208510610cad57601f0160051c0190615d8b565b386168b1565b50906122036040519283927f393b8ad200000000000000000000000000000000000000000000000000000000845260048401526040602484015260448301906152e6565b7f14c880ca0000000000000000000000000000000000000000000000000000000060005260046000fd5b67ffffffffffffffff166000818152600760205260409020549092919015616b585791616b5560e092616b2185616aad7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b97616ce9565b846000526008602052616ac4816040600020616fae565b616acd83616ce9565b846000526008602052616ae7836002604060002001616fae565b60405194855260208501906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b60808301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565ba1565b827f1e670e4b0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b616b8e615c11565b506fffffffffffffffffffffffffffffffff6060820151166fffffffffffffffffffffffffffffffff8083511691616beb6020850193616be5616bd863ffffffff87511642615862565b8560808901511690615d78565b90615a41565b80821015616c0457505b16825263ffffffff4216905290565b9050616bf5565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001660208201526020815261557f6040826154ee565b805160005b818110616c5757505050565b6001810180821161586f575b828110616c735750600101616c4b565b6001600160a01b03616c858386615ac9565b51166001600160a01b03616c998387615ac9565b511614616ca857600101616c63565b6001600160a01b03616cba8386615ac9565b51167fa1726e400000000000000000000000000000000000000000000000000000000060005260045260246000fd5b805115616d89576fffffffffffffffffffffffffffffffff6040820151166fffffffffffffffffffffffffffffffff60208301511610616d265750565b606490616d87604051917f8020d12400000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565bfd5b6fffffffffffffffffffffffffffffffff60408201511615801590616e11575b616db05750565b606490616d87604051917fd68af9cc00000000000000000000000000000000000000000000000000000000835260048301906fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b506fffffffffffffffffffffffffffffffff6020820151161515616da9565b6001600160a01b03616eb2911691604092600080855193616e5187866154ee565b602085527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564602086015260208151910182855af13d15616f5b573d91616e9683615511565b92616ea3875194856154ee565b83523d6000602085013e6178b0565b80519081616ebf57505050565b602080616ed0938301019101616798565b15616ed85750565b608490517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152fd5b6060916178b0565b906040519182815491828252602082019060005260206000209260005b818110616f95575050615972925003836154ee565b8454835260019485019487945060209093019201616f80565b7f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19916170e76060928054616feb63ffffffff8260801c1642615862565b9081617126575b50506fffffffffffffffffffffffffffffffff600181602086015116928281541680851060001461711e57508280855b16167fffffffffffffffffffffffffffffffff0000000000000000000000000000000082541617815561709b8651151582907fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff74ff0000000000000000000000000000000000000000835492151560a01b169116179055565b60408601517fffffffffffffffffffffffffffffffff0000000000000000000000000000000060809190911b16939092166fffffffffffffffffffffffffffffffff1692909217910155565b616b5560405180926fffffffffffffffffffffffffffffffff60408092805115158552826020820151166020860152015116910152565b838091617022565b6fffffffffffffffffffffffffffffffff9161715b8392836171546001880154948286169560801c90615d78565b9116615a41565b808210156171da57505b83547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff9290911692909216167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116174260801b73ffffffff00000000000000000000000000000000161781553880616ff2565b9050617165565b67ffffffffffffffff16617202816000526007602052604060002054151590565b1561724c57503360005260136020526040600020541561721e57565b7fd86ad9cf000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fa9902c7e0000000000000000000000000000000000000000000000000000000060005260045260246000fd5b9182549060ff8260a01c16158015617480575b61747a576fffffffffffffffffffffffffffffffff821691600185019081546172d163ffffffff6fffffffffffffffffffffffffffffffff83169360801c1642615862565b90816173dc575b505084811061739d57508383106173325750506173076fffffffffffffffffffffffffffffffff928392615862565b16167fffffffffffffffffffffffffffffffff00000000000000000000000000000000825416179055565b5460801c916173418185615862565b9260001981019080821161586f57617364617369926001600160a01b0396615a41565b616592565b7fd0c8d23a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b82856001600160a01b03927f1a76572a000000000000000000000000000000000000000000000000000000006000526004526024521660445260646000fd5b828692939611617450576173f792616be59160801c90615d78565b8084101561744b5750825b85547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff164260801b73ffffffff00000000000000000000000000000000161786559238806172d8565b617402565b7f9725942a0000000000000000000000000000000000000000000000000000000060005260046000fd5b50505050565b50821561728c565b80548210156157df5760005260206000200190600090565b8054906801000000000000000082101561544f57816174c79160016174de94018155617488565b81939154906000199060031b92831b921b19161790565b9055565b80600052600460205260406000205415600014615f75576175048160036174a0565b600354906000526004602052604060002055600190565b80600052601760205260406000205415600014615f755761753d8160166174a0565b601654906000526017602052604060002055600190565b80600052601360205260406000205415600014615f75576175768160126174a0565b601254906000526013602052604060002055600190565b80600052600760205260406000205415600014615f75576175af8160066174a0565b600654906000526007602052604060002055600190565b600082815260018201602052604090205461602c57806175e8836001936174a0565b80549260005201602052604060002055600190565b805480156176255760001901906176148282617488565b60001982549160031b1b1916905555565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600081815260046020526040902054801561602c57600019810181811161586f5760035490600019820191821161586f578181036176ae575b50505061769a60036175fd565b600052600460205260006040812055600190565b6176d06176bf6174c7936003617488565b90549060031b1c9283926003617488565b9055600052600460205260406000205538808061768d565b600081815260076020526040902054801561602c57600019810181811161586f5760065490600019820191821161586f57818103617742575b50505061772e60066175fd565b600052600760205260006040812055600190565b6177646177536174c7936006617488565b90549060031b1c9283926006617488565b90556000526007602052604060002055388080617721565b9060018201918160005282602052604060002054908115156000146178135760001982019180831161586f578154600019810190811161586f5783816177ca95036177dc575b5050506175fd565b60005260205260006040812055600190565b6177fc6177ec6174c79386617488565b90549060031b1c92839286617488565b9055600052846020526040600020553880806177c2565b50505050600090565b600081815260136020526040902054801561602c57600019810181811161586f5760125490600019820191821161586f57808203617876575b50505061786260126175fd565b600052601360205260006040812055600190565b6178986178876174c7936012617488565b90549060031b1c9283926012617488565b90556000526013602052604060002055388080617855565b9192901561792b57508151156178c4575090565b3b156178cd5790565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152fd5b82519091501561793e5750805190602001fd5b612203906040519182917f08c379a00000000000000000000000000000000000000000000000000000000083526020600484015260248301906152e656fea164736f6c634300081a000aac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b",
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRequiredInboundCCVs(opts *bind.CallOpts, arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRequiredInboundCCVs", arg0, sourceChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredInboundCCVs(&_SiloedUSDCTokenPool.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRequiredInboundCCVs(arg0 common.Address, sourceChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredInboundCCVs(&_SiloedUSDCTokenPool.CallOpts, arg0, sourceChainSelector, arg2, arg3, arg4)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRequiredOutboundCCVs(opts *bind.CallOpts, arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRequiredOutboundCCVs", arg0, destChainSelector, arg2, arg3, arg4)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredOutboundCCVs(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRequiredOutboundCCVs(arg0 common.Address, destChainSelector uint64, arg2 *big.Int, arg3 uint16, arg4 []byte) ([]common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRequiredOutboundCCVs(&_SiloedUSDCTokenPool.CallOpts, arg0, destChainSelector, arg2, arg3, arg4)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiloedUSDCTokenPool.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) GetRouter() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRouter(&_SiloedUSDCTokenPool.CallOpts)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolCallerSession) GetRouter() (common.Address, error) {
	return _SiloedUSDCTokenPool.Contract.GetRouter(&_SiloedUSDCTokenPool.CallOpts)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RenounceRole(&_SiloedUSDCTokenPool.TransactOpts, role, callerConfirmation)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.RenounceRole(&_SiloedUSDCTokenPool.TransactOpts, role, callerConfirmation)
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

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setRebalancer", newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRebalancer(&_SiloedUSDCTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetRebalancer(newRebalancer common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRebalancer(&_SiloedUSDCTokenPool.TransactOpts, newRebalancer)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.contract.Transact(opts, "setRouter", newRouter)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRouter(&_SiloedUSDCTokenPool.TransactOpts, newRouter)
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _SiloedUSDCTokenPool.Contract.SetRouter(&_SiloedUSDCTokenPool.TransactOpts, newRouter)
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
	RemoteChainSelector uint64
	OutboundCCVs        []common.Address
	InboundCCVs         []common.Address
	Raw                 types.Log
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

type SiloedUSDCTokenPoolRouterUpdatedIterator struct {
	Event *SiloedUSDCTokenPoolRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *SiloedUSDCTokenPoolRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiloedUSDCTokenPoolRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(SiloedUSDCTokenPoolRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *SiloedUSDCTokenPoolRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *SiloedUSDCTokenPoolRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type SiloedUSDCTokenPoolRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolRouterUpdatedIterator, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &SiloedUSDCTokenPoolRouterUpdatedIterator{contract: _SiloedUSDCTokenPool.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _SiloedUSDCTokenPool.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(SiloedUSDCTokenPoolRouterUpdated)
				if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_SiloedUSDCTokenPool *SiloedUSDCTokenPoolFilterer) ParseRouterUpdated(log types.Log) (*SiloedUSDCTokenPoolRouterUpdated, error) {
	event := new(SiloedUSDCTokenPoolRouterUpdated)
	if err := _SiloedUSDCTokenPool.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
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
	return common.HexToHash("0xb0897119e8510f887b892cbc4c8506fc51d9849fd90afae4fd065e705f2d0f6c")
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

func (SiloedUSDCTokenPoolOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (SiloedUSDCTokenPoolOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
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

func (SiloedUSDCTokenPoolRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
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
	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	RATELIMITERADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	GetAccumulatedFees(opts *bind.CallOpts) (*big.Int, error)

	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetAvailableTokens(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetChainRebalancer(opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentProposedCCTPChainMigration(opts *bind.CallOpts) (uint64, error)

	GetExcludedTokensByChain(opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error)

	GetRebalancer(opts *bind.CallOpts) (common.Address, error)

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

	GetUnsiloedLiquidity(opts *bind.CallOpts) (*big.Int, error)

	HasRateLimitAdminRole(opts *bind.CallOpts, account common.Address) (bool, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

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

	ApplyFinalityConfigUpdates(opts *bind.TransactOpts, finalityThreshold uint16, customFinalityTransferFeeBps uint16, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	ApplyTokenTransferFeeConfigUpdates(opts *bind.TransactOpts, tokenTransferFeeConfigArgs []TokenPoolTokenTransferFeeConfigArgs, destToUseDefaultFeeConfigs []uint64) (*types.Transaction, error)

	BurnLockedUSDC(opts *bind.TransactOpts) (*types.Transaction, error)

	CancelExistingCCTPMigrationProposal(opts *bind.TransactOpts) (*types.Transaction, error)

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

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RevokeRateLimitAdminRole(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetChainRateLimiterConfigs(opts *bind.TransactOpts, remoteChainSelectors []uint64, outboundConfigs []RateLimiterConfig, inboundConfigs []RateLimiterConfig) (*types.Transaction, error)

	SetCircleMigratorAddress(opts *bind.TransactOpts, migrator common.Address) (*types.Transaction, error)

	SetCustomFinalityRateLimitConfig(opts *bind.TransactOpts, rateLimitConfigArgs []TokenPoolCustomFinalityRateLimitConfigArgs) (*types.Transaction, error)

	SetRebalancer(opts *bind.TransactOpts, newRebalancer common.Address) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

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

	FilterCustomFinalityOutboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumedIterator, error)

	WatchCustomFinalityOutboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityOutboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomFinalityOutboundRateLimitConsumed, error)

	FilterCustomFinalityTransferInboundRateLimitConsumed(opts *bind.FilterOpts, remoteChainSelector []uint64) (*SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumedIterator, error)

	WatchCustomFinalityTransferInboundRateLimitConsumed(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, remoteChainSelector []uint64) (event.Subscription, error)

	ParseCustomFinalityTransferInboundRateLimitConsumed(log types.Log) (*SiloedUSDCTokenPoolCustomFinalityTransferInboundRateLimitConsumed, error)

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

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SiloedUSDCTokenPoolOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*SiloedUSDCTokenPoolOwnershipTransferred, error)

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

	FilterRouterUpdated(opts *bind.FilterOpts) (*SiloedUSDCTokenPoolRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *SiloedUSDCTokenPoolRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*SiloedUSDCTokenPoolRouterUpdated, error)

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
